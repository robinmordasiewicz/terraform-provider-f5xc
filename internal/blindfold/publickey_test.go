// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package blindfold

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetPublicKey_MockServer(t *testing.T) {
	tests := []struct {
		name           string
		responseCode   int
		responseBody   interface{}
		wantErr        bool
		wantErrContain string
		wantKeyVersion int
		wantTenant     string
	}{
		{
			name:         "successful response with envelope",
			responseCode: http.StatusOK,
			responseBody: APIEnvelope[PublicKey]{
				Data: PublicKey{
					KeyVersion:           1,
					ModulusBase64:        "dGVzdC1tb2R1bHVz", // base64 of "test-modulus"
					PublicExponentBase64: "AQAB",             // base64 of 65537
					Tenant:               "test-tenant",
				},
			},
			wantErr:        false,
			wantKeyVersion: 1,
			wantTenant:     "test-tenant",
		},
		{
			name:         "successful response with data wrapper",
			responseCode: http.StatusOK,
			responseBody: map[string]interface{}{
				"data": map[string]interface{}{
					"key_version":            3,
					"modulus_base64":         "dGVzdC1tb2R1bHVz",
					"public_exponent_base64": "AQAB",
					"tenant":                 "wrapped-tenant",
				},
			},
			wantErr:        false,
			wantKeyVersion: 3,
			wantTenant:     "wrapped-tenant",
		},
		{
			name:           "server error",
			responseCode:   http.StatusInternalServerError,
			responseBody:   map[string]string{"error": "internal error"},
			wantErr:        true,
			wantErrContain: "unexpected status 500",
		},
		{
			name:           "unauthorized",
			responseCode:   http.StatusUnauthorized,
			responseBody:   map[string]string{"error": "unauthorized"},
			wantErr:        true,
			wantErrContain: "unexpected status 401",
		},
		{
			name:         "missing modulus",
			responseCode: http.StatusOK,
			responseBody: APIEnvelope[PublicKey]{
				Data: PublicKey{
					KeyVersion:           1,
					PublicExponentBase64: "AQAB",
					Tenant:               "test-tenant",
				},
			},
			wantErr:        true,
			wantErrContain: "missing modulus",
		},
		{
			name:         "missing exponent",
			responseCode: http.StatusOK,
			responseBody: APIEnvelope[PublicKey]{
				Data: PublicKey{
					KeyVersion:    1,
					ModulusBase64: "dGVzdC1tb2R1bHVz",
					Tenant:        "test-tenant",
				},
			},
			wantErr:        true,
			wantErrContain: "missing public exponent",
		},
		{
			name:           "invalid JSON response",
			responseCode:   http.StatusOK,
			responseBody:   "not json",
			wantErr:        true,
			wantErrContain: "failed to parse response",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// Verify request method
				if r.Method != http.MethodGet {
					t.Errorf("expected GET, got %s", r.Method)
				}

				// Verify endpoint
				if r.URL.Path != PublicKeyEndpoint {
					t.Errorf("expected path %s, got %s", PublicKeyEndpoint, r.URL.Path)
				}

				// Verify Accept header
				if r.Header.Get("Accept") != "application/json" {
					t.Errorf("expected Accept: application/json, got %s", r.Header.Get("Accept"))
				}

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(tt.responseCode)

				var body []byte
				switch v := tt.responseBody.(type) {
				case string:
					body = []byte(v)
				default:
					body, _ = json.Marshal(v)
				}
				_, _ = w.Write(body)
			}))
			defer server.Close()

			client := &http.Client{Timeout: 5 * time.Second}
			ctx := context.Background()

			pubKey, err := GetPublicKey(ctx, client, server.URL)

			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				if tt.wantErrContain != "" && !contains(err.Error(), tt.wantErrContain) {
					t.Errorf("error %q should contain %q", err.Error(), tt.wantErrContain)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if pubKey.KeyVersion != tt.wantKeyVersion {
				t.Errorf("KeyVersion = %v, want %v", pubKey.KeyVersion, tt.wantKeyVersion)
			}
			if pubKey.Tenant != tt.wantTenant {
				t.Errorf("Tenant = %v, want %v", pubKey.Tenant, tt.wantTenant)
			}
		})
	}
}

