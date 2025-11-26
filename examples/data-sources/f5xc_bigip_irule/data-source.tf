# BigIP IRULE Data Source Example
# Retrieves information about an existing BigIP IRULE

# Look up an existing BigIP IRULE by name
data "f5xc_bigip_irule" "example" {
  name      = "example-bigip-irule"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "bigip_irule_id" {
#   value = data.f5xc_bigip_irule.example.id
# }
