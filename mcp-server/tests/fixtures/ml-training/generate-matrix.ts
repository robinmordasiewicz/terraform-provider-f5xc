#!/usr/bin/env npx tsx
// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

/**
 * ML Training Matrix Generator
 *
 * Generates reproducible input→output mappings for ML training context
 * by calling the MCP server tools and capturing their outputs.
 */

import * as fs from 'fs';
import * as path from 'path';
import { handleMetadata } from '../../../src/tools/metadata.js';
import { ResponseFormat } from '../../../src/types.js';

// Resources to include in the training matrix
const CORE_RESOURCES = [
  // Tier 1: Core Load Balancing
  'namespace',
  'origin_pool',
  'http_loadbalancer',
  'tcp_loadbalancer',
  'udp_loadbalancer',
  'cdn_loadbalancer',
  'healthcheck',

  // Tier 2: Security
  'app_firewall',
  'service_policy',
  'service_policy_rule',
  'rate_limiter',
  'rate_limiter_policy',
  'waf_exclusion_policy',

  // Tier 3: Infrastructure
  'virtual_site',
  'certificate',
  'certificate_chain',
  'trusted_ca_list',

  // Tier 4: API Security
  'api_definition',
  'api_discovery',
  'api_testing',
  'api_crawler',
  'app_api_group',
];

interface ResourceData {
  resource_name: string;
  category: string;
  prompt_templates: string[];
  expected_terraform?: string;
  mcp_syntax?: object;
  mcp_oneof_groups?: Array<{
    group_name: string;
    fields: string[];
    required: boolean;
  }>;
  terraform_validate: string;
  terraform_fmt: string;
  determinism_verified: boolean;
}

interface TrainingMatrix {
  version: string;
  generated_at: string;
  resources: Record<string, ResourceData>;
}

function getResourceCategory(resource: string): string {
  const categories: Record<string, string[]> = {
    load_balancing: ['http_loadbalancer', 'tcp_loadbalancer', 'udp_loadbalancer', 'cdn_loadbalancer', 'origin_pool', 'healthcheck'],
    security: ['app_firewall', 'service_policy', 'service_policy_rule', 'rate_limiter', 'rate_limiter_policy', 'waf_exclusion_policy'],
    infrastructure: ['namespace', 'virtual_site'],
    certificates: ['certificate', 'certificate_chain', 'trusted_ca_list'],
    api_security: ['api_definition', 'api_discovery', 'api_testing', 'api_crawler', 'app_api_group'],
  };

  for (const [category, resources] of Object.entries(categories)) {
    if (resources.includes(resource)) {
      return category;
    }
  }
  return 'other';
}

function getPromptTemplates(resource: string): string[] {
  const templates: Record<string, string[]> = {
    namespace: [
      'Create an F5XC namespace called {name}',
      'Set up a new namespace for {project}',
    ],
    origin_pool: [
      'Create an origin pool for {domain} on port {port}',
      'Set up an origin pool with TLS to {backend}',
    ],
    http_loadbalancer: [
      'Create an HTTP load balancer for {domain}',
      'Set up an HTTP LB on public VIP for {app}',
      'Create an HTTPS load balancer with WAF for {domain}',
    ],
    healthcheck: [
      'Create an HTTP health check for path {path}',
      'Set up a TCP health check on port {port}',
    ],
    app_firewall: [
      'Create a WAF policy with default settings',
      'Set up an application firewall for {app}',
    ],
  };

  return templates[resource] || [`Create an F5XC ${resource.replace(/_/g, ' ')}`];
}

