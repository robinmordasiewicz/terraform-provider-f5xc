# Blindfold Secret Management Guide - Variables
# Configure these variables to test different blindfold use cases.

# -----------------------------------------------------------------------------
# General Configuration
# -----------------------------------------------------------------------------

variable "namespace_name" {
  description = "Name of the namespace to use or create for resources"
  type        = string
  default     = "blindfold-guide"
}

variable "create_namespace" {
  description = "Set to true to create the namespace, false to use an existing one"
  type        = bool
  default     = true
}

variable "resource_prefix" {
  description = "Prefix for all resource names to ensure uniqueness"
  type        = string
  default     = "blindfold-example"
}

variable "labels" {
  description = "Labels to apply to all resources"
  type        = map(string)
  default = {
    managed_by = "terraform"
    guide      = "blindfold"
  }
}

# -----------------------------------------------------------------------------
# Example Toggle Flags
# -----------------------------------------------------------------------------
# Enable only the examples you want to test. Each example requires its own
# set of credentials or files to be provided.

variable "enable_certificate_example" {
  description = "Enable the TLS certificate example (requires server.crt and server.key in certs/)"
  type        = bool
  default     = false
}

variable "enable_aws_credentials_example" {
  description = "Enable the AWS credentials example (requires AWS access key and secret)"
  type        = bool
  default     = false
}

variable "enable_azure_credentials_example" {
  description = "Enable the Azure credentials example (requires Azure service principal)"
  type        = bool
  default     = false
}

variable "enable_gcp_credentials_example" {
  description = "Enable the GCP credentials example (requires GCP service account JSON file)"
  type        = bool
  default     = false
}

variable "enable_container_registry_example" {
  description = "Enable the container registry example (requires registry credentials)"
  type        = bool
  default     = false
}

# -----------------------------------------------------------------------------
# AWS Credentials (Example 2)
# -----------------------------------------------------------------------------

variable "aws_access_key_id" {
  description = "AWS access key ID for cloud credentials"
  type        = string
  default     = ""
  sensitive   = false # Access key ID is not secret
}

variable "aws_secret_access_key" {
  description = "AWS secret access key (will be encrypted with blindfold)"
  type        = string
  default     = ""
  sensitive   = true
}

# -----------------------------------------------------------------------------
# Azure Credentials (Example 3)
# -----------------------------------------------------------------------------

variable "azure_subscription_id" {
  description = "Azure subscription ID"
  type        = string
  default     = ""
}

variable "azure_tenant_id" {
  description = "Azure Active Directory tenant ID"
  type        = string
  default     = ""
}

variable "azure_client_id" {
  description = "Azure service principal client (application) ID"
  type        = string
  default     = ""
}

variable "azure_client_secret" {
  description = "Azure service principal client secret (will be encrypted with blindfold)"
  type        = string
  default     = ""
  sensitive   = true
}

# -----------------------------------------------------------------------------
# GCP Credentials (Example 4)
# -----------------------------------------------------------------------------

variable "gcp_credentials_file" {
  description = "Path to GCP service account JSON key file (will be encrypted with blindfold)"
  type        = string
  default     = ""
}

# -----------------------------------------------------------------------------
# Container Registry Credentials (Example 5)
# -----------------------------------------------------------------------------

variable "container_registry_server" {
  description = "Container registry server URL (e.g., docker.io, gcr.io, ghcr.io)"
  type        = string
  default     = "docker.io"
}

variable "container_registry_username" {
  description = "Container registry username"
  type        = string
  default     = ""
}

variable "container_registry_password" {
  description = "Container registry password or access token (will be encrypted with blindfold)"
  type        = string
  default     = ""
  sensitive   = true
}
