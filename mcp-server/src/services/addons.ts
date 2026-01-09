/**
 * Addon Service Module
 * Provides information about F5 Distributed Cloud addon services and activation workflows
 */

import { readFileSync as _readFileSync, existsSync as _existsSync } from 'fs';
import { join, dirname } from 'path';
import { fileURLToPath } from 'url';

// Addon service information from subscription metadata
const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

const PACKAGE_ROOT = join(__dirname, '..', '..'); // mcp-server/
const PROJECT_ROOT = join(PACKAGE_ROOT, '..'); // terraform-provider-f5xc/

// Paths for subscription metadata (contains addon service info)
const _BUNDLED_SUBSCRIPTION_METADATA = join(PACKAGE_ROOT, 'dist', 'subscription-tiers.json');
const _PROJECT_SUBSCRIPTION_METADATA = join(PROJECT_ROOT, 'tools', 'subscription-tiers.json');

// Known addon services with their activation types and descriptions
// These names are the actual API names from F5XC (validated via API testing)
// Format: f5xc-{service}-{tier}
const ADDON_SERVICES: Record<string, AddonServiceInfo> = {
  // Bot Defense variants
  'f5xc-bot-defense-standard': {
    name: 'f5xc-bot-defense-standard',
    displayName: 'Bot Defense (Standard)',
    description: 'Protect applications from automated attacks and bot traffic - Standard tier',
    tier: 'STANDARD',
    activationType: 'self',
    category: 'security',
  },
  'f5xc-bot-defense-advanced': {
    name: 'f5xc-bot-defense-advanced',
    displayName: 'Bot Defense (Advanced)',
    description: 'Protect applications from automated attacks and bot traffic - Advanced tier with ML',
    tier: 'ADVANCED',
    activationType: 'self',
    category: 'security',
  },
  // Client Side Defense
  'f5xc-client-side-defense-standard': {
    name: 'f5xc-client-side-defense-standard',
    displayName: 'Client Side Defense',
    description: 'Protect against client-side attacks like Magecart and formjacking',
    tier: 'STANDARD',
    activationType: 'self',
    category: 'security',
  },
  // WAAP (Web App and API Protection) - includes API Discovery, API Protection, Data Guard
  'f5xc-waap-standard': {
    name: 'f5xc-waap-standard',
    displayName: 'WAAP (Standard)',
    description: 'Web Application and API Protection - includes API Discovery and Data Guard',
    tier: 'STANDARD',
    activationType: 'self',
    category: 'security',
  },
  'f5xc-waap-advanced': {
    name: 'f5xc-waap-advanced',
    displayName: 'WAAP (Advanced)',
    description: 'Web Application and API Protection - Advanced tier with full API security',
    tier: 'ADVANCED',
    activationType: 'self',
    category: 'security',
  },
  // Malicious User Detection
  'f5xc-malicious-user-detection': {
    name: 'f5xc-malicious-user-detection',
    displayName: 'Malicious User Detection',
    description: 'Identify and block malicious user behavior patterns',
    tier: 'ADVANCED',
    activationType: 'self',
    category: 'security',
  },
  // Synthetic Monitoring
  'f5xc-synthetic-monitoring': {
    name: 'f5xc-synthetic-monitoring',
    displayName: 'Synthetic Monitoring',
    description: 'Monitor application availability and performance with synthetic tests',
    tier: 'STANDARD',
    activationType: 'self',
    category: 'observability',
  },
};

interface AddonServiceInfo {
  name: string;
  displayName: string;
  description: string;
  tier: 'NO_TIER' | 'STANDARD' | 'ADVANCED' | 'PREMIUM';
  activationType: 'self' | 'partial' | 'managed';
  category: string;
}

interface AddonListResult {
  total: number;
  services: AddonServiceInfo[];
}

interface AddonActivationCheck {
  addonService: string;
  displayName: string;
  tier: string;
  activationType: string;
  canActivate: boolean;
  steps: string[];
  terraformExample: string;
}

