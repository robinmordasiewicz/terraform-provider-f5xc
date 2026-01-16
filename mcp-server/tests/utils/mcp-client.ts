// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

/**
 * MCP Client Utility
 *
 * Provides a clean interface for invoking MCP server tools in tests.
 * Wraps the tool handlers for easy use in E2E testing.
 */

import { handleMetadata } from '../../src/tools/metadata.js';
import { handleDocs } from '../../src/tools/docs.js';
import { ResponseFormat } from '../../src/types.js';

// Type definitions matching the actual schema types
export type MetadataOperation =
  | 'oneof'
  | 'validation'
  | 'defaults'
  | 'enums'
  | 'attribute'
  | 'requires_replace'
  | 'tier'
  | 'dependencies'
  | 'troubleshoot'
  | 'summary'
  | 'syntax'
  | 'validate'
  | 'example'
  | 'mistakes';

export type DocsOperation = 'search' | 'get' | 'list';
export type DocType = 'resource' | 'data-source' | 'function' | 'guide';

export interface ValidationResult {
  valid: boolean;
  errors: Array<{
    type: string;
    field: string;
    message: string;
    found?: string;
    expected?: string;
  }>;
  warnings: Array<{
    type: string;
    field: string;
    message: string;
  }>;
}

/**
 * MCP Client for E2E testing
 *
 * Provides methods to invoke MCP tools and parse their responses.
 */
export class MCPClient {
  private responseFormat: ResponseFormat;

  constructor(format: ResponseFormat = ResponseFormat.JSON) {
    this.responseFormat = format;
  }

  // ===========================================================================
  // METADATA TOOL OPERATIONS
  // ===========================================================================

  /**
   * Get metadata for a specific operation and resource
   */
  async getMetadata(
    operation: MetadataOperation,
    options: {
      resource?: string;
      attribute?: string;
      config?: string;
      pattern?: string;
      error_code?: string;
      error_message?: string;
      tier?: 'STANDARD' | 'ADVANCED' | 'PREMIUM';
    } = {},
  ): Promise<string> {
    return handleMetadata({
      operation,
      resource: options.resource,
      attribute: options.attribute,
      config: options.config,
      pattern: options.pattern,
      error_code: options.error_code,
      error_message: options.error_message,
      tier: options.tier,
      response_format: this.responseFormat,
    });
  }

  /**
   * Get a working Terraform example for a resource
   */
  async getExample(resource: string, pattern?: string): Promise<string> {
    const result = await this.getMetadata('example', { resource, pattern });

    // If JSON format, extract the Terraform code
    if (this.responseFormat === ResponseFormat.JSON) {
      try {
        const data = JSON.parse(result);
        // The example is typically in a 'terraform' or 'code' field
        if (data.terraform) return data.terraform;
        if (data.code) return data.code;
        if (data.example) return data.example;
        // If there's a patterns array, get the first match
        if (data.patterns && Array.isArray(data.patterns)) {
          const matchingPattern = data.patterns.find(
            (p: { name: string; terraform?: string; code?: string }) =>
              !pattern || p.name === pattern,
          );
          if (matchingPattern) {
            return matchingPattern.terraform || matchingPattern.code || '';
          }
        }
      } catch {
        // If parsing fails, return raw result
      }
    }

    // For markdown, extract code blocks
    return this.extractTerraformCode(result);
  }

  /**
   * Get syntax guidance for a resource
   */
  async getSyntax(resource: string): Promise<string> {
    return this.getMetadata('syntax', { resource });
  }

  /**
   * Get OneOf groups for a resource (mutually exclusive fields)
   */
  async getOneOfGroups(resource: string): Promise<Record<string, { fields: string[]; description?: string }>> {
    const result = await this.getMetadata('oneof', { resource });

    if (this.responseFormat === ResponseFormat.JSON) {
      try {
        const data = JSON.parse(result);
        return data.groups || data.oneof_groups || {};
      } catch {
        return {};
      }
    }

    return {};
  }

