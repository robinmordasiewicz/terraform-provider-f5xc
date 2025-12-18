/**
 * Consolidated API Tool Handler
 *
 * Replaces: f5xc_terraform_search_api_specs, f5xc_terraform_get_api_spec,
 *           f5xc_terraform_find_endpoints, f5xc_terraform_get_schema_definition,
 *           f5xc_terraform_list_definitions
 *
 * Tool: f5xc_terraform_api
 */

import { ApiInput } from '../schemas/common.js';
import { ResponseFormat } from '../types.js';
import {
  searchApiSpecs,
  getApiSpec,
  findEndpoints,
  getSchemaDefinition,
  listDefinitions,
} from '../services/api-specs.js';

// =============================================================================
// HANDLER
// =============================================================================

/**
 * Handles the f5xc_terraform_api tool invocation
 * Routes to appropriate operation based on input.operation
 */
export async function handleApi(input: ApiInput): Promise<string> {
  const { operation, response_format } = input;

  switch (operation) {
    case 'search':
      return handleSearch(input, response_format);
    case 'get':
      return handleGet(input, response_format);
    case 'find_endpoints':
      return handleFindEndpoints(input, response_format);
    case 'get_definition':
      return handleGetDefinition(input, response_format);
    case 'list_definitions':
      return handleListDefinitions(input, response_format);
    default:
      throw new Error(`Unknown operation: ${operation}`);
  }
}

// =============================================================================
// OPERATION HANDLERS
// =============================================================================

function handleSearch(input: ApiInput, format: ResponseFormat): string {
  const { query, limit = 20 } = input;

  if (!query) {
    throw new Error('query is required for search operation');
  }

  const results = searchApiSpecs(query, limit);

  if (format === ResponseFormat.JSON) {
    return JSON.stringify(
      {
        operation: 'search',
        query,
        results_count: results.length,
        results: results.map((r) => ({
          name: r.name,
          type: r.type,
          score: r.score,
          snippet: r.snippet,
        })),
      },
      null,
      2,
    );
  }

  // Markdown format
  const lines: string[] = [
    `# API Spec Search: "${query}"`,
    '',
    `**Found**: ${results.length} specification(s)`,
    '',
  ];

  if (results.length === 0) {
    lines.push('No matching API specifications found.');
    lines.push('');
    lines.push('**Suggestions**:');
    lines.push("- Try different search terms (e.g., 'http_loadbalancer', 'namespace', 'waf')");
    lines.push("- Use `operation: 'get'` with a specific spec name");
    return lines.join('\n');
  }

  for (const result of results) {
    lines.push(`## ${result.name}`);
    lines.push(`- **Type**: ${result.type}`);
    lines.push(`- **Score**: ${result.score.toFixed(2)}`);
    if (result.snippet) {
      lines.push(`- **Info**: ${result.snippet.slice(0, 100)}...`);
    }
    lines.push('');
  }

  return lines.join('\n');
}

function handleGet(input: ApiInput, format: ResponseFormat): string {
  const name = input.name || input.spec_name;
  const { include_paths = true, include_definitions = false } = input;

  if (!name) {
    throw new Error('name or spec_name is required for get operation');
  }

  const spec = getApiSpec(name);

  if (!spec) {
    const error = {
      error: 'API specification not found',
      name,
      suggestion: "Use operation: 'search' to find available specifications",
    };
    return format === ResponseFormat.JSON
      ? JSON.stringify(error, null, 2)
      : `# Error: API Specification Not Found\n\n**Name**: ${name}\n\nUse \`operation: 'search'\` to find available specifications.`;
  }

  const content = spec.content;

  if (format === ResponseFormat.JSON) {
    const result: Record<string, unknown> = {
      operation: 'get',
      name: spec.name,
      schemaName: spec.schemaName,
      info: content?.info,
    };

    if (include_paths && content?.paths) {
      result.paths = Object.keys(content.paths).slice(0, 50);
      result.paths_count = Object.keys(content.paths).length;
    }

    if (include_definitions && content?.definitions) {
      result.definitions = Object.keys(content.definitions);
      result.definitions_count = Object.keys(content.definitions).length;
    }

    return JSON.stringify(result, null, 2);
  }

  // Markdown format
  const lines: string[] = [
    `# API Specification: ${spec.schemaName}`,
    '',
    `**File**: ${spec.name}`,
    '',
  ];

  if (content?.info) {
    lines.push('## Info');
    if (content.info.title) lines.push(`- **Title**: ${content.info.title}`);
    if (content.info.version) lines.push(`- **Version**: ${content.info.version}`);
    if (content.info.description) {
      lines.push(`- **Description**: ${content.info.description.slice(0, 200)}...`);
    }
    lines.push('');
  }

  if (include_paths && content?.paths) {
    const pathKeys = Object.keys(content.paths);
    lines.push(`## Endpoints (${pathKeys.length})`);
    lines.push('');
    for (const path of pathKeys.slice(0, 20)) {
      const methods = Object.keys(content.paths[path]).filter(
        (m) => ['get', 'post', 'put', 'delete', 'patch'].includes(m.toLowerCase()),
      );
      lines.push(`- \`${methods.join(', ').toUpperCase()}\` ${path}`);
    }
    if (pathKeys.length > 20) {
      lines.push(`- *... and ${pathKeys.length - 20} more endpoints*`);
    }
    lines.push('');
  }

  if (include_definitions && content?.definitions) {
    const defKeys = Object.keys(content.definitions);
    lines.push(`## Definitions (${defKeys.length})`);
    lines.push('');
    for (const def of defKeys.slice(0, 30)) {
      lines.push(`- ${def}`);
    }
    if (defKeys.length > 30) {
      lines.push(`- *... and ${defKeys.length - 30} more definitions*`);
    }
    lines.push('');
  }

  return lines.join('\n');
}

