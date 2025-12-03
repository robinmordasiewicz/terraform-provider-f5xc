---
page_title: "Guide: Blindfold Secret Management Functions"
subcategory: "Guides"
description: |-
  Learn how to securely encrypt secrets using F5XC Blindfold functions.
  Covers TLS certificates, cloud credentials, and container registry authentication.
---

# Blindfold Secret Management Functions

This guide walks you through using F5XC Blindfold functions to securely encrypt sensitive data. By the end, you'll understand how to:

- **Encrypt TLS private keys** - Secure certificate key storage
- **Protect cloud credentials** - AWS, Azure, and GCP secrets
- **Secure container registries** - Private image pull authentication
- **Understand the security model** - Local encryption, never transmitted unencrypted

## Overview

F5 Distributed Cloud Secret Management ("blindfold") provides client-side encryption for sensitive data. The blindfold functions encrypt your secrets locally using RSA-OAEP with SHA-256, meaning **your plaintext secrets never leave your machine unencrypted**.

### How It Works

```
┌─────────────────────────────────────────────────────────────────────┐
│                     Your Local Machine                              │
│                                                                     │
│  ┌──────────────┐    ┌──────────────────┐    ┌─────────────────┐   │
│  │ Secret       │───►│ Blindfold        │───►│ Encrypted       │   │
│  │ (plaintext)  │    │ Function         │    │ Ciphertext      │   │
│  └──────────────┘    │ (RSA-OAEP)       │    └────────┬────────┘   │
│                      └────────┬─────────┘             │            │
│                               │                       │            │
│                      ┌────────▼─────────┐             │            │
│                      │ F5XC Public Key  │             │            │
│                      │ (fetched once)   │             │            │
│                      └──────────────────┘             │            │
└──────────────────────────────────────────────────────│────────────┘
                                                        │
                                                        ▼
                                          ┌─────────────────────────┐
                                          │   F5 Distributed Cloud  │
                                          │   (stores ciphertext    │
                                          │    only)                │
                                          └─────────────────────────┘
```

### Security Properties

- **Local encryption**: Secrets encrypted on your machine before transmission
- **RSA-OAEP with SHA-256**: Industry-standard encryption algorithm
- **Policy-controlled decryption**: Only authorized F5XC services can decrypt
- **No plaintext storage**: F5XC never receives or stores plaintext secrets

## Prerequisites

Before you begin, ensure you have:

- **Terraform >= 1.8** - Provider-defined functions require Terraform 1.8 or later
- **F5 Distributed Cloud Account** - Sign up at <https://www.f5.com/cloud/products/distributed-cloud-console>
- **API Credentials** - Token or P12 certificate authentication configured

### Authentication Setup

Configure one of these authentication methods via environment variables:

**Option 1: API Token (Recommended for development)**
```bash
export F5XC_API_URL="https://your-tenant.console.ves.volterra.io/api"
export F5XC_API_TOKEN="your-api-token"
```

**Option 2: P12 Certificate (Recommended for production)**
```bash
export F5XC_API_URL="https://your-tenant.console.ves.volterra.io/api"
export F5XC_API_P12_FILE="/path/to/your-credentials.p12"
export F5XC_P12_PASSWORD="your-p12-password"  # pragma: allowlist secret
```

-> **Tip:** Add these to your shell profile (`~/.bashrc` or `~/.zshrc`) for persistence across terminal sessions.

## Quick Start

### Step 1: Clone the Repository

```bash
git clone https://github.com/robinmordasiewicz/terraform-provider-f5xc.git
cd terraform-provider-f5xc/examples/guides/blindfold
```

### Step 2: Configure Your Deployment

```bash
cp terraform.tfvars.example terraform.tfvars
```

Edit `terraform.tfvars` to enable the examples you want to test:

```hcl
# Enable the TLS certificate example
enable_certificate_example = true

# Namespace configuration
namespace_name   = "my-blindfold-test"
create_namespace = true
```

### Step 3: Generate Test Certificates (for certificate example)

```bash
openssl req -x509 -newkey rsa:2048 \
  -keyout certs/server.key \
  -out certs/server.crt \
  -days 365 -nodes \
  -subj "/CN=example.com"
```

### Step 4: Deploy

```bash
terraform init
terraform plan
terraform apply
```

Review the plan output, then type `yes` to confirm deployment.

### Step 5: Verify

1. Check the Terraform outputs for created resource names
2. Navigate to the F5XC Console
3. Verify the certificate shows encrypted private key (not plaintext)

## Understanding Blindfold Functions

The provider includes two blindfold functions:

### blindfold()

Encrypts base64-encoded plaintext:

```hcl
provider::f5xc::blindfold(plaintext, policy_name, namespace)
```

