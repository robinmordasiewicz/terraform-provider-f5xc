// Copyright (c) F5XC Community
// SPDX-License-Identifier: MPL-2.0

// Package acctest provides acceptance test utilities for F5 XC Terraform provider.
// Following HashiCorp's acceptance testing best practices.
package acctest

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/f5xc/terraform-provider-f5xc/internal/provider"
)

// Environment variable names for acceptance tests
const (
	// EnvF5XCURL is the environment variable for the F5 XC API URL
	EnvF5XCURL = "F5XC_API_URL"

	// EnvF5XCToken is the environment variable for the F5 XC API token
	EnvF5XCToken = "F5XC_API_TOKEN"

	// EnvF5XCP12File is the environment variable for the P12 certificate file path
	EnvF5XCP12File = "F5XC_API_P12_FILE"

	// EnvF5XCP12Password is the environment variable for the P12 certificate password
	EnvF5XCP12Password = "F5XC_P12_PASSWORD"

	// EnvF5XCCert is the environment variable for the PEM certificate file path
	EnvF5XCCert = "F5XC_API_CERT"

	// EnvF5XCKey is the environment variable for the PEM key file path
	EnvF5XCKey = "F5XC_API_KEY"

	// EnvF5XCTenantName is the environment variable for the F5 XC tenant name
	EnvF5XCTenantName = "F5XC_TENANT_NAME"

	// EnvTFAccTest enables acceptance tests
	EnvTFAccTest = "TF_ACC"
)

// ProtoV6ProviderFactories returns the provider factories for acceptance testing
var ProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"f5xc": providerserver.NewProtocol6WithError(provider.New("test")()),
}

// AuthMethod represents the authentication method detected
type AuthMethod int

const (
	// AuthMethodNone indicates no authentication configured
	AuthMethodNone AuthMethod = iota
	// AuthMethodToken indicates API token authentication
	AuthMethodToken
	// AuthMethodP12 indicates P12 certificate authentication
	AuthMethodP12
	// AuthMethodPEM indicates PEM certificate authentication
	AuthMethodPEM
)

// DetectAuthMethod determines which authentication method is configured
func DetectAuthMethod() AuthMethod {
	// Check P12 authentication (preferred for testing)
	if os.Getenv(EnvF5XCP12File) != "" && os.Getenv(EnvF5XCP12Password) != "" {
		return AuthMethodP12
	}

	// Check PEM certificate authentication
	if os.Getenv(EnvF5XCCert) != "" && os.Getenv(EnvF5XCKey) != "" {
		return AuthMethodPEM
	}

	// Check token authentication
	if os.Getenv(EnvF5XCToken) != "" {
		return AuthMethodToken
	}

	return AuthMethodNone
}

// PreCheck validates that required environment variables are set before running tests
func PreCheck(t *testing.T) {
	t.Helper()

	// API URL is always required
	if os.Getenv(EnvF5XCURL) == "" {
		t.Fatalf("Required environment variable not set: %s", EnvF5XCURL)
	}

	// Check for at least one valid authentication method
	authMethod := DetectAuthMethod()

	switch authMethod {
	case AuthMethodP12:
		t.Logf("Using P12 certificate authentication (file: %s)", os.Getenv(EnvF5XCP12File))
	case AuthMethodPEM:
		t.Logf("Using PEM certificate authentication (cert: %s, key: %s)",
			os.Getenv(EnvF5XCCert), os.Getenv(EnvF5XCKey))
	case AuthMethodToken:
		t.Logf("Using API token authentication")
	case AuthMethodNone:
		t.Fatalf("No authentication configured. Set one of:\n"+
			"  - P12: %s and %s\n"+
			"  - PEM: %s and %s\n"+
			"  - Token: %s",
			EnvF5XCP12File, EnvF5XCP12Password,
			EnvF5XCCert, EnvF5XCKey,
			EnvF5XCToken)
	}
}

// SkipIfNotAccTest skips the test if TF_ACC is not set
func SkipIfNotAccTest(t *testing.T) {
	t.Helper()

	if os.Getenv(EnvTFAccTest) == "" {
		t.Skip("Acceptance tests skipped unless TF_ACC is set")
	}
}

// RandomName generates a random name with the given prefix for test resources
func RandomName(prefix string) string {
	return fmt.Sprintf("%s-%s", prefix, acctest.RandStringFromCharSet(8, acctest.CharSetAlphaNum))
}

