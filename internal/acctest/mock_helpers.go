// Copyright (c) F5XC Community
// SPDX-License-Identifier: MPL-2.0

// Package acctest provides acceptance test utilities including mock server support.
package acctest

import (
	"fmt"
	"os"
	"regexp"
	"sync"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/f5xc/terraform-provider-f5xc/internal/mocks"
	"github.com/f5xc/terraform-provider-f5xc/internal/provider"
)

var (
	// globalMockServer is a shared mock server for all tests when F5XC_MOCK_MODE is set
	globalMockServer     *mocks.Server
	globalMockServerOnce sync.Once
	globalMockServerMu   sync.Mutex
	// originalEnvVars stores original environment variables before mock mode override
	originalEnvVars = make(map[string]string)
)

// MustCompileRegexp compiles a regular expression and panics if it fails.
// This is useful for test helpers where the regex is known at compile time.
func MustCompileRegexp(pattern string) *regexp.Regexp {
	return regexp.MustCompile(pattern)
}

// MockTestConfig holds configuration for mock-based tests
type MockTestConfig struct {
	Server        *mocks.Server
	ProviderFunc  func() (tfprotov6.ProviderServer, error)
	OriginalURL   string
	OriginalToken string
}

// SetupMockTest creates a mock server and configures the provider to use it.
// It returns a cleanup function that should be deferred.
//
// Usage:
//
//	func TestAccResource_withMock(t *testing.T) {
//	    mockCfg := acctest.SetupMockTest(t)
//	    defer mockCfg.Cleanup()
//
//	    // Pre-populate mock resources if needed
//	    mockCfg.Server.SetResource("/api/config/namespaces/system/aws_vpc_sites/my-site", ...)
//
//	    resource.Test(t, resource.TestCase{
//	        ProtoV6ProviderFactories: mockCfg.ProtoV6ProviderFactories(),
//	        Steps: []resource.TestStep{...},
//	    })
//	}
func SetupMockTest(t *testing.T) *MockTestConfig {
	t.Helper()

	// Log test category for reporting
	LogTestCategory(t, TestCategoryMock)

	// Create mock server
	server := mocks.NewServer()

	// Save original environment
	originalURL := os.Getenv("F5XC_API_URL")
	originalToken := os.Getenv("F5XC_API_TOKEN")

	// Configure environment to point to mock server
	_ = os.Setenv("F5XC_API_URL", server.URL())
	_ = os.Setenv("F5XC_API_TOKEN", "mock-token")

	return &MockTestConfig{
		Server:        server,
		OriginalURL:   originalURL,
		OriginalToken: originalToken,
	}
}

// Cleanup restores the original environment and closes the mock server
func (m *MockTestConfig) Cleanup() {
	// Restore original environment
	if m.OriginalURL != "" {
		_ = os.Setenv("F5XC_API_URL", m.OriginalURL)
	} else {
		_ = os.Unsetenv("F5XC_API_URL")
	}

	if m.OriginalToken != "" {
		_ = os.Setenv("F5XC_API_TOKEN", m.OriginalToken)
	} else {
		_ = os.Unsetenv("F5XC_API_TOKEN")
	}

	// Close mock server
	if m.Server != nil {
		m.Server.Close()
	}
}

// ProtoV6ProviderFactories returns provider factories for use in terraform-plugin-testing
func (m *MockTestConfig) ProtoV6ProviderFactories() map[string]func() (tfprotov6.ProviderServer, error) {
	return map[string]func() (tfprotov6.ProviderServer, error){
		"f5xc": providerserver.NewProtocol6WithError(provider.New("test")()),
	}
}

// MockProviderConfig returns a provider configuration block that uses the mock server
func (m *MockTestConfig) MockProviderConfig() string {
	return fmt.Sprintf(`
provider "f5xc" {
  api_url   = %q
  api_token = "mock-token"
}
`, m.Server.URL())
}

// PrePopulateResource adds a resource to the mock server before the test runs
func (m *MockTestConfig) PrePopulateResource(path string, response map[string]interface{}) {
	m.Server.SetResource(path, response)
}

// SetErrorForPath configures the mock server to return an error for a specific path
func (m *MockTestConfig) SetErrorForPath(path string, statusCode int, errorCode, message string) {
	m.Server.SetErrorResponse(path, statusCode, map[string]string{
		"code":    errorCode,
		"message": message,
	})
}

// MockTest runs a test with a mock server instead of the real F5 XC API.
// This is useful for testing resources that would otherwise require external credentials.
func MockTest(t *testing.T, testCase resource.TestCase) {
	t.Helper()

	mockCfg := SetupMockTest(t)
	defer mockCfg.Cleanup()

	// Override provider factories
	testCase.ProtoV6ProviderFactories = mockCfg.ProtoV6ProviderFactories()

	resource.Test(t, testCase)
}

// MockParallelTest runs a parallel test with a mock server
func MockParallelTest(t *testing.T, testCase resource.TestCase) {
	t.Helper()

	mockCfg := SetupMockTest(t)
	defer mockCfg.Cleanup()

	// Override provider factories
	testCase.ProtoV6ProviderFactories = mockCfg.ProtoV6ProviderFactories()

	resource.ParallelTest(t, testCase)
}

