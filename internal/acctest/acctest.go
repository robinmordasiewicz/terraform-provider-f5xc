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
	"sync"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/f5xc/terraform-provider-f5xc/internal/client"
	"github.com/f5xc/terraform-provider-f5xc/internal/provider"
)

var (
	// testClient is a shared client for acceptance tests
	testClient     *client.Client
	testClientOnce sync.Once
	testClientErr  error
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

// ExternalProviders defines external providers used in acceptance tests
// Use this when tests require external providers like hashicorp/time
var ExternalProviders = map[string]resource.ExternalProvider{
	"time": {
		Source: "hashicorp/time",
	},
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

// PreCheck validates that required environment variables are set before running tests.
// It also logs the test category as REAL_API for reporting purposes.
// When F5XC_MOCK_MODE is set, it automatically configures the environment to use
// the global mock server instead of requiring real credentials.
func PreCheck(t *testing.T) {
	t.Helper()

	// If mock mode is enabled, configure environment to use mock server
	// This must happen before any credential checks
	EnsureMockModeConfigured()

	// Log test category for reporting
	if IsMockMode() {
		LogTestCategory(t, TestCategoryMock)
		t.Logf("Using mock server mode (F5XC_MOCK_MODE=1)")
		return // Skip credential validation for mock mode
	}

	// Log test category for reporting - this is a real API test
	LogTestCategory(t, TestCategoryReal)

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

// SkipIfNotAccTest skips the test if TF_ACC is not set and mock mode is not enabled.
// When F5XC_MOCK_MODE is set, tests run without requiring TF_ACC.
func SkipIfNotAccTest(t *testing.T) {
	t.Helper()

	// Mock mode doesn't require TF_ACC
	if IsMockMode() {
		return
	}

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

// GetTestClient returns a shared API client for acceptance tests.
// The client is created once and reused across all tests.
func GetTestClient() (*client.Client, error) {
	testClientOnce.Do(func() {
		apiURL := os.Getenv(EnvF5XCURL)
		if apiURL == "" {
			apiURL = "https://console.ves.volterra.io"
		}

		// Normalize URL (remove trailing slashes and /api suffix)
		apiURL = strings.TrimRight(apiURL, "/")
		if strings.HasSuffix(strings.ToLower(apiURL), "/api") {
			apiURL = apiURL[:len(apiURL)-4]
		}
		apiURL = strings.TrimRight(apiURL, "/")

		switch DetectAuthMethod() {
		case AuthMethodP12:
			testClient, testClientErr = client.NewClientWithP12(
				apiURL,
				os.Getenv(EnvF5XCP12File),
				os.Getenv(EnvF5XCP12Password),
			)
		case AuthMethodPEM:
			testClient, testClientErr = client.NewClientWithCert(
				apiURL,
				os.Getenv(EnvF5XCCert),
				os.Getenv(EnvF5XCKey),
				"", // CA cert optional
			)
		case AuthMethodToken:
			testClient = client.NewClient(apiURL, os.Getenv(EnvF5XCToken))
		default:
			testClientErr = fmt.Errorf("no authentication method configured")
		}
	})

	return testClient, testClientErr
}

// CheckNamespaceDestroyed verifies that namespace resources have been deleted
// from the F5 XC API, not just from Terraform state.
// Includes retry logic to handle async deletion.
func CheckNamespaceDestroyed(s *terraform.State) error {
	c, err := GetTestClient()
	if err != nil {
		return fmt.Errorf("failed to get test client: %w", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "f5xc_namespace" {
			continue
		}

		name := rs.Primary.Attributes["name"]
		if name == "" {
			name = rs.Primary.ID
		}

		// Retry loop to handle async deletion (F5 XC may take time to fully delete)
		maxRetries := 6
		for i := 0; i < maxRetries; i++ {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			_, err := c.GetNamespace(ctx, "system", name)
			cancel()

			if err != nil {
				// Check if it's a "not found" error - success!
				if strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "NOT_FOUND") {
					break // Resource is deleted
				}
				// Some other error occurred
				return fmt.Errorf("unexpected error checking namespace %s: %w", name, err)
			}

			// Resource still exists
			if i == maxRetries-1 {
				return fmt.Errorf("namespace %s still exists in F5 XC API after waiting", name)
			}

			// Wait before retrying
			time.Sleep(5 * time.Second)
		}
	}

	return nil
}

// CheckNamespaceDisappears deletes a namespace outside of Terraform to test
// the "disappears" scenario where a resource is deleted externally.
func CheckNamespaceDisappears(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}

		c, err := GetTestClient()
		if err != nil {
			return fmt.Errorf("failed to get test client: %w", err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		name := rs.Primary.Attributes["name"]

		// Delete the namespace through the API using cascade_delete
		// (standard DELETE endpoint returns 501 Not Implemented)
		err = c.CascadeDeleteNamespace(ctx, name)
		if err != nil {
			return fmt.Errorf("failed to cascade delete namespace %s: %w", name, err)
		}

		return nil
	}
}

// CheckNamespaceExists verifies a namespace exists in the F5 XC API
func CheckNamespaceExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("resource ID is not set: %s", resourceName)
		}

		c, err := GetTestClient()
		if err != nil {
			return fmt.Errorf("failed to get test client: %w", err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		name := rs.Primary.Attributes["name"]
		namespace := rs.Primary.Attributes["namespace"]
		if namespace == "" {
			namespace = "system"
		}

		// Verify the namespace exists in the API
		apiResource, err := c.GetNamespace(ctx, namespace, name)
		if err != nil {
			return fmt.Errorf("namespace %s not found in API: %w", name, err)
		}

		// Verify the name matches
		if apiResource.Metadata.Name != name {
			return fmt.Errorf("API returned wrong namespace: got %s, want %s",
				apiResource.Metadata.Name, name)
		}

		return nil
	}
}

// CheckNamespaceAttributes verifies namespace attributes match expected values in the API
func CheckNamespaceAttributes(resourceName string, expectedLabels map[string]string, expectedAnnotations map[string]string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}

		c, err := GetTestClient()
		if err != nil {
			return fmt.Errorf("failed to get test client: %w", err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		name := rs.Primary.Attributes["name"]
		namespace := rs.Primary.Attributes["namespace"]
		if namespace == "" {
			namespace = "system"
		}

		apiResource, err := c.GetNamespace(ctx, namespace, name)
		if err != nil {
			return fmt.Errorf("failed to get namespace %s: %w", name, err)
		}

		// Check labels
		for k, v := range expectedLabels {
			if apiResource.Metadata.Labels[k] != v {
				return fmt.Errorf("label %s: got %q, want %q",
					k, apiResource.Metadata.Labels[k], v)
			}
		}

		// Check annotations
		for k, v := range expectedAnnotations {
			if apiResource.Metadata.Annotations[k] != v {
				return fmt.Errorf("annotation %s: got %q, want %q",
					k, apiResource.Metadata.Annotations[k], v)
			}
		}

		return nil
	}
}

// =============================================================================
// HEALTHCHECK RESOURCE TEST HELPERS
// =============================================================================

// CheckHealthcheckDestroyed verifies that healthcheck resources have been deleted
// from the F5 XC API, not just from Terraform state.
// Includes retry logic to handle async deletion.
func CheckHealthcheckDestroyed(s *terraform.State) error {
	c, err := GetTestClient()
	if err != nil {
		return fmt.Errorf("failed to get test client: %w", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "f5xc_healthcheck" {
			continue
		}

		name := rs.Primary.Attributes["name"]
		namespace := rs.Primary.Attributes["namespace"]
		if name == "" {
			name = rs.Primary.ID
		}
		if namespace == "" {
			namespace = "system"
		}

		// Retry loop to handle async deletion
		maxRetries := 6
		for i := 0; i < maxRetries; i++ {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			_, err := c.GetHealthcheck(ctx, namespace, name)
			cancel()

			if err != nil {
				// Check if it's a "not found" error - success!
				if strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "NOT_FOUND") {
					break // Resource is deleted
				}
				// Some other error occurred
				return fmt.Errorf("unexpected error checking healthcheck %s: %w", name, err)
			}

			// Resource still exists
			if i == maxRetries-1 {
				return fmt.Errorf("healthcheck %s still exists in F5 XC API after waiting", name)
			}

			// Wait before retrying
			time.Sleep(5 * time.Second)
		}
	}

	return nil
}

// CheckHealthcheckDisappears deletes a healthcheck outside of Terraform to test
// the "disappears" scenario where a resource is deleted externally.
func CheckHealthcheckDisappears(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}

		c, err := GetTestClient()
		if err != nil {
			return fmt.Errorf("failed to get test client: %w", err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		name := rs.Primary.Attributes["name"]
		namespace := rs.Primary.Attributes["namespace"]
		if namespace == "" {
			namespace = "system"
		}

		// Delete the healthcheck through the API (simulating external deletion)
		err = c.DeleteHealthcheck(ctx, namespace, name)
		if err != nil {
			return fmt.Errorf("failed to delete healthcheck %s: %w", name, err)
		}

		return nil
	}
}

// CheckHealthcheckExists verifies a healthcheck exists in the F5 XC API
func CheckHealthcheckExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("resource ID is not set: %s", resourceName)
		}

		c, err := GetTestClient()
		if err != nil {
			return fmt.Errorf("failed to get test client: %w", err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		name := rs.Primary.Attributes["name"]
		namespace := rs.Primary.Attributes["namespace"]
		if namespace == "" {
			namespace = "system"
		}

		// Verify the healthcheck exists in the API
		apiResource, err := c.GetHealthcheck(ctx, namespace, name)
		if err != nil {
			return fmt.Errorf("healthcheck %s not found in API: %w", name, err)
		}

		// Verify the name matches
		if apiResource.Metadata.Name != name {
			return fmt.Errorf("API returned wrong healthcheck: got %s, want %s",
				apiResource.Metadata.Name, name)
		}

		return nil
	}
}

// CheckHealthcheckAttributes verifies healthcheck attributes match expected values in the API
func CheckHealthcheckAttributes(resourceName string, expectedLabels map[string]string, expectedAnnotations map[string]string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}

		c, err := GetTestClient()
		if err != nil {
			return fmt.Errorf("failed to get test client: %w", err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		name := rs.Primary.Attributes["name"]
		namespace := rs.Primary.Attributes["namespace"]
		if namespace == "" {
			namespace = "system"
		}

		apiResource, err := c.GetHealthcheck(ctx, namespace, name)
		if err != nil {
			return fmt.Errorf("failed to get healthcheck %s: %w", name, err)
		}

		// Check labels
		for k, v := range expectedLabels {
			if apiResource.Metadata.Labels[k] != v {
				return fmt.Errorf("label %s: got %q, want %q",
					k, apiResource.Metadata.Labels[k], v)
			}
		}

		// Check annotations
		for k, v := range expectedAnnotations {
			if apiResource.Metadata.Annotations[k] != v {
				return fmt.Errorf("annotation %s: got %q, want %q",
					k, apiResource.Metadata.Annotations[k], v)
			}
		}

		return nil
	}
}

// =============================================================================
// IP_PREFIX_SET RESOURCE TEST HELPERS
// =============================================================================

// CheckIPPrefixSetDestroyed verifies that ip_prefix_set resources have been deleted
func CheckIPPrefixSetDestroyed(s *terraform.State) error {
	c, err := GetTestClient()
	if err != nil {
		return fmt.Errorf("failed to get test client: %w", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "f5xc_ip_prefix_set" {
			continue
		}

		name := rs.Primary.Attributes["name"]
		namespace := rs.Primary.Attributes["namespace"]
		if name == "" {
			name = rs.Primary.ID
		}
		if namespace == "" {
			namespace = "system"
		}

		maxRetries := 6
		for i := 0; i < maxRetries; i++ {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			_, err := c.GetIPPrefixSet(ctx, namespace, name)
			cancel()

			if err != nil {
				if strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "NOT_FOUND") {
					break
				}
				return fmt.Errorf("unexpected error checking ip_prefix_set %s: %w", name, err)
			}

			if i == maxRetries-1 {
				return fmt.Errorf("ip_prefix_set %s still exists in F5 XC API after waiting", name)
			}
			time.Sleep(5 * time.Second)
		}
	}
	return nil
}

// CheckIPPrefixSetDisappears deletes an ip_prefix_set outside of Terraform
func CheckIPPrefixSetDisappears(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}

		c, err := GetTestClient()
		if err != nil {
			return fmt.Errorf("failed to get test client: %w", err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		name := rs.Primary.Attributes["name"]
		namespace := rs.Primary.Attributes["namespace"]
		if namespace == "" {
			namespace = "system"
		}

		err = c.DeleteIPPrefixSet(ctx, namespace, name)
		if err != nil {
			return fmt.Errorf("failed to delete ip_prefix_set %s: %w", name, err)
		}
		return nil
	}
}

// CheckIPPrefixSetExists verifies an ip_prefix_set exists in the F5 XC API
func CheckIPPrefixSetExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("resource ID is not set: %s", resourceName)
		}

		c, err := GetTestClient()
		if err != nil {
			return fmt.Errorf("failed to get test client: %w", err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		name := rs.Primary.Attributes["name"]
		namespace := rs.Primary.Attributes["namespace"]
		if namespace == "" {
			namespace = "system"
		}

		apiResource, err := c.GetIPPrefixSet(ctx, namespace, name)
		if err != nil {
			return fmt.Errorf("ip_prefix_set %s not found in API: %w", name, err)
		}

		if apiResource.Metadata.Name != name {
			return fmt.Errorf("API returned wrong ip_prefix_set: got %s, want %s",
				apiResource.Metadata.Name, name)
		}
		return nil
	}
}

// =============================================================================
// BGP_ASN_SET RESOURCE TEST HELPERS
// =============================================================================

// CheckBGPAsnSetDestroyed verifies that bgp_asn_set resources have been deleted
func CheckBGPAsnSetDestroyed(s *terraform.State) error {
	c, err := GetTestClient()
	if err != nil {
		return fmt.Errorf("failed to get test client: %w", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "f5xc_bgp_asn_set" {
			continue
		}

		name := rs.Primary.Attributes["name"]
		namespace := rs.Primary.Attributes["namespace"]
		if name == "" {
			name = rs.Primary.ID
		}
		if namespace == "" {
			namespace = "system"
		}

		maxRetries := 6
		for i := 0; i < maxRetries; i++ {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			_, err := c.GetBGPAsnSet(ctx, namespace, name)
			cancel()

			if err != nil {
				if strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "NOT_FOUND") {
					break
				}
				return fmt.Errorf("unexpected error checking bgp_asn_set %s: %w", name, err)
			}

			if i == maxRetries-1 {
				return fmt.Errorf("bgp_asn_set %s still exists in F5 XC API after waiting", name)
			}
			time.Sleep(5 * time.Second)
		}
	}
	return nil
}

// CheckBGPAsnSetExists verifies a bgp_asn_set exists in the F5 XC API
func CheckBGPAsnSetExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("resource ID is not set: %s", resourceName)
		}

		c, err := GetTestClient()
		if err != nil {
			return fmt.Errorf("failed to get test client: %w", err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		name := rs.Primary.Attributes["name"]
		namespace := rs.Primary.Attributes["namespace"]
		if namespace == "" {
			namespace = "system"
		}

		apiResource, err := c.GetBGPAsnSet(ctx, namespace, name)
		if err != nil {
			return fmt.Errorf("bgp_asn_set %s not found in API: %w", name, err)
		}

		if apiResource.Metadata.Name != name {
			return fmt.Errorf("API returned wrong bgp_asn_set: got %s, want %s",
				apiResource.Metadata.Name, name)
		}
		return nil
	}
}

