// Package blindfold provides F5 Distributed Cloud Secret Management encryption.
//
// This package implements the blindfold encryption system used by F5XC to protect
// sensitive data such as TLS private keys, passwords, and other secrets. The
// encryption uses RSA-OAEP with SHA-256 and is compatible with the F5XC Wingman
// sidecar for decryption.
//
// This package is MANUALLY MAINTAINED and is NOT auto-generated from OpenAPI
// specifications. Changes to this package should be committed directly.
//
// # Architecture
//
// The blindfold system consists of three main components:
//
//   - Public Key: RSA public key fetched from F5XC Secret Management API
//   - Policy Document: Access control policy defining who can decrypt secrets
//   - Sealed Secret: Encrypted data that can only be decrypted by authorized clients
//
// # Usage
//
// Typical usage involves fetching the public key and policy document from F5XC,
// then using the Seal function to encrypt plaintext:
//
//	pubKey, err := blindfold.GetPublicKey(ctx, httpClient, baseURL)
//	policy, err := blindfold.GetSecretPolicyDocument(ctx, httpClient, baseURL, namespace, name)
//	sealed, err := blindfold.Seal(plaintext, pubKey, policy)
//
// The sealed output can then be used in Terraform configurations with
// blindfold_secret_info.location fields.
//
// # Security Considerations
//
// - Plaintext secrets are never transmitted to F5XC during encryption
// - Encryption happens locally using the public key
// - Only authorized clients (per the policy document) can decrypt
// - The sealed format includes key version for rotation support
package blindfold
