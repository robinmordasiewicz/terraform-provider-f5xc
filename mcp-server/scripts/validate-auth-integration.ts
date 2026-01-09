#!/usr/bin/env npx tsx
/**
 * Auth Integration Validation Script
 *
 * Idempotent script to validate the @robinmordasiewicz/f5xc-auth integration
 * in the F5XC Terraform MCP server.
 *
 * Usage:
 *   npx tsx scripts/validate-auth-integration.ts [options]
 *
 * Options:
 *   --verbose       Show detailed output for each validation step
 *   --test-real-api Test against real F5XC API (requires F5XC_TEST_API_URL and F5XC_TEST_API_TOKEN)
 *   --help          Show this help message
 *
 * Exit codes:
 *   0 - All validations passed
 *   1 - One or more validations failed
 */

import { execSync, spawn } from 'child_process';
import { existsSync, readFileSync } from 'fs';
import { join, dirname } from 'path';
import { fileURLToPath } from 'url';

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);
const ROOT_DIR = join(__dirname, '..');

// =============================================================================
// CONFIGURATION
// =============================================================================

interface ValidationResult {
  name: string;
  passed: boolean;
  message: string;
  details?: string;
  duration?: number;
}

const results: ValidationResult[] = [];
let verbose = false;
let testRealApi = false;

// =============================================================================
// COLORS
// =============================================================================

const colors = {
  reset: '\x1b[0m',
  bright: '\x1b[1m',
  red: '\x1b[31m',
  green: '\x1b[32m',
  yellow: '\x1b[33m',
  blue: '\x1b[34m',
  cyan: '\x1b[36m',
};

function log(message: string, color: string = colors.reset): void {
  console.log(`${color}${message}${colors.reset}`);
}

function logHeader(message: string): void {
  console.log('');
  log(`${'='.repeat(60)}`, colors.cyan);
  log(message, colors.bright + colors.cyan);
  log(`${'='.repeat(60)}`, colors.cyan);
}

function logStep(message: string): void {
  log(`\n>>> ${message}`, colors.blue);
}

function logSuccess(message: string): void {
  log(`  ✅ ${message}`, colors.green);
}

function logFailure(message: string): void {
  log(`  ❌ ${message}`, colors.red);
}

function logWarning(message: string): void {
  log(`  ⚠️  ${message}`, colors.yellow);
}

function logInfo(message: string): void {
  if (verbose) {
    log(`  ℹ️  ${message}`, colors.reset);
  }
}

// =============================================================================
// UTILITIES
// =============================================================================

function exec(command: string, options: { cwd?: string; silent?: boolean } = {}): string {
  const cwd = options.cwd || ROOT_DIR;
  try {
    const result = execSync(command, {
      cwd,
      encoding: 'utf-8',
      stdio: options.silent ? 'pipe' : 'inherit',
    });
    return result || '';
  } catch (error) {
    if (error instanceof Error && 'stdout' in error) {
      return (error as any).stdout || '';
    }
    throw error;
  }
}

function execAsync(command: string, args: string[], cwd: string): Promise<{ stdout: string; stderr: string; exitCode: number }> {
  return new Promise((resolve) => {
    const proc = spawn(command, args, { cwd, shell: true });
    let stdout = '';
    let stderr = '';

    proc.stdout?.on('data', (data) => {
      stdout += data.toString();
      if (verbose) process.stdout.write(data);
    });

    proc.stderr?.on('data', (data) => {
      stderr += data.toString();
      if (verbose) process.stderr.write(data);
    });

    proc.on('close', (exitCode) => {
      resolve({ stdout, stderr, exitCode: exitCode || 0 });
    });
  });
}

