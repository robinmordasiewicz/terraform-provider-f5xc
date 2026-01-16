// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

/**
 * MCP Validation Operation Tests
 *
 * Tests that verify the MCP server's `validate` operation correctly identifies
 * common Terraform syntax errors, particularly the distinction between
 * empty block syntax `{}` and boolean assignment `= true`.
 *
 * These tests ensure the MCP server provides DETERMINISTIC guidance that
 * catches mistakes BEFORE Terraform validation.
 */

import { describe, it, expect, beforeAll } from 'vitest';
import { MCPClient, createMCPClient } from '../utils/mcp-client.js';
import { ResponseFormat } from '../../src/types.js';

describe('MCP Validation Operation', () => {
  let mcpClient: MCPClient;

  beforeAll(() => {
    mcpClient = createMCPClient(ResponseFormat.JSON);
  });

  // ===========================================================================
  // BOOLEAN ASSIGNMENT ERROR DETECTION
  // ===========================================================================

  describe('Boolean Assignment Error Detection', () => {
    it('should detect boolean assignment error for disable_waf', async () => {
      const invalidConfig = `
resource "f5xc_http_loadbalancer" "test" {
  name      = "test-lb"
  namespace = "default"
  domains   = ["test.example.com"]

  # WRONG: Should be empty block, not boolean
  disable_waf = true
}`;
      const result = await mcpClient.validateConfig('http_loadbalancer', invalidConfig);

      // Should detect this as invalid
      expect(result.valid).toBe(false);
      // Should have at least one error about boolean assignment
      expect(result.errors.length).toBeGreaterThan(0);
    });

    it('should detect boolean assignment error for round_robin', async () => {
      const invalidConfig = `
resource "f5xc_http_loadbalancer" "test" {
  name      = "test-lb"
  namespace = "default"
  domains   = ["test.example.com"]

  # WRONG: Should be empty block
  round_robin = true
}`;
      const result = await mcpClient.validateConfig('http_loadbalancer', invalidConfig);

      expect(result.valid).toBe(false);
      expect(result.errors.length).toBeGreaterThan(0);
    });

    it('should detect boolean assignment error for advertise_on_public_default_vip', async () => {
      const invalidConfig = `
resource "f5xc_http_loadbalancer" "test" {
  name      = "test-lb"
  namespace = "default"
  domains   = ["test.example.com"]

  # WRONG: Should be empty block
  advertise_on_public_default_vip = true
}`;
      const result = await mcpClient.validateConfig('http_loadbalancer', invalidConfig);

      expect(result.valid).toBe(false);
      expect(result.errors.length).toBeGreaterThan(0);
    });

    it('should detect boolean assignment error for no_tls in origin_pool', async () => {
      const invalidConfig = `
resource "f5xc_origin_pool" "test" {
  name      = "test-pool"
  namespace = "default"

  origin_servers {
    public_name {
      dns_name = "example.com"
    }
  }

  # WRONG: Should be empty block
  no_tls = true
}`;
      const result = await mcpClient.validateConfig('origin_pool', invalidConfig);

      expect(result.valid).toBe(false);
      expect(result.errors.length).toBeGreaterThan(0);
    });
  });

  // ===========================================================================
  // CORRECT EMPTY BLOCK SYNTAX VALIDATION
  // ===========================================================================

  describe('Correct Empty Block Syntax', () => {
    it('should pass for correct disable_waf empty block syntax', async () => {
      const validConfig = `
resource "f5xc_http_loadbalancer" "test" {
  name      = "test-lb"
  namespace = "default"
  domains   = ["test.example.com"]

  # CORRECT: Empty block syntax
  disable_waf {}
}`;
      const result = await mcpClient.validateConfig('http_loadbalancer', validConfig);

      // The validate operation may report other issues (missing required fields)
      // but should NOT report boolean_assignment errors for this field
      const booleanErrors = result.errors.filter(e =>
        e.type === 'boolean_assignment' && e.field === 'disable_waf'
      );
      expect(booleanErrors).toHaveLength(0);
    });

    it('should pass for correct round_robin empty block syntax', async () => {
      const validConfig = `
resource "f5xc_http_loadbalancer" "test" {
  name      = "test-lb"
  namespace = "default"
  domains   = ["test.example.com"]

  # CORRECT: Empty block syntax
  round_robin {}
}`;
      const result = await mcpClient.validateConfig('http_loadbalancer', validConfig);

      const booleanErrors = result.errors.filter(e =>
        e.type === 'boolean_assignment' && e.field === 'round_robin'
      );
      expect(booleanErrors).toHaveLength(0);
    });

    it('should pass for correct no_tls empty block syntax', async () => {
      const validConfig = `
resource "f5xc_origin_pool" "test" {
  name      = "test-pool"
  namespace = "default"

  origin_servers {
    public_name {
      dns_name = "example.com"
    }
  }

  # CORRECT: Empty block syntax
  no_tls {}
}`;
      const result = await mcpClient.validateConfig('origin_pool', validConfig);

      const booleanErrors = result.errors.filter(e =>
        e.type === 'boolean_assignment' && e.field === 'no_tls'
      );
      expect(booleanErrors).toHaveLength(0);
    });
  });

  // ===========================================================================
  // MCP EXAMPLE OUTPUT VALIDATION
  // ===========================================================================

  describe('MCP Example Output Contains Correct Syntax', () => {
    it('should generate http_loadbalancer example with empty blocks', async () => {
      const example = await mcpClient.getExample('http_loadbalancer', 'basic');

      // Example should contain empty block syntax
      expect(example).toContain('disable_waf {}');
      expect(example).toContain('round_robin {}');
      expect(example).toContain('advertise_on_public_default_vip {}');

      // Example should NOT contain boolean assignment
      expect(example).not.toMatch(/disable_waf\s*=\s*true/);
      expect(example).not.toMatch(/round_robin\s*=\s*true/);
      expect(example).not.toMatch(/advertise_on_public_default_vip\s*=\s*true/);
    });

    it('should generate origin_pool example with empty blocks', async () => {
      const example = await mcpClient.getExample('origin_pool', 'basic');

      // Example should contain empty block syntax for TLS choice
      expect(example).toContain('no_tls {}');

      // Example should NOT contain boolean assignment
      expect(example).not.toMatch(/no_tls\s*=\s*true/);
    });

    it('should generate namespace example', async () => {
      const example = await mcpClient.getExample('namespace', 'basic');

      // Namespace example should have basic structure
      expect(example).toContain('f5xc_namespace');
      expect(example).toContain('name');
    });
  });

  // ===========================================================================
  // SYNTAX GUIDANCE VALIDATION
  // ===========================================================================

  describe('MCP Syntax Guidance', () => {
    it('should provide syntax guidance for http_loadbalancer', async () => {
      const syntax = await mcpClient.getSyntax('http_loadbalancer');

      // Syntax guidance should exist and not be empty
      expect(syntax).toBeTruthy();
      expect(syntax.length).toBeGreaterThan(0);
    });

    it('should provide syntax guidance for origin_pool', async () => {
      const syntax = await mcpClient.getSyntax('origin_pool');

      expect(syntax).toBeTruthy();
      expect(syntax.length).toBeGreaterThan(0);
    });
  });

  // ===========================================================================
  // ONEOF GROUPS VALIDATION
  // ===========================================================================

  describe('MCP OneOf Groups', () => {
    it('should return OneOf groups for http_loadbalancer', async () => {
      const oneOfGroups = await mcpClient.getOneOfGroups('http_loadbalancer');

      // Should have OneOf groups
      expect(Object.keys(oneOfGroups).length).toBeGreaterThan(0);

      // Should have waf_choice group
      expect(oneOfGroups).toHaveProperty('waf_choice');
      expect(oneOfGroups.waf_choice.fields).toContain('disable_waf');

      // Should have hash_policy_choice group
      expect(oneOfGroups).toHaveProperty('hash_policy_choice');
      expect(oneOfGroups.hash_policy_choice.fields).toContain('round_robin');

      // Should have advertise_choice group
      expect(oneOfGroups).toHaveProperty('advertise_choice');
      expect(oneOfGroups.advertise_choice.fields).toContain('advertise_on_public_default_vip');
    });

    it('should return OneOf groups for origin_pool', async () => {
      const oneOfGroups = await mcpClient.getOneOfGroups('origin_pool');

      // Should have tls_choice group
      const hasTlsChoice = Object.keys(oneOfGroups).some(key =>
        key.includes('tls') || oneOfGroups[key].fields?.includes('no_tls')
      );
      expect(hasTlsChoice).toBe(true);
    });
  });

  // ===========================================================================
  // COMMON MISTAKES GUIDANCE
  // ===========================================================================

  describe('MCP Common Mistakes Guidance', () => {
    it('should provide common mistakes for http_loadbalancer', async () => {
      const mistakes = await mcpClient.getCommonMistakes('http_loadbalancer');

      // Should return guidance
      expect(mistakes).toBeTruthy();
      expect(mistakes.length).toBeGreaterThan(0);

      // Should mention boolean assignment mistake
      const mentionsBooleanMistake =
        mistakes.toLowerCase().includes('boolean') ||
        mistakes.toLowerCase().includes('= true') ||
        mistakes.includes('{}');
      expect(mentionsBooleanMistake).toBe(true);
    });
  });

  // ===========================================================================
  // EXTENDED RESOURCES - TIER 1 & TIER 2
  // ===========================================================================

  describe('MCP Validation Operation - Extended Resources', () => {
    const tierOneResources = ['tcp_loadbalancer', 'cdn_loadbalancer'];
    const tierTwoResources = ['service_policy', 'rate_limiter', 'bot_defense_policy', 'service_policy_rule'];
    const tierThreeResources = ['virtual_site'];

    describe.each(tierOneResources)('Tier 1: %s', (resource) => {
      it(`should provide syntax guidance for ${resource}`, async () => {
        const syntax = await mcpClient.getSyntax(resource);

        expect(syntax).toBeTruthy();
        expect(syntax.length).toBeGreaterThan(0);
      });

      it(`should return OneOf groups for ${resource}`, async () => {
        const oneOfGroups = await mcpClient.getOneOfGroups(resource);

        expect(typeof oneOfGroups).toBe('object');
      });

      it(`should provide common mistakes for ${resource}`, async () => {
        const mistakes = await mcpClient.getCommonMistakes(resource);

        expect(mistakes).toBeDefined();
      });
    });

    describe.each(tierTwoResources)('Tier 2: %s', (resource) => {
      it(`should provide syntax guidance for ${resource}`, async () => {
        const syntax = await mcpClient.getSyntax(resource);

        expect(syntax).toBeTruthy();
        expect(syntax.length).toBeGreaterThan(0);
      });

      it(`should return OneOf groups for ${resource}`, async () => {
        const oneOfGroups = await mcpClient.getOneOfGroups(resource);

        expect(typeof oneOfGroups).toBe('object');
      });

      it(`should provide common mistakes for ${resource}`, async () => {
        const mistakes = await mcpClient.getCommonMistakes(resource);

        expect(mistakes).toBeDefined();
      });
    });

    describe.each(tierThreeResources)('Tier 3: %s', (resource) => {
      it(`should provide syntax guidance for ${resource}`, async () => {
        const syntax = await mcpClient.getSyntax(resource);

        expect(syntax).toBeTruthy();
        expect(syntax.length).toBeGreaterThan(0);
      });

      it(`should return OneOf groups for ${resource}`, async () => {
        const oneOfGroups = await mcpClient.getOneOfGroups(resource);

        expect(typeof oneOfGroups).toBe('object');
      });

      it(`should provide common mistakes for ${resource}`, async () => {
        const mistakes = await mcpClient.getCommonMistakes(resource);

        expect(mistakes).toBeDefined();
      });
    });
  });

  // ===========================================================================
  // SECURITY CATEGORIES - TIER 1 (Sprint 2)
  // ===========================================================================

  describe('MCP Validation Operation - Security Categories (Tier 1)', () => {
    const tier1SecurityResources = [
      'certificate',
      'certificate_chain',
      'api_definition',
      'api_discovery',
      'api_testing',
      'rate_limiter',
      'rate_limiter_policy',
      'healthcheck'
    ];

    describe.each(tier1SecurityResources)('Tier 1 Security: %s', (resource) => {
      it(`should provide syntax guidance for ${resource}`, async () => {
        const syntax = await mcpClient.getSyntax(resource);

        expect(syntax).toBeTruthy();
        expect(syntax.length).toBeGreaterThan(0);
      });

      it(`should return OneOf groups for ${resource}`, async () => {
        const oneOfGroups = await mcpClient.getOneOfGroups(resource);

        expect(typeof oneOfGroups).toBe('object');
      });

      it(`should provide common mistakes for ${resource}`, async () => {
        const mistakes = await mcpClient.getCommonMistakes(resource);

        expect(mistakes).toBeDefined();
      });
    });
  });

  // ===========================================================================
  // SECURITY CATEGORIES - TIER 2 (Sprint 2)
  // ===========================================================================

  describe('MCP Validation Operation - Security Categories (Tier 2)', () => {
    const tier2FirewallResources = [
      'enhanced_firewall_policy',
      'network_firewall',
      'waf_exclusion_policy',
      'service_policy',
      'service_policy_rule',
      'origin_pool'
    ];

    describe.each(tier2FirewallResources)('Tier 2 Firewall/Policy: %s', (resource) => {
      it(`should provide syntax guidance for ${resource}`, async () => {
        const syntax = await mcpClient.getSyntax(resource);

        expect(syntax).toBeTruthy();
        expect(syntax.length).toBeGreaterThan(0);
      });

      it(`should return OneOf groups for ${resource}`, async () => {
        const oneOfGroups = await mcpClient.getOneOfGroups(resource);

        expect(typeof oneOfGroups).toBe('object');
      });

      it(`should provide common mistakes for ${resource}`, async () => {
        const mistakes = await mcpClient.getCommonMistakes(resource);

        expect(mistakes).toBeDefined();
      });
    });
  });

  // ===========================================================================
  // Tier 3: Supporting Resources (Sprint 3)
  // ===========================================================================

  const tier3SupportingResources = [
    'trusted_ca_list',
    'api_crawler',
    'app_api_group',
    'udp_loadbalancer'
  ];

  describe.each(tier3SupportingResources)('Tier 3 Supporting: %s', (resource) => {
    it(`should provide syntax guidance for ${resource}`, async () => {
      const syntax = await mcpClient.getSyntax(resource);

      expect(syntax).toBeTruthy();
      expect(syntax.length).toBeGreaterThan(0);
    });

    it(`should return OneOf groups for ${resource}`, async () => {
      const oneOfGroups = await mcpClient.getOneOfGroups(resource);

      expect(typeof oneOfGroups).toBe('object');
    });

    it(`should provide common mistakes for ${resource}`, async () => {
      const mistakes = await mcpClient.getCommonMistakes(resource);

      expect(mistakes).toBeDefined();
    });
  });
});
