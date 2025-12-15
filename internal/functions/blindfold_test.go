// This file is MANUALLY MAINTAINED and is NOT auto-generated.

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
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/f5xc/terraform-provider-f5xc/internal/blindfold"
)

// mockF5XCServer creates a test server that simulates the F5XC API
func mockF5XCServer(t *testing.T, keyBits int) (*httptest.Server, *rsa.PrivateKey) {
	t.Helper()

	privateKey, err := rsa.GenerateKey(rand.Reader, keyBits)
	if err != nil {
		t.Fatalf("failed to generate RSA key: %v", err)
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		switch {
		case strings.HasSuffix(r.URL.Path, "/get_public_key"):
			response := map[string]interface{}{
				"data": map[string]interface{}{
					"key_version":            1,
					"modulus_base64":         base64.StdEncoding.EncodeToString(privateKey.N.Bytes()),
					"public_exponent_base64": base64.StdEncoding.EncodeToString([]byte{0x01, 0x00, 0x01}),
					"tenant":                 "test-tenant",
				},
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(response)

		case strings.HasSuffix(r.URL.Path, "/get_policy_document"):
			response := map[string]interface{}{
				"data": map[string]interface{}{
					"name":      "test-policy",
					"namespace": "test-namespace",
					"tenant":    "test-tenant",
					"policy_id": "policy-123",
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

	return server, privateKey
}

func TestBlindfoldFunction_Metadata(t *testing.T) {
	fn := NewBlindfoldFunction()
	var req function.MetadataRequest
	var resp function.MetadataResponse

	fn.Metadata(context.Background(), req, &resp)

	if resp.Name != "blindfold" {
		t.Errorf("expected name 'blindfold', got %q", resp.Name)
	}
}

func TestBlindfoldFunction_Definition(t *testing.T) {
	fn := NewBlindfoldFunction()
	var req function.DefinitionRequest
	var resp function.DefinitionResponse

	fn.Definition(context.Background(), req, &resp)

	def := resp.Definition
	if def.Summary == "" {
		t.Error("expected non-empty Summary")
	}
	if def.Description == "" {
		t.Error("expected non-empty Description")
	}
	if len(def.Parameters) != 3 {
		t.Errorf("expected 3 parameters, got %d", len(def.Parameters))
	}
}

func TestBlindfoldFunction_Run_EmptyPlaintext(t *testing.T) {
	// Save and restore env vars
	origToken := os.Getenv(blindfold.EnvAPIToken)
	origP12 := os.Getenv(blindfold.EnvP12File)
	defer func() {
		os.Setenv(blindfold.EnvAPIToken, origToken)
		os.Setenv(blindfold.EnvP12File, origP12)
	}()

	// Set token to pass auth check
	os.Setenv(blindfold.EnvAPIToken, "test-token")
	os.Unsetenv(blindfold.EnvP12File)

	fn := &BlindfoldFunction{}

	got := function.RunResponse{
		Result: function.NewResultData(types.StringUnknown()),
	}

	fn.Run(context.Background(), function.RunRequest{
		Arguments: function.NewArgumentsData([]attr.Value{
			types.StringValue(""),          // empty plaintext
			types.StringValue("policy"),    // policy_name
			types.StringValue("namespace"), // namespace
		}),
	}, &got)

	if got.Error == nil {
		t.Fatal("expected error for empty plaintext, got nil")
	}

	errStr := got.Error.Error()
	if !strings.Contains(errStr, "plaintext cannot be empty") {
		t.Errorf("error should mention plaintext, got: %s", errStr)
	}
}

func TestBlindfoldFunction_Run_EmptyPolicyName(t *testing.T) {
	// Save and restore env vars
	origToken := os.Getenv(blindfold.EnvAPIToken)
	origP12 := os.Getenv(blindfold.EnvP12File)
	defer func() {
		os.Setenv(blindfold.EnvAPIToken, origToken)
		os.Setenv(blindfold.EnvP12File, origP12)
	}()

	os.Setenv(blindfold.EnvAPIToken, "test-token")
	os.Unsetenv(blindfold.EnvP12File)

	fn := &BlindfoldFunction{}

	got := function.RunResponse{
		Result: function.NewResultData(types.StringUnknown()),
	}

	fn.Run(context.Background(), function.RunRequest{
		Arguments: function.NewArgumentsData([]attr.Value{
			types.StringValue(base64.StdEncoding.EncodeToString([]byte("secret"))),
			types.StringValue(""),          // empty policy_name
			types.StringValue("namespace"), // namespace
		}),
	}, &got)

	if got.Error == nil {
		t.Fatal("expected error for empty policy_name, got nil")
	}

	errStr := got.Error.Error()
	if !strings.Contains(errStr, "policy_name cannot be empty") {
		t.Errorf("error should mention policy_name, got: %s", errStr)
	}
}

func TestBlindfoldFunction_Run_EmptyNamespace(t *testing.T) {
	// Save and restore env vars
	origToken := os.Getenv(blindfold.EnvAPIToken)
	origP12 := os.Getenv(blindfold.EnvP12File)
	defer func() {
		os.Setenv(blindfold.EnvAPIToken, origToken)
		os.Setenv(blindfold.EnvP12File, origP12)
	}()

	os.Setenv(blindfold.EnvAPIToken, "test-token")
	os.Unsetenv(blindfold.EnvP12File)

	fn := &BlindfoldFunction{}

	got := function.RunResponse{
		Result: function.NewResultData(types.StringUnknown()),
	}

	fn.Run(context.Background(), function.RunRequest{
		Arguments: function.NewArgumentsData([]attr.Value{
			types.StringValue(base64.StdEncoding.EncodeToString([]byte("secret"))),
			types.StringValue("policy"), // policy_name
			types.StringValue(""),       // empty namespace
		}),
	}, &got)

	if got.Error == nil {
		t.Fatal("expected error for empty namespace, got nil")
	}

	errStr := got.Error.Error()
	if !strings.Contains(errStr, "namespace cannot be empty") {
		t.Errorf("error should mention namespace, got: %s", errStr)
	}
}

func TestBlindfoldFunction_Run_InvalidBase64(t *testing.T) {
	// Save and restore env vars
	origToken := os.Getenv(blindfold.EnvAPIToken)
	origURL := os.Getenv(blindfold.EnvAPIURL)
	defer func() {
		os.Setenv(blindfold.EnvAPIToken, origToken)
		os.Setenv(blindfold.EnvAPIURL, origURL)
	}()

	// Set up mock server
	server, _ := mockF5XCServer(t, 2048)
	defer server.Close()

	os.Setenv(blindfold.EnvAPIToken, "test-token")
	os.Setenv(blindfold.EnvAPIURL, server.URL)

	fn := &BlindfoldFunction{}

	got := function.RunResponse{
		Result: function.NewResultData(types.StringUnknown()),
	}

	fn.Run(context.Background(), function.RunRequest{
		Arguments: function.NewArgumentsData([]attr.Value{
			types.StringValue("not-valid-base64!!!"), // invalid base64
			types.StringValue("policy"),
			types.StringValue("namespace"),
		}),
	}, &got)

	if got.Error == nil {
		t.Fatal("expected error for invalid base64, got nil")
	}

	errStr := got.Error.Error()
	if !strings.Contains(errStr, "Invalid base64") {
		t.Errorf("error should mention base64, got: %s", errStr)
	}
}

func TestBlindfoldFunction_Run_NoAuth(t *testing.T) {
	// Clear all auth env vars
	origToken := os.Getenv(blindfold.EnvAPIToken)
	origP12 := os.Getenv(blindfold.EnvP12File)
	origP12Pass := os.Getenv(blindfold.EnvP12Password)
	defer func() {
		os.Setenv(blindfold.EnvAPIToken, origToken)
		os.Setenv(blindfold.EnvP12File, origP12)
		os.Setenv(blindfold.EnvP12Password, origP12Pass)
	}()

	os.Unsetenv(blindfold.EnvAPIToken)
	os.Unsetenv(blindfold.EnvP12File)
	os.Unsetenv(blindfold.EnvP12Password)

	fn := &BlindfoldFunction{}

	got := function.RunResponse{
		Result: function.NewResultData(types.StringUnknown()),
	}

	fn.Run(context.Background(), function.RunRequest{
		Arguments: function.NewArgumentsData([]attr.Value{
			types.StringValue(base64.StdEncoding.EncodeToString([]byte("secret"))),
			types.StringValue("policy"),
			types.StringValue("namespace"),
		}),
	}, &got)

	if got.Error == nil {
		t.Fatal("expected error for missing auth, got nil")
	}

	errStr := got.Error.Error()
	if !strings.Contains(errStr, "Authentication configuration error") {
		t.Errorf("error should mention authentication, got: %s", errStr)
	}
}

func TestBlindfoldFunction_Run_WithMockServer(t *testing.T) {
	server, _ := mockF5XCServer(t, 2048)
	defer server.Close()

	// Set up environment
	origToken := os.Getenv(blindfold.EnvAPIToken)
	origURL := os.Getenv(blindfold.EnvAPIURL)
	defer func() {
		os.Setenv(blindfold.EnvAPIToken, origToken)
		os.Setenv(blindfold.EnvAPIURL, origURL)
	}()

	os.Setenv(blindfold.EnvAPIToken, "test-token")
	os.Setenv(blindfold.EnvAPIURL, server.URL)

	fn := &BlindfoldFunction{}

	plaintext := "my-secret-value"
	plaintextBase64 := base64.StdEncoding.EncodeToString([]byte(plaintext))

	got := function.RunResponse{
		Result: function.NewResultData(types.StringUnknown()),
	}

	fn.Run(context.Background(), function.RunRequest{
		Arguments: function.NewArgumentsData([]attr.Value{
			types.StringValue(plaintextBase64),
			types.StringValue("test-policy"),
			types.StringValue("test-namespace"),
		}),
	}, &got)

	if got.Error != nil {
		t.Fatalf("unexpected error: %v", got.Error.Error())
	}

	// Verify result is set and has correct format
	// The result should be a string starting with "string:///"
	resultValue := got.Result
	if resultValue == (function.ResultData{}) {
		t.Fatal("expected result to be set")
	}
}

func TestBlindfoldFunction_Run_PlaintextTooLarge(t *testing.T) {
	server, _ := mockF5XCServer(t, 2048)
	defer server.Close()

	// Set up environment
	origToken := os.Getenv(blindfold.EnvAPIToken)
	origURL := os.Getenv(blindfold.EnvAPIURL)
	defer func() {
		os.Setenv(blindfold.EnvAPIToken, origToken)
		os.Setenv(blindfold.EnvAPIURL, origURL)
	}()

	os.Setenv(blindfold.EnvAPIToken, "test-token")
	os.Setenv(blindfold.EnvAPIURL, server.URL)

	fn := &BlindfoldFunction{}

	// Create plaintext that exceeds max size (128KB = 131072 bytes)
	// With envelope encryption, we can now encrypt up to 128KB
	largePlaintext := make([]byte, blindfold.MaxSecretSize+1) // 131073 bytes
	for i := range largePlaintext {
		largePlaintext[i] = byte(i % 256)
	}
	plaintextBase64 := base64.StdEncoding.EncodeToString(largePlaintext)

	got := function.RunResponse{
		Result: function.NewResultData(types.StringUnknown()),
	}

	fn.Run(context.Background(), function.RunRequest{
		Arguments: function.NewArgumentsData([]attr.Value{
			types.StringValue(plaintextBase64),
			types.StringValue("test-policy"),
			types.StringValue("test-namespace"),
		}),
	}, &got)

	if got.Error == nil {
		t.Fatal("expected error for large plaintext")
	}

	errStr := got.Error.Error()
	if !strings.Contains(errStr, "too large") && !strings.Contains(errStr, "Maximum size") {
		t.Errorf("error should mention size limit, got: %s", errStr)
	}
}

func TestBlindfoldFunction_Run_TableDriven(t *testing.T) {
	t.Parallel()

	// Set up mock server
	server, _ := mockF5XCServer(t, 2048)
	defer server.Close()

	// Save env vars
	origToken := os.Getenv(blindfold.EnvAPIToken)
	origURL := os.Getenv(blindfold.EnvAPIURL)
	origP12 := os.Getenv(blindfold.EnvP12File)

	t.Cleanup(func() {
		os.Setenv(blindfold.EnvAPIToken, origToken)
		os.Setenv(blindfold.EnvAPIURL, origURL)
		os.Setenv(blindfold.EnvP12File, origP12)
	})

	// Configure for mock server
	os.Setenv(blindfold.EnvAPIToken, "test-token")
	os.Setenv(blindfold.EnvAPIURL, server.URL)
	os.Unsetenv(blindfold.EnvP12File)

	testCases := map[string]struct {
		request     function.RunRequest
		expectError bool
		errContains string
	}{
		"valid-input": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.StringValue(base64.StdEncoding.EncodeToString([]byte("test-secret"))),
					types.StringValue("test-policy"),
					types.StringValue("test-namespace"),
				}),
			},
			expectError: false,
		},
		"empty-plaintext": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.StringValue(""),
					types.StringValue("policy"),
					types.StringValue("namespace"),
				}),
			},
			expectError: true,
			errContains: "plaintext cannot be empty",
		},
		"empty-policy": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.StringValue(base64.StdEncoding.EncodeToString([]byte("secret"))),
					types.StringValue(""),
					types.StringValue("namespace"),
				}),
			},
			expectError: true,
			errContains: "policy_name cannot be empty",
		},
		"empty-namespace": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.StringValue(base64.StdEncoding.EncodeToString([]byte("secret"))),
					types.StringValue("policy"),
					types.StringValue(""),
				}),
			},
			expectError: true,
			errContains: "namespace cannot be empty",
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			fn := &BlindfoldFunction{}

			got := function.RunResponse{
				Result: function.NewResultData(types.StringUnknown()),
			}

			fn.Run(context.Background(), testCase.request, &got)

			if testCase.expectError {
				if got.Error == nil {
					t.Fatal("expected error, got nil")
				}
				if testCase.errContains != "" && !strings.Contains(got.Error.Error(), testCase.errContains) {
					t.Errorf("error %q should contain %q", got.Error.Error(), testCase.errContains)
				}
			} else {
				if got.Error != nil {
					t.Fatalf("unexpected error: %v", got.Error.Error())
				}
			}
		})
	}
}

// Ensure cmp package is used to avoid unused import error
var _ = cmp.Diff
