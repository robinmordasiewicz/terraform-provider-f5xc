# Blindfold Secret Management Guide

This example demonstrates how to use F5XC Blindfold functions to securely encrypt secrets for use with F5 Distributed Cloud resources.

## Quick Start

### Prerequisites

- Terraform >= 1.8
- F5 Distributed Cloud account
- API credentials (token or P12 certificate)

### Setup

1. **Clone the repository:**
   ```bash
   git clone https://github.com/robinmordasiewicz/terraform-provider-f5xc.git
   cd terraform-provider-f5xc/examples/guides/blindfold
   ```

2. **Configure authentication:**
   ```bash
   export VES_API_URL="https://your-tenant.console.ves.volterra.io/api"
   export VES_API_TOKEN="your-api-token"
   ```

3. **Copy and edit the configuration:**
   ```bash
   cp terraform.tfvars.example terraform.tfvars
   # Edit terraform.tfvars with your values
   ```

4. **Deploy:**
   ```bash
   terraform init
   terraform plan
   terraform apply
   ```

## Use Cases

| Example | Description | Enable Flag |
|---------|-------------|-------------|
| TLS Certificate | Encrypt private keys for certificates | `enable_certificate_example` |
| AWS Credentials | Encrypt AWS secret access keys | `enable_aws_credentials_example` |
| Azure Credentials | Encrypt Azure client secrets | `enable_azure_credentials_example` |
| GCP Credentials | Encrypt GCP service account files | `enable_gcp_credentials_example` |
| Container Registry | Encrypt registry passwords | `enable_container_registry_example` |

## Security

All secrets are encrypted locally using RSA-OAEP with SHA-256. The plaintext never leaves your machine unencrypted.

## Clean Up

```bash
terraform destroy
```

## More Information

See the full [Blindfold Secret Management Guide](../../../docs/guides/blindfold.md) for detailed documentation.
