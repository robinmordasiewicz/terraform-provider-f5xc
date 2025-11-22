data "f5xc_http_loadbalancer" "existing" {
  name      = "production-lb"
  namespace = "application-services"
}

output "existing_lb_domains" {
  value = data.f5xc_http_loadbalancer.existing.domains
}

output "existing_lb_https" {
  value = data.f5xc_http_loadbalancer.existing.https
}

output "existing_lb_description" {
  value = data.f5xc_http_loadbalancer.existing.description
}
