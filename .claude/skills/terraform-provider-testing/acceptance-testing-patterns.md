# Terraform Provider Acceptance Testing Patterns

Comprehensive Go SDK patterns for testing Terraform providers using `terraform-plugin-testing`.

## Environment Setup

### Required Environment Variables

```bash
export TF_ACC=1                    # Enable acceptance tests
export TF_LOG=DEBUG                # Optional: Debug logging
export TF_LOG_PATH=/tmp/tf.log     # Optional: Log to file

# Provider-specific (example for AWS)
export AWS_REGION=us-west-2
export AWS_ACCESS_KEY_ID=xxx
export AWS_SECRET_ACCESS_KEY=xxx
```

### Provider Factory Setup

#### Protocol Version 6 (Plugin Framework)

```go
package provider_test

import (
    "github.com/hashicorp/terraform-plugin-framework/providerserver"
    "github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
    "example": providerserver.NewProtocol6WithError(New("test")()),
}
```

#### Protocol Version 5 (SDKv2)

```go
package provider_test

import (
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProviders = map[string]*schema.Provider{
    "example": Provider(),
}

var testAccProvider = Provider()
```

## TestCase Structure

### Complete TestCase Example

```go
func TestAccExampleThing_basic(t *testing.T) {
    t.Parallel()

    rName := acctest.RandomWithPrefix("tf-acc-test")

    resource.ParallelTest(t, resource.TestCase{
        // Prerequisites check
        PreCheck: func() {
            testAccPreCheck(t)
            testAccPreCheckRegion(t, "us-west-2")
        },

        // Provider instantiation
        ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,

        // Error handling
        ErrorCheck: acctest.ErrorCheck(t, names.ServiceEndpointID),

        // Cleanup verification
        CheckDestroy: testAccCheckExampleThingDestroy,

        // Test steps
        Steps: []resource.TestStep{
            {
                Config: testAccExampleThingConfig_basic(rName),
                Check: resource.ComposeAggregateTestCheckFunc(
                    testAccCheckExampleThingExists("example_thing.test"),
                    resource.TestCheckResourceAttr("example_thing.test", "name", rName),
                    resource.TestCheckResourceAttrSet("example_thing.test", "id"),
                    resource.TestCheckResourceAttr("example_thing.test", "enabled", "true"),
                ),
                ConfigStateChecks: []statecheck.StateCheck{
                    statecheck.ExpectKnownValue(
                        "example_thing.test",
                        tfjsonpath.New("name"),
                        knownvalue.StringExact(rName),
                    ),
                },
            },
            {
                Config:            testAccExampleThingConfig_basic(rName),
                ImportState:       true,
                ImportStateVerify: true,
            },
        },
    })
}
```

## PreCheck Patterns

### Basic PreCheck

```go
func testAccPreCheck(t *testing.T) {
    if v := os.Getenv("EXAMPLE_API_KEY"); v == "" {
        t.Fatal("EXAMPLE_API_KEY must be set for acceptance tests")
    }
    if v := os.Getenv("EXAMPLE_API_ENDPOINT"); v == "" {
        t.Fatal("EXAMPLE_API_ENDPOINT must be set for acceptance tests")
    }
}
```

### Region PreCheck

```go
func testAccPreCheckRegion(t *testing.T, region string) {
    if os.Getenv("AWS_REGION") != region {
        t.Skipf("Skipping test; AWS_REGION is not %s", region)
    }
}
```

### Service Availability PreCheck

```go
func testAccPreCheckService(t *testing.T) {
    conn := testAccProvider.Meta().(*Client)

    _, err := conn.ListResources(&ListResourcesInput{MaxResults: 1})
    if err != nil {
        t.Skipf("Service not available: %s", err)
    }
}
```

### Long-Running Test Guard

```go
func TestAccExampleThing_longRunning(t *testing.T) {
    if testing.Short() {
        t.Skip("skipping long-running test in short mode")
    }

    t.Parallel()
    // ... test implementation
}
```

## TestStep Patterns

### Create and Verify

```go
{
    Config: testAccConfig,
    Check: resource.ComposeTestCheckFunc(
        testAccCheckExists("example_thing.test"),
        resource.TestCheckResourceAttr("example_thing.test", "name", "test"),
    ),
}
```

