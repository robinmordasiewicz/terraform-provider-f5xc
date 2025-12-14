# DNS Domain Resource Example
# [Category: DNS] [Namespace: not_required] Manages DNS Domain in a given namespace. If one already exist it will give a error. in F5 Distributed Cloud.

# Basic DNS Domain configuration
resource "f5xc_dns_domain" "example" {
  name      = "example-dns-domain"
  namespace = "staging"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # Enable this option
  volterra_managed {
    # Configure volterra_managed settings
  }
}
