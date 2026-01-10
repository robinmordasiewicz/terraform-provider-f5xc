// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

/**
 * Auth Tool Acceptance Tests
 *
 * Tests the f5xc_terraform_auth tool behavior under different authentication states:
 * - Status operation in documentation vs authenticated mode
 * - List operation for profiles
 * - Validate operation with/without credentials
 */

import { describe, it, expect, beforeEach, afterEach, vi } from 'vitest';
import {
  CredentialManager,
  AuthMode,
  getProfileManager,
} from '@robinmordasiewicz/f5xc-auth';
import {
  clearF5XCEnvVars,
  setupDocumentationModeEnv,
  setupAuthenticatedModeEnv,
} from '../utils/ci-environment.js';
import {
  handleAuth,
  initializeAuth,
  getCredentialManager,
} from '../../src/tools/auth.js';
import { ResponseFormat } from '../../src/types.js';

describe('Auth Tool Acceptance Tests', () => {
  const originalEnv = process.env;

  beforeEach(() => {
    vi.clearAllMocks();
    process.env = { ...originalEnv };
  });

  afterEach(() => {
    process.env = originalEnv;
  });

  // ===========================================================================
  // STATUS OPERATION TESTS
  // ===========================================================================
  describe('Status Operation', () => {
    describe('Documentation Mode', () => {
      beforeEach(async () => {
        setupDocumentationModeEnv();
        await initializeAuth();
      });

      it('should return documentation mode status as JSON', async () => {
        const result = await handleAuth({
          operation: 'status',
          response_format: ResponseFormat.JSON,
        });

        const data = JSON.parse(result);
        expect(data.mode).toBe('none');
        expect(data.authenticated).toBe(false);
        expect(data.tenant).toBeNull();
        expect(data.api_url).toBeNull();
        expect(data.source).toBe('none (documentation mode)');
      });

      it('should return documentation mode status as markdown', async () => {
        const result = await handleAuth({
          operation: 'status',
          response_format: ResponseFormat.MARKDOWN,
        });

        expect(result).toContain('# F5XC Authentication Status');
        expect(result).toContain('**Mode** | none');
        expect(result).toContain('**Authenticated** | No');
        expect(result).toContain('Documentation Mode');
      });
    });

    describe('Authenticated Mode', () => {
      beforeEach(async () => {
        setupAuthenticatedModeEnv();
        await initializeAuth();
      });

      it('should return authenticated status as JSON', async () => {
        const result = await handleAuth({
          operation: 'status',
          response_format: ResponseFormat.JSON,
        });

        const data = JSON.parse(result);
        expect(data.mode).toBe('token');
        expect(data.authenticated).toBe(true);
        expect(data.tenant).toBe('test');
        expect(data.api_url).toBe('https://test.console.ves.volterra.io/api');
        expect(data.source).toBe('environment variables');
      });

      it('should return authenticated status as markdown', async () => {
        const result = await handleAuth({
          operation: 'status',
          response_format: ResponseFormat.MARKDOWN,
        });

        expect(result).toContain('# F5XC Authentication Status');
        expect(result).toContain('**Mode** | token');
        expect(result).toContain('**Authenticated** | Yes');
        expect(result).toContain('**Tenant** | test');
      });
    });

    describe('Custom Tenant', () => {
      beforeEach(async () => {
        setupAuthenticatedModeEnv({
          apiUrl: 'https://production.console.ves.volterra.io',
          apiToken: 'prod-token',
          namespace: 'prod-namespace',
        });
        await initializeAuth();
      });

      it('should extract tenant correctly', async () => {
        const result = await handleAuth({
          operation: 'status',
          response_format: ResponseFormat.JSON,
        });

        const data = JSON.parse(result);
        expect(data.tenant).toBe('production');
        expect(data.namespace).toBe('prod-namespace');
      });
    });
  });

  // ===========================================================================
  // LIST OPERATION TESTS
  // ===========================================================================
  describe('List Operation', () => {
    beforeEach(async () => {
      setupDocumentationModeEnv();
      await initializeAuth();
    });

    it('should return profile list as JSON', async () => {
      const result = await handleAuth({
        operation: 'list',
        response_format: ResponseFormat.JSON,
      });

      const data = JSON.parse(result);
      expect(data).toHaveProperty('profiles');
      expect(Array.isArray(data.profiles)).toBe(true);
      expect(data).toHaveProperty('active');
    });

    it('should return profile list as markdown', async () => {
      const result = await handleAuth({
        operation: 'list',
        response_format: ResponseFormat.MARKDOWN,
      });

      expect(result).toContain('# F5XC Profiles');
      // May contain "No profiles configured" or a table
      expect(
        result.includes('No profiles configured') || result.includes('| Profile |')
      ).toBe(true);
    });
  });

  // ===========================================================================
  // VALIDATE OPERATION TESTS
  // ===========================================================================
  describe('Validate Operation', () => {
    describe('Without Credentials', () => {
      beforeEach(async () => {
        setupDocumentationModeEnv();
        await initializeAuth();
      });

      it('should return validation failed as JSON', async () => {
        const result = await handleAuth({
          operation: 'validate',
          response_format: ResponseFormat.JSON,
        });

        const data = JSON.parse(result);
        expect(data.valid).toBe(false);
        expect(data.mode).toBe('none');
        expect(data.error).toContain('No credentials configured');
        expect(data.suggestion).toBeDefined();
      });

      it('should return validation failed as markdown', async () => {
        const result = await handleAuth({
          operation: 'validate',
          response_format: ResponseFormat.MARKDOWN,
        });

        expect(result).toContain('# Credential Validation');
        expect(result).toContain('**Result**: Not validated');
        expect(result).toContain('documentation mode');
      });
    });

    describe('With Credentials (Mock)', () => {
      beforeEach(async () => {
        setupAuthenticatedModeEnv();
        await initializeAuth();
      });

      it('should have initialized credential manager', () => {
        const credManager = getCredentialManager();
        expect(credManager).not.toBeNull();
        expect(credManager?.isAuthenticated()).toBe(true);
      });
    });
  });

  // ===========================================================================
  // SWITCH OPERATION TESTS
  // ===========================================================================
  describe('Switch Operation', () => {
    beforeEach(async () => {
      setupDocumentationModeEnv();
      await initializeAuth();
    });

    it('should require profile_name parameter', async () => {
      await expect(
        handleAuth({
          operation: 'switch',
          response_format: ResponseFormat.JSON,
        })
      ).rejects.toThrow('profile_name is required');
    });

    it('should throw for non-existent profile', async () => {
      await expect(
        handleAuth({
          operation: 'switch',
          profile_name: 'nonexistent-profile',
          response_format: ResponseFormat.JSON,
        })
      ).rejects.toThrow("Profile 'nonexistent-profile' not found");
    });
  });

  // ===========================================================================
  // TERRAFORM-ENV OPERATION TESTS
  // ===========================================================================
  describe('Terraform-Env Operation', () => {
    describe('Documentation Mode (No Credentials)', () => {
      beforeEach(async () => {
        setupDocumentationModeEnv();
        await initializeAuth();
      });

      it('should return error as JSON when not authenticated', async () => {
        const result = await handleAuth({
          operation: 'terraform-env',
          response_format: ResponseFormat.JSON,
        });

        const data = JSON.parse(result);
        expect(data.error).toBeDefined();
        expect(data.error).toContain('No credentials configured');
        expect(data.suggestion).toBeDefined();
      });

      it('should return error as markdown when not authenticated', async () => {
        const result = await handleAuth({
          operation: 'terraform-env',
          response_format: ResponseFormat.MARKDOWN,
        });

        expect(result).toContain('# Terraform Environment Variables');
        expect(result).toContain('**Error**: No credentials configured');
        expect(result).toContain('documentation mode');
      });
    });

    describe('Authenticated Mode', () => {
      beforeEach(async () => {
        setupAuthenticatedModeEnv();
        await initializeAuth();
      });

      it('should return shell exports as JSON', async () => {
        const result = await handleAuth({
          operation: 'terraform-env',
          output_type: 'shell',
          response_format: ResponseFormat.JSON,
        });

        const data = JSON.parse(result);
        expect(data.profile).toBeDefined();
        expect(data.auth_method).toBe('token');
        expect(data.variables).toBeDefined();
        expect(data.variables.F5XC_API_URL).toBe('https://test.console.ves.volterra.io/api');
        expect(data.variables.F5XC_API_TOKEN).toBe('test-token');
        expect(data.shell_command).toContain('export F5XC_API_URL=');
      });

      it('should return shell exports as markdown', async () => {
        const result = await handleAuth({
          operation: 'terraform-env',
          output_type: 'shell',
          response_format: ResponseFormat.MARKDOWN,
        });

        expect(result).toContain('# Terraform Environment Variables');
        expect(result).toContain('**Auth Method**: token');
        expect(result).toContain('```bash');
        expect(result).toContain('export F5XC_API_URL=');
        expect(result).toContain('export F5XC_API_TOKEN=');
        expect(result).toContain('## ⚠️ EXECUTE BEFORE TERRAFORM');
      });

      it('should return dotenv format as markdown', async () => {
        const result = await handleAuth({
          operation: 'terraform-env',
          output_type: 'dotenv',
          response_format: ResponseFormat.MARKDOWN,
        });

        expect(result).toContain('# Terraform Environment Variables (.env format)');
        expect(result).toContain('F5XC_API_URL=https://test.console.ves.volterra.io/api');
        expect(result).toContain('F5XC_API_TOKEN=test-token');
        expect(result).toContain('`.env` file');
      });

      it('should return JSON format as markdown', async () => {
        const result = await handleAuth({
          operation: 'terraform-env',
          output_type: 'json',
          response_format: ResponseFormat.MARKDOWN,
        });

        expect(result).toContain('# Terraform Environment Variables (JSON)');
        expect(result).toContain('```json');
        expect(result).toContain('"auth_method": "token"');
        expect(result).toContain('"F5XC_API_URL"');
      });

      it('should mask secrets when requested', async () => {
        const result = await handleAuth({
          operation: 'terraform-env',
          output_type: 'shell',
          mask_secrets: true,
          response_format: ResponseFormat.JSON,
        });

        const data = JSON.parse(result);
        expect(data.variables.F5XC_API_TOKEN).toContain('****');
        expect(data.variables.F5XC_API_TOKEN).not.toBe('test-token');
      });

      it('should not mask secrets by default', async () => {
        const result = await handleAuth({
          operation: 'terraform-env',
          output_type: 'shell',
          response_format: ResponseFormat.JSON,
        });

        const data = JSON.parse(result);
        expect(data.variables.F5XC_API_TOKEN).toBe('test-token');
        expect(data.variables.F5XC_API_TOKEN).not.toContain('****');
      });
    });

    describe('Custom Tenant', () => {
      beforeEach(async () => {
        setupAuthenticatedModeEnv({
          apiUrl: 'https://production.console.ves.volterra.io',
          apiToken: 'production-api-token-12345',
        });
        await initializeAuth();
      });

      it('should generate correct URL for custom tenant', async () => {
        const result = await handleAuth({
          operation: 'terraform-env',
          response_format: ResponseFormat.JSON,
        });

        const data = JSON.parse(result);
        expect(data.variables.F5XC_API_URL).toBe('https://production.console.ves.volterra.io/api');
        expect(data.variables.F5XC_API_TOKEN).toBe('production-api-token-12345');
      });

      it('should mask long tokens correctly', async () => {
        const result = await handleAuth({
          operation: 'terraform-env',
          mask_secrets: true,
          response_format: ResponseFormat.JSON,
        });

        const data = JSON.parse(result);
        expect(data.variables.F5XC_API_TOKEN).toMatch(/\*\*\*\*\.\.\.[\w-]{8}$/);
      });
    });

    describe('Output Type Defaults', () => {
      beforeEach(async () => {
        setupAuthenticatedModeEnv();
        await initializeAuth();
      });

      it('should default to shell output type', async () => {
        const result = await handleAuth({
          operation: 'terraform-env',
          response_format: ResponseFormat.MARKDOWN,
        });

        expect(result).toContain('```bash');
        expect(result).toContain('#!/bin/bash');
        expect(result).toContain('export F5XC_API_URL=');
      });
    });
  });

  // ===========================================================================
  // ERROR HANDLING TESTS
  // ===========================================================================
  describe('Error Handling', () => {
    beforeEach(async () => {
      setupDocumentationModeEnv();
      await initializeAuth();
    });

    it('should throw for unknown operation', async () => {
      await expect(
        handleAuth({
          operation: 'unknown' as any,
          response_format: ResponseFormat.JSON,
        })
      ).rejects.toThrow('Unknown operation');
    });
  });

  // ===========================================================================
  // CREDENTIAL MANAGER ACCESS TESTS
  // ===========================================================================
  describe('Credential Manager Access', () => {
    it('should expose credential manager after init', async () => {
      setupDocumentationModeEnv();
      await initializeAuth();

      const credManager = getCredentialManager();
      expect(credManager).not.toBeNull();
      expect(credManager).toBeInstanceOf(CredentialManager);
    });

    it('should return unauthenticated credential manager in doc mode', async () => {
      setupDocumentationModeEnv();
      await initializeAuth();

      const credManager = getCredentialManager();
      expect(credManager?.getAuthMode()).toBe(AuthMode.NONE);
      expect(credManager?.isAuthenticated()).toBe(false);
    });

    it('should return authenticated credential manager in auth mode', async () => {
      setupAuthenticatedModeEnv();
      await initializeAuth();

      const credManager = getCredentialManager();
      expect(credManager?.getAuthMode()).toBe(AuthMode.TOKEN);
      expect(credManager?.isAuthenticated()).toBe(true);
    });
  });
});

