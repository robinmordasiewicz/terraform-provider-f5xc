// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package blindfold

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetSecretPolicyDocument_MockServer(t *testing.T) {
	tests := []struct {
		name           string
		namespace      string
		policyName     string
		responseCode   int
		responseBody   interface{}
		wantErr        bool
		wantErrContain string
		wantPolicyID   string
		wantTenant     string
	}{
		{
			name:         "successful response with envelope",
			namespace:    "shared",
			policyName:   "ves-io-allow-volterra",
			responseCode: http.StatusOK,
			responseBody: APIEnvelope[SecretPolicyDocument]{
				Data: SecretPolicyDocument{
					Name:      "ves-io-allow-volterra",
					Namespace: "shared",
					Tenant:    "test-tenant",
					PolicyID:  "policy-id-12345",
					PolicyInfo: SecretPolicyInfo{
						Algo: "RSA-OAEP",
						Rules: []SecretPolicyRule{
							{Action: "ALLOW"},
						},
					},
				},
			},
			wantErr:      false,
			wantPolicyID: "policy-id-12345",
			wantTenant:   "test-tenant",
		},
		{
			name:         "successful response with data wrapper map",
			namespace:    "production",
			policyName:   "my-policy",
			responseCode: http.StatusOK,
			responseBody: map[string]interface{}{
				"data": map[string]interface{}{
					"name":      "my-policy",
					"namespace": "production",
					"tenant":    "prod-tenant",
					"policy_id": "policy-67890",
					"policy_info": map[string]interface{}{
						"algo": "RSA-OAEP",
					},
				},
			},
			wantErr:      false,
			wantPolicyID: "policy-67890",
			wantTenant:   "prod-tenant",
		},
		{
			name:           "policy not found",
			namespace:      "shared",
			policyName:     "nonexistent-policy",
			responseCode:   http.StatusNotFound,
			responseBody:   map[string]string{"error": "not found"},
			wantErr:        true,
			wantErrContain: "not found in namespace",
		},
		{
			name:           "server error",
			namespace:      "shared",
			policyName:     "test-policy",
			responseCode:   http.StatusInternalServerError,
			responseBody:   map[string]string{"error": "internal error"},
			wantErr:        true,
			wantErrContain: "unexpected status 500",
		},
		{
			name:           "unauthorized",
			namespace:      "shared",
			policyName:     "test-policy",
			responseCode:   http.StatusUnauthorized,
			responseBody:   map[string]string{"error": "unauthorized"},
			wantErr:        true,
			wantErrContain: "unexpected status 401",
		},
		{
			name:         "missing policy_id",
			namespace:    "shared",
			policyName:   "test-policy",
			responseCode: http.StatusOK,
			responseBody: APIEnvelope[SecretPolicyDocument]{
				Data: SecretPolicyDocument{
					Name:      "test-policy",
					Namespace: "shared",
					Tenant:    "test-tenant",
					// PolicyID intentionally missing
				},
			},
			wantErr:        true,
			wantErrContain: "missing policy_id",
		},
		{
			name:           "invalid JSON response",
			namespace:      "shared",
			policyName:     "test-policy",
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

				// Verify endpoint format
				expectedPath := fmt.Sprintf(PolicyDocumentEndpointFmt, tt.namespace, tt.policyName)
				if r.URL.Path != expectedPath {
					t.Errorf("expected path %s, got %s", expectedPath, r.URL.Path)
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

			policy, err := GetSecretPolicyDocument(ctx, client, server.URL, tt.namespace, tt.policyName)

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

			if policy.PolicyID != tt.wantPolicyID {
				t.Errorf("PolicyID = %v, want %v", policy.PolicyID, tt.wantPolicyID)
			}
			if policy.Tenant != tt.wantTenant {
				t.Errorf("Tenant = %v, want %v", policy.Tenant, tt.wantTenant)
			}
		})
	}
}

