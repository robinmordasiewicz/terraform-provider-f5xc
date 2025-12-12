# Mock Testing Infrastructure

This package provides a mock F5 XC API server for testing Terraform provider resources without requiring real cloud credentials or F5 XC infrastructure.

## Overview

The mock testing infrastructure enables:
- Testing resources that require external credentials (AWS, Azure, GCP)
- Testing resources that require special permissions (tenant configuration)
- Faster test execution without network latency
- CI/CD testing without secrets
- Error scenario testing (501, 403, timeouts)

## Quick Start

### Running Mock Tests

```bash
# Run all mock tests
VES_MOCK_MODE=1 go test -v ./internal/provider/ -run TestMock -timeout 5m

# Run specific mock test
VES_MOCK_MODE=1 go test -v ./internal/provider/ -run TestMockAWSVPCSiteResource_basic -timeout 5m
```

### Writing a Mock Test

```go
package provider_test

import (
    "testing"

    "github.com/hashicorp/terraform-plugin-testing/helper/resource"
    "github.com/f5xc/terraform-provider-f5xc/internal/acctest"
    "github.com/f5xc/terraform-provider-f5xc/internal/mocks"
)

func TestMockMyResource_basic(t *testing.T) {
    // Skip if not in mock mode
    acctest.SkipIfNoMockMode(t)

    // Setup mock server
    mockCfg := acctest.SetupMockTest(t)
    defer mockCfg.Cleanup()

    // Pre-populate dependencies if needed
    mockCfg.SetupNamespaceMock("my-namespace")

    // Run the test
    resource.Test(t, resource.TestCase{
        ProtoV6ProviderFactories: mockCfg.ProtoV6ProviderFactories(),
        Steps: []resource.TestStep{
            {
                Config: testAccMyResourceConfig(mockCfg),
                Check: resource.ComposeAggregateTestCheckFunc(
                    acctest.CheckResourceExists("f5xc_my_resource.test"),
                    resource.TestCheckResourceAttr("f5xc_my_resource.test", "name", "test"),
                ),
            },
        },
    })
}

func testAccMyResourceConfig(mockCfg *acctest.MockTestConfig) string {
    return acctest.ConfigCompose(
        mockCfg.MockProviderConfig(),
        `
resource "f5xc_my_resource" "test" {
  name      = "test"
  namespace = "my-namespace"
}
`)
}
```

## Package Structure

```
internal/mocks/
├── server.go      # Mock HTTP server implementation
├── fixtures.go    # Response generators for F5 XC resources
├── server_test.go # Unit tests for mock server
└── README.md      # This documentation

internal/acctest/
├── mock_helpers.go # Test helpers for mock-based testing
└── helpers.go      # Existing acceptance test helpers
```

## Core Components

### Mock Server (`server.go`)

The `Server` struct provides a complete mock F5 XC API:

```go
// Create a new mock server
server := mocks.NewServer()
defer server.Close()

// Get server URL for client configuration
url := server.URL()

// Pre-populate resources
server.SetResource("/api/config/namespaces/system/my_resources/name", response)

// Inject errors for testing error handling
server.SetErrorResponse("/api/path", 500, errorBody)

// Add custom handlers for complex scenarios
server.SetHandler("/api/custom/.*", customHandlerFunc)

// Inspect request log for verification
requests := server.GetRequestLog()
```

### Mock Fixtures (`fixtures.go`)

Pre-built response generators for common F5 XC resources:

| Function | Purpose |
|----------|---------|
| `NamespaceResponse(name, labels, annotations, description)` | Namespace resource |
| `AWSVPCSiteResponse(namespace, name, opts...)` | AWS VPC Site with options |
| `AzureVNETSiteResponse(namespace, name, opts...)` | Azure VNET Site |
| `GCPVPCSiteResponse(namespace, name, opts...)` | GCP VPC Site |
| `CloudCredentialsResponse(namespace, name, cloudType)` | Cloud credentials (aws/azure/gcp) |
| `OriginPoolResponse(namespace, name, opts...)` | Origin pool |
| `HTTPLoadBalancerResponse(namespace, name, domains)` | HTTP load balancer |
| `AppFirewallResponse(namespace, name)` | App firewall |
| `HealthcheckResponse(namespace, name, healthcheckType)` | Healthcheck |
| `TenantConfigurationResponse(name)` | Tenant configuration |
| `GenericResourceResponse(namespace, name, resourceType, spec)` | Any resource |

