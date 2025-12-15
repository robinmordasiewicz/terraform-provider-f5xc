// This file is MANUALLY MAINTAINED and is NOT auto-generated.

package blindfold

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
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
			// Verify envelope encryption fields are present
			if sealedSecret.EncryptedKey == "" {
				t.Error("encrypted_key should not be empty")
			}
			if sealedSecret.Nonce == "" {
				t.Error("nonce should not be empty")
			}
			if sealedSecret.Ciphertext == "" {
				t.Error("ciphertext should not be empty")
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

	// With envelope encryption, max size is now the API limit (128KB)
	if maxSize != MaxSecretSize {
		t.Errorf("max plaintext size: got %d, want %d", maxSize, MaxSecretSize)
	}
}

func TestMaxRSAPlaintextSize(t *testing.T) {
	_, pubKey := generateTestKeyPair(t, 2048)

	maxSize, err := MaxRSAPlaintextSize(pubKey)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// For 2048-bit key with SHA-256: 256 - 66 = 190 bytes
	expectedMax := 256 - 66
	if maxSize != expectedMax {
		t.Errorf("max RSA plaintext size: got %d, want %d", maxSize, expectedMax)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// TestSealOutputFormat validates that the sealed output matches F5XC expected format.
// The format is: string:///<base64-encoded-json>
// where the JSON contains envelope encryption fields: key_version, policy_id, tenant, encrypted_key, nonce, ciphertext
func TestSealOutputFormat(t *testing.T) {
	_, pubKey := generateTestKeyPair(t, 2048)
	policy := &SecretPolicyDocument{
		Name:      "ves-io-allow-volterra",
		Namespace: "shared",
		PolicyID:  "policy-id-12345",
		Tenant:    "test-tenant-xyz",
	}

	sealed, err := Seal([]byte("test-secret"), pubKey, policy)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Validate prefix
	if !strings.HasPrefix(sealed, "string:///") {
		t.Errorf("output must start with 'string:///', got prefix: %q", sealed[:min(20, len(sealed))])
	}

	// Extract and decode base64 payload
	payload := strings.TrimPrefix(sealed, "string:///")
	decoded, err := base64.StdEncoding.DecodeString(payload)
	if err != nil {
		t.Fatalf("payload must be valid base64: %v", err)
	}

	// Validate JSON structure matches SealedSecret
	var result SealedSecret
	if err := json.Unmarshal(decoded, &result); err != nil {
		t.Fatalf("payload must be valid JSON: %v", err)
	}

	// Validate required fields
	if result.KeyVersion != pubKey.KeyVersion {
		t.Errorf("key_version mismatch: got %d, want %d", result.KeyVersion, pubKey.KeyVersion)
	}
	if result.PolicyID != policy.PolicyID {
		t.Errorf("policy_id mismatch: got %q, want %q", result.PolicyID, policy.PolicyID)
	}
	if result.Tenant != pubKey.Tenant {
		t.Errorf("tenant mismatch: got %q, want %q", result.Tenant, pubKey.Tenant)
	}

	// Validate envelope encryption fields
	if result.EncryptedKey == "" {
		t.Error("encrypted_key field must not be empty")
	}
	if result.Nonce == "" {
		t.Error("nonce field must not be empty")
	}
	if result.Ciphertext == "" {
		t.Error("ciphertext field must not be empty")
	}

	// Validate encrypted_key is valid base64 (RSA-OAEP encrypted AES key)
	encKey, err := base64.StdEncoding.DecodeString(result.EncryptedKey)
	if err != nil {
		t.Errorf("encrypted_key field must be valid base64: %v", err)
	}
	// RSA-2048 produces 256 bytes of ciphertext
	if len(encKey) != 256 {
		t.Errorf("encrypted_key length: got %d, want 256 (RSA-2048)", len(encKey))
	}

	// Validate nonce is valid base64 (12 bytes for GCM)
	nonceBytes, err := base64.StdEncoding.DecodeString(result.Nonce)
	if err != nil {
		t.Errorf("nonce field must be valid base64: %v", err)
	}
	if len(nonceBytes) != GCMNonceSize {
		t.Errorf("nonce length: got %d, want %d", len(nonceBytes), GCMNonceSize)
	}

	// Validate ciphertext is valid base64
	_, err = base64.StdEncoding.DecodeString(result.Ciphertext)
	if err != nil {
		t.Errorf("ciphertext field must be valid base64: %v", err)
	}
}

// TestSealJSONFieldNames verifies the exact JSON field names used in output.
// With envelope encryption: key_version, policy_id, tenant, encrypted_key, nonce, ciphertext
func TestSealJSONFieldNames(t *testing.T) {
	_, pubKey := generateTestKeyPair(t, 2048)
	policy := &SecretPolicyDocument{
		PolicyID: "policy-123",
		Tenant:   "test-tenant",
	}

	sealed, err := Seal([]byte("test"), pubKey, policy)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	payload := strings.TrimPrefix(sealed, "string:///")
	decoded, err := base64.StdEncoding.DecodeString(payload)
	if err != nil {
		t.Fatalf("failed to decode: %v", err)
	}

	// Parse as generic map to check exact field names
	var rawJSON map[string]interface{}
	if err := json.Unmarshal(decoded, &rawJSON); err != nil {
		t.Fatalf("failed to parse JSON: %v", err)
	}

	requiredFields := []string{"key_version", "policy_id", "tenant", "encrypted_key", "nonce", "ciphertext"}
	for _, field := range requiredFields {
		if _, exists := rawJSON[field]; !exists {
			t.Errorf("missing required field: %q", field)
		}
	}

	// Ensure no extra fields
	if len(rawJSON) != len(requiredFields) {
		t.Errorf("expected %d fields, got %d: %v", len(requiredFields), len(rawJSON), rawJSON)
	}
}

// TestSealNonDeterministic verifies envelope encryption is non-deterministic.
// Same input encrypted twice should produce different ciphertext (due to random DEK and nonce).
func TestSealNonDeterministic(t *testing.T) {
	_, pubKey := generateTestKeyPair(t, 2048)
	policy := &SecretPolicyDocument{
		PolicyID: "policy-123",
	}

	plaintext := []byte("same-secret-value")

	sealed1, err := Seal(plaintext, pubKey, policy)
	if err != nil {
		t.Fatalf("first seal failed: %v", err)
	}

	sealed2, err := Seal(plaintext, pubKey, policy)
	if err != nil {
		t.Fatalf("second seal failed: %v", err)
	}

	if sealed1 == sealed2 {
		t.Error("envelope encryption should be non-deterministic; same plaintext produced identical output")
	}

	// Extract and compare the encrypted fields specifically
	extract := func(sealed string) (string, string, string) {
		payload := strings.TrimPrefix(sealed, "string:///")
		decoded, _ := base64.StdEncoding.DecodeString(payload)
		var result SealedSecret
		json.Unmarshal(decoded, &result)
		return result.EncryptedKey, result.Nonce, result.Ciphertext
	}

	key1, nonce1, ct1 := extract(sealed1)
	key2, nonce2, ct2 := extract(sealed2)

	// Each call generates a new random DEK
	if key1 == key2 {
		t.Error("encrypted_key should differ between calls due to random DEK")
	}

	// Each call generates a new random nonce
	if nonce1 == nonce2 {
		t.Error("nonce should differ between calls due to random generation")
	}

	// Ciphertext should differ due to different DEK/nonce
	if ct1 == ct2 {
		t.Error("ciphertext should differ between calls")
	}
}

// TestSealLargePlaintext verifies handling of large plaintexts with envelope encryption.
// With envelope encryption, we can now encrypt data up to 128KB (API limit).
func TestSealLargePlaintext(t *testing.T) {
	_, pubKey := generateTestKeyPair(t, 2048)
	policy := &SecretPolicyDocument{
		PolicyID: "policy-123",
	}

	// Test various large sizes that would fail with direct RSA-OAEP
	testSizes := []int{
		200,    // Just over old RSA limit (190)
		500,    // Typical API key size
		1700,   // RSA-2048 private key PEM
		2000,   // TLS certificate
		10000,  // 10KB
		50000,  // 50KB
		100000, // 100KB
	}

	for _, size := range testSizes {
		t.Run(strings.ReplaceAll(fmt.Sprintf("%d_bytes", size), "", ""), func(t *testing.T) {
			plaintext := make([]byte, size)
			for i := range plaintext {
				plaintext[i] = byte(i % 256)
			}

			sealed, err := Seal(plaintext, pubKey, policy)
			if err != nil {
				t.Errorf("should accept plaintext of %d bytes: %v", size, err)
				return
			}

			// Verify the output is valid
			if !strings.HasPrefix(sealed, "string:///") {
				t.Error("output format broken")
			}
		})
	}

	// Test exactly at max size (128KB)
	t.Run("exact max size 128KB", func(t *testing.T) {
		plaintext := make([]byte, MaxSecretSize)
		for i := range plaintext {
			plaintext[i] = byte(i % 256)
		}

		_, err := Seal(plaintext, pubKey, policy)
		if err != nil {
			t.Errorf("should accept plaintext at max size (%d bytes): %v", MaxSecretSize, err)
		}
	})

	// Test one byte over max size (should fail)
	t.Run("one over max size", func(t *testing.T) {
		plaintext := make([]byte, MaxSecretSize+1)
		_, err := Seal(plaintext, pubKey, policy)
		if err == nil {
			t.Error("should reject plaintext exceeding max size")
		}
		if !strings.Contains(err.Error(), "exceeds maximum allowed size") {
			t.Errorf("unexpected error message: %v", err)
		}
	})
}

// TestSealSpecialCharacters verifies handling of various special characters in plaintext.
func TestSealSpecialCharacters(t *testing.T) {
	_, pubKey := generateTestKeyPair(t, 2048)
	policy := &SecretPolicyDocument{
		PolicyID: "policy-123",
	}

	testCases := []struct {
		name      string
		plaintext []byte
	}{
		{"null bytes", []byte("before\x00after")},
		{"unicode", []byte("héllo wörld 你好")},
		{"newlines", []byte("line1\nline2\r\nline3")},
		{"json special", []byte(`{"key": "value", "nested": {"a": 1}}`)},
		{"base64 chars", []byte("abc+/=XYZ")},
		{"all printable ASCII", func() []byte {
			b := make([]byte, 95)
			for i := 0; i < 95; i++ {
				b[i] = byte(32 + i) // space through ~
			}
			return b
		}()},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sealed, err := Seal(tc.plaintext, pubKey, policy)
			if err != nil {
				t.Fatalf("failed to seal: %v", err)
			}

			// Verify format is still valid
			if !strings.HasPrefix(sealed, "string:///") {
				t.Error("output format broken")
			}

			// Verify we can decode and parse
			payload := strings.TrimPrefix(sealed, "string:///")
			decoded, err := base64.StdEncoding.DecodeString(payload)
			if err != nil {
				t.Fatalf("output not valid base64: %v", err)
			}

			var result SealedSecret
			if err := json.Unmarshal(decoded, &result); err != nil {
				t.Fatalf("output not valid JSON: %v", err)
			}
		})
	}
}

