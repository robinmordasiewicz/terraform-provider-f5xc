# Example: Query addon service details
# This data source retrieves information about an F5 Distributed Cloud addon service

data "f5xc_addon_service" "bot_defense" {
  name = "bot_defense"
}

# Output the addon service details
output "display_name" {
  description = "Human-readable name of the addon service"
  value       = data.f5xc_addon_service.bot_defense.display_name
}

output "tier" {
  description = "Required subscription tier (NO_TIER, BASIC, STANDARD, ADVANCED, PREMIUM)"
  value       = data.f5xc_addon_service.bot_defense.tier
}

output "activation_type" {
  description = "How the service is activated (self, partial, managed)"
  value       = data.f5xc_addon_service.bot_defense.activation_type
}

# Example: Check if self-activation is available
output "is_self_activatable" {
  description = "Whether the addon can be activated without manual intervention"
  value       = data.f5xc_addon_service.bot_defense.activation_type == "self"
}
