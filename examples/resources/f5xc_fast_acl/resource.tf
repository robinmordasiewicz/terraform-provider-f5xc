# Fast ACL Resource Example
# Manages object, object contains rules to protect site from denial of service It has destination{destination IP, destination port) and references to. in F5 Distributed Cloud.

# Basic Fast ACL configuration
resource "f5xc_fast_acl" "example" {
  name      = "example-fast-acl"
  namespace = "system"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # Type establishes a direct reference from one object(the r...
  protocol_policer {
    # Configure protocol_policer settings
  }
  # [OneOf: re_acl, site_acl] Fast ACL for RE. Fast ACL defin...
  re_acl {
    # Configure re_acl settings
  }
  # Enable this option
  all_public_vips {
    # Configure all_public_vips settings
  }
}
