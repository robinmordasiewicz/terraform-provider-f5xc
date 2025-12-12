# HTTP Load Balancer Guide - Variables
# =====================================
# Authentication is handled via environment variables:
#   VES_API_URL   - Your tenant URL
#   VES_API_TOKEN - Your API token

# -----------------------------------------------------------------------------
# Namespace Configuration
# -----------------------------------------------------------------------------

variable "namespace_name" {
  description = "Name for the namespace. Used for all resources in this deployment."
  type        = string
  default     = "demo"
}

variable "create_namespace" {
  description = "Whether to create a new namespace or use an existing one."
  type        = bool
  default     = true
}

# -----------------------------------------------------------------------------
# Application Configuration
# -----------------------------------------------------------------------------

variable "app_name" {
  description = "Application name prefix used for naming all resources."
  type        = string
  default     = "demo-app"
}

variable "domain" {
  description = "Domain name for the HTTP Load Balancer. Users will access your application at this domain."
  type        = string

  validation {
    condition     = can(regex("^[a-z0-9][a-z0-9.-]*[a-z0-9]$", var.domain))
    error_message = "Domain must be a valid DNS name (lowercase letters, numbers, dots, and hyphens)."
  }
}

# -----------------------------------------------------------------------------
# Origin Server Configuration
# -----------------------------------------------------------------------------

variable "origin_server" {
  description = "DNS name or IP address of your backend origin server."
  type        = string
}

variable "origin_port" {
  description = "Port number on the origin server."
  type        = number
  default     = 443

  validation {
    condition     = var.origin_port > 0 && var.origin_port <= 65535
    error_message = "Port must be between 1 and 65535."
  }
}

variable "health_check_path" {
  description = "HTTP path for health checks on the origin server."
  type        = string
  default     = "/health"
}

# -----------------------------------------------------------------------------
# Security Features
# -----------------------------------------------------------------------------

variable "enable_waf" {
  description = "Enable Web Application Firewall (WAF) in blocking mode."
  type        = bool
  default     = true
}

variable "enable_bot_defense" {
  description = "Enable Bot Defense to protect against automated attacks."
  type        = bool
  default     = true
}

variable "enable_rate_limiting" {
  description = "Enable rate limiting to prevent abuse."
  type        = bool
  default     = true
}

variable "rate_limit_requests" {
  description = "Maximum number of requests per minute before rate limiting kicks in."
  type        = number
  default     = 100

  validation {
    condition     = var.rate_limit_requests > 0
    error_message = "Rate limit must be greater than 0."
  }
}

# -----------------------------------------------------------------------------
# Labels
# -----------------------------------------------------------------------------

variable "labels" {
  description = "Labels to apply to all resources for organization and filtering."
  type        = map(string)
  default = {
    managed_by = "terraform"
    guide      = "http-loadbalancer"
  }
}
