/**
 * Type definitions for F5 Distributed Cloud MCP Server
 */

export interface ResourceDoc {
  name: string;
  path: string;
  type: 'resource' | 'data-source' | 'function' | 'guide';
  content?: string;
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
