// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

/**
 * Resource Validator Utility
 *
 * Validates that F5XC resources exist via direct API calls.
 * Used in E2E tests to verify Terraform apply created resources correctly.
 */

import axios, { AxiosInstance, AxiosError } from 'axios';

export interface ResourceDetails {
  name: string;
  namespace?: string;
  metadata?: Record<string, unknown>;
  spec?: Record<string, unknown>;
  system_metadata?: {
    uid?: string;
    creation_timestamp?: string;
    modification_timestamp?: string;
  };
}

export interface ValidationResult {
  exists: boolean;
  details?: ResourceDetails;
  error?: string;
}

/**
 * Resource Validator for F5XC API
 *
 * Makes direct API calls to verify resources exist and have expected configuration.
 */
export class ResourceValidator {
  private client: AxiosInstance;
  private apiUrl: string;

  constructor(apiUrl: string, apiToken: string) {
    this.apiUrl = apiUrl.replace(/\/$/, ''); // Remove trailing slash

    this.client = axios.create({
      baseURL: this.apiUrl,
      headers: {
        'Authorization': `APIToken ${apiToken}`,
        'Content-Type': 'application/json',
        'Accept': 'application/json',
      },
      timeout: 30000, // 30 second timeout
    });
  }

  // ===========================================================================
  // NAMESPACE OPERATIONS
  // ===========================================================================

  /**
   * Check if a namespace exists
   */
  async namespaceExists(name: string): Promise<boolean> {
    const result = await this.getNamespace(name);
    return result.exists;
  }

  /**
   * Get namespace details
   */
  async getNamespace(name: string): Promise<ValidationResult> {
    try {
      const response = await this.client.get(`/api/web/namespaces/${name}`);
      return {
        exists: true,
        details: this.extractResourceDetails(response.data),
      };
    } catch (error) {
      return this.handleApiError(error, 'namespace', name);
    }
  }

  /**
   * Delete a namespace (for cleanup)
   */
  async deleteNamespace(name: string): Promise<boolean> {
    try {
      await this.client.delete(`/api/web/namespaces/${name}`);
      return true;
    } catch {
      return false;
    }
  }

  // ===========================================================================
  // ORIGIN POOL OPERATIONS
  // ===========================================================================

  /**
   * Check if an origin pool exists
   */
  async originPoolExists(namespace: string, name: string): Promise<boolean> {
    const result = await this.getOriginPool(namespace, name);
    return result.exists;
  }

  /**
   * Get origin pool details
   */
  async getOriginPool(namespace: string, name: string): Promise<ValidationResult> {
    try {
      const response = await this.client.get(
        `/api/config/namespaces/${namespace}/origin_pools/${name}`,
      );
      return {
        exists: true,
        details: this.extractResourceDetails(response.data),
      };
    } catch (error) {
      return this.handleApiError(error, 'origin_pool', `${namespace}/${name}`);
    }
  }

  /**
   * Delete an origin pool (for cleanup)
   */
  async deleteOriginPool(namespace: string, name: string): Promise<boolean> {
    try {
      await this.client.delete(
        `/api/config/namespaces/${namespace}/origin_pools/${name}`,
      );
      return true;
    } catch {
      return false;
    }
  }

  // ===========================================================================
  // HTTP LOAD BALANCER OPERATIONS
  // ===========================================================================

  /**
   * Check if an HTTP load balancer exists
   */
  async httpLoadBalancerExists(namespace: string, name: string): Promise<boolean> {
    const result = await this.getHttpLoadBalancer(namespace, name);
    return result.exists;
  }

  /**
   * Get HTTP load balancer details
   */
  async getHttpLoadBalancer(namespace: string, name: string): Promise<ValidationResult> {
    try {
      const response = await this.client.get(
        `/api/config/namespaces/${namespace}/http_loadbalancers/${name}`,
      );
      return {
        exists: true,
        details: this.extractResourceDetails(response.data),
      };
    } catch (error) {
      return this.handleApiError(error, 'http_loadbalancer', `${namespace}/${name}`);
    }
  }

  /**
   * Delete an HTTP load balancer (for cleanup)
   */
  async deleteHttpLoadBalancer(namespace: string, name: string): Promise<boolean> {
    try {
      await this.client.delete(
        `/api/config/namespaces/${namespace}/http_loadbalancers/${name}`,
      );
      return true;
    } catch {
      return false;
    }
  }

  // ===========================================================================
  // APP FIREWALL (WAF) OPERATIONS
  // ===========================================================================

  /**
   * Check if an app firewall exists
   */
  async appFirewallExists(namespace: string, name: string): Promise<boolean> {
    const result = await this.getAppFirewall(namespace, name);
    return result.exists;
  }

