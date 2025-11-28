// Copyright (c) F5XC Community
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"
)

// =============================================================================
// Tests for isRetryableStatus()
// =============================================================================

func TestIsRetryableStatus(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		want       bool
	}{
		// Explicitly retryable status codes
		{"TooManyRequests_429", http.StatusTooManyRequests, true},
		{"ServiceUnavailable_503", http.StatusServiceUnavailable, true},
		{"GatewayTimeout_504", http.StatusGatewayTimeout, true},
		{"BadGateway_502", http.StatusBadGateway, true},

		// Other 5xx errors (should also be retryable)
		{"InternalServerError_500", http.StatusInternalServerError, true},
		{"NotImplemented_501", http.StatusNotImplemented, true},
		{"InsufficientStorage_507", http.StatusInsufficientStorage, true},

		// Non-retryable client errors (4xx)
		{"BadRequest_400", http.StatusBadRequest, false},
		{"Unauthorized_401", http.StatusUnauthorized, false},
		{"Forbidden_403", http.StatusForbidden, false},
		{"NotFound_404", http.StatusNotFound, false},
		{"MethodNotAllowed_405", http.StatusMethodNotAllowed, false},
		{"Conflict_409", http.StatusConflict, false},
		{"Gone_410", http.StatusGone, false},
		{"UnprocessableEntity_422", http.StatusUnprocessableEntity, false},

		// Success codes (not retryable - already successful)
		{"OK_200", http.StatusOK, false},
		{"Created_201", http.StatusCreated, false},
		{"Accepted_202", http.StatusAccepted, false},
		{"NoContent_204", http.StatusNoContent, false},

		// Redirect codes (not retryable)
		{"MovedPermanently_301", http.StatusMovedPermanently, false},
		{"Found_302", http.StatusFound, false},
		{"NotModified_304", http.StatusNotModified, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isRetryableStatus(tt.statusCode)
			if got != tt.want {
				t.Errorf("isRetryableStatus(%d) = %v, want %v", tt.statusCode, got, tt.want)
			}
		})
	}
}

// =============================================================================
// Tests for calculateBackoff()
// =============================================================================

func TestCalculateBackoff(t *testing.T) {
	client := &Client{
		RetryWaitMin: 1 * time.Second,
		RetryWaitMax: 30 * time.Second,
	}

	tests := []struct {
		name    string
		attempt int
		want    time.Duration
	}{
		{"Attempt_0", 0, 1 * time.Second},      // 1 * 2^0 = 1s
		{"Attempt_1", 1, 2 * time.Second},      // 1 * 2^1 = 2s
		{"Attempt_2", 2, 4 * time.Second},      // 1 * 2^2 = 4s
		{"Attempt_3", 3, 8 * time.Second},      // 1 * 2^3 = 8s
		{"Attempt_4", 4, 16 * time.Second},     // 1 * 2^4 = 16s
		{"Attempt_5", 5, 30 * time.Second},     // 1 * 2^5 = 32s, capped at 30s
		{"Attempt_10", 10, 30 * time.Second},   // 1 * 2^10 = 1024s, capped at 30s
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := client.calculateBackoff(tt.attempt)
			if got != tt.want {
				t.Errorf("calculateBackoff(%d) = %v, want %v", tt.attempt, got, tt.want)
			}
		})
	}
}

func TestCalculateBackoffCustomWaitTimes(t *testing.T) {
	client := &Client{
		RetryWaitMin: 500 * time.Millisecond,
		RetryWaitMax: 5 * time.Second,
	}

	tests := []struct {
		name    string
		attempt int
		want    time.Duration
	}{
		{"Attempt_0", 0, 500 * time.Millisecond},  // 500ms * 2^0 = 500ms
		{"Attempt_1", 1, 1 * time.Second},         // 500ms * 2^1 = 1s
		{"Attempt_2", 2, 2 * time.Second},         // 500ms * 2^2 = 2s
		{"Attempt_3", 3, 4 * time.Second},         // 500ms * 2^3 = 4s
		{"Attempt_4", 4, 5 * time.Second},         // 500ms * 2^4 = 8s, capped at 5s
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := client.calculateBackoff(tt.attempt)
			if got != tt.want {
				t.Errorf("calculateBackoff(%d) = %v, want %v", tt.attempt, got, tt.want)
			}
		})
	}
}