// SkipIfNoMockMode skips the test if F5XC_MOCK_MODE is not set.
// This allows running mock tests only when explicitly requested.
func SkipIfNoMockMode(t *testing.T) {
	t.Helper()

	if os.Getenv("F5XC_MOCK_MODE") == "" {
		t.Skip("Skipping mock test: F5XC_MOCK_MODE not set")
	}
}

// SkipIfRealAPI skips the test if running against the real API (F5XC_MOCK_MODE is not set).
// Use this for tests that have known API behavioral constraints that don't affect mock testing,
// such as ephemeral resource lifecycles, infrastructure dependencies, or staging environment limitations.
// The message should explain why this test cannot run against the real API.
func SkipIfRealAPI(t *testing.T, message string) {
	t.Helper()

	if os.Getenv("F5XC_MOCK_MODE") == "" {
		t.Skipf("Skipping on real API: %s", message)
	}
}

// IsMockMode returns true if mock mode is enabled
func IsMockMode() bool {
	return os.Getenv("F5XC_MOCK_MODE") != ""
}

// GetGlobalMockServer returns the global mock server, creating it if necessary.
// This ensures all tests share the same mock server when F5XC_MOCK_MODE is set.
func GetGlobalMockServer() *mocks.Server {
	globalMockServerOnce.Do(func() {
		globalMockServer = mocks.NewServer()

		// Store original environment variables
		globalMockServerMu.Lock()
		originalEnvVars["F5XC_API_URL"] = os.Getenv("F5XC_API_URL")
		originalEnvVars["F5XC_API_TOKEN"] = os.Getenv("F5XC_API_TOKEN")
		originalEnvVars["F5XC_API_P12_FILE"] = os.Getenv("F5XC_API_P12_FILE")
		originalEnvVars["F5XC_P12_PASSWORD"] = os.Getenv("F5XC_P12_PASSWORD")
		originalEnvVars["F5XC_API_CERT"] = os.Getenv("F5XC_API_CERT")
		originalEnvVars["F5XC_API_KEY"] = os.Getenv("F5XC_API_KEY")
		globalMockServerMu.Unlock()

		// Override environment to use mock server
		_ = os.Setenv("F5XC_API_URL", globalMockServer.URL())
		_ = os.Setenv("F5XC_API_TOKEN", "mock-token")
		// Clear P12/PEM credentials so the provider uses token auth
		_ = os.Unsetenv("F5XC_API_P12_FILE")
		_ = os.Unsetenv("F5XC_P12_PASSWORD")
		_ = os.Unsetenv("F5XC_API_CERT")
		_ = os.Unsetenv("F5XC_API_KEY")
	})
	return globalMockServer
}

// EnsureMockModeConfigured ensures that if F5XC_MOCK_MODE is set,
// the environment is configured to use the mock server.
// This should be called early in test setup.
func EnsureMockModeConfigured() {
	if IsMockMode() {
		GetGlobalMockServer()
	}
}

// RunWithMockOrReal runs a test with either mock or real API depending on environment.
// If TF_ACC and real credentials are set, uses real API.
// If F5XC_MOCK_MODE is set, uses mock server.
// Otherwise skips the test.
func RunWithMockOrReal(t *testing.T, testCase resource.TestCase, mockSetup func(*MockTestConfig)) {
	t.Helper()

	// Check if we can run real acceptance tests
	if os.Getenv(EnvTFAccTest) != "" {
		authMethod := DetectAuthMethod()
		if authMethod != AuthMethodNone && os.Getenv(EnvF5XCURL) != "" {
			// Run with real API
			testCase.ProtoV6ProviderFactories = ProtoV6ProviderFactories
			resource.Test(t, testCase)
			return
		}
	}

	// Check if mock mode is enabled
	if os.Getenv("F5XC_MOCK_MODE") != "" {
		mockCfg := SetupMockTest(t)
		defer mockCfg.Cleanup()

		// Allow test to customize mock setup
		if mockSetup != nil {
			mockSetup(mockCfg)
		}

		testCase.ProtoV6ProviderFactories = mockCfg.ProtoV6ProviderFactories()
		resource.Test(t, testCase)
		return
	}

	t.Skip("Skipping: neither TF_ACC with credentials nor F5XC_MOCK_MODE is set")
}

// MockResourceTestCase creates a test case configured for mock testing
func MockResourceTestCase(steps []resource.TestStep) resource.TestCase {
	return resource.TestCase{
		Steps: steps,
		// ProtoV6ProviderFactories will be set by MockTest/MockParallelTest
	}
}

// --- Helper functions for common mock scenarios ---

// SetupAWSCredentialsMock configures the mock server to provide fake AWS credentials
func (m *MockTestConfig) SetupAWSCredentialsMock(namespace, name string) {
	path := mocks.ResourcePath(namespace, "cloud_credentials", name)
	m.PrePopulateResource(path, mocks.CloudCredentialsResponse(namespace, name, "aws"))
}

