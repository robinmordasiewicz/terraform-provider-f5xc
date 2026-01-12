# F5 Distributed Cloud Terraform MCP Server

MCP (Model Context Protocol) server providing AI assistants with access to F5 Distributed Cloud Terraform provider documentation, 270+ OpenAPI specifications, subscription tier information, and addon service activation workflows.

## Quick Start

### Claude Code CLI (Recommended)

```bash
claude mcp add f5xc-terraform -- npx -y @robinmordasiewicz/f5xc-terraform-mcp
```

You should see:
```
Added stdio MCP server f5xc-terraform with command: npx -y @robinmordasiewicz/f5xc-terraform-mcp to local config
```

**Verify installation:**
```bash
claude mcp list
```

Look for this line in the output (among your other MCP servers):
```
f5xc-terraform: npx -y @robinmordasiewicz/f5xc-terraform-mcp - ✓ Connected
```

**Alternative verification** - get detailed server info:
```bash
claude mcp get f5xc-terraform
```

You should see:
```
f5xc-terraform:
  Status: ✓ Connected
  Type: stdio
  Command: npx
  Args: -y @robinmordasiewicz/f5xc-terraform-mcp
```

### VS Code (Cline/Continue)

Add to your VS Code settings (`.vscode/mcp.json` or global MCP settings):

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

### Claude Desktop

Add to your Claude Desktop config file:

**macOS:** `~/Library/Application Support/Claude/claude_desktop_config.json`
**Windows:** `%APPDATA%\Claude\claude_desktop_config.json`

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

## Available Tools

| Tool | Description |
|------|-------------|
| `f5xc_terraform_discover` | Discover available tools with schema details |
| `f5xc_terraform_docs` | Search, get, or list Terraform documentation |
| `f5xc_terraform_api` | Query 270+ F5XC OpenAPI specifications |
| `f5xc_terraform_subscription` | Check resource subscription tier requirements |
| `f5xc_terraform_addon` | List, check, and get workflows for addon services |
| `f5xc_terraform_metadata` | Query resource metadata, validation patterns, and syntax |
| `f5xc_terraform_get_summary` | Get provider documentation summary |
| `f5xc_terraform_auth` | Check authentication status and get Terraform config |

## Usage Examples

### Get Documentation for a Resource

```
f5xc_terraform_docs(operation="get", name="http_loadbalancer", type="resource")
```

### Search for Resources

```
f5xc_terraform_docs(operation="search", query="waf security")
```

### Check Subscription Requirements

```
f5xc_terraform_subscription(operation="resource", resource_name="http_loadbalancer")
```

### Get Correct Terraform Syntax

```
f5xc_terraform_metadata(operation="syntax", resource="http_loadbalancer")
```

### Validate Terraform Configuration

```
f5xc_terraform_metadata(operation="validate", resource="origin_pool", config="healthcheck { interval_seconds = 30 }")
```

### Get Authentication Configuration

```
f5xc_terraform_auth(operation="terraform-env", output_type="shell")
```

## Authentication

The MCP server supports multiple authentication methods for accessing F5XC:

### Using f5xc-auth Profiles (Recommended)

Install and configure [f5xc-auth](https://github.com/robinmordasiewicz/f5xc-auth):

```bash
npm install -g @robinmordasiewicz/f5xc-auth
f5xc-auth login
```

The MCP server will automatically detect configured profiles.

### Environment Variables

Set these environment variables before starting Claude Code:

```bash
export F5XC_API_URL="https://your-tenant.console.ves.volterra.io"
export F5XC_API_TOKEN="your-api-token"
```

### Using P12 Certificate

```bash
export F5XC_API_URL="https://your-tenant.console.ves.volterra.io"
export F5XC_P12_FILE="/path/to/certificate.p12"
export F5XC_P12_PASSWORD="certificate-password"  # pragma: allowlist secret
```

## Critical Provider Information

### Provider Source

**Always use `robinmordasiewicz/f5xc` - never use deprecated providers!**

```hcl
terraform {
  required_providers {
    f5xc = {
      source  = "robinmordasiewicz/f5xc"
      version = "~> 3.0"
    }
  }
}

provider "f5xc" {}
```

### Empty Block Syntax

This provider uses **empty blocks** `{}` for mutually exclusive options, NOT boolean values:

```hcl
# CORRECT
no_tls {}
advertise_on_public_default_vip {}
round_robin {}

# WRONG - These will cause errors
no_tls = true
advertise_on_public_default_vip = true
round_robin = true
```

## Troubleshooting

### Server Not Connected

1. Verify Node.js is installed (v18+):
   ```bash
   node --version
   ```

2. Check server status:
   ```bash
   claude mcp get f5xc-terraform
   ```

3. Test the server manually (should output JSON-RPC messages):
   ```bash
   npx -y @robinmordasiewicz/f5xc-terraform-mcp
   ```
   Press `Ctrl+C` to exit after confirming it starts.

### Remove and Re-add Server

```bash
claude mcp remove f5xc-terraform
claude mcp add f5xc-terraform -- npx -y @robinmordasiewicz/f5xc-terraform-mcp
```

### Server Shows "Failed to connect"

If `claude mcp list` shows `✗ Failed to connect`:

1. Clear npm cache and retry:
   ```bash
   npm cache clean --force
   claude mcp remove f5xc-terraform
   claude mcp add f5xc-terraform -- npx -y @robinmordasiewicz/f5xc-terraform-mcp
   ```

2. Restart Claude Code to reconnect all MCP servers.

## Requirements

- Node.js 18 or later
- npm (comes with Node.js)
- Claude Code, Cline, Continue, or Claude Desktop

## Links

- [Terraform Registry](https://registry.terraform.io/providers/robinmordasiewicz/f5xc/latest)
- [GitHub Repository](https://github.com/robinmordasiewicz/terraform-provider-f5xc)
- [npm Package](https://www.npmjs.com/package/@robinmordasiewicz/f5xc-terraform-mcp)
- [Claude Code MCP Documentation](https://docs.anthropic.com/en/docs/claude-code/mcp)

## License

MIT
