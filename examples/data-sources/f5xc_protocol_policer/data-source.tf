# Protocol Policer Data Source Example
# Retrieves information about an existing Protocol Policer

# Look up an existing Protocol Policer by name
data "f5xc_protocol_policer" "example" {
  name      = "example-protocol-policer"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "protocol_policer_id" {
#   value = data.f5xc_protocol_policer.example.id
# }
