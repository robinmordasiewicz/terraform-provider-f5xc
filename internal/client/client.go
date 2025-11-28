// Copyright (c) F5XC Community
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"time"

	"golang.org/x/crypto/pkcs12"

	f5xcerrors "github.com/f5xc/terraform-provider-f5xc/internal/errors"
)

// Client configuration defaults
const (
	DefaultTimeout        = 30 * time.Second
	DefaultMaxRetries     = 3
	DefaultRetryWaitMin   = 1 * time.Second
	DefaultRetryWaitMax   = 30 * time.Second
	DefaultRateLimitDelay = 60 * time.Second
)

// AuthType represents the authentication method used by the client
type AuthType int

const (
	// AuthTypeToken uses API token authentication
	AuthTypeToken AuthType = iota
	// AuthTypeCertificate uses TLS client certificate authentication
	AuthTypeCertificate
)

// Client manages communication with the F5 Distributed Cloud API
type Client struct {
	BaseURL      string
	APIToken     string
	AuthType     AuthType
	HTTPClient   *http.Client
	MaxRetries   int
	RetryWaitMin time.Duration
	RetryWaitMax time.Duration
}

// ClientOption allows customizing the client
type ClientOption func(*Client)

// WithTimeout sets the HTTP client timeout
func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) {
		c.HTTPClient.Timeout = timeout
	}
}

// WithMaxRetries sets the maximum number of retries
func WithMaxRetries(retries int) ClientOption {
	return func(c *Client) {
		c.MaxRetries = retries
	}
}

// WithRetryWait sets the retry wait bounds
func WithRetryWait(min, max time.Duration) ClientOption {
	return func(c *Client) {
		c.RetryWaitMin = min
		c.RetryWaitMax = max
	}
}

// WithTLSConfig sets a custom TLS configuration on the HTTP client
func WithTLSConfig(tlsConfig *tls.Config) ClientOption {
	return func(c *Client) {
		c.HTTPClient.Transport = &http.Transport{
			TLSClientConfig: tlsConfig,
		}
		c.AuthType = AuthTypeCertificate
	}
}

