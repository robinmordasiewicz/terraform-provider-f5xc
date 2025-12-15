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
	"path/filepath"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/f5xc/terraform-provider-f5xc/internal/blindfold"
)

// mockF5XCServerForFile creates a test server that simulates the F5XC API for file tests
func mockF5XCServerForFile(t *testing.T, keyBits int) (*httptest.Server, *rsa.PrivateKey) {
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

func TestBlindfoldFileFunction_Metadata(t *testing.T) {
	fn := NewBlindfoldFileFunction()
	var req function.MetadataRequest
	var resp function.MetadataResponse

	fn.Metadata(context.Background(), req, &resp)

	if resp.Name != "blindfold_file" {
		t.Errorf("expected name 'blindfold_file', got %q", resp.Name)
	}
}

func TestBlindfoldFileFunction_Definition(t *testing.T) {
	fn := NewBlindfoldFileFunction()
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

func TestBlindfoldFileFunction_Run_EmptyPath(t *testing.T) {
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

	fn := &BlindfoldFileFunction{}

	got := function.RunResponse{
		Result: function.NewResultData(types.StringUnknown()),
	}

	fn.Run(context.Background(), function.RunRequest{
		Arguments: function.NewArgumentsData([]attr.Value{
			types.StringValue(""),          // empty path
			types.StringValue("policy"),    // policy_name
			types.StringValue("namespace"), // namespace
		}),
	}, &got)

	if got.Error == nil {
		t.Fatal("expected error for empty path, got nil")
	}

	errStr := got.Error.Error()
	if !strings.Contains(errStr, "path cannot be empty") {
		t.Errorf("error should mention path, got: %s", errStr)
	}
}

func TestBlindfoldFileFunction_Run_EmptyPolicyName(t *testing.T) {
	// Save and restore env vars
	origToken := os.Getenv(blindfold.EnvAPIToken)
	origP12 := os.Getenv(blindfold.EnvP12File)
	defer func() {
		os.Setenv(blindfold.EnvAPIToken, origToken)
		os.Setenv(blindfold.EnvP12File, origP12)
	}()

	os.Setenv(blindfold.EnvAPIToken, "test-token")
	os.Unsetenv(blindfold.EnvP12File)

	fn := &BlindfoldFileFunction{}

	got := function.RunResponse{
		Result: function.NewResultData(types.StringUnknown()),
	}

	fn.Run(context.Background(), function.RunRequest{
		Arguments: function.NewArgumentsData([]attr.Value{
			types.StringValue("/some/path"), // path
			types.StringValue(""),           // empty policy_name
			types.StringValue("namespace"),  // namespace
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

func TestBlindfoldFileFunction_Run_EmptyNamespace(t *testing.T) {
	// Save and restore env vars
	origToken := os.Getenv(blindfold.EnvAPIToken)
	origP12 := os.Getenv(blindfold.EnvP12File)
	defer func() {
		os.Setenv(blindfold.EnvAPIToken, origToken)
		os.Setenv(blindfold.EnvP12File, origP12)
	}()

	os.Setenv(blindfold.EnvAPIToken, "test-token")
	os.Unsetenv(blindfold.EnvP12File)

	fn := &BlindfoldFileFunction{}

	got := function.RunResponse{
		Result: function.NewResultData(types.StringUnknown()),
	}

	fn.Run(context.Background(), function.RunRequest{
		Arguments: function.NewArgumentsData([]attr.Value{
			types.StringValue("/some/path"), // path
			types.StringValue("policy"),     // policy_name
			types.StringValue(""),           // empty namespace
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

func TestBlindfoldFileFunction_Run_NoAuth(t *testing.T) {
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

	fn := &BlindfoldFileFunction{}

	got := function.RunResponse{
		Result: function.NewResultData(types.StringUnknown()),
	}

	fn.Run(context.Background(), function.RunRequest{
		Arguments: function.NewArgumentsData([]attr.Value{
			types.StringValue("/some/path"),
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

func TestBlindfoldFileFunction_Run_FileNotFound(t *testing.T) {
	server, _ := mockF5XCServerForFile(t, 2048)
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

	fn := &BlindfoldFileFunction{}

	got := function.RunResponse{
		Result: function.NewResultData(types.StringUnknown()),
	}

	fn.Run(context.Background(), function.RunRequest{
		Arguments: function.NewArgumentsData([]attr.Value{
			types.StringValue("/nonexistent/path/to/file.txt"),
			types.StringValue("test-policy"),
			types.StringValue("test-namespace"),
		}),
	}, &got)

	if got.Error == nil {
		t.Fatal("expected error for file not found, got nil")
	}

	errStr := got.Error.Error()
	if !strings.Contains(errStr, "File not found") {
		t.Errorf("error should mention file not found, got: %s", errStr)
	}
}

func TestBlindfoldFileFunction_Run_DirectoryNotFile(t *testing.T) {
	server, _ := mockF5XCServerForFile(t, 2048)
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

	// Create a temporary directory
	tmpDir := t.TempDir()

	fn := &BlindfoldFileFunction{}

	got := function.RunResponse{
		Result: function.NewResultData(types.StringUnknown()),
	}

	fn.Run(context.Background(), function.RunRequest{
		Arguments: function.NewArgumentsData([]attr.Value{
			types.StringValue(tmpDir), // directory, not a file
			types.StringValue("test-policy"),
			types.StringValue("test-namespace"),
		}),
	}, &got)

	if got.Error == nil {
		t.Fatal("expected error for directory path, got nil")
	}

	errStr := got.Error.Error()
	if !strings.Contains(errStr, "not a regular file") {
		t.Errorf("error should mention not a regular file, got: %s", errStr)
	}
}

func TestBlindfoldFileFunction_Run_EmptyFile(t *testing.T) {
	server, _ := mockF5XCServerForFile(t, 2048)
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

	// Create a temporary empty file
	tmpDir := t.TempDir()
	emptyFile := filepath.Join(tmpDir, "empty.txt")
	if err := os.WriteFile(emptyFile, []byte{}, 0644); err != nil {
		t.Fatalf("failed to create empty file: %v", err)
	}

	fn := &BlindfoldFileFunction{}

	got := function.RunResponse{
		Result: function.NewResultData(types.StringUnknown()),
	}

	fn.Run(context.Background(), function.RunRequest{
		Arguments: function.NewArgumentsData([]attr.Value{
			types.StringValue(emptyFile),
			types.StringValue("test-policy"),
			types.StringValue("test-namespace"),
		}),
	}, &got)

	if got.Error == nil {
		t.Fatal("expected error for empty file, got nil")
	}

	errStr := got.Error.Error()
	if !strings.Contains(errStr, "File is empty") {
		t.Errorf("error should mention empty file, got: %s", errStr)
	}
}

func TestBlindfoldFileFunction_Run_WithMockServer(t *testing.T) {
	server, _ := mockF5XCServerForFile(t, 2048)
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

	// Create a temporary file with content
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "secret.txt")
	if err := os.WriteFile(testFile, []byte("my-secret-value"), 0644); err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}

	fn := &BlindfoldFileFunction{}

	got := function.RunResponse{
		Result: function.NewResultData(types.StringUnknown()),
	}

	fn.Run(context.Background(), function.RunRequest{
		Arguments: function.NewArgumentsData([]attr.Value{
			types.StringValue(testFile),
			types.StringValue("test-policy"),
			types.StringValue("test-namespace"),
		}),
	}, &got)

	if got.Error != nil {
		t.Fatalf("unexpected error: %v", got.Error.Error())
	}

	// Verify result is set
	resultValue := got.Result
	if resultValue == (function.ResultData{}) {
		t.Fatal("expected result to be set")
	}
}

func TestBlindfoldFileFunction_Run_FileTooLarge(t *testing.T) {
	server, _ := mockF5XCServerForFile(t, 2048)
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

	// Create a temporary file that exceeds max size (128KB API limit)
	// With envelope encryption, we can now encrypt up to 128KB
	tmpDir := t.TempDir()
	largeFile := filepath.Join(tmpDir, "large.txt")
	largeContent := make([]byte, blindfold.MaxSecretSize+1) // 131073 bytes
	for i := range largeContent {
		largeContent[i] = byte(i % 256)
	}
	if err := os.WriteFile(largeFile, largeContent, 0644); err != nil {
		t.Fatalf("failed to create large file: %v", err)
	}

	fn := &BlindfoldFileFunction{}

	got := function.RunResponse{
		Result: function.NewResultData(types.StringUnknown()),
	}

	fn.Run(context.Background(), function.RunRequest{
		Arguments: function.NewArgumentsData([]attr.Value{
			types.StringValue(largeFile),
			types.StringValue("test-policy"),
			types.StringValue("test-namespace"),
		}),
	}, &got)

	if got.Error == nil {
		t.Fatal("expected error for large file")
	}

	errStr := got.Error.Error()
	if !strings.Contains(errStr, "too large") {
		t.Errorf("error should mention size limit, got: %s", errStr)
	}
}

func TestBlindfoldFileFunction_Run_TableDriven(t *testing.T) {
	t.Parallel()

	// Set up mock server
	server, _ := mockF5XCServerForFile(t, 2048)
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

	// Create test files
	tmpDir := t.TempDir()
	validFile := filepath.Join(tmpDir, "valid.txt")
	if err := os.WriteFile(validFile, []byte("test-secret"), 0644); err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}

	testCases := map[string]struct {
		request     function.RunRequest
		expectError bool
		errContains string
	}{
		"valid-input": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.StringValue(validFile),
					types.StringValue("test-policy"),
					types.StringValue("test-namespace"),
				}),
			},
			expectError: false,
		},
		"empty-path": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.StringValue(""),
					types.StringValue("policy"),
					types.StringValue("namespace"),
				}),
			},
			expectError: true,
			errContains: "path cannot be empty",
		},
		"empty-policy": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.StringValue(validFile),
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
					types.StringValue(validFile),
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
			fn := &BlindfoldFileFunction{}

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

func TestBlindfoldFileFunction_Run_RelativePath(t *testing.T) {
	server, _ := mockF5XCServerForFile(t, 2048)
	defer server.Close()

	// Set up environment
	origToken := os.Getenv(blindfold.EnvAPIToken)
	origURL := os.Getenv(blindfold.EnvAPIURL)
	origWd, _ := os.Getwd()
	defer func() {
		os.Setenv(blindfold.EnvAPIToken, origToken)
		os.Setenv(blindfold.EnvAPIURL, origURL)
		os.Chdir(origWd)
	}()

	os.Setenv(blindfold.EnvAPIToken, "test-token")
	os.Setenv(blindfold.EnvAPIURL, server.URL)

	// Create a temporary directory and file, then change to that directory
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "secret.txt")
	if err := os.WriteFile(testFile, []byte("my-secret-value"), 0644); err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}

	// Change to the temp directory
	if err := os.Chdir(tmpDir); err != nil {
		t.Fatalf("failed to change directory: %v", err)
	}

	fn := &BlindfoldFileFunction{}

	got := function.RunResponse{
		Result: function.NewResultData(types.StringUnknown()),
	}

	// Use relative path
	fn.Run(context.Background(), function.RunRequest{
		Arguments: function.NewArgumentsData([]attr.Value{
			types.StringValue("secret.txt"), // relative path
			types.StringValue("test-policy"),
			types.StringValue("test-namespace"),
		}),
	}, &got)

	if got.Error != nil {
		t.Fatalf("unexpected error with relative path: %v", got.Error.Error())
	}

	// Verify result is set
	resultValue := got.Result
	if resultValue == (function.ResultData{}) {
		t.Fatal("expected result to be set")
	}
}

func TestBlindfoldFileFunction_Run_PathTraversal(t *testing.T) {
	server, _ := mockF5XCServerForFile(t, 2048)
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

	// Create a temporary file in a nested structure
	tmpDir := t.TempDir()
	subDir := filepath.Join(tmpDir, "subdir")
	if err := os.MkdirAll(subDir, 0755); err != nil {
		t.Fatalf("failed to create subdir: %v", err)
	}
	testFile := filepath.Join(tmpDir, "secret.txt")
	if err := os.WriteFile(testFile, []byte("my-secret-value"), 0644); err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}

	fn := &BlindfoldFileFunction{}

	got := function.RunResponse{
		Result: function.NewResultData(types.StringUnknown()),
	}

	// Use path with .. to traverse up
	pathWithTraversal := filepath.Join(subDir, "..", "secret.txt")
	fn.Run(context.Background(), function.RunRequest{
		Arguments: function.NewArgumentsData([]attr.Value{
			types.StringValue(pathWithTraversal),
			types.StringValue("test-policy"),
			types.StringValue("test-namespace"),
		}),
	}, &got)

	// filepath.Clean in the implementation should handle this correctly
	if got.Error != nil {
		t.Fatalf("unexpected error with path traversal: %v", got.Error.Error())
	}
}

func TestBlindfoldFileFunction_Run_BinaryFile(t *testing.T) {
	server, _ := mockF5XCServerForFile(t, 2048)
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

	// Create a temporary file with binary content
	tmpDir := t.TempDir()
	binaryFile := filepath.Join(tmpDir, "binary.bin")
	// Binary data with null bytes and non-printable characters
	binaryContent := []byte{0x00, 0x01, 0x02, 0x03, 0xFF, 0xFE, 0xFD, 'H', 'e', 'l', 'l', 'o'}
	if err := os.WriteFile(binaryFile, binaryContent, 0644); err != nil {
		t.Fatalf("failed to create binary file: %v", err)
	}

	fn := &BlindfoldFileFunction{}

	got := function.RunResponse{
		Result: function.NewResultData(types.StringUnknown()),
	}

	fn.Run(context.Background(), function.RunRequest{
		Arguments: function.NewArgumentsData([]attr.Value{
			types.StringValue(binaryFile),
			types.StringValue("test-policy"),
			types.StringValue("test-namespace"),
		}),
	}, &got)

	// Binary files should be encrypted successfully
	if got.Error != nil {
		t.Fatalf("unexpected error with binary file: %v", got.Error.Error())
	}

	// Verify result is set
	resultValue := got.Result
	if resultValue == (function.ResultData{}) {
		t.Fatal("expected result to be set")
	}
}
