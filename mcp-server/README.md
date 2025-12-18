# F5 Distributed Cloud Terraform MCP Server

MCP server providing AI assistants with access to F5 Distributed Cloud (F5XC) Terraform provider documentation, 270+ OpenAPI specifications, subscription tier information, and addon service workflows.

## Installation

### Option 1: npm (Requires Node.js 18+)

```bash
# Run directly with npx
npx @robinmordasiewicz/f5xc-terraform-mcp

# Or install globally
npm install -g @robinmordasiewicz/f5xc-terraform-mcp
f5xc-terraform-mcp
```

### Option 2: Binary Bundle (No Dependencies)

Download the standalone `.mcpb` binary - no Node.js or Docker required:

> **Note:** The `.mcpb` file is a standalone executable, NOT a VSCode extension. Do not try to install it through VSCode's extension manager. Configure it via `mcp.json` as shown in the Configuration section below.

1. Download from [GitHub Releases](https://github.com/robinmordasiewicz/terraform-provider-f5xc/releases/latest) - look for `f5xc-terraform-mcp-X.X.X.mcpb`

2. Make executable and run:
```bash
chmod +x f5xc-terraform-mcp-*.mcpb
./f5xc-terraform-mcp-*.mcpb
```

### Option 3: Portable Node.js (No Admin Rights)

For users without admin privileges who cannot install Node.js system-wide:

1. Download Node.js portable binaries from [nodejs.org/download/prebuilt-binaries](https://nodejs.org/en/download/prebuilt-binaries)
2. Extract to user directory (e.g., `~/node/`)
3. Add to PATH: `export PATH="$HOME/node/bin:$PATH"`
4. Use npx as normal

## Configuration

### Claude Desktop

Add to your Claude Desktop configuration (`~/Library/Application Support/Claude/claude_desktop_config.json` on macOS):

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

Or with the binary bundle:

```json
{
  "mcpServers": {
    "f5xc-terraform": {
      "command": "/path/to/f5xc-terraform-mcp.mcpb"
    }
  }
}
```

### Claude Code CLI

```bash
claude mcp add f5xc-terraform -- npx -y @robinmordasiewicz/f5xc-terraform-mcp
```

### VSCode with Cline/Continue

Add to your MCP settings:

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

The server provides 6 consolidated tools optimized for token efficiency:

| Tool | Description |
|------|-------------|
| `f5xc_terraform_discover` | Discovery meta-tool - lists available tools and their schemas |
| `f5xc_terraform_docs` | Search, get, or list provider documentation (resources, data sources, functions, guides) |
| `f5xc_terraform_api` | Search OpenAPI specs, get spec details, find endpoints, get schema definitions |
| `f5xc_terraform_subscription` | Get subscription tier requirements for resources |
| `f5xc_terraform_addon` | Get addon service activation information |
| `f5xc_terraform_get_summary` | Get provider overview with statistics |

### Example Usage

**Search documentation:**
```
f5xc_terraform_docs(operation: "search", query: "http_loadbalancer")
```

**Get full resource documentation:**
```
f5xc_terraform_docs(operation: "get", name: "http_loadbalancer", type: "resource")
```

**Find API endpoints:**
```
f5xc_terraform_api(operation: "find_endpoints", pattern: "/namespaces", method: "POST")
```

**Check subscription requirements:**
```
f5xc_terraform_subscription(operation: "get", resource: "bot_defense_advanced_policy")
```

## What's Included

- **144+ Terraform Resources** - Full documentation for all F5XC resources
- **270+ OpenAPI Specifications** - Complete API reference for F5XC services
- **Subscription Tier Info** - Which features require which subscription level
- **Addon Service Workflows** - Activation status and requirements
- **Example Configurations** - Working Terraform examples for each resource

## Version Synchronization

The MCP server version is automatically synchronized with the Terraform provider releases. Both always share the same version number (e.g., v2.15.2).

## Links

- [Terraform Provider](https://registry.terraform.io/providers/robinmordasiewicz/f5xc/latest)
- [GitHub Repository](https://github.com/robinmordasiewicz/terraform-provider-f5xc)
- [npm Package](https://www.npmjs.com/package/@robinmordasiewicz/f5xc-terraform-mcp)
- [F5 Distributed Cloud Documentation](https://docs.cloud.f5.com/)

## License

MIT