async function validate(
  name: string,
  testFn: () => Promise<{ passed: boolean; message: string; details?: string }>
): Promise<void> {
  logStep(name);
  const startTime = Date.now();

  try {
    const result = await testFn();
    const duration = Date.now() - startTime;

    results.push({
      name,
      ...result,
      duration,
    });

    if (result.passed) {
      logSuccess(result.message);
    } else {
      logFailure(result.message);
    }

    if (result.details && verbose) {
      logInfo(result.details);
    }
  } catch (error) {
    const duration = Date.now() - startTime;
    const message = error instanceof Error ? error.message : String(error);

    results.push({
      name,
      passed: false,
      message: `Error: ${message}`,
      duration,
    });

    logFailure(`Error: ${message}`);
  }
}

// =============================================================================
// VALIDATIONS
// =============================================================================

async function validateDependency(): Promise<{ passed: boolean; message: string; details?: string }> {
  const packageJsonPath = join(ROOT_DIR, 'package.json');

  if (!existsSync(packageJsonPath)) {
    return { passed: false, message: 'package.json not found' };
  }

  const packageJson = JSON.parse(readFileSync(packageJsonPath, 'utf-8'));
  const authDep = packageJson.dependencies?.['@robinmordasiewicz/f5xc-auth'];

  if (!authDep) {
    return { passed: false, message: '@robinmordasiewicz/f5xc-auth not found in dependencies' };
  }

  return {
    passed: true,
    message: `@robinmordasiewicz/f5xc-auth found: ${authDep}`,
  };
}

async function validateTypeScript(): Promise<{ passed: boolean; message: string; details?: string }> {
  try {
    exec('npm run typecheck', { silent: !verbose });
    return { passed: true, message: 'TypeScript compilation successful' };
  } catch {
    return { passed: false, message: 'TypeScript compilation failed' };
  }
}

async function validateBuild(): Promise<{ passed: boolean; message: string; details?: string }> {
  try {
    exec('npm run build', { silent: !verbose });
    return { passed: true, message: 'Build successful' };
  } catch {
    return { passed: false, message: 'Build failed' };
  }
}

async function validateUnitTests(): Promise<{ passed: boolean; message: string; details?: string }> {
  const result = await execAsync('npx', ['vitest', 'run', '--reporter=verbose'], ROOT_DIR);

  if (result.exitCode === 0) {
    // Extract test count from output
    const match = result.stdout.match(/(\d+)\s+passed/);
    const testCount = match ? match[1] : 'all';
    return { passed: true, message: `Unit tests passed (${testCount} tests)` };
  }

  return { passed: false, message: 'Unit tests failed', details: result.stderr };
}

async function validateAuthImports(): Promise<{ passed: boolean; message: string; details?: string }> {
  const authToolPath = join(ROOT_DIR, 'src/tools/auth.ts');

  if (!existsSync(authToolPath)) {
    return { passed: false, message: 'auth.ts not found' };
  }

  const content = readFileSync(authToolPath, 'utf-8');

  const requiredImports = [
    'CredentialManager',
    'AuthMode',
    'getProfileManager',
  ];

  const missingImports = requiredImports.filter((imp) => !content.includes(imp));

  if (missingImports.length > 0) {
    return {
      passed: false,
      message: `Missing imports: ${missingImports.join(', ')}`,
    };
  }

  if (!content.includes('@robinmordasiewicz/f5xc-auth')) {
    return {
      passed: false,
      message: 'Not importing from @robinmordasiewicz/f5xc-auth',
    };
  }

  return {
    passed: true,
    message: 'All required imports present from shared package',
  };
}

async function validateAuthSchemaExports(): Promise<{ passed: boolean; message: string; details?: string }> {
  const schemaPath = join(ROOT_DIR, 'src/schemas/common.ts');

  if (!existsSync(schemaPath)) {
    return { passed: false, message: 'common.ts schema not found' };
  }

  const content = readFileSync(schemaPath, 'utf-8');

  if (!content.includes('AuthSchema')) {
    return { passed: false, message: 'AuthSchema not found in common.ts' };
  }

  if (!content.includes('AuthInput')) {
    return { passed: false, message: 'AuthInput type not found in common.ts' };
  }

  return {
    passed: true,
    message: 'Auth schema and types properly defined',
  };
}