interface AddonWorkflow {
  addonService: string;
  activationType: string;
  description: string;
  prerequisites: string[];
  steps: WorkflowStep[];
  terraformConfig: string;
  estimatedTime: string;
  notes: string[];
}

interface WorkflowStep {
  step: number;
  action: string;
  description: string;
  terraformSnippet?: string;
}

/**
 * Convert addon service name to valid Terraform resource name
 * e.g., "f5xc-bot-defense-standard" -> "bot_defense_standard"
 */
function toTerraformResourceName(addonServiceName: string): string {
  return addonServiceName
    .replace(/^f5xc-/, '')  // Remove f5xc- prefix
    .replace(/-/g, '_');    // Replace hyphens with underscores
}

/**
 * List all available addon services with optional filtering
 */
export function listAddonServices(
  tierFilter?: 'STANDARD' | 'ADVANCED' | 'PREMIUM',
  activationTypeFilter?: 'self' | 'managed'
): AddonListResult {
  let services = Object.values(ADDON_SERVICES);

  // Apply tier filter
  if (tierFilter) {
    services = services.filter(s => s.tier === tierFilter);
  }

  // Apply activation type filter
  if (activationTypeFilter) {
    if (activationTypeFilter === 'managed') {
      services = services.filter(s => s.activationType === 'managed' || s.activationType === 'partial');
    } else {
      services = services.filter(s => s.activationType === activationTypeFilter);
    }
  }

  return {
    total: services.length,
    services,
  };
}

/**
 * Check activation requirements for a specific addon service
 */
export function checkAddonActivation(addonService: string): AddonActivationCheck | null {
  const service = ADDON_SERVICES[addonService];

  if (!service) {
    return null;
  }

  // Build activation steps based on type
  let steps: string[];
  let canActivate = true;

  switch (service.activationType) {
    case 'self':
      steps = [
        'Check activation status using data source',
        'Create addon_subscription resource',
        'Wait for subscription to be enabled',
        'Use addon features in your configuration',
      ];
      break;
    case 'partial':
      steps = [
        'Check activation status using data source',
        'Create addon_subscription with notification preference',
        'Wait for SRE team to process the request',
        'Monitor subscription state for SUBSCRIPTION_ENABLED',
        'Use addon features once enabled',
      ];
      canActivate = true; // User can initiate, but needs approval
      break;
    case 'managed':
      steps = [
        'Contact F5 Sales to discuss addon requirements',
        'Complete the subscription agreement',
        'Wait for F5 team to activate the addon',
        'Verify activation using data source',
        'Use addon features once enabled',
      ];
      canActivate = false; // Requires sales contact
      break;
    default:
      steps = ['Unknown activation type'];
  }

  const terraformResourceName = toTerraformResourceName(addonService);
  const terraformExample = `# Check if ${service.displayName} can be activated
data "f5xc_addon_service_activation_status" "${terraformResourceName}" {
  addon_service = "${addonService}"
}

# Create subscription if available
resource "f5xc_addon_subscription" "${terraformResourceName}" {
  count     = data.f5xc_addon_service_activation_status.${terraformResourceName}.can_activate ? 1 : 0
  name      = "${terraformResourceName.replace(/_/g, '-')}-subscription"
  namespace = "system"

  addon_service {
    name      = "${addonService}"
    namespace = "shared"
  }
}`;

  return {
    addonService: service.name,
    displayName: service.displayName,
    tier: service.tier,
    activationType: service.activationType,
    canActivate,
    steps,
    terraformExample,
  };
}

/**
 * Get detailed activation workflow for an addon service
 */
