// This file is MANUALLY MAINTAINED and is NOT auto-generated.
//
// Mock tests for blindfold functions that can run in CI/CD without real F5XC credentials.
// These tests use httptest servers to simulate the F5XC API.
//
// Run with: go test -v -run "TestMock" ./internal/functions/...

package functions

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/f5xc/terraform-provider-f5xc/internal/blindfold"
)

// mockServer creates a comprehensive mock F5XC server for CI/CD testing
type mockServer struct {
	server     *httptest.Server
	privateKey *rsa.PrivateKey
	keyBits    int
	// Configurable response behaviors
	publicKeyError bool
	policyError    bool
	policyNotFound bool
	serverError    bool
	invalidJSON    bool
	customPolicyID string
	customTenant   string
}

func newMockServer(t *testing.T, keyBits int) *mockServer {
	t.Helper()

	privateKey, err := rsa.GenerateKey(rand.Reader, keyBits)
	if err != nil {
		t.Fatalf("failed to generate RSA key: %v", err)
	}

	m := &mockServer{
		privateKey:     privateKey,
		keyBits:        keyBits,
		customPolicyID: "mock-policy-123",
		customTenant:   "mock-tenant",
	}

	m.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if m.serverError {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "internal server error"})
			return
		}

		if m.invalidJSON {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("not valid json"))
			return
		}

		switch {
		case strings.HasSuffix(r.URL.Path, "/get_public_key"):
			if m.publicKeyError {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]string{"error": "unauthorized"})
				return
			}
			response := map[string]interface{}{
				"data": map[string]interface{}{
					"key_version":            1,
					"modulus_base64":         base64.StdEncoding.EncodeToString(m.privateKey.N.Bytes()),
					"public_exponent_base64": base64.StdEncoding.EncodeToString([]byte{0x01, 0x00, 0x01}),
					"tenant":                 m.customTenant,
				},
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(response)

		case strings.HasSuffix(r.URL.Path, "/get_policy_document"):
			if m.policyNotFound {
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(map[string]string{"error": "policy not found"})
				return
			}
			if m.policyError {
				w.WriteHeader(http.StatusForbidden)
				json.NewEncoder(w).Encode(map[string]string{"error": "forbidden"})
				return
			}
			response := map[string]interface{}{
				"data": map[string]interface{}{
					"name":      "mock-policy",
					"namespace": "mock-namespace",
					"tenant":    m.customTenant,
					"policy_id": m.customPolicyID,
					"policy_info": map[string]interface{}{
						"algo": "RSA-OAEP",
					},
				},
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(response)

		default:
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "not found"})
		}
	}))

	return m
}

func (m *mockServer) Close() {
	m.server.Close()
}

func (m *mockServer) URL() string {
	return m.server.URL
}

// setupMockEnv configures environment for mock testing
func setupMockEnv(t *testing.T, serverURL string) func() {
	t.Helper()

	origToken := os.Getenv(blindfold.EnvAPIToken)
	origURL := os.Getenv(blindfold.EnvAPIURL)
	origP12 := os.Getenv(blindfold.EnvP12File)
	origP12Pass := os.Getenv(blindfold.EnvP12Password)

	os.Setenv(blindfold.EnvAPIToken, "mock-token")
	os.Setenv(blindfold.EnvAPIURL, serverURL)
	os.Unsetenv(blindfold.EnvP12File)
	os.Unsetenv(blindfold.EnvP12Password)

	return func() {
		os.Setenv(blindfold.EnvAPIToken, origToken)
		os.Setenv(blindfold.EnvAPIURL, origURL)
		os.Setenv(blindfold.EnvP12File, origP12)
		os.Setenv(blindfold.EnvP12Password, origP12Pass)
	}
}

