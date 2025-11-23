# Data Group Resource Example
# Manages data group in a given namespace. If one already exists it will give an error. in F5 Distributed Cloud.

# Basic Data Group configuration
resource "f5xc_data_group" "example" {
  name      = "example-data-group"
  namespace = "system"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # [OneOf: address_records, integer_records, string_records]...
  address_records {
    # Configure address_records settings
  }
  # Address records.
  records {
    # Configure records settings
  }
  # Integer record List. Data group with integer record List
  integer_records {
    # Configure integer_records settings
  }
}
