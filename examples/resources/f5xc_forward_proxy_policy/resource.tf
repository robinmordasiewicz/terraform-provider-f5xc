# Forward Proxy Policy Resource Example
# Manages a Forward Proxy Policy resource in F5 Distributed Cloud for forward proxy policy configuration.

# Basic Forward Proxy Policy configuration
resource "f5xc_forward_proxy_policy" "example" {
  name      = "example-forward-proxy-policy"
  namespace = "shared"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Forward Proxy Policy configuration
  proxy_label_selector {
    expressions = ["app in (web, api)"]
  }

  drp_http_connect {
    any_proxy {}
    rule_list {
      rules {
        metadata {
          name = "allow-external"
        }
        spec {
          action = "ALLOW"
          dst_list {
            any_dst {}
          }
        }
      }
    }
  }
}
