package blindfold

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/f5xc/terraform-provider-f5xc/internal/client"
)

func TestSizeAnalysis(t *testing.T) {
	apiURL := os.Getenv("F5XC_API_URL")
	p12File := os.Getenv("F5XC_P12_FILE")
	p12Password := os.Getenv("F5XC_P12_PASSWORD")

	if apiURL == "" || p12File == "" {
		t.Skip("Skipping: F5XC credentials not set")
	}

	c, err := client.NewClientWithP12(apiURL, p12File, p12Password)
	if err != nil {
		t.Fatalf("Error creating client: %v", err)
	}

	pubKey, err := GetPublicKey(context.Background(), c.HTTPClient, apiURL)
	if err != nil {
		t.Fatalf("Error getting public key: %v", err)
	}

	maxSize, err := MaxPlaintextSize(pubKey)
	if err != nil {
		t.Fatalf("Error getting max plaintext size: %v", err)
	}

	fmt.Printf("\n=== RSA Key Size Analysis ===\n")
	fmt.Printf("  Max plaintext size: %d bytes\n", maxSize)
	fmt.Printf("  Key version: %d\n", pubKey.KeyVersion)

	// Test encryption at exactly max size
	plaintext := []byte(strings.Repeat("A", maxSize))
	policy := &SecretPolicyDocument{
		PolicyID: "123",
	}

	sealed, err := Seal(plaintext, pubKey, policy)
	if err != nil {
		t.Fatalf("Failed to encrypt at max size (%d bytes): %v", maxSize, err)
	}
	fmt.Printf("  Successfully encrypted %d bytes\n", maxSize)
	fmt.Printf("  Sealed output length: %d chars\n", len(sealed))

	// Test encryption at max+1 to confirm limit
	plaintextTooLarge := []byte(strings.Repeat("A", maxSize+1))
	_, err = Seal(plaintextTooLarge, pubKey, policy)
	if err == nil {
		t.Logf("WARNING: Encryption succeeded at max+1 size (%d bytes) - limit may be wrong", maxSize+1)
	} else {
		fmt.Printf("  Correctly rejected %d bytes: %v\n", maxSize+1, err)
	}
}

func TestEncryptVariousSizes(t *testing.T) {
	apiURL := os.Getenv("F5XC_API_URL")
	p12File := os.Getenv("F5XC_P12_FILE")
	p12Password := os.Getenv("F5XC_P12_PASSWORD")

	if apiURL == "" || p12File == "" {
		t.Skip("Skipping: F5XC credentials not set")
	}

	c, err := client.NewClientWithP12(apiURL, p12File, p12Password)
	if err != nil {
		t.Fatalf("Error creating client: %v", err)
	}

	pubKey, err := GetPublicKey(context.Background(), c.HTTPClient, apiURL)
	if err != nil {
		t.Fatalf("Error getting public key: %v", err)
	}

	policy := &SecretPolicyDocument{
		PolicyID: "123",
	}

	testSizes := []int{10, 50, 100, 150, 180, 190, 191, 200, 250, 300, 500, 1000}

	fmt.Printf("\n=== Encryption Size Tests ===\n")
	for _, size := range testSizes {
		plaintext := []byte(strings.Repeat("X", size))
		_, err := Seal(plaintext, pubKey, policy)
		status := "✅"
		if err != nil {
			status = "❌"
		}
		fmt.Printf("  %s Size %4d bytes: %v\n", status, size, err)
	}
}

func TestTypicalSecretSizes(t *testing.T) {
	// Show sizes of typical secrets
	// NOTE: These are NOT real secrets - they are example strings to document typical sizes
	fmt.Printf("\n=== Typical Secret Sizes ===\n")

	// API key (example format, not real) pragma: allowlist secret
	apiKey := "sk-12345678901234567890123456789012345678901234567890" // pragma: allowlist secret
	fmt.Printf("  API key (50 chars): %d bytes\n", len(apiKey))

	// Database password (example format, not real) pragma: allowlist secret
	dbPassword := "MySecureP@ssw0rd!123" // pragma: allowlist secret
	fmt.Printf("  DB password (20 chars): %d bytes\n", len(dbPassword))

	// AWS Secret Key (AWS example key from docs) pragma: allowlist secret
	awsSecret := "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY" // pragma: allowlist secret
	fmt.Printf("  AWS secret key (40 chars): %d bytes\n", len(awsSecret))

	// PEM header line length (any PEM type)
	pemHeaderLen := len("-----BEGIN CERTIFICATE-----") // 27 bytes
	fmt.Printf("  PEM header line: %d bytes\n", pemHeaderLen)

	// Typical RSA-2048 private key PEM is ~1700 bytes
	fmt.Printf("  RSA-2048 private key PEM: ~1700 bytes\n")

	// Typical EC P-256 private key PEM is ~227 bytes
	fmt.Printf("  EC P-256 private key PEM: ~227 bytes\n")

	// Typical certificate is ~1200-2000 bytes
	fmt.Printf("  TLS certificate PEM: ~1200-2000 bytes\n")

	fmt.Printf("\n  RSA-OAEP with SHA-256 max: 190 bytes (RSA-2048)\n")
	fmt.Printf("  Conclusion: PEM files CANNOT be directly encrypted with RSA-OAEP\n")
	fmt.Printf("  F5XC vesctl must use hybrid encryption (AES + RSA-wrap) internally\n")
}
