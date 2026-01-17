// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

/**
 * Metadata Service for MCP Server
 *
 * Loads and provides access to resource metadata extracted during provider generation.
 * This enables AI assistants to make deterministic configuration decisions.
 *
 * Metadata includes:
 * - OneOf groups (mutually exclusive fields)
 * - Validation patterns (regex/range rules)
 * - Default values and recommendations
 * - Attribute metadata (types, requirements)
 */

import { readFileSync, existsSync } from 'fs';
import { join, dirname } from 'path';
import { fileURLToPath } from 'url';

// Get paths
const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

const PACKAGE_ROOT = join(__dirname, '..', '..'); // mcp-server/
const BUNDLED_METADATA = join(PACKAGE_ROOT, 'dist', 'metadata');
const PROJECT_ROOT = join(PACKAGE_ROOT, '..'); // terraform-provider-f5xc/
const PROJECT_METADATA = join(PROJECT_ROOT, 'tools', 'metadata');

// =============================================================================
// TYPES
// =============================================================================

export interface MetadataCollection {
  generated_at: string;
  version: string;
  resources: Record<string, ResourceMetadata>;
}

export interface ResourceMetadata {
  description: string;
  category?: string;
  tier?: string;
  import_format?: 'namespace/name' | 'name';
  oneof_groups?: Record<string, OneOfGroupInfo>;
  attributes: Record<string, AttributeMetadata>;
  dependencies?: DependencyInfo;
}

export interface OneOfGroupInfo {
  fields: string[];
  description?: string;
  default?: string;
}

export interface AttributeMetadata {
  type: string;
  required: boolean;
  optional?: boolean;
  computed?: boolean;
  sensitive?: boolean;
  is_block?: boolean;
  plan_modifier?: string;
  validation?: string;
  enum?: string[];
  default?: unknown;
  oneof_group?: string;
  description: string;
  server_default?: boolean;
  recommended_value?: string | number | boolean;
}

export interface DependencyInfo {
  references?: string[];
  referenced_by?: string[];
}

export interface ValidationPatterns {
  version: string;
  patterns: Record<string, ValidationPattern>;
}

export interface ValidationPattern {
  type: 'regex' | 'range';
  pattern?: string;
  min?: number;
  max?: number;
  description: string;
  examples?: string[];
  invalid?: string[];
}

export interface ErrorPatterns {
  version: string;
  description: string;
  error_codes: Record<string, ErrorCodeInfo>;
  common_patterns: Record<string, CommonErrorPattern>;
  diagnostic_tips: Record<string, string>;
}

export interface ErrorCodeInfo {
  status: number;
  causes: string[];
  remediation: string[];
  operations: string[];
}

export interface CommonErrorPattern {
  pattern: string;
  cause: string;
  remediation: string;
}

// Import operation metadata types from types.ts
import type {
  OperationsMetadataCollection,
  ResourceOperationInfo,
  OperationMetadata,
  DangerLevel,
  SideEffectsInfo,
  BestPracticesInfo,
  GuidedWorkflowInfo,
  ResponseTimeInfo,
} from '../types.js';

// =============================================================================
// CACHING
// =============================================================================

let resourceMetadataCache: MetadataCollection | null = null;
let validationPatternsCache: ValidationPatterns | null = null;
let errorPatternsCache: ErrorPatterns | null = null;
let operationsMetadataCache: OperationsMetadataCollection | null = null;

// =============================================================================
// PATH RESOLUTION
// =============================================================================

/**
 * Get the metadata directory path
 * Uses bundled metadata if available (npm), otherwise project metadata (development)
 */
function getMetadataPath(): string {
  if (existsSync(BUNDLED_METADATA)) {
    return BUNDLED_METADATA;
  }
  return PROJECT_METADATA;
}

// =============================================================================
// LOADING FUNCTIONS
// =============================================================================

/**
 * Load resource metadata from JSON file
 */
export function loadResourceMetadata(): MetadataCollection | null {
  if (resourceMetadataCache) {
    return resourceMetadataCache;
  }

  const metadataPath = join(getMetadataPath(), 'resource-metadata.json');

  if (!existsSync(metadataPath)) {
    console.error(`Resource metadata not found: ${metadataPath}`);
    return null;
  }

  try {
    const content = readFileSync(metadataPath, 'utf-8');
    resourceMetadataCache = JSON.parse(content) as MetadataCollection;
    return resourceMetadataCache;
  } catch (error) {
    console.error(`Failed to load resource metadata: ${error}`);
    return null;
  }
}

/**
 * Load validation patterns from JSON file
 */
