// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

// Package blindfold provides F5XC Secret Management encryption utilities.
package blindfold

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"net/http"
	"os"
	"time"

	"golang.org/x/crypto/pkcs12"
)

// AuthConfig holds configuration for F5XC API authentication.
type AuthConfig struct {
	// APIToken is the API token for bearer authentication.
	APIToken string
	// P12File is the path to the P12 certificate file.
	P12File string
	// P12Password is the password for the P12 certificate.
	P12Password string
	// BaseURL is the F5XC API base URL.
	BaseURL string
}

// AuthMethod represents the authentication method being used.
type AuthMethod int

const (
	// AuthMethodNone indicates no authentication is configured.
	AuthMethodNone AuthMethod = iota
	// AuthMethodToken indicates API token authentication.
	AuthMethodToken
	// AuthMethodP12 indicates P12 certificate authentication.
	AuthMethodP12
)

// String returns a human-readable name for the authentication method.
func (m AuthMethod) String() string {
	switch m {
	case AuthMethodToken:
		return "API Token"
	case AuthMethodP12:
		return "P12 Certificate"
	default:
		return "None"
	}
}

// AuthResult contains the HTTP client and metadata for authenticated requests.
type AuthResult struct {
	// Client is the configured HTTP client.
	Client *http.Client
	// Method indicates which authentication method was used.
	Method AuthMethod
	// BaseURL is the F5XC API base URL.
	BaseURL string
	// Token is the API token (only set for token auth).
	Token string
}

// Environment variable names for F5XC authentication.
// Using F5XC_* prefix for F5 Distributed Cloud branding.
const (
	EnvAPIURL      = "F5XC_API_URL"
	EnvAPIToken    = "F5XC_API_TOKEN"
	EnvP12File     = "F5XC_P12_FILE"
	EnvP12Password = "F5XC_P12_PASSWORD" // pragma: allowlist secret
)

// DefaultAPIURL is the default F5XC API URL.
const DefaultAPIURL = "https://console.ves.volterra.io"

// DefaultTimeout is the default HTTP client timeout.
const DefaultTimeout = 30 * time.Second

// GetAuthConfigFromEnv reads authentication configuration from environment variables.
// It checks for API token first, then falls back to P12 certificate authentication.
func GetAuthConfigFromEnv() (*AuthConfig, error) {
	config := &AuthConfig{
		APIToken:    os.Getenv(EnvAPIToken),
		P12File:     os.Getenv(EnvP12File),
		P12Password: os.Getenv(EnvP12Password),
		BaseURL:     os.Getenv(EnvAPIURL),
	}

	if config.BaseURL == "" {
		config.BaseURL = DefaultAPIURL
	}

	// Validate that at least one auth method is configured
	if config.APIToken == "" && config.P12File == "" {
		return nil, fmt.Errorf(
			"no F5XC authentication configured: set either %s for API token authentication "+
				"or %s and %s for P12 certificate authentication",
			EnvAPIToken, EnvP12File, EnvP12Password,
		)
	}

	return config, nil
}

// CreateAuthenticatedClient creates an HTTP client configured for F5XC API authentication.
// It prioritizes API token authentication over P12 certificate authentication.
//
// Priority:
//  1. API token (F5XC_API_TOKEN)
//  2. P12 certificate (F5XC_P12_FILE + F5XC_P12_PASSWORD)
//
// Returns an AuthResult containing the client and authentication metadata.
func CreateAuthenticatedClient(config *AuthConfig) (*AuthResult, error) {
	if config == nil {
		return nil, fmt.Errorf("authentication configuration is required")
	}

	// Priority 1: API Token authentication
	if config.APIToken != "" {
		client := &http.Client{
			Timeout: DefaultTimeout,
		}
		return &AuthResult{
			Client:  client,
			Method:  AuthMethodToken,
			BaseURL: config.BaseURL,
			Token:   config.APIToken,
		}, nil
	}

	// Priority 2: P12 Certificate authentication
	if config.P12File != "" {
		client, err := createHTTPClientFromP12(config.P12File, config.P12Password)
		if err != nil {
			return nil, fmt.Errorf("failed to create P12 authenticated client: %w", err)
		}
		return &AuthResult{
			Client:  client,
			Method:  AuthMethodP12,
			BaseURL: config.BaseURL,
		}, nil
	}

	return nil, fmt.Errorf(
		"no authentication method available: configure %s or %s",
		EnvAPIToken, EnvP12File,
	)
}

