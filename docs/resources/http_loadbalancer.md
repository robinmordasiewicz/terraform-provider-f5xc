---
page_title: "f5xc_http_loadbalancer Resource - terraform-provider-f5xc"
subcategory: "Load Balancing"
description: |-
  Manages a HTTP Load Balancer resource in F5 Distributed Cloud for load balancing HTTP/HTTPS traffic with advanced routing and security.
---

# f5xc_http_loadbalancer (Resource)

Manages a HTTP Load Balancer resource in F5 Distributed Cloud for load balancing HTTP/HTTPS traffic with advanced routing and security.

~> **Note** For more information about this resource, please refer to the [F5 XC API Documentation](https://docs.cloud.f5.com/docs/api/).

## Example Usage

```terraform
# HTTP Loadbalancer Resource Example
# Manages a HTTP Load Balancer resource in F5 Distributed Cloud for load balancing HTTP/HTTPS traffic with advanced routing and security.

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

<a id="add-location"></a>&#x2022; [`add_location`](#add-location) - Optional Bool<br>Add Location. X-example: true Appends header x-F5 Distributed Cloud-location = `<RE-site-name>` in responses. This configuration is ignored on CE sites

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
&#x2022; <a id="bot-defense"></a>[`bot_defense`](#bot-defense) - Optional Block<br>Bot Defense. This defines various configuration OPTIONS for Bot Defense Policy
<br><br>&#x2022; <a id="bot-defense-advanced"></a>[`bot_defense_advanced`](#bot-defense-advanced) - Optional Block<br>Bot Defense Advanced. Bot Defense Advanced

-> **One of the following:**
&#x2022; <a id="caching-policy"></a>[`caching_policy`](#caching-policy) - Optional Block<br>Caching Policies. Caching Policies for the CDN

-> **One of the following:**
&#x2022; <a id="captcha-challenge"></a>[`captcha_challenge`](#captcha-challenge) - Optional Block<br>Captcha Challenge Parameters. Enables loadbalancer to perform captcha challenge Captcha challenge will be based on Google Recaptcha. With this feature enabled, only clients that pass the captcha challenge will be allowed to complete the HTTP request. When loadbalancer is configured to do Captcha Challenge, it will redirect the browser to an HTML page on every new HTTP request. This HTML page will have captcha challenge embedded in it. Client will be allowed to make the request only if the captcha challenge is successful. Loadbalancer will tag response header with a cookie to avoid Captcha challenge for subsequent requests. CAPTCHA is mainly used as a security check to ensure only human users can pass through. Generally, computers or bots are not capable of solving a captcha. You can enable either Javascript challenge or Captcha challenge on a virtual host
<br><br>&#x2022; <a id="enable-challenge"></a>[`enable_challenge`](#enable-challenge) - Optional Block<br>Enable Malicious User Challenge. Configure auto mitigation i.e risk based challenges for malicious users
<br><br>&#x2022; <a id="js-challenge"></a>[`js_challenge`](#js-challenge) - Optional Block<br>Javascript Challenge Parameters. Enables loadbalancer to perform client browser compatibility test by redirecting to a page with Javascript. With this feature enabled, only clients that are capable of executing Javascript(mostly browsers) will be allowed to complete the HTTP request. When loadbalancer is configured to do Javascript Challenge, it will redirect the browser to an HTML page on every new HTTP request. This HTML page will have Javascript embedded in it. Loadbalancer chooses a set of random numbers for every new client and sends these numbers along with an encrypted answer with the request such that it embed these numbers as input in the Javascript. Javascript will run on the requestor browser and perform a complex Math operation. Script will submit the answer to loadbalancer. Loadbalancer will validate the answer by comparing the calculated answer with the decrypted answer (which was encrypted when it was sent back as reply) and allow the request to the upstream server only if the answer is correct. Loadbalancer will tag response header with a cookie to avoid Javascript challenge for subsequent requests. Javascript challenge serves following purposes \* Validate that the request is coming via a browser that is capable for running Javascript \* Force the browser to run a complex operation, f(X), that requires it to spend a large number of CPU cycles. This is to slow down a potential DOS attacker by making it difficult to launch a large request flood without having to spend even larger CPU cost at their end. You can enable either Javascript challenge or Captcha challenge on a virtual host
<br><br>&#x2022; <a id="no-challenge"></a>[`no_challenge`](#no-challenge) - Optional Block<br>Enable this option

-> **One of the following:**
&#x2022; <a id="client-side-defense"></a>[`client_side_defense`](#client-side-defense) - Optional Block<br>Client-Side Defense. This defines various configuration OPTIONS for Client-Side Defense Policy

-> **One of the following:**
&#x2022; <a id="cookie-stickiness"></a>[`cookie_stickiness`](#cookie-stickiness) - Optional Block<br>Hashing using Cookie. Two types of cookie affinity: 1. Passive. Takes a cookie that's present in the cookies header and hashes on its value. 2. Generated. Generates and sets a cookie with an expiration (TTL) on the first request from the client in its response to the client, based on the endpoint the request gets sent to. The client then presents this on the next and all subsequent requests. The hash of this is sufficient to ensure these requests GET sent to the same endpoint. The cookie is generated by hashing the source and destination ports and addresses so that multiple independent HTTP2 streams on the same connection will independently receive the same cookie, even if they arrive simultaneously
<br><br>&#x2022; <a id="least-active"></a>[`least_active`](#least-active) - Optional Block<br>Enable this option
<br><br>&#x2022; <a id="random"></a>[`random`](#random) - Optional Block<br>Enable this option
<br><br>&#x2022; <a id="ring-hash"></a>[`ring_hash`](#ring-hash) - Optional Block<br>Hash Policy List. List of hash policy rules
<br><br>&#x2022; <a id="round-robin"></a>[`round_robin`](#round-robin) - Optional Block<br>Enable this option

<a id="cors-policy"></a>&#x2022; [`cors_policy`](#cors-policy) - Optional Block<br>CORS Policy. Cross-Origin Resource Sharing requests configuration specified at Virtual-host or Route level. Route level configuration takes precedence. An example of an Cross origin HTTP request GET /resources/public-data/ HTTP/1.1 Host: bar.other User-Agent: Mozilla/5.0 (Macintosh; U; Intel MAC OS X 10.5; en-US; rv:1.9.1b3pre) Gecko/20081130 Minefield/3.1b3pre Accept: text/HTML,application/xhtml+XML,application/XML;q=0.9,*/*;q=0.8 Accept-Language: en-us,en;q=0.5 Accept-Encoding: gzip,deflate Accept-Charset: ISO-8859-1,utf-8;q=0.7,*;q=0.7 Connection: keep-alive Referrer: `HTTP://foo.example/examples/access-control/simplexsinvocation.HTML` Origin: `HTTP://foo.example` HTTP/1.1 200 OK Date: Mon, 01 Dec 2008 00:23:53 GMT Server: Apache/2.0.61 Access-Control-Allow-Origin: \* Keep-Alive: timeout=2, max=100 Connection: Keep-Alive Transfer-Encoding: chunked Content-Type: application/XML An example for cross origin HTTP OPTIONS request with Access-Control-Request-* header OPTIONS /resources/POST-here/ HTTP/1.1 Host: bar.other User-Agent: Mozilla/5.0 (Macintosh; U; Intel MAC OS X 10.5; en-US; rv:1.9.1b3pre) Gecko/20081130 Minefield/3.1b3pre Accept: text/HTML,application/xhtml+XML,application/XML;q=0.9,*/*;q=0.8 Accept-Language: en-us,en;q=0.5 Accept-Encoding: gzip,deflate Accept-Charset: ISO-8859-1,utf-8;q=0.7,*;q=0.7 Connection: keep-alive Origin: `HTTP://foo.example` Access-Control-Request-Method: POST Access-Control-Request-Headers: X-PINGOTHER, Content-Type HTTP/1.1 204 No Content Date: Mon, 01 Dec 2008 01:15:39 GMT Server: Apache/2.0.61 (Unix) Access-Control-Allow-Origin: `HTTP://foo.example` Access-Control-Allow-Methods: POST, GET, OPTIONS Access-Control-Allow-Headers: X-PINGOTHER, Content-Type Access-Control-Max-Age: 86400 Vary: Accept-Encoding, Origin Keep-Alive: timeout=2, max=100 Connection: Keep-Alive

<a id="csrf-policy"></a>&#x2022; [`csrf_policy`](#csrf-policy) - Optional Block<br>CSRF Policy. To mitigate CSRF attack , the policy checks where a request is coming from to determine if the request's origin is the same as its detination.the policy relies on two pieces of information used in determining if a request originated from the same host. 1. The origin that caused the user agent to issue the request (source origin). 2. The origin that the request is going to (target origin). When the policy evaluating a request, it ensures both pieces of information are present and compare their values. If the source origin is missing or origins do not match the request is rejected. The exception to this being if the source-origin has been added to they policy as valid. Because CSRF attacks specifically target state-changing requests, the policy only acts on the HTTP requests that have state-changing method (PUT,POST, etc.)

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

<a id="domains"></a>&#x2022; [`domains`](#domains) - Optional List<br>Domains. A list of Domains (host/authority header) that will be matched to load balancer. Supported Domains and search order: 1. Exact Domain names: `www.example.com.` 2. Domains starting with a Wildcard: \*.example.com. Not supported Domains: - Just a Wildcard: \* - A Wildcard and TLD with no root Domain: \*.com. - A Wildcard not matching a whole DNS label. E.g. \*.example.com and \*.bar.example.com are valid Wildcards however \*bar.example.com, \*-bar.example.com, and bar*.example.com are all invalid. Additional notes: A Wildcard will not match empty string. E.g. \*.example.com will match bar.example.com and baz-bar.example.com but not .example.com. The longest Wildcards match first. Only a single virtual host in the entire route configuration can match on \*. Also a Domain must be unique across all virtual hosts within an advertise policy. Domains are also used for SNI matching if the Loadbalancer type is HTTPS. Domains also indicate the list of names for which DNS resolution will be automatically resolved to IP addresses by the system

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

<a id="more-option"></a>&#x2022; [`more_option`](#more-option) - Optional Block<br>Advanced OPTIONS. This defines various OPTIONS to define a route

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

<a id="trusted-clients"></a>&#x2022; [`trusted_clients`](#trusted-clients) - Optional Block<br>Trusted Client Rules. Define rules to skip processing of one or more features such as WAF, Bot Defense etc. For clients

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

<a id="active-service-policies-policies"></a>&#x2022; [`policies`](#active-service-policies-policies) - Optional Block<br>Policies. Service Policies is a sequential engine where policies (and rules within the policy) are evaluated one after the other. It's important to define the correct order (policies evaluated from top to bottom in the list) for service policies, to GET the intended result. For each request, its characteristics are evaluated based on the match criteria in each service policy starting at the top. If there is a match in the current policy, then the policy takes effect, and no more policies are evaluated. Otherwise, the next policy is evaluated. If all policies are evaluated and none match, then the request will be denied by default<br>See [Policies](#active-service-policies-policies) below.

#### Active Service Policies Policies

A [`policies`](#active-service-policies-policies) block (within [`active_service_policies`](#active-service-policies)) supports the following:

<a id="active-service-policies-policies-name"></a>&#x2022; [`name`](#active-service-policies-policies-name) - Optional String<br>Name. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name

<a id="namespace-df0e5f"></a>&#x2022; [`namespace`](#namespace-df0e5f) - Optional String<br>Namespace. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace

<a id="active-service-policies-policies-tenant"></a>&#x2022; [`tenant`](#active-service-policies-policies-tenant) - Optional String<br>Tenant. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant

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

<a id="service-1fdc7a"></a>&#x2022; [`vk8s_service`](#service-1fdc7a) - Optional Block<br>VK8s Services on RE. This defines a reference to a RE site or virtual site where a load balancer could be advertised in the vK8s service network<br>See [Vk8s Service](#service-1fdc7a) below.

#### Advertise Custom Advertise Where Advertise On Public

An [`advertise_on_public`](#public-618a99) block (within [`advertise_custom.advertise_where`](#advertise-custom-advertise-where)) supports the following:

<a id="public-ip-d10b09"></a>&#x2022; [`public_ip`](#public-ip-d10b09) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Public IP](#public-ip-d10b09) below.

#### Advertise Custom Advertise Where Advertise On Public Public IP

<a id="deep-032ffb"></a>Deeply nested **IP** block collapsed for readability.

#### Advertise Custom Advertise Where Site

A [`site`](#advertise-custom-advertise-where-site) block (within [`advertise_custom.advertise_where`](#advertise-custom-advertise-where)) supports the following:

<a id="site-ip-78faa1"></a>&#x2022; [`ip`](#site-ip-78faa1) - Optional String<br>IP Address. Use given IP address as VIP on the site

<a id="network-5811a4"></a>&#x2022; [`network`](#network-5811a4) - Optional String  Defaults to `SITE_NETWORK_INSIDE_AND_OUTSIDE`<br>Possible values are `SITE_NETWORK_INSIDE_AND_OUTSIDE`, `SITE_NETWORK_INSIDE`, `SITE_NETWORK_OUTSIDE`, `SITE_NETWORK_SERVICE`, `SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP`, `SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP`, `SITE_NETWORK_IP_FABRIC`<br>[Enum: SITE_NETWORK_INSIDE_AND_OUTSIDE|SITE_NETWORK_INSIDE|SITE_NETWORK_OUTSIDE|SITE_NETWORK_SERVICE|SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP|SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP|SITE_NETWORK_IP_FABRIC] Site Network. This defines network types to be used on site All inside and outside networks. All inside and outside networks with internet VIP support. All inside networks. All outside networks. All outside networks with internet VIP support. VK8s service network. - SITE_NETWORK_IP_FABRIC: VER IP Fabric network for the site This Virtual network type is used for exposing virtual host on IP Fabric network on the VER site or for endpoint in IP Fabric network

<a id="site-7ecf1d"></a>&#x2022; [`site`](#site-7ecf1d) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Site](#site-7ecf1d) below.

#### Advertise Custom Advertise Where Site Site

A [`site`](#site-7ecf1d) block (within [`advertise_custom.advertise_where.site`](#advertise-custom-advertise-where-site)) supports the following:

<a id="name-201d26"></a>&#x2022; [`name`](#name-201d26) - Optional String<br>Name. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name

<a id="namespace-c3f40d"></a>&#x2022; [`namespace`](#namespace-c3f40d) - Optional String<br>Namespace. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace

<a id="tenant-8a632a"></a>&#x2022; [`tenant`](#tenant-8a632a) - Optional String<br>Tenant. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant

#### Advertise Custom Advertise Where Virtual Network

A [`virtual_network`](#network-a20be3) block (within [`advertise_custom.advertise_where`](#advertise-custom-advertise-where)) supports the following:

<a id="vip-26d874"></a>&#x2022; [`default_v6_vip`](#vip-26d874) - Optional Block<br>Enable this option

<a id="vip-c51931"></a>&#x2022; [`default_vip`](#vip-c51931) - Optional Block<br>Enable this option

<a id="vip-bb67d7"></a>&#x2022; [`specific_v6_vip`](#vip-bb67d7) - Optional String<br>Specific V6 VIP. Use given IPv6 address as VIP on virtual Network

<a id="vip-943090"></a>&#x2022; [`specific_vip`](#vip-943090) - Optional String<br>Specific V4 VIP. Use given IPv4 address as VIP on virtual Network

<a id="network-bff334"></a>&#x2022; [`virtual_network`](#network-bff334) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Virtual Network](#network-bff334) below.

#### Advertise Custom Advertise Where Virtual Network Virtual Network

<a id="deep-802ee3"></a>Deeply nested **Network** block collapsed for readability.

#### Advertise Custom Advertise Where Virtual Site

A [`virtual_site`](#site-5d39fd) block (within [`advertise_custom.advertise_where`](#advertise-custom-advertise-where)) supports the following:

<a id="network-15aca4"></a>&#x2022; [`network`](#network-15aca4) - Optional String  Defaults to `SITE_NETWORK_INSIDE_AND_OUTSIDE`<br>Possible values are `SITE_NETWORK_INSIDE_AND_OUTSIDE`, `SITE_NETWORK_INSIDE`, `SITE_NETWORK_OUTSIDE`, `SITE_NETWORK_SERVICE`, `SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP`, `SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP`, `SITE_NETWORK_IP_FABRIC`<br>[Enum: SITE_NETWORK_INSIDE_AND_OUTSIDE|SITE_NETWORK_INSIDE|SITE_NETWORK_OUTSIDE|SITE_NETWORK_SERVICE|SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP|SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP|SITE_NETWORK_IP_FABRIC] Site Network. This defines network types to be used on site All inside and outside networks. All inside and outside networks with internet VIP support. All inside networks. All outside networks. All outside networks with internet VIP support. VK8s service network. - SITE_NETWORK_IP_FABRIC: VER IP Fabric network for the site This Virtual network type is used for exposing virtual host on IP Fabric network on the VER site or for endpoint in IP Fabric network

<a id="site-04fd53"></a>&#x2022; [`virtual_site`](#site-04fd53) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Virtual Site](#site-04fd53) below.

#### Advertise Custom Advertise Where Virtual Site Virtual Site

<a id="deep-22557e"></a>Deeply nested **Site** block collapsed for readability.

#### Advertise Custom Advertise Where Virtual Site With VIP

<a id="deep-7807f9"></a>Deeply nested **VIP** block collapsed for readability.

#### Advertise Custom Advertise Where Virtual Site With VIP Virtual Site

<a id="deep-60188a"></a>Deeply nested **Site** block collapsed for readability.

#### Advertise Custom Advertise Where Vk8s Service

A [`vk8s_service`](#service-1fdc7a) block (within [`advertise_custom.advertise_where`](#advertise-custom-advertise-where)) supports the following:

<a id="site-ec8d32"></a>&#x2022; [`site`](#site-ec8d32) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Site](#site-ec8d32) below.

<a id="site-5fcbf9"></a>&#x2022; [`virtual_site`](#site-5fcbf9) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Virtual Site](#site-5fcbf9) below.

#### Advertise Custom Advertise Where Vk8s Service Site

A [`site`](#site-ec8d32) block (within [`advertise_custom.advertise_where.vk8s_service`](#service-1fdc7a)) supports the following:

<a id="name-950776"></a>&#x2022; [`name`](#name-950776) - Optional String<br>Name. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name

<a id="namespace-1faf25"></a>&#x2022; [`namespace`](#namespace-1faf25) - Optional String<br>Namespace. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace

<a id="tenant-98cf6a"></a>&#x2022; [`tenant`](#tenant-98cf6a) - Optional String<br>Tenant. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant

#### Advertise Custom Advertise Where Vk8s Service Virtual Site

<a id="deep-e5d00e"></a>Deeply nested **Site** block collapsed for readability.

#### Advertise On Public

An [`advertise_on_public`](#advertise-on-public) block supports the following:

<a id="advertise-on-public-public-ip"></a>&#x2022; [`public_ip`](#advertise-on-public-public-ip) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Public IP](#advertise-on-public-public-ip) below.

#### Advertise On Public Public IP

A [`public_ip`](#advertise-on-public-public-ip) block (within [`advertise_on_public`](#advertise-on-public)) supports the following:

<a id="advertise-on-public-public-ip-name"></a>&#x2022; [`name`](#advertise-on-public-public-ip-name) - Optional String<br>Name. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name

<a id="advertise-on-public-public-ip-namespace"></a>&#x2022; [`namespace`](#advertise-on-public-public-ip-namespace) - Optional String<br>Namespace. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace

<a id="advertise-on-public-public-ip-tenant"></a>&#x2022; [`tenant`](#advertise-on-public-public-ip-tenant) - Optional String<br>Tenant. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant

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

<a id="deep-208464"></a>Deeply nested **Method** block collapsed for readability.

#### API Protection Rules API Endpoint Rules Client Matcher

<a id="deep-ee8161"></a>Deeply nested **Matcher** block collapsed for readability.

#### API Protection Rules API Endpoint Rules Client Matcher Asn List

<a id="deep-fc97af"></a>Deeply nested **List** block collapsed for readability.

#### API Protection Rules API Endpoint Rules Client Matcher Asn Matcher

<a id="deep-cfdbbc"></a>Deeply nested **Matcher** block collapsed for readability.

#### API Protection Rules API Endpoint Rules Client Matcher Asn Matcher Asn Sets

<a id="deep-eeb64b"></a>Deeply nested **Sets** block collapsed for readability.

#### API Protection Rules API Endpoint Rules Client Matcher Client Selector

<a id="deep-c313e8"></a>Deeply nested **Selector** block collapsed for readability.

#### API Protection Rules API Endpoint Rules Client Matcher IP Matcher

<a id="deep-30148e"></a>Deeply nested **Matcher** block collapsed for readability.

#### API Protection Rules API Endpoint Rules Client Matcher IP Matcher Prefix Sets

<a id="deep-470eed"></a>Deeply nested **Sets** block collapsed for readability.

#### API Protection Rules API Endpoint Rules Client Matcher IP Prefix List

<a id="deep-779a4e"></a>Deeply nested **List** block collapsed for readability.

#### API Protection Rules API Endpoint Rules Client Matcher IP Threat Category List

<a id="deep-c487a2"></a>Deeply nested **List** block collapsed for readability.

#### API Protection Rules API Endpoint Rules Client Matcher TLS Fingerprint Matcher

<a id="deep-f6189e"></a>Deeply nested **Matcher** block collapsed for readability.

#### API Protection Rules API Endpoint Rules Metadata

A [`metadata`](#metadata-46451b) block (within [`api_protection_rules.api_endpoint_rules`](#api-protection-rules-api-endpoint-rules)) supports the following:

<a id="spec-5eb432"></a>&#x2022; [`description_spec`](#spec-5eb432) - Optional String<br>Description. Human readable description

<a id="name-af3d78"></a>&#x2022; [`name`](#name-af3d78) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### API Protection Rules API Endpoint Rules Request Matcher

<a id="deep-d99592"></a>Deeply nested **Matcher** block collapsed for readability.

#### API Protection Rules API Endpoint Rules Request Matcher Cookie Matchers

<a id="deep-a8337a"></a>Deeply nested **Matchers** block collapsed for readability.

#### API Protection Rules API Endpoint Rules Request Matcher Cookie Matchers Item

<a id="deep-094373"></a>Deeply nested **Item** block collapsed for readability.

#### API Protection Rules API Endpoint Rules Request Matcher Headers

<a id="deep-472c6f"></a>Deeply nested **Headers** block collapsed for readability.

#### API Protection Rules API Endpoint Rules Request Matcher Headers Item

<a id="deep-c4fea0"></a>Deeply nested **Item** block collapsed for readability.

#### API Protection Rules API Endpoint Rules Request Matcher JWT Claims

<a id="deep-17e736"></a>Deeply nested **Claims** block collapsed for readability.

#### API Protection Rules API Endpoint Rules Request Matcher JWT Claims Item

<a id="deep-b75331"></a>Deeply nested **Item** block collapsed for readability.

#### API Protection Rules API Endpoint Rules Request Matcher Query Params

<a id="deep-325cb1"></a>Deeply nested **Params** block collapsed for readability.

#### API Protection Rules API Endpoint Rules Request Matcher Query Params Item

<a id="deep-d72343"></a>Deeply nested **Item** block collapsed for readability.

#### API Protection Rules API Groups Rules

An [`api_groups_rules`](#api-protection-rules-api-groups-rules) block (within [`api_protection_rules`](#api-protection-rules)) supports the following:

<a id="action-fa62d7"></a>&#x2022; [`action`](#action-fa62d7) - Optional Block<br>API Protection Rule Action. The action to take if the input request matches the rule<br>See [Action](#action-fa62d7) below.

<a id="domain-b1276e"></a>&#x2022; [`any_domain`](#domain-b1276e) - Optional Block<br>Enable this option

<a id="group-a8b675"></a>&#x2022; [`api_group`](#group-a8b675) - Optional String<br>API Group. API groups derived from API Definition swaggers. For example oas-all-operations including all paths and methods from the swaggers, oas-base-URLs covering all requests under base-paths from the swaggers. Custom groups can be created if user tags paths or operations with 'x-F5 Distributed Cloud-API-group' extensions inside swaggers

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

<a id="deep-7bd2af"></a>Deeply nested **Matcher** block collapsed for readability.

#### API Protection Rules API Groups Rules Client Matcher Asn List

<a id="deep-4d0692"></a>Deeply nested **List** block collapsed for readability.

#### API Protection Rules API Groups Rules Client Matcher Asn Matcher

<a id="deep-9e8c02"></a>Deeply nested **Matcher** block collapsed for readability.

#### API Protection Rules API Groups Rules Client Matcher Asn Matcher Asn Sets

<a id="deep-a4292a"></a>Deeply nested **Sets** block collapsed for readability.

#### API Protection Rules API Groups Rules Client Matcher Client Selector

<a id="deep-0efe6a"></a>Deeply nested **Selector** block collapsed for readability.

#### API Protection Rules API Groups Rules Client Matcher IP Matcher

<a id="deep-51392a"></a>Deeply nested **Matcher** block collapsed for readability.

#### API Protection Rules API Groups Rules Client Matcher IP Matcher Prefix Sets

<a id="deep-c60247"></a>Deeply nested **Sets** block collapsed for readability.

#### API Protection Rules API Groups Rules Client Matcher IP Prefix List

<a id="deep-50bd47"></a>Deeply nested **List** block collapsed for readability.

#### API Protection Rules API Groups Rules Client Matcher IP Threat Category List

<a id="deep-1a8131"></a>Deeply nested **List** block collapsed for readability.

#### API Protection Rules API Groups Rules Client Matcher TLS Fingerprint Matcher

<a id="deep-e3bd0a"></a>Deeply nested **Matcher** block collapsed for readability.

#### API Protection Rules API Groups Rules Metadata

A [`metadata`](#metadata-b7fd60) block (within [`api_protection_rules.api_groups_rules`](#api-protection-rules-api-groups-rules)) supports the following:

<a id="spec-ccf62e"></a>&#x2022; [`description_spec`](#spec-ccf62e) - Optional String<br>Description. Human readable description

<a id="name-4148ef"></a>&#x2022; [`name`](#name-4148ef) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### API Protection Rules API Groups Rules Request Matcher

<a id="deep-623c26"></a>Deeply nested **Matcher** block collapsed for readability.

#### API Protection Rules API Groups Rules Request Matcher Cookie Matchers

<a id="deep-f88479"></a>Deeply nested **Matchers** block collapsed for readability.

#### API Protection Rules API Groups Rules Request Matcher Cookie Matchers Item

<a id="deep-0816fa"></a>Deeply nested **Item** block collapsed for readability.

#### API Protection Rules API Groups Rules Request Matcher Headers

<a id="deep-3e07d6"></a>Deeply nested **Headers** block collapsed for readability.

#### API Protection Rules API Groups Rules Request Matcher Headers Item

<a id="deep-bdd2f0"></a>Deeply nested **Item** block collapsed for readability.

#### API Protection Rules API Groups Rules Request Matcher JWT Claims

<a id="deep-cc6da8"></a>Deeply nested **Claims** block collapsed for readability.

#### API Protection Rules API Groups Rules Request Matcher JWT Claims Item

<a id="deep-7f9a23"></a>Deeply nested **Item** block collapsed for readability.

#### API Protection Rules API Groups Rules Request Matcher Query Params

<a id="deep-bef9d5"></a>Deeply nested **Params** block collapsed for readability.

#### API Protection Rules API Groups Rules Request Matcher Query Params Item

<a id="deep-3475f4"></a>Deeply nested **Item** block collapsed for readability.

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

<a id="deep-e1c30c"></a>Deeply nested **Method** block collapsed for readability.

#### API Rate Limit API Endpoint Rules Client Matcher

<a id="deep-89d214"></a>Deeply nested **Matcher** block collapsed for readability.

#### API Rate Limit API Endpoint Rules Client Matcher Asn List

<a id="deep-5b26f9"></a>Deeply nested **List** block collapsed for readability.

#### API Rate Limit API Endpoint Rules Client Matcher Asn Matcher

<a id="deep-fb456a"></a>Deeply nested **Matcher** block collapsed for readability.

#### API Rate Limit API Endpoint Rules Client Matcher Asn Matcher Asn Sets

<a id="deep-5b15e9"></a>Deeply nested **Sets** block collapsed for readability.

#### API Rate Limit API Endpoint Rules Client Matcher Client Selector

<a id="deep-2d36e2"></a>Deeply nested **Selector** block collapsed for readability.

#### API Rate Limit API Endpoint Rules Client Matcher IP Matcher

<a id="deep-c0f678"></a>Deeply nested **Matcher** block collapsed for readability.

#### API Rate Limit API Endpoint Rules Client Matcher IP Matcher Prefix Sets

<a id="deep-70b0c0"></a>Deeply nested **Sets** block collapsed for readability.

#### API Rate Limit API Endpoint Rules Client Matcher IP Prefix List

<a id="deep-5a2c8c"></a>Deeply nested **List** block collapsed for readability.

#### API Rate Limit API Endpoint Rules Client Matcher IP Threat Category List

<a id="deep-67d5ee"></a>Deeply nested **List** block collapsed for readability.

#### API Rate Limit API Endpoint Rules Client Matcher TLS Fingerprint Matcher

<a id="deep-09252e"></a>Deeply nested **Matcher** block collapsed for readability.

#### API Rate Limit API Endpoint Rules Inline Rate Limiter

<a id="deep-5cd429"></a>Deeply nested **Limiter** block collapsed for readability.

#### API Rate Limit API Endpoint Rules Inline Rate Limiter Ref User ID

<a id="deep-70ee5d"></a>Deeply nested **ID** block collapsed for readability.

#### API Rate Limit API Endpoint Rules Ref Rate Limiter

<a id="deep-1175ac"></a>Deeply nested **Limiter** block collapsed for readability.

#### API Rate Limit API Endpoint Rules Request Matcher

<a id="deep-b66b99"></a>Deeply nested **Matcher** block collapsed for readability.

#### API Rate Limit API Endpoint Rules Request Matcher Cookie Matchers

<a id="deep-137e54"></a>Deeply nested **Matchers** block collapsed for readability.

#### API Rate Limit API Endpoint Rules Request Matcher Cookie Matchers Item

<a id="deep-b2970b"></a>Deeply nested **Item** block collapsed for readability.

#### API Rate Limit API Endpoint Rules Request Matcher Headers

<a id="deep-78dc03"></a>Deeply nested **Headers** block collapsed for readability.

#### API Rate Limit API Endpoint Rules Request Matcher Headers Item

<a id="deep-a29608"></a>Deeply nested **Item** block collapsed for readability.

#### API Rate Limit API Endpoint Rules Request Matcher JWT Claims

<a id="deep-01f53c"></a>Deeply nested **Claims** block collapsed for readability.

#### API Rate Limit API Endpoint Rules Request Matcher JWT Claims Item

<a id="deep-1ae8a7"></a>Deeply nested **Item** block collapsed for readability.

#### API Rate Limit API Endpoint Rules Request Matcher Query Params

<a id="deep-bb70f4"></a>Deeply nested **Params** block collapsed for readability.

#### API Rate Limit API Endpoint Rules Request Matcher Query Params Item

<a id="deep-8ff910"></a>Deeply nested **Item** block collapsed for readability.

#### API Rate Limit Bypass Rate Limiting Rules

A [`bypass_rate_limiting_rules`](#rules-776e97) block (within [`api_rate_limit`](#api-rate-limit)) supports the following:

<a id="rules-51aa34"></a>&#x2022; [`bypass_rate_limiting_rules`](#rules-51aa34) - Optional Block<br>Bypass Rate Limiting. This category defines rules per URL or API group. If request matches any of these rules, skip Rate Limiting<br>See [Bypass Rate Limiting Rules](#rules-51aa34) below.

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules

<a id="deep-057f78"></a>Deeply nested **Rules** block collapsed for readability.

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules API Endpoint

<a id="deep-64f716"></a>Deeply nested **Endpoint** block collapsed for readability.

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules API Groups

<a id="deep-fc711c"></a>Deeply nested **Groups** block collapsed for readability.

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Client Matcher

<a id="deep-ffaf3a"></a>Deeply nested **Matcher** block collapsed for readability.

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Client Matcher Asn List

<a id="deep-242661"></a>Deeply nested **List** block collapsed for readability.

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Client Matcher Asn Matcher

<a id="deep-d519c0"></a>Deeply nested **Matcher** block collapsed for readability.

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Client Matcher Asn Matcher Asn Sets

<a id="deep-30e257"></a>Deeply nested **Sets** block collapsed for readability.

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Client Matcher Client Selector

<a id="deep-4a08b6"></a>Deeply nested **Selector** block collapsed for readability.

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Client Matcher IP Matcher

<a id="deep-16bba4"></a>Deeply nested **Matcher** block collapsed for readability.

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Client Matcher IP Matcher Prefix Sets

<a id="deep-da95b1"></a>Deeply nested **Sets** block collapsed for readability.

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Client Matcher IP Prefix List

<a id="deep-60dc79"></a>Deeply nested **List** block collapsed for readability.

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Client Matcher IP Threat Category List

<a id="deep-ed2c35"></a>Deeply nested **List** block collapsed for readability.

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Client Matcher TLS Fingerprint Matcher

<a id="deep-0e8717"></a>Deeply nested **Matcher** block collapsed for readability.

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Request Matcher

<a id="deep-9eed33"></a>Deeply nested **Matcher** block collapsed for readability.

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Request Matcher Cookie Matchers

<a id="deep-dd7483"></a>Deeply nested **Matchers** block collapsed for readability.

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Request Matcher Cookie Matchers Item

<a id="deep-746a10"></a>Deeply nested **Item** block collapsed for readability.

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Request Matcher Headers

<a id="deep-4e8a54"></a>Deeply nested **Headers** block collapsed for readability.

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Request Matcher Headers Item

<a id="deep-328c40"></a>Deeply nested **Item** block collapsed for readability.

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Request Matcher JWT Claims

<a id="deep-d34315"></a>Deeply nested **Claims** block collapsed for readability.

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Request Matcher JWT Claims Item

<a id="deep-9cd3b4"></a>Deeply nested **Item** block collapsed for readability.

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Request Matcher Query Params

<a id="deep-8194a5"></a>Deeply nested **Params** block collapsed for readability.

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Request Matcher Query Params Item

<a id="deep-4cdff3"></a>Deeply nested **Item** block collapsed for readability.

#### API Rate Limit Custom IP Allowed List

A [`custom_ip_allowed_list`](#api-rate-limit-custom-ip-allowed-list) block (within [`api_rate_limit`](#api-rate-limit)) supports the following:

<a id="prefixes-73df46"></a>&#x2022; [`rate_limiter_allowed_prefixes`](#prefixes-73df46) - Optional Block<br>List of IP Prefix Sets. References to ip_prefix_set objects. Requests from source IP addresses that are covered by one of the allowed IP Prefixes are not subjected to rate limiting<br>See [Rate Limiter Allowed Prefixes](#prefixes-73df46) below.

#### API Rate Limit Custom IP Allowed List Rate Limiter Allowed Prefixes

<a id="deep-85bfc8"></a>Deeply nested **Prefixes** block collapsed for readability.

#### API Rate Limit IP Allowed List

An [`ip_allowed_list`](#api-rate-limit-ip-allowed-list) block (within [`api_rate_limit`](#api-rate-limit)) supports the following:

<a id="api-rate-limit-ip-allowed-list-prefixes"></a>&#x2022; [`prefixes`](#api-rate-limit-ip-allowed-list-prefixes) - Optional List<br>IPv4 Prefix List. List of IPv4 prefixes that represent an endpoint

#### API Rate Limit Server URL Rules

A [`server_url_rules`](#api-rate-limit-server-url-rules) block (within [`api_rate_limit`](#api-rate-limit)) supports the following:

<a id="domain-0747c9"></a>&#x2022; [`any_domain`](#domain-0747c9) - Optional Block<br>Enable this option

<a id="group-15c11a"></a>&#x2022; [`api_group`](#group-15c11a) - Optional String<br>API Group. API groups derived from API Definition swaggers. For example oas-all-operations including all paths and methods from the swaggers, oas-base-URLs covering all requests under base-paths from the swaggers. Custom groups can be created if user tags paths or operations with 'x-F5 Distributed Cloud-API-group' extensions inside swaggers

<a id="path-44dbff"></a>&#x2022; [`base_path`](#path-44dbff) - Optional String<br>Base Path. Prefix of the request path

<a id="matcher-ed4b34"></a>&#x2022; [`client_matcher`](#matcher-ed4b34) - Optional Block<br>Client Matcher. Client conditions for matching a rule<br>See [Client Matcher](#matcher-ed4b34) below.

<a id="limiter-9faa53"></a>&#x2022; [`inline_rate_limiter`](#limiter-9faa53) - Optional Block<br>InlineRateLimiter<br>See [Inline Rate Limiter](#limiter-9faa53) below.

<a id="limiter-383ca9"></a>&#x2022; [`ref_rate_limiter`](#limiter-383ca9) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Ref Rate Limiter](#limiter-383ca9) below.

<a id="matcher-d0eea8"></a>&#x2022; [`request_matcher`](#matcher-d0eea8) - Optional Block<br>Request Matcher. Request conditions for matching a rule<br>See [Request Matcher](#matcher-d0eea8) below.

<a id="domain-dca9c1"></a>&#x2022; [`specific_domain`](#domain-dca9c1) - Optional String<br>Specific Domain. The rule will apply for a specific domain

#### API Rate Limit Server URL Rules Client Matcher

<a id="deep-ecbdb7"></a>Deeply nested **Matcher** block collapsed for readability.

#### API Rate Limit Server URL Rules Client Matcher Asn List

<a id="deep-e9df5c"></a>Deeply nested **List** block collapsed for readability.

#### API Rate Limit Server URL Rules Client Matcher Asn Matcher

<a id="deep-da7cfe"></a>Deeply nested **Matcher** block collapsed for readability.

#### API Rate Limit Server URL Rules Client Matcher Asn Matcher Asn Sets

<a id="deep-641e85"></a>Deeply nested **Sets** block collapsed for readability.

#### API Rate Limit Server URL Rules Client Matcher Client Selector

<a id="deep-1cd130"></a>Deeply nested **Selector** block collapsed for readability.

#### API Rate Limit Server URL Rules Client Matcher IP Matcher

<a id="deep-51ff2b"></a>Deeply nested **Matcher** block collapsed for readability.

#### API Rate Limit Server URL Rules Client Matcher IP Matcher Prefix Sets

<a id="deep-df562f"></a>Deeply nested **Sets** block collapsed for readability.

#### API Rate Limit Server URL Rules Client Matcher IP Prefix List

<a id="deep-d61862"></a>Deeply nested **List** block collapsed for readability.

#### API Rate Limit Server URL Rules Client Matcher IP Threat Category List

<a id="deep-d4b74b"></a>Deeply nested **List** block collapsed for readability.

#### API Rate Limit Server URL Rules Client Matcher TLS Fingerprint Matcher

<a id="deep-fab698"></a>Deeply nested **Matcher** block collapsed for readability.

#### API Rate Limit Server URL Rules Inline Rate Limiter

<a id="deep-758e39"></a>Deeply nested **Limiter** block collapsed for readability.

#### API Rate Limit Server URL Rules Inline Rate Limiter Ref User ID

<a id="deep-37d9ec"></a>Deeply nested **ID** block collapsed for readability.

#### API Rate Limit Server URL Rules Ref Rate Limiter

<a id="deep-6a8465"></a>Deeply nested **Limiter** block collapsed for readability.

#### API Rate Limit Server URL Rules Request Matcher

<a id="deep-f791bc"></a>Deeply nested **Matcher** block collapsed for readability.

#### API Rate Limit Server URL Rules Request Matcher Cookie Matchers

<a id="deep-11f932"></a>Deeply nested **Matchers** block collapsed for readability.

#### API Rate Limit Server URL Rules Request Matcher Cookie Matchers Item

<a id="deep-b08554"></a>Deeply nested **Item** block collapsed for readability.

#### API Rate Limit Server URL Rules Request Matcher Headers

<a id="deep-a738af"></a>Deeply nested **Headers** block collapsed for readability.

#### API Rate Limit Server URL Rules Request Matcher Headers Item

<a id="deep-28a0f1"></a>Deeply nested **Item** block collapsed for readability.

#### API Rate Limit Server URL Rules Request Matcher JWT Claims

<a id="deep-4069dc"></a>Deeply nested **Claims** block collapsed for readability.

#### API Rate Limit Server URL Rules Request Matcher JWT Claims Item

<a id="deep-e1e8f7"></a>Deeply nested **Item** block collapsed for readability.

#### API Rate Limit Server URL Rules Request Matcher Query Params

<a id="deep-015644"></a>Deeply nested **Params** block collapsed for readability.

#### API Rate Limit Server URL Rules Request Matcher Query Params Item

<a id="deep-ba63fa"></a>Deeply nested **Item** block collapsed for readability.

#### API Specification

An [`api_specification`](#api-specification) block supports the following:

<a id="api-specification-api-definition"></a>&#x2022; [`api_definition`](#api-specification-api-definition) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [API Definition](#api-specification-api-definition) below.

<a id="endpoints-4158a4"></a>&#x2022; [`validation_all_spec_endpoints`](#endpoints-4158a4) - Optional Block<br>API Inventory. Settings for API Inventory validation<br>See [Validation All Spec Endpoints](#endpoints-4158a4) below.

<a id="list-23b577"></a>&#x2022; [`validation_custom_list`](#list-23b577) - Optional Block<br>Custom List. Define API groups, base paths, or API endpoints and their OpenAPI validation modes. Any other API-endpoint not listed will act according to 'Fall Through Mode'<br>See [Validation Custom List](#list-23b577) below.

<a id="api-specification-validation-disabled"></a>&#x2022; [`validation_disabled`](#api-specification-validation-disabled) - Optional Block<br>Enable this option

#### API Specification API Definition

An [`api_definition`](#api-specification-api-definition) block (within [`api_specification`](#api-specification)) supports the following:

<a id="api-specification-api-definition-name"></a>&#x2022; [`name`](#api-specification-api-definition-name) - Optional String<br>Name. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name

<a id="namespace-c685bf"></a>&#x2022; [`namespace`](#namespace-c685bf) - Optional String<br>Namespace. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace

<a id="api-specification-api-definition-tenant"></a>&#x2022; [`tenant`](#api-specification-api-definition-tenant) - Optional String<br>Tenant. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant

#### API Specification Validation All Spec Endpoints

A [`validation_all_spec_endpoints`](#endpoints-4158a4) block (within [`api_specification`](#api-specification)) supports the following:

<a id="mode-8425c5"></a>&#x2022; [`fall_through_mode`](#mode-8425c5) - Optional Block<br>Fall Through Mode. Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a. Swagger) or doesn't have a specific rule in custom rules)<br>See [Fall Through Mode](#mode-8425c5) below.

<a id="settings-a83a93"></a>&#x2022; [`settings`](#settings-a83a93) - Optional Block<br>Common Settings. OpenAPI specification validation settings relevant for 'API Inventory' enforcement and for 'Custom list' enforcement<br>See [Settings](#settings-a83a93) below.

<a id="mode-cd4a1c"></a>&#x2022; [`validation_mode`](#mode-cd4a1c) - Optional Block<br>Validation Mode. Validation mode of OpenAPI specification. When a validation mismatch occurs on a request to one of the endpoints listed on the OpenAPI specification file (a.k.a. Swagger)<br>See [Validation Mode](#mode-cd4a1c) below.

#### API Specification Validation All Spec Endpoints Fall Through Mode

<a id="deep-93854f"></a>Deeply nested **Mode** block collapsed for readability.

#### API Specification Validation All Spec Endpoints Fall Through Mode Fall Through Mode Custom

<a id="deep-fcb4c6"></a>Deeply nested **Custom** block collapsed for readability.

#### API Specification Validation All Spec Endpoints Fall Through Mode Fall Through Mode Custom Open API Validation Rules

<a id="deep-5bd981"></a>Deeply nested **Rules** block collapsed for readability.

#### API Specification Validation All Spec Endpoints Fall Through Mode Fall Through Mode Custom Open API Validation Rules API Endpoint

<a id="deep-96d6ca"></a>Deeply nested **Endpoint** block collapsed for readability.

#### API Specification Validation All Spec Endpoints Fall Through Mode Fall Through Mode Custom Open API Validation Rules Metadata

<a id="deep-d275eb"></a>Deeply nested **Metadata** block collapsed for readability.

#### API Specification Validation All Spec Endpoints Settings

A [`settings`](#settings-a83a93) block (within [`api_specification.validation_all_spec_endpoints`](#endpoints-4158a4)) supports the following:

<a id="validation-462f95"></a>&#x2022; [`oversized_body_fail_validation`](#validation-462f95) - Optional Block<br>Enable this option

<a id="validation-7ffaab"></a>&#x2022; [`oversized_body_skip_validation`](#validation-7ffaab) - Optional Block<br>Enable this option

<a id="custom-8254df"></a>&#x2022; [`property_validation_settings_custom`](#custom-8254df) - Optional Block<br>Validation Property Settings. Custom property validation settings<br>See [Property Validation Settings Custom](#custom-8254df) below.

<a id="default-f746bd"></a>&#x2022; [`property_validation_settings_default`](#default-f746bd) - Optional Block<br>Enable this option

#### API Specification Validation All Spec Endpoints Settings Property Validation Settings Custom

<a id="deep-57507d"></a>Deeply nested **Custom** block collapsed for readability.

#### API Specification Validation All Spec Endpoints Settings Property Validation Settings Custom Query Parameters

<a id="deep-761ec3"></a>Deeply nested **Parameters** block collapsed for readability.

#### API Specification Validation All Spec Endpoints Validation Mode

<a id="deep-a84c66"></a>Deeply nested **Mode** block collapsed for readability.

#### API Specification Validation All Spec Endpoints Validation Mode Response Validation Mode Active

<a id="deep-7d440a"></a>Deeply nested **Active** block collapsed for readability.

#### API Specification Validation All Spec Endpoints Validation Mode Validation Mode Active

<a id="deep-dcf3e9"></a>Deeply nested **Active** block collapsed for readability.

#### API Specification Validation Custom List

A [`validation_custom_list`](#list-23b577) block (within [`api_specification`](#api-specification)) supports the following:

<a id="mode-146cc3"></a>&#x2022; [`fall_through_mode`](#mode-146cc3) - Optional Block<br>Fall Through Mode. Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a. Swagger) or doesn't have a specific rule in custom rules)<br>See [Fall Through Mode](#mode-146cc3) below.

<a id="rules-f51668"></a>&#x2022; [`open_api_validation_rules`](#rules-f51668) - Optional Block<br>Validation List<br>See [Open API Validation Rules](#rules-f51668) below.

<a id="settings-940e64"></a>&#x2022; [`settings`](#settings-940e64) - Optional Block<br>Common Settings. OpenAPI specification validation settings relevant for 'API Inventory' enforcement and for 'Custom list' enforcement<br>See [Settings](#settings-940e64) below.

#### API Specification Validation Custom List Fall Through Mode

<a id="deep-08870a"></a>Deeply nested **Mode** block collapsed for readability.

#### API Specification Validation Custom List Fall Through Mode Fall Through Mode Custom

<a id="deep-1af5fd"></a>Deeply nested **Custom** block collapsed for readability.

#### API Specification Validation Custom List Fall Through Mode Fall Through Mode Custom Open API Validation Rules

<a id="deep-e08601"></a>Deeply nested **Rules** block collapsed for readability.

#### API Specification Validation Custom List Fall Through Mode Fall Through Mode Custom Open API Validation Rules API Endpoint

<a id="deep-f88a1e"></a>Deeply nested **Endpoint** block collapsed for readability.

#### API Specification Validation Custom List Fall Through Mode Fall Through Mode Custom Open API Validation Rules Metadata

<a id="deep-8b9a6c"></a>Deeply nested **Metadata** block collapsed for readability.

#### API Specification Validation Custom List Open API Validation Rules

<a id="deep-59b908"></a>Deeply nested **Rules** block collapsed for readability.

#### API Specification Validation Custom List Open API Validation Rules API Endpoint

<a id="deep-287681"></a>Deeply nested **Endpoint** block collapsed for readability.

#### API Specification Validation Custom List Open API Validation Rules Metadata

<a id="deep-bdc351"></a>Deeply nested **Metadata** block collapsed for readability.

#### API Specification Validation Custom List Open API Validation Rules Validation Mode

<a id="deep-e4ce7e"></a>Deeply nested **Mode** block collapsed for readability.

#### API Specification Validation Custom List Open API Validation Rules Validation Mode Response Validation Mode Active

<a id="deep-c312ba"></a>Deeply nested **Active** block collapsed for readability.

#### API Specification Validation Custom List Open API Validation Rules Validation Mode Validation Mode Active

<a id="deep-0ccdf9"></a>Deeply nested **Active** block collapsed for readability.

#### API Specification Validation Custom List Settings

A [`settings`](#settings-940e64) block (within [`api_specification.validation_custom_list`](#list-23b577)) supports the following:

<a id="validation-cfaf7f"></a>&#x2022; [`oversized_body_fail_validation`](#validation-cfaf7f) - Optional Block<br>Enable this option

<a id="validation-0639fa"></a>&#x2022; [`oversized_body_skip_validation`](#validation-0639fa) - Optional Block<br>Enable this option

<a id="custom-8e6ea6"></a>&#x2022; [`property_validation_settings_custom`](#custom-8e6ea6) - Optional Block<br>Validation Property Settings. Custom property validation settings<br>See [Property Validation Settings Custom](#custom-8e6ea6) below.

<a id="default-baec50"></a>&#x2022; [`property_validation_settings_default`](#default-baec50) - Optional Block<br>Enable this option

#### API Specification Validation Custom List Settings Property Validation Settings Custom

<a id="deep-c74d74"></a>Deeply nested **Custom** block collapsed for readability.

#### API Specification Validation Custom List Settings Property Validation Settings Custom Query Parameters

<a id="deep-b42c34"></a>Deeply nested **Parameters** block collapsed for readability.

#### API Testing

An [`api_testing`](#api-testing) block supports the following:

<a id="api-testing-custom-header-value"></a>&#x2022; [`custom_header_value`](#api-testing-custom-header-value) - Optional String<br>Custom Header. Add x-F5-API-testing-identifier header value to prevent security flags on API testing traffic

<a id="api-testing-domains"></a>&#x2022; [`domains`](#api-testing-domains) - Optional Block<br>Testing Environments. Add and configure testing domains and credentials<br>See [Domains](#api-testing-domains) below.

<a id="api-testing-every-day"></a>&#x2022; [`every_day`](#api-testing-every-day) - Optional Block<br>Enable this option

<a id="api-testing-every-month"></a>&#x2022; [`every_month`](#api-testing-every-month) - Optional Block<br>Enable this option

<a id="api-testing-every-week"></a>&#x2022; [`every_week`](#api-testing-every-week) - Optional Block<br>Enable this option

#### API Testing Domains

A [`domains`](#api-testing-domains) block (within [`api_testing`](#api-testing)) supports the following:

<a id="methods-c3ca06"></a>&#x2022; [`allow_destructive_methods`](#methods-c3ca06) - Optional Bool<br>Use Destructive Methods (e.g., DELETE, PUT). Enable to allow API test to execute destructive methods. Be cautious as these can alter or DELETE data

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

<a id="deep-df8526"></a>Deeply nested **Info** block collapsed for readability.

#### API Testing Domains Credentials API Key Value Clear Secret Info

<a id="deep-ce1400"></a>Deeply nested **Info** block collapsed for readability.

#### API Testing Domains Credentials Basic Auth

A [`basic_auth`](#auth-4868f3) block (within [`api_testing.domains.credentials`](#api-testing-domains-credentials)) supports the following:

<a id="password-e6a065"></a>&#x2022; [`password`](#password-e6a065) - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Password](#password-e6a065) below.

<a id="user-2d5b74"></a>&#x2022; [`user`](#user-2d5b74) - Optional String<br>User

#### API Testing Domains Credentials Basic Auth Password

A [`password`](#password-e6a065) block (within [`api_testing.domains.credentials.basic_auth`](#auth-4868f3)) supports the following:

<a id="info-0decf2"></a>&#x2022; [`blindfold_secret_info`](#info-0decf2) - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#info-0decf2) below.

<a id="info-71b4da"></a>&#x2022; [`clear_secret_info`](#info-71b4da) - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#info-71b4da) below.

#### API Testing Domains Credentials Basic Auth Password Blindfold Secret Info

<a id="deep-f21ffa"></a>Deeply nested **Info** block collapsed for readability.

#### API Testing Domains Credentials Basic Auth Password Clear Secret Info

<a id="deep-2aa8d4"></a>Deeply nested **Info** block collapsed for readability.

#### API Testing Domains Credentials Bearer Token

A [`bearer_token`](#token-2a2002) block (within [`api_testing.domains.credentials`](#api-testing-domains-credentials)) supports the following:

<a id="token-7fcc22"></a>&#x2022; [`token`](#token-7fcc22) - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Token](#token-7fcc22) below.

#### API Testing Domains Credentials Bearer Token Token

A [`token`](#token-7fcc22) block (within [`api_testing.domains.credentials.bearer_token`](#token-2a2002)) supports the following:

<a id="info-ce1236"></a>&#x2022; [`blindfold_secret_info`](#info-ce1236) - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#info-ce1236) below.

<a id="info-5e8cda"></a>&#x2022; [`clear_secret_info`](#info-5e8cda) - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#info-5e8cda) below.

#### API Testing Domains Credentials Bearer Token Token Blindfold Secret Info

<a id="deep-7201df"></a>Deeply nested **Info** block collapsed for readability.

#### API Testing Domains Credentials Bearer Token Token Clear Secret Info

<a id="deep-9d2680"></a>Deeply nested **Info** block collapsed for readability.

#### API Testing Domains Credentials Login Endpoint

A [`login_endpoint`](#endpoint-08dc4d) block (within [`api_testing.domains.credentials`](#api-testing-domains-credentials)) supports the following:

<a id="payload-cb9ddd"></a>&#x2022; [`json_payload`](#payload-cb9ddd) - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [JSON Payload](#payload-cb9ddd) below.

<a id="method-ccde95"></a>&#x2022; [`method`](#method-ccde95) - Optional String  Defaults to `ANY`<br>See [HTTP Methods](#common-http-methods)<br> HTTP Method. Specifies the HTTP method used to access a resource. Any HTTP Method

<a id="path-fc4b22"></a>&#x2022; [`path`](#path-fc4b22) - Optional String<br>Path

<a id="key-04416f"></a>&#x2022; [`token_response_key`](#key-04416f) - Optional String<br>Token Response Key. Specifies how to handle the API response, extracting authentication tokens

#### API Testing Domains Credentials Login Endpoint JSON Payload

<a id="deep-d1d280"></a>Deeply nested **Payload** block collapsed for readability.

#### API Testing Domains Credentials Login Endpoint JSON Payload Blindfold Secret Info

<a id="deep-5f084b"></a>Deeply nested **Info** block collapsed for readability.

#### API Testing Domains Credentials Login Endpoint JSON Payload Clear Secret Info

<a id="deep-9c013b"></a>Deeply nested **Info** block collapsed for readability.

#### App Firewall

An [`app_firewall`](#app-firewall) block supports the following:

<a id="app-firewall-name"></a>&#x2022; [`name`](#app-firewall-name) - Optional String<br>Name. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name

<a id="app-firewall-namespace"></a>&#x2022; [`namespace`](#app-firewall-namespace) - Optional String<br>Namespace. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace

<a id="app-firewall-tenant"></a>&#x2022; [`tenant`](#app-firewall-tenant) - Optional String<br>Tenant. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant

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

<a id="bot-defense-policy"></a>&#x2022; [`policy`](#bot-defense-policy) - Optional Block<br>Bot Defense Policy. This defines various configuration OPTIONS for Bot Defense policy<br>See [Policy](#bot-defense-policy) below.

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

<a id="location-3a398d"></a>&#x2022; [`javascript_location`](#location-3a398d) - Optional String  Defaults to `AFTER_HEAD`<br>Possible values are `AFTER_HEAD`, `AFTER_TITLE_END`, `BEFORE_SCRIPT`<br>[Enum: AFTER_HEAD|AFTER_TITLE_END|BEFORE_SCRIPT] JavaScript Location. All inside networks. Insert JavaScript after `<HEAD>` tag Insert JavaScript after `</title>` tag. Insert JavaScript before first &lt;script> tag

#### Bot Defense Policy Js Insert All Pages Except

<a id="deep-5eada5"></a>Deeply nested **Except** block collapsed for readability.

#### Bot Defense Policy Js Insert All Pages Except Exclude List

<a id="deep-cc9d4f"></a>Deeply nested **List** block collapsed for readability.

#### Bot Defense Policy Js Insert All Pages Except Exclude List Domain

<a id="deep-58ae8c"></a>Deeply nested **Domain** block collapsed for readability.

#### Bot Defense Policy Js Insert All Pages Except Exclude List Metadata

<a id="deep-ac7397"></a>Deeply nested **Metadata** block collapsed for readability.

#### Bot Defense Policy Js Insert All Pages Except Exclude List Path

<a id="deep-9d17be"></a>Deeply nested **Path** block collapsed for readability.

#### Bot Defense Policy Js Insertion Rules

A [`js_insertion_rules`](#bot-defense-policy-js-insertion-rules) block (within [`bot_defense.policy`](#bot-defense-policy)) supports the following:

<a id="list-51668b"></a>&#x2022; [`exclude_list`](#list-51668b) - Optional Block<br>Exclude Paths. Optional JavaScript insertions exclude list of domain and path matchers<br>See [Exclude List](#list-51668b) below.

<a id="rules-15d983"></a>&#x2022; [`rules`](#rules-15d983) - Optional Block<br>JavaScript Insertions. Required list of pages to insert Bot Defense client JavaScript<br>See [Rules](#rules-15d983) below.

#### Bot Defense Policy Js Insertion Rules Exclude List

<a id="deep-5976c7"></a>Deeply nested **List** block collapsed for readability.

#### Bot Defense Policy Js Insertion Rules Exclude List Domain

<a id="deep-c9a8a1"></a>Deeply nested **Domain** block collapsed for readability.

#### Bot Defense Policy Js Insertion Rules Exclude List Metadata

<a id="deep-4b3d6a"></a>Deeply nested **Metadata** block collapsed for readability.

#### Bot Defense Policy Js Insertion Rules Exclude List Path

<a id="deep-fa04d6"></a>Deeply nested **Path** block collapsed for readability.

#### Bot Defense Policy Js Insertion Rules Rules

A [`rules`](#rules-15d983) block (within [`bot_defense.policy.js_insertion_rules`](#bot-defense-policy-js-insertion-rules)) supports the following:

<a id="domain-f27f00"></a>&#x2022; [`any_domain`](#domain-f27f00) - Optional Block<br>Enable this option

<a id="domain-834b0f"></a>&#x2022; [`domain`](#domain-834b0f) - Optional Block<br>Domains. Domains names<br>See [Domain](#domain-834b0f) below.

<a id="location-16277f"></a>&#x2022; [`javascript_location`](#location-16277f) - Optional String  Defaults to `AFTER_HEAD`<br>Possible values are `AFTER_HEAD`, `AFTER_TITLE_END`, `BEFORE_SCRIPT`<br>[Enum: AFTER_HEAD|AFTER_TITLE_END|BEFORE_SCRIPT] JavaScript Location. All inside networks. Insert JavaScript after `<HEAD>` tag Insert JavaScript after `</title>` tag. Insert JavaScript before first &lt;script> tag

<a id="metadata-e15703"></a>&#x2022; [`metadata`](#metadata-e15703) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#metadata-e15703) below.

<a id="path-711518"></a>&#x2022; [`path`](#path-711518) - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#path-711518) below.

#### Bot Defense Policy Js Insertion Rules Rules Domain

<a id="deep-72024a"></a>Deeply nested **Domain** block collapsed for readability.

#### Bot Defense Policy Js Insertion Rules Rules Metadata

<a id="deep-e9d554"></a>Deeply nested **Metadata** block collapsed for readability.

#### Bot Defense Policy Js Insertion Rules Rules Path

<a id="deep-3fdb24"></a>Deeply nested **Path** block collapsed for readability.

#### Bot Defense Policy Mobile Sdk Config

A [`mobile_sdk_config`](#bot-defense-policy-mobile-sdk-config) block (within [`bot_defense.policy`](#bot-defense-policy)) supports the following:

<a id="identifier-e34e48"></a>&#x2022; [`mobile_identifier`](#identifier-e34e48) - Optional Block<br>Mobile Traffic Identifier. Mobile traffic identifier type<br>See [Mobile Identifier](#identifier-e34e48) below.

#### Bot Defense Policy Mobile Sdk Config Mobile Identifier

<a id="deep-8f9ee4"></a>Deeply nested **Identifier** block collapsed for readability.

#### Bot Defense Policy Mobile Sdk Config Mobile Identifier Headers

<a id="deep-235998"></a>Deeply nested **Headers** block collapsed for readability.

#### Bot Defense Policy Mobile Sdk Config Mobile Identifier Headers Item

<a id="deep-ba14c6"></a>Deeply nested **Item** block collapsed for readability.

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

<a id="protocol-21c1f1"></a>&#x2022; [`protocol`](#protocol-21c1f1) - Optional String  Defaults to `BOTH`<br>Possible values are `BOTH`, `HTTP`, `HTTPS`<br>[Enum: BOTH|HTTP|HTTPS] URL Scheme. SchemeType is used to indicate URL scheme. - BOTH: BOTH URL scheme for HTTPS:// or HTTP://. - HTTP: HTTP URL scheme HTTP:// only. - HTTPS: HTTPS URL scheme HTTPS:// only

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

<a id="deep-aedbc5"></a>Deeply nested **Label** block collapsed for readability.

#### Bot Defense Policy Protected App Endpoints Flow Label Account Management

<a id="deep-e2bcec"></a>Deeply nested **Management** block collapsed for readability.

#### Bot Defense Policy Protected App Endpoints Flow Label Authentication

<a id="deep-9d3bff"></a>Deeply nested **Authentication** block collapsed for readability.

#### Bot Defense Policy Protected App Endpoints Flow Label Authentication Login

<a id="deep-f6b65a"></a>Deeply nested **Login** block collapsed for readability.

#### Bot Defense Policy Protected App Endpoints Flow Label Authentication Login Transaction Result

<a id="deep-1cd500"></a>Deeply nested **Result** block collapsed for readability.

#### Bot Defense Policy Protected App Endpoints Flow Label Authentication Login Transaction Result Failure Conditions

<a id="deep-95bfe3"></a>Deeply nested **Conditions** block collapsed for readability.

#### Bot Defense Policy Protected App Endpoints Flow Label Authentication Login Transaction Result Success Conditions

<a id="deep-0eb365"></a>Deeply nested **Conditions** block collapsed for readability.

#### Bot Defense Policy Protected App Endpoints Flow Label Financial Services

<a id="deep-e8aad0"></a>Deeply nested **Services** block collapsed for readability.

#### Bot Defense Policy Protected App Endpoints Flow Label Flight

<a id="deep-fe2123"></a>Deeply nested **Flight** block collapsed for readability.

#### Bot Defense Policy Protected App Endpoints Flow Label Profile Management

<a id="deep-e241e6"></a>Deeply nested **Management** block collapsed for readability.

#### Bot Defense Policy Protected App Endpoints Flow Label Search

<a id="deep-b86aa6"></a>Deeply nested **Search** block collapsed for readability.

#### Bot Defense Policy Protected App Endpoints Flow Label Shopping Gift Cards

<a id="deep-ae3c07"></a>Deeply nested **Cards** block collapsed for readability.

#### Bot Defense Policy Protected App Endpoints Headers

A [`headers`](#headers-986193) block (within [`bot_defense.policy.protected_app_endpoints`](#endpoints-01a2f3)) supports the following:

<a id="present-2e9857"></a>&#x2022; [`check_not_present`](#present-2e9857) - Optional Block<br>Enable this option

<a id="present-3a1075"></a>&#x2022; [`check_present`](#present-3a1075) - Optional Block<br>Enable this option

<a id="matcher-66fb69"></a>&#x2022; [`invert_matcher`](#matcher-66fb69) - Optional Bool<br>Invert Header Matcher. Invert the match result

<a id="item-ca0df2"></a>&#x2022; [`item`](#item-ca0df2) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#item-ca0df2) below.

<a id="name-34d16a"></a>&#x2022; [`name`](#name-34d16a) - Optional String<br>Header Name. A case-insensitive HTTP header name

#### Bot Defense Policy Protected App Endpoints Headers Item

<a id="deep-92cb3a"></a>Deeply nested **Item** block collapsed for readability.

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

<a id="deep-3c0d5a"></a>Deeply nested **Block** block collapsed for readability.

#### Bot Defense Policy Protected App Endpoints Mitigation Flag

<a id="deep-31a90f"></a>Deeply nested **Flag** block collapsed for readability.

#### Bot Defense Policy Protected App Endpoints Mitigation Flag Append Headers

<a id="deep-be96ae"></a>Deeply nested **Headers** block collapsed for readability.

#### Bot Defense Policy Protected App Endpoints Mitigation Redirect

<a id="deep-7565c4"></a>Deeply nested **Redirect** block collapsed for readability.

#### Bot Defense Policy Protected App Endpoints Path

A [`path`](#path-d5ee15) block (within [`bot_defense.policy.protected_app_endpoints`](#endpoints-01a2f3)) supports the following:

<a id="path-16664a"></a>&#x2022; [`path`](#path-16664a) - Optional String<br>Exact. Exact path value to match

<a id="prefix-5a090b"></a>&#x2022; [`prefix`](#prefix-5a090b) - Optional String<br>Prefix. Path prefix to match (e.g. The value / will match on all paths)

<a id="regex-1a10e9"></a>&#x2022; [`regex`](#regex-1a10e9) - Optional String<br>Regex. Regular expression of path match (e.g. The value .* will match on all paths)

#### Bot Defense Policy Protected App Endpoints Query Params

<a id="deep-0d34ba"></a>Deeply nested **Params** block collapsed for readability.

#### Bot Defense Policy Protected App Endpoints Query Params Item

<a id="deep-479116"></a>Deeply nested **Item** block collapsed for readability.

#### Bot Defense Policy Protected App Endpoints Web Mobile

<a id="deep-ee5137"></a>Deeply nested **Mobile** block collapsed for readability.

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

<a id="location-f54ccc"></a>&#x2022; [`javascript_location`](#location-f54ccc) - Optional String  Defaults to `AFTER_HEAD`<br>Possible values are `AFTER_HEAD`, `AFTER_TITLE_END`, `BEFORE_SCRIPT`<br>[Enum: AFTER_HEAD|AFTER_TITLE_END|BEFORE_SCRIPT] JavaScript Location. All inside networks. Insert JavaScript after `<HEAD>` tag Insert JavaScript after `</title>` tag. Insert JavaScript before first &lt;script> tag

#### Bot Defense Advanced Js Insert All Pages Except

<a id="deep-5cdaa3"></a>Deeply nested **Except** block collapsed for readability.

#### Bot Defense Advanced Js Insert All Pages Except Exclude List

<a id="deep-19163b"></a>Deeply nested **List** block collapsed for readability.

#### Bot Defense Advanced Js Insert All Pages Except Exclude List Domain

<a id="deep-10a658"></a>Deeply nested **Domain** block collapsed for readability.

#### Bot Defense Advanced Js Insert All Pages Except Exclude List Metadata

<a id="deep-b603fd"></a>Deeply nested **Metadata** block collapsed for readability.

#### Bot Defense Advanced Js Insert All Pages Except Exclude List Path

<a id="deep-f29bea"></a>Deeply nested **Path** block collapsed for readability.

#### Bot Defense Advanced Js Insertion Rules

A [`js_insertion_rules`](#bot-defense-advanced-js-insertion-rules) block (within [`bot_defense_advanced`](#bot-defense-advanced)) supports the following:

<a id="list-efb443"></a>&#x2022; [`exclude_list`](#list-efb443) - Optional Block<br>Exclude Paths. Optional JavaScript insertions exclude list of domain and path matchers<br>See [Exclude List](#list-efb443) below.

<a id="rules-24e5a0"></a>&#x2022; [`rules`](#rules-24e5a0) - Optional Block<br>JavaScript Insertions. Required list of pages to insert Bot Defense client JavaScript<br>See [Rules](#rules-24e5a0) below.

#### Bot Defense Advanced Js Insertion Rules Exclude List

<a id="deep-658a1f"></a>Deeply nested **List** block collapsed for readability.

#### Bot Defense Advanced Js Insertion Rules Exclude List Domain

<a id="deep-af9c74"></a>Deeply nested **Domain** block collapsed for readability.

#### Bot Defense Advanced Js Insertion Rules Exclude List Metadata

<a id="deep-29f8b7"></a>Deeply nested **Metadata** block collapsed for readability.

#### Bot Defense Advanced Js Insertion Rules Exclude List Path

<a id="deep-b98dbe"></a>Deeply nested **Path** block collapsed for readability.

#### Bot Defense Advanced Js Insertion Rules Rules

A [`rules`](#rules-24e5a0) block (within [`bot_defense_advanced.js_insertion_rules`](#bot-defense-advanced-js-insertion-rules)) supports the following:

<a id="domain-bd13eb"></a>&#x2022; [`any_domain`](#domain-bd13eb) - Optional Block<br>Enable this option

<a id="domain-ff2f2e"></a>&#x2022; [`domain`](#domain-ff2f2e) - Optional Block<br>Domains. Domains names<br>See [Domain](#domain-ff2f2e) below.

<a id="location-20f540"></a>&#x2022; [`javascript_location`](#location-20f540) - Optional String  Defaults to `AFTER_HEAD`<br>Possible values are `AFTER_HEAD`, `AFTER_TITLE_END`, `BEFORE_SCRIPT`<br>[Enum: AFTER_HEAD|AFTER_TITLE_END|BEFORE_SCRIPT] JavaScript Location. All inside networks. Insert JavaScript after `<HEAD>` tag Insert JavaScript after `</title>` tag. Insert JavaScript before first &lt;script> tag

<a id="metadata-43c6ee"></a>&#x2022; [`metadata`](#metadata-43c6ee) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#metadata-43c6ee) below.

<a id="path-a4408d"></a>&#x2022; [`path`](#path-a4408d) - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#path-a4408d) below.

#### Bot Defense Advanced Js Insertion Rules Rules Domain

<a id="deep-879809"></a>Deeply nested **Domain** block collapsed for readability.

#### Bot Defense Advanced Js Insertion Rules Rules Metadata

<a id="deep-b2bae1"></a>Deeply nested **Metadata** block collapsed for readability.

#### Bot Defense Advanced Js Insertion Rules Rules Path

<a id="deep-f7389d"></a>Deeply nested **Path** block collapsed for readability.

#### Bot Defense Advanced Mobile

A [`mobile`](#bot-defense-advanced-mobile) block (within [`bot_defense_advanced`](#bot-defense-advanced)) supports the following:

<a id="bot-defense-advanced-mobile-name"></a>&#x2022; [`name`](#bot-defense-advanced-mobile-name) - Optional String<br>Name. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name

<a id="bot-defense-advanced-mobile-namespace"></a>&#x2022; [`namespace`](#bot-defense-advanced-mobile-namespace) - Optional String<br>Namespace. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace

<a id="bot-defense-advanced-mobile-tenant"></a>&#x2022; [`tenant`](#bot-defense-advanced-mobile-tenant) - Optional String<br>Tenant. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant

#### Bot Defense Advanced Mobile Sdk Config

A [`mobile_sdk_config`](#bot-defense-advanced-mobile-sdk-config) block (within [`bot_defense_advanced`](#bot-defense-advanced)) supports the following:

<a id="identifier-163438"></a>&#x2022; [`mobile_identifier`](#identifier-163438) - Optional Block<br>Mobile Traffic Identifier. Mobile traffic identifier type<br>See [Mobile Identifier](#identifier-163438) below.

#### Bot Defense Advanced Mobile Sdk Config Mobile Identifier

<a id="deep-f229dd"></a>Deeply nested **Identifier** block collapsed for readability.

#### Bot Defense Advanced Mobile Sdk Config Mobile Identifier Headers

<a id="deep-6d616d"></a>Deeply nested **Headers** block collapsed for readability.

#### Bot Defense Advanced Mobile Sdk Config Mobile Identifier Headers Item

<a id="deep-83444d"></a>Deeply nested **Item** block collapsed for readability.

#### Bot Defense Advanced Web

A [`web`](#bot-defense-advanced-web) block (within [`bot_defense_advanced`](#bot-defense-advanced)) supports the following:

<a id="bot-defense-advanced-web-name"></a>&#x2022; [`name`](#bot-defense-advanced-web-name) - Optional String<br>Name. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name

<a id="bot-defense-advanced-web-namespace"></a>&#x2022; [`namespace`](#bot-defense-advanced-web-namespace) - Optional String<br>Namespace. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace

<a id="bot-defense-advanced-web-tenant"></a>&#x2022; [`tenant`](#bot-defense-advanced-web-tenant) - Optional String<br>Tenant. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant

#### Caching Policy

A [`caching_policy`](#caching-policy) block supports the following:

<a id="caching-policy-custom-cache-rule"></a>&#x2022; [`custom_cache_rule`](#caching-policy-custom-cache-rule) - Optional Block<br>Custom Cache Rules. Caching policies for CDN<br>See [Custom Cache Rule](#caching-policy-custom-cache-rule) below.

<a id="caching-policy-default-cache-action"></a>&#x2022; [`default_cache_action`](#caching-policy-default-cache-action) - Optional Block<br>Default Cache Behaviour. This defines a Default Cache Action<br>See [Default Cache Action](#caching-policy-default-cache-action) below.

#### Caching Policy Custom Cache Rule

A [`custom_cache_rule`](#caching-policy-custom-cache-rule) block (within [`caching_policy`](#caching-policy)) supports the following:

<a id="rules-e10c80"></a>&#x2022; [`cdn_cache_rules`](#rules-e10c80) - Optional Block<br>CDN Cache Rule. Reference to CDN Cache Rule configuration object<br>See [CDN Cache Rules](#rules-e10c80) below.

#### Caching Policy Custom Cache Rule CDN Cache Rules

<a id="deep-6f3e1b"></a>Deeply nested **Rules** block collapsed for readability.

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

<a id="client-side-defense-policy"></a>&#x2022; [`policy`](#client-side-defense-policy) - Optional Block<br>Client-Side Defense Policy. This defines various configuration OPTIONS for Client-Side Defense policy<br>See [Policy](#client-side-defense-policy) below.

#### Client Side Defense Policy

A [`policy`](#client-side-defense-policy) block (within [`client_side_defense`](#client-side-defense)) supports the following:

<a id="insert-683e69"></a>&#x2022; [`disable_js_insert`](#insert-683e69) - Optional Block<br>Enable this option

<a id="pages-38bd1c"></a>&#x2022; [`js_insert_all_pages`](#pages-38bd1c) - Optional Block<br>Enable this option

<a id="except-7bfe85"></a>&#x2022; [`js_insert_all_pages_except`](#except-7bfe85) - Optional Block<br>Insert JavaScript in All Pages with the Exceptions. Insert Client-Side Defense JavaScript in all pages with the exceptions<br>See [Js Insert All Pages Except](#except-7bfe85) below.

<a id="rules-ad3671"></a>&#x2022; [`js_insertion_rules`](#rules-ad3671) - Optional Block<br>JavaScript Custom Insertion Rules. This defines custom JavaScript insertion rules for Client-Side Defense Policy<br>See [Js Insertion Rules](#rules-ad3671) below.

#### Client Side Defense Policy Js Insert All Pages Except

<a id="deep-357348"></a>Deeply nested **Except** block collapsed for readability.

#### Client Side Defense Policy Js Insert All Pages Except Exclude List

<a id="deep-e071ac"></a>Deeply nested **List** block collapsed for readability.

#### Client Side Defense Policy Js Insert All Pages Except Exclude List Domain

<a id="deep-05664f"></a>Deeply nested **Domain** block collapsed for readability.

#### Client Side Defense Policy Js Insert All Pages Except Exclude List Metadata

<a id="deep-ec44bf"></a>Deeply nested **Metadata** block collapsed for readability.

#### Client Side Defense Policy Js Insert All Pages Except Exclude List Path

<a id="deep-eba2d4"></a>Deeply nested **Path** block collapsed for readability.

#### Client Side Defense Policy Js Insertion Rules

A [`js_insertion_rules`](#rules-ad3671) block (within [`client_side_defense.policy`](#client-side-defense-policy)) supports the following:

<a id="list-dfecb6"></a>&#x2022; [`exclude_list`](#list-dfecb6) - Optional Block<br>Exclude Paths. Optional JavaScript insertions exclude list of domain and path matchers<br>See [Exclude List](#list-dfecb6) below.

<a id="rules-6276bc"></a>&#x2022; [`rules`](#rules-6276bc) - Optional Block<br>JavaScript Insertions. Required list of pages to insert Client-Side Defense client JavaScript<br>See [Rules](#rules-6276bc) below.

#### Client Side Defense Policy Js Insertion Rules Exclude List

<a id="deep-1639b0"></a>Deeply nested **List** block collapsed for readability.

#### Client Side Defense Policy Js Insertion Rules Exclude List Domain

<a id="deep-949064"></a>Deeply nested **Domain** block collapsed for readability.

#### Client Side Defense Policy Js Insertion Rules Exclude List Metadata

<a id="deep-3af08a"></a>Deeply nested **Metadata** block collapsed for readability.

#### Client Side Defense Policy Js Insertion Rules Exclude List Path

<a id="deep-44c104"></a>Deeply nested **Path** block collapsed for readability.

#### Client Side Defense Policy Js Insertion Rules Rules

<a id="deep-ad681e"></a>Deeply nested **Rules** block collapsed for readability.

#### Client Side Defense Policy Js Insertion Rules Rules Domain

<a id="deep-15e025"></a>Deeply nested **Domain** block collapsed for readability.

#### Client Side Defense Policy Js Insertion Rules Rules Metadata

<a id="deep-6c411a"></a>Deeply nested **Metadata** block collapsed for readability.

#### Client Side Defense Policy Js Insertion Rules Rules Path

<a id="deep-880dda"></a>Deeply nested **Path** block collapsed for readability.

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

<a id="data-guard-rules-path-prefix"></a>&#x2022; [`prefix`](#data-guard-rules-path-prefix) - Optional String<br>Prefix. Path prefix to match (e.g. The value / will match on all paths)

<a id="data-guard-rules-path-regex"></a>&#x2022; [`regex`](#data-guard-rules-path-regex) - Optional String<br>Regex. Regular expression of path match (e.g. The value .* will match on all paths)

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

<a id="deep-4fa70d"></a>Deeply nested **List** block collapsed for readability.

#### DDOS Mitigation Rules DDOS Client Source Ja4 TLS Fingerprint Matcher

<a id="deep-ae859d"></a>Deeply nested **Matcher** block collapsed for readability.

#### DDOS Mitigation Rules DDOS Client Source TLS Fingerprint Matcher

<a id="deep-3d0bac"></a>Deeply nested **Matcher** block collapsed for readability.

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

<a id="default-pool-advanced-options"></a>&#x2022; [`advanced_options`](#default-pool-advanced-options) - Optional Block<br>Origin Pool Advanced OPTIONS. Configure Advanced OPTIONS for origin pool<br>See [Advanced Options](#default-pool-advanced-options) below.

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

<a id="subsets-5a741c"></a>&#x2022; [`enable_subsets`](#subsets-5a741c) - Optional Block<br>Origin Pool Subset OPTIONS. Configure subset OPTIONS for origin pool<br>See [Enable Subsets](#subsets-5a741c) below.

<a id="config-a0bc3c"></a>&#x2022; [`http1_config`](#config-a0bc3c) - Optional Block<br>HTTP/1.1 Protocol OPTIONS. HTTP/1.1 Protocol OPTIONS for upstream connections<br>See [Http1 Config](#config-a0bc3c) below.

<a id="options-fc9fd8"></a>&#x2022; [`http2_options`](#options-fc9fd8) - Optional Block<br>Http2 Protocol OPTIONS. Http2 Protocol OPTIONS for upstream connections<br>See [Http2 Options](#options-fc9fd8) below.

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

<a id="deep-d44dd0"></a>Deeply nested **Subset** block collapsed for readability.

#### Default Pool Advanced Options Enable Subsets Endpoint Subsets

<a id="deep-b9a0f0"></a>Deeply nested **Subsets** block collapsed for readability.

#### Default Pool Advanced Options Http1 Config

A [`http1_config`](#config-a0bc3c) block (within [`default_pool.advanced_options`](#default-pool-advanced-options)) supports the following:

<a id="transformation-cc4211"></a>&#x2022; [`header_transformation`](#transformation-cc4211) - Optional Block<br>Header Transformation. Header Transformation OPTIONS for HTTP/1.1 request/response headers<br>See [Header Transformation](#transformation-cc4211) below.

#### Default Pool Advanced Options Http1 Config Header Transformation

<a id="deep-62ae42"></a>Deeply nested **Transformation** block collapsed for readability.

#### Default Pool Advanced Options Http2 Options

A [`http2_options`](#options-fc9fd8) block (within [`default_pool.advanced_options`](#default-pool-advanced-options)) supports the following:

<a id="enabled-4e1198"></a>&#x2022; [`enabled`](#enabled-4e1198) - Optional Bool<br>HTTP2 Enabled. Enable/disable HTTP2 Protocol for upstream connections

#### Default Pool Advanced Options Outlier Detection

An [`outlier_detection`](#detection-c89e70) block (within [`default_pool.advanced_options`](#default-pool-advanced-options)) supports the following:

<a id="time-fd8e40"></a>&#x2022; [`base_ejection_time`](#time-fd8e40) - Optional Number  Defaults to `30000ms`  Specified in milliseconds<br>Base Ejection Time. The base time that a host is ejected for. The real time is equal to the base time multiplied by the number of times the host has been ejected. This causes hosts to GET ejected for longer periods if they continue to fail

<a id="5xx-1c4ec1"></a>&#x2022; [`consecutive_5xx`](#5xx-1c4ec1) - Optional Number  Defaults to `5`<br>Consecutive 5xx Count. If an upstream endpoint returns some number of consecutive 5xx, it will be ejected. Note that in this case a 5xx means an actual 5xx respond code, or an event that would cause the HTTP router to return one on the upstream’s behalf(reset, connection failure, etc.) consecutive_5xx indicates the number of consecutive 5xx responses required before a consecutive 5xx ejection occurs

<a id="failure-b19342"></a>&#x2022; [`consecutive_gateway_failure`](#failure-b19342) - Optional Number  Defaults to `5`<br>Consecutive Gateway Failure. If an upstream endpoint returns some number of consecutive “gateway errors” (502, 503 or 504 status code), it will be ejected. Note that this includes events that would cause the HTTP router to return one of these status codes on the upstream’s behalf (reset, connection failure, etc.). Consecutive_gateway_failure indicates the number of consecutive gateway failures before a consecutive gateway failure ejection occurs

<a id="interval-834cfc"></a>&#x2022; [`interval`](#interval-834cfc) - Optional Number  Defaults to `10000ms`  Specified in milliseconds<br>Interval. The time interval between ejection analysis sweeps. This can result in both new ejections as well as endpoints being returned to service

<a id="percent-9a52cc"></a>&#x2022; [`max_ejection_percent`](#percent-9a52cc) - Optional Number  Defaults to `10%`<br>Max Ejection Percentage. The maximum % of an upstream cluster that can be ejected due to outlier detection. but will eject at least one host regardless of the value

#### Default Pool Healthcheck

A [`healthcheck`](#default-pool-healthcheck) block (within [`default_pool`](#default-pool)) supports the following:

<a id="default-pool-healthcheck-name"></a>&#x2022; [`name`](#default-pool-healthcheck-name) - Optional String<br>Name. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name

<a id="default-pool-healthcheck-namespace"></a>&#x2022; [`namespace`](#default-pool-healthcheck-namespace) - Optional String<br>Namespace. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace

<a id="default-pool-healthcheck-tenant"></a>&#x2022; [`tenant`](#default-pool-healthcheck-tenant) - Optional String<br>Tenant. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant

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

<a id="pool-b708db"></a>&#x2022; [`snat_pool`](#pool-b708db) - Optional Block<br>SNAT Pool. SNAT Pool configuration<br>See [Snat Pool](#pool-b708db) below.

#### Default Pool Origin Servers Consul Service Site Locator

<a id="deep-530a54"></a>Deeply nested **Locator** block collapsed for readability.

#### Default Pool Origin Servers Consul Service Site Locator Site

<a id="deep-ccdb22"></a>Deeply nested **Site** block collapsed for readability.

#### Default Pool Origin Servers Consul Service Site Locator Virtual Site

<a id="deep-5c93cd"></a>Deeply nested **Site** block collapsed for readability.

#### Default Pool Origin Servers Consul Service Snat Pool

<a id="deep-4dada7"></a>Deeply nested **Pool** block collapsed for readability.

#### Default Pool Origin Servers Consul Service Snat Pool Snat Pool

<a id="deep-51125d"></a>Deeply nested **Pool** block collapsed for readability.

#### Default Pool Origin Servers Custom Endpoint Object

A [`custom_endpoint_object`](#object-12dd7f) block (within [`default_pool.origin_servers`](#default-pool-origin-servers)) supports the following:

<a id="endpoint-952bdd"></a>&#x2022; [`endpoint`](#endpoint-952bdd) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Endpoint](#endpoint-952bdd) below.

#### Default Pool Origin Servers Custom Endpoint Object Endpoint

<a id="deep-842b8b"></a>Deeply nested **Endpoint** block collapsed for readability.

#### Default Pool Origin Servers K8S Service

A [`k8s_service`](#default-pool-origin-servers-k8s-service) block (within [`default_pool.origin_servers`](#default-pool-origin-servers)) supports the following:

<a id="network-dfbf17"></a>&#x2022; [`inside_network`](#network-dfbf17) - Optional Block<br>Enable this option

<a id="network-d1b956"></a>&#x2022; [`outside_network`](#network-d1b956) - Optional Block<br>Enable this option

<a id="protocol-ffcd27"></a>&#x2022; [`protocol`](#protocol-ffcd27) - Optional String  Defaults to `PROTOCOL_TCP`<br>Possible values are `PROTOCOL_TCP`, `PROTOCOL_UDP`<br>[Enum: PROTOCOL_TCP|PROTOCOL_UDP] Protocol Type. Type of protocol - PROTOCOL_TCP: TCP - PROTOCOL_UDP: UDP

<a id="name-c77159"></a>&#x2022; [`service_name`](#name-c77159) - Optional String<br>Service Name. K8S service name of the origin server will be listed, including the namespace and cluster-ID. For vK8s services, you need to enter a string with the format servicename.namespace:cluster-ID. If the servicename is 'frontend', namespace is 'speedtest' and cluster-ID is 'prod', then you will enter 'frontend.speedtest:prod'. Both namespace and cluster-ID are optional

<a id="locator-8a5921"></a>&#x2022; [`site_locator`](#locator-8a5921) - Optional Block<br>Site or Virtual Site. This message defines a reference to a site or virtual site object<br>See [Site Locator](#locator-8a5921) below.

<a id="pool-0640ac"></a>&#x2022; [`snat_pool`](#pool-0640ac) - Optional Block<br>SNAT Pool. SNAT Pool configuration<br>See [Snat Pool](#pool-0640ac) below.

<a id="networks-03f764"></a>&#x2022; [`vk8s_networks`](#networks-03f764) - Optional Block<br>Enable this option

#### Default Pool Origin Servers K8S Service Site Locator

<a id="deep-2de3d2"></a>Deeply nested **Locator** block collapsed for readability.

#### Default Pool Origin Servers K8S Service Site Locator Site

<a id="deep-3efcb6"></a>Deeply nested **Site** block collapsed for readability.

#### Default Pool Origin Servers K8S Service Site Locator Virtual Site

<a id="deep-8b29b6"></a>Deeply nested **Site** block collapsed for readability.

#### Default Pool Origin Servers K8S Service Snat Pool

<a id="deep-b309ef"></a>Deeply nested **Pool** block collapsed for readability.

#### Default Pool Origin Servers K8S Service Snat Pool Snat Pool

<a id="deep-37c94b"></a>Deeply nested **Pool** block collapsed for readability.

#### Default Pool Origin Servers Private IP

A [`private_ip`](#default-pool-origin-servers-private-ip) block (within [`default_pool.origin_servers`](#default-pool-origin-servers)) supports the following:

<a id="network-704c7d"></a>&#x2022; [`inside_network`](#network-704c7d) - Optional Block<br>Enable this option

<a id="ip-ip-4a696b"></a>&#x2022; [`ip`](#ip-ip-4a696b) - Optional String<br>IP. Private IPv4 address

<a id="network-f44165"></a>&#x2022; [`outside_network`](#network-f44165) - Optional Block<br>Enable this option

<a id="segment-735aa1"></a>&#x2022; [`segment`](#segment-735aa1) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Segment](#segment-735aa1) below.

<a id="locator-1137c8"></a>&#x2022; [`site_locator`](#locator-1137c8) - Optional Block<br>Site or Virtual Site. This message defines a reference to a site or virtual site object<br>See [Site Locator](#locator-1137c8) below.

<a id="pool-916f33"></a>&#x2022; [`snat_pool`](#pool-916f33) - Optional Block<br>SNAT Pool. SNAT Pool configuration<br>See [Snat Pool](#pool-916f33) below.

#### Default Pool Origin Servers Private IP Segment

A [`segment`](#segment-735aa1) block (within [`default_pool.origin_servers.private_ip`](#default-pool-origin-servers-private-ip)) supports the following:

<a id="name-a7f063"></a>&#x2022; [`name`](#name-a7f063) - Optional String<br>Name. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name

<a id="namespace-4c09cc"></a>&#x2022; [`namespace`](#namespace-4c09cc) - Optional String<br>Namespace. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace

<a id="tenant-f972d4"></a>&#x2022; [`tenant`](#tenant-f972d4) - Optional String<br>Tenant. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant

#### Default Pool Origin Servers Private IP Site Locator

<a id="deep-ef72fc"></a>Deeply nested **Locator** block collapsed for readability.

#### Default Pool Origin Servers Private IP Site Locator Site

<a id="deep-5ff672"></a>Deeply nested **Site** block collapsed for readability.

#### Default Pool Origin Servers Private IP Site Locator Virtual Site

<a id="deep-9941ea"></a>Deeply nested **Site** block collapsed for readability.

#### Default Pool Origin Servers Private IP Snat Pool

<a id="deep-85b65f"></a>Deeply nested **Pool** block collapsed for readability.

#### Default Pool Origin Servers Private IP Snat Pool Snat Pool

<a id="deep-d9ddae"></a>Deeply nested **Pool** block collapsed for readability.

#### Default Pool Origin Servers Private Name

A [`private_name`](#name-966ae3) block (within [`default_pool.origin_servers`](#default-pool-origin-servers)) supports the following:

<a id="name-8a8021"></a>&#x2022; [`dns_name`](#name-8a8021) - Optional String<br>DNS Name. DNS Name

<a id="network-e9e813"></a>&#x2022; [`inside_network`](#network-e9e813) - Optional Block<br>Enable this option

<a id="network-873dcb"></a>&#x2022; [`outside_network`](#network-873dcb) - Optional Block<br>Enable this option

<a id="interval-615002"></a>&#x2022; [`refresh_interval`](#interval-615002) - Optional Number<br>DNS Refresh Interval. Interval for DNS refresh in seconds. Max value is 7 days as per `HTTPS://datatracker.ietf.org/doc/HTML/rfc8767.`

<a id="segment-8fe482"></a>&#x2022; [`segment`](#segment-8fe482) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Segment](#segment-8fe482) below.

<a id="locator-3db1ee"></a>&#x2022; [`site_locator`](#locator-3db1ee) - Optional Block<br>Site or Virtual Site. This message defines a reference to a site or virtual site object<br>See [Site Locator](#locator-3db1ee) below.

<a id="pool-6d884b"></a>&#x2022; [`snat_pool`](#pool-6d884b) - Optional Block<br>SNAT Pool. SNAT Pool configuration<br>See [Snat Pool](#pool-6d884b) below.

#### Default Pool Origin Servers Private Name Segment

A [`segment`](#segment-8fe482) block (within [`default_pool.origin_servers.private_name`](#name-966ae3)) supports the following:

<a id="name-88ec1c"></a>&#x2022; [`name`](#name-88ec1c) - Optional String<br>Name. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name

<a id="namespace-2f15e9"></a>&#x2022; [`namespace`](#namespace-2f15e9) - Optional String<br>Namespace. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace

<a id="tenant-1fee5d"></a>&#x2022; [`tenant`](#tenant-1fee5d) - Optional String<br>Tenant. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant

#### Default Pool Origin Servers Private Name Site Locator

<a id="deep-47d5a9"></a>Deeply nested **Locator** block collapsed for readability.

#### Default Pool Origin Servers Private Name Site Locator Site

<a id="deep-3e8686"></a>Deeply nested **Site** block collapsed for readability.

#### Default Pool Origin Servers Private Name Site Locator Virtual Site

<a id="deep-27cdc5"></a>Deeply nested **Site** block collapsed for readability.

#### Default Pool Origin Servers Private Name Snat Pool

<a id="deep-b3cd47"></a>Deeply nested **Pool** block collapsed for readability.

#### Default Pool Origin Servers Private Name Snat Pool Snat Pool

<a id="deep-82ea4a"></a>Deeply nested **Pool** block collapsed for readability.

#### Default Pool Origin Servers Public IP

A [`public_ip`](#default-pool-origin-servers-public-ip) block (within [`default_pool.origin_servers`](#default-pool-origin-servers)) supports the following:

<a id="ip-ip-dc19b2"></a>&#x2022; [`ip`](#ip-ip-dc19b2) - Optional String<br>Public IPv4. Public IPv4 address

#### Default Pool Origin Servers Public Name

A [`public_name`](#default-pool-origin-servers-public-name) block (within [`default_pool.origin_servers`](#default-pool-origin-servers)) supports the following:

<a id="name-7c3f95"></a>&#x2022; [`dns_name`](#name-7c3f95) - Optional String<br>DNS Name. DNS Name

<a id="interval-ac5170"></a>&#x2022; [`refresh_interval`](#interval-ac5170) - Optional Number<br>DNS Refresh Interval. Interval for DNS refresh in seconds. Max value is 7 days as per `HTTPS://datatracker.ietf.org/doc/HTML/rfc8767.`

#### Default Pool Origin Servers Vn Private IP

A [`vn_private_ip`](#private-ip-532445) block (within [`default_pool.origin_servers`](#default-pool-origin-servers)) supports the following:

<a id="ip-ip-c5d639"></a>&#x2022; [`ip`](#ip-ip-c5d639) - Optional String<br>IPv4. IPv4 address

<a id="network-56a203"></a>&#x2022; [`virtual_network`](#network-56a203) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Virtual Network](#network-56a203) below.

#### Default Pool Origin Servers Vn Private IP Virtual Network

<a id="deep-d6a69f"></a>Deeply nested **Network** block collapsed for readability.

#### Default Pool Origin Servers Vn Private Name

A [`vn_private_name`](#name-4a1747) block (within [`default_pool.origin_servers`](#default-pool-origin-servers)) supports the following:

<a id="name-10334e"></a>&#x2022; [`dns_name`](#name-10334e) - Optional String<br>DNS Name. DNS Name

<a id="network-bbece7"></a>&#x2022; [`private_network`](#network-bbece7) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Private Network](#network-bbece7) below.

#### Default Pool Origin Servers Vn Private Name Private Network

<a id="deep-f6c767"></a>Deeply nested **Network** block collapsed for readability.

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

<a id="default-pool-use-tls-tls-config"></a>&#x2022; [`tls_config`](#default-pool-use-tls-tls-config) - Optional Block<br>TLS Config. This defines various OPTIONS to configure TLS configuration parameters<br>See [TLS Config](#default-pool-use-tls-tls-config) below.

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

<a id="deep-d46d52"></a>Deeply nested **Security** block collapsed for readability.

#### Default Pool Use TLS Use mTLS

An [`use_mtls`](#default-pool-use-tls-use-mtls) block (within [`default_pool.use_tls`](#default-pool-use-tls)) supports the following:

<a id="certificates-5055f8"></a>&#x2022; [`tls_certificates`](#certificates-5055f8) - Optional Block<br>mTLS Client Certificate. mTLS Client Certificate<br>See [TLS Certificates](#certificates-5055f8) below.

#### Default Pool Use TLS Use mTLS TLS Certificates

<a id="deep-eafd3c"></a>Deeply nested **Certificates** block collapsed for readability.

#### Default Pool Use TLS Use mTLS TLS Certificates Custom Hash Algorithms

<a id="deep-3f8b43"></a>Deeply nested **Algorithms** block collapsed for readability.

#### Default Pool Use TLS Use mTLS TLS Certificates Private Key

<a id="deep-b960e8"></a>Deeply nested **Key** block collapsed for readability.

#### Default Pool Use TLS Use mTLS TLS Certificates Private Key Blindfold Secret Info

<a id="deep-1fc68e"></a>Deeply nested **Info** block collapsed for readability.

#### Default Pool Use TLS Use mTLS TLS Certificates Private Key Clear Secret Info

<a id="deep-724bf8"></a>Deeply nested **Info** block collapsed for readability.

#### Default Pool Use TLS Use mTLS Obj

An [`use_mtls_obj`](#default-pool-use-tls-use-mtls-obj) block (within [`default_pool.use_tls`](#default-pool-use-tls)) supports the following:

<a id="default-pool-use-tls-use-mtls-obj-name"></a>&#x2022; [`name`](#default-pool-use-tls-use-mtls-obj-name) - Optional String<br>Name. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name

<a id="namespace-5c1b91"></a>&#x2022; [`namespace`](#namespace-5c1b91) - Optional String<br>Namespace. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace

<a id="tenant-0b54b4"></a>&#x2022; [`tenant`](#tenant-0b54b4) - Optional String<br>Tenant. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant

#### Default Pool Use TLS Use Server Verification

An [`use_server_verification`](#verification-388853) block (within [`default_pool.use_tls`](#default-pool-use-tls)) supports the following:

<a id="trusted-ca-1da800"></a>&#x2022; [`trusted_ca`](#trusted-ca-1da800) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Trusted CA](#trusted-ca-1da800) below.

<a id="url-d5fda5"></a>&#x2022; [`trusted_ca_url`](#url-d5fda5) - Optional String<br>Inline Root CA Certificate (legacy). Upload a Root CA Certificate specifically for this Origin Pool for verification of server's certificate

#### Default Pool Use TLS Use Server Verification Trusted CA

<a id="deep-bae0f5"></a>Deeply nested **CA** block collapsed for readability.

#### Default Pool View Internal

A [`view_internal`](#default-pool-view-internal) block (within [`default_pool`](#default-pool)) supports the following:

<a id="default-pool-view-internal-name"></a>&#x2022; [`name`](#default-pool-view-internal-name) - Optional String<br>Name. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name

<a id="default-pool-view-internal-namespace"></a>&#x2022; [`namespace`](#default-pool-view-internal-namespace) - Optional String<br>Namespace. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace

<a id="default-pool-view-internal-tenant"></a>&#x2022; [`tenant`](#default-pool-view-internal-tenant) - Optional String<br>Tenant. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant

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

<a id="default-pool-list-pools-cluster-name"></a>&#x2022; [`name`](#default-pool-list-pools-cluster-name) - Optional String<br>Name. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name

<a id="namespace-4d4aed"></a>&#x2022; [`namespace`](#namespace-4d4aed) - Optional String<br>Namespace. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace

<a id="default-pool-list-pools-cluster-tenant"></a>&#x2022; [`tenant`](#default-pool-list-pools-cluster-tenant) - Optional String<br>Tenant. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant

#### Default Pool List Pools Pool

A [`pool`](#default-pool-list-pools-pool) block (within [`default_pool_list.pools`](#default-pool-list-pools)) supports the following:

<a id="default-pool-list-pools-pool-name"></a>&#x2022; [`name`](#default-pool-list-pools-pool-name) - Optional String<br>Name. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name

<a id="default-pool-list-pools-pool-namespace"></a>&#x2022; [`namespace`](#default-pool-list-pools-pool-namespace) - Optional String<br>Namespace. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace

<a id="default-pool-list-pools-pool-tenant"></a>&#x2022; [`tenant`](#default-pool-list-pools-pool-tenant) - Optional String<br>Tenant. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant

#### Default Route Pools

A [`default_route_pools`](#default-route-pools) block supports the following:

<a id="default-route-pools-cluster"></a>&#x2022; [`cluster`](#default-route-pools-cluster) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Cluster](#default-route-pools-cluster) below.

<a id="default-route-pools-endpoint-subsets"></a>&#x2022; [`endpoint_subsets`](#default-route-pools-endpoint-subsets) - Optional Block<br>Origin Servers Subsets. Upstream origin pool may be configured to divide its origin servers into subsets based on metadata attached to the origin servers. Routes may then specify the metadata that a endpoint must match in order to be selected by the load balancer For origin servers which are discovered in K8S or Consul cluster, the label of the service is merged with endpoint's labels. In case of Consul, the label is derived from the 'Tag' field. For labels that are common between configured endpoint and discovered service, labels from discovered service takes precedence. List of key-value pairs that will be used as matching metadata. Only those origin servers of upstream origin pool which match this metadata will be selected for load balancing

<a id="default-route-pools-pool"></a>&#x2022; [`pool`](#default-route-pools-pool) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Pool](#default-route-pools-pool) below.

<a id="default-route-pools-priority"></a>&#x2022; [`priority`](#default-route-pools-priority) - Optional Number<br>Priority. Priority of this origin pool, valid only with multiple origin pools. Value of 0 will make the pool as lowest priority origin pool Priority of 1 means highest priority and is considered active. When active origin pool is not available, lower priority origin pools are made active as per the increasing priority

<a id="default-route-pools-weight"></a>&#x2022; [`weight`](#default-route-pools-weight) - Optional Number<br>Weight. Weight of this origin pool, valid only with multiple origin pool. Value of 0 will disable the pool

#### Default Route Pools Cluster

A [`cluster`](#default-route-pools-cluster) block (within [`default_route_pools`](#default-route-pools)) supports the following:

<a id="default-route-pools-cluster-name"></a>&#x2022; [`name`](#default-route-pools-cluster-name) - Optional String<br>Name. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name

<a id="default-route-pools-cluster-namespace"></a>&#x2022; [`namespace`](#default-route-pools-cluster-namespace) - Optional String<br>Namespace. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace

<a id="default-route-pools-cluster-tenant"></a>&#x2022; [`tenant`](#default-route-pools-cluster-tenant) - Optional String<br>Tenant. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant

#### Default Route Pools Pool

A [`pool`](#default-route-pools-pool) block (within [`default_route_pools`](#default-route-pools)) supports the following:

<a id="default-route-pools-pool-name"></a>&#x2022; [`name`](#default-route-pools-pool-name) - Optional String<br>Name. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name

<a id="default-route-pools-pool-namespace"></a>&#x2022; [`namespace`](#default-route-pools-pool-namespace) - Optional String<br>Namespace. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace

<a id="default-route-pools-pool-tenant"></a>&#x2022; [`tenant`](#default-route-pools-pool-tenant) - Optional String<br>Tenant. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant

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

<a id="deep-038178"></a>Deeply nested **Config** block collapsed for readability.

#### Enable API Discovery API Crawler API Crawler Config Domains

<a id="deep-2dedc6"></a>Deeply nested **Domains** block collapsed for readability.

#### Enable API Discovery API Crawler API Crawler Config Domains Simple Login

<a id="deep-f94bfe"></a>Deeply nested **Login** block collapsed for readability.

#### Enable API Discovery API Crawler API Crawler Config Domains Simple Login Password

<a id="deep-dcac2b"></a>Deeply nested **Password** block collapsed for readability.

#### Enable API Discovery API Crawler API Crawler Config Domains Simple Login Password Blindfold Secret Info

<a id="deep-bb1337"></a>Deeply nested **Info** block collapsed for readability.

#### Enable API Discovery API Crawler API Crawler Config Domains Simple Login Password Clear Secret Info

<a id="deep-790308"></a>Deeply nested **Info** block collapsed for readability.

#### Enable API Discovery API Discovery From Code Scan

<a id="deep-65081a"></a>Deeply nested **Scan** block collapsed for readability.

#### Enable API Discovery API Discovery From Code Scan Code Base Integrations

<a id="deep-655889"></a>Deeply nested **Integrations** block collapsed for readability.

#### Enable API Discovery API Discovery From Code Scan Code Base Integrations Code Base Integration

<a id="deep-f966a5"></a>Deeply nested **Integration** block collapsed for readability.

#### Enable API Discovery API Discovery From Code Scan Code Base Integrations Selected Repos

<a id="deep-24c79f"></a>Deeply nested **Repos** block collapsed for readability.

#### Enable API Discovery Custom API Auth Discovery

A [`custom_api_auth_discovery`](#discovery-54db29) block (within [`enable_api_discovery`](#enable-api-discovery)) supports the following:

<a id="ref-a70328"></a>&#x2022; [`api_discovery_ref`](#ref-a70328) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [API Discovery Ref](#ref-a70328) below.

#### Enable API Discovery Custom API Auth Discovery API Discovery Ref

<a id="deep-af0566"></a>Deeply nested **Ref** block collapsed for readability.

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

<a id="name-3a9364"></a>&#x2022; [`name`](#name-3a9364) - Optional String<br>Name. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name

<a id="namespace-38ef32"></a>&#x2022; [`namespace`](#namespace-38ef32) - Optional String<br>Namespace. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace

<a id="tenant-78def2"></a>&#x2022; [`tenant`](#tenant-78def2) - Optional String<br>Tenant. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant

#### Enable IP Reputation

An [`enable_ip_reputation`](#enable-ip-reputation) block supports the following:

<a id="categories-bb360f"></a>&#x2022; [`ip_threat_categories`](#categories-bb360f) - Optional List  Defaults to `SPAM_SOURCES`<br>See [IP Threat Categories](#common-ip-threat-categories)<br>[Enum: SPAM_SOURCES|WINDOWS_EXPLOITS|WEB_ATTACKS|BOTNETS|SCANNERS|REPUTATION|PHISHING|PROXY|MOBILE_THREATS|TOR_PROXY|DENIAL_OF_SERVICE|NETWORK] List of IP Threat Categories to choose. If the source IP matches on atleast one of the enabled IP threat categories, the request will be denied

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

<a id="https-coalescing-options"></a>&#x2022; [`coalescing_options`](#https-coalescing-options) - Optional Block<br>TLS Coalescing OPTIONS. TLS connection coalescing configuration (not compatible with mTLS)<br>See [Coalescing Options](#https-coalescing-options) below.

<a id="https-connection-idle-timeout"></a>&#x2022; [`connection_idle_timeout`](#https-connection-idle-timeout) - Optional Number  Defaults to `2`  Specified in milliseconds<br>Connection Idle Timeout. The idle timeout for downstream connections. The idle timeout is defined as the period in which there are no active requests. When the idle timeout is reached the connection will be closed. Note that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive.  The minutes

<a id="https-default-header"></a>&#x2022; [`default_header`](#https-default-header) - Optional Block<br>Enable this option

<a id="https-default-loadbalancer"></a>&#x2022; [`default_loadbalancer`](#https-default-loadbalancer) - Optional Block<br>Enable this option

<a id="https-disable-path-normalize"></a>&#x2022; [`disable_path_normalize`](#https-disable-path-normalize) - Optional Block<br>Enable this option

<a id="https-enable-path-normalize"></a>&#x2022; [`enable_path_normalize`](#https-enable-path-normalize) - Optional Block<br>Enable this option

<a id="https-http-protocol-options"></a>&#x2022; [`http_protocol_options`](#https-http-protocol-options) - Optional Block<br>HTTP Protocol Configuration OPTIONS. HTTP protocol configuration OPTIONS for downstream connections<br>See [HTTP Protocol Options](#https-http-protocol-options) below.

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

<a id="only-46f3ca"></a>&#x2022; [`http_protocol_enable_v1_only`](#only-46f3ca) - Optional Block<br>HTTP/1.1 Protocol OPTIONS. HTTP/1.1 Protocol OPTIONS for downstream connections<br>See [HTTP Protocol Enable V1 Only](#only-46f3ca) below.

<a id="v1-v2-6f8b9b"></a>&#x2022; [`http_protocol_enable_v1_v2`](#v1-v2-6f8b9b) - Optional Block<br>Enable this option

<a id="only-5cefb3"></a>&#x2022; [`http_protocol_enable_v2_only`](#only-5cefb3) - Optional Block<br>Enable this option

#### HTTPS HTTP Protocol Options HTTP Protocol Enable V1 Only

<a id="deep-2a12f6"></a>Deeply nested **Only** block collapsed for readability.

#### HTTPS HTTP Protocol Options HTTP Protocol Enable V1 Only Header Transformation

<a id="deep-4134cf"></a>Deeply nested **Transformation** block collapsed for readability.

#### HTTPS TLS Cert Params

A [`tls_cert_params`](#https-tls-cert-params) block (within [`https`](#https)) supports the following:

<a id="https-tls-cert-params-certificates"></a>&#x2022; [`certificates`](#https-tls-cert-params-certificates) - Optional Block<br>Certificates. Select one or more certificates with any domain names<br>See [Certificates](#https-tls-cert-params-certificates) below.

<a id="https-tls-cert-params-no-mtls"></a>&#x2022; [`no_mtls`](#https-tls-cert-params-no-mtls) - Optional Block<br>Enable this option

<a id="https-tls-cert-params-tls-config"></a>&#x2022; [`tls_config`](#https-tls-cert-params-tls-config) - Optional Block<br>TLS Config. This defines various OPTIONS to configure TLS configuration parameters<br>See [TLS Config](#https-tls-cert-params-tls-config) below.

<a id="https-tls-cert-params-use-mtls"></a>&#x2022; [`use_mtls`](#https-tls-cert-params-use-mtls) - Optional Block<br>Clients TLS validation context. Validation context for downstream client TLS connections<br>See [Use mTLS](#https-tls-cert-params-use-mtls) below.

#### HTTPS TLS Cert Params Certificates

A [`certificates`](#https-tls-cert-params-certificates) block (within [`https.tls_cert_params`](#https-tls-cert-params)) supports the following:

<a id="https-tls-cert-params-certificates-name"></a>&#x2022; [`name`](#https-tls-cert-params-certificates-name) - Optional String<br>Name. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name

<a id="namespace-74e8ce"></a>&#x2022; [`namespace`](#namespace-74e8ce) - Optional String<br>Namespace. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace

<a id="tenant-7c270a"></a>&#x2022; [`tenant`](#tenant-7c270a) - Optional String<br>Tenant. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant

#### HTTPS TLS Cert Params TLS Config

A [`tls_config`](#https-tls-cert-params-tls-config) block (within [`https.tls_cert_params`](#https-tls-cert-params)) supports the following:

<a id="security-6452ce"></a>&#x2022; [`custom_security`](#security-6452ce) - Optional Block<br>Custom Ciphers. This defines TLS protocol config including min/max versions and allowed ciphers<br>See [Custom Security](#security-6452ce) below.

<a id="security-b6db5a"></a>&#x2022; [`default_security`](#security-b6db5a) - Optional Block<br>Enable this option

<a id="security-cbe12e"></a>&#x2022; [`low_security`](#security-cbe12e) - Optional Block<br>Enable this option

<a id="security-e410e3"></a>&#x2022; [`medium_security`](#security-e410e3) - Optional Block<br>Enable this option

#### HTTPS TLS Cert Params TLS Config Custom Security

<a id="deep-0971e8"></a>Deeply nested **Security** block collapsed for readability.

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

<a id="https-tls-cert-params-use-mtls-crl-name"></a>&#x2022; [`name`](#https-tls-cert-params-use-mtls-crl-name) - Optional String<br>Name. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name

<a id="namespace-4349dd"></a>&#x2022; [`namespace`](#namespace-4349dd) - Optional String<br>Namespace. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace

<a id="tenant-7e4839"></a>&#x2022; [`tenant`](#tenant-7e4839) - Optional String<br>Tenant. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant

#### HTTPS TLS Cert Params Use mTLS Trusted CA

<a id="deep-d14eb7"></a>Deeply nested **CA** block collapsed for readability.

#### HTTPS TLS Cert Params Use mTLS Xfcc Options

<a id="deep-9e1e47"></a>Deeply nested **Options** block collapsed for readability.

#### HTTPS TLS Parameters

A [`tls_parameters`](#https-tls-parameters) block (within [`https`](#https)) supports the following:

<a id="https-tls-parameters-no-mtls"></a>&#x2022; [`no_mtls`](#https-tls-parameters-no-mtls) - Optional Block<br>Enable this option

<a id="https-tls-parameters-tls-certificates"></a>&#x2022; [`tls_certificates`](#https-tls-parameters-tls-certificates) - Optional Block<br>TLS Certificates. Users can add one or more certificates that share the same set of domains. For example, domain.com and \*.domain.com - but use different signature algorithms<br>See [TLS Certificates](#https-tls-parameters-tls-certificates) below.

<a id="https-tls-parameters-tls-config"></a>&#x2022; [`tls_config`](#https-tls-parameters-tls-config) - Optional Block<br>TLS Config. This defines various OPTIONS to configure TLS configuration parameters<br>See [TLS Config](#https-tls-parameters-tls-config) below.

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

<a id="deep-adc018"></a>Deeply nested **Algorithms** block collapsed for readability.

#### HTTPS TLS Parameters TLS Certificates Private Key

A [`private_key`](#key-372460) block (within [`https.tls_parameters.tls_certificates`](#https-tls-parameters-tls-certificates)) supports the following:

<a id="info-0c9fbe"></a>&#x2022; [`blindfold_secret_info`](#info-0c9fbe) - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#info-0c9fbe) below.

<a id="info-556650"></a>&#x2022; [`clear_secret_info`](#info-556650) - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#info-556650) below.

#### HTTPS TLS Parameters TLS Certificates Private Key Blindfold Secret Info

<a id="deep-de8238"></a>Deeply nested **Info** block collapsed for readability.

#### HTTPS TLS Parameters TLS Certificates Private Key Clear Secret Info

<a id="deep-f7af42"></a>Deeply nested **Info** block collapsed for readability.

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

<a id="https-tls-parameters-use-mtls-crl-name"></a>&#x2022; [`name`](#https-tls-parameters-use-mtls-crl-name) - Optional String<br>Name. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name

<a id="namespace-ae83ff"></a>&#x2022; [`namespace`](#namespace-ae83ff) - Optional String<br>Namespace. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace

<a id="tenant-08da33"></a>&#x2022; [`tenant`](#tenant-08da33) - Optional String<br>Tenant. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant

#### HTTPS TLS Parameters Use mTLS Trusted CA

A [`trusted_ca`](#trusted-ca-264d37) block (within [`https.tls_parameters.use_mtls`](#https-tls-parameters-use-mtls)) supports the following:

<a id="name-7721af"></a>&#x2022; [`name`](#name-7721af) - Optional String<br>Name. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name

<a id="namespace-70022b"></a>&#x2022; [`namespace`](#namespace-70022b) - Optional String<br>Namespace. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace

<a id="tenant-99da02"></a>&#x2022; [`tenant`](#tenant-99da02) - Optional String<br>Tenant. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant

#### HTTPS TLS Parameters Use mTLS Xfcc Options

A [`xfcc_options`](#options-4d1e53) block (within [`https.tls_parameters.use_mtls`](#https-tls-parameters-use-mtls)) supports the following:

<a id="elements-52ccb1"></a>&#x2022; [`xfcc_header_elements`](#elements-52ccb1) - Optional List  Defaults to `XFCC_NONE`<br>Possible values are `XFCC_NONE`, `XFCC_CERT`, `XFCC_CHAIN`, `XFCC_SUBJECT`, `XFCC_URI`, `XFCC_DNS`<br>[Enum: XFCC_NONE|XFCC_CERT|XFCC_CHAIN|XFCC_SUBJECT|XFCC_URI|XFCC_DNS] XFCC Header Elements. X-Forwarded-Client-Cert header elements to be added to requests

#### HTTPS Auto Cert

A [`https_auto_cert`](#https-auto-cert) block supports the following:

<a id="https-auto-cert-add-hsts"></a>&#x2022; [`add_hsts`](#https-auto-cert-add-hsts) - Optional Bool<br>Add HSTS Header. Add HTTP Strict-Transport-Security response header

<a id="https-auto-cert-append-server-name"></a>&#x2022; [`append_server_name`](#https-auto-cert-append-server-name) - Optional String<br>Append header value. Define the header value for the header name “server”. If header value is already present, it is not overwritten and passed as-is

<a id="https-auto-cert-coalescing-options"></a>&#x2022; [`coalescing_options`](#https-auto-cert-coalescing-options) - Optional Block<br>TLS Coalescing OPTIONS. TLS connection coalescing configuration (not compatible with mTLS)<br>See [Coalescing Options](#https-auto-cert-coalescing-options) below.

<a id="https-auto-cert-connection-idle-timeout"></a>&#x2022; [`connection_idle_timeout`](#https-auto-cert-connection-idle-timeout) - Optional Number  Defaults to `2`  Specified in milliseconds<br>Connection Idle Timeout. The idle timeout for downstream connections. The idle timeout is defined as the period in which there are no active requests. When the idle timeout is reached the connection will be closed. Note that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive.  The minutes

<a id="https-auto-cert-default-header"></a>&#x2022; [`default_header`](#https-auto-cert-default-header) - Optional Block<br>Enable this option

<a id="https-auto-cert-default-loadbalancer"></a>&#x2022; [`default_loadbalancer`](#https-auto-cert-default-loadbalancer) - Optional Block<br>Enable this option

<a id="https-auto-cert-disable-path-normalize"></a>&#x2022; [`disable_path_normalize`](#https-auto-cert-disable-path-normalize) - Optional Block<br>Enable this option

<a id="https-auto-cert-enable-path-normalize"></a>&#x2022; [`enable_path_normalize`](#https-auto-cert-enable-path-normalize) - Optional Block<br>Enable this option

<a id="https-auto-cert-http-protocol-options"></a>&#x2022; [`http_protocol_options`](#https-auto-cert-http-protocol-options) - Optional Block<br>HTTP Protocol Configuration OPTIONS. HTTP protocol configuration OPTIONS for downstream connections<br>See [HTTP Protocol Options](#https-auto-cert-http-protocol-options) below.

<a id="https-auto-cert-http-redirect"></a>&#x2022; [`http_redirect`](#https-auto-cert-http-redirect) - Optional Bool<br>HTTP Redirect to HTTPS. Redirect HTTP traffic to HTTPS

<a id="https-auto-cert-no-mtls"></a>&#x2022; [`no_mtls`](#https-auto-cert-no-mtls) - Optional Block<br>Enable this option

<a id="loadbalancer-eb605c"></a>&#x2022; [`non_default_loadbalancer`](#loadbalancer-eb605c) - Optional Block<br>Enable this option

<a id="https-auto-cert-pass-through"></a>&#x2022; [`pass_through`](#https-auto-cert-pass-through) - Optional Block<br>Enable this option

<a id="https-auto-cert-port"></a>&#x2022; [`port`](#https-auto-cert-port) - Optional Number<br>HTTPS Listen Port. HTTPS port to Listen

<a id="https-auto-cert-port-ranges"></a>&#x2022; [`port_ranges`](#https-auto-cert-port-ranges) - Optional String<br>Port Ranges. A string containing a comma separated list of port ranges. Each port range consists of a single port or two ports separated by '-'

<a id="https-auto-cert-server-name"></a>&#x2022; [`server_name`](#https-auto-cert-server-name) - Optional String<br>Modify header value. Define the header value for the header name “server”. This will overwrite existing values, if any, for the server header

<a id="https-auto-cert-tls-config"></a>&#x2022; [`tls_config`](#https-auto-cert-tls-config) - Optional Block<br>TLS Config. This defines various OPTIONS to configure TLS configuration parameters<br>See [TLS Config](#https-auto-cert-tls-config) below.

<a id="https-auto-cert-use-mtls"></a>&#x2022; [`use_mtls`](#https-auto-cert-use-mtls) - Optional Block<br>Clients TLS validation context. Validation context for downstream client TLS connections<br>See [Use mTLS](#https-auto-cert-use-mtls) below.

#### HTTPS Auto Cert Coalescing Options

A [`coalescing_options`](#https-auto-cert-coalescing-options) block (within [`https_auto_cert`](#https-auto-cert)) supports the following:

<a id="coalescing-3c2270"></a>&#x2022; [`default_coalescing`](#coalescing-3c2270) - Optional Block<br>Enable this option

<a id="coalescing-010f02"></a>&#x2022; [`strict_coalescing`](#coalescing-010f02) - Optional Block<br>Enable this option

#### HTTPS Auto Cert HTTP Protocol Options

A [`http_protocol_options`](#https-auto-cert-http-protocol-options) block (within [`https_auto_cert`](#https-auto-cert)) supports the following:

<a id="only-d515de"></a>&#x2022; [`http_protocol_enable_v1_only`](#only-d515de) - Optional Block<br>HTTP/1.1 Protocol OPTIONS. HTTP/1.1 Protocol OPTIONS for downstream connections<br>See [HTTP Protocol Enable V1 Only](#only-d515de) below.

<a id="v1-v2-9e0811"></a>&#x2022; [`http_protocol_enable_v1_v2`](#v1-v2-9e0811) - Optional Block<br>Enable this option

<a id="only-65e5e2"></a>&#x2022; [`http_protocol_enable_v2_only`](#only-65e5e2) - Optional Block<br>Enable this option

#### HTTPS Auto Cert HTTP Protocol Options HTTP Protocol Enable V1 Only

<a id="deep-b510f7"></a>Deeply nested **Only** block collapsed for readability.

#### HTTPS Auto Cert HTTP Protocol Options HTTP Protocol Enable V1 Only Header Transformation

<a id="deep-65ebe3"></a>Deeply nested **Transformation** block collapsed for readability.

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

<a id="https-auto-cert-use-mtls-crl-name"></a>&#x2022; [`name`](#https-auto-cert-use-mtls-crl-name) - Optional String<br>Name. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name

<a id="https-auto-cert-use-mtls-crl-namespace"></a>&#x2022; [`namespace`](#https-auto-cert-use-mtls-crl-namespace) - Optional String<br>Namespace. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace

<a id="https-auto-cert-use-mtls-crl-tenant"></a>&#x2022; [`tenant`](#https-auto-cert-use-mtls-crl-tenant) - Optional String<br>Tenant. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant

#### HTTPS Auto Cert Use mTLS Trusted CA

A [`trusted_ca`](#https-auto-cert-use-mtls-trusted-ca) block (within [`https_auto_cert.use_mtls`](#https-auto-cert-use-mtls)) supports the following:

<a id="name-36a848"></a>&#x2022; [`name`](#name-36a848) - Optional String<br>Name. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name

<a id="namespace-1c0f1b"></a>&#x2022; [`namespace`](#namespace-1c0f1b) - Optional String<br>Namespace. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace

<a id="tenant-08b11a"></a>&#x2022; [`tenant`](#tenant-08b11a) - Optional String<br>Tenant. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant

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

<a id="names-2ccfbe"></a>&#x2022; [`claim_names`](#names-2ccfbe) - Optional List<br>Claim Names. Human-readable name for the resource

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

<a id="name-21c5e6"></a>&#x2022; [`name`](#name-21c5e6) - Optional String<br>Name. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name

<a id="namespace-622fcf"></a>&#x2022; [`namespace`](#namespace-622fcf) - Optional String<br>Namespace. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace

<a id="tenant-e6ac4d"></a>&#x2022; [`tenant`](#tenant-e6ac4d) - Optional String<br>Tenant. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant

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

<a id="methods-02fcb9"></a>&#x2022; [`http_methods`](#methods-02fcb9) - Optional List  Defaults to `ANY`<br>See [HTTP Methods](#common-http-methods)<br> HTTP Methods. Methods to be matched

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

<a id="deep-a500e1"></a>Deeply nested **Domain** block collapsed for readability.

#### Malware Protection Settings Malware Protection Rules Metadata

A [`metadata`](#metadata-20b73a) block (within [`malware_protection_settings.malware_protection_rules`](#rules-b2bf3e)) supports the following:

<a id="spec-d9f05e"></a>&#x2022; [`description_spec`](#spec-d9f05e) - Optional String<br>Description. Human readable description

<a id="name-34fca6"></a>&#x2022; [`name`](#name-34fca6) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### Malware Protection Settings Malware Protection Rules Path

A [`path`](#path-ce187c) block (within [`malware_protection_settings.malware_protection_rules`](#rules-b2bf3e)) supports the following:

<a id="path-2aac29"></a>&#x2022; [`path`](#path-2aac29) - Optional String<br>Exact. Exact path value to match

<a id="prefix-4ddbaa"></a>&#x2022; [`prefix`](#prefix-4ddbaa) - Optional String<br>Prefix. Path prefix to match (e.g. The value / will match on all paths)

<a id="regex-56e93f"></a>&#x2022; [`regex`](#regex-56e93f) - Optional String<br>Regex. Regular expression of path match (e.g. The value .* will match on all paths)

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

<a id="header-45cce4"></a>&#x2022; [`remove_accept_encoding_header`](#header-45cce4) - Optional Bool<br>Remove Accept-Encoding Header. If true, removes accept-encoding from the request headers before dispatching it to the upstream so that responses do not GET compressed before reaching the filter

#### More Option Request Cookies To Add

A [`request_cookies_to_add`](#more-option-request-cookies-to-add) block (within [`more_option`](#more-option)) supports the following:

<a id="more-option-request-cookies-to-add-name"></a>&#x2022; [`name`](#more-option-request-cookies-to-add-name) - Optional String<br>Name. Name of the cookie in Cookie header

<a id="overwrite-6d9c60"></a>&#x2022; [`overwrite`](#overwrite-6d9c60) - Optional Bool  Defaults to `do`<br>Overwrite. Should the value be overwritten? If true, the value is overwritten to existing values. not overwrite

<a id="value-2d505e"></a>&#x2022; [`secret_value`](#value-2d505e) - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Secret Value](#value-2d505e) below.

<a id="value-9c117c"></a>&#x2022; [`value`](#value-9c117c) - Optional String<br>Value. Value of the Cookie header

#### More Option Request Cookies To Add Secret Value

<a id="deep-9b01f7"></a>Deeply nested **Value** block collapsed for readability.

#### More Option Request Cookies To Add Secret Value Blindfold Secret Info

<a id="deep-50ee6a"></a>Deeply nested **Info** block collapsed for readability.

#### More Option Request Cookies To Add Secret Value Clear Secret Info

<a id="deep-7082bd"></a>Deeply nested **Info** block collapsed for readability.

#### More Option Request Headers To Add

A [`request_headers_to_add`](#more-option-request-headers-to-add) block (within [`more_option`](#more-option)) supports the following:

<a id="append-2047af"></a>&#x2022; [`append`](#append-2047af) - Optional Bool  Defaults to `do`<br>Append. Should the value be appended? If true, the value is appended to existing values. not append

<a id="more-option-request-headers-to-add-name"></a>&#x2022; [`name`](#more-option-request-headers-to-add-name) - Optional String<br>Name. Name of the HTTP header

<a id="value-d68008"></a>&#x2022; [`secret_value`](#value-d68008) - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Secret Value](#value-d68008) below.

<a id="value-7dbd74"></a>&#x2022; [`value`](#value-7dbd74) - Optional String<br>Value. Value of the HTTP header

#### More Option Request Headers To Add Secret Value

<a id="deep-a39cd4"></a>Deeply nested **Value** block collapsed for readability.

#### More Option Request Headers To Add Secret Value Blindfold Secret Info

<a id="deep-e1bed7"></a>Deeply nested **Info** block collapsed for readability.

#### More Option Request Headers To Add Secret Value Clear Secret Info

<a id="deep-d2f0a4"></a>Deeply nested **Info** block collapsed for readability.

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

<a id="deep-0ef1c7"></a>Deeply nested **Value** block collapsed for readability.

#### More Option Response Cookies To Add Secret Value Blindfold Secret Info

<a id="deep-f1b5a7"></a>Deeply nested **Info** block collapsed for readability.

#### More Option Response Cookies To Add Secret Value Clear Secret Info

<a id="deep-8afac9"></a>Deeply nested **Info** block collapsed for readability.

#### More Option Response Headers To Add

A [`response_headers_to_add`](#more-option-response-headers-to-add) block (within [`more_option`](#more-option)) supports the following:

<a id="append-f099c0"></a>&#x2022; [`append`](#append-f099c0) - Optional Bool  Defaults to `do`<br>Append. Should the value be appended? If true, the value is appended to existing values. not append

<a id="name-9f1ebf"></a>&#x2022; [`name`](#name-9f1ebf) - Optional String<br>Name. Name of the HTTP header

<a id="value-76f49e"></a>&#x2022; [`secret_value`](#value-76f49e) - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Secret Value](#value-76f49e) below.

<a id="value-06dc79"></a>&#x2022; [`value`](#value-06dc79) - Optional String<br>Value. Value of the HTTP header

#### More Option Response Headers To Add Secret Value

<a id="deep-141ce7"></a>Deeply nested **Value** block collapsed for readability.

#### More Option Response Headers To Add Secret Value Blindfold Secret Info

<a id="deep-141889"></a>Deeply nested **Info** block collapsed for readability.

#### More Option Response Headers To Add Secret Value Clear Secret Info

<a id="deep-400eae"></a>Deeply nested **Info** block collapsed for readability.

#### Origin Server Subset Rule List

An [`origin_server_subset_rule_list`](#origin-server-subset-rule-list) block supports the following:

<a id="rules-7a881b"></a>&#x2022; [`origin_server_subset_rules`](#rules-7a881b) - Optional Block<br>Origin Server Subset Rules. Origin Server Subset Rules allow users to define match condition on Client (IP address, ASN, Country), IP Reputation, Regional Edge names, Request for subset selection of origin servers. Origin Server Subset is a sequential engine where rules are evaluated one after the other. It's important to define the correct order for Origin Server Subset to GET the intended result, rules are evaluated from top to bottom in the list. When an Origin server subset rule is matched, then this selection rule takes effect and no more rules are evaluated<br>See [Origin Server Subset Rules](#rules-7a881b) below.

#### Origin Server Subset Rule List Origin Server Subset Rules

<a id="deep-e82219"></a>Deeply nested **Rules** block collapsed for readability.

#### Origin Server Subset Rule List Origin Server Subset Rules Asn List

<a id="deep-fd4e45"></a>Deeply nested **List** block collapsed for readability.

#### Origin Server Subset Rule List Origin Server Subset Rules Asn Matcher

<a id="deep-5db5db"></a>Deeply nested **Matcher** block collapsed for readability.

#### Origin Server Subset Rule List Origin Server Subset Rules Asn Matcher Asn Sets

<a id="deep-5a7316"></a>Deeply nested **Sets** block collapsed for readability.

#### Origin Server Subset Rule List Origin Server Subset Rules Client Selector

<a id="deep-92e33a"></a>Deeply nested **Selector** block collapsed for readability.

#### Origin Server Subset Rule List Origin Server Subset Rules IP Matcher

<a id="deep-ae40ab"></a>Deeply nested **Matcher** block collapsed for readability.

#### Origin Server Subset Rule List Origin Server Subset Rules IP Matcher Prefix Sets

<a id="deep-158295"></a>Deeply nested **Sets** block collapsed for readability.

#### Origin Server Subset Rule List Origin Server Subset Rules IP Prefix List

<a id="deep-b80c7e"></a>Deeply nested **List** block collapsed for readability.

#### Origin Server Subset Rule List Origin Server Subset Rules Metadata

<a id="deep-4083d2"></a>Deeply nested **Metadata** block collapsed for readability.

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

<a id="name-0fb02d"></a>&#x2022; [`name`](#name-0fb02d) - Optional String<br>Name. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name

<a id="namespace-490d76"></a>&#x2022; [`namespace`](#namespace-490d76) - Optional String<br>Namespace. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace

<a id="tenant-cf334a"></a>&#x2022; [`tenant`](#tenant-cf334a) - Optional String<br>Tenant. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant

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

<a id="selector-ca44f5"></a>&#x2022; [`client_selector`](#selector-ca44f5) - Optional Block<br>Label Selector. This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expresssions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects. Label selector is immutable. Expressions is a list of strings of label selection expression. Each string has ',' separated values which are 'AND' and all strings are logically 'OR'. BNF for expression string `<selector-syntax>` ::= `<requirement>` | `<requirement>` ',' `<selector-syntax>` `<requirement>` ::= [!] KEY [ `<set-based-restriction>` | `<exact-match-restriction>` ] `<set-based-restriction>` ::= '' | `<inclusion-exclusion>` `<value-set>` `<inclusion-exclusion>` ::= `<inclusion>` | `<exclusion>` `<exclusion>` ::= 'notin' `<inclusion>` ::= 'in' `<value-set>` ::= '(' `<values>` ')' `<values>` ::= VALUE | VALUE ',' `<values>` `<exact-match-restriction>` ::= ['='|'=='|'!='] VALUE<br>See [Client Selector](#selector-ca44f5) below.

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

<a id="deep-eb8ffe"></a>Deeply nested **Matchers** block collapsed for readability.

#### Policy Based Challenge Rule List Rules Spec Arg Matchers Item

<a id="deep-e03a12"></a>Deeply nested **Item** block collapsed for readability.

#### Policy Based Challenge Rule List Rules Spec Asn List

<a id="deep-8e9207"></a>Deeply nested **List** block collapsed for readability.

#### Policy Based Challenge Rule List Rules Spec Asn Matcher

<a id="deep-4e0cfb"></a>Deeply nested **Matcher** block collapsed for readability.

#### Policy Based Challenge Rule List Rules Spec Asn Matcher Asn Sets

<a id="deep-2a5120"></a>Deeply nested **Sets** block collapsed for readability.

#### Policy Based Challenge Rule List Rules Spec Body Matcher

<a id="deep-43ae26"></a>Deeply nested **Matcher** block collapsed for readability.

#### Policy Based Challenge Rule List Rules Spec Client Selector

<a id="deep-73214b"></a>Deeply nested **Selector** block collapsed for readability.

#### Policy Based Challenge Rule List Rules Spec Cookie Matchers

<a id="deep-946f25"></a>Deeply nested **Matchers** block collapsed for readability.

#### Policy Based Challenge Rule List Rules Spec Cookie Matchers Item

<a id="deep-939b70"></a>Deeply nested **Item** block collapsed for readability.

#### Policy Based Challenge Rule List Rules Spec Domain Matcher

<a id="deep-c2c201"></a>Deeply nested **Matcher** block collapsed for readability.

#### Policy Based Challenge Rule List Rules Spec Headers

<a id="deep-c62b4f"></a>Deeply nested **Headers** block collapsed for readability.

#### Policy Based Challenge Rule List Rules Spec Headers Item

<a id="deep-b6374c"></a>Deeply nested **Item** block collapsed for readability.

#### Policy Based Challenge Rule List Rules Spec HTTP Method

<a id="deep-dba95f"></a>Deeply nested **Method** block collapsed for readability.

#### Policy Based Challenge Rule List Rules Spec IP Matcher

<a id="deep-b37ae0"></a>Deeply nested **Matcher** block collapsed for readability.

#### Policy Based Challenge Rule List Rules Spec IP Matcher Prefix Sets

<a id="deep-73bd32"></a>Deeply nested **Sets** block collapsed for readability.

#### Policy Based Challenge Rule List Rules Spec IP Prefix List

<a id="deep-fb570d"></a>Deeply nested **List** block collapsed for readability.

#### Policy Based Challenge Rule List Rules Spec Path

<a id="deep-2ed6cf"></a>Deeply nested **Path** block collapsed for readability.

#### Policy Based Challenge Rule List Rules Spec Query Params

<a id="deep-39ee89"></a>Deeply nested **Params** block collapsed for readability.

#### Policy Based Challenge Rule List Rules Spec Query Params Item

<a id="deep-f76120"></a>Deeply nested **Item** block collapsed for readability.

#### Policy Based Challenge Rule List Rules Spec TLS Fingerprint Matcher

<a id="deep-2cb28a"></a>Deeply nested **Matcher** block collapsed for readability.

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

<a id="deep-e048af"></a>Deeply nested **Prefixes** block collapsed for readability.

#### Rate Limit IP Allowed List

An [`ip_allowed_list`](#rate-limit-ip-allowed-list) block (within [`rate_limit`](#rate-limit)) supports the following:

<a id="rate-limit-ip-allowed-list-prefixes"></a>&#x2022; [`prefixes`](#rate-limit-ip-allowed-list-prefixes) - Optional List<br>IPv4 Prefix List. List of IPv4 prefixes that represent an endpoint

#### Rate Limit Policies

A [`policies`](#rate-limit-policies) block (within [`rate_limit`](#rate-limit)) supports the following:

<a id="rate-limit-policies-policies"></a>&#x2022; [`policies`](#rate-limit-policies-policies) - Optional Block<br>Rate Limiter Policies. Ordered list of rate limiter policies<br>See [Policies](#rate-limit-policies-policies) below.

#### Rate Limit Policies Policies

A [`policies`](#rate-limit-policies-policies) block (within [`rate_limit.policies`](#rate-limit-policies)) supports the following:

<a id="rate-limit-policies-policies-name"></a>&#x2022; [`name`](#rate-limit-policies-policies-name) - Optional String<br>Name. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name

<a id="rate-limit-policies-policies-namespace"></a>&#x2022; [`namespace`](#rate-limit-policies-policies-namespace) - Optional String<br>Namespace. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace

<a id="rate-limit-policies-policies-tenant"></a>&#x2022; [`tenant`](#rate-limit-policies-policies-tenant) - Optional String<br>Tenant. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant

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

<a id="duration-617314"></a>&#x2022; [`duration`](#duration-617314) - Optional Number<br>Duration. Configuration parameter for duration

#### Rate Limit Rate Limiter Action Block Minutes

A [`minutes`](#minutes-c83f64) block (within [`rate_limit.rate_limiter.action_block`](#rate-limit-rate-limiter-action-block)) supports the following:

<a id="duration-534bd9"></a>&#x2022; [`duration`](#duration-534bd9) - Optional Number<br>Duration. Configuration parameter for duration

#### Rate Limit Rate Limiter Action Block Seconds

A [`seconds`](#seconds-8810ec) block (within [`rate_limit.rate_limiter.action_block`](#rate-limit-rate-limiter-action-block)) supports the following:

<a id="duration-dfe2a4"></a>&#x2022; [`duration`](#duration-dfe2a4) - Optional Number<br>Duration. Configuration parameter for duration

#### Ring Hash

A [`ring_hash`](#ring-hash) block supports the following:

<a id="ring-hash-hash-policy"></a>&#x2022; [`hash_policy`](#ring-hash-hash-policy) - Optional Block<br>Hash Policy. Specifies a list of hash policies to use for ring hash load balancing. Each hash policy is evaluated individually and the combined result is used to route the request<br>See [Hash Policy](#ring-hash-hash-policy) below.

#### Ring Hash Hash Policy

A [`hash_policy`](#ring-hash-hash-policy) block (within [`ring_hash`](#ring-hash)) supports the following:

<a id="ring-hash-hash-policy-cookie"></a>&#x2022; [`cookie`](#ring-hash-hash-policy-cookie) - Optional Block<br>Hashing using Cookie. Two types of cookie affinity: 1. Passive. Takes a cookie that's present in the cookies header and hashes on its value. 2. Generated. Generates and sets a cookie with an expiration (TTL) on the first request from the client in its response to the client, based on the endpoint the request gets sent to. The client then presents this on the next and all subsequent requests. The hash of this is sufficient to ensure these requests GET sent to the same endpoint. The cookie is generated by hashing the source and destination ports and addresses so that multiple independent HTTP2 streams on the same connection will independently receive the same cookie, even if they arrive simultaneously<br>See [Cookie](#ring-hash-hash-policy-cookie) below.

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

<a id="name-82f318"></a>&#x2022; [`name`](#name-82f318) - Optional String<br>Name. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name

<a id="namespace-7ec6c4"></a>&#x2022; [`namespace`](#namespace-7ec6c4) - Optional String<br>Namespace. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace

<a id="tenant-a123ba"></a>&#x2022; [`tenant`](#tenant-a123ba) - Optional String<br>Tenant. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant

#### Routes Direct Response Route

A [`direct_response_route`](#routes-direct-response-route) block (within [`routes`](#routes)) supports the following:

<a id="routes-direct-response-route-headers"></a>&#x2022; [`headers`](#routes-direct-response-route-headers) - Optional Block<br>Headers. List of (key, value) headers<br>See [Headers](#routes-direct-response-route-headers) below.

<a id="method-aec314"></a>&#x2022; [`http_method`](#method-aec314) - Optional String  Defaults to `ANY`<br>See [HTTP Methods](#common-http-methods)<br> HTTP Method. Specifies the HTTP method used to access a resource. Any HTTP Method

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

<a id="prefix-ff3976"></a>&#x2022; [`prefix`](#prefix-ff3976) - Optional String<br>Prefix. Path prefix to match (e.g. The value / will match on all paths)

<a id="routes-direct-response-route-path-regex"></a>&#x2022; [`regex`](#routes-direct-response-route-path-regex) - Optional String<br>Regex. Regular expression of path match (e.g. The value .* will match on all paths)

#### Routes Direct Response Route Route Direct Response

A [`route_direct_response`](#response-d0dcbd) block (within [`routes.direct_response_route`](#routes-direct-response-route)) supports the following:

<a id="encoded-a56f81"></a>&#x2022; [`response_body_encoded`](#encoded-a56f81) - Optional String<br>Response Body. Response body to send. Currently supported URL schemes is string:/// for which message should be encoded in Base64 format. The message can be either plain text or HTML. E.g. '`<p>` Access Denied `</p>`'. Base64 encoded string URL for this is string:///PHA+IEFjY2VzcyBEZW5pZWQgPC9wPg==

<a id="code-1bc88c"></a>&#x2022; [`response_code`](#code-1bc88c) - Optional Number<br>Response Code. Response code to send

#### Routes Redirect Route

A [`redirect_route`](#routes-redirect-route) block (within [`routes`](#routes)) supports the following:

<a id="routes-redirect-route-headers"></a>&#x2022; [`headers`](#routes-redirect-route-headers) - Optional Block<br>Headers. List of (key, value) headers<br>See [Headers](#routes-redirect-route-headers) below.

<a id="routes-redirect-route-http-method"></a>&#x2022; [`http_method`](#routes-redirect-route-http-method) - Optional String  Defaults to `ANY`<br>See [HTTP Methods](#common-http-methods)<br> HTTP Method. Specifies the HTTP method used to access a resource. Any HTTP Method

<a id="routes-redirect-route-incoming-port"></a>&#x2022; [`incoming_port`](#routes-redirect-route-incoming-port) - Optional Block<br>Port to Match. Port match of the request can be a range or a specific port<br>See [Incoming Port](#routes-redirect-route-incoming-port) below.

<a id="routes-redirect-route-path"></a>&#x2022; [`path`](#routes-redirect-route-path) - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#routes-redirect-route-path) below.

<a id="routes-redirect-route-route-redirect"></a>&#x2022; [`route_redirect`](#routes-redirect-route-route-redirect) - Optional Block<br>Redirect. Route redirect parameters when match action is redirect<br>See [Route Redirect](#routes-redirect-route-route-redirect) below.

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

<a id="routes-redirect-route-path-prefix"></a>&#x2022; [`prefix`](#routes-redirect-route-path-prefix) - Optional String<br>Prefix. Path prefix to match (e.g. The value / will match on all paths)

<a id="routes-redirect-route-path-regex"></a>&#x2022; [`regex`](#routes-redirect-route-path-regex) - Optional String<br>Regex. Regular expression of path match (e.g. The value .* will match on all paths)

#### Routes Redirect Route Route Redirect

A [`route_redirect`](#routes-redirect-route-route-redirect) block (within [`routes.redirect_route`](#routes-redirect-route)) supports the following:

<a id="redirect-bd044d"></a>&#x2022; [`host_redirect`](#redirect-bd044d) - Optional String<br>Host. Swap host part of incoming URL in redirect URL

<a id="redirect-2ae47a"></a>&#x2022; [`path_redirect`](#redirect-2ae47a) - Optional String<br>Path. swap path part of incoming URL in redirect URL

<a id="rewrite-a81c41"></a>&#x2022; [`prefix_rewrite`](#rewrite-a81c41) - Optional String<br>Prefix Rewrite. In Redirect response, the matched prefix (or path) should be swapped with this value. This option allows redirect URLs be dynamically created based on the request

<a id="redirect-f23979"></a>&#x2022; [`proto_redirect`](#redirect-f23979) - Optional String<br>Protocol. Swap protocol part of incoming URL in redirect URL The protocol can be swapped with either HTTP or HTTPS When incoming-proto option is specified, swapping of protocol is not done

<a id="params-0941dc"></a>&#x2022; [`remove_all_params`](#params-0941dc) - Optional Block<br>Enable this option

<a id="params-94a828"></a>&#x2022; [`replace_params`](#params-94a828) - Optional String<br>Replace All Parameters

<a id="code-d55c43"></a>&#x2022; [`response_code`](#code-d55c43) - Optional Number<br>Response Code. The HTTP status code to use in the redirect response

<a id="params-f96588"></a>&#x2022; [`retain_all_params`](#params-f96588) - Optional Block<br>Enable this option

#### Routes Simple Route

A [`simple_route`](#routes-simple-route) block (within [`routes`](#routes)) supports the following:

<a id="routes-simple-route-advanced-options"></a>&#x2022; [`advanced_options`](#routes-simple-route-advanced-options) - Optional Block<br>Advanced Route OPTIONS. Configure advanced OPTIONS for route like path rewrite, hash policy, etc<br>See [Advanced Options](#routes-simple-route-advanced-options) below.

<a id="routes-simple-route-auto-host-rewrite"></a>&#x2022; [`auto_host_rewrite`](#routes-simple-route-auto-host-rewrite) - Optional Block<br>Enable this option

<a id="rewrite-706535"></a>&#x2022; [`disable_host_rewrite`](#rewrite-706535) - Optional Block<br>Enable this option

<a id="routes-simple-route-headers"></a>&#x2022; [`headers`](#routes-simple-route-headers) - Optional Block<br>Headers. List of (key, value) headers<br>See [Headers](#routes-simple-route-headers) below.

<a id="routes-simple-route-host-rewrite"></a>&#x2022; [`host_rewrite`](#routes-simple-route-host-rewrite) - Optional String<br>Host Rewrite Value. Host header will be swapped with this value

<a id="routes-simple-route-http-method"></a>&#x2022; [`http_method`](#routes-simple-route-http-method) - Optional String  Defaults to `ANY`<br>See [HTTP Methods](#common-http-methods)<br> HTTP Method. Specifies the HTTP method used to access a resource. Any HTTP Method

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

<a id="policy-ba853e"></a>&#x2022; [`cors_policy`](#policy-ba853e) - Optional Block<br>CORS Policy. Cross-Origin Resource Sharing requests configuration specified at Virtual-host or Route level. Route level configuration takes precedence. An example of an Cross origin HTTP request GET /resources/public-data/ HTTP/1.1 Host: bar.other User-Agent: Mozilla/5.0 (Macintosh; U; Intel MAC OS X 10.5; en-US; rv:1.9.1b3pre) Gecko/20081130 Minefield/3.1b3pre Accept: text/HTML,application/xhtml+XML,application/XML;q=0.9,*/*;q=0.8 Accept-Language: en-us,en;q=0.5 Accept-Encoding: gzip,deflate Accept-Charset: ISO-8859-1,utf-8;q=0.7,*;q=0.7 Connection: keep-alive Referrer: `HTTP://foo.example/examples/access-control/simplexsinvocation.HTML` Origin: `HTTP://foo.example` HTTP/1.1 200 OK Date: Mon, 01 Dec 2008 00:23:53 GMT Server: Apache/2.0.61 Access-Control-Allow-Origin: \* Keep-Alive: timeout=2, max=100 Connection: Keep-Alive Transfer-Encoding: chunked Content-Type: application/XML An example for cross origin HTTP OPTIONS request with Access-Control-Request-* header OPTIONS /resources/POST-here/ HTTP/1.1 Host: bar.other User-Agent: Mozilla/5.0 (Macintosh; U; Intel MAC OS X 10.5; en-US; rv:1.9.1b3pre) Gecko/20081130 Minefield/3.1b3pre Accept: text/HTML,application/xhtml+XML,application/XML;q=0.9,*/*;q=0.8 Accept-Language: en-us,en;q=0.5 Accept-Encoding: gzip,deflate Accept-Charset: ISO-8859-1,utf-8;q=0.7,*;q=0.7 Connection: keep-alive Origin: `HTTP://foo.example` Access-Control-Request-Method: POST Access-Control-Request-Headers: X-PINGOTHER, Content-Type HTTP/1.1 204 No Content Date: Mon, 01 Dec 2008 01:15:39 GMT Server: Apache/2.0.61 (Unix) Access-Control-Allow-Origin: `HTTP://foo.example` Access-Control-Allow-Methods: POST, GET, OPTIONS Access-Control-Allow-Headers: X-PINGOTHER, Content-Type Access-Control-Max-Age: 86400 Vary: Accept-Encoding, Origin Keep-Alive: timeout=2, max=100 Connection: Keep-Alive<br>See [CORS Policy](#policy-ba853e) below.

<a id="policy-7816d7"></a>&#x2022; [`csrf_policy`](#policy-7816d7) - Optional Block<br>CSRF Policy. To mitigate CSRF attack , the policy checks where a request is coming from to determine if the request's origin is the same as its detination.the policy relies on two pieces of information used in determining if a request originated from the same host. 1. The origin that caused the user agent to issue the request (source origin). 2. The origin that the request is going to (target origin). When the policy evaluating a request, it ensures both pieces of information are present and compare their values. If the source origin is missing or origins do not match the request is rejected. The exception to this being if the source-origin has been added to they policy as valid. Because CSRF attacks specifically target state-changing requests, the policy only acts on the HTTP requests that have state-changing method (PUT,POST, etc.)<br>See [CSRF Policy](#policy-7816d7) below.

<a id="policy-70b68a"></a>&#x2022; [`default_retry_policy`](#policy-70b68a) - Optional Block<br>Enable this option

<a id="add-11129b"></a>&#x2022; [`disable_location_add`](#add-11129b) - Optional Bool<br>Disable Location Addition. Disables append of x-F5 Distributed Cloud-location = `<RE-site-name>` at route level, if it is configured at virtual-host level. This configuration is ignored on CE sites

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

<a id="name-dd1197"></a>&#x2022; [`name`](#name-dd1197) - Optional String<br>Name. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name

<a id="namespace-118663"></a>&#x2022; [`namespace`](#namespace-118663) - Optional String<br>Namespace. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace

<a id="tenant-6ccea5"></a>&#x2022; [`tenant`](#tenant-6ccea5) - Optional String<br>Tenant. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant

#### Routes Simple Route Advanced Options Bot Defense Javascript Injection

<a id="deep-2061b8"></a>Deeply nested **Injection** block collapsed for readability.

#### Routes Simple Route Advanced Options Bot Defense Javascript Injection Javascript Tags

<a id="deep-5972b3"></a>Deeply nested **Tags** block collapsed for readability.

#### Routes Simple Route Advanced Options Bot Defense Javascript Injection Javascript Tags Tag Attributes

<a id="deep-67ccd3"></a>Deeply nested **Attributes** block collapsed for readability.

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

<a id="deep-e656f6"></a>Deeply nested **List** block collapsed for readability.

#### Routes Simple Route Advanced Options Mirror Policy

A [`mirror_policy`](#policy-f5e84d) block (within [`routes.simple_route.advanced_options`](#routes-simple-route-advanced-options)) supports the following:

<a id="pool-8c75a0"></a>&#x2022; [`origin_pool`](#pool-8c75a0) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Origin Pool](#pool-8c75a0) below.

<a id="percent-99590a"></a>&#x2022; [`percent`](#percent-99590a) - Optional Block<br>Fractional Percent. Fraction used where sampling percentages are needed. Example sampled requests<br>See [Percent](#percent-99590a) below.

#### Routes Simple Route Advanced Options Mirror Policy Origin Pool

<a id="deep-5e0e39"></a>Deeply nested **Pool** block collapsed for readability.

#### Routes Simple Route Advanced Options Mirror Policy Percent

<a id="deep-8287f0"></a>Deeply nested **Percent** block collapsed for readability.

#### Routes Simple Route Advanced Options Regex Rewrite

A [`regex_rewrite`](#rewrite-c628a7) block (within [`routes.simple_route.advanced_options`](#routes-simple-route-advanced-options)) supports the following:

<a id="pattern-191576"></a>&#x2022; [`pattern`](#pattern-191576) - Optional String<br>Pattern. The regular expression used to find portions of a string that should be replaced

<a id="substitution-55c137"></a>&#x2022; [`substitution`](#substitution-55c137) - Optional String<br>Substitution. The string that should be substituted into matching portions of the subject string during a substitution operation to produce a new string

#### Routes Simple Route Advanced Options Request Cookies To Add

<a id="deep-ecb2d3"></a>Deeply nested **Add** block collapsed for readability.

#### Routes Simple Route Advanced Options Request Cookies To Add Secret Value

<a id="deep-ed6f4f"></a>Deeply nested **Value** block collapsed for readability.

#### Routes Simple Route Advanced Options Request Cookies To Add Secret Value Blindfold Secret Info

<a id="deep-f40bbd"></a>Deeply nested **Info** block collapsed for readability.

#### Routes Simple Route Advanced Options Request Cookies To Add Secret Value Clear Secret Info

<a id="deep-80c44d"></a>Deeply nested **Info** block collapsed for readability.

#### Routes Simple Route Advanced Options Request Headers To Add

<a id="deep-15e359"></a>Deeply nested **Add** block collapsed for readability.

#### Routes Simple Route Advanced Options Request Headers To Add Secret Value

<a id="deep-dc94f9"></a>Deeply nested **Value** block collapsed for readability.

#### Routes Simple Route Advanced Options Request Headers To Add Secret Value Blindfold Secret Info

<a id="deep-f976d3"></a>Deeply nested **Info** block collapsed for readability.

#### Routes Simple Route Advanced Options Request Headers To Add Secret Value Clear Secret Info

<a id="deep-e1345a"></a>Deeply nested **Info** block collapsed for readability.

#### Routes Simple Route Advanced Options Response Cookies To Add

<a id="deep-e44886"></a>Deeply nested **Add** block collapsed for readability.

#### Routes Simple Route Advanced Options Response Cookies To Add Secret Value

<a id="deep-a0b78b"></a>Deeply nested **Value** block collapsed for readability.

#### Routes Simple Route Advanced Options Response Cookies To Add Secret Value Blindfold Secret Info

<a id="deep-8eb600"></a>Deeply nested **Info** block collapsed for readability.

#### Routes Simple Route Advanced Options Response Cookies To Add Secret Value Clear Secret Info

<a id="deep-0efa32"></a>Deeply nested **Info** block collapsed for readability.

#### Routes Simple Route Advanced Options Response Headers To Add

<a id="deep-b11d53"></a>Deeply nested **Add** block collapsed for readability.

#### Routes Simple Route Advanced Options Response Headers To Add Secret Value

<a id="deep-c69bf4"></a>Deeply nested **Value** block collapsed for readability.

#### Routes Simple Route Advanced Options Response Headers To Add Secret Value Blindfold Secret Info

<a id="deep-a91b72"></a>Deeply nested **Info** block collapsed for readability.

#### Routes Simple Route Advanced Options Response Headers To Add Secret Value Clear Secret Info

<a id="deep-be0841"></a>Deeply nested **Info** block collapsed for readability.

#### Routes Simple Route Advanced Options Retry Policy

A [`retry_policy`](#policy-e40fa6) block (within [`routes.simple_route.advanced_options`](#routes-simple-route-advanced-options)) supports the following:

<a id="off-e4369f"></a>&#x2022; [`back_off`](#off-e4369f) - Optional Block<br>Retry BackOff Interval. Specifies parameters that control retry back off<br>See [Back Off](#off-e4369f) below.

<a id="retries-ee7703"></a>&#x2022; [`num_retries`](#retries-ee7703) - Optional Number  Defaults to `1`<br>Number of Retries. Specifies the allowed number of retries. Retries can be done any number of times. An exponential back-off algorithm is used between each retry

<a id="timeout-2485fd"></a>&#x2022; [`per_try_timeout`](#timeout-2485fd) - Optional Number<br>Per Try Timeout. Specifies a non-zero timeout per retry attempt. In milliseconds

<a id="codes-110133"></a>&#x2022; [`retriable_status_codes`](#codes-110133) - Optional List<br>Status Code to Retry. HTTP status codes that should trigger a retry in addition to those specified by retry_on

<a id="condition-28432a"></a>&#x2022; [`retry_condition`](#condition-28432a) - Optional List<br>Retry Condition. Specifies the conditions under which retry takes place. Retries can be on different types of condition depending on application requirements. For example, network failure, all 5xx response codes, idempotent 4xx response codes, etc The possible values are '5xx' : Retry will be done if the upstream server responds with any 5xx response code, or does not respond at all (disconnect/reset/read timeout). 'gateway-error' : Retry will be done only if the upstream server responds with 502, 503 or 504 responses (Included in 5xx) 'connect-failure' : Retry will be done if the request fails because of a connection failure to the upstream server (connect timeout, etc.). (Included in 5xx) 'refused-stream' : Retry is done if the upstream server resets the stream with a REFUSED_STREAM error code (Included in 5xx) 'retriable-4xx' : Retry is done if the upstream server responds with a retriable 4xx response code. The only response code in this category is HTTP CONFLICT (409) 'retriable-status-codes' : Retry is done if the upstream server responds with any response code matching one defined in retriable_status_codes field 'reset' : Retry is done if the upstream server does not respond at all (disconnect/reset/read timeout.)

#### Routes Simple Route Advanced Options Retry Policy Back Off

<a id="deep-8cef77"></a>Deeply nested **Off** block collapsed for readability.

#### Routes Simple Route Advanced Options Specific Hash Policy

<a id="deep-4a119f"></a>Deeply nested **Policy** block collapsed for readability.

#### Routes Simple Route Advanced Options Specific Hash Policy Hash Policy

<a id="deep-4b04ff"></a>Deeply nested **Policy** block collapsed for readability.

#### Routes Simple Route Advanced Options Specific Hash Policy Hash Policy Cookie

<a id="deep-44b5c0"></a>Deeply nested **Cookie** block collapsed for readability.

#### Routes Simple Route Advanced Options WAF Exclusion Policy

<a id="deep-40a4da"></a>Deeply nested **Policy** block collapsed for readability.

#### Routes Simple Route Advanced Options Web Socket Config

<a id="deep-d144b5"></a>Deeply nested **Config** block collapsed for readability.

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

<a id="name-82cffb"></a>&#x2022; [`name`](#name-82cffb) - Optional String<br>Name. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name

<a id="namespace-b73e1d"></a>&#x2022; [`namespace`](#namespace-b73e1d) - Optional String<br>Namespace. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace

<a id="tenant-511970"></a>&#x2022; [`tenant`](#tenant-511970) - Optional String<br>Tenant. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant

#### Routes Simple Route Origin Pools Pool

A [`pool`](#routes-simple-route-origin-pools-pool) block (within [`routes.simple_route.origin_pools`](#routes-simple-route-origin-pools)) supports the following:

<a id="name-7e0a9d"></a>&#x2022; [`name`](#name-7e0a9d) - Optional String<br>Name. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name

<a id="namespace-f2152d"></a>&#x2022; [`namespace`](#namespace-f2152d) - Optional String<br>Namespace. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace

<a id="tenant-06f782"></a>&#x2022; [`tenant`](#tenant-06f782) - Optional String<br>Tenant. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant

#### Routes Simple Route Path

A [`path`](#routes-simple-route-path) block (within [`routes.simple_route`](#routes-simple-route)) supports the following:

<a id="routes-simple-route-path-path"></a>&#x2022; [`path`](#routes-simple-route-path-path) - Optional String<br>Exact. Exact path value to match

<a id="routes-simple-route-path-prefix"></a>&#x2022; [`prefix`](#routes-simple-route-path-prefix) - Optional String<br>Prefix. Path prefix to match (e.g. The value / will match on all paths)

<a id="routes-simple-route-path-regex"></a>&#x2022; [`regex`](#routes-simple-route-path-regex) - Optional String<br>Regex. Regular expression of path match (e.g. The value .* will match on all paths)

#### Routes Simple Route Query Params

A [`query_params`](#routes-simple-route-query-params) block (within [`routes.simple_route`](#routes-simple-route)) supports the following:

<a id="params-aa1f17"></a>&#x2022; [`remove_all_params`](#params-aa1f17) - Optional Block<br>Enable this option

<a id="params-c3e5f1"></a>&#x2022; [`replace_params`](#params-c3e5f1) - Optional String<br>Replace All Parameters

<a id="params-bd2237"></a>&#x2022; [`retain_all_params`](#params-bd2237) - Optional Block<br>Enable this option

#### Sensitive Data Disclosure Rules

A [`sensitive_data_disclosure_rules`](#sensitive-data-disclosure-rules) block supports the following:

<a id="response-2680e4"></a>&#x2022; [`sensitive_data_types_in_response`](#response-2680e4) - Optional Block<br>Sensitive Data Exposure Rules. Sensitive Data Exposure Rules allows specifying rules to mask sensitive data fields in API responses<br>See [Sensitive Data Types In Response](#response-2680e4) below.

#### Sensitive Data Disclosure Rules Sensitive Data Types In Response

<a id="deep-44339a"></a>Deeply nested **Response** block collapsed for readability.

#### Sensitive Data Disclosure Rules Sensitive Data Types In Response API Endpoint

<a id="deep-cc989e"></a>Deeply nested **Endpoint** block collapsed for readability.

#### Sensitive Data Disclosure Rules Sensitive Data Types In Response Body

<a id="deep-cb4010"></a>Deeply nested **Body** block collapsed for readability.

#### Sensitive Data Policy

A [`sensitive_data_policy`](#sensitive-data-policy) block supports the following:

<a id="ref-55b260"></a>&#x2022; [`sensitive_data_policy_ref`](#ref-55b260) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Sensitive Data Policy Ref](#ref-55b260) below.

#### Sensitive Data Policy Sensitive Data Policy Ref

A [`sensitive_data_policy_ref`](#ref-55b260) block (within [`sensitive_data_policy`](#sensitive-data-policy)) supports the following:

<a id="name-d254a7"></a>&#x2022; [`name`](#name-d254a7) - Optional String<br>Name. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name

<a id="namespace-401387"></a>&#x2022; [`namespace`](#namespace-401387) - Optional String<br>Namespace. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace

<a id="tenant-d10cc7"></a>&#x2022; [`tenant`](#tenant-d10cc7) - Optional String<br>Tenant. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant

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

<a id="deep-989da5"></a>Deeply nested **Config** block collapsed for readability.

#### Single LB App Enable Discovery API Crawler API Crawler Config Domains

<a id="deep-fdc95f"></a>Deeply nested **Domains** block collapsed for readability.

#### Single LB App Enable Discovery API Crawler API Crawler Config Domains Simple Login

<a id="deep-cdb4fc"></a>Deeply nested **Login** block collapsed for readability.

#### Single LB App Enable Discovery API Crawler API Crawler Config Domains Simple Login Password

<a id="deep-6dac8b"></a>Deeply nested **Password** block collapsed for readability.

#### Single LB App Enable Discovery API Crawler API Crawler Config Domains Simple Login Password Blindfold Secret Info

<a id="deep-50ceec"></a>Deeply nested **Info** block collapsed for readability.

#### Single LB App Enable Discovery API Crawler API Crawler Config Domains Simple Login Password Clear Secret Info

<a id="deep-c26a90"></a>Deeply nested **Info** block collapsed for readability.

#### Single LB App Enable Discovery API Discovery From Code Scan

<a id="deep-7dbe06"></a>Deeply nested **Scan** block collapsed for readability.

#### Single LB App Enable Discovery API Discovery From Code Scan Code Base Integrations

<a id="deep-301c15"></a>Deeply nested **Integrations** block collapsed for readability.

#### Single LB App Enable Discovery API Discovery From Code Scan Code Base Integrations Code Base Integration

<a id="deep-a9865c"></a>Deeply nested **Integration** block collapsed for readability.

#### Single LB App Enable Discovery API Discovery From Code Scan Code Base Integrations Selected Repos

<a id="deep-19872d"></a>Deeply nested **Repos** block collapsed for readability.

#### Single LB App Enable Discovery Custom API Auth Discovery

<a id="deep-c1e6ff"></a>Deeply nested **Discovery** block collapsed for readability.

#### Single LB App Enable Discovery Custom API Auth Discovery API Discovery Ref

<a id="deep-7417e7"></a>Deeply nested **Ref** block collapsed for readability.

#### Single LB App Enable Discovery Discovered API Settings

<a id="deep-437161"></a>Deeply nested **Settings** block collapsed for readability.

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

<a id="user-identification-name"></a>&#x2022; [`name`](#user-identification-name) - Optional String<br>Name. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name

<a id="user-identification-namespace"></a>&#x2022; [`namespace`](#user-identification-namespace) - Optional String<br>Namespace. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace

<a id="user-identification-tenant"></a>&#x2022; [`tenant`](#user-identification-tenant) - Optional String<br>Tenant. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant

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

<a id="methods-19f73d"></a>&#x2022; [`methods`](#methods-19f73d) - Optional List  Defaults to `ANY`<br>See [HTTP Methods](#common-http-methods)<br> Methods. Methods to be matched

<a id="prefix-a857dd"></a>&#x2022; [`path_prefix`](#prefix-a857dd) - Optional String<br>Prefix. Path prefix to match (e.g. The value / will match on all paths)

<a id="regex-fdbacd"></a>&#x2022; [`path_regex`](#regex-fdbacd) - Optional String<br>Path Regex. Define the regex for the path. For example, the regex ^/.*$ will match on all paths

<a id="value-6f2f58"></a>&#x2022; [`suffix_value`](#value-6f2f58) - Optional String<br>Suffix Value. Suffix of domain name e.g 'xyz.com' will match '*.xyz.com' and 'xyz.com'

<a id="processing-8c8391"></a>&#x2022; [`waf_skip_processing`](#processing-8c8391) - Optional Block<br>Enable this option

#### WAF Exclusion WAF Exclusion Inline Rules Rules App Firewall Detection Control

<a id="deep-832ffb"></a>Deeply nested **Control** block collapsed for readability.

#### WAF Exclusion WAF Exclusion Inline Rules Rules App Firewall Detection Control Exclude Attack Type Contexts

<a id="deep-0e5af0"></a>Deeply nested **Contexts** block collapsed for readability.

#### WAF Exclusion WAF Exclusion Inline Rules Rules App Firewall Detection Control Exclude Bot Name Contexts

<a id="deep-22f5f8"></a>Deeply nested **Contexts** block collapsed for readability.

#### WAF Exclusion WAF Exclusion Inline Rules Rules App Firewall Detection Control Exclude Signature Contexts

<a id="deep-30cc06"></a>Deeply nested **Contexts** block collapsed for readability.

#### WAF Exclusion WAF Exclusion Inline Rules Rules App Firewall Detection Control Exclude Violation Contexts

<a id="deep-d6ba53"></a>Deeply nested **Contexts** block collapsed for readability.

#### WAF Exclusion WAF Exclusion Inline Rules Rules Metadata

<a id="deep-68f59c"></a>Deeply nested **Metadata** block collapsed for readability.

#### WAF Exclusion WAF Exclusion Policy

A [`waf_exclusion_policy`](#waf-exclusion-waf-exclusion-policy) block (within [`waf_exclusion`](#waf-exclusion)) supports the following:

<a id="waf-exclusion-waf-exclusion-policy-name"></a>&#x2022; [`name`](#waf-exclusion-waf-exclusion-policy-name) - Optional String<br>Name. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name

<a id="namespace-d8f030"></a>&#x2022; [`namespace`](#namespace-d8f030) - Optional String<br>Namespace. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace

<a id="tenant-d841f0"></a>&#x2022; [`tenant`](#tenant-d841f0) - Optional String<br>Tenant. When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant

---

## Common Types

The following type definitions are used throughout this resource. See the full definition here rather than repeated inline.

### Object Reference {#common-object-reference}

Object references establish a direct reference from one configuration object to another in F5 Distributed Cloud. References use the format `tenant/namespace/name`.

| Field | Type | Description |
| ----- | ---- | ----------- |
| `name` | String | Name of the referenced object |
| `namespace` | String | Namespace containing the referenced object |
| `tenant` | String | Tenant of the referenced object (system-managed) |

### Transformers {#common-transformers}

Transformers apply transformations to input values before matching. Multiple transformers can be applied in order.

| Value | Description |
| ----- | ----------- |
| `LOWER_CASE` | Convert to lowercase |
| `UPPER_CASE` | Convert to uppercase |
| `BASE64_DECODE` | Decode base64 content |
| `NORMALIZE_PATH` | Normalize URL path |
| `REMOVE_WHITESPACE` | Remove whitespace characters |
| `URL_DECODE` | Decode URL-encoded characters |
| `TRIM_LEFT` | Trim leading whitespace |
| `TRIM_RIGHT` | Trim trailing whitespace |
| `TRIM` | Trim both leading and trailing whitespace |

### HTTP Methods {#common-http-methods}

HTTP methods used for request matching.

| Value | Description |
| ----- | ----------- |
| `ANY` | Match any HTTP method |
| `GET` | HTTP GET request |
| `HEAD` | HTTP HEAD request |
| `POST` | HTTP POST request |
| `PUT` | HTTP PUT request |
| `DELETE` | HTTP DELETE request |
| `CONNECT` | HTTP CONNECT request |
| `OPTIONS` | HTTP OPTIONS request |
| `TRACE` | HTTP TRACE request |
| `PATCH` | HTTP PATCH request |
| `COPY` | HTTP COPY request (WebDAV) |

### TLS Fingerprints {#common-tls-fingerprints}

TLS fingerprint categories for malicious client detection.

| Value | Description |
| ----- | ----------- |
| `TLS_FINGERPRINT_NONE` | No fingerprint matching |
| `ANY_MALICIOUS_FINGERPRINT` | Match any known malicious fingerprint |
| `ADWARE` | Adware-associated fingerprints |
| `DRIDEX` | Dridex malware fingerprints |
| `GOOTKIT` | Gootkit malware fingerprints |
| `RANSOMWARE` | Ransomware-associated fingerprints |
| `TRICKBOT` | Trickbot malware fingerprints |

### IP Threat Categories {#common-ip-threat-categories}

IP address threat categories for security filtering.

| Value | Description |
| ----- | ----------- |
| `SPAM_SOURCES` | Known spam sources |
| `WINDOWS_EXPLOITS` | Windows exploit sources |
| `WEB_ATTACKS` | Web attack sources |
| `BOTNETS` | Known botnet IPs |
| `SCANNERS` | Network scanner IPs |
| `REPUTATION` | Poor reputation IPs |
| `PHISHING` | Phishing-related IPs |
| `PROXY` | Anonymous proxy IPs |
| `MOBILE_THREATS` | Mobile threat sources |
| `TOR_PROXY` | Tor exit nodes |
| `DENIAL_OF_SERVICE` | DoS attack sources |
| `NETWORK` | Known bad network ranges |

## Import

Import is supported using the following syntax:

```shell
# Import using namespace/name format
terraform import f5xc_http_loadbalancer.example system/example
```