export function loadValidationPatterns(): ValidationPatterns | null {
  if (validationPatternsCache) {
    return validationPatternsCache;
  }

  const patternsPath = join(getMetadataPath(), 'validation-patterns.json');

  if (!existsSync(patternsPath)) {
    console.error(`Validation patterns not found: ${patternsPath}`);
    return null;
  }

  try {
    const content = readFileSync(patternsPath, 'utf-8');
    validationPatternsCache = JSON.parse(content) as ValidationPatterns;
    return validationPatternsCache;
  } catch (error) {
    console.error(`Failed to load validation patterns: ${error}`);
    return null;
  }
}

/**
 * Load error patterns from JSON file
 */
export function loadErrorPatterns(): ErrorPatterns | null {
  if (errorPatternsCache) {
    return errorPatternsCache;
  }

  const errorPath = join(getMetadataPath(), 'error-patterns.json');

  if (!existsSync(errorPath)) {
    // Error patterns are optional
    return null;
  }

  try {
    const content = readFileSync(errorPath, 'utf-8');
    errorPatternsCache = JSON.parse(content) as ErrorPatterns;
    return errorPatternsCache;
  } catch (error) {
    console.error(`Failed to load error patterns: ${error}`);
    return null;
  }
}

/**
 * Load operations metadata from JSON file (v2.0.33 extensions)
 */
export function loadOperationsMetadata(): OperationsMetadataCollection | null {
  if (operationsMetadataCache) {
    return operationsMetadataCache;
  }

  const operationsPath = join(getMetadataPath(), 'operations-metadata.json');

  if (!existsSync(operationsPath)) {
    // Operations metadata is optional (may not be generated yet)
    return null;
  }

  try {
    const content = readFileSync(operationsPath, 'utf-8');
    operationsMetadataCache = JSON.parse(content) as OperationsMetadataCollection;
    return operationsMetadataCache;
  } catch (error) {
    console.error(`Failed to load operations metadata: ${error}`);
    return null;
  }
}

// =============================================================================
// QUERY FUNCTIONS
// =============================================================================

/**
 * Get metadata for a specific resource
 */
export function getResourceMetadata(resourceName: string): ResourceMetadata | null {
  const metadata = loadResourceMetadata();
  if (!metadata) return null;

  // Try exact match first, then with underscores
  const normalized = resourceName.toLowerCase().replace(/-/g, '_');
  return metadata.resources[normalized] || metadata.resources[resourceName] || null;
}

/**
 * Get OneOf groups for a resource
 */
export function getResourceOneOfGroups(resourceName: string): Record<string, OneOfGroupInfo> {
  const resource = getResourceMetadata(resourceName);
  return resource?.oneof_groups || {};
}

/**
 * Get attributes with defaults for a resource
 */
export function getAttributesWithDefaults(resourceName: string): Record<string, unknown> {
  const resource = getResourceMetadata(resourceName);
  if (!resource) return {};

  const defaults: Record<string, unknown> = {};
  for (const [name, attr] of Object.entries(resource.attributes)) {
    if (attr.default !== undefined) {
      defaults[name] = attr.default;
    }
  }
  return defaults;
}

/**
 * Get attributes with enum constraints for a resource
 */
export function getAttributesWithEnums(resourceName: string): Record<string, string[]> {
  const resource = getResourceMetadata(resourceName);
  if (!resource) return {};

  const enums: Record<string, string[]> = {};
  for (const [name, attr] of Object.entries(resource.attributes)) {
    if (attr.enum && attr.enum.length > 0) {
      enums[name] = attr.enum;
    }
  }
  return enums;
}

/**
 * Get required attributes for a resource
 */
export function getRequiredAttributes(resourceName: string): string[] {
  const resource = getResourceMetadata(resourceName);
  if (!resource) return [];

  return Object.entries(resource.attributes)
    .filter(([, attr]) => attr.required)
    .map(([name]) => name);
}

/**
 * Get attributes that require replacement when changed
 */
export function getRequiresReplaceAttributes(resourceName: string): string[] {
  const resource = getResourceMetadata(resourceName);
  if (!resource) return [];

  return Object.entries(resource.attributes)
    .filter(([, attr]) => attr.plan_modifier === 'RequiresReplace')
    .map(([name]) => name);
}

/**
 * Get a specific validation pattern
 */
export function getValidationPattern(patternName: string): ValidationPattern | null {
  const patterns = loadValidationPatterns();
  return patterns?.patterns[patternName] || null;
}

/**
 * Get all validation patterns
 */
export function getAllValidationPatterns(): Record<string, ValidationPattern> {
  const patterns = loadValidationPatterns();
  return patterns?.patterns || {};
}

