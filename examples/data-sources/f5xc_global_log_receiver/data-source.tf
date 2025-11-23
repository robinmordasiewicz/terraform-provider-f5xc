# Global Log Receiver Data Source Example
# Retrieves information about an existing Global Log Receiver

# Look up an existing Global Log Receiver by name
data "f5xc_global_log_receiver" "example" {
  name      = "example-global-log-receiver"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "global_log_receiver_id" {
#   value = data.f5xc_global_log_receiver.example.id
# }
