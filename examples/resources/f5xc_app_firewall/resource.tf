# App Firewall Resource Example
# Manages Application Firewall in F5 Distributed Cloud.

# Basic App Firewall configuration
resource "f5xc_app_firewall" "example" {
  name      = "example-app-firewall"
  namespace = "system"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Web Application Firewall configuration
  # Block malicious requests
  blocking {}

  # Use default detection settings
  use_default_blocking_page {}

  # Default bot defense configuration
  default_bot_setting {}

  # Default detection settings
  default_detection_settings {}

  # Allow all response codes
  allow_all_response_codes {}
}
