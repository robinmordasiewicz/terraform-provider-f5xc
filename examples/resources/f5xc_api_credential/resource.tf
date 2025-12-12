# API Credential Resource Example
# [Category: Authentication] [Namespace: required] Manages request specification. in F5 Distributed Cloud.

# Basic API Credential configuration
resource "f5xc_api_credential" "example" {
  name      = "example-api-credential"
  namespace = "system"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # API Credential configuration
  api_credential_type = "API_CERTIFICATE"

  # Expiration settings
  expiration_timestamp = "2025-12-31T23:59:59Z"

  # Active state
  active = true
}