/**
 * List all available resources in metadata
 */
export function listResourcesWithMetadata(): string[] {
  const metadata = loadResourceMetadata();
  return metadata ? Object.keys(metadata.resources) : [];
}

/**
 * Get metadata summary statistics
 */
export function getMetadataSummary(): {
  resourceCount: number;
  resourcesWithOneOf: number;
  validationPatternCount: number;
  generatedAt: string | null;
} {
  const metadata = loadResourceMetadata();
  const patterns = loadValidationPatterns();

  let resourcesWithOneOf = 0;
  if (metadata) {
    for (const resource of Object.values(metadata.resources)) {
      if (resource.oneof_groups && Object.keys(resource.oneof_groups).length > 0) {
        resourcesWithOneOf++;
      }
    }
  }

  return {
    resourceCount: metadata ? Object.keys(metadata.resources).length : 0,
    resourcesWithOneOf,
    validationPatternCount: patterns ? Object.keys(patterns.patterns).length : 0,
    generatedAt: metadata?.generated_at || null,
  };
}

/**
 * Check if metadata is available
 */
export function isMetadataAvailable(): boolean {
  return existsSync(join(getMetadataPath(), 'resource-metadata.json'));
}

// =============================================================================
// NEW QUERY FUNCTIONS (Phase 5 additions)
// =============================================================================

/**
 * Get subscription tier for a resource
 */
export function getResourceTier(resourceName: string): string | null {
  const resource = getResourceMetadata(resourceName);
  return resource?.tier || null;
}

/**
 * Get all resources requiring a specific tier
 */
export function getResourcesByTier(tier: string): string[] {
  const metadata = loadResourceMetadata();
  if (!metadata) return [];

  return Object.entries(metadata.resources)
    .filter(([, resource]) => resource.tier?.toUpperCase() === tier.toUpperCase())
    .map(([name]) => name);
}

/**
 * Get import format for a resource
 */
export function getResourceImportFormat(resourceName: string): string | null {
  const resource = getResourceMetadata(resourceName);
  return resource?.import_format || null;
}

/**
 * Get dependencies for a resource
 */
export function getResourceDependencies(resourceName: string): DependencyInfo | null {
  const resource = getResourceMetadata(resourceName);
  return resource?.dependencies || null;
}

/**
 * Compute creation order based on dependencies (topological sort)
 */
export function computeCreationOrder(resourceName: string): string[] {
  const resource = getResourceMetadata(resourceName);
  if (!resource?.dependencies?.references) {
    return [resourceName];
  }

  const visited = new Set<string>();
  const result: string[] = [];

  function visit(name: string) {
    if (visited.has(name)) return;
    visited.add(name);

    const res = getResourceMetadata(name);
    if (res?.dependencies?.references) {
      for (const dep of res.dependencies.references) {
        visit(dep);
      }
    }
    result.push(name);
  }

  visit(resourceName);
  return result;
}

/**
 * Get error code information
 */
export function getErrorCodeInfo(errorCode: string): ErrorCodeInfo | null {
  const patterns = loadErrorPatterns();
  return patterns?.error_codes[errorCode] || null;
}

/**
 * Find error by pattern match
 */
export function findErrorByPattern(errorMessage: string): {
  name: string;
  info: CommonErrorPattern;
} | null {
  const patterns = loadErrorPatterns();
  if (!patterns) return null;

  for (const [name, info] of Object.entries(patterns.common_patterns)) {
    if (errorMessage.toLowerCase().includes(info.pattern.toLowerCase())) {
      return { name, info };
    }
  }
  return null;
}

/**
 * Get all error codes
 */
export function getAllErrorCodes(): Record<string, ErrorCodeInfo> {
  const patterns = loadErrorPatterns();
  return patterns?.error_codes || {};
}

/**
 * Get diagnostic tips
 */
export function getDiagnosticTips(): Record<string, string> {
  const patterns = loadErrorPatterns();
  return patterns?.diagnostic_tips || {};
}

/**
 * Enhanced metadata summary with tier and dependency statistics
 */
