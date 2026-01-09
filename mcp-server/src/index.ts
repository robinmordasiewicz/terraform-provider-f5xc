#!/usr/bin/env node
/**
 * F5 Distributed Cloud Terraform Provider MCP Server
 *
 * Token-optimized MCP server with consolidated tools for AI assistants.
 *
 * This MCP server provides:
 * - Tool discovery meta-tool for lazy loading
 * - Consolidated documentation tool (search/get/list)
 * - Consolidated API specification tool (search/get/endpoints/definitions)
 * - Consolidated subscription tier tool
 * - Consolidated addon service tool
 * - Resource metadata tool for deterministic AI configuration generation
 * - Provider summary tool
 *
 * Token Optimization:
 * - 14 tools consolidated to 7 tools (~75% token reduction)
 * - Optimized descriptions using shared parameter descriptions
 * - Discovery meta-tool enables lazy schema loading
 *
 * Version Synchronization:
 * The npm package version is automatically synced with GitHub releases via CI/CD.
 * Both the Terraform provider and MCP server always share the same version number.
 *
 * @see https://github.com/robinmordasiewicz/terraform-provider-f5xc
 * @see https://www.npmjs.com/package/@robinmordasiewicz/f5xc-terraform-mcp
 */

import { createRequire } from 'module';
import { McpServer } from '@modelcontextprotocol/sdk/server/mcp.js';
import { StdioServerTransport } from '@modelcontextprotocol/sdk/server/stdio.js';

// Read version from package.json at runtime
const require = createRequire(import.meta.url);
const packageJson = require('../package.json');
const VERSION = packageJson.version;

import { ResponseFormat } from './types.js';
import { getDocumentationSummary } from './services/documentation.js';
import { getApiSpecsSummary } from './services/api-specs.js';

// Consolidated tool handlers
import { handleDiscover, DISCOVER_TOOL_DEFINITION } from './tools/discover.js';
import { handleDocs, DOCS_TOOL_DEFINITION } from './tools/docs.js';
import { handleApi, API_TOOL_DEFINITION } from './tools/api.js';
import { handleSubscription, SUBSCRIPTION_TOOL_DEFINITION } from './tools/subscription.js';
import { handleAddon, ADDON_TOOL_DEFINITION } from './tools/addon.js';
import { handleMetadata, METADATA_TOOL_DEFINITION } from './tools/metadata.js';
import { handleAuth, AUTH_TOOL_DEFINITION, initializeAuth } from './tools/auth.js';

// Consolidated schemas
import {
  DiscoverSchema,
  DocsSchema,
  ApiSchema,
  SubscriptionSchema,
  AddonSchema,
  MetadataSchema,
  AuthSchema,
  ResponseFormatSchema,
  type DiscoverInput,
  type DocsInput,
  type ApiInput,
  type SubscriptionInput,
  type AddonInput,
  type MetadataInput,
  type AuthInput,
} from './schemas/common.js';

// Legacy schemas for get_summary tool
import { GetSummarySchema, type GetSummaryInput } from './schemas/index.js';

// Constants
const CHARACTER_LIMIT = 50000; // Maximum response size

// Create MCP server instance
const server = new McpServer({
  name: 'f5xc-terraform-mcp',
  version: VERSION,
});

// =============================================================================
// TOOL 1: DISCOVERY META-TOOL (NEW)
// =============================================================================

server.registerTool(
  DISCOVER_TOOL_DEFINITION.name,
  {
    title: 'Discover F5XC Terraform Tools',
    description: DISCOVER_TOOL_DEFINITION.description,
    inputSchema: DiscoverSchema,
    annotations: {
      readOnlyHint: true,
      destructiveHint: false,
      idempotentHint: true,
      openWorldHint: false,
    },
  },
  async (params: DiscoverInput) => {
    const result = await handleDiscover(params);
    return {
      content: [{ type: 'text', text: result }],
    };
  },
);

// =============================================================================
// TOOL 2: CONSOLIDATED DOCUMENTATION TOOL
// =============================================================================

server.registerTool(
  DOCS_TOOL_DEFINITION.name,
  {
    title: 'F5XC Terraform Documentation',
    description: DOCS_TOOL_DEFINITION.description,
    inputSchema: DocsSchema,
    annotations: {
      readOnlyHint: true,
      destructiveHint: false,
      idempotentHint: true,
      openWorldHint: false,
    },
  },
  async (params: DocsInput) => {
    const result = await handleDocs(params);

    // Truncate if too long
    let textContent = result;
    if (textContent.length > CHARACTER_LIMIT) {
      textContent = textContent.slice(0, CHARACTER_LIMIT) + '\n\n... (truncated)';
    }

    return {
      content: [{ type: 'text', text: textContent }],
    };
  },
);

