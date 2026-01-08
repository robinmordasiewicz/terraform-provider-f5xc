# Advertise Policy Resource Example
# Manages a Advertise Policy resource in F5 Distributed Cloud for advertise_policy object controls how and where a service represented by a given virtual_host object is advertised to consumers. configuration.

# Basic Advertise Policy configuration
resource "f5xc_advertise_policy" "example" {
  name      = "example-advertise-policy"
  namespace = "staging"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # Optional. Public VIP to advertise This field is mutually ...
  public_ip {
    # Configure public_ip settings
  }
  # TLS configuration for downstream connections.
  tls_parameters {
    # Configure tls_parameters settings
  }
  # Can be used for messages where no values are needed.
  client_certificate_optional {
    # Configure client_certificate_optional settings
  }
}