export function getEnhancedMetadataSummary(): {
  resourceCount: number;
  resourcesWithOneOf: number;
  resourcesWithDependencies: number;
  resourcesWithTier: Record<string, number>;
  validationPatternCount: number;
  errorCodeCount: number;
  generatedAt: string | null;
} {
  const metadata = loadResourceMetadata();
  const patterns = loadValidationPatterns();
  const errors = loadErrorPatterns();

  let resourcesWithOneOf = 0;
  let resourcesWithDependencies = 0;
  const resourcesWithTier: Record<string, number> = {
    Standard: 0,
    Advanced: 0,
    Premium: 0,
  };

  if (metadata) {
    for (const resource of Object.values(metadata.resources)) {
      if (resource.oneof_groups && Object.keys(resource.oneof_groups).length > 0) {
        resourcesWithOneOf++;
      }
      if (resource.dependencies?.references && resource.dependencies.references.length > 0) {
        resourcesWithDependencies++;
      }
      if (resource.tier) {
        const tier = resource.tier.charAt(0).toUpperCase() + resource.tier.slice(1).toLowerCase();
        if (tier in resourcesWithTier) {
          resourcesWithTier[tier]++;
        }
      }
    }
  }

  return {
    resourceCount: metadata ? Object.keys(metadata.resources).length : 0,
    resourcesWithOneOf,
    resourcesWithDependencies,
    resourcesWithTier,
    validationPatternCount: patterns ? Object.keys(patterns.patterns).length : 0,
    errorCodeCount: errors ? Object.keys(errors.error_codes).length : 0,
    generatedAt: metadata?.generated_at || null,
  };
}

// =============================================================================
// SYNTAX GUIDANCE FUNCTIONS (Fix for block vs attribute confusion)
// =============================================================================

/**
 * Interface for syntax guidance response
 */
export interface AttributeSyntaxGuidance {
  attributeName: string;
  isBlock: boolean;
  isOneOf: boolean;
  oneOfGroup?: string;
  correctSyntax: string;
  incorrectSyntax: string;
  example: string;
  terraformType: string;
}

/**
 * Interface for complete resource syntax guide
 */
export interface ResourceSyntaxGuide {
  resource: string;
  totalBlocks: number;
  totalAttributes: number;
  totalOneOfGroups: number;
  blocks: string[];
  attributes: string[];
  oneOfGroups: Record<string, string[]>;
  syntaxGuide: string;
}

/**
 * Find which oneof group an attribute belongs to
 */
function findOneOfGroup(
  resourceName: string,
  attributeName: string
): string | undefined {
  const resource = getResourceMetadata(resourceName);
  if (!resource?.oneof_groups) return undefined;

  for (const [groupName, group] of Object.entries(resource.oneof_groups)) {
    if (group.fields.includes(attributeName)) {
      return groupName;
    }
  }
  return undefined;
}

/**
 * Check if a field references an empty object (ioschemaEmpty in OpenAPI)
 * These fields should use empty block syntax: field_name {}
 */
function isEmptyObjectField(
  resourceName: string,
  attributeName: string
): boolean {
  const resource = getResourceMetadata(resourceName);
  if (!resource) return false;

  const attr = resource.attributes[attributeName];
  if (!attr) return false;

  // Oneof fields that reference empty objects should use empty block syntax
  // Check if this is part of a oneof group and is an object type
  if (attr.oneof_group && attr.type === 'object') {
    return true;
  }

  return attr.is_block === true && attr.type === 'object';
}

/**
 * Get syntax guidance for a specific attribute
 * This helps AI assistants generate correct Terraform syntax
 */
export function getAttributeSyntaxGuidance(
  resourceName: string,
  attributeName: string
): AttributeSyntaxGuidance | null {
  const resource = getResourceMetadata(resourceName);
  if (!resource) return null;

  const attr = resource.attributes[attributeName];
  if (!attr) return null;

  const isBlock = attr.is_block || attr.type === 'object';
  const oneOfGroup = attr.oneof_group || findOneOfGroup(resourceName, attributeName);
  const isEmptyBlock = isEmptyObjectField(resourceName, attributeName);

  // Determine Terraform type for display
  let terraformType: string;
  if (isBlock) {
    if (attr.type === 'list') {
      terraformType = 'list(object)';
    } else if (attr.type === 'set') {
      terraformType = 'set(object)';
    } else if (attr.type === 'map') {
      terraformType = 'map(string)';
    } else {
      terraformType = 'object';
    }
  } else {
    terraformType = attr.type;
  }

  // Generate correct syntax
  let correctSyntax: string;
  let example: string;

  if (isBlock) {
    if (isEmptyBlock) {
      correctSyntax = `${attributeName} {}`;
      example = `${attributeName} {}`;
    } else {
      correctSyntax = `${attributeName} { ... }`;
      example = `${attributeName} {
  # nested properties
}`;
    }
  } else {
    if (attr.enum && attr.enum.length > 0) {
      correctSyntax = `${attributeName} = <${attr.enum.join(' | ')}>`;
      example = `${attributeName} = "${attr.enum[0]}"`;
    } else if (attr.type === 'string') {
      correctSyntax = `${attributeName} = "<string>"`;
      example = `${attributeName} = "example-value"`;
    } else if (attr.type === 'number' || attr.type === 'integer') {
      correctSyntax = `${attributeName} = <number>`;
      example = `${attributeName} = 42`;
    } else if (attr.type === 'bool') {
      correctSyntax = `${attributeName} = <true | false>`;
      example = `${attributeName} = true`;
    } else {
      correctSyntax = `${attributeName} = <value>`;
      example = `${attributeName} = <value>`;
    }
  }

  // Generate incorrect syntax warning
  let incorrectSyntax: string;
  if (isBlock) {
    if (isEmptyBlock) {
      incorrectSyntax = `${attributeName} = true  # WRONG - use block syntax, not boolean`;
    } else {
      incorrectSyntax = `${attributeName} = { ... }  # WRONG - use block syntax, not object assignment`;
    }
  } else {
    incorrectSyntax = `${attributeName} {}  # WRONG - attributes use assignment, not blocks`;
  }

  return {
    attributeName,
    isBlock,
    isOneOf: !!oneOfGroup,
    oneOfGroup,
    correctSyntax,
    incorrectSyntax,
    example,
    terraformType,
  };
}

