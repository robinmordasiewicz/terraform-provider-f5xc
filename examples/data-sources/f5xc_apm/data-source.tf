# Apm Data Source Example
# Retrieves information about an existing Apm

# Look up an existing Apm by name
data "f5xc_apm" "example" {
  name      = "example-apm"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "apm_id" {
#   value = data.f5xc_apm.example.id
# }
