// This file is MANUALLY MAINTAINED and is NOT auto-generated.

package blindfold

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"strings"
	"testing"
)

// generateTestKeyPair creates an RSA key pair for testing
func generateTestKeyPair(t *testing.T, bits int) (*rsa.PrivateKey, *PublicKey) {
	t.Helper()

	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		t.Fatalf("failed to generate RSA key: %v", err)
	}

	pubKey := &PublicKey{
		KeyVersion:           1,
		ModulusBase64:        base64.StdEncoding.EncodeToString(privateKey.N.Bytes()),
		PublicExponentBase64: base64.StdEncoding.EncodeToString([]byte{0x01, 0x00, 0x01}), // 65537
		Tenant:               "test-tenant",
	}

	return privateKey, pubKey
}

func TestSeal(t *testing.T) {
	_, pubKey := generateTestKeyPair(t, 2048)
	policy := &SecretPolicyDocument{
		Name:      "test-policy",
		Namespace: "test-namespace",
		PolicyID:  "policy-123",
		Tenant:    "test-tenant",
	}

	tests := []struct {
		name      string
		plaintext []byte
		pubKey    *PublicKey
		policy    *SecretPolicyDocument
		wantErr   bool
		errMsg    string
	}{
		{
			name:      "valid encryption",
			plaintext: []byte("secret-data"),
			pubKey:    pubKey,
			policy:    policy,
			wantErr:   false,
		},
		{
			name:      "nil public key",
			plaintext: []byte("secret-data"),
			pubKey:    nil,
			policy:    policy,
			wantErr:   true,
			errMsg:    "public key is required",
		},
		{
			name:      "nil policy",
			plaintext: []byte("secret-data"),
			pubKey:    pubKey,
			policy:    nil,
			wantErr:   true,
			errMsg:    "policy document is required",
		},
		{
			name:      "empty plaintext",
			plaintext: []byte{},
			pubKey:    pubKey,
			policy:    policy,
			wantErr:   true,
			errMsg:    "plaintext cannot be empty",
		},
		{
			name:      "binary data",
			plaintext: []byte{0x00, 0x01, 0x02, 0xff, 0xfe, 0xfd},
			pubKey:    pubKey,
			policy:    policy,
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sealed, err := Seal(tt.plaintext, tt.pubKey, tt.policy)

			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error containing %q, got nil", tt.errMsg)
				} else if !strings.Contains(err.Error(), tt.errMsg) {
					t.Errorf("expected error containing %q, got %q", tt.errMsg, err.Error())
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			// Verify format: string:///<base64>
			if !strings.HasPrefix(sealed, "string:///") {
				t.Errorf("sealed output should start with 'string:///', got %q", sealed[:min(20, len(sealed))])
			}

			// Verify we can decode the base64 payload
			payload := strings.TrimPrefix(sealed, "string:///")
			decoded, err := base64.StdEncoding.DecodeString(payload)
			if err != nil {
				t.Fatalf("failed to decode base64 payload: %v", err)
			}

			// Verify JSON structure
			var sealedSecret SealedSecret
			if err := json.Unmarshal(decoded, &sealedSecret); err != nil {
				t.Fatalf("failed to unmarshal sealed secret: %v", err)
			}

			if sealedSecret.KeyVersion != pubKey.KeyVersion {
				t.Errorf("key version mismatch: got %d, want %d", sealedSecret.KeyVersion, pubKey.KeyVersion)
			}
			if sealedSecret.PolicyID != policy.PolicyID {
				t.Errorf("policy ID mismatch: got %q, want %q", sealedSecret.PolicyID, policy.PolicyID)
			}
			if sealedSecret.Data == "" {
				t.Error("sealed data should not be empty")
			}
		})
	}
}

func TestSealBase64(t *testing.T) {
	_, pubKey := generateTestKeyPair(t, 2048)
	policy := &SecretPolicyDocument{
		Name:      "test-policy",
		Namespace: "test-namespace",
		PolicyID:  "policy-123",
	}

	plaintext := "my-secret-value"
	plaintextBase64 := base64.StdEncoding.EncodeToString([]byte(plaintext))

	sealed, err := SealBase64(plaintextBase64, pubKey, policy)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !strings.HasPrefix(sealed, "string:///") {
		t.Errorf("sealed output should start with 'string:///'")
	}
}

func TestSealBase64_InvalidBase64(t *testing.T) {
	_, pubKey := generateTestKeyPair(t, 2048)
	policy := &SecretPolicyDocument{
		PolicyID: "policy-123",
	}

	_, err := SealBase64("not-valid-base64!!!", pubKey, policy)
	if err == nil {
		t.Error("expected error for invalid base64")
	}
	if !strings.Contains(err.Error(), "failed to decode base64") {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestBuildRSAPublicKey(t *testing.T) {
	tests := []struct {
		name    string
		pubKey  *PublicKey
		wantErr bool
		errMsg  string
	}{
		{
			name: "valid key",
			pubKey: &PublicKey{
				ModulusBase64:        base64.StdEncoding.EncodeToString([]byte{0x00, 0x01, 0x02, 0x03}),
				PublicExponentBase64: base64.StdEncoding.EncodeToString([]byte{0x01, 0x00, 0x01}),
			},
			wantErr: false,
		},
		{
			name: "invalid modulus base64",
			pubKey: &PublicKey{
				ModulusBase64:        "not-valid-base64!!!",
				PublicExponentBase64: base64.StdEncoding.EncodeToString([]byte{0x01, 0x00, 0x01}),
			},
			wantErr: true,
			errMsg:  "failed to decode modulus",
		},
		{
			name: "invalid exponent base64",
			pubKey: &PublicKey{
				ModulusBase64:        base64.StdEncoding.EncodeToString([]byte{0x00, 0x01}),
				PublicExponentBase64: "not-valid-base64!!!",
			},
			wantErr: true,
			errMsg:  "failed to decode public exponent",
		},
		{
			name: "empty modulus",
			pubKey: &PublicKey{
				ModulusBase64:        base64.StdEncoding.EncodeToString([]byte{}),
				PublicExponentBase64: base64.StdEncoding.EncodeToString([]byte{0x01, 0x00, 0x01}),
			},
			wantErr: true,
			errMsg:  "modulus is empty",
		},
		{
			name: "empty exponent",
			pubKey: &PublicKey{
				ModulusBase64:        base64.StdEncoding.EncodeToString([]byte{0x01}),
				PublicExponentBase64: base64.StdEncoding.EncodeToString([]byte{}),
			},
			wantErr: true,
			errMsg:  "public exponent is empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := buildRSAPublicKey(tt.pubKey)

			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error containing %q, got nil", tt.errMsg)
				} else if !strings.Contains(err.Error(), tt.errMsg) {
					t.Errorf("expected error containing %q, got %q", tt.errMsg, err.Error())
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

func TestMaxPlaintextSize(t *testing.T) {
	_, pubKey := generateTestKeyPair(t, 2048)

	maxSize, err := MaxPlaintextSize(pubKey)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// For 2048-bit key with SHA-256: 256 - 66 = 190 bytes
	expectedMax := 256 - 66
	if maxSize != expectedMax {
		t.Errorf("max plaintext size: got %d, want %d", maxSize, expectedMax)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
