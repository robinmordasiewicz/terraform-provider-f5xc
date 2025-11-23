# Policer Resource Example
# Manages new policer with traffic rate limits in F5 Distributed Cloud.

# Basic Policer configuration
resource "f5xc_policer" "example" {
  name      = "example-policer"
  namespace = "system"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }
}
