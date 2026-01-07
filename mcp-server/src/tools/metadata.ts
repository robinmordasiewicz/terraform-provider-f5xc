/**
 * Consolidated Metadata Tool Handler
 *
 * Tool: f5xc_terraform_metadata
 *
 * Provides access to resource metadata for deterministic AI configuration generation.
 * Includes OneOf groups, validation patterns, defaults, and attribute information.
 */

import { MetadataInput } from '../schemas/common.js';
import { ResponseFormat } from '../types.js';
import {
  getResourceMetadata,
  getResourceOneOfGroups,
  getAttributesWithDefaults,
  getAttributesWithEnums,
  getRequiredAttributes,
  getRequiresReplaceAttributes,
  getValidationPattern,
  getAllValidationPatterns,
  listResourcesWithMetadata,
  getMetadataSummary,
  isMetadataAvailable,
  getResourceTier,
  getResourcesByTier,
  getResourceImportFormat,
  getResourceDependencies,
  computeCreationOrder,
  getErrorCodeInfo,
  findErrorByPattern,
  getAllErrorCodes,
  getDiagnosticTips,
  getEnhancedMetadataSummary,
} from '../services/metadata.js';

// =============================================================================
// HANDLER
// =============================================================================

/**
 * Handles the f5xc_terraform_metadata tool invocation
 * Routes to appropriate operation based on input.operation
 */
export async function handleMetadata(input: MetadataInput): Promise<string> {
  const { operation, response_format } = input;

  // Check if metadata is available
  if (!isMetadataAvailable()) {
    return formatError(
      'Metadata not available',
      'Metadata files have not been generated. Run the schema generator to create metadata.',
      response_format,
    );
  }

  switch (operation) {
    case 'oneof':
      return handleOneOf(input, response_format);
    case 'validation':
      return handleValidation(input, response_format);
    case 'defaults':
      return handleDefaults(input, response_format);
    case 'enums':
      return handleEnums(input, response_format);
    case 'attribute':
      return handleAttribute(input, response_format);
    case 'requires_replace':
      return handleRequiresReplace(input, response_format);
    case 'tier':
      return handleTier(input, response_format);
    case 'dependencies':
      return handleDependencies(input, response_format);
    case 'troubleshoot':
      return handleTroubleshoot(input, response_format);
    case 'summary':
      return handleSummary(response_format);
    default:
      throw new Error(`Unknown operation: ${operation}`);
  }
}

// =============================================================================
// OPERATION HANDLERS
// =============================================================================

function handleOneOf(input: MetadataInput, format: ResponseFormat): string {
  const { resource } = input;

  if (!resource) {
    return formatError(
      'Missing parameter',
      'resource parameter is required for oneof operation',
      format,
    );
  }

  const oneOfGroups = getResourceOneOfGroups(resource);

  if (Object.keys(oneOfGroups).length === 0) {
    return formatNoData(
      'No OneOf groups found',
      `Resource '${resource}' has no mutually exclusive field groups.`,
      format,
    );
  }

  if (format === ResponseFormat.JSON) {
    return JSON.stringify(
      {
        operation: 'oneof',
        resource,
        groups: oneOfGroups,
        total_groups: Object.keys(oneOfGroups).length,
      },
      null,
      2,
    );
  }

  // Markdown format
  const lines: string[] = [
    `# OneOf Groups: ${resource}`,
    '',
    'These field groups are **mutually exclusive** - only one field from each group can be set.',
    '',
  ];

  for (const [groupName, group] of Object.entries(oneOfGroups)) {
    lines.push(`## ${groupName}`);
    lines.push('');
    lines.push(`**Fields**: ${group.fields.join(', ')}`);
    if (group.default) {
      lines.push(`**Recommended Default**: \`${group.default}\``);
    }
    if (group.description) {
      lines.push(`**Description**: ${group.description}`);
    }
    lines.push('');
  }

  return lines.join('\n');
}

