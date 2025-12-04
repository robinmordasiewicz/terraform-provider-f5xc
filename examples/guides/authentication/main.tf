# F5 Distributed Cloud Provider - Authentication Example
# ======================================================
#
# This example demonstrates the three authentication methods supported
# by the F5XC Terraform provider. Uncomment the method you want to use.
#
# IMPORTANT: Never commit credentials to version control!
#
# QUICK START:
# 1. Choose your authentication method (environment variables recommended)
# 2. Set the required environment variables
# 3. Run: terraform init && terraform plan

terraform {
  required_version = ">= 1.8"

  required_providers {
    f5xc = {
      source = "robinmordasiewicz/f5xc"
    }
  }
}

# =============================================================================
# AUTHENTICATION METHOD 1: Environment Variables (Recommended)
# =============================================================================
#
# This is the recommended approach - credentials are passed via environment
# variables and the provider block remains empty.
#
# For API Token:
#   export F5XC_API_URL="https://your-tenant.console.ves.volterra.io/api"
#   export F5XC_API_TOKEN="your-api-token"
#
# For P12 Certificate:
#   export F5XC_API_URL="https://your-tenant.console.ves.volterra.io/api"
#   export F5XC_API_P12_FILE="/path/to/credentials.p12"
#   export F5XC_P12_PASSWORD="your-p12-password"  # pragma: allowlist secret
#
# For PEM Certificate:
#   export F5XC_API_URL="https://your-tenant.console.ves.volterra.io/api"
#   export F5XC_API_CERT="/path/to/certificate.pem"
#   export F5XC_API_KEY="/path/to/private-key.pem"

provider "f5xc" {
  # Authentication via environment variables
  # No explicit configuration needed
}

# =============================================================================
# AUTHENTICATION METHOD 2: Provider Configuration with Variables
# =============================================================================
#
# Uncomment this block and comment out the empty provider block above if you
# prefer explicit configuration. The actual values should come from
# variables (see variables.tf) populated via terraform.tfvars or TF_VAR_
# environment variables.
#
# provider "f5xc" {
#   api_url   = var.f5xc_api_url
#   api_token = var.f5xc_api_token
# }

# =============================================================================
# AUTHENTICATION METHOD 3: P12 Certificate via Variables
# =============================================================================
#
# For P12 certificate authentication with explicit configuration:
#
# provider "f5xc" {
#   api_url      = var.f5xc_api_url
#   api_p12_file = var.f5xc_api_p12_file
#   p12_password = var.f5xc_p12_password
# }

# =============================================================================
# AUTHENTICATION METHOD 4: PEM Certificate via Variables
# =============================================================================
#
# For PEM certificate authentication (extracted from P12):
#
# provider "f5xc" {
#   api_url  = var.f5xc_api_url
#   api_cert = var.f5xc_api_cert
#   api_key  = var.f5xc_api_key
# }

# =============================================================================
# Test Resource - Validates Authentication
# =============================================================================
#
# This data source validates that authentication is working correctly.
# It retrieves information about the "system" namespace which always exists.

data "f5xc_namespace" "system" {
  name = "system"
}

# Output the namespace to confirm authentication worked
output "authentication_test" {
  description = "Authentication successful - retrieved system namespace"
  value = {
    namespace   = data.f5xc_namespace.system.name
    description = data.f5xc_namespace.system.description
  }
}
