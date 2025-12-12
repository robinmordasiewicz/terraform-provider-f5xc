# Alert Policy Resource Example
# [Category: Monitoring] [Namespace: required] Manages a Alert Policy resource in F5 Distributed Cloud for alerting rules and notification policies.

# Basic Alert Policy configuration
resource "f5xc_alert_policy" "example" {
  name      = "example-alert-policy"
  namespace = "shared"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Alert Policy configuration
  # Alert receivers
  receivers {
    name      = "slack-receiver"
    namespace = "shared"
  }

  # Alert routes
  routes {
    any {}
    send {}
  }

  # Notification parameters
  notification_parameters {
    default {}
    group_wait     = "30s"
    group_interval = "1m"
  }
}
