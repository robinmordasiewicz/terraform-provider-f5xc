# F5 Distributed Cloud Provider - Addon Activation Outputs
# =========================================================

# =============================================================================
# ADDON SERVICE INFORMATION
# =============================================================================

output "bot_defense_info" {
  description = "Bot Defense addon service details"
  value = var.enable_bot_defense ? {
    display_name    = try(data.f5xc_addon_service.bot_defense[0].display_name, "N/A")
    tier            = try(data.f5xc_addon_service.bot_defense[0].tier, "N/A")
    activation_type = try(data.f5xc_addon_service.bot_defense[0].activation_type, "N/A")
  } : null
}

output "client_side_defense_info" {
  description = "Client Side Defense addon service details"
  value = var.enable_client_side_defense ? {
    display_name    = try(data.f5xc_addon_service.client_side_defense[0].display_name, "N/A")
    tier            = try(data.f5xc_addon_service.client_side_defense[0].tier, "N/A")
    activation_type = try(data.f5xc_addon_service.client_side_defense[0].activation_type, "N/A")
  } : null
}

output "waap_info" {
  description = "WAAP addon service details"
  value = var.enable_waap ? {
    display_name    = try(data.f5xc_addon_service.waap[0].display_name, "N/A")
    tier            = try(data.f5xc_addon_service.waap[0].tier, "N/A")
    activation_type = try(data.f5xc_addon_service.waap[0].activation_type, "N/A")
  } : null
}

# =============================================================================
# ACTIVATION STATUS
# =============================================================================

output "bot_defense_status" {
  description = "Bot Defense activation status"
  value = var.enable_bot_defense ? {
    state        = try(data.f5xc_addon_service_activation_status.bot_defense[0].state, "NOT_CHECKED")
    can_activate = try(data.f5xc_addon_service_activation_status.bot_defense[0].can_activate, false)
    message      = try(data.f5xc_addon_service_activation_status.bot_defense[0].message, "Not checked")
  } : null
}

output "client_side_defense_status" {
  description = "Client Side Defense activation status"
  value = var.enable_client_side_defense ? {
    state        = try(data.f5xc_addon_service_activation_status.client_side_defense[0].state, "NOT_CHECKED")
    can_activate = try(data.f5xc_addon_service_activation_status.client_side_defense[0].can_activate, false)
    message      = try(data.f5xc_addon_service_activation_status.client_side_defense[0].message, "Not checked")
  } : null
}

output "waap_status" {
  description = "WAAP activation status"
  value = var.enable_waap ? {
    state        = try(data.f5xc_addon_service_activation_status.waap[0].state, "NOT_CHECKED")
    can_activate = try(data.f5xc_addon_service_activation_status.waap[0].can_activate, false)
    message      = try(data.f5xc_addon_service_activation_status.waap[0].message, "Not checked")
  } : null
}

# =============================================================================
# SUBSCRIPTION RESULTS
# =============================================================================

output "activated_addons" {
  description = "List of addon subscriptions that were created"
  value = compact([
    length(f5xc_addon_subscription.bot_defense) > 0 ? "f5xc-bot-defense-standard" : "",
    length(f5xc_addon_subscription.client_side_defense) > 0 ? "f5xc-client-side-defense-standard" : "",
    length(f5xc_addon_subscription.waap) > 0 ? "f5xc-waap-standard" : "",
  ])
}

output "activation_summary" {
  description = "Summary of addon activation results"
  value = {
    total_requested = sum([
      var.enable_bot_defense ? 1 : 0,
      var.enable_client_side_defense ? 1 : 0,
      var.enable_waap ? 1 : 0,
    ])
    total_activated = sum([
      length(f5xc_addon_subscription.bot_defense),
      length(f5xc_addon_subscription.client_side_defense),
      length(f5xc_addon_subscription.waap),
    ])
    bot_defense = {
      requested = var.enable_bot_defense
      activated = length(f5xc_addon_subscription.bot_defense) > 0
    }
    client_side_defense = {
      requested = var.enable_client_side_defense
      activated = length(f5xc_addon_subscription.client_side_defense) > 0
    }
    waap = {
      requested = var.enable_waap
      activated = length(f5xc_addon_subscription.waap) > 0
    }
  }
}

# =============================================================================
# DEMO NAMESPACE
# =============================================================================

output "demo_namespace" {
  description = "Demo namespace details (if created)"
  value = var.create_demo_namespace && length(f5xc_namespace.demo) > 0 ? {
    name        = f5xc_namespace.demo[0].name
    description = f5xc_namespace.demo[0].description
  } : null
}