// LoadP12Certificate loads a PKCS#12 certificate file and returns a TLS config
func LoadP12Certificate(p12File, password string) (*tls.Config, error) {
	p12Data, err := os.ReadFile(p12File)
	if err != nil {
		return nil, fmt.Errorf("failed to read P12 file: %w", err)
	}

	blocks, err := pkcs12.ToPEM(p12Data, password)
	if err != nil {
		return nil, fmt.Errorf("failed to decode P12 file: %w", err)
	}

	var pemKey, pemCert []byte
	var caCerts []*x509.Certificate

	for _, block := range blocks {
		switch block.Type {
		case "PRIVATE KEY":
			// Use pem.EncodeToMemory to properly base64-encode the DER data
			pemKey = append(pemKey, pem.EncodeToMemory(block)...)
		case "CERTIFICATE":
			cert, err := x509.ParseCertificate(block.Bytes)
			if err != nil {
				return nil, fmt.Errorf("failed to parse certificate: %w", err)
			}
			// Check if this is a CA certificate or the client certificate
			if cert.IsCA {
				caCerts = append(caCerts, cert)
			} else {
				// Use pem.EncodeToMemory to properly base64-encode the DER data
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

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{tlsCert},
		RootCAs:      caCertPool,
		MinVersion:   tls.VersionTLS12,
	}

	return tlsConfig, nil
}

// LoadCertificateKeyPair loads PEM-encoded certificate and key files
func LoadCertificateKeyPair(certFile, keyFile, caFile string) (*tls.Config, error) {
	tlsCert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, fmt.Errorf("failed to load certificate/key pair: %w", err)
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{tlsCert},
		MinVersion:   tls.VersionTLS12,
	}

	if caFile != "" {
		caCert, err := os.ReadFile(caFile)
		if err != nil {
			return nil, fmt.Errorf("failed to read CA certificate: %w", err)
		}
		caCertPool := x509.NewCertPool()
		if !caCertPool.AppendCertsFromPEM(caCert) {
			return nil, fmt.Errorf("failed to parse CA certificate")
		}
		tlsConfig.RootCAs = caCertPool
	}

	return tlsConfig, nil
}

// NewClient creates a new F5 Distributed Cloud API client with token authentication
func NewClient(baseURL, apiToken string, opts ...ClientOption) *Client {
	c := &Client{
		BaseURL:      baseURL,
		APIToken:     apiToken,
		AuthType:     AuthTypeToken,
		MaxRetries:   DefaultMaxRetries,
		RetryWaitMin: DefaultRetryWaitMin,
		RetryWaitMax: DefaultRetryWaitMax,
		HTTPClient: &http.Client{
			Timeout: DefaultTimeout,
		},
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

// NewClientWithP12 creates a new F5 Distributed Cloud API client with P12 certificate authentication
func NewClientWithP12(baseURL, p12File, p12Password string, opts ...ClientOption) (*Client, error) {
	tlsConfig, err := LoadP12Certificate(p12File, p12Password)
	if err != nil {
		return nil, err
	}

	c := &Client{
		BaseURL:      baseURL,
		AuthType:     AuthTypeCertificate,
		MaxRetries:   DefaultMaxRetries,
		RetryWaitMin: DefaultRetryWaitMin,
		RetryWaitMax: DefaultRetryWaitMax,
		HTTPClient: &http.Client{
			Timeout: DefaultTimeout,
			Transport: &http.Transport{
				TLSClientConfig: tlsConfig,
			},
		},
	}

	for _, opt := range opts {
		opt(c)
	}

	return c, nil
}

// NewClientWithCert creates a new F5 Distributed Cloud API client with PEM certificate authentication
func NewClientWithCert(baseURL, certFile, keyFile, caFile string, opts ...ClientOption) (*Client, error) {
	tlsConfig, err := LoadCertificateKeyPair(certFile, keyFile, caFile)
	if err != nil {
		return nil, err
	}

	c := &Client{
		BaseURL:      baseURL,
		AuthType:     AuthTypeCertificate,
		MaxRetries:   DefaultMaxRetries,
		RetryWaitMin: DefaultRetryWaitMin,
		RetryWaitMax: DefaultRetryWaitMax,
		HTTPClient: &http.Client{
			Timeout: DefaultTimeout,
			Transport: &http.Transport{
				TLSClientConfig: tlsConfig,
			},
		},
	}

	for _, opt := range opts {
		opt(c)
	}

	return c, nil
}

// Metadata represents common metadata fields
type Metadata struct {
	Name        string            `json:"name"`
	Namespace   string            `json:"namespace,omitempty"`
	Labels      map[string]string `json:"labels,omitempty"`
	Annotations map[string]string `json:"annotations,omitempty"`
	Description string            `json:"description,omitempty"`
	Disable     bool              `json:"disable,omitempty"`
	UID         string            `json:"uid,omitempty"`
}

// isRetryableStatus returns true if the HTTP status code indicates a retryable error
func isRetryableStatus(statusCode int) bool {
	switch statusCode {
	case http.StatusTooManyRequests, // 429 - Rate limited
		http.StatusServiceUnavailable,   // 503
		http.StatusGatewayTimeout,       // 504
		http.StatusBadGateway:           // 502
		return true
	default:
		return statusCode >= 500
	}
}

// calculateBackoff returns the backoff duration for a given attempt
func (c *Client) calculateBackoff(attempt int) time.Duration {
	// Exponential backoff: min * 2^attempt, capped at max
	backoff := float64(c.RetryWaitMin) * math.Pow(2, float64(attempt))
	if backoff > float64(c.RetryWaitMax) {
		backoff = float64(c.RetryWaitMax)
	}
	return time.Duration(backoff)
}

// doRequest performs an HTTP request with retry logic and returns the response
func (c *Client) doRequest(ctx context.Context, method, path string, body interface{}) ([]byte, error) {
	var bodyBytes []byte
	if body != nil {
		var err error
		bodyBytes, err = json.Marshal(body)
		if err != nil {
			return nil, f5xcerrors.WrapError(err, "request", "marshal")
		}
	}

	url := c.BaseURL + path
	var lastErr error

	for attempt := 0; attempt <= c.MaxRetries; attempt++ {
		// Check context before making request
		if ctx.Err() != nil {
			return nil, f5xcerrors.NewTimeoutError("request", method+" "+path, ctx.Err())
		}

		var reqBody io.Reader
		if bodyBytes != nil {
			reqBody = bytes.NewBuffer(bodyBytes)
		}

		req, err := http.NewRequestWithContext(ctx, method, url, reqBody)
		if err != nil {
			return nil, f5xcerrors.WrapError(err, "request", "create")
		}

		// Only set Authorization header for token-based authentication
		// Certificate-based authentication uses TLS client certificates instead
		if c.AuthType == AuthTypeToken && c.APIToken != "" {
			req.Header.Set("Authorization", "APIToken "+c.APIToken)
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/json")

		resp, err := c.HTTPClient.Do(req)
		if err != nil {
			lastErr = f5xcerrors.NewNetworkError(err)
			// Network errors are retryable
			if attempt < c.MaxRetries {
				select {
				case <-ctx.Done():
					return nil, f5xcerrors.NewTimeoutError("request", method+" "+path, ctx.Err())
				case <-time.After(c.calculateBackoff(attempt)):
					continue
				}
			}
			return nil, lastErr
		}

		respBody, err := io.ReadAll(resp.Body)
		_ = resp.Body.Close() // Error intentionally ignored - body already read, cleanup only
		if err != nil {
			return nil, f5xcerrors.WrapError(err, "response", "read")
		}

		// Success
		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			return respBody, nil
		}

		// Create structured error
		apiErr := f5xcerrors.NewAPIError(resp.StatusCode, respBody, path, method)
		lastErr = apiErr

		// Check if error is retryable
		if !isRetryableStatus(resp.StatusCode) || attempt >= c.MaxRetries {
			return nil, apiErr
		}

		// Calculate backoff (use longer delay for rate limiting)
		backoff := c.calculateBackoff(attempt)
		if resp.StatusCode == http.StatusTooManyRequests {
			backoff = DefaultRateLimitDelay
		}

		select {
		case <-ctx.Done():
			return nil, f5xcerrors.NewTimeoutError("request", method+" "+path, ctx.Err())
		case <-time.After(backoff):
			// Continue to next retry
		}
	}

	return nil, lastErr
}

// Get performs a GET request
func (c *Client) Get(ctx context.Context, path string, result interface{}) error {
	body, err := c.doRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return err
	}
	if result != nil && len(body) > 0 {
		return json.Unmarshal(body, result)
	}
	return nil
}

// Post performs a POST request
func (c *Client) Post(ctx context.Context, path string, data, result interface{}) error {
	body, err := c.doRequest(ctx, http.MethodPost, path, data)
	if err != nil {
		return err
	}
	if result != nil && len(body) > 0 {
		return json.Unmarshal(body, result)
	}
	return nil
}

// Put performs a PUT request
func (c *Client) Put(ctx context.Context, path string, data, result interface{}) error {
	body, err := c.doRequest(ctx, http.MethodPut, path, data)
	if err != nil {
		return err
	}
	if result != nil && len(body) > 0 {
		return json.Unmarshal(body, result)
	}
	return nil
}

// Delete performs a DELETE request
func (c *Client) Delete(ctx context.Context, path string) error {
	_, err := c.doRequest(ctx, http.MethodDelete, path, nil)
	return err
}
