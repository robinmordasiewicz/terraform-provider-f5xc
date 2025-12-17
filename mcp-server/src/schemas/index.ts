/**
 * Zod schemas for MCP tool input validation
 */

import { z } from 'zod';
import { ResponseFormat } from '../types.js';

// Common pagination schema
export const PaginationSchema = z.object({
  limit: z.number()
    .int()
    .min(1)
    .max(100)
    .default(20)
    .describe('Maximum results to return (1-100)'),
  offset: z.number()
    .int()
    .min(0)
    .default(0)
    .describe('Number of results to skip for pagination'),
}).strict();

// Response format schema
export const ResponseFormatSchema = z.nativeEnum(ResponseFormat)
  .default(ResponseFormat.MARKDOWN)
  .describe('Output format: "markdown" for human-readable or "json" for machine-readable');

// Search documentation schema
export const SearchDocsSchema = z.object({
  query: z.string()
    .min(1, 'Query is required')
    .max(200, 'Query must not exceed 200 characters')
    .describe('Search query to find in documentation (resource names, attributes, descriptions)'),
  type: z.enum(['resource', 'data-source', 'function', 'guide'])
    .optional()
    .describe('Filter by documentation type'),
  limit: z.number()
    .int()
    .min(1)
    .max(50)
    .default(20)
    .describe('Maximum results to return'),
  response_format: ResponseFormatSchema,
}).strict();

// Get documentation schema
export const GetDocSchema = z.object({
  name: z.string()
    .min(1, 'Resource name is required')
    .describe('Name of the resource/data-source/function (e.g., "http_loadbalancer", "namespace", "blindfold")'),
  type: z.enum(['resource', 'data-source', 'function', 'guide'])
    .optional()
    .describe('Type of documentation to retrieve'),
  response_format: ResponseFormatSchema,
}).strict();

// List documentation schema
export const ListDocsSchema = z.object({
  type: z.enum(['resource', 'data-source', 'function', 'guide'])
    .optional()
    .describe('Filter by documentation type'),
  response_format: ResponseFormatSchema,
}).strict();

// Search API specs schema
export const SearchApiSpecsSchema = z.object({
  query: z.string()
    .min(1, 'Query is required')
    .max(200, 'Query must not exceed 200 characters')
    .describe('Search query for API specifications (schema names, endpoints)'),
  limit: z.number()
    .int()
    .min(1)
    .max(50)
    .default(20)
    .describe('Maximum results to return'),
  response_format: ResponseFormatSchema,
}).strict();

// Get API spec schema
export const GetApiSpecSchema = z.object({
  name: z.string()
    .min(1, 'Spec name is required')
    .describe('Name of the API specification (e.g., "http_loadbalancer", "app_firewall", "namespace")'),
  include_paths: z.boolean()
    .default(true)
    .describe('Whether to include endpoint paths in the response'),
  include_definitions: z.boolean()
    .default(false)
    .describe('Whether to include schema definitions (can be large)'),
  response_format: ResponseFormatSchema,
}).strict();

// Find endpoints schema
export const FindEndpointsSchema = z.object({
  pattern: z.string()
    .min(1, 'Pattern is required')
    .describe('URL pattern to search for (e.g., "/api/config/", "namespaces", "http_loadbalancers")'),
  method: z.enum(['GET', 'POST', 'PUT', 'DELETE', 'PATCH'])
    .optional()
    .describe('Filter by HTTP method'),
  limit: z.number()
    .int()
    .min(1)
    .max(100)
    .default(20)
    .describe('Maximum results to return'),
  response_format: ResponseFormatSchema,
}).strict();

// Get schema definition schema
export const GetSchemaDefSchema = z.object({
  spec_name: z.string()
    .min(1, 'Spec name is required')
    .describe('Name of the API specification containing the definition'),
  definition_name: z.string()
    .min(1, 'Definition name is required')
    .describe('Name of the schema definition to retrieve'),
  response_format: ResponseFormatSchema,
}).strict();

// List definitions schema
export const ListDefinitionsSchema = z.object({
  spec_name: z.string()
    .min(1, 'Spec name is required')
    .describe('Name of the API specification'),
  response_format: ResponseFormatSchema,
}).strict();

// Get summary schema
export const GetSummarySchema = z.object({
  response_format: ResponseFormatSchema,
}).strict();

// Get subscription info schema
export const GetSubscriptionInfoSchema = z.object({
  resource: z.string()
    .optional()
    .describe('Resource name to check (omit to get all Advanced tier resources)'),
  tier: z.enum(['STANDARD', 'ADVANCED'])
    .optional()
    .describe('Filter by subscription tier'),
  response_format: ResponseFormatSchema,
}).strict();

// Get property subscription info schema
export const GetPropertySubscriptionInfoSchema = z.object({
  resource: z.string()
    .min(1, 'Resource name is required')
    .describe('Resource name (e.g., "http_loadbalancer", "app_firewall")'),
  property: z.string()
    .optional()
    .describe('Property name to check (e.g., "enable_malicious_user_detection"). If omitted, returns all advanced features for the resource.'),
  response_format: ResponseFormatSchema,
}).strict();

// =============================================================================
// ADDON SERVICE SCHEMAS
// =============================================================================

// List addon services schema
export const ListAddonServicesSchema = z.object({
  tier: z.enum(['STANDARD', 'ADVANCED', 'PREMIUM'])
    .optional()
    .describe('Filter by subscription tier required'),
  activation_type: z.enum(['self', 'managed'])
    .optional()
    .describe('Filter by activation type (self = user can activate directly, managed = requires sales contact)'),
  response_format: ResponseFormatSchema,
}).strict();

// Check addon activation schema
export const CheckAddonActivationSchema = z.object({
  addon_service: z.string()
    .min(1, 'Addon service name is required')
    .describe('Name of the addon service (e.g., "bot_defense", "client_side_defense", "api_discovery")'),
  response_format: ResponseFormatSchema,
}).strict();

// Get addon workflow schema
export const GetAddonWorkflowSchema = z.object({
  addon_service: z.string()
    .min(1, 'Addon service name is required')
    .describe('Name of the addon service (e.g., "bot_defense", "client_side_defense")'),
  activation_type: z.enum(['self', 'partial', 'managed'])
    .optional()
    .describe('Override activation type for workflow (auto-detected if not specified)'),
  response_format: ResponseFormatSchema,
}).strict();

// Export type inference helpers
export type SearchDocsInput = z.infer<typeof SearchDocsSchema>;
export type GetDocInput = z.infer<typeof GetDocSchema>;
export type ListDocsInput = z.infer<typeof ListDocsSchema>;
export type SearchApiSpecsInput = z.infer<typeof SearchApiSpecsSchema>;
export type GetApiSpecInput = z.infer<typeof GetApiSpecSchema>;
export type FindEndpointsInput = z.infer<typeof FindEndpointsSchema>;
export type GetSchemaDefInput = z.infer<typeof GetSchemaDefSchema>;
export type ListDefinitionsInput = z.infer<typeof ListDefinitionsSchema>;
export type GetSummaryInput = z.infer<typeof GetSummarySchema>;
export type GetSubscriptionInfoInput = z.infer<typeof GetSubscriptionInfoSchema>;
export type GetPropertySubscriptionInfoInput = z.infer<typeof GetPropertySubscriptionInfoSchema>;
export type ListAddonServicesInput = z.infer<typeof ListAddonServicesSchema>;
export type CheckAddonActivationInput = z.infer<typeof CheckAddonActivationSchema>;
export type GetAddonWorkflowInput = z.infer<typeof GetAddonWorkflowSchema>;
