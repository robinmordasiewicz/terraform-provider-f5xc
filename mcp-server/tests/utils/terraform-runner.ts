// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

/**
 * Terraform Runner Utility
 *
 * Executes Terraform CLI commands for E2E testing.
 * Validates determinism by tracking apply attempts.
 */

import { execSync, spawn } from 'child_process';
import * as fs from 'fs';
import * as path from 'path';
import * as os from 'os';

export interface TerraformResult {
  success: boolean;
  output: string;
  errors: string[];
  exitCode: number;
}

export interface TerraformApplyResult extends TerraformResult {
  attempts: number;
}

export interface TerraformState {
  resources: Array<{
    type: string;
    name: string;
    values: Record<string, unknown>;
  }>;
}

/**
 * Check if Terraform CLI is available
 */
export function isTerraformAvailable(): boolean {
  try {
    execSync('terraform version', { stdio: 'pipe' });
    return true;
  } catch {
    return false;
  }
}

/**
 * Get Terraform version
 */
export function getTerraformVersion(): string | null {
  try {
    const output = execSync('terraform version -json', { encoding: 'utf-8' });
    const data = JSON.parse(output);
    return data.terraform_version;
  } catch {
    return null;
  }
}

/**
 * Terraform Runner class for E2E testing
 */
export class TerraformRunner {
  private workDir: string;
  private testName: string;
  private initialized: boolean = false;
  private providerPath: string;

  constructor(testName: string) {
    this.testName = testName;
    // Create unique working directory for this test
    const timestamp = Date.now();
    const sanitizedName = testName.replace(/[^a-zA-Z0-9]/g, '-');
    this.workDir = path.join(os.tmpdir(), 'f5xc-mcp-tests', `${sanitizedName}-${timestamp}`);

    // Path to the local provider build
    this.providerPath = path.resolve(__dirname, '../../../../terraform-provider-f5xc');
  }

  /**
   * Get the working directory path
   */
  getWorkDir(): string {
    return this.workDir;
  }

  /**
   * Setup the working directory
   */
  async setup(): Promise<void> {
    // Create working directory
    fs.mkdirSync(this.workDir, { recursive: true });

    // Create terraform configuration for local provider
    // dev_overrides needs the directory containing the provider binary
    // providerPath is already the project root (terraform-provider-f5xc/) containing the binary
    const terraformRc = `
provider_installation {
  dev_overrides {
    "registry.terraform.io/robinmordasiewicz/f5xc" = "${this.providerPath}"
  }
  direct {}
}
`;
    // Write .terraformrc to use local provider
    const rcPath = path.join(this.workDir, '.terraformrc');
    fs.writeFileSync(rcPath, terraformRc);

    // Set environment variable to use local terraformrc
    process.env.TF_CLI_CONFIG_FILE = rcPath;
  }

  /**
   * Write a Terraform configuration file
   */
  writeConfig(filename: string, content: string): void {
    const filePath = path.join(this.workDir, filename);
    fs.writeFileSync(filePath, content);
  }