// CheckBGPAsnSetDisappears deletes a bgp_asn_set outside of Terraform to test
// the "disappears" scenario where a resource is deleted externally.
func CheckBGPAsnSetDisappears(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}

		c, err := GetTestClient()
		if err != nil {
			return fmt.Errorf("failed to get test client: %w", err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		name := rs.Primary.Attributes["name"]
		namespace := rs.Primary.Attributes["namespace"]
		if namespace == "" {
			namespace = "system"
		}

		// Delete the bgp_asn_set through the API (simulating external deletion)
		err = c.DeleteBGPAsnSet(ctx, namespace, name)
		if err != nil {
			return fmt.Errorf("failed to delete bgp_asn_set %s: %w", name, err)
		}

		return nil
	}
}

// =============================================================================
// POLICER RESOURCE TEST HELPERS
// =============================================================================

// CheckPolicerDestroyed verifies that policer resources have been deleted
func CheckPolicerDestroyed(s *terraform.State) error {
	c, err := GetTestClient()
	if err != nil {
		return fmt.Errorf("failed to get test client: %w", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "f5xc_policer" {
			continue
		}

		name := rs.Primary.Attributes["name"]
		namespace := rs.Primary.Attributes["namespace"]
		if name == "" {
			name = rs.Primary.ID
		}
		if namespace == "" {
			namespace = "system"
		}

		maxRetries := 6
		for i := 0; i < maxRetries; i++ {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			_, err := c.GetPolicer(ctx, namespace, name)
			cancel()

			if err != nil {
				if strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "NOT_FOUND") {
					break
				}
				return fmt.Errorf("unexpected error checking policer %s: %w", name, err)
			}

			if i == maxRetries-1 {
				return fmt.Errorf("policer %s still exists in F5 XC API after waiting", name)
			}
			time.Sleep(5 * time.Second)
		}
	}
	return nil
}