/**
 * Generate a complete syntax guide for a resource
 */
export function generateTerraformSyntaxGuide(
  resourceName: string
): ResourceSyntaxGuide | null {
  const resource = getResourceMetadata(resourceName);
  if (!resource) return null;

  const blocks: string[] = [];
  const attributes: string[] = [];
  const oneOfGroups: Record<string, string[]> = {};

  // Categorize all attributes
  for (const [name, attr] of Object.entries(resource.attributes)) {
    if (attr.is_block || attr.type === 'object' || attr.type === 'list' || attr.type === 'set') {
      blocks.push(name);
    } else {
      attributes.push(name);
    }

    // Track oneof groups
    const group = attr.oneof_group || findOneOfGroup(resourceName, name);
    if (group) {
      if (!oneOfGroups[group]) {
        oneOfGroups[group] = [];
      }
      if (!oneOfGroups[group].includes(name)) {
        oneOfGroups[group].push(name);
      }
    }
  }

  // Generate syntax guide
  const syntaxGuide = generateSyntaxGuideMarkdown(resourceName, resource, blocks, attributes, oneOfGroups);

  return {
    resource: resourceName,
    totalBlocks: blocks.length,
    totalAttributes: attributes.length,
    totalOneOfGroups: Object.keys(oneOfGroups).length,
    blocks,
    attributes,
    oneOfGroups,
    syntaxGuide,
  };
}

/**
 * Generate markdown syntax guide for a resource
 */
function generateSyntaxGuideMarkdown(
  resourceName: string,
  resource: ResourceMetadata,
  blocks: string[],
  attributes: string[],
  oneOfGroups: Record<string, string[]>
): string {
  const lines: string[] = [
    `# Terraform Syntax Guide: ${resourceName}`,
    '',
    '## Overview',
    '',
    `Total: ${blocks.length} block(s), ${attributes.length} attribute(s), ${Object.keys(oneOfGroups).length} oneof group(s)`,
    '',
    '## Block-Type Attributes (use block syntax)',
    '',
    'These attributes require block syntax like `field_name { ... }` or `field_name {}`:',
    '',
  ];

  // List block attributes
  if (blocks.length > 0) {
    for (const block of blocks.sort()) {
      const guidance = getAttributeSyntaxGuidance(resourceName, block);
      if (guidance) {
        lines.push(`### \`${block}\``);
        lines.push('');
        lines.push(`- **Type**: ${guidance.terraformType}`);
        lines.push(`- **Correct Syntax**: \`${guidance.correctSyntax}\``);
        lines.push(`- **Example**: \`${guidance.example}\``);
        if (guidance.isOneOf) {
          lines.push(`- **OneOf Group**: \`${guidance.oneOfGroup}\``);
        }
        lines.push('');
      }
    }
  } else {
    lines.push('_No block-type attributes._');
    lines.push('');
  }

  lines.push('## Simple Attributes (use assignment syntax)');
  lines.push('');
  lines.push('These attributes use simple assignment like `field_name = value`:');
  lines.push('');

  // List simple attributes
  if (attributes.length > 0) {
    for (const attrName of attributes.sort()) {
      const guidance = getAttributeSyntaxGuidance(resourceName, attrName);
      const attr = resource.attributes[attrName];
      if (guidance) {
        lines.push(`### \`${attrName}\``);
        lines.push('');
        lines.push(`- **Type**: ${guidance.terraformType}`);
        lines.push(`- **Correct Syntax**: \`${guidance.correctSyntax}\``);
        lines.push(`- **Example**: \`${guidance.example}\``);
        if (attr.enum && attr.enum.length > 0) {
          lines.push(`- **Valid Values**: ${attr.enum.map((v: string) => `\`${v}\``).join(', ')}`);
        }
        lines.push('');
      }
    }
  } else {
    lines.push('_No simple attributes._');
    lines.push('');
  }

  // Document oneof groups
  if (Object.keys(oneOfGroups).length > 0) {
    lines.push('## OneOf Groups (Mutually Exclusive)');
    lines.push('');
    lines.push('These field groups are **mutually exclusive** - only one can be set:');
    lines.push('');

    for (const [groupName, fields] of Object.entries(oneOfGroups)) {
      lines.push(`### \`${groupName}\``);
      lines.push('');
      lines.push(`**Choose ONE**: ${fields.map(f => `\`${f}\``).join(', ')}`);
      lines.push('');

      // Show example for each option
      lines.push('**Examples**:');
      lines.push('');
      for (const field of fields) {
        const guidance = getAttributeSyntaxGuidance(resourceName, field);
        if (guidance) {
          lines.push(`\`\`\`hcl`);
          lines.push(`# ${field}`);
          lines.push(guidance.example);
          lines.push(`\`\`\``);
          lines.push('');
        }
      }
    }
  }

  return lines.join('\n');
}

