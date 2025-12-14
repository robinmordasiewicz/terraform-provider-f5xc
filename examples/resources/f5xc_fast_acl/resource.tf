# Fast ACL Resource Example
# [Namespace: required] Manages `fast_acl` object, `fast_acl` object contains rules to protect site from denial of service It has destination{destination IP, destination port) and references to `fast_acl_rule` in F5 Distributed Cloud.

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
  # Object reference. This type establishes a direct referenc...
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
