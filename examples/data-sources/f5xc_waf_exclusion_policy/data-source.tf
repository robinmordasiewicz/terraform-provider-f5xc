# WAF Exclusion Policy Data Source Example
# Retrieves information about an existing WAF Exclusion Policy

# Look up an existing WAF Exclusion Policy by name
data "f5xc_waf_exclusion_policy" "example" {
  name      = "example-waf-exclusion-policy"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "waf_exclusion_policy_id" {
#   value = data.f5xc_waf_exclusion_policy.example.id
# }
