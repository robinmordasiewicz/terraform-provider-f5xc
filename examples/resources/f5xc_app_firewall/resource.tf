# App Firewall Resource Example
# Manages Application Firewall in F5 Distributed Cloud.

# Basic App Firewall configuration
resource "f5xc_app_firewall" "example" {
  name      = "example-app-firewall"
  namespace = "shared"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  // One of the arguments from this list "blocking monitoring" must be set

  blocking {}

  // One of the arguments from this list "custom_blocking_page use_default_blocking_page" must be set

  use_default_blocking_page {}

  // One of the arguments from this list "bot_protection_setting default_bot_setting" must be set

  bot_protection_setting {
    malicious_bot_action  = "BLOCK"
    suspicious_bot_action = "REPORT"
    good_bot_action       = "REPORT"
  }

  // One of the arguments from this list "custom_detection_settings default_detection_settings" must be set

  default_detection_settings {}

  // One of the arguments from this list "allow_all_response_codes allowed_response_codes" must be set

  allow_all_response_codes {}
}
