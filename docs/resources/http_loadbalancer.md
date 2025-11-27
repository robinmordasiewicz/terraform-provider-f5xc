---
page_title: "f5xc_http_loadbalancer Resource - terraform-provider-f5xc"
subcategory: "Load Balancing"
description: |-
  Manages a HTTPLoadBalancer resource in F5 Distributed Cloud for load balancing HTTP/HTTPS traffic with advanced routing and security.
---

# f5xc_http_loadbalancer (Resource)

Manages a HTTPLoadBalancer resource in F5 Distributed Cloud for load balancing HTTP/HTTPS traffic with advanced routing and security.

~> **Note** Please refer to [HTTP Loadbalancer API docs](https://docs.cloud.f5.com/docs-v2/api/views-http-loadbalancer) to learn more.

## Example Usage

```terraform
# HTTP Loadbalancer Resource Example
# Manages a HTTPLoadBalancer resource in F5 Distributed Cloud for load balancing HTTP/HTTPS traffic with advanced routing and security.

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

  domains = ["app.example.com", "`www.example.com"`]

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
```

<!-- schema generated by tfplugindocs -->
## Argument Reference

### Metadata Argument Reference

&#x2022; [`name`](#name) - Required String<br>Name of the HTTPLoadBalancer. Must be unique within the namespace

&#x2022; [`namespace`](#namespace) - Required String<br>Namespace where the HTTPLoadBalancer will be created

&#x2022; [`annotations`](#annotations) - Optional Map<br>Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata

&#x2022; [`description`](#description) - Optional String<br>Human readable description for the object

&#x2022; [`disable`](#disable) - Optional Bool<br>A value of true will administratively disable the object

&#x2022; [`labels`](#labels) - Optional Map<br>Labels is a user defined key value map that can be attached to resources for organization and filtering

### Spec Argument Reference

-> **One of the following:**
&#x2022; [`active_service_policies`](#active-service-policies) - Optional Block<br>Service Policy List. List of service policies<br>See [Active Service Policies](#active-service-policies) below for details.
<br><br>&#x2022; [`no_service_policies`](#no-service-policies) - Optional Block<br>Empty. This can be used for messages where no values are needed
<br><br>&#x2022; [`service_policies_from_namespace`](#service-policies-from-namespace) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`add_location`](#add-location) - Optional Bool<br>Add Location. x-example: true Appends header x-volterra-location = <RE-site-name> in responses. This configuration is ignored on CE sites

-> **One of the following:**
&#x2022; [`advertise_custom`](#advertise-custom) - Optional Block<br>Advertise Custom. This defines a way to advertise a VIP on specific sites<br>See [Advertise Custom](#advertise-custom) below for details.
<br><br>&#x2022; [`advertise_on_public`](#advertise-on-public) - Optional Block<br>Advertise Public. This defines a way to advertise a load balancer on public. If optional public_ip is provided, it will only be advertised on RE sites where that public_ip is available<br>See [Advertise On Public](#advertise-on-public) below for details.
<br><br>&#x2022; [`advertise_on_public_default_vip`](#advertise-on-public-default-vip) - Optional Block<br>Empty. This can be used for messages where no values are needed
<br><br>&#x2022; [`do_not_advertise`](#do-not-advertise) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`api_protection_rules`](#api-protection-rules) - Optional Block<br>API Protection Rules. API Protection Rules<br>See [API Protection Rules](#api-protection-rules) below for details.

-> **One of the following:**
&#x2022; [`api_rate_limit`](#api-rate-limit) - Optional Block<br>APIRateLimit
<br><br>&#x2022; [`disable_rate_limit`](#disable-rate-limit) - Optional Block<br>Empty. This can be used for messages where no values are needed
<br><br>&#x2022; [`rate_limit`](#rate-limit) - Optional Block<br>RateLimitConfigType

-> **One of the following:**
&#x2022; [`api_specification`](#api-specification) - Optional Block<br>API Specification and Validation. Settings for API specification (API definition, OpenAPI validation, etc.)
<br><br>&#x2022; [`disable_api_definition`](#disable-api-definition) - Optional Block<br>Empty. This can be used for messages where no values are needed

-> **One of the following:**
&#x2022; [`api_testing`](#api-testing) - Optional Block<br>API Testing
<br><br>&#x2022; [`disable_api_testing`](#disable-api-testing) - Optional Block<br>Empty. This can be used for messages where no values are needed

-> **One of the following:**
&#x2022; [`app_firewall`](#app-firewall) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name
<br><br>&#x2022; [`disable_waf`](#disable-waf) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`blocked_clients`](#blocked-clients) - Optional Block<br>Client Blocking Rules. Define rules to block IP Prefixes or AS numbers

-> **One of the following:**
&#x2022; [`bot_defense`](#bot-defense) - Optional Block<br>Bot Defense. This defines various configuration options for Bot Defense Policy
<br><br>&#x2022; [`bot_defense_advanced`](#bot-defense-advanced) - Optional Block<br>Bot Defense Advanced. Bot Defense Advanced
<br><br>&#x2022; [`disable_bot_defense`](#disable-bot-defense) - Optional Block<br>Empty. This can be used for messages where no values are needed

-> **One of the following:**
&#x2022; [`caching_policy`](#caching-policy) - Optional Block<br>Caching Policies. x-required Caching Policies for the CDN
<br><br>&#x2022; [`disable_caching`](#disable-caching) - Optional Block<br>Empty. This can be used for messages where no values are needed

-> **One of the following:**
&#x2022; [`captcha_challenge`](#captcha-challenge) - Optional Block<br>Captcha Challenge Parameters. Enables loadbalancer to perform captcha challenge Captcha challenge will be based on Google Recaptcha. With this feature enabled, only clients that pass the captcha challenge will be allowed to complete the HTTP request. When loadbalancer is configured to do Captcha Challenge, it will redirect the browser to an HTML page on every new HTTP request. This HTML page will have captcha challenge embedded in it. Client will be allowed to make the request only if the captcha challenge is successful. Loadbalancer will tag response header with a cookie to avoid Captcha challenge for subsequent requests. CAPTCHA is mainly used as a security check to ensure only human users can pass through. Generally, computers or bots are not capable of solving a captcha. You can enable either Javascript challenge or Captcha challenge on a virtual host
<br><br>&#x2022; [`enable_challenge`](#enable-challenge) - Optional Block<br>Enable Malicious User Challenge. Configure auto mitigation i.e risk based challenges for malicious users
<br><br>&#x2022; [`js_challenge`](#js-challenge) - Optional Block<br>Javascript Challenge Parameters. Enables loadbalancer to perform client browser compatibility test by redirecting to a page with Javascript. With this feature enabled, only clients that are capable of executing Javascript(mostly browsers) will be allowed to complete the HTTP request. When loadbalancer is configured to do Javascript Challenge, it will redirect the browser to an HTML page on every new HTTP request. This HTML page will have Javascript embedded in it. Loadbalancer chooses a set of random numbers for every new client and sends these numbers along with an encrypted answer with the request such that it embed these numbers as input in the Javascript. Javascript will run on the requestor browser and perform a complex Math operation. Script will submit the answer to loadbalancer. Loadbalancer will validate the answer by comparing the calculated answer with the decrypted answer (which was encrypted when it was sent back as reply) and allow the request to the upstream server only if the answer is correct. Loadbalancer will tag response header with a cookie to avoid Javascript challenge for subsequent requests. Javascript challenge serves following purposes * Validate that the request is coming via a browser that is capable for running Javascript * Force the browser to run a complex operation, f(X), that requires it to spend a large number of CPU cycles. This is to slow down a potential DOS attacker by making it difficult to launch a large request flood without having to spend even larger CPU cost at their end. You can enable either Javascript challenge or Captcha challenge on a virtual host
<br><br>&#x2022; [`no_challenge`](#no-challenge) - Optional Block<br>Empty. This can be used for messages where no values are needed
<br><br>&#x2022; [`policy_based_challenge`](#policy-based-challenge) - Optional Block<br>Policy Based Challenge. Specifies the settings for policy rule based challenge

-> **One of the following:**
&#x2022; [`client_side_defense`](#client-side-defense) - Optional Block<br>Client-Side Defense. This defines various configuration options for Client-Side Defense Policy
<br><br>&#x2022; [`disable_client_side_defense`](#disable-client-side-defense) - Optional Block<br>Empty. This can be used for messages where no values are needed

-> **One of the following:**
&#x2022; [`cookie_stickiness`](#cookie-stickiness) - Optional Block<br>Hashing using Cookie. Two types of cookie affinity: 1. Passive. Takes a cookie that's present in the cookies header and hashes on its value. 2. Generated. Generates and sets a cookie with an expiration (TTL) on the first request from the client in its response to the client, based on the endpoint the request gets sent to. The client then presents this on the next and all subsequent requests. The hash of this is sufficient to ensure these requests get sent to the same endpoint. The cookie is generated by hashing the source and destination ports and addresses so that multiple independent HTTP2 streams on the same connection will independently receive the same cookie, even if they arrive simultaneously
<br><br>&#x2022; [`least_active`](#least-active) - Optional Block<br>Empty. This can be used for messages where no values are needed
<br><br>&#x2022; [`random`](#random) - Optional Block<br>Empty. This can be used for messages where no values are needed
<br><br>&#x2022; [`ring_hash`](#ring-hash) - Optional Block<br>Hash Policy List. List of hash policy rules
<br><br>&#x2022; [`round_robin`](#round-robin) - Optional Block<br>Empty. This can be used for messages where no values are needed
<br><br>&#x2022; [`source_ip_stickiness`](#source-ip-stickiness) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`cors_policy`](#cors-policy) - Optional Block<br>CORS Policy. Cross-Origin Resource Sharing requests configuration specified at Virtual-host or Route level. Route level configuration takes precedence. An example of an Cross origin HTTP request GET /resources/public-data/ HTTP/1.1 Host: bar.other User-Agent: Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.5; en-US; rv:1.9.1b3pre) Gecko/20081130 Minefield/3.1b3pre Accept: text/HTML,application/xhtml+XML,application/XML;q=0.9,*/*;q=0.8 Accept-Language: en-us,en;q=0.5 Accept-Encoding: gzip,deflate Accept-Charset: ISO-8859-1,utf-8;q=0.7,*;q=0.7 Connection: keep-alive Referrer: `HTTP://foo.example/examples/access-control/simpleXSInvocation.HTML` Origin: `HTTP://foo.example` HTTP/1.1 200 OK Date: Mon, 01 Dec 2008 00:23:53 GMT Server: Apache/2.0.61 Access-Control-Allow-Origin: * Keep-Alive: timeout=2, max=100 Connection: Keep-Alive Transfer-Encoding: chunked Content-Type: application/XML An example for cross origin HTTP OPTIONS request with Access-Control-Request-* header OPTIONS /resources/post-here/ HTTP/1.1 Host: bar.other User-Agent: Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.5; en-US; rv:1.9.1b3pre) Gecko/20081130 Minefield/3.1b3pre Accept: text/HTML,application/xhtml+XML,application/XML;q=0.9,*/*;q=0.8 Accept-Language: en-us,en;q=0.5 Accept-Encoding: gzip,deflate Accept-Charset: ISO-8859-1,utf-8;q=0.7,*;q=0.7 Connection: keep-alive Origin: `HTTP://foo.example` Access-Control-Request-Method: POST Access-Control-Request-Headers: X-PINGOTHER, Content-Type HTTP/1.1 204 No Content Date: Mon, 01 Dec 2008 01:15:39 GMT Server: Apache/2.0.61 (Unix) Access-Control-Allow-Origin: `HTTP://foo.example` Access-Control-Allow-Methods: POST, GET, OPTIONS Access-Control-Allow-Headers: X-PINGOTHER, Content-Type Access-Control-Max-Age: 86400 Vary: Accept-Encoding, Origin Keep-Alive: timeout=2, max=100 Connection: Keep-Alive

&#x2022; [`csrf_policy`](#csrf-policy) - Optional Block<br>CSRF Policy. To mitigate CSRF attack , the policy checks where a request is coming from to determine if the request's origin is the same as its detination.The policy relies on two pieces of information used in determining if a request originated from the same host. 1. The origin that caused the user agent to issue the request (source origin). 2. The origin that the request is going to (target origin). When the policy evaluating a request, it ensures both pieces of information are present and compare their values. If the source origin is missing or origins do not match the request is rejected. The exception to this being if the source-origin has been added to they policy as valid. Because CSRF attacks specifically target state-changing requests, the policy only acts on the HTTP requests that have state-changing method (PUT,POST, etc.)

&#x2022; [`data_guard_rules`](#data-guard-rules) - Optional Block<br>Data Guard Rules. Data Guard prevents responses from exposing sensitive information by masking the data. The system masks credit card numbers and social security numbers leaked from the application from within the HTTP response with a string of asterisks (*). Note: App Firewall should be enabled, to use Data Guard feature

&#x2022; [`ddos_mitigation_rules`](#ddos-mitigation-rules) - Optional Block<br>DDOS Mitigation Rules. Define manual mitigation rules to block L7 DDOS attacks

-> **One of the following:**
&#x2022; [`default_pool`](#default-pool) - Optional Block<br>Global Specification. Shape of the origin pool specification
<br><br>&#x2022; [`default_pool_list`](#default-pool-list) - Optional Block<br>Origin Pool List Type. List of Origin Pools

&#x2022; [`default_route_pools`](#default-route-pools) - Optional Block<br>Origin Pools. Origin Pools used when no route is specified (default route)

-> **One of the following:**
&#x2022; [`default_sensitive_data_policy`](#default-sensitive-data-policy) - Optional Block<br>Empty. This can be used for messages where no values are needed
<br><br>&#x2022; [`sensitive_data_policy`](#sensitive-data-policy) - Optional Block<br>Sensitive Data Discovery. Settings for data type policy

-> **One of the following:**
&#x2022; [`disable_api_discovery`](#disable-api-discovery) - Optional Block<br>Empty. This can be used for messages where no values are needed
<br><br>&#x2022; [`enable_api_discovery`](#enable-api-discovery) - Optional Block<br>API Discovery Setting. Specifies the settings used for API discovery

-> **One of the following:**
&#x2022; [`disable_ip_reputation`](#disable-ip-reputation) - Optional Block<br>Empty. This can be used for messages where no values are needed
<br><br>&#x2022; [`enable_ip_reputation`](#enable-ip-reputation) - Optional Block<br>IP Threat Category List. List of IP threat categories

-> **One of the following:**
&#x2022; [`disable_malicious_user_detection`](#disable-malicious-user-detection) - Optional Block<br>Empty. This can be used for messages where no values are needed
<br><br>&#x2022; [`enable_malicious_user_detection`](#enable-malicious-user-detection) - Optional Block<br>Empty. This can be used for messages where no values are needed

-> **One of the following:**
&#x2022; [`disable_malware_protection`](#disable-malware-protection) - Optional Block<br>Empty. This can be used for messages where no values are needed
<br><br>&#x2022; [`malware_protection_settings`](#malware-protection-settings) - Optional Block<br>Malware Protection Policy. Malware Protection protects Web Apps and APIs, from malicious file uploads by scanning files in real-time

-> **One of the following:**
&#x2022; [`disable_threat_mesh`](#disable-threat-mesh) - Optional Block<br>Empty. This can be used for messages where no values are needed
<br><br>&#x2022; [`enable_threat_mesh`](#enable-threat-mesh) - Optional Block<br>Empty. This can be used for messages where no values are needed

-> **One of the following:**
&#x2022; [`disable_trust_client_ip_headers`](#disable-trust-client-ip-headers) - Optional Block<br>Empty. This can be used for messages where no values are needed
<br><br>&#x2022; [`enable_trust_client_ip_headers`](#enable-trust-client-ip-headers) - Optional Block<br>Trust Client IP Headers List. List of Client IP Headers

&#x2022; [`domains`](#domains) - Optional List<br>Domains. A list of Domains (host/authority header) that will be matched to load balancer. Supported Domains and search order: 1. Exact Domain names: `www.foo.com.` 2. Domains starting with a Wildcard: *.foo.com. Not supported Domains: - Just a Wildcard: * - A Wildcard and TLD with no root Domain: *.com. - A Wildcard not matching a whole DNS label. e.g. *.foo.com and *.bar.foo.com are valid Wildcards however *bar.foo.com, *-bar.foo.com, and bar*.foo.com are all invalid. Additional notes: A Wildcard will not match empty string. e.g. *.foo.com will match bar.foo.com and baz-bar.foo.com but not .foo.com. The longest Wildcards match first. Only a single virtual host in the entire route configuration can match on *. Also a Domain must be unique across all virtual hosts within an advertise policy. Domains are also used for SNI matching if the Loadbalancer type is HTTPS. Domains also indicate the list of names for which DNS resolution will be automatically resolved to IP addresses by the system

&#x2022; [`graphql_rules`](#graphql-rules) - Optional Block<br>GraphQL Inspection. GraphQL is a query language and server-side runtime for APIs which provides a complete and understandable description of the data in API. GraphQL gives clients the power to ask for exactly what they need, makes it easier to evolve APIs over time, and enables powerful developer tools. Policy configuration to analyze GraphQL queries and prevent GraphQL tailored attacks

-> **One of the following:**
&#x2022; [`http`](#http) - Optional Block<br>HTTP Choice. Choice for selecting HTTP proxy
<br><br>&#x2022; [`https`](#https) - Optional Block<br>BYOC HTTPS Choice. Choice for selecting HTTP proxy with bring your own certificates
<br><br>&#x2022; [`https_auto_cert`](#https-auto-cert) - Optional Block<br>HTTPS with Auto Certs Choice. Choice for selecting HTTP proxy with bring your own certificates

&#x2022; [`jwt_validation`](#jwt-validation) - Optional Block<br>JWT Validation. JWT Validation stops JWT replay attacks and JWT tampering by cryptographically verifying incoming JWTs before they are passed to your API origin. JWT Validation will also stop requests with expired tokens or tokens that are not yet valid

-> **One of the following:**
&#x2022; [`l7_ddos_action_block`](#l7-ddos-action-block) - Optional Block<br>Empty. This can be used for messages where no values are needed
<br><br>&#x2022; [`l7_ddos_action_default`](#l7-ddos-action-default) - Optional Block<br>Empty. This can be used for messages where no values are needed
<br><br>&#x2022; [`l7_ddos_action_js_challenge`](#l7-ddos-action-js-challenge) - Optional Block<br>Javascript Challenge Parameters. Enables loadbalancer to perform client browser compatibility test by redirecting to a page with Javascript. With this feature enabled, only clients that are capable of executing Javascript(mostly browsers) will be allowed to complete the HTTP request. When loadbalancer is configured to do Javascript Challenge, it will redirect the browser to an HTML page on every new HTTP request. This HTML page will have Javascript embedded in it. Loadbalancer chooses a set of random numbers for every new client and sends these numbers along with an encrypted answer with the request such that it embed these numbers as input in the Javascript. Javascript will run on the requestor browser and perform a complex Math operation. Script will submit the answer to loadbalancer. Loadbalancer will validate the answer by comparing the calculated answer with the decrypted answer (which was encrypted when it was sent back as reply) and allow the request to the upstream server only if the answer is correct. Loadbalancer will tag response header with a cookie to avoid Javascript challenge for subsequent requests. Javascript challenge serves following purposes * Validate that the request is coming via a browser that is capable for running Javascript * Force the browser to run a complex operation, f(X), that requires it to spend a large number of CPU cycles. This is to slow down a potential DOS attacker by making it difficult to launch a large request flood without having to spend even larger CPU cost at their end. You can enable either Javascript challenge or Captcha challenge on a virtual host

&#x2022; [`l7_ddos_protection`](#l7-ddos-protection) - Optional Block<br>L7 DDOS Protection Settings. L7 DDOS protection is critical for safeguarding web applications, APIs, and services that are exposed to the internet from sophisticated, volumetric, application-level threats. Configure actions, thresholds and policies to apply during L7 DDOS attack

&#x2022; [`more_option`](#more-option) - Optional Block<br>Advanced Options. This defines various options to define a route

-> **One of the following:**
&#x2022; [`multi_lb_app`](#multi-lb-app) - Optional Block<br>Empty. This can be used for messages where no values are needed
<br><br>&#x2022; [`single_lb_app`](#single-lb-app) - Optional Block<br>Single Load Balancer App Setting. Specific settings for Machine learning analysis on this HTTP LB, independently from other LBs

&#x2022; [`origin_server_subset_rule_list`](#origin-server-subset-rule-list) - Optional Block<br>Origin Server Subset Rule List Type. List of Origin Pools

&#x2022; [`protected_cookies`](#protected-cookies) - Optional Block<br>Cookie Protection. Allows setting attributes (SameSite, Secure, and HttpOnly) on cookies in responses. Cookie Tampering Protection prevents attackers from modifying the value of session cookies. For Cookie Tampering Protection, enabling a web app firewall (WAF) is a prerequisite. The configured mode of WAF (monitoring or blocking) will be enforced on the request when cookie tampering is identified. Note: We recommend enabling Secure and HttpOnly attributes along with cookie tampering protection

&#x2022; [`routes`](#routes) - Optional Block<br>Routes. Routes allow users to define match condition on a path and/or HTTP method to either forward matching traffic to origin pool or redirect matching traffic to a different URL or respond directly to matching traffic

&#x2022; [`sensitive_data_disclosure_rules`](#sensitive-data-disclosure-rules) - Optional Block<br>Sensitive Data Exposure Rules. Sensitive Data Exposure Rules allows specifying rules to mask sensitive data fields in API responses

-> **One of the following:**
&#x2022; [`slow_ddos_mitigation`](#slow-ddos-mitigation) - Optional Block<br>Slow DDOS Mitigation. 'Slow and low' attacks tie up server resources, leaving none available for servicing requests from actual users
<br><br>&#x2022; [`system_default_timeouts`](#system-default-timeouts) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`timeouts`](#timeouts) - Optional Block

&#x2022; [`trusted_clients`](#trusted-clients) - Optional Block<br>Trusted Client Rules. Define rules to skip processing of one or more features such as WAF, Bot Defense etc. for clients

-> **One of the following:**
&#x2022; [`user_id_client_ip`](#user-id-client-ip) - Optional Block<br>Empty. This can be used for messages where no values are needed
<br><br>&#x2022; [`user_identification`](#user-identification) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name

&#x2022; [`waf_exclusion`](#waf-exclusion) - Optional Block<br>WAF Exclusion

### Attributes Reference

In addition to all arguments above, the following attributes are exported:

&#x2022; [`id`](#id) - Optional String<br>Unique identifier for the resource

---

#### Active Service Policies

An [`active_service_policies`](#active-service-policies) block supports the following:

&#x2022; [`policies`](#policies) - Optional Block<br>Policies. Service Policies is a sequential engine where policies (and rules within the policy) are evaluated one after the other. It's important to define the correct order (policies evaluated from top to bottom in the list) for service policies, to get the intended result. For each request, its characteristics are evaluated based on the match criteria in each service policy starting at the top. If there is a match in the current policy, then the policy takes effect, and no more policies are evaluated. Otherwise, the next policy is evaluated. If all policies are evaluated and none match, then the request will be denied by default<br>See [Policies](#active-service-policies-policies) below.

#### Active Service Policies Policies

A [`policies`](#active-service-policies-policies) block (within [`active_service_policies`](#active-service-policies)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Advertise Custom

An [`advertise_custom`](#advertise-custom) block supports the following:

&#x2022; [`advertise_where`](#advertise-where) - Optional Block<br>List of Sites to Advertise. Where should this load balancer be available<br>See [Advertise Where](#advertise-custom-advertise-where) below.

#### Advertise Custom Advertise Where

An [`advertise_where`](#advertise-custom-advertise-where) block (within [`advertise_custom`](#advertise-custom)) supports the following:

&#x2022; [`advertise_on_public`](#advertise-on-public) - Optional Block<br>Advertise Public. This defines a way to advertise a load balancer on public. If optional public_ip is provided, it will only be advertised on RE sites where that public_ip is available<br>See [Advertise On Public](#advertise-custom-advertise-where-advertise-on-public) below.

&#x2022; [`port`](#port) - Optional Number<br>Listen Port. Port to Listen

&#x2022; [`port_ranges`](#port-ranges) - Optional String<br>Listen Port Ranges. A string containing a comma separated list of port ranges. Each port range consists of a single port or two ports separated by '-'

&#x2022; [`site`](#site) - Optional Block<br>Site. This defines a reference to a CE site along with network type and an optional IP address where a load balancer could be advertised<br>See [Site](#advertise-custom-advertise-where-site) below.

&#x2022; [`use_default_port`](#use-default-port) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`virtual_network`](#virtual-network) - Optional Block<br>Virtual Network. Parameters to advertise on a given virtual network<br>See [Virtual Network](#advertise-custom-advertise-where-virtual-network) below.

&#x2022; [`virtual_site`](#virtual-site) - Optional Block<br>Virtual Site. This defines a reference to a customer site virtual site along with network type where a load balancer could be advertised<br>See [Virtual Site](#advertise-custom-advertise-where-virtual-site) below.

&#x2022; [`virtual_site_with_vip`](#virtual-site-with-vip) - Optional Block<br>Virtual Site with Specified VIP. This defines a reference to a customer site virtual site along with network type and IP where a load balancer could be advertised<br>See [Virtual Site With VIP](#advertise-custom-advertise-where-virtual-site-with-vip) below.

&#x2022; [`vk8s_service`](#vk8s-service) - Optional Block<br>vK8s Services on RE. This defines a reference to a RE site or virtual site where a load balancer could be advertised in the vK8s service network<br>See [Vk8s Service](#advertise-custom-advertise-where-vk8s-service) below.

#### Advertise Custom Advertise Where Advertise On Public

An [`advertise_on_public`](#advertise-custom-advertise-where-advertise-on-public) block (within [`advertise_custom.advertise_where`](#advertise-custom-advertise-where)) supports the following:

&#x2022; [`public_ip`](#public-ip) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Public IP](#advertise-custom-advertise-where-advertise-on-public-public-ip) below.

#### Advertise Custom Advertise Where Advertise On Public Public IP

A [`public_ip`](#advertise-custom-advertise-where-advertise-on-public-public-ip) block (within [`advertise_custom.advertise_where.advertise_on_public`](#advertise-custom-advertise-where-advertise-on-public)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Advertise Custom Advertise Where Site

A [`site`](#advertise-custom-advertise-where-site) block (within [`advertise_custom.advertise_where`](#advertise-custom-advertise-where)) supports the following:

&#x2022; [`ip`](#ip) - Optional String<br>IP Address. Use given IP address as VIP on the site

&#x2022; [`network`](#network) - Optional String  Defaults to `SITE_NETWORK_INSIDE_AND_OUTSIDE`<br>Possible values are `SITE_NETWORK_INSIDE_AND_OUTSIDE`, `SITE_NETWORK_INSIDE`, `SITE_NETWORK_OUTSIDE`, `SITE_NETWORK_SERVICE`, `SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP`, `SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP`, `SITE_NETWORK_IP_FABRIC`<br>Site Network. This defines network types to be used on site All inside and outside networks. All inside and outside networks with internet VIP support. All inside networks. All outside networks. All outside networks with internet VIP support. vK8s service network. - SITE_NETWORK_IP_FABRIC: VER IP Fabric network for the site This Virtual network type is used for exposing virtual host on IP Fabric network on the VER site or for endpoint in IP Fabric network

&#x2022; [`site`](#site) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Site](#advertise-custom-advertise-where-site-site) below.

#### Advertise Custom Advertise Where Site Site

A [`site`](#advertise-custom-advertise-where-site-site) block (within [`advertise_custom.advertise_where.site`](#advertise-custom-advertise-where-site)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Advertise Custom Advertise Where Virtual Network

A [`virtual_network`](#advertise-custom-advertise-where-virtual-network) block (within [`advertise_custom.advertise_where`](#advertise-custom-advertise-where)) supports the following:

&#x2022; [`default_v6_vip`](#default-v6-vip) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`default_vip`](#default-vip) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`specific_v6_vip`](#specific-v6-vip) - Optional String<br>Specific V6 VIP. Use given IPv6 address as VIP on virtual Network

&#x2022; [`specific_vip`](#specific-vip) - Optional String<br>Specific V4 VIP. Use given IPv4 address as VIP on virtual Network

&#x2022; [`virtual_network`](#virtual-network) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Virtual Network](#advertise-custom-advertise-where-virtual-network-virtual-network) below.

#### Advertise Custom Advertise Where Virtual Network Virtual Network

A [`virtual_network`](#advertise-custom-advertise-where-virtual-network-virtual-network) block (within [`advertise_custom.advertise_where.virtual_network`](#advertise-custom-advertise-where-virtual-network)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Advertise Custom Advertise Where Virtual Site

A [`virtual_site`](#advertise-custom-advertise-where-virtual-site) block (within [`advertise_custom.advertise_where`](#advertise-custom-advertise-where)) supports the following:

&#x2022; [`network`](#network) - Optional String  Defaults to `SITE_NETWORK_INSIDE_AND_OUTSIDE`<br>Possible values are `SITE_NETWORK_INSIDE_AND_OUTSIDE`, `SITE_NETWORK_INSIDE`, `SITE_NETWORK_OUTSIDE`, `SITE_NETWORK_SERVICE`, `SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP`, `SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP`, `SITE_NETWORK_IP_FABRIC`<br>Site Network. This defines network types to be used on site All inside and outside networks. All inside and outside networks with internet VIP support. All inside networks. All outside networks. All outside networks with internet VIP support. vK8s service network. - SITE_NETWORK_IP_FABRIC: VER IP Fabric network for the site This Virtual network type is used for exposing virtual host on IP Fabric network on the VER site or for endpoint in IP Fabric network

&#x2022; [`virtual_site`](#virtual-site) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Virtual Site](#advertise-custom-advertise-where-virtual-site-virtual-site) below.

#### Advertise Custom Advertise Where Virtual Site Virtual Site

A [`virtual_site`](#advertise-custom-advertise-where-virtual-site-virtual-site) block (within [`advertise_custom.advertise_where.virtual_site`](#advertise-custom-advertise-where-virtual-site)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Advertise Custom Advertise Where Virtual Site With VIP

A [`virtual_site_with_vip`](#advertise-custom-advertise-where-virtual-site-with-vip) block (within [`advertise_custom.advertise_where`](#advertise-custom-advertise-where)) supports the following:

&#x2022; [`ip`](#ip) - Optional String<br>IP Address. Use given IP address as VIP on the site

&#x2022; [`network`](#network) - Optional String  Defaults to `SITE_NETWORK_SPECIFIED_VIP_OUTSIDE`<br>Possible values are `SITE_NETWORK_SPECIFIED_VIP_OUTSIDE`, `SITE_NETWORK_SPECIFIED_VIP_INSIDE`<br>Site Network. This defines network types to be used on virtual-site with specified VIP All outside networks. All inside networks

&#x2022; [`virtual_site`](#virtual-site) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Virtual Site](#advertise-custom-advertise-where-virtual-site-with-vip-virtual-site) below.

#### Advertise Custom Advertise Where Virtual Site With VIP Virtual Site

A [`virtual_site`](#advertise-custom-advertise-where-virtual-site-with-vip-virtual-site) block (within [`advertise_custom.advertise_where.virtual_site_with_vip`](#advertise-custom-advertise-where-virtual-site-with-vip)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Advertise Custom Advertise Where Vk8s Service

A [`vk8s_service`](#advertise-custom-advertise-where-vk8s-service) block (within [`advertise_custom.advertise_where`](#advertise-custom-advertise-where)) supports the following:

&#x2022; [`site`](#site) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Site](#advertise-custom-advertise-where-vk8s-service-site) below.

&#x2022; [`virtual_site`](#virtual-site) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Virtual Site](#advertise-custom-advertise-where-vk8s-service-virtual-site) below.

#### Advertise Custom Advertise Where Vk8s Service Site

A [`site`](#advertise-custom-advertise-where-vk8s-service-site) block (within [`advertise_custom.advertise_where.vk8s_service`](#advertise-custom-advertise-where-vk8s-service)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Advertise Custom Advertise Where Vk8s Service Virtual Site

A [`virtual_site`](#advertise-custom-advertise-where-vk8s-service-virtual-site) block (within [`advertise_custom.advertise_where.vk8s_service`](#advertise-custom-advertise-where-vk8s-service)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Advertise On Public

An [`advertise_on_public`](#advertise-on-public) block supports the following:

&#x2022; [`public_ip`](#public-ip) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Public IP](#advertise-on-public-public-ip) below.

#### Advertise On Public Public IP

A [`public_ip`](#advertise-on-public-public-ip) block (within [`advertise_on_public`](#advertise-on-public)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### API Protection Rules

An [`api_protection_rules`](#api-protection-rules) block supports the following:

&#x2022; [`api_endpoint_rules`](#api-endpoint-rules) - Optional Block<br>API Endpoints. This category defines specific rules per API endpoints. If request matches any of these rules, skipping second category rules<br>See [API Endpoint Rules](#api-protection-rules-api-endpoint-rules) below.

&#x2022; [`api_groups_rules`](#api-groups-rules) - Optional Block<br>Server URLs and API Groups. This category includes rules per API group or Server URL. For API groups, refer to API Definition which includes API groups derived from uploaded swaggers<br>See [API Groups Rules](#api-protection-rules-api-groups-rules) below.

#### API Protection Rules API Endpoint Rules

An [`api_endpoint_rules`](#api-protection-rules-api-endpoint-rules) block (within [`api_protection_rules`](#api-protection-rules)) supports the following:

&#x2022; [`action`](#action) - Optional Block<br>API Protection Rule Action. The action to take if the input request matches the rule<br>See [Action](#api-protection-rules-api-endpoint-rules-action) below.

&#x2022; [`any_domain`](#any-domain) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`api_endpoint_method`](#api-endpoint-method) - Optional Block<br>HTTP Method Matcher. A HTTP method matcher specifies a list of methods to match an input HTTP method. The match is considered successful if the input method is a member of the list. The result of the match based on the method list is inverted if invert_matcher is true<br>See [API Endpoint Method](#api-protection-rules-api-endpoint-rules-api-endpoint-method) below.

&#x2022; [`api_endpoint_path`](#api-endpoint-path) - Optional String<br>API Endpoint. The endpoint (path) of the request

&#x2022; [`client_matcher`](#client-matcher) - Optional Block<br>Client Matcher. Client conditions for matching a rule<br>See [Client Matcher](#api-protection-rules-api-endpoint-rules-client-matcher) below.

&#x2022; [`metadata`](#metadata) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#api-protection-rules-api-endpoint-rules-metadata) below.

&#x2022; [`request_matcher`](#request-matcher) - Optional Block<br>Request Matcher. Request conditions for matching a rule<br>See [Request Matcher](#api-protection-rules-api-endpoint-rules-request-matcher) below.

&#x2022; [`specific_domain`](#specific-domain) - Optional String<br>Specific Domain. The rule will apply for a specific domain. For example: API.example.com

#### API Protection Rules API Endpoint Rules Action

An [`action`](#api-protection-rules-api-endpoint-rules-action) block (within [`api_protection_rules.api_endpoint_rules`](#api-protection-rules-api-endpoint-rules)) supports the following:

&#x2022; [`allow`](#allow) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`deny`](#deny) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### API Protection Rules API Endpoint Rules API Endpoint Method

An [`api_endpoint_method`](#api-protection-rules-api-endpoint-rules-api-endpoint-method) block (within [`api_protection_rules.api_endpoint_rules`](#api-protection-rules-api-endpoint-rules)) supports the following:

&#x2022; [`invert_matcher`](#invert-matcher) - Optional Bool<br>Invert Method Matcher. Invert the match result

&#x2022; [`methods`](#methods) - Optional List  Defaults to `ANY`<br>Possible values are `ANY`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`, `COPY`<br>Method List. List of methods values to match against

#### API Protection Rules API Endpoint Rules Client Matcher

A [`client_matcher`](#api-protection-rules-api-endpoint-rules-client-matcher) block (within [`api_protection_rules.api_endpoint_rules`](#api-protection-rules-api-endpoint-rules)) supports the following:

&#x2022; [`any_client`](#any-client) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`any_ip`](#any-ip) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`asn_list`](#asn-list) - Optional Block<br>ASN Match List. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer<br>See [Asn List](#api-protection-rules-api-endpoint-rules-client-matcher-asn-list) below.

&#x2022; [`asn_matcher`](#asn-matcher) - Optional Block<br>ASN Matcher. Match any AS number contained in the list of bgp_asn_sets<br>See [Asn Matcher](#api-protection-rules-api-endpoint-rules-client-matcher-asn-matcher) below.

&#x2022; [`client_selector`](#client-selector) - Optional Block<br>Label Selector. This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expresssions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects. Label selector is immutable. expressions is a list of strings of label selection expression. Each string has ',' separated values which are 'AND' and all strings are logically 'OR'. BNF for expression string <selector-syntax> ::= <requirement> | <requirement> ',' <selector-syntax> <requirement> ::= [!] KEY [ <set-based-restriction> | <exact-match-restriction> ] <set-based-restriction> ::= '' | <inclusion-exclusion> <value-set> <inclusion-exclusion> ::= <inclusion> | <exclusion> <exclusion> ::= 'notin' <inclusion> ::= 'in' <value-set> ::= '(' <values> ')' <values> ::= VALUE | VALUE ',' <values> <exact-match-restriction> ::= ['='|'=='|'!='] VALUE<br>See [Client Selector](#api-protection-rules-api-endpoint-rules-client-matcher-client-selector) below.

&#x2022; [`ip_matcher`](#ip-matcher) - Optional Block<br>IP Prefix Matcher. Match any IP prefix contained in the list of ip_prefix_sets. The result of the match is inverted if invert_matcher is true<br>See [IP Matcher](#api-protection-rules-api-endpoint-rules-client-matcher-ip-matcher) below.

&#x2022; [`ip_prefix_list`](#ip-prefix-list) - Optional Block<br>IP Prefix Match List. List of IP Prefix strings to match against<br>See [IP Prefix List](#api-protection-rules-api-endpoint-rules-client-matcher-ip-prefix-list) below.

&#x2022; [`ip_threat_category_list`](#ip-threat-category-list) - Optional Block<br>IP Threat Category List Type. List of IP threat categories<br>See [IP Threat Category List](#api-protection-rules-api-endpoint-rules-client-matcher-ip-threat-category-list) below.

&#x2022; [`tls_fingerprint_matcher`](#tls-fingerprint-matcher) - Optional Block<br>TLS Fingerprint Matcher. A TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of supported positve match criteria includes a list of known classes of TLS fingerprints and a list of exact values. The match is considered successful if either of these positive criteria are satisfied and the input fingerprint is not one of the excluded values<br>See [TLS Fingerprint Matcher](#api-protection-rules-api-endpoint-rules-client-matcher-tls-fingerprint-matcher) below.

#### API Protection Rules API Endpoint Rules Client Matcher Asn List

An [`asn_list`](#api-protection-rules-api-endpoint-rules-client-matcher-asn-list) block (within [`api_protection_rules.api_endpoint_rules.client_matcher`](#api-protection-rules-api-endpoint-rules-client-matcher)) supports the following:

&#x2022; [`as_numbers`](#as-numbers) - Optional List<br>AS Numbers. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer

#### API Protection Rules API Endpoint Rules Client Matcher Asn Matcher

An [`asn_matcher`](#api-protection-rules-api-endpoint-rules-client-matcher-asn-matcher) block (within [`api_protection_rules.api_endpoint_rules.client_matcher`](#api-protection-rules-api-endpoint-rules-client-matcher)) supports the following:

&#x2022; [`asn_sets`](#asn-sets) - Optional Block<br>BGP ASN Sets. A list of references to bgp_asn_set objects<br>See [Asn Sets](#api-protection-rules-api-endpoint-rules-client-matcher-asn-matcher-asn-sets) below.

#### API Protection Rules API Endpoint Rules Client Matcher Asn Matcher Asn Sets

An [`asn_sets`](#api-protection-rules-api-endpoint-rules-client-matcher-asn-matcher-asn-sets) block (within [`api_protection_rules.api_endpoint_rules.client_matcher.asn_matcher`](#api-protection-rules-api-endpoint-rules-client-matcher-asn-matcher)) supports the following:

&#x2022; [`kind`](#kind) - Optional String<br>Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route')

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

&#x2022; [`uid`](#uid) - Optional String<br>UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid

#### API Protection Rules API Endpoint Rules Client Matcher Client Selector

A [`client_selector`](#api-protection-rules-api-endpoint-rules-client-matcher-client-selector) block (within [`api_protection_rules.api_endpoint_rules.client_matcher`](#api-protection-rules-api-endpoint-rules-client-matcher)) supports the following:

&#x2022; [`expressions`](#expressions) - Optional List<br>Selector Expression. expressions contains the kubernetes style label expression for selections

#### API Protection Rules API Endpoint Rules Client Matcher IP Matcher

An [`ip_matcher`](#api-protection-rules-api-endpoint-rules-client-matcher-ip-matcher) block (within [`api_protection_rules.api_endpoint_rules.client_matcher`](#api-protection-rules-api-endpoint-rules-client-matcher)) supports the following:

&#x2022; [`invert_matcher`](#invert-matcher) - Optional Bool<br>Invert IP Matcher. Invert the match result

&#x2022; [`prefix_sets`](#prefix-sets) - Optional Block<br>IP Prefix Sets. A list of references to ip_prefix_set objects<br>See [Prefix Sets](#api-protection-rules-api-endpoint-rules-client-matcher-ip-matcher-prefix-sets) below.

#### API Protection Rules API Endpoint Rules Client Matcher IP Matcher Prefix Sets

A [`prefix_sets`](#api-protection-rules-api-endpoint-rules-client-matcher-ip-matcher-prefix-sets) block (within [`api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher`](#api-protection-rules-api-endpoint-rules-client-matcher-ip-matcher)) supports the following:

&#x2022; [`kind`](#kind) - Optional String<br>Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route')

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

&#x2022; [`uid`](#uid) - Optional String<br>UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid

#### API Protection Rules API Endpoint Rules Client Matcher IP Prefix List

An [`ip_prefix_list`](#api-protection-rules-api-endpoint-rules-client-matcher-ip-prefix-list) block (within [`api_protection_rules.api_endpoint_rules.client_matcher`](#api-protection-rules-api-endpoint-rules-client-matcher)) supports the following:

&#x2022; [`invert_match`](#invert-match) - Optional Bool<br>Invert Match Result. Invert the match result

&#x2022; [`ip_prefixes`](#ip-prefixes) - Optional List<br>IPv4 Prefix List. List of IPv4 prefix strings

#### API Protection Rules API Endpoint Rules Client Matcher IP Threat Category List

An [`ip_threat_category_list`](#api-protection-rules-api-endpoint-rules-client-matcher-ip-threat-category-list) block (within [`api_protection_rules.api_endpoint_rules.client_matcher`](#api-protection-rules-api-endpoint-rules-client-matcher)) supports the following:

&#x2022; [`ip_threat_categories`](#ip-threat-categories) - Optional List  Defaults to `SPAM_SOURCES`<br>Possible values are `SPAM_SOURCES`, `WINDOWS_EXPLOITS`, `WEB_ATTACKS`, `BOTNETS`, `SCANNERS`, `REPUTATION`, `PHISHING`, `PROXY`, `MOBILE_THREATS`, `TOR_PROXY`, `DENIAL_OF_SERVICE`, `NETWORK`<br>List of IP Threat Categories to choose. The IP threat categories is obtained from the list and is used to auto-generate equivalent label selection expressions

#### API Protection Rules API Endpoint Rules Client Matcher TLS Fingerprint Matcher

A [`tls_fingerprint_matcher`](#api-protection-rules-api-endpoint-rules-client-matcher-tls-fingerprint-matcher) block (within [`api_protection_rules.api_endpoint_rules.client_matcher`](#api-protection-rules-api-endpoint-rules-client-matcher)) supports the following:

&#x2022; [`classes`](#classes) - Optional List  Defaults to `TLS_FINGERPRINT_NONE`<br>Possible values are `TLS_FINGERPRINT_NONE`, `ANY_MALICIOUS_FINGERPRINT`, `ADWARE`, `ADWIND`, `DRIDEX`, `GOOTKIT`, `GOZI`, `JBIFROST`, `QUAKBOT`, `RANSOMWARE`, `TROLDESH`, `TOFSEE`, `TORRENTLOCKER`, `TRICKBOT`<br>TLS fingerprint classes. A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against

&#x2022; [`exact_values`](#exact-values) - Optional List<br>Exact Values. A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against

&#x2022; [`excluded_values`](#excluded-values) - Optional List<br>Excluded Values. A list of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can be used to skip known false positives when using one or more known TLS fingerprint classes in the enclosing matcher

#### API Protection Rules API Endpoint Rules Metadata

A [`metadata`](#api-protection-rules-api-endpoint-rules-metadata) block (within [`api_protection_rules.api_endpoint_rules`](#api-protection-rules-api-endpoint-rules)) supports the following:

&#x2022; [`description`](#description) - Optional String<br>Description. Human readable description

&#x2022; [`name`](#name) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### API Protection Rules API Endpoint Rules Request Matcher

A [`request_matcher`](#api-protection-rules-api-endpoint-rules-request-matcher) block (within [`api_protection_rules.api_endpoint_rules`](#api-protection-rules-api-endpoint-rules)) supports the following:

&#x2022; [`cookie_matchers`](#cookie-matchers) - Optional Block<br>Cookie Matchers. A list of predicates for all cookies that need to be matched. The criteria for matching each cookie is described in individual instances of CookieMatcherType. The actual cookie values are extracted from the request API as a list of strings for each cookie name. Note that all specified cookie matcher predicates must evaluate to true<br>See [Cookie Matchers](#api-protection-rules-api-endpoint-rules-request-matcher-cookie-matchers) below.

&#x2022; [`headers`](#headers) - Optional Block<br>HTTP Headers. A list of predicates for various HTTP headers that need to match. The criteria for matching each HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values are extracted from the request API as a list of strings for each HTTP header type. Note that all specified header predicates must evaluate to true<br>See [Headers](#api-protection-rules-api-endpoint-rules-request-matcher-headers) below.

&#x2022; [`jwt_claims`](#jwt-claims) - Optional Block<br>JWT Claims. A list of predicates for various JWT claims that need to match. The criteria for matching each JWT claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are extracted from the JWT payload as a list of strings. Note that all specified JWT claim predicates must evaluate to true. Note that this feature only works on LBs with JWT Validation feature enabled<br>See [JWT Claims](#api-protection-rules-api-endpoint-rules-request-matcher-jwt-claims) below.

&#x2022; [`query_params`](#query-params) - Optional Block<br>HTTP Query Parameters. A list of predicates for all query parameters that need to be matched. The criteria for matching each query parameter are described in individual instances of QueryParameterMatcherType. The actual query parameter values are extracted from the request API as a list of strings for each query parameter name. Note that all specified query parameter predicates must evaluate to true<br>See [Query Params](#api-protection-rules-api-endpoint-rules-request-matcher-query-params) below.

#### API Protection Rules API Endpoint Rules Request Matcher Cookie Matchers

A [`cookie_matchers`](#api-protection-rules-api-endpoint-rules-request-matcher-cookie-matchers) block (within [`api_protection_rules.api_endpoint_rules.request_matcher`](#api-protection-rules-api-endpoint-rules-request-matcher)) supports the following:

&#x2022; [`check_not_present`](#check-not-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`check_present`](#check-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`invert_matcher`](#invert-matcher) - Optional Bool<br>Invert Matcher. Invert Match of the expression defined

&#x2022; [`item`](#item) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#api-protection-rules-api-endpoint-rules-request-matcher-cookie-matchers-item) below.

&#x2022; [`name`](#name) - Optional String<br>Cookie Name. A case-sensitive cookie name

#### API Protection Rules API Endpoint Rules Request Matcher Cookie Matchers Item

An [`item`](#api-protection-rules-api-endpoint-rules-request-matcher-cookie-matchers-item) block (within [`api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers`](#api-protection-rules-api-endpoint-rules-request-matcher-cookie-matchers)) supports the following:

&#x2022; [`exact_values`](#exact-values) - Optional List<br>Exact Values. A list of exact values to match the input against

&#x2022; [`regex_values`](#regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

&#x2022; [`transformers`](#transformers) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### API Protection Rules API Endpoint Rules Request Matcher Headers

A [`headers`](#api-protection-rules-api-endpoint-rules-request-matcher-headers) block (within [`api_protection_rules.api_endpoint_rules.request_matcher`](#api-protection-rules-api-endpoint-rules-request-matcher)) supports the following:

&#x2022; [`check_not_present`](#check-not-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`check_present`](#check-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`invert_matcher`](#invert-matcher) - Optional Bool<br>Invert Header Matcher. Invert the match result

&#x2022; [`item`](#item) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#api-protection-rules-api-endpoint-rules-request-matcher-headers-item) below.

&#x2022; [`name`](#name) - Optional String<br>Header Name. A case-insensitive HTTP header name

#### API Protection Rules API Endpoint Rules Request Matcher Headers Item

An [`item`](#api-protection-rules-api-endpoint-rules-request-matcher-headers-item) block (within [`api_protection_rules.api_endpoint_rules.request_matcher.headers`](#api-protection-rules-api-endpoint-rules-request-matcher-headers)) supports the following:

&#x2022; [`exact_values`](#exact-values) - Optional List<br>Exact Values. A list of exact values to match the input against

&#x2022; [`regex_values`](#regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

&#x2022; [`transformers`](#transformers) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### API Protection Rules API Endpoint Rules Request Matcher JWT Claims

A [`jwt_claims`](#api-protection-rules-api-endpoint-rules-request-matcher-jwt-claims) block (within [`api_protection_rules.api_endpoint_rules.request_matcher`](#api-protection-rules-api-endpoint-rules-request-matcher)) supports the following:

&#x2022; [`check_not_present`](#check-not-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`check_present`](#check-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`invert_matcher`](#invert-matcher) - Optional Bool<br>Invert Matcher. Invert the match result

&#x2022; [`item`](#item) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#api-protection-rules-api-endpoint-rules-request-matcher-jwt-claims-item) below.

&#x2022; [`name`](#name) - Optional String<br>JWT Claim Name. JWT claim name

#### API Protection Rules API Endpoint Rules Request Matcher JWT Claims Item

An [`item`](#api-protection-rules-api-endpoint-rules-request-matcher-jwt-claims-item) block (within [`api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims`](#api-protection-rules-api-endpoint-rules-request-matcher-jwt-claims)) supports the following:

&#x2022; [`exact_values`](#exact-values) - Optional List<br>Exact Values. A list of exact values to match the input against

&#x2022; [`regex_values`](#regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

&#x2022; [`transformers`](#transformers) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### API Protection Rules API Endpoint Rules Request Matcher Query Params

A [`query_params`](#api-protection-rules-api-endpoint-rules-request-matcher-query-params) block (within [`api_protection_rules.api_endpoint_rules.request_matcher`](#api-protection-rules-api-endpoint-rules-request-matcher)) supports the following:

&#x2022; [`check_not_present`](#check-not-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`check_present`](#check-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`invert_matcher`](#invert-matcher) - Optional Bool<br>Invert Query Parameter Matcher. Invert the match result

&#x2022; [`item`](#item) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#api-protection-rules-api-endpoint-rules-request-matcher-query-params-item) below.

&#x2022; [`key`](#key) - Optional String<br>Query Parameter Name. A case-sensitive HTTP query parameter name

#### API Protection Rules API Endpoint Rules Request Matcher Query Params Item

An [`item`](#api-protection-rules-api-endpoint-rules-request-matcher-query-params-item) block (within [`api_protection_rules.api_endpoint_rules.request_matcher.query_params`](#api-protection-rules-api-endpoint-rules-request-matcher-query-params)) supports the following:

&#x2022; [`exact_values`](#exact-values) - Optional List<br>Exact Values. A list of exact values to match the input against

&#x2022; [`regex_values`](#regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

&#x2022; [`transformers`](#transformers) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### API Protection Rules API Groups Rules

An [`api_groups_rules`](#api-protection-rules-api-groups-rules) block (within [`api_protection_rules`](#api-protection-rules)) supports the following:

&#x2022; [`action`](#action) - Optional Block<br>API Protection Rule Action. The action to take if the input request matches the rule<br>See [Action](#api-protection-rules-api-groups-rules-action) below.

&#x2022; [`any_domain`](#any-domain) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`api_group`](#api-group) - Optional String<br>API Group. API groups derived from API Definition swaggers. For example oas-all-operations including all paths and methods from the swaggers, oas-base-urls covering all requests under base-paths from the swaggers. Custom groups can be created if user tags paths or operations with 'x-volterra-API-group' extensions inside swaggers

&#x2022; [`base_path`](#base-path) - Optional String<br>Base Path. Prefix of the request path. For example: /v1

&#x2022; [`client_matcher`](#client-matcher) - Optional Block<br>Client Matcher. Client conditions for matching a rule<br>See [Client Matcher](#api-protection-rules-api-groups-rules-client-matcher) below.

&#x2022; [`metadata`](#metadata) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#api-protection-rules-api-groups-rules-metadata) below.

&#x2022; [`request_matcher`](#request-matcher) - Optional Block<br>Request Matcher. Request conditions for matching a rule<br>See [Request Matcher](#api-protection-rules-api-groups-rules-request-matcher) below.

&#x2022; [`specific_domain`](#specific-domain) - Optional String<br>Specific Domain. The rule will apply for a specific domain. For example: API.example.com

#### API Protection Rules API Groups Rules Action

An [`action`](#api-protection-rules-api-groups-rules-action) block (within [`api_protection_rules.api_groups_rules`](#api-protection-rules-api-groups-rules)) supports the following:

&#x2022; [`allow`](#allow) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`deny`](#deny) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### API Protection Rules API Groups Rules Client Matcher

A [`client_matcher`](#api-protection-rules-api-groups-rules-client-matcher) block (within [`api_protection_rules.api_groups_rules`](#api-protection-rules-api-groups-rules)) supports the following:

&#x2022; [`any_client`](#any-client) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`any_ip`](#any-ip) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`asn_list`](#asn-list) - Optional Block<br>ASN Match List. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer<br>See [Asn List](#api-protection-rules-api-groups-rules-client-matcher-asn-list) below.

&#x2022; [`asn_matcher`](#asn-matcher) - Optional Block<br>ASN Matcher. Match any AS number contained in the list of bgp_asn_sets<br>See [Asn Matcher](#api-protection-rules-api-groups-rules-client-matcher-asn-matcher) below.

&#x2022; [`client_selector`](#client-selector) - Optional Block<br>Label Selector. This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expresssions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects. Label selector is immutable. expressions is a list of strings of label selection expression. Each string has ',' separated values which are 'AND' and all strings are logically 'OR'. BNF for expression string <selector-syntax> ::= <requirement> | <requirement> ',' <selector-syntax> <requirement> ::= [!] KEY [ <set-based-restriction> | <exact-match-restriction> ] <set-based-restriction> ::= '' | <inclusion-exclusion> <value-set> <inclusion-exclusion> ::= <inclusion> | <exclusion> <exclusion> ::= 'notin' <inclusion> ::= 'in' <value-set> ::= '(' <values> ')' <values> ::= VALUE | VALUE ',' <values> <exact-match-restriction> ::= ['='|'=='|'!='] VALUE<br>See [Client Selector](#api-protection-rules-api-groups-rules-client-matcher-client-selector) below.

&#x2022; [`ip_matcher`](#ip-matcher) - Optional Block<br>IP Prefix Matcher. Match any IP prefix contained in the list of ip_prefix_sets. The result of the match is inverted if invert_matcher is true<br>See [IP Matcher](#api-protection-rules-api-groups-rules-client-matcher-ip-matcher) below.

&#x2022; [`ip_prefix_list`](#ip-prefix-list) - Optional Block<br>IP Prefix Match List. List of IP Prefix strings to match against<br>See [IP Prefix List](#api-protection-rules-api-groups-rules-client-matcher-ip-prefix-list) below.

&#x2022; [`ip_threat_category_list`](#ip-threat-category-list) - Optional Block<br>IP Threat Category List Type. List of IP threat categories<br>See [IP Threat Category List](#api-protection-rules-api-groups-rules-client-matcher-ip-threat-category-list) below.

&#x2022; [`tls_fingerprint_matcher`](#tls-fingerprint-matcher) - Optional Block<br>TLS Fingerprint Matcher. A TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of supported positve match criteria includes a list of known classes of TLS fingerprints and a list of exact values. The match is considered successful if either of these positive criteria are satisfied and the input fingerprint is not one of the excluded values<br>See [TLS Fingerprint Matcher](#api-protection-rules-api-groups-rules-client-matcher-tls-fingerprint-matcher) below.

#### API Protection Rules API Groups Rules Client Matcher Asn List

An [`asn_list`](#api-protection-rules-api-groups-rules-client-matcher-asn-list) block (within [`api_protection_rules.api_groups_rules.client_matcher`](#api-protection-rules-api-groups-rules-client-matcher)) supports the following:

&#x2022; [`as_numbers`](#as-numbers) - Optional List<br>AS Numbers. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer

#### API Protection Rules API Groups Rules Client Matcher Asn Matcher

An [`asn_matcher`](#api-protection-rules-api-groups-rules-client-matcher-asn-matcher) block (within [`api_protection_rules.api_groups_rules.client_matcher`](#api-protection-rules-api-groups-rules-client-matcher)) supports the following:

&#x2022; [`asn_sets`](#asn-sets) - Optional Block<br>BGP ASN Sets. A list of references to bgp_asn_set objects<br>See [Asn Sets](#api-protection-rules-api-groups-rules-client-matcher-asn-matcher-asn-sets) below.

#### API Protection Rules API Groups Rules Client Matcher Asn Matcher Asn Sets

An [`asn_sets`](#api-protection-rules-api-groups-rules-client-matcher-asn-matcher-asn-sets) block (within [`api_protection_rules.api_groups_rules.client_matcher.asn_matcher`](#api-protection-rules-api-groups-rules-client-matcher-asn-matcher)) supports the following:

&#x2022; [`kind`](#kind) - Optional String<br>Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route')

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

&#x2022; [`uid`](#uid) - Optional String<br>UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid

#### API Protection Rules API Groups Rules Client Matcher Client Selector

A [`client_selector`](#api-protection-rules-api-groups-rules-client-matcher-client-selector) block (within [`api_protection_rules.api_groups_rules.client_matcher`](#api-protection-rules-api-groups-rules-client-matcher)) supports the following:

&#x2022; [`expressions`](#expressions) - Optional List<br>Selector Expression. expressions contains the kubernetes style label expression for selections

#### API Protection Rules API Groups Rules Client Matcher IP Matcher

An [`ip_matcher`](#api-protection-rules-api-groups-rules-client-matcher-ip-matcher) block (within [`api_protection_rules.api_groups_rules.client_matcher`](#api-protection-rules-api-groups-rules-client-matcher)) supports the following:

&#x2022; [`invert_matcher`](#invert-matcher) - Optional Bool<br>Invert IP Matcher. Invert the match result

&#x2022; [`prefix_sets`](#prefix-sets) - Optional Block<br>IP Prefix Sets. A list of references to ip_prefix_set objects<br>See [Prefix Sets](#api-protection-rules-api-groups-rules-client-matcher-ip-matcher-prefix-sets) below.

#### API Protection Rules API Groups Rules Client Matcher IP Matcher Prefix Sets

A [`prefix_sets`](#api-protection-rules-api-groups-rules-client-matcher-ip-matcher-prefix-sets) block (within [`api_protection_rules.api_groups_rules.client_matcher.ip_matcher`](#api-protection-rules-api-groups-rules-client-matcher-ip-matcher)) supports the following:

&#x2022; [`kind`](#kind) - Optional String<br>Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route')

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

&#x2022; [`uid`](#uid) - Optional String<br>UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid

#### API Protection Rules API Groups Rules Client Matcher IP Prefix List

An [`ip_prefix_list`](#api-protection-rules-api-groups-rules-client-matcher-ip-prefix-list) block (within [`api_protection_rules.api_groups_rules.client_matcher`](#api-protection-rules-api-groups-rules-client-matcher)) supports the following:

&#x2022; [`invert_match`](#invert-match) - Optional Bool<br>Invert Match Result. Invert the match result

&#x2022; [`ip_prefixes`](#ip-prefixes) - Optional List<br>IPv4 Prefix List. List of IPv4 prefix strings

#### API Protection Rules API Groups Rules Client Matcher IP Threat Category List

An [`ip_threat_category_list`](#api-protection-rules-api-groups-rules-client-matcher-ip-threat-category-list) block (within [`api_protection_rules.api_groups_rules.client_matcher`](#api-protection-rules-api-groups-rules-client-matcher)) supports the following:

&#x2022; [`ip_threat_categories`](#ip-threat-categories) - Optional List  Defaults to `SPAM_SOURCES`<br>Possible values are `SPAM_SOURCES`, `WINDOWS_EXPLOITS`, `WEB_ATTACKS`, `BOTNETS`, `SCANNERS`, `REPUTATION`, `PHISHING`, `PROXY`, `MOBILE_THREATS`, `TOR_PROXY`, `DENIAL_OF_SERVICE`, `NETWORK`<br>List of IP Threat Categories to choose. The IP threat categories is obtained from the list and is used to auto-generate equivalent label selection expressions

#### API Protection Rules API Groups Rules Client Matcher TLS Fingerprint Matcher

A [`tls_fingerprint_matcher`](#api-protection-rules-api-groups-rules-client-matcher-tls-fingerprint-matcher) block (within [`api_protection_rules.api_groups_rules.client_matcher`](#api-protection-rules-api-groups-rules-client-matcher)) supports the following:

&#x2022; [`classes`](#classes) - Optional List  Defaults to `TLS_FINGERPRINT_NONE`<br>Possible values are `TLS_FINGERPRINT_NONE`, `ANY_MALICIOUS_FINGERPRINT`, `ADWARE`, `ADWIND`, `DRIDEX`, `GOOTKIT`, `GOZI`, `JBIFROST`, `QUAKBOT`, `RANSOMWARE`, `TROLDESH`, `TOFSEE`, `TORRENTLOCKER`, `TRICKBOT`<br>TLS fingerprint classes. A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against

&#x2022; [`exact_values`](#exact-values) - Optional List<br>Exact Values. A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against

&#x2022; [`excluded_values`](#excluded-values) - Optional List<br>Excluded Values. A list of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can be used to skip known false positives when using one or more known TLS fingerprint classes in the enclosing matcher

#### API Protection Rules API Groups Rules Metadata

A [`metadata`](#api-protection-rules-api-groups-rules-metadata) block (within [`api_protection_rules.api_groups_rules`](#api-protection-rules-api-groups-rules)) supports the following:

&#x2022; [`description`](#description) - Optional String<br>Description. Human readable description

&#x2022; [`name`](#name) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### API Protection Rules API Groups Rules Request Matcher

A [`request_matcher`](#api-protection-rules-api-groups-rules-request-matcher) block (within [`api_protection_rules.api_groups_rules`](#api-protection-rules-api-groups-rules)) supports the following:

&#x2022; [`cookie_matchers`](#cookie-matchers) - Optional Block<br>Cookie Matchers. A list of predicates for all cookies that need to be matched. The criteria for matching each cookie is described in individual instances of CookieMatcherType. The actual cookie values are extracted from the request API as a list of strings for each cookie name. Note that all specified cookie matcher predicates must evaluate to true<br>See [Cookie Matchers](#api-protection-rules-api-groups-rules-request-matcher-cookie-matchers) below.

&#x2022; [`headers`](#headers) - Optional Block<br>HTTP Headers. A list of predicates for various HTTP headers that need to match. The criteria for matching each HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values are extracted from the request API as a list of strings for each HTTP header type. Note that all specified header predicates must evaluate to true<br>See [Headers](#api-protection-rules-api-groups-rules-request-matcher-headers) below.

&#x2022; [`jwt_claims`](#jwt-claims) - Optional Block<br>JWT Claims. A list of predicates for various JWT claims that need to match. The criteria for matching each JWT claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are extracted from the JWT payload as a list of strings. Note that all specified JWT claim predicates must evaluate to true. Note that this feature only works on LBs with JWT Validation feature enabled<br>See [JWT Claims](#api-protection-rules-api-groups-rules-request-matcher-jwt-claims) below.

&#x2022; [`query_params`](#query-params) - Optional Block<br>HTTP Query Parameters. A list of predicates for all query parameters that need to be matched. The criteria for matching each query parameter are described in individual instances of QueryParameterMatcherType. The actual query parameter values are extracted from the request API as a list of strings for each query parameter name. Note that all specified query parameter predicates must evaluate to true<br>See [Query Params](#api-protection-rules-api-groups-rules-request-matcher-query-params) below.

#### API Protection Rules API Groups Rules Request Matcher Cookie Matchers

A [`cookie_matchers`](#api-protection-rules-api-groups-rules-request-matcher-cookie-matchers) block (within [`api_protection_rules.api_groups_rules.request_matcher`](#api-protection-rules-api-groups-rules-request-matcher)) supports the following:

&#x2022; [`check_not_present`](#check-not-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`check_present`](#check-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`invert_matcher`](#invert-matcher) - Optional Bool<br>Invert Matcher. Invert Match of the expression defined

&#x2022; [`item`](#item) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#api-protection-rules-api-groups-rules-request-matcher-cookie-matchers-item) below.

&#x2022; [`name`](#name) - Optional String<br>Cookie Name. A case-sensitive cookie name

#### API Protection Rules API Groups Rules Request Matcher Cookie Matchers Item

An [`item`](#api-protection-rules-api-groups-rules-request-matcher-cookie-matchers-item) block (within [`api_protection_rules.api_groups_rules.request_matcher.cookie_matchers`](#api-protection-rules-api-groups-rules-request-matcher-cookie-matchers)) supports the following:

&#x2022; [`exact_values`](#exact-values) - Optional List<br>Exact Values. A list of exact values to match the input against

&#x2022; [`regex_values`](#regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

&#x2022; [`transformers`](#transformers) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### API Protection Rules API Groups Rules Request Matcher Headers

A [`headers`](#api-protection-rules-api-groups-rules-request-matcher-headers) block (within [`api_protection_rules.api_groups_rules.request_matcher`](#api-protection-rules-api-groups-rules-request-matcher)) supports the following:

&#x2022; [`check_not_present`](#check-not-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`check_present`](#check-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`invert_matcher`](#invert-matcher) - Optional Bool<br>Invert Header Matcher. Invert the match result

&#x2022; [`item`](#item) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#api-protection-rules-api-groups-rules-request-matcher-headers-item) below.

&#x2022; [`name`](#name) - Optional String<br>Header Name. A case-insensitive HTTP header name

#### API Protection Rules API Groups Rules Request Matcher Headers Item

An [`item`](#api-protection-rules-api-groups-rules-request-matcher-headers-item) block (within [`api_protection_rules.api_groups_rules.request_matcher.headers`](#api-protection-rules-api-groups-rules-request-matcher-headers)) supports the following:

&#x2022; [`exact_values`](#exact-values) - Optional List<br>Exact Values. A list of exact values to match the input against

&#x2022; [`regex_values`](#regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

&#x2022; [`transformers`](#transformers) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### API Protection Rules API Groups Rules Request Matcher JWT Claims

A [`jwt_claims`](#api-protection-rules-api-groups-rules-request-matcher-jwt-claims) block (within [`api_protection_rules.api_groups_rules.request_matcher`](#api-protection-rules-api-groups-rules-request-matcher)) supports the following:

&#x2022; [`check_not_present`](#check-not-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`check_present`](#check-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`invert_matcher`](#invert-matcher) - Optional Bool<br>Invert Matcher. Invert the match result

&#x2022; [`item`](#item) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#api-protection-rules-api-groups-rules-request-matcher-jwt-claims-item) below.

&#x2022; [`name`](#name) - Optional String<br>JWT Claim Name. JWT claim name

#### API Protection Rules API Groups Rules Request Matcher JWT Claims Item

An [`item`](#api-protection-rules-api-groups-rules-request-matcher-jwt-claims-item) block (within [`api_protection_rules.api_groups_rules.request_matcher.jwt_claims`](#api-protection-rules-api-groups-rules-request-matcher-jwt-claims)) supports the following:

&#x2022; [`exact_values`](#exact-values) - Optional List<br>Exact Values. A list of exact values to match the input against

&#x2022; [`regex_values`](#regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

&#x2022; [`transformers`](#transformers) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### API Protection Rules API Groups Rules Request Matcher Query Params

A [`query_params`](#api-protection-rules-api-groups-rules-request-matcher-query-params) block (within [`api_protection_rules.api_groups_rules.request_matcher`](#api-protection-rules-api-groups-rules-request-matcher)) supports the following:

&#x2022; [`check_not_present`](#check-not-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`check_present`](#check-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`invert_matcher`](#invert-matcher) - Optional Bool<br>Invert Query Parameter Matcher. Invert the match result

&#x2022; [`item`](#item) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#api-protection-rules-api-groups-rules-request-matcher-query-params-item) below.

&#x2022; [`key`](#key) - Optional String<br>Query Parameter Name. A case-sensitive HTTP query parameter name

#### API Protection Rules API Groups Rules Request Matcher Query Params Item

An [`item`](#api-protection-rules-api-groups-rules-request-matcher-query-params-item) block (within [`api_protection_rules.api_groups_rules.request_matcher.query_params`](#api-protection-rules-api-groups-rules-request-matcher-query-params)) supports the following:

&#x2022; [`exact_values`](#exact-values) - Optional List<br>Exact Values. A list of exact values to match the input against

&#x2022; [`regex_values`](#regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

&#x2022; [`transformers`](#transformers) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### API Rate Limit

An [`api_rate_limit`](#api-rate-limit) block supports the following:

&#x2022; [`api_endpoint_rules`](#api-endpoint-rules) - Optional Block<br>API Endpoints. Sets of rules for a specific endpoints. Order is matter as it uses first match policy. For creating rule that contain a whole domain or group of endpoints, please use the server URL rules above<br>See [API Endpoint Rules](#api-rate-limit-api-endpoint-rules) below.

&#x2022; [`bypass_rate_limiting_rules`](#bypass-rate-limiting-rules) - Optional Block<br>Bypass Rate Limiting. This category defines rules per URL or API group. If request matches any of these rules, skip Rate Limiting<br>See [Bypass Rate Limiting Rules](#api-rate-limit-bypass-rate-limiting-rules) below.

&#x2022; [`custom_ip_allowed_list`](#custom-ip-allowed-list) - Optional Block<br>Custom IP Allowed List. IP Allowed list using existing ip_prefix_set objects<br>See [Custom IP Allowed List](#api-rate-limit-custom-ip-allowed-list) below.

&#x2022; [`ip_allowed_list`](#ip-allowed-list) - Optional Block<br>IPv4 Prefix List. x-example: '192.168.20.0/24' List of IPv4 prefixes that represent an endpoint<br>See [IP Allowed List](#api-rate-limit-ip-allowed-list) below.

&#x2022; [`no_ip_allowed_list`](#no-ip-allowed-list) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`server_url_rules`](#server-url-rules) - Optional Block<br>Server URLs. Set of rules for entire domain or base path that contain multiple endpoints. Order is matter as it uses first match policy. For matching also specific endpoints you can use the API endpoint rules set bellow<br>See [Server URL Rules](#api-rate-limit-server-url-rules) below.

#### API Rate Limit API Endpoint Rules

An [`api_endpoint_rules`](#api-rate-limit-api-endpoint-rules) block (within [`api_rate_limit`](#api-rate-limit)) supports the following:

&#x2022; [`any_domain`](#any-domain) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`api_endpoint_method`](#api-endpoint-method) - Optional Block<br>HTTP Method Matcher. A HTTP method matcher specifies a list of methods to match an input HTTP method. The match is considered successful if the input method is a member of the list. The result of the match based on the method list is inverted if invert_matcher is true<br>See [API Endpoint Method](#api-rate-limit-api-endpoint-rules-api-endpoint-method) below.

&#x2022; [`api_endpoint_path`](#api-endpoint-path) - Optional String<br>API Endpoint. The endpoint (path) of the request

&#x2022; [`client_matcher`](#client-matcher) - Optional Block<br>Client Matcher. Client conditions for matching a rule<br>See [Client Matcher](#api-rate-limit-api-endpoint-rules-client-matcher) below.

&#x2022; [`inline_rate_limiter`](#inline-rate-limiter) - Optional Block<br>InlineRateLimiter<br>See [Inline Rate Limiter](#api-rate-limit-api-endpoint-rules-inline-rate-limiter) below.

&#x2022; [`ref_rate_limiter`](#ref-rate-limiter) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Ref Rate Limiter](#api-rate-limit-api-endpoint-rules-ref-rate-limiter) below.

&#x2022; [`request_matcher`](#request-matcher) - Optional Block<br>Request Matcher. Request conditions for matching a rule<br>See [Request Matcher](#api-rate-limit-api-endpoint-rules-request-matcher) below.

&#x2022; [`specific_domain`](#specific-domain) - Optional String<br>Specific Domain. The rule will apply for a specific domain

#### API Rate Limit API Endpoint Rules API Endpoint Method

An [`api_endpoint_method`](#api-rate-limit-api-endpoint-rules-api-endpoint-method) block (within [`api_rate_limit.api_endpoint_rules`](#api-rate-limit-api-endpoint-rules)) supports the following:

&#x2022; [`invert_matcher`](#invert-matcher) - Optional Bool<br>Invert Method Matcher. Invert the match result

&#x2022; [`methods`](#methods) - Optional List  Defaults to `ANY`<br>Possible values are `ANY`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`, `COPY`<br>Method List. List of methods values to match against

#### API Rate Limit API Endpoint Rules Client Matcher

A [`client_matcher`](#api-rate-limit-api-endpoint-rules-client-matcher) block (within [`api_rate_limit.api_endpoint_rules`](#api-rate-limit-api-endpoint-rules)) supports the following:

&#x2022; [`any_client`](#any-client) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`any_ip`](#any-ip) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`asn_list`](#asn-list) - Optional Block<br>ASN Match List. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer<br>See [Asn List](#api-rate-limit-api-endpoint-rules-client-matcher-asn-list) below.

&#x2022; [`asn_matcher`](#asn-matcher) - Optional Block<br>ASN Matcher. Match any AS number contained in the list of bgp_asn_sets<br>See [Asn Matcher](#api-rate-limit-api-endpoint-rules-client-matcher-asn-matcher) below.

&#x2022; [`client_selector`](#client-selector) - Optional Block<br>Label Selector. This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expresssions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects. Label selector is immutable. expressions is a list of strings of label selection expression. Each string has ',' separated values which are 'AND' and all strings are logically 'OR'. BNF for expression string <selector-syntax> ::= <requirement> | <requirement> ',' <selector-syntax> <requirement> ::= [!] KEY [ <set-based-restriction> | <exact-match-restriction> ] <set-based-restriction> ::= '' | <inclusion-exclusion> <value-set> <inclusion-exclusion> ::= <inclusion> | <exclusion> <exclusion> ::= 'notin' <inclusion> ::= 'in' <value-set> ::= '(' <values> ')' <values> ::= VALUE | VALUE ',' <values> <exact-match-restriction> ::= ['='|'=='|'!='] VALUE<br>See [Client Selector](#api-rate-limit-api-endpoint-rules-client-matcher-client-selector) below.

&#x2022; [`ip_matcher`](#ip-matcher) - Optional Block<br>IP Prefix Matcher. Match any IP prefix contained in the list of ip_prefix_sets. The result of the match is inverted if invert_matcher is true<br>See [IP Matcher](#api-rate-limit-api-endpoint-rules-client-matcher-ip-matcher) below.

&#x2022; [`ip_prefix_list`](#ip-prefix-list) - Optional Block<br>IP Prefix Match List. List of IP Prefix strings to match against<br>See [IP Prefix List](#api-rate-limit-api-endpoint-rules-client-matcher-ip-prefix-list) below.

&#x2022; [`ip_threat_category_list`](#ip-threat-category-list) - Optional Block<br>IP Threat Category List Type. List of IP threat categories<br>See [IP Threat Category List](#api-rate-limit-api-endpoint-rules-client-matcher-ip-threat-category-list) below.

&#x2022; [`tls_fingerprint_matcher`](#tls-fingerprint-matcher) - Optional Block<br>TLS Fingerprint Matcher. A TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of supported positve match criteria includes a list of known classes of TLS fingerprints and a list of exact values. The match is considered successful if either of these positive criteria are satisfied and the input fingerprint is not one of the excluded values<br>See [TLS Fingerprint Matcher](#api-rate-limit-api-endpoint-rules-client-matcher-tls-fingerprint-matcher) below.

#### API Rate Limit API Endpoint Rules Client Matcher Asn List

An [`asn_list`](#api-rate-limit-api-endpoint-rules-client-matcher-asn-list) block (within [`api_rate_limit.api_endpoint_rules.client_matcher`](#api-rate-limit-api-endpoint-rules-client-matcher)) supports the following:

&#x2022; [`as_numbers`](#as-numbers) - Optional List<br>AS Numbers. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer

#### API Rate Limit API Endpoint Rules Client Matcher Asn Matcher

An [`asn_matcher`](#api-rate-limit-api-endpoint-rules-client-matcher-asn-matcher) block (within [`api_rate_limit.api_endpoint_rules.client_matcher`](#api-rate-limit-api-endpoint-rules-client-matcher)) supports the following:

&#x2022; [`asn_sets`](#asn-sets) - Optional Block<br>BGP ASN Sets. A list of references to bgp_asn_set objects<br>See [Asn Sets](#api-rate-limit-api-endpoint-rules-client-matcher-asn-matcher-asn-sets) below.

#### API Rate Limit API Endpoint Rules Client Matcher Asn Matcher Asn Sets

An [`asn_sets`](#api-rate-limit-api-endpoint-rules-client-matcher-asn-matcher-asn-sets) block (within [`api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher`](#api-rate-limit-api-endpoint-rules-client-matcher-asn-matcher)) supports the following:

&#x2022; [`kind`](#kind) - Optional String<br>Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route')

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

&#x2022; [`uid`](#uid) - Optional String<br>UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid

#### API Rate Limit API Endpoint Rules Client Matcher Client Selector

A [`client_selector`](#api-rate-limit-api-endpoint-rules-client-matcher-client-selector) block (within [`api_rate_limit.api_endpoint_rules.client_matcher`](#api-rate-limit-api-endpoint-rules-client-matcher)) supports the following:

&#x2022; [`expressions`](#expressions) - Optional List<br>Selector Expression. expressions contains the kubernetes style label expression for selections

#### API Rate Limit API Endpoint Rules Client Matcher IP Matcher

An [`ip_matcher`](#api-rate-limit-api-endpoint-rules-client-matcher-ip-matcher) block (within [`api_rate_limit.api_endpoint_rules.client_matcher`](#api-rate-limit-api-endpoint-rules-client-matcher)) supports the following:

&#x2022; [`invert_matcher`](#invert-matcher) - Optional Bool<br>Invert IP Matcher. Invert the match result

&#x2022; [`prefix_sets`](#prefix-sets) - Optional Block<br>IP Prefix Sets. A list of references to ip_prefix_set objects<br>See [Prefix Sets](#api-rate-limit-api-endpoint-rules-client-matcher-ip-matcher-prefix-sets) below.

#### API Rate Limit API Endpoint Rules Client Matcher IP Matcher Prefix Sets

A [`prefix_sets`](#api-rate-limit-api-endpoint-rules-client-matcher-ip-matcher-prefix-sets) block (within [`api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher`](#api-rate-limit-api-endpoint-rules-client-matcher-ip-matcher)) supports the following:

&#x2022; [`kind`](#kind) - Optional String<br>Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route')

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

&#x2022; [`uid`](#uid) - Optional String<br>UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid

#### API Rate Limit API Endpoint Rules Client Matcher IP Prefix List

An [`ip_prefix_list`](#api-rate-limit-api-endpoint-rules-client-matcher-ip-prefix-list) block (within [`api_rate_limit.api_endpoint_rules.client_matcher`](#api-rate-limit-api-endpoint-rules-client-matcher)) supports the following:

&#x2022; [`invert_match`](#invert-match) - Optional Bool<br>Invert Match Result. Invert the match result

&#x2022; [`ip_prefixes`](#ip-prefixes) - Optional List<br>IPv4 Prefix List. List of IPv4 prefix strings

#### API Rate Limit API Endpoint Rules Client Matcher IP Threat Category List

An [`ip_threat_category_list`](#api-rate-limit-api-endpoint-rules-client-matcher-ip-threat-category-list) block (within [`api_rate_limit.api_endpoint_rules.client_matcher`](#api-rate-limit-api-endpoint-rules-client-matcher)) supports the following:

&#x2022; [`ip_threat_categories`](#ip-threat-categories) - Optional List  Defaults to `SPAM_SOURCES`<br>Possible values are `SPAM_SOURCES`, `WINDOWS_EXPLOITS`, `WEB_ATTACKS`, `BOTNETS`, `SCANNERS`, `REPUTATION`, `PHISHING`, `PROXY`, `MOBILE_THREATS`, `TOR_PROXY`, `DENIAL_OF_SERVICE`, `NETWORK`<br>List of IP Threat Categories to choose. The IP threat categories is obtained from the list and is used to auto-generate equivalent label selection expressions

#### API Rate Limit API Endpoint Rules Client Matcher TLS Fingerprint Matcher

A [`tls_fingerprint_matcher`](#api-rate-limit-api-endpoint-rules-client-matcher-tls-fingerprint-matcher) block (within [`api_rate_limit.api_endpoint_rules.client_matcher`](#api-rate-limit-api-endpoint-rules-client-matcher)) supports the following:

&#x2022; [`classes`](#classes) - Optional List  Defaults to `TLS_FINGERPRINT_NONE`<br>Possible values are `TLS_FINGERPRINT_NONE`, `ANY_MALICIOUS_FINGERPRINT`, `ADWARE`, `ADWIND`, `DRIDEX`, `GOOTKIT`, `GOZI`, `JBIFROST`, `QUAKBOT`, `RANSOMWARE`, `TROLDESH`, `TOFSEE`, `TORRENTLOCKER`, `TRICKBOT`<br>TLS fingerprint classes. A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against

&#x2022; [`exact_values`](#exact-values) - Optional List<br>Exact Values. A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against

&#x2022; [`excluded_values`](#excluded-values) - Optional List<br>Excluded Values. A list of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can be used to skip known false positives when using one or more known TLS fingerprint classes in the enclosing matcher

#### API Rate Limit API Endpoint Rules Inline Rate Limiter

An [`inline_rate_limiter`](#api-rate-limit-api-endpoint-rules-inline-rate-limiter) block (within [`api_rate_limit.api_endpoint_rules`](#api-rate-limit-api-endpoint-rules)) supports the following:

&#x2022; [`ref_user_id`](#ref-user-id) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Ref User Id](#api-rate-limit-api-endpoint-rules-inline-rate-limiter-ref-user-id) below.

&#x2022; [`threshold`](#threshold) - Optional Number<br>Threshold. The total number of allowed requests for 1 unit (e.g. SECOND/MINUTE/HOUR etc.) of the specified period

&#x2022; [`unit`](#unit) - Optional String  Defaults to `SECOND`<br>Possible values are `SECOND`, `MINUTE`, `HOUR`<br>Rate Limit Period Unit. Unit for the period per which the rate limit is applied. - SECOND: Second Rate limit period unit is seconds - MINUTE: Minute Rate limit period unit is minutes - HOUR: Hour Rate limit period unit is hours - DAY: Day Rate limit period unit is days

&#x2022; [`use_http_lb_user_id`](#use-http-lb-user-id) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### API Rate Limit API Endpoint Rules Inline Rate Limiter Ref User Id

A [`ref_user_id`](#api-rate-limit-api-endpoint-rules-inline-rate-limiter-ref-user-id) block (within [`api_rate_limit.api_endpoint_rules.inline_rate_limiter`](#api-rate-limit-api-endpoint-rules-inline-rate-limiter)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### API Rate Limit API Endpoint Rules Ref Rate Limiter

A [`ref_rate_limiter`](#api-rate-limit-api-endpoint-rules-ref-rate-limiter) block (within [`api_rate_limit.api_endpoint_rules`](#api-rate-limit-api-endpoint-rules)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### API Rate Limit API Endpoint Rules Request Matcher

A [`request_matcher`](#api-rate-limit-api-endpoint-rules-request-matcher) block (within [`api_rate_limit.api_endpoint_rules`](#api-rate-limit-api-endpoint-rules)) supports the following:

&#x2022; [`cookie_matchers`](#cookie-matchers) - Optional Block<br>Cookie Matchers. A list of predicates for all cookies that need to be matched. The criteria for matching each cookie is described in individual instances of CookieMatcherType. The actual cookie values are extracted from the request API as a list of strings for each cookie name. Note that all specified cookie matcher predicates must evaluate to true<br>See [Cookie Matchers](#api-rate-limit-api-endpoint-rules-request-matcher-cookie-matchers) below.

&#x2022; [`headers`](#headers) - Optional Block<br>HTTP Headers. A list of predicates for various HTTP headers that need to match. The criteria for matching each HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values are extracted from the request API as a list of strings for each HTTP header type. Note that all specified header predicates must evaluate to true<br>See [Headers](#api-rate-limit-api-endpoint-rules-request-matcher-headers) below.

&#x2022; [`jwt_claims`](#jwt-claims) - Optional Block<br>JWT Claims. A list of predicates for various JWT claims that need to match. The criteria for matching each JWT claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are extracted from the JWT payload as a list of strings. Note that all specified JWT claim predicates must evaluate to true. Note that this feature only works on LBs with JWT Validation feature enabled<br>See [JWT Claims](#api-rate-limit-api-endpoint-rules-request-matcher-jwt-claims) below.

&#x2022; [`query_params`](#query-params) - Optional Block<br>HTTP Query Parameters. A list of predicates for all query parameters that need to be matched. The criteria for matching each query parameter are described in individual instances of QueryParameterMatcherType. The actual query parameter values are extracted from the request API as a list of strings for each query parameter name. Note that all specified query parameter predicates must evaluate to true<br>See [Query Params](#api-rate-limit-api-endpoint-rules-request-matcher-query-params) below.

#### API Rate Limit API Endpoint Rules Request Matcher Cookie Matchers

A [`cookie_matchers`](#api-rate-limit-api-endpoint-rules-request-matcher-cookie-matchers) block (within [`api_rate_limit.api_endpoint_rules.request_matcher`](#api-rate-limit-api-endpoint-rules-request-matcher)) supports the following:

&#x2022; [`check_not_present`](#check-not-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`check_present`](#check-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`invert_matcher`](#invert-matcher) - Optional Bool<br>Invert Matcher. Invert Match of the expression defined

&#x2022; [`item`](#item) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#api-rate-limit-api-endpoint-rules-request-matcher-cookie-matchers-item) below.

&#x2022; [`name`](#name) - Optional String<br>Cookie Name. A case-sensitive cookie name

#### API Rate Limit API Endpoint Rules Request Matcher Cookie Matchers Item

An [`item`](#api-rate-limit-api-endpoint-rules-request-matcher-cookie-matchers-item) block (within [`api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers`](#api-rate-limit-api-endpoint-rules-request-matcher-cookie-matchers)) supports the following:

&#x2022; [`exact_values`](#exact-values) - Optional List<br>Exact Values. A list of exact values to match the input against

&#x2022; [`regex_values`](#regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

&#x2022; [`transformers`](#transformers) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### API Rate Limit API Endpoint Rules Request Matcher Headers

A [`headers`](#api-rate-limit-api-endpoint-rules-request-matcher-headers) block (within [`api_rate_limit.api_endpoint_rules.request_matcher`](#api-rate-limit-api-endpoint-rules-request-matcher)) supports the following:

&#x2022; [`check_not_present`](#check-not-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`check_present`](#check-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`invert_matcher`](#invert-matcher) - Optional Bool<br>Invert Header Matcher. Invert the match result

&#x2022; [`item`](#item) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#api-rate-limit-api-endpoint-rules-request-matcher-headers-item) below.

&#x2022; [`name`](#name) - Optional String<br>Header Name. A case-insensitive HTTP header name

#### API Rate Limit API Endpoint Rules Request Matcher Headers Item

An [`item`](#api-rate-limit-api-endpoint-rules-request-matcher-headers-item) block (within [`api_rate_limit.api_endpoint_rules.request_matcher.headers`](#api-rate-limit-api-endpoint-rules-request-matcher-headers)) supports the following:

&#x2022; [`exact_values`](#exact-values) - Optional List<br>Exact Values. A list of exact values to match the input against

&#x2022; [`regex_values`](#regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

&#x2022; [`transformers`](#transformers) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### API Rate Limit API Endpoint Rules Request Matcher JWT Claims

A [`jwt_claims`](#api-rate-limit-api-endpoint-rules-request-matcher-jwt-claims) block (within [`api_rate_limit.api_endpoint_rules.request_matcher`](#api-rate-limit-api-endpoint-rules-request-matcher)) supports the following:

&#x2022; [`check_not_present`](#check-not-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`check_present`](#check-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`invert_matcher`](#invert-matcher) - Optional Bool<br>Invert Matcher. Invert the match result

&#x2022; [`item`](#item) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#api-rate-limit-api-endpoint-rules-request-matcher-jwt-claims-item) below.

&#x2022; [`name`](#name) - Optional String<br>JWT Claim Name. JWT claim name

#### API Rate Limit API Endpoint Rules Request Matcher JWT Claims Item

An [`item`](#api-rate-limit-api-endpoint-rules-request-matcher-jwt-claims-item) block (within [`api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims`](#api-rate-limit-api-endpoint-rules-request-matcher-jwt-claims)) supports the following:

&#x2022; [`exact_values`](#exact-values) - Optional List<br>Exact Values. A list of exact values to match the input against

&#x2022; [`regex_values`](#regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

&#x2022; [`transformers`](#transformers) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### API Rate Limit API Endpoint Rules Request Matcher Query Params

A [`query_params`](#api-rate-limit-api-endpoint-rules-request-matcher-query-params) block (within [`api_rate_limit.api_endpoint_rules.request_matcher`](#api-rate-limit-api-endpoint-rules-request-matcher)) supports the following:

&#x2022; [`check_not_present`](#check-not-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`check_present`](#check-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`invert_matcher`](#invert-matcher) - Optional Bool<br>Invert Query Parameter Matcher. Invert the match result

&#x2022; [`item`](#item) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#api-rate-limit-api-endpoint-rules-request-matcher-query-params-item) below.

&#x2022; [`key`](#key) - Optional String<br>Query Parameter Name. A case-sensitive HTTP query parameter name

#### API Rate Limit API Endpoint Rules Request Matcher Query Params Item

An [`item`](#api-rate-limit-api-endpoint-rules-request-matcher-query-params-item) block (within [`api_rate_limit.api_endpoint_rules.request_matcher.query_params`](#api-rate-limit-api-endpoint-rules-request-matcher-query-params)) supports the following:

&#x2022; [`exact_values`](#exact-values) - Optional List<br>Exact Values. A list of exact values to match the input against

&#x2022; [`regex_values`](#regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

&#x2022; [`transformers`](#transformers) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### API Rate Limit Bypass Rate Limiting Rules

A [`bypass_rate_limiting_rules`](#api-rate-limit-bypass-rate-limiting-rules) block (within [`api_rate_limit`](#api-rate-limit)) supports the following:

&#x2022; [`bypass_rate_limiting_rules`](#bypass-rate-limiting-rules) - Optional Block<br>Bypass Rate Limiting. This category defines rules per URL or API group. If request matches any of these rules, skip Rate Limiting<br>See [Bypass Rate Limiting Rules](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules) below.

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules

A [`bypass_rate_limiting_rules`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules) block (within [`api_rate_limit.bypass_rate_limiting_rules`](#api-rate-limit-bypass-rate-limiting-rules)) supports the following:

&#x2022; [`any_domain`](#any-domain) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`any_url`](#any-url) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`api_endpoint`](#api-endpoint) - Optional Block<br>API Endpoint. This defines API endpoint<br>See [API Endpoint](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-api-endpoint) below.

&#x2022; [`api_groups`](#api-groups) - Optional Block<br>API Groups<br>See [API Groups](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-api-groups) below.

&#x2022; [`base_path`](#base-path) - Optional String<br>Base Path. The base path which this validation applies to

&#x2022; [`client_matcher`](#client-matcher) - Optional Block<br>Client Matcher. Client conditions for matching a rule<br>See [Client Matcher](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher) below.

&#x2022; [`request_matcher`](#request-matcher) - Optional Block<br>Request Matcher. Request conditions for matching a rule<br>See [Request Matcher](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher) below.

&#x2022; [`specific_domain`](#specific-domain) - Optional String<br>Specific Domain. The rule will apply for a specific domain. For example: API.example.com

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules API Endpoint

An [`api_endpoint`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-api-endpoint) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules)) supports the following:

&#x2022; [`methods`](#methods) - Optional List  Defaults to `ANY`<br>Possible values are `ANY`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`, `COPY`<br>Methods. Methods to be matched

&#x2022; [`path`](#path) - Optional String<br>Path. Path to be matched

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules API Groups

An [`api_groups`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-api-groups) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules)) supports the following:

&#x2022; [`api_groups`](#api-groups) - Optional List<br>API Groups

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Client Matcher

A [`client_matcher`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules)) supports the following:

&#x2022; [`any_client`](#any-client) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`any_ip`](#any-ip) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`asn_list`](#asn-list) - Optional Block<br>ASN Match List. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer<br>See [Asn List](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-asn-list) below.

&#x2022; [`asn_matcher`](#asn-matcher) - Optional Block<br>ASN Matcher. Match any AS number contained in the list of bgp_asn_sets<br>See [Asn Matcher](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-asn-matcher) below.

&#x2022; [`client_selector`](#client-selector) - Optional Block<br>Label Selector. This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expresssions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects. Label selector is immutable. expressions is a list of strings of label selection expression. Each string has ',' separated values which are 'AND' and all strings are logically 'OR'. BNF for expression string <selector-syntax> ::= <requirement> | <requirement> ',' <selector-syntax> <requirement> ::= [!] KEY [ <set-based-restriction> | <exact-match-restriction> ] <set-based-restriction> ::= '' | <inclusion-exclusion> <value-set> <inclusion-exclusion> ::= <inclusion> | <exclusion> <exclusion> ::= 'notin' <inclusion> ::= 'in' <value-set> ::= '(' <values> ')' <values> ::= VALUE | VALUE ',' <values> <exact-match-restriction> ::= ['='|'=='|'!='] VALUE<br>See [Client Selector](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-client-selector) below.

&#x2022; [`ip_matcher`](#ip-matcher) - Optional Block<br>IP Prefix Matcher. Match any IP prefix contained in the list of ip_prefix_sets. The result of the match is inverted if invert_matcher is true<br>See [IP Matcher](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-ip-matcher) below.

&#x2022; [`ip_prefix_list`](#ip-prefix-list) - Optional Block<br>IP Prefix Match List. List of IP Prefix strings to match against<br>See [IP Prefix List](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-ip-prefix-list) below.

&#x2022; [`ip_threat_category_list`](#ip-threat-category-list) - Optional Block<br>IP Threat Category List Type. List of IP threat categories<br>See [IP Threat Category List](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-ip-threat-category-list) below.

&#x2022; [`tls_fingerprint_matcher`](#tls-fingerprint-matcher) - Optional Block<br>TLS Fingerprint Matcher. A TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of supported positve match criteria includes a list of known classes of TLS fingerprints and a list of exact values. The match is considered successful if either of these positive criteria are satisfied and the input fingerprint is not one of the excluded values<br>See [TLS Fingerprint Matcher](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-tls-fingerprint-matcher) below.

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Client Matcher Asn List

An [`asn_list`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-asn-list) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher)) supports the following:

&#x2022; [`as_numbers`](#as-numbers) - Optional List<br>AS Numbers. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Client Matcher Asn Matcher

An [`asn_matcher`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-asn-matcher) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher)) supports the following:

&#x2022; [`asn_sets`](#asn-sets) - Optional Block<br>BGP ASN Sets. A list of references to bgp_asn_set objects<br>See [Asn Sets](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-asn-matcher-asn-sets) below.

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Client Matcher Asn Matcher Asn Sets

An [`asn_sets`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-asn-matcher-asn-sets) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-asn-matcher)) supports the following:

&#x2022; [`kind`](#kind) - Optional String<br>Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route')

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

&#x2022; [`uid`](#uid) - Optional String<br>UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Client Matcher Client Selector

A [`client_selector`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-client-selector) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher)) supports the following:

&#x2022; [`expressions`](#expressions) - Optional List<br>Selector Expression. expressions contains the kubernetes style label expression for selections

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Client Matcher IP Matcher

An [`ip_matcher`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-ip-matcher) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher)) supports the following:

&#x2022; [`invert_matcher`](#invert-matcher) - Optional Bool<br>Invert IP Matcher. Invert the match result

&#x2022; [`prefix_sets`](#prefix-sets) - Optional Block<br>IP Prefix Sets. A list of references to ip_prefix_set objects<br>See [Prefix Sets](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-ip-matcher-prefix-sets) below.

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Client Matcher IP Matcher Prefix Sets

A [`prefix_sets`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-ip-matcher-prefix-sets) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-ip-matcher)) supports the following:

&#x2022; [`kind`](#kind) - Optional String<br>Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route')

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

&#x2022; [`uid`](#uid) - Optional String<br>UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Client Matcher IP Prefix List

An [`ip_prefix_list`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-ip-prefix-list) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher)) supports the following:

&#x2022; [`invert_match`](#invert-match) - Optional Bool<br>Invert Match Result. Invert the match result

&#x2022; [`ip_prefixes`](#ip-prefixes) - Optional List<br>IPv4 Prefix List. List of IPv4 prefix strings

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Client Matcher IP Threat Category List

An [`ip_threat_category_list`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-ip-threat-category-list) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher)) supports the following:

&#x2022; [`ip_threat_categories`](#ip-threat-categories) - Optional List  Defaults to `SPAM_SOURCES`<br>Possible values are `SPAM_SOURCES`, `WINDOWS_EXPLOITS`, `WEB_ATTACKS`, `BOTNETS`, `SCANNERS`, `REPUTATION`, `PHISHING`, `PROXY`, `MOBILE_THREATS`, `TOR_PROXY`, `DENIAL_OF_SERVICE`, `NETWORK`<br>List of IP Threat Categories to choose. The IP threat categories is obtained from the list and is used to auto-generate equivalent label selection expressions

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Client Matcher TLS Fingerprint Matcher

A [`tls_fingerprint_matcher`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-tls-fingerprint-matcher) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher)) supports the following:

&#x2022; [`classes`](#classes) - Optional List  Defaults to `TLS_FINGERPRINT_NONE`<br>Possible values are `TLS_FINGERPRINT_NONE`, `ANY_MALICIOUS_FINGERPRINT`, `ADWARE`, `ADWIND`, `DRIDEX`, `GOOTKIT`, `GOZI`, `JBIFROST`, `QUAKBOT`, `RANSOMWARE`, `TROLDESH`, `TOFSEE`, `TORRENTLOCKER`, `TRICKBOT`<br>TLS fingerprint classes. A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against

&#x2022; [`exact_values`](#exact-values) - Optional List<br>Exact Values. A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against

&#x2022; [`excluded_values`](#excluded-values) - Optional List<br>Excluded Values. A list of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can be used to skip known false positives when using one or more known TLS fingerprint classes in the enclosing matcher

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Request Matcher

A [`request_matcher`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules)) supports the following:

&#x2022; [`cookie_matchers`](#cookie-matchers) - Optional Block<br>Cookie Matchers. A list of predicates for all cookies that need to be matched. The criteria for matching each cookie is described in individual instances of CookieMatcherType. The actual cookie values are extracted from the request API as a list of strings for each cookie name. Note that all specified cookie matcher predicates must evaluate to true<br>See [Cookie Matchers](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-cookie-matchers) below.

&#x2022; [`headers`](#headers) - Optional Block<br>HTTP Headers. A list of predicates for various HTTP headers that need to match. The criteria for matching each HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values are extracted from the request API as a list of strings for each HTTP header type. Note that all specified header predicates must evaluate to true<br>See [Headers](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-headers) below.

&#x2022; [`jwt_claims`](#jwt-claims) - Optional Block<br>JWT Claims. A list of predicates for various JWT claims that need to match. The criteria for matching each JWT claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are extracted from the JWT payload as a list of strings. Note that all specified JWT claim predicates must evaluate to true. Note that this feature only works on LBs with JWT Validation feature enabled<br>See [JWT Claims](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-jwt-claims) below.

&#x2022; [`query_params`](#query-params) - Optional Block<br>HTTP Query Parameters. A list of predicates for all query parameters that need to be matched. The criteria for matching each query parameter are described in individual instances of QueryParameterMatcherType. The actual query parameter values are extracted from the request API as a list of strings for each query parameter name. Note that all specified query parameter predicates must evaluate to true<br>See [Query Params](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-query-params) below.

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Request Matcher Cookie Matchers

A [`cookie_matchers`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-cookie-matchers) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher)) supports the following:

&#x2022; [`check_not_present`](#check-not-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`check_present`](#check-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`invert_matcher`](#invert-matcher) - Optional Bool<br>Invert Matcher. Invert Match of the expression defined

&#x2022; [`item`](#item) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-cookie-matchers-item) below.

&#x2022; [`name`](#name) - Optional String<br>Cookie Name. A case-sensitive cookie name

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Request Matcher Cookie Matchers Item

An [`item`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-cookie-matchers-item) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-cookie-matchers)) supports the following:

&#x2022; [`exact_values`](#exact-values) - Optional List<br>Exact Values. A list of exact values to match the input against

&#x2022; [`regex_values`](#regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

&#x2022; [`transformers`](#transformers) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Request Matcher Headers

A [`headers`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-headers) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher)) supports the following:

&#x2022; [`check_not_present`](#check-not-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`check_present`](#check-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`invert_matcher`](#invert-matcher) - Optional Bool<br>Invert Header Matcher. Invert the match result

&#x2022; [`item`](#item) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-headers-item) below.

&#x2022; [`name`](#name) - Optional String<br>Header Name. A case-insensitive HTTP header name

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Request Matcher Headers Item

An [`item`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-headers-item) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-headers)) supports the following:

&#x2022; [`exact_values`](#exact-values) - Optional List<br>Exact Values. A list of exact values to match the input against

&#x2022; [`regex_values`](#regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

&#x2022; [`transformers`](#transformers) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Request Matcher JWT Claims

A [`jwt_claims`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-jwt-claims) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher)) supports the following:

&#x2022; [`check_not_present`](#check-not-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`check_present`](#check-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`invert_matcher`](#invert-matcher) - Optional Bool<br>Invert Matcher. Invert the match result

&#x2022; [`item`](#item) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-jwt-claims-item) below.

&#x2022; [`name`](#name) - Optional String<br>JWT Claim Name. JWT claim name

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Request Matcher JWT Claims Item

An [`item`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-jwt-claims-item) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-jwt-claims)) supports the following:

&#x2022; [`exact_values`](#exact-values) - Optional List<br>Exact Values. A list of exact values to match the input against

&#x2022; [`regex_values`](#regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

&#x2022; [`transformers`](#transformers) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Request Matcher Query Params

A [`query_params`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-query-params) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher)) supports the following:

&#x2022; [`check_not_present`](#check-not-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`check_present`](#check-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`invert_matcher`](#invert-matcher) - Optional Bool<br>Invert Query Parameter Matcher. Invert the match result

&#x2022; [`item`](#item) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-query-params-item) below.

&#x2022; [`key`](#key) - Optional String<br>Query Parameter Name. A case-sensitive HTTP query parameter name

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Request Matcher Query Params Item

An [`item`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-query-params-item) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-query-params)) supports the following:

&#x2022; [`exact_values`](#exact-values) - Optional List<br>Exact Values. A list of exact values to match the input against

&#x2022; [`regex_values`](#regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

&#x2022; [`transformers`](#transformers) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### API Rate Limit Custom IP Allowed List

A [`custom_ip_allowed_list`](#api-rate-limit-custom-ip-allowed-list) block (within [`api_rate_limit`](#api-rate-limit)) supports the following:

&#x2022; [`rate_limiter_allowed_prefixes`](#rate-limiter-allowed-prefixes) - Optional Block<br>List of IP Prefix Sets. References to ip_prefix_set objects. Requests from source IP addresses that are covered by one of the allowed IP Prefixes are not subjected to rate limiting<br>See [Rate Limiter Allowed Prefixes](#api-rate-limit-custom-ip-allowed-list-rate-limiter-allowed-prefixes) below.

#### API Rate Limit Custom IP Allowed List Rate Limiter Allowed Prefixes

A [`rate_limiter_allowed_prefixes`](#api-rate-limit-custom-ip-allowed-list-rate-limiter-allowed-prefixes) block (within [`api_rate_limit.custom_ip_allowed_list`](#api-rate-limit-custom-ip-allowed-list)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### API Rate Limit IP Allowed List

An [`ip_allowed_list`](#api-rate-limit-ip-allowed-list) block (within [`api_rate_limit`](#api-rate-limit)) supports the following:

&#x2022; [`prefixes`](#prefixes) - Optional List<br>IPv4 Prefix List. List of IPv4 prefixes that represent an endpoint

#### API Rate Limit Server URL Rules

A [`server_url_rules`](#api-rate-limit-server-url-rules) block (within [`api_rate_limit`](#api-rate-limit)) supports the following:

&#x2022; [`any_domain`](#any-domain) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`api_group`](#api-group) - Optional String<br>API Group. API groups derived from API Definition swaggers. For example oas-all-operations including all paths and methods from the swaggers, oas-base-urls covering all requests under base-paths from the swaggers. Custom groups can be created if user tags paths or operations with 'x-volterra-API-group' extensions inside swaggers

&#x2022; [`base_path`](#base-path) - Optional String<br>Base Path. Prefix of the request path

&#x2022; [`client_matcher`](#client-matcher) - Optional Block<br>Client Matcher. Client conditions for matching a rule<br>See [Client Matcher](#api-rate-limit-server-url-rules-client-matcher) below.

&#x2022; [`inline_rate_limiter`](#inline-rate-limiter) - Optional Block<br>InlineRateLimiter<br>See [Inline Rate Limiter](#api-rate-limit-server-url-rules-inline-rate-limiter) below.

&#x2022; [`ref_rate_limiter`](#ref-rate-limiter) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Ref Rate Limiter](#api-rate-limit-server-url-rules-ref-rate-limiter) below.

&#x2022; [`request_matcher`](#request-matcher) - Optional Block<br>Request Matcher. Request conditions for matching a rule<br>See [Request Matcher](#api-rate-limit-server-url-rules-request-matcher) below.

&#x2022; [`specific_domain`](#specific-domain) - Optional String<br>Specific Domain. The rule will apply for a specific domain

#### API Rate Limit Server URL Rules Client Matcher

A [`client_matcher`](#api-rate-limit-server-url-rules-client-matcher) block (within [`api_rate_limit.server_url_rules`](#api-rate-limit-server-url-rules)) supports the following:

&#x2022; [`any_client`](#any-client) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`any_ip`](#any-ip) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`asn_list`](#asn-list) - Optional Block<br>ASN Match List. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer<br>See [Asn List](#api-rate-limit-server-url-rules-client-matcher-asn-list) below.

&#x2022; [`asn_matcher`](#asn-matcher) - Optional Block<br>ASN Matcher. Match any AS number contained in the list of bgp_asn_sets<br>See [Asn Matcher](#api-rate-limit-server-url-rules-client-matcher-asn-matcher) below.

&#x2022; [`client_selector`](#client-selector) - Optional Block<br>Label Selector. This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expresssions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects. Label selector is immutable. expressions is a list of strings of label selection expression. Each string has ',' separated values which are 'AND' and all strings are logically 'OR'. BNF for expression string <selector-syntax> ::= <requirement> | <requirement> ',' <selector-syntax> <requirement> ::= [!] KEY [ <set-based-restriction> | <exact-match-restriction> ] <set-based-restriction> ::= '' | <inclusion-exclusion> <value-set> <inclusion-exclusion> ::= <inclusion> | <exclusion> <exclusion> ::= 'notin' <inclusion> ::= 'in' <value-set> ::= '(' <values> ')' <values> ::= VALUE | VALUE ',' <values> <exact-match-restriction> ::= ['='|'=='|'!='] VALUE<br>See [Client Selector](#api-rate-limit-server-url-rules-client-matcher-client-selector) below.

&#x2022; [`ip_matcher`](#ip-matcher) - Optional Block<br>IP Prefix Matcher. Match any IP prefix contained in the list of ip_prefix_sets. The result of the match is inverted if invert_matcher is true<br>See [IP Matcher](#api-rate-limit-server-url-rules-client-matcher-ip-matcher) below.

&#x2022; [`ip_prefix_list`](#ip-prefix-list) - Optional Block<br>IP Prefix Match List. List of IP Prefix strings to match against<br>See [IP Prefix List](#api-rate-limit-server-url-rules-client-matcher-ip-prefix-list) below.

&#x2022; [`ip_threat_category_list`](#ip-threat-category-list) - Optional Block<br>IP Threat Category List Type. List of IP threat categories<br>See [IP Threat Category List](#api-rate-limit-server-url-rules-client-matcher-ip-threat-category-list) below.

&#x2022; [`tls_fingerprint_matcher`](#tls-fingerprint-matcher) - Optional Block<br>TLS Fingerprint Matcher. A TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of supported positve match criteria includes a list of known classes of TLS fingerprints and a list of exact values. The match is considered successful if either of these positive criteria are satisfied and the input fingerprint is not one of the excluded values<br>See [TLS Fingerprint Matcher](#api-rate-limit-server-url-rules-client-matcher-tls-fingerprint-matcher) below.

#### API Rate Limit Server URL Rules Client Matcher Asn List

An [`asn_list`](#api-rate-limit-server-url-rules-client-matcher-asn-list) block (within [`api_rate_limit.server_url_rules.client_matcher`](#api-rate-limit-server-url-rules-client-matcher)) supports the following:

&#x2022; [`as_numbers`](#as-numbers) - Optional List<br>AS Numbers. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer

#### API Rate Limit Server URL Rules Client Matcher Asn Matcher

An [`asn_matcher`](#api-rate-limit-server-url-rules-client-matcher-asn-matcher) block (within [`api_rate_limit.server_url_rules.client_matcher`](#api-rate-limit-server-url-rules-client-matcher)) supports the following:

&#x2022; [`asn_sets`](#asn-sets) - Optional Block<br>BGP ASN Sets. A list of references to bgp_asn_set objects<br>See [Asn Sets](#api-rate-limit-server-url-rules-client-matcher-asn-matcher-asn-sets) below.

#### API Rate Limit Server URL Rules Client Matcher Asn Matcher Asn Sets

An [`asn_sets`](#api-rate-limit-server-url-rules-client-matcher-asn-matcher-asn-sets) block (within [`api_rate_limit.server_url_rules.client_matcher.asn_matcher`](#api-rate-limit-server-url-rules-client-matcher-asn-matcher)) supports the following:

&#x2022; [`kind`](#kind) - Optional String<br>Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route')

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

&#x2022; [`uid`](#uid) - Optional String<br>UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid

#### API Rate Limit Server URL Rules Client Matcher Client Selector

A [`client_selector`](#api-rate-limit-server-url-rules-client-matcher-client-selector) block (within [`api_rate_limit.server_url_rules.client_matcher`](#api-rate-limit-server-url-rules-client-matcher)) supports the following:

&#x2022; [`expressions`](#expressions) - Optional List<br>Selector Expression. expressions contains the kubernetes style label expression for selections

#### API Rate Limit Server URL Rules Client Matcher IP Matcher

An [`ip_matcher`](#api-rate-limit-server-url-rules-client-matcher-ip-matcher) block (within [`api_rate_limit.server_url_rules.client_matcher`](#api-rate-limit-server-url-rules-client-matcher)) supports the following:

&#x2022; [`invert_matcher`](#invert-matcher) - Optional Bool<br>Invert IP Matcher. Invert the match result

&#x2022; [`prefix_sets`](#prefix-sets) - Optional Block<br>IP Prefix Sets. A list of references to ip_prefix_set objects<br>See [Prefix Sets](#api-rate-limit-server-url-rules-client-matcher-ip-matcher-prefix-sets) below.

#### API Rate Limit Server URL Rules Client Matcher IP Matcher Prefix Sets

A [`prefix_sets`](#api-rate-limit-server-url-rules-client-matcher-ip-matcher-prefix-sets) block (within [`api_rate_limit.server_url_rules.client_matcher.ip_matcher`](#api-rate-limit-server-url-rules-client-matcher-ip-matcher)) supports the following:

&#x2022; [`kind`](#kind) - Optional String<br>Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route')

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

&#x2022; [`uid`](#uid) - Optional String<br>UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid

#### API Rate Limit Server URL Rules Client Matcher IP Prefix List

An [`ip_prefix_list`](#api-rate-limit-server-url-rules-client-matcher-ip-prefix-list) block (within [`api_rate_limit.server_url_rules.client_matcher`](#api-rate-limit-server-url-rules-client-matcher)) supports the following:

&#x2022; [`invert_match`](#invert-match) - Optional Bool<br>Invert Match Result. Invert the match result

&#x2022; [`ip_prefixes`](#ip-prefixes) - Optional List<br>IPv4 Prefix List. List of IPv4 prefix strings

#### API Rate Limit Server URL Rules Client Matcher IP Threat Category List

An [`ip_threat_category_list`](#api-rate-limit-server-url-rules-client-matcher-ip-threat-category-list) block (within [`api_rate_limit.server_url_rules.client_matcher`](#api-rate-limit-server-url-rules-client-matcher)) supports the following:

&#x2022; [`ip_threat_categories`](#ip-threat-categories) - Optional List  Defaults to `SPAM_SOURCES`<br>Possible values are `SPAM_SOURCES`, `WINDOWS_EXPLOITS`, `WEB_ATTACKS`, `BOTNETS`, `SCANNERS`, `REPUTATION`, `PHISHING`, `PROXY`, `MOBILE_THREATS`, `TOR_PROXY`, `DENIAL_OF_SERVICE`, `NETWORK`<br>List of IP Threat Categories to choose. The IP threat categories is obtained from the list and is used to auto-generate equivalent label selection expressions

#### API Rate Limit Server URL Rules Client Matcher TLS Fingerprint Matcher

A [`tls_fingerprint_matcher`](#api-rate-limit-server-url-rules-client-matcher-tls-fingerprint-matcher) block (within [`api_rate_limit.server_url_rules.client_matcher`](#api-rate-limit-server-url-rules-client-matcher)) supports the following:

&#x2022; [`classes`](#classes) - Optional List  Defaults to `TLS_FINGERPRINT_NONE`<br>Possible values are `TLS_FINGERPRINT_NONE`, `ANY_MALICIOUS_FINGERPRINT`, `ADWARE`, `ADWIND`, `DRIDEX`, `GOOTKIT`, `GOZI`, `JBIFROST`, `QUAKBOT`, `RANSOMWARE`, `TROLDESH`, `TOFSEE`, `TORRENTLOCKER`, `TRICKBOT`<br>TLS fingerprint classes. A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against

&#x2022; [`exact_values`](#exact-values) - Optional List<br>Exact Values. A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against

&#x2022; [`excluded_values`](#excluded-values) - Optional List<br>Excluded Values. A list of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can be used to skip known false positives when using one or more known TLS fingerprint classes in the enclosing matcher

#### API Rate Limit Server URL Rules Inline Rate Limiter

An [`inline_rate_limiter`](#api-rate-limit-server-url-rules-inline-rate-limiter) block (within [`api_rate_limit.server_url_rules`](#api-rate-limit-server-url-rules)) supports the following:

&#x2022; [`ref_user_id`](#ref-user-id) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Ref User Id](#api-rate-limit-server-url-rules-inline-rate-limiter-ref-user-id) below.

&#x2022; [`threshold`](#threshold) - Optional Number<br>Threshold. The total number of allowed requests for 1 unit (e.g. SECOND/MINUTE/HOUR etc.) of the specified period

&#x2022; [`unit`](#unit) - Optional String  Defaults to `SECOND`<br>Possible values are `SECOND`, `MINUTE`, `HOUR`<br>Rate Limit Period Unit. Unit for the period per which the rate limit is applied. - SECOND: Second Rate limit period unit is seconds - MINUTE: Minute Rate limit period unit is minutes - HOUR: Hour Rate limit period unit is hours - DAY: Day Rate limit period unit is days

&#x2022; [`use_http_lb_user_id`](#use-http-lb-user-id) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### API Rate Limit Server URL Rules Inline Rate Limiter Ref User Id

A [`ref_user_id`](#api-rate-limit-server-url-rules-inline-rate-limiter-ref-user-id) block (within [`api_rate_limit.server_url_rules.inline_rate_limiter`](#api-rate-limit-server-url-rules-inline-rate-limiter)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### API Rate Limit Server URL Rules Ref Rate Limiter

A [`ref_rate_limiter`](#api-rate-limit-server-url-rules-ref-rate-limiter) block (within [`api_rate_limit.server_url_rules`](#api-rate-limit-server-url-rules)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### API Rate Limit Server URL Rules Request Matcher

A [`request_matcher`](#api-rate-limit-server-url-rules-request-matcher) block (within [`api_rate_limit.server_url_rules`](#api-rate-limit-server-url-rules)) supports the following:

&#x2022; [`cookie_matchers`](#cookie-matchers) - Optional Block<br>Cookie Matchers. A list of predicates for all cookies that need to be matched. The criteria for matching each cookie is described in individual instances of CookieMatcherType. The actual cookie values are extracted from the request API as a list of strings for each cookie name. Note that all specified cookie matcher predicates must evaluate to true<br>See [Cookie Matchers](#api-rate-limit-server-url-rules-request-matcher-cookie-matchers) below.

&#x2022; [`headers`](#headers) - Optional Block<br>HTTP Headers. A list of predicates for various HTTP headers that need to match. The criteria for matching each HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values are extracted from the request API as a list of strings for each HTTP header type. Note that all specified header predicates must evaluate to true<br>See [Headers](#api-rate-limit-server-url-rules-request-matcher-headers) below.

&#x2022; [`jwt_claims`](#jwt-claims) - Optional Block<br>JWT Claims. A list of predicates for various JWT claims that need to match. The criteria for matching each JWT claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are extracted from the JWT payload as a list of strings. Note that all specified JWT claim predicates must evaluate to true. Note that this feature only works on LBs with JWT Validation feature enabled<br>See [JWT Claims](#api-rate-limit-server-url-rules-request-matcher-jwt-claims) below.

&#x2022; [`query_params`](#query-params) - Optional Block<br>HTTP Query Parameters. A list of predicates for all query parameters that need to be matched. The criteria for matching each query parameter are described in individual instances of QueryParameterMatcherType. The actual query parameter values are extracted from the request API as a list of strings for each query parameter name. Note that all specified query parameter predicates must evaluate to true<br>See [Query Params](#api-rate-limit-server-url-rules-request-matcher-query-params) below.

#### API Rate Limit Server URL Rules Request Matcher Cookie Matchers

A [`cookie_matchers`](#api-rate-limit-server-url-rules-request-matcher-cookie-matchers) block (within [`api_rate_limit.server_url_rules.request_matcher`](#api-rate-limit-server-url-rules-request-matcher)) supports the following:

&#x2022; [`check_not_present`](#check-not-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`check_present`](#check-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`invert_matcher`](#invert-matcher) - Optional Bool<br>Invert Matcher. Invert Match of the expression defined

&#x2022; [`item`](#item) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#api-rate-limit-server-url-rules-request-matcher-cookie-matchers-item) below.

&#x2022; [`name`](#name) - Optional String<br>Cookie Name. A case-sensitive cookie name

#### API Rate Limit Server URL Rules Request Matcher Cookie Matchers Item

An [`item`](#api-rate-limit-server-url-rules-request-matcher-cookie-matchers-item) block (within [`api_rate_limit.server_url_rules.request_matcher.cookie_matchers`](#api-rate-limit-server-url-rules-request-matcher-cookie-matchers)) supports the following:

&#x2022; [`exact_values`](#exact-values) - Optional List<br>Exact Values. A list of exact values to match the input against

&#x2022; [`regex_values`](#regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

&#x2022; [`transformers`](#transformers) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### API Rate Limit Server URL Rules Request Matcher Headers

A [`headers`](#api-rate-limit-server-url-rules-request-matcher-headers) block (within [`api_rate_limit.server_url_rules.request_matcher`](#api-rate-limit-server-url-rules-request-matcher)) supports the following:

&#x2022; [`check_not_present`](#check-not-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`check_present`](#check-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`invert_matcher`](#invert-matcher) - Optional Bool<br>Invert Header Matcher. Invert the match result

&#x2022; [`item`](#item) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#api-rate-limit-server-url-rules-request-matcher-headers-item) below.

&#x2022; [`name`](#name) - Optional String<br>Header Name. A case-insensitive HTTP header name

#### API Rate Limit Server URL Rules Request Matcher Headers Item

An [`item`](#api-rate-limit-server-url-rules-request-matcher-headers-item) block (within [`api_rate_limit.server_url_rules.request_matcher.headers`](#api-rate-limit-server-url-rules-request-matcher-headers)) supports the following:

&#x2022; [`exact_values`](#exact-values) - Optional List<br>Exact Values. A list of exact values to match the input against

&#x2022; [`regex_values`](#regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

&#x2022; [`transformers`](#transformers) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### API Rate Limit Server URL Rules Request Matcher JWT Claims

A [`jwt_claims`](#api-rate-limit-server-url-rules-request-matcher-jwt-claims) block (within [`api_rate_limit.server_url_rules.request_matcher`](#api-rate-limit-server-url-rules-request-matcher)) supports the following:

&#x2022; [`check_not_present`](#check-not-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`check_present`](#check-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`invert_matcher`](#invert-matcher) - Optional Bool<br>Invert Matcher. Invert the match result

&#x2022; [`item`](#item) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#api-rate-limit-server-url-rules-request-matcher-jwt-claims-item) below.

&#x2022; [`name`](#name) - Optional String<br>JWT Claim Name. JWT claim name

#### API Rate Limit Server URL Rules Request Matcher JWT Claims Item

An [`item`](#api-rate-limit-server-url-rules-request-matcher-jwt-claims-item) block (within [`api_rate_limit.server_url_rules.request_matcher.jwt_claims`](#api-rate-limit-server-url-rules-request-matcher-jwt-claims)) supports the following:

&#x2022; [`exact_values`](#exact-values) - Optional List<br>Exact Values. A list of exact values to match the input against

&#x2022; [`regex_values`](#regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

&#x2022; [`transformers`](#transformers) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### API Rate Limit Server URL Rules Request Matcher Query Params

A [`query_params`](#api-rate-limit-server-url-rules-request-matcher-query-params) block (within [`api_rate_limit.server_url_rules.request_matcher`](#api-rate-limit-server-url-rules-request-matcher)) supports the following:

&#x2022; [`check_not_present`](#check-not-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`check_present`](#check-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`invert_matcher`](#invert-matcher) - Optional Bool<br>Invert Query Parameter Matcher. Invert the match result

&#x2022; [`item`](#item) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#api-rate-limit-server-url-rules-request-matcher-query-params-item) below.

&#x2022; [`key`](#key) - Optional String<br>Query Parameter Name. A case-sensitive HTTP query parameter name

#### API Rate Limit Server URL Rules Request Matcher Query Params Item

An [`item`](#api-rate-limit-server-url-rules-request-matcher-query-params-item) block (within [`api_rate_limit.server_url_rules.request_matcher.query_params`](#api-rate-limit-server-url-rules-request-matcher-query-params)) supports the following:

&#x2022; [`exact_values`](#exact-values) - Optional List<br>Exact Values. A list of exact values to match the input against

&#x2022; [`regex_values`](#regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

&#x2022; [`transformers`](#transformers) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### API Specification

An [`api_specification`](#api-specification) block supports the following:

&#x2022; [`api_definition`](#api-definition) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [API Definition](#api-specification-api-definition) below.

&#x2022; [`validation_all_spec_endpoints`](#validation-all-spec-endpoints) - Optional Block<br>API Inventory. Settings for API Inventory validation<br>See [Validation All Spec Endpoints](#api-specification-validation-all-spec-endpoints) below.

&#x2022; [`validation_custom_list`](#validation-custom-list) - Optional Block<br>Custom List. Define API groups, base paths, or API endpoints and their OpenAPI validation modes. Any other API-endpoint not listed will act according to 'Fall Through Mode'<br>See [Validation Custom List](#api-specification-validation-custom-list) below.

&#x2022; [`validation_disabled`](#validation-disabled) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### API Specification API Definition

An [`api_definition`](#api-specification-api-definition) block (within [`api_specification`](#api-specification)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### API Specification Validation All Spec Endpoints

A [`validation_all_spec_endpoints`](#api-specification-validation-all-spec-endpoints) block (within [`api_specification`](#api-specification)) supports the following:

&#x2022; [`fall_through_mode`](#fall-through-mode) - Optional Block<br>Fall Through Mode. x-required Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a. swagger) or doesn't have a specific rule in custom rules)<br>See [Fall Through Mode](#api-specification-validation-all-spec-endpoints-fall-through-mode) below.

&#x2022; [`settings`](#settings) - Optional Block<br>Common Settings. OpenAPI specification validation settings relevant for 'API Inventory' enforcement and for 'Custom list' enforcement<br>See [Settings](#api-specification-validation-all-spec-endpoints-settings) below.

&#x2022; [`validation_mode`](#validation-mode) - Optional Block<br>Validation Mode. x-required Validation mode of OpenAPI specification. When a validation mismatch occurs on a request to one of the endpoints listed on the OpenAPI specification file (a.k.a. swagger)<br>See [Validation Mode](#api-specification-validation-all-spec-endpoints-validation-mode) below.

#### API Specification Validation All Spec Endpoints Fall Through Mode

A [`fall_through_mode`](#api-specification-validation-all-spec-endpoints-fall-through-mode) block (within [`api_specification.validation_all_spec_endpoints`](#api-specification-validation-all-spec-endpoints)) supports the following:

&#x2022; [`fall_through_mode_allow`](#fall-through-mode-allow) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`fall_through_mode_custom`](#fall-through-mode-custom) - Optional Block<br>Custom Fall Through Mode. Define the fall through settings<br>See [Fall Through Mode Custom](#api-specification-validation-all-spec-endpoints-fall-through-mode-fall-through-mode-custom) below.

#### API Specification Validation All Spec Endpoints Fall Through Mode Fall Through Mode Custom

A [`fall_through_mode_custom`](#api-specification-validation-all-spec-endpoints-fall-through-mode-fall-through-mode-custom) block (within [`api_specification.validation_all_spec_endpoints.fall_through_mode`](#api-specification-validation-all-spec-endpoints-fall-through-mode)) supports the following:

&#x2022; [`open_api_validation_rules`](#open-api-validation-rules) - Optional Block<br>Custom Fall Through Rule List<br>See [Open API Validation Rules](#api-specification-validation-all-spec-endpoints-fall-through-mode-fall-through-mode-custom-open-api-validation-rules) below.

#### API Specification Validation All Spec Endpoints Fall Through Mode Fall Through Mode Custom Open API Validation Rules

An [`open_api_validation_rules`](#api-specification-validation-all-spec-endpoints-fall-through-mode-fall-through-mode-custom-open-api-validation-rules) block (within [`api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom`](#api-specification-validation-all-spec-endpoints-fall-through-mode-fall-through-mode-custom)) supports the following:

&#x2022; [`action_block`](#action-block) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`action_report`](#action-report) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`action_skip`](#action-skip) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`api_endpoint`](#api-endpoint) - Optional Block<br>API Endpoint. This defines API endpoint<br>See [API Endpoint](#api-specification-validation-all-spec-endpoints-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-api-endpoint) below.

&#x2022; [`api_group`](#api-group) - Optional String<br>API Group. The API group which this validation applies to

&#x2022; [`base_path`](#base-path) - Optional String<br>Base Path. The base path which this validation applies to

&#x2022; [`metadata`](#metadata) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#api-specification-validation-all-spec-endpoints-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-metadata) below.

#### API Specification Validation All Spec Endpoints Fall Through Mode Fall Through Mode Custom Open API Validation Rules API Endpoint

An [`api_endpoint`](#api-specification-validation-all-spec-endpoints-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-api-endpoint) block (within [`api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules`](#api-specification-validation-all-spec-endpoints-fall-through-mode-fall-through-mode-custom-open-api-validation-rules)) supports the following:

&#x2022; [`methods`](#methods) - Optional List  Defaults to `ANY`<br>Possible values are `ANY`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`, `COPY`<br>Methods. Methods to be matched

&#x2022; [`path`](#path) - Optional String<br>Path. Path to be matched

#### API Specification Validation All Spec Endpoints Fall Through Mode Fall Through Mode Custom Open API Validation Rules Metadata

A [`metadata`](#api-specification-validation-all-spec-endpoints-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-metadata) block (within [`api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules`](#api-specification-validation-all-spec-endpoints-fall-through-mode-fall-through-mode-custom-open-api-validation-rules)) supports the following:

&#x2022; [`description`](#description) - Optional String<br>Description. Human readable description

&#x2022; [`name`](#name) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### API Specification Validation All Spec Endpoints Settings

A [`settings`](#api-specification-validation-all-spec-endpoints-settings) block (within [`api_specification.validation_all_spec_endpoints`](#api-specification-validation-all-spec-endpoints)) supports the following:

&#x2022; [`oversized_body_fail_validation`](#oversized-body-fail-validation) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`oversized_body_skip_validation`](#oversized-body-skip-validation) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`property_validation_settings_custom`](#property-validation-settings-custom) - Optional Block<br>Validation Property Settings. Custom property validation settings<br>See [Property Validation Settings Custom](#api-specification-validation-all-spec-endpoints-settings-property-validation-settings-custom) below.

&#x2022; [`property_validation_settings_default`](#property-validation-settings-default) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### API Specification Validation All Spec Endpoints Settings Property Validation Settings Custom

A [`property_validation_settings_custom`](#api-specification-validation-all-spec-endpoints-settings-property-validation-settings-custom) block (within [`api_specification.validation_all_spec_endpoints.settings`](#api-specification-validation-all-spec-endpoints-settings)) supports the following:

&#x2022; [`query_parameters`](#query-parameters) - Optional Block<br>Validation Settings For Query Parameters. Custom settings for query parameters validation<br>See [Query Parameters](#api-specification-validation-all-spec-endpoints-settings-property-validation-settings-custom-query-parameters) below.

#### API Specification Validation All Spec Endpoints Settings Property Validation Settings Custom Query Parameters

A [`query_parameters`](#api-specification-validation-all-spec-endpoints-settings-property-validation-settings-custom-query-parameters) block (within [`api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom`](#api-specification-validation-all-spec-endpoints-settings-property-validation-settings-custom)) supports the following:

&#x2022; [`allow_additional_parameters`](#allow-additional-parameters) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`disallow_additional_parameters`](#disallow-additional-parameters) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### API Specification Validation All Spec Endpoints Validation Mode

A [`validation_mode`](#api-specification-validation-all-spec-endpoints-validation-mode) block (within [`api_specification.validation_all_spec_endpoints`](#api-specification-validation-all-spec-endpoints)) supports the following:

&#x2022; [`response_validation_mode_active`](#response-validation-mode-active) - Optional Block<br>Open API Validation Mode Active. Validation mode properties of response<br>See [Response Validation Mode Active](#api-specification-validation-all-spec-endpoints-validation-mode-response-validation-mode-active) below.

&#x2022; [`skip_response_validation`](#skip-response-validation) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`skip_validation`](#skip-validation) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`validation_mode_active`](#validation-mode-active) - Optional Block<br>Open API Validation Mode Active. Validation mode properties of request<br>See [Validation Mode Active](#api-specification-validation-all-spec-endpoints-validation-mode-validation-mode-active) below.

#### API Specification Validation All Spec Endpoints Validation Mode Response Validation Mode Active

A [`response_validation_mode_active`](#api-specification-validation-all-spec-endpoints-validation-mode-response-validation-mode-active) block (within [`api_specification.validation_all_spec_endpoints.validation_mode`](#api-specification-validation-all-spec-endpoints-validation-mode)) supports the following:

&#x2022; [`enforcement_block`](#enforcement-block) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`enforcement_report`](#enforcement-report) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`response_validation_properties`](#response-validation-properties) - Optional List  Defaults to `PROPERTY_QUERY_PARAMETERS`<br>Possible values are `PROPERTY_QUERY_PARAMETERS`, `PROPERTY_PATH_PARAMETERS`, `PROPERTY_CONTENT_TYPE`, `PROPERTY_COOKIE_PARAMETERS`, `PROPERTY_HTTP_HEADERS`, `PROPERTY_HTTP_BODY`, `PROPERTY_SECURITY_SCHEMA`, `PROPERTY_RESPONSE_CODE`<br>Response Validation Properties. List of properties of the response to validate according to the OpenAPI specification file (a.k.a. swagger)

#### API Specification Validation All Spec Endpoints Validation Mode Validation Mode Active

A [`validation_mode_active`](#api-specification-validation-all-spec-endpoints-validation-mode-validation-mode-active) block (within [`api_specification.validation_all_spec_endpoints.validation_mode`](#api-specification-validation-all-spec-endpoints-validation-mode)) supports the following:

&#x2022; [`enforcement_block`](#enforcement-block) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`enforcement_report`](#enforcement-report) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`request_validation_properties`](#request-validation-properties) - Optional List  Defaults to `PROPERTY_QUERY_PARAMETERS`<br>Possible values are `PROPERTY_QUERY_PARAMETERS`, `PROPERTY_PATH_PARAMETERS`, `PROPERTY_CONTENT_TYPE`, `PROPERTY_COOKIE_PARAMETERS`, `PROPERTY_HTTP_HEADERS`, `PROPERTY_HTTP_BODY`, `PROPERTY_SECURITY_SCHEMA`, `PROPERTY_RESPONSE_CODE`<br>Request Validation Properties. List of properties of the request to validate according to the OpenAPI specification file (a.k.a. swagger)

#### API Specification Validation Custom List

A [`validation_custom_list`](#api-specification-validation-custom-list) block (within [`api_specification`](#api-specification)) supports the following:

&#x2022; [`fall_through_mode`](#fall-through-mode) - Optional Block<br>Fall Through Mode. x-required Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a. swagger) or doesn't have a specific rule in custom rules)<br>See [Fall Through Mode](#api-specification-validation-custom-list-fall-through-mode) below.

&#x2022; [`open_api_validation_rules`](#open-api-validation-rules) - Optional Block<br>Validation List<br>See [Open API Validation Rules](#api-specification-validation-custom-list-open-api-validation-rules) below.

&#x2022; [`settings`](#settings) - Optional Block<br>Common Settings. OpenAPI specification validation settings relevant for 'API Inventory' enforcement and for 'Custom list' enforcement<br>See [Settings](#api-specification-validation-custom-list-settings) below.

#### API Specification Validation Custom List Fall Through Mode

A [`fall_through_mode`](#api-specification-validation-custom-list-fall-through-mode) block (within [`api_specification.validation_custom_list`](#api-specification-validation-custom-list)) supports the following:

&#x2022; [`fall_through_mode_allow`](#fall-through-mode-allow) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`fall_through_mode_custom`](#fall-through-mode-custom) - Optional Block<br>Custom Fall Through Mode. Define the fall through settings<br>See [Fall Through Mode Custom](#api-specification-validation-custom-list-fall-through-mode-fall-through-mode-custom) below.

#### API Specification Validation Custom List Fall Through Mode Fall Through Mode Custom

A [`fall_through_mode_custom`](#api-specification-validation-custom-list-fall-through-mode-fall-through-mode-custom) block (within [`api_specification.validation_custom_list.fall_through_mode`](#api-specification-validation-custom-list-fall-through-mode)) supports the following:

&#x2022; [`open_api_validation_rules`](#open-api-validation-rules) - Optional Block<br>Custom Fall Through Rule List<br>See [Open API Validation Rules](#api-specification-validation-custom-list-fall-through-mode-fall-through-mode-custom-open-api-validation-rules) below.

#### API Specification Validation Custom List Fall Through Mode Fall Through Mode Custom Open API Validation Rules

An [`open_api_validation_rules`](#api-specification-validation-custom-list-fall-through-mode-fall-through-mode-custom-open-api-validation-rules) block (within [`api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom`](#api-specification-validation-custom-list-fall-through-mode-fall-through-mode-custom)) supports the following:

&#x2022; [`action_block`](#action-block) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`action_report`](#action-report) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`action_skip`](#action-skip) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`api_endpoint`](#api-endpoint) - Optional Block<br>API Endpoint. This defines API endpoint<br>See [API Endpoint](#api-specification-validation-custom-list-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-api-endpoint) below.

&#x2022; [`api_group`](#api-group) - Optional String<br>API Group. The API group which this validation applies to

&#x2022; [`base_path`](#base-path) - Optional String<br>Base Path. The base path which this validation applies to

&#x2022; [`metadata`](#metadata) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#api-specification-validation-custom-list-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-metadata) below.

#### API Specification Validation Custom List Fall Through Mode Fall Through Mode Custom Open API Validation Rules API Endpoint

An [`api_endpoint`](#api-specification-validation-custom-list-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-api-endpoint) block (within [`api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules`](#api-specification-validation-custom-list-fall-through-mode-fall-through-mode-custom-open-api-validation-rules)) supports the following:

&#x2022; [`methods`](#methods) - Optional List  Defaults to `ANY`<br>Possible values are `ANY`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`, `COPY`<br>Methods. Methods to be matched

&#x2022; [`path`](#path) - Optional String<br>Path. Path to be matched

#### API Specification Validation Custom List Fall Through Mode Fall Through Mode Custom Open API Validation Rules Metadata

A [`metadata`](#api-specification-validation-custom-list-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-metadata) block (within [`api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules`](#api-specification-validation-custom-list-fall-through-mode-fall-through-mode-custom-open-api-validation-rules)) supports the following:

&#x2022; [`description`](#description) - Optional String<br>Description. Human readable description

&#x2022; [`name`](#name) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### API Specification Validation Custom List Open API Validation Rules

An [`open_api_validation_rules`](#api-specification-validation-custom-list-open-api-validation-rules) block (within [`api_specification.validation_custom_list`](#api-specification-validation-custom-list)) supports the following:

&#x2022; [`any_domain`](#any-domain) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`api_endpoint`](#api-endpoint) - Optional Block<br>API Endpoint. This defines API endpoint<br>See [API Endpoint](#api-specification-validation-custom-list-open-api-validation-rules-api-endpoint) below.

&#x2022; [`api_group`](#api-group) - Optional String<br>API Group. The API group which this validation applies to

&#x2022; [`base_path`](#base-path) - Optional String<br>Base Path. The base path which this validation applies to

&#x2022; [`metadata`](#metadata) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#api-specification-validation-custom-list-open-api-validation-rules-metadata) below.

&#x2022; [`specific_domain`](#specific-domain) - Optional String<br>Specific Domain. The rule will apply for a specific domain

&#x2022; [`validation_mode`](#validation-mode) - Optional Block<br>Validation Mode. x-required Validation mode of OpenAPI specification. When a validation mismatch occurs on a request to one of the endpoints listed on the OpenAPI specification file (a.k.a. swagger)<br>See [Validation Mode](#api-specification-validation-custom-list-open-api-validation-rules-validation-mode) below.

#### API Specification Validation Custom List Open API Validation Rules API Endpoint

An [`api_endpoint`](#api-specification-validation-custom-list-open-api-validation-rules-api-endpoint) block (within [`api_specification.validation_custom_list.open_api_validation_rules`](#api-specification-validation-custom-list-open-api-validation-rules)) supports the following:

&#x2022; [`methods`](#methods) - Optional List  Defaults to `ANY`<br>Possible values are `ANY`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`, `COPY`<br>Methods. Methods to be matched

&#x2022; [`path`](#path) - Optional String<br>Path. Path to be matched

#### API Specification Validation Custom List Open API Validation Rules Metadata

A [`metadata`](#api-specification-validation-custom-list-open-api-validation-rules-metadata) block (within [`api_specification.validation_custom_list.open_api_validation_rules`](#api-specification-validation-custom-list-open-api-validation-rules)) supports the following:

&#x2022; [`description`](#description) - Optional String<br>Description. Human readable description

&#x2022; [`name`](#name) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### API Specification Validation Custom List Open API Validation Rules Validation Mode

A [`validation_mode`](#api-specification-validation-custom-list-open-api-validation-rules-validation-mode) block (within [`api_specification.validation_custom_list.open_api_validation_rules`](#api-specification-validation-custom-list-open-api-validation-rules)) supports the following:

&#x2022; [`response_validation_mode_active`](#response-validation-mode-active) - Optional Block<br>Open API Validation Mode Active. Validation mode properties of response<br>See [Response Validation Mode Active](#api-specification-validation-custom-list-open-api-validation-rules-validation-mode-response-validation-mode-active) below.

&#x2022; [`skip_response_validation`](#skip-response-validation) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`skip_validation`](#skip-validation) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`validation_mode_active`](#validation-mode-active) - Optional Block<br>Open API Validation Mode Active. Validation mode properties of request<br>See [Validation Mode Active](#api-specification-validation-custom-list-open-api-validation-rules-validation-mode-validation-mode-active) below.

#### API Specification Validation Custom List Open API Validation Rules Validation Mode Response Validation Mode Active

A [`response_validation_mode_active`](#api-specification-validation-custom-list-open-api-validation-rules-validation-mode-response-validation-mode-active) block (within [`api_specification.validation_custom_list.open_api_validation_rules.validation_mode`](#api-specification-validation-custom-list-open-api-validation-rules-validation-mode)) supports the following:

&#x2022; [`enforcement_block`](#enforcement-block) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`enforcement_report`](#enforcement-report) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`response_validation_properties`](#response-validation-properties) - Optional List  Defaults to `PROPERTY_QUERY_PARAMETERS`<br>Possible values are `PROPERTY_QUERY_PARAMETERS`, `PROPERTY_PATH_PARAMETERS`, `PROPERTY_CONTENT_TYPE`, `PROPERTY_COOKIE_PARAMETERS`, `PROPERTY_HTTP_HEADERS`, `PROPERTY_HTTP_BODY`, `PROPERTY_SECURITY_SCHEMA`, `PROPERTY_RESPONSE_CODE`<br>Response Validation Properties. List of properties of the response to validate according to the OpenAPI specification file (a.k.a. swagger)

#### API Specification Validation Custom List Open API Validation Rules Validation Mode Validation Mode Active

A [`validation_mode_active`](#api-specification-validation-custom-list-open-api-validation-rules-validation-mode-validation-mode-active) block (within [`api_specification.validation_custom_list.open_api_validation_rules.validation_mode`](#api-specification-validation-custom-list-open-api-validation-rules-validation-mode)) supports the following:

&#x2022; [`enforcement_block`](#enforcement-block) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`enforcement_report`](#enforcement-report) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`request_validation_properties`](#request-validation-properties) - Optional List  Defaults to `PROPERTY_QUERY_PARAMETERS`<br>Possible values are `PROPERTY_QUERY_PARAMETERS`, `PROPERTY_PATH_PARAMETERS`, `PROPERTY_CONTENT_TYPE`, `PROPERTY_COOKIE_PARAMETERS`, `PROPERTY_HTTP_HEADERS`, `PROPERTY_HTTP_BODY`, `PROPERTY_SECURITY_SCHEMA`, `PROPERTY_RESPONSE_CODE`<br>Request Validation Properties. List of properties of the request to validate according to the OpenAPI specification file (a.k.a. swagger)

#### API Specification Validation Custom List Settings

A [`settings`](#api-specification-validation-custom-list-settings) block (within [`api_specification.validation_custom_list`](#api-specification-validation-custom-list)) supports the following:

&#x2022; [`oversized_body_fail_validation`](#oversized-body-fail-validation) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`oversized_body_skip_validation`](#oversized-body-skip-validation) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`property_validation_settings_custom`](#property-validation-settings-custom) - Optional Block<br>Validation Property Settings. Custom property validation settings<br>See [Property Validation Settings Custom](#api-specification-validation-custom-list-settings-property-validation-settings-custom) below.

&#x2022; [`property_validation_settings_default`](#property-validation-settings-default) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### API Specification Validation Custom List Settings Property Validation Settings Custom

A [`property_validation_settings_custom`](#api-specification-validation-custom-list-settings-property-validation-settings-custom) block (within [`api_specification.validation_custom_list.settings`](#api-specification-validation-custom-list-settings)) supports the following:

&#x2022; [`query_parameters`](#query-parameters) - Optional Block<br>Validation Settings For Query Parameters. Custom settings for query parameters validation<br>See [Query Parameters](#api-specification-validation-custom-list-settings-property-validation-settings-custom-query-parameters) below.

#### API Specification Validation Custom List Settings Property Validation Settings Custom Query Parameters

A [`query_parameters`](#api-specification-validation-custom-list-settings-property-validation-settings-custom-query-parameters) block (within [`api_specification.validation_custom_list.settings.property_validation_settings_custom`](#api-specification-validation-custom-list-settings-property-validation-settings-custom)) supports the following:

&#x2022; [`allow_additional_parameters`](#allow-additional-parameters) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`disallow_additional_parameters`](#disallow-additional-parameters) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### API Testing

An [`api_testing`](#api-testing) block supports the following:

&#x2022; [`custom_header_value`](#custom-header-value) - Optional String<br>Custom Header. Add x-f5-API-testing-identifier header value to prevent security flags on API testing traffic

&#x2022; [`domains`](#domains) - Optional Block<br>Testing Environments. Add and configure testing domains and credentials<br>See [Domains](#api-testing-domains) below.

&#x2022; [`every_day`](#every-day) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`every_month`](#every-month) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`every_week`](#every-week) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### API Testing Domains

A [`domains`](#api-testing-domains) block (within [`api_testing`](#api-testing)) supports the following:

&#x2022; [`allow_destructive_methods`](#allow-destructive-methods) - Optional Bool<br>Use Destructive Methods (e.g., DELETE, PUT). Enable to allow API test to execute destructive methods. Be cautious as these can alter or delete data

&#x2022; [`credentials`](#credentials) - Optional Block<br>Credentials. Add credentials for API testing to use in the selected environment<br>See [Credentials](#api-testing-domains-credentials) below.

&#x2022; [`domain`](#domain) - Optional String<br>Domain. Add your testing environment domain. Be aware that running tests on a production domain can impact live applications, as API testing cannot distinguish between production and testing environments

#### API Testing Domains Credentials

A [`credentials`](#api-testing-domains-credentials) block (within [`api_testing.domains`](#api-testing-domains)) supports the following:

&#x2022; [`admin`](#admin) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`api_key`](#api-key) - Optional Block<br>API Key<br>See [API Key](#api-testing-domains-credentials-api-key) below.

&#x2022; [`basic_auth`](#basic-auth) - Optional Block<br>Basic Authentication<br>See [Basic Auth](#api-testing-domains-credentials-basic-auth) below.

&#x2022; [`bearer_token`](#bearer-token) - Optional Block<br>Bearer<br>See [Bearer Token](#api-testing-domains-credentials-bearer-token) below.

&#x2022; [`credential_name`](#credential-name) - Optional String<br>Name. Enter a unique name for the credentials used in API testing

&#x2022; [`login_endpoint`](#login-endpoint) - Optional Block<br>Login Endpoint<br>See [Login Endpoint](#api-testing-domains-credentials-login-endpoint) below.

&#x2022; [`standard`](#standard) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### API Testing Domains Credentials API Key

An [`api_key`](#api-testing-domains-credentials-api-key) block (within [`api_testing.domains.credentials`](#api-testing-domains-credentials)) supports the following:

&#x2022; [`key`](#key) - Optional String<br>Key

&#x2022; [`value`](#value) - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Value](#api-testing-domains-credentials-api-key-value) below.

#### API Testing Domains Credentials API Key Value

A [`value`](#api-testing-domains-credentials-api-key-value) block (within [`api_testing.domains.credentials.api_key`](#api-testing-domains-credentials-api-key)) supports the following:

&#x2022; [`blindfold_secret_info`](#blindfold-secret-info) - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#api-testing-domains-credentials-api-key-value-blindfold-secret-info) below.

&#x2022; [`clear_secret_info`](#clear-secret-info) - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#api-testing-domains-credentials-api-key-value-clear-secret-info) below.

#### API Testing Domains Credentials API Key Value Blindfold Secret Info

A [`blindfold_secret_info`](#api-testing-domains-credentials-api-key-value-blindfold-secret-info) block (within [`api_testing.domains.credentials.api_key.value`](#api-testing-domains-credentials-api-key-value)) supports the following:

&#x2022; [`decryption_provider`](#decryption-provider) - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

&#x2022; [`location`](#location) - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

&#x2022; [`store_provider`](#store-provider) - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

#### API Testing Domains Credentials API Key Value Clear Secret Info

A [`clear_secret_info`](#api-testing-domains-credentials-api-key-value-clear-secret-info) block (within [`api_testing.domains.credentials.api_key.value`](#api-testing-domains-credentials-api-key-value)) supports the following:

&#x2022; [`provider_ref`](#provider-ref) - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

&#x2022; [`url`](#url) - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

#### API Testing Domains Credentials Basic Auth

A [`basic_auth`](#api-testing-domains-credentials-basic-auth) block (within [`api_testing.domains.credentials`](#api-testing-domains-credentials)) supports the following:

&#x2022; [`password`](#password) - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Password](#api-testing-domains-credentials-basic-auth-password) below.

&#x2022; [`user`](#user) - Optional String<br>User

#### API Testing Domains Credentials Basic Auth Password

A [`password`](#api-testing-domains-credentials-basic-auth-password) block (within [`api_testing.domains.credentials.basic_auth`](#api-testing-domains-credentials-basic-auth)) supports the following:

&#x2022; [`blindfold_secret_info`](#blindfold-secret-info) - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#api-testing-domains-credentials-basic-auth-password-blindfold-secret-info) below.

&#x2022; [`clear_secret_info`](#clear-secret-info) - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#api-testing-domains-credentials-basic-auth-password-clear-secret-info) below.

#### API Testing Domains Credentials Basic Auth Password Blindfold Secret Info

A [`blindfold_secret_info`](#api-testing-domains-credentials-basic-auth-password-blindfold-secret-info) block (within [`api_testing.domains.credentials.basic_auth.password`](#api-testing-domains-credentials-basic-auth-password)) supports the following:

&#x2022; [`decryption_provider`](#decryption-provider) - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

&#x2022; [`location`](#location) - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

&#x2022; [`store_provider`](#store-provider) - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

#### API Testing Domains Credentials Basic Auth Password Clear Secret Info

A [`clear_secret_info`](#api-testing-domains-credentials-basic-auth-password-clear-secret-info) block (within [`api_testing.domains.credentials.basic_auth.password`](#api-testing-domains-credentials-basic-auth-password)) supports the following:

&#x2022; [`provider_ref`](#provider-ref) - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

&#x2022; [`url`](#url) - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

#### API Testing Domains Credentials Bearer Token

A [`bearer_token`](#api-testing-domains-credentials-bearer-token) block (within [`api_testing.domains.credentials`](#api-testing-domains-credentials)) supports the following:

&#x2022; [`token`](#token) - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Token](#api-testing-domains-credentials-bearer-token-token) below.

#### API Testing Domains Credentials Bearer Token Token

A [`token`](#api-testing-domains-credentials-bearer-token-token) block (within [`api_testing.domains.credentials.bearer_token`](#api-testing-domains-credentials-bearer-token)) supports the following:

&#x2022; [`blindfold_secret_info`](#blindfold-secret-info) - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#api-testing-domains-credentials-bearer-token-token-blindfold-secret-info) below.

&#x2022; [`clear_secret_info`](#clear-secret-info) - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#api-testing-domains-credentials-bearer-token-token-clear-secret-info) below.

#### API Testing Domains Credentials Bearer Token Token Blindfold Secret Info

A [`blindfold_secret_info`](#api-testing-domains-credentials-bearer-token-token-blindfold-secret-info) block (within [`api_testing.domains.credentials.bearer_token.token`](#api-testing-domains-credentials-bearer-token-token)) supports the following:

&#x2022; [`decryption_provider`](#decryption-provider) - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

&#x2022; [`location`](#location) - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

&#x2022; [`store_provider`](#store-provider) - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

#### API Testing Domains Credentials Bearer Token Token Clear Secret Info

A [`clear_secret_info`](#api-testing-domains-credentials-bearer-token-token-clear-secret-info) block (within [`api_testing.domains.credentials.bearer_token.token`](#api-testing-domains-credentials-bearer-token-token)) supports the following:

&#x2022; [`provider_ref`](#provider-ref) - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

&#x2022; [`url`](#url) - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

#### API Testing Domains Credentials Login Endpoint

A [`login_endpoint`](#api-testing-domains-credentials-login-endpoint) block (within [`api_testing.domains.credentials`](#api-testing-domains-credentials)) supports the following:

&#x2022; [`json_payload`](#json-payload) - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [JSON Payload](#api-testing-domains-credentials-login-endpoint-json-payload) below.

&#x2022; [`method`](#method) - Optional String  Defaults to `ANY`<br>Possible values are `ANY`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`, `COPY`<br>HTTP Method. Specifies the HTTP method used to access a resource. Any HTTP Method

&#x2022; [`path`](#path) - Optional String<br>Path

&#x2022; [`token_response_key`](#token-response-key) - Optional String<br>Token Response Key. Specifies how to handle the API response, extracting authentication tokens

#### API Testing Domains Credentials Login Endpoint JSON Payload

A [`json_payload`](#api-testing-domains-credentials-login-endpoint-json-payload) block (within [`api_testing.domains.credentials.login_endpoint`](#api-testing-domains-credentials-login-endpoint)) supports the following:

&#x2022; [`blindfold_secret_info`](#blindfold-secret-info) - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#api-testing-domains-credentials-login-endpoint-json-payload-blindfold-secret-info) below.

&#x2022; [`clear_secret_info`](#clear-secret-info) - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#api-testing-domains-credentials-login-endpoint-json-payload-clear-secret-info) below.

#### API Testing Domains Credentials Login Endpoint JSON Payload Blindfold Secret Info

A [`blindfold_secret_info`](#api-testing-domains-credentials-login-endpoint-json-payload-blindfold-secret-info) block (within [`api_testing.domains.credentials.login_endpoint.json_payload`](#api-testing-domains-credentials-login-endpoint-json-payload)) supports the following:

&#x2022; [`decryption_provider`](#decryption-provider) - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

&#x2022; [`location`](#location) - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

&#x2022; [`store_provider`](#store-provider) - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

#### API Testing Domains Credentials Login Endpoint JSON Payload Clear Secret Info

A [`clear_secret_info`](#api-testing-domains-credentials-login-endpoint-json-payload-clear-secret-info) block (within [`api_testing.domains.credentials.login_endpoint.json_payload`](#api-testing-domains-credentials-login-endpoint-json-payload)) supports the following:

&#x2022; [`provider_ref`](#provider-ref) - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

&#x2022; [`url`](#url) - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

#### App Firewall

An [`app_firewall`](#app-firewall) block supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Blocked Clients

A [`blocked_clients`](#blocked-clients) block supports the following:

&#x2022; [`actions`](#actions) - Optional List  Defaults to `SKIP_PROCESSING_WAF`<br>Possible values are `SKIP_PROCESSING_WAF`, `SKIP_PROCESSING_BOT`, `SKIP_PROCESSING_MUM`, `SKIP_PROCESSING_IP_REPUTATION`, `SKIP_PROCESSING_API_PROTECTION`, `SKIP_PROCESSING_OAS_VALIDATION`, `SKIP_PROCESSING_DDOS_PROTECTION`, `SKIP_PROCESSING_THREAT_MESH`, `SKIP_PROCESSING_MALWARE_PROTECTION`<br>Actions. Actions that should be taken when client identifier matches the rule

&#x2022; [`as_number`](#as-number) - Optional Number<br>AS Number. RFC 6793 defined 4-byte AS number

&#x2022; [`bot_skip_processing`](#bot-skip-processing) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`expiration_timestamp`](#expiration-timestamp) - Optional String<br>Expiration Timestamp. The expiration_timestamp is the RFC 3339 format timestamp at which the containing rule is considered to be logically expired. The rule continues to exist in the configuration but is not applied anymore

&#x2022; [`http_header`](#http-header) - Optional Block<br>HTTP Header. Request header name and value pairs<br>See [HTTP Header](#blocked-clients-http-header) below.

&#x2022; [`ip_prefix`](#ip-prefix) - Optional String<br>IPv4 Prefix. IPv4 prefix string

&#x2022; [`ipv6_prefix`](#ipv6-prefix) - Optional String<br>IPv6 Prefix. IPv6 prefix string

&#x2022; [`metadata`](#metadata) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#blocked-clients-metadata) below.

&#x2022; [`skip_processing`](#skip-processing) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`user_identifier`](#user-identifier) - Optional String<br>User Identifier. Identify user based on user identifier. User identifier value needs to be copied from security event

&#x2022; [`waf_skip_processing`](#waf-skip-processing) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### Blocked Clients HTTP Header

A [`http_header`](#blocked-clients-http-header) block (within [`blocked_clients`](#blocked-clients)) supports the following:

&#x2022; [`headers`](#headers) - Optional Block<br>HTTP Headers. List of HTTP header name and value pairs<br>See [Headers](#blocked-clients-http-header-headers) below.

#### Blocked Clients HTTP Header Headers

A [`headers`](#blocked-clients-http-header-headers) block (within [`blocked_clients.http_header`](#blocked-clients-http-header)) supports the following:

&#x2022; [`exact`](#exact) - Optional String<br>Exact. Header value to match exactly

&#x2022; [`invert_match`](#invert-match) - Optional Bool<br>NOT of match. Invert the result of the match to detect missing header or non-matching value

&#x2022; [`name`](#name) - Optional String<br>Name. Name of the header

&#x2022; [`presence`](#presence) - Optional Bool<br>Presence. If true, check for presence of header

&#x2022; [`regex`](#regex) - Optional String<br>Regex. Regex match of the header value in re2 format

#### Blocked Clients Metadata

A [`metadata`](#blocked-clients-metadata) block (within [`blocked_clients`](#blocked-clients)) supports the following:

&#x2022; [`description`](#description) - Optional String<br>Description. Human readable description

&#x2022; [`name`](#name) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### Bot Defense

A [`bot_defense`](#bot-defense) block supports the following:

&#x2022; [`disable_cors_support`](#disable-cors-support) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`enable_cors_support`](#enable-cors-support) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`policy`](#policy) - Optional Block<br>Bot Defense Policy. This defines various configuration options for Bot Defense policy<br>See [Policy](#bot-defense-policy) below.

&#x2022; [`regional_endpoint`](#regional-endpoint) - Optional String  Defaults to `AUTO`<br>Possible values are `AUTO`, `US`, `EU`, `ASIA`<br>Bot Defense Region. Defines a selection for Bot Defense region - AUTO: AUTO Automatic selection based on client IP address - US: US US region - EU: EU European Union region - ASIA: ASIA Asia region

&#x2022; [`timeout`](#timeout) - Optional Number<br>Timeout. The timeout for the inference check, in milliseconds

#### Bot Defense Policy

A [`policy`](#bot-defense-policy) block (within [`bot_defense`](#bot-defense)) supports the following:

&#x2022; [`disable_js_insert`](#disable-js-insert) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`disable_mobile_sdk`](#disable-mobile-sdk) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`javascript_mode`](#javascript-mode) - Optional String  Defaults to `ASYNC_JS_NO_CACHING`<br>Possible values are `ASYNC_JS_NO_CACHING`, `ASYNC_JS_CACHING`, `SYNC_JS_NO_CACHING`, `SYNC_JS_CACHING`<br>Web Client JavaScript Mode. Web Client JavaScript Mode. Bot Defense JavaScript for telemetry collection is requested asynchronously, and it is non-cacheable Bot Defense JavaScript for telemetry collection is requested asynchronously, and it is cacheable Bot Defense JavaScript for telemetry collection is requested synchronously, and it is non-cacheable Bot Defense JavaScript for telemetry collection is requested synchronously, and it is cacheable

&#x2022; [`js_download_path`](#js-download-path) - Optional String<br>JavaScript Download Path. Customize Bot Defense Client JavaScript path. If not specified, default `/common.js`

&#x2022; [`js_insert_all_pages`](#js-insert-all-pages) - Optional Block<br>Insert Bot Defense JavaScript in All Pages. Insert Bot Defense JavaScript in all pages<br>See [Js Insert All Pages](#bot-defense-policy-js-insert-all-pages) below.

&#x2022; [`js_insert_all_pages_except`](#js-insert-all-pages-except) - Optional Block<br>Insert JavaScript in All Pages with the Exceptions. Insert Bot Defense JavaScript in all pages with the exceptions<br>See [Js Insert All Pages Except](#bot-defense-policy-js-insert-all-pages-except) below.

&#x2022; [`js_insertion_rules`](#js-insertion-rules) - Optional Block<br>JavaScript Custom Insertion Rules. This defines custom JavaScript insertion rules for Bot Defense Policy<br>See [Js Insertion Rules](#bot-defense-policy-js-insertion-rules) below.

&#x2022; [`mobile_sdk_config`](#mobile-sdk-config) - Optional Block<br>Mobile SDK Configuration. Mobile SDK configuration<br>See [Mobile Sdk Config](#bot-defense-policy-mobile-sdk-config) below.

&#x2022; [`protected_app_endpoints`](#protected-app-endpoints) - Optional Block<br>App Endpoint Type. List of protected endpoints. Limit: Approx '128 endpoints per Load Balancer (LB)' upto 4 LBs, '32 endpoints per LB' after 4 LBs<br>See [Protected App Endpoints](#bot-defense-policy-protected-app-endpoints) below.

#### Bot Defense Policy Js Insert All Pages

A [`js_insert_all_pages`](#bot-defense-policy-js-insert-all-pages) block (within [`bot_defense.policy`](#bot-defense-policy)) supports the following:

&#x2022; [`javascript_location`](#javascript-location) - Optional String  Defaults to `AFTER_HEAD`<br>Possible values are `AFTER_HEAD`, `AFTER_TITLE_END`, `BEFORE_SCRIPT`<br>JavaScript Location. All inside networks. Insert JavaScript after <head> tag Insert JavaScript after </title> tag. Insert JavaScript before first <script> tag

#### Bot Defense Policy Js Insert All Pages Except

A [`js_insert_all_pages_except`](#bot-defense-policy-js-insert-all-pages-except) block (within [`bot_defense.policy`](#bot-defense-policy)) supports the following:

&#x2022; [`exclude_list`](#exclude-list) - Optional Block<br>Exclude Pages. Optional JavaScript insertions exclude list of domain and path matchers<br>See [Exclude List](#bot-defense-policy-js-insert-all-pages-except-exclude-list) below.

&#x2022; [`javascript_location`](#javascript-location) - Optional String  Defaults to `AFTER_HEAD`<br>Possible values are `AFTER_HEAD`, `AFTER_TITLE_END`, `BEFORE_SCRIPT`<br>JavaScript Location. All inside networks. Insert JavaScript after <head> tag Insert JavaScript after </title> tag. Insert JavaScript before first <script> tag

#### Bot Defense Policy Js Insert All Pages Except Exclude List

An [`exclude_list`](#bot-defense-policy-js-insert-all-pages-except-exclude-list) block (within [`bot_defense.policy.js_insert_all_pages_except`](#bot-defense-policy-js-insert-all-pages-except)) supports the following:

&#x2022; [`any_domain`](#any-domain) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`domain`](#domain) - Optional Block<br>Domains. Domains names<br>See [Domain](#bot-defense-policy-js-insert-all-pages-except-exclude-list-domain) below.

&#x2022; [`metadata`](#metadata) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#bot-defense-policy-js-insert-all-pages-except-exclude-list-metadata) below.

&#x2022; [`path`](#path) - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#bot-defense-policy-js-insert-all-pages-except-exclude-list-path) below.

#### Bot Defense Policy Js Insert All Pages Except Exclude List Domain

A [`domain`](#bot-defense-policy-js-insert-all-pages-except-exclude-list-domain) block (within [`bot_defense.policy.js_insert_all_pages_except.exclude_list`](#bot-defense-policy-js-insert-all-pages-except-exclude-list)) supports the following:

&#x2022; [`exact_value`](#exact-value) - Optional String<br>Exact Value. Exact domain name

&#x2022; [`regex_value`](#regex-value) - Optional String<br>Regex Values of Domains. Regular Expression value for the domain name

&#x2022; [`suffix_value`](#suffix-value) - Optional String<br>Suffix Value. Suffix of domain name e.g 'xyz.com' will match '*.xyz.com' and 'xyz.com'

#### Bot Defense Policy Js Insert All Pages Except Exclude List Metadata

A [`metadata`](#bot-defense-policy-js-insert-all-pages-except-exclude-list-metadata) block (within [`bot_defense.policy.js_insert_all_pages_except.exclude_list`](#bot-defense-policy-js-insert-all-pages-except-exclude-list)) supports the following:

&#x2022; [`description`](#description) - Optional String<br>Description. Human readable description

&#x2022; [`name`](#name) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### Bot Defense Policy Js Insert All Pages Except Exclude List Path

A [`path`](#bot-defense-policy-js-insert-all-pages-except-exclude-list-path) block (within [`bot_defense.policy.js_insert_all_pages_except.exclude_list`](#bot-defense-policy-js-insert-all-pages-except-exclude-list)) supports the following:

&#x2022; [`path`](#path) - Optional String<br>Exact. Exact path value to match

&#x2022; [`prefix`](#prefix) - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

&#x2022; [`regex`](#regex) - Optional String<br>Regex. Regular expression of path match (e.g. the value .* will match on all paths)

#### Bot Defense Policy Js Insertion Rules

A [`js_insertion_rules`](#bot-defense-policy-js-insertion-rules) block (within [`bot_defense.policy`](#bot-defense-policy)) supports the following:

&#x2022; [`exclude_list`](#exclude-list) - Optional Block<br>Exclude Paths. Optional JavaScript insertions exclude list of domain and path matchers<br>See [Exclude List](#bot-defense-policy-js-insertion-rules-exclude-list) below.

&#x2022; [`rules`](#rules) - Optional Block<br>JavaScript Insertions. Required list of pages to insert Bot Defense client JavaScript<br>See [Rules](#bot-defense-policy-js-insertion-rules-rules) below.

#### Bot Defense Policy Js Insertion Rules Exclude List

An [`exclude_list`](#bot-defense-policy-js-insertion-rules-exclude-list) block (within [`bot_defense.policy.js_insertion_rules`](#bot-defense-policy-js-insertion-rules)) supports the following:

&#x2022; [`any_domain`](#any-domain) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`domain`](#domain) - Optional Block<br>Domains. Domains names<br>See [Domain](#bot-defense-policy-js-insertion-rules-exclude-list-domain) below.

&#x2022; [`metadata`](#metadata) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#bot-defense-policy-js-insertion-rules-exclude-list-metadata) below.

&#x2022; [`path`](#path) - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#bot-defense-policy-js-insertion-rules-exclude-list-path) below.

#### Bot Defense Policy Js Insertion Rules Exclude List Domain

A [`domain`](#bot-defense-policy-js-insertion-rules-exclude-list-domain) block (within [`bot_defense.policy.js_insertion_rules.exclude_list`](#bot-defense-policy-js-insertion-rules-exclude-list)) supports the following:

&#x2022; [`exact_value`](#exact-value) - Optional String<br>Exact Value. Exact domain name

&#x2022; [`regex_value`](#regex-value) - Optional String<br>Regex Values of Domains. Regular Expression value for the domain name

&#x2022; [`suffix_value`](#suffix-value) - Optional String<br>Suffix Value. Suffix of domain name e.g 'xyz.com' will match '*.xyz.com' and 'xyz.com'

#### Bot Defense Policy Js Insertion Rules Exclude List Metadata

A [`metadata`](#bot-defense-policy-js-insertion-rules-exclude-list-metadata) block (within [`bot_defense.policy.js_insertion_rules.exclude_list`](#bot-defense-policy-js-insertion-rules-exclude-list)) supports the following:

&#x2022; [`description`](#description) - Optional String<br>Description. Human readable description

&#x2022; [`name`](#name) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### Bot Defense Policy Js Insertion Rules Exclude List Path

A [`path`](#bot-defense-policy-js-insertion-rules-exclude-list-path) block (within [`bot_defense.policy.js_insertion_rules.exclude_list`](#bot-defense-policy-js-insertion-rules-exclude-list)) supports the following:

&#x2022; [`path`](#path) - Optional String<br>Exact. Exact path value to match

&#x2022; [`prefix`](#prefix) - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

&#x2022; [`regex`](#regex) - Optional String<br>Regex. Regular expression of path match (e.g. the value .* will match on all paths)

#### Bot Defense Policy Js Insertion Rules Rules

A [`rules`](#bot-defense-policy-js-insertion-rules-rules) block (within [`bot_defense.policy.js_insertion_rules`](#bot-defense-policy-js-insertion-rules)) supports the following:

&#x2022; [`any_domain`](#any-domain) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`domain`](#domain) - Optional Block<br>Domains. Domains names<br>See [Domain](#bot-defense-policy-js-insertion-rules-rules-domain) below.

&#x2022; [`javascript_location`](#javascript-location) - Optional String  Defaults to `AFTER_HEAD`<br>Possible values are `AFTER_HEAD`, `AFTER_TITLE_END`, `BEFORE_SCRIPT`<br>JavaScript Location. All inside networks. Insert JavaScript after <head> tag Insert JavaScript after </title> tag. Insert JavaScript before first <script> tag

&#x2022; [`metadata`](#metadata) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#bot-defense-policy-js-insertion-rules-rules-metadata) below.

&#x2022; [`path`](#path) - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#bot-defense-policy-js-insertion-rules-rules-path) below.

#### Bot Defense Policy Js Insertion Rules Rules Domain

A [`domain`](#bot-defense-policy-js-insertion-rules-rules-domain) block (within [`bot_defense.policy.js_insertion_rules.rules`](#bot-defense-policy-js-insertion-rules-rules)) supports the following:

&#x2022; [`exact_value`](#exact-value) - Optional String<br>Exact Value. Exact domain name

&#x2022; [`regex_value`](#regex-value) - Optional String<br>Regex Values of Domains. Regular Expression value for the domain name

&#x2022; [`suffix_value`](#suffix-value) - Optional String<br>Suffix Value. Suffix of domain name e.g 'xyz.com' will match '*.xyz.com' and 'xyz.com'

#### Bot Defense Policy Js Insertion Rules Rules Metadata

A [`metadata`](#bot-defense-policy-js-insertion-rules-rules-metadata) block (within [`bot_defense.policy.js_insertion_rules.rules`](#bot-defense-policy-js-insertion-rules-rules)) supports the following:

&#x2022; [`description`](#description) - Optional String<br>Description. Human readable description

&#x2022; [`name`](#name) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### Bot Defense Policy Js Insertion Rules Rules Path

A [`path`](#bot-defense-policy-js-insertion-rules-rules-path) block (within [`bot_defense.policy.js_insertion_rules.rules`](#bot-defense-policy-js-insertion-rules-rules)) supports the following:

&#x2022; [`path`](#path) - Optional String<br>Exact. Exact path value to match

&#x2022; [`prefix`](#prefix) - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

&#x2022; [`regex`](#regex) - Optional String<br>Regex. Regular expression of path match (e.g. the value .* will match on all paths)

#### Bot Defense Policy Mobile Sdk Config

A [`mobile_sdk_config`](#bot-defense-policy-mobile-sdk-config) block (within [`bot_defense.policy`](#bot-defense-policy)) supports the following:

&#x2022; [`mobile_identifier`](#mobile-identifier) - Optional Block<br>Mobile Traffic Identifier. Mobile traffic identifier type<br>See [Mobile Identifier](#bot-defense-policy-mobile-sdk-config-mobile-identifier) below.

#### Bot Defense Policy Mobile Sdk Config Mobile Identifier

A [`mobile_identifier`](#bot-defense-policy-mobile-sdk-config-mobile-identifier) block (within [`bot_defense.policy.mobile_sdk_config`](#bot-defense-policy-mobile-sdk-config)) supports the following:

&#x2022; [`headers`](#headers) - Optional Block<br>Headers. Headers that can be used to identify mobile traffic<br>See [Headers](#bot-defense-policy-mobile-sdk-config-mobile-identifier-headers) below.

#### Bot Defense Policy Mobile Sdk Config Mobile Identifier Headers

A [`headers`](#bot-defense-policy-mobile-sdk-config-mobile-identifier-headers) block (within [`bot_defense.policy.mobile_sdk_config.mobile_identifier`](#bot-defense-policy-mobile-sdk-config-mobile-identifier)) supports the following:

&#x2022; [`check_not_present`](#check-not-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`check_present`](#check-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`item`](#item) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#bot-defense-policy-mobile-sdk-config-mobile-identifier-headers-item) below.

&#x2022; [`name`](#name) - Optional String<br>Header Name. A case-insensitive HTTP header name

#### Bot Defense Policy Mobile Sdk Config Mobile Identifier Headers Item

An [`item`](#bot-defense-policy-mobile-sdk-config-mobile-identifier-headers-item) block (within [`bot_defense.policy.mobile_sdk_config.mobile_identifier.headers`](#bot-defense-policy-mobile-sdk-config-mobile-identifier-headers)) supports the following:

&#x2022; [`exact_values`](#exact-values) - Optional List<br>Exact Values. A list of exact values to match the input against

&#x2022; [`regex_values`](#regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

&#x2022; [`transformers`](#transformers) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### Bot Defense Policy Protected App Endpoints

A [`protected_app_endpoints`](#bot-defense-policy-protected-app-endpoints) block (within [`bot_defense.policy`](#bot-defense-policy)) supports the following:

&#x2022; [`allow_good_bots`](#allow-good-bots) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`any_domain`](#any-domain) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`domain`](#domain) - Optional Block<br>Domains. Domains names<br>See [Domain](#bot-defense-policy-protected-app-endpoints-domain) below.

&#x2022; [`flow_label`](#flow-label) - Optional Block<br>Bot Defense Flow Label Category. Bot Defense Flow Label Category allows to associate traffic with selected category<br>See [Flow Label](#bot-defense-policy-protected-app-endpoints-flow-label) below.

&#x2022; [`headers`](#headers) - Optional Block<br>HTTP Headers. A list of predicates for various HTTP headers that need to match. The criteria for matching each HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values are extracted from the request API as a list of strings for each HTTP header type. Note that all specified header predicates must evaluate to true<br>See [Headers](#bot-defense-policy-protected-app-endpoints-headers) below.

&#x2022; [`http_methods`](#http-methods) - Optional List  Defaults to `METHOD_ANY`<br>Possible values are `METHOD_ANY`, `METHOD_GET`, `METHOD_POST`, `METHOD_PUT`, `METHOD_PATCH`, `METHOD_DELETE`, `METHOD_GET_DOCUMENT`<br>HTTP Methods. List of HTTP methods

&#x2022; [`metadata`](#metadata) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#bot-defense-policy-protected-app-endpoints-metadata) below.

&#x2022; [`mitigate_good_bots`](#mitigate-good-bots) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`mitigation`](#mitigation) - Optional Block<br>Bot Mitigation Action. Modify Bot Defense behavior for a matching request<br>See [Mitigation](#bot-defense-policy-protected-app-endpoints-mitigation) below.

&#x2022; [`mobile`](#mobile) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`path`](#path) - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#bot-defense-policy-protected-app-endpoints-path) below.

&#x2022; [`protocol`](#protocol) - Optional String  Defaults to `BOTH`<br>Possible values are `BOTH`, `HTTP`, `HTTPS`<br>URL Scheme. SchemeType is used to indicate URL scheme. - BOTH: BOTH URL scheme for HTTPS:// or `HTTP://.` - HTTP: HTTP URL scheme HTTP:// only. - HTTPS: HTTPS URL scheme HTTPS:// only

&#x2022; [`query_params`](#query-params) - Optional Block<br>HTTP Query Parameters. A list of predicates for all query parameters that need to be matched. The criteria for matching each query parameter are described in individual instances of QueryParameterMatcherType. The actual query parameter values are extracted from the request API as a list of strings for each query parameter name. Note that all specified query parameter predicates must evaluate to true<br>See [Query Params](#bot-defense-policy-protected-app-endpoints-query-params) below.

&#x2022; [`undefined_flow_label`](#undefined-flow-label) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`web`](#web) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`web_mobile`](#web-mobile) - Optional Block<br>Web and Mobile traffic type. Web and Mobile traffic type<br>See [Web Mobile](#bot-defense-policy-protected-app-endpoints-web-mobile) below.

#### Bot Defense Policy Protected App Endpoints Domain

A [`domain`](#bot-defense-policy-protected-app-endpoints-domain) block (within [`bot_defense.policy.protected_app_endpoints`](#bot-defense-policy-protected-app-endpoints)) supports the following:

&#x2022; [`exact_value`](#exact-value) - Optional String<br>Exact Value. Exact domain name

&#x2022; [`regex_value`](#regex-value) - Optional String<br>Regex Values of Domains. Regular Expression value for the domain name

&#x2022; [`suffix_value`](#suffix-value) - Optional String<br>Suffix Value. Suffix of domain name e.g 'xyz.com' will match '*.xyz.com' and 'xyz.com'

#### Bot Defense Policy Protected App Endpoints Flow Label

A [`flow_label`](#bot-defense-policy-protected-app-endpoints-flow-label) block (within [`bot_defense.policy.protected_app_endpoints`](#bot-defense-policy-protected-app-endpoints)) supports the following:

&#x2022; [`account_management`](#account-management) - Optional Block<br>Bot Defense Flow Label Account Management Category. Bot Defense Flow Label Account Management Category<br>See [Account Management](#bot-defense-policy-protected-app-endpoints-flow-label-account-management) below.

&#x2022; [`authentication`](#authentication) - Optional Block<br>Bot Defense Flow Label Authentication Category. Bot Defense Flow Label Authentication Category<br>See [Authentication](#bot-defense-policy-protected-app-endpoints-flow-label-authentication) below.

&#x2022; [`financial_services`](#financial-services) - Optional Block<br>Bot Defense Flow Label Financial Services Category. Bot Defense Flow Label Financial Services Category<br>See [Financial Services](#bot-defense-policy-protected-app-endpoints-flow-label-financial-services) below.

&#x2022; [`flight`](#flight) - Optional Block<br>Bot Defense Flow Label Flight Category. Bot Defense Flow Label Flight Category<br>See [Flight](#bot-defense-policy-protected-app-endpoints-flow-label-flight) below.

&#x2022; [`profile_management`](#profile-management) - Optional Block<br>Bot Defense Flow Label Profile Management Category. Bot Defense Flow Label Profile Management Category<br>See [Profile Management](#bot-defense-policy-protected-app-endpoints-flow-label-profile-management) below.

&#x2022; [`search`](#search) - Optional Block<br>Bot Defense Flow Label Search Category. Bot Defense Flow Label Search Category<br>See [Search](#bot-defense-policy-protected-app-endpoints-flow-label-search) below.

&#x2022; [`shopping_gift_cards`](#shopping-gift-cards) - Optional Block<br>Bot Defense Flow Label Shopping & Gift Cards Category. Bot Defense Flow Label Shopping & Gift Cards Category<br>See [Shopping Gift Cards](#bot-defense-policy-protected-app-endpoints-flow-label-shopping-gift-cards) below.

#### Bot Defense Policy Protected App Endpoints Flow Label Account Management

An [`account_management`](#bot-defense-policy-protected-app-endpoints-flow-label-account-management) block (within [`bot_defense.policy.protected_app_endpoints.flow_label`](#bot-defense-policy-protected-app-endpoints-flow-label)) supports the following:

&#x2022; [`create`](#create) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`password_reset`](#password-reset) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### Bot Defense Policy Protected App Endpoints Flow Label Authentication

An [`authentication`](#bot-defense-policy-protected-app-endpoints-flow-label-authentication) block (within [`bot_defense.policy.protected_app_endpoints.flow_label`](#bot-defense-policy-protected-app-endpoints-flow-label)) supports the following:

&#x2022; [`login`](#login) - Optional Block<br>Bot Defense Transaction Result. Bot Defense Transaction Result<br>See [Login](#bot-defense-policy-protected-app-endpoints-flow-label-authentication-login) below.

&#x2022; [`login_mfa`](#login-mfa) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`login_partner`](#login-partner) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`logout`](#logout) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`token_refresh`](#token-refresh) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### Bot Defense Policy Protected App Endpoints Flow Label Authentication Login

A [`login`](#bot-defense-policy-protected-app-endpoints-flow-label-authentication-login) block (within [`bot_defense.policy.protected_app_endpoints.flow_label.authentication`](#bot-defense-policy-protected-app-endpoints-flow-label-authentication)) supports the following:

&#x2022; [`disable_transaction_result`](#disable-transaction-result) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`transaction_result`](#transaction-result) - Optional Block<br>Bot Defense Transaction Result Type. Bot Defense Transaction ResultType<br>See [Transaction Result](#bot-defense-policy-protected-app-endpoints-flow-label-authentication-login-transaction-result) below.

#### Bot Defense Policy Protected App Endpoints Flow Label Authentication Login Transaction Result

A [`transaction_result`](#bot-defense-policy-protected-app-endpoints-flow-label-authentication-login-transaction-result) block (within [`bot_defense.policy.protected_app_endpoints.flow_label.authentication.login`](#bot-defense-policy-protected-app-endpoints-flow-label-authentication-login)) supports the following:

&#x2022; [`failure_conditions`](#failure-conditions) - Optional Block<br>Failure Conditions. Failure Conditions<br>See [Failure Conditions](#bot-defense-policy-protected-app-endpoints-flow-label-authentication-login-transaction-result-failure-conditions) below.

&#x2022; [`success_conditions`](#success-conditions) - Optional Block<br>Success Conditions. Success Conditions<br>See [Success Conditions](#bot-defense-policy-protected-app-endpoints-flow-label-authentication-login-transaction-result-success-conditions) below.

#### Bot Defense Policy Protected App Endpoints Flow Label Authentication Login Transaction Result Failure Conditions

A [`failure_conditions`](#bot-defense-policy-protected-app-endpoints-flow-label-authentication-login-transaction-result-failure-conditions) block (within [`bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result`](#bot-defense-policy-protected-app-endpoints-flow-label-authentication-login-transaction-result)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Header Name. A case-insensitive HTTP header name

&#x2022; [`regex_values`](#regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

&#x2022; [`status`](#status) - Optional String  Defaults to `EmptyStatusCode`<br>Possible values are `EmptyStatusCode`, `Continue`, `OK`, `Created`, `Accepted`, `NonAuthoritativeInformation`, `NoContent`, `ResetContent`, `PartialContent`, `MultiStatus`, `AlreadyReported`, `IMUsed`, `MultipleChoices`, `MovedPermanently`, `Found`, `SeeOther`, `NotModified`, `UseProxy`, `TemporaryRedirect`, `PermanentRedirect`, `BadRequest`, `Unauthorized`, `PaymentRequired`, `Forbidden`, `NotFound`, `MethodNotAllowed`, `NotAcceptable`, `ProxyAuthenticationRequired`, `RequestTimeout`, `Conflict`, `Gone`, `LengthRequired`, `PreconditionFailed`, `PayloadTooLarge`, `URITooLong`, `UnsupportedMediaType`, `RangeNotSatisfiable`, `ExpectationFailed`, `MisdirectedRequest`, `UnprocessableEntity`, `Locked`, `FailedDependency`, `UpgradeRequired`, `PreconditionRequired`, `TooManyRequests`, `RequestHeaderFieldsTooLarge`, `InternalServerError`, `NotImplemented`, `BadGateway`, `ServiceUnavailable`, `GatewayTimeout`, `HTTPVersionNotSupported`, `VariantAlsoNegotiates`, `InsufficientStorage`, `LoopDetected`, `NotExtended`, `NetworkAuthenticationRequired`<br>HTTP Status Code. HTTP response status codes EmptyStatusCode response codes means it is not specified Continue status code OK status code Created status code Accepted status code Non Authoritative Information status code No Content status code Reset Content status code Partial Content status code Multi Status status code Already Reported status code Im Used status code Multiple Choices status code Moved Permanently status code Found status code See Other status code Not Modified status code Use Proxy status code Temporary Redirect status code Permanent Redirect status code Bad Request status code Unauthorized status code Payment Required status code Forbidden status code Not Found status code Method Not Allowed status code Not Acceptable status code Proxy Authentication Required status code Request Timeout status code Conflict status code Gone status code Length Required status code Precondition Failed status code Payload Too Large status code URI Too Long status code Unsupported Media Type status code Range Not Satisfiable status code Expectation Failed status code Misdirected Request status code Unprocessable Entity status code Locked status code Failed Dependency status code Upgrade Required status code Precondition Required status code Too Many Requests status code Request Header Fields Too Large status code Internal Server Error status code Not Implemented status code Bad Gateway status code Service Unavailable status code Gateway Timeout status code HTTP Version Not Supported status code Variant Also Negotiates status code Insufficient Storage status code Loop Detected status code Not Extended status code Network Authentication Required status code

#### Bot Defense Policy Protected App Endpoints Flow Label Authentication Login Transaction Result Success Conditions

A [`success_conditions`](#bot-defense-policy-protected-app-endpoints-flow-label-authentication-login-transaction-result-success-conditions) block (within [`bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result`](#bot-defense-policy-protected-app-endpoints-flow-label-authentication-login-transaction-result)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Header Name. A case-insensitive HTTP header name

&#x2022; [`regex_values`](#regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

&#x2022; [`status`](#status) - Optional String  Defaults to `EmptyStatusCode`<br>Possible values are `EmptyStatusCode`, `Continue`, `OK`, `Created`, `Accepted`, `NonAuthoritativeInformation`, `NoContent`, `ResetContent`, `PartialContent`, `MultiStatus`, `AlreadyReported`, `IMUsed`, `MultipleChoices`, `MovedPermanently`, `Found`, `SeeOther`, `NotModified`, `UseProxy`, `TemporaryRedirect`, `PermanentRedirect`, `BadRequest`, `Unauthorized`, `PaymentRequired`, `Forbidden`, `NotFound`, `MethodNotAllowed`, `NotAcceptable`, `ProxyAuthenticationRequired`, `RequestTimeout`, `Conflict`, `Gone`, `LengthRequired`, `PreconditionFailed`, `PayloadTooLarge`, `URITooLong`, `UnsupportedMediaType`, `RangeNotSatisfiable`, `ExpectationFailed`, `MisdirectedRequest`, `UnprocessableEntity`, `Locked`, `FailedDependency`, `UpgradeRequired`, `PreconditionRequired`, `TooManyRequests`, `RequestHeaderFieldsTooLarge`, `InternalServerError`, `NotImplemented`, `BadGateway`, `ServiceUnavailable`, `GatewayTimeout`, `HTTPVersionNotSupported`, `VariantAlsoNegotiates`, `InsufficientStorage`, `LoopDetected`, `NotExtended`, `NetworkAuthenticationRequired`<br>HTTP Status Code. HTTP response status codes EmptyStatusCode response codes means it is not specified Continue status code OK status code Created status code Accepted status code Non Authoritative Information status code No Content status code Reset Content status code Partial Content status code Multi Status status code Already Reported status code Im Used status code Multiple Choices status code Moved Permanently status code Found status code See Other status code Not Modified status code Use Proxy status code Temporary Redirect status code Permanent Redirect status code Bad Request status code Unauthorized status code Payment Required status code Forbidden status code Not Found status code Method Not Allowed status code Not Acceptable status code Proxy Authentication Required status code Request Timeout status code Conflict status code Gone status code Length Required status code Precondition Failed status code Payload Too Large status code URI Too Long status code Unsupported Media Type status code Range Not Satisfiable status code Expectation Failed status code Misdirected Request status code Unprocessable Entity status code Locked status code Failed Dependency status code Upgrade Required status code Precondition Required status code Too Many Requests status code Request Header Fields Too Large status code Internal Server Error status code Not Implemented status code Bad Gateway status code Service Unavailable status code Gateway Timeout status code HTTP Version Not Supported status code Variant Also Negotiates status code Insufficient Storage status code Loop Detected status code Not Extended status code Network Authentication Required status code

#### Bot Defense Policy Protected App Endpoints Flow Label Financial Services

A [`financial_services`](#bot-defense-policy-protected-app-endpoints-flow-label-financial-services) block (within [`bot_defense.policy.protected_app_endpoints.flow_label`](#bot-defense-policy-protected-app-endpoints-flow-label)) supports the following:

&#x2022; [`apply`](#apply) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`money_transfer`](#money-transfer) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### Bot Defense Policy Protected App Endpoints Flow Label Flight

A [`flight`](#bot-defense-policy-protected-app-endpoints-flow-label-flight) block (within [`bot_defense.policy.protected_app_endpoints.flow_label`](#bot-defense-policy-protected-app-endpoints-flow-label)) supports the following:

&#x2022; [`checkin`](#checkin) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### Bot Defense Policy Protected App Endpoints Flow Label Profile Management

A [`profile_management`](#bot-defense-policy-protected-app-endpoints-flow-label-profile-management) block (within [`bot_defense.policy.protected_app_endpoints.flow_label`](#bot-defense-policy-protected-app-endpoints-flow-label)) supports the following:

&#x2022; [`create`](#create) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`update`](#update) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`view`](#view) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### Bot Defense Policy Protected App Endpoints Flow Label Search

A [`search`](#bot-defense-policy-protected-app-endpoints-flow-label-search) block (within [`bot_defense.policy.protected_app_endpoints.flow_label`](#bot-defense-policy-protected-app-endpoints-flow-label)) supports the following:

&#x2022; [`flight_search`](#flight-search) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`product_search`](#product-search) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`reservation_search`](#reservation-search) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`room_search`](#room-search) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### Bot Defense Policy Protected App Endpoints Flow Label Shopping Gift Cards

A [`shopping_gift_cards`](#bot-defense-policy-protected-app-endpoints-flow-label-shopping-gift-cards) block (within [`bot_defense.policy.protected_app_endpoints.flow_label`](#bot-defense-policy-protected-app-endpoints-flow-label)) supports the following:

&#x2022; [`gift_card_make_purchase_with_gift_card`](#gift-card-make-purchase-with-gift-card) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`gift_card_validation`](#gift-card-validation) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`shop_add_to_cart`](#shop-add-to-cart) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`shop_checkout`](#shop-checkout) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`shop_choose_seat`](#shop-choose-seat) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`shop_enter_drawing_submission`](#shop-enter-drawing-submission) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`shop_make_payment`](#shop-make-payment) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`shop_order`](#shop-order) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`shop_price_inquiry`](#shop-price-inquiry) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`shop_promo_code_validation`](#shop-promo-code-validation) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`shop_purchase_gift_card`](#shop-purchase-gift-card) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`shop_update_quantity`](#shop-update-quantity) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### Bot Defense Policy Protected App Endpoints Headers

A [`headers`](#bot-defense-policy-protected-app-endpoints-headers) block (within [`bot_defense.policy.protected_app_endpoints`](#bot-defense-policy-protected-app-endpoints)) supports the following:

&#x2022; [`check_not_present`](#check-not-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`check_present`](#check-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`invert_matcher`](#invert-matcher) - Optional Bool<br>Invert Header Matcher. Invert the match result

&#x2022; [`item`](#item) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#bot-defense-policy-protected-app-endpoints-headers-item) below.

&#x2022; [`name`](#name) - Optional String<br>Header Name. A case-insensitive HTTP header name

#### Bot Defense Policy Protected App Endpoints Headers Item

An [`item`](#bot-defense-policy-protected-app-endpoints-headers-item) block (within [`bot_defense.policy.protected_app_endpoints.headers`](#bot-defense-policy-protected-app-endpoints-headers)) supports the following:

&#x2022; [`exact_values`](#exact-values) - Optional List<br>Exact Values. A list of exact values to match the input against

&#x2022; [`regex_values`](#regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

&#x2022; [`transformers`](#transformers) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### Bot Defense Policy Protected App Endpoints Metadata

A [`metadata`](#bot-defense-policy-protected-app-endpoints-metadata) block (within [`bot_defense.policy.protected_app_endpoints`](#bot-defense-policy-protected-app-endpoints)) supports the following:

&#x2022; [`description`](#description) - Optional String<br>Description. Human readable description

&#x2022; [`name`](#name) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### Bot Defense Policy Protected App Endpoints Mitigation

A [`mitigation`](#bot-defense-policy-protected-app-endpoints-mitigation) block (within [`bot_defense.policy.protected_app_endpoints`](#bot-defense-policy-protected-app-endpoints)) supports the following:

&#x2022; [`block`](#block) - Optional Block<br>Block bot mitigation. Block request and respond with custom content<br>See [Block](#bot-defense-policy-protected-app-endpoints-mitigation-block) below.

&#x2022; [`flag`](#flag) - Optional Block<br>Select Flag Bot Mitigation Action. Flag mitigation action<br>See [Flag](#bot-defense-policy-protected-app-endpoints-mitigation-flag) below.

&#x2022; [`redirect`](#redirect) - Optional Block<br>Redirect bot mitigation. Redirect request to a custom URI<br>See [Redirect](#bot-defense-policy-protected-app-endpoints-mitigation-redirect) below.

#### Bot Defense Policy Protected App Endpoints Mitigation Block

A [`block`](#bot-defense-policy-protected-app-endpoints-mitigation-block) block (within [`bot_defense.policy.protected_app_endpoints.mitigation`](#bot-defense-policy-protected-app-endpoints-mitigation)) supports the following:

&#x2022; [`body`](#body) - Optional String<br>Body. Custom body message is of type uri_ref. Currently supported URL schemes is string:///. For string:/// scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. 'Your request was blocked' or it can be HTML paragraph or a body string encoded as base64 string E.g. '<p> Your request was blocked </p>'. Base64 encoded string for this HTML is 'LzxwPiBZb3VyIHJlcXVlc3Qgd2FzIGJsb2NrZWQgPC9wPg=='

&#x2022; [`status`](#status) - Optional String  Defaults to `EmptyStatusCode`<br>Possible values are `EmptyStatusCode`, `Continue`, `OK`, `Created`, `Accepted`, `NonAuthoritativeInformation`, `NoContent`, `ResetContent`, `PartialContent`, `MultiStatus`, `AlreadyReported`, `IMUsed`, `MultipleChoices`, `MovedPermanently`, `Found`, `SeeOther`, `NotModified`, `UseProxy`, `TemporaryRedirect`, `PermanentRedirect`, `BadRequest`, `Unauthorized`, `PaymentRequired`, `Forbidden`, `NotFound`, `MethodNotAllowed`, `NotAcceptable`, `ProxyAuthenticationRequired`, `RequestTimeout`, `Conflict`, `Gone`, `LengthRequired`, `PreconditionFailed`, `PayloadTooLarge`, `URITooLong`, `UnsupportedMediaType`, `RangeNotSatisfiable`, `ExpectationFailed`, `MisdirectedRequest`, `UnprocessableEntity`, `Locked`, `FailedDependency`, `UpgradeRequired`, `PreconditionRequired`, `TooManyRequests`, `RequestHeaderFieldsTooLarge`, `InternalServerError`, `NotImplemented`, `BadGateway`, `ServiceUnavailable`, `GatewayTimeout`, `HTTPVersionNotSupported`, `VariantAlsoNegotiates`, `InsufficientStorage`, `LoopDetected`, `NotExtended`, `NetworkAuthenticationRequired`<br>HTTP Status Code. HTTP response status codes EmptyStatusCode response codes means it is not specified Continue status code OK status code Created status code Accepted status code Non Authoritative Information status code No Content status code Reset Content status code Partial Content status code Multi Status status code Already Reported status code Im Used status code Multiple Choices status code Moved Permanently status code Found status code See Other status code Not Modified status code Use Proxy status code Temporary Redirect status code Permanent Redirect status code Bad Request status code Unauthorized status code Payment Required status code Forbidden status code Not Found status code Method Not Allowed status code Not Acceptable status code Proxy Authentication Required status code Request Timeout status code Conflict status code Gone status code Length Required status code Precondition Failed status code Payload Too Large status code URI Too Long status code Unsupported Media Type status code Range Not Satisfiable status code Expectation Failed status code Misdirected Request status code Unprocessable Entity status code Locked status code Failed Dependency status code Upgrade Required status code Precondition Required status code Too Many Requests status code Request Header Fields Too Large status code Internal Server Error status code Not Implemented status code Bad Gateway status code Service Unavailable status code Gateway Timeout status code HTTP Version Not Supported status code Variant Also Negotiates status code Insufficient Storage status code Loop Detected status code Not Extended status code Network Authentication Required status code

#### Bot Defense Policy Protected App Endpoints Mitigation Flag

A [`flag`](#bot-defense-policy-protected-app-endpoints-mitigation-flag) block (within [`bot_defense.policy.protected_app_endpoints.mitigation`](#bot-defense-policy-protected-app-endpoints-mitigation)) supports the following:

&#x2022; [`append_headers`](#append-headers) - Optional Block<br>Append Flag Mitigation Headers. Append flag mitigation headers to forwarded request<br>See [Append Headers](#bot-defense-policy-protected-app-endpoints-mitigation-flag-append-headers) below.

&#x2022; [`no_headers`](#no-headers) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### Bot Defense Policy Protected App Endpoints Mitigation Flag Append Headers

An [`append_headers`](#bot-defense-policy-protected-app-endpoints-mitigation-flag-append-headers) block (within [`bot_defense.policy.protected_app_endpoints.mitigation.flag`](#bot-defense-policy-protected-app-endpoints-mitigation-flag)) supports the following:

&#x2022; [`auto_type_header_name`](#auto-type-header-name) - Optional String<br>Automation Type Header Name. A case-insensitive HTTP header name

&#x2022; [`inference_header_name`](#inference-header-name) - Optional String<br>Inference Header Name. A case-insensitive HTTP header name

#### Bot Defense Policy Protected App Endpoints Mitigation Redirect

A [`redirect`](#bot-defense-policy-protected-app-endpoints-mitigation-redirect) block (within [`bot_defense.policy.protected_app_endpoints.mitigation`](#bot-defense-policy-protected-app-endpoints-mitigation)) supports the following:

&#x2022; [`uri`](#uri) - Optional String<br>URI. URI location for redirect may be relative or absolute

#### Bot Defense Policy Protected App Endpoints Path

A [`path`](#bot-defense-policy-protected-app-endpoints-path) block (within [`bot_defense.policy.protected_app_endpoints`](#bot-defense-policy-protected-app-endpoints)) supports the following:

&#x2022; [`path`](#path) - Optional String<br>Exact. Exact path value to match

&#x2022; [`prefix`](#prefix) - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

&#x2022; [`regex`](#regex) - Optional String<br>Regex. Regular expression of path match (e.g. the value .* will match on all paths)

#### Bot Defense Policy Protected App Endpoints Query Params

A [`query_params`](#bot-defense-policy-protected-app-endpoints-query-params) block (within [`bot_defense.policy.protected_app_endpoints`](#bot-defense-policy-protected-app-endpoints)) supports the following:

&#x2022; [`check_not_present`](#check-not-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`check_present`](#check-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`invert_matcher`](#invert-matcher) - Optional Bool<br>Invert Query Parameter Matcher. Invert the match result

&#x2022; [`item`](#item) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#bot-defense-policy-protected-app-endpoints-query-params-item) below.

&#x2022; [`key`](#key) - Optional String<br>Query Parameter Name. A case-sensitive HTTP query parameter name

#### Bot Defense Policy Protected App Endpoints Query Params Item

An [`item`](#bot-defense-policy-protected-app-endpoints-query-params-item) block (within [`bot_defense.policy.protected_app_endpoints.query_params`](#bot-defense-policy-protected-app-endpoints-query-params)) supports the following:

&#x2022; [`exact_values`](#exact-values) - Optional List<br>Exact Values. A list of exact values to match the input against

&#x2022; [`regex_values`](#regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

&#x2022; [`transformers`](#transformers) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### Bot Defense Policy Protected App Endpoints Web Mobile

A [`web_mobile`](#bot-defense-policy-protected-app-endpoints-web-mobile) block (within [`bot_defense.policy.protected_app_endpoints`](#bot-defense-policy-protected-app-endpoints)) supports the following:

&#x2022; [`mobile_identifier`](#mobile-identifier) - Optional String  Defaults to `HEADERS`<br>Mobile Identifier. Mobile identifier type - HEADERS: Headers Headers. The only possible value is `HEADERS`

#### Bot Defense Advanced

A [`bot_defense_advanced`](#bot-defense-advanced) block supports the following:

&#x2022; [`disable_js_insert`](#disable-js-insert) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`disable_mobile_sdk`](#disable-mobile-sdk) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`js_insert_all_pages`](#js-insert-all-pages) - Optional Block<br>Insert Bot Defense JavaScript in All Pages. Insert Bot Defense JavaScript in all pages<br>See [Js Insert All Pages](#bot-defense-advanced-js-insert-all-pages) below.

&#x2022; [`js_insert_all_pages_except`](#js-insert-all-pages-except) - Optional Block<br>Insert JavaScript in All Pages with the Exceptions. Insert Bot Defense JavaScript in all pages with the exceptions<br>See [Js Insert All Pages Except](#bot-defense-advanced-js-insert-all-pages-except) below.

&#x2022; [`js_insertion_rules`](#js-insertion-rules) - Optional Block<br>JavaScript Custom Insertion Rules. This defines custom JavaScript insertion rules for Bot Defense Policy<br>See [Js Insertion Rules](#bot-defense-advanced-js-insertion-rules) below.

&#x2022; [`mobile`](#mobile) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Mobile](#bot-defense-advanced-mobile) below.

&#x2022; [`mobile_sdk_config`](#mobile-sdk-config) - Optional Block<br>Mobile Request Identifier Headers. Mobile Request Identifier Headers<br>See [Mobile Sdk Config](#bot-defense-advanced-mobile-sdk-config) below.

&#x2022; [`web`](#web) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Web](#bot-defense-advanced-web) below.

#### Bot Defense Advanced Js Insert All Pages

A [`js_insert_all_pages`](#bot-defense-advanced-js-insert-all-pages) block (within [`bot_defense_advanced`](#bot-defense-advanced)) supports the following:

&#x2022; [`javascript_location`](#javascript-location) - Optional String  Defaults to `AFTER_HEAD`<br>Possible values are `AFTER_HEAD`, `AFTER_TITLE_END`, `BEFORE_SCRIPT`<br>JavaScript Location. All inside networks. Insert JavaScript after <head> tag Insert JavaScript after </title> tag. Insert JavaScript before first <script> tag

#### Bot Defense Advanced Js Insert All Pages Except

A [`js_insert_all_pages_except`](#bot-defense-advanced-js-insert-all-pages-except) block (within [`bot_defense_advanced`](#bot-defense-advanced)) supports the following:

&#x2022; [`exclude_list`](#exclude-list) - Optional Block<br>Exclude Pages. Optional JavaScript insertions exclude list of domain and path matchers<br>See [Exclude List](#bot-defense-advanced-js-insert-all-pages-except-exclude-list) below.

&#x2022; [`javascript_location`](#javascript-location) - Optional String  Defaults to `AFTER_HEAD`<br>Possible values are `AFTER_HEAD`, `AFTER_TITLE_END`, `BEFORE_SCRIPT`<br>JavaScript Location. All inside networks. Insert JavaScript after <head> tag Insert JavaScript after </title> tag. Insert JavaScript before first <script> tag

#### Bot Defense Advanced Js Insert All Pages Except Exclude List

An [`exclude_list`](#bot-defense-advanced-js-insert-all-pages-except-exclude-list) block (within [`bot_defense_advanced.js_insert_all_pages_except`](#bot-defense-advanced-js-insert-all-pages-except)) supports the following:

&#x2022; [`any_domain`](#any-domain) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`domain`](#domain) - Optional Block<br>Domains. Domains names<br>See [Domain](#bot-defense-advanced-js-insert-all-pages-except-exclude-list-domain) below.

&#x2022; [`metadata`](#metadata) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#bot-defense-advanced-js-insert-all-pages-except-exclude-list-metadata) below.

&#x2022; [`path`](#path) - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#bot-defense-advanced-js-insert-all-pages-except-exclude-list-path) below.

#### Bot Defense Advanced Js Insert All Pages Except Exclude List Domain

A [`domain`](#bot-defense-advanced-js-insert-all-pages-except-exclude-list-domain) block (within [`bot_defense_advanced.js_insert_all_pages_except.exclude_list`](#bot-defense-advanced-js-insert-all-pages-except-exclude-list)) supports the following:

&#x2022; [`exact_value`](#exact-value) - Optional String<br>Exact Value. Exact domain name

&#x2022; [`regex_value`](#regex-value) - Optional String<br>Regex Values of Domains. Regular Expression value for the domain name

&#x2022; [`suffix_value`](#suffix-value) - Optional String<br>Suffix Value. Suffix of domain name e.g 'xyz.com' will match '*.xyz.com' and 'xyz.com'

#### Bot Defense Advanced Js Insert All Pages Except Exclude List Metadata

A [`metadata`](#bot-defense-advanced-js-insert-all-pages-except-exclude-list-metadata) block (within [`bot_defense_advanced.js_insert_all_pages_except.exclude_list`](#bot-defense-advanced-js-insert-all-pages-except-exclude-list)) supports the following:

&#x2022; [`description`](#description) - Optional String<br>Description. Human readable description

&#x2022; [`name`](#name) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### Bot Defense Advanced Js Insert All Pages Except Exclude List Path

A [`path`](#bot-defense-advanced-js-insert-all-pages-except-exclude-list-path) block (within [`bot_defense_advanced.js_insert_all_pages_except.exclude_list`](#bot-defense-advanced-js-insert-all-pages-except-exclude-list)) supports the following:

&#x2022; [`path`](#path) - Optional String<br>Exact. Exact path value to match

&#x2022; [`prefix`](#prefix) - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

&#x2022; [`regex`](#regex) - Optional String<br>Regex. Regular expression of path match (e.g. the value .* will match on all paths)

#### Bot Defense Advanced Js Insertion Rules

A [`js_insertion_rules`](#bot-defense-advanced-js-insertion-rules) block (within [`bot_defense_advanced`](#bot-defense-advanced)) supports the following:

&#x2022; [`exclude_list`](#exclude-list) - Optional Block<br>Exclude Paths. Optional JavaScript insertions exclude list of domain and path matchers<br>See [Exclude List](#bot-defense-advanced-js-insertion-rules-exclude-list) below.

&#x2022; [`rules`](#rules) - Optional Block<br>JavaScript Insertions. Required list of pages to insert Bot Defense client JavaScript<br>See [Rules](#bot-defense-advanced-js-insertion-rules-rules) below.

#### Bot Defense Advanced Js Insertion Rules Exclude List

An [`exclude_list`](#bot-defense-advanced-js-insertion-rules-exclude-list) block (within [`bot_defense_advanced.js_insertion_rules`](#bot-defense-advanced-js-insertion-rules)) supports the following:

&#x2022; [`any_domain`](#any-domain) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`domain`](#domain) - Optional Block<br>Domains. Domains names<br>See [Domain](#bot-defense-advanced-js-insertion-rules-exclude-list-domain) below.

&#x2022; [`metadata`](#metadata) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#bot-defense-advanced-js-insertion-rules-exclude-list-metadata) below.

&#x2022; [`path`](#path) - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#bot-defense-advanced-js-insertion-rules-exclude-list-path) below.

#### Bot Defense Advanced Js Insertion Rules Exclude List Domain

A [`domain`](#bot-defense-advanced-js-insertion-rules-exclude-list-domain) block (within [`bot_defense_advanced.js_insertion_rules.exclude_list`](#bot-defense-advanced-js-insertion-rules-exclude-list)) supports the following:

&#x2022; [`exact_value`](#exact-value) - Optional String<br>Exact Value. Exact domain name

&#x2022; [`regex_value`](#regex-value) - Optional String<br>Regex Values of Domains. Regular Expression value for the domain name

&#x2022; [`suffix_value`](#suffix-value) - Optional String<br>Suffix Value. Suffix of domain name e.g 'xyz.com' will match '*.xyz.com' and 'xyz.com'

#### Bot Defense Advanced Js Insertion Rules Exclude List Metadata

A [`metadata`](#bot-defense-advanced-js-insertion-rules-exclude-list-metadata) block (within [`bot_defense_advanced.js_insertion_rules.exclude_list`](#bot-defense-advanced-js-insertion-rules-exclude-list)) supports the following:

&#x2022; [`description`](#description) - Optional String<br>Description. Human readable description

&#x2022; [`name`](#name) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### Bot Defense Advanced Js Insertion Rules Exclude List Path

A [`path`](#bot-defense-advanced-js-insertion-rules-exclude-list-path) block (within [`bot_defense_advanced.js_insertion_rules.exclude_list`](#bot-defense-advanced-js-insertion-rules-exclude-list)) supports the following:

&#x2022; [`path`](#path) - Optional String<br>Exact. Exact path value to match

&#x2022; [`prefix`](#prefix) - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

&#x2022; [`regex`](#regex) - Optional String<br>Regex. Regular expression of path match (e.g. the value .* will match on all paths)

#### Bot Defense Advanced Js Insertion Rules Rules

A [`rules`](#bot-defense-advanced-js-insertion-rules-rules) block (within [`bot_defense_advanced.js_insertion_rules`](#bot-defense-advanced-js-insertion-rules)) supports the following:

&#x2022; [`any_domain`](#any-domain) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`domain`](#domain) - Optional Block<br>Domains. Domains names<br>See [Domain](#bot-defense-advanced-js-insertion-rules-rules-domain) below.

&#x2022; [`javascript_location`](#javascript-location) - Optional String  Defaults to `AFTER_HEAD`<br>Possible values are `AFTER_HEAD`, `AFTER_TITLE_END`, `BEFORE_SCRIPT`<br>JavaScript Location. All inside networks. Insert JavaScript after <head> tag Insert JavaScript after </title> tag. Insert JavaScript before first <script> tag

&#x2022; [`metadata`](#metadata) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#bot-defense-advanced-js-insertion-rules-rules-metadata) below.

&#x2022; [`path`](#path) - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#bot-defense-advanced-js-insertion-rules-rules-path) below.

#### Bot Defense Advanced Js Insertion Rules Rules Domain

A [`domain`](#bot-defense-advanced-js-insertion-rules-rules-domain) block (within [`bot_defense_advanced.js_insertion_rules.rules`](#bot-defense-advanced-js-insertion-rules-rules)) supports the following:

&#x2022; [`exact_value`](#exact-value) - Optional String<br>Exact Value. Exact domain name

&#x2022; [`regex_value`](#regex-value) - Optional String<br>Regex Values of Domains. Regular Expression value for the domain name

&#x2022; [`suffix_value`](#suffix-value) - Optional String<br>Suffix Value. Suffix of domain name e.g 'xyz.com' will match '*.xyz.com' and 'xyz.com'

#### Bot Defense Advanced Js Insertion Rules Rules Metadata

A [`metadata`](#bot-defense-advanced-js-insertion-rules-rules-metadata) block (within [`bot_defense_advanced.js_insertion_rules.rules`](#bot-defense-advanced-js-insertion-rules-rules)) supports the following:

&#x2022; [`description`](#description) - Optional String<br>Description. Human readable description

&#x2022; [`name`](#name) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### Bot Defense Advanced Js Insertion Rules Rules Path

A [`path`](#bot-defense-advanced-js-insertion-rules-rules-path) block (within [`bot_defense_advanced.js_insertion_rules.rules`](#bot-defense-advanced-js-insertion-rules-rules)) supports the following:

&#x2022; [`path`](#path) - Optional String<br>Exact. Exact path value to match

&#x2022; [`prefix`](#prefix) - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

&#x2022; [`regex`](#regex) - Optional String<br>Regex. Regular expression of path match (e.g. the value .* will match on all paths)

#### Bot Defense Advanced Mobile

A [`mobile`](#bot-defense-advanced-mobile) block (within [`bot_defense_advanced`](#bot-defense-advanced)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Bot Defense Advanced Mobile Sdk Config

A [`mobile_sdk_config`](#bot-defense-advanced-mobile-sdk-config) block (within [`bot_defense_advanced`](#bot-defense-advanced)) supports the following:

&#x2022; [`mobile_identifier`](#mobile-identifier) - Optional Block<br>Mobile Traffic Identifier. Mobile traffic identifier type<br>See [Mobile Identifier](#bot-defense-advanced-mobile-sdk-config-mobile-identifier) below.

#### Bot Defense Advanced Mobile Sdk Config Mobile Identifier

A [`mobile_identifier`](#bot-defense-advanced-mobile-sdk-config-mobile-identifier) block (within [`bot_defense_advanced.mobile_sdk_config`](#bot-defense-advanced-mobile-sdk-config)) supports the following:

&#x2022; [`headers`](#headers) - Optional Block<br>Headers. Headers that can be used to identify mobile traffic<br>See [Headers](#bot-defense-advanced-mobile-sdk-config-mobile-identifier-headers) below.

#### Bot Defense Advanced Mobile Sdk Config Mobile Identifier Headers

A [`headers`](#bot-defense-advanced-mobile-sdk-config-mobile-identifier-headers) block (within [`bot_defense_advanced.mobile_sdk_config.mobile_identifier`](#bot-defense-advanced-mobile-sdk-config-mobile-identifier)) supports the following:

&#x2022; [`check_not_present`](#check-not-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`check_present`](#check-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`item`](#item) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#bot-defense-advanced-mobile-sdk-config-mobile-identifier-headers-item) below.

&#x2022; [`name`](#name) - Optional String<br>Header Name. A case-insensitive HTTP header name

#### Bot Defense Advanced Mobile Sdk Config Mobile Identifier Headers Item

An [`item`](#bot-defense-advanced-mobile-sdk-config-mobile-identifier-headers-item) block (within [`bot_defense_advanced.mobile_sdk_config.mobile_identifier.headers`](#bot-defense-advanced-mobile-sdk-config-mobile-identifier-headers)) supports the following:

&#x2022; [`exact_values`](#exact-values) - Optional List<br>Exact Values. A list of exact values to match the input against

&#x2022; [`regex_values`](#regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

&#x2022; [`transformers`](#transformers) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### Bot Defense Advanced Web

A [`web`](#bot-defense-advanced-web) block (within [`bot_defense_advanced`](#bot-defense-advanced)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Caching Policy

A [`caching_policy`](#caching-policy) block supports the following:

&#x2022; [`custom_cache_rule`](#custom-cache-rule) - Optional Block<br>Custom Cache Rules. Caching policies for CDN<br>See [Custom Cache Rule](#caching-policy-custom-cache-rule) below.

&#x2022; [`default_cache_action`](#default-cache-action) - Optional Block<br>Default Cache Behaviour. This defines a Default Cache Action<br>See [Default Cache Action](#caching-policy-default-cache-action) below.

#### Caching Policy Custom Cache Rule

A [`custom_cache_rule`](#caching-policy-custom-cache-rule) block (within [`caching_policy`](#caching-policy)) supports the following:

&#x2022; [`cdn_cache_rules`](#cdn-cache-rules) - Optional Block<br>CDN Cache Rule. Reference to CDN Cache Rule configuration object<br>See [CDN Cache Rules](#caching-policy-custom-cache-rule-cdn-cache-rules) below.

#### Caching Policy Custom Cache Rule CDN Cache Rules

A [`cdn_cache_rules`](#caching-policy-custom-cache-rule-cdn-cache-rules) block (within [`caching_policy.custom_cache_rule`](#caching-policy-custom-cache-rule)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Caching Policy Default Cache Action

A [`default_cache_action`](#caching-policy-default-cache-action) block (within [`caching_policy`](#caching-policy)) supports the following:

&#x2022; [`cache_disabled`](#cache-disabled) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`cache_ttl_default`](#cache-ttl-default) - Optional String<br>Fallback Cache TTL (d/ h/ m). Use Cache TTL Provided by Origin, and set a contigency TTL value in case one is not provided

&#x2022; [`cache_ttl_override`](#cache-ttl-override) - Optional String<br>Override Cache TTL (d/ h/ m/ s). Always override the Cahce TTL provided by Origin

#### Captcha Challenge

A [`captcha_challenge`](#captcha-challenge) block supports the following:

&#x2022; [`cookie_expiry`](#cookie-expiry) - Optional Number<br>Cookie Expiration Period. Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new challenge

&#x2022; [`custom_page`](#custom-page) - Optional String<br>Custom message for Captcha Challenge. Custom message is of type uri_ref. Currently supported URL schemes is string:///. For string:/// scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. 'Please Wait.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '<p> Please Wait </p>'. Base64 encoded string for this HTML is 'PHA+IFBsZWFzZSBXYWl0IDwvcD4='

#### Client Side Defense

A [`client_side_defense`](#client-side-defense) block supports the following:

&#x2022; [`policy`](#policy) - Optional Block<br>Client-Side Defense Policy. This defines various configuration options for Client-Side Defense policy<br>See [Policy](#client-side-defense-policy) below.

#### Client Side Defense Policy

A [`policy`](#client-side-defense-policy) block (within [`client_side_defense`](#client-side-defense)) supports the following:

&#x2022; [`disable_js_insert`](#disable-js-insert) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`js_insert_all_pages`](#js-insert-all-pages) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`js_insert_all_pages_except`](#js-insert-all-pages-except) - Optional Block<br>Insert JavaScript in All Pages with the Exceptions. Insert Client-Side Defense JavaScript in all pages with the exceptions<br>See [Js Insert All Pages Except](#client-side-defense-policy-js-insert-all-pages-except) below.

&#x2022; [`js_insertion_rules`](#js-insertion-rules) - Optional Block<br>JavaScript Custom Insertion Rules. This defines custom JavaScript insertion rules for Client-Side Defense Policy<br>See [Js Insertion Rules](#client-side-defense-policy-js-insertion-rules) below.

#### Client Side Defense Policy Js Insert All Pages Except

A [`js_insert_all_pages_except`](#client-side-defense-policy-js-insert-all-pages-except) block (within [`client_side_defense.policy`](#client-side-defense-policy)) supports the following:

&#x2022; [`exclude_list`](#exclude-list) - Optional Block<br>Exclude Pages. Optional JavaScript insertions exclude list of domain and path matchers<br>See [Exclude List](#client-side-defense-policy-js-insert-all-pages-except-exclude-list) below.

#### Client Side Defense Policy Js Insert All Pages Except Exclude List

An [`exclude_list`](#client-side-defense-policy-js-insert-all-pages-except-exclude-list) block (within [`client_side_defense.policy.js_insert_all_pages_except`](#client-side-defense-policy-js-insert-all-pages-except)) supports the following:

&#x2022; [`any_domain`](#any-domain) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`domain`](#domain) - Optional Block<br>Domains. Domains names<br>See [Domain](#client-side-defense-policy-js-insert-all-pages-except-exclude-list-domain) below.

&#x2022; [`metadata`](#metadata) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#client-side-defense-policy-js-insert-all-pages-except-exclude-list-metadata) below.

&#x2022; [`path`](#path) - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#client-side-defense-policy-js-insert-all-pages-except-exclude-list-path) below.

#### Client Side Defense Policy Js Insert All Pages Except Exclude List Domain

A [`domain`](#client-side-defense-policy-js-insert-all-pages-except-exclude-list-domain) block (within [`client_side_defense.policy.js_insert_all_pages_except.exclude_list`](#client-side-defense-policy-js-insert-all-pages-except-exclude-list)) supports the following:

&#x2022; [`exact_value`](#exact-value) - Optional String<br>Exact Value. Exact domain name

&#x2022; [`regex_value`](#regex-value) - Optional String<br>Regex Values of Domains. Regular Expression value for the domain name

&#x2022; [`suffix_value`](#suffix-value) - Optional String<br>Suffix Value. Suffix of domain name e.g 'xyz.com' will match '*.xyz.com' and 'xyz.com'

#### Client Side Defense Policy Js Insert All Pages Except Exclude List Metadata

A [`metadata`](#client-side-defense-policy-js-insert-all-pages-except-exclude-list-metadata) block (within [`client_side_defense.policy.js_insert_all_pages_except.exclude_list`](#client-side-defense-policy-js-insert-all-pages-except-exclude-list)) supports the following:

&#x2022; [`description`](#description) - Optional String<br>Description. Human readable description

&#x2022; [`name`](#name) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### Client Side Defense Policy Js Insert All Pages Except Exclude List Path

A [`path`](#client-side-defense-policy-js-insert-all-pages-except-exclude-list-path) block (within [`client_side_defense.policy.js_insert_all_pages_except.exclude_list`](#client-side-defense-policy-js-insert-all-pages-except-exclude-list)) supports the following:

&#x2022; [`path`](#path) - Optional String<br>Exact. Exact path value to match

&#x2022; [`prefix`](#prefix) - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

&#x2022; [`regex`](#regex) - Optional String<br>Regex. Regular expression of path match (e.g. the value .* will match on all paths)

#### Client Side Defense Policy Js Insertion Rules

A [`js_insertion_rules`](#client-side-defense-policy-js-insertion-rules) block (within [`client_side_defense.policy`](#client-side-defense-policy)) supports the following:

&#x2022; [`exclude_list`](#exclude-list) - Optional Block<br>Exclude Paths. Optional JavaScript insertions exclude list of domain and path matchers<br>See [Exclude List](#client-side-defense-policy-js-insertion-rules-exclude-list) below.

&#x2022; [`rules`](#rules) - Optional Block<br>JavaScript Insertions. Required list of pages to insert Client-Side Defense client JavaScript<br>See [Rules](#client-side-defense-policy-js-insertion-rules-rules) below.

#### Client Side Defense Policy Js Insertion Rules Exclude List

An [`exclude_list`](#client-side-defense-policy-js-insertion-rules-exclude-list) block (within [`client_side_defense.policy.js_insertion_rules`](#client-side-defense-policy-js-insertion-rules)) supports the following:

&#x2022; [`any_domain`](#any-domain) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`domain`](#domain) - Optional Block<br>Domains. Domains names<br>See [Domain](#client-side-defense-policy-js-insertion-rules-exclude-list-domain) below.

&#x2022; [`metadata`](#metadata) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#client-side-defense-policy-js-insertion-rules-exclude-list-metadata) below.

&#x2022; [`path`](#path) - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#client-side-defense-policy-js-insertion-rules-exclude-list-path) below.

#### Client Side Defense Policy Js Insertion Rules Exclude List Domain

A [`domain`](#client-side-defense-policy-js-insertion-rules-exclude-list-domain) block (within [`client_side_defense.policy.js_insertion_rules.exclude_list`](#client-side-defense-policy-js-insertion-rules-exclude-list)) supports the following:

&#x2022; [`exact_value`](#exact-value) - Optional String<br>Exact Value. Exact domain name

&#x2022; [`regex_value`](#regex-value) - Optional String<br>Regex Values of Domains. Regular Expression value for the domain name

&#x2022; [`suffix_value`](#suffix-value) - Optional String<br>Suffix Value. Suffix of domain name e.g 'xyz.com' will match '*.xyz.com' and 'xyz.com'

#### Client Side Defense Policy Js Insertion Rules Exclude List Metadata

A [`metadata`](#client-side-defense-policy-js-insertion-rules-exclude-list-metadata) block (within [`client_side_defense.policy.js_insertion_rules.exclude_list`](#client-side-defense-policy-js-insertion-rules-exclude-list)) supports the following:

&#x2022; [`description`](#description) - Optional String<br>Description. Human readable description

&#x2022; [`name`](#name) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### Client Side Defense Policy Js Insertion Rules Exclude List Path

A [`path`](#client-side-defense-policy-js-insertion-rules-exclude-list-path) block (within [`client_side_defense.policy.js_insertion_rules.exclude_list`](#client-side-defense-policy-js-insertion-rules-exclude-list)) supports the following:

&#x2022; [`path`](#path) - Optional String<br>Exact. Exact path value to match

&#x2022; [`prefix`](#prefix) - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

&#x2022; [`regex`](#regex) - Optional String<br>Regex. Regular expression of path match (e.g. the value .* will match on all paths)

#### Client Side Defense Policy Js Insertion Rules Rules

A [`rules`](#client-side-defense-policy-js-insertion-rules-rules) block (within [`client_side_defense.policy.js_insertion_rules`](#client-side-defense-policy-js-insertion-rules)) supports the following:

&#x2022; [`any_domain`](#any-domain) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`domain`](#domain) - Optional Block<br>Domains. Domains names<br>See [Domain](#client-side-defense-policy-js-insertion-rules-rules-domain) below.

&#x2022; [`metadata`](#metadata) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#client-side-defense-policy-js-insertion-rules-rules-metadata) below.

&#x2022; [`path`](#path) - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#client-side-defense-policy-js-insertion-rules-rules-path) below.

#### Client Side Defense Policy Js Insertion Rules Rules Domain

A [`domain`](#client-side-defense-policy-js-insertion-rules-rules-domain) block (within [`client_side_defense.policy.js_insertion_rules.rules`](#client-side-defense-policy-js-insertion-rules-rules)) supports the following:

&#x2022; [`exact_value`](#exact-value) - Optional String<br>Exact Value. Exact domain name

&#x2022; [`regex_value`](#regex-value) - Optional String<br>Regex Values of Domains. Regular Expression value for the domain name

&#x2022; [`suffix_value`](#suffix-value) - Optional String<br>Suffix Value. Suffix of domain name e.g 'xyz.com' will match '*.xyz.com' and 'xyz.com'

#### Client Side Defense Policy Js Insertion Rules Rules Metadata

A [`metadata`](#client-side-defense-policy-js-insertion-rules-rules-metadata) block (within [`client_side_defense.policy.js_insertion_rules.rules`](#client-side-defense-policy-js-insertion-rules-rules)) supports the following:

&#x2022; [`description`](#description) - Optional String<br>Description. Human readable description

&#x2022; [`name`](#name) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### Client Side Defense Policy Js Insertion Rules Rules Path

A [`path`](#client-side-defense-policy-js-insertion-rules-rules-path) block (within [`client_side_defense.policy.js_insertion_rules.rules`](#client-side-defense-policy-js-insertion-rules-rules)) supports the following:

&#x2022; [`path`](#path) - Optional String<br>Exact. Exact path value to match

&#x2022; [`prefix`](#prefix) - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

&#x2022; [`regex`](#regex) - Optional String<br>Regex. Regular expression of path match (e.g. the value .* will match on all paths)

#### Cookie Stickiness

A [`cookie_stickiness`](#cookie-stickiness) block supports the following:

&#x2022; [`add_httponly`](#add-httponly) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`add_secure`](#add-secure) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`ignore_httponly`](#ignore-httponly) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`ignore_samesite`](#ignore-samesite) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`ignore_secure`](#ignore-secure) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`name`](#name) - Optional String<br>Name. The name of the cookie that will be used to obtain the hash key. If the cookie is not present and TTL below is not set, no hash will be produced

&#x2022; [`path`](#path) - Optional String<br>Path. The name of the path for the cookie. If no path is specified here, no path will be set for the cookie

&#x2022; [`samesite_lax`](#samesite-lax) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`samesite_none`](#samesite-none) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`samesite_strict`](#samesite-strict) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`ttl`](#ttl) - Optional Number<br>TTL. If specified, a cookie with the TTL will be generated if the cookie is not present. If the TTL is present and zero, the generated cookie will be a session cookie. TTL value is in milliseconds

#### CORS Policy

A [`cors_policy`](#cors-policy) block supports the following:

&#x2022; [`allow_credentials`](#allow-credentials) - Optional Bool<br>Allow Credentials. Specifies whether the resource allows credentials

&#x2022; [`allow_headers`](#allow-headers) - Optional String<br>Allow Headers. Specifies the content for the access-control-allow-headers header

&#x2022; [`allow_methods`](#allow-methods) - Optional String<br>Allow Methods. Specifies the content for the access-control-allow-methods header

&#x2022; [`allow_origin`](#allow-origin) - Optional List<br>Allow Origin. Specifies the origins that will be allowed to do CORS requests. An origin is allowed if either allow_origin or allow_origin_regex match

&#x2022; [`allow_origin_regex`](#allow-origin-regex) - Optional List<br>Allow Origin Regex. Specifies regex patterns that match allowed origins. An origin is allowed if either allow_origin or allow_origin_regex match

&#x2022; [`disabled`](#disabled) - Optional Bool<br>Disabled. Disable the CorsPolicy for a particular route. This is useful when virtual-host has CorsPolicy, but we need to disable it on a specific route. The value of this field is ignored for virtual-host

&#x2022; [`expose_headers`](#expose-headers) - Optional String<br>Expose Headers. Specifies the content for the access-control-expose-headers header

&#x2022; [`maximum_age`](#maximum-age) - Optional Number<br>Maximum Age. Specifies the content for the access-control-max-age header in seconds. This indicates the maximum number of seconds the results can be cached A value of -1 will disable caching. Maximum permitted value is 86400 seconds (24 hours)

#### CSRF Policy

A [`csrf_policy`](#csrf-policy) block supports the following:

&#x2022; [`all_load_balancer_domains`](#all-load-balancer-domains) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`custom_domain_list`](#custom-domain-list) - Optional Block<br>Domain name list. List of domain names used for Host header matching<br>See [Custom Domain List](#csrf-policy-custom-domain-list) below.

&#x2022; [`disabled`](#disabled) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### CSRF Policy Custom Domain List

A [`custom_domain_list`](#csrf-policy-custom-domain-list) block (within [`csrf_policy`](#csrf-policy)) supports the following:

&#x2022; [`domains`](#domains) - Optional List<br>Domain names. A list of domain names that will be matched to loadbalancer. These domains are not used for SNI match. Wildcard names are supported in the suffix or prefix form

#### Data Guard Rules

A [`data_guard_rules`](#data-guard-rules) block supports the following:

&#x2022; [`any_domain`](#any-domain) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`apply_data_guard`](#apply-data-guard) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`exact_value`](#exact-value) - Optional String<br>Exact Value. Exact domain name

&#x2022; [`metadata`](#metadata) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#data-guard-rules-metadata) below.

&#x2022; [`path`](#path) - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#data-guard-rules-path) below.

&#x2022; [`skip_data_guard`](#skip-data-guard) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`suffix_value`](#suffix-value) - Optional String<br>Suffix Value. Suffix of domain name e.g 'xyz.com' will match '*.xyz.com' and 'xyz.com'

#### Data Guard Rules Metadata

A [`metadata`](#data-guard-rules-metadata) block (within [`data_guard_rules`](#data-guard-rules)) supports the following:

&#x2022; [`description`](#description) - Optional String<br>Description. Human readable description

&#x2022; [`name`](#name) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### Data Guard Rules Path

A [`path`](#data-guard-rules-path) block (within [`data_guard_rules`](#data-guard-rules)) supports the following:

&#x2022; [`path`](#path) - Optional String<br>Exact. Exact path value to match

&#x2022; [`prefix`](#prefix) - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

&#x2022; [`regex`](#regex) - Optional String<br>Regex. Regular expression of path match (e.g. the value .* will match on all paths)

#### DDOS Mitigation Rules

A [`ddos_mitigation_rules`](#ddos-mitigation-rules) block supports the following:

&#x2022; [`block`](#block) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`ddos_client_source`](#ddos-client-source) - Optional Block<br>DDOS Client Source Choice. DDOS Mitigation sources to be blocked<br>See [DDOS Client Source](#ddos-mitigation-rules-ddos-client-source) below.

&#x2022; [`expiration_timestamp`](#expiration-timestamp) - Optional String<br>Expiration Timestamp. The expiration_timestamp is the RFC 3339 format timestamp at which the containing rule is considered to be logically expired. The rule continues to exist in the configuration but is not applied anymore

&#x2022; [`ip_prefix_list`](#ip-prefix-list) - Optional Block<br>IP Prefix Match List. List of IP Prefix strings to match against<br>See [IP Prefix List](#ddos-mitigation-rules-ip-prefix-list) below.

&#x2022; [`metadata`](#metadata) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#ddos-mitigation-rules-metadata) below.

#### DDOS Mitigation Rules DDOS Client Source

A [`ddos_client_source`](#ddos-mitigation-rules-ddos-client-source) block (within [`ddos_mitigation_rules`](#ddos-mitigation-rules)) supports the following:

&#x2022; [`asn_list`](#asn-list) - Optional Block<br>ASN Match List. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer<br>See [Asn List](#ddos-mitigation-rules-ddos-client-source-asn-list) below.

&#x2022; [`country_list`](#country-list) - Optional List  Defaults to `COUNTRY_NONE`<br>Possible values are `COUNTRY_NONE`, `COUNTRY_AD`, `COUNTRY_AE`, `COUNTRY_AF`, `COUNTRY_AG`, `COUNTRY_AI`, `COUNTRY_AL`, `COUNTRY_AM`, `COUNTRY_AN`, `COUNTRY_AO`, `COUNTRY_AQ`, `COUNTRY_AR`, `COUNTRY_AS`, `COUNTRY_AT`, `COUNTRY_AU`, `COUNTRY_AW`, `COUNTRY_AX`, `COUNTRY_AZ`, `COUNTRY_BA`, `COUNTRY_BB`, `COUNTRY_BD`, `COUNTRY_BE`, `COUNTRY_BF`, `COUNTRY_BG`, `COUNTRY_BH`, `COUNTRY_BI`, `COUNTRY_BJ`, `COUNTRY_BL`, `COUNTRY_BM`, `COUNTRY_BN`, `COUNTRY_BO`, `COUNTRY_BQ`, `COUNTRY_BR`, `COUNTRY_BS`, `COUNTRY_BT`, `COUNTRY_BV`, `COUNTRY_BW`, `COUNTRY_BY`, `COUNTRY_BZ`, `COUNTRY_CA`, `COUNTRY_CC`, `COUNTRY_CD`, `COUNTRY_CF`, `COUNTRY_CG`, `COUNTRY_CH`, `COUNTRY_CI`, `COUNTRY_CK`, `COUNTRY_CL`, `COUNTRY_CM`, `COUNTRY_CN`, `COUNTRY_CO`, `COUNTRY_CR`, `COUNTRY_CS`, `COUNTRY_CU`, `COUNTRY_CV`, `COUNTRY_CW`, `COUNTRY_CX`, `COUNTRY_CY`, `COUNTRY_CZ`, `COUNTRY_DE`, `COUNTRY_DJ`, `COUNTRY_DK`, `COUNTRY_DM`, `COUNTRY_DO`, `COUNTRY_DZ`, `COUNTRY_EC`, `COUNTRY_EE`, `COUNTRY_EG`, `COUNTRY_EH`, `COUNTRY_ER`, `COUNTRY_ES`, `COUNTRY_ET`, `COUNTRY_FI`, `COUNTRY_FJ`, `COUNTRY_FK`, `COUNTRY_FM`, `COUNTRY_FO`, `COUNTRY_FR`, `COUNTRY_GA`, `COUNTRY_GB`, `COUNTRY_GD`, `COUNTRY_GE`, `COUNTRY_GF`, `COUNTRY_GG`, `COUNTRY_GH`, `COUNTRY_GI`, `COUNTRY_GL`, `COUNTRY_GM`, `COUNTRY_GN`, `COUNTRY_GP`, `COUNTRY_GQ`, `COUNTRY_GR`, `COUNTRY_GS`, `COUNTRY_GT`, `COUNTRY_GU`, `COUNTRY_GW`, `COUNTRY_GY`, `COUNTRY_HK`, `COUNTRY_HM`, `COUNTRY_HN`, `COUNTRY_HR`, `COUNTRY_HT`, `COUNTRY_HU`, `COUNTRY_ID`, `COUNTRY_IE`, `COUNTRY_IL`, `COUNTRY_IM`, `COUNTRY_IN`, `COUNTRY_IO`, `COUNTRY_IQ`, `COUNTRY_IR`, `COUNTRY_IS`, `COUNTRY_IT`, `COUNTRY_JE`, `COUNTRY_JM`, `COUNTRY_JO`, `COUNTRY_JP`, `COUNTRY_KE`, `COUNTRY_KG`, `COUNTRY_KH`, `COUNTRY_KI`, `COUNTRY_KM`, `COUNTRY_KN`, `COUNTRY_KP`, `COUNTRY_KR`, `COUNTRY_KW`, `COUNTRY_KY`, `COUNTRY_KZ`, `COUNTRY_LA`, `COUNTRY_LB`, `COUNTRY_LC`, `COUNTRY_LI`, `COUNTRY_LK`, `COUNTRY_LR`, `COUNTRY_LS`, `COUNTRY_LT`, `COUNTRY_LU`, `COUNTRY_LV`, `COUNTRY_LY`, `COUNTRY_MA`, `COUNTRY_MC`, `COUNTRY_MD`, `COUNTRY_ME`, `COUNTRY_MF`, `COUNTRY_MG`, `COUNTRY_MH`, `COUNTRY_MK`, `COUNTRY_ML`, `COUNTRY_MM`, `COUNTRY_MN`, `COUNTRY_MO`, `COUNTRY_MP`, `COUNTRY_MQ`, `COUNTRY_MR`, `COUNTRY_MS`, `COUNTRY_MT`, `COUNTRY_MU`, `COUNTRY_MV`, `COUNTRY_MW`, `COUNTRY_MX`, `COUNTRY_MY`, `COUNTRY_MZ`, `COUNTRY_NA`, `COUNTRY_NC`, `COUNTRY_NE`, `COUNTRY_NF`, `COUNTRY_NG`, `COUNTRY_NI`, `COUNTRY_NL`, `COUNTRY_NO`, `COUNTRY_NP`, `COUNTRY_NR`, `COUNTRY_NU`, `COUNTRY_NZ`, `COUNTRY_OM`, `COUNTRY_PA`, `COUNTRY_PE`, `COUNTRY_PF`, `COUNTRY_PG`, `COUNTRY_PH`, `COUNTRY_PK`, `COUNTRY_PL`, `COUNTRY_PM`, `COUNTRY_PN`, `COUNTRY_PR`, `COUNTRY_PS`, `COUNTRY_PT`, `COUNTRY_PW`, `COUNTRY_PY`, `COUNTRY_QA`, `COUNTRY_RE`, `COUNTRY_RO`, `COUNTRY_RS`, `COUNTRY_RU`, `COUNTRY_RW`, `COUNTRY_SA`, `COUNTRY_SB`, `COUNTRY_SC`, `COUNTRY_SD`, `COUNTRY_SE`, `COUNTRY_SG`, `COUNTRY_SH`, `COUNTRY_SI`, `COUNTRY_SJ`, `COUNTRY_SK`, `COUNTRY_SL`, `COUNTRY_SM`, `COUNTRY_SN`, `COUNTRY_SO`, `COUNTRY_SR`, `COUNTRY_SS`, `COUNTRY_ST`, `COUNTRY_SV`, `COUNTRY_SX`, `COUNTRY_SY`, `COUNTRY_SZ`, `COUNTRY_TC`, `COUNTRY_TD`, `COUNTRY_TF`, `COUNTRY_TG`, `COUNTRY_TH`, `COUNTRY_TJ`, `COUNTRY_TK`, `COUNTRY_TL`, `COUNTRY_TM`, `COUNTRY_TN`, `COUNTRY_TO`, `COUNTRY_TR`, `COUNTRY_TT`, `COUNTRY_TV`, `COUNTRY_TW`, `COUNTRY_TZ`, `COUNTRY_UA`, `COUNTRY_UG`, `COUNTRY_UM`, `COUNTRY_US`, `COUNTRY_UY`, `COUNTRY_UZ`, `COUNTRY_VA`, `COUNTRY_VC`, `COUNTRY_VE`, `COUNTRY_VG`, `COUNTRY_VI`, `COUNTRY_VN`, `COUNTRY_VU`, `COUNTRY_WF`, `COUNTRY_WS`, `COUNTRY_XK`, `COUNTRY_XT`, `COUNTRY_YE`, `COUNTRY_YT`, `COUNTRY_ZA`, `COUNTRY_ZM`, `COUNTRY_ZW`<br>Country List. Sources that are located in one of the countries in the given list

&#x2022; [`ja4_tls_fingerprint_matcher`](#ja4-tls-fingerprint-matcher) - Optional Block<br>JA4 TLS Fingerprint Matcher. An extended version of JA3 that includes additional fields for more comprehensive fingerprinting of SSL/TLS clients and potentially has a different structure and length<br>See [Ja4 TLS Fingerprint Matcher](#ddos-mitigation-rules-ddos-client-source-ja4-tls-fingerprint-matcher) below.

&#x2022; [`tls_fingerprint_matcher`](#tls-fingerprint-matcher) - Optional Block<br>TLS Fingerprint Matcher. A TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of supported positve match criteria includes a list of known classes of TLS fingerprints and a list of exact values. The match is considered successful if either of these positive criteria are satisfied and the input fingerprint is not one of the excluded values<br>See [TLS Fingerprint Matcher](#ddos-mitigation-rules-ddos-client-source-tls-fingerprint-matcher) below.

#### DDOS Mitigation Rules DDOS Client Source Asn List

An [`asn_list`](#ddos-mitigation-rules-ddos-client-source-asn-list) block (within [`ddos_mitigation_rules.ddos_client_source`](#ddos-mitigation-rules-ddos-client-source)) supports the following:

&#x2022; [`as_numbers`](#as-numbers) - Optional List<br>AS Numbers. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer

#### DDOS Mitigation Rules DDOS Client Source Ja4 TLS Fingerprint Matcher

A [`ja4_tls_fingerprint_matcher`](#ddos-mitigation-rules-ddos-client-source-ja4-tls-fingerprint-matcher) block (within [`ddos_mitigation_rules.ddos_client_source`](#ddos-mitigation-rules-ddos-client-source)) supports the following:

&#x2022; [`exact_values`](#exact-values) - Optional List<br>Exact Values. A list of exact JA4 TLS fingerprint to match the input JA4 TLS fingerprint against

#### DDOS Mitigation Rules DDOS Client Source TLS Fingerprint Matcher

A [`tls_fingerprint_matcher`](#ddos-mitigation-rules-ddos-client-source-tls-fingerprint-matcher) block (within [`ddos_mitigation_rules.ddos_client_source`](#ddos-mitigation-rules-ddos-client-source)) supports the following:

&#x2022; [`classes`](#classes) - Optional List  Defaults to `TLS_FINGERPRINT_NONE`<br>Possible values are `TLS_FINGERPRINT_NONE`, `ANY_MALICIOUS_FINGERPRINT`, `ADWARE`, `ADWIND`, `DRIDEX`, `GOOTKIT`, `GOZI`, `JBIFROST`, `QUAKBOT`, `RANSOMWARE`, `TROLDESH`, `TOFSEE`, `TORRENTLOCKER`, `TRICKBOT`<br>TLS fingerprint classes. A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against

&#x2022; [`exact_values`](#exact-values) - Optional List<br>Exact Values. A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against

&#x2022; [`excluded_values`](#excluded-values) - Optional List<br>Excluded Values. A list of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can be used to skip known false positives when using one or more known TLS fingerprint classes in the enclosing matcher

#### DDOS Mitigation Rules IP Prefix List

An [`ip_prefix_list`](#ddos-mitigation-rules-ip-prefix-list) block (within [`ddos_mitigation_rules`](#ddos-mitigation-rules)) supports the following:

&#x2022; [`invert_match`](#invert-match) - Optional Bool<br>Invert Match Result. Invert the match result

&#x2022; [`ip_prefixes`](#ip-prefixes) - Optional List<br>IPv4 Prefix List. List of IPv4 prefix strings

#### DDOS Mitigation Rules Metadata

A [`metadata`](#ddos-mitigation-rules-metadata) block (within [`ddos_mitigation_rules`](#ddos-mitigation-rules)) supports the following:

&#x2022; [`description`](#description) - Optional String<br>Description. Human readable description

&#x2022; [`name`](#name) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### Default Pool

A [`default_pool`](#default-pool) block supports the following:

&#x2022; [`advanced_options`](#advanced-options) - Optional Block<br>Origin Pool Advanced Options. Configure Advanced options for origin pool<br>See [Advanced Options](#default-pool-advanced-options) below.

&#x2022; [`automatic_port`](#automatic-port) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`endpoint_selection`](#endpoint-selection) - Optional String  Defaults to `DISTRIBUTED`<br>Possible values are `DISTRIBUTED`, `LOCAL_ONLY`, `LOCAL_PREFERRED`<br>Endpoint Selection Policy. Policy for selection of endpoints from local site/remote site/both Consider both remote and local endpoints for load balancing LOCAL_ONLY: Consider only local endpoints for load balancing Enable this policy to load balance ONLY among locally discovered endpoints Prefer the local endpoints for load balancing. If local endpoints are not present remote endpoints will be considered

&#x2022; [`health_check_port`](#health-check-port) - Optional Number<br>Health check port. Port used for performing health check

&#x2022; [`healthcheck`](#healthcheck) - Optional Block<br>Health Check object. Reference to healthcheck configuration objects<br>See [Healthcheck](#default-pool-healthcheck) below.

&#x2022; [`lb_port`](#lb-port) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`loadbalancer_algorithm`](#loadbalancer-algorithm) - Optional String  Defaults to `ROUND_ROBIN`<br>Possible values are `ROUND_ROBIN`, `LEAST_REQUEST`, `RING_HASH`, `RANDOM`, `LB_OVERRIDE`<br>Load Balancer Algorithm. Different load balancing algorithms supported When a connection to a endpoint in an upstream cluster is required, the load balancer uses loadbalancer_algorithm to determine which host is selected. - ROUND_ROBIN: ROUND_ROBIN Policy in which each healthy/available upstream endpoint is selected in round robin order. - LEAST_REQUEST: LEAST_REQUEST Policy in which loadbalancer picks the upstream endpoint which has the fewest active requests - RING_HASH: RING_HASH Policy implements consistent hashing to upstream endpoints using ring hash of endpoint names Hash of the incoming request is calculated using request hash policy. The ring/modulo hash load balancer implements consistent hashing to upstream hosts. The algorithm is based on mapping all hosts onto a circle such that the addition or removal of a host from the host set changes only affect 1/N requests. This technique is also commonly known as “ketama” hashing. A consistent hashing load balancer is only effective when protocol routing is used that specifies a value to hash on. The minimum ring size governs the replication factor for each host in the ring. For example, if the minimum ring size is 1024 and there are 16 hosts, each host will be replicated 64 times. - RANDOM: RANDOM Policy in which each available upstream endpoint is selected in random order. The random load balancer selects a random healthy host. The random load balancer generally performs better than round robin if no health checking policy is configured. Random selection avoids bias towards the host in the set that comes after a failed host. - LB_OVERRIDE: Load Balancer Override Hash policy is taken from from the load balancer which is using this origin pool

&#x2022; [`no_tls`](#no-tls) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`origin_servers`](#origin-servers) - Optional Block<br>Origin Servers. List of origin servers in this pool<br>See [Origin Servers](#default-pool-origin-servers) below.

&#x2022; [`port`](#port) - Optional Number<br>Port. Endpoint service is available on this port

&#x2022; [`same_as_endpoint_port`](#same-as-endpoint-port) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`upstream_conn_pool_reuse_type`](#upstream-conn-pool-reuse-type) - Optional Block<br>Select upstream connection pool reuse state. Select upstream connection pool reuse state for every downstream connection. This configuration choice is for HTTP(S) LB only<br>See [Upstream Conn Pool Reuse Type](#default-pool-upstream-conn-pool-reuse-type) below.

&#x2022; [`use_tls`](#use-tls) - Optional Block<br>TLS Parameters for Origin Servers. Upstream TLS Parameters<br>See [Use TLS](#default-pool-use-tls) below.

&#x2022; [`view_internal`](#view-internal) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [View Internal](#default-pool-view-internal) below.

#### Default Pool Advanced Options

An [`advanced_options`](#default-pool-advanced-options) block (within [`default_pool`](#default-pool)) supports the following:

&#x2022; [`auto_http_config`](#auto-http-config) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`circuit_breaker`](#circuit-breaker) - Optional Block<br>Circuit Breaker. CircuitBreaker provides a mechanism for watching failures in upstream connections or requests and if the failures reach a certain threshold, automatically fail subsequent requests which allows to apply back pressure on downstream quickly<br>See [Circuit Breaker](#default-pool-advanced-options-circuit-breaker) below.

&#x2022; [`connection_timeout`](#connection-timeout) - Optional Number  Defaults to `2`  Specified in milliseconds<br>Connection Timeout. The timeout for new network connections to endpoints in the cluster.  The seconds

&#x2022; [`default_circuit_breaker`](#default-circuit-breaker) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`disable_circuit_breaker`](#disable-circuit-breaker) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`disable_lb_source_ip_persistance`](#disable-lb-source-ip-persistance) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`disable_outlier_detection`](#disable-outlier-detection) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`disable_proxy_protocol`](#disable-proxy-protocol) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`disable_subsets`](#disable-subsets) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`enable_lb_source_ip_persistance`](#enable-lb-source-ip-persistance) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`enable_subsets`](#enable-subsets) - Optional Block<br>Origin Pool Subset Options. Configure subset options for origin pool<br>See [Enable Subsets](#default-pool-advanced-options-enable-subsets) below.

&#x2022; [`http1_config`](#http1-config) - Optional Block<br>HTTP/1.1 Protocol Options. HTTP/1.1 Protocol options for upstream connections<br>See [Http1 Config](#default-pool-advanced-options-http1-config) below.

&#x2022; [`http2_options`](#http2-options) - Optional Block<br>Http2 Protocol Options. Http2 Protocol options for upstream connections<br>See [Http2 Options](#default-pool-advanced-options-http2-options) below.

&#x2022; [`http_idle_timeout`](#http-idle-timeout) - Optional Number  Defaults to `5`  Specified in milliseconds<br>HTTP Idle Timeout. The idle timeout for upstream connection pool connections. The idle timeout is defined as the period in which there are no active requests. When the idle timeout is reached the connection will be closed. Note that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive.  The minutes

&#x2022; [`no_panic_threshold`](#no-panic-threshold) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`outlier_detection`](#outlier-detection) - Optional Block<br>Outlier Detection. Outlier detection and ejection is the process of dynamically determining whether some number of hosts in an upstream cluster are performing unlike the others and removing them from the healthy load balancing set. Outlier detection is a form of passive health checking. Algorithm 1. A endpoint is determined to be an outlier (based on configured number of consecutive_5xx or consecutive_gateway_failures) . 2. If no endpoints have been ejected, loadbalancer will eject the host immediately. Otherwise, it checks to make sure the number of ejected hosts is below the allowed threshold (specified via max_ejection_percent setting). If the number of ejected hosts is above the threshold, the host is not ejected. 3. The endpoint is ejected for some number of milliseconds. Ejection means that the endpoint is marked unhealthy and will not be used during load balancing. The number of milliseconds is equal to the base_ejection_time value multiplied by the number of times the host has been ejected. 4. An ejected endpoint will automatically be brought back into service after the ejection time has been satisfied<br>See [Outlier Detection](#default-pool-advanced-options-outlier-detection) below.

&#x2022; [`panic_threshold`](#panic-threshold) - Optional Number<br>Panic threshold. x-example:'25' Configure a threshold (percentage of unhealthy endpoints) below which all endpoints will be considered for load balancing ignoring its health status

&#x2022; [`proxy_protocol_v1`](#proxy-protocol-v1) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`proxy_protocol_v2`](#proxy-protocol-v2) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### Default Pool Advanced Options Circuit Breaker

A [`circuit_breaker`](#default-pool-advanced-options-circuit-breaker) block (within [`default_pool.advanced_options`](#default-pool-advanced-options)) supports the following:

&#x2022; [`connection_limit`](#connection-limit) - Optional Number<br>Connection Limit. The maximum number of connections that loadbalancer will establish to all hosts in an upstream cluster. In practice this is only applicable to TCP and HTTP/1.1 clusters since HTTP/2 uses a single connection to each host. Remove endpoint out of load balancing decision, if number of connections reach connection limit

&#x2022; [`max_requests`](#max-requests) - Optional Number<br>Maximum Request Count. The maximum number of requests that can be outstanding to all hosts in a cluster at any given time. In practice this is applicable to HTTP/2 clusters since HTTP/1.1 clusters are governed by the maximum connections (connection_limit). Remove endpoint out of load balancing decision, if requests exceed this count

&#x2022; [`pending_requests`](#pending-requests) - Optional Number<br>Pending Requests. The maximum number of requests that will be queued while waiting for a ready connection pool connection. Since HTTP/2 requests are sent over a single connection, this circuit breaker only comes into play as the initial connection is created, as requests will be multiplexed immediately afterwards. For HTTP/1.1, requests are added to the list of pending requests whenever there aren’t enough upstream connections available to immediately dispatch the request, so this circuit breaker will remain in play for the lifetime of the process. Remove endpoint out of load balancing decision, if pending request reach pending_request

&#x2022; [`priority`](#priority) - Optional String  Defaults to `DEFAULT`<br>Possible values are `DEFAULT`, `HIGH`<br>Routing Priority. Priority routing for each request. Different connection pools are used based on the priority selected for the request. Also, circuit-breaker configuration at destination cluster is chosen based on selected priority. Default routing mechanism High-Priority routing mechanism

&#x2022; [`retries`](#retries) - Optional Number<br>Retry Count. The maximum number of retries that can be outstanding to all hosts in a cluster at any given time. Remove endpoint out of load balancing decision, if retries for request exceed this count

#### Default Pool Advanced Options Enable Subsets

An [`enable_subsets`](#default-pool-advanced-options-enable-subsets) block (within [`default_pool.advanced_options`](#default-pool-advanced-options)) supports the following:

&#x2022; [`any_endpoint`](#any-endpoint) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`default_subset`](#default-subset) - Optional Block<br>Origin Pool Default Subset. Default Subset definition<br>See [Default Subset](#default-pool-advanced-options-enable-subsets-default-subset) below.

&#x2022; [`endpoint_subsets`](#endpoint-subsets) - Optional Block<br>Origin Server Subsets Classes. List of subset class. Subsets class is defined using list of keys. Every unique combination of values of these keys form a subset withing the class<br>See [Endpoint Subsets](#default-pool-advanced-options-enable-subsets-endpoint-subsets) below.

&#x2022; [`fail_request`](#fail-request) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### Default Pool Advanced Options Enable Subsets Default Subset

A [`default_subset`](#default-pool-advanced-options-enable-subsets-default-subset) block (within [`default_pool.advanced_options.enable_subsets`](#default-pool-advanced-options-enable-subsets)) supports the following:

&#x2022; [`default_subset`](#default-subset) - Optional Block<br>Default Subset for Origin Pool. List of key-value pairs that define default subset. which gets used when route specifies no metadata or no subset matching the metadata exists

#### Default Pool Advanced Options Enable Subsets Endpoint Subsets

An [`endpoint_subsets`](#default-pool-advanced-options-enable-subsets-endpoint-subsets) block (within [`default_pool.advanced_options.enable_subsets`](#default-pool-advanced-options-enable-subsets)) supports the following:

&#x2022; [`keys`](#keys) - Optional List<br>Keys. List of keys that define a cluster subset class

#### Default Pool Advanced Options Http1 Config

A [`http1_config`](#default-pool-advanced-options-http1-config) block (within [`default_pool.advanced_options`](#default-pool-advanced-options)) supports the following:

&#x2022; [`header_transformation`](#header-transformation) - Optional Block<br>Header Transformation. Header Transformation options for HTTP/1.1 request/response headers<br>See [Header Transformation](#default-pool-advanced-options-http1-config-header-transformation) below.

#### Default Pool Advanced Options Http1 Config Header Transformation

A [`header_transformation`](#default-pool-advanced-options-http1-config-header-transformation) block (within [`default_pool.advanced_options.http1_config`](#default-pool-advanced-options-http1-config)) supports the following:

&#x2022; [`default_header_transformation`](#default-header-transformation) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`legacy_header_transformation`](#legacy-header-transformation) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`preserve_case_header_transformation`](#preserve-case-header-transformation) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`proper_case_header_transformation`](#proper-case-header-transformation) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### Default Pool Advanced Options Http2 Options

A [`http2_options`](#default-pool-advanced-options-http2-options) block (within [`default_pool.advanced_options`](#default-pool-advanced-options)) supports the following:

&#x2022; [`enabled`](#enabled) - Optional Bool<br>HTTP2 Enabled. Enable/disable HTTP2 Protocol for upstream connections

#### Default Pool Advanced Options Outlier Detection

An [`outlier_detection`](#default-pool-advanced-options-outlier-detection) block (within [`default_pool.advanced_options`](#default-pool-advanced-options)) supports the following:

&#x2022; [`base_ejection_time`](#base-ejection-time) - Optional Number  Defaults to `30000ms`  Specified in milliseconds<br>Base Ejection Time. The base time that a host is ejected for. The real time is equal to the base time multiplied by the number of times the host has been ejected. This causes hosts to get ejected for longer periods if they continue to fail

&#x2022; [`consecutive_5xx`](#consecutive-5xx) - Optional Number  Defaults to `5`<br>Consecutive 5xx Count. If an upstream endpoint returns some number of consecutive 5xx, it will be ejected. Note that in this case a 5xx means an actual 5xx respond code, or an event that would cause the HTTP router to return one on the upstream’s behalf(reset, connection failure, etc.) consecutive_5xx indicates the number of consecutive 5xx responses required before a consecutive 5xx ejection occurs

&#x2022; [`consecutive_gateway_failure`](#consecutive-gateway-failure) - Optional Number  Defaults to `5`<br>Consecutive Gateway Failure. If an upstream endpoint returns some number of consecutive “gateway errors” (502, 503 or 504 status code), it will be ejected. Note that this includes events that would cause the HTTP router to return one of these status codes on the upstream’s behalf (reset, connection failure, etc.). consecutive_gateway_failure indicates the number of consecutive gateway failures before a consecutive gateway failure ejection occurs

&#x2022; [`interval`](#interval) - Optional Number  Defaults to `10000ms`  Specified in milliseconds<br>Interval. The time interval between ejection analysis sweeps. This can result in both new ejections as well as endpoints being returned to service

&#x2022; [`max_ejection_percent`](#max-ejection-percent) - Optional Number  Defaults to `10%`<br>Max Ejection Percentage. The maximum % of an upstream cluster that can be ejected due to outlier detection. but will eject at least one host regardless of the value

#### Default Pool Healthcheck

A [`healthcheck`](#default-pool-healthcheck) block (within [`default_pool`](#default-pool)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Default Pool Origin Servers

An [`origin_servers`](#default-pool-origin-servers) block (within [`default_pool`](#default-pool)) supports the following:

&#x2022; [`cbip_service`](#cbip-service) - Optional Block<br>Discovered Classic BIG-IP Service Name. Specify origin server with Classic BIG-IP Service (Virtual Server)<br>See [Cbip Service](#default-pool-origin-servers-cbip-service) below.

&#x2022; [`consul_service`](#consul-service) - Optional Block<br>Consul Service Name on given Sites. Specify origin server with Hashi Corp Consul service name and site information<br>See [Consul Service](#default-pool-origin-servers-consul-service) below.

&#x2022; [`custom_endpoint_object`](#custom-endpoint-object) - Optional Block<br>Custom Endpoint Object for Origin Server. Specify origin server with a reference to endpoint object<br>See [Custom Endpoint Object](#default-pool-origin-servers-custom-endpoint-object) below.

&#x2022; [`k8s_service`](#k8s-service) - Optional Block<br>K8s Service Name on given Sites. Specify origin server with K8s service name and site information<br>See [K8s Service](#default-pool-origin-servers-k8s-service) below.

&#x2022; [`labels`](#labels) - Optional Block<br>Origin Server Labels. Add Labels for this origin server, these labels can be used to form subset

&#x2022; [`private_ip`](#private-ip) - Optional Block<br>IP address on given Sites. Specify origin server with private or public IP address and site information<br>See [Private IP](#default-pool-origin-servers-private-ip) below.

&#x2022; [`private_name`](#private-name) - Optional Block<br>DNS Name on given Sites. Specify origin server with private or public DNS name and site information<br>See [Private Name](#default-pool-origin-servers-private-name) below.

&#x2022; [`public_ip`](#public-ip) - Optional Block<br>Public IP. Specify origin server with public IP address<br>See [Public IP](#default-pool-origin-servers-public-ip) below.

&#x2022; [`public_name`](#public-name) - Optional Block<br>Public DNS Name. Specify origin server with public DNS name<br>See [Public Name](#default-pool-origin-servers-public-name) below.

&#x2022; [`vn_private_ip`](#vn-private-ip) - Optional Block<br>IP address Virtual Network. Specify origin server with IP on Virtual Network<br>See [Vn Private IP](#default-pool-origin-servers-vn-private-ip) below.

&#x2022; [`vn_private_name`](#vn-private-name) - Optional Block<br>DNS Name on Virtual Network. Specify origin server with DNS name on Virtual Network<br>See [Vn Private Name](#default-pool-origin-servers-vn-private-name) below.

#### Default Pool Origin Servers Cbip Service

A [`cbip_service`](#default-pool-origin-servers-cbip-service) block (within [`default_pool.origin_servers`](#default-pool-origin-servers)) supports the following:

&#x2022; [`service_name`](#service-name) - Optional String<br>Service Name. Name of the discovered Classic BIG-IP virtual server to be used as origin

#### Default Pool Origin Servers Consul Service

A [`consul_service`](#default-pool-origin-servers-consul-service) block (within [`default_pool.origin_servers`](#default-pool-origin-servers)) supports the following:

&#x2022; [`inside_network`](#inside-network) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`outside_network`](#outside-network) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`service_name`](#service-name) - Optional String<br>Service Name. Consul service name of this origin server will be listed, including cluster-id. The format is servicename:cluster-id

&#x2022; [`site_locator`](#site-locator) - Optional Block<br>Site or Virtual Site. This message defines a reference to a site or virtual site object<br>See [Site Locator](#default-pool-origin-servers-consul-service-site-locator) below.

&#x2022; [`snat_pool`](#snat-pool) - Optional Block<br>Snat Pool. Snat Pool configuration<br>See [Snat Pool](#default-pool-origin-servers-consul-service-snat-pool) below.

#### Default Pool Origin Servers Consul Service Site Locator

A [`site_locator`](#default-pool-origin-servers-consul-service-site-locator) block (within [`default_pool.origin_servers.consul_service`](#default-pool-origin-servers-consul-service)) supports the following:

&#x2022; [`site`](#site) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Site](#default-pool-origin-servers-consul-service-site-locator-site) below.

&#x2022; [`virtual_site`](#virtual-site) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Virtual Site](#default-pool-origin-servers-consul-service-site-locator-virtual-site) below.

#### Default Pool Origin Servers Consul Service Site Locator Site

A [`site`](#default-pool-origin-servers-consul-service-site-locator-site) block (within [`default_pool.origin_servers.consul_service.site_locator`](#default-pool-origin-servers-consul-service-site-locator)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Default Pool Origin Servers Consul Service Site Locator Virtual Site

A [`virtual_site`](#default-pool-origin-servers-consul-service-site-locator-virtual-site) block (within [`default_pool.origin_servers.consul_service.site_locator`](#default-pool-origin-servers-consul-service-site-locator)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Default Pool Origin Servers Consul Service Snat Pool

A [`snat_pool`](#default-pool-origin-servers-consul-service-snat-pool) block (within [`default_pool.origin_servers.consul_service`](#default-pool-origin-servers-consul-service)) supports the following:

&#x2022; [`no_snat_pool`](#no-snat-pool) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`snat_pool`](#snat-pool) - Optional Block<br>IPv4 Prefix List. x-example: '192.168.20.0/24' List of IPv4 prefixes that represent an endpoint<br>See [Snat Pool](#default-pool-origin-servers-consul-service-snat-pool-snat-pool) below.

#### Default Pool Origin Servers Consul Service Snat Pool Snat Pool

A [`snat_pool`](#default-pool-origin-servers-consul-service-snat-pool-snat-pool) block (within [`default_pool.origin_servers.consul_service.snat_pool`](#default-pool-origin-servers-consul-service-snat-pool)) supports the following:

&#x2022; [`prefixes`](#prefixes) - Optional List<br>IPv4 Prefix List. List of IPv4 prefixes that represent an endpoint

#### Default Pool Origin Servers Custom Endpoint Object

A [`custom_endpoint_object`](#default-pool-origin-servers-custom-endpoint-object) block (within [`default_pool.origin_servers`](#default-pool-origin-servers)) supports the following:

&#x2022; [`endpoint`](#endpoint) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Endpoint](#default-pool-origin-servers-custom-endpoint-object-endpoint) below.

#### Default Pool Origin Servers Custom Endpoint Object Endpoint

An [`endpoint`](#default-pool-origin-servers-custom-endpoint-object-endpoint) block (within [`default_pool.origin_servers.custom_endpoint_object`](#default-pool-origin-servers-custom-endpoint-object)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Default Pool Origin Servers K8s Service

A [`k8s_service`](#default-pool-origin-servers-k8s-service) block (within [`default_pool.origin_servers`](#default-pool-origin-servers)) supports the following:

&#x2022; [`inside_network`](#inside-network) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`outside_network`](#outside-network) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`protocol`](#protocol) - Optional String  Defaults to `PROTOCOL_TCP`<br>Possible values are `PROTOCOL_TCP`, `PROTOCOL_UDP`<br>Protocol Type. Type of protocol - PROTOCOL_TCP: TCP - PROTOCOL_UDP: UDP

&#x2022; [`service_name`](#service-name) - Optional String<br>Service Name. K8s service name of the origin server will be listed, including the namespace and cluster-id. For vK8s services, you need to enter a string with the format servicename.namespace:cluster-id. If the servicename is 'frontend', namespace is 'speedtest' and cluster-id is 'prod', then you will enter 'frontend.speedtest:prod'. Both namespace and cluster-id are optional

&#x2022; [`site_locator`](#site-locator) - Optional Block<br>Site or Virtual Site. This message defines a reference to a site or virtual site object<br>See [Site Locator](#default-pool-origin-servers-k8s-service-site-locator) below.

&#x2022; [`snat_pool`](#snat-pool) - Optional Block<br>Snat Pool. Snat Pool configuration<br>See [Snat Pool](#default-pool-origin-servers-k8s-service-snat-pool) below.

&#x2022; [`vk8s_networks`](#vk8s-networks) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### Default Pool Origin Servers K8s Service Site Locator

A [`site_locator`](#default-pool-origin-servers-k8s-service-site-locator) block (within [`default_pool.origin_servers.k8s_service`](#default-pool-origin-servers-k8s-service)) supports the following:

&#x2022; [`site`](#site) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Site](#default-pool-origin-servers-k8s-service-site-locator-site) below.

&#x2022; [`virtual_site`](#virtual-site) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Virtual Site](#default-pool-origin-servers-k8s-service-site-locator-virtual-site) below.

#### Default Pool Origin Servers K8s Service Site Locator Site

A [`site`](#default-pool-origin-servers-k8s-service-site-locator-site) block (within [`default_pool.origin_servers.k8s_service.site_locator`](#default-pool-origin-servers-k8s-service-site-locator)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Default Pool Origin Servers K8s Service Site Locator Virtual Site

A [`virtual_site`](#default-pool-origin-servers-k8s-service-site-locator-virtual-site) block (within [`default_pool.origin_servers.k8s_service.site_locator`](#default-pool-origin-servers-k8s-service-site-locator)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Default Pool Origin Servers K8s Service Snat Pool

A [`snat_pool`](#default-pool-origin-servers-k8s-service-snat-pool) block (within [`default_pool.origin_servers.k8s_service`](#default-pool-origin-servers-k8s-service)) supports the following:

&#x2022; [`no_snat_pool`](#no-snat-pool) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`snat_pool`](#snat-pool) - Optional Block<br>IPv4 Prefix List. x-example: '192.168.20.0/24' List of IPv4 prefixes that represent an endpoint<br>See [Snat Pool](#default-pool-origin-servers-k8s-service-snat-pool-snat-pool) below.

#### Default Pool Origin Servers K8s Service Snat Pool Snat Pool

A [`snat_pool`](#default-pool-origin-servers-k8s-service-snat-pool-snat-pool) block (within [`default_pool.origin_servers.k8s_service.snat_pool`](#default-pool-origin-servers-k8s-service-snat-pool)) supports the following:

&#x2022; [`prefixes`](#prefixes) - Optional List<br>IPv4 Prefix List. List of IPv4 prefixes that represent an endpoint

#### Default Pool Origin Servers Private IP

A [`private_ip`](#default-pool-origin-servers-private-ip) block (within [`default_pool.origin_servers`](#default-pool-origin-servers)) supports the following:

&#x2022; [`inside_network`](#inside-network) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`ip`](#ip) - Optional String<br>IP. Private IPv4 address

&#x2022; [`outside_network`](#outside-network) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`segment`](#segment) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Segment](#default-pool-origin-servers-private-ip-segment) below.

&#x2022; [`site_locator`](#site-locator) - Optional Block<br>Site or Virtual Site. This message defines a reference to a site or virtual site object<br>See [Site Locator](#default-pool-origin-servers-private-ip-site-locator) below.

&#x2022; [`snat_pool`](#snat-pool) - Optional Block<br>Snat Pool. Snat Pool configuration<br>See [Snat Pool](#default-pool-origin-servers-private-ip-snat-pool) below.

#### Default Pool Origin Servers Private IP Segment

A [`segment`](#default-pool-origin-servers-private-ip-segment) block (within [`default_pool.origin_servers.private_ip`](#default-pool-origin-servers-private-ip)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Default Pool Origin Servers Private IP Site Locator

A [`site_locator`](#default-pool-origin-servers-private-ip-site-locator) block (within [`default_pool.origin_servers.private_ip`](#default-pool-origin-servers-private-ip)) supports the following:

&#x2022; [`site`](#site) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Site](#default-pool-origin-servers-private-ip-site-locator-site) below.

&#x2022; [`virtual_site`](#virtual-site) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Virtual Site](#default-pool-origin-servers-private-ip-site-locator-virtual-site) below.

#### Default Pool Origin Servers Private IP Site Locator Site

A [`site`](#default-pool-origin-servers-private-ip-site-locator-site) block (within [`default_pool.origin_servers.private_ip.site_locator`](#default-pool-origin-servers-private-ip-site-locator)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Default Pool Origin Servers Private IP Site Locator Virtual Site

A [`virtual_site`](#default-pool-origin-servers-private-ip-site-locator-virtual-site) block (within [`default_pool.origin_servers.private_ip.site_locator`](#default-pool-origin-servers-private-ip-site-locator)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Default Pool Origin Servers Private IP Snat Pool

A [`snat_pool`](#default-pool-origin-servers-private-ip-snat-pool) block (within [`default_pool.origin_servers.private_ip`](#default-pool-origin-servers-private-ip)) supports the following:

&#x2022; [`no_snat_pool`](#no-snat-pool) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`snat_pool`](#snat-pool) - Optional Block<br>IPv4 Prefix List. x-example: '192.168.20.0/24' List of IPv4 prefixes that represent an endpoint<br>See [Snat Pool](#default-pool-origin-servers-private-ip-snat-pool-snat-pool) below.

#### Default Pool Origin Servers Private IP Snat Pool Snat Pool

A [`snat_pool`](#default-pool-origin-servers-private-ip-snat-pool-snat-pool) block (within [`default_pool.origin_servers.private_ip.snat_pool`](#default-pool-origin-servers-private-ip-snat-pool)) supports the following:

&#x2022; [`prefixes`](#prefixes) - Optional List<br>IPv4 Prefix List. List of IPv4 prefixes that represent an endpoint

#### Default Pool Origin Servers Private Name

A [`private_name`](#default-pool-origin-servers-private-name) block (within [`default_pool.origin_servers`](#default-pool-origin-servers)) supports the following:

&#x2022; [`dns_name`](#dns-name) - Optional String<br>DNS Name. DNS Name

&#x2022; [`inside_network`](#inside-network) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`outside_network`](#outside-network) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`refresh_interval`](#refresh-interval) - Optional Number<br>DNS Refresh Interval. Interval for DNS refresh in seconds. Max value is 7 days as per `HTTPS://datatracker.ietf.org/doc/HTML/rfc8767`

&#x2022; [`segment`](#segment) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Segment](#default-pool-origin-servers-private-name-segment) below.

&#x2022; [`site_locator`](#site-locator) - Optional Block<br>Site or Virtual Site. This message defines a reference to a site or virtual site object<br>See [Site Locator](#default-pool-origin-servers-private-name-site-locator) below.

&#x2022; [`snat_pool`](#snat-pool) - Optional Block<br>Snat Pool. Snat Pool configuration<br>See [Snat Pool](#default-pool-origin-servers-private-name-snat-pool) below.

#### Default Pool Origin Servers Private Name Segment

A [`segment`](#default-pool-origin-servers-private-name-segment) block (within [`default_pool.origin_servers.private_name`](#default-pool-origin-servers-private-name)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Default Pool Origin Servers Private Name Site Locator

A [`site_locator`](#default-pool-origin-servers-private-name-site-locator) block (within [`default_pool.origin_servers.private_name`](#default-pool-origin-servers-private-name)) supports the following:

&#x2022; [`site`](#site) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Site](#default-pool-origin-servers-private-name-site-locator-site) below.

&#x2022; [`virtual_site`](#virtual-site) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Virtual Site](#default-pool-origin-servers-private-name-site-locator-virtual-site) below.

#### Default Pool Origin Servers Private Name Site Locator Site

A [`site`](#default-pool-origin-servers-private-name-site-locator-site) block (within [`default_pool.origin_servers.private_name.site_locator`](#default-pool-origin-servers-private-name-site-locator)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Default Pool Origin Servers Private Name Site Locator Virtual Site

A [`virtual_site`](#default-pool-origin-servers-private-name-site-locator-virtual-site) block (within [`default_pool.origin_servers.private_name.site_locator`](#default-pool-origin-servers-private-name-site-locator)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Default Pool Origin Servers Private Name Snat Pool

A [`snat_pool`](#default-pool-origin-servers-private-name-snat-pool) block (within [`default_pool.origin_servers.private_name`](#default-pool-origin-servers-private-name)) supports the following:

&#x2022; [`no_snat_pool`](#no-snat-pool) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`snat_pool`](#snat-pool) - Optional Block<br>IPv4 Prefix List. x-example: '192.168.20.0/24' List of IPv4 prefixes that represent an endpoint<br>See [Snat Pool](#default-pool-origin-servers-private-name-snat-pool-snat-pool) below.

#### Default Pool Origin Servers Private Name Snat Pool Snat Pool

A [`snat_pool`](#default-pool-origin-servers-private-name-snat-pool-snat-pool) block (within [`default_pool.origin_servers.private_name.snat_pool`](#default-pool-origin-servers-private-name-snat-pool)) supports the following:

&#x2022; [`prefixes`](#prefixes) - Optional List<br>IPv4 Prefix List. List of IPv4 prefixes that represent an endpoint

#### Default Pool Origin Servers Public IP

A [`public_ip`](#default-pool-origin-servers-public-ip) block (within [`default_pool.origin_servers`](#default-pool-origin-servers)) supports the following:

&#x2022; [`ip`](#ip) - Optional String<br>Public IPv4. Public IPv4 address

#### Default Pool Origin Servers Public Name

A [`public_name`](#default-pool-origin-servers-public-name) block (within [`default_pool.origin_servers`](#default-pool-origin-servers)) supports the following:

&#x2022; [`dns_name`](#dns-name) - Optional String<br>DNS Name. DNS Name

&#x2022; [`refresh_interval`](#refresh-interval) - Optional Number<br>DNS Refresh Interval. Interval for DNS refresh in seconds. Max value is 7 days as per `HTTPS://datatracker.ietf.org/doc/HTML/rfc8767`

#### Default Pool Origin Servers Vn Private IP

A [`vn_private_ip`](#default-pool-origin-servers-vn-private-ip) block (within [`default_pool.origin_servers`](#default-pool-origin-servers)) supports the following:

&#x2022; [`ip`](#ip) - Optional String<br>IPv4. IPv4 address

&#x2022; [`virtual_network`](#virtual-network) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Virtual Network](#default-pool-origin-servers-vn-private-ip-virtual-network) below.

#### Default Pool Origin Servers Vn Private IP Virtual Network

A [`virtual_network`](#default-pool-origin-servers-vn-private-ip-virtual-network) block (within [`default_pool.origin_servers.vn_private_ip`](#default-pool-origin-servers-vn-private-ip)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Default Pool Origin Servers Vn Private Name

A [`vn_private_name`](#default-pool-origin-servers-vn-private-name) block (within [`default_pool.origin_servers`](#default-pool-origin-servers)) supports the following:

&#x2022; [`dns_name`](#dns-name) - Optional String<br>DNS Name. DNS Name

&#x2022; [`private_network`](#private-network) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Private Network](#default-pool-origin-servers-vn-private-name-private-network) below.

#### Default Pool Origin Servers Vn Private Name Private Network

A [`private_network`](#default-pool-origin-servers-vn-private-name-private-network) block (within [`default_pool.origin_servers.vn_private_name`](#default-pool-origin-servers-vn-private-name)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Default Pool Upstream Conn Pool Reuse Type

An [`upstream_conn_pool_reuse_type`](#default-pool-upstream-conn-pool-reuse-type) block (within [`default_pool`](#default-pool)) supports the following:

&#x2022; [`disable_conn_pool_reuse`](#disable-conn-pool-reuse) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`enable_conn_pool_reuse`](#enable-conn-pool-reuse) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### Default Pool Use TLS

An [`use_tls`](#default-pool-use-tls) block (within [`default_pool`](#default-pool)) supports the following:

&#x2022; [`default_session_key_caching`](#default-session-key-caching) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`disable_session_key_caching`](#disable-session-key-caching) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`disable_sni`](#disable-sni) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`max_session_keys`](#max-session-keys) - Optional Number<br>Max Session Keys Cached. x-example:'25' Number of session keys that are cached

&#x2022; [`no_mtls`](#no-mtls) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`skip_server_verification`](#skip-server-verification) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`sni`](#sni) - Optional String<br>SNI Value. SNI value to be used

&#x2022; [`tls_config`](#tls-config) - Optional Block<br>TLS Config. This defines various options to configure TLS configuration parameters<br>See [TLS Config](#default-pool-use-tls-tls-config) below.

&#x2022; [`use_host_header_as_sni`](#use-host-header-as-sni) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`use_mtls`](#use-mtls) - Optional Block<br>mTLS Certificate. mTLS Client Certificate<br>See [Use mTLS](#default-pool-use-tls-use-mtls) below.

&#x2022; [`use_mtls_obj`](#use-mtls-obj) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Use mTLS Obj](#default-pool-use-tls-use-mtls-obj) below.

&#x2022; [`use_server_verification`](#use-server-verification) - Optional Block<br>TLS Validation Context for Origin Servers. Upstream TLS Validation Context<br>See [Use Server Verification](#default-pool-use-tls-use-server-verification) below.

&#x2022; [`volterra_trusted_ca`](#volterra-trusted-ca) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### Default Pool Use TLS TLS Config

A [`tls_config`](#default-pool-use-tls-tls-config) block (within [`default_pool.use_tls`](#default-pool-use-tls)) supports the following:

&#x2022; [`custom_security`](#custom-security) - Optional Block<br>Custom Ciphers. This defines TLS protocol config including min/max versions and allowed ciphers<br>See [Custom Security](#default-pool-use-tls-tls-config-custom-security) below.

&#x2022; [`default_security`](#default-security) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`low_security`](#low-security) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`medium_security`](#medium-security) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### Default Pool Use TLS TLS Config Custom Security

A [`custom_security`](#default-pool-use-tls-tls-config-custom-security) block (within [`default_pool.use_tls.tls_config`](#default-pool-use-tls-tls-config)) supports the following:

&#x2022; [`cipher_suites`](#cipher-suites) - Optional List<br>Cipher Suites. The TLS listener will only support the specified cipher list

&#x2022; [`max_version`](#max-version) - Optional String  Defaults to `TLS_AUTO`<br>Possible values are `TLS_AUTO`, `TLSv1_0`, `TLSv1_1`, `TLSv1_2`, `TLSv1_3`<br>TLS Protocol. TlsProtocol is enumeration of supported TLS versions F5 Distributed Cloud will choose the optimal TLS version

&#x2022; [`min_version`](#min-version) - Optional String  Defaults to `TLS_AUTO`<br>Possible values are `TLS_AUTO`, `TLSv1_0`, `TLSv1_1`, `TLSv1_2`, `TLSv1_3`<br>TLS Protocol. TlsProtocol is enumeration of supported TLS versions F5 Distributed Cloud will choose the optimal TLS version

#### Default Pool Use TLS Use mTLS

An [`use_mtls`](#default-pool-use-tls-use-mtls) block (within [`default_pool.use_tls`](#default-pool-use-tls)) supports the following:

&#x2022; [`tls_certificates`](#tls-certificates) - Optional Block<br>mTLS Client Certificate. mTLS Client Certificate<br>See [TLS Certificates](#default-pool-use-tls-use-mtls-tls-certificates) below.

#### Default Pool Use TLS Use mTLS TLS Certificates

A [`tls_certificates`](#default-pool-use-tls-use-mtls-tls-certificates) block (within [`default_pool.use_tls.use_mtls`](#default-pool-use-tls-use-mtls)) supports the following:

&#x2022; [`certificate_url`](#certificate-url) - Optional String<br>Certificate. TLS certificate. Certificate or certificate chain in PEM format including the PEM headers

&#x2022; [`custom_hash_algorithms`](#custom-hash-algorithms) - Optional Block<br>Hash Algorithms. Specifies the hash algorithms to be used<br>See [Custom Hash Algorithms](#default-pool-use-tls-use-mtls-tls-certificates-custom-hash-algorithms) below.

&#x2022; [`description`](#description) - Optional String<br>Description. Description for the certificate

&#x2022; [`disable_ocsp_stapling`](#disable-ocsp-stapling) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`private_key`](#private-key) - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Private Key](#default-pool-use-tls-use-mtls-tls-certificates-private-key) below.

&#x2022; [`use_system_defaults`](#use-system-defaults) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### Default Pool Use TLS Use mTLS TLS Certificates Custom Hash Algorithms

A [`custom_hash_algorithms`](#default-pool-use-tls-use-mtls-tls-certificates-custom-hash-algorithms) block (within [`default_pool.use_tls.use_mtls.tls_certificates`](#default-pool-use-tls-use-mtls-tls-certificates)) supports the following:

&#x2022; [`hash_algorithms`](#hash-algorithms) - Optional List  Defaults to `INVALID_HASH_ALGORITHM`<br>Possible values are `INVALID_HASH_ALGORITHM`, `SHA256`, `SHA1`<br>Hash Algorithms. Ordered list of hash algorithms to be used

#### Default Pool Use TLS Use mTLS TLS Certificates Private Key

A [`private_key`](#default-pool-use-tls-use-mtls-tls-certificates-private-key) block (within [`default_pool.use_tls.use_mtls.tls_certificates`](#default-pool-use-tls-use-mtls-tls-certificates)) supports the following:

&#x2022; [`blindfold_secret_info`](#blindfold-secret-info) - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#default-pool-use-tls-use-mtls-tls-certificates-private-key-blindfold-secret-info) below.

&#x2022; [`clear_secret_info`](#clear-secret-info) - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#default-pool-use-tls-use-mtls-tls-certificates-private-key-clear-secret-info) below.

#### Default Pool Use TLS Use mTLS TLS Certificates Private Key Blindfold Secret Info

A [`blindfold_secret_info`](#default-pool-use-tls-use-mtls-tls-certificates-private-key-blindfold-secret-info) block (within [`default_pool.use_tls.use_mtls.tls_certificates.private_key`](#default-pool-use-tls-use-mtls-tls-certificates-private-key)) supports the following:

&#x2022; [`decryption_provider`](#decryption-provider) - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

&#x2022; [`location`](#location) - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

&#x2022; [`store_provider`](#store-provider) - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

#### Default Pool Use TLS Use mTLS TLS Certificates Private Key Clear Secret Info

A [`clear_secret_info`](#default-pool-use-tls-use-mtls-tls-certificates-private-key-clear-secret-info) block (within [`default_pool.use_tls.use_mtls.tls_certificates.private_key`](#default-pool-use-tls-use-mtls-tls-certificates-private-key)) supports the following:

&#x2022; [`provider_ref`](#provider-ref) - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

&#x2022; [`url`](#url) - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

#### Default Pool Use TLS Use mTLS Obj

An [`use_mtls_obj`](#default-pool-use-tls-use-mtls-obj) block (within [`default_pool.use_tls`](#default-pool-use-tls)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Default Pool Use TLS Use Server Verification

An [`use_server_verification`](#default-pool-use-tls-use-server-verification) block (within [`default_pool.use_tls`](#default-pool-use-tls)) supports the following:

&#x2022; [`trusted_ca`](#trusted-ca) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Trusted CA](#default-pool-use-tls-use-server-verification-trusted-ca) below.

&#x2022; [`trusted_ca_url`](#trusted-ca-url) - Optional String<br>Inline Root CA Certificate (legacy). Upload a Root CA Certificate specifically for this Origin Pool for verification of server's certificate

#### Default Pool Use TLS Use Server Verification Trusted CA

A [`trusted_ca`](#default-pool-use-tls-use-server-verification-trusted-ca) block (within [`default_pool.use_tls.use_server_verification`](#default-pool-use-tls-use-server-verification)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Default Pool View Internal

A [`view_internal`](#default-pool-view-internal) block (within [`default_pool`](#default-pool)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Default Pool List

A [`default_pool_list`](#default-pool-list) block supports the following:

&#x2022; [`pools`](#pools) - Optional Block<br>Origin Pools. List of Origin Pools<br>See [Pools](#default-pool-list-pools) below.

#### Default Pool List Pools

A [`pools`](#default-pool-list-pools) block (within [`default_pool_list`](#default-pool-list)) supports the following:

&#x2022; [`cluster`](#cluster) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Cluster](#default-pool-list-pools-cluster) below.

&#x2022; [`endpoint_subsets`](#endpoint-subsets) - Optional Block<br>Origin Servers Subsets. Upstream origin pool may be configured to divide its origin servers into subsets based on metadata attached to the origin servers. Routes may then specify the metadata that a endpoint must match in order to be selected by the load balancer For origin servers which are discovered in K8S or Consul cluster, the label of the service is merged with endpoint's labels. In case of Consul, the label is derived from the 'Tag' field. For labels that are common between configured endpoint and discovered service, labels from discovered service takes precedence. List of key-value pairs that will be used as matching metadata. Only those origin servers of upstream origin pool which match this metadata will be selected for load balancing

&#x2022; [`pool`](#pool) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Pool](#default-pool-list-pools-pool) below.

&#x2022; [`priority`](#priority) - Optional Number<br>Priority. Priority of this origin pool, valid only with multiple origin pools. Value of 0 will make the pool as lowest priority origin pool Priority of 1 means highest priority and is considered active. When active origin pool is not available, lower priority origin pools are made active as per the increasing priority

&#x2022; [`weight`](#weight) - Optional Number<br>Weight. Weight of this origin pool, valid only with multiple origin pool. Value of 0 will disable the pool

#### Default Pool List Pools Cluster

A [`cluster`](#default-pool-list-pools-cluster) block (within [`default_pool_list.pools`](#default-pool-list-pools)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Default Pool List Pools Pool

A [`pool`](#default-pool-list-pools-pool) block (within [`default_pool_list.pools`](#default-pool-list-pools)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Default Route Pools

A [`default_route_pools`](#default-route-pools) block supports the following:

&#x2022; [`cluster`](#cluster) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Cluster](#default-route-pools-cluster) below.

&#x2022; [`endpoint_subsets`](#endpoint-subsets) - Optional Block<br>Origin Servers Subsets. Upstream origin pool may be configured to divide its origin servers into subsets based on metadata attached to the origin servers. Routes may then specify the metadata that a endpoint must match in order to be selected by the load balancer For origin servers which are discovered in K8S or Consul cluster, the label of the service is merged with endpoint's labels. In case of Consul, the label is derived from the 'Tag' field. For labels that are common between configured endpoint and discovered service, labels from discovered service takes precedence. List of key-value pairs that will be used as matching metadata. Only those origin servers of upstream origin pool which match this metadata will be selected for load balancing

&#x2022; [`pool`](#pool) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Pool](#default-route-pools-pool) below.

&#x2022; [`priority`](#priority) - Optional Number<br>Priority. Priority of this origin pool, valid only with multiple origin pools. Value of 0 will make the pool as lowest priority origin pool Priority of 1 means highest priority and is considered active. When active origin pool is not available, lower priority origin pools are made active as per the increasing priority

&#x2022; [`weight`](#weight) - Optional Number<br>Weight. Weight of this origin pool, valid only with multiple origin pool. Value of 0 will disable the pool

#### Default Route Pools Cluster

A [`cluster`](#default-route-pools-cluster) block (within [`default_route_pools`](#default-route-pools)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Default Route Pools Pool

A [`pool`](#default-route-pools-pool) block (within [`default_route_pools`](#default-route-pools)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Enable API Discovery

An [`enable_api_discovery`](#enable-api-discovery) block supports the following:

&#x2022; [`api_crawler`](#api-crawler) - Optional Block<br>API Crawling. API Crawler message<br>See [API Crawler](#enable-api-discovery-api-crawler) below.

&#x2022; [`api_discovery_from_code_scan`](#api-discovery-from-code-scan) - Optional Block<br>Select Code Base and Repositories. x-required<br>See [API Discovery From Code Scan](#enable-api-discovery-api-discovery-from-code-scan) below.

&#x2022; [`custom_api_auth_discovery`](#custom-api-auth-discovery) - Optional Block<br>API Discovery Advanced Settings. API Discovery Advanced settings<br>See [Custom API Auth Discovery](#enable-api-discovery-custom-api-auth-discovery) below.

&#x2022; [`default_api_auth_discovery`](#default-api-auth-discovery) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`disable_learn_from_redirect_traffic`](#disable-learn-from-redirect-traffic) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`discovered_api_settings`](#discovered-api-settings) - Optional Block<br>Discovered API Settings. x-example: '2' Configure Discovered API Settings<br>See [Discovered API Settings](#enable-api-discovery-discovered-api-settings) below.

&#x2022; [`enable_learn_from_redirect_traffic`](#enable-learn-from-redirect-traffic) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### Enable API Discovery API Crawler

An [`api_crawler`](#enable-api-discovery-api-crawler) block (within [`enable_api_discovery`](#enable-api-discovery)) supports the following:

&#x2022; [`api_crawler_config`](#api-crawler-config) - Optional Block<br>Crawler Configure<br>See [API Crawler Config](#enable-api-discovery-api-crawler-api-crawler-config) below.

&#x2022; [`disable_api_crawler`](#disable-api-crawler) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### Enable API Discovery API Crawler API Crawler Config

An [`api_crawler_config`](#enable-api-discovery-api-crawler-api-crawler-config) block (within [`enable_api_discovery.api_crawler`](#enable-api-discovery-api-crawler)) supports the following:

&#x2022; [`domains`](#domains) - Optional Block<br>Domains to Crawl. Enter domains and their credentials to allow authenticated API crawling. You can only include domains you own that are associated with this Load Balancer<br>See [Domains](#enable-api-discovery-api-crawler-api-crawler-config-domains) below.

#### Enable API Discovery API Crawler API Crawler Config Domains

A [`domains`](#enable-api-discovery-api-crawler-api-crawler-config-domains) block (within [`enable_api_discovery.api_crawler.api_crawler_config`](#enable-api-discovery-api-crawler-api-crawler-config)) supports the following:

&#x2022; [`domain`](#domain) - Optional String<br>Domain. Select the domain to execute API Crawling with given credentials

&#x2022; [`simple_login`](#simple-login) - Optional Block<br>Simple Login<br>See [Simple Login](#enable-api-discovery-api-crawler-api-crawler-config-domains-simple-login) below.

#### Enable API Discovery API Crawler API Crawler Config Domains Simple Login

A [`simple_login`](#enable-api-discovery-api-crawler-api-crawler-config-domains-simple-login) block (within [`enable_api_discovery.api_crawler.api_crawler_config.domains`](#enable-api-discovery-api-crawler-api-crawler-config-domains)) supports the following:

&#x2022; [`password`](#password) - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Password](#enable-api-discovery-api-crawler-api-crawler-config-domains-simple-login-password) below.

&#x2022; [`user`](#user) - Optional String<br>User. Enter the username to assign credentials for the selected domain to crawl

#### Enable API Discovery API Crawler API Crawler Config Domains Simple Login Password

A [`password`](#enable-api-discovery-api-crawler-api-crawler-config-domains-simple-login-password) block (within [`enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login`](#enable-api-discovery-api-crawler-api-crawler-config-domains-simple-login)) supports the following:

&#x2022; [`blindfold_secret_info`](#blindfold-secret-info) - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#enable-api-discovery-api-crawler-api-crawler-config-domains-simple-login-password-blindfold-secret-info) below.

&#x2022; [`clear_secret_info`](#clear-secret-info) - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#enable-api-discovery-api-crawler-api-crawler-config-domains-simple-login-password-clear-secret-info) below.

#### Enable API Discovery API Crawler API Crawler Config Domains Simple Login Password Blindfold Secret Info

A [`blindfold_secret_info`](#enable-api-discovery-api-crawler-api-crawler-config-domains-simple-login-password-blindfold-secret-info) block (within [`enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password`](#enable-api-discovery-api-crawler-api-crawler-config-domains-simple-login-password)) supports the following:

&#x2022; [`decryption_provider`](#decryption-provider) - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

&#x2022; [`location`](#location) - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

&#x2022; [`store_provider`](#store-provider) - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

#### Enable API Discovery API Crawler API Crawler Config Domains Simple Login Password Clear Secret Info

A [`clear_secret_info`](#enable-api-discovery-api-crawler-api-crawler-config-domains-simple-login-password-clear-secret-info) block (within [`enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password`](#enable-api-discovery-api-crawler-api-crawler-config-domains-simple-login-password)) supports the following:

&#x2022; [`provider_ref`](#provider-ref) - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

&#x2022; [`url`](#url) - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

#### Enable API Discovery API Discovery From Code Scan

An [`api_discovery_from_code_scan`](#enable-api-discovery-api-discovery-from-code-scan) block (within [`enable_api_discovery`](#enable-api-discovery)) supports the following:

&#x2022; [`code_base_integrations`](#code-base-integrations) - Optional Block<br>Select Code Base Integrations<br>See [Code Base Integrations](#enable-api-discovery-api-discovery-from-code-scan-code-base-integrations) below.

#### Enable API Discovery API Discovery From Code Scan Code Base Integrations

A [`code_base_integrations`](#enable-api-discovery-api-discovery-from-code-scan-code-base-integrations) block (within [`enable_api_discovery.api_discovery_from_code_scan`](#enable-api-discovery-api-discovery-from-code-scan)) supports the following:

&#x2022; [`all_repos`](#all-repos) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`code_base_integration`](#code-base-integration) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Code Base Integration](#enable-api-discovery-api-discovery-from-code-scan-code-base-integrations-code-base-integration) below.

&#x2022; [`selected_repos`](#selected-repos) - Optional Block<br>API Code Repositories. Select which API repositories represent the LB applications<br>See [Selected Repos](#enable-api-discovery-api-discovery-from-code-scan-code-base-integrations-selected-repos) below.

#### Enable API Discovery API Discovery From Code Scan Code Base Integrations Code Base Integration

A [`code_base_integration`](#enable-api-discovery-api-discovery-from-code-scan-code-base-integrations-code-base-integration) block (within [`enable_api_discovery.api_discovery_from_code_scan.code_base_integrations`](#enable-api-discovery-api-discovery-from-code-scan-code-base-integrations)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Enable API Discovery API Discovery From Code Scan Code Base Integrations Selected Repos

A [`selected_repos`](#enable-api-discovery-api-discovery-from-code-scan-code-base-integrations-selected-repos) block (within [`enable_api_discovery.api_discovery_from_code_scan.code_base_integrations`](#enable-api-discovery-api-discovery-from-code-scan-code-base-integrations)) supports the following:

&#x2022; [`api_code_repo`](#api-code-repo) - Optional List<br>API Code Repository. Code repository which contain API endpoints

#### Enable API Discovery Custom API Auth Discovery

A [`custom_api_auth_discovery`](#enable-api-discovery-custom-api-auth-discovery) block (within [`enable_api_discovery`](#enable-api-discovery)) supports the following:

&#x2022; [`api_discovery_ref`](#api-discovery-ref) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [API Discovery Ref](#enable-api-discovery-custom-api-auth-discovery-api-discovery-ref) below.

#### Enable API Discovery Custom API Auth Discovery API Discovery Ref

An [`api_discovery_ref`](#enable-api-discovery-custom-api-auth-discovery-api-discovery-ref) block (within [`enable_api_discovery.custom_api_auth_discovery`](#enable-api-discovery-custom-api-auth-discovery)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Enable API Discovery Discovered API Settings

A [`discovered_api_settings`](#enable-api-discovery-discovered-api-settings) block (within [`enable_api_discovery`](#enable-api-discovery)) supports the following:

&#x2022; [`purge_duration_for_inactive_discovered_apis`](#purge-duration-for-inactive-discovered-apis) - Optional Number<br>Purge Duration for Inactive Discovered APIs from Traffic. Inactive discovered API will be deleted after configured duration

#### Enable Challenge

An [`enable_challenge`](#enable-challenge) block supports the following:

&#x2022; [`captcha_challenge_parameters`](#captcha-challenge-parameters) - Optional Block<br>Captcha Challenge Parameters. Enables loadbalancer to perform captcha challenge Captcha challenge will be based on Google Recaptcha. With this feature enabled, only clients that pass the captcha challenge will be allowed to complete the HTTP request. When loadbalancer is configured to do Captcha Challenge, it will redirect the browser to an HTML page on every new HTTP request. This HTML page will have captcha challenge embedded in it. Client will be allowed to make the request only if the captcha challenge is successful. Loadbalancer will tag response header with a cookie to avoid Captcha challenge for subsequent requests. CAPTCHA is mainly used as a security check to ensure only human users can pass through. Generally, computers or bots are not capable of solving a captcha. You can enable either Javascript challenge or Captcha challenge on a virtual host<br>See [Captcha Challenge Parameters](#enable-challenge-captcha-challenge-parameters) below.

&#x2022; [`default_captcha_challenge_parameters`](#default-captcha-challenge-parameters) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`default_js_challenge_parameters`](#default-js-challenge-parameters) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`default_mitigation_settings`](#default-mitigation-settings) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`js_challenge_parameters`](#js-challenge-parameters) - Optional Block<br>Javascript Challenge Parameters. Enables loadbalancer to perform client browser compatibility test by redirecting to a page with Javascript. With this feature enabled, only clients that are capable of executing Javascript(mostly browsers) will be allowed to complete the HTTP request. When loadbalancer is configured to do Javascript Challenge, it will redirect the browser to an HTML page on every new HTTP request. This HTML page will have Javascript embedded in it. Loadbalancer chooses a set of random numbers for every new client and sends these numbers along with an encrypted answer with the request such that it embed these numbers as input in the Javascript. Javascript will run on the requestor browser and perform a complex Math operation. Script will submit the answer to loadbalancer. Loadbalancer will validate the answer by comparing the calculated answer with the decrypted answer (which was encrypted when it was sent back as reply) and allow the request to the upstream server only if the answer is correct. Loadbalancer will tag response header with a cookie to avoid Javascript challenge for subsequent requests. Javascript challenge serves following purposes * Validate that the request is coming via a browser that is capable for running Javascript * Force the browser to run a complex operation, f(X), that requires it to spend a large number of CPU cycles. This is to slow down a potential DOS attacker by making it difficult to launch a large request flood without having to spend even larger CPU cost at their end. You can enable either Javascript challenge or Captcha challenge on a virtual host<br>See [Js Challenge Parameters](#enable-challenge-js-challenge-parameters) below.

&#x2022; [`malicious_user_mitigation`](#malicious-user-mitigation) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Malicious User Mitigation](#enable-challenge-malicious-user-mitigation) below.

#### Enable Challenge Captcha Challenge Parameters

A [`captcha_challenge_parameters`](#enable-challenge-captcha-challenge-parameters) block (within [`enable_challenge`](#enable-challenge)) supports the following:

&#x2022; [`cookie_expiry`](#cookie-expiry) - Optional Number<br>Cookie Expiration Period. Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new challenge

&#x2022; [`custom_page`](#custom-page) - Optional String<br>Custom message for Captcha Challenge. Custom message is of type uri_ref. Currently supported URL schemes is string:///. For string:/// scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. 'Please Wait.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '<p> Please Wait </p>'. Base64 encoded string for this HTML is 'PHA+IFBsZWFzZSBXYWl0IDwvcD4='

#### Enable Challenge Js Challenge Parameters

A [`js_challenge_parameters`](#enable-challenge-js-challenge-parameters) block (within [`enable_challenge`](#enable-challenge)) supports the following:

&#x2022; [`cookie_expiry`](#cookie-expiry) - Optional Number<br>Cookie Expiration Period. Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new challenge

&#x2022; [`custom_page`](#custom-page) - Optional String<br>Custom Message for Javascript Challenge. Custom message is of type uri_ref. Currently supported URL schemes is string:///. For string:/// scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. 'Please Wait.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '<p> Please Wait </p>'. Base64 encoded string for this HTML is 'PHA+IFBsZWFzZSBXYWl0IDwvcD4='

&#x2022; [`js_script_delay`](#js-script-delay) - Optional Number<br>Javascript Delay. Delay introduced by Javascript, in milliseconds

#### Enable Challenge Malicious User Mitigation

A [`malicious_user_mitigation`](#enable-challenge-malicious-user-mitigation) block (within [`enable_challenge`](#enable-challenge)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Enable IP Reputation

An [`enable_ip_reputation`](#enable-ip-reputation) block supports the following:

&#x2022; [`ip_threat_categories`](#ip-threat-categories) - Optional List  Defaults to `SPAM_SOURCES`<br>Possible values are `SPAM_SOURCES`, `WINDOWS_EXPLOITS`, `WEB_ATTACKS`, `BOTNETS`, `SCANNERS`, `REPUTATION`, `PHISHING`, `PROXY`, `MOBILE_THREATS`, `TOR_PROXY`, `DENIAL_OF_SERVICE`, `NETWORK`<br>List of IP Threat Categories to choose. If the source IP matches on atleast one of the enabled IP threat categories, the request will be denied

#### Enable Trust Client IP Headers

An [`enable_trust_client_ip_headers`](#enable-trust-client-ip-headers) block supports the following:

&#x2022; [`client_ip_headers`](#client-ip-headers) - Optional List<br>Client IP Headers. Define the list of one or more Client IP Headers. Headers will be used in order from top to bottom, meaning if the first header is not present in the request, the system will proceed to check for the second header, and so on, until one of the listed headers is found. If none of the defined headers exist, or the value is not an IP address, then the system will use the source IP of the packet. If multiple defined headers with different names are present in the request, the value of the first header name in the configuration will be used. If multiple defined headers with the same name are present in the request, values of all those headers will be combined. The system will read the right-most IP address from header, if there are multiple IP addresses in the header value. For X-Forwarded-For header, the system will read the IP address(rightmost - 1), as the client IP

#### GraphQL Rules

A [`graphql_rules`](#graphql-rules) block supports the following:

&#x2022; [`any_domain`](#any-domain) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`exact_path`](#exact-path) - Optional String  Defaults to `/GraphQL`<br>Path. Specifies the exact path to GraphQL endpoint

&#x2022; [`exact_value`](#exact-value) - Optional String<br>Exact Value. Exact domain name

&#x2022; [`graphql_settings`](#graphql-settings) - Optional Block<br>GraphQL Settings. GraphQL configuration<br>See [GraphQL Settings](#graphql-rules-graphql-settings) below.

&#x2022; [`metadata`](#metadata) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#graphql-rules-metadata) below.

&#x2022; [`method_get`](#method-get) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`method_post`](#method-post) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`suffix_value`](#suffix-value) - Optional String<br>Suffix Value. Suffix of domain name e.g 'xyz.com' will match '*.xyz.com' and 'xyz.com'

#### GraphQL Rules GraphQL Settings

A [`graphql_settings`](#graphql-rules-graphql-settings) block (within [`graphql_rules`](#graphql-rules)) supports the following:

&#x2022; [`disable_introspection`](#disable-introspection) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`enable_introspection`](#enable-introspection) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`max_batched_queries`](#max-batched-queries) - Optional Number<br>Maximum Batched Queries. Specify maximum number of queries in a single batched request

&#x2022; [`max_depth`](#max-depth) - Optional Number<br>Maximum Structure Depth. Specify maximum depth for the GraphQL query

&#x2022; [`max_total_length`](#max-total-length) - Optional Number<br>Maximum Total Length. Specify maximum length in bytes for the GraphQL query

#### GraphQL Rules Metadata

A [`metadata`](#graphql-rules-metadata) block (within [`graphql_rules`](#graphql-rules)) supports the following:

&#x2022; [`description`](#description) - Optional String<br>Description. Human readable description

&#x2022; [`name`](#name) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### HTTP

A [`http`](#http) block supports the following:

&#x2022; [`dns_volterra_managed`](#dns-volterra-managed) - Optional Bool<br>Automatically Manage DNS Records. DNS records for domains will be managed automatically by F5 Distributed Cloud. As a prerequisite, the domain must be delegated to F5 Distributed Cloud using Delegated domain feature or a DNS CNAME record should be created in your DNS provider's portal

&#x2022; [`port`](#port) - Optional Number<br>HTTP Listen Port. HTTP port to Listen

&#x2022; [`port_ranges`](#port-ranges) - Optional String<br>Port Ranges. A string containing a comma separated list of port ranges. Each port range consists of a single port or two ports separated by '-'

#### HTTPS

A [`https`](#https) block supports the following:

&#x2022; [`add_hsts`](#add-hsts) - Optional Bool<br>Add HSTS Header. Add HTTP Strict-Transport-Security response header

&#x2022; [`append_server_name`](#append-server-name) - Optional String<br>Append header value. Define the header value for the header name “server”. If header value is already present, it is not overwritten and passed as-is

&#x2022; [`coalescing_options`](#coalescing-options) - Optional Block<br>TLS Coalescing Options. TLS connection coalescing configuration (not compatible with mTLS)<br>See [Coalescing Options](#https-coalescing-options) below.

&#x2022; [`connection_idle_timeout`](#connection-idle-timeout) - Optional Number  Defaults to `2`  Specified in milliseconds<br>Connection Idle Timeout. The idle timeout for downstream connections. The idle timeout is defined as the period in which there are no active requests. When the idle timeout is reached the connection will be closed. Note that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive.  The minutes

&#x2022; [`default_header`](#default-header) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`default_loadbalancer`](#default-loadbalancer) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`disable_path_normalize`](#disable-path-normalize) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`enable_path_normalize`](#enable-path-normalize) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`http_protocol_options`](#http-protocol-options) - Optional Block<br>HTTP Protocol Configuration Options. HTTP protocol configuration options for downstream connections<br>See [HTTP Protocol Options](#https-http-protocol-options) below.

&#x2022; [`http_redirect`](#http-redirect) - Optional Bool<br>HTTP Redirect to HTTPS. Redirect HTTP traffic to HTTPS

&#x2022; [`non_default_loadbalancer`](#non-default-loadbalancer) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`pass_through`](#pass-through) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`port`](#port) - Optional Number<br>HTTPS Port. HTTPS port to Listen

&#x2022; [`port_ranges`](#port-ranges) - Optional String<br>Port Ranges. A string containing a comma separated list of port ranges. Each port range consists of a single port or two ports separated by '-'

&#x2022; [`server_name`](#server-name) - Optional String<br>Modify header value. Define the header value for the header name “server”. This will overwrite existing values, if any, for the server header

&#x2022; [`tls_cert_params`](#tls-cert-params) - Optional Block<br>TLS Parameters. Select TLS Parameters and Certificates<br>See [TLS Cert Params](#https-tls-cert-params) below.

&#x2022; [`tls_parameters`](#tls-parameters) - Optional Block<br>Inline TLS Parameters. Inline TLS parameters<br>See [TLS Parameters](#https-tls-parameters) below.

#### HTTPS Coalescing Options

A [`coalescing_options`](#https-coalescing-options) block (within [`https`](#https)) supports the following:

&#x2022; [`default_coalescing`](#default-coalescing) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`strict_coalescing`](#strict-coalescing) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### HTTPS HTTP Protocol Options

A [`http_protocol_options`](#https-http-protocol-options) block (within [`https`](#https)) supports the following:

&#x2022; [`http_protocol_enable_v1_only`](#http-protocol-enable-v1-only) - Optional Block<br>HTTP/1.1 Protocol Options. HTTP/1.1 Protocol options for downstream connections<br>See [HTTP Protocol Enable V1 Only](#https-http-protocol-options-http-protocol-enable-v1-only) below.

&#x2022; [`http_protocol_enable_v1_v2`](#http-protocol-enable-v1-v2) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`http_protocol_enable_v2_only`](#http-protocol-enable-v2-only) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### HTTPS HTTP Protocol Options HTTP Protocol Enable V1 Only

A [`http_protocol_enable_v1_only`](#https-http-protocol-options-http-protocol-enable-v1-only) block (within [`https.http_protocol_options`](#https-http-protocol-options)) supports the following:

&#x2022; [`header_transformation`](#header-transformation) - Optional Block<br>Header Transformation. Header Transformation options for HTTP/1.1 request/response headers<br>See [Header Transformation](#https-http-protocol-options-http-protocol-enable-v1-only-header-transformation) below.

#### HTTPS HTTP Protocol Options HTTP Protocol Enable V1 Only Header Transformation

A [`header_transformation`](#https-http-protocol-options-http-protocol-enable-v1-only-header-transformation) block (within [`https.http_protocol_options.http_protocol_enable_v1_only`](#https-http-protocol-options-http-protocol-enable-v1-only)) supports the following:

&#x2022; [`default_header_transformation`](#default-header-transformation) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`legacy_header_transformation`](#legacy-header-transformation) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`preserve_case_header_transformation`](#preserve-case-header-transformation) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`proper_case_header_transformation`](#proper-case-header-transformation) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### HTTPS TLS Cert Params

A [`tls_cert_params`](#https-tls-cert-params) block (within [`https`](#https)) supports the following:

&#x2022; [`certificates`](#certificates) - Optional Block<br>Certificates. Select one or more certificates with any domain names<br>See [Certificates](#https-tls-cert-params-certificates) below.

&#x2022; [`no_mtls`](#no-mtls) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`tls_config`](#tls-config) - Optional Block<br>TLS Config. This defines various options to configure TLS configuration parameters<br>See [TLS Config](#https-tls-cert-params-tls-config) below.

&#x2022; [`use_mtls`](#use-mtls) - Optional Block<br>Clients TLS validation context. Validation context for downstream client TLS connections<br>See [Use mTLS](#https-tls-cert-params-use-mtls) below.

#### HTTPS TLS Cert Params Certificates

A [`certificates`](#https-tls-cert-params-certificates) block (within [`https.tls_cert_params`](#https-tls-cert-params)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### HTTPS TLS Cert Params TLS Config

A [`tls_config`](#https-tls-cert-params-tls-config) block (within [`https.tls_cert_params`](#https-tls-cert-params)) supports the following:

&#x2022; [`custom_security`](#custom-security) - Optional Block<br>Custom Ciphers. This defines TLS protocol config including min/max versions and allowed ciphers<br>See [Custom Security](#https-tls-cert-params-tls-config-custom-security) below.

&#x2022; [`default_security`](#default-security) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`low_security`](#low-security) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`medium_security`](#medium-security) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### HTTPS TLS Cert Params TLS Config Custom Security

A [`custom_security`](#https-tls-cert-params-tls-config-custom-security) block (within [`https.tls_cert_params.tls_config`](#https-tls-cert-params-tls-config)) supports the following:

&#x2022; [`cipher_suites`](#cipher-suites) - Optional List<br>Cipher Suites. The TLS listener will only support the specified cipher list

&#x2022; [`max_version`](#max-version) - Optional String  Defaults to `TLS_AUTO`<br>Possible values are `TLS_AUTO`, `TLSv1_0`, `TLSv1_1`, `TLSv1_2`, `TLSv1_3`<br>TLS Protocol. TlsProtocol is enumeration of supported TLS versions F5 Distributed Cloud will choose the optimal TLS version

&#x2022; [`min_version`](#min-version) - Optional String  Defaults to `TLS_AUTO`<br>Possible values are `TLS_AUTO`, `TLSv1_0`, `TLSv1_1`, `TLSv1_2`, `TLSv1_3`<br>TLS Protocol. TlsProtocol is enumeration of supported TLS versions F5 Distributed Cloud will choose the optimal TLS version

#### HTTPS TLS Cert Params Use mTLS

An [`use_mtls`](#https-tls-cert-params-use-mtls) block (within [`https.tls_cert_params`](#https-tls-cert-params)) supports the following:

&#x2022; [`client_certificate_optional`](#client-certificate-optional) - Optional Bool<br>Client Certificate Optional. Client certificate is optional. If the client has provided a certificate, the load balancer will verify it. If certification verification fails, the connection will be terminated. If the client does not provide a certificate, the connection will be accepted

&#x2022; [`crl`](#crl) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [CRL](#https-tls-cert-params-use-mtls-crl) below.

&#x2022; [`no_crl`](#no-crl) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`trusted_ca`](#trusted-ca) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Trusted CA](#https-tls-cert-params-use-mtls-trusted-ca) below.

&#x2022; [`trusted_ca_url`](#trusted-ca-url) - Optional String<br>Inline Root CA Certificate (legacy). Upload a Root CA Certificate specifically for this Load Balancer

&#x2022; [`xfcc_disabled`](#xfcc-disabled) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`xfcc_options`](#xfcc-options) - Optional Block<br>XFCC Header Elements. X-Forwarded-Client-Cert header elements to be added to requests<br>See [Xfcc Options](#https-tls-cert-params-use-mtls-xfcc-options) below.

#### HTTPS TLS Cert Params Use mTLS CRL

A [`crl`](#https-tls-cert-params-use-mtls-crl) block (within [`https.tls_cert_params.use_mtls`](#https-tls-cert-params-use-mtls)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### HTTPS TLS Cert Params Use mTLS Trusted CA

A [`trusted_ca`](#https-tls-cert-params-use-mtls-trusted-ca) block (within [`https.tls_cert_params.use_mtls`](#https-tls-cert-params-use-mtls)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### HTTPS TLS Cert Params Use mTLS Xfcc Options

A [`xfcc_options`](#https-tls-cert-params-use-mtls-xfcc-options) block (within [`https.tls_cert_params.use_mtls`](#https-tls-cert-params-use-mtls)) supports the following:

&#x2022; [`xfcc_header_elements`](#xfcc-header-elements) - Optional List  Defaults to `XFCC_NONE`<br>Possible values are `XFCC_NONE`, `XFCC_CERT`, `XFCC_CHAIN`, `XFCC_SUBJECT`, `XFCC_URI`, `XFCC_DNS`<br>XFCC Header Elements. X-Forwarded-Client-Cert header elements to be added to requests

#### HTTPS TLS Parameters

A [`tls_parameters`](#https-tls-parameters) block (within [`https`](#https)) supports the following:

&#x2022; [`no_mtls`](#no-mtls) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`tls_certificates`](#tls-certificates) - Optional Block<br>TLS Certificates. Users can add one or more certificates that share the same set of domains. for example, domain.com and *.domain.com - but use different signature algorithms<br>See [TLS Certificates](#https-tls-parameters-tls-certificates) below.

&#x2022; [`tls_config`](#tls-config) - Optional Block<br>TLS Config. This defines various options to configure TLS configuration parameters<br>See [TLS Config](#https-tls-parameters-tls-config) below.

&#x2022; [`use_mtls`](#use-mtls) - Optional Block<br>Clients TLS validation context. Validation context for downstream client TLS connections<br>See [Use mTLS](#https-tls-parameters-use-mtls) below.

#### HTTPS TLS Parameters TLS Certificates

A [`tls_certificates`](#https-tls-parameters-tls-certificates) block (within [`https.tls_parameters`](#https-tls-parameters)) supports the following:

&#x2022; [`certificate_url`](#certificate-url) - Optional String<br>Certificate. TLS certificate. Certificate or certificate chain in PEM format including the PEM headers

&#x2022; [`custom_hash_algorithms`](#custom-hash-algorithms) - Optional Block<br>Hash Algorithms. Specifies the hash algorithms to be used<br>See [Custom Hash Algorithms](#https-tls-parameters-tls-certificates-custom-hash-algorithms) below.

&#x2022; [`description`](#description) - Optional String<br>Description. Description for the certificate

&#x2022; [`disable_ocsp_stapling`](#disable-ocsp-stapling) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`private_key`](#private-key) - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Private Key](#https-tls-parameters-tls-certificates-private-key) below.

&#x2022; [`use_system_defaults`](#use-system-defaults) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### HTTPS TLS Parameters TLS Certificates Custom Hash Algorithms

A [`custom_hash_algorithms`](#https-tls-parameters-tls-certificates-custom-hash-algorithms) block (within [`https.tls_parameters.tls_certificates`](#https-tls-parameters-tls-certificates)) supports the following:

&#x2022; [`hash_algorithms`](#hash-algorithms) - Optional List  Defaults to `INVALID_HASH_ALGORITHM`<br>Possible values are `INVALID_HASH_ALGORITHM`, `SHA256`, `SHA1`<br>Hash Algorithms. Ordered list of hash algorithms to be used

#### HTTPS TLS Parameters TLS Certificates Private Key

A [`private_key`](#https-tls-parameters-tls-certificates-private-key) block (within [`https.tls_parameters.tls_certificates`](#https-tls-parameters-tls-certificates)) supports the following:

&#x2022; [`blindfold_secret_info`](#blindfold-secret-info) - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#https-tls-parameters-tls-certificates-private-key-blindfold-secret-info) below.

&#x2022; [`clear_secret_info`](#clear-secret-info) - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#https-tls-parameters-tls-certificates-private-key-clear-secret-info) below.

#### HTTPS TLS Parameters TLS Certificates Private Key Blindfold Secret Info

A [`blindfold_secret_info`](#https-tls-parameters-tls-certificates-private-key-blindfold-secret-info) block (within [`https.tls_parameters.tls_certificates.private_key`](#https-tls-parameters-tls-certificates-private-key)) supports the following:

&#x2022; [`decryption_provider`](#decryption-provider) - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

&#x2022; [`location`](#location) - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

&#x2022; [`store_provider`](#store-provider) - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

#### HTTPS TLS Parameters TLS Certificates Private Key Clear Secret Info

A [`clear_secret_info`](#https-tls-parameters-tls-certificates-private-key-clear-secret-info) block (within [`https.tls_parameters.tls_certificates.private_key`](#https-tls-parameters-tls-certificates-private-key)) supports the following:

&#x2022; [`provider_ref`](#provider-ref) - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

&#x2022; [`url`](#url) - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

#### HTTPS TLS Parameters TLS Config

A [`tls_config`](#https-tls-parameters-tls-config) block (within [`https.tls_parameters`](#https-tls-parameters)) supports the following:

&#x2022; [`custom_security`](#custom-security) - Optional Block<br>Custom Ciphers. This defines TLS protocol config including min/max versions and allowed ciphers<br>See [Custom Security](#https-tls-parameters-tls-config-custom-security) below.

&#x2022; [`default_security`](#default-security) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`low_security`](#low-security) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`medium_security`](#medium-security) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### HTTPS TLS Parameters TLS Config Custom Security

A [`custom_security`](#https-tls-parameters-tls-config-custom-security) block (within [`https.tls_parameters.tls_config`](#https-tls-parameters-tls-config)) supports the following:

&#x2022; [`cipher_suites`](#cipher-suites) - Optional List<br>Cipher Suites. The TLS listener will only support the specified cipher list

&#x2022; [`max_version`](#max-version) - Optional String  Defaults to `TLS_AUTO`<br>Possible values are `TLS_AUTO`, `TLSv1_0`, `TLSv1_1`, `TLSv1_2`, `TLSv1_3`<br>TLS Protocol. TlsProtocol is enumeration of supported TLS versions F5 Distributed Cloud will choose the optimal TLS version

&#x2022; [`min_version`](#min-version) - Optional String  Defaults to `TLS_AUTO`<br>Possible values are `TLS_AUTO`, `TLSv1_0`, `TLSv1_1`, `TLSv1_2`, `TLSv1_3`<br>TLS Protocol. TlsProtocol is enumeration of supported TLS versions F5 Distributed Cloud will choose the optimal TLS version

#### HTTPS TLS Parameters Use mTLS

An [`use_mtls`](#https-tls-parameters-use-mtls) block (within [`https.tls_parameters`](#https-tls-parameters)) supports the following:

&#x2022; [`client_certificate_optional`](#client-certificate-optional) - Optional Bool<br>Client Certificate Optional. Client certificate is optional. If the client has provided a certificate, the load balancer will verify it. If certification verification fails, the connection will be terminated. If the client does not provide a certificate, the connection will be accepted

&#x2022; [`crl`](#crl) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [CRL](#https-tls-parameters-use-mtls-crl) below.

&#x2022; [`no_crl`](#no-crl) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`trusted_ca`](#trusted-ca) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Trusted CA](#https-tls-parameters-use-mtls-trusted-ca) below.

&#x2022; [`trusted_ca_url`](#trusted-ca-url) - Optional String<br>Inline Root CA Certificate (legacy). Upload a Root CA Certificate specifically for this Load Balancer

&#x2022; [`xfcc_disabled`](#xfcc-disabled) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`xfcc_options`](#xfcc-options) - Optional Block<br>XFCC Header Elements. X-Forwarded-Client-Cert header elements to be added to requests<br>See [Xfcc Options](#https-tls-parameters-use-mtls-xfcc-options) below.

#### HTTPS TLS Parameters Use mTLS CRL

A [`crl`](#https-tls-parameters-use-mtls-crl) block (within [`https.tls_parameters.use_mtls`](#https-tls-parameters-use-mtls)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### HTTPS TLS Parameters Use mTLS Trusted CA

A [`trusted_ca`](#https-tls-parameters-use-mtls-trusted-ca) block (within [`https.tls_parameters.use_mtls`](#https-tls-parameters-use-mtls)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### HTTPS TLS Parameters Use mTLS Xfcc Options

A [`xfcc_options`](#https-tls-parameters-use-mtls-xfcc-options) block (within [`https.tls_parameters.use_mtls`](#https-tls-parameters-use-mtls)) supports the following:

&#x2022; [`xfcc_header_elements`](#xfcc-header-elements) - Optional List  Defaults to `XFCC_NONE`<br>Possible values are `XFCC_NONE`, `XFCC_CERT`, `XFCC_CHAIN`, `XFCC_SUBJECT`, `XFCC_URI`, `XFCC_DNS`<br>XFCC Header Elements. X-Forwarded-Client-Cert header elements to be added to requests

#### HTTPS Auto Cert

A [`https_auto_cert`](#https-auto-cert) block supports the following:

&#x2022; [`add_hsts`](#add-hsts) - Optional Bool<br>Add HSTS Header. Add HTTP Strict-Transport-Security response header

&#x2022; [`append_server_name`](#append-server-name) - Optional String<br>Append header value. Define the header value for the header name “server”. If header value is already present, it is not overwritten and passed as-is

&#x2022; [`coalescing_options`](#coalescing-options) - Optional Block<br>TLS Coalescing Options. TLS connection coalescing configuration (not compatible with mTLS)<br>See [Coalescing Options](#https-auto-cert-coalescing-options) below.

&#x2022; [`connection_idle_timeout`](#connection-idle-timeout) - Optional Number  Defaults to `2`  Specified in milliseconds<br>Connection Idle Timeout. The idle timeout for downstream connections. The idle timeout is defined as the period in which there are no active requests. When the idle timeout is reached the connection will be closed. Note that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive.  The minutes

&#x2022; [`default_header`](#default-header) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`default_loadbalancer`](#default-loadbalancer) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`disable_path_normalize`](#disable-path-normalize) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`enable_path_normalize`](#enable-path-normalize) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`http_protocol_options`](#http-protocol-options) - Optional Block<br>HTTP Protocol Configuration Options. HTTP protocol configuration options for downstream connections<br>See [HTTP Protocol Options](#https-auto-cert-http-protocol-options) below.

&#x2022; [`http_redirect`](#http-redirect) - Optional Bool<br>HTTP Redirect to HTTPS. Redirect HTTP traffic to HTTPS

&#x2022; [`no_mtls`](#no-mtls) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`non_default_loadbalancer`](#non-default-loadbalancer) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`pass_through`](#pass-through) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`port`](#port) - Optional Number<br>HTTPS Listen Port. HTTPS port to Listen

&#x2022; [`port_ranges`](#port-ranges) - Optional String<br>Port Ranges. A string containing a comma separated list of port ranges. Each port range consists of a single port or two ports separated by '-'

&#x2022; [`server_name`](#server-name) - Optional String<br>Modify header value. Define the header value for the header name “server”. This will overwrite existing values, if any, for the server header

&#x2022; [`tls_config`](#tls-config) - Optional Block<br>TLS Config. This defines various options to configure TLS configuration parameters<br>See [TLS Config](#https-auto-cert-tls-config) below.

&#x2022; [`use_mtls`](#use-mtls) - Optional Block<br>Clients TLS validation context. Validation context for downstream client TLS connections<br>See [Use mTLS](#https-auto-cert-use-mtls) below.

#### HTTPS Auto Cert Coalescing Options

A [`coalescing_options`](#https-auto-cert-coalescing-options) block (within [`https_auto_cert`](#https-auto-cert)) supports the following:

&#x2022; [`default_coalescing`](#default-coalescing) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`strict_coalescing`](#strict-coalescing) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### HTTPS Auto Cert HTTP Protocol Options

A [`http_protocol_options`](#https-auto-cert-http-protocol-options) block (within [`https_auto_cert`](#https-auto-cert)) supports the following:

&#x2022; [`http_protocol_enable_v1_only`](#http-protocol-enable-v1-only) - Optional Block<br>HTTP/1.1 Protocol Options. HTTP/1.1 Protocol options for downstream connections<br>See [HTTP Protocol Enable V1 Only](#https-auto-cert-http-protocol-options-http-protocol-enable-v1-only) below.

&#x2022; [`http_protocol_enable_v1_v2`](#http-protocol-enable-v1-v2) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`http_protocol_enable_v2_only`](#http-protocol-enable-v2-only) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### HTTPS Auto Cert HTTP Protocol Options HTTP Protocol Enable V1 Only

A [`http_protocol_enable_v1_only`](#https-auto-cert-http-protocol-options-http-protocol-enable-v1-only) block (within [`https_auto_cert.http_protocol_options`](#https-auto-cert-http-protocol-options)) supports the following:

&#x2022; [`header_transformation`](#header-transformation) - Optional Block<br>Header Transformation. Header Transformation options for HTTP/1.1 request/response headers<br>See [Header Transformation](#https-auto-cert-http-protocol-options-http-protocol-enable-v1-only-header-transformation) below.

#### HTTPS Auto Cert HTTP Protocol Options HTTP Protocol Enable V1 Only Header Transformation

A [`header_transformation`](#https-auto-cert-http-protocol-options-http-protocol-enable-v1-only-header-transformation) block (within [`https_auto_cert.http_protocol_options.http_protocol_enable_v1_only`](#https-auto-cert-http-protocol-options-http-protocol-enable-v1-only)) supports the following:

&#x2022; [`default_header_transformation`](#default-header-transformation) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`legacy_header_transformation`](#legacy-header-transformation) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`preserve_case_header_transformation`](#preserve-case-header-transformation) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`proper_case_header_transformation`](#proper-case-header-transformation) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### HTTPS Auto Cert TLS Config

A [`tls_config`](#https-auto-cert-tls-config) block (within [`https_auto_cert`](#https-auto-cert)) supports the following:

&#x2022; [`custom_security`](#custom-security) - Optional Block<br>Custom Ciphers. This defines TLS protocol config including min/max versions and allowed ciphers<br>See [Custom Security](#https-auto-cert-tls-config-custom-security) below.

&#x2022; [`default_security`](#default-security) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`low_security`](#low-security) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`medium_security`](#medium-security) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### HTTPS Auto Cert TLS Config Custom Security

A [`custom_security`](#https-auto-cert-tls-config-custom-security) block (within [`https_auto_cert.tls_config`](#https-auto-cert-tls-config)) supports the following:

&#x2022; [`cipher_suites`](#cipher-suites) - Optional List<br>Cipher Suites. The TLS listener will only support the specified cipher list

&#x2022; [`max_version`](#max-version) - Optional String  Defaults to `TLS_AUTO`<br>Possible values are `TLS_AUTO`, `TLSv1_0`, `TLSv1_1`, `TLSv1_2`, `TLSv1_3`<br>TLS Protocol. TlsProtocol is enumeration of supported TLS versions F5 Distributed Cloud will choose the optimal TLS version

&#x2022; [`min_version`](#min-version) - Optional String  Defaults to `TLS_AUTO`<br>Possible values are `TLS_AUTO`, `TLSv1_0`, `TLSv1_1`, `TLSv1_2`, `TLSv1_3`<br>TLS Protocol. TlsProtocol is enumeration of supported TLS versions F5 Distributed Cloud will choose the optimal TLS version

#### HTTPS Auto Cert Use mTLS

An [`use_mtls`](#https-auto-cert-use-mtls) block (within [`https_auto_cert`](#https-auto-cert)) supports the following:

&#x2022; [`client_certificate_optional`](#client-certificate-optional) - Optional Bool<br>Client Certificate Optional. Client certificate is optional. If the client has provided a certificate, the load balancer will verify it. If certification verification fails, the connection will be terminated. If the client does not provide a certificate, the connection will be accepted

&#x2022; [`crl`](#crl) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [CRL](#https-auto-cert-use-mtls-crl) below.

&#x2022; [`no_crl`](#no-crl) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`trusted_ca`](#trusted-ca) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Trusted CA](#https-auto-cert-use-mtls-trusted-ca) below.

&#x2022; [`trusted_ca_url`](#trusted-ca-url) - Optional String<br>Inline Root CA Certificate (legacy). Upload a Root CA Certificate specifically for this Load Balancer

&#x2022; [`xfcc_disabled`](#xfcc-disabled) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`xfcc_options`](#xfcc-options) - Optional Block<br>XFCC Header Elements. X-Forwarded-Client-Cert header elements to be added to requests<br>See [Xfcc Options](#https-auto-cert-use-mtls-xfcc-options) below.

#### HTTPS Auto Cert Use mTLS CRL

A [`crl`](#https-auto-cert-use-mtls-crl) block (within [`https_auto_cert.use_mtls`](#https-auto-cert-use-mtls)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### HTTPS Auto Cert Use mTLS Trusted CA

A [`trusted_ca`](#https-auto-cert-use-mtls-trusted-ca) block (within [`https_auto_cert.use_mtls`](#https-auto-cert-use-mtls)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### HTTPS Auto Cert Use mTLS Xfcc Options

A [`xfcc_options`](#https-auto-cert-use-mtls-xfcc-options) block (within [`https_auto_cert.use_mtls`](#https-auto-cert-use-mtls)) supports the following:

&#x2022; [`xfcc_header_elements`](#xfcc-header-elements) - Optional List  Defaults to `XFCC_NONE`<br>Possible values are `XFCC_NONE`, `XFCC_CERT`, `XFCC_CHAIN`, `XFCC_SUBJECT`, `XFCC_URI`, `XFCC_DNS`<br>XFCC Header Elements. X-Forwarded-Client-Cert header elements to be added to requests

#### Js Challenge

A [`js_challenge`](#js-challenge) block supports the following:

&#x2022; [`cookie_expiry`](#cookie-expiry) - Optional Number<br>Cookie Expiration Period. Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new challenge

&#x2022; [`custom_page`](#custom-page) - Optional String<br>Custom Message for Javascript Challenge. Custom message is of type uri_ref. Currently supported URL schemes is string:///. For string:/// scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. 'Please Wait.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '<p> Please Wait </p>'. Base64 encoded string for this HTML is 'PHA+IFBsZWFzZSBXYWl0IDwvcD4='

&#x2022; [`js_script_delay`](#js-script-delay) - Optional Number<br>Javascript Delay. Delay introduced by Javascript, in milliseconds

#### JWT Validation

A [`jwt_validation`](#jwt-validation) block supports the following:

&#x2022; [`action`](#action) - Optional Block<br>Action<br>See [Action](#jwt-validation-action) below.

&#x2022; [`jwks_config`](#jwks-config) - Optional Block<br>JSON Web Key Set (JWKS). The JSON Web Key Set (JWKS) is a set of keys used to verify JSON Web Token (JWT) issued by the Authorization Server. See RFC 7517 for more details<br>See [Jwks Config](#jwt-validation-jwks-config) below.

&#x2022; [`mandatory_claims`](#mandatory-claims) - Optional Block<br>Mandatory Claims. Configurable Validation of mandatory Claims<br>See [Mandatory Claims](#jwt-validation-mandatory-claims) below.

&#x2022; [`reserved_claims`](#reserved-claims) - Optional Block<br>Reserved claims configuration. Configurable Validation of reserved Claims<br>See [Reserved Claims](#jwt-validation-reserved-claims) below.

&#x2022; [`target`](#target) - Optional Block<br>Target. Define endpoints for which JWT token validation will be performed<br>See [Target](#jwt-validation-target) below.

&#x2022; [`token_location`](#token-location) - Optional Block<br>Token Location. Location of JWT in HTTP request<br>See [Token Location](#jwt-validation-token-location) below.

#### JWT Validation Action

An [`action`](#jwt-validation-action) block (within [`jwt_validation`](#jwt-validation)) supports the following:

&#x2022; [`block`](#block) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`report`](#report) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### JWT Validation Jwks Config

A [`jwks_config`](#jwt-validation-jwks-config) block (within [`jwt_validation`](#jwt-validation)) supports the following:

&#x2022; [`cleartext`](#cleartext) - Optional String<br>JSON Web Key Set (JWKS). The JSON Web Key Set (JWKS) is a set of keys used to verify JSON Web Token (JWT) issued by the Authorization Server. See RFC 7517 for more details

#### JWT Validation Mandatory Claims

A [`mandatory_claims`](#jwt-validation-mandatory-claims) block (within [`jwt_validation`](#jwt-validation)) supports the following:

&#x2022; [`claim_names`](#claim-names) - Optional List<br>Claim Names

#### JWT Validation Reserved Claims

A [`reserved_claims`](#jwt-validation-reserved-claims) block (within [`jwt_validation`](#jwt-validation)) supports the following:

&#x2022; [`audience`](#audience) - Optional Block<br>Audiences<br>See [Audience](#jwt-validation-reserved-claims-audience) below.

&#x2022; [`audience_disable`](#audience-disable) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`issuer`](#issuer) - Optional String<br>Exact Match

&#x2022; [`issuer_disable`](#issuer-disable) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`validate_period_disable`](#validate-period-disable) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`validate_period_enable`](#validate-period-enable) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### JWT Validation Reserved Claims Audience

An [`audience`](#jwt-validation-reserved-claims-audience) block (within [`jwt_validation.reserved_claims`](#jwt-validation-reserved-claims)) supports the following:

&#x2022; [`audiences`](#audiences) - Optional List<br>Values

#### JWT Validation Target

A [`target`](#jwt-validation-target) block (within [`jwt_validation`](#jwt-validation)) supports the following:

&#x2022; [`all_endpoint`](#all-endpoint) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`api_groups`](#api-groups) - Optional Block<br>API Groups<br>See [API Groups](#jwt-validation-target-api-groups) below.

&#x2022; [`base_paths`](#base-paths) - Optional Block<br>Base Paths<br>See [Base Paths](#jwt-validation-target-base-paths) below.

#### JWT Validation Target API Groups

An [`api_groups`](#jwt-validation-target-api-groups) block (within [`jwt_validation.target`](#jwt-validation-target)) supports the following:

&#x2022; [`api_groups`](#api-groups) - Optional List<br>API Groups

#### JWT Validation Target Base Paths

A [`base_paths`](#jwt-validation-target-base-paths) block (within [`jwt_validation.target`](#jwt-validation-target)) supports the following:

&#x2022; [`base_paths`](#base-paths) - Optional List<br>Prefix Values

#### JWT Validation Token Location

A [`token_location`](#jwt-validation-token-location) block (within [`jwt_validation`](#jwt-validation)) supports the following:

&#x2022; [`bearer_token`](#bearer-token) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### L7 DDOS Action Js Challenge

A [`l7_ddos_action_js_challenge`](#l7-ddos-action-js-challenge) block supports the following:

&#x2022; [`cookie_expiry`](#cookie-expiry) - Optional Number<br>Cookie Expiration Period. Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new challenge

&#x2022; [`custom_page`](#custom-page) - Optional String<br>Custom Message for Javascript Challenge. Custom message is of type uri_ref. Currently supported URL schemes is string:///. For string:/// scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. 'Please Wait.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '<p> Please Wait </p>'. Base64 encoded string for this HTML is 'PHA+IFBsZWFzZSBXYWl0IDwvcD4='

&#x2022; [`js_script_delay`](#js-script-delay) - Optional Number<br>Javascript Delay. Delay introduced by Javascript, in milliseconds

#### L7 DDOS Protection

A [`l7_ddos_protection`](#l7-ddos-protection) block supports the following:

&#x2022; [`clientside_action_captcha_challenge`](#clientside-action-captcha-challenge) - Optional Block<br>Captcha Challenge Parameters. Enables loadbalancer to perform captcha challenge Captcha challenge will be based on Google Recaptcha. With this feature enabled, only clients that pass the captcha challenge will be allowed to complete the HTTP request. When loadbalancer is configured to do Captcha Challenge, it will redirect the browser to an HTML page on every new HTTP request. This HTML page will have captcha challenge embedded in it. Client will be allowed to make the request only if the captcha challenge is successful. Loadbalancer will tag response header with a cookie to avoid Captcha challenge for subsequent requests. CAPTCHA is mainly used as a security check to ensure only human users can pass through. Generally, computers or bots are not capable of solving a captcha. You can enable either Javascript challenge or Captcha challenge on a virtual host<br>See [Clientside Action Captcha Challenge](#l7-ddos-protection-clientside-action-captcha-challenge) below.

&#x2022; [`clientside_action_js_challenge`](#clientside-action-js-challenge) - Optional Block<br>Javascript Challenge Parameters. Enables loadbalancer to perform client browser compatibility test by redirecting to a page with Javascript. With this feature enabled, only clients that are capable of executing Javascript(mostly browsers) will be allowed to complete the HTTP request. When loadbalancer is configured to do Javascript Challenge, it will redirect the browser to an HTML page on every new HTTP request. This HTML page will have Javascript embedded in it. Loadbalancer chooses a set of random numbers for every new client and sends these numbers along with an encrypted answer with the request such that it embed these numbers as input in the Javascript. Javascript will run on the requestor browser and perform a complex Math operation. Script will submit the answer to loadbalancer. Loadbalancer will validate the answer by comparing the calculated answer with the decrypted answer (which was encrypted when it was sent back as reply) and allow the request to the upstream server only if the answer is correct. Loadbalancer will tag response header with a cookie to avoid Javascript challenge for subsequent requests. Javascript challenge serves following purposes * Validate that the request is coming via a browser that is capable for running Javascript * Force the browser to run a complex operation, f(X), that requires it to spend a large number of CPU cycles. This is to slow down a potential DOS attacker by making it difficult to launch a large request flood without having to spend even larger CPU cost at their end. You can enable either Javascript challenge or Captcha challenge on a virtual host<br>See [Clientside Action Js Challenge](#l7-ddos-protection-clientside-action-js-challenge) below.

&#x2022; [`clientside_action_none`](#clientside-action-none) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`ddos_policy_custom`](#ddos-policy-custom) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [DDOS Policy Custom](#l7-ddos-protection-ddos-policy-custom) below.

&#x2022; [`ddos_policy_none`](#ddos-policy-none) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`default_rps_threshold`](#default-rps-threshold) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`mitigation_block`](#mitigation-block) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`mitigation_captcha_challenge`](#mitigation-captcha-challenge) - Optional Block<br>Captcha Challenge Parameters. Enables loadbalancer to perform captcha challenge Captcha challenge will be based on Google Recaptcha. With this feature enabled, only clients that pass the captcha challenge will be allowed to complete the HTTP request. When loadbalancer is configured to do Captcha Challenge, it will redirect the browser to an HTML page on every new HTTP request. This HTML page will have captcha challenge embedded in it. Client will be allowed to make the request only if the captcha challenge is successful. Loadbalancer will tag response header with a cookie to avoid Captcha challenge for subsequent requests. CAPTCHA is mainly used as a security check to ensure only human users can pass through. Generally, computers or bots are not capable of solving a captcha. You can enable either Javascript challenge or Captcha challenge on a virtual host<br>See [Mitigation Captcha Challenge](#l7-ddos-protection-mitigation-captcha-challenge) below.

&#x2022; [`mitigation_js_challenge`](#mitigation-js-challenge) - Optional Block<br>Javascript Challenge Parameters. Enables loadbalancer to perform client browser compatibility test by redirecting to a page with Javascript. With this feature enabled, only clients that are capable of executing Javascript(mostly browsers) will be allowed to complete the HTTP request. When loadbalancer is configured to do Javascript Challenge, it will redirect the browser to an HTML page on every new HTTP request. This HTML page will have Javascript embedded in it. Loadbalancer chooses a set of random numbers for every new client and sends these numbers along with an encrypted answer with the request such that it embed these numbers as input in the Javascript. Javascript will run on the requestor browser and perform a complex Math operation. Script will submit the answer to loadbalancer. Loadbalancer will validate the answer by comparing the calculated answer with the decrypted answer (which was encrypted when it was sent back as reply) and allow the request to the upstream server only if the answer is correct. Loadbalancer will tag response header with a cookie to avoid Javascript challenge for subsequent requests. Javascript challenge serves following purposes * Validate that the request is coming via a browser that is capable for running Javascript * Force the browser to run a complex operation, f(X), that requires it to spend a large number of CPU cycles. This is to slow down a potential DOS attacker by making it difficult to launch a large request flood without having to spend even larger CPU cost at their end. You can enable either Javascript challenge or Captcha challenge on a virtual host<br>See [Mitigation Js Challenge](#l7-ddos-protection-mitigation-js-challenge) below.

&#x2022; [`rps_threshold`](#rps-threshold) - Optional Number<br>Custom. Configure custom RPS threshold

#### L7 DDOS Protection Clientside Action Captcha Challenge

A [`clientside_action_captcha_challenge`](#l7-ddos-protection-clientside-action-captcha-challenge) block (within [`l7_ddos_protection`](#l7-ddos-protection)) supports the following:

&#x2022; [`cookie_expiry`](#cookie-expiry) - Optional Number<br>Cookie Expiration Period. Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new challenge

&#x2022; [`custom_page`](#custom-page) - Optional String<br>Custom message for Captcha Challenge. Custom message is of type uri_ref. Currently supported URL schemes is string:///. For string:/// scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. 'Please Wait.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '<p> Please Wait </p>'. Base64 encoded string for this HTML is 'PHA+IFBsZWFzZSBXYWl0IDwvcD4='

#### L7 DDOS Protection Clientside Action Js Challenge

A [`clientside_action_js_challenge`](#l7-ddos-protection-clientside-action-js-challenge) block (within [`l7_ddos_protection`](#l7-ddos-protection)) supports the following:

&#x2022; [`cookie_expiry`](#cookie-expiry) - Optional Number<br>Cookie Expiration Period. Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new challenge

&#x2022; [`custom_page`](#custom-page) - Optional String<br>Custom Message for Javascript Challenge. Custom message is of type uri_ref. Currently supported URL schemes is string:///. For string:/// scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. 'Please Wait.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '<p> Please Wait </p>'. Base64 encoded string for this HTML is 'PHA+IFBsZWFzZSBXYWl0IDwvcD4='

&#x2022; [`js_script_delay`](#js-script-delay) - Optional Number<br>Javascript Delay. Delay introduced by Javascript, in milliseconds

#### L7 DDOS Protection DDOS Policy Custom

A [`ddos_policy_custom`](#l7-ddos-protection-ddos-policy-custom) block (within [`l7_ddos_protection`](#l7-ddos-protection)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### L7 DDOS Protection Mitigation Captcha Challenge

A [`mitigation_captcha_challenge`](#l7-ddos-protection-mitigation-captcha-challenge) block (within [`l7_ddos_protection`](#l7-ddos-protection)) supports the following:

&#x2022; [`cookie_expiry`](#cookie-expiry) - Optional Number<br>Cookie Expiration Period. Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new challenge

&#x2022; [`custom_page`](#custom-page) - Optional String<br>Custom message for Captcha Challenge. Custom message is of type uri_ref. Currently supported URL schemes is string:///. For string:/// scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. 'Please Wait.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '<p> Please Wait </p>'. Base64 encoded string for this HTML is 'PHA+IFBsZWFzZSBXYWl0IDwvcD4='

#### L7 DDOS Protection Mitigation Js Challenge

A [`mitigation_js_challenge`](#l7-ddos-protection-mitigation-js-challenge) block (within [`l7_ddos_protection`](#l7-ddos-protection)) supports the following:

&#x2022; [`cookie_expiry`](#cookie-expiry) - Optional Number<br>Cookie Expiration Period. Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new challenge

&#x2022; [`custom_page`](#custom-page) - Optional String<br>Custom Message for Javascript Challenge. Custom message is of type uri_ref. Currently supported URL schemes is string:///. For string:/// scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. 'Please Wait.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '<p> Please Wait </p>'. Base64 encoded string for this HTML is 'PHA+IFBsZWFzZSBXYWl0IDwvcD4='

&#x2022; [`js_script_delay`](#js-script-delay) - Optional Number<br>Javascript Delay. Delay introduced by Javascript, in milliseconds

#### Malware Protection Settings

A [`malware_protection_settings`](#malware-protection-settings) block supports the following:

&#x2022; [`malware_protection_rules`](#malware-protection-rules) - Optional Block<br>Malware Detection Rules. Configure the match criteria to trigger Malware Protection Scan<br>See [Malware Protection Rules](#malware-protection-settings-malware-protection-rules) below.

#### Malware Protection Settings Malware Protection Rules

A [`malware_protection_rules`](#malware-protection-settings-malware-protection-rules) block (within [`malware_protection_settings`](#malware-protection-settings)) supports the following:

&#x2022; [`action`](#action) - Optional Block<br>Action<br>See [Action](#malware-protection-settings-malware-protection-rules-action) below.

&#x2022; [`domain`](#domain) - Optional Block<br>Domain to Match. Domain to be matched<br>See [Domain](#malware-protection-settings-malware-protection-rules-domain) below.

&#x2022; [`http_methods`](#http-methods) - Optional List  Defaults to `ANY`<br>Possible values are `ANY`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`, `COPY`<br>HTTP Methods. Methods to be matched

&#x2022; [`metadata`](#metadata) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#malware-protection-settings-malware-protection-rules-metadata) below.

&#x2022; [`path`](#path) - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#malware-protection-settings-malware-protection-rules-path) below.

#### Malware Protection Settings Malware Protection Rules Action

An [`action`](#malware-protection-settings-malware-protection-rules-action) block (within [`malware_protection_settings.malware_protection_rules`](#malware-protection-settings-malware-protection-rules)) supports the following:

&#x2022; [`block`](#block) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`report`](#report) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### Malware Protection Settings Malware Protection Rules Domain

A [`domain`](#malware-protection-settings-malware-protection-rules-domain) block (within [`malware_protection_settings.malware_protection_rules`](#malware-protection-settings-malware-protection-rules)) supports the following:

&#x2022; [`any_domain`](#any-domain) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`domain`](#domain) - Optional Block<br>Domains. Domains names<br>See [Domain](#malware-protection-settings-malware-protection-rules-domain-domain) below.

#### Malware Protection Settings Malware Protection Rules Domain Domain

A [`domain`](#malware-protection-settings-malware-protection-rules-domain-domain) block (within [`malware_protection_settings.malware_protection_rules.domain`](#malware-protection-settings-malware-protection-rules-domain)) supports the following:

&#x2022; [`exact_value`](#exact-value) - Optional String<br>Exact Value. Exact domain name

&#x2022; [`regex_value`](#regex-value) - Optional String<br>Regex Values of Domains. Regular Expression value for the domain name

&#x2022; [`suffix_value`](#suffix-value) - Optional String<br>Suffix Value. Suffix of domain name e.g 'xyz.com' will match '*.xyz.com' and 'xyz.com'

#### Malware Protection Settings Malware Protection Rules Metadata

A [`metadata`](#malware-protection-settings-malware-protection-rules-metadata) block (within [`malware_protection_settings.malware_protection_rules`](#malware-protection-settings-malware-protection-rules)) supports the following:

&#x2022; [`description`](#description) - Optional String<br>Description. Human readable description

&#x2022; [`name`](#name) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### Malware Protection Settings Malware Protection Rules Path

A [`path`](#malware-protection-settings-malware-protection-rules-path) block (within [`malware_protection_settings.malware_protection_rules`](#malware-protection-settings-malware-protection-rules)) supports the following:

&#x2022; [`path`](#path) - Optional String<br>Exact. Exact path value to match

&#x2022; [`prefix`](#prefix) - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

&#x2022; [`regex`](#regex) - Optional String<br>Regex. Regular expression of path match (e.g. the value .* will match on all paths)

#### More Option

A [`more_option`](#more-option) block supports the following:

&#x2022; [`buffer_policy`](#buffer-policy) - Optional Block<br>Buffer Configuration. Some upstream applications are not capable of handling streamed data. This config enables buffering the entire request before sending to upstream application. We can specify the maximum buffer size and buffer interval with this config. Buffering can be enabled and disabled at VirtualHost and Route levels Route level buffer configuration takes precedence<br>See [Buffer Policy](#more-option-buffer-policy) below.

&#x2022; [`compression_params`](#compression-params) - Optional Block<br>Compression Parameters. Enables loadbalancer to compress dispatched data from an upstream service upon client request. The content is compressed and then sent to the client with the appropriate headers if either response and request allow. Only GZIP compression is supported. By default compression will be skipped when: A request does NOT contain accept-encoding header. A request includes accept-encoding header, but it does not contain “gzip” or “*”. A request includes accept-encoding with “gzip” or “*” with the weight “q=0”. Note that the “gzip” will have a higher weight then “*”. For example, if accept-encoding is “gzip;q=0,*;q=1”, the filter will not compress. But if the header is set to “*;q=0,gzip;q=1”, the filter will compress. A request whose accept-encoding header includes “identity”. A response contains a content-encoding header. A response contains a cache-control header whose value includes “no-transform”. A response contains a transfer-encoding header whose value includes “gzip”. A response does not contain a content-type value that matches one of the selected mime-types, which default to application/javascript, application/JSON, application/xhtml+XML, image/svg+XML, text/CSS, text/HTML, text/plain, text/XML. Neither content-length nor transfer-encoding headers are present in the response. Response size is smaller than 30 bytes (only applicable when transfer-encoding is not chunked). When compression is applied: The content-length is removed from response headers. Response headers contain “transfer-encoding: chunked” and do not contain “content-encoding” header. The “vary: accept-encoding” header is inserted on every response. GZIP Compression Level: A value which is optimal balance between speed of compression and amount of compression is chosen<br>See [Compression Params](#more-option-compression-params) below.

&#x2022; [`custom_errors`](#custom-errors) - Optional Block<br>Custom Error Responses. Map of integer error codes as keys and string values that can be used to provide custom HTTP pages for each error code. Key of the map can be either response code class or HTTP Error code. Response code classes for key is configured as follows 3 -- for 3xx response code class 4 -- for 4xx response code class 5 -- for 5xx response code class Value of the map is string which represents custom HTTP responses. Specific response code takes preference when both response code and response code class matches for a request

&#x2022; [`disable_default_error_pages`](#disable-default-error-pages) - Optional Bool<br>Disable Default Error Pages. Disable the use of default F5XC error pages

&#x2022; [`disable_path_normalize`](#disable-path-normalize) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`enable_path_normalize`](#enable-path-normalize) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`idle_timeout`](#idle-timeout) - Optional Number<br>Idle Timeout. The amount of time that a stream can exist without upstream or downstream activity, in milliseconds. The stream is terminated with a HTTP 504 (Gateway Timeout) error code if no upstream response header has been received, otherwise the stream is reset

&#x2022; [`max_request_header_size`](#max-request-header-size) - Optional Number<br>Maximum Request Header Size. The maximum request header size for downstream connections, in KiB. A HTTP 431 (Request Header Fields Too Large) error code is sent for requests that exceed this size. If multiple load balancers share the same advertise_policy, the highest value configured across all such load balancers is used for all the load balancers in question

&#x2022; [`request_cookies_to_add`](#request-cookies-to-add) - Optional Block<br>Add Cookies in Cookie Header. Cookies are key-value pairs to be added to HTTP request being routed towards upstream. Cookies specified at this level are applied after cookies from matched Route are applied<br>See [Request Cookies To Add](#more-option-request-cookies-to-add) below.

&#x2022; [`request_cookies_to_remove`](#request-cookies-to-remove) - Optional List<br>Remove Cookies from Cookie Header. List of keys of Cookies to be removed from the HTTP request being sent towards upstream

&#x2022; [`request_headers_to_add`](#request-headers-to-add) - Optional Block<br>Add Request Headers. Headers are key-value pairs to be added to HTTP request being routed towards upstream. Headers specified at this level are applied after headers from matched Route are applied<br>See [Request Headers To Add](#more-option-request-headers-to-add) below.

&#x2022; [`request_headers_to_remove`](#request-headers-to-remove) - Optional List<br>Remove Request Headers. List of keys of Headers to be removed from the HTTP request being sent towards upstream

&#x2022; [`response_cookies_to_add`](#response-cookies-to-add) - Optional Block<br>Add Set-Cookie Headers. Cookies are name-value pairs along with optional attribute parameters to be added to HTTP response being sent towards downstream. Cookies specified at this level are applied after cookies from matched Route are applied<br>See [Response Cookies To Add](#more-option-response-cookies-to-add) below.

&#x2022; [`response_cookies_to_remove`](#response-cookies-to-remove) - Optional List<br>Remove Cookies from Set-Cookie Headers. List of name of Cookies to be removed from the HTTP response being sent towards downstream. Entire set-cookie header will be removed

&#x2022; [`response_headers_to_add`](#response-headers-to-add) - Optional Block<br>Add Response Headers. Headers are key-value pairs to be added to HTTP response being sent towards downstream. Headers specified at this level are applied after headers from matched Route are applied<br>See [Response Headers To Add](#more-option-response-headers-to-add) below.

&#x2022; [`response_headers_to_remove`](#response-headers-to-remove) - Optional List<br>Remove Response Headers. List of keys of Headers to be removed from the HTTP response being sent towards downstream

#### More Option Buffer Policy

A [`buffer_policy`](#more-option-buffer-policy) block (within [`more_option`](#more-option)) supports the following:

&#x2022; [`disabled`](#disabled) - Optional Bool<br>Disable. Disable buffering for a particular route. This is useful when virtual-host has buffering, but we need to disable it on a specific route. The value of this field is ignored for virtual-host

&#x2022; [`max_request_bytes`](#max-request-bytes) - Optional Number<br>Max Request Bytes. The maximum request size that the filter will buffer before the connection manager will stop buffering and return a RequestEntityTooLarge (413) response

#### More Option Compression Params

A [`compression_params`](#more-option-compression-params) block (within [`more_option`](#more-option)) supports the following:

&#x2022; [`content_length`](#content-length) - Optional Number  Defaults to `30`<br>Content Length. Minimum response length, in bytes, which will trigger compression. The

&#x2022; [`content_type`](#content-type) - Optional List<br>Content Type. Set of strings that allows specifying which mime-types yield compression When this field is not defined, compression will be applied to the following mime-types: 'application/javascript' 'application/JSON', 'application/xhtml+XML' 'image/svg+XML' 'text/CSS' 'text/HTML' 'text/plain' 'text/XML'

&#x2022; [`disable_on_etag_header`](#disable-on-etag-header) - Optional Bool<br>Disable On Etag Header. If true, disables compression when the response contains an etag header. When it is false, weak etags will be preserved and the ones that require strong validation will be removed

&#x2022; [`remove_accept_encoding_header`](#remove-accept-encoding-header) - Optional Bool<br>Remove Accept-Encoding Header. If true, removes accept-encoding from the request headers before dispatching it to the upstream so that responses do not get compressed before reaching the filter

#### More Option Request Cookies To Add

A [`request_cookies_to_add`](#more-option-request-cookies-to-add) block (within [`more_option`](#more-option)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. Name of the cookie in Cookie header

&#x2022; [`overwrite`](#overwrite) - Optional Bool  Defaults to `do`<br>Overwrite. Should the value be overwritten? If true, the value is overwritten to existing values. not overwrite

&#x2022; [`secret_value`](#secret-value) - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Secret Value](#more-option-request-cookies-to-add-secret-value) below.

&#x2022; [`value`](#value) - Optional String<br>Value. Value of the Cookie header

#### More Option Request Cookies To Add Secret Value

A [`secret_value`](#more-option-request-cookies-to-add-secret-value) block (within [`more_option.request_cookies_to_add`](#more-option-request-cookies-to-add)) supports the following:

&#x2022; [`blindfold_secret_info`](#blindfold-secret-info) - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#more-option-request-cookies-to-add-secret-value-blindfold-secret-info) below.

&#x2022; [`clear_secret_info`](#clear-secret-info) - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#more-option-request-cookies-to-add-secret-value-clear-secret-info) below.

#### More Option Request Cookies To Add Secret Value Blindfold Secret Info

A [`blindfold_secret_info`](#more-option-request-cookies-to-add-secret-value-blindfold-secret-info) block (within [`more_option.request_cookies_to_add.secret_value`](#more-option-request-cookies-to-add-secret-value)) supports the following:

&#x2022; [`decryption_provider`](#decryption-provider) - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

&#x2022; [`location`](#location) - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

&#x2022; [`store_provider`](#store-provider) - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

#### More Option Request Cookies To Add Secret Value Clear Secret Info

A [`clear_secret_info`](#more-option-request-cookies-to-add-secret-value-clear-secret-info) block (within [`more_option.request_cookies_to_add.secret_value`](#more-option-request-cookies-to-add-secret-value)) supports the following:

&#x2022; [`provider_ref`](#provider-ref) - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

&#x2022; [`url`](#url) - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

#### More Option Request Headers To Add

A [`request_headers_to_add`](#more-option-request-headers-to-add) block (within [`more_option`](#more-option)) supports the following:

&#x2022; [`append`](#append) - Optional Bool  Defaults to `do`<br>Append. Should the value be appended? If true, the value is appended to existing values. not append

&#x2022; [`name`](#name) - Optional String<br>Name. Name of the HTTP header

&#x2022; [`secret_value`](#secret-value) - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Secret Value](#more-option-request-headers-to-add-secret-value) below.

&#x2022; [`value`](#value) - Optional String<br>Value. Value of the HTTP header

#### More Option Request Headers To Add Secret Value

A [`secret_value`](#more-option-request-headers-to-add-secret-value) block (within [`more_option.request_headers_to_add`](#more-option-request-headers-to-add)) supports the following:

&#x2022; [`blindfold_secret_info`](#blindfold-secret-info) - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#more-option-request-headers-to-add-secret-value-blindfold-secret-info) below.

&#x2022; [`clear_secret_info`](#clear-secret-info) - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#more-option-request-headers-to-add-secret-value-clear-secret-info) below.

#### More Option Request Headers To Add Secret Value Blindfold Secret Info

A [`blindfold_secret_info`](#more-option-request-headers-to-add-secret-value-blindfold-secret-info) block (within [`more_option.request_headers_to_add.secret_value`](#more-option-request-headers-to-add-secret-value)) supports the following:

&#x2022; [`decryption_provider`](#decryption-provider) - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

&#x2022; [`location`](#location) - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

&#x2022; [`store_provider`](#store-provider) - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

#### More Option Request Headers To Add Secret Value Clear Secret Info

A [`clear_secret_info`](#more-option-request-headers-to-add-secret-value-clear-secret-info) block (within [`more_option.request_headers_to_add.secret_value`](#more-option-request-headers-to-add-secret-value)) supports the following:

&#x2022; [`provider_ref`](#provider-ref) - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

&#x2022; [`url`](#url) - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

#### More Option Response Cookies To Add

A [`response_cookies_to_add`](#more-option-response-cookies-to-add) block (within [`more_option`](#more-option)) supports the following:

&#x2022; [`add_domain`](#add-domain) - Optional String<br>Add Domain. Add domain attribute

&#x2022; [`add_expiry`](#add-expiry) - Optional String<br>Add expiry. Add expiry attribute

&#x2022; [`add_httponly`](#add-httponly) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`add_partitioned`](#add-partitioned) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`add_path`](#add-path) - Optional String<br>Add path. Add path attribute

&#x2022; [`add_secure`](#add-secure) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`ignore_domain`](#ignore-domain) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`ignore_expiry`](#ignore-expiry) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`ignore_httponly`](#ignore-httponly) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`ignore_max_age`](#ignore-max-age) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`ignore_partitioned`](#ignore-partitioned) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`ignore_path`](#ignore-path) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`ignore_samesite`](#ignore-samesite) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`ignore_secure`](#ignore-secure) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`ignore_value`](#ignore-value) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`max_age_value`](#max-age-value) - Optional Number<br>Add Max Age. Add max age attribute

&#x2022; [`name`](#name) - Optional String<br>Name. Name of the cookie in Cookie header

&#x2022; [`overwrite`](#overwrite) - Optional Bool  Defaults to `do`<br>Overwrite. Should the value be overwritten? If true, the value is overwritten to existing values. not overwrite

&#x2022; [`samesite_lax`](#samesite-lax) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`samesite_none`](#samesite-none) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`samesite_strict`](#samesite-strict) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`secret_value`](#secret-value) - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Secret Value](#more-option-response-cookies-to-add-secret-value) below.

&#x2022; [`value`](#value) - Optional String<br>Value. Value of the Cookie header

#### More Option Response Cookies To Add Secret Value

A [`secret_value`](#more-option-response-cookies-to-add-secret-value) block (within [`more_option.response_cookies_to_add`](#more-option-response-cookies-to-add)) supports the following:

&#x2022; [`blindfold_secret_info`](#blindfold-secret-info) - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#more-option-response-cookies-to-add-secret-value-blindfold-secret-info) below.

&#x2022; [`clear_secret_info`](#clear-secret-info) - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#more-option-response-cookies-to-add-secret-value-clear-secret-info) below.

#### More Option Response Cookies To Add Secret Value Blindfold Secret Info

A [`blindfold_secret_info`](#more-option-response-cookies-to-add-secret-value-blindfold-secret-info) block (within [`more_option.response_cookies_to_add.secret_value`](#more-option-response-cookies-to-add-secret-value)) supports the following:

&#x2022; [`decryption_provider`](#decryption-provider) - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

&#x2022; [`location`](#location) - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

&#x2022; [`store_provider`](#store-provider) - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

#### More Option Response Cookies To Add Secret Value Clear Secret Info

A [`clear_secret_info`](#more-option-response-cookies-to-add-secret-value-clear-secret-info) block (within [`more_option.response_cookies_to_add.secret_value`](#more-option-response-cookies-to-add-secret-value)) supports the following:

&#x2022; [`provider_ref`](#provider-ref) - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

&#x2022; [`url`](#url) - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

#### More Option Response Headers To Add

A [`response_headers_to_add`](#more-option-response-headers-to-add) block (within [`more_option`](#more-option)) supports the following:

&#x2022; [`append`](#append) - Optional Bool  Defaults to `do`<br>Append. Should the value be appended? If true, the value is appended to existing values. not append

&#x2022; [`name`](#name) - Optional String<br>Name. Name of the HTTP header

&#x2022; [`secret_value`](#secret-value) - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Secret Value](#more-option-response-headers-to-add-secret-value) below.

&#x2022; [`value`](#value) - Optional String<br>Value. Value of the HTTP header

#### More Option Response Headers To Add Secret Value

A [`secret_value`](#more-option-response-headers-to-add-secret-value) block (within [`more_option.response_headers_to_add`](#more-option-response-headers-to-add)) supports the following:

&#x2022; [`blindfold_secret_info`](#blindfold-secret-info) - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#more-option-response-headers-to-add-secret-value-blindfold-secret-info) below.

&#x2022; [`clear_secret_info`](#clear-secret-info) - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#more-option-response-headers-to-add-secret-value-clear-secret-info) below.

#### More Option Response Headers To Add Secret Value Blindfold Secret Info

A [`blindfold_secret_info`](#more-option-response-headers-to-add-secret-value-blindfold-secret-info) block (within [`more_option.response_headers_to_add.secret_value`](#more-option-response-headers-to-add-secret-value)) supports the following:

&#x2022; [`decryption_provider`](#decryption-provider) - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

&#x2022; [`location`](#location) - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

&#x2022; [`store_provider`](#store-provider) - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

#### More Option Response Headers To Add Secret Value Clear Secret Info

A [`clear_secret_info`](#more-option-response-headers-to-add-secret-value-clear-secret-info) block (within [`more_option.response_headers_to_add.secret_value`](#more-option-response-headers-to-add-secret-value)) supports the following:

&#x2022; [`provider_ref`](#provider-ref) - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

&#x2022; [`url`](#url) - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

#### Origin Server Subset Rule List

An [`origin_server_subset_rule_list`](#origin-server-subset-rule-list) block supports the following:

&#x2022; [`origin_server_subset_rules`](#origin-server-subset-rules) - Optional Block<br>Origin Server Subset Rules. Origin Server Subset Rules allow users to define match condition on Client (IP address, ASN, Country), IP Reputation, Regional Edge names, Request for subset selection of origin servers. Origin Server Subset is a sequential engine where rules are evaluated one after the other. It's important to define the correct order for Origin Server Subset to get the intended result, rules are evaluated from top to bottom in the list. When an Origin server subset rule is matched, then this selection rule takes effect and no more rules are evaluated<br>See [Origin Server Subset Rules](#origin-server-subset-rule-list-origin-server-subset-rules) below.

#### Origin Server Subset Rule List Origin Server Subset Rules

An [`origin_server_subset_rules`](#origin-server-subset-rule-list-origin-server-subset-rules) block (within [`origin_server_subset_rule_list`](#origin-server-subset-rule-list)) supports the following:

&#x2022; [`any_asn`](#any-asn) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`any_ip`](#any-ip) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`asn_list`](#asn-list) - Optional Block<br>ASN Match List. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer<br>See [Asn List](#origin-server-subset-rule-list-origin-server-subset-rules-asn-list) below.

&#x2022; [`asn_matcher`](#asn-matcher) - Optional Block<br>ASN Matcher. Match any AS number contained in the list of bgp_asn_sets<br>See [Asn Matcher](#origin-server-subset-rule-list-origin-server-subset-rules-asn-matcher) below.

&#x2022; [`client_selector`](#client-selector) - Optional Block<br>Label Selector. This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expresssions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects. Label selector is immutable. expressions is a list of strings of label selection expression. Each string has ',' separated values which are 'AND' and all strings are logically 'OR'. BNF for expression string <selector-syntax> ::= <requirement> | <requirement> ',' <selector-syntax> <requirement> ::= [!] KEY [ <set-based-restriction> | <exact-match-restriction> ] <set-based-restriction> ::= '' | <inclusion-exclusion> <value-set> <inclusion-exclusion> ::= <inclusion> | <exclusion> <exclusion> ::= 'notin' <inclusion> ::= 'in' <value-set> ::= '(' <values> ')' <values> ::= VALUE | VALUE ',' <values> <exact-match-restriction> ::= ['='|'=='|'!='] VALUE<br>See [Client Selector](#origin-server-subset-rule-list-origin-server-subset-rules-client-selector) below.

&#x2022; [`country_codes`](#country-codes) - Optional List  Defaults to `COUNTRY_NONE`<br>Possible values are `COUNTRY_NONE`, `COUNTRY_AD`, `COUNTRY_AE`, `COUNTRY_AF`, `COUNTRY_AG`, `COUNTRY_AI`, `COUNTRY_AL`, `COUNTRY_AM`, `COUNTRY_AN`, `COUNTRY_AO`, `COUNTRY_AQ`, `COUNTRY_AR`, `COUNTRY_AS`, `COUNTRY_AT`, `COUNTRY_AU`, `COUNTRY_AW`, `COUNTRY_AX`, `COUNTRY_AZ`, `COUNTRY_BA`, `COUNTRY_BB`, `COUNTRY_BD`, `COUNTRY_BE`, `COUNTRY_BF`, `COUNTRY_BG`, `COUNTRY_BH`, `COUNTRY_BI`, `COUNTRY_BJ`, `COUNTRY_BL`, `COUNTRY_BM`, `COUNTRY_BN`, `COUNTRY_BO`, `COUNTRY_BQ`, `COUNTRY_BR`, `COUNTRY_BS`, `COUNTRY_BT`, `COUNTRY_BV`, `COUNTRY_BW`, `COUNTRY_BY`, `COUNTRY_BZ`, `COUNTRY_CA`, `COUNTRY_CC`, `COUNTRY_CD`, `COUNTRY_CF`, `COUNTRY_CG`, `COUNTRY_CH`, `COUNTRY_CI`, `COUNTRY_CK`, `COUNTRY_CL`, `COUNTRY_CM`, `COUNTRY_CN`, `COUNTRY_CO`, `COUNTRY_CR`, `COUNTRY_CS`, `COUNTRY_CU`, `COUNTRY_CV`, `COUNTRY_CW`, `COUNTRY_CX`, `COUNTRY_CY`, `COUNTRY_CZ`, `COUNTRY_DE`, `COUNTRY_DJ`, `COUNTRY_DK`, `COUNTRY_DM`, `COUNTRY_DO`, `COUNTRY_DZ`, `COUNTRY_EC`, `COUNTRY_EE`, `COUNTRY_EG`, `COUNTRY_EH`, `COUNTRY_ER`, `COUNTRY_ES`, `COUNTRY_ET`, `COUNTRY_FI`, `COUNTRY_FJ`, `COUNTRY_FK`, `COUNTRY_FM`, `COUNTRY_FO`, `COUNTRY_FR`, `COUNTRY_GA`, `COUNTRY_GB`, `COUNTRY_GD`, `COUNTRY_GE`, `COUNTRY_GF`, `COUNTRY_GG`, `COUNTRY_GH`, `COUNTRY_GI`, `COUNTRY_GL`, `COUNTRY_GM`, `COUNTRY_GN`, `COUNTRY_GP`, `COUNTRY_GQ`, `COUNTRY_GR`, `COUNTRY_GS`, `COUNTRY_GT`, `COUNTRY_GU`, `COUNTRY_GW`, `COUNTRY_GY`, `COUNTRY_HK`, `COUNTRY_HM`, `COUNTRY_HN`, `COUNTRY_HR`, `COUNTRY_HT`, `COUNTRY_HU`, `COUNTRY_ID`, `COUNTRY_IE`, `COUNTRY_IL`, `COUNTRY_IM`, `COUNTRY_IN`, `COUNTRY_IO`, `COUNTRY_IQ`, `COUNTRY_IR`, `COUNTRY_IS`, `COUNTRY_IT`, `COUNTRY_JE`, `COUNTRY_JM`, `COUNTRY_JO`, `COUNTRY_JP`, `COUNTRY_KE`, `COUNTRY_KG`, `COUNTRY_KH`, `COUNTRY_KI`, `COUNTRY_KM`, `COUNTRY_KN`, `COUNTRY_KP`, `COUNTRY_KR`, `COUNTRY_KW`, `COUNTRY_KY`, `COUNTRY_KZ`, `COUNTRY_LA`, `COUNTRY_LB`, `COUNTRY_LC`, `COUNTRY_LI`, `COUNTRY_LK`, `COUNTRY_LR`, `COUNTRY_LS`, `COUNTRY_LT`, `COUNTRY_LU`, `COUNTRY_LV`, `COUNTRY_LY`, `COUNTRY_MA`, `COUNTRY_MC`, `COUNTRY_MD`, `COUNTRY_ME`, `COUNTRY_MF`, `COUNTRY_MG`, `COUNTRY_MH`, `COUNTRY_MK`, `COUNTRY_ML`, `COUNTRY_MM`, `COUNTRY_MN`, `COUNTRY_MO`, `COUNTRY_MP`, `COUNTRY_MQ`, `COUNTRY_MR`, `COUNTRY_MS`, `COUNTRY_MT`, `COUNTRY_MU`, `COUNTRY_MV`, `COUNTRY_MW`, `COUNTRY_MX`, `COUNTRY_MY`, `COUNTRY_MZ`, `COUNTRY_NA`, `COUNTRY_NC`, `COUNTRY_NE`, `COUNTRY_NF`, `COUNTRY_NG`, `COUNTRY_NI`, `COUNTRY_NL`, `COUNTRY_NO`, `COUNTRY_NP`, `COUNTRY_NR`, `COUNTRY_NU`, `COUNTRY_NZ`, `COUNTRY_OM`, `COUNTRY_PA`, `COUNTRY_PE`, `COUNTRY_PF`, `COUNTRY_PG`, `COUNTRY_PH`, `COUNTRY_PK`, `COUNTRY_PL`, `COUNTRY_PM`, `COUNTRY_PN`, `COUNTRY_PR`, `COUNTRY_PS`, `COUNTRY_PT`, `COUNTRY_PW`, `COUNTRY_PY`, `COUNTRY_QA`, `COUNTRY_RE`, `COUNTRY_RO`, `COUNTRY_RS`, `COUNTRY_RU`, `COUNTRY_RW`, `COUNTRY_SA`, `COUNTRY_SB`, `COUNTRY_SC`, `COUNTRY_SD`, `COUNTRY_SE`, `COUNTRY_SG`, `COUNTRY_SH`, `COUNTRY_SI`, `COUNTRY_SJ`, `COUNTRY_SK`, `COUNTRY_SL`, `COUNTRY_SM`, `COUNTRY_SN`, `COUNTRY_SO`, `COUNTRY_SR`, `COUNTRY_SS`, `COUNTRY_ST`, `COUNTRY_SV`, `COUNTRY_SX`, `COUNTRY_SY`, `COUNTRY_SZ`, `COUNTRY_TC`, `COUNTRY_TD`, `COUNTRY_TF`, `COUNTRY_TG`, `COUNTRY_TH`, `COUNTRY_TJ`, `COUNTRY_TK`, `COUNTRY_TL`, `COUNTRY_TM`, `COUNTRY_TN`, `COUNTRY_TO`, `COUNTRY_TR`, `COUNTRY_TT`, `COUNTRY_TV`, `COUNTRY_TW`, `COUNTRY_TZ`, `COUNTRY_UA`, `COUNTRY_UG`, `COUNTRY_UM`, `COUNTRY_US`, `COUNTRY_UY`, `COUNTRY_UZ`, `COUNTRY_VA`, `COUNTRY_VC`, `COUNTRY_VE`, `COUNTRY_VG`, `COUNTRY_VI`, `COUNTRY_VN`, `COUNTRY_VU`, `COUNTRY_WF`, `COUNTRY_WS`, `COUNTRY_XK`, `COUNTRY_XT`, `COUNTRY_YE`, `COUNTRY_YT`, `COUNTRY_ZA`, `COUNTRY_ZM`, `COUNTRY_ZW`<br>Country Codes List. List of Country Codes

&#x2022; [`ip_matcher`](#ip-matcher) - Optional Block<br>IP Prefix Matcher. Match any IP prefix contained in the list of ip_prefix_sets. The result of the match is inverted if invert_matcher is true<br>See [IP Matcher](#origin-server-subset-rule-list-origin-server-subset-rules-ip-matcher) below.

&#x2022; [`ip_prefix_list`](#ip-prefix-list) - Optional Block<br>IP Prefix Match List. List of IP Prefix strings to match against<br>See [IP Prefix List](#origin-server-subset-rule-list-origin-server-subset-rules-ip-prefix-list) below.

&#x2022; [`metadata`](#metadata) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#origin-server-subset-rule-list-origin-server-subset-rules-metadata) below.

&#x2022; [`none`](#none) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`origin_server_subsets_action`](#origin-server-subsets-action) - Optional Block<br>Action. Add labels to select one or more origin servers. Note: The pre-requisite settings to be configured in the origin pool are: 1. Add labels to origin servers 2. Enable subset load balancing in the Origin Server Subsets section and configure keys in origin server subsets classes

&#x2022; [`re_name_list`](#re-name-list) - Optional List<br>RE Names. List of RE names for match

#### Origin Server Subset Rule List Origin Server Subset Rules Asn List

An [`asn_list`](#origin-server-subset-rule-list-origin-server-subset-rules-asn-list) block (within [`origin_server_subset_rule_list.origin_server_subset_rules`](#origin-server-subset-rule-list-origin-server-subset-rules)) supports the following:

&#x2022; [`as_numbers`](#as-numbers) - Optional List<br>AS Numbers. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer

#### Origin Server Subset Rule List Origin Server Subset Rules Asn Matcher

An [`asn_matcher`](#origin-server-subset-rule-list-origin-server-subset-rules-asn-matcher) block (within [`origin_server_subset_rule_list.origin_server_subset_rules`](#origin-server-subset-rule-list-origin-server-subset-rules)) supports the following:

&#x2022; [`asn_sets`](#asn-sets) - Optional Block<br>BGP ASN Sets. A list of references to bgp_asn_set objects<br>See [Asn Sets](#origin-server-subset-rule-list-origin-server-subset-rules-asn-matcher-asn-sets) below.

#### Origin Server Subset Rule List Origin Server Subset Rules Asn Matcher Asn Sets

An [`asn_sets`](#origin-server-subset-rule-list-origin-server-subset-rules-asn-matcher-asn-sets) block (within [`origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher`](#origin-server-subset-rule-list-origin-server-subset-rules-asn-matcher)) supports the following:

&#x2022; [`kind`](#kind) - Optional String<br>Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route')

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

&#x2022; [`uid`](#uid) - Optional String<br>UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid

#### Origin Server Subset Rule List Origin Server Subset Rules Client Selector

A [`client_selector`](#origin-server-subset-rule-list-origin-server-subset-rules-client-selector) block (within [`origin_server_subset_rule_list.origin_server_subset_rules`](#origin-server-subset-rule-list-origin-server-subset-rules)) supports the following:

&#x2022; [`expressions`](#expressions) - Optional List<br>Selector Expression. expressions contains the kubernetes style label expression for selections

#### Origin Server Subset Rule List Origin Server Subset Rules IP Matcher

An [`ip_matcher`](#origin-server-subset-rule-list-origin-server-subset-rules-ip-matcher) block (within [`origin_server_subset_rule_list.origin_server_subset_rules`](#origin-server-subset-rule-list-origin-server-subset-rules)) supports the following:

&#x2022; [`invert_matcher`](#invert-matcher) - Optional Bool<br>Invert IP Matcher. Invert the match result

&#x2022; [`prefix_sets`](#prefix-sets) - Optional Block<br>IP Prefix Sets. A list of references to ip_prefix_set objects<br>See [Prefix Sets](#origin-server-subset-rule-list-origin-server-subset-rules-ip-matcher-prefix-sets) below.

#### Origin Server Subset Rule List Origin Server Subset Rules IP Matcher Prefix Sets

A [`prefix_sets`](#origin-server-subset-rule-list-origin-server-subset-rules-ip-matcher-prefix-sets) block (within [`origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher`](#origin-server-subset-rule-list-origin-server-subset-rules-ip-matcher)) supports the following:

&#x2022; [`kind`](#kind) - Optional String<br>Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route')

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

&#x2022; [`uid`](#uid) - Optional String<br>UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid

#### Origin Server Subset Rule List Origin Server Subset Rules IP Prefix List

An [`ip_prefix_list`](#origin-server-subset-rule-list-origin-server-subset-rules-ip-prefix-list) block (within [`origin_server_subset_rule_list.origin_server_subset_rules`](#origin-server-subset-rule-list-origin-server-subset-rules)) supports the following:

&#x2022; [`invert_match`](#invert-match) - Optional Bool<br>Invert Match Result. Invert the match result

&#x2022; [`ip_prefixes`](#ip-prefixes) - Optional List<br>IPv4 Prefix List. List of IPv4 prefix strings

#### Origin Server Subset Rule List Origin Server Subset Rules Metadata

A [`metadata`](#origin-server-subset-rule-list-origin-server-subset-rules-metadata) block (within [`origin_server_subset_rule_list.origin_server_subset_rules`](#origin-server-subset-rule-list-origin-server-subset-rules)) supports the following:

&#x2022; [`description`](#description) - Optional String<br>Description. Human readable description

&#x2022; [`name`](#name) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### Policy Based Challenge

A [`policy_based_challenge`](#policy-based-challenge) block supports the following:

&#x2022; [`always_enable_captcha_challenge`](#always-enable-captcha-challenge) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`always_enable_js_challenge`](#always-enable-js-challenge) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`captcha_challenge_parameters`](#captcha-challenge-parameters) - Optional Block<br>Captcha Challenge Parameters. Enables loadbalancer to perform captcha challenge Captcha challenge will be based on Google Recaptcha. With this feature enabled, only clients that pass the captcha challenge will be allowed to complete the HTTP request. When loadbalancer is configured to do Captcha Challenge, it will redirect the browser to an HTML page on every new HTTP request. This HTML page will have captcha challenge embedded in it. Client will be allowed to make the request only if the captcha challenge is successful. Loadbalancer will tag response header with a cookie to avoid Captcha challenge for subsequent requests. CAPTCHA is mainly used as a security check to ensure only human users can pass through. Generally, computers or bots are not capable of solving a captcha. You can enable either Javascript challenge or Captcha challenge on a virtual host<br>See [Captcha Challenge Parameters](#policy-based-challenge-captcha-challenge-parameters) below.

&#x2022; [`default_captcha_challenge_parameters`](#default-captcha-challenge-parameters) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`default_js_challenge_parameters`](#default-js-challenge-parameters) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`default_mitigation_settings`](#default-mitigation-settings) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`default_temporary_blocking_parameters`](#default-temporary-blocking-parameters) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`js_challenge_parameters`](#js-challenge-parameters) - Optional Block<br>Javascript Challenge Parameters. Enables loadbalancer to perform client browser compatibility test by redirecting to a page with Javascript. With this feature enabled, only clients that are capable of executing Javascript(mostly browsers) will be allowed to complete the HTTP request. When loadbalancer is configured to do Javascript Challenge, it will redirect the browser to an HTML page on every new HTTP request. This HTML page will have Javascript embedded in it. Loadbalancer chooses a set of random numbers for every new client and sends these numbers along with an encrypted answer with the request such that it embed these numbers as input in the Javascript. Javascript will run on the requestor browser and perform a complex Math operation. Script will submit the answer to loadbalancer. Loadbalancer will validate the answer by comparing the calculated answer with the decrypted answer (which was encrypted when it was sent back as reply) and allow the request to the upstream server only if the answer is correct. Loadbalancer will tag response header with a cookie to avoid Javascript challenge for subsequent requests. Javascript challenge serves following purposes * Validate that the request is coming via a browser that is capable for running Javascript * Force the browser to run a complex operation, f(X), that requires it to spend a large number of CPU cycles. This is to slow down a potential DOS attacker by making it difficult to launch a large request flood without having to spend even larger CPU cost at their end. You can enable either Javascript challenge or Captcha challenge on a virtual host<br>See [Js Challenge Parameters](#policy-based-challenge-js-challenge-parameters) below.

&#x2022; [`malicious_user_mitigation`](#malicious-user-mitigation) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Malicious User Mitigation](#policy-based-challenge-malicious-user-mitigation) below.

&#x2022; [`no_challenge`](#no-challenge) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`rule_list`](#rule-list) - Optional Block<br>Challenge Rule List. List of challenge rules to be used in policy based challenge<br>See [Rule List](#policy-based-challenge-rule-list) below.

&#x2022; [`temporary_user_blocking`](#temporary-user-blocking) - Optional Block<br>Temporary User Blocking. Specifies configuration for temporary user blocking resulting from user behavior analysis. When Malicious User Mitigation is enabled from service policy rules, users' accessing the application will be analyzed for malicious activity and the configured mitigation actions will be taken on identified malicious users. These mitigation actions include setting up temporary blocking on that user. This configuration specifies settings on how that blocking should be done by the loadbalancer<br>See [Temporary User Blocking](#policy-based-challenge-temporary-user-blocking) below.

#### Policy Based Challenge Captcha Challenge Parameters

A [`captcha_challenge_parameters`](#policy-based-challenge-captcha-challenge-parameters) block (within [`policy_based_challenge`](#policy-based-challenge)) supports the following:

&#x2022; [`cookie_expiry`](#cookie-expiry) - Optional Number<br>Cookie Expiration Period. Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new challenge

&#x2022; [`custom_page`](#custom-page) - Optional String<br>Custom message for Captcha Challenge. Custom message is of type uri_ref. Currently supported URL schemes is string:///. For string:/// scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. 'Please Wait.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '<p> Please Wait </p>'. Base64 encoded string for this HTML is 'PHA+IFBsZWFzZSBXYWl0IDwvcD4='

#### Policy Based Challenge Js Challenge Parameters

A [`js_challenge_parameters`](#policy-based-challenge-js-challenge-parameters) block (within [`policy_based_challenge`](#policy-based-challenge)) supports the following:

&#x2022; [`cookie_expiry`](#cookie-expiry) - Optional Number<br>Cookie Expiration Period. Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new challenge

&#x2022; [`custom_page`](#custom-page) - Optional String<br>Custom Message for Javascript Challenge. Custom message is of type uri_ref. Currently supported URL schemes is string:///. For string:/// scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. 'Please Wait.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '<p> Please Wait </p>'. Base64 encoded string for this HTML is 'PHA+IFBsZWFzZSBXYWl0IDwvcD4='

&#x2022; [`js_script_delay`](#js-script-delay) - Optional Number<br>Javascript Delay. Delay introduced by Javascript, in milliseconds

#### Policy Based Challenge Malicious User Mitigation

A [`malicious_user_mitigation`](#policy-based-challenge-malicious-user-mitigation) block (within [`policy_based_challenge`](#policy-based-challenge)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Policy Based Challenge Rule List

A [`rule_list`](#policy-based-challenge-rule-list) block (within [`policy_based_challenge`](#policy-based-challenge)) supports the following:

&#x2022; [`rules`](#rules) - Optional Block<br>Rules. Rules that specify the match conditions and challenge type to be launched. When a challenge type is selected to be always enabled, these rules can be used to disable challenge or launch a different challenge for requests that match the specified conditions<br>See [Rules](#policy-based-challenge-rule-list-rules) below.

#### Policy Based Challenge Rule List Rules

A [`rules`](#policy-based-challenge-rule-list-rules) block (within [`policy_based_challenge.rule_list`](#policy-based-challenge-rule-list)) supports the following:

&#x2022; [`metadata`](#metadata) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#policy-based-challenge-rule-list-rules-metadata) below.

&#x2022; [`spec`](#spec) - Optional Block<br>Challenge Rule Specification. A Challenge Rule consists of an unordered list of predicates and an action. The predicates are evaluated against a set of input fields that are extracted from or derived from an L7 request API. A request API is considered to match the rule if all predicates in the rule evaluate to true for that request. Any predicates that are not specified in a rule are implicitly considered to be true. If a request API matches a challenge rule, the configured challenge is enforced<br>See [Spec](#policy-based-challenge-rule-list-rules-spec) below.

#### Policy Based Challenge Rule List Rules Metadata

A [`metadata`](#policy-based-challenge-rule-list-rules-metadata) block (within [`policy_based_challenge.rule_list.rules`](#policy-based-challenge-rule-list-rules)) supports the following:

&#x2022; [`description`](#description) - Optional String<br>Description. Human readable description

&#x2022; [`name`](#name) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### Policy Based Challenge Rule List Rules Spec

A [`spec`](#policy-based-challenge-rule-list-rules-spec) block (within [`policy_based_challenge.rule_list.rules`](#policy-based-challenge-rule-list-rules)) supports the following:

&#x2022; [`any_asn`](#any-asn) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`any_client`](#any-client) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`any_ip`](#any-ip) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`arg_matchers`](#arg-matchers) - Optional Block<br>A list of predicates for all POST args that need to be matched. The criteria for matching each arg are described in individual instances of ArgMatcherType. The actual arg values are extracted from the request API as a list of strings for each arg selector name. Note that all specified arg matcher predicates must evaluate to true<br>See [Arg Matchers](#policy-based-challenge-rule-list-rules-spec-arg-matchers) below.

&#x2022; [`asn_list`](#asn-list) - Optional Block<br>ASN Match List. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer<br>See [Asn List](#policy-based-challenge-rule-list-rules-spec-asn-list) below.

&#x2022; [`asn_matcher`](#asn-matcher) - Optional Block<br>ASN Matcher. Match any AS number contained in the list of bgp_asn_sets<br>See [Asn Matcher](#policy-based-challenge-rule-list-rules-spec-asn-matcher) below.

&#x2022; [`body_matcher`](#body-matcher) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Body Matcher](#policy-based-challenge-rule-list-rules-spec-body-matcher) below.

&#x2022; [`client_selector`](#client-selector) - Optional Block<br>Label Selector. This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expresssions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects. Label selector is immutable. expressions is a list of strings of label selection expression. Each string has ',' separated values which are 'AND' and all strings are logically 'OR'. BNF for expression string <selector-syntax> ::= <requirement> | <requirement> ',' <selector-syntax> <requirement> ::= [!] KEY [ <set-based-restriction> | <exact-match-restriction> ] <set-based-restriction> ::= '' | <inclusion-exclusion> <value-set> <inclusion-exclusion> ::= <inclusion> | <exclusion> <exclusion> ::= 'notin' <inclusion> ::= 'in' <value-set> ::= '(' <values> ')' <values> ::= VALUE | VALUE ',' <values> <exact-match-restriction> ::= ['='|'=='|'!='] VALUE<br>See [Client Selector](#policy-based-challenge-rule-list-rules-spec-client-selector) below.

&#x2022; [`cookie_matchers`](#cookie-matchers) - Optional Block<br>A list of predicates for all cookies that need to be matched. The criteria for matching each cookie is described in individual instances of CookieMatcherType. The actual cookie values are extracted from the request API as a list of strings for each cookie name. Note that all specified cookie matcher predicates must evaluate to true<br>See [Cookie Matchers](#policy-based-challenge-rule-list-rules-spec-cookie-matchers) below.

&#x2022; [`disable_challenge`](#disable-challenge) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`domain_matcher`](#domain-matcher) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Domain Matcher](#policy-based-challenge-rule-list-rules-spec-domain-matcher) below.

&#x2022; [`enable_captcha_challenge`](#enable-captcha-challenge) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`enable_javascript_challenge`](#enable-javascript-challenge) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`expiration_timestamp`](#expiration-timestamp) - Optional String<br>The expiration_timestamp is the RFC 3339 format timestamp at which the containing rule is considered to be logically expired. The rule continues to exist in the configuration but is not applied anymore

&#x2022; [`headers`](#headers) - Optional Block<br>A list of predicates for various HTTP headers that need to match. The criteria for matching each HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values are extracted from the request API as a list of strings for each HTTP header type. Note that all specified header predicates must evaluate to true<br>See [Headers](#policy-based-challenge-rule-list-rules-spec-headers) below.

&#x2022; [`http_method`](#http-method) - Optional Block<br>HTTP Method Matcher. A HTTP method matcher specifies a list of methods to match an input HTTP method. The match is considered successful if the input method is a member of the list. The result of the match based on the method list is inverted if invert_matcher is true<br>See [HTTP Method](#policy-based-challenge-rule-list-rules-spec-http-method) below.

&#x2022; [`ip_matcher`](#ip-matcher) - Optional Block<br>IP Prefix Matcher. Match any IP prefix contained in the list of ip_prefix_sets. The result of the match is inverted if invert_matcher is true<br>See [IP Matcher](#policy-based-challenge-rule-list-rules-spec-ip-matcher) below.

&#x2022; [`ip_prefix_list`](#ip-prefix-list) - Optional Block<br>IP Prefix Match List. List of IP Prefix strings to match against<br>See [IP Prefix List](#policy-based-challenge-rule-list-rules-spec-ip-prefix-list) below.

&#x2022; [`path`](#path) - Optional Block<br>Path Matcher. A path matcher specifies multiple criteria for matching an HTTP path string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of path prefixes, a list of exact path values and a list of regular expressions<br>See [Path](#policy-based-challenge-rule-list-rules-spec-path) below.

&#x2022; [`query_params`](#query-params) - Optional Block<br>A list of predicates for all query parameters that need to be matched. The criteria for matching each query parameter are described in individual instances of QueryParameterMatcherType. The actual query parameter values are extracted from the request API as a list of strings for each query parameter name. Note that all specified query parameter predicates must evaluate to true<br>See [Query Params](#policy-based-challenge-rule-list-rules-spec-query-params) below.

&#x2022; [`tls_fingerprint_matcher`](#tls-fingerprint-matcher) - Optional Block<br>TLS Fingerprint Matcher. A TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of supported positve match criteria includes a list of known classes of TLS fingerprints and a list of exact values. The match is considered successful if either of these positive criteria are satisfied and the input fingerprint is not one of the excluded values<br>See [TLS Fingerprint Matcher](#policy-based-challenge-rule-list-rules-spec-tls-fingerprint-matcher) below.

#### Policy Based Challenge Rule List Rules Spec Arg Matchers

An [`arg_matchers`](#policy-based-challenge-rule-list-rules-spec-arg-matchers) block (within [`policy_based_challenge.rule_list.rules.spec`](#policy-based-challenge-rule-list-rules-spec)) supports the following:

&#x2022; [`check_not_present`](#check-not-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`check_present`](#check-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`invert_matcher`](#invert-matcher) - Optional Bool<br>Invert Matcher. Invert Match of the expression defined

&#x2022; [`item`](#item) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#policy-based-challenge-rule-list-rules-spec-arg-matchers-item) below.

&#x2022; [`name`](#name) - Optional String<br>Argument Name. x-example: 'phones[_]' x-example: 'cars.make.toyota.models[1]' x-example: 'cars.make.honda.models[_]' x-example: 'cars.make[_].models[_]' A case-sensitive JSON path in the HTTP request body

#### Policy Based Challenge Rule List Rules Spec Arg Matchers Item

An [`item`](#policy-based-challenge-rule-list-rules-spec-arg-matchers-item) block (within [`policy_based_challenge.rule_list.rules.spec.arg_matchers`](#policy-based-challenge-rule-list-rules-spec-arg-matchers)) supports the following:

&#x2022; [`exact_values`](#exact-values) - Optional List<br>Exact Values. A list of exact values to match the input against

&#x2022; [`regex_values`](#regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

&#x2022; [`transformers`](#transformers) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### Policy Based Challenge Rule List Rules Spec Asn List

An [`asn_list`](#policy-based-challenge-rule-list-rules-spec-asn-list) block (within [`policy_based_challenge.rule_list.rules.spec`](#policy-based-challenge-rule-list-rules-spec)) supports the following:

&#x2022; [`as_numbers`](#as-numbers) - Optional List<br>AS Numbers. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer

#### Policy Based Challenge Rule List Rules Spec Asn Matcher

An [`asn_matcher`](#policy-based-challenge-rule-list-rules-spec-asn-matcher) block (within [`policy_based_challenge.rule_list.rules.spec`](#policy-based-challenge-rule-list-rules-spec)) supports the following:

&#x2022; [`asn_sets`](#asn-sets) - Optional Block<br>BGP ASN Sets. A list of references to bgp_asn_set objects<br>See [Asn Sets](#policy-based-challenge-rule-list-rules-spec-asn-matcher-asn-sets) below.

#### Policy Based Challenge Rule List Rules Spec Asn Matcher Asn Sets

An [`asn_sets`](#policy-based-challenge-rule-list-rules-spec-asn-matcher-asn-sets) block (within [`policy_based_challenge.rule_list.rules.spec.asn_matcher`](#policy-based-challenge-rule-list-rules-spec-asn-matcher)) supports the following:

&#x2022; [`kind`](#kind) - Optional String<br>Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route')

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

&#x2022; [`uid`](#uid) - Optional String<br>UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid

#### Policy Based Challenge Rule List Rules Spec Body Matcher

A [`body_matcher`](#policy-based-challenge-rule-list-rules-spec-body-matcher) block (within [`policy_based_challenge.rule_list.rules.spec`](#policy-based-challenge-rule-list-rules-spec)) supports the following:

&#x2022; [`exact_values`](#exact-values) - Optional List<br>Exact Values. A list of exact values to match the input against

&#x2022; [`regex_values`](#regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

&#x2022; [`transformers`](#transformers) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### Policy Based Challenge Rule List Rules Spec Client Selector

A [`client_selector`](#policy-based-challenge-rule-list-rules-spec-client-selector) block (within [`policy_based_challenge.rule_list.rules.spec`](#policy-based-challenge-rule-list-rules-spec)) supports the following:

&#x2022; [`expressions`](#expressions) - Optional List<br>Selector Expression. expressions contains the kubernetes style label expression for selections

#### Policy Based Challenge Rule List Rules Spec Cookie Matchers

A [`cookie_matchers`](#policy-based-challenge-rule-list-rules-spec-cookie-matchers) block (within [`policy_based_challenge.rule_list.rules.spec`](#policy-based-challenge-rule-list-rules-spec)) supports the following:

&#x2022; [`check_not_present`](#check-not-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`check_present`](#check-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`invert_matcher`](#invert-matcher) - Optional Bool<br>Invert Matcher. Invert Match of the expression defined

&#x2022; [`item`](#item) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#policy-based-challenge-rule-list-rules-spec-cookie-matchers-item) below.

&#x2022; [`name`](#name) - Optional String<br>Cookie Name. A case-sensitive cookie name

#### Policy Based Challenge Rule List Rules Spec Cookie Matchers Item

An [`item`](#policy-based-challenge-rule-list-rules-spec-cookie-matchers-item) block (within [`policy_based_challenge.rule_list.rules.spec.cookie_matchers`](#policy-based-challenge-rule-list-rules-spec-cookie-matchers)) supports the following:

&#x2022; [`exact_values`](#exact-values) - Optional List<br>Exact Values. A list of exact values to match the input against

&#x2022; [`regex_values`](#regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

&#x2022; [`transformers`](#transformers) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### Policy Based Challenge Rule List Rules Spec Domain Matcher

A [`domain_matcher`](#policy-based-challenge-rule-list-rules-spec-domain-matcher) block (within [`policy_based_challenge.rule_list.rules.spec`](#policy-based-challenge-rule-list-rules-spec)) supports the following:

&#x2022; [`exact_values`](#exact-values) - Optional List<br>Exact Values. A list of exact values to match the input against

&#x2022; [`regex_values`](#regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

#### Policy Based Challenge Rule List Rules Spec Headers

A [`headers`](#policy-based-challenge-rule-list-rules-spec-headers) block (within [`policy_based_challenge.rule_list.rules.spec`](#policy-based-challenge-rule-list-rules-spec)) supports the following:

&#x2022; [`check_not_present`](#check-not-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`check_present`](#check-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`invert_matcher`](#invert-matcher) - Optional Bool<br>Invert Header Matcher. Invert the match result

&#x2022; [`item`](#item) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#policy-based-challenge-rule-list-rules-spec-headers-item) below.

&#x2022; [`name`](#name) - Optional String<br>Header Name. A case-insensitive HTTP header name

#### Policy Based Challenge Rule List Rules Spec Headers Item

An [`item`](#policy-based-challenge-rule-list-rules-spec-headers-item) block (within [`policy_based_challenge.rule_list.rules.spec.headers`](#policy-based-challenge-rule-list-rules-spec-headers)) supports the following:

&#x2022; [`exact_values`](#exact-values) - Optional List<br>Exact Values. A list of exact values to match the input against

&#x2022; [`regex_values`](#regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

&#x2022; [`transformers`](#transformers) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### Policy Based Challenge Rule List Rules Spec HTTP Method

A [`http_method`](#policy-based-challenge-rule-list-rules-spec-http-method) block (within [`policy_based_challenge.rule_list.rules.spec`](#policy-based-challenge-rule-list-rules-spec)) supports the following:

&#x2022; [`invert_matcher`](#invert-matcher) - Optional Bool<br>Invert Method Matcher. Invert the match result

&#x2022; [`methods`](#methods) - Optional List  Defaults to `ANY`<br>Possible values are `ANY`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`, `COPY`<br>Method List. List of methods values to match against

#### Policy Based Challenge Rule List Rules Spec IP Matcher

An [`ip_matcher`](#policy-based-challenge-rule-list-rules-spec-ip-matcher) block (within [`policy_based_challenge.rule_list.rules.spec`](#policy-based-challenge-rule-list-rules-spec)) supports the following:

&#x2022; [`invert_matcher`](#invert-matcher) - Optional Bool<br>Invert IP Matcher. Invert the match result

&#x2022; [`prefix_sets`](#prefix-sets) - Optional Block<br>IP Prefix Sets. A list of references to ip_prefix_set objects<br>See [Prefix Sets](#policy-based-challenge-rule-list-rules-spec-ip-matcher-prefix-sets) below.

#### Policy Based Challenge Rule List Rules Spec IP Matcher Prefix Sets

A [`prefix_sets`](#policy-based-challenge-rule-list-rules-spec-ip-matcher-prefix-sets) block (within [`policy_based_challenge.rule_list.rules.spec.ip_matcher`](#policy-based-challenge-rule-list-rules-spec-ip-matcher)) supports the following:

&#x2022; [`kind`](#kind) - Optional String<br>Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route')

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

&#x2022; [`uid`](#uid) - Optional String<br>UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid

#### Policy Based Challenge Rule List Rules Spec IP Prefix List

An [`ip_prefix_list`](#policy-based-challenge-rule-list-rules-spec-ip-prefix-list) block (within [`policy_based_challenge.rule_list.rules.spec`](#policy-based-challenge-rule-list-rules-spec)) supports the following:

&#x2022; [`invert_match`](#invert-match) - Optional Bool<br>Invert Match Result. Invert the match result

&#x2022; [`ip_prefixes`](#ip-prefixes) - Optional List<br>IPv4 Prefix List. List of IPv4 prefix strings

#### Policy Based Challenge Rule List Rules Spec Path

A [`path`](#policy-based-challenge-rule-list-rules-spec-path) block (within [`policy_based_challenge.rule_list.rules.spec`](#policy-based-challenge-rule-list-rules-spec)) supports the following:

&#x2022; [`exact_values`](#exact-values) - Optional List<br>Exact Values. A list of exact path values to match the input HTTP path against

&#x2022; [`invert_matcher`](#invert-matcher) - Optional Bool<br>Invert Path Matcher. Invert the match result

&#x2022; [`prefix_values`](#prefix-values) - Optional List<br>Prefix Values. A list of path prefix values to match the input HTTP path against

&#x2022; [`regex_values`](#regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input HTTP path against

&#x2022; [`suffix_values`](#suffix-values) - Optional List<br>Suffix Values. A list of path suffix values to match the input HTTP path against

&#x2022; [`transformers`](#transformers) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### Policy Based Challenge Rule List Rules Spec Query Params

A [`query_params`](#policy-based-challenge-rule-list-rules-spec-query-params) block (within [`policy_based_challenge.rule_list.rules.spec`](#policy-based-challenge-rule-list-rules-spec)) supports the following:

&#x2022; [`check_not_present`](#check-not-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`check_present`](#check-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`invert_matcher`](#invert-matcher) - Optional Bool<br>Invert Query Parameter Matcher. Invert the match result

&#x2022; [`item`](#item) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#policy-based-challenge-rule-list-rules-spec-query-params-item) below.

&#x2022; [`key`](#key) - Optional String<br>Query Parameter Name. A case-sensitive HTTP query parameter name

#### Policy Based Challenge Rule List Rules Spec Query Params Item

An [`item`](#policy-based-challenge-rule-list-rules-spec-query-params-item) block (within [`policy_based_challenge.rule_list.rules.spec.query_params`](#policy-based-challenge-rule-list-rules-spec-query-params)) supports the following:

&#x2022; [`exact_values`](#exact-values) - Optional List<br>Exact Values. A list of exact values to match the input against

&#x2022; [`regex_values`](#regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

&#x2022; [`transformers`](#transformers) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### Policy Based Challenge Rule List Rules Spec TLS Fingerprint Matcher

A [`tls_fingerprint_matcher`](#policy-based-challenge-rule-list-rules-spec-tls-fingerprint-matcher) block (within [`policy_based_challenge.rule_list.rules.spec`](#policy-based-challenge-rule-list-rules-spec)) supports the following:

&#x2022; [`classes`](#classes) - Optional List  Defaults to `TLS_FINGERPRINT_NONE`<br>Possible values are `TLS_FINGERPRINT_NONE`, `ANY_MALICIOUS_FINGERPRINT`, `ADWARE`, `ADWIND`, `DRIDEX`, `GOOTKIT`, `GOZI`, `JBIFROST`, `QUAKBOT`, `RANSOMWARE`, `TROLDESH`, `TOFSEE`, `TORRENTLOCKER`, `TRICKBOT`<br>TLS fingerprint classes. A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against

&#x2022; [`exact_values`](#exact-values) - Optional List<br>Exact Values. A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against

&#x2022; [`excluded_values`](#excluded-values) - Optional List<br>Excluded Values. A list of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can be used to skip known false positives when using one or more known TLS fingerprint classes in the enclosing matcher

#### Policy Based Challenge Temporary User Blocking

A [`temporary_user_blocking`](#policy-based-challenge-temporary-user-blocking) block (within [`policy_based_challenge`](#policy-based-challenge)) supports the following:

&#x2022; [`custom_page`](#custom-page) - Optional String<br>Custom Message for Temporary Blocking. Custom message is of type `uri_ref`. Currently supported URL schemes is `string:///`. For `string:///` scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. 'Blocked.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '<p> Blocked </p>'. Base64 encoded string for this HTML is 'PHA+IFBsZWFzZSBXYWl0IDwvcD4='

#### Protected Cookies

A [`protected_cookies`](#protected-cookies) block supports the following:

&#x2022; [`add_httponly`](#add-httponly) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`add_secure`](#add-secure) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`disable_tampering_protection`](#disable-tampering-protection) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`enable_tampering_protection`](#enable-tampering-protection) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`ignore_httponly`](#ignore-httponly) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`ignore_max_age`](#ignore-max-age) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`ignore_samesite`](#ignore-samesite) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`ignore_secure`](#ignore-secure) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`max_age_value`](#max-age-value) - Optional Number<br>Add Max Age. Add max age attribute

&#x2022; [`name`](#name) - Optional String<br>Cookie Name. Name of the Cookie

&#x2022; [`samesite_lax`](#samesite-lax) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`samesite_none`](#samesite-none) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`samesite_strict`](#samesite-strict) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### Rate Limit

A [`rate_limit`](#rate-limit) block supports the following:

&#x2022; [`custom_ip_allowed_list`](#custom-ip-allowed-list) - Optional Block<br>Custom IP Allowed List. IP Allowed list using existing ip_prefix_set objects<br>See [Custom IP Allowed List](#rate-limit-custom-ip-allowed-list) below.

&#x2022; [`ip_allowed_list`](#ip-allowed-list) - Optional Block<br>IPv4 Prefix List. x-example: '192.168.20.0/24' List of IPv4 prefixes that represent an endpoint<br>See [IP Allowed List](#rate-limit-ip-allowed-list) below.

&#x2022; [`no_ip_allowed_list`](#no-ip-allowed-list) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`no_policies`](#no-policies) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`policies`](#policies) - Optional Block<br>Rate Limiter Policy List. List of rate limiter policies to be applied<br>See [Policies](#rate-limit-policies) below.

&#x2022; [`rate_limiter`](#rate-limiter) - Optional Block<br>Rate Limit Value. A tuple consisting of a rate limit period unit and the total number of allowed requests for that period<br>See [Rate Limiter](#rate-limit-rate-limiter) below.

#### Rate Limit Custom IP Allowed List

A [`custom_ip_allowed_list`](#rate-limit-custom-ip-allowed-list) block (within [`rate_limit`](#rate-limit)) supports the following:

&#x2022; [`rate_limiter_allowed_prefixes`](#rate-limiter-allowed-prefixes) - Optional Block<br>List of IP Prefix Sets. References to ip_prefix_set objects. Requests from source IP addresses that are covered by one of the allowed IP Prefixes are not subjected to rate limiting<br>See [Rate Limiter Allowed Prefixes](#rate-limit-custom-ip-allowed-list-rate-limiter-allowed-prefixes) below.

#### Rate Limit Custom IP Allowed List Rate Limiter Allowed Prefixes

A [`rate_limiter_allowed_prefixes`](#rate-limit-custom-ip-allowed-list-rate-limiter-allowed-prefixes) block (within [`rate_limit.custom_ip_allowed_list`](#rate-limit-custom-ip-allowed-list)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Rate Limit IP Allowed List

An [`ip_allowed_list`](#rate-limit-ip-allowed-list) block (within [`rate_limit`](#rate-limit)) supports the following:

&#x2022; [`prefixes`](#prefixes) - Optional List<br>IPv4 Prefix List. List of IPv4 prefixes that represent an endpoint

#### Rate Limit Policies

A [`policies`](#rate-limit-policies) block (within [`rate_limit`](#rate-limit)) supports the following:

&#x2022; [`policies`](#policies) - Optional Block<br>Rate Limiter Policies. Ordered list of rate limiter policies<br>See [Policies](#rate-limit-policies-policies) below.

#### Rate Limit Policies Policies

A [`policies`](#rate-limit-policies-policies) block (within [`rate_limit.policies`](#rate-limit-policies)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Rate Limit Rate Limiter

A [`rate_limiter`](#rate-limit-rate-limiter) block (within [`rate_limit`](#rate-limit)) supports the following:

&#x2022; [`action_block`](#action-block) - Optional Block<br>Rate Limit Block Action. Action where a user is blocked from making further requests after exceeding rate limit threshold<br>See [Action Block](#rate-limit-rate-limiter-action-block) below.

&#x2022; [`burst_multiplier`](#burst-multiplier) - Optional Number<br>Burst Multiplier. The maximum burst of requests to accommodate, expressed as a multiple of the rate

&#x2022; [`disabled`](#disabled) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`leaky_bucket`](#leaky-bucket) - Optional Block<br>Leaky Bucket Rate Limiter. Leaky-Bucket is the default rate limiter algorithm for F5

&#x2022; [`period_multiplier`](#period-multiplier) - Optional Number<br>Periods. This setting, combined with Per Period units, provides a duration

&#x2022; [`token_bucket`](#token-bucket) - Optional Block<br>Token Bucket Rate Limiter. Token-Bucket is a rate limiter algorithm that is stricter with enforcing limits

&#x2022; [`total_number`](#total-number) - Optional Number<br>Number Of Requests. The total number of allowed requests per rate-limiting period

&#x2022; [`unit`](#unit) - Optional String  Defaults to `SECOND`<br>Possible values are `SECOND`, `MINUTE`, `HOUR`<br>Rate Limit Period Unit. Unit for the period per which the rate limit is applied. - SECOND: Second Rate limit period unit is seconds - MINUTE: Minute Rate limit period unit is minutes - HOUR: Hour Rate limit period unit is hours - DAY: Day Rate limit period unit is days

#### Rate Limit Rate Limiter Action Block

An [`action_block`](#rate-limit-rate-limiter-action-block) block (within [`rate_limit.rate_limiter`](#rate-limit-rate-limiter)) supports the following:

&#x2022; [`hours`](#hours) - Optional Block<br>Hours. Input Duration Hours<br>See [Hours](#rate-limit-rate-limiter-action-block-hours) below.

&#x2022; [`minutes`](#minutes) - Optional Block<br>Minutes. Input Duration Minutes<br>See [Minutes](#rate-limit-rate-limiter-action-block-minutes) below.

&#x2022; [`seconds`](#seconds) - Optional Block<br>Seconds. Input Duration Seconds<br>See [Seconds](#rate-limit-rate-limiter-action-block-seconds) below.

#### Rate Limit Rate Limiter Action Block Hours

A [`hours`](#rate-limit-rate-limiter-action-block-hours) block (within [`rate_limit.rate_limiter.action_block`](#rate-limit-rate-limiter-action-block)) supports the following:

&#x2022; [`duration`](#duration) - Optional Number<br>Duration

#### Rate Limit Rate Limiter Action Block Minutes

A [`minutes`](#rate-limit-rate-limiter-action-block-minutes) block (within [`rate_limit.rate_limiter.action_block`](#rate-limit-rate-limiter-action-block)) supports the following:

&#x2022; [`duration`](#duration) - Optional Number<br>Duration

#### Rate Limit Rate Limiter Action Block Seconds

A [`seconds`](#rate-limit-rate-limiter-action-block-seconds) block (within [`rate_limit.rate_limiter.action_block`](#rate-limit-rate-limiter-action-block)) supports the following:

&#x2022; [`duration`](#duration) - Optional Number<br>Duration

#### Ring Hash

A [`ring_hash`](#ring-hash) block supports the following:

&#x2022; [`hash_policy`](#hash-policy) - Optional Block<br>Hash Policy. Specifies a list of hash policies to use for ring hash load balancing. Each hash policy is evaluated individually and the combined result is used to route the request<br>See [Hash Policy](#ring-hash-hash-policy) below.

#### Ring Hash Hash Policy

A [`hash_policy`](#ring-hash-hash-policy) block (within [`ring_hash`](#ring-hash)) supports the following:

&#x2022; [`cookie`](#cookie) - Optional Block<br>Hashing using Cookie. Two types of cookie affinity: 1. Passive. Takes a cookie that's present in the cookies header and hashes on its value. 2. Generated. Generates and sets a cookie with an expiration (TTL) on the first request from the client in its response to the client, based on the endpoint the request gets sent to. The client then presents this on the next and all subsequent requests. The hash of this is sufficient to ensure these requests get sent to the same endpoint. The cookie is generated by hashing the source and destination ports and addresses so that multiple independent HTTP2 streams on the same connection will independently receive the same cookie, even if they arrive simultaneously<br>See [Cookie](#ring-hash-hash-policy-cookie) below.

&#x2022; [`header_name`](#header-name) - Optional String<br>Header Name. The name or key of the request header that will be used to obtain the hash key

&#x2022; [`source_ip`](#source-ip) - Optional Bool<br>Source IP. Hash based on source IP address

&#x2022; [`terminal`](#terminal) - Optional Bool<br>Terminal. Specify if its a terminal policy

#### Ring Hash Hash Policy Cookie

A [`cookie`](#ring-hash-hash-policy-cookie) block (within [`ring_hash.hash_policy`](#ring-hash-hash-policy)) supports the following:

&#x2022; [`add_httponly`](#add-httponly) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`add_secure`](#add-secure) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`ignore_httponly`](#ignore-httponly) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`ignore_samesite`](#ignore-samesite) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`ignore_secure`](#ignore-secure) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`name`](#name) - Optional String<br>Name. The name of the cookie that will be used to obtain the hash key. If the cookie is not present and TTL below is not set, no hash will be produced

&#x2022; [`path`](#path) - Optional String<br>Path. The name of the path for the cookie. If no path is specified here, no path will be set for the cookie

&#x2022; [`samesite_lax`](#samesite-lax) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`samesite_none`](#samesite-none) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`samesite_strict`](#samesite-strict) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`ttl`](#ttl) - Optional Number<br>TTL. If specified, a cookie with the TTL will be generated if the cookie is not present. If the TTL is present and zero, the generated cookie will be a session cookie. TTL value is in milliseconds

#### Routes

A [`routes`](#routes) block supports the following:

&#x2022; [`custom_route_object`](#custom-route-object) - Optional Block<br>Custom Route Object. A custom route uses a route object created outside of this view<br>See [Custom Route Object](#routes-custom-route-object) below.

&#x2022; [`direct_response_route`](#direct-response-route) - Optional Block<br>Direct Response Route. A direct response route matches on path, incoming header, incoming port and/or HTTP method and responds directly to the matching traffic<br>See [Direct Response Route](#routes-direct-response-route) below.

&#x2022; [`redirect_route`](#redirect-route) - Optional Block<br>Redirect Route. A redirect route matches on path, incoming header, incoming port and/or HTTP method and redirects the matching traffic to a different URL<br>See [Redirect Route](#routes-redirect-route) below.

&#x2022; [`simple_route`](#simple-route) - Optional Block<br>Simple Route. A simple route matches on path, incoming header, incoming port and/or HTTP method and forwards the matching traffic to the associated pools<br>See [Simple Route](#routes-simple-route) below.

#### Routes Custom Route Object

A [`custom_route_object`](#routes-custom-route-object) block (within [`routes`](#routes)) supports the following:

&#x2022; [`route_ref`](#route-ref) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Route Ref](#routes-custom-route-object-route-ref) below.

#### Routes Custom Route Object Route Ref

A [`route_ref`](#routes-custom-route-object-route-ref) block (within [`routes.custom_route_object`](#routes-custom-route-object)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Routes Direct Response Route

A [`direct_response_route`](#routes-direct-response-route) block (within [`routes`](#routes)) supports the following:

&#x2022; [`headers`](#headers) - Optional Block<br>Headers. List of (key, value) headers<br>See [Headers](#routes-direct-response-route-headers) below.

&#x2022; [`http_method`](#http-method) - Optional String  Defaults to `ANY`<br>Possible values are `ANY`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`, `COPY`<br>HTTP Method. Specifies the HTTP method used to access a resource. Any HTTP Method

&#x2022; [`incoming_port`](#incoming-port) - Optional Block<br>Port to Match. Port match of the request can be a range or a specific port<br>See [Incoming Port](#routes-direct-response-route-incoming-port) below.

&#x2022; [`path`](#path) - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#routes-direct-response-route-path) below.

&#x2022; [`route_direct_response`](#route-direct-response) - Optional Block<br>Direct Response. Send this direct response in case of route match action is direct response<br>See [Route Direct Response](#routes-direct-response-route-route-direct-response) below.

#### Routes Direct Response Route Headers

A [`headers`](#routes-direct-response-route-headers) block (within [`routes.direct_response_route`](#routes-direct-response-route)) supports the following:

&#x2022; [`exact`](#exact) - Optional String<br>Exact. Header value to match exactly

&#x2022; [`invert_match`](#invert-match) - Optional Bool<br>NOT of match. Invert the result of the match to detect missing header or non-matching value

&#x2022; [`name`](#name) - Optional String<br>Name. Name of the header

&#x2022; [`presence`](#presence) - Optional Bool<br>Presence. If true, check for presence of header

&#x2022; [`regex`](#regex) - Optional String<br>Regex. Regex match of the header value in re2 format

#### Routes Direct Response Route Incoming Port

An [`incoming_port`](#routes-direct-response-route-incoming-port) block (within [`routes.direct_response_route`](#routes-direct-response-route)) supports the following:

&#x2022; [`no_port_match`](#no-port-match) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`port`](#port) - Optional Number<br>Port. Exact Port to match

&#x2022; [`port_ranges`](#port-ranges) - Optional String<br>Port range. Port range to match

#### Routes Direct Response Route Path

A [`path`](#routes-direct-response-route-path) block (within [`routes.direct_response_route`](#routes-direct-response-route)) supports the following:

&#x2022; [`path`](#path) - Optional String<br>Exact. Exact path value to match

&#x2022; [`prefix`](#prefix) - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

&#x2022; [`regex`](#regex) - Optional String<br>Regex. Regular expression of path match (e.g. the value .* will match on all paths)

#### Routes Direct Response Route Route Direct Response

A [`route_direct_response`](#routes-direct-response-route-route-direct-response) block (within [`routes.direct_response_route`](#routes-direct-response-route)) supports the following:

&#x2022; [`response_body_encoded`](#response-body-encoded) - Optional String<br>Response Body. Response body to send. Currently supported URL schemes is string:/// for which message should be encoded in Base64 format. The message can be either plain text or HTML. E.g. '<p> Access Denied </p>'. Base64 encoded string URL for this is string:///PHA+IEFjY2VzcyBEZW5pZWQgPC9wPg==

&#x2022; [`response_code`](#response-code) - Optional Number<br>Response Code. response code to send

#### Routes Redirect Route

A [`redirect_route`](#routes-redirect-route) block (within [`routes`](#routes)) supports the following:

&#x2022; [`headers`](#headers) - Optional Block<br>Headers. List of (key, value) headers<br>See [Headers](#routes-redirect-route-headers) below.

&#x2022; [`http_method`](#http-method) - Optional String  Defaults to `ANY`<br>Possible values are `ANY`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`, `COPY`<br>HTTP Method. Specifies the HTTP method used to access a resource. Any HTTP Method

&#x2022; [`incoming_port`](#incoming-port) - Optional Block<br>Port to Match. Port match of the request can be a range or a specific port<br>See [Incoming Port](#routes-redirect-route-incoming-port) below.

&#x2022; [`path`](#path) - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#routes-redirect-route-path) below.

&#x2022; [`route_redirect`](#route-redirect) - Optional Block<br>Redirect. route redirect parameters when match action is redirect<br>See [Route Redirect](#routes-redirect-route-route-redirect) below.

#### Routes Redirect Route Headers

A [`headers`](#routes-redirect-route-headers) block (within [`routes.redirect_route`](#routes-redirect-route)) supports the following:

&#x2022; [`exact`](#exact) - Optional String<br>Exact. Header value to match exactly

&#x2022; [`invert_match`](#invert-match) - Optional Bool<br>NOT of match. Invert the result of the match to detect missing header or non-matching value

&#x2022; [`name`](#name) - Optional String<br>Name. Name of the header

&#x2022; [`presence`](#presence) - Optional Bool<br>Presence. If true, check for presence of header

&#x2022; [`regex`](#regex) - Optional String<br>Regex. Regex match of the header value in re2 format

#### Routes Redirect Route Incoming Port

An [`incoming_port`](#routes-redirect-route-incoming-port) block (within [`routes.redirect_route`](#routes-redirect-route)) supports the following:

&#x2022; [`no_port_match`](#no-port-match) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`port`](#port) - Optional Number<br>Port. Exact Port to match

&#x2022; [`port_ranges`](#port-ranges) - Optional String<br>Port range. Port range to match

#### Routes Redirect Route Path

A [`path`](#routes-redirect-route-path) block (within [`routes.redirect_route`](#routes-redirect-route)) supports the following:

&#x2022; [`path`](#path) - Optional String<br>Exact. Exact path value to match

&#x2022; [`prefix`](#prefix) - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

&#x2022; [`regex`](#regex) - Optional String<br>Regex. Regular expression of path match (e.g. the value .* will match on all paths)

#### Routes Redirect Route Route Redirect

A [`route_redirect`](#routes-redirect-route-route-redirect) block (within [`routes.redirect_route`](#routes-redirect-route)) supports the following:

&#x2022; [`host_redirect`](#host-redirect) - Optional String<br>Host. swap host part of incoming URL in redirect URL

&#x2022; [`path_redirect`](#path-redirect) - Optional String<br>Path. swap path part of incoming URL in redirect URL

&#x2022; [`prefix_rewrite`](#prefix-rewrite) - Optional String<br>Prefix Rewrite. In Redirect response, the matched prefix (or path) should be swapped with this value. This option allows redirect URLs be dynamically created based on the request

&#x2022; [`proto_redirect`](#proto-redirect) - Optional String<br>Protocol. swap protocol part of incoming URL in redirect URL The protocol can be swapped with either HTTP or HTTPS When incoming-proto option is specified, swapping of protocol is not done

&#x2022; [`remove_all_params`](#remove-all-params) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`replace_params`](#replace-params) - Optional String<br>Replace All Parameters

&#x2022; [`response_code`](#response-code) - Optional Number<br>Response Code. The HTTP status code to use in the redirect response

&#x2022; [`retain_all_params`](#retain-all-params) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### Routes Simple Route

A [`simple_route`](#routes-simple-route) block (within [`routes`](#routes)) supports the following:

&#x2022; [`advanced_options`](#advanced-options) - Optional Block<br>Advanced Route Options. Configure advanced options for route like path rewrite, hash policy, etc<br>See [Advanced Options](#routes-simple-route-advanced-options) below.

&#x2022; [`auto_host_rewrite`](#auto-host-rewrite) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`disable_host_rewrite`](#disable-host-rewrite) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`headers`](#headers) - Optional Block<br>Headers. List of (key, value) headers<br>See [Headers](#routes-simple-route-headers) below.

&#x2022; [`host_rewrite`](#host-rewrite) - Optional String<br>Host Rewrite Value. Host header will be swapped with this value

&#x2022; [`http_method`](#http-method) - Optional String  Defaults to `ANY`<br>Possible values are `ANY`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`, `COPY`<br>HTTP Method. Specifies the HTTP method used to access a resource. Any HTTP Method

&#x2022; [`incoming_port`](#incoming-port) - Optional Block<br>Port to Match. Port match of the request can be a range or a specific port<br>See [Incoming Port](#routes-simple-route-incoming-port) below.

&#x2022; [`origin_pools`](#origin-pools) - Optional Block<br>Origin Pools. Origin Pools for this route<br>See [Origin Pools](#routes-simple-route-origin-pools) below.

&#x2022; [`path`](#path) - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#routes-simple-route-path) below.

&#x2022; [`query_params`](#query-params) - Optional Block<br>Query Parameters. Handling of incoming query parameters in simple route<br>See [Query Params](#routes-simple-route-query-params) below.

#### Routes Simple Route Advanced Options

An [`advanced_options`](#routes-simple-route-advanced-options) block (within [`routes.simple_route`](#routes-simple-route)) supports the following:

&#x2022; [`app_firewall`](#app-firewall) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [App Firewall](#routes-simple-route-advanced-options-app-firewall) below.

&#x2022; [`bot_defense_javascript_injection`](#bot-defense-javascript-injection) - Optional Block<br>Bot Defense Javascript Injection Configuration for inline deployments. Bot Defense Javascript Injection Configuration for inline bot defense deployments<br>See [Bot Defense Javascript Injection](#routes-simple-route-advanced-options-bot-defense-javascript-injection) below.

&#x2022; [`buffer_policy`](#buffer-policy) - Optional Block<br>Buffer Configuration. Some upstream applications are not capable of handling streamed data. This config enables buffering the entire request before sending to upstream application. We can specify the maximum buffer size and buffer interval with this config. Buffering can be enabled and disabled at VirtualHost and Route levels Route level buffer configuration takes precedence<br>See [Buffer Policy](#routes-simple-route-advanced-options-buffer-policy) below.

&#x2022; [`common_buffering`](#common-buffering) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`common_hash_policy`](#common-hash-policy) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`cors_policy`](#cors-policy) - Optional Block<br>CORS Policy. Cross-Origin Resource Sharing requests configuration specified at Virtual-host or Route level. Route level configuration takes precedence. An example of an Cross origin HTTP request GET /resources/public-data/ HTTP/1.1 Host: bar.other User-Agent: Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.5; en-US; rv:1.9.1b3pre) Gecko/20081130 Minefield/3.1b3pre Accept: text/HTML,application/xhtml+XML,application/XML;q=0.9,*/*;q=0.8 Accept-Language: en-us,en;q=0.5 Accept-Encoding: gzip,deflate Accept-Charset: ISO-8859-1,utf-8;q=0.7,*;q=0.7 Connection: keep-alive Referrer: `HTTP://foo.example/examples/access-control/simpleXSInvocation.HTML` Origin: `HTTP://foo.example` HTTP/1.1 200 OK Date: Mon, 01 Dec 2008 00:23:53 GMT Server: Apache/2.0.61 Access-Control-Allow-Origin: * Keep-Alive: timeout=2, max=100 Connection: Keep-Alive Transfer-Encoding: chunked Content-Type: application/XML An example for cross origin HTTP OPTIONS request with Access-Control-Request-* header OPTIONS /resources/post-here/ HTTP/1.1 Host: bar.other User-Agent: Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.5; en-US; rv:1.9.1b3pre) Gecko/20081130 Minefield/3.1b3pre Accept: text/HTML,application/xhtml+XML,application/XML;q=0.9,*/*;q=0.8 Accept-Language: en-us,en;q=0.5 Accept-Encoding: gzip,deflate Accept-Charset: ISO-8859-1,utf-8;q=0.7,*;q=0.7 Connection: keep-alive Origin: `HTTP://foo.example` Access-Control-Request-Method: POST Access-Control-Request-Headers: X-PINGOTHER, Content-Type HTTP/1.1 204 No Content Date: Mon, 01 Dec 2008 01:15:39 GMT Server: Apache/2.0.61 (Unix) Access-Control-Allow-Origin: `HTTP://foo.example` Access-Control-Allow-Methods: POST, GET, OPTIONS Access-Control-Allow-Headers: X-PINGOTHER, Content-Type Access-Control-Max-Age: 86400 Vary: Accept-Encoding, Origin Keep-Alive: timeout=2, max=100 Connection: Keep-Alive<br>See [CORS Policy](#routes-simple-route-advanced-options-cors-policy) below.

&#x2022; [`csrf_policy`](#csrf-policy) - Optional Block<br>CSRF Policy. To mitigate CSRF attack , the policy checks where a request is coming from to determine if the request's origin is the same as its detination.The policy relies on two pieces of information used in determining if a request originated from the same host. 1. The origin that caused the user agent to issue the request (source origin). 2. The origin that the request is going to (target origin). When the policy evaluating a request, it ensures both pieces of information are present and compare their values. If the source origin is missing or origins do not match the request is rejected. The exception to this being if the source-origin has been added to they policy as valid. Because CSRF attacks specifically target state-changing requests, the policy only acts on the HTTP requests that have state-changing method (PUT,POST, etc.)<br>See [CSRF Policy](#routes-simple-route-advanced-options-csrf-policy) below.

&#x2022; [`default_retry_policy`](#default-retry-policy) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`disable_location_add`](#disable-location-add) - Optional Bool<br>Disable Location Addition. disables append of x-volterra-location = <RE-site-name> at route level, if it is configured at virtual-host level. This configuration is ignored on CE sites

&#x2022; [`disable_mirroring`](#disable-mirroring) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`disable_prefix_rewrite`](#disable-prefix-rewrite) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`disable_spdy`](#disable-spdy) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`disable_waf`](#disable-waf) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`disable_web_socket_config`](#disable-web-socket-config) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`do_not_retract_cluster`](#do-not-retract-cluster) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`enable_spdy`](#enable-spdy) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`endpoint_subsets`](#endpoint-subsets) - Optional Block<br>Origin Servers Subsets. Upstream origin pool may be configured to divide its origin servers into subsets based on metadata attached to the origin servers. Routes may then specify the metadata that a endpoint must match in order to be selected by the load balancer For origin servers which are discovered in K8S or Consul cluster, the label of the service is merged with endpoint's labels. In case of Consul, the label is derived from the 'Tag' field. For labels that are common between configured endpoint and discovered service, labels from discovered service takes precedence. List of key-value pairs that will be used as matching metadata. Only those origin servers of upstream origin pool which match this metadata will be selected for load balancing

&#x2022; [`inherited_bot_defense_javascript_injection`](#inherited-bot-defense-javascript-injection) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`inherited_waf`](#inherited-waf) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`inherited_waf_exclusion`](#inherited-waf-exclusion) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`mirror_policy`](#mirror-policy) - Optional Block<br>Mirror Policy. MirrorPolicy is used for shadowing traffic from one origin pool to another. The approach used is 'fire and forget', meaning it will not wait for the shadow origin pool to respond before returning the response from the primary origin pool. All normal statistics are collected for the shadow origin pool making this feature useful for testing and troubleshooting<br>See [Mirror Policy](#routes-simple-route-advanced-options-mirror-policy) below.

&#x2022; [`no_retry_policy`](#no-retry-policy) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`prefix_rewrite`](#prefix-rewrite) - Optional String<br>Enable Prefix Rewrite. prefix_rewrite indicates that during forwarding, the matched prefix (or path) should be swapped with its value. When using regex path matching, the entire path (not including the query string) will be swapped with this value

&#x2022; [`priority`](#priority) - Optional String  Defaults to `DEFAULT`<br>Possible values are `DEFAULT`, `HIGH`<br>Routing Priority. Priority routing for each request. Different connection pools are used based on the priority selected for the request. Also, circuit-breaker configuration at destination cluster is chosen based on selected priority. Default routing mechanism High-Priority routing mechanism

&#x2022; [`regex_rewrite`](#regex-rewrite) - Optional Block<br>Regex Match Rewrite. RegexMatchRewrite describes how to match a string and then produce a new string using a regular expression and a substitution string<br>See [Regex Rewrite](#routes-simple-route-advanced-options-regex-rewrite) below.

&#x2022; [`request_cookies_to_add`](#request-cookies-to-add) - Optional Block<br>Add Cookies in Cookie Header. Cookies are key-value pairs to be added to HTTP request being routed towards upstream. Cookies specified at this level are applied after cookies from matched Route are applied<br>See [Request Cookies To Add](#routes-simple-route-advanced-options-request-cookies-to-add) below.

&#x2022; [`request_cookies_to_remove`](#request-cookies-to-remove) - Optional List<br>Remove Cookies from Cookie Header. List of keys of Cookies to be removed from the HTTP request being sent towards upstream

&#x2022; [`request_headers_to_add`](#request-headers-to-add) - Optional Block<br>Add Request Headers. Headers are key-value pairs to be added to HTTP request being routed towards upstream<br>See [Request Headers To Add](#routes-simple-route-advanced-options-request-headers-to-add) below.

&#x2022; [`request_headers_to_remove`](#request-headers-to-remove) - Optional List<br>Remove Request Headers. List of keys of Headers to be removed from the HTTP request being sent towards upstream

&#x2022; [`response_cookies_to_add`](#response-cookies-to-add) - Optional Block<br>Add Set-Cookie Headers. Cookies are name-value pairs along with optional attribute parameters to be added to HTTP response being sent towards downstream. Cookies specified at this level are applied after cookies from matched Route are applied<br>See [Response Cookies To Add](#routes-simple-route-advanced-options-response-cookies-to-add) below.

&#x2022; [`response_cookies_to_remove`](#response-cookies-to-remove) - Optional List<br>Remove Cookies from Set-Cookie Headers. List of name of Cookies to be removed from the HTTP response being sent towards downstream. Entire set-cookie header will be removed

&#x2022; [`response_headers_to_add`](#response-headers-to-add) - Optional Block<br>Add Response Headers. Headers are key-value pairs to be added to HTTP response being sent towards downstream<br>See [Response Headers To Add](#routes-simple-route-advanced-options-response-headers-to-add) below.

&#x2022; [`response_headers_to_remove`](#response-headers-to-remove) - Optional List<br>Remove Response Headers. List of keys of Headers to be removed from the HTTP response being sent towards downstream

&#x2022; [`retract_cluster`](#retract-cluster) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`retry_policy`](#retry-policy) - Optional Block<br>Retry Policy. Retry policy configuration for route destination<br>See [Retry Policy](#routes-simple-route-advanced-options-retry-policy) below.

&#x2022; [`specific_hash_policy`](#specific-hash-policy) - Optional Block<br>Hash Policy List. List of hash policy rules<br>See [Specific Hash Policy](#routes-simple-route-advanced-options-specific-hash-policy) below.

&#x2022; [`timeout`](#timeout) - Optional Number<br>Timeout. The timeout for the route including all retries, in milliseconds. Should be set to a high value or 0 (infinite timeout) for server-side streaming

&#x2022; [`waf_exclusion_policy`](#waf-exclusion-policy) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [WAF Exclusion Policy](#routes-simple-route-advanced-options-waf-exclusion-policy) below.

&#x2022; [`web_socket_config`](#web-socket-config) - Optional Block<br>WebSocket Configuration. Configuration to allow WebSocket Request headers of such upgrade looks like below 'connection', 'Upgrade' 'upgrade', 'WebSocket' With configuration to allow WebSocket upgrade, ADC will produce following response 'HTTP/1.1 101 Switching Protocols 'Upgrade': 'WebSocket' 'Connection': 'Upgrade'<br>See [Web Socket Config](#routes-simple-route-advanced-options-web-socket-config) below.

#### Routes Simple Route Advanced Options App Firewall

An [`app_firewall`](#routes-simple-route-advanced-options-app-firewall) block (within [`routes.simple_route.advanced_options`](#routes-simple-route-advanced-options)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Routes Simple Route Advanced Options Bot Defense Javascript Injection

A [`bot_defense_javascript_injection`](#routes-simple-route-advanced-options-bot-defense-javascript-injection) block (within [`routes.simple_route.advanced_options`](#routes-simple-route-advanced-options)) supports the following:

&#x2022; [`javascript_location`](#javascript-location) - Optional String  Defaults to `AFTER_HEAD`<br>Possible values are `AFTER_HEAD`, `AFTER_TITLE_END`, `BEFORE_SCRIPT`<br>JavaScript Location. All inside networks. Insert JavaScript after <head> tag Insert JavaScript after </title> tag. Insert JavaScript before first <script> tag

&#x2022; [`javascript_tags`](#javascript-tags) - Optional Block<br>JavaScript Tags. Select Add item to configure your javascript tag. If adding both Bot Adv and Fraud, the Bot Javascript should be added first<br>See [Javascript Tags](#routes-simple-route-advanced-options-bot-defense-javascript-injection-javascript-tags) below.

#### Routes Simple Route Advanced Options Bot Defense Javascript Injection Javascript Tags

A [`javascript_tags`](#routes-simple-route-advanced-options-bot-defense-javascript-injection-javascript-tags) block (within [`routes.simple_route.advanced_options.bot_defense_javascript_injection`](#routes-simple-route-advanced-options-bot-defense-javascript-injection)) supports the following:

&#x2022; [`javascript_url`](#javascript-url) - Optional String<br>URL. Please enter the full URL (include domain and path), or relative path

&#x2022; [`tag_attributes`](#tag-attributes) - Optional Block<br>Tag Attributes. Add the tag attributes you want to include in your Javascript tag<br>See [Tag Attributes](#routes-simple-route-advanced-options-bot-defense-javascript-injection-javascript-tags-tag-attributes) below.

#### Routes Simple Route Advanced Options Bot Defense Javascript Injection Javascript Tags Tag Attributes

A [`tag_attributes`](#routes-simple-route-advanced-options-bot-defense-javascript-injection-javascript-tags-tag-attributes) block (within [`routes.simple_route.advanced_options.bot_defense_javascript_injection.javascript_tags`](#routes-simple-route-advanced-options-bot-defense-javascript-injection-javascript-tags)) supports the following:

&#x2022; [`javascript_tag`](#javascript-tag) - Optional String  Defaults to `JS_ATTR_ID`<br>Possible values are `JS_ATTR_ID`, `JS_ATTR_CID`, `JS_ATTR_CN`, `JS_ATTR_API_DOMAIN`, `JS_ATTR_API_URL`, `JS_ATTR_API_PATH`, `JS_ATTR_ASYNC`, `JS_ATTR_DEFER`<br>Tag Attribute Name. Select from one of the predefined tag attributes

&#x2022; [`tag_value`](#tag-value) - Optional String<br>Value. Add the tag attribute value

#### Routes Simple Route Advanced Options Buffer Policy

A [`buffer_policy`](#routes-simple-route-advanced-options-buffer-policy) block (within [`routes.simple_route.advanced_options`](#routes-simple-route-advanced-options)) supports the following:

&#x2022; [`disabled`](#disabled) - Optional Bool<br>Disable. Disable buffering for a particular route. This is useful when virtual-host has buffering, but we need to disable it on a specific route. The value of this field is ignored for virtual-host

&#x2022; [`max_request_bytes`](#max-request-bytes) - Optional Number<br>Max Request Bytes. The maximum request size that the filter will buffer before the connection manager will stop buffering and return a RequestEntityTooLarge (413) response

#### Routes Simple Route Advanced Options CORS Policy

A [`cors_policy`](#routes-simple-route-advanced-options-cors-policy) block (within [`routes.simple_route.advanced_options`](#routes-simple-route-advanced-options)) supports the following:

&#x2022; [`allow_credentials`](#allow-credentials) - Optional Bool<br>Allow Credentials. Specifies whether the resource allows credentials

&#x2022; [`allow_headers`](#allow-headers) - Optional String<br>Allow Headers. Specifies the content for the access-control-allow-headers header

&#x2022; [`allow_methods`](#allow-methods) - Optional String<br>Allow Methods. Specifies the content for the access-control-allow-methods header

&#x2022; [`allow_origin`](#allow-origin) - Optional List<br>Allow Origin. Specifies the origins that will be allowed to do CORS requests. An origin is allowed if either allow_origin or allow_origin_regex match

&#x2022; [`allow_origin_regex`](#allow-origin-regex) - Optional List<br>Allow Origin Regex. Specifies regex patterns that match allowed origins. An origin is allowed if either allow_origin or allow_origin_regex match

&#x2022; [`disabled`](#disabled) - Optional Bool<br>Disabled. Disable the CorsPolicy for a particular route. This is useful when virtual-host has CorsPolicy, but we need to disable it on a specific route. The value of this field is ignored for virtual-host

&#x2022; [`expose_headers`](#expose-headers) - Optional String<br>Expose Headers. Specifies the content for the access-control-expose-headers header

&#x2022; [`maximum_age`](#maximum-age) - Optional Number<br>Maximum Age. Specifies the content for the access-control-max-age header in seconds. This indicates the maximum number of seconds the results can be cached A value of -1 will disable caching. Maximum permitted value is 86400 seconds (24 hours)

#### Routes Simple Route Advanced Options CSRF Policy

A [`csrf_policy`](#routes-simple-route-advanced-options-csrf-policy) block (within [`routes.simple_route.advanced_options`](#routes-simple-route-advanced-options)) supports the following:

&#x2022; [`all_load_balancer_domains`](#all-load-balancer-domains) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`custom_domain_list`](#custom-domain-list) - Optional Block<br>Domain name list. List of domain names used for Host header matching<br>See [Custom Domain List](#routes-simple-route-advanced-options-csrf-policy-custom-domain-list) below.

&#x2022; [`disabled`](#disabled) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### Routes Simple Route Advanced Options CSRF Policy Custom Domain List

A [`custom_domain_list`](#routes-simple-route-advanced-options-csrf-policy-custom-domain-list) block (within [`routes.simple_route.advanced_options.csrf_policy`](#routes-simple-route-advanced-options-csrf-policy)) supports the following:

&#x2022; [`domains`](#domains) - Optional List<br>Domain names. A list of domain names that will be matched to loadbalancer. These domains are not used for SNI match. Wildcard names are supported in the suffix or prefix form

#### Routes Simple Route Advanced Options Mirror Policy

A [`mirror_policy`](#routes-simple-route-advanced-options-mirror-policy) block (within [`routes.simple_route.advanced_options`](#routes-simple-route-advanced-options)) supports the following:

&#x2022; [`origin_pool`](#origin-pool) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Origin Pool](#routes-simple-route-advanced-options-mirror-policy-origin-pool) below.

&#x2022; [`percent`](#percent) - Optional Block<br>Fractional Percent. Fraction used where sampling percentages are needed. example sampled requests<br>See [Percent](#routes-simple-route-advanced-options-mirror-policy-percent) below.

#### Routes Simple Route Advanced Options Mirror Policy Origin Pool

An [`origin_pool`](#routes-simple-route-advanced-options-mirror-policy-origin-pool) block (within [`routes.simple_route.advanced_options.mirror_policy`](#routes-simple-route-advanced-options-mirror-policy)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Routes Simple Route Advanced Options Mirror Policy Percent

A [`percent`](#routes-simple-route-advanced-options-mirror-policy-percent) block (within [`routes.simple_route.advanced_options.mirror_policy`](#routes-simple-route-advanced-options-mirror-policy)) supports the following:

&#x2022; [`denominator`](#denominator) - Optional String  Defaults to `HUNDRED`<br>Possible values are `HUNDRED`, `TEN_THOUSAND`, `MILLION`<br>Denominator. Denominator used in fraction where sampling percentages are needed. example sampled requests Use hundred as denominator Use ten thousand as denominator Use million as denominator

&#x2022; [`numerator`](#numerator) - Optional Number<br>Numerator. sampled parts per denominator. If denominator was 10000, then value of 5 will be 5 in 10000

#### Routes Simple Route Advanced Options Regex Rewrite

A [`regex_rewrite`](#routes-simple-route-advanced-options-regex-rewrite) block (within [`routes.simple_route.advanced_options`](#routes-simple-route-advanced-options)) supports the following:

&#x2022; [`pattern`](#pattern) - Optional String<br>Pattern. The regular expression used to find portions of a string that should be replaced

&#x2022; [`substitution`](#substitution) - Optional String<br>Substitution. The string that should be substituted into matching portions of the subject string during a substitution operation to produce a new string

#### Routes Simple Route Advanced Options Request Cookies To Add

A [`request_cookies_to_add`](#routes-simple-route-advanced-options-request-cookies-to-add) block (within [`routes.simple_route.advanced_options`](#routes-simple-route-advanced-options)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. Name of the cookie in Cookie header

&#x2022; [`overwrite`](#overwrite) - Optional Bool  Defaults to `do`<br>Overwrite. Should the value be overwritten? If true, the value is overwritten to existing values. not overwrite

&#x2022; [`secret_value`](#secret-value) - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Secret Value](#routes-simple-route-advanced-options-request-cookies-to-add-secret-value) below.

&#x2022; [`value`](#value) - Optional String<br>Value. Value of the Cookie header

#### Routes Simple Route Advanced Options Request Cookies To Add Secret Value

A [`secret_value`](#routes-simple-route-advanced-options-request-cookies-to-add-secret-value) block (within [`routes.simple_route.advanced_options.request_cookies_to_add`](#routes-simple-route-advanced-options-request-cookies-to-add)) supports the following:

&#x2022; [`blindfold_secret_info`](#blindfold-secret-info) - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#routes-simple-route-advanced-options-request-cookies-to-add-secret-value-blindfold-secret-info) below.

&#x2022; [`clear_secret_info`](#clear-secret-info) - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#routes-simple-route-advanced-options-request-cookies-to-add-secret-value-clear-secret-info) below.

#### Routes Simple Route Advanced Options Request Cookies To Add Secret Value Blindfold Secret Info

A [`blindfold_secret_info`](#routes-simple-route-advanced-options-request-cookies-to-add-secret-value-blindfold-secret-info) block (within [`routes.simple_route.advanced_options.request_cookies_to_add.secret_value`](#routes-simple-route-advanced-options-request-cookies-to-add-secret-value)) supports the following:

&#x2022; [`decryption_provider`](#decryption-provider) - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

&#x2022; [`location`](#location) - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

&#x2022; [`store_provider`](#store-provider) - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

#### Routes Simple Route Advanced Options Request Cookies To Add Secret Value Clear Secret Info

A [`clear_secret_info`](#routes-simple-route-advanced-options-request-cookies-to-add-secret-value-clear-secret-info) block (within [`routes.simple_route.advanced_options.request_cookies_to_add.secret_value`](#routes-simple-route-advanced-options-request-cookies-to-add-secret-value)) supports the following:

&#x2022; [`provider_ref`](#provider-ref) - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

&#x2022; [`url`](#url) - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

#### Routes Simple Route Advanced Options Request Headers To Add

A [`request_headers_to_add`](#routes-simple-route-advanced-options-request-headers-to-add) block (within [`routes.simple_route.advanced_options`](#routes-simple-route-advanced-options)) supports the following:

&#x2022; [`append`](#append) - Optional Bool  Defaults to `do`<br>Append. Should the value be appended? If true, the value is appended to existing values. not append

&#x2022; [`name`](#name) - Optional String<br>Name. Name of the HTTP header

&#x2022; [`secret_value`](#secret-value) - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Secret Value](#routes-simple-route-advanced-options-request-headers-to-add-secret-value) below.

&#x2022; [`value`](#value) - Optional String<br>Value. Value of the HTTP header

#### Routes Simple Route Advanced Options Request Headers To Add Secret Value

A [`secret_value`](#routes-simple-route-advanced-options-request-headers-to-add-secret-value) block (within [`routes.simple_route.advanced_options.request_headers_to_add`](#routes-simple-route-advanced-options-request-headers-to-add)) supports the following:

&#x2022; [`blindfold_secret_info`](#blindfold-secret-info) - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#routes-simple-route-advanced-options-request-headers-to-add-secret-value-blindfold-secret-info) below.

&#x2022; [`clear_secret_info`](#clear-secret-info) - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#routes-simple-route-advanced-options-request-headers-to-add-secret-value-clear-secret-info) below.

#### Routes Simple Route Advanced Options Request Headers To Add Secret Value Blindfold Secret Info

A [`blindfold_secret_info`](#routes-simple-route-advanced-options-request-headers-to-add-secret-value-blindfold-secret-info) block (within [`routes.simple_route.advanced_options.request_headers_to_add.secret_value`](#routes-simple-route-advanced-options-request-headers-to-add-secret-value)) supports the following:

&#x2022; [`decryption_provider`](#decryption-provider) - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

&#x2022; [`location`](#location) - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

&#x2022; [`store_provider`](#store-provider) - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

#### Routes Simple Route Advanced Options Request Headers To Add Secret Value Clear Secret Info

A [`clear_secret_info`](#routes-simple-route-advanced-options-request-headers-to-add-secret-value-clear-secret-info) block (within [`routes.simple_route.advanced_options.request_headers_to_add.secret_value`](#routes-simple-route-advanced-options-request-headers-to-add-secret-value)) supports the following:

&#x2022; [`provider_ref`](#provider-ref) - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

&#x2022; [`url`](#url) - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

#### Routes Simple Route Advanced Options Response Cookies To Add

A [`response_cookies_to_add`](#routes-simple-route-advanced-options-response-cookies-to-add) block (within [`routes.simple_route.advanced_options`](#routes-simple-route-advanced-options)) supports the following:

&#x2022; [`add_domain`](#add-domain) - Optional String<br>Add Domain. Add domain attribute

&#x2022; [`add_expiry`](#add-expiry) - Optional String<br>Add expiry. Add expiry attribute

&#x2022; [`add_httponly`](#add-httponly) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`add_partitioned`](#add-partitioned) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`add_path`](#add-path) - Optional String<br>Add path. Add path attribute

&#x2022; [`add_secure`](#add-secure) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`ignore_domain`](#ignore-domain) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`ignore_expiry`](#ignore-expiry) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`ignore_httponly`](#ignore-httponly) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`ignore_max_age`](#ignore-max-age) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`ignore_partitioned`](#ignore-partitioned) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`ignore_path`](#ignore-path) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`ignore_samesite`](#ignore-samesite) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`ignore_secure`](#ignore-secure) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`ignore_value`](#ignore-value) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`max_age_value`](#max-age-value) - Optional Number<br>Add Max Age. Add max age attribute

&#x2022; [`name`](#name) - Optional String<br>Name. Name of the cookie in Cookie header

&#x2022; [`overwrite`](#overwrite) - Optional Bool  Defaults to `do`<br>Overwrite. Should the value be overwritten? If true, the value is overwritten to existing values. not overwrite

&#x2022; [`samesite_lax`](#samesite-lax) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`samesite_none`](#samesite-none) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`samesite_strict`](#samesite-strict) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`secret_value`](#secret-value) - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Secret Value](#routes-simple-route-advanced-options-response-cookies-to-add-secret-value) below.

&#x2022; [`value`](#value) - Optional String<br>Value. Value of the Cookie header

#### Routes Simple Route Advanced Options Response Cookies To Add Secret Value

A [`secret_value`](#routes-simple-route-advanced-options-response-cookies-to-add-secret-value) block (within [`routes.simple_route.advanced_options.response_cookies_to_add`](#routes-simple-route-advanced-options-response-cookies-to-add)) supports the following:

&#x2022; [`blindfold_secret_info`](#blindfold-secret-info) - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#routes-simple-route-advanced-options-response-cookies-to-add-secret-value-blindfold-secret-info) below.

&#x2022; [`clear_secret_info`](#clear-secret-info) - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#routes-simple-route-advanced-options-response-cookies-to-add-secret-value-clear-secret-info) below.

#### Routes Simple Route Advanced Options Response Cookies To Add Secret Value Blindfold Secret Info

A [`blindfold_secret_info`](#routes-simple-route-advanced-options-response-cookies-to-add-secret-value-blindfold-secret-info) block (within [`routes.simple_route.advanced_options.response_cookies_to_add.secret_value`](#routes-simple-route-advanced-options-response-cookies-to-add-secret-value)) supports the following:

&#x2022; [`decryption_provider`](#decryption-provider) - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

&#x2022; [`location`](#location) - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

&#x2022; [`store_provider`](#store-provider) - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

#### Routes Simple Route Advanced Options Response Cookies To Add Secret Value Clear Secret Info

A [`clear_secret_info`](#routes-simple-route-advanced-options-response-cookies-to-add-secret-value-clear-secret-info) block (within [`routes.simple_route.advanced_options.response_cookies_to_add.secret_value`](#routes-simple-route-advanced-options-response-cookies-to-add-secret-value)) supports the following:

&#x2022; [`provider_ref`](#provider-ref) - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

&#x2022; [`url`](#url) - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

#### Routes Simple Route Advanced Options Response Headers To Add

A [`response_headers_to_add`](#routes-simple-route-advanced-options-response-headers-to-add) block (within [`routes.simple_route.advanced_options`](#routes-simple-route-advanced-options)) supports the following:

&#x2022; [`append`](#append) - Optional Bool  Defaults to `do`<br>Append. Should the value be appended? If true, the value is appended to existing values. not append

&#x2022; [`name`](#name) - Optional String<br>Name. Name of the HTTP header

&#x2022; [`secret_value`](#secret-value) - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Secret Value](#routes-simple-route-advanced-options-response-headers-to-add-secret-value) below.

&#x2022; [`value`](#value) - Optional String<br>Value. Value of the HTTP header

#### Routes Simple Route Advanced Options Response Headers To Add Secret Value

A [`secret_value`](#routes-simple-route-advanced-options-response-headers-to-add-secret-value) block (within [`routes.simple_route.advanced_options.response_headers_to_add`](#routes-simple-route-advanced-options-response-headers-to-add)) supports the following:

&#x2022; [`blindfold_secret_info`](#blindfold-secret-info) - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#routes-simple-route-advanced-options-response-headers-to-add-secret-value-blindfold-secret-info) below.

&#x2022; [`clear_secret_info`](#clear-secret-info) - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#routes-simple-route-advanced-options-response-headers-to-add-secret-value-clear-secret-info) below.

#### Routes Simple Route Advanced Options Response Headers To Add Secret Value Blindfold Secret Info

A [`blindfold_secret_info`](#routes-simple-route-advanced-options-response-headers-to-add-secret-value-blindfold-secret-info) block (within [`routes.simple_route.advanced_options.response_headers_to_add.secret_value`](#routes-simple-route-advanced-options-response-headers-to-add-secret-value)) supports the following:

&#x2022; [`decryption_provider`](#decryption-provider) - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

&#x2022; [`location`](#location) - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

&#x2022; [`store_provider`](#store-provider) - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

#### Routes Simple Route Advanced Options Response Headers To Add Secret Value Clear Secret Info

A [`clear_secret_info`](#routes-simple-route-advanced-options-response-headers-to-add-secret-value-clear-secret-info) block (within [`routes.simple_route.advanced_options.response_headers_to_add.secret_value`](#routes-simple-route-advanced-options-response-headers-to-add-secret-value)) supports the following:

&#x2022; [`provider_ref`](#provider-ref) - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

&#x2022; [`url`](#url) - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

#### Routes Simple Route Advanced Options Retry Policy

A [`retry_policy`](#routes-simple-route-advanced-options-retry-policy) block (within [`routes.simple_route.advanced_options`](#routes-simple-route-advanced-options)) supports the following:

&#x2022; [`back_off`](#back-off) - Optional Block<br>Retry BackOff Interval. Specifies parameters that control retry back off<br>See [Back Off](#routes-simple-route-advanced-options-retry-policy-back-off) below.

&#x2022; [`num_retries`](#num-retries) - Optional Number  Defaults to `1`<br>Number of Retries. Specifies the allowed number of retries. Retries can be done any number of times. An exponential back-off algorithm is used between each retry

&#x2022; [`per_try_timeout`](#per-try-timeout) - Optional Number<br>Per Try Timeout. Specifies a non-zero timeout per retry attempt. In milliseconds

&#x2022; [`retriable_status_codes`](#retriable-status-codes) - Optional List<br>Status Code to Retry. HTTP status codes that should trigger a retry in addition to those specified by retry_on

&#x2022; [`retry_condition`](#retry-condition) - Optional List<br>Retry Condition. Specifies the conditions under which retry takes place. Retries can be on different types of condition depending on application requirements. For example, network failure, all 5xx response codes, idempotent 4xx response codes, etc The possible values are '5xx' : Retry will be done if the upstream server responds with any 5xx response code, or does not respond at all (disconnect/reset/read timeout). 'gateway-error' : Retry will be done only if the upstream server responds with 502, 503 or 504 responses (Included in 5xx) 'connect-failure' : Retry will be done if the request fails because of a connection failure to the upstream server (connect timeout, etc.). (Included in 5xx) 'refused-stream' : Retry is done if the upstream server resets the stream with a REFUSED_STREAM error code (Included in 5xx) 'retriable-4xx' : Retry is done if the upstream server responds with a retriable 4xx response code. The only response code in this category is HTTP CONFLICT (409) 'retriable-status-codes' : Retry is done if the upstream server responds with any response code matching one defined in retriable_status_codes field 'reset' : Retry is done if the upstream server does not respond at all (disconnect/reset/read timeout.)

#### Routes Simple Route Advanced Options Retry Policy Back Off

A [`back_off`](#routes-simple-route-advanced-options-retry-policy-back-off) block (within [`routes.simple_route.advanced_options.retry_policy`](#routes-simple-route-advanced-options-retry-policy)) supports the following:

&#x2022; [`base_interval`](#base-interval) - Optional Number<br>Base Retry Interval. Specifies the base interval between retries in milliseconds

&#x2022; [`max_interval`](#max-interval) - Optional Number  Defaults to `10`<br>Maximum Retry Interval. Specifies the maximum interval between retries in milliseconds. This parameter is optional, but must be greater than or equal to the base_interval if set. The times the base_interval

#### Routes Simple Route Advanced Options Specific Hash Policy

A [`specific_hash_policy`](#routes-simple-route-advanced-options-specific-hash-policy) block (within [`routes.simple_route.advanced_options`](#routes-simple-route-advanced-options)) supports the following:

&#x2022; [`hash_policy`](#hash-policy) - Optional Block<br>Hash Policy. Specifies a list of hash policies to use for ring hash load balancing. Each hash policy is evaluated individually and the combined result is used to route the request<br>See [Hash Policy](#routes-simple-route-advanced-options-specific-hash-policy-hash-policy) below.

#### Routes Simple Route Advanced Options Specific Hash Policy Hash Policy

A [`hash_policy`](#routes-simple-route-advanced-options-specific-hash-policy-hash-policy) block (within [`routes.simple_route.advanced_options.specific_hash_policy`](#routes-simple-route-advanced-options-specific-hash-policy)) supports the following:

&#x2022; [`cookie`](#cookie) - Optional Block<br>Hashing using Cookie. Two types of cookie affinity: 1. Passive. Takes a cookie that's present in the cookies header and hashes on its value. 2. Generated. Generates and sets a cookie with an expiration (TTL) on the first request from the client in its response to the client, based on the endpoint the request gets sent to. The client then presents this on the next and all subsequent requests. The hash of this is sufficient to ensure these requests get sent to the same endpoint. The cookie is generated by hashing the source and destination ports and addresses so that multiple independent HTTP2 streams on the same connection will independently receive the same cookie, even if they arrive simultaneously<br>See [Cookie](#routes-simple-route-advanced-options-specific-hash-policy-hash-policy-cookie) below.

&#x2022; [`header_name`](#header-name) - Optional String<br>Header Name. The name or key of the request header that will be used to obtain the hash key

&#x2022; [`source_ip`](#source-ip) - Optional Bool<br>Source IP. Hash based on source IP address

&#x2022; [`terminal`](#terminal) - Optional Bool<br>Terminal. Specify if its a terminal policy

#### Routes Simple Route Advanced Options Specific Hash Policy Hash Policy Cookie

A [`cookie`](#routes-simple-route-advanced-options-specific-hash-policy-hash-policy-cookie) block (within [`routes.simple_route.advanced_options.specific_hash_policy.hash_policy`](#routes-simple-route-advanced-options-specific-hash-policy-hash-policy)) supports the following:

&#x2022; [`add_httponly`](#add-httponly) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`add_secure`](#add-secure) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`ignore_httponly`](#ignore-httponly) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`ignore_samesite`](#ignore-samesite) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`ignore_secure`](#ignore-secure) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`name`](#name) - Optional String<br>Name. The name of the cookie that will be used to obtain the hash key. If the cookie is not present and TTL below is not set, no hash will be produced

&#x2022; [`path`](#path) - Optional String<br>Path. The name of the path for the cookie. If no path is specified here, no path will be set for the cookie

&#x2022; [`samesite_lax`](#samesite-lax) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`samesite_none`](#samesite-none) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`samesite_strict`](#samesite-strict) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`ttl`](#ttl) - Optional Number<br>TTL. If specified, a cookie with the TTL will be generated if the cookie is not present. If the TTL is present and zero, the generated cookie will be a session cookie. TTL value is in milliseconds

#### Routes Simple Route Advanced Options WAF Exclusion Policy

A [`waf_exclusion_policy`](#routes-simple-route-advanced-options-waf-exclusion-policy) block (within [`routes.simple_route.advanced_options`](#routes-simple-route-advanced-options)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Routes Simple Route Advanced Options Web Socket Config

A [`web_socket_config`](#routes-simple-route-advanced-options-web-socket-config) block (within [`routes.simple_route.advanced_options`](#routes-simple-route-advanced-options)) supports the following:

&#x2022; [`use_websocket`](#use-websocket) - Optional Bool<br>Use WebSocket. Specifies that the HTTP client connection to this route is allowed to upgrade to a WebSocket connection

#### Routes Simple Route Headers

A [`headers`](#routes-simple-route-headers) block (within [`routes.simple_route`](#routes-simple-route)) supports the following:

&#x2022; [`exact`](#exact) - Optional String<br>Exact. Header value to match exactly

&#x2022; [`invert_match`](#invert-match) - Optional Bool<br>NOT of match. Invert the result of the match to detect missing header or non-matching value

&#x2022; [`name`](#name) - Optional String<br>Name. Name of the header

&#x2022; [`presence`](#presence) - Optional Bool<br>Presence. If true, check for presence of header

&#x2022; [`regex`](#regex) - Optional String<br>Regex. Regex match of the header value in re2 format

#### Routes Simple Route Incoming Port

An [`incoming_port`](#routes-simple-route-incoming-port) block (within [`routes.simple_route`](#routes-simple-route)) supports the following:

&#x2022; [`no_port_match`](#no-port-match) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`port`](#port) - Optional Number<br>Port. Exact Port to match

&#x2022; [`port_ranges`](#port-ranges) - Optional String<br>Port range. Port range to match

#### Routes Simple Route Origin Pools

An [`origin_pools`](#routes-simple-route-origin-pools) block (within [`routes.simple_route`](#routes-simple-route)) supports the following:

&#x2022; [`cluster`](#cluster) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Cluster](#routes-simple-route-origin-pools-cluster) below.

&#x2022; [`endpoint_subsets`](#endpoint-subsets) - Optional Block<br>Origin Servers Subsets. Upstream origin pool may be configured to divide its origin servers into subsets based on metadata attached to the origin servers. Routes may then specify the metadata that a endpoint must match in order to be selected by the load balancer For origin servers which are discovered in K8S or Consul cluster, the label of the service is merged with endpoint's labels. In case of Consul, the label is derived from the 'Tag' field. For labels that are common between configured endpoint and discovered service, labels from discovered service takes precedence. List of key-value pairs that will be used as matching metadata. Only those origin servers of upstream origin pool which match this metadata will be selected for load balancing

&#x2022; [`pool`](#pool) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Pool](#routes-simple-route-origin-pools-pool) below.

&#x2022; [`priority`](#priority) - Optional Number<br>Priority. Priority of this origin pool, valid only with multiple origin pools. Value of 0 will make the pool as lowest priority origin pool Priority of 1 means highest priority and is considered active. When active origin pool is not available, lower priority origin pools are made active as per the increasing priority

&#x2022; [`weight`](#weight) - Optional Number<br>Weight. Weight of this origin pool, valid only with multiple origin pool. Value of 0 will disable the pool

#### Routes Simple Route Origin Pools Cluster

A [`cluster`](#routes-simple-route-origin-pools-cluster) block (within [`routes.simple_route.origin_pools`](#routes-simple-route-origin-pools)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Routes Simple Route Origin Pools Pool

A [`pool`](#routes-simple-route-origin-pools-pool) block (within [`routes.simple_route.origin_pools`](#routes-simple-route-origin-pools)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Routes Simple Route Path

A [`path`](#routes-simple-route-path) block (within [`routes.simple_route`](#routes-simple-route)) supports the following:

&#x2022; [`path`](#path) - Optional String<br>Exact. Exact path value to match

&#x2022; [`prefix`](#prefix) - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

&#x2022; [`regex`](#regex) - Optional String<br>Regex. Regular expression of path match (e.g. the value .* will match on all paths)

#### Routes Simple Route Query Params

A [`query_params`](#routes-simple-route-query-params) block (within [`routes.simple_route`](#routes-simple-route)) supports the following:

&#x2022; [`remove_all_params`](#remove-all-params) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`replace_params`](#replace-params) - Optional String<br>Replace All Parameters

&#x2022; [`retain_all_params`](#retain-all-params) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### Sensitive Data Disclosure Rules

A [`sensitive_data_disclosure_rules`](#sensitive-data-disclosure-rules) block supports the following:

&#x2022; [`sensitive_data_types_in_response`](#sensitive-data-types-in-response) - Optional Block<br>Sensitive Data Exposure Rules. Sensitive Data Exposure Rules allows specifying rules to mask sensitive data fields in API responses<br>See [Sensitive Data Types In Response](#sensitive-data-disclosure-rules-sensitive-data-types-in-response) below.

#### Sensitive Data Disclosure Rules Sensitive Data Types In Response

A [`sensitive_data_types_in_response`](#sensitive-data-disclosure-rules-sensitive-data-types-in-response) block (within [`sensitive_data_disclosure_rules`](#sensitive-data-disclosure-rules)) supports the following:

&#x2022; [`api_endpoint`](#api-endpoint) - Optional Block<br>API Endpoint. This defines API endpoint<br>See [API Endpoint](#sensitive-data-disclosure-rules-sensitive-data-types-in-response-api-endpoint) below.

&#x2022; [`body`](#body) - Optional Block<br>Body Section Masking Options. Options for HTTP Body Masking<br>See [Body](#sensitive-data-disclosure-rules-sensitive-data-types-in-response-body) below.

&#x2022; [`mask`](#mask) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`report`](#report) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### Sensitive Data Disclosure Rules Sensitive Data Types In Response API Endpoint

An [`api_endpoint`](#sensitive-data-disclosure-rules-sensitive-data-types-in-response-api-endpoint) block (within [`sensitive_data_disclosure_rules.sensitive_data_types_in_response`](#sensitive-data-disclosure-rules-sensitive-data-types-in-response)) supports the following:

&#x2022; [`methods`](#methods) - Optional List  Defaults to `ANY`<br>Possible values are `ANY`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`, `COPY`<br>Methods. Methods to be matched

&#x2022; [`path`](#path) - Optional String<br>Path. Path to be matched

#### Sensitive Data Disclosure Rules Sensitive Data Types In Response Body

A [`body`](#sensitive-data-disclosure-rules-sensitive-data-types-in-response-body) block (within [`sensitive_data_disclosure_rules.sensitive_data_types_in_response`](#sensitive-data-disclosure-rules-sensitive-data-types-in-response)) supports the following:

&#x2022; [`fields`](#fields) - Optional List<br>Values. List of JSON Path field values. Use square brackets with an underscore [_] to indicate array elements (e.g., person.emails[_]). To reference JSON keys that contain spaces, enclose the entire path in double quotes. For example: 'person.first name'

#### Sensitive Data Policy

A [`sensitive_data_policy`](#sensitive-data-policy) block supports the following:

&#x2022; [`sensitive_data_policy_ref`](#sensitive-data-policy-ref) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Sensitive Data Policy Ref](#sensitive-data-policy-sensitive-data-policy-ref) below.

#### Sensitive Data Policy Sensitive Data Policy Ref

A [`sensitive_data_policy_ref`](#sensitive-data-policy-sensitive-data-policy-ref) block (within [`sensitive_data_policy`](#sensitive-data-policy)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Single LB App

A [`single_lb_app`](#single-lb-app) block supports the following:

&#x2022; [`disable_discovery`](#disable-discovery) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`disable_malicious_user_detection`](#disable-malicious-user-detection) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`enable_discovery`](#enable-discovery) - Optional Block<br>API Discovery Setting. Specifies the settings used for API discovery<br>See [Enable Discovery](#single-lb-app-enable-discovery) below.

&#x2022; [`enable_malicious_user_detection`](#enable-malicious-user-detection) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### Single LB App Enable Discovery

An [`enable_discovery`](#single-lb-app-enable-discovery) block (within [`single_lb_app`](#single-lb-app)) supports the following:

&#x2022; [`api_crawler`](#api-crawler) - Optional Block<br>API Crawling. API Crawler message<br>See [API Crawler](#single-lb-app-enable-discovery-api-crawler) below.

&#x2022; [`api_discovery_from_code_scan`](#api-discovery-from-code-scan) - Optional Block<br>Select Code Base and Repositories. x-required<br>See [API Discovery From Code Scan](#single-lb-app-enable-discovery-api-discovery-from-code-scan) below.

&#x2022; [`custom_api_auth_discovery`](#custom-api-auth-discovery) - Optional Block<br>API Discovery Advanced Settings. API Discovery Advanced settings<br>See [Custom API Auth Discovery](#single-lb-app-enable-discovery-custom-api-auth-discovery) below.

&#x2022; [`default_api_auth_discovery`](#default-api-auth-discovery) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`disable_learn_from_redirect_traffic`](#disable-learn-from-redirect-traffic) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`discovered_api_settings`](#discovered-api-settings) - Optional Block<br>Discovered API Settings. x-example: '2' Configure Discovered API Settings<br>See [Discovered API Settings](#single-lb-app-enable-discovery-discovered-api-settings) below.

&#x2022; [`enable_learn_from_redirect_traffic`](#enable-learn-from-redirect-traffic) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### Single LB App Enable Discovery API Crawler

An [`api_crawler`](#single-lb-app-enable-discovery-api-crawler) block (within [`single_lb_app.enable_discovery`](#single-lb-app-enable-discovery)) supports the following:

&#x2022; [`api_crawler_config`](#api-crawler-config) - Optional Block<br>Crawler Configure<br>See [API Crawler Config](#single-lb-app-enable-discovery-api-crawler-api-crawler-config) below.

&#x2022; [`disable_api_crawler`](#disable-api-crawler) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### Single LB App Enable Discovery API Crawler API Crawler Config

An [`api_crawler_config`](#single-lb-app-enable-discovery-api-crawler-api-crawler-config) block (within [`single_lb_app.enable_discovery.api_crawler`](#single-lb-app-enable-discovery-api-crawler)) supports the following:

&#x2022; [`domains`](#domains) - Optional Block<br>Domains to Crawl. Enter domains and their credentials to allow authenticated API crawling. You can only include domains you own that are associated with this Load Balancer<br>See [Domains](#single-lb-app-enable-discovery-api-crawler-api-crawler-config-domains) below.

#### Single LB App Enable Discovery API Crawler API Crawler Config Domains

A [`domains`](#single-lb-app-enable-discovery-api-crawler-api-crawler-config-domains) block (within [`single_lb_app.enable_discovery.api_crawler.api_crawler_config`](#single-lb-app-enable-discovery-api-crawler-api-crawler-config)) supports the following:

&#x2022; [`domain`](#domain) - Optional String<br>Domain. Select the domain to execute API Crawling with given credentials

&#x2022; [`simple_login`](#simple-login) - Optional Block<br>Simple Login<br>See [Simple Login](#single-lb-app-enable-discovery-api-crawler-api-crawler-config-domains-simple-login) below.

#### Single LB App Enable Discovery API Crawler API Crawler Config Domains Simple Login

A [`simple_login`](#single-lb-app-enable-discovery-api-crawler-api-crawler-config-domains-simple-login) block (within [`single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains`](#single-lb-app-enable-discovery-api-crawler-api-crawler-config-domains)) supports the following:

&#x2022; [`password`](#password) - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Password](#single-lb-app-enable-discovery-api-crawler-api-crawler-config-domains-simple-login-password) below.

&#x2022; [`user`](#user) - Optional String<br>User. Enter the username to assign credentials for the selected domain to crawl

#### Single LB App Enable Discovery API Crawler API Crawler Config Domains Simple Login Password

A [`password`](#single-lb-app-enable-discovery-api-crawler-api-crawler-config-domains-simple-login-password) block (within [`single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login`](#single-lb-app-enable-discovery-api-crawler-api-crawler-config-domains-simple-login)) supports the following:

&#x2022; [`blindfold_secret_info`](#blindfold-secret-info) - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#single-lb-app-enable-discovery-api-crawler-api-crawler-config-domains-simple-login-password-blindfold-secret-info) below.

&#x2022; [`clear_secret_info`](#clear-secret-info) - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#single-lb-app-enable-discovery-api-crawler-api-crawler-config-domains-simple-login-password-clear-secret-info) below.

#### Single LB App Enable Discovery API Crawler API Crawler Config Domains Simple Login Password Blindfold Secret Info

A [`blindfold_secret_info`](#single-lb-app-enable-discovery-api-crawler-api-crawler-config-domains-simple-login-password-blindfold-secret-info) block (within [`single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password`](#single-lb-app-enable-discovery-api-crawler-api-crawler-config-domains-simple-login-password)) supports the following:

&#x2022; [`decryption_provider`](#decryption-provider) - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

&#x2022; [`location`](#location) - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

&#x2022; [`store_provider`](#store-provider) - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

#### Single LB App Enable Discovery API Crawler API Crawler Config Domains Simple Login Password Clear Secret Info

A [`clear_secret_info`](#single-lb-app-enable-discovery-api-crawler-api-crawler-config-domains-simple-login-password-clear-secret-info) block (within [`single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password`](#single-lb-app-enable-discovery-api-crawler-api-crawler-config-domains-simple-login-password)) supports the following:

&#x2022; [`provider_ref`](#provider-ref) - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

&#x2022; [`url`](#url) - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

#### Single LB App Enable Discovery API Discovery From Code Scan

An [`api_discovery_from_code_scan`](#single-lb-app-enable-discovery-api-discovery-from-code-scan) block (within [`single_lb_app.enable_discovery`](#single-lb-app-enable-discovery)) supports the following:

&#x2022; [`code_base_integrations`](#code-base-integrations) - Optional Block<br>Select Code Base Integrations<br>See [Code Base Integrations](#single-lb-app-enable-discovery-api-discovery-from-code-scan-code-base-integrations) below.

#### Single LB App Enable Discovery API Discovery From Code Scan Code Base Integrations

A [`code_base_integrations`](#single-lb-app-enable-discovery-api-discovery-from-code-scan-code-base-integrations) block (within [`single_lb_app.enable_discovery.api_discovery_from_code_scan`](#single-lb-app-enable-discovery-api-discovery-from-code-scan)) supports the following:

&#x2022; [`all_repos`](#all-repos) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`code_base_integration`](#code-base-integration) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Code Base Integration](#single-lb-app-enable-discovery-api-discovery-from-code-scan-code-base-integrations-code-base-integration) below.

&#x2022; [`selected_repos`](#selected-repos) - Optional Block<br>API Code Repositories. Select which API repositories represent the LB applications<br>See [Selected Repos](#single-lb-app-enable-discovery-api-discovery-from-code-scan-code-base-integrations-selected-repos) below.

#### Single LB App Enable Discovery API Discovery From Code Scan Code Base Integrations Code Base Integration

A [`code_base_integration`](#single-lb-app-enable-discovery-api-discovery-from-code-scan-code-base-integrations-code-base-integration) block (within [`single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations`](#single-lb-app-enable-discovery-api-discovery-from-code-scan-code-base-integrations)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Single LB App Enable Discovery API Discovery From Code Scan Code Base Integrations Selected Repos

A [`selected_repos`](#single-lb-app-enable-discovery-api-discovery-from-code-scan-code-base-integrations-selected-repos) block (within [`single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations`](#single-lb-app-enable-discovery-api-discovery-from-code-scan-code-base-integrations)) supports the following:

&#x2022; [`api_code_repo`](#api-code-repo) - Optional List<br>API Code Repository. Code repository which contain API endpoints

#### Single LB App Enable Discovery Custom API Auth Discovery

A [`custom_api_auth_discovery`](#single-lb-app-enable-discovery-custom-api-auth-discovery) block (within [`single_lb_app.enable_discovery`](#single-lb-app-enable-discovery)) supports the following:

&#x2022; [`api_discovery_ref`](#api-discovery-ref) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [API Discovery Ref](#single-lb-app-enable-discovery-custom-api-auth-discovery-api-discovery-ref) below.

#### Single LB App Enable Discovery Custom API Auth Discovery API Discovery Ref

An [`api_discovery_ref`](#single-lb-app-enable-discovery-custom-api-auth-discovery-api-discovery-ref) block (within [`single_lb_app.enable_discovery.custom_api_auth_discovery`](#single-lb-app-enable-discovery-custom-api-auth-discovery)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Single LB App Enable Discovery Discovered API Settings

A [`discovered_api_settings`](#single-lb-app-enable-discovery-discovered-api-settings) block (within [`single_lb_app.enable_discovery`](#single-lb-app-enable-discovery)) supports the following:

&#x2022; [`purge_duration_for_inactive_discovered_apis`](#purge-duration-for-inactive-discovered-apis) - Optional Number<br>Purge Duration for Inactive Discovered APIs from Traffic. Inactive discovered API will be deleted after configured duration

#### Slow DDOS Mitigation

A [`slow_ddos_mitigation`](#slow-ddos-mitigation) block supports the following:

&#x2022; [`disable_request_timeout`](#disable-request-timeout) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`request_headers_timeout`](#request-headers-timeout) - Optional Number  Defaults to `10000`<br>Request Headers Timeout. The amount of time the client has to send only the headers on the request stream before the stream is cancelled. The milliseconds. This setting provides protection against Slowloris attacks

&#x2022; [`request_timeout`](#request-timeout) - Optional Number<br>Custom Timeout

#### Timeouts

A [`timeouts`](#timeouts) block supports the following:

&#x2022; [`create`](#create) - Optional String<br>A string that can be [parsed as a duration](`HTTPS://pkg.go.dev/time#ParseDuration`) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours)

&#x2022; [`delete`](#delete) - Optional String<br>A string that can be [parsed as a duration](`HTTPS://pkg.go.dev/time#ParseDuration`) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are saved into state before the destroy operation occurs

&#x2022; [`read`](#read) - Optional String<br>A string that can be [parsed as a duration](`HTTPS://pkg.go.dev/time#ParseDuration`) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh is enabled

&#x2022; [`update`](#update) - Optional String<br>A string that can be [parsed as a duration](`HTTPS://pkg.go.dev/time#ParseDuration`) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours)

#### Trusted Clients

A [`trusted_clients`](#trusted-clients) block supports the following:

&#x2022; [`actions`](#actions) - Optional List  Defaults to `SKIP_PROCESSING_WAF`<br>Possible values are `SKIP_PROCESSING_WAF`, `SKIP_PROCESSING_BOT`, `SKIP_PROCESSING_MUM`, `SKIP_PROCESSING_IP_REPUTATION`, `SKIP_PROCESSING_API_PROTECTION`, `SKIP_PROCESSING_OAS_VALIDATION`, `SKIP_PROCESSING_DDOS_PROTECTION`, `SKIP_PROCESSING_THREAT_MESH`, `SKIP_PROCESSING_MALWARE_PROTECTION`<br>Actions. Actions that should be taken when client identifier matches the rule

&#x2022; [`as_number`](#as-number) - Optional Number<br>AS Number. RFC 6793 defined 4-byte AS number

&#x2022; [`bot_skip_processing`](#bot-skip-processing) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`expiration_timestamp`](#expiration-timestamp) - Optional String<br>Expiration Timestamp. The expiration_timestamp is the RFC 3339 format timestamp at which the containing rule is considered to be logically expired. The rule continues to exist in the configuration but is not applied anymore

&#x2022; [`http_header`](#http-header) - Optional Block<br>HTTP Header. Request header name and value pairs<br>See [HTTP Header](#trusted-clients-http-header) below.

&#x2022; [`ip_prefix`](#ip-prefix) - Optional String<br>IPv4 Prefix. IPv4 prefix string

&#x2022; [`ipv6_prefix`](#ipv6-prefix) - Optional String<br>IPv6 Prefix. IPv6 prefix string

&#x2022; [`metadata`](#metadata) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#trusted-clients-metadata) below.

&#x2022; [`skip_processing`](#skip-processing) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`user_identifier`](#user-identifier) - Optional String<br>User Identifier. Identify user based on user identifier. User identifier value needs to be copied from security event

&#x2022; [`waf_skip_processing`](#waf-skip-processing) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### Trusted Clients HTTP Header

A [`http_header`](#trusted-clients-http-header) block (within [`trusted_clients`](#trusted-clients)) supports the following:

&#x2022; [`headers`](#headers) - Optional Block<br>HTTP Headers. List of HTTP header name and value pairs<br>See [Headers](#trusted-clients-http-header-headers) below.

#### Trusted Clients HTTP Header Headers

A [`headers`](#trusted-clients-http-header-headers) block (within [`trusted_clients.http_header`](#trusted-clients-http-header)) supports the following:

&#x2022; [`exact`](#exact) - Optional String<br>Exact. Header value to match exactly

&#x2022; [`invert_match`](#invert-match) - Optional Bool<br>NOT of match. Invert the result of the match to detect missing header or non-matching value

&#x2022; [`name`](#name) - Optional String<br>Name. Name of the header

&#x2022; [`presence`](#presence) - Optional Bool<br>Presence. If true, check for presence of header

&#x2022; [`regex`](#regex) - Optional String<br>Regex. Regex match of the header value in re2 format

#### Trusted Clients Metadata

A [`metadata`](#trusted-clients-metadata) block (within [`trusted_clients`](#trusted-clients)) supports the following:

&#x2022; [`description`](#description) - Optional String<br>Description. Human readable description

&#x2022; [`name`](#name) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### User Identification

An [`user_identification`](#user-identification) block supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### WAF Exclusion

A [`waf_exclusion`](#waf-exclusion) block supports the following:

&#x2022; [`waf_exclusion_inline_rules`](#waf-exclusion-inline-rules) - Optional Block<br>WAF Exclusion Inline Rules. A list of WAF exclusion rules that will be applied inline<br>See [WAF Exclusion Inline Rules](#waf-exclusion-waf-exclusion-inline-rules) below.

&#x2022; [`waf_exclusion_policy`](#waf-exclusion-policy) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [WAF Exclusion Policy](#waf-exclusion-waf-exclusion-policy) below.

#### WAF Exclusion WAF Exclusion Inline Rules

A [`waf_exclusion_inline_rules`](#waf-exclusion-waf-exclusion-inline-rules) block (within [`waf_exclusion`](#waf-exclusion)) supports the following:

&#x2022; [`rules`](#rules) - Optional Block<br>WAF Exclusion Rules. An ordered list of WAF Exclusions specific to this Load Balancer<br>See [Rules](#waf-exclusion-waf-exclusion-inline-rules-rules) below.

#### WAF Exclusion WAF Exclusion Inline Rules Rules

A [`rules`](#waf-exclusion-waf-exclusion-inline-rules-rules) block (within [`waf_exclusion.waf_exclusion_inline_rules`](#waf-exclusion-waf-exclusion-inline-rules)) supports the following:

&#x2022; [`any_domain`](#any-domain) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`any_path`](#any-path) - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; [`app_firewall_detection_control`](#app-firewall-detection-control) - Optional Block<br>App Firewall Detection Control. Define the list of Signature IDs, Violations, Attack Types and Bot Names that should be excluded from triggering on the defined match criteria<br>See [App Firewall Detection Control](#waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control) below.

&#x2022; [`exact_value`](#exact-value) - Optional String<br>Exact Value. Exact domain name

&#x2022; [`expiration_timestamp`](#expiration-timestamp) - Optional String<br>Expiration Timestamp. The expiration_timestamp is the RFC 3339 format timestamp at which the containing rule is considered to be logically expired. The rule continues to exist in the configuration but is not applied anymore

&#x2022; [`metadata`](#metadata) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#waf-exclusion-waf-exclusion-inline-rules-rules-metadata) below.

&#x2022; [`methods`](#methods) - Optional List  Defaults to `ANY`<br>Possible values are `ANY`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`, `COPY`<br>Methods. methods to be matched

&#x2022; [`path_prefix`](#path-prefix) - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

&#x2022; [`path_regex`](#path-regex) - Optional String<br>Path Regex. Define the regex for the path. For example, the regex ^/.*$ will match on all paths

&#x2022; [`suffix_value`](#suffix-value) - Optional String<br>Suffix Value. Suffix of domain name e.g 'xyz.com' will match '*.xyz.com' and 'xyz.com'

&#x2022; [`waf_skip_processing`](#waf-skip-processing) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### WAF Exclusion WAF Exclusion Inline Rules Rules App Firewall Detection Control

An [`app_firewall_detection_control`](#waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control) block (within [`waf_exclusion.waf_exclusion_inline_rules.rules`](#waf-exclusion-waf-exclusion-inline-rules-rules)) supports the following:

&#x2022; [`exclude_attack_type_contexts`](#exclude-attack-type-contexts) - Optional Block<br>Attack Types. Attack Types to be excluded for the defined match criteria<br>See [Exclude Attack Type Contexts](#waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control-exclude-attack-type-contexts) below.

&#x2022; [`exclude_bot_name_contexts`](#exclude-bot-name-contexts) - Optional Block<br>Bot Names. Bot Names to be excluded for the defined match criteria<br>See [Exclude Bot Name Contexts](#waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control-exclude-bot-name-contexts) below.

&#x2022; [`exclude_signature_contexts`](#exclude-signature-contexts) - Optional Block<br>Signature IDs. Signature IDs to be excluded for the defined match criteria<br>See [Exclude Signature Contexts](#waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control-exclude-signature-contexts) below.

&#x2022; [`exclude_violation_contexts`](#exclude-violation-contexts) - Optional Block<br>Violations. Violations to be excluded for the defined match criteria<br>See [Exclude Violation Contexts](#waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control-exclude-violation-contexts) below.

#### WAF Exclusion WAF Exclusion Inline Rules Rules App Firewall Detection Control Exclude Attack Type Contexts

An [`exclude_attack_type_contexts`](#waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control-exclude-attack-type-contexts) block (within [`waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control`](#waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control)) supports the following:

&#x2022; [`context`](#context) - Optional String  Defaults to `CONTEXT_ANY`<br>Possible values are `CONTEXT_ANY`, `CONTEXT_BODY`, `CONTEXT_REQUEST`, `CONTEXT_RESPONSE`, `CONTEXT_PARAMETER`, `CONTEXT_HEADER`, `CONTEXT_COOKIE`, `CONTEXT_URL`, `CONTEXT_URI`<br>WAF Exclusion Context Options. The available contexts for Exclusion rules. - CONTEXT_ANY: CONTEXT_ANY Detection will be excluded for all contexts. - CONTEXT_BODY: CONTEXT_BODY Detection will be excluded for the request body. - CONTEXT_REQUEST: CONTEXT_REQUEST Detection will be excluded for the request. - CONTEXT_RESPONSE: CONTEXT_RESPONSE - CONTEXT_PARAMETER: CONTEXT_PARAMETER Detection will be excluded for the parameters. The parameter name is required in the Context name field. If the field is left empty, the detection will be excluded for all parameters. - CONTEXT_HEADER: CONTEXT_HEADER Detection will be excluded for the headers. The header name is required in the Context name field. If the field is left empty, the detection will be excluded for all headers. - CONTEXT_COOKIE: CONTEXT_COOKIE Detection will be excluded for the cookies. The cookie name is required in the Context name field. If the field is left empty, the detection will be excluded for all cookies. - CONTEXT_URL: CONTEXT_URL Detection will be excluded for the request URL. - CONTEXT_URI: CONTEXT_URI

&#x2022; [`context_name`](#context-name) - Optional String<br>Context Name. Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an wildcard asterisk (*)

&#x2022; [`exclude_attack_type`](#exclude-attack-type) - Optional String  Defaults to `ATTACK_TYPE_NONE`<br>Possible values are `ATTACK_TYPE_NONE`, `ATTACK_TYPE_NON_BROWSER_CLIENT`, `ATTACK_TYPE_OTHER_APPLICATION_ATTACKS`, `ATTACK_TYPE_TROJAN_BACKDOOR_SPYWARE`, `ATTACK_TYPE_DETECTION_EVASION`, `ATTACK_TYPE_VULNERABILITY_SCAN`, `ATTACK_TYPE_ABUSE_OF_FUNCTIONALITY`, `ATTACK_TYPE_AUTHENTICATION_AUTHORIZATION_ATTACKS`, `ATTACK_TYPE_BUFFER_OVERFLOW`, `ATTACK_TYPE_PREDICTABLE_RESOURCE_LOCATION`, `ATTACK_TYPE_INFORMATION_LEAKAGE`, `ATTACK_TYPE_DIRECTORY_INDEXING`, `ATTACK_TYPE_PATH_TRAVERSAL`, `ATTACK_TYPE_XPATH_INJECTION`, `ATTACK_TYPE_LDAP_INJECTION`, `ATTACK_TYPE_SERVER_SIDE_CODE_INJECTION`, `ATTACK_TYPE_COMMAND_EXECUTION`, `ATTACK_TYPE_SQL_INJECTION`, `ATTACK_TYPE_CROSS_SITE_SCRIPTING`, `ATTACK_TYPE_DENIAL_OF_SERVICE`, `ATTACK_TYPE_HTTP_PARSER_ATTACK`, `ATTACK_TYPE_SESSION_HIJACKING`, `ATTACK_TYPE_HTTP_RESPONSE_SPLITTING`, `ATTACK_TYPE_FORCEFUL_BROWSING`, `ATTACK_TYPE_REMOTE_FILE_INCLUDE`, `ATTACK_TYPE_MALICIOUS_FILE_UPLOAD`, `ATTACK_TYPE_GRAPHQL_PARSER_ATTACK`<br>Attack Types. List of all Attack Types ATTACK_TYPE_NONE ATTACK_TYPE_NON_BROWSER_CLIENT ATTACK_TYPE_OTHER_APPLICATION_ATTACKS ATTACK_TYPE_TROJAN_BACKDOOR_SPYWARE ATTACK_TYPE_DETECTION_EVASION ATTACK_TYPE_VULNERABILITY_SCAN ATTACK_TYPE_ABUSE_OF_FUNCTIONALITY ATTACK_TYPE_AUTHENTICATION_AUTHORIZATION_ATTACKS ATTACK_TYPE_BUFFER_OVERFLOW ATTACK_TYPE_PREDICTABLE_RESOURCE_LOCATION ATTACK_TYPE_INFORMATION_LEAKAGE ATTACK_TYPE_DIRECTORY_INDEXING ATTACK_TYPE_PATH_TRAVERSAL ATTACK_TYPE_XPATH_INJECTION ATTACK_TYPE_LDAP_INJECTION ATTACK_TYPE_SERVER_SIDE_CODE_INJECTION ATTACK_TYPE_COMMAND_EXECUTION ATTACK_TYPE_SQL_INJECTION ATTACK_TYPE_CROSS_SITE_SCRIPTING ATTACK_TYPE_DENIAL_OF_SERVICE ATTACK_TYPE_HTTP_PARSER_ATTACK ATTACK_TYPE_SESSION_HIJACKING ATTACK_TYPE_HTTP_RESPONSE_SPLITTING ATTACK_TYPE_FORCEFUL_BROWSING ATTACK_TYPE_REMOTE_FILE_INCLUDE ATTACK_TYPE_MALICIOUS_FILE_UPLOAD ATTACK_TYPE_GRAPHQL_PARSER_ATTACK

#### WAF Exclusion WAF Exclusion Inline Rules Rules App Firewall Detection Control Exclude Bot Name Contexts

An [`exclude_bot_name_contexts`](#waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control-exclude-bot-name-contexts) block (within [`waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control`](#waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control)) supports the following:

&#x2022; [`bot_name`](#bot-name) - Optional String<br>Bot Name

#### WAF Exclusion WAF Exclusion Inline Rules Rules App Firewall Detection Control Exclude Signature Contexts

An [`exclude_signature_contexts`](#waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control-exclude-signature-contexts) block (within [`waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control`](#waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control)) supports the following:

&#x2022; [`context`](#context) - Optional String  Defaults to `CONTEXT_ANY`<br>Possible values are `CONTEXT_ANY`, `CONTEXT_BODY`, `CONTEXT_REQUEST`, `CONTEXT_RESPONSE`, `CONTEXT_PARAMETER`, `CONTEXT_HEADER`, `CONTEXT_COOKIE`, `CONTEXT_URL`, `CONTEXT_URI`<br>WAF Exclusion Context Options. The available contexts for Exclusion rules. - CONTEXT_ANY: CONTEXT_ANY Detection will be excluded for all contexts. - CONTEXT_BODY: CONTEXT_BODY Detection will be excluded for the request body. - CONTEXT_REQUEST: CONTEXT_REQUEST Detection will be excluded for the request. - CONTEXT_RESPONSE: CONTEXT_RESPONSE - CONTEXT_PARAMETER: CONTEXT_PARAMETER Detection will be excluded for the parameters. The parameter name is required in the Context name field. If the field is left empty, the detection will be excluded for all parameters. - CONTEXT_HEADER: CONTEXT_HEADER Detection will be excluded for the headers. The header name is required in the Context name field. If the field is left empty, the detection will be excluded for all headers. - CONTEXT_COOKIE: CONTEXT_COOKIE Detection will be excluded for the cookies. The cookie name is required in the Context name field. If the field is left empty, the detection will be excluded for all cookies. - CONTEXT_URL: CONTEXT_URL Detection will be excluded for the request URL. - CONTEXT_URI: CONTEXT_URI

&#x2022; [`context_name`](#context-name) - Optional String<br>Context Name. Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an wildcard asterisk (*)

&#x2022; [`signature_id`](#signature-id) - Optional Number<br>SignatureID. The allowed values for signature id are 0 and in the range of 200000001-299999999. 0 implies that all signatures will be excluded for the specified context

#### WAF Exclusion WAF Exclusion Inline Rules Rules App Firewall Detection Control Exclude Violation Contexts

An [`exclude_violation_contexts`](#waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control-exclude-violation-contexts) block (within [`waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control`](#waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control)) supports the following:

&#x2022; [`context`](#context) - Optional String  Defaults to `CONTEXT_ANY`<br>Possible values are `CONTEXT_ANY`, `CONTEXT_BODY`, `CONTEXT_REQUEST`, `CONTEXT_RESPONSE`, `CONTEXT_PARAMETER`, `CONTEXT_HEADER`, `CONTEXT_COOKIE`, `CONTEXT_URL`, `CONTEXT_URI`<br>WAF Exclusion Context Options. The available contexts for Exclusion rules. - CONTEXT_ANY: CONTEXT_ANY Detection will be excluded for all contexts. - CONTEXT_BODY: CONTEXT_BODY Detection will be excluded for the request body. - CONTEXT_REQUEST: CONTEXT_REQUEST Detection will be excluded for the request. - CONTEXT_RESPONSE: CONTEXT_RESPONSE - CONTEXT_PARAMETER: CONTEXT_PARAMETER Detection will be excluded for the parameters. The parameter name is required in the Context name field. If the field is left empty, the detection will be excluded for all parameters. - CONTEXT_HEADER: CONTEXT_HEADER Detection will be excluded for the headers. The header name is required in the Context name field. If the field is left empty, the detection will be excluded for all headers. - CONTEXT_COOKIE: CONTEXT_COOKIE Detection will be excluded for the cookies. The cookie name is required in the Context name field. If the field is left empty, the detection will be excluded for all cookies. - CONTEXT_URL: CONTEXT_URL Detection will be excluded for the request URL. - CONTEXT_URI: CONTEXT_URI

&#x2022; [`context_name`](#context-name) - Optional String<br>Context Name. Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an wildcard asterisk (*)

&#x2022; [`exclude_violation`](#exclude-violation) - Optional String  Defaults to `VIOL_NONE`<br>Possible values are `VIOL_NONE`, `VIOL_FILETYPE`, `VIOL_METHOD`, `VIOL_MANDATORY_HEADER`, `VIOL_HTTP_RESPONSE_STATUS`, `VIOL_REQUEST_MAX_LENGTH`, `VIOL_FILE_UPLOAD`, `VIOL_FILE_UPLOAD_IN_BODY`, `VIOL_XML_MALFORMED`, `VIOL_JSON_MALFORMED`, `VIOL_ASM_COOKIE_MODIFIED`, `VIOL_HTTP_PROTOCOL_MULTIPLE_HOST_HEADERS`, `VIOL_HTTP_PROTOCOL_BAD_HOST_HEADER_VALUE`, `VIOL_HTTP_PROTOCOL_UNPARSABLE_REQUEST_CONTENT`, `VIOL_HTTP_PROTOCOL_NULL_IN_REQUEST`, `VIOL_HTTP_PROTOCOL_BAD_HTTP_VERSION`, `VIOL_HTTP_PROTOCOL_SEVERAL_CONTENT_LENGTH_HEADERS`, `VIOL_EVASION_DIRECTORY_TRAVERSALS`, `VIOL_MALFORMED_REQUEST`, `VIOL_EVASION_MULTIPLE_DECODING`, `VIOL_DATA_GUARD`, `VIOL_EVASION_APACHE_WHITESPACE`, `VIOL_COOKIE_MODIFIED`, `VIOL_EVASION_IIS_UNICODE_CODEPOINTS`, `VIOL_EVASION_IIS_BACKSLASHES`, `VIOL_EVASION_PERCENT_U_DECODING`, `VIOL_EVASION_BARE_BYTE_DECODING`, `VIOL_EVASION_BAD_UNESCAPE`, `VIOL_HTTP_PROTOCOL_BODY_IN_GET_OR_HEAD_REQUEST`, `VIOL_ENCODING`, `VIOL_COOKIE_MALFORMED`, `VIOL_GRAPHQL_FORMAT`, `VIOL_GRAPHQL_MALFORMED`, `VIOL_GRAPHQL_INTROSPECTION_QUERY`<br>App Firewall Violation Type. List of all supported Violation Types VIOL_NONE VIOL_FILETYPE VIOL_METHOD VIOL_MANDATORY_HEADER VIOL_HTTP_RESPONSE_STATUS VIOL_REQUEST_MAX_LENGTH VIOL_FILE_UPLOAD VIOL_FILE_UPLOAD_IN_BODY VIOL_XML_MALFORMED VIOL_JSON_MALFORMED VIOL_ASM_COOKIE_MODIFIED VIOL_HTTP_PROTOCOL_MULTIPLE_HOST_HEADERS VIOL_HTTP_PROTOCOL_BAD_HOST_HEADER_VALUE VIOL_HTTP_PROTOCOL_UNPARSABLE_REQUEST_CONTENT VIOL_HTTP_PROTOCOL_NULL_IN_REQUEST VIOL_HTTP_PROTOCOL_BAD_HTTP_VERSION VIOL_HTTP_PROTOCOL_CRLF_CHARACTERS_BEFORE_REQUEST_START VIOL_HTTP_PROTOCOL_NO_HOST_HEADER_IN_HTTP_1_1_REQUEST VIOL_HTTP_PROTOCOL_BAD_MULTIPART_PARAMETERS_PARSING VIOL_HTTP_PROTOCOL_SEVERAL_CONTENT_LENGTH_HEADERS VIOL_HTTP_PROTOCOL_CONTENT_LENGTH_SHOULD_BE_A_POSITIVE_NUMBER VIOL_EVASION_DIRECTORY_TRAVERSALS VIOL_MALFORMED_REQUEST VIOL_EVASION_MULTIPLE_DECODING VIOL_DATA_GUARD VIOL_EVASION_APACHE_WHITESPACE VIOL_COOKIE_MODIFIED VIOL_EVASION_IIS_UNICODE_CODEPOINTS VIOL_EVASION_IIS_BACKSLASHES VIOL_EVASION_PERCENT_U_DECODING VIOL_EVASION_BARE_BYTE_DECODING VIOL_EVASION_BAD_UNESCAPE VIOL_HTTP_PROTOCOL_BAD_MULTIPART_FORMDATA_REQUEST_PARSING VIOL_HTTP_PROTOCOL_BODY_IN_GET_OR_HEAD_REQUEST VIOL_HTTP_PROTOCOL_HIGH_ASCII_CHARACTERS_IN_HEADERS VIOL_ENCODING VIOL_COOKIE_MALFORMED VIOL_GRAPHQL_FORMAT VIOL_GRAPHQL_MALFORMED VIOL_GRAPHQL_INTROSPECTION_QUERY

#### WAF Exclusion WAF Exclusion Inline Rules Rules Metadata

A [`metadata`](#waf-exclusion-waf-exclusion-inline-rules-rules-metadata) block (within [`waf_exclusion.waf_exclusion_inline_rules.rules`](#waf-exclusion-waf-exclusion-inline-rules-rules)) supports the following:

&#x2022; [`description`](#description) - Optional String<br>Description. Human readable description

&#x2022; [`name`](#name) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### WAF Exclusion WAF Exclusion Policy

A [`waf_exclusion_policy`](#waf-exclusion-waf-exclusion-policy) block (within [`waf_exclusion`](#waf-exclusion)) supports the following:

&#x2022; [`name`](#name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; [`namespace`](#namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; [`tenant`](#tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

## Import

Import is supported using the following syntax:

```shell
# Import using namespace/name format
terraform import f5xc_http_loadbalancer.example system/example
```
