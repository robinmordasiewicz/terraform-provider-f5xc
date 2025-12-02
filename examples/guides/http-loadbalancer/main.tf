# HTTP Load Balancer with Full Security Features
# ===============================================
# This configuration deploys:
# - Namespace (optional)
# - Health Check
# - Origin Pool
# - Application Firewall (WAF)
# - Rate Limiter
# - HTTP Load Balancer with bot defense, JS challenge, and more

terraform {
  required_version = ">= 1.8"
  required_providers {
    f5xc = {
      source = "robinmordasiewicz/f5xc"
    }
  }
}

# Provider configuration uses environment variables:
# F5XC_API_URL   - Your tenant URL (e.g., https://tenant.console.ves.volterra.io/api)
# F5XC_API_TOKEN - Your API token
provider "f5xc" {}

# -----------------------------------------------------------------------------
# Local Values
# -----------------------------------------------------------------------------

locals {
  # Use created namespace or existing one based on variable
  namespace = var.create_namespace ? f5xc_namespace.this[0].name : var.namespace_name

  # Common labels for all resources
  common_labels = merge(var.labels, {
    app = var.app_name
  })
}

# -----------------------------------------------------------------------------
# Namespace (Optional)
# -----------------------------------------------------------------------------

resource "f5xc_namespace" "this" {
  count = var.create_namespace ? 1 : 0

  name        = var.namespace_name
  namespace   = "system"
  description = "Namespace for ${var.app_name} HTTP Load Balancer guide"

  labels = local.common_labels
}

# -----------------------------------------------------------------------------
# Health Check
# -----------------------------------------------------------------------------

resource "f5xc_healthcheck" "this" {
  name      = "${var.app_name}-healthcheck"
  namespace = local.namespace

  labels = local.common_labels

  http_health_check {
    use_origin_server_name {}
    path = var.health_check_path
  }

  healthy_threshold   = 3
  unhealthy_threshold = 3
  interval            = 15
  timeout             = 5

  depends_on = [f5xc_namespace.this]
}

# -----------------------------------------------------------------------------
# Origin Pool
# -----------------------------------------------------------------------------

resource "f5xc_origin_pool" "this" {
  name      = "${var.app_name}-origin-pool"
  namespace = local.namespace

  labels = local.common_labels

  origin_servers {
    public_name {
      dns_name         = var.origin_server
      refresh_interval = 60
    }
  }

  port = var.origin_port

  use_tls {
    use_host_header_as_sni {}

    tls_config {
      default_security {}
    }

    no_mtls {}

    volterra_trusted_ca {}
  }

  healthcheck {
    name      = f5xc_healthcheck.this.name
    namespace = local.namespace
  }

  endpoint_selection     = "LOCAL_PREFERRED"
  loadbalancer_algorithm = "ROUND_ROBIN"

  depends_on = [f5xc_namespace.this]
}

# -----------------------------------------------------------------------------
# Application Firewall (WAF)
# -----------------------------------------------------------------------------

resource "f5xc_app_firewall" "this" {
  count = var.enable_waf ? 1 : 0

  name      = "${var.app_name}-waf"
  namespace = local.namespace

  labels = local.common_labels

  # WAF in blocking mode - actively blocks attacks
  blocking {}

  use_default_blocking_page {}

  bot_protection_setting {
    malicious_bot_action  = "BLOCK"
    suspicious_bot_action = "REPORT"
    good_bot_action       = "REPORT"
  }

  default_detection_settings {}

  allow_all_response_codes {}

  depends_on = [f5xc_namespace.this]
}

# Note: Rate limiting is configured inline within the HTTP Load Balancer
# A standalone f5xc_rate_limiter resource can be used for reusable rate limiting policies

# -----------------------------------------------------------------------------
# HTTP Load Balancer
# -----------------------------------------------------------------------------

resource "f5xc_http_loadbalancer" "this" {
  name      = "${var.app_name}-lb"
  namespace = local.namespace

  labels = local.common_labels

  # Domain configuration
  domains = [var.domain]

  # Advertise on F5 Distributed Cloud's global network
  advertise_on_public_default_vip {}

  # HTTPS with automatic certificate management
  https_auto_cert {
    http_redirect = true
    add_hsts      = true

    default_header {}

    tls_config {
      default_security {}
    }

    no_mtls {}
  }

  # Default route to origin pool
  default_route_pools {
    pool {
      name      = f5xc_origin_pool.this.name
      namespace = local.namespace
    }
    weight   = 1
    priority = 1
  }

  # Load balancing algorithm
  round_robin {}

  # JavaScript challenge for bot protection
  js_challenge {
    cookie_expiry   = 3600
    js_script_delay = 5000
  }

  # WAF configuration (conditional)
  dynamic "app_firewall" {
    for_each = var.enable_waf ? [1] : []
    content {
      name      = f5xc_app_firewall.this[0].name
      namespace = local.namespace
    }
  }

  # If WAF is disabled, explicitly disable it
  dynamic "disable_waf" {
    for_each = var.enable_waf ? [] : [1]
    content {}
  }

  # Rate limiting (conditional)
  dynamic "rate_limit" {
    for_each = var.enable_rate_limiting ? [1] : []
    content {
      rate_limiter {
        total_number     = var.rate_limit_requests
        unit             = "MINUTE"
        burst_multiplier = 10
        token_bucket {}
      }
      no_ip_allowed_list {}
    }
  }

  # If rate limiting is disabled
  dynamic "disable_rate_limit" {
    for_each = var.enable_rate_limiting ? [] : [1]
    content {}
  }

  # Bot defense (conditional)
  dynamic "bot_defense" {
    for_each = var.enable_bot_defense ? [1] : []
    content {
      policy {
        js_insert_all_pages {
          javascript_location = "AFTER_HEAD"
        }
        disable_mobile_sdk {}
      }
      regional_endpoint = "US"
      timeout           = 1000
    }
  }

  # If bot defense is disabled
  dynamic "disable_bot_defense" {
    for_each = var.enable_bot_defense ? [] : [1]
    content {}
  }

  # Enable additional security features
  enable_malicious_user_detection {}
  enable_threat_mesh {}

  # API settings
  disable_api_definition {}
  disable_api_discovery {}
  disable_api_testing {}

  # Data protection
  default_sensitive_data_policy {}

  # Service policies
  no_service_policies {}

  # Client IP trust
  disable_trust_client_ip_headers {}

  # User identification (using client IP)
  user_id_client_ip {}

  # Malware protection (disabled for this guide)
  disable_malware_protection {}

  depends_on = [f5xc_namespace.this]
}
