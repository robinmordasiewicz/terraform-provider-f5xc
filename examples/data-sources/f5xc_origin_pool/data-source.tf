data "f5xc_origin_pool" "existing" {
  name      = "production-backend"
  namespace = "application-services"
}

output "existing_pool_servers" {
  value = data.f5xc_origin_pool.existing.origin_servers
}

output "existing_pool_port" {
  value = data.f5xc_origin_pool.existing.port
}

output "existing_pool_lb_method" {
  value = data.f5xc_origin_pool.existing.loadbalancer_method
}