/**
 * Get all attributes that are oneof choices with empty block syntax
 * Useful for detecting common mistakes
 */
export function getOneOfBlockAttributes(resourceName: string): string[] {
  const resource = getResourceMetadata(resourceName);
  if (!resource) return [];

  const result: string[] = [];
  for (const [name, attr] of Object.entries(resource.attributes)) {
    if (attr.oneof_group && attr.type === 'object') {
      result.push(name);
    }
  }
  return result;
}

/**
 * Validate a Terraform configuration snippet for common syntax errors
 */
export function validateConfigurationSyntax(
  resourceName: string,
  configurationSnippet: string
): {
  valid: boolean;
  errors: string[];
  suggestions: string[];
} {
  const resource = getResourceMetadata(resourceName);
  if (!resource) {
    return {
      valid: false,
      errors: [`Resource '${resourceName}' not found in metadata`],
      suggestions: [],
    };
  }

  const errors: string[] = [];
  const suggestions: string[] = [];

  // Check for common mistakes with block-type attributes
  for (const [name, attr] of Object.entries(resource.attributes)) {
    const isBlock = attr.is_block || attr.type === 'object' || attr.type === 'list' || attr.type === 'set';
    const isOneOfBlock = attr.oneof_group && attr.type === 'object';

    if (isBlock) {
      // Check for incorrect boolean assignment
      const booleanAssignmentPattern = new RegExp(`\\b${name}\\s*=\\s*(true|false)`, 'gi');
      if (booleanAssignmentPattern.test(configurationSnippet)) {
        const guidance = getAttributeSyntaxGuidance(resourceName, name);
        if (guidance) {
          errors.push(`Incorrect syntax: \`${name}\` is a block-type attribute, not a boolean`);
          suggestions.push(`Use \`${guidance.correctSyntax}\` instead of \`${guidance.incorrectSyntax.split('  #')[0].trim()}\``);
        }
      }

      // Check for object assignment instead of block
      const objectAssignmentPattern = new RegExp(`\\b${name}\\s*=\\s*\\{`, 'gi');
      if (objectAssignmentPattern.test(configurationSnippet) && isOneOfBlock) {
        errors.push(`Incorrect syntax: \`${name}\` should use empty block syntax`);
        suggestions.push(`Use \`${name} {} instead of ${name} = { ... }\``);
      }
    } else {
      // Check for block syntax used for simple attributes
      const blockSyntaxPattern = new RegExp(`\\b${name}\\s*\\{`, 'gi');
      if (blockSyntaxPattern.test(configurationSnippet)) {
        errors.push(`Incorrect syntax: \`${name}\` is a simple attribute, not a block`);
        const guidance = getAttributeSyntaxGuidance(resourceName, name);
        if (guidance) {
          suggestions.push(`Use \`${guidance.correctSyntax}\` instead of block syntax`);
        }
      }
    }
  }

  return {
    valid: errors.length === 0,
    errors,
    suggestions,
  };
}

/**
 * Get a quick reference for all oneof groups in a resource
 */
