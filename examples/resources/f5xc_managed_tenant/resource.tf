# Managed Tenant Resource Example
# Manages managed_tenant config instance. Name of the object is name of the tenant that is allowed to manage. in F5 Distributed Cloud.

# Basic Managed Tenant configuration
resource "f5xc_managed_tenant" "example" {
  name      = "example-managed-tenant"
  namespace = "system"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
    # Group Mapping. List of local user group association to us...
    groups {
      # Configure groups settings
    }
    # Object reference. This type establishes a direct referenc...
    group {
      # Configure group settings
    }
}
