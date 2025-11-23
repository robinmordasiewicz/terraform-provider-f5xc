# Child Tenant Manager Resource Example
# Manages child_tenant_manager config instance. Name of the object is the name of the child tenant manager to be created. in F5 Distributed Cloud.

# Basic Child Tenant Manager configuration
resource "f5xc_child_tenant_manager" "example" {
  name      = "example-child-tenant-manager"
  namespace = "system"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
    # Group Mapping. The Group Mapping field is used to associa...
    group_assignments {
      # Configure group_assignments settings
    }
    # Object reference. This type establishes a direct referenc...
    group {
      # Configure group settings
    }
    # Object reference. This type establishes a direct referenc...
    tenant_owner_group {
      # Configure tenant_owner_group settings
    }
}
