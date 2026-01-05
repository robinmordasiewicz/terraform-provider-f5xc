# Proxy Resource Example
# Manages a Proxy resource in F5 Distributed Cloud for tcp loadbalancer create specification. configuration.

# Basic Proxy configuration
resource "f5xc_proxy" "example" {
  name      = "example-proxy"
  namespace = "staging"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Proxy configuration
  proxy_url = "http://proxy.example.com:8080"
}