// ===========================================================================
// AUTH TOOL CAPABILITY MATRIX
// ===========================================================================
describe('Auth Tool Capability Matrix', () => {
  const originalEnv = process.env;

  beforeEach(() => {
    vi.clearAllMocks();
    process.env = { ...originalEnv };
  });

  afterEach(() => {
    process.env = originalEnv;
  });

  const capabilityMatrix: Array<{
    name: string;
    setup: () => void;
    expectedMode: string;
    expectedAuth: boolean;
    expectedSource: string;
  }> = [
    {
      name: 'Unauthenticated → Documentation mode',
      setup: () => setupDocumentationModeEnv(),
      expectedMode: 'none',
      expectedAuth: false,
      expectedSource: 'none (documentation mode)',
    },
    {
      name: 'Token authenticated → Execution mode',
      setup: () => setupAuthenticatedModeEnv(),
      expectedMode: 'token',
      expectedAuth: true,
      expectedSource: 'environment variables',
    },
    {
      name: 'Custom tenant → Full capabilities',
      setup: () => setupAuthenticatedModeEnv({
        apiUrl: 'https://production.console.ves.volterra.io',
        apiToken: 'prod-token',
      }),
      expectedMode: 'token',
      expectedAuth: true,
      expectedSource: 'environment variables',
    },
  ];

  capabilityMatrix.forEach(({ name, setup, expectedMode, expectedAuth, expectedSource }) => {
    it(`Capability Matrix: ${name}`, async () => {
      setup();
      await initializeAuth();

      const result = await handleAuth({
        operation: 'status',
        response_format: ResponseFormat.JSON,
      });

      const data = JSON.parse(result);
      expect(data.mode).toBe(expectedMode);
      expect(data.authenticated).toBe(expectedAuth);
      expect(data.source).toBe(expectedSource);
    });
  });
});

