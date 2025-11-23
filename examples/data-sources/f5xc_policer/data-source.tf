# Policer Data Source Example
# Retrieves information about an existing Policer

# Look up an existing Policer by name
data "f5xc_policer" "example" {
  name      = "example-policer"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "policer_id" {
#   value = data.f5xc_policer.example.id
# }
