# F5 Distributed Cloud Terraform Provider MCP Server

A Model Context Protocol (MCP) server that provides AI assistants with comprehensive access to the F5 Distributed Cloud (F5XC) Terraform provider documentation and OpenAPI specifications.

## Features

- **144+ Resource Documentation**: Complete Terraform resource docs with arguments, attributes, and examples
- **270+ OpenAPI Specifications**: Full F5XC API specifications for all services
- **Intelligent Search**: Search across documentation and API specs with relevance scoring
- **Schema Exploration**: Browse and query API schema definitions
- **Endpoint Discovery**: Find API endpoints by pattern across all specifications

## Installation

Choose the installation method that best fits your environment:

| Method | Best For | Requirements |
|--------|----------|--------------|
| [VSCode MCP Gallery](#vscode-mcp-gallery) | VSCode users (easiest) | VSCode 1.99+ |
| [npx (Recommended)](#from-npm) | Developers with Node.js | Node.js 18+ |
| [MCPB Bundle](#mcpb-bundle-no-nodejs-required) | Corporate laptops (no Node.js) | None |
| [From Source](#from-source) | Contributors | Node.js 18+, npm |

### VSCode MCP Gallery

The easiest way to install for VSCode users. Choose from multiple options:

**Option A: MCP Gallery Search (Recommended for VSCode 1.105+)**

1. Open VSCode
2. Open the Extensions view (`Ctrl+Shift+X` / `Cmd+Shift+X`)
3. Type `@MCP` in the search box to filter MCP servers
4. Search for `f5xc` or `F5 Distributed Cloud`
5. Click **Install**

**Option B: Command Palette**

1. Open VSCode
2. Press `Ctrl+Shift+P` / `Cmd+Shift+P` to open Command Palette
3. Run `MCP: Add Server`
4. Select `npm package`
5. Enter: `@robinmordasiewicz/f5xc-terraform-mcp`
6. Choose scope: **Global** (all workspaces) or **Workspace** (this project only)

### From npm

```bash
npm install -g @robinmordasiewicz/f5xc-terraform-mcp
```

Or run directly with npx (no installation required):

```bash
npx @robinmordasiewicz/f5xc-terraform-mcp
```

### MCPB Bundle (No Node.js Required)

For corporate environments where Node.js cannot be installed. The MCPB bundle is fully self-contained with all dependencies included.

**For VSCode:**

1. Download the latest `.mcpb` file from [GitHub Releases](https://github.com/robinmordasiewicz/terraform-provider-f5xc/releases)
2. In VSCode, press `Ctrl+Shift+P` / `Cmd+Shift+P`
3. Run `MCP: Add Server`
4. Drag and drop the `.mcpb` file, or select it when prompted
5. Run `MCP: List Servers` to verify installation

**For Claude Desktop:**

1. Download the latest `.mcpb` file from [GitHub Releases](https://github.com/robinmordasiewicz/terraform-provider-f5xc/releases)
2. Double-click the file to install, or drag it into Claude Desktop
3. Restart Claude Desktop to activate

**File**: `f5xc-terraform-mcp-X.Y.Z.mcpb`

-> **Note:** MCPB bundles are automatically built and attached to each GitHub Release. No npm or Node.js installation is required.

### From Source

```bash
cd mcp-server
npm install
npm run build
```

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

Or if running from source:

```json
{
  "mcpServers": {
    "f5xc": {
      "command": "node",
      "args": ["/path/to/terraform-provider-f5xc/mcp-server/dist/index.js"]
    }
  }
}
```

### Claude Code (CLI)

Install the MCP server with a single command:

```bash
claude mcp add f5xc-terraform -- npx -y @robinmordasiewicz/f5xc-terraform-mcp
```

**Scope Options:**

- `--scope local` (default): Available only in the current directory
- `--scope project`: Shared with anyone who clones the repository (saved in `.mcp.json`)
- `--scope user`: Available in all your Claude Code sessions

Examples with different scopes:

```bash
# User-wide installation (recommended for personal use)
claude mcp add --scope user f5xc-terraform -- npx -y @robinmordasiewicz/f5xc-terraform-mcp

# Project-specific installation (for team collaboration)
claude mcp add --scope project f5xc-terraform -- npx -y @robinmordasiewicz/f5xc-terraform-mcp
```

**Verify Installation:**

```bash
claude mcp list
```

You should see `f5xc-terraform` listed with a `✓ Connected` status.

**Remove Server:**

```bash
claude mcp remove f5xc-terraform
```

### Visual Studio Code (with GitHub Copilot)

VS Code 1.99+ supports MCP servers through GitHub Copilot. Configure by creating a `.vscode/mcp.json` file in your workspace:

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

**Alternative: Command Palette**

1. Open Command Palette (`Ctrl+Shift+P` / `Cmd+Shift+P`)
2. Run `MCP: Add Server`
3. Select `npm package`
4. Enter: `@robinmordasiewicz/f5xc-terraform-mcp`

**Alternative: Command Line**

```bash
code --add-mcp "{\"name\":\"f5xc-terraform\",\"command\":\"npx\",\"args\":[\"-y\",\"@robinmordasiewicz/f5xc-terraform-mcp\"]}"
```

**Global Configuration (User Settings)**

To make the MCP server available across all workspaces, add to your VS Code user settings (`settings.json`):

```json
{
  "mcp": {
    "servers": {
      "f5xc-terraform": {
        "command": "npx",
        "args": ["-y", "@robinmordasiewicz/f5xc-terraform-mcp"]
      }
    }
  }
}
```

**Verify VSCode Installation:**

1. Press `Ctrl+Shift+P` / `Cmd+Shift+P`
2. Run `MCP: List Servers`
3. Look for `f5xc-terraform` with a green status indicator

**Troubleshooting VSCode:**

- **Server not appearing**: Restart VSCode after adding the server configuration
- **Connection issues**: Check the Output panel (`View > Output`) and select "MCP" from the dropdown
- **Node.js not found**: Use the [MCPB Bundle](#mcpb-bundle-no-nodejs-required) method instead

## Available Tools

### Documentation Tools

| Tool | Description |
|------|-------------|
| `f5xc_terraform_search_docs` | Search provider documentation by keyword |
| `f5xc_terraform_get_doc` | Get complete documentation for a resource |
| `f5xc_terraform_list_docs` | List all available documentation |

### API Specification Tools

| Tool | Description |
|------|-------------|
| `f5xc_terraform_search_api_specs` | Search OpenAPI specifications |
| `f5xc_terraform_get_api_spec` | Get a specific API specification |
| `f5xc_terraform_find_endpoints` | Find API endpoints by URL pattern |
| `f5xc_terraform_get_schema_definition` | Get a schema definition from a spec |
| `f5xc_terraform_list_definitions` | List all definitions in a spec |

### Subscription Tier Tools

| Tool | Description |
|------|-------------|
| `f5xc_terraform_get_subscription_info` | Get subscription tier requirements for resources |
| `f5xc_terraform_get_property_subscription_info` | Get property-level subscription tier indicators |

### Addon Service Tools

| Tool | Description |
|------|-------------|
| `f5xc_terraform_addon_list_services` | List available addon services with activation requirements |
| `f5xc_terraform_addon_check_activation` | Check if an addon service is activated for the tenant |
| `f5xc_terraform_addon_activation_workflow` | Get activation workflow and Terraform code for addons |

### Utility Tools

| Tool | Description |
|------|-------------|
| `f5xc_terraform_get_summary` | Get overview of all available docs and specs |

## Usage Examples

### Find HTTP Load Balancer Documentation

```
User: How do I configure an HTTP load balancer in F5XC with Terraform?

Claude: [Uses f5xc_terraform_search_docs with query "http_loadbalancer"]
        [Uses f5xc_terraform_get_doc with name "http_loadbalancer"]
```

### Discover API Endpoints

```
User: What API endpoints are available for managing namespaces?

Claude: [Uses f5xc_terraform_find_endpoints with pattern "/namespaces"]
```

### Explore Schema Definitions

```
User: What fields are available in the app_firewall configuration?

Claude: [Uses f5xc_terraform_get_api_spec with name "app_firewall"]
        [Uses f5xc_terraform_list_definitions with spec_name "app_firewall"]
        [Uses f5xc_terraform_get_schema_definition for specific type]
```

## Development

### Build

```bash
npm run build
```

### Development Mode (with auto-reload)

```bash
npm run dev
```

### Type Check

```bash
npm run typecheck
```

## Architecture

```
mcp-server/
├── src/
│   ├── index.ts          # MCP server entry point
│   ├── types.ts          # TypeScript type definitions
│   ├── schemas/          # Zod validation schemas
│   │   └── index.ts
│   └── services/         # Core services
│       ├── documentation.ts   # Doc loading and search
│       └── api-specs.ts       # OpenAPI spec handling
├── package.json
├── tsconfig.json
└── README.md
```

## Resources Served

### Documentation Types

- **Resources**: Terraform resources (http_loadbalancer, origin_pool, namespace, etc.)
- **Data Sources**: Terraform data sources for reading existing resources
- **Functions**: Provider-defined functions (blindfold, blindfold_file)
- **Guides**: Step-by-step tutorials and how-to guides

### API Specifications

All F5 Distributed Cloud public APIs including:
- HTTP/TCP Load Balancers
- Origin Pools
- Application Firewalls (WAF)
- Namespaces
- DNS Management
- Network Policies
- Cloud Sites (AWS, Azure, GCP)
- And 260+ more...

## License

MIT

## Contributing

Contributions welcome! Please see the main [terraform-provider-f5xc](https://github.com/robinmordasiewicz/terraform-provider-f5xc) repository for contribution guidelines.
