# Quota Data Source Example
# Retrieves information about an existing Quota

# Look up an existing Quota by name
data "f5xc_quota" "example" {
  name      = "example-quota"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "quota_id" {
#   value = data.f5xc_quota.example.id
# }
