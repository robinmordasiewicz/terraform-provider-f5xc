---
page_title: "Guide: Addon Service Activation"
subcategory: "Guides"
description: |-
  Learn how to activate F5XC addon services using Terraform.
  Covers Bot Defense, Client Side Defense, WAAP, and more.
---

# Addon Service Activation

This guide walks you through activating F5 Distributed Cloud addon services using Terraform. By the end, you'll understand how to:

- **Check activation eligibility** - Determine if an addon can be activated
- **Activate self-service addons** - Bot Defense, Client Side Defense, etc.
- **Handle managed activation** - Services requiring sales contact
- **Monitor activation status** - Track subscription state changes

## Overview

F5 Distributed Cloud addon services are additional security and performance features that can be activated for your tenant. These include:

| Addon Service                        | Description                                   | Tier Required |
| ------------------------------------ | --------------------------------------------- | ------------- |
| `f5xc-bot-defense-standard`          | Protect applications from automated attacks   | STANDARD      |
| `f5xc-bot-defense-advanced`          | Bot defense with advanced ML detection        | ADVANCED      |
| `f5xc-client-side-defense-standard`  | Protect against Magecart and formjacking      | STANDARD      |
| `f5xc-waap-standard`                 | Web App and API Protection with API Discovery | STANDARD      |
| `f5xc-waap-advanced`                 | WAAP with full API security features          | ADVANCED      |
| `f5xc-malicious-user-detection`      | Identify malicious user behavior patterns     | ADVANCED      |
| `f5xc-synthetic-monitoring`          | Monitor application availability              | STANDARD      |

### Activation Types

Addon services have different activation types that determine how they can be activated:

```text
┌─────────────────────────────────────────────────────────────────────┐
│                     Activation Types                                │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│  SELF-ACTIVATION                                                    │
│  ┌──────────────┐    ┌──────────────┐    ┌──────────────┐          │
│  │ Check Status │───►│ Create       │───►│ Active       │          │
│  │ (AS_NONE)    │    │ Subscription │    │ (AS_SUBSCRIBED) │       │
│  └──────────────┘    └──────────────┘    └──────────────┘          │
│  User can activate directly via Terraform                           │
│                                                                     │
│  PARTIALLY MANAGED                                                  │
│  ┌──────────────┐    ┌──────────────┐    ┌──────────────┐          │
│  │ Check Status │───►│ Request      │───►│ SRE Review   │          │
│  │ (AS_NONE)    │    │ Subscription │    │ (AS_PENDING) │          │
│  └──────────────┘    └──────────────┘    └──────┬───────┘          │
│                                                  │                  │
│                                          ┌──────▼───────┐          │
│                                          │ Active       │          │
│                                          │ (AS_SUBSCRIBED) │       │
│                                          └──────────────┘          │
│  User initiates, SRE team processes                                 │
│                                                                     │
│  FULLY MANAGED                                                      │
│  ┌──────────────┐    ┌──────────────┐    ┌──────────────┐          │
│  │ Contact      │───►│ Sales        │───►│ F5 Activates │          │
│  │ F5 Sales     │    │ Agreement    │    │ Addon        │          │
│  └──────────────┘    └──────────────┘    └──────────────┘          │
│  Requires sales engagement                                          │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

## Prerequisites

Before you begin, ensure you have:

- **Terraform >= 1.0** - The F5XC provider requires Terraform 1.0 or later
- **F5 Distributed Cloud Account** - Sign up at <https://www.f5.com/cloud/products/distributed-cloud-console>
- **API Credentials** - Token or P12 certificate authentication configured
- **Appropriate Subscription Tier** - Most addon services require ADVANCED tier

### Authentication Setup

Configure one of these authentication methods via environment variables:

#### Option 1: API Token (Recommended for development)

```bash
export F5XC_API_URL="https://your-tenant.console.ves.volterra.io"
export F5XC_API_TOKEN="your-api-token"
```

#### Option 2: P12 Certificate (Recommended for production)

```bash
export F5XC_API_URL="https://your-tenant.console.ves.volterra.io"
export F5XC_P12_FILE="/path/to/your-credentials.p12"
export F5XC_P12_PASSWORD="your-p12-password"  # pragma: allowlist secret
```

## Quick Start

### Step 1: Clone the Repository

```bash
git clone https://github.com/robinmordasiewicz/terraform-provider-f5xc.git
cd terraform-provider-f5xc/examples/guides/addon-activation
```

### Step 2: Configure Your Deployment

```bash
cp terraform.tfvars.example terraform.tfvars
```

Edit `terraform.tfvars` to enable the addon services you want to activate:

```hcl
# Enable Bot Defense activation
enable_bot_defense = true

