// Copyright (c) F5XC Community
// SPDX-License-Identifier: MPL-2.0

package acctest

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/base64"
	"encoding/pem"
	"math/big"
	"time"
)

// TestCertificates holds generated test certificates for acceptance tests
type TestCertificates struct {
	// RootCA is the root CA certificate (PEM format)
	RootCA string
	// RootCAKey is the root CA private key (PEM format)
	RootCAKey string
	// IntermediateCA is the intermediate CA certificate (PEM format)
	IntermediateCA string
	// IntermediateCAKey is the intermediate CA private key (PEM format)
	IntermediateCAKey string
	// ServerCert is the server certificate (PEM format)
	ServerCert string
	// ServerKey is the server private key (PEM format)
	ServerKey string

	// Base64 encoded versions for F5 XC API (string:/// format)
	// RootCABase64 is the root CA certificate base64 encoded
	RootCABase64 string
	// IntermediateCABase64 is the intermediate CA certificate base64 encoded
	IntermediateCABase64 string
	// ServerCertBase64 is the server certificate base64 encoded
	ServerCertBase64 string
	// ServerKeyBase64 is the server private key base64 encoded
	ServerKeyBase64 string
}

// GenerateTestCertificates creates a complete certificate chain for testing:
// - Root CA (self-signed, CA:TRUE)
// - Intermediate CA (signed by Root CA, CA:TRUE, pathlen:0)
// - Server Certificate (signed by Intermediate CA, for test.example.com)
//
// All certificates are valid for 365 days from generation time.
// Returns TestCertificates with both PEM and Base64-encoded formats.
func GenerateTestCertificates() (*TestCertificates, error) {
	certs := &TestCertificates{}

	// Generate Root CA
	rootCAKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}

	rootCATemplate := &x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Country:            []string{"US"},
			Province:           []string{"Test"},
			Locality:           []string{"Test"},
			Organization:       []string{"TestCA"},
			OrganizationalUnit: []string{"Testing"},
			CommonName:         "Test Root CA",
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(1, 0, 0), // 1 year
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign | x509.KeyUsageCRLSign,
		BasicConstraintsValid: true,
		IsCA:                  true,
		MaxPathLen:            2,
	}

	rootCADER, err := x509.CreateCertificate(rand.Reader, rootCATemplate, rootCATemplate, &rootCAKey.PublicKey, rootCAKey)
	if err != nil {
		return nil, err
	}

	certs.RootCA = encodeCertToPEM(rootCADER)
	certs.RootCAKey = encodeKeyToPEM(rootCAKey)
	certs.RootCABase64 = base64.StdEncoding.EncodeToString([]byte(certs.RootCA))

	// Parse root CA for signing
	rootCACert, err := x509.ParseCertificate(rootCADER)
	if err != nil {
		return nil, err
	}

	// Generate Intermediate CA
	intermediateCAKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}

	intermediateCATemplate := &x509.Certificate{
		SerialNumber: big.NewInt(2),
		Subject: pkix.Name{
			Country:            []string{"US"},
			Province:           []string{"Test"},
			Locality:           []string{"Test"},
			Organization:       []string{"TestCA"},
			OrganizationalUnit: []string{"Testing"},
			CommonName:         "Test Intermediate CA",
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(1, 0, 0), // 1 year
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign | x509.KeyUsageCRLSign,
		BasicConstraintsValid: true,
		IsCA:                  true,
		MaxPathLenZero:        true, // pathlen:0
	}

	intermediateCADER, err := x509.CreateCertificate(rand.Reader, intermediateCATemplate, rootCACert, &intermediateCAKey.PublicKey, rootCAKey)
	if err != nil {
		return nil, err
	}

	certs.IntermediateCA = encodeCertToPEM(intermediateCADER)
	certs.IntermediateCAKey = encodeKeyToPEM(intermediateCAKey)
	certs.IntermediateCABase64 = base64.StdEncoding.EncodeToString([]byte(certs.IntermediateCA))

	// Parse intermediate CA for signing
	intermediateCACert, err := x509.ParseCertificate(intermediateCADER)
	if err != nil {
		return nil, err
	}

	// Generate Server Certificate
	serverKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}

	serverTemplate := &x509.Certificate{
		SerialNumber: big.NewInt(3),
		Subject: pkix.Name{
			Country:            []string{"US"},
			Province:           []string{"Test"},
			Locality:           []string{"Test"},
			Organization:       []string{"TestOrg"},
			OrganizationalUnit: []string{"Testing"},
			CommonName:         "test.example.com",
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(1, 0, 0), // 1 year
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		IsCA:                  false,
		DNSNames:              []string{"test.example.com"},
	}

	serverCertDER, err := x509.CreateCertificate(rand.Reader, serverTemplate, intermediateCACert, &serverKey.PublicKey, intermediateCAKey)
	if err != nil {
		return nil, err
	}

	certs.ServerCert = encodeCertToPEM(serverCertDER)
	certs.ServerKey = encodeKeyToPEM(serverKey)
	certs.ServerCertBase64 = base64.StdEncoding.EncodeToString([]byte(certs.ServerCert))
	certs.ServerKeyBase64 = base64.StdEncoding.EncodeToString([]byte(certs.ServerKey))

	return certs, nil
}

// encodeCertToPEM encodes a DER certificate to PEM format
func encodeCertToPEM(certDER []byte) string {
	var buf bytes.Buffer
	_ = pem.Encode(&buf, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certDER,
	})
	return buf.String()
}

// encodeKeyToPEM encodes an RSA private key to PEM format
func encodeKeyToPEM(key *rsa.PrivateKey) string {
	keyBytes, err := x509.MarshalPKCS8PrivateKey(key)
	if err != nil {
		// Fall back to PKCS1 format if PKCS8 fails
		keyBytes = x509.MarshalPKCS1PrivateKey(key)
	}
	var buf bytes.Buffer
	_ = pem.Encode(&buf, &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: keyBytes,
	})
	return buf.String()
}

// MustGenerateTestCertificates generates test certificates and panics on error.
// This is useful in test setup where error handling is not practical.
func MustGenerateTestCertificates() *TestCertificates {
	certs, err := GenerateTestCertificates()
	if err != nil {
		panic("failed to generate test certificates: " + err.Error())
	}
	return certs
}
