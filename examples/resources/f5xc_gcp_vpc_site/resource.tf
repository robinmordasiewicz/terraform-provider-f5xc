# GCP VPC Site Resource Example
# Manages a GCPVPCSite resource in F5 Distributed Cloud for deploying F5 sites within Google Cloud VPC environments.

# Basic GCP VPC Site configuration
resource "f5xc_gcp_vpc_site" "example" {
  name      = "example-gcp-vpc-site"
  namespace = "system"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # GCP VPC Site configuration
  gcp_region = "us-west1"

  # GCP credentials reference
  cloud_credentials {
    name      = "gcp-credentials"
    namespace = "system"
  }

  # Instance type
  instance_type = "n1-standard-4"

  # Ingress/Egress gateway
  ingress_egress_gw {
    gcp_certified_hw = "gcp-byol-multi-nic-voltmesh"
    node_number      = 1
    inside_network {
      new_network {
        name = "inside-network"
      }
    }
    outside_network {
      new_network {
        name = "outside-network"
      }
    }
    inside_subnet {
      new_subnet {
        subnet_name  = "inside-subnet"
        primary_ipv4 = "10.0.1.0/24"
      }
    }
    outside_subnet {
      new_subnet {
        subnet_name  = "outside-subnet"
        primary_ipv4 = "10.0.2.0/24"
      }
    }
  }

  # No worker nodes by default
  no_worker_nodes {}
}
