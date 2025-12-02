# Blindfold Secret Management Guide - Outputs
# These outputs provide useful information and next steps after deployment.

# -----------------------------------------------------------------------------
# Namespace Information
# -----------------------------------------------------------------------------

output "namespace" {
  description = "The namespace where resources were created"
  value       = local.namespace
}

# -----------------------------------------------------------------------------
# Certificate Outputs
# -----------------------------------------------------------------------------

output "certificate_name" {
  description = "Name of the created certificate (if enabled)"
  value       = var.enable_certificate_example ? f5xc_certificate.example[0].name : null
}

# -----------------------------------------------------------------------------
# Cloud Credentials Outputs
# -----------------------------------------------------------------------------

output "aws_credentials_name" {
  description = "Name of the AWS cloud credentials (if enabled)"
  value       = var.enable_aws_credentials_example ? f5xc_cloud_credentials.aws[0].name : null
}

output "azure_credentials_name" {
  description = "Name of the Azure cloud credentials (if enabled)"
  value       = var.enable_azure_credentials_example ? f5xc_cloud_credentials.azure[0].name : null
}

output "gcp_credentials_name" {
  description = "Name of the GCP cloud credentials (if enabled)"
  value       = var.enable_gcp_credentials_example ? f5xc_cloud_credentials.gcp[0].name : null
}

# -----------------------------------------------------------------------------
# Container Registry Outputs
# -----------------------------------------------------------------------------

output "container_registry_name" {
  description = "Name of the container registry (if enabled)"
  value       = var.enable_container_registry_example ? f5xc_container_registry.example[0].name : null
}

# -----------------------------------------------------------------------------
# Next Steps
# -----------------------------------------------------------------------------

output "next_steps" {
  description = "Helpful next steps after deployment"
  value       = <<-EOT

    ============================================================================
    BLINDFOLD SECRET MANAGEMENT - DEPLOYMENT COMPLETE
    ============================================================================

    Resources Created:
    %{if var.enable_certificate_example~}
    - Certificate: ${f5xc_certificate.example[0].name}
    %{endif~}
    %{if var.enable_aws_credentials_example~}
    - AWS Credentials: ${f5xc_cloud_credentials.aws[0].name}
    %{endif~}
    %{if var.enable_azure_credentials_example~}
    - Azure Credentials: ${f5xc_cloud_credentials.azure[0].name}
    %{endif~}
    %{if var.enable_gcp_credentials_example~}
    - GCP Credentials: ${f5xc_cloud_credentials.gcp[0].name}
    %{endif~}
    %{if var.enable_container_registry_example~}
    - Container Registry: ${f5xc_container_registry.example[0].name}
    %{endif~}

    Verify in F5XC Console:
    1. Navigate to your F5 Distributed Cloud Console
    2. Go to the appropriate section:
       - Certificates: Multi-Cloud Network Connect > Manage > Load Balancers > Certificates
       - Cloud Credentials: Cloud and Edge Sites > Manage > Site Management > Cloud Credentials
       - Container Registries: Distributed Apps > Manage > Service Discovery > Container Registries
    3. Verify your resources show encrypted secrets (not plaintext)

    Security Notes:
    - All secrets were encrypted locally before transmission
    - F5XC stores only the encrypted ciphertext
    - Only authorized F5XC services can decrypt using the SecretPolicy

    Clean Up:
    To remove all resources: terraform destroy

    ============================================================================

  EOT
}
