# Quick Start

This guide walks you through creating your first F5XC resource with Terraform.

## Prerequisites

- [Terraform](https://www.terraform.io/downloads.html) >= 1.0
- F5 Distributed Cloud account
- [API token](authentication.md)

## Create a Namespace

```hcl
terraform {
  required_providers {
    f5xc = {
      source  = "robinmordasiewicz/f5xc"
      version = "~> 3.0"
    }
  }
}

provider "f5xc" {
  api_url = "https://console.ves.volterra.io"
}

resource "f5xc_namespace" "example" {
  name = "my-terraform-namespace"
}
```

## Apply the Configuration

1. Initialize Terraform:

```bash
terraform init
```

2. Review the plan:

```bash
terraform plan
```

3. Apply the configuration:

```bash
terraform apply
```

## Import Existing Resources

You can import existing F5XC resources:

```bash
terraform import f5xc_namespace.example my-existing-namespace
```

## Next Steps

- Explore [guides](../guides/index.md) for more complex scenarios
- Browse available [resources](../resources/index.md)
- Check out [data sources](../data-sources/index.md)
