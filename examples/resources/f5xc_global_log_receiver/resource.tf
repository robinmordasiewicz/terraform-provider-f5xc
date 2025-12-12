# Global Log Receiver Resource Example
# [Category: Monitoring] [Namespace: required] Manages a Global Log Receiver resource in F5 Distributed Cloud for global log aggregation settings.

# Basic Global Log Receiver configuration
resource "f5xc_global_log_receiver" "example" {
  name      = "example-global-log-receiver"
  namespace = "system"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # [OneOf: audit_logs, dns_logs, request_logs, security_even...
  audit_logs {
    # Configure audit_logs settings
  }
  # [OneOf: aws_cloud_watch_receiver, azure_event_hubs_receiv...
  aws_cloud_watch_receiver {
    # Configure aws_cloud_watch_receiver settings
  }
  # Object reference. This type establishes a direct referenc...
  aws_cred {
    # Configure aws_cred settings
  }
}