### Update Configuration

```go
{
    Config: testAccConfigUpdated,
    Check: resource.ComposeTestCheckFunc(
        testAccCheckExists("example_thing.test"),
        resource.TestCheckResourceAttr("example_thing.test", "name", "test-updated"),
    ),
}
```

### Import State

```go
{
    Config:            testAccConfig,
    ImportState:       true,
    ImportStateVerify: true,
    ImportStateVerifyIgnore: []string{
        "password",        // Sensitive field not returned
        "force_destroy",   // Local-only attribute
    },
}
```

### Expect Error

```go
{
    Config:      testAccInvalidConfig,
    ExpectError: regexp.MustCompile(`name must be between 1 and 64 characters`),
}
```

### Expect Non-Empty Plan

```go
{
    Config:             testAccConfigIncomplete,
    ExpectNonEmptyPlan: true,
    Check: resource.ComposeTestCheckFunc(
        testAccCheckExists("example_thing.test"),
    ),
}
```

## State Check Functions (Modern Pattern)

### ExpectKnownValue

```go
import (
    "github.com/hashicorp/terraform-plugin-testing/knownvalue"
    "github.com/hashicorp/terraform-plugin-testing/statecheck"
    "github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

ConfigStateChecks: []statecheck.StateCheck{
    // String exact match
    statecheck.ExpectKnownValue(
        "example_thing.test",
        tfjsonpath.New("name"),
        knownvalue.StringExact("my-thing"),
    ),

    // String regex match
    statecheck.ExpectKnownValue(
        "example_thing.test",
        tfjsonpath.New("arn"),
        knownvalue.StringRegexp(regexp.MustCompile(`^arn:aws:example:`)),
    ),

    // Boolean check
    statecheck.ExpectKnownValue(
        "example_thing.test",
        tfjsonpath.New("enabled"),
        knownvalue.Bool(true),
    ),

    // Integer check
    statecheck.ExpectKnownValue(
        "example_thing.test",
        tfjsonpath.New("count"),
        knownvalue.Int64Exact(5),
    ),

    // Null check
    statecheck.ExpectKnownValue(
        "example_thing.test",
        tfjsonpath.New("optional_field"),
        knownvalue.Null(),
    ),

    // Not null check
    statecheck.ExpectKnownValue(
        "example_thing.test",
        tfjsonpath.New("id"),
        knownvalue.NotNull(),
    ),
}
```

### Nested Attribute Paths

```go
// Check nested object attribute
statecheck.ExpectKnownValue(
    "example_thing.test",
    tfjsonpath.New("config").AtMapKey("setting"),
    knownvalue.StringExact("value"),
)

// Check list element
statecheck.ExpectKnownValue(
    "example_thing.test",
    tfjsonpath.New("items").AtSliceIndex(0).AtMapKey("name"),
    knownvalue.StringExact("first-item"),
)
```

### ExpectKnownOutputValue

```go
statecheck.ExpectKnownOutputValue(
    "thing_id",
    knownvalue.StringRegexp(regexp.MustCompile(`^thing-`)),
)

statecheck.ExpectKnownOutputValueAtPath(
    "thing_config",
    tfjsonpath.New("settings").AtMapKey("enabled"),
    knownvalue.Bool(true),
)
```

## Custom State Checks

### Basic Custom Check

```go
type expectResourceExists struct {
    resourceAddress string
}

func (e expectResourceExists) CheckState(ctx context.Context, req statecheck.CheckStateRequest, resp *statecheck.CheckStateResponse) {
    var found bool

    for _, resource := range req.State.Values.RootModule.Resources {
        if resource.Address == e.resourceAddress {
            found = true
            break
        }
    }

    if !found {
        resp.Error = fmt.Errorf("resource %s not found in state", e.resourceAddress)
    }
}

func ExpectResourceExists(resourceAddress string) statecheck.StateCheck {
    return expectResourceExists{resourceAddress: resourceAddress}
}
```

### Custom Check with Remote API Verification

