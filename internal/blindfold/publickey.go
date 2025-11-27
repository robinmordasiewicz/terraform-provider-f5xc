// This file is MANUALLY MAINTAINED and is NOT auto-generated.

package blindfold

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	// PublicKeyEndpoint is the F5XC API endpoint for fetching the encryption public key
	PublicKeyEndpoint = "/api/secret_management/get_public_key"
)

// GetPublicKey fetches the encryption public key from F5XC Secret Management API.
// The public key is used for RSA-OAEP encryption of secrets.
//
// Parameters:
//   - ctx: Context for cancellation and timeouts
//   - httpClient: Configured HTTP client with authentication
//   - baseURL: F5XC API base URL (e.g., "https://tenant.console.ves.volterra.io/api")
//
// Returns:
//   - PublicKey containing RSA modulus and exponent
//   - Error if the request fails or response is invalid
func GetPublicKey(ctx context.Context, httpClient *http.Client, baseURL string) (*PublicKey, error) {
	return GetPublicKeyWithVersion(ctx, httpClient, baseURL, 0)
}

// GetPublicKeyWithVersion fetches a specific version of the encryption public key.
// Use version 0 or negative to get the current/latest key.
//
// Parameters:
//   - ctx: Context for cancellation and timeouts
//   - httpClient: Configured HTTP client with authentication
//   - baseURL: F5XC API base URL
//   - version: Specific key version to fetch, or 0 for latest
//
// Returns:
//   - PublicKey containing RSA modulus and exponent
//   - Error if the request fails or response is invalid
func GetPublicKeyWithVersion(ctx context.Context, httpClient *http.Client, baseURL string, version int) (*PublicKey, error) {
	if httpClient == nil {
		return nil, fmt.Errorf("HTTP client is required")
	}
	if baseURL == "" {
		return nil, fmt.Errorf("base URL is required")
	}

	// Build URL with optional version parameter
	url := baseURL + PublicKeyEndpoint
	if version > 0 {
		url = fmt.Sprintf("%s?key_version=%d", url, version)
	}

	// Create request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Accept", "application/json")

	// Execute request
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	// Check status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status %d: %s", resp.StatusCode, string(body))
	}

	// Parse response - F5XC wraps responses in an envelope
	var envelope APIEnvelope[PublicKey]
	if err := json.Unmarshal(body, &envelope); err != nil {
		// Try parsing without envelope (direct response)
		var pubKey PublicKey
		if err2 := json.Unmarshal(body, &pubKey); err2 != nil {
			return nil, fmt.Errorf("failed to parse response: %w (envelope: %v)", err2, err)
		}
		return &pubKey, nil
	}

	// Validate the response
	if envelope.Data.ModulusBase64 == "" {
		return nil, fmt.Errorf("response missing modulus")
	}
	if envelope.Data.PublicExponentBase64 == "" {
		return nil, fmt.Errorf("response missing public exponent")
	}

	return &envelope.Data, nil
}
