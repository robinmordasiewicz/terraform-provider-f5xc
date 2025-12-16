# Filter Set Resource Example
# Manages a Filter Set resource in F5 Distributed Cloud for create configuration.

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
  # Filter Fields. list of fields and their values selected b...
  filter_fields {
    # Configure filter_fields settings
  }
  # Filter Date/Time Range Field. either an absolute time ran...
  date_field {
    # Configure date_field settings
  }
  # Date Range. Date range is for selecting a date range
  absolute {
    # Configure absolute settings
  }
}
