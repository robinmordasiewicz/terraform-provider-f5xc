#!/usr/bin/env node
/**
 * F5 Distributed Cloud Terraform Provider MCP Server
 *
 * This MCP server provides AI assistants with access to:
 * - Terraform provider documentation (resources, data sources, functions, guides)
 * - F5 Distributed Cloud OpenAPI specifications (270+ specs)
 * - Search and query capabilities for both documentation and API specs
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
import {
  searchDocumentation,
  getDocumentation,
  listDocumentation,
  getDocumentationSummary,
  getResourceSubscriptionInfo,
  getAdvancedTierResources,
  getSubscriptionSummary,
  getPropertySubscriptionInfo,
  getResourceAdvancedProperties,
} from './services/documentation.js';
import {
  searchApiSpecs,
  getApiSpec,
  listApiSpecs,
  findEndpoints,
  getSchemaDefinition,
  listDefinitions,
  getApiSpecsSummary,
} from './services/api-specs.js';
import {
  listAddonServices,
  checkAddonActivation,
  getAddonWorkflow,
} from './services/addons.js';
import {
  SearchDocsSchema,
  GetDocSchema,
  ListDocsSchema,
  SearchApiSpecsSchema,
  GetApiSpecSchema,
  FindEndpointsSchema,
  GetSchemaDefSchema,
  ListDefinitionsSchema,
  GetSummarySchema,
  GetSubscriptionInfoSchema,
  GetPropertySubscriptionInfoSchema,
  ListAddonServicesSchema,
  CheckAddonActivationSchema,
  GetAddonWorkflowSchema,
  type SearchDocsInput,
  type GetDocInput,
  type ListDocsInput,
  type SearchApiSpecsInput,
  type GetApiSpecInput,
  type FindEndpointsInput,
  type GetSchemaDefInput,
  type ListDefinitionsInput,
  type GetSummaryInput,
  type GetSubscriptionInfoInput,
  type GetPropertySubscriptionInfoInput,
  type ListAddonServicesInput,
  type CheckAddonActivationInput,
  type GetAddonWorkflowInput,
} from './schemas/index.js';

// Constants
const CHARACTER_LIMIT = 50000; // Maximum response size

// Create MCP server instance
const server = new McpServer({
  name: 'f5xc-terraform-mcp',
  version: VERSION,
});

// =============================================================================
// DOCUMENTATION TOOLS
// =============================================================================

server.registerTool(
  'f5xc_terraform_search_docs',
  {
    title: 'Search F5XC Documentation',
    description: `Search the F5 Distributed Cloud Terraform provider documentation.

Searches across all resource documentation, data sources, functions, and guides
to find relevant information about configuring F5XC resources with Terraform.

Args:
  - query (string): Search terms (resource names, attributes, descriptions)
  - type (optional): Filter by 'resource', 'data-source', 'function', or 'guide'
  - limit (number): Maximum results (default: 20, max: 50)
  - response_format: 'markdown' or 'json'

Returns:
  List of matching documentation with relevance scores and snippets.

Examples:
  - "http_loadbalancer" -> Find HTTP load balancer documentation
  - "origin pool" -> Find origin pool related docs
  - "waf" -> Find Web Application Firewall docs`,
    inputSchema: SearchDocsSchema,
    annotations: {
      readOnlyHint: true,
      destructiveHint: false,
      idempotentHint: true,
      openWorldHint: false,
    },
  },
  async (params: SearchDocsInput) => {
    const results = searchDocumentation(params.query, params.type, params.limit);

    if (results.length === 0) {
      return {
        content: [{
          type: 'text',
          text: `No documentation found matching "${params.query}"`,
        }],
      };
    }

    const output = {
      query: params.query,
      total: results.length,
      results: results.map(r => ({
        name: r.name,
        type: r.type,
        score: r.score,
        snippet: r.snippet,
      })),
    };

    let textContent: string;
    if (params.response_format === ResponseFormat.MARKDOWN) {
      const lines = [
        `# Documentation Search Results: "${params.query}"`,
        '',
        `Found ${results.length} results:`,
        '',
      ];
      for (const result of results) {
        lines.push(`## ${result.name} (${result.type})`);
        lines.push(`Score: ${result.score}`);
        lines.push(result.snippet);
        lines.push('');
      }
      textContent = lines.join('\n');
    } else {
      textContent = JSON.stringify(output, null, 2);
    }

    return {
      content: [{ type: 'text', text: textContent }],
      structuredContent: output,
    };
  }
);

server.registerTool(
  'f5xc_terraform_get_doc',
  {
    title: 'Get F5XC Resource Documentation',
    description: `Get complete documentation for a specific F5XC Terraform resource.

Retrieves the full markdown documentation including:
- Resource description and usage
- Argument reference (all configurable attributes)
- Attribute reference (computed/read-only attributes)
- Import instructions
- Example configurations

Args:
  - name (string): Resource name (e.g., "http_loadbalancer", "namespace", "origin_pool")
  - type (optional): 'resource', 'data-source', 'function', or 'guide'
  - response_format: 'markdown' or 'json'

Returns:
  Complete documentation content.

Examples:
  - name="http_loadbalancer" -> HTTP Load Balancer resource docs
  - name="app_firewall" -> Application Firewall (WAF) docs
  - name="blindfold", type="function" -> Blindfold encryption function docs`,
    inputSchema: GetDocSchema,
    annotations: {
      readOnlyHint: true,
      destructiveHint: false,
      idempotentHint: true,
      openWorldHint: false,
    },
  },
  async (params: GetDocInput) => {
    const doc = getDocumentation(params.name, params.type);

    if (!doc) {
      return {
        content: [{
          type: 'text',
          text: `Documentation not found for "${params.name}"${params.type ? ` (type: ${params.type})` : ''}. Use f5xc_terraform_search_docs to find available documentation.`,
        }],
      };
    }

    const output = {
      name: doc.name,
      type: doc.type,
      path: doc.path,
      content: doc.content,
    };

    let textContent: string;
    if (params.response_format === ResponseFormat.MARKDOWN) {
      textContent = doc.content || `# ${doc.name}\n\nNo content available.`;
    } else {
      textContent = JSON.stringify(output, null, 2);
    }

    // Truncate if too long
    if (textContent.length > CHARACTER_LIMIT) {
      textContent = textContent.slice(0, CHARACTER_LIMIT) + '\n\n... (truncated)';
    }

    return {
      content: [{ type: 'text', text: textContent }],
      structuredContent: output,
    };
  }
);

server.registerTool(
  'f5xc_terraform_list_docs',
  {
    title: 'List F5XC Documentation',
    description: `List all available F5 Distributed Cloud Terraform provider documentation.

Lists all resources, data sources, functions, and guides available in the provider.

Args:
  - type (optional): Filter by 'resource', 'data-source', 'function', or 'guide'
  - response_format: 'markdown' or 'json'

Returns:
  List of all available documentation items.`,
    inputSchema: ListDocsSchema,
    annotations: {
      readOnlyHint: true,
      destructiveHint: false,
      idempotentHint: true,
      openWorldHint: false,
    },
  },
  async (params: ListDocsInput) => {
    const docs = listDocumentation(params.type);

    const output = {
      total: docs.length,
      type_filter: params.type || 'all',
      items: docs.map(d => ({
        name: d.name,
        type: d.type,
      })),
    };

    let textContent: string;
    if (params.response_format === ResponseFormat.MARKDOWN) {
      const lines = [
        `# F5XC Terraform Documentation`,
        '',
        `Total: ${docs.length} items${params.type ? ` (filtered: ${params.type})` : ''}`,
        '',
      ];

      // Group by type
      const byType: Record<string, string[]> = {};
      for (const doc of docs) {
        if (!byType[doc.type]) byType[doc.type] = [];
        byType[doc.type].push(doc.name);
      }

      for (const [type, names] of Object.entries(byType)) {
        lines.push(`## ${type}s (${names.length})`);
        lines.push(names.sort().join(', '));
        lines.push('');
      }

      textContent = lines.join('\n');
    } else {
      textContent = JSON.stringify(output, null, 2);
    }

    return {
      content: [{ type: 'text', text: textContent }],
      structuredContent: output,
    };
  }
);

// =============================================================================
// API SPECIFICATION TOOLS
// =============================================================================

server.registerTool(
  'f5xc_terraform_search_api_specs',
  {
    title: 'Search F5XC API Specifications',
    description: `Search the F5 Distributed Cloud OpenAPI specifications.

The provider includes 270+ OpenAPI specs covering all F5XC API endpoints.
Use this to find API specifications for specific services or operations.

Args:
  - query (string): Search terms (schema names, service names)
  - limit (number): Maximum results (default: 20, max: 50)
  - response_format: 'markdown' or 'json'

Returns:
  List of matching API specifications with titles and descriptions.

Examples:
  - "http_loadbalancer" -> Find HTTP LB API spec
  - "namespace" -> Find namespace management API
  - "waf" or "app_firewall" -> Find WAF API specs`,
    inputSchema: SearchApiSpecsSchema,
    annotations: {
      readOnlyHint: true,
      destructiveHint: false,
      idempotentHint: true,
      openWorldHint: false,
    },
  },
  async (params: SearchApiSpecsInput) => {
    const results = searchApiSpecs(params.query, params.limit);

    if (results.length === 0) {
      return {
        content: [{
          type: 'text',
          text: `No API specifications found matching "${params.query}"`,
        }],
      };
    }

    const output = {
      query: params.query,
      total: results.length,
      results: results.map(r => ({
        name: r.name,
        snippet: r.snippet,
        score: r.score,
      })),
    };

    let textContent: string;
    if (params.response_format === ResponseFormat.MARKDOWN) {
      const lines = [
        `# API Specification Search: "${params.query}"`,
        '',
        `Found ${results.length} specifications:`,
        '',
      ];
      for (const result of results) {
        lines.push(`## ${result.name}`);
        lines.push(result.snippet);
        lines.push('');
      }
      textContent = lines.join('\n');
    } else {
      textContent = JSON.stringify(output, null, 2);
    }

    return {
      content: [{ type: 'text', text: textContent }],
      structuredContent: output,
    };
  }
);

server.registerTool(
  'f5xc_terraform_get_api_spec',
  {
    title: 'Get F5XC API Specification',
    description: `Get a specific F5 Distributed Cloud OpenAPI specification.

Retrieves the complete OpenAPI spec including:
- API info and description
- Available endpoints and methods
- Request/response schemas
- Authentication requirements

Args:
  - name (string): Spec name (e.g., "http_loadbalancer", "namespace", "app_firewall")
  - include_paths (boolean): Include endpoint paths (default: true)
  - include_definitions (boolean): Include schema definitions (default: false, can be large)
  - response_format: 'markdown' or 'json'

Returns:
  OpenAPI specification content.`,
    inputSchema: GetApiSpecSchema,
    annotations: {
      readOnlyHint: true,
      destructiveHint: false,
      idempotentHint: true,
      openWorldHint: false,
    },
  },
  async (params: GetApiSpecInput) => {
    const spec = getApiSpec(params.name);

    if (!spec?.content) {
      return {
        content: [{
          type: 'text',
          text: `API specification not found for "${params.name}". Use f5xc_terraform_search_api_specs to find available specs.`,
        }],
      };
    }

    const content = spec.content;
    const output: Record<string, unknown> = {
      name: spec.schemaName,
      info: content.info,
    };

    if (params.include_paths && content.paths) {
      output.paths = Object.keys(content.paths).map(path => {
        const pathItem = content.paths![path];
        const methods: string[] = [];
        if (pathItem.get) methods.push('GET');
        if (pathItem.post) methods.push('POST');
        if (pathItem.put) methods.push('PUT');
        if (pathItem.delete) methods.push('DELETE');
        if (pathItem.patch) methods.push('PATCH');
        return { path, methods };
      });
    }

    if (params.include_definitions) {
      output.definitions = content.definitions || content.components?.schemas || {};
    }

    let textContent: string;
    if (params.response_format === ResponseFormat.MARKDOWN) {
      const lines = [
        `# ${content.info?.title || spec.schemaName}`,
        '',
        content.info?.description || '',
        '',
        `**Version**: ${content.info?.version || 'N/A'}`,
        '',
      ];

      if (params.include_paths && content.paths) {
        lines.push('## Endpoints');
        lines.push('');
        for (const [path, pathItem] of Object.entries(content.paths)) {
          const methods: string[] = [];
          if (pathItem.get) methods.push('GET');
          if (pathItem.post) methods.push('POST');
          if (pathItem.put) methods.push('PUT');
          if (pathItem.delete) methods.push('DELETE');
          if (pathItem.patch) methods.push('PATCH');
          lines.push(`- \`${methods.join('|')} ${path}\``);
        }
        lines.push('');
      }

      textContent = lines.join('\n');
    } else {
      textContent = JSON.stringify(output, null, 2);
    }

    // Truncate if too long
    if (textContent.length > CHARACTER_LIMIT) {
      textContent = textContent.slice(0, CHARACTER_LIMIT) + '\n\n... (truncated, use include_definitions=false for smaller response)';
    }

    return {
      content: [{ type: 'text', text: textContent }],
      structuredContent: output,
    };
  }
);

server.registerTool(
  'f5xc_terraform_find_endpoints',
  {
    title: 'Find F5XC API Endpoints',
    description: `Find API endpoints across all F5XC OpenAPI specifications.

Searches through 270+ API specs to find endpoints matching a pattern.
Useful for discovering which APIs to use for specific operations.

Args:
  - pattern (string): URL pattern to search (e.g., "/namespaces", "http_loadbalancer")
  - method (optional): Filter by HTTP method (GET, POST, PUT, DELETE, PATCH)
  - limit (number): Maximum results (default: 20, max: 100)
  - response_format: 'markdown' or 'json'

Returns:
  List of matching endpoints with spec name, path, method, and description.

Examples:
  - pattern="/namespaces" -> All namespace-related endpoints
  - pattern="config", method="POST" -> POST endpoints for configuration`,
    inputSchema: FindEndpointsSchema,
    annotations: {
      readOnlyHint: true,
      destructiveHint: false,
      idempotentHint: true,
      openWorldHint: false,
    },
  },
  async (params: FindEndpointsInput) => {
    const endpoints = findEndpoints(params.pattern, params.method, params.limit);

    if (endpoints.length === 0) {
      return {
        content: [{
          type: 'text',
          text: `No endpoints found matching pattern "${params.pattern}"${params.method ? ` with method ${params.method}` : ''}`,
        }],
      };
    }

    const output = {
      pattern: params.pattern,
      method_filter: params.method || 'all',
      total: endpoints.length,
      endpoints,
    };

    let textContent: string;
    if (params.response_format === ResponseFormat.MARKDOWN) {
      const lines = [
        `# API Endpoints: "${params.pattern}"`,
        '',
        `Found ${endpoints.length} endpoints:`,
        '',
      ];
      for (const ep of endpoints) {
        lines.push(`## ${ep.method} ${ep.path}`);
        lines.push(`**Spec**: ${ep.specName}`);
        if (ep.summary) lines.push(`**Summary**: ${ep.summary}`);
        if (ep.operationId) lines.push(`**Operation ID**: ${ep.operationId}`);
        lines.push('');
      }
      textContent = lines.join('\n');
    } else {
      textContent = JSON.stringify(output, null, 2);
    }

    return {
      content: [{ type: 'text', text: textContent }],
      structuredContent: output,
    };
  }
);

server.registerTool(
  'f5xc_terraform_get_schema_definition',
  {
    title: 'Get API Schema Definition',
    description: `Get a specific schema definition from an F5XC OpenAPI spec.

Retrieves the JSON schema for a specific type definition, useful for
understanding the structure of API request/response objects.

Args:
  - spec_name (string): Name of the API spec
  - definition_name (string): Name of the schema definition
  - response_format: 'markdown' or 'json'

Returns:
  JSON schema definition with properties, types, and descriptions.`,
    inputSchema: GetSchemaDefSchema,
    annotations: {
      readOnlyHint: true,
      destructiveHint: false,
      idempotentHint: true,
      openWorldHint: false,
    },
  },
  async (params: GetSchemaDefInput) => {
    const definition = getSchemaDefinition(params.spec_name, params.definition_name);

    if (!definition) {
      return {
        content: [{
          type: 'text',
          text: `Schema definition "${params.definition_name}" not found in spec "${params.spec_name}". Use f5xc_terraform_list_definitions to see available definitions.`,
        }],
      };
    }

    const output = {
      spec_name: params.spec_name,
      definition_name: params.definition_name,
      schema: definition,
    };

    let textContent: string;
    if (params.response_format === ResponseFormat.MARKDOWN) {
      const lines = [
        `# Schema: ${params.definition_name}`,
        '',
        `**Spec**: ${params.spec_name}`,
        '',
        '```json',
        JSON.stringify(definition, null, 2),
        '```',
      ];
      textContent = lines.join('\n');
    } else {
      textContent = JSON.stringify(output, null, 2);
    }

    // Truncate if too long
    if (textContent.length > CHARACTER_LIMIT) {
      textContent = textContent.slice(0, CHARACTER_LIMIT) + '\n\n... (truncated)';
    }

    return {
      content: [{ type: 'text', text: textContent }],
      structuredContent: output,
    };
  }
);

server.registerTool(
  'f5xc_terraform_list_definitions',
  {
    title: 'List API Schema Definitions',
    description: `List all schema definitions in an F5XC OpenAPI spec.

Lists the names of all type definitions available in a specification,
useful for discovering what schemas are available to query.

Args:
  - spec_name (string): Name of the API spec
  - response_format: 'markdown' or 'json'

Returns:
  List of definition names in the spec.`,
    inputSchema: ListDefinitionsSchema,
    annotations: {
      readOnlyHint: true,
      destructiveHint: false,
      idempotentHint: true,
      openWorldHint: false,
    },
  },
  async (params: ListDefinitionsInput) => {
    const definitions = listDefinitions(params.spec_name);

    if (definitions.length === 0) {
      return {
        content: [{
          type: 'text',
          text: `No definitions found in spec "${params.spec_name}". Use f5xc_terraform_search_api_specs to find available specs.`,
        }],
      };
    }

    const output = {
      spec_name: params.spec_name,
      total: definitions.length,
      definitions,
    };

    let textContent: string;
    if (params.response_format === ResponseFormat.MARKDOWN) {
      const lines = [
        `# Schema Definitions: ${params.spec_name}`,
        '',
        `Total: ${definitions.length} definitions`,
        '',
        definitions.sort().join(', '),
      ];
      textContent = lines.join('\n');
    } else {
      textContent = JSON.stringify(output, null, 2);
    }

    return {
      content: [{ type: 'text', text: textContent }],
      structuredContent: output,
    };
  }
);

// =============================================================================
// SUMMARY TOOL
// =============================================================================

server.registerTool(
  'f5xc_terraform_get_summary',
  {
    title: 'Get F5XC Provider Summary',
    description: `Get a summary of all available F5 Distributed Cloud Terraform provider documentation and API specifications.

Provides an overview of:
- Total number of resources, data sources, functions, and guides
- Total number of API specifications
- Categories and counts

Useful as a starting point to understand what's available.

Args:
  - response_format: 'markdown' or 'json'

Returns:
  Summary statistics of all documentation and API specs.`,
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
      lines.push('### Available Tools');
      lines.push('');
      lines.push('- `f5xc_terraform_search_docs` - Search documentation');
      lines.push('- `f5xc_terraform_get_doc` - Get specific resource documentation');
      lines.push('- `f5xc_terraform_list_docs` - List all documentation');
      lines.push('- `f5xc_terraform_search_api_specs` - Search API specifications');
      lines.push('- `f5xc_terraform_get_api_spec` - Get specific API spec');
      lines.push('- `f5xc_terraform_find_endpoints` - Find API endpoints');
      lines.push('- `f5xc_terraform_get_schema_definition` - Get schema definition');
      lines.push('- `f5xc_terraform_list_definitions` - List definitions in a spec');
      lines.push('- `f5xc_terraform_get_subscription_info` - Check resource subscription tier requirements');
      lines.push('- `f5xc_terraform_get_property_subscription_info` - Check property-level subscription tier requirements');
      lines.push('- `f5xc_terraform_addon_list_services` - List all addon services with tier requirements');
      lines.push('- `f5xc_terraform_addon_check_activation` - Check activation requirements for an addon');
      lines.push('- `f5xc_terraform_addon_activation_workflow` - Get detailed activation workflow with Terraform config');

      textContent = lines.join('\n');
    } else {
      textContent = JSON.stringify(output, null, 2);
    }

    return {
      content: [{ type: 'text', text: textContent }],
      structuredContent: output,
    };
  }
);

// =============================================================================
// SUBSCRIPTION TIER TOOL
// =============================================================================

server.registerTool(
  'f5xc_terraform_get_subscription_info',
  {
    title: 'Get F5XC Subscription Tier Info',
    description: `Get subscription tier requirements for F5 Distributed Cloud resources.

Returns information about which resources require an Advanced subscription tier.
Only Standard and Advanced subscription tiers are available.
Resources not requiring Advanced are available with Standard subscription.

Args:
  - resource (optional): Specific resource name to check
  - tier (optional): Filter by 'STANDARD' or 'ADVANCED'
  - response_format: 'markdown' or 'json'

Returns:
  - For specific resource: Tier requirement and any Advanced-only features
  - For no resource specified: List of all resources requiring Advanced tier

Examples:
  - resource="http_loadbalancer" -> Shows features requiring Advanced
  - tier="ADVANCED" -> Lists all resources requiring Advanced subscription
  - No args -> Summary of Advanced tier requirements`,
    inputSchema: GetSubscriptionInfoSchema,
    annotations: {
      readOnlyHint: true,
      destructiveHint: false,
      idempotentHint: true,
      openWorldHint: false,
    },
  },
  async (params: GetSubscriptionInfoInput) => {
    // Query for specific resource
    if (params.resource) {
      const info = getResourceSubscriptionInfo(params.resource);

      if (!info) {
        return {
          content: [{
            type: 'text',
            text: `No subscription information found for resource "${params.resource}". The resource may not exist or may not have tier metadata.`,
          }],
        };
      }

      const output = {
        resource: params.resource,
        tier: info.tier,
        service: info.service,
        requires_advanced: info.requiresAdvanced,
        advanced_features: info.advancedFeatures || [],
      };

      let textContent: string;
      if (params.response_format === ResponseFormat.MARKDOWN) {
        const lines = [
          `# Subscription Info: ${params.resource}`,
          '',
          `**Minimum Tier**: ${info.tier}`,
          `**Service**: ${info.service}`,
          `**Requires Advanced**: ${info.requiresAdvanced ? 'Yes' : 'No'}`,
          '',
        ];

        if (info.advancedFeatures && info.advancedFeatures.length > 0) {
          lines.push('## Features Requiring Advanced Subscription');
          lines.push('');
          for (const feature of info.advancedFeatures) {
            lines.push(`- \`${feature}\``);
          }
        } else if (!info.requiresAdvanced) {
          lines.push('This resource is fully available with a Standard subscription.');
        }

        textContent = lines.join('\n');
      } else {
        textContent = JSON.stringify(output, null, 2);
      }

      return {
        content: [{ type: 'text', text: textContent }],
        structuredContent: output,
      };
    }

    // List all Advanced tier resources
    const advancedResources = getAdvancedTierResources();
    const summary = getSubscriptionSummary();

    // Filter by tier if specified
    let filteredResources = advancedResources;
    if (params.tier === 'ADVANCED') {
      filteredResources = advancedResources.filter(r => r.subscriptionTier === 'ADVANCED');
    }

    const output = {
      summary: {
        total_resources: summary.totalResources,
        advanced_only_resources: summary.advancedOnlyResources,
        resources_with_advanced_features: summary.resourcesWithAdvancedFeatures,
        advanced_features: summary.advancedFeaturesList,
      },
      resources: filteredResources.map(r => ({
        name: r.name,
        type: r.type,
        tier: r.subscriptionTier,
        service: r.addonService,
        advanced_features: r.advancedFeatures,
      })),
    };

    let textContent: string;
    if (params.response_format === ResponseFormat.MARKDOWN) {
      const lines = [
        '# F5XC Subscription Tier Requirements',
        '',
        '## Summary',
        '',
        `- **Total Resources**: ${summary.totalResources}`,
        `- **Resources Requiring Advanced**: ${summary.advancedOnlyResources}`,
        `- **Resources with Advanced Features**: ${summary.resourcesWithAdvancedFeatures}`,
        '',
      ];

      if (filteredResources.length > 0) {
        lines.push('## Resources Requiring Advanced Subscription');
        lines.push('');
        for (const res of filteredResources) {
          lines.push(`### ${res.name} (${res.type})`);
          if (res.subscriptionTier === 'ADVANCED') {
            lines.push('**Requires Advanced subscription**');
          }
          if (res.advancedFeatures && res.advancedFeatures.length > 0) {
            lines.push('');
            lines.push('Advanced-only features:');
            for (const feat of res.advancedFeatures) {
              lines.push(`- \`${feat}\``);
            }
          }
          lines.push('');
        }
      }

      if (summary.advancedFeaturesList.length > 0) {
        lines.push('## All Advanced Features');
        lines.push('');
        lines.push(summary.advancedFeaturesList.map(f => `\`${f}\``).join(', '));
      }

      textContent = lines.join('\n');
    } else {
      textContent = JSON.stringify(output, null, 2);
    }

    return {
      content: [{ type: 'text', text: textContent }],
      structuredContent: output,
    };
  }
);

// =============================================================================
// PROPERTY SUBSCRIPTION TIER TOOL
// =============================================================================

server.registerTool(
  'f5xc_terraform_get_property_subscription_info',
  {
    title: 'Get Property Subscription Tier Info',
    description: `Get subscription tier requirements for a specific property within an F5XC resource.

Checks if a specific property/attribute requires the Advanced subscription tier.
This is useful for determining if a particular configuration option is available
with your current subscription level.

Args:
  - resource (string): Resource name (e.g., "http_loadbalancer", "app_firewall")
  - property (optional): Property name to check (e.g., "enable_malicious_user_detection")
    If omitted, returns all advanced features for the resource
  - response_format: 'markdown' or 'json'

Returns:
  - For specific property: Whether it requires Advanced subscription and matched feature
  - For no property: List of all Advanced-only features for the resource

Examples:
  - resource="http_loadbalancer", property="enable_malicious_user_detection" -> Advanced required
  - resource="http_loadbalancer", property="domains" -> Standard (no advanced requirement)
  - resource="http_loadbalancer" -> Lists all advanced features for HTTP LB`,
    inputSchema: GetPropertySubscriptionInfoSchema,
    annotations: {
      readOnlyHint: true,
      destructiveHint: false,
      idempotentHint: true,
      openWorldHint: false,
    },
  },
  async (params: GetPropertySubscriptionInfoInput) => {
    // Query for specific property
    if (params.property) {
      const info = getPropertySubscriptionInfo(params.resource, params.property);

      if (!info) {
        return {
          content: [{
            type: 'text',
            text: `No subscription information found for resource "${params.resource}". The resource may not exist or may not have tier metadata.`,
          }],
        };
      }

      const output = {
        resource: info.resourceName,
        property: info.propertyName,
        requires_advanced: info.requiresAdvanced,
        matched_feature: info.matchedFeature || null,
        resource_tier: info.resourceTier,
        service: info.service,
      };

      let textContent: string;
      if (params.response_format === ResponseFormat.MARKDOWN) {
        const lines = [
          `# Property Subscription Info`,
          '',
          `**Resource**: \`${info.resourceName}\``,
          `**Property**: \`${info.propertyName}\``,
          '',
          `**Requires Advanced Subscription**: ${info.requiresAdvanced ? '**Yes**' : 'No'}`,
        ];

        if (info.matchedFeature) {
          lines.push(`**Matched Feature**: \`${info.matchedFeature}\``);
          lines.push('');
          lines.push('> ⚠️ This property requires an Advanced subscription to use.');
        } else {
          lines.push('');
          lines.push('✓ This property is available with a Standard subscription.');
        }

        lines.push('');
        lines.push(`**Resource Tier**: ${info.resourceTier}`);
        lines.push(`**Service**: ${info.service}`);

        textContent = lines.join('\n');
      } else {
        textContent = JSON.stringify(output, null, 2);
      }

      return {
        content: [{ type: 'text', text: textContent }],
        structuredContent: output,
      };
    }

    // List all advanced features for the resource
    const resourceInfo = getResourceAdvancedProperties(params.resource);

    if (!resourceInfo) {
      return {
        content: [{
          type: 'text',
          text: `No subscription information found for resource "${params.resource}". The resource may not exist or may not have tier metadata.`,
        }],
      };
    }

    const output = {
      resource: resourceInfo.resourceName,
      resource_tier: resourceInfo.resourceTier,
      service: resourceInfo.service,
      advanced_features: resourceInfo.advancedFeatures,
      has_advanced_features: resourceInfo.advancedFeatures.length > 0,
    };

    let textContent: string;
    if (params.response_format === ResponseFormat.MARKDOWN) {
      const lines = [
        `# Advanced Features: ${resourceInfo.resourceName}`,
        '',
        `**Resource Tier**: ${resourceInfo.resourceTier}`,
        `**Service**: ${resourceInfo.service}`,
        '',
      ];

      if (resourceInfo.advancedFeatures.length > 0) {
        lines.push('## Properties Requiring Advanced Subscription');
        lines.push('');
        lines.push('The following features/properties require an Advanced subscription:');
        lines.push('');
        for (const feature of resourceInfo.advancedFeatures) {
          lines.push(`- \`${feature}\``);
        }
        lines.push('');
        lines.push('> Note: Property names may vary (e.g., `enable_malicious_user_detection` for feature `malicious_user_detection`)');
      } else {
        lines.push('✓ All properties of this resource are available with a Standard subscription.');
      }

      textContent = lines.join('\n');
    } else {
      textContent = JSON.stringify(output, null, 2);
    }

    return {
      content: [{ type: 'text', text: textContent }],
      structuredContent: output,
    };
  }
);

// =============================================================================
// ADDON SERVICE TOOLS
// =============================================================================

server.registerTool(
  'f5xc_terraform_addon_list_services',
  {
    title: 'List F5XC Addon Services',
    description: `List all available F5 Distributed Cloud addon services with their activation types and tier requirements.

Addon services are additional features that can be activated for your F5XC tenant,
such as Bot Defense, Client Side Defense, API Discovery, and more.

Args:
  - tier (optional): Filter by 'STANDARD', 'ADVANCED', or 'PREMIUM'
  - activation_type (optional): Filter by 'self' (user-activated) or 'managed' (requires sales contact)
  - response_format: 'markdown' or 'json'

Returns:
  List of addon services with name, tier requirement, activation type, and description.

Examples:
  - tier="ADVANCED" -> List all Advanced tier addon services
  - activation_type="self" -> List all self-activatable services`,
    inputSchema: ListAddonServicesSchema,
    annotations: {
      readOnlyHint: true,
      destructiveHint: false,
      idempotentHint: true,
      openWorldHint: false,
    },
  },
  async (params: ListAddonServicesInput) => {
    const result = listAddonServices(params.tier, params.activation_type);

    const output = {
      total: result.total,
      filters: {
        tier: params.tier || 'all',
        activation_type: params.activation_type || 'all',
      },
      services: result.services,
    };

    let textContent: string;
    if (params.response_format === ResponseFormat.MARKDOWN) {
      const lines = [
        '# F5XC Addon Services',
        '',
        `Total: ${result.total} services`,
        '',
      ];

      for (const service of result.services) {
        lines.push(`## ${service.displayName} (\`${service.name}\`)`);
        lines.push('');
        lines.push(service.description);
        lines.push('');
        lines.push(`- **Tier**: ${service.tier}`);
        lines.push(`- **Activation Type**: ${service.activationType}`);
        lines.push(`- **Category**: ${service.category}`);
        lines.push('');
      }

      textContent = lines.join('\n');
    } else {
      textContent = JSON.stringify(output, null, 2);
    }

    return {
      content: [{ type: 'text', text: textContent }],
      structuredContent: output,
    };
  }
);

server.registerTool(
  'f5xc_terraform_addon_check_activation',
  {
    title: 'Check Addon Activation Requirements',
    description: `Check activation requirements for a specific F5XC addon service.

Provides information about whether an addon service can be activated,
the required subscription tier, and step-by-step activation guidance.

Args:
  - addon_service (string): Name of the addon service (e.g., "bot_defense", "client_side_defense")
  - response_format: 'markdown' or 'json'

Returns:
  - Activation status (can activate, requirements)
  - Required subscription tier
  - Activation type (self, partial, managed)
  - Step-by-step activation process
  - Terraform example configuration

Examples:
  - addon_service="bot_defense" -> Bot Defense activation requirements
  - addon_service="api_discovery" -> API Discovery activation info`,
    inputSchema: CheckAddonActivationSchema,
    annotations: {
      readOnlyHint: true,
      destructiveHint: false,
      idempotentHint: true,
      openWorldHint: false,
    },
  },
  async (params: CheckAddonActivationInput) => {
    const result = checkAddonActivation(params.addon_service);

    if (!result) {
      return {
        content: [{
          type: 'text',
          text: `Addon service "${params.addon_service}" not found. Use f5xc_terraform_addon_list_services to see available services.`,
        }],
      };
    }

    const output = {
      addon_service: result.addonService,
      display_name: result.displayName,
      tier: result.tier,
      activation_type: result.activationType,
      can_activate: result.canActivate,
      steps: result.steps,
      terraform_example: result.terraformExample,
    };

    let textContent: string;
    if (params.response_format === ResponseFormat.MARKDOWN) {
      const lines = [
        `# Activation Requirements: ${result.displayName}`,
        '',
        `**Addon Service**: \`${result.addonService}\``,
        `**Required Tier**: ${result.tier}`,
        `**Activation Type**: ${result.activationType}`,
        `**Can Activate Directly**: ${result.canActivate ? 'Yes' : 'No (requires sales contact)'}`,
        '',
        '## Activation Steps',
        '',
      ];

      result.steps.forEach((step, index) => {
        lines.push(`${index + 1}. ${step}`);
      });

      lines.push('');
      lines.push('## Terraform Configuration');
      lines.push('');
      lines.push('```hcl');
      lines.push(result.terraformExample);
      lines.push('```');

      textContent = lines.join('\n');
    } else {
      textContent = JSON.stringify(output, null, 2);
    }

    return {
      content: [{ type: 'text', text: textContent }],
      structuredContent: output,
    };
  }
);

server.registerTool(
  'f5xc_terraform_addon_activation_workflow',
  {
    title: 'Get Addon Activation Workflow',
    description: `Get a detailed step-by-step activation workflow for an F5XC addon service.

Provides comprehensive guidance for activating an addon service including:
- Prerequisites and requirements
- Step-by-step instructions with Terraform snippets
- Complete Terraform configuration
- Estimated activation time
- Important notes and considerations

Args:
  - addon_service (string): Name of the addon service
  - activation_type (optional): Override workflow type ('self', 'partial', 'managed')
  - response_format: 'markdown' or 'json'

Returns:
  Complete activation workflow with Terraform configuration examples.

Examples:
  - addon_service="bot_defense" -> Self-activation workflow
  - addon_service="bot_defense", activation_type="partial" -> Partial managed workflow`,
    inputSchema: GetAddonWorkflowSchema,
    annotations: {
      readOnlyHint: true,
      destructiveHint: false,
      idempotentHint: true,
      openWorldHint: false,
    },
  },
  async (params: GetAddonWorkflowInput) => {
    const result = getAddonWorkflow(params.addon_service, params.activation_type);

    if (!result) {
      return {
        content: [{
          type: 'text',
          text: `Addon service "${params.addon_service}" not found or invalid activation type. Use f5xc_terraform_addon_list_services to see available services.`,
        }],
      };
    }

    const output = {
      addon_service: result.addonService,
      activation_type: result.activationType,
      description: result.description,
      prerequisites: result.prerequisites,
      steps: result.steps,
      terraform_config: result.terraformConfig,
      estimated_time: result.estimatedTime,
      notes: result.notes,
    };

    let textContent: string;
    if (params.response_format === ResponseFormat.MARKDOWN) {
      const lines = [
        `# Activation Workflow: ${result.addonService}`,
        '',
        result.description,
        '',
        `**Activation Type**: ${result.activationType}`,
        `**Estimated Time**: ${result.estimatedTime}`,
        '',
        '## Prerequisites',
        '',
      ];

      for (const prereq of result.prerequisites) {
        lines.push(`- ${prereq}`);
      }

      lines.push('');
      lines.push('## Step-by-Step Instructions');
      lines.push('');

      for (const step of result.steps) {
        lines.push(`### Step ${step.step}: ${step.action}`);
        lines.push('');
        lines.push(step.description);
        if (step.terraformSnippet) {
          lines.push('');
          lines.push('```hcl');
          lines.push(step.terraformSnippet);
          lines.push('```');
        }
        lines.push('');
      }

      lines.push('## Complete Terraform Configuration');
      lines.push('');
      lines.push('```hcl');
      lines.push(result.terraformConfig);
      lines.push('```');
      lines.push('');
      lines.push('## Notes');
      lines.push('');

      for (const note of result.notes) {
        lines.push(`- ${note}`);
      }

      textContent = lines.join('\n');
    } else {
      textContent = JSON.stringify(output, null, 2);
    }

    // Truncate if too long
    if (textContent.length > CHARACTER_LIMIT) {
      textContent = textContent.slice(0, CHARACTER_LIMIT) + '\n\n... (truncated)';
    }

    return {
      content: [{ type: 'text', text: textContent }],
      structuredContent: output,
    };
  }
);

// =============================================================================
// SERVER STARTUP
// =============================================================================

async function main() {
  const transport = new StdioServerTransport();
  await server.connect(transport);
  console.error('F5XC Terraform MCP server running on stdio');
}

main().catch(error => {
  console.error('Server error:', error);
  process.exit(1);
});
