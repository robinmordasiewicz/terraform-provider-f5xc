// This file is MANUALLY MAINTAINED and is NOT auto-generated.

package blindfold

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"time"
)

const (
	// PublicKeyEndpoint is the F5XC API endpoint for fetching the encryption public key
	PublicKeyEndpoint = "/api/secret_management/get_public_key"

	// Retry configuration for rate-limited requests
	defaultMaxRetries     = 3
	defaultRetryWaitMin   = 1 * time.Second
	defaultRetryWaitMax   = 30 * time.Second
	defaultRateLimitDelay = 60 * time.Second
)

// GetPublicKey fetches the encryption public key from F5XC Secret Management API.
// The public key is used for RSA-OAEP encryption of secrets.
//
// Parameters:
//   - ctx: Context for cancellation and timeouts
//   - httpClient: Configured HTTP client with authentication
//   - baseURL: F5XC API base URL (e.g., "https://tenant.console.ves.volterra.io")
//
// Returns:
//   - PublicKey containing RSA modulus and exponent
//   - Error if the request fails or response is invalid
func GetPublicKey(ctx context.Context, httpClient *http.Client, baseURL string) (*PublicKey, error) {
	return GetPublicKeyWithVersion(ctx, httpClient, baseURL, 0)
}

// isRetryableStatus returns true if the HTTP status code indicates a retryable error
func isRetryableStatus(statusCode int) bool {
	switch statusCode {
	case http.StatusTooManyRequests, // 429 - Rate limited
		http.StatusServiceUnavailable, // 503
		http.StatusGatewayTimeout,     // 504
		http.StatusBadGateway:         // 502
		return true
	default:
		return false
	}
}

// calculateBackoff returns the backoff duration for a given attempt using exponential backoff
func calculateBackoff(attempt int) time.Duration {
	// Exponential backoff: min * 2^attempt, capped at max
	backoff := float64(defaultRetryWaitMin) * math.Pow(2, float64(attempt))
	if backoff > float64(defaultRetryWaitMax) {
		backoff = float64(defaultRetryWaitMax)
	}
	return time.Duration(backoff)
}

// GetPublicKeyWithVersion fetches a specific version of the encryption public key.
// Use version 0 or negative to get the current/latest key.
//
// This function includes retry logic with exponential backoff for handling
// rate-limited (429) and temporary server errors (502, 503, 504).
//
// Parameters:
//   - ctx: Context for cancellation and timeouts
//   - httpClient: Configured HTTP client with authentication
//   - baseURL: F5XC API base URL
//   - version: Specific key version to fetch, or 0 for latest
//
// Returns:
//   - PublicKey containing RSA modulus and exponent
//   - Error if the request fails after all retries or response is invalid
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

	var lastErr error

	for attempt := 0; attempt <= defaultMaxRetries; attempt++ {
		// Check context before making request
		if ctx.Err() != nil {
			return nil, fmt.Errorf("context cancelled: %w", ctx.Err())
		}

		// Create request for each attempt (request body may be consumed)
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create request: %w", err)
		}
		req.Header.Set("Accept", "application/json")

		// Execute request
		resp, err := httpClient.Do(req)
		if err != nil {
			lastErr = fmt.Errorf("request failed: %w", err)
			// Network errors are retryable
			if attempt < defaultMaxRetries {
				select {
				case <-ctx.Done():
					return nil, fmt.Errorf("context cancelled during retry: %w", ctx.Err())
				case <-time.After(calculateBackoff(attempt)):
					continue
				}
			}
			return nil, lastErr
		}

		// Read response body
		body, err := io.ReadAll(resp.Body)
		_ = resp.Body.Close()
		if err != nil {
			return nil, fmt.Errorf("failed to read response: %w", err)
		}

		// Check for successful response
		if resp.StatusCode == http.StatusOK {
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

		// Check if error is retryable
		lastErr = fmt.Errorf("unexpected status %d: %s", resp.StatusCode, string(body))
		if !isRetryableStatus(resp.StatusCode) || attempt >= defaultMaxRetries {
			return nil, lastErr
		}

		// Calculate backoff (use longer delay for rate limiting)
		backoff := calculateBackoff(attempt)
		if resp.StatusCode == http.StatusTooManyRequests {
			backoff = defaultRateLimitDelay
		}

		// Wait before retry
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("context cancelled during retry: %w", ctx.Err())
		case <-time.After(backoff):
			continue
		}
	}

	return nil, lastErr
}
