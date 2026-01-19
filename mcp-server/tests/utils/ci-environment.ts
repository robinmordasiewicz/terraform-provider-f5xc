// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

/**
 * CI Environment Utilities
 *
 * Helpers for detecting CI environment and setting up test scenarios
 * for authenticated vs unauthenticated testing.
 */

/**
 * Detect if running in CI environment
 */
export function isCI(): boolean {
  return (
    process.env.CI === 'true' ||
    process.env.GITHUB_ACTIONS === 'true' ||
    process.env.GITLAB_CI === 'true' ||
    process.env.JENKINS_URL !== undefined ||
    process.env.CIRCLECI === 'true'
  );
}

/**
 * List of F5XC environment variables that affect authentication
 */
const F5XC_ENV_VARS = [
  'F5XC_API_URL',
  'F5XC_API_TOKEN',
  'F5XC_P12_FILE',
  'F5XC_P12_PASSWORD',
  'F5XC_CERT',
  'F5XC_KEY',
  'F5XC_NAMESPACE',
  'F5XC_TENANT',
  'F5XC_PROFILE',
];

/**
 * Clear all F5XC-related environment variables
 */
export function clearF5XCEnvVars(): void {
  for (const envVar of F5XC_ENV_VARS) {
    delete process.env[envVar];
  }
}

/**
 * Set up environment for documentation mode (no credentials)
 */
export function setupDocumentationModeEnv(): void {
  clearF5XCEnvVars();
  // Set XDG_CONFIG_HOME to a non-existent path to prevent profile loading
  process.env.XDG_CONFIG_HOME = '/tmp/__nonexistent_test_config__';
}

/**
 * Set up environment for authenticated mode with token
 */
export function setupAuthenticatedModeEnv(options?: {
  apiUrl?: string;
  apiToken?: string;
  namespace?: string;
}): void {
  clearF5XCEnvVars();
  process.env.F5XC_API_URL = options?.apiUrl || 'https://test.console.ves.volterra.io';
  process.env.F5XC_API_TOKEN = options?.apiToken || 'test-token';
  if (options?.namespace) {
    process.env.F5XC_NAMESPACE = options.namespace;
  }
  // Set XDG_CONFIG_HOME to a non-existent path to prevent profile interference
  process.env.XDG_CONFIG_HOME = '/tmp/__nonexistent_test_config__';
}

/**
 * Check if real API testing is possible (has valid credentials)
 * Supports both F5XC_API_* and F5XC_TEST_API_* naming conventions
 */
export function hasRealCredentials(): boolean {
  return Boolean(
    (process.env.F5XC_API_URL || process.env.F5XC_TEST_API_URL) &&
    (process.env.F5XC_API_TOKEN || process.env.F5XC_TEST_API_TOKEN)
  );
}

/**
 * Get real test credentials if available
 * Supports both F5XC_API_* and F5XC_TEST_API_* naming conventions
 */
export function getRealCredentials(): { apiUrl: string; apiToken: string } | null {
  if (!hasRealCredentials()) {
    return null;
  }
  return {
    apiUrl: process.env.F5XC_API_URL || process.env.F5XC_TEST_API_URL!,
    apiToken: process.env.F5XC_API_TOKEN || process.env.F5XC_TEST_API_TOKEN!,
  };
}

// =============================================================================
// E2E TEST HELPERS
// =============================================================================

/**
 * Check if E2E tests should run
 * Requires credentials and optionally Terraform CLI
 */
export function shouldRunE2ETests(): boolean {
  return hasRealCredentials();
}

/**
 * Get test namespace prefix for E2E tests
 * Default: 'mcp-e2e-test'
 */
export function getTestNamespacePrefix(): string {
  return process.env.F5XC_TEST_NAMESPACE_PREFIX || 'mcp-e2e-test';
}

/**
 * Check if test cleanup should be performed
 * Default: true
 */
export function shouldCleanup(): boolean {
  const cleanup = process.env.F5XC_TEST_CLEANUP;
  if (cleanup === undefined) return true;
  return cleanup.toLowerCase() !== 'false';
}

/**
 * Generate a unique test resource name
 */
export function generateTestResourceName(prefix?: string): string {
  const basePrefix = prefix || getTestNamespacePrefix();
  const timestamp = Date.now();
  const random = Math.random().toString(36).substring(2, 8);
  return `${basePrefix}-${timestamp}-${random}`;
}

/**
 * Setup environment for E2E testing
 * Uses real credentials from F5XC_TEST_* variables
 */
export function setupE2ETestEnv(): {
  apiUrl: string;
  apiToken: string;
  namespace: string;
} | null {
  const credentials = getRealCredentials();
  if (!credentials) {
    return null;
  }

  // Set up environment for Terraform
  process.env.F5XC_API_URL = credentials.apiUrl;
  process.env.F5XC_API_TOKEN = credentials.apiToken;

  // Generate unique namespace for this test run
  const namespace = generateTestResourceName();

  return {
    ...credentials,
    namespace,
  };
}

/**
 * Skip message for E2E tests when credentials are not available
 */
export const E2E_SKIP_MESSAGE =
  'E2E tests require F5XC_TEST_API_URL and F5XC_TEST_API_TOKEN environment variables';

/**
 * Skip message for Terraform tests when CLI is not available
 */
export const TERRAFORM_SKIP_MESSAGE =
  'Terraform CLI is not available - skipping Terraform-based tests';
