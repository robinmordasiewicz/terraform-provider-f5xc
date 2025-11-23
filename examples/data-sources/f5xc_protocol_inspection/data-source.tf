# Protocol Inspection Data Source Example
# Retrieves information about an existing Protocol Inspection

# Look up an existing Protocol Inspection by name
data "f5xc_protocol_inspection" "example" {
  name      = "example-protocol-inspection"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "protocol_inspection_id" {
#   value = data.f5xc_protocol_inspection.example.id
# }