function handleValidation(input: MetadataInput, format: ResponseFormat): string {
  const { pattern } = input;

  // If pattern specified, return that specific pattern
  if (pattern) {
    const validationPattern = getValidationPattern(pattern);

    if (!validationPattern) {
      return formatError(
        'Pattern not found',
        `Validation pattern '${pattern}' not found. Available: name, domain, port, namespace`,
        format,
      );
    }

    if (format === ResponseFormat.JSON) {
      return JSON.stringify(
        {
          operation: 'validation',
          pattern,
          ...validationPattern,
        },
        null,
        2,
      );
    }

    const lines: string[] = [
      `# Validation Pattern: ${pattern}`,
      '',
      `**Type**: ${validationPattern.type}`,
      '',
    ];

    if (validationPattern.pattern) {
      lines.push(`**Pattern**: \`${validationPattern.pattern}\``);
    }
    if (validationPattern.min !== undefined) {
      lines.push(`**Min**: ${validationPattern.min}`);
    }
    if (validationPattern.max !== undefined) {
      lines.push(`**Max**: ${validationPattern.max}`);
    }
    lines.push('');
    lines.push(`**Description**: ${validationPattern.description}`);
    lines.push('');

    if (validationPattern.examples && validationPattern.examples.length > 0) {
      lines.push('**Valid Examples**:');
      for (const ex of validationPattern.examples) {
        lines.push(`- \`${ex}\``);
      }
      lines.push('');
    }

    if (validationPattern.invalid && validationPattern.invalid.length > 0) {
      lines.push('**Invalid Examples**:');
      for (const ex of validationPattern.invalid) {
        lines.push(`- \`${ex}\``);
      }
    }

    return lines.join('\n');
  }

  // Return all patterns
  const allPatterns = getAllValidationPatterns();

  if (format === ResponseFormat.JSON) {
    return JSON.stringify(
      {
        operation: 'validation',
        patterns: allPatterns,
        total: Object.keys(allPatterns).length,
      },
      null,
      2,
    );
  }

  const lines: string[] = ['# Validation Patterns', '', '| Pattern | Type | Description |', '|---------|------|-------------|'];

  for (const [name, pat] of Object.entries(allPatterns)) {
    lines.push(`| ${name} | ${pat.type} | ${pat.description.slice(0, 50)}... |`);
  }

  lines.push('');
  lines.push('Use `pattern` parameter to get full details for a specific pattern.');

  return lines.join('\n');
}

function handleDefaults(input: MetadataInput, format: ResponseFormat): string {
  const { resource } = input;

  if (!resource) {
    return formatError(
      'Missing parameter',
      'resource parameter is required for defaults operation',
      format,
    );
  }

  const defaults = getAttributesWithDefaults(resource);

  if (Object.keys(defaults).length === 0) {
    return formatNoData(
      'No defaults found',
      `Resource '${resource}' has no attributes with default values.`,
      format,
    );
  }

  if (format === ResponseFormat.JSON) {
    return JSON.stringify(
      {
        operation: 'defaults',
        resource,
        defaults,
        total: Object.keys(defaults).length,
      },
      null,
      2,
    );
  }

  const lines: string[] = [
    `# Default Values: ${resource}`,
    '',
    'These attributes have default values when not specified:',
    '',
    '| Attribute | Default Value |',
    '|-----------|---------------|',
  ];

  for (const [name, value] of Object.entries(defaults)) {
    const displayValue = typeof value === 'string' ? `"${value}"` : JSON.stringify(value);
    lines.push(`| ${name} | ${displayValue} |`);
  }

  return lines.join('\n');
}

