# Log Receiver Data Source Example
# Retrieves information about an existing Log Receiver

# Look up an existing Log Receiver by name
data "f5xc_log_receiver" "example" {
  name      = "example-log-receiver"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "log_receiver_id" {
#   value = data.f5xc_log_receiver.example.id
# }

# Example: Reference log receiver in site configuration
# resource "f5xc_securemesh_site_v2" "example" {
#   name      = "example-site"
#   namespace = "system"
#
#   log_receiver {
#     name      = data.f5xc_log_receiver.example.name
#     namespace = data.f5xc_log_receiver.example.namespace
#   }
# }
