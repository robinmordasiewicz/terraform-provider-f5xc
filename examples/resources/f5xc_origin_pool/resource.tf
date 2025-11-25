# Origin Pool Resource Example
# Manages a OriginPool resource in F5 Distributed Cloud for defining backend server pools for load balancer targets.

# Basic Origin Pool configuration
resource "f5xc_origin_pool" "example" {
  name      = "example-origin-pool"
  namespace = "system"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  // Origin servers configuration
  origin_servers {
    // One of the arguments from this list "consul_service custom_endpoint_object k8s_service private_ip private_name public_ip public_name vn_private_ip vn_private_name" must be set

    public_name {
      dns_name         = "origin.example.com"
      refresh_interval = 60
    }

    labels = {
      "app" = "backend"
    }
  }

  origin_servers {
    // One of the arguments from this list "consul_service custom_endpoint_object k8s_service private_ip private_name public_ip public_name vn_private_ip vn_private_name" must be set

    k8s_service {
      service_name = "backend-svc"

      // One of the arguments from this list "inside_network outside_network vk8s_networks" must be set

      vk8s_networks {}

      site_locator {
        // One of the arguments from this list "site virtual_site" must be set

        site {
          name      = "example-site"
          namespace = "system"
        }
      }
    }
  }

  port = 443

  // One of the arguments from this list "no_tls use_tls" must be set

  use_tls {
    // One of the arguments from this list "disable_sni sni use_host_header_as_sni" must be set

    sni = "backend.example.com"

    tls_config {
      // One of the arguments from this list "custom_security default_security low_security medium_security" must be set

      default_security {}
    }

    // One of the arguments from this list "no_mtls use_mtls use_mtls_obj" must be set

    no_mtls {}

    // One of the arguments from this list "skip_server_verification use_server_verification volterra_trusted_ca" must be set

    volterra_trusted_ca {}
  }

  // Health check configuration
  healthcheck {
    name      = "example-healthcheck"
    namespace = "system"
  }

  // Load balancing configuration
  endpoint_selection     = "LOCAL_PREFERRED"
  loadbalancer_algorithm = "ROUND_ROBIN"
}
