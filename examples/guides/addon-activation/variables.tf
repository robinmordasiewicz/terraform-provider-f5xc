# F5 Distributed Cloud Provider - Addon Activation Variables
# ===========================================================

# =============================================================================
# ADDON ENABLEMENT TOGGLES
# =============================================================================

variable "enable_bot_defense" {
  description = "Whether to check and potentially activate Bot Defense addon"
  type        = bool
  default     = false
}

variable "enable_client_side_defense" {
  description = "Whether to check and potentially activate Client Side Defense addon"
  type        = bool
  default     = false
}

variable "enable_waap" {
  description = "Whether to check and potentially activate WAAP (Web App and API Protection) addon"
  type        = bool
  default     = false
}

# =============================================================================
# ACTIVATION SETTINGS
# =============================================================================

variable "activation_wait_time" {
  description = "Time to wait after activation before proceeding (e.g., '30s', '1m')"
  type        = string
  default     = "30s"

  validation {
    condition     = can(regex("^[0-9]+[smh]$", var.activation_wait_time))
    error_message = "The activation_wait_time must be a duration string like '30s', '1m', or '1h'."
  }
}

# =============================================================================
# DEMO RESOURCES
# =============================================================================

variable "create_demo_namespace" {
  description = "Whether to create a demo namespace (useful for testing)"
  type        = bool
  default     = false
}

variable "demo_namespace_name" {
  description = "Name for the demo namespace"
  type        = string
  default     = "addon-demo"

  validation {
    condition     = can(regex("^[a-z][a-z0-9-]*[a-z0-9]$", var.demo_namespace_name))
    error_message = "The namespace name must start with a letter, end with alphanumeric, and contain only lowercase letters, numbers, and hyphens."
  }
}
