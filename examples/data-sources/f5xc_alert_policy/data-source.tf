# Alert Policy Data Source Example
# Retrieves information about an existing Alert Policy

# Look up an existing Alert Policy by name
data "f5xc_alert_policy" "example" {
  name      = "example-alert-policy"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "alert_policy_id" {
#   value = data.f5xc_alert_policy.example.id
# }
