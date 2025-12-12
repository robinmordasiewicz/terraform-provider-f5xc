# F5 Distributed Cloud Terraform Provider MCP Server

A Model Context Protocol (MCP) server that provides AI assistants with comprehensive access to the F5 Distributed Cloud (F5XC) Terraform provider documentation and OpenAPI specifications.

## Features

- **144+ Resource Documentation**: Complete Terraform resource docs with arguments, attributes, and examples
- **270+ OpenAPI Specifications**: Full F5XC API specifications for all services
- **Intelligent Search**: Search across documentation and API specs with relevance scoring
- **Schema Exploration**: Browse and query API schema definitions
- **Endpoint Discovery**: Find API endpoints by pattern across all specifications

## Installation

### From npm

```bash
npm install -g @robinmordasiewicz/f5xc-terraform-mcp
```

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