export function getOneOfGroupsQuickReference(resourceName: string): Record<string, {
  fields: string[];
  default?: string;
  description: string;
  syntax: Record<string, string>;
}> {
  const resource = getResourceMetadata(resourceName);
  if (!resource || !resource.oneof_groups) return {};

  const result: Record<string, {
    fields: string[];
    default?: string;
    description: string;
    syntax: Record<string, string>;
  }> = {};

  for (const [groupName, group] of Object.entries(resource.oneof_groups)) {
    const syntax: Record<string, string> = {};

    for (const field of group.fields) {
      const guidance = getAttributeSyntaxGuidance(resourceName, field);
      if (guidance) {
        syntax[field] = guidance.example;
      }
    }

    result[groupName] = {
      fields: group.fields,
      default: group.default,
      description: group.description || `OneOf group: ${group.fields.join(', ')}`,
      syntax,
    };
  }

  return result;
}

/**
 * Get a quick reference for all block-type attributes across all resources
 * Useful for AI assistants to understand the pattern
 */
export function getAllBlockAttributesQuickReference(): {
  totalResources: number;
  totalBlockAttributes: number;
  resources: Record<string, {
    blockCount: number;
    oneOfGroupsCount: number;
    sampleBlockAttributes: string[];
  }>;
} {
  const metadata = loadResourceMetadata();
  if (!metadata) {
    return {
      totalResources: 0,
      totalBlockAttributes: 0,
      resources: {},
    };
  }

  const resources: Record<string, {
    blockCount: number;
    oneOfGroupsCount: number;
    sampleBlockAttributes: string[];
  }> = {};

  let totalBlockAttributes = 0;

  for (const [resourceName, resource] of Object.entries(metadata.resources)) {
    let blockCount = 0;
    const sampleBlockAttributes: string[] = [];

    for (const [attrName, attr] of Object.entries(resource.attributes)) {
      if (attr.is_block || attr.type === 'object' || attr.type === 'list' || attr.type === 'set') {
        blockCount++;
        totalBlockAttributes++;
        if (sampleBlockAttributes.length < 5) {
          sampleBlockAttributes.push(attrName);
        }
      }
    }

    resources[resourceName] = {
      blockCount,
      oneOfGroupsCount: resource.oneof_groups ? Object.keys(resource.oneof_groups).length : 0,
      sampleBlockAttributes,
    };
  }

  return {
    totalResources: Object.keys(metadata.resources).length,
    totalBlockAttributes,
    resources,
  };
}

/**
 * Get common block attribute patterns across all resources
 * This helps AI assistants understand when to use block syntax
 */
export function getCommonBlockPatterns(): {
  emptyOneOfBlocks: string[];
  nestedBlockAttributes: string[];
  listBlockAttributes: string[];
} {
  const metadata = loadResourceMetadata();
  if (!metadata) {
    return {
      emptyOneOfBlocks: [],
      nestedBlockAttributes: [],
      listBlockAttributes: [],
    };
  }

  const emptyOneOfBlocks: Set<string> = new Set();
  const nestedBlockAttributes: Set<string> = new Set();
  const listBlockAttributes: Set<string> = new Set();

  for (const resource of Object.values(metadata.resources)) {
    for (const [attrName, attr] of Object.entries(resource.attributes)) {
      // Empty oneof blocks (disable_*, enable_* patterns)
      if (attr.oneof_group && attr.type === 'object') {
        emptyOneOfBlocks.add(attrName);
      }
      // Non-empty blocks (with nested properties)
      else if (attr.is_block && attr.type === 'object' && !attr.oneof_group) {
        nestedBlockAttributes.add(attrName);
      }
      // List blocks
      else if (attr.type === 'list' || attr.type === 'set') {
        listBlockAttributes.add(attrName);
      }
    }
  }

  return {
    emptyOneOfBlocks: Array.from(emptyOneOfBlocks).sort(),
    nestedBlockAttributes: Array.from(nestedBlockAttributes).sort(),
    listBlockAttributes: Array.from(listBlockAttributes).sort(),
  };
}

// =============================================================================
// OPERATION METADATA FUNCTIONS (v2.0.33 extensions)
// =============================================================================

/**
 * Get operation metadata for a specific resource
 */
export function getResourceOperationInfo(resourceName: string): ResourceOperationInfo | null {
  const metadata = loadOperationsMetadata();
  if (!metadata) return null;

  const normalized = resourceName.toLowerCase().replace(/-/g, '_');
  return metadata.resources[normalized] || metadata.resources[resourceName] || null;
}

/**
 * Get danger level for a specific operation
 */
export function getOperationDangerLevel(
  resourceName: string,
  operation: 'create' | 'read' | 'update' | 'delete' | 'list',
): DangerLevel | null {
  const resourceOps = getResourceOperationInfo(resourceName);
  return resourceOps?.operations[operation]?.danger_level || null;
}

