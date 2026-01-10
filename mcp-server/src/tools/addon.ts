// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

/**
 * Consolidated Addon Tool Handler
 *
 * Replaces: f5xc_terraform_addon_list_services, f5xc_terraform_addon_check_activation,
 *           f5xc_terraform_addon_activation_workflow
 *
 * Tool: f5xc_terraform_addon
 */

import { AddonInput } from '../schemas/common.js';
import { ResponseFormat } from '../types.js';
import {
  listAddonServices,
  checkAddonActivation,
  getAddonWorkflow,
} from '../services/addons.js';

// =============================================================================
// HANDLER
// =============================================================================

/**
 * Handles the f5xc_terraform_addon tool invocation
 * Routes to appropriate operation based on input.operation
 */
export async function handleAddon(input: AddonInput): Promise<string> {
  const { operation, response_format } = input;

  switch (operation) {
    case 'list':
      return handleList(input, response_format);
    case 'check':
      return handleCheck(input, response_format);
    case 'workflow':
      return handleWorkflow(input, response_format);
    default:
      throw new Error(`Unknown operation: ${operation}`);
  }
}

// =============================================================================
// OPERATION HANDLERS
// =============================================================================

function handleList(input: AddonInput, format: ResponseFormat): string {
  const { tier, activation_type } = input;

  // Map tier from schema enum to service format
  const tierFilter = tier as 'STANDARD' | 'ADVANCED' | 'PREMIUM' | undefined;
  const activationFilter = activation_type === 'partial' ? 'managed' : (activation_type as 'self' | 'managed' | undefined);

  const result = listAddonServices(tierFilter, activationFilter);

  if (format === ResponseFormat.JSON) {
    return JSON.stringify(
      {
        operation: 'list',
        filters: {
          tier: tier || 'all',
          activation_type: activation_type || 'all',
        },
        total: result.total,
        services: result.services,
      },
      null,
      2,
    );
  }

  // Markdown format
  const lines: string[] = [
    '# F5XC Addon Services',
    '',
    tier ? `**Tier Filter**: ${tier}` : '',
    activation_type ? `**Activation Filter**: ${activation_type}` : '',
    `**Total**: ${result.total} service(s)`,
    '',
  ];

  if (result.total === 0) {
    lines.push('No addon services match the specified filters.');
    return lines.join('\n');
  }

  // Group by category
  const byCategory = new Map<string, typeof result.services>();
  for (const service of result.services) {
    const cat = service.category || 'other';
    if (!byCategory.has(cat)) {
      byCategory.set(cat, []);
    }
    byCategory.get(cat)!.push(service);
  }

  for (const [category, services] of byCategory) {
    lines.push(`## ${category.charAt(0).toUpperCase() + category.slice(1)}`);
    lines.push('');
    for (const service of services) {
      lines.push(`### ${service.displayName}`);
      lines.push(`- **Name**: \`${service.name}\``);
      lines.push(`- **Tier**: ${service.tier}`);
      lines.push(`- **Activation**: ${service.activationType}`);
      lines.push(`- ${service.description}`);
      lines.push('');
    }
  }

  return lines.join('\n');
}

function handleCheck(input: AddonInput, format: ResponseFormat): string {
  const { service_name } = input;

  if (!service_name) {
    throw new Error('service_name is required for check operation');
  }

  const result = checkAddonActivation(service_name);

  if (!result) {
    const error = {
      error: 'Addon service not found',
      service_name,
      suggestion: "Use operation: 'list' to see available addon services",
    };
    return format === ResponseFormat.JSON
      ? JSON.stringify(error, null, 2)
      : `# Error: Addon Service Not Found\n\n**Service**: ${service_name}\n\nUse \`operation: 'list'\` to see available addon services.`;
  }

  if (format === ResponseFormat.JSON) {
    return JSON.stringify(
      {
        operation: 'check',
        service_name: result.addonService,
        display_name: result.displayName,
        tier: result.tier,
        activation_type: result.activationType,
        can_activate: result.canActivate,
        steps: result.steps,
        terraform_example: result.terraformExample,
      },
      null,
      2,
    );
  }

  // Markdown format
  const lines: string[] = [
    `# Addon Activation Check: ${result.displayName}`,
    '',
    `**Service Name**: \`${result.addonService}\``,
    `**Tier Required**: ${result.tier}`,
    `**Activation Type**: ${result.activationType}`,
    `**Can Self-Activate**: ${result.canActivate ? '✅ Yes' : '❌ No (requires sales contact)'}`,
    '',
    '## Activation Steps',
    '',
  ];

  for (let i = 0; i < result.steps.length; i++) {
    lines.push(`${i + 1}. ${result.steps[i]}`);
  }

  lines.push('');
  lines.push('## Terraform Example');
  lines.push('');
  lines.push('```hcl');
  lines.push(result.terraformExample);
  lines.push('```');

  return lines.join('\n');
}

function handleWorkflow(input: AddonInput, format: ResponseFormat): string {
  const { service_name, activation_type } = input;

  if (!service_name) {
    throw new Error('service_name is required for workflow operation');
  }

  const result = getAddonWorkflow(service_name, activation_type);

  if (!result) {
    const error = {
      error: 'Addon service not found or workflow not available',
      service_name,
      suggestion: "Use operation: 'list' to see available addon services",
    };
    return format === ResponseFormat.JSON
      ? JSON.stringify(error, null, 2)
      : `# Error: Workflow Not Available\n\n**Service**: ${service_name}\n\nUse \`operation: 'list'\` to see available addon services.`;
  }

  if (format === ResponseFormat.JSON) {
    return JSON.stringify(
      {
        operation: 'workflow',
        service_name: result.addonService,
        activation_type: result.activationType,
        description: result.description,
        prerequisites: result.prerequisites,
        steps: result.steps,
        terraform_config: result.terraformConfig,
        estimated_time: result.estimatedTime,
        notes: result.notes,
      },
      null,
      2,
    );
  }

  // Markdown format
  const lines: string[] = [
    `# Activation Workflow: ${result.addonService}`,
    '',
    `**Activation Type**: ${result.activationType}`,
    `**Estimated Time**: ${result.estimatedTime}`,
    '',
    result.description,
    '',
    '## Prerequisites',
    '',
  ];

  for (const prereq of result.prerequisites) {
    lines.push(`- ${prereq}`);
  }

  lines.push('');
  lines.push('## Steps');
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

  if (result.notes.length > 0) {
    lines.push('## Notes');
    lines.push('');
    for (const note of result.notes) {
      lines.push(`> ${note}`);
    }
  }

  return lines.join('\n');
}

// =============================================================================
// TOOL DEFINITION
// =============================================================================

export const ADDON_TOOL_DEFINITION = {
  name: 'f5xc_terraform_addon',
  description: 'List, check activation, and get workflows for F5XC addon services',
};
