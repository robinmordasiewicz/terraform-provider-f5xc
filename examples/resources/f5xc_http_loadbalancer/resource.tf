# HTTP Loadbalancer Resource Example
# [Category: Load Balancing] [Namespace: required] [DependsOn: namespace, origin_pool] Manages a HTTP Load Balancer resource in F5 Distributed Cloud for load balancing HTTP/HTTPS traffic with advanced routing and security.

# Basic HTTP Loadbalancer configuration
resource "f5xc_http_loadbalancer" "example" {
  name      = "example-http-loadbalancer"
  namespace = "staging"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  // One of the arguments from this list "advertise_custom advertise_on_public advertise_on_public_default_vip do_not_advertise" must be set

  advertise_on_public_default_vip = true

  // One of the arguments from this list "api_specification disable_api_definition" must be set

  disable_api_definition = true

  // One of the arguments from this list "disable_api_discovery enable_api_discovery" must be set

  enable_api_discovery {
    // One of the arguments from this list "api_discovery_from_code_scan api_discovery_from_discovered_schema api_discovery_from_live_traffic" must be set

    api_discovery_from_live_traffic {}

    discovered_api_settings {
      purge_duration_for_inactive_discovered_apis = "30"
    }

    // One of the arguments from this list "disable_learn_from_redirect_traffic enable_learn_from_redirect_traffic" must be set

    disable_learn_from_redirect_traffic = true
  }

  // One of the arguments from this list "api_testing disable_api_testing" must be set

  disable_api_testing = true

  // One of the arguments from this list "captcha_challenge enable_challenge js_challenge no_challenge policy_based_challenge" must be set

  js_challenge {
    cookie_expiry   = 3600
    custom_page     = ""
    js_script_delay = 5000
  }

  domains = ["app.example.com", "www.example.com"]

  // One of the arguments from this list "cookie_stickiness least_active random ring_hash round_robin source_ip_stickiness" must be set

  round_robin = true

  // One of the arguments from this list "http https https_auto_cert" must be set

  https_auto_cert {
    http_redirect = true
    add_hsts      = true

    // One of the arguments from this list "default_header no_headers server_name" must be set

    default_header {}

    tls_config {
      // One of the arguments from this list "custom_security default_security low_security medium_security" must be set

      default_security {}
    }

    // One of the arguments from this list "no_mtls use_mtls" must be set

    no_mtls {}
  }

  // One of the arguments from this list "disable_malicious_user_detection enable_malicious_user_detection" must be set

  enable_malicious_user_detection = true

  // One of the arguments from this list "disable_malware_protection malware_protection_settings" must be set

  disable_malware_protection = true

  // One of the arguments from this list "api_rate_limit disable_rate_limit rate_limit" must be set

  rate_limit {
    rate_limiter {
      name      = "example-rate-limiter"
      namespace = "shared"
    }
    no_ip_allowed_list {}
  }

  // One of the arguments from this list "default_sensitive_data_policy sensitive_data_policy" must be set

  default_sensitive_data_policy = true

  // One of the arguments from this list "active_service_policies no_service_policies service_policies_from_namespace" must be set

  active_service_policies {
    policies {
      name      = "example-service-policy"
      namespace = "shared"
    }
  }

  // One of the arguments from this list "disable_threat_mesh enable_threat_mesh" must be set

  enable_threat_mesh = true

  // One of the arguments from this list "disable_trust_client_ip_headers enable_trust_client_ip_headers" must be set

  disable_trust_client_ip_headers = true

  // One of the arguments from this list "user_id_client_ip user_identification" must be set

  user_identification {
    name      = "example-user-identification"
    namespace = "shared"
  }

  // One of the arguments from this list "app_firewall disable_waf" must be set

  app_firewall {
    name      = "example-app-firewall"
    namespace = "shared"
  }

  // One of the arguments from this list "bot_defense bot_defense_advanced disable_bot_defense" must be set

  bot_defense {
    policy {
      // One of the arguments from this list "js_download_path js_insert_all_pages js_insert_all_pages_except" must be set

      js_insert_all_pages {
        javascript_location = "AFTER_HEAD"
      }

      // One of the arguments from this list "disable_mobile_sdk enable_mobile_sdk" must be set

      disable_mobile_sdk {}
    }
    regional_endpoint = "US"
    timeout           = 1000
  }

  // Default route pools configuration
  default_route_pools {
    pool {
      name      = "example-origin-pool"
      namespace = "staging"
    }
    weight   = 1
    priority = 1
  }
}
