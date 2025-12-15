// This file is MANUALLY MAINTAINED and is NOT auto-generated.

package blindfold

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/big"
)

// Seal encrypts plaintext using RSA-OAEP with SHA-256 and the provided public key.
// The result is a base64-encoded JSON structure suitable for use in
// blindfold_secret_info.location fields with the "string:///" prefix.
//
// Parameters:
//   - plaintext: Raw bytes to encrypt (not base64-encoded)
//   - pubKey: Public key fetched from F5XC Secret Management API
//   - policy: Policy document defining decryption access control
//
// Returns:
//   - Sealed secret string in format "string:///<base64-encoded-sealed-json>"
//   - Error if encryption fails
func Seal(plaintext []byte, pubKey *PublicKey, policy *SecretPolicyDocument) (string, error) {
	if pubKey == nil {
		return "", fmt.Errorf("public key is required")
	}
	if policy == nil {
		return "", fmt.Errorf("policy document is required")
	}
	if len(plaintext) == 0 {
		return "", fmt.Errorf("plaintext cannot be empty")
	}

	// Build RSA public key from components
	rsaPubKey, err := buildRSAPublicKey(pubKey)
	if err != nil {
		return "", fmt.Errorf("failed to build RSA public key: %w", err)
	}

	// Encrypt using RSA-OAEP with SHA-256
	ciphertext, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		rsaPubKey,
		plaintext,
		nil, // No label
	)
	if err != nil {
		return "", fmt.Errorf("RSA-OAEP encryption failed: %w", err)
	}

	// Create sealed secret structure
	sealed := SealedSecret{
		KeyVersion: pubKey.KeyVersion,
		PolicyID:   policy.PolicyID,
		Tenant:     pubKey.Tenant,
		Data:       base64.StdEncoding.EncodeToString(ciphertext),
	}

	// Serialize to JSON
	sealedJSON, err := json.Marshal(sealed)
	if err != nil {
		return "", fmt.Errorf("failed to serialize sealed secret: %w", err)
	}

	// Return in the format expected by F5XC: string:///<base64>
	return "string:///" + base64.StdEncoding.EncodeToString(sealedJSON), nil
}

// SealBase64 is a convenience function that accepts base64-encoded plaintext.
// This is useful when the input is already base64-encoded (e.g., from Terraform's
// base64encode function).
func SealBase64(plaintextBase64 string, pubKey *PublicKey, policy *SecretPolicyDocument) (string, error) {
	plaintext, err := base64.StdEncoding.DecodeString(plaintextBase64)
	if err != nil {
		return "", fmt.Errorf("failed to decode base64 plaintext: %w", err)
	}
	return Seal(plaintext, pubKey, policy)
}

// buildRSAPublicKey constructs an RSA public key from the F5XC public key components.
// F5XC uses a custom format where the exponent bytes may include a version prefix byte (0x03).
// When present, this prefix must be stripped to extract the actual RSA exponent.
func buildRSAPublicKey(pubKey *PublicKey) (*rsa.PublicKey, error) {
	// Decode modulus from base64
	modulusBytes, err := base64.StdEncoding.DecodeString(pubKey.ModulusBase64)
	if err != nil {
		return nil, fmt.Errorf("failed to decode modulus: %w", err)
	}
	if len(modulusBytes) == 0 {
		return nil, fmt.Errorf("modulus is empty")
	}

	// Decode public exponent from base64
	exponentBytes, err := base64.StdEncoding.DecodeString(pubKey.PublicExponentBase64)
	if err != nil {
		return nil, fmt.Errorf("failed to decode public exponent: %w", err)
	}
	if len(exponentBytes) == 0 {
		return nil, fmt.Errorf("public exponent is empty")
	}

	// F5XC exponent format: The first byte (0x03) is a version/type indicator
	// that must be stripped to get the actual RSA exponent value.
	// Example: base64 "AzzBVP0=" decodes to [0x03, 0x3c, 0xc1, 0x54, 0xfd]
	//          With prefix: 13904205053 (doesn't fit in int32)
	//          Without prefix (0x3cc154fd): 1019303165 (valid RSA exponent)
	actualExponentBytes := exponentBytes
	if len(exponentBytes) > 1 && exponentBytes[0] == 0x03 {
		actualExponentBytes = exponentBytes[1:]
	}

	// Convert to big integers
	modulus := new(big.Int).SetBytes(modulusBytes)
	exponent := new(big.Int).SetBytes(actualExponentBytes)

	// Validate exponent fits in int (should be 65537 typically, but F5XC uses larger values)
	if !exponent.IsInt64() {
		return nil, fmt.Errorf("public exponent too large to fit in int64")
	}

	// Check Go crypto/rsa limit (must be <= 2^31-1)
	maxExponent := int64(1<<31 - 1)
	if exponent.Int64() > maxExponent {
		return nil, fmt.Errorf("public exponent %d exceeds Go crypto/rsa limit (%d)", exponent.Int64(), maxExponent)
	}

	e := int(exponent.Int64())
	if e <= 0 {
		return nil, fmt.Errorf("public exponent must be positive")
	}

	return &rsa.PublicKey{
		N: modulus,
		E: e,
	}, nil
}

// MaxPlaintextSize returns the maximum plaintext size that can be encrypted
// with the given public key using RSA-OAEP with SHA-256.
// For RSA-OAEP: max = keySize - 2*hashSize - 2
// With SHA-256 (32 bytes): max = keySize - 66
func MaxPlaintextSize(pubKey *PublicKey) (int, error) {
	rsaPubKey, err := buildRSAPublicKey(pubKey)
	if err != nil {
		return 0, err
	}

	keySize := rsaPubKey.Size() // Size in bytes
	hashSize := sha256.Size     // 32 bytes for SHA-256

	// RSA-OAEP overhead: 2*hashSize + 2
	maxSize := keySize - 2*hashSize - 2
	if maxSize <= 0 {
		return 0, fmt.Errorf("key too small for RSA-OAEP with SHA-256")
	}

	return maxSize, nil
}