func TestGetSecretPolicyDocument_URLEncoding(t *testing.T) {
	// Test that special characters in namespace and policy name are properly URL encoded
	tests := []struct {
		name            string
		namespace       string
		policyName      string
		wantEncodedPath string
	}{
		{
			name:            "normal names",
			namespace:       "shared",
			policyName:      "ves-io-allow-volterra",
			wantEncodedPath: "/api/secret_management/namespaces/shared/secret_policys/ves-io-allow-volterra/get_policy_document",
		},
		{
			name:            "names with spaces (url encoded)",
			namespace:       "my namespace",
			policyName:      "my policy",
			wantEncodedPath: "/api/secret_management/namespaces/my%20namespace/secret_policys/my%20policy/get_policy_document",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var receivedPath string
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				receivedPath = r.URL.EscapedPath()

				response := APIEnvelope[SecretPolicyDocument]{
					Data: SecretPolicyDocument{
						PolicyID: "test-policy-id",
					},
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(response)
			}))
			defer server.Close()

			client := &http.Client{Timeout: 5 * time.Second}
			ctx := context.Background()

			_, err := GetSecretPolicyDocument(ctx, client, server.URL, tt.namespace, tt.policyName)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if receivedPath != tt.wantEncodedPath {
				t.Errorf("path = %q, want %q", receivedPath, tt.wantEncodedPath)
			}
		})
	}
}

func TestGetSecretPolicyDocument_NilClient(t *testing.T) {
	ctx := context.Background()
	_, err := GetSecretPolicyDocument(ctx, nil, "https://example.com", "shared", "policy")
	if err == nil {
		t.Fatal("expected error for nil client")
	}
	if !contains(err.Error(), "HTTP client is required") {
		t.Errorf("error should mention HTTP client, got: %v", err)
	}
}

func TestGetSecretPolicyDocument_EmptyBaseURL(t *testing.T) {
	ctx := context.Background()
	client := &http.Client{Timeout: 5 * time.Second}
	_, err := GetSecretPolicyDocument(ctx, client, "", "shared", "policy")
	if err == nil {
		t.Fatal("expected error for empty base URL")
	}
	if !contains(err.Error(), "base URL is required") {
		t.Errorf("error should mention base URL, got: %v", err)
	}
}

func TestGetSecretPolicyDocument_EmptyNamespace(t *testing.T) {
	ctx := context.Background()
	client := &http.Client{Timeout: 5 * time.Second}
	_, err := GetSecretPolicyDocument(ctx, client, "https://example.com", "", "policy")
	if err == nil {
		t.Fatal("expected error for empty namespace")
	}
	if !contains(err.Error(), "namespace is required") {
		t.Errorf("error should mention namespace, got: %v", err)
	}
}

func TestGetSecretPolicyDocument_EmptyPolicyName(t *testing.T) {
	ctx := context.Background()
	client := &http.Client{Timeout: 5 * time.Second}
	_, err := GetSecretPolicyDocument(ctx, client, "https://example.com", "shared", "")
	if err == nil {
		t.Fatal("expected error for empty policy name")
	}
	if !contains(err.Error(), "policy name is required") {
		t.Errorf("error should mention policy name, got: %v", err)
	}
}

func TestGetSecretPolicyDocument_ContextCancellation(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate slow response
		time.Sleep(100 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := &http.Client{Timeout: 5 * time.Second}
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel immediately

	_, err := GetSecretPolicyDocument(ctx, client, server.URL, "shared", "policy")
	if err == nil {
		t.Fatal("expected error for cancelled context")
	}
}

func TestGetSecretPolicyDocument_NetworkError(t *testing.T) {
	client := &http.Client{Timeout: 1 * time.Second}
	ctx := context.Background()

	// Use an invalid URL that will fail to connect
	_, err := GetSecretPolicyDocument(ctx, client, "http://localhost:99999", "shared", "policy")
	if err == nil {
		t.Fatal("expected network error")
	}
	if !contains(err.Error(), "request failed") {
		t.Errorf("error should mention request failure, got: %v", err)
	}
}

func TestGetSecretPolicyDocument_NamespaceFallback(t *testing.T) {
	// Test that name and namespace are set from parameters when not in response
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := APIEnvelope[SecretPolicyDocument]{
			Data: SecretPolicyDocument{
				// Name and Namespace intentionally not set
				PolicyID: "test-policy-id",
				Tenant:   "test-tenant",
			},
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := &http.Client{Timeout: 5 * time.Second}
	ctx := context.Background()

	policy, err := GetSecretPolicyDocument(ctx, client, server.URL, "my-namespace", "my-policy")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if policy.Name != "my-policy" {
		t.Errorf("Name = %q, want %q", policy.Name, "my-policy")
	}
	if policy.Namespace != "my-namespace" {
		t.Errorf("Namespace = %q, want %q", policy.Namespace, "my-namespace")
	}
}
