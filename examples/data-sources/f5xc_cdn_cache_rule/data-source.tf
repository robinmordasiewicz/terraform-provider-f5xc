# Cdn Cache Rule Data Source Example
# Retrieves information about an existing Cdn Cache Rule

# Look up an existing Cdn Cache Rule by name
data "f5xc_cdn_cache_rule" "example" {
  name      = "example-cdn-cache-rule"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "cdn_cache_rule_id" {
#   value = data.f5xc_cdn_cache_rule.example.id
# }
