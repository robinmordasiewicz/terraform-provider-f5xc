// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

/**
 * Consolidated Subscription Tool Handler
 *
 * Replaces: f5xc_terraform_get_subscription_info, f5xc_terraform_get_property_subscription_info
 *
 * Tool: f5xc_terraform_subscription
 */

import { SubscriptionInput } from '../schemas/common.js';
import { ResponseFormat } from '../types.js';
import {
  getResourceSubscriptionInfo,
  getPropertySubscriptionInfo,
  getAdvancedTierResources,
  getSubscriptionSummary,
} from '../services/documentation.js';

// =============================================================================
// HANDLER
// =============================================================================

/**
 * Handles the f5xc_terraform_subscription tool invocation
 * Routes to appropriate operation based on input.operation
 */
export async function handleSubscription(input: SubscriptionInput): Promise<string> {
  const { operation, response_format } = input;

  switch (operation) {
    case 'resource':
      return handleResource(input, response_format);
    case 'property':
      return handleProperty(input, response_format);
    default:
      throw new Error(`Unknown operation: ${operation}`);
  }
}

// =============================================================================
// OPERATION HANDLERS
// =============================================================================

function handleResource(input: SubscriptionInput, format: ResponseFormat): string {
  const { resource_name, tier } = input;

  // If no resource specified, return summary or filtered list
  if (!resource_name) {
    return handleResourceSummary(tier, format);
  }

  const info = getResourceSubscriptionInfo(resource_name);

  if (!info) {
    const error = {
      error: 'Resource not found in subscription metadata',
      resource: resource_name,
      suggestion: 'Use operation: "resource" without resource_name to see all resources',
    };
    return format === ResponseFormat.JSON
      ? JSON.stringify(error, null, 2)
      : `# Error: Resource Not Found\n\n**Resource**: ${resource_name}\n\nUse \`operation: 'resource'\` without resource_name to see all resources.`;
  }

  if (format === ResponseFormat.JSON) {
    return JSON.stringify(
      {
        operation: 'resource',
        resource_name,
        tier: info.tier,
        service: info.service,
        advanced_features: info.advancedFeatures,
        requires_advanced: info.tier === 'ADVANCED',
      },
      null,
      2,
    );
  }

  // Markdown format
  const lines: string[] = [
    `# Subscription Info: ${resource_name}`,
    '',
    `**Tier**: ${info.tier}`,
    `**Service**: ${info.service}`,
    '',
  ];

  if (info.tier === 'ADVANCED') {
    lines.push('> ⚠️ This resource requires an **Advanced** subscription tier.');
    lines.push('');
  }

  if (info.advancedFeatures && info.advancedFeatures.length > 0) {
    lines.push('## Advanced Features');
    lines.push('');
    lines.push('The following features require Advanced tier:');
    for (const feature of info.advancedFeatures) {
      lines.push(`- ${feature}`);
    }
    lines.push('');
  }

  return lines.join('\n');
}

function handleResourceSummary(tier: string | undefined, format: ResponseFormat): string {
  const summary = getSubscriptionSummary();
  const advancedResources = getAdvancedTierResources();

  if (format === ResponseFormat.JSON) {
    return JSON.stringify(
      {
        operation: 'resource',
        filter: tier || 'all',
        summary: {
          total_resources: summary.totalResources,
          advanced_only: summary.advancedOnlyResources,
          with_advanced_features: summary.resourcesWithAdvancedFeatures,
        },
        advanced_resources:
          !tier || tier === 'ADVANCED'
            ? advancedResources.map((r) => ({
                name: r.name,
                type: r.type,
              }))
            : undefined,
      },
      null,
      2,
    );
  }

  // Markdown format
  const lines: string[] = [
    '# F5XC Subscription Tiers Overview',
    '',
    '## Summary',
    '',
    `- **Total Resources**: ${summary.totalResources}`,
    `- **Advanced Only**: ${summary.advancedOnlyResources} resources`,
    `- **With Advanced Features**: ${summary.resourcesWithAdvancedFeatures} resources`,
    '',
  ];

  if (!tier || tier === 'ADVANCED') {
    lines.push('## Advanced Tier Resources');
    lines.push('');
    lines.push('These resources require an Advanced subscription:');
    lines.push('');
    for (const resource of advancedResources.slice(0, 30)) {
      lines.push(`- **${resource.name}** (${resource.type})`);
    }
    if (advancedResources.length > 30) {
      lines.push(`- *... and ${advancedResources.length - 30} more*`);
    }
    lines.push('');
  }

  return lines.join('\n');
}

