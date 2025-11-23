# Data Type Data Source Example
# Retrieves information about an existing Data Type

# Look up an existing Data Type by name
data "f5xc_data_type" "example" {
  name      = "example-data-type"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "data_type_id" {
#   value = data.f5xc_data_type.example.id
# }
