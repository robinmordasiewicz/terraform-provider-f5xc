// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

/**
 * API Specifications Service
 * Loads and manages F5 Distributed Cloud OpenAPI specifications
 */

import { readFileSync, existsSync, readdirSync } from 'fs';
import { join, dirname, basename as _basename } from 'path';
import { fileURLToPath } from 'url';
import type { ApiSpec, OpenAPISpec, SearchResult, SchemaDefinition } from '../types.js';

// Get the package root directory
const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

// When installed via npm, specs are in dist/docs/specifications/api/ relative to package root
// When running in development, specs are in parent project's docs/specifications/api/
const PACKAGE_ROOT = join(__dirname, '..', '..'); // mcp-server/
const BUNDLED_SPECS = join(PACKAGE_ROOT, 'dist', 'docs', 'specifications', 'api');
const PROJECT_ROOT = join(PACKAGE_ROOT, '..'); // terraform-provider-f5xc/
const PROJECT_SPECS = join(PROJECT_ROOT, 'docs', 'specifications', 'api');

// Use bundled specs if available (npm install), otherwise fall back to project specs (development)
function getApiSpecsPath(): string {
  if (existsSync(BUNDLED_SPECS)) {
    return BUNDLED_SPECS;
  }
  return PROJECT_SPECS;
}

// API specifications path
const API_SPECS_PATH = getApiSpecsPath();

// Cache for loaded specifications
const specsCache = new Map<string, ApiSpec>();
let allSpecs: ApiSpec[] | null = null;

/**
 * Parse schema name from filename
 * Example: docs-cloud-f5-com.0019.public.ves.io.schema.app_firewall.ves-swagger.json
 * Returns: app_firewall
 */
function parseSchemaName(filename: string): string {
  // Remove .ves-swagger.json suffix
  const withoutSuffix = filename.replace('.ves-swagger.json', '');
  // Split by dots and get the meaningful parts
  const parts = withoutSuffix.split('.');
  // Find the schema name (usually after 'schema' or the last meaningful part)
  const schemaIndex = parts.findIndex(p => p === 'schema');
  if (schemaIndex !== -1 && schemaIndex < parts.length - 1) {
    // Get everything after 'schema'
    return parts.slice(schemaIndex + 1).join('_');
  }
  // Fallback: use last part
  return parts[parts.length - 1];
}

/**
 * Load all API specification metadata
 */
export function loadAllApiSpecs(): ApiSpec[] {
  if (allSpecs) {
    return allSpecs;
  }

  allSpecs = [];

  if (!existsSync(API_SPECS_PATH)) {
    console.error(`API specs path not found: ${API_SPECS_PATH}`);
    return allSpecs;
  }

  const files = readdirSync(API_SPECS_PATH).filter(f => f.endsWith('.json'));

  for (const file of files) {
    const schemaName = parseSchemaName(file);
    allSpecs.push({
      name: file,
      path: join(API_SPECS_PATH, file),
      schemaName,
    });
  }

  return allSpecs;
}

/**
 * Get a specific API specification by schema name or filename
 */
export function getApiSpec(identifier: string): ApiSpec | null {
  if (specsCache.has(identifier)) {
    return specsCache.get(identifier)!;
  }

  const specs = loadAllApiSpecs();
  const identifierLower = identifier.toLowerCase();

  // Try to find by schema name first
  let spec = specs.find(s =>
    s.schemaName.toLowerCase() === identifierLower ||
    s.schemaName.toLowerCase().replace(/_/g, '-') === identifierLower ||
    s.schemaName.toLowerCase().includes(identifierLower)
  );

  // Then try by filename
  if (!spec) {
    spec = specs.find(s => s.name.toLowerCase().includes(identifierLower));
  }

  if (spec && existsSync(spec.path)) {
    try {
      const content = JSON.parse(readFileSync(spec.path, 'utf-8')) as OpenAPISpec;
      spec = { ...spec, content };
      specsCache.set(identifier, spec);
      return spec;
    } catch (error) {
      console.error(`Error parsing API spec ${spec.path}:`, error);
      return null;
    }
  }

  return null;
}

/**
 * List all available API specifications
 */
export function listApiSpecs(): ApiSpec[] {
  return loadAllApiSpecs();
}

/**
 * Search API specifications
 */
