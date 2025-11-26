# APM Data Source Example
# Retrieves information about an existing APM

# Look up an existing APM by name
data "f5xc_apm" "example" {
  name      = "example-apm"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "apm_id" {
#   value = data.f5xc_apm.example.id
# }
