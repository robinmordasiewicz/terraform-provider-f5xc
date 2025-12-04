# F5 Distributed Cloud Provider - Authentication Example

This example demonstrates the three authentication methods supported by the F5XC Terraform provider.

## Quick Start

### Option 1: Environment Variables (Recommended)

```bash
# For API Token authentication
export F5XC_API_URL="https://your-tenant.console.ves.volterra.io/api"
export F5XC_API_TOKEN="your-api-token"

# Initialize and apply
terraform init
terraform plan
```

### Option 2: Using tfvars File

```bash
# Copy the example configuration
cp terraform.tfvars.example terraform.tfvars

# Edit terraform.tfvars with your values
# Then run:
terraform init
terraform plan
```

## Authentication Methods

| Method | Environment Variables | Security Level |
|--------|----------------------|----------------|
| API Token | `F5XC_API_TOKEN` | Standard (one-way TLS) |
| P12 Certificate | `F5XC_API_P12_FILE`, `F5XC_P12_PASSWORD` | High (mTLS) |
| PEM Certificate | `F5XC_API_CERT`, `F5XC_API_KEY` | High (mTLS) |

## Files

| File | Description |
|------|-------------|
| `main.tf` | Provider configuration with authentication options |
| `variables.tf` | Variable definitions for explicit configuration |
| `outputs.tf` | Outputs to verify authentication |
| `terraform.tfvars.example` | Example variable values (copy to terraform.tfvars) |

## Creating Credentials

### API Token

1. Log in to F5 Distributed Cloud Console
2. Go to **Administration** → **Personal Management** → **Credentials**
3. Click **+ Add Credentials**
4. Select **API Token**, set expiry, and click **Generate**
5. Copy the token immediately (it won't be shown again)

### P12 Certificate

1. Log in to F5 Distributed Cloud Console
2. Go to **Administration** → **Personal Management** → **Credentials**
3. Click **+ Add Credentials**
4. Select **API Certificate**, enter a password
5. Click **Download** to get the `.p12` file

### PEM Certificate (from P12)

```bash
# Extract certificate
openssl pkcs12 -in credentials.p12 -nodes -nokeys -out cert.pem

# Extract private key
openssl pkcs12 -in credentials.p12 -nodes -nocerts -out key.pem
```

## Verification

After running `terraform plan`, you should see:

```
data.f5xc_namespace.system: Reading...
data.f5xc_namespace.system: Read complete

Changes to Outputs:
  + authentication_test = {
      + description = "F5 system namespace"
      + namespace   = "system"
    }
```

This confirms authentication is working correctly.

## Troubleshooting

### 401 Unauthorized

- Verify your API URL ends with `/api`
- Check that your token hasn't expired
- Ensure environment variables are exported (not just set)

### Certificate Errors

- Verify the P12 password is correct
- Use absolute paths for certificate files
- Check file permissions (should be readable)

## Documentation

- [Full Authentication Guide](https://registry.terraform.io/providers/robinmordasiewicz/f5xc/latest/docs/guides/authentication)
- [F5 Credentials Documentation](https://docs.cloud.f5.com/docs-v2/administration/how-tos/user-mgmt/Credentials)
