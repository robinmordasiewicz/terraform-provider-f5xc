# IRULE Data Source Example
# Retrieves information about an existing IRULE

# Look up an existing IRULE by name
data "f5xc_irule" "example" {
  name      = "example-irule"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "irule_id" {
#   value = data.f5xc_irule.example.id
# }
