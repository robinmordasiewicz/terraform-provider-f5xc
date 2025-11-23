# Proxy Resource Example
# Manages a Proxy resource in F5 Distributed Cloud for tcp loadbalancer create configuration.

# Basic Proxy configuration
resource "f5xc_proxy" "example" {
  name      = "example-proxy"
  namespace = "system"

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
