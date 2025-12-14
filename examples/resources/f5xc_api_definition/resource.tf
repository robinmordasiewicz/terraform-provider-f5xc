# API Definition Resource Example
# [Category: API Security] [Namespace: required] [DependsOn: namespace] Manages API Definition. in F5 Distributed Cloud.

# Basic API Definition configuration
resource "f5xc_api_definition" "example" {
  name      = "example-api-definition"
  namespace = "shared"

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
