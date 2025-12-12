---
page_title: "Guide: Authentication Methods"
subcategory: "Guides"
description: |-
  Comprehensive guide to authenticating the F5 Distributed Cloud Terraform
  provider using API tokens, P12 certificates, or PEM certificates.
---

# Authentication Methods

This guide walks you through configuring authentication for the F5 Distributed Cloud Terraform provider. By the end, you'll understand:

- **Three authentication methods** - API Token, P12 Certificate, and PEM Certificate
- **Configuration approaches** - Environment variables, provider blocks, and tfvars files
- **CI/CD integration** - GitHub Actions workflows with secure credential handling
- **Security best practices** - Protecting credentials in different environments

## Prerequisites

Before you begin, ensure you have:

- **F5 Distributed Cloud Account** - Sign up at <https://www.f5.com/cloud/products/distributed-cloud-console>
- **Terraform >= 1.8** - Download and install from <https://www.terraform.io/downloads>
- **Console Access** - Ability to create credentials in the F5XC Console

## Creating Credentials in F5 Distributed Cloud

### Personal Credentials

Personal credentials are tied to your user account and are ideal for development and testing.

#### Creating an API Token

1. Open the F5 Distributed Cloud Console homepage
2. Select **Administration** from the main menu
3. Select **Personal Management** → **Credentials**
4. Click **+ Add Credentials**
5. Enter a descriptive name (e.g., "terraform-dev-token")
6. Select **API Token** from the "Credential type" dropdown
7. Choose an expiry date from the calendar
8. Click **Generate** to create the token
9. Click **Copy** to copy the token value

!> **Warning:** Copy and save your token immediately. You cannot retrieve it after closing this dialog.

10. Click **Done** to exit

#### Creating a P12 Certificate

