# Irule Data Source Example
# Retrieves information about an existing Irule

# Look up an existing Irule by name
data "f5xc_irule" "example" {
  name      = "example-irule"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "irule_id" {
#   value = data.f5xc_irule.example.id
# }
