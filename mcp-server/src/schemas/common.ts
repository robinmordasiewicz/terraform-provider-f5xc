/**
 * Common schemas, descriptions, and utilities for token optimization
 *
 * This module provides:
 * - Shared parameter descriptions to reduce token usage
 * - Description optimization function for truncation
 * - Common Zod schemas used across consolidated tools
 */

import { z } from 'zod';
import { ResponseFormat } from '../types.js';

// =============================================================================
// COMMON PARAMETER DESCRIPTIONS
// =============================================================================

/**
 * Shared parameter descriptions for token optimization.
 * Use these instead of inline descriptions for common parameters.
 */
export const COMMON_PARAM_DESCRIPTIONS: Record<string, string> = {
  // Namespacing
  namespace: "Target namespace (e.g., 'default')",
  'metadata.namespace': 'Target namespace for the resource',

  // Naming
  name: 'Resource name',
  'metadata.name': 'Resource name identifier',
  resource_name: 'Terraform resource name',
  spec_name: "OpenAPI spec name (e.g., 'http_loadbalancer')",
  definition_name: 'Schema definition name',
  service_name: 'Addon service identifier',
  addon_service: 'Addon service name',

  // Operations
  operation: 'Action to perform',
  query: 'Search query string',
  pattern: 'URL pattern to search',

  // Output control
  response_format: "Output: 'markdown' or 'json'",
  limit: 'Max results (1-50, default: 20)',

  // Type filters
  type: 'Filter: resource|data-source|function|guide',
  tier: 'Subscription tier filter',
  method: 'HTTP method filter',

  // Flags
  include_paths: 'Include API endpoint paths',
  include_definitions: 'Include schema definitions',
  verbose: 'Return detailed information',
};

// =============================================================================
// DESCRIPTION OPTIMIZATION
// =============================================================================

/**
 * Optimizes parameter descriptions for token efficiency.
 *
 * Strategy:
 * 1. Use shared description if available
 * 2. Use first sentence if under 100 chars
 * 3. Truncate with ellipsis if over 100 chars
 *
 * @param name - Parameter name (may match shared descriptions)
 * @param description - Original description
 * @returns Optimized description (max 100 chars)
 */
export function optimizeDescription(name: string, description: string): string {
  // Use common description if available
  if (COMMON_PARAM_DESCRIPTIONS[name]) {
    return COMMON_PARAM_DESCRIPTIONS[name];
  }

  // If already short, return as-is
  if (description.length <= 100) {
    return description;
  }

  // Try to use first sentence if it fits
  const firstSentence = description.split(/[.\n]/)[0];
  if (firstSentence && firstSentence.length <= 100) {
    return firstSentence;
  }

  // Truncate with ellipsis
  return description.slice(0, 97) + '...';
}

// =============================================================================
// COMMON SCHEMAS
// =============================================================================

/**
 * Response format schema - shared across all tools
 */
export const ResponseFormatSchema = z.nativeEnum(ResponseFormat)
  .default(ResponseFormat.MARKDOWN)
  .describe(COMMON_PARAM_DESCRIPTIONS.response_format);

/**
 * Document type enum - shared across docs tools
 */
export const DocTypeEnum = z.enum(['resource', 'data-source', 'function', 'guide']);

/**
 * HTTP method enum - shared across API tools
 */
export const HttpMethodEnum = z.enum(['GET', 'POST', 'PUT', 'DELETE', 'PATCH']);

/**
 * Subscription tier enum - shared across subscription and addon tools
 */
export const SubscriptionTierEnum = z.enum(['STANDARD', 'ADVANCED', 'PREMIUM']);

// =============================================================================
// CONSOLIDATED TOOL SCHEMAS
// =============================================================================

/**
 * Discovery meta-tool schema
 * Enables lazy loading of full tool schemas
 */
export const DiscoverSchema = z.object({
  category: z.enum(['docs', 'api', 'subscription', 'addon', 'all'])
    .optional()
    .default('all')
    .describe('Filter tools by category'),
  verbose: z.boolean()
    .optional()
    .default(false)
    .describe(COMMON_PARAM_DESCRIPTIONS.verbose),
  response_format: ResponseFormatSchema,
}).strict();

/**
 * Consolidated documentation tool schema
 * Replaces: search_docs, get_doc, list_docs
 */
export const DocsSchema = z.object({
  operation: z.enum(['search', 'get', 'list'])
    .describe(COMMON_PARAM_DESCRIPTIONS.operation),
  query: z.string()
    .min(1)
    .max(200)
    .optional()
    .describe(COMMON_PARAM_DESCRIPTIONS.query),
  name: z.string()
    .min(1)
    .optional()
    .describe(COMMON_PARAM_DESCRIPTIONS.name),
  type: DocTypeEnum
    .optional()
    .describe(COMMON_PARAM_DESCRIPTIONS.type),
  limit: z.number()
    .int()
    .min(1)
    .max(50)
    .optional()
    .default(20)
    .describe(COMMON_PARAM_DESCRIPTIONS.limit),
  response_format: ResponseFormatSchema,
}).strict();

