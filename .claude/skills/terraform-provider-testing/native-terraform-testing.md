# Native Terraform Testing Framework

Complete guide to `.tftest.hcl` files introduced in Terraform 1.6+ with mocking support from 1.7+.

## File Structure and Discovery

Terraform discovers test files with these extensions:
- `.tftest.hcl` (HCL format - recommended)
- `.tftest.json` (JSON format)

Default search locations:
- Root directory of the module
- `tests/` subdirectory

```
module/
├── main.tf
├── variables.tf
├── outputs.tf
└── tests/
    ├── unit.tftest.hcl
    ├── integration.tftest.hcl
    └── mocks/
        └── aws.tfmock.hcl
```

## Test File Syntax

### Root-Level Blocks

```hcl
# Optional: Configure test execution
test {
  parallel = true  # Run eligible run blocks in parallel
}

# Variables available to all run blocks
variables {
  environment = "test"
  region      = "us-west-2"
}

# Provider configuration
provider "aws" {
  region = "us-west-2"
}

# One or more run blocks
run "test_name" {
  # Test configuration
}
```

## Run Block Configuration

### Basic Run Block

```hcl
run "validate_bucket_name" {
  command = plan  # or apply (default)

  variables {
    bucket_name = "my-test-bucket"
  }

  assert {
    condition     = aws_s3_bucket.main.bucket == "my-test-bucket"
    error_message = "Bucket name does not match expected value"
  }

  assert {
    condition     = aws_s3_bucket.main.tags["Environment"] == "test"
    error_message = "Environment tag not set correctly"
  }
}
```

### Run Block Attributes

| Attribute | Description | Default |
|-----------|-------------|---------|
| `command` | `plan` or `apply` | `apply` |
| `parallel` | Override test-level parallel setting | inherited |
| `variables` | Variables for this run only | - |
| `providers` | Provider mappings | - |
| `module` | Alternate module to execute | main config |

## Assertions

### Basic Assertions

```hcl
run "validate_configuration" {
  command = plan

  assert {
    condition     = var.instance_type != ""
    error_message = "Instance type cannot be empty"
  }

  assert {
    condition     = length(var.allowed_cidrs) > 0
    error_message = "At least one CIDR must be specified"
  }

  assert {
    condition     = can(regex("^[a-z][a-z0-9-]*$", var.name))
    error_message = "Name must start with letter and contain only lowercase alphanumeric and hyphens"
  }
}
```

### Assertions with Computed Values

```hcl
run "validate_computed_values" {
  command = apply  # Required for computed values

  assert {
    condition     = aws_instance.main.public_ip != ""
    error_message = "Public IP should be assigned"
  }

  assert {
    condition     = length(aws_instance.main.id) > 0
    error_message = "Instance ID should be generated"
  }
}
```

### Assertions with Functions

```hcl
run "validate_json_output" {
  command = plan

  assert {
    condition     = try(jsondecode(local_file.config.content), null) != null
    error_message = "File content is not valid JSON"
  }

  assert {
    condition     = contains(keys(jsondecode(local_file.config.content)), "version")
    error_message = "JSON must contain 'version' key"
  }
}
```

## Mock Providers (Terraform 1.7+)

### Basic Mock Provider

```hcl
# Mock the AWS provider - no credentials needed
mock_provider "aws" {}

run "test_without_credentials" {
  command = plan

  assert {
    condition     = aws_s3_bucket.main.bucket == var.bucket_name
    error_message = "Bucket name mismatch"
  }
}
```

### Mock Provider with Alias

```hcl
provider "aws" {}  # Real provider

mock_provider "aws" {
  alias = "fake"
}

run "use_real_provider" {
  providers = {
    aws = aws
  }
  # Uses real AWS provider
}

run "use_mocked_provider" {
  providers = {
    aws = aws.fake
  }
  # Uses mocked provider
}
```

### Mock Resources with Defaults

