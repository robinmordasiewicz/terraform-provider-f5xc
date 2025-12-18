/**
 * Consolidated Documentation Tool Handler
 *
 * Replaces: f5xc_terraform_search_docs, f5xc_terraform_get_doc, f5xc_terraform_list_docs
 *
 * Tool: f5xc_terraform_docs
 */

import { DocsInput } from '../schemas/common.js';
import { ResponseFormat } from '../types.js';
import {
  searchDocumentation,
  getDocumentation,
  listDocumentation,
} from '../services/documentation.js';

// =============================================================================
// HANDLER
// =============================================================================

/**
 * Handles the f5xc_terraform_docs tool invocation
 * Routes to appropriate operation based on input.operation
 */
export async function handleDocs(input: DocsInput): Promise<string> {
  const { operation, response_format } = input;

  switch (operation) {
    case 'search':
      return handleSearch(input, response_format);
    case 'get':
      return handleGet(input, response_format);
    case 'list':
      return handleList(input, response_format);
    default:
      throw new Error(`Unknown operation: ${operation}`);
  }
}

// =============================================================================
// OPERATION HANDLERS
// =============================================================================

function handleSearch(input: DocsInput, format: ResponseFormat): string {
  const { query, type, limit = 20 } = input;

  if (!query) {
    throw new Error('query is required for search operation');
  }

  const results = searchDocumentation(query, type, limit);

  if (format === ResponseFormat.JSON) {
    return JSON.stringify(
      {
        operation: 'search',
        query,
        type: type || 'all',
        results_count: results.length,
        results,
      },
      null,
      2,
    );
  }

  // Markdown format
  const lines: string[] = [
    `# Search Results: "${query}"`,
    '',
    `**Found**: ${results.length} result(s)`,
    type ? `**Type Filter**: ${type}` : '',
    '',
  ];

  if (results.length === 0) {
    lines.push('No matching documentation found.');
    lines.push('');
    lines.push('**Suggestions**:');
    lines.push('- Try different search terms');
    lines.push("- Use `operation: 'list'` to see all available documentation");
    return lines.join('\n');
  }

  for (const result of results) {
    lines.push(`## ${result.name}`);
    lines.push(`- **Type**: ${result.type}`);
    lines.push(`- **Score**: ${result.score.toFixed(2)}`);
    if (result.snippet) {
      lines.push(`- **Snippet**: ${result.snippet.slice(0, 100)}...`);
    }
    lines.push('');
  }

  return lines.join('\n');
}

function handleGet(input: DocsInput, format: ResponseFormat): string {
  const { name, type } = input;

  if (!name) {
    throw new Error('name is required for get operation');
  }

  const doc = getDocumentation(name, type);

  if (!doc) {
    const error = {
      error: 'Documentation not found',
      resource: name,
      type: type || 'any',
      suggestion: "Use operation: 'search' or 'list' to find available documentation",
    };
    return format === ResponseFormat.JSON
      ? JSON.stringify(error, null, 2)
      : `# Error: Documentation Not Found\n\n**Resource**: ${name}\n**Type**: ${type || 'any'}\n\nUse \`operation: 'list'\` to see available documentation.`;
  }

  if (format === ResponseFormat.JSON) {
    return JSON.stringify(
      {
        operation: 'get',
        name: doc.name,
        type: doc.type,
        path: doc.path,
        content: doc.content,
      },
      null,
      2,
    );
  }

  // Return the full markdown content
  return doc.content || `# ${doc.name}\n\nNo content available.`;
}

function handleList(input: DocsInput, format: ResponseFormat): string {
  const { type } = input;

  const docs = listDocumentation(type);

  if (format === ResponseFormat.JSON) {
    return JSON.stringify(
      {
        operation: 'list',
        type: type || 'all',
        total: docs.length,
        items: docs.map((d) => ({
          name: d.name,
          type: d.type,
          subscriptionTier: d.subscriptionTier,
        })),
      },
      null,
      2,
    );
  }

  // Markdown format
  const lines: string[] = [
    `# F5XC Terraform Documentation`,
    '',
    `**Total**: ${docs.length} items`,
    type ? `**Type Filter**: ${type}` : '',
    '',
  ];

  // Group by type
  const byType = new Map<string, typeof docs>();
  for (const doc of docs) {
    const t = doc.type;
    if (!byType.has(t)) {
      byType.set(t, []);
    }
    byType.get(t)!.push(doc);
  }

  for (const [docType, typeDocs] of byType) {
    lines.push(`## ${docType}s (${typeDocs.length})`);
    lines.push('');
    for (const doc of typeDocs.slice(0, 20)) {
      // Limit to 20 per type for readability
      const tierInfo = doc.subscriptionTier ? ` [${doc.subscriptionTier}]` : '';
      lines.push(`- **${doc.name}**${tierInfo}`);
    }
    if (typeDocs.length > 20) {
      lines.push(`- *... and ${typeDocs.length - 20} more*`);
    }
    lines.push('');
  }

  return lines.join('\n');
}

// =============================================================================
// TOOL DEFINITION
// =============================================================================

export const DOCS_TOOL_DEFINITION = {
  name: 'f5xc_terraform_docs',
  description: 'F5XC Terraform docs - search, get, or list resources/data-sources/functions/guides',
};
