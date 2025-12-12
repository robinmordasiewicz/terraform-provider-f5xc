---
page_title: "Guide: Authentication Methods"
subcategory: "Guides"
description: |-
  Comprehensive guide to authenticating the F5 Distributed Cloud Terraform
  provider using API tokens, P12 certificates, or PEM certificates.
---

# Authentication Methods

This guide covers authentication configuration for the F5 Distributed Cloud Terraform provider.

## Quick Reference

| Method | Complexity | Best For | Security |
|--------|------------|----------|----------|
| API Token | Simplest | Development, quick testing | Bearer token over TLS |
| P12 Certificate | Moderate | Production, CI/CD | Mutual TLS (mTLS) |
| PEM Certificate | Advanced | When tooling requires PEM format | Derived from P12, mTLS |

## Prerequisites

- **F5 Distributed Cloud Account** - Sign up at <https://www.f5.com/cloud/products/distributed-cloud-console>
- **Terraform >= 1.8** - Download from <https://www.terraform.io/downloads>
- **Console Access** - Ability to create credentials in the F5XC Console

## Creating Credentials in F5 Distributed Cloud

### Personal Credentials

Personal credentials are tied to your user account and ideal for development.

#### Creating an API Token

1. Open the F5 Distributed Cloud Console
2. Navigate to **Administration** → **Personal Management** → **Credentials**
3. Click **+ Add Credentials**
4. Enter a name (e.g., "terraform-dev-token")
5. Select **API Token** from the dropdown
6. Choose an expiry date
7. Click **Generate**, then **Copy** the token value

!> **Warning:** Copy and save your token immediately. You cannot retrieve it after closing this dialog.

#### Creating a P12 Certificate

1. Navigate to **Administration** → **Personal Management** → **Credentials**
2. Click **+ Add Credentials**
3. Enter a name (e.g., "terraform-dev-cert")
4. Select **API Certificate** from the dropdown
5. Enter and confirm a password
6. Select an expiry date
7. Click **Download** to get the `.p12` file

-> **Tip:** Store the P12 file securely and never commit it to version control.

### Service Credentials (IAM)

Service credentials are managed through IAM and recommended for production. They can be scoped to specific roles and namespaces.

#### Creating a Service API Token

1. Navigate to **Administration** → **IAM** → **Service Credentials**
2. Click **+ Add Service Credentials**
3. Enter a name (e.g., "terraform-cicd-token")
4. Select **API Token** from the dropdown
5. Optionally assign roles and namespaces to limit scope
6. Choose an expiry date
7. Click **Generate**, then **Copy** the token value

#### Creating a Service P12 Certificate

1. Navigate to **Administration** → **IAM** → **Service Credentials**
2. Click **+ Add Service Credentials**
3. Enter a name
4. Select **API Certificate** from the dropdown
5. Optionally assign roles and namespaces
6. Enter and confirm a password
7. Select an expiry date
8. Click **Download** to get the `.p12` file

## Authentication Methods

### Method 1: API Token Authentication (Simplest)

API tokens provide bearer token authentication over TLS. This is the quickest way to get started.

**Using Environment Variables:**

```bash
export VES_API_URL="https://your-tenant.console.ves.volterra.io"
export VES_API_TOKEN="your-api-token"
```

**Using Provider Configuration:**

```hcl
provider "f5xc" {
  api_url   = var.f5xc_api_url
  api_token = var.f5xc_api_token
}
```

### Method 2: P12 Certificate Authentication (Recommended for Production)

P12 certificates provide mutual TLS (mTLS) authentication, where both client and server verify each other's identity.

**Using Environment Variables:**

```bash
export VES_API_URL="https://your-tenant.console.ves.volterra.io"
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

### Method 3: PEM Certificate Authentication (Derived from P12)

PEM authentication uses separate certificate and private key files. Since F5XC only provides P12 certificates, you must extract PEM files using OpenSSL.

**When to use this method:** Only when your tooling specifically requires PEM format instead of P12.

**Step 1: Extract PEM files from P12:**

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

**Step 2: Configure the provider:**

**Using Environment Variables:**

```bash
export VES_API_URL="https://your-tenant.console.ves.volterra.io"
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