// createHTTPClientFromP12 creates an HTTP client configured with P12 certificate authentication.
// This implementation handles P12 files with certificate chains (multiple certificates).
func createHTTPClientFromP12(p12File, password string) (*http.Client, error) {
	p12Data, err := os.ReadFile(p12File)
	if err != nil {
		return nil, fmt.Errorf("failed to read P12 file %q: %w", p12File, err)
	}

	// Use ToPEM to handle P12 files with certificate chains
	blocks, err := pkcs12.ToPEM(p12Data, password)
	if err != nil {
		return nil, fmt.Errorf("failed to decode P12 file: %w", err)
	}

	var pemKey, pemCert []byte
	var caCerts []*x509.Certificate

	for _, block := range blocks {
		switch block.Type {
		case "PRIVATE KEY":
			pemKey = append(pemKey, pem.EncodeToMemory(block)...)
		case "CERTIFICATE":
			cert, err := x509.ParseCertificate(block.Bytes)
			if err != nil {
				return nil, fmt.Errorf("failed to parse certificate: %w", err)
			}
			// Distinguish between CA certificates and client certificate
			if cert.IsCA {
				caCerts = append(caCerts, cert)
			} else {
				pemCert = append(pemCert, pem.EncodeToMemory(block)...)
			}
		}
	}

	if len(pemKey) == 0 || len(pemCert) == 0 {
		return nil, fmt.Errorf("P12 file missing required key or certificate")
	}

	tlsCert, err := tls.X509KeyPair(pemCert, pemKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create X509 key pair: %w", err)
	}

	// Create CA pool starting with system CAs, then add P12 CAs
	caCertPool, err := x509.SystemCertPool()
	if err != nil {
		// Fall back to empty pool if system certs unavailable
		caCertPool = x509.NewCertPool()
	}
	for _, caCert := range caCerts {
		caCertPool.AddCert(caCert)
	}

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			Certificates: []tls.Certificate{tlsCert},
			RootCAs:      caCertPool,
			MinVersion:   tls.VersionTLS12,
		},
	}

	return &http.Client{
		Timeout:   DefaultTimeout,
		Transport: transport,
	}, nil
}

// SetAuthorizationHeader sets the appropriate authorization header on the request
// based on the authentication method. For token auth, it adds the Bearer token.
// For P12 auth, no header is needed as the certificate is used at the TLS layer.
func (r *AuthResult) SetAuthorizationHeader(req *http.Request) {
	if r.Method == AuthMethodToken && r.Token != "" {
		req.Header.Set("Authorization", "APIToken "+r.Token)
	}
	// P12 authentication uses client certificates at the TLS layer,
	// so no Authorization header is needed.
}

// NewRequest creates a new HTTP request with the appropriate authentication.
func (r *AuthResult) NewRequest(method, url string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	r.SetAuthorizationHeader(req)
	return req, nil
}

// ValidateConfig validates the authentication configuration.
func ValidateConfig(config *AuthConfig) error {
	if config == nil {
		return fmt.Errorf("configuration is nil")
	}

	if config.BaseURL == "" {
		return fmt.Errorf("base URL is required")
	}

	// Check for valid auth method
	hasToken := config.APIToken != ""
	hasP12 := config.P12File != ""

	if !hasToken && !hasP12 {
		return fmt.Errorf("no authentication method configured")
	}

	// If P12 is configured, validate the file exists
	if hasP12 {
		if _, err := os.Stat(config.P12File); os.IsNotExist(err) {
			return fmt.Errorf("P12 file does not exist: %s", config.P12File)
		}
	}

	return nil
}
