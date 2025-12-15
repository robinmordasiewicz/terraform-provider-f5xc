# WAF Exclusion Policy Resource Example
# Manages WAF exclusion policy in F5 Distributed Cloud.

# Basic WAF Exclusion Policy configuration
resource "f5xc_waf_exclusion_policy" "example" {
  name      = "example-waf-exclusion-policy"
  namespace = "shared"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # WAF Exclusion Rules. An ordered list of rules.
  waf_exclusion_rules {
    # Configure waf_exclusion_rules settings
  }
  # Enable this option
  any_domain {
    # Configure any_domain settings
  }
  # Enable this option
  any_path {
    # Configure any_path settings
  }
}
