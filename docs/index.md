---
page_title: "F5XC Provider"
description: |-
  Terraform provider for F5 Distributed Cloud (F5XC) enabling infrastructure as code for load balancers, security policies, sites, and networking. Community-maintained provider built from public F5 API documentation.
---

# F5XC Provider

The F5XC Terraform provider enables infrastructure as code management for F5 Distributed Cloud (F5XC) resources. Configure HTTP/TCP load balancers, origin pools, application firewalls, service policies, cloud sites, and more through declarative Terraform configurations.

This is a community-maintained provider built from public F5 API documentation.

## Requirements

| Name      | Version |
|-----------|---------|
| terraform | >= 1.8  |

-> **Note:** This provider uses provider-defined functions which require Terraform 1.8 or later. For details, see the [Functions](/docs/functions) documentation.

## Authenticating to F5 Distributed Cloud

The F5XC Terraform provider supports multiple authentication methods:

1. **API Token** - Simplest method using a personal API token
2. **P12 Certificate** - Certificate-based authentication using PKCS#12 bundle
3. **PEM Certificate** - Certificate-based authentication using separate cert/key files

Learn more about [how to generate API credentials](https://docs.cloud.f5.com/docs/how-to/user-mgmt/credentials).

## Example Usage

```terraform
# Configure the F5XC Provider with API Token Authentication
provider "f5xc" {
  api_url   = "https://your-tenant.console.ves.volterra.io"
  api_token = var.f5xc_api_token
}

# Alternatively, use environment variables:
# export F5XC_API_URL="https://your-tenant.console.ves.volterra.io"
# export F5XC_API_TOKEN="your-api-token"

variable "f5xc_api_token" {
  description = "F5 Distributed Cloud API Token"
  type        = string
  sensitive   = true
}

# Or use P12 Certificate Authentication:
# provider "f5xc" {
#   api_url      = "https://your-tenant.console.ves.volterra.io"
#   api_p12_file = "/path/to/certificate.p12"
#   p12_password = var.f5xc_p12_password
# }
#
# Environment variables for P12 authentication:
# export F5XC_API_URL="https://your-tenant.console.ves.volterra.io"
# export F5XC_P12_FILE="/path/to/certificate.p12"
# export F5XC_P12_PASSWORD="your-p12-password"
```

## Argument Reference

### Required (one of the following authentication methods)

* `api_token` - F5 Distributed Cloud API Token (`String`, Sensitive). Can also be set via `F5XC_API_TOKEN` environment variable.

* `api_p12_file` - Path to PKCS#12 certificate bundle file (`String`). Can also be set via `F5XC_P12_FILE` environment variable. Requires `p12_password`.

* `api_cert` and `api_key` - Paths to PEM-encoded certificate and private key files (`String`). Can also be set via `F5XC_CERT` and `F5XC_KEY` environment variables.

### Optional

* `api_url` - F5 Distributed Cloud API URL (`String`). Base URL **without** `/api` suffix. Defaults to `https://console.ves.volterra.io`. Can also be set via `F5XC_API_URL` environment variable.

* `p12_password` - Password for PKCS#12 certificate bundle (`String`, Sensitive). Required when using `api_p12_file`. Can also be set via `F5XC_P12_PASSWORD` environment variable.

* `api_ca_cert` - Path to PEM-encoded CA certificate file (`String`). Optional, used for server certificate verification. Can also be set via `F5XC_CACERT` environment variable.

## Authentication Options

### Option 1: API Token Authentication

The simplest authentication method using a personal API token.

**Provider Configuration:**

```hcl
provider "f5xc" {
  api_url   = "https://your-tenant.console.ves.volterra.io"
  api_token = var.f5xc_api_token
}
```

**Environment Variables:**

```bash
export F5XC_API_URL="https://your-tenant.console.ves.volterra.io"
export F5XC_API_TOKEN="your-api-token"
```

### Option 2: P12 Certificate Authentication

Certificate-based authentication using a PKCS#12 bundle downloaded from F5 Distributed Cloud.

**Provider Configuration:**

```hcl
provider "f5xc" {
  api_url      = "https://your-tenant.console.ves.volterra.io"
  api_p12_file = "/path/to/certificate.p12"
  p12_password = var.f5xc_p12_password
}
```

**Environment Variables:**

```bash
export F5XC_API_URL="https://your-tenant.console.ves.volterra.io"
export F5XC_P12_FILE="/path/to/certificate.p12"
export F5XC_P12_PASSWORD="your-p12-password"
```

### Option 3: PEM Certificate Authentication

Certificate-based authentication using separate PEM-encoded certificate and key files.

**Provider Configuration:**

```hcl
provider "f5xc" {
  api_url     = "https://your-tenant.console.ves.volterra.io"
  api_cert    = "/path/to/certificate.crt"
  api_key     = "/path/to/private.key"
  api_ca_cert = "/path/to/ca-certificate.crt"  # Optional
}
```

**Environment Variables:**

```bash
export F5XC_API_URL="https://your-tenant.console.ves.volterra.io"
export F5XC_CERT="/path/to/certificate.crt"
export F5XC_KEY="/path/to/private.key"
export F5XC_CACERT="/path/to/ca-certificate.crt"  # Optional
```

-> **Note:** Environment variables are the recommended approach for CI/CD pipelines and to avoid storing sensitive credentials in version control.

## Getting Started

1. **Generate API Credentials**: Navigate to your F5 Distributed Cloud console, go to **Administration** > **Personal Management** > **Credentials**, and create either an API Token or download a certificate bundle.

2. **Configure the Provider**: Add the provider configuration to your Terraform files using one of the authentication options above.

3. **Create Resources**: Start managing F5XC resources like namespaces, load balancers, and origin pools.

### Example: Create a Namespace

```hcl
resource "f5xc_namespace" "example" {
  name = "example-namespace"
}
```

### Example: Create an HTTP Load Balancer

```hcl
resource "f5xc_http_loadbalancer" "example" {
  name      = "example-load-balancer"
  namespace = "example-namespace"
  domains   = ["example.com"]
}
```

## MCP Configuration

This provider includes a Model Context Protocol (MCP) server that enables AI assistants like Claude to interact with F5 Distributed Cloud resources. The MCP server provides schema information, documentation, and example configurations.

### Quick Setup for Claude Desktop

Add the following to your Claude Desktop configuration file:

**macOS:** `~/Library/Application Support/Claude/claude_desktop_config.json`
**Windows:** `%APPDATA%\Claude\claude_desktop_config.json`

```json
{
  "mcpServers": {
    "f5xc-terraform": {
      "command": "npx",
      "args": ["@robinmordasiewicz/f5xc-terraform-mcp"]
    }
  }
}
```

### Quick Setup for Claude Code (CLI)

Install the MCP server with a single command:

```bash
claude mcp add f5xc-terraform -- npx -y @robinmordasiewicz/f5xc-terraform-mcp
```

**Scope Options:**

* `--scope local` (default) - Available only in the current directory
* `--scope user` - Available in all your Claude Code sessions
* `--scope project` - Shared with anyone who clones the repository (saved in `.mcp.json`)

Example with user scope (recommended for personal use):

```bash
claude mcp add --scope user f5xc-terraform -- npx -y @robinmordasiewicz/f5xc-terraform-mcp
```

**Verify Installation:**

```bash
claude mcp list
```

You should see `f5xc-terraform` listed with a `✓ Connected` status.

### Quick Setup for Visual Studio Code

VS Code 1.99+ supports MCP servers through GitHub Copilot. Choose the installation method that best fits your environment:

#### Option 1: Workspace Configuration (Recommended)

Create a `.vscode/mcp.json` file in your workspace:

```json
{
  "servers": {
    "f5xc-terraform": {
      "command": "npx",
      "args": ["-y", "@robinmordasiewicz/f5xc-terraform-mcp"]
    }
  }
}
```

#### Option 2: Corporate Environments (No Node.js Required)

For environments where npm/Node.js cannot be installed:

1. Download the latest `.mcpb` bundle from [GitHub Releases](https://github.com/robinmordasiewicz/terraform-provider-f5xc/releases)
2. Place the file in a known location (e.g., `~/.mcp/f5xc-terraform-mcp.mcpb`)
3. Create a `.vscode/mcp.json` file:

```json
{
  "servers": {
    "f5xc-terraform": {
      "command": "/path/to/f5xc-terraform-mcp.mcpb"
    }
  }
}
```

#### Verify Installation

1. Press `Ctrl+Shift+P` / `Cmd+Shift+P`
2. Run `MCP: List Servers`
3. Look for `f5xc-terraform` with a green status indicator

For complete MCP server documentation including all available tools and advanced configuration options, see the [NPM package page](https://www.npmjs.com/package/@robinmordasiewicz/f5xc-terraform-mcp).

## Resources and Data Sources

Browse the documentation sidebar for the complete list of resources and data sources organized by category.

<!-- Template version: 1.0.1 - MCP token optimization (#592) -->
