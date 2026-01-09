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
import { DocsSchema, ApiSchema, SubscriptionSchema, AddonSchema, AuthSchema } from '../schemas/common.js';

// =============================================================================
// CRITICAL PROVIDER CONSTANTS
// =============================================================================

/**
 * The ONLY correct Terraform provider source for F5 Distributed Cloud.
 * AI assistants MUST use this source - never volterraedge/volterra (deprecated).
 */
const TERRAFORM_PROVIDER_BLOCK = `terraform {
  required_providers {
    f5xc = {
      source  = "robinmordasiewicz/f5xc"
      version = "~> 3.0"
    }
  }
}

provider "f5xc" {}`;

// =============================================================================
// TOOL METADATA
// =============================================================================

interface ToolInfo {
  name: string;
  description: string;
  category: 'docs' | 'api' | 'subscription' | 'addon' | 'auth' | 'summary';
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
  {
    name: 'f5xc_terraform_auth',
    description: 'Authentication status, profiles, validation, and Terraform env/provider config',
    category: 'auth',
    operations: ['status', 'list', 'switch', 'validate', 'terraform-env', 'terraform-block'],
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
  critical_provider_warning: {
    message: string;
    correct_source: string;
    deprecated_sources: string[];
    registry_url: string;
    terraform_block: string;
  };
  critical_auth_workflow: {
    message: string;
    step_1: string;
    step_2: string;
    step_3: string;
    common_error: string;
  };
  critical_syntax_warning: {
    message: string;
    common_mistake: string;
    correct_syntax: string;
    recommendation: string;
  };
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
    critical_provider_warning: {
      message: 'ALWAYS use robinmordasiewicz/f5xc - NEVER use volterraedge/volterra (deprecated legacy provider)',
      correct_source: 'robinmordasiewicz/f5xc',
      deprecated_sources: ['volterraedge/volterra', 'hashicorp/volterra'],
      registry_url: 'https://registry.terraform.io/providers/robinmordasiewicz/f5xc/latest',
      terraform_block: TERRAFORM_PROVIDER_BLOCK,
    },
    critical_auth_workflow: {
      message: 'MUST authenticate BEFORE running terraform plan/apply',
      step_1: 'Call f5xc_terraform_auth(operation="terraform-env") to get credentials',
      step_2: 'Set F5XC_API_URL and F5XC_API_TOKEN environment variables',
      step_3: 'Then run terraform init, plan, apply',
      common_error: 'Skipping authentication causes "Missing Authentication Configuration" error',
    },
    critical_syntax_warning: {
      message: 'This provider uses empty blocks {} for mutually exclusive options, NOT boolean values',
      common_mistake: 'no_tls = true, advertise_on_public_default_vip = true, round_robin = true',
      correct_syntax: 'no_tls {}, advertise_on_public_default_vip {}, round_robin {}',
      recommendation: 'Always query f5xc_terraform_metadata(operation="syntax", resource="...") BEFORE writing Terraform config',
    },
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
    // ==========================================================================
    // CRITICAL PROVIDER SOURCE WARNING - MUST READ FIRST
    // ==========================================================================
    '---',
    '',
    '## 🚨 CRITICAL: Provider Source (READ FIRST)',
    '',
    '**ALWAYS use `robinmordasiewicz/f5xc` - NEVER use deprecated providers!**',
    '',
    '| ✅ Correct | ❌ WRONG (Deprecated) |',
    '|------------|----------------------|',
    '| `robinmordasiewicz/f5xc` | ~~volterraedge/volterra~~ |',
    '| | ~~hashicorp/volterra~~ |',
    '',
    '> **Why?** The `volterraedge/volterra` provider was the legacy Volterra provider before F5 acquired',
    '> Volterra and rebranded it to F5 Distributed Cloud. That provider is deprecated and unmaintained.',
    '',
    '### Required Terraform Configuration:',
    '```hcl',
    'terraform {',
    '  required_providers {',
    '    f5xc = {',
    '      source  = "robinmordasiewicz/f5xc"',
    '      version = "~> 3.0"',
    '    }',
    '  }',
    '}',
    '',
    'provider "f5xc" {}',
    '```',
    '',
    // ==========================================================================
    // CRITICAL SYNTAX PATTERN WARNING
    // ==========================================================================
    '---',
    '',
    '## ⚠️ CRITICAL: Empty Block Syntax Pattern',
    '',
    'This provider uses **empty blocks** `{}` for mutually exclusive options, NOT boolean values.',
    '',
    '### Common AI Mistake (WRONG):',
    '```hcl',
    '# WRONG - These are NOT boolean attributes!',
    'no_tls = true                          # ERROR',
    'advertise_on_public_default_vip = true # ERROR',
    'round_robin = true                     # ERROR',
    'disable_api_definition = true          # ERROR',
    'no_challenge = true                    # ERROR',
    '```',
    '',
    '### Correct Syntax:',
    '```hcl',
    '# CORRECT - Use empty block syntax for OneOf choices',
    'no_tls {}',
    'advertise_on_public_default_vip {}',
    'round_robin {}',
    'disable_api_definition {}',
    'no_challenge {}',
    '```',
    '',
    '### Before Writing ANY Resource:',
    '1. Query `f5xc_terraform_metadata(operation="syntax", resource="<name>")` to get correct syntax',
    '2. Query `f5xc_terraform_metadata(operation="example", resource="<name>")` to get complete examples',
    '3. Use `f5xc_terraform_metadata(operation="validate", resource="<name>", config="...")` to check your config',
    '',
    // ==========================================================================
    // CRITICAL AUTHENTICATION WORKFLOW
    // ==========================================================================
    '---',
    '',
    '## 🔐 CRITICAL: Authentication Workflow',
    '',
    '**BEFORE running `terraform plan` or `terraform apply`, you MUST authenticate:**',
    '',
    '### Step 1: Get Authentication Credentials',
    '```',
    'f5xc_terraform_auth(operation="terraform-env", output_type="shell")',
    '```',
    '',
    '### Step 2: Set Environment Variables',
    'Copy and paste the export commands from Step 1:',
    '```bash',
    'export F5XC_API_URL="https://your-tenant.console.ves.volterra.io/api"',
    'export F5XC_API_TOKEN="your-api-token"',
    '```',
    '',
    '### Step 3: Run Terraform',
    '```bash',
    'terraform init',
    'terraform plan',
    '```',
    '',
    '**⚠️ Without authentication, terraform will fail with "Missing Authentication Configuration".**',
    '',
    '---',
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
    case 'f5xc_terraform_auth':
      return schemaToJson(AuthSchema);
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
