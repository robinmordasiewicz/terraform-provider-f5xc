// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

/**
 * Operation Metadata Service for F5 XC MCP Server
 *
 * Provides access to operation-level metadata from v2.0.33 API specifications.
 * This includes danger levels, side effects, response times, and best practices.
 */

import * as fs from 'fs';
import * as path from 'path';
import {
  OperationsMetadataCollection,
  ResourceOperationInfo,
  OperationMetadata,
  DangerLevel,
  SideEffectsInfo,
  BestPracticesInfo,
  GuidedWorkflowInfo,
  ResponseTimeInfo,
} from '../types.js';

export class OperationsService {
  private operationsMetadata: OperationsMetadataCollection | null = null;
  private metadataPath: string;
  private loaded = false;

  constructor(basePath: string = process.cwd()) {
    // Try multiple paths to find operations-metadata.json
    const possiblePaths = [
      path.join(basePath, 'tools', 'metadata', 'operations-metadata.json'),
      path.join(basePath, '..', 'tools', 'metadata', 'operations-metadata.json'),
      path.join(__dirname, '..', '..', '..', 'tools', 'metadata', 'operations-metadata.json'),
    ];

    this.metadataPath = possiblePaths[0];
    for (const p of possiblePaths) {
      if (fs.existsSync(p)) {
        this.metadataPath = p;
        break;
      }
    }
  }

  /**
   * Load operations metadata from JSON file
   */
  private async loadMetadata(): Promise<void> {
    if (this.loaded) return;

    try {
      if (fs.existsSync(this.metadataPath)) {
        const content = fs.readFileSync(this.metadataPath, 'utf-8');
        this.operationsMetadata = JSON.parse(content);
        this.loaded = true;
      } else {
        // Initialize with empty collection if file doesn't exist
        this.operationsMetadata = {
          generated_at: new Date().toISOString(),
          version: '1.0.0',
          resources: {},
        };
        this.loaded = true;
      }
    } catch (error) {
      console.error('Failed to load operations metadata:', error);
      this.operationsMetadata = {
        generated_at: new Date().toISOString(),
        version: '1.0.0',
        resources: {},
      };
      this.loaded = true;
    }
  }

  /**
   * Get all operation metadata for a resource
   */
  async getResourceOperations(resource: string): Promise<ResourceOperationInfo | null> {
    await this.loadMetadata();
    return this.operationsMetadata?.resources[resource] || null;
  }

  /**
   * Get danger level for a specific operation
   */
  async getDangerLevel(
    resource: string,
    operation: 'create' | 'read' | 'update' | 'delete' | 'list',
  ): Promise<DangerLevel | undefined> {
    const resourceOps = await this.getResourceOperations(resource);
    return resourceOps?.operations[operation]?.danger_level;
  }

  /**
   * Get side effects for a specific operation
   */
  async getSideEffects(
    resource: string,
    operation: 'create' | 'read' | 'update' | 'delete' | 'list',
  ): Promise<SideEffectsInfo | undefined> {
    const resourceOps = await this.getResourceOperations(resource);
    return resourceOps?.operations[operation]?.side_effects;
  }

  /**
   * Get response time metrics for a specific operation
   */
  async getResponseTime(
    resource: string,
    operation: 'create' | 'read' | 'update' | 'delete' | 'list',
  ): Promise<ResponseTimeInfo | undefined> {
    const resourceOps = await this.getResourceOperations(resource);
    return resourceOps?.operations[operation]?.discovered_response_time;
  }

  /**
   * Check if an operation requires confirmation
   */
  async requiresConfirmation(
    resource: string,
    operation: 'create' | 'read' | 'update' | 'delete' | 'list',
  ): Promise<boolean> {
    const resourceOps = await this.getResourceOperations(resource);
    return resourceOps?.operations[operation]?.confirmation_required || false;
  }

