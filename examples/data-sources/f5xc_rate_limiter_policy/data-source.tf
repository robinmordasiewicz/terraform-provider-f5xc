# Rate Limiter Policy Data Source Example
# Retrieves information about an existing Rate Limiter Policy

# Look up an existing Rate Limiter Policy by name
data "f5xc_rate_limiter_policy" "example" {
  name      = "example-rate-limiter-policy"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "rate_limiter_policy_id" {
#   value = data.f5xc_rate_limiter_policy.example.id
# }