func TestGetPublicKeyWithVersion_MockServer(t *testing.T) {
	tests := []struct {
		name           string
		version        int
		wantQueryParam string
		wantHasVersion bool
		wantVersionNum int
	}{
		{
			name:           "no version (current key)",
			version:        0,
			wantQueryParam: "",
			wantHasVersion: false,
		},
		{
			name:           "specific version",
			version:        5,
			wantQueryParam: "key_version=5",
			wantHasVersion: true,
			wantVersionNum: 5,
		},
		{
			name:           "negative version (treated as current)",
			version:        -1,
			wantQueryParam: "",
			wantHasVersion: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var receivedQuery string
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				receivedQuery = r.URL.RawQuery

				response := APIEnvelope[PublicKey]{
					Data: PublicKey{
						KeyVersion:           tt.version,
						ModulusBase64:        "dGVzdC1tb2R1bHVz",
						PublicExponentBase64: "AQAB",
						Tenant:               "test-tenant",
					},
				}

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(response)
			}))
			defer server.Close()

			client := &http.Client{Timeout: 5 * time.Second}
			ctx := context.Background()

			_, err := GetPublicKeyWithVersion(ctx, client, server.URL, tt.version)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if tt.wantHasVersion {
				if receivedQuery != tt.wantQueryParam {
					t.Errorf("query = %q, want %q", receivedQuery, tt.wantQueryParam)
				}
			} else {
				if receivedQuery != "" {
					t.Errorf("expected no query params, got %q", receivedQuery)
				}
			}
		})
	}
}

func TestGetPublicKey_NilClient(t *testing.T) {
	ctx := context.Background()
	_, err := GetPublicKey(ctx, nil, "https://example.com")
	if err == nil {
		t.Fatal("expected error for nil client")
	}
	if !contains(err.Error(), "HTTP client is required") {
		t.Errorf("error should mention HTTP client, got: %v", err)
	}
}

func TestGetPublicKey_EmptyBaseURL(t *testing.T) {
	ctx := context.Background()
	client := &http.Client{Timeout: 5 * time.Second}
	_, err := GetPublicKey(ctx, client, "")
	if err == nil {
		t.Fatal("expected error for empty base URL")
	}
	if !contains(err.Error(), "base URL is required") {
		t.Errorf("error should mention base URL, got: %v", err)
	}
}

func TestGetPublicKey_ContextCancellation(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate slow response
		time.Sleep(100 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := &http.Client{Timeout: 5 * time.Second}
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel immediately

	_, err := GetPublicKey(ctx, client, server.URL)
	if err == nil {
		t.Fatal("expected error for cancelled context")
	}
}

func TestGetPublicKey_NetworkError(t *testing.T) {
	client := &http.Client{Timeout: 1 * time.Second}
	ctx := context.Background()

	// Use an invalid URL that will fail to connect
	_, err := GetPublicKey(ctx, client, "http://localhost:99999")
	if err == nil {
		t.Fatal("expected network error")
	}
	if !contains(err.Error(), "request failed") {
		t.Errorf("error should mention request failure, got: %v", err)
	}
}

// contains is a helper function to check if a string contains a substring
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > 0 && containsHelper(s, substr))
}

func containsHelper(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

// Tests for retry behavior

func TestGetPublicKey_RetryOnRateLimit(t *testing.T) {
	attempts := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempts++
		if attempts < 2 {
			// First request: return rate limit
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte(`{"error": "rate limited"}`))
			return
		}
		// Second request: return success
		response := APIEnvelope[PublicKey]{
			Data: PublicKey{
				KeyVersion:           1,
				ModulusBase64:        "dGVzdC1tb2R1bHVz",
				PublicExponentBase64: "AQAB",
				Tenant:               "test-tenant",
			},
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := &http.Client{Timeout: 5 * time.Second}
	ctx := context.Background()

	pubKey, err := GetPublicKey(ctx, client, server.URL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if pubKey == nil {
		t.Fatal("expected public key, got nil")
	}
	if attempts != 2 {
		t.Errorf("expected 2 attempts, got %d", attempts)
	}
}

func TestGetPublicKey_RetryOnServerError(t *testing.T) {
	retryableErrors := []int{
		http.StatusBadGateway,         // 502
		http.StatusServiceUnavailable, // 503
		http.StatusGatewayTimeout,     // 504
	}

	for _, statusCode := range retryableErrors {
		t.Run(http.StatusText(statusCode), func(t *testing.T) {
			attempts := 0
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				attempts++
				if attempts < 2 {
					w.WriteHeader(statusCode)
					w.Write([]byte(`{"error": "server error"}`))
					return
				}
				response := APIEnvelope[PublicKey]{
					Data: PublicKey{
						KeyVersion:           1,
						ModulusBase64:        "dGVzdC1tb2R1bHVz",
						PublicExponentBase64: "AQAB",
						Tenant:               "test-tenant",
					},
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(response)
			}))
			defer server.Close()

			client := &http.Client{Timeout: 5 * time.Second}
			ctx := context.Background()

			pubKey, err := GetPublicKey(ctx, client, server.URL)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if pubKey == nil {
				t.Fatal("expected public key, got nil")
			}
			if attempts != 2 {
				t.Errorf("expected 2 attempts, got %d", attempts)
			}
		})
	}
}

