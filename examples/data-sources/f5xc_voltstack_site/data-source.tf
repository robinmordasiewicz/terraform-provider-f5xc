# Voltstack Site Data Source Example
# Retrieves information about an existing Voltstack Site

# Look up an existing Voltstack Site by name
data "f5xc_voltstack_site" "example" {
  name      = "example-voltstack-site"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "voltstack_site_id" {
#   value = data.f5xc_voltstack_site.example.id
# }
