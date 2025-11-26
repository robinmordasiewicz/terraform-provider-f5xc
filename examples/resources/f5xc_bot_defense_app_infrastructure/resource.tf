# Bot Defense App Infrastructure Resource Example
# Manages Bot Defense App Infrastructure in a given namespace. in F5 Distributed Cloud.

# Basic Bot Defense App Infrastructure configuration
resource "f5xc_bot_defense_app_infrastructure" "example" {
  name      = "example-bot-defense-app-infrastructure"
  namespace = "shared"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # [OneOf: cloud_hosted, data_center_hosted] F5 Hosted. Infr...
  cloud_hosted {
    # Configure cloud_hosted settings
  }
  # Egress. Egress
  egress {
    # Configure egress settings
  }
  # Ingress. Ingress
  ingress {
    # Configure ingress settings
  }
}
