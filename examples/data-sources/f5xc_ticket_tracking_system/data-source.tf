# Ticket Tracking System Data Source Example
# Retrieves information about an existing Ticket Tracking System

# Look up an existing Ticket Tracking System by name
data "f5xc_ticket_tracking_system" "example" {
  name      = "example-ticket-tracking-system"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "ticket_tracking_system_id" {
#   value = data.f5xc_ticket_tracking_system.example.id
# }