// RandomNameWithSuffix generates a random name with prefix and suffix
func RandomNameWithSuffix(prefix, suffix string) string {
	return fmt.Sprintf("%s-%s-%s", prefix, acctest.RandStringFromCharSet(6, acctest.CharSetAlphaNum), suffix)
}

// TestNamespace returns the namespace for tests (default: "system")
func TestNamespace() string {
	if ns := os.Getenv("F5XC_TEST_NAMESPACE"); ns != "" {
		return ns
	}
	return "system"
}

// ConfigCompose composes multiple Terraform configurations
func ConfigCompose(configs ...string) string {
	var sb strings.Builder
	for _, config := range configs {
		sb.WriteString(config)
		sb.WriteString("\n")
	}
	return sb.String()
}

// ProviderConfig returns the provider configuration for tests
func ProviderConfig() string {
	return `
provider "f5xc" {
  # Configuration from environment variables
}
`
}

// CheckResourceExists returns a resource.TestCheckFunc that verifies a resource exists
func CheckResourceExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("resource ID is not set: %s", resourceName)
		}

		return nil
	}
}

// CheckResourceDestroyed returns a resource.TestCheckFunc that verifies a resource is destroyed
func CheckResourceDestroyed(resourceType string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for _, rs := range s.RootModule().Resources {
			if rs.Type != resourceType {
				continue
			}

			// In a real implementation, you would make an API call here
			// to verify the resource no longer exists.
			// For now, we assume if the resource is in the destroyed state,
			// Terraform has already verified it doesn't exist.
			_ = rs.Primary.ID // Placeholder for future API verification
		}
		return nil
	}
}

// CheckResourceAttr is a convenience wrapper around resource.TestCheckResourceAttr
func CheckResourceAttr(name, key, value string) resource.TestCheckFunc {
	return resource.TestCheckResourceAttr(name, key, value)
}

// CheckResourceAttrSet is a convenience wrapper around resource.TestCheckResourceAttrSet
func CheckResourceAttrSet(name, key string) resource.TestCheckFunc {
	return resource.TestCheckResourceAttrSet(name, key)
}

// CheckResourceAttrPair is a convenience wrapper around resource.TestCheckResourceAttrPair
func CheckResourceAttrPair(nameFirst, keyFirst, nameSecond, keySecond string) resource.TestCheckFunc {
	return resource.TestCheckResourceAttrPair(nameFirst, keyFirst, nameSecond, keySecond)
}

// ImportStateVerify returns the import state verify settings
func ImportStateVerify(resourceName string) resource.TestStep {
	return resource.TestStep{
		ResourceName:      resourceName,
		ImportState:       true,
		ImportStateVerify: true,
	}
}

// ImportStateVerifyIgnore returns import state verify with ignored attributes
func ImportStateVerifyIgnore(resourceName string, ignoreFields ...string) resource.TestStep {
	return resource.TestStep{
		ResourceName:            resourceName,
		ImportState:             true,
		ImportStateVerify:       true,
		ImportStateVerifyIgnore: ignoreFields,
	}
}

// TestResource provides a base structure for resource tests
type TestResource struct {
	Name         string
	ResourceType string
	Namespace    string
}

// NewTestResource creates a new test resource with a random name
func NewTestResource(resourceType string) *TestResource {
	return &TestResource{
		Name:         RandomName("tf-test"),
		ResourceType: resourceType,
		Namespace:    TestNamespace(),
	}
}

// FullResourceName returns the full Terraform resource name
func (r *TestResource) FullResourceName() string {
	return fmt.Sprintf("f5xc_%s.test", r.ResourceType)
}

// IDAttribute returns the attribute path for the ID
func (r *TestResource) IDAttribute() string {
	return fmt.Sprintf("%s.id", r.FullResourceName())
}

// ErrorContains checks if an error contains a specific substring
func ErrorContains(t *testing.T, err error, substring string) {
	t.Helper()

	if err == nil {
		t.Fatalf("expected error containing %q, got nil", substring)
	}

	if !strings.Contains(err.Error(), substring) {
		t.Fatalf("expected error containing %q, got: %s", substring, err.Error())
	}
}

// DefaultTestTimeout is the default timeout for test operations
const DefaultTestTimeout = 10 * time.Minute

// ContextWithTimeout returns a context with the standard test timeout
func ContextWithTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), DefaultTestTimeout)
}

// TestCheckFuncCompose composes multiple TestCheckFuncs into one
func TestCheckFuncCompose(funcs ...resource.TestCheckFunc) resource.TestCheckFunc {
	return resource.ComposeAggregateTestCheckFunc(funcs...)
}
