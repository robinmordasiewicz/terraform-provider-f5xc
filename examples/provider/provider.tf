# Configure the F5XC Provider
provider "f5xc" {
  api_url = "https://your-tenant.console.ves.volterra.io/api"
  api_p12 = "/path/to/certificate.p12"
}

# Alternatively, use environment variables:
# export F5XC_API_URL="https://your-tenant.console.ves.volterra.io/api"
# export F5XC_API_P12="/path/to/certificate.p12"
