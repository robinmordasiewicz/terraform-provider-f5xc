# Filter Set Resource Example
# Manages specification. in F5 Distributed Cloud.

# Basic Filter Set configuration
resource "f5xc_filter_set" "example" {
  name      = "example-filter-set"
  namespace = "shared"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # List of fields and their values selected by the user .
  filter_fields {
    # Configure filter_fields settings
  }
  # Either an absolute time range or a relative time interval.
  date_field {
    # Configure date_field settings
  }
  # Date range is for selecting a date range.
  absolute {
    # Configure absolute settings
  }
}
