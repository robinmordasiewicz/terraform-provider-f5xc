// Copyright (c) F5XC Community
// SPDX-License-Identifier: MPL-2.0

package acctest

import (
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"strings"
	"testing"
)

func TestGenerateTestCertificates(t *testing.T) {
	certs, err := GenerateTestCertificates()
	if err != nil {
		t.Fatalf("failed to generate test certificates: %v", err)
	}

	// Verify Root CA
	t.Run("RootCA", func(t *testing.T) {
		if certs.RootCA == "" {
			t.Error("RootCA is empty")
		}
		if !strings.Contains(certs.RootCA, "BEGIN CERTIFICATE") {
			t.Error("RootCA is not in PEM format")
		}

		cert := parseCertPEM(t, certs.RootCA)
		if !cert.IsCA {
			t.Error("RootCA is not a CA certificate")
		}
		if cert.Subject.CommonName != "Test Root CA" {
			t.Errorf("RootCA CN = %q, want %q", cert.Subject.CommonName, "Test Root CA")
		}
	})

	// Verify Intermediate CA
	t.Run("IntermediateCA", func(t *testing.T) {
		if certs.IntermediateCA == "" {
			t.Error("IntermediateCA is empty")
		}
		if !strings.Contains(certs.IntermediateCA, "BEGIN CERTIFICATE") {
			t.Error("IntermediateCA is not in PEM format")
		}

		cert := parseCertPEM(t, certs.IntermediateCA)
		if !cert.IsCA {
			t.Error("IntermediateCA is not a CA certificate")
		}
		if cert.Subject.CommonName != "Test Intermediate CA" {
			t.Errorf("IntermediateCA CN = %q, want %q", cert.Subject.CommonName, "Test Intermediate CA")
		}
	})

	// Verify Server Certificate
	t.Run("ServerCert", func(t *testing.T) {
		if certs.ServerCert == "" {
			t.Error("ServerCert is empty")
		}
		if !strings.Contains(certs.ServerCert, "BEGIN CERTIFICATE") {
			t.Error("ServerCert is not in PEM format")
		}

		cert := parseCertPEM(t, certs.ServerCert)
		if cert.IsCA {
			t.Error("ServerCert should not be a CA certificate")
		}
		if cert.Subject.CommonName != "test.example.com" {
			t.Errorf("ServerCert CN = %q, want %q", cert.Subject.CommonName, "test.example.com")
		}
	})

	// Verify Server Key
	t.Run("ServerKey", func(t *testing.T) {
		if certs.ServerKey == "" {
			t.Error("ServerKey is empty")
		}
		if !strings.Contains(certs.ServerKey, "PRIVATE KEY") {
			t.Error("ServerKey is not in PEM format")
		}
	})

	// Verify Base64 encoded versions
	t.Run("Base64Encoding", func(t *testing.T) {
		// Verify IntermediateCABase64 decodes to the same PEM
		decoded, err := base64.StdEncoding.DecodeString(certs.IntermediateCABase64)
		if err != nil {
			t.Errorf("failed to decode IntermediateCABase64: %v", err)
		}
		if string(decoded) != certs.IntermediateCA {
			t.Error("IntermediateCABase64 does not decode to IntermediateCA")
		}

		// Verify ServerCertBase64 decodes to the same PEM
		decoded, err = base64.StdEncoding.DecodeString(certs.ServerCertBase64)
		if err != nil {
			t.Errorf("failed to decode ServerCertBase64: %v", err)
		}
		if string(decoded) != certs.ServerCert {
			t.Error("ServerCertBase64 does not decode to ServerCert")
		}

		// Verify ServerKeyBase64 decodes to the same PEM
		decoded, err = base64.StdEncoding.DecodeString(certs.ServerKeyBase64)
		if err != nil {
			t.Errorf("failed to decode ServerKeyBase64: %v", err)
		}
		if string(decoded) != certs.ServerKey {
			t.Error("ServerKeyBase64 does not decode to ServerKey")
		}
	})
}

func TestMustGenerateTestCertificates(t *testing.T) {
	// Should not panic
	certs := MustGenerateTestCertificates()
	if certs == nil {
		t.Error("MustGenerateTestCertificates returned nil")
	}
}

// parseCertPEM parses a PEM-encoded certificate and returns the x509 certificate
func parseCertPEM(t *testing.T, certPEM string) *x509.Certificate {
	t.Helper()
	block, _ := pem.Decode([]byte(certPEM))
	if block == nil {
		t.Fatal("failed to decode PEM block")
		return nil // staticcheck: unreachable but satisfies nil check analysis
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		t.Fatalf("failed to parse certificate: %v", err)
	}
	return cert
}
