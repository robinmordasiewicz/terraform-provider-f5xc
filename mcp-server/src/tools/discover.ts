/**
 * Discovery Meta-Tool Handler
 *
 * Enables lazy loading of tool schemas for token efficiency.
 * AI assistants can discover available tools without loading full schemas.
 *
 * Tool: f5xc_terraform_discover
 */

import { DiscoverInput } from '../schemas/common.js';
import { ResponseFormat } from '../types.js';
import { DocsSchema, ApiSchema, SubscriptionSchema, AddonSchema } from '../schemas/common.js';

// =============================================================================
// TOOL METADATA
// =============================================================================

interface ToolInfo {
  name: string;
  description: string;
  category: 'docs' | 'api' | 'subscription' | 'addon' | 'summary';
  operations?: string[];
}

const TOOL_REGISTRY: ToolInfo[] = [
  {
    name: 'f5xc_terraform_docs',
    description: 'Search, get, list F5XC Terraform docs',
    category: 'docs',
    operations: ['search', 'get', 'list'],
  },
  {
    name: 'f5xc_terraform_api',
    description: 'Query 270+ OpenAPI specs and schemas',
    category: 'api',
    operations: ['search', 'get', 'find_endpoints', 'get_definition', 'list_definitions'],
  },
  {
    name: 'f5xc_terraform_subscription',
    description: 'Check resource subscription tiers',
    category: 'subscription',
    operations: ['resource', 'property'],
  },
  {
    name: 'f5xc_terraform_addon',
    description: 'List/check/activate addon services',
    category: 'addon',
    operations: ['list', 'check', 'workflow'],
  },
  {
    name: 'f5xc_terraform_get_summary',
    description: 'Get provider documentation summary',
    category: 'summary',
  },
];

// =============================================================================
// HANDLER
// =============================================================================

/**
 * Handles the f5xc_terraform_discover tool invocation
 */
export async function handleDiscover(input: DiscoverInput): Promise<string> {
  const { category, verbose, response_format } = input;

  // Filter tools by category
  const tools = category === 'all'
    ? TOOL_REGISTRY
    : TOOL_REGISTRY.filter((t) => t.category === category);

  // Build response
  const response = buildResponse(tools, verbose, response_format);

  return response;
}

function buildResponse(
  tools: ToolInfo[],
  verbose: boolean,
  format: ResponseFormat,
): string {
  if (format === ResponseFormat.JSON) {
    return JSON.stringify(buildJsonResponse(tools, verbose), null, 2);
  }

  return buildMarkdownResponse(tools, verbose);
}

interface DiscoverResponse {
  provider: {
    name: string;
    registry_url: string;
    github_url: string;
  };
  tools: Array<{
    name: string;
    description: string;
    operations?: string[];
    schema?: object;
  }>;
  total_tools: number;
  token_estimate: string;
  usage_hint: string;
}

function buildJsonResponse(tools: ToolInfo[], verbose: boolean): DiscoverResponse {
  return {
    provider: {
      name: 'robinmordasiewicz/f5xc',
      registry_url: 'https://registry.terraform.io/providers/robinmordasiewicz/f5xc/latest',
      github_url: 'https://github.com/robinmordasiewicz/terraform-provider-f5xc',
    },
    tools: tools.map((t) => {
      const tool: DiscoverResponse['tools'][0] = {
        name: t.name,
        description: t.description,
      };

      if (verbose) {
        tool.operations = t.operations;
        tool.schema = getSchemaForTool(t.name);
      }

      return tool;
    }),
    total_tools: tools.length,
    token_estimate: verbose ? '~2,000 tokens' : '~150 tokens',
    usage_hint: verbose
      ? 'Full schemas included for reference'
      : 'Use verbose=true for full schemas',
  };
}

function buildMarkdownResponse(tools: ToolInfo[], verbose: boolean): string {
  const lines: string[] = [
    '# F5XC Terraform MCP Tools',
    '',
    '**Provider**: `robinmordasiewicz/f5xc`',
    '**Registry**: https://registry.terraform.io/providers/robinmordasiewicz/f5xc/latest',
    '',
    `**Available Tools**: ${tools.length}`,
    `**Token Estimate**: ${verbose ? '~2,000' : '~150'} tokens`,
    '',
    '## Tools',
    '',
  ];

  for (const tool of tools) {
    lines.push(`### ${tool.name}`);
    lines.push(tool.description);

    if (tool.operations) {
      lines.push(`- **Operations**: ${tool.operations.join(', ')}`);
    }

    if (verbose) {
      lines.push('');
      lines.push('**Schema**:');
      lines.push('```json');
      lines.push(JSON.stringify(getSchemaForTool(tool.name), null, 2));
      lines.push('```');
    }

    lines.push('');
  }

  if (!verbose) {
    lines.push('---');
    lines.push('*Use `verbose=true` to see full schemas for each tool.*');
  }

  return lines.join('\n');
}

function getSchemaForTool(toolName: string): object {
  switch (toolName) {
    case 'f5xc_terraform_docs':
      return schemaToJson(DocsSchema);
    case 'f5xc_terraform_api':
      return schemaToJson(ApiSchema);
    case 'f5xc_terraform_subscription':
      return schemaToJson(SubscriptionSchema);
    case 'f5xc_terraform_addon':
      return schemaToJson(AddonSchema);
    case 'f5xc_terraform_get_summary':
      return { response_format: { type: 'string', enum: ['markdown', 'json'] } };
    default:
      return {};
  }
}

/**
 * Converts a Zod schema to a simplified JSON representation
 * for documentation purposes
 */
function schemaToJson(schema: { shape?: object }): object {
  if (!schema.shape) {
    return { type: 'object' };
  }

  const result: Record<string, object> = {};

  for (const [key, value] of Object.entries(schema.shape)) {
    result[key] = describeZodType(value);
  }

  return result;
}

function describeZodType(zodType: unknown): object {
  // Extract type info from Zod schema
  const desc = (zodType as { description?: string }).description;
  const typeDef = (zodType as { _def?: { typeName?: string; innerType?: unknown; values?: string[] } })._def;

  if (!typeDef) {
    return { type: 'unknown' };
  }

  const result: Record<string, unknown> = {};

  switch (typeDef.typeName) {
    case 'ZodString':
      result.type = 'string';
      break;
    case 'ZodNumber':
      result.type = 'number';
      break;
    case 'ZodBoolean':
      result.type = 'boolean';
      break;
    case 'ZodEnum':
      result.type = 'string';
      result.enum = typeDef.values;
      break;
    case 'ZodNativeEnum':
      result.type = 'string';
      result.enum = ['markdown', 'json'];
      break;
    case 'ZodOptional':
      return { ...describeZodType(typeDef.innerType), optional: true };
    case 'ZodDefault':
      return { ...describeZodType(typeDef.innerType) };
    default:
      result.type = 'unknown';
  }

  if (desc) {
    result.description = desc;
  }

  return result;
}

// =============================================================================
// TOOL DEFINITION
// =============================================================================

export const DISCOVER_TOOL_DEFINITION = {
  name: 'f5xc_terraform_discover',
  description: 'Discover available F5XC Terraform MCP tools with optional schema details',
};
