# Credential Files

Alternative methods for providing authentication credentials.

## P12 Certificate Files

P12 certificates can be generated from the F5 Distributed Cloud Console:

1. Navigate to **Account Settings** → **Credentials**
2. Click **Add Credentials**
3. Select **API Certificate**
4. Download the `.p12` file
5. Note the password provided

Store the certificate securely and reference it via environment variable:

```bash
export F5XC_P12_FILE="/secure/path/certificate.p12"
export F5XC_P12_PASSWORD="provided-password"  # pragma: allowlist secret
```

## Security Best Practices

!!! warning "Security"
    - Never commit credential files to version control
    - Use environment variables instead of hardcoding credentials
    - Store P12 files with restrictive permissions (`chmod 600`)
    - Rotate credentials regularly

## See Also

- [Environment variables](environment-variables.md)
- [Authentication guide](../getting-started/authentication.md)