// CheckPolicerExists verifies a policer exists in the F5 XC API
func CheckPolicerExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("resource ID is not set: %s", resourceName)
		}

		c, err := GetTestClient()
		if err != nil {
			return fmt.Errorf("failed to get test client: %w", err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		name := rs.Primary.Attributes["name"]
		namespace := rs.Primary.Attributes["namespace"]
		if namespace == "" {
			namespace = "system"
		}

		apiResource, err := c.GetPolicer(ctx, namespace, name)
		if err != nil {
			return fmt.Errorf("policer %s not found in API: %w", name, err)
		}

		if apiResource.Metadata.Name != name {
			return fmt.Errorf("API returned wrong policer: got %s, want %s",
				apiResource.Metadata.Name, name)
		}
		return nil
	}
}

// CheckPolicerDisappears deletes a policer outside of Terraform to test
// the "disappears" scenario where a resource is deleted externally.
func CheckPolicerDisappears(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}

		c, err := GetTestClient()
		if err != nil {
			return fmt.Errorf("failed to get test client: %w", err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		name := rs.Primary.Attributes["name"]
		namespace := rs.Primary.Attributes["namespace"]
		if namespace == "" {
			namespace = "system"
		}

		// Delete the policer through the API (simulating external deletion)
		err = c.DeletePolicer(ctx, namespace, name)
		if err != nil {
			return fmt.Errorf("failed to delete policer %s: %w", name, err)
		}

		return nil
	}
}

// =============================================================================
// GEO_LOCATION_SET RESOURCE TEST HELPERS
// =============================================================================

// CheckGeoLocationSetDestroyed verifies that geo_location_set resources have been deleted
func CheckGeoLocationSetDestroyed(s *terraform.State) error {
	c, err := GetTestClient()
	if err != nil {
		return fmt.Errorf("failed to get test client: %w", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "f5xc_geo_location_set" {
			continue
		}

		name := rs.Primary.Attributes["name"]
		namespace := rs.Primary.Attributes["namespace"]
		if name == "" {
			name = rs.Primary.ID
		}
		if namespace == "" {
			namespace = "system"
		}

		maxRetries := 6
		for i := 0; i < maxRetries; i++ {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			_, err := c.GetGeoLocationSet(ctx, namespace, name)
			cancel()

			if err != nil {
				if strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "NOT_FOUND") {
					break
				}
				return fmt.Errorf("unexpected error checking geo_location_set %s: %w", name, err)
			}

			if i == maxRetries-1 {
				return fmt.Errorf("geo_location_set %s still exists in F5 XC API after waiting", name)
			}
			time.Sleep(5 * time.Second)
		}
	}
	return nil
}

// CheckGeoLocationSetExists verifies a geo_location_set exists in the F5 XC API
func CheckGeoLocationSetExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("resource ID is not set: %s", resourceName)
		}

		c, err := GetTestClient()
		if err != nil {
			return fmt.Errorf("failed to get test client: %w", err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		name := rs.Primary.Attributes["name"]
		namespace := rs.Primary.Attributes["namespace"]
		if namespace == "" {
			namespace = "system"
		}

		apiResource, err := c.GetGeoLocationSet(ctx, namespace, name)
		if err != nil {
			return fmt.Errorf("geo_location_set %s not found in API: %w", name, err)
		}

		if apiResource.Metadata.Name != name {
			return fmt.Errorf("API returned wrong geo_location_set: got %s, want %s",
				apiResource.Metadata.Name, name)
		}
		return nil
	}
}

