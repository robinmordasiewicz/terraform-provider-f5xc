# Service Policy Data Source Example
# Retrieves information about an existing Service Policy

# Look up an existing Service Policy by name
data "f5xc_service_policy" "example" {
  name      = "example-service-policy"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "service_policy_id" {
#   value = data.f5xc_service_policy.example.id
# }

# Example: Reference service policy in HTTP load balancer
# resource "f5xc_http_loadbalancer" "example" {
#   name      = "policy-protected-lb"
#   namespace = "system"
#
#   active_service_policies {
#     policies {
#       name      = data.f5xc_service_policy.example.name
#       namespace = data.f5xc_service_policy.example.namespace
#     }
#   }
# }
