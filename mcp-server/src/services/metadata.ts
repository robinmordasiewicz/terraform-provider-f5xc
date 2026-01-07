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

// =============================================================================
// CACHING
// =============================================================================

let resourceMetadataCache: MetadataCollection | null = null;
let validationPatternsCache: ValidationPatterns | null = null;
let errorPatternsCache: ErrorPatterns | null = null;

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
