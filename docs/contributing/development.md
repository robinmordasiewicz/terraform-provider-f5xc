# Development Guide

Thank you for your interest in contributing to the F5 Distributed Cloud Terraform Provider!

## Prerequisites

- Go 1.21 or later
- Terraform 1.0 or later
- F5 Distributed Cloud account with API access

## Repository Structure

```
.
├── internal/
│   ├── provider/       # Resource and data source implementations
│   ├── client/         # F5XC API client
│   ├── functions/      # Provider-defined functions
│   └── blindfold/      # Secret Management library
├── tools/              # Code generation utilities
├── docs/               # Generated Terraform Registry documentation
├── examples/           # Example Terraform configurations
└── .github/workflows/  # CI/CD automation
```

## Getting Started

1. Clone the repository:

```bash
git clone https://github.com/robinmordasiewicz/terraform-provider-f5xc.git
cd terraform-provider-f5xc
```

2. Build the provider:

```bash
go build -o terraform-provider-f5xc
```

3. Run tests:

```bash
go test ./...
```

## Development Workflow

See [CLAUDE.md](https://github.com/robinmordasiewicz/terraform-provider-f5xc/blob/main/CLAUDE.md) for detailed workflow guidelines including:

- Issue-first development
- Branch naming conventions
- Conventional commits
- CI/CD automation rules

## Adding New Resources

Resources are generated from F5's OpenAPI specifications using tools in the `tools/` directory.

See the [CLAUDE.md](https://github.com/robinmordasiewicz/terraform-provider-f5xc/blob/main/CLAUDE.md) for the complete process.

## See Also

- [Testing guide](testing.md)
- [Documentation guide](documentation.md)
- [CLAUDE.md](https://github.com/robinmordasiewicz/terraform-provider-f5xc/blob/main/CLAUDE.md) - Complete development constitution
