# Healthcheck Resource Example
# Manages a Healthcheck resource in F5 Distributed Cloud for healthcheck object defines method to determine if the given endpoint is healthy. single healthcheck object can be referred to by one or many cluster objects. configuration.

# Basic Healthcheck configuration
resource "f5xc_healthcheck" "example" {
  name      = "example-healthcheck"
  namespace = "staging"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  // One of the arguments from this list "http_health_check tcp_health_check udp_icmp_health_check" must be set

  http_health_check {
    // One of the arguments from this list "host_header use_origin_server_name" must be set

    use_origin_server_name {}

    path                  = "/health"
    use_http2             = false
    expected_status_codes = ["200"]

    // One of the arguments from this list "headers request_headers_to_remove" must be set

    headers = {
      "x-health-check" = "true"
    }
  }

  healthy_threshold   = 3
  unhealthy_threshold = 3
  interval            = 15
  timeout             = 5
}