  /**
   * Write provider configuration
   */
  writeProviderConfig(apiUrl: string): void {
    const providerConfig = `
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
    this.writeConfig('provider.tf', providerConfig);
  }

  /**
   * Run terraform init
   * Note: When using dev_overrides, init may fail because it tries to query
   * the registry. In this case, we return success since dev_overrides loads
   * the provider directly without needing init.
   */
  async init(): Promise<TerraformResult> {
    const result = await this.runCommand(['init', '-no-color', '-input=false']);

    // When using dev_overrides, terraform init may fail because it tries to
    // query the registry, but validate/apply will still work because the
    // provider is loaded directly from the dev_overrides path.
    // Check if the error is just about provider registry query failure.
    if (!result.success) {
      const output = result.output.toLowerCase();
      if (output.includes('provider development overrides are in effect') ||
          output.includes('skip terraform init when using provider development overrides')) {
        // This is expected with dev_overrides - treat as success
        this.initialized = true;
        return {
          success: true,
          output: result.output,
          errors: [],
          exitCode: 0,
        };
      }
    }

    this.initialized = result.success;
    return result;
  }

  /**
   * Run terraform validate
   */
  async validate(): Promise<TerraformResult> {
    const result = await this.runCommand(['validate', '-no-color', '-json']);

    // Parse JSON output for validation
    try {
      const jsonResult = JSON.parse(result.output);
      return {
        success: jsonResult.valid === true,
        output: result.output,
        errors: jsonResult.diagnostics?.filter((d: { severity: string }) => d.severity === 'error')
          .map((d: { summary: string; detail?: string }) => `${d.summary}: ${d.detail || ''}`) || [],
        exitCode: result.exitCode,
      };
    } catch {
      // If JSON parsing fails, use the raw result
      return result;
    }
  }

  /**
   * Run terraform plan
   */
  async plan(): Promise<TerraformResult> {
    return this.runCommand(['plan', '-no-color', '-input=false']);
  }

  /**
   * Run terraform apply - MUST succeed on first attempt for determinism test
   */
  async apply(): Promise<TerraformApplyResult> {
    const result = await this.runCommand(['apply', '-auto-approve', '-no-color', '-input=false']);

    return {
      ...result,
      attempts: 1, // We only allow one attempt - determinism requirement
    };
  }

  /**
   * Run terraform destroy
   */
  async destroy(): Promise<TerraformResult> {
    return this.runCommand(['destroy', '-auto-approve', '-no-color', '-input=false']);
  }

  /**
   * Run terraform fmt -check to verify configuration is formatted
   * Returns success if config is already properly formatted
   */
  async fmt(): Promise<TerraformResult> {
    const result = await this.runCommand(['fmt', '-check', '-diff', '-no-color']);
    return {
      success: result.exitCode === 0,
      output: result.output,
      errors: result.exitCode !== 0 ? ['Configuration is not properly formatted. Run: terraform fmt'] : [],
      exitCode: result.exitCode,
    };
  }

  /**
   * Run terraform fmt to format configuration in place
   * Returns the list of files that were modified
   */
  async fmtWrite(): Promise<TerraformResult> {
    return this.runCommand(['fmt', '-no-color']);
  }

  /**
   * Get terraform state
   */
  async getState(): Promise<TerraformState | null> {
    const result = await this.runCommand(['show', '-json']);
    if (!result.success) {
      return null;
    }

    try {
      const state = JSON.parse(result.output);
      const resources: TerraformState['resources'] = [];

      if (state.values?.root_module?.resources) {
        for (const resource of state.values.root_module.resources) {
          resources.push({
            type: resource.type,
            name: resource.name,
            values: resource.values,
          });
        }
      }

      return { resources };
    } catch {
      return null;
    }
  }

  /**
   * Get output values
   */
  async getOutputs(): Promise<Record<string, unknown>> {
    const result = await this.runCommand(['output', '-json']);
    if (!result.success) {
      return {};
    }

    try {
      return JSON.parse(result.output);
    } catch {
      return {};
    }
  }

  /**
   * Cleanup working directory
   */
  async cleanup(): Promise<void> {
    try {
      // First try to destroy resources
      await this.destroy();
    } catch {
      // Ignore destroy errors during cleanup
    }

    try {
      // Remove working directory
      fs.rmSync(this.workDir, { recursive: true, force: true });
    } catch {
      // Ignore cleanup errors
    }
  }

  /**
   * Run a terraform command
   */
  private async runCommand(args: string[]): Promise<TerraformResult> {
    return new Promise((resolve) => {
      const stdout: string[] = [];
      const stderr: string[] = [];

      const proc = spawn('terraform', args, {
        cwd: this.workDir,
        env: {
          ...process.env,
          TF_IN_AUTOMATION: 'true',
          TF_CLI_CONFIG_FILE: path.join(this.workDir, '.terraformrc'),
        },
      });

      proc.stdout.on('data', (data) => {
        stdout.push(data.toString());
      });

      proc.stderr.on('data', (data) => {
        stderr.push(data.toString());
      });

      proc.on('close', (code) => {
        const output = stdout.join('');
        const errors = stderr.join('');

        resolve({
          success: code === 0,
          output: output + errors,
          errors: errors ? [errors] : [],
          exitCode: code || 0,
        });
      });

      proc.on('error', (error) => {
        resolve({
          success: false,
          output: '',
          errors: [error.message],
          exitCode: 1,
        });
      });
    });
  }
}

/**
 * Skip test if Terraform is not available
 */
export function skipIfNoTerraform(): void {
  if (!isTerraformAvailable()) {
    throw new Error('Terraform CLI is not available - skipping E2E tests');
  }
}
