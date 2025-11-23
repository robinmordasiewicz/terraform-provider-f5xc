# Healthcheck Resource Example
# Manages a Healthcheck resource in F5 Distributed Cloud for healthcheck object defines method to determine if the given endpoint is healthy. single healthcheck object can be referred to by one or many cluster objects. configuration.

# Basic Healthcheck configuration
resource "f5xc_healthcheck" "example" {
  name      = "example-healthcheck"
  namespace = "system"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Health Check specific configuration
  http_health_check {
    use_origin_server_name {}
    path                  = "/health"
    use_http2             = false
    expected_status_codes = ["200"]
  }

  healthy_threshold   = 3
  unhealthy_threshold = 3
  interval            = 15
  timeout             = 5
}