```hcl
mock_provider "aws" {
  mock_resource "aws_s3_bucket" {
    defaults = {
      arn                 = "arn:aws:s3:::mock-bucket"
      bucket              = "mock-bucket"
      bucket_domain_name  = "mock-bucket.s3.amazonaws.com"
      hosted_zone_id      = "Z3AQBSTGFYJSTF"
      region              = "us-east-1"
    }
  }

  mock_resource "aws_instance" {
    defaults = {
      id               = "i-mock12345678"
      public_ip        = "192.168.1.1"
      private_ip       = "10.0.0.1"
      availability_zone = "us-east-1a"
    }
  }
}
```

### Mock Data Sources

```hcl
mock_provider "aws" {
  mock_data "aws_caller_identity" {
    defaults = {
      account_id = "123456789012"
      arn        = "arn:aws:iam::123456789012:root"
      user_id    = "AIDAMOCKUSER"
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
      names = ["us-west-2a", "us-west-2b", "us-west-2c"]
      zone_ids = ["usw2-az1", "usw2-az2", "usw2-az3"]
    }
  }
}
```

## Override Blocks

### Override Specific Resources

```hcl
mock_provider "aws" {}

# Override at file level (applies to all run blocks)
override_resource {
  target = aws_s3_bucket.specific_bucket
  values = {
    bucket = "overridden-bucket-name"
    arn    = "arn:aws:s3:::overridden-bucket-name"
  }
}

run "test_with_override" {
  command = plan

  assert {
    condition     = aws_s3_bucket.specific_bucket.bucket == "overridden-bucket-name"
    error_message = "Override not applied"
  }
}
```

### Override Data Sources

```hcl
override_data {
  target = data.aws_ami.latest
  values = {
    id           = "ami-mock12345"
    name         = "mock-ami"
    architecture = "x86_64"
    owner_id     = "123456789012"
  }
}
```

### Override Modules

```hcl
override_module {
  target = module.network
  outputs = {
    vpc_id            = "vpc-mock12345"
    private_subnet_ids = ["subnet-priv-a", "subnet-priv-b"]
    public_subnet_ids  = ["subnet-pub-a", "subnet-pub-b"]
    nat_gateway_ids    = ["nat-mock123"]
  }
}

run "test_with_module_override" {
  command = plan

  assert {
    condition     = aws_instance.main.subnet_id == "subnet-priv-a"
    error_message = "Instance should use first private subnet"
  }
}
```

### Run Block Level Overrides

```hcl
# File-level override
override_data {
  target = data.aws_region.current
  values = { name = "us-east-1" }
}

run "test_east_region" {
  # Uses file-level override (us-east-1)
  assert {
    condition     = data.aws_region.current.name == "us-east-1"
    error_message = "Should be us-east-1"
  }
}

run "test_west_region" {
  # Run-level override takes precedence
  override_data {
    target = data.aws_region.current
    values = { name = "us-west-2" }
  }

  assert {
    condition     = data.aws_region.current.name == "us-west-2"
    error_message = "Should be us-west-2"
  }
}
```

## Shared Mock Files (.tfmock.hcl)

Create reusable mock definitions in separate files:

```hcl
# tests/mocks/aws.tfmock.hcl
mock_resource "aws_s3_bucket" {
  defaults = {
    arn    = "arn:aws:s3:::shared-mock"
    bucket = "shared-mock"
  }
}

mock_data "aws_caller_identity" {
  defaults = {
    account_id = "123456789012"
  }
}
```

Reference in test files:

```hcl
mock_provider "aws" {
  source = "./mocks/aws"
}

run "test_with_shared_mocks" {
  command = plan
  # Uses mocks from aws.tfmock.hcl
}
```

## Testing Modules

### Setup Module Pattern