  /**
   * Validate a Terraform configuration
   */
  async validateConfig(resource: string, config: string): Promise<ValidationResult> {
    const result = await this.getMetadata('validate', { resource, config });

    if (this.responseFormat === ResponseFormat.JSON) {
      try {
        const data = JSON.parse(result);
        return {
          valid: data.valid ?? data.errors?.length === 0 ?? true,
          errors: data.errors || [],
          warnings: data.warnings || [],
        };
      } catch {
        return { valid: true, errors: [], warnings: [] };
      }
    }

    // For markdown, check for error indicators
    const hasErrors = result.toLowerCase().includes('error') ||
                      result.toLowerCase().includes('invalid');
    return {
      valid: !hasErrors,
      errors: hasErrors ? [{ type: 'unknown', field: 'unknown', message: result }] : [],
      warnings: [],
    };
  }

  /**
   * Get common mistakes for a resource
   */
  async getCommonMistakes(resource: string): Promise<string> {
    return this.getMetadata('mistakes', { resource });
  }

  /**
   * Get validation patterns for a resource
   */
  async getValidationPatterns(resource: string, pattern?: string): Promise<string> {
    return this.getMetadata('validation', { resource, pattern });
  }

  /**
   * Get default values for resource attributes
   */
  async getDefaults(resource: string): Promise<string> {
    return this.getMetadata('defaults', { resource });
  }

  /**
   * Get enum values for resource attributes
   */
  async getEnums(resource: string): Promise<string> {
    return this.getMetadata('enums', { resource });
  }

  /**
   * Get resource dependencies and creation order
   */
  async getDependencies(resource: string): Promise<string> {
    return this.getMetadata('dependencies', { resource });
  }

  /**
   * Get troubleshooting guidance for an error
   */
  async troubleshoot(errorCode?: string, errorMessage?: string): Promise<string> {
    return this.getMetadata('troubleshoot', { error_code: errorCode, error_message: errorMessage });
  }

  // ===========================================================================
  // DOCUMENTATION TOOL OPERATIONS
  // ===========================================================================

  /**
   * Get documentation for a specific resource
   */
  async getDocs(
    operation: DocsOperation,
    options: {
      name?: string;
      query?: string;
      type?: DocType;
      limit?: number;
    } = {},
  ): Promise<string> {
    return handleDocs({
      operation,
      name: options.name,
      query: options.query,
      type: options.type,
      limit: options.limit ?? 20,
      response_format: this.responseFormat,
    });
  }

  /**
   * Get full documentation for a resource
   */
  async getResourceDoc(name: string): Promise<string> {
    return this.getDocs('get', { name, type: 'resource' });
  }

  /**
   * Search documentation
   */
  async searchDocs(query: string, type?: DocType, limit?: number): Promise<string> {
    return this.getDocs('search', { query, type, limit });
  }

  /**
   * List all documentation
   */
  async listDocs(type?: DocType): Promise<string> {
    return this.getDocs('list', { type });
  }

  // ===========================================================================
  // HELPER METHODS
  // ===========================================================================

  /**
   * Extract Terraform code from markdown response
   */
  private extractTerraformCode(markdown: string): string {
    // Match HCL/Terraform code blocks
    const codeBlockRegex = /```(?:hcl|terraform|tf)?\n([\s\S]*?)```/g;
    const matches: string[] = [];

    let match;
    while ((match = codeBlockRegex.exec(markdown)) !== null) {
      matches.push(match[1].trim());
    }

    if (matches.length > 0) {
      return matches.join('\n\n');
    }

    // If no code blocks found, return the raw markdown
    return markdown;
  }

  /**
   * Generate a complete Terraform configuration for testing
   * Combines provider config with resource config
   */
  generateCompleteConfig(
    apiUrl: string,
    resourceConfig: string,
  ): string {
    const providerBlock = `
terraform {
  required_providers {
    f5xc = {
      source  = "registry.terraform.io/robinmordasiewicz/f5xc"
      version = "~> 3.0"
    }
  }
}

provider "f5xc" {
  api_url = "${apiUrl}"
  # Token is read from F5XC_API_TOKEN environment variable
}
`;
    return providerBlock + '\n' + resourceConfig;
  }

  /**
   * Generate a unique resource name for testing
   */
  static generateTestResourceName(prefix: string = 'mcp-test'): string {
    const timestamp = Date.now();
    const random = Math.random().toString(36).substring(2, 8);
    return `${prefix}-${timestamp}-${random}`;
  }
}

/**
 * Factory function for creating MCP clients
 */
export function createMCPClient(format: ResponseFormat = ResponseFormat.JSON): MCPClient {
  return new MCPClient(format);
}