~> **Note:** For server certificate verification, specify a CA certificate using `VES_CACERT` environment variable or `api_ca_cert` provider attribute.

## Environment Variable Reference

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

```bash
# Add to ~/.bashrc or ~/.zshrc
export VES_API_URL="https://your-tenant.console.ves.volterra.io"
export VES_API_TOKEN="your-api-token"
```

Then reload: `source ~/.zshrc` or `source ~/.bashrc`

### Authentication Priority

When multiple methods are configured, the provider uses this priority:

1. **P12 Certificate** - If `api_p12_file` is set (requires `p12_password`)
2. **PEM Certificate** - If both `api_cert` and `api_key` are set
3. **API Token** - If `api_token` is set
4. **Error** - If none are provided

## CI/CD Integration

### GitHub Actions with API Token

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
      - uses: actions/checkout@v4

      - uses: hashicorp/setup-terraform@v3
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

**GitHub Secrets to configure:**

| Secret Name | Value |
|-------------|-------|
| `VES_API_URL` | `https://your-tenant.console.ves.volterra.io` |
| `VES_API_TOKEN` | Your API token value |

### GitHub Actions with P12 Certificate

For production deployments requiring mTLS:

```yaml
name: Terraform Deploy with P12

on:
  push:
    branches: [main]

jobs:
  terraform:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - uses: hashicorp/setup-terraform@v3
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
        env:
          VES_API_URL: ${{ secrets.VES_API_URL }}
          VES_P12_FILE: /tmp/f5xc-credentials.p12
          VES_P12_PASSWORD: ${{ secrets.VES_P12_PASSWORD }}
        run: terraform apply -auto-approve tfplan

      - name: Cleanup
        if: always()
        run: rm -f /tmp/f5xc-credentials.p12
```

**Encoding P12 for GitHub Secrets:**

```bash
# On macOS
base64 -i your-credentials.p12 | pbcopy

# On Linux
base64 -w 0 your-credentials.p12
```

**GitHub Secrets to configure:**

| Secret Name | Value |
|-------------|-------|
| `VES_API_URL` | `https://your-tenant.console.ves.volterra.io` |
| `VES_P12_BASE64` | Base64-encoded P12 file contents |
| `VES_P12_PASSWORD` | Password for the P12 file |

## Security Best Practices

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

## Troubleshooting

### Authentication Failed (401 Unauthorized)

1. Verify API URL does **NOT** include `/api` suffix (e.g., `https://tenant.console.ves.volterra.io`)
2. Check token hasn't expired
3. Verify token copied correctly (no whitespace)
4. Ensure environment variables are exported:

```bash
echo $VES_API_URL
echo $VES_API_TOKEN
```

### Certificate Verification Failed

1. Verify P12 password is correct
2. Check file path is absolute or correctly relative
3. Test P12 file:

```bash
openssl pkcs12 -in your-credentials.p12 -nokeys -info
```

### Permission Denied (403 Forbidden)

1. Verify credential has required permissions
2. For Service Credentials, check role and namespace assignments
3. Some operations require specific system roles

### Environment Variables Not Working

1. Ensure variables are exported:

```bash
export VES_API_TOKEN="token"  # Correct
VES_API_TOKEN="token"         # Won't work
```

2. Verify spelling is exact (case-sensitive)
3. Check for hidden characters in values

## Revoking Credentials

1. Navigate to **Administration** → **Personal Management** → **Credentials** (or **IAM** → **Service Credentials**)
2. Find the credential
3. Click **Actions** (three dots) → **Force Expiry**

## Next Steps

- [HTTP Load Balancer Guide](http-loadbalancer) - Deploy your first load balancer
- [Blindfold Functions Guide](blindfold) - Secure secret management
- [Namespace Resource](../resources/namespace) - Organize your resources

## Support

- [Provider Documentation](../index)
- [F5 Distributed Cloud Docs](https://docs.cloud.f5.com/)
- [F5 Credentials Guide](https://docs.cloud.f5.com/docs-v2/administration/how-tos/user-mgmt/Credentials)
- [GitHub Issues](https://github.com/robinmordasiewicz/terraform-provider-f5xc/issues)