/**
 * Get side effects for a specific operation
 */
export function getOperationSideEffects(
  resourceName: string,
  operation: 'create' | 'read' | 'update' | 'delete' | 'list',
): SideEffectsInfo | null {
  const resourceOps = getResourceOperationInfo(resourceName);
  return resourceOps?.operations[operation]?.side_effects || null;
}

/**
 * Get response time metrics for a specific operation
 */
export function getOperationResponseTime(
  resourceName: string,
  operation: 'create' | 'read' | 'update' | 'delete' | 'list',
): ResponseTimeInfo | null {
  const resourceOps = getResourceOperationInfo(resourceName);
  return resourceOps?.operations[operation]?.discovered_response_time || null;
}

/**
 * Check if an operation requires confirmation
 */
export function operationRequiresConfirmation(
  resourceName: string,
  operation: 'create' | 'read' | 'update' | 'delete' | 'list',
): boolean {
  const resourceOps = getResourceOperationInfo(resourceName);
  return resourceOps?.operations[operation]?.confirmation_required || false;
}

/**
 * Get required fields for an operation
 */
export function getOperationRequiredFields(
  resourceName: string,
  operation: 'create' | 'read' | 'update' | 'delete' | 'list',
): string[] {
  const resourceOps = getResourceOperationInfo(resourceName);
  return resourceOps?.operations[operation]?.required_fields || [];
}

/**
 * Get best practices for a resource
 */
export function getResourceBestPractices(resourceName: string): BestPracticesInfo | null {
  const resourceOps = getResourceOperationInfo(resourceName);
  return resourceOps?.best_practices || null;
}

/**
 * Get guided workflows for a resource
 */
export function getResourceGuidedWorkflows(resourceName: string): GuidedWorkflowInfo[] {
  const resourceOps = getResourceOperationInfo(resourceName);
  return resourceOps?.guided_workflows || [];
}

/**
 * Get complete operation metadata
 */
export function getOperationMetadataDetail(
  resourceName: string,
  operation: 'create' | 'read' | 'update' | 'delete' | 'list',
): OperationMetadata | null {
  const resourceOps = getResourceOperationInfo(resourceName);
  return resourceOps?.operations[operation] || null;
}

/**
 * Get all high-danger operations across all resources
 */
export function getHighDangerOperations(): Array<{
  resource: string;
  operation: string;
  dangerLevel: DangerLevel;
  confirmationRequired: boolean;
}> {
  const metadata = loadOperationsMetadata();
  if (!metadata) return [];

  const results: Array<{
    resource: string;
    operation: string;
    dangerLevel: DangerLevel;
    confirmationRequired: boolean;
  }> = [];

  for (const [resourceName, resourceOps] of Object.entries(metadata.resources)) {
    for (const [opName, opMeta] of Object.entries(resourceOps.operations)) {
      if (opMeta.danger_level === 'high' || opMeta.danger_level === 'critical') {
        results.push({
          resource: resourceName,
          operation: opName,
          dangerLevel: opMeta.danger_level,
          confirmationRequired: opMeta.confirmation_required || false,
        });
      }
    }
  }

  return results;
}

/**
 * Get operations metadata summary
 */
export function getOperationsMetadataSummary(): {
  available: boolean;
  totalResources: number;
  totalOperations: number;
  byDangerLevel: Record<string, number>;
  confirmationRequired: number;
} {
  const metadata = loadOperationsMetadata();
  if (!metadata) {
    return {
      available: false,
      totalResources: 0,
      totalOperations: 0,
      byDangerLevel: {},
      confirmationRequired: 0,
    };
  }

  const summary = {
    available: true,
    totalResources: Object.keys(metadata.resources).length,
    totalOperations: 0,
    byDangerLevel: {} as Record<string, number>,
    confirmationRequired: 0,
  };

  for (const resourceOps of Object.values(metadata.resources)) {
    for (const opMeta of Object.values(resourceOps.operations)) {
      summary.totalOperations++;
      if (opMeta.danger_level) {
        summary.byDangerLevel[opMeta.danger_level] =
          (summary.byDangerLevel[opMeta.danger_level] || 0) + 1;
      }
      if (opMeta.confirmation_required) {
        summary.confirmationRequired++;
      }
    }
  }

  return summary;
}

/**
 * Check if operations metadata is available
 */
export function isOperationsMetadataAvailable(): boolean {
  const metadata = loadOperationsMetadata();
  return metadata !== null && Object.keys(metadata.resources).length > 0;
}
