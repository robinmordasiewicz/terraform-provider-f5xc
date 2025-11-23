# Tpm Api Key Data Source Example
# Retrieves information about an existing Tpm Api Key

# Look up an existing Tpm Api Key by name
data "f5xc_tpm_api_key" "example" {
  name      = "example-tpm-api-key"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "tpm_api_key_id" {
#   value = data.f5xc_tpm_api_key.example.id
# }