  /**
   * Get app firewall details
   */
  async getAppFirewall(namespace: string, name: string): Promise<ValidationResult> {
    try {
      const response = await this.client.get(
        `/api/config/namespaces/${namespace}/app_firewalls/${name}`,
      );
      return {
        exists: true,
        details: this.extractResourceDetails(response.data),
      };
    } catch (error) {
      return this.handleApiError(error, 'app_firewall', `${namespace}/${name}`);
    }
  }

  /**
   * Delete an app firewall (for cleanup)
   */
  async deleteAppFirewall(namespace: string, name: string): Promise<boolean> {
    try {
      await this.client.delete(
        `/api/config/namespaces/${namespace}/app_firewalls/${name}`,
      );
      return true;
    } catch {
      return false;
    }
  }

  // ===========================================================================
  // HEALTHCHECK OPERATIONS
  // ===========================================================================

  /**
   * Check if a healthcheck exists
   */
  async healthcheckExists(namespace: string, name: string): Promise<boolean> {
    const result = await this.getHealthcheck(namespace, name);
    return result.exists;
  }

  /**
   * Get healthcheck details
   */
  async getHealthcheck(namespace: string, name: string): Promise<ValidationResult> {
    try {
      const response = await this.client.get(
        `/api/config/namespaces/${namespace}/healthchecks/${name}`,
      );
      return {
        exists: true,
        details: this.extractResourceDetails(response.data),
      };
    } catch (error) {
      return this.handleApiError(error, 'healthcheck', `${namespace}/${name}`);
    }
  }

  /**
   * Delete a healthcheck (for cleanup)
   */
  async deleteHealthcheck(namespace: string, name: string): Promise<boolean> {
    try {
      await this.client.delete(
        `/api/config/namespaces/${namespace}/healthchecks/${name}`,
      );
      return true;
    } catch {
      return false;
    }
  }

  // ===========================================================================
  // GENERIC RESOURCE OPERATIONS
  // ===========================================================================

  /**
   * Get any resource by type, namespace, and name
   */
  async getResource(
    resourceType: string,
    namespace: string,
    name: string,
  ): Promise<ValidationResult> {
    // Map resource types to API paths
    const pathMap: Record<string, string> = {
      namespace: `/api/web/namespaces/${name}`,
      origin_pool: `/api/config/namespaces/${namespace}/origin_pools/${name}`,
      http_loadbalancer: `/api/config/namespaces/${namespace}/http_loadbalancers/${name}`,
      app_firewall: `/api/config/namespaces/${namespace}/app_firewalls/${name}`,
      healthcheck: `/api/config/namespaces/${namespace}/healthchecks/${name}`,
      rate_limiter: `/api/config/namespaces/${namespace}/rate_limiters/${name}`,
      service_policy: `/api/config/namespaces/${namespace}/service_policys/${name}`,
    };

    const path = pathMap[resourceType];
    if (!path) {
      return {
        exists: false,
        error: `Unknown resource type: ${resourceType}`,
      };
    }

    try {
      const response = await this.client.get(path);
      return {
        exists: true,
        details: this.extractResourceDetails(response.data),
      };
    } catch (error) {
      return this.handleApiError(error, resourceType, `${namespace}/${name}`);
    }
  }

  /**
   * Check if any resource exists
   */
  async resourceExists(
    resourceType: string,
    namespace: string,
    name: string,
  ): Promise<boolean> {
    const result = await this.getResource(resourceType, namespace, name);
    return result.exists;
  }

  // ===========================================================================
  // VALIDATION HELPERS
  // ===========================================================================

  /**
   * Validate resource has expected spec values
   */
  validateSpec(
    details: ResourceDetails,
    expectedSpec: Record<string, unknown>,
  ): { valid: boolean; mismatches: string[] } {
    const mismatches: string[] = [];

    for (const [key, expectedValue] of Object.entries(expectedSpec)) {
      const actualValue = details.spec?.[key];

      if (JSON.stringify(actualValue) !== JSON.stringify(expectedValue)) {
        mismatches.push(
          `${key}: expected ${JSON.stringify(expectedValue)}, got ${JSON.stringify(actualValue)}`,
        );
      }
    }

    return {
      valid: mismatches.length === 0,
      mismatches,
    };
  }

  /**
   * Wait for a resource to be ready (with timeout)
   */
  async waitForResource(
    resourceType: string,
    namespace: string,
    name: string,
    timeoutMs: number = 60000,
    intervalMs: number = 2000,
  ): Promise<ValidationResult> {
    const startTime = Date.now();

    while (Date.now() - startTime < timeoutMs) {
      const result = await this.getResource(resourceType, namespace, name);
      if (result.exists) {
        return result;
      }

      await this.delay(intervalMs);
    }

    return {
      exists: false,
      error: `Timeout waiting for ${resourceType} ${namespace}/${name}`,
    };
  }