func TestMockBlindfoldFunction_Success(t *testing.T) {
	mock := newMockServer(t, 2048)
	defer mock.Close()

	cleanup := setupMockEnv(t, mock.URL())
	defer cleanup()

	fn := &BlindfoldFunction{}

	plaintext := "mock-test-secret"
	plaintextBase64 := base64.StdEncoding.EncodeToString([]byte(plaintext))

	got := function.RunResponse{
		Result: function.NewResultData(types.StringUnknown()),
	}

	fn.Run(context.Background(), function.RunRequest{
		Arguments: function.NewArgumentsData([]attr.Value{
			types.StringValue(plaintextBase64),
			types.StringValue("mock-policy"),
			types.StringValue("mock-namespace"),
		}),
	}, &got)

	if got.Error != nil {
		t.Fatalf("unexpected error: %v", got.Error.Error())
	}
}

func TestMockBlindfoldFunction_PublicKeyError(t *testing.T) {
	mock := newMockServer(t, 2048)
	mock.publicKeyError = true
	defer mock.Close()

	cleanup := setupMockEnv(t, mock.URL())
	defer cleanup()

	fn := &BlindfoldFunction{}

	plaintext := "test-secret"
	plaintextBase64 := base64.StdEncoding.EncodeToString([]byte(plaintext))

	got := function.RunResponse{
		Result: function.NewResultData(types.StringUnknown()),
	}

	fn.Run(context.Background(), function.RunRequest{
		Arguments: function.NewArgumentsData([]attr.Value{
			types.StringValue(plaintextBase64),
			types.StringValue("policy"),
			types.StringValue("namespace"),
		}),
	}, &got)

	if got.Error == nil {
		t.Fatal("expected error for public key failure, got nil")
	}

	errStr := got.Error.Error()
	if !strings.Contains(errStr, "public key") {
		t.Errorf("error should mention public key, got: %s", errStr)
	}
}

func TestMockBlindfoldFunction_PolicyNotFound(t *testing.T) {
	mock := newMockServer(t, 2048)
	mock.policyNotFound = true
	defer mock.Close()

	cleanup := setupMockEnv(t, mock.URL())
	defer cleanup()

	fn := &BlindfoldFunction{}

	plaintext := "test-secret"
	plaintextBase64 := base64.StdEncoding.EncodeToString([]byte(plaintext))

	got := function.RunResponse{
		Result: function.NewResultData(types.StringUnknown()),
	}

	fn.Run(context.Background(), function.RunRequest{
		Arguments: function.NewArgumentsData([]attr.Value{
			types.StringValue(plaintextBase64),
			types.StringValue("nonexistent-policy"),
			types.StringValue("namespace"),
		}),
	}, &got)

	if got.Error == nil {
		t.Fatal("expected error for policy not found, got nil")
	}

	errStr := got.Error.Error()
	if !strings.Contains(errStr, "policy") {
		t.Errorf("error should mention policy, got: %s", errStr)
	}
}

func TestMockBlindfoldFunction_PolicyForbidden(t *testing.T) {
	mock := newMockServer(t, 2048)
	mock.policyError = true
	defer mock.Close()

	cleanup := setupMockEnv(t, mock.URL())
	defer cleanup()

	fn := &BlindfoldFunction{}

	plaintext := "test-secret"
	plaintextBase64 := base64.StdEncoding.EncodeToString([]byte(plaintext))

	got := function.RunResponse{
		Result: function.NewResultData(types.StringUnknown()),
	}

	fn.Run(context.Background(), function.RunRequest{
		Arguments: function.NewArgumentsData([]attr.Value{
			types.StringValue(plaintextBase64),
			types.StringValue("forbidden-policy"),
			types.StringValue("namespace"),
		}),
	}, &got)

	if got.Error == nil {
		t.Fatal("expected error for forbidden policy, got nil")
	}
}

