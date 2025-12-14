# F5 Distributed Cloud Provider - Authentication Variables
# =========================================================
#
# These variables are used when configuring the provider with explicit
# values instead of environment variables. All sensitive variables are
# marked as such to prevent accidental exposure in logs and output.
#
# IMPORTANT: Never commit real credential values to version control!

# -----------------------------------------------------------------------------
# API URL Configuration
# -----------------------------------------------------------------------------

variable "f5xc_api_url" {
  description = "F5 Distributed Cloud API URL. Format: https://your-tenant.console.ves.volterra.io"
  type        = string
  default     = ""

  validation {
    condition     = var.f5xc_api_url == "" || can(regex("^https://.*\\.console\\.ves\\.volterra\\.io(/api)?$", var.f5xc_api_url))
    error_message = "API URL must be in format: https://your-tenant.console.ves.volterra.io"
  }
}

# -----------------------------------------------------------------------------
# API Token Authentication
# -----------------------------------------------------------------------------

variable "f5xc_api_token" {
  description = "F5 Distributed Cloud API token for bearer authentication. Create in Console under Administration > Personal Management > Credentials."
  type        = string
  default     = ""
  sensitive   = true
}

# -----------------------------------------------------------------------------
# P12 Certificate Authentication
# -----------------------------------------------------------------------------

variable "f5xc_p12_file" {
  description = "Path to the P12 certificate file. Download from Console under Administration > Personal Management > Credentials."
  type        = string
  default     = ""

  validation {
    condition     = var.f5xc_p12_file == "" || can(regex(".*\\.p12$", var.f5xc_p12_file))
    error_message = "P12 file path must end with .p12 extension."
  }
}

variable "f5xc_p12_password" {
  description = "Password for the P12 certificate file. Set when creating the certificate in the Console."
  type        = string
  default     = ""
  sensitive   = true
}

# -----------------------------------------------------------------------------
# PEM Certificate Authentication
# -----------------------------------------------------------------------------

variable "f5xc_cert" {
  description = "Path to the PEM certificate file. Extract from P12 using: openssl pkcs12 -in creds.p12 -nodes -nokeys -out cert.pem"
  type        = string
  default     = ""

  validation {
    condition     = var.f5xc_cert == "" || can(regex(".*\\.(pem|cert|crt)$", var.f5xc_cert))
    error_message = "Certificate file path must end with .pem, .cert, or .crt extension."
  }
}

variable "f5xc_key" {
  description = "Path to the PEM private key file. Extract from P12 using: openssl pkcs12 -in creds.p12 -nodes -nocerts -out key.pem"
  type        = string
  default     = ""
  sensitive   = true
}

variable "f5xc_cacert" {
  description = "Optional path to CA certificate for server verification."
  type        = string
  default     = ""
}
