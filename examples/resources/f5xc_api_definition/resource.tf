# Api Definition Resource Example
# Manages a APIDefinition resource in F5 Distributed Cloud for x-required create api definition. configuration.

# Basic Api Definition configuration
resource "f5xc_api_definition" "example" {
  name      = "example-api-definition"
  namespace = "system"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # API Definition configuration
  # OpenAPI spec
  swagger_specs = ["string:///base64-openapi-spec"]

  # Non-validation mode
  non_validation_mode {}
}
