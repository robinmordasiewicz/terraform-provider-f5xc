# Ticket Tracking System Resource Example
# [Namespace: required] Manages Ticket Tracking System in F5 Distributed Cloud.

# Basic Ticket Tracking System configuration
resource "f5xc_ticket_tracking_system" "example" {
  name      = "example-ticket-tracking-system"
  namespace = "staging"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # Jira Configuration Type.
  jira_config {
    # Configure jira_config settings
  }
  # JIRA Ad-hoc REST API Configuration Type. v3 API Basic Aut...
  adhoc_rest_api {
    # Configure adhoc_rest_api settings
  }
}