```hcl
# Create prerequisite infrastructure
run "setup" {
  command = apply

  module {
    source = "./testing/setup"
  }

  variables {
    resource_group_name = "test-rg"
    location            = "eastus"
  }
}

# Test main configuration using setup outputs
run "test_main_config" {
  command = apply

  variables {
    resource_group_name = run.setup.resource_group_name
    location            = run.setup.location
  }

  assert {
    condition     = azurerm_storage_account.main.resource_group_name == run.setup.resource_group_name
    error_message = "Wrong resource group"
  }
}
```

### Helper Module for Data Lookups

```hcl
# testing/loader/main.tf
variable "url" { type = string }

data "http" "response" {
  url = var.url
}

output "status_code" {
  value = data.http.response.status_code
}
```

```hcl
# tests/integration.tftest.hcl
run "deploy" {
  command = apply
}

run "verify_endpoint" {
  command = apply

  module {
    source = "./testing/loader"
  }

  variables {
    url = run.deploy.endpoint_url
  }

  assert {
    condition     = run.verify_endpoint.status_code == 200
    error_message = "Endpoint should return 200"
  }
}
```

## Variable Precedence

Variables are evaluated in this order (later overrides earlier):
1. `terraform.tfvars` / `*.auto.tfvars`
2. Test file root-level `variables` block
3. Run block `variables` block

```hcl
# Root-level variables (apply to all runs)
variables {
  environment = "test"
  region      = "us-east-1"
}

run "uses_root_level" {
  command = plan
  # Uses environment = "test", region = "us-east-1"
}

run "overrides_region" {
  command = plan

  variables {
    region = "us-west-2"
  }
  # Uses environment = "test", region = "us-west-2"
}
```

## Parallel Execution (Terraform 1.7+)

```hcl
test {
  parallel = true  # Enable parallel execution for all eligible runs
}

run "independent_test_1" {
  command = plan
  # Runs in parallel with other independent runs
}

run "independent_test_2" {
  command = plan
  # Runs in parallel with independent_test_1
}

run "depends_on_previous" {
  command = plan
  parallel = false  # Force sequential execution

  variables {
    value = run.independent_test_1.output_value
  }
  # Runs after independent_test_1 completes
}
```

## Expecting Errors

```hcl
run "expect_validation_error" {
  command = plan

  variables {
    instance_type = "invalid-type"
  }

  expect_failures = [
    var.instance_type,  # Expect validation error on this variable
  ]
}

run "expect_resource_error" {
  command = plan

  expect_failures = [
    aws_instance.main,  # Expect precondition failure on this resource
  ]
}
```

## State Management

### Shared State Between Runs

```hcl
run "create_resource" {
  command = apply
  # Creates resources, generates state
}

run "verify_resource" {
  command = plan
  # Uses state from previous run

  assert {
    condition     = aws_s3_bucket.main.id != ""
    error_message = "Bucket should exist from previous run"
  }
}
```

### Explicit State Key

```hcl
run "setup" {
  command   = apply
  state_key = "shared"

  module {
    source = "./testing/setup"
  }
}

run "main" {
  command   = apply
  state_key = "shared"  # Share state with setup
}
```

## CLI Commands

```bash
# Run all tests
terraform test

# Filter to specific test file
terraform test -filter=tests/unit.tftest.hcl

# Verbose output (shows plan/state)
terraform test -verbose

# Custom test directory
terraform test -test-directory=my_tests

# JSON output for CI/CD
terraform test -json

# Pass variables
terraform test -var="environment=staging"
terraform test -var-file=test.tfvars

# Cloud execution (HCP Terraform)
terraform test -cloud-run=app.terraform.io/org/registry/module/provider
```

## Best Practices

1. **Use `command = plan`** for unit tests - faster, no infrastructure cost
2. **Organize mocks** in `.tfmock.hcl` files for reusability
3. **Test variable validation** with `expect_failures`
4. **Use setup modules** for prerequisite infrastructure
5. **Name tests descriptively** - test names appear in output
6. **Parallelize independent tests** with `parallel = true`
7. **Clean assertions** - one logical check per assert block
8. **Reference run outputs** with `run.<name>.<attribute>` syntax