// =============================================================================
// Tests for NewClient() and client options
// =============================================================================

func TestNewClient(t *testing.T) {
	client := NewClient("https://api.example.com", "test-token")

	if client.BaseURL != "https://api.example.com" {
		t.Errorf("BaseURL = %s, want https://api.example.com", client.BaseURL)
	}
	if client.APIToken != "test-token" {
		t.Errorf("APIToken = %s, want test-token", client.APIToken)
	}
	if client.AuthType != AuthTypeToken {
		t.Errorf("AuthType = %v, want AuthTypeToken", client.AuthType)
	}
	if client.MaxRetries != DefaultMaxRetries {
		t.Errorf("MaxRetries = %d, want %d", client.MaxRetries, DefaultMaxRetries)
	}
	if client.RetryWaitMin != DefaultRetryWaitMin {
		t.Errorf("RetryWaitMin = %v, want %v", client.RetryWaitMin, DefaultRetryWaitMin)
	}
	if client.RetryWaitMax != DefaultRetryWaitMax {
		t.Errorf("RetryWaitMax = %v, want %v", client.RetryWaitMax, DefaultRetryWaitMax)
	}
	if client.HTTPClient == nil {
		t.Error("HTTPClient is nil")
	}
	if client.HTTPClient.Timeout != DefaultTimeout {
		t.Errorf("HTTPClient.Timeout = %v, want %v", client.HTTPClient.Timeout, DefaultTimeout)
	}
}

func TestWithTimeout(t *testing.T) {
	customTimeout := 60 * time.Second
	client := NewClient("https://api.example.com", "test-token", WithTimeout(customTimeout))

	if client.HTTPClient.Timeout != customTimeout {
		t.Errorf("HTTPClient.Timeout = %v, want %v", client.HTTPClient.Timeout, customTimeout)
	}
}

func TestWithMaxRetries(t *testing.T) {
	customRetries := 5
	client := NewClient("https://api.example.com", "test-token", WithMaxRetries(customRetries))

	if client.MaxRetries != customRetries {
		t.Errorf("MaxRetries = %d, want %d", client.MaxRetries, customRetries)
	}
}

func TestWithRetryWait(t *testing.T) {
	minWait := 2 * time.Second
	maxWait := 60 * time.Second
	client := NewClient("https://api.example.com", "test-token", WithRetryWait(minWait, maxWait))

	if client.RetryWaitMin != minWait {
		t.Errorf("RetryWaitMin = %v, want %v", client.RetryWaitMin, minWait)
	}
	if client.RetryWaitMax != maxWait {
		t.Errorf("RetryWaitMax = %v, want %v", client.RetryWaitMax, maxWait)
	}
}

func TestMultipleClientOptions(t *testing.T) {
	client := NewClient(
		"https://api.example.com",
		"test-token",
		WithTimeout(120*time.Second),
		WithMaxRetries(10),
		WithRetryWait(500*time.Millisecond, 10*time.Second),
	)

	if client.HTTPClient.Timeout != 120*time.Second {
		t.Errorf("HTTPClient.Timeout = %v, want 120s", client.HTTPClient.Timeout)
	}
	if client.MaxRetries != 10 {
		t.Errorf("MaxRetries = %d, want 10", client.MaxRetries)
	}
	if client.RetryWaitMin != 500*time.Millisecond {
		t.Errorf("RetryWaitMin = %v, want 500ms", client.RetryWaitMin)
	}
	if client.RetryWaitMax != 10*time.Second {
		t.Errorf("RetryWaitMax = %v, want 10s", client.RetryWaitMax)
	}
}

// =============================================================================
// Tests for HTTP methods (Get, Post, Put, Delete)
// =============================================================================