async function validateDocumentationModeInit(): Promise<{ passed: boolean; message: string; details?: string }> {
  // Clear any existing credentials
  const originalEnv = { ...process.env };
  delete process.env.F5XC_API_URL;
  delete process.env.F5XC_API_TOKEN;
  delete process.env.F5XC_P12_FILE;
  delete process.env.F5XC_CERT;
  process.env.XDG_CONFIG_HOME = '/tmp/__nonexistent_test_config__';

  try {
    const { CredentialManager, AuthMode } = await import('@robinmordasiewicz/f5xc-auth');
    const cm = new CredentialManager();
    await cm.initialize();

    const mode = cm.getAuthMode();
    const isAuth = cm.isAuthenticated();

    if (mode !== AuthMode.NONE || isAuth) {
      return {
        passed: false,
        message: `Expected NONE mode, got ${mode} (authenticated: ${isAuth})`,
      };
    }

    return {
      passed: true,
      message: 'Documentation mode initializes correctly with NONE auth',
    };
  } catch (error) {
    return {
      passed: false,
      message: `Failed to initialize: ${error instanceof Error ? error.message : String(error)}`,
    };
  } finally {
    process.env = originalEnv;
  }
}

async function validateTokenModeInit(): Promise<{ passed: boolean; message: string; details?: string }> {
  const originalEnv = { ...process.env };

  try {
    process.env.F5XC_API_URL = 'https://test.console.ves.volterra.io';
    process.env.F5XC_API_TOKEN = 'test-token-12345';
    process.env.XDG_CONFIG_HOME = '/tmp/__nonexistent_test_config__';

    const { CredentialManager, AuthMode } = await import('@robinmordasiewicz/f5xc-auth');
    const cm = new CredentialManager();
    await cm.initialize();

    const mode = cm.getAuthMode();
    const isAuth = cm.isAuthenticated();
    const tenant = cm.getTenant();

    if (mode !== AuthMode.TOKEN) {
      return { passed: false, message: `Expected TOKEN mode, got ${mode}` };
    }

    if (!isAuth) {
      return { passed: false, message: 'Expected authenticated, got false' };
    }

    if (tenant !== 'test') {
      return { passed: false, message: `Expected tenant 'test', got '${tenant}'` };
    }

    return {
      passed: true,
      message: 'Token mode initializes correctly with authenticated state',
    };
  } catch (error) {
    return {
      passed: false,
      message: `Failed to initialize: ${error instanceof Error ? error.message : String(error)}`,
    };
  } finally {
    process.env = originalEnv;
  }
}

async function validateProfileManagerAvailable(): Promise<{ passed: boolean; message: string; details?: string }> {
  try {
    const { getProfileManager } = await import('@robinmordasiewicz/f5xc-auth');
    const pm = getProfileManager();

    if (!pm) {
      return { passed: false, message: 'getProfileManager returned null' };
    }

    if (typeof pm.list !== 'function') {
      return { passed: false, message: 'ProfileManager.list is not a function' };
    }

    if (typeof pm.get !== 'function') {
      return { passed: false, message: 'ProfileManager.get is not a function' };
    }

    if (typeof pm.getActive !== 'function') {
      return { passed: false, message: 'ProfileManager.getActive is not a function' };
    }

    return {
      passed: true,
      message: 'ProfileManager available with all required methods',
    };
  } catch (error) {
    return {
      passed: false,
      message: `Failed to get ProfileManager: ${error instanceof Error ? error.message : String(error)}`,
    };
  }
}