### Test Helpers (`mock_helpers.go`)

High-level helpers for common test patterns:

```go
// Setup mock test environment
mockCfg := acctest.SetupMockTest(t)
defer mockCfg.Cleanup()

// Get provider factories for test case
factories := mockCfg.ProtoV6ProviderFactories()

// Get provider config block for Terraform
config := mockCfg.MockProviderConfig()

// Setup common dependencies
mockCfg.SetupNamespaceMock("ns-name")
mockCfg.SetupAWSCredentialsMock("ns", "creds-name")
mockCfg.SetupAWSVPCSiteMock("ns", "site-name", mocks.WithAWSRegion("us-west-2"))

// Simulate errors
mockCfg.Simulate501NotImplemented("system", "resource_type")
mockCfg.Simulate403Forbidden("system", "resource_type")
```

## Test Patterns

### Pattern 1: Basic Resource CRUD

Test create, read, update, delete operations:

```go
func TestMockResource_basic(t *testing.T) {
    acctest.SkipIfNoMockMode(t)

    mockCfg := acctest.SetupMockTest(t)
    defer mockCfg.Cleanup()

    resource.Test(t, resource.TestCase{
        ProtoV6ProviderFactories: mockCfg.ProtoV6ProviderFactories(),
        Steps: []resource.TestStep{
            // Create
            {
                Config: configCreate,
                Check:  checkCreate,
            },
            // Update
            {
                Config: configUpdate,
                Check:  checkUpdate,
            },
            // Import
            {
                ResourceName:      resourceName,
                ImportState:       true,
                ImportStateVerify: true,
            },
        },
    })
}
```

### Pattern 2: Resources with Dependencies

Test resources that depend on other resources:

```go
func TestMockAWSVPCSite_withCredentials(t *testing.T) {
    acctest.SkipIfNoMockMode(t)

    mockCfg := acctest.SetupMockTest(t)
    defer mockCfg.Cleanup()

    // Pre-populate dependencies
    mockCfg.SetupNamespaceMock(nsName)
    mockCfg.SetupAWSCredentialsMock(nsName, credsName)

    resource.Test(t, resource.TestCase{
        ProtoV6ProviderFactories: mockCfg.ProtoV6ProviderFactories(),
        Steps: []resource.TestStep{
            {
                Config: configWithCredentials,
                Check:  checks,
            },
        },
    })
}
```

### Pattern 3: Error Handling

Test how resources handle API errors:

```go
func TestMockResource_errorHandling(t *testing.T) {
    acctest.SkipIfNoMockMode(t)

    mockCfg := acctest.SetupMockTest(t)
    defer mockCfg.Cleanup()

    // Simulate 501 NOT_IMPLEMENTED
    mockCfg.Simulate501NotImplemented("system", "resources")

    resource.Test(t, resource.TestCase{
        ProtoV6ProviderFactories: mockCfg.ProtoV6ProviderFactories(),
        Steps: []resource.TestStep{
            {
                Config:      config,
                ExpectError: acctest.MustCompileRegexp(`(?i)(501|not implemented)`),
            },
        },
    })
}
```

### Pattern 4: Data Source Testing

Test data sources that read existing resources:

```go
func TestMockDataSource_basic(t *testing.T) {
    acctest.SkipIfNoMockMode(t)

    mockCfg := acctest.SetupMockTest(t)
    defer mockCfg.Cleanup()

    // Pre-populate the resource to read
    path := mocks.ResourcePath("system", "resources", "existing")
    mockCfg.PrePopulateResource(path, mocks.GenericResourceResponse(...))

    resource.Test(t, resource.TestCase{
        ProtoV6ProviderFactories: mockCfg.ProtoV6ProviderFactories(),
        Steps: []resource.TestStep{
            {
                Config: dataSourceConfig,
                Check:  dataSourceChecks,
            },
        },
    })
}
```

### Pattern 5: Hybrid Mock/Real Testing

Test with mock OR real API depending on environment:

```go
func TestResource_hybridMode(t *testing.T) {
    testCase := resource.TestCase{
        Steps: []resource.TestStep{
            {Config: config, Check: checks},
        },
    }

    // Runs with real API if TF_ACC + credentials set
    // Runs with mock if VES_MOCK_MODE=1 set
    // Skips if neither available
    acctest.RunWithMockOrReal(t, testCase, func(mockCfg *acctest.MockTestConfig) {
        // Optional: customize mock setup
        mockCfg.SetupNamespaceMock("test")
    })
}
```

