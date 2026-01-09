# Installation

## Terraform Registry

The F5 Distributed Cloud provider is available on the [Terraform Registry](https://registry.terraform.io/providers/robinmordasiewicz/f5xc/latest).

Add the provider to your Terraform configuration:

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
```

Run `terraform init` to download and install the provider:

```bash
terraform init
```

## Version Constraints

We recommend using version constraints to ensure compatibility:

- `~> 3.0` - Any version in the 3.x series
- `>= 3.0.0` - Version 3.0.0 or higher
- `= 3.0.0` - Exactly version 3.0.0

## Next Steps

- [Configure authentication](authentication.md)
- [Quick start guide](quick-start.md)
