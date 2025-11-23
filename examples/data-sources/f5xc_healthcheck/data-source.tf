# Healthcheck Data Source Example
# Retrieves information about an existing Healthcheck

# Look up an existing Healthcheck by name
data "f5xc_healthcheck" "example" {
  name      = "example-healthcheck"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "healthcheck_id" {
#   value = data.f5xc_healthcheck.example.id
# }

# Example: Reference healthcheck in origin pool
# resource "f5xc_origin_pool" "example" {
#   name      = "example-pool"
#   namespace = "system"
#
#   healthcheck {
#     name      = data.f5xc_healthcheck.example.name
#     namespace = data.f5xc_healthcheck.example.namespace
#   }
# }
