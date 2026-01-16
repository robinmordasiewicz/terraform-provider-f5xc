// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

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
  getMetadataSummary as _getMetadataSummary,
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
  getAttributeSyntaxGuidance,
  generateTerraformSyntaxGuide,
  validateConfigurationSyntax as _validateConfigurationSyntax,
  getOneOfGroupsQuickReference as _getOneOfGroupsQuickReference,
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
    case 'syntax':
      return handleSyntax(input, response_format);
    case 'validate':
      return handleValidate(input, response_format);
    case 'example':
      return handleExample(input, response_format);
    case 'mistakes':
      return handleMistakes(input, response_format);
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
    lines.push('**Syntax Examples**:');
    lines.push('');
    for (const field of group.fields) {
      const guidance = getAttributeSyntaxGuidance(resource, field);
      if (guidance) {
        lines.push(`### \`${field}\``);
        lines.push('');
        lines.push('```hcl');
        lines.push(guidance.example);
        lines.push('```');
        lines.push('');
      }
    }
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

  // Add syntax guidance for this attribute
  const syntaxGuidance = getAttributeSyntaxGuidance(resource, attribute);
  if (syntaxGuidance) {
    lines.push('');
    lines.push('## Terraform Syntax');
    lines.push('');
    lines.push('**Correct Syntax**: `' + syntaxGuidance.correctSyntax + '`');
    lines.push('');
    lines.push('**Example**:');
    lines.push('```hcl');
    lines.push(syntaxGuidance.example);
    lines.push('```');
    lines.push('');

    if (syntaxGuidance.isBlock) {
      lines.push('**Important**: This is a **block-type** attribute. Use block syntax, not assignment.');
      lines.push('');
      lines.push('**Common Mistake**: `' + syntaxGuidance.incorrectSyntax.split('  #')[0].trim() + '`');
      lines.push('');
      lines.push(syntaxGuidance.incorrectSyntax);
    }
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
  lines.push('| `syntax` | Get Terraform syntax guidance (block vs attribute) |');
  lines.push('| `summary` | This summary |');

  return lines.join('\n');
}

// =============================================================================
// SYNTAX GUIDANCE HANDLER
// =============================================================================

function handleSyntax(input: MetadataInput, format: ResponseFormat): string {
  const { resource, attribute } = input;

  if (!resource) {
    return formatError(
      'Missing parameter',
      'resource parameter is required for syntax operation',
      format,
    );
  }

  // If attribute specified, return guidance for that specific attribute
  if (attribute) {
    const guidance = getAttributeSyntaxGuidance(resource, attribute);

    if (!guidance) {
      return formatError(
        'Attribute not found',
        `Attribute '${attribute}' not found in resource '${resource}'`,
        format,
      );
    }

    if (format === ResponseFormat.JSON) {
      return JSON.stringify(
        {
          operation: 'syntax',
          resource,
          attribute,
          ...guidance,
        },
        null,
        2,
      );
    }

    // Markdown format for specific attribute
    const lines: string[] = [
      `# Syntax Guidance: ${resource}.${attribute}`,
      '',
      '## Correct Usage',
      '',
      `**Type**: ${guidance.terraformType}`,
      '',
      `**Correct Syntax**: \`${guidance.correctSyntax}\``,
      '',
      '**Example**:',
      '```hcl',
      guidance.example,
      '```',
      '',
    ];

    if (guidance.isOneOf) {
      lines.push('## OneOf Group');
      lines.push('');
      lines.push(`This attribute is part of OneOf group: \`${guidance.oneOfGroup}\``);
      lines.push('');
      lines.push('**Mutually Exclusive**: Only one attribute from this group can be set.');
      lines.push('');
    }

    lines.push('## Common Mistakes');
    lines.push('');
    lines.push(`**Incorrect**: \`${guidance.incorrectSyntax.split('  #')[0].trim()}\``);
    lines.push('');
    lines.push(guidance.incorrectSyntax);
    lines.push('');

    return lines.join('\n');
  }

  // Return full syntax guide for the resource
  const guide = generateTerraformSyntaxGuide(resource);

  if (!guide) {
    return formatError(
      'Resource not found',
      `Resource '${resource}' not found in metadata`,
      format,
    );
  }

  if (format === ResponseFormat.JSON) {
    return JSON.stringify(
      {
        operation: 'syntax',
        resource,
        totalBlocks: guide.totalBlocks,
        totalAttributes: guide.totalAttributes,
        totalOneOfGroups: guide.totalOneOfGroups,
        blocks: guide.blocks,
        attributes: guide.attributes,
        oneOfGroups: guide.oneOfGroups,
      },
      null,
      2,
    );
  }

  // Return the generated markdown guide
  return guide.syntaxGuide;
}

// =============================================================================
// VALIDATE OPERATION HANDLER
// =============================================================================

interface ValidationError {
  type: 'boolean_assignment' | 'unsupported_argument' | 'unsupported_block' | 'reference_inline';
  field: string;
  line?: number;
  location?: string;
  found: string;
  expected: string;
  reason: string;
  suggestions?: string[];
  referenceResource?: string;
}

/**
 * Validates a Terraform config snippet for common syntax errors
 * Phase 2: Now detects unsupported arguments/blocks, not just boolean assignment
 */