/**
 * Consolidated API tool schema
 * Replaces: search_api_specs, get_api_spec, find_endpoints, get_schema_definition, list_definitions
 */
export const ApiSchema = z.object({
  operation: z.enum(['search', 'get', 'find_endpoints', 'get_definition', 'list_definitions'])
    .describe(COMMON_PARAM_DESCRIPTIONS.operation),
  query: z.string()
    .min(1)
    .max(200)
    .optional()
    .describe(COMMON_PARAM_DESCRIPTIONS.query),
  pattern: z.string()
    .min(1)
    .optional()
    .describe(COMMON_PARAM_DESCRIPTIONS.pattern),
  name: z.string()
    .min(1)
    .optional()
    .describe(COMMON_PARAM_DESCRIPTIONS.spec_name),
  spec_name: z.string()
    .min(1)
    .optional()
    .describe(COMMON_PARAM_DESCRIPTIONS.spec_name),
  definition_name: z.string()
    .min(1)
    .optional()
    .describe(COMMON_PARAM_DESCRIPTIONS.definition_name),
  method: HttpMethodEnum
    .optional()
    .describe(COMMON_PARAM_DESCRIPTIONS.method),
  include_paths: z.boolean()
    .optional()
    .default(true)
    .describe(COMMON_PARAM_DESCRIPTIONS.include_paths),
  include_definitions: z.boolean()
    .optional()
    .default(false)
    .describe(COMMON_PARAM_DESCRIPTIONS.include_definitions),
  limit: z.number()
    .int()
    .min(1)
    .max(100)
    .optional()
    .default(20)
    .describe(COMMON_PARAM_DESCRIPTIONS.limit),
  response_format: ResponseFormatSchema,
}).strict();

/**
 * Consolidated subscription tool schema
 * Replaces: get_subscription_info, get_property_subscription_info
 */
export const SubscriptionSchema = z.object({
  operation: z.enum(['resource', 'property'])
    .describe(COMMON_PARAM_DESCRIPTIONS.operation),
  resource_name: z.string()
    .optional()
    .describe(COMMON_PARAM_DESCRIPTIONS.resource_name),
  property_path: z.string()
    .optional()
    .describe('Property path to check (e.g., "enable_malicious_user_detection")'),
  tier: SubscriptionTierEnum
    .optional()
    .describe(COMMON_PARAM_DESCRIPTIONS.tier),
  response_format: ResponseFormatSchema,
}).strict();

/**
 * Consolidated addon tool schema
 * Replaces: addon_list_services, addon_check_activation, addon_activation_workflow
 */
export const AddonSchema = z.object({
  operation: z.enum(['list', 'check', 'workflow'])
    .describe(COMMON_PARAM_DESCRIPTIONS.operation),
  service_name: z.string()
    .min(1)
    .optional()
    .describe(COMMON_PARAM_DESCRIPTIONS.service_name),
  tier: SubscriptionTierEnum
    .optional()
    .describe(COMMON_PARAM_DESCRIPTIONS.tier),
  activation_type: z.enum(['self', 'partial', 'managed'])
    .optional()
    .describe('Filter or override activation type'),
  response_format: ResponseFormatSchema,
}).strict();

/**
 * Consolidated metadata tool schema
 * Provides access to resource metadata for deterministic AI configuration generation
 */
export const MetadataSchema = z.object({
  operation: z.enum(['oneof', 'validation', 'defaults', 'enums', 'attribute', 'requires_replace', 'tier', 'dependencies', 'troubleshoot', 'summary', 'syntax'])
    .describe(COMMON_PARAM_DESCRIPTIONS.operation),
  resource: z.string()
    .min(1)
    .optional()
    .describe('Resource name (e.g., "http_loadbalancer", "namespace")'),
  attribute: z.string()
    .min(1)
    .optional()
    .describe('Attribute name for attribute operation'),
  pattern: z.string()
    .min(1)
    .optional()
    .describe('Validation pattern name (e.g., "name", "domain", "port")'),
  tier: SubscriptionTierEnum
    .optional()
    .describe('Filter resources by subscription tier'),
  error_code: z.string()
    .min(1)
    .optional()
    .describe('Error code for troubleshoot operation (e.g., "NOT_FOUND", "FORBIDDEN")'),
  error_message: z.string()
    .min(1)
    .optional()
    .describe('Error message pattern for troubleshoot operation'),
  response_format: ResponseFormatSchema,
}).strict();

// =============================================================================
// TYPE EXPORTS
// =============================================================================

export type DiscoverInput = z.infer<typeof DiscoverSchema>;
export type DocsInput = z.infer<typeof DocsSchema>;
export type ApiInput = z.infer<typeof ApiSchema>;
export type SubscriptionInput = z.infer<typeof SubscriptionSchema>;
export type AddonInput = z.infer<typeof AddonSchema>;
export type MetadataInput = z.infer<typeof MetadataSchema>;