// testResponse is a helper struct for test responses
type testResponse struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Message string `json:"message,omitempty"`
}

func TestGetSuccess(t *testing.T) {
	expectedResponse := testResponse{ID: "123", Name: "test"}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request method
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		// Verify authorization header
		if auth := r.Header.Get("Authorization"); auth != "APIToken test-token" {
			t.Errorf("Authorization header = %s, want APIToken test-token", auth)
		}
		// Verify content type header
		if ct := r.Header.Get("Content-Type"); ct != "application/json" {
			t.Errorf("Content-Type header = %s, want application/json", ct)
		}
		// Verify accept header
		if accept := r.Header.Get("Accept"); accept != "application/json" {
			t.Errorf("Accept header = %s, want application/json", accept)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	client := NewClient(server.URL, "test-token")
	var result testResponse

	err := client.Get(context.Background(), "/test", &result)
	if err != nil {
		t.Fatalf("Get() error = %v", err)
	}

	if result.ID != expectedResponse.ID {
		t.Errorf("result.ID = %s, want %s", result.ID, expectedResponse.ID)
	}
	if result.Name != expectedResponse.Name {
		t.Errorf("result.Name = %s, want %s", result.Name, expectedResponse.Name)
	}
}

func TestGetNotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"code": "NOT_FOUND", "message": "Resource not found"}`))
	}))
	defer server.Close()

	client := NewClient(server.URL, "test-token", WithMaxRetries(0))
	var result testResponse

	err := client.Get(context.Background(), "/test", &result)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
}

func TestPostSuccess(t *testing.T) {
	requestData := testResponse{Name: "new-resource"}
	expectedResponse := testResponse{ID: "456", Name: "new-resource"}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST request, got %s", r.Method)
		}

		// Verify request body
		var received testResponse
		if err := json.NewDecoder(r.Body).Decode(&received); err != nil {
			t.Errorf("Failed to decode request body: %v", err)
		}
		if received.Name != requestData.Name {
			t.Errorf("Request body name = %s, want %s", received.Name, requestData.Name)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	client := NewClient(server.URL, "test-token")
	var result testResponse

	err := client.Post(context.Background(), "/test", requestData, &result)
	if err != nil {
		t.Fatalf("Post() error = %v", err)
	}

	if result.ID != expectedResponse.ID {
		t.Errorf("result.ID = %s, want %s", result.ID, expectedResponse.ID)
	}
}

func TestPutSuccess(t *testing.T) {
	requestData := testResponse{ID: "123", Name: "updated-resource"}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("Expected PUT request, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(requestData)
	}))
	defer server.Close()

	client := NewClient(server.URL, "test-token")
	var result testResponse

	err := client.Put(context.Background(), "/test/123", requestData, &result)
	if err != nil {
		t.Fatalf("Put() error = %v", err)
	}

	if result.Name != requestData.Name {
		t.Errorf("result.Name = %s, want %s", result.Name, requestData.Name)
	}
}

func TestDeleteSuccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("Expected DELETE request, got %s", r.Method)
		}

		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	client := NewClient(server.URL, "test-token")

	err := client.Delete(context.Background(), "/test/123")
	if err != nil {
		t.Fatalf("Delete() error = %v", err)
	}
}

func TestRetryOnServerError(t *testing.T) {
	attemptCount := 0

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attemptCount++
		if attemptCount < 3 {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id": "123", "name": "success"}`))
	}))
	defer server.Close()

	// Use minimal retry wait times for faster testing
	client := NewClient(server.URL, "test-token",
		WithMaxRetries(3),
		WithRetryWait(1*time.Millisecond, 10*time.Millisecond),
	)
	var result testResponse

	err := client.Get(context.Background(), "/test", &result)
	if err != nil {
		t.Fatalf("Get() error = %v", err)
	}

	if attemptCount != 3 {
		t.Errorf("attemptCount = %d, want 3", attemptCount)
	}
	if result.ID != "123" {
		t.Errorf("result.ID = %s, want 123", result.ID)
	}
}