  // ===========================================================================
  // CLEANUP HELPERS
  // ===========================================================================

  /**
   * Clean up all test resources in a namespace
   */
  async cleanupNamespace(
    namespace: string,
    options: {
      deleteNamespace?: boolean;
      resourceTypes?: string[];
    } = {},
  ): Promise<{ deleted: string[]; errors: string[] }> {
    const deleted: string[] = [];
    const errors: string[] = [];

    const resourceTypes = options.resourceTypes || [
      'http_loadbalancer',
      'origin_pool',
      'app_firewall',
      'healthcheck',
    ];

    // Delete resources in reverse dependency order
    for (const resourceType of resourceTypes) {
      try {
        const listPath = this.getListPath(resourceType, namespace);
        if (!listPath) continue;

        const response = await this.client.get(listPath);
        const items = response.data?.items || [];

        for (const item of items) {
          const name = item?.metadata?.name || item?.name;
          if (name) {
            try {
              await this.deleteResource(resourceType, namespace, name);
              deleted.push(`${resourceType}/${name}`);
            } catch (e) {
              errors.push(`Failed to delete ${resourceType}/${name}: ${e}`);
            }
          }
        }
      } catch {
        // Ignore list errors
      }
    }

    // Delete namespace if requested
    if (options.deleteNamespace) {
      try {
        await this.deleteNamespace(namespace);
        deleted.push(`namespace/${namespace}`);
      } catch (e) {
        errors.push(`Failed to delete namespace: ${e}`);
      }
    }

    return { deleted, errors };
  }

  /**
   * Delete any resource by type
   */
  private async deleteResource(
    resourceType: string,
    namespace: string,
    name: string,
  ): Promise<boolean> {
    const deleteMap: Record<string, () => Promise<boolean>> = {
      namespace: () => this.deleteNamespace(name),
      origin_pool: () => this.deleteOriginPool(namespace, name),
      http_loadbalancer: () => this.deleteHttpLoadBalancer(namespace, name),
      app_firewall: () => this.deleteAppFirewall(namespace, name),
      healthcheck: () => this.deleteHealthcheck(namespace, name),
    };

    const deleteFn = deleteMap[resourceType];
    if (!deleteFn) {
      return false;
    }

    return deleteFn();
  }

  /**
   * Get list path for a resource type
   */
  private getListPath(resourceType: string, namespace: string): string | null {
    const pathMap: Record<string, string> = {
      origin_pool: `/api/config/namespaces/${namespace}/origin_pools`,
      http_loadbalancer: `/api/config/namespaces/${namespace}/http_loadbalancers`,
      app_firewall: `/api/config/namespaces/${namespace}/app_firewalls`,
      healthcheck: `/api/config/namespaces/${namespace}/healthchecks`,
    };

    return pathMap[resourceType] || null;
  }

  // ===========================================================================
  // PRIVATE HELPERS
  // ===========================================================================

  /**
   * Extract resource details from API response
   */
  private extractResourceDetails(data: Record<string, unknown>): ResourceDetails {
    return {
      name: (data.metadata as Record<string, unknown>)?.name as string || data.name as string || '',
      namespace: (data.metadata as Record<string, unknown>)?.namespace as string,
      metadata: data.metadata as Record<string, unknown>,
      spec: data.spec as Record<string, unknown>,
      system_metadata: data.system_metadata as ResourceDetails['system_metadata'],
    };
  }

  /**
   * Handle API errors
   */
  private handleApiError(
    error: unknown,
    resourceType: string,
    identifier: string,
  ): ValidationResult {
    if (axios.isAxiosError(error)) {
      const axiosError = error as AxiosError;

      if (axiosError.response?.status === 404) {
        return {
          exists: false,
          error: `${resourceType} '${identifier}' not found`,
        };
      }

      return {
        exists: false,
        error: `API error for ${resourceType} '${identifier}': ${axiosError.message}`,
      };
    }

    return {
      exists: false,
      error: `Unknown error checking ${resourceType} '${identifier}': ${error}`,
    };
  }

  /**
   * Delay helper
   */
  private delay(ms: number): Promise<void> {
    return new Promise((resolve) => setTimeout(resolve, ms));
  }
}

/**
 * Factory function for creating resource validators
 */
export function createResourceValidator(apiUrl: string, apiToken: string): ResourceValidator {
  return new ResourceValidator(apiUrl, apiToken);
}
