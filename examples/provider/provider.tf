# Configure the F5XC Provider
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
