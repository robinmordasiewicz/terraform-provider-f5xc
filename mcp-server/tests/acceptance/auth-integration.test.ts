// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

/**
 * End-to-End Auth Integration Tests
 *
 * Validates the @robinmordasiewicz/f5xc-auth package integration
 * in the F5XC Terraform MCP server.
 *
 * Test Matrix:
 * - Unauthenticated (Documentation Mode)
 * - Authenticated (Execution Mode) with Token
 * - Profile-based authentication
 * - Environment variable priority cascade
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

describe('Auth Integration Tests', () => {
  const originalEnv = process.env;

  beforeEach(() => {
    vi.clearAllMocks();
    process.env = { ...originalEnv };
  });

  afterEach(() => {
    process.env = originalEnv;
  });

  // ===========================================================================
  // SCENARIO 1: Documentation Mode (No Credentials)
  // ===========================================================================
  describe('Scenario: Documentation Mode (Unauthenticated)', () => {
    beforeEach(() => {
      setupDocumentationModeEnv();
    });

    it('should initialize CredentialManager in NONE mode', async () => {
      const credentialManager = new CredentialManager();
      await credentialManager.initialize();

      expect(credentialManager.getAuthMode()).toBe(AuthMode.NONE);
      expect(credentialManager.isAuthenticated()).toBe(false);
    });

    it('should return null for API URL when not authenticated', async () => {
      const credentialManager = new CredentialManager();
      await credentialManager.initialize();

      expect(credentialManager.getApiUrl()).toBeNull();
    });

    it('should return null for tenant when not authenticated', async () => {
      const credentialManager = new CredentialManager();
      await credentialManager.initialize();

      expect(credentialManager.getTenant()).toBeNull();
    });

    it('should return null for token when not authenticated', async () => {
      const credentialManager = new CredentialManager();
      await credentialManager.initialize();

      expect(credentialManager.getToken()).toBeNull();
    });

    it('should return null for namespace when not configured', async () => {
      const credentialManager = new CredentialManager();
      await credentialManager.initialize();

      expect(credentialManager.getNamespace()).toBeNull();
    });
  });

  // ===========================================================================
  // SCENARIO 2: Token Authentication
  // ===========================================================================
  describe('Scenario: Token Authentication', () => {
    beforeEach(() => {
      setupAuthenticatedModeEnv();
    });

    it('should initialize CredentialManager in TOKEN mode', async () => {
      const credentialManager = new CredentialManager();
      await credentialManager.initialize();

      expect(credentialManager.getAuthMode()).toBe(AuthMode.TOKEN);
      expect(credentialManager.isAuthenticated()).toBe(true);
    });

    it('should return configured API URL', async () => {
      const credentialManager = new CredentialManager();
      await credentialManager.initialize();

      // Note: f5xc-auth normalizes API URLs to include /api path
      expect(credentialManager.getApiUrl()).toBe('https://test.console.ves.volterra.io/api');
    });

    it('should extract tenant from API URL', async () => {
      const credentialManager = new CredentialManager();
      await credentialManager.initialize();

      expect(credentialManager.getTenant()).toBe('test');
    });

    it('should return configured token', async () => {
      const credentialManager = new CredentialManager();
      await credentialManager.initialize();

      expect(credentialManager.getToken()).toBe('test-token');
    });

    it('should handle custom namespace', async () => {
      process.env.F5XC_NAMESPACE = 'custom-namespace';

      const credentialManager = new CredentialManager();
      await credentialManager.initialize();

      expect(credentialManager.getNamespace()).toBe('custom-namespace');
    });
  });

  // ===========================================================================
  // SCENARIO 3: Environment Variable Priority
  // ===========================================================================
  describe('Scenario: Environment Variable Priority Cascade', () => {
    it('should prioritize env vars over profile settings', async () => {
      // Set up env vars with specific values
      setupAuthenticatedModeEnv({
        apiUrl: 'https://env-tenant.console.ves.volterra.io',
        apiToken: 'env-token',
      });

      const credentialManager = new CredentialManager();
      await credentialManager.initialize();

      // Env vars should take priority (note: f5xc-auth normalizes with /api)
      expect(credentialManager.getApiUrl()).toBe('https://env-tenant.console.ves.volterra.io/api');
      expect(credentialManager.getToken()).toBe('env-token');
      expect(credentialManager.getTenant()).toBe('env-tenant');
    });

    it('should fall back to documentation mode when no credentials', async () => {
      clearF5XCEnvVars();
      setupDocumentationModeEnv();

      const credentialManager = new CredentialManager();
      await credentialManager.initialize();

      expect(credentialManager.getAuthMode()).toBe(AuthMode.NONE);
    });
  });

  // ===========================================================================
  // SCENARIO 4: Credential Manager Reload
  // ===========================================================================
  describe('Scenario: Credential Manager Reload', () => {
    it('should reload credentials when environment changes', async () => {
      // Start in documentation mode
      setupDocumentationModeEnv();

      const credentialManager = new CredentialManager();
      await credentialManager.initialize();

      expect(credentialManager.getAuthMode()).toBe(AuthMode.NONE);

      // Change to authenticated mode
      setupAuthenticatedModeEnv();
      await credentialManager.reload();

      expect(credentialManager.getAuthMode()).toBe(AuthMode.TOKEN);
      expect(credentialManager.isAuthenticated()).toBe(true);
    });

    it('should transition from authenticated to documentation mode', async () => {
      // Start in authenticated mode
      setupAuthenticatedModeEnv();

      const credentialManager = new CredentialManager();
      await credentialManager.initialize();

      expect(credentialManager.getAuthMode()).toBe(AuthMode.TOKEN);

      // Clear credentials
      setupDocumentationModeEnv();
      await credentialManager.reload();

      expect(credentialManager.getAuthMode()).toBe(AuthMode.NONE);
      expect(credentialManager.isAuthenticated()).toBe(false);
    });
  });

  // ===========================================================================
  // SCENARIO 5: Multiple Credential Manager Instances
  // ===========================================================================
  describe('Scenario: Multiple Credential Manager Instances', () => {
    it('should allow multiple independent instances', async () => {
      setupAuthenticatedModeEnv({
        apiUrl: 'https://tenant-a.console.ves.volterra.io',
        apiToken: 'token-a',
      });

      const manager1 = new CredentialManager();
      await manager1.initialize();

      // Change environment
      setupAuthenticatedModeEnv({
        apiUrl: 'https://tenant-b.console.ves.volterra.io',
        apiToken: 'token-b',
      });

      const manager2 = new CredentialManager();
      await manager2.initialize();

      // First instance should retain original values
      expect(manager1.getTenant()).toBe('tenant-a');

      // Second instance should have new values
      expect(manager2.getTenant()).toBe('tenant-b');
    });
  });

  // ===========================================================================
  // SCENARIO 6: Profile Manager Integration
  // ===========================================================================
  describe('Scenario: Profile Manager Integration', () => {
    it('should access ProfileManager from shared package', () => {
      const profileManager = getProfileManager();

      expect(profileManager).toBeDefined();
      expect(typeof profileManager.list).toBe('function');
      expect(typeof profileManager.get).toBe('function');
      expect(typeof profileManager.getActive).toBe('function');
      expect(typeof profileManager.setActive).toBe('function');
    });

    it('should list profiles (may be empty)', async () => {
      setupDocumentationModeEnv();

      const profileManager = getProfileManager();
      const profiles = await profileManager.list();

      expect(Array.isArray(profiles)).toBe(true);
    });

    it('should get active profile (may be null)', async () => {
      setupDocumentationModeEnv();

      const profileManager = getProfileManager();
      const active = await profileManager.getActive();

      // In test environment, no active profile expected
      expect(active === null || typeof active === 'string').toBe(true);
    });
  });

  // ===========================================================================
  // SCENARIO 7: Error Handling
  // ===========================================================================
  describe('Scenario: Error Handling', () => {
    it('should handle invalid API URL gracefully', async () => {
      process.env.F5XC_API_URL = 'not-a-valid-url';
      process.env.F5XC_API_TOKEN = 'test-token';

      const credentialManager = new CredentialManager();
      await credentialManager.initialize();

      // Should still initialize but may have issues with URL parsing
      // The credential manager should not throw
      expect(credentialManager).toBeDefined();
    });

    it('should handle empty token gracefully', async () => {
      process.env.F5XC_API_URL = 'https://test.console.ves.volterra.io';
      process.env.F5XC_API_TOKEN = '';

      const credentialManager = new CredentialManager();
      await credentialManager.initialize();

      // Empty token should result in documentation mode
      expect(credentialManager.getAuthMode()).toBe(AuthMode.NONE);
    });

    it('should handle missing API URL with token', async () => {
      delete process.env.F5XC_API_URL;
      process.env.F5XC_API_TOKEN = 'test-token';

      const credentialManager = new CredentialManager();
      await credentialManager.initialize();

      // Without API URL, should be in documentation mode
      expect(credentialManager.getAuthMode()).toBe(AuthMode.NONE);
    });
  });

  // ===========================================================================
  // SCENARIO 8: AuthMode Enum Values
  // ===========================================================================
  describe('Scenario: AuthMode Enum Validation', () => {
    it('should have correct string values for AuthMode', () => {
      expect(AuthMode.NONE).toBe('none');
      expect(AuthMode.TOKEN).toBe('token');
      expect(AuthMode.CERTIFICATE).toBe('certificate');
    });

    it('should return string mode values (not numeric)', async () => {
      setupAuthenticatedModeEnv();

      const credentialManager = new CredentialManager();
      await credentialManager.initialize();

      const mode = credentialManager.getAuthMode();

      // Mode should be a string, not a number
      expect(typeof mode).toBe('string');
      expect(['none', 'token', 'certificate']).toContain(mode);
    });
  });
});

// ===========================================================================
// ACCEPTANCE TEST MATRIX
// ===========================================================================
describe('Auth Integration Test Matrix', () => {
  const originalEnv = process.env;

  beforeEach(() => {
    process.env = { ...originalEnv };
  });

  afterEach(() => {
    process.env = originalEnv;
  });

  const testMatrix: Array<{
    name: string;
    setup: () => void;
    expectedMode: AuthMode;
    expectedAuthenticated: boolean;
    expectedTenant: string | null;
  }> = [
    {
      name: 'No credentials → Documentation mode',
      setup: () => setupDocumentationModeEnv(),
      expectedMode: AuthMode.NONE,
      expectedAuthenticated: false,
      expectedTenant: null,
    },
    {
      name: 'Token only → Execution mode',
      setup: () => setupAuthenticatedModeEnv(),
      expectedMode: AuthMode.TOKEN,
      expectedAuthenticated: true,
      expectedTenant: 'test',
    },
    {
      name: 'Custom tenant URL → Extract tenant',
      setup: () => setupAuthenticatedModeEnv({
        apiUrl: 'https://custom-tenant.console.ves.volterra.io',
        apiToken: 'custom-token',
      }),
      expectedMode: AuthMode.TOKEN,
      expectedAuthenticated: true,
      expectedTenant: 'custom-tenant',
    },
    {
      name: 'Missing token → Documentation mode',
      setup: () => {
        clearF5XCEnvVars();
        process.env.F5XC_API_URL = 'https://test.console.ves.volterra.io';
        // No token set
        process.env.XDG_CONFIG_HOME = '/tmp/__nonexistent_test_config__';
      },
      expectedMode: AuthMode.NONE,
      expectedAuthenticated: false,
      expectedTenant: null,
    },
    {
      name: 'Empty token → Documentation mode',
      setup: () => {
        clearF5XCEnvVars();
        process.env.F5XC_API_URL = 'https://test.console.ves.volterra.io';
        process.env.F5XC_API_TOKEN = '';
        process.env.XDG_CONFIG_HOME = '/tmp/__nonexistent_test_config__';
      },
      expectedMode: AuthMode.NONE,
      expectedAuthenticated: false,
      expectedTenant: null,
    },
  ];

  testMatrix.forEach(({ name, setup, expectedMode, expectedAuthenticated, expectedTenant }) => {
    it(`Matrix: ${name}`, async () => {
      setup();

      const credentialManager = new CredentialManager();
      await credentialManager.initialize();

      expect(credentialManager.getAuthMode()).toBe(expectedMode);
      expect(credentialManager.isAuthenticated()).toBe(expectedAuthenticated);
      expect(credentialManager.getTenant()).toBe(expectedTenant);
    });
  });
});