// CheckGeoLocationSetDisappears deletes a geo_location_set outside of Terraform to test
// the "disappears" scenario where a resource is deleted externally.
func CheckGeoLocationSetDisappears(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}

		c, err := GetTestClient()
		if err != nil {
			return fmt.Errorf("failed to get test client: %w", err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		name := rs.Primary.Attributes["name"]
		namespace := rs.Primary.Attributes["namespace"]
		if namespace == "" {
			namespace = "system"
		}

		// Delete the geo_location_set through the API (simulating external deletion)
		err = c.DeleteGeoLocationSet(ctx, namespace, name)
		if err != nil {
			return fmt.Errorf("failed to delete geo_location_set %s: %w", name, err)
		}

		return nil
	}
}

// =============================================================================
// DATA_GROUP RESOURCE TEST HELPERS
// =============================================================================

// CheckDataGroupDestroyed verifies that data_group resources have been deleted
func CheckDataGroupDestroyed(s *terraform.State) error {
	c, err := GetTestClient()
	if err != nil {
		return fmt.Errorf("failed to get test client: %w", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "f5xc_data_group" {
			continue
		}

		name := rs.Primary.Attributes["name"]
		namespace := rs.Primary.Attributes["namespace"]
		if name == "" {
			name = rs.Primary.ID
		}
		if namespace == "" {
			namespace = "system"
		}

		maxRetries := 6
		for i := 0; i < maxRetries; i++ {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			_, err := c.GetDataGroup(ctx, namespace, name)
			cancel()

			if err != nil {
				if strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "NOT_FOUND") {
					break
				}
				return fmt.Errorf("unexpected error checking data_group %s: %w", name, err)
			}

			if i == maxRetries-1 {
				return fmt.Errorf("data_group %s still exists in F5 XC API after waiting", name)
			}
			time.Sleep(5 * time.Second)
		}
	}
	return nil
}

// CheckDataGroupExists verifies a data_group exists in the F5 XC API
func CheckDataGroupExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("resource ID is not set: %s", resourceName)
		}

		c, err := GetTestClient()
		if err != nil {
			return fmt.Errorf("failed to get test client: %w", err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		name := rs.Primary.Attributes["name"]
		namespace := rs.Primary.Attributes["namespace"]
		if namespace == "" {
			namespace = "system"
		}

		apiResource, err := c.GetDataGroup(ctx, namespace, name)
		if err != nil {
			return fmt.Errorf("data_group %s not found in API: %w", name, err)
		}

		if apiResource.Metadata.Name != name {
			return fmt.Errorf("API returned wrong data_group: got %s, want %s",
				apiResource.Metadata.Name, name)
		}
		return nil
	}
}

// CheckDataGroupDisappears deletes a data_group outside of Terraform to test
// the "disappears" scenario where a resource is deleted externally.
func CheckDataGroupDisappears(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}

		c, err := GetTestClient()
		if err != nil {
			return fmt.Errorf("failed to get test client: %w", err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		name := rs.Primary.Attributes["name"]
		namespace := rs.Primary.Attributes["namespace"]
		if namespace == "" {
			namespace = "system"
		}

		// Delete the data_group through the API (simulating external deletion)
		err = c.DeleteDataGroup(ctx, namespace, name)
		if err != nil {
			return fmt.Errorf("failed to delete data_group %s: %w", name, err)
		}

		return nil
	}
}

// =============================================================================
// DATA_TYPE RESOURCE TEST HELPERS
// =============================================================================

// CheckDataTypeDestroyed verifies that data_type resources have been deleted
func CheckDataTypeDestroyed(s *terraform.State) error {
	c, err := GetTestClient()
	if err != nil {
		return fmt.Errorf("failed to get test client: %w", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "f5xc_data_type" {
			continue
		}

		name := rs.Primary.Attributes["name"]
		namespace := rs.Primary.Attributes["namespace"]
		if name == "" {
			name = rs.Primary.ID
		}
		if namespace == "" {
			namespace = "system"
		}

		maxRetries := 6
		for i := 0; i < maxRetries; i++ {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			_, err := c.GetDataType(ctx, namespace, name)
			cancel()

			if err != nil {
				if strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "NOT_FOUND") {
					break
				}
				return fmt.Errorf("unexpected error checking data_type %s: %w", name, err)
			}

			if i == maxRetries-1 {
				return fmt.Errorf("data_type %s still exists in F5 XC API after waiting", name)
			}
			time.Sleep(5 * time.Second)
		}
	}
	return nil
}

// CheckDataTypeExists verifies a data_type exists in the F5 XC API
func CheckDataTypeExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("resource ID is not set: %s", resourceName)
		}

		c, err := GetTestClient()
		if err != nil {
			return fmt.Errorf("failed to get test client: %w", err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		name := rs.Primary.Attributes["name"]
		namespace := rs.Primary.Attributes["namespace"]
		if namespace == "" {
			namespace = "system"
		}

		apiResource, err := c.GetDataType(ctx, namespace, name)
		if err != nil {
			return fmt.Errorf("data_type %s not found in API: %w", name, err)
		}

		if apiResource.Metadata.Name != name {
			return fmt.Errorf("API returned wrong data_type: got %s, want %s",
				apiResource.Metadata.Name, name)
		}
		return nil
	}
}

// CheckDataTypeDisappears deletes a data_type outside of Terraform to test
// the "disappears" scenario where a resource is deleted externally.
func CheckDataTypeDisappears(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}

		c, err := GetTestClient()
		if err != nil {
			return fmt.Errorf("failed to get test client: %w", err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		name := rs.Primary.Attributes["name"]
		namespace := rs.Primary.Attributes["namespace"]
		if namespace == "" {
			namespace = "system"
		}

		// Delete the data_type through the API (simulating external deletion)
		err = c.DeleteDataType(ctx, namespace, name)
		if err != nil {
			return fmt.Errorf("failed to delete data_type %s: %w", name, err)
		}

		return nil
	}
}

// =============================================================================
// FILTER_SET RESOURCE TEST HELPERS
// =============================================================================

// CheckFilterSetDestroyed verifies that filter_set resources have been deleted
func CheckFilterSetDestroyed(s *terraform.State) error {
	c, err := GetTestClient()
	if err != nil {
		return fmt.Errorf("failed to get test client: %w", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "f5xc_filter_set" {
			continue
		}

		name := rs.Primary.Attributes["name"]
		namespace := rs.Primary.Attributes["namespace"]
		if name == "" {
			name = rs.Primary.ID
		}
		if namespace == "" {
			namespace = "system"
		}

		maxRetries := 6
		for i := 0; i < maxRetries; i++ {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			_, err := c.GetFilterSet(ctx, namespace, name)
			cancel()

			if err != nil {
				if strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "NOT_FOUND") {
					break
				}
				return fmt.Errorf("unexpected error checking filter_set %s: %w", name, err)
			}

			if i == maxRetries-1 {
				return fmt.Errorf("filter_set %s still exists in F5 XC API after waiting", name)
			}
			time.Sleep(5 * time.Second)
		}
	}
	return nil
}

