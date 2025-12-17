# Addon Activation Example

This example demonstrates how to activate F5 Distributed Cloud addon services using Terraform.

## Overview

F5XC addon services are additional security and performance features that can be activated for your tenant. This example shows how to:

1. Query addon service details and tier requirements
2. Check activation eligibility for your tenant
3. Create addon subscriptions to activate services
4. Handle activation waiting and dependencies

## Supported Addon Services

| Addon Service                       | Description                                   | Tier Required |
| ----------------------------------- | --------------------------------------------- | ------------- |
| `f5xc-bot-defense-standard`         | Protect applications from automated attacks   | STANDARD      |
| `f5xc-bot-defense-advanced`         | Bot defense with advanced ML detection        | ADVANCED      |
| `f5xc-client-side-defense-standard` | Protect against Magecart and formjacking      | STANDARD      |
| `f5xc-waap-standard`                | Web App and API Protection with API Discovery | STANDARD      |
| `f5xc-waap-advanced`                | WAAP with full API security features          | ADVANCED      |

## Prerequisites

- Terraform >= 1.0
- F5 Distributed Cloud account with appropriate subscription tier
- API credentials configured

## Quick Start

### 1. Configure Authentication

Set environment variables for authentication:

```bash
# Option 1: API Token
export F5XC_API_URL="https://your-tenant.console.ves.volterra.io"
export F5XC_API_TOKEN="your-api-token"

# Option 2: P12 Certificate
export F5XC_API_URL="https://your-tenant.console.ves.volterra.io"
export F5XC_P12_FILE="/path/to/credentials.p12"
export F5XC_P12_PASSWORD="your-p12-password"  # pragma: allowlist secret
```

### 2. Configure Variables

```bash
cp terraform.tfvars.example terraform.tfvars
```

Edit `terraform.tfvars` to enable desired addons:

```hcl
enable_bot_defense = true
enable_client_side_defense = false
enable_waap = false
```

### 3. Deploy

```bash
terraform init
terraform plan
terraform apply
```

## Outputs

After applying, you'll see:

- **Addon service info**: Details about each addon (tier, activation type)
- **Activation status**: Current state and whether activation is possible
- **Activation summary**: Which addons were activated

## How It Works

1. **Check eligibility**: The example uses `f5xc_addon_service_activation_status` data source to check if each addon can be activated
2. **Conditional activation**: Subscriptions are only created if the addon is available and not already active
3. **Wait for propagation**: An optional delay ensures addons are fully active before dependent resources are created

## Customization

### Activate Additional Addons

To activate other addon services, add similar blocks to `main.tf`:

```hcl
# Example: Adding WAAP Advanced tier
data "f5xc_addon_service_activation_status" "waap_advanced" {
  count         = var.enable_waap_advanced ? 1 : 0
  addon_service = "f5xc-waap-advanced"
}

resource "f5xc_addon_subscription" "waap_advanced" {
  count = (
    var.enable_waap_advanced &&
    length(data.f5xc_addon_service_activation_status.waap_advanced) > 0 &&
    data.f5xc_addon_service_activation_status.waap_advanced[0].can_activate &&
    data.f5xc_addon_service_activation_status.waap_advanced[0].state == "AS_NONE"
  ) ? 1 : 0

  name      = "waap-advanced-subscription"
  namespace = "system"

  addon_service {
    name      = "f5xc-waap-advanced"
    namespace = "shared"
  }
}
```

### Adjust Wait Time

If addons need more time to activate:

```hcl
activation_wait_time = "1m"  # Increase to 1 minute
```

## Troubleshooting

### Addon Not Activating

1. Check the activation status output for the `can_activate` and `state` values
2. Verify your subscription tier supports the addon
3. Check F5XC console for any pending approvals

### State Shows "AS_PENDING"

Some addons require SRE approval. Wait for the approval process to complete.

### State Shows "AS_ERROR"

Contact F5 support with your tenant ID and the specific error message.

## Related Documentation

- [Addon Activation Guide](https://registry.terraform.io/providers/robinmordasiewicz/f5xc/latest/docs/guides/addon-activation)
- [f5xc_addon_service Data Source](https://registry.terraform.io/providers/robinmordasiewicz/f5xc/latest/docs/data-sources/addon_service)
- [f5xc_addon_service_activation_status Data Source](https://registry.terraform.io/providers/robinmordasiewicz/f5xc/latest/docs/data-sources/addon_service_activation_status)
- [f5xc_addon_subscription Resource](https://registry.terraform.io/providers/robinmordasiewicz/f5xc/latest/docs/resources/addon_subscription)