// TestMaxRSAPlaintextSize_VariousKeySizes tests RSA max plaintext calculation for different key sizes.
func TestMaxRSAPlaintextSize_VariousKeySizes(t *testing.T) {
	tests := []struct {
		keyBits     int
		wantMaxSize int // keySize - 2*hashSize - 2 = keySize - 66 for SHA-256
	}{
		{2048, 256 - 66}, // 190 bytes
		{3072, 384 - 66}, // 318 bytes
		{4096, 512 - 66}, // 446 bytes
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d_bit", tt.keyBits), func(t *testing.T) {
			_, pubKey := generateTestKeyPair(t, tt.keyBits)

			maxSize, err := MaxRSAPlaintextSize(pubKey)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if maxSize != tt.wantMaxSize {
				t.Errorf("max RSA size for %d-bit key: got %d, want %d", tt.keyBits, maxSize, tt.wantMaxSize)
			}
		})
	}
}

// TestMaxPlaintextSize_AllKeySizes verifies MaxPlaintextSize returns API limit for all key sizes.
func TestMaxPlaintextSize_AllKeySizes(t *testing.T) {
	keySizes := []int{2048, 3072, 4096}

	for _, bits := range keySizes {
		t.Run(fmt.Sprintf("%d_bit", bits), func(t *testing.T) {
			_, pubKey := generateTestKeyPair(t, bits)

			maxSize, err := MaxPlaintextSize(pubKey)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			// With envelope encryption, all key sizes support the same max (API limit)
			if maxSize != MaxSecretSize {
				t.Errorf("max size for %d-bit key: got %d, want %d", bits, maxSize, MaxSecretSize)
			}
		})
	}
}
