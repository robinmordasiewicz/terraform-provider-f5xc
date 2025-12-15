---
page_title: "Guide: Advanced HTTP Load Balancer Security"
subcategory: "Guides"
description: |-
  Advanced guide to deploying a fully-secured HTTP Load Balancer with all security
  controls including WAF, Data Guard, IP Reputation, Malicious User Detection, and
  Threat Mesh using F5 Distributed Cloud and Terraform.
---

# Advanced HTTP Load Balancer Security

This guide extends the [basic HTTP Load Balancer guide](http-loadbalancer) with advanced security features for production deployments requiring comprehensive protection against sophisticated threats.

By following this guide, you'll deploy an HTTP Load Balancer with **11 security controls**:

| Security Layer | Feature | Protection |
|----------------|---------|------------|
| **Perimeter** | IP Reputation | Blocks known malicious IPs by threat category |
| **Perimeter** | Threat Mesh | Global threat intelligence sharing |
| **Bot Defense** | JavaScript Challenge | Client-side bot detection |
| **Bot Defense** | Malicious User Detection | Behavioral analysis and risk scoring |
| **Application** | Web Application Firewall | Blocks SQLi, XSS, and OWASP Top 10 |
| **Application** | Bot Protection Settings | Signature-based bot classification |
| **Rate Control** | Rate Limiting | Prevents abuse with configurable thresholds |
| **Data Protection** | Data Guard | Masks sensitive data (CC, SSN) in responses |

## Prerequisites

Before you begin, ensure you have:

