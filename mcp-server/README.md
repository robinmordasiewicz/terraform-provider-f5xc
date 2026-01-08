# F5 Distributed Cloud Terraform MCP Server

[![npm version](https://img.shields.io/npm/v/@robinmordasiewicz/f5xc-terraform-mcp)](https://www.npmjs.com/package/@robinmordasiewicz/f5xc-terraform-mcp)
[![MIT License](https://img.shields.io/badge/License-MIT-yellow)](LICENSE)

Token-optimized Model Context Protocol (MCP) server for F5 Distributed Cloud Terraform provider documentation, API specifications, and subscription management.

This MCP server exposes tools for AI assistants to interact with the [F5 Distributed Cloud Terraform Provider](https://github.com/robinmordasiewicz/terraform-provider-f5xc), enabling intelligent infrastructure-as-code assistance.

## Features

- **270+ OpenAPI Specifications** - Complete F5XC API documentation for AI-assisted configuration
- **250+ Terraform Resources** - Full documentation for all provider resources and data sources
- **Token-Optimized Design** - 14 tools consolidated to 7 (~75% token reduction)
- **Subscription Tier Management** - Check feature availability across STANDARD/ADVANCED/PREMIUM tiers
- **Addon Service Workflows** - Manage F5XC service activations
- **Resource Metadata** - Deterministic AI configuration generation with validation rules, defaults, and one-of schemas

## Installation

### Global Installation

```bash
npm install -g @robinmordasiewicz/f5xc-terraform-mcp
```

### Usage with Claude Desktop

Add to your `claude_desktop_config.json`:

```json
{
  "mcpServers": {
    "f5xc-terraform": {
      "command": "npx",
      "args": ["-y", "@robinmordasiewicz/f5xc-terraform-mcp"]
    }
  }
}
```

### Manual CLI Usage

```bash
# Run the MCP server
npx @robinmordasiewicz/f5xc-terraform-mcp

# Or install globally and run directly
f5xc-terraform-mcp
```

## Available Tools

| Tool | Description |
|------|-------------|
| `f5xc_terraform_discover` | Meta-tool for discovering available tools with optional schema details |
| `f5xc_terraform_docs` | Search, get, or list documentation (resources, data-sources, functions, guides) |
| `f5xc_terraform_api` | Query 270+ OpenAPI specs (search, get, find endpoints, schema definitions) |
| `f5xc_terraform_subscription` | Check subscription tier requirements for resources and properties |
| `f5xc_terraform_addon` | Manage addon services (list, check activation, workflow) |
| `f5xc_terraform_metadata` | Resource metadata for deterministic AI config generation |
| `f5xc_terraform_get_summary` | Get provider summary with documentation and API specs overview |

## Tool Details

### Documentation Tool (`f5xc_terraform_docs`)

Search and retrieve Terraform provider documentation.

```typescript
// Search for resources
await f5xc_terraform_docs({
  operation: "search",
  query: "http_loadbalancer",
  type: "resource",
  limit: 10
});

// Get specific resource documentation
await f5xc_terraform_docs({
  operation: "get",
  name: "http_loadbalancer",
  type: "resource"
});

// List all resources
await f5xc_terraform_docs({
  operation: "list",
  type: "resource"
});
```

### API Specification Tool (`f5xc_terraform_api`)

Query OpenAPI specifications for API details.

```typescript
// Search API specs
await f5xc_terraform_api({
  operation: "search",
  query: "http_loadbalancer",
  limit: 10
});

// Get full spec with endpoints
await f5xc_terraform_api({
  operation: "get",
  name: "http_loadbalancer",
  include_paths: true,
  include_definitions: false
});

// Find endpoints matching pattern
await f5xc_terraform_api({
  operation: "find_endpoints",
  pattern: "/namespaces",
  method: "GET"
});

// Get schema definition
await f5xc_terraform_api({
  operation: "get_definition",
  spec_name: "http_loadbalancer",
  definition_name: "HTTPLoadBalancerSpec"
});
```

### Subscription Tool (`f5xc_terraform_subscription`)

Check subscription tier requirements.

```typescript
// Get resource tier requirements
await f5xc_terraform_subscription({
  operation: "resource",
  resource_name: "http_loadbalancer"
});

// Check if property requires Advanced tier
await f5xc_terraform_subscription({
  operation: "property",
  resource_name: "http_loadbalancer",
  property_path: "enable_malicious_user_detection"
});
```

### Addon Service Tool (`f5xc_terraform_addon`)

Manage F5XC addon services.

```typescript
// List available addon services
await f5xc_terraform_addon({
  operation: "list"
});

// Check service activation status
await f5xc_terraform_addon({
  operation: "check",
  service_name: "wasm"
});

// Get activation workflow guidance
await f5xc_terraform_addon({
  operation: "workflow",
  service_name: "security_app_protect"
});
```

### Metadata Tool (`f5xc_terraform_metadata`)

Get resource metadata for deterministic AI configuration.

```typescript
// Get one-of field options
await f5xc_terraform_metadata({
  operation: "oneof",
  resource: "http_loadbalancer",
  attribute: "http_type"
});

// Get validation patterns
await f5xc_terraform_metadata({
  operation: "validation",
  pattern: "name"
});

// Get default values
await f5xc_terraform_metadata({
  operation: "defaults",
  resource: "namespace"
});

// Get enums
await f5xc_terraform_metadata({
  operation: "enums",
  resource: "http_loadbalancer"
});

// Check attribute properties
await f5xc_terraform_metadata({
  operation: "attribute",
  resource: "http_loadbalancer",
  attribute: "name"
});

// Check requires_replace
await f5xc_terraform_metadata({
  operation: "requires_replace",
  resource: "http_loadbalancer",
  attribute: "name"
});

// Get resource dependencies
await f5xc_terraform_metadata({
  operation: "dependencies",
  resource: "http_loadbalancer"
});

// Get troubleshooting info
await f5xc_terraform_metadata({
  operation: "troubleshoot",
  error_code: "NOT_FOUND"
});

// Get full resource summary
await f5xc_terraform_metadata({
  operation: "summary",
  resource: "namespace"
});
```

## Response Format

All tools support both `markdown` (default) and `json` response formats:

```typescript
await f5xc_terraform_get_summary({
  response_format: "json"
});
```

## Token Optimization

This MCP server uses a token-optimized design:

- **14 original tools consolidated to 7** (~75% token reduction)
- **Discovery meta-tool** enables lazy schema loading
- **Shared parameter descriptions** reduce schema size
- **Response truncation** for large payloads (50,000 character limit)

## F5XC Provider Quick Start

For Terraform configuration:

```hcl
terraform {
  required_providers {
    f5xc = {
      source  = "robinmordasiewicz/f5xc"
      version = ">= 2.0.0"
    }
  }
}

provider "f5xc" {
  api_url   = "https://your-tenant.console.ves.volterra.io"
  api_token = var.f5xc_api_token
}

# Create a namespace
resource "f5xc_namespace" "example" {
  name = "my-namespace"
}
```

## Resources Supported

- **250+ Terraform Resources** - Full CRUD for F5XC services
- **40+ Data Sources** - Read-only access to F5XC configurations
- **2 Provider Functions** - `blindfold` and `blindfold_file` for secret management
- **Multiple Guides** - Step-by-step tutorials for common use cases

## Subscription Tiers

F5XC resources are organized by subscription tier:

| Tier | Description |
|------|-------------|
| **STANDARD** | Core networking and security features |
| **ADVANCED** | Enhanced WAF, bot defense, and advanced security |
| **PREMIUM** | Enterprise features and dedicated support |

Use the subscription tool to check which tier a resource or feature requires.

## Documentation

- [Terraform Provider Documentation](https://registry.terraform.io/providers/robinmordasiewicz/f5xc/latest/docs)
- [F5 Distributed Cloud Docs](https://docs.cloud.f5.com/)
- [MCP Protocol Documentation](https://modelcontextprotocol.io/)

## Version Synchronization

This npm package version is automatically synchronized with GitHub releases. Both the Terraform provider and MCP server share the same version number for consistency.

## Development

```bash
# Clone the repository
git clone https://github.com/robinmordasiewicz/terraform-provider-f5xc.git

# Navigate to MCP server directory
cd mcp-server

# Install dependencies
npm install

# Build the server
npm run build

# Run in development mode with hot reload
npm run dev
```

## License

This project is licensed under the [MIT License](LICENSE).

## Support

- [GitHub Issues](https://github.com/robinmordasiewicz/terraform-provider-f5xc/issues)
- [F5 Community](https://community.f5.com/)