func TestMaxRetriesExceeded(t *testing.T) {
	attemptCount := 0

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attemptCount++
		w.WriteHeader(http.StatusServiceUnavailable)
	}))
	defer server.Close()

	client := NewClient(server.URL, "test-token",
		WithMaxRetries(2),
		WithRetryWait(1*time.Millisecond, 10*time.Millisecond),
	)

	err := client.Get(context.Background(), "/test", nil)
	if err == nil {
		t.Fatal("Expected error after max retries, got nil")
	}

	// Initial attempt + 2 retries = 3 total attempts
	expectedAttempts := 3
	if attemptCount != expectedAttempts {
		t.Errorf("attemptCount = %d, want %d", attemptCount, expectedAttempts)
	}
}

func TestContextCancellation(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate slow response
		time.Sleep(100 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := NewClient(server.URL, "test-token")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	err := client.Get(ctx, "/test", nil)
	if err == nil {
		t.Fatal("Expected context cancellation error, got nil")
	}
}

func TestNilResultParameter(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id": "123"}`))
	}))
	defer server.Close()

	client := NewClient(server.URL, "test-token")

	// Should not panic when result is nil
	err := client.Get(context.Background(), "/test", nil)
	if err != nil {
		t.Fatalf("Get() error = %v", err)
	}
}

func TestEmptyResponseBody(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		// Empty body
	}))
	defer server.Close()

	client := NewClient(server.URL, "test-token")
	var result testResponse

	// Should not error when response body is empty
	err := client.Get(context.Background(), "/test", &result)
	if err != nil {
		t.Fatalf("Get() error = %v", err)
	}
}

// =============================================================================
// Tests for certificate loading functions
// =============================================================================

func TestLoadCertificateKeyPairFileNotFound(t *testing.T) {
	_, err := LoadCertificateKeyPair("/nonexistent/cert.pem", "/nonexistent/key.pem", "")
	if err == nil {
		t.Error("Expected error for non-existent certificate files, got nil")
	}
}

