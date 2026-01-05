# Rate Limiter Policy Resource Example
# Manages a Rate Limiter Policy resource in F5 Distributed Cloud for rate limiter policy create specification. configuration.

# Basic Rate Limiter Policy configuration
resource "f5xc_rate_limiter_policy" "example" {
  name      = "example-rate-limiter-policy"
  namespace = "shared"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # [OneOf: any_server, server_name, server_name_matcher, ser...
  any_server {
    # Configure any_server settings
  }
  # Rules. A list of RateLimiterRules that are evaluated sequ...
  rules {
    # Configure rules settings
  }
}