  /**
   * Get required fields for an operation
   */
  async getRequiredFields(
    resource: string,
    operation: 'create' | 'read' | 'update' | 'delete' | 'list',
  ): Promise<string[]> {
    const resourceOps = await this.getResourceOperations(resource);
    return resourceOps?.operations[operation]?.required_fields || [];
  }

  /**
   * Get best practices for a resource
   */
  async getBestPractices(resource: string): Promise<BestPracticesInfo | undefined> {
    const resourceOps = await this.getResourceOperations(resource);
    return resourceOps?.best_practices;
  }

  /**
   * Get guided workflows for a resource
   */
  async getGuidedWorkflows(resource: string): Promise<GuidedWorkflowInfo[]> {
    const resourceOps = await this.getResourceOperations(resource);
    return resourceOps?.guided_workflows || [];
  }

  /**
   * Get complete operation metadata for a specific operation
   */
  async getOperationMetadata(
    resource: string,
    operation: 'create' | 'read' | 'update' | 'delete' | 'list',
  ): Promise<OperationMetadata | undefined> {
    const resourceOps = await this.getResourceOperations(resource);
    return resourceOps?.operations[operation];
  }

  /**
   * Get all high-danger operations across all resources
   */
  async getHighDangerOperations(): Promise<
    Array<{ resource: string; operation: string; metadata: OperationMetadata }>
  > {
    await this.loadMetadata();
    const results: Array<{
      resource: string;
      operation: string;
      metadata: OperationMetadata;
    }> = [];

    if (!this.operationsMetadata) return results;

    for (const [resourceName, resourceOps] of Object.entries(
      this.operationsMetadata.resources,
    )) {
      for (const [opName, opMeta] of Object.entries(resourceOps.operations)) {
        if (opMeta.danger_level === 'high' || opMeta.danger_level === 'critical') {
          results.push({
            resource: resourceName,
            operation: opName,
            metadata: opMeta,
          });
        }
      }
    }

    return results;
  }

  /**
   * Get all operations requiring confirmation
   */
  async getConfirmationRequiredOperations(): Promise<
    Array<{ resource: string; operation: string; metadata: OperationMetadata }>
  > {
    await this.loadMetadata();
    const results: Array<{
      resource: string;
      operation: string;
      metadata: OperationMetadata;
    }> = [];

    if (!this.operationsMetadata) return results;

    for (const [resourceName, resourceOps] of Object.entries(
      this.operationsMetadata.resources,
    )) {
      for (const [opName, opMeta] of Object.entries(resourceOps.operations)) {
        if (opMeta.confirmation_required) {
          results.push({
            resource: resourceName,
            operation: opName,
            metadata: opMeta,
          });
        }
      }
    }

    return results;
  }

  /**
   * Get all resources with operation metadata
   */
  async getResourcesWithOperationMetadata(): Promise<string[]> {
    await this.loadMetadata();
    return Object.keys(this.operationsMetadata?.resources || {});
  }

  /**
   * Get summary statistics about operation metadata
   */
  async getSummary(): Promise<{
    totalResources: number;
    totalOperations: number;
    byDangerLevel: Record<string, number>;
    confirmationRequired: number;
  }> {
    await this.loadMetadata();

    const summary = {
      totalResources: 0,
      totalOperations: 0,
      byDangerLevel: {} as Record<string, number>,
      confirmationRequired: 0,
    };

    if (!this.operationsMetadata) return summary;

    summary.totalResources = Object.keys(this.operationsMetadata.resources).length;

    for (const resourceOps of Object.values(this.operationsMetadata.resources)) {
      for (const opMeta of Object.values(resourceOps.operations)) {
        summary.totalOperations++;
        if (opMeta.danger_level) {
          summary.byDangerLevel[opMeta.danger_level] =
            (summary.byDangerLevel[opMeta.danger_level] || 0) + 1;
        }
        if (opMeta.confirmation_required) {
          summary.confirmationRequired++;
        }
      }
    }

    return summary;
  }
}