async function generateResourceData(resource: string): Promise<ResourceData> {
  console.log(`Processing: ${resource}`);

  // Get example from MCP - handleMetadata returns string directly
  let example = '';
  try {
    const exampleResult = await handleMetadata({
      operation: 'example',
      resource,
      pattern: 'basic',
      response_format: ResponseFormat.MARKDOWN,
    });
    // handleMetadata returns string directly, not an object with content array
    example = typeof exampleResult === 'string' ? exampleResult : '';
  } catch (e) {
    console.warn(`  Warning: No example for ${resource}`);
  }

  // Get syntax from MCP
  let syntax = {};
  try {
    const syntaxResult = await handleMetadata({
      operation: 'syntax',
      resource,
      response_format: ResponseFormat.JSON,
    });
    // handleMetadata returns string directly
    const syntaxText = typeof syntaxResult === 'string' ? syntaxResult : '{}';
    try {
      syntax = JSON.parse(syntaxText);
    } catch {
      // If not valid JSON, store as raw text
      syntax = { raw: syntaxText };
    }
  } catch (e) {
    console.warn(`  Warning: No syntax for ${resource}`);
  }

  // Get OneOf groups from MCP
  let oneofGroups: ResourceData['mcp_oneof_groups'] = [];
  try {
    const oneofResult = await handleMetadata({
      operation: 'oneof',
      resource,
      response_format: ResponseFormat.JSON,
    });
    // handleMetadata returns string directly
    const oneofText = typeof oneofResult === 'string' ? oneofResult : '{}';
    try {
      const oneofData = JSON.parse(oneofText);
      if (oneofData && typeof oneofData === 'object') {
        oneofGroups = Object.entries(oneofData).map(([name, data]: [string, any]) => ({
          group_name: name,
          fields: data.fields || [],
          required: data.required || false,
        }));
      }
    } catch {
      // If not valid JSON, skip
    }
  } catch (e) {
    console.warn(`  Warning: No OneOf groups for ${resource}`);
  }

  return {
    resource_name: resource,
    category: getResourceCategory(resource),
    prompt_templates: getPromptTemplates(resource),
    expected_terraform: example,
    mcp_syntax: syntax,
    mcp_oneof_groups: oneofGroups,
    terraform_validate: 'pending', // Will be filled by test runner
    terraform_fmt: 'pending', // Will be filled by test runner
    determinism_verified: false, // Will be verified by test runner
  };
}

async function main() {
  console.log('Generating ML Training Matrix...\n');

  const matrix: TrainingMatrix = {
    version: '3.20.0',
    generated_at: new Date().toISOString(),
    resources: {},
  };

  for (const resource of CORE_RESOURCES) {
    try {
      matrix.resources[resource] = await generateResourceData(resource);
      console.log(`  ✓ ${resource}`);
    } catch (error) {
      console.error(`  ✗ ${resource}: ${error}`);
    }
  }

  // Write outputs
  const outputDir = path.dirname(new URL(import.meta.url).pathname);

  // Write main matrix
  const matrixPath = path.join(outputDir, 'ml-training-data.json');
  fs.writeFileSync(matrixPath, JSON.stringify(matrix, null, 2));
  console.log(`\nWrote matrix to: ${matrixPath}`);

  // Write individual examples
  const examplesDir = path.join(outputDir, 'mcp-outputs/examples');
  for (const [resource, data] of Object.entries(matrix.resources)) {
    if (data.expected_terraform) {
      const examplePath = path.join(examplesDir, `${resource}.tf`);
      fs.writeFileSync(examplePath, data.expected_terraform);
    }
  }
  console.log(`Wrote ${Object.keys(matrix.resources).length} example files`);

  // Write syntax outputs
  const syntaxDir = path.join(outputDir, 'mcp-outputs/syntax');
  for (const [resource, data] of Object.entries(matrix.resources)) {
    if (data.mcp_syntax && Object.keys(data.mcp_syntax).length > 0) {
      const syntaxPath = path.join(syntaxDir, `${resource}.json`);
      fs.writeFileSync(syntaxPath, JSON.stringify(data.mcp_syntax, null, 2));
    }
  }

  // Write OneOf outputs
  const oneofDir = path.join(outputDir, 'mcp-outputs/oneof');
  for (const [resource, data] of Object.entries(matrix.resources)) {
    if (data.mcp_oneof_groups && data.mcp_oneof_groups.length > 0) {
      const oneofPath = path.join(oneofDir, `${resource}.json`);
      fs.writeFileSync(oneofPath, JSON.stringify(data.mcp_oneof_groups, null, 2));
    }
  }

  console.log('\n✓ ML Training Matrix generation complete');
}

main().catch(console.error);
