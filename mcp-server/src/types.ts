// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

/**
 * Type definitions for F5 Distributed Cloud MCP Server
 */

export interface ResourceDoc {
  name: string;
  path: string;
  type: 'resource' | 'data-source' | 'function' | 'guide' | 'provider';
  content?: string;
  // Subscription tier fields
  subscriptionTier?: SubscriptionTier;
  addonService?: string;
  advancedFeatures?: string[];
}

/**
 * F5 Distributed Cloud subscription tiers
 * Note: Only Standard and Advanced tiers are currently available
 * NO_TIER indicates the feature is included in all subscriptions
 */
export type SubscriptionTier = 'NO_TIER' | 'STANDARD' | 'ADVANCED';

/**
 * Service information from F5 XC catalog
 */
export interface ServiceInfo {
  tier: SubscriptionTier;
  display_name: string;
  group_name?: string;
}

/**
 * Resource metadata including subscription requirements
 */
export interface ResourceMetadata {
  service: string;
  minimum_tier: SubscriptionTier;
  advanced_features?: string[];
}

/**
 * Full subscription metadata structure
 */
export interface SubscriptionMetadata {
  generated_at: string;
  source: string;
  services: Record<string, ServiceInfo>;
  resources: Record<string, ResourceMetadata>;
}

export interface ApiSpec {
  name: string;
  path: string;
  schemaName: string;
  content?: OpenAPISpec;
}

export interface OpenAPISpec {
  swagger?: string;
  openapi?: string;
  info?: {
    title?: string;
    description?: string;
    version?: string;
  };
  paths?: Record<string, PathItem>;
  definitions?: Record<string, SchemaDefinition>;
  components?: {
    schemas?: Record<string, SchemaDefinition>;
  };
}

export interface PathItem {
  get?: Operation;
  post?: Operation;
  put?: Operation;
  delete?: Operation;
  patch?: Operation;
}

export interface Operation {
  summary?: string;
  description?: string;
  operationId?: string;
  tags?: string[];
  parameters?: Parameter[];
  requestBody?: RequestBody;
  responses?: Record<string, Response>;
}

export interface Parameter {
  name: string;
  in: 'path' | 'query' | 'header' | 'body';
  description?: string;
  required?: boolean;
  type?: string;
  schema?: SchemaDefinition;
}

export interface RequestBody {
  description?: string;
  required?: boolean;
  content?: Record<string, { schema?: SchemaDefinition }>;
}

export interface Response {
  description?: string;
  schema?: SchemaDefinition;
  content?: Record<string, { schema?: SchemaDefinition }>;
}

export interface SchemaDefinition {
  type?: string;
  description?: string;
  properties?: Record<string, SchemaDefinition>;
  items?: SchemaDefinition;
  required?: string[];
  enum?: string[];
  $ref?: string;
  allOf?: SchemaDefinition[];
  oneOf?: SchemaDefinition[];
  anyOf?: SchemaDefinition[];
  // F5 vendor extensions (x-ves-*)
  'x-displayname'?: string;
  'x-ves-example'?: string;
  // Enrichment extensions (x-f5xc-*)
  'x-f5xc-category'?: string;
  'x-f5xc-requires-tier'?: string;
  'x-f5xc-complexity'?: string;
  'x-f5xc-example'?: string;
  'x-f5xc-description-short'?: string;
  'x-f5xc-description-medium'?: string;
  'x-f5xc-use-cases'?: string[];
  'x-f5xc-related-domains'?: string[];
  'x-f5xc-is-preview'?: boolean;
}

export interface SearchResult {
  name: string;
  type: string;
  path: string;
  snippet: string;
  score: number;
}

export enum ResponseFormat {
  MARKDOWN = 'markdown',
  JSON = 'json'
}

// =============================================================================
// Operation Metadata Types (v2.0.33 extensions)
// =============================================================================

/**
 * Danger level classification for operations
 */
export type DangerLevel = 'low' | 'medium' | 'high' | 'critical';

/**
 * Collection of operation metadata for all resources
 */
export interface OperationsMetadataCollection {
  generated_at: string;
  version: string;
  resources: Record<string, ResourceOperationInfo>;
}

/**
 * Operation metadata for a single resource
 */
export interface ResourceOperationInfo {
  resource: string;
  base_path?: string;
  operations: Record<string, OperationMetadata>;  // key: "create" | "read" | "update" | "delete" | "list"
  best_practices?: BestPracticesInfo;
  guided_workflows?: GuidedWorkflowInfo[];
}

/**
 * Operation-level metadata from x-f5xc-* extensions
 */
export interface OperationMetadata {
  method: string;  // HTTP method (POST, GET, PUT, DELETE)
  path: string;    // API path
  danger_level?: DangerLevel;
  discovered_response_time?: ResponseTimeInfo;
  required_fields?: string[];
  confirmation_required?: boolean;
  side_effects?: SideEffectsInfo;
  purpose?: string;
}

/**
 * Response time metrics from API discovery
 */
export interface ResponseTimeInfo {
  p50_ms: number;
  p95_ms: number;
  p99_ms: number;
  sample_count: number;
  source: 'measured' | 'estimate';
}

/**
 * Side effects of an operation
 */
export interface SideEffectsInfo {
  creates?: string[];
  modifies?: string[];
  deletes?: string[];
}

/**
 * Best practices information
 */
export interface BestPracticesInfo {
  common_errors?: CommonErrorInfo[];
}

/**
 * Common error and its resolution
 */
export interface CommonErrorInfo {
  code: number;
  message: string;
  resolution: string;
  prevention?: string;
}

/**
 * Guided workflow for resource creation
 */
export interface GuidedWorkflowInfo {
  name: string;
  description: string;
  steps: WorkflowStepInfo[];
}

/**
 * Step in a guided workflow
 */
export interface WorkflowStepInfo {
  order: number;
  action: string;
  description?: string;
  fields?: string[];
  validation?: string;
}