func TestMockBlindfoldFunction_ServerError(t *testing.T) {
	mock := newMockServer(t, 2048)
	mock.serverError = true
	defer mock.Close()

	cleanup := setupMockEnv(t, mock.URL())
	defer cleanup()

	fn := &BlindfoldFunction{}

	plaintext := "test-secret"
	plaintextBase64 := base64.StdEncoding.EncodeToString([]byte(plaintext))

	got := function.RunResponse{
		Result: function.NewResultData(types.StringUnknown()),
	}

	fn.Run(context.Background(), function.RunRequest{
		Arguments: function.NewArgumentsData([]attr.Value{
			types.StringValue(plaintextBase64),
			types.StringValue("policy"),
			types.StringValue("namespace"),
		}),
	}, &got)

	if got.Error == nil {
		t.Fatal("expected error for server error, got nil")
	}
}

func TestMockBlindfoldFunction_DifferentKeySizes(t *testing.T) {
	keySizes := []int{2048, 3072, 4096}

	for _, keySize := range keySizes {
		t.Run(strings.ReplaceAll(string(rune(keySize)), "", ""), func(t *testing.T) {
			mock := newMockServer(t, keySize)
			defer mock.Close()

			cleanup := setupMockEnv(t, mock.URL())
			defer cleanup()

			fn := &BlindfoldFunction{}

			// Use a small plaintext that will work with any key size
			plaintext := "small-secret"
			plaintextBase64 := base64.StdEncoding.EncodeToString([]byte(plaintext))

			got := function.RunResponse{
				Result: function.NewResultData(types.StringUnknown()),
			}

			fn.Run(context.Background(), function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.StringValue(plaintextBase64),
					types.StringValue("policy"),
					types.StringValue("namespace"),
				}),
			}, &got)

			if got.Error != nil {
				t.Errorf("key size %d failed: %v", keySize, got.Error.Error())
			}
		})
	}
}

func TestMockBlindfoldFileFunction_Success(t *testing.T) {
	mock := newMockServer(t, 2048)
	defer mock.Close()

	cleanup := setupMockEnv(t, mock.URL())
	defer cleanup()

	// Create test file
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "secret.txt")
	if err := os.WriteFile(testFile, []byte("file-secret-content"), 0644); err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}

	fn := &BlindfoldFileFunction{}

	got := function.RunResponse{
		Result: function.NewResultData(types.StringUnknown()),
	}

	fn.Run(context.Background(), function.RunRequest{
		Arguments: function.NewArgumentsData([]attr.Value{
			types.StringValue(testFile),
			types.StringValue("mock-policy"),
			types.StringValue("mock-namespace"),
		}),
	}, &got)

	if got.Error != nil {
		t.Fatalf("unexpected error: %v", got.Error.Error())
	}
}

func TestMockBlindfoldFileFunction_VariousFileTypes(t *testing.T) {
	mock := newMockServer(t, 2048)
	defer mock.Close()

	cleanup := setupMockEnv(t, mock.URL())
	defer cleanup()

	tmpDir := t.TempDir()

	testCases := []struct {
		name    string
		content []byte
	}{
		{"text", []byte("plain text content")},
		{"json", []byte(`{"key": "value", "nested": {"a": 1}}`)},
		{"pem", []byte("-----BEGIN CERTIFICATE-----\nMIIB...\n-----END CERTIFICATE-----")},
		{"binary", []byte{0x00, 0x01, 0x02, 0xFF, 0xFE, 0xFD}},
		{"unicode", []byte("Unicode: 你好世界 🌍 émojis")},
		{"newlines", []byte("line1\nline2\r\nline3\rline4")},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testFile := filepath.Join(tmpDir, tc.name+".txt")
			if err := os.WriteFile(testFile, tc.content, 0644); err != nil {
				t.Fatalf("failed to create test file: %v", err)
			}

			fn := &BlindfoldFileFunction{}

			got := function.RunResponse{
				Result: function.NewResultData(types.StringUnknown()),
			}

			fn.Run(context.Background(), function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.StringValue(testFile),
					types.StringValue("policy"),
					types.StringValue("namespace"),
				}),
			}, &got)

			if got.Error != nil {
				t.Errorf("file type %s failed: %v", tc.name, got.Error.Error())
			}
		})
	}
}

