# Authentication

The F5XC provider supports two authentication methods:

## API Token (Recommended)

Set the `F5XC_API_TOKEN` environment variable:

```bash
export F5XC_API_TOKEN="your-api-token"
```

Or configure it in the provider block:

```hcl
provider "f5xc" {
  api_token = "your-api-token"  # Not recommended - use environment variable
  api_url   = "https://console.ves.volterra.io"
}
```

### Obtaining an API Token

1. Log in to F5 Distributed Cloud Console
2. Navigate to **Account Settings** → **Credentials** → **Add Credentials**
3. Select **API Token**
4. Save the generated token securely

## P12 Certificate

Alternatively, use P12 certificate authentication:

```bash
export F5XC_P12_FILE="/path/to/certificate.p12"
export F5XC_P12_PASSWORD="certificate-password"  # pragma: allowlist secret
```

Or in the provider configuration:

```hcl
provider "f5xc" {
  p12_file     = "/path/to/certificate.p12"
  p12_password = "certificate-password"  # pragma: allowlist secret # Not recommended - use environment variable
  api_url      = "https://console.ves.volterra.io"
}
```

## Environment Variables

See the full list of [environment variables](../configuration/environment-variables.md).

## Next Steps

- [Quick start guide](quick-start.md)
- [Provider configuration](../configuration/environment-variables.md)
