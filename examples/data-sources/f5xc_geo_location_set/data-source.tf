# Geo Location Set Data Source Example
# Retrieves information about an existing Geo Location Set

# Look up an existing Geo Location Set by name
data "f5xc_geo_location_set" "example" {
  name      = "example-geo-location-set"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "geo_location_set_id" {
#   value = data.f5xc_geo_location_set.example.id
# }