func TestMockBlindfoldFunction_ConcurrentRequests(t *testing.T) {
	mock := newMockServer(t, 2048)
	defer mock.Close()

	cleanup := setupMockEnv(t, mock.URL())
	defer cleanup()

	// Test concurrent encryption requests
	const numRequests = 10
	results := make(chan error, numRequests)

	for i := 0; i < numRequests; i++ {
		go func(index int) {
			fn := &BlindfoldFunction{}

			plaintext := strings.Repeat("secret-", index+1)
			plaintextBase64 := base64.StdEncoding.EncodeToString([]byte(plaintext))

			got := function.RunResponse{
				Result: function.NewResultData(types.StringUnknown()),
			}

			fn.Run(context.Background(), function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.StringValue(plaintextBase64),
					types.StringValue("policy"),
					types.StringValue("namespace"),
				}),
			}, &got)

			if got.Error != nil {
				results <- got.Error
			} else {
				results <- nil
			}
		}(i)
	}

	// Collect results
	var errors []error
	for i := 0; i < numRequests; i++ {
		if err := <-results; err != nil {
			errors = append(errors, err)
		}
	}

	if len(errors) > 0 {
		t.Errorf("concurrent requests had %d errors: %v", len(errors), errors)
	}
}

func TestMockBlindfoldFunction_CustomTenantAndPolicy(t *testing.T) {
	mock := newMockServer(t, 2048)
	mock.customTenant = "custom-tenant-xyz"
	mock.customPolicyID = "custom-policy-abc"
	defer mock.Close()

	cleanup := setupMockEnv(t, mock.URL())
	defer cleanup()

	fn := &BlindfoldFunction{}

	plaintext := "test-secret"
	plaintextBase64 := base64.StdEncoding.EncodeToString([]byte(plaintext))

	got := function.RunResponse{
		Result: function.NewResultData(types.StringUnknown()),
	}

	fn.Run(context.Background(), function.RunRequest{
		Arguments: function.NewArgumentsData([]attr.Value{
			types.StringValue(plaintextBase64),
			types.StringValue("my-policy"),
			types.StringValue("my-namespace"),
		}),
	}, &got)

	if got.Error != nil {
		t.Fatalf("unexpected error: %v", got.Error.Error())
	}
}

func TestMockBlindfoldFunction_EmptyInputValidation(t *testing.T) {
	mock := newMockServer(t, 2048)
	defer mock.Close()

	cleanup := setupMockEnv(t, mock.URL())
	defer cleanup()

	testCases := []struct {
		name        string
		plaintext   string
		policyName  string
		namespace   string
		errContains string
	}{
		{"empty plaintext", "", "policy", "namespace", "plaintext cannot be empty"},
		{"empty policy", "c2VjcmV0", "", "namespace", "policy_name cannot be empty"},
		{"empty namespace", "c2VjcmV0", "policy", "", "namespace cannot be empty"},
		{"whitespace plaintext", "   ", "policy", "namespace", ""}, // This is valid base64 of spaces
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fn := &BlindfoldFunction{}

			got := function.RunResponse{
				Result: function.NewResultData(types.StringUnknown()),
			}

			fn.Run(context.Background(), function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.StringValue(tc.plaintext),
					types.StringValue(tc.policyName),
					types.StringValue(tc.namespace),
				}),
			}, &got)

			if tc.errContains != "" {
				if got.Error == nil {
					t.Fatalf("expected error containing %q, got nil", tc.errContains)
				}
				if !strings.Contains(got.Error.Error(), tc.errContains) {
					t.Errorf("error should contain %q, got: %s", tc.errContains, got.Error.Error())
				}
			}
		})
	}
}
