# Addon Subscription Data Source Example
# Retrieves information about an existing Addon Subscription

# Look up an existing Addon Subscription by name
data "f5xc_addon_subscription" "example" {
  name      = "example-addon-subscription"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "addon_subscription_id" {
#   value = data.f5xc_addon_subscription.example.id
# }
