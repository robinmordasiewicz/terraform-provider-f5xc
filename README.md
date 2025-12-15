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
export F5XC_API_TOKEN="your-api-token"
```

## Documentation

- [Provider Documentation](https://registry.terraform.io/providers/robinmordasiewicz/f5xc/latest/docs)
- [F5 Distributed Cloud API](https://docs.cloud.f5.com/)

## Contributing

This provider uses automated workflows for documentation generation and releases:

- **Documentation**: Auto-generated from OpenAPI specs on merge to main
- **Releases**: Triggered automatically using [semantic versioning](https://semver.org/) based on conventional commits
- **Testing**: All PRs require passing tests and lint checks

All substantive changes (code, documentation, configuration) automatically trigger a new release with appropriate version bumping.

For detailed contribution guidelines, see [CLAUDE.md](CLAUDE.md).

## License

[Mozilla Public License 2.0](LICENSE)
