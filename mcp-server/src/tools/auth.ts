/**
 * Authentication Tool Handler
 *
 * Provides authentication status, profile management, and credential validation
 * for the F5XC Terraform MCP server.
 *
 * Uses the shared @robinmordasiewicz/f5xc-auth package for unified authentication
 * across F5XC MCP servers.
 *
 * Tool: f5xc_terraform_auth
 */

import axios from 'axios';
import { AuthInput } from '../schemas/common.js';
import { ResponseFormat } from '../types.js';
import {
  CredentialManager,
  AuthMode,
  getProfileManager,
  type Profile,
} from '@robinmordasiewicz/f5xc-auth';

// =============================================================================
// TOOL DEFINITION
// =============================================================================

export const AUTH_TOOL_DEFINITION = {
  name: 'f5xc_terraform_auth',
  description: 'Check authentication status, list/switch profiles, validate credentials',
};

// =============================================================================
// MODULE-LEVEL STATE
// =============================================================================

let credentialManager: CredentialManager | null = null;

/**
 * Initialize the credential manager (called from server startup)
 */
export async function initializeAuth(): Promise<CredentialManager> {
  credentialManager = new CredentialManager();
  await credentialManager.initialize();
  return credentialManager;
}

/**
 * Get the current credential manager instance
 */
export function getCredentialManager(): CredentialManager | null {
  return credentialManager;
}

// =============================================================================
// HANDLER
// =============================================================================

/**
 * Handles the f5xc_terraform_auth tool invocation
 */
export async function handleAuth(input: AuthInput): Promise<string> {
  const { operation, profile_name, output_type, mask_secrets, response_format } = input;

  switch (operation) {
    case 'status':
      return handleStatus(response_format);
    case 'list':
      return handleList(response_format);
    case 'switch':
      return handleSwitch(profile_name, response_format);
    case 'validate':
      return handleValidate(response_format);
    case 'terraform-env':
      return handleTerraformEnv(output_type || 'shell', mask_secrets || false, response_format);
    default:
      throw new Error(`Unknown operation: ${operation}`);
  }
}

// =============================================================================
// OPERATION HANDLERS
// =============================================================================

/**
 * Handle 'status' operation - show current auth state
 */
async function handleStatus(format: ResponseFormat): Promise<string> {
  if (!credentialManager) {
    throw new Error('Auth not initialized. Server startup issue.');
  }

  const profileManager = getProfileManager();
  const activeProfile = await profileManager.getActive();

  const status = {
    mode: credentialManager.getAuthMode(),
    authenticated: credentialManager.isAuthenticated(),
    tenant: credentialManager.getTenant(),
    api_url: credentialManager.getApiUrl(),
    namespace: credentialManager.getNamespace(),
    active_profile: activeProfile,
    source: getAuthSource(credentialManager),
  };

  if (format === ResponseFormat.JSON) {
    return JSON.stringify(status, null, 2);
  }

  return formatStatusMarkdown(status);
}

function getAuthSource(cm: CredentialManager): string {
  if (cm.getAuthMode() === AuthMode.NONE) {
    return 'none (documentation mode)';
  }

  // Check for environment variables
  if (process.env.F5XC_API_TOKEN || process.env.F5XC_P12_FILE || process.env.F5XC_CERT) {
    return 'environment variables';
  }

  return 'profile';
}

function formatStatusMarkdown(status: {
  mode: string;
  authenticated: boolean;
  tenant: string | null;
  api_url: string | null;
  namespace: string | null;
  active_profile: string | null;
  source: string;
}): string {
  const lines = [
    '# F5XC Authentication Status',
    '',
    '| Property | Value |',
    '|----------|-------|',
    `| **Mode** | ${status.mode} |`,
    `| **Authenticated** | ${status.authenticated ? 'Yes' : 'No'} |`,
    `| **Source** | ${status.source} |`,
  ];

  if (status.tenant) {
    lines.push(`| **Tenant** | ${status.tenant} |`);
  }

  if (status.api_url) {
    lines.push(`| **API URL** | ${status.api_url} |`);
  }

  if (status.namespace) {
    lines.push(`| **Default Namespace** | ${status.namespace} |`);
  }

  if (status.active_profile) {
    lines.push(`| **Active Profile** | ${status.active_profile} |`);
  }

  lines.push('');

  if (!status.authenticated) {
    lines.push('> **Documentation Mode**: No credentials configured.');
    lines.push('> The MCP server is operating in documentation-only mode.');
    lines.push('> To enable API access, configure credentials via environment variables or profiles.');
    lines.push('');
    lines.push('## Configuration Options');
    lines.push('');
    lines.push('### Environment Variables');
    lines.push('```bash');
    lines.push('export F5XC_API_URL="https://your-tenant.console.ves.volterra.io"');
    lines.push('export F5XC_API_TOKEN="your-api-token"');
    lines.push('```');
    lines.push('');
    lines.push('### Profile Configuration');
    lines.push('```bash');
    lines.push('# Create profile at ~/.config/f5xc/profiles/<name>.json');
    lines.push('```');
  }

  return lines.join('\n');
}

