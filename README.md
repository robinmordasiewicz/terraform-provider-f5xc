# Terraform Provider for F5 Distributed Cloud

Community Terraform provider for [F5 Distributed Cloud](https://www.f5.com/cloud).

## Installation

```terraform
terraform {
  required_providers {
    f5xc = {
      source  = "robinmordasiewicz/f5xc"
    }
  }
}

provider "f5xc" {}
```

## Authentication

Set your API token as an environment variable:

```bash
export VES_API_TOKEN="your-api-token"
```

## Documentation

- [Provider Documentation](https://registry.terraform.io/providers/robinmordasiewicz/f5xc/latest/docs)
- [F5 Distributed Cloud API](https://docs.cloud.f5.com/)

## License

[Mozilla Public License 2.0](LICENSE)