export function searchApiSpecs(query: string, limit: number = 20): SearchResult[] {
  const specs = loadAllApiSpecs();
  const results: SearchResult[] = [];
  const queryLower = query.toLowerCase();
  const queryTerms = queryLower.split(/\s+/).filter(t => t.length > 0);

  for (const spec of specs) {
    let score = 0;
    const schemaNameLower = spec.schemaName.toLowerCase();
    const filenameLower = spec.name.toLowerCase();

    // Exact schema name match
    if (schemaNameLower === queryLower) {
      score += 100;
    }
    // Schema name contains query
    else if (schemaNameLower.includes(queryLower)) {
      score += 50;
    }
    // Filename contains query
    else if (filenameLower.includes(queryLower)) {
      score += 30;
    }
    // Any term matches
    else {
      for (const term of queryTerms) {
        if (schemaNameLower.includes(term)) {
          score += 20;
        }
        if (filenameLower.includes(term)) {
          score += 10;
        }
      }
    }

    if (score > 0) {
      // Try to load spec to get additional info
      let snippet = `Schema: ${spec.schemaName}`;
      try {
        if (existsSync(spec.path)) {
          const content = JSON.parse(readFileSync(spec.path, 'utf-8')) as OpenAPISpec;
          if (content.info?.title) {
            snippet = content.info.title;
          }
          if (content.info?.description) {
            snippet += ` - ${content.info.description.slice(0, 100)}...`;
          }
        }
      } catch {
        // Ignore parse errors for search
      }

      results.push({
        name: spec.schemaName,
        type: 'api-spec',
        path: spec.path,
        snippet,
        score,
      });
    }
  }

  // Sort by score descending
  results.sort((a, b) => b.score - a.score);

  return results.slice(0, limit);
}

/**
 * Find API endpoints matching a pattern
 */
export function findEndpoints(
  pattern: string,
  method?: string,
  limit: number = 20
): Array<{
  specName: string;
  path: string;
  method: string;
  summary?: string;
  operationId?: string;
}> {
  const specs = loadAllApiSpecs();
  const results: Array<{
    specName: string;
    path: string;
    method: string;
    summary?: string;
    operationId?: string;
  }> = [];
  const patternLower = pattern.toLowerCase();

  for (const spec of specs) {
    if (!existsSync(spec.path)) continue;

    try {
      const content = JSON.parse(readFileSync(spec.path, 'utf-8')) as OpenAPISpec;
      const paths = content.paths || {};

      for (const [pathStr, pathItem] of Object.entries(paths)) {
        if (!pathStr.toLowerCase().includes(patternLower)) continue;

        const methods = ['get', 'post', 'put', 'delete', 'patch'] as const;
        for (const m of methods) {
          if (method && m.toLowerCase() !== method.toLowerCase()) continue;

          const operation = pathItem[m];
          if (operation) {
            results.push({
              specName: spec.schemaName,
              path: pathStr,
              method: m.toUpperCase(),
              summary: operation.summary,
              operationId: operation.operationId,
            });
          }
        }

        if (results.length >= limit) break;
      }
    } catch {
      // Skip specs that fail to parse
    }

    if (results.length >= limit) break;
  }

  return results;
}

/**
 * Get schema definition from a spec
 */
export function getSchemaDefinition(
  specIdentifier: string,
  definitionName: string
): SchemaDefinition | null {
  const spec = getApiSpec(specIdentifier);
  if (!spec?.content) return null;

  // Check Swagger 2.0 definitions
  if (spec.content.definitions?.[definitionName]) {
    return spec.content.definitions[definitionName];
  }

  // Check OpenAPI 3.0 component schemas
  if (spec.content.components?.schemas?.[definitionName]) {
    return spec.content.components.schemas[definitionName];
  }

  // Try partial match
  const definitions = spec.content.definitions || spec.content.components?.schemas || {};
  for (const [name, schema] of Object.entries(definitions)) {
    if (name.toLowerCase().includes(definitionName.toLowerCase())) {
      return schema;
    }
  }

  return null;
}

/**
 * List all definitions in a spec
 */
export function listDefinitions(specIdentifier: string): string[] {
  const spec = getApiSpec(specIdentifier);
  if (!spec?.content) return [];

  const definitions = spec.content.definitions || spec.content.components?.schemas || {};
  return Object.keys(definitions);
}

/**
 * Get API specifications summary
 */
export function getApiSpecsSummary(): {
  total: number;
  categories: Record<string, number>;
} {
  const specs = loadAllApiSpecs();
  const categories: Record<string, number> = {};

  for (const spec of specs) {
    // Extract category from schema name (e.g., 'views', 'api_sec', etc.)
    const parts = spec.name.split('.');
    const schemaIndex = parts.findIndex(p => p === 'schema');
    if (schemaIndex !== -1 && schemaIndex < parts.length - 1) {
      const category = parts[schemaIndex + 1];
      categories[category] = (categories[category] || 0) + 1;
    }
  }

  return {
    total: specs.length,
    categories,
  };
}
