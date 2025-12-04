# Terraform Test Harness Templates

Copy-paste templates for quick test setup.

## Native Terraform Test Templates

### Basic Unit Test Template

```hcl
# tests/unit.tftest.hcl

variables {
  name        = "test-resource"
  environment = "test"
}

run "validate_basic_config" {
  command = plan

  assert {
    condition     = var.name != ""
    error_message = "Name cannot be empty"
  }

  assert {
    condition     = contains(["dev", "test", "prod"], var.environment)
    error_message = "Environment must be dev, test, or prod"
  }
}

run "validate_resource_naming" {
  command = plan

  assert {
    condition     = local.resource_name == "test-resource-test"
    error_message = "Resource name format incorrect"
  }
}
```

### Integration Test with Mocks Template

```hcl
# tests/integration.tftest.hcl

mock_provider "aws" {
  mock_resource "aws_s3_bucket" {
    defaults = {
      arn                 = "arn:aws:s3:::mock-bucket"
      bucket              = "mock-bucket"
      bucket_domain_name  = "mock-bucket.s3.amazonaws.com"
    }
  }

  mock_data "aws_caller_identity" {
    defaults = {
      account_id = "123456789012"
      arn        = "arn:aws:iam::123456789012:user/test"
    }
  }

  mock_data "aws_region" {
    defaults = {
      name = "us-west-2"
    }
  }
}

variables {
  bucket_name = "test-bucket"
  environment = "test"
}

run "create_s3_bucket" {
  command = plan

  assert {
    condition     = aws_s3_bucket.main.bucket == var.bucket_name
    error_message = "Bucket name mismatch"
  }

  assert {
    condition     = aws_s3_bucket.main.tags["Environment"] == var.environment
    error_message = "Environment tag not set"
  }
}

run "verify_bucket_policy" {
  command = plan

  assert {
    condition     = can(jsondecode(aws_s3_bucket_policy.main.policy))
    error_message = "Bucket policy is not valid JSON"
  }
}
```

### Module Override Template

```hcl
# tests/with_overrides.tftest.hcl

mock_provider "aws" {}

override_module {
  target = module.vpc
  outputs = {
    vpc_id             = "vpc-mock12345"
    private_subnet_ids = ["subnet-priv-a", "subnet-priv-b"]
    public_subnet_ids  = ["subnet-pub-a", "subnet-pub-b"]
  }
}

override_data {
  target = data.aws_availability_zones.available
  values = {
    names = ["us-west-2a", "us-west-2b", "us-west-2c"]
  }
}

run "test_with_vpc_dependency" {
  command = plan

  assert {
    condition     = aws_instance.main.subnet_id == "subnet-priv-a"
    error_message = "Instance should be in first private subnet"
  }
}
```

### Shared Mock File Template

```hcl
# tests/mocks/aws.tfmock.hcl

mock_resource "aws_s3_bucket" {
  defaults = {
    arn                 = "arn:aws:s3:::mock-bucket"
    bucket              = "mock-bucket"
    bucket_domain_name  = "mock-bucket.s3.amazonaws.com"
    hosted_zone_id      = "Z3AQBSTGFYJSTF"
    region              = "us-west-2"
  }
}

mock_resource "aws_instance" {
  defaults = {
    id                = "i-mock12345678"
    arn               = "arn:aws:ec2:us-west-2:123456789012:instance/i-mock12345678"
    public_ip         = "54.123.45.67"
    private_ip        = "10.0.1.100"
    availability_zone = "us-west-2a"
    instance_state    = "running"
  }
}

mock_resource "aws_iam_role" {
  defaults = {
    arn         = "arn:aws:iam::123456789012:role/mock-role"
    id          = "mock-role"
    unique_id   = "AROAMOCKROLEID"
    create_date = "2024-01-01T00:00:00Z"
  }
}

mock_data "aws_caller_identity" {
  defaults = {
    account_id = "123456789012"
    arn        = "arn:aws:iam::123456789012:user/test"
    user_id    = "AIDAMOCKUSERID"
  }
}

mock_data "aws_region" {
  defaults = {
    name        = "us-west-2"
    description = "US West (Oregon)"
  }
}

mock_data "aws_availability_zones" {
  defaults = {
    names    = ["us-west-2a", "us-west-2b", "us-west-2c"]
    zone_ids = ["usw2-az1", "usw2-az2", "usw2-az3"]
  }
}
```

