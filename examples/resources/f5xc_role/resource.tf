# Role Resource Example
# Manages role in F5 Distributed Cloud.

# Basic Role configuration
resource "f5xc_role" "example" {
  name      = "example-role"
  namespace = "system"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Role configuration
  role_type = "CUSTOM"

  # API groups for this role
  api_groups {
    name = "read-http-lb"
    api_group_elements {
      api_group     = "ves.io.schema.views.http_loadbalancer"
      resource_type = "http_loadbalancer"
      verbs {
        get  = true
        list = true
      }
    }
  }
}
