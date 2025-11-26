# Cdn Cache Rule Resource Example
# Manages a CDNCacheRule resource in F5 Distributed Cloud for cdn loadbalancer configuration.

# Basic Cdn Cache Rule configuration
resource "f5xc_cdn_cache_rule" "example" {
  name      = "example-cdn-cache-rule"
  namespace = "staging"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # Cache Rule. This defines a CDN Cache Rule
  cache_rules {
    # Configure cache_rules settings
  }
  # Empty. This can be used for messages where no values are ...
  cache_bypass {
    # Configure cache_bypass settings
  }
  # Cache Action Options. List of options for Cache Action
  eligible_for_cache {
    # Configure eligible_for_cache settings
  }
}