# Enable Client Side Defense
enable_client_side_defense = false
```

### Step 3: Initialize and Apply

```bash
terraform init
terraform plan
terraform apply
```

## Checking Activation Eligibility

Before attempting to activate an addon service, check if it's available for your tenant.

### Using the Activation Status Data Source

```hcl
# Check if Bot Defense can be activated
data "f5xc_addon_service_activation_status" "bot_defense" {
  addon_service = "f5xc-bot-defense-standard"
}

output "bot_defense_status" {
  value = {
    state        = data.f5xc_addon_service_activation_status.bot_defense.state
    can_activate = data.f5xc_addon_service_activation_status.bot_defense.can_activate
    message      = data.f5xc_addon_service_activation_status.bot_defense.message
  }
}
```

### State Values

| State           | Description            | Can Activate?        |
| --------------- | ---------------------- | -------------------- |
| `AS_NONE`       | Service not subscribed | Yes                  |
| `AS_PENDING`    | Activation in progress | No (wait)            |
| `AS_SUBSCRIBED` | Already active         | Already done         |
| `AS_ERROR`      | Subscription error     | No (contact support) |

### Querying Addon Service Details

```hcl
# Get detailed information about an addon service
data "f5xc_addon_service" "bot_defense" {
  name = "f5xc-bot-defense-standard"
}

output "addon_details" {
  value = {
    display_name    = data.f5xc_addon_service.bot_defense.display_name
    tier            = data.f5xc_addon_service.bot_defense.tier
    activation_type = data.f5xc_addon_service.bot_defense.activation_type
  }
}
```

## Self-Activation Workflow

For addon services with `self` activation type, you can activate directly via Terraform.

### Basic Self-Activation

```hcl
# Step 1: Check if we can activate
data "f5xc_addon_service_activation_status" "bot_defense" {
  addon_service = "f5xc-bot-defense-standard"
}

# Step 2: Create subscription only if available
resource "f5xc_addon_subscription" "bot_defense" {
  count = data.f5xc_addon_service_activation_status.bot_defense.can_activate && data.f5xc_addon_service_activation_status.bot_defense.state == "AS_NONE" ? 1 : 0

  name      = "bot-defense-subscription"
  namespace = "system"

  addon_service {
    name      = "f5xc-bot-defense-standard"
    namespace = "shared"
  }
}

output "activation_result" {
  value = length(f5xc_addon_subscription.bot_defense) > 0 ? "Activated" : "Not activated (check status)"
}
```

### Multiple Addon Activation

```hcl
locals {
  # Define the addons you want to activate
  addons_to_activate = [
    "f5xc-bot-defense-standard",
    "f5xc-client-side-defense-standard",
    "f5xc-waap-standard",
  ]
}

# Check activation status for each
data "f5xc_addon_service_activation_status" "addons" {
  for_each      = toset(local.addons_to_activate)
  addon_service = each.value
}

# Create subscriptions for available addons
resource "f5xc_addon_subscription" "addons" {
  for_each = {
    for addon in local.addons_to_activate :
    addon => addon
    if data.f5xc_addon_service_activation_status.addons[addon].can_activate && data.f5xc_addon_service_activation_status.addons[addon].state == "AS_NONE"
  }

  name      = "${replace(replace(each.value, "f5xc-", ""), "-standard", "")}-subscription"
  namespace = "system"

  addon_service {
    name      = each.value
    namespace = "shared"
  }
}
```

## Waiting for Activation

Some addons may take time to activate, especially those with partial management. Here are patterns for handling this.

### Pattern 1: Using terraform_data with Precondition

```hcl
# Check status after subscription
data "f5xc_addon_service_activation_status" "bot_defense_status" {
  addon_service = "f5xc-bot-defense-standard"

  depends_on = [f5xc_addon_subscription.bot_defense]
}

# Validate activation succeeded
resource "terraform_data" "validate_activation" {
  lifecycle {
    precondition {
      condition     = data.f5xc_addon_service_activation_status.bot_defense_status.state == "AS_SUBSCRIBED"
      error_message = "Bot Defense activation not yet complete. Current state: ${data.f5xc_addon_service_activation_status.bot_defense_status.state}"
    }
  }
}
```

### Pattern 2: Using time_sleep for Simple Delays

```hcl
resource "f5xc_addon_subscription" "bot_defense" {
  name      = "bot-defense-subscription"
  namespace = "system"

  addon_service {
    name      = "f5xc-bot-defense-standard"
    namespace = "shared"
  }
}

# Wait for activation to propagate
resource "time_sleep" "wait_for_activation" {
  depends_on = [f5xc_addon_subscription.bot_defense]

  create_duration = "30s"
}