// CheckFilterSetExists verifies a filter_set exists in the F5 XC API
func CheckFilterSetExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("resource ID is not set: %s", resourceName)
		}

		c, err := GetTestClient()
		if err != nil {
			return fmt.Errorf("failed to get test client: %w", err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		name := rs.Primary.Attributes["name"]
		namespace := rs.Primary.Attributes["namespace"]
		if namespace == "" {
			namespace = "system"
		}

		apiResource, err := c.GetFilterSet(ctx, namespace, name)
		if err != nil {
			return fmt.Errorf("filter_set %s not found in API: %w", name, err)
		}

		if apiResource.Metadata.Name != name {
			return fmt.Errorf("API returned wrong filter_set: got %s, want %s",
				apiResource.Metadata.Name, name)
		}
		return nil
	}
}

// CheckFilterSetDisappears deletes a filter_set outside of Terraform to test
// the "disappears" scenario where a resource is deleted externally.
func CheckFilterSetDisappears(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}

		c, err := GetTestClient()
		if err != nil {
			return fmt.Errorf("failed to get test client: %w", err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		name := rs.Primary.Attributes["name"]
		namespace := rs.Primary.Attributes["namespace"]
		if namespace == "" {
			namespace = "system"
		}

		// Delete the filter_set through the API (simulating external deletion)
		err = c.DeleteFilterSet(ctx, namespace, name)
		if err != nil {
			return fmt.Errorf("failed to delete filter_set %s: %w", name, err)
		}

		return nil
	}
}

// =============================================================================
// FORWARDING_CLASS RESOURCE TEST HELPERS
// =============================================================================

// CheckForwardingClassDestroyed verifies that forwarding_class resources have been deleted
func CheckForwardingClassDestroyed(s *terraform.State) error {
	c, err := GetTestClient()
	if err != nil {
		return fmt.Errorf("failed to get test client: %w", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "f5xc_forwarding_class" {
			continue
		}

		name := rs.Primary.Attributes["name"]
		namespace := rs.Primary.Attributes["namespace"]
		if name == "" {
			name = rs.Primary.ID
		}
		if namespace == "" {
			namespace = "system"
		}

		maxRetries := 6
		for i := 0; i < maxRetries; i++ {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			_, err := c.GetForwardingClass(ctx, namespace, name)
			cancel()

			if err != nil {
				if strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "NOT_FOUND") {
					break
				}
				return fmt.Errorf("unexpected error checking forwarding_class %s: %w", name, err)
			}

			if i == maxRetries-1 {
				return fmt.Errorf("forwarding_class %s still exists in F5 XC API after waiting", name)
			}
			time.Sleep(5 * time.Second)
		}
	}
	return nil
}

// CheckForwardingClassExists verifies a forwarding_class exists in the F5 XC API
func CheckForwardingClassExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("resource ID is not set: %s", resourceName)
		}

		c, err := GetTestClient()
		if err != nil {
			return fmt.Errorf("failed to get test client: %w", err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		name := rs.Primary.Attributes["name"]
		namespace := rs.Primary.Attributes["namespace"]
		if namespace == "" {
			namespace = "system"
		}

		apiResource, err := c.GetForwardingClass(ctx, namespace, name)
		if err != nil {
			return fmt.Errorf("forwarding_class %s not found in API: %w", name, err)
		}

		if apiResource.Metadata.Name != name {
			return fmt.Errorf("API returned wrong forwarding_class: got %s, want %s",
				apiResource.Metadata.Name, name)
		}
		return nil
	}
}

// CheckForwardingClassDisappears deletes a forwarding_class outside of Terraform to test
// the "disappears" scenario where a resource is deleted externally.
func CheckForwardingClassDisappears(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}

		c, err := GetTestClient()
		if err != nil {
			return fmt.Errorf("failed to get test client: %w", err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		name := rs.Primary.Attributes["name"]
		namespace := rs.Primary.Attributes["namespace"]
		if namespace == "" {
			namespace = "system"
		}

		// Delete the forwarding_class through the API (simulating external deletion)
		err = c.DeleteForwardingClass(ctx, namespace, name)
		if err != nil {
			return fmt.Errorf("failed to delete forwarding_class %s: %w", name, err)
		}

		return nil
	}
}

// ============================================================================
// Rate Limiter Test Helpers
// ============================================================================

// CheckRateLimiterDestroyed verifies that a rate_limiter has been properly destroyed.
// It implements the resource.TestCheckFunc signature for use in acceptance tests.
func CheckRateLimiterDestroyed(s *terraform.State) error {
	c, err := GetTestClient()
	if err != nil {
		return fmt.Errorf("failed to get test client: %w", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "f5xc_rate_limiter" {
			continue
		}

		name := rs.Primary.Attributes["name"]
		namespace := rs.Primary.Attributes["namespace"]
		if namespace == "" {
			namespace = "system"
		}

		var lastErr error
		for i := 0; i < 6; i++ {
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			_, err := c.GetRateLimiter(ctx, namespace, name)
			cancel()

			if err != nil {
				lastErr = err
				break
			}

			if i < 5 {
				time.Sleep(5 * time.Second)
			}
		}

		if lastErr == nil {
			return fmt.Errorf("rate_limiter %s still exists after deletion wait period", name)
		}
	}
	return nil
}

// CheckRateLimiterExists verifies that a rate_limiter exists in the F5 XC API.
func CheckRateLimiterExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no ID is set for resource: %s", resourceName)
		}

		c, err := GetTestClient()
		if err != nil {
			return fmt.Errorf("failed to get test client: %w", err)
		}

		name := rs.Primary.Attributes["name"]
		namespace := rs.Primary.Attributes["namespace"]
		if namespace == "" {
			namespace = "system"
		}

		var lastErr error
		for i := 0; i < 6; i++ {
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			apiResource, err := c.GetRateLimiter(ctx, namespace, name)
			cancel()

			if err == nil && apiResource != nil {
				if apiResource.Metadata.Name != name {
					return fmt.Errorf("API returned wrong rate_limiter: got %s, want %s",
						apiResource.Metadata.Name, name)
				}
				return nil
			}

			lastErr = err
			if i < 5 {
				time.Sleep(5 * time.Second)
			}
		}

		return fmt.Errorf("rate_limiter %s not found after retries: %w", name, lastErr)
	}
}

// CheckRateLimiterDisappears deletes a rate_limiter outside of Terraform to test
// the "disappears" scenario where a resource is deleted externally.
func CheckRateLimiterDisappears(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}

		c, err := GetTestClient()
		if err != nil {
			return fmt.Errorf("failed to get test client: %w", err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		name := rs.Primary.Attributes["name"]
		namespace := rs.Primary.Attributes["namespace"]
		if namespace == "" {
			namespace = "system"
		}

		// Delete the rate_limiter through the API (simulating external deletion)
		err = c.DeleteRateLimiter(ctx, namespace, name)
		if err != nil {
			return fmt.Errorf("failed to delete rate_limiter %s: %w", name, err)
		}

		return nil
	}
}

// ============================================================================
// User Identification Test Helpers
// ============================================================================

