# Child Tenant Resource Example
# Manages child_tenant config instance. Name of the object is the name of the child tenant to be created. in F5 Distributed Cloud.

# Basic Child Tenant configuration
resource "f5xc_child_tenant" "example" {
  name      = "example-child-tenant"
  namespace = "staging"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # Object reference. This type establishes a direct referenc...
  child_tenant_manager {
    # Configure child_tenant_manager settings
  }
  # Contact. Instance of one single contact that can be used ...
  contact_detail {
    # Configure contact_detail settings
  }
  # Customer Info. Optional details for the new child tenant
  customer_info {
    # Configure customer_info settings
  }
}
