# Tpm Manager Data Source Example
# Retrieves information about an existing Tpm Manager

# Look up an existing Tpm Manager by name
data "f5xc_tpm_manager" "example" {
  name      = "example-tpm-manager"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "tpm_manager_id" {
#   value = data.f5xc_tpm_manager.example.id
# }
