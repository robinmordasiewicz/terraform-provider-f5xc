# App Setting Resource Example
# Manages App setting configuration in namespace metadata.namespace in F5 Distributed Cloud.

# Basic App Setting configuration
resource "f5xc_app_setting" "example" {
  name      = "example-app-setting"
  namespace = "system"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
    # Customize AppType For This Namespace. List of settings to...
    app_type_settings {
      # Configure app_type_settings settings
    }
    # AppType. The AppType of App instance in current Namespace...
    app_type_ref {
      # Configure app_type_ref settings
    }
    # API Discovery. Settings specifying how API Discovery will...
    business_logic_markup_setting {
      # Configure business_logic_markup_setting settings
    }
}
