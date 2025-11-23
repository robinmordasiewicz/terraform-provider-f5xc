# App Api Group Data Source Example
# Retrieves information about an existing App Api Group

# Look up an existing App Api Group by name
data "f5xc_app_api_group" "example" {
  name      = "example-app-api-group"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "app_api_group_id" {
#   value = data.f5xc_app_api_group.example.id
# }
