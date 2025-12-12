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

## AI-Enhanced Schema

For AI agents and tooling, an enhanced schema wrapper is available that adds an `ai_hints` namespace with OneOf constraints, enum values, resource categories, and recommended defaults:

```bash
terraform providers schema -json | go run tools/terraform-schema-ai/main.go -stdin
```

The wrapper enriches the standard Terraform schema with:

- **OneOf groups**: Mutually exclusive fields extracted from OpenAPI specs
- **Recommended defaults**: Best practice default choices for OneOf constraints
- **Resource categories**: Load Balancing, Security, DNS, Networking, etc.
- **Dependencies**: Common resource creation order
- **Enum values**: Valid values for enum fields

The original `provider_schemas` output remains unchanged; AI hints are added in a separate `ai_hints` namespace.

## License

[Mozilla Public License 2.0](LICENSE)