- **F5 Distributed Cloud Account** - Sign up at <https://www.f5.com/cloud/products/distributed-cloud-console>
- **API Token** - Generate credentials from the F5XC Console ([documentation](https://docs.cloud.f5.com/docs/how-to/user-mgmt/credentials))
- **Terraform >= 1.8** - Download from <https://www.terraform.io/downloads>
- **Namespace** - An existing namespace or permissions to create one
- **Backend Origin** - Your application server accessible from the internet

-> **Tip:** Review the [Authentication Guide](authentication) for detailed credential setup instructions.

## Complete Configuration

The following configuration creates a production-ready HTTP Load Balancer with all security features enabled.

### Provider Configuration

```hcl
terraform {
  required_version = ">= 1.0"
  required_providers {
    f5xc = {
      source  = "robinmordasiewicz/f5xc"
      version = ">= 2.5"
    }
  }
}

provider "f5xc" {
  api_token = var.api_token
  api_url   = var.api_url
}
```

### Variables

```hcl
variable "api_token" {
  description = "F5 XC API token for authentication"
  type        = string
  sensitive   = true
}

variable "api_url" {
  description = "F5 XC API URL (e.g., https://your-tenant.console.ves.volterra.io/api)"
  type        = string
}

variable "namespace" {
  description = "F5 XC namespace for the load balancer"
  type        = string
  default     = "default"
}

variable "name_prefix" {
  description = "Prefix for resource names"
  type        = string
  default     = "secure-app"
}

variable "domain" {
  description = "Domain for the load balancer"
  type        = string
}

variable "origin_server" {
  description = "Backend origin server DNS name"
  type        = string
}
```

### Web Application Firewall

The WAF provides signature-based attack detection with configurable bot protection. For detailed WAF configuration options, see [Create Web Application Firewall](https://docs.cloud.f5.com/docs-v2/web-app-and-api-protection/how-to/app-security/application-firewall).

```hcl
resource "f5xc_app_firewall" "waf" {
  name      = "${var.name_prefix}-waf"
  namespace = var.namespace

  # Blocking mode actively mitigates threats
  # Use monitoring {} for detection-only mode
  blocking {}

  detection_settings {
    signature_selection_setting {
      default_attack_type_settings {}
      high_medium_accuracy_signatures {}
    }
    enable_suppression {}
    enable_threat_campaigns {}

    # Bot protection with graduated response
    bot_protection_setting {
      malicious_bot_action  = "BLOCK"
      suspicious_bot_action = "REPORT"
      good_bot_action       = "REPORT"
    }
  }
}
```

~> **Note:** The default enforcement mode is `monitoring`, meaning threats are logged but not blocked. Use `blocking {}` for production deployments. See [WAF Enforcement Modes](https://docs.cloud.f5.com/docs-v2/web-app-and-api-protection/how-to/app-security/application-firewall) for details.

### Health Check

Configure active health monitoring for your origin servers:

```hcl
resource "f5xc_healthcheck" "http" {
  name      = "${var.name_prefix}-healthcheck"
  namespace = var.namespace

  http_health_check {
    path                  = "/health"
    expected_status_codes = ["200"]
  }

  timeout             = 3
  interval            = 15
  unhealthy_threshold = 3
  healthy_threshold   = 2
}
```

### Origin Pool

The origin pool defines your backend servers. For additional origin pool options, see [Origin Pools](https://docs.cloud.f5.com/docs-v2/multi-cloud-app-connect/how-to/load-balance/create-http-load-balancer).

```hcl
resource "f5xc_origin_pool" "backend" {
  name      = "${var.name_prefix}-origin-pool"
  namespace = var.namespace

  origin_servers {
    public_name {
      dns_name = var.origin_server
    }
  }

  port = 443

  use_tls {
    skip_server_verification {}
    tls_config {
      default_security {}
    }
    sni = var.origin_server
  }

  endpoint_selection     = "LOCAL_PREFERRED"
  loadbalancer_algorithm = "ROUND_ROBIN"

  healthcheck {
    name      = f5xc_healthcheck.http.name
    namespace = var.namespace
  }
}
```

### HTTP Load Balancer with All Security Features

This is the main resource that brings together all security controls:

```hcl
resource "f5xc_http_loadbalancer" "app" {
  name      = "${var.name_prefix}-lb"
  namespace = var.namespace
  domains   = [var.domain]

  http {
    port = 80
  }

  advertise_on_public_default_vip {}

  default_route_pools {
    pool {
      name      = f5xc_origin_pool.backend.name
      namespace = var.namespace
    }
    weight = 1
  }

  round_robin {}

  # ─────────────────────────────────────────────────────────────────────────────
  # WAF Configuration
  # ─────────────────────────────────────────────────────────────────────────────
  app_firewall {
    name      = f5xc_app_firewall.waf.name
    namespace = var.namespace
  }

  # ─────────────────────────────────────────────────────────────────────────────
  # Rate Limiting
  # Prevents abuse by limiting requests per client IP
  # See: https://docs.cloud.f5.com/docs/how-to/advanced-security/user-rate-limit
  # ─────────────────────────────────────────────────────────────────────────────
  rate_limit {
    no_ip_allowed_list {}
    rate_limiter {
      total_number     = 100
      unit             = "MINUTE"
      burst_multiplier = 2
      leaky_bucket {}
    }
  }

  # ─────────────────────────────────────────────────────────────────────────────
  # IP Reputation Filtering
  # Blocks IPs based on threat intelligence categories
  # See: https://docs.cloud.f5.com/docs/how-to/advanced-security/configure-ip-reputation
  # ─────────────────────────────────────────────────────────────────────────────
  enable_ip_reputation {
    ip_threat_categories = [
      "SPAM_SOURCES",
      "WEB_ATTACKS",
      "BOTNETS",
      "SCANNERS",
      "PHISHING",
      "PROXY",
      "TOR_PROXY",
      "DENIAL_OF_SERVICE"
    ]
  }

  # ─────────────────────────────────────────────────────────────────────────────
  # JavaScript Challenge
  # Client-side bot detection using JS challenge
  # ─────────────────────────────────────────────────────────────────────────────
  js_challenge {
    js_script_delay = 1000
    cookie_expiry   = 3600
  }

  # ─────────────────────────────────────────────────────────────────────────────
  # Data Guard
  # Masks sensitive data (credit cards, SSN) in responses
  # Requires WAF to be enabled
  # ─────────────────────────────────────────────────────────────────────────────
  data_guard_rules {
    metadata {
      name             = "${var.name_prefix}-data-guard"
      description_spec = "Mask sensitive data in all responses"
    }
    any_domain {}
    path {
      prefix = "/"
    }
    apply_data_guard {}
  }

  # ─────────────────────────────────────────────────────────────────────────────
  # Malicious User Detection
  # Behavioral analysis with risk scoring
  # See: https://docs.cloud.f5.com/docs-v2/web-app-and-api-protection/how-to/adv-security/malicious-users
  # ─────────────────────────────────────────────────────────────────────────────
  enable_malicious_user_detection {}

  # ─────────────────────────────────────────────────────────────────────────────
  # Threat Mesh
  # Global threat intelligence sharing across F5XC network
  # ─────────────────────────────────────────────────────────────────────────────
  enable_threat_mesh {}

  labels = {
    environment = "production"
    managed_by  = "terraform"
    security    = "advanced"
  }
}
```

## Understanding Each Security Feature

### IP Reputation Service

The IP Reputation service maintains a continuously-updated database of known malicious IP addresses. When enabled, requests from IPs matching configured threat categories are automatically blocked.

| Threat Category | Description |
|-----------------|-------------|
| `SPAM_SOURCES` | Known spam-sending IP addresses |
| `WEB_ATTACKS` | IPs involved in web-based attacks |
| `BOTNETS` | Command & control and infected hosts |
| `SCANNERS` | Reconnaissance, probes, brute force |
| `PHISHING` | Phishing and fraud operations |
| `PROXY` | Anonymous proxy services |
| `TOR_PROXY` | Tor exit nodes |
| `DENIAL_OF_SERVICE` | DoS and DDoS sources |

-> **Tip:** Start with all categories enabled, then selectively disable based on your application requirements. For example, disable `TOR_PROXY` if you need to support privacy-focused users.

### Data Guard

Data Guard automatically detects and masks sensitive data in HTTP responses before they reach clients. This protects against accidental data exposure such as:

- Credit card numbers (PAN)
- Social Security Numbers (SSN)
- Custom patterns (configurable)

!> **Important:** Data Guard requires WAF to be enabled. If you disable WAF, Data Guard will not function.

### Malicious User Detection

This feature uses behavioral analysis to identify potentially malicious users based on:

- **Rate Limiting Violations** - Exceeding configured rate limits
- **WAF Violations** - Triggering WAF rules
- **Bot Detection Signals** - Failing JavaScript challenges
- **Threat Intelligence** - IP reputation matches

Users are assigned a risk score, and mitigation actions can be configured based on thresholds.

### Threat Mesh

Threat Mesh enables sharing of threat intelligence across the F5 Distributed Cloud network. When a threat is detected at one customer's load balancer, that intelligence can protect all participating customers.

## Configuration Variations

### Conditional Security Features

Use Terraform variables to make security features configurable:

```hcl
variable "enable_waf" {
  description = "Enable WAF protection"
  type        = bool
  default     = true
}

variable "enable_data_guard" {
  description = "Enable Data Guard (requires WAF)"
  type        = bool
  default     = true
}

variable "enable_ip_reputation" {
  description = "Enable IP Reputation filtering"
  type        = bool
  default     = true
}

variable "ip_threat_categories" {
  description = "IP threat categories to block"
  type        = list(string)
  default = [
    "SPAM_SOURCES",
    "WEB_ATTACKS",
    "BOTNETS",
    "SCANNERS"
  ]
}
```

Then use dynamic blocks in the load balancer:

```hcl
resource "f5xc_http_loadbalancer" "app" {
  # ... base configuration ...

  dynamic "app_firewall" {
    for_each = var.enable_waf ? [1] : []
    content {
      name      = f5xc_app_firewall.waf[0].name
      namespace = var.namespace
    }
  }

  dynamic "disable_waf" {
    for_each = var.enable_waf ? [] : [1]
    content {}
  }

  dynamic "enable_ip_reputation" {
    for_each = var.enable_ip_reputation ? [1] : []
    content {
      ip_threat_categories = var.ip_threat_categories
    }
  }

  dynamic "disable_ip_reputation" {
    for_each = var.enable_ip_reputation ? [] : [1]
    content {}
  }

  dynamic "data_guard_rules" {
    for_each = var.enable_data_guard && var.enable_waf ? [1] : []
    content {
      metadata {
        name             = "${var.name_prefix}-data-guard"
        description_spec = "Mask sensitive data"
      }
      any_domain {}
      path {
        prefix = "/"
      }
      apply_data_guard {}
    }
  }
}
```

### WAF Monitoring Mode

For initial deployment or debugging, use monitoring mode instead of blocking:

```hcl
resource "f5xc_app_firewall" "waf" {
  name      = "${var.name_prefix}-waf"
  namespace = var.namespace

  # Monitoring mode - detect but don't block
  monitoring {}

  detection_settings {
    # ... same detection settings ...
  }
}
```

### Custom Rate Limiting

Adjust rate limiting based on your application's traffic patterns:

```hcl
variable "rate_limit_requests" {
  description = "Number of requests allowed per rate limit period"
  type        = number
  default     = 100
}

variable "rate_limit_unit" {
  description = "Rate limit period: SECOND, MINUTE, or HOUR"
  type        = string
  default     = "MINUTE"

  validation {
    condition     = contains(["SECOND", "MINUTE", "HOUR"], var.rate_limit_unit)
    error_message = "Rate limit unit must be SECOND, MINUTE, or HOUR."
  }
}
```

## Outputs

Add outputs to retrieve deployment information:

```hcl
output "load_balancer_name" {
  description = "Name of the HTTP load balancer"
  value       = f5xc_http_loadbalancer.app.name
}

output "security_summary" {
  description = "Summary of enabled security controls"
  value = {
    waf_enabled              = var.enable_waf
    waf_mode                 = var.enable_waf ? "blocking" : "disabled"
    rate_limiting            = "${var.rate_limit_requests} per ${var.rate_limit_unit}"
    ip_reputation            = var.enable_ip_reputation
    data_guard               = var.enable_data_guard && var.enable_waf
    malicious_user_detection = true
    threat_mesh              = true
    js_challenge             = true
  }
}
```

## Troubleshooting

### Data Guard Not Masking Data

**Symptom:** Sensitive data appears in responses despite Data Guard being configured.

**Solutions:**

1. Verify WAF is enabled (Data Guard requires WAF)
2. Check the path configuration matches your application routes
3. Verify the response content type is text-based (HTML, JSON, XML)

### IP Reputation Blocking Legitimate Users

**Symptom:** Users from corporate networks or VPNs are being blocked.

**Solutions:**

1. Review blocked requests in Security Analytics
2. Consider removing `PROXY` category if your users use VPNs
3. Add IP allow lists for known-good networks:

```hcl
rate_limit {
  ip_allowed_list {
    prefixes = ["10.0.0.0/8", "192.168.0.0/16"]
  }
  rate_limiter {
    # ... configuration ...
  }
}
```

### JavaScript Challenge Breaking Application

**Symptom:** API calls or mobile apps fail with JavaScript challenge.

**Solutions:**

1. Use `no_challenge {}` instead of `js_challenge {}` for API-only endpoints
2. Configure trusted client rules to bypass JS challenge for specific clients
3. Consider using `captcha_challenge {}` for interactive applications

## Security Best Practices

1. **Start with monitoring mode** - Deploy WAF in monitoring mode first to understand your traffic patterns
2. **Review security analytics** - Regularly review blocked requests in the F5XC Console
3. **Tune gradually** - Enable features one at a time and monitor impact
4. **Use all layers** - Defense in depth requires multiple security controls
5. **Keep Terraform state secure** - Use remote state with encryption for production

## Related Documentation

### F5 Distributed Cloud Documentation

- [Create HTTP Load Balancer](https://docs.cloud.f5.com/docs-v2/multi-cloud-app-connect/how-to/load-balance/create-http-load-balancer)
- [Web Application Firewall](https://docs.cloud.f5.com/docs-v2/web-app-and-api-protection/how-to/app-security/application-firewall)
- [IP Reputation Service](https://docs.cloud.f5.com/docs/how-to/advanced-security/configure-ip-reputation)
- [Rate Limiting](https://docs.cloud.f5.com/docs/how-to/advanced-security/user-rate-limit)
- [Bot Defense](https://docs.cloud.f5.com/docs/how-to/advanced-security/bot-defense)
- [Malicious User Detection](https://docs.cloud.f5.com/docs-v2/web-app-and-api-protection/how-to/adv-security/malicious-users)

### Provider Resources

- [f5xc_http_loadbalancer](../resources/http_loadbalancer)
- [f5xc_app_firewall](../resources/app_firewall)
- [f5xc_origin_pool](../resources/origin_pool)
- [f5xc_healthcheck](../resources/healthcheck)

## Support

- **Provider Issues:** [GitHub Issues](https://github.com/robinmordasiewicz/terraform-provider-f5xc/issues)
- **F5 Support:** [F5 Distributed Cloud Support](https://docs.cloud.f5.com/docs/support)