/**
 * Handle 'list' operation - list available profiles
 */
async function handleList(format: ResponseFormat): Promise<string> {
  const profileManager = getProfileManager();
  const profiles = await profileManager.list();
  const activeProfile = await profileManager.getActive();

  const profileList = profiles.map((p) => ({
    name: p.name,
    api_url: p.apiUrl,
    tenant: extractTenant(p.apiUrl),
    namespace: p.defaultNamespace || null,
    auth_method: getProfileAuthMethod(p),
    is_active: p.name === activeProfile,
  }));

  if (format === ResponseFormat.JSON) {
    return JSON.stringify({ profiles: profileList, active: activeProfile }, null, 2);
  }

  return formatProfileListMarkdown(profileList, activeProfile);
}

function extractTenant(apiUrl: string): string | null {
  const match = apiUrl.match(/https:\/\/([^.]+)\.console\.ves\.volterra\.io/);
  return match ? match[1] : null;
}

function getProfileAuthMethod(profile: Profile): string {
  if (profile.apiToken) return 'token';
  if (profile.p12Bundle) return 'p12';
  if (profile.cert && profile.key) return 'cert';
  return 'none';
}

function formatProfileListMarkdown(
  profiles: Array<{
    name: string;
    api_url: string;
    tenant: string | null;
    namespace: string | null;
    auth_method: string;
    is_active: boolean;
  }>,
  activeProfile: string | null,
): string {
  if (profiles.length === 0) {
    return [
      '# F5XC Profiles',
      '',
      'No profiles configured.',
      '',
      'Create a profile at `~/.config/f5xc/profiles/<name>.json`:',
      '',
      '```json',
      '{',
      '  "name": "production",',
      '  "apiUrl": "https://your-tenant.console.ves.volterra.io/api",',
      '  "apiToken": "your-api-token",',
      '  "defaultNamespace": "my-namespace"',
      '}',
      '```',
    ].join('\n');
  }

  const lines = [
    '# F5XC Profiles',
    '',
    `Active: ${activeProfile || 'none'}`,
    '',
    '| Profile | Tenant | Namespace | Auth | Active |',
    '|---------|--------|-----------|------|--------|',
  ];

  for (const p of profiles) {
    const active = p.is_active ? '**Yes**' : 'No';
    lines.push(
      `| ${p.name} | ${p.tenant || '-'} | ${p.namespace || '-'} | ${p.auth_method} | ${active} |`,
    );
  }

  lines.push('');
  lines.push(
    '> Use `operation: "switch", profile_name: "<name>"` to change the active profile.',
  );

  return lines.join('\n');
}

/**
 * Handle 'switch' operation - switch to a different profile
 */
async function handleSwitch(
  profileName: string | undefined,
  format: ResponseFormat,
): Promise<string> {
  if (!profileName) {
    throw new Error('profile_name is required for switch operation');
  }

  const profileManager = getProfileManager();

  // Check if profile exists
  const profile = await profileManager.get(profileName);
  if (!profile) {
    throw new Error(`Profile '${profileName}' not found`);
  }

  // Set as active profile
  const result = await profileManager.setActive(profileName);
  if (!result.success) {
    throw new Error(`Failed to switch profile: ${result.message}`);
  }

  // Reload credential manager with new profile
  if (credentialManager) {
    await credentialManager.reload();
  }

  const response = {
    success: true,
    profile: profileName,
    message: `Switched to profile '${profileName}'`,
    tenant: credentialManager?.getTenant() || null,
    authenticated: credentialManager?.isAuthenticated() || false,
  };

  if (format === ResponseFormat.JSON) {
    return JSON.stringify(response, null, 2);
  }

  const lines = [
    '# Profile Switched',
    '',
    `Successfully switched to profile: **${profileName}**`,
    '',
  ];

  if (response.authenticated) {
    lines.push(`- Tenant: ${response.tenant}`);
    lines.push('- Status: Authenticated');
  } else {
    lines.push('- Status: Not authenticated (check credentials)');
  }

  return lines.join('\n');
}

/**
 * Handle 'validate' operation - test API connectivity
 */