```go
type verifyRemoteState struct {
    resourceAddress string
    apiClient       APIClient
}

func (v verifyRemoteState) CheckState(ctx context.Context, req statecheck.CheckStateRequest, resp *statecheck.CheckStateResponse) {
    // Find resource in state
    var resourceID string
    for _, resource := range req.State.Values.RootModule.Resources {
        if resource.Address == v.resourceAddress {
            id, ok := resource.AttributeValues["id"].(string)
            if !ok {
                resp.Error = fmt.Errorf("id attribute not found")
                return
            }
            resourceID = id
            break
        }
    }

    if resourceID == "" {
        resp.Error = fmt.Errorf("resource %s not found", v.resourceAddress)
        return
    }

    // Verify with remote API
    remoteResource, err := v.apiClient.Get(ctx, resourceID)
    if err != nil {
        resp.Error = fmt.Errorf("failed to get remote resource: %w", err)
        return
    }

    if remoteResource == nil {
        resp.Error = fmt.Errorf("resource %s not found in remote API", resourceID)
        return
    }
}

func VerifyRemoteState(resourceAddress string, client APIClient) statecheck.StateCheck {
    return verifyRemoteState{
        resourceAddress: resourceAddress,
        apiClient:       client,
    }
}
```

## Legacy Check Functions

### TestCheckResourceAttr

```go
// Exact value match
resource.TestCheckResourceAttr("example_thing.test", "name", "expected-name")

// Check attribute is set (any value)
resource.TestCheckResourceAttrSet("example_thing.test", "id")

// Check attribute pair (values match between resources)
resource.TestCheckResourceAttrPair(
    "example_thing.test", "vpc_id",
    "aws_vpc.test", "id",
)

// Check no attribute exists
resource.TestCheckNoResourceAttr("example_thing.test", "deprecated_field")
```

### ComposeTestCheckFunc

```go
Check: resource.ComposeTestCheckFunc(
    testAccCheckExists("example_thing.test"),
    resource.TestCheckResourceAttr("example_thing.test", "name", rName),
    resource.TestCheckResourceAttr("example_thing.test", "enabled", "true"),
    resource.TestCheckResourceAttrSet("example_thing.test", "arn"),
),

// Aggregate version (continues even if one check fails)
Check: resource.ComposeAggregateTestCheckFunc(
    // ... same checks
),
```

### Custom Check Function

```go
func testAccCheckExampleThingExists(resourceName string) resource.TestCheckFunc {
    return func(s *terraform.State) error {
        rs, ok := s.RootModule().Resources[resourceName]
        if !ok {
            return fmt.Errorf("resource not found: %s", resourceName)
        }

        if rs.Primary.ID == "" {
            return fmt.Errorf("resource ID not set")
        }

        conn := testAccProvider.Meta().(*Client)
        _, err := conn.GetThing(rs.Primary.ID)
        if err != nil {
            return fmt.Errorf("error getting thing: %w", err)
        }

        return nil
    }
}
```

## CheckDestroy Implementation

```go
func testAccCheckExampleThingDestroy(s *terraform.State) error {
    conn := testAccProvider.Meta().(*Client)

    for _, rs := range s.RootModule().Resources {
        if rs.Type != "example_thing" {
            continue
        }

        _, err := conn.GetThing(rs.Primary.ID)
        if err == nil {
            return fmt.Errorf("Example Thing %s still exists", rs.Primary.ID)
        }

        // Check for expected "not found" error
        if !isResourceNotFoundError(err) {
            return fmt.Errorf("unexpected error checking thing destroy: %w", err)
        }
    }

    return nil
}
```

## Test Configuration Helpers

### Basic Config Function

```go
func testAccExampleThingConfig_basic(rName string) string {
    return fmt.Sprintf(`
resource "example_thing" "test" {
  name = %[1]q
}
`, rName)
}
```

### Config with Dependencies

```go
func testAccExampleThingConfig_withVPC(rName string) string {
    return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.0.0.0/16"

  tags = {
    Name = %[1]q
  }
}

resource "example_thing" "test" {
  name   = %[1]q
  vpc_id = aws_vpc.test.id
}
`, rName)
}
```

### Config Composition

```go
func testAccExampleThingConfig_base(rName string) string {
    return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.0.0.0/16"
  tags = { Name = %[1]q }
}
`, rName)
}

func testAccExampleThingConfig_full(rName, description string) string {
    return acctest.ConfigCompose(
        testAccExampleThingConfig_base(rName),
        fmt.Sprintf(`
