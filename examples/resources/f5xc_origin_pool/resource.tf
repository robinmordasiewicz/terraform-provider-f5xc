resource "f5xc_namespace" "app" {
  name        = "application-services"
  description = "Namespace for application services"
}

resource "f5xc_origin_pool" "backend" {
  name      = "backend-pool"
  namespace = f5xc_namespace.app.name

  description = "Backend server pool for application"

  origin_servers = [
    "10.1.1.10",
    "10.1.1.11",
    "10.1.1.12"
  ]

  port                 = 8080
  loadbalancer_method  = "ROUND_ROBIN"

  labels = {
    environment = "production"
    tier        = "backend"
    managed-by  = "terraform"
  }
}

output "origin_pool_id" {
  value = f5xc_origin_pool.backend.id
}

output "origin_pool_servers" {
  value = f5xc_origin_pool.backend.origin_servers
}