function handleProperty(input: SubscriptionInput, format: ResponseFormat): string {
  const { resource_name, property_path } = input;

  if (!resource_name) {
    throw new Error('resource_name is required for property operation');
  }

  // If no property specified, get all advanced properties for the resource
  if (!property_path) {
    return handleResourceProperties(resource_name, format);
  }

  const info = getPropertySubscriptionInfo(resource_name, property_path);

  if (!info) {
    const error = {
      error: 'Resource not found in subscription metadata',
      resource: resource_name,
      property: property_path,
    };
    return format === ResponseFormat.JSON
      ? JSON.stringify(error, null, 2)
      : `# Error: Resource Not Found\n\n**Resource**: ${resource_name}`;
  }

  if (format === ResponseFormat.JSON) {
    return JSON.stringify(
      {
        operation: 'property',
        resource_name: info.resourceName,
        property_path: info.propertyName,
        requires_advanced: info.requiresAdvanced,
        matched_feature: info.matchedFeature,
        resource_tier: info.resourceTier,
        service: info.service,
      },
      null,
      2,
    );
  }

  // Markdown format
  const lines: string[] = [
    `# Property Subscription Info`,
    '',
    `**Resource**: ${resource_name}`,
    `**Property**: ${property_path}`,
    '',
  ];

  if (info.requiresAdvanced) {
    lines.push('> ⚠️ This property requires an **Advanced** subscription tier.');
    lines.push('');
    lines.push(`**Feature**: ${info.matchedFeature || property_path}`);
  } else {
    lines.push('✅ This property is available on **Standard** tier.');
  }

  return lines.join('\n');
}

function handleResourceProperties(resourceName: string, format: ResponseFormat): string {
  const info = getResourceSubscriptionInfo(resourceName);

  if (!info) {
    const error = {
      error: 'Resource not found',
      resource: resourceName,
    };
    return format === ResponseFormat.JSON
      ? JSON.stringify(error, null, 2)
      : `# Error: Resource Not Found\n\n**Resource**: ${resourceName}`;
  }

  const advancedFeatures = info.advancedFeatures || [];

  if (format === ResponseFormat.JSON) {
    return JSON.stringify(
      {
        operation: 'property',
        resource_name: resourceName,
        resource_tier: info.tier,
        advanced_properties: advancedFeatures,
        advanced_count: advancedFeatures.length,
      },
      null,
      2,
    );
  }

  // Markdown format
  const lines: string[] = [
    `# Advanced Properties: ${resourceName}`,
    '',
    `**Resource Tier**: ${info.tier}`,
    `**Advanced Properties**: ${advancedFeatures.length}`,
    '',
  ];

  if (advancedFeatures.length === 0) {
    lines.push('No advanced-tier properties found for this resource.');
    lines.push('All properties are available on Standard tier.');
  } else {
    lines.push('## Properties Requiring Advanced Tier');
    lines.push('');
    for (const feature of advancedFeatures) {
      lines.push(`- ${feature}`);
    }
  }

  return lines.join('\n');
}

// =============================================================================
// TOOL DEFINITION
// =============================================================================

export const SUBSCRIPTION_TOOL_DEFINITION = {
  name: 'f5xc_terraform_subscription',
  description: 'Check F5XC resource subscription tiers and advanced feature requirements',
};