resource "example_thing" "test" {
  name        = %[1]q
  description = %[2]q
  vpc_id      = aws_vpc.test.id
}
`, rName, description),
    )
}
```

## Test Naming Conventions

### Resource Tests

```go
// Basic test - minimal configuration
func TestAccExampleThing_basic(t *testing.T)

// Full configuration with all attributes
func TestAccExampleThing_full(t *testing.T)

// Update test - create then update
func TestAccExampleThing_update(t *testing.T)

// Disappears test - handle external deletion
func TestAccExampleThing_disappears(t *testing.T)

// Specific attribute tests
func TestAccExampleThing_tags(t *testing.T)
func TestAccExampleThing_description(t *testing.T)

// Error handling tests
func TestAccExampleThing_invalidName(t *testing.T)
```

### Data Source Tests

```go
func TestAccExampleThingDataSource_basic(t *testing.T)
func TestAccExampleThingDataSource_byName(t *testing.T)
func TestAccExampleThingDataSource_byID(t *testing.T)
```

## Test Sweepers

### Sweeper Registration

```go
func init() {
    resource.AddTestSweepers("example_thing", &resource.Sweeper{
        Name: "example_thing",
        F:    sweepThings,
        Dependencies: []string{
            "example_dependent_thing",
        },
    })
}

func sweepThings(region string) error {
    client, err := sharedClientForRegion(region)
    if err != nil {
        return fmt.Errorf("error getting client: %w", err)
    }

    conn := client.(*Client)
    things, err := conn.ListThings(&ListThingsInput{})
    if err != nil {
        return fmt.Errorf("error listing things: %w", err)
    }

    for _, thing := range things {
        // Only delete test resources
        if !strings.HasPrefix(thing.Name, "tf-acc-test") {
            continue
        }

        _, err := conn.DeleteThing(&DeleteThingInput{ID: thing.ID})
        if err != nil {
            log.Printf("[ERROR] Error deleting thing %s: %s", thing.ID, err)
            continue
        }
    }

    return nil
}
```

### Running Sweepers

```bash
go test -v ./internal/service/example/ -sweep=us-west-2 -sweep-run=example_thing
```

## Serial Test Execution

### Test Groups

```go
func TestAccExampleThing_serial(t *testing.T) {
    t.Parallel()

    testCases := map[string]map[string]func(t *testing.T){
        "Thing": {
            "basic":      testAccExampleThing_basic,
            "disappears": testAccExampleThing_disappears,
            "update":     testAccExampleThing_update,
        },
        "OtherResource": {
            "basic": testAccOtherResource_basic,
        },
    }

    acctest.RunSerialTests2Levels(t, testCases, 0)
}

func testAccExampleThing_basic(t *testing.T) {
    // Implementation without t.Parallel()
}
```

## Best Practices Summary

1. **Always use `t.Parallel()`** and `resource.ParallelTest()` for concurrent execution
2. **Generate unique names** with `acctest.RandomWithPrefix("tf-acc-test")`
3. **Implement CheckDestroy** to verify cleanup
4. **Use ConfigStateChecks** for modern state validation
5. **Add import tests** for all resources
6. **Implement sweepers** for resource cleanup
7. **Skip long tests** with `-short` flag support
8. **Organize configs** with composition helpers
9. **Test edge cases** with ExpectError patterns
10. **Verify remote state** with custom checks when needed

## Go Testing Best Practices

### Table-Driven Tests for Multiple Scenarios