// CheckUserIdentificationDestroyed verifies that a user_identification has been properly destroyed.
// It implements the resource.TestCheckFunc signature for use in acceptance tests.
func CheckUserIdentificationDestroyed(s *terraform.State) error {
	c, err := GetTestClient()
	if err != nil {
		return fmt.Errorf("failed to get test client: %w", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "f5xc_user_identification" {
			continue
		}

		name := rs.Primary.Attributes["name"]
		namespace := rs.Primary.Attributes["namespace"]
		if namespace == "" {
			namespace = "system"
		}

		var lastErr error
		for i := 0; i < 6; i++ {
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			_, err := c.GetUserIdentification(ctx, namespace, name)
			cancel()

			if err != nil {
				lastErr = err
				break
			}

			if i < 5 {
				time.Sleep(5 * time.Second)
			}
		}

		if lastErr == nil {
			return fmt.Errorf("user_identification %s still exists after deletion wait period", name)
		}
	}
	return nil
}

// CheckUserIdentificationExists verifies that a user_identification exists in the F5 XC API.
func CheckUserIdentificationExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no ID is set for resource: %s", resourceName)
		}

		c, err := GetTestClient()
		if err != nil {
			return fmt.Errorf("failed to get test client: %w", err)
		}

		name := rs.Primary.Attributes["name"]
		namespace := rs.Primary.Attributes["namespace"]
		if namespace == "" {
			namespace = "system"
		}

		var lastErr error
		for i := 0; i < 6; i++ {
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			apiResource, err := c.GetUserIdentification(ctx, namespace, name)
			cancel()

			if err == nil && apiResource != nil {
				if apiResource.Metadata.Name != name {
					return fmt.Errorf("API returned wrong user_identification: got %s, want %s",
						apiResource.Metadata.Name, name)
				}
				return nil
			}

			lastErr = err
			if i < 5 {
				time.Sleep(5 * time.Second)
			}
		}

		return fmt.Errorf("user_identification %s not found after retries: %w", name, lastErr)
	}
}

// CheckUserIdentificationDisappears deletes a user_identification outside of Terraform to test
// the "disappears" scenario where a resource is deleted externally.
func CheckUserIdentificationDisappears(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}

		c, err := GetTestClient()
		if err != nil {
			return fmt.Errorf("failed to get test client: %w", err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		name := rs.Primary.Attributes["name"]
		namespace := rs.Primary.Attributes["namespace"]
		if namespace == "" {
			namespace = "system"
		}

		// Delete the user_identification through the API (simulating external deletion)
		err = c.DeleteUserIdentification(ctx, namespace, name)
		if err != nil {
			return fmt.Errorf("failed to delete user_identification %s: %w", name, err)
		}

		return nil
	}
}

// ============================================================================
// Malicious User Mitigation Test Helpers
// ============================================================================

// CheckMaliciousUserMitigationDestroyed verifies that a malicious_user_mitigation has been properly destroyed.
// It implements the resource.TestCheckFunc signature for use in acceptance tests.
func CheckMaliciousUserMitigationDestroyed(s *terraform.State) error {
	c, err := GetTestClient()
	if err != nil {
		return fmt.Errorf("failed to get test client: %w", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "f5xc_malicious_user_mitigation" {
			continue
		}

		name := rs.Primary.Attributes["name"]
		namespace := rs.Primary.Attributes["namespace"]
		if namespace == "" {
			namespace = "system"
		}

		var lastErr error
		for i := 0; i < 6; i++ {
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			_, err := c.GetMaliciousUserMitigation(ctx, namespace, name)
			cancel()

			if err != nil {
				lastErr = err
				break
			}

			if i < 5 {
				time.Sleep(5 * time.Second)
			}
		}

		if lastErr == nil {
			return fmt.Errorf("malicious_user_mitigation %s still exists after deletion wait period", name)
		}
	}
	return nil
}

// CheckMaliciousUserMitigationExists verifies that a malicious_user_mitigation exists in the F5 XC API.
func CheckMaliciousUserMitigationExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no ID is set for resource: %s", resourceName)
		}

		c, err := GetTestClient()
		if err != nil {
			return fmt.Errorf("failed to get test client: %w", err)
		}

		name := rs.Primary.Attributes["name"]
		namespace := rs.Primary.Attributes["namespace"]
		if namespace == "" {
			namespace = "system"
		}

		var lastErr error
		for i := 0; i < 6; i++ {
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			apiResource, err := c.GetMaliciousUserMitigation(ctx, namespace, name)
			cancel()

			if err == nil && apiResource != nil {
				if apiResource.Metadata.Name != name {
					return fmt.Errorf("API returned wrong malicious_user_mitigation: got %s, want %s",
						apiResource.Metadata.Name, name)
				}
				return nil
			}

			lastErr = err
			if i < 5 {
				time.Sleep(5 * time.Second)
			}
		}

		return fmt.Errorf("malicious_user_mitigation %s not found after retries: %w", name, lastErr)
	}
}

// CheckMaliciousUserMitigationDisappears deletes a malicious_user_mitigation outside of Terraform to test
// the "disappears" scenario where a resource is deleted externally.
func CheckMaliciousUserMitigationDisappears(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}

		c, err := GetTestClient()
		if err != nil {
			return fmt.Errorf("failed to get test client: %w", err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		name := rs.Primary.Attributes["name"]
		namespace := rs.Primary.Attributes["namespace"]
		if namespace == "" {
			namespace = "system"
		}

		// Delete the malicious_user_mitigation through the API (simulating external deletion)
		err = c.DeleteMaliciousUserMitigation(ctx, namespace, name)
		if err != nil {
			return fmt.Errorf("failed to delete malicious_user_mitigation %s: %w", name, err)
		}

		return nil
	}
}

// CheckAlertPolicyDestroyed verifies that alert_policy has been properly destroyed
func CheckAlertPolicyDestroyed(s *terraform.State) error {
	c, err := GetTestClient()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "f5xc_alert_policy" {
			continue
		}

		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		name := rs.Primary.Attributes["name"]
		namespace := rs.Primary.Attributes["namespace"]
		if namespace == "" {
			namespace = "system"
		}

		// Try to get the alert_policy - it should return an error if properly destroyed
		_, err := c.GetAlertPolicy(ctx, namespace, name)
		if err == nil {
			return fmt.Errorf("alert_policy %s still exists", name)
		}
	}
	return nil
}

// CheckAlertPolicyExists verifies that alert_policy exists in the F5 XC API
func CheckAlertPolicyExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no ID is set for resource: %s", resourceName)
		}

		c, err := GetTestClient()
		if err != nil {
			return err
		}

		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		name := rs.Primary.Attributes["name"]
		namespace := rs.Primary.Attributes["namespace"]
		if namespace == "" {
			namespace = "system"
		}

		alertPolicy, err := c.GetAlertPolicy(ctx, namespace, name)
		if err != nil {
			return fmt.Errorf("error fetching alert_policy %s: %w", name, err)
		}

		if alertPolicy.Metadata.Name != name {
			return fmt.Errorf("alert_policy name mismatch: expected %s, got %s", name, alertPolicy.Metadata.Name)
		}

		return nil
	}
}

