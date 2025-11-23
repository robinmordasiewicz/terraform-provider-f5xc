# App Type Data Source Example
# Retrieves information about an existing App Type

# Look up an existing App Type by name
data "f5xc_app_type" "example" {
  name      = "example-app-type"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "app_type_id" {
#   value = data.f5xc_app_type.example.id
# }
