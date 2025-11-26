# Oidc Provider Resource Example
# Manages a OidcProvider resource in F5 Distributed Cloud for customcreatespectype is the spec to create oidc provider configuration.

# Basic Oidc Provider configuration
resource "f5xc_oidc_provider" "example" {
  name      = "example-oidc-provider"
  namespace = "staging"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # [OneOf: azure_oidc_spec_type, google_oidc_spec_type, oidc...
  azure_oidc_spec_type {
    # Configure azure_oidc_spec_type settings
  }
  # Google OIDC Spec Type. GoogleOIDCSpecType specifies the a...
  google_oidc_spec_type {
    # Configure google_oidc_spec_type settings
  }
  # OpenID Connect v1.0 Spec Type. OIDCV10SpecType specifies ...
  oidc_v10_spec_type {
    # Configure oidc_v10_spec_type settings
  }
}