function handleEnums(input: MetadataInput, format: ResponseFormat): string {
  const { resource } = input;

  if (!resource) {
    return formatError(
      'Missing parameter',
      'resource parameter is required for enums operation',
      format,
    );
  }

  const enums = getAttributesWithEnums(resource);

  if (Object.keys(enums).length === 0) {
    return formatNoData(
      'No enums found',
      `Resource '${resource}' has no attributes with enum constraints.`,
      format,
    );
  }

  if (format === ResponseFormat.JSON) {
    return JSON.stringify(
      {
        operation: 'enums',
        resource,
        enums,
        total: Object.keys(enums).length,
      },
      null,
      2,
    );
  }

  const lines: string[] = [
    `# Enum Constraints: ${resource}`,
    '',
    'These attributes only accept specific values:',
    '',
  ];

  for (const [name, values] of Object.entries(enums)) {
    lines.push(`## ${name}`);
    lines.push('');
    lines.push(`**Valid values**: ${values.map(v => `\`${v}\``).join(', ')}`);
    lines.push('');
  }

  return lines.join('\n');
}

function handleAttribute(input: MetadataInput, format: ResponseFormat): string {
  const { resource, attribute } = input;

  if (!resource || !attribute) {
    return formatError(
      'Missing parameter',
      'Both resource and attribute parameters are required for attribute operation',
      format,
    );
  }

  const resourceMeta = getResourceMetadata(resource);

  if (!resourceMeta) {
    return formatError(
      'Resource not found',
      `Resource '${resource}' not found in metadata`,
      format,
    );
  }

  const attrMeta = resourceMeta.attributes[attribute];

  if (!attrMeta) {
    const availableAttrs = Object.keys(resourceMeta.attributes).slice(0, 10);
    return formatError(
      'Attribute not found',
      `Attribute '${attribute}' not found in ${resource}. Available: ${availableAttrs.join(', ')}...`,
      format,
    );
  }

  if (format === ResponseFormat.JSON) {
    return JSON.stringify(
      {
        operation: 'attribute',
        resource,
        attribute,
        ...attrMeta,
      },
      null,
      2,
    );
  }

  const lines: string[] = [
    `# Attribute: ${resource}.${attribute}`,
    '',
    `**Type**: ${attrMeta.type}`,
    `**Required**: ${attrMeta.required ? 'Yes' : 'No'}`,
  ];

  if (attrMeta.computed) lines.push(`**Computed**: Yes`);
  if (attrMeta.sensitive) lines.push(`**Sensitive**: Yes`);
  if (attrMeta.is_block) lines.push(`**Block Type**: Yes`);
  if (attrMeta.plan_modifier) lines.push(`**Plan Modifier**: ${attrMeta.plan_modifier}`);
  if (attrMeta.validation) lines.push(`**Validation**: ${attrMeta.validation}`);
  if (attrMeta.oneof_group) lines.push(`**OneOf Group**: ${attrMeta.oneof_group}`);

  lines.push('');
  lines.push(`**Description**: ${attrMeta.description}`);

  if (attrMeta.enum && attrMeta.enum.length > 0) {
    lines.push('');
    lines.push(`**Valid Values**: ${attrMeta.enum.map(v => `\`${v}\``).join(', ')}`);
  }

  if (attrMeta.default !== undefined) {
    lines.push('');
    lines.push(`**Default**: ${JSON.stringify(attrMeta.default)}`);
  }

  return lines.join('\n');
}

function handleRequiresReplace(input: MetadataInput, format: ResponseFormat): string {
  const { resource } = input;

  if (!resource) {
    return formatError(
      'Missing parameter',
      'resource parameter is required for requires_replace operation',
      format,
    );
  }

  const attrs = getRequiresReplaceAttributes(resource);
  const required = getRequiredAttributes(resource);

  if (format === ResponseFormat.JSON) {
    return JSON.stringify(
      {
        operation: 'requires_replace',
        resource,
        requires_replace_attributes: attrs,
        required_attributes: required,
      },
      null,
      2,
    );
  }

  const lines: string[] = [
    `# Important Attributes: ${resource}`,
    '',
    '## Requires Replacement',
    '',
    'Changing these attributes will destroy and recreate the resource:',
    '',
  ];

  if (attrs.length > 0) {
    for (const attr of attrs) {
      lines.push(`- \`${attr}\``);
    }
  } else {
    lines.push('*No attributes require replacement.*');
  }

  lines.push('');
  lines.push('## Required Attributes');
  lines.push('');
  lines.push('These attributes must be specified:');
  lines.push('');

  for (const attr of required) {
    lines.push(`- \`${attr}\``);
  }

  return lines.join('\n');
}

