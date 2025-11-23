# Fleet Data Source Example
# Retrieves information about an existing Fleet

# Look up an existing Fleet by name
data "f5xc_fleet" "example" {
  name      = "example-fleet"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "fleet_id" {
#   value = data.f5xc_fleet.example.id
# }
