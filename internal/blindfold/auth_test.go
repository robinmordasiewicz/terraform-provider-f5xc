// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package blindfold

import (
	"net/http"
	"os"
	"path/filepath"
	"testing"
)

func TestAuthMethod_String(t *testing.T) {
	tests := []struct {
		name   string
		method AuthMethod
		want   string
	}{
		{"None", AuthMethodNone, "None"},
		{"Token", AuthMethodToken, "API Token"},
		{"P12", AuthMethodP12, "P12 Certificate"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.method.String(); got != tt.want {
				t.Errorf("AuthMethod.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAuthConfigFromEnv(t *testing.T) {
	// Save original env vars
	origToken := os.Getenv(EnvAPIToken)
	origP12 := os.Getenv(EnvP12File)
	origP12Pass := os.Getenv(EnvP12Password)
	origURL := os.Getenv(EnvAPIURL)

	// Cleanup after test
	defer func() {
		os.Setenv(EnvAPIToken, origToken)
		os.Setenv(EnvP12File, origP12)
		os.Setenv(EnvP12Password, origP12Pass)
		os.Setenv(EnvAPIURL, origURL)
	}()

	tests := []struct {
		name        string
		envVars     map[string]string
		wantErr     bool
		wantToken   string
		wantP12     string
		wantBaseURL string
	}{
		{
			name: "token auth configured",
			envVars: map[string]string{
				EnvAPIToken: "test-token",
				EnvAPIURL:   "https://test.example.com/api",
			},
			wantErr:     false,
			wantToken:   "test-token",
			wantBaseURL: "https://test.example.com/api",
		},
		{
			name: "P12 auth configured",
			envVars: map[string]string{
				EnvP12File:     "/path/to/cert.p12",
				EnvP12Password: "password",
				EnvAPIURL:      "https://test.example.com/api",
			},
			wantErr:     false,
			wantP12:     "/path/to/cert.p12",
			wantBaseURL: "https://test.example.com/api",
		},
		{
			name: "default URL when not set",
			envVars: map[string]string{
				EnvAPIToken: "test-token",
			},
			wantErr:     false,
			wantToken:   "test-token",
			wantBaseURL: DefaultAPIURL,
		},
		{
			name:    "no auth configured",
			envVars: map[string]string{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Clear all env vars
			os.Unsetenv(EnvAPIToken)
			os.Unsetenv(EnvP12File)
			os.Unsetenv(EnvP12Password)
			os.Unsetenv(EnvAPIURL)

			// Set test env vars
			for k, v := range tt.envVars {
				os.Setenv(k, v)
			}

			config, err := GetAuthConfigFromEnv()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAuthConfigFromEnv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			if config.APIToken != tt.wantToken {
				t.Errorf("APIToken = %v, want %v", config.APIToken, tt.wantToken)
			}
			if config.P12File != tt.wantP12 {
				t.Errorf("P12File = %v, want %v", config.P12File, tt.wantP12)
			}
			if config.BaseURL != tt.wantBaseURL {
				t.Errorf("BaseURL = %v, want %v", config.BaseURL, tt.wantBaseURL)
			}
		})
	}
}

func TestCreateAuthenticatedClient(t *testing.T) {
	tests := []struct {
		name       string
		config     *AuthConfig
		wantErr    bool
		wantMethod AuthMethod
	}{
		{
			name:    "nil config",
			config:  nil,
			wantErr: true,
		},
		{
			name: "token auth",
			config: &AuthConfig{
				APIToken: "test-token",
				BaseURL:  "https://test.example.com/api",
			},
			wantErr:    false,
			wantMethod: AuthMethodToken,
		},
		{
			name: "no auth method",
			config: &AuthConfig{
				BaseURL: "https://test.example.com/api",
			},
			wantErr: true,
		},
		{
			name: "token takes priority over P12",
			config: &AuthConfig{
				APIToken: "test-token",
				P12File:  "/path/to/cert.p12",
				BaseURL:  "https://test.example.com/api",
			},
			wantErr:    false,
			wantMethod: AuthMethodToken,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := CreateAuthenticatedClient(tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateAuthenticatedClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			if result.Method != tt.wantMethod {
				t.Errorf("Method = %v, want %v", result.Method, tt.wantMethod)
			}
			if result.Client == nil {
				t.Error("Client is nil")
			}
		})
	}
}

func TestCreateAuthenticatedClient_P12(t *testing.T) {
	// Skip if no test P12 file available
	// This test would require a real P12 file to be created for testing
	t.Run("invalid P12 file", func(t *testing.T) {
		config := &AuthConfig{
			P12File:     "/nonexistent/path.p12",
			P12Password: "password",
			BaseURL:     "https://test.example.com/api",
		}

		_, err := CreateAuthenticatedClient(config)
		if err == nil {
			t.Error("expected error for nonexistent P12 file")
		}
	})
}

func TestAuthResult_SetAuthorizationHeader(t *testing.T) {
	tests := []struct {
		name       string
		result     *AuthResult
		wantHeader string
	}{
		{
			name: "token auth sets header",
			result: &AuthResult{
				Method: AuthMethodToken,
				Token:  "test-token",
			},
			wantHeader: "APIToken test-token",
		},
		{
			name: "P12 auth no header",
			result: &AuthResult{
				Method: AuthMethodP12,
			},
			wantHeader: "",
		},
		{
			name: "no auth no header",
			result: &AuthResult{
				Method: AuthMethodNone,
			},
			wantHeader: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "https://example.com", nil)
			tt.result.SetAuthorizationHeader(req)

			got := req.Header.Get("Authorization")
			if got != tt.wantHeader {
				t.Errorf("Authorization header = %v, want %v", got, tt.wantHeader)
			}
		})
	}
}

func TestAuthResult_NewRequest(t *testing.T) {
	result := &AuthResult{
		Method: AuthMethodToken,
		Token:  "test-token",
	}

	req, err := result.NewRequest("GET", "https://example.com/api/test")
	if err != nil {
		t.Fatalf("NewRequest() error = %v", err)
	}

	if req.Method != "GET" {
		t.Errorf("Method = %v, want GET", req.Method)
	}

	authHeader := req.Header.Get("Authorization")
	if authHeader != "APIToken test-token" {
		t.Errorf("Authorization = %v, want APIToken test-token", authHeader)
	}
}

func TestValidateConfig(t *testing.T) {
	// Create a temporary file for P12 validation test
	tmpDir := t.TempDir()
	tmpP12 := filepath.Join(tmpDir, "test.p12")
	if err := os.WriteFile(tmpP12, []byte("dummy"), 0600); err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}

	tests := []struct {
		name    string
		config  *AuthConfig
		wantErr bool
	}{
		{
			name:    "nil config",
			config:  nil,
			wantErr: true,
		},
		{
			name: "empty base URL",
			config: &AuthConfig{
				APIToken: "token",
			},
			wantErr: true,
		},
		{
			name: "no auth method",
			config: &AuthConfig{
				BaseURL: "https://example.com",
			},
			wantErr: true,
		},
		{
			name: "valid token auth",
			config: &AuthConfig{
				APIToken: "token",
				BaseURL:  "https://example.com",
			},
			wantErr: false,
		},
		{
			name: "P12 file not found",
			config: &AuthConfig{
				P12File: "/nonexistent/path.p12",
				BaseURL: "https://example.com",
			},
			wantErr: true,
		},
		{
			name: "valid P12 auth with existing file",
			config: &AuthConfig{
				P12File:     tmpP12,
				P12Password: "password",
				BaseURL:     "https://example.com",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateConfig(tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