function handleTier(input: MetadataInput, format: ResponseFormat): string {
  const { resource, tier } = input;

  // If resource specified, return tier for that resource
  if (resource) {
    const resourceTier = getResourceTier(resource);
    const importFormat = getResourceImportFormat(resource);

    if (!resourceTier) {
      return formatNoData(
        'Tier not found',
        `Resource '${resource}' does not have tier information in metadata.`,
        format,
      );
    }

    if (format === ResponseFormat.JSON) {
      return JSON.stringify(
        {
          operation: 'tier',
          resource,
          tier: resourceTier,
          import_format: importFormat,
        },
        null,
        2,
      );
    }

    const lines: string[] = [
      `# Subscription Tier: ${resource}`,
      '',
      `**Tier**: ${resourceTier}`,
      `**Import Format**: \`${importFormat || 'unknown'}\``,
      '',
      '## Tier Descriptions',
      '',
      '- **STANDARD**: Included with base F5 XC subscription',
      '- **ADVANCED**: Requires Advanced tier subscription',
      '- **PREMIUM**: Requires Premium tier subscription',
    ];

    return lines.join('\n');
  }

  // If tier specified, return all resources with that tier
  if (tier) {
    const resources = getResourcesByTier(tier);

    if (resources.length === 0) {
      return formatNoData(
        'No resources found',
        `No resources found requiring tier '${tier}'.`,
        format,
      );
    }

    if (format === ResponseFormat.JSON) {
      return JSON.stringify(
        {
          operation: 'tier',
          tier,
          resources,
          total: resources.length,
        },
        null,
        2,
      );
    }

    const lines: string[] = [
      `# Resources Requiring ${tier} Tier`,
      '',
      `Total: ${resources.length} resources`,
      '',
      '| Resource |',
      '|----------|',
    ];

    for (const res of resources) {
      lines.push(`| ${res} |`);
    }

    return lines.join('\n');
  }

  // No resource or tier specified - return summary of tiers
  const summary = getEnhancedMetadataSummary();

  if (format === ResponseFormat.JSON) {
    return JSON.stringify(
      {
        operation: 'tier',
        tier_distribution: summary.resourcesWithTier,
        total_resources: summary.resourceCount,
      },
      null,
      2,
    );
  }

  const lines: string[] = [
    '# Subscription Tier Distribution',
    '',
    '| Tier | Resource Count |',
    '|------|----------------|',
  ];

  for (const [tierName, count] of Object.entries(summary.resourcesWithTier)) {
    lines.push(`| ${tierName} | ${count} |`);
  }

  lines.push('');
  lines.push(`Total resources: ${summary.resourceCount}`);
  lines.push('');
  lines.push('Use `resource` parameter to check a specific resource, or `tier` parameter to list resources by tier.');

  return lines.join('\n');
}