func TestGetPublicKey_NoRetryOnClientError(t *testing.T) {
	nonRetryableErrors := []int{
		http.StatusBadRequest,   // 400
		http.StatusUnauthorized, // 401
		http.StatusForbidden,    // 403
		http.StatusNotFound,     // 404
	}

	for _, statusCode := range nonRetryableErrors {
		t.Run(http.StatusText(statusCode), func(t *testing.T) {
			attempts := 0
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				attempts++
				w.WriteHeader(statusCode)
				w.Write([]byte(`{"error": "client error"}`))
			}))
			defer server.Close()

			client := &http.Client{Timeout: 5 * time.Second}
			ctx := context.Background()

			_, err := GetPublicKey(ctx, client, server.URL)
			if err == nil {
				t.Fatal("expected error")
			}
			if attempts != 1 {
				t.Errorf("expected 1 attempt (no retry), got %d", attempts)
			}
		})
	}
}

func TestGetPublicKey_MaxRetriesExceeded(t *testing.T) {
	attempts := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempts++
		// Always return rate limit
		w.WriteHeader(http.StatusTooManyRequests)
		w.Write([]byte(`{"error": "rate limited"}`))
	}))
	defer server.Close()

	client := &http.Client{Timeout: 5 * time.Second}
	ctx := context.Background()

	_, err := GetPublicKey(ctx, client, server.URL)
	if err == nil {
		t.Fatal("expected error after max retries")
	}
	// Should have tried defaultMaxRetries + 1 times (initial + retries)
	expectedAttempts := defaultMaxRetries + 1
	if attempts != expectedAttempts {
		t.Errorf("expected %d attempts, got %d", expectedAttempts, attempts)
	}
}

func TestIsRetryableStatus(t *testing.T) {
	tests := []struct {
		statusCode int
		want       bool
	}{
		{http.StatusOK, false},
		{http.StatusBadRequest, false},
		{http.StatusUnauthorized, false},
		{http.StatusForbidden, false},
		{http.StatusNotFound, false},
		{http.StatusTooManyRequests, true}, // 429
		{http.StatusInternalServerError, false},
		{http.StatusBadGateway, true},         // 502
		{http.StatusServiceUnavailable, true}, // 503
		{http.StatusGatewayTimeout, true},     // 504
	}

	for _, tt := range tests {
		t.Run(http.StatusText(tt.statusCode), func(t *testing.T) {
			got := isRetryableStatus(tt.statusCode)
			if got != tt.want {
				t.Errorf("isRetryableStatus(%d) = %v, want %v", tt.statusCode, got, tt.want)
			}
		})
	}
}

func TestCalculateBackoff(t *testing.T) {
	tests := []struct {
		attempt int
		wantMin time.Duration
		wantMax time.Duration
	}{
		{0, 1 * time.Second, 1 * time.Second},    // 1s * 2^0 = 1s
		{1, 2 * time.Second, 2 * time.Second},    // 1s * 2^1 = 2s
		{2, 4 * time.Second, 4 * time.Second},    // 1s * 2^2 = 4s
		{3, 8 * time.Second, 8 * time.Second},    // 1s * 2^3 = 8s
		{4, 16 * time.Second, 16 * time.Second},  // 1s * 2^4 = 16s
		{5, 30 * time.Second, 30 * time.Second},  // capped at max
		{10, 30 * time.Second, 30 * time.Second}, // capped at max
	}

	for _, tt := range tests {
		t.Run("attempt_"+string(rune('0'+tt.attempt)), func(t *testing.T) {
			got := calculateBackoff(tt.attempt)
			if got < tt.wantMin || got > tt.wantMax {
				t.Errorf("calculateBackoff(%d) = %v, want between %v and %v", tt.attempt, got, tt.wantMin, tt.wantMax)
			}
		})
	}
}
