# Site Resource Example
# Manages virtual site object in given namespace. in F5 Distributed Cloud.

# Basic Site configuration
resource "f5xc_site" "example" {
  name      = "example-site"
  namespace = "staging"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # Label Selector. This type can be used to establish a 'sel...
  site_selector {
    # Configure site_selector settings
  }
}
