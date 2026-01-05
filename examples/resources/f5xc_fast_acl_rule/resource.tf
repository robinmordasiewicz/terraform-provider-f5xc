# Fast ACL Rule Resource Example
# Manages new Fast ACL rule, `fast_acl_rule` has specification to match source IP, source port and action to apply. in F5 Distributed Cloud.

# Basic Fast ACL Rule configuration
resource "f5xc_fast_acl_rule" "example" {
  name      = "example-fast-acl-rule"
  namespace = "system"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # Action. FastAclRuleAction specifies possible action to be...
  action {
    # Configure action settings
  }
  # Policer Reference. Reference to policer object.
  policer_action {
    # Configure policer_action settings
  }
  # Reference. A policer direct reference.
  ref {
    # Configure ref settings
  }
}
