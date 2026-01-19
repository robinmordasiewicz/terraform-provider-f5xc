// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

/**
 * ML Training Matrix Generator
 *
 * Generates a reproducible input→output mapping for ML training context.
 * Uses MCP server metadata to create known outcomes for all supported resources.
 */

import * as fs from 'fs';
import * as path from 'path';
import { fileURLToPath } from 'url';

// Import MCP metadata handler
const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

// Navigate to dist directory
const distPath = path.resolve(__dirname, '../../../dist');

interface ResourceEntry {
  prompt_templates: string[];
  expected_terraform: string;
  mcp_syntax: Record<string, unknown> | null;
  mcp_oneof_groups: string[];
  terraform_validate: 'success' | 'pending' | 'error';
  terraform_fmt: 'formatted' | 'pending' | 'error';
}

interface MLTrainingMatrix {
  version: string;
  generated_at: string;
  provider: string;
  resources: Record<string, ResourceEntry>;
}

// Resources to include in ML training matrix
const TRAINING_RESOURCES = [
  'namespace',
  'origin_pool',
  'http_loadbalancer',
  'healthcheck',
  'app_firewall',
  'tcp_loadbalancer',
  'service_policy',
  'route',
  'virtual_host',
  'certificate',
  'dns_zone',
  'dns_record',
  'alert_policy',
  'alert_receiver',
  'forward_proxy_policy',
  'network_policy',
  'rate_limiter_policy',
  'waf_exclusion_policy',
  'user_identification',
  'fast_acl',
  'fast_acl_rule',
  'service_policy_rule',
];

async function generateMatrix(): Promise<void> {
  console.log('Loading MCP metadata handler...');

  // Dynamic import of the metadata module
  const metadataModule = await import(path.join(distPath, 'tools/metadata.js'));
  const { handleMetadata } = metadataModule;

  const matrix: MLTrainingMatrix = {
    version: '3.20.0',
    generated_at: new Date().toISOString(),
    provider: 'robinmordasiewicz/f5xc',
    resources: {},
  };

  const examplesDir = path.join(__dirname, 'mcp-outputs/examples');

  console.log(`Processing ${TRAINING_RESOURCES.length} resources...`);

  for (const resource of TRAINING_RESOURCES) {
    console.log(`  - ${resource}`);

    try {
      // Get example from MCP
      const exampleResult = await handleMetadata({
        operation: 'example',
        resource,
        pattern: 'basic',
        response_format: 'json',
      });

      // handleMetadata returns a string directly
      const example = typeof exampleResult === 'string' ? exampleResult : '';

      // Get syntax info
      let syntax: Record<string, unknown> | null = null;
      try {
        const syntaxResult = await handleMetadata({
          operation: 'syntax',
          resource,
          response_format: 'json',
        });
        if (typeof syntaxResult === 'string') {
          try {
            syntax = JSON.parse(syntaxResult);
          } catch {
            // Not valid JSON, skip
          }
        }
      } catch {
        // Syntax not available for this resource
      }

      // Get oneof groups
      let oneofGroups: string[] = [];
      try {
        const oneofResult = await handleMetadata({
          operation: 'oneof',
          resource,
          response_format: 'json',
        });
        if (typeof oneofResult === 'string') {
          try {
            const parsed = JSON.parse(oneofResult);
            if (Array.isArray(parsed)) {
              oneofGroups = parsed.map((g: { name?: string }) => g.name || '').filter(Boolean);
            }
          } catch {
            // Not valid JSON, skip
          }
        }
      } catch {
        // OneOf not available for this resource
      }

      // Create prompt templates
      const promptTemplates = [
        `Create a Terraform configuration for ${resource.replace(/_/g, ' ')} called {name}`,
        `Using f5xc-terraform MCP, generate ${resource} resource with name {name}`,
      ];

      matrix.resources[resource] = {
        prompt_templates: promptTemplates,
        expected_terraform: example,
        mcp_syntax: syntax,
        mcp_oneof_groups: oneofGroups,
        terraform_validate: example ? 'pending' : 'error',
        terraform_fmt: example ? 'pending' : 'error',
      };

      // Write example to file
      if (example) {
        const examplePath = path.join(examplesDir, `${resource}.tf`);
        fs.writeFileSync(examplePath, example);
      }
    } catch (error) {
      console.error(`    Error processing ${resource}:`, error);
      matrix.resources[resource] = {
        prompt_templates: [],
        expected_terraform: '',
        mcp_syntax: null,
        mcp_oneof_groups: [],
        terraform_validate: 'error',
        terraform_fmt: 'error',
      };
    }
  }

  // Write matrix to file
  const matrixPath = path.join(__dirname, 'ml-training-matrix.json');
  fs.writeFileSync(matrixPath, JSON.stringify(matrix, null, 2));

  console.log(`\nGenerated ML training matrix:`);
  console.log(`  - Resources: ${Object.keys(matrix.resources).length}`);
  console.log(`  - Output: ${matrixPath}`);
  console.log(`  - Examples: ${examplesDir}`);
}

// Run generator
generateMatrix().catch(console.error);
