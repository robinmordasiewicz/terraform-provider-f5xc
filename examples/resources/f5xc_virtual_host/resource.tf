# Virtual Host Resource Example
# Manages virtual host in a given namespace. in F5 Distributed Cloud.

# Basic Virtual Host configuration
resource "f5xc_virtual_host" "example" {
  name      = "example-virtual-host"
  namespace = "staging"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # Advertise Policies. Advertise Policy allows you to define...
  advertise_policies {
    # Configure advertise_policies settings
  }
  # [OneOf: authentication, no_authentication; Default: no_au...
  authentication {
    # Configure authentication settings
  }
  # Reference to Authentication Object. Reference to Authenti...
  auth_config {
    # Configure auth_config settings
  }
}
