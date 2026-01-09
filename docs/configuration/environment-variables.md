# Environment Variables

The F5XC provider supports the following environment variables:

## Authentication

### `F5XC_API_TOKEN`

**Type**: String  
**Required**: One of `F5XC_API_TOKEN` or `F5XC_P12_FILE` is required

API token for F5 Distributed Cloud authentication.

```bash
export F5XC_API_TOKEN="your-api-token"
```

### `F5XC_P12_FILE`

**Type**: String  
**Required**: Alternative to `F5XC_API_TOKEN`

Path to P12 certificate file for authentication.

```bash
export F5XC_P12_FILE="/path/to/certificate.p12"
```

### `F5XC_P12_PASSWORD`

**Type**: String  
**Required**: When using `F5XC_P12_FILE`

Password for the P12 certificate file.

```bash
export F5XC_P12_PASSWORD="certificate-password"  # pragma: allowlist secret
```

## Configuration

### `F5XC_API_URL`

**Type**: String  
**Default**: `https://console.ves.volterra.io`

API endpoint URL for F5 Distributed Cloud.

```bash
export F5XC_API_URL="https://console.ves.volterra.io"
```

## Testing

### `TF_ACC`

**Type**: Boolean  
**Default**: `false`

Enable acceptance tests. Set to `1` when running acceptance tests.

```bash
TF_ACC=1 go test ./... -v
```

## See Also

- [Authentication guide](../getting-started/authentication.md)
- [Credential files](credential-files.md)