function handleFindEndpoints(input: ApiInput, format: ResponseFormat): string {
  const { pattern, method, limit = 20 } = input;

  if (!pattern) {
    throw new Error('pattern is required for find_endpoints operation');
  }

  const results = findEndpoints(pattern, method, limit);

  if (format === ResponseFormat.JSON) {
    return JSON.stringify(
      {
        operation: 'find_endpoints',
        pattern,
        method: method || 'any',
        results_count: results.length,
        results,
      },
      null,
      2,
    );
  }

  // Markdown format
  const lines: string[] = [
    `# Endpoint Search: "${pattern}"`,
    '',
    method ? `**Method Filter**: ${method}` : '',
    `**Found**: ${results.length} endpoint(s)`,
    '',
  ];

  if (results.length === 0) {
    lines.push('No matching endpoints found.');
    return lines.join('\n');
  }

  for (const result of results) {
    lines.push(`## \`${result.method}\` ${result.path}`);
    lines.push(`- **Spec**: ${result.specName}`);
    if (result.summary) {
      lines.push(`- **Summary**: ${result.summary}`);
    }
    lines.push('');
  }

  return lines.join('\n');
}

function handleGetDefinition(input: ApiInput, format: ResponseFormat): string {
  const { spec_name, definition_name } = input;

  if (!spec_name) {
    throw new Error('spec_name is required for get_definition operation');
  }
  if (!definition_name) {
    throw new Error('definition_name is required for get_definition operation');
  }

  const definition = getSchemaDefinition(spec_name, definition_name);

  if (!definition) {
    const error = {
      error: 'Schema definition not found',
      spec_name,
      definition_name,
      suggestion: "Use operation: 'list_definitions' to see available definitions",
    };
    return format === ResponseFormat.JSON
      ? JSON.stringify(error, null, 2)
      : `# Error: Schema Definition Not Found\n\n**Spec**: ${spec_name}\n**Definition**: ${definition_name}\n\nUse \`operation: 'list_definitions'\` to see available definitions.`;
  }

  if (format === ResponseFormat.JSON) {
    return JSON.stringify(
      {
        operation: 'get_definition',
        spec_name,
        definition_name,
        definition,
      },
      null,
      2,
    );
  }

  // Markdown format
  const lines: string[] = [
    `# Schema Definition: ${definition_name}`,
    '',
    `**Spec**: ${spec_name}`,
    '',
    '```json',
    JSON.stringify(definition, null, 2),
    '```',
  ];

  return lines.join('\n');
}

function handleListDefinitions(input: ApiInput, format: ResponseFormat): string {
  const { spec_name } = input;

  if (!spec_name) {
    throw new Error('spec_name is required for list_definitions operation');
  }

  const definitions = listDefinitions(spec_name);

  if (format === ResponseFormat.JSON) {
    return JSON.stringify(
      {
        operation: 'list_definitions',
        spec_name,
        count: definitions.length,
        definitions,
      },
      null,
      2,
    );
  }

  // Markdown format
  const lines: string[] = [
    `# Schema Definitions: ${spec_name}`,
    '',
    `**Total**: ${definitions.length} definition(s)`,
    '',
  ];

  if (definitions.length === 0) {
    lines.push('No definitions found in this specification.');
    return lines.join('\n');
  }

  for (const def of definitions.slice(0, 50)) {
    lines.push(`- ${def}`);
  }

  if (definitions.length > 50) {
    lines.push(`- *... and ${definitions.length - 50} more definitions*`);
  }

  return lines.join('\n');
}

// =============================================================================
// TOOL DEFINITION
// =============================================================================

export const API_TOOL_DEFINITION = {
  name: 'f5xc_terraform_api',
  description: 'Query 270+ F5XC OpenAPI specs - search, get, find endpoints, schema definitions',
};
