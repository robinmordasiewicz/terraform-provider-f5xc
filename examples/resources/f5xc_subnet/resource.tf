# Subnet Resource Example
# [Namespace: required] Manages a Subnet resource in F5 Distributed Cloud for subnet object contains configuration for an interface of a vm/pod. it is created in user or shared namespace. configuration.

# Basic Subnet configuration
resource "f5xc_subnet" "example" {
  name      = "example-subnet"
  namespace = "system"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # [OneOf: connect_to_layer2, connect_to_slo, isolated_nw] S...
  connect_to_layer2 {
    # Configure connect_to_layer2 settings
  }
  # Object reference. This type establishes a direct referenc...
  layer2_intf_ref {
    # Configure layer2_intf_ref settings
  }
  # Enable this option
  connect_to_slo {
    # Configure connect_to_slo settings
  }
}