// =============================================================================
// TOOL 3: CONSOLIDATED API TOOL
// =============================================================================

server.registerTool(
  API_TOOL_DEFINITION.name,
  {
    title: 'F5XC OpenAPI Specifications',
    description: API_TOOL_DEFINITION.description,
    inputSchema: ApiSchema,
    annotations: {
      readOnlyHint: true,
      destructiveHint: false,
      idempotentHint: true,
      openWorldHint: false,
    },
  },
  async (params: ApiInput) => {
    const result = await handleApi(params);

    // Truncate if too long
    let textContent = result;
    if (textContent.length > CHARACTER_LIMIT) {
      textContent =
        textContent.slice(0, CHARACTER_LIMIT) +
        '\n\n... (truncated, use include_definitions=false for smaller response)';
    }

    return {
      content: [{ type: 'text', text: textContent }],
    };
  },
);

// =============================================================================
// TOOL 4: CONSOLIDATED SUBSCRIPTION TOOL
// =============================================================================

server.registerTool(
  SUBSCRIPTION_TOOL_DEFINITION.name,
  {
    title: 'F5XC Subscription Tiers',
    description: SUBSCRIPTION_TOOL_DEFINITION.description,
    inputSchema: SubscriptionSchema,
    annotations: {
      readOnlyHint: true,
      destructiveHint: false,
      idempotentHint: true,
      openWorldHint: false,
    },
  },
  async (params: SubscriptionInput) => {
    const result = await handleSubscription(params);
    return {
      content: [{ type: 'text', text: result }],
    };
  },
);

// =============================================================================
// TOOL 5: CONSOLIDATED ADDON TOOL
// =============================================================================

server.registerTool(
  ADDON_TOOL_DEFINITION.name,
  {
    title: 'F5XC Addon Services',
    description: ADDON_TOOL_DEFINITION.description,
    inputSchema: AddonSchema,
    annotations: {
      readOnlyHint: true,
      destructiveHint: false,
      idempotentHint: true,
      openWorldHint: false,
    },
  },
  async (params: AddonInput) => {
    const result = await handleAddon(params);

    // Truncate if too long
    let textContent = result;
    if (textContent.length > CHARACTER_LIMIT) {
      textContent = textContent.slice(0, CHARACTER_LIMIT) + '\n\n... (truncated)';
    }

    return {
      content: [{ type: 'text', text: textContent }],
    };
  },
);

// =============================================================================
// TOOL 6: METADATA TOOL (NEW)
// =============================================================================

server.registerTool(
  METADATA_TOOL_DEFINITION.name,
  {
    title: 'F5XC Resource Metadata',
    description: METADATA_TOOL_DEFINITION.description,
    inputSchema: MetadataSchema,
    annotations: {
      readOnlyHint: true,
      destructiveHint: false,
      idempotentHint: true,
      openWorldHint: false,
    },
  },
  async (params: MetadataInput) => {
    const result = await handleMetadata(params);
    return {
      content: [{ type: 'text', text: result }],
    };
  },
);

// =============================================================================
// TOOL 7: SUMMARY TOOL (KEEP AS-IS)
// =============================================================================

