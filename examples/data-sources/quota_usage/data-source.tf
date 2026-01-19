# Query quota usage for the system namespace
data "f5xc_quota_usage" "system" {
  namespace = "system"
}

# Output quota information for healthcheck resources
output "healthcheck_quota" {
  description = "Current healthcheck quota usage in system namespace"
  value = {
    used      = data.f5xc_quota_usage.system.objects["healthcheck"].usage
    limit     = data.f5xc_quota_usage.system.objects["healthcheck"].limit
    available = data.f5xc_quota_usage.system.objects["healthcheck"].available
  }
}

# Output quota information for HTTP load balancers
output "http_loadbalancer_quota" {
  description = "Current HTTP load balancer quota usage in system namespace"
  value = {
    used      = data.f5xc_quota_usage.system.objects["http_loadbalancer"].usage
    limit     = data.f5xc_quota_usage.system.objects["http_loadbalancer"].limit
    available = data.f5xc_quota_usage.system.objects["http_loadbalancer"].available
  }
}

# Example: Use quota data with lifecycle precondition to prevent creation
# when quota is exhausted
resource "f5xc_healthcheck" "example" {
  name      = "example-healthcheck"
  namespace = "system"

  http_health_check {
    path = "/health"
  }

  lifecycle {
    precondition {
      condition     = data.f5xc_quota_usage.system.objects["healthcheck"].available > 0
      error_message = "Cannot create healthcheck: quota exhausted in system namespace (${data.f5xc_quota_usage.system.objects["healthcheck"].usage}/${data.f5xc_quota_usage.system.objects["healthcheck"].limit} used)"
    }
  }
}