async function handleValidate(format: ResponseFormat): Promise<string> {
  if (!credentialManager) {
    throw new Error('Auth not initialized. Server startup issue.');
  }

  if (!credentialManager.isAuthenticated()) {
    const response = {
      valid: false,
      mode: credentialManager.getAuthMode(),
      error: 'No credentials configured (documentation mode)',
      suggestion: 'Configure F5XC_API_TOKEN environment variable or create a profile',
    };

    if (format === ResponseFormat.JSON) {
      return JSON.stringify(response, null, 2);
    }

    return [
      '# Credential Validation',
      '',
      '**Result**: Not validated',
      '',
      'No credentials are configured. The MCP server is operating in documentation mode.',
      '',
      '## To enable API access:',
      '',
      '1. Set environment variables:',
      '   ```bash',
      '   export F5XC_API_URL="https://your-tenant.console.ves.volterra.io"',
      '   export F5XC_API_TOKEN="your-api-token"',
      '   ```',
      '',
      '2. Or create a profile at `~/.config/f5xc/profiles/<name>.json`',
    ].join('\n');
  }

  // Test API connectivity
  const apiUrl = credentialManager.getApiUrl();
  const token = credentialManager.getToken();

  if (!apiUrl || !token) {
    throw new Error('Missing API URL or token');
  }

  try {
    // Make a simple API call to verify credentials
    const response = await axios.get(`${apiUrl}/web/namespaces`, {
      headers: {
        Authorization: `APIToken ${token}`,
        'Content-Type': 'application/json',
      },
      timeout: 10000,
    });

    const namespaceCount = response.data?.items?.length || 0;

    const result = {
      valid: true,
      tenant: credentialManager.getTenant(),
      api_url: apiUrl,
      namespaces_found: namespaceCount,
      message: `Successfully connected to F5XC API. Found ${namespaceCount} namespace(s).`,
    };

    if (format === ResponseFormat.JSON) {
      return JSON.stringify(result, null, 2);
    }

    return [
      '# Credential Validation',
      '',
      '**Result**: Valid',
      '',
      '| Property | Value |',
      '|----------|-------|',
      `| **Tenant** | ${result.tenant} |`,
      `| **API URL** | ${result.api_url} |`,
      `| **Namespaces Found** | ${result.namespaces_found} |`,
      '',
      'Credentials are valid and API is accessible.',
    ].join('\n');
  } catch (error) {
    const axiosError = error as { response?: { status: number; data?: { message?: string } }; message?: string };

    const result = {
      valid: false,
      tenant: credentialManager.getTenant(),
      api_url: apiUrl,
      error: axiosError.response?.data?.message || axiosError.message || 'Unknown error',
      status_code: axiosError.response?.status || null,
    };

    if (format === ResponseFormat.JSON) {
      return JSON.stringify(result, null, 2);
    }

    return [
      '# Credential Validation',
      '',
      '**Result**: Failed',
      '',
      '| Property | Value |',
      '|----------|-------|',
      `| **Tenant** | ${result.tenant} |`,
      `| **API URL** | ${result.api_url} |`,
      `| **Error** | ${result.error} |`,
      result.status_code ? `| **Status Code** | ${result.status_code} |` : '',
      '',
      '## Troubleshooting',
      '',
      '- Verify your API token is valid and not expired',
      '- Check that the API URL is correct',
      '- Ensure network connectivity to F5 Distributed Cloud',
    ]
      .filter((l) => l)
      .join('\n');
  }
}

// =============================================================================
// TERRAFORM-ENV OPERATION
// =============================================================================

/**
 * Environment variables interface for Terraform provider
 */
interface TerraformEnvVars {
  F5XC_API_URL?: string;
  F5XC_API_TOKEN?: string;
  F5XC_P12_FILE?: string;
  F5XC_P12_PASSWORD?: string;
  F5XC_CERT?: string;
  F5XC_KEY?: string;
  F5XC_CACERT?: string;
}

/**
 * Mask sensitive values for display
 */
function maskSecret(value: string, showLast = 8): string {
  if (value.length <= showLast) {
    return '****';
  }
  return '****...' + value.slice(-showLast);
}

/**
 * Build environment variables from credential manager
 */
function buildEnvironmentVariables(
  cm: CredentialManager,
  maskSecrets: boolean,
): TerraformEnvVars {
  const vars: TerraformEnvVars = {};

  // Always include API URL
  const apiUrl = cm.getApiUrl();
  if (apiUrl) {
    vars.F5XC_API_URL = apiUrl;
  }

  // Handle auth method
  const authMode = cm.getAuthMode();

  if (authMode === AuthMode.TOKEN) {
    const token = cm.getToken();
    if (token) {
      vars.F5XC_API_TOKEN = maskSecrets ? maskSecret(token) : token;
    }
  } else if (authMode === AuthMode.CERTIFICATE) {
    // For certificate auth, we need the file paths from the profile
    // The CredentialManager loads the certificate data, but we need paths for env vars
    // Check if credentials came from profile to get file paths
    const profileManager = getProfileManager();
    const activeProfileName = cm.getActiveProfile ? cm.getActiveProfile() : null;

    // For now, indicate that cert-based auth requires file paths
    // The profile stores the paths, not the data, so we can retrieve them
    vars.F5XC_P12_FILE = '<configure-p12-path>';
    vars.F5XC_CERT = '<configure-cert-path>';
    vars.F5XC_KEY = '<configure-key-path>';
  }

  // Optional namespace is not a provider env var, skip it
  // F5XC_NAMESPACE is used by the auth package, not the Terraform provider directly

  return vars;
}

