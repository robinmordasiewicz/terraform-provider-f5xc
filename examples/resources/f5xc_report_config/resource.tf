# Report Config Resource Example
# Manages a ReportConfig resource in F5 Distributed Cloud for report configuration is used to schedule report generation at a later point in time. configuration.

# Basic Report Config configuration
resource "f5xc_report_config" "example" {
  name      = "example-report-config"
  namespace = "system"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # Report recipients. Report recipients
  report_recipients {
    # Configure report_recipients settings
  }
  # User Groups. Select one or more user groups, to which the...
  user_groups {
    # Configure user_groups settings
  }
  # Report Type Waap. Report Type Waap
  waap {
    # Configure waap settings
  }
}
