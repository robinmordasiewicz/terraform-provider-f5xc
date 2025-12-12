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

variable "ves_api_url" {
  description = "F5 Distributed Cloud API URL. Format: https://your-tenant.console.ves.volterra.io/api"
  type        = string
  default     = ""

  validation {
    condition     = var.ves_api_url == "" || can(regex("^https://.*\\.console\\.ves\\.volterra\\.io/api$", var.ves_api_url))
    error_message = "API URL must be in format: https://your-tenant.console.ves.volterra.io/api"
  }
}

# -----------------------------------------------------------------------------
# API Token Authentication
# -----------------------------------------------------------------------------

variable "ves_api_token" {
  description = "F5 Distributed Cloud API token for bearer authentication. Create in Console under Administration > Personal Management > Credentials."
  type        = string
  default     = ""
  sensitive   = true
}

# -----------------------------------------------------------------------------
# P12 Certificate Authentication
# -----------------------------------------------------------------------------

variable "ves_p12_file" {
  description = "Path to the P12 certificate file. Download from Console under Administration > Personal Management > Credentials."
  type        = string
  default     = ""

  validation {
    condition     = var.ves_p12_file == "" || can(regex(".*\\.p12$", var.ves_p12_file))
    error_message = "P12 file path must end with .p12 extension."
  }
}

variable "ves_p12_password" {
  description = "Password for the P12 certificate file. Set when creating the certificate in the Console."
  type        = string
  default     = ""
  sensitive   = true
}

# -----------------------------------------------------------------------------
# PEM Certificate Authentication
# -----------------------------------------------------------------------------

variable "ves_cert" {
  description = "Path to the PEM certificate file. Extract from P12 using: openssl pkcs12 -in creds.p12 -nodes -nokeys -out cert.pem"
  type        = string
  default     = ""

  validation {
    condition     = var.ves_cert == "" || can(regex(".*\\.(pem|cert|crt)$", var.ves_cert))
    error_message = "Certificate file path must end with .pem, .cert, or .crt extension."
  }
}

variable "ves_key" {
  description = "Path to the PEM private key file. Extract from P12 using: openssl pkcs12 -in creds.p12 -nodes -nocerts -out key.pem"
  type        = string
  default     = ""
  sensitive   = true
}

variable "ves_cacert" {
  description = "Optional path to CA certificate for server verification."
  type        = string
  default     = ""
}