```go
func TestAccExampleResource_variations(t *testing.T) {
    t.Parallel()

    testCases := map[string]struct {
        config        func(string) string
        checks        []resource.TestCheckFunc
        expectError   *regexp.Regexp
        skipCondition func() (bool, string)
    }{
        "basic": {
            config: testAccExampleConfig_basic,
            checks: []resource.TestCheckFunc{
                resource.TestCheckResourceAttrSet("example_resource.test", "id"),
            },
        },
        "with_optional": {
            config: testAccExampleConfig_withOptional,
            checks: []resource.TestCheckFunc{
                resource.TestCheckResourceAttr("example_resource.test", "optional", "value"),
            },
        },
        "invalid_name": {
            config:      testAccExampleConfig_invalidName,
            expectError: regexp.MustCompile(`name must be alphanumeric`),
        },
    }

    for name, tc := range testCases {
        tc := tc // capture range variable
        t.Run(name, func(t *testing.T) {
            t.Parallel()

            if tc.skipCondition != nil {
                if skip, msg := tc.skipCondition(); skip {
                    t.Skip(msg)
                }
            }

            rName := acctest.RandomWithPrefix("tf-acc-test")

            steps := []resource.TestStep{
                {
                    Config:      tc.config(rName),
                    Check:       resource.ComposeAggregateTestCheckFunc(tc.checks...),
                    ExpectError: tc.expectError,
                },
            }

            resource.ParallelTest(t, resource.TestCase{
                PreCheck:                 func() { testAccPreCheck(t) },
                ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
                CheckDestroy:             testAccCheckExampleDestroy,
                Steps:                    steps,
            })
        })
    }
}
```

### Error Handling in Test Helpers

```go
// ✅ Wrap errors with context
func testAccCheckResourceExists(resourceName string) resource.TestCheckFunc {
    return func(s *terraform.State) error {
        rs, ok := s.RootModule().Resources[resourceName]
        if !ok {
            return fmt.Errorf("resource not found in state: %s", resourceName)
        }

        if rs.Primary.ID == "" {
            return fmt.Errorf("resource %s has empty ID", resourceName)
        }

        conn := testAccProvider.Meta().(*Client)
        _, err := conn.GetResource(context.Background(), rs.Primary.ID)
        if err != nil {
            return fmt.Errorf("error fetching resource %s (id: %s): %w",
                resourceName, rs.Primary.ID, err)
        }

        return nil
    }
}

// ✅ Use custom error types for better diagnostics
type TestCheckError struct {
    ResourceName string
    CheckType    string
    Details      string
    Underlying   error
}

func (e *TestCheckError) Error() string {
    if e.Underlying != nil {
        return fmt.Sprintf("%s check failed for %s: %s (%v)",
            e.CheckType, e.ResourceName, e.Details, e.Underlying)
    }
    return fmt.Sprintf("%s check failed for %s: %s",
        e.CheckType, e.ResourceName, e.Details)
}
```

### Context Propagation in Tests

```go
// ✅ Use context with timeout for API operations
func testAccCheckResourceExistsWithTimeout(resourceName string) resource.TestCheckFunc {
    return func(s *terraform.State) error {
        ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
        defer cancel()

        rs, ok := s.RootModule().Resources[resourceName]
        if !ok {
            return fmt.Errorf("resource not found: %s", resourceName)
        }

        conn := testAccProvider.Meta().(*Client)
        _, err := conn.GetResource(ctx, rs.Primary.ID)
        if err != nil {
            if ctx.Err() == context.DeadlineExceeded {
                return fmt.Errorf("timeout waiting for resource %s", resourceName)
            }
            return fmt.Errorf("error getting resource: %w", err)
        }

        return nil
    }
}
```

### Benchmarking Test Operations

```go
func BenchmarkTestAccCheckExists(b *testing.B) {
    // Setup: create a real resource first
    client := testAccProvider.Meta().(*Client)
    resource, err := client.CreateResource(context.Background(), CreateRequest{
        Name: "bench-test-" + acctest.RandString(8),
    })
    if err != nil {
        b.Fatal(err)
    }
    defer client.DeleteResource(context.Background(), resource.ID)

    // Create mock state
    state := &terraform.State{
        Modules: []*terraform.ModuleState{
            {
                Path: []string{"root"},
                Resources: map[string]*terraform.ResourceState{
                    "example_resource.test": {
                        Primary: &terraform.InstanceState{
                            ID: resource.ID,
                        },
                    },
                },
            },
        },
    }

    b.ResetTimer()
    b.ReportAllocs()

    for i := 0; i < b.N; i++ {
        checkFunc := testAccCheckResourceExists("example_resource.test")
        if err := checkFunc(state); err != nil {
            b.Fatal(err)
        }
    }
}
```

