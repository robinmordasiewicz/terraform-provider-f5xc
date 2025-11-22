resource "f5xc_namespace" "app" {
  name        = "application-services"
  description = "Namespace for application services"
}

resource "f5xc_http_loadbalancer" "example" {
  name        = "example-lb"
  namespace   = f5xc_namespace.app.name
  description = "Example HTTP Load Balancer for web application"

  domains = [
    "www.example.com",
    "example.com"
  ]

  https      = true
  http_port  = 80
  https_port = 443

  labels = {
    environment = "production"
    app         = "web"
    managed-by  = "terraform"
  }
}

output "load_balancer_id" {
  value = f5xc_http_loadbalancer.example.id
}

output "load_balancer_domains" {
  value = f5xc_http_loadbalancer.example.domains
}