/**
 * Format shell export statements
 */
function formatShellExports(
  envVars: TerraformEnvVars,
  profileName: string | null,
  authMode: string,
): string {
  const lines: string[] = [
    '#!/bin/bash',
    '# F5XC Terraform Provider Configuration',
    `# Profile: ${profileName || 'environment'}`,
    `# Auth Method: ${authMode}`,
    '# WARNING: Contains sensitive credentials - do not share or commit',
    '',
  ];

  for (const [key, value] of Object.entries(envVars)) {
    if (value) {
      lines.push(`export ${key}="${value}"`);
    }
  }

  return lines.join('\n');
}

/**
 * Format .env file content
 */
function formatDotEnv(
  envVars: TerraformEnvVars,
  profileName: string | null,
): string {
  const lines: string[] = [
    '# F5XC Terraform Provider Configuration',
    `# Profile: ${profileName || 'environment'}`,
    '',
  ];

  for (const [key, value] of Object.entries(envVars)) {
    if (value) {
      lines.push(`${key}=${value}`);
    }
  }

  return lines.join('\n');
}

/**
 * Format JSON output
 */
function formatEnvJson(
  envVars: TerraformEnvVars,
  profileName: string | null,
  authMode: string,
): string {
  return JSON.stringify({
    profile: profileName,
    auth_method: authMode,
    variables: envVars,
    shell_command: Object.entries(envVars)
      .filter(([_, v]) => v)
      .map(([k, v]) => `export ${k}="${v}"`)
      .join('; '),
  }, null, 2);
}

/**
 * Handle 'terraform-env' operation - generate Terraform environment variables
 */
async function handleTerraformEnv(
  outputType: 'shell' | 'dotenv' | 'json',
  maskSecrets: boolean,
  format: ResponseFormat,
): Promise<string> {
  if (!credentialManager) {
    throw new Error('Auth not initialized. Server startup issue.');
  }

  if (!credentialManager.isAuthenticated()) {
    const error = {
      error: 'No credentials configured (documentation mode)',
      suggestion: 'Configure F5XC_API_TOKEN environment variable or create a profile at ~/.config/f5xc/profiles/<name>.json',
    };

    if (format === ResponseFormat.JSON) {
      return JSON.stringify(error, null, 2);
    }

    return [
      '# Terraform Environment Variables',
      '',
      '**Error**: No credentials configured',
      '',
      'The MCP server is operating in documentation mode. To generate environment variables:',
      '',
      '1. Set environment variables directly:',
      '   ```bash',
      '   export F5XC_API_URL="https://your-tenant.console.ves.volterra.io"',
      '   export F5XC_API_TOKEN="your-api-token"',
      '   ```',
      '',
      '2. Or create a profile at `~/.config/f5xc/profiles/<name>.json`',
    ].join('\n');
  }

  const profileManager = getProfileManager();
  const activeProfile = await profileManager.getActive();
  const authMode = credentialManager.getAuthMode();
  const envVars = buildEnvironmentVariables(credentialManager, maskSecrets);

  // For JSON response format, always return JSON structure
  if (format === ResponseFormat.JSON) {
    return formatEnvJson(envVars, activeProfile, authMode);
  }

  // For Markdown response format, format based on output_type
  switch (outputType) {
    case 'shell':
      return [
        '# Terraform Environment Variables',
        '',
        `**Profile**: ${activeProfile || 'environment'}`,
        `**Auth Method**: ${authMode}`,
        '',
        '```bash',
        formatShellExports(envVars, activeProfile, authMode),
        '```',
        '',
        '## Usage',
        '',
        'Copy and paste the export commands into your terminal, or run:',
        '',
        '```bash',
        '# Source directly (if saved to file)',
        'source ./f5xc-env.sh',
        '',
        '# Then run Terraform',
        'terraform plan',
        '```',
      ].join('\n');

    case 'dotenv':
      return [
        '# Terraform Environment Variables (.env format)',
        '',
        `**Profile**: ${activeProfile || 'environment'}`,
        '',
        '```',
        formatDotEnv(envVars, activeProfile),
        '```',
        '',
        '## Usage',
        '',
        'Save this content to a `.env` file and use with dotenv-compatible tools.',
      ].join('\n');

    case 'json':
      return [
        '# Terraform Environment Variables (JSON)',
        '',
        '```json',
        formatEnvJson(envVars, activeProfile, authMode),
        '```',
      ].join('\n');

    default:
      throw new Error(`Unknown output type: ${outputType}`);
  }
}