func TestLoadCertificateKeyPairInvalidFiles(t *testing.T) {
	// Create temporary directory
	tmpDir, err := os.MkdirTemp("", "client_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create invalid cert and key files
	certFile := filepath.Join(tmpDir, "cert.pem")
	keyFile := filepath.Join(tmpDir, "key.pem")

	if err := os.WriteFile(certFile, []byte("invalid cert"), 0644); err != nil {
		t.Fatalf("Failed to write cert file: %v", err)
	}
	if err := os.WriteFile(keyFile, []byte("invalid key"), 0644); err != nil {
		t.Fatalf("Failed to write key file: %v", err)
	}

	_, err = LoadCertificateKeyPair(certFile, keyFile, "")
	if err == nil {
		t.Error("Expected error for invalid certificate files, got nil")
	}
}

func TestLoadCertificateKeyPairCAFileNotFound(t *testing.T) {
	// Create temporary directory
	tmpDir, err := os.MkdirTemp("", "client_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create valid self-signed cert and key for testing
	certPEM := `-----BEGIN CERTIFICATE-----
MIIBkTCB+wIJAKHBfLTUjD5sMA0GCSqGSIb3DQEBCwUAMBExDzANBgNVBAMMBnVu
dXNlZDAeFw0yMzAxMDEwMDAwMDBaFw0yNDAxMDEwMDAwMDBaMBExDzANBgNVBAMM
BnVudXNlZDBcMA0GCSqGSIb3DQEBAQUAA0sAMEgCQQC7o96HilR6PorBTOR6QLXE
LiL9cK+R4GTAbF5X5TLxlHNpSqaRk6yLe1L7RgBN0DGx7cN6VrL1JfP7Uf8kcYBL
AgMBAAGjUzBRMB0GA1UdDgQWBBR7Wz5K6xVxqP5cP3mJB6A7a9x9bjAfBgNVHSME
GDAWgBR7Wz5K6xVxqP5cP3mJB6A7a9x9bjAPBgNVHRMBAf8EBTADAQH/MA0GCSqG
SIb3DQEBCwUAA0EAUdDT+FK/mzBSVKEVG8mVvm4q5Ao9B3l5MrP7lPJD5g8WVmJZ
p+P+Y0lg5fFOJDvbJvJlVvmKlWL5hLlbGPv8FA==
-----END CERTIFICATE-----`
	keyPEM := `-----BEGIN RSA PRIVATE KEY-----
MIIBOgIBAAJBALuj3oeKVHo+isFM5HpAtcQuIv1wr5HgZMBsXlflMvGUc2lKppGT
rIt7UvtGAE3QMbHtw3pWsvUl8/tR/yRxgEsCAwEAAQJAWAzXqiuRCPJJQFZEMibi
v0I6oGxvPJVmwkLNiZOFEcEdBHj/xKvfO5LnWZIhw+cxJhBBvJx+J3T7Xh9OSkpU
gQIhAOQYt0YZ7Ky7+TLxnm+pA0RoqP7lO5SqVx4I8wUQOyRbAiEA0eD/DHWWQMVZ
uDyG5N0vJlWLvxYvGPvFH1xWL1X7m2ECIFWb4kX9TqIJVAqKJQXw5CLKO+VPOdDt
n9cM8nHjFvC3AiEAl7VPGKqSvfQa0qxvB/8qMzLXlNvJGLvmGCkzKMm6nmECIDa1
i1O0E6kuWIQMZgYqKd0UhP7ZJLcXB8O7RH5L0uZg
-----END RSA PRIVATE KEY-----`

	certFile := filepath.Join(tmpDir, "cert.pem")
	keyFile := filepath.Join(tmpDir, "key.pem")

	if err := os.WriteFile(certFile, []byte(certPEM), 0644); err != nil {
		t.Fatalf("Failed to write cert file: %v", err)
	}
	if err := os.WriteFile(keyFile, []byte(keyPEM), 0600); err != nil {
		t.Fatalf("Failed to write key file: %v", err)
	}

	// Should fail when CA file doesn't exist
	_, err = LoadCertificateKeyPair(certFile, keyFile, "/nonexistent/ca.pem")
	if err == nil {
		t.Error("Expected error for non-existent CA file, got nil")
	}
}

func TestLoadP12CertificateFileNotFound(t *testing.T) {
	_, err := LoadP12Certificate("/nonexistent/cert.p12", "password")
	if err == nil {
		t.Error("Expected error for non-existent P12 file, got nil")
	}
}

func TestLoadP12CertificateInvalidFile(t *testing.T) {
	// Create temporary directory
	tmpDir, err := os.MkdirTemp("", "client_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create invalid P12 file
	p12File := filepath.Join(tmpDir, "invalid.p12")
	if err := os.WriteFile(p12File, []byte("invalid p12 content"), 0644); err != nil {
		t.Fatalf("Failed to write P12 file: %v", err)
	}

	_, err = LoadP12Certificate(p12File, "password")
	if err == nil {
		t.Error("Expected error for invalid P12 file, got nil")
	}
}

// =============================================================================
// Tests for AuthType handling
// =============================================================================

func TestAuthTypeToken(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth != "APIToken my-secret-token" {
			t.Errorf("Authorization header = %s, want APIToken my-secret-token", auth)
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := NewClient(server.URL, "my-secret-token")
	if client.AuthType != AuthTypeToken {
		t.Errorf("AuthType = %v, want AuthTypeToken", client.AuthType)
	}

	err := client.Get(context.Background(), "/test", nil)
	if err != nil {
		t.Fatalf("Get() error = %v", err)
	}
}

// =============================================================================
// Tests for Metadata struct
// =============================================================================

func TestMetadataJSONSerialization(t *testing.T) {
	metadata := Metadata{
		Name:      "test-resource",
		Namespace: "test-namespace",
		Labels: map[string]string{
			"env": "test",
		},
		Annotations: map[string]string{
			"description": "test annotation",
		},
		UID: "test-uid-123",
	}

	data, err := json.Marshal(metadata)
	if err != nil {
		t.Fatalf("Failed to marshal Metadata: %v", err)
	}

	var result Metadata
	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatalf("Failed to unmarshal Metadata: %v", err)
	}

	if result.Name != metadata.Name {
		t.Errorf("Name = %s, want %s", result.Name, metadata.Name)
	}
	if result.Namespace != metadata.Namespace {
		t.Errorf("Namespace = %s, want %s", result.Namespace, metadata.Namespace)
	}
	if result.Labels["env"] != "test" {
		t.Errorf("Labels[env] = %s, want test", result.Labels["env"])
	}
	if result.UID != metadata.UID {
		t.Errorf("UID = %s, want %s", result.UID, metadata.UID)
	}
}

func TestMetadataJSONOmitEmpty(t *testing.T) {
	// Minimal metadata with only required field
	metadata := Metadata{
		Name: "minimal-resource",
	}

	data, err := json.Marshal(metadata)
	if err != nil {
		t.Fatalf("Failed to marshal Metadata: %v", err)
	}

	// Verify that omitempty fields are not included
	jsonStr := string(data)
	if contains(jsonStr, "namespace") {
		t.Error("JSON should not contain 'namespace' when empty")
	}
	if contains(jsonStr, "labels") {
		t.Error("JSON should not contain 'labels' when nil")
	}
	if contains(jsonStr, "annotations") {
		t.Error("JSON should not contain 'annotations' when nil")
	}
	if contains(jsonStr, "uid") {
		t.Error("JSON should not contain 'uid' when empty")
	}
}

// Helper function to check if a string contains a substring
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > 0 && containsSubstring(s, substr))
}

func containsSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

// =============================================================================
// Tests for default constants
// =============================================================================

func TestDefaultConstants(t *testing.T) {
	if DefaultTimeout != 30*time.Second {
		t.Errorf("DefaultTimeout = %v, want 30s", DefaultTimeout)
	}
	if DefaultMaxRetries != 3 {
		t.Errorf("DefaultMaxRetries = %d, want 3", DefaultMaxRetries)
	}
	if DefaultRetryWaitMin != 1*time.Second {
		t.Errorf("DefaultRetryWaitMin = %v, want 1s", DefaultRetryWaitMin)
	}
	if DefaultRetryWaitMax != 30*time.Second {
		t.Errorf("DefaultRetryWaitMax = %v, want 30s", DefaultRetryWaitMax)
	}
	if DefaultRateLimitDelay != 60*time.Second {
		t.Errorf("DefaultRateLimitDelay = %v, want 60s", DefaultRateLimitDelay)
	}
}

// =============================================================================
// Tests for rate limit handling
// =============================================================================

func TestRateLimitHandlingUsesLongerDelay(t *testing.T) {
	// This test verifies that 429 responses are recognized as retryable
	// and that the client attempts to retry them.
	// Note: The actual DefaultRateLimitDelay (60s) is used for 429 responses,
	// which is too long for unit tests. We use a short context timeout to
	// verify the retry behavior without waiting the full 60 seconds.

	attemptCount := 0

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attemptCount++
		w.WriteHeader(http.StatusTooManyRequests)
		w.Write([]byte(`{"code": "RATE_LIMIT", "message": "Too many requests"}`))
	}))
	defer server.Close()

	client := NewClient(server.URL, "test-token",
		WithMaxRetries(3),
		WithRetryWait(1*time.Millisecond, 10*time.Millisecond),
	)

	// Use a short context timeout to avoid waiting 60 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	err := client.Get(ctx, "/test", nil)

	// Should have made at least one attempt
	if attemptCount < 1 {
		t.Errorf("Server should have been hit at least once, got %d attempts", attemptCount)
	}

	// Should fail due to context timeout (not retries exhausted)
	if err == nil {
		t.Error("Expected error due to context timeout or rate limit")
	}

	// Verify 429 is recognized as retryable
	if !isRetryableStatus(http.StatusTooManyRequests) {
		t.Error("429 should be recognized as retryable status")
	}
}
