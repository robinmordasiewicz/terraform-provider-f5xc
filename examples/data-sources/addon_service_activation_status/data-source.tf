# Example: Check addon service activation status
# Use this to determine if an addon service can be activated

data "f5xc_addon_service_activation_status" "bot_defense" {
  addon_service = "bot_defense"
}

# Output the activation status
output "state" {
  description = "Current subscription state (AS_NONE, AS_PENDING, AS_SUBSCRIBED, AS_ERROR)"
  value       = data.f5xc_addon_service_activation_status.bot_defense.state
}

output "can_activate" {
  description = "Whether the addon can be activated"
  value       = data.f5xc_addon_service_activation_status.bot_defense.can_activate
}

output "status_message" {
  description = "Human-readable status message"
  value       = data.f5xc_addon_service_activation_status.bot_defense.message
}

# Example: Conditional subscription based on activation status
# Only create the subscription if the addon is available for activation
resource "f5xc_addon_subscription" "bot_defense" {
  count = data.f5xc_addon_service_activation_status.bot_defense.can_activate && data.f5xc_addon_service_activation_status.bot_defense.state == "AS_NONE" ? 1 : 0

  name      = "my-bot-defense-subscription"
  namespace = "system"

  addon_service {
    name      = "bot_defense"
    namespace = "shared"
  }
}