function handleDependencies(input: MetadataInput, format: ResponseFormat): string {
  const { resource } = input;

  if (!resource) {
    return formatError(
      'Missing parameter',
      'resource parameter is required for dependencies operation',
      format,
    );
  }

  const dependencies = getResourceDependencies(resource);
  const creationOrder = computeCreationOrder(resource);
  const importFormat = getResourceImportFormat(resource);

  if (format === ResponseFormat.JSON) {
    return JSON.stringify(
      {
        operation: 'dependencies',
        resource,
        import_format: importFormat,
        dependencies: dependencies || { references: [], referenced_by: [] },
        creation_order: creationOrder,
      },
      null,
      2,
    );
  }

  const lines: string[] = [
    `# Dependencies: ${resource}`,
    '',
    `**Import Format**: \`${importFormat || 'unknown'}\``,
    '',
  ];

  if (dependencies?.references && dependencies.references.length > 0) {
    lines.push('## Required Dependencies');
    lines.push('');
    lines.push('This resource depends on:');
    lines.push('');
    for (const dep of dependencies.references) {
      lines.push(`- \`${dep}\``);
    }
    lines.push('');
  } else {
    lines.push('## Required Dependencies');
    lines.push('');
    lines.push('*No dependencies detected.*');
    lines.push('');
  }

  if (dependencies?.referenced_by && dependencies.referenced_by.length > 0) {
    lines.push('## Referenced By');
    lines.push('');
    lines.push('This resource is used by:');
    lines.push('');
    for (const ref of dependencies.referenced_by) {
      lines.push(`- \`${ref}\``);
    }
    lines.push('');
  }

  lines.push('## Creation Order');
  lines.push('');
  lines.push('Create resources in this order:');
  lines.push('');
  for (let i = 0; i < creationOrder.length; i++) {
    lines.push(`${i + 1}. \`${creationOrder[i]}\``);
  }

  return lines.join('\n');
}

function handleTroubleshoot(input: MetadataInput, format: ResponseFormat): string {
  const { error_code, error_message } = input;

  // If error_code specified, return that error's info
  if (error_code) {
    const errorInfo = getErrorCodeInfo(error_code);

    if (!errorInfo) {
      const availableCodes = Object.keys(getAllErrorCodes());
      return formatError(
        'Error code not found',
        `Error code '${error_code}' not found. Available: ${availableCodes.join(', ')}`,
        format,
      );
    }

    if (format === ResponseFormat.JSON) {
      return JSON.stringify(
        {
          operation: 'troubleshoot',
          error_code,
          ...errorInfo,
        },
        null,
        2,
      );
    }

    const lines: string[] = [
      `# Error: ${error_code}`,
      '',
      `**HTTP Status**: ${errorInfo.status}`,
      '',
      '## Common Causes',
      '',
    ];

    for (const cause of errorInfo.causes) {
      lines.push(`- ${cause}`);
    }

    lines.push('');
    lines.push('## Remediation Steps');
    lines.push('');

    for (const step of errorInfo.remediation) {
      lines.push(`- ${step}`);
    }

    lines.push('');
    lines.push(`**Affected Operations**: ${errorInfo.operations.join(', ')}`);

    return lines.join('\n');
  }

  // If error_message specified, try pattern matching
  if (error_message) {
    const match = findErrorByPattern(error_message);

    if (!match) {
      return formatNoData(
        'No pattern match',
        `Could not find a matching error pattern for: "${error_message.slice(0, 100)}..."`,
        format,
      );
    }

    if (format === ResponseFormat.JSON) {
      return JSON.stringify(
        {
          operation: 'troubleshoot',
          matched_pattern: match.name,
          ...match.info,
        },
        null,
        2,
      );
    }

    const lines: string[] = [
      `# Matched Pattern: ${match.name}`,
      '',
      `**Pattern**: \`${match.info.pattern}\``,
      '',
      `**Cause**: ${match.info.cause}`,
      '',
      `**Remediation**: ${match.info.remediation}`,
    ];

    return lines.join('\n');
  }

  // No error specified - return all error codes and tips
  const allErrors = getAllErrorCodes();
  const tips = getDiagnosticTips();

  if (format === ResponseFormat.JSON) {
    return JSON.stringify(
      {
        operation: 'troubleshoot',
        error_codes: allErrors,
        diagnostic_tips: tips,
        total_codes: Object.keys(allErrors).length,
      },
      null,
      2,
    );
  }

  const lines: string[] = [
    '# Error Troubleshooting Guide',
    '',
    '## Error Codes',
    '',
    '| Code | Status | Primary Cause |',
    '|------|--------|---------------|',
  ];

  for (const [code, info] of Object.entries(allErrors)) {
    lines.push(`| ${code} | ${info.status} | ${info.causes[0]} |`);
  }

  lines.push('');
  lines.push('## Diagnostic Tips');
  lines.push('');

  for (const [key, tip] of Object.entries(tips)) {
    lines.push(`- **${key}**: ${tip}`);
  }

  lines.push('');
  lines.push('Use `error_code` parameter for detailed info, or `error_message` for pattern matching.');

  return lines.join('\n');
}

