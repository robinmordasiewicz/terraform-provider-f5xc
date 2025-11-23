# Service Policy Rule Resource Example
# Manages service_policy_rule creates a new object in the storage backend for metadata.namespace. in F5 Distributed Cloud.

# Basic Service Policy Rule configuration
resource "f5xc_service_policy_rule" "example" {
  name      = "example-service-policy-rule"
  namespace = "system"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
    # [OneOf: any_asn, asn_list, asn_matcher] Empty. This can b...
    any_asn {
      # Configure any_asn settings
    }
    # [OneOf: any_client, client_name, client_name_matcher, cli...
    any_client {
      # Configure any_client settings
    }
    # [OneOf: any_ip, ip_matcher, ip_prefix_list] Empty. This c...
    any_ip {
      # Configure any_ip settings
    }
}