## Go SDK Test Templates

### Provider Test Setup Template

```go
// internal/provider/provider_test.go
package provider_test

import (
    "os"
    "testing"

    "github.com/hashicorp/terraform-plugin-framework/providerserver"
    "github.com/hashicorp/terraform-plugin-go/tfprotov6"
    "github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// testAccProtoV6ProviderFactories for Protocol Version 6
var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
    "example": providerserver.NewProtocol6WithError(New("test")()),
}

func testAccPreCheck(t *testing.T) {
    t.Helper()

    if v := os.Getenv("EXAMPLE_API_KEY"); v == "" {
        t.Fatal("EXAMPLE_API_KEY must be set for acceptance tests")
    }
    if v := os.Getenv("EXAMPLE_API_ENDPOINT"); v == "" {
        t.Fatal("EXAMPLE_API_ENDPOINT must be set for acceptance tests")
    }
}

func TestProvider(t *testing.T) {
    t.Parallel()

    resource.Test(t, resource.TestCase{
        ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
        Steps: []resource.TestStep{
            {
                Config: testAccProviderConfig,
            },
        },
    })
}

const testAccProviderConfig = `
provider "example" {}
`
```

### Resource Test Template

```go
// internal/service/example/resource_thing_test.go
package example_test

import (
    "fmt"
    "regexp"
    "testing"

    "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
    "github.com/hashicorp/terraform-plugin-testing/helper/resource"
    "github.com/hashicorp/terraform-plugin-testing/knownvalue"
    "github.com/hashicorp/terraform-plugin-testing/statecheck"
    "github.com/hashicorp/terraform-plugin-testing/terraform"
    "github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccExampleThing_basic(t *testing.T) {
    t.Parallel()

    rName := acctest.RandomWithPrefix("tf-acc-test")

    resource.ParallelTest(t, resource.TestCase{
        PreCheck:                 func() { testAccPreCheck(t) },
        ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
        CheckDestroy:             testAccCheckExampleThingDestroy,
        Steps: []resource.TestStep{
            {
                Config: testAccExampleThingConfig_basic(rName),
                Check: resource.ComposeAggregateTestCheckFunc(
                    testAccCheckExampleThingExists("example_thing.test"),
                    resource.TestCheckResourceAttr("example_thing.test", "name", rName),
                    resource.TestCheckResourceAttrSet("example_thing.test", "id"),
                ),
                ConfigStateChecks: []statecheck.StateCheck{
                    statecheck.ExpectKnownValue(
                        "example_thing.test",
                        tfjsonpath.New("name"),
                        knownvalue.StringExact(rName),
                    ),
                    statecheck.ExpectKnownValue(
                        "example_thing.test",
                        tfjsonpath.New("enabled"),
                        knownvalue.Bool(true),
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

func TestAccExampleThing_update(t *testing.T) {
    t.Parallel()

    rName := acctest.RandomWithPrefix("tf-acc-test")
    rNameUpdated := rName + "-updated"

    resource.ParallelTest(t, resource.TestCase{
        PreCheck:                 func() { testAccPreCheck(t) },
        ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
        CheckDestroy:             testAccCheckExampleThingDestroy,
        Steps: []resource.TestStep{
            {
                Config: testAccExampleThingConfig_basic(rName),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckExampleThingExists("example_thing.test"),
                    resource.TestCheckResourceAttr("example_thing.test", "name", rName),
                ),
            },
            {
                Config: testAccExampleThingConfig_basic(rNameUpdated),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckExampleThingExists("example_thing.test"),
                    resource.TestCheckResourceAttr("example_thing.test", "name", rNameUpdated),
                ),
            },
        },
    })
}

func TestAccExampleThing_full(t *testing.T) {
    t.Parallel()

    rName := acctest.RandomWithPrefix("tf-acc-test")

    resource.ParallelTest(t, resource.TestCase{
        PreCheck:                 func() { testAccPreCheck(t) },
        ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
        CheckDestroy:             testAccCheckExampleThingDestroy,
        Steps: []resource.TestStep{
            {
                Config: testAccExampleThingConfig_full(rName),
                Check: resource.ComposeAggregateTestCheckFunc(
                    testAccCheckExampleThingExists("example_thing.test"),
                    resource.TestCheckResourceAttr("example_thing.test", "name", rName),
                    resource.TestCheckResourceAttr("example_thing.test", "description", "Test description"),
                    resource.TestCheckResourceAttr("example_thing.test", "enabled", "true"),
                    resource.TestCheckResourceAttr("example_thing.test", "tags.%", "2"),
                    resource.TestCheckResourceAttr("example_thing.test", "tags.Environment", "test"),
                ),
            },
        },
    })
}

func TestAccExampleThing_disappears(t *testing.T) {
    t.Parallel()

    rName := acctest.RandomWithPrefix("tf-acc-test")

    resource.ParallelTest(t, resource.TestCase{
        PreCheck:                 func() { testAccPreCheck(t) },
        ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
        CheckDestroy:             testAccCheckExampleThingDestroy,
        Steps: []resource.TestStep{
            {
                Config: testAccExampleThingConfig_basic(rName),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckExampleThingExists("example_thing.test"),
                    acctest.CheckResourceDisappears(testAccProvider, resourceExampleThing(), "example_thing.test"),
                ),
                ExpectNonEmptyPlan: true,
            },
        },
    })
}

func TestAccExampleThing_invalidName(t *testing.T) {
    t.Parallel()

    resource.ParallelTest(t, resource.TestCase{
        PreCheck:                 func() { testAccPreCheck(t) },
        ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
        Steps: []resource.TestStep{
            {
                Config:      testAccExampleThingConfig_basic("invalid name with spaces"),
                ExpectError: regexp.MustCompile(`name must contain only alphanumeric characters and hyphens`),
            },
        },
    })
}

// Helper functions

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

        if !isResourceNotFoundError(err) {
            return fmt.Errorf("unexpected error: %w", err)
        }
    }

    return nil
}

// Config generators

func testAccExampleThingConfig_basic(rName string) string {
    return fmt.Sprintf(`