server.registerTool(
  'f5xc_terraform_get_summary',
  {
    title: 'F5XC Provider Summary',
    description: 'Get provider documentation and API specs summary',
    inputSchema: GetSummarySchema,
    annotations: {
      readOnlyHint: true,
      destructiveHint: false,
      idempotentHint: true,
      openWorldHint: false,
    },
  },
  async (params: GetSummaryInput) => {
    const docsSummary = getDocumentationSummary();
    const apiSummary = getApiSpecsSummary();

    const output = {
      provider: {
        name: 'robinmordasiewicz/f5xc',
        registry_url: 'https://registry.terraform.io/providers/robinmordasiewicz/f5xc/latest',
        github_url: 'https://github.com/robinmordasiewicz/terraform-provider-f5xc',
        npm_package: '@robinmordasiewicz/f5xc-terraform-mcp',
      },
      documentation: {
        total: Object.values(docsSummary).reduce((a, b) => a + b, 0),
        by_type: docsSummary,
      },
      api_specifications: {
        total: apiSummary.total,
        categories: apiSummary.categories,
      },
    };

    let textContent: string;
    if (params.response_format === ResponseFormat.MARKDOWN) {
      const lines = [
        '# F5 Distributed Cloud Terraform Provider',
        '',
        '## Provider Information',
        '',
        '| Property | Value |',
        '|----------|-------|',
        '| **Provider Name** | `robinmordasiewicz/f5xc` |',
        '| **Registry URL** | https://registry.terraform.io/providers/robinmordasiewicz/f5xc/latest |',
        '| **GitHub** | https://github.com/robinmordasiewicz/terraform-provider-f5xc |',
        '| **npm Package** | `@robinmordasiewicz/f5xc-terraform-mcp` |',
        '',
        '## Quick Start',
        '',
        '### Terraform Configuration',
        '',
        '```hcl',
        'terraform {',
        '  required_providers {',
        '    f5xc = {',
        '      source = "robinmordasiewicz/f5xc"',
        '    }',
        '  }',
        '}',
        '',
        'provider "f5xc" {',
        '  # API Token authentication (recommended)',
        '  api_url   = "https://your-tenant.console.ves.volterra.io"',
        '  api_token = var.f5xc_api_token',
        '}',
        '```',
        '',
        '### Environment Variables',
        '',
        '| Variable | Description |',
        '|----------|-------------|',
        '| `F5XC_API_URL` | F5 Distributed Cloud API URL |',
        '| `F5XC_API_TOKEN` | API Token for authentication |',
        '| `F5XC_P12_FILE` | Path to P12 certificate file |',
        '| `F5XC_P12_PASSWORD` | Password for P12 certificate |',
        '| `F5XC_CERT` | Path to PEM certificate file |',
        '| `F5XC_KEY` | Path to PEM private key file |',
        '| `F5XC_CACERT` | Path to CA certificate file |',
        '',
        '## Documentation',
        '',
        `Total: ${output.documentation.total} items`,
        '',
      ];

      for (const [type, count] of Object.entries(docsSummary)) {
        lines.push(`- **${type}s**: ${count}`);
      }

      lines.push('');
      lines.push('## API Specifications');
      lines.push('');
      lines.push(`Total: ${apiSummary.total} OpenAPI specs`);
      lines.push('');
      lines.push('## Available Tools (Token-Optimized)');
      lines.push('');
      lines.push(
        '- `f5xc_terraform_discover` - Discover available tools with optional schema details',
      );
      lines.push(
        '- `f5xc_terraform_docs` - Search, get, or list documentation (operations: search, get, list)',
      );
      lines.push(
        '- `f5xc_terraform_api` - Query API specs (operations: search, get, find_endpoints, get_definition, list_definitions)',
      );
      lines.push(
        '- `f5xc_terraform_subscription` - Check subscription tiers (operations: resource, property)',
      );
      lines.push(
        '- `f5xc_terraform_addon` - Addon services (operations: list, check, workflow)',
      );
      lines.push(
        '- `f5xc_terraform_metadata` - Resource metadata for deterministic AI config generation (operations: oneof, validation, defaults, enums, attribute, requires_replace, tier, dependencies, troubleshoot, syntax, validate, example, mistakes, summary)',
      );
      lines.push('- `f5xc_terraform_get_summary` - This summary');
      lines.push(
        '- `f5xc_terraform_auth` - Authentication status, profiles, and validation (operations: status, list, switch, validate)',
      );
      lines.push('');
      lines.push('> **Token Optimization**: 14 tools consolidated to 8 tools (~70% reduction)');
      lines.push('');
      lines.push('> **Note**: Use `f5xc_terraform_docs` with `operation: "get", name: "provider"` for complete provider documentation.');

      textContent = lines.join('\n');
    } else {
      textContent = JSON.stringify(output, null, 2);
    }

    return {
      content: [{ type: 'text', text: textContent }],
      structuredContent: output,
    };
  },
);

// =============================================================================
// TOOL 8: AUTH TOOL
// =============================================================================

server.registerTool(
  AUTH_TOOL_DEFINITION.name,
  {
    title: 'F5XC Authentication',
    description: AUTH_TOOL_DEFINITION.description,
    inputSchema: AuthSchema,
    annotations: {
      readOnlyHint: true,
      destructiveHint: false,
      idempotentHint: true,
      openWorldHint: false,
    },
  },
  async (params: AuthInput) => {
    const result = await handleAuth(params);
    return {
      content: [{ type: 'text', text: result }],
    };
  },
);

// =============================================================================
// SERVER STARTUP
// =============================================================================

async function main() {
  // Initialize authentication system
  const credManager = await initializeAuth();
  const mode = credManager.isAuthenticated() ? 'authenticated' : 'documentation';
  const tenant = credManager.getTenant();

  const transport = new StdioServerTransport();
  await server.connect(transport);

  if (tenant) {
    console.error(`F5XC Terraform MCP server running (${mode} mode, tenant: ${tenant})`);
  } else {
    console.error(`F5XC Terraform MCP server running (${mode} mode)`);
  }
}

main().catch((error) => {
  console.error('Server error:', error);
  process.exit(1);
});
