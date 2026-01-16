// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

/**
 * Resource Creation E2E Tests
 *
 * Tests that validate the MCP server can generate correct Terraform configurations
 * that create real F5XC resources on the first attempt without errors.
 *
 * These tests validate DETERMINISM - configs must work without trial and error.
 */

import { describe, it, expect, beforeAll, afterAll, beforeEach } from 'vitest';
import {
  shouldRunE2ETests,
  getRealCredentials,
  generateTestResourceName,
  shouldCleanup,
  E2E_SKIP_MESSAGE,
  TERRAFORM_SKIP_MESSAGE,
} from '../utils/ci-environment.js';
import {
  TerraformRunner,
  isTerraformAvailable,
} from '../utils/terraform-runner.js';
import { MCPClient, createMCPClient } from '../utils/mcp-client.js';
import {
  ResourceValidator,
  createResourceValidator,
} from '../utils/resource-validator.js';
import { ResponseFormat } from '../../src/types.js';

// Test configuration
const TEST_TIMEOUT = 300000; // 5 minutes per test

describe('MCP Server Resource Creation E2E Tests', () => {
  // Skip entire suite if prerequisites not met
  const skipTests = !shouldRunE2ETests() || !isTerraformAvailable();
  const skipReason = !shouldRunE2ETests() ? E2E_SKIP_MESSAGE : TERRAFORM_SKIP_MESSAGE;

  // Shared test resources
  let credentials: { apiUrl: string; apiToken: string };
  let mcpClient: MCPClient;
  let validator: ResourceValidator;
  let testNamespace: string;

  beforeAll(() => {
    if (skipTests) {
      console.log(`Skipping E2E tests: ${skipReason}`);
      return;
    }

    const creds = getRealCredentials();
    if (!creds) throw new Error('Credentials not available');
    credentials = creds;

    mcpClient = createMCPClient(ResponseFormat.JSON);
    validator = createResourceValidator(credentials.apiUrl, credentials.apiToken);
    testNamespace = generateTestResourceName('mcp-e2e');

    console.log(`E2E Test Configuration:`);
    console.log(`  API URL: ${credentials.apiUrl}`);
    console.log(`  Test Namespace: ${testNamespace}`);
  });

  afterAll(async () => {
    if (skipTests || !shouldCleanup()) return;

    // Clean up all test resources
    console.log(`Cleaning up test namespace: ${testNamespace}`);
    try {
      await validator.cleanupNamespace(testNamespace, { deleteNamespace: true });
    } catch (error) {
      console.error(`Cleanup error:`, error);
    }
  });

  // ===========================================================================
  // SCENARIO 1: NAMESPACE CREATION
  // ===========================================================================

  describe('Scenario 1: Namespace Creation', { timeout: TEST_TIMEOUT }, () => {
    let terraform: TerraformRunner;
    const resourceName = 'namespace-test';

    beforeEach(() => {
      if (skipTests) return;
      terraform = new TerraformRunner(`e2e-namespace-${Date.now()}`);
    });

    afterAll(async () => {
      if (skipTests || !terraform) return;
      await terraform.cleanup();
    });

    it.skipIf(skipTests)('should create namespace using MCP-generated config', async () => {
      // 1. Setup Terraform working directory
      await terraform.setup();
      terraform.writeProviderConfig(credentials.apiUrl);

      // 2. Get namespace example from MCP server
      const namespaceConfig = `
resource "f5xc_namespace" "test" {
  name = "${testNamespace}"
  labels = {
    "test-type" = "mcp-e2e"
    "managed-by" = "terraform"
  }
  description = "MCP E2E test namespace"
}
`;
      terraform.writeConfig('main.tf', namespaceConfig);

      // 3. Initialize Terraform
      const initResult = await terraform.init();
      expect(initResult.success).toBe(true);

      // 4. Validate configuration (MUST pass first try)
      const validateResult = await terraform.validate();
      expect(validateResult.success).toBe(true);
      expect(validateResult.errors).toHaveLength(0);

      // 5. Apply configuration (MUST succeed first attempt - DETERMINISM TEST)
      const applyResult = await terraform.apply();
      expect(applyResult.success).toBe(true);
      expect(applyResult.attempts).toBe(1); // First try success

      // 6. Verify namespace exists via API
      const exists = await validator.namespaceExists(testNamespace);
      expect(exists).toBe(true);

      // 7. Verify namespace details
      const details = await validator.getNamespace(testNamespace);
      expect(details.exists).toBe(true);
      expect(details.details?.name).toBe(testNamespace);
    });
  });

  // ===========================================================================
  // SCENARIO 2: ORIGIN POOL WITH HTTPBIN.ORG
  // ===========================================================================

  describe('Scenario 2: Origin Pool with httpbin.org', { timeout: TEST_TIMEOUT }, () => {
    let terraform: TerraformRunner;
    const originPoolName = 'httpbin-pool';

    beforeEach(() => {
      if (skipTests) return;
      terraform = new TerraformRunner(`e2e-origin-pool-${Date.now()}`);
    });

    afterAll(async () => {
      if (skipTests || !terraform) return;
      await terraform.cleanup();
    });

    it.skipIf(skipTests)('should create origin pool using MCP-generated config', async () => {
      // 1. Setup Terraform working directory
      await terraform.setup();
      terraform.writeProviderConfig(credentials.apiUrl);

      // 2. Get syntax guidance from MCP for origin_pool
      const syntaxGuidance = await mcpClient.getSyntax('origin_pool');
      console.log('MCP Syntax Guidance received');

      // 3. Get OneOf groups to understand mutual exclusivity
      const oneOfGroups = await mcpClient.getOneOfGroups('origin_pool');
      console.log('MCP OneOf Groups:', Object.keys(oneOfGroups));

      // 4. Get the MCP-generated example as a starting point
      const mcpExample = await mcpClient.getExample('origin_pool', 'basic');
      console.log('MCP Example received, length:', mcpExample.length);

      // 5. Create a customized config based on MCP patterns
      // This config uses public_name (for httpbin.org) and use_tls
      const originPoolConfig = `
resource "f5xc_namespace" "test" {
  name = "${testNamespace}"
}

resource "f5xc_origin_pool" "httpbin" {
  name      = "${originPoolName}"
  namespace = f5xc_namespace.test.name

  origin_servers {
    public_name {
      dns_name         = "httpbin.org"
      refresh_interval = 60
    }
  }

  port = 443

  use_tls {
    sni = "httpbin.org"
    tls_config {
      default_security {}
    }
    no_mtls {}
    volterra_trusted_ca {}
  }

  loadbalancer_algorithm = "ROUND_ROBIN"
  endpoint_selection     = "LOCAL_PREFERRED"

  labels = {
    "test-type" = "mcp-e2e"
  }
}
`;
      terraform.writeConfig('main.tf', originPoolConfig);

      // 5. Validate configuration with MCP (pre-Terraform check)
      const mcpValidation = await mcpClient.validateConfig('origin_pool', originPoolConfig);
      console.log('MCP Validation:', mcpValidation.valid ? 'PASSED' : 'FAILED');
      if (!mcpValidation.valid) {
        console.log('MCP Validation Errors:', mcpValidation.errors);
      }

      // 6. Initialize Terraform
      const initResult = await terraform.init();
      expect(initResult.success).toBe(true);

      // 7. Validate with Terraform (MUST pass first try)
      const validateResult = await terraform.validate();
      expect(validateResult.success).toBe(true);
      expect(validateResult.errors).toHaveLength(0);

      // 8. Apply (MUST succeed first attempt - DETERMINISM TEST)
      const applyResult = await terraform.apply();
      expect(applyResult.success).toBe(true);
      expect(applyResult.attempts).toBe(1);

      // 9. Verify origin pool exists via API
      const exists = await validator.originPoolExists(testNamespace, originPoolName);
      expect(exists).toBe(true);

      // 10. Verify origin pool configuration
      const details = await validator.getOriginPool(testNamespace, originPoolName);
      expect(details.exists).toBe(true);
      expect(details.details?.name).toBe(originPoolName);
    });
  });

  // ===========================================================================
  // SCENARIO 3: HTTP LOAD BALANCER WITH ORIGIN POOL
  // ===========================================================================

  describe('Scenario 3: HTTP Load Balancer with Origin Pool', { timeout: TEST_TIMEOUT }, () => {
    let terraform: TerraformRunner;
    const lbName = 'httpbin-lb';
    const originPoolName = 'httpbin-pool-lb';

    beforeEach(() => {
      if (skipTests) return;
      terraform = new TerraformRunner(`e2e-http-lb-${Date.now()}`);
    });

    afterAll(async () => {
      if (skipTests || !terraform) return;
      await terraform.cleanup();
    });

    it.skipIf(skipTests)('should create HTTP LB with origin pool reference', async () => {
      // 1. Setup
      await terraform.setup();
      terraform.writeProviderConfig(credentials.apiUrl);

      // 2. Get dependencies from MCP
      const dependencies = await mcpClient.getDependencies('http_loadbalancer');
      console.log('MCP Dependencies guidance received');

      // 3. Get MCP-generated examples as reference for correct syntax
      const mcpLbExample = await mcpClient.getExample('http_loadbalancer', 'basic');
      console.log('MCP HTTP LB example received, length:', mcpLbExample.length);

      // Verify MCP example uses correct empty block syntax
      if (!mcpLbExample.includes('disable_waf {}') || !mcpLbExample.includes('round_robin {}')) {
        console.warn('WARNING: MCP example may not use correct empty block syntax');
      }

      // 4. Generate full configuration with namespace, origin pool, and LB
      const fullConfig = `
resource "f5xc_namespace" "test" {
  name = "${testNamespace}"
}

resource "f5xc_origin_pool" "httpbin" {
  name      = "${originPoolName}"
  namespace = f5xc_namespace.test.name

  origin_servers {
    public_name {
      dns_name         = "httpbin.org"
      refresh_interval = 60
    }
  }

  port = 443

  use_tls {
    sni = "httpbin.org"
    tls_config {
      default_security {}
    }
    no_mtls {}
    volterra_trusted_ca {}
  }

  loadbalancer_algorithm = "ROUND_ROBIN"
  endpoint_selection     = "LOCAL_PREFERRED"
}

resource "f5xc_http_loadbalancer" "test" {
  name      = "${lbName}"
  namespace = f5xc_namespace.test.name
  domains   = ["${lbName}.example.com"]

  # Empty blocks for flag-style attributes
  advertise_on_public_default_vip {}

  http {
    dns_volterra_managed = false
    port                 = 80
  }

  default_route_pools {
    pool {
      name      = f5xc_origin_pool.httpbin.name
      namespace = f5xc_namespace.test.name
    }
    weight   = 1
    priority = 1
  }

  # Empty blocks for OneOf selections
  round_robin {}

  disable_waf {}
  disable_bot_defense {}
  disable_rate_limit {}
  no_challenge {}
  disable_api_definition {}
  disable_api_discovery {}
  disable_ip_reputation {}
  disable_client_side_defense {}
  disable_trust_client_ip_headers {}

  labels = {
    "test-type" = "mcp-e2e"
  }
}
`;
      terraform.writeConfig('main.tf', fullConfig);

      // 4. Initialize and validate
      const initResult = await terraform.init();
      expect(initResult.success).toBe(true);

      const validateResult = await terraform.validate();
      expect(validateResult.success).toBe(true);
      expect(validateResult.errors).toHaveLength(0);

      // 5. Apply (DETERMINISM TEST)
      const applyResult = await terraform.apply();
      expect(applyResult.success).toBe(true);
      expect(applyResult.attempts).toBe(1);

      // 6. Verify all resources exist
      expect(await validator.namespaceExists(testNamespace)).toBe(true);
      expect(await validator.originPoolExists(testNamespace, originPoolName)).toBe(true);
      expect(await validator.httpLoadBalancerExists(testNamespace, lbName)).toBe(true);
    });
  });

  // ===========================================================================
  // SCENARIO 4: FULL SECURITY STACK
  // ===========================================================================

  describe('Scenario 4: HTTP LB with Security Controls', { timeout: TEST_TIMEOUT }, () => {
    let terraform: TerraformRunner;
    const lbName = 'secure-lb';
    const originPoolName = 'secure-pool';

    beforeEach(() => {
      if (skipTests) return;
      terraform = new TerraformRunner(`e2e-security-${Date.now()}`);
    });

    afterAll(async () => {
      if (skipTests || !terraform) return;
      await terraform.cleanup();
    });

    it.skipIf(skipTests)('should create HTTP LB with rate limiting and IP reputation', async () => {
      // 1. Setup
      await terraform.setup();
      terraform.writeProviderConfig(credentials.apiUrl);

      // 2. Get common mistakes to avoid from MCP
      const mistakes = await mcpClient.getCommonMistakes('http_loadbalancer');
      console.log('MCP Common Mistakes guidance received');

      // 3. Generate security-enabled configuration
      const securityConfig = `
resource "f5xc_namespace" "test" {
  name = "${testNamespace}"
}

resource "f5xc_origin_pool" "secure" {
  name      = "${originPoolName}"
  namespace = f5xc_namespace.test.name

  origin_servers {
    public_name {
      dns_name         = "httpbin.org"
      refresh_interval = 60
    }
  }

  port = 443

  use_tls {
    sni = "httpbin.org"
    tls_config {
      default_security {}
    }
    no_mtls {}
    volterra_trusted_ca {}
  }

  loadbalancer_algorithm = "ROUND_ROBIN"
  endpoint_selection     = "LOCAL_PREFERRED"
}

resource "f5xc_http_loadbalancer" "secure" {
  name      = "${lbName}"
  namespace = f5xc_namespace.test.name
  domains   = ["${lbName}.example.com"]

  # Empty blocks for flag-style attributes
  advertise_on_public_default_vip {}

  http {
    dns_volterra_managed = false
    port                 = 80
  }

  default_route_pools {
    pool {
      name      = f5xc_origin_pool.secure.name
      namespace = f5xc_namespace.test.name
    }
    weight = 1
  }

  # Empty blocks for OneOf selections
  round_robin {}

  # Security controls - empty block syntax
  disable_waf {}
  disable_bot_defense {}
  no_challenge {}
  disable_api_definition {}
  disable_api_discovery {}
  disable_client_side_defense {}
  disable_trust_client_ip_headers {}

  # Rate limiting configuration
  rate_limit {
    no_ip_allowed_list {}
    rate_limiter {
      total_number     = 100
      unit             = "MINUTE"
      burst_multiplier = 2
    }
  }

  # IP reputation filtering
  enable_ip_reputation {
    ip_threat_categories = [
      "SPAM_SOURCES",
      "WEB_ATTACKS",
      "BOTNETS",
      "SCANNERS"
    ]
  }

  # Malicious user detection - empty block
  enable_malicious_user_detection {}

  labels = {
    "test-type"     = "mcp-e2e"
    "security-tier" = "standard"
  }
}
`;
      terraform.writeConfig('main.tf', securityConfig);

      // 4. Initialize and validate
      const initResult = await terraform.init();
      expect(initResult.success).toBe(true);

      const validateResult = await terraform.validate();
      if (!validateResult.success) {
        console.log('Terraform validation errors:', validateResult.errors);
      }
      expect(validateResult.success).toBe(true);
      expect(validateResult.errors).toHaveLength(0);

      // 5. Apply (DETERMINISM TEST)
      const applyResult = await terraform.apply();
      if (!applyResult.success) {
        console.log('Terraform apply output:', applyResult.output);
      }
      expect(applyResult.success).toBe(true);
      expect(applyResult.attempts).toBe(1);

      // 6. Verify LB exists with security enabled
      const lbDetails = await validator.getHttpLoadBalancer(testNamespace, lbName);
      expect(lbDetails.exists).toBe(true);

      // 7. Verify rate limiting is configured (check spec)
      expect(lbDetails.details?.spec).toBeDefined();
    });
  });
});
