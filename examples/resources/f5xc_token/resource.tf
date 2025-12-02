# Token Resource Example
# Manages new token. token object is used to manage site admission. User must generate token before provisioning and pass this token to site during it's registration. in F5 Distributed Cloud.

# Basic Token configuration
resource "f5xc_token" "example" {
  name      = "example-token"
  namespace = "system"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Token configuration
  token_type = "REGISTRATION_TOKEN"
}