function handleValidate(input: MetadataInput, format: ResponseFormat): string {
  const { resource, config } = input;

  if (!resource) {
    return formatError(
      'Missing parameter',
      'resource parameter is required for validate operation',
      format,
    );
  }

  if (!config) {
    return formatError(
      'Missing parameter',
      'config parameter is required for validate operation. Provide a Terraform config snippet to validate.',
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

  const errors: ValidationError[] = [];
  const validFields = Object.keys(resourceMeta.attributes);

  // Phase 1: Check for boolean assignment errors (existing)
  const booleanAssignmentPattern = /^\s*(\w+)\s*=\s*(true|false)\s*$/gm;
  let match;

  while ((match = booleanAssignmentPattern.exec(config)) !== null) {
    const fieldName = match[1];
    const value = match[2];

    const attrMeta = resourceMeta.attributes[fieldName];
    if (attrMeta && attrMeta.is_block && attrMeta.type === 'object') {
      errors.push({
        type: 'boolean_assignment',
        field: fieldName,
        found: `${fieldName} = ${value}`,
        expected: `${fieldName} {}`,
        reason: attrMeta.oneof_group
          ? `This is a OneOf choice field in group '${attrMeta.oneof_group}', not a boolean`
          : 'This is a block-type attribute, not a boolean',
      });
    }
  }

  // Phase 2: Parse HCL and check for unsupported fields
  const parsed = parseHCLForFields(config);

  // Check each parsed field against the schema
  for (const field of parsed.fields) {
    // Skip standard fields that are always valid at root level
    if (!field.parent && ['name', 'namespace', 'description', 'labels', 'annotations'].includes(field.name)) {
      continue;
    }

    // Check if field exists in resource schema
    const parentMeta = field.parent ? resourceMeta.attributes[field.parent] : null;

    // Strategy 1: Check if parent is a KNOWN reference block (handles nested cases like pool inside origin_pools_weights)
    // This catches cases where the parent block name is in KNOWN_REFERENCE_BLOCKS even if
    // we don't have schema metadata for it (e.g., nested blocks)
    if (field.parent && KNOWN_REFERENCE_BLOCKS.has(field.parent)) {
      const allowedFields = ['name', 'namespace', 'tenant'];
      if (!allowedFields.includes(field.name)) {
        const refResource = getReferenceResourceName(field.parent);
        errors.push({
          type: field.type === 'block' ? 'unsupported_block' : 'reference_inline',
          field: field.name,
          line: field.line,
          location: `Inside \`${field.parent}\` block`,
          found: field.type === 'block' ? `${field.name} {}` : `${field.name} = ${field.value || '...'}`,
          expected: `Only name, namespace, tenant are valid inside ${field.parent}`,
          reason: `\`${field.parent}\` is a reference block - it references a separate resource, not inline configuration`,
          suggestions: [`Create a separate ${refResource} resource with these parameters`],
          referenceResource: refResource,
        });
      }
      continue; // Skip further checks for this field
    }

    // Strategy 2: Check parent using schema metadata (handles top-level blocks)
    if (field.parent && parentMeta) {
      // Field is inside a parent block - check if parent is a reference block via schema
      if (isReferenceBlock(parentMeta, field.parent)) {
        const allowedFields = ['name', 'namespace', 'tenant'];
        if (!allowedFields.includes(field.name)) {
          const refResource = getReferenceResourceName(field.parent);
          errors.push({
            type: field.type === 'block' ? 'unsupported_block' : 'reference_inline',
            field: field.name,
            line: field.line,
            location: `Inside \`${field.parent}\` block`,
            found: field.type === 'block' ? `${field.name} {}` : `${field.name} = ${field.value || '...'}`,
            expected: `Only name, namespace, tenant are valid inside ${field.parent}`,
            reason: `\`${field.parent}\` is a reference block - it references a separate resource, not inline configuration`,
            suggestions: [`Create a separate ${refResource} resource with these parameters`],
            referenceResource: refResource,
          });
        }
      }
      // Note: For non-reference blocks without nested_attributes in schema,
      // we can't validate nested fields - skip validation for these cases
    } else if (!field.parent) {
      // Field is at root level - check against resource schema
      if (!resourceMeta.attributes[field.name]) {
        const similar = findSimilarFieldNames(field.name, validFields);
        errors.push({
          type: field.type === 'block' ? 'unsupported_block' : 'unsupported_argument',
          field: field.name,
          line: field.line,
          found: field.type === 'block' ? `${field.name} {}` : `${field.name} = ${field.value || '...'}`,
          expected: `Valid attributes for ${resource}`,
          reason: `An ${field.type === 'block' ? 'block' : 'argument'} named "${field.name}" is not expected here`,
          suggestions: similar.length > 0 ? [`Did you mean: ${similar.join(', ')}?`] : undefined,
        });
      }
    }
  }

  // Build response
  if (errors.length === 0) {
    if (format === ResponseFormat.JSON) {
      return JSON.stringify({
        operation: 'validate',
        resource,
        status: 'valid',
        errors: [],
        message: 'No syntax errors detected in the provided configuration.',
      }, null, 2);
    }

    return [
      `# Validation Results: ${resource}`,
      '',
      '✅ **No syntax errors detected**',
      '',
      'The provided configuration appears to use correct syntax.',
      '',
      '*Note: This validates common syntax patterns only. Run `terraform validate` for complete validation.*',
    ].join('\n');
  }

  // Errors found
  if (format === ResponseFormat.JSON) {
    return JSON.stringify({
      operation: 'validate',
      resource,
      status: 'invalid',
      errors,
      error_count: errors.length,
    }, null, 2);
  }

  const lines: string[] = [
    `# Validation Results: ${errors.length} Error(s) Found`,
    '',
  ];

  for (let i = 0; i < errors.length; i++) {
    const err = errors[i];
    const errorTitle = getErrorTitle(err.type);
    lines.push(`## Error ${i + 1}: ${errorTitle}`);
    lines.push('');
    lines.push(`- **Field**: \`${err.field}\``);
    if (err.line) {
      lines.push(`- **Line**: ${err.line}`);
    }
    if (err.location) {
      lines.push(`- **Location**: ${err.location}`);
    }
    lines.push(`- **Found**: \`${err.found}\``);
    lines.push(`- **Expected**: ${err.expected}`);
    lines.push(`- **Reason**: ${err.reason}`);
    if (err.suggestions && err.suggestions.length > 0) {
      lines.push(`- **Suggestion**: ${err.suggestions.join('; ')}`);
    }
    lines.push('');
  }

  // Add corrected pattern for reference block errors
  const refErrors = errors.filter(e => e.type === 'reference_inline' || (e.type === 'unsupported_block' && e.referenceResource));
  if (refErrors.length > 0) {
    const refErr = refErrors[0];
    const parentBlock = refErr.location?.match(/`(\w+)`/)?.[1] || 'unknown';
    const refResource = refErr.referenceResource || `f5xc_${parentBlock}`;

    lines.push('## Correct Pattern');
    lines.push('');
    lines.push('```hcl');
    lines.push(`# Create the ${parentBlock} resource separately`);
    lines.push(`resource "${refResource}" "example" {`);
    lines.push(`  name      = "my-${parentBlock}"`);
    lines.push('  namespace = "default"');
    lines.push('');
    lines.push('  # Add your configuration here');
    lines.push('  # ...');
    lines.push('}');
    lines.push('');
    lines.push(`# Reference it in ${resource}`);
    lines.push(`resource "f5xc_${resource}" "example" {`);
    lines.push(`  ${parentBlock} {`);
    lines.push(`    name      = ${refResource}.example.name`);
    lines.push('    namespace = "default"');
    lines.push('  }');
    lines.push('}');
    lines.push('```');
  } else {
    // For non-reference errors, show corrected config
    const boolErrors = errors.filter(e => e.type === 'boolean_assignment');
    if (boolErrors.length > 0) {
      lines.push('## Corrected Configuration');
      lines.push('');
      lines.push('```hcl');
      lines.push(generateCorrectedConfig(config, boolErrors));
      lines.push('```');
    }
  }

  return lines.join('\n');
}

/**
 * Get human-readable error title
 */
function getErrorTitle(type: ValidationError['type']): string {
  switch (type) {
    case 'boolean_assignment':
      return 'Invalid boolean assignment';
    case 'unsupported_argument':
      return 'Unsupported argument';
    case 'unsupported_block':
      return 'Unsupported block type';
    case 'reference_inline':
      return 'Inline configuration in reference block';
    default:
      return 'Validation error';
  }
}

/**
 * Generates corrected config by replacing boolean assignments with empty blocks
 */
function generateCorrectedConfig(
  config: string,
  errors: ValidationError[],
): string {
  let corrected = config;
  for (const err of errors) {
    if (err.type === 'boolean_assignment') {
      // Replace "field = true" or "field = false" with "field {}"
      const pattern = new RegExp(`\\b${err.field}\\s*=\\s*(true|false)`, 'g');
      corrected = corrected.replace(pattern, `${err.field} {}`);
    }
  }
  return corrected;
}

// =============================================================================
// EXAMPLE OPERATION HANDLER
// =============================================================================

/**
 * Generates complete, syntactically correct Terraform examples
 */
function handleExample(input: MetadataInput, format: ResponseFormat): string {
  const { resource, pattern = 'basic' } = input;

  if (!resource) {
    return formatError(
      'Missing parameter',
      'resource parameter is required for example operation',
      format,
    );
  }

  // Get example based on resource type and pattern
  const example = getExampleForResource(resource, pattern);

  if (!example) {
    return formatError(
      'Example not found',
      `No example available for resource '${resource}' with pattern '${pattern}'. Available patterns: basic, with_waf, with_bot_defense, with_rate_limiting, full`,
      format,
    );
  }

  if (format === ResponseFormat.JSON) {
    return JSON.stringify({
      operation: 'example',
      resource,
      pattern,
      ...example,
    }, null, 2);
  }

  return example.markdown;
}

/**
 * Returns example configuration for a resource
 */
function getExampleForResource(
  resource: string,
  pattern: string,
): { markdown: string; terraform: string; description: string } | null {
  // HTTP Load Balancer examples
  if (resource === 'http_loadbalancer') {
    return getHttpLoadbalancerExample(pattern);
  }

  // Origin Pool examples
  if (resource === 'origin_pool') {
    return getOriginPoolExample(pattern);
  }

  // Namespace example
  if (resource === 'namespace') {
    return {
      description: 'Basic namespace configuration',
      terraform: `resource "f5xc_namespace" "example" {
  name = "my-namespace"
}`,
      markdown: [
        '# Complete Example: Namespace (Basic Pattern)',
        '',
        'Creates a namespace in F5 Distributed Cloud.',
        '',
        '```hcl',
        'resource "f5xc_namespace" "example" {',
        '  name = "my-namespace"',
        '}',
        '```',
        '',
        '## Key Syntax Notes',
        '',
        '- Namespace is the simplest resource - just needs a name',
        '- Most other resources require a namespace reference',
      ].join('\n'),
    };
  }

  // App Firewall example
  if (resource === 'app_firewall') {
    return getAppFirewallExample(pattern);
  }

  // TCP Load Balancer example
  if (resource === 'tcp_loadbalancer') {
    return getTcpLoadbalancerExample(pattern);
  }

  // UDP Load Balancer example
  if (resource === 'udp_loadbalancer') {
    return getUdpLoadbalancerExample(pattern);
  }

  // Certificate example
  if (resource === 'certificate') {
    return getCertificateExample(pattern);
  }

  // Generic fallback - generate from syntax guide
  const syntaxGuide = generateTerraformSyntaxGuide(resource);
  if (syntaxGuide) {
    const blocks = syntaxGuide.blocks.slice(0, 5);
    const attrs = syntaxGuide.attributes.slice(0, 5);

    const exampleLines = [
      `resource "f5xc_${resource}" "example" {`,
      '  name      = "example-' + resource.replace(/_/g, '-') + '"',
      '  namespace = "default"',
      '',
    ];

    // Add required OneOf selections
    for (const [groupName, fields] of Object.entries(syntaxGuide.oneOfGroups)) {
      if (fields.length > 0) {
        exampleLines.push(`  # OneOf group: ${groupName} - choose one`);
        exampleLines.push(`  ${fields[0]} {}`);
        exampleLines.push('');
      }
    }

    exampleLines.push('}');

    return {
      description: `Auto-generated example for ${resource}`,
      terraform: exampleLines.join('\n'),
      markdown: [
        `# Complete Example: ${resource} (Auto-Generated)`,
        '',
        '⚠️ **Note**: This is an auto-generated example. Verify required fields in documentation.',
        '',
        '```hcl',
        exampleLines.join('\n'),
        '```',
        '',
        '## Block-Type Attributes (use empty block syntax)',
        '',
        blocks.map(b => `- \`${b} {}\``).join('\n'),
        '',
        '## Simple Attributes',
        '',
        attrs.map(a => `- \`${a} = <value>\``).join('\n'),
      ].join('\n'),
    };
  }

  return null;
}

function getTcpLoadbalancerExample(pattern: string): { markdown: string; terraform: string; description: string } {
  const baseExample = `resource "f5xc_tcp_loadbalancer" "example" {
  name      = "my-tcp-lb"
  namespace = "default"

  # DNS configuration
  dns_volterra_managed = true

  # Origin pool reference
  origin_pools_weights {
    pool {
      name      = "my-origin-pool"
      namespace = "default"
    }
    weight   = 1
    priority = 1
  }

  # Listen port
  listen_port = 8080

  # Advertise configuration - use empty block syntax
  advertise_on_public_default_vip {}
}`;

  if (pattern === 'basic') {
    return {
      description: 'Basic TCP load balancer configuration with origin pool',
      terraform: baseExample,
      markdown: [
        '# Complete Example: TCP Load Balancer (Basic Pattern)',
        '',
        'Creates a TCP load balancer in F5 Distributed Cloud.',
        '',
        '## Required Resources',
        '',
        '```hcl',
        baseExample,
        '```',
        '',
        '## Key Syntax Notes',
        '',
        '1. **origin_pools_weights**: References existing origin pool',
        '2. **advertise_on_public_default_vip**: Use empty block syntax {}',
        '3. **dns_volterra_managed**: Enable F5XC-managed DNS',
        '4. **listen_port**: TCP port for load balancer to listen on',
        '',
        '## Important Notes',
        '',
        '- Origin pool must exist before creating TCP load balancer',
        '- Use weight and priority for traffic distribution',
        '- Empty block syntax for advertise configuration choices',
      ].join('\n'),
    };
  }

  // Fallback
  return {
    description: 'TCP load balancer example',
    terraform: baseExample,
    markdown: `# TCP Load Balancer Example\n\n\`\`\`hcl\n${baseExample}\n\`\`\``,
  };
}

function getUdpLoadbalancerExample(pattern: string): { markdown: string; terraform: string; description: string } {
  const baseExample = `resource "f5xc_udp_loadbalancer" "example" {
  name      = "my-udp-lb"
  namespace = "default"

  # DNS configuration
  dns_volterra_managed = true

  # Origin pool reference
  origin_pools_weights {
    pool {
      name      = "my-origin-pool"
      namespace = "default"
    }
    weight   = 1
    priority = 1
  }

  # Port range (single port or range like "53" or "5000-5010")
  port_ranges = "53"

  # Advertise configuration - use empty block syntax
  advertise_on_public_default_vip {}
}`;

  if (pattern === 'basic') {
    return {
      description: 'Basic UDP load balancer configuration with origin pool',
      terraform: baseExample,
      markdown: [
        '# Complete Example: UDP Load Balancer (Basic Pattern)',
        '',
        'Creates a UDP load balancer in F5 Distributed Cloud.',
        '',
        '## Required Resources',
        '',
        '```hcl',
        baseExample,
        '```',
        '',
        '## Key Syntax Notes',
        '',
        '1. **origin_pools_weights**: References existing origin pool',
        '2. **advertise_on_public_default_vip**: Use empty block syntax {}',
        '3. **dns_volterra_managed**: Enable F5XC-managed DNS',
        '4. **port_ranges**: Single port ("53") or range ("5000-5010")',
        '',
        '## Important Notes',
        '',
        '- Origin pool must exist before creating UDP load balancer',
        '- Use weight and priority for traffic distribution',
        '- Empty block syntax for advertise configuration choices',
        '- Port ranges support single port or range format',
      ].join('\n'),
    };
  }

  // Fallback
  return {
    description: 'UDP load balancer example',
    terraform: baseExample,
    markdown: `# UDP Load Balancer Example\n\n\`\`\`hcl\n${baseExample}\n\`\`\``,
  };
}

function getCertificateExample(pattern: string): { markdown: string; terraform: string; description: string } {
  const baseExample = `resource "f5xc_certificate" "example" {
  name      = "my-certificate"
  namespace = "default"

  # Certificate data in PEM format (replace with actual certificate)
  certificate_url = "string:///-----BEGIN CERTIFICATE-----\\nMIIDdzCCAl+gAwIBAgIEAgAAuTANBgkqhkiG9w0BAQUFADBaMQswCQYDVQQGEwJJRTESMBAGA1UEChMJQmFsdGltb3JlMRMwEQYDVQQLEwpDeWJlclRydXN0MSIwIAYDVQQDExlCYWx0aW1vcmUgQ3liZXJUcnVzdCBSb290MB4XDTAwMDUxMjE4NDYwMFoXDTI1MDUxMjIzNTkwMFowWjELMAkGA1UEBhMCSUUxEjAQBgNVBAoTCUJhbHRpbW9yZTETMBEGA1UECxMKQ3liZXJUcnVzdDEiMCAGA1UEAxMZQmFsdGltb3JlIEN5YmVyVHJ1c3QgUm9vdDCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKMEuyKrmD1X6CZymrV51Cni4eiVgLGw41uOKymaZN+hXe2wCQVt2yguzmKiYv60iNoS6zjrIZ3AQSsBUnuId9Mcj8e6uYi1agnnc+gRQKfRzMpijS3ljwumUNKoUMMo6vWrJYeKmpYcqWe4PwzV9\/lSEy\/CG9VwcPCPwBLKBsua4dnKM3p31vjsufFoREJIE9LAwqSuXmD+tqYF\/LTdB1kC1FkYmGP1pWPgkAx9XbIGevOF6uvUA65ehD5f\/xXtabz5OTZydc93Uk3zyZAsuT3lySNTPx8kmCFcB5kpvcY67Oduhjprl3RjM71oGDHweI12v\/yejl0qhqdNkNwnGjkCAwEAAaNFMEMwHQYDVR0OBBYEFOWdWTCCR1jMrPoIVDaGezq1BE3wMBIGA1UdEwEB\/wQIMAYBAf8CAQMwDgYDVR0PAQH\/BAQDAgEGMA0GCSqGSIb3DQEBBQUAA4IBAQCFDF2O5G9RaEIFoN27TyclhAO992T9Ldcw46QQF+vaKSm2eT929hkTI7gQCvlYpNRhcL0EYWoSihfVCr3FvDB81ukMJY2GQE\/szKN+OMY3EU\/t3WgxjkzSswF07r51XgdIGn9w\/xZchMB5hbgF\/X++ZRGjD8ACtPhSNzkE1akxehi\/oCr0Epn3o0WC4zxe9Z2etciefC7IpJ5OCBRLbf1wbWsaY71k5h+3zvDyny67G7fyUIhzksLi4xaNmjICq44Y3ekQEe5+NauQrz4wlHrQMz2nZQ\/1\/I6eYs9HRCwBXbsdtTLSR9I4LtD+gdwyah617jzV\/OeBHRnDJELqYzmp\\n-----END CERTIFICATE-----"
}`;

  if (pattern === 'basic') {
    return {
      description: 'Basic certificate configuration with TLS certificate and private key',
      terraform: baseExample,
      markdown: [
        '# Complete Example: Certificate (Basic Pattern)',
        '',
        'Creates a TLS certificate resource in F5 Distributed Cloud.',
        '',
        '## Required Resources',
        '',
        '```hcl',
        baseExample,
        '```',
        '',
        '## Key Syntax Notes',
        '',
        '1. **certificate_url**: PEM-encoded certificate with string:/// prefix',
        '2. **private_key**: Nested block with blindfold_secret_info_internal',
        '3. **location**: PEM-encoded private key with string:/// prefix',
        '',
        '## Important Notes',
        '',
        '- Replace MIICert... with your actual certificate content',
        '- Replace MIIEprivate... with your actual private key content',
        '- PEM format must include BEGIN/END markers',
        '- Use \\n for line breaks in the string',
      ].join('\n'),
    };
  }

  // Fallback
  return {
    description: 'Certificate example',
    terraform: baseExample,
    markdown: `# Certificate Example\n\n\`\`\`hcl\n${baseExample}\n\`\`\``,
  };
}

function getHttpLoadbalancerExample(pattern: string): { markdown: string; terraform: string; description: string } {
  const baseExample = `# 1. Create the origin pool
resource "f5xc_origin_pool" "example" {
  name      = "httpbin-pool"
  namespace = "default"

  # TLS choice - use empty block syntax
  no_tls {}

  # Port choice - use empty block syntax
  automatic_port {}

  origin_servers {
    public_name {
      dns_name = "httpbin.org"
    }
  }

  loadbalancer_algorithm = "ROUND_ROBIN"
}

# 2. Create the HTTP load balancer
resource "f5xc_http_loadbalancer" "example" {
  name      = "httpbin-lb"
  namespace = "default"
  domains   = ["httpbin.example.com"]

  # Load balancer type - use empty block syntax
  https_auto_cert {
    http_redirect = true
    add_hsts      = false
  }

  # Advertising - use empty block syntax
  advertise_on_public_default_vip {}

  # Origin pool reference
  default_route_pools {
    pool {
      name      = f5xc_origin_pool.example.name
      namespace = "default"
    }
  }

  # Security options - use empty block syntax
  disable_waf {}
  no_challenge {}
  disable_rate_limit {}
  no_service_policies {}
  disable_bot_defense {}
  disable_api_definition {}
  disable_api_discovery {}

  # Load balancing algorithm - use empty block syntax
  round_robin {}
}`;

  if (pattern === 'basic') {
    return {
      description: 'Basic HTTP Load Balancer with origin pool pointing to httpbin.org',
      terraform: baseExample,
      markdown: [
        '# Complete Example: HTTP Load Balancer (Basic Pattern)',
        '',
        'Creates an HTTP load balancer with an origin pool pointing to httpbin.org.',
        '',
        '## Required Resources',
        '',
        '```hcl',
        baseExample,
        '```',
        '',
        '## Key Syntax Notes',
        '',
        '1. **Empty blocks for OneOf choices**: `no_tls {}`, `advertise_on_public_default_vip {}`',
        '2. **Nested blocks with values**: `https_auto_cert { http_redirect = true }`',
        '3. **Reference syntax**: `f5xc_origin_pool.example.name`',
        '',
        '## Common Mistakes to Avoid',
        '',
        '```hcl',
        '# WRONG - Do not use boolean assignment for these fields!',
        'no_tls = true                          # ERROR: use no_tls {}',
        'advertise_on_public_default_vip = true # ERROR: use advertise_on_public_default_vip {}',
        'round_robin = true                     # ERROR: use round_robin {}',
        '```',
      ].join('\n'),
    };
  }

  if (pattern === 'with_waf') {
    const wafExample = baseExample.replace(
      '  disable_waf {}',
      `  app_firewall {
    name      = f5xc_app_firewall.example.name
    namespace = "default"
  }`,
    );
    const fullExample = `# 0. Create App Firewall first
resource "f5xc_app_firewall" "example" {
  name      = "my-waf"
  namespace = "default"

  # Use blocking mode (not monitoring)
  blocking {}
}

${wafExample}`;

    return {
      description: 'HTTP Load Balancer with WAF/App Firewall protection',
      terraform: fullExample,
      markdown: [
        '# Complete Example: HTTP Load Balancer with WAF',
        '',
        'Creates an HTTP load balancer with Web Application Firewall protection.',
        '',
        '```hcl',
        fullExample,
        '```',
        '',
        '## WAF Configuration Notes',
        '',
        '- Create `f5xc_app_firewall` resource first',
        '- Reference it in `http_loadbalancer.app_firewall` block',
        '- Use `blocking {}` for enforcement, `monitoring {}` for detection-only',
      ].join('\n'),
    };
  }

  // Default to basic
  return getHttpLoadbalancerExample('basic');
}

function getOriginPoolExample(_pattern: string): { markdown: string; terraform: string; description: string } {
  const example = `# Optional: Create a healthcheck resource first
# (healthcheck is a REFERENCE BLOCK - only name/namespace/tenant are valid)
resource "f5xc_healthcheck" "example" {
  name      = "my-healthcheck"
  namespace = "default"

  # Health check type - select ONE
  http_health_check {
    use_origin_server_name {}
    path = "/health"
  }
  # OR: tcp_health_check {}
  # OR: grpc_health_check { ... }

  healthy_threshold   = 2
  unhealthy_threshold = 3
  interval            = 15
  timeout             = 5
}

resource "f5xc_origin_pool" "example" {
  name      = "my-origin-pool"
  namespace = "default"

  # TLS Choice - select ONE (use empty block syntax)
  # Option 1: No TLS (plain HTTP to origin)
  no_tls {}

  # Option 2: Use TLS (uncomment and configure)
  # use_tls {
  #   # Nested TLS configuration
  #   skip_server_verification {}
  #   # OR
  #   # use_server_verification {}
  # }

  # Port Choice - select ONE
  # Option 1: Automatic port (80 for HTTP, 443 for HTTPS)
  automatic_port {}

  # Option 2: Same as endpoint port
  # same_as_endpoint_port {}

  # Option 3: Explicit port number (use attribute, not block)
  # port = 8080

  # Origin servers (at least one required)
  origin_servers {
    # Public DNS name
    public_name {
      dns_name = "api.example.com"
    }
  }

  # Healthcheck - REFERENCE BLOCK (only name/namespace/tenant!)
  # This references the separate f5xc_healthcheck resource above
  # DO NOT put configuration parameters here (interval, timeout, etc.)
  healthcheck {
    name      = f5xc_healthcheck.example.name
    namespace = "default"
  }

  # Load balancing algorithm (attribute, not block)
  loadbalancer_algorithm = "ROUND_ROBIN"
}`;

  return {
    description: 'Origin Pool with healthcheck reference and common configuration patterns',
    terraform: example,
    markdown: [
      '# Complete Example: Origin Pool with Healthcheck',
      '',
      'Creates an origin pool with healthcheck reference and multiple configuration options.',
      '',
      '```hcl',
      example,
      '```',
      '',
      '## Reference Blocks vs Inline Blocks',
      '',
      '**IMPORTANT**: The `healthcheck` field is a **reference block**, NOT an inline configuration block.',
      '',
      '### Reference Blocks (only accept name/namespace/tenant)',
      '- `healthcheck {}` - References a separate `f5xc_healthcheck` resource',
      '- Configuration parameters (interval, timeout, path, etc.) go in the separate resource',
      '',
      '### Common Error (WRONG):',
      '```hcl',
      '# DO NOT DO THIS - healthcheck is a reference, not inline config',
      'healthcheck {',
      '  interval_seconds = 30  # ERROR!',
      '  http_request {}        # ERROR!',
      '}',
      '```',
      '',
      '## OneOf Groups Explained',
      '',
      '### tls_choice',
      '- `no_tls {}` - Plain HTTP to origin (use empty block)',
      '- `use_tls { ... }` - TLS to origin (use block with nested config)',
      '',
      '### port_choice',
      '- `automatic_port {}` - Auto-select based on TLS (use empty block)',
      '- `same_as_endpoint_port {}` - Use endpoint\'s port (use empty block)',
      '- `port = 8080` - Explicit port (use attribute assignment)',
      '',
      '## Key Syntax Pattern',
      '',
      'Block-type fields like `no_tls`, `automatic_port` use `{}` syntax.',
      'Attribute fields like `port` use `= value` syntax.',
      'Reference fields like `healthcheck` only accept `name`/`namespace`/`tenant`.',
    ].join('\n'),
  };
}

function getAppFirewallExample(_pattern: string): { markdown: string; terraform: string; description: string } {
  const example = `resource "f5xc_app_firewall" "example" {
  name      = "my-waf-policy"
  namespace = "default"

  # Enforcement mode - select ONE (use empty block syntax)
  # Option 1: Blocking mode (actively blocks attacks)
  blocking {}

  # Option 2: Monitoring mode (logs only, no blocking)
  # monitoring {}

  # Anonymization setting - select ONE
  default_anonymization {}
  # OR: custom_anonymization { ... }
  # OR: disable_anonymization {}

  # Bot protection settings - select ONE
  default_bot_setting {}
  # OR: custom_bot_protection_setting { ... }

  # Detection settings - select ONE
  default_detection_settings {}
  # OR: detection_settings { ... }
}`;

  return {
    description: 'App Firewall (WAF) with common configuration patterns',
    terraform: example,
    markdown: [
      '# Complete Example: App Firewall (WAF)',
      '',
      'Creates a Web Application Firewall policy.',
      '',
      '```hcl',
      example,
      '```',
      '',
      '## Key OneOf Groups',
      '',
      '### enforcement_mode_choice',
      '- `blocking {}` - Actively block attacks',
      '- `monitoring {}` - Detection only, no blocking',
      '',
      '### anonymization_setting',
      '- `default_anonymization {}` - Use default settings',
      '- `custom_anonymization { ... }` - Custom configuration',
      '- `disable_anonymization {}` - No anonymization',
      '',
      '## Important',
      '',
      'All OneOf options use **empty block syntax** `{}`, not boolean assignment.',
    ].join('\n'),
  };
}

// =============================================================================
// MISTAKES OPERATION HANDLER
// =============================================================================

/**
 * Returns common syntax mistakes and how to fix them
 */
function handleMistakes(input: MetadataInput, format: ResponseFormat): string {
  const { resource } = input;

  const mistakes = getCommonMistakes(resource);

  if (format === ResponseFormat.JSON) {
    return JSON.stringify({
      operation: 'mistakes',
      resource: resource || 'all',
      mistakes,
      total: mistakes.length,
    }, null, 2);
  }

  const lines: string[] = [
    '# Common Terraform Syntax Mistakes for F5XC Provider',
    '',
  ];

  if (resource) {
    lines.push(`**Resource**: ${resource}`);
    lines.push('');
  }

  for (let i = 0; i < mistakes.length; i++) {
    const m = mistakes[i];
    lines.push(`## Mistake #${i + 1}: ${m.title}`);
    lines.push('');
    lines.push(`**Severity**: ${m.severity}`);
    lines.push('');
    lines.push('**Wrong**:');
    lines.push('```hcl');
    lines.push(m.wrong);
    lines.push('```');
    lines.push('');
    lines.push('**Correct**:');
    lines.push('```hcl');
    lines.push(m.correct);
    lines.push('```');
    lines.push('');
    lines.push(`**Explanation**: ${m.explanation}`);
    lines.push('');
    if (m.detection) {
      lines.push(`**How to detect**: ${m.detection}`);
      lines.push('');
    }
  }

  return lines.join('\n');
}

interface CommonMistake {
  title: string;
  severity: 'Critical' | 'High' | 'Medium';
  wrong: string;
  correct: string;
  explanation: string;
  detection?: string;
  resources?: string[];
}

function getCommonMistakes(resource?: string): CommonMistake[] {
  const allMistakes: CommonMistake[] = [
    {
      title: 'Boolean Assignment Instead of Empty Block',
      severity: 'Critical',
      wrong: `no_tls = true
advertise_on_public_default_vip = true
round_robin = true
disable_api_definition = true
no_challenge = true`,
      correct: `no_tls {}
advertise_on_public_default_vip {}
round_robin {}
disable_api_definition {}
no_challenge {}`,
      explanation: 'OneOf choice fields in F5XC provider are modeled as blocks, not booleans. Use empty block syntax {} to select an option.',
      detection: 'Query f5xc_terraform_metadata(operation="syntax", resource="...") and check "Block-Type Attributes" section.',
      resources: ['http_loadbalancer', 'origin_pool', 'tcp_loadbalancer', 'app_firewall'],
    },
    {
      title: 'Missing Required OneOf Selection',
      severity: 'Critical',
      wrong: `resource "f5xc_http_loadbalancer" "example" {
  name      = "test"
  namespace = "default"
  domains   = ["example.com"]
  # Missing loadbalancer_type selection!
}`,
      correct: `resource "f5xc_http_loadbalancer" "example" {
  name      = "test"
  namespace = "default"
  domains   = ["example.com"]

  # Required: one of http/https/https_auto_cert
  https_auto_cert {
    http_redirect = true
  }
}`,
      explanation: 'Some OneOf groups have a required selection. For http_loadbalancer, you must specify http {}, https {}, or https_auto_cert {}.',
      detection: 'Query f5xc_terraform_metadata(operation="oneof", resource="...") to see required OneOf groups.',
      resources: ['http_loadbalancer'],
    },
    {
      title: 'Confusing Block Attribute with Simple Attribute',
      severity: 'High',
      wrong: `# Using 'port' thinking it's part of port_choice OneOf
port {}  # ERROR: port is a simple attribute

# Using loadbalancer_algorithm as a block
round_robin {}  # This is for hash_policy_choice, not algorithm`,
      correct: `# port is a simple integer attribute
port = 8080

# loadbalancer_algorithm is a string attribute
loadbalancer_algorithm = "ROUND_ROBIN"

# hash_policy_choice uses blocks for stickiness
round_robin {}  # OR source_ip_stickiness {} etc.`,
      explanation: 'Not all similarly-named fields work the same way. Some are simple attributes (use =), others are blocks (use {}).',
      detection: 'Query f5xc_terraform_metadata(operation="attribute", resource="...", attribute="...") to check the type.',
      resources: ['origin_pool', 'http_loadbalancer'],
    },
    {
      title: 'Nested Block Syntax Errors',
      severity: 'Medium',
      wrong: `use_tls {
  no_mtls = true              # ERROR
  volterra_trusted_ca = true  # ERROR
}`,
      correct: `use_tls {
  no_mtls {}              # Empty block syntax
  volterra_trusted_ca {}  # Empty block syntax

  tls_config {
    default_security {}   # Also empty block
  }
}`,
      explanation: 'Empty block syntax applies to nested fields too, not just top-level attributes.',
      detection: 'Recursively check nested fields using the syntax operation.',
      resources: ['origin_pool', 'http_loadbalancer'],
    },
    {
      title: 'Inline Configuration in Reference Block',
      severity: 'Critical',
      wrong: `# WRONG - trying to configure healthcheck inline
resource "f5xc_origin_pool" "example" {
  name      = "my-pool"
  namespace = "default"

  healthcheck {
    interval_seconds    = 30   # ERROR: not a valid attribute
    timeout_seconds     = 5    # ERROR: not a valid attribute
    healthy_threshold   = 2    # ERROR: not a valid attribute
    unhealthy_threshold = 3    # ERROR: not a valid attribute

    http_request {             # ERROR: not a valid block
      path = "/health"
    }
  }
}`,
      correct: `# CORRECT - create healthcheck as separate resource, then reference it
# Step 1: Create the healthcheck resource
resource "f5xc_healthcheck" "example" {
  name      = "my-healthcheck"
  namespace = "default"

  http_health_check {
    use_origin_server_name {}
    path = "/health"
  }

  healthy_threshold   = 2
  unhealthy_threshold = 3
  interval            = 30
  timeout             = 5
}

# Step 2: Reference it by name in origin_pool
resource "f5xc_origin_pool" "example" {
  name      = "my-pool"
  namespace = "default"

  healthcheck {
    name      = f5xc_healthcheck.example.name
    namespace = "default"
  }
}`,
      explanation: `Some fields like "healthcheck" are REFERENCE BLOCKS - they reference separate resources, not inline configuration.
Reference blocks ONLY accept "name", "namespace", and optionally "tenant" - NOT configuration parameters.
The actual configuration must be done in a separate resource (e.g., f5xc_healthcheck).`,
      detection: `Query f5xc_terraform_metadata(operation="attribute", resource="origin_pool", attribute="healthcheck")
and check if it's a reference block (only has name/namespace/tenant nested attributes).
Or use f5xc_terraform_metadata(operation="validate", resource="origin_pool", config="...") to detect this error.`,
      resources: ['origin_pool', 'http_loadbalancer', 'tcp_loadbalancer'],
    },
  ];

  // Filter by resource if specified
  if (resource) {
    return allMistakes.filter(m =>
      !m.resources || m.resources.includes(resource),
    );
  }

  return allMistakes;
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
// HCL PARSING AND VALIDATION HELPERS (Phase 2)
// =============================================================================

interface ParsedField {
  name: string;
  type: 'attribute' | 'block';
  line: number;
  parent?: string;  // Parent block name if nested
  value?: string;   // For attributes, the assigned value
}

interface ParsedBlock {
  name: string;
  line: number;
  parent?: string;
  children: ParsedField[];
}

/**
 * Parse HCL config to extract all field/block names for validation
 */
function parseHCLForFields(hcl: string): { fields: ParsedField[]; blocks: ParsedBlock[] } {
  const fields: ParsedField[] = [];
  const blocks: ParsedBlock[] = [];
  const lines = hcl.split('\n');

  // Track block nesting
  const blockStack: { name: string; line: number }[] = [];
  let _braceCount = 0;

  for (let lineNum = 0; lineNum < lines.length; lineNum++) {
    const line = lines[lineNum];
    const trimmed = line.trim();

    // Skip comments and empty lines
    if (trimmed.startsWith('#') || trimmed.startsWith('//') || trimmed === '') {
      continue;
    }

    // Count braces to track nesting
    const openBraces = (line.match(/\{/g) || []).length;
    const closeBraces = (line.match(/\}/g) || []).length;

    // Check for block definition: name { or name "label" {
    const blockMatch = trimmed.match(/^(\w+)\s*(?:"[^"]*"\s*)?\{(.*)$/);
    if (blockMatch) {
      const blockName = blockMatch[1];
      const inlineContent = blockMatch[2] || '';
      // Skip resource/data/provider/variable declarations
      if (!['resource', 'data', 'provider', 'variable', 'output', 'locals', 'terraform', 'module'].includes(blockName)) {
        const parent = blockStack.length > 0 ? blockStack[blockStack.length - 1].name : undefined;
        blocks.push({
          name: blockName,
          line: lineNum + 1,
          parent,
          children: [],
        });
        fields.push({
          name: blockName,
          type: 'block',
          line: lineNum + 1,
          parent,
        });
      }
      blockStack.push({ name: blockMatch[1], line: lineNum + 1 });

      // Handle inline content after opening brace (e.g., "block { attr = value }")
      if (inlineContent.trim() && !inlineContent.trim().startsWith('}')) {
        const currentBlockName = blockStack[blockStack.length - 1].name;
        // Extract attributes from inline content
        const inlineAttrMatch = inlineContent.match(/(\w+)\s*=\s*([^}]+)/);
        if (inlineAttrMatch) {
          const inlineAttrName = inlineAttrMatch[1];
          const inlineAttrValue = inlineAttrMatch[2].trim();
          fields.push({
            name: inlineAttrName,
            type: 'attribute',
            line: lineNum + 1,
            parent: currentBlockName,
            value: inlineAttrValue,
          });
        }
        // Check for nested blocks in inline content
        const inlineBlockMatch = inlineContent.match(/(\w+)\s*\{/);
        if (inlineBlockMatch && !inlineAttrMatch) {
          fields.push({
            name: inlineBlockMatch[1],
            type: 'block',
            line: lineNum + 1,
            parent: currentBlockName,
          });
        }
      }
    }

    // Check for attribute assignment: name = value (only if not part of a block definition)
    const attrMatch = trimmed.match(/^(\w+)\s*=\s*(.+?)$/);
    if (attrMatch && !blockMatch) {
      const attrName = attrMatch[1];
      const attrValue = attrMatch[2];
      const parent = blockStack.length > 0 ? blockStack[blockStack.length - 1].name : undefined;

      // Skip top-level resource attributes like 'name', 'namespace' at resource level
      if (blockStack.length > 0 || !['resource', 'data'].includes(blockStack[0]?.name || '')) {
        fields.push({
          name: attrName,
          type: 'attribute',
          line: lineNum + 1,
          parent: parent !== 'resource' ? parent : undefined,
          value: attrValue,
        });
      }
    }

    // Track brace closing
    _braceCount += openBraces - closeBraces;
    if (closeBraces > 0 && blockStack.length > 0) {
      for (let i = 0; i < closeBraces; i++) {
        if (blockStack.length > 0) {
          blockStack.pop();
        }
      }
    }
  }

  return { fields, blocks };
}

/**
 * Known reference block field names in F5XC provider.
 * These blocks only accept name/namespace/tenant, not inline configuration.
 *
 * COMPREHENSIVE LIST - derived from analysis of all 98 resources in the provider.
 * This list is used for deterministic detection when description patterns are unclear.
 */
const KNOWN_REFERENCE_BLOCKS = new Set([
  // Security references
  'app_firewall',
  'bot_defense_policy',
  'csrf_policy',
  'data_guard_rules',
  'dos_protection',
  'graphql_rules',
  'ip_reputation',
  'jwt_validation',
  'malicious_user_mitigation',
  'protected_cookies',
  'slow_ddos_mitigation',
  'trusted_clients',
  'usb_policy',
  'waf_exclusion_rules',

  // Load balancer and pool references
  'healthcheck',
  'origin_pool',
  'pool',
  'cluster',

  // Policy references
  'rate_limiter',
  'rate_limiter_allowed_prefixes',
  'service_policy',
  'policer',

  // Identity references
  'user_identification',
  'authentication',
  'ad_client',

  // API references
  'api_definition',
  'api_definitions',

  // Certificate references
  'certificate',
  'certificates',
  'crl',
  'trusted_ca',
  'trusted_ca_list',

  // Infrastructure references
  'site',
  'virtual_site',
  'virtual_network',
  'segment',
  'network_connector',
  'cloud_credentials',
  'aws_cred',
  'azure_cred',
  'gcp_cred',
  'k8s_cluster',
  'log_receiver',

  // Site group references
  'dc_cluster_group',
  'dc_cluster_group_inside',
  'dc_cluster_group_sli',
  'dc_cluster_group_slo',
  'ce_site_reference',

  // Selector references (name/namespace pattern)
  'client_selector',
  'server_selector',
  'site_selector',
  'prefix_selector',
  'proxy_label_selector',
  'ip_prefix_set',

  // Advertise location references
  'public_ip',
  'where',

  // DR references
  'sli_to_global_dr',
  'slo_to_global_dr',

  // Nested reference blocks (commonly found inside other blocks)
  'policies',  // inside active_service_policies
]);

/**
 * Description patterns that indicate a reference block.
 * These are exact patterns from F5XC API documentation.
 */
const REFERENCE_DESCRIPTION_PATTERNS = [
  // Exact F5XC pattern for object references
  'type establishes a direct reference from one object',
  'establishes a direct reference',
  'such a reference is in form of tenant/namespace/name',
  'reference is in form of tenant/namespace/name',

  // Simpler reference patterns
  'reference to',
  'references to',
  'reference object',
  'refers to another',

  // Configuration object reference patterns
  'configuration object(e.g. virtual_host) refers to another',
  'when a configuration object',
];

/**
 * Check if an attribute is a reference block (only accepts name/namespace/tenant)
 *
 * Detection strategy (in order of precedence):
 * 1. Known reference block names (deterministic, hardcoded list)
 * 2. Description pattern matching (F5XC documentation patterns)
 * 3. Type + block flag analysis
 *
 * This function is critical for preventing AI trial-and-error by correctly
 * identifying blocks that ONLY accept name/namespace/tenant.
 */
function isReferenceBlock(attrMeta: {
  type: string;
  is_block?: boolean;
  description?: string;
}, fieldName?: string): boolean {
  // Strategy 1: Check if it's a known reference block (most reliable)
  if (fieldName && KNOWN_REFERENCE_BLOCKS.has(fieldName)) {
    return true;
  }

  // Strategy 2: Must be a block type for description-based detection
  if (!attrMeta.is_block && attrMeta.type !== 'object' && attrMeta.type !== 'list') {
    return false;
  }

  // Strategy 3: Check description for F5XC reference patterns
  const desc = (attrMeta.description || '').toLowerCase();

  // Check against all known reference patterns
  for (const pattern of REFERENCE_DESCRIPTION_PATTERNS) {
    if (desc.includes(pattern.toLowerCase())) {
      return true;
    }
  }

  return false;
}

/**
 * Calculate Levenshtein distance between two strings
 */
function levenshteinDistance(str1: string, str2: string): number {
  const m = str1.length;
  const n = str2.length;

  // Create distance matrix
  const dp: number[][] = Array(m + 1).fill(null).map(() => Array(n + 1).fill(0));

  // Initialize base cases
  for (let i = 0; i <= m; i++) dp[i][0] = i;
  for (let j = 0; j <= n; j++) dp[0][j] = j;

  // Fill in the rest of the matrix
  for (let i = 1; i <= m; i++) {
    for (let j = 1; j <= n; j++) {
      if (str1[i - 1] === str2[j - 1]) {
        dp[i][j] = dp[i - 1][j - 1];
      } else {
        dp[i][j] = 1 + Math.min(
          dp[i - 1][j],     // deletion
          dp[i][j - 1],     // insertion
          dp[i - 1][j - 1], // substitution
        );
      }
    }
  }

  return dp[m][n];
}

/**
 * Find similar field names for "did you mean?" suggestions
 */
function findSimilarFieldNames(typo: string, validFields: string[], maxDistance: number = 3): string[] {
  return validFields
    .map(f => ({ field: f, distance: levenshteinDistance(typo.toLowerCase(), f.toLowerCase()) }))
    .filter(x => x.distance <= maxDistance && x.distance > 0)
    .sort((a, b) => a.distance - b.distance)
    .slice(0, 3)
    .map(x => x.field);
}

/**
 * Get the resource that a reference block points to.
 * Comprehensive mapping for all known reference blocks.
 */
function getReferenceResourceName(fieldName: string): string {
  // Comprehensive reference field mappings
  const mappings: Record<string, string> = {
    // Security resources
    'app_firewall': 'f5xc_app_firewall',
    'waf': 'f5xc_app_firewall',
    'bot_defense_policy': 'f5xc_bot_defense_policy',
    'malicious_user_mitigation': 'f5xc_malicious_user_mitigation',
    'usb_policy': 'f5xc_usb_policy',

    // Load balancer and pool resources
    'healthcheck': 'f5xc_healthcheck',
    'origin_pool': 'f5xc_origin_pool',
    'pool': 'f5xc_origin_pool',
    'cluster': 'f5xc_cluster',

    // Policy resources
    'service_policy': 'f5xc_service_policy',
    'policies': 'f5xc_service_policy',
    'rate_limiter': 'f5xc_rate_limiter',
    'policer': 'f5xc_policer',

    // Identity resources
    'user_identification': 'f5xc_user_identification',
    'authentication': 'f5xc_authentication',

    // API resources
    'api_definition': 'f5xc_api_definition',
    'api_definitions': 'f5xc_api_definition',

    // Certificate resources
    'certificate': 'f5xc_certificate',
    'certificates': 'f5xc_certificate',
    'crl': 'f5xc_crl',
    'trusted_ca': 'f5xc_trusted_ca_list',
    'trusted_ca_list': 'f5xc_trusted_ca_list',

    // Infrastructure resources
    'site': 'f5xc_site',
    'virtual_site': 'f5xc_virtual_site',
    'virtual_network': 'f5xc_virtual_network',
    'segment': 'f5xc_segment',
    'network_connector': 'f5xc_network_connector',
    'cloud_credentials': 'f5xc_cloud_credentials',
    'aws_cred': 'f5xc_cloud_credentials',
    'azure_cred': 'f5xc_cloud_credentials',
    'gcp_cred': 'f5xc_cloud_credentials',
    'k8s_cluster': 'f5xc_cluster',
    'log_receiver': 'f5xc_log_receiver',

    // Site group resources
    'dc_cluster_group': 'f5xc_dc_cluster_group',
    'dc_cluster_group_inside': 'f5xc_dc_cluster_group',
    'dc_cluster_group_sli': 'f5xc_dc_cluster_group',
    'dc_cluster_group_slo': 'f5xc_dc_cluster_group',

    // IP/Prefix resources
    'ip_prefix_set': 'f5xc_ip_prefix_set',
    'public_ip': 'f5xc_cloud_elastic_ip',
  };

  return mappings[fieldName] || `f5xc_${fieldName}`;
}

// =============================================================================
// TOOL DEFINITION
// =============================================================================

export const METADATA_TOOL_DEFINITION = {
  name: 'f5xc_terraform_metadata',
  description: 'Query resource metadata for deterministic Terraform configuration generation. Operations: oneof (mutually exclusive fields), validation (regex/range patterns), defaults, enums, attribute (full metadata), requires_replace, tier (subscription requirements), dependencies (creation order), troubleshoot (error remediation), syntax (correct Terraform block vs attribute syntax), validate (check config for syntax errors), example (generate complete working examples), mistakes (common errors and fixes), summary.',
};
