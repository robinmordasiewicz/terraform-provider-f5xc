# Workload Resource Example
# Manages workload_flavor. in F5 Distributed Cloud.

# Basic Workload configuration
resource "f5xc_workload" "example" {
  name      = "example-workload"
  namespace = "staging"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Workload configuration
  # Container configuration
  containers {
    name = "web"
    image {
      name        = "nginx"
      public      = {}
      pull_policy = "IMAGE_PULL_POLICY_ALWAYS"
    }
  }

  # Deploy on regional edge
  deploy_on_re {
    virtual_site {
      name      = "example-virtual-site"
      namespace = "system"
    }
  }
}