function handleSummary(format: ResponseFormat): string {
  const summary = getEnhancedMetadataSummary();
  const resources = listResourcesWithMetadata();

  if (format === ResponseFormat.JSON) {
    return JSON.stringify(
      {
        operation: 'summary',
        ...summary,
        sample_resources: resources.slice(0, 20),
      },
      null,
      2,
    );
  }

  const lines: string[] = [
    '# Metadata Summary',
    '',
    '| Metric | Value |',
    '|--------|-------|',
    `| Resources with Metadata | ${summary.resourceCount} |`,
    `| Resources with OneOf Groups | ${summary.resourcesWithOneOf} |`,
    `| Resources with Dependencies | ${summary.resourcesWithDependencies} |`,
    `| Validation Patterns | ${summary.validationPatternCount} |`,
    `| Error Codes Indexed | ${summary.errorCodeCount} |`,
    `| Generated At | ${summary.generatedAt || 'Unknown'} |`,
    '',
    '## Tier Distribution',
    '',
    '| Tier | Count |',
    '|------|-------|',
  ];

  for (const [tierName, count] of Object.entries(summary.resourcesWithTier)) {
    lines.push(`| ${tierName} | ${count} |`);
  }

  lines.push('');
  lines.push('## Available Operations');
  lines.push('');
  lines.push('| Operation | Description |');
  lines.push('|-----------|-------------|');
  lines.push('| `oneof` | Get mutually exclusive field groups for a resource |');
  lines.push('| `validation` | Get validation patterns (regex/range rules) |');
  lines.push('| `defaults` | Get default values for a resource |');
  lines.push('| `enums` | Get enum constraints for a resource |');
  lines.push('| `attribute` | Get full metadata for a specific attribute |');
  lines.push('| `requires_replace` | Get attributes that trigger resource replacement |');
  lines.push('| `tier` | Get subscription tier requirements for resources |');
  lines.push('| `dependencies` | Get resource dependencies and creation order |');
  lines.push('| `troubleshoot` | Get error codes, causes, and remediation steps |');
  lines.push('| `summary` | This summary |');

  return lines.join('\n');
}

// =============================================================================
// HELPERS
// =============================================================================

function formatError(title: string, message: string, format: ResponseFormat): string {
  if (format === ResponseFormat.JSON) {
    return JSON.stringify({ error: title, message }, null, 2);
  }
  return `# Error: ${title}\n\n${message}`;
}

function formatNoData(title: string, message: string, format: ResponseFormat): string {
  if (format === ResponseFormat.JSON) {
    return JSON.stringify({ info: title, message }, null, 2);
  }
  return `# ${title}\n\n${message}`;
}

// =============================================================================
// TOOL DEFINITION
// =============================================================================

export const METADATA_TOOL_DEFINITION = {
  name: 'f5xc_terraform_metadata',
  description: 'Query resource metadata for deterministic Terraform configuration generation. Operations: oneof (mutually exclusive fields), validation (regex/range patterns), defaults, enums, attribute (full metadata), requires_replace, tier (subscription requirements), dependencies (creation order), troubleshoot (error remediation), summary.',
};