export function getAddonWorkflow(
  addonService: string,
  activationType?: 'self' | 'partial' | 'managed'
): AddonWorkflow | null {
  const service = ADDON_SERVICES[addonService];

  if (!service) {
    return null;
  }

  // Use provided activation type or default to the service's type
  const effectiveType = activationType || service.activationType;
  const tfResourceName = toTerraformResourceName(addonService);

  let workflow: AddonWorkflow;

  switch (effectiveType) {
    case 'self':
      workflow = {
        addonService: service.name,
        activationType: 'self',
        description: `Self-activation workflow for ${service.displayName}. User can activate directly without manual intervention.`,
        prerequisites: [
          'Valid F5 Distributed Cloud tenant',
          `${service.tier} subscription tier or higher`,
          'Appropriate IAM permissions for addon management',
        ],
        steps: [
          {
            step: 1,
            action: 'Check Activation Status',
            description: 'Verify the addon can be activated for your tenant',
            terraformSnippet: `data "f5xc_addon_service_activation_status" "${tfResourceName}" {
  addon_service = "${addonService}"
}

output "${tfResourceName}_can_activate" {
  value = data.f5xc_addon_service_activation_status.${tfResourceName}.can_activate
}`,
          },
          {
            step: 2,
            action: 'Create Subscription',
            description: 'Create the addon subscription resource',
            terraformSnippet: `resource "f5xc_addon_subscription" "${tfResourceName}" {
  name      = "${tfResourceName.replace(/_/g, '-')}-subscription"
  namespace = "system"

  addon_service {
    name      = "${addonService}"
    namespace = "shared"
  }
}`,
          },
          {
            step: 3,
            action: 'Verify Activation',
            description: 'Confirm the subscription is enabled',
            terraformSnippet: `# After apply, the subscription should be in SUBSCRIPTION_ENABLED state
# Check status:
# terraform show | grep status`,
          },
          {
            step: 4,
            action: 'Use Addon Features',
            description: `Enable ${service.displayName} features in your load balancer or other resources`,
          },
        ],
        terraformConfig: generateFullTerraformConfig(service),
        estimatedTime: '1-5 minutes',
        notes: [
          'Self-activation is typically immediate',
          'Some features may require additional configuration',
          `Ensure your subscription tier is ${service.tier} or higher`,
        ],
      };
      break;

    case 'partial':
      workflow = {
        addonService: service.name,
        activationType: 'partial',
        description: `Partially managed activation for ${service.displayName}. User initiates request, requires SRE processing.`,
        prerequisites: [
          'Valid F5 Distributed Cloud tenant',
          `${service.tier} subscription tier or higher`,
          'Appropriate IAM permissions',
          'Email for notification preferences',
        ],
        steps: [
          {
            step: 1,
            action: 'Check Activation Status',
            description: 'Verify the addon can be requested',
            terraformSnippet: `data "f5xc_addon_service_activation_status" "${tfResourceName}" {
  addon_service = "${addonService}"
}`,
          },
          {
            step: 2,
            action: 'Create Subscription Request',
            description: 'Submit subscription request with notification preferences',
            terraformSnippet: `resource "f5xc_addon_subscription" "${tfResourceName}" {
  name      = "${tfResourceName.replace(/_/g, '-')}-subscription"
  namespace = "system"

  addon_service {
    name      = "${addonService}"
    namespace = "shared"
  }

  notification_preference {
    emails {
      email_ids = ["admin@example.com"]
    }
  }
}`,
          },
          {
            step: 3,
            action: 'Wait for Processing',
            description: 'Monitor subscription state for SUBSCRIPTION_ENABLED',
          },
          {
            step: 4,
            action: 'Use Addon Features',
            description: 'Once enabled, configure addon features in your resources',
          },
        ],
        terraformConfig: generateFullTerraformConfig(service, true),
        estimatedTime: '1-24 hours (depends on SRE processing)',
        notes: [
          'Request will be in SUBSCRIPTION_PENDING state initially',
          'SRE team will process and approve the request',
          'You will receive email notification when enabled',
        ],
      };
      break;

    case 'managed':
      workflow = {
        addonService: service.name,
        activationType: 'managed',
        description: `Fully managed activation for ${service.displayName}. Requires contacting F5 Sales.`,
        prerequisites: [
          'Valid F5 Distributed Cloud tenant',
          'Sales agreement for the addon service',
          'Premium support contract (recommended)',
        ],
        steps: [
          {
            step: 1,
            action: 'Contact Sales',
            description: 'Reach out to F5 Sales to discuss requirements and pricing',
          },
          {
            step: 2,
            action: 'Complete Agreement',
            description: 'Sign the subscription agreement with F5',
          },
          {
            step: 3,
            action: 'F5 Activation',
            description: 'F5 team will activate the addon for your tenant',
          },
          {
            step: 4,
            action: 'Verify in Terraform',
            description: 'Check activation status and create data source reference',
            terraformSnippet: `data "f5xc_addon_service_activation_status" "${tfResourceName}" {
  addon_service = "${addonService}"
}

output "is_active" {
  value = data.f5xc_addon_service_activation_status.${tfResourceName}.state == "AS_SUBSCRIBED"
}`,
          },
        ],
        terraformConfig: `# For managed addons, activation is handled by F5 Sales
# Once activated, you can reference the addon in your configuration

data "f5xc_addon_service" "${tfResourceName}" {
  name = "${addonService}"
}

data "f5xc_addon_service_activation_status" "${tfResourceName}" {
  addon_service = "${addonService}"
}`,
        estimatedTime: 'Days to weeks (depends on sales process)',
        notes: [
          'This addon requires a sales agreement',
          'Contact F5 Sales: https://www.f5.com/products/get-f5',
          'Once activated, use data sources to verify status',
        ],
      };
      break;

    default:
      return null;
  }

  return workflow;
}