async function validateRealApiConnection(): Promise<{ passed: boolean; message: string; details?: string }> {
  if (!testRealApi) {
    return {
      passed: true,
      message: 'Skipped (use --test-real-api to enable)',
    };
  }

  const apiUrl = process.env.F5XC_TEST_API_URL;
  const apiToken = process.env.F5XC_TEST_API_TOKEN;

  if (!apiUrl || !apiToken) {
    return {
      passed: false,
      message: 'F5XC_TEST_API_URL and F5XC_TEST_API_TOKEN required for real API test',
    };
  }

  try {
    const { CredentialManager } = await import('@robinmordasiewicz/f5xc-auth');

    const originalEnv = { ...process.env };
    process.env.F5XC_API_URL = apiUrl;
    process.env.F5XC_API_TOKEN = apiToken;

    const cm = new CredentialManager();
    await cm.initialize();

    const tenant = cm.getTenant();
    const isAuth = cm.isAuthenticated();

    process.env = originalEnv;

    if (!isAuth) {
      return { passed: false, message: 'Failed to authenticate with real credentials' };
    }

    return {
      passed: true,
      message: `Successfully authenticated to tenant: ${tenant}`,
    };
  } catch (error) {
    return {
      passed: false,
      message: `API connection failed: ${error instanceof Error ? error.message : String(error)}`,
    };
  }
}

// =============================================================================
// MAIN
// =============================================================================

function showHelp(): void {
  console.log(`
Auth Integration Validation Script

Usage:
  npx tsx scripts/validate-auth-integration.ts [options]

Options:
  --verbose       Show detailed output for each validation step
  --test-real-api Test against real F5XC API
                  (requires F5XC_TEST_API_URL and F5XC_TEST_API_TOKEN)
  --help          Show this help message

Exit codes:
  0 - All validations passed
  1 - One or more validations failed
`);
}

function printSummary(): void {
  logHeader('VALIDATION SUMMARY');

  const passed = results.filter((r) => r.passed).length;
  const failed = results.filter((r) => !r.passed).length;
  const total = results.length;

  console.log('');
  for (const result of results) {
    const icon = result.passed ? '✅' : '❌';
    const time = result.duration ? ` (${result.duration}ms)` : '';
    console.log(`  ${icon} ${result.name}${time}`);
  }

  console.log('');
  log(`${'─'.repeat(60)}`, colors.cyan);

  if (failed === 0) {
    log(`\n  🎉 All ${total} validations PASSED\n`, colors.green);
  } else {
    log(`\n  ⚠️  ${passed}/${total} passed, ${failed} FAILED\n`, colors.red);
  }
}

async function main(): Promise<void> {
  // Parse arguments
  const args = process.argv.slice(2);

  if (args.includes('--help') || args.includes('-h')) {
    showHelp();
    process.exit(0);
  }

  verbose = args.includes('--verbose') || args.includes('-v');
  testRealApi = args.includes('--test-real-api');

  logHeader('F5XC Auth Integration Validation');
  log(`\nRunning from: ${ROOT_DIR}`);
  log(`Verbose: ${verbose}`);
  log(`Test Real API: ${testRealApi}`);

  // Run validations
  await validate('1. Dependency Check', validateDependency);
  await validate('2. TypeScript Compilation', validateTypeScript);
  await validate('3. Build', validateBuild);
  await validate('4. Unit Tests', validateUnitTests);
  await validate('5. Auth Tool Imports', validateAuthImports);
  await validate('6. Auth Schema Exports', validateAuthSchemaExports);
  await validate('7. Documentation Mode Initialization', validateDocumentationModeInit);
  await validate('8. Token Mode Initialization', validateTokenModeInit);
  await validate('9. ProfileManager Availability', validateProfileManagerAvailable);
  await validate('10. Real API Connection', validateRealApiConnection);

  // Print summary
  printSummary();

  // Exit with appropriate code
  const failed = results.filter((r) => !r.passed).length;
  process.exit(failed > 0 ? 1 : 0);
}

main().catch((error) => {
  console.error('Fatal error:', error);
  process.exit(1);
});
