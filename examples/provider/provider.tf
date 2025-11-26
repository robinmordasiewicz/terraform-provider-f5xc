# Configure the F5XC Provider with API Token Authentication
provider "f5xc" {
  api_url   = "https://your-tenant.console.ves.volterra.io/api"
  api_token = var.f5xc_api_token
}

# Alternatively, use environment variables:
# export F5XC_API_URL="https://your-tenant.console.ves.volterra.io/api"
# export F5XC_API_TOKEN="your-api-token"

variable "f5xc_api_token" {
  description = "F5 Distributed Cloud API Token"
  type        = string
  sensitive   = true
}

# Or use P12 Certificate Authentication:
# provider "f5xc" {
#   api_url      = "https://your-tenant.console.ves.volterra.io/api"
#   api_p12_file = "/path/to/certificate.p12"
#   p12_password = var.f5xc_p12_password
# }
#
# Environment variables for P12 authentication:
# export F5XC_API_URL="https://your-tenant.console.ves.volterra.io/api"
# export F5XC_API_P12_FILE="/path/to/certificate.p12"
# export F5XC_P12_PASSWORD="your-p12-password"
