# Geo Location Set Resource Example
# [Namespace: required] Manages Geolocation Set in F5 Distributed Cloud.

# Basic Geo Location Set configuration
resource "f5xc_geo_location_set" "example" {
  name      = "example-geo-location-set"
  namespace = "shared"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Geo Location Set configuration
  country_codes = ["US", "CA", "GB"]
}