| Parameter | Type | Description |
|-----------|------|-------------|
| `plaintext` | string | Base64-encoded secret to encrypt |
| `policy_name` | string | Name of the SecretPolicy |
| `namespace` | string | Namespace containing the policy |

**Example:**
```hcl
location = provider::f5xc::blindfold(
  base64encode(var.my_secret),
  "ves-io-allow-volterra",
  "shared"
)
```

### blindfold_file()

Reads a file and encrypts its contents:

```hcl
provider::f5xc::blindfold_file(path, policy_name, namespace)
```

| Parameter | Type | Description |
|-----------|------|-------------|
| `path` | string | Path to the file to encrypt |
| `policy_name` | string | Name of the SecretPolicy |
| `namespace` | string | Namespace containing the policy |

**Example:**
```hcl
location = provider::f5xc::blindfold_file(
  "${path.module}/certs/server.key",
  "ves-io-allow-volterra",
  "shared"
)
```

### Built-in SecretPolicy

Every F5XC tenant includes a default policy: `ves-io-allow-volterra` in the `shared` namespace. This policy allows Volterra (F5XC) services to decrypt secrets.

## Use Cases

### TLS Certificate Private Key

Store TLS certificates for load balancers. Note that TLS private keys are typically too large (>1000 bytes) for direct blindfold encryption (~190 byte limit). Use `clear_secret_info` for TLS keys:

```hcl
resource "f5xc_certificate" "example" {
  name      = "my-certificate"
  namespace = "shared"

  certificate_url = "string:///${base64encode(file("${path.module}/certs/server.crt"))}"

  # TLS private keys are too large for blindfold encryption
  # Use clear_secret_info - the key is transmitted securely over HTTPS
  private_key {
    clear_secret_info {
      url = "string:///${base64encode(file("${path.module}/certs/server.key"))}"
    }
  }

  use_system_defaults {}
}
```

~> **Size Limitation:** Blindfold encryption uses RSA-OAEP which limits plaintext to ~190 bytes. TLS private keys exceed this limit. Use `clear_secret_info` for certificates - the data is transmitted securely over HTTPS. For secrets under 190 bytes (API keys, passwords), use blindfold functions as shown below.

### AWS Cloud Credentials

Protect AWS secret access keys for VPC site deployments:

```hcl
resource "f5xc_cloud_credentials" "aws" {
  name      = "aws-credentials"
  namespace = "system"

  aws_secret_key {
    access_key = var.aws_access_key_id

    secret_key {
      blindfold_secret_info {
        location = provider::f5xc::blindfold(
          base64encode(var.aws_secret_access_key),
          "ves-io-allow-volterra",
          "shared"
        )
      }
    }
  }
}
```

### Azure Cloud Credentials

Secure Azure service principal client secrets:

```hcl
resource "f5xc_cloud_credentials" "azure" {
  name      = "azure-credentials"
  namespace = "system"

  azure_client_secret {
    subscription_id = var.azure_subscription_id
    tenant_id       = var.azure_tenant_id
    client_id       = var.azure_client_id

    client_secret {
      blindfold_secret_info {
        location = provider::f5xc::blindfold(
          base64encode(var.azure_client_secret),
          "ves-io-allow-volterra",
          "shared"
        )
      }
    }
  }
}
```

### GCP Cloud Credentials

Encrypt GCP service account JSON key files:

```hcl
resource "f5xc_cloud_credentials" "gcp" {
  name      = "gcp-credentials"
  namespace = "system"

  gcp_cred_file {
    credential_file {
      blindfold_secret_info {
        location = provider::f5xc::blindfold_file(
          var.gcp_credentials_file,
          "ves-io-allow-volterra",
          "shared"
        )
      }
    }
  }
}
```

### Container Registry Authentication

Protect container registry passwords for private image pulls:

```hcl
resource "f5xc_container_registry" "example" {
  name      = "docker-registry"
  namespace = "shared"

  registry  = "docker.io"
  user_name = var.registry_username

  password {
    blindfold_secret_info {
      location = provider::f5xc::blindfold(
        base64encode(var.registry_password),
        "ves-io-allow-volterra",
        "shared"
      )
    }
  }
}
```

## Configuration Options

### Using Custom SecretPolicies

While the built-in `ves-io-allow-volterra` policy works for most cases, you can create custom policies for fine-grained access control:

```hcl
locals {
  policy_name = "my-custom-policy"
  policy_ns   = "my-namespace"
}

# Reference your custom policy
location = provider::f5xc::blindfold(
  base64encode(var.secret),
  local.policy_name,
  local.policy_ns
)
```

### Encrypting Multiple Secrets with for_each