## Adding New Mock Fixtures

To add support for a new resource type:

1. Add response generator in `fixtures.go`:

```go
func MyResourceResponse(namespace, name string) map[string]interface{} {
    return map[string]interface{}{
        "metadata": map[string]interface{}{
            "name":      name,
            "namespace": namespace,
            "uid":       generateUID(),
        },
        "spec": map[string]interface{}{
            // Resource-specific fields
        },
        "system_metadata": systemMetadata(),
    }
}
```

2. Add setup helper in `mock_helpers.go`:

```go
func (m *MockTestConfig) SetupMyResourceMock(namespace, name string) {
    path := mocks.ResourcePath(namespace, "my_resources", name)
    m.PrePopulateResource(path, mocks.MyResourceResponse(namespace, name))
}
```

3. Use in tests:

```go
mockCfg.SetupMyResourceMock("system", "test-resource")
```

## Environment Variables

| Variable | Purpose |
|----------|---------|
| `VES_MOCK_MODE=1` | Enable mock testing mode |
| `TF_ACC=1` | Enable real acceptance tests |
| `VES_API_URL` | Real API URL (for non-mock tests) |
| `VES_API_TOKEN` | Real API token (for non-mock tests) |

## Test Categories and Reporting

Tests are automatically categorized by their naming convention:

| Prefix | Category | Description |
|--------|----------|-------------|
| `TestAcc*` | `REAL_API` | Tests against real F5 XC API endpoints |
| `TestMock*` | `MOCK_API` | Tests against local mock server |
| `Test*` (other) | `UNIT` | Unit tests without external dependencies |

### Makefile Targets

```bash
# Run only real API tests
make testacc-real

# Run only mock tests
make testacc-mock

# Run both with categorized report
make testacc-all

# Generate report from previous run
make test-report

# Generate markdown report
make test-report-md
```

### Using the Test Report Tool

```bash
# Generate text report
go test -json ./internal/provider/... | go run tools/test-report/main.go

# Generate markdown report
go test -json ./internal/provider/... | go run tools/test-report/main.go -format=markdown

# Generate JSON report
go test -json ./internal/provider/... | go run tools/test-report/main.go -format=json

# Show all tests, not just summary
go test -json ./internal/provider/... | go run tools/test-report/main.go -all
```

### Example Report Output

```
==============================================================================
                           TEST SUMMARY REPORT
==============================================================================
Timestamp: 2025-12-01T12:00:00Z
------------------------------------------------------------------------------
[PASS] TOTAL: 50 tests | 45 passed | 2 failed | 3 skipped
------------------------------------------------------------------------------
BY CATEGORY:
  [PASS] REAL_API  :   30 tests |   28 passed |    0 failed |    2 skipped
  [FAIL] MOCK_API  :   12 tests |   10 passed |    2 failed |    0 skipped
  [PASS] UNIT      :    8 tests |    7 passed |    0 failed |    1 skipped
------------------------------------------------------------------------------
FAILED TESTS:
  [MOCK_API] TestMockResource_errorHandling (2.50s)
  [MOCK_API] TestMockResource_timeout (5.00s)
------------------------------------------------------------------------------
==============================================================================
```

## Best Practices

1. **Always use `SkipIfNoMockMode(t)`** at the start of mock tests
2. **Always call `defer mockCfg.Cleanup()`** to restore environment
3. **Pre-populate dependencies** before testing dependent resources
4. **Use `ConfigCompose`** to combine provider config with resource config
5. **Test error scenarios** using `Simulate*` helpers
6. **Follow naming convention**: `TestMock{Resource}_{scenario}`

## Troubleshooting

### "Missing metadata" error
The mock server expects requests to include a `metadata` block with `name`. Ensure your Terraform config includes required fields.

### "Resource not found" error
Pre-populate the resource before running the test:
```go
mockCfg.PrePopulateResource(path, response)
```

### Tests not running
Check that `VES_MOCK_MODE=1` is set:
```bash
VES_MOCK_MODE=1 go test -v ./internal/provider/ -run TestMock
```

### Cascade delete not working
The mock server handles `/cascade_delete` endpoints specially. Ensure the path pattern matches F5 XC conventions.