# Use the addon feature after waiting
resource "f5xc_http_loadbalancer" "with_bot_defense" {
  depends_on = [time_sleep.wait_for_activation]
  # ... configuration with bot defense enabled
}
```

### Pattern 3: External Verification Script

For critical deployments, you may want to verify activation before proceeding:

```hcl
resource "null_resource" "verify_activation" {
  depends_on = [f5xc_addon_subscription.bot_defense]

  provisioner "local-exec" {
    command = <<-EOT
      for i in {1..30}; do
        status=$(curl -s -H "Authorization: APIToken $F5XC_API_TOKEN" \
          "$F5XC_API_URL/api/web/namespaces/system/addon_services/f5xc-bot-defense-standard/activation-status" \
          | jq -r '.state')
        if [ "$status" = "AS_SUBSCRIBED" ]; then
          echo "Activation complete!"
          exit 0
        fi
        echo "Waiting for activation... (attempt $i/30, status: $status)"
        sleep 10
      done
      echo "Activation timeout"
      exit 1
    EOT
  }
}
```

## Managed Activation Workflow

For addon services requiring sales contact, use Terraform to monitor status after F5 activates the service.

### Verifying Managed Addon Status

```hcl
# For managed addons, just check status (don't try to create subscription)
data "f5xc_addon_service_activation_status" "managed_addon" {
  addon_service = "some_managed_addon"
}

output "managed_addon_status" {
  value = {
    active  = data.f5xc_addon_service_activation_status.managed_addon.state == "AS_SUBSCRIBED"
    message = data.f5xc_addon_service_activation_status.managed_addon.message
  }
}

# Use conditional logic based on activation status
resource "f5xc_http_loadbalancer" "with_managed_feature" {
  count = data.f5xc_addon_service_activation_status.managed_addon.state == "AS_SUBSCRIBED" ? 1 : 0

  # Configuration that uses the managed addon feature
  name      = "lb-with-managed-addon"
  namespace = "production"
  # ... rest of configuration
}
```

## Using Addon Features

Once an addon is activated, you can use its features in your configurations.

### Bot Defense in HTTP Load Balancer

```hcl
resource "f5xc_http_loadbalancer" "with_bot_defense" {
  depends_on = [f5xc_addon_subscription.bot_defense]

  name      = "my-protected-app"
  namespace = "production"

  domains = ["app.example.com"]

  default_route_pools {
    pool {
      name      = f5xc_origin_pool.backend.name
      namespace = "production"
    }
    weight = 1
  }

  # Enable Bot Defense
  bot_defense {
    policy {
      name      = "my-bot-policy"
      namespace = "shared"
    }
  }

  http {
    port = 80
  }
}
```

### Client Side Defense

```hcl
resource "f5xc_http_loadbalancer" "with_csd" {
  depends_on = [f5xc_addon_subscription.client_side_defense]

  name      = "my-protected-app"
  namespace = "production"

  domains = ["app.example.com"]

  # Enable Client Side Defense
  enable_client_side_defense = true

  # ... rest of configuration
}
```

## Troubleshooting

### Common Issues

#### Access denied when creating subscription

- Verify your API token has addon management permissions
- Check that your subscription tier supports the addon

#### Activation stuck in AS_PENDING

- For partially managed addons, contact F5 support
- For self-activation, wait and retry after a few minutes

#### State shows AS_ERROR

- Check F5XC console for detailed error messages
- Contact F5 support with your tenant ID

### Debugging Tips

```hcl
# Output detailed status for debugging
output "debug_addon_status" {
  value = {
    addon_service = "f5xc-bot-defense-standard"
    state         = data.f5xc_addon_service_activation_status.bot_defense.state
    can_activate  = data.f5xc_addon_service_activation_status.bot_defense.can_activate
    message       = data.f5xc_addon_service_activation_status.bot_defense.message
  }
}
```

## Best Practices

1. **Always check eligibility first** - Use the activation status data source before attempting activation
2. **Use conditional resource creation** - Only create subscriptions when `can_activate` is true
3. **Handle dependencies properly** - Use `depends_on` to ensure addons are active before using features
4. **Monitor activation state** - For partially managed addons, monitor the state for completion
5. **Document addon requirements** - Clearly document which addons your configuration requires

## Complete Example

See the [addon-activation example](https://github.com/robinmordasiewicz/terraform-provider-f5xc/tree/main/examples/guides/addon-activation) for a complete, working Terraform configuration.

## Related Resources

- [f5xc_addon_service Data Source](../data-sources/addon_service)
- [f5xc_addon_service_activation_status Data Source](../data-sources/addon_service_activation_status)
- [f5xc_addon_subscription Resource](../resources/addon_subscription)
- [HTTP Load Balancer Resource](../resources/http_loadbalancer)