```hcl
variable "secrets" {
  type = map(string)
  default = {
    "api-key"    = "secret1"
    "auth-token" = "secret2"
  }
}

locals {
  encrypted_secrets = {
    for name, value in var.secrets :
    name => provider::f5xc::blindfold(
      base64encode(value),
      "ves-io-allow-volterra",
      "shared"
    )
  }
}
```

## Technical Details

### Size Limitations

RSA-OAEP encryption has a maximum plaintext size based on the key size:

| Key Size | Maximum Plaintext |
|----------|-------------------|
| 2048-bit | ~190 bytes |
| 4096-bit | ~446 bytes |

~> **Note:** If your secret exceeds the size limit, consider splitting it or using a different approach. The function will return a clear error message if the plaintext is too large.

### Output Format

The blindfold functions return a sealed secret string with the `string:///` prefix followed by a base64-encoded JSON structure:

```
string:///eyJrZXlfdmVyc2lvbiI6InYxLjIuMyIsInBvbGljeV9pZCI6InNoYXJlZC92ZXMtaW8tYWxsb3ctdm9sdGVycmEiLCJ0ZW5hbnQiOiJ5b3VyLXRlbmFudCIsImRhdGEiOiJBQkNERUYxMjM0NTY3ODkwLi4uIn0=
```

When base64-decoded, the sealed JSON contains these fields:

```json
{
  "key_version": "v1.2.3",
  "policy_id": "shared/ves-io-allow-volterra",
  "tenant": "your-tenant",
  "data": "ABCDEF1234567890..."
}
```

Field descriptions:
- `key_version`: Public key version used for encryption
- `policy_id`: Reference to the SecretPolicy (namespace/name format)
- `tenant`: Your F5XC tenant identifier
- `data`: Base64-encoded RSA-OAEP ciphertext

### Function Behavior

- **Idempotent**: Same input produces different output (due to random padding)
- **Network required**: Functions fetch the public key from F5XC API
- **Caching**: Public keys are cached for the Terraform run

## Troubleshooting

### Authentication Configuration Error

**Symptom:** Error message about missing authentication configuration.

**Solution:**
```bash
# Verify environment variables are set
echo $F5XC_API_URL
echo $F5XC_API_TOKEN  # or F5XC_API_P12_FILE

# Set them if missing
export F5XC_API_URL="https://your-tenant.console.ves.volterra.io/api"
export F5XC_API_TOKEN="your-api-token"
```

### Policy Not Found

**Symptom:** Error about secret policy not found.

**Solutions:**

1. Use the built-in policy:
   ```hcl
   policy_name = "ves-io-allow-volterra"
   namespace   = "shared"
   ```

2. Verify your custom policy exists in F5XC Console

### Plaintext Too Large

**Symptom:** Error indicating plaintext exceeds maximum size.

**Solutions:**

1. Verify your secret size:
   ```bash
   wc -c < your-secret-file
   ```

2. For large files (like some GCP credentials), consider:
   - Extracting only the private key portion
   - Using F5XC's external secret management integration

### File Not Found

**Symptom:** Error about file not found when using `blindfold_file()`.

**Solutions:**

1. Use `${path.module}` for relative paths:
   ```hcl
   location = provider::f5xc::blindfold_file(
     "${path.module}/certs/server.key",  # Correct
     ...
   )
   ```

2. Verify the file exists and is readable

### Invalid Base64

**Symptom:** Error about invalid base64 encoding.

**Solution:** Ensure you're base64-encoding your plaintext:
```hcl
# Correct
location = provider::f5xc::blindfold(
  base64encode(var.secret),  # base64encode() wraps the secret
  ...
)

# Incorrect
location = provider::f5xc::blindfold(
  var.secret,  # Raw secret will fail
  ...
)
```

## Clean Up

To remove all resources created by this guide:

```bash
terraform destroy
```

Type `yes` to confirm destruction.

!> **Warning:** This will immediately remove all created resources. Ensure you have backups of any certificates or credentials you need.

## Next Steps

Now that you understand blindfold encryption, explore related resources:

- [Certificate Resource](../resources/certificate) - Full certificate management
- [Cloud Credentials Resource](../resources/cloud_credentials) - Cloud provider authentication
- [HTTP Load Balancer Guide](./http-loadbalancer) - Use certificates in load balancers
- [blindfold Function Reference](../functions/blindfold) - Function API details
- [blindfold_file Function Reference](../functions/blindfold_file) - Function API details

## Support

- **Provider Documentation:** [F5XC Provider](../index)
- **F5 Documentation:** [F5 Distributed Cloud Docs](https://docs.cloud.f5.com/)
- **Secret Management:** [F5XC Secret Management](https://docs.cloud.f5.com/docs/how-to/secrets-management)
- **Issues:** [GitHub Issues](https://github.com/robinmordasiewicz/terraform-provider-f5xc/issues)