resource "example_thing" "test" {
  name = %[1]q
}
`, rName)
}

func testAccExampleThingConfig_full(rName string) string {
    return fmt.Sprintf(`
resource "example_thing" "test" {
  name        = %[1]q
  description = "Test description"
  enabled     = true

  tags = {
    Environment = "test"
    Name        = %[1]q
  }
}
`, rName)
}
```

### Data Source Test Template

```go
// internal/service/example/data_source_thing_test.go
package example_test

import (
    "fmt"
    "testing"

    "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
    "github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccExampleThingDataSource_basic(t *testing.T) {
    t.Parallel()

    rName := acctest.RandomWithPrefix("tf-acc-test")

    resource.ParallelTest(t, resource.TestCase{
        PreCheck:                 func() { testAccPreCheck(t) },
        ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
        Steps: []resource.TestStep{
            {
                Config: testAccExampleThingDataSourceConfig_basic(rName),
                Check: resource.ComposeAggregateTestCheckFunc(
                    resource.TestCheckResourceAttr("data.example_thing.test", "name", rName),
                    resource.TestCheckResourceAttrSet("data.example_thing.test", "id"),
                    resource.TestCheckResourceAttrPair(
                        "data.example_thing.test", "id",
                        "example_thing.test", "id",
                    ),
                ),
            },
        },
    })
}

func TestAccExampleThingDataSource_byID(t *testing.T) {
    t.Parallel()

    rName := acctest.RandomWithPrefix("tf-acc-test")

    resource.ParallelTest(t, resource.TestCase{
        PreCheck:                 func() { testAccPreCheck(t) },
        ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
        Steps: []resource.TestStep{
            {
                Config: testAccExampleThingDataSourceConfig_byID(rName),
                Check: resource.ComposeAggregateTestCheckFunc(
                    resource.TestCheckResourceAttr("data.example_thing.by_id", "name", rName),
                    resource.TestCheckResourceAttrPair(
                        "data.example_thing.by_id", "id",
                        "example_thing.test", "id",
                    ),
                ),
            },
        },
    })
}

func testAccExampleThingDataSourceConfig_basic(rName string) string {
    return fmt.Sprintf(`
resource "example_thing" "test" {
  name = %[1]q
}

