# Provider Functions

Provider-defined functions offer utility capabilities beyond resource management.

!!! info "Terraform 1.8+ Required"
    Provider functions require Terraform 1.8.0 or later.

## Available Functions

### Secret Management

- **`provider::f5xc::blindfold`** - Encrypt base64-encoded plaintext using F5XC Secret Management
- **`provider::f5xc::blindfold_file`** - Read and encrypt file contents using F5XC Secret Management

## Usage Example

```hcl
terraform {
  required_version = ">= 1.8.0"
  required_providers {
    f5xc = {
      source  = "robinmordasiewicz/f5xc"
      version = "~> 3.0"
    }
  }
}

locals {
  encrypted = provider::f5xc::blindfold(
    base64encode("my-secret-value"),
    "secret-policy-name",
    "namespace"
  )
}
```

## Function Documentation

For complete function documentation, see the [Terraform Registry](https://registry.terraform.io/providers/robinmordasiewicz/f5xc/latest/docs/functions).