// ===========================================================================
// STATE TRANSITION TESTS
// ===========================================================================
describe('Auth State Transitions', () => {
  const originalEnv = process.env;

  beforeEach(() => {
    vi.clearAllMocks();
    process.env = { ...originalEnv };
  });

  afterEach(() => {
    process.env = originalEnv;
  });

  it('should transition from documentation to authenticated mode', async () => {
    // Start in documentation mode
    setupDocumentationModeEnv();
    await initializeAuth();

    let result = await handleAuth({
      operation: 'status',
      response_format: ResponseFormat.JSON,
    });
    let data = JSON.parse(result);
    expect(data.authenticated).toBe(false);

    // Change to authenticated mode
    setupAuthenticatedModeEnv();
    await initializeAuth();

    result = await handleAuth({
      operation: 'status',
      response_format: ResponseFormat.JSON,
    });
    data = JSON.parse(result);
    expect(data.authenticated).toBe(true);
  });

  it('should transition from authenticated to documentation mode', async () => {
    // Start in authenticated mode
    setupAuthenticatedModeEnv();
    await initializeAuth();

    let result = await handleAuth({
      operation: 'status',
      response_format: ResponseFormat.JSON,
    });
    let data = JSON.parse(result);
    expect(data.authenticated).toBe(true);

    // Change to documentation mode
    setupDocumentationModeEnv();
    await initializeAuth();

    result = await handleAuth({
      operation: 'status',
      response_format: ResponseFormat.JSON,
    });
    data = JSON.parse(result);
    expect(data.authenticated).toBe(false);
  });

  it('should handle multiple tenant switches', async () => {
    const tenants = ['tenant-a', 'tenant-b', 'tenant-c'];

    for (const tenant of tenants) {
      setupAuthenticatedModeEnv({
        apiUrl: `https://${tenant}.console.ves.volterra.io`,
        apiToken: `${tenant}-token`,
      });
      await initializeAuth();

      const result = await handleAuth({
        operation: 'status',
        response_format: ResponseFormat.JSON,
      });
      const data = JSON.parse(result);
      expect(data.tenant).toBe(tenant);
    }
  });
});