### Interface-Based Testing for Mocking

```go
// ✅ Define interface for testable operations
type ResourceClient interface {
    GetResource(ctx context.Context, id string) (*Resource, error)
    CreateResource(ctx context.Context, req CreateRequest) (*Resource, error)
    DeleteResource(ctx context.Context, id string) error
}

// ✅ Mock implementation for unit tests
type MockResourceClient struct {
    GetResourceFunc    func(ctx context.Context, id string) (*Resource, error)
    CreateResourceFunc func(ctx context.Context, req CreateRequest) (*Resource, error)
    DeleteResourceFunc func(ctx context.Context, id string) error
}

func (m *MockResourceClient) GetResource(ctx context.Context, id string) (*Resource, error) {
    if m.GetResourceFunc != nil {
        return m.GetResourceFunc(ctx, id)
    }
    return &Resource{ID: id}, nil
}

// ✅ Use in tests
func TestResourceLogic(t *testing.T) {
    mock := &MockResourceClient{
        GetResourceFunc: func(ctx context.Context, id string) (*Resource, error) {
            if id == "not-found" {
                return nil, ErrNotFound
            }
            return &Resource{ID: id, Name: "test"}, nil
        },
    }

    // Test with mock...
}
```

### Test Helpers with t.Helper()

```go
// ✅ Mark helper functions with t.Helper()
func assertResourceExists(t *testing.T, s *terraform.State, resourceName string) {
    t.Helper()

    rs, ok := s.RootModule().Resources[resourceName]
    if !ok {
        t.Fatalf("resource not found: %s", resourceName)
    }
    if rs.Primary.ID == "" {
        t.Fatalf("resource %s has empty ID", resourceName)
    }
}

func assertNoError(t *testing.T, err error, msg string) {
    t.Helper()
    if err != nil {
        t.Fatalf("%s: %v", msg, err)
    }
}

func assertEqual(t *testing.T, expected, actual interface{}, msg string) {
    t.Helper()
    if !reflect.DeepEqual(expected, actual) {
        t.Errorf("%s: expected %v, got %v", msg, expected, actual)
    }
}
```

### Test Cleanup with t.Cleanup()

```go
func TestAccResource_withCleanup(t *testing.T) {
    t.Parallel()

    // Create prerequisite resource
    client := testAccProvider.Meta().(*Client)
    prereq, err := client.CreateResource(context.Background(), CreateRequest{
        Name: "prereq-" + acctest.RandString(8),
    })
    if err != nil {
        t.Fatal(err)
    }

    // Register cleanup - runs even if test fails
    t.Cleanup(func() {
        if err := client.DeleteResource(context.Background(), prereq.ID); err != nil {
            t.Logf("cleanup failed: %v", err)
        }
    })

    // Run test with prerequisite...
}
```

### Parallel Test Isolation

```go
// ✅ Each parallel test uses unique resources
func TestAccResource_parallel(t *testing.T) {
    t.Parallel()

    // Use unique prefix for this test run
    prefix := fmt.Sprintf("tf-acc-%s", acctest.RandString(8))

    t.Run("create", func(t *testing.T) {
        t.Parallel()
        rName := prefix + "-create"
        // Test with rName...
    })

    t.Run("update", func(t *testing.T) {
        t.Parallel()
        rName := prefix + "-update"
        // Test with rName...
    })
}
```

### Race Condition Detection

```bash
# Run tests with race detector
go test -race -v ./internal/provider/...

# Run specific test with race detection
go test -race -v -run TestAccResource_basic ./internal/provider/...
```

### Code Coverage for Tests

```bash
# Generate coverage report
go test -v -coverprofile=coverage.out ./internal/provider/...

# View coverage in browser
go tool cover -html=coverage.out

# Get coverage percentage
go tool cover -func=coverage.out | grep total

# Coverage for acceptance tests
TF_ACC=1 go test -v -coverprofile=coverage.out ./internal/provider/...
```
