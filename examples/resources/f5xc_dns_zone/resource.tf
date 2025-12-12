# DNS Zone Resource Example
# [Category: DNS] [Namespace: not_required] Manages a DNS Zone resource in F5 Distributed Cloud.

# Basic DNS Zone configuration
resource "f5xc_dns_zone" "example" {
  name      = "example-dns-zone"
  namespace = "staging"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # DNS Zone configuration
  # Primary DNS zone
  primary {
    soa_record_parameters {
      refresh = 86400
      retry   = 7200
      expire  = 3600000
      ttl     = 86400
      neg_ttl = 1800
    }
    default_rr_set_group {}
    default_soa_parameters {}
    dnssec_mode {
      disable {}
    }
  }
}