// SetupAzureCredentialsMock configures the mock server to provide fake Azure credentials
func (m *MockTestConfig) SetupAzureCredentialsMock(namespace, name string) {
	path := mocks.ResourcePath(namespace, "cloud_credentials", name)
	m.PrePopulateResource(path, mocks.CloudCredentialsResponse(namespace, name, "azure"))
}

// SetupGCPCredentialsMock configures the mock server to provide fake GCP credentials
func (m *MockTestConfig) SetupGCPCredentialsMock(namespace, name string) {
	path := mocks.ResourcePath(namespace, "cloud_credentials", name)
	m.PrePopulateResource(path, mocks.CloudCredentialsResponse(namespace, name, "gcp"))
}

// SetupNamespaceMock configures the mock server to provide a namespace
func (m *MockTestConfig) SetupNamespaceMock(name string) {
	path := mocks.ResourcePath("system", "namespaces", name)
	m.PrePopulateResource(path, mocks.NamespaceResponse(name, nil, nil, ""))
}

// SetupAWSVPCSiteMock configures the mock server to provide an AWS VPC site
func (m *MockTestConfig) SetupAWSVPCSiteMock(namespace, name string, opts ...mocks.AWSVPCSiteOption) {
	path := mocks.ResourcePath(namespace, "aws_vpc_sites", name)
	m.PrePopulateResource(path, mocks.AWSVPCSiteResponse(namespace, name, opts...))
}

// SetupAzureVNETSiteMock configures the mock server to provide an Azure VNET site
func (m *MockTestConfig) SetupAzureVNETSiteMock(namespace, name string, opts ...mocks.AzureVNETSiteOption) {
	path := mocks.ResourcePath(namespace, "azure_vnet_sites", name)
	m.PrePopulateResource(path, mocks.AzureVNETSiteResponse(namespace, name, opts...))
}

// SetupGCPVPCSiteMock configures the mock server to provide a GCP VPC site
func (m *MockTestConfig) SetupGCPVPCSiteMock(namespace, name string, opts ...mocks.GCPVPCSiteOption) {
	path := mocks.ResourcePath(namespace, "gcp_vpc_sites", name)
	m.PrePopulateResource(path, mocks.GCPVPCSiteResponse(namespace, name, opts...))
}

// SetupOriginPoolMock configures the mock server to provide an origin pool
func (m *MockTestConfig) SetupOriginPoolMock(namespace, name string, opts ...mocks.OriginPoolOption) {
	path := mocks.ResourcePath(namespace, "origin_pools", name)
	m.PrePopulateResource(path, mocks.OriginPoolResponse(namespace, name, opts...))
}

// SetupHealthcheckMock configures the mock server to provide a healthcheck
func (m *MockTestConfig) SetupHealthcheckMock(namespace, name, healthcheckType string) {
	path := mocks.ResourcePath(namespace, "healthchecks", name)
	m.PrePopulateResource(path, mocks.HealthcheckResponse(namespace, name, healthcheckType))
}

// SetupAppFirewallMock configures the mock server to provide an app firewall
func (m *MockTestConfig) SetupAppFirewallMock(namespace, name string) {
	path := mocks.ResourcePath(namespace, "app_firewalls", name)
	m.PrePopulateResource(path, mocks.AppFirewallResponse(namespace, name))
}

// SetupHTTPLoadBalancerMock configures the mock server to provide an HTTP load balancer
func (m *MockTestConfig) SetupHTTPLoadBalancerMock(namespace, name string, domains []string) {
	path := mocks.ResourcePath(namespace, "http_loadbalancers", name)
	m.PrePopulateResource(path, mocks.HTTPLoadBalancerResponse(namespace, name, domains))
}

// SetupGenericResourceMock configures the mock server to provide a generic resource
func (m *MockTestConfig) SetupGenericResourceMock(namespace, resourceType, name string, spec map[string]interface{}) {
	path := mocks.ResourcePath(namespace, resourceType, name)
	m.PrePopulateResource(path, mocks.GenericResourceResponse(namespace, name, resourceType, spec))
}

// SimulateAPIError configures the mock server to return an error for all operations on a resource type
func (m *MockTestConfig) SimulateAPIError(namespace, resourceType string, statusCode int, errorCode, message string) {
	// Set error for list operations
	listPath := mocks.ListPath(namespace, resourceType)
	m.Server.SetErrorResponse(listPath, statusCode, map[string]string{
		"code":    errorCode,
		"message": message,
	})
}

// Simulate501NotImplemented configures the mock server to return 501 for a specific resource type
// This is useful for testing resources where the API returns 501 in certain environments
func (m *MockTestConfig) Simulate501NotImplemented(namespace, resourceType string) {
	m.SimulateAPIError(namespace, resourceType, 501, "NOT_IMPLEMENTED", "Operation not implemented")
}

// Simulate403Forbidden configures the mock server to return 403 for a specific resource type
func (m *MockTestConfig) Simulate403Forbidden(namespace, resourceType string) {
	m.SimulateAPIError(namespace, resourceType, 403, "FORBIDDEN", "Access denied")
}
