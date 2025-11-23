# Cloud Elastic Ip Resource Example
# Manages Cloud Elastic IP creates Cloud Elastic IP object Object is attached to a site in F5 Distributed Cloud.

# Basic Cloud Elastic Ip configuration
resource "f5xc_cloud_elastic_ip" "example" {
  name      = "example-cloud-elastic-ip"
  namespace = "system"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # Site Reference. Site to which this cloud elastic ip objec...
  site_ref {
    # Configure site_ref settings
  }
}