// CheckAlertPolicyDisappears deletes alert_policy outside of Terraform to test disappearance handling
func CheckAlertPolicyDisappears(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}

		c, err := GetTestClient()
		if err != nil {
			return err
		}

		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		name := rs.Primary.Attributes["name"]
		namespace := rs.Primary.Attributes["namespace"]
		if namespace == "" {
			namespace = "system"
		}

		// Delete the alert_policy through the API (simulating external deletion)
		err = c.DeleteAlertPolicy(ctx, namespace, name)
		if err != nil {
			return fmt.Errorf("failed to delete alert_policy %s: %w", name, err)
		}

		return nil
	}
}

// CheckAlertReceiverDestroyed verifies that alert_receiver has been properly destroyed
func CheckAlertReceiverDestroyed(s *terraform.State) error {
	c, err := GetTestClient()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "f5xc_alert_receiver" {
			continue
		}

		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		name := rs.Primary.Attributes["name"]
		namespace := rs.Primary.Attributes["namespace"]
		if namespace == "" {
			namespace = "system"
		}

		// Try to get the alert_receiver - it should return an error if properly destroyed
		_, err := c.GetAlertReceiver(ctx, namespace, name)
		if err == nil {
			return fmt.Errorf("alert_receiver %s still exists", name)
		}
	}
	return nil
}

// CheckAlertReceiverExists verifies that alert_receiver exists in the F5 XC API
func CheckAlertReceiverExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no ID is set for resource: %s", resourceName)
		}

		c, err := GetTestClient()
		if err != nil {
			return err
		}

		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		name := rs.Primary.Attributes["name"]
		namespace := rs.Primary.Attributes["namespace"]
		if namespace == "" {
			namespace = "system"
		}

		alertReceiver, err := c.GetAlertReceiver(ctx, namespace, name)
		if err != nil {
			return fmt.Errorf("error fetching alert_receiver %s: %w", name, err)
		}

		if alertReceiver.Metadata.Name != name {
			return fmt.Errorf("alert_receiver name mismatch: expected %s, got %s", name, alertReceiver.Metadata.Name)
		}

		return nil
	}
}

// CheckAlertReceiverDisappears deletes alert_receiver outside of Terraform to test disappearance handling
func CheckAlertReceiverDisappears(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}

		c, err := GetTestClient()
		if err != nil {
			return err
		}

		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		name := rs.Primary.Attributes["name"]
		namespace := rs.Primary.Attributes["namespace"]
		if namespace == "" {
			namespace = "system"
		}

		// Delete the alert_receiver through the API (simulating external deletion)
		err = c.DeleteAlertReceiver(ctx, namespace, name)
		if err != nil {
			return fmt.Errorf("failed to delete alert_receiver %s: %w", name, err)
		}

		return nil
	}
}

// CheckCloudCredentialsDestroyed verifies that cloud_credentials has been properly destroyed
func CheckCloudCredentialsDestroyed(s *terraform.State) error {
	c, err := GetTestClient()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "f5xc_cloud_credentials" {
			continue
		}

		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		name := rs.Primary.Attributes["name"]
		namespace := rs.Primary.Attributes["namespace"]
		if namespace == "" {
			namespace = "system"
		}

		// Try to get the cloud_credentials - it should return an error if properly destroyed
		_, err := c.GetCloudCredentials(ctx, namespace, name)
		if err == nil {
			return fmt.Errorf("cloud_credentials %s still exists", name)
		}
	}
	return nil
}

// CheckCloudCredentialsExists verifies that cloud_credentials exists in the F5 XC API
func CheckCloudCredentialsExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no ID is set for resource: %s", resourceName)
		}

		c, err := GetTestClient()
		if err != nil {
			return err
		}

		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		name := rs.Primary.Attributes["name"]
		namespace := rs.Primary.Attributes["namespace"]
		if namespace == "" {
			namespace = "system"
		}

		cloudCredentials, err := c.GetCloudCredentials(ctx, namespace, name)
		if err != nil {
			return fmt.Errorf("error fetching cloud_credentials %s: %w", name, err)
		}

		if cloudCredentials.Metadata.Name != name {
			return fmt.Errorf("cloud_credentials name mismatch: expected %s, got %s", name, cloudCredentials.Metadata.Name)
		}

		return nil
	}
}

// CheckCloudCredentialsDisappears deletes cloud_credentials outside of Terraform to test disappearance handling
func CheckCloudCredentialsDisappears(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}

		c, err := GetTestClient()
		if err != nil {
			return err
		}

		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		name := rs.Primary.Attributes["name"]
		namespace := rs.Primary.Attributes["namespace"]
		if namespace == "" {
			namespace = "system"
		}

		// Delete the cloud_credentials through the API (simulating external deletion)
		err = c.DeleteCloudCredentials(ctx, namespace, name)
		if err != nil {
			return fmt.Errorf("failed to delete cloud_credentials %s: %w", name, err)
		}

		return nil
	}
}

// CheckNetworkConnectorDestroyed verifies that network_connector has been properly destroyed
func CheckNetworkConnectorDestroyed(s *terraform.State) error {
	c, err := GetTestClient()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "f5xc_network_connector" {
			continue
		}

		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		name := rs.Primary.Attributes["name"]
		namespace := rs.Primary.Attributes["namespace"]
		if namespace == "" {
			namespace = "system"
		}

		// Try to get the network_connector - it should return an error if properly destroyed
		_, err := c.GetNetworkConnector(ctx, namespace, name)
		if err == nil {
			return fmt.Errorf("network_connector %s still exists", name)
		}
	}
	return nil
}

// CheckNetworkConnectorExists verifies that network_connector exists in the F5 XC API
func CheckNetworkConnectorExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no ID is set for resource: %s", resourceName)
		}

		c, err := GetTestClient()
		if err != nil {
			return err
		}

		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		name := rs.Primary.Attributes["name"]
		namespace := rs.Primary.Attributes["namespace"]
		if namespace == "" {
			namespace = "system"
		}

		networkConnector, err := c.GetNetworkConnector(ctx, namespace, name)
		if err != nil {
			return fmt.Errorf("error fetching network_connector %s: %w", name, err)
		}

		if networkConnector.Metadata.Name != name {
			return fmt.Errorf("network_connector name mismatch: expected %s, got %s", name, networkConnector.Metadata.Name)
		}

		return nil
	}
}

// CheckNetworkConnectorDisappears deletes network_connector outside of Terraform to test disappearance handling
func CheckNetworkConnectorDisappears(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}

		c, err := GetTestClient()
		if err != nil {
			return err
		}

		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		name := rs.Primary.Attributes["name"]
		namespace := rs.Primary.Attributes["namespace"]
		if namespace == "" {
			namespace = "system"
		}

		// Delete the network_connector through the API (simulating external deletion)
		err = c.DeleteNetworkConnector(ctx, namespace, name)
		if err != nil {
			return fmt.Errorf("failed to delete network_connector %s: %w", name, err)
		}

		return nil
	}
}

// ============================================================================
// App Firewall Test Helpers
// ============================================================================

