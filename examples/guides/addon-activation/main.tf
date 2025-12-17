# F5 Distributed Cloud Provider - Addon Activation Example
# =========================================================
#
# This example demonstrates how to activate F5XC addon services
# using Terraform. It shows the complete workflow from checking
# eligibility to activating and using addon features.
#
# QUICK START:
# 1. Configure authentication (environment variables recommended)
# 2. Copy terraform.tfvars.example to terraform.tfvars
# 3. Enable desired addons in terraform.tfvars
# 4. Run: terraform init && terraform apply

terraform {
  required_version = ">= 1.0"

  required_providers {
    f5xc = {
      source = "robinmordasiewicz/f5xc"
    }
    time = {
      source  = "hashicorp/time"
      version = ">= 0.9.0"
    }
  }
}

# =============================================================================
# PROVIDER CONFIGURATION
# =============================================================================
#
# Authentication via environment variables (recommended):
#   export F5XC_API_URL="https://your-tenant.console.ves.volterra.io"
#   export F5XC_API_TOKEN="your-api-token"
#
# Or for P12 certificate:
#   export F5XC_API_URL="https://your-tenant.console.ves.volterra.io"
#   export F5XC_P12_FILE="/path/to/credentials.p12"
#   export F5XC_P12_PASSWORD="your-p12-password"  # pragma: allowlist secret

provider "f5xc" {
  # Authentication via environment variables
}

# =============================================================================
# ADDON SERVICE INFORMATION
# =============================================================================

# Get details about Bot Defense addon service
data "f5xc_addon_service" "bot_defense" {
  count = var.enable_bot_defense ? 1 : 0
  name  = "f5xc-bot-defense-standard"
}

# Get details about Client Side Defense addon service
data "f5xc_addon_service" "client_side_defense" {
  count = var.enable_client_side_defense ? 1 : 0
  name  = "f5xc-client-side-defense-standard"
}

# Get details about WAAP (Web App and API Protection) addon service
data "f5xc_addon_service" "waap" {
  count = var.enable_waap ? 1 : 0
  name  = "f5xc-waap-standard"
}

# =============================================================================
# ACTIVATION STATUS CHECKS
# =============================================================================

# Check Bot Defense activation status
data "f5xc_addon_service_activation_status" "bot_defense" {
  count         = var.enable_bot_defense ? 1 : 0
  addon_service = "f5xc-bot-defense-standard"
}

# Check Client Side Defense activation status
data "f5xc_addon_service_activation_status" "client_side_defense" {
  count         = var.enable_client_side_defense ? 1 : 0
  addon_service = "f5xc-client-side-defense-standard"
}

# Check WAAP activation status
data "f5xc_addon_service_activation_status" "waap" {
  count         = var.enable_waap ? 1 : 0
  addon_service = "f5xc-waap-standard"
}

# =============================================================================
# ADDON SUBSCRIPTIONS
# =============================================================================

# Activate Bot Defense if enabled and available
resource "f5xc_addon_subscription" "bot_defense" {
  count = (
    var.enable_bot_defense &&
    length(data.f5xc_addon_service_activation_status.bot_defense) > 0 &&
    data.f5xc_addon_service_activation_status.bot_defense[0].can_activate &&
    data.f5xc_addon_service_activation_status.bot_defense[0].state == "AS_NONE"
  ) ? 1 : 0

  name      = "bot-defense-subscription"
  namespace = "system"

  addon_service {
    name      = "f5xc-bot-defense-standard"
    namespace = "shared"
  }
}

# Activate Client Side Defense if enabled and available
resource "f5xc_addon_subscription" "client_side_defense" {
  count = (
    var.enable_client_side_defense &&
    length(data.f5xc_addon_service_activation_status.client_side_defense) > 0 &&
    data.f5xc_addon_service_activation_status.client_side_defense[0].can_activate &&
    data.f5xc_addon_service_activation_status.client_side_defense[0].state == "AS_NONE"
  ) ? 1 : 0

  name      = "client-side-defense-subscription"
  namespace = "system"

  addon_service {
    name      = "f5xc-client-side-defense-standard"
    namespace = "shared"
  }
}

# Activate WAAP if enabled and available
resource "f5xc_addon_subscription" "waap" {
  count = (
    var.enable_waap &&
    length(data.f5xc_addon_service_activation_status.waap) > 0 &&
    data.f5xc_addon_service_activation_status.waap[0].can_activate &&
    data.f5xc_addon_service_activation_status.waap[0].state == "AS_NONE"
  ) ? 1 : 0

  name      = "waap-subscription"
  namespace = "system"

  addon_service {
    name      = "f5xc-waap-standard"
    namespace = "shared"
  }
}

# =============================================================================
# WAIT FOR ACTIVATION (Optional)
# =============================================================================

# Wait for activation to propagate before using features
resource "time_sleep" "wait_for_activation" {
  count = (
    length(f5xc_addon_subscription.bot_defense) > 0 ||
    length(f5xc_addon_subscription.client_side_defense) > 0 ||
    length(f5xc_addon_subscription.waap) > 0
  ) ? 1 : 0

  depends_on = [
    f5xc_addon_subscription.bot_defense,
    f5xc_addon_subscription.client_side_defense,
    f5xc_addon_subscription.waap,
  ]

  create_duration = var.activation_wait_time
}

# =============================================================================
# EXAMPLE NAMESPACE (for demonstration)
# =============================================================================

resource "f5xc_namespace" "demo" {
  count       = var.create_demo_namespace ? 1 : 0
  name        = var.demo_namespace_name
  description = "Namespace for addon activation demonstration"

  depends_on = [time_sleep.wait_for_activation]
}
