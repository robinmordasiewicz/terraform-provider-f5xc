# Data Group Data Source Example
# Retrieves information about an existing Data Group

# Look up an existing Data Group by name
data "f5xc_data_group" "example" {
  name      = "example-data-group"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "data_group_id" {
#   value = data.f5xc_data_group.example.id
# }
