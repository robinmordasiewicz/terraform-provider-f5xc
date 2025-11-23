# Tpm Category Data Source Example
# Retrieves information about an existing Tpm Category

# Look up an existing Tpm Category by name
data "f5xc_tpm_category" "example" {
  name      = "example-tpm-category"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "tpm_category_id" {
#   value = data.f5xc_tpm_category.example.id
# }