1. Navigate to **Administration** → **Personal Management** → **Credentials**
2. Click **+ Add Credentials**
3. Enter a certificate name (e.g., "terraform-dev-cert")
4. Select **API Certificate** from the "Credential type" dropdown
5. Enter and confirm a password (you'll need this password to use the certificate)
6. Select an expiry date from the calendar
7. Click **Download** to generate and download the `.p12` file

-> **Tip:** Store the P12 file securely and never commit it to version control.

### Service Credentials (IAM)

Service credentials are managed through IAM and recommended for production and CI/CD pipelines. They can be scoped to specific roles and namespaces.

#### Creating a Service API Token

1. Navigate to **Administration** → **IAM** → **Service Credentials**
2. Click **+ Add Service Credentials**
3. Enter a descriptive name (e.g., "terraform-cicd-token")
4. Select **API Token** from the "Credential type" dropdown
5. Optionally assign the credential to groups
6. Optionally assign specific roles and namespaces to limit scope
7. Choose an expiry date
8. Click **Generate**, then **Copy** the token value
9. Click **Done** to exit

#### Creating a Service P12 Certificate

1. Navigate to **Administration** → **IAM** → **Service Credentials**
2. Click **+ Add Service Credentials**
3. Enter a certificate name
4. Select **API Certificate** from the "Credential type" dropdown
5. Optionally assign to groups, roles, and namespaces
6. Enter and confirm a password
7. Select an expiry date
8. Click **Download** to generate the `.p12` file

## Authentication Methods

The provider supports three authentication methods. Choose based on your security requirements and operational context.

### Method 1: P12 Certificate Authentication (Recommended for Production)

P12 certificates provide mutual TLS (mTLS) authentication, where both the client and server verify each other's identity.

**Using Environment Variables:**

```bash
export VES_API_URL="https://your-tenant.console.ves.volterra.io/api"
export VES_P12_FILE="/path/to/your-credentials.p12"
export VES_P12_PASSWORD="your-p12-password"  # pragma: allowlist secret
```

**Using Provider Configuration:**

```hcl
provider "f5xc" {
  api_url      = var.f5xc_api_url
  api_p12_file = var.f5xc_api_p12_file
  p12_password = var.f5xc_p12_password
}
```

### Method 2: PEM Certificate Authentication

PEM certificate authentication uses separate certificate and private key files extracted from a P12 file. This is useful when your tooling prefers PEM format.

**Extracting PEM from P12:**

```bash
# Create a directory for certificates
mkdir -p certs

# Extract the certificate
openssl pkcs12 -in ~/your-tenant.console.ves.volterra.io.api-creds.p12 \
  -nodes -nokeys -out certs/f5xc.cert
# Enter Import Password: <your p12 password>

# Extract the private key
openssl pkcs12 -in ~/your-tenant.console.ves.volterra.io.api-creds.p12 \
  -nodes -nocerts -out certs/f5xc.key
# Enter Import Password: <your p12 password>
```

**Using Environment Variables:**

```bash
export VES_API_URL="https://your-tenant.console.ves.volterra.io/api"
export VES_CERT="/path/to/certs/f5xc.cert"
export VES_KEY="/path/to/certs/f5xc.key"
```

**Using Provider Configuration:**

```hcl
provider "f5xc" {
  api_url  = var.f5xc_api_url
  api_cert = var.f5xc_api_cert
  api_key  = var.f5xc_api_key
}
```

~> **Note:** If you need to verify the server certificate, you can also specify a CA certificate using `VES_CACERT` environment variable or `api_ca_cert` provider attribute.

### Method 3: API Token Authentication (Simplest)

API tokens provide the simplest authentication method using bearer token authentication over TLS.

**Using Environment Variables:**

```bash
export VES_API_URL="https://your-tenant.console.ves.volterra.io/api"
export VES_API_TOKEN="your-api-token"
```

**Using Provider Configuration:**

```hcl
provider "f5xc" {
  api_url   = var.f5xc_api_url
  api_token = var.f5xc_api_token
}
```

## Configuration Approaches

### Environment Variables (Recommended)

Environment variables are the recommended approach because they:

- Keep credentials out of configuration files
- Work consistently across local development and CI/CD
- Are easy to rotate without changing code
- Never risk being committed to version control

**Complete Environment Variable Reference:**

| Variable | Description | Required |
|----------|-------------|----------|
| `VES_API_URL` | F5XC tenant API URL | Yes |
| `VES_API_TOKEN` | API token for bearer authentication | One of: token, P12, or PEM |
| `VES_P12_FILE` | Path to P12 certificate file | With `VES_P12_PASSWORD` |
| `VES_P12_PASSWORD` | Password for P12 file | With `VES_P12_FILE` |
| `VES_CERT` | Path to PEM certificate file | With `VES_KEY` |
| `VES_KEY` | Path to PEM private key file | With `VES_CERT` |
| `VES_CACERT` | Path to CA certificate for server verification | No |

**Adding to Shell Profile:**

For persistence across terminal sessions, add exports to your shell profile:

```bash
# Add to ~/.bashrc or ~/.zshrc
export VES_API_URL="https://your-tenant.console.ves.volterra.io/api"
export VES_API_TOKEN="your-api-token"
```

Then reload your shell: `source ~/.zshrc` or `source ~/.bashrc`

### Provider Configuration Block

When you need explicit configuration or want to use Terraform variables:

```hcl
# Using environment variables (recommended)
provider "f5xc" {}

# Explicit configuration with variables
provider "f5xc" {
  api_url   = var.f5xc_api_url
  api_token = var.f5xc_api_token  # Marked sensitive in variables.tf
}
```

### Using terraform.tfvars

You can define variable values in a `terraform.tfvars` file, but **never commit files containing credentials**:

```hcl
# terraform.tfvars - DO NOT COMMIT THIS FILE
f5xc_api_url   = "https://your-tenant.console.ves.volterra.io/api"
f5xc_api_token = "your-api-token"
```

!> **Warning:** Add `*.tfvars` to your `.gitignore` if it contains credentials. Better yet, use environment variables for sensitive values.

### Authentication Priority

When multiple authentication methods are configured, the provider uses this priority:

1. **P12 Certificate** - If `api_p12_file` is set (requires `p12_password`)
2. **PEM Certificate** - If both `api_cert` and `api_key` are set
3. **API Token** - If `api_token` is set
4. **Error** - If none of the above are provided

## CI/CD Integration

### GitHub Actions with API Token

This is the simplest approach for GitHub Actions:

```yaml
name: Terraform Deploy

on:
  push:
    branches: [main]
  pull_request:
    branches: [main]

jobs:
  terraform:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout code
        uses: actions/checkout@v4

      - name: Setup Terraform
        uses: hashicorp/setup-terraform@v3
        with:
          terraform_version: "1.8.0"

      - name: Terraform Init
        run: terraform init

      - name: Terraform Plan
        env:
          VES_API_URL: ${{ secrets.VES_API_URL }}
          VES_API_TOKEN: ${{ secrets.VES_API_TOKEN }}
        run: terraform plan -out=tfplan

      - name: Terraform Apply
        if: github.ref == 'refs/heads/main' && github.event_name == 'push'
        env:
          VES_API_URL: ${{ secrets.VES_API_URL }}
          VES_API_TOKEN: ${{ secrets.VES_API_TOKEN }}
        run: terraform apply -auto-approve tfplan
```

**Setting Up GitHub Secrets:**

1. Navigate to your repository on GitHub
2. Go to **Settings** → **Secrets and variables** → **Actions**
3. Click **New repository secret**
4. Add the following secrets:

| Secret Name | Value |
|-------------|-------|
| `VES_API_URL` | `https://your-tenant.console.ves.volterra.io/api` |
| `VES_API_TOKEN` | Your API token value |

### GitHub Actions with P12 Certificate (More Secure)

For production deployments requiring mTLS:

```yaml
name: Terraform Deploy with P12 Certificate

on:
  push:
    branches: [main]
  pull_request:
    branches: [main]

jobs:
  terraform:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout code
        uses: actions/checkout@v4

      - name: Setup Terraform
        uses: hashicorp/setup-terraform@v3
        with:
          terraform_version: "1.8.0"

      - name: Setup P12 Certificate
        run: |
          echo "${{ secrets.VES_P12_BASE64 }}" | base64 -d > /tmp/f5xc-credentials.p12
          chmod 600 /tmp/f5xc-credentials.p12

      - name: Terraform Init
        run: terraform init

      - name: Terraform Plan
        env:
          VES_API_URL: ${{ secrets.VES_API_URL }}
          VES_P12_FILE: /tmp/f5xc-credentials.p12
          VES_P12_PASSWORD: ${{ secrets.VES_P12_PASSWORD }}
        run: terraform plan -out=tfplan

      - name: Terraform Apply
        if: github.ref == 'refs/heads/main' && github.event_name == 'push'
        env:
          VES_API_URL: ${{ secrets.VES_API_URL }}
          VES_P12_FILE: /tmp/f5xc-credentials.p12
          VES_P12_PASSWORD: ${{ secrets.VES_P12_PASSWORD }}
        run: terraform apply -auto-approve tfplan

      - name: Cleanup Credentials
        if: always()
        run: rm -f /tmp/f5xc-credentials.p12
```

**Encoding P12 for GitHub Secrets:**

Before adding the P12 file as a secret, encode it as base64:

```bash
# On macOS
base64 -i your-credentials.p12 | pbcopy
# The base64 string is now in your clipboard

# On Linux
base64 -w 0 your-credentials.p12
# Copy the output
```

**Setting Up GitHub Secrets for P12:**

| Secret Name | Value |
|-------------|-------|
| `VES_API_URL` | `https://your-tenant.console.ves.volterra.io/api` |
| `VES_P12_BASE64` | Base64-encoded P12 file contents |
| `VES_P12_PASSWORD` | Password for the P12 file |

## Security Best Practices

### Credential Protection

- **Never commit credentials** to version control. Add `*.tfvars`, `*.p12`, and `*.pem` to `.gitignore`
- **Use environment variables** for sensitive values in local development
- **Use GitHub Secrets** or equivalent for CI/CD pipelines
- **Limit credential scope** using Service Credentials with specific roles and namespaces

### Choosing the Right Method

| Use Case | Recommended Method | Reason |
|----------|-------------------|--------|
| Local development | API Token | Simplest setup |
| CI/CD pipelines | P12 Certificate | mTLS security |
| Production automation | Service Credentials | Role-based access control |
| Quick testing | API Token | Fast iteration |

### Credential Expiry

When creating credentials, consider:

- **Development credentials**: Shorter expiry (30-90 days)
- **CI/CD credentials**: Align with your rotation schedule
- **Use Service Credentials**: Easier to manage across teams

## Troubleshooting

### Authentication Failed (401 Unauthorized)

**Symptom:** Terraform returns 401 errors.

**Solutions:**

1. Verify your API URL is correct (includes `/api` suffix)
2. Check that your token hasn't expired
3. Verify the token was copied correctly (no extra whitespace)
4. Ensure environment variables are set in the current shell

```bash
# Verify environment variables are set
echo $VES_API_URL
echo $VES_API_TOKEN
```

### Certificate Verification Failed

**Symptom:** TLS/SSL certificate errors.

**Solutions:**

1. Verify the P12 password is correct
2. Check that the P12 file path is absolute or correctly relative
3. Ensure the P12 file hasn't been corrupted during transfer

```bash
# Test P12 file with OpenSSL
openssl pkcs12 -in your-credentials.p12 -nokeys -info
```

### Permission Denied

**Symptom:** 403 Forbidden errors on specific operations.

**Solutions:**

1. Verify your credential has the required permissions
2. For Service Credentials, check role and namespace assignments
3. Some operations require specific system roles

### Environment Variables Not Working

**Symptom:** Provider asks for credentials despite environment variables being set.

**Solutions:**

1. Ensure variables are exported (not just set)

   ```bash
   export VES_API_TOKEN="token"  # Correct
   VES_API_TOKEN="token"         # Won't work with Terraform
   ```

2. Verify spelling matches exactly (case-sensitive)
3. Check for hidden characters or quotes in the value

## Clean Up

To revoke credentials when no longer needed:

1. Navigate to **Administration** → **Personal Management** → **Credentials** (or **IAM** → **Service Credentials**)
2. Find the credential in the list
3. Click the **Actions** menu (three dots)
4. Select **Force Expiry** to immediately invalidate the credential

## Next Steps

Now that you have authentication configured, explore:

- [HTTP Load Balancer Guide](http-loadbalancer) - Deploy your first load balancer
- [Blindfold Functions Guide](blindfold) - Secure secret management
- [Namespace Resource](../resources/namespace) - Organize your resources
- [Origin Pool Resource](../resources/origin_pool) - Configure backend servers

## Support

- **Provider Documentation:** [F5XC Provider](../index)
- **F5 Documentation:** [F5 Distributed Cloud Docs](https://docs.cloud.f5.com/)
- **Credential Management:** [F5 Credentials Guide](https://docs.cloud.f5.com/docs-v2/administration/how-tos/user-mgmt/Credentials)
- **Issues:** [GitHub Issues](https://github.com/robinmordasiewicz/terraform-provider-f5xc/issues)