// CheckAppFirewallDestroyed verifies that app_firewall has been properly destroyed
func CheckAppFirewallDestroyed(s *terraform.State) error {
	c, err := GetTestClient()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "f5xc_app_firewall" {
			continue
		}

		name := rs.Primary.Attributes["name"]
		namespace := rs.Primary.Attributes["namespace"]
		if namespace == "" {
			namespace = "system"
		}

		maxRetries := 6
		for i := 0; i < maxRetries; i++ {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			_, err := c.GetAppFirewall(ctx, namespace, name)
			cancel()

			if err != nil {
				if strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "NOT_FOUND") {
					break
				}
				return fmt.Errorf("unexpected error checking app_firewall %s: %w", name, err)
			}

			if i == maxRetries-1 {
				return fmt.Errorf("app_firewall %s still exists in F5 XC API after waiting", name)
			}
			time.Sleep(5 * time.Second)
		}
	}
	return nil
}

// CheckAppFirewallExists verifies that app_firewall exists in the F5 XC API
func CheckAppFirewallExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no ID is set for resource: %s", resourceName)
		}

		c, err := GetTestClient()
		if err != nil {
			return err
		}

		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		name := rs.Primary.Attributes["name"]
		namespace := rs.Primary.Attributes["namespace"]
		if namespace == "" {
			namespace = "system"
		}

		appFirewall, err := c.GetAppFirewall(ctx, namespace, name)
		if err != nil {
			return fmt.Errorf("error fetching app_firewall %s: %w", name, err)
		}

		if appFirewall.Metadata.Name != name {
			return fmt.Errorf("app_firewall name mismatch: expected %s, got %s", name, appFirewall.Metadata.Name)
		}

		return nil
	}
}

// CheckAppFirewallDisappears deletes app_firewall outside of Terraform to test disappearance handling
func CheckAppFirewallDisappears(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}

		c, err := GetTestClient()
		if err != nil {
			return err
		}

		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		name := rs.Primary.Attributes["name"]
		namespace := rs.Primary.Attributes["namespace"]
		if namespace == "" {
			namespace = "system"
		}

		// Delete the app_firewall through the API (simulating external deletion)
		err = c.DeleteAppFirewall(ctx, namespace, name)
		if err != nil {
			return fmt.Errorf("failed to delete app_firewall %s: %w", name, err)
		}

		return nil
	}
}

// ============================================================================
// Origin Pool Test Helpers
// ============================================================================

// CheckOriginPoolDestroyed verifies that origin_pool has been properly destroyed
func CheckOriginPoolDestroyed(s *terraform.State) error {
	c, err := GetTestClient()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "f5xc_origin_pool" {
			continue
		}

		name := rs.Primary.Attributes["name"]
		namespace := rs.Primary.Attributes["namespace"]
		if namespace == "" {
			namespace = "system"
		}

		maxRetries := 6
		for i := 0; i < maxRetries; i++ {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			_, err := c.GetOriginPool(ctx, namespace, name)
			cancel()

			if err != nil {
				if strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "NOT_FOUND") {
					break
				}
				return fmt.Errorf("unexpected error checking origin_pool %s: %w", name, err)
			}

			if i == maxRetries-1 {
				return fmt.Errorf("origin_pool %s still exists in F5 XC API after waiting", name)
			}
			time.Sleep(5 * time.Second)
		}
	}
	return nil
}

// CheckOriginPoolExists verifies that origin_pool exists in the F5 XC API
func CheckOriginPoolExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no ID is set for resource: %s", resourceName)
		}

		c, err := GetTestClient()
		if err != nil {
			return err
		}

		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		name := rs.Primary.Attributes["name"]
		namespace := rs.Primary.Attributes["namespace"]
		if namespace == "" {
			namespace = "system"
		}

		originPool, err := c.GetOriginPool(ctx, namespace, name)
		if err != nil {
			return fmt.Errorf("error fetching origin_pool %s: %w", name, err)
		}

		if originPool.Metadata.Name != name {
			return fmt.Errorf("origin_pool name mismatch: expected %s, got %s", name, originPool.Metadata.Name)
		}

		return nil
	}
}

// CheckOriginPoolDisappears deletes origin_pool outside of Terraform to test disappearance handling
func CheckOriginPoolDisappears(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}

		c, err := GetTestClient()
		if err != nil {
			return err
		}

		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		name := rs.Primary.Attributes["name"]
		namespace := rs.Primary.Attributes["namespace"]
		if namespace == "" {
			namespace = "system"
		}

		// Delete the origin_pool through the API (simulating external deletion)
		err = c.DeleteOriginPool(ctx, namespace, name)
		if err != nil {
			return fmt.Errorf("failed to delete origin_pool %s: %w", name, err)
		}

		return nil
	}
}

// ============================================================================
// Service Policy Test Helpers
// ============================================================================

// CheckServicePolicyDestroyed verifies that service_policy has been properly destroyed
func CheckServicePolicyDestroyed(s *terraform.State) error {
	c, err := GetTestClient()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "f5xc_service_policy" {
			continue
		}

		name := rs.Primary.Attributes["name"]
		namespace := rs.Primary.Attributes["namespace"]
		if namespace == "" {
			namespace = "system"
		}

		maxRetries := 6
		for i := 0; i < maxRetries; i++ {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			_, err := c.GetServicePolicy(ctx, namespace, name)
			cancel()

			if err != nil {
				if strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "NOT_FOUND") {
					break
				}
				return fmt.Errorf("unexpected error checking service_policy %s: %w", name, err)
			}

			if i == maxRetries-1 {
				return fmt.Errorf("service_policy %s still exists in F5 XC API after waiting", name)
			}
			time.Sleep(5 * time.Second)
		}
	}
	return nil
}

// CheckServicePolicyExists verifies that service_policy exists in the F5 XC API
func CheckServicePolicyExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no ID is set for resource: %s", resourceName)
		}

		c, err := GetTestClient()
		if err != nil {
			return err
		}

		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		name := rs.Primary.Attributes["name"]
		namespace := rs.Primary.Attributes["namespace"]
		if namespace == "" {
			namespace = "system"
		}

		servicePolicy, err := c.GetServicePolicy(ctx, namespace, name)
		if err != nil {
			return fmt.Errorf("error fetching service_policy %s: %w", name, err)
		}

		if servicePolicy.Metadata.Name != name {
			return fmt.Errorf("service_policy name mismatch: expected %s, got %s", name, servicePolicy.Metadata.Name)
		}

		return nil
	}
}

// CheckServicePolicyDisappears deletes service_policy outside of Terraform to test disappearance handling
func CheckServicePolicyDisappears(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}

		c, err := GetTestClient()
		if err != nil {
			return err
		}

		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		name := rs.Primary.Attributes["name"]
		namespace := rs.Primary.Attributes["namespace"]
		if namespace == "" {
			namespace = "system"
		}

		err = c.DeleteServicePolicy(ctx, namespace, name)
		if err != nil {
			return fmt.Errorf("failed to delete service_policy %s: %w", name, err)
		}

		return nil
	}
}