/**
 * Generate a full Terraform configuration for an addon
 */
function generateFullTerraformConfig(service: AddonServiceInfo, includeNotification = false): string {
  const serviceName = service.name;  // Full API name like "f5xc-bot-defense-standard"
  const resourceName = toTerraformResourceName(serviceName);  // Terraform-safe name like "bot_defense_standard"

  let config = `# ${service.displayName} Activation Configuration
# Tier Required: ${service.tier}
# Activation Type: ${service.activationType}

terraform {
  required_providers {
    f5xc = {
      source  = "registry.terraform.io/robinmordasiewicz/f5xc"
      version = ">= 0.1.0"
    }
  }
}

provider "f5xc" {
  # Configure via environment variables:
  # F5XC_API_URL - API endpoint (optional)
  # F5XC_API_TOKEN - API token (or use P12 cert)
  # F5XC_P12_FILE - Path to P12 certificate
  # F5XC_P12_PASSWORD - P12 certificate password
}

# Step 1: Check activation status
data "f5xc_addon_service_activation_status" "${serviceName}" {
  addon_service = "${serviceName}"
}

# Step 2: Get addon service details
data "f5xc_addon_service" "${serviceName}" {
  name = "${serviceName}"
}

# Step 3: Create subscription (only if can_activate is true)
resource "f5xc_addon_subscription" "${serviceName}" {
  count     = data.f5xc_addon_service_activation_status.${serviceName}.can_activate ? 1 : 0
  name      = "${resourceName}-subscription"
  namespace = "system"

  addon_service {
    name      = "${serviceName}"
    namespace = "shared"
  }`;

  if (includeNotification) {
    config += `

  notification_preference {
    emails {
      email_ids = ["admin@example.com"]
    }
  }`;
  }

  config += `
}

# Outputs
output "addon_tier" {
  description = "Required subscription tier"
  value       = data.f5xc_addon_service.${serviceName}.tier
}

output "activation_type" {
  description = "Activation type (self, partial, managed)"
  value       = data.f5xc_addon_service.${serviceName}.activation_type
}

output "can_activate" {
  description = "Whether the addon can be activated"
  value       = data.f5xc_addon_service_activation_status.${serviceName}.can_activate
}

output "current_state" {
  description = "Current subscription state"
  value       = data.f5xc_addon_service_activation_status.${serviceName}.state
}`;

  return config;
}

/**
 * Get list of addon service names
 */
export function getAddonServiceNames(): string[] {
  return Object.keys(ADDON_SERVICES);
}

/**
 * Check if an addon service exists
 */
export function addonServiceExists(name: string): boolean {
  return name in ADDON_SERVICES;
}