data "example_thing" "test" {
  name = example_thing.test.name
}
`, rName)
}

func testAccExampleThingDataSourceConfig_byID(rName string) string {
    return fmt.Sprintf(`
resource "example_thing" "test" {
  name = %[1]q
}

data "example_thing" "by_id" {
  id = example_thing.test.id
}
`, rName)
}
```

### Test Sweeper Template

```go
// internal/service/example/sweep.go
package example

import (
    "fmt"
    "log"
    "strings"

    "github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func init() {
    resource.AddTestSweepers("example_thing", &resource.Sweeper{
        Name: "example_thing",
        F:    sweepThings,
        Dependencies: []string{
            "example_thing_association", // Sweep associations first
        },
    })
}

func sweepThings(region string) error {
    client, err := sharedClientForRegion(region)
    if err != nil {
        return fmt.Errorf("error getting client: %w", err)
    }

    conn := client.(*Client)

    input := &ListThingsInput{
        MaxResults: 100,
    }

    var sweeperErrs []error

    for {
        output, err := conn.ListThings(input)
        if err != nil {
            return fmt.Errorf("error listing things: %w", err)
        }

        for _, thing := range output.Things {
            // Only sweep test resources
            if !strings.HasPrefix(thing.Name, "tf-acc-test") {
                continue
            }

            log.Printf("[INFO] Deleting Example Thing: %s", thing.ID)

            _, err := conn.DeleteThing(&DeleteThingInput{
                ID: thing.ID,
            })
            if err != nil {
                sweeperErrs = append(sweeperErrs, fmt.Errorf(
                    "error deleting thing %s: %w", thing.ID, err,
                ))
                continue
            }
        }

        if output.NextToken == "" {
            break
        }
        input.NextToken = output.NextToken
    }

    if len(sweeperErrs) > 0 {
        return fmt.Errorf("sweeper errors: %v", sweeperErrs)
    }

    return nil
}
```

### Custom State Check Template

```go
// internal/service/example/state_checks_test.go
package example_test

import (
    "context"
    "fmt"

    "github.com/hashicorp/terraform-plugin-testing/statecheck"
)

type expectRemoteResourceExists struct {
    resourceAddress string
    apiClient       APIClient
}

var _ statecheck.StateCheck = expectRemoteResourceExists{}

func (e expectRemoteResourceExists) CheckState(ctx context.Context, req statecheck.CheckStateRequest, resp *statecheck.CheckStateResponse) {
    // Find the resource in state
    var resourceID string
    for _, resource := range req.State.Values.RootModule.Resources {
        if resource.Address == e.resourceAddress {
            id, ok := resource.AttributeValues["id"].(string)
            if !ok {
                resp.Error = fmt.Errorf("id attribute not found or not a string")
                return
            }
            resourceID = id
            break
        }
    }

    if resourceID == "" {
        resp.Error = fmt.Errorf("resource %s not found in state", e.resourceAddress)
        return
    }

    // Verify resource exists in remote API
    _, err := e.apiClient.GetThing(ctx, resourceID)
    if err != nil {
        resp.Error = fmt.Errorf("resource %s not found in remote API: %w", resourceID, err)
        return
    }
}

func ExpectRemoteResourceExists(resourceAddress string, client APIClient) statecheck.StateCheck {
    return expectRemoteResourceExists{
        resourceAddress: resourceAddress,
        apiClient:       client,
    }
}

// Usage:
// ConfigStateChecks: []statecheck.StateCheck{
//     ExpectRemoteResourceExists("example_thing.test", testAccClient),
// },
```

## Makefile Template

```makefile
# Makefile for Terraform Provider

GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)
PKG_NAME=internal

default: build

build:
	go build -o terraform-provider-example

test:
	go test -v ./...

testacc:
	TF_ACC=1 go test -v ./$(PKG_NAME)/... -timeout 120m

testacc-short:
	TF_ACC=1 go test -v -short ./$(PKG_NAME)/... -timeout 30m

sweep:
	@echo "WARNING: This will destroy infrastructure. Use with caution."
	go test ./$(PKG_NAME)/... -sweep=us-west-2 -v

fmt:
	gofmt -w $(GOFMT_FILES)

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

lint:
	golangci-lint run ./...

.PHONY: build test testacc testacc-short sweep fmt fmtcheck lint
```
