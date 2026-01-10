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
