# Dns Zone Data Source Example
# Retrieves information about an existing Dns Zone

# Look up an existing Dns Zone by name
data "f5xc_dns_zone" "example" {
  name      = "example-dns-zone"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "dns_zone_id" {
#   value = data.f5xc_dns_zone.example.id
# }

# Example: Reference DNS zone in DNS load balancer
# resource "f5xc_dns_load_balancer" "example" {
#   name      = "example-dns-lb"
#   namespace = "system"
#
#   dns_zone {
#     name      = data.f5xc_dns_zone.example.name
#     namespace = data.f5xc_dns_zone.example.namespace
#   }
# }
