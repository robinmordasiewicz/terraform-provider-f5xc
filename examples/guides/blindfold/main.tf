# Blindfold Secret Management Guide Examples
# This file demonstrates all common use cases for F5XC blindfold encryption.
#
# IMPORTANT: These examples use toggle variables to enable/disable each use case.
# Set the appropriate enable_* variable to true for the use case you want to test.

terraform {
  required_version = ">= 1.8"
  required_providers {
    f5xc = {
      source = "robinmordasiewicz/f5xc"
    }
  }
}

provider "f5xc" {
  # Authentication configured via environment variables:
  # - F5XC_API_URL (required)
  # - F5XC_API_TOKEN (token auth) OR
  # - F5XC_P12_FILE + F5XC_P12_PASSWORD (P12 cert auth)
}

# -----------------------------------------------------------------------------
# Common Configuration
# -----------------------------------------------------------------------------

locals {
  # Use created namespace or existing one based on configuration
  namespace = var.create_namespace ? f5xc_namespace.this[0].name : var.namespace_name

  # Built-in SecretPolicy that allows Volterra services to decrypt
  # This policy exists in every F5XC tenant by default
  policy_name = "ves-io-allow-volterra"
  policy_ns   = "shared"
}

# Optional: Create a dedicated namespace for testing
resource "f5xc_namespace" "this" {
  count = var.create_namespace ? 1 : 0
  name  = var.namespace_name
}

# -----------------------------------------------------------------------------
# Example 1: TLS Certificate with Private Key
# -----------------------------------------------------------------------------
# Use Case: TLS certificate storage for load balancers and other F5XC services.
#
# NOTE: TLS private keys are typically too large (>1000 bytes) for direct
# blindfold encryption, which is limited to ~190 bytes. For TLS certificates,
# use clear_secret_info which transmits the key securely over HTTPS.
#
# For secrets that fit within the size limit (API keys, passwords, tokens),
# use blindfold() or blindfold_file() as shown in the other examples.

resource "f5xc_certificate" "example" {
  count     = var.enable_certificate_example ? 1 : 0
  name      = "${var.resource_prefix}-cert"
  namespace = local.namespace

  # Certificate chain in PEM format (base64 encoded)
  certificate_url = "string:///${base64encode(file("${path.module}/certs/server.crt"))}"

  # Private key using clear_secret_info (appropriate for TLS keys)
  # The key is transmitted securely over HTTPS to F5XC
  private_key {
    clear_secret_info {
      url = "string:///${base64encode(file("${path.module}/certs/server.key"))}"
    }
  }

  # Use system defaults for OCSP stapling
  use_system_defaults {}

  labels = var.labels
}

# -----------------------------------------------------------------------------
# Example 2: AWS Cloud Credentials
# -----------------------------------------------------------------------------
# Use Case: Store AWS credentials for F5XC to access AWS services (VPC sites,
# cloud connectors, etc.). The secret access key is encrypted with blindfold.
#
# The blindfold() function takes base64-encoded plaintext and encrypts it.

resource "f5xc_cloud_credentials" "aws" {
  count     = var.enable_aws_credentials_example ? 1 : 0
  name      = "${var.resource_prefix}-aws-creds"
  namespace = "system" # Cloud credentials must be in system namespace

  aws_secret_key {
    access_key = var.aws_access_key_id

    secret_key {
      blindfold_secret_info {
        location = provider::f5xc::blindfold(
          base64encode(var.aws_secret_access_key),
          local.policy_name,
          local.policy_ns
        )
      }
    }
  }

  labels = var.labels
}

# -----------------------------------------------------------------------------
# Example 3: Azure Cloud Credentials
# -----------------------------------------------------------------------------
# Use Case: Store Azure service principal credentials for F5XC to manage
# Azure resources. The client secret is encrypted with blindfold.

resource "f5xc_cloud_credentials" "azure" {
  count     = var.enable_azure_credentials_example ? 1 : 0
  name      = "${var.resource_prefix}-azure-creds"
  namespace = "system" # Cloud credentials must be in system namespace

  azure_client_secret {
    subscription_id = var.azure_subscription_id
    tenant_id       = var.azure_tenant_id
    client_id       = var.azure_client_id

    client_secret {
      blindfold_secret_info {
        location = provider::f5xc::blindfold(
          base64encode(var.azure_client_secret),
          local.policy_name,
          local.policy_ns
        )
      }
    }
  }

  labels = var.labels
}

# -----------------------------------------------------------------------------
# Example 4: GCP Cloud Credentials
# -----------------------------------------------------------------------------
# Use Case: Store GCP service account credentials for F5XC to manage GCP
# resources. The entire service account JSON key file is encrypted.
#
# Note: GCP service account JSON files can be large. Ensure the file is under
# the RSA-OAEP size limit (~190 bytes for 2048-bit keys).

resource "f5xc_cloud_credentials" "gcp" {
  count     = var.enable_gcp_credentials_example ? 1 : 0
  name      = "${var.resource_prefix}-gcp-creds"
  namespace = "system" # Cloud credentials must be in system namespace

  gcp_cred_file {
    credential_file {
      blindfold_secret_info {
        location = provider::f5xc::blindfold_file(
          var.gcp_credentials_file,
          local.policy_name,
          local.policy_ns
        )
      }
    }
  }

  labels = var.labels
}

# -----------------------------------------------------------------------------
# Example 5: Container Registry Credentials
# -----------------------------------------------------------------------------
# Use Case: Store container registry authentication for pulling private images
# in F5XC app deployments. The password/token is encrypted with blindfold.

resource "f5xc_container_registry" "example" {
  count     = var.enable_container_registry_example ? 1 : 0
  name      = "${var.resource_prefix}-registry"
  namespace = local.namespace

  registry = var.container_registry_server

  password {
    blindfold_secret_info {
      location = provider::f5xc::blindfold(
        base64encode(var.container_registry_password),
        local.policy_name,
        local.policy_ns
      )
    }
  }

  user_name = var.container_registry_username

  labels = var.labels
}
