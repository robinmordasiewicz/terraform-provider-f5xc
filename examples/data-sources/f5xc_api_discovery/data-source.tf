# API Discovery Data Source Example
# Retrieves information about an existing API Discovery

# Look up an existing API Discovery by name
data "f5xc_api_discovery" "example" {
  name      = "example-api-discovery"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "api_discovery_id" {
#   value = data.f5xc_api_discovery.example.id
# }
