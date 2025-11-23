# Alert Receiver Data Source Example
# Retrieves information about an existing Alert Receiver

# Look up an existing Alert Receiver by name
data "f5xc_alert_receiver" "example" {
  name      = "example-alert-receiver"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "alert_receiver_id" {
#   value = data.f5xc_alert_receiver.example.id
# }

# Example: Reference alert receiver in alert policy
# resource "f5xc_alert_policy" "example" {
#   name      = "example-policy"
#   namespace = "system"
#
#   receivers {
#     name      = data.f5xc_alert_receiver.example.name
#     namespace = data.f5xc_alert_receiver.example.namespace
#   }
# }
