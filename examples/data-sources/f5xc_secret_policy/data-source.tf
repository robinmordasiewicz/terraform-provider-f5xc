# Secret Policy Data Source Example
# Retrieves information about an existing Secret Policy

# Look up an existing Secret Policy by name
data "f5xc_secret_policy" "example" {
  name      = "example-secret-policy"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "secret_policy_id" {
#   value = data.f5xc_secret_policy.example.id
# }
