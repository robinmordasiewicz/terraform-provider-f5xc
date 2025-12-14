---
page_title: "f5xc_http_loadbalancer Resource - terraform-provider-f5xc"
subcategory: "Load Balancing"
description: |-
  [Category: Load Balancing] [Namespace: required] [DependsOn: namespace, origin_pool] Manages a HTTP Load Balancer resource in F5 Distributed Cloud for load balancing HTTP/HTTPS traffic with advanced routing and security.
---

# f5xc_http_loadbalancer (Resource)

[Category: Load Balancing] [Namespace: required] [DependsOn: namespace, origin_pool] Manages a HTTP Load Balancer resource in F5 Distributed Cloud for load balancing HTTP/HTTPS traffic with advanced routing and security.

~> **Note** Please refer to [HTTP Loadbalancer API docs](https://docs.cloud.f5.com/docs-v2/api/views-http-loadbalancer) to learn more.

## Example Usage

```terraform
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

<a id="name"></a>&#x2022; [`name`](#name) - Required String<br>Name of the HTTP Load Balancer. Must be unique within the namespace

<a id="namespace"></a>&#x2022; [`namespace`](#namespace) - Required String<br>Namespace where the HTTP Load Balancer will be created

<a id="annotations"></a>&#x2022; [`annotations`](#annotations) - Optional Map<br>Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata

<a id="description"></a>&#x2022; [`description`](#description) - Optional String<br>Human readable description for the object

<a id="disable"></a>&#x2022; [`disable`](#disable) - Optional Bool<br>A value of true will administratively disable the object

<a id="labels"></a>&#x2022; [`labels`](#labels) - Optional Map<br>Labels is a user defined key value map that can be attached to resources for organization and filtering

### Spec Argument Reference

-> **One of the following:**
&#x2022; <a id="active-service-policies"></a>[`active_service_policies`](#active-service-policies) - Optional Block<br>Service Policy List. List of service policies<br>See [Active Service Policies](#active-service-policies) below for details.
<br><br>&#x2022; <a id="no-service-policies"></a>[`no_service_policies`](#no-service-policies) - Optional Block<br>Enable this option

<a id="add-location"></a>&#x2022; [`add_location`](#add-location) - Optional Bool<br>Add Location. Appends header x-volterra-location = `<RE-site-name>` in responses. This configuration is ignored on CE sites

-> **One of the following:**
&#x2022; <a id="advertise-custom"></a>[`advertise_custom`](#advertise-custom) - Optional Block<br>Advertise Custom. This defines a way to advertise a VIP on specific sites<br>See [Advertise Custom](#advertise-custom) below for details.
<br><br>&#x2022; <a id="advertise-on-public"></a>[`advertise_on_public`](#advertise-on-public) - Optional Block<br>Advertise Public. This defines a way to advertise a load balancer on public. If optional public_ip is provided, it will only be advertised on RE sites where that public_ip is available<br>See [Advertise On Public](#advertise-on-public) below for details.
<br><br>&#x2022; <a id="advertise-on-public-default-vip"></a>[`advertise_on_public_default_vip`](#advertise-on-public-default-vip) - Optional Block<br>Enable this option

<a id="api-protection-rules"></a>&#x2022; [`api_protection_rules`](#api-protection-rules) - Optional Block<br>API Protection Rules. API Protection Rules<br>See [API Protection Rules](#api-protection-rules) below for details.

-> **One of the following:**
&#x2022; <a id="api-rate-limit"></a>[`api_rate_limit`](#api-rate-limit) - Optional Block<br>APIRateLimit
<br><br>&#x2022; <a id="disable-rate-limit"></a>[`disable_rate_limit`](#disable-rate-limit) - Optional Block<br>Enable this option

-> **One of the following:**
&#x2022; <a id="api-specification"></a>[`api_specification`](#api-specification) - Optional Block<br>API Specification and Validation. Settings for API specification (API definition, OpenAPI validation, etc.)

-> **One of the following:**
&#x2022; <a id="api-testing"></a>[`api_testing`](#api-testing) - Optional Block<br>API Testing

-> **One of the following:**
&#x2022; <a id="app-firewall"></a>[`app_firewall`](#app-firewall) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name

<a id="blocked-clients"></a>&#x2022; [`blocked_clients`](#blocked-clients) - Optional Block<br>Client Blocking Rules. Define rules to block IP Prefixes or AS numbers

-> **One of the following:**
&#x2022; <a id="bot-defense"></a>[`bot_defense`](#bot-defense) - Optional Block<br>Bot Defense. This defines various configuration options for Bot Defense Policy
<br><br>&#x2022; <a id="bot-defense-advanced"></a>[`bot_defense_advanced`](#bot-defense-advanced) - Optional Block<br>Bot Defense Advanced. Bot Defense Advanced

-> **One of the following:**
&#x2022; <a id="caching-policy"></a>[`caching_policy`](#caching-policy) - Optional Block<br>Caching Policies.Caching Policies for the CDN

-> **One of the following:**
&#x2022; <a id="captcha-challenge"></a>[`captcha_challenge`](#captcha-challenge) - Optional Block<br>Captcha Challenge Parameters. Enables loadbalancer to perform captcha challenge Captcha challenge will be based on Google Recaptcha. With this feature enabled, only clients that pass the captcha challenge will be allowed to complete the HTTP request. When loadbalancer is configured to do Captcha Challenge, it will redirect the browser to an HTML page on every new HTTP request. This HTML page will have captcha challenge embedded in it. Client will be allowed to make the request only if the captcha challenge is successful. Loadbalancer will tag response header with a cookie to avoid Captcha challenge for subsequent requests. CAPTCHA is mainly used as a security check to ensure only human users can pass through. Generally, computers or bots are not capable of solving a captcha. You can enable either Javascript challenge or Captcha challenge on a virtual host
<br><br>&#x2022; <a id="enable-challenge"></a>[`enable_challenge`](#enable-challenge) - Optional Block<br>Enable Malicious User Challenge. Configure auto mitigation i.e risk based challenges for malicious users
<br><br>&#x2022; <a id="js-challenge"></a>[`js_challenge`](#js-challenge) - Optional Block<br>Javascript Challenge Parameters. Enables loadbalancer to perform client browser compatibility test by redirecting to a page with Javascript. With this feature enabled, only clients that are capable of executing Javascript(mostly browsers) will be allowed to complete the HTTP request. When loadbalancer is configured to do Javascript Challenge, it will redirect the browser to an HTML page on every new HTTP request. This HTML page will have Javascript embedded in it. Loadbalancer chooses a set of random numbers for every new client and sends these numbers along with an encrypted answer with the request such that it embed these numbers as input in the Javascript. Javascript will run on the requestor browser and perform a complex Math operation. Script will submit the answer to loadbalancer. Loadbalancer will validate the answer by comparing the calculated answer with the decrypted answer (which was encrypted when it was sent back as reply) and allow the request to the upstream server only if the answer is correct. Loadbalancer will tag response header with a cookie to avoid Javascript challenge for subsequent requests. Javascript challenge serves following purposes \* Validate that the request is coming via a browser that is capable for running Javascript \* Force the browser to run a complex operation, f(X), that requires it to spend a large number of CPU cycles. This is to slow down a potential DOS attacker by making it difficult to launch a large request flood without having to spend even larger CPU cost at their end. You can enable either Javascript challenge or Captcha challenge on a virtual host
<br><br>&#x2022; <a id="no-challenge"></a>[`no_challenge`](#no-challenge) - Optional Block<br>Enable this option

-> **One of the following:**
&#x2022; <a id="client-side-defense"></a>[`client_side_defense`](#client-side-defense) - Optional Block<br>Client-Side Defense. This defines various configuration options for Client-Side Defense Policy

-> **One of the following:**
&#x2022; <a id="cookie-stickiness"></a>[`cookie_stickiness`](#cookie-stickiness) - Optional Block<br>Hashing using Cookie. Two types of cookie affinity: 1. Passive. Takes a cookie that's present in the cookies header and hashes on its value. 2. Generated. Generates and sets a cookie with an expiration (TTL) on the first request from the client in its response to the client, based on the endpoint the request gets sent to. The client then presents this on the next and all subsequent requests. The hash of this is sufficient to ensure these requests get sent to the same endpoint. The cookie is generated by hashing the source and destination ports and addresses so that multiple independent HTTP2 streams on the same connection will independently receive the same cookie, even if they arrive simultaneously
<br><br>&#x2022; <a id="least-active"></a>[`least_active`](#least-active) - Optional Block<br>Enable this option
<br><br>&#x2022; <a id="random"></a>[`random`](#random) - Optional Block<br>Enable this option
<br><br>&#x2022; <a id="ring-hash"></a>[`ring_hash`](#ring-hash) - Optional Block<br>Hash Policy List. List of hash policy rules
<br><br>&#x2022; <a id="round-robin"></a>[`round_robin`](#round-robin) - Optional Block<br>Enable this option

<a id="cors-policy"></a>&#x2022; [`cors_policy`](#cors-policy) - Optional Block<br>CORS Policy. Cross-Origin Resource Sharing requests configuration specified at Virtual-host or Route level. Route level configuration takes precedence. An example of an Cross origin HTTP request GET /resources/public-data/ HTTP/1.1 Host: bar.other User-Agent: Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.5; en-US; rv:1.9.1b3pre) Gecko/20081130 Minefield/3.1b3pre Accept: text/HTML,application/xhtml+XML,application/XML;q=0.9,*/*;q=0.8 Accept-Language: en-us,en;q=0.5 Accept-Encoding: gzip,deflate Accept-Charset: ISO-8859-1,utf-8;q=0.7,*;q=0.7 Connection: keep-alive Referrer: `HTTP://foo.example/examples/access-control/simpleXSInvocation.HTML` Origin: `HTTP://foo.example` HTTP/1.1 200 OK Date: Mon, 01 Dec 2008 00:23:53 GMT Server: Apache/2.0.61 Access-Control-Allow-Origin: \* Keep-Alive: timeout=2, max=100 Connection: Keep-Alive Transfer-Encoding: chunked Content-Type: application/XML An example for cross origin HTTP OPTIONS request with Access-Control-Request-* header OPTIONS /resources/post-here/ HTTP/1.1 Host: bar.other User-Agent: Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.5; en-US; rv:1.9.1b3pre) Gecko/20081130 Minefield/3.1b3pre Accept: text/HTML,application/xhtml+XML,application/XML;q=0.9,*/*;q=0.8 Accept-Language: en-us,en;q=0.5 Accept-Encoding: gzip,deflate Accept-Charset: ISO-8859-1,utf-8;q=0.7,*;q=0.7 Connection: keep-alive Origin: `HTTP://foo.example` Access-Control-Request-Method: POST Access-Control-Request-Headers: X-PINGOTHER, Content-Type HTTP/1.1 204 No Content Date: Mon, 01 Dec 2008 01:15:39 GMT Server: Apache/2.0.61 (Unix) Access-Control-Allow-Origin: `HTTP://foo.example` Access-Control-Allow-Methods: POST, GET, OPTIONS Access-Control-Allow-Headers: X-PINGOTHER, Content-Type Access-Control-Max-Age: 86400 Vary: Accept-Encoding, Origin Keep-Alive: timeout=2, max=100 Connection: Keep-Alive

<a id="csrf-policy"></a>&#x2022; [`csrf_policy`](#csrf-policy) - Optional Block<br>CSRF Policy. To mitigate CSRF attack , the policy checks where a request is coming from to determine if the request's origin is the same as its detination.The policy relies on two pieces of information used in determining if a request originated from the same host. 1. The origin that caused the user agent to issue the request (source origin). 2. The origin that the request is going to (target origin). When the policy evaluating a request, it ensures both pieces of information are present and compare their values. If the source origin is missing or origins do not match the request is rejected. The exception to this being if the source-origin has been added to they policy as valid. Because CSRF attacks specifically target state-changing requests, the policy only acts on the HTTP requests that have state-changing method (PUT,POST, etc.)

<a id="data-guard-rules"></a>&#x2022; [`data_guard_rules`](#data-guard-rules) - Optional Block<br>Data Guard Rules. Data Guard prevents responses from exposing sensitive information by masking the data. The system masks credit card numbers and social security numbers leaked from the application from within the HTTP response with a string of asterisks (*). Note: App Firewall should be enabled, to use Data Guard feature

<a id="ddos-mitigation-rules"></a>&#x2022; [`ddos_mitigation_rules`](#ddos-mitigation-rules) - Optional Block<br>DDOS Mitigation Rules. Define manual mitigation rules to block L7 DDOS attacks

-> **One of the following:**
&#x2022; <a id="default-pool"></a>[`default_pool`](#default-pool) - Optional Block<br>Global Specification. Shape of the origin pool specification

<a id="default-pool-list"></a>&#x2022; [`default_pool_list`](#default-pool-list) - Optional Block<br>Origin Pool List Type. List of Origin Pools

<a id="default-route-pools"></a>&#x2022; [`default_route_pools`](#default-route-pools) - Optional Block<br>Origin Pools. Origin Pools used when no route is specified (default route)

-> **One of the following:**
&#x2022; <a id="default-sensitive-data-policy"></a>[`default_sensitive_data_policy`](#default-sensitive-data-policy) - Optional Block<br>Enable this option

<a id="disable-api-definition"></a>&#x2022; [`disable_api_definition`](#disable-api-definition) - Optional Block<br>Enable this option

-> **One of the following:**
&#x2022; <a id="disable-api-discovery"></a>[`disable_api_discovery`](#disable-api-discovery) - Optional Block<br>Enable this option

<a id="disable-api-testing"></a>&#x2022; [`disable_api_testing`](#disable-api-testing) - Optional Block<br>Enable this option

<a id="disable-bot-defense"></a>&#x2022; [`disable_bot_defense`](#disable-bot-defense) - Optional Block<br>Enable this option

<a id="disable-caching"></a>&#x2022; [`disable_caching`](#disable-caching) - Optional Block<br>Enable this option

<a id="disable-client-side-defense"></a>&#x2022; [`disable_client_side_defense`](#disable-client-side-defense) - Optional Block<br>Enable this option

-> **One of the following:**
&#x2022; <a id="disable-ip-reputation"></a>[`disable_ip_reputation`](#disable-ip-reputation) - Optional Block<br>Enable this option

-> **One of the following:**
&#x2022; <a id="disable-malicious-user-detection"></a>[`disable_malicious_user_detection`](#disable-malicious-user-detection) - Optional Block<br>Enable this option

-> **One of the following:**
&#x2022; <a id="disable-malware-protection"></a>[`disable_malware_protection`](#disable-malware-protection) - Optional Block<br>Enable this option

-> **One of the following:**
&#x2022; <a id="disable-threat-mesh"></a>[`disable_threat_mesh`](#disable-threat-mesh) - Optional Block<br>Enable this option

-> **One of the following:**
&#x2022; <a id="disable-trust-client-ip-headers"></a>[`disable_trust_client_ip_headers`](#disable-trust-client-ip-headers) - Optional Block<br>Enable this option

<a id="disable-waf"></a>&#x2022; [`disable_waf`](#disable-waf) - Optional Block<br>Enable this option

<a id="do-not-advertise"></a>&#x2022; [`do_not_advertise`](#do-not-advertise) - Optional Block<br>Enable this option

<a id="domains"></a>&#x2022; [`domains`](#domains) - Optional List<br>Domains. A list of Domains (host/authority header) that will be matched to load balancer. Supported Domains and search order: 1. Exact Domain names: `www.foo.com.` 2. Domains starting with a Wildcard: \*.foo.com. Not supported Domains: - Just a Wildcard: \* - A Wildcard and TLD with no root Domain: \*.com. - A Wildcard not matching a whole DNS label. e.g. \*.foo.com and \*.bar.foo.com are valid Wildcards however \*bar.foo.com, \*-bar.foo.com, and bar*.foo.com are all invalid. Additional notes: A Wildcard will not match empty string. e.g. \*.foo.com will match bar.foo.com and baz-bar.foo.com but not .foo.com. The longest Wildcards match first. Only a single virtual host in the entire route configuration can match on \*. Also a Domain must be unique across all virtual hosts within an advertise policy. Domains are also used for SNI matching if the Loadbalancer type is HTTPS. Domains also indicate the list of names for which DNS resolution will be automatically resolved to IP addresses by the system

<a id="enable-api-discovery"></a>&#x2022; [`enable_api_discovery`](#enable-api-discovery) - Optional Block<br>API Discovery Setting. Specifies the settings used for API discovery

<a id="enable-ip-reputation"></a>&#x2022; [`enable_ip_reputation`](#enable-ip-reputation) - Optional Block<br>IP Threat Category List. List of IP threat categories

<a id="enable-malicious-user-detection"></a>&#x2022; [`enable_malicious_user_detection`](#enable-malicious-user-detection) - Optional Block<br>Enable this option

<a id="enable-threat-mesh"></a>&#x2022; [`enable_threat_mesh`](#enable-threat-mesh) - Optional Block<br>Enable this option

<a id="enable-trust-client-ip-headers"></a>&#x2022; [`enable_trust_client_ip_headers`](#enable-trust-client-ip-headers) - Optional Block<br>Trust Client IP Headers List. List of Client IP Headers

<a id="graphql-rules"></a>&#x2022; [`graphql_rules`](#graphql-rules) - Optional Block<br>GraphQL Inspection. GraphQL is a query language and server-side runtime for APIs which provides a complete and understandable description of the data in API. GraphQL gives clients the power to ask for exactly what they need, makes it easier to evolve APIs over time, and enables powerful developer tools. Policy configuration to analyze GraphQL queries and prevent GraphQL tailored attacks

-> **One of the following:**
&#x2022; <a id="http"></a>[`http`](#http) - Optional Block<br>HTTP Choice. Choice for selecting HTTP proxy
<br><br>&#x2022; <a id="https"></a>[`https`](#https) - Optional Block<br>BYOC HTTPS Choice. Choice for selecting HTTP proxy with bring your own certificates

<a id="https-auto-cert"></a>&#x2022; [`https_auto_cert`](#https-auto-cert) - Optional Block<br>HTTPS with Auto Certs Choice. Choice for selecting HTTP proxy with bring your own certificates

<a id="jwt-validation"></a>&#x2022; [`jwt_validation`](#jwt-validation) - Optional Block<br>JWT Validation. JWT Validation stops JWT replay attacks and JWT tampering by cryptographically verifying incoming JWTs before they are passed to your API origin. JWT Validation will also stop requests with expired tokens or tokens that are not yet valid

-> **One of the following:**
&#x2022; <a id="l7-ddos-action-block"></a>[`l7_ddos_action_block`](#l7-ddos-action-block) - Optional Block<br>Enable this option
<br><br>&#x2022; <a id="l7-ddos-action-default"></a>[`l7_ddos_action_default`](#l7-ddos-action-default) - Optional Block<br>Enable this option

<a id="l7-ddos-action-js-challenge"></a>&#x2022; [`l7_ddos_action_js_challenge`](#l7-ddos-action-js-challenge) - Optional Block<br>Javascript Challenge Parameters. Enables loadbalancer to perform client browser compatibility test by redirecting to a page with Javascript. With this feature enabled, only clients that are capable of executing Javascript(mostly browsers) will be allowed to complete the HTTP request. When loadbalancer is configured to do Javascript Challenge, it will redirect the browser to an HTML page on every new HTTP request. This HTML page will have Javascript embedded in it. Loadbalancer chooses a set of random numbers for every new client and sends these numbers along with an encrypted answer with the request such that it embed these numbers as input in the Javascript. Javascript will run on the requestor browser and perform a complex Math operation. Script will submit the answer to loadbalancer. Loadbalancer will validate the answer by comparing the calculated answer with the decrypted answer (which was encrypted when it was sent back as reply) and allow the request to the upstream server only if the answer is correct. Loadbalancer will tag response header with a cookie to avoid Javascript challenge for subsequent requests. Javascript challenge serves following purposes \* Validate that the request is coming via a browser that is capable for running Javascript \* Force the browser to run a complex operation, f(X), that requires it to spend a large number of CPU cycles. This is to slow down a potential DOS attacker by making it difficult to launch a large request flood without having to spend even larger CPU cost at their end. You can enable either Javascript challenge or Captcha challenge on a virtual host

<a id="l7-ddos-protection"></a>&#x2022; [`l7_ddos_protection`](#l7-ddos-protection) - Optional Block<br>L7 DDOS Protection Settings. L7 DDOS protection is critical for safeguarding web applications, APIs, and services that are exposed to the internet from sophisticated, volumetric, application-level threats. Configure actions, thresholds and policies to apply during L7 DDOS attack

<a id="malware-protection-settings"></a>&#x2022; [`malware_protection_settings`](#malware-protection-settings) - Optional Block<br>Malware Protection Policy. Malware Protection protects Web Apps and APIs, from malicious file uploads by scanning files in real-time

<a id="more-option"></a>&#x2022; [`more_option`](#more-option) - Optional Block<br>Advanced Options. This defines various options to define a route

-> **One of the following:**
&#x2022; <a id="multi-lb-app"></a>[`multi_lb_app`](#multi-lb-app) - Optional Block<br>Enable this option
<br><br>&#x2022; <a id="single-lb-app"></a>[`single_lb_app`](#single-lb-app) - Optional Block<br>Single Load Balancer App Setting. Specific settings for Machine learning analysis on this HTTP LB, independently from other LBs

<a id="origin-server-subset-rule-list"></a>&#x2022; [`origin_server_subset_rule_list`](#origin-server-subset-rule-list) - Optional Block<br>Origin Server Subset Rule List Type. List of Origin Pools

<a id="policy-based-challenge"></a>&#x2022; [`policy_based_challenge`](#policy-based-challenge) - Optional Block<br>Policy Based Challenge. Specifies the settings for policy rule based challenge

<a id="protected-cookies"></a>&#x2022; [`protected_cookies`](#protected-cookies) - Optional Block<br>Cookie Protection. Allows setting attributes (SameSite, Secure, and HttpOnly) on cookies in responses. Cookie Tampering Protection prevents attackers from modifying the value of session cookies. For Cookie Tampering Protection, enabling a web app firewall (WAF) is a prerequisite. The configured mode of WAF (monitoring or blocking) will be enforced on the request when cookie tampering is identified. Note: We recommend enabling Secure and HttpOnly attributes along with cookie tampering protection

<a id="rate-limit"></a>&#x2022; [`rate_limit`](#rate-limit) - Optional Block<br>RateLimitConfigType

<a id="routes"></a>&#x2022; [`routes`](#routes) - Optional Block<br>Routes. Routes allow users to define match condition on a path and/or HTTP method to either forward matching traffic to origin pool or redirect matching traffic to a different URL or respond directly to matching traffic

<a id="sensitive-data-disclosure-rules"></a>&#x2022; [`sensitive_data_disclosure_rules`](#sensitive-data-disclosure-rules) - Optional Block<br>Sensitive Data Exposure Rules. Sensitive Data Exposure Rules allows specifying rules to mask sensitive data fields in API responses

<a id="sensitive-data-policy"></a>&#x2022; [`sensitive_data_policy`](#sensitive-data-policy) - Optional Block<br>Sensitive Data Discovery. Settings for data type policy

<a id="service-policies-from-namespace"></a>&#x2022; [`service_policies_from_namespace`](#service-policies-from-namespace) - Optional Block<br>Enable this option

-> **One of the following:**
&#x2022; <a id="slow-ddos-mitigation"></a>[`slow_ddos_mitigation`](#slow-ddos-mitigation) - Optional Block<br>Slow DDOS Mitigation. 'Slow and low' attacks tie up server resources, leaving none available for servicing requests from actual users

<a id="source-ip-stickiness"></a>&#x2022; [`source_ip_stickiness`](#source-ip-stickiness) - Optional Block<br>Enable this option

<a id="system-default-timeouts"></a>&#x2022; [`system_default_timeouts`](#system-default-timeouts) - Optional Block<br>Enable this option

<a id="timeouts"></a>&#x2022; [`timeouts`](#timeouts) - Optional Block

<a id="trusted-clients"></a>&#x2022; [`trusted_clients`](#trusted-clients) - Optional Block<br>Trusted Client Rules. Define rules to skip processing of one or more features such as WAF, Bot Defense etc. for clients

-> **One of the following:**
&#x2022; <a id="user-id-client-ip"></a>[`user_id_client_ip`](#user-id-client-ip) - Optional Block<br>Enable this option
<br><br>&#x2022; <a id="user-identification"></a>[`user_identification`](#user-identification) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name

<a id="waf-exclusion"></a>&#x2022; [`waf_exclusion`](#waf-exclusion) - Optional Block<br>WAF Exclusion

### Attributes Reference

In addition to all arguments above, the following attributes are exported:

<a id="id"></a>&#x2022; [`id`](#id) - Optional String<br>Unique identifier for the resource

---

#### Active Service Policies

An [`active_service_policies`](#active-service-policies) block supports the following:

<a id="active-service-policies-policies"></a>&#x2022; [`policies`](#active-service-policies-policies) - Optional Block<br>Policies. Service Policies is a sequential engine where policies (and rules within the policy) are evaluated one after the other. It's important to define the correct order (policies evaluated from top to bottom in the list) for service policies, to get the intended result. For each request, its characteristics are evaluated based on the match criteria in each service policy starting at the top. If there is a match in the current policy, then the policy takes effect, and no more policies are evaluated. Otherwise, the next policy is evaluated. If all policies are evaluated and none match, then the request will be denied by default<br>See [Policies](#active-service-policies-policies) below.

#### Active Service Policies Policies

A [`policies`](#active-service-policies-policies) block (within [`active_service_policies`](#active-service-policies)) supports the following:

<a id="active-service-policies-policies-name"></a>&#x2022; [`name`](#active-service-policies-policies-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-df0e5f"></a>&#x2022; [`namespace`](#namespace-df0e5f) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="active-service-policies-policies-tenant"></a>&#x2022; [`tenant`](#active-service-policies-policies-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Advertise Custom

An [`advertise_custom`](#advertise-custom) block supports the following:

<a id="advertise-custom-advertise-where"></a>&#x2022; [`advertise_where`](#advertise-custom-advertise-where) - Optional Block<br>List of Sites to Advertise. Where should this load balancer be available<br>See [Advertise Where](#advertise-custom-advertise-where) below.

#### Advertise Custom Advertise Where

An [`advertise_where`](#advertise-custom-advertise-where) block (within [`advertise_custom`](#advertise-custom)) supports the following:

<a id="public-618a99"></a>&#x2022; [`advertise_on_public`](#public-618a99) - Optional Block<br>Advertise Public. This defines a way to advertise a load balancer on public. If optional public_ip is provided, it will only be advertised on RE sites where that public_ip is available<br>See [Advertise On Public](#public-618a99) below.

<a id="advertise-custom-advertise-where-port"></a>&#x2022; [`port`](#advertise-custom-advertise-where-port) - Optional Number<br>Listen Port. Port to Listen

<a id="ranges-7cbec3"></a>&#x2022; [`port_ranges`](#ranges-7cbec3) - Optional String<br>Listen Port Ranges. A string containing a comma separated list of port ranges. Each port range consists of a single port or two ports separated by '-'

<a id="advertise-custom-advertise-where-site"></a>&#x2022; [`site`](#advertise-custom-advertise-where-site) - Optional Block<br>Site. This defines a reference to a CE site along with network type and an optional IP address where a load balancer could be advertised<br>See [Site](#advertise-custom-advertise-where-site) below.

<a id="port-b19c4f"></a>&#x2022; [`use_default_port`](#port-b19c4f) - Optional Block<br>Enable this option

<a id="network-a20be3"></a>&#x2022; [`virtual_network`](#network-a20be3) - Optional Block<br>Virtual Network. Parameters to advertise on a given virtual network<br>See [Virtual Network](#network-a20be3) below.

<a id="site-5d39fd"></a>&#x2022; [`virtual_site`](#site-5d39fd) - Optional Block<br>Virtual Site. This defines a reference to a customer site virtual site along with network type where a load balancer could be advertised<br>See [Virtual Site](#site-5d39fd) below.

<a id="vip-870b0b"></a>&#x2022; [`virtual_site_with_vip`](#vip-870b0b) - Optional Block<br>Virtual Site with Specified VIP. This defines a reference to a customer site virtual site along with network type and IP where a load balancer could be advertised<br>See [Virtual Site With VIP](#vip-870b0b) below.

<a id="service-1fdc7a"></a>&#x2022; [`vk8s_service`](#service-1fdc7a) - Optional Block<br>vK8s Services on RE. This defines a reference to a RE site or virtual site where a load balancer could be advertised in the vK8s service network<br>See [Vk8s Service](#service-1fdc7a) below.

#### Advertise Custom Advertise Where Advertise On Public

An [`advertise_on_public`](#public-618a99) block (within [`advertise_custom.advertise_where`](#advertise-custom-advertise-where)) supports the following:

<a id="public-ip-d10b09"></a>&#x2022; [`public_ip`](#public-ip-d10b09) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Public IP](#public-ip-d10b09) below.

#### Advertise Custom Advertise Where Advertise On Public Public IP

A [`public_ip`](#public-ip-d10b09) block (within [`advertise_custom.advertise_where.advertise_on_public`](#public-618a99)) supports the following:

<a id="name-4126f8"></a>&#x2022; [`name`](#name-4126f8) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-edf1ff"></a>&#x2022; [`namespace`](#namespace-edf1ff) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-ac4633"></a>&#x2022; [`tenant`](#tenant-ac4633) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Advertise Custom Advertise Where Site

A [`site`](#advertise-custom-advertise-where-site) block (within [`advertise_custom.advertise_where`](#advertise-custom-advertise-where)) supports the following:

<a id="site-ip-78faa1"></a>&#x2022; [`ip`](#site-ip-78faa1) - Optional String<br>IP Address. Use given IP address as VIP on the site

<a id="network-5811a4"></a>&#x2022; [`network`](#network-5811a4) - Optional String  Defaults to `SITE_NETWORK_INSIDE_AND_OUTSIDE`<br>Possible values are `SITE_NETWORK_INSIDE_AND_OUTSIDE`, `SITE_NETWORK_INSIDE`, `SITE_NETWORK_OUTSIDE`, `SITE_NETWORK_SERVICE`, `SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP`, `SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP`, `SITE_NETWORK_IP_FABRIC`<br>[Enum: SITE_NETWORK_INSIDE_AND_OUTSIDE|SITE_NETWORK_INSIDE|SITE_NETWORK_OUTSIDE|SITE_NETWORK_SERVICE|SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP|SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP|SITE_NETWORK_IP_FABRIC] Site Network. This defines network types to be used on site All inside and outside networks. All inside and outside networks with internet VIP support. All inside networks. All outside networks. All outside networks with internet VIP support. vK8s service network. - SITE_NETWORK_IP_FABRIC: VER IP Fabric network for the site This Virtual network type is used for exposing virtual host on IP Fabric network on the VER site or for endpoint in IP Fabric network

<a id="site-7ecf1d"></a>&#x2022; [`site`](#site-7ecf1d) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Site](#site-7ecf1d) below.

#### Advertise Custom Advertise Where Site Site

A [`site`](#site-7ecf1d) block (within [`advertise_custom.advertise_where.site`](#advertise-custom-advertise-where-site)) supports the following:

<a id="name-201d26"></a>&#x2022; [`name`](#name-201d26) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-c3f40d"></a>&#x2022; [`namespace`](#namespace-c3f40d) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-8a632a"></a>&#x2022; [`tenant`](#tenant-8a632a) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Advertise Custom Advertise Where Virtual Network

A [`virtual_network`](#network-a20be3) block (within [`advertise_custom.advertise_where`](#advertise-custom-advertise-where)) supports the following:

<a id="vip-26d874"></a>&#x2022; [`default_v6_vip`](#vip-26d874) - Optional Block<br>Enable this option

<a id="vip-c51931"></a>&#x2022; [`default_vip`](#vip-c51931) - Optional Block<br>Enable this option

<a id="vip-bb67d7"></a>&#x2022; [`specific_v6_vip`](#vip-bb67d7) - Optional String<br>Specific V6 VIP. Use given IPv6 address as VIP on virtual Network

<a id="vip-943090"></a>&#x2022; [`specific_vip`](#vip-943090) - Optional String<br>Specific V4 VIP. Use given IPv4 address as VIP on virtual Network

<a id="network-bff334"></a>&#x2022; [`virtual_network`](#network-bff334) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Virtual Network](#network-bff334) below.

#### Advertise Custom Advertise Where Virtual Network Virtual Network

A [`virtual_network`](#network-bff334) block (within [`advertise_custom.advertise_where.virtual_network`](#network-a20be3)) supports the following:

<a id="name-5596bc"></a>&#x2022; [`name`](#name-5596bc) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-030577"></a>&#x2022; [`namespace`](#namespace-030577) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-72f925"></a>&#x2022; [`tenant`](#tenant-72f925) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Advertise Custom Advertise Where Virtual Site

A [`virtual_site`](#site-5d39fd) block (within [`advertise_custom.advertise_where`](#advertise-custom-advertise-where)) supports the following:

<a id="network-15aca4"></a>&#x2022; [`network`](#network-15aca4) - Optional String  Defaults to `SITE_NETWORK_INSIDE_AND_OUTSIDE`<br>Possible values are `SITE_NETWORK_INSIDE_AND_OUTSIDE`, `SITE_NETWORK_INSIDE`, `SITE_NETWORK_OUTSIDE`, `SITE_NETWORK_SERVICE`, `SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP`, `SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP`, `SITE_NETWORK_IP_FABRIC`<br>[Enum: SITE_NETWORK_INSIDE_AND_OUTSIDE|SITE_NETWORK_INSIDE|SITE_NETWORK_OUTSIDE|SITE_NETWORK_SERVICE|SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP|SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP|SITE_NETWORK_IP_FABRIC] Site Network. This defines network types to be used on site All inside and outside networks. All inside and outside networks with internet VIP support. All inside networks. All outside networks. All outside networks with internet VIP support. vK8s service network. - SITE_NETWORK_IP_FABRIC: VER IP Fabric network for the site This Virtual network type is used for exposing virtual host on IP Fabric network on the VER site or for endpoint in IP Fabric network

<a id="site-04fd53"></a>&#x2022; [`virtual_site`](#site-04fd53) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Virtual Site](#site-04fd53) below.

#### Advertise Custom Advertise Where Virtual Site Virtual Site

A [`virtual_site`](#site-04fd53) block (within [`advertise_custom.advertise_where.virtual_site`](#site-5d39fd)) supports the following:

<a id="name-b7ccc7"></a>&#x2022; [`name`](#name-b7ccc7) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-a4ffcf"></a>&#x2022; [`namespace`](#namespace-a4ffcf) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-637b28"></a>&#x2022; [`tenant`](#tenant-637b28) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Advertise Custom Advertise Where Virtual Site With VIP

A [`virtual_site_with_vip`](#vip-870b0b) block (within [`advertise_custom.advertise_where`](#advertise-custom-advertise-where)) supports the following:

<a id="vip-ip-4850ab"></a>&#x2022; [`ip`](#vip-ip-4850ab) - Optional String<br>IP Address. Use given IP address as VIP on the site

<a id="network-8b2765"></a>&#x2022; [`network`](#network-8b2765) - Optional String  Defaults to `SITE_NETWORK_SPECIFIED_VIP_OUTSIDE`<br>Possible values are `SITE_NETWORK_SPECIFIED_VIP_OUTSIDE`, `SITE_NETWORK_SPECIFIED_VIP_INSIDE`<br>[Enum: SITE_NETWORK_SPECIFIED_VIP_OUTSIDE|SITE_NETWORK_SPECIFIED_VIP_INSIDE] Site Network. This defines network types to be used on virtual-site with specified VIP All outside networks. All inside networks

<a id="site-ac753e"></a>&#x2022; [`virtual_site`](#site-ac753e) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Virtual Site](#site-ac753e) below.

#### Advertise Custom Advertise Where Virtual Site With VIP Virtual Site

A [`virtual_site`](#site-ac753e) block (within [`advertise_custom.advertise_where.virtual_site_with_vip`](#vip-870b0b)) supports the following:

<a id="name-5f7f0d"></a>&#x2022; [`name`](#name-5f7f0d) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-414bc8"></a>&#x2022; [`namespace`](#namespace-414bc8) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-130ed4"></a>&#x2022; [`tenant`](#tenant-130ed4) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Advertise Custom Advertise Where Vk8s Service

A [`vk8s_service`](#service-1fdc7a) block (within [`advertise_custom.advertise_where`](#advertise-custom-advertise-where)) supports the following:

<a id="site-ec8d32"></a>&#x2022; [`site`](#site-ec8d32) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Site](#site-ec8d32) below.

<a id="site-5fcbf9"></a>&#x2022; [`virtual_site`](#site-5fcbf9) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Virtual Site](#site-5fcbf9) below.

#### Advertise Custom Advertise Where Vk8s Service Site

A [`site`](#site-ec8d32) block (within [`advertise_custom.advertise_where.vk8s_service`](#service-1fdc7a)) supports the following:

<a id="name-950776"></a>&#x2022; [`name`](#name-950776) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-1faf25"></a>&#x2022; [`namespace`](#namespace-1faf25) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-98cf6a"></a>&#x2022; [`tenant`](#tenant-98cf6a) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Advertise Custom Advertise Where Vk8s Service Virtual Site

A [`virtual_site`](#site-5fcbf9) block (within [`advertise_custom.advertise_where.vk8s_service`](#service-1fdc7a)) supports the following:

<a id="name-1cf7c0"></a>&#x2022; [`name`](#name-1cf7c0) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-3dbb7e"></a>&#x2022; [`namespace`](#namespace-3dbb7e) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-38ddda"></a>&#x2022; [`tenant`](#tenant-38ddda) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Advertise On Public

An [`advertise_on_public`](#advertise-on-public) block supports the following:

<a id="advertise-on-public-public-ip"></a>&#x2022; [`public_ip`](#advertise-on-public-public-ip) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Public IP](#advertise-on-public-public-ip) below.

#### Advertise On Public Public IP

A [`public_ip`](#advertise-on-public-public-ip) block (within [`advertise_on_public`](#advertise-on-public)) supports the following:

<a id="advertise-on-public-public-ip-name"></a>&#x2022; [`name`](#advertise-on-public-public-ip-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="advertise-on-public-public-ip-namespace"></a>&#x2022; [`namespace`](#advertise-on-public-public-ip-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="advertise-on-public-public-ip-tenant"></a>&#x2022; [`tenant`](#advertise-on-public-public-ip-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### API Protection Rules

An [`api_protection_rules`](#api-protection-rules) block supports the following:

<a id="api-protection-rules-api-endpoint-rules"></a>&#x2022; [`api_endpoint_rules`](#api-protection-rules-api-endpoint-rules) - Optional Block<br>API Endpoints. This category defines specific rules per API endpoints. If request matches any of these rules, skipping second category rules<br>See [API Endpoint Rules](#api-protection-rules-api-endpoint-rules) below.

<a id="api-protection-rules-api-groups-rules"></a>&#x2022; [`api_groups_rules`](#api-protection-rules-api-groups-rules) - Optional Block<br>Server URLs and API Groups. This category includes rules per API group or Server URL. For API groups, refer to API Definition which includes API groups derived from uploaded swaggers<br>See [API Groups Rules](#api-protection-rules-api-groups-rules) below.

#### API Protection Rules API Endpoint Rules

An [`api_endpoint_rules`](#api-protection-rules-api-endpoint-rules) block (within [`api_protection_rules`](#api-protection-rules)) supports the following:

<a id="action-389797"></a>&#x2022; [`action`](#action-389797) - Optional Block<br>API Protection Rule Action. The action to take if the input request matches the rule<br>See [Action](#action-389797) below.

<a id="domain-c69c3a"></a>&#x2022; [`any_domain`](#domain-c69c3a) - Optional Block<br>Enable this option

<a id="method-361974"></a>&#x2022; [`api_endpoint_method`](#method-361974) - Optional Block<br>HTTP Method Matcher. A HTTP method matcher specifies a list of methods to match an input HTTP method. The match is considered successful if the input method is a member of the list. The result of the match based on the method list is inverted if invert_matcher is true<br>See [API Endpoint Method](#method-361974) below.

<a id="path-64d754"></a>&#x2022; [`api_endpoint_path`](#path-64d754) - Optional String<br>API Endpoint. The endpoint (path) of the request

<a id="matcher-af9b22"></a>&#x2022; [`client_matcher`](#matcher-af9b22) - Optional Block<br>Client Matcher. Client conditions for matching a rule<br>See [Client Matcher](#matcher-af9b22) below.

<a id="metadata-46451b"></a>&#x2022; [`metadata`](#metadata-46451b) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#metadata-46451b) below.

<a id="matcher-25d16b"></a>&#x2022; [`request_matcher`](#matcher-25d16b) - Optional Block<br>Request Matcher. Request conditions for matching a rule<br>See [Request Matcher](#matcher-25d16b) below.

<a id="domain-745a34"></a>&#x2022; [`specific_domain`](#domain-745a34) - Optional String<br>Specific Domain. The rule will apply for a specific domain. For example: API.example.com

#### API Protection Rules API Endpoint Rules Action

An [`action`](#action-389797) block (within [`api_protection_rules.api_endpoint_rules`](#api-protection-rules-api-endpoint-rules)) supports the following:

<a id="allow-9ca9d1"></a>&#x2022; [`allow`](#allow-9ca9d1) - Optional Block<br>Enable this option

<a id="deny-ec80de"></a>&#x2022; [`deny`](#deny-ec80de) - Optional Block<br>Enable this option

#### API Protection Rules API Endpoint Rules API Endpoint Method

An [`api_endpoint_method`](#method-361974) block (within [`api_protection_rules.api_endpoint_rules`](#api-protection-rules-api-endpoint-rules)) supports the following:

<a id="matcher-50f4d0"></a>&#x2022; [`invert_matcher`](#matcher-50f4d0) - Optional Bool<br>Invert Method Matcher. Invert the match result

<a id="methods-756da7"></a>&#x2022; [`methods`](#methods-756da7) - Optional List  Defaults to `ANY`<br>Possible values are `ANY`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`, `COPY`<br>[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY] Method List. List of methods values to match against

#### API Protection Rules API Endpoint Rules Client Matcher

A [`client_matcher`](#matcher-af9b22) block (within [`api_protection_rules.api_endpoint_rules`](#api-protection-rules-api-endpoint-rules)) supports the following:

<a id="client-978075"></a>&#x2022; [`any_client`](#client-978075) - Optional Block<br>Enable this option

<a id="any-ip-f91673"></a>&#x2022; [`any_ip`](#any-ip-f91673) - Optional Block<br>Enable this option

<a id="list-68c8bb"></a>&#x2022; [`asn_list`](#list-68c8bb) - Optional Block<br>ASN Match List. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer<br>See [Asn List](#list-68c8bb) below.

<a id="matcher-96a63e"></a>&#x2022; [`asn_matcher`](#matcher-96a63e) - Optional Block<br>ASN Matcher. Match any AS number contained in the list of bgp_asn_sets<br>See [Asn Matcher](#matcher-96a63e) below.

<a id="selector-a427f6"></a>&#x2022; [`client_selector`](#selector-a427f6) - Optional Block<br>Label Selector. This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expresssions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects. Label selector is immutable. expressions is a list of strings of label selection expression. Each string has ',' separated values which are 'AND' and all strings are logically 'OR'. BNF for expression string `<selector-syntax>` ::= `<requirement>` | `<requirement>` ',' `<selector-syntax>` `<requirement>` ::= [!] KEY [ `<set-based-restriction>` | `<exact-match-restriction>` ] `<set-based-restriction>` ::= '' | `<inclusion-exclusion>` `<value-set>` `<inclusion-exclusion>` ::= `<inclusion>` | `<exclusion>` `<exclusion>` ::= 'notin' `<inclusion>` ::= 'in' `<value-set>` ::= '(' `<values>` ')' `<values>` ::= VALUE | VALUE ',' `<values>` `<exact-match-restriction>` ::= ['='|'=='|'!='] VALUE<br>See [Client Selector](#selector-a427f6) below.

<a id="matcher-21e09e"></a>&#x2022; [`ip_matcher`](#matcher-21e09e) - Optional Block<br>IP Prefix Matcher. Match any IP prefix contained in the list of ip_prefix_sets. The result of the match is inverted if invert_matcher is true<br>See [IP Matcher](#matcher-21e09e) below.

<a id="list-1cb46a"></a>&#x2022; [`ip_prefix_list`](#list-1cb46a) - Optional Block<br>IP Prefix Match List. List of IP Prefix strings to match against<br>See [IP Prefix List](#list-1cb46a) below.

<a id="list-1acba1"></a>&#x2022; [`ip_threat_category_list`](#list-1acba1) - Optional Block<br>IP Threat Category List Type. List of IP threat categories<br>See [IP Threat Category List](#list-1acba1) below.

<a id="matcher-e92a66"></a>&#x2022; [`tls_fingerprint_matcher`](#matcher-e92a66) - Optional Block<br>TLS Fingerprint Matcher. A TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of supported positve match criteria includes a list of known classes of TLS fingerprints and a list of exact values. The match is considered successful if either of these positive criteria are satisfied and the input fingerprint is not one of the excluded values<br>See [TLS Fingerprint Matcher](#matcher-e92a66) below.

#### API Protection Rules API Endpoint Rules Client Matcher Asn List

An [`asn_list`](#list-68c8bb) block (within [`api_protection_rules.api_endpoint_rules.client_matcher`](#matcher-af9b22)) supports the following:

<a id="numbers-10b755"></a>&#x2022; [`as_numbers`](#numbers-10b755) - Optional List<br>AS Numbers. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer

#### API Protection Rules API Endpoint Rules Client Matcher Asn Matcher

An [`asn_matcher`](#matcher-96a63e) block (within [`api_protection_rules.api_endpoint_rules.client_matcher`](#matcher-af9b22)) supports the following:

<a id="sets-c9f052"></a>&#x2022; [`asn_sets`](#sets-c9f052) - Optional Block<br>BGP ASN Sets. A list of references to bgp_asn_set objects<br>See [Asn Sets](#sets-c9f052) below.

#### API Protection Rules API Endpoint Rules Client Matcher Asn Matcher Asn Sets

An [`asn_sets`](#sets-c9f052) block (within [`api_protection_rules.api_endpoint_rules.client_matcher.asn_matcher`](#matcher-96a63e)) supports the following:

<a id="kind-224d5f"></a>&#x2022; [`kind`](#kind-224d5f) - Optional String<br>Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route')

<a id="name-a13538"></a>&#x2022; [`name`](#name-a13538) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-1168dc"></a>&#x2022; [`namespace`](#namespace-1168dc) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-4d3914"></a>&#x2022; [`tenant`](#tenant-4d3914) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="uid-29bb76"></a>&#x2022; [`uid`](#uid-29bb76) - Optional String<br>UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid

#### API Protection Rules API Endpoint Rules Client Matcher Client Selector

A [`client_selector`](#selector-a427f6) block (within [`api_protection_rules.api_endpoint_rules.client_matcher`](#matcher-af9b22)) supports the following:

<a id="expressions-9a6713"></a>&#x2022; [`expressions`](#expressions-9a6713) - Optional List<br>Selector Expression. expressions contains the kubernetes style label expression for selections

#### API Protection Rules API Endpoint Rules Client Matcher IP Matcher

An [`ip_matcher`](#matcher-21e09e) block (within [`api_protection_rules.api_endpoint_rules.client_matcher`](#matcher-af9b22)) supports the following:

<a id="matcher-fe38a7"></a>&#x2022; [`invert_matcher`](#matcher-fe38a7) - Optional Bool<br>Invert IP Matcher. Invert the match result

<a id="sets-fb81b9"></a>&#x2022; [`prefix_sets`](#sets-fb81b9) - Optional Block<br>IP Prefix Sets. A list of references to ip_prefix_set objects<br>See [Prefix Sets](#sets-fb81b9) below.

#### API Protection Rules API Endpoint Rules Client Matcher IP Matcher Prefix Sets

A [`prefix_sets`](#sets-fb81b9) block (within [`api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher`](#matcher-21e09e)) supports the following:

<a id="kind-0809f8"></a>&#x2022; [`kind`](#kind-0809f8) - Optional String<br>Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route')

<a id="name-e5dfab"></a>&#x2022; [`name`](#name-e5dfab) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-6a977e"></a>&#x2022; [`namespace`](#namespace-6a977e) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-bcc048"></a>&#x2022; [`tenant`](#tenant-bcc048) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="uid-495add"></a>&#x2022; [`uid`](#uid-495add) - Optional String<br>UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid

#### API Protection Rules API Endpoint Rules Client Matcher IP Prefix List

An [`ip_prefix_list`](#list-1cb46a) block (within [`api_protection_rules.api_endpoint_rules.client_matcher`](#matcher-af9b22)) supports the following:

<a id="match-a9eb3d"></a>&#x2022; [`invert_match`](#match-a9eb3d) - Optional Bool<br>Invert Match Result. Invert the match result

<a id="prefixes-ea4a30"></a>&#x2022; [`ip_prefixes`](#prefixes-ea4a30) - Optional List<br>IPv4 Prefix List. List of IPv4 prefix strings

#### API Protection Rules API Endpoint Rules Client Matcher IP Threat Category List

An [`ip_threat_category_list`](#list-1acba1) block (within [`api_protection_rules.api_endpoint_rules.client_matcher`](#matcher-af9b22)) supports the following:

<a id="categories-e445ac"></a>&#x2022; [`ip_threat_categories`](#categories-e445ac) - Optional List  Defaults to `SPAM_SOURCES`<br>Possible values are `SPAM_SOURCES`, `WINDOWS_EXPLOITS`, `WEB_ATTACKS`, `BOTNETS`, `SCANNERS`, `REPUTATION`, `PHISHING`, `PROXY`, `MOBILE_THREATS`, `TOR_PROXY`, `DENIAL_OF_SERVICE`, `NETWORK`<br>[Enum: SPAM_SOURCES|WINDOWS_EXPLOITS|WEB_ATTACKS|BOTNETS|SCANNERS|REPUTATION|PHISHING|PROXY|MOBILE_THREATS|TOR_PROXY|DENIAL_OF_SERVICE|NETWORK] List of IP Threat Categories to choose. The IP threat categories is obtained from the list and is used to auto-generate equivalent label selection expressions

#### API Protection Rules API Endpoint Rules Client Matcher TLS Fingerprint Matcher

A [`tls_fingerprint_matcher`](#matcher-e92a66) block (within [`api_protection_rules.api_endpoint_rules.client_matcher`](#matcher-af9b22)) supports the following:

<a id="classes-0f7a74"></a>&#x2022; [`classes`](#classes-0f7a74) - Optional List  Defaults to `TLS_FINGERPRINT_NONE`<br>Possible values are `TLS_FINGERPRINT_NONE`, `ANY_MALICIOUS_FINGERPRINT`, `ADWARE`, `ADWIND`, `DRIDEX`, `GOOTKIT`, `GOZI`, `JBIFROST`, `QUAKBOT`, `RANSOMWARE`, `TROLDESH`, `TOFSEE`, `TORRENTLOCKER`, `TRICKBOT`<br>[Enum: TLS_FINGERPRINT_NONE|ANY_MALICIOUS_FINGERPRINT|ADWARE|ADWIND|DRIDEX|GOOTKIT|GOZI|JBIFROST|QUAKBOT|RANSOMWARE|TROLDESH|TOFSEE|TORRENTLOCKER|TRICKBOT] TLS fingerprint classes. A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against

<a id="values-05e60c"></a>&#x2022; [`exact_values`](#values-05e60c) - Optional List<br>Exact Values. A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against

<a id="values-c92a3a"></a>&#x2022; [`excluded_values`](#values-c92a3a) - Optional List<br>Excluded Values. A list of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can be used to skip known false positives when using one or more known TLS fingerprint classes in the enclosing matcher

#### API Protection Rules API Endpoint Rules Metadata

A [`metadata`](#metadata-46451b) block (within [`api_protection_rules.api_endpoint_rules`](#api-protection-rules-api-endpoint-rules)) supports the following:

<a id="spec-5eb432"></a>&#x2022; [`description_spec`](#spec-5eb432) - Optional String<br>Description. Human readable description

<a id="name-af3d78"></a>&#x2022; [`name`](#name-af3d78) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### API Protection Rules API Endpoint Rules Request Matcher

A [`request_matcher`](#matcher-25d16b) block (within [`api_protection_rules.api_endpoint_rules`](#api-protection-rules-api-endpoint-rules)) supports the following:

<a id="matchers-b1b3fa"></a>&#x2022; [`cookie_matchers`](#matchers-b1b3fa) - Optional Block<br>Cookie Matchers. A list of predicates for all cookies that need to be matched. The criteria for matching each cookie is described in individual instances of CookieMatcherType. The actual cookie values are extracted from the request API as a list of strings for each cookie name. Note that all specified cookie matcher predicates must evaluate to true<br>See [Cookie Matchers](#matchers-b1b3fa) below.

<a id="headers-67bd91"></a>&#x2022; [`headers`](#headers-67bd91) - Optional Block<br>HTTP Headers. A list of predicates for various HTTP headers that need to match. The criteria for matching each HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values are extracted from the request API as a list of strings for each HTTP header type. Note that all specified header predicates must evaluate to true<br>See [Headers](#headers-67bd91) below.

<a id="claims-d53204"></a>&#x2022; [`jwt_claims`](#claims-d53204) - Optional Block<br>JWT Claims. A list of predicates for various JWT claims that need to match. The criteria for matching each JWT claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are extracted from the JWT payload as a list of strings. Note that all specified JWT claim predicates must evaluate to true. Note that this feature only works on LBs with JWT Validation feature enabled<br>See [JWT Claims](#claims-d53204) below.

<a id="params-24d19c"></a>&#x2022; [`query_params`](#params-24d19c) - Optional Block<br>HTTP Query Parameters. A list of predicates for all query parameters that need to be matched. The criteria for matching each query parameter are described in individual instances of QueryParameterMatcherType. The actual query parameter values are extracted from the request API as a list of strings for each query parameter name. Note that all specified query parameter predicates must evaluate to true<br>See [Query Params](#params-24d19c) below.

#### API Protection Rules API Endpoint Rules Request Matcher Cookie Matchers

A [`cookie_matchers`](#matchers-b1b3fa) block (within [`api_protection_rules.api_endpoint_rules.request_matcher`](#matcher-25d16b)) supports the following:

<a id="present-3eca19"></a>&#x2022; [`check_not_present`](#present-3eca19) - Optional Block<br>Enable this option

<a id="present-c08ac5"></a>&#x2022; [`check_present`](#present-c08ac5) - Optional Block<br>Enable this option

<a id="matcher-09cbb6"></a>&#x2022; [`invert_matcher`](#matcher-09cbb6) - Optional Bool<br>Invert Matcher. Invert Match of the expression defined

<a id="item-174883"></a>&#x2022; [`item`](#item-174883) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#item-174883) below.

<a id="name-3b244c"></a>&#x2022; [`name`](#name-3b244c) - Optional String<br>Cookie Name. A case-sensitive cookie name

#### API Protection Rules API Endpoint Rules Request Matcher Cookie Matchers Item

An [`item`](#item-174883) block (within [`api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers`](#matchers-b1b3fa)) supports the following:

<a id="values-ed8412"></a>&#x2022; [`exact_values`](#values-ed8412) - Optional List<br>Exact Values. A list of exact values to match the input against

<a id="values-c7eaf4"></a>&#x2022; [`regex_values`](#values-c7eaf4) - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="transformers-c8c9a1"></a>&#x2022; [`transformers`](#transformers-c8c9a1) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>[Enum: LOWER_CASE|UPPER_CASE|BASE64_DECODE|NORMALIZE_PATH|REMOVE_WHITESPACE|URL_DECODE|TRIM_LEFT|TRIM_RIGHT|TRIM] Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### API Protection Rules API Endpoint Rules Request Matcher Headers

A [`headers`](#headers-67bd91) block (within [`api_protection_rules.api_endpoint_rules.request_matcher`](#matcher-25d16b)) supports the following:

<a id="present-f03e29"></a>&#x2022; [`check_not_present`](#present-f03e29) - Optional Block<br>Enable this option

<a id="present-9371fe"></a>&#x2022; [`check_present`](#present-9371fe) - Optional Block<br>Enable this option

<a id="matcher-036f4a"></a>&#x2022; [`invert_matcher`](#matcher-036f4a) - Optional Bool<br>Invert Header Matcher. Invert the match result

<a id="item-7e3514"></a>&#x2022; [`item`](#item-7e3514) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#item-7e3514) below.

<a id="name-a2b99d"></a>&#x2022; [`name`](#name-a2b99d) - Optional String<br>Header Name. A case-insensitive HTTP header name

#### API Protection Rules API Endpoint Rules Request Matcher Headers Item

An [`item`](#item-7e3514) block (within [`api_protection_rules.api_endpoint_rules.request_matcher.headers`](#headers-67bd91)) supports the following:

<a id="values-17b805"></a>&#x2022; [`exact_values`](#values-17b805) - Optional List<br>Exact Values. A list of exact values to match the input against

<a id="values-003282"></a>&#x2022; [`regex_values`](#values-003282) - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="transformers-54238d"></a>&#x2022; [`transformers`](#transformers-54238d) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>[Enum: LOWER_CASE|UPPER_CASE|BASE64_DECODE|NORMALIZE_PATH|REMOVE_WHITESPACE|URL_DECODE|TRIM_LEFT|TRIM_RIGHT|TRIM] Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### API Protection Rules API Endpoint Rules Request Matcher JWT Claims

A [`jwt_claims`](#claims-d53204) block (within [`api_protection_rules.api_endpoint_rules.request_matcher`](#matcher-25d16b)) supports the following:

<a id="present-abfe35"></a>&#x2022; [`check_not_present`](#present-abfe35) - Optional Block<br>Enable this option

<a id="present-dd34db"></a>&#x2022; [`check_present`](#present-dd34db) - Optional Block<br>Enable this option

<a id="matcher-c0c86a"></a>&#x2022; [`invert_matcher`](#matcher-c0c86a) - Optional Bool<br>Invert Matcher. Invert the match result

<a id="item-379a9f"></a>&#x2022; [`item`](#item-379a9f) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#item-379a9f) below.

<a id="name-d7fe9e"></a>&#x2022; [`name`](#name-d7fe9e) - Optional String<br>JWT Claim Name. JWT claim name

#### API Protection Rules API Endpoint Rules Request Matcher JWT Claims Item

An [`item`](#item-379a9f) block (within [`api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims`](#claims-d53204)) supports the following:

<a id="values-31b737"></a>&#x2022; [`exact_values`](#values-31b737) - Optional List<br>Exact Values. A list of exact values to match the input against

<a id="values-2843f6"></a>&#x2022; [`regex_values`](#values-2843f6) - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="transformers-a2eca2"></a>&#x2022; [`transformers`](#transformers-a2eca2) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>[Enum: LOWER_CASE|UPPER_CASE|BASE64_DECODE|NORMALIZE_PATH|REMOVE_WHITESPACE|URL_DECODE|TRIM_LEFT|TRIM_RIGHT|TRIM] Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### API Protection Rules API Endpoint Rules Request Matcher Query Params

A [`query_params`](#params-24d19c) block (within [`api_protection_rules.api_endpoint_rules.request_matcher`](#matcher-25d16b)) supports the following:

<a id="present-0e9ebe"></a>&#x2022; [`check_not_present`](#present-0e9ebe) - Optional Block<br>Enable this option

<a id="present-c54e0e"></a>&#x2022; [`check_present`](#present-c54e0e) - Optional Block<br>Enable this option

<a id="matcher-453f50"></a>&#x2022; [`invert_matcher`](#matcher-453f50) - Optional Bool<br>Invert Query Parameter Matcher. Invert the match result

<a id="item-f2caab"></a>&#x2022; [`item`](#item-f2caab) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#item-f2caab) below.

<a id="key-6d44ea"></a>&#x2022; [`key`](#key-6d44ea) - Optional String<br>Query Parameter Name. A case-sensitive HTTP query parameter name

#### API Protection Rules API Endpoint Rules Request Matcher Query Params Item

An [`item`](#item-f2caab) block (within [`api_protection_rules.api_endpoint_rules.request_matcher.query_params`](#params-24d19c)) supports the following:

<a id="values-d4c3cb"></a>&#x2022; [`exact_values`](#values-d4c3cb) - Optional List<br>Exact Values. A list of exact values to match the input against

<a id="values-423960"></a>&#x2022; [`regex_values`](#values-423960) - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="transformers-f4daba"></a>&#x2022; [`transformers`](#transformers-f4daba) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>[Enum: LOWER_CASE|UPPER_CASE|BASE64_DECODE|NORMALIZE_PATH|REMOVE_WHITESPACE|URL_DECODE|TRIM_LEFT|TRIM_RIGHT|TRIM] Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### API Protection Rules API Groups Rules

An [`api_groups_rules`](#api-protection-rules-api-groups-rules) block (within [`api_protection_rules`](#api-protection-rules)) supports the following:

<a id="action-fa62d7"></a>&#x2022; [`action`](#action-fa62d7) - Optional Block<br>API Protection Rule Action. The action to take if the input request matches the rule<br>See [Action](#action-fa62d7) below.

<a id="domain-b1276e"></a>&#x2022; [`any_domain`](#domain-b1276e) - Optional Block<br>Enable this option

<a id="group-a8b675"></a>&#x2022; [`api_group`](#group-a8b675) - Optional String<br>API Group. API groups derived from API Definition swaggers. For example oas-all-operations including all paths and methods from the swaggers, oas-base-urls covering all requests under base-paths from the swaggers. Custom groups can be created if user tags paths or operations with 'x-volterra-API-group' extensions inside swaggers

<a id="path-f85c4e"></a>&#x2022; [`base_path`](#path-f85c4e) - Optional String<br>Base Path. Prefix of the request path. For example: /v1

<a id="matcher-92ba86"></a>&#x2022; [`client_matcher`](#matcher-92ba86) - Optional Block<br>Client Matcher. Client conditions for matching a rule<br>See [Client Matcher](#matcher-92ba86) below.

<a id="metadata-b7fd60"></a>&#x2022; [`metadata`](#metadata-b7fd60) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#metadata-b7fd60) below.

<a id="matcher-58ddde"></a>&#x2022; [`request_matcher`](#matcher-58ddde) - Optional Block<br>Request Matcher. Request conditions for matching a rule<br>See [Request Matcher](#matcher-58ddde) below.

<a id="domain-3d8e2d"></a>&#x2022; [`specific_domain`](#domain-3d8e2d) - Optional String<br>Specific Domain. The rule will apply for a specific domain. For example: API.example.com

#### API Protection Rules API Groups Rules Action

An [`action`](#action-fa62d7) block (within [`api_protection_rules.api_groups_rules`](#api-protection-rules-api-groups-rules)) supports the following:

<a id="allow-eba8d3"></a>&#x2022; [`allow`](#allow-eba8d3) - Optional Block<br>Enable this option

<a id="deny-99c219"></a>&#x2022; [`deny`](#deny-99c219) - Optional Block<br>Enable this option

#### API Protection Rules API Groups Rules Client Matcher

A [`client_matcher`](#matcher-92ba86) block (within [`api_protection_rules.api_groups_rules`](#api-protection-rules-api-groups-rules)) supports the following:

<a id="client-9b9805"></a>&#x2022; [`any_client`](#client-9b9805) - Optional Block<br>Enable this option

<a id="any-ip-28f0aa"></a>&#x2022; [`any_ip`](#any-ip-28f0aa) - Optional Block<br>Enable this option

<a id="list-85a574"></a>&#x2022; [`asn_list`](#list-85a574) - Optional Block<br>ASN Match List. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer<br>See [Asn List](#list-85a574) below.

<a id="matcher-0187be"></a>&#x2022; [`asn_matcher`](#matcher-0187be) - Optional Block<br>ASN Matcher. Match any AS number contained in the list of bgp_asn_sets<br>See [Asn Matcher](#matcher-0187be) below.

<a id="selector-90369d"></a>&#x2022; [`client_selector`](#selector-90369d) - Optional Block<br>Label Selector. This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expresssions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects. Label selector is immutable. expressions is a list of strings of label selection expression. Each string has ',' separated values which are 'AND' and all strings are logically 'OR'. BNF for expression string `<selector-syntax>` ::= `<requirement>` | `<requirement>` ',' `<selector-syntax>` `<requirement>` ::= [!] KEY [ `<set-based-restriction>` | `<exact-match-restriction>` ] `<set-based-restriction>` ::= '' | `<inclusion-exclusion>` `<value-set>` `<inclusion-exclusion>` ::= `<inclusion>` | `<exclusion>` `<exclusion>` ::= 'notin' `<inclusion>` ::= 'in' `<value-set>` ::= '(' `<values>` ')' `<values>` ::= VALUE | VALUE ',' `<values>` `<exact-match-restriction>` ::= ['='|'=='|'!='] VALUE<br>See [Client Selector](#selector-90369d) below.

<a id="matcher-a763d9"></a>&#x2022; [`ip_matcher`](#matcher-a763d9) - Optional Block<br>IP Prefix Matcher. Match any IP prefix contained in the list of ip_prefix_sets. The result of the match is inverted if invert_matcher is true<br>See [IP Matcher](#matcher-a763d9) below.

<a id="list-7bdd48"></a>&#x2022; [`ip_prefix_list`](#list-7bdd48) - Optional Block<br>IP Prefix Match List. List of IP Prefix strings to match against<br>See [IP Prefix List](#list-7bdd48) below.

<a id="list-955575"></a>&#x2022; [`ip_threat_category_list`](#list-955575) - Optional Block<br>IP Threat Category List Type. List of IP threat categories<br>See [IP Threat Category List](#list-955575) below.

<a id="matcher-975eab"></a>&#x2022; [`tls_fingerprint_matcher`](#matcher-975eab) - Optional Block<br>TLS Fingerprint Matcher. A TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of supported positve match criteria includes a list of known classes of TLS fingerprints and a list of exact values. The match is considered successful if either of these positive criteria are satisfied and the input fingerprint is not one of the excluded values<br>See [TLS Fingerprint Matcher](#matcher-975eab) below.

#### API Protection Rules API Groups Rules Client Matcher Asn List

An [`asn_list`](#list-85a574) block (within [`api_protection_rules.api_groups_rules.client_matcher`](#matcher-92ba86)) supports the following:

<a id="numbers-85bf35"></a>&#x2022; [`as_numbers`](#numbers-85bf35) - Optional List<br>AS Numbers. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer

#### API Protection Rules API Groups Rules Client Matcher Asn Matcher

An [`asn_matcher`](#matcher-0187be) block (within [`api_protection_rules.api_groups_rules.client_matcher`](#matcher-92ba86)) supports the following:

<a id="sets-1b1b2c"></a>&#x2022; [`asn_sets`](#sets-1b1b2c) - Optional Block<br>BGP ASN Sets. A list of references to bgp_asn_set objects<br>See [Asn Sets](#sets-1b1b2c) below.

#### API Protection Rules API Groups Rules Client Matcher Asn Matcher Asn Sets

An [`asn_sets`](#sets-1b1b2c) block (within [`api_protection_rules.api_groups_rules.client_matcher.asn_matcher`](#matcher-0187be)) supports the following:

<a id="kind-556ecf"></a>&#x2022; [`kind`](#kind-556ecf) - Optional String<br>Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route')

<a id="name-9e3120"></a>&#x2022; [`name`](#name-9e3120) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-a9d0b5"></a>&#x2022; [`namespace`](#namespace-a9d0b5) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-161805"></a>&#x2022; [`tenant`](#tenant-161805) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="uid-1c9bd7"></a>&#x2022; [`uid`](#uid-1c9bd7) - Optional String<br>UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid

#### API Protection Rules API Groups Rules Client Matcher Client Selector

A [`client_selector`](#selector-90369d) block (within [`api_protection_rules.api_groups_rules.client_matcher`](#matcher-92ba86)) supports the following:

<a id="expressions-303af5"></a>&#x2022; [`expressions`](#expressions-303af5) - Optional List<br>Selector Expression. expressions contains the kubernetes style label expression for selections

#### API Protection Rules API Groups Rules Client Matcher IP Matcher

An [`ip_matcher`](#matcher-a763d9) block (within [`api_protection_rules.api_groups_rules.client_matcher`](#matcher-92ba86)) supports the following:

<a id="matcher-92f5c2"></a>&#x2022; [`invert_matcher`](#matcher-92f5c2) - Optional Bool<br>Invert IP Matcher. Invert the match result

<a id="sets-775479"></a>&#x2022; [`prefix_sets`](#sets-775479) - Optional Block<br>IP Prefix Sets. A list of references to ip_prefix_set objects<br>See [Prefix Sets](#sets-775479) below.

#### API Protection Rules API Groups Rules Client Matcher IP Matcher Prefix Sets

A [`prefix_sets`](#sets-775479) block (within [`api_protection_rules.api_groups_rules.client_matcher.ip_matcher`](#matcher-a763d9)) supports the following:

<a id="kind-829795"></a>&#x2022; [`kind`](#kind-829795) - Optional String<br>Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route')

<a id="name-b546eb"></a>&#x2022; [`name`](#name-b546eb) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-27de4d"></a>&#x2022; [`namespace`](#namespace-27de4d) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-c119cd"></a>&#x2022; [`tenant`](#tenant-c119cd) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="uid-e470a6"></a>&#x2022; [`uid`](#uid-e470a6) - Optional String<br>UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid

#### API Protection Rules API Groups Rules Client Matcher IP Prefix List

An [`ip_prefix_list`](#list-7bdd48) block (within [`api_protection_rules.api_groups_rules.client_matcher`](#matcher-92ba86)) supports the following:

<a id="match-4f1909"></a>&#x2022; [`invert_match`](#match-4f1909) - Optional Bool<br>Invert Match Result. Invert the match result

<a id="prefixes-927d73"></a>&#x2022; [`ip_prefixes`](#prefixes-927d73) - Optional List<br>IPv4 Prefix List. List of IPv4 prefix strings

#### API Protection Rules API Groups Rules Client Matcher IP Threat Category List

An [`ip_threat_category_list`](#list-955575) block (within [`api_protection_rules.api_groups_rules.client_matcher`](#matcher-92ba86)) supports the following:

<a id="categories-fd5336"></a>&#x2022; [`ip_threat_categories`](#categories-fd5336) - Optional List  Defaults to `SPAM_SOURCES`<br>Possible values are `SPAM_SOURCES`, `WINDOWS_EXPLOITS`, `WEB_ATTACKS`, `BOTNETS`, `SCANNERS`, `REPUTATION`, `PHISHING`, `PROXY`, `MOBILE_THREATS`, `TOR_PROXY`, `DENIAL_OF_SERVICE`, `NETWORK`<br>[Enum: SPAM_SOURCES|WINDOWS_EXPLOITS|WEB_ATTACKS|BOTNETS|SCANNERS|REPUTATION|PHISHING|PROXY|MOBILE_THREATS|TOR_PROXY|DENIAL_OF_SERVICE|NETWORK] List of IP Threat Categories to choose. The IP threat categories is obtained from the list and is used to auto-generate equivalent label selection expressions

#### API Protection Rules API Groups Rules Client Matcher TLS Fingerprint Matcher

A [`tls_fingerprint_matcher`](#matcher-975eab) block (within [`api_protection_rules.api_groups_rules.client_matcher`](#matcher-92ba86)) supports the following:

<a id="classes-73763e"></a>&#x2022; [`classes`](#classes-73763e) - Optional List  Defaults to `TLS_FINGERPRINT_NONE`<br>Possible values are `TLS_FINGERPRINT_NONE`, `ANY_MALICIOUS_FINGERPRINT`, `ADWARE`, `ADWIND`, `DRIDEX`, `GOOTKIT`, `GOZI`, `JBIFROST`, `QUAKBOT`, `RANSOMWARE`, `TROLDESH`, `TOFSEE`, `TORRENTLOCKER`, `TRICKBOT`<br>[Enum: TLS_FINGERPRINT_NONE|ANY_MALICIOUS_FINGERPRINT|ADWARE|ADWIND|DRIDEX|GOOTKIT|GOZI|JBIFROST|QUAKBOT|RANSOMWARE|TROLDESH|TOFSEE|TORRENTLOCKER|TRICKBOT] TLS fingerprint classes. A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against

<a id="values-eb1458"></a>&#x2022; [`exact_values`](#values-eb1458) - Optional List<br>Exact Values. A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against

<a id="values-4c4633"></a>&#x2022; [`excluded_values`](#values-4c4633) - Optional List<br>Excluded Values. A list of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can be used to skip known false positives when using one or more known TLS fingerprint classes in the enclosing matcher

#### API Protection Rules API Groups Rules Metadata

A [`metadata`](#metadata-b7fd60) block (within [`api_protection_rules.api_groups_rules`](#api-protection-rules-api-groups-rules)) supports the following:

<a id="spec-ccf62e"></a>&#x2022; [`description_spec`](#spec-ccf62e) - Optional String<br>Description. Human readable description

<a id="name-4148ef"></a>&#x2022; [`name`](#name-4148ef) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### API Protection Rules API Groups Rules Request Matcher

A [`request_matcher`](#matcher-58ddde) block (within [`api_protection_rules.api_groups_rules`](#api-protection-rules-api-groups-rules)) supports the following:

<a id="matchers-39d3e2"></a>&#x2022; [`cookie_matchers`](#matchers-39d3e2) - Optional Block<br>Cookie Matchers. A list of predicates for all cookies that need to be matched. The criteria for matching each cookie is described in individual instances of CookieMatcherType. The actual cookie values are extracted from the request API as a list of strings for each cookie name. Note that all specified cookie matcher predicates must evaluate to true<br>See [Cookie Matchers](#matchers-39d3e2) below.

<a id="headers-057c89"></a>&#x2022; [`headers`](#headers-057c89) - Optional Block<br>HTTP Headers. A list of predicates for various HTTP headers that need to match. The criteria for matching each HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values are extracted from the request API as a list of strings for each HTTP header type. Note that all specified header predicates must evaluate to true<br>See [Headers](#headers-057c89) below.

<a id="claims-954a59"></a>&#x2022; [`jwt_claims`](#claims-954a59) - Optional Block<br>JWT Claims. A list of predicates for various JWT claims that need to match. The criteria for matching each JWT claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are extracted from the JWT payload as a list of strings. Note that all specified JWT claim predicates must evaluate to true. Note that this feature only works on LBs with JWT Validation feature enabled<br>See [JWT Claims](#claims-954a59) below.

<a id="params-0e278f"></a>&#x2022; [`query_params`](#params-0e278f) - Optional Block<br>HTTP Query Parameters. A list of predicates for all query parameters that need to be matched. The criteria for matching each query parameter are described in individual instances of QueryParameterMatcherType. The actual query parameter values are extracted from the request API as a list of strings for each query parameter name. Note that all specified query parameter predicates must evaluate to true<br>See [Query Params](#params-0e278f) below.

#### API Protection Rules API Groups Rules Request Matcher Cookie Matchers

A [`cookie_matchers`](#matchers-39d3e2) block (within [`api_protection_rules.api_groups_rules.request_matcher`](#matcher-58ddde)) supports the following:

<a id="present-705c01"></a>&#x2022; [`check_not_present`](#present-705c01) - Optional Block<br>Enable this option

<a id="present-a826ac"></a>&#x2022; [`check_present`](#present-a826ac) - Optional Block<br>Enable this option

<a id="matcher-92b61c"></a>&#x2022; [`invert_matcher`](#matcher-92b61c) - Optional Bool<br>Invert Matcher. Invert Match of the expression defined

<a id="item-5daff3"></a>&#x2022; [`item`](#item-5daff3) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#item-5daff3) below.

<a id="name-3fdad7"></a>&#x2022; [`name`](#name-3fdad7) - Optional String<br>Cookie Name. A case-sensitive cookie name

#### API Protection Rules API Groups Rules Request Matcher Cookie Matchers Item

An [`item`](#item-5daff3) block (within [`api_protection_rules.api_groups_rules.request_matcher.cookie_matchers`](#matchers-39d3e2)) supports the following:

<a id="values-fcea74"></a>&#x2022; [`exact_values`](#values-fcea74) - Optional List<br>Exact Values. A list of exact values to match the input against

<a id="values-8a2b84"></a>&#x2022; [`regex_values`](#values-8a2b84) - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="transformers-e23344"></a>&#x2022; [`transformers`](#transformers-e23344) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>[Enum: LOWER_CASE|UPPER_CASE|BASE64_DECODE|NORMALIZE_PATH|REMOVE_WHITESPACE|URL_DECODE|TRIM_LEFT|TRIM_RIGHT|TRIM] Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### API Protection Rules API Groups Rules Request Matcher Headers

A [`headers`](#headers-057c89) block (within [`api_protection_rules.api_groups_rules.request_matcher`](#matcher-58ddde)) supports the following:

<a id="present-ae8200"></a>&#x2022; [`check_not_present`](#present-ae8200) - Optional Block<br>Enable this option

<a id="present-e949a8"></a>&#x2022; [`check_present`](#present-e949a8) - Optional Block<br>Enable this option

<a id="matcher-1d9a89"></a>&#x2022; [`invert_matcher`](#matcher-1d9a89) - Optional Bool<br>Invert Header Matcher. Invert the match result

<a id="item-cc1169"></a>&#x2022; [`item`](#item-cc1169) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#item-cc1169) below.

<a id="name-a446bf"></a>&#x2022; [`name`](#name-a446bf) - Optional String<br>Header Name. A case-insensitive HTTP header name

#### API Protection Rules API Groups Rules Request Matcher Headers Item

An [`item`](#item-cc1169) block (within [`api_protection_rules.api_groups_rules.request_matcher.headers`](#headers-057c89)) supports the following:

<a id="values-203803"></a>&#x2022; [`exact_values`](#values-203803) - Optional List<br>Exact Values. A list of exact values to match the input against

<a id="values-939c6a"></a>&#x2022; [`regex_values`](#values-939c6a) - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="transformers-a38138"></a>&#x2022; [`transformers`](#transformers-a38138) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>[Enum: LOWER_CASE|UPPER_CASE|BASE64_DECODE|NORMALIZE_PATH|REMOVE_WHITESPACE|URL_DECODE|TRIM_LEFT|TRIM_RIGHT|TRIM] Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### API Protection Rules API Groups Rules Request Matcher JWT Claims

A [`jwt_claims`](#claims-954a59) block (within [`api_protection_rules.api_groups_rules.request_matcher`](#matcher-58ddde)) supports the following:

<a id="present-98f8dd"></a>&#x2022; [`check_not_present`](#present-98f8dd) - Optional Block<br>Enable this option

<a id="present-44c204"></a>&#x2022; [`check_present`](#present-44c204) - Optional Block<br>Enable this option

<a id="matcher-73853f"></a>&#x2022; [`invert_matcher`](#matcher-73853f) - Optional Bool<br>Invert Matcher. Invert the match result

<a id="item-4b994e"></a>&#x2022; [`item`](#item-4b994e) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#item-4b994e) below.

<a id="name-29aac4"></a>&#x2022; [`name`](#name-29aac4) - Optional String<br>JWT Claim Name. JWT claim name

#### API Protection Rules API Groups Rules Request Matcher JWT Claims Item

An [`item`](#item-4b994e) block (within [`api_protection_rules.api_groups_rules.request_matcher.jwt_claims`](#claims-954a59)) supports the following:

<a id="values-52617f"></a>&#x2022; [`exact_values`](#values-52617f) - Optional List<br>Exact Values. A list of exact values to match the input against

<a id="values-677016"></a>&#x2022; [`regex_values`](#values-677016) - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="transformers-0d2f06"></a>&#x2022; [`transformers`](#transformers-0d2f06) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>[Enum: LOWER_CASE|UPPER_CASE|BASE64_DECODE|NORMALIZE_PATH|REMOVE_WHITESPACE|URL_DECODE|TRIM_LEFT|TRIM_RIGHT|TRIM] Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### API Protection Rules API Groups Rules Request Matcher Query Params

A [`query_params`](#params-0e278f) block (within [`api_protection_rules.api_groups_rules.request_matcher`](#matcher-58ddde)) supports the following:

<a id="present-cfb48e"></a>&#x2022; [`check_not_present`](#present-cfb48e) - Optional Block<br>Enable this option

<a id="present-ba6401"></a>&#x2022; [`check_present`](#present-ba6401) - Optional Block<br>Enable this option

<a id="matcher-b98619"></a>&#x2022; [`invert_matcher`](#matcher-b98619) - Optional Bool<br>Invert Query Parameter Matcher. Invert the match result

<a id="item-6fd7a8"></a>&#x2022; [`item`](#item-6fd7a8) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#item-6fd7a8) below.

<a id="key-854ad3"></a>&#x2022; [`key`](#key-854ad3) - Optional String<br>Query Parameter Name. A case-sensitive HTTP query parameter name

#### API Protection Rules API Groups Rules Request Matcher Query Params Item

An [`item`](#item-6fd7a8) block (within [`api_protection_rules.api_groups_rules.request_matcher.query_params`](#params-0e278f)) supports the following:

<a id="values-d52b3b"></a>&#x2022; [`exact_values`](#values-d52b3b) - Optional List<br>Exact Values. A list of exact values to match the input against

<a id="values-b02c56"></a>&#x2022; [`regex_values`](#values-b02c56) - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="transformers-24c720"></a>&#x2022; [`transformers`](#transformers-24c720) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>[Enum: LOWER_CASE|UPPER_CASE|BASE64_DECODE|NORMALIZE_PATH|REMOVE_WHITESPACE|URL_DECODE|TRIM_LEFT|TRIM_RIGHT|TRIM] Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### API Rate Limit

An [`api_rate_limit`](#api-rate-limit) block supports the following:

<a id="api-rate-limit-api-endpoint-rules"></a>&#x2022; [`api_endpoint_rules`](#api-rate-limit-api-endpoint-rules) - Optional Block<br>API Endpoints. Sets of rules for a specific endpoints. Order is matter as it uses first match policy. For creating rule that contain a whole domain or group of endpoints, please use the server URL rules above<br>See [API Endpoint Rules](#api-rate-limit-api-endpoint-rules) below.

<a id="rules-776e97"></a>&#x2022; [`bypass_rate_limiting_rules`](#rules-776e97) - Optional Block<br>Bypass Rate Limiting. This category defines rules per URL or API group. If request matches any of these rules, skip Rate Limiting<br>See [Bypass Rate Limiting Rules](#rules-776e97) below.

<a id="api-rate-limit-custom-ip-allowed-list"></a>&#x2022; [`custom_ip_allowed_list`](#api-rate-limit-custom-ip-allowed-list) - Optional Block<br>Custom IP Allowed List. IP Allowed list using existing ip_prefix_set objects<br>See [Custom IP Allowed List](#api-rate-limit-custom-ip-allowed-list) below.

<a id="api-rate-limit-ip-allowed-list"></a>&#x2022; [`ip_allowed_list`](#api-rate-limit-ip-allowed-list) - Optional Block<br>IPv4 Prefix List. List of IPv4 prefixes that represent an endpoint<br>See [IP Allowed List](#api-rate-limit-ip-allowed-list) below.

<a id="api-rate-limit-no-ip-allowed-list"></a>&#x2022; [`no_ip_allowed_list`](#api-rate-limit-no-ip-allowed-list) - Optional Block<br>Enable this option

<a id="api-rate-limit-server-url-rules"></a>&#x2022; [`server_url_rules`](#api-rate-limit-server-url-rules) - Optional Block<br>Server URLs. Set of rules for entire domain or base path that contain multiple endpoints. Order is matter as it uses first match policy. For matching also specific endpoints you can use the API endpoint rules set bellow<br>See [Server URL Rules](#api-rate-limit-server-url-rules) below.

#### API Rate Limit API Endpoint Rules

An [`api_endpoint_rules`](#api-rate-limit-api-endpoint-rules) block (within [`api_rate_limit`](#api-rate-limit)) supports the following:

<a id="domain-cf087c"></a>&#x2022; [`any_domain`](#domain-cf087c) - Optional Block<br>Enable this option

<a id="method-1e49b0"></a>&#x2022; [`api_endpoint_method`](#method-1e49b0) - Optional Block<br>HTTP Method Matcher. A HTTP method matcher specifies a list of methods to match an input HTTP method. The match is considered successful if the input method is a member of the list. The result of the match based on the method list is inverted if invert_matcher is true<br>See [API Endpoint Method](#method-1e49b0) below.

<a id="path-297bf2"></a>&#x2022; [`api_endpoint_path`](#path-297bf2) - Optional String<br>API Endpoint. The endpoint (path) of the request

<a id="matcher-794c7c"></a>&#x2022; [`client_matcher`](#matcher-794c7c) - Optional Block<br>Client Matcher. Client conditions for matching a rule<br>See [Client Matcher](#matcher-794c7c) below.

<a id="limiter-38a124"></a>&#x2022; [`inline_rate_limiter`](#limiter-38a124) - Optional Block<br>InlineRateLimiter<br>See [Inline Rate Limiter](#limiter-38a124) below.

<a id="limiter-f23897"></a>&#x2022; [`ref_rate_limiter`](#limiter-f23897) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Ref Rate Limiter](#limiter-f23897) below.

<a id="matcher-869fa1"></a>&#x2022; [`request_matcher`](#matcher-869fa1) - Optional Block<br>Request Matcher. Request conditions for matching a rule<br>See [Request Matcher](#matcher-869fa1) below.

<a id="domain-1ce4ba"></a>&#x2022; [`specific_domain`](#domain-1ce4ba) - Optional String<br>Specific Domain. The rule will apply for a specific domain

#### API Rate Limit API Endpoint Rules API Endpoint Method

An [`api_endpoint_method`](#method-1e49b0) block (within [`api_rate_limit.api_endpoint_rules`](#api-rate-limit-api-endpoint-rules)) supports the following:

<a id="matcher-25dd70"></a>&#x2022; [`invert_matcher`](#matcher-25dd70) - Optional Bool<br>Invert Method Matcher. Invert the match result

<a id="methods-bf7e55"></a>&#x2022; [`methods`](#methods-bf7e55) - Optional List  Defaults to `ANY`<br>Possible values are `ANY`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`, `COPY`<br>[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY] Method List. List of methods values to match against

#### API Rate Limit API Endpoint Rules Client Matcher

A [`client_matcher`](#matcher-794c7c) block (within [`api_rate_limit.api_endpoint_rules`](#api-rate-limit-api-endpoint-rules)) supports the following:

<a id="client-dd4b14"></a>&#x2022; [`any_client`](#client-dd4b14) - Optional Block<br>Enable this option

<a id="any-ip-2507e5"></a>&#x2022; [`any_ip`](#any-ip-2507e5) - Optional Block<br>Enable this option

<a id="list-541161"></a>&#x2022; [`asn_list`](#list-541161) - Optional Block<br>ASN Match List. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer<br>See [Asn List](#list-541161) below.

<a id="matcher-cd99ef"></a>&#x2022; [`asn_matcher`](#matcher-cd99ef) - Optional Block<br>ASN Matcher. Match any AS number contained in the list of bgp_asn_sets<br>See [Asn Matcher](#matcher-cd99ef) below.

<a id="selector-1b39eb"></a>&#x2022; [`client_selector`](#selector-1b39eb) - Optional Block<br>Label Selector. This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expresssions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects. Label selector is immutable. expressions is a list of strings of label selection expression. Each string has ',' separated values which are 'AND' and all strings are logically 'OR'. BNF for expression string `<selector-syntax>` ::= `<requirement>` | `<requirement>` ',' `<selector-syntax>` `<requirement>` ::= [!] KEY [ `<set-based-restriction>` | `<exact-match-restriction>` ] `<set-based-restriction>` ::= '' | `<inclusion-exclusion>` `<value-set>` `<inclusion-exclusion>` ::= `<inclusion>` | `<exclusion>` `<exclusion>` ::= 'notin' `<inclusion>` ::= 'in' `<value-set>` ::= '(' `<values>` ')' `<values>` ::= VALUE | VALUE ',' `<values>` `<exact-match-restriction>` ::= ['='|'=='|'!='] VALUE<br>See [Client Selector](#selector-1b39eb) below.

<a id="matcher-2d1e1b"></a>&#x2022; [`ip_matcher`](#matcher-2d1e1b) - Optional Block<br>IP Prefix Matcher. Match any IP prefix contained in the list of ip_prefix_sets. The result of the match is inverted if invert_matcher is true<br>See [IP Matcher](#matcher-2d1e1b) below.

<a id="list-19e2d7"></a>&#x2022; [`ip_prefix_list`](#list-19e2d7) - Optional Block<br>IP Prefix Match List. List of IP Prefix strings to match against<br>See [IP Prefix List](#list-19e2d7) below.

<a id="list-d4ce55"></a>&#x2022; [`ip_threat_category_list`](#list-d4ce55) - Optional Block<br>IP Threat Category List Type. List of IP threat categories<br>See [IP Threat Category List](#list-d4ce55) below.

<a id="matcher-ab7cce"></a>&#x2022; [`tls_fingerprint_matcher`](#matcher-ab7cce) - Optional Block<br>TLS Fingerprint Matcher. A TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of supported positve match criteria includes a list of known classes of TLS fingerprints and a list of exact values. The match is considered successful if either of these positive criteria are satisfied and the input fingerprint is not one of the excluded values<br>See [TLS Fingerprint Matcher](#matcher-ab7cce) below.

#### API Rate Limit API Endpoint Rules Client Matcher Asn List

An [`asn_list`](#list-541161) block (within [`api_rate_limit.api_endpoint_rules.client_matcher`](#matcher-794c7c)) supports the following:

<a id="numbers-7bb86e"></a>&#x2022; [`as_numbers`](#numbers-7bb86e) - Optional List<br>AS Numbers. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer

#### API Rate Limit API Endpoint Rules Client Matcher Asn Matcher

An [`asn_matcher`](#matcher-cd99ef) block (within [`api_rate_limit.api_endpoint_rules.client_matcher`](#matcher-794c7c)) supports the following:

<a id="sets-d85457"></a>&#x2022; [`asn_sets`](#sets-d85457) - Optional Block<br>BGP ASN Sets. A list of references to bgp_asn_set objects<br>See [Asn Sets](#sets-d85457) below.

#### API Rate Limit API Endpoint Rules Client Matcher Asn Matcher Asn Sets

An [`asn_sets`](#sets-d85457) block (within [`api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher`](#matcher-cd99ef)) supports the following:

<a id="kind-d515ee"></a>&#x2022; [`kind`](#kind-d515ee) - Optional String<br>Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route')

<a id="name-8f5645"></a>&#x2022; [`name`](#name-8f5645) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-e278b9"></a>&#x2022; [`namespace`](#namespace-e278b9) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-e59381"></a>&#x2022; [`tenant`](#tenant-e59381) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="uid-3ba47f"></a>&#x2022; [`uid`](#uid-3ba47f) - Optional String<br>UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid

#### API Rate Limit API Endpoint Rules Client Matcher Client Selector

A [`client_selector`](#selector-1b39eb) block (within [`api_rate_limit.api_endpoint_rules.client_matcher`](#matcher-794c7c)) supports the following:

<a id="expressions-f101e1"></a>&#x2022; [`expressions`](#expressions-f101e1) - Optional List<br>Selector Expression. expressions contains the kubernetes style label expression for selections

#### API Rate Limit API Endpoint Rules Client Matcher IP Matcher

An [`ip_matcher`](#matcher-2d1e1b) block (within [`api_rate_limit.api_endpoint_rules.client_matcher`](#matcher-794c7c)) supports the following:

<a id="matcher-91fa13"></a>&#x2022; [`invert_matcher`](#matcher-91fa13) - Optional Bool<br>Invert IP Matcher. Invert the match result

<a id="sets-cb5183"></a>&#x2022; [`prefix_sets`](#sets-cb5183) - Optional Block<br>IP Prefix Sets. A list of references to ip_prefix_set objects<br>See [Prefix Sets](#sets-cb5183) below.

#### API Rate Limit API Endpoint Rules Client Matcher IP Matcher Prefix Sets

A [`prefix_sets`](#sets-cb5183) block (within [`api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher`](#matcher-2d1e1b)) supports the following:

<a id="kind-ccd934"></a>&#x2022; [`kind`](#kind-ccd934) - Optional String<br>Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route')

<a id="name-19bb1b"></a>&#x2022; [`name`](#name-19bb1b) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-c8c75b"></a>&#x2022; [`namespace`](#namespace-c8c75b) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-f57c6e"></a>&#x2022; [`tenant`](#tenant-f57c6e) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="uid-97a716"></a>&#x2022; [`uid`](#uid-97a716) - Optional String<br>UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid

#### API Rate Limit API Endpoint Rules Client Matcher IP Prefix List

An [`ip_prefix_list`](#list-19e2d7) block (within [`api_rate_limit.api_endpoint_rules.client_matcher`](#matcher-794c7c)) supports the following:

<a id="match-169cde"></a>&#x2022; [`invert_match`](#match-169cde) - Optional Bool<br>Invert Match Result. Invert the match result

<a id="prefixes-c54632"></a>&#x2022; [`ip_prefixes`](#prefixes-c54632) - Optional List<br>IPv4 Prefix List. List of IPv4 prefix strings

#### API Rate Limit API Endpoint Rules Client Matcher IP Threat Category List

An [`ip_threat_category_list`](#list-d4ce55) block (within [`api_rate_limit.api_endpoint_rules.client_matcher`](#matcher-794c7c)) supports the following:

<a id="categories-33fa18"></a>&#x2022; [`ip_threat_categories`](#categories-33fa18) - Optional List  Defaults to `SPAM_SOURCES`<br>Possible values are `SPAM_SOURCES`, `WINDOWS_EXPLOITS`, `WEB_ATTACKS`, `BOTNETS`, `SCANNERS`, `REPUTATION`, `PHISHING`, `PROXY`, `MOBILE_THREATS`, `TOR_PROXY`, `DENIAL_OF_SERVICE`, `NETWORK`<br>[Enum: SPAM_SOURCES|WINDOWS_EXPLOITS|WEB_ATTACKS|BOTNETS|SCANNERS|REPUTATION|PHISHING|PROXY|MOBILE_THREATS|TOR_PROXY|DENIAL_OF_SERVICE|NETWORK] List of IP Threat Categories to choose. The IP threat categories is obtained from the list and is used to auto-generate equivalent label selection expressions

#### API Rate Limit API Endpoint Rules Client Matcher TLS Fingerprint Matcher

A [`tls_fingerprint_matcher`](#matcher-ab7cce) block (within [`api_rate_limit.api_endpoint_rules.client_matcher`](#matcher-794c7c)) supports the following:

<a id="classes-fe5ffc"></a>&#x2022; [`classes`](#classes-fe5ffc) - Optional List  Defaults to `TLS_FINGERPRINT_NONE`<br>Possible values are `TLS_FINGERPRINT_NONE`, `ANY_MALICIOUS_FINGERPRINT`, `ADWARE`, `ADWIND`, `DRIDEX`, `GOOTKIT`, `GOZI`, `JBIFROST`, `QUAKBOT`, `RANSOMWARE`, `TROLDESH`, `TOFSEE`, `TORRENTLOCKER`, `TRICKBOT`<br>[Enum: TLS_FINGERPRINT_NONE|ANY_MALICIOUS_FINGERPRINT|ADWARE|ADWIND|DRIDEX|GOOTKIT|GOZI|JBIFROST|QUAKBOT|RANSOMWARE|TROLDESH|TOFSEE|TORRENTLOCKER|TRICKBOT] TLS fingerprint classes. A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against

<a id="values-442dc5"></a>&#x2022; [`exact_values`](#values-442dc5) - Optional List<br>Exact Values. A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against

<a id="values-ea7eff"></a>&#x2022; [`excluded_values`](#values-ea7eff) - Optional List<br>Excluded Values. A list of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can be used to skip known false positives when using one or more known TLS fingerprint classes in the enclosing matcher

#### API Rate Limit API Endpoint Rules Inline Rate Limiter

An [`inline_rate_limiter`](#limiter-38a124) block (within [`api_rate_limit.api_endpoint_rules`](#api-rate-limit-api-endpoint-rules)) supports the following:

<a id="user-id-48be18"></a>&#x2022; [`ref_user_id`](#user-id-48be18) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Ref User ID](#user-id-48be18) below.

<a id="threshold-e13d5c"></a>&#x2022; [`threshold`](#threshold-e13d5c) - Optional Number<br>Threshold. The total number of allowed requests for 1 unit (e.g. SECOND/MINUTE/HOUR etc.) of the specified period

<a id="unit-4402df"></a>&#x2022; [`unit`](#unit-4402df) - Optional String  Defaults to `SECOND`<br>Possible values are `SECOND`, `MINUTE`, `HOUR`<br>[Enum: SECOND|MINUTE|HOUR] Rate Limit Period Unit. Unit for the period per which the rate limit is applied. - SECOND: Second Rate limit period unit is seconds - MINUTE: Minute Rate limit period unit is minutes - HOUR: Hour Rate limit period unit is hours - DAY: Day Rate limit period unit is days

<a id="user-id-ddc28d"></a>&#x2022; [`use_http_lb_user_id`](#user-id-ddc28d) - Optional Block<br>Enable this option

#### API Rate Limit API Endpoint Rules Inline Rate Limiter Ref User ID

A [`ref_user_id`](#user-id-48be18) block (within [`api_rate_limit.api_endpoint_rules.inline_rate_limiter`](#limiter-38a124)) supports the following:

<a id="name-44d974"></a>&#x2022; [`name`](#name-44d974) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-7c0a17"></a>&#x2022; [`namespace`](#namespace-7c0a17) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-568dd6"></a>&#x2022; [`tenant`](#tenant-568dd6) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### API Rate Limit API Endpoint Rules Ref Rate Limiter

A [`ref_rate_limiter`](#limiter-f23897) block (within [`api_rate_limit.api_endpoint_rules`](#api-rate-limit-api-endpoint-rules)) supports the following:

<a id="name-feb57b"></a>&#x2022; [`name`](#name-feb57b) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-8702e9"></a>&#x2022; [`namespace`](#namespace-8702e9) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-93c30b"></a>&#x2022; [`tenant`](#tenant-93c30b) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### API Rate Limit API Endpoint Rules Request Matcher

A [`request_matcher`](#matcher-869fa1) block (within [`api_rate_limit.api_endpoint_rules`](#api-rate-limit-api-endpoint-rules)) supports the following:

<a id="matchers-77386d"></a>&#x2022; [`cookie_matchers`](#matchers-77386d) - Optional Block<br>Cookie Matchers. A list of predicates for all cookies that need to be matched. The criteria for matching each cookie is described in individual instances of CookieMatcherType. The actual cookie values are extracted from the request API as a list of strings for each cookie name. Note that all specified cookie matcher predicates must evaluate to true<br>See [Cookie Matchers](#matchers-77386d) below.

<a id="headers-4b4f60"></a>&#x2022; [`headers`](#headers-4b4f60) - Optional Block<br>HTTP Headers. A list of predicates for various HTTP headers that need to match. The criteria for matching each HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values are extracted from the request API as a list of strings for each HTTP header type. Note that all specified header predicates must evaluate to true<br>See [Headers](#headers-4b4f60) below.

<a id="claims-eecbd4"></a>&#x2022; [`jwt_claims`](#claims-eecbd4) - Optional Block<br>JWT Claims. A list of predicates for various JWT claims that need to match. The criteria for matching each JWT claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are extracted from the JWT payload as a list of strings. Note that all specified JWT claim predicates must evaluate to true. Note that this feature only works on LBs with JWT Validation feature enabled<br>See [JWT Claims](#claims-eecbd4) below.

<a id="params-153b4d"></a>&#x2022; [`query_params`](#params-153b4d) - Optional Block<br>HTTP Query Parameters. A list of predicates for all query parameters that need to be matched. The criteria for matching each query parameter are described in individual instances of QueryParameterMatcherType. The actual query parameter values are extracted from the request API as a list of strings for each query parameter name. Note that all specified query parameter predicates must evaluate to true<br>See [Query Params](#params-153b4d) below.

#### API Rate Limit API Endpoint Rules Request Matcher Cookie Matchers

A [`cookie_matchers`](#matchers-77386d) block (within [`api_rate_limit.api_endpoint_rules.request_matcher`](#matcher-869fa1)) supports the following:

<a id="present-50637d"></a>&#x2022; [`check_not_present`](#present-50637d) - Optional Block<br>Enable this option

<a id="present-334dd3"></a>&#x2022; [`check_present`](#present-334dd3) - Optional Block<br>Enable this option

<a id="matcher-3e072a"></a>&#x2022; [`invert_matcher`](#matcher-3e072a) - Optional Bool<br>Invert Matcher. Invert Match of the expression defined

<a id="item-8fd79b"></a>&#x2022; [`item`](#item-8fd79b) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#item-8fd79b) below.

<a id="name-0761c7"></a>&#x2022; [`name`](#name-0761c7) - Optional String<br>Cookie Name. A case-sensitive cookie name

#### API Rate Limit API Endpoint Rules Request Matcher Cookie Matchers Item

An [`item`](#item-8fd79b) block (within [`api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers`](#matchers-77386d)) supports the following:

<a id="values-da29fa"></a>&#x2022; [`exact_values`](#values-da29fa) - Optional List<br>Exact Values. A list of exact values to match the input against

<a id="values-cd25b5"></a>&#x2022; [`regex_values`](#values-cd25b5) - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="transformers-8d7fe4"></a>&#x2022; [`transformers`](#transformers-8d7fe4) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>[Enum: LOWER_CASE|UPPER_CASE|BASE64_DECODE|NORMALIZE_PATH|REMOVE_WHITESPACE|URL_DECODE|TRIM_LEFT|TRIM_RIGHT|TRIM] Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### API Rate Limit API Endpoint Rules Request Matcher Headers

A [`headers`](#headers-4b4f60) block (within [`api_rate_limit.api_endpoint_rules.request_matcher`](#matcher-869fa1)) supports the following:

<a id="present-4ee0a5"></a>&#x2022; [`check_not_present`](#present-4ee0a5) - Optional Block<br>Enable this option

<a id="present-a33cbd"></a>&#x2022; [`check_present`](#present-a33cbd) - Optional Block<br>Enable this option

<a id="matcher-633667"></a>&#x2022; [`invert_matcher`](#matcher-633667) - Optional Bool<br>Invert Header Matcher. Invert the match result

<a id="item-ad545a"></a>&#x2022; [`item`](#item-ad545a) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#item-ad545a) below.

<a id="name-9984d9"></a>&#x2022; [`name`](#name-9984d9) - Optional String<br>Header Name. A case-insensitive HTTP header name

#### API Rate Limit API Endpoint Rules Request Matcher Headers Item

An [`item`](#item-ad545a) block (within [`api_rate_limit.api_endpoint_rules.request_matcher.headers`](#headers-4b4f60)) supports the following:

<a id="values-637456"></a>&#x2022; [`exact_values`](#values-637456) - Optional List<br>Exact Values. A list of exact values to match the input against

<a id="values-2c77df"></a>&#x2022; [`regex_values`](#values-2c77df) - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="transformers-c3442a"></a>&#x2022; [`transformers`](#transformers-c3442a) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>[Enum: LOWER_CASE|UPPER_CASE|BASE64_DECODE|NORMALIZE_PATH|REMOVE_WHITESPACE|URL_DECODE|TRIM_LEFT|TRIM_RIGHT|TRIM] Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### API Rate Limit API Endpoint Rules Request Matcher JWT Claims

A [`jwt_claims`](#claims-eecbd4) block (within [`api_rate_limit.api_endpoint_rules.request_matcher`](#matcher-869fa1)) supports the following:

<a id="present-935ed5"></a>&#x2022; [`check_not_present`](#present-935ed5) - Optional Block<br>Enable this option

<a id="present-fcd929"></a>&#x2022; [`check_present`](#present-fcd929) - Optional Block<br>Enable this option

<a id="matcher-cd173e"></a>&#x2022; [`invert_matcher`](#matcher-cd173e) - Optional Bool<br>Invert Matcher. Invert the match result

<a id="item-69e131"></a>&#x2022; [`item`](#item-69e131) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#item-69e131) below.

<a id="name-b37439"></a>&#x2022; [`name`](#name-b37439) - Optional String<br>JWT Claim Name. JWT claim name

#### API Rate Limit API Endpoint Rules Request Matcher JWT Claims Item

An [`item`](#item-69e131) block (within [`api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims`](#claims-eecbd4)) supports the following:

<a id="values-959e73"></a>&#x2022; [`exact_values`](#values-959e73) - Optional List<br>Exact Values. A list of exact values to match the input against

<a id="values-e5f104"></a>&#x2022; [`regex_values`](#values-e5f104) - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="transformers-cedd59"></a>&#x2022; [`transformers`](#transformers-cedd59) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>[Enum: LOWER_CASE|UPPER_CASE|BASE64_DECODE|NORMALIZE_PATH|REMOVE_WHITESPACE|URL_DECODE|TRIM_LEFT|TRIM_RIGHT|TRIM] Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### API Rate Limit API Endpoint Rules Request Matcher Query Params

A [`query_params`](#params-153b4d) block (within [`api_rate_limit.api_endpoint_rules.request_matcher`](#matcher-869fa1)) supports the following:

<a id="present-6edbea"></a>&#x2022; [`check_not_present`](#present-6edbea) - Optional Block<br>Enable this option

<a id="present-1cdabc"></a>&#x2022; [`check_present`](#present-1cdabc) - Optional Block<br>Enable this option

<a id="matcher-c07a30"></a>&#x2022; [`invert_matcher`](#matcher-c07a30) - Optional Bool<br>Invert Query Parameter Matcher. Invert the match result

<a id="item-b0554c"></a>&#x2022; [`item`](#item-b0554c) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#item-b0554c) below.

<a id="key-c4c42f"></a>&#x2022; [`key`](#key-c4c42f) - Optional String<br>Query Parameter Name. A case-sensitive HTTP query parameter name

#### API Rate Limit API Endpoint Rules Request Matcher Query Params Item

An [`item`](#item-b0554c) block (within [`api_rate_limit.api_endpoint_rules.request_matcher.query_params`](#params-153b4d)) supports the following:

<a id="values-36f490"></a>&#x2022; [`exact_values`](#values-36f490) - Optional List<br>Exact Values. A list of exact values to match the input against

<a id="values-6e357a"></a>&#x2022; [`regex_values`](#values-6e357a) - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="transformers-03d03b"></a>&#x2022; [`transformers`](#transformers-03d03b) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>[Enum: LOWER_CASE|UPPER_CASE|BASE64_DECODE|NORMALIZE_PATH|REMOVE_WHITESPACE|URL_DECODE|TRIM_LEFT|TRIM_RIGHT|TRIM] Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### API Rate Limit Bypass Rate Limiting Rules

A [`bypass_rate_limiting_rules`](#rules-776e97) block (within [`api_rate_limit`](#api-rate-limit)) supports the following:

<a id="rules-51aa34"></a>&#x2022; [`bypass_rate_limiting_rules`](#rules-51aa34) - Optional Block<br>Bypass Rate Limiting. This category defines rules per URL or API group. If request matches any of these rules, skip Rate Limiting<br>See [Bypass Rate Limiting Rules](#rules-51aa34) below.

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules

A [`bypass_rate_limiting_rules`](#rules-51aa34) block (within [`api_rate_limit.bypass_rate_limiting_rules`](#rules-776e97)) supports the following:

<a id="domain-0985ea"></a>&#x2022; [`any_domain`](#domain-0985ea) - Optional Block<br>Enable this option

<a id="url-7b53df"></a>&#x2022; [`any_url`](#url-7b53df) - Optional Block<br>Enable this option

<a id="endpoint-e28aa4"></a>&#x2022; [`api_endpoint`](#endpoint-e28aa4) - Optional Block<br>API Endpoint. This defines API endpoint<br>See [API Endpoint](#endpoint-e28aa4) below.

<a id="groups-c92822"></a>&#x2022; [`api_groups`](#groups-c92822) - Optional Block<br>API Groups<br>See [API Groups](#groups-c92822) below.

<a id="path-b16510"></a>&#x2022; [`base_path`](#path-b16510) - Optional String<br>Base Path. The base path which this validation applies to

<a id="matcher-a9da18"></a>&#x2022; [`client_matcher`](#matcher-a9da18) - Optional Block<br>Client Matcher. Client conditions for matching a rule<br>See [Client Matcher](#matcher-a9da18) below.

<a id="matcher-e9bb4d"></a>&#x2022; [`request_matcher`](#matcher-e9bb4d) - Optional Block<br>Request Matcher. Request conditions for matching a rule<br>See [Request Matcher](#matcher-e9bb4d) below.

<a id="domain-451df1"></a>&#x2022; [`specific_domain`](#domain-451df1) - Optional String<br>Specific Domain. The rule will apply for a specific domain. For example: API.example.com

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules API Endpoint

An [`api_endpoint`](#endpoint-e28aa4) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules`](#rules-51aa34)) supports the following:

<a id="methods-2f7610"></a>&#x2022; [`methods`](#methods-2f7610) - Optional List  Defaults to `ANY`<br>Possible values are `ANY`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`, `COPY`<br>[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY] Methods. Methods to be matched

<a id="path-79e5a9"></a>&#x2022; [`path`](#path-79e5a9) - Optional String<br>Path. Path to be matched

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules API Groups

An [`api_groups`](#groups-c92822) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules`](#rules-51aa34)) supports the following:

<a id="groups-56ebad"></a>&#x2022; [`api_groups`](#groups-56ebad) - Optional List<br>API Groups

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Client Matcher

A [`client_matcher`](#matcher-a9da18) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules`](#rules-51aa34)) supports the following:

<a id="client-10be2f"></a>&#x2022; [`any_client`](#client-10be2f) - Optional Block<br>Enable this option

<a id="any-ip-7c4970"></a>&#x2022; [`any_ip`](#any-ip-7c4970) - Optional Block<br>Enable this option

<a id="list-221e4b"></a>&#x2022; [`asn_list`](#list-221e4b) - Optional Block<br>ASN Match List. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer<br>See [Asn List](#list-221e4b) below.

<a id="matcher-d64a47"></a>&#x2022; [`asn_matcher`](#matcher-d64a47) - Optional Block<br>ASN Matcher. Match any AS number contained in the list of bgp_asn_sets<br>See [Asn Matcher](#matcher-d64a47) below.

<a id="selector-8bcea5"></a>&#x2022; [`client_selector`](#selector-8bcea5) - Optional Block<br>Label Selector. This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expresssions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects. Label selector is immutable. expressions is a list of strings of label selection expression. Each string has ',' separated values which are 'AND' and all strings are logically 'OR'. BNF for expression string `<selector-syntax>` ::= `<requirement>` | `<requirement>` ',' `<selector-syntax>` `<requirement>` ::= [!] KEY [ `<set-based-restriction>` | `<exact-match-restriction>` ] `<set-based-restriction>` ::= '' | `<inclusion-exclusion>` `<value-set>` `<inclusion-exclusion>` ::= `<inclusion>` | `<exclusion>` `<exclusion>` ::= 'notin' `<inclusion>` ::= 'in' `<value-set>` ::= '(' `<values>` ')' `<values>` ::= VALUE | VALUE ',' `<values>` `<exact-match-restriction>` ::= ['='|'=='|'!='] VALUE<br>See [Client Selector](#selector-8bcea5) below.

<a id="matcher-273263"></a>&#x2022; [`ip_matcher`](#matcher-273263) - Optional Block<br>IP Prefix Matcher. Match any IP prefix contained in the list of ip_prefix_sets. The result of the match is inverted if invert_matcher is true<br>See [IP Matcher](#matcher-273263) below.

<a id="list-3ef91d"></a>&#x2022; [`ip_prefix_list`](#list-3ef91d) - Optional Block<br>IP Prefix Match List. List of IP Prefix strings to match against<br>See [IP Prefix List](#list-3ef91d) below.

<a id="list-94743d"></a>&#x2022; [`ip_threat_category_list`](#list-94743d) - Optional Block<br>IP Threat Category List Type. List of IP threat categories<br>See [IP Threat Category List](#list-94743d) below.

<a id="matcher-c87ce2"></a>&#x2022; [`tls_fingerprint_matcher`](#matcher-c87ce2) - Optional Block<br>TLS Fingerprint Matcher. A TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of supported positve match criteria includes a list of known classes of TLS fingerprints and a list of exact values. The match is considered successful if either of these positive criteria are satisfied and the input fingerprint is not one of the excluded values<br>See [TLS Fingerprint Matcher](#matcher-c87ce2) below.

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Client Matcher Asn List

An [`asn_list`](#list-221e4b) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher`](#matcher-a9da18)) supports the following:

<a id="numbers-9c8ce1"></a>&#x2022; [`as_numbers`](#numbers-9c8ce1) - Optional List<br>AS Numbers. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Client Matcher Asn Matcher

An [`asn_matcher`](#matcher-d64a47) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher`](#matcher-a9da18)) supports the following:

<a id="sets-489c65"></a>&#x2022; [`asn_sets`](#sets-489c65) - Optional Block<br>BGP ASN Sets. A list of references to bgp_asn_set objects<br>See [Asn Sets](#sets-489c65) below.

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Client Matcher Asn Matcher Asn Sets

An [`asn_sets`](#sets-489c65) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher`](#matcher-d64a47)) supports the following:

<a id="kind-cf32d0"></a>&#x2022; [`kind`](#kind-cf32d0) - Optional String<br>Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route')

<a id="name-3a9c74"></a>&#x2022; [`name`](#name-3a9c74) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-d9cfc4"></a>&#x2022; [`namespace`](#namespace-d9cfc4) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-c5db47"></a>&#x2022; [`tenant`](#tenant-c5db47) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="uid-d6cdd9"></a>&#x2022; [`uid`](#uid-d6cdd9) - Optional String<br>UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Client Matcher Client Selector

A [`client_selector`](#selector-8bcea5) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher`](#matcher-a9da18)) supports the following:

<a id="expressions-e48729"></a>&#x2022; [`expressions`](#expressions-e48729) - Optional List<br>Selector Expression. expressions contains the kubernetes style label expression for selections

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Client Matcher IP Matcher

An [`ip_matcher`](#matcher-273263) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher`](#matcher-a9da18)) supports the following:

<a id="matcher-939f7d"></a>&#x2022; [`invert_matcher`](#matcher-939f7d) - Optional Bool<br>Invert IP Matcher. Invert the match result

<a id="sets-377781"></a>&#x2022; [`prefix_sets`](#sets-377781) - Optional Block<br>IP Prefix Sets. A list of references to ip_prefix_set objects<br>See [Prefix Sets](#sets-377781) below.

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Client Matcher IP Matcher Prefix Sets

A [`prefix_sets`](#sets-377781) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher`](#matcher-273263)) supports the following:

<a id="kind-566b63"></a>&#x2022; [`kind`](#kind-566b63) - Optional String<br>Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route')

<a id="name-80111f"></a>&#x2022; [`name`](#name-80111f) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-39558d"></a>&#x2022; [`namespace`](#namespace-39558d) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-80e0d0"></a>&#x2022; [`tenant`](#tenant-80e0d0) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="uid-232063"></a>&#x2022; [`uid`](#uid-232063) - Optional String<br>UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Client Matcher IP Prefix List

An [`ip_prefix_list`](#list-3ef91d) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher`](#matcher-a9da18)) supports the following:

<a id="match-710e3e"></a>&#x2022; [`invert_match`](#match-710e3e) - Optional Bool<br>Invert Match Result. Invert the match result

<a id="prefixes-fe1028"></a>&#x2022; [`ip_prefixes`](#prefixes-fe1028) - Optional List<br>IPv4 Prefix List. List of IPv4 prefix strings

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Client Matcher IP Threat Category List

An [`ip_threat_category_list`](#list-94743d) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher`](#matcher-a9da18)) supports the following:

<a id="categories-abf09a"></a>&#x2022; [`ip_threat_categories`](#categories-abf09a) - Optional List  Defaults to `SPAM_SOURCES`<br>Possible values are `SPAM_SOURCES`, `WINDOWS_EXPLOITS`, `WEB_ATTACKS`, `BOTNETS`, `SCANNERS`, `REPUTATION`, `PHISHING`, `PROXY`, `MOBILE_THREATS`, `TOR_PROXY`, `DENIAL_OF_SERVICE`, `NETWORK`<br>[Enum: SPAM_SOURCES|WINDOWS_EXPLOITS|WEB_ATTACKS|BOTNETS|SCANNERS|REPUTATION|PHISHING|PROXY|MOBILE_THREATS|TOR_PROXY|DENIAL_OF_SERVICE|NETWORK] List of IP Threat Categories to choose. The IP threat categories is obtained from the list and is used to auto-generate equivalent label selection expressions

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Client Matcher TLS Fingerprint Matcher

A [`tls_fingerprint_matcher`](#matcher-c87ce2) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher`](#matcher-a9da18)) supports the following:

<a id="classes-5a5e36"></a>&#x2022; [`classes`](#classes-5a5e36) - Optional List  Defaults to `TLS_FINGERPRINT_NONE`<br>Possible values are `TLS_FINGERPRINT_NONE`, `ANY_MALICIOUS_FINGERPRINT`, `ADWARE`, `ADWIND`, `DRIDEX`, `GOOTKIT`, `GOZI`, `JBIFROST`, `QUAKBOT`, `RANSOMWARE`, `TROLDESH`, `TOFSEE`, `TORRENTLOCKER`, `TRICKBOT`<br>[Enum: TLS_FINGERPRINT_NONE|ANY_MALICIOUS_FINGERPRINT|ADWARE|ADWIND|DRIDEX|GOOTKIT|GOZI|JBIFROST|QUAKBOT|RANSOMWARE|TROLDESH|TOFSEE|TORRENTLOCKER|TRICKBOT] TLS fingerprint classes. A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against

<a id="values-918a08"></a>&#x2022; [`exact_values`](#values-918a08) - Optional List<br>Exact Values. A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against

<a id="values-a7d1eb"></a>&#x2022; [`excluded_values`](#values-a7d1eb) - Optional List<br>Excluded Values. A list of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can be used to skip known false positives when using one or more known TLS fingerprint classes in the enclosing matcher

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Request Matcher

A [`request_matcher`](#matcher-e9bb4d) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules`](#rules-51aa34)) supports the following:

<a id="matchers-f34f88"></a>&#x2022; [`cookie_matchers`](#matchers-f34f88) - Optional Block<br>Cookie Matchers. A list of predicates for all cookies that need to be matched. The criteria for matching each cookie is described in individual instances of CookieMatcherType. The actual cookie values are extracted from the request API as a list of strings for each cookie name. Note that all specified cookie matcher predicates must evaluate to true<br>See [Cookie Matchers](#matchers-f34f88) below.

<a id="headers-4161e4"></a>&#x2022; [`headers`](#headers-4161e4) - Optional Block<br>HTTP Headers. A list of predicates for various HTTP headers that need to match. The criteria for matching each HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values are extracted from the request API as a list of strings for each HTTP header type. Note that all specified header predicates must evaluate to true<br>See [Headers](#headers-4161e4) below.

<a id="claims-13ffa7"></a>&#x2022; [`jwt_claims`](#claims-13ffa7) - Optional Block<br>JWT Claims. A list of predicates for various JWT claims that need to match. The criteria for matching each JWT claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are extracted from the JWT payload as a list of strings. Note that all specified JWT claim predicates must evaluate to true. Note that this feature only works on LBs with JWT Validation feature enabled<br>See [JWT Claims](#claims-13ffa7) below.

<a id="params-bfd454"></a>&#x2022; [`query_params`](#params-bfd454) - Optional Block<br>HTTP Query Parameters. A list of predicates for all query parameters that need to be matched. The criteria for matching each query parameter are described in individual instances of QueryParameterMatcherType. The actual query parameter values are extracted from the request API as a list of strings for each query parameter name. Note that all specified query parameter predicates must evaluate to true<br>See [Query Params](#params-bfd454) below.

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Request Matcher Cookie Matchers

A [`cookie_matchers`](#matchers-f34f88) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher`](#matcher-e9bb4d)) supports the following:

<a id="present-25beab"></a>&#x2022; [`check_not_present`](#present-25beab) - Optional Block<br>Enable this option

<a id="present-c78615"></a>&#x2022; [`check_present`](#present-c78615) - Optional Block<br>Enable this option

<a id="matcher-e0ea9b"></a>&#x2022; [`invert_matcher`](#matcher-e0ea9b) - Optional Bool<br>Invert Matcher. Invert Match of the expression defined

<a id="item-69d8e4"></a>&#x2022; [`item`](#item-69d8e4) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#item-69d8e4) below.

<a id="name-2d8fee"></a>&#x2022; [`name`](#name-2d8fee) - Optional String<br>Cookie Name. A case-sensitive cookie name

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Request Matcher Cookie Matchers Item

An [`item`](#item-69d8e4) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers`](#matchers-f34f88)) supports the following:

<a id="values-9bae32"></a>&#x2022; [`exact_values`](#values-9bae32) - Optional List<br>Exact Values. A list of exact values to match the input against

<a id="values-9362bb"></a>&#x2022; [`regex_values`](#values-9362bb) - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="transformers-a8c5ff"></a>&#x2022; [`transformers`](#transformers-a8c5ff) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>[Enum: LOWER_CASE|UPPER_CASE|BASE64_DECODE|NORMALIZE_PATH|REMOVE_WHITESPACE|URL_DECODE|TRIM_LEFT|TRIM_RIGHT|TRIM] Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Request Matcher Headers

A [`headers`](#headers-4161e4) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher`](#matcher-e9bb4d)) supports the following:

<a id="present-ae8198"></a>&#x2022; [`check_not_present`](#present-ae8198) - Optional Block<br>Enable this option

<a id="present-2f4647"></a>&#x2022; [`check_present`](#present-2f4647) - Optional Block<br>Enable this option

<a id="matcher-65dc3a"></a>&#x2022; [`invert_matcher`](#matcher-65dc3a) - Optional Bool<br>Invert Header Matcher. Invert the match result

<a id="item-4706a3"></a>&#x2022; [`item`](#item-4706a3) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#item-4706a3) below.

<a id="name-b0a7d3"></a>&#x2022; [`name`](#name-b0a7d3) - Optional String<br>Header Name. A case-insensitive HTTP header name

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Request Matcher Headers Item

An [`item`](#item-4706a3) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers`](#headers-4161e4)) supports the following:

<a id="values-3baf40"></a>&#x2022; [`exact_values`](#values-3baf40) - Optional List<br>Exact Values. A list of exact values to match the input against

<a id="values-ce8cc2"></a>&#x2022; [`regex_values`](#values-ce8cc2) - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="transformers-e84879"></a>&#x2022; [`transformers`](#transformers-e84879) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>[Enum: LOWER_CASE|UPPER_CASE|BASE64_DECODE|NORMALIZE_PATH|REMOVE_WHITESPACE|URL_DECODE|TRIM_LEFT|TRIM_RIGHT|TRIM] Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Request Matcher JWT Claims

A [`jwt_claims`](#claims-13ffa7) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher`](#matcher-e9bb4d)) supports the following:

<a id="present-0e108d"></a>&#x2022; [`check_not_present`](#present-0e108d) - Optional Block<br>Enable this option

<a id="present-0ac17d"></a>&#x2022; [`check_present`](#present-0ac17d) - Optional Block<br>Enable this option

<a id="matcher-43cde9"></a>&#x2022; [`invert_matcher`](#matcher-43cde9) - Optional Bool<br>Invert Matcher. Invert the match result

<a id="item-ca8646"></a>&#x2022; [`item`](#item-ca8646) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#item-ca8646) below.

<a id="name-3618ff"></a>&#x2022; [`name`](#name-3618ff) - Optional String<br>JWT Claim Name. JWT claim name

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Request Matcher JWT Claims Item

An [`item`](#item-ca8646) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims`](#claims-13ffa7)) supports the following:

<a id="values-1ac86b"></a>&#x2022; [`exact_values`](#values-1ac86b) - Optional List<br>Exact Values. A list of exact values to match the input against

<a id="values-44ceff"></a>&#x2022; [`regex_values`](#values-44ceff) - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="transformers-c7f562"></a>&#x2022; [`transformers`](#transformers-c7f562) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>[Enum: LOWER_CASE|UPPER_CASE|BASE64_DECODE|NORMALIZE_PATH|REMOVE_WHITESPACE|URL_DECODE|TRIM_LEFT|TRIM_RIGHT|TRIM] Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Request Matcher Query Params

A [`query_params`](#params-bfd454) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher`](#matcher-e9bb4d)) supports the following:

<a id="present-928f2f"></a>&#x2022; [`check_not_present`](#present-928f2f) - Optional Block<br>Enable this option

<a id="present-0997e1"></a>&#x2022; [`check_present`](#present-0997e1) - Optional Block<br>Enable this option

<a id="matcher-94883f"></a>&#x2022; [`invert_matcher`](#matcher-94883f) - Optional Bool<br>Invert Query Parameter Matcher. Invert the match result

<a id="item-1cc059"></a>&#x2022; [`item`](#item-1cc059) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#item-1cc059) below.

<a id="key-4059d5"></a>&#x2022; [`key`](#key-4059d5) - Optional String<br>Query Parameter Name. A case-sensitive HTTP query parameter name

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Request Matcher Query Params Item

An [`item`](#item-1cc059) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params`](#params-bfd454)) supports the following:

<a id="values-1b4d75"></a>&#x2022; [`exact_values`](#values-1b4d75) - Optional List<br>Exact Values. A list of exact values to match the input against

<a id="values-76def7"></a>&#x2022; [`regex_values`](#values-76def7) - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="transformers-9a24cd"></a>&#x2022; [`transformers`](#transformers-9a24cd) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>[Enum: LOWER_CASE|UPPER_CASE|BASE64_DECODE|NORMALIZE_PATH|REMOVE_WHITESPACE|URL_DECODE|TRIM_LEFT|TRIM_RIGHT|TRIM] Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### API Rate Limit Custom IP Allowed List

A [`custom_ip_allowed_list`](#api-rate-limit-custom-ip-allowed-list) block (within [`api_rate_limit`](#api-rate-limit)) supports the following:

<a id="prefixes-73df46"></a>&#x2022; [`rate_limiter_allowed_prefixes`](#prefixes-73df46) - Optional Block<br>List of IP Prefix Sets. References to ip_prefix_set objects. Requests from source IP addresses that are covered by one of the allowed IP Prefixes are not subjected to rate limiting<br>See [Rate Limiter Allowed Prefixes](#prefixes-73df46) below.

#### API Rate Limit Custom IP Allowed List Rate Limiter Allowed Prefixes

A [`rate_limiter_allowed_prefixes`](#prefixes-73df46) block (within [`api_rate_limit.custom_ip_allowed_list`](#api-rate-limit-custom-ip-allowed-list)) supports the following:

<a id="name-3a08ca"></a>&#x2022; [`name`](#name-3a08ca) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-8714e1"></a>&#x2022; [`namespace`](#namespace-8714e1) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-89acd3"></a>&#x2022; [`tenant`](#tenant-89acd3) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### API Rate Limit IP Allowed List

An [`ip_allowed_list`](#api-rate-limit-ip-allowed-list) block (within [`api_rate_limit`](#api-rate-limit)) supports the following:

<a id="api-rate-limit-ip-allowed-list-prefixes"></a>&#x2022; [`prefixes`](#api-rate-limit-ip-allowed-list-prefixes) - Optional List<br>IPv4 Prefix List. List of IPv4 prefixes that represent an endpoint

#### API Rate Limit Server URL Rules

A [`server_url_rules`](#api-rate-limit-server-url-rules) block (within [`api_rate_limit`](#api-rate-limit)) supports the following:

<a id="domain-0747c9"></a>&#x2022; [`any_domain`](#domain-0747c9) - Optional Block<br>Enable this option

<a id="group-15c11a"></a>&#x2022; [`api_group`](#group-15c11a) - Optional String<br>API Group. API groups derived from API Definition swaggers. For example oas-all-operations including all paths and methods from the swaggers, oas-base-urls covering all requests under base-paths from the swaggers. Custom groups can be created if user tags paths or operations with 'x-volterra-API-group' extensions inside swaggers

<a id="path-44dbff"></a>&#x2022; [`base_path`](#path-44dbff) - Optional String<br>Base Path. Prefix of the request path

<a id="matcher-ed4b34"></a>&#x2022; [`client_matcher`](#matcher-ed4b34) - Optional Block<br>Client Matcher. Client conditions for matching a rule<br>See [Client Matcher](#matcher-ed4b34) below.

<a id="limiter-9faa53"></a>&#x2022; [`inline_rate_limiter`](#limiter-9faa53) - Optional Block<br>InlineRateLimiter<br>See [Inline Rate Limiter](#limiter-9faa53) below.

<a id="limiter-383ca9"></a>&#x2022; [`ref_rate_limiter`](#limiter-383ca9) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Ref Rate Limiter](#limiter-383ca9) below.

<a id="matcher-d0eea8"></a>&#x2022; [`request_matcher`](#matcher-d0eea8) - Optional Block<br>Request Matcher. Request conditions for matching a rule<br>See [Request Matcher](#matcher-d0eea8) below.

<a id="domain-dca9c1"></a>&#x2022; [`specific_domain`](#domain-dca9c1) - Optional String<br>Specific Domain. The rule will apply for a specific domain

#### API Rate Limit Server URL Rules Client Matcher

A [`client_matcher`](#matcher-ed4b34) block (within [`api_rate_limit.server_url_rules`](#api-rate-limit-server-url-rules)) supports the following:

<a id="client-d95ee8"></a>&#x2022; [`any_client`](#client-d95ee8) - Optional Block<br>Enable this option

<a id="any-ip-e752c7"></a>&#x2022; [`any_ip`](#any-ip-e752c7) - Optional Block<br>Enable this option

<a id="list-52dae1"></a>&#x2022; [`asn_list`](#list-52dae1) - Optional Block<br>ASN Match List. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer<br>See [Asn List](#list-52dae1) below.

<a id="matcher-9643c3"></a>&#x2022; [`asn_matcher`](#matcher-9643c3) - Optional Block<br>ASN Matcher. Match any AS number contained in the list of bgp_asn_sets<br>See [Asn Matcher](#matcher-9643c3) below.

<a id="selector-75ec07"></a>&#x2022; [`client_selector`](#selector-75ec07) - Optional Block<br>Label Selector. This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expresssions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects. Label selector is immutable. expressions is a list of strings of label selection expression. Each string has ',' separated values which are 'AND' and all strings are logically 'OR'. BNF for expression string `<selector-syntax>` ::= `<requirement>` | `<requirement>` ',' `<selector-syntax>` `<requirement>` ::= [!] KEY [ `<set-based-restriction>` | `<exact-match-restriction>` ] `<set-based-restriction>` ::= '' | `<inclusion-exclusion>` `<value-set>` `<inclusion-exclusion>` ::= `<inclusion>` | `<exclusion>` `<exclusion>` ::= 'notin' `<inclusion>` ::= 'in' `<value-set>` ::= '(' `<values>` ')' `<values>` ::= VALUE | VALUE ',' `<values>` `<exact-match-restriction>` ::= ['='|'=='|'!='] VALUE<br>See [Client Selector](#selector-75ec07) below.

<a id="matcher-74485b"></a>&#x2022; [`ip_matcher`](#matcher-74485b) - Optional Block<br>IP Prefix Matcher. Match any IP prefix contained in the list of ip_prefix_sets. The result of the match is inverted if invert_matcher is true<br>See [IP Matcher](#matcher-74485b) below.

<a id="list-047465"></a>&#x2022; [`ip_prefix_list`](#list-047465) - Optional Block<br>IP Prefix Match List. List of IP Prefix strings to match against<br>See [IP Prefix List](#list-047465) below.

<a id="list-ac4c85"></a>&#x2022; [`ip_threat_category_list`](#list-ac4c85) - Optional Block<br>IP Threat Category List Type. List of IP threat categories<br>See [IP Threat Category List](#list-ac4c85) below.

<a id="matcher-d896f3"></a>&#x2022; [`tls_fingerprint_matcher`](#matcher-d896f3) - Optional Block<br>TLS Fingerprint Matcher. A TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of supported positve match criteria includes a list of known classes of TLS fingerprints and a list of exact values. The match is considered successful if either of these positive criteria are satisfied and the input fingerprint is not one of the excluded values<br>See [TLS Fingerprint Matcher](#matcher-d896f3) below.

#### API Rate Limit Server URL Rules Client Matcher Asn List

An [`asn_list`](#list-52dae1) block (within [`api_rate_limit.server_url_rules.client_matcher`](#matcher-ed4b34)) supports the following:

<a id="numbers-826050"></a>&#x2022; [`as_numbers`](#numbers-826050) - Optional List<br>AS Numbers. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer

#### API Rate Limit Server URL Rules Client Matcher Asn Matcher

An [`asn_matcher`](#matcher-9643c3) block (within [`api_rate_limit.server_url_rules.client_matcher`](#matcher-ed4b34)) supports the following:

<a id="sets-2c4adf"></a>&#x2022; [`asn_sets`](#sets-2c4adf) - Optional Block<br>BGP ASN Sets. A list of references to bgp_asn_set objects<br>See [Asn Sets](#sets-2c4adf) below.

#### API Rate Limit Server URL Rules Client Matcher Asn Matcher Asn Sets

An [`asn_sets`](#sets-2c4adf) block (within [`api_rate_limit.server_url_rules.client_matcher.asn_matcher`](#matcher-9643c3)) supports the following:

<a id="kind-901b29"></a>&#x2022; [`kind`](#kind-901b29) - Optional String<br>Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route')

<a id="name-b807d9"></a>&#x2022; [`name`](#name-b807d9) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-f14602"></a>&#x2022; [`namespace`](#namespace-f14602) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-d3fbe6"></a>&#x2022; [`tenant`](#tenant-d3fbe6) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="uid-e2fb5d"></a>&#x2022; [`uid`](#uid-e2fb5d) - Optional String<br>UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid

#### API Rate Limit Server URL Rules Client Matcher Client Selector

A [`client_selector`](#selector-75ec07) block (within [`api_rate_limit.server_url_rules.client_matcher`](#matcher-ed4b34)) supports the following:

<a id="expressions-6e94c4"></a>&#x2022; [`expressions`](#expressions-6e94c4) - Optional List<br>Selector Expression. expressions contains the kubernetes style label expression for selections

#### API Rate Limit Server URL Rules Client Matcher IP Matcher

An [`ip_matcher`](#matcher-74485b) block (within [`api_rate_limit.server_url_rules.client_matcher`](#matcher-ed4b34)) supports the following:

<a id="matcher-348dd9"></a>&#x2022; [`invert_matcher`](#matcher-348dd9) - Optional Bool<br>Invert IP Matcher. Invert the match result

<a id="sets-13bcb5"></a>&#x2022; [`prefix_sets`](#sets-13bcb5) - Optional Block<br>IP Prefix Sets. A list of references to ip_prefix_set objects<br>See [Prefix Sets](#sets-13bcb5) below.

#### API Rate Limit Server URL Rules Client Matcher IP Matcher Prefix Sets

A [`prefix_sets`](#sets-13bcb5) block (within [`api_rate_limit.server_url_rules.client_matcher.ip_matcher`](#matcher-74485b)) supports the following:

<a id="kind-6b5e2d"></a>&#x2022; [`kind`](#kind-6b5e2d) - Optional String<br>Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route')

<a id="name-f97673"></a>&#x2022; [`name`](#name-f97673) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-77b4df"></a>&#x2022; [`namespace`](#namespace-77b4df) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-2de84a"></a>&#x2022; [`tenant`](#tenant-2de84a) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="uid-b93ef7"></a>&#x2022; [`uid`](#uid-b93ef7) - Optional String<br>UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid

#### API Rate Limit Server URL Rules Client Matcher IP Prefix List

An [`ip_prefix_list`](#list-047465) block (within [`api_rate_limit.server_url_rules.client_matcher`](#matcher-ed4b34)) supports the following:

<a id="match-2b3904"></a>&#x2022; [`invert_match`](#match-2b3904) - Optional Bool<br>Invert Match Result. Invert the match result

<a id="prefixes-4fcef5"></a>&#x2022; [`ip_prefixes`](#prefixes-4fcef5) - Optional List<br>IPv4 Prefix List. List of IPv4 prefix strings

#### API Rate Limit Server URL Rules Client Matcher IP Threat Category List

An [`ip_threat_category_list`](#list-ac4c85) block (within [`api_rate_limit.server_url_rules.client_matcher`](#matcher-ed4b34)) supports the following:

<a id="categories-f2e19c"></a>&#x2022; [`ip_threat_categories`](#categories-f2e19c) - Optional List  Defaults to `SPAM_SOURCES`<br>Possible values are `SPAM_SOURCES`, `WINDOWS_EXPLOITS`, `WEB_ATTACKS`, `BOTNETS`, `SCANNERS`, `REPUTATION`, `PHISHING`, `PROXY`, `MOBILE_THREATS`, `TOR_PROXY`, `DENIAL_OF_SERVICE`, `NETWORK`<br>[Enum: SPAM_SOURCES|WINDOWS_EXPLOITS|WEB_ATTACKS|BOTNETS|SCANNERS|REPUTATION|PHISHING|PROXY|MOBILE_THREATS|TOR_PROXY|DENIAL_OF_SERVICE|NETWORK] List of IP Threat Categories to choose. The IP threat categories is obtained from the list and is used to auto-generate equivalent label selection expressions

#### API Rate Limit Server URL Rules Client Matcher TLS Fingerprint Matcher

A [`tls_fingerprint_matcher`](#matcher-d896f3) block (within [`api_rate_limit.server_url_rules.client_matcher`](#matcher-ed4b34)) supports the following:

<a id="classes-c89726"></a>&#x2022; [`classes`](#classes-c89726) - Optional List  Defaults to `TLS_FINGERPRINT_NONE`<br>Possible values are `TLS_FINGERPRINT_NONE`, `ANY_MALICIOUS_FINGERPRINT`, `ADWARE`, `ADWIND`, `DRIDEX`, `GOOTKIT`, `GOZI`, `JBIFROST`, `QUAKBOT`, `RANSOMWARE`, `TROLDESH`, `TOFSEE`, `TORRENTLOCKER`, `TRICKBOT`<br>[Enum: TLS_FINGERPRINT_NONE|ANY_MALICIOUS_FINGERPRINT|ADWARE|ADWIND|DRIDEX|GOOTKIT|GOZI|JBIFROST|QUAKBOT|RANSOMWARE|TROLDESH|TOFSEE|TORRENTLOCKER|TRICKBOT] TLS fingerprint classes. A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against

<a id="values-e9c2ed"></a>&#x2022; [`exact_values`](#values-e9c2ed) - Optional List<br>Exact Values. A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against

<a id="values-c878a8"></a>&#x2022; [`excluded_values`](#values-c878a8) - Optional List<br>Excluded Values. A list of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can be used to skip known false positives when using one or more known TLS fingerprint classes in the enclosing matcher

#### API Rate Limit Server URL Rules Inline Rate Limiter

An [`inline_rate_limiter`](#limiter-9faa53) block (within [`api_rate_limit.server_url_rules`](#api-rate-limit-server-url-rules)) supports the following:

<a id="user-id-2410cf"></a>&#x2022; [`ref_user_id`](#user-id-2410cf) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Ref User ID](#user-id-2410cf) below.

<a id="threshold-e1f6ce"></a>&#x2022; [`threshold`](#threshold-e1f6ce) - Optional Number<br>Threshold. The total number of allowed requests for 1 unit (e.g. SECOND/MINUTE/HOUR etc.) of the specified period

<a id="unit-23f142"></a>&#x2022; [`unit`](#unit-23f142) - Optional String  Defaults to `SECOND`<br>Possible values are `SECOND`, `MINUTE`, `HOUR`<br>[Enum: SECOND|MINUTE|HOUR] Rate Limit Period Unit. Unit for the period per which the rate limit is applied. - SECOND: Second Rate limit period unit is seconds - MINUTE: Minute Rate limit period unit is minutes - HOUR: Hour Rate limit period unit is hours - DAY: Day Rate limit period unit is days

<a id="user-id-560a62"></a>&#x2022; [`use_http_lb_user_id`](#user-id-560a62) - Optional Block<br>Enable this option

#### API Rate Limit Server URL Rules Inline Rate Limiter Ref User ID

A [`ref_user_id`](#user-id-2410cf) block (within [`api_rate_limit.server_url_rules.inline_rate_limiter`](#limiter-9faa53)) supports the following:

<a id="name-91d5b8"></a>&#x2022; [`name`](#name-91d5b8) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-9a5eae"></a>&#x2022; [`namespace`](#namespace-9a5eae) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-4c9142"></a>&#x2022; [`tenant`](#tenant-4c9142) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### API Rate Limit Server URL Rules Ref Rate Limiter

A [`ref_rate_limiter`](#limiter-383ca9) block (within [`api_rate_limit.server_url_rules`](#api-rate-limit-server-url-rules)) supports the following:

<a id="name-33d109"></a>&#x2022; [`name`](#name-33d109) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-bdf110"></a>&#x2022; [`namespace`](#namespace-bdf110) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-3c7e29"></a>&#x2022; [`tenant`](#tenant-3c7e29) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### API Rate Limit Server URL Rules Request Matcher

A [`request_matcher`](#matcher-d0eea8) block (within [`api_rate_limit.server_url_rules`](#api-rate-limit-server-url-rules)) supports the following:

<a id="matchers-834089"></a>&#x2022; [`cookie_matchers`](#matchers-834089) - Optional Block<br>Cookie Matchers. A list of predicates for all cookies that need to be matched. The criteria for matching each cookie is described in individual instances of CookieMatcherType. The actual cookie values are extracted from the request API as a list of strings for each cookie name. Note that all specified cookie matcher predicates must evaluate to true<br>See [Cookie Matchers](#matchers-834089) below.

<a id="headers-f3e5bd"></a>&#x2022; [`headers`](#headers-f3e5bd) - Optional Block<br>HTTP Headers. A list of predicates for various HTTP headers that need to match. The criteria for matching each HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values are extracted from the request API as a list of strings for each HTTP header type. Note that all specified header predicates must evaluate to true<br>See [Headers](#headers-f3e5bd) below.

<a id="claims-12c338"></a>&#x2022; [`jwt_claims`](#claims-12c338) - Optional Block<br>JWT Claims. A list of predicates for various JWT claims that need to match. The criteria for matching each JWT claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are extracted from the JWT payload as a list of strings. Note that all specified JWT claim predicates must evaluate to true. Note that this feature only works on LBs with JWT Validation feature enabled<br>See [JWT Claims](#claims-12c338) below.

<a id="params-176d8b"></a>&#x2022; [`query_params`](#params-176d8b) - Optional Block<br>HTTP Query Parameters. A list of predicates for all query parameters that need to be matched. The criteria for matching each query parameter are described in individual instances of QueryParameterMatcherType. The actual query parameter values are extracted from the request API as a list of strings for each query parameter name. Note that all specified query parameter predicates must evaluate to true<br>See [Query Params](#params-176d8b) below.

#### API Rate Limit Server URL Rules Request Matcher Cookie Matchers

A [`cookie_matchers`](#matchers-834089) block (within [`api_rate_limit.server_url_rules.request_matcher`](#matcher-d0eea8)) supports the following:

<a id="present-4a314d"></a>&#x2022; [`check_not_present`](#present-4a314d) - Optional Block<br>Enable this option

<a id="present-bfd192"></a>&#x2022; [`check_present`](#present-bfd192) - Optional Block<br>Enable this option

<a id="matcher-5a5e9f"></a>&#x2022; [`invert_matcher`](#matcher-5a5e9f) - Optional Bool<br>Invert Matcher. Invert Match of the expression defined

<a id="item-4904af"></a>&#x2022; [`item`](#item-4904af) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#item-4904af) below.

<a id="name-764850"></a>&#x2022; [`name`](#name-764850) - Optional String<br>Cookie Name. A case-sensitive cookie name

#### API Rate Limit Server URL Rules Request Matcher Cookie Matchers Item

An [`item`](#item-4904af) block (within [`api_rate_limit.server_url_rules.request_matcher.cookie_matchers`](#matchers-834089)) supports the following:

<a id="values-1a1200"></a>&#x2022; [`exact_values`](#values-1a1200) - Optional List<br>Exact Values. A list of exact values to match the input against

<a id="values-d6047d"></a>&#x2022; [`regex_values`](#values-d6047d) - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="transformers-c3e045"></a>&#x2022; [`transformers`](#transformers-c3e045) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>[Enum: LOWER_CASE|UPPER_CASE|BASE64_DECODE|NORMALIZE_PATH|REMOVE_WHITESPACE|URL_DECODE|TRIM_LEFT|TRIM_RIGHT|TRIM] Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### API Rate Limit Server URL Rules Request Matcher Headers

A [`headers`](#headers-f3e5bd) block (within [`api_rate_limit.server_url_rules.request_matcher`](#matcher-d0eea8)) supports the following:

<a id="present-b8a223"></a>&#x2022; [`check_not_present`](#present-b8a223) - Optional Block<br>Enable this option

<a id="present-de3982"></a>&#x2022; [`check_present`](#present-de3982) - Optional Block<br>Enable this option

<a id="matcher-7ed807"></a>&#x2022; [`invert_matcher`](#matcher-7ed807) - Optional Bool<br>Invert Header Matcher. Invert the match result

<a id="item-8635a7"></a>&#x2022; [`item`](#item-8635a7) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#item-8635a7) below.

<a id="name-7ac50b"></a>&#x2022; [`name`](#name-7ac50b) - Optional String<br>Header Name. A case-insensitive HTTP header name

#### API Rate Limit Server URL Rules Request Matcher Headers Item

An [`item`](#item-8635a7) block (within [`api_rate_limit.server_url_rules.request_matcher.headers`](#headers-f3e5bd)) supports the following:

<a id="values-bb1d87"></a>&#x2022; [`exact_values`](#values-bb1d87) - Optional List<br>Exact Values. A list of exact values to match the input against

<a id="values-99be35"></a>&#x2022; [`regex_values`](#values-99be35) - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="transformers-17e679"></a>&#x2022; [`transformers`](#transformers-17e679) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>[Enum: LOWER_CASE|UPPER_CASE|BASE64_DECODE|NORMALIZE_PATH|REMOVE_WHITESPACE|URL_DECODE|TRIM_LEFT|TRIM_RIGHT|TRIM] Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### API Rate Limit Server URL Rules Request Matcher JWT Claims

A [`jwt_claims`](#claims-12c338) block (within [`api_rate_limit.server_url_rules.request_matcher`](#matcher-d0eea8)) supports the following:

<a id="present-b2dfb7"></a>&#x2022; [`check_not_present`](#present-b2dfb7) - Optional Block<br>Enable this option

<a id="present-54a989"></a>&#x2022; [`check_present`](#present-54a989) - Optional Block<br>Enable this option

<a id="matcher-a26813"></a>&#x2022; [`invert_matcher`](#matcher-a26813) - Optional Bool<br>Invert Matcher. Invert the match result

<a id="item-b8fbc9"></a>&#x2022; [`item`](#item-b8fbc9) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#item-b8fbc9) below.

<a id="name-eb48e1"></a>&#x2022; [`name`](#name-eb48e1) - Optional String<br>JWT Claim Name. JWT claim name

#### API Rate Limit Server URL Rules Request Matcher JWT Claims Item

An [`item`](#item-b8fbc9) block (within [`api_rate_limit.server_url_rules.request_matcher.jwt_claims`](#claims-12c338)) supports the following:

<a id="values-db3fa8"></a>&#x2022; [`exact_values`](#values-db3fa8) - Optional List<br>Exact Values. A list of exact values to match the input against

<a id="values-477cf3"></a>&#x2022; [`regex_values`](#values-477cf3) - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="transformers-9a6b31"></a>&#x2022; [`transformers`](#transformers-9a6b31) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>[Enum: LOWER_CASE|UPPER_CASE|BASE64_DECODE|NORMALIZE_PATH|REMOVE_WHITESPACE|URL_DECODE|TRIM_LEFT|TRIM_RIGHT|TRIM] Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### API Rate Limit Server URL Rules Request Matcher Query Params

A [`query_params`](#params-176d8b) block (within [`api_rate_limit.server_url_rules.request_matcher`](#matcher-d0eea8)) supports the following:

<a id="present-3444e4"></a>&#x2022; [`check_not_present`](#present-3444e4) - Optional Block<br>Enable this option

<a id="present-e16059"></a>&#x2022; [`check_present`](#present-e16059) - Optional Block<br>Enable this option

<a id="matcher-1e387b"></a>&#x2022; [`invert_matcher`](#matcher-1e387b) - Optional Bool<br>Invert Query Parameter Matcher. Invert the match result

<a id="item-b75d42"></a>&#x2022; [`item`](#item-b75d42) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#item-b75d42) below.

<a id="key-049198"></a>&#x2022; [`key`](#key-049198) - Optional String<br>Query Parameter Name. A case-sensitive HTTP query parameter name

#### API Rate Limit Server URL Rules Request Matcher Query Params Item

An [`item`](#item-b75d42) block (within [`api_rate_limit.server_url_rules.request_matcher.query_params`](#params-176d8b)) supports the following:

<a id="values-1feeac"></a>&#x2022; [`exact_values`](#values-1feeac) - Optional List<br>Exact Values. A list of exact values to match the input against

<a id="values-2381e5"></a>&#x2022; [`regex_values`](#values-2381e5) - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="transformers-8b8f7c"></a>&#x2022; [`transformers`](#transformers-8b8f7c) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>[Enum: LOWER_CASE|UPPER_CASE|BASE64_DECODE|NORMALIZE_PATH|REMOVE_WHITESPACE|URL_DECODE|TRIM_LEFT|TRIM_RIGHT|TRIM] Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### API Specification

An [`api_specification`](#api-specification) block supports the following:

<a id="api-specification-api-definition"></a>&#x2022; [`api_definition`](#api-specification-api-definition) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [API Definition](#api-specification-api-definition) below.

<a id="endpoints-4158a4"></a>&#x2022; [`validation_all_spec_endpoints`](#endpoints-4158a4) - Optional Block<br>API Inventory. Settings for API Inventory validation<br>See [Validation All Spec Endpoints](#endpoints-4158a4) below.

<a id="list-23b577"></a>&#x2022; [`validation_custom_list`](#list-23b577) - Optional Block<br>Custom List. Define API groups, base paths, or API endpoints and their OpenAPI validation modes. Any other API-endpoint not listed will act according to 'Fall Through Mode'<br>See [Validation Custom List](#list-23b577) below.

<a id="api-specification-validation-disabled"></a>&#x2022; [`validation_disabled`](#api-specification-validation-disabled) - Optional Block<br>Enable this option

#### API Specification API Definition

An [`api_definition`](#api-specification-api-definition) block (within [`api_specification`](#api-specification)) supports the following:

<a id="api-specification-api-definition-name"></a>&#x2022; [`name`](#api-specification-api-definition-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-c685bf"></a>&#x2022; [`namespace`](#namespace-c685bf) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="api-specification-api-definition-tenant"></a>&#x2022; [`tenant`](#api-specification-api-definition-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### API Specification Validation All Spec Endpoints

A [`validation_all_spec_endpoints`](#endpoints-4158a4) block (within [`api_specification`](#api-specification)) supports the following:

<a id="mode-8425c5"></a>&#x2022; [`fall_through_mode`](#mode-8425c5) - Optional Block<br>Fall Through Mode.Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a. swagger) or doesn't have a specific rule in custom rules)<br>See [Fall Through Mode](#mode-8425c5) below.

<a id="settings-a83a93"></a>&#x2022; [`settings`](#settings-a83a93) - Optional Block<br>Common Settings. OpenAPI specification validation settings relevant for 'API Inventory' enforcement and for 'Custom list' enforcement<br>See [Settings](#settings-a83a93) below.

<a id="mode-cd4a1c"></a>&#x2022; [`validation_mode`](#mode-cd4a1c) - Optional Block<br>Validation Mode.Validation mode of OpenAPI specification. When a validation mismatch occurs on a request to one of the endpoints listed on the OpenAPI specification file (a.k.a. swagger)<br>See [Validation Mode](#mode-cd4a1c) below.

#### API Specification Validation All Spec Endpoints Fall Through Mode

A [`fall_through_mode`](#mode-8425c5) block (within [`api_specification.validation_all_spec_endpoints`](#endpoints-4158a4)) supports the following:

<a id="allow-fe1e6a"></a>&#x2022; [`fall_through_mode_allow`](#allow-fe1e6a) - Optional Block<br>Enable this option

<a id="custom-aadcaa"></a>&#x2022; [`fall_through_mode_custom`](#custom-aadcaa) - Optional Block<br>Custom Fall Through Mode. Define the fall through settings<br>See [Fall Through Mode Custom](#custom-aadcaa) below.

#### API Specification Validation All Spec Endpoints Fall Through Mode Fall Through Mode Custom

A [`fall_through_mode_custom`](#custom-aadcaa) block (within [`api_specification.validation_all_spec_endpoints.fall_through_mode`](#mode-8425c5)) supports the following:

<a id="rules-7e1bb3"></a>&#x2022; [`open_api_validation_rules`](#rules-7e1bb3) - Optional Block<br>Custom Fall Through Rule List<br>See [Open API Validation Rules](#rules-7e1bb3) below.

#### API Specification Validation All Spec Endpoints Fall Through Mode Fall Through Mode Custom Open API Validation Rules

An [`open_api_validation_rules`](#rules-7e1bb3) block (within [`api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom`](#custom-aadcaa)) supports the following:

<a id="block-392345"></a>&#x2022; [`action_block`](#block-392345) - Optional Block<br>Enable this option

<a id="report-70f264"></a>&#x2022; [`action_report`](#report-70f264) - Optional Block<br>Enable this option

<a id="skip-5ad739"></a>&#x2022; [`action_skip`](#skip-5ad739) - Optional Block<br>Enable this option

<a id="endpoint-eb6d3c"></a>&#x2022; [`api_endpoint`](#endpoint-eb6d3c) - Optional Block<br>API Endpoint. This defines API endpoint<br>See [API Endpoint](#endpoint-eb6d3c) below.

<a id="group-ba04ab"></a>&#x2022; [`api_group`](#group-ba04ab) - Optional String<br>API Group. The API group which this validation applies to

<a id="path-822099"></a>&#x2022; [`base_path`](#path-822099) - Optional String<br>Base Path. The base path which this validation applies to

<a id="metadata-868613"></a>&#x2022; [`metadata`](#metadata-868613) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#metadata-868613) below.

#### API Specification Validation All Spec Endpoints Fall Through Mode Fall Through Mode Custom Open API Validation Rules API Endpoint

An [`api_endpoint`](#endpoint-eb6d3c) block (within [`api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules`](#rules-7e1bb3)) supports the following:

<a id="methods-782949"></a>&#x2022; [`methods`](#methods-782949) - Optional List  Defaults to `ANY`<br>Possible values are `ANY`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`, `COPY`<br>[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY] Methods. Methods to be matched

<a id="path-f12bb6"></a>&#x2022; [`path`](#path-f12bb6) - Optional String<br>Path. Path to be matched

#### API Specification Validation All Spec Endpoints Fall Through Mode Fall Through Mode Custom Open API Validation Rules Metadata

A [`metadata`](#metadata-868613) block (within [`api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules`](#rules-7e1bb3)) supports the following:

<a id="spec-d3d0fc"></a>&#x2022; [`description_spec`](#spec-d3d0fc) - Optional String<br>Description. Human readable description

<a id="name-7c2811"></a>&#x2022; [`name`](#name-7c2811) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### API Specification Validation All Spec Endpoints Settings

A [`settings`](#settings-a83a93) block (within [`api_specification.validation_all_spec_endpoints`](#endpoints-4158a4)) supports the following:

<a id="validation-462f95"></a>&#x2022; [`oversized_body_fail_validation`](#validation-462f95) - Optional Block<br>Enable this option

<a id="validation-7ffaab"></a>&#x2022; [`oversized_body_skip_validation`](#validation-7ffaab) - Optional Block<br>Enable this option

<a id="custom-8254df"></a>&#x2022; [`property_validation_settings_custom`](#custom-8254df) - Optional Block<br>Validation Property Settings. Custom property validation settings<br>See [Property Validation Settings Custom](#custom-8254df) below.

<a id="default-f746bd"></a>&#x2022; [`property_validation_settings_default`](#default-f746bd) - Optional Block<br>Enable this option

#### API Specification Validation All Spec Endpoints Settings Property Validation Settings Custom

A [`property_validation_settings_custom`](#custom-8254df) block (within [`api_specification.validation_all_spec_endpoints.settings`](#settings-a83a93)) supports the following:

<a id="parameters-83e343"></a>&#x2022; [`query_parameters`](#parameters-83e343) - Optional Block<br>Validation Settings For Query Parameters. Custom settings for query parameters validation<br>See [Query Parameters](#parameters-83e343) below.

#### API Specification Validation All Spec Endpoints Settings Property Validation Settings Custom Query Parameters

A [`query_parameters`](#parameters-83e343) block (within [`api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom`](#custom-8254df)) supports the following:

<a id="parameters-788bd4"></a>&#x2022; [`allow_additional_parameters`](#parameters-788bd4) - Optional Block<br>Enable this option

<a id="parameters-84cc51"></a>&#x2022; [`disallow_additional_parameters`](#parameters-84cc51) - Optional Block<br>Enable this option

#### API Specification Validation All Spec Endpoints Validation Mode

A [`validation_mode`](#mode-cd4a1c) block (within [`api_specification.validation_all_spec_endpoints`](#endpoints-4158a4)) supports the following:

<a id="active-df510e"></a>&#x2022; [`response_validation_mode_active`](#active-df510e) - Optional Block<br>Open API Validation Mode Active. Validation mode properties of response<br>See [Response Validation Mode Active](#active-df510e) below.

<a id="validation-5ae35e"></a>&#x2022; [`skip_response_validation`](#validation-5ae35e) - Optional Block<br>Enable this option

<a id="validation-a6bc43"></a>&#x2022; [`skip_validation`](#validation-a6bc43) - Optional Block<br>Enable this option

<a id="active-876e02"></a>&#x2022; [`validation_mode_active`](#active-876e02) - Optional Block<br>Open API Validation Mode Active. Validation mode properties of request<br>See [Validation Mode Active](#active-876e02) below.

#### API Specification Validation All Spec Endpoints Validation Mode Response Validation Mode Active

A [`response_validation_mode_active`](#active-df510e) block (within [`api_specification.validation_all_spec_endpoints.validation_mode`](#mode-cd4a1c)) supports the following:

<a id="block-ff5c27"></a>&#x2022; [`enforcement_block`](#block-ff5c27) - Optional Block<br>Enable this option

<a id="report-0b6c08"></a>&#x2022; [`enforcement_report`](#report-0b6c08) - Optional Block<br>Enable this option

<a id="properties-138811"></a>&#x2022; [`response_validation_properties`](#properties-138811) - Optional List  Defaults to `PROPERTY_QUERY_PARAMETERS`<br>Possible values are `PROPERTY_QUERY_PARAMETERS`, `PROPERTY_PATH_PARAMETERS`, `PROPERTY_CONTENT_TYPE`, `PROPERTY_COOKIE_PARAMETERS`, `PROPERTY_HTTP_HEADERS`, `PROPERTY_HTTP_BODY`, `PROPERTY_SECURITY_SCHEMA`, `PROPERTY_RESPONSE_CODE`<br>[Enum: PROPERTY_QUERY_PARAMETERS|PROPERTY_PATH_PARAMETERS|PROPERTY_CONTENT_TYPE|PROPERTY_COOKIE_PARAMETERS|PROPERTY_HTTP_HEADERS|PROPERTY_HTTP_BODY|PROPERTY_SECURITY_SCHEMA|PROPERTY_RESPONSE_CODE] Response Validation Properties. List of properties of the response to validate according to the OpenAPI specification file (a.k.a. swagger)

#### API Specification Validation All Spec Endpoints Validation Mode Validation Mode Active

A [`validation_mode_active`](#active-876e02) block (within [`api_specification.validation_all_spec_endpoints.validation_mode`](#mode-cd4a1c)) supports the following:

<a id="block-cb8976"></a>&#x2022; [`enforcement_block`](#block-cb8976) - Optional Block<br>Enable this option

<a id="report-c50e43"></a>&#x2022; [`enforcement_report`](#report-c50e43) - Optional Block<br>Enable this option

<a id="properties-029aa9"></a>&#x2022; [`request_validation_properties`](#properties-029aa9) - Optional List  Defaults to `PROPERTY_QUERY_PARAMETERS`<br>Possible values are `PROPERTY_QUERY_PARAMETERS`, `PROPERTY_PATH_PARAMETERS`, `PROPERTY_CONTENT_TYPE`, `PROPERTY_COOKIE_PARAMETERS`, `PROPERTY_HTTP_HEADERS`, `PROPERTY_HTTP_BODY`, `PROPERTY_SECURITY_SCHEMA`, `PROPERTY_RESPONSE_CODE`<br>[Enum: PROPERTY_QUERY_PARAMETERS|PROPERTY_PATH_PARAMETERS|PROPERTY_CONTENT_TYPE|PROPERTY_COOKIE_PARAMETERS|PROPERTY_HTTP_HEADERS|PROPERTY_HTTP_BODY|PROPERTY_SECURITY_SCHEMA|PROPERTY_RESPONSE_CODE] Request Validation Properties. List of properties of the request to validate according to the OpenAPI specification file (a.k.a. swagger)

#### API Specification Validation Custom List

A [`validation_custom_list`](#list-23b577) block (within [`api_specification`](#api-specification)) supports the following:

<a id="mode-146cc3"></a>&#x2022; [`fall_through_mode`](#mode-146cc3) - Optional Block<br>Fall Through Mode.Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a. swagger) or doesn't have a specific rule in custom rules)<br>See [Fall Through Mode](#mode-146cc3) below.

<a id="rules-f51668"></a>&#x2022; [`open_api_validation_rules`](#rules-f51668) - Optional Block<br>Validation List<br>See [Open API Validation Rules](#rules-f51668) below.

<a id="settings-940e64"></a>&#x2022; [`settings`](#settings-940e64) - Optional Block<br>Common Settings. OpenAPI specification validation settings relevant for 'API Inventory' enforcement and for 'Custom list' enforcement<br>See [Settings](#settings-940e64) below.

#### API Specification Validation Custom List Fall Through Mode

A [`fall_through_mode`](#mode-146cc3) block (within [`api_specification.validation_custom_list`](#list-23b577)) supports the following:

<a id="allow-c0dd39"></a>&#x2022; [`fall_through_mode_allow`](#allow-c0dd39) - Optional Block<br>Enable this option

<a id="custom-c29bcd"></a>&#x2022; [`fall_through_mode_custom`](#custom-c29bcd) - Optional Block<br>Custom Fall Through Mode. Define the fall through settings<br>See [Fall Through Mode Custom](#custom-c29bcd) below.

#### API Specification Validation Custom List Fall Through Mode Fall Through Mode Custom

A [`fall_through_mode_custom`](#custom-c29bcd) block (within [`api_specification.validation_custom_list.fall_through_mode`](#mode-146cc3)) supports the following:

<a id="rules-ed6696"></a>&#x2022; [`open_api_validation_rules`](#rules-ed6696) - Optional Block<br>Custom Fall Through Rule List<br>See [Open API Validation Rules](#rules-ed6696) below.

#### API Specification Validation Custom List Fall Through Mode Fall Through Mode Custom Open API Validation Rules

An [`open_api_validation_rules`](#rules-ed6696) block (within [`api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom`](#custom-c29bcd)) supports the following:

<a id="block-31614e"></a>&#x2022; [`action_block`](#block-31614e) - Optional Block<br>Enable this option

<a id="report-e29f47"></a>&#x2022; [`action_report`](#report-e29f47) - Optional Block<br>Enable this option

<a id="skip-c4580b"></a>&#x2022; [`action_skip`](#skip-c4580b) - Optional Block<br>Enable this option

<a id="endpoint-997b3f"></a>&#x2022; [`api_endpoint`](#endpoint-997b3f) - Optional Block<br>API Endpoint. This defines API endpoint<br>See [API Endpoint](#endpoint-997b3f) below.

<a id="group-e515e8"></a>&#x2022; [`api_group`](#group-e515e8) - Optional String<br>API Group. The API group which this validation applies to

<a id="path-835b18"></a>&#x2022; [`base_path`](#path-835b18) - Optional String<br>Base Path. The base path which this validation applies to

<a id="metadata-6c686f"></a>&#x2022; [`metadata`](#metadata-6c686f) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#metadata-6c686f) below.

#### API Specification Validation Custom List Fall Through Mode Fall Through Mode Custom Open API Validation Rules API Endpoint

An [`api_endpoint`](#endpoint-997b3f) block (within [`api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules`](#rules-ed6696)) supports the following:

<a id="methods-9c3511"></a>&#x2022; [`methods`](#methods-9c3511) - Optional List  Defaults to `ANY`<br>Possible values are `ANY`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`, `COPY`<br>[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY] Methods. Methods to be matched

<a id="path-0f1169"></a>&#x2022; [`path`](#path-0f1169) - Optional String<br>Path. Path to be matched

#### API Specification Validation Custom List Fall Through Mode Fall Through Mode Custom Open API Validation Rules Metadata

A [`metadata`](#metadata-6c686f) block (within [`api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules`](#rules-ed6696)) supports the following:

<a id="spec-a68d3b"></a>&#x2022; [`description_spec`](#spec-a68d3b) - Optional String<br>Description. Human readable description

<a id="name-f578bb"></a>&#x2022; [`name`](#name-f578bb) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### API Specification Validation Custom List Open API Validation Rules

An [`open_api_validation_rules`](#rules-f51668) block (within [`api_specification.validation_custom_list`](#list-23b577)) supports the following:

<a id="domain-b31fd1"></a>&#x2022; [`any_domain`](#domain-b31fd1) - Optional Block<br>Enable this option

<a id="endpoint-1f50db"></a>&#x2022; [`api_endpoint`](#endpoint-1f50db) - Optional Block<br>API Endpoint. This defines API endpoint<br>See [API Endpoint](#endpoint-1f50db) below.

<a id="group-ba8ad2"></a>&#x2022; [`api_group`](#group-ba8ad2) - Optional String<br>API Group. The API group which this validation applies to

<a id="path-ca1339"></a>&#x2022; [`base_path`](#path-ca1339) - Optional String<br>Base Path. The base path which this validation applies to

<a id="metadata-304b10"></a>&#x2022; [`metadata`](#metadata-304b10) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#metadata-304b10) below.

<a id="domain-221c29"></a>&#x2022; [`specific_domain`](#domain-221c29) - Optional String<br>Specific Domain. The rule will apply for a specific domain

<a id="mode-79470e"></a>&#x2022; [`validation_mode`](#mode-79470e) - Optional Block<br>Validation Mode.Validation mode of OpenAPI specification. When a validation mismatch occurs on a request to one of the endpoints listed on the OpenAPI specification file (a.k.a. swagger)<br>See [Validation Mode](#mode-79470e) below.

#### API Specification Validation Custom List Open API Validation Rules API Endpoint

An [`api_endpoint`](#endpoint-1f50db) block (within [`api_specification.validation_custom_list.open_api_validation_rules`](#rules-f51668)) supports the following:

<a id="methods-acc02e"></a>&#x2022; [`methods`](#methods-acc02e) - Optional List  Defaults to `ANY`<br>Possible values are `ANY`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`, `COPY`<br>[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY] Methods. Methods to be matched

<a id="path-cb14d1"></a>&#x2022; [`path`](#path-cb14d1) - Optional String<br>Path. Path to be matched

#### API Specification Validation Custom List Open API Validation Rules Metadata

A [`metadata`](#metadata-304b10) block (within [`api_specification.validation_custom_list.open_api_validation_rules`](#rules-f51668)) supports the following:

<a id="spec-cef192"></a>&#x2022; [`description_spec`](#spec-cef192) - Optional String<br>Description. Human readable description

<a id="name-5500dd"></a>&#x2022; [`name`](#name-5500dd) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### API Specification Validation Custom List Open API Validation Rules Validation Mode

A [`validation_mode`](#mode-79470e) block (within [`api_specification.validation_custom_list.open_api_validation_rules`](#rules-f51668)) supports the following:

<a id="active-871b48"></a>&#x2022; [`response_validation_mode_active`](#active-871b48) - Optional Block<br>Open API Validation Mode Active. Validation mode properties of response<br>See [Response Validation Mode Active](#active-871b48) below.

<a id="validation-6f1b64"></a>&#x2022; [`skip_response_validation`](#validation-6f1b64) - Optional Block<br>Enable this option

<a id="validation-902520"></a>&#x2022; [`skip_validation`](#validation-902520) - Optional Block<br>Enable this option

<a id="active-984dc6"></a>&#x2022; [`validation_mode_active`](#active-984dc6) - Optional Block<br>Open API Validation Mode Active. Validation mode properties of request<br>See [Validation Mode Active](#active-984dc6) below.

#### API Specification Validation Custom List Open API Validation Rules Validation Mode Response Validation Mode Active

A [`response_validation_mode_active`](#active-871b48) block (within [`api_specification.validation_custom_list.open_api_validation_rules.validation_mode`](#mode-79470e)) supports the following:

<a id="block-410bdc"></a>&#x2022; [`enforcement_block`](#block-410bdc) - Optional Block<br>Enable this option

<a id="report-129a90"></a>&#x2022; [`enforcement_report`](#report-129a90) - Optional Block<br>Enable this option

<a id="properties-b162cc"></a>&#x2022; [`response_validation_properties`](#properties-b162cc) - Optional List  Defaults to `PROPERTY_QUERY_PARAMETERS`<br>Possible values are `PROPERTY_QUERY_PARAMETERS`, `PROPERTY_PATH_PARAMETERS`, `PROPERTY_CONTENT_TYPE`, `PROPERTY_COOKIE_PARAMETERS`, `PROPERTY_HTTP_HEADERS`, `PROPERTY_HTTP_BODY`, `PROPERTY_SECURITY_SCHEMA`, `PROPERTY_RESPONSE_CODE`<br>[Enum: PROPERTY_QUERY_PARAMETERS|PROPERTY_PATH_PARAMETERS|PROPERTY_CONTENT_TYPE|PROPERTY_COOKIE_PARAMETERS|PROPERTY_HTTP_HEADERS|PROPERTY_HTTP_BODY|PROPERTY_SECURITY_SCHEMA|PROPERTY_RESPONSE_CODE] Response Validation Properties. List of properties of the response to validate according to the OpenAPI specification file (a.k.a. swagger)

#### API Specification Validation Custom List Open API Validation Rules Validation Mode Validation Mode Active

A [`validation_mode_active`](#active-984dc6) block (within [`api_specification.validation_custom_list.open_api_validation_rules.validation_mode`](#mode-79470e)) supports the following:

<a id="block-d25b95"></a>&#x2022; [`enforcement_block`](#block-d25b95) - Optional Block<br>Enable this option

<a id="report-dda104"></a>&#x2022; [`enforcement_report`](#report-dda104) - Optional Block<br>Enable this option

<a id="properties-aae899"></a>&#x2022; [`request_validation_properties`](#properties-aae899) - Optional List  Defaults to `PROPERTY_QUERY_PARAMETERS`<br>Possible values are `PROPERTY_QUERY_PARAMETERS`, `PROPERTY_PATH_PARAMETERS`, `PROPERTY_CONTENT_TYPE`, `PROPERTY_COOKIE_PARAMETERS`, `PROPERTY_HTTP_HEADERS`, `PROPERTY_HTTP_BODY`, `PROPERTY_SECURITY_SCHEMA`, `PROPERTY_RESPONSE_CODE`<br>[Enum: PROPERTY_QUERY_PARAMETERS|PROPERTY_PATH_PARAMETERS|PROPERTY_CONTENT_TYPE|PROPERTY_COOKIE_PARAMETERS|PROPERTY_HTTP_HEADERS|PROPERTY_HTTP_BODY|PROPERTY_SECURITY_SCHEMA|PROPERTY_RESPONSE_CODE] Request Validation Properties. List of properties of the request to validate according to the OpenAPI specification file (a.k.a. swagger)

#### API Specification Validation Custom List Settings

A [`settings`](#settings-940e64) block (within [`api_specification.validation_custom_list`](#list-23b577)) supports the following:

<a id="validation-cfaf7f"></a>&#x2022; [`oversized_body_fail_validation`](#validation-cfaf7f) - Optional Block<br>Enable this option

<a id="validation-0639fa"></a>&#x2022; [`oversized_body_skip_validation`](#validation-0639fa) - Optional Block<br>Enable this option

<a id="custom-8e6ea6"></a>&#x2022; [`property_validation_settings_custom`](#custom-8e6ea6) - Optional Block<br>Validation Property Settings. Custom property validation settings<br>See [Property Validation Settings Custom](#custom-8e6ea6) below.

<a id="default-baec50"></a>&#x2022; [`property_validation_settings_default`](#default-baec50) - Optional Block<br>Enable this option

#### API Specification Validation Custom List Settings Property Validation Settings Custom

A [`property_validation_settings_custom`](#custom-8e6ea6) block (within [`api_specification.validation_custom_list.settings`](#settings-940e64)) supports the following:

<a id="parameters-bb35d2"></a>&#x2022; [`query_parameters`](#parameters-bb35d2) - Optional Block<br>Validation Settings For Query Parameters. Custom settings for query parameters validation<br>See [Query Parameters](#parameters-bb35d2) below.

#### API Specification Validation Custom List Settings Property Validation Settings Custom Query Parameters

A [`query_parameters`](#parameters-bb35d2) block (within [`api_specification.validation_custom_list.settings.property_validation_settings_custom`](#custom-8e6ea6)) supports the following:

<a id="parameters-547273"></a>&#x2022; [`allow_additional_parameters`](#parameters-547273) - Optional Block<br>Enable this option

<a id="parameters-22e36d"></a>&#x2022; [`disallow_additional_parameters`](#parameters-22e36d) - Optional Block<br>Enable this option

#### API Testing

An [`api_testing`](#api-testing) block supports the following:

<a id="api-testing-custom-header-value"></a>&#x2022; [`custom_header_value`](#api-testing-custom-header-value) - Optional String<br>Custom Header. Add x-f5-API-testing-identifier header value to prevent security flags on API testing traffic

<a id="api-testing-domains"></a>&#x2022; [`domains`](#api-testing-domains) - Optional Block<br>Testing Environments. Add and configure testing domains and credentials<br>See [Domains](#api-testing-domains) below.

<a id="api-testing-every-day"></a>&#x2022; [`every_day`](#api-testing-every-day) - Optional Block<br>Enable this option

<a id="api-testing-every-month"></a>&#x2022; [`every_month`](#api-testing-every-month) - Optional Block<br>Enable this option

<a id="api-testing-every-week"></a>&#x2022; [`every_week`](#api-testing-every-week) - Optional Block<br>Enable this option

#### API Testing Domains

A [`domains`](#api-testing-domains) block (within [`api_testing`](#api-testing)) supports the following:

<a id="methods-c3ca06"></a>&#x2022; [`allow_destructive_methods`](#methods-c3ca06) - Optional Bool<br>Use Destructive Methods (e.g., DELETE, PUT). Enable to allow API test to execute destructive methods. Be cautious as these can alter or delete data

<a id="api-testing-domains-credentials"></a>&#x2022; [`credentials`](#api-testing-domains-credentials) - Optional Block<br>Credentials. Add credentials for API testing to use in the selected environment<br>See [Credentials](#api-testing-domains-credentials) below.

<a id="api-testing-domains-domain"></a>&#x2022; [`domain`](#api-testing-domains-domain) - Optional String<br>Domain. Add your testing environment domain. Be aware that running tests on a production domain can impact live applications, as API testing cannot distinguish between production and testing environments

#### API Testing Domains Credentials

A [`credentials`](#api-testing-domains-credentials) block (within [`api_testing.domains`](#api-testing-domains)) supports the following:

<a id="api-testing-domains-credentials-admin"></a>&#x2022; [`admin`](#api-testing-domains-credentials-admin) - Optional Block<br>Enable this option

<a id="api-testing-domains-credentials-api-key"></a>&#x2022; [`api_key`](#api-testing-domains-credentials-api-key) - Optional Block<br>API Key<br>See [API Key](#api-testing-domains-credentials-api-key) below.

<a id="auth-4868f3"></a>&#x2022; [`basic_auth`](#auth-4868f3) - Optional Block<br>Basic Authentication<br>See [Basic Auth](#auth-4868f3) below.

<a id="token-2a2002"></a>&#x2022; [`bearer_token`](#token-2a2002) - Optional Block<br>Bearer<br>See [Bearer Token](#token-2a2002) below.

<a id="name-8eb15c"></a>&#x2022; [`credential_name`](#name-8eb15c) - Optional String<br>Name. Enter a unique name for the credentials used in API testing

<a id="endpoint-08dc4d"></a>&#x2022; [`login_endpoint`](#endpoint-08dc4d) - Optional Block<br>Login Endpoint<br>See [Login Endpoint](#endpoint-08dc4d) below.

<a id="standard-8b74a1"></a>&#x2022; [`standard`](#standard-8b74a1) - Optional Block<br>Enable this option

#### API Testing Domains Credentials API Key

An [`api_key`](#api-testing-domains-credentials-api-key) block (within [`api_testing.domains.credentials`](#api-testing-domains-credentials)) supports the following:

<a id="key-a622d1"></a>&#x2022; [`key`](#key-a622d1) - Optional String<br>Key

<a id="value-f4a033"></a>&#x2022; [`value`](#value-f4a033) - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Value](#value-f4a033) below.

#### API Testing Domains Credentials API Key Value

A [`value`](#value-f4a033) block (within [`api_testing.domains.credentials.api_key`](#api-testing-domains-credentials-api-key)) supports the following:

<a id="info-b2f866"></a>&#x2022; [`blindfold_secret_info`](#info-b2f866) - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#info-b2f866) below.

<a id="info-d18192"></a>&#x2022; [`clear_secret_info`](#info-d18192) - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#info-d18192) below.

#### API Testing Domains Credentials API Key Value Blindfold Secret Info

A [`blindfold_secret_info`](#info-b2f866) block (within [`api_testing.domains.credentials.api_key.value`](#value-f4a033)) supports the following:

<a id="provider-1fe897"></a>&#x2022; [`decryption_provider`](#provider-1fe897) - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

<a id="location-9f76a4"></a>&#x2022; [`location`](#location-9f76a4) - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

<a id="provider-cff879"></a>&#x2022; [`store_provider`](#provider-cff879) - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

#### API Testing Domains Credentials API Key Value Clear Secret Info

A [`clear_secret_info`](#info-d18192) block (within [`api_testing.domains.credentials.api_key.value`](#value-f4a033)) supports the following:

<a id="ref-5643b5"></a>&#x2022; [`provider_ref`](#ref-5643b5) - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

<a id="url-6994ed"></a>&#x2022; [`url`](#url-6994ed) - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

#### API Testing Domains Credentials Basic Auth

A [`basic_auth`](#auth-4868f3) block (within [`api_testing.domains.credentials`](#api-testing-domains-credentials)) supports the following:

<a id="password-e6a065"></a>&#x2022; [`password`](#password-e6a065) - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Password](#password-e6a065) below.

<a id="user-2d5b74"></a>&#x2022; [`user`](#user-2d5b74) - Optional String<br>User

#### API Testing Domains Credentials Basic Auth Password

A [`password`](#password-e6a065) block (within [`api_testing.domains.credentials.basic_auth`](#auth-4868f3)) supports the following:

<a id="info-0decf2"></a>&#x2022; [`blindfold_secret_info`](#info-0decf2) - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#info-0decf2) below.

<a id="info-71b4da"></a>&#x2022; [`clear_secret_info`](#info-71b4da) - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#info-71b4da) below.

#### API Testing Domains Credentials Basic Auth Password Blindfold Secret Info

A [`blindfold_secret_info`](#info-0decf2) block (within [`api_testing.domains.credentials.basic_auth.password`](#password-e6a065)) supports the following:

<a id="provider-493e59"></a>&#x2022; [`decryption_provider`](#provider-493e59) - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

<a id="location-2fffc7"></a>&#x2022; [`location`](#location-2fffc7) - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

<a id="provider-e4aced"></a>&#x2022; [`store_provider`](#provider-e4aced) - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

#### API Testing Domains Credentials Basic Auth Password Clear Secret Info

A [`clear_secret_info`](#info-71b4da) block (within [`api_testing.domains.credentials.basic_auth.password`](#password-e6a065)) supports the following:

<a id="ref-3909a9"></a>&#x2022; [`provider_ref`](#ref-3909a9) - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

<a id="url-b5d1ef"></a>&#x2022; [`url`](#url-b5d1ef) - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

#### API Testing Domains Credentials Bearer Token

A [`bearer_token`](#token-2a2002) block (within [`api_testing.domains.credentials`](#api-testing-domains-credentials)) supports the following:

<a id="token-7fcc22"></a>&#x2022; [`token`](#token-7fcc22) - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Token](#token-7fcc22) below.

#### API Testing Domains Credentials Bearer Token Token

A [`token`](#token-7fcc22) block (within [`api_testing.domains.credentials.bearer_token`](#token-2a2002)) supports the following:

<a id="info-ce1236"></a>&#x2022; [`blindfold_secret_info`](#info-ce1236) - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#info-ce1236) below.

<a id="info-5e8cda"></a>&#x2022; [`clear_secret_info`](#info-5e8cda) - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#info-5e8cda) below.

#### API Testing Domains Credentials Bearer Token Token Blindfold Secret Info

A [`blindfold_secret_info`](#info-ce1236) block (within [`api_testing.domains.credentials.bearer_token.token`](#token-7fcc22)) supports the following:

<a id="provider-11e2c2"></a>&#x2022; [`decryption_provider`](#provider-11e2c2) - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

<a id="location-d3cf52"></a>&#x2022; [`location`](#location-d3cf52) - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

<a id="provider-44e876"></a>&#x2022; [`store_provider`](#provider-44e876) - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

#### API Testing Domains Credentials Bearer Token Token Clear Secret Info

A [`clear_secret_info`](#info-5e8cda) block (within [`api_testing.domains.credentials.bearer_token.token`](#token-7fcc22)) supports the following:

<a id="ref-0dfaf8"></a>&#x2022; [`provider_ref`](#ref-0dfaf8) - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

<a id="url-05f153"></a>&#x2022; [`url`](#url-05f153) - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

#### API Testing Domains Credentials Login Endpoint

A [`login_endpoint`](#endpoint-08dc4d) block (within [`api_testing.domains.credentials`](#api-testing-domains-credentials)) supports the following:

<a id="payload-cb9ddd"></a>&#x2022; [`json_payload`](#payload-cb9ddd) - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [JSON Payload](#payload-cb9ddd) below.

<a id="method-ccde95"></a>&#x2022; [`method`](#method-ccde95) - Optional String  Defaults to `ANY`<br>Possible values are `ANY`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`, `COPY`<br>[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY] HTTP Method. Specifies the HTTP method used to access a resource. Any HTTP Method

<a id="path-fc4b22"></a>&#x2022; [`path`](#path-fc4b22) - Optional String<br>Path

<a id="key-04416f"></a>&#x2022; [`token_response_key`](#key-04416f) - Optional String<br>Token Response Key. Specifies how to handle the API response, extracting authentication tokens

#### API Testing Domains Credentials Login Endpoint JSON Payload

A [`json_payload`](#payload-cb9ddd) block (within [`api_testing.domains.credentials.login_endpoint`](#endpoint-08dc4d)) supports the following:

<a id="info-2f70ce"></a>&#x2022; [`blindfold_secret_info`](#info-2f70ce) - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#info-2f70ce) below.

<a id="info-2574ee"></a>&#x2022; [`clear_secret_info`](#info-2574ee) - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#info-2574ee) below.

#### API Testing Domains Credentials Login Endpoint JSON Payload Blindfold Secret Info

A [`blindfold_secret_info`](#info-2f70ce) block (within [`api_testing.domains.credentials.login_endpoint.json_payload`](#payload-cb9ddd)) supports the following:

<a id="provider-e2e9cb"></a>&#x2022; [`decryption_provider`](#provider-e2e9cb) - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

<a id="location-2cd17b"></a>&#x2022; [`location`](#location-2cd17b) - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

<a id="provider-eceebd"></a>&#x2022; [`store_provider`](#provider-eceebd) - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

#### API Testing Domains Credentials Login Endpoint JSON Payload Clear Secret Info

A [`clear_secret_info`](#info-2574ee) block (within [`api_testing.domains.credentials.login_endpoint.json_payload`](#payload-cb9ddd)) supports the following:

<a id="ref-f10ba7"></a>&#x2022; [`provider_ref`](#ref-f10ba7) - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

<a id="url-856dba"></a>&#x2022; [`url`](#url-856dba) - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

#### App Firewall

An [`app_firewall`](#app-firewall) block supports the following:

<a id="app-firewall-name"></a>&#x2022; [`name`](#app-firewall-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="app-firewall-namespace"></a>&#x2022; [`namespace`](#app-firewall-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="app-firewall-tenant"></a>&#x2022; [`tenant`](#app-firewall-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Blocked Clients

A [`blocked_clients`](#blocked-clients) block supports the following:

<a id="blocked-clients-actions"></a>&#x2022; [`actions`](#blocked-clients-actions) - Optional List  Defaults to `SKIP_PROCESSING_WAF`<br>Possible values are `SKIP_PROCESSING_WAF`, `SKIP_PROCESSING_BOT`, `SKIP_PROCESSING_MUM`, `SKIP_PROCESSING_IP_REPUTATION`, `SKIP_PROCESSING_API_PROTECTION`, `SKIP_PROCESSING_OAS_VALIDATION`, `SKIP_PROCESSING_DDOS_PROTECTION`, `SKIP_PROCESSING_THREAT_MESH`, `SKIP_PROCESSING_MALWARE_PROTECTION`<br>[Enum: SKIP_PROCESSING_WAF|SKIP_PROCESSING_BOT|SKIP_PROCESSING_MUM|SKIP_PROCESSING_IP_REPUTATION|SKIP_PROCESSING_API_PROTECTION|SKIP_PROCESSING_OAS_VALIDATION|SKIP_PROCESSING_DDOS_PROTECTION|SKIP_PROCESSING_THREAT_MESH|SKIP_PROCESSING_MALWARE_PROTECTION] Actions. Actions that should be taken when client identifier matches the rule

<a id="blocked-clients-as-number"></a>&#x2022; [`as_number`](#blocked-clients-as-number) - Optional Number<br>AS Number. RFC 6793 defined 4-byte AS number

<a id="blocked-clients-bot-skip-processing"></a>&#x2022; [`bot_skip_processing`](#blocked-clients-bot-skip-processing) - Optional Block<br>Enable this option

<a id="blocked-clients-expiration-timestamp"></a>&#x2022; [`expiration_timestamp`](#blocked-clients-expiration-timestamp) - Optional String<br>Expiration Timestamp. The expiration_timestamp is the RFC 3339 format timestamp at which the containing rule is considered to be logically expired. The rule continues to exist in the configuration but is not applied anymore

<a id="blocked-clients-http-header"></a>&#x2022; [`http_header`](#blocked-clients-http-header) - Optional Block<br>HTTP Header. Request header name and value pairs<br>See [HTTP Header](#blocked-clients-http-header) below.

<a id="blocked-clients-ip-prefix"></a>&#x2022; [`ip_prefix`](#blocked-clients-ip-prefix) - Optional String<br>IPv4 Prefix. IPv4 prefix string

<a id="blocked-clients-ipv6-prefix"></a>&#x2022; [`ipv6_prefix`](#blocked-clients-ipv6-prefix) - Optional String<br>IPv6 Prefix. IPv6 prefix string

<a id="blocked-clients-metadata"></a>&#x2022; [`metadata`](#blocked-clients-metadata) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#blocked-clients-metadata) below.

<a id="blocked-clients-skip-processing"></a>&#x2022; [`skip_processing`](#blocked-clients-skip-processing) - Optional Block<br>Enable this option

<a id="blocked-clients-user-identifier"></a>&#x2022; [`user_identifier`](#blocked-clients-user-identifier) - Optional String<br>User Identifier. Identify user based on user identifier. User identifier value needs to be copied from security event

<a id="blocked-clients-waf-skip-processing"></a>&#x2022; [`waf_skip_processing`](#blocked-clients-waf-skip-processing) - Optional Block<br>Enable this option

#### Blocked Clients HTTP Header

A [`http_header`](#blocked-clients-http-header) block (within [`blocked_clients`](#blocked-clients)) supports the following:

<a id="blocked-clients-http-header-headers"></a>&#x2022; [`headers`](#blocked-clients-http-header-headers) - Optional Block<br>HTTP Headers. List of HTTP header name and value pairs<br>See [Headers](#blocked-clients-http-header-headers) below.

#### Blocked Clients HTTP Header Headers

A [`headers`](#blocked-clients-http-header-headers) block (within [`blocked_clients.http_header`](#blocked-clients-http-header)) supports the following:

<a id="exact-a1dbef"></a>&#x2022; [`exact`](#exact-a1dbef) - Optional String<br>Exact. Header value to match exactly

<a id="match-b2ef8e"></a>&#x2022; [`invert_match`](#match-b2ef8e) - Optional Bool<br>NOT of match. Invert the result of the match to detect missing header or non-matching value

<a id="name-dc6d90"></a>&#x2022; [`name`](#name-dc6d90) - Optional String<br>Name. Name of the header

<a id="presence-659464"></a>&#x2022; [`presence`](#presence-659464) - Optional Bool<br>Presence. If true, check for presence of header

<a id="regex-6757d0"></a>&#x2022; [`regex`](#regex-6757d0) - Optional String<br>Regex. Regex match of the header value in re2 format

#### Blocked Clients Metadata

A [`metadata`](#blocked-clients-metadata) block (within [`blocked_clients`](#blocked-clients)) supports the following:

<a id="spec-b8ecbc"></a>&#x2022; [`description_spec`](#spec-b8ecbc) - Optional String<br>Description. Human readable description

<a id="blocked-clients-metadata-name"></a>&#x2022; [`name`](#blocked-clients-metadata-name) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### Bot Defense

A [`bot_defense`](#bot-defense) block supports the following:

<a id="bot-defense-disable-cors-support"></a>&#x2022; [`disable_cors_support`](#bot-defense-disable-cors-support) - Optional Block<br>Enable this option

<a id="bot-defense-enable-cors-support"></a>&#x2022; [`enable_cors_support`](#bot-defense-enable-cors-support) - Optional Block<br>Enable this option

<a id="bot-defense-policy"></a>&#x2022; [`policy`](#bot-defense-policy) - Optional Block<br>Bot Defense Policy. This defines various configuration options for Bot Defense policy<br>See [Policy](#bot-defense-policy) below.

<a id="bot-defense-regional-endpoint"></a>&#x2022; [`regional_endpoint`](#bot-defense-regional-endpoint) - Optional String  Defaults to `AUTO`<br>Possible values are `AUTO`, `US`, `EU`, `ASIA`<br>[Enum: AUTO|US|EU|ASIA] Bot Defense Region. Defines a selection for Bot Defense region - AUTO: AUTO Automatic selection based on client IP address - US: US US region - EU: EU European Union region - ASIA: ASIA Asia region

<a id="bot-defense-timeout"></a>&#x2022; [`timeout`](#bot-defense-timeout) - Optional Number<br>Timeout. The timeout for the inference check, in milliseconds

#### Bot Defense Policy

A [`policy`](#bot-defense-policy) block (within [`bot_defense`](#bot-defense)) supports the following:

<a id="bot-defense-policy-disable-js-insert"></a>&#x2022; [`disable_js_insert`](#bot-defense-policy-disable-js-insert) - Optional Block<br>Enable this option

<a id="bot-defense-policy-disable-mobile-sdk"></a>&#x2022; [`disable_mobile_sdk`](#bot-defense-policy-disable-mobile-sdk) - Optional Block<br>Enable this option

<a id="bot-defense-policy-javascript-mode"></a>&#x2022; [`javascript_mode`](#bot-defense-policy-javascript-mode) - Optional String  Defaults to `ASYNC_JS_NO_CACHING`<br>Possible values are `ASYNC_JS_NO_CACHING`, `ASYNC_JS_CACHING`, `SYNC_JS_NO_CACHING`, `SYNC_JS_CACHING`<br>[Enum: ASYNC_JS_NO_CACHING|ASYNC_JS_CACHING|SYNC_JS_NO_CACHING|SYNC_JS_CACHING] Web Client JavaScript Mode. Web Client JavaScript Mode. Bot Defense JavaScript for telemetry collection is requested asynchronously, and it is non-cacheable Bot Defense JavaScript for telemetry collection is requested asynchronously, and it is cacheable Bot Defense JavaScript for telemetry collection is requested synchronously, and it is non-cacheable Bot Defense JavaScript for telemetry collection is requested synchronously, and it is cacheable

<a id="bot-defense-policy-js-download-path"></a>&#x2022; [`js_download_path`](#bot-defense-policy-js-download-path) - Optional String<br>JavaScript Download Path. Customize Bot Defense Client JavaScript path. If not specified, default `/common.js`

<a id="bot-defense-policy-js-insert-all-pages"></a>&#x2022; [`js_insert_all_pages`](#bot-defense-policy-js-insert-all-pages) - Optional Block<br>Insert Bot Defense JavaScript in All Pages. Insert Bot Defense JavaScript in all pages<br>See [Js Insert All Pages](#bot-defense-policy-js-insert-all-pages) below.

<a id="except-2f0f51"></a>&#x2022; [`js_insert_all_pages_except`](#except-2f0f51) - Optional Block<br>Insert JavaScript in All Pages with the Exceptions. Insert Bot Defense JavaScript in all pages with the exceptions<br>See [Js Insert All Pages Except](#except-2f0f51) below.

<a id="bot-defense-policy-js-insertion-rules"></a>&#x2022; [`js_insertion_rules`](#bot-defense-policy-js-insertion-rules) - Optional Block<br>JavaScript Custom Insertion Rules. This defines custom JavaScript insertion rules for Bot Defense Policy<br>See [Js Insertion Rules](#bot-defense-policy-js-insertion-rules) below.

<a id="bot-defense-policy-mobile-sdk-config"></a>&#x2022; [`mobile_sdk_config`](#bot-defense-policy-mobile-sdk-config) - Optional Block<br>Mobile SDK Configuration. Mobile SDK configuration<br>See [Mobile Sdk Config](#bot-defense-policy-mobile-sdk-config) below.

<a id="endpoints-01a2f3"></a>&#x2022; [`protected_app_endpoints`](#endpoints-01a2f3) - Optional Block<br>App Endpoint Type. List of protected endpoints. Limit: Approx '128 endpoints per Load Balancer (LB)' upto 4 LBs, '32 endpoints per LB' after 4 LBs<br>See [Protected App Endpoints](#endpoints-01a2f3) below.

#### Bot Defense Policy Js Insert All Pages

A [`js_insert_all_pages`](#bot-defense-policy-js-insert-all-pages) block (within [`bot_defense.policy`](#bot-defense-policy)) supports the following:

<a id="location-3a398d"></a>&#x2022; [`javascript_location`](#location-3a398d) - Optional String  Defaults to `AFTER_HEAD`<br>Possible values are `AFTER_HEAD`, `AFTER_TITLE_END`, `BEFORE_SCRIPT`<br>[Enum: AFTER_HEAD|AFTER_TITLE_END|BEFORE_SCRIPT] JavaScript Location. All inside networks. Insert JavaScript after `<head>` tag Insert JavaScript after `</title>` tag. Insert JavaScript before first `<script>` tag

#### Bot Defense Policy Js Insert All Pages Except

A [`js_insert_all_pages_except`](#except-2f0f51) block (within [`bot_defense.policy`](#bot-defense-policy)) supports the following:

<a id="list-c11e8a"></a>&#x2022; [`exclude_list`](#list-c11e8a) - Optional Block<br>Exclude Pages. Optional JavaScript insertions exclude list of domain and path matchers<br>See [Exclude List](#list-c11e8a) below.

<a id="location-7d08dc"></a>&#x2022; [`javascript_location`](#location-7d08dc) - Optional String  Defaults to `AFTER_HEAD`<br>Possible values are `AFTER_HEAD`, `AFTER_TITLE_END`, `BEFORE_SCRIPT`<br>[Enum: AFTER_HEAD|AFTER_TITLE_END|BEFORE_SCRIPT] JavaScript Location. All inside networks. Insert JavaScript after `<head>` tag Insert JavaScript after `</title>` tag. Insert JavaScript before first `<script>` tag

#### Bot Defense Policy Js Insert All Pages Except Exclude List

An [`exclude_list`](#list-c11e8a) block (within [`bot_defense.policy.js_insert_all_pages_except`](#except-2f0f51)) supports the following:

<a id="domain-73ce27"></a>&#x2022; [`any_domain`](#domain-73ce27) - Optional Block<br>Enable this option

<a id="domain-503442"></a>&#x2022; [`domain`](#domain-503442) - Optional Block<br>Domains. Domains names<br>See [Domain](#domain-503442) below.

<a id="metadata-f70b11"></a>&#x2022; [`metadata`](#metadata-f70b11) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#metadata-f70b11) below.

<a id="path-e8b4e3"></a>&#x2022; [`path`](#path-e8b4e3) - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#path-e8b4e3) below.

#### Bot Defense Policy Js Insert All Pages Except Exclude List Domain

A [`domain`](#domain-503442) block (within [`bot_defense.policy.js_insert_all_pages_except.exclude_list`](#list-c11e8a)) supports the following:

<a id="value-64604c"></a>&#x2022; [`exact_value`](#value-64604c) - Optional String<br>Exact Value. Exact domain name

<a id="value-a33f5c"></a>&#x2022; [`regex_value`](#value-a33f5c) - Optional String<br>Regex Values of Domains. Regular Expression value for the domain name

<a id="value-ae4d1e"></a>&#x2022; [`suffix_value`](#value-ae4d1e) - Optional String<br>Suffix Value. Suffix of domain name e.g 'xyz.com' will match '*.xyz.com' and 'xyz.com'

#### Bot Defense Policy Js Insert All Pages Except Exclude List Metadata

A [`metadata`](#metadata-f70b11) block (within [`bot_defense.policy.js_insert_all_pages_except.exclude_list`](#list-c11e8a)) supports the following:

<a id="spec-aabb87"></a>&#x2022; [`description_spec`](#spec-aabb87) - Optional String<br>Description. Human readable description

<a id="name-644012"></a>&#x2022; [`name`](#name-644012) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### Bot Defense Policy Js Insert All Pages Except Exclude List Path

A [`path`](#path-e8b4e3) block (within [`bot_defense.policy.js_insert_all_pages_except.exclude_list`](#list-c11e8a)) supports the following:

<a id="path-39fb60"></a>&#x2022; [`path`](#path-39fb60) - Optional String<br>Exact. Exact path value to match

<a id="prefix-fe375b"></a>&#x2022; [`prefix`](#prefix-fe375b) - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

<a id="regex-93d4f7"></a>&#x2022; [`regex`](#regex-93d4f7) - Optional String<br>Regex. Regular expression of path match (e.g. the value .* will match on all paths)

#### Bot Defense Policy Js Insertion Rules

A [`js_insertion_rules`](#bot-defense-policy-js-insertion-rules) block (within [`bot_defense.policy`](#bot-defense-policy)) supports the following:

<a id="list-51668b"></a>&#x2022; [`exclude_list`](#list-51668b) - Optional Block<br>Exclude Paths. Optional JavaScript insertions exclude list of domain and path matchers<br>See [Exclude List](#list-51668b) below.

<a id="rules-15d983"></a>&#x2022; [`rules`](#rules-15d983) - Optional Block<br>JavaScript Insertions. Required list of pages to insert Bot Defense client JavaScript<br>See [Rules](#rules-15d983) below.

#### Bot Defense Policy Js Insertion Rules Exclude List

An [`exclude_list`](#list-51668b) block (within [`bot_defense.policy.js_insertion_rules`](#bot-defense-policy-js-insertion-rules)) supports the following:

<a id="domain-090b66"></a>&#x2022; [`any_domain`](#domain-090b66) - Optional Block<br>Enable this option

<a id="domain-47cfd3"></a>&#x2022; [`domain`](#domain-47cfd3) - Optional Block<br>Domains. Domains names<br>See [Domain](#domain-47cfd3) below.

<a id="metadata-7d33fd"></a>&#x2022; [`metadata`](#metadata-7d33fd) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#metadata-7d33fd) below.

<a id="path-a9cb42"></a>&#x2022; [`path`](#path-a9cb42) - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#path-a9cb42) below.

#### Bot Defense Policy Js Insertion Rules Exclude List Domain

A [`domain`](#domain-47cfd3) block (within [`bot_defense.policy.js_insertion_rules.exclude_list`](#list-51668b)) supports the following:

<a id="value-19618a"></a>&#x2022; [`exact_value`](#value-19618a) - Optional String<br>Exact Value. Exact domain name

<a id="value-84ab50"></a>&#x2022; [`regex_value`](#value-84ab50) - Optional String<br>Regex Values of Domains. Regular Expression value for the domain name

<a id="value-f83edf"></a>&#x2022; [`suffix_value`](#value-f83edf) - Optional String<br>Suffix Value. Suffix of domain name e.g 'xyz.com' will match '*.xyz.com' and 'xyz.com'

#### Bot Defense Policy Js Insertion Rules Exclude List Metadata

A [`metadata`](#metadata-7d33fd) block (within [`bot_defense.policy.js_insertion_rules.exclude_list`](#list-51668b)) supports the following:

<a id="spec-1f6f3a"></a>&#x2022; [`description_spec`](#spec-1f6f3a) - Optional String<br>Description. Human readable description

<a id="name-487812"></a>&#x2022; [`name`](#name-487812) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### Bot Defense Policy Js Insertion Rules Exclude List Path

A [`path`](#path-a9cb42) block (within [`bot_defense.policy.js_insertion_rules.exclude_list`](#list-51668b)) supports the following:

<a id="path-0e9a9e"></a>&#x2022; [`path`](#path-0e9a9e) - Optional String<br>Exact. Exact path value to match

<a id="prefix-635824"></a>&#x2022; [`prefix`](#prefix-635824) - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

<a id="regex-1d9ec1"></a>&#x2022; [`regex`](#regex-1d9ec1) - Optional String<br>Regex. Regular expression of path match (e.g. the value .* will match on all paths)

#### Bot Defense Policy Js Insertion Rules Rules

A [`rules`](#rules-15d983) block (within [`bot_defense.policy.js_insertion_rules`](#bot-defense-policy-js-insertion-rules)) supports the following:

<a id="domain-f27f00"></a>&#x2022; [`any_domain`](#domain-f27f00) - Optional Block<br>Enable this option

<a id="domain-834b0f"></a>&#x2022; [`domain`](#domain-834b0f) - Optional Block<br>Domains. Domains names<br>See [Domain](#domain-834b0f) below.

<a id="location-16277f"></a>&#x2022; [`javascript_location`](#location-16277f) - Optional String  Defaults to `AFTER_HEAD`<br>Possible values are `AFTER_HEAD`, `AFTER_TITLE_END`, `BEFORE_SCRIPT`<br>[Enum: AFTER_HEAD|AFTER_TITLE_END|BEFORE_SCRIPT] JavaScript Location. All inside networks. Insert JavaScript after `<head>` tag Insert JavaScript after `</title>` tag. Insert JavaScript before first `<script>` tag

<a id="metadata-e15703"></a>&#x2022; [`metadata`](#metadata-e15703) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#metadata-e15703) below.

<a id="path-711518"></a>&#x2022; [`path`](#path-711518) - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#path-711518) below.

#### Bot Defense Policy Js Insertion Rules Rules Domain

A [`domain`](#domain-834b0f) block (within [`bot_defense.policy.js_insertion_rules.rules`](#rules-15d983)) supports the following:

<a id="value-761413"></a>&#x2022; [`exact_value`](#value-761413) - Optional String<br>Exact Value. Exact domain name

<a id="value-626e98"></a>&#x2022; [`regex_value`](#value-626e98) - Optional String<br>Regex Values of Domains. Regular Expression value for the domain name

<a id="value-a64bf5"></a>&#x2022; [`suffix_value`](#value-a64bf5) - Optional String<br>Suffix Value. Suffix of domain name e.g 'xyz.com' will match '*.xyz.com' and 'xyz.com'

#### Bot Defense Policy Js Insertion Rules Rules Metadata

A [`metadata`](#metadata-e15703) block (within [`bot_defense.policy.js_insertion_rules.rules`](#rules-15d983)) supports the following:

<a id="spec-bd4771"></a>&#x2022; [`description_spec`](#spec-bd4771) - Optional String<br>Description. Human readable description

<a id="name-1bd3d4"></a>&#x2022; [`name`](#name-1bd3d4) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### Bot Defense Policy Js Insertion Rules Rules Path

A [`path`](#path-711518) block (within [`bot_defense.policy.js_insertion_rules.rules`](#rules-15d983)) supports the following:

<a id="path-6d550e"></a>&#x2022; [`path`](#path-6d550e) - Optional String<br>Exact. Exact path value to match

<a id="prefix-cee2d1"></a>&#x2022; [`prefix`](#prefix-cee2d1) - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

<a id="regex-597dd8"></a>&#x2022; [`regex`](#regex-597dd8) - Optional String<br>Regex. Regular expression of path match (e.g. the value .* will match on all paths)

#### Bot Defense Policy Mobile Sdk Config

A [`mobile_sdk_config`](#bot-defense-policy-mobile-sdk-config) block (within [`bot_defense.policy`](#bot-defense-policy)) supports the following:

<a id="identifier-e34e48"></a>&#x2022; [`mobile_identifier`](#identifier-e34e48) - Optional Block<br>Mobile Traffic Identifier. Mobile traffic identifier type<br>See [Mobile Identifier](#identifier-e34e48) below.

#### Bot Defense Policy Mobile Sdk Config Mobile Identifier

A [`mobile_identifier`](#identifier-e34e48) block (within [`bot_defense.policy.mobile_sdk_config`](#bot-defense-policy-mobile-sdk-config)) supports the following:

<a id="headers-529e3c"></a>&#x2022; [`headers`](#headers-529e3c) - Optional Block<br>Headers. Headers that can be used to identify mobile traffic<br>See [Headers](#headers-529e3c) below.

#### Bot Defense Policy Mobile Sdk Config Mobile Identifier Headers

A [`headers`](#headers-529e3c) block (within [`bot_defense.policy.mobile_sdk_config.mobile_identifier`](#identifier-e34e48)) supports the following:

<a id="present-c23fa6"></a>&#x2022; [`check_not_present`](#present-c23fa6) - Optional Block<br>Enable this option

<a id="present-1d1c99"></a>&#x2022; [`check_present`](#present-1d1c99) - Optional Block<br>Enable this option

<a id="item-6622e2"></a>&#x2022; [`item`](#item-6622e2) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#item-6622e2) below.

<a id="name-581daa"></a>&#x2022; [`name`](#name-581daa) - Optional String<br>Header Name. A case-insensitive HTTP header name

#### Bot Defense Policy Mobile Sdk Config Mobile Identifier Headers Item

An [`item`](#item-6622e2) block (within [`bot_defense.policy.mobile_sdk_config.mobile_identifier.headers`](#headers-529e3c)) supports the following:

<a id="values-f1a647"></a>&#x2022; [`exact_values`](#values-f1a647) - Optional List<br>Exact Values. A list of exact values to match the input against

<a id="values-102e01"></a>&#x2022; [`regex_values`](#values-102e01) - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="transformers-e54cc5"></a>&#x2022; [`transformers`](#transformers-e54cc5) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>[Enum: LOWER_CASE|UPPER_CASE|BASE64_DECODE|NORMALIZE_PATH|REMOVE_WHITESPACE|URL_DECODE|TRIM_LEFT|TRIM_RIGHT|TRIM] Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### Bot Defense Policy Protected App Endpoints

A [`protected_app_endpoints`](#endpoints-01a2f3) block (within [`bot_defense.policy`](#bot-defense-policy)) supports the following:

<a id="bots-cfdb6f"></a>&#x2022; [`allow_good_bots`](#bots-cfdb6f) - Optional Block<br>Enable this option

<a id="domain-f4f253"></a>&#x2022; [`any_domain`](#domain-f4f253) - Optional Block<br>Enable this option

<a id="domain-18bf1a"></a>&#x2022; [`domain`](#domain-18bf1a) - Optional Block<br>Domains. Domains names<br>See [Domain](#domain-18bf1a) below.

<a id="label-244fef"></a>&#x2022; [`flow_label`](#label-244fef) - Optional Block<br>Bot Defense Flow Label Category. Bot Defense Flow Label Category allows to associate traffic with selected category<br>See [Flow Label](#label-244fef) below.

<a id="headers-986193"></a>&#x2022; [`headers`](#headers-986193) - Optional Block<br>HTTP Headers. A list of predicates for various HTTP headers that need to match. The criteria for matching each HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values are extracted from the request API as a list of strings for each HTTP header type. Note that all specified header predicates must evaluate to true<br>See [Headers](#headers-986193) below.

<a id="methods-2d1fa7"></a>&#x2022; [`http_methods`](#methods-2d1fa7) - Optional List  Defaults to `METHOD_ANY`<br>Possible values are `METHOD_ANY`, `METHOD_GET`, `METHOD_POST`, `METHOD_PUT`, `METHOD_PATCH`, `METHOD_DELETE`, `METHOD_GET_DOCUMENT`<br>[Enum: METHOD_ANY|METHOD_GET|METHOD_POST|METHOD_PUT|METHOD_PATCH|METHOD_DELETE|METHOD_GET_DOCUMENT] HTTP Methods. List of HTTP methods

<a id="metadata-c93137"></a>&#x2022; [`metadata`](#metadata-c93137) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#metadata-c93137) below.

<a id="bots-5c9c05"></a>&#x2022; [`mitigate_good_bots`](#bots-5c9c05) - Optional Block<br>Enable this option

<a id="mitigation-cc96eb"></a>&#x2022; [`mitigation`](#mitigation-cc96eb) - Optional Block<br>Bot Mitigation Action. Modify Bot Defense behavior for a matching request<br>See [Mitigation](#mitigation-cc96eb) below.

<a id="mobile-2839a0"></a>&#x2022; [`mobile`](#mobile-2839a0) - Optional Block<br>Enable this option

<a id="path-d5ee15"></a>&#x2022; [`path`](#path-d5ee15) - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#path-d5ee15) below.

<a id="protocol-21c1f1"></a>&#x2022; [`protocol`](#protocol-21c1f1) - Optional String  Defaults to `BOTH`<br>Possible values are `BOTH`, `HTTP`, `HTTPS`<br>[Enum: BOTH|HTTP|HTTPS] URL Scheme. SchemeType is used to indicate URL scheme. - BOTH: BOTH URL scheme for HTTPS:// or `HTTP://.` - HTTP: HTTP URL scheme HTTP:// only. - HTTPS: HTTPS URL scheme HTTPS:// only

<a id="params-8f5791"></a>&#x2022; [`query_params`](#params-8f5791) - Optional Block<br>HTTP Query Parameters. A list of predicates for all query parameters that need to be matched. The criteria for matching each query parameter are described in individual instances of QueryParameterMatcherType. The actual query parameter values are extracted from the request API as a list of strings for each query parameter name. Note that all specified query parameter predicates must evaluate to true<br>See [Query Params](#params-8f5791) below.

<a id="label-a84f6e"></a>&#x2022; [`undefined_flow_label`](#label-a84f6e) - Optional Block<br>Enable this option

<a id="web-a33d3d"></a>&#x2022; [`web`](#web-a33d3d) - Optional Block<br>Enable this option

<a id="mobile-0ffdfb"></a>&#x2022; [`web_mobile`](#mobile-0ffdfb) - Optional Block<br>Web and Mobile traffic type. Web and Mobile traffic type<br>See [Web Mobile](#mobile-0ffdfb) below.

#### Bot Defense Policy Protected App Endpoints Domain

A [`domain`](#domain-18bf1a) block (within [`bot_defense.policy.protected_app_endpoints`](#endpoints-01a2f3)) supports the following:

<a id="value-d5b836"></a>&#x2022; [`exact_value`](#value-d5b836) - Optional String<br>Exact Value. Exact domain name

<a id="value-4e4756"></a>&#x2022; [`regex_value`](#value-4e4756) - Optional String<br>Regex Values of Domains. Regular Expression value for the domain name

<a id="value-89654e"></a>&#x2022; [`suffix_value`](#value-89654e) - Optional String<br>Suffix Value. Suffix of domain name e.g 'xyz.com' will match '*.xyz.com' and 'xyz.com'

#### Bot Defense Policy Protected App Endpoints Flow Label

A [`flow_label`](#label-244fef) block (within [`bot_defense.policy.protected_app_endpoints`](#endpoints-01a2f3)) supports the following:

<a id="management-d237e9"></a>&#x2022; [`account_management`](#management-d237e9) - Optional Block<br>Bot Defense Flow Label Account Management Category. Bot Defense Flow Label Account Management Category<br>See [Account Management](#management-d237e9) below.

<a id="authentication-60331f"></a>&#x2022; [`authentication`](#authentication-60331f) - Optional Block<br>Bot Defense Flow Label Authentication Category. Bot Defense Flow Label Authentication Category<br>See [Authentication](#authentication-60331f) below.

<a id="services-acd29e"></a>&#x2022; [`financial_services`](#services-acd29e) - Optional Block<br>Bot Defense Flow Label Financial Services Category. Bot Defense Flow Label Financial Services Category<br>See [Financial Services](#services-acd29e) below.

<a id="flight-0c8cf6"></a>&#x2022; [`flight`](#flight-0c8cf6) - Optional Block<br>Bot Defense Flow Label Flight Category. Bot Defense Flow Label Flight Category<br>See [Flight](#flight-0c8cf6) below.

<a id="management-9be6b5"></a>&#x2022; [`profile_management`](#management-9be6b5) - Optional Block<br>Bot Defense Flow Label Profile Management Category. Bot Defense Flow Label Profile Management Category<br>See [Profile Management](#management-9be6b5) below.

<a id="search-d60360"></a>&#x2022; [`search`](#search-d60360) - Optional Block<br>Bot Defense Flow Label Search Category. Bot Defense Flow Label Search Category<br>See [Search](#search-d60360) below.

<a id="cards-f10d47"></a>&#x2022; [`shopping_gift_cards`](#cards-f10d47) - Optional Block<br>Bot Defense Flow Label Shopping & Gift Cards Category. Bot Defense Flow Label Shopping & Gift Cards Category<br>See [Shopping Gift Cards](#cards-f10d47) below.

#### Bot Defense Policy Protected App Endpoints Flow Label Account Management

An [`account_management`](#management-d237e9) block (within [`bot_defense.policy.protected_app_endpoints.flow_label`](#label-244fef)) supports the following:

<a id="create-c8c685"></a>&#x2022; [`create`](#create-c8c685) - Optional Block<br>Enable this option

<a id="reset-862ec4"></a>&#x2022; [`password_reset`](#reset-862ec4) - Optional Block<br>Enable this option

#### Bot Defense Policy Protected App Endpoints Flow Label Authentication

An [`authentication`](#authentication-60331f) block (within [`bot_defense.policy.protected_app_endpoints.flow_label`](#label-244fef)) supports the following:

<a id="login-79d094"></a>&#x2022; [`login`](#login-79d094) - Optional Block<br>Bot Defense Transaction Result. Bot Defense Transaction Result<br>See [Login](#login-79d094) below.

<a id="mfa-43e4fe"></a>&#x2022; [`login_mfa`](#mfa-43e4fe) - Optional Block<br>Enable this option

<a id="partner-c06e70"></a>&#x2022; [`login_partner`](#partner-c06e70) - Optional Block<br>Enable this option

<a id="logout-01c637"></a>&#x2022; [`logout`](#logout-01c637) - Optional Block<br>Enable this option

<a id="refresh-89934b"></a>&#x2022; [`token_refresh`](#refresh-89934b) - Optional Block<br>Enable this option

#### Bot Defense Policy Protected App Endpoints Flow Label Authentication Login

A [`login`](#login-79d094) block (within [`bot_defense.policy.protected_app_endpoints.flow_label.authentication`](#authentication-60331f)) supports the following:

<a id="result-60e1f0"></a>&#x2022; [`disable_transaction_result`](#result-60e1f0) - Optional Block<br>Enable this option

<a id="result-c2e927"></a>&#x2022; [`transaction_result`](#result-c2e927) - Optional Block<br>Bot Defense Transaction Result Type. Bot Defense Transaction ResultType<br>See [Transaction Result](#result-c2e927) below.

#### Bot Defense Policy Protected App Endpoints Flow Label Authentication Login Transaction Result

A [`transaction_result`](#result-c2e927) block (within [`bot_defense.policy.protected_app_endpoints.flow_label.authentication.login`](#login-79d094)) supports the following:

<a id="conditions-86ce87"></a>&#x2022; [`failure_conditions`](#conditions-86ce87) - Optional Block<br>Failure Conditions. Failure Conditions<br>See [Failure Conditions](#conditions-86ce87) below.

<a id="conditions-0b5152"></a>&#x2022; [`success_conditions`](#conditions-0b5152) - Optional Block<br>Success Conditions. Success Conditions<br>See [Success Conditions](#conditions-0b5152) below.

#### Bot Defense Policy Protected App Endpoints Flow Label Authentication Login Transaction Result Failure Conditions

A [`failure_conditions`](#conditions-86ce87) block (within [`bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result`](#result-c2e927)) supports the following:

<a id="name-eed9c2"></a>&#x2022; [`name`](#name-eed9c2) - Optional String<br>Header Name. A case-insensitive HTTP header name

<a id="values-15e71a"></a>&#x2022; [`regex_values`](#values-15e71a) - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="status-b492f7"></a>&#x2022; [`status`](#status-b492f7) - Optional String  Defaults to `EmptyStatusCode`<br>Possible values are `EmptyStatusCode`, `Continue`, `OK`, `Created`, `Accepted`, `NonAuthoritativeInformation`, `NoContent`, `ResetContent`, `PartialContent`, `MultiStatus`, `AlreadyReported`, `IMUsed`, `MultipleChoices`, `MovedPermanently`, `Found`, `SeeOther`, `NotModified`, `UseProxy`, `TemporaryRedirect`, `PermanentRedirect`, `BadRequest`, `Unauthorized`, `PaymentRequired`, `Forbidden`, `NotFound`, `MethodNotAllowed`, `NotAcceptable`, `ProxyAuthenticationRequired`, `RequestTimeout`, `Conflict`, `Gone`, `LengthRequired`, `PreconditionFailed`, `PayloadTooLarge`, `URITooLong`, `UnsupportedMediaType`, `RangeNotSatisfiable`, `ExpectationFailed`, `MisdirectedRequest`, `UnprocessableEntity`, `Locked`, `FailedDependency`, `UpgradeRequired`, `PreconditionRequired`, `TooManyRequests`, `RequestHeaderFieldsTooLarge`, `InternalServerError`, `NotImplemented`, `BadGateway`, `ServiceUnavailable`, `GatewayTimeout`, `HTTPVersionNotSupported`, `VariantAlsoNegotiates`, `InsufficientStorage`, `LoopDetected`, `NotExtended`, `NetworkAuthenticationRequired`<br>[Enum: EmptyStatusCode|Continue|OK|Created|Accepted|NonAuthoritativeInformation|NoContent|ResetContent|PartialContent|MultiStatus|AlreadyReported|IMUsed|MultipleChoices|MovedPermanently|Found|SeeOther|NotModified|UseProxy|TemporaryRedirect|PermanentRedirect|BadRequest|Unauthorized|PaymentRequired|Forbidden|NotFound|MethodNotAllowed|NotAcceptable|ProxyAuthenticationRequired|RequestTimeout|Conflict|Gone|LengthRequired|PreconditionFailed|PayloadTooLarge|URITooLong|UnsupportedMediaType|RangeNotSatisfiable|ExpectationFailed|MisdirectedRequest|UnprocessableEntity|Locked|FailedDependency|UpgradeRequired|PreconditionRequired|TooManyRequests|RequestHeaderFieldsTooLarge|InternalServerError|NotImplemented|BadGateway|ServiceUnavailable|GatewayTimeout|HTTPVersionNotSupported|VariantAlsoNegotiates|InsufficientStorage|LoopDetected|NotExtended|NetworkAuthenticationRequired] HTTP Status Code. HTTP response status codes EmptyStatusCode response codes means it is not specified Continue status code OK status code Created status code Accepted status code Non Authoritative Information status code No Content status code Reset Content status code Partial Content status code Multi Status status code Already Reported status code Im Used status code Multiple Choices status code Moved Permanently status code Found status code See Other status code Not Modified status code Use Proxy status code Temporary Redirect status code Permanent Redirect status code Bad Request status code Unauthorized status code Payment Required status code Forbidden status code Not Found status code Method Not Allowed status code Not Acceptable status code Proxy Authentication Required status code Request Timeout status code Conflict status code Gone status code Length Required status code Precondition Failed status code Payload Too Large status code URI Too Long status code Unsupported Media Type status code Range Not Satisfiable status code Expectation Failed status code Misdirected Request status code Unprocessable Entity status code Locked status code Failed Dependency status code Upgrade Required status code Precondition Required status code Too Many Requests status code Request Header Fields Too Large status code Internal Server Error status code Not Implemented status code Bad Gateway status code Service Unavailable status code Gateway Timeout status code HTTP Version Not Supported status code Variant Also Negotiates status code Insufficient Storage status code Loop Detected status code Not Extended status code Network Authentication Required status code

#### Bot Defense Policy Protected App Endpoints Flow Label Authentication Login Transaction Result Success Conditions

A [`success_conditions`](#conditions-0b5152) block (within [`bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result`](#result-c2e927)) supports the following:

<a id="name-98fa41"></a>&#x2022; [`name`](#name-98fa41) - Optional String<br>Header Name. A case-insensitive HTTP header name

<a id="values-9be26f"></a>&#x2022; [`regex_values`](#values-9be26f) - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="status-c08615"></a>&#x2022; [`status`](#status-c08615) - Optional String  Defaults to `EmptyStatusCode`<br>Possible values are `EmptyStatusCode`, `Continue`, `OK`, `Created`, `Accepted`, `NonAuthoritativeInformation`, `NoContent`, `ResetContent`, `PartialContent`, `MultiStatus`, `AlreadyReported`, `IMUsed`, `MultipleChoices`, `MovedPermanently`, `Found`, `SeeOther`, `NotModified`, `UseProxy`, `TemporaryRedirect`, `PermanentRedirect`, `BadRequest`, `Unauthorized`, `PaymentRequired`, `Forbidden`, `NotFound`, `MethodNotAllowed`, `NotAcceptable`, `ProxyAuthenticationRequired`, `RequestTimeout`, `Conflict`, `Gone`, `LengthRequired`, `PreconditionFailed`, `PayloadTooLarge`, `URITooLong`, `UnsupportedMediaType`, `RangeNotSatisfiable`, `ExpectationFailed`, `MisdirectedRequest`, `UnprocessableEntity`, `Locked`, `FailedDependency`, `UpgradeRequired`, `PreconditionRequired`, `TooManyRequests`, `RequestHeaderFieldsTooLarge`, `InternalServerError`, `NotImplemented`, `BadGateway`, `ServiceUnavailable`, `GatewayTimeout`, `HTTPVersionNotSupported`, `VariantAlsoNegotiates`, `InsufficientStorage`, `LoopDetected`, `NotExtended`, `NetworkAuthenticationRequired`<br>[Enum: EmptyStatusCode|Continue|OK|Created|Accepted|NonAuthoritativeInformation|NoContent|ResetContent|PartialContent|MultiStatus|AlreadyReported|IMUsed|MultipleChoices|MovedPermanently|Found|SeeOther|NotModified|UseProxy|TemporaryRedirect|PermanentRedirect|BadRequest|Unauthorized|PaymentRequired|Forbidden|NotFound|MethodNotAllowed|NotAcceptable|ProxyAuthenticationRequired|RequestTimeout|Conflict|Gone|LengthRequired|PreconditionFailed|PayloadTooLarge|URITooLong|UnsupportedMediaType|RangeNotSatisfiable|ExpectationFailed|MisdirectedRequest|UnprocessableEntity|Locked|FailedDependency|UpgradeRequired|PreconditionRequired|TooManyRequests|RequestHeaderFieldsTooLarge|InternalServerError|NotImplemented|BadGateway|ServiceUnavailable|GatewayTimeout|HTTPVersionNotSupported|VariantAlsoNegotiates|InsufficientStorage|LoopDetected|NotExtended|NetworkAuthenticationRequired] HTTP Status Code. HTTP response status codes EmptyStatusCode response codes means it is not specified Continue status code OK status code Created status code Accepted status code Non Authoritative Information status code No Content status code Reset Content status code Partial Content status code Multi Status status code Already Reported status code Im Used status code Multiple Choices status code Moved Permanently status code Found status code See Other status code Not Modified status code Use Proxy status code Temporary Redirect status code Permanent Redirect status code Bad Request status code Unauthorized status code Payment Required status code Forbidden status code Not Found status code Method Not Allowed status code Not Acceptable status code Proxy Authentication Required status code Request Timeout status code Conflict status code Gone status code Length Required status code Precondition Failed status code Payload Too Large status code URI Too Long status code Unsupported Media Type status code Range Not Satisfiable status code Expectation Failed status code Misdirected Request status code Unprocessable Entity status code Locked status code Failed Dependency status code Upgrade Required status code Precondition Required status code Too Many Requests status code Request Header Fields Too Large status code Internal Server Error status code Not Implemented status code Bad Gateway status code Service Unavailable status code Gateway Timeout status code HTTP Version Not Supported status code Variant Also Negotiates status code Insufficient Storage status code Loop Detected status code Not Extended status code Network Authentication Required status code

#### Bot Defense Policy Protected App Endpoints Flow Label Financial Services

A [`financial_services`](#services-acd29e) block (within [`bot_defense.policy.protected_app_endpoints.flow_label`](#label-244fef)) supports the following:

<a id="apply-9bb5b8"></a>&#x2022; [`apply`](#apply-9bb5b8) - Optional Block<br>Enable this option

<a id="transfer-ec9dc0"></a>&#x2022; [`money_transfer`](#transfer-ec9dc0) - Optional Block<br>Enable this option

#### Bot Defense Policy Protected App Endpoints Flow Label Flight

A [`flight`](#flight-0c8cf6) block (within [`bot_defense.policy.protected_app_endpoints.flow_label`](#label-244fef)) supports the following:

<a id="checkin-f1e656"></a>&#x2022; [`checkin`](#checkin-f1e656) - Optional Block<br>Enable this option

#### Bot Defense Policy Protected App Endpoints Flow Label Profile Management

A [`profile_management`](#management-9be6b5) block (within [`bot_defense.policy.protected_app_endpoints.flow_label`](#label-244fef)) supports the following:

<a id="create-63ed29"></a>&#x2022; [`create`](#create-63ed29) - Optional Block<br>Enable this option

<a id="update-c7e26c"></a>&#x2022; [`update`](#update-c7e26c) - Optional Block<br>Enable this option

<a id="view-2c180c"></a>&#x2022; [`view`](#view-2c180c) - Optional Block<br>Enable this option

#### Bot Defense Policy Protected App Endpoints Flow Label Search

A [`search`](#search-d60360) block (within [`bot_defense.policy.protected_app_endpoints.flow_label`](#label-244fef)) supports the following:

<a id="search-de8571"></a>&#x2022; [`flight_search`](#search-de8571) - Optional Block<br>Enable this option

<a id="search-389b2b"></a>&#x2022; [`product_search`](#search-389b2b) - Optional Block<br>Enable this option

<a id="search-0f9951"></a>&#x2022; [`reservation_search`](#search-0f9951) - Optional Block<br>Enable this option

<a id="search-3917b2"></a>&#x2022; [`room_search`](#search-3917b2) - Optional Block<br>Enable this option

#### Bot Defense Policy Protected App Endpoints Flow Label Shopping Gift Cards

A [`shopping_gift_cards`](#cards-f10d47) block (within [`bot_defense.policy.protected_app_endpoints.flow_label`](#label-244fef)) supports the following:

<a id="card-a0f859"></a>&#x2022; [`gift_card_make_purchase_with_gift_card`](#card-a0f859) - Optional Block<br>Enable this option

<a id="validation-1fa308"></a>&#x2022; [`gift_card_validation`](#validation-1fa308) - Optional Block<br>Enable this option

<a id="cart-6e5d88"></a>&#x2022; [`shop_add_to_cart`](#cart-6e5d88) - Optional Block<br>Enable this option

<a id="checkout-e4a04e"></a>&#x2022; [`shop_checkout`](#checkout-e4a04e) - Optional Block<br>Enable this option

<a id="seat-cb52da"></a>&#x2022; [`shop_choose_seat`](#seat-cb52da) - Optional Block<br>Enable this option

<a id="submission-f6e144"></a>&#x2022; [`shop_enter_drawing_submission`](#submission-f6e144) - Optional Block<br>Enable this option

<a id="payment-d25ab7"></a>&#x2022; [`shop_make_payment`](#payment-d25ab7) - Optional Block<br>Enable this option

<a id="order-c19bbe"></a>&#x2022; [`shop_order`](#order-c19bbe) - Optional Block<br>Enable this option

<a id="inquiry-9ca8c2"></a>&#x2022; [`shop_price_inquiry`](#inquiry-9ca8c2) - Optional Block<br>Enable this option

<a id="validation-b02840"></a>&#x2022; [`shop_promo_code_validation`](#validation-b02840) - Optional Block<br>Enable this option

<a id="card-2cf94e"></a>&#x2022; [`shop_purchase_gift_card`](#card-2cf94e) - Optional Block<br>Enable this option

<a id="quantity-4339b1"></a>&#x2022; [`shop_update_quantity`](#quantity-4339b1) - Optional Block<br>Enable this option

#### Bot Defense Policy Protected App Endpoints Headers

A [`headers`](#headers-986193) block (within [`bot_defense.policy.protected_app_endpoints`](#endpoints-01a2f3)) supports the following:

<a id="present-2e9857"></a>&#x2022; [`check_not_present`](#present-2e9857) - Optional Block<br>Enable this option

<a id="present-3a1075"></a>&#x2022; [`check_present`](#present-3a1075) - Optional Block<br>Enable this option

<a id="matcher-66fb69"></a>&#x2022; [`invert_matcher`](#matcher-66fb69) - Optional Bool<br>Invert Header Matcher. Invert the match result

<a id="item-ca0df2"></a>&#x2022; [`item`](#item-ca0df2) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#item-ca0df2) below.

<a id="name-34d16a"></a>&#x2022; [`name`](#name-34d16a) - Optional String<br>Header Name. A case-insensitive HTTP header name

#### Bot Defense Policy Protected App Endpoints Headers Item

An [`item`](#item-ca0df2) block (within [`bot_defense.policy.protected_app_endpoints.headers`](#headers-986193)) supports the following:

<a id="values-13c944"></a>&#x2022; [`exact_values`](#values-13c944) - Optional List<br>Exact Values. A list of exact values to match the input against

<a id="values-b0d727"></a>&#x2022; [`regex_values`](#values-b0d727) - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="transformers-eb2f5b"></a>&#x2022; [`transformers`](#transformers-eb2f5b) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>[Enum: LOWER_CASE|UPPER_CASE|BASE64_DECODE|NORMALIZE_PATH|REMOVE_WHITESPACE|URL_DECODE|TRIM_LEFT|TRIM_RIGHT|TRIM] Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### Bot Defense Policy Protected App Endpoints Metadata

A [`metadata`](#metadata-c93137) block (within [`bot_defense.policy.protected_app_endpoints`](#endpoints-01a2f3)) supports the following:

<a id="spec-e767de"></a>&#x2022; [`description_spec`](#spec-e767de) - Optional String<br>Description. Human readable description

<a id="name-324216"></a>&#x2022; [`name`](#name-324216) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### Bot Defense Policy Protected App Endpoints Mitigation

A [`mitigation`](#mitigation-cc96eb) block (within [`bot_defense.policy.protected_app_endpoints`](#endpoints-01a2f3)) supports the following:

<a id="block-d25b81"></a>&#x2022; [`block`](#block-d25b81) - Optional Block<br>Block bot mitigation. Block request and respond with custom content<br>See [Block](#block-d25b81) below.

<a id="flag-50a52b"></a>&#x2022; [`flag`](#flag-50a52b) - Optional Block<br>Select Flag Bot Mitigation Action. Flag mitigation action<br>See [Flag](#flag-50a52b) below.

<a id="redirect-2c8f41"></a>&#x2022; [`redirect`](#redirect-2c8f41) - Optional Block<br>Redirect bot mitigation. Redirect request to a custom URI<br>See [Redirect](#redirect-2c8f41) below.

#### Bot Defense Policy Protected App Endpoints Mitigation Block

A [`block`](#block-d25b81) block (within [`bot_defense.policy.protected_app_endpoints.mitigation`](#mitigation-cc96eb)) supports the following:

<a id="body-fe6d39"></a>&#x2022; [`body`](#body-fe6d39) - Optional String<br>Body. Custom body message is of type uri_ref. Currently supported URL schemes is string:///. For string:/// scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. 'Your request was blocked' or it can be HTML paragraph or a body string encoded as base64 string E.g. '`<p>` Your request was blocked `</p>`'. Base64 encoded string for this HTML is 'LzxwPiBZb3VyIHJlcXVlc3Qgd2FzIGJsb2NrZWQgPC9wPg=='

<a id="status-590093"></a>&#x2022; [`status`](#status-590093) - Optional String  Defaults to `EmptyStatusCode`<br>Possible values are `EmptyStatusCode`, `Continue`, `OK`, `Created`, `Accepted`, `NonAuthoritativeInformation`, `NoContent`, `ResetContent`, `PartialContent`, `MultiStatus`, `AlreadyReported`, `IMUsed`, `MultipleChoices`, `MovedPermanently`, `Found`, `SeeOther`, `NotModified`, `UseProxy`, `TemporaryRedirect`, `PermanentRedirect`, `BadRequest`, `Unauthorized`, `PaymentRequired`, `Forbidden`, `NotFound`, `MethodNotAllowed`, `NotAcceptable`, `ProxyAuthenticationRequired`, `RequestTimeout`, `Conflict`, `Gone`, `LengthRequired`, `PreconditionFailed`, `PayloadTooLarge`, `URITooLong`, `UnsupportedMediaType`, `RangeNotSatisfiable`, `ExpectationFailed`, `MisdirectedRequest`, `UnprocessableEntity`, `Locked`, `FailedDependency`, `UpgradeRequired`, `PreconditionRequired`, `TooManyRequests`, `RequestHeaderFieldsTooLarge`, `InternalServerError`, `NotImplemented`, `BadGateway`, `ServiceUnavailable`, `GatewayTimeout`, `HTTPVersionNotSupported`, `VariantAlsoNegotiates`, `InsufficientStorage`, `LoopDetected`, `NotExtended`, `NetworkAuthenticationRequired`<br>[Enum: EmptyStatusCode|Continue|OK|Created|Accepted|NonAuthoritativeInformation|NoContent|ResetContent|PartialContent|MultiStatus|AlreadyReported|IMUsed|MultipleChoices|MovedPermanently|Found|SeeOther|NotModified|UseProxy|TemporaryRedirect|PermanentRedirect|BadRequest|Unauthorized|PaymentRequired|Forbidden|NotFound|MethodNotAllowed|NotAcceptable|ProxyAuthenticationRequired|RequestTimeout|Conflict|Gone|LengthRequired|PreconditionFailed|PayloadTooLarge|URITooLong|UnsupportedMediaType|RangeNotSatisfiable|ExpectationFailed|MisdirectedRequest|UnprocessableEntity|Locked|FailedDependency|UpgradeRequired|PreconditionRequired|TooManyRequests|RequestHeaderFieldsTooLarge|InternalServerError|NotImplemented|BadGateway|ServiceUnavailable|GatewayTimeout|HTTPVersionNotSupported|VariantAlsoNegotiates|InsufficientStorage|LoopDetected|NotExtended|NetworkAuthenticationRequired] HTTP Status Code. HTTP response status codes EmptyStatusCode response codes means it is not specified Continue status code OK status code Created status code Accepted status code Non Authoritative Information status code No Content status code Reset Content status code Partial Content status code Multi Status status code Already Reported status code Im Used status code Multiple Choices status code Moved Permanently status code Found status code See Other status code Not Modified status code Use Proxy status code Temporary Redirect status code Permanent Redirect status code Bad Request status code Unauthorized status code Payment Required status code Forbidden status code Not Found status code Method Not Allowed status code Not Acceptable status code Proxy Authentication Required status code Request Timeout status code Conflict status code Gone status code Length Required status code Precondition Failed status code Payload Too Large status code URI Too Long status code Unsupported Media Type status code Range Not Satisfiable status code Expectation Failed status code Misdirected Request status code Unprocessable Entity status code Locked status code Failed Dependency status code Upgrade Required status code Precondition Required status code Too Many Requests status code Request Header Fields Too Large status code Internal Server Error status code Not Implemented status code Bad Gateway status code Service Unavailable status code Gateway Timeout status code HTTP Version Not Supported status code Variant Also Negotiates status code Insufficient Storage status code Loop Detected status code Not Extended status code Network Authentication Required status code

#### Bot Defense Policy Protected App Endpoints Mitigation Flag

A [`flag`](#flag-50a52b) block (within [`bot_defense.policy.protected_app_endpoints.mitigation`](#mitigation-cc96eb)) supports the following:

<a id="headers-cba7f7"></a>&#x2022; [`append_headers`](#headers-cba7f7) - Optional Block<br>Append Flag Mitigation Headers. Append flag mitigation headers to forwarded request<br>See [Append Headers](#headers-cba7f7) below.

<a id="headers-25974d"></a>&#x2022; [`no_headers`](#headers-25974d) - Optional Block<br>Enable this option

#### Bot Defense Policy Protected App Endpoints Mitigation Flag Append Headers

An [`append_headers`](#headers-cba7f7) block (within [`bot_defense.policy.protected_app_endpoints.mitigation.flag`](#flag-50a52b)) supports the following:

<a id="name-c64f18"></a>&#x2022; [`auto_type_header_name`](#name-c64f18) - Optional String<br>Automation Type Header Name. A case-insensitive HTTP header name

<a id="name-66a056"></a>&#x2022; [`inference_header_name`](#name-66a056) - Optional String<br>Inference Header Name. A case-insensitive HTTP header name

#### Bot Defense Policy Protected App Endpoints Mitigation Redirect

A [`redirect`](#redirect-2c8f41) block (within [`bot_defense.policy.protected_app_endpoints.mitigation`](#mitigation-cc96eb)) supports the following:

<a id="uri-56c0f3"></a>&#x2022; [`uri`](#uri-56c0f3) - Optional String<br>URI. URI location for redirect may be relative or absolute

#### Bot Defense Policy Protected App Endpoints Path

A [`path`](#path-d5ee15) block (within [`bot_defense.policy.protected_app_endpoints`](#endpoints-01a2f3)) supports the following:

<a id="path-16664a"></a>&#x2022; [`path`](#path-16664a) - Optional String<br>Exact. Exact path value to match

<a id="prefix-5a090b"></a>&#x2022; [`prefix`](#prefix-5a090b) - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

<a id="regex-1a10e9"></a>&#x2022; [`regex`](#regex-1a10e9) - Optional String<br>Regex. Regular expression of path match (e.g. the value .* will match on all paths)

#### Bot Defense Policy Protected App Endpoints Query Params

A [`query_params`](#params-8f5791) block (within [`bot_defense.policy.protected_app_endpoints`](#endpoints-01a2f3)) supports the following:

<a id="present-966c69"></a>&#x2022; [`check_not_present`](#present-966c69) - Optional Block<br>Enable this option

<a id="present-f5250d"></a>&#x2022; [`check_present`](#present-f5250d) - Optional Block<br>Enable this option

<a id="matcher-6361dd"></a>&#x2022; [`invert_matcher`](#matcher-6361dd) - Optional Bool<br>Invert Query Parameter Matcher. Invert the match result

<a id="item-92c230"></a>&#x2022; [`item`](#item-92c230) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#item-92c230) below.

<a id="key-2452b1"></a>&#x2022; [`key`](#key-2452b1) - Optional String<br>Query Parameter Name. A case-sensitive HTTP query parameter name

#### Bot Defense Policy Protected App Endpoints Query Params Item

An [`item`](#item-92c230) block (within [`bot_defense.policy.protected_app_endpoints.query_params`](#params-8f5791)) supports the following:

<a id="values-b9ca65"></a>&#x2022; [`exact_values`](#values-b9ca65) - Optional List<br>Exact Values. A list of exact values to match the input against

<a id="values-836047"></a>&#x2022; [`regex_values`](#values-836047) - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="transformers-794fb2"></a>&#x2022; [`transformers`](#transformers-794fb2) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>[Enum: LOWER_CASE|UPPER_CASE|BASE64_DECODE|NORMALIZE_PATH|REMOVE_WHITESPACE|URL_DECODE|TRIM_LEFT|TRIM_RIGHT|TRIM] Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### Bot Defense Policy Protected App Endpoints Web Mobile

A [`web_mobile`](#mobile-0ffdfb) block (within [`bot_defense.policy.protected_app_endpoints`](#endpoints-01a2f3)) supports the following:

<a id="identifier-0e0f05"></a>&#x2022; [`mobile_identifier`](#identifier-0e0f05) - Optional String  Defaults to `HEADERS`<br>[Enum: HEADERS] Mobile Identifier. Mobile identifier type - HEADERS: Headers Headers. The only possible value is `HEADERS`

#### Bot Defense Advanced

A [`bot_defense_advanced`](#bot-defense-advanced) block supports the following:

<a id="bot-defense-advanced-disable-js-insert"></a>&#x2022; [`disable_js_insert`](#bot-defense-advanced-disable-js-insert) - Optional Block<br>Enable this option

<a id="bot-defense-advanced-disable-mobile-sdk"></a>&#x2022; [`disable_mobile_sdk`](#bot-defense-advanced-disable-mobile-sdk) - Optional Block<br>Enable this option

<a id="pages-27f3ea"></a>&#x2022; [`js_insert_all_pages`](#pages-27f3ea) - Optional Block<br>Insert Bot Defense JavaScript in All Pages. Insert Bot Defense JavaScript in all pages<br>See [Js Insert All Pages](#pages-27f3ea) below.

<a id="except-cd2acd"></a>&#x2022; [`js_insert_all_pages_except`](#except-cd2acd) - Optional Block<br>Insert JavaScript in All Pages with the Exceptions. Insert Bot Defense JavaScript in all pages with the exceptions<br>See [Js Insert All Pages Except](#except-cd2acd) below.

<a id="bot-defense-advanced-js-insertion-rules"></a>&#x2022; [`js_insertion_rules`](#bot-defense-advanced-js-insertion-rules) - Optional Block<br>JavaScript Custom Insertion Rules. This defines custom JavaScript insertion rules for Bot Defense Policy<br>See [Js Insertion Rules](#bot-defense-advanced-js-insertion-rules) below.

<a id="bot-defense-advanced-mobile"></a>&#x2022; [`mobile`](#bot-defense-advanced-mobile) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Mobile](#bot-defense-advanced-mobile) below.

<a id="bot-defense-advanced-mobile-sdk-config"></a>&#x2022; [`mobile_sdk_config`](#bot-defense-advanced-mobile-sdk-config) - Optional Block<br>Mobile Request Identifier Headers. Mobile Request Identifier Headers<br>See [Mobile Sdk Config](#bot-defense-advanced-mobile-sdk-config) below.

<a id="bot-defense-advanced-web"></a>&#x2022; [`web`](#bot-defense-advanced-web) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Web](#bot-defense-advanced-web) below.

#### Bot Defense Advanced Js Insert All Pages

A [`js_insert_all_pages`](#pages-27f3ea) block (within [`bot_defense_advanced`](#bot-defense-advanced)) supports the following:

<a id="location-f54ccc"></a>&#x2022; [`javascript_location`](#location-f54ccc) - Optional String  Defaults to `AFTER_HEAD`<br>Possible values are `AFTER_HEAD`, `AFTER_TITLE_END`, `BEFORE_SCRIPT`<br>[Enum: AFTER_HEAD|AFTER_TITLE_END|BEFORE_SCRIPT] JavaScript Location. All inside networks. Insert JavaScript after `<head>` tag Insert JavaScript after `</title>` tag. Insert JavaScript before first `<script>` tag

#### Bot Defense Advanced Js Insert All Pages Except

A [`js_insert_all_pages_except`](#except-cd2acd) block (within [`bot_defense_advanced`](#bot-defense-advanced)) supports the following:

<a id="list-e251bf"></a>&#x2022; [`exclude_list`](#list-e251bf) - Optional Block<br>Exclude Pages. Optional JavaScript insertions exclude list of domain and path matchers<br>See [Exclude List](#list-e251bf) below.

<a id="location-76cc97"></a>&#x2022; [`javascript_location`](#location-76cc97) - Optional String  Defaults to `AFTER_HEAD`<br>Possible values are `AFTER_HEAD`, `AFTER_TITLE_END`, `BEFORE_SCRIPT`<br>[Enum: AFTER_HEAD|AFTER_TITLE_END|BEFORE_SCRIPT] JavaScript Location. All inside networks. Insert JavaScript after `<head>` tag Insert JavaScript after `</title>` tag. Insert JavaScript before first `<script>` tag

#### Bot Defense Advanced Js Insert All Pages Except Exclude List

An [`exclude_list`](#list-e251bf) block (within [`bot_defense_advanced.js_insert_all_pages_except`](#except-cd2acd)) supports the following:

<a id="domain-6fbb02"></a>&#x2022; [`any_domain`](#domain-6fbb02) - Optional Block<br>Enable this option

<a id="domain-172e0e"></a>&#x2022; [`domain`](#domain-172e0e) - Optional Block<br>Domains. Domains names<br>See [Domain](#domain-172e0e) below.

<a id="metadata-9fea08"></a>&#x2022; [`metadata`](#metadata-9fea08) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#metadata-9fea08) below.

<a id="path-2abaf7"></a>&#x2022; [`path`](#path-2abaf7) - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#path-2abaf7) below.

#### Bot Defense Advanced Js Insert All Pages Except Exclude List Domain

A [`domain`](#domain-172e0e) block (within [`bot_defense_advanced.js_insert_all_pages_except.exclude_list`](#list-e251bf)) supports the following:

<a id="value-2ff9db"></a>&#x2022; [`exact_value`](#value-2ff9db) - Optional String<br>Exact Value. Exact domain name

<a id="value-449383"></a>&#x2022; [`regex_value`](#value-449383) - Optional String<br>Regex Values of Domains. Regular Expression value for the domain name

<a id="value-1ebcbf"></a>&#x2022; [`suffix_value`](#value-1ebcbf) - Optional String<br>Suffix Value. Suffix of domain name e.g 'xyz.com' will match '*.xyz.com' and 'xyz.com'

#### Bot Defense Advanced Js Insert All Pages Except Exclude List Metadata

A [`metadata`](#metadata-9fea08) block (within [`bot_defense_advanced.js_insert_all_pages_except.exclude_list`](#list-e251bf)) supports the following:

<a id="spec-465ab4"></a>&#x2022; [`description_spec`](#spec-465ab4) - Optional String<br>Description. Human readable description

<a id="name-0b1f86"></a>&#x2022; [`name`](#name-0b1f86) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### Bot Defense Advanced Js Insert All Pages Except Exclude List Path

A [`path`](#path-2abaf7) block (within [`bot_defense_advanced.js_insert_all_pages_except.exclude_list`](#list-e251bf)) supports the following:

<a id="path-b0e485"></a>&#x2022; [`path`](#path-b0e485) - Optional String<br>Exact. Exact path value to match

<a id="prefix-b0164b"></a>&#x2022; [`prefix`](#prefix-b0164b) - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

<a id="regex-09a490"></a>&#x2022; [`regex`](#regex-09a490) - Optional String<br>Regex. Regular expression of path match (e.g. the value .* will match on all paths)

#### Bot Defense Advanced Js Insertion Rules

A [`js_insertion_rules`](#bot-defense-advanced-js-insertion-rules) block (within [`bot_defense_advanced`](#bot-defense-advanced)) supports the following:

<a id="list-efb443"></a>&#x2022; [`exclude_list`](#list-efb443) - Optional Block<br>Exclude Paths. Optional JavaScript insertions exclude list of domain and path matchers<br>See [Exclude List](#list-efb443) below.

<a id="rules-24e5a0"></a>&#x2022; [`rules`](#rules-24e5a0) - Optional Block<br>JavaScript Insertions. Required list of pages to insert Bot Defense client JavaScript<br>See [Rules](#rules-24e5a0) below.

#### Bot Defense Advanced Js Insertion Rules Exclude List

An [`exclude_list`](#list-efb443) block (within [`bot_defense_advanced.js_insertion_rules`](#bot-defense-advanced-js-insertion-rules)) supports the following:

<a id="domain-d9459c"></a>&#x2022; [`any_domain`](#domain-d9459c) - Optional Block<br>Enable this option

<a id="domain-c7e6a5"></a>&#x2022; [`domain`](#domain-c7e6a5) - Optional Block<br>Domains. Domains names<br>See [Domain](#domain-c7e6a5) below.

<a id="metadata-62a4f8"></a>&#x2022; [`metadata`](#metadata-62a4f8) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#metadata-62a4f8) below.

<a id="path-a587a8"></a>&#x2022; [`path`](#path-a587a8) - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#path-a587a8) below.

#### Bot Defense Advanced Js Insertion Rules Exclude List Domain

A [`domain`](#domain-c7e6a5) block (within [`bot_defense_advanced.js_insertion_rules.exclude_list`](#list-efb443)) supports the following:

<a id="value-e2cd94"></a>&#x2022; [`exact_value`](#value-e2cd94) - Optional String<br>Exact Value. Exact domain name

<a id="value-a38a1e"></a>&#x2022; [`regex_value`](#value-a38a1e) - Optional String<br>Regex Values of Domains. Regular Expression value for the domain name

<a id="value-88bdc2"></a>&#x2022; [`suffix_value`](#value-88bdc2) - Optional String<br>Suffix Value. Suffix of domain name e.g 'xyz.com' will match '*.xyz.com' and 'xyz.com'

#### Bot Defense Advanced Js Insertion Rules Exclude List Metadata

A [`metadata`](#metadata-62a4f8) block (within [`bot_defense_advanced.js_insertion_rules.exclude_list`](#list-efb443)) supports the following:

<a id="spec-59d921"></a>&#x2022; [`description_spec`](#spec-59d921) - Optional String<br>Description. Human readable description

<a id="name-42d756"></a>&#x2022; [`name`](#name-42d756) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### Bot Defense Advanced Js Insertion Rules Exclude List Path

A [`path`](#path-a587a8) block (within [`bot_defense_advanced.js_insertion_rules.exclude_list`](#list-efb443)) supports the following:

<a id="path-45571c"></a>&#x2022; [`path`](#path-45571c) - Optional String<br>Exact. Exact path value to match

<a id="prefix-4ca2ff"></a>&#x2022; [`prefix`](#prefix-4ca2ff) - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

<a id="regex-a19933"></a>&#x2022; [`regex`](#regex-a19933) - Optional String<br>Regex. Regular expression of path match (e.g. the value .* will match on all paths)

#### Bot Defense Advanced Js Insertion Rules Rules

A [`rules`](#rules-24e5a0) block (within [`bot_defense_advanced.js_insertion_rules`](#bot-defense-advanced-js-insertion-rules)) supports the following:

<a id="domain-bd13eb"></a>&#x2022; [`any_domain`](#domain-bd13eb) - Optional Block<br>Enable this option

<a id="domain-ff2f2e"></a>&#x2022; [`domain`](#domain-ff2f2e) - Optional Block<br>Domains. Domains names<br>See [Domain](#domain-ff2f2e) below.

<a id="location-20f540"></a>&#x2022; [`javascript_location`](#location-20f540) - Optional String  Defaults to `AFTER_HEAD`<br>Possible values are `AFTER_HEAD`, `AFTER_TITLE_END`, `BEFORE_SCRIPT`<br>[Enum: AFTER_HEAD|AFTER_TITLE_END|BEFORE_SCRIPT] JavaScript Location. All inside networks. Insert JavaScript after `<head>` tag Insert JavaScript after `</title>` tag. Insert JavaScript before first `<script>` tag

<a id="metadata-43c6ee"></a>&#x2022; [`metadata`](#metadata-43c6ee) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#metadata-43c6ee) below.

<a id="path-a4408d"></a>&#x2022; [`path`](#path-a4408d) - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#path-a4408d) below.

#### Bot Defense Advanced Js Insertion Rules Rules Domain

A [`domain`](#domain-ff2f2e) block (within [`bot_defense_advanced.js_insertion_rules.rules`](#rules-24e5a0)) supports the following:

<a id="value-592e4c"></a>&#x2022; [`exact_value`](#value-592e4c) - Optional String<br>Exact Value. Exact domain name

<a id="value-14d019"></a>&#x2022; [`regex_value`](#value-14d019) - Optional String<br>Regex Values of Domains. Regular Expression value for the domain name

<a id="value-8745f0"></a>&#x2022; [`suffix_value`](#value-8745f0) - Optional String<br>Suffix Value. Suffix of domain name e.g 'xyz.com' will match '*.xyz.com' and 'xyz.com'

#### Bot Defense Advanced Js Insertion Rules Rules Metadata

A [`metadata`](#metadata-43c6ee) block (within [`bot_defense_advanced.js_insertion_rules.rules`](#rules-24e5a0)) supports the following:

<a id="spec-644a5a"></a>&#x2022; [`description_spec`](#spec-644a5a) - Optional String<br>Description. Human readable description

<a id="name-e019ea"></a>&#x2022; [`name`](#name-e019ea) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### Bot Defense Advanced Js Insertion Rules Rules Path

A [`path`](#path-a4408d) block (within [`bot_defense_advanced.js_insertion_rules.rules`](#rules-24e5a0)) supports the following:

<a id="path-d4bd7b"></a>&#x2022; [`path`](#path-d4bd7b) - Optional String<br>Exact. Exact path value to match

<a id="prefix-714e31"></a>&#x2022; [`prefix`](#prefix-714e31) - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

<a id="regex-dddc59"></a>&#x2022; [`regex`](#regex-dddc59) - Optional String<br>Regex. Regular expression of path match (e.g. the value .* will match on all paths)

#### Bot Defense Advanced Mobile

A [`mobile`](#bot-defense-advanced-mobile) block (within [`bot_defense_advanced`](#bot-defense-advanced)) supports the following:

<a id="bot-defense-advanced-mobile-name"></a>&#x2022; [`name`](#bot-defense-advanced-mobile-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="bot-defense-advanced-mobile-namespace"></a>&#x2022; [`namespace`](#bot-defense-advanced-mobile-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="bot-defense-advanced-mobile-tenant"></a>&#x2022; [`tenant`](#bot-defense-advanced-mobile-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Bot Defense Advanced Mobile Sdk Config

A [`mobile_sdk_config`](#bot-defense-advanced-mobile-sdk-config) block (within [`bot_defense_advanced`](#bot-defense-advanced)) supports the following:

<a id="identifier-163438"></a>&#x2022; [`mobile_identifier`](#identifier-163438) - Optional Block<br>Mobile Traffic Identifier. Mobile traffic identifier type<br>See [Mobile Identifier](#identifier-163438) below.

#### Bot Defense Advanced Mobile Sdk Config Mobile Identifier

A [`mobile_identifier`](#identifier-163438) block (within [`bot_defense_advanced.mobile_sdk_config`](#bot-defense-advanced-mobile-sdk-config)) supports the following:

<a id="headers-dd786e"></a>&#x2022; [`headers`](#headers-dd786e) - Optional Block<br>Headers. Headers that can be used to identify mobile traffic<br>See [Headers](#headers-dd786e) below.

#### Bot Defense Advanced Mobile Sdk Config Mobile Identifier Headers

A [`headers`](#headers-dd786e) block (within [`bot_defense_advanced.mobile_sdk_config.mobile_identifier`](#identifier-163438)) supports the following:

<a id="present-dfe97f"></a>&#x2022; [`check_not_present`](#present-dfe97f) - Optional Block<br>Enable this option

<a id="present-9a0d45"></a>&#x2022; [`check_present`](#present-9a0d45) - Optional Block<br>Enable this option

<a id="item-78ae4d"></a>&#x2022; [`item`](#item-78ae4d) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#item-78ae4d) below.

<a id="name-8f313e"></a>&#x2022; [`name`](#name-8f313e) - Optional String<br>Header Name. A case-insensitive HTTP header name

#### Bot Defense Advanced Mobile Sdk Config Mobile Identifier Headers Item

An [`item`](#item-78ae4d) block (within [`bot_defense_advanced.mobile_sdk_config.mobile_identifier.headers`](#headers-dd786e)) supports the following:

<a id="values-4ff98b"></a>&#x2022; [`exact_values`](#values-4ff98b) - Optional List<br>Exact Values. A list of exact values to match the input against

<a id="values-e3e723"></a>&#x2022; [`regex_values`](#values-e3e723) - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="transformers-e6f04e"></a>&#x2022; [`transformers`](#transformers-e6f04e) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>[Enum: LOWER_CASE|UPPER_CASE|BASE64_DECODE|NORMALIZE_PATH|REMOVE_WHITESPACE|URL_DECODE|TRIM_LEFT|TRIM_RIGHT|TRIM] Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### Bot Defense Advanced Web

A [`web`](#bot-defense-advanced-web) block (within [`bot_defense_advanced`](#bot-defense-advanced)) supports the following:

<a id="bot-defense-advanced-web-name"></a>&#x2022; [`name`](#bot-defense-advanced-web-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="bot-defense-advanced-web-namespace"></a>&#x2022; [`namespace`](#bot-defense-advanced-web-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="bot-defense-advanced-web-tenant"></a>&#x2022; [`tenant`](#bot-defense-advanced-web-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Caching Policy

A [`caching_policy`](#caching-policy) block supports the following:

<a id="caching-policy-custom-cache-rule"></a>&#x2022; [`custom_cache_rule`](#caching-policy-custom-cache-rule) - Optional Block<br>Custom Cache Rules. Caching policies for CDN<br>See [Custom Cache Rule](#caching-policy-custom-cache-rule) below.

<a id="caching-policy-default-cache-action"></a>&#x2022; [`default_cache_action`](#caching-policy-default-cache-action) - Optional Block<br>Default Cache Behaviour. This defines a Default Cache Action<br>See [Default Cache Action](#caching-policy-default-cache-action) below.

#### Caching Policy Custom Cache Rule

A [`custom_cache_rule`](#caching-policy-custom-cache-rule) block (within [`caching_policy`](#caching-policy)) supports the following:

<a id="rules-e10c80"></a>&#x2022; [`cdn_cache_rules`](#rules-e10c80) - Optional Block<br>CDN Cache Rule. Reference to CDN Cache Rule configuration object<br>See [CDN Cache Rules](#rules-e10c80) below.

#### Caching Policy Custom Cache Rule CDN Cache Rules

A [`cdn_cache_rules`](#rules-e10c80) block (within [`caching_policy.custom_cache_rule`](#caching-policy-custom-cache-rule)) supports the following:

<a id="name-60e3d9"></a>&#x2022; [`name`](#name-60e3d9) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-d73f50"></a>&#x2022; [`namespace`](#namespace-d73f50) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-6fd72f"></a>&#x2022; [`tenant`](#tenant-6fd72f) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Caching Policy Default Cache Action

A [`default_cache_action`](#caching-policy-default-cache-action) block (within [`caching_policy`](#caching-policy)) supports the following:

<a id="disabled-39759c"></a>&#x2022; [`cache_disabled`](#disabled-39759c) - Optional Block<br>Enable this option

<a id="default-2899d2"></a>&#x2022; [`cache_ttl_default`](#default-2899d2) - Optional String<br>Fallback Cache TTL (d/ h/ m). Use Cache TTL Provided by Origin, and set a contigency TTL value in case one is not provided

<a id="override-3c128f"></a>&#x2022; [`cache_ttl_override`](#override-3c128f) - Optional String<br>Override Cache TTL (d/ h/ m/ s). Always override the Cahce TTL provided by Origin

#### Captcha Challenge

A [`captcha_challenge`](#captcha-challenge) block supports the following:

<a id="captcha-challenge-cookie-expiry"></a>&#x2022; [`cookie_expiry`](#captcha-challenge-cookie-expiry) - Optional Number<br>Cookie Expiration Period. Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new challenge

<a id="captcha-challenge-custom-page"></a>&#x2022; [`custom_page`](#captcha-challenge-custom-page) - Optional String<br>Custom message for Captcha Challenge. Custom message is of type uri_ref. Currently supported URL schemes is string:///. For string:/// scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. 'Please Wait.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '`<p>` Please Wait `</p>`'. Base64 encoded string for this HTML is 'PHA+IFBsZWFzZSBXYWl0IDwvcD4='

#### Client Side Defense

A [`client_side_defense`](#client-side-defense) block supports the following:

<a id="client-side-defense-policy"></a>&#x2022; [`policy`](#client-side-defense-policy) - Optional Block<br>Client-Side Defense Policy. This defines various configuration options for Client-Side Defense policy<br>See [Policy](#client-side-defense-policy) below.

#### Client Side Defense Policy

A [`policy`](#client-side-defense-policy) block (within [`client_side_defense`](#client-side-defense)) supports the following:

<a id="insert-683e69"></a>&#x2022; [`disable_js_insert`](#insert-683e69) - Optional Block<br>Enable this option

<a id="pages-38bd1c"></a>&#x2022; [`js_insert_all_pages`](#pages-38bd1c) - Optional Block<br>Enable this option

<a id="except-7bfe85"></a>&#x2022; [`js_insert_all_pages_except`](#except-7bfe85) - Optional Block<br>Insert JavaScript in All Pages with the Exceptions. Insert Client-Side Defense JavaScript in all pages with the exceptions<br>See [Js Insert All Pages Except](#except-7bfe85) below.

<a id="rules-ad3671"></a>&#x2022; [`js_insertion_rules`](#rules-ad3671) - Optional Block<br>JavaScript Custom Insertion Rules. This defines custom JavaScript insertion rules for Client-Side Defense Policy<br>See [Js Insertion Rules](#rules-ad3671) below.

#### Client Side Defense Policy Js Insert All Pages Except

A [`js_insert_all_pages_except`](#except-7bfe85) block (within [`client_side_defense.policy`](#client-side-defense-policy)) supports the following:

<a id="list-fc1c50"></a>&#x2022; [`exclude_list`](#list-fc1c50) - Optional Block<br>Exclude Pages. Optional JavaScript insertions exclude list of domain and path matchers<br>See [Exclude List](#list-fc1c50) below.

#### Client Side Defense Policy Js Insert All Pages Except Exclude List

An [`exclude_list`](#list-fc1c50) block (within [`client_side_defense.policy.js_insert_all_pages_except`](#except-7bfe85)) supports the following:

<a id="domain-cfab55"></a>&#x2022; [`any_domain`](#domain-cfab55) - Optional Block<br>Enable this option

<a id="domain-15fe0c"></a>&#x2022; [`domain`](#domain-15fe0c) - Optional Block<br>Domains. Domains names<br>See [Domain](#domain-15fe0c) below.

<a id="metadata-50baa8"></a>&#x2022; [`metadata`](#metadata-50baa8) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#metadata-50baa8) below.

<a id="path-82c392"></a>&#x2022; [`path`](#path-82c392) - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#path-82c392) below.

#### Client Side Defense Policy Js Insert All Pages Except Exclude List Domain

A [`domain`](#domain-15fe0c) block (within [`client_side_defense.policy.js_insert_all_pages_except.exclude_list`](#list-fc1c50)) supports the following:

<a id="value-f9285b"></a>&#x2022; [`exact_value`](#value-f9285b) - Optional String<br>Exact Value. Exact domain name

<a id="value-4207c5"></a>&#x2022; [`regex_value`](#value-4207c5) - Optional String<br>Regex Values of Domains. Regular Expression value for the domain name

<a id="value-2a8824"></a>&#x2022; [`suffix_value`](#value-2a8824) - Optional String<br>Suffix Value. Suffix of domain name e.g 'xyz.com' will match '*.xyz.com' and 'xyz.com'

#### Client Side Defense Policy Js Insert All Pages Except Exclude List Metadata

A [`metadata`](#metadata-50baa8) block (within [`client_side_defense.policy.js_insert_all_pages_except.exclude_list`](#list-fc1c50)) supports the following:

<a id="spec-28351b"></a>&#x2022; [`description_spec`](#spec-28351b) - Optional String<br>Description. Human readable description

<a id="name-3e26d9"></a>&#x2022; [`name`](#name-3e26d9) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### Client Side Defense Policy Js Insert All Pages Except Exclude List Path

A [`path`](#path-82c392) block (within [`client_side_defense.policy.js_insert_all_pages_except.exclude_list`](#list-fc1c50)) supports the following:

<a id="path-390fcf"></a>&#x2022; [`path`](#path-390fcf) - Optional String<br>Exact. Exact path value to match

<a id="prefix-d5139b"></a>&#x2022; [`prefix`](#prefix-d5139b) - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

<a id="regex-0010b2"></a>&#x2022; [`regex`](#regex-0010b2) - Optional String<br>Regex. Regular expression of path match (e.g. the value .* will match on all paths)

#### Client Side Defense Policy Js Insertion Rules

A [`js_insertion_rules`](#rules-ad3671) block (within [`client_side_defense.policy`](#client-side-defense-policy)) supports the following:

<a id="list-dfecb6"></a>&#x2022; [`exclude_list`](#list-dfecb6) - Optional Block<br>Exclude Paths. Optional JavaScript insertions exclude list of domain and path matchers<br>See [Exclude List](#list-dfecb6) below.

<a id="rules-6276bc"></a>&#x2022; [`rules`](#rules-6276bc) - Optional Block<br>JavaScript Insertions. Required list of pages to insert Client-Side Defense client JavaScript<br>See [Rules](#rules-6276bc) below.

#### Client Side Defense Policy Js Insertion Rules Exclude List

An [`exclude_list`](#list-dfecb6) block (within [`client_side_defense.policy.js_insertion_rules`](#rules-ad3671)) supports the following:

<a id="domain-7b414f"></a>&#x2022; [`any_domain`](#domain-7b414f) - Optional Block<br>Enable this option

<a id="domain-7c7a7c"></a>&#x2022; [`domain`](#domain-7c7a7c) - Optional Block<br>Domains. Domains names<br>See [Domain](#domain-7c7a7c) below.

<a id="metadata-bd2353"></a>&#x2022; [`metadata`](#metadata-bd2353) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#metadata-bd2353) below.

<a id="path-962e59"></a>&#x2022; [`path`](#path-962e59) - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#path-962e59) below.

#### Client Side Defense Policy Js Insertion Rules Exclude List Domain

A [`domain`](#domain-7c7a7c) block (within [`client_side_defense.policy.js_insertion_rules.exclude_list`](#list-dfecb6)) supports the following:

<a id="value-f64365"></a>&#x2022; [`exact_value`](#value-f64365) - Optional String<br>Exact Value. Exact domain name

<a id="value-b28460"></a>&#x2022; [`regex_value`](#value-b28460) - Optional String<br>Regex Values of Domains. Regular Expression value for the domain name

<a id="value-9e64cb"></a>&#x2022; [`suffix_value`](#value-9e64cb) - Optional String<br>Suffix Value. Suffix of domain name e.g 'xyz.com' will match '*.xyz.com' and 'xyz.com'

#### Client Side Defense Policy Js Insertion Rules Exclude List Metadata

A [`metadata`](#metadata-bd2353) block (within [`client_side_defense.policy.js_insertion_rules.exclude_list`](#list-dfecb6)) supports the following:

<a id="spec-13fe6d"></a>&#x2022; [`description_spec`](#spec-13fe6d) - Optional String<br>Description. Human readable description

<a id="name-b1fc08"></a>&#x2022; [`name`](#name-b1fc08) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### Client Side Defense Policy Js Insertion Rules Exclude List Path

A [`path`](#path-962e59) block (within [`client_side_defense.policy.js_insertion_rules.exclude_list`](#list-dfecb6)) supports the following:

<a id="path-c659e3"></a>&#x2022; [`path`](#path-c659e3) - Optional String<br>Exact. Exact path value to match

<a id="prefix-d88a45"></a>&#x2022; [`prefix`](#prefix-d88a45) - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

<a id="regex-424dc4"></a>&#x2022; [`regex`](#regex-424dc4) - Optional String<br>Regex. Regular expression of path match (e.g. the value .* will match on all paths)

#### Client Side Defense Policy Js Insertion Rules Rules

A [`rules`](#rules-6276bc) block (within [`client_side_defense.policy.js_insertion_rules`](#rules-ad3671)) supports the following:

<a id="domain-a7aac7"></a>&#x2022; [`any_domain`](#domain-a7aac7) - Optional Block<br>Enable this option

<a id="domain-4b295f"></a>&#x2022; [`domain`](#domain-4b295f) - Optional Block<br>Domains. Domains names<br>See [Domain](#domain-4b295f) below.

<a id="metadata-60fc86"></a>&#x2022; [`metadata`](#metadata-60fc86) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#metadata-60fc86) below.

<a id="path-71b688"></a>&#x2022; [`path`](#path-71b688) - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#path-71b688) below.

#### Client Side Defense Policy Js Insertion Rules Rules Domain

A [`domain`](#domain-4b295f) block (within [`client_side_defense.policy.js_insertion_rules.rules`](#rules-6276bc)) supports the following:

<a id="value-2357e4"></a>&#x2022; [`exact_value`](#value-2357e4) - Optional String<br>Exact Value. Exact domain name

<a id="value-a4f4ae"></a>&#x2022; [`regex_value`](#value-a4f4ae) - Optional String<br>Regex Values of Domains. Regular Expression value for the domain name

<a id="value-a7894b"></a>&#x2022; [`suffix_value`](#value-a7894b) - Optional String<br>Suffix Value. Suffix of domain name e.g 'xyz.com' will match '*.xyz.com' and 'xyz.com'

#### Client Side Defense Policy Js Insertion Rules Rules Metadata

A [`metadata`](#metadata-60fc86) block (within [`client_side_defense.policy.js_insertion_rules.rules`](#rules-6276bc)) supports the following:

<a id="spec-f78d91"></a>&#x2022; [`description_spec`](#spec-f78d91) - Optional String<br>Description. Human readable description

<a id="name-1afae7"></a>&#x2022; [`name`](#name-1afae7) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### Client Side Defense Policy Js Insertion Rules Rules Path

A [`path`](#path-71b688) block (within [`client_side_defense.policy.js_insertion_rules.rules`](#rules-6276bc)) supports the following:

<a id="path-df67b6"></a>&#x2022; [`path`](#path-df67b6) - Optional String<br>Exact. Exact path value to match

<a id="prefix-7c1e21"></a>&#x2022; [`prefix`](#prefix-7c1e21) - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

<a id="regex-248fb3"></a>&#x2022; [`regex`](#regex-248fb3) - Optional String<br>Regex. Regular expression of path match (e.g. the value .* will match on all paths)

#### Cookie Stickiness

A [`cookie_stickiness`](#cookie-stickiness) block supports the following:

<a id="cookie-stickiness-add-httponly"></a>&#x2022; [`add_httponly`](#cookie-stickiness-add-httponly) - Optional Block<br>Enable this option

<a id="cookie-stickiness-add-secure"></a>&#x2022; [`add_secure`](#cookie-stickiness-add-secure) - Optional Block<br>Enable this option

<a id="cookie-stickiness-ignore-httponly"></a>&#x2022; [`ignore_httponly`](#cookie-stickiness-ignore-httponly) - Optional Block<br>Enable this option

<a id="cookie-stickiness-ignore-samesite"></a>&#x2022; [`ignore_samesite`](#cookie-stickiness-ignore-samesite) - Optional Block<br>Enable this option

<a id="cookie-stickiness-ignore-secure"></a>&#x2022; [`ignore_secure`](#cookie-stickiness-ignore-secure) - Optional Block<br>Enable this option

<a id="cookie-stickiness-name"></a>&#x2022; [`name`](#cookie-stickiness-name) - Optional String<br>Name. The name of the cookie that will be used to obtain the hash key. If the cookie is not present and TTL below is not set, no hash will be produced

<a id="cookie-stickiness-path"></a>&#x2022; [`path`](#cookie-stickiness-path) - Optional String<br>Path. The name of the path for the cookie. If no path is specified here, no path will be set for the cookie

<a id="cookie-stickiness-samesite-lax"></a>&#x2022; [`samesite_lax`](#cookie-stickiness-samesite-lax) - Optional Block<br>Enable this option

<a id="cookie-stickiness-samesite-none"></a>&#x2022; [`samesite_none`](#cookie-stickiness-samesite-none) - Optional Block<br>Enable this option

<a id="cookie-stickiness-samesite-strict"></a>&#x2022; [`samesite_strict`](#cookie-stickiness-samesite-strict) - Optional Block<br>Enable this option

<a id="cookie-stickiness-ttl"></a>&#x2022; [`ttl`](#cookie-stickiness-ttl) - Optional Number<br>TTL. If specified, a cookie with the TTL will be generated if the cookie is not present. If the TTL is present and zero, the generated cookie will be a session cookie. TTL value is in milliseconds

#### CORS Policy

A [`cors_policy`](#cors-policy) block supports the following:

<a id="cors-policy-allow-credentials"></a>&#x2022; [`allow_credentials`](#cors-policy-allow-credentials) - Optional Bool<br>Allow Credentials. Specifies whether the resource allows credentials

<a id="cors-policy-allow-headers"></a>&#x2022; [`allow_headers`](#cors-policy-allow-headers) - Optional String<br>Allow Headers. Specifies the content for the access-control-allow-headers header

<a id="cors-policy-allow-methods"></a>&#x2022; [`allow_methods`](#cors-policy-allow-methods) - Optional String<br>Allow Methods. Specifies the content for the access-control-allow-methods header

<a id="cors-policy-allow-origin"></a>&#x2022; [`allow_origin`](#cors-policy-allow-origin) - Optional List<br>Allow Origin. Specifies the origins that will be allowed to do CORS requests. An origin is allowed if either allow_origin or allow_origin_regex match

<a id="cors-policy-allow-origin-regex"></a>&#x2022; [`allow_origin_regex`](#cors-policy-allow-origin-regex) - Optional List<br>Allow Origin Regex. Specifies regex patterns that match allowed origins. An origin is allowed if either allow_origin or allow_origin_regex match

<a id="cors-policy-disabled"></a>&#x2022; [`disabled`](#cors-policy-disabled) - Optional Bool<br>Disabled. Disable the CorsPolicy for a particular route. This is useful when virtual-host has CorsPolicy, but we need to disable it on a specific route. The value of this field is ignored for virtual-host

<a id="cors-policy-expose-headers"></a>&#x2022; [`expose_headers`](#cors-policy-expose-headers) - Optional String<br>Expose Headers. Specifies the content for the access-control-expose-headers header

<a id="cors-policy-maximum-age"></a>&#x2022; [`maximum_age`](#cors-policy-maximum-age) - Optional Number<br>Maximum Age. Specifies the content for the access-control-max-age header in seconds. This indicates the maximum number of seconds the results can be cached A value of -1 will disable caching. Maximum permitted value is 86400 seconds (24 hours)

#### CSRF Policy

A [`csrf_policy`](#csrf-policy) block supports the following:

<a id="csrf-policy-all-load-balancer-domains"></a>&#x2022; [`all_load_balancer_domains`](#csrf-policy-all-load-balancer-domains) - Optional Block<br>Enable this option

<a id="csrf-policy-custom-domain-list"></a>&#x2022; [`custom_domain_list`](#csrf-policy-custom-domain-list) - Optional Block<br>Domain name list. List of domain names used for Host header matching<br>See [Custom Domain List](#csrf-policy-custom-domain-list) below.

<a id="csrf-policy-disabled"></a>&#x2022; [`disabled`](#csrf-policy-disabled) - Optional Block<br>Enable this option

#### CSRF Policy Custom Domain List

A [`custom_domain_list`](#csrf-policy-custom-domain-list) block (within [`csrf_policy`](#csrf-policy)) supports the following:

<a id="csrf-policy-custom-domain-list-domains"></a>&#x2022; [`domains`](#csrf-policy-custom-domain-list-domains) - Optional List<br>Domain names. A list of domain names that will be matched to loadbalancer. These domains are not used for SNI match. Wildcard names are supported in the suffix or prefix form

#### Data Guard Rules

A [`data_guard_rules`](#data-guard-rules) block supports the following:

<a id="data-guard-rules-any-domain"></a>&#x2022; [`any_domain`](#data-guard-rules-any-domain) - Optional Block<br>Enable this option

<a id="data-guard-rules-apply-data-guard"></a>&#x2022; [`apply_data_guard`](#data-guard-rules-apply-data-guard) - Optional Block<br>Enable this option

<a id="data-guard-rules-exact-value"></a>&#x2022; [`exact_value`](#data-guard-rules-exact-value) - Optional String<br>Exact Value. Exact domain name

<a id="data-guard-rules-metadata"></a>&#x2022; [`metadata`](#data-guard-rules-metadata) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#data-guard-rules-metadata) below.

<a id="data-guard-rules-path"></a>&#x2022; [`path`](#data-guard-rules-path) - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#data-guard-rules-path) below.

<a id="data-guard-rules-skip-data-guard"></a>&#x2022; [`skip_data_guard`](#data-guard-rules-skip-data-guard) - Optional Block<br>Enable this option

<a id="data-guard-rules-suffix-value"></a>&#x2022; [`suffix_value`](#data-guard-rules-suffix-value) - Optional String<br>Suffix Value. Suffix of domain name e.g 'xyz.com' will match '*.xyz.com' and 'xyz.com'

#### Data Guard Rules Metadata

A [`metadata`](#data-guard-rules-metadata) block (within [`data_guard_rules`](#data-guard-rules)) supports the following:

<a id="spec-bca77c"></a>&#x2022; [`description_spec`](#spec-bca77c) - Optional String<br>Description. Human readable description

<a id="data-guard-rules-metadata-name"></a>&#x2022; [`name`](#data-guard-rules-metadata-name) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### Data Guard Rules Path

A [`path`](#data-guard-rules-path) block (within [`data_guard_rules`](#data-guard-rules)) supports the following:

<a id="data-guard-rules-path-path"></a>&#x2022; [`path`](#data-guard-rules-path-path) - Optional String<br>Exact. Exact path value to match

<a id="data-guard-rules-path-prefix"></a>&#x2022; [`prefix`](#data-guard-rules-path-prefix) - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

<a id="data-guard-rules-path-regex"></a>&#x2022; [`regex`](#data-guard-rules-path-regex) - Optional String<br>Regex. Regular expression of path match (e.g. the value .* will match on all paths)

#### DDOS Mitigation Rules

A [`ddos_mitigation_rules`](#ddos-mitigation-rules) block supports the following:

<a id="ddos-mitigation-rules-block"></a>&#x2022; [`block`](#ddos-mitigation-rules-block) - Optional Block<br>Enable this option

<a id="source-02aa55"></a>&#x2022; [`ddos_client_source`](#source-02aa55) - Optional Block<br>DDOS Client Source Choice. DDOS Mitigation sources to be blocked<br>See [DDOS Client Source](#source-02aa55) below.

<a id="timestamp-bd6f49"></a>&#x2022; [`expiration_timestamp`](#timestamp-bd6f49) - Optional String<br>Expiration Timestamp. The expiration_timestamp is the RFC 3339 format timestamp at which the containing rule is considered to be logically expired. The rule continues to exist in the configuration but is not applied anymore

<a id="ddos-mitigation-rules-ip-prefix-list"></a>&#x2022; [`ip_prefix_list`](#ddos-mitigation-rules-ip-prefix-list) - Optional Block<br>IP Prefix Match List. List of IP Prefix strings to match against<br>See [IP Prefix List](#ddos-mitigation-rules-ip-prefix-list) below.

<a id="ddos-mitigation-rules-metadata"></a>&#x2022; [`metadata`](#ddos-mitigation-rules-metadata) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#ddos-mitigation-rules-metadata) below.

#### DDOS Mitigation Rules DDOS Client Source

A [`ddos_client_source`](#source-02aa55) block (within [`ddos_mitigation_rules`](#ddos-mitigation-rules)) supports the following:

<a id="list-20cb78"></a>&#x2022; [`asn_list`](#list-20cb78) - Optional Block<br>ASN Match List. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer<br>See [Asn List](#list-20cb78) below.

<a id="list-78d261"></a>&#x2022; [`country_list`](#list-78d261) - Optional List  Defaults to `COUNTRY_NONE`<br>Possible values are `COUNTRY_NONE`, `COUNTRY_AD`, `COUNTRY_AE`, `COUNTRY_AF`, `COUNTRY_AG`, `COUNTRY_AI`, `COUNTRY_AL`, `COUNTRY_AM`, `COUNTRY_AN`, `COUNTRY_AO`, `COUNTRY_AQ`, `COUNTRY_AR`, `COUNTRY_AS`, `COUNTRY_AT`, `COUNTRY_AU`, `COUNTRY_AW`, `COUNTRY_AX`, `COUNTRY_AZ`, `COUNTRY_BA`, `COUNTRY_BB`, `COUNTRY_BD`, `COUNTRY_BE`, `COUNTRY_BF`, `COUNTRY_BG`, `COUNTRY_BH`, `COUNTRY_BI`, `COUNTRY_BJ`, `COUNTRY_BL`, `COUNTRY_BM`, `COUNTRY_BN`, `COUNTRY_BO`, `COUNTRY_BQ`, `COUNTRY_BR`, `COUNTRY_BS`, `COUNTRY_BT`, `COUNTRY_BV`, `COUNTRY_BW`, `COUNTRY_BY`, `COUNTRY_BZ`, `COUNTRY_CA`, `COUNTRY_CC`, `COUNTRY_CD`, `COUNTRY_CF`, `COUNTRY_CG`, `COUNTRY_CH`, `COUNTRY_CI`, `COUNTRY_CK`, `COUNTRY_CL`, `COUNTRY_CM`, `COUNTRY_CN`, `COUNTRY_CO`, `COUNTRY_CR`, `COUNTRY_CS`, `COUNTRY_CU`, `COUNTRY_CV`, `COUNTRY_CW`, `COUNTRY_CX`, `COUNTRY_CY`, `COUNTRY_CZ`, `COUNTRY_DE`, `COUNTRY_DJ`, `COUNTRY_DK`, `COUNTRY_DM`, `COUNTRY_DO`, `COUNTRY_DZ`, `COUNTRY_EC`, `COUNTRY_EE`, `COUNTRY_EG`, `COUNTRY_EH`, `COUNTRY_ER`, `COUNTRY_ES`, `COUNTRY_ET`, `COUNTRY_FI`, `COUNTRY_FJ`, `COUNTRY_FK`, `COUNTRY_FM`, `COUNTRY_FO`, `COUNTRY_FR`, `COUNTRY_GA`, `COUNTRY_GB`, `COUNTRY_GD`, `COUNTRY_GE`, `COUNTRY_GF`, `COUNTRY_GG`, `COUNTRY_GH`, `COUNTRY_GI`, `COUNTRY_GL`, `COUNTRY_GM`, `COUNTRY_GN`, `COUNTRY_GP`, `COUNTRY_GQ`, `COUNTRY_GR`, `COUNTRY_GS`, `COUNTRY_GT`, `COUNTRY_GU`, `COUNTRY_GW`, `COUNTRY_GY`, `COUNTRY_HK`, `COUNTRY_HM`, `COUNTRY_HN`, `COUNTRY_HR`, `COUNTRY_HT`, `COUNTRY_HU`, `COUNTRY_ID`, `COUNTRY_IE`, `COUNTRY_IL`, `COUNTRY_IM`, `COUNTRY_IN`, `COUNTRY_IO`, `COUNTRY_IQ`, `COUNTRY_IR`, `COUNTRY_IS`, `COUNTRY_IT`, `COUNTRY_JE`, `COUNTRY_JM`, `COUNTRY_JO`, `COUNTRY_JP`, `COUNTRY_KE`, `COUNTRY_KG`, `COUNTRY_KH`, `COUNTRY_KI`, `COUNTRY_KM`, `COUNTRY_KN`, `COUNTRY_KP`, `COUNTRY_KR`, `COUNTRY_KW`, `COUNTRY_KY`, `COUNTRY_KZ`, `COUNTRY_LA`, `COUNTRY_LB`, `COUNTRY_LC`, `COUNTRY_LI`, `COUNTRY_LK`, `COUNTRY_LR`, `COUNTRY_LS`, `COUNTRY_LT`, `COUNTRY_LU`, `COUNTRY_LV`, `COUNTRY_LY`, `COUNTRY_MA`, `COUNTRY_MC`, `COUNTRY_MD`, `COUNTRY_ME`, `COUNTRY_MF`, `COUNTRY_MG`, `COUNTRY_MH`, `COUNTRY_MK`, `COUNTRY_ML`, `COUNTRY_MM`, `COUNTRY_MN`, `COUNTRY_MO`, `COUNTRY_MP`, `COUNTRY_MQ`, `COUNTRY_MR`, `COUNTRY_MS`, `COUNTRY_MT`, `COUNTRY_MU`, `COUNTRY_MV`, `COUNTRY_MW`, `COUNTRY_MX`, `COUNTRY_MY`, `COUNTRY_MZ`, `COUNTRY_NA`, `COUNTRY_NC`, `COUNTRY_NE`, `COUNTRY_NF`, `COUNTRY_NG`, `COUNTRY_NI`, `COUNTRY_NL`, `COUNTRY_NO`, `COUNTRY_NP`, `COUNTRY_NR`, `COUNTRY_NU`, `COUNTRY_NZ`, `COUNTRY_OM`, `COUNTRY_PA`, `COUNTRY_PE`, `COUNTRY_PF`, `COUNTRY_PG`, `COUNTRY_PH`, `COUNTRY_PK`, `COUNTRY_PL`, `COUNTRY_PM`, `COUNTRY_PN`, `COUNTRY_PR`, `COUNTRY_PS`, `COUNTRY_PT`, `COUNTRY_PW`, `COUNTRY_PY`, `COUNTRY_QA`, `COUNTRY_RE`, `COUNTRY_RO`, `COUNTRY_RS`, `COUNTRY_RU`, `COUNTRY_RW`, `COUNTRY_SA`, `COUNTRY_SB`, `COUNTRY_SC`, `COUNTRY_SD`, `COUNTRY_SE`, `COUNTRY_SG`, `COUNTRY_SH`, `COUNTRY_SI`, `COUNTRY_SJ`, `COUNTRY_SK`, `COUNTRY_SL`, `COUNTRY_SM`, `COUNTRY_SN`, `COUNTRY_SO`, `COUNTRY_SR`, `COUNTRY_SS`, `COUNTRY_ST`, `COUNTRY_SV`, `COUNTRY_SX`, `COUNTRY_SY`, `COUNTRY_SZ`, `COUNTRY_TC`, `COUNTRY_TD`, `COUNTRY_TF`, `COUNTRY_TG`, `COUNTRY_TH`, `COUNTRY_TJ`, `COUNTRY_TK`, `COUNTRY_TL`, `COUNTRY_TM`, `COUNTRY_TN`, `COUNTRY_TO`, `COUNTRY_TR`, `COUNTRY_TT`, `COUNTRY_TV`, `COUNTRY_TW`, `COUNTRY_TZ`, `COUNTRY_UA`, `COUNTRY_UG`, `COUNTRY_UM`, `COUNTRY_US`, `COUNTRY_UY`, `COUNTRY_UZ`, `COUNTRY_VA`, `COUNTRY_VC`, `COUNTRY_VE`, `COUNTRY_VG`, `COUNTRY_VI`, `COUNTRY_VN`, `COUNTRY_VU`, `COUNTRY_WF`, `COUNTRY_WS`, `COUNTRY_XK`, `COUNTRY_XT`, `COUNTRY_YE`, `COUNTRY_YT`, `COUNTRY_ZA`, `COUNTRY_ZM`, `COUNTRY_ZW`<br>[Enum: COUNTRY_NONE|COUNTRY_AD|COUNTRY_AE|COUNTRY_AF|COUNTRY_AG|COUNTRY_AI|COUNTRY_AL|COUNTRY_AM|COUNTRY_AN|COUNTRY_AO|COUNTRY_AQ|COUNTRY_AR|COUNTRY_AS|COUNTRY_AT|COUNTRY_AU|COUNTRY_AW|COUNTRY_AX|COUNTRY_AZ|COUNTRY_BA|COUNTRY_BB|COUNTRY_BD|COUNTRY_BE|COUNTRY_BF|COUNTRY_BG|COUNTRY_BH|COUNTRY_BI|COUNTRY_BJ|COUNTRY_BL|COUNTRY_BM|COUNTRY_BN|COUNTRY_BO|COUNTRY_BQ|COUNTRY_BR|COUNTRY_BS|COUNTRY_BT|COUNTRY_BV|COUNTRY_BW|COUNTRY_BY|COUNTRY_BZ|COUNTRY_CA|COUNTRY_CC|COUNTRY_CD|COUNTRY_CF|COUNTRY_CG|COUNTRY_CH|COUNTRY_CI|COUNTRY_CK|COUNTRY_CL|COUNTRY_CM|COUNTRY_CN|COUNTRY_CO|COUNTRY_CR|COUNTRY_CS|COUNTRY_CU|COUNTRY_CV|COUNTRY_CW|COUNTRY_CX|COUNTRY_CY|COUNTRY_CZ|COUNTRY_DE|COUNTRY_DJ|COUNTRY_DK|COUNTRY_DM|COUNTRY_DO|COUNTRY_DZ|COUNTRY_EC|COUNTRY_EE|COUNTRY_EG|COUNTRY_EH|COUNTRY_ER|COUNTRY_ES|COUNTRY_ET|COUNTRY_FI|COUNTRY_FJ|COUNTRY_FK|COUNTRY_FM|COUNTRY_FO|COUNTRY_FR|COUNTRY_GA|COUNTRY_GB|COUNTRY_GD|COUNTRY_GE|COUNTRY_GF|COUNTRY_GG|COUNTRY_GH|COUNTRY_GI|COUNTRY_GL|COUNTRY_GM|COUNTRY_GN|COUNTRY_GP|COUNTRY_GQ|COUNTRY_GR|COUNTRY_GS|COUNTRY_GT|COUNTRY_GU|COUNTRY_GW|COUNTRY_GY|COUNTRY_HK|COUNTRY_HM|COUNTRY_HN|COUNTRY_HR|COUNTRY_HT|COUNTRY_HU|COUNTRY_ID|COUNTRY_IE|COUNTRY_IL|COUNTRY_IM|COUNTRY_IN|COUNTRY_IO|COUNTRY_IQ|COUNTRY_IR|COUNTRY_IS|COUNTRY_IT|COUNTRY_JE|COUNTRY_JM|COUNTRY_JO|COUNTRY_JP|COUNTRY_KE|COUNTRY_KG|COUNTRY_KH|COUNTRY_KI|COUNTRY_KM|COUNTRY_KN|COUNTRY_KP|COUNTRY_KR|COUNTRY_KW|COUNTRY_KY|COUNTRY_KZ|COUNTRY_LA|COUNTRY_LB|COUNTRY_LC|COUNTRY_LI|COUNTRY_LK|COUNTRY_LR|COUNTRY_LS|COUNTRY_LT|COUNTRY_LU|COUNTRY_LV|COUNTRY_LY|COUNTRY_MA|COUNTRY_MC|COUNTRY_MD|COUNTRY_ME|COUNTRY_MF|COUNTRY_MG|COUNTRY_MH|COUNTRY_MK|COUNTRY_ML|COUNTRY_MM|COUNTRY_MN|COUNTRY_MO|COUNTRY_MP|COUNTRY_MQ|COUNTRY_MR|COUNTRY_MS|COUNTRY_MT|COUNTRY_MU|COUNTRY_MV|COUNTRY_MW|COUNTRY_MX|COUNTRY_MY|COUNTRY_MZ|COUNTRY_NA|COUNTRY_NC|COUNTRY_NE|COUNTRY_NF|COUNTRY_NG|COUNTRY_NI|COUNTRY_NL|COUNTRY_NO|COUNTRY_NP|COUNTRY_NR|COUNTRY_NU|COUNTRY_NZ|COUNTRY_OM|COUNTRY_PA|COUNTRY_PE|COUNTRY_PF|COUNTRY_PG|COUNTRY_PH|COUNTRY_PK|COUNTRY_PL|COUNTRY_PM|COUNTRY_PN|COUNTRY_PR|COUNTRY_PS|COUNTRY_PT|COUNTRY_PW|COUNTRY_PY|COUNTRY_QA|COUNTRY_RE|COUNTRY_RO|COUNTRY_RS|COUNTRY_RU|COUNTRY_RW|COUNTRY_SA|COUNTRY_SB|COUNTRY_SC|COUNTRY_SD|COUNTRY_SE|COUNTRY_SG|COUNTRY_SH|COUNTRY_SI|COUNTRY_SJ|COUNTRY_SK|COUNTRY_SL|COUNTRY_SM|COUNTRY_SN|COUNTRY_SO|COUNTRY_SR|COUNTRY_SS|COUNTRY_ST|COUNTRY_SV|COUNTRY_SX|COUNTRY_SY|COUNTRY_SZ|COUNTRY_TC|COUNTRY_TD|COUNTRY_TF|COUNTRY_TG|COUNTRY_TH|COUNTRY_TJ|COUNTRY_TK|COUNTRY_TL|COUNTRY_TM|COUNTRY_TN|COUNTRY_TO|COUNTRY_TR|COUNTRY_TT|COUNTRY_TV|COUNTRY_TW|COUNTRY_TZ|COUNTRY_UA|COUNTRY_UG|COUNTRY_UM|COUNTRY_US|COUNTRY_UY|COUNTRY_UZ|COUNTRY_VA|COUNTRY_VC|COUNTRY_VE|COUNTRY_VG|COUNTRY_VI|COUNTRY_VN|COUNTRY_VU|COUNTRY_WF|COUNTRY_WS|COUNTRY_XK|COUNTRY_XT|COUNTRY_YE|COUNTRY_YT|COUNTRY_ZA|COUNTRY_ZM|COUNTRY_ZW] Country List. Sources that are located in one of the countries in the given list

<a id="matcher-a7a10e"></a>&#x2022; [`ja4_tls_fingerprint_matcher`](#matcher-a7a10e) - Optional Block<br>JA4 TLS Fingerprint Matcher. An extended version of JA3 that includes additional fields for more comprehensive fingerprinting of SSL/TLS clients and potentially has a different structure and length<br>See [Ja4 TLS Fingerprint Matcher](#matcher-a7a10e) below.

<a id="matcher-d4dd17"></a>&#x2022; [`tls_fingerprint_matcher`](#matcher-d4dd17) - Optional Block<br>TLS Fingerprint Matcher. A TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of supported positve match criteria includes a list of known classes of TLS fingerprints and a list of exact values. The match is considered successful if either of these positive criteria are satisfied and the input fingerprint is not one of the excluded values<br>See [TLS Fingerprint Matcher](#matcher-d4dd17) below.

#### DDOS Mitigation Rules DDOS Client Source Asn List

An [`asn_list`](#list-20cb78) block (within [`ddos_mitigation_rules.ddos_client_source`](#source-02aa55)) supports the following:

<a id="numbers-a2cb60"></a>&#x2022; [`as_numbers`](#numbers-a2cb60) - Optional List<br>AS Numbers. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer

#### DDOS Mitigation Rules DDOS Client Source Ja4 TLS Fingerprint Matcher

A [`ja4_tls_fingerprint_matcher`](#matcher-a7a10e) block (within [`ddos_mitigation_rules.ddos_client_source`](#source-02aa55)) supports the following:

<a id="values-d0a266"></a>&#x2022; [`exact_values`](#values-d0a266) - Optional List<br>Exact Values. A list of exact JA4 TLS fingerprint to match the input JA4 TLS fingerprint against

#### DDOS Mitigation Rules DDOS Client Source TLS Fingerprint Matcher

A [`tls_fingerprint_matcher`](#matcher-d4dd17) block (within [`ddos_mitigation_rules.ddos_client_source`](#source-02aa55)) supports the following:

<a id="classes-b8db1d"></a>&#x2022; [`classes`](#classes-b8db1d) - Optional List  Defaults to `TLS_FINGERPRINT_NONE`<br>Possible values are `TLS_FINGERPRINT_NONE`, `ANY_MALICIOUS_FINGERPRINT`, `ADWARE`, `ADWIND`, `DRIDEX`, `GOOTKIT`, `GOZI`, `JBIFROST`, `QUAKBOT`, `RANSOMWARE`, `TROLDESH`, `TOFSEE`, `TORRENTLOCKER`, `TRICKBOT`<br>[Enum: TLS_FINGERPRINT_NONE|ANY_MALICIOUS_FINGERPRINT|ADWARE|ADWIND|DRIDEX|GOOTKIT|GOZI|JBIFROST|QUAKBOT|RANSOMWARE|TROLDESH|TOFSEE|TORRENTLOCKER|TRICKBOT] TLS fingerprint classes. A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against

<a id="values-2a6e72"></a>&#x2022; [`exact_values`](#values-2a6e72) - Optional List<br>Exact Values. A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against

<a id="values-7f5411"></a>&#x2022; [`excluded_values`](#values-7f5411) - Optional List<br>Excluded Values. A list of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can be used to skip known false positives when using one or more known TLS fingerprint classes in the enclosing matcher

#### DDOS Mitigation Rules IP Prefix List

An [`ip_prefix_list`](#ddos-mitigation-rules-ip-prefix-list) block (within [`ddos_mitigation_rules`](#ddos-mitigation-rules)) supports the following:

<a id="match-f3f64f"></a>&#x2022; [`invert_match`](#match-f3f64f) - Optional Bool<br>Invert Match Result. Invert the match result

<a id="prefixes-aed245"></a>&#x2022; [`ip_prefixes`](#prefixes-aed245) - Optional List<br>IPv4 Prefix List. List of IPv4 prefix strings

#### DDOS Mitigation Rules Metadata

A [`metadata`](#ddos-mitigation-rules-metadata) block (within [`ddos_mitigation_rules`](#ddos-mitigation-rules)) supports the following:

<a id="spec-f95573"></a>&#x2022; [`description_spec`](#spec-f95573) - Optional String<br>Description. Human readable description

<a id="ddos-mitigation-rules-metadata-name"></a>&#x2022; [`name`](#ddos-mitigation-rules-metadata-name) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### Default Pool

A [`default_pool`](#default-pool) block supports the following:

<a id="default-pool-advanced-options"></a>&#x2022; [`advanced_options`](#default-pool-advanced-options) - Optional Block<br>Origin Pool Advanced Options. Configure Advanced options for origin pool<br>See [Advanced Options](#default-pool-advanced-options) below.

<a id="default-pool-automatic-port"></a>&#x2022; [`automatic_port`](#default-pool-automatic-port) - Optional Block<br>Enable this option

<a id="default-pool-endpoint-selection"></a>&#x2022; [`endpoint_selection`](#default-pool-endpoint-selection) - Optional String  Defaults to `DISTRIBUTED`<br>Possible values are `DISTRIBUTED`, `LOCAL_ONLY`, `LOCAL_PREFERRED`<br>[Enum: DISTRIBUTED|LOCAL_ONLY|LOCAL_PREFERRED] Endpoint Selection Policy. Policy for selection of endpoints from local site/remote site/both Consider both remote and local endpoints for load balancing LOCAL_ONLY: Consider only local endpoints for load balancing Enable this policy to load balance ONLY among locally discovered endpoints Prefer the local endpoints for load balancing. If local endpoints are not present remote endpoints will be considered

<a id="default-pool-health-check-port"></a>&#x2022; [`health_check_port`](#default-pool-health-check-port) - Optional Number<br>Health check port. Port used for performing health check

<a id="default-pool-healthcheck"></a>&#x2022; [`healthcheck`](#default-pool-healthcheck) - Optional Block<br>Health Check object. Reference to healthcheck configuration objects<br>See [Healthcheck](#default-pool-healthcheck) below.

<a id="default-pool-lb-port"></a>&#x2022; [`lb_port`](#default-pool-lb-port) - Optional Block<br>Enable this option

<a id="default-pool-loadbalancer-algorithm"></a>&#x2022; [`loadbalancer_algorithm`](#default-pool-loadbalancer-algorithm) - Optional String  Defaults to `ROUND_ROBIN`<br>Possible values are `ROUND_ROBIN`, `LEAST_REQUEST`, `RING_HASH`, `RANDOM`, `LB_OVERRIDE`<br>[Enum: ROUND_ROBIN|LEAST_REQUEST|RING_HASH|RANDOM|LB_OVERRIDE] Load Balancer Algorithm. Different load balancing algorithms supported When a connection to a endpoint in an upstream cluster is required, the load balancer uses loadbalancer_algorithm to determine which host is selected. - ROUND_ROBIN: ROUND_ROBIN Policy in which each healthy/available upstream endpoint is selected in round robin order. - LEAST_REQUEST: LEAST_REQUEST Policy in which loadbalancer picks the upstream endpoint which has the fewest active requests - RING_HASH: RING_HASH Policy implements consistent hashing to upstream endpoints using ring hash of endpoint names Hash of the incoming request is calculated using request hash policy. The ring/modulo hash load balancer implements consistent hashing to upstream hosts. The algorithm is based on mapping all hosts onto a circle such that the addition or removal of a host from the host set changes only affect 1/N requests. This technique is also commonly known as “ketama” hashing. A consistent hashing load balancer is only effective when protocol routing is used that specifies a value to hash on. The minimum ring size governs the replication factor for each host in the ring. For example, if the minimum ring size is 1024 and there are 16 hosts, each host will be replicated 64 times. - RANDOM: RANDOM Policy in which each available upstream endpoint is selected in random order. The random load balancer selects a random healthy host. The random load balancer generally performs better than round robin if no health checking policy is configured. Random selection avoids bias towards the host in the set that comes after a failed host. - LB_OVERRIDE: Load Balancer Override Hash policy is taken from from the load balancer which is using this origin pool

<a id="default-pool-no-tls"></a>&#x2022; [`no_tls`](#default-pool-no-tls) - Optional Block<br>Enable this option

<a id="default-pool-origin-servers"></a>&#x2022; [`origin_servers`](#default-pool-origin-servers) - Optional Block<br>Origin Servers. List of origin servers in this pool<br>See [Origin Servers](#default-pool-origin-servers) below.

<a id="default-pool-port"></a>&#x2022; [`port`](#default-pool-port) - Optional Number<br>Port. Endpoint service is available on this port

<a id="default-pool-same-as-endpoint-port"></a>&#x2022; [`same_as_endpoint_port`](#default-pool-same-as-endpoint-port) - Optional Block<br>Enable this option

<a id="type-2756f7"></a>&#x2022; [`upstream_conn_pool_reuse_type`](#type-2756f7) - Optional Block<br>Select upstream connection pool reuse state. Select upstream connection pool reuse state for every downstream connection. This configuration choice is for HTTP(S) LB only<br>See [Upstream Conn Pool Reuse Type](#type-2756f7) below.

<a id="default-pool-use-tls"></a>&#x2022; [`use_tls`](#default-pool-use-tls) - Optional Block<br>TLS Parameters for Origin Servers. Upstream TLS Parameters<br>See [Use TLS](#default-pool-use-tls) below.

<a id="default-pool-view-internal"></a>&#x2022; [`view_internal`](#default-pool-view-internal) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [View Internal](#default-pool-view-internal) below.

#### Default Pool Advanced Options

An [`advanced_options`](#default-pool-advanced-options) block (within [`default_pool`](#default-pool)) supports the following:

<a id="config-48f56b"></a>&#x2022; [`auto_http_config`](#config-48f56b) - Optional Block<br>Enable this option

<a id="breaker-8f5df4"></a>&#x2022; [`circuit_breaker`](#breaker-8f5df4) - Optional Block<br>Circuit Breaker. CircuitBreaker provides a mechanism for watching failures in upstream connections or requests and if the failures reach a certain threshold, automatically fail subsequent requests which allows to apply back pressure on downstream quickly<br>See [Circuit Breaker](#breaker-8f5df4) below.

<a id="timeout-8cd873"></a>&#x2022; [`connection_timeout`](#timeout-8cd873) - Optional Number  Defaults to `2`  Specified in milliseconds<br>Connection Timeout. The timeout for new network connections to endpoints in the cluster.  The seconds

<a id="breaker-db5d25"></a>&#x2022; [`default_circuit_breaker`](#breaker-db5d25) - Optional Block<br>Enable this option

<a id="breaker-03c951"></a>&#x2022; [`disable_circuit_breaker`](#breaker-03c951) - Optional Block<br>Enable this option

<a id="persistance-83b4c4"></a>&#x2022; [`disable_lb_source_ip_persistance`](#persistance-83b4c4) - Optional Block<br>Enable this option

<a id="detection-46546c"></a>&#x2022; [`disable_outlier_detection`](#detection-46546c) - Optional Block<br>Enable this option

<a id="protocol-d614b9"></a>&#x2022; [`disable_proxy_protocol`](#protocol-d614b9) - Optional Block<br>Enable this option

<a id="subsets-b0bd38"></a>&#x2022; [`disable_subsets`](#subsets-b0bd38) - Optional Block<br>Enable this option

<a id="persistance-d799ee"></a>&#x2022; [`enable_lb_source_ip_persistance`](#persistance-d799ee) - Optional Block<br>Enable this option

<a id="subsets-5a741c"></a>&#x2022; [`enable_subsets`](#subsets-5a741c) - Optional Block<br>Origin Pool Subset Options. Configure subset options for origin pool<br>See [Enable Subsets](#subsets-5a741c) below.

<a id="config-a0bc3c"></a>&#x2022; [`http1_config`](#config-a0bc3c) - Optional Block<br>HTTP/1.1 Protocol Options. HTTP/1.1 Protocol options for upstream connections<br>See [Http1 Config](#config-a0bc3c) below.

<a id="options-fc9fd8"></a>&#x2022; [`http2_options`](#options-fc9fd8) - Optional Block<br>Http2 Protocol Options. Http2 Protocol options for upstream connections<br>See [Http2 Options](#options-fc9fd8) below.

<a id="timeout-3807d9"></a>&#x2022; [`http_idle_timeout`](#timeout-3807d9) - Optional Number  Defaults to `5`  Specified in milliseconds<br>HTTP Idle Timeout. The idle timeout for upstream connection pool connections. The idle timeout is defined as the period in which there are no active requests. When the idle timeout is reached the connection will be closed. Note that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive.  The minutes

<a id="threshold-4ef07a"></a>&#x2022; [`no_panic_threshold`](#threshold-4ef07a) - Optional Block<br>Enable this option

<a id="detection-c89e70"></a>&#x2022; [`outlier_detection`](#detection-c89e70) - Optional Block<br>Outlier Detection. Outlier detection and ejection is the process of dynamically determining whether some number of hosts in an upstream cluster are performing unlike the others and removing them from the healthy load balancing set. Outlier detection is a form of passive health checking. Algorithm 1. A endpoint is determined to be an outlier (based on configured number of consecutive_5xx or consecutive_gateway_failures) . 2. If no endpoints have been ejected, loadbalancer will eject the host immediately. Otherwise, it checks to make sure the number of ejected hosts is below the allowed threshold (specified via max_ejection_percent setting). If the number of ejected hosts is above the threshold, the host is not ejected. 3. The endpoint is ejected for some number of milliseconds. Ejection means that the endpoint is marked unhealthy and will not be used during load balancing. The number of milliseconds is equal to the base_ejection_time value multiplied by the number of times the host has been ejected. 4. An ejected endpoint will automatically be brought back into service after the ejection time has been satisfied<br>See [Outlier Detection](#detection-c89e70) below.

<a id="threshold-61a03f"></a>&#x2022; [`panic_threshold`](#threshold-61a03f) - Optional Number<br>Panic threshold. Configure a threshold (percentage of unhealthy endpoints) below which all endpoints will be considered for load balancing ignoring its health status

<a id="protocol-v1-de8613"></a>&#x2022; [`proxy_protocol_v1`](#protocol-v1-de8613) - Optional Block<br>Enable this option

<a id="protocol-v2-5b9d69"></a>&#x2022; [`proxy_protocol_v2`](#protocol-v2-5b9d69) - Optional Block<br>Enable this option

#### Default Pool Advanced Options Circuit Breaker

A [`circuit_breaker`](#breaker-8f5df4) block (within [`default_pool.advanced_options`](#default-pool-advanced-options)) supports the following:

<a id="limit-866c11"></a>&#x2022; [`connection_limit`](#limit-866c11) - Optional Number<br>Connection Limit. The maximum number of connections that loadbalancer will establish to all hosts in an upstream cluster. In practice this is only applicable to TCP and HTTP/1.1 clusters since HTTP/2 uses a single connection to each host. Remove endpoint out of load balancing decision, if number of connections reach connection limit

<a id="requests-ec6e61"></a>&#x2022; [`max_requests`](#requests-ec6e61) - Optional Number<br>Maximum Request Count. The maximum number of requests that can be outstanding to all hosts in a cluster at any given time. In practice this is applicable to HTTP/2 clusters since HTTP/1.1 clusters are governed by the maximum connections (connection_limit). Remove endpoint out of load balancing decision, if requests exceed this count

<a id="requests-d174a1"></a>&#x2022; [`pending_requests`](#requests-d174a1) - Optional Number<br>Pending Requests. The maximum number of requests that will be queued while waiting for a ready connection pool connection. Since HTTP/2 requests are sent over a single connection, this circuit breaker only comes into play as the initial connection is created, as requests will be multiplexed immediately afterwards. For HTTP/1.1, requests are added to the list of pending requests whenever there aren’t enough upstream connections available to immediately dispatch the request, so this circuit breaker will remain in play for the lifetime of the process. Remove endpoint out of load balancing decision, if pending request reach pending_request

<a id="priority-80164b"></a>&#x2022; [`priority`](#priority-80164b) - Optional String  Defaults to `DEFAULT`<br>Possible values are `DEFAULT`, `HIGH`<br>[Enum: DEFAULT|HIGH] Routing Priority. Priority routing for each request. Different connection pools are used based on the priority selected for the request. Also, circuit-breaker configuration at destination cluster is chosen based on selected priority. Default routing mechanism High-Priority routing mechanism

<a id="retries-a3c574"></a>&#x2022; [`retries`](#retries-a3c574) - Optional Number<br>Retry Count. The maximum number of retries that can be outstanding to all hosts in a cluster at any given time. Remove endpoint out of load balancing decision, if retries for request exceed this count

#### Default Pool Advanced Options Enable Subsets

An [`enable_subsets`](#subsets-5a741c) block (within [`default_pool.advanced_options`](#default-pool-advanced-options)) supports the following:

<a id="endpoint-6dd4c1"></a>&#x2022; [`any_endpoint`](#endpoint-6dd4c1) - Optional Block<br>Enable this option

<a id="subset-1bf539"></a>&#x2022; [`default_subset`](#subset-1bf539) - Optional Block<br>Origin Pool Default Subset. Default Subset definition<br>See [Default Subset](#subset-1bf539) below.

<a id="subsets-5171e5"></a>&#x2022; [`endpoint_subsets`](#subsets-5171e5) - Optional Block<br>Origin Server Subsets Classes. List of subset class. Subsets class is defined using list of keys. Every unique combination of values of these keys form a subset withing the class<br>See [Endpoint Subsets](#subsets-5171e5) below.

<a id="request-26d37d"></a>&#x2022; [`fail_request`](#request-26d37d) - Optional Block<br>Enable this option

#### Default Pool Advanced Options Enable Subsets Default Subset

A [`default_subset`](#subset-1bf539) block (within [`default_pool.advanced_options.enable_subsets`](#subsets-5a741c)) supports the following:

<a id="subset-c104f5"></a>&#x2022; [`default_subset`](#subset-c104f5) - Optional Block<br>Default Subset for Origin Pool. List of key-value pairs that define default subset. which gets used when route specifies no metadata or no subset matching the metadata exists

#### Default Pool Advanced Options Enable Subsets Endpoint Subsets

An [`endpoint_subsets`](#subsets-5171e5) block (within [`default_pool.advanced_options.enable_subsets`](#subsets-5a741c)) supports the following:

<a id="keys-1d8cbc"></a>&#x2022; [`keys`](#keys-1d8cbc) - Optional List<br>Keys. List of keys that define a cluster subset class

#### Default Pool Advanced Options Http1 Config

A [`http1_config`](#config-a0bc3c) block (within [`default_pool.advanced_options`](#default-pool-advanced-options)) supports the following:

<a id="transformation-cc4211"></a>&#x2022; [`header_transformation`](#transformation-cc4211) - Optional Block<br>Header Transformation. Header Transformation options for HTTP/1.1 request/response headers<br>See [Header Transformation](#transformation-cc4211) below.

#### Default Pool Advanced Options Http1 Config Header Transformation

A [`header_transformation`](#transformation-cc4211) block (within [`default_pool.advanced_options.http1_config`](#config-a0bc3c)) supports the following:

<a id="transformation-310fb4"></a>&#x2022; [`default_header_transformation`](#transformation-310fb4) - Optional Block<br>Enable this option

<a id="transformation-c1482f"></a>&#x2022; [`legacy_header_transformation`](#transformation-c1482f) - Optional Block<br>Enable this option

<a id="transformation-cb8377"></a>&#x2022; [`preserve_case_header_transformation`](#transformation-cb8377) - Optional Block<br>Enable this option

<a id="transformation-af9e0d"></a>&#x2022; [`proper_case_header_transformation`](#transformation-af9e0d) - Optional Block<br>Enable this option

#### Default Pool Advanced Options Http2 Options

A [`http2_options`](#options-fc9fd8) block (within [`default_pool.advanced_options`](#default-pool-advanced-options)) supports the following:

<a id="enabled-4e1198"></a>&#x2022; [`enabled`](#enabled-4e1198) - Optional Bool<br>HTTP2 Enabled. Enable/disable HTTP2 Protocol for upstream connections

#### Default Pool Advanced Options Outlier Detection

An [`outlier_detection`](#detection-c89e70) block (within [`default_pool.advanced_options`](#default-pool-advanced-options)) supports the following:

<a id="time-fd8e40"></a>&#x2022; [`base_ejection_time`](#time-fd8e40) - Optional Number  Defaults to `30000ms`  Specified in milliseconds<br>Base Ejection Time. The base time that a host is ejected for. The real time is equal to the base time multiplied by the number of times the host has been ejected. This causes hosts to get ejected for longer periods if they continue to fail

<a id="5xx-1c4ec1"></a>&#x2022; [`consecutive_5xx`](#5xx-1c4ec1) - Optional Number  Defaults to `5`<br>Consecutive 5xx Count. If an upstream endpoint returns some number of consecutive 5xx, it will be ejected. Note that in this case a 5xx means an actual 5xx respond code, or an event that would cause the HTTP router to return one on the upstream’s behalf(reset, connection failure, etc.) consecutive_5xx indicates the number of consecutive 5xx responses required before a consecutive 5xx ejection occurs

<a id="failure-b19342"></a>&#x2022; [`consecutive_gateway_failure`](#failure-b19342) - Optional Number  Defaults to `5`<br>Consecutive Gateway Failure. If an upstream endpoint returns some number of consecutive “gateway errors” (502, 503 or 504 status code), it will be ejected. Note that this includes events that would cause the HTTP router to return one of these status codes on the upstream’s behalf (reset, connection failure, etc.). consecutive_gateway_failure indicates the number of consecutive gateway failures before a consecutive gateway failure ejection occurs

<a id="interval-834cfc"></a>&#x2022; [`interval`](#interval-834cfc) - Optional Number  Defaults to `10000ms`  Specified in milliseconds<br>Interval. The time interval between ejection analysis sweeps. This can result in both new ejections as well as endpoints being returned to service

<a id="percent-9a52cc"></a>&#x2022; [`max_ejection_percent`](#percent-9a52cc) - Optional Number  Defaults to `10%`<br>Max Ejection Percentage. The maximum % of an upstream cluster that can be ejected due to outlier detection. but will eject at least one host regardless of the value

#### Default Pool Healthcheck

A [`healthcheck`](#default-pool-healthcheck) block (within [`default_pool`](#default-pool)) supports the following:

<a id="default-pool-healthcheck-name"></a>&#x2022; [`name`](#default-pool-healthcheck-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="default-pool-healthcheck-namespace"></a>&#x2022; [`namespace`](#default-pool-healthcheck-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="default-pool-healthcheck-tenant"></a>&#x2022; [`tenant`](#default-pool-healthcheck-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Default Pool Origin Servers

An [`origin_servers`](#default-pool-origin-servers) block (within [`default_pool`](#default-pool)) supports the following:

<a id="service-060a80"></a>&#x2022; [`cbip_service`](#service-060a80) - Optional Block<br>Discovered Classic BIG-IP Service Name. Specify origin server with Classic BIG-IP Service (Virtual Server)<br>See [Cbip Service](#service-060a80) below.

<a id="service-799005"></a>&#x2022; [`consul_service`](#service-799005) - Optional Block<br>Consul Service Name on given Sites. Specify origin server with Hashi Corp Consul service name and site information<br>See [Consul Service](#service-799005) below.

<a id="object-12dd7f"></a>&#x2022; [`custom_endpoint_object`](#object-12dd7f) - Optional Block<br>Custom Endpoint Object for Origin Server. Specify origin server with a reference to endpoint object<br>See [Custom Endpoint Object](#object-12dd7f) below.

<a id="default-pool-origin-servers-k8s-service"></a>&#x2022; [`k8s_service`](#default-pool-origin-servers-k8s-service) - Optional Block<br>K8S Service Name on given Sites. Specify origin server with K8S service name and site information<br>See [K8S Service](#default-pool-origin-servers-k8s-service) below.

<a id="default-pool-origin-servers-labels"></a>&#x2022; [`labels`](#default-pool-origin-servers-labels) - Optional Block<br>Origin Server Labels. Add Labels for this origin server, these labels can be used to form subset

<a id="default-pool-origin-servers-private-ip"></a>&#x2022; [`private_ip`](#default-pool-origin-servers-private-ip) - Optional Block<br>IP address on given Sites. Specify origin server with private or public IP address and site information<br>See [Private IP](#default-pool-origin-servers-private-ip) below.

<a id="name-966ae3"></a>&#x2022; [`private_name`](#name-966ae3) - Optional Block<br>DNS Name on given Sites. Specify origin server with private or public DNS name and site information<br>See [Private Name](#name-966ae3) below.

<a id="default-pool-origin-servers-public-ip"></a>&#x2022; [`public_ip`](#default-pool-origin-servers-public-ip) - Optional Block<br>Public IP. Specify origin server with public IP address<br>See [Public IP](#default-pool-origin-servers-public-ip) below.

<a id="default-pool-origin-servers-public-name"></a>&#x2022; [`public_name`](#default-pool-origin-servers-public-name) - Optional Block<br>Public DNS Name. Specify origin server with public DNS name<br>See [Public Name](#default-pool-origin-servers-public-name) below.

<a id="private-ip-532445"></a>&#x2022; [`vn_private_ip`](#private-ip-532445) - Optional Block<br>IP address Virtual Network. Specify origin server with IP on Virtual Network<br>See [Vn Private IP](#private-ip-532445) below.

<a id="name-4a1747"></a>&#x2022; [`vn_private_name`](#name-4a1747) - Optional Block<br>DNS Name on Virtual Network. Specify origin server with DNS name on Virtual Network<br>See [Vn Private Name](#name-4a1747) below.

#### Default Pool Origin Servers Cbip Service

A [`cbip_service`](#service-060a80) block (within [`default_pool.origin_servers`](#default-pool-origin-servers)) supports the following:

<a id="name-231ea0"></a>&#x2022; [`service_name`](#name-231ea0) - Optional String<br>Service Name. Name of the discovered Classic BIG-IP virtual server to be used as origin

#### Default Pool Origin Servers Consul Service

A [`consul_service`](#service-799005) block (within [`default_pool.origin_servers`](#default-pool-origin-servers)) supports the following:

<a id="network-654a28"></a>&#x2022; [`inside_network`](#network-654a28) - Optional Block<br>Enable this option

<a id="network-b1e5db"></a>&#x2022; [`outside_network`](#network-b1e5db) - Optional Block<br>Enable this option

<a id="name-5d42b9"></a>&#x2022; [`service_name`](#name-5d42b9) - Optional String<br>Service Name. Consul service name of this origin server will be listed, including cluster-ID. The format is servicename:cluster-ID

<a id="locator-7261e5"></a>&#x2022; [`site_locator`](#locator-7261e5) - Optional Block<br>Site or Virtual Site. This message defines a reference to a site or virtual site object<br>See [Site Locator](#locator-7261e5) below.

<a id="pool-b708db"></a>&#x2022; [`snat_pool`](#pool-b708db) - Optional Block<br>Snat Pool. Snat Pool configuration<br>See [Snat Pool](#pool-b708db) below.

#### Default Pool Origin Servers Consul Service Site Locator

A [`site_locator`](#locator-7261e5) block (within [`default_pool.origin_servers.consul_service`](#service-799005)) supports the following:

<a id="site-6ad63d"></a>&#x2022; [`site`](#site-6ad63d) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Site](#site-6ad63d) below.

<a id="site-009268"></a>&#x2022; [`virtual_site`](#site-009268) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Virtual Site](#site-009268) below.

#### Default Pool Origin Servers Consul Service Site Locator Site

A [`site`](#site-6ad63d) block (within [`default_pool.origin_servers.consul_service.site_locator`](#locator-7261e5)) supports the following:

<a id="name-5a5e90"></a>&#x2022; [`name`](#name-5a5e90) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-4accbf"></a>&#x2022; [`namespace`](#namespace-4accbf) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-0a2a17"></a>&#x2022; [`tenant`](#tenant-0a2a17) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Default Pool Origin Servers Consul Service Site Locator Virtual Site

A [`virtual_site`](#site-009268) block (within [`default_pool.origin_servers.consul_service.site_locator`](#locator-7261e5)) supports the following:

<a id="name-347220"></a>&#x2022; [`name`](#name-347220) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-d4ae84"></a>&#x2022; [`namespace`](#namespace-d4ae84) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-d740f9"></a>&#x2022; [`tenant`](#tenant-d740f9) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Default Pool Origin Servers Consul Service Snat Pool

A [`snat_pool`](#pool-b708db) block (within [`default_pool.origin_servers.consul_service`](#service-799005)) supports the following:

<a id="pool-c3480b"></a>&#x2022; [`no_snat_pool`](#pool-c3480b) - Optional Block<br>Enable this option

<a id="pool-61958b"></a>&#x2022; [`snat_pool`](#pool-61958b) - Optional Block<br>IPv4 Prefix List. List of IPv4 prefixes that represent an endpoint<br>See [Snat Pool](#pool-61958b) below.

#### Default Pool Origin Servers Consul Service Snat Pool Snat Pool

A [`snat_pool`](#pool-61958b) block (within [`default_pool.origin_servers.consul_service.snat_pool`](#pool-b708db)) supports the following:

<a id="prefixes-f065a9"></a>&#x2022; [`prefixes`](#prefixes-f065a9) - Optional List<br>IPv4 Prefix List. List of IPv4 prefixes that represent an endpoint

#### Default Pool Origin Servers Custom Endpoint Object

A [`custom_endpoint_object`](#object-12dd7f) block (within [`default_pool.origin_servers`](#default-pool-origin-servers)) supports the following:

<a id="endpoint-952bdd"></a>&#x2022; [`endpoint`](#endpoint-952bdd) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Endpoint](#endpoint-952bdd) below.

#### Default Pool Origin Servers Custom Endpoint Object Endpoint

An [`endpoint`](#endpoint-952bdd) block (within [`default_pool.origin_servers.custom_endpoint_object`](#object-12dd7f)) supports the following:

<a id="name-2281f2"></a>&#x2022; [`name`](#name-2281f2) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-06a905"></a>&#x2022; [`namespace`](#namespace-06a905) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-111c1e"></a>&#x2022; [`tenant`](#tenant-111c1e) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Default Pool Origin Servers K8S Service

A [`k8s_service`](#default-pool-origin-servers-k8s-service) block (within [`default_pool.origin_servers`](#default-pool-origin-servers)) supports the following:

<a id="network-dfbf17"></a>&#x2022; [`inside_network`](#network-dfbf17) - Optional Block<br>Enable this option

<a id="network-d1b956"></a>&#x2022; [`outside_network`](#network-d1b956) - Optional Block<br>Enable this option

<a id="protocol-ffcd27"></a>&#x2022; [`protocol`](#protocol-ffcd27) - Optional String  Defaults to `PROTOCOL_TCP`<br>Possible values are `PROTOCOL_TCP`, `PROTOCOL_UDP`<br>[Enum: PROTOCOL_TCP|PROTOCOL_UDP] Protocol Type. Type of protocol - PROTOCOL_TCP: TCP - PROTOCOL_UDP: UDP

<a id="name-c77159"></a>&#x2022; [`service_name`](#name-c77159) - Optional String<br>Service Name. K8S service name of the origin server will be listed, including the namespace and cluster-ID. For vK8s services, you need to enter a string with the format servicename.namespace:cluster-ID. If the servicename is 'frontend', namespace is 'speedtest' and cluster-ID is 'prod', then you will enter 'frontend.speedtest:prod'. Both namespace and cluster-ID are optional

<a id="locator-8a5921"></a>&#x2022; [`site_locator`](#locator-8a5921) - Optional Block<br>Site or Virtual Site. This message defines a reference to a site or virtual site object<br>See [Site Locator](#locator-8a5921) below.

<a id="pool-0640ac"></a>&#x2022; [`snat_pool`](#pool-0640ac) - Optional Block<br>Snat Pool. Snat Pool configuration<br>See [Snat Pool](#pool-0640ac) below.

<a id="networks-03f764"></a>&#x2022; [`vk8s_networks`](#networks-03f764) - Optional Block<br>Enable this option

#### Default Pool Origin Servers K8S Service Site Locator

A [`site_locator`](#locator-8a5921) block (within [`default_pool.origin_servers.k8s_service`](#default-pool-origin-servers-k8s-service)) supports the following:

<a id="site-69e713"></a>&#x2022; [`site`](#site-69e713) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Site](#site-69e713) below.

<a id="site-a3829f"></a>&#x2022; [`virtual_site`](#site-a3829f) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Virtual Site](#site-a3829f) below.

#### Default Pool Origin Servers K8S Service Site Locator Site

A [`site`](#site-69e713) block (within [`default_pool.origin_servers.k8s_service.site_locator`](#locator-8a5921)) supports the following:

<a id="name-de9835"></a>&#x2022; [`name`](#name-de9835) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-25c371"></a>&#x2022; [`namespace`](#namespace-25c371) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-cd3052"></a>&#x2022; [`tenant`](#tenant-cd3052) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Default Pool Origin Servers K8S Service Site Locator Virtual Site

A [`virtual_site`](#site-a3829f) block (within [`default_pool.origin_servers.k8s_service.site_locator`](#locator-8a5921)) supports the following:

<a id="name-15b617"></a>&#x2022; [`name`](#name-15b617) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-f78ef8"></a>&#x2022; [`namespace`](#namespace-f78ef8) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-4d300e"></a>&#x2022; [`tenant`](#tenant-4d300e) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Default Pool Origin Servers K8S Service Snat Pool

A [`snat_pool`](#pool-0640ac) block (within [`default_pool.origin_servers.k8s_service`](#default-pool-origin-servers-k8s-service)) supports the following:

<a id="pool-f2a6d4"></a>&#x2022; [`no_snat_pool`](#pool-f2a6d4) - Optional Block<br>Enable this option

<a id="pool-6caefd"></a>&#x2022; [`snat_pool`](#pool-6caefd) - Optional Block<br>IPv4 Prefix List. List of IPv4 prefixes that represent an endpoint<br>See [Snat Pool](#pool-6caefd) below.

#### Default Pool Origin Servers K8S Service Snat Pool Snat Pool

A [`snat_pool`](#pool-6caefd) block (within [`default_pool.origin_servers.k8s_service.snat_pool`](#pool-0640ac)) supports the following:

<a id="prefixes-fe8376"></a>&#x2022; [`prefixes`](#prefixes-fe8376) - Optional List<br>IPv4 Prefix List. List of IPv4 prefixes that represent an endpoint

#### Default Pool Origin Servers Private IP

A [`private_ip`](#default-pool-origin-servers-private-ip) block (within [`default_pool.origin_servers`](#default-pool-origin-servers)) supports the following:

<a id="network-704c7d"></a>&#x2022; [`inside_network`](#network-704c7d) - Optional Block<br>Enable this option

<a id="ip-ip-4a696b"></a>&#x2022; [`ip`](#ip-ip-4a696b) - Optional String<br>IP. Private IPv4 address

<a id="network-f44165"></a>&#x2022; [`outside_network`](#network-f44165) - Optional Block<br>Enable this option

<a id="segment-735aa1"></a>&#x2022; [`segment`](#segment-735aa1) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Segment](#segment-735aa1) below.

<a id="locator-1137c8"></a>&#x2022; [`site_locator`](#locator-1137c8) - Optional Block<br>Site or Virtual Site. This message defines a reference to a site or virtual site object<br>See [Site Locator](#locator-1137c8) below.

<a id="pool-916f33"></a>&#x2022; [`snat_pool`](#pool-916f33) - Optional Block<br>Snat Pool. Snat Pool configuration<br>See [Snat Pool](#pool-916f33) below.

#### Default Pool Origin Servers Private IP Segment

A [`segment`](#segment-735aa1) block (within [`default_pool.origin_servers.private_ip`](#default-pool-origin-servers-private-ip)) supports the following:

<a id="name-a7f063"></a>&#x2022; [`name`](#name-a7f063) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-4c09cc"></a>&#x2022; [`namespace`](#namespace-4c09cc) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-f972d4"></a>&#x2022; [`tenant`](#tenant-f972d4) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Default Pool Origin Servers Private IP Site Locator

A [`site_locator`](#locator-1137c8) block (within [`default_pool.origin_servers.private_ip`](#default-pool-origin-servers-private-ip)) supports the following:

<a id="site-65ee28"></a>&#x2022; [`site`](#site-65ee28) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Site](#site-65ee28) below.

<a id="site-4a5d0a"></a>&#x2022; [`virtual_site`](#site-4a5d0a) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Virtual Site](#site-4a5d0a) below.

#### Default Pool Origin Servers Private IP Site Locator Site

A [`site`](#site-65ee28) block (within [`default_pool.origin_servers.private_ip.site_locator`](#locator-1137c8)) supports the following:

<a id="name-820020"></a>&#x2022; [`name`](#name-820020) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-be09c6"></a>&#x2022; [`namespace`](#namespace-be09c6) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-65aac2"></a>&#x2022; [`tenant`](#tenant-65aac2) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Default Pool Origin Servers Private IP Site Locator Virtual Site

A [`virtual_site`](#site-4a5d0a) block (within [`default_pool.origin_servers.private_ip.site_locator`](#locator-1137c8)) supports the following:

<a id="name-5f1167"></a>&#x2022; [`name`](#name-5f1167) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-72fe71"></a>&#x2022; [`namespace`](#namespace-72fe71) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-d05c1a"></a>&#x2022; [`tenant`](#tenant-d05c1a) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Default Pool Origin Servers Private IP Snat Pool

A [`snat_pool`](#pool-916f33) block (within [`default_pool.origin_servers.private_ip`](#default-pool-origin-servers-private-ip)) supports the following:

<a id="pool-2e25b6"></a>&#x2022; [`no_snat_pool`](#pool-2e25b6) - Optional Block<br>Enable this option

<a id="pool-3655b2"></a>&#x2022; [`snat_pool`](#pool-3655b2) - Optional Block<br>IPv4 Prefix List. List of IPv4 prefixes that represent an endpoint<br>See [Snat Pool](#pool-3655b2) below.

#### Default Pool Origin Servers Private IP Snat Pool Snat Pool

A [`snat_pool`](#pool-3655b2) block (within [`default_pool.origin_servers.private_ip.snat_pool`](#pool-916f33)) supports the following:

<a id="prefixes-c6dd3e"></a>&#x2022; [`prefixes`](#prefixes-c6dd3e) - Optional List<br>IPv4 Prefix List. List of IPv4 prefixes that represent an endpoint

#### Default Pool Origin Servers Private Name

A [`private_name`](#name-966ae3) block (within [`default_pool.origin_servers`](#default-pool-origin-servers)) supports the following:

<a id="name-8a8021"></a>&#x2022; [`dns_name`](#name-8a8021) - Optional String<br>DNS Name. DNS Name

<a id="network-e9e813"></a>&#x2022; [`inside_network`](#network-e9e813) - Optional Block<br>Enable this option

<a id="network-873dcb"></a>&#x2022; [`outside_network`](#network-873dcb) - Optional Block<br>Enable this option

<a id="interval-615002"></a>&#x2022; [`refresh_interval`](#interval-615002) - Optional Number<br>DNS Refresh Interval. Interval for DNS refresh in seconds. Max value is 7 days as per `HTTPS://datatracker.ietf.org/doc/HTML/rfc8767`

<a id="segment-8fe482"></a>&#x2022; [`segment`](#segment-8fe482) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Segment](#segment-8fe482) below.

<a id="locator-3db1ee"></a>&#x2022; [`site_locator`](#locator-3db1ee) - Optional Block<br>Site or Virtual Site. This message defines a reference to a site or virtual site object<br>See [Site Locator](#locator-3db1ee) below.

<a id="pool-6d884b"></a>&#x2022; [`snat_pool`](#pool-6d884b) - Optional Block<br>Snat Pool. Snat Pool configuration<br>See [Snat Pool](#pool-6d884b) below.

#### Default Pool Origin Servers Private Name Segment

A [`segment`](#segment-8fe482) block (within [`default_pool.origin_servers.private_name`](#name-966ae3)) supports the following:

<a id="name-88ec1c"></a>&#x2022; [`name`](#name-88ec1c) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-2f15e9"></a>&#x2022; [`namespace`](#namespace-2f15e9) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-1fee5d"></a>&#x2022; [`tenant`](#tenant-1fee5d) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Default Pool Origin Servers Private Name Site Locator

A [`site_locator`](#locator-3db1ee) block (within [`default_pool.origin_servers.private_name`](#name-966ae3)) supports the following:

<a id="site-6955ea"></a>&#x2022; [`site`](#site-6955ea) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Site](#site-6955ea) below.

<a id="site-3c7bcd"></a>&#x2022; [`virtual_site`](#site-3c7bcd) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Virtual Site](#site-3c7bcd) below.

#### Default Pool Origin Servers Private Name Site Locator Site

A [`site`](#site-6955ea) block (within [`default_pool.origin_servers.private_name.site_locator`](#locator-3db1ee)) supports the following:

<a id="name-6cddba"></a>&#x2022; [`name`](#name-6cddba) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-215f18"></a>&#x2022; [`namespace`](#namespace-215f18) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-a48599"></a>&#x2022; [`tenant`](#tenant-a48599) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Default Pool Origin Servers Private Name Site Locator Virtual Site

A [`virtual_site`](#site-3c7bcd) block (within [`default_pool.origin_servers.private_name.site_locator`](#locator-3db1ee)) supports the following:

<a id="name-b8101c"></a>&#x2022; [`name`](#name-b8101c) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-6ee63a"></a>&#x2022; [`namespace`](#namespace-6ee63a) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-ab19fa"></a>&#x2022; [`tenant`](#tenant-ab19fa) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Default Pool Origin Servers Private Name Snat Pool

A [`snat_pool`](#pool-6d884b) block (within [`default_pool.origin_servers.private_name`](#name-966ae3)) supports the following:

<a id="pool-ec9835"></a>&#x2022; [`no_snat_pool`](#pool-ec9835) - Optional Block<br>Enable this option

<a id="pool-d9e993"></a>&#x2022; [`snat_pool`](#pool-d9e993) - Optional Block<br>IPv4 Prefix List. List of IPv4 prefixes that represent an endpoint<br>See [Snat Pool](#pool-d9e993) below.

#### Default Pool Origin Servers Private Name Snat Pool Snat Pool

A [`snat_pool`](#pool-d9e993) block (within [`default_pool.origin_servers.private_name.snat_pool`](#pool-6d884b)) supports the following:

<a id="prefixes-47d8d5"></a>&#x2022; [`prefixes`](#prefixes-47d8d5) - Optional List<br>IPv4 Prefix List. List of IPv4 prefixes that represent an endpoint

#### Default Pool Origin Servers Public IP

A [`public_ip`](#default-pool-origin-servers-public-ip) block (within [`default_pool.origin_servers`](#default-pool-origin-servers)) supports the following:

<a id="ip-ip-dc19b2"></a>&#x2022; [`ip`](#ip-ip-dc19b2) - Optional String<br>Public IPv4. Public IPv4 address

#### Default Pool Origin Servers Public Name

A [`public_name`](#default-pool-origin-servers-public-name) block (within [`default_pool.origin_servers`](#default-pool-origin-servers)) supports the following:

<a id="name-7c3f95"></a>&#x2022; [`dns_name`](#name-7c3f95) - Optional String<br>DNS Name. DNS Name

<a id="interval-ac5170"></a>&#x2022; [`refresh_interval`](#interval-ac5170) - Optional Number<br>DNS Refresh Interval. Interval for DNS refresh in seconds. Max value is 7 days as per `HTTPS://datatracker.ietf.org/doc/HTML/rfc8767`

#### Default Pool Origin Servers Vn Private IP

A [`vn_private_ip`](#private-ip-532445) block (within [`default_pool.origin_servers`](#default-pool-origin-servers)) supports the following:

<a id="ip-ip-c5d639"></a>&#x2022; [`ip`](#ip-ip-c5d639) - Optional String<br>IPv4. IPv4 address

<a id="network-56a203"></a>&#x2022; [`virtual_network`](#network-56a203) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Virtual Network](#network-56a203) below.

#### Default Pool Origin Servers Vn Private IP Virtual Network

A [`virtual_network`](#network-56a203) block (within [`default_pool.origin_servers.vn_private_ip`](#private-ip-532445)) supports the following:

<a id="name-324ace"></a>&#x2022; [`name`](#name-324ace) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-2680c8"></a>&#x2022; [`namespace`](#namespace-2680c8) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-bf4e7c"></a>&#x2022; [`tenant`](#tenant-bf4e7c) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Default Pool Origin Servers Vn Private Name

A [`vn_private_name`](#name-4a1747) block (within [`default_pool.origin_servers`](#default-pool-origin-servers)) supports the following:

<a id="name-10334e"></a>&#x2022; [`dns_name`](#name-10334e) - Optional String<br>DNS Name. DNS Name

<a id="network-bbece7"></a>&#x2022; [`private_network`](#network-bbece7) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Private Network](#network-bbece7) below.

#### Default Pool Origin Servers Vn Private Name Private Network

A [`private_network`](#network-bbece7) block (within [`default_pool.origin_servers.vn_private_name`](#name-4a1747)) supports the following:

<a id="name-a1668f"></a>&#x2022; [`name`](#name-a1668f) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-8f343b"></a>&#x2022; [`namespace`](#namespace-8f343b) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-e40ade"></a>&#x2022; [`tenant`](#tenant-e40ade) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Default Pool Upstream Conn Pool Reuse Type

An [`upstream_conn_pool_reuse_type`](#type-2756f7) block (within [`default_pool`](#default-pool)) supports the following:

<a id="reuse-6660fb"></a>&#x2022; [`disable_conn_pool_reuse`](#reuse-6660fb) - Optional Block<br>Enable this option

<a id="reuse-52dccb"></a>&#x2022; [`enable_conn_pool_reuse`](#reuse-52dccb) - Optional Block<br>Enable this option

#### Default Pool Use TLS

An [`use_tls`](#default-pool-use-tls) block (within [`default_pool`](#default-pool)) supports the following:

<a id="caching-6d4585"></a>&#x2022; [`default_session_key_caching`](#caching-6d4585) - Optional Block<br>Enable this option

<a id="caching-e75c13"></a>&#x2022; [`disable_session_key_caching`](#caching-e75c13) - Optional Block<br>Enable this option

<a id="default-pool-use-tls-disable-sni"></a>&#x2022; [`disable_sni`](#default-pool-use-tls-disable-sni) - Optional Block<br>Enable this option

<a id="default-pool-use-tls-max-session-keys"></a>&#x2022; [`max_session_keys`](#default-pool-use-tls-max-session-keys) - Optional Number<br>Max Session Keys Cached. Number of session keys that are cached

<a id="default-pool-use-tls-no-mtls"></a>&#x2022; [`no_mtls`](#default-pool-use-tls-no-mtls) - Optional Block<br>Enable this option

<a id="verification-a40775"></a>&#x2022; [`skip_server_verification`](#verification-a40775) - Optional Block<br>Enable this option

<a id="default-pool-use-tls-sni"></a>&#x2022; [`sni`](#default-pool-use-tls-sni) - Optional String<br>SNI Value. SNI value to be used

<a id="default-pool-use-tls-tls-config"></a>&#x2022; [`tls_config`](#default-pool-use-tls-tls-config) - Optional Block<br>TLS Config. This defines various options to configure TLS configuration parameters<br>See [TLS Config](#default-pool-use-tls-tls-config) below.

<a id="sni-a63eaf"></a>&#x2022; [`use_host_header_as_sni`](#sni-a63eaf) - Optional Block<br>Enable this option

<a id="default-pool-use-tls-use-mtls"></a>&#x2022; [`use_mtls`](#default-pool-use-tls-use-mtls) - Optional Block<br>mTLS Certificate. mTLS Client Certificate<br>See [Use mTLS](#default-pool-use-tls-use-mtls) below.

<a id="default-pool-use-tls-use-mtls-obj"></a>&#x2022; [`use_mtls_obj`](#default-pool-use-tls-use-mtls-obj) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Use mTLS Obj](#default-pool-use-tls-use-mtls-obj) below.

<a id="verification-388853"></a>&#x2022; [`use_server_verification`](#verification-388853) - Optional Block<br>TLS Validation Context for Origin Servers. Upstream TLS Validation Context<br>See [Use Server Verification](#verification-388853) below.

<a id="trusted-ca-e7a557"></a>&#x2022; [`volterra_trusted_ca`](#trusted-ca-e7a557) - Optional Block<br>Enable this option

#### Default Pool Use TLS TLS Config

A [`tls_config`](#default-pool-use-tls-tls-config) block (within [`default_pool.use_tls`](#default-pool-use-tls)) supports the following:

<a id="security-e7acc0"></a>&#x2022; [`custom_security`](#security-e7acc0) - Optional Block<br>Custom Ciphers. This defines TLS protocol config including min/max versions and allowed ciphers<br>See [Custom Security](#security-e7acc0) below.

<a id="security-b532bd"></a>&#x2022; [`default_security`](#security-b532bd) - Optional Block<br>Enable this option

<a id="security-33124a"></a>&#x2022; [`low_security`](#security-33124a) - Optional Block<br>Enable this option

<a id="security-ea8e0f"></a>&#x2022; [`medium_security`](#security-ea8e0f) - Optional Block<br>Enable this option

#### Default Pool Use TLS TLS Config Custom Security

A [`custom_security`](#security-e7acc0) block (within [`default_pool.use_tls.tls_config`](#default-pool-use-tls-tls-config)) supports the following:

<a id="suites-501641"></a>&#x2022; [`cipher_suites`](#suites-501641) - Optional List<br>Cipher Suites. The TLS listener will only support the specified cipher list

<a id="version-c70f29"></a>&#x2022; [`max_version`](#version-c70f29) - Optional String  Defaults to `TLS_AUTO`<br>Possible values are `TLS_AUTO`, `TLSv1_0`, `TLSv1_1`, `TLSv1_2`, `TLSv1_3`<br>[Enum: TLS_AUTO|TLSv1_0|TLSv1_1|TLSv1_2|TLSv1_3] TLS Protocol. TlsProtocol is enumeration of supported TLS versions F5 Distributed Cloud will choose the optimal TLS version

<a id="version-ed4870"></a>&#x2022; [`min_version`](#version-ed4870) - Optional String  Defaults to `TLS_AUTO`<br>Possible values are `TLS_AUTO`, `TLSv1_0`, `TLSv1_1`, `TLSv1_2`, `TLSv1_3`<br>[Enum: TLS_AUTO|TLSv1_0|TLSv1_1|TLSv1_2|TLSv1_3] TLS Protocol. TlsProtocol is enumeration of supported TLS versions F5 Distributed Cloud will choose the optimal TLS version

#### Default Pool Use TLS Use mTLS

An [`use_mtls`](#default-pool-use-tls-use-mtls) block (within [`default_pool.use_tls`](#default-pool-use-tls)) supports the following:

<a id="certificates-5055f8"></a>&#x2022; [`tls_certificates`](#certificates-5055f8) - Optional Block<br>mTLS Client Certificate. mTLS Client Certificate<br>See [TLS Certificates](#certificates-5055f8) below.

#### Default Pool Use TLS Use mTLS TLS Certificates

A [`tls_certificates`](#certificates-5055f8) block (within [`default_pool.use_tls.use_mtls`](#default-pool-use-tls-use-mtls)) supports the following:

<a id="url-85a8c8"></a>&#x2022; [`certificate_url`](#url-85a8c8) - Optional String<br>Certificate. TLS certificate. Certificate or certificate chain in PEM format including the PEM headers

<a id="algorithms-461e80"></a>&#x2022; [`custom_hash_algorithms`](#algorithms-461e80) - Optional Block<br>Hash Algorithms. Specifies the hash algorithms to be used<br>See [Custom Hash Algorithms](#algorithms-461e80) below.

<a id="spec-fa2bd3"></a>&#x2022; [`description_spec`](#spec-fa2bd3) - Optional String<br>Description. Description for the certificate

<a id="stapling-2533ef"></a>&#x2022; [`disable_ocsp_stapling`](#stapling-2533ef) - Optional Block<br>Enable this option

<a id="key-4883cf"></a>&#x2022; [`private_key`](#key-4883cf) - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Private Key](#key-4883cf) below.

<a id="defaults-834d93"></a>&#x2022; [`use_system_defaults`](#defaults-834d93) - Optional Block<br>Enable this option

#### Default Pool Use TLS Use mTLS TLS Certificates Custom Hash Algorithms

A [`custom_hash_algorithms`](#algorithms-461e80) block (within [`default_pool.use_tls.use_mtls.tls_certificates`](#certificates-5055f8)) supports the following:

<a id="algorithms-4cdfb3"></a>&#x2022; [`hash_algorithms`](#algorithms-4cdfb3) - Optional List  Defaults to `INVALID_HASH_ALGORITHM`<br>Possible values are `INVALID_HASH_ALGORITHM`, `SHA256`, `SHA1`<br>[Enum: INVALID_HASH_ALGORITHM|SHA256|SHA1] Hash Algorithms. Ordered list of hash algorithms to be used

#### Default Pool Use TLS Use mTLS TLS Certificates Private Key

A [`private_key`](#key-4883cf) block (within [`default_pool.use_tls.use_mtls.tls_certificates`](#certificates-5055f8)) supports the following:

<a id="info-1359a1"></a>&#x2022; [`blindfold_secret_info`](#info-1359a1) - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#info-1359a1) below.

<a id="info-6d26ad"></a>&#x2022; [`clear_secret_info`](#info-6d26ad) - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#info-6d26ad) below.

#### Default Pool Use TLS Use mTLS TLS Certificates Private Key Blindfold Secret Info

A [`blindfold_secret_info`](#info-1359a1) block (within [`default_pool.use_tls.use_mtls.tls_certificates.private_key`](#key-4883cf)) supports the following:

<a id="provider-9cd752"></a>&#x2022; [`decryption_provider`](#provider-9cd752) - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

<a id="location-84fe0a"></a>&#x2022; [`location`](#location-84fe0a) - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

<a id="provider-ca4ccb"></a>&#x2022; [`store_provider`](#provider-ca4ccb) - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

#### Default Pool Use TLS Use mTLS TLS Certificates Private Key Clear Secret Info

A [`clear_secret_info`](#info-6d26ad) block (within [`default_pool.use_tls.use_mtls.tls_certificates.private_key`](#key-4883cf)) supports the following:

<a id="ref-adc982"></a>&#x2022; [`provider_ref`](#ref-adc982) - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

<a id="url-befe21"></a>&#x2022; [`url`](#url-befe21) - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

#### Default Pool Use TLS Use mTLS Obj

An [`use_mtls_obj`](#default-pool-use-tls-use-mtls-obj) block (within [`default_pool.use_tls`](#default-pool-use-tls)) supports the following:

<a id="default-pool-use-tls-use-mtls-obj-name"></a>&#x2022; [`name`](#default-pool-use-tls-use-mtls-obj-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-5c1b91"></a>&#x2022; [`namespace`](#namespace-5c1b91) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-0b54b4"></a>&#x2022; [`tenant`](#tenant-0b54b4) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Default Pool Use TLS Use Server Verification

An [`use_server_verification`](#verification-388853) block (within [`default_pool.use_tls`](#default-pool-use-tls)) supports the following:

<a id="trusted-ca-1da800"></a>&#x2022; [`trusted_ca`](#trusted-ca-1da800) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Trusted CA](#trusted-ca-1da800) below.

<a id="url-d5fda5"></a>&#x2022; [`trusted_ca_url`](#url-d5fda5) - Optional String<br>Inline Root CA Certificate (legacy). Upload a Root CA Certificate specifically for this Origin Pool for verification of server's certificate

#### Default Pool Use TLS Use Server Verification Trusted CA

A [`trusted_ca`](#trusted-ca-1da800) block (within [`default_pool.use_tls.use_server_verification`](#verification-388853)) supports the following:

<a id="name-5db771"></a>&#x2022; [`name`](#name-5db771) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-5085fc"></a>&#x2022; [`namespace`](#namespace-5085fc) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-e12ebb"></a>&#x2022; [`tenant`](#tenant-e12ebb) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Default Pool View Internal

A [`view_internal`](#default-pool-view-internal) block (within [`default_pool`](#default-pool)) supports the following:

<a id="default-pool-view-internal-name"></a>&#x2022; [`name`](#default-pool-view-internal-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="default-pool-view-internal-namespace"></a>&#x2022; [`namespace`](#default-pool-view-internal-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="default-pool-view-internal-tenant"></a>&#x2022; [`tenant`](#default-pool-view-internal-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Default Pool List

A [`default_pool_list`](#default-pool-list) block supports the following:

<a id="default-pool-list-pools"></a>&#x2022; [`pools`](#default-pool-list-pools) - Optional Block<br>Origin Pools. List of Origin Pools<br>See [Pools](#default-pool-list-pools) below.

#### Default Pool List Pools

A [`pools`](#default-pool-list-pools) block (within [`default_pool_list`](#default-pool-list)) supports the following:

<a id="default-pool-list-pools-cluster"></a>&#x2022; [`cluster`](#default-pool-list-pools-cluster) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Cluster](#default-pool-list-pools-cluster) below.

<a id="subsets-f5b698"></a>&#x2022; [`endpoint_subsets`](#subsets-f5b698) - Optional Block<br>Origin Servers Subsets. Upstream origin pool may be configured to divide its origin servers into subsets based on metadata attached to the origin servers. Routes may then specify the metadata that a endpoint must match in order to be selected by the load balancer For origin servers which are discovered in K8S or Consul cluster, the label of the service is merged with endpoint's labels. In case of Consul, the label is derived from the 'Tag' field. For labels that are common between configured endpoint and discovered service, labels from discovered service takes precedence. List of key-value pairs that will be used as matching metadata. Only those origin servers of upstream origin pool which match this metadata will be selected for load balancing

<a id="default-pool-list-pools-pool"></a>&#x2022; [`pool`](#default-pool-list-pools-pool) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Pool](#default-pool-list-pools-pool) below.

<a id="default-pool-list-pools-priority"></a>&#x2022; [`priority`](#default-pool-list-pools-priority) - Optional Number<br>Priority. Priority of this origin pool, valid only with multiple origin pools. Value of 0 will make the pool as lowest priority origin pool Priority of 1 means highest priority and is considered active. When active origin pool is not available, lower priority origin pools are made active as per the increasing priority

<a id="default-pool-list-pools-weight"></a>&#x2022; [`weight`](#default-pool-list-pools-weight) - Optional Number<br>Weight. Weight of this origin pool, valid only with multiple origin pool. Value of 0 will disable the pool

#### Default Pool List Pools Cluster

A [`cluster`](#default-pool-list-pools-cluster) block (within [`default_pool_list.pools`](#default-pool-list-pools)) supports the following:

<a id="default-pool-list-pools-cluster-name"></a>&#x2022; [`name`](#default-pool-list-pools-cluster-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-4d4aed"></a>&#x2022; [`namespace`](#namespace-4d4aed) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="default-pool-list-pools-cluster-tenant"></a>&#x2022; [`tenant`](#default-pool-list-pools-cluster-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Default Pool List Pools Pool

A [`pool`](#default-pool-list-pools-pool) block (within [`default_pool_list.pools`](#default-pool-list-pools)) supports the following:

<a id="default-pool-list-pools-pool-name"></a>&#x2022; [`name`](#default-pool-list-pools-pool-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="default-pool-list-pools-pool-namespace"></a>&#x2022; [`namespace`](#default-pool-list-pools-pool-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="default-pool-list-pools-pool-tenant"></a>&#x2022; [`tenant`](#default-pool-list-pools-pool-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Default Route Pools

A [`default_route_pools`](#default-route-pools) block supports the following:

<a id="default-route-pools-cluster"></a>&#x2022; [`cluster`](#default-route-pools-cluster) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Cluster](#default-route-pools-cluster) below.

<a id="default-route-pools-endpoint-subsets"></a>&#x2022; [`endpoint_subsets`](#default-route-pools-endpoint-subsets) - Optional Block<br>Origin Servers Subsets. Upstream origin pool may be configured to divide its origin servers into subsets based on metadata attached to the origin servers. Routes may then specify the metadata that a endpoint must match in order to be selected by the load balancer For origin servers which are discovered in K8S or Consul cluster, the label of the service is merged with endpoint's labels. In case of Consul, the label is derived from the 'Tag' field. For labels that are common between configured endpoint and discovered service, labels from discovered service takes precedence. List of key-value pairs that will be used as matching metadata. Only those origin servers of upstream origin pool which match this metadata will be selected for load balancing

<a id="default-route-pools-pool"></a>&#x2022; [`pool`](#default-route-pools-pool) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Pool](#default-route-pools-pool) below.

<a id="default-route-pools-priority"></a>&#x2022; [`priority`](#default-route-pools-priority) - Optional Number<br>Priority. Priority of this origin pool, valid only with multiple origin pools. Value of 0 will make the pool as lowest priority origin pool Priority of 1 means highest priority and is considered active. When active origin pool is not available, lower priority origin pools are made active as per the increasing priority

<a id="default-route-pools-weight"></a>&#x2022; [`weight`](#default-route-pools-weight) - Optional Number<br>Weight. Weight of this origin pool, valid only with multiple origin pool. Value of 0 will disable the pool

#### Default Route Pools Cluster

A [`cluster`](#default-route-pools-cluster) block (within [`default_route_pools`](#default-route-pools)) supports the following:

<a id="default-route-pools-cluster-name"></a>&#x2022; [`name`](#default-route-pools-cluster-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="default-route-pools-cluster-namespace"></a>&#x2022; [`namespace`](#default-route-pools-cluster-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="default-route-pools-cluster-tenant"></a>&#x2022; [`tenant`](#default-route-pools-cluster-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Default Route Pools Pool

A [`pool`](#default-route-pools-pool) block (within [`default_route_pools`](#default-route-pools)) supports the following:

<a id="default-route-pools-pool-name"></a>&#x2022; [`name`](#default-route-pools-pool-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="default-route-pools-pool-namespace"></a>&#x2022; [`namespace`](#default-route-pools-pool-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="default-route-pools-pool-tenant"></a>&#x2022; [`tenant`](#default-route-pools-pool-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Enable API Discovery

An [`enable_api_discovery`](#enable-api-discovery) block supports the following:

<a id="enable-api-discovery-api-crawler"></a>&#x2022; [`api_crawler`](#enable-api-discovery-api-crawler) - Optional Block<br>API Crawling. API Crawler message<br>See [API Crawler](#enable-api-discovery-api-crawler) below.

<a id="scan-930604"></a>&#x2022; [`api_discovery_from_code_scan`](#scan-930604) - Optional Block<br>Select Code Base and Repositories<br>See [API Discovery From Code Scan](#scan-930604) below.

<a id="discovery-54db29"></a>&#x2022; [`custom_api_auth_discovery`](#discovery-54db29) - Optional Block<br>API Discovery Advanced Settings. API Discovery Advanced settings<br>See [Custom API Auth Discovery](#discovery-54db29) below.

<a id="discovery-29517f"></a>&#x2022; [`default_api_auth_discovery`](#discovery-29517f) - Optional Block<br>Enable this option

<a id="traffic-90c445"></a>&#x2022; [`disable_learn_from_redirect_traffic`](#traffic-90c445) - Optional Block<br>Enable this option

<a id="settings-c31c55"></a>&#x2022; [`discovered_api_settings`](#settings-c31c55) - Optional Block<br>Discovered API Settings. Configure Discovered API Settings<br>See [Discovered API Settings](#settings-c31c55) below.

<a id="traffic-074877"></a>&#x2022; [`enable_learn_from_redirect_traffic`](#traffic-074877) - Optional Block<br>Enable this option

#### Enable API Discovery API Crawler

An [`api_crawler`](#enable-api-discovery-api-crawler) block (within [`enable_api_discovery`](#enable-api-discovery)) supports the following:

<a id="config-1070d6"></a>&#x2022; [`api_crawler_config`](#config-1070d6) - Optional Block<br>Crawler Configure<br>See [API Crawler Config](#config-1070d6) below.

<a id="crawler-167f20"></a>&#x2022; [`disable_api_crawler`](#crawler-167f20) - Optional Block<br>Enable this option

#### Enable API Discovery API Crawler API Crawler Config

An [`api_crawler_config`](#config-1070d6) block (within [`enable_api_discovery.api_crawler`](#enable-api-discovery-api-crawler)) supports the following:

<a id="domains-5b24a2"></a>&#x2022; [`domains`](#domains-5b24a2) - Optional Block<br>Domains to Crawl. Enter domains and their credentials to allow authenticated API crawling. You can only include domains you own that are associated with this Load Balancer<br>See [Domains](#domains-5b24a2) below.

#### Enable API Discovery API Crawler API Crawler Config Domains

A [`domains`](#domains-5b24a2) block (within [`enable_api_discovery.api_crawler.api_crawler_config`](#config-1070d6)) supports the following:

<a id="domain-101008"></a>&#x2022; [`domain`](#domain-101008) - Optional String<br>Domain. Select the domain to execute API Crawling with given credentials

<a id="login-d7ed1c"></a>&#x2022; [`simple_login`](#login-d7ed1c) - Optional Block<br>Simple Login<br>See [Simple Login](#login-d7ed1c) below.

#### Enable API Discovery API Crawler API Crawler Config Domains Simple Login

A [`simple_login`](#login-d7ed1c) block (within [`enable_api_discovery.api_crawler.api_crawler_config.domains`](#domains-5b24a2)) supports the following:

<a id="password-6dce3d"></a>&#x2022; [`password`](#password-6dce3d) - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Password](#password-6dce3d) below.

<a id="user-6538fc"></a>&#x2022; [`user`](#user-6538fc) - Optional String<br>User. Enter the username to assign credentials for the selected domain to crawl

#### Enable API Discovery API Crawler API Crawler Config Domains Simple Login Password

A [`password`](#password-6dce3d) block (within [`enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login`](#login-d7ed1c)) supports the following:

<a id="info-0086db"></a>&#x2022; [`blindfold_secret_info`](#info-0086db) - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#info-0086db) below.

<a id="info-e77ed8"></a>&#x2022; [`clear_secret_info`](#info-e77ed8) - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#info-e77ed8) below.

#### Enable API Discovery API Crawler API Crawler Config Domains Simple Login Password Blindfold Secret Info

A [`blindfold_secret_info`](#info-0086db) block (within [`enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password`](#password-6dce3d)) supports the following:

<a id="provider-5866c6"></a>&#x2022; [`decryption_provider`](#provider-5866c6) - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

<a id="location-c9ff51"></a>&#x2022; [`location`](#location-c9ff51) - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

<a id="provider-a85d09"></a>&#x2022; [`store_provider`](#provider-a85d09) - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

#### Enable API Discovery API Crawler API Crawler Config Domains Simple Login Password Clear Secret Info

A [`clear_secret_info`](#info-e77ed8) block (within [`enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password`](#password-6dce3d)) supports the following:

<a id="ref-6d12d4"></a>&#x2022; [`provider_ref`](#ref-6d12d4) - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

<a id="url-54d163"></a>&#x2022; [`url`](#url-54d163) - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

#### Enable API Discovery API Discovery From Code Scan

An [`api_discovery_from_code_scan`](#scan-930604) block (within [`enable_api_discovery`](#enable-api-discovery)) supports the following:

<a id="integrations-684fd9"></a>&#x2022; [`code_base_integrations`](#integrations-684fd9) - Optional Block<br>Select Code Base Integrations<br>See [Code Base Integrations](#integrations-684fd9) below.

#### Enable API Discovery API Discovery From Code Scan Code Base Integrations

A [`code_base_integrations`](#integrations-684fd9) block (within [`enable_api_discovery.api_discovery_from_code_scan`](#scan-930604)) supports the following:

<a id="repos-6dd9b2"></a>&#x2022; [`all_repos`](#repos-6dd9b2) - Optional Block<br>Enable this option

<a id="integration-65ad07"></a>&#x2022; [`code_base_integration`](#integration-65ad07) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Code Base Integration](#integration-65ad07) below.

<a id="repos-85b753"></a>&#x2022; [`selected_repos`](#repos-85b753) - Optional Block<br>API Code Repositories. Select which API repositories represent the LB applications<br>See [Selected Repos](#repos-85b753) below.

#### Enable API Discovery API Discovery From Code Scan Code Base Integrations Code Base Integration

A [`code_base_integration`](#integration-65ad07) block (within [`enable_api_discovery.api_discovery_from_code_scan.code_base_integrations`](#integrations-684fd9)) supports the following:

<a id="name-c1c22e"></a>&#x2022; [`name`](#name-c1c22e) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-758b47"></a>&#x2022; [`namespace`](#namespace-758b47) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-ab614a"></a>&#x2022; [`tenant`](#tenant-ab614a) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Enable API Discovery API Discovery From Code Scan Code Base Integrations Selected Repos

A [`selected_repos`](#repos-85b753) block (within [`enable_api_discovery.api_discovery_from_code_scan.code_base_integrations`](#integrations-684fd9)) supports the following:

<a id="repo-27b7de"></a>&#x2022; [`api_code_repo`](#repo-27b7de) - Optional List<br>API Code Repository. Code repository which contain API endpoints

#### Enable API Discovery Custom API Auth Discovery

A [`custom_api_auth_discovery`](#discovery-54db29) block (within [`enable_api_discovery`](#enable-api-discovery)) supports the following:

<a id="ref-a70328"></a>&#x2022; [`api_discovery_ref`](#ref-a70328) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [API Discovery Ref](#ref-a70328) below.

#### Enable API Discovery Custom API Auth Discovery API Discovery Ref

An [`api_discovery_ref`](#ref-a70328) block (within [`enable_api_discovery.custom_api_auth_discovery`](#discovery-54db29)) supports the following:

<a id="name-168227"></a>&#x2022; [`name`](#name-168227) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-3af87c"></a>&#x2022; [`namespace`](#namespace-3af87c) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-e22b6f"></a>&#x2022; [`tenant`](#tenant-e22b6f) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Enable API Discovery Discovered API Settings

A [`discovered_api_settings`](#settings-c31c55) block (within [`enable_api_discovery`](#enable-api-discovery)) supports the following:

<a id="apis-cd00eb"></a>&#x2022; [`purge_duration_for_inactive_discovered_apis`](#apis-cd00eb) - Optional Number<br>Purge Duration for Inactive Discovered APIs from Traffic. Inactive discovered API will be deleted after configured duration

#### Enable Challenge

An [`enable_challenge`](#enable-challenge) block supports the following:

<a id="parameters-13a9c7"></a>&#x2022; [`captcha_challenge_parameters`](#parameters-13a9c7) - Optional Block<br>Captcha Challenge Parameters. Enables loadbalancer to perform captcha challenge Captcha challenge will be based on Google Recaptcha. With this feature enabled, only clients that pass the captcha challenge will be allowed to complete the HTTP request. When loadbalancer is configured to do Captcha Challenge, it will redirect the browser to an HTML page on every new HTTP request. This HTML page will have captcha challenge embedded in it. Client will be allowed to make the request only if the captcha challenge is successful. Loadbalancer will tag response header with a cookie to avoid Captcha challenge for subsequent requests. CAPTCHA is mainly used as a security check to ensure only human users can pass through. Generally, computers or bots are not capable of solving a captcha. You can enable either Javascript challenge or Captcha challenge on a virtual host<br>See [Captcha Challenge Parameters](#parameters-13a9c7) below.

<a id="parameters-247f74"></a>&#x2022; [`default_captcha_challenge_parameters`](#parameters-247f74) - Optional Block<br>Enable this option

<a id="parameters-e2729d"></a>&#x2022; [`default_js_challenge_parameters`](#parameters-e2729d) - Optional Block<br>Enable this option

<a id="settings-f4fda5"></a>&#x2022; [`default_mitigation_settings`](#settings-f4fda5) - Optional Block<br>Enable this option

<a id="parameters-6f7506"></a>&#x2022; [`js_challenge_parameters`](#parameters-6f7506) - Optional Block<br>Javascript Challenge Parameters. Enables loadbalancer to perform client browser compatibility test by redirecting to a page with Javascript. With this feature enabled, only clients that are capable of executing Javascript(mostly browsers) will be allowed to complete the HTTP request. When loadbalancer is configured to do Javascript Challenge, it will redirect the browser to an HTML page on every new HTTP request. This HTML page will have Javascript embedded in it. Loadbalancer chooses a set of random numbers for every new client and sends these numbers along with an encrypted answer with the request such that it embed these numbers as input in the Javascript. Javascript will run on the requestor browser and perform a complex Math operation. Script will submit the answer to loadbalancer. Loadbalancer will validate the answer by comparing the calculated answer with the decrypted answer (which was encrypted when it was sent back as reply) and allow the request to the upstream server only if the answer is correct. Loadbalancer will tag response header with a cookie to avoid Javascript challenge for subsequent requests. Javascript challenge serves following purposes \* Validate that the request is coming via a browser that is capable for running Javascript \* Force the browser to run a complex operation, f(X), that requires it to spend a large number of CPU cycles. This is to slow down a potential DOS attacker by making it difficult to launch a large request flood without having to spend even larger CPU cost at their end. You can enable either Javascript challenge or Captcha challenge on a virtual host<br>See [Js Challenge Parameters](#parameters-6f7506) below.

<a id="mitigation-b3e04b"></a>&#x2022; [`malicious_user_mitigation`](#mitigation-b3e04b) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Malicious User Mitigation](#mitigation-b3e04b) below.

#### Enable Challenge Captcha Challenge Parameters

A [`captcha_challenge_parameters`](#parameters-13a9c7) block (within [`enable_challenge`](#enable-challenge)) supports the following:

<a id="expiry-af25d3"></a>&#x2022; [`cookie_expiry`](#expiry-af25d3) - Optional Number<br>Cookie Expiration Period. Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new challenge

<a id="page-1f96cf"></a>&#x2022; [`custom_page`](#page-1f96cf) - Optional String<br>Custom message for Captcha Challenge. Custom message is of type uri_ref. Currently supported URL schemes is string:///. For string:/// scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. 'Please Wait.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '`<p>` Please Wait `</p>`'. Base64 encoded string for this HTML is 'PHA+IFBsZWFzZSBXYWl0IDwvcD4='

#### Enable Challenge Js Challenge Parameters

A [`js_challenge_parameters`](#parameters-6f7506) block (within [`enable_challenge`](#enable-challenge)) supports the following:

<a id="expiry-c03358"></a>&#x2022; [`cookie_expiry`](#expiry-c03358) - Optional Number<br>Cookie Expiration Period. Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new challenge

<a id="page-831ba9"></a>&#x2022; [`custom_page`](#page-831ba9) - Optional String<br>Custom Message for Javascript Challenge. Custom message is of type uri_ref. Currently supported URL schemes is string:///. For string:/// scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. 'Please Wait.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '`<p>` Please Wait `</p>`'. Base64 encoded string for this HTML is 'PHA+IFBsZWFzZSBXYWl0IDwvcD4='

<a id="delay-a5405d"></a>&#x2022; [`js_script_delay`](#delay-a5405d) - Optional Number<br>Javascript Delay. Delay introduced by Javascript, in milliseconds

#### Enable Challenge Malicious User Mitigation

A [`malicious_user_mitigation`](#mitigation-b3e04b) block (within [`enable_challenge`](#enable-challenge)) supports the following:

<a id="name-3a9364"></a>&#x2022; [`name`](#name-3a9364) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-38ef32"></a>&#x2022; [`namespace`](#namespace-38ef32) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-78def2"></a>&#x2022; [`tenant`](#tenant-78def2) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Enable IP Reputation

An [`enable_ip_reputation`](#enable-ip-reputation) block supports the following:

<a id="categories-bb360f"></a>&#x2022; [`ip_threat_categories`](#categories-bb360f) - Optional List  Defaults to `SPAM_SOURCES`<br>Possible values are `SPAM_SOURCES`, `WINDOWS_EXPLOITS`, `WEB_ATTACKS`, `BOTNETS`, `SCANNERS`, `REPUTATION`, `PHISHING`, `PROXY`, `MOBILE_THREATS`, `TOR_PROXY`, `DENIAL_OF_SERVICE`, `NETWORK`<br>[Enum: SPAM_SOURCES|WINDOWS_EXPLOITS|WEB_ATTACKS|BOTNETS|SCANNERS|REPUTATION|PHISHING|PROXY|MOBILE_THREATS|TOR_PROXY|DENIAL_OF_SERVICE|NETWORK] List of IP Threat Categories to choose. If the source IP matches on atleast one of the enabled IP threat categories, the request will be denied

#### Enable Trust Client IP Headers

An [`enable_trust_client_ip_headers`](#enable-trust-client-ip-headers) block supports the following:

<a id="headers-cb3a8e"></a>&#x2022; [`client_ip_headers`](#headers-cb3a8e) - Optional List<br>Client IP Headers. Define the list of one or more Client IP Headers. Headers will be used in order from top to bottom, meaning if the first header is not present in the request, the system will proceed to check for the second header, and so on, until one of the listed headers is found. If none of the defined headers exist, or the value is not an IP address, then the system will use the source IP of the packet. If multiple defined headers with different names are present in the request, the value of the first header name in the configuration will be used. If multiple defined headers with the same name are present in the request, values of all those headers will be combined. The system will read the right-most IP address from header, if there are multiple IP addresses in the header value. For X-Forwarded-For header, the system will read the IP address(rightmost - 1), as the client IP

#### GraphQL Rules

A [`graphql_rules`](#graphql-rules) block supports the following:

<a id="graphql-rules-any-domain"></a>&#x2022; [`any_domain`](#graphql-rules-any-domain) - Optional Block<br>Enable this option

<a id="graphql-rules-exact-path"></a>&#x2022; [`exact_path`](#graphql-rules-exact-path) - Optional String  Defaults to `/GraphQL`<br>Path. Specifies the exact path to GraphQL endpoint

<a id="graphql-rules-exact-value"></a>&#x2022; [`exact_value`](#graphql-rules-exact-value) - Optional String<br>Exact Value. Exact domain name

<a id="graphql-rules-graphql-settings"></a>&#x2022; [`graphql_settings`](#graphql-rules-graphql-settings) - Optional Block<br>GraphQL Settings. GraphQL configuration<br>See [GraphQL Settings](#graphql-rules-graphql-settings) below.

<a id="graphql-rules-metadata"></a>&#x2022; [`metadata`](#graphql-rules-metadata) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#graphql-rules-metadata) below.

<a id="graphql-rules-method-get"></a>&#x2022; [`method_get`](#graphql-rules-method-get) - Optional Block<br>Enable this option

<a id="graphql-rules-method-post"></a>&#x2022; [`method_post`](#graphql-rules-method-post) - Optional Block<br>Enable this option

<a id="graphql-rules-suffix-value"></a>&#x2022; [`suffix_value`](#graphql-rules-suffix-value) - Optional String<br>Suffix Value. Suffix of domain name e.g 'xyz.com' will match '*.xyz.com' and 'xyz.com'

#### GraphQL Rules GraphQL Settings

A [`graphql_settings`](#graphql-rules-graphql-settings) block (within [`graphql_rules`](#graphql-rules)) supports the following:

<a id="introspection-492a5f"></a>&#x2022; [`disable_introspection`](#introspection-492a5f) - Optional Block<br>Enable this option

<a id="introspection-762fd0"></a>&#x2022; [`enable_introspection`](#introspection-762fd0) - Optional Block<br>Enable this option

<a id="queries-f5cdb7"></a>&#x2022; [`max_batched_queries`](#queries-f5cdb7) - Optional Number<br>Maximum Batched Queries. Specify maximum number of queries in a single batched request

<a id="depth-42541b"></a>&#x2022; [`max_depth`](#depth-42541b) - Optional Number<br>Maximum Structure Depth. Specify maximum depth for the GraphQL query

<a id="length-21ac73"></a>&#x2022; [`max_total_length`](#length-21ac73) - Optional Number<br>Maximum Total Length. Specify maximum length in bytes for the GraphQL query

#### GraphQL Rules Metadata

A [`metadata`](#graphql-rules-metadata) block (within [`graphql_rules`](#graphql-rules)) supports the following:

<a id="graphql-rules-metadata-description-spec"></a>&#x2022; [`description_spec`](#graphql-rules-metadata-description-spec) - Optional String<br>Description. Human readable description

<a id="graphql-rules-metadata-name"></a>&#x2022; [`name`](#graphql-rules-metadata-name) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### HTTP

A [`http`](#http) block supports the following:

<a id="http-dns-volterra-managed"></a>&#x2022; [`dns_volterra_managed`](#http-dns-volterra-managed) - Optional Bool<br>Automatically Manage DNS Records. DNS records for domains will be managed automatically by F5 Distributed Cloud. As a prerequisite, the domain must be delegated to F5 Distributed Cloud using Delegated domain feature or a DNS CNAME record should be created in your DNS provider's portal

<a id="http-port"></a>&#x2022; [`port`](#http-port) - Optional Number<br>HTTP Listen Port. HTTP port to Listen

<a id="http-port-ranges"></a>&#x2022; [`port_ranges`](#http-port-ranges) - Optional String<br>Port Ranges. A string containing a comma separated list of port ranges. Each port range consists of a single port or two ports separated by '-'

#### HTTPS

A [`https`](#https) block supports the following:

<a id="https-add-hsts"></a>&#x2022; [`add_hsts`](#https-add-hsts) - Optional Bool<br>Add HSTS Header. Add HTTP Strict-Transport-Security response header

<a id="https-append-server-name"></a>&#x2022; [`append_server_name`](#https-append-server-name) - Optional String<br>Append header value. Define the header value for the header name “server”. If header value is already present, it is not overwritten and passed as-is

<a id="https-coalescing-options"></a>&#x2022; [`coalescing_options`](#https-coalescing-options) - Optional Block<br>TLS Coalescing Options. TLS connection coalescing configuration (not compatible with mTLS)<br>See [Coalescing Options](#https-coalescing-options) below.

<a id="https-connection-idle-timeout"></a>&#x2022; [`connection_idle_timeout`](#https-connection-idle-timeout) - Optional Number  Defaults to `2`  Specified in milliseconds<br>Connection Idle Timeout. The idle timeout for downstream connections. The idle timeout is defined as the period in which there are no active requests. When the idle timeout is reached the connection will be closed. Note that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive.  The minutes

<a id="https-default-header"></a>&#x2022; [`default_header`](#https-default-header) - Optional Block<br>Enable this option

<a id="https-default-loadbalancer"></a>&#x2022; [`default_loadbalancer`](#https-default-loadbalancer) - Optional Block<br>Enable this option

<a id="https-disable-path-normalize"></a>&#x2022; [`disable_path_normalize`](#https-disable-path-normalize) - Optional Block<br>Enable this option

<a id="https-enable-path-normalize"></a>&#x2022; [`enable_path_normalize`](#https-enable-path-normalize) - Optional Block<br>Enable this option

<a id="https-http-protocol-options"></a>&#x2022; [`http_protocol_options`](#https-http-protocol-options) - Optional Block<br>HTTP Protocol Configuration Options. HTTP protocol configuration options for downstream connections<br>See [HTTP Protocol Options](#https-http-protocol-options) below.

<a id="https-http-redirect"></a>&#x2022; [`http_redirect`](#https-http-redirect) - Optional Bool<br>HTTP Redirect to HTTPS. Redirect HTTP traffic to HTTPS

<a id="https-non-default-loadbalancer"></a>&#x2022; [`non_default_loadbalancer`](#https-non-default-loadbalancer) - Optional Block<br>Enable this option

<a id="https-pass-through"></a>&#x2022; [`pass_through`](#https-pass-through) - Optional Block<br>Enable this option

<a id="https-port"></a>&#x2022; [`port`](#https-port) - Optional Number<br>HTTPS Port. HTTPS port to Listen

<a id="https-port-ranges"></a>&#x2022; [`port_ranges`](#https-port-ranges) - Optional String<br>Port Ranges. A string containing a comma separated list of port ranges. Each port range consists of a single port or two ports separated by '-'

<a id="https-server-name"></a>&#x2022; [`server_name`](#https-server-name) - Optional String<br>Modify header value. Define the header value for the header name “server”. This will overwrite existing values, if any, for the server header

<a id="https-tls-cert-params"></a>&#x2022; [`tls_cert_params`](#https-tls-cert-params) - Optional Block<br>TLS Parameters. Select TLS Parameters and Certificates<br>See [TLS Cert Params](#https-tls-cert-params) below.

<a id="https-tls-parameters"></a>&#x2022; [`tls_parameters`](#https-tls-parameters) - Optional Block<br>Inline TLS Parameters. Inline TLS parameters<br>See [TLS Parameters](#https-tls-parameters) below.

#### HTTPS Coalescing Options

A [`coalescing_options`](#https-coalescing-options) block (within [`https`](#https)) supports the following:

<a id="coalescing-f90c69"></a>&#x2022; [`default_coalescing`](#coalescing-f90c69) - Optional Block<br>Enable this option

<a id="coalescing-c5278e"></a>&#x2022; [`strict_coalescing`](#coalescing-c5278e) - Optional Block<br>Enable this option

#### HTTPS HTTP Protocol Options

A [`http_protocol_options`](#https-http-protocol-options) block (within [`https`](#https)) supports the following:

<a id="only-46f3ca"></a>&#x2022; [`http_protocol_enable_v1_only`](#only-46f3ca) - Optional Block<br>HTTP/1.1 Protocol Options. HTTP/1.1 Protocol options for downstream connections<br>See [HTTP Protocol Enable V1 Only](#only-46f3ca) below.

<a id="v1-v2-6f8b9b"></a>&#x2022; [`http_protocol_enable_v1_v2`](#v1-v2-6f8b9b) - Optional Block<br>Enable this option

<a id="only-5cefb3"></a>&#x2022; [`http_protocol_enable_v2_only`](#only-5cefb3) - Optional Block<br>Enable this option

#### HTTPS HTTP Protocol Options HTTP Protocol Enable V1 Only

A [`http_protocol_enable_v1_only`](#only-46f3ca) block (within [`https.http_protocol_options`](#https-http-protocol-options)) supports the following:

<a id="transformation-0f7183"></a>&#x2022; [`header_transformation`](#transformation-0f7183) - Optional Block<br>Header Transformation. Header Transformation options for HTTP/1.1 request/response headers<br>See [Header Transformation](#transformation-0f7183) below.

#### HTTPS HTTP Protocol Options HTTP Protocol Enable V1 Only Header Transformation

A [`header_transformation`](#transformation-0f7183) block (within [`https.http_protocol_options.http_protocol_enable_v1_only`](#only-46f3ca)) supports the following:

<a id="transformation-58cde9"></a>&#x2022; [`default_header_transformation`](#transformation-58cde9) - Optional Block<br>Enable this option

<a id="transformation-1bd32a"></a>&#x2022; [`legacy_header_transformation`](#transformation-1bd32a) - Optional Block<br>Enable this option

<a id="transformation-edda67"></a>&#x2022; [`preserve_case_header_transformation`](#transformation-edda67) - Optional Block<br>Enable this option

<a id="transformation-b085be"></a>&#x2022; [`proper_case_header_transformation`](#transformation-b085be) - Optional Block<br>Enable this option

#### HTTPS TLS Cert Params

A [`tls_cert_params`](#https-tls-cert-params) block (within [`https`](#https)) supports the following:

<a id="https-tls-cert-params-certificates"></a>&#x2022; [`certificates`](#https-tls-cert-params-certificates) - Optional Block<br>Certificates. Select one or more certificates with any domain names<br>See [Certificates](#https-tls-cert-params-certificates) below.

<a id="https-tls-cert-params-no-mtls"></a>&#x2022; [`no_mtls`](#https-tls-cert-params-no-mtls) - Optional Block<br>Enable this option

<a id="https-tls-cert-params-tls-config"></a>&#x2022; [`tls_config`](#https-tls-cert-params-tls-config) - Optional Block<br>TLS Config. This defines various options to configure TLS configuration parameters<br>See [TLS Config](#https-tls-cert-params-tls-config) below.

<a id="https-tls-cert-params-use-mtls"></a>&#x2022; [`use_mtls`](#https-tls-cert-params-use-mtls) - Optional Block<br>Clients TLS validation context. Validation context for downstream client TLS connections<br>See [Use mTLS](#https-tls-cert-params-use-mtls) below.

#### HTTPS TLS Cert Params Certificates

A [`certificates`](#https-tls-cert-params-certificates) block (within [`https.tls_cert_params`](#https-tls-cert-params)) supports the following:

<a id="https-tls-cert-params-certificates-name"></a>&#x2022; [`name`](#https-tls-cert-params-certificates-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-74e8ce"></a>&#x2022; [`namespace`](#namespace-74e8ce) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-7c270a"></a>&#x2022; [`tenant`](#tenant-7c270a) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### HTTPS TLS Cert Params TLS Config

A [`tls_config`](#https-tls-cert-params-tls-config) block (within [`https.tls_cert_params`](#https-tls-cert-params)) supports the following:

<a id="security-6452ce"></a>&#x2022; [`custom_security`](#security-6452ce) - Optional Block<br>Custom Ciphers. This defines TLS protocol config including min/max versions and allowed ciphers<br>See [Custom Security](#security-6452ce) below.

<a id="security-b6db5a"></a>&#x2022; [`default_security`](#security-b6db5a) - Optional Block<br>Enable this option

<a id="security-cbe12e"></a>&#x2022; [`low_security`](#security-cbe12e) - Optional Block<br>Enable this option

<a id="security-e410e3"></a>&#x2022; [`medium_security`](#security-e410e3) - Optional Block<br>Enable this option

#### HTTPS TLS Cert Params TLS Config Custom Security

A [`custom_security`](#security-6452ce) block (within [`https.tls_cert_params.tls_config`](#https-tls-cert-params-tls-config)) supports the following:

<a id="suites-2dd1c6"></a>&#x2022; [`cipher_suites`](#suites-2dd1c6) - Optional List<br>Cipher Suites. The TLS listener will only support the specified cipher list

<a id="version-17f0b1"></a>&#x2022; [`max_version`](#version-17f0b1) - Optional String  Defaults to `TLS_AUTO`<br>Possible values are `TLS_AUTO`, `TLSv1_0`, `TLSv1_1`, `TLSv1_2`, `TLSv1_3`<br>[Enum: TLS_AUTO|TLSv1_0|TLSv1_1|TLSv1_2|TLSv1_3] TLS Protocol. TlsProtocol is enumeration of supported TLS versions F5 Distributed Cloud will choose the optimal TLS version

<a id="version-1bef47"></a>&#x2022; [`min_version`](#version-1bef47) - Optional String  Defaults to `TLS_AUTO`<br>Possible values are `TLS_AUTO`, `TLSv1_0`, `TLSv1_1`, `TLSv1_2`, `TLSv1_3`<br>[Enum: TLS_AUTO|TLSv1_0|TLSv1_1|TLSv1_2|TLSv1_3] TLS Protocol. TlsProtocol is enumeration of supported TLS versions F5 Distributed Cloud will choose the optimal TLS version

#### HTTPS TLS Cert Params Use mTLS

An [`use_mtls`](#https-tls-cert-params-use-mtls) block (within [`https.tls_cert_params`](#https-tls-cert-params)) supports the following:

<a id="optional-fd9757"></a>&#x2022; [`client_certificate_optional`](#optional-fd9757) - Optional Bool<br>Client Certificate Optional. Client certificate is optional. If the client has provided a certificate, the load balancer will verify it. If certification verification fails, the connection will be terminated. If the client does not provide a certificate, the connection will be accepted

<a id="https-tls-cert-params-use-mtls-crl"></a>&#x2022; [`crl`](#https-tls-cert-params-use-mtls-crl) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [CRL](#https-tls-cert-params-use-mtls-crl) below.

<a id="https-tls-cert-params-use-mtls-no-crl"></a>&#x2022; [`no_crl`](#https-tls-cert-params-use-mtls-no-crl) - Optional Block<br>Enable this option

<a id="trusted-ca-2ba851"></a>&#x2022; [`trusted_ca`](#trusted-ca-2ba851) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Trusted CA](#trusted-ca-2ba851) below.

<a id="url-2b1433"></a>&#x2022; [`trusted_ca_url`](#url-2b1433) - Optional String<br>Inline Root CA Certificate (legacy). Upload a Root CA Certificate specifically for this Load Balancer

<a id="disabled-bc6638"></a>&#x2022; [`xfcc_disabled`](#disabled-bc6638) - Optional Block<br>Enable this option

<a id="options-8f161e"></a>&#x2022; [`xfcc_options`](#options-8f161e) - Optional Block<br>XFCC Header Elements. X-Forwarded-Client-Cert header elements to be added to requests<br>See [Xfcc Options](#options-8f161e) below.

#### HTTPS TLS Cert Params Use mTLS CRL

A [`crl`](#https-tls-cert-params-use-mtls-crl) block (within [`https.tls_cert_params.use_mtls`](#https-tls-cert-params-use-mtls)) supports the following:

<a id="https-tls-cert-params-use-mtls-crl-name"></a>&#x2022; [`name`](#https-tls-cert-params-use-mtls-crl-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-4349dd"></a>&#x2022; [`namespace`](#namespace-4349dd) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-7e4839"></a>&#x2022; [`tenant`](#tenant-7e4839) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### HTTPS TLS Cert Params Use mTLS Trusted CA

A [`trusted_ca`](#trusted-ca-2ba851) block (within [`https.tls_cert_params.use_mtls`](#https-tls-cert-params-use-mtls)) supports the following:

<a id="name-d64fa2"></a>&#x2022; [`name`](#name-d64fa2) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-ba276e"></a>&#x2022; [`namespace`](#namespace-ba276e) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-8677c1"></a>&#x2022; [`tenant`](#tenant-8677c1) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### HTTPS TLS Cert Params Use mTLS Xfcc Options

A [`xfcc_options`](#options-8f161e) block (within [`https.tls_cert_params.use_mtls`](#https-tls-cert-params-use-mtls)) supports the following:

<a id="elements-f88d19"></a>&#x2022; [`xfcc_header_elements`](#elements-f88d19) - Optional List  Defaults to `XFCC_NONE`<br>Possible values are `XFCC_NONE`, `XFCC_CERT`, `XFCC_CHAIN`, `XFCC_SUBJECT`, `XFCC_URI`, `XFCC_DNS`<br>[Enum: XFCC_NONE|XFCC_CERT|XFCC_CHAIN|XFCC_SUBJECT|XFCC_URI|XFCC_DNS] XFCC Header Elements. X-Forwarded-Client-Cert header elements to be added to requests

#### HTTPS TLS Parameters

A [`tls_parameters`](#https-tls-parameters) block (within [`https`](#https)) supports the following:

<a id="https-tls-parameters-no-mtls"></a>&#x2022; [`no_mtls`](#https-tls-parameters-no-mtls) - Optional Block<br>Enable this option

<a id="https-tls-parameters-tls-certificates"></a>&#x2022; [`tls_certificates`](#https-tls-parameters-tls-certificates) - Optional Block<br>TLS Certificates. Users can add one or more certificates that share the same set of domains. for example, domain.com and \*.domain.com - but use different signature algorithms<br>See [TLS Certificates](#https-tls-parameters-tls-certificates) below.

<a id="https-tls-parameters-tls-config"></a>&#x2022; [`tls_config`](#https-tls-parameters-tls-config) - Optional Block<br>TLS Config. This defines various options to configure TLS configuration parameters<br>See [TLS Config](#https-tls-parameters-tls-config) below.

<a id="https-tls-parameters-use-mtls"></a>&#x2022; [`use_mtls`](#https-tls-parameters-use-mtls) - Optional Block<br>Clients TLS validation context. Validation context for downstream client TLS connections<br>See [Use mTLS](#https-tls-parameters-use-mtls) below.

#### HTTPS TLS Parameters TLS Certificates

A [`tls_certificates`](#https-tls-parameters-tls-certificates) block (within [`https.tls_parameters`](#https-tls-parameters)) supports the following:

<a id="url-86f976"></a>&#x2022; [`certificate_url`](#url-86f976) - Optional String<br>Certificate. TLS certificate. Certificate or certificate chain in PEM format including the PEM headers

<a id="algorithms-9ccf95"></a>&#x2022; [`custom_hash_algorithms`](#algorithms-9ccf95) - Optional Block<br>Hash Algorithms. Specifies the hash algorithms to be used<br>See [Custom Hash Algorithms](#algorithms-9ccf95) below.

<a id="spec-77c99c"></a>&#x2022; [`description_spec`](#spec-77c99c) - Optional String<br>Description. Description for the certificate

<a id="stapling-fd931a"></a>&#x2022; [`disable_ocsp_stapling`](#stapling-fd931a) - Optional Block<br>Enable this option

<a id="key-372460"></a>&#x2022; [`private_key`](#key-372460) - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Private Key](#key-372460) below.

<a id="defaults-2777d9"></a>&#x2022; [`use_system_defaults`](#defaults-2777d9) - Optional Block<br>Enable this option

#### HTTPS TLS Parameters TLS Certificates Custom Hash Algorithms

A [`custom_hash_algorithms`](#algorithms-9ccf95) block (within [`https.tls_parameters.tls_certificates`](#https-tls-parameters-tls-certificates)) supports the following:

<a id="algorithms-d09bed"></a>&#x2022; [`hash_algorithms`](#algorithms-d09bed) - Optional List  Defaults to `INVALID_HASH_ALGORITHM`<br>Possible values are `INVALID_HASH_ALGORITHM`, `SHA256`, `SHA1`<br>[Enum: INVALID_HASH_ALGORITHM|SHA256|SHA1] Hash Algorithms. Ordered list of hash algorithms to be used

#### HTTPS TLS Parameters TLS Certificates Private Key

A [`private_key`](#key-372460) block (within [`https.tls_parameters.tls_certificates`](#https-tls-parameters-tls-certificates)) supports the following:

<a id="info-0c9fbe"></a>&#x2022; [`blindfold_secret_info`](#info-0c9fbe) - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#info-0c9fbe) below.

<a id="info-556650"></a>&#x2022; [`clear_secret_info`](#info-556650) - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#info-556650) below.

#### HTTPS TLS Parameters TLS Certificates Private Key Blindfold Secret Info

A [`blindfold_secret_info`](#info-0c9fbe) block (within [`https.tls_parameters.tls_certificates.private_key`](#key-372460)) supports the following:

<a id="provider-c727f0"></a>&#x2022; [`decryption_provider`](#provider-c727f0) - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

<a id="location-558d7e"></a>&#x2022; [`location`](#location-558d7e) - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

<a id="provider-31d2e0"></a>&#x2022; [`store_provider`](#provider-31d2e0) - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

#### HTTPS TLS Parameters TLS Certificates Private Key Clear Secret Info

A [`clear_secret_info`](#info-556650) block (within [`https.tls_parameters.tls_certificates.private_key`](#key-372460)) supports the following:

<a id="ref-1fb9e3"></a>&#x2022; [`provider_ref`](#ref-1fb9e3) - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

<a id="url-78aab7"></a>&#x2022; [`url`](#url-78aab7) - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

#### HTTPS TLS Parameters TLS Config

A [`tls_config`](#https-tls-parameters-tls-config) block (within [`https.tls_parameters`](#https-tls-parameters)) supports the following:

<a id="security-775274"></a>&#x2022; [`custom_security`](#security-775274) - Optional Block<br>Custom Ciphers. This defines TLS protocol config including min/max versions and allowed ciphers<br>See [Custom Security](#security-775274) below.

<a id="security-fd7aef"></a>&#x2022; [`default_security`](#security-fd7aef) - Optional Block<br>Enable this option

<a id="security-c7d5df"></a>&#x2022; [`low_security`](#security-c7d5df) - Optional Block<br>Enable this option

<a id="security-3f26ed"></a>&#x2022; [`medium_security`](#security-3f26ed) - Optional Block<br>Enable this option

#### HTTPS TLS Parameters TLS Config Custom Security

A [`custom_security`](#security-775274) block (within [`https.tls_parameters.tls_config`](#https-tls-parameters-tls-config)) supports the following:

<a id="suites-4e37c7"></a>&#x2022; [`cipher_suites`](#suites-4e37c7) - Optional List<br>Cipher Suites. The TLS listener will only support the specified cipher list

<a id="version-746a9d"></a>&#x2022; [`max_version`](#version-746a9d) - Optional String  Defaults to `TLS_AUTO`<br>Possible values are `TLS_AUTO`, `TLSv1_0`, `TLSv1_1`, `TLSv1_2`, `TLSv1_3`<br>[Enum: TLS_AUTO|TLSv1_0|TLSv1_1|TLSv1_2|TLSv1_3] TLS Protocol. TlsProtocol is enumeration of supported TLS versions F5 Distributed Cloud will choose the optimal TLS version

<a id="version-64a607"></a>&#x2022; [`min_version`](#version-64a607) - Optional String  Defaults to `TLS_AUTO`<br>Possible values are `TLS_AUTO`, `TLSv1_0`, `TLSv1_1`, `TLSv1_2`, `TLSv1_3`<br>[Enum: TLS_AUTO|TLSv1_0|TLSv1_1|TLSv1_2|TLSv1_3] TLS Protocol. TlsProtocol is enumeration of supported TLS versions F5 Distributed Cloud will choose the optimal TLS version

#### HTTPS TLS Parameters Use mTLS

An [`use_mtls`](#https-tls-parameters-use-mtls) block (within [`https.tls_parameters`](#https-tls-parameters)) supports the following:

<a id="optional-9a0326"></a>&#x2022; [`client_certificate_optional`](#optional-9a0326) - Optional Bool<br>Client Certificate Optional. Client certificate is optional. If the client has provided a certificate, the load balancer will verify it. If certification verification fails, the connection will be terminated. If the client does not provide a certificate, the connection will be accepted

<a id="https-tls-parameters-use-mtls-crl"></a>&#x2022; [`crl`](#https-tls-parameters-use-mtls-crl) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [CRL](#https-tls-parameters-use-mtls-crl) below.

<a id="https-tls-parameters-use-mtls-no-crl"></a>&#x2022; [`no_crl`](#https-tls-parameters-use-mtls-no-crl) - Optional Block<br>Enable this option

<a id="trusted-ca-264d37"></a>&#x2022; [`trusted_ca`](#trusted-ca-264d37) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Trusted CA](#trusted-ca-264d37) below.

<a id="url-bc2530"></a>&#x2022; [`trusted_ca_url`](#url-bc2530) - Optional String<br>Inline Root CA Certificate (legacy). Upload a Root CA Certificate specifically for this Load Balancer

<a id="disabled-5c360d"></a>&#x2022; [`xfcc_disabled`](#disabled-5c360d) - Optional Block<br>Enable this option

<a id="options-4d1e53"></a>&#x2022; [`xfcc_options`](#options-4d1e53) - Optional Block<br>XFCC Header Elements. X-Forwarded-Client-Cert header elements to be added to requests<br>See [Xfcc Options](#options-4d1e53) below.

#### HTTPS TLS Parameters Use mTLS CRL

A [`crl`](#https-tls-parameters-use-mtls-crl) block (within [`https.tls_parameters.use_mtls`](#https-tls-parameters-use-mtls)) supports the following:

<a id="https-tls-parameters-use-mtls-crl-name"></a>&#x2022; [`name`](#https-tls-parameters-use-mtls-crl-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-ae83ff"></a>&#x2022; [`namespace`](#namespace-ae83ff) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-08da33"></a>&#x2022; [`tenant`](#tenant-08da33) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### HTTPS TLS Parameters Use mTLS Trusted CA

A [`trusted_ca`](#trusted-ca-264d37) block (within [`https.tls_parameters.use_mtls`](#https-tls-parameters-use-mtls)) supports the following:

<a id="name-7721af"></a>&#x2022; [`name`](#name-7721af) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-70022b"></a>&#x2022; [`namespace`](#namespace-70022b) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-99da02"></a>&#x2022; [`tenant`](#tenant-99da02) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### HTTPS TLS Parameters Use mTLS Xfcc Options

A [`xfcc_options`](#options-4d1e53) block (within [`https.tls_parameters.use_mtls`](#https-tls-parameters-use-mtls)) supports the following:

<a id="elements-52ccb1"></a>&#x2022; [`xfcc_header_elements`](#elements-52ccb1) - Optional List  Defaults to `XFCC_NONE`<br>Possible values are `XFCC_NONE`, `XFCC_CERT`, `XFCC_CHAIN`, `XFCC_SUBJECT`, `XFCC_URI`, `XFCC_DNS`<br>[Enum: XFCC_NONE|XFCC_CERT|XFCC_CHAIN|XFCC_SUBJECT|XFCC_URI|XFCC_DNS] XFCC Header Elements. X-Forwarded-Client-Cert header elements to be added to requests

#### HTTPS Auto Cert

A [`https_auto_cert`](#https-auto-cert) block supports the following:

<a id="https-auto-cert-add-hsts"></a>&#x2022; [`add_hsts`](#https-auto-cert-add-hsts) - Optional Bool<br>Add HSTS Header. Add HTTP Strict-Transport-Security response header

<a id="https-auto-cert-append-server-name"></a>&#x2022; [`append_server_name`](#https-auto-cert-append-server-name) - Optional String<br>Append header value. Define the header value for the header name “server”. If header value is already present, it is not overwritten and passed as-is

<a id="https-auto-cert-coalescing-options"></a>&#x2022; [`coalescing_options`](#https-auto-cert-coalescing-options) - Optional Block<br>TLS Coalescing Options. TLS connection coalescing configuration (not compatible with mTLS)<br>See [Coalescing Options](#https-auto-cert-coalescing-options) below.

<a id="https-auto-cert-connection-idle-timeout"></a>&#x2022; [`connection_idle_timeout`](#https-auto-cert-connection-idle-timeout) - Optional Number  Defaults to `2`  Specified in milliseconds<br>Connection Idle Timeout. The idle timeout for downstream connections. The idle timeout is defined as the period in which there are no active requests. When the idle timeout is reached the connection will be closed. Note that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive.  The minutes

<a id="https-auto-cert-default-header"></a>&#x2022; [`default_header`](#https-auto-cert-default-header) - Optional Block<br>Enable this option

<a id="https-auto-cert-default-loadbalancer"></a>&#x2022; [`default_loadbalancer`](#https-auto-cert-default-loadbalancer) - Optional Block<br>Enable this option

<a id="https-auto-cert-disable-path-normalize"></a>&#x2022; [`disable_path_normalize`](#https-auto-cert-disable-path-normalize) - Optional Block<br>Enable this option

<a id="https-auto-cert-enable-path-normalize"></a>&#x2022; [`enable_path_normalize`](#https-auto-cert-enable-path-normalize) - Optional Block<br>Enable this option

<a id="https-auto-cert-http-protocol-options"></a>&#x2022; [`http_protocol_options`](#https-auto-cert-http-protocol-options) - Optional Block<br>HTTP Protocol Configuration Options. HTTP protocol configuration options for downstream connections<br>See [HTTP Protocol Options](#https-auto-cert-http-protocol-options) below.

<a id="https-auto-cert-http-redirect"></a>&#x2022; [`http_redirect`](#https-auto-cert-http-redirect) - Optional Bool<br>HTTP Redirect to HTTPS. Redirect HTTP traffic to HTTPS

<a id="https-auto-cert-no-mtls"></a>&#x2022; [`no_mtls`](#https-auto-cert-no-mtls) - Optional Block<br>Enable this option

<a id="loadbalancer-eb605c"></a>&#x2022; [`non_default_loadbalancer`](#loadbalancer-eb605c) - Optional Block<br>Enable this option

<a id="https-auto-cert-pass-through"></a>&#x2022; [`pass_through`](#https-auto-cert-pass-through) - Optional Block<br>Enable this option

<a id="https-auto-cert-port"></a>&#x2022; [`port`](#https-auto-cert-port) - Optional Number<br>HTTPS Listen Port. HTTPS port to Listen

<a id="https-auto-cert-port-ranges"></a>&#x2022; [`port_ranges`](#https-auto-cert-port-ranges) - Optional String<br>Port Ranges. A string containing a comma separated list of port ranges. Each port range consists of a single port or two ports separated by '-'

<a id="https-auto-cert-server-name"></a>&#x2022; [`server_name`](#https-auto-cert-server-name) - Optional String<br>Modify header value. Define the header value for the header name “server”. This will overwrite existing values, if any, for the server header

<a id="https-auto-cert-tls-config"></a>&#x2022; [`tls_config`](#https-auto-cert-tls-config) - Optional Block<br>TLS Config. This defines various options to configure TLS configuration parameters<br>See [TLS Config](#https-auto-cert-tls-config) below.

<a id="https-auto-cert-use-mtls"></a>&#x2022; [`use_mtls`](#https-auto-cert-use-mtls) - Optional Block<br>Clients TLS validation context. Validation context for downstream client TLS connections<br>See [Use mTLS](#https-auto-cert-use-mtls) below.

#### HTTPS Auto Cert Coalescing Options

A [`coalescing_options`](#https-auto-cert-coalescing-options) block (within [`https_auto_cert`](#https-auto-cert)) supports the following:

<a id="coalescing-3c2270"></a>&#x2022; [`default_coalescing`](#coalescing-3c2270) - Optional Block<br>Enable this option

<a id="coalescing-010f02"></a>&#x2022; [`strict_coalescing`](#coalescing-010f02) - Optional Block<br>Enable this option

#### HTTPS Auto Cert HTTP Protocol Options

A [`http_protocol_options`](#https-auto-cert-http-protocol-options) block (within [`https_auto_cert`](#https-auto-cert)) supports the following:

<a id="only-d515de"></a>&#x2022; [`http_protocol_enable_v1_only`](#only-d515de) - Optional Block<br>HTTP/1.1 Protocol Options. HTTP/1.1 Protocol options for downstream connections<br>See [HTTP Protocol Enable V1 Only](#only-d515de) below.

<a id="v1-v2-9e0811"></a>&#x2022; [`http_protocol_enable_v1_v2`](#v1-v2-9e0811) - Optional Block<br>Enable this option

<a id="only-65e5e2"></a>&#x2022; [`http_protocol_enable_v2_only`](#only-65e5e2) - Optional Block<br>Enable this option

#### HTTPS Auto Cert HTTP Protocol Options HTTP Protocol Enable V1 Only

A [`http_protocol_enable_v1_only`](#only-d515de) block (within [`https_auto_cert.http_protocol_options`](#https-auto-cert-http-protocol-options)) supports the following:

<a id="transformation-8a0d5e"></a>&#x2022; [`header_transformation`](#transformation-8a0d5e) - Optional Block<br>Header Transformation. Header Transformation options for HTTP/1.1 request/response headers<br>See [Header Transformation](#transformation-8a0d5e) below.

#### HTTPS Auto Cert HTTP Protocol Options HTTP Protocol Enable V1 Only Header Transformation

A [`header_transformation`](#transformation-8a0d5e) block (within [`https_auto_cert.http_protocol_options.http_protocol_enable_v1_only`](#only-d515de)) supports the following:

<a id="transformation-36cba3"></a>&#x2022; [`default_header_transformation`](#transformation-36cba3) - Optional Block<br>Enable this option

<a id="transformation-491c67"></a>&#x2022; [`legacy_header_transformation`](#transformation-491c67) - Optional Block<br>Enable this option

<a id="transformation-3874df"></a>&#x2022; [`preserve_case_header_transformation`](#transformation-3874df) - Optional Block<br>Enable this option

<a id="transformation-1d5671"></a>&#x2022; [`proper_case_header_transformation`](#transformation-1d5671) - Optional Block<br>Enable this option

#### HTTPS Auto Cert TLS Config

A [`tls_config`](#https-auto-cert-tls-config) block (within [`https_auto_cert`](#https-auto-cert)) supports the following:

<a id="security-7a53da"></a>&#x2022; [`custom_security`](#security-7a53da) - Optional Block<br>Custom Ciphers. This defines TLS protocol config including min/max versions and allowed ciphers<br>See [Custom Security](#security-7a53da) below.

<a id="security-121c52"></a>&#x2022; [`default_security`](#security-121c52) - Optional Block<br>Enable this option

<a id="https-auto-cert-tls-config-low-security"></a>&#x2022; [`low_security`](#https-auto-cert-tls-config-low-security) - Optional Block<br>Enable this option

<a id="security-5e1ea1"></a>&#x2022; [`medium_security`](#security-5e1ea1) - Optional Block<br>Enable this option

#### HTTPS Auto Cert TLS Config Custom Security

A [`custom_security`](#security-7a53da) block (within [`https_auto_cert.tls_config`](#https-auto-cert-tls-config)) supports the following:

<a id="suites-c101c3"></a>&#x2022; [`cipher_suites`](#suites-c101c3) - Optional List<br>Cipher Suites. The TLS listener will only support the specified cipher list

<a id="version-f1d807"></a>&#x2022; [`max_version`](#version-f1d807) - Optional String  Defaults to `TLS_AUTO`<br>Possible values are `TLS_AUTO`, `TLSv1_0`, `TLSv1_1`, `TLSv1_2`, `TLSv1_3`<br>[Enum: TLS_AUTO|TLSv1_0|TLSv1_1|TLSv1_2|TLSv1_3] TLS Protocol. TlsProtocol is enumeration of supported TLS versions F5 Distributed Cloud will choose the optimal TLS version

<a id="version-b42e30"></a>&#x2022; [`min_version`](#version-b42e30) - Optional String  Defaults to `TLS_AUTO`<br>Possible values are `TLS_AUTO`, `TLSv1_0`, `TLSv1_1`, `TLSv1_2`, `TLSv1_3`<br>[Enum: TLS_AUTO|TLSv1_0|TLSv1_1|TLSv1_2|TLSv1_3] TLS Protocol. TlsProtocol is enumeration of supported TLS versions F5 Distributed Cloud will choose the optimal TLS version

#### HTTPS Auto Cert Use mTLS

An [`use_mtls`](#https-auto-cert-use-mtls) block (within [`https_auto_cert`](#https-auto-cert)) supports the following:

<a id="optional-4b03b7"></a>&#x2022; [`client_certificate_optional`](#optional-4b03b7) - Optional Bool<br>Client Certificate Optional. Client certificate is optional. If the client has provided a certificate, the load balancer will verify it. If certification verification fails, the connection will be terminated. If the client does not provide a certificate, the connection will be accepted

<a id="https-auto-cert-use-mtls-crl"></a>&#x2022; [`crl`](#https-auto-cert-use-mtls-crl) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [CRL](#https-auto-cert-use-mtls-crl) below.

<a id="https-auto-cert-use-mtls-no-crl"></a>&#x2022; [`no_crl`](#https-auto-cert-use-mtls-no-crl) - Optional Block<br>Enable this option

<a id="https-auto-cert-use-mtls-trusted-ca"></a>&#x2022; [`trusted_ca`](#https-auto-cert-use-mtls-trusted-ca) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Trusted CA](#https-auto-cert-use-mtls-trusted-ca) below.

<a id="https-auto-cert-use-mtls-trusted-ca-url"></a>&#x2022; [`trusted_ca_url`](#https-auto-cert-use-mtls-trusted-ca-url) - Optional String<br>Inline Root CA Certificate (legacy). Upload a Root CA Certificate specifically for this Load Balancer

<a id="https-auto-cert-use-mtls-xfcc-disabled"></a>&#x2022; [`xfcc_disabled`](#https-auto-cert-use-mtls-xfcc-disabled) - Optional Block<br>Enable this option

<a id="https-auto-cert-use-mtls-xfcc-options"></a>&#x2022; [`xfcc_options`](#https-auto-cert-use-mtls-xfcc-options) - Optional Block<br>XFCC Header Elements. X-Forwarded-Client-Cert header elements to be added to requests<br>See [Xfcc Options](#https-auto-cert-use-mtls-xfcc-options) below.

#### HTTPS Auto Cert Use mTLS CRL

A [`crl`](#https-auto-cert-use-mtls-crl) block (within [`https_auto_cert.use_mtls`](#https-auto-cert-use-mtls)) supports the following:

<a id="https-auto-cert-use-mtls-crl-name"></a>&#x2022; [`name`](#https-auto-cert-use-mtls-crl-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="https-auto-cert-use-mtls-crl-namespace"></a>&#x2022; [`namespace`](#https-auto-cert-use-mtls-crl-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="https-auto-cert-use-mtls-crl-tenant"></a>&#x2022; [`tenant`](#https-auto-cert-use-mtls-crl-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### HTTPS Auto Cert Use mTLS Trusted CA

A [`trusted_ca`](#https-auto-cert-use-mtls-trusted-ca) block (within [`https_auto_cert.use_mtls`](#https-auto-cert-use-mtls)) supports the following:

<a id="name-36a848"></a>&#x2022; [`name`](#name-36a848) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-1c0f1b"></a>&#x2022; [`namespace`](#namespace-1c0f1b) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-08b11a"></a>&#x2022; [`tenant`](#tenant-08b11a) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### HTTPS Auto Cert Use mTLS Xfcc Options

A [`xfcc_options`](#https-auto-cert-use-mtls-xfcc-options) block (within [`https_auto_cert.use_mtls`](#https-auto-cert-use-mtls)) supports the following:

<a id="elements-3cb07c"></a>&#x2022; [`xfcc_header_elements`](#elements-3cb07c) - Optional List  Defaults to `XFCC_NONE`<br>Possible values are `XFCC_NONE`, `XFCC_CERT`, `XFCC_CHAIN`, `XFCC_SUBJECT`, `XFCC_URI`, `XFCC_DNS`<br>[Enum: XFCC_NONE|XFCC_CERT|XFCC_CHAIN|XFCC_SUBJECT|XFCC_URI|XFCC_DNS] XFCC Header Elements. X-Forwarded-Client-Cert header elements to be added to requests

#### Js Challenge

A [`js_challenge`](#js-challenge) block supports the following:

<a id="js-challenge-cookie-expiry"></a>&#x2022; [`cookie_expiry`](#js-challenge-cookie-expiry) - Optional Number<br>Cookie Expiration Period. Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new challenge

<a id="js-challenge-custom-page"></a>&#x2022; [`custom_page`](#js-challenge-custom-page) - Optional String<br>Custom Message for Javascript Challenge. Custom message is of type uri_ref. Currently supported URL schemes is string:///. For string:/// scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. 'Please Wait.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '`<p>` Please Wait `</p>`'. Base64 encoded string for this HTML is 'PHA+IFBsZWFzZSBXYWl0IDwvcD4='

<a id="js-challenge-js-script-delay"></a>&#x2022; [`js_script_delay`](#js-challenge-js-script-delay) - Optional Number<br>Javascript Delay. Delay introduced by Javascript, in milliseconds

#### JWT Validation

A [`jwt_validation`](#jwt-validation) block supports the following:

<a id="jwt-validation-action"></a>&#x2022; [`action`](#jwt-validation-action) - Optional Block<br>Action<br>See [Action](#jwt-validation-action) below.

<a id="jwt-validation-jwks-config"></a>&#x2022; [`jwks_config`](#jwt-validation-jwks-config) - Optional Block<br>JSON Web Key Set (JWKS). The JSON Web Key Set (JWKS) is a set of keys used to verify JSON Web Token (JWT) issued by the Authorization Server. See RFC 7517 for more details<br>See [Jwks Config](#jwt-validation-jwks-config) below.

<a id="jwt-validation-mandatory-claims"></a>&#x2022; [`mandatory_claims`](#jwt-validation-mandatory-claims) - Optional Block<br>Mandatory Claims. Configurable Validation of mandatory Claims<br>See [Mandatory Claims](#jwt-validation-mandatory-claims) below.

<a id="jwt-validation-reserved-claims"></a>&#x2022; [`reserved_claims`](#jwt-validation-reserved-claims) - Optional Block<br>Reserved claims configuration. Configurable Validation of reserved Claims<br>See [Reserved Claims](#jwt-validation-reserved-claims) below.

<a id="jwt-validation-target"></a>&#x2022; [`target`](#jwt-validation-target) - Optional Block<br>Target. Define endpoints for which JWT token validation will be performed<br>See [Target](#jwt-validation-target) below.

<a id="jwt-validation-token-location"></a>&#x2022; [`token_location`](#jwt-validation-token-location) - Optional Block<br>Token Location. Location of JWT in HTTP request<br>See [Token Location](#jwt-validation-token-location) below.

#### JWT Validation Action

An [`action`](#jwt-validation-action) block (within [`jwt_validation`](#jwt-validation)) supports the following:

<a id="jwt-validation-action-block"></a>&#x2022; [`block`](#jwt-validation-action-block) - Optional Block<br>Enable this option

<a id="jwt-validation-action-report"></a>&#x2022; [`report`](#jwt-validation-action-report) - Optional Block<br>Enable this option

#### JWT Validation Jwks Config

A [`jwks_config`](#jwt-validation-jwks-config) block (within [`jwt_validation`](#jwt-validation)) supports the following:

<a id="jwt-validation-jwks-config-cleartext"></a>&#x2022; [`cleartext`](#jwt-validation-jwks-config-cleartext) - Optional String<br>JSON Web Key Set (JWKS). The JSON Web Key Set (JWKS) is a set of keys used to verify JSON Web Token (JWT) issued by the Authorization Server. See RFC 7517 for more details

#### JWT Validation Mandatory Claims

A [`mandatory_claims`](#jwt-validation-mandatory-claims) block (within [`jwt_validation`](#jwt-validation)) supports the following:

<a id="names-2ccfbe"></a>&#x2022; [`claim_names`](#names-2ccfbe) - Optional List<br>Claim Names

#### JWT Validation Reserved Claims

A [`reserved_claims`](#jwt-validation-reserved-claims) block (within [`jwt_validation`](#jwt-validation)) supports the following:

<a id="jwt-validation-reserved-claims-audience"></a>&#x2022; [`audience`](#jwt-validation-reserved-claims-audience) - Optional Block<br>Audiences<br>See [Audience](#jwt-validation-reserved-claims-audience) below.

<a id="disable-dcfb50"></a>&#x2022; [`audience_disable`](#disable-dcfb50) - Optional Block<br>Enable this option

<a id="jwt-validation-reserved-claims-issuer"></a>&#x2022; [`issuer`](#jwt-validation-reserved-claims-issuer) - Optional String<br>Exact Match

<a id="disable-c89c1c"></a>&#x2022; [`issuer_disable`](#disable-c89c1c) - Optional Block<br>Enable this option

<a id="disable-5d3cb1"></a>&#x2022; [`validate_period_disable`](#disable-5d3cb1) - Optional Block<br>Enable this option

<a id="enable-66243b"></a>&#x2022; [`validate_period_enable`](#enable-66243b) - Optional Block<br>Enable this option

#### JWT Validation Reserved Claims Audience

An [`audience`](#jwt-validation-reserved-claims-audience) block (within [`jwt_validation.reserved_claims`](#jwt-validation-reserved-claims)) supports the following:

<a id="audiences-a34853"></a>&#x2022; [`audiences`](#audiences-a34853) - Optional List<br>Values

#### JWT Validation Target

A [`target`](#jwt-validation-target) block (within [`jwt_validation`](#jwt-validation)) supports the following:

<a id="jwt-validation-target-all-endpoint"></a>&#x2022; [`all_endpoint`](#jwt-validation-target-all-endpoint) - Optional Block<br>Enable this option

<a id="jwt-validation-target-api-groups"></a>&#x2022; [`api_groups`](#jwt-validation-target-api-groups) - Optional Block<br>API Groups<br>See [API Groups](#jwt-validation-target-api-groups) below.

<a id="jwt-validation-target-base-paths"></a>&#x2022; [`base_paths`](#jwt-validation-target-base-paths) - Optional Block<br>Base Paths<br>See [Base Paths](#jwt-validation-target-base-paths) below.

#### JWT Validation Target API Groups

An [`api_groups`](#jwt-validation-target-api-groups) block (within [`jwt_validation.target`](#jwt-validation-target)) supports the following:

<a id="groups-057782"></a>&#x2022; [`api_groups`](#groups-057782) - Optional List<br>API Groups

#### JWT Validation Target Base Paths

A [`base_paths`](#jwt-validation-target-base-paths) block (within [`jwt_validation.target`](#jwt-validation-target)) supports the following:

<a id="paths-b433d5"></a>&#x2022; [`base_paths`](#paths-b433d5) - Optional List<br>Prefix Values

#### JWT Validation Token Location

A [`token_location`](#jwt-validation-token-location) block (within [`jwt_validation`](#jwt-validation)) supports the following:

<a id="token-e5c0e3"></a>&#x2022; [`bearer_token`](#token-e5c0e3) - Optional Block<br>Enable this option

#### L7 DDOS Action Js Challenge

A [`l7_ddos_action_js_challenge`](#l7-ddos-action-js-challenge) block supports the following:

<a id="expiry-2697a0"></a>&#x2022; [`cookie_expiry`](#expiry-2697a0) - Optional Number<br>Cookie Expiration Period. Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new challenge

<a id="l7-ddos-action-js-challenge-custom-page"></a>&#x2022; [`custom_page`](#l7-ddos-action-js-challenge-custom-page) - Optional String<br>Custom Message for Javascript Challenge. Custom message is of type uri_ref. Currently supported URL schemes is string:///. For string:/// scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. 'Please Wait.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '`<p>` Please Wait `</p>`'. Base64 encoded string for this HTML is 'PHA+IFBsZWFzZSBXYWl0IDwvcD4='

<a id="delay-88f51d"></a>&#x2022; [`js_script_delay`](#delay-88f51d) - Optional Number<br>Javascript Delay. Delay introduced by Javascript, in milliseconds

#### L7 DDOS Protection

A [`l7_ddos_protection`](#l7-ddos-protection) block supports the following:

<a id="challenge-84ab9e"></a>&#x2022; [`clientside_action_captcha_challenge`](#challenge-84ab9e) - Optional Block<br>Captcha Challenge Parameters. Enables loadbalancer to perform captcha challenge Captcha challenge will be based on Google Recaptcha. With this feature enabled, only clients that pass the captcha challenge will be allowed to complete the HTTP request. When loadbalancer is configured to do Captcha Challenge, it will redirect the browser to an HTML page on every new HTTP request. This HTML page will have captcha challenge embedded in it. Client will be allowed to make the request only if the captcha challenge is successful. Loadbalancer will tag response header with a cookie to avoid Captcha challenge for subsequent requests. CAPTCHA is mainly used as a security check to ensure only human users can pass through. Generally, computers or bots are not capable of solving a captcha. You can enable either Javascript challenge or Captcha challenge on a virtual host<br>See [Clientside Action Captcha Challenge](#challenge-84ab9e) below.

<a id="challenge-1070c2"></a>&#x2022; [`clientside_action_js_challenge`](#challenge-1070c2) - Optional Block<br>Javascript Challenge Parameters. Enables loadbalancer to perform client browser compatibility test by redirecting to a page with Javascript. With this feature enabled, only clients that are capable of executing Javascript(mostly browsers) will be allowed to complete the HTTP request. When loadbalancer is configured to do Javascript Challenge, it will redirect the browser to an HTML page on every new HTTP request. This HTML page will have Javascript embedded in it. Loadbalancer chooses a set of random numbers for every new client and sends these numbers along with an encrypted answer with the request such that it embed these numbers as input in the Javascript. Javascript will run on the requestor browser and perform a complex Math operation. Script will submit the answer to loadbalancer. Loadbalancer will validate the answer by comparing the calculated answer with the decrypted answer (which was encrypted when it was sent back as reply) and allow the request to the upstream server only if the answer is correct. Loadbalancer will tag response header with a cookie to avoid Javascript challenge for subsequent requests. Javascript challenge serves following purposes \* Validate that the request is coming via a browser that is capable for running Javascript \* Force the browser to run a complex operation, f(X), that requires it to spend a large number of CPU cycles. This is to slow down a potential DOS attacker by making it difficult to launch a large request flood without having to spend even larger CPU cost at their end. You can enable either Javascript challenge or Captcha challenge on a virtual host<br>See [Clientside Action Js Challenge](#challenge-1070c2) below.

<a id="none-88961b"></a>&#x2022; [`clientside_action_none`](#none-88961b) - Optional Block<br>Enable this option

<a id="l7-ddos-protection-ddos-policy-custom"></a>&#x2022; [`ddos_policy_custom`](#l7-ddos-protection-ddos-policy-custom) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [DDOS Policy Custom](#l7-ddos-protection-ddos-policy-custom) below.

<a id="l7-ddos-protection-ddos-policy-none"></a>&#x2022; [`ddos_policy_none`](#l7-ddos-protection-ddos-policy-none) - Optional Block<br>Enable this option

<a id="threshold-332758"></a>&#x2022; [`default_rps_threshold`](#threshold-332758) - Optional Block<br>Enable this option

<a id="l7-ddos-protection-mitigation-block"></a>&#x2022; [`mitigation_block`](#l7-ddos-protection-mitigation-block) - Optional Block<br>Enable this option

<a id="challenge-62fb67"></a>&#x2022; [`mitigation_captcha_challenge`](#challenge-62fb67) - Optional Block<br>Captcha Challenge Parameters. Enables loadbalancer to perform captcha challenge Captcha challenge will be based on Google Recaptcha. With this feature enabled, only clients that pass the captcha challenge will be allowed to complete the HTTP request. When loadbalancer is configured to do Captcha Challenge, it will redirect the browser to an HTML page on every new HTTP request. This HTML page will have captcha challenge embedded in it. Client will be allowed to make the request only if the captcha challenge is successful. Loadbalancer will tag response header with a cookie to avoid Captcha challenge for subsequent requests. CAPTCHA is mainly used as a security check to ensure only human users can pass through. Generally, computers or bots are not capable of solving a captcha. You can enable either Javascript challenge or Captcha challenge on a virtual host<br>See [Mitigation Captcha Challenge](#challenge-62fb67) below.

<a id="challenge-2a2755"></a>&#x2022; [`mitigation_js_challenge`](#challenge-2a2755) - Optional Block<br>Javascript Challenge Parameters. Enables loadbalancer to perform client browser compatibility test by redirecting to a page with Javascript. With this feature enabled, only clients that are capable of executing Javascript(mostly browsers) will be allowed to complete the HTTP request. When loadbalancer is configured to do Javascript Challenge, it will redirect the browser to an HTML page on every new HTTP request. This HTML page will have Javascript embedded in it. Loadbalancer chooses a set of random numbers for every new client and sends these numbers along with an encrypted answer with the request such that it embed these numbers as input in the Javascript. Javascript will run on the requestor browser and perform a complex Math operation. Script will submit the answer to loadbalancer. Loadbalancer will validate the answer by comparing the calculated answer with the decrypted answer (which was encrypted when it was sent back as reply) and allow the request to the upstream server only if the answer is correct. Loadbalancer will tag response header with a cookie to avoid Javascript challenge for subsequent requests. Javascript challenge serves following purposes \* Validate that the request is coming via a browser that is capable for running Javascript \* Force the browser to run a complex operation, f(X), that requires it to spend a large number of CPU cycles. This is to slow down a potential DOS attacker by making it difficult to launch a large request flood without having to spend even larger CPU cost at their end. You can enable either Javascript challenge or Captcha challenge on a virtual host<br>See [Mitigation Js Challenge](#challenge-2a2755) below.

<a id="l7-ddos-protection-rps-threshold"></a>&#x2022; [`rps_threshold`](#l7-ddos-protection-rps-threshold) - Optional Number<br>Custom. Configure custom RPS threshold

#### L7 DDOS Protection Clientside Action Captcha Challenge

A [`clientside_action_captcha_challenge`](#challenge-84ab9e) block (within [`l7_ddos_protection`](#l7-ddos-protection)) supports the following:

<a id="expiry-758337"></a>&#x2022; [`cookie_expiry`](#expiry-758337) - Optional Number<br>Cookie Expiration Period. Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new challenge

<a id="page-774877"></a>&#x2022; [`custom_page`](#page-774877) - Optional String<br>Custom message for Captcha Challenge. Custom message is of type uri_ref. Currently supported URL schemes is string:///. For string:/// scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. 'Please Wait.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '`<p>` Please Wait `</p>`'. Base64 encoded string for this HTML is 'PHA+IFBsZWFzZSBXYWl0IDwvcD4='

#### L7 DDOS Protection Clientside Action Js Challenge

A [`clientside_action_js_challenge`](#challenge-1070c2) block (within [`l7_ddos_protection`](#l7-ddos-protection)) supports the following:

<a id="expiry-fb1ca5"></a>&#x2022; [`cookie_expiry`](#expiry-fb1ca5) - Optional Number<br>Cookie Expiration Period. Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new challenge

<a id="page-957fdc"></a>&#x2022; [`custom_page`](#page-957fdc) - Optional String<br>Custom Message for Javascript Challenge. Custom message is of type uri_ref. Currently supported URL schemes is string:///. For string:/// scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. 'Please Wait.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '`<p>` Please Wait `</p>`'. Base64 encoded string for this HTML is 'PHA+IFBsZWFzZSBXYWl0IDwvcD4='

<a id="delay-cc7ed2"></a>&#x2022; [`js_script_delay`](#delay-cc7ed2) - Optional Number<br>Javascript Delay. Delay introduced by Javascript, in milliseconds

#### L7 DDOS Protection DDOS Policy Custom

A [`ddos_policy_custom`](#l7-ddos-protection-ddos-policy-custom) block (within [`l7_ddos_protection`](#l7-ddos-protection)) supports the following:

<a id="name-21c5e6"></a>&#x2022; [`name`](#name-21c5e6) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-622fcf"></a>&#x2022; [`namespace`](#namespace-622fcf) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-e6ac4d"></a>&#x2022; [`tenant`](#tenant-e6ac4d) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### L7 DDOS Protection Mitigation Captcha Challenge

A [`mitigation_captcha_challenge`](#challenge-62fb67) block (within [`l7_ddos_protection`](#l7-ddos-protection)) supports the following:

<a id="expiry-dcbde2"></a>&#x2022; [`cookie_expiry`](#expiry-dcbde2) - Optional Number<br>Cookie Expiration Period. Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new challenge

<a id="page-0f52d0"></a>&#x2022; [`custom_page`](#page-0f52d0) - Optional String<br>Custom message for Captcha Challenge. Custom message is of type uri_ref. Currently supported URL schemes is string:///. For string:/// scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. 'Please Wait.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '`<p>` Please Wait `</p>`'. Base64 encoded string for this HTML is 'PHA+IFBsZWFzZSBXYWl0IDwvcD4='

#### L7 DDOS Protection Mitigation Js Challenge

A [`mitigation_js_challenge`](#challenge-2a2755) block (within [`l7_ddos_protection`](#l7-ddos-protection)) supports the following:

<a id="expiry-299338"></a>&#x2022; [`cookie_expiry`](#expiry-299338) - Optional Number<br>Cookie Expiration Period. Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new challenge

<a id="page-180c41"></a>&#x2022; [`custom_page`](#page-180c41) - Optional String<br>Custom Message for Javascript Challenge. Custom message is of type uri_ref. Currently supported URL schemes is string:///. For string:/// scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. 'Please Wait.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '`<p>` Please Wait `</p>`'. Base64 encoded string for this HTML is 'PHA+IFBsZWFzZSBXYWl0IDwvcD4='

<a id="delay-9c5ffc"></a>&#x2022; [`js_script_delay`](#delay-9c5ffc) - Optional Number<br>Javascript Delay. Delay introduced by Javascript, in milliseconds

#### Malware Protection Settings

A [`malware_protection_settings`](#malware-protection-settings) block supports the following:

<a id="rules-b2bf3e"></a>&#x2022; [`malware_protection_rules`](#rules-b2bf3e) - Optional Block<br>Malware Detection Rules. Configure the match criteria to trigger Malware Protection Scan<br>See [Malware Protection Rules](#rules-b2bf3e) below.

#### Malware Protection Settings Malware Protection Rules

A [`malware_protection_rules`](#rules-b2bf3e) block (within [`malware_protection_settings`](#malware-protection-settings)) supports the following:

<a id="action-f0dc04"></a>&#x2022; [`action`](#action-f0dc04) - Optional Block<br>Action<br>See [Action](#action-f0dc04) below.

<a id="domain-7b5aea"></a>&#x2022; [`domain`](#domain-7b5aea) - Optional Block<br>Domain to Match. Domain to be matched<br>See [Domain](#domain-7b5aea) below.

<a id="methods-02fcb9"></a>&#x2022; [`http_methods`](#methods-02fcb9) - Optional List  Defaults to `ANY`<br>Possible values are `ANY`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`, `COPY`<br>[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY] HTTP Methods. Methods to be matched

<a id="metadata-20b73a"></a>&#x2022; [`metadata`](#metadata-20b73a) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#metadata-20b73a) below.

<a id="path-ce187c"></a>&#x2022; [`path`](#path-ce187c) - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#path-ce187c) below.

#### Malware Protection Settings Malware Protection Rules Action

An [`action`](#action-f0dc04) block (within [`malware_protection_settings.malware_protection_rules`](#rules-b2bf3e)) supports the following:

<a id="block-97253f"></a>&#x2022; [`block`](#block-97253f) - Optional Block<br>Enable this option

<a id="report-055d2b"></a>&#x2022; [`report`](#report-055d2b) - Optional Block<br>Enable this option

#### Malware Protection Settings Malware Protection Rules Domain

A [`domain`](#domain-7b5aea) block (within [`malware_protection_settings.malware_protection_rules`](#rules-b2bf3e)) supports the following:

<a id="domain-5794b8"></a>&#x2022; [`any_domain`](#domain-5794b8) - Optional Block<br>Enable this option

<a id="domain-168628"></a>&#x2022; [`domain`](#domain-168628) - Optional Block<br>Domains. Domains names<br>See [Domain](#domain-168628) below.

#### Malware Protection Settings Malware Protection Rules Domain Domain

A [`domain`](#domain-168628) block (within [`malware_protection_settings.malware_protection_rules.domain`](#domain-7b5aea)) supports the following:

<a id="value-d1c5e9"></a>&#x2022; [`exact_value`](#value-d1c5e9) - Optional String<br>Exact Value. Exact domain name

<a id="value-b30045"></a>&#x2022; [`regex_value`](#value-b30045) - Optional String<br>Regex Values of Domains. Regular Expression value for the domain name

<a id="value-332cf0"></a>&#x2022; [`suffix_value`](#value-332cf0) - Optional String<br>Suffix Value. Suffix of domain name e.g 'xyz.com' will match '*.xyz.com' and 'xyz.com'

#### Malware Protection Settings Malware Protection Rules Metadata

A [`metadata`](#metadata-20b73a) block (within [`malware_protection_settings.malware_protection_rules`](#rules-b2bf3e)) supports the following:

<a id="spec-d9f05e"></a>&#x2022; [`description_spec`](#spec-d9f05e) - Optional String<br>Description. Human readable description

<a id="name-34fca6"></a>&#x2022; [`name`](#name-34fca6) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### Malware Protection Settings Malware Protection Rules Path

A [`path`](#path-ce187c) block (within [`malware_protection_settings.malware_protection_rules`](#rules-b2bf3e)) supports the following:

<a id="path-2aac29"></a>&#x2022; [`path`](#path-2aac29) - Optional String<br>Exact. Exact path value to match

<a id="prefix-4ddbaa"></a>&#x2022; [`prefix`](#prefix-4ddbaa) - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

<a id="regex-56e93f"></a>&#x2022; [`regex`](#regex-56e93f) - Optional String<br>Regex. Regular expression of path match (e.g. the value .* will match on all paths)

#### More Option

A [`more_option`](#more-option) block supports the following:

<a id="more-option-buffer-policy"></a>&#x2022; [`buffer_policy`](#more-option-buffer-policy) - Optional Block<br>Buffer Configuration. Some upstream applications are not capable of handling streamed data. This config enables buffering the entire request before sending to upstream application. We can specify the maximum buffer size and buffer interval with this config. Buffering can be enabled and disabled at VirtualHost and Route levels Route level buffer configuration takes precedence<br>See [Buffer Policy](#more-option-buffer-policy) below.

<a id="more-option-compression-params"></a>&#x2022; [`compression_params`](#more-option-compression-params) - Optional Block<br>Compression Parameters. Enables loadbalancer to compress dispatched data from an upstream service upon client request. The content is compressed and then sent to the client with the appropriate headers if either response and request allow. Only GZIP compression is supported. By default compression will be skipped when: A request does NOT contain accept-encoding header. A request includes accept-encoding header, but it does not contain “gzip” or “*”. A request includes accept-encoding with “gzip” or “*” with the weight “q=0”. Note that the “gzip” will have a higher weight then “*”. For example, if accept-encoding is “gzip;q=0,*;q=1”, the filter will not compress. But if the header is set to “*;q=0,gzip;q=1”, the filter will compress. A request whose accept-encoding header includes “identity”. A response contains a content-encoding header. A response contains a cache-control header whose value includes “no-transform”. A response contains a transfer-encoding header whose value includes “gzip”. A response does not contain a content-type value that matches one of the selected mime-types, which default to application/javascript, application/JSON, application/xhtml+XML, image/svg+XML, text/CSS, text/HTML, text/plain, text/XML. Neither content-length nor transfer-encoding headers are present in the response. Response size is smaller than 30 bytes (only applicable when transfer-encoding is not chunked). When compression is applied: The content-length is removed from response headers. Response headers contain “transfer-encoding: chunked” and do not contain “content-encoding” header. The “vary: accept-encoding” header is inserted on every response. GZIP Compression Level: A value which is optimal balance between speed of compression and amount of compression is chosen<br>See [Compression Params](#more-option-compression-params) below.

<a id="more-option-custom-errors"></a>&#x2022; [`custom_errors`](#more-option-custom-errors) - Optional Block<br>Custom Error Responses. Map of integer error codes as keys and string values that can be used to provide custom HTTP pages for each error code. Key of the map can be either response code class or HTTP Error code. Response code classes for key is configured as follows 3 -- for 3xx response code class 4 -- for 4xx response code class 5 -- for 5xx response code class Value of the map is string which represents custom HTTP responses. Specific response code takes preference when both response code and response code class matches for a request

<a id="more-option-disable-default-error-pages"></a>&#x2022; [`disable_default_error_pages`](#more-option-disable-default-error-pages) - Optional Bool<br>Disable Default Error Pages. Disable the use of default F5XC error pages

<a id="more-option-disable-path-normalize"></a>&#x2022; [`disable_path_normalize`](#more-option-disable-path-normalize) - Optional Block<br>Enable this option

<a id="more-option-enable-path-normalize"></a>&#x2022; [`enable_path_normalize`](#more-option-enable-path-normalize) - Optional Block<br>Enable this option

<a id="more-option-idle-timeout"></a>&#x2022; [`idle_timeout`](#more-option-idle-timeout) - Optional Number<br>Idle Timeout. The amount of time that a stream can exist without upstream or downstream activity, in milliseconds. The stream is terminated with a HTTP 504 (Gateway Timeout) error code if no upstream response header has been received, otherwise the stream is reset

<a id="more-option-max-request-header-size"></a>&#x2022; [`max_request_header_size`](#more-option-max-request-header-size) - Optional Number<br>Maximum Request Header Size. The maximum request header size for downstream connections, in KiB. A HTTP 431 (Request Header Fields Too Large) error code is sent for requests that exceed this size. If multiple load balancers share the same advertise_policy, the highest value configured across all such load balancers is used for all the load balancers in question

<a id="more-option-request-cookies-to-add"></a>&#x2022; [`request_cookies_to_add`](#more-option-request-cookies-to-add) - Optional Block<br>Add Cookies in Cookie Header. Cookies are key-value pairs to be added to HTTP request being routed towards upstream. Cookies specified at this level are applied after cookies from matched Route are applied<br>See [Request Cookies To Add](#more-option-request-cookies-to-add) below.

<a id="more-option-request-cookies-to-remove"></a>&#x2022; [`request_cookies_to_remove`](#more-option-request-cookies-to-remove) - Optional List<br>Remove Cookies from Cookie Header. List of keys of Cookies to be removed from the HTTP request being sent towards upstream

<a id="more-option-request-headers-to-add"></a>&#x2022; [`request_headers_to_add`](#more-option-request-headers-to-add) - Optional Block<br>Add Request Headers. Headers are key-value pairs to be added to HTTP request being routed towards upstream. Headers specified at this level are applied after headers from matched Route are applied<br>See [Request Headers To Add](#more-option-request-headers-to-add) below.

<a id="more-option-request-headers-to-remove"></a>&#x2022; [`request_headers_to_remove`](#more-option-request-headers-to-remove) - Optional List<br>Remove Request Headers. List of keys of Headers to be removed from the HTTP request being sent towards upstream

<a id="more-option-response-cookies-to-add"></a>&#x2022; [`response_cookies_to_add`](#more-option-response-cookies-to-add) - Optional Block<br>Add Set-Cookie Headers. Cookies are name-value pairs along with optional attribute parameters to be added to HTTP response being sent towards downstream. Cookies specified at this level are applied after cookies from matched Route are applied<br>See [Response Cookies To Add](#more-option-response-cookies-to-add) below.

<a id="more-option-response-cookies-to-remove"></a>&#x2022; [`response_cookies_to_remove`](#more-option-response-cookies-to-remove) - Optional List<br>Remove Cookies from Set-Cookie Headers. List of name of Cookies to be removed from the HTTP response being sent towards downstream. Entire set-cookie header will be removed

<a id="more-option-response-headers-to-add"></a>&#x2022; [`response_headers_to_add`](#more-option-response-headers-to-add) - Optional Block<br>Add Response Headers. Headers are key-value pairs to be added to HTTP response being sent towards downstream. Headers specified at this level are applied after headers from matched Route are applied<br>See [Response Headers To Add](#more-option-response-headers-to-add) below.

<a id="more-option-response-headers-to-remove"></a>&#x2022; [`response_headers_to_remove`](#more-option-response-headers-to-remove) - Optional List<br>Remove Response Headers. List of keys of Headers to be removed from the HTTP response being sent towards downstream

#### More Option Buffer Policy

A [`buffer_policy`](#more-option-buffer-policy) block (within [`more_option`](#more-option)) supports the following:

<a id="more-option-buffer-policy-disabled"></a>&#x2022; [`disabled`](#more-option-buffer-policy-disabled) - Optional Bool<br>Disable. Disable buffering for a particular route. This is useful when virtual-host has buffering, but we need to disable it on a specific route. The value of this field is ignored for virtual-host

<a id="bytes-49c09c"></a>&#x2022; [`max_request_bytes`](#bytes-49c09c) - Optional Number<br>Max Request Bytes. The maximum request size that the filter will buffer before the connection manager will stop buffering and return a RequestEntityTooLarge (413) response

#### More Option Compression Params

A [`compression_params`](#more-option-compression-params) block (within [`more_option`](#more-option)) supports the following:

<a id="length-5f90b4"></a>&#x2022; [`content_length`](#length-5f90b4) - Optional Number  Defaults to `30`<br>Content Length. Minimum response length, in bytes, which will trigger compression. The

<a id="type-163d4d"></a>&#x2022; [`content_type`](#type-163d4d) - Optional List<br>Content Type. Set of strings that allows specifying which mime-types yield compression When this field is not defined, compression will be applied to the following mime-types: 'application/javascript' 'application/JSON', 'application/xhtml+XML' 'image/svg+XML' 'text/CSS' 'text/HTML' 'text/plain' 'text/XML'

<a id="header-8c49b4"></a>&#x2022; [`disable_on_etag_header`](#header-8c49b4) - Optional Bool<br>Disable On Etag Header. If true, disables compression when the response contains an etag header. When it is false, weak etags will be preserved and the ones that require strong validation will be removed

<a id="header-45cce4"></a>&#x2022; [`remove_accept_encoding_header`](#header-45cce4) - Optional Bool<br>Remove Accept-Encoding Header. If true, removes accept-encoding from the request headers before dispatching it to the upstream so that responses do not get compressed before reaching the filter

#### More Option Request Cookies To Add

A [`request_cookies_to_add`](#more-option-request-cookies-to-add) block (within [`more_option`](#more-option)) supports the following:

<a id="more-option-request-cookies-to-add-name"></a>&#x2022; [`name`](#more-option-request-cookies-to-add-name) - Optional String<br>Name. Name of the cookie in Cookie header

<a id="overwrite-6d9c60"></a>&#x2022; [`overwrite`](#overwrite-6d9c60) - Optional Bool  Defaults to `do`<br>Overwrite. Should the value be overwritten? If true, the value is overwritten to existing values. not overwrite

<a id="value-2d505e"></a>&#x2022; [`secret_value`](#value-2d505e) - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Secret Value](#value-2d505e) below.

<a id="value-9c117c"></a>&#x2022; [`value`](#value-9c117c) - Optional String<br>Value. Value of the Cookie header

#### More Option Request Cookies To Add Secret Value

A [`secret_value`](#value-2d505e) block (within [`more_option.request_cookies_to_add`](#more-option-request-cookies-to-add)) supports the following:

<a id="info-63d8c8"></a>&#x2022; [`blindfold_secret_info`](#info-63d8c8) - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#info-63d8c8) below.

<a id="info-a40feb"></a>&#x2022; [`clear_secret_info`](#info-a40feb) - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#info-a40feb) below.

#### More Option Request Cookies To Add Secret Value Blindfold Secret Info

A [`blindfold_secret_info`](#info-63d8c8) block (within [`more_option.request_cookies_to_add.secret_value`](#value-2d505e)) supports the following:

<a id="provider-24ac9b"></a>&#x2022; [`decryption_provider`](#provider-24ac9b) - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

<a id="location-14eb0e"></a>&#x2022; [`location`](#location-14eb0e) - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

<a id="provider-65257a"></a>&#x2022; [`store_provider`](#provider-65257a) - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

#### More Option Request Cookies To Add Secret Value Clear Secret Info

A [`clear_secret_info`](#info-a40feb) block (within [`more_option.request_cookies_to_add.secret_value`](#value-2d505e)) supports the following:

<a id="ref-2107b7"></a>&#x2022; [`provider_ref`](#ref-2107b7) - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

<a id="url-d34f46"></a>&#x2022; [`url`](#url-d34f46) - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

#### More Option Request Headers To Add

A [`request_headers_to_add`](#more-option-request-headers-to-add) block (within [`more_option`](#more-option)) supports the following:

<a id="append-2047af"></a>&#x2022; [`append`](#append-2047af) - Optional Bool  Defaults to `do`<br>Append. Should the value be appended? If true, the value is appended to existing values. not append

<a id="more-option-request-headers-to-add-name"></a>&#x2022; [`name`](#more-option-request-headers-to-add-name) - Optional String<br>Name. Name of the HTTP header

<a id="value-d68008"></a>&#x2022; [`secret_value`](#value-d68008) - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Secret Value](#value-d68008) below.

<a id="value-7dbd74"></a>&#x2022; [`value`](#value-7dbd74) - Optional String<br>Value. Value of the HTTP header

#### More Option Request Headers To Add Secret Value

A [`secret_value`](#value-d68008) block (within [`more_option.request_headers_to_add`](#more-option-request-headers-to-add)) supports the following:

<a id="info-42792e"></a>&#x2022; [`blindfold_secret_info`](#info-42792e) - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#info-42792e) below.

<a id="info-052c46"></a>&#x2022; [`clear_secret_info`](#info-052c46) - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#info-052c46) below.

#### More Option Request Headers To Add Secret Value Blindfold Secret Info

A [`blindfold_secret_info`](#info-42792e) block (within [`more_option.request_headers_to_add.secret_value`](#value-d68008)) supports the following:

<a id="provider-818163"></a>&#x2022; [`decryption_provider`](#provider-818163) - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

<a id="location-39d072"></a>&#x2022; [`location`](#location-39d072) - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

<a id="provider-745200"></a>&#x2022; [`store_provider`](#provider-745200) - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

#### More Option Request Headers To Add Secret Value Clear Secret Info

A [`clear_secret_info`](#info-052c46) block (within [`more_option.request_headers_to_add.secret_value`](#value-d68008)) supports the following:

<a id="ref-15a211"></a>&#x2022; [`provider_ref`](#ref-15a211) - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

<a id="url-2e7228"></a>&#x2022; [`url`](#url-2e7228) - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

#### More Option Response Cookies To Add

A [`response_cookies_to_add`](#more-option-response-cookies-to-add) block (within [`more_option`](#more-option)) supports the following:

<a id="domain-90efef"></a>&#x2022; [`add_domain`](#domain-90efef) - Optional String<br>Add Domain. Add domain attribute

<a id="expiry-5d54ee"></a>&#x2022; [`add_expiry`](#expiry-5d54ee) - Optional String<br>Add expiry. Add expiry attribute

<a id="httponly-8439bf"></a>&#x2022; [`add_httponly`](#httponly-8439bf) - Optional Block<br>Enable this option

<a id="partitioned-781e49"></a>&#x2022; [`add_partitioned`](#partitioned-781e49) - Optional Block<br>Enable this option

<a id="path-e18695"></a>&#x2022; [`add_path`](#path-e18695) - Optional String<br>Add path. Add path attribute

<a id="secure-e3baa0"></a>&#x2022; [`add_secure`](#secure-e3baa0) - Optional Block<br>Enable this option

<a id="domain-6328c3"></a>&#x2022; [`ignore_domain`](#domain-6328c3) - Optional Block<br>Enable this option

<a id="expiry-49396f"></a>&#x2022; [`ignore_expiry`](#expiry-49396f) - Optional Block<br>Enable this option

<a id="httponly-6fac42"></a>&#x2022; [`ignore_httponly`](#httponly-6fac42) - Optional Block<br>Enable this option

<a id="age-d6a859"></a>&#x2022; [`ignore_max_age`](#age-d6a859) - Optional Block<br>Enable this option

<a id="partitioned-f4bce0"></a>&#x2022; [`ignore_partitioned`](#partitioned-f4bce0) - Optional Block<br>Enable this option

<a id="path-f6b7e0"></a>&#x2022; [`ignore_path`](#path-f6b7e0) - Optional Block<br>Enable this option

<a id="samesite-aceec3"></a>&#x2022; [`ignore_samesite`](#samesite-aceec3) - Optional Block<br>Enable this option

<a id="secure-fe1099"></a>&#x2022; [`ignore_secure`](#secure-fe1099) - Optional Block<br>Enable this option

<a id="value-bbe342"></a>&#x2022; [`ignore_value`](#value-bbe342) - Optional Block<br>Enable this option

<a id="value-cb17ee"></a>&#x2022; [`max_age_value`](#value-cb17ee) - Optional Number<br>Add Max Age. Add max age attribute

<a id="name-9a2258"></a>&#x2022; [`name`](#name-9a2258) - Optional String<br>Name. Name of the cookie in Cookie header

<a id="overwrite-16498a"></a>&#x2022; [`overwrite`](#overwrite-16498a) - Optional Bool  Defaults to `do`<br>Overwrite. Should the value be overwritten? If true, the value is overwritten to existing values. not overwrite

<a id="lax-9d99f8"></a>&#x2022; [`samesite_lax`](#lax-9d99f8) - Optional Block<br>Enable this option

<a id="none-4de9cb"></a>&#x2022; [`samesite_none`](#none-4de9cb) - Optional Block<br>Enable this option

<a id="strict-d87273"></a>&#x2022; [`samesite_strict`](#strict-d87273) - Optional Block<br>Enable this option

<a id="value-3191e4"></a>&#x2022; [`secret_value`](#value-3191e4) - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Secret Value](#value-3191e4) below.

<a id="value-ff8a1e"></a>&#x2022; [`value`](#value-ff8a1e) - Optional String<br>Value. Value of the Cookie header

#### More Option Response Cookies To Add Secret Value

A [`secret_value`](#value-3191e4) block (within [`more_option.response_cookies_to_add`](#more-option-response-cookies-to-add)) supports the following:

<a id="info-dfeddd"></a>&#x2022; [`blindfold_secret_info`](#info-dfeddd) - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#info-dfeddd) below.

<a id="info-55a3e0"></a>&#x2022; [`clear_secret_info`](#info-55a3e0) - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#info-55a3e0) below.

#### More Option Response Cookies To Add Secret Value Blindfold Secret Info

A [`blindfold_secret_info`](#info-dfeddd) block (within [`more_option.response_cookies_to_add.secret_value`](#value-3191e4)) supports the following:

<a id="provider-d9b0df"></a>&#x2022; [`decryption_provider`](#provider-d9b0df) - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

<a id="location-c7d2a1"></a>&#x2022; [`location`](#location-c7d2a1) - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

<a id="provider-22fd33"></a>&#x2022; [`store_provider`](#provider-22fd33) - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

#### More Option Response Cookies To Add Secret Value Clear Secret Info

A [`clear_secret_info`](#info-55a3e0) block (within [`more_option.response_cookies_to_add.secret_value`](#value-3191e4)) supports the following:

<a id="ref-ebd966"></a>&#x2022; [`provider_ref`](#ref-ebd966) - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

<a id="url-43bc69"></a>&#x2022; [`url`](#url-43bc69) - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

#### More Option Response Headers To Add

A [`response_headers_to_add`](#more-option-response-headers-to-add) block (within [`more_option`](#more-option)) supports the following:

<a id="append-f099c0"></a>&#x2022; [`append`](#append-f099c0) - Optional Bool  Defaults to `do`<br>Append. Should the value be appended? If true, the value is appended to existing values. not append

<a id="name-9f1ebf"></a>&#x2022; [`name`](#name-9f1ebf) - Optional String<br>Name. Name of the HTTP header

<a id="value-76f49e"></a>&#x2022; [`secret_value`](#value-76f49e) - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Secret Value](#value-76f49e) below.

<a id="value-06dc79"></a>&#x2022; [`value`](#value-06dc79) - Optional String<br>Value. Value of the HTTP header

#### More Option Response Headers To Add Secret Value

A [`secret_value`](#value-76f49e) block (within [`more_option.response_headers_to_add`](#more-option-response-headers-to-add)) supports the following:

<a id="info-6414c7"></a>&#x2022; [`blindfold_secret_info`](#info-6414c7) - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#info-6414c7) below.

<a id="info-1049c1"></a>&#x2022; [`clear_secret_info`](#info-1049c1) - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#info-1049c1) below.

#### More Option Response Headers To Add Secret Value Blindfold Secret Info

A [`blindfold_secret_info`](#info-6414c7) block (within [`more_option.response_headers_to_add.secret_value`](#value-76f49e)) supports the following:

<a id="provider-490cd7"></a>&#x2022; [`decryption_provider`](#provider-490cd7) - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

<a id="location-a47166"></a>&#x2022; [`location`](#location-a47166) - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

<a id="provider-fbf086"></a>&#x2022; [`store_provider`](#provider-fbf086) - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

#### More Option Response Headers To Add Secret Value Clear Secret Info

A [`clear_secret_info`](#info-1049c1) block (within [`more_option.response_headers_to_add.secret_value`](#value-76f49e)) supports the following:

<a id="ref-dadee9"></a>&#x2022; [`provider_ref`](#ref-dadee9) - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

<a id="url-c32192"></a>&#x2022; [`url`](#url-c32192) - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

#### Origin Server Subset Rule List

An [`origin_server_subset_rule_list`](#origin-server-subset-rule-list) block supports the following:

<a id="rules-7a881b"></a>&#x2022; [`origin_server_subset_rules`](#rules-7a881b) - Optional Block<br>Origin Server Subset Rules. Origin Server Subset Rules allow users to define match condition on Client (IP address, ASN, Country), IP Reputation, Regional Edge names, Request for subset selection of origin servers. Origin Server Subset is a sequential engine where rules are evaluated one after the other. It's important to define the correct order for Origin Server Subset to get the intended result, rules are evaluated from top to bottom in the list. When an Origin server subset rule is matched, then this selection rule takes effect and no more rules are evaluated<br>See [Origin Server Subset Rules](#rules-7a881b) below.

#### Origin Server Subset Rule List Origin Server Subset Rules

An [`origin_server_subset_rules`](#rules-7a881b) block (within [`origin_server_subset_rule_list`](#origin-server-subset-rule-list)) supports the following:

<a id="asn-7b608d"></a>&#x2022; [`any_asn`](#asn-7b608d) - Optional Block<br>Enable this option

<a id="any-ip-dce6f3"></a>&#x2022; [`any_ip`](#any-ip-dce6f3) - Optional Block<br>Enable this option

<a id="list-d0328c"></a>&#x2022; [`asn_list`](#list-d0328c) - Optional Block<br>ASN Match List. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer<br>See [Asn List](#list-d0328c) below.

<a id="matcher-87350d"></a>&#x2022; [`asn_matcher`](#matcher-87350d) - Optional Block<br>ASN Matcher. Match any AS number contained in the list of bgp_asn_sets<br>See [Asn Matcher](#matcher-87350d) below.

<a id="selector-98315c"></a>&#x2022; [`client_selector`](#selector-98315c) - Optional Block<br>Label Selector. This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expresssions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects. Label selector is immutable. expressions is a list of strings of label selection expression. Each string has ',' separated values which are 'AND' and all strings are logically 'OR'. BNF for expression string `<selector-syntax>` ::= `<requirement>` | `<requirement>` ',' `<selector-syntax>` `<requirement>` ::= [!] KEY [ `<set-based-restriction>` | `<exact-match-restriction>` ] `<set-based-restriction>` ::= '' | `<inclusion-exclusion>` `<value-set>` `<inclusion-exclusion>` ::= `<inclusion>` | `<exclusion>` `<exclusion>` ::= 'notin' `<inclusion>` ::= 'in' `<value-set>` ::= '(' `<values>` ')' `<values>` ::= VALUE | VALUE ',' `<values>` `<exact-match-restriction>` ::= ['='|'=='|'!='] VALUE<br>See [Client Selector](#selector-98315c) below.

<a id="codes-c20af6"></a>&#x2022; [`country_codes`](#codes-c20af6) - Optional List  Defaults to `COUNTRY_NONE`<br>Possible values are `COUNTRY_NONE`, `COUNTRY_AD`, `COUNTRY_AE`, `COUNTRY_AF`, `COUNTRY_AG`, `COUNTRY_AI`, `COUNTRY_AL`, `COUNTRY_AM`, `COUNTRY_AN`, `COUNTRY_AO`, `COUNTRY_AQ`, `COUNTRY_AR`, `COUNTRY_AS`, `COUNTRY_AT`, `COUNTRY_AU`, `COUNTRY_AW`, `COUNTRY_AX`, `COUNTRY_AZ`, `COUNTRY_BA`, `COUNTRY_BB`, `COUNTRY_BD`, `COUNTRY_BE`, `COUNTRY_BF`, `COUNTRY_BG`, `COUNTRY_BH`, `COUNTRY_BI`, `COUNTRY_BJ`, `COUNTRY_BL`, `COUNTRY_BM`, `COUNTRY_BN`, `COUNTRY_BO`, `COUNTRY_BQ`, `COUNTRY_BR`, `COUNTRY_BS`, `COUNTRY_BT`, `COUNTRY_BV`, `COUNTRY_BW`, `COUNTRY_BY`, `COUNTRY_BZ`, `COUNTRY_CA`, `COUNTRY_CC`, `COUNTRY_CD`, `COUNTRY_CF`, `COUNTRY_CG`, `COUNTRY_CH`, `COUNTRY_CI`, `COUNTRY_CK`, `COUNTRY_CL`, `COUNTRY_CM`, `COUNTRY_CN`, `COUNTRY_CO`, `COUNTRY_CR`, `COUNTRY_CS`, `COUNTRY_CU`, `COUNTRY_CV`, `COUNTRY_CW`, `COUNTRY_CX`, `COUNTRY_CY`, `COUNTRY_CZ`, `COUNTRY_DE`, `COUNTRY_DJ`, `COUNTRY_DK`, `COUNTRY_DM`, `COUNTRY_DO`, `COUNTRY_DZ`, `COUNTRY_EC`, `COUNTRY_EE`, `COUNTRY_EG`, `COUNTRY_EH`, `COUNTRY_ER`, `COUNTRY_ES`, `COUNTRY_ET`, `COUNTRY_FI`, `COUNTRY_FJ`, `COUNTRY_FK`, `COUNTRY_FM`, `COUNTRY_FO`, `COUNTRY_FR`, `COUNTRY_GA`, `COUNTRY_GB`, `COUNTRY_GD`, `COUNTRY_GE`, `COUNTRY_GF`, `COUNTRY_GG`, `COUNTRY_GH`, `COUNTRY_GI`, `COUNTRY_GL`, `COUNTRY_GM`, `COUNTRY_GN`, `COUNTRY_GP`, `COUNTRY_GQ`, `COUNTRY_GR`, `COUNTRY_GS`, `COUNTRY_GT`, `COUNTRY_GU`, `COUNTRY_GW`, `COUNTRY_GY`, `COUNTRY_HK`, `COUNTRY_HM`, `COUNTRY_HN`, `COUNTRY_HR`, `COUNTRY_HT`, `COUNTRY_HU`, `COUNTRY_ID`, `COUNTRY_IE`, `COUNTRY_IL`, `COUNTRY_IM`, `COUNTRY_IN`, `COUNTRY_IO`, `COUNTRY_IQ`, `COUNTRY_IR`, `COUNTRY_IS`, `COUNTRY_IT`, `COUNTRY_JE`, `COUNTRY_JM`, `COUNTRY_JO`, `COUNTRY_JP`, `COUNTRY_KE`, `COUNTRY_KG`, `COUNTRY_KH`, `COUNTRY_KI`, `COUNTRY_KM`, `COUNTRY_KN`, `COUNTRY_KP`, `COUNTRY_KR`, `COUNTRY_KW`, `COUNTRY_KY`, `COUNTRY_KZ`, `COUNTRY_LA`, `COUNTRY_LB`, `COUNTRY_LC`, `COUNTRY_LI`, `COUNTRY_LK`, `COUNTRY_LR`, `COUNTRY_LS`, `COUNTRY_LT`, `COUNTRY_LU`, `COUNTRY_LV`, `COUNTRY_LY`, `COUNTRY_MA`, `COUNTRY_MC`, `COUNTRY_MD`, `COUNTRY_ME`, `COUNTRY_MF`, `COUNTRY_MG`, `COUNTRY_MH`, `COUNTRY_MK`, `COUNTRY_ML`, `COUNTRY_MM`, `COUNTRY_MN`, `COUNTRY_MO`, `COUNTRY_MP`, `COUNTRY_MQ`, `COUNTRY_MR`, `COUNTRY_MS`, `COUNTRY_MT`, `COUNTRY_MU`, `COUNTRY_MV`, `COUNTRY_MW`, `COUNTRY_MX`, `COUNTRY_MY`, `COUNTRY_MZ`, `COUNTRY_NA`, `COUNTRY_NC`, `COUNTRY_NE`, `COUNTRY_NF`, `COUNTRY_NG`, `COUNTRY_NI`, `COUNTRY_NL`, `COUNTRY_NO`, `COUNTRY_NP`, `COUNTRY_NR`, `COUNTRY_NU`, `COUNTRY_NZ`, `COUNTRY_OM`, `COUNTRY_PA`, `COUNTRY_PE`, `COUNTRY_PF`, `COUNTRY_PG`, `COUNTRY_PH`, `COUNTRY_PK`, `COUNTRY_PL`, `COUNTRY_PM`, `COUNTRY_PN`, `COUNTRY_PR`, `COUNTRY_PS`, `COUNTRY_PT`, `COUNTRY_PW`, `COUNTRY_PY`, `COUNTRY_QA`, `COUNTRY_RE`, `COUNTRY_RO`, `COUNTRY_RS`, `COUNTRY_RU`, `COUNTRY_RW`, `COUNTRY_SA`, `COUNTRY_SB`, `COUNTRY_SC`, `COUNTRY_SD`, `COUNTRY_SE`, `COUNTRY_SG`, `COUNTRY_SH`, `COUNTRY_SI`, `COUNTRY_SJ`, `COUNTRY_SK`, `COUNTRY_SL`, `COUNTRY_SM`, `COUNTRY_SN`, `COUNTRY_SO`, `COUNTRY_SR`, `COUNTRY_SS`, `COUNTRY_ST`, `COUNTRY_SV`, `COUNTRY_SX`, `COUNTRY_SY`, `COUNTRY_SZ`, `COUNTRY_TC`, `COUNTRY_TD`, `COUNTRY_TF`, `COUNTRY_TG`, `COUNTRY_TH`, `COUNTRY_TJ`, `COUNTRY_TK`, `COUNTRY_TL`, `COUNTRY_TM`, `COUNTRY_TN`, `COUNTRY_TO`, `COUNTRY_TR`, `COUNTRY_TT`, `COUNTRY_TV`, `COUNTRY_TW`, `COUNTRY_TZ`, `COUNTRY_UA`, `COUNTRY_UG`, `COUNTRY_UM`, `COUNTRY_US`, `COUNTRY_UY`, `COUNTRY_UZ`, `COUNTRY_VA`, `COUNTRY_VC`, `COUNTRY_VE`, `COUNTRY_VG`, `COUNTRY_VI`, `COUNTRY_VN`, `COUNTRY_VU`, `COUNTRY_WF`, `COUNTRY_WS`, `COUNTRY_XK`, `COUNTRY_XT`, `COUNTRY_YE`, `COUNTRY_YT`, `COUNTRY_ZA`, `COUNTRY_ZM`, `COUNTRY_ZW`<br>[Enum: COUNTRY_NONE|COUNTRY_AD|COUNTRY_AE|COUNTRY_AF|COUNTRY_AG|COUNTRY_AI|COUNTRY_AL|COUNTRY_AM|COUNTRY_AN|COUNTRY_AO|COUNTRY_AQ|COUNTRY_AR|COUNTRY_AS|COUNTRY_AT|COUNTRY_AU|COUNTRY_AW|COUNTRY_AX|COUNTRY_AZ|COUNTRY_BA|COUNTRY_BB|COUNTRY_BD|COUNTRY_BE|COUNTRY_BF|COUNTRY_BG|COUNTRY_BH|COUNTRY_BI|COUNTRY_BJ|COUNTRY_BL|COUNTRY_BM|COUNTRY_BN|COUNTRY_BO|COUNTRY_BQ|COUNTRY_BR|COUNTRY_BS|COUNTRY_BT|COUNTRY_BV|COUNTRY_BW|COUNTRY_BY|COUNTRY_BZ|COUNTRY_CA|COUNTRY_CC|COUNTRY_CD|COUNTRY_CF|COUNTRY_CG|COUNTRY_CH|COUNTRY_CI|COUNTRY_CK|COUNTRY_CL|COUNTRY_CM|COUNTRY_CN|COUNTRY_CO|COUNTRY_CR|COUNTRY_CS|COUNTRY_CU|COUNTRY_CV|COUNTRY_CW|COUNTRY_CX|COUNTRY_CY|COUNTRY_CZ|COUNTRY_DE|COUNTRY_DJ|COUNTRY_DK|COUNTRY_DM|COUNTRY_DO|COUNTRY_DZ|COUNTRY_EC|COUNTRY_EE|COUNTRY_EG|COUNTRY_EH|COUNTRY_ER|COUNTRY_ES|COUNTRY_ET|COUNTRY_FI|COUNTRY_FJ|COUNTRY_FK|COUNTRY_FM|COUNTRY_FO|COUNTRY_FR|COUNTRY_GA|COUNTRY_GB|COUNTRY_GD|COUNTRY_GE|COUNTRY_GF|COUNTRY_GG|COUNTRY_GH|COUNTRY_GI|COUNTRY_GL|COUNTRY_GM|COUNTRY_GN|COUNTRY_GP|COUNTRY_GQ|COUNTRY_GR|COUNTRY_GS|COUNTRY_GT|COUNTRY_GU|COUNTRY_GW|COUNTRY_GY|COUNTRY_HK|COUNTRY_HM|COUNTRY_HN|COUNTRY_HR|COUNTRY_HT|COUNTRY_HU|COUNTRY_ID|COUNTRY_IE|COUNTRY_IL|COUNTRY_IM|COUNTRY_IN|COUNTRY_IO|COUNTRY_IQ|COUNTRY_IR|COUNTRY_IS|COUNTRY_IT|COUNTRY_JE|COUNTRY_JM|COUNTRY_JO|COUNTRY_JP|COUNTRY_KE|COUNTRY_KG|COUNTRY_KH|COUNTRY_KI|COUNTRY_KM|COUNTRY_KN|COUNTRY_KP|COUNTRY_KR|COUNTRY_KW|COUNTRY_KY|COUNTRY_KZ|COUNTRY_LA|COUNTRY_LB|COUNTRY_LC|COUNTRY_LI|COUNTRY_LK|COUNTRY_LR|COUNTRY_LS|COUNTRY_LT|COUNTRY_LU|COUNTRY_LV|COUNTRY_LY|COUNTRY_MA|COUNTRY_MC|COUNTRY_MD|COUNTRY_ME|COUNTRY_MF|COUNTRY_MG|COUNTRY_MH|COUNTRY_MK|COUNTRY_ML|COUNTRY_MM|COUNTRY_MN|COUNTRY_MO|COUNTRY_MP|COUNTRY_MQ|COUNTRY_MR|COUNTRY_MS|COUNTRY_MT|COUNTRY_MU|COUNTRY_MV|COUNTRY_MW|COUNTRY_MX|COUNTRY_MY|COUNTRY_MZ|COUNTRY_NA|COUNTRY_NC|COUNTRY_NE|COUNTRY_NF|COUNTRY_NG|COUNTRY_NI|COUNTRY_NL|COUNTRY_NO|COUNTRY_NP|COUNTRY_NR|COUNTRY_NU|COUNTRY_NZ|COUNTRY_OM|COUNTRY_PA|COUNTRY_PE|COUNTRY_PF|COUNTRY_PG|COUNTRY_PH|COUNTRY_PK|COUNTRY_PL|COUNTRY_PM|COUNTRY_PN|COUNTRY_PR|COUNTRY_PS|COUNTRY_PT|COUNTRY_PW|COUNTRY_PY|COUNTRY_QA|COUNTRY_RE|COUNTRY_RO|COUNTRY_RS|COUNTRY_RU|COUNTRY_RW|COUNTRY_SA|COUNTRY_SB|COUNTRY_SC|COUNTRY_SD|COUNTRY_SE|COUNTRY_SG|COUNTRY_SH|COUNTRY_SI|COUNTRY_SJ|COUNTRY_SK|COUNTRY_SL|COUNTRY_SM|COUNTRY_SN|COUNTRY_SO|COUNTRY_SR|COUNTRY_SS|COUNTRY_ST|COUNTRY_SV|COUNTRY_SX|COUNTRY_SY|COUNTRY_SZ|COUNTRY_TC|COUNTRY_TD|COUNTRY_TF|COUNTRY_TG|COUNTRY_TH|COUNTRY_TJ|COUNTRY_TK|COUNTRY_TL|COUNTRY_TM|COUNTRY_TN|COUNTRY_TO|COUNTRY_TR|COUNTRY_TT|COUNTRY_TV|COUNTRY_TW|COUNTRY_TZ|COUNTRY_UA|COUNTRY_UG|COUNTRY_UM|COUNTRY_US|COUNTRY_UY|COUNTRY_UZ|COUNTRY_VA|COUNTRY_VC|COUNTRY_VE|COUNTRY_VG|COUNTRY_VI|COUNTRY_VN|COUNTRY_VU|COUNTRY_WF|COUNTRY_WS|COUNTRY_XK|COUNTRY_XT|COUNTRY_YE|COUNTRY_YT|COUNTRY_ZA|COUNTRY_ZM|COUNTRY_ZW] Country Codes List. List of Country Codes

<a id="matcher-85c46f"></a>&#x2022; [`ip_matcher`](#matcher-85c46f) - Optional Block<br>IP Prefix Matcher. Match any IP prefix contained in the list of ip_prefix_sets. The result of the match is inverted if invert_matcher is true<br>See [IP Matcher](#matcher-85c46f) below.

<a id="list-227268"></a>&#x2022; [`ip_prefix_list`](#list-227268) - Optional Block<br>IP Prefix Match List. List of IP Prefix strings to match against<br>See [IP Prefix List](#list-227268) below.

<a id="metadata-341112"></a>&#x2022; [`metadata`](#metadata-341112) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#metadata-341112) below.

<a id="none-b1132f"></a>&#x2022; [`none`](#none-b1132f) - Optional Block<br>Enable this option

<a id="action-194a52"></a>&#x2022; [`origin_server_subsets_action`](#action-194a52) - Optional Block<br>Action. Add labels to select one or more origin servers. Note: The pre-requisite settings to be configured in the origin pool are: 1. Add labels to origin servers 2. Enable subset load balancing in the Origin Server Subsets section and configure keys in origin server subsets classes

<a id="list-9ae7d2"></a>&#x2022; [`re_name_list`](#list-9ae7d2) - Optional List<br>RE Names. List of RE names for match

#### Origin Server Subset Rule List Origin Server Subset Rules Asn List

An [`asn_list`](#list-d0328c) block (within [`origin_server_subset_rule_list.origin_server_subset_rules`](#rules-7a881b)) supports the following:

<a id="numbers-77c506"></a>&#x2022; [`as_numbers`](#numbers-77c506) - Optional List<br>AS Numbers. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer

#### Origin Server Subset Rule List Origin Server Subset Rules Asn Matcher

An [`asn_matcher`](#matcher-87350d) block (within [`origin_server_subset_rule_list.origin_server_subset_rules`](#rules-7a881b)) supports the following:

<a id="sets-d0ded4"></a>&#x2022; [`asn_sets`](#sets-d0ded4) - Optional Block<br>BGP ASN Sets. A list of references to bgp_asn_set objects<br>See [Asn Sets](#sets-d0ded4) below.

#### Origin Server Subset Rule List Origin Server Subset Rules Asn Matcher Asn Sets

An [`asn_sets`](#sets-d0ded4) block (within [`origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher`](#matcher-87350d)) supports the following:

<a id="kind-d45a2f"></a>&#x2022; [`kind`](#kind-d45a2f) - Optional String<br>Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route')

<a id="name-51f0ce"></a>&#x2022; [`name`](#name-51f0ce) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-8189ee"></a>&#x2022; [`namespace`](#namespace-8189ee) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-5f4eba"></a>&#x2022; [`tenant`](#tenant-5f4eba) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="uid-fe6057"></a>&#x2022; [`uid`](#uid-fe6057) - Optional String<br>UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid

#### Origin Server Subset Rule List Origin Server Subset Rules Client Selector

A [`client_selector`](#selector-98315c) block (within [`origin_server_subset_rule_list.origin_server_subset_rules`](#rules-7a881b)) supports the following:

<a id="expressions-c0ca8c"></a>&#x2022; [`expressions`](#expressions-c0ca8c) - Optional List<br>Selector Expression. expressions contains the kubernetes style label expression for selections

#### Origin Server Subset Rule List Origin Server Subset Rules IP Matcher

An [`ip_matcher`](#matcher-85c46f) block (within [`origin_server_subset_rule_list.origin_server_subset_rules`](#rules-7a881b)) supports the following:

<a id="matcher-f6e828"></a>&#x2022; [`invert_matcher`](#matcher-f6e828) - Optional Bool<br>Invert IP Matcher. Invert the match result

<a id="sets-bbd7a8"></a>&#x2022; [`prefix_sets`](#sets-bbd7a8) - Optional Block<br>IP Prefix Sets. A list of references to ip_prefix_set objects<br>See [Prefix Sets](#sets-bbd7a8) below.

#### Origin Server Subset Rule List Origin Server Subset Rules IP Matcher Prefix Sets

A [`prefix_sets`](#sets-bbd7a8) block (within [`origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher`](#matcher-85c46f)) supports the following:

<a id="kind-dee008"></a>&#x2022; [`kind`](#kind-dee008) - Optional String<br>Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route')

<a id="name-8fb4c6"></a>&#x2022; [`name`](#name-8fb4c6) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-4402ae"></a>&#x2022; [`namespace`](#namespace-4402ae) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-2521fb"></a>&#x2022; [`tenant`](#tenant-2521fb) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="uid-b4d427"></a>&#x2022; [`uid`](#uid-b4d427) - Optional String<br>UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid

#### Origin Server Subset Rule List Origin Server Subset Rules IP Prefix List

An [`ip_prefix_list`](#list-227268) block (within [`origin_server_subset_rule_list.origin_server_subset_rules`](#rules-7a881b)) supports the following:

<a id="match-5ace95"></a>&#x2022; [`invert_match`](#match-5ace95) - Optional Bool<br>Invert Match Result. Invert the match result

<a id="prefixes-32abf0"></a>&#x2022; [`ip_prefixes`](#prefixes-32abf0) - Optional List<br>IPv4 Prefix List. List of IPv4 prefix strings

#### Origin Server Subset Rule List Origin Server Subset Rules Metadata

A [`metadata`](#metadata-341112) block (within [`origin_server_subset_rule_list.origin_server_subset_rules`](#rules-7a881b)) supports the following:

<a id="spec-b59444"></a>&#x2022; [`description_spec`](#spec-b59444) - Optional String<br>Description. Human readable description

<a id="name-4eb15a"></a>&#x2022; [`name`](#name-4eb15a) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### Policy Based Challenge

A [`policy_based_challenge`](#policy-based-challenge) block supports the following:

<a id="challenge-a53c7e"></a>&#x2022; [`always_enable_captcha_challenge`](#challenge-a53c7e) - Optional Block<br>Enable this option

<a id="challenge-3ba035"></a>&#x2022; [`always_enable_js_challenge`](#challenge-3ba035) - Optional Block<br>Enable this option

<a id="parameters-699e87"></a>&#x2022; [`captcha_challenge_parameters`](#parameters-699e87) - Optional Block<br>Captcha Challenge Parameters. Enables loadbalancer to perform captcha challenge Captcha challenge will be based on Google Recaptcha. With this feature enabled, only clients that pass the captcha challenge will be allowed to complete the HTTP request. When loadbalancer is configured to do Captcha Challenge, it will redirect the browser to an HTML page on every new HTTP request. This HTML page will have captcha challenge embedded in it. Client will be allowed to make the request only if the captcha challenge is successful. Loadbalancer will tag response header with a cookie to avoid Captcha challenge for subsequent requests. CAPTCHA is mainly used as a security check to ensure only human users can pass through. Generally, computers or bots are not capable of solving a captcha. You can enable either Javascript challenge or Captcha challenge on a virtual host<br>See [Captcha Challenge Parameters](#parameters-699e87) below.

<a id="parameters-1afe14"></a>&#x2022; [`default_captcha_challenge_parameters`](#parameters-1afe14) - Optional Block<br>Enable this option

<a id="parameters-d11492"></a>&#x2022; [`default_js_challenge_parameters`](#parameters-d11492) - Optional Block<br>Enable this option

<a id="settings-3c8e74"></a>&#x2022; [`default_mitigation_settings`](#settings-3c8e74) - Optional Block<br>Enable this option

<a id="parameters-f17f1a"></a>&#x2022; [`default_temporary_blocking_parameters`](#parameters-f17f1a) - Optional Block<br>Enable this option

<a id="parameters-65055e"></a>&#x2022; [`js_challenge_parameters`](#parameters-65055e) - Optional Block<br>Javascript Challenge Parameters. Enables loadbalancer to perform client browser compatibility test by redirecting to a page with Javascript. With this feature enabled, only clients that are capable of executing Javascript(mostly browsers) will be allowed to complete the HTTP request. When loadbalancer is configured to do Javascript Challenge, it will redirect the browser to an HTML page on every new HTTP request. This HTML page will have Javascript embedded in it. Loadbalancer chooses a set of random numbers for every new client and sends these numbers along with an encrypted answer with the request such that it embed these numbers as input in the Javascript. Javascript will run on the requestor browser and perform a complex Math operation. Script will submit the answer to loadbalancer. Loadbalancer will validate the answer by comparing the calculated answer with the decrypted answer (which was encrypted when it was sent back as reply) and allow the request to the upstream server only if the answer is correct. Loadbalancer will tag response header with a cookie to avoid Javascript challenge for subsequent requests. Javascript challenge serves following purposes \* Validate that the request is coming via a browser that is capable for running Javascript \* Force the browser to run a complex operation, f(X), that requires it to spend a large number of CPU cycles. This is to slow down a potential DOS attacker by making it difficult to launch a large request flood without having to spend even larger CPU cost at their end. You can enable either Javascript challenge or Captcha challenge on a virtual host<br>See [Js Challenge Parameters](#parameters-65055e) below.

<a id="mitigation-d19aea"></a>&#x2022; [`malicious_user_mitigation`](#mitigation-d19aea) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Malicious User Mitigation](#mitigation-d19aea) below.

<a id="policy-based-challenge-no-challenge"></a>&#x2022; [`no_challenge`](#policy-based-challenge-no-challenge) - Optional Block<br>Enable this option

<a id="policy-based-challenge-rule-list"></a>&#x2022; [`rule_list`](#policy-based-challenge-rule-list) - Optional Block<br>Challenge Rule List. List of challenge rules to be used in policy based challenge<br>See [Rule List](#policy-based-challenge-rule-list) below.

<a id="blocking-9fdca7"></a>&#x2022; [`temporary_user_blocking`](#blocking-9fdca7) - Optional Block<br>Temporary User Blocking. Specifies configuration for temporary user blocking resulting from user behavior analysis. When Malicious User Mitigation is enabled from service policy rules, users' accessing the application will be analyzed for malicious activity and the configured mitigation actions will be taken on identified malicious users. These mitigation actions include setting up temporary blocking on that user. This configuration specifies settings on how that blocking should be done by the loadbalancer<br>See [Temporary User Blocking](#blocking-9fdca7) below.

#### Policy Based Challenge Captcha Challenge Parameters

A [`captcha_challenge_parameters`](#parameters-699e87) block (within [`policy_based_challenge`](#policy-based-challenge)) supports the following:

<a id="expiry-fff199"></a>&#x2022; [`cookie_expiry`](#expiry-fff199) - Optional Number<br>Cookie Expiration Period. Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new challenge

<a id="page-46537d"></a>&#x2022; [`custom_page`](#page-46537d) - Optional String<br>Custom message for Captcha Challenge. Custom message is of type uri_ref. Currently supported URL schemes is string:///. For string:/// scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. 'Please Wait.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '`<p>` Please Wait `</p>`'. Base64 encoded string for this HTML is 'PHA+IFBsZWFzZSBXYWl0IDwvcD4='

#### Policy Based Challenge Js Challenge Parameters

A [`js_challenge_parameters`](#parameters-65055e) block (within [`policy_based_challenge`](#policy-based-challenge)) supports the following:

<a id="expiry-745058"></a>&#x2022; [`cookie_expiry`](#expiry-745058) - Optional Number<br>Cookie Expiration Period. Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new challenge

<a id="page-59809e"></a>&#x2022; [`custom_page`](#page-59809e) - Optional String<br>Custom Message for Javascript Challenge. Custom message is of type uri_ref. Currently supported URL schemes is string:///. For string:/// scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. 'Please Wait.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '`<p>` Please Wait `</p>`'. Base64 encoded string for this HTML is 'PHA+IFBsZWFzZSBXYWl0IDwvcD4='

<a id="delay-6a6ceb"></a>&#x2022; [`js_script_delay`](#delay-6a6ceb) - Optional Number<br>Javascript Delay. Delay introduced by Javascript, in milliseconds

#### Policy Based Challenge Malicious User Mitigation

A [`malicious_user_mitigation`](#mitigation-d19aea) block (within [`policy_based_challenge`](#policy-based-challenge)) supports the following:

<a id="name-0fb02d"></a>&#x2022; [`name`](#name-0fb02d) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-490d76"></a>&#x2022; [`namespace`](#namespace-490d76) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-cf334a"></a>&#x2022; [`tenant`](#tenant-cf334a) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Policy Based Challenge Rule List

A [`rule_list`](#policy-based-challenge-rule-list) block (within [`policy_based_challenge`](#policy-based-challenge)) supports the following:

<a id="policy-based-challenge-rule-list-rules"></a>&#x2022; [`rules`](#policy-based-challenge-rule-list-rules) - Optional Block<br>Rules. Rules that specify the match conditions and challenge type to be launched. When a challenge type is selected to be always enabled, these rules can be used to disable challenge or launch a different challenge for requests that match the specified conditions<br>See [Rules](#policy-based-challenge-rule-list-rules) below.

#### Policy Based Challenge Rule List Rules

A [`rules`](#policy-based-challenge-rule-list-rules) block (within [`policy_based_challenge.rule_list`](#policy-based-challenge-rule-list)) supports the following:

<a id="metadata-72ce94"></a>&#x2022; [`metadata`](#metadata-72ce94) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#metadata-72ce94) below.

<a id="spec-fbd0f9"></a>&#x2022; [`spec`](#spec-fbd0f9) - Optional Block<br>Challenge Rule Specification. A Challenge Rule consists of an unordered list of predicates and an action. The predicates are evaluated against a set of input fields that are extracted from or derived from an L7 request API. A request API is considered to match the rule if all predicates in the rule evaluate to true for that request. Any predicates that are not specified in a rule are implicitly considered to be true. If a request API matches a challenge rule, the configured challenge is enforced<br>See [Spec](#spec-fbd0f9) below.

#### Policy Based Challenge Rule List Rules Metadata

A [`metadata`](#metadata-72ce94) block (within [`policy_based_challenge.rule_list.rules`](#policy-based-challenge-rule-list-rules)) supports the following:

<a id="spec-760f37"></a>&#x2022; [`description_spec`](#spec-760f37) - Optional String<br>Description. Human readable description

<a id="name-44607a"></a>&#x2022; [`name`](#name-44607a) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### Policy Based Challenge Rule List Rules Spec

A [`spec`](#spec-fbd0f9) block (within [`policy_based_challenge.rule_list.rules`](#policy-based-challenge-rule-list-rules)) supports the following:

<a id="asn-cae05d"></a>&#x2022; [`any_asn`](#asn-cae05d) - Optional Block<br>Enable this option

<a id="client-df7cdb"></a>&#x2022; [`any_client`](#client-df7cdb) - Optional Block<br>Enable this option

<a id="any-ip-6a2554"></a>&#x2022; [`any_ip`](#any-ip-6a2554) - Optional Block<br>Enable this option

<a id="matchers-86dff2"></a>&#x2022; [`arg_matchers`](#matchers-86dff2) - Optional Block<br>A list of predicates for all POST args that need to be matched. The criteria for matching each arg are described in individual instances of ArgMatcherType. The actual arg values are extracted from the request API as a list of strings for each arg selector name. Note that all specified arg matcher predicates must evaluate to true<br>See [Arg Matchers](#matchers-86dff2) below.

<a id="list-628bd4"></a>&#x2022; [`asn_list`](#list-628bd4) - Optional Block<br>ASN Match List. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer<br>See [Asn List](#list-628bd4) below.

<a id="matcher-6b840c"></a>&#x2022; [`asn_matcher`](#matcher-6b840c) - Optional Block<br>ASN Matcher. Match any AS number contained in the list of bgp_asn_sets<br>See [Asn Matcher](#matcher-6b840c) below.

<a id="matcher-4075fc"></a>&#x2022; [`body_matcher`](#matcher-4075fc) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Body Matcher](#matcher-4075fc) below.

<a id="selector-ca44f5"></a>&#x2022; [`client_selector`](#selector-ca44f5) - Optional Block<br>Label Selector. This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expresssions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects. Label selector is immutable. expressions is a list of strings of label selection expression. Each string has ',' separated values which are 'AND' and all strings are logically 'OR'. BNF for expression string `<selector-syntax>` ::= `<requirement>` | `<requirement>` ',' `<selector-syntax>` `<requirement>` ::= [!] KEY [ `<set-based-restriction>` | `<exact-match-restriction>` ] `<set-based-restriction>` ::= '' | `<inclusion-exclusion>` `<value-set>` `<inclusion-exclusion>` ::= `<inclusion>` | `<exclusion>` `<exclusion>` ::= 'notin' `<inclusion>` ::= 'in' `<value-set>` ::= '(' `<values>` ')' `<values>` ::= VALUE | VALUE ',' `<values>` `<exact-match-restriction>` ::= ['='|'=='|'!='] VALUE<br>See [Client Selector](#selector-ca44f5) below.

<a id="matchers-cb349b"></a>&#x2022; [`cookie_matchers`](#matchers-cb349b) - Optional Block<br>A list of predicates for all cookies that need to be matched. The criteria for matching each cookie is described in individual instances of CookieMatcherType. The actual cookie values are extracted from the request API as a list of strings for each cookie name. Note that all specified cookie matcher predicates must evaluate to true<br>See [Cookie Matchers](#matchers-cb349b) below.

<a id="challenge-fbd9a1"></a>&#x2022; [`disable_challenge`](#challenge-fbd9a1) - Optional Block<br>Enable this option

<a id="matcher-888f5a"></a>&#x2022; [`domain_matcher`](#matcher-888f5a) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Domain Matcher](#matcher-888f5a) below.

<a id="challenge-e0353b"></a>&#x2022; [`enable_captcha_challenge`](#challenge-e0353b) - Optional Block<br>Enable this option

<a id="challenge-3644c3"></a>&#x2022; [`enable_javascript_challenge`](#challenge-3644c3) - Optional Block<br>Enable this option

<a id="timestamp-6d26e1"></a>&#x2022; [`expiration_timestamp`](#timestamp-6d26e1) - Optional String<br>The expiration_timestamp is the RFC 3339 format timestamp at which the containing rule is considered to be logically expired. The rule continues to exist in the configuration but is not applied anymore

<a id="headers-1bea3b"></a>&#x2022; [`headers`](#headers-1bea3b) - Optional Block<br>A list of predicates for various HTTP headers that need to match. The criteria for matching each HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values are extracted from the request API as a list of strings for each HTTP header type. Note that all specified header predicates must evaluate to true<br>See [Headers](#headers-1bea3b) below.

<a id="method-9ab722"></a>&#x2022; [`http_method`](#method-9ab722) - Optional Block<br>HTTP Method Matcher. A HTTP method matcher specifies a list of methods to match an input HTTP method. The match is considered successful if the input method is a member of the list. The result of the match based on the method list is inverted if invert_matcher is true<br>See [HTTP Method](#method-9ab722) below.

<a id="matcher-82616b"></a>&#x2022; [`ip_matcher`](#matcher-82616b) - Optional Block<br>IP Prefix Matcher. Match any IP prefix contained in the list of ip_prefix_sets. The result of the match is inverted if invert_matcher is true<br>See [IP Matcher](#matcher-82616b) below.

<a id="list-537143"></a>&#x2022; [`ip_prefix_list`](#list-537143) - Optional Block<br>IP Prefix Match List. List of IP Prefix strings to match against<br>See [IP Prefix List](#list-537143) below.

<a id="path-acb3cc"></a>&#x2022; [`path`](#path-acb3cc) - Optional Block<br>Path Matcher. A path matcher specifies multiple criteria for matching an HTTP path string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of path prefixes, a list of exact path values and a list of regular expressions<br>See [Path](#path-acb3cc) below.

<a id="params-04b1ad"></a>&#x2022; [`query_params`](#params-04b1ad) - Optional Block<br>A list of predicates for all query parameters that need to be matched. The criteria for matching each query parameter are described in individual instances of QueryParameterMatcherType. The actual query parameter values are extracted from the request API as a list of strings for each query parameter name. Note that all specified query parameter predicates must evaluate to true<br>See [Query Params](#params-04b1ad) below.

<a id="matcher-3cbc4b"></a>&#x2022; [`tls_fingerprint_matcher`](#matcher-3cbc4b) - Optional Block<br>TLS Fingerprint Matcher. A TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of supported positve match criteria includes a list of known classes of TLS fingerprints and a list of exact values. The match is considered successful if either of these positive criteria are satisfied and the input fingerprint is not one of the excluded values<br>See [TLS Fingerprint Matcher](#matcher-3cbc4b) below.

#### Policy Based Challenge Rule List Rules Spec Arg Matchers

An [`arg_matchers`](#matchers-86dff2) block (within [`policy_based_challenge.rule_list.rules.spec`](#spec-fbd0f9)) supports the following:

<a id="present-9fee6d"></a>&#x2022; [`check_not_present`](#present-9fee6d) - Optional Block<br>Enable this option

<a id="present-07ddc9"></a>&#x2022; [`check_present`](#present-07ddc9) - Optional Block<br>Enable this option

<a id="matcher-e00de3"></a>&#x2022; [`invert_matcher`](#matcher-e00de3) - Optional Bool<br>Invert Matcher. Invert Match of the expression defined

<a id="item-ab8776"></a>&#x2022; [`item`](#item-ab8776) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#item-ab8776) below.

<a id="name-965072"></a>&#x2022; [`name`](#name-965072) - Optional String<br>Argument Name. A case-sensitive JSON path in the HTTP request body

#### Policy Based Challenge Rule List Rules Spec Arg Matchers Item

An [`item`](#item-ab8776) block (within [`policy_based_challenge.rule_list.rules.spec.arg_matchers`](#matchers-86dff2)) supports the following:

<a id="values-b5c3b2"></a>&#x2022; [`exact_values`](#values-b5c3b2) - Optional List<br>Exact Values. A list of exact values to match the input against

<a id="values-8f235d"></a>&#x2022; [`regex_values`](#values-8f235d) - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="transformers-707b5a"></a>&#x2022; [`transformers`](#transformers-707b5a) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>[Enum: LOWER_CASE|UPPER_CASE|BASE64_DECODE|NORMALIZE_PATH|REMOVE_WHITESPACE|URL_DECODE|TRIM_LEFT|TRIM_RIGHT|TRIM] Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### Policy Based Challenge Rule List Rules Spec Asn List

An [`asn_list`](#list-628bd4) block (within [`policy_based_challenge.rule_list.rules.spec`](#spec-fbd0f9)) supports the following:

<a id="numbers-deb154"></a>&#x2022; [`as_numbers`](#numbers-deb154) - Optional List<br>AS Numbers. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer

#### Policy Based Challenge Rule List Rules Spec Asn Matcher

An [`asn_matcher`](#matcher-6b840c) block (within [`policy_based_challenge.rule_list.rules.spec`](#spec-fbd0f9)) supports the following:

<a id="sets-be0fa6"></a>&#x2022; [`asn_sets`](#sets-be0fa6) - Optional Block<br>BGP ASN Sets. A list of references to bgp_asn_set objects<br>See [Asn Sets](#sets-be0fa6) below.

#### Policy Based Challenge Rule List Rules Spec Asn Matcher Asn Sets

An [`asn_sets`](#sets-be0fa6) block (within [`policy_based_challenge.rule_list.rules.spec.asn_matcher`](#matcher-6b840c)) supports the following:

<a id="kind-fe53d2"></a>&#x2022; [`kind`](#kind-fe53d2) - Optional String<br>Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route')

<a id="name-064ac5"></a>&#x2022; [`name`](#name-064ac5) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-b9224d"></a>&#x2022; [`namespace`](#namespace-b9224d) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-d8d9f8"></a>&#x2022; [`tenant`](#tenant-d8d9f8) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="uid-67bb42"></a>&#x2022; [`uid`](#uid-67bb42) - Optional String<br>UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid

#### Policy Based Challenge Rule List Rules Spec Body Matcher

A [`body_matcher`](#matcher-4075fc) block (within [`policy_based_challenge.rule_list.rules.spec`](#spec-fbd0f9)) supports the following:

<a id="values-597ee8"></a>&#x2022; [`exact_values`](#values-597ee8) - Optional List<br>Exact Values. A list of exact values to match the input against

<a id="values-4179b7"></a>&#x2022; [`regex_values`](#values-4179b7) - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="transformers-71023a"></a>&#x2022; [`transformers`](#transformers-71023a) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>[Enum: LOWER_CASE|UPPER_CASE|BASE64_DECODE|NORMALIZE_PATH|REMOVE_WHITESPACE|URL_DECODE|TRIM_LEFT|TRIM_RIGHT|TRIM] Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### Policy Based Challenge Rule List Rules Spec Client Selector

A [`client_selector`](#selector-ca44f5) block (within [`policy_based_challenge.rule_list.rules.spec`](#spec-fbd0f9)) supports the following:

<a id="expressions-4e019d"></a>&#x2022; [`expressions`](#expressions-4e019d) - Optional List<br>Selector Expression. expressions contains the kubernetes style label expression for selections

#### Policy Based Challenge Rule List Rules Spec Cookie Matchers

A [`cookie_matchers`](#matchers-cb349b) block (within [`policy_based_challenge.rule_list.rules.spec`](#spec-fbd0f9)) supports the following:

<a id="present-00c6c9"></a>&#x2022; [`check_not_present`](#present-00c6c9) - Optional Block<br>Enable this option

<a id="present-df779f"></a>&#x2022; [`check_present`](#present-df779f) - Optional Block<br>Enable this option

<a id="matcher-aec0f9"></a>&#x2022; [`invert_matcher`](#matcher-aec0f9) - Optional Bool<br>Invert Matcher. Invert Match of the expression defined

<a id="item-951d5f"></a>&#x2022; [`item`](#item-951d5f) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#item-951d5f) below.

<a id="name-9bddbc"></a>&#x2022; [`name`](#name-9bddbc) - Optional String<br>Cookie Name. A case-sensitive cookie name

#### Policy Based Challenge Rule List Rules Spec Cookie Matchers Item

An [`item`](#item-951d5f) block (within [`policy_based_challenge.rule_list.rules.spec.cookie_matchers`](#matchers-cb349b)) supports the following:

<a id="values-093591"></a>&#x2022; [`exact_values`](#values-093591) - Optional List<br>Exact Values. A list of exact values to match the input against

<a id="values-447dbd"></a>&#x2022; [`regex_values`](#values-447dbd) - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="transformers-d98df1"></a>&#x2022; [`transformers`](#transformers-d98df1) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>[Enum: LOWER_CASE|UPPER_CASE|BASE64_DECODE|NORMALIZE_PATH|REMOVE_WHITESPACE|URL_DECODE|TRIM_LEFT|TRIM_RIGHT|TRIM] Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### Policy Based Challenge Rule List Rules Spec Domain Matcher

A [`domain_matcher`](#matcher-888f5a) block (within [`policy_based_challenge.rule_list.rules.spec`](#spec-fbd0f9)) supports the following:

<a id="values-3f5560"></a>&#x2022; [`exact_values`](#values-3f5560) - Optional List<br>Exact Values. A list of exact values to match the input against

<a id="values-f1499c"></a>&#x2022; [`regex_values`](#values-f1499c) - Optional List<br>Regex Values. A list of regular expressions to match the input against

#### Policy Based Challenge Rule List Rules Spec Headers

A [`headers`](#headers-1bea3b) block (within [`policy_based_challenge.rule_list.rules.spec`](#spec-fbd0f9)) supports the following:

<a id="present-55942e"></a>&#x2022; [`check_not_present`](#present-55942e) - Optional Block<br>Enable this option

<a id="present-4db691"></a>&#x2022; [`check_present`](#present-4db691) - Optional Block<br>Enable this option

<a id="matcher-1277ca"></a>&#x2022; [`invert_matcher`](#matcher-1277ca) - Optional Bool<br>Invert Header Matcher. Invert the match result

<a id="item-f6ed80"></a>&#x2022; [`item`](#item-f6ed80) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#item-f6ed80) below.

<a id="name-c77daf"></a>&#x2022; [`name`](#name-c77daf) - Optional String<br>Header Name. A case-insensitive HTTP header name

#### Policy Based Challenge Rule List Rules Spec Headers Item

An [`item`](#item-f6ed80) block (within [`policy_based_challenge.rule_list.rules.spec.headers`](#headers-1bea3b)) supports the following:

<a id="values-fc746b"></a>&#x2022; [`exact_values`](#values-fc746b) - Optional List<br>Exact Values. A list of exact values to match the input against

<a id="values-52c029"></a>&#x2022; [`regex_values`](#values-52c029) - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="transformers-12bdd9"></a>&#x2022; [`transformers`](#transformers-12bdd9) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>[Enum: LOWER_CASE|UPPER_CASE|BASE64_DECODE|NORMALIZE_PATH|REMOVE_WHITESPACE|URL_DECODE|TRIM_LEFT|TRIM_RIGHT|TRIM] Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### Policy Based Challenge Rule List Rules Spec HTTP Method

A [`http_method`](#method-9ab722) block (within [`policy_based_challenge.rule_list.rules.spec`](#spec-fbd0f9)) supports the following:

<a id="matcher-7f896b"></a>&#x2022; [`invert_matcher`](#matcher-7f896b) - Optional Bool<br>Invert Method Matcher. Invert the match result

<a id="methods-7a62d5"></a>&#x2022; [`methods`](#methods-7a62d5) - Optional List  Defaults to `ANY`<br>Possible values are `ANY`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`, `COPY`<br>[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY] Method List. List of methods values to match against

#### Policy Based Challenge Rule List Rules Spec IP Matcher

An [`ip_matcher`](#matcher-82616b) block (within [`policy_based_challenge.rule_list.rules.spec`](#spec-fbd0f9)) supports the following:

<a id="matcher-82cd23"></a>&#x2022; [`invert_matcher`](#matcher-82cd23) - Optional Bool<br>Invert IP Matcher. Invert the match result

<a id="sets-9dc015"></a>&#x2022; [`prefix_sets`](#sets-9dc015) - Optional Block<br>IP Prefix Sets. A list of references to ip_prefix_set objects<br>See [Prefix Sets](#sets-9dc015) below.

#### Policy Based Challenge Rule List Rules Spec IP Matcher Prefix Sets

A [`prefix_sets`](#sets-9dc015) block (within [`policy_based_challenge.rule_list.rules.spec.ip_matcher`](#matcher-82616b)) supports the following:

<a id="kind-190f46"></a>&#x2022; [`kind`](#kind-190f46) - Optional String<br>Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route')

<a id="name-434336"></a>&#x2022; [`name`](#name-434336) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-2b573f"></a>&#x2022; [`namespace`](#namespace-2b573f) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-86f089"></a>&#x2022; [`tenant`](#tenant-86f089) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="uid-2efcf4"></a>&#x2022; [`uid`](#uid-2efcf4) - Optional String<br>UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid

#### Policy Based Challenge Rule List Rules Spec IP Prefix List

An [`ip_prefix_list`](#list-537143) block (within [`policy_based_challenge.rule_list.rules.spec`](#spec-fbd0f9)) supports the following:

<a id="match-c9ff32"></a>&#x2022; [`invert_match`](#match-c9ff32) - Optional Bool<br>Invert Match Result. Invert the match result

<a id="prefixes-607b2d"></a>&#x2022; [`ip_prefixes`](#prefixes-607b2d) - Optional List<br>IPv4 Prefix List. List of IPv4 prefix strings

#### Policy Based Challenge Rule List Rules Spec Path

A [`path`](#path-acb3cc) block (within [`policy_based_challenge.rule_list.rules.spec`](#spec-fbd0f9)) supports the following:

<a id="values-485c76"></a>&#x2022; [`exact_values`](#values-485c76) - Optional List<br>Exact Values. A list of exact path values to match the input HTTP path against

<a id="matcher-003880"></a>&#x2022; [`invert_matcher`](#matcher-003880) - Optional Bool<br>Invert Path Matcher. Invert the match result

<a id="values-083d9f"></a>&#x2022; [`prefix_values`](#values-083d9f) - Optional List<br>Prefix Values. A list of path prefix values to match the input HTTP path against

<a id="values-4b2fb8"></a>&#x2022; [`regex_values`](#values-4b2fb8) - Optional List<br>Regex Values. A list of regular expressions to match the input HTTP path against

<a id="values-401ec9"></a>&#x2022; [`suffix_values`](#values-401ec9) - Optional List<br>Suffix Values. A list of path suffix values to match the input HTTP path against

<a id="transformers-c5ad13"></a>&#x2022; [`transformers`](#transformers-c5ad13) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>[Enum: LOWER_CASE|UPPER_CASE|BASE64_DECODE|NORMALIZE_PATH|REMOVE_WHITESPACE|URL_DECODE|TRIM_LEFT|TRIM_RIGHT|TRIM] Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### Policy Based Challenge Rule List Rules Spec Query Params

A [`query_params`](#params-04b1ad) block (within [`policy_based_challenge.rule_list.rules.spec`](#spec-fbd0f9)) supports the following:

<a id="present-3df5a5"></a>&#x2022; [`check_not_present`](#present-3df5a5) - Optional Block<br>Enable this option

<a id="present-aa54e6"></a>&#x2022; [`check_present`](#present-aa54e6) - Optional Block<br>Enable this option

<a id="matcher-c30f88"></a>&#x2022; [`invert_matcher`](#matcher-c30f88) - Optional Bool<br>Invert Query Parameter Matcher. Invert the match result

<a id="item-51a3bb"></a>&#x2022; [`item`](#item-51a3bb) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#item-51a3bb) below.

<a id="key-3709d7"></a>&#x2022; [`key`](#key-3709d7) - Optional String<br>Query Parameter Name. A case-sensitive HTTP query parameter name

#### Policy Based Challenge Rule List Rules Spec Query Params Item

An [`item`](#item-51a3bb) block (within [`policy_based_challenge.rule_list.rules.spec.query_params`](#params-04b1ad)) supports the following:

<a id="values-f4b36f"></a>&#x2022; [`exact_values`](#values-f4b36f) - Optional List<br>Exact Values. A list of exact values to match the input against

<a id="values-1b7a32"></a>&#x2022; [`regex_values`](#values-1b7a32) - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="transformers-f4e714"></a>&#x2022; [`transformers`](#transformers-f4e714) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>[Enum: LOWER_CASE|UPPER_CASE|BASE64_DECODE|NORMALIZE_PATH|REMOVE_WHITESPACE|URL_DECODE|TRIM_LEFT|TRIM_RIGHT|TRIM] Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### Policy Based Challenge Rule List Rules Spec TLS Fingerprint Matcher

A [`tls_fingerprint_matcher`](#matcher-3cbc4b) block (within [`policy_based_challenge.rule_list.rules.spec`](#spec-fbd0f9)) supports the following:

<a id="classes-e328f3"></a>&#x2022; [`classes`](#classes-e328f3) - Optional List  Defaults to `TLS_FINGERPRINT_NONE`<br>Possible values are `TLS_FINGERPRINT_NONE`, `ANY_MALICIOUS_FINGERPRINT`, `ADWARE`, `ADWIND`, `DRIDEX`, `GOOTKIT`, `GOZI`, `JBIFROST`, `QUAKBOT`, `RANSOMWARE`, `TROLDESH`, `TOFSEE`, `TORRENTLOCKER`, `TRICKBOT`<br>[Enum: TLS_FINGERPRINT_NONE|ANY_MALICIOUS_FINGERPRINT|ADWARE|ADWIND|DRIDEX|GOOTKIT|GOZI|JBIFROST|QUAKBOT|RANSOMWARE|TROLDESH|TOFSEE|TORRENTLOCKER|TRICKBOT] TLS fingerprint classes. A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against

<a id="values-02374f"></a>&#x2022; [`exact_values`](#values-02374f) - Optional List<br>Exact Values. A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against

<a id="values-2fc745"></a>&#x2022; [`excluded_values`](#values-2fc745) - Optional List<br>Excluded Values. A list of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can be used to skip known false positives when using one or more known TLS fingerprint classes in the enclosing matcher

#### Policy Based Challenge Temporary User Blocking

A [`temporary_user_blocking`](#blocking-9fdca7) block (within [`policy_based_challenge`](#policy-based-challenge)) supports the following:

<a id="page-dc34c2"></a>&#x2022; [`custom_page`](#page-dc34c2) - Optional String<br>Custom Message for Temporary Blocking. Custom message is of type `uri_ref`. Currently supported URL schemes is `string:///`. For `string:///` scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. 'Blocked.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '`<p>` Blocked `</p>`'. Base64 encoded string for this HTML is 'PHA+IFBsZWFzZSBXYWl0IDwvcD4='

#### Protected Cookies

A [`protected_cookies`](#protected-cookies) block supports the following:

<a id="protected-cookies-add-httponly"></a>&#x2022; [`add_httponly`](#protected-cookies-add-httponly) - Optional Block<br>Enable this option

<a id="protected-cookies-add-secure"></a>&#x2022; [`add_secure`](#protected-cookies-add-secure) - Optional Block<br>Enable this option

<a id="protection-51c741"></a>&#x2022; [`disable_tampering_protection`](#protection-51c741) - Optional Block<br>Enable this option

<a id="protection-d59c9f"></a>&#x2022; [`enable_tampering_protection`](#protection-d59c9f) - Optional Block<br>Enable this option

<a id="protected-cookies-ignore-httponly"></a>&#x2022; [`ignore_httponly`](#protected-cookies-ignore-httponly) - Optional Block<br>Enable this option

<a id="protected-cookies-ignore-max-age"></a>&#x2022; [`ignore_max_age`](#protected-cookies-ignore-max-age) - Optional Block<br>Enable this option

<a id="protected-cookies-ignore-samesite"></a>&#x2022; [`ignore_samesite`](#protected-cookies-ignore-samesite) - Optional Block<br>Enable this option

<a id="protected-cookies-ignore-secure"></a>&#x2022; [`ignore_secure`](#protected-cookies-ignore-secure) - Optional Block<br>Enable this option

<a id="protected-cookies-max-age-value"></a>&#x2022; [`max_age_value`](#protected-cookies-max-age-value) - Optional Number<br>Add Max Age. Add max age attribute

<a id="protected-cookies-name"></a>&#x2022; [`name`](#protected-cookies-name) - Optional String<br>Cookie Name. Name of the Cookie

<a id="protected-cookies-samesite-lax"></a>&#x2022; [`samesite_lax`](#protected-cookies-samesite-lax) - Optional Block<br>Enable this option

<a id="protected-cookies-samesite-none"></a>&#x2022; [`samesite_none`](#protected-cookies-samesite-none) - Optional Block<br>Enable this option

<a id="protected-cookies-samesite-strict"></a>&#x2022; [`samesite_strict`](#protected-cookies-samesite-strict) - Optional Block<br>Enable this option

#### Rate Limit

A [`rate_limit`](#rate-limit) block supports the following:

<a id="rate-limit-custom-ip-allowed-list"></a>&#x2022; [`custom_ip_allowed_list`](#rate-limit-custom-ip-allowed-list) - Optional Block<br>Custom IP Allowed List. IP Allowed list using existing ip_prefix_set objects<br>See [Custom IP Allowed List](#rate-limit-custom-ip-allowed-list) below.

<a id="rate-limit-ip-allowed-list"></a>&#x2022; [`ip_allowed_list`](#rate-limit-ip-allowed-list) - Optional Block<br>IPv4 Prefix List. List of IPv4 prefixes that represent an endpoint<br>See [IP Allowed List](#rate-limit-ip-allowed-list) below.

<a id="rate-limit-no-ip-allowed-list"></a>&#x2022; [`no_ip_allowed_list`](#rate-limit-no-ip-allowed-list) - Optional Block<br>Enable this option

<a id="rate-limit-no-policies"></a>&#x2022; [`no_policies`](#rate-limit-no-policies) - Optional Block<br>Enable this option

<a id="rate-limit-policies"></a>&#x2022; [`policies`](#rate-limit-policies) - Optional Block<br>Rate Limiter Policy List. List of rate limiter policies to be applied<br>See [Policies](#rate-limit-policies) below.

<a id="rate-limit-rate-limiter"></a>&#x2022; [`rate_limiter`](#rate-limit-rate-limiter) - Optional Block<br>Rate Limit Value. A tuple consisting of a rate limit period unit and the total number of allowed requests for that period<br>See [Rate Limiter](#rate-limit-rate-limiter) below.

#### Rate Limit Custom IP Allowed List

A [`custom_ip_allowed_list`](#rate-limit-custom-ip-allowed-list) block (within [`rate_limit`](#rate-limit)) supports the following:

<a id="prefixes-266335"></a>&#x2022; [`rate_limiter_allowed_prefixes`](#prefixes-266335) - Optional Block<br>List of IP Prefix Sets. References to ip_prefix_set objects. Requests from source IP addresses that are covered by one of the allowed IP Prefixes are not subjected to rate limiting<br>See [Rate Limiter Allowed Prefixes](#prefixes-266335) below.

#### Rate Limit Custom IP Allowed List Rate Limiter Allowed Prefixes

A [`rate_limiter_allowed_prefixes`](#prefixes-266335) block (within [`rate_limit.custom_ip_allowed_list`](#rate-limit-custom-ip-allowed-list)) supports the following:

<a id="name-2e45f5"></a>&#x2022; [`name`](#name-2e45f5) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-cab019"></a>&#x2022; [`namespace`](#namespace-cab019) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-d99ce2"></a>&#x2022; [`tenant`](#tenant-d99ce2) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Rate Limit IP Allowed List

An [`ip_allowed_list`](#rate-limit-ip-allowed-list) block (within [`rate_limit`](#rate-limit)) supports the following:

<a id="rate-limit-ip-allowed-list-prefixes"></a>&#x2022; [`prefixes`](#rate-limit-ip-allowed-list-prefixes) - Optional List<br>IPv4 Prefix List. List of IPv4 prefixes that represent an endpoint

#### Rate Limit Policies

A [`policies`](#rate-limit-policies) block (within [`rate_limit`](#rate-limit)) supports the following:

<a id="rate-limit-policies-policies"></a>&#x2022; [`policies`](#rate-limit-policies-policies) - Optional Block<br>Rate Limiter Policies. Ordered list of rate limiter policies<br>See [Policies](#rate-limit-policies-policies) below.

#### Rate Limit Policies Policies

A [`policies`](#rate-limit-policies-policies) block (within [`rate_limit.policies`](#rate-limit-policies)) supports the following:

<a id="rate-limit-policies-policies-name"></a>&#x2022; [`name`](#rate-limit-policies-policies-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="rate-limit-policies-policies-namespace"></a>&#x2022; [`namespace`](#rate-limit-policies-policies-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="rate-limit-policies-policies-tenant"></a>&#x2022; [`tenant`](#rate-limit-policies-policies-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Rate Limit Rate Limiter

A [`rate_limiter`](#rate-limit-rate-limiter) block (within [`rate_limit`](#rate-limit)) supports the following:

<a id="rate-limit-rate-limiter-action-block"></a>&#x2022; [`action_block`](#rate-limit-rate-limiter-action-block) - Optional Block<br>Rate Limit Block Action. Action where a user is blocked from making further requests after exceeding rate limit threshold<br>See [Action Block](#rate-limit-rate-limiter-action-block) below.

<a id="multiplier-1bc2e7"></a>&#x2022; [`burst_multiplier`](#multiplier-1bc2e7) - Optional Number<br>Burst Multiplier. The maximum burst of requests to accommodate, expressed as a multiple of the rate

<a id="rate-limit-rate-limiter-disabled"></a>&#x2022; [`disabled`](#rate-limit-rate-limiter-disabled) - Optional Block<br>Enable this option

<a id="rate-limit-rate-limiter-leaky-bucket"></a>&#x2022; [`leaky_bucket`](#rate-limit-rate-limiter-leaky-bucket) - Optional Block<br>Leaky Bucket Rate Limiter. Leaky-Bucket is the default rate limiter algorithm for F5

<a id="multiplier-07ace4"></a>&#x2022; [`period_multiplier`](#multiplier-07ace4) - Optional Number<br>Periods. This setting, combined with Per Period units, provides a duration

<a id="rate-limit-rate-limiter-token-bucket"></a>&#x2022; [`token_bucket`](#rate-limit-rate-limiter-token-bucket) - Optional Block<br>Token Bucket Rate Limiter. Token-Bucket is a rate limiter algorithm that is stricter with enforcing limits

<a id="rate-limit-rate-limiter-total-number"></a>&#x2022; [`total_number`](#rate-limit-rate-limiter-total-number) - Optional Number<br>Number Of Requests. The total number of allowed requests per rate-limiting period

<a id="rate-limit-rate-limiter-unit"></a>&#x2022; [`unit`](#rate-limit-rate-limiter-unit) - Optional String  Defaults to `SECOND`<br>Possible values are `SECOND`, `MINUTE`, `HOUR`<br>[Enum: SECOND|MINUTE|HOUR] Rate Limit Period Unit. Unit for the period per which the rate limit is applied. - SECOND: Second Rate limit period unit is seconds - MINUTE: Minute Rate limit period unit is minutes - HOUR: Hour Rate limit period unit is hours - DAY: Day Rate limit period unit is days

#### Rate Limit Rate Limiter Action Block

An [`action_block`](#rate-limit-rate-limiter-action-block) block (within [`rate_limit.rate_limiter`](#rate-limit-rate-limiter)) supports the following:

<a id="hours-fe2333"></a>&#x2022; [`hours`](#hours-fe2333) - Optional Block<br>Hours. Input Duration Hours<br>See [Hours](#hours-fe2333) below.

<a id="minutes-c83f64"></a>&#x2022; [`minutes`](#minutes-c83f64) - Optional Block<br>Minutes. Input Duration Minutes<br>See [Minutes](#minutes-c83f64) below.

<a id="seconds-8810ec"></a>&#x2022; [`seconds`](#seconds-8810ec) - Optional Block<br>Seconds. Input Duration Seconds<br>See [Seconds](#seconds-8810ec) below.

#### Rate Limit Rate Limiter Action Block Hours

A [`hours`](#hours-fe2333) block (within [`rate_limit.rate_limiter.action_block`](#rate-limit-rate-limiter-action-block)) supports the following:

<a id="duration-617314"></a>&#x2022; [`duration`](#duration-617314) - Optional Number<br>Duration

#### Rate Limit Rate Limiter Action Block Minutes

A [`minutes`](#minutes-c83f64) block (within [`rate_limit.rate_limiter.action_block`](#rate-limit-rate-limiter-action-block)) supports the following:

<a id="duration-534bd9"></a>&#x2022; [`duration`](#duration-534bd9) - Optional Number<br>Duration

#### Rate Limit Rate Limiter Action Block Seconds

A [`seconds`](#seconds-8810ec) block (within [`rate_limit.rate_limiter.action_block`](#rate-limit-rate-limiter-action-block)) supports the following:

<a id="duration-dfe2a4"></a>&#x2022; [`duration`](#duration-dfe2a4) - Optional Number<br>Duration

#### Ring Hash

A [`ring_hash`](#ring-hash) block supports the following:

<a id="ring-hash-hash-policy"></a>&#x2022; [`hash_policy`](#ring-hash-hash-policy) - Optional Block<br>Hash Policy. Specifies a list of hash policies to use for ring hash load balancing. Each hash policy is evaluated individually and the combined result is used to route the request<br>See [Hash Policy](#ring-hash-hash-policy) below.

#### Ring Hash Hash Policy

A [`hash_policy`](#ring-hash-hash-policy) block (within [`ring_hash`](#ring-hash)) supports the following:

<a id="ring-hash-hash-policy-cookie"></a>&#x2022; [`cookie`](#ring-hash-hash-policy-cookie) - Optional Block<br>Hashing using Cookie. Two types of cookie affinity: 1. Passive. Takes a cookie that's present in the cookies header and hashes on its value. 2. Generated. Generates and sets a cookie with an expiration (TTL) on the first request from the client in its response to the client, based on the endpoint the request gets sent to. The client then presents this on the next and all subsequent requests. The hash of this is sufficient to ensure these requests get sent to the same endpoint. The cookie is generated by hashing the source and destination ports and addresses so that multiple independent HTTP2 streams on the same connection will independently receive the same cookie, even if they arrive simultaneously<br>See [Cookie](#ring-hash-hash-policy-cookie) below.

<a id="ring-hash-hash-policy-header-name"></a>&#x2022; [`header_name`](#ring-hash-hash-policy-header-name) - Optional String<br>Header Name. The name or key of the request header that will be used to obtain the hash key

<a id="ring-hash-hash-policy-source-ip"></a>&#x2022; [`source_ip`](#ring-hash-hash-policy-source-ip) - Optional Bool<br>Source IP. Hash based on source IP address

<a id="ring-hash-hash-policy-terminal"></a>&#x2022; [`terminal`](#ring-hash-hash-policy-terminal) - Optional Bool<br>Terminal. Specify if its a terminal policy

#### Ring Hash Hash Policy Cookie

A [`cookie`](#ring-hash-hash-policy-cookie) block (within [`ring_hash.hash_policy`](#ring-hash-hash-policy)) supports the following:

<a id="httponly-54c45b"></a>&#x2022; [`add_httponly`](#httponly-54c45b) - Optional Block<br>Enable this option

<a id="ring-hash-hash-policy-cookie-add-secure"></a>&#x2022; [`add_secure`](#ring-hash-hash-policy-cookie-add-secure) - Optional Block<br>Enable this option

<a id="httponly-7f2aea"></a>&#x2022; [`ignore_httponly`](#httponly-7f2aea) - Optional Block<br>Enable this option

<a id="samesite-106140"></a>&#x2022; [`ignore_samesite`](#samesite-106140) - Optional Block<br>Enable this option

<a id="secure-febf51"></a>&#x2022; [`ignore_secure`](#secure-febf51) - Optional Block<br>Enable this option

<a id="ring-hash-hash-policy-cookie-name"></a>&#x2022; [`name`](#ring-hash-hash-policy-cookie-name) - Optional String<br>Name. The name of the cookie that will be used to obtain the hash key. If the cookie is not present and TTL below is not set, no hash will be produced

<a id="ring-hash-hash-policy-cookie-path"></a>&#x2022; [`path`](#ring-hash-hash-policy-cookie-path) - Optional String<br>Path. The name of the path for the cookie. If no path is specified here, no path will be set for the cookie

<a id="lax-749b7e"></a>&#x2022; [`samesite_lax`](#lax-749b7e) - Optional Block<br>Enable this option

<a id="none-5bbed3"></a>&#x2022; [`samesite_none`](#none-5bbed3) - Optional Block<br>Enable this option

<a id="strict-3e550d"></a>&#x2022; [`samesite_strict`](#strict-3e550d) - Optional Block<br>Enable this option

<a id="ring-hash-hash-policy-cookie-ttl"></a>&#x2022; [`ttl`](#ring-hash-hash-policy-cookie-ttl) - Optional Number<br>TTL. If specified, a cookie with the TTL will be generated if the cookie is not present. If the TTL is present and zero, the generated cookie will be a session cookie. TTL value is in milliseconds

#### Routes

A [`routes`](#routes) block supports the following:

<a id="routes-custom-route-object"></a>&#x2022; [`custom_route_object`](#routes-custom-route-object) - Optional Block<br>Custom Route Object. A custom route uses a route object created outside of this view<br>See [Custom Route Object](#routes-custom-route-object) below.

<a id="routes-direct-response-route"></a>&#x2022; [`direct_response_route`](#routes-direct-response-route) - Optional Block<br>Direct Response Route. A direct response route matches on path, incoming header, incoming port and/or HTTP method and responds directly to the matching traffic<br>See [Direct Response Route](#routes-direct-response-route) below.

<a id="routes-redirect-route"></a>&#x2022; [`redirect_route`](#routes-redirect-route) - Optional Block<br>Redirect Route. A redirect route matches on path, incoming header, incoming port and/or HTTP method and redirects the matching traffic to a different URL<br>See [Redirect Route](#routes-redirect-route) below.

<a id="routes-simple-route"></a>&#x2022; [`simple_route`](#routes-simple-route) - Optional Block<br>Simple Route. A simple route matches on path, incoming header, incoming port and/or HTTP method and forwards the matching traffic to the associated pools<br>See [Simple Route](#routes-simple-route) below.

#### Routes Custom Route Object

A [`custom_route_object`](#routes-custom-route-object) block (within [`routes`](#routes)) supports the following:

<a id="routes-custom-route-object-route-ref"></a>&#x2022; [`route_ref`](#routes-custom-route-object-route-ref) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Route Ref](#routes-custom-route-object-route-ref) below.

#### Routes Custom Route Object Route Ref

A [`route_ref`](#routes-custom-route-object-route-ref) block (within [`routes.custom_route_object`](#routes-custom-route-object)) supports the following:

<a id="name-82f318"></a>&#x2022; [`name`](#name-82f318) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-7ec6c4"></a>&#x2022; [`namespace`](#namespace-7ec6c4) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-a123ba"></a>&#x2022; [`tenant`](#tenant-a123ba) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Routes Direct Response Route

A [`direct_response_route`](#routes-direct-response-route) block (within [`routes`](#routes)) supports the following:

<a id="routes-direct-response-route-headers"></a>&#x2022; [`headers`](#routes-direct-response-route-headers) - Optional Block<br>Headers. List of (key, value) headers<br>See [Headers](#routes-direct-response-route-headers) below.

<a id="method-aec314"></a>&#x2022; [`http_method`](#method-aec314) - Optional String  Defaults to `ANY`<br>Possible values are `ANY`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`, `COPY`<br>[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY] HTTP Method. Specifies the HTTP method used to access a resource. Any HTTP Method

<a id="port-9bcff1"></a>&#x2022; [`incoming_port`](#port-9bcff1) - Optional Block<br>Port to Match. Port match of the request can be a range or a specific port<br>See [Incoming Port](#port-9bcff1) below.

<a id="routes-direct-response-route-path"></a>&#x2022; [`path`](#routes-direct-response-route-path) - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#routes-direct-response-route-path) below.

<a id="response-d0dcbd"></a>&#x2022; [`route_direct_response`](#response-d0dcbd) - Optional Block<br>Direct Response. Send this direct response in case of route match action is direct response<br>See [Route Direct Response](#response-d0dcbd) below.

#### Routes Direct Response Route Headers

A [`headers`](#routes-direct-response-route-headers) block (within [`routes.direct_response_route`](#routes-direct-response-route)) supports the following:

<a id="exact-23a9da"></a>&#x2022; [`exact`](#exact-23a9da) - Optional String<br>Exact. Header value to match exactly

<a id="match-16404c"></a>&#x2022; [`invert_match`](#match-16404c) - Optional Bool<br>NOT of match. Invert the result of the match to detect missing header or non-matching value

<a id="name-77fba0"></a>&#x2022; [`name`](#name-77fba0) - Optional String<br>Name. Name of the header

<a id="presence-72a565"></a>&#x2022; [`presence`](#presence-72a565) - Optional Bool<br>Presence. If true, check for presence of header

<a id="regex-4cdd3d"></a>&#x2022; [`regex`](#regex-4cdd3d) - Optional String<br>Regex. Regex match of the header value in re2 format

#### Routes Direct Response Route Incoming Port

An [`incoming_port`](#port-9bcff1) block (within [`routes.direct_response_route`](#routes-direct-response-route)) supports the following:

<a id="match-ba3425"></a>&#x2022; [`no_port_match`](#match-ba3425) - Optional Block<br>Enable this option

<a id="port-9debaf"></a>&#x2022; [`port`](#port-9debaf) - Optional Number<br>Port. Exact Port to match

<a id="ranges-b7be24"></a>&#x2022; [`port_ranges`](#ranges-b7be24) - Optional String<br>Port range. Port range to match

#### Routes Direct Response Route Path

A [`path`](#routes-direct-response-route-path) block (within [`routes.direct_response_route`](#routes-direct-response-route)) supports the following:

<a id="routes-direct-response-route-path-path"></a>&#x2022; [`path`](#routes-direct-response-route-path-path) - Optional String<br>Exact. Exact path value to match

<a id="prefix-ff3976"></a>&#x2022; [`prefix`](#prefix-ff3976) - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

<a id="routes-direct-response-route-path-regex"></a>&#x2022; [`regex`](#routes-direct-response-route-path-regex) - Optional String<br>Regex. Regular expression of path match (e.g. the value .* will match on all paths)

#### Routes Direct Response Route Route Direct Response

A [`route_direct_response`](#response-d0dcbd) block (within [`routes.direct_response_route`](#routes-direct-response-route)) supports the following:

<a id="encoded-a56f81"></a>&#x2022; [`response_body_encoded`](#encoded-a56f81) - Optional String<br>Response Body. Response body to send. Currently supported URL schemes is string:/// for which message should be encoded in Base64 format. The message can be either plain text or HTML. E.g. '`<p>` Access Denied `</p>`'. Base64 encoded string URL for this is string:///PHA+IEFjY2VzcyBEZW5pZWQgPC9wPg==

<a id="code-1bc88c"></a>&#x2022; [`response_code`](#code-1bc88c) - Optional Number<br>Response Code. response code to send

#### Routes Redirect Route

A [`redirect_route`](#routes-redirect-route) block (within [`routes`](#routes)) supports the following:

<a id="routes-redirect-route-headers"></a>&#x2022; [`headers`](#routes-redirect-route-headers) - Optional Block<br>Headers. List of (key, value) headers<br>See [Headers](#routes-redirect-route-headers) below.

<a id="routes-redirect-route-http-method"></a>&#x2022; [`http_method`](#routes-redirect-route-http-method) - Optional String  Defaults to `ANY`<br>Possible values are `ANY`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`, `COPY`<br>[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY] HTTP Method. Specifies the HTTP method used to access a resource. Any HTTP Method

<a id="routes-redirect-route-incoming-port"></a>&#x2022; [`incoming_port`](#routes-redirect-route-incoming-port) - Optional Block<br>Port to Match. Port match of the request can be a range or a specific port<br>See [Incoming Port](#routes-redirect-route-incoming-port) below.

<a id="routes-redirect-route-path"></a>&#x2022; [`path`](#routes-redirect-route-path) - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#routes-redirect-route-path) below.

<a id="routes-redirect-route-route-redirect"></a>&#x2022; [`route_redirect`](#routes-redirect-route-route-redirect) - Optional Block<br>Redirect. route redirect parameters when match action is redirect<br>See [Route Redirect](#routes-redirect-route-route-redirect) below.

#### Routes Redirect Route Headers

A [`headers`](#routes-redirect-route-headers) block (within [`routes.redirect_route`](#routes-redirect-route)) supports the following:

<a id="routes-redirect-route-headers-exact"></a>&#x2022; [`exact`](#routes-redirect-route-headers-exact) - Optional String<br>Exact. Header value to match exactly

<a id="match-b95584"></a>&#x2022; [`invert_match`](#match-b95584) - Optional Bool<br>NOT of match. Invert the result of the match to detect missing header or non-matching value

<a id="routes-redirect-route-headers-name"></a>&#x2022; [`name`](#routes-redirect-route-headers-name) - Optional String<br>Name. Name of the header

<a id="routes-redirect-route-headers-presence"></a>&#x2022; [`presence`](#routes-redirect-route-headers-presence) - Optional Bool<br>Presence. If true, check for presence of header

<a id="routes-redirect-route-headers-regex"></a>&#x2022; [`regex`](#routes-redirect-route-headers-regex) - Optional String<br>Regex. Regex match of the header value in re2 format

#### Routes Redirect Route Incoming Port

An [`incoming_port`](#routes-redirect-route-incoming-port) block (within [`routes.redirect_route`](#routes-redirect-route)) supports the following:

<a id="match-56f681"></a>&#x2022; [`no_port_match`](#match-56f681) - Optional Block<br>Enable this option

<a id="port-02a88d"></a>&#x2022; [`port`](#port-02a88d) - Optional Number<br>Port. Exact Port to match

<a id="ranges-33e473"></a>&#x2022; [`port_ranges`](#ranges-33e473) - Optional String<br>Port range. Port range to match

#### Routes Redirect Route Path

A [`path`](#routes-redirect-route-path) block (within [`routes.redirect_route`](#routes-redirect-route)) supports the following:

<a id="routes-redirect-route-path-path"></a>&#x2022; [`path`](#routes-redirect-route-path-path) - Optional String<br>Exact. Exact path value to match

<a id="routes-redirect-route-path-prefix"></a>&#x2022; [`prefix`](#routes-redirect-route-path-prefix) - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

<a id="routes-redirect-route-path-regex"></a>&#x2022; [`regex`](#routes-redirect-route-path-regex) - Optional String<br>Regex. Regular expression of path match (e.g. the value .* will match on all paths)

#### Routes Redirect Route Route Redirect

A [`route_redirect`](#routes-redirect-route-route-redirect) block (within [`routes.redirect_route`](#routes-redirect-route)) supports the following:

<a id="redirect-bd044d"></a>&#x2022; [`host_redirect`](#redirect-bd044d) - Optional String<br>Host. swap host part of incoming URL in redirect URL

<a id="redirect-2ae47a"></a>&#x2022; [`path_redirect`](#redirect-2ae47a) - Optional String<br>Path. swap path part of incoming URL in redirect URL

<a id="rewrite-a81c41"></a>&#x2022; [`prefix_rewrite`](#rewrite-a81c41) - Optional String<br>Prefix Rewrite. In Redirect response, the matched prefix (or path) should be swapped with this value. This option allows redirect URLs be dynamically created based on the request

<a id="redirect-f23979"></a>&#x2022; [`proto_redirect`](#redirect-f23979) - Optional String<br>Protocol. swap protocol part of incoming URL in redirect URL The protocol can be swapped with either HTTP or HTTPS When incoming-proto option is specified, swapping of protocol is not done

<a id="params-0941dc"></a>&#x2022; [`remove_all_params`](#params-0941dc) - Optional Block<br>Enable this option

<a id="params-94a828"></a>&#x2022; [`replace_params`](#params-94a828) - Optional String<br>Replace All Parameters

<a id="code-d55c43"></a>&#x2022; [`response_code`](#code-d55c43) - Optional Number<br>Response Code. The HTTP status code to use in the redirect response

<a id="params-f96588"></a>&#x2022; [`retain_all_params`](#params-f96588) - Optional Block<br>Enable this option

#### Routes Simple Route

A [`simple_route`](#routes-simple-route) block (within [`routes`](#routes)) supports the following:

<a id="routes-simple-route-advanced-options"></a>&#x2022; [`advanced_options`](#routes-simple-route-advanced-options) - Optional Block<br>Advanced Route Options. Configure advanced options for route like path rewrite, hash policy, etc<br>See [Advanced Options](#routes-simple-route-advanced-options) below.

<a id="routes-simple-route-auto-host-rewrite"></a>&#x2022; [`auto_host_rewrite`](#routes-simple-route-auto-host-rewrite) - Optional Block<br>Enable this option

<a id="rewrite-706535"></a>&#x2022; [`disable_host_rewrite`](#rewrite-706535) - Optional Block<br>Enable this option

<a id="routes-simple-route-headers"></a>&#x2022; [`headers`](#routes-simple-route-headers) - Optional Block<br>Headers. List of (key, value) headers<br>See [Headers](#routes-simple-route-headers) below.

<a id="routes-simple-route-host-rewrite"></a>&#x2022; [`host_rewrite`](#routes-simple-route-host-rewrite) - Optional String<br>Host Rewrite Value. Host header will be swapped with this value

<a id="routes-simple-route-http-method"></a>&#x2022; [`http_method`](#routes-simple-route-http-method) - Optional String  Defaults to `ANY`<br>Possible values are `ANY`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`, `COPY`<br>[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY] HTTP Method. Specifies the HTTP method used to access a resource. Any HTTP Method

<a id="routes-simple-route-incoming-port"></a>&#x2022; [`incoming_port`](#routes-simple-route-incoming-port) - Optional Block<br>Port to Match. Port match of the request can be a range or a specific port<br>See [Incoming Port](#routes-simple-route-incoming-port) below.

<a id="routes-simple-route-origin-pools"></a>&#x2022; [`origin_pools`](#routes-simple-route-origin-pools) - Optional Block<br>Origin Pools. Origin Pools for this route<br>See [Origin Pools](#routes-simple-route-origin-pools) below.

<a id="routes-simple-route-path"></a>&#x2022; [`path`](#routes-simple-route-path) - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#routes-simple-route-path) below.

<a id="routes-simple-route-query-params"></a>&#x2022; [`query_params`](#routes-simple-route-query-params) - Optional Block<br>Query Parameters. Handling of incoming query parameters in simple route<br>See [Query Params](#routes-simple-route-query-params) below.

#### Routes Simple Route Advanced Options

An [`advanced_options`](#routes-simple-route-advanced-options) block (within [`routes.simple_route`](#routes-simple-route)) supports the following:

<a id="firewall-b8f7c9"></a>&#x2022; [`app_firewall`](#firewall-b8f7c9) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [App Firewall](#firewall-b8f7c9) below.

<a id="injection-23f0bd"></a>&#x2022; [`bot_defense_javascript_injection`](#injection-23f0bd) - Optional Block<br>Bot Defense Javascript Injection Configuration for inline deployments. Bot Defense Javascript Injection Configuration for inline bot defense deployments<br>See [Bot Defense Javascript Injection](#injection-23f0bd) below.

<a id="policy-23a3f6"></a>&#x2022; [`buffer_policy`](#policy-23a3f6) - Optional Block<br>Buffer Configuration. Some upstream applications are not capable of handling streamed data. This config enables buffering the entire request before sending to upstream application. We can specify the maximum buffer size and buffer interval with this config. Buffering can be enabled and disabled at VirtualHost and Route levels Route level buffer configuration takes precedence<br>See [Buffer Policy](#policy-23a3f6) below.

<a id="buffering-44c193"></a>&#x2022; [`common_buffering`](#buffering-44c193) - Optional Block<br>Enable this option

<a id="policy-b912b0"></a>&#x2022; [`common_hash_policy`](#policy-b912b0) - Optional Block<br>Enable this option

<a id="policy-ba853e"></a>&#x2022; [`cors_policy`](#policy-ba853e) - Optional Block<br>CORS Policy. Cross-Origin Resource Sharing requests configuration specified at Virtual-host or Route level. Route level configuration takes precedence. An example of an Cross origin HTTP request GET /resources/public-data/ HTTP/1.1 Host: bar.other User-Agent: Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.5; en-US; rv:1.9.1b3pre) Gecko/20081130 Minefield/3.1b3pre Accept: text/HTML,application/xhtml+XML,application/XML;q=0.9,*/*;q=0.8 Accept-Language: en-us,en;q=0.5 Accept-Encoding: gzip,deflate Accept-Charset: ISO-8859-1,utf-8;q=0.7,*;q=0.7 Connection: keep-alive Referrer: `HTTP://foo.example/examples/access-control/simpleXSInvocation.HTML` Origin: `HTTP://foo.example` HTTP/1.1 200 OK Date: Mon, 01 Dec 2008 00:23:53 GMT Server: Apache/2.0.61 Access-Control-Allow-Origin: \* Keep-Alive: timeout=2, max=100 Connection: Keep-Alive Transfer-Encoding: chunked Content-Type: application/XML An example for cross origin HTTP OPTIONS request with Access-Control-Request-* header OPTIONS /resources/post-here/ HTTP/1.1 Host: bar.other User-Agent: Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.5; en-US; rv:1.9.1b3pre) Gecko/20081130 Minefield/3.1b3pre Accept: text/HTML,application/xhtml+XML,application/XML;q=0.9,*/*;q=0.8 Accept-Language: en-us,en;q=0.5 Accept-Encoding: gzip,deflate Accept-Charset: ISO-8859-1,utf-8;q=0.7,*;q=0.7 Connection: keep-alive Origin: `HTTP://foo.example` Access-Control-Request-Method: POST Access-Control-Request-Headers: X-PINGOTHER, Content-Type HTTP/1.1 204 No Content Date: Mon, 01 Dec 2008 01:15:39 GMT Server: Apache/2.0.61 (Unix) Access-Control-Allow-Origin: `HTTP://foo.example` Access-Control-Allow-Methods: POST, GET, OPTIONS Access-Control-Allow-Headers: X-PINGOTHER, Content-Type Access-Control-Max-Age: 86400 Vary: Accept-Encoding, Origin Keep-Alive: timeout=2, max=100 Connection: Keep-Alive<br>See [CORS Policy](#policy-ba853e) below.

<a id="policy-7816d7"></a>&#x2022; [`csrf_policy`](#policy-7816d7) - Optional Block<br>CSRF Policy. To mitigate CSRF attack , the policy checks where a request is coming from to determine if the request's origin is the same as its detination.The policy relies on two pieces of information used in determining if a request originated from the same host. 1. The origin that caused the user agent to issue the request (source origin). 2. The origin that the request is going to (target origin). When the policy evaluating a request, it ensures both pieces of information are present and compare their values. If the source origin is missing or origins do not match the request is rejected. The exception to this being if the source-origin has been added to they policy as valid. Because CSRF attacks specifically target state-changing requests, the policy only acts on the HTTP requests that have state-changing method (PUT,POST, etc.)<br>See [CSRF Policy](#policy-7816d7) below.

<a id="policy-70b68a"></a>&#x2022; [`default_retry_policy`](#policy-70b68a) - Optional Block<br>Enable this option

<a id="add-11129b"></a>&#x2022; [`disable_location_add`](#add-11129b) - Optional Bool<br>Disable Location Addition. disables append of x-volterra-location = `<RE-site-name>` at route level, if it is configured at virtual-host level. This configuration is ignored on CE sites

<a id="mirroring-e37294"></a>&#x2022; [`disable_mirroring`](#mirroring-e37294) - Optional Block<br>Enable this option

<a id="rewrite-8c52ee"></a>&#x2022; [`disable_prefix_rewrite`](#rewrite-8c52ee) - Optional Block<br>Enable this option

<a id="spdy-c4a11a"></a>&#x2022; [`disable_spdy`](#spdy-c4a11a) - Optional Block<br>Enable this option

<a id="waf-afaac0"></a>&#x2022; [`disable_waf`](#waf-afaac0) - Optional Block<br>Enable this option

<a id="config-b3faa9"></a>&#x2022; [`disable_web_socket_config`](#config-b3faa9) - Optional Block<br>Enable this option

<a id="cluster-f8e26f"></a>&#x2022; [`do_not_retract_cluster`](#cluster-f8e26f) - Optional Block<br>Enable this option

<a id="spdy-676c6f"></a>&#x2022; [`enable_spdy`](#spdy-676c6f) - Optional Block<br>Enable this option

<a id="subsets-b6a9d9"></a>&#x2022; [`endpoint_subsets`](#subsets-b6a9d9) - Optional Block<br>Origin Servers Subsets. Upstream origin pool may be configured to divide its origin servers into subsets based on metadata attached to the origin servers. Routes may then specify the metadata that a endpoint must match in order to be selected by the load balancer For origin servers which are discovered in K8S or Consul cluster, the label of the service is merged with endpoint's labels. In case of Consul, the label is derived from the 'Tag' field. For labels that are common between configured endpoint and discovered service, labels from discovered service takes precedence. List of key-value pairs that will be used as matching metadata. Only those origin servers of upstream origin pool which match this metadata will be selected for load balancing

<a id="injection-46208e"></a>&#x2022; [`inherited_bot_defense_javascript_injection`](#injection-46208e) - Optional Block<br>Enable this option

<a id="waf-0043f0"></a>&#x2022; [`inherited_waf`](#waf-0043f0) - Optional Block<br>Enable this option

<a id="exclusion-0ba7d9"></a>&#x2022; [`inherited_waf_exclusion`](#exclusion-0ba7d9) - Optional Block<br>Enable this option

<a id="policy-f5e84d"></a>&#x2022; [`mirror_policy`](#policy-f5e84d) - Optional Block<br>Mirror Policy. MirrorPolicy is used for shadowing traffic from one origin pool to another. The approach used is 'fire and forget', meaning it will not wait for the shadow origin pool to respond before returning the response from the primary origin pool. All normal statistics are collected for the shadow origin pool making this feature useful for testing and troubleshooting<br>See [Mirror Policy](#policy-f5e84d) below.

<a id="policy-ad7b2b"></a>&#x2022; [`no_retry_policy`](#policy-ad7b2b) - Optional Block<br>Enable this option

<a id="rewrite-ffbe86"></a>&#x2022; [`prefix_rewrite`](#rewrite-ffbe86) - Optional String<br>Enable Prefix Rewrite. prefix_rewrite indicates that during forwarding, the matched prefix (or path) should be swapped with its value. When using regex path matching, the entire path (not including the query string) will be swapped with this value

<a id="priority-f6c73d"></a>&#x2022; [`priority`](#priority-f6c73d) - Optional String  Defaults to `DEFAULT`<br>Possible values are `DEFAULT`, `HIGH`<br>[Enum: DEFAULT|HIGH] Routing Priority. Priority routing for each request. Different connection pools are used based on the priority selected for the request. Also, circuit-breaker configuration at destination cluster is chosen based on selected priority. Default routing mechanism High-Priority routing mechanism

<a id="rewrite-c628a7"></a>&#x2022; [`regex_rewrite`](#rewrite-c628a7) - Optional Block<br>Regex Match Rewrite. RegexMatchRewrite describes how to match a string and then produce a new string using a regular expression and a substitution string<br>See [Regex Rewrite](#rewrite-c628a7) below.

<a id="add-818453"></a>&#x2022; [`request_cookies_to_add`](#add-818453) - Optional Block<br>Add Cookies in Cookie Header. Cookies are key-value pairs to be added to HTTP request being routed towards upstream. Cookies specified at this level are applied after cookies from matched Route are applied<br>See [Request Cookies To Add](#add-818453) below.

<a id="remove-8d3e2e"></a>&#x2022; [`request_cookies_to_remove`](#remove-8d3e2e) - Optional List<br>Remove Cookies from Cookie Header. List of keys of Cookies to be removed from the HTTP request being sent towards upstream

<a id="add-b0a13c"></a>&#x2022; [`request_headers_to_add`](#add-b0a13c) - Optional Block<br>Add Request Headers. Headers are key-value pairs to be added to HTTP request being routed towards upstream<br>See [Request Headers To Add](#add-b0a13c) below.

<a id="remove-68d81f"></a>&#x2022; [`request_headers_to_remove`](#remove-68d81f) - Optional List<br>Remove Request Headers. List of keys of Headers to be removed from the HTTP request being sent towards upstream

<a id="add-b87fc7"></a>&#x2022; [`response_cookies_to_add`](#add-b87fc7) - Optional Block<br>Add Set-Cookie Headers. Cookies are name-value pairs along with optional attribute parameters to be added to HTTP response being sent towards downstream. Cookies specified at this level are applied after cookies from matched Route are applied<br>See [Response Cookies To Add](#add-b87fc7) below.

<a id="remove-e3e521"></a>&#x2022; [`response_cookies_to_remove`](#remove-e3e521) - Optional List<br>Remove Cookies from Set-Cookie Headers. List of name of Cookies to be removed from the HTTP response being sent towards downstream. Entire set-cookie header will be removed

<a id="add-9fbd2a"></a>&#x2022; [`response_headers_to_add`](#add-9fbd2a) - Optional Block<br>Add Response Headers. Headers are key-value pairs to be added to HTTP response being sent towards downstream<br>See [Response Headers To Add](#add-9fbd2a) below.

<a id="remove-5c99fb"></a>&#x2022; [`response_headers_to_remove`](#remove-5c99fb) - Optional List<br>Remove Response Headers. List of keys of Headers to be removed from the HTTP response being sent towards downstream

<a id="cluster-3cb556"></a>&#x2022; [`retract_cluster`](#cluster-3cb556) - Optional Block<br>Enable this option

<a id="policy-e40fa6"></a>&#x2022; [`retry_policy`](#policy-e40fa6) - Optional Block<br>Retry Policy. Retry policy configuration for route destination<br>See [Retry Policy](#policy-e40fa6) below.

<a id="policy-2b5b42"></a>&#x2022; [`specific_hash_policy`](#policy-2b5b42) - Optional Block<br>Hash Policy List. List of hash policy rules<br>See [Specific Hash Policy](#policy-2b5b42) below.

<a id="timeout-57487b"></a>&#x2022; [`timeout`](#timeout-57487b) - Optional Number<br>Timeout. The timeout for the route including all retries, in milliseconds. Should be set to a high value or 0 (infinite timeout) for server-side streaming

<a id="policy-4e9c59"></a>&#x2022; [`waf_exclusion_policy`](#policy-4e9c59) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [WAF Exclusion Policy](#policy-4e9c59) below.

<a id="config-8876be"></a>&#x2022; [`web_socket_config`](#config-8876be) - Optional Block<br>WebSocket Configuration. Configuration to allow WebSocket Request headers of such upgrade looks like below 'connection', 'Upgrade' 'upgrade', 'WebSocket' With configuration to allow WebSocket upgrade, ADC will produce following response 'HTTP/1.1 101 Switching Protocols 'Upgrade': 'WebSocket' 'Connection': 'Upgrade'<br>See [Web Socket Config](#config-8876be) below.

#### Routes Simple Route Advanced Options App Firewall

An [`app_firewall`](#firewall-b8f7c9) block (within [`routes.simple_route.advanced_options`](#routes-simple-route-advanced-options)) supports the following:

<a id="name-dd1197"></a>&#x2022; [`name`](#name-dd1197) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-118663"></a>&#x2022; [`namespace`](#namespace-118663) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-6ccea5"></a>&#x2022; [`tenant`](#tenant-6ccea5) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Routes Simple Route Advanced Options Bot Defense Javascript Injection

A [`bot_defense_javascript_injection`](#injection-23f0bd) block (within [`routes.simple_route.advanced_options`](#routes-simple-route-advanced-options)) supports the following:

<a id="location-c6e119"></a>&#x2022; [`javascript_location`](#location-c6e119) - Optional String  Defaults to `AFTER_HEAD`<br>Possible values are `AFTER_HEAD`, `AFTER_TITLE_END`, `BEFORE_SCRIPT`<br>[Enum: AFTER_HEAD|AFTER_TITLE_END|BEFORE_SCRIPT] JavaScript Location. All inside networks. Insert JavaScript after `<head>` tag Insert JavaScript after `</title>` tag. Insert JavaScript before first `<script>` tag

<a id="tags-ae7abb"></a>&#x2022; [`javascript_tags`](#tags-ae7abb) - Optional Block<br>JavaScript Tags. Select Add item to configure your javascript tag. If adding both Bot Adv and Fraud, the Bot Javascript should be added first<br>See [Javascript Tags](#tags-ae7abb) below.

#### Routes Simple Route Advanced Options Bot Defense Javascript Injection Javascript Tags

A [`javascript_tags`](#tags-ae7abb) block (within [`routes.simple_route.advanced_options.bot_defense_javascript_injection`](#injection-23f0bd)) supports the following:

<a id="url-7720a6"></a>&#x2022; [`javascript_url`](#url-7720a6) - Optional String<br>URL. Please enter the full URL (include domain and path), or relative path

<a id="attributes-9c576a"></a>&#x2022; [`tag_attributes`](#attributes-9c576a) - Optional Block<br>Tag Attributes. Add the tag attributes you want to include in your Javascript tag<br>See [Tag Attributes](#attributes-9c576a) below.

#### Routes Simple Route Advanced Options Bot Defense Javascript Injection Javascript Tags Tag Attributes

A [`tag_attributes`](#attributes-9c576a) block (within [`routes.simple_route.advanced_options.bot_defense_javascript_injection.javascript_tags`](#tags-ae7abb)) supports the following:

<a id="tag-3220a1"></a>&#x2022; [`javascript_tag`](#tag-3220a1) - Optional String  Defaults to `JS_ATTR_ID`<br>Possible values are `JS_ATTR_ID`, `JS_ATTR_CID`, `JS_ATTR_CN`, `JS_ATTR_API_DOMAIN`, `JS_ATTR_API_URL`, `JS_ATTR_API_PATH`, `JS_ATTR_ASYNC`, `JS_ATTR_DEFER`<br>[Enum: JS_ATTR_ID|JS_ATTR_CID|JS_ATTR_CN|JS_ATTR_API_DOMAIN|JS_ATTR_API_URL|JS_ATTR_API_PATH|JS_ATTR_ASYNC|JS_ATTR_DEFER] Tag Attribute Name. Select from one of the predefined tag attributes

<a id="value-1c74d0"></a>&#x2022; [`tag_value`](#value-1c74d0) - Optional String<br>Value. Add the tag attribute value

#### Routes Simple Route Advanced Options Buffer Policy

A [`buffer_policy`](#policy-23a3f6) block (within [`routes.simple_route.advanced_options`](#routes-simple-route-advanced-options)) supports the following:

<a id="disabled-8985b5"></a>&#x2022; [`disabled`](#disabled-8985b5) - Optional Bool<br>Disable. Disable buffering for a particular route. This is useful when virtual-host has buffering, but we need to disable it on a specific route. The value of this field is ignored for virtual-host

<a id="bytes-aad0b4"></a>&#x2022; [`max_request_bytes`](#bytes-aad0b4) - Optional Number<br>Max Request Bytes. The maximum request size that the filter will buffer before the connection manager will stop buffering and return a RequestEntityTooLarge (413) response

#### Routes Simple Route Advanced Options CORS Policy

A [`cors_policy`](#policy-ba853e) block (within [`routes.simple_route.advanced_options`](#routes-simple-route-advanced-options)) supports the following:

<a id="credentials-5c92f8"></a>&#x2022; [`allow_credentials`](#credentials-5c92f8) - Optional Bool<br>Allow Credentials. Specifies whether the resource allows credentials

<a id="headers-706468"></a>&#x2022; [`allow_headers`](#headers-706468) - Optional String<br>Allow Headers. Specifies the content for the access-control-allow-headers header

<a id="methods-8c987d"></a>&#x2022; [`allow_methods`](#methods-8c987d) - Optional String<br>Allow Methods. Specifies the content for the access-control-allow-methods header

<a id="origin-fbf2bc"></a>&#x2022; [`allow_origin`](#origin-fbf2bc) - Optional List<br>Allow Origin. Specifies the origins that will be allowed to do CORS requests. An origin is allowed if either allow_origin or allow_origin_regex match

<a id="regex-0641e8"></a>&#x2022; [`allow_origin_regex`](#regex-0641e8) - Optional List<br>Allow Origin Regex. Specifies regex patterns that match allowed origins. An origin is allowed if either allow_origin or allow_origin_regex match

<a id="disabled-cd4fb6"></a>&#x2022; [`disabled`](#disabled-cd4fb6) - Optional Bool<br>Disabled. Disable the CorsPolicy for a particular route. This is useful when virtual-host has CorsPolicy, but we need to disable it on a specific route. The value of this field is ignored for virtual-host

<a id="headers-c03d50"></a>&#x2022; [`expose_headers`](#headers-c03d50) - Optional String<br>Expose Headers. Specifies the content for the access-control-expose-headers header

<a id="age-eabbcc"></a>&#x2022; [`maximum_age`](#age-eabbcc) - Optional Number<br>Maximum Age. Specifies the content for the access-control-max-age header in seconds. This indicates the maximum number of seconds the results can be cached A value of -1 will disable caching. Maximum permitted value is 86400 seconds (24 hours)

#### Routes Simple Route Advanced Options CSRF Policy

A [`csrf_policy`](#policy-7816d7) block (within [`routes.simple_route.advanced_options`](#routes-simple-route-advanced-options)) supports the following:

<a id="domains-b58044"></a>&#x2022; [`all_load_balancer_domains`](#domains-b58044) - Optional Block<br>Enable this option

<a id="list-c11aec"></a>&#x2022; [`custom_domain_list`](#list-c11aec) - Optional Block<br>Domain name list. List of domain names used for Host header matching<br>See [Custom Domain List](#list-c11aec) below.

<a id="disabled-ac6077"></a>&#x2022; [`disabled`](#disabled-ac6077) - Optional Block<br>Enable this option

#### Routes Simple Route Advanced Options CSRF Policy Custom Domain List

A [`custom_domain_list`](#list-c11aec) block (within [`routes.simple_route.advanced_options.csrf_policy`](#policy-7816d7)) supports the following:

<a id="domains-17cc07"></a>&#x2022; [`domains`](#domains-17cc07) - Optional List<br>Domain names. A list of domain names that will be matched to loadbalancer. These domains are not used for SNI match. Wildcard names are supported in the suffix or prefix form

#### Routes Simple Route Advanced Options Mirror Policy

A [`mirror_policy`](#policy-f5e84d) block (within [`routes.simple_route.advanced_options`](#routes-simple-route-advanced-options)) supports the following:

<a id="pool-8c75a0"></a>&#x2022; [`origin_pool`](#pool-8c75a0) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Origin Pool](#pool-8c75a0) below.

<a id="percent-99590a"></a>&#x2022; [`percent`](#percent-99590a) - Optional Block<br>Fractional Percent. Fraction used where sampling percentages are needed. example sampled requests<br>See [Percent](#percent-99590a) below.

#### Routes Simple Route Advanced Options Mirror Policy Origin Pool

An [`origin_pool`](#pool-8c75a0) block (within [`routes.simple_route.advanced_options.mirror_policy`](#policy-f5e84d)) supports the following:

<a id="name-07b8c1"></a>&#x2022; [`name`](#name-07b8c1) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-69a88c"></a>&#x2022; [`namespace`](#namespace-69a88c) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-1c63d9"></a>&#x2022; [`tenant`](#tenant-1c63d9) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Routes Simple Route Advanced Options Mirror Policy Percent

A [`percent`](#percent-99590a) block (within [`routes.simple_route.advanced_options.mirror_policy`](#policy-f5e84d)) supports the following:

<a id="denominator-bc2f36"></a>&#x2022; [`denominator`](#denominator-bc2f36) - Optional String  Defaults to `HUNDRED`<br>Possible values are `HUNDRED`, `TEN_THOUSAND`, `MILLION`<br>[Enum: HUNDRED|TEN_THOUSAND|MILLION] Denominator. Denominator used in fraction where sampling percentages are needed. example sampled requests Use hundred as denominator Use ten thousand as denominator Use million as denominator

<a id="numerator-63dd3d"></a>&#x2022; [`numerator`](#numerator-63dd3d) - Optional Number<br>Numerator. sampled parts per denominator. If denominator was 10000, then value of 5 will be 5 in 10000

#### Routes Simple Route Advanced Options Regex Rewrite

A [`regex_rewrite`](#rewrite-c628a7) block (within [`routes.simple_route.advanced_options`](#routes-simple-route-advanced-options)) supports the following:

<a id="pattern-191576"></a>&#x2022; [`pattern`](#pattern-191576) - Optional String<br>Pattern. The regular expression used to find portions of a string that should be replaced

<a id="substitution-55c137"></a>&#x2022; [`substitution`](#substitution-55c137) - Optional String<br>Substitution. The string that should be substituted into matching portions of the subject string during a substitution operation to produce a new string

#### Routes Simple Route Advanced Options Request Cookies To Add

A [`request_cookies_to_add`](#add-818453) block (within [`routes.simple_route.advanced_options`](#routes-simple-route-advanced-options)) supports the following:

<a id="name-3366db"></a>&#x2022; [`name`](#name-3366db) - Optional String<br>Name. Name of the cookie in Cookie header

<a id="overwrite-9f0365"></a>&#x2022; [`overwrite`](#overwrite-9f0365) - Optional Bool  Defaults to `do`<br>Overwrite. Should the value be overwritten? If true, the value is overwritten to existing values. not overwrite

<a id="value-b6b480"></a>&#x2022; [`secret_value`](#value-b6b480) - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Secret Value](#value-b6b480) below.

<a id="value-c8a3cf"></a>&#x2022; [`value`](#value-c8a3cf) - Optional String<br>Value. Value of the Cookie header

#### Routes Simple Route Advanced Options Request Cookies To Add Secret Value

A [`secret_value`](#value-b6b480) block (within [`routes.simple_route.advanced_options.request_cookies_to_add`](#add-818453)) supports the following:

<a id="info-595c5e"></a>&#x2022; [`blindfold_secret_info`](#info-595c5e) - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#info-595c5e) below.

<a id="info-748acf"></a>&#x2022; [`clear_secret_info`](#info-748acf) - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#info-748acf) below.

#### Routes Simple Route Advanced Options Request Cookies To Add Secret Value Blindfold Secret Info

A [`blindfold_secret_info`](#info-595c5e) block (within [`routes.simple_route.advanced_options.request_cookies_to_add.secret_value`](#value-b6b480)) supports the following:

<a id="provider-8a22f4"></a>&#x2022; [`decryption_provider`](#provider-8a22f4) - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

<a id="location-1d3277"></a>&#x2022; [`location`](#location-1d3277) - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

<a id="provider-fd40ab"></a>&#x2022; [`store_provider`](#provider-fd40ab) - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

#### Routes Simple Route Advanced Options Request Cookies To Add Secret Value Clear Secret Info

A [`clear_secret_info`](#info-748acf) block (within [`routes.simple_route.advanced_options.request_cookies_to_add.secret_value`](#value-b6b480)) supports the following:

<a id="ref-b36cbb"></a>&#x2022; [`provider_ref`](#ref-b36cbb) - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

<a id="url-d3ffc3"></a>&#x2022; [`url`](#url-d3ffc3) - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

#### Routes Simple Route Advanced Options Request Headers To Add

A [`request_headers_to_add`](#add-b0a13c) block (within [`routes.simple_route.advanced_options`](#routes-simple-route-advanced-options)) supports the following:

<a id="append-cba980"></a>&#x2022; [`append`](#append-cba980) - Optional Bool  Defaults to `do`<br>Append. Should the value be appended? If true, the value is appended to existing values. not append

<a id="name-9fd523"></a>&#x2022; [`name`](#name-9fd523) - Optional String<br>Name. Name of the HTTP header

<a id="value-9a671c"></a>&#x2022; [`secret_value`](#value-9a671c) - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Secret Value](#value-9a671c) below.

<a id="value-e4a3b5"></a>&#x2022; [`value`](#value-e4a3b5) - Optional String<br>Value. Value of the HTTP header

#### Routes Simple Route Advanced Options Request Headers To Add Secret Value

A [`secret_value`](#value-9a671c) block (within [`routes.simple_route.advanced_options.request_headers_to_add`](#add-b0a13c)) supports the following:

<a id="info-4b823c"></a>&#x2022; [`blindfold_secret_info`](#info-4b823c) - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#info-4b823c) below.

<a id="info-d6bcd0"></a>&#x2022; [`clear_secret_info`](#info-d6bcd0) - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#info-d6bcd0) below.

#### Routes Simple Route Advanced Options Request Headers To Add Secret Value Blindfold Secret Info

A [`blindfold_secret_info`](#info-4b823c) block (within [`routes.simple_route.advanced_options.request_headers_to_add.secret_value`](#value-9a671c)) supports the following:

<a id="provider-5ade69"></a>&#x2022; [`decryption_provider`](#provider-5ade69) - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

<a id="location-1a5f3c"></a>&#x2022; [`location`](#location-1a5f3c) - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

<a id="provider-fb8118"></a>&#x2022; [`store_provider`](#provider-fb8118) - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

#### Routes Simple Route Advanced Options Request Headers To Add Secret Value Clear Secret Info

A [`clear_secret_info`](#info-d6bcd0) block (within [`routes.simple_route.advanced_options.request_headers_to_add.secret_value`](#value-9a671c)) supports the following:

<a id="ref-50f150"></a>&#x2022; [`provider_ref`](#ref-50f150) - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

<a id="url-facf7e"></a>&#x2022; [`url`](#url-facf7e) - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

#### Routes Simple Route Advanced Options Response Cookies To Add

A [`response_cookies_to_add`](#add-b87fc7) block (within [`routes.simple_route.advanced_options`](#routes-simple-route-advanced-options)) supports the following:

<a id="domain-a5e720"></a>&#x2022; [`add_domain`](#domain-a5e720) - Optional String<br>Add Domain. Add domain attribute

<a id="expiry-8a2f2f"></a>&#x2022; [`add_expiry`](#expiry-8a2f2f) - Optional String<br>Add expiry. Add expiry attribute

<a id="httponly-eb4086"></a>&#x2022; [`add_httponly`](#httponly-eb4086) - Optional Block<br>Enable this option

<a id="partitioned-fae05f"></a>&#x2022; [`add_partitioned`](#partitioned-fae05f) - Optional Block<br>Enable this option

<a id="path-598e89"></a>&#x2022; [`add_path`](#path-598e89) - Optional String<br>Add path. Add path attribute

<a id="secure-704f76"></a>&#x2022; [`add_secure`](#secure-704f76) - Optional Block<br>Enable this option

<a id="domain-06950b"></a>&#x2022; [`ignore_domain`](#domain-06950b) - Optional Block<br>Enable this option

<a id="expiry-452914"></a>&#x2022; [`ignore_expiry`](#expiry-452914) - Optional Block<br>Enable this option

<a id="httponly-565d2f"></a>&#x2022; [`ignore_httponly`](#httponly-565d2f) - Optional Block<br>Enable this option

<a id="age-0ac812"></a>&#x2022; [`ignore_max_age`](#age-0ac812) - Optional Block<br>Enable this option

<a id="partitioned-60bc5b"></a>&#x2022; [`ignore_partitioned`](#partitioned-60bc5b) - Optional Block<br>Enable this option

<a id="path-b87dc7"></a>&#x2022; [`ignore_path`](#path-b87dc7) - Optional Block<br>Enable this option

<a id="samesite-9e77d8"></a>&#x2022; [`ignore_samesite`](#samesite-9e77d8) - Optional Block<br>Enable this option

<a id="secure-02848e"></a>&#x2022; [`ignore_secure`](#secure-02848e) - Optional Block<br>Enable this option

<a id="value-f79f42"></a>&#x2022; [`ignore_value`](#value-f79f42) - Optional Block<br>Enable this option

<a id="value-8fcccf"></a>&#x2022; [`max_age_value`](#value-8fcccf) - Optional Number<br>Add Max Age. Add max age attribute

<a id="name-ca6654"></a>&#x2022; [`name`](#name-ca6654) - Optional String<br>Name. Name of the cookie in Cookie header

<a id="overwrite-e595bf"></a>&#x2022; [`overwrite`](#overwrite-e595bf) - Optional Bool  Defaults to `do`<br>Overwrite. Should the value be overwritten? If true, the value is overwritten to existing values. not overwrite

<a id="lax-a80ae4"></a>&#x2022; [`samesite_lax`](#lax-a80ae4) - Optional Block<br>Enable this option

<a id="none-6fd04e"></a>&#x2022; [`samesite_none`](#none-6fd04e) - Optional Block<br>Enable this option

<a id="strict-f8f2b1"></a>&#x2022; [`samesite_strict`](#strict-f8f2b1) - Optional Block<br>Enable this option

<a id="value-8f0f8f"></a>&#x2022; [`secret_value`](#value-8f0f8f) - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Secret Value](#value-8f0f8f) below.

<a id="value-b54fc1"></a>&#x2022; [`value`](#value-b54fc1) - Optional String<br>Value. Value of the Cookie header

#### Routes Simple Route Advanced Options Response Cookies To Add Secret Value

A [`secret_value`](#value-8f0f8f) block (within [`routes.simple_route.advanced_options.response_cookies_to_add`](#add-b87fc7)) supports the following:

<a id="info-b052ee"></a>&#x2022; [`blindfold_secret_info`](#info-b052ee) - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#info-b052ee) below.

<a id="info-9a64b1"></a>&#x2022; [`clear_secret_info`](#info-9a64b1) - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#info-9a64b1) below.

#### Routes Simple Route Advanced Options Response Cookies To Add Secret Value Blindfold Secret Info

A [`blindfold_secret_info`](#info-b052ee) block (within [`routes.simple_route.advanced_options.response_cookies_to_add.secret_value`](#value-8f0f8f)) supports the following:

<a id="provider-a0fc39"></a>&#x2022; [`decryption_provider`](#provider-a0fc39) - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

<a id="location-d79de4"></a>&#x2022; [`location`](#location-d79de4) - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

<a id="provider-94a6cd"></a>&#x2022; [`store_provider`](#provider-94a6cd) - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

#### Routes Simple Route Advanced Options Response Cookies To Add Secret Value Clear Secret Info

A [`clear_secret_info`](#info-9a64b1) block (within [`routes.simple_route.advanced_options.response_cookies_to_add.secret_value`](#value-8f0f8f)) supports the following:

<a id="ref-392858"></a>&#x2022; [`provider_ref`](#ref-392858) - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

<a id="url-a53238"></a>&#x2022; [`url`](#url-a53238) - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

#### Routes Simple Route Advanced Options Response Headers To Add

A [`response_headers_to_add`](#add-9fbd2a) block (within [`routes.simple_route.advanced_options`](#routes-simple-route-advanced-options)) supports the following:

<a id="append-83cfb2"></a>&#x2022; [`append`](#append-83cfb2) - Optional Bool  Defaults to `do`<br>Append. Should the value be appended? If true, the value is appended to existing values. not append

<a id="name-be5965"></a>&#x2022; [`name`](#name-be5965) - Optional String<br>Name. Name of the HTTP header

<a id="value-beb59b"></a>&#x2022; [`secret_value`](#value-beb59b) - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Secret Value](#value-beb59b) below.

<a id="value-05076d"></a>&#x2022; [`value`](#value-05076d) - Optional String<br>Value. Value of the HTTP header

#### Routes Simple Route Advanced Options Response Headers To Add Secret Value

A [`secret_value`](#value-beb59b) block (within [`routes.simple_route.advanced_options.response_headers_to_add`](#add-9fbd2a)) supports the following:

<a id="info-d130f8"></a>&#x2022; [`blindfold_secret_info`](#info-d130f8) - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#info-d130f8) below.

<a id="info-f99ba9"></a>&#x2022; [`clear_secret_info`](#info-f99ba9) - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#info-f99ba9) below.

#### Routes Simple Route Advanced Options Response Headers To Add Secret Value Blindfold Secret Info

A [`blindfold_secret_info`](#info-d130f8) block (within [`routes.simple_route.advanced_options.response_headers_to_add.secret_value`](#value-beb59b)) supports the following:

<a id="provider-7f573e"></a>&#x2022; [`decryption_provider`](#provider-7f573e) - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

<a id="location-779737"></a>&#x2022; [`location`](#location-779737) - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

<a id="provider-6d60f6"></a>&#x2022; [`store_provider`](#provider-6d60f6) - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

#### Routes Simple Route Advanced Options Response Headers To Add Secret Value Clear Secret Info

A [`clear_secret_info`](#info-f99ba9) block (within [`routes.simple_route.advanced_options.response_headers_to_add.secret_value`](#value-beb59b)) supports the following:

<a id="ref-c11d41"></a>&#x2022; [`provider_ref`](#ref-c11d41) - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

<a id="url-88123e"></a>&#x2022; [`url`](#url-88123e) - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

#### Routes Simple Route Advanced Options Retry Policy

A [`retry_policy`](#policy-e40fa6) block (within [`routes.simple_route.advanced_options`](#routes-simple-route-advanced-options)) supports the following:

<a id="off-e4369f"></a>&#x2022; [`back_off`](#off-e4369f) - Optional Block<br>Retry BackOff Interval. Specifies parameters that control retry back off<br>See [Back Off](#off-e4369f) below.

<a id="retries-ee7703"></a>&#x2022; [`num_retries`](#retries-ee7703) - Optional Number  Defaults to `1`<br>Number of Retries. Specifies the allowed number of retries. Retries can be done any number of times. An exponential back-off algorithm is used between each retry

<a id="timeout-2485fd"></a>&#x2022; [`per_try_timeout`](#timeout-2485fd) - Optional Number<br>Per Try Timeout. Specifies a non-zero timeout per retry attempt. In milliseconds

<a id="codes-110133"></a>&#x2022; [`retriable_status_codes`](#codes-110133) - Optional List<br>Status Code to Retry. HTTP status codes that should trigger a retry in addition to those specified by retry_on

<a id="condition-28432a"></a>&#x2022; [`retry_condition`](#condition-28432a) - Optional List<br>Retry Condition. Specifies the conditions under which retry takes place. Retries can be on different types of condition depending on application requirements. For example, network failure, all 5xx response codes, idempotent 4xx response codes, etc The possible values are '5xx' : Retry will be done if the upstream server responds with any 5xx response code, or does not respond at all (disconnect/reset/read timeout). 'gateway-error' : Retry will be done only if the upstream server responds with 502, 503 or 504 responses (Included in 5xx) 'connect-failure' : Retry will be done if the request fails because of a connection failure to the upstream server (connect timeout, etc.). (Included in 5xx) 'refused-stream' : Retry is done if the upstream server resets the stream with a REFUSED_STREAM error code (Included in 5xx) 'retriable-4xx' : Retry is done if the upstream server responds with a retriable 4xx response code. The only response code in this category is HTTP CONFLICT (409) 'retriable-status-codes' : Retry is done if the upstream server responds with any response code matching one defined in retriable_status_codes field 'reset' : Retry is done if the upstream server does not respond at all (disconnect/reset/read timeout.)

#### Routes Simple Route Advanced Options Retry Policy Back Off

A [`back_off`](#off-e4369f) block (within [`routes.simple_route.advanced_options.retry_policy`](#policy-e40fa6)) supports the following:

<a id="interval-eb135e"></a>&#x2022; [`base_interval`](#interval-eb135e) - Optional Number<br>Base Retry Interval. Specifies the base interval between retries in milliseconds

<a id="interval-f1eeb7"></a>&#x2022; [`max_interval`](#interval-f1eeb7) - Optional Number  Defaults to `10`<br>Maximum Retry Interval. Specifies the maximum interval between retries in milliseconds. This parameter is optional, but must be greater than or equal to the base_interval if set. The times the base_interval

#### Routes Simple Route Advanced Options Specific Hash Policy

A [`specific_hash_policy`](#policy-2b5b42) block (within [`routes.simple_route.advanced_options`](#routes-simple-route-advanced-options)) supports the following:

<a id="policy-3743d7"></a>&#x2022; [`hash_policy`](#policy-3743d7) - Optional Block<br>Hash Policy. Specifies a list of hash policies to use for ring hash load balancing. Each hash policy is evaluated individually and the combined result is used to route the request<br>See [Hash Policy](#policy-3743d7) below.

#### Routes Simple Route Advanced Options Specific Hash Policy Hash Policy

A [`hash_policy`](#policy-3743d7) block (within [`routes.simple_route.advanced_options.specific_hash_policy`](#policy-2b5b42)) supports the following:

<a id="cookie-c89776"></a>&#x2022; [`cookie`](#cookie-c89776) - Optional Block<br>Hashing using Cookie. Two types of cookie affinity: 1. Passive. Takes a cookie that's present in the cookies header and hashes on its value. 2. Generated. Generates and sets a cookie with an expiration (TTL) on the first request from the client in its response to the client, based on the endpoint the request gets sent to. The client then presents this on the next and all subsequent requests. The hash of this is sufficient to ensure these requests get sent to the same endpoint. The cookie is generated by hashing the source and destination ports and addresses so that multiple independent HTTP2 streams on the same connection will independently receive the same cookie, even if they arrive simultaneously<br>See [Cookie](#cookie-c89776) below.

<a id="name-d32f88"></a>&#x2022; [`header_name`](#name-d32f88) - Optional String<br>Header Name. The name or key of the request header that will be used to obtain the hash key

<a id="source-ip-fde636"></a>&#x2022; [`source_ip`](#source-ip-fde636) - Optional Bool<br>Source IP. Hash based on source IP address

<a id="terminal-63fd33"></a>&#x2022; [`terminal`](#terminal-63fd33) - Optional Bool<br>Terminal. Specify if its a terminal policy

#### Routes Simple Route Advanced Options Specific Hash Policy Hash Policy Cookie

A [`cookie`](#cookie-c89776) block (within [`routes.simple_route.advanced_options.specific_hash_policy.hash_policy`](#policy-3743d7)) supports the following:

<a id="httponly-c4f2e7"></a>&#x2022; [`add_httponly`](#httponly-c4f2e7) - Optional Block<br>Enable this option

<a id="secure-870d63"></a>&#x2022; [`add_secure`](#secure-870d63) - Optional Block<br>Enable this option

<a id="httponly-d97d73"></a>&#x2022; [`ignore_httponly`](#httponly-d97d73) - Optional Block<br>Enable this option

<a id="samesite-66f244"></a>&#x2022; [`ignore_samesite`](#samesite-66f244) - Optional Block<br>Enable this option

<a id="secure-6baa9f"></a>&#x2022; [`ignore_secure`](#secure-6baa9f) - Optional Block<br>Enable this option

<a id="name-d8ff3e"></a>&#x2022; [`name`](#name-d8ff3e) - Optional String<br>Name. The name of the cookie that will be used to obtain the hash key. If the cookie is not present and TTL below is not set, no hash will be produced

<a id="path-bade23"></a>&#x2022; [`path`](#path-bade23) - Optional String<br>Path. The name of the path for the cookie. If no path is specified here, no path will be set for the cookie

<a id="lax-938303"></a>&#x2022; [`samesite_lax`](#lax-938303) - Optional Block<br>Enable this option

<a id="none-caacf2"></a>&#x2022; [`samesite_none`](#none-caacf2) - Optional Block<br>Enable this option

<a id="strict-8b314e"></a>&#x2022; [`samesite_strict`](#strict-8b314e) - Optional Block<br>Enable this option

<a id="ttl-6e530c"></a>&#x2022; [`ttl`](#ttl-6e530c) - Optional Number<br>TTL. If specified, a cookie with the TTL will be generated if the cookie is not present. If the TTL is present and zero, the generated cookie will be a session cookie. TTL value is in milliseconds

#### Routes Simple Route Advanced Options WAF Exclusion Policy

A [`waf_exclusion_policy`](#policy-4e9c59) block (within [`routes.simple_route.advanced_options`](#routes-simple-route-advanced-options)) supports the following:

<a id="name-d7d4c4"></a>&#x2022; [`name`](#name-d7d4c4) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-7479f5"></a>&#x2022; [`namespace`](#namespace-7479f5) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-6c2697"></a>&#x2022; [`tenant`](#tenant-6c2697) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Routes Simple Route Advanced Options Web Socket Config

A [`web_socket_config`](#config-8876be) block (within [`routes.simple_route.advanced_options`](#routes-simple-route-advanced-options)) supports the following:

<a id="websocket-3b58f3"></a>&#x2022; [`use_websocket`](#websocket-3b58f3) - Optional Bool<br>Use WebSocket. Specifies that the HTTP client connection to this route is allowed to upgrade to a WebSocket connection

#### Routes Simple Route Headers

A [`headers`](#routes-simple-route-headers) block (within [`routes.simple_route`](#routes-simple-route)) supports the following:

<a id="routes-simple-route-headers-exact"></a>&#x2022; [`exact`](#routes-simple-route-headers-exact) - Optional String<br>Exact. Header value to match exactly

<a id="match-29c396"></a>&#x2022; [`invert_match`](#match-29c396) - Optional Bool<br>NOT of match. Invert the result of the match to detect missing header or non-matching value

<a id="routes-simple-route-headers-name"></a>&#x2022; [`name`](#routes-simple-route-headers-name) - Optional String<br>Name. Name of the header

<a id="routes-simple-route-headers-presence"></a>&#x2022; [`presence`](#routes-simple-route-headers-presence) - Optional Bool<br>Presence. If true, check for presence of header

<a id="routes-simple-route-headers-regex"></a>&#x2022; [`regex`](#routes-simple-route-headers-regex) - Optional String<br>Regex. Regex match of the header value in re2 format

#### Routes Simple Route Incoming Port

An [`incoming_port`](#routes-simple-route-incoming-port) block (within [`routes.simple_route`](#routes-simple-route)) supports the following:

<a id="match-80b1e5"></a>&#x2022; [`no_port_match`](#match-80b1e5) - Optional Block<br>Enable this option

<a id="routes-simple-route-incoming-port-port"></a>&#x2022; [`port`](#routes-simple-route-incoming-port-port) - Optional Number<br>Port. Exact Port to match

<a id="ranges-b32092"></a>&#x2022; [`port_ranges`](#ranges-b32092) - Optional String<br>Port range. Port range to match

#### Routes Simple Route Origin Pools

An [`origin_pools`](#routes-simple-route-origin-pools) block (within [`routes.simple_route`](#routes-simple-route)) supports the following:

<a id="cluster-094c5e"></a>&#x2022; [`cluster`](#cluster-094c5e) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Cluster](#cluster-094c5e) below.

<a id="subsets-73b90d"></a>&#x2022; [`endpoint_subsets`](#subsets-73b90d) - Optional Block<br>Origin Servers Subsets. Upstream origin pool may be configured to divide its origin servers into subsets based on metadata attached to the origin servers. Routes may then specify the metadata that a endpoint must match in order to be selected by the load balancer For origin servers which are discovered in K8S or Consul cluster, the label of the service is merged with endpoint's labels. In case of Consul, the label is derived from the 'Tag' field. For labels that are common between configured endpoint and discovered service, labels from discovered service takes precedence. List of key-value pairs that will be used as matching metadata. Only those origin servers of upstream origin pool which match this metadata will be selected for load balancing

<a id="routes-simple-route-origin-pools-pool"></a>&#x2022; [`pool`](#routes-simple-route-origin-pools-pool) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Pool](#routes-simple-route-origin-pools-pool) below.

<a id="priority-b8cd45"></a>&#x2022; [`priority`](#priority-b8cd45) - Optional Number<br>Priority. Priority of this origin pool, valid only with multiple origin pools. Value of 0 will make the pool as lowest priority origin pool Priority of 1 means highest priority and is considered active. When active origin pool is not available, lower priority origin pools are made active as per the increasing priority

<a id="routes-simple-route-origin-pools-weight"></a>&#x2022; [`weight`](#routes-simple-route-origin-pools-weight) - Optional Number<br>Weight. Weight of this origin pool, valid only with multiple origin pool. Value of 0 will disable the pool

#### Routes Simple Route Origin Pools Cluster

A [`cluster`](#cluster-094c5e) block (within [`routes.simple_route.origin_pools`](#routes-simple-route-origin-pools)) supports the following:

<a id="name-82cffb"></a>&#x2022; [`name`](#name-82cffb) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-b73e1d"></a>&#x2022; [`namespace`](#namespace-b73e1d) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-511970"></a>&#x2022; [`tenant`](#tenant-511970) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Routes Simple Route Origin Pools Pool

A [`pool`](#routes-simple-route-origin-pools-pool) block (within [`routes.simple_route.origin_pools`](#routes-simple-route-origin-pools)) supports the following:

<a id="name-7e0a9d"></a>&#x2022; [`name`](#name-7e0a9d) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-f2152d"></a>&#x2022; [`namespace`](#namespace-f2152d) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-06f782"></a>&#x2022; [`tenant`](#tenant-06f782) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Routes Simple Route Path

A [`path`](#routes-simple-route-path) block (within [`routes.simple_route`](#routes-simple-route)) supports the following:

<a id="routes-simple-route-path-path"></a>&#x2022; [`path`](#routes-simple-route-path-path) - Optional String<br>Exact. Exact path value to match

<a id="routes-simple-route-path-prefix"></a>&#x2022; [`prefix`](#routes-simple-route-path-prefix) - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

<a id="routes-simple-route-path-regex"></a>&#x2022; [`regex`](#routes-simple-route-path-regex) - Optional String<br>Regex. Regular expression of path match (e.g. the value .* will match on all paths)

#### Routes Simple Route Query Params

A [`query_params`](#routes-simple-route-query-params) block (within [`routes.simple_route`](#routes-simple-route)) supports the following:

<a id="params-aa1f17"></a>&#x2022; [`remove_all_params`](#params-aa1f17) - Optional Block<br>Enable this option

<a id="params-c3e5f1"></a>&#x2022; [`replace_params`](#params-c3e5f1) - Optional String<br>Replace All Parameters

<a id="params-bd2237"></a>&#x2022; [`retain_all_params`](#params-bd2237) - Optional Block<br>Enable this option

#### Sensitive Data Disclosure Rules

A [`sensitive_data_disclosure_rules`](#sensitive-data-disclosure-rules) block supports the following:

<a id="response-2680e4"></a>&#x2022; [`sensitive_data_types_in_response`](#response-2680e4) - Optional Block<br>Sensitive Data Exposure Rules. Sensitive Data Exposure Rules allows specifying rules to mask sensitive data fields in API responses<br>See [Sensitive Data Types In Response](#response-2680e4) below.

#### Sensitive Data Disclosure Rules Sensitive Data Types In Response

A [`sensitive_data_types_in_response`](#response-2680e4) block (within [`sensitive_data_disclosure_rules`](#sensitive-data-disclosure-rules)) supports the following:

<a id="endpoint-0fdc68"></a>&#x2022; [`api_endpoint`](#endpoint-0fdc68) - Optional Block<br>API Endpoint. This defines API endpoint<br>See [API Endpoint](#endpoint-0fdc68) below.

<a id="body-4757be"></a>&#x2022; [`body`](#body-4757be) - Optional Block<br>Body Section Masking Options. Options for HTTP Body Masking<br>See [Body](#body-4757be) below.

<a id="mask-87d5f1"></a>&#x2022; [`mask`](#mask-87d5f1) - Optional Block<br>Enable this option

<a id="report-71fe99"></a>&#x2022; [`report`](#report-71fe99) - Optional Block<br>Enable this option

#### Sensitive Data Disclosure Rules Sensitive Data Types In Response API Endpoint

An [`api_endpoint`](#endpoint-0fdc68) block (within [`sensitive_data_disclosure_rules.sensitive_data_types_in_response`](#response-2680e4)) supports the following:

<a id="methods-a6eecd"></a>&#x2022; [`methods`](#methods-a6eecd) - Optional List  Defaults to `ANY`<br>Possible values are `ANY`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`, `COPY`<br>[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY] Methods. Methods to be matched

<a id="path-8e5892"></a>&#x2022; [`path`](#path-8e5892) - Optional String<br>Path. Path to be matched

#### Sensitive Data Disclosure Rules Sensitive Data Types In Response Body

A [`body`](#body-4757be) block (within [`sensitive_data_disclosure_rules.sensitive_data_types_in_response`](#response-2680e4)) supports the following:

<a id="fields-bb6c3f"></a>&#x2022; [`fields`](#fields-bb6c3f) - Optional List<br>Values. List of JSON Path field values. Use square brackets with an underscore [\_] to indicate array elements (e.g., person.emails[\_]). To reference JSON keys that contain spaces, enclose the entire path in double quotes. For example: 'person.first name'

#### Sensitive Data Policy

A [`sensitive_data_policy`](#sensitive-data-policy) block supports the following:

<a id="ref-55b260"></a>&#x2022; [`sensitive_data_policy_ref`](#ref-55b260) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Sensitive Data Policy Ref](#ref-55b260) below.

#### Sensitive Data Policy Sensitive Data Policy Ref

A [`sensitive_data_policy_ref`](#ref-55b260) block (within [`sensitive_data_policy`](#sensitive-data-policy)) supports the following:

<a id="name-d254a7"></a>&#x2022; [`name`](#name-d254a7) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-401387"></a>&#x2022; [`namespace`](#namespace-401387) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-d10cc7"></a>&#x2022; [`tenant`](#tenant-d10cc7) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Single LB App

A [`single_lb_app`](#single-lb-app) block supports the following:

<a id="single-lb-app-disable-discovery"></a>&#x2022; [`disable_discovery`](#single-lb-app-disable-discovery) - Optional Block<br>Enable this option

<a id="detection-d482d0"></a>&#x2022; [`disable_malicious_user_detection`](#detection-d482d0) - Optional Block<br>Enable this option

<a id="single-lb-app-enable-discovery"></a>&#x2022; [`enable_discovery`](#single-lb-app-enable-discovery) - Optional Block<br>API Discovery Setting. Specifies the settings used for API discovery<br>See [Enable Discovery](#single-lb-app-enable-discovery) below.

<a id="detection-280554"></a>&#x2022; [`enable_malicious_user_detection`](#detection-280554) - Optional Block<br>Enable this option

#### Single LB App Enable Discovery

An [`enable_discovery`](#single-lb-app-enable-discovery) block (within [`single_lb_app`](#single-lb-app)) supports the following:

<a id="crawler-cb748a"></a>&#x2022; [`api_crawler`](#crawler-cb748a) - Optional Block<br>API Crawling. API Crawler message<br>See [API Crawler](#crawler-cb748a) below.

<a id="scan-c86c24"></a>&#x2022; [`api_discovery_from_code_scan`](#scan-c86c24) - Optional Block<br>Select Code Base and Repositories<br>See [API Discovery From Code Scan](#scan-c86c24) below.

<a id="discovery-3993cd"></a>&#x2022; [`custom_api_auth_discovery`](#discovery-3993cd) - Optional Block<br>API Discovery Advanced Settings. API Discovery Advanced settings<br>See [Custom API Auth Discovery](#discovery-3993cd) below.

<a id="discovery-e02c6d"></a>&#x2022; [`default_api_auth_discovery`](#discovery-e02c6d) - Optional Block<br>Enable this option

<a id="traffic-7e1631"></a>&#x2022; [`disable_learn_from_redirect_traffic`](#traffic-7e1631) - Optional Block<br>Enable this option

<a id="settings-e36cc7"></a>&#x2022; [`discovered_api_settings`](#settings-e36cc7) - Optional Block<br>Discovered API Settings. Configure Discovered API Settings<br>See [Discovered API Settings](#settings-e36cc7) below.

<a id="traffic-ebfb24"></a>&#x2022; [`enable_learn_from_redirect_traffic`](#traffic-ebfb24) - Optional Block<br>Enable this option

#### Single LB App Enable Discovery API Crawler

An [`api_crawler`](#crawler-cb748a) block (within [`single_lb_app.enable_discovery`](#single-lb-app-enable-discovery)) supports the following:

<a id="config-3110d5"></a>&#x2022; [`api_crawler_config`](#config-3110d5) - Optional Block<br>Crawler Configure<br>See [API Crawler Config](#config-3110d5) below.

<a id="crawler-f73dfe"></a>&#x2022; [`disable_api_crawler`](#crawler-f73dfe) - Optional Block<br>Enable this option

#### Single LB App Enable Discovery API Crawler API Crawler Config

An [`api_crawler_config`](#config-3110d5) block (within [`single_lb_app.enable_discovery.api_crawler`](#crawler-cb748a)) supports the following:

<a id="domains-fafd03"></a>&#x2022; [`domains`](#domains-fafd03) - Optional Block<br>Domains to Crawl. Enter domains and their credentials to allow authenticated API crawling. You can only include domains you own that are associated with this Load Balancer<br>See [Domains](#domains-fafd03) below.

#### Single LB App Enable Discovery API Crawler API Crawler Config Domains

A [`domains`](#domains-fafd03) block (within [`single_lb_app.enable_discovery.api_crawler.api_crawler_config`](#config-3110d5)) supports the following:

<a id="domain-d4b18c"></a>&#x2022; [`domain`](#domain-d4b18c) - Optional String<br>Domain. Select the domain to execute API Crawling with given credentials

<a id="login-a68d0b"></a>&#x2022; [`simple_login`](#login-a68d0b) - Optional Block<br>Simple Login<br>See [Simple Login](#login-a68d0b) below.

#### Single LB App Enable Discovery API Crawler API Crawler Config Domains Simple Login

A [`simple_login`](#login-a68d0b) block (within [`single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains`](#domains-fafd03)) supports the following:

<a id="password-a28dbf"></a>&#x2022; [`password`](#password-a28dbf) - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Password](#password-a28dbf) below.

<a id="user-f4e1c2"></a>&#x2022; [`user`](#user-f4e1c2) - Optional String<br>User. Enter the username to assign credentials for the selected domain to crawl

#### Single LB App Enable Discovery API Crawler API Crawler Config Domains Simple Login Password

A [`password`](#password-a28dbf) block (within [`single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login`](#login-a68d0b)) supports the following:

<a id="info-a7bf15"></a>&#x2022; [`blindfold_secret_info`](#info-a7bf15) - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#info-a7bf15) below.

<a id="info-b4d182"></a>&#x2022; [`clear_secret_info`](#info-b4d182) - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#info-b4d182) below.

#### Single LB App Enable Discovery API Crawler API Crawler Config Domains Simple Login Password Blindfold Secret Info

A [`blindfold_secret_info`](#info-a7bf15) block (within [`single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password`](#password-a28dbf)) supports the following:

<a id="provider-3a3dc7"></a>&#x2022; [`decryption_provider`](#provider-3a3dc7) - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

<a id="location-c9e17c"></a>&#x2022; [`location`](#location-c9e17c) - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

<a id="provider-d6bf00"></a>&#x2022; [`store_provider`](#provider-d6bf00) - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

#### Single LB App Enable Discovery API Crawler API Crawler Config Domains Simple Login Password Clear Secret Info

A [`clear_secret_info`](#info-b4d182) block (within [`single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password`](#password-a28dbf)) supports the following:

<a id="ref-28caaf"></a>&#x2022; [`provider_ref`](#ref-28caaf) - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

<a id="url-9f4d43"></a>&#x2022; [`url`](#url-9f4d43) - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

#### Single LB App Enable Discovery API Discovery From Code Scan

An [`api_discovery_from_code_scan`](#scan-c86c24) block (within [`single_lb_app.enable_discovery`](#single-lb-app-enable-discovery)) supports the following:

<a id="integrations-0d9a39"></a>&#x2022; [`code_base_integrations`](#integrations-0d9a39) - Optional Block<br>Select Code Base Integrations<br>See [Code Base Integrations](#integrations-0d9a39) below.

#### Single LB App Enable Discovery API Discovery From Code Scan Code Base Integrations

A [`code_base_integrations`](#integrations-0d9a39) block (within [`single_lb_app.enable_discovery.api_discovery_from_code_scan`](#scan-c86c24)) supports the following:

<a id="repos-e3fdd1"></a>&#x2022; [`all_repos`](#repos-e3fdd1) - Optional Block<br>Enable this option

<a id="integration-7aa3ed"></a>&#x2022; [`code_base_integration`](#integration-7aa3ed) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Code Base Integration](#integration-7aa3ed) below.

<a id="repos-6aacd7"></a>&#x2022; [`selected_repos`](#repos-6aacd7) - Optional Block<br>API Code Repositories. Select which API repositories represent the LB applications<br>See [Selected Repos](#repos-6aacd7) below.

#### Single LB App Enable Discovery API Discovery From Code Scan Code Base Integrations Code Base Integration

A [`code_base_integration`](#integration-7aa3ed) block (within [`single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations`](#integrations-0d9a39)) supports the following:

<a id="name-25d13c"></a>&#x2022; [`name`](#name-25d13c) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-e1fd2c"></a>&#x2022; [`namespace`](#namespace-e1fd2c) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-a678e2"></a>&#x2022; [`tenant`](#tenant-a678e2) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Single LB App Enable Discovery API Discovery From Code Scan Code Base Integrations Selected Repos

A [`selected_repos`](#repos-6aacd7) block (within [`single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations`](#integrations-0d9a39)) supports the following:

<a id="repo-b055b9"></a>&#x2022; [`api_code_repo`](#repo-b055b9) - Optional List<br>API Code Repository. Code repository which contain API endpoints

#### Single LB App Enable Discovery Custom API Auth Discovery

A [`custom_api_auth_discovery`](#discovery-3993cd) block (within [`single_lb_app.enable_discovery`](#single-lb-app-enable-discovery)) supports the following:

<a id="ref-9b1b8a"></a>&#x2022; [`api_discovery_ref`](#ref-9b1b8a) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [API Discovery Ref](#ref-9b1b8a) below.

#### Single LB App Enable Discovery Custom API Auth Discovery API Discovery Ref

An [`api_discovery_ref`](#ref-9b1b8a) block (within [`single_lb_app.enable_discovery.custom_api_auth_discovery`](#discovery-3993cd)) supports the following:

<a id="name-ff2ba8"></a>&#x2022; [`name`](#name-ff2ba8) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-01d21f"></a>&#x2022; [`namespace`](#namespace-01d21f) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-5faa9d"></a>&#x2022; [`tenant`](#tenant-5faa9d) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Single LB App Enable Discovery Discovered API Settings

A [`discovered_api_settings`](#settings-e36cc7) block (within [`single_lb_app.enable_discovery`](#single-lb-app-enable-discovery)) supports the following:

<a id="apis-7c9721"></a>&#x2022; [`purge_duration_for_inactive_discovered_apis`](#apis-7c9721) - Optional Number<br>Purge Duration for Inactive Discovered APIs from Traffic. Inactive discovered API will be deleted after configured duration

#### Slow DDOS Mitigation

A [`slow_ddos_mitigation`](#slow-ddos-mitigation) block supports the following:

<a id="timeout-81071e"></a>&#x2022; [`disable_request_timeout`](#timeout-81071e) - Optional Block<br>Enable this option

<a id="timeout-da89d3"></a>&#x2022; [`request_headers_timeout`](#timeout-da89d3) - Optional Number  Defaults to `10000`<br>Request Headers Timeout. The amount of time the client has to send only the headers on the request stream before the stream is cancelled. The milliseconds. This setting provides protection against Slowloris attacks

<a id="slow-ddos-mitigation-request-timeout"></a>&#x2022; [`request_timeout`](#slow-ddos-mitigation-request-timeout) - Optional Number<br>Custom Timeout

#### Timeouts

A [`timeouts`](#timeouts) block supports the following:

<a id="timeouts-create"></a>&#x2022; [`create`](#timeouts-create) - Optional String (Defaults to `10 minutes`)<br>Used when creating the resource

<a id="timeouts-delete"></a>&#x2022; [`delete`](#timeouts-delete) - Optional String (Defaults to `10 minutes`)<br>Used when deleting the resource

<a id="timeouts-read"></a>&#x2022; [`read`](#timeouts-read) - Optional String (Defaults to `5 minutes`)<br>Used when retrieving the resource

<a id="timeouts-update"></a>&#x2022; [`update`](#timeouts-update) - Optional String (Defaults to `10 minutes`)<br>Used when updating the resource

#### Trusted Clients

A [`trusted_clients`](#trusted-clients) block supports the following:

<a id="trusted-clients-actions"></a>&#x2022; [`actions`](#trusted-clients-actions) - Optional List  Defaults to `SKIP_PROCESSING_WAF`<br>Possible values are `SKIP_PROCESSING_WAF`, `SKIP_PROCESSING_BOT`, `SKIP_PROCESSING_MUM`, `SKIP_PROCESSING_IP_REPUTATION`, `SKIP_PROCESSING_API_PROTECTION`, `SKIP_PROCESSING_OAS_VALIDATION`, `SKIP_PROCESSING_DDOS_PROTECTION`, `SKIP_PROCESSING_THREAT_MESH`, `SKIP_PROCESSING_MALWARE_PROTECTION`<br>[Enum: SKIP_PROCESSING_WAF|SKIP_PROCESSING_BOT|SKIP_PROCESSING_MUM|SKIP_PROCESSING_IP_REPUTATION|SKIP_PROCESSING_API_PROTECTION|SKIP_PROCESSING_OAS_VALIDATION|SKIP_PROCESSING_DDOS_PROTECTION|SKIP_PROCESSING_THREAT_MESH|SKIP_PROCESSING_MALWARE_PROTECTION] Actions. Actions that should be taken when client identifier matches the rule

<a id="trusted-clients-as-number"></a>&#x2022; [`as_number`](#trusted-clients-as-number) - Optional Number<br>AS Number. RFC 6793 defined 4-byte AS number

<a id="trusted-clients-bot-skip-processing"></a>&#x2022; [`bot_skip_processing`](#trusted-clients-bot-skip-processing) - Optional Block<br>Enable this option

<a id="trusted-clients-expiration-timestamp"></a>&#x2022; [`expiration_timestamp`](#trusted-clients-expiration-timestamp) - Optional String<br>Expiration Timestamp. The expiration_timestamp is the RFC 3339 format timestamp at which the containing rule is considered to be logically expired. The rule continues to exist in the configuration but is not applied anymore

<a id="trusted-clients-http-header"></a>&#x2022; [`http_header`](#trusted-clients-http-header) - Optional Block<br>HTTP Header. Request header name and value pairs<br>See [HTTP Header](#trusted-clients-http-header) below.

<a id="trusted-clients-ip-prefix"></a>&#x2022; [`ip_prefix`](#trusted-clients-ip-prefix) - Optional String<br>IPv4 Prefix. IPv4 prefix string

<a id="trusted-clients-ipv6-prefix"></a>&#x2022; [`ipv6_prefix`](#trusted-clients-ipv6-prefix) - Optional String<br>IPv6 Prefix. IPv6 prefix string

<a id="trusted-clients-metadata"></a>&#x2022; [`metadata`](#trusted-clients-metadata) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#trusted-clients-metadata) below.

<a id="trusted-clients-skip-processing"></a>&#x2022; [`skip_processing`](#trusted-clients-skip-processing) - Optional Block<br>Enable this option

<a id="trusted-clients-user-identifier"></a>&#x2022; [`user_identifier`](#trusted-clients-user-identifier) - Optional String<br>User Identifier. Identify user based on user identifier. User identifier value needs to be copied from security event

<a id="trusted-clients-waf-skip-processing"></a>&#x2022; [`waf_skip_processing`](#trusted-clients-waf-skip-processing) - Optional Block<br>Enable this option

#### Trusted Clients HTTP Header

A [`http_header`](#trusted-clients-http-header) block (within [`trusted_clients`](#trusted-clients)) supports the following:

<a id="trusted-clients-http-header-headers"></a>&#x2022; [`headers`](#trusted-clients-http-header-headers) - Optional Block<br>HTTP Headers. List of HTTP header name and value pairs<br>See [Headers](#trusted-clients-http-header-headers) below.

#### Trusted Clients HTTP Header Headers

A [`headers`](#trusted-clients-http-header-headers) block (within [`trusted_clients.http_header`](#trusted-clients-http-header)) supports the following:

<a id="exact-1a048f"></a>&#x2022; [`exact`](#exact-1a048f) - Optional String<br>Exact. Header value to match exactly

<a id="match-4d5659"></a>&#x2022; [`invert_match`](#match-4d5659) - Optional Bool<br>NOT of match. Invert the result of the match to detect missing header or non-matching value

<a id="name-b3a383"></a>&#x2022; [`name`](#name-b3a383) - Optional String<br>Name. Name of the header

<a id="presence-a73dd8"></a>&#x2022; [`presence`](#presence-a73dd8) - Optional Bool<br>Presence. If true, check for presence of header

<a id="regex-d6b675"></a>&#x2022; [`regex`](#regex-d6b675) - Optional String<br>Regex. Regex match of the header value in re2 format

#### Trusted Clients Metadata

A [`metadata`](#trusted-clients-metadata) block (within [`trusted_clients`](#trusted-clients)) supports the following:

<a id="spec-766a6d"></a>&#x2022; [`description_spec`](#spec-766a6d) - Optional String<br>Description. Human readable description

<a id="trusted-clients-metadata-name"></a>&#x2022; [`name`](#trusted-clients-metadata-name) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### User Identification

An [`user_identification`](#user-identification) block supports the following:

<a id="user-identification-name"></a>&#x2022; [`name`](#user-identification-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="user-identification-namespace"></a>&#x2022; [`namespace`](#user-identification-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="user-identification-tenant"></a>&#x2022; [`tenant`](#user-identification-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### WAF Exclusion

A [`waf_exclusion`](#waf-exclusion) block supports the following:

<a id="rules-6d8efc"></a>&#x2022; [`waf_exclusion_inline_rules`](#rules-6d8efc) - Optional Block<br>WAF Exclusion Inline Rules. A list of WAF exclusion rules that will be applied inline<br>See [WAF Exclusion Inline Rules](#rules-6d8efc) below.

<a id="waf-exclusion-waf-exclusion-policy"></a>&#x2022; [`waf_exclusion_policy`](#waf-exclusion-waf-exclusion-policy) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [WAF Exclusion Policy](#waf-exclusion-waf-exclusion-policy) below.

#### WAF Exclusion WAF Exclusion Inline Rules

A [`waf_exclusion_inline_rules`](#rules-6d8efc) block (within [`waf_exclusion`](#waf-exclusion)) supports the following:

<a id="rules-28cf34"></a>&#x2022; [`rules`](#rules-28cf34) - Optional Block<br>WAF Exclusion Rules. An ordered list of WAF Exclusions specific to this Load Balancer<br>See [Rules](#rules-28cf34) below.

#### WAF Exclusion WAF Exclusion Inline Rules Rules

A [`rules`](#rules-28cf34) block (within [`waf_exclusion.waf_exclusion_inline_rules`](#rules-6d8efc)) supports the following:

<a id="domain-3f85e2"></a>&#x2022; [`any_domain`](#domain-3f85e2) - Optional Block<br>Enable this option

<a id="path-f75bfe"></a>&#x2022; [`any_path`](#path-f75bfe) - Optional Block<br>Enable this option

<a id="control-0cb52d"></a>&#x2022; [`app_firewall_detection_control`](#control-0cb52d) - Optional Block<br>App Firewall Detection Control. Define the list of Signature IDs, Violations, Attack Types and Bot Names that should be excluded from triggering on the defined match criteria<br>See [App Firewall Detection Control](#control-0cb52d) below.

<a id="value-451fbf"></a>&#x2022; [`exact_value`](#value-451fbf) - Optional String<br>Exact Value. Exact domain name

<a id="timestamp-423d81"></a>&#x2022; [`expiration_timestamp`](#timestamp-423d81) - Optional String<br>Expiration Timestamp. The expiration_timestamp is the RFC 3339 format timestamp at which the containing rule is considered to be logically expired. The rule continues to exist in the configuration but is not applied anymore

<a id="metadata-09584f"></a>&#x2022; [`metadata`](#metadata-09584f) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#metadata-09584f) below.

<a id="methods-19f73d"></a>&#x2022; [`methods`](#methods-19f73d) - Optional List  Defaults to `ANY`<br>Possible values are `ANY`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`, `COPY`<br>[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY] Methods. methods to be matched

<a id="prefix-a857dd"></a>&#x2022; [`path_prefix`](#prefix-a857dd) - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

<a id="regex-fdbacd"></a>&#x2022; [`path_regex`](#regex-fdbacd) - Optional String<br>Path Regex. Define the regex for the path. For example, the regex ^/.*$ will match on all paths

<a id="value-6f2f58"></a>&#x2022; [`suffix_value`](#value-6f2f58) - Optional String<br>Suffix Value. Suffix of domain name e.g 'xyz.com' will match '*.xyz.com' and 'xyz.com'

<a id="processing-8c8391"></a>&#x2022; [`waf_skip_processing`](#processing-8c8391) - Optional Block<br>Enable this option

#### WAF Exclusion WAF Exclusion Inline Rules Rules App Firewall Detection Control

An [`app_firewall_detection_control`](#control-0cb52d) block (within [`waf_exclusion.waf_exclusion_inline_rules.rules`](#rules-28cf34)) supports the following:

<a id="contexts-6197b1"></a>&#x2022; [`exclude_attack_type_contexts`](#contexts-6197b1) - Optional Block<br>Attack Types. Attack Types to be excluded for the defined match criteria<br>See [Exclude Attack Type Contexts](#contexts-6197b1) below.

<a id="contexts-e832e4"></a>&#x2022; [`exclude_bot_name_contexts`](#contexts-e832e4) - Optional Block<br>Bot Names. Bot Names to be excluded for the defined match criteria<br>See [Exclude Bot Name Contexts](#contexts-e832e4) below.

<a id="contexts-0794ff"></a>&#x2022; [`exclude_signature_contexts`](#contexts-0794ff) - Optional Block<br>Signature IDs. Signature IDs to be excluded for the defined match criteria<br>See [Exclude Signature Contexts](#contexts-0794ff) below.

<a id="contexts-29dd68"></a>&#x2022; [`exclude_violation_contexts`](#contexts-29dd68) - Optional Block<br>Violations. Violations to be excluded for the defined match criteria<br>See [Exclude Violation Contexts](#contexts-29dd68) below.

#### WAF Exclusion WAF Exclusion Inline Rules Rules App Firewall Detection Control Exclude Attack Type Contexts

An [`exclude_attack_type_contexts`](#contexts-6197b1) block (within [`waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control`](#control-0cb52d)) supports the following:

<a id="context-b0f79f"></a>&#x2022; [`context`](#context-b0f79f) - Optional String  Defaults to `CONTEXT_ANY`<br>Possible values are `CONTEXT_ANY`, `CONTEXT_BODY`, `CONTEXT_REQUEST`, `CONTEXT_RESPONSE`, `CONTEXT_PARAMETER`, `CONTEXT_HEADER`, `CONTEXT_COOKIE`, `CONTEXT_URL`, `CONTEXT_URI`<br>[Enum: CONTEXT_ANY|CONTEXT_BODY|CONTEXT_REQUEST|CONTEXT_RESPONSE|CONTEXT_PARAMETER|CONTEXT_HEADER|CONTEXT_COOKIE|CONTEXT_URL|CONTEXT_URI] WAF Exclusion Context Options. The available contexts for Exclusion rules. - CONTEXT_ANY: CONTEXT_ANY Detection will be excluded for all contexts. - CONTEXT_BODY: CONTEXT_BODY Detection will be excluded for the request body. - CONTEXT_REQUEST: CONTEXT_REQUEST Detection will be excluded for the request. - CONTEXT_RESPONSE: CONTEXT_RESPONSE - CONTEXT_PARAMETER: CONTEXT_PARAMETER Detection will be excluded for the parameters. The parameter name is required in the Context name field. If the field is left empty, the detection will be excluded for all parameters. - CONTEXT_HEADER: CONTEXT_HEADER Detection will be excluded for the headers. The header name is required in the Context name field. If the field is left empty, the detection will be excluded for all headers. - CONTEXT_COOKIE: CONTEXT_COOKIE Detection will be excluded for the cookies. The cookie name is required in the Context name field. If the field is left empty, the detection will be excluded for all cookies. - CONTEXT_URL: CONTEXT_URL Detection will be excluded for the request URL. - CONTEXT_URI: CONTEXT_URI

<a id="name-114c4e"></a>&#x2022; [`context_name`](#name-114c4e) - Optional String<br>Context Name. Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an wildcard asterisk (*)

<a id="type-541ccc"></a>&#x2022; [`exclude_attack_type`](#type-541ccc) - Optional String  Defaults to `ATTACK_TYPE_NONE`<br>Possible values are `ATTACK_TYPE_NONE`, `ATTACK_TYPE_NON_BROWSER_CLIENT`, `ATTACK_TYPE_OTHER_APPLICATION_ATTACKS`, `ATTACK_TYPE_TROJAN_BACKDOOR_SPYWARE`, `ATTACK_TYPE_DETECTION_EVASION`, `ATTACK_TYPE_VULNERABILITY_SCAN`, `ATTACK_TYPE_ABUSE_OF_FUNCTIONALITY`, `ATTACK_TYPE_AUTHENTICATION_AUTHORIZATION_ATTACKS`, `ATTACK_TYPE_BUFFER_OVERFLOW`, `ATTACK_TYPE_PREDICTABLE_RESOURCE_LOCATION`, `ATTACK_TYPE_INFORMATION_LEAKAGE`, `ATTACK_TYPE_DIRECTORY_INDEXING`, `ATTACK_TYPE_PATH_TRAVERSAL`, `ATTACK_TYPE_XPATH_INJECTION`, `ATTACK_TYPE_LDAP_INJECTION`, `ATTACK_TYPE_SERVER_SIDE_CODE_INJECTION`, `ATTACK_TYPE_COMMAND_EXECUTION`, `ATTACK_TYPE_SQL_INJECTION`, `ATTACK_TYPE_CROSS_SITE_SCRIPTING`, `ATTACK_TYPE_DENIAL_OF_SERVICE`, `ATTACK_TYPE_HTTP_PARSER_ATTACK`, `ATTACK_TYPE_SESSION_HIJACKING`, `ATTACK_TYPE_HTTP_RESPONSE_SPLITTING`, `ATTACK_TYPE_FORCEFUL_BROWSING`, `ATTACK_TYPE_REMOTE_FILE_INCLUDE`, `ATTACK_TYPE_MALICIOUS_FILE_UPLOAD`, `ATTACK_TYPE_GRAPHQL_PARSER_ATTACK`<br>[Enum: ATTACK_TYPE_NONE|ATTACK_TYPE_NON_BROWSER_CLIENT|ATTACK_TYPE_OTHER_APPLICATION_ATTACKS|ATTACK_TYPE_TROJAN_BACKDOOR_SPYWARE|ATTACK_TYPE_DETECTION_EVASION|ATTACK_TYPE_VULNERABILITY_SCAN|ATTACK_TYPE_ABUSE_OF_FUNCTIONALITY|ATTACK_TYPE_AUTHENTICATION_AUTHORIZATION_ATTACKS|ATTACK_TYPE_BUFFER_OVERFLOW|ATTACK_TYPE_PREDICTABLE_RESOURCE_LOCATION|ATTACK_TYPE_INFORMATION_LEAKAGE|ATTACK_TYPE_DIRECTORY_INDEXING|ATTACK_TYPE_PATH_TRAVERSAL|ATTACK_TYPE_XPATH_INJECTION|ATTACK_TYPE_LDAP_INJECTION|ATTACK_TYPE_SERVER_SIDE_CODE_INJECTION|ATTACK_TYPE_COMMAND_EXECUTION|ATTACK_TYPE_SQL_INJECTION|ATTACK_TYPE_CROSS_SITE_SCRIPTING|ATTACK_TYPE_DENIAL_OF_SERVICE|ATTACK_TYPE_HTTP_PARSER_ATTACK|ATTACK_TYPE_SESSION_HIJACKING|ATTACK_TYPE_HTTP_RESPONSE_SPLITTING|ATTACK_TYPE_FORCEFUL_BROWSING|ATTACK_TYPE_REMOTE_FILE_INCLUDE|ATTACK_TYPE_MALICIOUS_FILE_UPLOAD|ATTACK_TYPE_GRAPHQL_PARSER_ATTACK] Attack Types. List of all Attack Types ATTACK_TYPE_NONE ATTACK_TYPE_NON_BROWSER_CLIENT ATTACK_TYPE_OTHER_APPLICATION_ATTACKS ATTACK_TYPE_TROJAN_BACKDOOR_SPYWARE ATTACK_TYPE_DETECTION_EVASION ATTACK_TYPE_VULNERABILITY_SCAN ATTACK_TYPE_ABUSE_OF_FUNCTIONALITY ATTACK_TYPE_AUTHENTICATION_AUTHORIZATION_ATTACKS ATTACK_TYPE_BUFFER_OVERFLOW ATTACK_TYPE_PREDICTABLE_RESOURCE_LOCATION ATTACK_TYPE_INFORMATION_LEAKAGE ATTACK_TYPE_DIRECTORY_INDEXING ATTACK_TYPE_PATH_TRAVERSAL ATTACK_TYPE_XPATH_INJECTION ATTACK_TYPE_LDAP_INJECTION ATTACK_TYPE_SERVER_SIDE_CODE_INJECTION ATTACK_TYPE_COMMAND_EXECUTION ATTACK_TYPE_SQL_INJECTION ATTACK_TYPE_CROSS_SITE_SCRIPTING ATTACK_TYPE_DENIAL_OF_SERVICE ATTACK_TYPE_HTTP_PARSER_ATTACK ATTACK_TYPE_SESSION_HIJACKING ATTACK_TYPE_HTTP_RESPONSE_SPLITTING ATTACK_TYPE_FORCEFUL_BROWSING ATTACK_TYPE_REMOTE_FILE_INCLUDE ATTACK_TYPE_MALICIOUS_FILE_UPLOAD ATTACK_TYPE_GRAPHQL_PARSER_ATTACK

#### WAF Exclusion WAF Exclusion Inline Rules Rules App Firewall Detection Control Exclude Bot Name Contexts

An [`exclude_bot_name_contexts`](#contexts-e832e4) block (within [`waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control`](#control-0cb52d)) supports the following:

<a id="name-1d3dba"></a>&#x2022; [`bot_name`](#name-1d3dba) - Optional String<br>Bot Name

#### WAF Exclusion WAF Exclusion Inline Rules Rules App Firewall Detection Control Exclude Signature Contexts

An [`exclude_signature_contexts`](#contexts-0794ff) block (within [`waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control`](#control-0cb52d)) supports the following:

<a id="context-e1f5a0"></a>&#x2022; [`context`](#context-e1f5a0) - Optional String  Defaults to `CONTEXT_ANY`<br>Possible values are `CONTEXT_ANY`, `CONTEXT_BODY`, `CONTEXT_REQUEST`, `CONTEXT_RESPONSE`, `CONTEXT_PARAMETER`, `CONTEXT_HEADER`, `CONTEXT_COOKIE`, `CONTEXT_URL`, `CONTEXT_URI`<br>[Enum: CONTEXT_ANY|CONTEXT_BODY|CONTEXT_REQUEST|CONTEXT_RESPONSE|CONTEXT_PARAMETER|CONTEXT_HEADER|CONTEXT_COOKIE|CONTEXT_URL|CONTEXT_URI] WAF Exclusion Context Options. The available contexts for Exclusion rules. - CONTEXT_ANY: CONTEXT_ANY Detection will be excluded for all contexts. - CONTEXT_BODY: CONTEXT_BODY Detection will be excluded for the request body. - CONTEXT_REQUEST: CONTEXT_REQUEST Detection will be excluded for the request. - CONTEXT_RESPONSE: CONTEXT_RESPONSE - CONTEXT_PARAMETER: CONTEXT_PARAMETER Detection will be excluded for the parameters. The parameter name is required in the Context name field. If the field is left empty, the detection will be excluded for all parameters. - CONTEXT_HEADER: CONTEXT_HEADER Detection will be excluded for the headers. The header name is required in the Context name field. If the field is left empty, the detection will be excluded for all headers. - CONTEXT_COOKIE: CONTEXT_COOKIE Detection will be excluded for the cookies. The cookie name is required in the Context name field. If the field is left empty, the detection will be excluded for all cookies. - CONTEXT_URL: CONTEXT_URL Detection will be excluded for the request URL. - CONTEXT_URI: CONTEXT_URI

<a id="name-efd12c"></a>&#x2022; [`context_name`](#name-efd12c) - Optional String<br>Context Name. Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an wildcard asterisk (*)

<a id="signature-id-f725d3"></a>&#x2022; [`signature_id`](#signature-id-f725d3) - Optional Number<br>SignatureID. The allowed values for signature ID are 0 and in the range of 200000001-299999999. 0 implies that all signatures will be excluded for the specified context

#### WAF Exclusion WAF Exclusion Inline Rules Rules App Firewall Detection Control Exclude Violation Contexts

An [`exclude_violation_contexts`](#contexts-29dd68) block (within [`waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control`](#control-0cb52d)) supports the following:

<a id="context-5543b9"></a>&#x2022; [`context`](#context-5543b9) - Optional String  Defaults to `CONTEXT_ANY`<br>Possible values are `CONTEXT_ANY`, `CONTEXT_BODY`, `CONTEXT_REQUEST`, `CONTEXT_RESPONSE`, `CONTEXT_PARAMETER`, `CONTEXT_HEADER`, `CONTEXT_COOKIE`, `CONTEXT_URL`, `CONTEXT_URI`<br>[Enum: CONTEXT_ANY|CONTEXT_BODY|CONTEXT_REQUEST|CONTEXT_RESPONSE|CONTEXT_PARAMETER|CONTEXT_HEADER|CONTEXT_COOKIE|CONTEXT_URL|CONTEXT_URI] WAF Exclusion Context Options. The available contexts for Exclusion rules. - CONTEXT_ANY: CONTEXT_ANY Detection will be excluded for all contexts. - CONTEXT_BODY: CONTEXT_BODY Detection will be excluded for the request body. - CONTEXT_REQUEST: CONTEXT_REQUEST Detection will be excluded for the request. - CONTEXT_RESPONSE: CONTEXT_RESPONSE - CONTEXT_PARAMETER: CONTEXT_PARAMETER Detection will be excluded for the parameters. The parameter name is required in the Context name field. If the field is left empty, the detection will be excluded for all parameters. - CONTEXT_HEADER: CONTEXT_HEADER Detection will be excluded for the headers. The header name is required in the Context name field. If the field is left empty, the detection will be excluded for all headers. - CONTEXT_COOKIE: CONTEXT_COOKIE Detection will be excluded for the cookies. The cookie name is required in the Context name field. If the field is left empty, the detection will be excluded for all cookies. - CONTEXT_URL: CONTEXT_URL Detection will be excluded for the request URL. - CONTEXT_URI: CONTEXT_URI

<a id="name-b96b20"></a>&#x2022; [`context_name`](#name-b96b20) - Optional String<br>Context Name. Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an wildcard asterisk (*)

<a id="violation-53384e"></a>&#x2022; [`exclude_violation`](#violation-53384e) - Optional String  Defaults to `VIOL_NONE`<br>Possible values are `VIOL_NONE`, `VIOL_FILETYPE`, `VIOL_METHOD`, `VIOL_MANDATORY_HEADER`, `VIOL_HTTP_RESPONSE_STATUS`, `VIOL_REQUEST_MAX_LENGTH`, `VIOL_FILE_UPLOAD`, `VIOL_FILE_UPLOAD_IN_BODY`, `VIOL_XML_MALFORMED`, `VIOL_JSON_MALFORMED`, `VIOL_ASM_COOKIE_MODIFIED`, `VIOL_HTTP_PROTOCOL_MULTIPLE_HOST_HEADERS`, `VIOL_HTTP_PROTOCOL_BAD_HOST_HEADER_VALUE`, `VIOL_HTTP_PROTOCOL_UNPARSABLE_REQUEST_CONTENT`, `VIOL_HTTP_PROTOCOL_NULL_IN_REQUEST`, `VIOL_HTTP_PROTOCOL_BAD_HTTP_VERSION`, `VIOL_HTTP_PROTOCOL_SEVERAL_CONTENT_LENGTH_HEADERS`, `VIOL_EVASION_DIRECTORY_TRAVERSALS`, `VIOL_MALFORMED_REQUEST`, `VIOL_EVASION_MULTIPLE_DECODING`, `VIOL_DATA_GUARD`, `VIOL_EVASION_APACHE_WHITESPACE`, `VIOL_COOKIE_MODIFIED`, `VIOL_EVASION_IIS_UNICODE_CODEPOINTS`, `VIOL_EVASION_IIS_BACKSLASHES`, `VIOL_EVASION_PERCENT_U_DECODING`, `VIOL_EVASION_BARE_BYTE_DECODING`, `VIOL_EVASION_BAD_UNESCAPE`, `VIOL_HTTP_PROTOCOL_BODY_IN_GET_OR_HEAD_REQUEST`, `VIOL_ENCODING`, `VIOL_COOKIE_MALFORMED`, `VIOL_GRAPHQL_FORMAT`, `VIOL_GRAPHQL_MALFORMED`, `VIOL_GRAPHQL_INTROSPECTION_QUERY`<br>[Enum: VIOL_NONE|VIOL_FILETYPE|VIOL_METHOD|VIOL_MANDATORY_HEADER|VIOL_HTTP_RESPONSE_STATUS|VIOL_REQUEST_MAX_LENGTH|VIOL_FILE_UPLOAD|VIOL_FILE_UPLOAD_IN_BODY|VIOL_XML_MALFORMED|VIOL_JSON_MALFORMED|VIOL_ASM_COOKIE_MODIFIED|VIOL_HTTP_PROTOCOL_MULTIPLE_HOST_HEADERS|VIOL_HTTP_PROTOCOL_BAD_HOST_HEADER_VALUE|VIOL_HTTP_PROTOCOL_UNPARSABLE_REQUEST_CONTENT|VIOL_HTTP_PROTOCOL_NULL_IN_REQUEST|VIOL_HTTP_PROTOCOL_BAD_HTTP_VERSION|VIOL_HTTP_PROTOCOL_SEVERAL_CONTENT_LENGTH_HEADERS|VIOL_EVASION_DIRECTORY_TRAVERSALS|VIOL_MALFORMED_REQUEST|VIOL_EVASION_MULTIPLE_DECODING|VIOL_DATA_GUARD|VIOL_EVASION_APACHE_WHITESPACE|VIOL_COOKIE_MODIFIED|VIOL_EVASION_IIS_UNICODE_CODEPOINTS|VIOL_EVASION_IIS_BACKSLASHES|VIOL_EVASION_PERCENT_U_DECODING|VIOL_EVASION_BARE_BYTE_DECODING|VIOL_EVASION_BAD_UNESCAPE|VIOL_HTTP_PROTOCOL_BODY_IN_GET_OR_HEAD_REQUEST|VIOL_ENCODING|VIOL_COOKIE_MALFORMED|VIOL_GRAPHQL_FORMAT|VIOL_GRAPHQL_MALFORMED|VIOL_GRAPHQL_INTROSPECTION_QUERY] App Firewall Violation Type. List of all supported Violation Types VIOL_NONE VIOL_FILETYPE VIOL_METHOD VIOL_MANDATORY_HEADER VIOL_HTTP_RESPONSE_STATUS VIOL_REQUEST_MAX_LENGTH VIOL_FILE_UPLOAD VIOL_FILE_UPLOAD_IN_BODY VIOL_XML_MALFORMED VIOL_JSON_MALFORMED VIOL_ASM_COOKIE_MODIFIED VIOL_HTTP_PROTOCOL_MULTIPLE_HOST_HEADERS VIOL_HTTP_PROTOCOL_BAD_HOST_HEADER_VALUE VIOL_HTTP_PROTOCOL_UNPARSABLE_REQUEST_CONTENT VIOL_HTTP_PROTOCOL_NULL_IN_REQUEST VIOL_HTTP_PROTOCOL_BAD_HTTP_VERSION VIOL_HTTP_PROTOCOL_CRLF_CHARACTERS_BEFORE_REQUEST_START VIOL_HTTP_PROTOCOL_NO_HOST_HEADER_IN_HTTP_1_1_REQUEST VIOL_HTTP_PROTOCOL_BAD_MULTIPART_PARAMETERS_PARSING VIOL_HTTP_PROTOCOL_SEVERAL_CONTENT_LENGTH_HEADERS VIOL_HTTP_PROTOCOL_CONTENT_LENGTH_SHOULD_BE_A_POSITIVE_NUMBER VIOL_EVASION_DIRECTORY_TRAVERSALS VIOL_MALFORMED_REQUEST VIOL_EVASION_MULTIPLE_DECODING VIOL_DATA_GUARD VIOL_EVASION_APACHE_WHITESPACE VIOL_COOKIE_MODIFIED VIOL_EVASION_IIS_UNICODE_CODEPOINTS VIOL_EVASION_IIS_BACKSLASHES VIOL_EVASION_PERCENT_U_DECODING VIOL_EVASION_BARE_BYTE_DECODING VIOL_EVASION_BAD_UNESCAPE VIOL_HTTP_PROTOCOL_BAD_MULTIPART_FORMDATA_REQUEST_PARSING VIOL_HTTP_PROTOCOL_BODY_IN_GET_OR_HEAD_REQUEST VIOL_HTTP_PROTOCOL_HIGH_ASCII_CHARACTERS_IN_HEADERS VIOL_ENCODING VIOL_COOKIE_MALFORMED VIOL_GRAPHQL_FORMAT VIOL_GRAPHQL_MALFORMED VIOL_GRAPHQL_INTROSPECTION_QUERY

#### WAF Exclusion WAF Exclusion Inline Rules Rules Metadata

A [`metadata`](#metadata-09584f) block (within [`waf_exclusion.waf_exclusion_inline_rules.rules`](#rules-28cf34)) supports the following:

<a id="spec-942e33"></a>&#x2022; [`description_spec`](#spec-942e33) - Optional String<br>Description. Human readable description

<a id="name-a43295"></a>&#x2022; [`name`](#name-a43295) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### WAF Exclusion WAF Exclusion Policy

A [`waf_exclusion_policy`](#waf-exclusion-waf-exclusion-policy) block (within [`waf_exclusion`](#waf-exclusion)) supports the following:

<a id="waf-exclusion-waf-exclusion-policy-name"></a>&#x2022; [`name`](#waf-exclusion-waf-exclusion-policy-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="namespace-d8f030"></a>&#x2022; [`namespace`](#namespace-d8f030) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="tenant-d841f0"></a>&#x2022; [`tenant`](#tenant-d841f0) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

## Import

Import is supported using the following syntax:

```shell
# Import using namespace/name format
terraform import f5xc_http_loadbalancer.example system/example
```
