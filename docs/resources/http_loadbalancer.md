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

<a id="name"></a>&#x2022; [`name`](#name) - Required String<br>Name of the HTTPLoadBalancer. Must be unique within the namespace

<a id="namespace"></a>&#x2022; [`namespace`](#namespace) - Required String<br>Namespace where the HTTPLoadBalancer will be created

<a id="annotations"></a>&#x2022; [`annotations`](#annotations) - Optional Map<br>Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata

<a id="description"></a>&#x2022; [`description`](#description) - Optional String<br>Human readable description for the object

<a id="disable"></a>&#x2022; [`disable`](#disable) - Optional Bool<br>A value of true will administratively disable the object

<a id="labels"></a>&#x2022; [`labels`](#labels) - Optional Map<br>Labels is a user defined key value map that can be attached to resources for organization and filtering

### Spec Argument Reference

-> **One of the following:**
&#x2022; <a id="active-service-policies"></a>[`active_service_policies`](#active-service-policies) - Optional Block<br>Service Policy List. List of service policies<br>See [Active Service Policies](#active-service-policies) below for details.
<br><br>&#x2022; <a id="no-service-policies"></a>[`no_service_policies`](#no-service-policies) - Optional Block<br>Empty. This can be used for messages where no values are needed
<br><br>&#x2022; <a id="service-policies-from-namespace"></a>[`service_policies_from_namespace`](#service-policies-from-namespace) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="add-location"></a>&#x2022; [`add_location`](#add-location) - Optional Bool<br>Add Location. x-example: true Appends header x-volterra-location = `<RE-site-name>` in responses. This configuration is ignored on CE sites

-> **One of the following:**
&#x2022; <a id="advertise-custom"></a>[`advertise_custom`](#advertise-custom) - Optional Block<br>Advertise Custom. This defines a way to advertise a VIP on specific sites<br>See [Advertise Custom](#advertise-custom) below for details.
<br><br>&#x2022; <a id="advertise-on-public"></a>[`advertise_on_public`](#advertise-on-public) - Optional Block<br>Advertise Public. This defines a way to advertise a load balancer on public. If optional public_ip is provided, it will only be advertised on RE sites where that public_ip is available<br>See [Advertise On Public](#advertise-on-public) below for details.
<br><br>&#x2022; <a id="advertise-on-public-default-vip"></a>[`advertise_on_public_default_vip`](#advertise-on-public-default-vip) - Optional Block<br>Empty. This can be used for messages where no values are needed
<br><br>&#x2022; <a id="do-not-advertise"></a>[`do_not_advertise`](#do-not-advertise) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-protection-rules"></a>&#x2022; [`api_protection_rules`](#api-protection-rules) - Optional Block<br>API Protection Rules. API Protection Rules<br>See [API Protection Rules](#api-protection-rules) below for details.

-> **One of the following:**
&#x2022; <a id="api-rate-limit"></a>[`api_rate_limit`](#api-rate-limit) - Optional Block<br>APIRateLimit
<br><br>&#x2022; <a id="disable-rate-limit"></a>[`disable_rate_limit`](#disable-rate-limit) - Optional Block<br>Empty. This can be used for messages where no values are needed
<br><br>&#x2022; <a id="rate-limit"></a>[`rate_limit`](#rate-limit) - Optional Block<br>RateLimitConfigType

-> **One of the following:**
&#x2022; <a id="api-specification"></a>[`api_specification`](#api-specification) - Optional Block<br>API Specification and Validation. Settings for API specification (API definition, OpenAPI validation, etc.)
<br><br>&#x2022; <a id="disable-api-definition"></a>[`disable_api_definition`](#disable-api-definition) - Optional Block<br>Empty. This can be used for messages where no values are needed

-> **One of the following:**
&#x2022; <a id="api-testing"></a>[`api_testing`](#api-testing) - Optional Block<br>API Testing
<br><br>&#x2022; <a id="disable-api-testing"></a>[`disable_api_testing`](#disable-api-testing) - Optional Block<br>Empty. This can be used for messages where no values are needed

-> **One of the following:**
&#x2022; <a id="app-firewall"></a>[`app_firewall`](#app-firewall) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name
<br><br>&#x2022; <a id="disable-waf"></a>[`disable_waf`](#disable-waf) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="blocked-clients"></a>&#x2022; [`blocked_clients`](#blocked-clients) - Optional Block<br>Client Blocking Rules. Define rules to block IP Prefixes or AS numbers

-> **One of the following:**
&#x2022; <a id="bot-defense"></a>[`bot_defense`](#bot-defense) - Optional Block<br>Bot Defense. This defines various configuration options for Bot Defense Policy
<br><br>&#x2022; <a id="bot-defense-advanced"></a>[`bot_defense_advanced`](#bot-defense-advanced) - Optional Block<br>Bot Defense Advanced. Bot Defense Advanced
<br><br>&#x2022; <a id="disable-bot-defense"></a>[`disable_bot_defense`](#disable-bot-defense) - Optional Block<br>Empty. This can be used for messages where no values are needed

-> **One of the following:**
&#x2022; <a id="caching-policy"></a>[`caching_policy`](#caching-policy) - Optional Block<br>Caching Policies. x-required Caching Policies for the CDN
<br><br>&#x2022; <a id="disable-caching"></a>[`disable_caching`](#disable-caching) - Optional Block<br>Empty. This can be used for messages where no values are needed

-> **One of the following:**
&#x2022; <a id="captcha-challenge"></a>[`captcha_challenge`](#captcha-challenge) - Optional Block<br>Captcha Challenge Parameters. Enables loadbalancer to perform captcha challenge Captcha challenge will be based on Google Recaptcha. With this feature enabled, only clients that pass the captcha challenge will be allowed to complete the HTTP request. When loadbalancer is configured to do Captcha Challenge, it will redirect the browser to an HTML page on every new HTTP request. This HTML page will have captcha challenge embedded in it. Client will be allowed to make the request only if the captcha challenge is successful. Loadbalancer will tag response header with a cookie to avoid Captcha challenge for subsequent requests. CAPTCHA is mainly used as a security check to ensure only human users can pass through. Generally, computers or bots are not capable of solving a captcha. You can enable either Javascript challenge or Captcha challenge on a virtual host
<br><br>&#x2022; <a id="enable-challenge"></a>[`enable_challenge`](#enable-challenge) - Optional Block<br>Enable Malicious User Challenge. Configure auto mitigation i.e risk based challenges for malicious users
<br><br>&#x2022; <a id="js-challenge"></a>[`js_challenge`](#js-challenge) - Optional Block<br>Javascript Challenge Parameters. Enables loadbalancer to perform client browser compatibility test by redirecting to a page with Javascript. With this feature enabled, only clients that are capable of executing Javascript(mostly browsers) will be allowed to complete the HTTP request. When loadbalancer is configured to do Javascript Challenge, it will redirect the browser to an HTML page on every new HTTP request. This HTML page will have Javascript embedded in it. Loadbalancer chooses a set of random numbers for every new client and sends these numbers along with an encrypted answer with the request such that it embed these numbers as input in the Javascript. Javascript will run on the requestor browser and perform a complex Math operation. Script will submit the answer to loadbalancer. Loadbalancer will validate the answer by comparing the calculated answer with the decrypted answer (which was encrypted when it was sent back as reply) and allow the request to the upstream server only if the answer is correct. Loadbalancer will tag response header with a cookie to avoid Javascript challenge for subsequent requests. Javascript challenge serves following purposes * Validate that the request is coming via a browser that is capable for running Javascript * Force the browser to run a complex operation, f(X), that requires it to spend a large number of CPU cycles. This is to slow down a potential DOS attacker by making it difficult to launch a large request flood without having to spend even larger CPU cost at their end. You can enable either Javascript challenge or Captcha challenge on a virtual host
<br><br>&#x2022; <a id="no-challenge"></a>[`no_challenge`](#no-challenge) - Optional Block<br>Empty. This can be used for messages where no values are needed
<br><br>&#x2022; <a id="policy-based-challenge"></a>[`policy_based_challenge`](#policy-based-challenge) - Optional Block<br>Policy Based Challenge. Specifies the settings for policy rule based challenge

-> **One of the following:**
&#x2022; <a id="client-side-defense"></a>[`client_side_defense`](#client-side-defense) - Optional Block<br>Client-Side Defense. This defines various configuration options for Client-Side Defense Policy
<br><br>&#x2022; <a id="disable-client-side-defense"></a>[`disable_client_side_defense`](#disable-client-side-defense) - Optional Block<br>Empty. This can be used for messages where no values are needed

-> **One of the following:**
&#x2022; <a id="cookie-stickiness"></a>[`cookie_stickiness`](#cookie-stickiness) - Optional Block<br>Hashing using Cookie. Two types of cookie affinity: 1. Passive. Takes a cookie that's present in the cookies header and hashes on its value. 2. Generated. Generates and sets a cookie with an expiration (TTL) on the first request from the client in its response to the client, based on the endpoint the request gets sent to. The client then presents this on the next and all subsequent requests. The hash of this is sufficient to ensure these requests get sent to the same endpoint. The cookie is generated by hashing the source and destination ports and addresses so that multiple independent HTTP2 streams on the same connection will independently receive the same cookie, even if they arrive simultaneously
<br><br>&#x2022; <a id="least-active"></a>[`least_active`](#least-active) - Optional Block<br>Empty. This can be used for messages where no values are needed
<br><br>&#x2022; <a id="random"></a>[`random`](#random) - Optional Block<br>Empty. This can be used for messages where no values are needed
<br><br>&#x2022; <a id="ring-hash"></a>[`ring_hash`](#ring-hash) - Optional Block<br>Hash Policy List. List of hash policy rules
<br><br>&#x2022; <a id="round-robin"></a>[`round_robin`](#round-robin) - Optional Block<br>Empty. This can be used for messages where no values are needed
<br><br>&#x2022; <a id="source-ip-stickiness"></a>[`source_ip_stickiness`](#source-ip-stickiness) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="cors-policy"></a>&#x2022; [`cors_policy`](#cors-policy) - Optional Block<br>CORS Policy. Cross-Origin Resource Sharing requests configuration specified at Virtual-host or Route level. Route level configuration takes precedence. An example of an Cross origin HTTP request GET /resources/public-data/ HTTP/1.1 Host: bar.other User-Agent: Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.5; en-US; rv:1.9.1b3pre) Gecko/20081130 Minefield/3.1b3pre Accept: text/HTML,application/xhtml+XML,application/XML;q=0.9,*/*;q=0.8 Accept-Language: en-us,en;q=0.5 Accept-Encoding: gzip,deflate Accept-Charset: ISO-8859-1,utf-8;q=0.7,*;q=0.7 Connection: keep-alive Referrer: `HTTP://foo.example/examples/access-control/simpleXSInvocation.HTML` Origin: `HTTP://foo.example` HTTP/1.1 200 OK Date: Mon, 01 Dec 2008 00:23:53 GMT Server: Apache/2.0.61 Access-Control-Allow-Origin: * Keep-Alive: timeout=2, max=100 Connection: Keep-Alive Transfer-Encoding: chunked Content-Type: application/XML An example for cross origin HTTP OPTIONS request with Access-Control-Request-* header OPTIONS /resources/post-here/ HTTP/1.1 Host: bar.other User-Agent: Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.5; en-US; rv:1.9.1b3pre) Gecko/20081130 Minefield/3.1b3pre Accept: text/HTML,application/xhtml+XML,application/XML;q=0.9,*/*;q=0.8 Accept-Language: en-us,en;q=0.5 Accept-Encoding: gzip,deflate Accept-Charset: ISO-8859-1,utf-8;q=0.7,*;q=0.7 Connection: keep-alive Origin: `HTTP://foo.example` Access-Control-Request-Method: POST Access-Control-Request-Headers: X-PINGOTHER, Content-Type HTTP/1.1 204 No Content Date: Mon, 01 Dec 2008 01:15:39 GMT Server: Apache/2.0.61 (Unix) Access-Control-Allow-Origin: `HTTP://foo.example` Access-Control-Allow-Methods: POST, GET, OPTIONS Access-Control-Allow-Headers: X-PINGOTHER, Content-Type Access-Control-Max-Age: 86400 Vary: Accept-Encoding, Origin Keep-Alive: timeout=2, max=100 Connection: Keep-Alive

<a id="csrf-policy"></a>&#x2022; [`csrf_policy`](#csrf-policy) - Optional Block<br>CSRF Policy. To mitigate CSRF attack , the policy checks where a request is coming from to determine if the request's origin is the same as its detination.The policy relies on two pieces of information used in determining if a request originated from the same host. 1. The origin that caused the user agent to issue the request (source origin). 2. The origin that the request is going to (target origin). When the policy evaluating a request, it ensures both pieces of information are present and compare their values. If the source origin is missing or origins do not match the request is rejected. The exception to this being if the source-origin has been added to they policy as valid. Because CSRF attacks specifically target state-changing requests, the policy only acts on the HTTP requests that have state-changing method (PUT,POST, etc.)

<a id="data-guard-rules"></a>&#x2022; [`data_guard_rules`](#data-guard-rules) - Optional Block<br>Data Guard Rules. Data Guard prevents responses from exposing sensitive information by masking the data. The system masks credit card numbers and social security numbers leaked from the application from within the HTTP response with a string of asterisks (*). Note: App Firewall should be enabled, to use Data Guard feature

<a id="ddos-mitigation-rules"></a>&#x2022; [`ddos_mitigation_rules`](#ddos-mitigation-rules) - Optional Block<br>DDOS Mitigation Rules. Define manual mitigation rules to block L7 DDOS attacks

-> **One of the following:**
&#x2022; <a id="default-pool"></a>[`default_pool`](#default-pool) - Optional Block<br>Global Specification. Shape of the origin pool specification
<br><br>&#x2022; <a id="default-pool-list"></a>[`default_pool_list`](#default-pool-list) - Optional Block<br>Origin Pool List Type. List of Origin Pools

<a id="default-route-pools"></a>&#x2022; [`default_route_pools`](#default-route-pools) - Optional Block<br>Origin Pools. Origin Pools used when no route is specified (default route)

-> **One of the following:**
&#x2022; <a id="default-sensitive-data-policy"></a>[`default_sensitive_data_policy`](#default-sensitive-data-policy) - Optional Block<br>Empty. This can be used for messages where no values are needed
<br><br>&#x2022; <a id="sensitive-data-policy"></a>[`sensitive_data_policy`](#sensitive-data-policy) - Optional Block<br>Sensitive Data Discovery. Settings for data type policy

-> **One of the following:**
&#x2022; <a id="disable-api-discovery"></a>[`disable_api_discovery`](#disable-api-discovery) - Optional Block<br>Empty. This can be used for messages where no values are needed
<br><br>&#x2022; <a id="enable-api-discovery"></a>[`enable_api_discovery`](#enable-api-discovery) - Optional Block<br>API Discovery Setting. Specifies the settings used for API discovery

-> **One of the following:**
&#x2022; <a id="disable-ip-reputation"></a>[`disable_ip_reputation`](#disable-ip-reputation) - Optional Block<br>Empty. This can be used for messages where no values are needed
<br><br>&#x2022; <a id="enable-ip-reputation"></a>[`enable_ip_reputation`](#enable-ip-reputation) - Optional Block<br>IP Threat Category List. List of IP threat categories

-> **One of the following:**
&#x2022; <a id="disable-malicious-user-detection"></a>[`disable_malicious_user_detection`](#disable-malicious-user-detection) - Optional Block<br>Empty. This can be used for messages where no values are needed
<br><br>&#x2022; <a id="enable-malicious-user-detection"></a>[`enable_malicious_user_detection`](#enable-malicious-user-detection) - Optional Block<br>Empty. This can be used for messages where no values are needed

-> **One of the following:**
&#x2022; <a id="disable-malware-protection"></a>[`disable_malware_protection`](#disable-malware-protection) - Optional Block<br>Empty. This can be used for messages where no values are needed
<br><br>&#x2022; <a id="malware-protection-settings"></a>[`malware_protection_settings`](#malware-protection-settings) - Optional Block<br>Malware Protection Policy. Malware Protection protects Web Apps and APIs, from malicious file uploads by scanning files in real-time

-> **One of the following:**
&#x2022; <a id="disable-threat-mesh"></a>[`disable_threat_mesh`](#disable-threat-mesh) - Optional Block<br>Empty. This can be used for messages where no values are needed
<br><br>&#x2022; <a id="enable-threat-mesh"></a>[`enable_threat_mesh`](#enable-threat-mesh) - Optional Block<br>Empty. This can be used for messages where no values are needed

-> **One of the following:**
&#x2022; <a id="disable-trust-client-ip-headers"></a>[`disable_trust_client_ip_headers`](#disable-trust-client-ip-headers) - Optional Block<br>Empty. This can be used for messages where no values are needed
<br><br>&#x2022; <a id="enable-trust-client-ip-headers"></a>[`enable_trust_client_ip_headers`](#enable-trust-client-ip-headers) - Optional Block<br>Trust Client IP Headers List. List of Client IP Headers

<a id="domains"></a>&#x2022; [`domains`](#domains) - Optional List<br>Domains. A list of Domains (host/authority header) that will be matched to load balancer. Supported Domains and search order: 1. Exact Domain names: `www.foo.com.` 2. Domains starting with a Wildcard: *.foo.com. Not supported Domains: - Just a Wildcard: * - A Wildcard and TLD with no root Domain: *.com. - A Wildcard not matching a whole DNS label. e.g. *.foo.com and *.bar.foo.com are valid Wildcards however *bar.foo.com, *-bar.foo.com, and bar*.foo.com are all invalid. Additional notes: A Wildcard will not match empty string. e.g. *.foo.com will match bar.foo.com and baz-bar.foo.com but not .foo.com. The longest Wildcards match first. Only a single virtual host in the entire route configuration can match on *. Also a Domain must be unique across all virtual hosts within an advertise policy. Domains are also used for SNI matching if the Loadbalancer type is HTTPS. Domains also indicate the list of names for which DNS resolution will be automatically resolved to IP addresses by the system

<a id="graphql-rules"></a>&#x2022; [`graphql_rules`](#graphql-rules) - Optional Block<br>GraphQL Inspection. GraphQL is a query language and server-side runtime for APIs which provides a complete and understandable description of the data in API. GraphQL gives clients the power to ask for exactly what they need, makes it easier to evolve APIs over time, and enables powerful developer tools. Policy configuration to analyze GraphQL queries and prevent GraphQL tailored attacks

-> **One of the following:**
&#x2022; <a id="http"></a>[`http`](#http) - Optional Block<br>HTTP Choice. Choice for selecting HTTP proxy
<br><br>&#x2022; <a id="https"></a>[`https`](#https) - Optional Block<br>BYOC HTTPS Choice. Choice for selecting HTTP proxy with bring your own certificates
<br><br>&#x2022; <a id="https-auto-cert"></a>[`https_auto_cert`](#https-auto-cert) - Optional Block<br>HTTPS with Auto Certs Choice. Choice for selecting HTTP proxy with bring your own certificates

<a id="jwt-validation"></a>&#x2022; [`jwt_validation`](#jwt-validation) - Optional Block<br>JWT Validation. JWT Validation stops JWT replay attacks and JWT tampering by cryptographically verifying incoming JWTs before they are passed to your API origin. JWT Validation will also stop requests with expired tokens or tokens that are not yet valid

-> **One of the following:**
&#x2022; <a id="l7-ddos-action-block"></a>[`l7_ddos_action_block`](#l7-ddos-action-block) - Optional Block<br>Empty. This can be used for messages where no values are needed
<br><br>&#x2022; <a id="l7-ddos-action-default"></a>[`l7_ddos_action_default`](#l7-ddos-action-default) - Optional Block<br>Empty. This can be used for messages where no values are needed
<br><br>&#x2022; <a id="l7-ddos-action-js-challenge"></a>[`l7_ddos_action_js_challenge`](#l7-ddos-action-js-challenge) - Optional Block<br>Javascript Challenge Parameters. Enables loadbalancer to perform client browser compatibility test by redirecting to a page with Javascript. With this feature enabled, only clients that are capable of executing Javascript(mostly browsers) will be allowed to complete the HTTP request. When loadbalancer is configured to do Javascript Challenge, it will redirect the browser to an HTML page on every new HTTP request. This HTML page will have Javascript embedded in it. Loadbalancer chooses a set of random numbers for every new client and sends these numbers along with an encrypted answer with the request such that it embed these numbers as input in the Javascript. Javascript will run on the requestor browser and perform a complex Math operation. Script will submit the answer to loadbalancer. Loadbalancer will validate the answer by comparing the calculated answer with the decrypted answer (which was encrypted when it was sent back as reply) and allow the request to the upstream server only if the answer is correct. Loadbalancer will tag response header with a cookie to avoid Javascript challenge for subsequent requests. Javascript challenge serves following purposes * Validate that the request is coming via a browser that is capable for running Javascript * Force the browser to run a complex operation, f(X), that requires it to spend a large number of CPU cycles. This is to slow down a potential DOS attacker by making it difficult to launch a large request flood without having to spend even larger CPU cost at their end. You can enable either Javascript challenge or Captcha challenge on a virtual host

<a id="l7-ddos-protection"></a>&#x2022; [`l7_ddos_protection`](#l7-ddos-protection) - Optional Block<br>L7 DDOS Protection Settings. L7 DDOS protection is critical for safeguarding web applications, APIs, and services that are exposed to the internet from sophisticated, volumetric, application-level threats. Configure actions, thresholds and policies to apply during L7 DDOS attack

<a id="more-option"></a>&#x2022; [`more_option`](#more-option) - Optional Block<br>Advanced Options. This defines various options to define a route

-> **One of the following:**
&#x2022; <a id="multi-lb-app"></a>[`multi_lb_app`](#multi-lb-app) - Optional Block<br>Empty. This can be used for messages where no values are needed
<br><br>&#x2022; <a id="single-lb-app"></a>[`single_lb_app`](#single-lb-app) - Optional Block<br>Single Load Balancer App Setting. Specific settings for Machine learning analysis on this HTTP LB, independently from other LBs

<a id="origin-server-subset-rule-list"></a>&#x2022; [`origin_server_subset_rule_list`](#origin-server-subset-rule-list) - Optional Block<br>Origin Server Subset Rule List Type. List of Origin Pools

<a id="protected-cookies"></a>&#x2022; [`protected_cookies`](#protected-cookies) - Optional Block<br>Cookie Protection. Allows setting attributes (SameSite, Secure, and HttpOnly) on cookies in responses. Cookie Tampering Protection prevents attackers from modifying the value of session cookies. For Cookie Tampering Protection, enabling a web app firewall (WAF) is a prerequisite. The configured mode of WAF (monitoring or blocking) will be enforced on the request when cookie tampering is identified. Note: We recommend enabling Secure and HttpOnly attributes along with cookie tampering protection

<a id="routes"></a>&#x2022; [`routes`](#routes) - Optional Block<br>Routes. Routes allow users to define match condition on a path and/or HTTP method to either forward matching traffic to origin pool or redirect matching traffic to a different URL or respond directly to matching traffic

<a id="sensitive-data-disclosure-rules"></a>&#x2022; [`sensitive_data_disclosure_rules`](#sensitive-data-disclosure-rules) - Optional Block<br>Sensitive Data Exposure Rules. Sensitive Data Exposure Rules allows specifying rules to mask sensitive data fields in API responses

-> **One of the following:**
&#x2022; <a id="slow-ddos-mitigation"></a>[`slow_ddos_mitigation`](#slow-ddos-mitigation) - Optional Block<br>Slow DDOS Mitigation. 'Slow and low' attacks tie up server resources, leaving none available for servicing requests from actual users
<br><br>&#x2022; <a id="system-default-timeouts"></a>[`system_default_timeouts`](#system-default-timeouts) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="timeouts"></a>&#x2022; [`timeouts`](#timeouts) - Optional Block

<a id="trusted-clients"></a>&#x2022; [`trusted_clients`](#trusted-clients) - Optional Block<br>Trusted Client Rules. Define rules to skip processing of one or more features such as WAF, Bot Defense etc. for clients

-> **One of the following:**
&#x2022; <a id="user-id-client-ip"></a>[`user_id_client_ip`](#user-id-client-ip) - Optional Block<br>Empty. This can be used for messages where no values are needed
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

<a id="active-service-policies-policies-namespace"></a>&#x2022; [`namespace`](#active-service-policies-policies-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="active-service-policies-policies-tenant"></a>&#x2022; [`tenant`](#active-service-policies-policies-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Advertise Custom

An [`advertise_custom`](#advertise-custom) block supports the following:

<a id="advertise-custom-advertise-where"></a>&#x2022; [`advertise_where`](#advertise-custom-advertise-where) - Optional Block<br>List of Sites to Advertise. Where should this load balancer be available<br>See [Advertise Where](#advertise-custom-advertise-where) below.

#### Advertise Custom Advertise Where

An [`advertise_where`](#advertise-custom-advertise-where) block (within [`advertise_custom`](#advertise-custom)) supports the following:

<a id="advertise-custom-advertise-where-advertise-on-public"></a>&#x2022; [`advertise_on_public`](#advertise-custom-advertise-where-advertise-on-public) - Optional Block<br>Advertise Public. This defines a way to advertise a load balancer on public. If optional public_ip is provided, it will only be advertised on RE sites where that public_ip is available<br>See [Advertise On Public](#advertise-custom-advertise-where-advertise-on-public) below.

<a id="advertise-custom-advertise-where-port"></a>&#x2022; [`port`](#advertise-custom-advertise-where-port) - Optional Number<br>Listen Port. Port to Listen

<a id="advertise-custom-advertise-where-port-ranges"></a>&#x2022; [`port_ranges`](#advertise-custom-advertise-where-port-ranges) - Optional String<br>Listen Port Ranges. A string containing a comma separated list of port ranges. Each port range consists of a single port or two ports separated by '-'

<a id="advertise-custom-advertise-where-site"></a>&#x2022; [`site`](#advertise-custom-advertise-where-site) - Optional Block<br>Site. This defines a reference to a CE site along with network type and an optional IP address where a load balancer could be advertised<br>See [Site](#advertise-custom-advertise-where-site) below.

<a id="advertise-custom-advertise-where-use-default-port"></a>&#x2022; [`use_default_port`](#advertise-custom-advertise-where-use-default-port) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="advertise-custom-advertise-where-virtual-network"></a>&#x2022; [`virtual_network`](#advertise-custom-advertise-where-virtual-network) - Optional Block<br>Virtual Network. Parameters to advertise on a given virtual network<br>See [Virtual Network](#advertise-custom-advertise-where-virtual-network) below.

<a id="advertise-custom-advertise-where-virtual-site"></a>&#x2022; [`virtual_site`](#advertise-custom-advertise-where-virtual-site) - Optional Block<br>Virtual Site. This defines a reference to a customer site virtual site along with network type where a load balancer could be advertised<br>See [Virtual Site](#advertise-custom-advertise-where-virtual-site) below.

<a id="advertise-custom-advertise-where-virtual-site-with-vip"></a>&#x2022; [`virtual_site_with_vip`](#advertise-custom-advertise-where-virtual-site-with-vip) - Optional Block<br>Virtual Site with Specified VIP. This defines a reference to a customer site virtual site along with network type and IP where a load balancer could be advertised<br>See [Virtual Site With VIP](#advertise-custom-advertise-where-virtual-site-with-vip) below.

<a id="advertise-custom-advertise-where-vk8s-service"></a>&#x2022; [`vk8s_service`](#advertise-custom-advertise-where-vk8s-service) - Optional Block<br>vK8s Services on RE. This defines a reference to a RE site or virtual site where a load balancer could be advertised in the vK8s service network<br>See [Vk8s Service](#advertise-custom-advertise-where-vk8s-service) below.

#### Advertise Custom Advertise Where Advertise On Public

An [`advertise_on_public`](#advertise-custom-advertise-where-advertise-on-public) block (within [`advertise_custom.advertise_where`](#advertise-custom-advertise-where)) supports the following:

<a id="advertise-custom-advertise-where-advertise-on-public-public-ip"></a>&#x2022; [`public_ip`](#advertise-custom-advertise-where-advertise-on-public-public-ip) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Public IP](#advertise-custom-advertise-where-advertise-on-public-public-ip) below.

#### Advertise Custom Advertise Where Advertise On Public Public IP

A [`public_ip`](#advertise-custom-advertise-where-advertise-on-public-public-ip) block (within [`advertise_custom.advertise_where.advertise_on_public`](#advertise-custom-advertise-where-advertise-on-public)) supports the following:

<a id="advertise-custom-advertise-where-advertise-on-public-public-ip-name"></a>&#x2022; [`name`](#advertise-custom-advertise-where-advertise-on-public-public-ip-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="advertise-custom-advertise-where-advertise-on-public-public-ip-namespace"></a>&#x2022; [`namespace`](#advertise-custom-advertise-where-advertise-on-public-public-ip-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="advertise-custom-advertise-where-advertise-on-public-public-ip-tenant"></a>&#x2022; [`tenant`](#advertise-custom-advertise-where-advertise-on-public-public-ip-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Advertise Custom Advertise Where Site

A [`site`](#advertise-custom-advertise-where-site) block (within [`advertise_custom.advertise_where`](#advertise-custom-advertise-where)) supports the following:

<a id="advertise-custom-advertise-where-site-ip"></a>&#x2022; [`ip`](#advertise-custom-advertise-where-site-ip) - Optional String<br>IP Address. Use given IP address as VIP on the site

<a id="advertise-custom-advertise-where-site-network"></a>&#x2022; [`network`](#advertise-custom-advertise-where-site-network) - Optional String  Defaults to `SITE_NETWORK_INSIDE_AND_OUTSIDE`<br>Possible values are `SITE_NETWORK_INSIDE_AND_OUTSIDE`, `SITE_NETWORK_INSIDE`, `SITE_NETWORK_OUTSIDE`, `SITE_NETWORK_SERVICE`, `SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP`, `SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP`, `SITE_NETWORK_IP_FABRIC`<br>Site Network. This defines network types to be used on site All inside and outside networks. All inside and outside networks with internet VIP support. All inside networks. All outside networks. All outside networks with internet VIP support. vK8s service network. - SITE_NETWORK_IP_FABRIC: VER IP Fabric network for the site This Virtual network type is used for exposing virtual host on IP Fabric network on the VER site or for endpoint in IP Fabric network

<a id="advertise-custom-advertise-where-site-site"></a>&#x2022; [`site`](#advertise-custom-advertise-where-site-site) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Site](#advertise-custom-advertise-where-site-site) below.

#### Advertise Custom Advertise Where Site Site

A [`site`](#advertise-custom-advertise-where-site-site) block (within [`advertise_custom.advertise_where.site`](#advertise-custom-advertise-where-site)) supports the following:

<a id="advertise-custom-advertise-where-site-site-name"></a>&#x2022; [`name`](#advertise-custom-advertise-where-site-site-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="advertise-custom-advertise-where-site-site-namespace"></a>&#x2022; [`namespace`](#advertise-custom-advertise-where-site-site-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="advertise-custom-advertise-where-site-site-tenant"></a>&#x2022; [`tenant`](#advertise-custom-advertise-where-site-site-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Advertise Custom Advertise Where Virtual Network

A [`virtual_network`](#advertise-custom-advertise-where-virtual-network) block (within [`advertise_custom.advertise_where`](#advertise-custom-advertise-where)) supports the following:

<a id="advertise-custom-advertise-where-virtual-network-default-v6-vip"></a>&#x2022; [`default_v6_vip`](#advertise-custom-advertise-where-virtual-network-default-v6-vip) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="advertise-custom-advertise-where-virtual-network-default-vip"></a>&#x2022; [`default_vip`](#advertise-custom-advertise-where-virtual-network-default-vip) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="advertise-custom-advertise-where-virtual-network-specific-v6-vip"></a>&#x2022; [`specific_v6_vip`](#advertise-custom-advertise-where-virtual-network-specific-v6-vip) - Optional String<br>Specific V6 VIP. Use given IPv6 address as VIP on virtual Network

<a id="advertise-custom-advertise-where-virtual-network-specific-vip"></a>&#x2022; [`specific_vip`](#advertise-custom-advertise-where-virtual-network-specific-vip) - Optional String<br>Specific V4 VIP. Use given IPv4 address as VIP on virtual Network

<a id="advertise-custom-advertise-where-virtual-network-virtual-network"></a>&#x2022; [`virtual_network`](#advertise-custom-advertise-where-virtual-network-virtual-network) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Virtual Network](#advertise-custom-advertise-where-virtual-network-virtual-network) below.

#### Advertise Custom Advertise Where Virtual Network Virtual Network

A [`virtual_network`](#advertise-custom-advertise-where-virtual-network-virtual-network) block (within [`advertise_custom.advertise_where.virtual_network`](#advertise-custom-advertise-where-virtual-network)) supports the following:

<a id="advertise-custom-advertise-where-virtual-network-virtual-network-name"></a>&#x2022; [`name`](#advertise-custom-advertise-where-virtual-network-virtual-network-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="advertise-custom-advertise-where-virtual-network-virtual-network-namespace"></a>&#x2022; [`namespace`](#advertise-custom-advertise-where-virtual-network-virtual-network-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="advertise-custom-advertise-where-virtual-network-virtual-network-tenant"></a>&#x2022; [`tenant`](#advertise-custom-advertise-where-virtual-network-virtual-network-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Advertise Custom Advertise Where Virtual Site

A [`virtual_site`](#advertise-custom-advertise-where-virtual-site) block (within [`advertise_custom.advertise_where`](#advertise-custom-advertise-where)) supports the following:

<a id="advertise-custom-advertise-where-virtual-site-network"></a>&#x2022; [`network`](#advertise-custom-advertise-where-virtual-site-network) - Optional String  Defaults to `SITE_NETWORK_INSIDE_AND_OUTSIDE`<br>Possible values are `SITE_NETWORK_INSIDE_AND_OUTSIDE`, `SITE_NETWORK_INSIDE`, `SITE_NETWORK_OUTSIDE`, `SITE_NETWORK_SERVICE`, `SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP`, `SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP`, `SITE_NETWORK_IP_FABRIC`<br>Site Network. This defines network types to be used on site All inside and outside networks. All inside and outside networks with internet VIP support. All inside networks. All outside networks. All outside networks with internet VIP support. vK8s service network. - SITE_NETWORK_IP_FABRIC: VER IP Fabric network for the site This Virtual network type is used for exposing virtual host on IP Fabric network on the VER site or for endpoint in IP Fabric network

<a id="advertise-custom-advertise-where-virtual-site-virtual-site"></a>&#x2022; [`virtual_site`](#advertise-custom-advertise-where-virtual-site-virtual-site) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Virtual Site](#advertise-custom-advertise-where-virtual-site-virtual-site) below.

#### Advertise Custom Advertise Where Virtual Site Virtual Site

A [`virtual_site`](#advertise-custom-advertise-where-virtual-site-virtual-site) block (within [`advertise_custom.advertise_where.virtual_site`](#advertise-custom-advertise-where-virtual-site)) supports the following:

<a id="advertise-custom-advertise-where-virtual-site-virtual-site-name"></a>&#x2022; [`name`](#advertise-custom-advertise-where-virtual-site-virtual-site-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="advertise-custom-advertise-where-virtual-site-virtual-site-namespace"></a>&#x2022; [`namespace`](#advertise-custom-advertise-where-virtual-site-virtual-site-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="advertise-custom-advertise-where-virtual-site-virtual-site-tenant"></a>&#x2022; [`tenant`](#advertise-custom-advertise-where-virtual-site-virtual-site-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Advertise Custom Advertise Where Virtual Site With VIP

A [`virtual_site_with_vip`](#advertise-custom-advertise-where-virtual-site-with-vip) block (within [`advertise_custom.advertise_where`](#advertise-custom-advertise-where)) supports the following:

<a id="advertise-custom-advertise-where-virtual-site-with-vip-ip"></a>&#x2022; [`ip`](#advertise-custom-advertise-where-virtual-site-with-vip-ip) - Optional String<br>IP Address. Use given IP address as VIP on the site

<a id="advertise-custom-advertise-where-virtual-site-with-vip-network"></a>&#x2022; [`network`](#advertise-custom-advertise-where-virtual-site-with-vip-network) - Optional String  Defaults to `SITE_NETWORK_SPECIFIED_VIP_OUTSIDE`<br>Possible values are `SITE_NETWORK_SPECIFIED_VIP_OUTSIDE`, `SITE_NETWORK_SPECIFIED_VIP_INSIDE`<br>Site Network. This defines network types to be used on virtual-site with specified VIP All outside networks. All inside networks

<a id="advertise-custom-advertise-where-virtual-site-with-vip-virtual-site"></a>&#x2022; [`virtual_site`](#advertise-custom-advertise-where-virtual-site-with-vip-virtual-site) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Virtual Site](#advertise-custom-advertise-where-virtual-site-with-vip-virtual-site) below.

#### Advertise Custom Advertise Where Virtual Site With VIP Virtual Site

A [`virtual_site`](#advertise-custom-advertise-where-virtual-site-with-vip-virtual-site) block (within [`advertise_custom.advertise_where.virtual_site_with_vip`](#advertise-custom-advertise-where-virtual-site-with-vip)) supports the following:

<a id="advertise-custom-advertise-where-virtual-site-with-vip-virtual-site-name"></a>&#x2022; [`name`](#advertise-custom-advertise-where-virtual-site-with-vip-virtual-site-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="advertise-custom-advertise-where-virtual-site-with-vip-virtual-site-namespace"></a>&#x2022; [`namespace`](#advertise-custom-advertise-where-virtual-site-with-vip-virtual-site-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="advertise-custom-advertise-where-virtual-site-with-vip-virtual-site-tenant"></a>&#x2022; [`tenant`](#advertise-custom-advertise-where-virtual-site-with-vip-virtual-site-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Advertise Custom Advertise Where Vk8s Service

A [`vk8s_service`](#advertise-custom-advertise-where-vk8s-service) block (within [`advertise_custom.advertise_where`](#advertise-custom-advertise-where)) supports the following:

<a id="advertise-custom-advertise-where-vk8s-service-site"></a>&#x2022; [`site`](#advertise-custom-advertise-where-vk8s-service-site) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Site](#advertise-custom-advertise-where-vk8s-service-site) below.

<a id="advertise-custom-advertise-where-vk8s-service-virtual-site"></a>&#x2022; [`virtual_site`](#advertise-custom-advertise-where-vk8s-service-virtual-site) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Virtual Site](#advertise-custom-advertise-where-vk8s-service-virtual-site) below.

#### Advertise Custom Advertise Where Vk8s Service Site

A [`site`](#advertise-custom-advertise-where-vk8s-service-site) block (within [`advertise_custom.advertise_where.vk8s_service`](#advertise-custom-advertise-where-vk8s-service)) supports the following:

<a id="advertise-custom-advertise-where-vk8s-service-site-name"></a>&#x2022; [`name`](#advertise-custom-advertise-where-vk8s-service-site-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="advertise-custom-advertise-where-vk8s-service-site-namespace"></a>&#x2022; [`namespace`](#advertise-custom-advertise-where-vk8s-service-site-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="advertise-custom-advertise-where-vk8s-service-site-tenant"></a>&#x2022; [`tenant`](#advertise-custom-advertise-where-vk8s-service-site-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Advertise Custom Advertise Where Vk8s Service Virtual Site

A [`virtual_site`](#advertise-custom-advertise-where-vk8s-service-virtual-site) block (within [`advertise_custom.advertise_where.vk8s_service`](#advertise-custom-advertise-where-vk8s-service)) supports the following:

<a id="advertise-custom-advertise-where-vk8s-service-virtual-site-name"></a>&#x2022; [`name`](#advertise-custom-advertise-where-vk8s-service-virtual-site-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="advertise-custom-advertise-where-vk8s-service-virtual-site-namespace"></a>&#x2022; [`namespace`](#advertise-custom-advertise-where-vk8s-service-virtual-site-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="advertise-custom-advertise-where-vk8s-service-virtual-site-tenant"></a>&#x2022; [`tenant`](#advertise-custom-advertise-where-vk8s-service-virtual-site-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

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

<a id="api-protection-rules-api-endpoint-rules-action"></a>&#x2022; [`action`](#api-protection-rules-api-endpoint-rules-action) - Optional Block<br>API Protection Rule Action. The action to take if the input request matches the rule<br>See [Action](#api-protection-rules-api-endpoint-rules-action) below.

<a id="api-protection-rules-api-endpoint-rules-any-domain"></a>&#x2022; [`any_domain`](#api-protection-rules-api-endpoint-rules-any-domain) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-protection-rules-api-endpoint-rules-api-endpoint-method"></a>&#x2022; [`api_endpoint_method`](#api-protection-rules-api-endpoint-rules-api-endpoint-method) - Optional Block<br>HTTP Method Matcher. A HTTP method matcher specifies a list of methods to match an input HTTP method. The match is considered successful if the input method is a member of the list. The result of the match based on the method list is inverted if invert_matcher is true<br>See [API Endpoint Method](#api-protection-rules-api-endpoint-rules-api-endpoint-method) below.

<a id="api-protection-rules-api-endpoint-rules-api-endpoint-path"></a>&#x2022; [`api_endpoint_path`](#api-protection-rules-api-endpoint-rules-api-endpoint-path) - Optional String<br>API Endpoint. The endpoint (path) of the request

<a id="api-protection-rules-api-endpoint-rules-client-matcher"></a>&#x2022; [`client_matcher`](#api-protection-rules-api-endpoint-rules-client-matcher) - Optional Block<br>Client Matcher. Client conditions for matching a rule<br>See [Client Matcher](#api-protection-rules-api-endpoint-rules-client-matcher) below.

<a id="api-protection-rules-api-endpoint-rules-metadata"></a>&#x2022; [`metadata`](#api-protection-rules-api-endpoint-rules-metadata) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#api-protection-rules-api-endpoint-rules-metadata) below.

<a id="api-protection-rules-api-endpoint-rules-request-matcher"></a>&#x2022; [`request_matcher`](#api-protection-rules-api-endpoint-rules-request-matcher) - Optional Block<br>Request Matcher. Request conditions for matching a rule<br>See [Request Matcher](#api-protection-rules-api-endpoint-rules-request-matcher) below.

<a id="api-protection-rules-api-endpoint-rules-specific-domain"></a>&#x2022; [`specific_domain`](#api-protection-rules-api-endpoint-rules-specific-domain) - Optional String<br>Specific Domain. The rule will apply for a specific domain. For example: API.example.com

#### API Protection Rules API Endpoint Rules Action

An [`action`](#api-protection-rules-api-endpoint-rules-action) block (within [`api_protection_rules.api_endpoint_rules`](#api-protection-rules-api-endpoint-rules)) supports the following:

<a id="api-protection-rules-api-endpoint-rules-action-allow"></a>&#x2022; [`allow`](#api-protection-rules-api-endpoint-rules-action-allow) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-protection-rules-api-endpoint-rules-action-deny"></a>&#x2022; [`deny`](#api-protection-rules-api-endpoint-rules-action-deny) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### API Protection Rules API Endpoint Rules API Endpoint Method

An [`api_endpoint_method`](#api-protection-rules-api-endpoint-rules-api-endpoint-method) block (within [`api_protection_rules.api_endpoint_rules`](#api-protection-rules-api-endpoint-rules)) supports the following:

<a id="api-protection-rules-api-endpoint-rules-api-endpoint-method-invert-matcher"></a>&#x2022; [`invert_matcher`](#api-protection-rules-api-endpoint-rules-api-endpoint-method-invert-matcher) - Optional Bool<br>Invert Method Matcher. Invert the match result

<a id="api-protection-rules-api-endpoint-rules-api-endpoint-method-methods"></a>&#x2022; [`methods`](#api-protection-rules-api-endpoint-rules-api-endpoint-method-methods) - Optional List  Defaults to `ANY`<br>Possible values are `ANY`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`, `COPY`<br>Method List. List of methods values to match against

#### API Protection Rules API Endpoint Rules Client Matcher

A [`client_matcher`](#api-protection-rules-api-endpoint-rules-client-matcher) block (within [`api_protection_rules.api_endpoint_rules`](#api-protection-rules-api-endpoint-rules)) supports the following:

<a id="api-protection-rules-api-endpoint-rules-client-matcher-any-client"></a>&#x2022; [`any_client`](#api-protection-rules-api-endpoint-rules-client-matcher-any-client) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-protection-rules-api-endpoint-rules-client-matcher-any-ip"></a>&#x2022; [`any_ip`](#api-protection-rules-api-endpoint-rules-client-matcher-any-ip) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-protection-rules-api-endpoint-rules-client-matcher-asn-list"></a>&#x2022; [`asn_list`](#api-protection-rules-api-endpoint-rules-client-matcher-asn-list) - Optional Block<br>ASN Match List. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer<br>See [Asn List](#api-protection-rules-api-endpoint-rules-client-matcher-asn-list) below.

<a id="api-protection-rules-api-endpoint-rules-client-matcher-asn-matcher"></a>&#x2022; [`asn_matcher`](#api-protection-rules-api-endpoint-rules-client-matcher-asn-matcher) - Optional Block<br>ASN Matcher. Match any AS number contained in the list of bgp_asn_sets<br>See [Asn Matcher](#api-protection-rules-api-endpoint-rules-client-matcher-asn-matcher) below.

<a id="api-protection-rules-api-endpoint-rules-client-matcher-client-selector"></a>&#x2022; [`client_selector`](#api-protection-rules-api-endpoint-rules-client-matcher-client-selector) - Optional Block<br>Label Selector. This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expresssions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects. Label selector is immutable. expressions is a list of strings of label selection expression. Each string has ',' separated values which are 'AND' and all strings are logically 'OR'. BNF for expression string `<selector-syntax>` ::= `<requirement>` | `<requirement>` ',' `<selector-syntax>` `<requirement>` ::= [!] KEY [ `<set-based-restriction>` | `<exact-match-restriction>` ] `<set-based-restriction>` ::= '' | `<inclusion-exclusion>` `<value-set>` `<inclusion-exclusion>` ::= `<inclusion>` | `<exclusion>` `<exclusion>` ::= 'notin' `<inclusion>` ::= 'in' `<value-set>` ::= '(' `<values>` ')' `<values>` ::= VALUE | VALUE ',' `<values>` `<exact-match-restriction>` ::= ['='|'=='|'!='] VALUE<br>See [Client Selector](#api-protection-rules-api-endpoint-rules-client-matcher-client-selector) below.

<a id="api-protection-rules-api-endpoint-rules-client-matcher-ip-matcher"></a>&#x2022; [`ip_matcher`](#api-protection-rules-api-endpoint-rules-client-matcher-ip-matcher) - Optional Block<br>IP Prefix Matcher. Match any IP prefix contained in the list of ip_prefix_sets. The result of the match is inverted if invert_matcher is true<br>See [IP Matcher](#api-protection-rules-api-endpoint-rules-client-matcher-ip-matcher) below.

<a id="api-protection-rules-api-endpoint-rules-client-matcher-ip-prefix-list"></a>&#x2022; [`ip_prefix_list`](#api-protection-rules-api-endpoint-rules-client-matcher-ip-prefix-list) - Optional Block<br>IP Prefix Match List. List of IP Prefix strings to match against<br>See [IP Prefix List](#api-protection-rules-api-endpoint-rules-client-matcher-ip-prefix-list) below.

<a id="api-protection-rules-api-endpoint-rules-client-matcher-ip-threat-category-list"></a>&#x2022; [`ip_threat_category_list`](#api-protection-rules-api-endpoint-rules-client-matcher-ip-threat-category-list) - Optional Block<br>IP Threat Category List Type. List of IP threat categories<br>See [IP Threat Category List](#api-protection-rules-api-endpoint-rules-client-matcher-ip-threat-category-list) below.

<a id="api-protection-rules-api-endpoint-rules-client-matcher-tls-fingerprint-matcher"></a>&#x2022; [`tls_fingerprint_matcher`](#api-protection-rules-api-endpoint-rules-client-matcher-tls-fingerprint-matcher) - Optional Block<br>TLS Fingerprint Matcher. A TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of supported positve match criteria includes a list of known classes of TLS fingerprints and a list of exact values. The match is considered successful if either of these positive criteria are satisfied and the input fingerprint is not one of the excluded values<br>See [TLS Fingerprint Matcher](#api-protection-rules-api-endpoint-rules-client-matcher-tls-fingerprint-matcher) below.

#### API Protection Rules API Endpoint Rules Client Matcher Asn List

An [`asn_list`](#api-protection-rules-api-endpoint-rules-client-matcher-asn-list) block (within [`api_protection_rules.api_endpoint_rules.client_matcher`](#api-protection-rules-api-endpoint-rules-client-matcher)) supports the following:

<a id="api-protection-rules-api-endpoint-rules-client-matcher-asn-list-as-numbers"></a>&#x2022; [`as_numbers`](#api-protection-rules-api-endpoint-rules-client-matcher-asn-list-as-numbers) - Optional List<br>AS Numbers. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer

#### API Protection Rules API Endpoint Rules Client Matcher Asn Matcher

An [`asn_matcher`](#api-protection-rules-api-endpoint-rules-client-matcher-asn-matcher) block (within [`api_protection_rules.api_endpoint_rules.client_matcher`](#api-protection-rules-api-endpoint-rules-client-matcher)) supports the following:

<a id="api-protection-rules-api-endpoint-rules-client-matcher-asn-matcher-asn-sets"></a>&#x2022; [`asn_sets`](#api-protection-rules-api-endpoint-rules-client-matcher-asn-matcher-asn-sets) - Optional Block<br>BGP ASN Sets. A list of references to bgp_asn_set objects<br>See [Asn Sets](#api-protection-rules-api-endpoint-rules-client-matcher-asn-matcher-asn-sets) below.

#### API Protection Rules API Endpoint Rules Client Matcher Asn Matcher Asn Sets

An [`asn_sets`](#api-protection-rules-api-endpoint-rules-client-matcher-asn-matcher-asn-sets) block (within [`api_protection_rules.api_endpoint_rules.client_matcher.asn_matcher`](#api-protection-rules-api-endpoint-rules-client-matcher-asn-matcher)) supports the following:

<a id="api-protection-rules-api-endpoint-rules-client-matcher-asn-matcher-asn-sets-kind"></a>&#x2022; [`kind`](#api-protection-rules-api-endpoint-rules-client-matcher-asn-matcher-asn-sets-kind) - Optional String<br>Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route')

<a id="api-protection-rules-api-endpoint-rules-client-matcher-asn-matcher-asn-sets-name"></a>&#x2022; [`name`](#api-protection-rules-api-endpoint-rules-client-matcher-asn-matcher-asn-sets-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="api-protection-rules-api-endpoint-rules-client-matcher-asn-matcher-asn-sets-namespace"></a>&#x2022; [`namespace`](#api-protection-rules-api-endpoint-rules-client-matcher-asn-matcher-asn-sets-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="api-protection-rules-api-endpoint-rules-client-matcher-asn-matcher-asn-sets-tenant"></a>&#x2022; [`tenant`](#api-protection-rules-api-endpoint-rules-client-matcher-asn-matcher-asn-sets-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="api-protection-rules-api-endpoint-rules-client-matcher-asn-matcher-asn-sets-uid"></a>&#x2022; [`uid`](#api-protection-rules-api-endpoint-rules-client-matcher-asn-matcher-asn-sets-uid) - Optional String<br>UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid

#### API Protection Rules API Endpoint Rules Client Matcher Client Selector

A [`client_selector`](#api-protection-rules-api-endpoint-rules-client-matcher-client-selector) block (within [`api_protection_rules.api_endpoint_rules.client_matcher`](#api-protection-rules-api-endpoint-rules-client-matcher)) supports the following:

<a id="api-protection-rules-api-endpoint-rules-client-matcher-client-selector-expressions"></a>&#x2022; [`expressions`](#api-protection-rules-api-endpoint-rules-client-matcher-client-selector-expressions) - Optional List<br>Selector Expression. expressions contains the kubernetes style label expression for selections

#### API Protection Rules API Endpoint Rules Client Matcher IP Matcher

An [`ip_matcher`](#api-protection-rules-api-endpoint-rules-client-matcher-ip-matcher) block (within [`api_protection_rules.api_endpoint_rules.client_matcher`](#api-protection-rules-api-endpoint-rules-client-matcher)) supports the following:

<a id="api-protection-rules-api-endpoint-rules-client-matcher-ip-matcher-invert-matcher"></a>&#x2022; [`invert_matcher`](#api-protection-rules-api-endpoint-rules-client-matcher-ip-matcher-invert-matcher) - Optional Bool<br>Invert IP Matcher. Invert the match result

<a id="api-protection-rules-api-endpoint-rules-client-matcher-ip-matcher-prefix-sets"></a>&#x2022; [`prefix_sets`](#api-protection-rules-api-endpoint-rules-client-matcher-ip-matcher-prefix-sets) - Optional Block<br>IP Prefix Sets. A list of references to ip_prefix_set objects<br>See [Prefix Sets](#api-protection-rules-api-endpoint-rules-client-matcher-ip-matcher-prefix-sets) below.

#### API Protection Rules API Endpoint Rules Client Matcher IP Matcher Prefix Sets

A [`prefix_sets`](#api-protection-rules-api-endpoint-rules-client-matcher-ip-matcher-prefix-sets) block (within [`api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher`](#api-protection-rules-api-endpoint-rules-client-matcher-ip-matcher)) supports the following:

<a id="api-protection-rules-api-endpoint-rules-client-matcher-ip-matcher-prefix-sets-kind"></a>&#x2022; [`kind`](#api-protection-rules-api-endpoint-rules-client-matcher-ip-matcher-prefix-sets-kind) - Optional String<br>Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route')

<a id="api-protection-rules-api-endpoint-rules-client-matcher-ip-matcher-prefix-sets-name"></a>&#x2022; [`name`](#api-protection-rules-api-endpoint-rules-client-matcher-ip-matcher-prefix-sets-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="api-protection-rules-api-endpoint-rules-client-matcher-ip-matcher-prefix-sets-namespace"></a>&#x2022; [`namespace`](#api-protection-rules-api-endpoint-rules-client-matcher-ip-matcher-prefix-sets-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="api-protection-rules-api-endpoint-rules-client-matcher-ip-matcher-prefix-sets-tenant"></a>&#x2022; [`tenant`](#api-protection-rules-api-endpoint-rules-client-matcher-ip-matcher-prefix-sets-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="api-protection-rules-api-endpoint-rules-client-matcher-ip-matcher-prefix-sets-uid"></a>&#x2022; [`uid`](#api-protection-rules-api-endpoint-rules-client-matcher-ip-matcher-prefix-sets-uid) - Optional String<br>UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid

#### API Protection Rules API Endpoint Rules Client Matcher IP Prefix List

An [`ip_prefix_list`](#api-protection-rules-api-endpoint-rules-client-matcher-ip-prefix-list) block (within [`api_protection_rules.api_endpoint_rules.client_matcher`](#api-protection-rules-api-endpoint-rules-client-matcher)) supports the following:

<a id="api-protection-rules-api-endpoint-rules-client-matcher-ip-prefix-list-invert-match"></a>&#x2022; [`invert_match`](#api-protection-rules-api-endpoint-rules-client-matcher-ip-prefix-list-invert-match) - Optional Bool<br>Invert Match Result. Invert the match result

<a id="api-protection-rules-api-endpoint-rules-client-matcher-ip-prefix-list-ip-prefixes"></a>&#x2022; [`ip_prefixes`](#api-protection-rules-api-endpoint-rules-client-matcher-ip-prefix-list-ip-prefixes) - Optional List<br>IPv4 Prefix List. List of IPv4 prefix strings

#### API Protection Rules API Endpoint Rules Client Matcher IP Threat Category List

An [`ip_threat_category_list`](#api-protection-rules-api-endpoint-rules-client-matcher-ip-threat-category-list) block (within [`api_protection_rules.api_endpoint_rules.client_matcher`](#api-protection-rules-api-endpoint-rules-client-matcher)) supports the following:

<a id="api-protection-rules-api-endpoint-rules-client-matcher-ip-threat-category-list-ip-threat-categories"></a>&#x2022; [`ip_threat_categories`](#api-protection-rules-api-endpoint-rules-client-matcher-ip-threat-category-list-ip-threat-categories) - Optional List  Defaults to `SPAM_SOURCES`<br>Possible values are `SPAM_SOURCES`, `WINDOWS_EXPLOITS`, `WEB_ATTACKS`, `BOTNETS`, `SCANNERS`, `REPUTATION`, `PHISHING`, `PROXY`, `MOBILE_THREATS`, `TOR_PROXY`, `DENIAL_OF_SERVICE`, `NETWORK`<br>List of IP Threat Categories to choose. The IP threat categories is obtained from the list and is used to auto-generate equivalent label selection expressions

#### API Protection Rules API Endpoint Rules Client Matcher TLS Fingerprint Matcher

A [`tls_fingerprint_matcher`](#api-protection-rules-api-endpoint-rules-client-matcher-tls-fingerprint-matcher) block (within [`api_protection_rules.api_endpoint_rules.client_matcher`](#api-protection-rules-api-endpoint-rules-client-matcher)) supports the following:

<a id="api-protection-rules-api-endpoint-rules-client-matcher-tls-fingerprint-matcher-classes"></a>&#x2022; [`classes`](#api-protection-rules-api-endpoint-rules-client-matcher-tls-fingerprint-matcher-classes) - Optional List  Defaults to `TLS_FINGERPRINT_NONE`<br>Possible values are `TLS_FINGERPRINT_NONE`, `ANY_MALICIOUS_FINGERPRINT`, `ADWARE`, `ADWIND`, `DRIDEX`, `GOOTKIT`, `GOZI`, `JBIFROST`, `QUAKBOT`, `RANSOMWARE`, `TROLDESH`, `TOFSEE`, `TORRENTLOCKER`, `TRICKBOT`<br>TLS fingerprint classes. A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against

<a id="api-protection-rules-api-endpoint-rules-client-matcher-tls-fingerprint-matcher-exact-values"></a>&#x2022; [`exact_values`](#api-protection-rules-api-endpoint-rules-client-matcher-tls-fingerprint-matcher-exact-values) - Optional List<br>Exact Values. A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against

<a id="api-protection-rules-api-endpoint-rules-client-matcher-tls-fingerprint-matcher-excluded-values"></a>&#x2022; [`excluded_values`](#api-protection-rules-api-endpoint-rules-client-matcher-tls-fingerprint-matcher-excluded-values) - Optional List<br>Excluded Values. A list of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can be used to skip known false positives when using one or more known TLS fingerprint classes in the enclosing matcher

#### API Protection Rules API Endpoint Rules Metadata

A [`metadata`](#api-protection-rules-api-endpoint-rules-metadata) block (within [`api_protection_rules.api_endpoint_rules`](#api-protection-rules-api-endpoint-rules)) supports the following:

<a id="api-protection-rules-api-endpoint-rules-metadata-description-spec"></a>&#x2022; [`description_spec`](#api-protection-rules-api-endpoint-rules-metadata-description-spec) - Optional String<br>Description. Human readable description

<a id="api-protection-rules-api-endpoint-rules-metadata-name"></a>&#x2022; [`name`](#api-protection-rules-api-endpoint-rules-metadata-name) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### API Protection Rules API Endpoint Rules Request Matcher

A [`request_matcher`](#api-protection-rules-api-endpoint-rules-request-matcher) block (within [`api_protection_rules.api_endpoint_rules`](#api-protection-rules-api-endpoint-rules)) supports the following:

<a id="api-protection-rules-api-endpoint-rules-request-matcher-cookie-matchers"></a>&#x2022; [`cookie_matchers`](#api-protection-rules-api-endpoint-rules-request-matcher-cookie-matchers) - Optional Block<br>Cookie Matchers. A list of predicates for all cookies that need to be matched. The criteria for matching each cookie is described in individual instances of CookieMatcherType. The actual cookie values are extracted from the request API as a list of strings for each cookie name. Note that all specified cookie matcher predicates must evaluate to true<br>See [Cookie Matchers](#api-protection-rules-api-endpoint-rules-request-matcher-cookie-matchers) below.

<a id="api-protection-rules-api-endpoint-rules-request-matcher-headers"></a>&#x2022; [`headers`](#api-protection-rules-api-endpoint-rules-request-matcher-headers) - Optional Block<br>HTTP Headers. A list of predicates for various HTTP headers that need to match. The criteria for matching each HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values are extracted from the request API as a list of strings for each HTTP header type. Note that all specified header predicates must evaluate to true<br>See [Headers](#api-protection-rules-api-endpoint-rules-request-matcher-headers) below.

<a id="api-protection-rules-api-endpoint-rules-request-matcher-jwt-claims"></a>&#x2022; [`jwt_claims`](#api-protection-rules-api-endpoint-rules-request-matcher-jwt-claims) - Optional Block<br>JWT Claims. A list of predicates for various JWT claims that need to match. The criteria for matching each JWT claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are extracted from the JWT payload as a list of strings. Note that all specified JWT claim predicates must evaluate to true. Note that this feature only works on LBs with JWT Validation feature enabled<br>See [JWT Claims](#api-protection-rules-api-endpoint-rules-request-matcher-jwt-claims) below.

<a id="api-protection-rules-api-endpoint-rules-request-matcher-query-params"></a>&#x2022; [`query_params`](#api-protection-rules-api-endpoint-rules-request-matcher-query-params) - Optional Block<br>HTTP Query Parameters. A list of predicates for all query parameters that need to be matched. The criteria for matching each query parameter are described in individual instances of QueryParameterMatcherType. The actual query parameter values are extracted from the request API as a list of strings for each query parameter name. Note that all specified query parameter predicates must evaluate to true<br>See [Query Params](#api-protection-rules-api-endpoint-rules-request-matcher-query-params) below.

#### API Protection Rules API Endpoint Rules Request Matcher Cookie Matchers

A [`cookie_matchers`](#api-protection-rules-api-endpoint-rules-request-matcher-cookie-matchers) block (within [`api_protection_rules.api_endpoint_rules.request_matcher`](#api-protection-rules-api-endpoint-rules-request-matcher)) supports the following:

<a id="api-protection-rules-api-endpoint-rules-request-matcher-cookie-matchers-check-not-present"></a>&#x2022; [`check_not_present`](#api-protection-rules-api-endpoint-rules-request-matcher-cookie-matchers-check-not-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-protection-rules-api-endpoint-rules-request-matcher-cookie-matchers-check-present"></a>&#x2022; [`check_present`](#api-protection-rules-api-endpoint-rules-request-matcher-cookie-matchers-check-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-protection-rules-api-endpoint-rules-request-matcher-cookie-matchers-invert-matcher"></a>&#x2022; [`invert_matcher`](#api-protection-rules-api-endpoint-rules-request-matcher-cookie-matchers-invert-matcher) - Optional Bool<br>Invert Matcher. Invert Match of the expression defined

<a id="api-protection-rules-api-endpoint-rules-request-matcher-cookie-matchers-item"></a>&#x2022; [`item`](#api-protection-rules-api-endpoint-rules-request-matcher-cookie-matchers-item) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#api-protection-rules-api-endpoint-rules-request-matcher-cookie-matchers-item) below.

<a id="api-protection-rules-api-endpoint-rules-request-matcher-cookie-matchers-name"></a>&#x2022; [`name`](#api-protection-rules-api-endpoint-rules-request-matcher-cookie-matchers-name) - Optional String<br>Cookie Name. A case-sensitive cookie name

#### API Protection Rules API Endpoint Rules Request Matcher Cookie Matchers Item

An [`item`](#api-protection-rules-api-endpoint-rules-request-matcher-cookie-matchers-item) block (within [`api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers`](#api-protection-rules-api-endpoint-rules-request-matcher-cookie-matchers)) supports the following:

<a id="api-protection-rules-api-endpoint-rules-request-matcher-cookie-matchers-item-exact-values"></a>&#x2022; [`exact_values`](#api-protection-rules-api-endpoint-rules-request-matcher-cookie-matchers-item-exact-values) - Optional List<br>Exact Values. A list of exact values to match the input against

<a id="api-protection-rules-api-endpoint-rules-request-matcher-cookie-matchers-item-regex-values"></a>&#x2022; [`regex_values`](#api-protection-rules-api-endpoint-rules-request-matcher-cookie-matchers-item-regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="api-protection-rules-api-endpoint-rules-request-matcher-cookie-matchers-item-transformers"></a>&#x2022; [`transformers`](#api-protection-rules-api-endpoint-rules-request-matcher-cookie-matchers-item-transformers) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### API Protection Rules API Endpoint Rules Request Matcher Headers

A [`headers`](#api-protection-rules-api-endpoint-rules-request-matcher-headers) block (within [`api_protection_rules.api_endpoint_rules.request_matcher`](#api-protection-rules-api-endpoint-rules-request-matcher)) supports the following:

<a id="api-protection-rules-api-endpoint-rules-request-matcher-headers-check-not-present"></a>&#x2022; [`check_not_present`](#api-protection-rules-api-endpoint-rules-request-matcher-headers-check-not-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-protection-rules-api-endpoint-rules-request-matcher-headers-check-present"></a>&#x2022; [`check_present`](#api-protection-rules-api-endpoint-rules-request-matcher-headers-check-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-protection-rules-api-endpoint-rules-request-matcher-headers-invert-matcher"></a>&#x2022; [`invert_matcher`](#api-protection-rules-api-endpoint-rules-request-matcher-headers-invert-matcher) - Optional Bool<br>Invert Header Matcher. Invert the match result

<a id="api-protection-rules-api-endpoint-rules-request-matcher-headers-item"></a>&#x2022; [`item`](#api-protection-rules-api-endpoint-rules-request-matcher-headers-item) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#api-protection-rules-api-endpoint-rules-request-matcher-headers-item) below.

<a id="api-protection-rules-api-endpoint-rules-request-matcher-headers-name"></a>&#x2022; [`name`](#api-protection-rules-api-endpoint-rules-request-matcher-headers-name) - Optional String<br>Header Name. A case-insensitive HTTP header name

#### API Protection Rules API Endpoint Rules Request Matcher Headers Item

An [`item`](#api-protection-rules-api-endpoint-rules-request-matcher-headers-item) block (within [`api_protection_rules.api_endpoint_rules.request_matcher.headers`](#api-protection-rules-api-endpoint-rules-request-matcher-headers)) supports the following:

<a id="api-protection-rules-api-endpoint-rules-request-matcher-headers-item-exact-values"></a>&#x2022; [`exact_values`](#api-protection-rules-api-endpoint-rules-request-matcher-headers-item-exact-values) - Optional List<br>Exact Values. A list of exact values to match the input against

<a id="api-protection-rules-api-endpoint-rules-request-matcher-headers-item-regex-values"></a>&#x2022; [`regex_values`](#api-protection-rules-api-endpoint-rules-request-matcher-headers-item-regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="api-protection-rules-api-endpoint-rules-request-matcher-headers-item-transformers"></a>&#x2022; [`transformers`](#api-protection-rules-api-endpoint-rules-request-matcher-headers-item-transformers) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### API Protection Rules API Endpoint Rules Request Matcher JWT Claims

A [`jwt_claims`](#api-protection-rules-api-endpoint-rules-request-matcher-jwt-claims) block (within [`api_protection_rules.api_endpoint_rules.request_matcher`](#api-protection-rules-api-endpoint-rules-request-matcher)) supports the following:

<a id="api-protection-rules-api-endpoint-rules-request-matcher-jwt-claims-check-not-present"></a>&#x2022; [`check_not_present`](#api-protection-rules-api-endpoint-rules-request-matcher-jwt-claims-check-not-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-protection-rules-api-endpoint-rules-request-matcher-jwt-claims-check-present"></a>&#x2022; [`check_present`](#api-protection-rules-api-endpoint-rules-request-matcher-jwt-claims-check-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-protection-rules-api-endpoint-rules-request-matcher-jwt-claims-invert-matcher"></a>&#x2022; [`invert_matcher`](#api-protection-rules-api-endpoint-rules-request-matcher-jwt-claims-invert-matcher) - Optional Bool<br>Invert Matcher. Invert the match result

<a id="api-protection-rules-api-endpoint-rules-request-matcher-jwt-claims-item"></a>&#x2022; [`item`](#api-protection-rules-api-endpoint-rules-request-matcher-jwt-claims-item) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#api-protection-rules-api-endpoint-rules-request-matcher-jwt-claims-item) below.

<a id="api-protection-rules-api-endpoint-rules-request-matcher-jwt-claims-name"></a>&#x2022; [`name`](#api-protection-rules-api-endpoint-rules-request-matcher-jwt-claims-name) - Optional String<br>JWT Claim Name. JWT claim name

#### API Protection Rules API Endpoint Rules Request Matcher JWT Claims Item

An [`item`](#api-protection-rules-api-endpoint-rules-request-matcher-jwt-claims-item) block (within [`api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims`](#api-protection-rules-api-endpoint-rules-request-matcher-jwt-claims)) supports the following:

<a id="api-protection-rules-api-endpoint-rules-request-matcher-jwt-claims-item-exact-values"></a>&#x2022; [`exact_values`](#api-protection-rules-api-endpoint-rules-request-matcher-jwt-claims-item-exact-values) - Optional List<br>Exact Values. A list of exact values to match the input against

<a id="api-protection-rules-api-endpoint-rules-request-matcher-jwt-claims-item-regex-values"></a>&#x2022; [`regex_values`](#api-protection-rules-api-endpoint-rules-request-matcher-jwt-claims-item-regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="api-protection-rules-api-endpoint-rules-request-matcher-jwt-claims-item-transformers"></a>&#x2022; [`transformers`](#api-protection-rules-api-endpoint-rules-request-matcher-jwt-claims-item-transformers) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### API Protection Rules API Endpoint Rules Request Matcher Query Params

A [`query_params`](#api-protection-rules-api-endpoint-rules-request-matcher-query-params) block (within [`api_protection_rules.api_endpoint_rules.request_matcher`](#api-protection-rules-api-endpoint-rules-request-matcher)) supports the following:

<a id="api-protection-rules-api-endpoint-rules-request-matcher-query-params-check-not-present"></a>&#x2022; [`check_not_present`](#api-protection-rules-api-endpoint-rules-request-matcher-query-params-check-not-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-protection-rules-api-endpoint-rules-request-matcher-query-params-check-present"></a>&#x2022; [`check_present`](#api-protection-rules-api-endpoint-rules-request-matcher-query-params-check-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-protection-rules-api-endpoint-rules-request-matcher-query-params-invert-matcher"></a>&#x2022; [`invert_matcher`](#api-protection-rules-api-endpoint-rules-request-matcher-query-params-invert-matcher) - Optional Bool<br>Invert Query Parameter Matcher. Invert the match result

<a id="api-protection-rules-api-endpoint-rules-request-matcher-query-params-item"></a>&#x2022; [`item`](#api-protection-rules-api-endpoint-rules-request-matcher-query-params-item) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#api-protection-rules-api-endpoint-rules-request-matcher-query-params-item) below.

<a id="api-protection-rules-api-endpoint-rules-request-matcher-query-params-key"></a>&#x2022; [`key`](#api-protection-rules-api-endpoint-rules-request-matcher-query-params-key) - Optional String<br>Query Parameter Name. A case-sensitive HTTP query parameter name

#### API Protection Rules API Endpoint Rules Request Matcher Query Params Item

An [`item`](#api-protection-rules-api-endpoint-rules-request-matcher-query-params-item) block (within [`api_protection_rules.api_endpoint_rules.request_matcher.query_params`](#api-protection-rules-api-endpoint-rules-request-matcher-query-params)) supports the following:

<a id="api-protection-rules-api-endpoint-rules-request-matcher-query-params-item-exact-values"></a>&#x2022; [`exact_values`](#api-protection-rules-api-endpoint-rules-request-matcher-query-params-item-exact-values) - Optional List<br>Exact Values. A list of exact values to match the input against

<a id="api-protection-rules-api-endpoint-rules-request-matcher-query-params-item-regex-values"></a>&#x2022; [`regex_values`](#api-protection-rules-api-endpoint-rules-request-matcher-query-params-item-regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="api-protection-rules-api-endpoint-rules-request-matcher-query-params-item-transformers"></a>&#x2022; [`transformers`](#api-protection-rules-api-endpoint-rules-request-matcher-query-params-item-transformers) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### API Protection Rules API Groups Rules

An [`api_groups_rules`](#api-protection-rules-api-groups-rules) block (within [`api_protection_rules`](#api-protection-rules)) supports the following:

<a id="api-protection-rules-api-groups-rules-action"></a>&#x2022; [`action`](#api-protection-rules-api-groups-rules-action) - Optional Block<br>API Protection Rule Action. The action to take if the input request matches the rule<br>See [Action](#api-protection-rules-api-groups-rules-action) below.

<a id="api-protection-rules-api-groups-rules-any-domain"></a>&#x2022; [`any_domain`](#api-protection-rules-api-groups-rules-any-domain) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-protection-rules-api-groups-rules-api-group"></a>&#x2022; [`api_group`](#api-protection-rules-api-groups-rules-api-group) - Optional String<br>API Group. API groups derived from API Definition swaggers. For example oas-all-operations including all paths and methods from the swaggers, oas-base-urls covering all requests under base-paths from the swaggers. Custom groups can be created if user tags paths or operations with 'x-volterra-API-group' extensions inside swaggers

<a id="api-protection-rules-api-groups-rules-base-path"></a>&#x2022; [`base_path`](#api-protection-rules-api-groups-rules-base-path) - Optional String<br>Base Path. Prefix of the request path. For example: /v1

<a id="api-protection-rules-api-groups-rules-client-matcher"></a>&#x2022; [`client_matcher`](#api-protection-rules-api-groups-rules-client-matcher) - Optional Block<br>Client Matcher. Client conditions for matching a rule<br>See [Client Matcher](#api-protection-rules-api-groups-rules-client-matcher) below.

<a id="api-protection-rules-api-groups-rules-metadata"></a>&#x2022; [`metadata`](#api-protection-rules-api-groups-rules-metadata) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#api-protection-rules-api-groups-rules-metadata) below.

<a id="api-protection-rules-api-groups-rules-request-matcher"></a>&#x2022; [`request_matcher`](#api-protection-rules-api-groups-rules-request-matcher) - Optional Block<br>Request Matcher. Request conditions for matching a rule<br>See [Request Matcher](#api-protection-rules-api-groups-rules-request-matcher) below.

<a id="api-protection-rules-api-groups-rules-specific-domain"></a>&#x2022; [`specific_domain`](#api-protection-rules-api-groups-rules-specific-domain) - Optional String<br>Specific Domain. The rule will apply for a specific domain. For example: API.example.com

#### API Protection Rules API Groups Rules Action

An [`action`](#api-protection-rules-api-groups-rules-action) block (within [`api_protection_rules.api_groups_rules`](#api-protection-rules-api-groups-rules)) supports the following:

<a id="api-protection-rules-api-groups-rules-action-allow"></a>&#x2022; [`allow`](#api-protection-rules-api-groups-rules-action-allow) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-protection-rules-api-groups-rules-action-deny"></a>&#x2022; [`deny`](#api-protection-rules-api-groups-rules-action-deny) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### API Protection Rules API Groups Rules Client Matcher

A [`client_matcher`](#api-protection-rules-api-groups-rules-client-matcher) block (within [`api_protection_rules.api_groups_rules`](#api-protection-rules-api-groups-rules)) supports the following:

<a id="api-protection-rules-api-groups-rules-client-matcher-any-client"></a>&#x2022; [`any_client`](#api-protection-rules-api-groups-rules-client-matcher-any-client) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-protection-rules-api-groups-rules-client-matcher-any-ip"></a>&#x2022; [`any_ip`](#api-protection-rules-api-groups-rules-client-matcher-any-ip) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-protection-rules-api-groups-rules-client-matcher-asn-list"></a>&#x2022; [`asn_list`](#api-protection-rules-api-groups-rules-client-matcher-asn-list) - Optional Block<br>ASN Match List. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer<br>See [Asn List](#api-protection-rules-api-groups-rules-client-matcher-asn-list) below.

<a id="api-protection-rules-api-groups-rules-client-matcher-asn-matcher"></a>&#x2022; [`asn_matcher`](#api-protection-rules-api-groups-rules-client-matcher-asn-matcher) - Optional Block<br>ASN Matcher. Match any AS number contained in the list of bgp_asn_sets<br>See [Asn Matcher](#api-protection-rules-api-groups-rules-client-matcher-asn-matcher) below.

<a id="api-protection-rules-api-groups-rules-client-matcher-client-selector"></a>&#x2022; [`client_selector`](#api-protection-rules-api-groups-rules-client-matcher-client-selector) - Optional Block<br>Label Selector. This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expresssions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects. Label selector is immutable. expressions is a list of strings of label selection expression. Each string has ',' separated values which are 'AND' and all strings are logically 'OR'. BNF for expression string `<selector-syntax>` ::= `<requirement>` | `<requirement>` ',' `<selector-syntax>` `<requirement>` ::= [!] KEY [ `<set-based-restriction>` | `<exact-match-restriction>` ] `<set-based-restriction>` ::= '' | `<inclusion-exclusion>` `<value-set>` `<inclusion-exclusion>` ::= `<inclusion>` | `<exclusion>` `<exclusion>` ::= 'notin' `<inclusion>` ::= 'in' `<value-set>` ::= '(' `<values>` ')' `<values>` ::= VALUE | VALUE ',' `<values>` `<exact-match-restriction>` ::= ['='|'=='|'!='] VALUE<br>See [Client Selector](#api-protection-rules-api-groups-rules-client-matcher-client-selector) below.

<a id="api-protection-rules-api-groups-rules-client-matcher-ip-matcher"></a>&#x2022; [`ip_matcher`](#api-protection-rules-api-groups-rules-client-matcher-ip-matcher) - Optional Block<br>IP Prefix Matcher. Match any IP prefix contained in the list of ip_prefix_sets. The result of the match is inverted if invert_matcher is true<br>See [IP Matcher](#api-protection-rules-api-groups-rules-client-matcher-ip-matcher) below.

<a id="api-protection-rules-api-groups-rules-client-matcher-ip-prefix-list"></a>&#x2022; [`ip_prefix_list`](#api-protection-rules-api-groups-rules-client-matcher-ip-prefix-list) - Optional Block<br>IP Prefix Match List. List of IP Prefix strings to match against<br>See [IP Prefix List](#api-protection-rules-api-groups-rules-client-matcher-ip-prefix-list) below.

<a id="api-protection-rules-api-groups-rules-client-matcher-ip-threat-category-list"></a>&#x2022; [`ip_threat_category_list`](#api-protection-rules-api-groups-rules-client-matcher-ip-threat-category-list) - Optional Block<br>IP Threat Category List Type. List of IP threat categories<br>See [IP Threat Category List](#api-protection-rules-api-groups-rules-client-matcher-ip-threat-category-list) below.

<a id="api-protection-rules-api-groups-rules-client-matcher-tls-fingerprint-matcher"></a>&#x2022; [`tls_fingerprint_matcher`](#api-protection-rules-api-groups-rules-client-matcher-tls-fingerprint-matcher) - Optional Block<br>TLS Fingerprint Matcher. A TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of supported positve match criteria includes a list of known classes of TLS fingerprints and a list of exact values. The match is considered successful if either of these positive criteria are satisfied and the input fingerprint is not one of the excluded values<br>See [TLS Fingerprint Matcher](#api-protection-rules-api-groups-rules-client-matcher-tls-fingerprint-matcher) below.

#### API Protection Rules API Groups Rules Client Matcher Asn List

An [`asn_list`](#api-protection-rules-api-groups-rules-client-matcher-asn-list) block (within [`api_protection_rules.api_groups_rules.client_matcher`](#api-protection-rules-api-groups-rules-client-matcher)) supports the following:

<a id="api-protection-rules-api-groups-rules-client-matcher-asn-list-as-numbers"></a>&#x2022; [`as_numbers`](#api-protection-rules-api-groups-rules-client-matcher-asn-list-as-numbers) - Optional List<br>AS Numbers. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer

#### API Protection Rules API Groups Rules Client Matcher Asn Matcher

An [`asn_matcher`](#api-protection-rules-api-groups-rules-client-matcher-asn-matcher) block (within [`api_protection_rules.api_groups_rules.client_matcher`](#api-protection-rules-api-groups-rules-client-matcher)) supports the following:

<a id="api-protection-rules-api-groups-rules-client-matcher-asn-matcher-asn-sets"></a>&#x2022; [`asn_sets`](#api-protection-rules-api-groups-rules-client-matcher-asn-matcher-asn-sets) - Optional Block<br>BGP ASN Sets. A list of references to bgp_asn_set objects<br>See [Asn Sets](#api-protection-rules-api-groups-rules-client-matcher-asn-matcher-asn-sets) below.

#### API Protection Rules API Groups Rules Client Matcher Asn Matcher Asn Sets

An [`asn_sets`](#api-protection-rules-api-groups-rules-client-matcher-asn-matcher-asn-sets) block (within [`api_protection_rules.api_groups_rules.client_matcher.asn_matcher`](#api-protection-rules-api-groups-rules-client-matcher-asn-matcher)) supports the following:

<a id="api-protection-rules-api-groups-rules-client-matcher-asn-matcher-asn-sets-kind"></a>&#x2022; [`kind`](#api-protection-rules-api-groups-rules-client-matcher-asn-matcher-asn-sets-kind) - Optional String<br>Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route')

<a id="api-protection-rules-api-groups-rules-client-matcher-asn-matcher-asn-sets-name"></a>&#x2022; [`name`](#api-protection-rules-api-groups-rules-client-matcher-asn-matcher-asn-sets-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="api-protection-rules-api-groups-rules-client-matcher-asn-matcher-asn-sets-namespace"></a>&#x2022; [`namespace`](#api-protection-rules-api-groups-rules-client-matcher-asn-matcher-asn-sets-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="api-protection-rules-api-groups-rules-client-matcher-asn-matcher-asn-sets-tenant"></a>&#x2022; [`tenant`](#api-protection-rules-api-groups-rules-client-matcher-asn-matcher-asn-sets-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="api-protection-rules-api-groups-rules-client-matcher-asn-matcher-asn-sets-uid"></a>&#x2022; [`uid`](#api-protection-rules-api-groups-rules-client-matcher-asn-matcher-asn-sets-uid) - Optional String<br>UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid

#### API Protection Rules API Groups Rules Client Matcher Client Selector

A [`client_selector`](#api-protection-rules-api-groups-rules-client-matcher-client-selector) block (within [`api_protection_rules.api_groups_rules.client_matcher`](#api-protection-rules-api-groups-rules-client-matcher)) supports the following:

<a id="api-protection-rules-api-groups-rules-client-matcher-client-selector-expressions"></a>&#x2022; [`expressions`](#api-protection-rules-api-groups-rules-client-matcher-client-selector-expressions) - Optional List<br>Selector Expression. expressions contains the kubernetes style label expression for selections

#### API Protection Rules API Groups Rules Client Matcher IP Matcher

An [`ip_matcher`](#api-protection-rules-api-groups-rules-client-matcher-ip-matcher) block (within [`api_protection_rules.api_groups_rules.client_matcher`](#api-protection-rules-api-groups-rules-client-matcher)) supports the following:

<a id="api-protection-rules-api-groups-rules-client-matcher-ip-matcher-invert-matcher"></a>&#x2022; [`invert_matcher`](#api-protection-rules-api-groups-rules-client-matcher-ip-matcher-invert-matcher) - Optional Bool<br>Invert IP Matcher. Invert the match result

<a id="api-protection-rules-api-groups-rules-client-matcher-ip-matcher-prefix-sets"></a>&#x2022; [`prefix_sets`](#api-protection-rules-api-groups-rules-client-matcher-ip-matcher-prefix-sets) - Optional Block<br>IP Prefix Sets. A list of references to ip_prefix_set objects<br>See [Prefix Sets](#api-protection-rules-api-groups-rules-client-matcher-ip-matcher-prefix-sets) below.

#### API Protection Rules API Groups Rules Client Matcher IP Matcher Prefix Sets

A [`prefix_sets`](#api-protection-rules-api-groups-rules-client-matcher-ip-matcher-prefix-sets) block (within [`api_protection_rules.api_groups_rules.client_matcher.ip_matcher`](#api-protection-rules-api-groups-rules-client-matcher-ip-matcher)) supports the following:

<a id="api-protection-rules-api-groups-rules-client-matcher-ip-matcher-prefix-sets-kind"></a>&#x2022; [`kind`](#api-protection-rules-api-groups-rules-client-matcher-ip-matcher-prefix-sets-kind) - Optional String<br>Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route')

<a id="api-protection-rules-api-groups-rules-client-matcher-ip-matcher-prefix-sets-name"></a>&#x2022; [`name`](#api-protection-rules-api-groups-rules-client-matcher-ip-matcher-prefix-sets-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="api-protection-rules-api-groups-rules-client-matcher-ip-matcher-prefix-sets-namespace"></a>&#x2022; [`namespace`](#api-protection-rules-api-groups-rules-client-matcher-ip-matcher-prefix-sets-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="api-protection-rules-api-groups-rules-client-matcher-ip-matcher-prefix-sets-tenant"></a>&#x2022; [`tenant`](#api-protection-rules-api-groups-rules-client-matcher-ip-matcher-prefix-sets-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="api-protection-rules-api-groups-rules-client-matcher-ip-matcher-prefix-sets-uid"></a>&#x2022; [`uid`](#api-protection-rules-api-groups-rules-client-matcher-ip-matcher-prefix-sets-uid) - Optional String<br>UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid

#### API Protection Rules API Groups Rules Client Matcher IP Prefix List

An [`ip_prefix_list`](#api-protection-rules-api-groups-rules-client-matcher-ip-prefix-list) block (within [`api_protection_rules.api_groups_rules.client_matcher`](#api-protection-rules-api-groups-rules-client-matcher)) supports the following:

<a id="api-protection-rules-api-groups-rules-client-matcher-ip-prefix-list-invert-match"></a>&#x2022; [`invert_match`](#api-protection-rules-api-groups-rules-client-matcher-ip-prefix-list-invert-match) - Optional Bool<br>Invert Match Result. Invert the match result

<a id="api-protection-rules-api-groups-rules-client-matcher-ip-prefix-list-ip-prefixes"></a>&#x2022; [`ip_prefixes`](#api-protection-rules-api-groups-rules-client-matcher-ip-prefix-list-ip-prefixes) - Optional List<br>IPv4 Prefix List. List of IPv4 prefix strings

#### API Protection Rules API Groups Rules Client Matcher IP Threat Category List

An [`ip_threat_category_list`](#api-protection-rules-api-groups-rules-client-matcher-ip-threat-category-list) block (within [`api_protection_rules.api_groups_rules.client_matcher`](#api-protection-rules-api-groups-rules-client-matcher)) supports the following:

<a id="api-protection-rules-api-groups-rules-client-matcher-ip-threat-category-list-ip-threat-categories"></a>&#x2022; [`ip_threat_categories`](#api-protection-rules-api-groups-rules-client-matcher-ip-threat-category-list-ip-threat-categories) - Optional List  Defaults to `SPAM_SOURCES`<br>Possible values are `SPAM_SOURCES`, `WINDOWS_EXPLOITS`, `WEB_ATTACKS`, `BOTNETS`, `SCANNERS`, `REPUTATION`, `PHISHING`, `PROXY`, `MOBILE_THREATS`, `TOR_PROXY`, `DENIAL_OF_SERVICE`, `NETWORK`<br>List of IP Threat Categories to choose. The IP threat categories is obtained from the list and is used to auto-generate equivalent label selection expressions

#### API Protection Rules API Groups Rules Client Matcher TLS Fingerprint Matcher

A [`tls_fingerprint_matcher`](#api-protection-rules-api-groups-rules-client-matcher-tls-fingerprint-matcher) block (within [`api_protection_rules.api_groups_rules.client_matcher`](#api-protection-rules-api-groups-rules-client-matcher)) supports the following:

<a id="api-protection-rules-api-groups-rules-client-matcher-tls-fingerprint-matcher-classes"></a>&#x2022; [`classes`](#api-protection-rules-api-groups-rules-client-matcher-tls-fingerprint-matcher-classes) - Optional List  Defaults to `TLS_FINGERPRINT_NONE`<br>Possible values are `TLS_FINGERPRINT_NONE`, `ANY_MALICIOUS_FINGERPRINT`, `ADWARE`, `ADWIND`, `DRIDEX`, `GOOTKIT`, `GOZI`, `JBIFROST`, `QUAKBOT`, `RANSOMWARE`, `TROLDESH`, `TOFSEE`, `TORRENTLOCKER`, `TRICKBOT`<br>TLS fingerprint classes. A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against

<a id="api-protection-rules-api-groups-rules-client-matcher-tls-fingerprint-matcher-exact-values"></a>&#x2022; [`exact_values`](#api-protection-rules-api-groups-rules-client-matcher-tls-fingerprint-matcher-exact-values) - Optional List<br>Exact Values. A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against

<a id="api-protection-rules-api-groups-rules-client-matcher-tls-fingerprint-matcher-excluded-values"></a>&#x2022; [`excluded_values`](#api-protection-rules-api-groups-rules-client-matcher-tls-fingerprint-matcher-excluded-values) - Optional List<br>Excluded Values. A list of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can be used to skip known false positives when using one or more known TLS fingerprint classes in the enclosing matcher

#### API Protection Rules API Groups Rules Metadata

A [`metadata`](#api-protection-rules-api-groups-rules-metadata) block (within [`api_protection_rules.api_groups_rules`](#api-protection-rules-api-groups-rules)) supports the following:

<a id="api-protection-rules-api-groups-rules-metadata-description-spec"></a>&#x2022; [`description_spec`](#api-protection-rules-api-groups-rules-metadata-description-spec) - Optional String<br>Description. Human readable description

<a id="api-protection-rules-api-groups-rules-metadata-name"></a>&#x2022; [`name`](#api-protection-rules-api-groups-rules-metadata-name) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### API Protection Rules API Groups Rules Request Matcher

A [`request_matcher`](#api-protection-rules-api-groups-rules-request-matcher) block (within [`api_protection_rules.api_groups_rules`](#api-protection-rules-api-groups-rules)) supports the following:

<a id="api-protection-rules-api-groups-rules-request-matcher-cookie-matchers"></a>&#x2022; [`cookie_matchers`](#api-protection-rules-api-groups-rules-request-matcher-cookie-matchers) - Optional Block<br>Cookie Matchers. A list of predicates for all cookies that need to be matched. The criteria for matching each cookie is described in individual instances of CookieMatcherType. The actual cookie values are extracted from the request API as a list of strings for each cookie name. Note that all specified cookie matcher predicates must evaluate to true<br>See [Cookie Matchers](#api-protection-rules-api-groups-rules-request-matcher-cookie-matchers) below.

<a id="api-protection-rules-api-groups-rules-request-matcher-headers"></a>&#x2022; [`headers`](#api-protection-rules-api-groups-rules-request-matcher-headers) - Optional Block<br>HTTP Headers. A list of predicates for various HTTP headers that need to match. The criteria for matching each HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values are extracted from the request API as a list of strings for each HTTP header type. Note that all specified header predicates must evaluate to true<br>See [Headers](#api-protection-rules-api-groups-rules-request-matcher-headers) below.

<a id="api-protection-rules-api-groups-rules-request-matcher-jwt-claims"></a>&#x2022; [`jwt_claims`](#api-protection-rules-api-groups-rules-request-matcher-jwt-claims) - Optional Block<br>JWT Claims. A list of predicates for various JWT claims that need to match. The criteria for matching each JWT claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are extracted from the JWT payload as a list of strings. Note that all specified JWT claim predicates must evaluate to true. Note that this feature only works on LBs with JWT Validation feature enabled<br>See [JWT Claims](#api-protection-rules-api-groups-rules-request-matcher-jwt-claims) below.

<a id="api-protection-rules-api-groups-rules-request-matcher-query-params"></a>&#x2022; [`query_params`](#api-protection-rules-api-groups-rules-request-matcher-query-params) - Optional Block<br>HTTP Query Parameters. A list of predicates for all query parameters that need to be matched. The criteria for matching each query parameter are described in individual instances of QueryParameterMatcherType. The actual query parameter values are extracted from the request API as a list of strings for each query parameter name. Note that all specified query parameter predicates must evaluate to true<br>See [Query Params](#api-protection-rules-api-groups-rules-request-matcher-query-params) below.

#### API Protection Rules API Groups Rules Request Matcher Cookie Matchers

A [`cookie_matchers`](#api-protection-rules-api-groups-rules-request-matcher-cookie-matchers) block (within [`api_protection_rules.api_groups_rules.request_matcher`](#api-protection-rules-api-groups-rules-request-matcher)) supports the following:

<a id="api-protection-rules-api-groups-rules-request-matcher-cookie-matchers-check-not-present"></a>&#x2022; [`check_not_present`](#api-protection-rules-api-groups-rules-request-matcher-cookie-matchers-check-not-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-protection-rules-api-groups-rules-request-matcher-cookie-matchers-check-present"></a>&#x2022; [`check_present`](#api-protection-rules-api-groups-rules-request-matcher-cookie-matchers-check-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-protection-rules-api-groups-rules-request-matcher-cookie-matchers-invert-matcher"></a>&#x2022; [`invert_matcher`](#api-protection-rules-api-groups-rules-request-matcher-cookie-matchers-invert-matcher) - Optional Bool<br>Invert Matcher. Invert Match of the expression defined

<a id="api-protection-rules-api-groups-rules-request-matcher-cookie-matchers-item"></a>&#x2022; [`item`](#api-protection-rules-api-groups-rules-request-matcher-cookie-matchers-item) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#api-protection-rules-api-groups-rules-request-matcher-cookie-matchers-item) below.

<a id="api-protection-rules-api-groups-rules-request-matcher-cookie-matchers-name"></a>&#x2022; [`name`](#api-protection-rules-api-groups-rules-request-matcher-cookie-matchers-name) - Optional String<br>Cookie Name. A case-sensitive cookie name

#### API Protection Rules API Groups Rules Request Matcher Cookie Matchers Item

An [`item`](#api-protection-rules-api-groups-rules-request-matcher-cookie-matchers-item) block (within [`api_protection_rules.api_groups_rules.request_matcher.cookie_matchers`](#api-protection-rules-api-groups-rules-request-matcher-cookie-matchers)) supports the following:

<a id="api-protection-rules-api-groups-rules-request-matcher-cookie-matchers-item-exact-values"></a>&#x2022; [`exact_values`](#api-protection-rules-api-groups-rules-request-matcher-cookie-matchers-item-exact-values) - Optional List<br>Exact Values. A list of exact values to match the input against

<a id="api-protection-rules-api-groups-rules-request-matcher-cookie-matchers-item-regex-values"></a>&#x2022; [`regex_values`](#api-protection-rules-api-groups-rules-request-matcher-cookie-matchers-item-regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="api-protection-rules-api-groups-rules-request-matcher-cookie-matchers-item-transformers"></a>&#x2022; [`transformers`](#api-protection-rules-api-groups-rules-request-matcher-cookie-matchers-item-transformers) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### API Protection Rules API Groups Rules Request Matcher Headers

A [`headers`](#api-protection-rules-api-groups-rules-request-matcher-headers) block (within [`api_protection_rules.api_groups_rules.request_matcher`](#api-protection-rules-api-groups-rules-request-matcher)) supports the following:

<a id="api-protection-rules-api-groups-rules-request-matcher-headers-check-not-present"></a>&#x2022; [`check_not_present`](#api-protection-rules-api-groups-rules-request-matcher-headers-check-not-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-protection-rules-api-groups-rules-request-matcher-headers-check-present"></a>&#x2022; [`check_present`](#api-protection-rules-api-groups-rules-request-matcher-headers-check-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-protection-rules-api-groups-rules-request-matcher-headers-invert-matcher"></a>&#x2022; [`invert_matcher`](#api-protection-rules-api-groups-rules-request-matcher-headers-invert-matcher) - Optional Bool<br>Invert Header Matcher. Invert the match result

<a id="api-protection-rules-api-groups-rules-request-matcher-headers-item"></a>&#x2022; [`item`](#api-protection-rules-api-groups-rules-request-matcher-headers-item) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#api-protection-rules-api-groups-rules-request-matcher-headers-item) below.

<a id="api-protection-rules-api-groups-rules-request-matcher-headers-name"></a>&#x2022; [`name`](#api-protection-rules-api-groups-rules-request-matcher-headers-name) - Optional String<br>Header Name. A case-insensitive HTTP header name

#### API Protection Rules API Groups Rules Request Matcher Headers Item

An [`item`](#api-protection-rules-api-groups-rules-request-matcher-headers-item) block (within [`api_protection_rules.api_groups_rules.request_matcher.headers`](#api-protection-rules-api-groups-rules-request-matcher-headers)) supports the following:

<a id="api-protection-rules-api-groups-rules-request-matcher-headers-item-exact-values"></a>&#x2022; [`exact_values`](#api-protection-rules-api-groups-rules-request-matcher-headers-item-exact-values) - Optional List<br>Exact Values. A list of exact values to match the input against

<a id="api-protection-rules-api-groups-rules-request-matcher-headers-item-regex-values"></a>&#x2022; [`regex_values`](#api-protection-rules-api-groups-rules-request-matcher-headers-item-regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="api-protection-rules-api-groups-rules-request-matcher-headers-item-transformers"></a>&#x2022; [`transformers`](#api-protection-rules-api-groups-rules-request-matcher-headers-item-transformers) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### API Protection Rules API Groups Rules Request Matcher JWT Claims

A [`jwt_claims`](#api-protection-rules-api-groups-rules-request-matcher-jwt-claims) block (within [`api_protection_rules.api_groups_rules.request_matcher`](#api-protection-rules-api-groups-rules-request-matcher)) supports the following:

<a id="api-protection-rules-api-groups-rules-request-matcher-jwt-claims-check-not-present"></a>&#x2022; [`check_not_present`](#api-protection-rules-api-groups-rules-request-matcher-jwt-claims-check-not-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-protection-rules-api-groups-rules-request-matcher-jwt-claims-check-present"></a>&#x2022; [`check_present`](#api-protection-rules-api-groups-rules-request-matcher-jwt-claims-check-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-protection-rules-api-groups-rules-request-matcher-jwt-claims-invert-matcher"></a>&#x2022; [`invert_matcher`](#api-protection-rules-api-groups-rules-request-matcher-jwt-claims-invert-matcher) - Optional Bool<br>Invert Matcher. Invert the match result

<a id="api-protection-rules-api-groups-rules-request-matcher-jwt-claims-item"></a>&#x2022; [`item`](#api-protection-rules-api-groups-rules-request-matcher-jwt-claims-item) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#api-protection-rules-api-groups-rules-request-matcher-jwt-claims-item) below.

<a id="api-protection-rules-api-groups-rules-request-matcher-jwt-claims-name"></a>&#x2022; [`name`](#api-protection-rules-api-groups-rules-request-matcher-jwt-claims-name) - Optional String<br>JWT Claim Name. JWT claim name

#### API Protection Rules API Groups Rules Request Matcher JWT Claims Item

An [`item`](#api-protection-rules-api-groups-rules-request-matcher-jwt-claims-item) block (within [`api_protection_rules.api_groups_rules.request_matcher.jwt_claims`](#api-protection-rules-api-groups-rules-request-matcher-jwt-claims)) supports the following:

<a id="api-protection-rules-api-groups-rules-request-matcher-jwt-claims-item-exact-values"></a>&#x2022; [`exact_values`](#api-protection-rules-api-groups-rules-request-matcher-jwt-claims-item-exact-values) - Optional List<br>Exact Values. A list of exact values to match the input against

<a id="api-protection-rules-api-groups-rules-request-matcher-jwt-claims-item-regex-values"></a>&#x2022; [`regex_values`](#api-protection-rules-api-groups-rules-request-matcher-jwt-claims-item-regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="api-protection-rules-api-groups-rules-request-matcher-jwt-claims-item-transformers"></a>&#x2022; [`transformers`](#api-protection-rules-api-groups-rules-request-matcher-jwt-claims-item-transformers) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### API Protection Rules API Groups Rules Request Matcher Query Params

A [`query_params`](#api-protection-rules-api-groups-rules-request-matcher-query-params) block (within [`api_protection_rules.api_groups_rules.request_matcher`](#api-protection-rules-api-groups-rules-request-matcher)) supports the following:

<a id="api-protection-rules-api-groups-rules-request-matcher-query-params-check-not-present"></a>&#x2022; [`check_not_present`](#api-protection-rules-api-groups-rules-request-matcher-query-params-check-not-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-protection-rules-api-groups-rules-request-matcher-query-params-check-present"></a>&#x2022; [`check_present`](#api-protection-rules-api-groups-rules-request-matcher-query-params-check-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-protection-rules-api-groups-rules-request-matcher-query-params-invert-matcher"></a>&#x2022; [`invert_matcher`](#api-protection-rules-api-groups-rules-request-matcher-query-params-invert-matcher) - Optional Bool<br>Invert Query Parameter Matcher. Invert the match result

<a id="api-protection-rules-api-groups-rules-request-matcher-query-params-item"></a>&#x2022; [`item`](#api-protection-rules-api-groups-rules-request-matcher-query-params-item) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#api-protection-rules-api-groups-rules-request-matcher-query-params-item) below.

<a id="api-protection-rules-api-groups-rules-request-matcher-query-params-key"></a>&#x2022; [`key`](#api-protection-rules-api-groups-rules-request-matcher-query-params-key) - Optional String<br>Query Parameter Name. A case-sensitive HTTP query parameter name

#### API Protection Rules API Groups Rules Request Matcher Query Params Item

An [`item`](#api-protection-rules-api-groups-rules-request-matcher-query-params-item) block (within [`api_protection_rules.api_groups_rules.request_matcher.query_params`](#api-protection-rules-api-groups-rules-request-matcher-query-params)) supports the following:

<a id="api-protection-rules-api-groups-rules-request-matcher-query-params-item-exact-values"></a>&#x2022; [`exact_values`](#api-protection-rules-api-groups-rules-request-matcher-query-params-item-exact-values) - Optional List<br>Exact Values. A list of exact values to match the input against

<a id="api-protection-rules-api-groups-rules-request-matcher-query-params-item-regex-values"></a>&#x2022; [`regex_values`](#api-protection-rules-api-groups-rules-request-matcher-query-params-item-regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="api-protection-rules-api-groups-rules-request-matcher-query-params-item-transformers"></a>&#x2022; [`transformers`](#api-protection-rules-api-groups-rules-request-matcher-query-params-item-transformers) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### API Rate Limit

An [`api_rate_limit`](#api-rate-limit) block supports the following:

<a id="api-rate-limit-api-endpoint-rules"></a>&#x2022; [`api_endpoint_rules`](#api-rate-limit-api-endpoint-rules) - Optional Block<br>API Endpoints. Sets of rules for a specific endpoints. Order is matter as it uses first match policy. For creating rule that contain a whole domain or group of endpoints, please use the server URL rules above<br>See [API Endpoint Rules](#api-rate-limit-api-endpoint-rules) below.

<a id="api-rate-limit-bypass-rate-limiting-rules"></a>&#x2022; [`bypass_rate_limiting_rules`](#api-rate-limit-bypass-rate-limiting-rules) - Optional Block<br>Bypass Rate Limiting. This category defines rules per URL or API group. If request matches any of these rules, skip Rate Limiting<br>See [Bypass Rate Limiting Rules](#api-rate-limit-bypass-rate-limiting-rules) below.

<a id="api-rate-limit-custom-ip-allowed-list"></a>&#x2022; [`custom_ip_allowed_list`](#api-rate-limit-custom-ip-allowed-list) - Optional Block<br>Custom IP Allowed List. IP Allowed list using existing ip_prefix_set objects<br>See [Custom IP Allowed List](#api-rate-limit-custom-ip-allowed-list) below.

<a id="api-rate-limit-ip-allowed-list"></a>&#x2022; [`ip_allowed_list`](#api-rate-limit-ip-allowed-list) - Optional Block<br>IPv4 Prefix List. x-example: '192.168.20.0/24' List of IPv4 prefixes that represent an endpoint<br>See [IP Allowed List](#api-rate-limit-ip-allowed-list) below.

<a id="api-rate-limit-no-ip-allowed-list"></a>&#x2022; [`no_ip_allowed_list`](#api-rate-limit-no-ip-allowed-list) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-rate-limit-server-url-rules"></a>&#x2022; [`server_url_rules`](#api-rate-limit-server-url-rules) - Optional Block<br>Server URLs. Set of rules for entire domain or base path that contain multiple endpoints. Order is matter as it uses first match policy. For matching also specific endpoints you can use the API endpoint rules set bellow<br>See [Server URL Rules](#api-rate-limit-server-url-rules) below.

#### API Rate Limit API Endpoint Rules

An [`api_endpoint_rules`](#api-rate-limit-api-endpoint-rules) block (within [`api_rate_limit`](#api-rate-limit)) supports the following:

<a id="api-rate-limit-api-endpoint-rules-any-domain"></a>&#x2022; [`any_domain`](#api-rate-limit-api-endpoint-rules-any-domain) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-rate-limit-api-endpoint-rules-api-endpoint-method"></a>&#x2022; [`api_endpoint_method`](#api-rate-limit-api-endpoint-rules-api-endpoint-method) - Optional Block<br>HTTP Method Matcher. A HTTP method matcher specifies a list of methods to match an input HTTP method. The match is considered successful if the input method is a member of the list. The result of the match based on the method list is inverted if invert_matcher is true<br>See [API Endpoint Method](#api-rate-limit-api-endpoint-rules-api-endpoint-method) below.

<a id="api-rate-limit-api-endpoint-rules-api-endpoint-path"></a>&#x2022; [`api_endpoint_path`](#api-rate-limit-api-endpoint-rules-api-endpoint-path) - Optional String<br>API Endpoint. The endpoint (path) of the request

<a id="api-rate-limit-api-endpoint-rules-client-matcher"></a>&#x2022; [`client_matcher`](#api-rate-limit-api-endpoint-rules-client-matcher) - Optional Block<br>Client Matcher. Client conditions for matching a rule<br>See [Client Matcher](#api-rate-limit-api-endpoint-rules-client-matcher) below.

<a id="api-rate-limit-api-endpoint-rules-inline-rate-limiter"></a>&#x2022; [`inline_rate_limiter`](#api-rate-limit-api-endpoint-rules-inline-rate-limiter) - Optional Block<br>InlineRateLimiter<br>See [Inline Rate Limiter](#api-rate-limit-api-endpoint-rules-inline-rate-limiter) below.

<a id="api-rate-limit-api-endpoint-rules-ref-rate-limiter"></a>&#x2022; [`ref_rate_limiter`](#api-rate-limit-api-endpoint-rules-ref-rate-limiter) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Ref Rate Limiter](#api-rate-limit-api-endpoint-rules-ref-rate-limiter) below.

<a id="api-rate-limit-api-endpoint-rules-request-matcher"></a>&#x2022; [`request_matcher`](#api-rate-limit-api-endpoint-rules-request-matcher) - Optional Block<br>Request Matcher. Request conditions for matching a rule<br>See [Request Matcher](#api-rate-limit-api-endpoint-rules-request-matcher) below.

<a id="api-rate-limit-api-endpoint-rules-specific-domain"></a>&#x2022; [`specific_domain`](#api-rate-limit-api-endpoint-rules-specific-domain) - Optional String<br>Specific Domain. The rule will apply for a specific domain

#### API Rate Limit API Endpoint Rules API Endpoint Method

An [`api_endpoint_method`](#api-rate-limit-api-endpoint-rules-api-endpoint-method) block (within [`api_rate_limit.api_endpoint_rules`](#api-rate-limit-api-endpoint-rules)) supports the following:

<a id="api-rate-limit-api-endpoint-rules-api-endpoint-method-invert-matcher"></a>&#x2022; [`invert_matcher`](#api-rate-limit-api-endpoint-rules-api-endpoint-method-invert-matcher) - Optional Bool<br>Invert Method Matcher. Invert the match result

<a id="api-rate-limit-api-endpoint-rules-api-endpoint-method-methods"></a>&#x2022; [`methods`](#api-rate-limit-api-endpoint-rules-api-endpoint-method-methods) - Optional List  Defaults to `ANY`<br>Possible values are `ANY`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`, `COPY`<br>Method List. List of methods values to match against

#### API Rate Limit API Endpoint Rules Client Matcher

A [`client_matcher`](#api-rate-limit-api-endpoint-rules-client-matcher) block (within [`api_rate_limit.api_endpoint_rules`](#api-rate-limit-api-endpoint-rules)) supports the following:

<a id="api-rate-limit-api-endpoint-rules-client-matcher-any-client"></a>&#x2022; [`any_client`](#api-rate-limit-api-endpoint-rules-client-matcher-any-client) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-rate-limit-api-endpoint-rules-client-matcher-any-ip"></a>&#x2022; [`any_ip`](#api-rate-limit-api-endpoint-rules-client-matcher-any-ip) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-rate-limit-api-endpoint-rules-client-matcher-asn-list"></a>&#x2022; [`asn_list`](#api-rate-limit-api-endpoint-rules-client-matcher-asn-list) - Optional Block<br>ASN Match List. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer<br>See [Asn List](#api-rate-limit-api-endpoint-rules-client-matcher-asn-list) below.

<a id="api-rate-limit-api-endpoint-rules-client-matcher-asn-matcher"></a>&#x2022; [`asn_matcher`](#api-rate-limit-api-endpoint-rules-client-matcher-asn-matcher) - Optional Block<br>ASN Matcher. Match any AS number contained in the list of bgp_asn_sets<br>See [Asn Matcher](#api-rate-limit-api-endpoint-rules-client-matcher-asn-matcher) below.

<a id="api-rate-limit-api-endpoint-rules-client-matcher-client-selector"></a>&#x2022; [`client_selector`](#api-rate-limit-api-endpoint-rules-client-matcher-client-selector) - Optional Block<br>Label Selector. This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expresssions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects. Label selector is immutable. expressions is a list of strings of label selection expression. Each string has ',' separated values which are 'AND' and all strings are logically 'OR'. BNF for expression string `<selector-syntax>` ::= `<requirement>` | `<requirement>` ',' `<selector-syntax>` `<requirement>` ::= [!] KEY [ `<set-based-restriction>` | `<exact-match-restriction>` ] `<set-based-restriction>` ::= '' | `<inclusion-exclusion>` `<value-set>` `<inclusion-exclusion>` ::= `<inclusion>` | `<exclusion>` `<exclusion>` ::= 'notin' `<inclusion>` ::= 'in' `<value-set>` ::= '(' `<values>` ')' `<values>` ::= VALUE | VALUE ',' `<values>` `<exact-match-restriction>` ::= ['='|'=='|'!='] VALUE<br>See [Client Selector](#api-rate-limit-api-endpoint-rules-client-matcher-client-selector) below.

<a id="api-rate-limit-api-endpoint-rules-client-matcher-ip-matcher"></a>&#x2022; [`ip_matcher`](#api-rate-limit-api-endpoint-rules-client-matcher-ip-matcher) - Optional Block<br>IP Prefix Matcher. Match any IP prefix contained in the list of ip_prefix_sets. The result of the match is inverted if invert_matcher is true<br>See [IP Matcher](#api-rate-limit-api-endpoint-rules-client-matcher-ip-matcher) below.

<a id="api-rate-limit-api-endpoint-rules-client-matcher-ip-prefix-list"></a>&#x2022; [`ip_prefix_list`](#api-rate-limit-api-endpoint-rules-client-matcher-ip-prefix-list) - Optional Block<br>IP Prefix Match List. List of IP Prefix strings to match against<br>See [IP Prefix List](#api-rate-limit-api-endpoint-rules-client-matcher-ip-prefix-list) below.

<a id="api-rate-limit-api-endpoint-rules-client-matcher-ip-threat-category-list"></a>&#x2022; [`ip_threat_category_list`](#api-rate-limit-api-endpoint-rules-client-matcher-ip-threat-category-list) - Optional Block<br>IP Threat Category List Type. List of IP threat categories<br>See [IP Threat Category List](#api-rate-limit-api-endpoint-rules-client-matcher-ip-threat-category-list) below.

<a id="api-rate-limit-api-endpoint-rules-client-matcher-tls-fingerprint-matcher"></a>&#x2022; [`tls_fingerprint_matcher`](#api-rate-limit-api-endpoint-rules-client-matcher-tls-fingerprint-matcher) - Optional Block<br>TLS Fingerprint Matcher. A TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of supported positve match criteria includes a list of known classes of TLS fingerprints and a list of exact values. The match is considered successful if either of these positive criteria are satisfied and the input fingerprint is not one of the excluded values<br>See [TLS Fingerprint Matcher](#api-rate-limit-api-endpoint-rules-client-matcher-tls-fingerprint-matcher) below.

#### API Rate Limit API Endpoint Rules Client Matcher Asn List

An [`asn_list`](#api-rate-limit-api-endpoint-rules-client-matcher-asn-list) block (within [`api_rate_limit.api_endpoint_rules.client_matcher`](#api-rate-limit-api-endpoint-rules-client-matcher)) supports the following:

<a id="api-rate-limit-api-endpoint-rules-client-matcher-asn-list-as-numbers"></a>&#x2022; [`as_numbers`](#api-rate-limit-api-endpoint-rules-client-matcher-asn-list-as-numbers) - Optional List<br>AS Numbers. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer

#### API Rate Limit API Endpoint Rules Client Matcher Asn Matcher

An [`asn_matcher`](#api-rate-limit-api-endpoint-rules-client-matcher-asn-matcher) block (within [`api_rate_limit.api_endpoint_rules.client_matcher`](#api-rate-limit-api-endpoint-rules-client-matcher)) supports the following:

<a id="api-rate-limit-api-endpoint-rules-client-matcher-asn-matcher-asn-sets"></a>&#x2022; [`asn_sets`](#api-rate-limit-api-endpoint-rules-client-matcher-asn-matcher-asn-sets) - Optional Block<br>BGP ASN Sets. A list of references to bgp_asn_set objects<br>See [Asn Sets](#api-rate-limit-api-endpoint-rules-client-matcher-asn-matcher-asn-sets) below.

#### API Rate Limit API Endpoint Rules Client Matcher Asn Matcher Asn Sets

An [`asn_sets`](#api-rate-limit-api-endpoint-rules-client-matcher-asn-matcher-asn-sets) block (within [`api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher`](#api-rate-limit-api-endpoint-rules-client-matcher-asn-matcher)) supports the following:

<a id="api-rate-limit-api-endpoint-rules-client-matcher-asn-matcher-asn-sets-kind"></a>&#x2022; [`kind`](#api-rate-limit-api-endpoint-rules-client-matcher-asn-matcher-asn-sets-kind) - Optional String<br>Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route')

<a id="api-rate-limit-api-endpoint-rules-client-matcher-asn-matcher-asn-sets-name"></a>&#x2022; [`name`](#api-rate-limit-api-endpoint-rules-client-matcher-asn-matcher-asn-sets-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="api-rate-limit-api-endpoint-rules-client-matcher-asn-matcher-asn-sets-namespace"></a>&#x2022; [`namespace`](#api-rate-limit-api-endpoint-rules-client-matcher-asn-matcher-asn-sets-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="api-rate-limit-api-endpoint-rules-client-matcher-asn-matcher-asn-sets-tenant"></a>&#x2022; [`tenant`](#api-rate-limit-api-endpoint-rules-client-matcher-asn-matcher-asn-sets-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="api-rate-limit-api-endpoint-rules-client-matcher-asn-matcher-asn-sets-uid"></a>&#x2022; [`uid`](#api-rate-limit-api-endpoint-rules-client-matcher-asn-matcher-asn-sets-uid) - Optional String<br>UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid

#### API Rate Limit API Endpoint Rules Client Matcher Client Selector

A [`client_selector`](#api-rate-limit-api-endpoint-rules-client-matcher-client-selector) block (within [`api_rate_limit.api_endpoint_rules.client_matcher`](#api-rate-limit-api-endpoint-rules-client-matcher)) supports the following:

<a id="api-rate-limit-api-endpoint-rules-client-matcher-client-selector-expressions"></a>&#x2022; [`expressions`](#api-rate-limit-api-endpoint-rules-client-matcher-client-selector-expressions) - Optional List<br>Selector Expression. expressions contains the kubernetes style label expression for selections

#### API Rate Limit API Endpoint Rules Client Matcher IP Matcher

An [`ip_matcher`](#api-rate-limit-api-endpoint-rules-client-matcher-ip-matcher) block (within [`api_rate_limit.api_endpoint_rules.client_matcher`](#api-rate-limit-api-endpoint-rules-client-matcher)) supports the following:

<a id="api-rate-limit-api-endpoint-rules-client-matcher-ip-matcher-invert-matcher"></a>&#x2022; [`invert_matcher`](#api-rate-limit-api-endpoint-rules-client-matcher-ip-matcher-invert-matcher) - Optional Bool<br>Invert IP Matcher. Invert the match result

<a id="api-rate-limit-api-endpoint-rules-client-matcher-ip-matcher-prefix-sets"></a>&#x2022; [`prefix_sets`](#api-rate-limit-api-endpoint-rules-client-matcher-ip-matcher-prefix-sets) - Optional Block<br>IP Prefix Sets. A list of references to ip_prefix_set objects<br>See [Prefix Sets](#api-rate-limit-api-endpoint-rules-client-matcher-ip-matcher-prefix-sets) below.

#### API Rate Limit API Endpoint Rules Client Matcher IP Matcher Prefix Sets

A [`prefix_sets`](#api-rate-limit-api-endpoint-rules-client-matcher-ip-matcher-prefix-sets) block (within [`api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher`](#api-rate-limit-api-endpoint-rules-client-matcher-ip-matcher)) supports the following:

<a id="api-rate-limit-api-endpoint-rules-client-matcher-ip-matcher-prefix-sets-kind"></a>&#x2022; [`kind`](#api-rate-limit-api-endpoint-rules-client-matcher-ip-matcher-prefix-sets-kind) - Optional String<br>Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route')

<a id="api-rate-limit-api-endpoint-rules-client-matcher-ip-matcher-prefix-sets-name"></a>&#x2022; [`name`](#api-rate-limit-api-endpoint-rules-client-matcher-ip-matcher-prefix-sets-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="api-rate-limit-api-endpoint-rules-client-matcher-ip-matcher-prefix-sets-namespace"></a>&#x2022; [`namespace`](#api-rate-limit-api-endpoint-rules-client-matcher-ip-matcher-prefix-sets-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="api-rate-limit-api-endpoint-rules-client-matcher-ip-matcher-prefix-sets-tenant"></a>&#x2022; [`tenant`](#api-rate-limit-api-endpoint-rules-client-matcher-ip-matcher-prefix-sets-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="api-rate-limit-api-endpoint-rules-client-matcher-ip-matcher-prefix-sets-uid"></a>&#x2022; [`uid`](#api-rate-limit-api-endpoint-rules-client-matcher-ip-matcher-prefix-sets-uid) - Optional String<br>UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid

#### API Rate Limit API Endpoint Rules Client Matcher IP Prefix List

An [`ip_prefix_list`](#api-rate-limit-api-endpoint-rules-client-matcher-ip-prefix-list) block (within [`api_rate_limit.api_endpoint_rules.client_matcher`](#api-rate-limit-api-endpoint-rules-client-matcher)) supports the following:

<a id="api-rate-limit-api-endpoint-rules-client-matcher-ip-prefix-list-invert-match"></a>&#x2022; [`invert_match`](#api-rate-limit-api-endpoint-rules-client-matcher-ip-prefix-list-invert-match) - Optional Bool<br>Invert Match Result. Invert the match result

<a id="api-rate-limit-api-endpoint-rules-client-matcher-ip-prefix-list-ip-prefixes"></a>&#x2022; [`ip_prefixes`](#api-rate-limit-api-endpoint-rules-client-matcher-ip-prefix-list-ip-prefixes) - Optional List<br>IPv4 Prefix List. List of IPv4 prefix strings

#### API Rate Limit API Endpoint Rules Client Matcher IP Threat Category List

An [`ip_threat_category_list`](#api-rate-limit-api-endpoint-rules-client-matcher-ip-threat-category-list) block (within [`api_rate_limit.api_endpoint_rules.client_matcher`](#api-rate-limit-api-endpoint-rules-client-matcher)) supports the following:

<a id="api-rate-limit-api-endpoint-rules-client-matcher-ip-threat-category-list-ip-threat-categories"></a>&#x2022; [`ip_threat_categories`](#api-rate-limit-api-endpoint-rules-client-matcher-ip-threat-category-list-ip-threat-categories) - Optional List  Defaults to `SPAM_SOURCES`<br>Possible values are `SPAM_SOURCES`, `WINDOWS_EXPLOITS`, `WEB_ATTACKS`, `BOTNETS`, `SCANNERS`, `REPUTATION`, `PHISHING`, `PROXY`, `MOBILE_THREATS`, `TOR_PROXY`, `DENIAL_OF_SERVICE`, `NETWORK`<br>List of IP Threat Categories to choose. The IP threat categories is obtained from the list and is used to auto-generate equivalent label selection expressions

#### API Rate Limit API Endpoint Rules Client Matcher TLS Fingerprint Matcher

A [`tls_fingerprint_matcher`](#api-rate-limit-api-endpoint-rules-client-matcher-tls-fingerprint-matcher) block (within [`api_rate_limit.api_endpoint_rules.client_matcher`](#api-rate-limit-api-endpoint-rules-client-matcher)) supports the following:

<a id="api-rate-limit-api-endpoint-rules-client-matcher-tls-fingerprint-matcher-classes"></a>&#x2022; [`classes`](#api-rate-limit-api-endpoint-rules-client-matcher-tls-fingerprint-matcher-classes) - Optional List  Defaults to `TLS_FINGERPRINT_NONE`<br>Possible values are `TLS_FINGERPRINT_NONE`, `ANY_MALICIOUS_FINGERPRINT`, `ADWARE`, `ADWIND`, `DRIDEX`, `GOOTKIT`, `GOZI`, `JBIFROST`, `QUAKBOT`, `RANSOMWARE`, `TROLDESH`, `TOFSEE`, `TORRENTLOCKER`, `TRICKBOT`<br>TLS fingerprint classes. A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against

<a id="api-rate-limit-api-endpoint-rules-client-matcher-tls-fingerprint-matcher-exact-values"></a>&#x2022; [`exact_values`](#api-rate-limit-api-endpoint-rules-client-matcher-tls-fingerprint-matcher-exact-values) - Optional List<br>Exact Values. A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against

<a id="api-rate-limit-api-endpoint-rules-client-matcher-tls-fingerprint-matcher-excluded-values"></a>&#x2022; [`excluded_values`](#api-rate-limit-api-endpoint-rules-client-matcher-tls-fingerprint-matcher-excluded-values) - Optional List<br>Excluded Values. A list of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can be used to skip known false positives when using one or more known TLS fingerprint classes in the enclosing matcher

#### API Rate Limit API Endpoint Rules Inline Rate Limiter

An [`inline_rate_limiter`](#api-rate-limit-api-endpoint-rules-inline-rate-limiter) block (within [`api_rate_limit.api_endpoint_rules`](#api-rate-limit-api-endpoint-rules)) supports the following:

<a id="api-rate-limit-api-endpoint-rules-inline-rate-limiter-ref-user-id"></a>&#x2022; [`ref_user_id`](#api-rate-limit-api-endpoint-rules-inline-rate-limiter-ref-user-id) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Ref User Id](#api-rate-limit-api-endpoint-rules-inline-rate-limiter-ref-user-id) below.

<a id="api-rate-limit-api-endpoint-rules-inline-rate-limiter-threshold"></a>&#x2022; [`threshold`](#api-rate-limit-api-endpoint-rules-inline-rate-limiter-threshold) - Optional Number<br>Threshold. The total number of allowed requests for 1 unit (e.g. SECOND/MINUTE/HOUR etc.) of the specified period

<a id="api-rate-limit-api-endpoint-rules-inline-rate-limiter-unit"></a>&#x2022; [`unit`](#api-rate-limit-api-endpoint-rules-inline-rate-limiter-unit) - Optional String  Defaults to `SECOND`<br>Possible values are `SECOND`, `MINUTE`, `HOUR`<br>Rate Limit Period Unit. Unit for the period per which the rate limit is applied. - SECOND: Second Rate limit period unit is seconds - MINUTE: Minute Rate limit period unit is minutes - HOUR: Hour Rate limit period unit is hours - DAY: Day Rate limit period unit is days

<a id="api-rate-limit-api-endpoint-rules-inline-rate-limiter-use-http-lb-user-id"></a>&#x2022; [`use_http_lb_user_id`](#api-rate-limit-api-endpoint-rules-inline-rate-limiter-use-http-lb-user-id) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### API Rate Limit API Endpoint Rules Inline Rate Limiter Ref User Id

A [`ref_user_id`](#api-rate-limit-api-endpoint-rules-inline-rate-limiter-ref-user-id) block (within [`api_rate_limit.api_endpoint_rules.inline_rate_limiter`](#api-rate-limit-api-endpoint-rules-inline-rate-limiter)) supports the following:

<a id="api-rate-limit-api-endpoint-rules-inline-rate-limiter-ref-user-id-name"></a>&#x2022; [`name`](#api-rate-limit-api-endpoint-rules-inline-rate-limiter-ref-user-id-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="api-rate-limit-api-endpoint-rules-inline-rate-limiter-ref-user-id-namespace"></a>&#x2022; [`namespace`](#api-rate-limit-api-endpoint-rules-inline-rate-limiter-ref-user-id-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="api-rate-limit-api-endpoint-rules-inline-rate-limiter-ref-user-id-tenant"></a>&#x2022; [`tenant`](#api-rate-limit-api-endpoint-rules-inline-rate-limiter-ref-user-id-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### API Rate Limit API Endpoint Rules Ref Rate Limiter

A [`ref_rate_limiter`](#api-rate-limit-api-endpoint-rules-ref-rate-limiter) block (within [`api_rate_limit.api_endpoint_rules`](#api-rate-limit-api-endpoint-rules)) supports the following:

<a id="api-rate-limit-api-endpoint-rules-ref-rate-limiter-name"></a>&#x2022; [`name`](#api-rate-limit-api-endpoint-rules-ref-rate-limiter-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="api-rate-limit-api-endpoint-rules-ref-rate-limiter-namespace"></a>&#x2022; [`namespace`](#api-rate-limit-api-endpoint-rules-ref-rate-limiter-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="api-rate-limit-api-endpoint-rules-ref-rate-limiter-tenant"></a>&#x2022; [`tenant`](#api-rate-limit-api-endpoint-rules-ref-rate-limiter-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### API Rate Limit API Endpoint Rules Request Matcher

A [`request_matcher`](#api-rate-limit-api-endpoint-rules-request-matcher) block (within [`api_rate_limit.api_endpoint_rules`](#api-rate-limit-api-endpoint-rules)) supports the following:

<a id="api-rate-limit-api-endpoint-rules-request-matcher-cookie-matchers"></a>&#x2022; [`cookie_matchers`](#api-rate-limit-api-endpoint-rules-request-matcher-cookie-matchers) - Optional Block<br>Cookie Matchers. A list of predicates for all cookies that need to be matched. The criteria for matching each cookie is described in individual instances of CookieMatcherType. The actual cookie values are extracted from the request API as a list of strings for each cookie name. Note that all specified cookie matcher predicates must evaluate to true<br>See [Cookie Matchers](#api-rate-limit-api-endpoint-rules-request-matcher-cookie-matchers) below.

<a id="api-rate-limit-api-endpoint-rules-request-matcher-headers"></a>&#x2022; [`headers`](#api-rate-limit-api-endpoint-rules-request-matcher-headers) - Optional Block<br>HTTP Headers. A list of predicates for various HTTP headers that need to match. The criteria for matching each HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values are extracted from the request API as a list of strings for each HTTP header type. Note that all specified header predicates must evaluate to true<br>See [Headers](#api-rate-limit-api-endpoint-rules-request-matcher-headers) below.

<a id="api-rate-limit-api-endpoint-rules-request-matcher-jwt-claims"></a>&#x2022; [`jwt_claims`](#api-rate-limit-api-endpoint-rules-request-matcher-jwt-claims) - Optional Block<br>JWT Claims. A list of predicates for various JWT claims that need to match. The criteria for matching each JWT claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are extracted from the JWT payload as a list of strings. Note that all specified JWT claim predicates must evaluate to true. Note that this feature only works on LBs with JWT Validation feature enabled<br>See [JWT Claims](#api-rate-limit-api-endpoint-rules-request-matcher-jwt-claims) below.

<a id="api-rate-limit-api-endpoint-rules-request-matcher-query-params"></a>&#x2022; [`query_params`](#api-rate-limit-api-endpoint-rules-request-matcher-query-params) - Optional Block<br>HTTP Query Parameters. A list of predicates for all query parameters that need to be matched. The criteria for matching each query parameter are described in individual instances of QueryParameterMatcherType. The actual query parameter values are extracted from the request API as a list of strings for each query parameter name. Note that all specified query parameter predicates must evaluate to true<br>See [Query Params](#api-rate-limit-api-endpoint-rules-request-matcher-query-params) below.

#### API Rate Limit API Endpoint Rules Request Matcher Cookie Matchers

A [`cookie_matchers`](#api-rate-limit-api-endpoint-rules-request-matcher-cookie-matchers) block (within [`api_rate_limit.api_endpoint_rules.request_matcher`](#api-rate-limit-api-endpoint-rules-request-matcher)) supports the following:

<a id="api-rate-limit-api-endpoint-rules-request-matcher-cookie-matchers-check-not-present"></a>&#x2022; [`check_not_present`](#api-rate-limit-api-endpoint-rules-request-matcher-cookie-matchers-check-not-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-rate-limit-api-endpoint-rules-request-matcher-cookie-matchers-check-present"></a>&#x2022; [`check_present`](#api-rate-limit-api-endpoint-rules-request-matcher-cookie-matchers-check-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-rate-limit-api-endpoint-rules-request-matcher-cookie-matchers-invert-matcher"></a>&#x2022; [`invert_matcher`](#api-rate-limit-api-endpoint-rules-request-matcher-cookie-matchers-invert-matcher) - Optional Bool<br>Invert Matcher. Invert Match of the expression defined

<a id="api-rate-limit-api-endpoint-rules-request-matcher-cookie-matchers-item"></a>&#x2022; [`item`](#api-rate-limit-api-endpoint-rules-request-matcher-cookie-matchers-item) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#api-rate-limit-api-endpoint-rules-request-matcher-cookie-matchers-item) below.

<a id="api-rate-limit-api-endpoint-rules-request-matcher-cookie-matchers-name"></a>&#x2022; [`name`](#api-rate-limit-api-endpoint-rules-request-matcher-cookie-matchers-name) - Optional String<br>Cookie Name. A case-sensitive cookie name

#### API Rate Limit API Endpoint Rules Request Matcher Cookie Matchers Item

An [`item`](#api-rate-limit-api-endpoint-rules-request-matcher-cookie-matchers-item) block (within [`api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers`](#api-rate-limit-api-endpoint-rules-request-matcher-cookie-matchers)) supports the following:

<a id="api-rate-limit-api-endpoint-rules-request-matcher-cookie-matchers-item-exact-values"></a>&#x2022; [`exact_values`](#api-rate-limit-api-endpoint-rules-request-matcher-cookie-matchers-item-exact-values) - Optional List<br>Exact Values. A list of exact values to match the input against

<a id="api-rate-limit-api-endpoint-rules-request-matcher-cookie-matchers-item-regex-values"></a>&#x2022; [`regex_values`](#api-rate-limit-api-endpoint-rules-request-matcher-cookie-matchers-item-regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="api-rate-limit-api-endpoint-rules-request-matcher-cookie-matchers-item-transformers"></a>&#x2022; [`transformers`](#api-rate-limit-api-endpoint-rules-request-matcher-cookie-matchers-item-transformers) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### API Rate Limit API Endpoint Rules Request Matcher Headers

A [`headers`](#api-rate-limit-api-endpoint-rules-request-matcher-headers) block (within [`api_rate_limit.api_endpoint_rules.request_matcher`](#api-rate-limit-api-endpoint-rules-request-matcher)) supports the following:

<a id="api-rate-limit-api-endpoint-rules-request-matcher-headers-check-not-present"></a>&#x2022; [`check_not_present`](#api-rate-limit-api-endpoint-rules-request-matcher-headers-check-not-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-rate-limit-api-endpoint-rules-request-matcher-headers-check-present"></a>&#x2022; [`check_present`](#api-rate-limit-api-endpoint-rules-request-matcher-headers-check-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-rate-limit-api-endpoint-rules-request-matcher-headers-invert-matcher"></a>&#x2022; [`invert_matcher`](#api-rate-limit-api-endpoint-rules-request-matcher-headers-invert-matcher) - Optional Bool<br>Invert Header Matcher. Invert the match result

<a id="api-rate-limit-api-endpoint-rules-request-matcher-headers-item"></a>&#x2022; [`item`](#api-rate-limit-api-endpoint-rules-request-matcher-headers-item) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#api-rate-limit-api-endpoint-rules-request-matcher-headers-item) below.

<a id="api-rate-limit-api-endpoint-rules-request-matcher-headers-name"></a>&#x2022; [`name`](#api-rate-limit-api-endpoint-rules-request-matcher-headers-name) - Optional String<br>Header Name. A case-insensitive HTTP header name

#### API Rate Limit API Endpoint Rules Request Matcher Headers Item

An [`item`](#api-rate-limit-api-endpoint-rules-request-matcher-headers-item) block (within [`api_rate_limit.api_endpoint_rules.request_matcher.headers`](#api-rate-limit-api-endpoint-rules-request-matcher-headers)) supports the following:

<a id="api-rate-limit-api-endpoint-rules-request-matcher-headers-item-exact-values"></a>&#x2022; [`exact_values`](#api-rate-limit-api-endpoint-rules-request-matcher-headers-item-exact-values) - Optional List<br>Exact Values. A list of exact values to match the input against

<a id="api-rate-limit-api-endpoint-rules-request-matcher-headers-item-regex-values"></a>&#x2022; [`regex_values`](#api-rate-limit-api-endpoint-rules-request-matcher-headers-item-regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="api-rate-limit-api-endpoint-rules-request-matcher-headers-item-transformers"></a>&#x2022; [`transformers`](#api-rate-limit-api-endpoint-rules-request-matcher-headers-item-transformers) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### API Rate Limit API Endpoint Rules Request Matcher JWT Claims

A [`jwt_claims`](#api-rate-limit-api-endpoint-rules-request-matcher-jwt-claims) block (within [`api_rate_limit.api_endpoint_rules.request_matcher`](#api-rate-limit-api-endpoint-rules-request-matcher)) supports the following:

<a id="api-rate-limit-api-endpoint-rules-request-matcher-jwt-claims-check-not-present"></a>&#x2022; [`check_not_present`](#api-rate-limit-api-endpoint-rules-request-matcher-jwt-claims-check-not-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-rate-limit-api-endpoint-rules-request-matcher-jwt-claims-check-present"></a>&#x2022; [`check_present`](#api-rate-limit-api-endpoint-rules-request-matcher-jwt-claims-check-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-rate-limit-api-endpoint-rules-request-matcher-jwt-claims-invert-matcher"></a>&#x2022; [`invert_matcher`](#api-rate-limit-api-endpoint-rules-request-matcher-jwt-claims-invert-matcher) - Optional Bool<br>Invert Matcher. Invert the match result

<a id="api-rate-limit-api-endpoint-rules-request-matcher-jwt-claims-item"></a>&#x2022; [`item`](#api-rate-limit-api-endpoint-rules-request-matcher-jwt-claims-item) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#api-rate-limit-api-endpoint-rules-request-matcher-jwt-claims-item) below.

<a id="api-rate-limit-api-endpoint-rules-request-matcher-jwt-claims-name"></a>&#x2022; [`name`](#api-rate-limit-api-endpoint-rules-request-matcher-jwt-claims-name) - Optional String<br>JWT Claim Name. JWT claim name

#### API Rate Limit API Endpoint Rules Request Matcher JWT Claims Item

An [`item`](#api-rate-limit-api-endpoint-rules-request-matcher-jwt-claims-item) block (within [`api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims`](#api-rate-limit-api-endpoint-rules-request-matcher-jwt-claims)) supports the following:

<a id="api-rate-limit-api-endpoint-rules-request-matcher-jwt-claims-item-exact-values"></a>&#x2022; [`exact_values`](#api-rate-limit-api-endpoint-rules-request-matcher-jwt-claims-item-exact-values) - Optional List<br>Exact Values. A list of exact values to match the input against

<a id="api-rate-limit-api-endpoint-rules-request-matcher-jwt-claims-item-regex-values"></a>&#x2022; [`regex_values`](#api-rate-limit-api-endpoint-rules-request-matcher-jwt-claims-item-regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="api-rate-limit-api-endpoint-rules-request-matcher-jwt-claims-item-transformers"></a>&#x2022; [`transformers`](#api-rate-limit-api-endpoint-rules-request-matcher-jwt-claims-item-transformers) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### API Rate Limit API Endpoint Rules Request Matcher Query Params

A [`query_params`](#api-rate-limit-api-endpoint-rules-request-matcher-query-params) block (within [`api_rate_limit.api_endpoint_rules.request_matcher`](#api-rate-limit-api-endpoint-rules-request-matcher)) supports the following:

<a id="api-rate-limit-api-endpoint-rules-request-matcher-query-params-check-not-present"></a>&#x2022; [`check_not_present`](#api-rate-limit-api-endpoint-rules-request-matcher-query-params-check-not-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-rate-limit-api-endpoint-rules-request-matcher-query-params-check-present"></a>&#x2022; [`check_present`](#api-rate-limit-api-endpoint-rules-request-matcher-query-params-check-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-rate-limit-api-endpoint-rules-request-matcher-query-params-invert-matcher"></a>&#x2022; [`invert_matcher`](#api-rate-limit-api-endpoint-rules-request-matcher-query-params-invert-matcher) - Optional Bool<br>Invert Query Parameter Matcher. Invert the match result

<a id="api-rate-limit-api-endpoint-rules-request-matcher-query-params-item"></a>&#x2022; [`item`](#api-rate-limit-api-endpoint-rules-request-matcher-query-params-item) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#api-rate-limit-api-endpoint-rules-request-matcher-query-params-item) below.

<a id="api-rate-limit-api-endpoint-rules-request-matcher-query-params-key"></a>&#x2022; [`key`](#api-rate-limit-api-endpoint-rules-request-matcher-query-params-key) - Optional String<br>Query Parameter Name. A case-sensitive HTTP query parameter name

#### API Rate Limit API Endpoint Rules Request Matcher Query Params Item

An [`item`](#api-rate-limit-api-endpoint-rules-request-matcher-query-params-item) block (within [`api_rate_limit.api_endpoint_rules.request_matcher.query_params`](#api-rate-limit-api-endpoint-rules-request-matcher-query-params)) supports the following:

<a id="api-rate-limit-api-endpoint-rules-request-matcher-query-params-item-exact-values"></a>&#x2022; [`exact_values`](#api-rate-limit-api-endpoint-rules-request-matcher-query-params-item-exact-values) - Optional List<br>Exact Values. A list of exact values to match the input against

<a id="api-rate-limit-api-endpoint-rules-request-matcher-query-params-item-regex-values"></a>&#x2022; [`regex_values`](#api-rate-limit-api-endpoint-rules-request-matcher-query-params-item-regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="api-rate-limit-api-endpoint-rules-request-matcher-query-params-item-transformers"></a>&#x2022; [`transformers`](#api-rate-limit-api-endpoint-rules-request-matcher-query-params-item-transformers) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### API Rate Limit Bypass Rate Limiting Rules

A [`bypass_rate_limiting_rules`](#api-rate-limit-bypass-rate-limiting-rules) block (within [`api_rate_limit`](#api-rate-limit)) supports the following:

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules"></a>&#x2022; [`bypass_rate_limiting_rules`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules) - Optional Block<br>Bypass Rate Limiting. This category defines rules per URL or API group. If request matches any of these rules, skip Rate Limiting<br>See [Bypass Rate Limiting Rules](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules) below.

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules

A [`bypass_rate_limiting_rules`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules) block (within [`api_rate_limit.bypass_rate_limiting_rules`](#api-rate-limit-bypass-rate-limiting-rules)) supports the following:

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-any-domain"></a>&#x2022; [`any_domain`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-any-domain) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-any-url"></a>&#x2022; [`any_url`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-any-url) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-api-endpoint"></a>&#x2022; [`api_endpoint`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-api-endpoint) - Optional Block<br>API Endpoint. This defines API endpoint<br>See [API Endpoint](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-api-endpoint) below.

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-api-groups"></a>&#x2022; [`api_groups`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-api-groups) - Optional Block<br>API Groups<br>See [API Groups](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-api-groups) below.

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-base-path"></a>&#x2022; [`base_path`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-base-path) - Optional String<br>Base Path. The base path which this validation applies to

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher"></a>&#x2022; [`client_matcher`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher) - Optional Block<br>Client Matcher. Client conditions for matching a rule<br>See [Client Matcher](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher) below.

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher"></a>&#x2022; [`request_matcher`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher) - Optional Block<br>Request Matcher. Request conditions for matching a rule<br>See [Request Matcher](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher) below.

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-specific-domain"></a>&#x2022; [`specific_domain`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-specific-domain) - Optional String<br>Specific Domain. The rule will apply for a specific domain. For example: API.example.com

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules API Endpoint

An [`api_endpoint`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-api-endpoint) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules)) supports the following:

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-api-endpoint-methods"></a>&#x2022; [`methods`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-api-endpoint-methods) - Optional List  Defaults to `ANY`<br>Possible values are `ANY`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`, `COPY`<br>Methods. Methods to be matched

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-api-endpoint-path"></a>&#x2022; [`path`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-api-endpoint-path) - Optional String<br>Path. Path to be matched

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules API Groups

An [`api_groups`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-api-groups) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules)) supports the following:

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-api-groups-api-groups"></a>&#x2022; [`api_groups`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-api-groups-api-groups) - Optional List<br>API Groups

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Client Matcher

A [`client_matcher`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules)) supports the following:

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-any-client"></a>&#x2022; [`any_client`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-any-client) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-any-ip"></a>&#x2022; [`any_ip`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-any-ip) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-asn-list"></a>&#x2022; [`asn_list`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-asn-list) - Optional Block<br>ASN Match List. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer<br>See [Asn List](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-asn-list) below.

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-asn-matcher"></a>&#x2022; [`asn_matcher`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-asn-matcher) - Optional Block<br>ASN Matcher. Match any AS number contained in the list of bgp_asn_sets<br>See [Asn Matcher](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-asn-matcher) below.

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-client-selector"></a>&#x2022; [`client_selector`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-client-selector) - Optional Block<br>Label Selector. This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expresssions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects. Label selector is immutable. expressions is a list of strings of label selection expression. Each string has ',' separated values which are 'AND' and all strings are logically 'OR'. BNF for expression string `<selector-syntax>` ::= `<requirement>` | `<requirement>` ',' `<selector-syntax>` `<requirement>` ::= [!] KEY [ `<set-based-restriction>` | `<exact-match-restriction>` ] `<set-based-restriction>` ::= '' | `<inclusion-exclusion>` `<value-set>` `<inclusion-exclusion>` ::= `<inclusion>` | `<exclusion>` `<exclusion>` ::= 'notin' `<inclusion>` ::= 'in' `<value-set>` ::= '(' `<values>` ')' `<values>` ::= VALUE | VALUE ',' `<values>` `<exact-match-restriction>` ::= ['='|'=='|'!='] VALUE<br>See [Client Selector](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-client-selector) below.

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-ip-matcher"></a>&#x2022; [`ip_matcher`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-ip-matcher) - Optional Block<br>IP Prefix Matcher. Match any IP prefix contained in the list of ip_prefix_sets. The result of the match is inverted if invert_matcher is true<br>See [IP Matcher](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-ip-matcher) below.

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-ip-prefix-list"></a>&#x2022; [`ip_prefix_list`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-ip-prefix-list) - Optional Block<br>IP Prefix Match List. List of IP Prefix strings to match against<br>See [IP Prefix List](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-ip-prefix-list) below.

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-ip-threat-category-list"></a>&#x2022; [`ip_threat_category_list`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-ip-threat-category-list) - Optional Block<br>IP Threat Category List Type. List of IP threat categories<br>See [IP Threat Category List](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-ip-threat-category-list) below.

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-tls-fingerprint-matcher"></a>&#x2022; [`tls_fingerprint_matcher`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-tls-fingerprint-matcher) - Optional Block<br>TLS Fingerprint Matcher. A TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of supported positve match criteria includes a list of known classes of TLS fingerprints and a list of exact values. The match is considered successful if either of these positive criteria are satisfied and the input fingerprint is not one of the excluded values<br>See [TLS Fingerprint Matcher](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-tls-fingerprint-matcher) below.

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Client Matcher Asn List

An [`asn_list`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-asn-list) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher)) supports the following:

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-asn-list-as-numbers"></a>&#x2022; [`as_numbers`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-asn-list-as-numbers) - Optional List<br>AS Numbers. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Client Matcher Asn Matcher

An [`asn_matcher`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-asn-matcher) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher)) supports the following:

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-asn-matcher-asn-sets"></a>&#x2022; [`asn_sets`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-asn-matcher-asn-sets) - Optional Block<br>BGP ASN Sets. A list of references to bgp_asn_set objects<br>See [Asn Sets](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-asn-matcher-asn-sets) below.

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Client Matcher Asn Matcher Asn Sets

An [`asn_sets`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-asn-matcher-asn-sets) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-asn-matcher)) supports the following:

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-asn-matcher-asn-sets-kind"></a>&#x2022; [`kind`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-asn-matcher-asn-sets-kind) - Optional String<br>Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route')

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-asn-matcher-asn-sets-name"></a>&#x2022; [`name`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-asn-matcher-asn-sets-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-asn-matcher-asn-sets-namespace"></a>&#x2022; [`namespace`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-asn-matcher-asn-sets-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-asn-matcher-asn-sets-tenant"></a>&#x2022; [`tenant`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-asn-matcher-asn-sets-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-asn-matcher-asn-sets-uid"></a>&#x2022; [`uid`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-asn-matcher-asn-sets-uid) - Optional String<br>UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Client Matcher Client Selector

A [`client_selector`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-client-selector) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher)) supports the following:

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-client-selector-expressions"></a>&#x2022; [`expressions`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-client-selector-expressions) - Optional List<br>Selector Expression. expressions contains the kubernetes style label expression for selections

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Client Matcher IP Matcher

An [`ip_matcher`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-ip-matcher) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher)) supports the following:

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-ip-matcher-invert-matcher"></a>&#x2022; [`invert_matcher`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-ip-matcher-invert-matcher) - Optional Bool<br>Invert IP Matcher. Invert the match result

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-ip-matcher-prefix-sets"></a>&#x2022; [`prefix_sets`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-ip-matcher-prefix-sets) - Optional Block<br>IP Prefix Sets. A list of references to ip_prefix_set objects<br>See [Prefix Sets](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-ip-matcher-prefix-sets) below.

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Client Matcher IP Matcher Prefix Sets

A [`prefix_sets`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-ip-matcher-prefix-sets) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-ip-matcher)) supports the following:

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-ip-matcher-prefix-sets-kind"></a>&#x2022; [`kind`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-ip-matcher-prefix-sets-kind) - Optional String<br>Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route')

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-ip-matcher-prefix-sets-name"></a>&#x2022; [`name`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-ip-matcher-prefix-sets-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-ip-matcher-prefix-sets-namespace"></a>&#x2022; [`namespace`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-ip-matcher-prefix-sets-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-ip-matcher-prefix-sets-tenant"></a>&#x2022; [`tenant`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-ip-matcher-prefix-sets-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-ip-matcher-prefix-sets-uid"></a>&#x2022; [`uid`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-ip-matcher-prefix-sets-uid) - Optional String<br>UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Client Matcher IP Prefix List

An [`ip_prefix_list`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-ip-prefix-list) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher)) supports the following:

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-ip-prefix-list-invert-match"></a>&#x2022; [`invert_match`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-ip-prefix-list-invert-match) - Optional Bool<br>Invert Match Result. Invert the match result

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-ip-prefix-list-ip-prefixes"></a>&#x2022; [`ip_prefixes`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-ip-prefix-list-ip-prefixes) - Optional List<br>IPv4 Prefix List. List of IPv4 prefix strings

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Client Matcher IP Threat Category List

An [`ip_threat_category_list`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-ip-threat-category-list) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher)) supports the following:

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-ip-threat-category-list-ip-threat-categories"></a>&#x2022; [`ip_threat_categories`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-ip-threat-category-list-ip-threat-categories) - Optional List  Defaults to `SPAM_SOURCES`<br>Possible values are `SPAM_SOURCES`, `WINDOWS_EXPLOITS`, `WEB_ATTACKS`, `BOTNETS`, `SCANNERS`, `REPUTATION`, `PHISHING`, `PROXY`, `MOBILE_THREATS`, `TOR_PROXY`, `DENIAL_OF_SERVICE`, `NETWORK`<br>List of IP Threat Categories to choose. The IP threat categories is obtained from the list and is used to auto-generate equivalent label selection expressions

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Client Matcher TLS Fingerprint Matcher

A [`tls_fingerprint_matcher`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-tls-fingerprint-matcher) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher)) supports the following:

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-tls-fingerprint-matcher-classes"></a>&#x2022; [`classes`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-tls-fingerprint-matcher-classes) - Optional List  Defaults to `TLS_FINGERPRINT_NONE`<br>Possible values are `TLS_FINGERPRINT_NONE`, `ANY_MALICIOUS_FINGERPRINT`, `ADWARE`, `ADWIND`, `DRIDEX`, `GOOTKIT`, `GOZI`, `JBIFROST`, `QUAKBOT`, `RANSOMWARE`, `TROLDESH`, `TOFSEE`, `TORRENTLOCKER`, `TRICKBOT`<br>TLS fingerprint classes. A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-tls-fingerprint-matcher-exact-values"></a>&#x2022; [`exact_values`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-tls-fingerprint-matcher-exact-values) - Optional List<br>Exact Values. A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-tls-fingerprint-matcher-excluded-values"></a>&#x2022; [`excluded_values`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-tls-fingerprint-matcher-excluded-values) - Optional List<br>Excluded Values. A list of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can be used to skip known false positives when using one or more known TLS fingerprint classes in the enclosing matcher

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Request Matcher

A [`request_matcher`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules)) supports the following:

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-cookie-matchers"></a>&#x2022; [`cookie_matchers`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-cookie-matchers) - Optional Block<br>Cookie Matchers. A list of predicates for all cookies that need to be matched. The criteria for matching each cookie is described in individual instances of CookieMatcherType. The actual cookie values are extracted from the request API as a list of strings for each cookie name. Note that all specified cookie matcher predicates must evaluate to true<br>See [Cookie Matchers](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-cookie-matchers) below.

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-headers"></a>&#x2022; [`headers`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-headers) - Optional Block<br>HTTP Headers. A list of predicates for various HTTP headers that need to match. The criteria for matching each HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values are extracted from the request API as a list of strings for each HTTP header type. Note that all specified header predicates must evaluate to true<br>See [Headers](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-headers) below.

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-jwt-claims"></a>&#x2022; [`jwt_claims`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-jwt-claims) - Optional Block<br>JWT Claims. A list of predicates for various JWT claims that need to match. The criteria for matching each JWT claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are extracted from the JWT payload as a list of strings. Note that all specified JWT claim predicates must evaluate to true. Note that this feature only works on LBs with JWT Validation feature enabled<br>See [JWT Claims](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-jwt-claims) below.

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-query-params"></a>&#x2022; [`query_params`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-query-params) - Optional Block<br>HTTP Query Parameters. A list of predicates for all query parameters that need to be matched. The criteria for matching each query parameter are described in individual instances of QueryParameterMatcherType. The actual query parameter values are extracted from the request API as a list of strings for each query parameter name. Note that all specified query parameter predicates must evaluate to true<br>See [Query Params](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-query-params) below.

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Request Matcher Cookie Matchers

A [`cookie_matchers`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-cookie-matchers) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher)) supports the following:

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-cookie-matchers-check-not-present"></a>&#x2022; [`check_not_present`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-cookie-matchers-check-not-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-cookie-matchers-check-present"></a>&#x2022; [`check_present`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-cookie-matchers-check-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-cookie-matchers-invert-matcher"></a>&#x2022; [`invert_matcher`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-cookie-matchers-invert-matcher) - Optional Bool<br>Invert Matcher. Invert Match of the expression defined

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-cookie-matchers-item"></a>&#x2022; [`item`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-cookie-matchers-item) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-cookie-matchers-item) below.

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-cookie-matchers-name"></a>&#x2022; [`name`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-cookie-matchers-name) - Optional String<br>Cookie Name. A case-sensitive cookie name

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Request Matcher Cookie Matchers Item

An [`item`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-cookie-matchers-item) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-cookie-matchers)) supports the following:

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-cookie-matchers-item-exact-values"></a>&#x2022; [`exact_values`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-cookie-matchers-item-exact-values) - Optional List<br>Exact Values. A list of exact values to match the input against

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-cookie-matchers-item-regex-values"></a>&#x2022; [`regex_values`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-cookie-matchers-item-regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-cookie-matchers-item-transformers"></a>&#x2022; [`transformers`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-cookie-matchers-item-transformers) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Request Matcher Headers

A [`headers`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-headers) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher)) supports the following:

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-headers-check-not-present"></a>&#x2022; [`check_not_present`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-headers-check-not-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-headers-check-present"></a>&#x2022; [`check_present`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-headers-check-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-headers-invert-matcher"></a>&#x2022; [`invert_matcher`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-headers-invert-matcher) - Optional Bool<br>Invert Header Matcher. Invert the match result

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-headers-item"></a>&#x2022; [`item`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-headers-item) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-headers-item) below.

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-headers-name"></a>&#x2022; [`name`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-headers-name) - Optional String<br>Header Name. A case-insensitive HTTP header name

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Request Matcher Headers Item

An [`item`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-headers-item) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-headers)) supports the following:

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-headers-item-exact-values"></a>&#x2022; [`exact_values`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-headers-item-exact-values) - Optional List<br>Exact Values. A list of exact values to match the input against

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-headers-item-regex-values"></a>&#x2022; [`regex_values`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-headers-item-regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-headers-item-transformers"></a>&#x2022; [`transformers`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-headers-item-transformers) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Request Matcher JWT Claims

A [`jwt_claims`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-jwt-claims) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher)) supports the following:

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-jwt-claims-check-not-present"></a>&#x2022; [`check_not_present`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-jwt-claims-check-not-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-jwt-claims-check-present"></a>&#x2022; [`check_present`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-jwt-claims-check-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-jwt-claims-invert-matcher"></a>&#x2022; [`invert_matcher`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-jwt-claims-invert-matcher) - Optional Bool<br>Invert Matcher. Invert the match result

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-jwt-claims-item"></a>&#x2022; [`item`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-jwt-claims-item) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-jwt-claims-item) below.

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-jwt-claims-name"></a>&#x2022; [`name`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-jwt-claims-name) - Optional String<br>JWT Claim Name. JWT claim name

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Request Matcher JWT Claims Item

An [`item`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-jwt-claims-item) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-jwt-claims)) supports the following:

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-jwt-claims-item-exact-values"></a>&#x2022; [`exact_values`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-jwt-claims-item-exact-values) - Optional List<br>Exact Values. A list of exact values to match the input against

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-jwt-claims-item-regex-values"></a>&#x2022; [`regex_values`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-jwt-claims-item-regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-jwt-claims-item-transformers"></a>&#x2022; [`transformers`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-jwt-claims-item-transformers) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Request Matcher Query Params

A [`query_params`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-query-params) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher)) supports the following:

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-query-params-check-not-present"></a>&#x2022; [`check_not_present`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-query-params-check-not-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-query-params-check-present"></a>&#x2022; [`check_present`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-query-params-check-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-query-params-invert-matcher"></a>&#x2022; [`invert_matcher`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-query-params-invert-matcher) - Optional Bool<br>Invert Query Parameter Matcher. Invert the match result

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-query-params-item"></a>&#x2022; [`item`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-query-params-item) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-query-params-item) below.

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-query-params-key"></a>&#x2022; [`key`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-query-params-key) - Optional String<br>Query Parameter Name. A case-sensitive HTTP query parameter name

#### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Request Matcher Query Params Item

An [`item`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-query-params-item) block (within [`api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-query-params)) supports the following:

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-query-params-item-exact-values"></a>&#x2022; [`exact_values`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-query-params-item-exact-values) - Optional List<br>Exact Values. A list of exact values to match the input against

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-query-params-item-regex-values"></a>&#x2022; [`regex_values`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-query-params-item-regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-query-params-item-transformers"></a>&#x2022; [`transformers`](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-query-params-item-transformers) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### API Rate Limit Custom IP Allowed List

A [`custom_ip_allowed_list`](#api-rate-limit-custom-ip-allowed-list) block (within [`api_rate_limit`](#api-rate-limit)) supports the following:

<a id="api-rate-limit-custom-ip-allowed-list-rate-limiter-allowed-prefixes"></a>&#x2022; [`rate_limiter_allowed_prefixes`](#api-rate-limit-custom-ip-allowed-list-rate-limiter-allowed-prefixes) - Optional Block<br>List of IP Prefix Sets. References to ip_prefix_set objects. Requests from source IP addresses that are covered by one of the allowed IP Prefixes are not subjected to rate limiting<br>See [Rate Limiter Allowed Prefixes](#api-rate-limit-custom-ip-allowed-list-rate-limiter-allowed-prefixes) below.

#### API Rate Limit Custom IP Allowed List Rate Limiter Allowed Prefixes

A [`rate_limiter_allowed_prefixes`](#api-rate-limit-custom-ip-allowed-list-rate-limiter-allowed-prefixes) block (within [`api_rate_limit.custom_ip_allowed_list`](#api-rate-limit-custom-ip-allowed-list)) supports the following:

<a id="api-rate-limit-custom-ip-allowed-list-rate-limiter-allowed-prefixes-name"></a>&#x2022; [`name`](#api-rate-limit-custom-ip-allowed-list-rate-limiter-allowed-prefixes-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="api-rate-limit-custom-ip-allowed-list-rate-limiter-allowed-prefixes-namespace"></a>&#x2022; [`namespace`](#api-rate-limit-custom-ip-allowed-list-rate-limiter-allowed-prefixes-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="api-rate-limit-custom-ip-allowed-list-rate-limiter-allowed-prefixes-tenant"></a>&#x2022; [`tenant`](#api-rate-limit-custom-ip-allowed-list-rate-limiter-allowed-prefixes-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### API Rate Limit IP Allowed List

An [`ip_allowed_list`](#api-rate-limit-ip-allowed-list) block (within [`api_rate_limit`](#api-rate-limit)) supports the following:

<a id="api-rate-limit-ip-allowed-list-prefixes"></a>&#x2022; [`prefixes`](#api-rate-limit-ip-allowed-list-prefixes) - Optional List<br>IPv4 Prefix List. List of IPv4 prefixes that represent an endpoint

#### API Rate Limit Server URL Rules

A [`server_url_rules`](#api-rate-limit-server-url-rules) block (within [`api_rate_limit`](#api-rate-limit)) supports the following:

<a id="api-rate-limit-server-url-rules-any-domain"></a>&#x2022; [`any_domain`](#api-rate-limit-server-url-rules-any-domain) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-rate-limit-server-url-rules-api-group"></a>&#x2022; [`api_group`](#api-rate-limit-server-url-rules-api-group) - Optional String<br>API Group. API groups derived from API Definition swaggers. For example oas-all-operations including all paths and methods from the swaggers, oas-base-urls covering all requests under base-paths from the swaggers. Custom groups can be created if user tags paths or operations with 'x-volterra-API-group' extensions inside swaggers

<a id="api-rate-limit-server-url-rules-base-path"></a>&#x2022; [`base_path`](#api-rate-limit-server-url-rules-base-path) - Optional String<br>Base Path. Prefix of the request path

<a id="api-rate-limit-server-url-rules-client-matcher"></a>&#x2022; [`client_matcher`](#api-rate-limit-server-url-rules-client-matcher) - Optional Block<br>Client Matcher. Client conditions for matching a rule<br>See [Client Matcher](#api-rate-limit-server-url-rules-client-matcher) below.

<a id="api-rate-limit-server-url-rules-inline-rate-limiter"></a>&#x2022; [`inline_rate_limiter`](#api-rate-limit-server-url-rules-inline-rate-limiter) - Optional Block<br>InlineRateLimiter<br>See [Inline Rate Limiter](#api-rate-limit-server-url-rules-inline-rate-limiter) below.

<a id="api-rate-limit-server-url-rules-ref-rate-limiter"></a>&#x2022; [`ref_rate_limiter`](#api-rate-limit-server-url-rules-ref-rate-limiter) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Ref Rate Limiter](#api-rate-limit-server-url-rules-ref-rate-limiter) below.

<a id="api-rate-limit-server-url-rules-request-matcher"></a>&#x2022; [`request_matcher`](#api-rate-limit-server-url-rules-request-matcher) - Optional Block<br>Request Matcher. Request conditions for matching a rule<br>See [Request Matcher](#api-rate-limit-server-url-rules-request-matcher) below.

<a id="api-rate-limit-server-url-rules-specific-domain"></a>&#x2022; [`specific_domain`](#api-rate-limit-server-url-rules-specific-domain) - Optional String<br>Specific Domain. The rule will apply for a specific domain

#### API Rate Limit Server URL Rules Client Matcher

A [`client_matcher`](#api-rate-limit-server-url-rules-client-matcher) block (within [`api_rate_limit.server_url_rules`](#api-rate-limit-server-url-rules)) supports the following:

<a id="api-rate-limit-server-url-rules-client-matcher-any-client"></a>&#x2022; [`any_client`](#api-rate-limit-server-url-rules-client-matcher-any-client) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-rate-limit-server-url-rules-client-matcher-any-ip"></a>&#x2022; [`any_ip`](#api-rate-limit-server-url-rules-client-matcher-any-ip) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-rate-limit-server-url-rules-client-matcher-asn-list"></a>&#x2022; [`asn_list`](#api-rate-limit-server-url-rules-client-matcher-asn-list) - Optional Block<br>ASN Match List. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer<br>See [Asn List](#api-rate-limit-server-url-rules-client-matcher-asn-list) below.

<a id="api-rate-limit-server-url-rules-client-matcher-asn-matcher"></a>&#x2022; [`asn_matcher`](#api-rate-limit-server-url-rules-client-matcher-asn-matcher) - Optional Block<br>ASN Matcher. Match any AS number contained in the list of bgp_asn_sets<br>See [Asn Matcher](#api-rate-limit-server-url-rules-client-matcher-asn-matcher) below.

<a id="api-rate-limit-server-url-rules-client-matcher-client-selector"></a>&#x2022; [`client_selector`](#api-rate-limit-server-url-rules-client-matcher-client-selector) - Optional Block<br>Label Selector. This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expresssions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects. Label selector is immutable. expressions is a list of strings of label selection expression. Each string has ',' separated values which are 'AND' and all strings are logically 'OR'. BNF for expression string `<selector-syntax>` ::= `<requirement>` | `<requirement>` ',' `<selector-syntax>` `<requirement>` ::= [!] KEY [ `<set-based-restriction>` | `<exact-match-restriction>` ] `<set-based-restriction>` ::= '' | `<inclusion-exclusion>` `<value-set>` `<inclusion-exclusion>` ::= `<inclusion>` | `<exclusion>` `<exclusion>` ::= 'notin' `<inclusion>` ::= 'in' `<value-set>` ::= '(' `<values>` ')' `<values>` ::= VALUE | VALUE ',' `<values>` `<exact-match-restriction>` ::= ['='|'=='|'!='] VALUE<br>See [Client Selector](#api-rate-limit-server-url-rules-client-matcher-client-selector) below.

<a id="api-rate-limit-server-url-rules-client-matcher-ip-matcher"></a>&#x2022; [`ip_matcher`](#api-rate-limit-server-url-rules-client-matcher-ip-matcher) - Optional Block<br>IP Prefix Matcher. Match any IP prefix contained in the list of ip_prefix_sets. The result of the match is inverted if invert_matcher is true<br>See [IP Matcher](#api-rate-limit-server-url-rules-client-matcher-ip-matcher) below.

<a id="api-rate-limit-server-url-rules-client-matcher-ip-prefix-list"></a>&#x2022; [`ip_prefix_list`](#api-rate-limit-server-url-rules-client-matcher-ip-prefix-list) - Optional Block<br>IP Prefix Match List. List of IP Prefix strings to match against<br>See [IP Prefix List](#api-rate-limit-server-url-rules-client-matcher-ip-prefix-list) below.

<a id="api-rate-limit-server-url-rules-client-matcher-ip-threat-category-list"></a>&#x2022; [`ip_threat_category_list`](#api-rate-limit-server-url-rules-client-matcher-ip-threat-category-list) - Optional Block<br>IP Threat Category List Type. List of IP threat categories<br>See [IP Threat Category List](#api-rate-limit-server-url-rules-client-matcher-ip-threat-category-list) below.

<a id="api-rate-limit-server-url-rules-client-matcher-tls-fingerprint-matcher"></a>&#x2022; [`tls_fingerprint_matcher`](#api-rate-limit-server-url-rules-client-matcher-tls-fingerprint-matcher) - Optional Block<br>TLS Fingerprint Matcher. A TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of supported positve match criteria includes a list of known classes of TLS fingerprints and a list of exact values. The match is considered successful if either of these positive criteria are satisfied and the input fingerprint is not one of the excluded values<br>See [TLS Fingerprint Matcher](#api-rate-limit-server-url-rules-client-matcher-tls-fingerprint-matcher) below.

#### API Rate Limit Server URL Rules Client Matcher Asn List

An [`asn_list`](#api-rate-limit-server-url-rules-client-matcher-asn-list) block (within [`api_rate_limit.server_url_rules.client_matcher`](#api-rate-limit-server-url-rules-client-matcher)) supports the following:

<a id="api-rate-limit-server-url-rules-client-matcher-asn-list-as-numbers"></a>&#x2022; [`as_numbers`](#api-rate-limit-server-url-rules-client-matcher-asn-list-as-numbers) - Optional List<br>AS Numbers. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer

#### API Rate Limit Server URL Rules Client Matcher Asn Matcher

An [`asn_matcher`](#api-rate-limit-server-url-rules-client-matcher-asn-matcher) block (within [`api_rate_limit.server_url_rules.client_matcher`](#api-rate-limit-server-url-rules-client-matcher)) supports the following:

<a id="api-rate-limit-server-url-rules-client-matcher-asn-matcher-asn-sets"></a>&#x2022; [`asn_sets`](#api-rate-limit-server-url-rules-client-matcher-asn-matcher-asn-sets) - Optional Block<br>BGP ASN Sets. A list of references to bgp_asn_set objects<br>See [Asn Sets](#api-rate-limit-server-url-rules-client-matcher-asn-matcher-asn-sets) below.

#### API Rate Limit Server URL Rules Client Matcher Asn Matcher Asn Sets

An [`asn_sets`](#api-rate-limit-server-url-rules-client-matcher-asn-matcher-asn-sets) block (within [`api_rate_limit.server_url_rules.client_matcher.asn_matcher`](#api-rate-limit-server-url-rules-client-matcher-asn-matcher)) supports the following:

<a id="api-rate-limit-server-url-rules-client-matcher-asn-matcher-asn-sets-kind"></a>&#x2022; [`kind`](#api-rate-limit-server-url-rules-client-matcher-asn-matcher-asn-sets-kind) - Optional String<br>Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route')

<a id="api-rate-limit-server-url-rules-client-matcher-asn-matcher-asn-sets-name"></a>&#x2022; [`name`](#api-rate-limit-server-url-rules-client-matcher-asn-matcher-asn-sets-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="api-rate-limit-server-url-rules-client-matcher-asn-matcher-asn-sets-namespace"></a>&#x2022; [`namespace`](#api-rate-limit-server-url-rules-client-matcher-asn-matcher-asn-sets-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="api-rate-limit-server-url-rules-client-matcher-asn-matcher-asn-sets-tenant"></a>&#x2022; [`tenant`](#api-rate-limit-server-url-rules-client-matcher-asn-matcher-asn-sets-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="api-rate-limit-server-url-rules-client-matcher-asn-matcher-asn-sets-uid"></a>&#x2022; [`uid`](#api-rate-limit-server-url-rules-client-matcher-asn-matcher-asn-sets-uid) - Optional String<br>UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid

#### API Rate Limit Server URL Rules Client Matcher Client Selector

A [`client_selector`](#api-rate-limit-server-url-rules-client-matcher-client-selector) block (within [`api_rate_limit.server_url_rules.client_matcher`](#api-rate-limit-server-url-rules-client-matcher)) supports the following:

<a id="api-rate-limit-server-url-rules-client-matcher-client-selector-expressions"></a>&#x2022; [`expressions`](#api-rate-limit-server-url-rules-client-matcher-client-selector-expressions) - Optional List<br>Selector Expression. expressions contains the kubernetes style label expression for selections

#### API Rate Limit Server URL Rules Client Matcher IP Matcher

An [`ip_matcher`](#api-rate-limit-server-url-rules-client-matcher-ip-matcher) block (within [`api_rate_limit.server_url_rules.client_matcher`](#api-rate-limit-server-url-rules-client-matcher)) supports the following:

<a id="api-rate-limit-server-url-rules-client-matcher-ip-matcher-invert-matcher"></a>&#x2022; [`invert_matcher`](#api-rate-limit-server-url-rules-client-matcher-ip-matcher-invert-matcher) - Optional Bool<br>Invert IP Matcher. Invert the match result

<a id="api-rate-limit-server-url-rules-client-matcher-ip-matcher-prefix-sets"></a>&#x2022; [`prefix_sets`](#api-rate-limit-server-url-rules-client-matcher-ip-matcher-prefix-sets) - Optional Block<br>IP Prefix Sets. A list of references to ip_prefix_set objects<br>See [Prefix Sets](#api-rate-limit-server-url-rules-client-matcher-ip-matcher-prefix-sets) below.

#### API Rate Limit Server URL Rules Client Matcher IP Matcher Prefix Sets

A [`prefix_sets`](#api-rate-limit-server-url-rules-client-matcher-ip-matcher-prefix-sets) block (within [`api_rate_limit.server_url_rules.client_matcher.ip_matcher`](#api-rate-limit-server-url-rules-client-matcher-ip-matcher)) supports the following:

<a id="api-rate-limit-server-url-rules-client-matcher-ip-matcher-prefix-sets-kind"></a>&#x2022; [`kind`](#api-rate-limit-server-url-rules-client-matcher-ip-matcher-prefix-sets-kind) - Optional String<br>Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route')

<a id="api-rate-limit-server-url-rules-client-matcher-ip-matcher-prefix-sets-name"></a>&#x2022; [`name`](#api-rate-limit-server-url-rules-client-matcher-ip-matcher-prefix-sets-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="api-rate-limit-server-url-rules-client-matcher-ip-matcher-prefix-sets-namespace"></a>&#x2022; [`namespace`](#api-rate-limit-server-url-rules-client-matcher-ip-matcher-prefix-sets-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="api-rate-limit-server-url-rules-client-matcher-ip-matcher-prefix-sets-tenant"></a>&#x2022; [`tenant`](#api-rate-limit-server-url-rules-client-matcher-ip-matcher-prefix-sets-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="api-rate-limit-server-url-rules-client-matcher-ip-matcher-prefix-sets-uid"></a>&#x2022; [`uid`](#api-rate-limit-server-url-rules-client-matcher-ip-matcher-prefix-sets-uid) - Optional String<br>UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid

#### API Rate Limit Server URL Rules Client Matcher IP Prefix List

An [`ip_prefix_list`](#api-rate-limit-server-url-rules-client-matcher-ip-prefix-list) block (within [`api_rate_limit.server_url_rules.client_matcher`](#api-rate-limit-server-url-rules-client-matcher)) supports the following:

<a id="api-rate-limit-server-url-rules-client-matcher-ip-prefix-list-invert-match"></a>&#x2022; [`invert_match`](#api-rate-limit-server-url-rules-client-matcher-ip-prefix-list-invert-match) - Optional Bool<br>Invert Match Result. Invert the match result

<a id="api-rate-limit-server-url-rules-client-matcher-ip-prefix-list-ip-prefixes"></a>&#x2022; [`ip_prefixes`](#api-rate-limit-server-url-rules-client-matcher-ip-prefix-list-ip-prefixes) - Optional List<br>IPv4 Prefix List. List of IPv4 prefix strings

#### API Rate Limit Server URL Rules Client Matcher IP Threat Category List

An [`ip_threat_category_list`](#api-rate-limit-server-url-rules-client-matcher-ip-threat-category-list) block (within [`api_rate_limit.server_url_rules.client_matcher`](#api-rate-limit-server-url-rules-client-matcher)) supports the following:

<a id="api-rate-limit-server-url-rules-client-matcher-ip-threat-category-list-ip-threat-categories"></a>&#x2022; [`ip_threat_categories`](#api-rate-limit-server-url-rules-client-matcher-ip-threat-category-list-ip-threat-categories) - Optional List  Defaults to `SPAM_SOURCES`<br>Possible values are `SPAM_SOURCES`, `WINDOWS_EXPLOITS`, `WEB_ATTACKS`, `BOTNETS`, `SCANNERS`, `REPUTATION`, `PHISHING`, `PROXY`, `MOBILE_THREATS`, `TOR_PROXY`, `DENIAL_OF_SERVICE`, `NETWORK`<br>List of IP Threat Categories to choose. The IP threat categories is obtained from the list and is used to auto-generate equivalent label selection expressions

#### API Rate Limit Server URL Rules Client Matcher TLS Fingerprint Matcher

A [`tls_fingerprint_matcher`](#api-rate-limit-server-url-rules-client-matcher-tls-fingerprint-matcher) block (within [`api_rate_limit.server_url_rules.client_matcher`](#api-rate-limit-server-url-rules-client-matcher)) supports the following:

<a id="api-rate-limit-server-url-rules-client-matcher-tls-fingerprint-matcher-classes"></a>&#x2022; [`classes`](#api-rate-limit-server-url-rules-client-matcher-tls-fingerprint-matcher-classes) - Optional List  Defaults to `TLS_FINGERPRINT_NONE`<br>Possible values are `TLS_FINGERPRINT_NONE`, `ANY_MALICIOUS_FINGERPRINT`, `ADWARE`, `ADWIND`, `DRIDEX`, `GOOTKIT`, `GOZI`, `JBIFROST`, `QUAKBOT`, `RANSOMWARE`, `TROLDESH`, `TOFSEE`, `TORRENTLOCKER`, `TRICKBOT`<br>TLS fingerprint classes. A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against

<a id="api-rate-limit-server-url-rules-client-matcher-tls-fingerprint-matcher-exact-values"></a>&#x2022; [`exact_values`](#api-rate-limit-server-url-rules-client-matcher-tls-fingerprint-matcher-exact-values) - Optional List<br>Exact Values. A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against

<a id="api-rate-limit-server-url-rules-client-matcher-tls-fingerprint-matcher-excluded-values"></a>&#x2022; [`excluded_values`](#api-rate-limit-server-url-rules-client-matcher-tls-fingerprint-matcher-excluded-values) - Optional List<br>Excluded Values. A list of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can be used to skip known false positives when using one or more known TLS fingerprint classes in the enclosing matcher

#### API Rate Limit Server URL Rules Inline Rate Limiter

An [`inline_rate_limiter`](#api-rate-limit-server-url-rules-inline-rate-limiter) block (within [`api_rate_limit.server_url_rules`](#api-rate-limit-server-url-rules)) supports the following:

<a id="api-rate-limit-server-url-rules-inline-rate-limiter-ref-user-id"></a>&#x2022; [`ref_user_id`](#api-rate-limit-server-url-rules-inline-rate-limiter-ref-user-id) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Ref User Id](#api-rate-limit-server-url-rules-inline-rate-limiter-ref-user-id) below.

<a id="api-rate-limit-server-url-rules-inline-rate-limiter-threshold"></a>&#x2022; [`threshold`](#api-rate-limit-server-url-rules-inline-rate-limiter-threshold) - Optional Number<br>Threshold. The total number of allowed requests for 1 unit (e.g. SECOND/MINUTE/HOUR etc.) of the specified period

<a id="api-rate-limit-server-url-rules-inline-rate-limiter-unit"></a>&#x2022; [`unit`](#api-rate-limit-server-url-rules-inline-rate-limiter-unit) - Optional String  Defaults to `SECOND`<br>Possible values are `SECOND`, `MINUTE`, `HOUR`<br>Rate Limit Period Unit. Unit for the period per which the rate limit is applied. - SECOND: Second Rate limit period unit is seconds - MINUTE: Minute Rate limit period unit is minutes - HOUR: Hour Rate limit period unit is hours - DAY: Day Rate limit period unit is days

<a id="api-rate-limit-server-url-rules-inline-rate-limiter-use-http-lb-user-id"></a>&#x2022; [`use_http_lb_user_id`](#api-rate-limit-server-url-rules-inline-rate-limiter-use-http-lb-user-id) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### API Rate Limit Server URL Rules Inline Rate Limiter Ref User Id

A [`ref_user_id`](#api-rate-limit-server-url-rules-inline-rate-limiter-ref-user-id) block (within [`api_rate_limit.server_url_rules.inline_rate_limiter`](#api-rate-limit-server-url-rules-inline-rate-limiter)) supports the following:

<a id="api-rate-limit-server-url-rules-inline-rate-limiter-ref-user-id-name"></a>&#x2022; [`name`](#api-rate-limit-server-url-rules-inline-rate-limiter-ref-user-id-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="api-rate-limit-server-url-rules-inline-rate-limiter-ref-user-id-namespace"></a>&#x2022; [`namespace`](#api-rate-limit-server-url-rules-inline-rate-limiter-ref-user-id-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="api-rate-limit-server-url-rules-inline-rate-limiter-ref-user-id-tenant"></a>&#x2022; [`tenant`](#api-rate-limit-server-url-rules-inline-rate-limiter-ref-user-id-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### API Rate Limit Server URL Rules Ref Rate Limiter

A [`ref_rate_limiter`](#api-rate-limit-server-url-rules-ref-rate-limiter) block (within [`api_rate_limit.server_url_rules`](#api-rate-limit-server-url-rules)) supports the following:

<a id="api-rate-limit-server-url-rules-ref-rate-limiter-name"></a>&#x2022; [`name`](#api-rate-limit-server-url-rules-ref-rate-limiter-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="api-rate-limit-server-url-rules-ref-rate-limiter-namespace"></a>&#x2022; [`namespace`](#api-rate-limit-server-url-rules-ref-rate-limiter-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="api-rate-limit-server-url-rules-ref-rate-limiter-tenant"></a>&#x2022; [`tenant`](#api-rate-limit-server-url-rules-ref-rate-limiter-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### API Rate Limit Server URL Rules Request Matcher

A [`request_matcher`](#api-rate-limit-server-url-rules-request-matcher) block (within [`api_rate_limit.server_url_rules`](#api-rate-limit-server-url-rules)) supports the following:

<a id="api-rate-limit-server-url-rules-request-matcher-cookie-matchers"></a>&#x2022; [`cookie_matchers`](#api-rate-limit-server-url-rules-request-matcher-cookie-matchers) - Optional Block<br>Cookie Matchers. A list of predicates for all cookies that need to be matched. The criteria for matching each cookie is described in individual instances of CookieMatcherType. The actual cookie values are extracted from the request API as a list of strings for each cookie name. Note that all specified cookie matcher predicates must evaluate to true<br>See [Cookie Matchers](#api-rate-limit-server-url-rules-request-matcher-cookie-matchers) below.

<a id="api-rate-limit-server-url-rules-request-matcher-headers"></a>&#x2022; [`headers`](#api-rate-limit-server-url-rules-request-matcher-headers) - Optional Block<br>HTTP Headers. A list of predicates for various HTTP headers that need to match. The criteria for matching each HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values are extracted from the request API as a list of strings for each HTTP header type. Note that all specified header predicates must evaluate to true<br>See [Headers](#api-rate-limit-server-url-rules-request-matcher-headers) below.

<a id="api-rate-limit-server-url-rules-request-matcher-jwt-claims"></a>&#x2022; [`jwt_claims`](#api-rate-limit-server-url-rules-request-matcher-jwt-claims) - Optional Block<br>JWT Claims. A list of predicates for various JWT claims that need to match. The criteria for matching each JWT claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are extracted from the JWT payload as a list of strings. Note that all specified JWT claim predicates must evaluate to true. Note that this feature only works on LBs with JWT Validation feature enabled<br>See [JWT Claims](#api-rate-limit-server-url-rules-request-matcher-jwt-claims) below.

<a id="api-rate-limit-server-url-rules-request-matcher-query-params"></a>&#x2022; [`query_params`](#api-rate-limit-server-url-rules-request-matcher-query-params) - Optional Block<br>HTTP Query Parameters. A list of predicates for all query parameters that need to be matched. The criteria for matching each query parameter are described in individual instances of QueryParameterMatcherType. The actual query parameter values are extracted from the request API as a list of strings for each query parameter name. Note that all specified query parameter predicates must evaluate to true<br>See [Query Params](#api-rate-limit-server-url-rules-request-matcher-query-params) below.

#### API Rate Limit Server URL Rules Request Matcher Cookie Matchers

A [`cookie_matchers`](#api-rate-limit-server-url-rules-request-matcher-cookie-matchers) block (within [`api_rate_limit.server_url_rules.request_matcher`](#api-rate-limit-server-url-rules-request-matcher)) supports the following:

<a id="api-rate-limit-server-url-rules-request-matcher-cookie-matchers-check-not-present"></a>&#x2022; [`check_not_present`](#api-rate-limit-server-url-rules-request-matcher-cookie-matchers-check-not-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-rate-limit-server-url-rules-request-matcher-cookie-matchers-check-present"></a>&#x2022; [`check_present`](#api-rate-limit-server-url-rules-request-matcher-cookie-matchers-check-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-rate-limit-server-url-rules-request-matcher-cookie-matchers-invert-matcher"></a>&#x2022; [`invert_matcher`](#api-rate-limit-server-url-rules-request-matcher-cookie-matchers-invert-matcher) - Optional Bool<br>Invert Matcher. Invert Match of the expression defined

<a id="api-rate-limit-server-url-rules-request-matcher-cookie-matchers-item"></a>&#x2022; [`item`](#api-rate-limit-server-url-rules-request-matcher-cookie-matchers-item) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#api-rate-limit-server-url-rules-request-matcher-cookie-matchers-item) below.

<a id="api-rate-limit-server-url-rules-request-matcher-cookie-matchers-name"></a>&#x2022; [`name`](#api-rate-limit-server-url-rules-request-matcher-cookie-matchers-name) - Optional String<br>Cookie Name. A case-sensitive cookie name

#### API Rate Limit Server URL Rules Request Matcher Cookie Matchers Item

An [`item`](#api-rate-limit-server-url-rules-request-matcher-cookie-matchers-item) block (within [`api_rate_limit.server_url_rules.request_matcher.cookie_matchers`](#api-rate-limit-server-url-rules-request-matcher-cookie-matchers)) supports the following:

<a id="api-rate-limit-server-url-rules-request-matcher-cookie-matchers-item-exact-values"></a>&#x2022; [`exact_values`](#api-rate-limit-server-url-rules-request-matcher-cookie-matchers-item-exact-values) - Optional List<br>Exact Values. A list of exact values to match the input against

<a id="api-rate-limit-server-url-rules-request-matcher-cookie-matchers-item-regex-values"></a>&#x2022; [`regex_values`](#api-rate-limit-server-url-rules-request-matcher-cookie-matchers-item-regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="api-rate-limit-server-url-rules-request-matcher-cookie-matchers-item-transformers"></a>&#x2022; [`transformers`](#api-rate-limit-server-url-rules-request-matcher-cookie-matchers-item-transformers) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### API Rate Limit Server URL Rules Request Matcher Headers

A [`headers`](#api-rate-limit-server-url-rules-request-matcher-headers) block (within [`api_rate_limit.server_url_rules.request_matcher`](#api-rate-limit-server-url-rules-request-matcher)) supports the following:

<a id="api-rate-limit-server-url-rules-request-matcher-headers-check-not-present"></a>&#x2022; [`check_not_present`](#api-rate-limit-server-url-rules-request-matcher-headers-check-not-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-rate-limit-server-url-rules-request-matcher-headers-check-present"></a>&#x2022; [`check_present`](#api-rate-limit-server-url-rules-request-matcher-headers-check-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-rate-limit-server-url-rules-request-matcher-headers-invert-matcher"></a>&#x2022; [`invert_matcher`](#api-rate-limit-server-url-rules-request-matcher-headers-invert-matcher) - Optional Bool<br>Invert Header Matcher. Invert the match result

<a id="api-rate-limit-server-url-rules-request-matcher-headers-item"></a>&#x2022; [`item`](#api-rate-limit-server-url-rules-request-matcher-headers-item) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#api-rate-limit-server-url-rules-request-matcher-headers-item) below.

<a id="api-rate-limit-server-url-rules-request-matcher-headers-name"></a>&#x2022; [`name`](#api-rate-limit-server-url-rules-request-matcher-headers-name) - Optional String<br>Header Name. A case-insensitive HTTP header name

#### API Rate Limit Server URL Rules Request Matcher Headers Item

An [`item`](#api-rate-limit-server-url-rules-request-matcher-headers-item) block (within [`api_rate_limit.server_url_rules.request_matcher.headers`](#api-rate-limit-server-url-rules-request-matcher-headers)) supports the following:

<a id="api-rate-limit-server-url-rules-request-matcher-headers-item-exact-values"></a>&#x2022; [`exact_values`](#api-rate-limit-server-url-rules-request-matcher-headers-item-exact-values) - Optional List<br>Exact Values. A list of exact values to match the input against

<a id="api-rate-limit-server-url-rules-request-matcher-headers-item-regex-values"></a>&#x2022; [`regex_values`](#api-rate-limit-server-url-rules-request-matcher-headers-item-regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="api-rate-limit-server-url-rules-request-matcher-headers-item-transformers"></a>&#x2022; [`transformers`](#api-rate-limit-server-url-rules-request-matcher-headers-item-transformers) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### API Rate Limit Server URL Rules Request Matcher JWT Claims

A [`jwt_claims`](#api-rate-limit-server-url-rules-request-matcher-jwt-claims) block (within [`api_rate_limit.server_url_rules.request_matcher`](#api-rate-limit-server-url-rules-request-matcher)) supports the following:

<a id="api-rate-limit-server-url-rules-request-matcher-jwt-claims-check-not-present"></a>&#x2022; [`check_not_present`](#api-rate-limit-server-url-rules-request-matcher-jwt-claims-check-not-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-rate-limit-server-url-rules-request-matcher-jwt-claims-check-present"></a>&#x2022; [`check_present`](#api-rate-limit-server-url-rules-request-matcher-jwt-claims-check-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-rate-limit-server-url-rules-request-matcher-jwt-claims-invert-matcher"></a>&#x2022; [`invert_matcher`](#api-rate-limit-server-url-rules-request-matcher-jwt-claims-invert-matcher) - Optional Bool<br>Invert Matcher. Invert the match result

<a id="api-rate-limit-server-url-rules-request-matcher-jwt-claims-item"></a>&#x2022; [`item`](#api-rate-limit-server-url-rules-request-matcher-jwt-claims-item) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#api-rate-limit-server-url-rules-request-matcher-jwt-claims-item) below.

<a id="api-rate-limit-server-url-rules-request-matcher-jwt-claims-name"></a>&#x2022; [`name`](#api-rate-limit-server-url-rules-request-matcher-jwt-claims-name) - Optional String<br>JWT Claim Name. JWT claim name

#### API Rate Limit Server URL Rules Request Matcher JWT Claims Item

An [`item`](#api-rate-limit-server-url-rules-request-matcher-jwt-claims-item) block (within [`api_rate_limit.server_url_rules.request_matcher.jwt_claims`](#api-rate-limit-server-url-rules-request-matcher-jwt-claims)) supports the following:

<a id="api-rate-limit-server-url-rules-request-matcher-jwt-claims-item-exact-values"></a>&#x2022; [`exact_values`](#api-rate-limit-server-url-rules-request-matcher-jwt-claims-item-exact-values) - Optional List<br>Exact Values. A list of exact values to match the input against

<a id="api-rate-limit-server-url-rules-request-matcher-jwt-claims-item-regex-values"></a>&#x2022; [`regex_values`](#api-rate-limit-server-url-rules-request-matcher-jwt-claims-item-regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="api-rate-limit-server-url-rules-request-matcher-jwt-claims-item-transformers"></a>&#x2022; [`transformers`](#api-rate-limit-server-url-rules-request-matcher-jwt-claims-item-transformers) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### API Rate Limit Server URL Rules Request Matcher Query Params

A [`query_params`](#api-rate-limit-server-url-rules-request-matcher-query-params) block (within [`api_rate_limit.server_url_rules.request_matcher`](#api-rate-limit-server-url-rules-request-matcher)) supports the following:

<a id="api-rate-limit-server-url-rules-request-matcher-query-params-check-not-present"></a>&#x2022; [`check_not_present`](#api-rate-limit-server-url-rules-request-matcher-query-params-check-not-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-rate-limit-server-url-rules-request-matcher-query-params-check-present"></a>&#x2022; [`check_present`](#api-rate-limit-server-url-rules-request-matcher-query-params-check-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-rate-limit-server-url-rules-request-matcher-query-params-invert-matcher"></a>&#x2022; [`invert_matcher`](#api-rate-limit-server-url-rules-request-matcher-query-params-invert-matcher) - Optional Bool<br>Invert Query Parameter Matcher. Invert the match result

<a id="api-rate-limit-server-url-rules-request-matcher-query-params-item"></a>&#x2022; [`item`](#api-rate-limit-server-url-rules-request-matcher-query-params-item) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#api-rate-limit-server-url-rules-request-matcher-query-params-item) below.

<a id="api-rate-limit-server-url-rules-request-matcher-query-params-key"></a>&#x2022; [`key`](#api-rate-limit-server-url-rules-request-matcher-query-params-key) - Optional String<br>Query Parameter Name. A case-sensitive HTTP query parameter name

#### API Rate Limit Server URL Rules Request Matcher Query Params Item

An [`item`](#api-rate-limit-server-url-rules-request-matcher-query-params-item) block (within [`api_rate_limit.server_url_rules.request_matcher.query_params`](#api-rate-limit-server-url-rules-request-matcher-query-params)) supports the following:

<a id="api-rate-limit-server-url-rules-request-matcher-query-params-item-exact-values"></a>&#x2022; [`exact_values`](#api-rate-limit-server-url-rules-request-matcher-query-params-item-exact-values) - Optional List<br>Exact Values. A list of exact values to match the input against

<a id="api-rate-limit-server-url-rules-request-matcher-query-params-item-regex-values"></a>&#x2022; [`regex_values`](#api-rate-limit-server-url-rules-request-matcher-query-params-item-regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="api-rate-limit-server-url-rules-request-matcher-query-params-item-transformers"></a>&#x2022; [`transformers`](#api-rate-limit-server-url-rules-request-matcher-query-params-item-transformers) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### API Specification

An [`api_specification`](#api-specification) block supports the following:

<a id="api-specification-api-definition"></a>&#x2022; [`api_definition`](#api-specification-api-definition) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [API Definition](#api-specification-api-definition) below.

<a id="api-specification-validation-all-spec-endpoints"></a>&#x2022; [`validation_all_spec_endpoints`](#api-specification-validation-all-spec-endpoints) - Optional Block<br>API Inventory. Settings for API Inventory validation<br>See [Validation All Spec Endpoints](#api-specification-validation-all-spec-endpoints) below.

<a id="api-specification-validation-custom-list"></a>&#x2022; [`validation_custom_list`](#api-specification-validation-custom-list) - Optional Block<br>Custom List. Define API groups, base paths, or API endpoints and their OpenAPI validation modes. Any other API-endpoint not listed will act according to 'Fall Through Mode'<br>See [Validation Custom List](#api-specification-validation-custom-list) below.

<a id="api-specification-validation-disabled"></a>&#x2022; [`validation_disabled`](#api-specification-validation-disabled) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### API Specification API Definition

An [`api_definition`](#api-specification-api-definition) block (within [`api_specification`](#api-specification)) supports the following:

<a id="api-specification-api-definition-name"></a>&#x2022; [`name`](#api-specification-api-definition-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="api-specification-api-definition-namespace"></a>&#x2022; [`namespace`](#api-specification-api-definition-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="api-specification-api-definition-tenant"></a>&#x2022; [`tenant`](#api-specification-api-definition-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### API Specification Validation All Spec Endpoints

A [`validation_all_spec_endpoints`](#api-specification-validation-all-spec-endpoints) block (within [`api_specification`](#api-specification)) supports the following:

<a id="api-specification-validation-all-spec-endpoints-fall-through-mode"></a>&#x2022; [`fall_through_mode`](#api-specification-validation-all-spec-endpoints-fall-through-mode) - Optional Block<br>Fall Through Mode. x-required Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a. swagger) or doesn't have a specific rule in custom rules)<br>See [Fall Through Mode](#api-specification-validation-all-spec-endpoints-fall-through-mode) below.

<a id="api-specification-validation-all-spec-endpoints-settings"></a>&#x2022; [`settings`](#api-specification-validation-all-spec-endpoints-settings) - Optional Block<br>Common Settings. OpenAPI specification validation settings relevant for 'API Inventory' enforcement and for 'Custom list' enforcement<br>See [Settings](#api-specification-validation-all-spec-endpoints-settings) below.

<a id="api-specification-validation-all-spec-endpoints-validation-mode"></a>&#x2022; [`validation_mode`](#api-specification-validation-all-spec-endpoints-validation-mode) - Optional Block<br>Validation Mode. x-required Validation mode of OpenAPI specification. When a validation mismatch occurs on a request to one of the endpoints listed on the OpenAPI specification file (a.k.a. swagger)<br>See [Validation Mode](#api-specification-validation-all-spec-endpoints-validation-mode) below.

#### API Specification Validation All Spec Endpoints Fall Through Mode

A [`fall_through_mode`](#api-specification-validation-all-spec-endpoints-fall-through-mode) block (within [`api_specification.validation_all_spec_endpoints`](#api-specification-validation-all-spec-endpoints)) supports the following:

<a id="api-specification-validation-all-spec-endpoints-fall-through-mode-fall-through-mode-allow"></a>&#x2022; [`fall_through_mode_allow`](#api-specification-validation-all-spec-endpoints-fall-through-mode-fall-through-mode-allow) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-specification-validation-all-spec-endpoints-fall-through-mode-fall-through-mode-custom"></a>&#x2022; [`fall_through_mode_custom`](#api-specification-validation-all-spec-endpoints-fall-through-mode-fall-through-mode-custom) - Optional Block<br>Custom Fall Through Mode. Define the fall through settings<br>See [Fall Through Mode Custom](#api-specification-validation-all-spec-endpoints-fall-through-mode-fall-through-mode-custom) below.

#### API Specification Validation All Spec Endpoints Fall Through Mode Fall Through Mode Custom

A [`fall_through_mode_custom`](#api-specification-validation-all-spec-endpoints-fall-through-mode-fall-through-mode-custom) block (within [`api_specification.validation_all_spec_endpoints.fall_through_mode`](#api-specification-validation-all-spec-endpoints-fall-through-mode)) supports the following:

<a id="api-specification-validation-all-spec-endpoints-fall-through-mode-fall-through-mode-custom-open-api-validation-rules"></a>&#x2022; [`open_api_validation_rules`](#api-specification-validation-all-spec-endpoints-fall-through-mode-fall-through-mode-custom-open-api-validation-rules) - Optional Block<br>Custom Fall Through Rule List<br>See [Open API Validation Rules](#api-specification-validation-all-spec-endpoints-fall-through-mode-fall-through-mode-custom-open-api-validation-rules) below.

#### API Specification Validation All Spec Endpoints Fall Through Mode Fall Through Mode Custom Open API Validation Rules

An [`open_api_validation_rules`](#api-specification-validation-all-spec-endpoints-fall-through-mode-fall-through-mode-custom-open-api-validation-rules) block (within [`api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom`](#api-specification-validation-all-spec-endpoints-fall-through-mode-fall-through-mode-custom)) supports the following:

<a id="api-specification-validation-all-spec-endpoints-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-action-block"></a>&#x2022; [`action_block`](#api-specification-validation-all-spec-endpoints-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-action-block) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-specification-validation-all-spec-endpoints-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-action-report"></a>&#x2022; [`action_report`](#api-specification-validation-all-spec-endpoints-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-action-report) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-specification-validation-all-spec-endpoints-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-action-skip"></a>&#x2022; [`action_skip`](#api-specification-validation-all-spec-endpoints-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-action-skip) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-specification-validation-all-spec-endpoints-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-api-endpoint"></a>&#x2022; [`api_endpoint`](#api-specification-validation-all-spec-endpoints-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-api-endpoint) - Optional Block<br>API Endpoint. This defines API endpoint<br>See [API Endpoint](#api-specification-validation-all-spec-endpoints-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-api-endpoint) below.

<a id="api-specification-validation-all-spec-endpoints-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-api-group"></a>&#x2022; [`api_group`](#api-specification-validation-all-spec-endpoints-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-api-group) - Optional String<br>API Group. The API group which this validation applies to

<a id="api-specification-validation-all-spec-endpoints-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-base-path"></a>&#x2022; [`base_path`](#api-specification-validation-all-spec-endpoints-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-base-path) - Optional String<br>Base Path. The base path which this validation applies to

<a id="api-specification-validation-all-spec-endpoints-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-metadata"></a>&#x2022; [`metadata`](#api-specification-validation-all-spec-endpoints-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-metadata) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#api-specification-validation-all-spec-endpoints-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-metadata) below.

#### API Specification Validation All Spec Endpoints Fall Through Mode Fall Through Mode Custom Open API Validation Rules API Endpoint

An [`api_endpoint`](#api-specification-validation-all-spec-endpoints-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-api-endpoint) block (within [`api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules`](#api-specification-validation-all-spec-endpoints-fall-through-mode-fall-through-mode-custom-open-api-validation-rules)) supports the following:

<a id="api-specification-validation-all-spec-endpoints-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-api-endpoint-methods"></a>&#x2022; [`methods`](#api-specification-validation-all-spec-endpoints-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-api-endpoint-methods) - Optional List  Defaults to `ANY`<br>Possible values are `ANY`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`, `COPY`<br>Methods. Methods to be matched

<a id="api-specification-validation-all-spec-endpoints-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-api-endpoint-path"></a>&#x2022; [`path`](#api-specification-validation-all-spec-endpoints-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-api-endpoint-path) - Optional String<br>Path. Path to be matched

#### API Specification Validation All Spec Endpoints Fall Through Mode Fall Through Mode Custom Open API Validation Rules Metadata

A [`metadata`](#api-specification-validation-all-spec-endpoints-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-metadata) block (within [`api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules`](#api-specification-validation-all-spec-endpoints-fall-through-mode-fall-through-mode-custom-open-api-validation-rules)) supports the following:

<a id="api-specification-validation-all-spec-endpoints-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-metadata-description-spec"></a>&#x2022; [`description_spec`](#api-specification-validation-all-spec-endpoints-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-metadata-description-spec) - Optional String<br>Description. Human readable description

<a id="api-specification-validation-all-spec-endpoints-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-metadata-name"></a>&#x2022; [`name`](#api-specification-validation-all-spec-endpoints-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-metadata-name) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### API Specification Validation All Spec Endpoints Settings

A [`settings`](#api-specification-validation-all-spec-endpoints-settings) block (within [`api_specification.validation_all_spec_endpoints`](#api-specification-validation-all-spec-endpoints)) supports the following:

<a id="api-specification-validation-all-spec-endpoints-settings-oversized-body-fail-validation"></a>&#x2022; [`oversized_body_fail_validation`](#api-specification-validation-all-spec-endpoints-settings-oversized-body-fail-validation) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-specification-validation-all-spec-endpoints-settings-oversized-body-skip-validation"></a>&#x2022; [`oversized_body_skip_validation`](#api-specification-validation-all-spec-endpoints-settings-oversized-body-skip-validation) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-specification-validation-all-spec-endpoints-settings-property-validation-settings-custom"></a>&#x2022; [`property_validation_settings_custom`](#api-specification-validation-all-spec-endpoints-settings-property-validation-settings-custom) - Optional Block<br>Validation Property Settings. Custom property validation settings<br>See [Property Validation Settings Custom](#api-specification-validation-all-spec-endpoints-settings-property-validation-settings-custom) below.

<a id="api-specification-validation-all-spec-endpoints-settings-property-validation-settings-default"></a>&#x2022; [`property_validation_settings_default`](#api-specification-validation-all-spec-endpoints-settings-property-validation-settings-default) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### API Specification Validation All Spec Endpoints Settings Property Validation Settings Custom

A [`property_validation_settings_custom`](#api-specification-validation-all-spec-endpoints-settings-property-validation-settings-custom) block (within [`api_specification.validation_all_spec_endpoints.settings`](#api-specification-validation-all-spec-endpoints-settings)) supports the following:

<a id="api-specification-validation-all-spec-endpoints-settings-property-validation-settings-custom-query-parameters"></a>&#x2022; [`query_parameters`](#api-specification-validation-all-spec-endpoints-settings-property-validation-settings-custom-query-parameters) - Optional Block<br>Validation Settings For Query Parameters. Custom settings for query parameters validation<br>See [Query Parameters](#api-specification-validation-all-spec-endpoints-settings-property-validation-settings-custom-query-parameters) below.

#### API Specification Validation All Spec Endpoints Settings Property Validation Settings Custom Query Parameters

A [`query_parameters`](#api-specification-validation-all-spec-endpoints-settings-property-validation-settings-custom-query-parameters) block (within [`api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom`](#api-specification-validation-all-spec-endpoints-settings-property-validation-settings-custom)) supports the following:

<a id="api-specification-validation-all-spec-endpoints-settings-property-validation-settings-custom-query-parameters-allow-additional-parameters"></a>&#x2022; [`allow_additional_parameters`](#api-specification-validation-all-spec-endpoints-settings-property-validation-settings-custom-query-parameters-allow-additional-parameters) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-specification-validation-all-spec-endpoints-settings-property-validation-settings-custom-query-parameters-disallow-additional-parameters"></a>&#x2022; [`disallow_additional_parameters`](#api-specification-validation-all-spec-endpoints-settings-property-validation-settings-custom-query-parameters-disallow-additional-parameters) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### API Specification Validation All Spec Endpoints Validation Mode

A [`validation_mode`](#api-specification-validation-all-spec-endpoints-validation-mode) block (within [`api_specification.validation_all_spec_endpoints`](#api-specification-validation-all-spec-endpoints)) supports the following:

<a id="api-specification-validation-all-spec-endpoints-validation-mode-response-validation-mode-active"></a>&#x2022; [`response_validation_mode_active`](#api-specification-validation-all-spec-endpoints-validation-mode-response-validation-mode-active) - Optional Block<br>Open API Validation Mode Active. Validation mode properties of response<br>See [Response Validation Mode Active](#api-specification-validation-all-spec-endpoints-validation-mode-response-validation-mode-active) below.

<a id="api-specification-validation-all-spec-endpoints-validation-mode-skip-response-validation"></a>&#x2022; [`skip_response_validation`](#api-specification-validation-all-spec-endpoints-validation-mode-skip-response-validation) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-specification-validation-all-spec-endpoints-validation-mode-skip-validation"></a>&#x2022; [`skip_validation`](#api-specification-validation-all-spec-endpoints-validation-mode-skip-validation) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-specification-validation-all-spec-endpoints-validation-mode-validation-mode-active"></a>&#x2022; [`validation_mode_active`](#api-specification-validation-all-spec-endpoints-validation-mode-validation-mode-active) - Optional Block<br>Open API Validation Mode Active. Validation mode properties of request<br>See [Validation Mode Active](#api-specification-validation-all-spec-endpoints-validation-mode-validation-mode-active) below.

#### API Specification Validation All Spec Endpoints Validation Mode Response Validation Mode Active

A [`response_validation_mode_active`](#api-specification-validation-all-spec-endpoints-validation-mode-response-validation-mode-active) block (within [`api_specification.validation_all_spec_endpoints.validation_mode`](#api-specification-validation-all-spec-endpoints-validation-mode)) supports the following:

<a id="api-specification-validation-all-spec-endpoints-validation-mode-response-validation-mode-active-enforcement-block"></a>&#x2022; [`enforcement_block`](#api-specification-validation-all-spec-endpoints-validation-mode-response-validation-mode-active-enforcement-block) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-specification-validation-all-spec-endpoints-validation-mode-response-validation-mode-active-enforcement-report"></a>&#x2022; [`enforcement_report`](#api-specification-validation-all-spec-endpoints-validation-mode-response-validation-mode-active-enforcement-report) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-specification-validation-all-spec-endpoints-validation-mode-response-validation-mode-active-response-validation-properties"></a>&#x2022; [`response_validation_properties`](#api-specification-validation-all-spec-endpoints-validation-mode-response-validation-mode-active-response-validation-properties) - Optional List  Defaults to `PROPERTY_QUERY_PARAMETERS`<br>Possible values are `PROPERTY_QUERY_PARAMETERS`, `PROPERTY_PATH_PARAMETERS`, `PROPERTY_CONTENT_TYPE`, `PROPERTY_COOKIE_PARAMETERS`, `PROPERTY_HTTP_HEADERS`, `PROPERTY_HTTP_BODY`, `PROPERTY_SECURITY_SCHEMA`, `PROPERTY_RESPONSE_CODE`<br>Response Validation Properties. List of properties of the response to validate according to the OpenAPI specification file (a.k.a. swagger)

#### API Specification Validation All Spec Endpoints Validation Mode Validation Mode Active

A [`validation_mode_active`](#api-specification-validation-all-spec-endpoints-validation-mode-validation-mode-active) block (within [`api_specification.validation_all_spec_endpoints.validation_mode`](#api-specification-validation-all-spec-endpoints-validation-mode)) supports the following:

<a id="api-specification-validation-all-spec-endpoints-validation-mode-validation-mode-active-enforcement-block"></a>&#x2022; [`enforcement_block`](#api-specification-validation-all-spec-endpoints-validation-mode-validation-mode-active-enforcement-block) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-specification-validation-all-spec-endpoints-validation-mode-validation-mode-active-enforcement-report"></a>&#x2022; [`enforcement_report`](#api-specification-validation-all-spec-endpoints-validation-mode-validation-mode-active-enforcement-report) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-specification-validation-all-spec-endpoints-validation-mode-validation-mode-active-request-validation-properties"></a>&#x2022; [`request_validation_properties`](#api-specification-validation-all-spec-endpoints-validation-mode-validation-mode-active-request-validation-properties) - Optional List  Defaults to `PROPERTY_QUERY_PARAMETERS`<br>Possible values are `PROPERTY_QUERY_PARAMETERS`, `PROPERTY_PATH_PARAMETERS`, `PROPERTY_CONTENT_TYPE`, `PROPERTY_COOKIE_PARAMETERS`, `PROPERTY_HTTP_HEADERS`, `PROPERTY_HTTP_BODY`, `PROPERTY_SECURITY_SCHEMA`, `PROPERTY_RESPONSE_CODE`<br>Request Validation Properties. List of properties of the request to validate according to the OpenAPI specification file (a.k.a. swagger)

#### API Specification Validation Custom List

A [`validation_custom_list`](#api-specification-validation-custom-list) block (within [`api_specification`](#api-specification)) supports the following:

<a id="api-specification-validation-custom-list-fall-through-mode"></a>&#x2022; [`fall_through_mode`](#api-specification-validation-custom-list-fall-through-mode) - Optional Block<br>Fall Through Mode. x-required Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a. swagger) or doesn't have a specific rule in custom rules)<br>See [Fall Through Mode](#api-specification-validation-custom-list-fall-through-mode) below.

<a id="api-specification-validation-custom-list-open-api-validation-rules"></a>&#x2022; [`open_api_validation_rules`](#api-specification-validation-custom-list-open-api-validation-rules) - Optional Block<br>Validation List<br>See [Open API Validation Rules](#api-specification-validation-custom-list-open-api-validation-rules) below.

<a id="api-specification-validation-custom-list-settings"></a>&#x2022; [`settings`](#api-specification-validation-custom-list-settings) - Optional Block<br>Common Settings. OpenAPI specification validation settings relevant for 'API Inventory' enforcement and for 'Custom list' enforcement<br>See [Settings](#api-specification-validation-custom-list-settings) below.

#### API Specification Validation Custom List Fall Through Mode

A [`fall_through_mode`](#api-specification-validation-custom-list-fall-through-mode) block (within [`api_specification.validation_custom_list`](#api-specification-validation-custom-list)) supports the following:

<a id="api-specification-validation-custom-list-fall-through-mode-fall-through-mode-allow"></a>&#x2022; [`fall_through_mode_allow`](#api-specification-validation-custom-list-fall-through-mode-fall-through-mode-allow) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-specification-validation-custom-list-fall-through-mode-fall-through-mode-custom"></a>&#x2022; [`fall_through_mode_custom`](#api-specification-validation-custom-list-fall-through-mode-fall-through-mode-custom) - Optional Block<br>Custom Fall Through Mode. Define the fall through settings<br>See [Fall Through Mode Custom](#api-specification-validation-custom-list-fall-through-mode-fall-through-mode-custom) below.

#### API Specification Validation Custom List Fall Through Mode Fall Through Mode Custom

A [`fall_through_mode_custom`](#api-specification-validation-custom-list-fall-through-mode-fall-through-mode-custom) block (within [`api_specification.validation_custom_list.fall_through_mode`](#api-specification-validation-custom-list-fall-through-mode)) supports the following:

<a id="api-specification-validation-custom-list-fall-through-mode-fall-through-mode-custom-open-api-validation-rules"></a>&#x2022; [`open_api_validation_rules`](#api-specification-validation-custom-list-fall-through-mode-fall-through-mode-custom-open-api-validation-rules) - Optional Block<br>Custom Fall Through Rule List<br>See [Open API Validation Rules](#api-specification-validation-custom-list-fall-through-mode-fall-through-mode-custom-open-api-validation-rules) below.

#### API Specification Validation Custom List Fall Through Mode Fall Through Mode Custom Open API Validation Rules

An [`open_api_validation_rules`](#api-specification-validation-custom-list-fall-through-mode-fall-through-mode-custom-open-api-validation-rules) block (within [`api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom`](#api-specification-validation-custom-list-fall-through-mode-fall-through-mode-custom)) supports the following:

<a id="api-specification-validation-custom-list-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-action-block"></a>&#x2022; [`action_block`](#api-specification-validation-custom-list-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-action-block) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-specification-validation-custom-list-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-action-report"></a>&#x2022; [`action_report`](#api-specification-validation-custom-list-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-action-report) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-specification-validation-custom-list-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-action-skip"></a>&#x2022; [`action_skip`](#api-specification-validation-custom-list-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-action-skip) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-specification-validation-custom-list-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-api-endpoint"></a>&#x2022; [`api_endpoint`](#api-specification-validation-custom-list-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-api-endpoint) - Optional Block<br>API Endpoint. This defines API endpoint<br>See [API Endpoint](#api-specification-validation-custom-list-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-api-endpoint) below.

<a id="api-specification-validation-custom-list-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-api-group"></a>&#x2022; [`api_group`](#api-specification-validation-custom-list-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-api-group) - Optional String<br>API Group. The API group which this validation applies to

<a id="api-specification-validation-custom-list-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-base-path"></a>&#x2022; [`base_path`](#api-specification-validation-custom-list-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-base-path) - Optional String<br>Base Path. The base path which this validation applies to

<a id="api-specification-validation-custom-list-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-metadata"></a>&#x2022; [`metadata`](#api-specification-validation-custom-list-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-metadata) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#api-specification-validation-custom-list-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-metadata) below.

#### API Specification Validation Custom List Fall Through Mode Fall Through Mode Custom Open API Validation Rules API Endpoint

An [`api_endpoint`](#api-specification-validation-custom-list-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-api-endpoint) block (within [`api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules`](#api-specification-validation-custom-list-fall-through-mode-fall-through-mode-custom-open-api-validation-rules)) supports the following:

<a id="api-specification-validation-custom-list-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-api-endpoint-methods"></a>&#x2022; [`methods`](#api-specification-validation-custom-list-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-api-endpoint-methods) - Optional List  Defaults to `ANY`<br>Possible values are `ANY`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`, `COPY`<br>Methods. Methods to be matched

<a id="api-specification-validation-custom-list-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-api-endpoint-path"></a>&#x2022; [`path`](#api-specification-validation-custom-list-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-api-endpoint-path) - Optional String<br>Path. Path to be matched

#### API Specification Validation Custom List Fall Through Mode Fall Through Mode Custom Open API Validation Rules Metadata

A [`metadata`](#api-specification-validation-custom-list-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-metadata) block (within [`api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules`](#api-specification-validation-custom-list-fall-through-mode-fall-through-mode-custom-open-api-validation-rules)) supports the following:

<a id="api-specification-validation-custom-list-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-metadata-description-spec"></a>&#x2022; [`description_spec`](#api-specification-validation-custom-list-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-metadata-description-spec) - Optional String<br>Description. Human readable description

<a id="api-specification-validation-custom-list-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-metadata-name"></a>&#x2022; [`name`](#api-specification-validation-custom-list-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-metadata-name) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### API Specification Validation Custom List Open API Validation Rules

An [`open_api_validation_rules`](#api-specification-validation-custom-list-open-api-validation-rules) block (within [`api_specification.validation_custom_list`](#api-specification-validation-custom-list)) supports the following:

<a id="api-specification-validation-custom-list-open-api-validation-rules-any-domain"></a>&#x2022; [`any_domain`](#api-specification-validation-custom-list-open-api-validation-rules-any-domain) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-specification-validation-custom-list-open-api-validation-rules-api-endpoint"></a>&#x2022; [`api_endpoint`](#api-specification-validation-custom-list-open-api-validation-rules-api-endpoint) - Optional Block<br>API Endpoint. This defines API endpoint<br>See [API Endpoint](#api-specification-validation-custom-list-open-api-validation-rules-api-endpoint) below.

<a id="api-specification-validation-custom-list-open-api-validation-rules-api-group"></a>&#x2022; [`api_group`](#api-specification-validation-custom-list-open-api-validation-rules-api-group) - Optional String<br>API Group. The API group which this validation applies to

<a id="api-specification-validation-custom-list-open-api-validation-rules-base-path"></a>&#x2022; [`base_path`](#api-specification-validation-custom-list-open-api-validation-rules-base-path) - Optional String<br>Base Path. The base path which this validation applies to

<a id="api-specification-validation-custom-list-open-api-validation-rules-metadata"></a>&#x2022; [`metadata`](#api-specification-validation-custom-list-open-api-validation-rules-metadata) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#api-specification-validation-custom-list-open-api-validation-rules-metadata) below.

<a id="api-specification-validation-custom-list-open-api-validation-rules-specific-domain"></a>&#x2022; [`specific_domain`](#api-specification-validation-custom-list-open-api-validation-rules-specific-domain) - Optional String<br>Specific Domain. The rule will apply for a specific domain

<a id="api-specification-validation-custom-list-open-api-validation-rules-validation-mode"></a>&#x2022; [`validation_mode`](#api-specification-validation-custom-list-open-api-validation-rules-validation-mode) - Optional Block<br>Validation Mode. x-required Validation mode of OpenAPI specification. When a validation mismatch occurs on a request to one of the endpoints listed on the OpenAPI specification file (a.k.a. swagger)<br>See [Validation Mode](#api-specification-validation-custom-list-open-api-validation-rules-validation-mode) below.

#### API Specification Validation Custom List Open API Validation Rules API Endpoint

An [`api_endpoint`](#api-specification-validation-custom-list-open-api-validation-rules-api-endpoint) block (within [`api_specification.validation_custom_list.open_api_validation_rules`](#api-specification-validation-custom-list-open-api-validation-rules)) supports the following:

<a id="api-specification-validation-custom-list-open-api-validation-rules-api-endpoint-methods"></a>&#x2022; [`methods`](#api-specification-validation-custom-list-open-api-validation-rules-api-endpoint-methods) - Optional List  Defaults to `ANY`<br>Possible values are `ANY`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`, `COPY`<br>Methods. Methods to be matched

<a id="api-specification-validation-custom-list-open-api-validation-rules-api-endpoint-path"></a>&#x2022; [`path`](#api-specification-validation-custom-list-open-api-validation-rules-api-endpoint-path) - Optional String<br>Path. Path to be matched

#### API Specification Validation Custom List Open API Validation Rules Metadata

A [`metadata`](#api-specification-validation-custom-list-open-api-validation-rules-metadata) block (within [`api_specification.validation_custom_list.open_api_validation_rules`](#api-specification-validation-custom-list-open-api-validation-rules)) supports the following:

<a id="api-specification-validation-custom-list-open-api-validation-rules-metadata-description-spec"></a>&#x2022; [`description_spec`](#api-specification-validation-custom-list-open-api-validation-rules-metadata-description-spec) - Optional String<br>Description. Human readable description

<a id="api-specification-validation-custom-list-open-api-validation-rules-metadata-name"></a>&#x2022; [`name`](#api-specification-validation-custom-list-open-api-validation-rules-metadata-name) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### API Specification Validation Custom List Open API Validation Rules Validation Mode

A [`validation_mode`](#api-specification-validation-custom-list-open-api-validation-rules-validation-mode) block (within [`api_specification.validation_custom_list.open_api_validation_rules`](#api-specification-validation-custom-list-open-api-validation-rules)) supports the following:

<a id="api-specification-validation-custom-list-open-api-validation-rules-validation-mode-response-validation-mode-active"></a>&#x2022; [`response_validation_mode_active`](#api-specification-validation-custom-list-open-api-validation-rules-validation-mode-response-validation-mode-active) - Optional Block<br>Open API Validation Mode Active. Validation mode properties of response<br>See [Response Validation Mode Active](#api-specification-validation-custom-list-open-api-validation-rules-validation-mode-response-validation-mode-active) below.

<a id="api-specification-validation-custom-list-open-api-validation-rules-validation-mode-skip-response-validation"></a>&#x2022; [`skip_response_validation`](#api-specification-validation-custom-list-open-api-validation-rules-validation-mode-skip-response-validation) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-specification-validation-custom-list-open-api-validation-rules-validation-mode-skip-validation"></a>&#x2022; [`skip_validation`](#api-specification-validation-custom-list-open-api-validation-rules-validation-mode-skip-validation) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-specification-validation-custom-list-open-api-validation-rules-validation-mode-validation-mode-active"></a>&#x2022; [`validation_mode_active`](#api-specification-validation-custom-list-open-api-validation-rules-validation-mode-validation-mode-active) - Optional Block<br>Open API Validation Mode Active. Validation mode properties of request<br>See [Validation Mode Active](#api-specification-validation-custom-list-open-api-validation-rules-validation-mode-validation-mode-active) below.

#### API Specification Validation Custom List Open API Validation Rules Validation Mode Response Validation Mode Active

A [`response_validation_mode_active`](#api-specification-validation-custom-list-open-api-validation-rules-validation-mode-response-validation-mode-active) block (within [`api_specification.validation_custom_list.open_api_validation_rules.validation_mode`](#api-specification-validation-custom-list-open-api-validation-rules-validation-mode)) supports the following:

<a id="api-specification-validation-custom-list-open-api-validation-rules-validation-mode-response-validation-mode-active-enforcement-block"></a>&#x2022; [`enforcement_block`](#api-specification-validation-custom-list-open-api-validation-rules-validation-mode-response-validation-mode-active-enforcement-block) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-specification-validation-custom-list-open-api-validation-rules-validation-mode-response-validation-mode-active-enforcement-report"></a>&#x2022; [`enforcement_report`](#api-specification-validation-custom-list-open-api-validation-rules-validation-mode-response-validation-mode-active-enforcement-report) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-specification-validation-custom-list-open-api-validation-rules-validation-mode-response-validation-mode-active-response-validation-properties"></a>&#x2022; [`response_validation_properties`](#api-specification-validation-custom-list-open-api-validation-rules-validation-mode-response-validation-mode-active-response-validation-properties) - Optional List  Defaults to `PROPERTY_QUERY_PARAMETERS`<br>Possible values are `PROPERTY_QUERY_PARAMETERS`, `PROPERTY_PATH_PARAMETERS`, `PROPERTY_CONTENT_TYPE`, `PROPERTY_COOKIE_PARAMETERS`, `PROPERTY_HTTP_HEADERS`, `PROPERTY_HTTP_BODY`, `PROPERTY_SECURITY_SCHEMA`, `PROPERTY_RESPONSE_CODE`<br>Response Validation Properties. List of properties of the response to validate according to the OpenAPI specification file (a.k.a. swagger)

#### API Specification Validation Custom List Open API Validation Rules Validation Mode Validation Mode Active

A [`validation_mode_active`](#api-specification-validation-custom-list-open-api-validation-rules-validation-mode-validation-mode-active) block (within [`api_specification.validation_custom_list.open_api_validation_rules.validation_mode`](#api-specification-validation-custom-list-open-api-validation-rules-validation-mode)) supports the following:

<a id="api-specification-validation-custom-list-open-api-validation-rules-validation-mode-validation-mode-active-enforcement-block"></a>&#x2022; [`enforcement_block`](#api-specification-validation-custom-list-open-api-validation-rules-validation-mode-validation-mode-active-enforcement-block) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-specification-validation-custom-list-open-api-validation-rules-validation-mode-validation-mode-active-enforcement-report"></a>&#x2022; [`enforcement_report`](#api-specification-validation-custom-list-open-api-validation-rules-validation-mode-validation-mode-active-enforcement-report) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-specification-validation-custom-list-open-api-validation-rules-validation-mode-validation-mode-active-request-validation-properties"></a>&#x2022; [`request_validation_properties`](#api-specification-validation-custom-list-open-api-validation-rules-validation-mode-validation-mode-active-request-validation-properties) - Optional List  Defaults to `PROPERTY_QUERY_PARAMETERS`<br>Possible values are `PROPERTY_QUERY_PARAMETERS`, `PROPERTY_PATH_PARAMETERS`, `PROPERTY_CONTENT_TYPE`, `PROPERTY_COOKIE_PARAMETERS`, `PROPERTY_HTTP_HEADERS`, `PROPERTY_HTTP_BODY`, `PROPERTY_SECURITY_SCHEMA`, `PROPERTY_RESPONSE_CODE`<br>Request Validation Properties. List of properties of the request to validate according to the OpenAPI specification file (a.k.a. swagger)

#### API Specification Validation Custom List Settings

A [`settings`](#api-specification-validation-custom-list-settings) block (within [`api_specification.validation_custom_list`](#api-specification-validation-custom-list)) supports the following:

<a id="api-specification-validation-custom-list-settings-oversized-body-fail-validation"></a>&#x2022; [`oversized_body_fail_validation`](#api-specification-validation-custom-list-settings-oversized-body-fail-validation) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-specification-validation-custom-list-settings-oversized-body-skip-validation"></a>&#x2022; [`oversized_body_skip_validation`](#api-specification-validation-custom-list-settings-oversized-body-skip-validation) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-specification-validation-custom-list-settings-property-validation-settings-custom"></a>&#x2022; [`property_validation_settings_custom`](#api-specification-validation-custom-list-settings-property-validation-settings-custom) - Optional Block<br>Validation Property Settings. Custom property validation settings<br>See [Property Validation Settings Custom](#api-specification-validation-custom-list-settings-property-validation-settings-custom) below.

<a id="api-specification-validation-custom-list-settings-property-validation-settings-default"></a>&#x2022; [`property_validation_settings_default`](#api-specification-validation-custom-list-settings-property-validation-settings-default) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### API Specification Validation Custom List Settings Property Validation Settings Custom

A [`property_validation_settings_custom`](#api-specification-validation-custom-list-settings-property-validation-settings-custom) block (within [`api_specification.validation_custom_list.settings`](#api-specification-validation-custom-list-settings)) supports the following:

<a id="api-specification-validation-custom-list-settings-property-validation-settings-custom-query-parameters"></a>&#x2022; [`query_parameters`](#api-specification-validation-custom-list-settings-property-validation-settings-custom-query-parameters) - Optional Block<br>Validation Settings For Query Parameters. Custom settings for query parameters validation<br>See [Query Parameters](#api-specification-validation-custom-list-settings-property-validation-settings-custom-query-parameters) below.

#### API Specification Validation Custom List Settings Property Validation Settings Custom Query Parameters

A [`query_parameters`](#api-specification-validation-custom-list-settings-property-validation-settings-custom-query-parameters) block (within [`api_specification.validation_custom_list.settings.property_validation_settings_custom`](#api-specification-validation-custom-list-settings-property-validation-settings-custom)) supports the following:

<a id="api-specification-validation-custom-list-settings-property-validation-settings-custom-query-parameters-allow-additional-parameters"></a>&#x2022; [`allow_additional_parameters`](#api-specification-validation-custom-list-settings-property-validation-settings-custom-query-parameters-allow-additional-parameters) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-specification-validation-custom-list-settings-property-validation-settings-custom-query-parameters-disallow-additional-parameters"></a>&#x2022; [`disallow_additional_parameters`](#api-specification-validation-custom-list-settings-property-validation-settings-custom-query-parameters-disallow-additional-parameters) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### API Testing

An [`api_testing`](#api-testing) block supports the following:

<a id="api-testing-custom-header-value"></a>&#x2022; [`custom_header_value`](#api-testing-custom-header-value) - Optional String<br>Custom Header. Add x-f5-API-testing-identifier header value to prevent security flags on API testing traffic

<a id="api-testing-domains"></a>&#x2022; [`domains`](#api-testing-domains) - Optional Block<br>Testing Environments. Add and configure testing domains and credentials<br>See [Domains](#api-testing-domains) below.

<a id="api-testing-every-day"></a>&#x2022; [`every_day`](#api-testing-every-day) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-testing-every-month"></a>&#x2022; [`every_month`](#api-testing-every-month) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-testing-every-week"></a>&#x2022; [`every_week`](#api-testing-every-week) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### API Testing Domains

A [`domains`](#api-testing-domains) block (within [`api_testing`](#api-testing)) supports the following:

<a id="api-testing-domains-allow-destructive-methods"></a>&#x2022; [`allow_destructive_methods`](#api-testing-domains-allow-destructive-methods) - Optional Bool<br>Use Destructive Methods (e.g., DELETE, PUT). Enable to allow API test to execute destructive methods. Be cautious as these can alter or delete data

<a id="api-testing-domains-credentials"></a>&#x2022; [`credentials`](#api-testing-domains-credentials) - Optional Block<br>Credentials. Add credentials for API testing to use in the selected environment<br>See [Credentials](#api-testing-domains-credentials) below.

<a id="api-testing-domains-domain"></a>&#x2022; [`domain`](#api-testing-domains-domain) - Optional String<br>Domain. Add your testing environment domain. Be aware that running tests on a production domain can impact live applications, as API testing cannot distinguish between production and testing environments

#### API Testing Domains Credentials

A [`credentials`](#api-testing-domains-credentials) block (within [`api_testing.domains`](#api-testing-domains)) supports the following:

<a id="api-testing-domains-credentials-admin"></a>&#x2022; [`admin`](#api-testing-domains-credentials-admin) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-testing-domains-credentials-api-key"></a>&#x2022; [`api_key`](#api-testing-domains-credentials-api-key) - Optional Block<br>API Key<br>See [API Key](#api-testing-domains-credentials-api-key) below.

<a id="api-testing-domains-credentials-basic-auth"></a>&#x2022; [`basic_auth`](#api-testing-domains-credentials-basic-auth) - Optional Block<br>Basic Authentication<br>See [Basic Auth](#api-testing-domains-credentials-basic-auth) below.

<a id="api-testing-domains-credentials-bearer-token"></a>&#x2022; [`bearer_token`](#api-testing-domains-credentials-bearer-token) - Optional Block<br>Bearer<br>See [Bearer Token](#api-testing-domains-credentials-bearer-token) below.

<a id="api-testing-domains-credentials-credential-name"></a>&#x2022; [`credential_name`](#api-testing-domains-credentials-credential-name) - Optional String<br>Name. Enter a unique name for the credentials used in API testing

<a id="api-testing-domains-credentials-login-endpoint"></a>&#x2022; [`login_endpoint`](#api-testing-domains-credentials-login-endpoint) - Optional Block<br>Login Endpoint<br>See [Login Endpoint](#api-testing-domains-credentials-login-endpoint) below.

<a id="api-testing-domains-credentials-standard"></a>&#x2022; [`standard`](#api-testing-domains-credentials-standard) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### API Testing Domains Credentials API Key

An [`api_key`](#api-testing-domains-credentials-api-key) block (within [`api_testing.domains.credentials`](#api-testing-domains-credentials)) supports the following:

<a id="api-testing-domains-credentials-api-key-key"></a>&#x2022; [`key`](#api-testing-domains-credentials-api-key-key) - Optional String<br>Key

<a id="api-testing-domains-credentials-api-key-value"></a>&#x2022; [`value`](#api-testing-domains-credentials-api-key-value) - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Value](#api-testing-domains-credentials-api-key-value) below.

#### API Testing Domains Credentials API Key Value

A [`value`](#api-testing-domains-credentials-api-key-value) block (within [`api_testing.domains.credentials.api_key`](#api-testing-domains-credentials-api-key)) supports the following:

<a id="api-testing-domains-credentials-api-key-value-blindfold-secret-info"></a>&#x2022; [`blindfold_secret_info`](#api-testing-domains-credentials-api-key-value-blindfold-secret-info) - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#api-testing-domains-credentials-api-key-value-blindfold-secret-info) below.

<a id="api-testing-domains-credentials-api-key-value-clear-secret-info"></a>&#x2022; [`clear_secret_info`](#api-testing-domains-credentials-api-key-value-clear-secret-info) - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#api-testing-domains-credentials-api-key-value-clear-secret-info) below.

#### API Testing Domains Credentials API Key Value Blindfold Secret Info

A [`blindfold_secret_info`](#api-testing-domains-credentials-api-key-value-blindfold-secret-info) block (within [`api_testing.domains.credentials.api_key.value`](#api-testing-domains-credentials-api-key-value)) supports the following:

<a id="api-testing-domains-credentials-api-key-value-blindfold-secret-info-decryption-provider"></a>&#x2022; [`decryption_provider`](#api-testing-domains-credentials-api-key-value-blindfold-secret-info-decryption-provider) - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

<a id="api-testing-domains-credentials-api-key-value-blindfold-secret-info-location"></a>&#x2022; [`location`](#api-testing-domains-credentials-api-key-value-blindfold-secret-info-location) - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

<a id="api-testing-domains-credentials-api-key-value-blindfold-secret-info-store-provider"></a>&#x2022; [`store_provider`](#api-testing-domains-credentials-api-key-value-blindfold-secret-info-store-provider) - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

#### API Testing Domains Credentials API Key Value Clear Secret Info

A [`clear_secret_info`](#api-testing-domains-credentials-api-key-value-clear-secret-info) block (within [`api_testing.domains.credentials.api_key.value`](#api-testing-domains-credentials-api-key-value)) supports the following:

<a id="api-testing-domains-credentials-api-key-value-clear-secret-info-provider-ref"></a>&#x2022; [`provider_ref`](#api-testing-domains-credentials-api-key-value-clear-secret-info-provider-ref) - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

<a id="api-testing-domains-credentials-api-key-value-clear-secret-info-url"></a>&#x2022; [`url`](#api-testing-domains-credentials-api-key-value-clear-secret-info-url) - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

#### API Testing Domains Credentials Basic Auth

A [`basic_auth`](#api-testing-domains-credentials-basic-auth) block (within [`api_testing.domains.credentials`](#api-testing-domains-credentials)) supports the following:

<a id="api-testing-domains-credentials-basic-auth-password"></a>&#x2022; [`password`](#api-testing-domains-credentials-basic-auth-password) - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Password](#api-testing-domains-credentials-basic-auth-password) below.

<a id="api-testing-domains-credentials-basic-auth-user"></a>&#x2022; [`user`](#api-testing-domains-credentials-basic-auth-user) - Optional String<br>User

#### API Testing Domains Credentials Basic Auth Password

A [`password`](#api-testing-domains-credentials-basic-auth-password) block (within [`api_testing.domains.credentials.basic_auth`](#api-testing-domains-credentials-basic-auth)) supports the following:

<a id="api-testing-domains-credentials-basic-auth-password-blindfold-secret-info"></a>&#x2022; [`blindfold_secret_info`](#api-testing-domains-credentials-basic-auth-password-blindfold-secret-info) - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#api-testing-domains-credentials-basic-auth-password-blindfold-secret-info) below.

<a id="api-testing-domains-credentials-basic-auth-password-clear-secret-info"></a>&#x2022; [`clear_secret_info`](#api-testing-domains-credentials-basic-auth-password-clear-secret-info) - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#api-testing-domains-credentials-basic-auth-password-clear-secret-info) below.

#### API Testing Domains Credentials Basic Auth Password Blindfold Secret Info

A [`blindfold_secret_info`](#api-testing-domains-credentials-basic-auth-password-blindfold-secret-info) block (within [`api_testing.domains.credentials.basic_auth.password`](#api-testing-domains-credentials-basic-auth-password)) supports the following:

<a id="api-testing-domains-credentials-basic-auth-password-blindfold-secret-info-decryption-provider"></a>&#x2022; [`decryption_provider`](#api-testing-domains-credentials-basic-auth-password-blindfold-secret-info-decryption-provider) - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

<a id="api-testing-domains-credentials-basic-auth-password-blindfold-secret-info-location"></a>&#x2022; [`location`](#api-testing-domains-credentials-basic-auth-password-blindfold-secret-info-location) - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

<a id="api-testing-domains-credentials-basic-auth-password-blindfold-secret-info-store-provider"></a>&#x2022; [`store_provider`](#api-testing-domains-credentials-basic-auth-password-blindfold-secret-info-store-provider) - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

#### API Testing Domains Credentials Basic Auth Password Clear Secret Info

A [`clear_secret_info`](#api-testing-domains-credentials-basic-auth-password-clear-secret-info) block (within [`api_testing.domains.credentials.basic_auth.password`](#api-testing-domains-credentials-basic-auth-password)) supports the following:

<a id="api-testing-domains-credentials-basic-auth-password-clear-secret-info-provider-ref"></a>&#x2022; [`provider_ref`](#api-testing-domains-credentials-basic-auth-password-clear-secret-info-provider-ref) - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

<a id="api-testing-domains-credentials-basic-auth-password-clear-secret-info-url"></a>&#x2022; [`url`](#api-testing-domains-credentials-basic-auth-password-clear-secret-info-url) - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

#### API Testing Domains Credentials Bearer Token

A [`bearer_token`](#api-testing-domains-credentials-bearer-token) block (within [`api_testing.domains.credentials`](#api-testing-domains-credentials)) supports the following:

<a id="api-testing-domains-credentials-bearer-token-token"></a>&#x2022; [`token`](#api-testing-domains-credentials-bearer-token-token) - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Token](#api-testing-domains-credentials-bearer-token-token) below.

#### API Testing Domains Credentials Bearer Token Token

A [`token`](#api-testing-domains-credentials-bearer-token-token) block (within [`api_testing.domains.credentials.bearer_token`](#api-testing-domains-credentials-bearer-token)) supports the following:

<a id="api-testing-domains-credentials-bearer-token-token-blindfold-secret-info"></a>&#x2022; [`blindfold_secret_info`](#api-testing-domains-credentials-bearer-token-token-blindfold-secret-info) - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#api-testing-domains-credentials-bearer-token-token-blindfold-secret-info) below.

<a id="api-testing-domains-credentials-bearer-token-token-clear-secret-info"></a>&#x2022; [`clear_secret_info`](#api-testing-domains-credentials-bearer-token-token-clear-secret-info) - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#api-testing-domains-credentials-bearer-token-token-clear-secret-info) below.

#### API Testing Domains Credentials Bearer Token Token Blindfold Secret Info

A [`blindfold_secret_info`](#api-testing-domains-credentials-bearer-token-token-blindfold-secret-info) block (within [`api_testing.domains.credentials.bearer_token.token`](#api-testing-domains-credentials-bearer-token-token)) supports the following:

<a id="api-testing-domains-credentials-bearer-token-token-blindfold-secret-info-decryption-provider"></a>&#x2022; [`decryption_provider`](#api-testing-domains-credentials-bearer-token-token-blindfold-secret-info-decryption-provider) - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

<a id="api-testing-domains-credentials-bearer-token-token-blindfold-secret-info-location"></a>&#x2022; [`location`](#api-testing-domains-credentials-bearer-token-token-blindfold-secret-info-location) - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

<a id="api-testing-domains-credentials-bearer-token-token-blindfold-secret-info-store-provider"></a>&#x2022; [`store_provider`](#api-testing-domains-credentials-bearer-token-token-blindfold-secret-info-store-provider) - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

#### API Testing Domains Credentials Bearer Token Token Clear Secret Info

A [`clear_secret_info`](#api-testing-domains-credentials-bearer-token-token-clear-secret-info) block (within [`api_testing.domains.credentials.bearer_token.token`](#api-testing-domains-credentials-bearer-token-token)) supports the following:

<a id="api-testing-domains-credentials-bearer-token-token-clear-secret-info-provider-ref"></a>&#x2022; [`provider_ref`](#api-testing-domains-credentials-bearer-token-token-clear-secret-info-provider-ref) - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

<a id="api-testing-domains-credentials-bearer-token-token-clear-secret-info-url"></a>&#x2022; [`url`](#api-testing-domains-credentials-bearer-token-token-clear-secret-info-url) - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

#### API Testing Domains Credentials Login Endpoint

A [`login_endpoint`](#api-testing-domains-credentials-login-endpoint) block (within [`api_testing.domains.credentials`](#api-testing-domains-credentials)) supports the following:

<a id="api-testing-domains-credentials-login-endpoint-json-payload"></a>&#x2022; [`json_payload`](#api-testing-domains-credentials-login-endpoint-json-payload) - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [JSON Payload](#api-testing-domains-credentials-login-endpoint-json-payload) below.

<a id="api-testing-domains-credentials-login-endpoint-method"></a>&#x2022; [`method`](#api-testing-domains-credentials-login-endpoint-method) - Optional String  Defaults to `ANY`<br>Possible values are `ANY`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`, `COPY`<br>HTTP Method. Specifies the HTTP method used to access a resource. Any HTTP Method

<a id="api-testing-domains-credentials-login-endpoint-path"></a>&#x2022; [`path`](#api-testing-domains-credentials-login-endpoint-path) - Optional String<br>Path

<a id="api-testing-domains-credentials-login-endpoint-token-response-key"></a>&#x2022; [`token_response_key`](#api-testing-domains-credentials-login-endpoint-token-response-key) - Optional String<br>Token Response Key. Specifies how to handle the API response, extracting authentication tokens

#### API Testing Domains Credentials Login Endpoint JSON Payload

A [`json_payload`](#api-testing-domains-credentials-login-endpoint-json-payload) block (within [`api_testing.domains.credentials.login_endpoint`](#api-testing-domains-credentials-login-endpoint)) supports the following:

<a id="api-testing-domains-credentials-login-endpoint-json-payload-blindfold-secret-info"></a>&#x2022; [`blindfold_secret_info`](#api-testing-domains-credentials-login-endpoint-json-payload-blindfold-secret-info) - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#api-testing-domains-credentials-login-endpoint-json-payload-blindfold-secret-info) below.

<a id="api-testing-domains-credentials-login-endpoint-json-payload-clear-secret-info"></a>&#x2022; [`clear_secret_info`](#api-testing-domains-credentials-login-endpoint-json-payload-clear-secret-info) - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#api-testing-domains-credentials-login-endpoint-json-payload-clear-secret-info) below.

#### API Testing Domains Credentials Login Endpoint JSON Payload Blindfold Secret Info

A [`blindfold_secret_info`](#api-testing-domains-credentials-login-endpoint-json-payload-blindfold-secret-info) block (within [`api_testing.domains.credentials.login_endpoint.json_payload`](#api-testing-domains-credentials-login-endpoint-json-payload)) supports the following:

<a id="api-testing-domains-credentials-login-endpoint-json-payload-blindfold-secret-info-decryption-provider"></a>&#x2022; [`decryption_provider`](#api-testing-domains-credentials-login-endpoint-json-payload-blindfold-secret-info-decryption-provider) - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

<a id="api-testing-domains-credentials-login-endpoint-json-payload-blindfold-secret-info-location"></a>&#x2022; [`location`](#api-testing-domains-credentials-login-endpoint-json-payload-blindfold-secret-info-location) - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

<a id="api-testing-domains-credentials-login-endpoint-json-payload-blindfold-secret-info-store-provider"></a>&#x2022; [`store_provider`](#api-testing-domains-credentials-login-endpoint-json-payload-blindfold-secret-info-store-provider) - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

#### API Testing Domains Credentials Login Endpoint JSON Payload Clear Secret Info

A [`clear_secret_info`](#api-testing-domains-credentials-login-endpoint-json-payload-clear-secret-info) block (within [`api_testing.domains.credentials.login_endpoint.json_payload`](#api-testing-domains-credentials-login-endpoint-json-payload)) supports the following:

<a id="api-testing-domains-credentials-login-endpoint-json-payload-clear-secret-info-provider-ref"></a>&#x2022; [`provider_ref`](#api-testing-domains-credentials-login-endpoint-json-payload-clear-secret-info-provider-ref) - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

<a id="api-testing-domains-credentials-login-endpoint-json-payload-clear-secret-info-url"></a>&#x2022; [`url`](#api-testing-domains-credentials-login-endpoint-json-payload-clear-secret-info-url) - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

#### App Firewall

An [`app_firewall`](#app-firewall) block supports the following:

<a id="app-firewall-name"></a>&#x2022; [`name`](#app-firewall-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="app-firewall-namespace"></a>&#x2022; [`namespace`](#app-firewall-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="app-firewall-tenant"></a>&#x2022; [`tenant`](#app-firewall-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Blocked Clients

A [`blocked_clients`](#blocked-clients) block supports the following:

<a id="blocked-clients-actions"></a>&#x2022; [`actions`](#blocked-clients-actions) - Optional List  Defaults to `SKIP_PROCESSING_WAF`<br>Possible values are `SKIP_PROCESSING_WAF`, `SKIP_PROCESSING_BOT`, `SKIP_PROCESSING_MUM`, `SKIP_PROCESSING_IP_REPUTATION`, `SKIP_PROCESSING_API_PROTECTION`, `SKIP_PROCESSING_OAS_VALIDATION`, `SKIP_PROCESSING_DDOS_PROTECTION`, `SKIP_PROCESSING_THREAT_MESH`, `SKIP_PROCESSING_MALWARE_PROTECTION`<br>Actions. Actions that should be taken when client identifier matches the rule

<a id="blocked-clients-as-number"></a>&#x2022; [`as_number`](#blocked-clients-as-number) - Optional Number<br>AS Number. RFC 6793 defined 4-byte AS number

<a id="blocked-clients-bot-skip-processing"></a>&#x2022; [`bot_skip_processing`](#blocked-clients-bot-skip-processing) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="blocked-clients-expiration-timestamp"></a>&#x2022; [`expiration_timestamp`](#blocked-clients-expiration-timestamp) - Optional String<br>Expiration Timestamp. The expiration_timestamp is the RFC 3339 format timestamp at which the containing rule is considered to be logically expired. The rule continues to exist in the configuration but is not applied anymore

<a id="blocked-clients-http-header"></a>&#x2022; [`http_header`](#blocked-clients-http-header) - Optional Block<br>HTTP Header. Request header name and value pairs<br>See [HTTP Header](#blocked-clients-http-header) below.

<a id="blocked-clients-ip-prefix"></a>&#x2022; [`ip_prefix`](#blocked-clients-ip-prefix) - Optional String<br>IPv4 Prefix. IPv4 prefix string

<a id="blocked-clients-ipv6-prefix"></a>&#x2022; [`ipv6_prefix`](#blocked-clients-ipv6-prefix) - Optional String<br>IPv6 Prefix. IPv6 prefix string

<a id="blocked-clients-metadata"></a>&#x2022; [`metadata`](#blocked-clients-metadata) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#blocked-clients-metadata) below.

<a id="blocked-clients-skip-processing"></a>&#x2022; [`skip_processing`](#blocked-clients-skip-processing) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="blocked-clients-user-identifier"></a>&#x2022; [`user_identifier`](#blocked-clients-user-identifier) - Optional String<br>User Identifier. Identify user based on user identifier. User identifier value needs to be copied from security event

<a id="blocked-clients-waf-skip-processing"></a>&#x2022; [`waf_skip_processing`](#blocked-clients-waf-skip-processing) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### Blocked Clients HTTP Header

A [`http_header`](#blocked-clients-http-header) block (within [`blocked_clients`](#blocked-clients)) supports the following:

<a id="blocked-clients-http-header-headers"></a>&#x2022; [`headers`](#blocked-clients-http-header-headers) - Optional Block<br>HTTP Headers. List of HTTP header name and value pairs<br>See [Headers](#blocked-clients-http-header-headers) below.

#### Blocked Clients HTTP Header Headers

A [`headers`](#blocked-clients-http-header-headers) block (within [`blocked_clients.http_header`](#blocked-clients-http-header)) supports the following:

<a id="blocked-clients-http-header-headers-exact"></a>&#x2022; [`exact`](#blocked-clients-http-header-headers-exact) - Optional String<br>Exact. Header value to match exactly

<a id="blocked-clients-http-header-headers-invert-match"></a>&#x2022; [`invert_match`](#blocked-clients-http-header-headers-invert-match) - Optional Bool<br>NOT of match. Invert the result of the match to detect missing header or non-matching value

<a id="blocked-clients-http-header-headers-name"></a>&#x2022; [`name`](#blocked-clients-http-header-headers-name) - Optional String<br>Name. Name of the header

<a id="blocked-clients-http-header-headers-presence"></a>&#x2022; [`presence`](#blocked-clients-http-header-headers-presence) - Optional Bool<br>Presence. If true, check for presence of header

<a id="blocked-clients-http-header-headers-regex"></a>&#x2022; [`regex`](#blocked-clients-http-header-headers-regex) - Optional String<br>Regex. Regex match of the header value in re2 format

#### Blocked Clients Metadata

A [`metadata`](#blocked-clients-metadata) block (within [`blocked_clients`](#blocked-clients)) supports the following:

<a id="blocked-clients-metadata-description-spec"></a>&#x2022; [`description_spec`](#blocked-clients-metadata-description-spec) - Optional String<br>Description. Human readable description

<a id="blocked-clients-metadata-name"></a>&#x2022; [`name`](#blocked-clients-metadata-name) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### Bot Defense

A [`bot_defense`](#bot-defense) block supports the following:

<a id="bot-defense-disable-cors-support"></a>&#x2022; [`disable_cors_support`](#bot-defense-disable-cors-support) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="bot-defense-enable-cors-support"></a>&#x2022; [`enable_cors_support`](#bot-defense-enable-cors-support) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="bot-defense-policy"></a>&#x2022; [`policy`](#bot-defense-policy) - Optional Block<br>Bot Defense Policy. This defines various configuration options for Bot Defense policy<br>See [Policy](#bot-defense-policy) below.

<a id="bot-defense-regional-endpoint"></a>&#x2022; [`regional_endpoint`](#bot-defense-regional-endpoint) - Optional String  Defaults to `AUTO`<br>Possible values are `AUTO`, `US`, `EU`, `ASIA`<br>Bot Defense Region. Defines a selection for Bot Defense region - AUTO: AUTO Automatic selection based on client IP address - US: US US region - EU: EU European Union region - ASIA: ASIA Asia region

<a id="bot-defense-timeout"></a>&#x2022; [`timeout`](#bot-defense-timeout) - Optional Number<br>Timeout. The timeout for the inference check, in milliseconds

#### Bot Defense Policy

A [`policy`](#bot-defense-policy) block (within [`bot_defense`](#bot-defense)) supports the following:

<a id="bot-defense-policy-disable-js-insert"></a>&#x2022; [`disable_js_insert`](#bot-defense-policy-disable-js-insert) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="bot-defense-policy-disable-mobile-sdk"></a>&#x2022; [`disable_mobile_sdk`](#bot-defense-policy-disable-mobile-sdk) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="bot-defense-policy-javascript-mode"></a>&#x2022; [`javascript_mode`](#bot-defense-policy-javascript-mode) - Optional String  Defaults to `ASYNC_JS_NO_CACHING`<br>Possible values are `ASYNC_JS_NO_CACHING`, `ASYNC_JS_CACHING`, `SYNC_JS_NO_CACHING`, `SYNC_JS_CACHING`<br>Web Client JavaScript Mode. Web Client JavaScript Mode. Bot Defense JavaScript for telemetry collection is requested asynchronously, and it is non-cacheable Bot Defense JavaScript for telemetry collection is requested asynchronously, and it is cacheable Bot Defense JavaScript for telemetry collection is requested synchronously, and it is non-cacheable Bot Defense JavaScript for telemetry collection is requested synchronously, and it is cacheable

<a id="bot-defense-policy-js-download-path"></a>&#x2022; [`js_download_path`](#bot-defense-policy-js-download-path) - Optional String<br>JavaScript Download Path. Customize Bot Defense Client JavaScript path. If not specified, default `/common.js`

<a id="bot-defense-policy-js-insert-all-pages"></a>&#x2022; [`js_insert_all_pages`](#bot-defense-policy-js-insert-all-pages) - Optional Block<br>Insert Bot Defense JavaScript in All Pages. Insert Bot Defense JavaScript in all pages<br>See [Js Insert All Pages](#bot-defense-policy-js-insert-all-pages) below.

<a id="bot-defense-policy-js-insert-all-pages-except"></a>&#x2022; [`js_insert_all_pages_except`](#bot-defense-policy-js-insert-all-pages-except) - Optional Block<br>Insert JavaScript in All Pages with the Exceptions. Insert Bot Defense JavaScript in all pages with the exceptions<br>See [Js Insert All Pages Except](#bot-defense-policy-js-insert-all-pages-except) below.

<a id="bot-defense-policy-js-insertion-rules"></a>&#x2022; [`js_insertion_rules`](#bot-defense-policy-js-insertion-rules) - Optional Block<br>JavaScript Custom Insertion Rules. This defines custom JavaScript insertion rules for Bot Defense Policy<br>See [Js Insertion Rules](#bot-defense-policy-js-insertion-rules) below.

<a id="bot-defense-policy-mobile-sdk-config"></a>&#x2022; [`mobile_sdk_config`](#bot-defense-policy-mobile-sdk-config) - Optional Block<br>Mobile SDK Configuration. Mobile SDK configuration<br>See [Mobile Sdk Config](#bot-defense-policy-mobile-sdk-config) below.

<a id="bot-defense-policy-protected-app-endpoints"></a>&#x2022; [`protected_app_endpoints`](#bot-defense-policy-protected-app-endpoints) - Optional Block<br>App Endpoint Type. List of protected endpoints. Limit: Approx '128 endpoints per Load Balancer (LB)' upto 4 LBs, '32 endpoints per LB' after 4 LBs<br>See [Protected App Endpoints](#bot-defense-policy-protected-app-endpoints) below.

#### Bot Defense Policy Js Insert All Pages

A [`js_insert_all_pages`](#bot-defense-policy-js-insert-all-pages) block (within [`bot_defense.policy`](#bot-defense-policy)) supports the following:

<a id="bot-defense-policy-js-insert-all-pages-javascript-location"></a>&#x2022; [`javascript_location`](#bot-defense-policy-js-insert-all-pages-javascript-location) - Optional String  Defaults to `AFTER_HEAD`<br>Possible values are `AFTER_HEAD`, `AFTER_TITLE_END`, `BEFORE_SCRIPT`<br>JavaScript Location. All inside networks. Insert JavaScript after `<head>` tag Insert JavaScript after `</title>` tag. Insert JavaScript before first `<script>` tag

#### Bot Defense Policy Js Insert All Pages Except

A [`js_insert_all_pages_except`](#bot-defense-policy-js-insert-all-pages-except) block (within [`bot_defense.policy`](#bot-defense-policy)) supports the following:

<a id="bot-defense-policy-js-insert-all-pages-except-exclude-list"></a>&#x2022; [`exclude_list`](#bot-defense-policy-js-insert-all-pages-except-exclude-list) - Optional Block<br>Exclude Pages. Optional JavaScript insertions exclude list of domain and path matchers<br>See [Exclude List](#bot-defense-policy-js-insert-all-pages-except-exclude-list) below.

<a id="bot-defense-policy-js-insert-all-pages-except-javascript-location"></a>&#x2022; [`javascript_location`](#bot-defense-policy-js-insert-all-pages-except-javascript-location) - Optional String  Defaults to `AFTER_HEAD`<br>Possible values are `AFTER_HEAD`, `AFTER_TITLE_END`, `BEFORE_SCRIPT`<br>JavaScript Location. All inside networks. Insert JavaScript after `<head>` tag Insert JavaScript after `</title>` tag. Insert JavaScript before first `<script>` tag

#### Bot Defense Policy Js Insert All Pages Except Exclude List

An [`exclude_list`](#bot-defense-policy-js-insert-all-pages-except-exclude-list) block (within [`bot_defense.policy.js_insert_all_pages_except`](#bot-defense-policy-js-insert-all-pages-except)) supports the following:

<a id="bot-defense-policy-js-insert-all-pages-except-exclude-list-any-domain"></a>&#x2022; [`any_domain`](#bot-defense-policy-js-insert-all-pages-except-exclude-list-any-domain) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="bot-defense-policy-js-insert-all-pages-except-exclude-list-domain"></a>&#x2022; [`domain`](#bot-defense-policy-js-insert-all-pages-except-exclude-list-domain) - Optional Block<br>Domains. Domains names<br>See [Domain](#bot-defense-policy-js-insert-all-pages-except-exclude-list-domain) below.

<a id="bot-defense-policy-js-insert-all-pages-except-exclude-list-metadata"></a>&#x2022; [`metadata`](#bot-defense-policy-js-insert-all-pages-except-exclude-list-metadata) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#bot-defense-policy-js-insert-all-pages-except-exclude-list-metadata) below.

<a id="bot-defense-policy-js-insert-all-pages-except-exclude-list-path"></a>&#x2022; [`path`](#bot-defense-policy-js-insert-all-pages-except-exclude-list-path) - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#bot-defense-policy-js-insert-all-pages-except-exclude-list-path) below.

#### Bot Defense Policy Js Insert All Pages Except Exclude List Domain

A [`domain`](#bot-defense-policy-js-insert-all-pages-except-exclude-list-domain) block (within [`bot_defense.policy.js_insert_all_pages_except.exclude_list`](#bot-defense-policy-js-insert-all-pages-except-exclude-list)) supports the following:

<a id="bot-defense-policy-js-insert-all-pages-except-exclude-list-domain-exact-value"></a>&#x2022; [`exact_value`](#bot-defense-policy-js-insert-all-pages-except-exclude-list-domain-exact-value) - Optional String<br>Exact Value. Exact domain name

<a id="bot-defense-policy-js-insert-all-pages-except-exclude-list-domain-regex-value"></a>&#x2022; [`regex_value`](#bot-defense-policy-js-insert-all-pages-except-exclude-list-domain-regex-value) - Optional String<br>Regex Values of Domains. Regular Expression value for the domain name

<a id="bot-defense-policy-js-insert-all-pages-except-exclude-list-domain-suffix-value"></a>&#x2022; [`suffix_value`](#bot-defense-policy-js-insert-all-pages-except-exclude-list-domain-suffix-value) - Optional String<br>Suffix Value. Suffix of domain name e.g 'xyz.com' will match '*.xyz.com' and 'xyz.com'

#### Bot Defense Policy Js Insert All Pages Except Exclude List Metadata

A [`metadata`](#bot-defense-policy-js-insert-all-pages-except-exclude-list-metadata) block (within [`bot_defense.policy.js_insert_all_pages_except.exclude_list`](#bot-defense-policy-js-insert-all-pages-except-exclude-list)) supports the following:

<a id="bot-defense-policy-js-insert-all-pages-except-exclude-list-metadata-description-spec"></a>&#x2022; [`description_spec`](#bot-defense-policy-js-insert-all-pages-except-exclude-list-metadata-description-spec) - Optional String<br>Description. Human readable description

<a id="bot-defense-policy-js-insert-all-pages-except-exclude-list-metadata-name"></a>&#x2022; [`name`](#bot-defense-policy-js-insert-all-pages-except-exclude-list-metadata-name) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### Bot Defense Policy Js Insert All Pages Except Exclude List Path

A [`path`](#bot-defense-policy-js-insert-all-pages-except-exclude-list-path) block (within [`bot_defense.policy.js_insert_all_pages_except.exclude_list`](#bot-defense-policy-js-insert-all-pages-except-exclude-list)) supports the following:

<a id="bot-defense-policy-js-insert-all-pages-except-exclude-list-path-path"></a>&#x2022; [`path`](#bot-defense-policy-js-insert-all-pages-except-exclude-list-path-path) - Optional String<br>Exact. Exact path value to match

<a id="bot-defense-policy-js-insert-all-pages-except-exclude-list-path-prefix"></a>&#x2022; [`prefix`](#bot-defense-policy-js-insert-all-pages-except-exclude-list-path-prefix) - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

<a id="bot-defense-policy-js-insert-all-pages-except-exclude-list-path-regex"></a>&#x2022; [`regex`](#bot-defense-policy-js-insert-all-pages-except-exclude-list-path-regex) - Optional String<br>Regex. Regular expression of path match (e.g. the value .* will match on all paths)

#### Bot Defense Policy Js Insertion Rules

A [`js_insertion_rules`](#bot-defense-policy-js-insertion-rules) block (within [`bot_defense.policy`](#bot-defense-policy)) supports the following:

<a id="bot-defense-policy-js-insertion-rules-exclude-list"></a>&#x2022; [`exclude_list`](#bot-defense-policy-js-insertion-rules-exclude-list) - Optional Block<br>Exclude Paths. Optional JavaScript insertions exclude list of domain and path matchers<br>See [Exclude List](#bot-defense-policy-js-insertion-rules-exclude-list) below.

<a id="bot-defense-policy-js-insertion-rules-rules"></a>&#x2022; [`rules`](#bot-defense-policy-js-insertion-rules-rules) - Optional Block<br>JavaScript Insertions. Required list of pages to insert Bot Defense client JavaScript<br>See [Rules](#bot-defense-policy-js-insertion-rules-rules) below.

#### Bot Defense Policy Js Insertion Rules Exclude List

An [`exclude_list`](#bot-defense-policy-js-insertion-rules-exclude-list) block (within [`bot_defense.policy.js_insertion_rules`](#bot-defense-policy-js-insertion-rules)) supports the following:

<a id="bot-defense-policy-js-insertion-rules-exclude-list-any-domain"></a>&#x2022; [`any_domain`](#bot-defense-policy-js-insertion-rules-exclude-list-any-domain) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="bot-defense-policy-js-insertion-rules-exclude-list-domain"></a>&#x2022; [`domain`](#bot-defense-policy-js-insertion-rules-exclude-list-domain) - Optional Block<br>Domains. Domains names<br>See [Domain](#bot-defense-policy-js-insertion-rules-exclude-list-domain) below.

<a id="bot-defense-policy-js-insertion-rules-exclude-list-metadata"></a>&#x2022; [`metadata`](#bot-defense-policy-js-insertion-rules-exclude-list-metadata) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#bot-defense-policy-js-insertion-rules-exclude-list-metadata) below.

<a id="bot-defense-policy-js-insertion-rules-exclude-list-path"></a>&#x2022; [`path`](#bot-defense-policy-js-insertion-rules-exclude-list-path) - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#bot-defense-policy-js-insertion-rules-exclude-list-path) below.

#### Bot Defense Policy Js Insertion Rules Exclude List Domain

A [`domain`](#bot-defense-policy-js-insertion-rules-exclude-list-domain) block (within [`bot_defense.policy.js_insertion_rules.exclude_list`](#bot-defense-policy-js-insertion-rules-exclude-list)) supports the following:

<a id="bot-defense-policy-js-insertion-rules-exclude-list-domain-exact-value"></a>&#x2022; [`exact_value`](#bot-defense-policy-js-insertion-rules-exclude-list-domain-exact-value) - Optional String<br>Exact Value. Exact domain name

<a id="bot-defense-policy-js-insertion-rules-exclude-list-domain-regex-value"></a>&#x2022; [`regex_value`](#bot-defense-policy-js-insertion-rules-exclude-list-domain-regex-value) - Optional String<br>Regex Values of Domains. Regular Expression value for the domain name

<a id="bot-defense-policy-js-insertion-rules-exclude-list-domain-suffix-value"></a>&#x2022; [`suffix_value`](#bot-defense-policy-js-insertion-rules-exclude-list-domain-suffix-value) - Optional String<br>Suffix Value. Suffix of domain name e.g 'xyz.com' will match '*.xyz.com' and 'xyz.com'

#### Bot Defense Policy Js Insertion Rules Exclude List Metadata

A [`metadata`](#bot-defense-policy-js-insertion-rules-exclude-list-metadata) block (within [`bot_defense.policy.js_insertion_rules.exclude_list`](#bot-defense-policy-js-insertion-rules-exclude-list)) supports the following:

<a id="bot-defense-policy-js-insertion-rules-exclude-list-metadata-description-spec"></a>&#x2022; [`description_spec`](#bot-defense-policy-js-insertion-rules-exclude-list-metadata-description-spec) - Optional String<br>Description. Human readable description

<a id="bot-defense-policy-js-insertion-rules-exclude-list-metadata-name"></a>&#x2022; [`name`](#bot-defense-policy-js-insertion-rules-exclude-list-metadata-name) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### Bot Defense Policy Js Insertion Rules Exclude List Path

A [`path`](#bot-defense-policy-js-insertion-rules-exclude-list-path) block (within [`bot_defense.policy.js_insertion_rules.exclude_list`](#bot-defense-policy-js-insertion-rules-exclude-list)) supports the following:

<a id="bot-defense-policy-js-insertion-rules-exclude-list-path-path"></a>&#x2022; [`path`](#bot-defense-policy-js-insertion-rules-exclude-list-path-path) - Optional String<br>Exact. Exact path value to match

<a id="bot-defense-policy-js-insertion-rules-exclude-list-path-prefix"></a>&#x2022; [`prefix`](#bot-defense-policy-js-insertion-rules-exclude-list-path-prefix) - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

<a id="bot-defense-policy-js-insertion-rules-exclude-list-path-regex"></a>&#x2022; [`regex`](#bot-defense-policy-js-insertion-rules-exclude-list-path-regex) - Optional String<br>Regex. Regular expression of path match (e.g. the value .* will match on all paths)

#### Bot Defense Policy Js Insertion Rules Rules

A [`rules`](#bot-defense-policy-js-insertion-rules-rules) block (within [`bot_defense.policy.js_insertion_rules`](#bot-defense-policy-js-insertion-rules)) supports the following:

<a id="bot-defense-policy-js-insertion-rules-rules-any-domain"></a>&#x2022; [`any_domain`](#bot-defense-policy-js-insertion-rules-rules-any-domain) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="bot-defense-policy-js-insertion-rules-rules-domain"></a>&#x2022; [`domain`](#bot-defense-policy-js-insertion-rules-rules-domain) - Optional Block<br>Domains. Domains names<br>See [Domain](#bot-defense-policy-js-insertion-rules-rules-domain) below.

<a id="bot-defense-policy-js-insertion-rules-rules-javascript-location"></a>&#x2022; [`javascript_location`](#bot-defense-policy-js-insertion-rules-rules-javascript-location) - Optional String  Defaults to `AFTER_HEAD`<br>Possible values are `AFTER_HEAD`, `AFTER_TITLE_END`, `BEFORE_SCRIPT`<br>JavaScript Location. All inside networks. Insert JavaScript after `<head>` tag Insert JavaScript after `</title>` tag. Insert JavaScript before first `<script>` tag

<a id="bot-defense-policy-js-insertion-rules-rules-metadata"></a>&#x2022; [`metadata`](#bot-defense-policy-js-insertion-rules-rules-metadata) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#bot-defense-policy-js-insertion-rules-rules-metadata) below.

<a id="bot-defense-policy-js-insertion-rules-rules-path"></a>&#x2022; [`path`](#bot-defense-policy-js-insertion-rules-rules-path) - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#bot-defense-policy-js-insertion-rules-rules-path) below.

#### Bot Defense Policy Js Insertion Rules Rules Domain

A [`domain`](#bot-defense-policy-js-insertion-rules-rules-domain) block (within [`bot_defense.policy.js_insertion_rules.rules`](#bot-defense-policy-js-insertion-rules-rules)) supports the following:

<a id="bot-defense-policy-js-insertion-rules-rules-domain-exact-value"></a>&#x2022; [`exact_value`](#bot-defense-policy-js-insertion-rules-rules-domain-exact-value) - Optional String<br>Exact Value. Exact domain name

<a id="bot-defense-policy-js-insertion-rules-rules-domain-regex-value"></a>&#x2022; [`regex_value`](#bot-defense-policy-js-insertion-rules-rules-domain-regex-value) - Optional String<br>Regex Values of Domains. Regular Expression value for the domain name

<a id="bot-defense-policy-js-insertion-rules-rules-domain-suffix-value"></a>&#x2022; [`suffix_value`](#bot-defense-policy-js-insertion-rules-rules-domain-suffix-value) - Optional String<br>Suffix Value. Suffix of domain name e.g 'xyz.com' will match '*.xyz.com' and 'xyz.com'

#### Bot Defense Policy Js Insertion Rules Rules Metadata

A [`metadata`](#bot-defense-policy-js-insertion-rules-rules-metadata) block (within [`bot_defense.policy.js_insertion_rules.rules`](#bot-defense-policy-js-insertion-rules-rules)) supports the following:

<a id="bot-defense-policy-js-insertion-rules-rules-metadata-description-spec"></a>&#x2022; [`description_spec`](#bot-defense-policy-js-insertion-rules-rules-metadata-description-spec) - Optional String<br>Description. Human readable description

<a id="bot-defense-policy-js-insertion-rules-rules-metadata-name"></a>&#x2022; [`name`](#bot-defense-policy-js-insertion-rules-rules-metadata-name) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### Bot Defense Policy Js Insertion Rules Rules Path

A [`path`](#bot-defense-policy-js-insertion-rules-rules-path) block (within [`bot_defense.policy.js_insertion_rules.rules`](#bot-defense-policy-js-insertion-rules-rules)) supports the following:

<a id="bot-defense-policy-js-insertion-rules-rules-path-path"></a>&#x2022; [`path`](#bot-defense-policy-js-insertion-rules-rules-path-path) - Optional String<br>Exact. Exact path value to match

<a id="bot-defense-policy-js-insertion-rules-rules-path-prefix"></a>&#x2022; [`prefix`](#bot-defense-policy-js-insertion-rules-rules-path-prefix) - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

<a id="bot-defense-policy-js-insertion-rules-rules-path-regex"></a>&#x2022; [`regex`](#bot-defense-policy-js-insertion-rules-rules-path-regex) - Optional String<br>Regex. Regular expression of path match (e.g. the value .* will match on all paths)

#### Bot Defense Policy Mobile Sdk Config

A [`mobile_sdk_config`](#bot-defense-policy-mobile-sdk-config) block (within [`bot_defense.policy`](#bot-defense-policy)) supports the following:

<a id="bot-defense-policy-mobile-sdk-config-mobile-identifier"></a>&#x2022; [`mobile_identifier`](#bot-defense-policy-mobile-sdk-config-mobile-identifier) - Optional Block<br>Mobile Traffic Identifier. Mobile traffic identifier type<br>See [Mobile Identifier](#bot-defense-policy-mobile-sdk-config-mobile-identifier) below.

#### Bot Defense Policy Mobile Sdk Config Mobile Identifier

A [`mobile_identifier`](#bot-defense-policy-mobile-sdk-config-mobile-identifier) block (within [`bot_defense.policy.mobile_sdk_config`](#bot-defense-policy-mobile-sdk-config)) supports the following:

<a id="bot-defense-policy-mobile-sdk-config-mobile-identifier-headers"></a>&#x2022; [`headers`](#bot-defense-policy-mobile-sdk-config-mobile-identifier-headers) - Optional Block<br>Headers. Headers that can be used to identify mobile traffic<br>See [Headers](#bot-defense-policy-mobile-sdk-config-mobile-identifier-headers) below.

#### Bot Defense Policy Mobile Sdk Config Mobile Identifier Headers

A [`headers`](#bot-defense-policy-mobile-sdk-config-mobile-identifier-headers) block (within [`bot_defense.policy.mobile_sdk_config.mobile_identifier`](#bot-defense-policy-mobile-sdk-config-mobile-identifier)) supports the following:

<a id="bot-defense-policy-mobile-sdk-config-mobile-identifier-headers-check-not-present"></a>&#x2022; [`check_not_present`](#bot-defense-policy-mobile-sdk-config-mobile-identifier-headers-check-not-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="bot-defense-policy-mobile-sdk-config-mobile-identifier-headers-check-present"></a>&#x2022; [`check_present`](#bot-defense-policy-mobile-sdk-config-mobile-identifier-headers-check-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="bot-defense-policy-mobile-sdk-config-mobile-identifier-headers-item"></a>&#x2022; [`item`](#bot-defense-policy-mobile-sdk-config-mobile-identifier-headers-item) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#bot-defense-policy-mobile-sdk-config-mobile-identifier-headers-item) below.

<a id="bot-defense-policy-mobile-sdk-config-mobile-identifier-headers-name"></a>&#x2022; [`name`](#bot-defense-policy-mobile-sdk-config-mobile-identifier-headers-name) - Optional String<br>Header Name. A case-insensitive HTTP header name

#### Bot Defense Policy Mobile Sdk Config Mobile Identifier Headers Item

An [`item`](#bot-defense-policy-mobile-sdk-config-mobile-identifier-headers-item) block (within [`bot_defense.policy.mobile_sdk_config.mobile_identifier.headers`](#bot-defense-policy-mobile-sdk-config-mobile-identifier-headers)) supports the following:

<a id="bot-defense-policy-mobile-sdk-config-mobile-identifier-headers-item-exact-values"></a>&#x2022; [`exact_values`](#bot-defense-policy-mobile-sdk-config-mobile-identifier-headers-item-exact-values) - Optional List<br>Exact Values. A list of exact values to match the input against

<a id="bot-defense-policy-mobile-sdk-config-mobile-identifier-headers-item-regex-values"></a>&#x2022; [`regex_values`](#bot-defense-policy-mobile-sdk-config-mobile-identifier-headers-item-regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="bot-defense-policy-mobile-sdk-config-mobile-identifier-headers-item-transformers"></a>&#x2022; [`transformers`](#bot-defense-policy-mobile-sdk-config-mobile-identifier-headers-item-transformers) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### Bot Defense Policy Protected App Endpoints

A [`protected_app_endpoints`](#bot-defense-policy-protected-app-endpoints) block (within [`bot_defense.policy`](#bot-defense-policy)) supports the following:

<a id="bot-defense-policy-protected-app-endpoints-allow-good-bots"></a>&#x2022; [`allow_good_bots`](#bot-defense-policy-protected-app-endpoints-allow-good-bots) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="bot-defense-policy-protected-app-endpoints-any-domain"></a>&#x2022; [`any_domain`](#bot-defense-policy-protected-app-endpoints-any-domain) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="bot-defense-policy-protected-app-endpoints-domain"></a>&#x2022; [`domain`](#bot-defense-policy-protected-app-endpoints-domain) - Optional Block<br>Domains. Domains names<br>See [Domain](#bot-defense-policy-protected-app-endpoints-domain) below.

<a id="bot-defense-policy-protected-app-endpoints-flow-label"></a>&#x2022; [`flow_label`](#bot-defense-policy-protected-app-endpoints-flow-label) - Optional Block<br>Bot Defense Flow Label Category. Bot Defense Flow Label Category allows to associate traffic with selected category<br>See [Flow Label](#bot-defense-policy-protected-app-endpoints-flow-label) below.

<a id="bot-defense-policy-protected-app-endpoints-headers"></a>&#x2022; [`headers`](#bot-defense-policy-protected-app-endpoints-headers) - Optional Block<br>HTTP Headers. A list of predicates for various HTTP headers that need to match. The criteria for matching each HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values are extracted from the request API as a list of strings for each HTTP header type. Note that all specified header predicates must evaluate to true<br>See [Headers](#bot-defense-policy-protected-app-endpoints-headers) below.

<a id="bot-defense-policy-protected-app-endpoints-http-methods"></a>&#x2022; [`http_methods`](#bot-defense-policy-protected-app-endpoints-http-methods) - Optional List  Defaults to `METHOD_ANY`<br>Possible values are `METHOD_ANY`, `METHOD_GET`, `METHOD_POST`, `METHOD_PUT`, `METHOD_PATCH`, `METHOD_DELETE`, `METHOD_GET_DOCUMENT`<br>HTTP Methods. List of HTTP methods

<a id="bot-defense-policy-protected-app-endpoints-metadata"></a>&#x2022; [`metadata`](#bot-defense-policy-protected-app-endpoints-metadata) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#bot-defense-policy-protected-app-endpoints-metadata) below.

<a id="bot-defense-policy-protected-app-endpoints-mitigate-good-bots"></a>&#x2022; [`mitigate_good_bots`](#bot-defense-policy-protected-app-endpoints-mitigate-good-bots) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="bot-defense-policy-protected-app-endpoints-mitigation"></a>&#x2022; [`mitigation`](#bot-defense-policy-protected-app-endpoints-mitigation) - Optional Block<br>Bot Mitigation Action. Modify Bot Defense behavior for a matching request<br>See [Mitigation](#bot-defense-policy-protected-app-endpoints-mitigation) below.

<a id="bot-defense-policy-protected-app-endpoints-mobile"></a>&#x2022; [`mobile`](#bot-defense-policy-protected-app-endpoints-mobile) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="bot-defense-policy-protected-app-endpoints-path"></a>&#x2022; [`path`](#bot-defense-policy-protected-app-endpoints-path) - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#bot-defense-policy-protected-app-endpoints-path) below.

<a id="bot-defense-policy-protected-app-endpoints-protocol"></a>&#x2022; [`protocol`](#bot-defense-policy-protected-app-endpoints-protocol) - Optional String  Defaults to `BOTH`<br>Possible values are `BOTH`, `HTTP`, `HTTPS`<br>URL Scheme. SchemeType is used to indicate URL scheme. - BOTH: BOTH URL scheme for HTTPS:// or `HTTP://.` - HTTP: HTTP URL scheme HTTP:// only. - HTTPS: HTTPS URL scheme HTTPS:// only

<a id="bot-defense-policy-protected-app-endpoints-query-params"></a>&#x2022; [`query_params`](#bot-defense-policy-protected-app-endpoints-query-params) - Optional Block<br>HTTP Query Parameters. A list of predicates for all query parameters that need to be matched. The criteria for matching each query parameter are described in individual instances of QueryParameterMatcherType. The actual query parameter values are extracted from the request API as a list of strings for each query parameter name. Note that all specified query parameter predicates must evaluate to true<br>See [Query Params](#bot-defense-policy-protected-app-endpoints-query-params) below.

<a id="bot-defense-policy-protected-app-endpoints-undefined-flow-label"></a>&#x2022; [`undefined_flow_label`](#bot-defense-policy-protected-app-endpoints-undefined-flow-label) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="bot-defense-policy-protected-app-endpoints-web"></a>&#x2022; [`web`](#bot-defense-policy-protected-app-endpoints-web) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="bot-defense-policy-protected-app-endpoints-web-mobile"></a>&#x2022; [`web_mobile`](#bot-defense-policy-protected-app-endpoints-web-mobile) - Optional Block<br>Web and Mobile traffic type. Web and Mobile traffic type<br>See [Web Mobile](#bot-defense-policy-protected-app-endpoints-web-mobile) below.

#### Bot Defense Policy Protected App Endpoints Domain

A [`domain`](#bot-defense-policy-protected-app-endpoints-domain) block (within [`bot_defense.policy.protected_app_endpoints`](#bot-defense-policy-protected-app-endpoints)) supports the following:

<a id="bot-defense-policy-protected-app-endpoints-domain-exact-value"></a>&#x2022; [`exact_value`](#bot-defense-policy-protected-app-endpoints-domain-exact-value) - Optional String<br>Exact Value. Exact domain name

<a id="bot-defense-policy-protected-app-endpoints-domain-regex-value"></a>&#x2022; [`regex_value`](#bot-defense-policy-protected-app-endpoints-domain-regex-value) - Optional String<br>Regex Values of Domains. Regular Expression value for the domain name

<a id="bot-defense-policy-protected-app-endpoints-domain-suffix-value"></a>&#x2022; [`suffix_value`](#bot-defense-policy-protected-app-endpoints-domain-suffix-value) - Optional String<br>Suffix Value. Suffix of domain name e.g 'xyz.com' will match '*.xyz.com' and 'xyz.com'

#### Bot Defense Policy Protected App Endpoints Flow Label

A [`flow_label`](#bot-defense-policy-protected-app-endpoints-flow-label) block (within [`bot_defense.policy.protected_app_endpoints`](#bot-defense-policy-protected-app-endpoints)) supports the following:

<a id="bot-defense-policy-protected-app-endpoints-flow-label-account-management"></a>&#x2022; [`account_management`](#bot-defense-policy-protected-app-endpoints-flow-label-account-management) - Optional Block<br>Bot Defense Flow Label Account Management Category. Bot Defense Flow Label Account Management Category<br>See [Account Management](#bot-defense-policy-protected-app-endpoints-flow-label-account-management) below.

<a id="bot-defense-policy-protected-app-endpoints-flow-label-authentication"></a>&#x2022; [`authentication`](#bot-defense-policy-protected-app-endpoints-flow-label-authentication) - Optional Block<br>Bot Defense Flow Label Authentication Category. Bot Defense Flow Label Authentication Category<br>See [Authentication](#bot-defense-policy-protected-app-endpoints-flow-label-authentication) below.

<a id="bot-defense-policy-protected-app-endpoints-flow-label-financial-services"></a>&#x2022; [`financial_services`](#bot-defense-policy-protected-app-endpoints-flow-label-financial-services) - Optional Block<br>Bot Defense Flow Label Financial Services Category. Bot Defense Flow Label Financial Services Category<br>See [Financial Services](#bot-defense-policy-protected-app-endpoints-flow-label-financial-services) below.

<a id="bot-defense-policy-protected-app-endpoints-flow-label-flight"></a>&#x2022; [`flight`](#bot-defense-policy-protected-app-endpoints-flow-label-flight) - Optional Block<br>Bot Defense Flow Label Flight Category. Bot Defense Flow Label Flight Category<br>See [Flight](#bot-defense-policy-protected-app-endpoints-flow-label-flight) below.

<a id="bot-defense-policy-protected-app-endpoints-flow-label-profile-management"></a>&#x2022; [`profile_management`](#bot-defense-policy-protected-app-endpoints-flow-label-profile-management) - Optional Block<br>Bot Defense Flow Label Profile Management Category. Bot Defense Flow Label Profile Management Category<br>See [Profile Management](#bot-defense-policy-protected-app-endpoints-flow-label-profile-management) below.

<a id="bot-defense-policy-protected-app-endpoints-flow-label-search"></a>&#x2022; [`search`](#bot-defense-policy-protected-app-endpoints-flow-label-search) - Optional Block<br>Bot Defense Flow Label Search Category. Bot Defense Flow Label Search Category<br>See [Search](#bot-defense-policy-protected-app-endpoints-flow-label-search) below.

<a id="bot-defense-policy-protected-app-endpoints-flow-label-shopping-gift-cards"></a>&#x2022; [`shopping_gift_cards`](#bot-defense-policy-protected-app-endpoints-flow-label-shopping-gift-cards) - Optional Block<br>Bot Defense Flow Label Shopping & Gift Cards Category. Bot Defense Flow Label Shopping & Gift Cards Category<br>See [Shopping Gift Cards](#bot-defense-policy-protected-app-endpoints-flow-label-shopping-gift-cards) below.

#### Bot Defense Policy Protected App Endpoints Flow Label Account Management

An [`account_management`](#bot-defense-policy-protected-app-endpoints-flow-label-account-management) block (within [`bot_defense.policy.protected_app_endpoints.flow_label`](#bot-defense-policy-protected-app-endpoints-flow-label)) supports the following:

<a id="bot-defense-policy-protected-app-endpoints-flow-label-account-management-create"></a>&#x2022; [`create`](#bot-defense-policy-protected-app-endpoints-flow-label-account-management-create) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="bot-defense-policy-protected-app-endpoints-flow-label-account-management-password-reset"></a>&#x2022; [`password_reset`](#bot-defense-policy-protected-app-endpoints-flow-label-account-management-password-reset) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### Bot Defense Policy Protected App Endpoints Flow Label Authentication

An [`authentication`](#bot-defense-policy-protected-app-endpoints-flow-label-authentication) block (within [`bot_defense.policy.protected_app_endpoints.flow_label`](#bot-defense-policy-protected-app-endpoints-flow-label)) supports the following:

<a id="bot-defense-policy-protected-app-endpoints-flow-label-authentication-login"></a>&#x2022; [`login`](#bot-defense-policy-protected-app-endpoints-flow-label-authentication-login) - Optional Block<br>Bot Defense Transaction Result. Bot Defense Transaction Result<br>See [Login](#bot-defense-policy-protected-app-endpoints-flow-label-authentication-login) below.

<a id="bot-defense-policy-protected-app-endpoints-flow-label-authentication-login-mfa"></a>&#x2022; [`login_mfa`](#bot-defense-policy-protected-app-endpoints-flow-label-authentication-login-mfa) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="bot-defense-policy-protected-app-endpoints-flow-label-authentication-login-partner"></a>&#x2022; [`login_partner`](#bot-defense-policy-protected-app-endpoints-flow-label-authentication-login-partner) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="bot-defense-policy-protected-app-endpoints-flow-label-authentication-logout"></a>&#x2022; [`logout`](#bot-defense-policy-protected-app-endpoints-flow-label-authentication-logout) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="bot-defense-policy-protected-app-endpoints-flow-label-authentication-token-refresh"></a>&#x2022; [`token_refresh`](#bot-defense-policy-protected-app-endpoints-flow-label-authentication-token-refresh) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### Bot Defense Policy Protected App Endpoints Flow Label Authentication Login

A [`login`](#bot-defense-policy-protected-app-endpoints-flow-label-authentication-login) block (within [`bot_defense.policy.protected_app_endpoints.flow_label.authentication`](#bot-defense-policy-protected-app-endpoints-flow-label-authentication)) supports the following:

<a id="bot-defense-policy-protected-app-endpoints-flow-label-authentication-login-disable-transaction-result"></a>&#x2022; [`disable_transaction_result`](#bot-defense-policy-protected-app-endpoints-flow-label-authentication-login-disable-transaction-result) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="bot-defense-policy-protected-app-endpoints-flow-label-authentication-login-transaction-result"></a>&#x2022; [`transaction_result`](#bot-defense-policy-protected-app-endpoints-flow-label-authentication-login-transaction-result) - Optional Block<br>Bot Defense Transaction Result Type. Bot Defense Transaction ResultType<br>See [Transaction Result](#bot-defense-policy-protected-app-endpoints-flow-label-authentication-login-transaction-result) below.

#### Bot Defense Policy Protected App Endpoints Flow Label Authentication Login Transaction Result

A [`transaction_result`](#bot-defense-policy-protected-app-endpoints-flow-label-authentication-login-transaction-result) block (within [`bot_defense.policy.protected_app_endpoints.flow_label.authentication.login`](#bot-defense-policy-protected-app-endpoints-flow-label-authentication-login)) supports the following:

<a id="bot-defense-policy-protected-app-endpoints-flow-label-authentication-login-transaction-result-failure-conditions"></a>&#x2022; [`failure_conditions`](#bot-defense-policy-protected-app-endpoints-flow-label-authentication-login-transaction-result-failure-conditions) - Optional Block<br>Failure Conditions. Failure Conditions<br>See [Failure Conditions](#bot-defense-policy-protected-app-endpoints-flow-label-authentication-login-transaction-result-failure-conditions) below.

<a id="bot-defense-policy-protected-app-endpoints-flow-label-authentication-login-transaction-result-success-conditions"></a>&#x2022; [`success_conditions`](#bot-defense-policy-protected-app-endpoints-flow-label-authentication-login-transaction-result-success-conditions) - Optional Block<br>Success Conditions. Success Conditions<br>See [Success Conditions](#bot-defense-policy-protected-app-endpoints-flow-label-authentication-login-transaction-result-success-conditions) below.

#### Bot Defense Policy Protected App Endpoints Flow Label Authentication Login Transaction Result Failure Conditions

A [`failure_conditions`](#bot-defense-policy-protected-app-endpoints-flow-label-authentication-login-transaction-result-failure-conditions) block (within [`bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result`](#bot-defense-policy-protected-app-endpoints-flow-label-authentication-login-transaction-result)) supports the following:

<a id="bot-defense-policy-protected-app-endpoints-flow-label-authentication-login-transaction-result-failure-conditions-name"></a>&#x2022; [`name`](#bot-defense-policy-protected-app-endpoints-flow-label-authentication-login-transaction-result-failure-conditions-name) - Optional String<br>Header Name. A case-insensitive HTTP header name

<a id="bot-defense-policy-protected-app-endpoints-flow-label-authentication-login-transaction-result-failure-conditions-regex-values"></a>&#x2022; [`regex_values`](#bot-defense-policy-protected-app-endpoints-flow-label-authentication-login-transaction-result-failure-conditions-regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="bot-defense-policy-protected-app-endpoints-flow-label-authentication-login-transaction-result-failure-conditions-status"></a>&#x2022; [`status`](#bot-defense-policy-protected-app-endpoints-flow-label-authentication-login-transaction-result-failure-conditions-status) - Optional String  Defaults to `EmptyStatusCode`<br>Possible values are `EmptyStatusCode`, `Continue`, `OK`, `Created`, `Accepted`, `NonAuthoritativeInformation`, `NoContent`, `ResetContent`, `PartialContent`, `MultiStatus`, `AlreadyReported`, `IMUsed`, `MultipleChoices`, `MovedPermanently`, `Found`, `SeeOther`, `NotModified`, `UseProxy`, `TemporaryRedirect`, `PermanentRedirect`, `BadRequest`, `Unauthorized`, `PaymentRequired`, `Forbidden`, `NotFound`, `MethodNotAllowed`, `NotAcceptable`, `ProxyAuthenticationRequired`, `RequestTimeout`, `Conflict`, `Gone`, `LengthRequired`, `PreconditionFailed`, `PayloadTooLarge`, `URITooLong`, `UnsupportedMediaType`, `RangeNotSatisfiable`, `ExpectationFailed`, `MisdirectedRequest`, `UnprocessableEntity`, `Locked`, `FailedDependency`, `UpgradeRequired`, `PreconditionRequired`, `TooManyRequests`, `RequestHeaderFieldsTooLarge`, `InternalServerError`, `NotImplemented`, `BadGateway`, `ServiceUnavailable`, `GatewayTimeout`, `HTTPVersionNotSupported`, `VariantAlsoNegotiates`, `InsufficientStorage`, `LoopDetected`, `NotExtended`, `NetworkAuthenticationRequired`<br>HTTP Status Code. HTTP response status codes EmptyStatusCode response codes means it is not specified Continue status code OK status code Created status code Accepted status code Non Authoritative Information status code No Content status code Reset Content status code Partial Content status code Multi Status status code Already Reported status code Im Used status code Multiple Choices status code Moved Permanently status code Found status code See Other status code Not Modified status code Use Proxy status code Temporary Redirect status code Permanent Redirect status code Bad Request status code Unauthorized status code Payment Required status code Forbidden status code Not Found status code Method Not Allowed status code Not Acceptable status code Proxy Authentication Required status code Request Timeout status code Conflict status code Gone status code Length Required status code Precondition Failed status code Payload Too Large status code URI Too Long status code Unsupported Media Type status code Range Not Satisfiable status code Expectation Failed status code Misdirected Request status code Unprocessable Entity status code Locked status code Failed Dependency status code Upgrade Required status code Precondition Required status code Too Many Requests status code Request Header Fields Too Large status code Internal Server Error status code Not Implemented status code Bad Gateway status code Service Unavailable status code Gateway Timeout status code HTTP Version Not Supported status code Variant Also Negotiates status code Insufficient Storage status code Loop Detected status code Not Extended status code Network Authentication Required status code

#### Bot Defense Policy Protected App Endpoints Flow Label Authentication Login Transaction Result Success Conditions

A [`success_conditions`](#bot-defense-policy-protected-app-endpoints-flow-label-authentication-login-transaction-result-success-conditions) block (within [`bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result`](#bot-defense-policy-protected-app-endpoints-flow-label-authentication-login-transaction-result)) supports the following:

<a id="bot-defense-policy-protected-app-endpoints-flow-label-authentication-login-transaction-result-success-conditions-name"></a>&#x2022; [`name`](#bot-defense-policy-protected-app-endpoints-flow-label-authentication-login-transaction-result-success-conditions-name) - Optional String<br>Header Name. A case-insensitive HTTP header name

<a id="bot-defense-policy-protected-app-endpoints-flow-label-authentication-login-transaction-result-success-conditions-regex-values"></a>&#x2022; [`regex_values`](#bot-defense-policy-protected-app-endpoints-flow-label-authentication-login-transaction-result-success-conditions-regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="bot-defense-policy-protected-app-endpoints-flow-label-authentication-login-transaction-result-success-conditions-status"></a>&#x2022; [`status`](#bot-defense-policy-protected-app-endpoints-flow-label-authentication-login-transaction-result-success-conditions-status) - Optional String  Defaults to `EmptyStatusCode`<br>Possible values are `EmptyStatusCode`, `Continue`, `OK`, `Created`, `Accepted`, `NonAuthoritativeInformation`, `NoContent`, `ResetContent`, `PartialContent`, `MultiStatus`, `AlreadyReported`, `IMUsed`, `MultipleChoices`, `MovedPermanently`, `Found`, `SeeOther`, `NotModified`, `UseProxy`, `TemporaryRedirect`, `PermanentRedirect`, `BadRequest`, `Unauthorized`, `PaymentRequired`, `Forbidden`, `NotFound`, `MethodNotAllowed`, `NotAcceptable`, `ProxyAuthenticationRequired`, `RequestTimeout`, `Conflict`, `Gone`, `LengthRequired`, `PreconditionFailed`, `PayloadTooLarge`, `URITooLong`, `UnsupportedMediaType`, `RangeNotSatisfiable`, `ExpectationFailed`, `MisdirectedRequest`, `UnprocessableEntity`, `Locked`, `FailedDependency`, `UpgradeRequired`, `PreconditionRequired`, `TooManyRequests`, `RequestHeaderFieldsTooLarge`, `InternalServerError`, `NotImplemented`, `BadGateway`, `ServiceUnavailable`, `GatewayTimeout`, `HTTPVersionNotSupported`, `VariantAlsoNegotiates`, `InsufficientStorage`, `LoopDetected`, `NotExtended`, `NetworkAuthenticationRequired`<br>HTTP Status Code. HTTP response status codes EmptyStatusCode response codes means it is not specified Continue status code OK status code Created status code Accepted status code Non Authoritative Information status code No Content status code Reset Content status code Partial Content status code Multi Status status code Already Reported status code Im Used status code Multiple Choices status code Moved Permanently status code Found status code See Other status code Not Modified status code Use Proxy status code Temporary Redirect status code Permanent Redirect status code Bad Request status code Unauthorized status code Payment Required status code Forbidden status code Not Found status code Method Not Allowed status code Not Acceptable status code Proxy Authentication Required status code Request Timeout status code Conflict status code Gone status code Length Required status code Precondition Failed status code Payload Too Large status code URI Too Long status code Unsupported Media Type status code Range Not Satisfiable status code Expectation Failed status code Misdirected Request status code Unprocessable Entity status code Locked status code Failed Dependency status code Upgrade Required status code Precondition Required status code Too Many Requests status code Request Header Fields Too Large status code Internal Server Error status code Not Implemented status code Bad Gateway status code Service Unavailable status code Gateway Timeout status code HTTP Version Not Supported status code Variant Also Negotiates status code Insufficient Storage status code Loop Detected status code Not Extended status code Network Authentication Required status code

#### Bot Defense Policy Protected App Endpoints Flow Label Financial Services

A [`financial_services`](#bot-defense-policy-protected-app-endpoints-flow-label-financial-services) block (within [`bot_defense.policy.protected_app_endpoints.flow_label`](#bot-defense-policy-protected-app-endpoints-flow-label)) supports the following:

<a id="bot-defense-policy-protected-app-endpoints-flow-label-financial-services-apply"></a>&#x2022; [`apply`](#bot-defense-policy-protected-app-endpoints-flow-label-financial-services-apply) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="bot-defense-policy-protected-app-endpoints-flow-label-financial-services-money-transfer"></a>&#x2022; [`money_transfer`](#bot-defense-policy-protected-app-endpoints-flow-label-financial-services-money-transfer) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### Bot Defense Policy Protected App Endpoints Flow Label Flight

A [`flight`](#bot-defense-policy-protected-app-endpoints-flow-label-flight) block (within [`bot_defense.policy.protected_app_endpoints.flow_label`](#bot-defense-policy-protected-app-endpoints-flow-label)) supports the following:

<a id="bot-defense-policy-protected-app-endpoints-flow-label-flight-checkin"></a>&#x2022; [`checkin`](#bot-defense-policy-protected-app-endpoints-flow-label-flight-checkin) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### Bot Defense Policy Protected App Endpoints Flow Label Profile Management

A [`profile_management`](#bot-defense-policy-protected-app-endpoints-flow-label-profile-management) block (within [`bot_defense.policy.protected_app_endpoints.flow_label`](#bot-defense-policy-protected-app-endpoints-flow-label)) supports the following:

<a id="bot-defense-policy-protected-app-endpoints-flow-label-profile-management-create"></a>&#x2022; [`create`](#bot-defense-policy-protected-app-endpoints-flow-label-profile-management-create) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="bot-defense-policy-protected-app-endpoints-flow-label-profile-management-update"></a>&#x2022; [`update`](#bot-defense-policy-protected-app-endpoints-flow-label-profile-management-update) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="bot-defense-policy-protected-app-endpoints-flow-label-profile-management-view"></a>&#x2022; [`view`](#bot-defense-policy-protected-app-endpoints-flow-label-profile-management-view) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### Bot Defense Policy Protected App Endpoints Flow Label Search

A [`search`](#bot-defense-policy-protected-app-endpoints-flow-label-search) block (within [`bot_defense.policy.protected_app_endpoints.flow_label`](#bot-defense-policy-protected-app-endpoints-flow-label)) supports the following:

<a id="bot-defense-policy-protected-app-endpoints-flow-label-search-flight-search"></a>&#x2022; [`flight_search`](#bot-defense-policy-protected-app-endpoints-flow-label-search-flight-search) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="bot-defense-policy-protected-app-endpoints-flow-label-search-product-search"></a>&#x2022; [`product_search`](#bot-defense-policy-protected-app-endpoints-flow-label-search-product-search) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="bot-defense-policy-protected-app-endpoints-flow-label-search-reservation-search"></a>&#x2022; [`reservation_search`](#bot-defense-policy-protected-app-endpoints-flow-label-search-reservation-search) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="bot-defense-policy-protected-app-endpoints-flow-label-search-room-search"></a>&#x2022; [`room_search`](#bot-defense-policy-protected-app-endpoints-flow-label-search-room-search) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### Bot Defense Policy Protected App Endpoints Flow Label Shopping Gift Cards

A [`shopping_gift_cards`](#bot-defense-policy-protected-app-endpoints-flow-label-shopping-gift-cards) block (within [`bot_defense.policy.protected_app_endpoints.flow_label`](#bot-defense-policy-protected-app-endpoints-flow-label)) supports the following:

<a id="bot-defense-policy-protected-app-endpoints-flow-label-shopping-gift-cards-gift-card-make-purchase-with-gift-card"></a>&#x2022; [`gift_card_make_purchase_with_gift_card`](#bot-defense-policy-protected-app-endpoints-flow-label-shopping-gift-cards-gift-card-make-purchase-with-gift-card) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="bot-defense-policy-protected-app-endpoints-flow-label-shopping-gift-cards-gift-card-validation"></a>&#x2022; [`gift_card_validation`](#bot-defense-policy-protected-app-endpoints-flow-label-shopping-gift-cards-gift-card-validation) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="bot-defense-policy-protected-app-endpoints-flow-label-shopping-gift-cards-shop-add-to-cart"></a>&#x2022; [`shop_add_to_cart`](#bot-defense-policy-protected-app-endpoints-flow-label-shopping-gift-cards-shop-add-to-cart) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="bot-defense-policy-protected-app-endpoints-flow-label-shopping-gift-cards-shop-checkout"></a>&#x2022; [`shop_checkout`](#bot-defense-policy-protected-app-endpoints-flow-label-shopping-gift-cards-shop-checkout) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="bot-defense-policy-protected-app-endpoints-flow-label-shopping-gift-cards-shop-choose-seat"></a>&#x2022; [`shop_choose_seat`](#bot-defense-policy-protected-app-endpoints-flow-label-shopping-gift-cards-shop-choose-seat) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="bot-defense-policy-protected-app-endpoints-flow-label-shopping-gift-cards-shop-enter-drawing-submission"></a>&#x2022; [`shop_enter_drawing_submission`](#bot-defense-policy-protected-app-endpoints-flow-label-shopping-gift-cards-shop-enter-drawing-submission) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="bot-defense-policy-protected-app-endpoints-flow-label-shopping-gift-cards-shop-make-payment"></a>&#x2022; [`shop_make_payment`](#bot-defense-policy-protected-app-endpoints-flow-label-shopping-gift-cards-shop-make-payment) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="bot-defense-policy-protected-app-endpoints-flow-label-shopping-gift-cards-shop-order"></a>&#x2022; [`shop_order`](#bot-defense-policy-protected-app-endpoints-flow-label-shopping-gift-cards-shop-order) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="bot-defense-policy-protected-app-endpoints-flow-label-shopping-gift-cards-shop-price-inquiry"></a>&#x2022; [`shop_price_inquiry`](#bot-defense-policy-protected-app-endpoints-flow-label-shopping-gift-cards-shop-price-inquiry) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="bot-defense-policy-protected-app-endpoints-flow-label-shopping-gift-cards-shop-promo-code-validation"></a>&#x2022; [`shop_promo_code_validation`](#bot-defense-policy-protected-app-endpoints-flow-label-shopping-gift-cards-shop-promo-code-validation) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="bot-defense-policy-protected-app-endpoints-flow-label-shopping-gift-cards-shop-purchase-gift-card"></a>&#x2022; [`shop_purchase_gift_card`](#bot-defense-policy-protected-app-endpoints-flow-label-shopping-gift-cards-shop-purchase-gift-card) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="bot-defense-policy-protected-app-endpoints-flow-label-shopping-gift-cards-shop-update-quantity"></a>&#x2022; [`shop_update_quantity`](#bot-defense-policy-protected-app-endpoints-flow-label-shopping-gift-cards-shop-update-quantity) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### Bot Defense Policy Protected App Endpoints Headers

A [`headers`](#bot-defense-policy-protected-app-endpoints-headers) block (within [`bot_defense.policy.protected_app_endpoints`](#bot-defense-policy-protected-app-endpoints)) supports the following:

<a id="bot-defense-policy-protected-app-endpoints-headers-check-not-present"></a>&#x2022; [`check_not_present`](#bot-defense-policy-protected-app-endpoints-headers-check-not-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="bot-defense-policy-protected-app-endpoints-headers-check-present"></a>&#x2022; [`check_present`](#bot-defense-policy-protected-app-endpoints-headers-check-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="bot-defense-policy-protected-app-endpoints-headers-invert-matcher"></a>&#x2022; [`invert_matcher`](#bot-defense-policy-protected-app-endpoints-headers-invert-matcher) - Optional Bool<br>Invert Header Matcher. Invert the match result

<a id="bot-defense-policy-protected-app-endpoints-headers-item"></a>&#x2022; [`item`](#bot-defense-policy-protected-app-endpoints-headers-item) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#bot-defense-policy-protected-app-endpoints-headers-item) below.

<a id="bot-defense-policy-protected-app-endpoints-headers-name"></a>&#x2022; [`name`](#bot-defense-policy-protected-app-endpoints-headers-name) - Optional String<br>Header Name. A case-insensitive HTTP header name

#### Bot Defense Policy Protected App Endpoints Headers Item

An [`item`](#bot-defense-policy-protected-app-endpoints-headers-item) block (within [`bot_defense.policy.protected_app_endpoints.headers`](#bot-defense-policy-protected-app-endpoints-headers)) supports the following:

<a id="bot-defense-policy-protected-app-endpoints-headers-item-exact-values"></a>&#x2022; [`exact_values`](#bot-defense-policy-protected-app-endpoints-headers-item-exact-values) - Optional List<br>Exact Values. A list of exact values to match the input against

<a id="bot-defense-policy-protected-app-endpoints-headers-item-regex-values"></a>&#x2022; [`regex_values`](#bot-defense-policy-protected-app-endpoints-headers-item-regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="bot-defense-policy-protected-app-endpoints-headers-item-transformers"></a>&#x2022; [`transformers`](#bot-defense-policy-protected-app-endpoints-headers-item-transformers) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### Bot Defense Policy Protected App Endpoints Metadata

A [`metadata`](#bot-defense-policy-protected-app-endpoints-metadata) block (within [`bot_defense.policy.protected_app_endpoints`](#bot-defense-policy-protected-app-endpoints)) supports the following:

<a id="bot-defense-policy-protected-app-endpoints-metadata-description-spec"></a>&#x2022; [`description_spec`](#bot-defense-policy-protected-app-endpoints-metadata-description-spec) - Optional String<br>Description. Human readable description

<a id="bot-defense-policy-protected-app-endpoints-metadata-name"></a>&#x2022; [`name`](#bot-defense-policy-protected-app-endpoints-metadata-name) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### Bot Defense Policy Protected App Endpoints Mitigation

A [`mitigation`](#bot-defense-policy-protected-app-endpoints-mitigation) block (within [`bot_defense.policy.protected_app_endpoints`](#bot-defense-policy-protected-app-endpoints)) supports the following:

<a id="bot-defense-policy-protected-app-endpoints-mitigation-block"></a>&#x2022; [`block`](#bot-defense-policy-protected-app-endpoints-mitigation-block) - Optional Block<br>Block bot mitigation. Block request and respond with custom content<br>See [Block](#bot-defense-policy-protected-app-endpoints-mitigation-block) below.

<a id="bot-defense-policy-protected-app-endpoints-mitigation-flag"></a>&#x2022; [`flag`](#bot-defense-policy-protected-app-endpoints-mitigation-flag) - Optional Block<br>Select Flag Bot Mitigation Action. Flag mitigation action<br>See [Flag](#bot-defense-policy-protected-app-endpoints-mitigation-flag) below.

<a id="bot-defense-policy-protected-app-endpoints-mitigation-redirect"></a>&#x2022; [`redirect`](#bot-defense-policy-protected-app-endpoints-mitigation-redirect) - Optional Block<br>Redirect bot mitigation. Redirect request to a custom URI<br>See [Redirect](#bot-defense-policy-protected-app-endpoints-mitigation-redirect) below.

#### Bot Defense Policy Protected App Endpoints Mitigation Block

A [`block`](#bot-defense-policy-protected-app-endpoints-mitigation-block) block (within [`bot_defense.policy.protected_app_endpoints.mitigation`](#bot-defense-policy-protected-app-endpoints-mitigation)) supports the following:

<a id="bot-defense-policy-protected-app-endpoints-mitigation-block-body"></a>&#x2022; [`body`](#bot-defense-policy-protected-app-endpoints-mitigation-block-body) - Optional String<br>Body. Custom body message is of type uri_ref. Currently supported URL schemes is string:///. For string:/// scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. 'Your request was blocked' or it can be HTML paragraph or a body string encoded as base64 string E.g. '`<p>` Your request was blocked `</p>`'. Base64 encoded string for this HTML is 'LzxwPiBZb3VyIHJlcXVlc3Qgd2FzIGJsb2NrZWQgPC9wPg=='

<a id="bot-defense-policy-protected-app-endpoints-mitigation-block-status"></a>&#x2022; [`status`](#bot-defense-policy-protected-app-endpoints-mitigation-block-status) - Optional String  Defaults to `EmptyStatusCode`<br>Possible values are `EmptyStatusCode`, `Continue`, `OK`, `Created`, `Accepted`, `NonAuthoritativeInformation`, `NoContent`, `ResetContent`, `PartialContent`, `MultiStatus`, `AlreadyReported`, `IMUsed`, `MultipleChoices`, `MovedPermanently`, `Found`, `SeeOther`, `NotModified`, `UseProxy`, `TemporaryRedirect`, `PermanentRedirect`, `BadRequest`, `Unauthorized`, `PaymentRequired`, `Forbidden`, `NotFound`, `MethodNotAllowed`, `NotAcceptable`, `ProxyAuthenticationRequired`, `RequestTimeout`, `Conflict`, `Gone`, `LengthRequired`, `PreconditionFailed`, `PayloadTooLarge`, `URITooLong`, `UnsupportedMediaType`, `RangeNotSatisfiable`, `ExpectationFailed`, `MisdirectedRequest`, `UnprocessableEntity`, `Locked`, `FailedDependency`, `UpgradeRequired`, `PreconditionRequired`, `TooManyRequests`, `RequestHeaderFieldsTooLarge`, `InternalServerError`, `NotImplemented`, `BadGateway`, `ServiceUnavailable`, `GatewayTimeout`, `HTTPVersionNotSupported`, `VariantAlsoNegotiates`, `InsufficientStorage`, `LoopDetected`, `NotExtended`, `NetworkAuthenticationRequired`<br>HTTP Status Code. HTTP response status codes EmptyStatusCode response codes means it is not specified Continue status code OK status code Created status code Accepted status code Non Authoritative Information status code No Content status code Reset Content status code Partial Content status code Multi Status status code Already Reported status code Im Used status code Multiple Choices status code Moved Permanently status code Found status code See Other status code Not Modified status code Use Proxy status code Temporary Redirect status code Permanent Redirect status code Bad Request status code Unauthorized status code Payment Required status code Forbidden status code Not Found status code Method Not Allowed status code Not Acceptable status code Proxy Authentication Required status code Request Timeout status code Conflict status code Gone status code Length Required status code Precondition Failed status code Payload Too Large status code URI Too Long status code Unsupported Media Type status code Range Not Satisfiable status code Expectation Failed status code Misdirected Request status code Unprocessable Entity status code Locked status code Failed Dependency status code Upgrade Required status code Precondition Required status code Too Many Requests status code Request Header Fields Too Large status code Internal Server Error status code Not Implemented status code Bad Gateway status code Service Unavailable status code Gateway Timeout status code HTTP Version Not Supported status code Variant Also Negotiates status code Insufficient Storage status code Loop Detected status code Not Extended status code Network Authentication Required status code

#### Bot Defense Policy Protected App Endpoints Mitigation Flag

A [`flag`](#bot-defense-policy-protected-app-endpoints-mitigation-flag) block (within [`bot_defense.policy.protected_app_endpoints.mitigation`](#bot-defense-policy-protected-app-endpoints-mitigation)) supports the following:

<a id="bot-defense-policy-protected-app-endpoints-mitigation-flag-append-headers"></a>&#x2022; [`append_headers`](#bot-defense-policy-protected-app-endpoints-mitigation-flag-append-headers) - Optional Block<br>Append Flag Mitigation Headers. Append flag mitigation headers to forwarded request<br>See [Append Headers](#bot-defense-policy-protected-app-endpoints-mitigation-flag-append-headers) below.

<a id="bot-defense-policy-protected-app-endpoints-mitigation-flag-no-headers"></a>&#x2022; [`no_headers`](#bot-defense-policy-protected-app-endpoints-mitigation-flag-no-headers) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### Bot Defense Policy Protected App Endpoints Mitigation Flag Append Headers

An [`append_headers`](#bot-defense-policy-protected-app-endpoints-mitigation-flag-append-headers) block (within [`bot_defense.policy.protected_app_endpoints.mitigation.flag`](#bot-defense-policy-protected-app-endpoints-mitigation-flag)) supports the following:

<a id="bot-defense-policy-protected-app-endpoints-mitigation-flag-append-headers-auto-type-header-name"></a>&#x2022; [`auto_type_header_name`](#bot-defense-policy-protected-app-endpoints-mitigation-flag-append-headers-auto-type-header-name) - Optional String<br>Automation Type Header Name. A case-insensitive HTTP header name

<a id="bot-defense-policy-protected-app-endpoints-mitigation-flag-append-headers-inference-header-name"></a>&#x2022; [`inference_header_name`](#bot-defense-policy-protected-app-endpoints-mitigation-flag-append-headers-inference-header-name) - Optional String<br>Inference Header Name. A case-insensitive HTTP header name

#### Bot Defense Policy Protected App Endpoints Mitigation Redirect

A [`redirect`](#bot-defense-policy-protected-app-endpoints-mitigation-redirect) block (within [`bot_defense.policy.protected_app_endpoints.mitigation`](#bot-defense-policy-protected-app-endpoints-mitigation)) supports the following:

<a id="bot-defense-policy-protected-app-endpoints-mitigation-redirect-uri"></a>&#x2022; [`uri`](#bot-defense-policy-protected-app-endpoints-mitigation-redirect-uri) - Optional String<br>URI. URI location for redirect may be relative or absolute

#### Bot Defense Policy Protected App Endpoints Path

A [`path`](#bot-defense-policy-protected-app-endpoints-path) block (within [`bot_defense.policy.protected_app_endpoints`](#bot-defense-policy-protected-app-endpoints)) supports the following:

<a id="bot-defense-policy-protected-app-endpoints-path-path"></a>&#x2022; [`path`](#bot-defense-policy-protected-app-endpoints-path-path) - Optional String<br>Exact. Exact path value to match

<a id="bot-defense-policy-protected-app-endpoints-path-prefix"></a>&#x2022; [`prefix`](#bot-defense-policy-protected-app-endpoints-path-prefix) - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

<a id="bot-defense-policy-protected-app-endpoints-path-regex"></a>&#x2022; [`regex`](#bot-defense-policy-protected-app-endpoints-path-regex) - Optional String<br>Regex. Regular expression of path match (e.g. the value .* will match on all paths)

#### Bot Defense Policy Protected App Endpoints Query Params

A [`query_params`](#bot-defense-policy-protected-app-endpoints-query-params) block (within [`bot_defense.policy.protected_app_endpoints`](#bot-defense-policy-protected-app-endpoints)) supports the following:

<a id="bot-defense-policy-protected-app-endpoints-query-params-check-not-present"></a>&#x2022; [`check_not_present`](#bot-defense-policy-protected-app-endpoints-query-params-check-not-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="bot-defense-policy-protected-app-endpoints-query-params-check-present"></a>&#x2022; [`check_present`](#bot-defense-policy-protected-app-endpoints-query-params-check-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="bot-defense-policy-protected-app-endpoints-query-params-invert-matcher"></a>&#x2022; [`invert_matcher`](#bot-defense-policy-protected-app-endpoints-query-params-invert-matcher) - Optional Bool<br>Invert Query Parameter Matcher. Invert the match result

<a id="bot-defense-policy-protected-app-endpoints-query-params-item"></a>&#x2022; [`item`](#bot-defense-policy-protected-app-endpoints-query-params-item) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#bot-defense-policy-protected-app-endpoints-query-params-item) below.

<a id="bot-defense-policy-protected-app-endpoints-query-params-key"></a>&#x2022; [`key`](#bot-defense-policy-protected-app-endpoints-query-params-key) - Optional String<br>Query Parameter Name. A case-sensitive HTTP query parameter name

#### Bot Defense Policy Protected App Endpoints Query Params Item

An [`item`](#bot-defense-policy-protected-app-endpoints-query-params-item) block (within [`bot_defense.policy.protected_app_endpoints.query_params`](#bot-defense-policy-protected-app-endpoints-query-params)) supports the following:

<a id="bot-defense-policy-protected-app-endpoints-query-params-item-exact-values"></a>&#x2022; [`exact_values`](#bot-defense-policy-protected-app-endpoints-query-params-item-exact-values) - Optional List<br>Exact Values. A list of exact values to match the input against

<a id="bot-defense-policy-protected-app-endpoints-query-params-item-regex-values"></a>&#x2022; [`regex_values`](#bot-defense-policy-protected-app-endpoints-query-params-item-regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="bot-defense-policy-protected-app-endpoints-query-params-item-transformers"></a>&#x2022; [`transformers`](#bot-defense-policy-protected-app-endpoints-query-params-item-transformers) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### Bot Defense Policy Protected App Endpoints Web Mobile

A [`web_mobile`](#bot-defense-policy-protected-app-endpoints-web-mobile) block (within [`bot_defense.policy.protected_app_endpoints`](#bot-defense-policy-protected-app-endpoints)) supports the following:

<a id="bot-defense-policy-protected-app-endpoints-web-mobile-mobile-identifier"></a>&#x2022; [`mobile_identifier`](#bot-defense-policy-protected-app-endpoints-web-mobile-mobile-identifier) - Optional String  Defaults to `HEADERS`<br>Mobile Identifier. Mobile identifier type - HEADERS: Headers Headers. The only possible value is `HEADERS`

#### Bot Defense Advanced

A [`bot_defense_advanced`](#bot-defense-advanced) block supports the following:

<a id="bot-defense-advanced-disable-js-insert"></a>&#x2022; [`disable_js_insert`](#bot-defense-advanced-disable-js-insert) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="bot-defense-advanced-disable-mobile-sdk"></a>&#x2022; [`disable_mobile_sdk`](#bot-defense-advanced-disable-mobile-sdk) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="bot-defense-advanced-js-insert-all-pages"></a>&#x2022; [`js_insert_all_pages`](#bot-defense-advanced-js-insert-all-pages) - Optional Block<br>Insert Bot Defense JavaScript in All Pages. Insert Bot Defense JavaScript in all pages<br>See [Js Insert All Pages](#bot-defense-advanced-js-insert-all-pages) below.

<a id="bot-defense-advanced-js-insert-all-pages-except"></a>&#x2022; [`js_insert_all_pages_except`](#bot-defense-advanced-js-insert-all-pages-except) - Optional Block<br>Insert JavaScript in All Pages with the Exceptions. Insert Bot Defense JavaScript in all pages with the exceptions<br>See [Js Insert All Pages Except](#bot-defense-advanced-js-insert-all-pages-except) below.

<a id="bot-defense-advanced-js-insertion-rules"></a>&#x2022; [`js_insertion_rules`](#bot-defense-advanced-js-insertion-rules) - Optional Block<br>JavaScript Custom Insertion Rules. This defines custom JavaScript insertion rules for Bot Defense Policy<br>See [Js Insertion Rules](#bot-defense-advanced-js-insertion-rules) below.

<a id="bot-defense-advanced-mobile"></a>&#x2022; [`mobile`](#bot-defense-advanced-mobile) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Mobile](#bot-defense-advanced-mobile) below.

<a id="bot-defense-advanced-mobile-sdk-config"></a>&#x2022; [`mobile_sdk_config`](#bot-defense-advanced-mobile-sdk-config) - Optional Block<br>Mobile Request Identifier Headers. Mobile Request Identifier Headers<br>See [Mobile Sdk Config](#bot-defense-advanced-mobile-sdk-config) below.

<a id="bot-defense-advanced-web"></a>&#x2022; [`web`](#bot-defense-advanced-web) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Web](#bot-defense-advanced-web) below.

#### Bot Defense Advanced Js Insert All Pages

A [`js_insert_all_pages`](#bot-defense-advanced-js-insert-all-pages) block (within [`bot_defense_advanced`](#bot-defense-advanced)) supports the following:

<a id="bot-defense-advanced-js-insert-all-pages-javascript-location"></a>&#x2022; [`javascript_location`](#bot-defense-advanced-js-insert-all-pages-javascript-location) - Optional String  Defaults to `AFTER_HEAD`<br>Possible values are `AFTER_HEAD`, `AFTER_TITLE_END`, `BEFORE_SCRIPT`<br>JavaScript Location. All inside networks. Insert JavaScript after `<head>` tag Insert JavaScript after `</title>` tag. Insert JavaScript before first `<script>` tag

#### Bot Defense Advanced Js Insert All Pages Except

A [`js_insert_all_pages_except`](#bot-defense-advanced-js-insert-all-pages-except) block (within [`bot_defense_advanced`](#bot-defense-advanced)) supports the following:

<a id="bot-defense-advanced-js-insert-all-pages-except-exclude-list"></a>&#x2022; [`exclude_list`](#bot-defense-advanced-js-insert-all-pages-except-exclude-list) - Optional Block<br>Exclude Pages. Optional JavaScript insertions exclude list of domain and path matchers<br>See [Exclude List](#bot-defense-advanced-js-insert-all-pages-except-exclude-list) below.

<a id="bot-defense-advanced-js-insert-all-pages-except-javascript-location"></a>&#x2022; [`javascript_location`](#bot-defense-advanced-js-insert-all-pages-except-javascript-location) - Optional String  Defaults to `AFTER_HEAD`<br>Possible values are `AFTER_HEAD`, `AFTER_TITLE_END`, `BEFORE_SCRIPT`<br>JavaScript Location. All inside networks. Insert JavaScript after `<head>` tag Insert JavaScript after `</title>` tag. Insert JavaScript before first `<script>` tag

#### Bot Defense Advanced Js Insert All Pages Except Exclude List

An [`exclude_list`](#bot-defense-advanced-js-insert-all-pages-except-exclude-list) block (within [`bot_defense_advanced.js_insert_all_pages_except`](#bot-defense-advanced-js-insert-all-pages-except)) supports the following:

<a id="bot-defense-advanced-js-insert-all-pages-except-exclude-list-any-domain"></a>&#x2022; [`any_domain`](#bot-defense-advanced-js-insert-all-pages-except-exclude-list-any-domain) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="bot-defense-advanced-js-insert-all-pages-except-exclude-list-domain"></a>&#x2022; [`domain`](#bot-defense-advanced-js-insert-all-pages-except-exclude-list-domain) - Optional Block<br>Domains. Domains names<br>See [Domain](#bot-defense-advanced-js-insert-all-pages-except-exclude-list-domain) below.

<a id="bot-defense-advanced-js-insert-all-pages-except-exclude-list-metadata"></a>&#x2022; [`metadata`](#bot-defense-advanced-js-insert-all-pages-except-exclude-list-metadata) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#bot-defense-advanced-js-insert-all-pages-except-exclude-list-metadata) below.

<a id="bot-defense-advanced-js-insert-all-pages-except-exclude-list-path"></a>&#x2022; [`path`](#bot-defense-advanced-js-insert-all-pages-except-exclude-list-path) - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#bot-defense-advanced-js-insert-all-pages-except-exclude-list-path) below.

#### Bot Defense Advanced Js Insert All Pages Except Exclude List Domain

A [`domain`](#bot-defense-advanced-js-insert-all-pages-except-exclude-list-domain) block (within [`bot_defense_advanced.js_insert_all_pages_except.exclude_list`](#bot-defense-advanced-js-insert-all-pages-except-exclude-list)) supports the following:

<a id="bot-defense-advanced-js-insert-all-pages-except-exclude-list-domain-exact-value"></a>&#x2022; [`exact_value`](#bot-defense-advanced-js-insert-all-pages-except-exclude-list-domain-exact-value) - Optional String<br>Exact Value. Exact domain name

<a id="bot-defense-advanced-js-insert-all-pages-except-exclude-list-domain-regex-value"></a>&#x2022; [`regex_value`](#bot-defense-advanced-js-insert-all-pages-except-exclude-list-domain-regex-value) - Optional String<br>Regex Values of Domains. Regular Expression value for the domain name

<a id="bot-defense-advanced-js-insert-all-pages-except-exclude-list-domain-suffix-value"></a>&#x2022; [`suffix_value`](#bot-defense-advanced-js-insert-all-pages-except-exclude-list-domain-suffix-value) - Optional String<br>Suffix Value. Suffix of domain name e.g 'xyz.com' will match '*.xyz.com' and 'xyz.com'

#### Bot Defense Advanced Js Insert All Pages Except Exclude List Metadata

A [`metadata`](#bot-defense-advanced-js-insert-all-pages-except-exclude-list-metadata) block (within [`bot_defense_advanced.js_insert_all_pages_except.exclude_list`](#bot-defense-advanced-js-insert-all-pages-except-exclude-list)) supports the following:

<a id="bot-defense-advanced-js-insert-all-pages-except-exclude-list-metadata-description-spec"></a>&#x2022; [`description_spec`](#bot-defense-advanced-js-insert-all-pages-except-exclude-list-metadata-description-spec) - Optional String<br>Description. Human readable description

<a id="bot-defense-advanced-js-insert-all-pages-except-exclude-list-metadata-name"></a>&#x2022; [`name`](#bot-defense-advanced-js-insert-all-pages-except-exclude-list-metadata-name) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### Bot Defense Advanced Js Insert All Pages Except Exclude List Path

A [`path`](#bot-defense-advanced-js-insert-all-pages-except-exclude-list-path) block (within [`bot_defense_advanced.js_insert_all_pages_except.exclude_list`](#bot-defense-advanced-js-insert-all-pages-except-exclude-list)) supports the following:

<a id="bot-defense-advanced-js-insert-all-pages-except-exclude-list-path-path"></a>&#x2022; [`path`](#bot-defense-advanced-js-insert-all-pages-except-exclude-list-path-path) - Optional String<br>Exact. Exact path value to match

<a id="bot-defense-advanced-js-insert-all-pages-except-exclude-list-path-prefix"></a>&#x2022; [`prefix`](#bot-defense-advanced-js-insert-all-pages-except-exclude-list-path-prefix) - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

<a id="bot-defense-advanced-js-insert-all-pages-except-exclude-list-path-regex"></a>&#x2022; [`regex`](#bot-defense-advanced-js-insert-all-pages-except-exclude-list-path-regex) - Optional String<br>Regex. Regular expression of path match (e.g. the value .* will match on all paths)

#### Bot Defense Advanced Js Insertion Rules

A [`js_insertion_rules`](#bot-defense-advanced-js-insertion-rules) block (within [`bot_defense_advanced`](#bot-defense-advanced)) supports the following:

<a id="bot-defense-advanced-js-insertion-rules-exclude-list"></a>&#x2022; [`exclude_list`](#bot-defense-advanced-js-insertion-rules-exclude-list) - Optional Block<br>Exclude Paths. Optional JavaScript insertions exclude list of domain and path matchers<br>See [Exclude List](#bot-defense-advanced-js-insertion-rules-exclude-list) below.

<a id="bot-defense-advanced-js-insertion-rules-rules"></a>&#x2022; [`rules`](#bot-defense-advanced-js-insertion-rules-rules) - Optional Block<br>JavaScript Insertions. Required list of pages to insert Bot Defense client JavaScript<br>See [Rules](#bot-defense-advanced-js-insertion-rules-rules) below.

#### Bot Defense Advanced Js Insertion Rules Exclude List

An [`exclude_list`](#bot-defense-advanced-js-insertion-rules-exclude-list) block (within [`bot_defense_advanced.js_insertion_rules`](#bot-defense-advanced-js-insertion-rules)) supports the following:

<a id="bot-defense-advanced-js-insertion-rules-exclude-list-any-domain"></a>&#x2022; [`any_domain`](#bot-defense-advanced-js-insertion-rules-exclude-list-any-domain) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="bot-defense-advanced-js-insertion-rules-exclude-list-domain"></a>&#x2022; [`domain`](#bot-defense-advanced-js-insertion-rules-exclude-list-domain) - Optional Block<br>Domains. Domains names<br>See [Domain](#bot-defense-advanced-js-insertion-rules-exclude-list-domain) below.

<a id="bot-defense-advanced-js-insertion-rules-exclude-list-metadata"></a>&#x2022; [`metadata`](#bot-defense-advanced-js-insertion-rules-exclude-list-metadata) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#bot-defense-advanced-js-insertion-rules-exclude-list-metadata) below.

<a id="bot-defense-advanced-js-insertion-rules-exclude-list-path"></a>&#x2022; [`path`](#bot-defense-advanced-js-insertion-rules-exclude-list-path) - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#bot-defense-advanced-js-insertion-rules-exclude-list-path) below.

#### Bot Defense Advanced Js Insertion Rules Exclude List Domain

A [`domain`](#bot-defense-advanced-js-insertion-rules-exclude-list-domain) block (within [`bot_defense_advanced.js_insertion_rules.exclude_list`](#bot-defense-advanced-js-insertion-rules-exclude-list)) supports the following:

<a id="bot-defense-advanced-js-insertion-rules-exclude-list-domain-exact-value"></a>&#x2022; [`exact_value`](#bot-defense-advanced-js-insertion-rules-exclude-list-domain-exact-value) - Optional String<br>Exact Value. Exact domain name

<a id="bot-defense-advanced-js-insertion-rules-exclude-list-domain-regex-value"></a>&#x2022; [`regex_value`](#bot-defense-advanced-js-insertion-rules-exclude-list-domain-regex-value) - Optional String<br>Regex Values of Domains. Regular Expression value for the domain name

<a id="bot-defense-advanced-js-insertion-rules-exclude-list-domain-suffix-value"></a>&#x2022; [`suffix_value`](#bot-defense-advanced-js-insertion-rules-exclude-list-domain-suffix-value) - Optional String<br>Suffix Value. Suffix of domain name e.g 'xyz.com' will match '*.xyz.com' and 'xyz.com'

#### Bot Defense Advanced Js Insertion Rules Exclude List Metadata

A [`metadata`](#bot-defense-advanced-js-insertion-rules-exclude-list-metadata) block (within [`bot_defense_advanced.js_insertion_rules.exclude_list`](#bot-defense-advanced-js-insertion-rules-exclude-list)) supports the following:

<a id="bot-defense-advanced-js-insertion-rules-exclude-list-metadata-description-spec"></a>&#x2022; [`description_spec`](#bot-defense-advanced-js-insertion-rules-exclude-list-metadata-description-spec) - Optional String<br>Description. Human readable description

<a id="bot-defense-advanced-js-insertion-rules-exclude-list-metadata-name"></a>&#x2022; [`name`](#bot-defense-advanced-js-insertion-rules-exclude-list-metadata-name) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### Bot Defense Advanced Js Insertion Rules Exclude List Path

A [`path`](#bot-defense-advanced-js-insertion-rules-exclude-list-path) block (within [`bot_defense_advanced.js_insertion_rules.exclude_list`](#bot-defense-advanced-js-insertion-rules-exclude-list)) supports the following:

<a id="bot-defense-advanced-js-insertion-rules-exclude-list-path-path"></a>&#x2022; [`path`](#bot-defense-advanced-js-insertion-rules-exclude-list-path-path) - Optional String<br>Exact. Exact path value to match

<a id="bot-defense-advanced-js-insertion-rules-exclude-list-path-prefix"></a>&#x2022; [`prefix`](#bot-defense-advanced-js-insertion-rules-exclude-list-path-prefix) - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

<a id="bot-defense-advanced-js-insertion-rules-exclude-list-path-regex"></a>&#x2022; [`regex`](#bot-defense-advanced-js-insertion-rules-exclude-list-path-regex) - Optional String<br>Regex. Regular expression of path match (e.g. the value .* will match on all paths)

#### Bot Defense Advanced Js Insertion Rules Rules

A [`rules`](#bot-defense-advanced-js-insertion-rules-rules) block (within [`bot_defense_advanced.js_insertion_rules`](#bot-defense-advanced-js-insertion-rules)) supports the following:

<a id="bot-defense-advanced-js-insertion-rules-rules-any-domain"></a>&#x2022; [`any_domain`](#bot-defense-advanced-js-insertion-rules-rules-any-domain) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="bot-defense-advanced-js-insertion-rules-rules-domain"></a>&#x2022; [`domain`](#bot-defense-advanced-js-insertion-rules-rules-domain) - Optional Block<br>Domains. Domains names<br>See [Domain](#bot-defense-advanced-js-insertion-rules-rules-domain) below.

<a id="bot-defense-advanced-js-insertion-rules-rules-javascript-location"></a>&#x2022; [`javascript_location`](#bot-defense-advanced-js-insertion-rules-rules-javascript-location) - Optional String  Defaults to `AFTER_HEAD`<br>Possible values are `AFTER_HEAD`, `AFTER_TITLE_END`, `BEFORE_SCRIPT`<br>JavaScript Location. All inside networks. Insert JavaScript after `<head>` tag Insert JavaScript after `</title>` tag. Insert JavaScript before first `<script>` tag

<a id="bot-defense-advanced-js-insertion-rules-rules-metadata"></a>&#x2022; [`metadata`](#bot-defense-advanced-js-insertion-rules-rules-metadata) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#bot-defense-advanced-js-insertion-rules-rules-metadata) below.

<a id="bot-defense-advanced-js-insertion-rules-rules-path"></a>&#x2022; [`path`](#bot-defense-advanced-js-insertion-rules-rules-path) - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#bot-defense-advanced-js-insertion-rules-rules-path) below.

#### Bot Defense Advanced Js Insertion Rules Rules Domain

A [`domain`](#bot-defense-advanced-js-insertion-rules-rules-domain) block (within [`bot_defense_advanced.js_insertion_rules.rules`](#bot-defense-advanced-js-insertion-rules-rules)) supports the following:

<a id="bot-defense-advanced-js-insertion-rules-rules-domain-exact-value"></a>&#x2022; [`exact_value`](#bot-defense-advanced-js-insertion-rules-rules-domain-exact-value) - Optional String<br>Exact Value. Exact domain name

<a id="bot-defense-advanced-js-insertion-rules-rules-domain-regex-value"></a>&#x2022; [`regex_value`](#bot-defense-advanced-js-insertion-rules-rules-domain-regex-value) - Optional String<br>Regex Values of Domains. Regular Expression value for the domain name

<a id="bot-defense-advanced-js-insertion-rules-rules-domain-suffix-value"></a>&#x2022; [`suffix_value`](#bot-defense-advanced-js-insertion-rules-rules-domain-suffix-value) - Optional String<br>Suffix Value. Suffix of domain name e.g 'xyz.com' will match '*.xyz.com' and 'xyz.com'

#### Bot Defense Advanced Js Insertion Rules Rules Metadata

A [`metadata`](#bot-defense-advanced-js-insertion-rules-rules-metadata) block (within [`bot_defense_advanced.js_insertion_rules.rules`](#bot-defense-advanced-js-insertion-rules-rules)) supports the following:

<a id="bot-defense-advanced-js-insertion-rules-rules-metadata-description-spec"></a>&#x2022; [`description_spec`](#bot-defense-advanced-js-insertion-rules-rules-metadata-description-spec) - Optional String<br>Description. Human readable description

<a id="bot-defense-advanced-js-insertion-rules-rules-metadata-name"></a>&#x2022; [`name`](#bot-defense-advanced-js-insertion-rules-rules-metadata-name) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### Bot Defense Advanced Js Insertion Rules Rules Path

A [`path`](#bot-defense-advanced-js-insertion-rules-rules-path) block (within [`bot_defense_advanced.js_insertion_rules.rules`](#bot-defense-advanced-js-insertion-rules-rules)) supports the following:

<a id="bot-defense-advanced-js-insertion-rules-rules-path-path"></a>&#x2022; [`path`](#bot-defense-advanced-js-insertion-rules-rules-path-path) - Optional String<br>Exact. Exact path value to match

<a id="bot-defense-advanced-js-insertion-rules-rules-path-prefix"></a>&#x2022; [`prefix`](#bot-defense-advanced-js-insertion-rules-rules-path-prefix) - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

<a id="bot-defense-advanced-js-insertion-rules-rules-path-regex"></a>&#x2022; [`regex`](#bot-defense-advanced-js-insertion-rules-rules-path-regex) - Optional String<br>Regex. Regular expression of path match (e.g. the value .* will match on all paths)

#### Bot Defense Advanced Mobile

A [`mobile`](#bot-defense-advanced-mobile) block (within [`bot_defense_advanced`](#bot-defense-advanced)) supports the following:

<a id="bot-defense-advanced-mobile-name"></a>&#x2022; [`name`](#bot-defense-advanced-mobile-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="bot-defense-advanced-mobile-namespace"></a>&#x2022; [`namespace`](#bot-defense-advanced-mobile-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="bot-defense-advanced-mobile-tenant"></a>&#x2022; [`tenant`](#bot-defense-advanced-mobile-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Bot Defense Advanced Mobile Sdk Config

A [`mobile_sdk_config`](#bot-defense-advanced-mobile-sdk-config) block (within [`bot_defense_advanced`](#bot-defense-advanced)) supports the following:

<a id="bot-defense-advanced-mobile-sdk-config-mobile-identifier"></a>&#x2022; [`mobile_identifier`](#bot-defense-advanced-mobile-sdk-config-mobile-identifier) - Optional Block<br>Mobile Traffic Identifier. Mobile traffic identifier type<br>See [Mobile Identifier](#bot-defense-advanced-mobile-sdk-config-mobile-identifier) below.

#### Bot Defense Advanced Mobile Sdk Config Mobile Identifier

A [`mobile_identifier`](#bot-defense-advanced-mobile-sdk-config-mobile-identifier) block (within [`bot_defense_advanced.mobile_sdk_config`](#bot-defense-advanced-mobile-sdk-config)) supports the following:

<a id="bot-defense-advanced-mobile-sdk-config-mobile-identifier-headers"></a>&#x2022; [`headers`](#bot-defense-advanced-mobile-sdk-config-mobile-identifier-headers) - Optional Block<br>Headers. Headers that can be used to identify mobile traffic<br>See [Headers](#bot-defense-advanced-mobile-sdk-config-mobile-identifier-headers) below.

#### Bot Defense Advanced Mobile Sdk Config Mobile Identifier Headers

A [`headers`](#bot-defense-advanced-mobile-sdk-config-mobile-identifier-headers) block (within [`bot_defense_advanced.mobile_sdk_config.mobile_identifier`](#bot-defense-advanced-mobile-sdk-config-mobile-identifier)) supports the following:

<a id="bot-defense-advanced-mobile-sdk-config-mobile-identifier-headers-check-not-present"></a>&#x2022; [`check_not_present`](#bot-defense-advanced-mobile-sdk-config-mobile-identifier-headers-check-not-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="bot-defense-advanced-mobile-sdk-config-mobile-identifier-headers-check-present"></a>&#x2022; [`check_present`](#bot-defense-advanced-mobile-sdk-config-mobile-identifier-headers-check-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="bot-defense-advanced-mobile-sdk-config-mobile-identifier-headers-item"></a>&#x2022; [`item`](#bot-defense-advanced-mobile-sdk-config-mobile-identifier-headers-item) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#bot-defense-advanced-mobile-sdk-config-mobile-identifier-headers-item) below.

<a id="bot-defense-advanced-mobile-sdk-config-mobile-identifier-headers-name"></a>&#x2022; [`name`](#bot-defense-advanced-mobile-sdk-config-mobile-identifier-headers-name) - Optional String<br>Header Name. A case-insensitive HTTP header name

#### Bot Defense Advanced Mobile Sdk Config Mobile Identifier Headers Item

An [`item`](#bot-defense-advanced-mobile-sdk-config-mobile-identifier-headers-item) block (within [`bot_defense_advanced.mobile_sdk_config.mobile_identifier.headers`](#bot-defense-advanced-mobile-sdk-config-mobile-identifier-headers)) supports the following:

<a id="bot-defense-advanced-mobile-sdk-config-mobile-identifier-headers-item-exact-values"></a>&#x2022; [`exact_values`](#bot-defense-advanced-mobile-sdk-config-mobile-identifier-headers-item-exact-values) - Optional List<br>Exact Values. A list of exact values to match the input against

<a id="bot-defense-advanced-mobile-sdk-config-mobile-identifier-headers-item-regex-values"></a>&#x2022; [`regex_values`](#bot-defense-advanced-mobile-sdk-config-mobile-identifier-headers-item-regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="bot-defense-advanced-mobile-sdk-config-mobile-identifier-headers-item-transformers"></a>&#x2022; [`transformers`](#bot-defense-advanced-mobile-sdk-config-mobile-identifier-headers-item-transformers) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

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

<a id="caching-policy-custom-cache-rule-cdn-cache-rules"></a>&#x2022; [`cdn_cache_rules`](#caching-policy-custom-cache-rule-cdn-cache-rules) - Optional Block<br>CDN Cache Rule. Reference to CDN Cache Rule configuration object<br>See [CDN Cache Rules](#caching-policy-custom-cache-rule-cdn-cache-rules) below.

#### Caching Policy Custom Cache Rule CDN Cache Rules

A [`cdn_cache_rules`](#caching-policy-custom-cache-rule-cdn-cache-rules) block (within [`caching_policy.custom_cache_rule`](#caching-policy-custom-cache-rule)) supports the following:

<a id="caching-policy-custom-cache-rule-cdn-cache-rules-name"></a>&#x2022; [`name`](#caching-policy-custom-cache-rule-cdn-cache-rules-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="caching-policy-custom-cache-rule-cdn-cache-rules-namespace"></a>&#x2022; [`namespace`](#caching-policy-custom-cache-rule-cdn-cache-rules-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="caching-policy-custom-cache-rule-cdn-cache-rules-tenant"></a>&#x2022; [`tenant`](#caching-policy-custom-cache-rule-cdn-cache-rules-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Caching Policy Default Cache Action

A [`default_cache_action`](#caching-policy-default-cache-action) block (within [`caching_policy`](#caching-policy)) supports the following:

<a id="caching-policy-default-cache-action-cache-disabled"></a>&#x2022; [`cache_disabled`](#caching-policy-default-cache-action-cache-disabled) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="caching-policy-default-cache-action-cache-ttl-default"></a>&#x2022; [`cache_ttl_default`](#caching-policy-default-cache-action-cache-ttl-default) - Optional String<br>Fallback Cache TTL (d/ h/ m). Use Cache TTL Provided by Origin, and set a contigency TTL value in case one is not provided

<a id="caching-policy-default-cache-action-cache-ttl-override"></a>&#x2022; [`cache_ttl_override`](#caching-policy-default-cache-action-cache-ttl-override) - Optional String<br>Override Cache TTL (d/ h/ m/ s). Always override the Cahce TTL provided by Origin

#### Captcha Challenge

A [`captcha_challenge`](#captcha-challenge) block supports the following:

<a id="captcha-challenge-cookie-expiry"></a>&#x2022; [`cookie_expiry`](#captcha-challenge-cookie-expiry) - Optional Number<br>Cookie Expiration Period. Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new challenge

<a id="captcha-challenge-custom-page"></a>&#x2022; [`custom_page`](#captcha-challenge-custom-page) - Optional String<br>Custom message for Captcha Challenge. Custom message is of type uri_ref. Currently supported URL schemes is string:///. For string:/// scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. 'Please Wait.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '`<p>` Please Wait `</p>`'. Base64 encoded string for this HTML is 'PHA+IFBsZWFzZSBXYWl0IDwvcD4='

#### Client Side Defense

A [`client_side_defense`](#client-side-defense) block supports the following:

<a id="client-side-defense-policy"></a>&#x2022; [`policy`](#client-side-defense-policy) - Optional Block<br>Client-Side Defense Policy. This defines various configuration options for Client-Side Defense policy<br>See [Policy](#client-side-defense-policy) below.

#### Client Side Defense Policy

A [`policy`](#client-side-defense-policy) block (within [`client_side_defense`](#client-side-defense)) supports the following:

<a id="client-side-defense-policy-disable-js-insert"></a>&#x2022; [`disable_js_insert`](#client-side-defense-policy-disable-js-insert) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="client-side-defense-policy-js-insert-all-pages"></a>&#x2022; [`js_insert_all_pages`](#client-side-defense-policy-js-insert-all-pages) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="client-side-defense-policy-js-insert-all-pages-except"></a>&#x2022; [`js_insert_all_pages_except`](#client-side-defense-policy-js-insert-all-pages-except) - Optional Block<br>Insert JavaScript in All Pages with the Exceptions. Insert Client-Side Defense JavaScript in all pages with the exceptions<br>See [Js Insert All Pages Except](#client-side-defense-policy-js-insert-all-pages-except) below.

<a id="client-side-defense-policy-js-insertion-rules"></a>&#x2022; [`js_insertion_rules`](#client-side-defense-policy-js-insertion-rules) - Optional Block<br>JavaScript Custom Insertion Rules. This defines custom JavaScript insertion rules for Client-Side Defense Policy<br>See [Js Insertion Rules](#client-side-defense-policy-js-insertion-rules) below.

#### Client Side Defense Policy Js Insert All Pages Except

A [`js_insert_all_pages_except`](#client-side-defense-policy-js-insert-all-pages-except) block (within [`client_side_defense.policy`](#client-side-defense-policy)) supports the following:

<a id="client-side-defense-policy-js-insert-all-pages-except-exclude-list"></a>&#x2022; [`exclude_list`](#client-side-defense-policy-js-insert-all-pages-except-exclude-list) - Optional Block<br>Exclude Pages. Optional JavaScript insertions exclude list of domain and path matchers<br>See [Exclude List](#client-side-defense-policy-js-insert-all-pages-except-exclude-list) below.

#### Client Side Defense Policy Js Insert All Pages Except Exclude List

An [`exclude_list`](#client-side-defense-policy-js-insert-all-pages-except-exclude-list) block (within [`client_side_defense.policy.js_insert_all_pages_except`](#client-side-defense-policy-js-insert-all-pages-except)) supports the following:

<a id="client-side-defense-policy-js-insert-all-pages-except-exclude-list-any-domain"></a>&#x2022; [`any_domain`](#client-side-defense-policy-js-insert-all-pages-except-exclude-list-any-domain) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="client-side-defense-policy-js-insert-all-pages-except-exclude-list-domain"></a>&#x2022; [`domain`](#client-side-defense-policy-js-insert-all-pages-except-exclude-list-domain) - Optional Block<br>Domains. Domains names<br>See [Domain](#client-side-defense-policy-js-insert-all-pages-except-exclude-list-domain) below.

<a id="client-side-defense-policy-js-insert-all-pages-except-exclude-list-metadata"></a>&#x2022; [`metadata`](#client-side-defense-policy-js-insert-all-pages-except-exclude-list-metadata) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#client-side-defense-policy-js-insert-all-pages-except-exclude-list-metadata) below.

<a id="client-side-defense-policy-js-insert-all-pages-except-exclude-list-path"></a>&#x2022; [`path`](#client-side-defense-policy-js-insert-all-pages-except-exclude-list-path) - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#client-side-defense-policy-js-insert-all-pages-except-exclude-list-path) below.

#### Client Side Defense Policy Js Insert All Pages Except Exclude List Domain

A [`domain`](#client-side-defense-policy-js-insert-all-pages-except-exclude-list-domain) block (within [`client_side_defense.policy.js_insert_all_pages_except.exclude_list`](#client-side-defense-policy-js-insert-all-pages-except-exclude-list)) supports the following:

<a id="client-side-defense-policy-js-insert-all-pages-except-exclude-list-domain-exact-value"></a>&#x2022; [`exact_value`](#client-side-defense-policy-js-insert-all-pages-except-exclude-list-domain-exact-value) - Optional String<br>Exact Value. Exact domain name

<a id="client-side-defense-policy-js-insert-all-pages-except-exclude-list-domain-regex-value"></a>&#x2022; [`regex_value`](#client-side-defense-policy-js-insert-all-pages-except-exclude-list-domain-regex-value) - Optional String<br>Regex Values of Domains. Regular Expression value for the domain name

<a id="client-side-defense-policy-js-insert-all-pages-except-exclude-list-domain-suffix-value"></a>&#x2022; [`suffix_value`](#client-side-defense-policy-js-insert-all-pages-except-exclude-list-domain-suffix-value) - Optional String<br>Suffix Value. Suffix of domain name e.g 'xyz.com' will match '*.xyz.com' and 'xyz.com'

#### Client Side Defense Policy Js Insert All Pages Except Exclude List Metadata

A [`metadata`](#client-side-defense-policy-js-insert-all-pages-except-exclude-list-metadata) block (within [`client_side_defense.policy.js_insert_all_pages_except.exclude_list`](#client-side-defense-policy-js-insert-all-pages-except-exclude-list)) supports the following:

<a id="client-side-defense-policy-js-insert-all-pages-except-exclude-list-metadata-description-spec"></a>&#x2022; [`description_spec`](#client-side-defense-policy-js-insert-all-pages-except-exclude-list-metadata-description-spec) - Optional String<br>Description. Human readable description

<a id="client-side-defense-policy-js-insert-all-pages-except-exclude-list-metadata-name"></a>&#x2022; [`name`](#client-side-defense-policy-js-insert-all-pages-except-exclude-list-metadata-name) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### Client Side Defense Policy Js Insert All Pages Except Exclude List Path

A [`path`](#client-side-defense-policy-js-insert-all-pages-except-exclude-list-path) block (within [`client_side_defense.policy.js_insert_all_pages_except.exclude_list`](#client-side-defense-policy-js-insert-all-pages-except-exclude-list)) supports the following:

<a id="client-side-defense-policy-js-insert-all-pages-except-exclude-list-path-path"></a>&#x2022; [`path`](#client-side-defense-policy-js-insert-all-pages-except-exclude-list-path-path) - Optional String<br>Exact. Exact path value to match

<a id="client-side-defense-policy-js-insert-all-pages-except-exclude-list-path-prefix"></a>&#x2022; [`prefix`](#client-side-defense-policy-js-insert-all-pages-except-exclude-list-path-prefix) - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

<a id="client-side-defense-policy-js-insert-all-pages-except-exclude-list-path-regex"></a>&#x2022; [`regex`](#client-side-defense-policy-js-insert-all-pages-except-exclude-list-path-regex) - Optional String<br>Regex. Regular expression of path match (e.g. the value .* will match on all paths)

#### Client Side Defense Policy Js Insertion Rules

A [`js_insertion_rules`](#client-side-defense-policy-js-insertion-rules) block (within [`client_side_defense.policy`](#client-side-defense-policy)) supports the following:

<a id="client-side-defense-policy-js-insertion-rules-exclude-list"></a>&#x2022; [`exclude_list`](#client-side-defense-policy-js-insertion-rules-exclude-list) - Optional Block<br>Exclude Paths. Optional JavaScript insertions exclude list of domain and path matchers<br>See [Exclude List](#client-side-defense-policy-js-insertion-rules-exclude-list) below.

<a id="client-side-defense-policy-js-insertion-rules-rules"></a>&#x2022; [`rules`](#client-side-defense-policy-js-insertion-rules-rules) - Optional Block<br>JavaScript Insertions. Required list of pages to insert Client-Side Defense client JavaScript<br>See [Rules](#client-side-defense-policy-js-insertion-rules-rules) below.

#### Client Side Defense Policy Js Insertion Rules Exclude List

An [`exclude_list`](#client-side-defense-policy-js-insertion-rules-exclude-list) block (within [`client_side_defense.policy.js_insertion_rules`](#client-side-defense-policy-js-insertion-rules)) supports the following:

<a id="client-side-defense-policy-js-insertion-rules-exclude-list-any-domain"></a>&#x2022; [`any_domain`](#client-side-defense-policy-js-insertion-rules-exclude-list-any-domain) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="client-side-defense-policy-js-insertion-rules-exclude-list-domain"></a>&#x2022; [`domain`](#client-side-defense-policy-js-insertion-rules-exclude-list-domain) - Optional Block<br>Domains. Domains names<br>See [Domain](#client-side-defense-policy-js-insertion-rules-exclude-list-domain) below.

<a id="client-side-defense-policy-js-insertion-rules-exclude-list-metadata"></a>&#x2022; [`metadata`](#client-side-defense-policy-js-insertion-rules-exclude-list-metadata) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#client-side-defense-policy-js-insertion-rules-exclude-list-metadata) below.

<a id="client-side-defense-policy-js-insertion-rules-exclude-list-path"></a>&#x2022; [`path`](#client-side-defense-policy-js-insertion-rules-exclude-list-path) - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#client-side-defense-policy-js-insertion-rules-exclude-list-path) below.

#### Client Side Defense Policy Js Insertion Rules Exclude List Domain

A [`domain`](#client-side-defense-policy-js-insertion-rules-exclude-list-domain) block (within [`client_side_defense.policy.js_insertion_rules.exclude_list`](#client-side-defense-policy-js-insertion-rules-exclude-list)) supports the following:

<a id="client-side-defense-policy-js-insertion-rules-exclude-list-domain-exact-value"></a>&#x2022; [`exact_value`](#client-side-defense-policy-js-insertion-rules-exclude-list-domain-exact-value) - Optional String<br>Exact Value. Exact domain name

<a id="client-side-defense-policy-js-insertion-rules-exclude-list-domain-regex-value"></a>&#x2022; [`regex_value`](#client-side-defense-policy-js-insertion-rules-exclude-list-domain-regex-value) - Optional String<br>Regex Values of Domains. Regular Expression value for the domain name

<a id="client-side-defense-policy-js-insertion-rules-exclude-list-domain-suffix-value"></a>&#x2022; [`suffix_value`](#client-side-defense-policy-js-insertion-rules-exclude-list-domain-suffix-value) - Optional String<br>Suffix Value. Suffix of domain name e.g 'xyz.com' will match '*.xyz.com' and 'xyz.com'

#### Client Side Defense Policy Js Insertion Rules Exclude List Metadata

A [`metadata`](#client-side-defense-policy-js-insertion-rules-exclude-list-metadata) block (within [`client_side_defense.policy.js_insertion_rules.exclude_list`](#client-side-defense-policy-js-insertion-rules-exclude-list)) supports the following:

<a id="client-side-defense-policy-js-insertion-rules-exclude-list-metadata-description-spec"></a>&#x2022; [`description_spec`](#client-side-defense-policy-js-insertion-rules-exclude-list-metadata-description-spec) - Optional String<br>Description. Human readable description

<a id="client-side-defense-policy-js-insertion-rules-exclude-list-metadata-name"></a>&#x2022; [`name`](#client-side-defense-policy-js-insertion-rules-exclude-list-metadata-name) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### Client Side Defense Policy Js Insertion Rules Exclude List Path

A [`path`](#client-side-defense-policy-js-insertion-rules-exclude-list-path) block (within [`client_side_defense.policy.js_insertion_rules.exclude_list`](#client-side-defense-policy-js-insertion-rules-exclude-list)) supports the following:

<a id="client-side-defense-policy-js-insertion-rules-exclude-list-path-path"></a>&#x2022; [`path`](#client-side-defense-policy-js-insertion-rules-exclude-list-path-path) - Optional String<br>Exact. Exact path value to match

<a id="client-side-defense-policy-js-insertion-rules-exclude-list-path-prefix"></a>&#x2022; [`prefix`](#client-side-defense-policy-js-insertion-rules-exclude-list-path-prefix) - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

<a id="client-side-defense-policy-js-insertion-rules-exclude-list-path-regex"></a>&#x2022; [`regex`](#client-side-defense-policy-js-insertion-rules-exclude-list-path-regex) - Optional String<br>Regex. Regular expression of path match (e.g. the value .* will match on all paths)

#### Client Side Defense Policy Js Insertion Rules Rules

A [`rules`](#client-side-defense-policy-js-insertion-rules-rules) block (within [`client_side_defense.policy.js_insertion_rules`](#client-side-defense-policy-js-insertion-rules)) supports the following:

<a id="client-side-defense-policy-js-insertion-rules-rules-any-domain"></a>&#x2022; [`any_domain`](#client-side-defense-policy-js-insertion-rules-rules-any-domain) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="client-side-defense-policy-js-insertion-rules-rules-domain"></a>&#x2022; [`domain`](#client-side-defense-policy-js-insertion-rules-rules-domain) - Optional Block<br>Domains. Domains names<br>See [Domain](#client-side-defense-policy-js-insertion-rules-rules-domain) below.

<a id="client-side-defense-policy-js-insertion-rules-rules-metadata"></a>&#x2022; [`metadata`](#client-side-defense-policy-js-insertion-rules-rules-metadata) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#client-side-defense-policy-js-insertion-rules-rules-metadata) below.

<a id="client-side-defense-policy-js-insertion-rules-rules-path"></a>&#x2022; [`path`](#client-side-defense-policy-js-insertion-rules-rules-path) - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#client-side-defense-policy-js-insertion-rules-rules-path) below.

#### Client Side Defense Policy Js Insertion Rules Rules Domain

A [`domain`](#client-side-defense-policy-js-insertion-rules-rules-domain) block (within [`client_side_defense.policy.js_insertion_rules.rules`](#client-side-defense-policy-js-insertion-rules-rules)) supports the following:

<a id="client-side-defense-policy-js-insertion-rules-rules-domain-exact-value"></a>&#x2022; [`exact_value`](#client-side-defense-policy-js-insertion-rules-rules-domain-exact-value) - Optional String<br>Exact Value. Exact domain name

<a id="client-side-defense-policy-js-insertion-rules-rules-domain-regex-value"></a>&#x2022; [`regex_value`](#client-side-defense-policy-js-insertion-rules-rules-domain-regex-value) - Optional String<br>Regex Values of Domains. Regular Expression value for the domain name

<a id="client-side-defense-policy-js-insertion-rules-rules-domain-suffix-value"></a>&#x2022; [`suffix_value`](#client-side-defense-policy-js-insertion-rules-rules-domain-suffix-value) - Optional String<br>Suffix Value. Suffix of domain name e.g 'xyz.com' will match '*.xyz.com' and 'xyz.com'

#### Client Side Defense Policy Js Insertion Rules Rules Metadata

A [`metadata`](#client-side-defense-policy-js-insertion-rules-rules-metadata) block (within [`client_side_defense.policy.js_insertion_rules.rules`](#client-side-defense-policy-js-insertion-rules-rules)) supports the following:

<a id="client-side-defense-policy-js-insertion-rules-rules-metadata-description-spec"></a>&#x2022; [`description_spec`](#client-side-defense-policy-js-insertion-rules-rules-metadata-description-spec) - Optional String<br>Description. Human readable description

<a id="client-side-defense-policy-js-insertion-rules-rules-metadata-name"></a>&#x2022; [`name`](#client-side-defense-policy-js-insertion-rules-rules-metadata-name) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### Client Side Defense Policy Js Insertion Rules Rules Path

A [`path`](#client-side-defense-policy-js-insertion-rules-rules-path) block (within [`client_side_defense.policy.js_insertion_rules.rules`](#client-side-defense-policy-js-insertion-rules-rules)) supports the following:

<a id="client-side-defense-policy-js-insertion-rules-rules-path-path"></a>&#x2022; [`path`](#client-side-defense-policy-js-insertion-rules-rules-path-path) - Optional String<br>Exact. Exact path value to match

<a id="client-side-defense-policy-js-insertion-rules-rules-path-prefix"></a>&#x2022; [`prefix`](#client-side-defense-policy-js-insertion-rules-rules-path-prefix) - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

<a id="client-side-defense-policy-js-insertion-rules-rules-path-regex"></a>&#x2022; [`regex`](#client-side-defense-policy-js-insertion-rules-rules-path-regex) - Optional String<br>Regex. Regular expression of path match (e.g. the value .* will match on all paths)

#### Cookie Stickiness

A [`cookie_stickiness`](#cookie-stickiness) block supports the following:

<a id="cookie-stickiness-add-httponly"></a>&#x2022; [`add_httponly`](#cookie-stickiness-add-httponly) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="cookie-stickiness-add-secure"></a>&#x2022; [`add_secure`](#cookie-stickiness-add-secure) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="cookie-stickiness-ignore-httponly"></a>&#x2022; [`ignore_httponly`](#cookie-stickiness-ignore-httponly) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="cookie-stickiness-ignore-samesite"></a>&#x2022; [`ignore_samesite`](#cookie-stickiness-ignore-samesite) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="cookie-stickiness-ignore-secure"></a>&#x2022; [`ignore_secure`](#cookie-stickiness-ignore-secure) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="cookie-stickiness-name"></a>&#x2022; [`name`](#cookie-stickiness-name) - Optional String<br>Name. The name of the cookie that will be used to obtain the hash key. If the cookie is not present and TTL below is not set, no hash will be produced

<a id="cookie-stickiness-path"></a>&#x2022; [`path`](#cookie-stickiness-path) - Optional String<br>Path. The name of the path for the cookie. If no path is specified here, no path will be set for the cookie

<a id="cookie-stickiness-samesite-lax"></a>&#x2022; [`samesite_lax`](#cookie-stickiness-samesite-lax) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="cookie-stickiness-samesite-none"></a>&#x2022; [`samesite_none`](#cookie-stickiness-samesite-none) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="cookie-stickiness-samesite-strict"></a>&#x2022; [`samesite_strict`](#cookie-stickiness-samesite-strict) - Optional Block<br>Empty. This can be used for messages where no values are needed

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

<a id="csrf-policy-all-load-balancer-domains"></a>&#x2022; [`all_load_balancer_domains`](#csrf-policy-all-load-balancer-domains) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="csrf-policy-custom-domain-list"></a>&#x2022; [`custom_domain_list`](#csrf-policy-custom-domain-list) - Optional Block<br>Domain name list. List of domain names used for Host header matching<br>See [Custom Domain List](#csrf-policy-custom-domain-list) below.

<a id="csrf-policy-disabled"></a>&#x2022; [`disabled`](#csrf-policy-disabled) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### CSRF Policy Custom Domain List

A [`custom_domain_list`](#csrf-policy-custom-domain-list) block (within [`csrf_policy`](#csrf-policy)) supports the following:

<a id="csrf-policy-custom-domain-list-domains"></a>&#x2022; [`domains`](#csrf-policy-custom-domain-list-domains) - Optional List<br>Domain names. A list of domain names that will be matched to loadbalancer. These domains are not used for SNI match. Wildcard names are supported in the suffix or prefix form

#### Data Guard Rules

A [`data_guard_rules`](#data-guard-rules) block supports the following:

<a id="data-guard-rules-any-domain"></a>&#x2022; [`any_domain`](#data-guard-rules-any-domain) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="data-guard-rules-apply-data-guard"></a>&#x2022; [`apply_data_guard`](#data-guard-rules-apply-data-guard) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="data-guard-rules-exact-value"></a>&#x2022; [`exact_value`](#data-guard-rules-exact-value) - Optional String<br>Exact Value. Exact domain name

<a id="data-guard-rules-metadata"></a>&#x2022; [`metadata`](#data-guard-rules-metadata) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#data-guard-rules-metadata) below.

<a id="data-guard-rules-path"></a>&#x2022; [`path`](#data-guard-rules-path) - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#data-guard-rules-path) below.

<a id="data-guard-rules-skip-data-guard"></a>&#x2022; [`skip_data_guard`](#data-guard-rules-skip-data-guard) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="data-guard-rules-suffix-value"></a>&#x2022; [`suffix_value`](#data-guard-rules-suffix-value) - Optional String<br>Suffix Value. Suffix of domain name e.g 'xyz.com' will match '*.xyz.com' and 'xyz.com'

#### Data Guard Rules Metadata

A [`metadata`](#data-guard-rules-metadata) block (within [`data_guard_rules`](#data-guard-rules)) supports the following:

<a id="data-guard-rules-metadata-description-spec"></a>&#x2022; [`description_spec`](#data-guard-rules-metadata-description-spec) - Optional String<br>Description. Human readable description

<a id="data-guard-rules-metadata-name"></a>&#x2022; [`name`](#data-guard-rules-metadata-name) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### Data Guard Rules Path

A [`path`](#data-guard-rules-path) block (within [`data_guard_rules`](#data-guard-rules)) supports the following:

<a id="data-guard-rules-path-path"></a>&#x2022; [`path`](#data-guard-rules-path-path) - Optional String<br>Exact. Exact path value to match

<a id="data-guard-rules-path-prefix"></a>&#x2022; [`prefix`](#data-guard-rules-path-prefix) - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

<a id="data-guard-rules-path-regex"></a>&#x2022; [`regex`](#data-guard-rules-path-regex) - Optional String<br>Regex. Regular expression of path match (e.g. the value .* will match on all paths)

#### DDOS Mitigation Rules

A [`ddos_mitigation_rules`](#ddos-mitigation-rules) block supports the following:

<a id="ddos-mitigation-rules-block"></a>&#x2022; [`block`](#ddos-mitigation-rules-block) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="ddos-mitigation-rules-ddos-client-source"></a>&#x2022; [`ddos_client_source`](#ddos-mitigation-rules-ddos-client-source) - Optional Block<br>DDOS Client Source Choice. DDOS Mitigation sources to be blocked<br>See [DDOS Client Source](#ddos-mitigation-rules-ddos-client-source) below.

<a id="ddos-mitigation-rules-expiration-timestamp"></a>&#x2022; [`expiration_timestamp`](#ddos-mitigation-rules-expiration-timestamp) - Optional String<br>Expiration Timestamp. The expiration_timestamp is the RFC 3339 format timestamp at which the containing rule is considered to be logically expired. The rule continues to exist in the configuration but is not applied anymore

<a id="ddos-mitigation-rules-ip-prefix-list"></a>&#x2022; [`ip_prefix_list`](#ddos-mitigation-rules-ip-prefix-list) - Optional Block<br>IP Prefix Match List. List of IP Prefix strings to match against<br>See [IP Prefix List](#ddos-mitigation-rules-ip-prefix-list) below.

<a id="ddos-mitigation-rules-metadata"></a>&#x2022; [`metadata`](#ddos-mitigation-rules-metadata) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#ddos-mitigation-rules-metadata) below.

#### DDOS Mitigation Rules DDOS Client Source

A [`ddos_client_source`](#ddos-mitigation-rules-ddos-client-source) block (within [`ddos_mitigation_rules`](#ddos-mitigation-rules)) supports the following:

<a id="ddos-mitigation-rules-ddos-client-source-asn-list"></a>&#x2022; [`asn_list`](#ddos-mitigation-rules-ddos-client-source-asn-list) - Optional Block<br>ASN Match List. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer<br>See [Asn List](#ddos-mitigation-rules-ddos-client-source-asn-list) below.

<a id="ddos-mitigation-rules-ddos-client-source-country-list"></a>&#x2022; [`country_list`](#ddos-mitigation-rules-ddos-client-source-country-list) - Optional List  Defaults to `COUNTRY_NONE`<br>Possible values are `COUNTRY_NONE`, `COUNTRY_AD`, `COUNTRY_AE`, `COUNTRY_AF`, `COUNTRY_AG`, `COUNTRY_AI`, `COUNTRY_AL`, `COUNTRY_AM`, `COUNTRY_AN`, `COUNTRY_AO`, `COUNTRY_AQ`, `COUNTRY_AR`, `COUNTRY_AS`, `COUNTRY_AT`, `COUNTRY_AU`, `COUNTRY_AW`, `COUNTRY_AX`, `COUNTRY_AZ`, `COUNTRY_BA`, `COUNTRY_BB`, `COUNTRY_BD`, `COUNTRY_BE`, `COUNTRY_BF`, `COUNTRY_BG`, `COUNTRY_BH`, `COUNTRY_BI`, `COUNTRY_BJ`, `COUNTRY_BL`, `COUNTRY_BM`, `COUNTRY_BN`, `COUNTRY_BO`, `COUNTRY_BQ`, `COUNTRY_BR`, `COUNTRY_BS`, `COUNTRY_BT`, `COUNTRY_BV`, `COUNTRY_BW`, `COUNTRY_BY`, `COUNTRY_BZ`, `COUNTRY_CA`, `COUNTRY_CC`, `COUNTRY_CD`, `COUNTRY_CF`, `COUNTRY_CG`, `COUNTRY_CH`, `COUNTRY_CI`, `COUNTRY_CK`, `COUNTRY_CL`, `COUNTRY_CM`, `COUNTRY_CN`, `COUNTRY_CO`, `COUNTRY_CR`, `COUNTRY_CS`, `COUNTRY_CU`, `COUNTRY_CV`, `COUNTRY_CW`, `COUNTRY_CX`, `COUNTRY_CY`, `COUNTRY_CZ`, `COUNTRY_DE`, `COUNTRY_DJ`, `COUNTRY_DK`, `COUNTRY_DM`, `COUNTRY_DO`, `COUNTRY_DZ`, `COUNTRY_EC`, `COUNTRY_EE`, `COUNTRY_EG`, `COUNTRY_EH`, `COUNTRY_ER`, `COUNTRY_ES`, `COUNTRY_ET`, `COUNTRY_FI`, `COUNTRY_FJ`, `COUNTRY_FK`, `COUNTRY_FM`, `COUNTRY_FO`, `COUNTRY_FR`, `COUNTRY_GA`, `COUNTRY_GB`, `COUNTRY_GD`, `COUNTRY_GE`, `COUNTRY_GF`, `COUNTRY_GG`, `COUNTRY_GH`, `COUNTRY_GI`, `COUNTRY_GL`, `COUNTRY_GM`, `COUNTRY_GN`, `COUNTRY_GP`, `COUNTRY_GQ`, `COUNTRY_GR`, `COUNTRY_GS`, `COUNTRY_GT`, `COUNTRY_GU`, `COUNTRY_GW`, `COUNTRY_GY`, `COUNTRY_HK`, `COUNTRY_HM`, `COUNTRY_HN`, `COUNTRY_HR`, `COUNTRY_HT`, `COUNTRY_HU`, `COUNTRY_ID`, `COUNTRY_IE`, `COUNTRY_IL`, `COUNTRY_IM`, `COUNTRY_IN`, `COUNTRY_IO`, `COUNTRY_IQ`, `COUNTRY_IR`, `COUNTRY_IS`, `COUNTRY_IT`, `COUNTRY_JE`, `COUNTRY_JM`, `COUNTRY_JO`, `COUNTRY_JP`, `COUNTRY_KE`, `COUNTRY_KG`, `COUNTRY_KH`, `COUNTRY_KI`, `COUNTRY_KM`, `COUNTRY_KN`, `COUNTRY_KP`, `COUNTRY_KR`, `COUNTRY_KW`, `COUNTRY_KY`, `COUNTRY_KZ`, `COUNTRY_LA`, `COUNTRY_LB`, `COUNTRY_LC`, `COUNTRY_LI`, `COUNTRY_LK`, `COUNTRY_LR`, `COUNTRY_LS`, `COUNTRY_LT`, `COUNTRY_LU`, `COUNTRY_LV`, `COUNTRY_LY`, `COUNTRY_MA`, `COUNTRY_MC`, `COUNTRY_MD`, `COUNTRY_ME`, `COUNTRY_MF`, `COUNTRY_MG`, `COUNTRY_MH`, `COUNTRY_MK`, `COUNTRY_ML`, `COUNTRY_MM`, `COUNTRY_MN`, `COUNTRY_MO`, `COUNTRY_MP`, `COUNTRY_MQ`, `COUNTRY_MR`, `COUNTRY_MS`, `COUNTRY_MT`, `COUNTRY_MU`, `COUNTRY_MV`, `COUNTRY_MW`, `COUNTRY_MX`, `COUNTRY_MY`, `COUNTRY_MZ`, `COUNTRY_NA`, `COUNTRY_NC`, `COUNTRY_NE`, `COUNTRY_NF`, `COUNTRY_NG`, `COUNTRY_NI`, `COUNTRY_NL`, `COUNTRY_NO`, `COUNTRY_NP`, `COUNTRY_NR`, `COUNTRY_NU`, `COUNTRY_NZ`, `COUNTRY_OM`, `COUNTRY_PA`, `COUNTRY_PE`, `COUNTRY_PF`, `COUNTRY_PG`, `COUNTRY_PH`, `COUNTRY_PK`, `COUNTRY_PL`, `COUNTRY_PM`, `COUNTRY_PN`, `COUNTRY_PR`, `COUNTRY_PS`, `COUNTRY_PT`, `COUNTRY_PW`, `COUNTRY_PY`, `COUNTRY_QA`, `COUNTRY_RE`, `COUNTRY_RO`, `COUNTRY_RS`, `COUNTRY_RU`, `COUNTRY_RW`, `COUNTRY_SA`, `COUNTRY_SB`, `COUNTRY_SC`, `COUNTRY_SD`, `COUNTRY_SE`, `COUNTRY_SG`, `COUNTRY_SH`, `COUNTRY_SI`, `COUNTRY_SJ`, `COUNTRY_SK`, `COUNTRY_SL`, `COUNTRY_SM`, `COUNTRY_SN`, `COUNTRY_SO`, `COUNTRY_SR`, `COUNTRY_SS`, `COUNTRY_ST`, `COUNTRY_SV`, `COUNTRY_SX`, `COUNTRY_SY`, `COUNTRY_SZ`, `COUNTRY_TC`, `COUNTRY_TD`, `COUNTRY_TF`, `COUNTRY_TG`, `COUNTRY_TH`, `COUNTRY_TJ`, `COUNTRY_TK`, `COUNTRY_TL`, `COUNTRY_TM`, `COUNTRY_TN`, `COUNTRY_TO`, `COUNTRY_TR`, `COUNTRY_TT`, `COUNTRY_TV`, `COUNTRY_TW`, `COUNTRY_TZ`, `COUNTRY_UA`, `COUNTRY_UG`, `COUNTRY_UM`, `COUNTRY_US`, `COUNTRY_UY`, `COUNTRY_UZ`, `COUNTRY_VA`, `COUNTRY_VC`, `COUNTRY_VE`, `COUNTRY_VG`, `COUNTRY_VI`, `COUNTRY_VN`, `COUNTRY_VU`, `COUNTRY_WF`, `COUNTRY_WS`, `COUNTRY_XK`, `COUNTRY_XT`, `COUNTRY_YE`, `COUNTRY_YT`, `COUNTRY_ZA`, `COUNTRY_ZM`, `COUNTRY_ZW`<br>Country List. Sources that are located in one of the countries in the given list

<a id="ddos-mitigation-rules-ddos-client-source-ja4-tls-fingerprint-matcher"></a>&#x2022; [`ja4_tls_fingerprint_matcher`](#ddos-mitigation-rules-ddos-client-source-ja4-tls-fingerprint-matcher) - Optional Block<br>JA4 TLS Fingerprint Matcher. An extended version of JA3 that includes additional fields for more comprehensive fingerprinting of SSL/TLS clients and potentially has a different structure and length<br>See [Ja4 TLS Fingerprint Matcher](#ddos-mitigation-rules-ddos-client-source-ja4-tls-fingerprint-matcher) below.

<a id="ddos-mitigation-rules-ddos-client-source-tls-fingerprint-matcher"></a>&#x2022; [`tls_fingerprint_matcher`](#ddos-mitigation-rules-ddos-client-source-tls-fingerprint-matcher) - Optional Block<br>TLS Fingerprint Matcher. A TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of supported positve match criteria includes a list of known classes of TLS fingerprints and a list of exact values. The match is considered successful if either of these positive criteria are satisfied and the input fingerprint is not one of the excluded values<br>See [TLS Fingerprint Matcher](#ddos-mitigation-rules-ddos-client-source-tls-fingerprint-matcher) below.

#### DDOS Mitigation Rules DDOS Client Source Asn List

An [`asn_list`](#ddos-mitigation-rules-ddos-client-source-asn-list) block (within [`ddos_mitigation_rules.ddos_client_source`](#ddos-mitigation-rules-ddos-client-source)) supports the following:

<a id="ddos-mitigation-rules-ddos-client-source-asn-list-as-numbers"></a>&#x2022; [`as_numbers`](#ddos-mitigation-rules-ddos-client-source-asn-list-as-numbers) - Optional List<br>AS Numbers. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer

#### DDOS Mitigation Rules DDOS Client Source Ja4 TLS Fingerprint Matcher

A [`ja4_tls_fingerprint_matcher`](#ddos-mitigation-rules-ddos-client-source-ja4-tls-fingerprint-matcher) block (within [`ddos_mitigation_rules.ddos_client_source`](#ddos-mitigation-rules-ddos-client-source)) supports the following:

<a id="ddos-mitigation-rules-ddos-client-source-ja4-tls-fingerprint-matcher-exact-values"></a>&#x2022; [`exact_values`](#ddos-mitigation-rules-ddos-client-source-ja4-tls-fingerprint-matcher-exact-values) - Optional List<br>Exact Values. A list of exact JA4 TLS fingerprint to match the input JA4 TLS fingerprint against

#### DDOS Mitigation Rules DDOS Client Source TLS Fingerprint Matcher

A [`tls_fingerprint_matcher`](#ddos-mitigation-rules-ddos-client-source-tls-fingerprint-matcher) block (within [`ddos_mitigation_rules.ddos_client_source`](#ddos-mitigation-rules-ddos-client-source)) supports the following:

<a id="ddos-mitigation-rules-ddos-client-source-tls-fingerprint-matcher-classes"></a>&#x2022; [`classes`](#ddos-mitigation-rules-ddos-client-source-tls-fingerprint-matcher-classes) - Optional List  Defaults to `TLS_FINGERPRINT_NONE`<br>Possible values are `TLS_FINGERPRINT_NONE`, `ANY_MALICIOUS_FINGERPRINT`, `ADWARE`, `ADWIND`, `DRIDEX`, `GOOTKIT`, `GOZI`, `JBIFROST`, `QUAKBOT`, `RANSOMWARE`, `TROLDESH`, `TOFSEE`, `TORRENTLOCKER`, `TRICKBOT`<br>TLS fingerprint classes. A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against

<a id="ddos-mitigation-rules-ddos-client-source-tls-fingerprint-matcher-exact-values"></a>&#x2022; [`exact_values`](#ddos-mitigation-rules-ddos-client-source-tls-fingerprint-matcher-exact-values) - Optional List<br>Exact Values. A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against

<a id="ddos-mitigation-rules-ddos-client-source-tls-fingerprint-matcher-excluded-values"></a>&#x2022; [`excluded_values`](#ddos-mitigation-rules-ddos-client-source-tls-fingerprint-matcher-excluded-values) - Optional List<br>Excluded Values. A list of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can be used to skip known false positives when using one or more known TLS fingerprint classes in the enclosing matcher

#### DDOS Mitigation Rules IP Prefix List

An [`ip_prefix_list`](#ddos-mitigation-rules-ip-prefix-list) block (within [`ddos_mitigation_rules`](#ddos-mitigation-rules)) supports the following:

<a id="ddos-mitigation-rules-ip-prefix-list-invert-match"></a>&#x2022; [`invert_match`](#ddos-mitigation-rules-ip-prefix-list-invert-match) - Optional Bool<br>Invert Match Result. Invert the match result

<a id="ddos-mitigation-rules-ip-prefix-list-ip-prefixes"></a>&#x2022; [`ip_prefixes`](#ddos-mitigation-rules-ip-prefix-list-ip-prefixes) - Optional List<br>IPv4 Prefix List. List of IPv4 prefix strings

#### DDOS Mitigation Rules Metadata

A [`metadata`](#ddos-mitigation-rules-metadata) block (within [`ddos_mitigation_rules`](#ddos-mitigation-rules)) supports the following:

<a id="ddos-mitigation-rules-metadata-description-spec"></a>&#x2022; [`description_spec`](#ddos-mitigation-rules-metadata-description-spec) - Optional String<br>Description. Human readable description

<a id="ddos-mitigation-rules-metadata-name"></a>&#x2022; [`name`](#ddos-mitigation-rules-metadata-name) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### Default Pool

A [`default_pool`](#default-pool) block supports the following:

<a id="default-pool-advanced-options"></a>&#x2022; [`advanced_options`](#default-pool-advanced-options) - Optional Block<br>Origin Pool Advanced Options. Configure Advanced options for origin pool<br>See [Advanced Options](#default-pool-advanced-options) below.

<a id="default-pool-automatic-port"></a>&#x2022; [`automatic_port`](#default-pool-automatic-port) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="default-pool-endpoint-selection"></a>&#x2022; [`endpoint_selection`](#default-pool-endpoint-selection) - Optional String  Defaults to `DISTRIBUTED`<br>Possible values are `DISTRIBUTED`, `LOCAL_ONLY`, `LOCAL_PREFERRED`<br>Endpoint Selection Policy. Policy for selection of endpoints from local site/remote site/both Consider both remote and local endpoints for load balancing LOCAL_ONLY: Consider only local endpoints for load balancing Enable this policy to load balance ONLY among locally discovered endpoints Prefer the local endpoints for load balancing. If local endpoints are not present remote endpoints will be considered

<a id="default-pool-health-check-port"></a>&#x2022; [`health_check_port`](#default-pool-health-check-port) - Optional Number<br>Health check port. Port used for performing health check

<a id="default-pool-healthcheck"></a>&#x2022; [`healthcheck`](#default-pool-healthcheck) - Optional Block<br>Health Check object. Reference to healthcheck configuration objects<br>See [Healthcheck](#default-pool-healthcheck) below.

<a id="default-pool-lb-port"></a>&#x2022; [`lb_port`](#default-pool-lb-port) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="default-pool-loadbalancer-algorithm"></a>&#x2022; [`loadbalancer_algorithm`](#default-pool-loadbalancer-algorithm) - Optional String  Defaults to `ROUND_ROBIN`<br>Possible values are `ROUND_ROBIN`, `LEAST_REQUEST`, `RING_HASH`, `RANDOM`, `LB_OVERRIDE`<br>Load Balancer Algorithm. Different load balancing algorithms supported When a connection to a endpoint in an upstream cluster is required, the load balancer uses loadbalancer_algorithm to determine which host is selected. - ROUND_ROBIN: ROUND_ROBIN Policy in which each healthy/available upstream endpoint is selected in round robin order. - LEAST_REQUEST: LEAST_REQUEST Policy in which loadbalancer picks the upstream endpoint which has the fewest active requests - RING_HASH: RING_HASH Policy implements consistent hashing to upstream endpoints using ring hash of endpoint names Hash of the incoming request is calculated using request hash policy. The ring/modulo hash load balancer implements consistent hashing to upstream hosts. The algorithm is based on mapping all hosts onto a circle such that the addition or removal of a host from the host set changes only affect 1/N requests. This technique is also commonly known as “ketama” hashing. A consistent hashing load balancer is only effective when protocol routing is used that specifies a value to hash on. The minimum ring size governs the replication factor for each host in the ring. For example, if the minimum ring size is 1024 and there are 16 hosts, each host will be replicated 64 times. - RANDOM: RANDOM Policy in which each available upstream endpoint is selected in random order. The random load balancer selects a random healthy host. The random load balancer generally performs better than round robin if no health checking policy is configured. Random selection avoids bias towards the host in the set that comes after a failed host. - LB_OVERRIDE: Load Balancer Override Hash policy is taken from from the load balancer which is using this origin pool

<a id="default-pool-no-tls"></a>&#x2022; [`no_tls`](#default-pool-no-tls) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="default-pool-origin-servers"></a>&#x2022; [`origin_servers`](#default-pool-origin-servers) - Optional Block<br>Origin Servers. List of origin servers in this pool<br>See [Origin Servers](#default-pool-origin-servers) below.

<a id="default-pool-port"></a>&#x2022; [`port`](#default-pool-port) - Optional Number<br>Port. Endpoint service is available on this port

<a id="default-pool-same-as-endpoint-port"></a>&#x2022; [`same_as_endpoint_port`](#default-pool-same-as-endpoint-port) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="default-pool-upstream-conn-pool-reuse-type"></a>&#x2022; [`upstream_conn_pool_reuse_type`](#default-pool-upstream-conn-pool-reuse-type) - Optional Block<br>Select upstream connection pool reuse state. Select upstream connection pool reuse state for every downstream connection. This configuration choice is for HTTP(S) LB only<br>See [Upstream Conn Pool Reuse Type](#default-pool-upstream-conn-pool-reuse-type) below.

<a id="default-pool-use-tls"></a>&#x2022; [`use_tls`](#default-pool-use-tls) - Optional Block<br>TLS Parameters for Origin Servers. Upstream TLS Parameters<br>See [Use TLS](#default-pool-use-tls) below.

<a id="default-pool-view-internal"></a>&#x2022; [`view_internal`](#default-pool-view-internal) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [View Internal](#default-pool-view-internal) below.

#### Default Pool Advanced Options

An [`advanced_options`](#default-pool-advanced-options) block (within [`default_pool`](#default-pool)) supports the following:

<a id="default-pool-advanced-options-auto-http-config"></a>&#x2022; [`auto_http_config`](#default-pool-advanced-options-auto-http-config) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="default-pool-advanced-options-circuit-breaker"></a>&#x2022; [`circuit_breaker`](#default-pool-advanced-options-circuit-breaker) - Optional Block<br>Circuit Breaker. CircuitBreaker provides a mechanism for watching failures in upstream connections or requests and if the failures reach a certain threshold, automatically fail subsequent requests which allows to apply back pressure on downstream quickly<br>See [Circuit Breaker](#default-pool-advanced-options-circuit-breaker) below.

<a id="default-pool-advanced-options-connection-timeout"></a>&#x2022; [`connection_timeout`](#default-pool-advanced-options-connection-timeout) - Optional Number  Defaults to `2`  Specified in milliseconds<br>Connection Timeout. The timeout for new network connections to endpoints in the cluster.  The seconds

<a id="default-pool-advanced-options-default-circuit-breaker"></a>&#x2022; [`default_circuit_breaker`](#default-pool-advanced-options-default-circuit-breaker) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="default-pool-advanced-options-disable-circuit-breaker"></a>&#x2022; [`disable_circuit_breaker`](#default-pool-advanced-options-disable-circuit-breaker) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="default-pool-advanced-options-disable-lb-source-ip-persistance"></a>&#x2022; [`disable_lb_source_ip_persistance`](#default-pool-advanced-options-disable-lb-source-ip-persistance) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="default-pool-advanced-options-disable-outlier-detection"></a>&#x2022; [`disable_outlier_detection`](#default-pool-advanced-options-disable-outlier-detection) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="default-pool-advanced-options-disable-proxy-protocol"></a>&#x2022; [`disable_proxy_protocol`](#default-pool-advanced-options-disable-proxy-protocol) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="default-pool-advanced-options-disable-subsets"></a>&#x2022; [`disable_subsets`](#default-pool-advanced-options-disable-subsets) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="default-pool-advanced-options-enable-lb-source-ip-persistance"></a>&#x2022; [`enable_lb_source_ip_persistance`](#default-pool-advanced-options-enable-lb-source-ip-persistance) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="default-pool-advanced-options-enable-subsets"></a>&#x2022; [`enable_subsets`](#default-pool-advanced-options-enable-subsets) - Optional Block<br>Origin Pool Subset Options. Configure subset options for origin pool<br>See [Enable Subsets](#default-pool-advanced-options-enable-subsets) below.

<a id="default-pool-advanced-options-http1-config"></a>&#x2022; [`http1_config`](#default-pool-advanced-options-http1-config) - Optional Block<br>HTTP/1.1 Protocol Options. HTTP/1.1 Protocol options for upstream connections<br>See [Http1 Config](#default-pool-advanced-options-http1-config) below.

<a id="default-pool-advanced-options-http2-options"></a>&#x2022; [`http2_options`](#default-pool-advanced-options-http2-options) - Optional Block<br>Http2 Protocol Options. Http2 Protocol options for upstream connections<br>See [Http2 Options](#default-pool-advanced-options-http2-options) below.

<a id="default-pool-advanced-options-http-idle-timeout"></a>&#x2022; [`http_idle_timeout`](#default-pool-advanced-options-http-idle-timeout) - Optional Number  Defaults to `5`  Specified in milliseconds<br>HTTP Idle Timeout. The idle timeout for upstream connection pool connections. The idle timeout is defined as the period in which there are no active requests. When the idle timeout is reached the connection will be closed. Note that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive.  The minutes

<a id="default-pool-advanced-options-no-panic-threshold"></a>&#x2022; [`no_panic_threshold`](#default-pool-advanced-options-no-panic-threshold) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="default-pool-advanced-options-outlier-detection"></a>&#x2022; [`outlier_detection`](#default-pool-advanced-options-outlier-detection) - Optional Block<br>Outlier Detection. Outlier detection and ejection is the process of dynamically determining whether some number of hosts in an upstream cluster are performing unlike the others and removing them from the healthy load balancing set. Outlier detection is a form of passive health checking. Algorithm 1. A endpoint is determined to be an outlier (based on configured number of consecutive_5xx or consecutive_gateway_failures) . 2. If no endpoints have been ejected, loadbalancer will eject the host immediately. Otherwise, it checks to make sure the number of ejected hosts is below the allowed threshold (specified via max_ejection_percent setting). If the number of ejected hosts is above the threshold, the host is not ejected. 3. The endpoint is ejected for some number of milliseconds. Ejection means that the endpoint is marked unhealthy and will not be used during load balancing. The number of milliseconds is equal to the base_ejection_time value multiplied by the number of times the host has been ejected. 4. An ejected endpoint will automatically be brought back into service after the ejection time has been satisfied<br>See [Outlier Detection](#default-pool-advanced-options-outlier-detection) below.

<a id="default-pool-advanced-options-panic-threshold"></a>&#x2022; [`panic_threshold`](#default-pool-advanced-options-panic-threshold) - Optional Number<br>Panic threshold. x-example:'25' Configure a threshold (percentage of unhealthy endpoints) below which all endpoints will be considered for load balancing ignoring its health status

<a id="default-pool-advanced-options-proxy-protocol-v1"></a>&#x2022; [`proxy_protocol_v1`](#default-pool-advanced-options-proxy-protocol-v1) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="default-pool-advanced-options-proxy-protocol-v2"></a>&#x2022; [`proxy_protocol_v2`](#default-pool-advanced-options-proxy-protocol-v2) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### Default Pool Advanced Options Circuit Breaker

A [`circuit_breaker`](#default-pool-advanced-options-circuit-breaker) block (within [`default_pool.advanced_options`](#default-pool-advanced-options)) supports the following:

<a id="default-pool-advanced-options-circuit-breaker-connection-limit"></a>&#x2022; [`connection_limit`](#default-pool-advanced-options-circuit-breaker-connection-limit) - Optional Number<br>Connection Limit. The maximum number of connections that loadbalancer will establish to all hosts in an upstream cluster. In practice this is only applicable to TCP and HTTP/1.1 clusters since HTTP/2 uses a single connection to each host. Remove endpoint out of load balancing decision, if number of connections reach connection limit

<a id="default-pool-advanced-options-circuit-breaker-max-requests"></a>&#x2022; [`max_requests`](#default-pool-advanced-options-circuit-breaker-max-requests) - Optional Number<br>Maximum Request Count. The maximum number of requests that can be outstanding to all hosts in a cluster at any given time. In practice this is applicable to HTTP/2 clusters since HTTP/1.1 clusters are governed by the maximum connections (connection_limit). Remove endpoint out of load balancing decision, if requests exceed this count

<a id="default-pool-advanced-options-circuit-breaker-pending-requests"></a>&#x2022; [`pending_requests`](#default-pool-advanced-options-circuit-breaker-pending-requests) - Optional Number<br>Pending Requests. The maximum number of requests that will be queued while waiting for a ready connection pool connection. Since HTTP/2 requests are sent over a single connection, this circuit breaker only comes into play as the initial connection is created, as requests will be multiplexed immediately afterwards. For HTTP/1.1, requests are added to the list of pending requests whenever there aren’t enough upstream connections available to immediately dispatch the request, so this circuit breaker will remain in play for the lifetime of the process. Remove endpoint out of load balancing decision, if pending request reach pending_request

<a id="default-pool-advanced-options-circuit-breaker-priority"></a>&#x2022; [`priority`](#default-pool-advanced-options-circuit-breaker-priority) - Optional String  Defaults to `DEFAULT`<br>Possible values are `DEFAULT`, `HIGH`<br>Routing Priority. Priority routing for each request. Different connection pools are used based on the priority selected for the request. Also, circuit-breaker configuration at destination cluster is chosen based on selected priority. Default routing mechanism High-Priority routing mechanism

<a id="default-pool-advanced-options-circuit-breaker-retries"></a>&#x2022; [`retries`](#default-pool-advanced-options-circuit-breaker-retries) - Optional Number<br>Retry Count. The maximum number of retries that can be outstanding to all hosts in a cluster at any given time. Remove endpoint out of load balancing decision, if retries for request exceed this count

#### Default Pool Advanced Options Enable Subsets

An [`enable_subsets`](#default-pool-advanced-options-enable-subsets) block (within [`default_pool.advanced_options`](#default-pool-advanced-options)) supports the following:

<a id="default-pool-advanced-options-enable-subsets-any-endpoint"></a>&#x2022; [`any_endpoint`](#default-pool-advanced-options-enable-subsets-any-endpoint) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="default-pool-advanced-options-enable-subsets-default-subset"></a>&#x2022; [`default_subset`](#default-pool-advanced-options-enable-subsets-default-subset) - Optional Block<br>Origin Pool Default Subset. Default Subset definition<br>See [Default Subset](#default-pool-advanced-options-enable-subsets-default-subset) below.

<a id="default-pool-advanced-options-enable-subsets-endpoint-subsets"></a>&#x2022; [`endpoint_subsets`](#default-pool-advanced-options-enable-subsets-endpoint-subsets) - Optional Block<br>Origin Server Subsets Classes. List of subset class. Subsets class is defined using list of keys. Every unique combination of values of these keys form a subset withing the class<br>See [Endpoint Subsets](#default-pool-advanced-options-enable-subsets-endpoint-subsets) below.

<a id="default-pool-advanced-options-enable-subsets-fail-request"></a>&#x2022; [`fail_request`](#default-pool-advanced-options-enable-subsets-fail-request) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### Default Pool Advanced Options Enable Subsets Default Subset

A [`default_subset`](#default-pool-advanced-options-enable-subsets-default-subset) block (within [`default_pool.advanced_options.enable_subsets`](#default-pool-advanced-options-enable-subsets)) supports the following:

<a id="default-pool-advanced-options-enable-subsets-default-subset-default-subset"></a>&#x2022; [`default_subset`](#default-pool-advanced-options-enable-subsets-default-subset-default-subset) - Optional Block<br>Default Subset for Origin Pool. List of key-value pairs that define default subset. which gets used when route specifies no metadata or no subset matching the metadata exists

#### Default Pool Advanced Options Enable Subsets Endpoint Subsets

An [`endpoint_subsets`](#default-pool-advanced-options-enable-subsets-endpoint-subsets) block (within [`default_pool.advanced_options.enable_subsets`](#default-pool-advanced-options-enable-subsets)) supports the following:

<a id="default-pool-advanced-options-enable-subsets-endpoint-subsets-keys"></a>&#x2022; [`keys`](#default-pool-advanced-options-enable-subsets-endpoint-subsets-keys) - Optional List<br>Keys. List of keys that define a cluster subset class

#### Default Pool Advanced Options Http1 Config

A [`http1_config`](#default-pool-advanced-options-http1-config) block (within [`default_pool.advanced_options`](#default-pool-advanced-options)) supports the following:

<a id="default-pool-advanced-options-http1-config-header-transformation"></a>&#x2022; [`header_transformation`](#default-pool-advanced-options-http1-config-header-transformation) - Optional Block<br>Header Transformation. Header Transformation options for HTTP/1.1 request/response headers<br>See [Header Transformation](#default-pool-advanced-options-http1-config-header-transformation) below.

#### Default Pool Advanced Options Http1 Config Header Transformation

A [`header_transformation`](#default-pool-advanced-options-http1-config-header-transformation) block (within [`default_pool.advanced_options.http1_config`](#default-pool-advanced-options-http1-config)) supports the following:

<a id="default-pool-advanced-options-http1-config-header-transformation-default-header-transformation"></a>&#x2022; [`default_header_transformation`](#default-pool-advanced-options-http1-config-header-transformation-default-header-transformation) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="default-pool-advanced-options-http1-config-header-transformation-legacy-header-transformation"></a>&#x2022; [`legacy_header_transformation`](#default-pool-advanced-options-http1-config-header-transformation-legacy-header-transformation) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="default-pool-advanced-options-http1-config-header-transformation-preserve-case-header-transformation"></a>&#x2022; [`preserve_case_header_transformation`](#default-pool-advanced-options-http1-config-header-transformation-preserve-case-header-transformation) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="default-pool-advanced-options-http1-config-header-transformation-proper-case-header-transformation"></a>&#x2022; [`proper_case_header_transformation`](#default-pool-advanced-options-http1-config-header-transformation-proper-case-header-transformation) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### Default Pool Advanced Options Http2 Options

A [`http2_options`](#default-pool-advanced-options-http2-options) block (within [`default_pool.advanced_options`](#default-pool-advanced-options)) supports the following:

<a id="default-pool-advanced-options-http2-options-enabled"></a>&#x2022; [`enabled`](#default-pool-advanced-options-http2-options-enabled) - Optional Bool<br>HTTP2 Enabled. Enable/disable HTTP2 Protocol for upstream connections

#### Default Pool Advanced Options Outlier Detection

An [`outlier_detection`](#default-pool-advanced-options-outlier-detection) block (within [`default_pool.advanced_options`](#default-pool-advanced-options)) supports the following:

<a id="default-pool-advanced-options-outlier-detection-base-ejection-time"></a>&#x2022; [`base_ejection_time`](#default-pool-advanced-options-outlier-detection-base-ejection-time) - Optional Number  Defaults to `30000ms`  Specified in milliseconds<br>Base Ejection Time. The base time that a host is ejected for. The real time is equal to the base time multiplied by the number of times the host has been ejected. This causes hosts to get ejected for longer periods if they continue to fail

<a id="default-pool-advanced-options-outlier-detection-consecutive-5xx"></a>&#x2022; [`consecutive_5xx`](#default-pool-advanced-options-outlier-detection-consecutive-5xx) - Optional Number  Defaults to `5`<br>Consecutive 5xx Count. If an upstream endpoint returns some number of consecutive 5xx, it will be ejected. Note that in this case a 5xx means an actual 5xx respond code, or an event that would cause the HTTP router to return one on the upstream’s behalf(reset, connection failure, etc.) consecutive_5xx indicates the number of consecutive 5xx responses required before a consecutive 5xx ejection occurs

<a id="default-pool-advanced-options-outlier-detection-consecutive-gateway-failure"></a>&#x2022; [`consecutive_gateway_failure`](#default-pool-advanced-options-outlier-detection-consecutive-gateway-failure) - Optional Number  Defaults to `5`<br>Consecutive Gateway Failure. If an upstream endpoint returns some number of consecutive “gateway errors” (502, 503 or 504 status code), it will be ejected. Note that this includes events that would cause the HTTP router to return one of these status codes on the upstream’s behalf (reset, connection failure, etc.). consecutive_gateway_failure indicates the number of consecutive gateway failures before a consecutive gateway failure ejection occurs

<a id="default-pool-advanced-options-outlier-detection-interval"></a>&#x2022; [`interval`](#default-pool-advanced-options-outlier-detection-interval) - Optional Number  Defaults to `10000ms`  Specified in milliseconds<br>Interval. The time interval between ejection analysis sweeps. This can result in both new ejections as well as endpoints being returned to service

<a id="default-pool-advanced-options-outlier-detection-max-ejection-percent"></a>&#x2022; [`max_ejection_percent`](#default-pool-advanced-options-outlier-detection-max-ejection-percent) - Optional Number  Defaults to `10%`<br>Max Ejection Percentage. The maximum % of an upstream cluster that can be ejected due to outlier detection. but will eject at least one host regardless of the value

#### Default Pool Healthcheck

A [`healthcheck`](#default-pool-healthcheck) block (within [`default_pool`](#default-pool)) supports the following:

<a id="default-pool-healthcheck-name"></a>&#x2022; [`name`](#default-pool-healthcheck-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="default-pool-healthcheck-namespace"></a>&#x2022; [`namespace`](#default-pool-healthcheck-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="default-pool-healthcheck-tenant"></a>&#x2022; [`tenant`](#default-pool-healthcheck-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Default Pool Origin Servers

An [`origin_servers`](#default-pool-origin-servers) block (within [`default_pool`](#default-pool)) supports the following:

<a id="default-pool-origin-servers-cbip-service"></a>&#x2022; [`cbip_service`](#default-pool-origin-servers-cbip-service) - Optional Block<br>Discovered Classic BIG-IP Service Name. Specify origin server with Classic BIG-IP Service (Virtual Server)<br>See [Cbip Service](#default-pool-origin-servers-cbip-service) below.

<a id="default-pool-origin-servers-consul-service"></a>&#x2022; [`consul_service`](#default-pool-origin-servers-consul-service) - Optional Block<br>Consul Service Name on given Sites. Specify origin server with Hashi Corp Consul service name and site information<br>See [Consul Service](#default-pool-origin-servers-consul-service) below.

<a id="default-pool-origin-servers-custom-endpoint-object"></a>&#x2022; [`custom_endpoint_object`](#default-pool-origin-servers-custom-endpoint-object) - Optional Block<br>Custom Endpoint Object for Origin Server. Specify origin server with a reference to endpoint object<br>See [Custom Endpoint Object](#default-pool-origin-servers-custom-endpoint-object) below.

<a id="default-pool-origin-servers-k8s-service"></a>&#x2022; [`k8s_service`](#default-pool-origin-servers-k8s-service) - Optional Block<br>K8s Service Name on given Sites. Specify origin server with K8s service name and site information<br>See [K8s Service](#default-pool-origin-servers-k8s-service) below.

<a id="default-pool-origin-servers-labels"></a>&#x2022; [`labels`](#default-pool-origin-servers-labels) - Optional Block<br>Origin Server Labels. Add Labels for this origin server, these labels can be used to form subset

<a id="default-pool-origin-servers-private-ip"></a>&#x2022; [`private_ip`](#default-pool-origin-servers-private-ip) - Optional Block<br>IP address on given Sites. Specify origin server with private or public IP address and site information<br>See [Private IP](#default-pool-origin-servers-private-ip) below.

<a id="default-pool-origin-servers-private-name"></a>&#x2022; [`private_name`](#default-pool-origin-servers-private-name) - Optional Block<br>DNS Name on given Sites. Specify origin server with private or public DNS name and site information<br>See [Private Name](#default-pool-origin-servers-private-name) below.

<a id="default-pool-origin-servers-public-ip"></a>&#x2022; [`public_ip`](#default-pool-origin-servers-public-ip) - Optional Block<br>Public IP. Specify origin server with public IP address<br>See [Public IP](#default-pool-origin-servers-public-ip) below.

<a id="default-pool-origin-servers-public-name"></a>&#x2022; [`public_name`](#default-pool-origin-servers-public-name) - Optional Block<br>Public DNS Name. Specify origin server with public DNS name<br>See [Public Name](#default-pool-origin-servers-public-name) below.

<a id="default-pool-origin-servers-vn-private-ip"></a>&#x2022; [`vn_private_ip`](#default-pool-origin-servers-vn-private-ip) - Optional Block<br>IP address Virtual Network. Specify origin server with IP on Virtual Network<br>See [Vn Private IP](#default-pool-origin-servers-vn-private-ip) below.

<a id="default-pool-origin-servers-vn-private-name"></a>&#x2022; [`vn_private_name`](#default-pool-origin-servers-vn-private-name) - Optional Block<br>DNS Name on Virtual Network. Specify origin server with DNS name on Virtual Network<br>See [Vn Private Name](#default-pool-origin-servers-vn-private-name) below.

#### Default Pool Origin Servers Cbip Service

A [`cbip_service`](#default-pool-origin-servers-cbip-service) block (within [`default_pool.origin_servers`](#default-pool-origin-servers)) supports the following:

<a id="default-pool-origin-servers-cbip-service-service-name"></a>&#x2022; [`service_name`](#default-pool-origin-servers-cbip-service-service-name) - Optional String<br>Service Name. Name of the discovered Classic BIG-IP virtual server to be used as origin

#### Default Pool Origin Servers Consul Service

A [`consul_service`](#default-pool-origin-servers-consul-service) block (within [`default_pool.origin_servers`](#default-pool-origin-servers)) supports the following:

<a id="default-pool-origin-servers-consul-service-inside-network"></a>&#x2022; [`inside_network`](#default-pool-origin-servers-consul-service-inside-network) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="default-pool-origin-servers-consul-service-outside-network"></a>&#x2022; [`outside_network`](#default-pool-origin-servers-consul-service-outside-network) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="default-pool-origin-servers-consul-service-service-name"></a>&#x2022; [`service_name`](#default-pool-origin-servers-consul-service-service-name) - Optional String<br>Service Name. Consul service name of this origin server will be listed, including cluster-id. The format is servicename:cluster-id

<a id="default-pool-origin-servers-consul-service-site-locator"></a>&#x2022; [`site_locator`](#default-pool-origin-servers-consul-service-site-locator) - Optional Block<br>Site or Virtual Site. This message defines a reference to a site or virtual site object<br>See [Site Locator](#default-pool-origin-servers-consul-service-site-locator) below.

<a id="default-pool-origin-servers-consul-service-snat-pool"></a>&#x2022; [`snat_pool`](#default-pool-origin-servers-consul-service-snat-pool) - Optional Block<br>Snat Pool. Snat Pool configuration<br>See [Snat Pool](#default-pool-origin-servers-consul-service-snat-pool) below.

#### Default Pool Origin Servers Consul Service Site Locator

A [`site_locator`](#default-pool-origin-servers-consul-service-site-locator) block (within [`default_pool.origin_servers.consul_service`](#default-pool-origin-servers-consul-service)) supports the following:

<a id="default-pool-origin-servers-consul-service-site-locator-site"></a>&#x2022; [`site`](#default-pool-origin-servers-consul-service-site-locator-site) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Site](#default-pool-origin-servers-consul-service-site-locator-site) below.

<a id="default-pool-origin-servers-consul-service-site-locator-virtual-site"></a>&#x2022; [`virtual_site`](#default-pool-origin-servers-consul-service-site-locator-virtual-site) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Virtual Site](#default-pool-origin-servers-consul-service-site-locator-virtual-site) below.

#### Default Pool Origin Servers Consul Service Site Locator Site

A [`site`](#default-pool-origin-servers-consul-service-site-locator-site) block (within [`default_pool.origin_servers.consul_service.site_locator`](#default-pool-origin-servers-consul-service-site-locator)) supports the following:

<a id="default-pool-origin-servers-consul-service-site-locator-site-name"></a>&#x2022; [`name`](#default-pool-origin-servers-consul-service-site-locator-site-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="default-pool-origin-servers-consul-service-site-locator-site-namespace"></a>&#x2022; [`namespace`](#default-pool-origin-servers-consul-service-site-locator-site-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="default-pool-origin-servers-consul-service-site-locator-site-tenant"></a>&#x2022; [`tenant`](#default-pool-origin-servers-consul-service-site-locator-site-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Default Pool Origin Servers Consul Service Site Locator Virtual Site

A [`virtual_site`](#default-pool-origin-servers-consul-service-site-locator-virtual-site) block (within [`default_pool.origin_servers.consul_service.site_locator`](#default-pool-origin-servers-consul-service-site-locator)) supports the following:

<a id="default-pool-origin-servers-consul-service-site-locator-virtual-site-name"></a>&#x2022; [`name`](#default-pool-origin-servers-consul-service-site-locator-virtual-site-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="default-pool-origin-servers-consul-service-site-locator-virtual-site-namespace"></a>&#x2022; [`namespace`](#default-pool-origin-servers-consul-service-site-locator-virtual-site-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="default-pool-origin-servers-consul-service-site-locator-virtual-site-tenant"></a>&#x2022; [`tenant`](#default-pool-origin-servers-consul-service-site-locator-virtual-site-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Default Pool Origin Servers Consul Service Snat Pool

A [`snat_pool`](#default-pool-origin-servers-consul-service-snat-pool) block (within [`default_pool.origin_servers.consul_service`](#default-pool-origin-servers-consul-service)) supports the following:

<a id="default-pool-origin-servers-consul-service-snat-pool-no-snat-pool"></a>&#x2022; [`no_snat_pool`](#default-pool-origin-servers-consul-service-snat-pool-no-snat-pool) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="default-pool-origin-servers-consul-service-snat-pool-snat-pool"></a>&#x2022; [`snat_pool`](#default-pool-origin-servers-consul-service-snat-pool-snat-pool) - Optional Block<br>IPv4 Prefix List. x-example: '192.168.20.0/24' List of IPv4 prefixes that represent an endpoint<br>See [Snat Pool](#default-pool-origin-servers-consul-service-snat-pool-snat-pool) below.

#### Default Pool Origin Servers Consul Service Snat Pool Snat Pool

A [`snat_pool`](#default-pool-origin-servers-consul-service-snat-pool-snat-pool) block (within [`default_pool.origin_servers.consul_service.snat_pool`](#default-pool-origin-servers-consul-service-snat-pool)) supports the following:

<a id="default-pool-origin-servers-consul-service-snat-pool-snat-pool-prefixes"></a>&#x2022; [`prefixes`](#default-pool-origin-servers-consul-service-snat-pool-snat-pool-prefixes) - Optional List<br>IPv4 Prefix List. List of IPv4 prefixes that represent an endpoint

#### Default Pool Origin Servers Custom Endpoint Object

A [`custom_endpoint_object`](#default-pool-origin-servers-custom-endpoint-object) block (within [`default_pool.origin_servers`](#default-pool-origin-servers)) supports the following:

<a id="default-pool-origin-servers-custom-endpoint-object-endpoint"></a>&#x2022; [`endpoint`](#default-pool-origin-servers-custom-endpoint-object-endpoint) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Endpoint](#default-pool-origin-servers-custom-endpoint-object-endpoint) below.

#### Default Pool Origin Servers Custom Endpoint Object Endpoint

An [`endpoint`](#default-pool-origin-servers-custom-endpoint-object-endpoint) block (within [`default_pool.origin_servers.custom_endpoint_object`](#default-pool-origin-servers-custom-endpoint-object)) supports the following:

<a id="default-pool-origin-servers-custom-endpoint-object-endpoint-name"></a>&#x2022; [`name`](#default-pool-origin-servers-custom-endpoint-object-endpoint-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="default-pool-origin-servers-custom-endpoint-object-endpoint-namespace"></a>&#x2022; [`namespace`](#default-pool-origin-servers-custom-endpoint-object-endpoint-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="default-pool-origin-servers-custom-endpoint-object-endpoint-tenant"></a>&#x2022; [`tenant`](#default-pool-origin-servers-custom-endpoint-object-endpoint-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Default Pool Origin Servers K8s Service

A [`k8s_service`](#default-pool-origin-servers-k8s-service) block (within [`default_pool.origin_servers`](#default-pool-origin-servers)) supports the following:

<a id="default-pool-origin-servers-k8s-service-inside-network"></a>&#x2022; [`inside_network`](#default-pool-origin-servers-k8s-service-inside-network) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="default-pool-origin-servers-k8s-service-outside-network"></a>&#x2022; [`outside_network`](#default-pool-origin-servers-k8s-service-outside-network) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="default-pool-origin-servers-k8s-service-protocol"></a>&#x2022; [`protocol`](#default-pool-origin-servers-k8s-service-protocol) - Optional String  Defaults to `PROTOCOL_TCP`<br>Possible values are `PROTOCOL_TCP`, `PROTOCOL_UDP`<br>Protocol Type. Type of protocol - PROTOCOL_TCP: TCP - PROTOCOL_UDP: UDP

<a id="default-pool-origin-servers-k8s-service-service-name"></a>&#x2022; [`service_name`](#default-pool-origin-servers-k8s-service-service-name) - Optional String<br>Service Name. K8s service name of the origin server will be listed, including the namespace and cluster-id. For vK8s services, you need to enter a string with the format servicename.namespace:cluster-id. If the servicename is 'frontend', namespace is 'speedtest' and cluster-id is 'prod', then you will enter 'frontend.speedtest:prod'. Both namespace and cluster-id are optional

<a id="default-pool-origin-servers-k8s-service-site-locator"></a>&#x2022; [`site_locator`](#default-pool-origin-servers-k8s-service-site-locator) - Optional Block<br>Site or Virtual Site. This message defines a reference to a site or virtual site object<br>See [Site Locator](#default-pool-origin-servers-k8s-service-site-locator) below.

<a id="default-pool-origin-servers-k8s-service-snat-pool"></a>&#x2022; [`snat_pool`](#default-pool-origin-servers-k8s-service-snat-pool) - Optional Block<br>Snat Pool. Snat Pool configuration<br>See [Snat Pool](#default-pool-origin-servers-k8s-service-snat-pool) below.

<a id="default-pool-origin-servers-k8s-service-vk8s-networks"></a>&#x2022; [`vk8s_networks`](#default-pool-origin-servers-k8s-service-vk8s-networks) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### Default Pool Origin Servers K8s Service Site Locator

A [`site_locator`](#default-pool-origin-servers-k8s-service-site-locator) block (within [`default_pool.origin_servers.k8s_service`](#default-pool-origin-servers-k8s-service)) supports the following:

<a id="default-pool-origin-servers-k8s-service-site-locator-site"></a>&#x2022; [`site`](#default-pool-origin-servers-k8s-service-site-locator-site) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Site](#default-pool-origin-servers-k8s-service-site-locator-site) below.

<a id="default-pool-origin-servers-k8s-service-site-locator-virtual-site"></a>&#x2022; [`virtual_site`](#default-pool-origin-servers-k8s-service-site-locator-virtual-site) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Virtual Site](#default-pool-origin-servers-k8s-service-site-locator-virtual-site) below.

#### Default Pool Origin Servers K8s Service Site Locator Site

A [`site`](#default-pool-origin-servers-k8s-service-site-locator-site) block (within [`default_pool.origin_servers.k8s_service.site_locator`](#default-pool-origin-servers-k8s-service-site-locator)) supports the following:

<a id="default-pool-origin-servers-k8s-service-site-locator-site-name"></a>&#x2022; [`name`](#default-pool-origin-servers-k8s-service-site-locator-site-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="default-pool-origin-servers-k8s-service-site-locator-site-namespace"></a>&#x2022; [`namespace`](#default-pool-origin-servers-k8s-service-site-locator-site-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="default-pool-origin-servers-k8s-service-site-locator-site-tenant"></a>&#x2022; [`tenant`](#default-pool-origin-servers-k8s-service-site-locator-site-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Default Pool Origin Servers K8s Service Site Locator Virtual Site

A [`virtual_site`](#default-pool-origin-servers-k8s-service-site-locator-virtual-site) block (within [`default_pool.origin_servers.k8s_service.site_locator`](#default-pool-origin-servers-k8s-service-site-locator)) supports the following:

<a id="default-pool-origin-servers-k8s-service-site-locator-virtual-site-name"></a>&#x2022; [`name`](#default-pool-origin-servers-k8s-service-site-locator-virtual-site-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="default-pool-origin-servers-k8s-service-site-locator-virtual-site-namespace"></a>&#x2022; [`namespace`](#default-pool-origin-servers-k8s-service-site-locator-virtual-site-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="default-pool-origin-servers-k8s-service-site-locator-virtual-site-tenant"></a>&#x2022; [`tenant`](#default-pool-origin-servers-k8s-service-site-locator-virtual-site-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Default Pool Origin Servers K8s Service Snat Pool

A [`snat_pool`](#default-pool-origin-servers-k8s-service-snat-pool) block (within [`default_pool.origin_servers.k8s_service`](#default-pool-origin-servers-k8s-service)) supports the following:

<a id="default-pool-origin-servers-k8s-service-snat-pool-no-snat-pool"></a>&#x2022; [`no_snat_pool`](#default-pool-origin-servers-k8s-service-snat-pool-no-snat-pool) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="default-pool-origin-servers-k8s-service-snat-pool-snat-pool"></a>&#x2022; [`snat_pool`](#default-pool-origin-servers-k8s-service-snat-pool-snat-pool) - Optional Block<br>IPv4 Prefix List. x-example: '192.168.20.0/24' List of IPv4 prefixes that represent an endpoint<br>See [Snat Pool](#default-pool-origin-servers-k8s-service-snat-pool-snat-pool) below.

#### Default Pool Origin Servers K8s Service Snat Pool Snat Pool

A [`snat_pool`](#default-pool-origin-servers-k8s-service-snat-pool-snat-pool) block (within [`default_pool.origin_servers.k8s_service.snat_pool`](#default-pool-origin-servers-k8s-service-snat-pool)) supports the following:

<a id="default-pool-origin-servers-k8s-service-snat-pool-snat-pool-prefixes"></a>&#x2022; [`prefixes`](#default-pool-origin-servers-k8s-service-snat-pool-snat-pool-prefixes) - Optional List<br>IPv4 Prefix List. List of IPv4 prefixes that represent an endpoint

#### Default Pool Origin Servers Private IP

A [`private_ip`](#default-pool-origin-servers-private-ip) block (within [`default_pool.origin_servers`](#default-pool-origin-servers)) supports the following:

<a id="default-pool-origin-servers-private-ip-inside-network"></a>&#x2022; [`inside_network`](#default-pool-origin-servers-private-ip-inside-network) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="default-pool-origin-servers-private-ip-ip"></a>&#x2022; [`ip`](#default-pool-origin-servers-private-ip-ip) - Optional String<br>IP. Private IPv4 address

<a id="default-pool-origin-servers-private-ip-outside-network"></a>&#x2022; [`outside_network`](#default-pool-origin-servers-private-ip-outside-network) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="default-pool-origin-servers-private-ip-segment"></a>&#x2022; [`segment`](#default-pool-origin-servers-private-ip-segment) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Segment](#default-pool-origin-servers-private-ip-segment) below.

<a id="default-pool-origin-servers-private-ip-site-locator"></a>&#x2022; [`site_locator`](#default-pool-origin-servers-private-ip-site-locator) - Optional Block<br>Site or Virtual Site. This message defines a reference to a site or virtual site object<br>See [Site Locator](#default-pool-origin-servers-private-ip-site-locator) below.

<a id="default-pool-origin-servers-private-ip-snat-pool"></a>&#x2022; [`snat_pool`](#default-pool-origin-servers-private-ip-snat-pool) - Optional Block<br>Snat Pool. Snat Pool configuration<br>See [Snat Pool](#default-pool-origin-servers-private-ip-snat-pool) below.

#### Default Pool Origin Servers Private IP Segment

A [`segment`](#default-pool-origin-servers-private-ip-segment) block (within [`default_pool.origin_servers.private_ip`](#default-pool-origin-servers-private-ip)) supports the following:

<a id="default-pool-origin-servers-private-ip-segment-name"></a>&#x2022; [`name`](#default-pool-origin-servers-private-ip-segment-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="default-pool-origin-servers-private-ip-segment-namespace"></a>&#x2022; [`namespace`](#default-pool-origin-servers-private-ip-segment-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="default-pool-origin-servers-private-ip-segment-tenant"></a>&#x2022; [`tenant`](#default-pool-origin-servers-private-ip-segment-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Default Pool Origin Servers Private IP Site Locator

A [`site_locator`](#default-pool-origin-servers-private-ip-site-locator) block (within [`default_pool.origin_servers.private_ip`](#default-pool-origin-servers-private-ip)) supports the following:

<a id="default-pool-origin-servers-private-ip-site-locator-site"></a>&#x2022; [`site`](#default-pool-origin-servers-private-ip-site-locator-site) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Site](#default-pool-origin-servers-private-ip-site-locator-site) below.

<a id="default-pool-origin-servers-private-ip-site-locator-virtual-site"></a>&#x2022; [`virtual_site`](#default-pool-origin-servers-private-ip-site-locator-virtual-site) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Virtual Site](#default-pool-origin-servers-private-ip-site-locator-virtual-site) below.

#### Default Pool Origin Servers Private IP Site Locator Site

A [`site`](#default-pool-origin-servers-private-ip-site-locator-site) block (within [`default_pool.origin_servers.private_ip.site_locator`](#default-pool-origin-servers-private-ip-site-locator)) supports the following:

<a id="default-pool-origin-servers-private-ip-site-locator-site-name"></a>&#x2022; [`name`](#default-pool-origin-servers-private-ip-site-locator-site-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="default-pool-origin-servers-private-ip-site-locator-site-namespace"></a>&#x2022; [`namespace`](#default-pool-origin-servers-private-ip-site-locator-site-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="default-pool-origin-servers-private-ip-site-locator-site-tenant"></a>&#x2022; [`tenant`](#default-pool-origin-servers-private-ip-site-locator-site-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Default Pool Origin Servers Private IP Site Locator Virtual Site

A [`virtual_site`](#default-pool-origin-servers-private-ip-site-locator-virtual-site) block (within [`default_pool.origin_servers.private_ip.site_locator`](#default-pool-origin-servers-private-ip-site-locator)) supports the following:

<a id="default-pool-origin-servers-private-ip-site-locator-virtual-site-name"></a>&#x2022; [`name`](#default-pool-origin-servers-private-ip-site-locator-virtual-site-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="default-pool-origin-servers-private-ip-site-locator-virtual-site-namespace"></a>&#x2022; [`namespace`](#default-pool-origin-servers-private-ip-site-locator-virtual-site-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="default-pool-origin-servers-private-ip-site-locator-virtual-site-tenant"></a>&#x2022; [`tenant`](#default-pool-origin-servers-private-ip-site-locator-virtual-site-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Default Pool Origin Servers Private IP Snat Pool

A [`snat_pool`](#default-pool-origin-servers-private-ip-snat-pool) block (within [`default_pool.origin_servers.private_ip`](#default-pool-origin-servers-private-ip)) supports the following:

<a id="default-pool-origin-servers-private-ip-snat-pool-no-snat-pool"></a>&#x2022; [`no_snat_pool`](#default-pool-origin-servers-private-ip-snat-pool-no-snat-pool) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="default-pool-origin-servers-private-ip-snat-pool-snat-pool"></a>&#x2022; [`snat_pool`](#default-pool-origin-servers-private-ip-snat-pool-snat-pool) - Optional Block<br>IPv4 Prefix List. x-example: '192.168.20.0/24' List of IPv4 prefixes that represent an endpoint<br>See [Snat Pool](#default-pool-origin-servers-private-ip-snat-pool-snat-pool) below.

#### Default Pool Origin Servers Private IP Snat Pool Snat Pool

A [`snat_pool`](#default-pool-origin-servers-private-ip-snat-pool-snat-pool) block (within [`default_pool.origin_servers.private_ip.snat_pool`](#default-pool-origin-servers-private-ip-snat-pool)) supports the following:

<a id="default-pool-origin-servers-private-ip-snat-pool-snat-pool-prefixes"></a>&#x2022; [`prefixes`](#default-pool-origin-servers-private-ip-snat-pool-snat-pool-prefixes) - Optional List<br>IPv4 Prefix List. List of IPv4 prefixes that represent an endpoint

#### Default Pool Origin Servers Private Name

A [`private_name`](#default-pool-origin-servers-private-name) block (within [`default_pool.origin_servers`](#default-pool-origin-servers)) supports the following:

<a id="default-pool-origin-servers-private-name-dns-name"></a>&#x2022; [`dns_name`](#default-pool-origin-servers-private-name-dns-name) - Optional String<br>DNS Name. DNS Name

<a id="default-pool-origin-servers-private-name-inside-network"></a>&#x2022; [`inside_network`](#default-pool-origin-servers-private-name-inside-network) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="default-pool-origin-servers-private-name-outside-network"></a>&#x2022; [`outside_network`](#default-pool-origin-servers-private-name-outside-network) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="default-pool-origin-servers-private-name-refresh-interval"></a>&#x2022; [`refresh_interval`](#default-pool-origin-servers-private-name-refresh-interval) - Optional Number<br>DNS Refresh Interval. Interval for DNS refresh in seconds. Max value is 7 days as per `HTTPS://datatracker.ietf.org/doc/HTML/rfc8767`

<a id="default-pool-origin-servers-private-name-segment"></a>&#x2022; [`segment`](#default-pool-origin-servers-private-name-segment) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Segment](#default-pool-origin-servers-private-name-segment) below.

<a id="default-pool-origin-servers-private-name-site-locator"></a>&#x2022; [`site_locator`](#default-pool-origin-servers-private-name-site-locator) - Optional Block<br>Site or Virtual Site. This message defines a reference to a site or virtual site object<br>See [Site Locator](#default-pool-origin-servers-private-name-site-locator) below.

<a id="default-pool-origin-servers-private-name-snat-pool"></a>&#x2022; [`snat_pool`](#default-pool-origin-servers-private-name-snat-pool) - Optional Block<br>Snat Pool. Snat Pool configuration<br>See [Snat Pool](#default-pool-origin-servers-private-name-snat-pool) below.

#### Default Pool Origin Servers Private Name Segment

A [`segment`](#default-pool-origin-servers-private-name-segment) block (within [`default_pool.origin_servers.private_name`](#default-pool-origin-servers-private-name)) supports the following:

<a id="default-pool-origin-servers-private-name-segment-name"></a>&#x2022; [`name`](#default-pool-origin-servers-private-name-segment-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="default-pool-origin-servers-private-name-segment-namespace"></a>&#x2022; [`namespace`](#default-pool-origin-servers-private-name-segment-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="default-pool-origin-servers-private-name-segment-tenant"></a>&#x2022; [`tenant`](#default-pool-origin-servers-private-name-segment-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Default Pool Origin Servers Private Name Site Locator

A [`site_locator`](#default-pool-origin-servers-private-name-site-locator) block (within [`default_pool.origin_servers.private_name`](#default-pool-origin-servers-private-name)) supports the following:

<a id="default-pool-origin-servers-private-name-site-locator-site"></a>&#x2022; [`site`](#default-pool-origin-servers-private-name-site-locator-site) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Site](#default-pool-origin-servers-private-name-site-locator-site) below.

<a id="default-pool-origin-servers-private-name-site-locator-virtual-site"></a>&#x2022; [`virtual_site`](#default-pool-origin-servers-private-name-site-locator-virtual-site) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Virtual Site](#default-pool-origin-servers-private-name-site-locator-virtual-site) below.

#### Default Pool Origin Servers Private Name Site Locator Site

A [`site`](#default-pool-origin-servers-private-name-site-locator-site) block (within [`default_pool.origin_servers.private_name.site_locator`](#default-pool-origin-servers-private-name-site-locator)) supports the following:

<a id="default-pool-origin-servers-private-name-site-locator-site-name"></a>&#x2022; [`name`](#default-pool-origin-servers-private-name-site-locator-site-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="default-pool-origin-servers-private-name-site-locator-site-namespace"></a>&#x2022; [`namespace`](#default-pool-origin-servers-private-name-site-locator-site-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="default-pool-origin-servers-private-name-site-locator-site-tenant"></a>&#x2022; [`tenant`](#default-pool-origin-servers-private-name-site-locator-site-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Default Pool Origin Servers Private Name Site Locator Virtual Site

A [`virtual_site`](#default-pool-origin-servers-private-name-site-locator-virtual-site) block (within [`default_pool.origin_servers.private_name.site_locator`](#default-pool-origin-servers-private-name-site-locator)) supports the following:

<a id="default-pool-origin-servers-private-name-site-locator-virtual-site-name"></a>&#x2022; [`name`](#default-pool-origin-servers-private-name-site-locator-virtual-site-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="default-pool-origin-servers-private-name-site-locator-virtual-site-namespace"></a>&#x2022; [`namespace`](#default-pool-origin-servers-private-name-site-locator-virtual-site-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="default-pool-origin-servers-private-name-site-locator-virtual-site-tenant"></a>&#x2022; [`tenant`](#default-pool-origin-servers-private-name-site-locator-virtual-site-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Default Pool Origin Servers Private Name Snat Pool

A [`snat_pool`](#default-pool-origin-servers-private-name-snat-pool) block (within [`default_pool.origin_servers.private_name`](#default-pool-origin-servers-private-name)) supports the following:

<a id="default-pool-origin-servers-private-name-snat-pool-no-snat-pool"></a>&#x2022; [`no_snat_pool`](#default-pool-origin-servers-private-name-snat-pool-no-snat-pool) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="default-pool-origin-servers-private-name-snat-pool-snat-pool"></a>&#x2022; [`snat_pool`](#default-pool-origin-servers-private-name-snat-pool-snat-pool) - Optional Block<br>IPv4 Prefix List. x-example: '192.168.20.0/24' List of IPv4 prefixes that represent an endpoint<br>See [Snat Pool](#default-pool-origin-servers-private-name-snat-pool-snat-pool) below.

#### Default Pool Origin Servers Private Name Snat Pool Snat Pool

A [`snat_pool`](#default-pool-origin-servers-private-name-snat-pool-snat-pool) block (within [`default_pool.origin_servers.private_name.snat_pool`](#default-pool-origin-servers-private-name-snat-pool)) supports the following:

<a id="default-pool-origin-servers-private-name-snat-pool-snat-pool-prefixes"></a>&#x2022; [`prefixes`](#default-pool-origin-servers-private-name-snat-pool-snat-pool-prefixes) - Optional List<br>IPv4 Prefix List. List of IPv4 prefixes that represent an endpoint

#### Default Pool Origin Servers Public IP

A [`public_ip`](#default-pool-origin-servers-public-ip) block (within [`default_pool.origin_servers`](#default-pool-origin-servers)) supports the following:

<a id="default-pool-origin-servers-public-ip-ip"></a>&#x2022; [`ip`](#default-pool-origin-servers-public-ip-ip) - Optional String<br>Public IPv4. Public IPv4 address

#### Default Pool Origin Servers Public Name

A [`public_name`](#default-pool-origin-servers-public-name) block (within [`default_pool.origin_servers`](#default-pool-origin-servers)) supports the following:

<a id="default-pool-origin-servers-public-name-dns-name"></a>&#x2022; [`dns_name`](#default-pool-origin-servers-public-name-dns-name) - Optional String<br>DNS Name. DNS Name

<a id="default-pool-origin-servers-public-name-refresh-interval"></a>&#x2022; [`refresh_interval`](#default-pool-origin-servers-public-name-refresh-interval) - Optional Number<br>DNS Refresh Interval. Interval for DNS refresh in seconds. Max value is 7 days as per `HTTPS://datatracker.ietf.org/doc/HTML/rfc8767`

#### Default Pool Origin Servers Vn Private IP

A [`vn_private_ip`](#default-pool-origin-servers-vn-private-ip) block (within [`default_pool.origin_servers`](#default-pool-origin-servers)) supports the following:

<a id="default-pool-origin-servers-vn-private-ip-ip"></a>&#x2022; [`ip`](#default-pool-origin-servers-vn-private-ip-ip) - Optional String<br>IPv4. IPv4 address

<a id="default-pool-origin-servers-vn-private-ip-virtual-network"></a>&#x2022; [`virtual_network`](#default-pool-origin-servers-vn-private-ip-virtual-network) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Virtual Network](#default-pool-origin-servers-vn-private-ip-virtual-network) below.

#### Default Pool Origin Servers Vn Private IP Virtual Network

A [`virtual_network`](#default-pool-origin-servers-vn-private-ip-virtual-network) block (within [`default_pool.origin_servers.vn_private_ip`](#default-pool-origin-servers-vn-private-ip)) supports the following:

<a id="default-pool-origin-servers-vn-private-ip-virtual-network-name"></a>&#x2022; [`name`](#default-pool-origin-servers-vn-private-ip-virtual-network-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="default-pool-origin-servers-vn-private-ip-virtual-network-namespace"></a>&#x2022; [`namespace`](#default-pool-origin-servers-vn-private-ip-virtual-network-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="default-pool-origin-servers-vn-private-ip-virtual-network-tenant"></a>&#x2022; [`tenant`](#default-pool-origin-servers-vn-private-ip-virtual-network-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Default Pool Origin Servers Vn Private Name

A [`vn_private_name`](#default-pool-origin-servers-vn-private-name) block (within [`default_pool.origin_servers`](#default-pool-origin-servers)) supports the following:

<a id="default-pool-origin-servers-vn-private-name-dns-name"></a>&#x2022; [`dns_name`](#default-pool-origin-servers-vn-private-name-dns-name) - Optional String<br>DNS Name. DNS Name

<a id="default-pool-origin-servers-vn-private-name-private-network"></a>&#x2022; [`private_network`](#default-pool-origin-servers-vn-private-name-private-network) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Private Network](#default-pool-origin-servers-vn-private-name-private-network) below.

#### Default Pool Origin Servers Vn Private Name Private Network

A [`private_network`](#default-pool-origin-servers-vn-private-name-private-network) block (within [`default_pool.origin_servers.vn_private_name`](#default-pool-origin-servers-vn-private-name)) supports the following:

<a id="default-pool-origin-servers-vn-private-name-private-network-name"></a>&#x2022; [`name`](#default-pool-origin-servers-vn-private-name-private-network-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="default-pool-origin-servers-vn-private-name-private-network-namespace"></a>&#x2022; [`namespace`](#default-pool-origin-servers-vn-private-name-private-network-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="default-pool-origin-servers-vn-private-name-private-network-tenant"></a>&#x2022; [`tenant`](#default-pool-origin-servers-vn-private-name-private-network-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Default Pool Upstream Conn Pool Reuse Type

An [`upstream_conn_pool_reuse_type`](#default-pool-upstream-conn-pool-reuse-type) block (within [`default_pool`](#default-pool)) supports the following:

<a id="default-pool-upstream-conn-pool-reuse-type-disable-conn-pool-reuse"></a>&#x2022; [`disable_conn_pool_reuse`](#default-pool-upstream-conn-pool-reuse-type-disable-conn-pool-reuse) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="default-pool-upstream-conn-pool-reuse-type-enable-conn-pool-reuse"></a>&#x2022; [`enable_conn_pool_reuse`](#default-pool-upstream-conn-pool-reuse-type-enable-conn-pool-reuse) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### Default Pool Use TLS

An [`use_tls`](#default-pool-use-tls) block (within [`default_pool`](#default-pool)) supports the following:

<a id="default-pool-use-tls-default-session-key-caching"></a>&#x2022; [`default_session_key_caching`](#default-pool-use-tls-default-session-key-caching) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="default-pool-use-tls-disable-session-key-caching"></a>&#x2022; [`disable_session_key_caching`](#default-pool-use-tls-disable-session-key-caching) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="default-pool-use-tls-disable-sni"></a>&#x2022; [`disable_sni`](#default-pool-use-tls-disable-sni) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="default-pool-use-tls-max-session-keys"></a>&#x2022; [`max_session_keys`](#default-pool-use-tls-max-session-keys) - Optional Number<br>Max Session Keys Cached. x-example:'25' Number of session keys that are cached

<a id="default-pool-use-tls-no-mtls"></a>&#x2022; [`no_mtls`](#default-pool-use-tls-no-mtls) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="default-pool-use-tls-skip-server-verification"></a>&#x2022; [`skip_server_verification`](#default-pool-use-tls-skip-server-verification) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="default-pool-use-tls-sni"></a>&#x2022; [`sni`](#default-pool-use-tls-sni) - Optional String<br>SNI Value. SNI value to be used

<a id="default-pool-use-tls-tls-config"></a>&#x2022; [`tls_config`](#default-pool-use-tls-tls-config) - Optional Block<br>TLS Config. This defines various options to configure TLS configuration parameters<br>See [TLS Config](#default-pool-use-tls-tls-config) below.

<a id="default-pool-use-tls-use-host-header-as-sni"></a>&#x2022; [`use_host_header_as_sni`](#default-pool-use-tls-use-host-header-as-sni) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="default-pool-use-tls-use-mtls"></a>&#x2022; [`use_mtls`](#default-pool-use-tls-use-mtls) - Optional Block<br>mTLS Certificate. mTLS Client Certificate<br>See [Use mTLS](#default-pool-use-tls-use-mtls) below.

<a id="default-pool-use-tls-use-mtls-obj"></a>&#x2022; [`use_mtls_obj`](#default-pool-use-tls-use-mtls-obj) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Use mTLS Obj](#default-pool-use-tls-use-mtls-obj) below.

<a id="default-pool-use-tls-use-server-verification"></a>&#x2022; [`use_server_verification`](#default-pool-use-tls-use-server-verification) - Optional Block<br>TLS Validation Context for Origin Servers. Upstream TLS Validation Context<br>See [Use Server Verification](#default-pool-use-tls-use-server-verification) below.

<a id="default-pool-use-tls-volterra-trusted-ca"></a>&#x2022; [`volterra_trusted_ca`](#default-pool-use-tls-volterra-trusted-ca) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### Default Pool Use TLS TLS Config

A [`tls_config`](#default-pool-use-tls-tls-config) block (within [`default_pool.use_tls`](#default-pool-use-tls)) supports the following:

<a id="default-pool-use-tls-tls-config-custom-security"></a>&#x2022; [`custom_security`](#default-pool-use-tls-tls-config-custom-security) - Optional Block<br>Custom Ciphers. This defines TLS protocol config including min/max versions and allowed ciphers<br>See [Custom Security](#default-pool-use-tls-tls-config-custom-security) below.

<a id="default-pool-use-tls-tls-config-default-security"></a>&#x2022; [`default_security`](#default-pool-use-tls-tls-config-default-security) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="default-pool-use-tls-tls-config-low-security"></a>&#x2022; [`low_security`](#default-pool-use-tls-tls-config-low-security) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="default-pool-use-tls-tls-config-medium-security"></a>&#x2022; [`medium_security`](#default-pool-use-tls-tls-config-medium-security) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### Default Pool Use TLS TLS Config Custom Security

A [`custom_security`](#default-pool-use-tls-tls-config-custom-security) block (within [`default_pool.use_tls.tls_config`](#default-pool-use-tls-tls-config)) supports the following:

<a id="default-pool-use-tls-tls-config-custom-security-cipher-suites"></a>&#x2022; [`cipher_suites`](#default-pool-use-tls-tls-config-custom-security-cipher-suites) - Optional List<br>Cipher Suites. The TLS listener will only support the specified cipher list

<a id="default-pool-use-tls-tls-config-custom-security-max-version"></a>&#x2022; [`max_version`](#default-pool-use-tls-tls-config-custom-security-max-version) - Optional String  Defaults to `TLS_AUTO`<br>Possible values are `TLS_AUTO`, `TLSv1_0`, `TLSv1_1`, `TLSv1_2`, `TLSv1_3`<br>TLS Protocol. TlsProtocol is enumeration of supported TLS versions F5 Distributed Cloud will choose the optimal TLS version

<a id="default-pool-use-tls-tls-config-custom-security-min-version"></a>&#x2022; [`min_version`](#default-pool-use-tls-tls-config-custom-security-min-version) - Optional String  Defaults to `TLS_AUTO`<br>Possible values are `TLS_AUTO`, `TLSv1_0`, `TLSv1_1`, `TLSv1_2`, `TLSv1_3`<br>TLS Protocol. TlsProtocol is enumeration of supported TLS versions F5 Distributed Cloud will choose the optimal TLS version

#### Default Pool Use TLS Use mTLS

An [`use_mtls`](#default-pool-use-tls-use-mtls) block (within [`default_pool.use_tls`](#default-pool-use-tls)) supports the following:

<a id="default-pool-use-tls-use-mtls-tls-certificates"></a>&#x2022; [`tls_certificates`](#default-pool-use-tls-use-mtls-tls-certificates) - Optional Block<br>mTLS Client Certificate. mTLS Client Certificate<br>See [TLS Certificates](#default-pool-use-tls-use-mtls-tls-certificates) below.

#### Default Pool Use TLS Use mTLS TLS Certificates

A [`tls_certificates`](#default-pool-use-tls-use-mtls-tls-certificates) block (within [`default_pool.use_tls.use_mtls`](#default-pool-use-tls-use-mtls)) supports the following:

<a id="default-pool-use-tls-use-mtls-tls-certificates-certificate-url"></a>&#x2022; [`certificate_url`](#default-pool-use-tls-use-mtls-tls-certificates-certificate-url) - Optional String<br>Certificate. TLS certificate. Certificate or certificate chain in PEM format including the PEM headers

<a id="default-pool-use-tls-use-mtls-tls-certificates-custom-hash-algorithms"></a>&#x2022; [`custom_hash_algorithms`](#default-pool-use-tls-use-mtls-tls-certificates-custom-hash-algorithms) - Optional Block<br>Hash Algorithms. Specifies the hash algorithms to be used<br>See [Custom Hash Algorithms](#default-pool-use-tls-use-mtls-tls-certificates-custom-hash-algorithms) below.

<a id="default-pool-use-tls-use-mtls-tls-certificates-description-spec"></a>&#x2022; [`description_spec`](#default-pool-use-tls-use-mtls-tls-certificates-description-spec) - Optional String<br>Description. Description for the certificate

<a id="default-pool-use-tls-use-mtls-tls-certificates-disable-ocsp-stapling"></a>&#x2022; [`disable_ocsp_stapling`](#default-pool-use-tls-use-mtls-tls-certificates-disable-ocsp-stapling) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="default-pool-use-tls-use-mtls-tls-certificates-private-key"></a>&#x2022; [`private_key`](#default-pool-use-tls-use-mtls-tls-certificates-private-key) - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Private Key](#default-pool-use-tls-use-mtls-tls-certificates-private-key) below.

<a id="default-pool-use-tls-use-mtls-tls-certificates-use-system-defaults"></a>&#x2022; [`use_system_defaults`](#default-pool-use-tls-use-mtls-tls-certificates-use-system-defaults) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### Default Pool Use TLS Use mTLS TLS Certificates Custom Hash Algorithms

A [`custom_hash_algorithms`](#default-pool-use-tls-use-mtls-tls-certificates-custom-hash-algorithms) block (within [`default_pool.use_tls.use_mtls.tls_certificates`](#default-pool-use-tls-use-mtls-tls-certificates)) supports the following:

<a id="default-pool-use-tls-use-mtls-tls-certificates-custom-hash-algorithms-hash-algorithms"></a>&#x2022; [`hash_algorithms`](#default-pool-use-tls-use-mtls-tls-certificates-custom-hash-algorithms-hash-algorithms) - Optional List  Defaults to `INVALID_HASH_ALGORITHM`<br>Possible values are `INVALID_HASH_ALGORITHM`, `SHA256`, `SHA1`<br>Hash Algorithms. Ordered list of hash algorithms to be used

#### Default Pool Use TLS Use mTLS TLS Certificates Private Key

A [`private_key`](#default-pool-use-tls-use-mtls-tls-certificates-private-key) block (within [`default_pool.use_tls.use_mtls.tls_certificates`](#default-pool-use-tls-use-mtls-tls-certificates)) supports the following:

<a id="default-pool-use-tls-use-mtls-tls-certificates-private-key-blindfold-secret-info"></a>&#x2022; [`blindfold_secret_info`](#default-pool-use-tls-use-mtls-tls-certificates-private-key-blindfold-secret-info) - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#default-pool-use-tls-use-mtls-tls-certificates-private-key-blindfold-secret-info) below.

<a id="default-pool-use-tls-use-mtls-tls-certificates-private-key-clear-secret-info"></a>&#x2022; [`clear_secret_info`](#default-pool-use-tls-use-mtls-tls-certificates-private-key-clear-secret-info) - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#default-pool-use-tls-use-mtls-tls-certificates-private-key-clear-secret-info) below.

#### Default Pool Use TLS Use mTLS TLS Certificates Private Key Blindfold Secret Info

A [`blindfold_secret_info`](#default-pool-use-tls-use-mtls-tls-certificates-private-key-blindfold-secret-info) block (within [`default_pool.use_tls.use_mtls.tls_certificates.private_key`](#default-pool-use-tls-use-mtls-tls-certificates-private-key)) supports the following:

<a id="default-pool-use-tls-use-mtls-tls-certificates-private-key-blindfold-secret-info-decryption-provider"></a>&#x2022; [`decryption_provider`](#default-pool-use-tls-use-mtls-tls-certificates-private-key-blindfold-secret-info-decryption-provider) - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

<a id="default-pool-use-tls-use-mtls-tls-certificates-private-key-blindfold-secret-info-location"></a>&#x2022; [`location`](#default-pool-use-tls-use-mtls-tls-certificates-private-key-blindfold-secret-info-location) - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

<a id="default-pool-use-tls-use-mtls-tls-certificates-private-key-blindfold-secret-info-store-provider"></a>&#x2022; [`store_provider`](#default-pool-use-tls-use-mtls-tls-certificates-private-key-blindfold-secret-info-store-provider) - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

#### Default Pool Use TLS Use mTLS TLS Certificates Private Key Clear Secret Info

A [`clear_secret_info`](#default-pool-use-tls-use-mtls-tls-certificates-private-key-clear-secret-info) block (within [`default_pool.use_tls.use_mtls.tls_certificates.private_key`](#default-pool-use-tls-use-mtls-tls-certificates-private-key)) supports the following:

<a id="default-pool-use-tls-use-mtls-tls-certificates-private-key-clear-secret-info-provider-ref"></a>&#x2022; [`provider_ref`](#default-pool-use-tls-use-mtls-tls-certificates-private-key-clear-secret-info-provider-ref) - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

<a id="default-pool-use-tls-use-mtls-tls-certificates-private-key-clear-secret-info-url"></a>&#x2022; [`url`](#default-pool-use-tls-use-mtls-tls-certificates-private-key-clear-secret-info-url) - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

#### Default Pool Use TLS Use mTLS Obj

An [`use_mtls_obj`](#default-pool-use-tls-use-mtls-obj) block (within [`default_pool.use_tls`](#default-pool-use-tls)) supports the following:

<a id="default-pool-use-tls-use-mtls-obj-name"></a>&#x2022; [`name`](#default-pool-use-tls-use-mtls-obj-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="default-pool-use-tls-use-mtls-obj-namespace"></a>&#x2022; [`namespace`](#default-pool-use-tls-use-mtls-obj-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="default-pool-use-tls-use-mtls-obj-tenant"></a>&#x2022; [`tenant`](#default-pool-use-tls-use-mtls-obj-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Default Pool Use TLS Use Server Verification

An [`use_server_verification`](#default-pool-use-tls-use-server-verification) block (within [`default_pool.use_tls`](#default-pool-use-tls)) supports the following:

<a id="default-pool-use-tls-use-server-verification-trusted-ca"></a>&#x2022; [`trusted_ca`](#default-pool-use-tls-use-server-verification-trusted-ca) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Trusted CA](#default-pool-use-tls-use-server-verification-trusted-ca) below.

<a id="default-pool-use-tls-use-server-verification-trusted-ca-url"></a>&#x2022; [`trusted_ca_url`](#default-pool-use-tls-use-server-verification-trusted-ca-url) - Optional String<br>Inline Root CA Certificate (legacy). Upload a Root CA Certificate specifically for this Origin Pool for verification of server's certificate

#### Default Pool Use TLS Use Server Verification Trusted CA

A [`trusted_ca`](#default-pool-use-tls-use-server-verification-trusted-ca) block (within [`default_pool.use_tls.use_server_verification`](#default-pool-use-tls-use-server-verification)) supports the following:

<a id="default-pool-use-tls-use-server-verification-trusted-ca-name"></a>&#x2022; [`name`](#default-pool-use-tls-use-server-verification-trusted-ca-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="default-pool-use-tls-use-server-verification-trusted-ca-namespace"></a>&#x2022; [`namespace`](#default-pool-use-tls-use-server-verification-trusted-ca-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="default-pool-use-tls-use-server-verification-trusted-ca-tenant"></a>&#x2022; [`tenant`](#default-pool-use-tls-use-server-verification-trusted-ca-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

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

<a id="default-pool-list-pools-endpoint-subsets"></a>&#x2022; [`endpoint_subsets`](#default-pool-list-pools-endpoint-subsets) - Optional Block<br>Origin Servers Subsets. Upstream origin pool may be configured to divide its origin servers into subsets based on metadata attached to the origin servers. Routes may then specify the metadata that a endpoint must match in order to be selected by the load balancer For origin servers which are discovered in K8S or Consul cluster, the label of the service is merged with endpoint's labels. In case of Consul, the label is derived from the 'Tag' field. For labels that are common between configured endpoint and discovered service, labels from discovered service takes precedence. List of key-value pairs that will be used as matching metadata. Only those origin servers of upstream origin pool which match this metadata will be selected for load balancing

<a id="default-pool-list-pools-pool"></a>&#x2022; [`pool`](#default-pool-list-pools-pool) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Pool](#default-pool-list-pools-pool) below.

<a id="default-pool-list-pools-priority"></a>&#x2022; [`priority`](#default-pool-list-pools-priority) - Optional Number<br>Priority. Priority of this origin pool, valid only with multiple origin pools. Value of 0 will make the pool as lowest priority origin pool Priority of 1 means highest priority and is considered active. When active origin pool is not available, lower priority origin pools are made active as per the increasing priority

<a id="default-pool-list-pools-weight"></a>&#x2022; [`weight`](#default-pool-list-pools-weight) - Optional Number<br>Weight. Weight of this origin pool, valid only with multiple origin pool. Value of 0 will disable the pool

#### Default Pool List Pools Cluster

A [`cluster`](#default-pool-list-pools-cluster) block (within [`default_pool_list.pools`](#default-pool-list-pools)) supports the following:

<a id="default-pool-list-pools-cluster-name"></a>&#x2022; [`name`](#default-pool-list-pools-cluster-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="default-pool-list-pools-cluster-namespace"></a>&#x2022; [`namespace`](#default-pool-list-pools-cluster-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

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

<a id="enable-api-discovery-api-discovery-from-code-scan"></a>&#x2022; [`api_discovery_from_code_scan`](#enable-api-discovery-api-discovery-from-code-scan) - Optional Block<br>Select Code Base and Repositories. x-required<br>See [API Discovery From Code Scan](#enable-api-discovery-api-discovery-from-code-scan) below.

<a id="enable-api-discovery-custom-api-auth-discovery"></a>&#x2022; [`custom_api_auth_discovery`](#enable-api-discovery-custom-api-auth-discovery) - Optional Block<br>API Discovery Advanced Settings. API Discovery Advanced settings<br>See [Custom API Auth Discovery](#enable-api-discovery-custom-api-auth-discovery) below.

<a id="enable-api-discovery-default-api-auth-discovery"></a>&#x2022; [`default_api_auth_discovery`](#enable-api-discovery-default-api-auth-discovery) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="enable-api-discovery-disable-learn-from-redirect-traffic"></a>&#x2022; [`disable_learn_from_redirect_traffic`](#enable-api-discovery-disable-learn-from-redirect-traffic) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="enable-api-discovery-discovered-api-settings"></a>&#x2022; [`discovered_api_settings`](#enable-api-discovery-discovered-api-settings) - Optional Block<br>Discovered API Settings. x-example: '2' Configure Discovered API Settings<br>See [Discovered API Settings](#enable-api-discovery-discovered-api-settings) below.

<a id="enable-api-discovery-enable-learn-from-redirect-traffic"></a>&#x2022; [`enable_learn_from_redirect_traffic`](#enable-api-discovery-enable-learn-from-redirect-traffic) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### Enable API Discovery API Crawler

An [`api_crawler`](#enable-api-discovery-api-crawler) block (within [`enable_api_discovery`](#enable-api-discovery)) supports the following:

<a id="enable-api-discovery-api-crawler-api-crawler-config"></a>&#x2022; [`api_crawler_config`](#enable-api-discovery-api-crawler-api-crawler-config) - Optional Block<br>Crawler Configure<br>See [API Crawler Config](#enable-api-discovery-api-crawler-api-crawler-config) below.

<a id="enable-api-discovery-api-crawler-disable-api-crawler"></a>&#x2022; [`disable_api_crawler`](#enable-api-discovery-api-crawler-disable-api-crawler) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### Enable API Discovery API Crawler API Crawler Config

An [`api_crawler_config`](#enable-api-discovery-api-crawler-api-crawler-config) block (within [`enable_api_discovery.api_crawler`](#enable-api-discovery-api-crawler)) supports the following:

<a id="enable-api-discovery-api-crawler-api-crawler-config-domains"></a>&#x2022; [`domains`](#enable-api-discovery-api-crawler-api-crawler-config-domains) - Optional Block<br>Domains to Crawl. Enter domains and their credentials to allow authenticated API crawling. You can only include domains you own that are associated with this Load Balancer<br>See [Domains](#enable-api-discovery-api-crawler-api-crawler-config-domains) below.

#### Enable API Discovery API Crawler API Crawler Config Domains

A [`domains`](#enable-api-discovery-api-crawler-api-crawler-config-domains) block (within [`enable_api_discovery.api_crawler.api_crawler_config`](#enable-api-discovery-api-crawler-api-crawler-config)) supports the following:

<a id="enable-api-discovery-api-crawler-api-crawler-config-domains-domain"></a>&#x2022; [`domain`](#enable-api-discovery-api-crawler-api-crawler-config-domains-domain) - Optional String<br>Domain. Select the domain to execute API Crawling with given credentials

<a id="enable-api-discovery-api-crawler-api-crawler-config-domains-simple-login"></a>&#x2022; [`simple_login`](#enable-api-discovery-api-crawler-api-crawler-config-domains-simple-login) - Optional Block<br>Simple Login<br>See [Simple Login](#enable-api-discovery-api-crawler-api-crawler-config-domains-simple-login) below.

#### Enable API Discovery API Crawler API Crawler Config Domains Simple Login

A [`simple_login`](#enable-api-discovery-api-crawler-api-crawler-config-domains-simple-login) block (within [`enable_api_discovery.api_crawler.api_crawler_config.domains`](#enable-api-discovery-api-crawler-api-crawler-config-domains)) supports the following:

<a id="enable-api-discovery-api-crawler-api-crawler-config-domains-simple-login-password"></a>&#x2022; [`password`](#enable-api-discovery-api-crawler-api-crawler-config-domains-simple-login-password) - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Password](#enable-api-discovery-api-crawler-api-crawler-config-domains-simple-login-password) below.

<a id="enable-api-discovery-api-crawler-api-crawler-config-domains-simple-login-user"></a>&#x2022; [`user`](#enable-api-discovery-api-crawler-api-crawler-config-domains-simple-login-user) - Optional String<br>User. Enter the username to assign credentials for the selected domain to crawl

#### Enable API Discovery API Crawler API Crawler Config Domains Simple Login Password

A [`password`](#enable-api-discovery-api-crawler-api-crawler-config-domains-simple-login-password) block (within [`enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login`](#enable-api-discovery-api-crawler-api-crawler-config-domains-simple-login)) supports the following:

<a id="enable-api-discovery-api-crawler-api-crawler-config-domains-simple-login-password-blindfold-secret-info"></a>&#x2022; [`blindfold_secret_info`](#enable-api-discovery-api-crawler-api-crawler-config-domains-simple-login-password-blindfold-secret-info) - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#enable-api-discovery-api-crawler-api-crawler-config-domains-simple-login-password-blindfold-secret-info) below.

<a id="enable-api-discovery-api-crawler-api-crawler-config-domains-simple-login-password-clear-secret-info"></a>&#x2022; [`clear_secret_info`](#enable-api-discovery-api-crawler-api-crawler-config-domains-simple-login-password-clear-secret-info) - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#enable-api-discovery-api-crawler-api-crawler-config-domains-simple-login-password-clear-secret-info) below.

#### Enable API Discovery API Crawler API Crawler Config Domains Simple Login Password Blindfold Secret Info

A [`blindfold_secret_info`](#enable-api-discovery-api-crawler-api-crawler-config-domains-simple-login-password-blindfold-secret-info) block (within [`enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password`](#enable-api-discovery-api-crawler-api-crawler-config-domains-simple-login-password)) supports the following:

<a id="enable-api-discovery-api-crawler-api-crawler-config-domains-simple-login-password-blindfold-secret-info-decryption-provider"></a>&#x2022; [`decryption_provider`](#enable-api-discovery-api-crawler-api-crawler-config-domains-simple-login-password-blindfold-secret-info-decryption-provider) - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

<a id="enable-api-discovery-api-crawler-api-crawler-config-domains-simple-login-password-blindfold-secret-info-location"></a>&#x2022; [`location`](#enable-api-discovery-api-crawler-api-crawler-config-domains-simple-login-password-blindfold-secret-info-location) - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

<a id="enable-api-discovery-api-crawler-api-crawler-config-domains-simple-login-password-blindfold-secret-info-store-provider"></a>&#x2022; [`store_provider`](#enable-api-discovery-api-crawler-api-crawler-config-domains-simple-login-password-blindfold-secret-info-store-provider) - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

#### Enable API Discovery API Crawler API Crawler Config Domains Simple Login Password Clear Secret Info

A [`clear_secret_info`](#enable-api-discovery-api-crawler-api-crawler-config-domains-simple-login-password-clear-secret-info) block (within [`enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password`](#enable-api-discovery-api-crawler-api-crawler-config-domains-simple-login-password)) supports the following:

<a id="enable-api-discovery-api-crawler-api-crawler-config-domains-simple-login-password-clear-secret-info-provider-ref"></a>&#x2022; [`provider_ref`](#enable-api-discovery-api-crawler-api-crawler-config-domains-simple-login-password-clear-secret-info-provider-ref) - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

<a id="enable-api-discovery-api-crawler-api-crawler-config-domains-simple-login-password-clear-secret-info-url"></a>&#x2022; [`url`](#enable-api-discovery-api-crawler-api-crawler-config-domains-simple-login-password-clear-secret-info-url) - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

#### Enable API Discovery API Discovery From Code Scan

An [`api_discovery_from_code_scan`](#enable-api-discovery-api-discovery-from-code-scan) block (within [`enable_api_discovery`](#enable-api-discovery)) supports the following:

<a id="enable-api-discovery-api-discovery-from-code-scan-code-base-integrations"></a>&#x2022; [`code_base_integrations`](#enable-api-discovery-api-discovery-from-code-scan-code-base-integrations) - Optional Block<br>Select Code Base Integrations<br>See [Code Base Integrations](#enable-api-discovery-api-discovery-from-code-scan-code-base-integrations) below.

#### Enable API Discovery API Discovery From Code Scan Code Base Integrations

A [`code_base_integrations`](#enable-api-discovery-api-discovery-from-code-scan-code-base-integrations) block (within [`enable_api_discovery.api_discovery_from_code_scan`](#enable-api-discovery-api-discovery-from-code-scan)) supports the following:

<a id="enable-api-discovery-api-discovery-from-code-scan-code-base-integrations-all-repos"></a>&#x2022; [`all_repos`](#enable-api-discovery-api-discovery-from-code-scan-code-base-integrations-all-repos) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="enable-api-discovery-api-discovery-from-code-scan-code-base-integrations-code-base-integration"></a>&#x2022; [`code_base_integration`](#enable-api-discovery-api-discovery-from-code-scan-code-base-integrations-code-base-integration) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Code Base Integration](#enable-api-discovery-api-discovery-from-code-scan-code-base-integrations-code-base-integration) below.

<a id="enable-api-discovery-api-discovery-from-code-scan-code-base-integrations-selected-repos"></a>&#x2022; [`selected_repos`](#enable-api-discovery-api-discovery-from-code-scan-code-base-integrations-selected-repos) - Optional Block<br>API Code Repositories. Select which API repositories represent the LB applications<br>See [Selected Repos](#enable-api-discovery-api-discovery-from-code-scan-code-base-integrations-selected-repos) below.

#### Enable API Discovery API Discovery From Code Scan Code Base Integrations Code Base Integration

A [`code_base_integration`](#enable-api-discovery-api-discovery-from-code-scan-code-base-integrations-code-base-integration) block (within [`enable_api_discovery.api_discovery_from_code_scan.code_base_integrations`](#enable-api-discovery-api-discovery-from-code-scan-code-base-integrations)) supports the following:

<a id="enable-api-discovery-api-discovery-from-code-scan-code-base-integrations-code-base-integration-name"></a>&#x2022; [`name`](#enable-api-discovery-api-discovery-from-code-scan-code-base-integrations-code-base-integration-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="enable-api-discovery-api-discovery-from-code-scan-code-base-integrations-code-base-integration-namespace"></a>&#x2022; [`namespace`](#enable-api-discovery-api-discovery-from-code-scan-code-base-integrations-code-base-integration-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="enable-api-discovery-api-discovery-from-code-scan-code-base-integrations-code-base-integration-tenant"></a>&#x2022; [`tenant`](#enable-api-discovery-api-discovery-from-code-scan-code-base-integrations-code-base-integration-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Enable API Discovery API Discovery From Code Scan Code Base Integrations Selected Repos

A [`selected_repos`](#enable-api-discovery-api-discovery-from-code-scan-code-base-integrations-selected-repos) block (within [`enable_api_discovery.api_discovery_from_code_scan.code_base_integrations`](#enable-api-discovery-api-discovery-from-code-scan-code-base-integrations)) supports the following:

<a id="enable-api-discovery-api-discovery-from-code-scan-code-base-integrations-selected-repos-api-code-repo"></a>&#x2022; [`api_code_repo`](#enable-api-discovery-api-discovery-from-code-scan-code-base-integrations-selected-repos-api-code-repo) - Optional List<br>API Code Repository. Code repository which contain API endpoints

#### Enable API Discovery Custom API Auth Discovery

A [`custom_api_auth_discovery`](#enable-api-discovery-custom-api-auth-discovery) block (within [`enable_api_discovery`](#enable-api-discovery)) supports the following:

<a id="enable-api-discovery-custom-api-auth-discovery-api-discovery-ref"></a>&#x2022; [`api_discovery_ref`](#enable-api-discovery-custom-api-auth-discovery-api-discovery-ref) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [API Discovery Ref](#enable-api-discovery-custom-api-auth-discovery-api-discovery-ref) below.

#### Enable API Discovery Custom API Auth Discovery API Discovery Ref

An [`api_discovery_ref`](#enable-api-discovery-custom-api-auth-discovery-api-discovery-ref) block (within [`enable_api_discovery.custom_api_auth_discovery`](#enable-api-discovery-custom-api-auth-discovery)) supports the following:

<a id="enable-api-discovery-custom-api-auth-discovery-api-discovery-ref-name"></a>&#x2022; [`name`](#enable-api-discovery-custom-api-auth-discovery-api-discovery-ref-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="enable-api-discovery-custom-api-auth-discovery-api-discovery-ref-namespace"></a>&#x2022; [`namespace`](#enable-api-discovery-custom-api-auth-discovery-api-discovery-ref-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="enable-api-discovery-custom-api-auth-discovery-api-discovery-ref-tenant"></a>&#x2022; [`tenant`](#enable-api-discovery-custom-api-auth-discovery-api-discovery-ref-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Enable API Discovery Discovered API Settings

A [`discovered_api_settings`](#enable-api-discovery-discovered-api-settings) block (within [`enable_api_discovery`](#enable-api-discovery)) supports the following:

<a id="enable-api-discovery-discovered-api-settings-purge-duration-for-inactive-discovered-apis"></a>&#x2022; [`purge_duration_for_inactive_discovered_apis`](#enable-api-discovery-discovered-api-settings-purge-duration-for-inactive-discovered-apis) - Optional Number<br>Purge Duration for Inactive Discovered APIs from Traffic. Inactive discovered API will be deleted after configured duration

#### Enable Challenge

An [`enable_challenge`](#enable-challenge) block supports the following:

<a id="enable-challenge-captcha-challenge-parameters"></a>&#x2022; [`captcha_challenge_parameters`](#enable-challenge-captcha-challenge-parameters) - Optional Block<br>Captcha Challenge Parameters. Enables loadbalancer to perform captcha challenge Captcha challenge will be based on Google Recaptcha. With this feature enabled, only clients that pass the captcha challenge will be allowed to complete the HTTP request. When loadbalancer is configured to do Captcha Challenge, it will redirect the browser to an HTML page on every new HTTP request. This HTML page will have captcha challenge embedded in it. Client will be allowed to make the request only if the captcha challenge is successful. Loadbalancer will tag response header with a cookie to avoid Captcha challenge for subsequent requests. CAPTCHA is mainly used as a security check to ensure only human users can pass through. Generally, computers or bots are not capable of solving a captcha. You can enable either Javascript challenge or Captcha challenge on a virtual host<br>See [Captcha Challenge Parameters](#enable-challenge-captcha-challenge-parameters) below.

<a id="enable-challenge-default-captcha-challenge-parameters"></a>&#x2022; [`default_captcha_challenge_parameters`](#enable-challenge-default-captcha-challenge-parameters) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="enable-challenge-default-js-challenge-parameters"></a>&#x2022; [`default_js_challenge_parameters`](#enable-challenge-default-js-challenge-parameters) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="enable-challenge-default-mitigation-settings"></a>&#x2022; [`default_mitigation_settings`](#enable-challenge-default-mitigation-settings) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="enable-challenge-js-challenge-parameters"></a>&#x2022; [`js_challenge_parameters`](#enable-challenge-js-challenge-parameters) - Optional Block<br>Javascript Challenge Parameters. Enables loadbalancer to perform client browser compatibility test by redirecting to a page with Javascript. With this feature enabled, only clients that are capable of executing Javascript(mostly browsers) will be allowed to complete the HTTP request. When loadbalancer is configured to do Javascript Challenge, it will redirect the browser to an HTML page on every new HTTP request. This HTML page will have Javascript embedded in it. Loadbalancer chooses a set of random numbers for every new client and sends these numbers along with an encrypted answer with the request such that it embed these numbers as input in the Javascript. Javascript will run on the requestor browser and perform a complex Math operation. Script will submit the answer to loadbalancer. Loadbalancer will validate the answer by comparing the calculated answer with the decrypted answer (which was encrypted when it was sent back as reply) and allow the request to the upstream server only if the answer is correct. Loadbalancer will tag response header with a cookie to avoid Javascript challenge for subsequent requests. Javascript challenge serves following purposes * Validate that the request is coming via a browser that is capable for running Javascript * Force the browser to run a complex operation, f(X), that requires it to spend a large number of CPU cycles. This is to slow down a potential DOS attacker by making it difficult to launch a large request flood without having to spend even larger CPU cost at their end. You can enable either Javascript challenge or Captcha challenge on a virtual host<br>See [Js Challenge Parameters](#enable-challenge-js-challenge-parameters) below.

<a id="enable-challenge-malicious-user-mitigation"></a>&#x2022; [`malicious_user_mitigation`](#enable-challenge-malicious-user-mitigation) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Malicious User Mitigation](#enable-challenge-malicious-user-mitigation) below.

#### Enable Challenge Captcha Challenge Parameters

A [`captcha_challenge_parameters`](#enable-challenge-captcha-challenge-parameters) block (within [`enable_challenge`](#enable-challenge)) supports the following:

<a id="enable-challenge-captcha-challenge-parameters-cookie-expiry"></a>&#x2022; [`cookie_expiry`](#enable-challenge-captcha-challenge-parameters-cookie-expiry) - Optional Number<br>Cookie Expiration Period. Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new challenge

<a id="enable-challenge-captcha-challenge-parameters-custom-page"></a>&#x2022; [`custom_page`](#enable-challenge-captcha-challenge-parameters-custom-page) - Optional String<br>Custom message for Captcha Challenge. Custom message is of type uri_ref. Currently supported URL schemes is string:///. For string:/// scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. 'Please Wait.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '`<p>` Please Wait `</p>`'. Base64 encoded string for this HTML is 'PHA+IFBsZWFzZSBXYWl0IDwvcD4='

#### Enable Challenge Js Challenge Parameters

A [`js_challenge_parameters`](#enable-challenge-js-challenge-parameters) block (within [`enable_challenge`](#enable-challenge)) supports the following:

<a id="enable-challenge-js-challenge-parameters-cookie-expiry"></a>&#x2022; [`cookie_expiry`](#enable-challenge-js-challenge-parameters-cookie-expiry) - Optional Number<br>Cookie Expiration Period. Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new challenge

<a id="enable-challenge-js-challenge-parameters-custom-page"></a>&#x2022; [`custom_page`](#enable-challenge-js-challenge-parameters-custom-page) - Optional String<br>Custom Message for Javascript Challenge. Custom message is of type uri_ref. Currently supported URL schemes is string:///. For string:/// scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. 'Please Wait.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '`<p>` Please Wait `</p>`'. Base64 encoded string for this HTML is 'PHA+IFBsZWFzZSBXYWl0IDwvcD4='

<a id="enable-challenge-js-challenge-parameters-js-script-delay"></a>&#x2022; [`js_script_delay`](#enable-challenge-js-challenge-parameters-js-script-delay) - Optional Number<br>Javascript Delay. Delay introduced by Javascript, in milliseconds

#### Enable Challenge Malicious User Mitigation

A [`malicious_user_mitigation`](#enable-challenge-malicious-user-mitigation) block (within [`enable_challenge`](#enable-challenge)) supports the following:

<a id="enable-challenge-malicious-user-mitigation-name"></a>&#x2022; [`name`](#enable-challenge-malicious-user-mitigation-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="enable-challenge-malicious-user-mitigation-namespace"></a>&#x2022; [`namespace`](#enable-challenge-malicious-user-mitigation-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="enable-challenge-malicious-user-mitigation-tenant"></a>&#x2022; [`tenant`](#enable-challenge-malicious-user-mitigation-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Enable IP Reputation

An [`enable_ip_reputation`](#enable-ip-reputation) block supports the following:

<a id="enable-ip-reputation-ip-threat-categories"></a>&#x2022; [`ip_threat_categories`](#enable-ip-reputation-ip-threat-categories) - Optional List  Defaults to `SPAM_SOURCES`<br>Possible values are `SPAM_SOURCES`, `WINDOWS_EXPLOITS`, `WEB_ATTACKS`, `BOTNETS`, `SCANNERS`, `REPUTATION`, `PHISHING`, `PROXY`, `MOBILE_THREATS`, `TOR_PROXY`, `DENIAL_OF_SERVICE`, `NETWORK`<br>List of IP Threat Categories to choose. If the source IP matches on atleast one of the enabled IP threat categories, the request will be denied

#### Enable Trust Client IP Headers

An [`enable_trust_client_ip_headers`](#enable-trust-client-ip-headers) block supports the following:

<a id="enable-trust-client-ip-headers-client-ip-headers"></a>&#x2022; [`client_ip_headers`](#enable-trust-client-ip-headers-client-ip-headers) - Optional List<br>Client IP Headers. Define the list of one or more Client IP Headers. Headers will be used in order from top to bottom, meaning if the first header is not present in the request, the system will proceed to check for the second header, and so on, until one of the listed headers is found. If none of the defined headers exist, or the value is not an IP address, then the system will use the source IP of the packet. If multiple defined headers with different names are present in the request, the value of the first header name in the configuration will be used. If multiple defined headers with the same name are present in the request, values of all those headers will be combined. The system will read the right-most IP address from header, if there are multiple IP addresses in the header value. For X-Forwarded-For header, the system will read the IP address(rightmost - 1), as the client IP

#### GraphQL Rules

A [`graphql_rules`](#graphql-rules) block supports the following:

<a id="graphql-rules-any-domain"></a>&#x2022; [`any_domain`](#graphql-rules-any-domain) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="graphql-rules-exact-path"></a>&#x2022; [`exact_path`](#graphql-rules-exact-path) - Optional String  Defaults to `/GraphQL`<br>Path. Specifies the exact path to GraphQL endpoint

<a id="graphql-rules-exact-value"></a>&#x2022; [`exact_value`](#graphql-rules-exact-value) - Optional String<br>Exact Value. Exact domain name

<a id="graphql-rules-graphql-settings"></a>&#x2022; [`graphql_settings`](#graphql-rules-graphql-settings) - Optional Block<br>GraphQL Settings. GraphQL configuration<br>See [GraphQL Settings](#graphql-rules-graphql-settings) below.

<a id="graphql-rules-metadata"></a>&#x2022; [`metadata`](#graphql-rules-metadata) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#graphql-rules-metadata) below.

<a id="graphql-rules-method-get"></a>&#x2022; [`method_get`](#graphql-rules-method-get) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="graphql-rules-method-post"></a>&#x2022; [`method_post`](#graphql-rules-method-post) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="graphql-rules-suffix-value"></a>&#x2022; [`suffix_value`](#graphql-rules-suffix-value) - Optional String<br>Suffix Value. Suffix of domain name e.g 'xyz.com' will match '*.xyz.com' and 'xyz.com'

#### GraphQL Rules GraphQL Settings

A [`graphql_settings`](#graphql-rules-graphql-settings) block (within [`graphql_rules`](#graphql-rules)) supports the following:

<a id="graphql-rules-graphql-settings-disable-introspection"></a>&#x2022; [`disable_introspection`](#graphql-rules-graphql-settings-disable-introspection) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="graphql-rules-graphql-settings-enable-introspection"></a>&#x2022; [`enable_introspection`](#graphql-rules-graphql-settings-enable-introspection) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="graphql-rules-graphql-settings-max-batched-queries"></a>&#x2022; [`max_batched_queries`](#graphql-rules-graphql-settings-max-batched-queries) - Optional Number<br>Maximum Batched Queries. Specify maximum number of queries in a single batched request

<a id="graphql-rules-graphql-settings-max-depth"></a>&#x2022; [`max_depth`](#graphql-rules-graphql-settings-max-depth) - Optional Number<br>Maximum Structure Depth. Specify maximum depth for the GraphQL query

<a id="graphql-rules-graphql-settings-max-total-length"></a>&#x2022; [`max_total_length`](#graphql-rules-graphql-settings-max-total-length) - Optional Number<br>Maximum Total Length. Specify maximum length in bytes for the GraphQL query

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

<a id="https-default-header"></a>&#x2022; [`default_header`](#https-default-header) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="https-default-loadbalancer"></a>&#x2022; [`default_loadbalancer`](#https-default-loadbalancer) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="https-disable-path-normalize"></a>&#x2022; [`disable_path_normalize`](#https-disable-path-normalize) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="https-enable-path-normalize"></a>&#x2022; [`enable_path_normalize`](#https-enable-path-normalize) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="https-http-protocol-options"></a>&#x2022; [`http_protocol_options`](#https-http-protocol-options) - Optional Block<br>HTTP Protocol Configuration Options. HTTP protocol configuration options for downstream connections<br>See [HTTP Protocol Options](#https-http-protocol-options) below.

<a id="https-http-redirect"></a>&#x2022; [`http_redirect`](#https-http-redirect) - Optional Bool<br>HTTP Redirect to HTTPS. Redirect HTTP traffic to HTTPS

<a id="https-non-default-loadbalancer"></a>&#x2022; [`non_default_loadbalancer`](#https-non-default-loadbalancer) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="https-pass-through"></a>&#x2022; [`pass_through`](#https-pass-through) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="https-port"></a>&#x2022; [`port`](#https-port) - Optional Number<br>HTTPS Port. HTTPS port to Listen

<a id="https-port-ranges"></a>&#x2022; [`port_ranges`](#https-port-ranges) - Optional String<br>Port Ranges. A string containing a comma separated list of port ranges. Each port range consists of a single port or two ports separated by '-'

<a id="https-server-name"></a>&#x2022; [`server_name`](#https-server-name) - Optional String<br>Modify header value. Define the header value for the header name “server”. This will overwrite existing values, if any, for the server header

<a id="https-tls-cert-params"></a>&#x2022; [`tls_cert_params`](#https-tls-cert-params) - Optional Block<br>TLS Parameters. Select TLS Parameters and Certificates<br>See [TLS Cert Params](#https-tls-cert-params) below.

<a id="https-tls-parameters"></a>&#x2022; [`tls_parameters`](#https-tls-parameters) - Optional Block<br>Inline TLS Parameters. Inline TLS parameters<br>See [TLS Parameters](#https-tls-parameters) below.

#### HTTPS Coalescing Options

A [`coalescing_options`](#https-coalescing-options) block (within [`https`](#https)) supports the following:

<a id="https-coalescing-options-default-coalescing"></a>&#x2022; [`default_coalescing`](#https-coalescing-options-default-coalescing) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="https-coalescing-options-strict-coalescing"></a>&#x2022; [`strict_coalescing`](#https-coalescing-options-strict-coalescing) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### HTTPS HTTP Protocol Options

A [`http_protocol_options`](#https-http-protocol-options) block (within [`https`](#https)) supports the following:

<a id="https-http-protocol-options-http-protocol-enable-v1-only"></a>&#x2022; [`http_protocol_enable_v1_only`](#https-http-protocol-options-http-protocol-enable-v1-only) - Optional Block<br>HTTP/1.1 Protocol Options. HTTP/1.1 Protocol options for downstream connections<br>See [HTTP Protocol Enable V1 Only](#https-http-protocol-options-http-protocol-enable-v1-only) below.

<a id="https-http-protocol-options-http-protocol-enable-v1-v2"></a>&#x2022; [`http_protocol_enable_v1_v2`](#https-http-protocol-options-http-protocol-enable-v1-v2) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="https-http-protocol-options-http-protocol-enable-v2-only"></a>&#x2022; [`http_protocol_enable_v2_only`](#https-http-protocol-options-http-protocol-enable-v2-only) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### HTTPS HTTP Protocol Options HTTP Protocol Enable V1 Only

A [`http_protocol_enable_v1_only`](#https-http-protocol-options-http-protocol-enable-v1-only) block (within [`https.http_protocol_options`](#https-http-protocol-options)) supports the following:

<a id="https-http-protocol-options-http-protocol-enable-v1-only-header-transformation"></a>&#x2022; [`header_transformation`](#https-http-protocol-options-http-protocol-enable-v1-only-header-transformation) - Optional Block<br>Header Transformation. Header Transformation options for HTTP/1.1 request/response headers<br>See [Header Transformation](#https-http-protocol-options-http-protocol-enable-v1-only-header-transformation) below.

#### HTTPS HTTP Protocol Options HTTP Protocol Enable V1 Only Header Transformation

A [`header_transformation`](#https-http-protocol-options-http-protocol-enable-v1-only-header-transformation) block (within [`https.http_protocol_options.http_protocol_enable_v1_only`](#https-http-protocol-options-http-protocol-enable-v1-only)) supports the following:

<a id="https-http-protocol-options-http-protocol-enable-v1-only-header-transformation-default-header-transformation"></a>&#x2022; [`default_header_transformation`](#https-http-protocol-options-http-protocol-enable-v1-only-header-transformation-default-header-transformation) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="https-http-protocol-options-http-protocol-enable-v1-only-header-transformation-legacy-header-transformation"></a>&#x2022; [`legacy_header_transformation`](#https-http-protocol-options-http-protocol-enable-v1-only-header-transformation-legacy-header-transformation) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="https-http-protocol-options-http-protocol-enable-v1-only-header-transformation-preserve-case-header-transformation"></a>&#x2022; [`preserve_case_header_transformation`](#https-http-protocol-options-http-protocol-enable-v1-only-header-transformation-preserve-case-header-transformation) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="https-http-protocol-options-http-protocol-enable-v1-only-header-transformation-proper-case-header-transformation"></a>&#x2022; [`proper_case_header_transformation`](#https-http-protocol-options-http-protocol-enable-v1-only-header-transformation-proper-case-header-transformation) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### HTTPS TLS Cert Params

A [`tls_cert_params`](#https-tls-cert-params) block (within [`https`](#https)) supports the following:

<a id="https-tls-cert-params-certificates"></a>&#x2022; [`certificates`](#https-tls-cert-params-certificates) - Optional Block<br>Certificates. Select one or more certificates with any domain names<br>See [Certificates](#https-tls-cert-params-certificates) below.

<a id="https-tls-cert-params-no-mtls"></a>&#x2022; [`no_mtls`](#https-tls-cert-params-no-mtls) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="https-tls-cert-params-tls-config"></a>&#x2022; [`tls_config`](#https-tls-cert-params-tls-config) - Optional Block<br>TLS Config. This defines various options to configure TLS configuration parameters<br>See [TLS Config](#https-tls-cert-params-tls-config) below.

<a id="https-tls-cert-params-use-mtls"></a>&#x2022; [`use_mtls`](#https-tls-cert-params-use-mtls) - Optional Block<br>Clients TLS validation context. Validation context for downstream client TLS connections<br>See [Use mTLS](#https-tls-cert-params-use-mtls) below.

#### HTTPS TLS Cert Params Certificates

A [`certificates`](#https-tls-cert-params-certificates) block (within [`https.tls_cert_params`](#https-tls-cert-params)) supports the following:

<a id="https-tls-cert-params-certificates-name"></a>&#x2022; [`name`](#https-tls-cert-params-certificates-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="https-tls-cert-params-certificates-namespace"></a>&#x2022; [`namespace`](#https-tls-cert-params-certificates-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="https-tls-cert-params-certificates-tenant"></a>&#x2022; [`tenant`](#https-tls-cert-params-certificates-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### HTTPS TLS Cert Params TLS Config

A [`tls_config`](#https-tls-cert-params-tls-config) block (within [`https.tls_cert_params`](#https-tls-cert-params)) supports the following:

<a id="https-tls-cert-params-tls-config-custom-security"></a>&#x2022; [`custom_security`](#https-tls-cert-params-tls-config-custom-security) - Optional Block<br>Custom Ciphers. This defines TLS protocol config including min/max versions and allowed ciphers<br>See [Custom Security](#https-tls-cert-params-tls-config-custom-security) below.

<a id="https-tls-cert-params-tls-config-default-security"></a>&#x2022; [`default_security`](#https-tls-cert-params-tls-config-default-security) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="https-tls-cert-params-tls-config-low-security"></a>&#x2022; [`low_security`](#https-tls-cert-params-tls-config-low-security) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="https-tls-cert-params-tls-config-medium-security"></a>&#x2022; [`medium_security`](#https-tls-cert-params-tls-config-medium-security) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### HTTPS TLS Cert Params TLS Config Custom Security

A [`custom_security`](#https-tls-cert-params-tls-config-custom-security) block (within [`https.tls_cert_params.tls_config`](#https-tls-cert-params-tls-config)) supports the following:

<a id="https-tls-cert-params-tls-config-custom-security-cipher-suites"></a>&#x2022; [`cipher_suites`](#https-tls-cert-params-tls-config-custom-security-cipher-suites) - Optional List<br>Cipher Suites. The TLS listener will only support the specified cipher list

<a id="https-tls-cert-params-tls-config-custom-security-max-version"></a>&#x2022; [`max_version`](#https-tls-cert-params-tls-config-custom-security-max-version) - Optional String  Defaults to `TLS_AUTO`<br>Possible values are `TLS_AUTO`, `TLSv1_0`, `TLSv1_1`, `TLSv1_2`, `TLSv1_3`<br>TLS Protocol. TlsProtocol is enumeration of supported TLS versions F5 Distributed Cloud will choose the optimal TLS version

<a id="https-tls-cert-params-tls-config-custom-security-min-version"></a>&#x2022; [`min_version`](#https-tls-cert-params-tls-config-custom-security-min-version) - Optional String  Defaults to `TLS_AUTO`<br>Possible values are `TLS_AUTO`, `TLSv1_0`, `TLSv1_1`, `TLSv1_2`, `TLSv1_3`<br>TLS Protocol. TlsProtocol is enumeration of supported TLS versions F5 Distributed Cloud will choose the optimal TLS version

#### HTTPS TLS Cert Params Use mTLS

An [`use_mtls`](#https-tls-cert-params-use-mtls) block (within [`https.tls_cert_params`](#https-tls-cert-params)) supports the following:

<a id="https-tls-cert-params-use-mtls-client-certificate-optional"></a>&#x2022; [`client_certificate_optional`](#https-tls-cert-params-use-mtls-client-certificate-optional) - Optional Bool<br>Client Certificate Optional. Client certificate is optional. If the client has provided a certificate, the load balancer will verify it. If certification verification fails, the connection will be terminated. If the client does not provide a certificate, the connection will be accepted

<a id="https-tls-cert-params-use-mtls-crl"></a>&#x2022; [`crl`](#https-tls-cert-params-use-mtls-crl) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [CRL](#https-tls-cert-params-use-mtls-crl) below.

<a id="https-tls-cert-params-use-mtls-no-crl"></a>&#x2022; [`no_crl`](#https-tls-cert-params-use-mtls-no-crl) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="https-tls-cert-params-use-mtls-trusted-ca"></a>&#x2022; [`trusted_ca`](#https-tls-cert-params-use-mtls-trusted-ca) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Trusted CA](#https-tls-cert-params-use-mtls-trusted-ca) below.

<a id="https-tls-cert-params-use-mtls-trusted-ca-url"></a>&#x2022; [`trusted_ca_url`](#https-tls-cert-params-use-mtls-trusted-ca-url) - Optional String<br>Inline Root CA Certificate (legacy). Upload a Root CA Certificate specifically for this Load Balancer

<a id="https-tls-cert-params-use-mtls-xfcc-disabled"></a>&#x2022; [`xfcc_disabled`](#https-tls-cert-params-use-mtls-xfcc-disabled) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="https-tls-cert-params-use-mtls-xfcc-options"></a>&#x2022; [`xfcc_options`](#https-tls-cert-params-use-mtls-xfcc-options) - Optional Block<br>XFCC Header Elements. X-Forwarded-Client-Cert header elements to be added to requests<br>See [Xfcc Options](#https-tls-cert-params-use-mtls-xfcc-options) below.

#### HTTPS TLS Cert Params Use mTLS CRL

A [`crl`](#https-tls-cert-params-use-mtls-crl) block (within [`https.tls_cert_params.use_mtls`](#https-tls-cert-params-use-mtls)) supports the following:

<a id="https-tls-cert-params-use-mtls-crl-name"></a>&#x2022; [`name`](#https-tls-cert-params-use-mtls-crl-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="https-tls-cert-params-use-mtls-crl-namespace"></a>&#x2022; [`namespace`](#https-tls-cert-params-use-mtls-crl-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="https-tls-cert-params-use-mtls-crl-tenant"></a>&#x2022; [`tenant`](#https-tls-cert-params-use-mtls-crl-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### HTTPS TLS Cert Params Use mTLS Trusted CA

A [`trusted_ca`](#https-tls-cert-params-use-mtls-trusted-ca) block (within [`https.tls_cert_params.use_mtls`](#https-tls-cert-params-use-mtls)) supports the following:

<a id="https-tls-cert-params-use-mtls-trusted-ca-name"></a>&#x2022; [`name`](#https-tls-cert-params-use-mtls-trusted-ca-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="https-tls-cert-params-use-mtls-trusted-ca-namespace"></a>&#x2022; [`namespace`](#https-tls-cert-params-use-mtls-trusted-ca-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="https-tls-cert-params-use-mtls-trusted-ca-tenant"></a>&#x2022; [`tenant`](#https-tls-cert-params-use-mtls-trusted-ca-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### HTTPS TLS Cert Params Use mTLS Xfcc Options

A [`xfcc_options`](#https-tls-cert-params-use-mtls-xfcc-options) block (within [`https.tls_cert_params.use_mtls`](#https-tls-cert-params-use-mtls)) supports the following:

<a id="https-tls-cert-params-use-mtls-xfcc-options-xfcc-header-elements"></a>&#x2022; [`xfcc_header_elements`](#https-tls-cert-params-use-mtls-xfcc-options-xfcc-header-elements) - Optional List  Defaults to `XFCC_NONE`<br>Possible values are `XFCC_NONE`, `XFCC_CERT`, `XFCC_CHAIN`, `XFCC_SUBJECT`, `XFCC_URI`, `XFCC_DNS`<br>XFCC Header Elements. X-Forwarded-Client-Cert header elements to be added to requests

#### HTTPS TLS Parameters

A [`tls_parameters`](#https-tls-parameters) block (within [`https`](#https)) supports the following:

<a id="https-tls-parameters-no-mtls"></a>&#x2022; [`no_mtls`](#https-tls-parameters-no-mtls) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="https-tls-parameters-tls-certificates"></a>&#x2022; [`tls_certificates`](#https-tls-parameters-tls-certificates) - Optional Block<br>TLS Certificates. Users can add one or more certificates that share the same set of domains. for example, domain.com and *.domain.com - but use different signature algorithms<br>See [TLS Certificates](#https-tls-parameters-tls-certificates) below.

<a id="https-tls-parameters-tls-config"></a>&#x2022; [`tls_config`](#https-tls-parameters-tls-config) - Optional Block<br>TLS Config. This defines various options to configure TLS configuration parameters<br>See [TLS Config](#https-tls-parameters-tls-config) below.

<a id="https-tls-parameters-use-mtls"></a>&#x2022; [`use_mtls`](#https-tls-parameters-use-mtls) - Optional Block<br>Clients TLS validation context. Validation context for downstream client TLS connections<br>See [Use mTLS](#https-tls-parameters-use-mtls) below.

#### HTTPS TLS Parameters TLS Certificates

A [`tls_certificates`](#https-tls-parameters-tls-certificates) block (within [`https.tls_parameters`](#https-tls-parameters)) supports the following:

<a id="https-tls-parameters-tls-certificates-certificate-url"></a>&#x2022; [`certificate_url`](#https-tls-parameters-tls-certificates-certificate-url) - Optional String<br>Certificate. TLS certificate. Certificate or certificate chain in PEM format including the PEM headers

<a id="https-tls-parameters-tls-certificates-custom-hash-algorithms"></a>&#x2022; [`custom_hash_algorithms`](#https-tls-parameters-tls-certificates-custom-hash-algorithms) - Optional Block<br>Hash Algorithms. Specifies the hash algorithms to be used<br>See [Custom Hash Algorithms](#https-tls-parameters-tls-certificates-custom-hash-algorithms) below.

<a id="https-tls-parameters-tls-certificates-description-spec"></a>&#x2022; [`description_spec`](#https-tls-parameters-tls-certificates-description-spec) - Optional String<br>Description. Description for the certificate

<a id="https-tls-parameters-tls-certificates-disable-ocsp-stapling"></a>&#x2022; [`disable_ocsp_stapling`](#https-tls-parameters-tls-certificates-disable-ocsp-stapling) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="https-tls-parameters-tls-certificates-private-key"></a>&#x2022; [`private_key`](#https-tls-parameters-tls-certificates-private-key) - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Private Key](#https-tls-parameters-tls-certificates-private-key) below.

<a id="https-tls-parameters-tls-certificates-use-system-defaults"></a>&#x2022; [`use_system_defaults`](#https-tls-parameters-tls-certificates-use-system-defaults) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### HTTPS TLS Parameters TLS Certificates Custom Hash Algorithms

A [`custom_hash_algorithms`](#https-tls-parameters-tls-certificates-custom-hash-algorithms) block (within [`https.tls_parameters.tls_certificates`](#https-tls-parameters-tls-certificates)) supports the following:

<a id="https-tls-parameters-tls-certificates-custom-hash-algorithms-hash-algorithms"></a>&#x2022; [`hash_algorithms`](#https-tls-parameters-tls-certificates-custom-hash-algorithms-hash-algorithms) - Optional List  Defaults to `INVALID_HASH_ALGORITHM`<br>Possible values are `INVALID_HASH_ALGORITHM`, `SHA256`, `SHA1`<br>Hash Algorithms. Ordered list of hash algorithms to be used

#### HTTPS TLS Parameters TLS Certificates Private Key

A [`private_key`](#https-tls-parameters-tls-certificates-private-key) block (within [`https.tls_parameters.tls_certificates`](#https-tls-parameters-tls-certificates)) supports the following:

<a id="https-tls-parameters-tls-certificates-private-key-blindfold-secret-info"></a>&#x2022; [`blindfold_secret_info`](#https-tls-parameters-tls-certificates-private-key-blindfold-secret-info) - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#https-tls-parameters-tls-certificates-private-key-blindfold-secret-info) below.

<a id="https-tls-parameters-tls-certificates-private-key-clear-secret-info"></a>&#x2022; [`clear_secret_info`](#https-tls-parameters-tls-certificates-private-key-clear-secret-info) - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#https-tls-parameters-tls-certificates-private-key-clear-secret-info) below.

#### HTTPS TLS Parameters TLS Certificates Private Key Blindfold Secret Info

A [`blindfold_secret_info`](#https-tls-parameters-tls-certificates-private-key-blindfold-secret-info) block (within [`https.tls_parameters.tls_certificates.private_key`](#https-tls-parameters-tls-certificates-private-key)) supports the following:

<a id="https-tls-parameters-tls-certificates-private-key-blindfold-secret-info-decryption-provider"></a>&#x2022; [`decryption_provider`](#https-tls-parameters-tls-certificates-private-key-blindfold-secret-info-decryption-provider) - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

<a id="https-tls-parameters-tls-certificates-private-key-blindfold-secret-info-location"></a>&#x2022; [`location`](#https-tls-parameters-tls-certificates-private-key-blindfold-secret-info-location) - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

<a id="https-tls-parameters-tls-certificates-private-key-blindfold-secret-info-store-provider"></a>&#x2022; [`store_provider`](#https-tls-parameters-tls-certificates-private-key-blindfold-secret-info-store-provider) - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

#### HTTPS TLS Parameters TLS Certificates Private Key Clear Secret Info

A [`clear_secret_info`](#https-tls-parameters-tls-certificates-private-key-clear-secret-info) block (within [`https.tls_parameters.tls_certificates.private_key`](#https-tls-parameters-tls-certificates-private-key)) supports the following:

<a id="https-tls-parameters-tls-certificates-private-key-clear-secret-info-provider-ref"></a>&#x2022; [`provider_ref`](#https-tls-parameters-tls-certificates-private-key-clear-secret-info-provider-ref) - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

<a id="https-tls-parameters-tls-certificates-private-key-clear-secret-info-url"></a>&#x2022; [`url`](#https-tls-parameters-tls-certificates-private-key-clear-secret-info-url) - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

#### HTTPS TLS Parameters TLS Config

A [`tls_config`](#https-tls-parameters-tls-config) block (within [`https.tls_parameters`](#https-tls-parameters)) supports the following:

<a id="https-tls-parameters-tls-config-custom-security"></a>&#x2022; [`custom_security`](#https-tls-parameters-tls-config-custom-security) - Optional Block<br>Custom Ciphers. This defines TLS protocol config including min/max versions and allowed ciphers<br>See [Custom Security](#https-tls-parameters-tls-config-custom-security) below.

<a id="https-tls-parameters-tls-config-default-security"></a>&#x2022; [`default_security`](#https-tls-parameters-tls-config-default-security) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="https-tls-parameters-tls-config-low-security"></a>&#x2022; [`low_security`](#https-tls-parameters-tls-config-low-security) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="https-tls-parameters-tls-config-medium-security"></a>&#x2022; [`medium_security`](#https-tls-parameters-tls-config-medium-security) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### HTTPS TLS Parameters TLS Config Custom Security

A [`custom_security`](#https-tls-parameters-tls-config-custom-security) block (within [`https.tls_parameters.tls_config`](#https-tls-parameters-tls-config)) supports the following:

<a id="https-tls-parameters-tls-config-custom-security-cipher-suites"></a>&#x2022; [`cipher_suites`](#https-tls-parameters-tls-config-custom-security-cipher-suites) - Optional List<br>Cipher Suites. The TLS listener will only support the specified cipher list

<a id="https-tls-parameters-tls-config-custom-security-max-version"></a>&#x2022; [`max_version`](#https-tls-parameters-tls-config-custom-security-max-version) - Optional String  Defaults to `TLS_AUTO`<br>Possible values are `TLS_AUTO`, `TLSv1_0`, `TLSv1_1`, `TLSv1_2`, `TLSv1_3`<br>TLS Protocol. TlsProtocol is enumeration of supported TLS versions F5 Distributed Cloud will choose the optimal TLS version

<a id="https-tls-parameters-tls-config-custom-security-min-version"></a>&#x2022; [`min_version`](#https-tls-parameters-tls-config-custom-security-min-version) - Optional String  Defaults to `TLS_AUTO`<br>Possible values are `TLS_AUTO`, `TLSv1_0`, `TLSv1_1`, `TLSv1_2`, `TLSv1_3`<br>TLS Protocol. TlsProtocol is enumeration of supported TLS versions F5 Distributed Cloud will choose the optimal TLS version

#### HTTPS TLS Parameters Use mTLS

An [`use_mtls`](#https-tls-parameters-use-mtls) block (within [`https.tls_parameters`](#https-tls-parameters)) supports the following:

<a id="https-tls-parameters-use-mtls-client-certificate-optional"></a>&#x2022; [`client_certificate_optional`](#https-tls-parameters-use-mtls-client-certificate-optional) - Optional Bool<br>Client Certificate Optional. Client certificate is optional. If the client has provided a certificate, the load balancer will verify it. If certification verification fails, the connection will be terminated. If the client does not provide a certificate, the connection will be accepted

<a id="https-tls-parameters-use-mtls-crl"></a>&#x2022; [`crl`](#https-tls-parameters-use-mtls-crl) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [CRL](#https-tls-parameters-use-mtls-crl) below.

<a id="https-tls-parameters-use-mtls-no-crl"></a>&#x2022; [`no_crl`](#https-tls-parameters-use-mtls-no-crl) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="https-tls-parameters-use-mtls-trusted-ca"></a>&#x2022; [`trusted_ca`](#https-tls-parameters-use-mtls-trusted-ca) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Trusted CA](#https-tls-parameters-use-mtls-trusted-ca) below.

<a id="https-tls-parameters-use-mtls-trusted-ca-url"></a>&#x2022; [`trusted_ca_url`](#https-tls-parameters-use-mtls-trusted-ca-url) - Optional String<br>Inline Root CA Certificate (legacy). Upload a Root CA Certificate specifically for this Load Balancer

<a id="https-tls-parameters-use-mtls-xfcc-disabled"></a>&#x2022; [`xfcc_disabled`](#https-tls-parameters-use-mtls-xfcc-disabled) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="https-tls-parameters-use-mtls-xfcc-options"></a>&#x2022; [`xfcc_options`](#https-tls-parameters-use-mtls-xfcc-options) - Optional Block<br>XFCC Header Elements. X-Forwarded-Client-Cert header elements to be added to requests<br>See [Xfcc Options](#https-tls-parameters-use-mtls-xfcc-options) below.

#### HTTPS TLS Parameters Use mTLS CRL

A [`crl`](#https-tls-parameters-use-mtls-crl) block (within [`https.tls_parameters.use_mtls`](#https-tls-parameters-use-mtls)) supports the following:

<a id="https-tls-parameters-use-mtls-crl-name"></a>&#x2022; [`name`](#https-tls-parameters-use-mtls-crl-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="https-tls-parameters-use-mtls-crl-namespace"></a>&#x2022; [`namespace`](#https-tls-parameters-use-mtls-crl-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="https-tls-parameters-use-mtls-crl-tenant"></a>&#x2022; [`tenant`](#https-tls-parameters-use-mtls-crl-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### HTTPS TLS Parameters Use mTLS Trusted CA

A [`trusted_ca`](#https-tls-parameters-use-mtls-trusted-ca) block (within [`https.tls_parameters.use_mtls`](#https-tls-parameters-use-mtls)) supports the following:

<a id="https-tls-parameters-use-mtls-trusted-ca-name"></a>&#x2022; [`name`](#https-tls-parameters-use-mtls-trusted-ca-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="https-tls-parameters-use-mtls-trusted-ca-namespace"></a>&#x2022; [`namespace`](#https-tls-parameters-use-mtls-trusted-ca-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="https-tls-parameters-use-mtls-trusted-ca-tenant"></a>&#x2022; [`tenant`](#https-tls-parameters-use-mtls-trusted-ca-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### HTTPS TLS Parameters Use mTLS Xfcc Options

A [`xfcc_options`](#https-tls-parameters-use-mtls-xfcc-options) block (within [`https.tls_parameters.use_mtls`](#https-tls-parameters-use-mtls)) supports the following:

<a id="https-tls-parameters-use-mtls-xfcc-options-xfcc-header-elements"></a>&#x2022; [`xfcc_header_elements`](#https-tls-parameters-use-mtls-xfcc-options-xfcc-header-elements) - Optional List  Defaults to `XFCC_NONE`<br>Possible values are `XFCC_NONE`, `XFCC_CERT`, `XFCC_CHAIN`, `XFCC_SUBJECT`, `XFCC_URI`, `XFCC_DNS`<br>XFCC Header Elements. X-Forwarded-Client-Cert header elements to be added to requests

#### HTTPS Auto Cert

A [`https_auto_cert`](#https-auto-cert) block supports the following:

<a id="https-auto-cert-add-hsts"></a>&#x2022; [`add_hsts`](#https-auto-cert-add-hsts) - Optional Bool<br>Add HSTS Header. Add HTTP Strict-Transport-Security response header

<a id="https-auto-cert-append-server-name"></a>&#x2022; [`append_server_name`](#https-auto-cert-append-server-name) - Optional String<br>Append header value. Define the header value for the header name “server”. If header value is already present, it is not overwritten and passed as-is

<a id="https-auto-cert-coalescing-options"></a>&#x2022; [`coalescing_options`](#https-auto-cert-coalescing-options) - Optional Block<br>TLS Coalescing Options. TLS connection coalescing configuration (not compatible with mTLS)<br>See [Coalescing Options](#https-auto-cert-coalescing-options) below.

<a id="https-auto-cert-connection-idle-timeout"></a>&#x2022; [`connection_idle_timeout`](#https-auto-cert-connection-idle-timeout) - Optional Number  Defaults to `2`  Specified in milliseconds<br>Connection Idle Timeout. The idle timeout for downstream connections. The idle timeout is defined as the period in which there are no active requests. When the idle timeout is reached the connection will be closed. Note that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive.  The minutes

<a id="https-auto-cert-default-header"></a>&#x2022; [`default_header`](#https-auto-cert-default-header) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="https-auto-cert-default-loadbalancer"></a>&#x2022; [`default_loadbalancer`](#https-auto-cert-default-loadbalancer) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="https-auto-cert-disable-path-normalize"></a>&#x2022; [`disable_path_normalize`](#https-auto-cert-disable-path-normalize) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="https-auto-cert-enable-path-normalize"></a>&#x2022; [`enable_path_normalize`](#https-auto-cert-enable-path-normalize) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="https-auto-cert-http-protocol-options"></a>&#x2022; [`http_protocol_options`](#https-auto-cert-http-protocol-options) - Optional Block<br>HTTP Protocol Configuration Options. HTTP protocol configuration options for downstream connections<br>See [HTTP Protocol Options](#https-auto-cert-http-protocol-options) below.

<a id="https-auto-cert-http-redirect"></a>&#x2022; [`http_redirect`](#https-auto-cert-http-redirect) - Optional Bool<br>HTTP Redirect to HTTPS. Redirect HTTP traffic to HTTPS

<a id="https-auto-cert-no-mtls"></a>&#x2022; [`no_mtls`](#https-auto-cert-no-mtls) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="https-auto-cert-non-default-loadbalancer"></a>&#x2022; [`non_default_loadbalancer`](#https-auto-cert-non-default-loadbalancer) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="https-auto-cert-pass-through"></a>&#x2022; [`pass_through`](#https-auto-cert-pass-through) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="https-auto-cert-port"></a>&#x2022; [`port`](#https-auto-cert-port) - Optional Number<br>HTTPS Listen Port. HTTPS port to Listen

<a id="https-auto-cert-port-ranges"></a>&#x2022; [`port_ranges`](#https-auto-cert-port-ranges) - Optional String<br>Port Ranges. A string containing a comma separated list of port ranges. Each port range consists of a single port or two ports separated by '-'

<a id="https-auto-cert-server-name"></a>&#x2022; [`server_name`](#https-auto-cert-server-name) - Optional String<br>Modify header value. Define the header value for the header name “server”. This will overwrite existing values, if any, for the server header

<a id="https-auto-cert-tls-config"></a>&#x2022; [`tls_config`](#https-auto-cert-tls-config) - Optional Block<br>TLS Config. This defines various options to configure TLS configuration parameters<br>See [TLS Config](#https-auto-cert-tls-config) below.

<a id="https-auto-cert-use-mtls"></a>&#x2022; [`use_mtls`](#https-auto-cert-use-mtls) - Optional Block<br>Clients TLS validation context. Validation context for downstream client TLS connections<br>See [Use mTLS](#https-auto-cert-use-mtls) below.

#### HTTPS Auto Cert Coalescing Options

A [`coalescing_options`](#https-auto-cert-coalescing-options) block (within [`https_auto_cert`](#https-auto-cert)) supports the following:

<a id="https-auto-cert-coalescing-options-default-coalescing"></a>&#x2022; [`default_coalescing`](#https-auto-cert-coalescing-options-default-coalescing) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="https-auto-cert-coalescing-options-strict-coalescing"></a>&#x2022; [`strict_coalescing`](#https-auto-cert-coalescing-options-strict-coalescing) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### HTTPS Auto Cert HTTP Protocol Options

A [`http_protocol_options`](#https-auto-cert-http-protocol-options) block (within [`https_auto_cert`](#https-auto-cert)) supports the following:

<a id="https-auto-cert-http-protocol-options-http-protocol-enable-v1-only"></a>&#x2022; [`http_protocol_enable_v1_only`](#https-auto-cert-http-protocol-options-http-protocol-enable-v1-only) - Optional Block<br>HTTP/1.1 Protocol Options. HTTP/1.1 Protocol options for downstream connections<br>See [HTTP Protocol Enable V1 Only](#https-auto-cert-http-protocol-options-http-protocol-enable-v1-only) below.

<a id="https-auto-cert-http-protocol-options-http-protocol-enable-v1-v2"></a>&#x2022; [`http_protocol_enable_v1_v2`](#https-auto-cert-http-protocol-options-http-protocol-enable-v1-v2) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="https-auto-cert-http-protocol-options-http-protocol-enable-v2-only"></a>&#x2022; [`http_protocol_enable_v2_only`](#https-auto-cert-http-protocol-options-http-protocol-enable-v2-only) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### HTTPS Auto Cert HTTP Protocol Options HTTP Protocol Enable V1 Only

A [`http_protocol_enable_v1_only`](#https-auto-cert-http-protocol-options-http-protocol-enable-v1-only) block (within [`https_auto_cert.http_protocol_options`](#https-auto-cert-http-protocol-options)) supports the following:

<a id="https-auto-cert-http-protocol-options-http-protocol-enable-v1-only-header-transformation"></a>&#x2022; [`header_transformation`](#https-auto-cert-http-protocol-options-http-protocol-enable-v1-only-header-transformation) - Optional Block<br>Header Transformation. Header Transformation options for HTTP/1.1 request/response headers<br>See [Header Transformation](#https-auto-cert-http-protocol-options-http-protocol-enable-v1-only-header-transformation) below.

#### HTTPS Auto Cert HTTP Protocol Options HTTP Protocol Enable V1 Only Header Transformation

A [`header_transformation`](#https-auto-cert-http-protocol-options-http-protocol-enable-v1-only-header-transformation) block (within [`https_auto_cert.http_protocol_options.http_protocol_enable_v1_only`](#https-auto-cert-http-protocol-options-http-protocol-enable-v1-only)) supports the following:

<a id="https-auto-cert-http-protocol-options-http-protocol-enable-v1-only-header-transformation-default-header-transformation"></a>&#x2022; [`default_header_transformation`](#https-auto-cert-http-protocol-options-http-protocol-enable-v1-only-header-transformation-default-header-transformation) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="https-auto-cert-http-protocol-options-http-protocol-enable-v1-only-header-transformation-legacy-header-transformation"></a>&#x2022; [`legacy_header_transformation`](#https-auto-cert-http-protocol-options-http-protocol-enable-v1-only-header-transformation-legacy-header-transformation) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="https-auto-cert-http-protocol-options-http-protocol-enable-v1-only-header-transformation-preserve-case-header-transformation"></a>&#x2022; [`preserve_case_header_transformation`](#https-auto-cert-http-protocol-options-http-protocol-enable-v1-only-header-transformation-preserve-case-header-transformation) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="https-auto-cert-http-protocol-options-http-protocol-enable-v1-only-header-transformation-proper-case-header-transformation"></a>&#x2022; [`proper_case_header_transformation`](#https-auto-cert-http-protocol-options-http-protocol-enable-v1-only-header-transformation-proper-case-header-transformation) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### HTTPS Auto Cert TLS Config

A [`tls_config`](#https-auto-cert-tls-config) block (within [`https_auto_cert`](#https-auto-cert)) supports the following:

<a id="https-auto-cert-tls-config-custom-security"></a>&#x2022; [`custom_security`](#https-auto-cert-tls-config-custom-security) - Optional Block<br>Custom Ciphers. This defines TLS protocol config including min/max versions and allowed ciphers<br>See [Custom Security](#https-auto-cert-tls-config-custom-security) below.

<a id="https-auto-cert-tls-config-default-security"></a>&#x2022; [`default_security`](#https-auto-cert-tls-config-default-security) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="https-auto-cert-tls-config-low-security"></a>&#x2022; [`low_security`](#https-auto-cert-tls-config-low-security) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="https-auto-cert-tls-config-medium-security"></a>&#x2022; [`medium_security`](#https-auto-cert-tls-config-medium-security) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### HTTPS Auto Cert TLS Config Custom Security

A [`custom_security`](#https-auto-cert-tls-config-custom-security) block (within [`https_auto_cert.tls_config`](#https-auto-cert-tls-config)) supports the following:

<a id="https-auto-cert-tls-config-custom-security-cipher-suites"></a>&#x2022; [`cipher_suites`](#https-auto-cert-tls-config-custom-security-cipher-suites) - Optional List<br>Cipher Suites. The TLS listener will only support the specified cipher list

<a id="https-auto-cert-tls-config-custom-security-max-version"></a>&#x2022; [`max_version`](#https-auto-cert-tls-config-custom-security-max-version) - Optional String  Defaults to `TLS_AUTO`<br>Possible values are `TLS_AUTO`, `TLSv1_0`, `TLSv1_1`, `TLSv1_2`, `TLSv1_3`<br>TLS Protocol. TlsProtocol is enumeration of supported TLS versions F5 Distributed Cloud will choose the optimal TLS version

<a id="https-auto-cert-tls-config-custom-security-min-version"></a>&#x2022; [`min_version`](#https-auto-cert-tls-config-custom-security-min-version) - Optional String  Defaults to `TLS_AUTO`<br>Possible values are `TLS_AUTO`, `TLSv1_0`, `TLSv1_1`, `TLSv1_2`, `TLSv1_3`<br>TLS Protocol. TlsProtocol is enumeration of supported TLS versions F5 Distributed Cloud will choose the optimal TLS version

#### HTTPS Auto Cert Use mTLS

An [`use_mtls`](#https-auto-cert-use-mtls) block (within [`https_auto_cert`](#https-auto-cert)) supports the following:

<a id="https-auto-cert-use-mtls-client-certificate-optional"></a>&#x2022; [`client_certificate_optional`](#https-auto-cert-use-mtls-client-certificate-optional) - Optional Bool<br>Client Certificate Optional. Client certificate is optional. If the client has provided a certificate, the load balancer will verify it. If certification verification fails, the connection will be terminated. If the client does not provide a certificate, the connection will be accepted

<a id="https-auto-cert-use-mtls-crl"></a>&#x2022; [`crl`](#https-auto-cert-use-mtls-crl) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [CRL](#https-auto-cert-use-mtls-crl) below.

<a id="https-auto-cert-use-mtls-no-crl"></a>&#x2022; [`no_crl`](#https-auto-cert-use-mtls-no-crl) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="https-auto-cert-use-mtls-trusted-ca"></a>&#x2022; [`trusted_ca`](#https-auto-cert-use-mtls-trusted-ca) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Trusted CA](#https-auto-cert-use-mtls-trusted-ca) below.

<a id="https-auto-cert-use-mtls-trusted-ca-url"></a>&#x2022; [`trusted_ca_url`](#https-auto-cert-use-mtls-trusted-ca-url) - Optional String<br>Inline Root CA Certificate (legacy). Upload a Root CA Certificate specifically for this Load Balancer

<a id="https-auto-cert-use-mtls-xfcc-disabled"></a>&#x2022; [`xfcc_disabled`](#https-auto-cert-use-mtls-xfcc-disabled) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="https-auto-cert-use-mtls-xfcc-options"></a>&#x2022; [`xfcc_options`](#https-auto-cert-use-mtls-xfcc-options) - Optional Block<br>XFCC Header Elements. X-Forwarded-Client-Cert header elements to be added to requests<br>See [Xfcc Options](#https-auto-cert-use-mtls-xfcc-options) below.

#### HTTPS Auto Cert Use mTLS CRL

A [`crl`](#https-auto-cert-use-mtls-crl) block (within [`https_auto_cert.use_mtls`](#https-auto-cert-use-mtls)) supports the following:

<a id="https-auto-cert-use-mtls-crl-name"></a>&#x2022; [`name`](#https-auto-cert-use-mtls-crl-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="https-auto-cert-use-mtls-crl-namespace"></a>&#x2022; [`namespace`](#https-auto-cert-use-mtls-crl-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="https-auto-cert-use-mtls-crl-tenant"></a>&#x2022; [`tenant`](#https-auto-cert-use-mtls-crl-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### HTTPS Auto Cert Use mTLS Trusted CA

A [`trusted_ca`](#https-auto-cert-use-mtls-trusted-ca) block (within [`https_auto_cert.use_mtls`](#https-auto-cert-use-mtls)) supports the following:

<a id="https-auto-cert-use-mtls-trusted-ca-name"></a>&#x2022; [`name`](#https-auto-cert-use-mtls-trusted-ca-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="https-auto-cert-use-mtls-trusted-ca-namespace"></a>&#x2022; [`namespace`](#https-auto-cert-use-mtls-trusted-ca-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="https-auto-cert-use-mtls-trusted-ca-tenant"></a>&#x2022; [`tenant`](#https-auto-cert-use-mtls-trusted-ca-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### HTTPS Auto Cert Use mTLS Xfcc Options

A [`xfcc_options`](#https-auto-cert-use-mtls-xfcc-options) block (within [`https_auto_cert.use_mtls`](#https-auto-cert-use-mtls)) supports the following:

<a id="https-auto-cert-use-mtls-xfcc-options-xfcc-header-elements"></a>&#x2022; [`xfcc_header_elements`](#https-auto-cert-use-mtls-xfcc-options-xfcc-header-elements) - Optional List  Defaults to `XFCC_NONE`<br>Possible values are `XFCC_NONE`, `XFCC_CERT`, `XFCC_CHAIN`, `XFCC_SUBJECT`, `XFCC_URI`, `XFCC_DNS`<br>XFCC Header Elements. X-Forwarded-Client-Cert header elements to be added to requests

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

<a id="jwt-validation-action-block"></a>&#x2022; [`block`](#jwt-validation-action-block) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="jwt-validation-action-report"></a>&#x2022; [`report`](#jwt-validation-action-report) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### JWT Validation Jwks Config

A [`jwks_config`](#jwt-validation-jwks-config) block (within [`jwt_validation`](#jwt-validation)) supports the following:

<a id="jwt-validation-jwks-config-cleartext"></a>&#x2022; [`cleartext`](#jwt-validation-jwks-config-cleartext) - Optional String<br>JSON Web Key Set (JWKS). The JSON Web Key Set (JWKS) is a set of keys used to verify JSON Web Token (JWT) issued by the Authorization Server. See RFC 7517 for more details

#### JWT Validation Mandatory Claims

A [`mandatory_claims`](#jwt-validation-mandatory-claims) block (within [`jwt_validation`](#jwt-validation)) supports the following:

<a id="jwt-validation-mandatory-claims-claim-names"></a>&#x2022; [`claim_names`](#jwt-validation-mandatory-claims-claim-names) - Optional List<br>Claim Names

#### JWT Validation Reserved Claims

A [`reserved_claims`](#jwt-validation-reserved-claims) block (within [`jwt_validation`](#jwt-validation)) supports the following:

<a id="jwt-validation-reserved-claims-audience"></a>&#x2022; [`audience`](#jwt-validation-reserved-claims-audience) - Optional Block<br>Audiences<br>See [Audience](#jwt-validation-reserved-claims-audience) below.

<a id="jwt-validation-reserved-claims-audience-disable"></a>&#x2022; [`audience_disable`](#jwt-validation-reserved-claims-audience-disable) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="jwt-validation-reserved-claims-issuer"></a>&#x2022; [`issuer`](#jwt-validation-reserved-claims-issuer) - Optional String<br>Exact Match

<a id="jwt-validation-reserved-claims-issuer-disable"></a>&#x2022; [`issuer_disable`](#jwt-validation-reserved-claims-issuer-disable) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="jwt-validation-reserved-claims-validate-period-disable"></a>&#x2022; [`validate_period_disable`](#jwt-validation-reserved-claims-validate-period-disable) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="jwt-validation-reserved-claims-validate-period-enable"></a>&#x2022; [`validate_period_enable`](#jwt-validation-reserved-claims-validate-period-enable) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### JWT Validation Reserved Claims Audience

An [`audience`](#jwt-validation-reserved-claims-audience) block (within [`jwt_validation.reserved_claims`](#jwt-validation-reserved-claims)) supports the following:

<a id="jwt-validation-reserved-claims-audience-audiences"></a>&#x2022; [`audiences`](#jwt-validation-reserved-claims-audience-audiences) - Optional List<br>Values

#### JWT Validation Target

A [`target`](#jwt-validation-target) block (within [`jwt_validation`](#jwt-validation)) supports the following:

<a id="jwt-validation-target-all-endpoint"></a>&#x2022; [`all_endpoint`](#jwt-validation-target-all-endpoint) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="jwt-validation-target-api-groups"></a>&#x2022; [`api_groups`](#jwt-validation-target-api-groups) - Optional Block<br>API Groups<br>See [API Groups](#jwt-validation-target-api-groups) below.

<a id="jwt-validation-target-base-paths"></a>&#x2022; [`base_paths`](#jwt-validation-target-base-paths) - Optional Block<br>Base Paths<br>See [Base Paths](#jwt-validation-target-base-paths) below.

#### JWT Validation Target API Groups

An [`api_groups`](#jwt-validation-target-api-groups) block (within [`jwt_validation.target`](#jwt-validation-target)) supports the following:

<a id="jwt-validation-target-api-groups-api-groups"></a>&#x2022; [`api_groups`](#jwt-validation-target-api-groups-api-groups) - Optional List<br>API Groups

#### JWT Validation Target Base Paths

A [`base_paths`](#jwt-validation-target-base-paths) block (within [`jwt_validation.target`](#jwt-validation-target)) supports the following:

<a id="jwt-validation-target-base-paths-base-paths"></a>&#x2022; [`base_paths`](#jwt-validation-target-base-paths-base-paths) - Optional List<br>Prefix Values

#### JWT Validation Token Location

A [`token_location`](#jwt-validation-token-location) block (within [`jwt_validation`](#jwt-validation)) supports the following:

<a id="jwt-validation-token-location-bearer-token"></a>&#x2022; [`bearer_token`](#jwt-validation-token-location-bearer-token) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### L7 DDOS Action Js Challenge

A [`l7_ddos_action_js_challenge`](#l7-ddos-action-js-challenge) block supports the following:

<a id="l7-ddos-action-js-challenge-cookie-expiry"></a>&#x2022; [`cookie_expiry`](#l7-ddos-action-js-challenge-cookie-expiry) - Optional Number<br>Cookie Expiration Period. Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new challenge

<a id="l7-ddos-action-js-challenge-custom-page"></a>&#x2022; [`custom_page`](#l7-ddos-action-js-challenge-custom-page) - Optional String<br>Custom Message for Javascript Challenge. Custom message is of type uri_ref. Currently supported URL schemes is string:///. For string:/// scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. 'Please Wait.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '`<p>` Please Wait `</p>`'. Base64 encoded string for this HTML is 'PHA+IFBsZWFzZSBXYWl0IDwvcD4='

<a id="l7-ddos-action-js-challenge-js-script-delay"></a>&#x2022; [`js_script_delay`](#l7-ddos-action-js-challenge-js-script-delay) - Optional Number<br>Javascript Delay. Delay introduced by Javascript, in milliseconds

#### L7 DDOS Protection

A [`l7_ddos_protection`](#l7-ddos-protection) block supports the following:

<a id="l7-ddos-protection-clientside-action-captcha-challenge"></a>&#x2022; [`clientside_action_captcha_challenge`](#l7-ddos-protection-clientside-action-captcha-challenge) - Optional Block<br>Captcha Challenge Parameters. Enables loadbalancer to perform captcha challenge Captcha challenge will be based on Google Recaptcha. With this feature enabled, only clients that pass the captcha challenge will be allowed to complete the HTTP request. When loadbalancer is configured to do Captcha Challenge, it will redirect the browser to an HTML page on every new HTTP request. This HTML page will have captcha challenge embedded in it. Client will be allowed to make the request only if the captcha challenge is successful. Loadbalancer will tag response header with a cookie to avoid Captcha challenge for subsequent requests. CAPTCHA is mainly used as a security check to ensure only human users can pass through. Generally, computers or bots are not capable of solving a captcha. You can enable either Javascript challenge or Captcha challenge on a virtual host<br>See [Clientside Action Captcha Challenge](#l7-ddos-protection-clientside-action-captcha-challenge) below.

<a id="l7-ddos-protection-clientside-action-js-challenge"></a>&#x2022; [`clientside_action_js_challenge`](#l7-ddos-protection-clientside-action-js-challenge) - Optional Block<br>Javascript Challenge Parameters. Enables loadbalancer to perform client browser compatibility test by redirecting to a page with Javascript. With this feature enabled, only clients that are capable of executing Javascript(mostly browsers) will be allowed to complete the HTTP request. When loadbalancer is configured to do Javascript Challenge, it will redirect the browser to an HTML page on every new HTTP request. This HTML page will have Javascript embedded in it. Loadbalancer chooses a set of random numbers for every new client and sends these numbers along with an encrypted answer with the request such that it embed these numbers as input in the Javascript. Javascript will run on the requestor browser and perform a complex Math operation. Script will submit the answer to loadbalancer. Loadbalancer will validate the answer by comparing the calculated answer with the decrypted answer (which was encrypted when it was sent back as reply) and allow the request to the upstream server only if the answer is correct. Loadbalancer will tag response header with a cookie to avoid Javascript challenge for subsequent requests. Javascript challenge serves following purposes * Validate that the request is coming via a browser that is capable for running Javascript * Force the browser to run a complex operation, f(X), that requires it to spend a large number of CPU cycles. This is to slow down a potential DOS attacker by making it difficult to launch a large request flood without having to spend even larger CPU cost at their end. You can enable either Javascript challenge or Captcha challenge on a virtual host<br>See [Clientside Action Js Challenge](#l7-ddos-protection-clientside-action-js-challenge) below.

<a id="l7-ddos-protection-clientside-action-none"></a>&#x2022; [`clientside_action_none`](#l7-ddos-protection-clientside-action-none) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="l7-ddos-protection-ddos-policy-custom"></a>&#x2022; [`ddos_policy_custom`](#l7-ddos-protection-ddos-policy-custom) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [DDOS Policy Custom](#l7-ddos-protection-ddos-policy-custom) below.

<a id="l7-ddos-protection-ddos-policy-none"></a>&#x2022; [`ddos_policy_none`](#l7-ddos-protection-ddos-policy-none) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="l7-ddos-protection-default-rps-threshold"></a>&#x2022; [`default_rps_threshold`](#l7-ddos-protection-default-rps-threshold) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="l7-ddos-protection-mitigation-block"></a>&#x2022; [`mitigation_block`](#l7-ddos-protection-mitigation-block) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="l7-ddos-protection-mitigation-captcha-challenge"></a>&#x2022; [`mitigation_captcha_challenge`](#l7-ddos-protection-mitigation-captcha-challenge) - Optional Block<br>Captcha Challenge Parameters. Enables loadbalancer to perform captcha challenge Captcha challenge will be based on Google Recaptcha. With this feature enabled, only clients that pass the captcha challenge will be allowed to complete the HTTP request. When loadbalancer is configured to do Captcha Challenge, it will redirect the browser to an HTML page on every new HTTP request. This HTML page will have captcha challenge embedded in it. Client will be allowed to make the request only if the captcha challenge is successful. Loadbalancer will tag response header with a cookie to avoid Captcha challenge for subsequent requests. CAPTCHA is mainly used as a security check to ensure only human users can pass through. Generally, computers or bots are not capable of solving a captcha. You can enable either Javascript challenge or Captcha challenge on a virtual host<br>See [Mitigation Captcha Challenge](#l7-ddos-protection-mitigation-captcha-challenge) below.

<a id="l7-ddos-protection-mitigation-js-challenge"></a>&#x2022; [`mitigation_js_challenge`](#l7-ddos-protection-mitigation-js-challenge) - Optional Block<br>Javascript Challenge Parameters. Enables loadbalancer to perform client browser compatibility test by redirecting to a page with Javascript. With this feature enabled, only clients that are capable of executing Javascript(mostly browsers) will be allowed to complete the HTTP request. When loadbalancer is configured to do Javascript Challenge, it will redirect the browser to an HTML page on every new HTTP request. This HTML page will have Javascript embedded in it. Loadbalancer chooses a set of random numbers for every new client and sends these numbers along with an encrypted answer with the request such that it embed these numbers as input in the Javascript. Javascript will run on the requestor browser and perform a complex Math operation. Script will submit the answer to loadbalancer. Loadbalancer will validate the answer by comparing the calculated answer with the decrypted answer (which was encrypted when it was sent back as reply) and allow the request to the upstream server only if the answer is correct. Loadbalancer will tag response header with a cookie to avoid Javascript challenge for subsequent requests. Javascript challenge serves following purposes * Validate that the request is coming via a browser that is capable for running Javascript * Force the browser to run a complex operation, f(X), that requires it to spend a large number of CPU cycles. This is to slow down a potential DOS attacker by making it difficult to launch a large request flood without having to spend even larger CPU cost at their end. You can enable either Javascript challenge or Captcha challenge on a virtual host<br>See [Mitigation Js Challenge](#l7-ddos-protection-mitigation-js-challenge) below.

<a id="l7-ddos-protection-rps-threshold"></a>&#x2022; [`rps_threshold`](#l7-ddos-protection-rps-threshold) - Optional Number<br>Custom. Configure custom RPS threshold

#### L7 DDOS Protection Clientside Action Captcha Challenge

A [`clientside_action_captcha_challenge`](#l7-ddos-protection-clientside-action-captcha-challenge) block (within [`l7_ddos_protection`](#l7-ddos-protection)) supports the following:

<a id="l7-ddos-protection-clientside-action-captcha-challenge-cookie-expiry"></a>&#x2022; [`cookie_expiry`](#l7-ddos-protection-clientside-action-captcha-challenge-cookie-expiry) - Optional Number<br>Cookie Expiration Period. Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new challenge

<a id="l7-ddos-protection-clientside-action-captcha-challenge-custom-page"></a>&#x2022; [`custom_page`](#l7-ddos-protection-clientside-action-captcha-challenge-custom-page) - Optional String<br>Custom message for Captcha Challenge. Custom message is of type uri_ref. Currently supported URL schemes is string:///. For string:/// scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. 'Please Wait.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '`<p>` Please Wait `</p>`'. Base64 encoded string for this HTML is 'PHA+IFBsZWFzZSBXYWl0IDwvcD4='

#### L7 DDOS Protection Clientside Action Js Challenge

A [`clientside_action_js_challenge`](#l7-ddos-protection-clientside-action-js-challenge) block (within [`l7_ddos_protection`](#l7-ddos-protection)) supports the following:

<a id="l7-ddos-protection-clientside-action-js-challenge-cookie-expiry"></a>&#x2022; [`cookie_expiry`](#l7-ddos-protection-clientside-action-js-challenge-cookie-expiry) - Optional Number<br>Cookie Expiration Period. Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new challenge

<a id="l7-ddos-protection-clientside-action-js-challenge-custom-page"></a>&#x2022; [`custom_page`](#l7-ddos-protection-clientside-action-js-challenge-custom-page) - Optional String<br>Custom Message for Javascript Challenge. Custom message is of type uri_ref. Currently supported URL schemes is string:///. For string:/// scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. 'Please Wait.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '`<p>` Please Wait `</p>`'. Base64 encoded string for this HTML is 'PHA+IFBsZWFzZSBXYWl0IDwvcD4='

<a id="l7-ddos-protection-clientside-action-js-challenge-js-script-delay"></a>&#x2022; [`js_script_delay`](#l7-ddos-protection-clientside-action-js-challenge-js-script-delay) - Optional Number<br>Javascript Delay. Delay introduced by Javascript, in milliseconds

#### L7 DDOS Protection DDOS Policy Custom

A [`ddos_policy_custom`](#l7-ddos-protection-ddos-policy-custom) block (within [`l7_ddos_protection`](#l7-ddos-protection)) supports the following:

<a id="l7-ddos-protection-ddos-policy-custom-name"></a>&#x2022; [`name`](#l7-ddos-protection-ddos-policy-custom-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="l7-ddos-protection-ddos-policy-custom-namespace"></a>&#x2022; [`namespace`](#l7-ddos-protection-ddos-policy-custom-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="l7-ddos-protection-ddos-policy-custom-tenant"></a>&#x2022; [`tenant`](#l7-ddos-protection-ddos-policy-custom-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### L7 DDOS Protection Mitigation Captcha Challenge

A [`mitigation_captcha_challenge`](#l7-ddos-protection-mitigation-captcha-challenge) block (within [`l7_ddos_protection`](#l7-ddos-protection)) supports the following:

<a id="l7-ddos-protection-mitigation-captcha-challenge-cookie-expiry"></a>&#x2022; [`cookie_expiry`](#l7-ddos-protection-mitigation-captcha-challenge-cookie-expiry) - Optional Number<br>Cookie Expiration Period. Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new challenge

<a id="l7-ddos-protection-mitigation-captcha-challenge-custom-page"></a>&#x2022; [`custom_page`](#l7-ddos-protection-mitigation-captcha-challenge-custom-page) - Optional String<br>Custom message for Captcha Challenge. Custom message is of type uri_ref. Currently supported URL schemes is string:///. For string:/// scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. 'Please Wait.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '`<p>` Please Wait `</p>`'. Base64 encoded string for this HTML is 'PHA+IFBsZWFzZSBXYWl0IDwvcD4='

#### L7 DDOS Protection Mitigation Js Challenge

A [`mitigation_js_challenge`](#l7-ddos-protection-mitigation-js-challenge) block (within [`l7_ddos_protection`](#l7-ddos-protection)) supports the following:

<a id="l7-ddos-protection-mitigation-js-challenge-cookie-expiry"></a>&#x2022; [`cookie_expiry`](#l7-ddos-protection-mitigation-js-challenge-cookie-expiry) - Optional Number<br>Cookie Expiration Period. Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new challenge

<a id="l7-ddos-protection-mitigation-js-challenge-custom-page"></a>&#x2022; [`custom_page`](#l7-ddos-protection-mitigation-js-challenge-custom-page) - Optional String<br>Custom Message for Javascript Challenge. Custom message is of type uri_ref. Currently supported URL schemes is string:///. For string:/// scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. 'Please Wait.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '`<p>` Please Wait `</p>`'. Base64 encoded string for this HTML is 'PHA+IFBsZWFzZSBXYWl0IDwvcD4='

<a id="l7-ddos-protection-mitigation-js-challenge-js-script-delay"></a>&#x2022; [`js_script_delay`](#l7-ddos-protection-mitigation-js-challenge-js-script-delay) - Optional Number<br>Javascript Delay. Delay introduced by Javascript, in milliseconds

#### Malware Protection Settings

A [`malware_protection_settings`](#malware-protection-settings) block supports the following:

<a id="malware-protection-settings-malware-protection-rules"></a>&#x2022; [`malware_protection_rules`](#malware-protection-settings-malware-protection-rules) - Optional Block<br>Malware Detection Rules. Configure the match criteria to trigger Malware Protection Scan<br>See [Malware Protection Rules](#malware-protection-settings-malware-protection-rules) below.

#### Malware Protection Settings Malware Protection Rules

A [`malware_protection_rules`](#malware-protection-settings-malware-protection-rules) block (within [`malware_protection_settings`](#malware-protection-settings)) supports the following:

<a id="malware-protection-settings-malware-protection-rules-action"></a>&#x2022; [`action`](#malware-protection-settings-malware-protection-rules-action) - Optional Block<br>Action<br>See [Action](#malware-protection-settings-malware-protection-rules-action) below.

<a id="malware-protection-settings-malware-protection-rules-domain"></a>&#x2022; [`domain`](#malware-protection-settings-malware-protection-rules-domain) - Optional Block<br>Domain to Match. Domain to be matched<br>See [Domain](#malware-protection-settings-malware-protection-rules-domain) below.

<a id="malware-protection-settings-malware-protection-rules-http-methods"></a>&#x2022; [`http_methods`](#malware-protection-settings-malware-protection-rules-http-methods) - Optional List  Defaults to `ANY`<br>Possible values are `ANY`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`, `COPY`<br>HTTP Methods. Methods to be matched

<a id="malware-protection-settings-malware-protection-rules-metadata"></a>&#x2022; [`metadata`](#malware-protection-settings-malware-protection-rules-metadata) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#malware-protection-settings-malware-protection-rules-metadata) below.

<a id="malware-protection-settings-malware-protection-rules-path"></a>&#x2022; [`path`](#malware-protection-settings-malware-protection-rules-path) - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#malware-protection-settings-malware-protection-rules-path) below.

#### Malware Protection Settings Malware Protection Rules Action

An [`action`](#malware-protection-settings-malware-protection-rules-action) block (within [`malware_protection_settings.malware_protection_rules`](#malware-protection-settings-malware-protection-rules)) supports the following:

<a id="malware-protection-settings-malware-protection-rules-action-block"></a>&#x2022; [`block`](#malware-protection-settings-malware-protection-rules-action-block) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="malware-protection-settings-malware-protection-rules-action-report"></a>&#x2022; [`report`](#malware-protection-settings-malware-protection-rules-action-report) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### Malware Protection Settings Malware Protection Rules Domain

A [`domain`](#malware-protection-settings-malware-protection-rules-domain) block (within [`malware_protection_settings.malware_protection_rules`](#malware-protection-settings-malware-protection-rules)) supports the following:

<a id="malware-protection-settings-malware-protection-rules-domain-any-domain"></a>&#x2022; [`any_domain`](#malware-protection-settings-malware-protection-rules-domain-any-domain) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="malware-protection-settings-malware-protection-rules-domain-domain"></a>&#x2022; [`domain`](#malware-protection-settings-malware-protection-rules-domain-domain) - Optional Block<br>Domains. Domains names<br>See [Domain](#malware-protection-settings-malware-protection-rules-domain-domain) below.

#### Malware Protection Settings Malware Protection Rules Domain Domain

A [`domain`](#malware-protection-settings-malware-protection-rules-domain-domain) block (within [`malware_protection_settings.malware_protection_rules.domain`](#malware-protection-settings-malware-protection-rules-domain)) supports the following:

<a id="malware-protection-settings-malware-protection-rules-domain-domain-exact-value"></a>&#x2022; [`exact_value`](#malware-protection-settings-malware-protection-rules-domain-domain-exact-value) - Optional String<br>Exact Value. Exact domain name

<a id="malware-protection-settings-malware-protection-rules-domain-domain-regex-value"></a>&#x2022; [`regex_value`](#malware-protection-settings-malware-protection-rules-domain-domain-regex-value) - Optional String<br>Regex Values of Domains. Regular Expression value for the domain name

<a id="malware-protection-settings-malware-protection-rules-domain-domain-suffix-value"></a>&#x2022; [`suffix_value`](#malware-protection-settings-malware-protection-rules-domain-domain-suffix-value) - Optional String<br>Suffix Value. Suffix of domain name e.g 'xyz.com' will match '*.xyz.com' and 'xyz.com'

#### Malware Protection Settings Malware Protection Rules Metadata

A [`metadata`](#malware-protection-settings-malware-protection-rules-metadata) block (within [`malware_protection_settings.malware_protection_rules`](#malware-protection-settings-malware-protection-rules)) supports the following:

<a id="malware-protection-settings-malware-protection-rules-metadata-description-spec"></a>&#x2022; [`description_spec`](#malware-protection-settings-malware-protection-rules-metadata-description-spec) - Optional String<br>Description. Human readable description

<a id="malware-protection-settings-malware-protection-rules-metadata-name"></a>&#x2022; [`name`](#malware-protection-settings-malware-protection-rules-metadata-name) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### Malware Protection Settings Malware Protection Rules Path

A [`path`](#malware-protection-settings-malware-protection-rules-path) block (within [`malware_protection_settings.malware_protection_rules`](#malware-protection-settings-malware-protection-rules)) supports the following:

<a id="malware-protection-settings-malware-protection-rules-path-path"></a>&#x2022; [`path`](#malware-protection-settings-malware-protection-rules-path-path) - Optional String<br>Exact. Exact path value to match

<a id="malware-protection-settings-malware-protection-rules-path-prefix"></a>&#x2022; [`prefix`](#malware-protection-settings-malware-protection-rules-path-prefix) - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

<a id="malware-protection-settings-malware-protection-rules-path-regex"></a>&#x2022; [`regex`](#malware-protection-settings-malware-protection-rules-path-regex) - Optional String<br>Regex. Regular expression of path match (e.g. the value .* will match on all paths)

#### More Option

A [`more_option`](#more-option) block supports the following:

<a id="more-option-buffer-policy"></a>&#x2022; [`buffer_policy`](#more-option-buffer-policy) - Optional Block<br>Buffer Configuration. Some upstream applications are not capable of handling streamed data. This config enables buffering the entire request before sending to upstream application. We can specify the maximum buffer size and buffer interval with this config. Buffering can be enabled and disabled at VirtualHost and Route levels Route level buffer configuration takes precedence<br>See [Buffer Policy](#more-option-buffer-policy) below.

<a id="more-option-compression-params"></a>&#x2022; [`compression_params`](#more-option-compression-params) - Optional Block<br>Compression Parameters. Enables loadbalancer to compress dispatched data from an upstream service upon client request. The content is compressed and then sent to the client with the appropriate headers if either response and request allow. Only GZIP compression is supported. By default compression will be skipped when: A request does NOT contain accept-encoding header. A request includes accept-encoding header, but it does not contain “gzip” or “*”. A request includes accept-encoding with “gzip” or “*” with the weight “q=0”. Note that the “gzip” will have a higher weight then “*”. For example, if accept-encoding is “gzip;q=0,*;q=1”, the filter will not compress. But if the header is set to “*;q=0,gzip;q=1”, the filter will compress. A request whose accept-encoding header includes “identity”. A response contains a content-encoding header. A response contains a cache-control header whose value includes “no-transform”. A response contains a transfer-encoding header whose value includes “gzip”. A response does not contain a content-type value that matches one of the selected mime-types, which default to application/javascript, application/JSON, application/xhtml+XML, image/svg+XML, text/CSS, text/HTML, text/plain, text/XML. Neither content-length nor transfer-encoding headers are present in the response. Response size is smaller than 30 bytes (only applicable when transfer-encoding is not chunked). When compression is applied: The content-length is removed from response headers. Response headers contain “transfer-encoding: chunked” and do not contain “content-encoding” header. The “vary: accept-encoding” header is inserted on every response. GZIP Compression Level: A value which is optimal balance between speed of compression and amount of compression is chosen<br>See [Compression Params](#more-option-compression-params) below.

<a id="more-option-custom-errors"></a>&#x2022; [`custom_errors`](#more-option-custom-errors) - Optional Block<br>Custom Error Responses. Map of integer error codes as keys and string values that can be used to provide custom HTTP pages for each error code. Key of the map can be either response code class or HTTP Error code. Response code classes for key is configured as follows 3 -- for 3xx response code class 4 -- for 4xx response code class 5 -- for 5xx response code class Value of the map is string which represents custom HTTP responses. Specific response code takes preference when both response code and response code class matches for a request

<a id="more-option-disable-default-error-pages"></a>&#x2022; [`disable_default_error_pages`](#more-option-disable-default-error-pages) - Optional Bool<br>Disable Default Error Pages. Disable the use of default F5XC error pages

<a id="more-option-disable-path-normalize"></a>&#x2022; [`disable_path_normalize`](#more-option-disable-path-normalize) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="more-option-enable-path-normalize"></a>&#x2022; [`enable_path_normalize`](#more-option-enable-path-normalize) - Optional Block<br>Empty. This can be used for messages where no values are needed

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

<a id="more-option-buffer-policy-max-request-bytes"></a>&#x2022; [`max_request_bytes`](#more-option-buffer-policy-max-request-bytes) - Optional Number<br>Max Request Bytes. The maximum request size that the filter will buffer before the connection manager will stop buffering and return a RequestEntityTooLarge (413) response

#### More Option Compression Params

A [`compression_params`](#more-option-compression-params) block (within [`more_option`](#more-option)) supports the following:

<a id="more-option-compression-params-content-length"></a>&#x2022; [`content_length`](#more-option-compression-params-content-length) - Optional Number  Defaults to `30`<br>Content Length. Minimum response length, in bytes, which will trigger compression. The

<a id="more-option-compression-params-content-type"></a>&#x2022; [`content_type`](#more-option-compression-params-content-type) - Optional List<br>Content Type. Set of strings that allows specifying which mime-types yield compression When this field is not defined, compression will be applied to the following mime-types: 'application/javascript' 'application/JSON', 'application/xhtml+XML' 'image/svg+XML' 'text/CSS' 'text/HTML' 'text/plain' 'text/XML'

<a id="more-option-compression-params-disable-on-etag-header"></a>&#x2022; [`disable_on_etag_header`](#more-option-compression-params-disable-on-etag-header) - Optional Bool<br>Disable On Etag Header. If true, disables compression when the response contains an etag header. When it is false, weak etags will be preserved and the ones that require strong validation will be removed

<a id="more-option-compression-params-remove-accept-encoding-header"></a>&#x2022; [`remove_accept_encoding_header`](#more-option-compression-params-remove-accept-encoding-header) - Optional Bool<br>Remove Accept-Encoding Header. If true, removes accept-encoding from the request headers before dispatching it to the upstream so that responses do not get compressed before reaching the filter

#### More Option Request Cookies To Add

A [`request_cookies_to_add`](#more-option-request-cookies-to-add) block (within [`more_option`](#more-option)) supports the following:

<a id="more-option-request-cookies-to-add-name"></a>&#x2022; [`name`](#more-option-request-cookies-to-add-name) - Optional String<br>Name. Name of the cookie in Cookie header

<a id="more-option-request-cookies-to-add-overwrite"></a>&#x2022; [`overwrite`](#more-option-request-cookies-to-add-overwrite) - Optional Bool  Defaults to `do`<br>Overwrite. Should the value be overwritten? If true, the value is overwritten to existing values. not overwrite

<a id="more-option-request-cookies-to-add-secret-value"></a>&#x2022; [`secret_value`](#more-option-request-cookies-to-add-secret-value) - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Secret Value](#more-option-request-cookies-to-add-secret-value) below.

<a id="more-option-request-cookies-to-add-value"></a>&#x2022; [`value`](#more-option-request-cookies-to-add-value) - Optional String<br>Value. Value of the Cookie header

#### More Option Request Cookies To Add Secret Value

A [`secret_value`](#more-option-request-cookies-to-add-secret-value) block (within [`more_option.request_cookies_to_add`](#more-option-request-cookies-to-add)) supports the following:

<a id="more-option-request-cookies-to-add-secret-value-blindfold-secret-info"></a>&#x2022; [`blindfold_secret_info`](#more-option-request-cookies-to-add-secret-value-blindfold-secret-info) - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#more-option-request-cookies-to-add-secret-value-blindfold-secret-info) below.

<a id="more-option-request-cookies-to-add-secret-value-clear-secret-info"></a>&#x2022; [`clear_secret_info`](#more-option-request-cookies-to-add-secret-value-clear-secret-info) - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#more-option-request-cookies-to-add-secret-value-clear-secret-info) below.

#### More Option Request Cookies To Add Secret Value Blindfold Secret Info

A [`blindfold_secret_info`](#more-option-request-cookies-to-add-secret-value-blindfold-secret-info) block (within [`more_option.request_cookies_to_add.secret_value`](#more-option-request-cookies-to-add-secret-value)) supports the following:

<a id="more-option-request-cookies-to-add-secret-value-blindfold-secret-info-decryption-provider"></a>&#x2022; [`decryption_provider`](#more-option-request-cookies-to-add-secret-value-blindfold-secret-info-decryption-provider) - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

<a id="more-option-request-cookies-to-add-secret-value-blindfold-secret-info-location"></a>&#x2022; [`location`](#more-option-request-cookies-to-add-secret-value-blindfold-secret-info-location) - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

<a id="more-option-request-cookies-to-add-secret-value-blindfold-secret-info-store-provider"></a>&#x2022; [`store_provider`](#more-option-request-cookies-to-add-secret-value-blindfold-secret-info-store-provider) - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

#### More Option Request Cookies To Add Secret Value Clear Secret Info

A [`clear_secret_info`](#more-option-request-cookies-to-add-secret-value-clear-secret-info) block (within [`more_option.request_cookies_to_add.secret_value`](#more-option-request-cookies-to-add-secret-value)) supports the following:

<a id="more-option-request-cookies-to-add-secret-value-clear-secret-info-provider-ref"></a>&#x2022; [`provider_ref`](#more-option-request-cookies-to-add-secret-value-clear-secret-info-provider-ref) - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

<a id="more-option-request-cookies-to-add-secret-value-clear-secret-info-url"></a>&#x2022; [`url`](#more-option-request-cookies-to-add-secret-value-clear-secret-info-url) - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

#### More Option Request Headers To Add

A [`request_headers_to_add`](#more-option-request-headers-to-add) block (within [`more_option`](#more-option)) supports the following:

<a id="more-option-request-headers-to-add-append"></a>&#x2022; [`append`](#more-option-request-headers-to-add-append) - Optional Bool  Defaults to `do`<br>Append. Should the value be appended? If true, the value is appended to existing values. not append

<a id="more-option-request-headers-to-add-name"></a>&#x2022; [`name`](#more-option-request-headers-to-add-name) - Optional String<br>Name. Name of the HTTP header

<a id="more-option-request-headers-to-add-secret-value"></a>&#x2022; [`secret_value`](#more-option-request-headers-to-add-secret-value) - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Secret Value](#more-option-request-headers-to-add-secret-value) below.

<a id="more-option-request-headers-to-add-value"></a>&#x2022; [`value`](#more-option-request-headers-to-add-value) - Optional String<br>Value. Value of the HTTP header

#### More Option Request Headers To Add Secret Value

A [`secret_value`](#more-option-request-headers-to-add-secret-value) block (within [`more_option.request_headers_to_add`](#more-option-request-headers-to-add)) supports the following:

<a id="more-option-request-headers-to-add-secret-value-blindfold-secret-info"></a>&#x2022; [`blindfold_secret_info`](#more-option-request-headers-to-add-secret-value-blindfold-secret-info) - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#more-option-request-headers-to-add-secret-value-blindfold-secret-info) below.

<a id="more-option-request-headers-to-add-secret-value-clear-secret-info"></a>&#x2022; [`clear_secret_info`](#more-option-request-headers-to-add-secret-value-clear-secret-info) - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#more-option-request-headers-to-add-secret-value-clear-secret-info) below.

#### More Option Request Headers To Add Secret Value Blindfold Secret Info

A [`blindfold_secret_info`](#more-option-request-headers-to-add-secret-value-blindfold-secret-info) block (within [`more_option.request_headers_to_add.secret_value`](#more-option-request-headers-to-add-secret-value)) supports the following:

<a id="more-option-request-headers-to-add-secret-value-blindfold-secret-info-decryption-provider"></a>&#x2022; [`decryption_provider`](#more-option-request-headers-to-add-secret-value-blindfold-secret-info-decryption-provider) - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

<a id="more-option-request-headers-to-add-secret-value-blindfold-secret-info-location"></a>&#x2022; [`location`](#more-option-request-headers-to-add-secret-value-blindfold-secret-info-location) - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

<a id="more-option-request-headers-to-add-secret-value-blindfold-secret-info-store-provider"></a>&#x2022; [`store_provider`](#more-option-request-headers-to-add-secret-value-blindfold-secret-info-store-provider) - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

#### More Option Request Headers To Add Secret Value Clear Secret Info

A [`clear_secret_info`](#more-option-request-headers-to-add-secret-value-clear-secret-info) block (within [`more_option.request_headers_to_add.secret_value`](#more-option-request-headers-to-add-secret-value)) supports the following:

<a id="more-option-request-headers-to-add-secret-value-clear-secret-info-provider-ref"></a>&#x2022; [`provider_ref`](#more-option-request-headers-to-add-secret-value-clear-secret-info-provider-ref) - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

<a id="more-option-request-headers-to-add-secret-value-clear-secret-info-url"></a>&#x2022; [`url`](#more-option-request-headers-to-add-secret-value-clear-secret-info-url) - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

#### More Option Response Cookies To Add

A [`response_cookies_to_add`](#more-option-response-cookies-to-add) block (within [`more_option`](#more-option)) supports the following:

<a id="more-option-response-cookies-to-add-add-domain"></a>&#x2022; [`add_domain`](#more-option-response-cookies-to-add-add-domain) - Optional String<br>Add Domain. Add domain attribute

<a id="more-option-response-cookies-to-add-add-expiry"></a>&#x2022; [`add_expiry`](#more-option-response-cookies-to-add-add-expiry) - Optional String<br>Add expiry. Add expiry attribute

<a id="more-option-response-cookies-to-add-add-httponly"></a>&#x2022; [`add_httponly`](#more-option-response-cookies-to-add-add-httponly) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="more-option-response-cookies-to-add-add-partitioned"></a>&#x2022; [`add_partitioned`](#more-option-response-cookies-to-add-add-partitioned) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="more-option-response-cookies-to-add-add-path"></a>&#x2022; [`add_path`](#more-option-response-cookies-to-add-add-path) - Optional String<br>Add path. Add path attribute

<a id="more-option-response-cookies-to-add-add-secure"></a>&#x2022; [`add_secure`](#more-option-response-cookies-to-add-add-secure) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="more-option-response-cookies-to-add-ignore-domain"></a>&#x2022; [`ignore_domain`](#more-option-response-cookies-to-add-ignore-domain) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="more-option-response-cookies-to-add-ignore-expiry"></a>&#x2022; [`ignore_expiry`](#more-option-response-cookies-to-add-ignore-expiry) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="more-option-response-cookies-to-add-ignore-httponly"></a>&#x2022; [`ignore_httponly`](#more-option-response-cookies-to-add-ignore-httponly) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="more-option-response-cookies-to-add-ignore-max-age"></a>&#x2022; [`ignore_max_age`](#more-option-response-cookies-to-add-ignore-max-age) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="more-option-response-cookies-to-add-ignore-partitioned"></a>&#x2022; [`ignore_partitioned`](#more-option-response-cookies-to-add-ignore-partitioned) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="more-option-response-cookies-to-add-ignore-path"></a>&#x2022; [`ignore_path`](#more-option-response-cookies-to-add-ignore-path) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="more-option-response-cookies-to-add-ignore-samesite"></a>&#x2022; [`ignore_samesite`](#more-option-response-cookies-to-add-ignore-samesite) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="more-option-response-cookies-to-add-ignore-secure"></a>&#x2022; [`ignore_secure`](#more-option-response-cookies-to-add-ignore-secure) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="more-option-response-cookies-to-add-ignore-value"></a>&#x2022; [`ignore_value`](#more-option-response-cookies-to-add-ignore-value) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="more-option-response-cookies-to-add-max-age-value"></a>&#x2022; [`max_age_value`](#more-option-response-cookies-to-add-max-age-value) - Optional Number<br>Add Max Age. Add max age attribute

<a id="more-option-response-cookies-to-add-name"></a>&#x2022; [`name`](#more-option-response-cookies-to-add-name) - Optional String<br>Name. Name of the cookie in Cookie header

<a id="more-option-response-cookies-to-add-overwrite"></a>&#x2022; [`overwrite`](#more-option-response-cookies-to-add-overwrite) - Optional Bool  Defaults to `do`<br>Overwrite. Should the value be overwritten? If true, the value is overwritten to existing values. not overwrite

<a id="more-option-response-cookies-to-add-samesite-lax"></a>&#x2022; [`samesite_lax`](#more-option-response-cookies-to-add-samesite-lax) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="more-option-response-cookies-to-add-samesite-none"></a>&#x2022; [`samesite_none`](#more-option-response-cookies-to-add-samesite-none) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="more-option-response-cookies-to-add-samesite-strict"></a>&#x2022; [`samesite_strict`](#more-option-response-cookies-to-add-samesite-strict) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="more-option-response-cookies-to-add-secret-value"></a>&#x2022; [`secret_value`](#more-option-response-cookies-to-add-secret-value) - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Secret Value](#more-option-response-cookies-to-add-secret-value) below.

<a id="more-option-response-cookies-to-add-value"></a>&#x2022; [`value`](#more-option-response-cookies-to-add-value) - Optional String<br>Value. Value of the Cookie header

#### More Option Response Cookies To Add Secret Value

A [`secret_value`](#more-option-response-cookies-to-add-secret-value) block (within [`more_option.response_cookies_to_add`](#more-option-response-cookies-to-add)) supports the following:

<a id="more-option-response-cookies-to-add-secret-value-blindfold-secret-info"></a>&#x2022; [`blindfold_secret_info`](#more-option-response-cookies-to-add-secret-value-blindfold-secret-info) - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#more-option-response-cookies-to-add-secret-value-blindfold-secret-info) below.

<a id="more-option-response-cookies-to-add-secret-value-clear-secret-info"></a>&#x2022; [`clear_secret_info`](#more-option-response-cookies-to-add-secret-value-clear-secret-info) - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#more-option-response-cookies-to-add-secret-value-clear-secret-info) below.

#### More Option Response Cookies To Add Secret Value Blindfold Secret Info

A [`blindfold_secret_info`](#more-option-response-cookies-to-add-secret-value-blindfold-secret-info) block (within [`more_option.response_cookies_to_add.secret_value`](#more-option-response-cookies-to-add-secret-value)) supports the following:

<a id="more-option-response-cookies-to-add-secret-value-blindfold-secret-info-decryption-provider"></a>&#x2022; [`decryption_provider`](#more-option-response-cookies-to-add-secret-value-blindfold-secret-info-decryption-provider) - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

<a id="more-option-response-cookies-to-add-secret-value-blindfold-secret-info-location"></a>&#x2022; [`location`](#more-option-response-cookies-to-add-secret-value-blindfold-secret-info-location) - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

<a id="more-option-response-cookies-to-add-secret-value-blindfold-secret-info-store-provider"></a>&#x2022; [`store_provider`](#more-option-response-cookies-to-add-secret-value-blindfold-secret-info-store-provider) - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

#### More Option Response Cookies To Add Secret Value Clear Secret Info

A [`clear_secret_info`](#more-option-response-cookies-to-add-secret-value-clear-secret-info) block (within [`more_option.response_cookies_to_add.secret_value`](#more-option-response-cookies-to-add-secret-value)) supports the following:

<a id="more-option-response-cookies-to-add-secret-value-clear-secret-info-provider-ref"></a>&#x2022; [`provider_ref`](#more-option-response-cookies-to-add-secret-value-clear-secret-info-provider-ref) - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

<a id="more-option-response-cookies-to-add-secret-value-clear-secret-info-url"></a>&#x2022; [`url`](#more-option-response-cookies-to-add-secret-value-clear-secret-info-url) - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

#### More Option Response Headers To Add

A [`response_headers_to_add`](#more-option-response-headers-to-add) block (within [`more_option`](#more-option)) supports the following:

<a id="more-option-response-headers-to-add-append"></a>&#x2022; [`append`](#more-option-response-headers-to-add-append) - Optional Bool  Defaults to `do`<br>Append. Should the value be appended? If true, the value is appended to existing values. not append

<a id="more-option-response-headers-to-add-name"></a>&#x2022; [`name`](#more-option-response-headers-to-add-name) - Optional String<br>Name. Name of the HTTP header

<a id="more-option-response-headers-to-add-secret-value"></a>&#x2022; [`secret_value`](#more-option-response-headers-to-add-secret-value) - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Secret Value](#more-option-response-headers-to-add-secret-value) below.

<a id="more-option-response-headers-to-add-value"></a>&#x2022; [`value`](#more-option-response-headers-to-add-value) - Optional String<br>Value. Value of the HTTP header

#### More Option Response Headers To Add Secret Value

A [`secret_value`](#more-option-response-headers-to-add-secret-value) block (within [`more_option.response_headers_to_add`](#more-option-response-headers-to-add)) supports the following:

<a id="more-option-response-headers-to-add-secret-value-blindfold-secret-info"></a>&#x2022; [`blindfold_secret_info`](#more-option-response-headers-to-add-secret-value-blindfold-secret-info) - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#more-option-response-headers-to-add-secret-value-blindfold-secret-info) below.

<a id="more-option-response-headers-to-add-secret-value-clear-secret-info"></a>&#x2022; [`clear_secret_info`](#more-option-response-headers-to-add-secret-value-clear-secret-info) - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#more-option-response-headers-to-add-secret-value-clear-secret-info) below.

#### More Option Response Headers To Add Secret Value Blindfold Secret Info

A [`blindfold_secret_info`](#more-option-response-headers-to-add-secret-value-blindfold-secret-info) block (within [`more_option.response_headers_to_add.secret_value`](#more-option-response-headers-to-add-secret-value)) supports the following:

<a id="more-option-response-headers-to-add-secret-value-blindfold-secret-info-decryption-provider"></a>&#x2022; [`decryption_provider`](#more-option-response-headers-to-add-secret-value-blindfold-secret-info-decryption-provider) - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

<a id="more-option-response-headers-to-add-secret-value-blindfold-secret-info-location"></a>&#x2022; [`location`](#more-option-response-headers-to-add-secret-value-blindfold-secret-info-location) - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

<a id="more-option-response-headers-to-add-secret-value-blindfold-secret-info-store-provider"></a>&#x2022; [`store_provider`](#more-option-response-headers-to-add-secret-value-blindfold-secret-info-store-provider) - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

#### More Option Response Headers To Add Secret Value Clear Secret Info

A [`clear_secret_info`](#more-option-response-headers-to-add-secret-value-clear-secret-info) block (within [`more_option.response_headers_to_add.secret_value`](#more-option-response-headers-to-add-secret-value)) supports the following:

<a id="more-option-response-headers-to-add-secret-value-clear-secret-info-provider-ref"></a>&#x2022; [`provider_ref`](#more-option-response-headers-to-add-secret-value-clear-secret-info-provider-ref) - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

<a id="more-option-response-headers-to-add-secret-value-clear-secret-info-url"></a>&#x2022; [`url`](#more-option-response-headers-to-add-secret-value-clear-secret-info-url) - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

#### Origin Server Subset Rule List

An [`origin_server_subset_rule_list`](#origin-server-subset-rule-list) block supports the following:

<a id="origin-server-subset-rule-list-origin-server-subset-rules"></a>&#x2022; [`origin_server_subset_rules`](#origin-server-subset-rule-list-origin-server-subset-rules) - Optional Block<br>Origin Server Subset Rules. Origin Server Subset Rules allow users to define match condition on Client (IP address, ASN, Country), IP Reputation, Regional Edge names, Request for subset selection of origin servers. Origin Server Subset is a sequential engine where rules are evaluated one after the other. It's important to define the correct order for Origin Server Subset to get the intended result, rules are evaluated from top to bottom in the list. When an Origin server subset rule is matched, then this selection rule takes effect and no more rules are evaluated<br>See [Origin Server Subset Rules](#origin-server-subset-rule-list-origin-server-subset-rules) below.

#### Origin Server Subset Rule List Origin Server Subset Rules

An [`origin_server_subset_rules`](#origin-server-subset-rule-list-origin-server-subset-rules) block (within [`origin_server_subset_rule_list`](#origin-server-subset-rule-list)) supports the following:

<a id="origin-server-subset-rule-list-origin-server-subset-rules-any-asn"></a>&#x2022; [`any_asn`](#origin-server-subset-rule-list-origin-server-subset-rules-any-asn) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="origin-server-subset-rule-list-origin-server-subset-rules-any-ip"></a>&#x2022; [`any_ip`](#origin-server-subset-rule-list-origin-server-subset-rules-any-ip) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="origin-server-subset-rule-list-origin-server-subset-rules-asn-list"></a>&#x2022; [`asn_list`](#origin-server-subset-rule-list-origin-server-subset-rules-asn-list) - Optional Block<br>ASN Match List. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer<br>See [Asn List](#origin-server-subset-rule-list-origin-server-subset-rules-asn-list) below.

<a id="origin-server-subset-rule-list-origin-server-subset-rules-asn-matcher"></a>&#x2022; [`asn_matcher`](#origin-server-subset-rule-list-origin-server-subset-rules-asn-matcher) - Optional Block<br>ASN Matcher. Match any AS number contained in the list of bgp_asn_sets<br>See [Asn Matcher](#origin-server-subset-rule-list-origin-server-subset-rules-asn-matcher) below.

<a id="origin-server-subset-rule-list-origin-server-subset-rules-client-selector"></a>&#x2022; [`client_selector`](#origin-server-subset-rule-list-origin-server-subset-rules-client-selector) - Optional Block<br>Label Selector. This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expresssions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects. Label selector is immutable. expressions is a list of strings of label selection expression. Each string has ',' separated values which are 'AND' and all strings are logically 'OR'. BNF for expression string `<selector-syntax>` ::= `<requirement>` | `<requirement>` ',' `<selector-syntax>` `<requirement>` ::= [!] KEY [ `<set-based-restriction>` | `<exact-match-restriction>` ] `<set-based-restriction>` ::= '' | `<inclusion-exclusion>` `<value-set>` `<inclusion-exclusion>` ::= `<inclusion>` | `<exclusion>` `<exclusion>` ::= 'notin' `<inclusion>` ::= 'in' `<value-set>` ::= '(' `<values>` ')' `<values>` ::= VALUE | VALUE ',' `<values>` `<exact-match-restriction>` ::= ['='|'=='|'!='] VALUE<br>See [Client Selector](#origin-server-subset-rule-list-origin-server-subset-rules-client-selector) below.

<a id="origin-server-subset-rule-list-origin-server-subset-rules-country-codes"></a>&#x2022; [`country_codes`](#origin-server-subset-rule-list-origin-server-subset-rules-country-codes) - Optional List  Defaults to `COUNTRY_NONE`<br>Possible values are `COUNTRY_NONE`, `COUNTRY_AD`, `COUNTRY_AE`, `COUNTRY_AF`, `COUNTRY_AG`, `COUNTRY_AI`, `COUNTRY_AL`, `COUNTRY_AM`, `COUNTRY_AN`, `COUNTRY_AO`, `COUNTRY_AQ`, `COUNTRY_AR`, `COUNTRY_AS`, `COUNTRY_AT`, `COUNTRY_AU`, `COUNTRY_AW`, `COUNTRY_AX`, `COUNTRY_AZ`, `COUNTRY_BA`, `COUNTRY_BB`, `COUNTRY_BD`, `COUNTRY_BE`, `COUNTRY_BF`, `COUNTRY_BG`, `COUNTRY_BH`, `COUNTRY_BI`, `COUNTRY_BJ`, `COUNTRY_BL`, `COUNTRY_BM`, `COUNTRY_BN`, `COUNTRY_BO`, `COUNTRY_BQ`, `COUNTRY_BR`, `COUNTRY_BS`, `COUNTRY_BT`, `COUNTRY_BV`, `COUNTRY_BW`, `COUNTRY_BY`, `COUNTRY_BZ`, `COUNTRY_CA`, `COUNTRY_CC`, `COUNTRY_CD`, `COUNTRY_CF`, `COUNTRY_CG`, `COUNTRY_CH`, `COUNTRY_CI`, `COUNTRY_CK`, `COUNTRY_CL`, `COUNTRY_CM`, `COUNTRY_CN`, `COUNTRY_CO`, `COUNTRY_CR`, `COUNTRY_CS`, `COUNTRY_CU`, `COUNTRY_CV`, `COUNTRY_CW`, `COUNTRY_CX`, `COUNTRY_CY`, `COUNTRY_CZ`, `COUNTRY_DE`, `COUNTRY_DJ`, `COUNTRY_DK`, `COUNTRY_DM`, `COUNTRY_DO`, `COUNTRY_DZ`, `COUNTRY_EC`, `COUNTRY_EE`, `COUNTRY_EG`, `COUNTRY_EH`, `COUNTRY_ER`, `COUNTRY_ES`, `COUNTRY_ET`, `COUNTRY_FI`, `COUNTRY_FJ`, `COUNTRY_FK`, `COUNTRY_FM`, `COUNTRY_FO`, `COUNTRY_FR`, `COUNTRY_GA`, `COUNTRY_GB`, `COUNTRY_GD`, `COUNTRY_GE`, `COUNTRY_GF`, `COUNTRY_GG`, `COUNTRY_GH`, `COUNTRY_GI`, `COUNTRY_GL`, `COUNTRY_GM`, `COUNTRY_GN`, `COUNTRY_GP`, `COUNTRY_GQ`, `COUNTRY_GR`, `COUNTRY_GS`, `COUNTRY_GT`, `COUNTRY_GU`, `COUNTRY_GW`, `COUNTRY_GY`, `COUNTRY_HK`, `COUNTRY_HM`, `COUNTRY_HN`, `COUNTRY_HR`, `COUNTRY_HT`, `COUNTRY_HU`, `COUNTRY_ID`, `COUNTRY_IE`, `COUNTRY_IL`, `COUNTRY_IM`, `COUNTRY_IN`, `COUNTRY_IO`, `COUNTRY_IQ`, `COUNTRY_IR`, `COUNTRY_IS`, `COUNTRY_IT`, `COUNTRY_JE`, `COUNTRY_JM`, `COUNTRY_JO`, `COUNTRY_JP`, `COUNTRY_KE`, `COUNTRY_KG`, `COUNTRY_KH`, `COUNTRY_KI`, `COUNTRY_KM`, `COUNTRY_KN`, `COUNTRY_KP`, `COUNTRY_KR`, `COUNTRY_KW`, `COUNTRY_KY`, `COUNTRY_KZ`, `COUNTRY_LA`, `COUNTRY_LB`, `COUNTRY_LC`, `COUNTRY_LI`, `COUNTRY_LK`, `COUNTRY_LR`, `COUNTRY_LS`, `COUNTRY_LT`, `COUNTRY_LU`, `COUNTRY_LV`, `COUNTRY_LY`, `COUNTRY_MA`, `COUNTRY_MC`, `COUNTRY_MD`, `COUNTRY_ME`, `COUNTRY_MF`, `COUNTRY_MG`, `COUNTRY_MH`, `COUNTRY_MK`, `COUNTRY_ML`, `COUNTRY_MM`, `COUNTRY_MN`, `COUNTRY_MO`, `COUNTRY_MP`, `COUNTRY_MQ`, `COUNTRY_MR`, `COUNTRY_MS`, `COUNTRY_MT`, `COUNTRY_MU`, `COUNTRY_MV`, `COUNTRY_MW`, `COUNTRY_MX`, `COUNTRY_MY`, `COUNTRY_MZ`, `COUNTRY_NA`, `COUNTRY_NC`, `COUNTRY_NE`, `COUNTRY_NF`, `COUNTRY_NG`, `COUNTRY_NI`, `COUNTRY_NL`, `COUNTRY_NO`, `COUNTRY_NP`, `COUNTRY_NR`, `COUNTRY_NU`, `COUNTRY_NZ`, `COUNTRY_OM`, `COUNTRY_PA`, `COUNTRY_PE`, `COUNTRY_PF`, `COUNTRY_PG`, `COUNTRY_PH`, `COUNTRY_PK`, `COUNTRY_PL`, `COUNTRY_PM`, `COUNTRY_PN`, `COUNTRY_PR`, `COUNTRY_PS`, `COUNTRY_PT`, `COUNTRY_PW`, `COUNTRY_PY`, `COUNTRY_QA`, `COUNTRY_RE`, `COUNTRY_RO`, `COUNTRY_RS`, `COUNTRY_RU`, `COUNTRY_RW`, `COUNTRY_SA`, `COUNTRY_SB`, `COUNTRY_SC`, `COUNTRY_SD`, `COUNTRY_SE`, `COUNTRY_SG`, `COUNTRY_SH`, `COUNTRY_SI`, `COUNTRY_SJ`, `COUNTRY_SK`, `COUNTRY_SL`, `COUNTRY_SM`, `COUNTRY_SN`, `COUNTRY_SO`, `COUNTRY_SR`, `COUNTRY_SS`, `COUNTRY_ST`, `COUNTRY_SV`, `COUNTRY_SX`, `COUNTRY_SY`, `COUNTRY_SZ`, `COUNTRY_TC`, `COUNTRY_TD`, `COUNTRY_TF`, `COUNTRY_TG`, `COUNTRY_TH`, `COUNTRY_TJ`, `COUNTRY_TK`, `COUNTRY_TL`, `COUNTRY_TM`, `COUNTRY_TN`, `COUNTRY_TO`, `COUNTRY_TR`, `COUNTRY_TT`, `COUNTRY_TV`, `COUNTRY_TW`, `COUNTRY_TZ`, `COUNTRY_UA`, `COUNTRY_UG`, `COUNTRY_UM`, `COUNTRY_US`, `COUNTRY_UY`, `COUNTRY_UZ`, `COUNTRY_VA`, `COUNTRY_VC`, `COUNTRY_VE`, `COUNTRY_VG`, `COUNTRY_VI`, `COUNTRY_VN`, `COUNTRY_VU`, `COUNTRY_WF`, `COUNTRY_WS`, `COUNTRY_XK`, `COUNTRY_XT`, `COUNTRY_YE`, `COUNTRY_YT`, `COUNTRY_ZA`, `COUNTRY_ZM`, `COUNTRY_ZW`<br>Country Codes List. List of Country Codes

<a id="origin-server-subset-rule-list-origin-server-subset-rules-ip-matcher"></a>&#x2022; [`ip_matcher`](#origin-server-subset-rule-list-origin-server-subset-rules-ip-matcher) - Optional Block<br>IP Prefix Matcher. Match any IP prefix contained in the list of ip_prefix_sets. The result of the match is inverted if invert_matcher is true<br>See [IP Matcher](#origin-server-subset-rule-list-origin-server-subset-rules-ip-matcher) below.

<a id="origin-server-subset-rule-list-origin-server-subset-rules-ip-prefix-list"></a>&#x2022; [`ip_prefix_list`](#origin-server-subset-rule-list-origin-server-subset-rules-ip-prefix-list) - Optional Block<br>IP Prefix Match List. List of IP Prefix strings to match against<br>See [IP Prefix List](#origin-server-subset-rule-list-origin-server-subset-rules-ip-prefix-list) below.

<a id="origin-server-subset-rule-list-origin-server-subset-rules-metadata"></a>&#x2022; [`metadata`](#origin-server-subset-rule-list-origin-server-subset-rules-metadata) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#origin-server-subset-rule-list-origin-server-subset-rules-metadata) below.

<a id="origin-server-subset-rule-list-origin-server-subset-rules-none"></a>&#x2022; [`none`](#origin-server-subset-rule-list-origin-server-subset-rules-none) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="origin-server-subset-rule-list-origin-server-subset-rules-origin-server-subsets-action"></a>&#x2022; [`origin_server_subsets_action`](#origin-server-subset-rule-list-origin-server-subset-rules-origin-server-subsets-action) - Optional Block<br>Action. Add labels to select one or more origin servers. Note: The pre-requisite settings to be configured in the origin pool are: 1. Add labels to origin servers 2. Enable subset load balancing in the Origin Server Subsets section and configure keys in origin server subsets classes

<a id="origin-server-subset-rule-list-origin-server-subset-rules-re-name-list"></a>&#x2022; [`re_name_list`](#origin-server-subset-rule-list-origin-server-subset-rules-re-name-list) - Optional List<br>RE Names. List of RE names for match

#### Origin Server Subset Rule List Origin Server Subset Rules Asn List

An [`asn_list`](#origin-server-subset-rule-list-origin-server-subset-rules-asn-list) block (within [`origin_server_subset_rule_list.origin_server_subset_rules`](#origin-server-subset-rule-list-origin-server-subset-rules)) supports the following:

<a id="origin-server-subset-rule-list-origin-server-subset-rules-asn-list-as-numbers"></a>&#x2022; [`as_numbers`](#origin-server-subset-rule-list-origin-server-subset-rules-asn-list-as-numbers) - Optional List<br>AS Numbers. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer

#### Origin Server Subset Rule List Origin Server Subset Rules Asn Matcher

An [`asn_matcher`](#origin-server-subset-rule-list-origin-server-subset-rules-asn-matcher) block (within [`origin_server_subset_rule_list.origin_server_subset_rules`](#origin-server-subset-rule-list-origin-server-subset-rules)) supports the following:

<a id="origin-server-subset-rule-list-origin-server-subset-rules-asn-matcher-asn-sets"></a>&#x2022; [`asn_sets`](#origin-server-subset-rule-list-origin-server-subset-rules-asn-matcher-asn-sets) - Optional Block<br>BGP ASN Sets. A list of references to bgp_asn_set objects<br>See [Asn Sets](#origin-server-subset-rule-list-origin-server-subset-rules-asn-matcher-asn-sets) below.

#### Origin Server Subset Rule List Origin Server Subset Rules Asn Matcher Asn Sets

An [`asn_sets`](#origin-server-subset-rule-list-origin-server-subset-rules-asn-matcher-asn-sets) block (within [`origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher`](#origin-server-subset-rule-list-origin-server-subset-rules-asn-matcher)) supports the following:

<a id="origin-server-subset-rule-list-origin-server-subset-rules-asn-matcher-asn-sets-kind"></a>&#x2022; [`kind`](#origin-server-subset-rule-list-origin-server-subset-rules-asn-matcher-asn-sets-kind) - Optional String<br>Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route')

<a id="origin-server-subset-rule-list-origin-server-subset-rules-asn-matcher-asn-sets-name"></a>&#x2022; [`name`](#origin-server-subset-rule-list-origin-server-subset-rules-asn-matcher-asn-sets-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="origin-server-subset-rule-list-origin-server-subset-rules-asn-matcher-asn-sets-namespace"></a>&#x2022; [`namespace`](#origin-server-subset-rule-list-origin-server-subset-rules-asn-matcher-asn-sets-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="origin-server-subset-rule-list-origin-server-subset-rules-asn-matcher-asn-sets-tenant"></a>&#x2022; [`tenant`](#origin-server-subset-rule-list-origin-server-subset-rules-asn-matcher-asn-sets-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="origin-server-subset-rule-list-origin-server-subset-rules-asn-matcher-asn-sets-uid"></a>&#x2022; [`uid`](#origin-server-subset-rule-list-origin-server-subset-rules-asn-matcher-asn-sets-uid) - Optional String<br>UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid

#### Origin Server Subset Rule List Origin Server Subset Rules Client Selector

A [`client_selector`](#origin-server-subset-rule-list-origin-server-subset-rules-client-selector) block (within [`origin_server_subset_rule_list.origin_server_subset_rules`](#origin-server-subset-rule-list-origin-server-subset-rules)) supports the following:

<a id="origin-server-subset-rule-list-origin-server-subset-rules-client-selector-expressions"></a>&#x2022; [`expressions`](#origin-server-subset-rule-list-origin-server-subset-rules-client-selector-expressions) - Optional List<br>Selector Expression. expressions contains the kubernetes style label expression for selections

#### Origin Server Subset Rule List Origin Server Subset Rules IP Matcher

An [`ip_matcher`](#origin-server-subset-rule-list-origin-server-subset-rules-ip-matcher) block (within [`origin_server_subset_rule_list.origin_server_subset_rules`](#origin-server-subset-rule-list-origin-server-subset-rules)) supports the following:

<a id="origin-server-subset-rule-list-origin-server-subset-rules-ip-matcher-invert-matcher"></a>&#x2022; [`invert_matcher`](#origin-server-subset-rule-list-origin-server-subset-rules-ip-matcher-invert-matcher) - Optional Bool<br>Invert IP Matcher. Invert the match result

<a id="origin-server-subset-rule-list-origin-server-subset-rules-ip-matcher-prefix-sets"></a>&#x2022; [`prefix_sets`](#origin-server-subset-rule-list-origin-server-subset-rules-ip-matcher-prefix-sets) - Optional Block<br>IP Prefix Sets. A list of references to ip_prefix_set objects<br>See [Prefix Sets](#origin-server-subset-rule-list-origin-server-subset-rules-ip-matcher-prefix-sets) below.

#### Origin Server Subset Rule List Origin Server Subset Rules IP Matcher Prefix Sets

A [`prefix_sets`](#origin-server-subset-rule-list-origin-server-subset-rules-ip-matcher-prefix-sets) block (within [`origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher`](#origin-server-subset-rule-list-origin-server-subset-rules-ip-matcher)) supports the following:

<a id="origin-server-subset-rule-list-origin-server-subset-rules-ip-matcher-prefix-sets-kind"></a>&#x2022; [`kind`](#origin-server-subset-rule-list-origin-server-subset-rules-ip-matcher-prefix-sets-kind) - Optional String<br>Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route')

<a id="origin-server-subset-rule-list-origin-server-subset-rules-ip-matcher-prefix-sets-name"></a>&#x2022; [`name`](#origin-server-subset-rule-list-origin-server-subset-rules-ip-matcher-prefix-sets-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="origin-server-subset-rule-list-origin-server-subset-rules-ip-matcher-prefix-sets-namespace"></a>&#x2022; [`namespace`](#origin-server-subset-rule-list-origin-server-subset-rules-ip-matcher-prefix-sets-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="origin-server-subset-rule-list-origin-server-subset-rules-ip-matcher-prefix-sets-tenant"></a>&#x2022; [`tenant`](#origin-server-subset-rule-list-origin-server-subset-rules-ip-matcher-prefix-sets-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="origin-server-subset-rule-list-origin-server-subset-rules-ip-matcher-prefix-sets-uid"></a>&#x2022; [`uid`](#origin-server-subset-rule-list-origin-server-subset-rules-ip-matcher-prefix-sets-uid) - Optional String<br>UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid

#### Origin Server Subset Rule List Origin Server Subset Rules IP Prefix List

An [`ip_prefix_list`](#origin-server-subset-rule-list-origin-server-subset-rules-ip-prefix-list) block (within [`origin_server_subset_rule_list.origin_server_subset_rules`](#origin-server-subset-rule-list-origin-server-subset-rules)) supports the following:

<a id="origin-server-subset-rule-list-origin-server-subset-rules-ip-prefix-list-invert-match"></a>&#x2022; [`invert_match`](#origin-server-subset-rule-list-origin-server-subset-rules-ip-prefix-list-invert-match) - Optional Bool<br>Invert Match Result. Invert the match result

<a id="origin-server-subset-rule-list-origin-server-subset-rules-ip-prefix-list-ip-prefixes"></a>&#x2022; [`ip_prefixes`](#origin-server-subset-rule-list-origin-server-subset-rules-ip-prefix-list-ip-prefixes) - Optional List<br>IPv4 Prefix List. List of IPv4 prefix strings

#### Origin Server Subset Rule List Origin Server Subset Rules Metadata

A [`metadata`](#origin-server-subset-rule-list-origin-server-subset-rules-metadata) block (within [`origin_server_subset_rule_list.origin_server_subset_rules`](#origin-server-subset-rule-list-origin-server-subset-rules)) supports the following:

<a id="origin-server-subset-rule-list-origin-server-subset-rules-metadata-description-spec"></a>&#x2022; [`description_spec`](#origin-server-subset-rule-list-origin-server-subset-rules-metadata-description-spec) - Optional String<br>Description. Human readable description

<a id="origin-server-subset-rule-list-origin-server-subset-rules-metadata-name"></a>&#x2022; [`name`](#origin-server-subset-rule-list-origin-server-subset-rules-metadata-name) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### Policy Based Challenge

A [`policy_based_challenge`](#policy-based-challenge) block supports the following:

<a id="policy-based-challenge-always-enable-captcha-challenge"></a>&#x2022; [`always_enable_captcha_challenge`](#policy-based-challenge-always-enable-captcha-challenge) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="policy-based-challenge-always-enable-js-challenge"></a>&#x2022; [`always_enable_js_challenge`](#policy-based-challenge-always-enable-js-challenge) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="policy-based-challenge-captcha-challenge-parameters"></a>&#x2022; [`captcha_challenge_parameters`](#policy-based-challenge-captcha-challenge-parameters) - Optional Block<br>Captcha Challenge Parameters. Enables loadbalancer to perform captcha challenge Captcha challenge will be based on Google Recaptcha. With this feature enabled, only clients that pass the captcha challenge will be allowed to complete the HTTP request. When loadbalancer is configured to do Captcha Challenge, it will redirect the browser to an HTML page on every new HTTP request. This HTML page will have captcha challenge embedded in it. Client will be allowed to make the request only if the captcha challenge is successful. Loadbalancer will tag response header with a cookie to avoid Captcha challenge for subsequent requests. CAPTCHA is mainly used as a security check to ensure only human users can pass through. Generally, computers or bots are not capable of solving a captcha. You can enable either Javascript challenge or Captcha challenge on a virtual host<br>See [Captcha Challenge Parameters](#policy-based-challenge-captcha-challenge-parameters) below.

<a id="policy-based-challenge-default-captcha-challenge-parameters"></a>&#x2022; [`default_captcha_challenge_parameters`](#policy-based-challenge-default-captcha-challenge-parameters) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="policy-based-challenge-default-js-challenge-parameters"></a>&#x2022; [`default_js_challenge_parameters`](#policy-based-challenge-default-js-challenge-parameters) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="policy-based-challenge-default-mitigation-settings"></a>&#x2022; [`default_mitigation_settings`](#policy-based-challenge-default-mitigation-settings) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="policy-based-challenge-default-temporary-blocking-parameters"></a>&#x2022; [`default_temporary_blocking_parameters`](#policy-based-challenge-default-temporary-blocking-parameters) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="policy-based-challenge-js-challenge-parameters"></a>&#x2022; [`js_challenge_parameters`](#policy-based-challenge-js-challenge-parameters) - Optional Block<br>Javascript Challenge Parameters. Enables loadbalancer to perform client browser compatibility test by redirecting to a page with Javascript. With this feature enabled, only clients that are capable of executing Javascript(mostly browsers) will be allowed to complete the HTTP request. When loadbalancer is configured to do Javascript Challenge, it will redirect the browser to an HTML page on every new HTTP request. This HTML page will have Javascript embedded in it. Loadbalancer chooses a set of random numbers for every new client and sends these numbers along with an encrypted answer with the request such that it embed these numbers as input in the Javascript. Javascript will run on the requestor browser and perform a complex Math operation. Script will submit the answer to loadbalancer. Loadbalancer will validate the answer by comparing the calculated answer with the decrypted answer (which was encrypted when it was sent back as reply) and allow the request to the upstream server only if the answer is correct. Loadbalancer will tag response header with a cookie to avoid Javascript challenge for subsequent requests. Javascript challenge serves following purposes * Validate that the request is coming via a browser that is capable for running Javascript * Force the browser to run a complex operation, f(X), that requires it to spend a large number of CPU cycles. This is to slow down a potential DOS attacker by making it difficult to launch a large request flood without having to spend even larger CPU cost at their end. You can enable either Javascript challenge or Captcha challenge on a virtual host<br>See [Js Challenge Parameters](#policy-based-challenge-js-challenge-parameters) below.

<a id="policy-based-challenge-malicious-user-mitigation"></a>&#x2022; [`malicious_user_mitigation`](#policy-based-challenge-malicious-user-mitigation) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Malicious User Mitigation](#policy-based-challenge-malicious-user-mitigation) below.

<a id="policy-based-challenge-no-challenge"></a>&#x2022; [`no_challenge`](#policy-based-challenge-no-challenge) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="policy-based-challenge-rule-list"></a>&#x2022; [`rule_list`](#policy-based-challenge-rule-list) - Optional Block<br>Challenge Rule List. List of challenge rules to be used in policy based challenge<br>See [Rule List](#policy-based-challenge-rule-list) below.

<a id="policy-based-challenge-temporary-user-blocking"></a>&#x2022; [`temporary_user_blocking`](#policy-based-challenge-temporary-user-blocking) - Optional Block<br>Temporary User Blocking. Specifies configuration for temporary user blocking resulting from user behavior analysis. When Malicious User Mitigation is enabled from service policy rules, users' accessing the application will be analyzed for malicious activity and the configured mitigation actions will be taken on identified malicious users. These mitigation actions include setting up temporary blocking on that user. This configuration specifies settings on how that blocking should be done by the loadbalancer<br>See [Temporary User Blocking](#policy-based-challenge-temporary-user-blocking) below.

#### Policy Based Challenge Captcha Challenge Parameters

A [`captcha_challenge_parameters`](#policy-based-challenge-captcha-challenge-parameters) block (within [`policy_based_challenge`](#policy-based-challenge)) supports the following:

<a id="policy-based-challenge-captcha-challenge-parameters-cookie-expiry"></a>&#x2022; [`cookie_expiry`](#policy-based-challenge-captcha-challenge-parameters-cookie-expiry) - Optional Number<br>Cookie Expiration Period. Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new challenge

<a id="policy-based-challenge-captcha-challenge-parameters-custom-page"></a>&#x2022; [`custom_page`](#policy-based-challenge-captcha-challenge-parameters-custom-page) - Optional String<br>Custom message for Captcha Challenge. Custom message is of type uri_ref. Currently supported URL schemes is string:///. For string:/// scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. 'Please Wait.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '`<p>` Please Wait `</p>`'. Base64 encoded string for this HTML is 'PHA+IFBsZWFzZSBXYWl0IDwvcD4='

#### Policy Based Challenge Js Challenge Parameters

A [`js_challenge_parameters`](#policy-based-challenge-js-challenge-parameters) block (within [`policy_based_challenge`](#policy-based-challenge)) supports the following:

<a id="policy-based-challenge-js-challenge-parameters-cookie-expiry"></a>&#x2022; [`cookie_expiry`](#policy-based-challenge-js-challenge-parameters-cookie-expiry) - Optional Number<br>Cookie Expiration Period. Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new challenge

<a id="policy-based-challenge-js-challenge-parameters-custom-page"></a>&#x2022; [`custom_page`](#policy-based-challenge-js-challenge-parameters-custom-page) - Optional String<br>Custom Message for Javascript Challenge. Custom message is of type uri_ref. Currently supported URL schemes is string:///. For string:/// scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. 'Please Wait.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '`<p>` Please Wait `</p>`'. Base64 encoded string for this HTML is 'PHA+IFBsZWFzZSBXYWl0IDwvcD4='

<a id="policy-based-challenge-js-challenge-parameters-js-script-delay"></a>&#x2022; [`js_script_delay`](#policy-based-challenge-js-challenge-parameters-js-script-delay) - Optional Number<br>Javascript Delay. Delay introduced by Javascript, in milliseconds

#### Policy Based Challenge Malicious User Mitigation

A [`malicious_user_mitigation`](#policy-based-challenge-malicious-user-mitigation) block (within [`policy_based_challenge`](#policy-based-challenge)) supports the following:

<a id="policy-based-challenge-malicious-user-mitigation-name"></a>&#x2022; [`name`](#policy-based-challenge-malicious-user-mitigation-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="policy-based-challenge-malicious-user-mitigation-namespace"></a>&#x2022; [`namespace`](#policy-based-challenge-malicious-user-mitigation-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="policy-based-challenge-malicious-user-mitigation-tenant"></a>&#x2022; [`tenant`](#policy-based-challenge-malicious-user-mitigation-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Policy Based Challenge Rule List

A [`rule_list`](#policy-based-challenge-rule-list) block (within [`policy_based_challenge`](#policy-based-challenge)) supports the following:

<a id="policy-based-challenge-rule-list-rules"></a>&#x2022; [`rules`](#policy-based-challenge-rule-list-rules) - Optional Block<br>Rules. Rules that specify the match conditions and challenge type to be launched. When a challenge type is selected to be always enabled, these rules can be used to disable challenge or launch a different challenge for requests that match the specified conditions<br>See [Rules](#policy-based-challenge-rule-list-rules) below.

#### Policy Based Challenge Rule List Rules

A [`rules`](#policy-based-challenge-rule-list-rules) block (within [`policy_based_challenge.rule_list`](#policy-based-challenge-rule-list)) supports the following:

<a id="policy-based-challenge-rule-list-rules-metadata"></a>&#x2022; [`metadata`](#policy-based-challenge-rule-list-rules-metadata) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#policy-based-challenge-rule-list-rules-metadata) below.

<a id="policy-based-challenge-rule-list-rules-spec"></a>&#x2022; [`spec`](#policy-based-challenge-rule-list-rules-spec) - Optional Block<br>Challenge Rule Specification. A Challenge Rule consists of an unordered list of predicates and an action. The predicates are evaluated against a set of input fields that are extracted from or derived from an L7 request API. A request API is considered to match the rule if all predicates in the rule evaluate to true for that request. Any predicates that are not specified in a rule are implicitly considered to be true. If a request API matches a challenge rule, the configured challenge is enforced<br>See [Spec](#policy-based-challenge-rule-list-rules-spec) below.

#### Policy Based Challenge Rule List Rules Metadata

A [`metadata`](#policy-based-challenge-rule-list-rules-metadata) block (within [`policy_based_challenge.rule_list.rules`](#policy-based-challenge-rule-list-rules)) supports the following:

<a id="policy-based-challenge-rule-list-rules-metadata-description-spec"></a>&#x2022; [`description_spec`](#policy-based-challenge-rule-list-rules-metadata-description-spec) - Optional String<br>Description. Human readable description

<a id="policy-based-challenge-rule-list-rules-metadata-name"></a>&#x2022; [`name`](#policy-based-challenge-rule-list-rules-metadata-name) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### Policy Based Challenge Rule List Rules Spec

A [`spec`](#policy-based-challenge-rule-list-rules-spec) block (within [`policy_based_challenge.rule_list.rules`](#policy-based-challenge-rule-list-rules)) supports the following:

<a id="policy-based-challenge-rule-list-rules-spec-any-asn"></a>&#x2022; [`any_asn`](#policy-based-challenge-rule-list-rules-spec-any-asn) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="policy-based-challenge-rule-list-rules-spec-any-client"></a>&#x2022; [`any_client`](#policy-based-challenge-rule-list-rules-spec-any-client) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="policy-based-challenge-rule-list-rules-spec-any-ip"></a>&#x2022; [`any_ip`](#policy-based-challenge-rule-list-rules-spec-any-ip) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="policy-based-challenge-rule-list-rules-spec-arg-matchers"></a>&#x2022; [`arg_matchers`](#policy-based-challenge-rule-list-rules-spec-arg-matchers) - Optional Block<br>A list of predicates for all POST args that need to be matched. The criteria for matching each arg are described in individual instances of ArgMatcherType. The actual arg values are extracted from the request API as a list of strings for each arg selector name. Note that all specified arg matcher predicates must evaluate to true<br>See [Arg Matchers](#policy-based-challenge-rule-list-rules-spec-arg-matchers) below.

<a id="policy-based-challenge-rule-list-rules-spec-asn-list"></a>&#x2022; [`asn_list`](#policy-based-challenge-rule-list-rules-spec-asn-list) - Optional Block<br>ASN Match List. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer<br>See [Asn List](#policy-based-challenge-rule-list-rules-spec-asn-list) below.

<a id="policy-based-challenge-rule-list-rules-spec-asn-matcher"></a>&#x2022; [`asn_matcher`](#policy-based-challenge-rule-list-rules-spec-asn-matcher) - Optional Block<br>ASN Matcher. Match any AS number contained in the list of bgp_asn_sets<br>See [Asn Matcher](#policy-based-challenge-rule-list-rules-spec-asn-matcher) below.

<a id="policy-based-challenge-rule-list-rules-spec-body-matcher"></a>&#x2022; [`body_matcher`](#policy-based-challenge-rule-list-rules-spec-body-matcher) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Body Matcher](#policy-based-challenge-rule-list-rules-spec-body-matcher) below.

<a id="policy-based-challenge-rule-list-rules-spec-client-selector"></a>&#x2022; [`client_selector`](#policy-based-challenge-rule-list-rules-spec-client-selector) - Optional Block<br>Label Selector. This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expresssions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects. Label selector is immutable. expressions is a list of strings of label selection expression. Each string has ',' separated values which are 'AND' and all strings are logically 'OR'. BNF for expression string `<selector-syntax>` ::= `<requirement>` | `<requirement>` ',' `<selector-syntax>` `<requirement>` ::= [!] KEY [ `<set-based-restriction>` | `<exact-match-restriction>` ] `<set-based-restriction>` ::= '' | `<inclusion-exclusion>` `<value-set>` `<inclusion-exclusion>` ::= `<inclusion>` | `<exclusion>` `<exclusion>` ::= 'notin' `<inclusion>` ::= 'in' `<value-set>` ::= '(' `<values>` ')' `<values>` ::= VALUE | VALUE ',' `<values>` `<exact-match-restriction>` ::= ['='|'=='|'!='] VALUE<br>See [Client Selector](#policy-based-challenge-rule-list-rules-spec-client-selector) below.

<a id="policy-based-challenge-rule-list-rules-spec-cookie-matchers"></a>&#x2022; [`cookie_matchers`](#policy-based-challenge-rule-list-rules-spec-cookie-matchers) - Optional Block<br>A list of predicates for all cookies that need to be matched. The criteria for matching each cookie is described in individual instances of CookieMatcherType. The actual cookie values are extracted from the request API as a list of strings for each cookie name. Note that all specified cookie matcher predicates must evaluate to true<br>See [Cookie Matchers](#policy-based-challenge-rule-list-rules-spec-cookie-matchers) below.

<a id="policy-based-challenge-rule-list-rules-spec-disable-challenge"></a>&#x2022; [`disable_challenge`](#policy-based-challenge-rule-list-rules-spec-disable-challenge) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="policy-based-challenge-rule-list-rules-spec-domain-matcher"></a>&#x2022; [`domain_matcher`](#policy-based-challenge-rule-list-rules-spec-domain-matcher) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Domain Matcher](#policy-based-challenge-rule-list-rules-spec-domain-matcher) below.

<a id="policy-based-challenge-rule-list-rules-spec-enable-captcha-challenge"></a>&#x2022; [`enable_captcha_challenge`](#policy-based-challenge-rule-list-rules-spec-enable-captcha-challenge) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="policy-based-challenge-rule-list-rules-spec-enable-javascript-challenge"></a>&#x2022; [`enable_javascript_challenge`](#policy-based-challenge-rule-list-rules-spec-enable-javascript-challenge) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="policy-based-challenge-rule-list-rules-spec-expiration-timestamp"></a>&#x2022; [`expiration_timestamp`](#policy-based-challenge-rule-list-rules-spec-expiration-timestamp) - Optional String<br>The expiration_timestamp is the RFC 3339 format timestamp at which the containing rule is considered to be logically expired. The rule continues to exist in the configuration but is not applied anymore

<a id="policy-based-challenge-rule-list-rules-spec-headers"></a>&#x2022; [`headers`](#policy-based-challenge-rule-list-rules-spec-headers) - Optional Block<br>A list of predicates for various HTTP headers that need to match. The criteria for matching each HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values are extracted from the request API as a list of strings for each HTTP header type. Note that all specified header predicates must evaluate to true<br>See [Headers](#policy-based-challenge-rule-list-rules-spec-headers) below.

<a id="policy-based-challenge-rule-list-rules-spec-http-method"></a>&#x2022; [`http_method`](#policy-based-challenge-rule-list-rules-spec-http-method) - Optional Block<br>HTTP Method Matcher. A HTTP method matcher specifies a list of methods to match an input HTTP method. The match is considered successful if the input method is a member of the list. The result of the match based on the method list is inverted if invert_matcher is true<br>See [HTTP Method](#policy-based-challenge-rule-list-rules-spec-http-method) below.

<a id="policy-based-challenge-rule-list-rules-spec-ip-matcher"></a>&#x2022; [`ip_matcher`](#policy-based-challenge-rule-list-rules-spec-ip-matcher) - Optional Block<br>IP Prefix Matcher. Match any IP prefix contained in the list of ip_prefix_sets. The result of the match is inverted if invert_matcher is true<br>See [IP Matcher](#policy-based-challenge-rule-list-rules-spec-ip-matcher) below.

<a id="policy-based-challenge-rule-list-rules-spec-ip-prefix-list"></a>&#x2022; [`ip_prefix_list`](#policy-based-challenge-rule-list-rules-spec-ip-prefix-list) - Optional Block<br>IP Prefix Match List. List of IP Prefix strings to match against<br>See [IP Prefix List](#policy-based-challenge-rule-list-rules-spec-ip-prefix-list) below.

<a id="policy-based-challenge-rule-list-rules-spec-path"></a>&#x2022; [`path`](#policy-based-challenge-rule-list-rules-spec-path) - Optional Block<br>Path Matcher. A path matcher specifies multiple criteria for matching an HTTP path string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of path prefixes, a list of exact path values and a list of regular expressions<br>See [Path](#policy-based-challenge-rule-list-rules-spec-path) below.

<a id="policy-based-challenge-rule-list-rules-spec-query-params"></a>&#x2022; [`query_params`](#policy-based-challenge-rule-list-rules-spec-query-params) - Optional Block<br>A list of predicates for all query parameters that need to be matched. The criteria for matching each query parameter are described in individual instances of QueryParameterMatcherType. The actual query parameter values are extracted from the request API as a list of strings for each query parameter name. Note that all specified query parameter predicates must evaluate to true<br>See [Query Params](#policy-based-challenge-rule-list-rules-spec-query-params) below.

<a id="policy-based-challenge-rule-list-rules-spec-tls-fingerprint-matcher"></a>&#x2022; [`tls_fingerprint_matcher`](#policy-based-challenge-rule-list-rules-spec-tls-fingerprint-matcher) - Optional Block<br>TLS Fingerprint Matcher. A TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of supported positve match criteria includes a list of known classes of TLS fingerprints and a list of exact values. The match is considered successful if either of these positive criteria are satisfied and the input fingerprint is not one of the excluded values<br>See [TLS Fingerprint Matcher](#policy-based-challenge-rule-list-rules-spec-tls-fingerprint-matcher) below.

#### Policy Based Challenge Rule List Rules Spec Arg Matchers

An [`arg_matchers`](#policy-based-challenge-rule-list-rules-spec-arg-matchers) block (within [`policy_based_challenge.rule_list.rules.spec`](#policy-based-challenge-rule-list-rules-spec)) supports the following:

<a id="policy-based-challenge-rule-list-rules-spec-arg-matchers-check-not-present"></a>&#x2022; [`check_not_present`](#policy-based-challenge-rule-list-rules-spec-arg-matchers-check-not-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="policy-based-challenge-rule-list-rules-spec-arg-matchers-check-present"></a>&#x2022; [`check_present`](#policy-based-challenge-rule-list-rules-spec-arg-matchers-check-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="policy-based-challenge-rule-list-rules-spec-arg-matchers-invert-matcher"></a>&#x2022; [`invert_matcher`](#policy-based-challenge-rule-list-rules-spec-arg-matchers-invert-matcher) - Optional Bool<br>Invert Matcher. Invert Match of the expression defined

<a id="policy-based-challenge-rule-list-rules-spec-arg-matchers-item"></a>&#x2022; [`item`](#policy-based-challenge-rule-list-rules-spec-arg-matchers-item) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#policy-based-challenge-rule-list-rules-spec-arg-matchers-item) below.

<a id="policy-based-challenge-rule-list-rules-spec-arg-matchers-name"></a>&#x2022; [`name`](#policy-based-challenge-rule-list-rules-spec-arg-matchers-name) - Optional String<br>Argument Name. x-example: 'phones[_]' x-example: 'cars.make.toyota.models[1]' x-example: 'cars.make.honda.models[_]' x-example: 'cars.make[_].models[_]' A case-sensitive JSON path in the HTTP request body

#### Policy Based Challenge Rule List Rules Spec Arg Matchers Item

An [`item`](#policy-based-challenge-rule-list-rules-spec-arg-matchers-item) block (within [`policy_based_challenge.rule_list.rules.spec.arg_matchers`](#policy-based-challenge-rule-list-rules-spec-arg-matchers)) supports the following:

<a id="policy-based-challenge-rule-list-rules-spec-arg-matchers-item-exact-values"></a>&#x2022; [`exact_values`](#policy-based-challenge-rule-list-rules-spec-arg-matchers-item-exact-values) - Optional List<br>Exact Values. A list of exact values to match the input against

<a id="policy-based-challenge-rule-list-rules-spec-arg-matchers-item-regex-values"></a>&#x2022; [`regex_values`](#policy-based-challenge-rule-list-rules-spec-arg-matchers-item-regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="policy-based-challenge-rule-list-rules-spec-arg-matchers-item-transformers"></a>&#x2022; [`transformers`](#policy-based-challenge-rule-list-rules-spec-arg-matchers-item-transformers) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### Policy Based Challenge Rule List Rules Spec Asn List

An [`asn_list`](#policy-based-challenge-rule-list-rules-spec-asn-list) block (within [`policy_based_challenge.rule_list.rules.spec`](#policy-based-challenge-rule-list-rules-spec)) supports the following:

<a id="policy-based-challenge-rule-list-rules-spec-asn-list-as-numbers"></a>&#x2022; [`as_numbers`](#policy-based-challenge-rule-list-rules-spec-asn-list-as-numbers) - Optional List<br>AS Numbers. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer

#### Policy Based Challenge Rule List Rules Spec Asn Matcher

An [`asn_matcher`](#policy-based-challenge-rule-list-rules-spec-asn-matcher) block (within [`policy_based_challenge.rule_list.rules.spec`](#policy-based-challenge-rule-list-rules-spec)) supports the following:

<a id="policy-based-challenge-rule-list-rules-spec-asn-matcher-asn-sets"></a>&#x2022; [`asn_sets`](#policy-based-challenge-rule-list-rules-spec-asn-matcher-asn-sets) - Optional Block<br>BGP ASN Sets. A list of references to bgp_asn_set objects<br>See [Asn Sets](#policy-based-challenge-rule-list-rules-spec-asn-matcher-asn-sets) below.

#### Policy Based Challenge Rule List Rules Spec Asn Matcher Asn Sets

An [`asn_sets`](#policy-based-challenge-rule-list-rules-spec-asn-matcher-asn-sets) block (within [`policy_based_challenge.rule_list.rules.spec.asn_matcher`](#policy-based-challenge-rule-list-rules-spec-asn-matcher)) supports the following:

<a id="policy-based-challenge-rule-list-rules-spec-asn-matcher-asn-sets-kind"></a>&#x2022; [`kind`](#policy-based-challenge-rule-list-rules-spec-asn-matcher-asn-sets-kind) - Optional String<br>Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route')

<a id="policy-based-challenge-rule-list-rules-spec-asn-matcher-asn-sets-name"></a>&#x2022; [`name`](#policy-based-challenge-rule-list-rules-spec-asn-matcher-asn-sets-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="policy-based-challenge-rule-list-rules-spec-asn-matcher-asn-sets-namespace"></a>&#x2022; [`namespace`](#policy-based-challenge-rule-list-rules-spec-asn-matcher-asn-sets-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="policy-based-challenge-rule-list-rules-spec-asn-matcher-asn-sets-tenant"></a>&#x2022; [`tenant`](#policy-based-challenge-rule-list-rules-spec-asn-matcher-asn-sets-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="policy-based-challenge-rule-list-rules-spec-asn-matcher-asn-sets-uid"></a>&#x2022; [`uid`](#policy-based-challenge-rule-list-rules-spec-asn-matcher-asn-sets-uid) - Optional String<br>UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid

#### Policy Based Challenge Rule List Rules Spec Body Matcher

A [`body_matcher`](#policy-based-challenge-rule-list-rules-spec-body-matcher) block (within [`policy_based_challenge.rule_list.rules.spec`](#policy-based-challenge-rule-list-rules-spec)) supports the following:

<a id="policy-based-challenge-rule-list-rules-spec-body-matcher-exact-values"></a>&#x2022; [`exact_values`](#policy-based-challenge-rule-list-rules-spec-body-matcher-exact-values) - Optional List<br>Exact Values. A list of exact values to match the input against

<a id="policy-based-challenge-rule-list-rules-spec-body-matcher-regex-values"></a>&#x2022; [`regex_values`](#policy-based-challenge-rule-list-rules-spec-body-matcher-regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="policy-based-challenge-rule-list-rules-spec-body-matcher-transformers"></a>&#x2022; [`transformers`](#policy-based-challenge-rule-list-rules-spec-body-matcher-transformers) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### Policy Based Challenge Rule List Rules Spec Client Selector

A [`client_selector`](#policy-based-challenge-rule-list-rules-spec-client-selector) block (within [`policy_based_challenge.rule_list.rules.spec`](#policy-based-challenge-rule-list-rules-spec)) supports the following:

<a id="policy-based-challenge-rule-list-rules-spec-client-selector-expressions"></a>&#x2022; [`expressions`](#policy-based-challenge-rule-list-rules-spec-client-selector-expressions) - Optional List<br>Selector Expression. expressions contains the kubernetes style label expression for selections

#### Policy Based Challenge Rule List Rules Spec Cookie Matchers

A [`cookie_matchers`](#policy-based-challenge-rule-list-rules-spec-cookie-matchers) block (within [`policy_based_challenge.rule_list.rules.spec`](#policy-based-challenge-rule-list-rules-spec)) supports the following:

<a id="policy-based-challenge-rule-list-rules-spec-cookie-matchers-check-not-present"></a>&#x2022; [`check_not_present`](#policy-based-challenge-rule-list-rules-spec-cookie-matchers-check-not-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="policy-based-challenge-rule-list-rules-spec-cookie-matchers-check-present"></a>&#x2022; [`check_present`](#policy-based-challenge-rule-list-rules-spec-cookie-matchers-check-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="policy-based-challenge-rule-list-rules-spec-cookie-matchers-invert-matcher"></a>&#x2022; [`invert_matcher`](#policy-based-challenge-rule-list-rules-spec-cookie-matchers-invert-matcher) - Optional Bool<br>Invert Matcher. Invert Match of the expression defined

<a id="policy-based-challenge-rule-list-rules-spec-cookie-matchers-item"></a>&#x2022; [`item`](#policy-based-challenge-rule-list-rules-spec-cookie-matchers-item) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#policy-based-challenge-rule-list-rules-spec-cookie-matchers-item) below.

<a id="policy-based-challenge-rule-list-rules-spec-cookie-matchers-name"></a>&#x2022; [`name`](#policy-based-challenge-rule-list-rules-spec-cookie-matchers-name) - Optional String<br>Cookie Name. A case-sensitive cookie name

#### Policy Based Challenge Rule List Rules Spec Cookie Matchers Item

An [`item`](#policy-based-challenge-rule-list-rules-spec-cookie-matchers-item) block (within [`policy_based_challenge.rule_list.rules.spec.cookie_matchers`](#policy-based-challenge-rule-list-rules-spec-cookie-matchers)) supports the following:

<a id="policy-based-challenge-rule-list-rules-spec-cookie-matchers-item-exact-values"></a>&#x2022; [`exact_values`](#policy-based-challenge-rule-list-rules-spec-cookie-matchers-item-exact-values) - Optional List<br>Exact Values. A list of exact values to match the input against

<a id="policy-based-challenge-rule-list-rules-spec-cookie-matchers-item-regex-values"></a>&#x2022; [`regex_values`](#policy-based-challenge-rule-list-rules-spec-cookie-matchers-item-regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="policy-based-challenge-rule-list-rules-spec-cookie-matchers-item-transformers"></a>&#x2022; [`transformers`](#policy-based-challenge-rule-list-rules-spec-cookie-matchers-item-transformers) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### Policy Based Challenge Rule List Rules Spec Domain Matcher

A [`domain_matcher`](#policy-based-challenge-rule-list-rules-spec-domain-matcher) block (within [`policy_based_challenge.rule_list.rules.spec`](#policy-based-challenge-rule-list-rules-spec)) supports the following:

<a id="policy-based-challenge-rule-list-rules-spec-domain-matcher-exact-values"></a>&#x2022; [`exact_values`](#policy-based-challenge-rule-list-rules-spec-domain-matcher-exact-values) - Optional List<br>Exact Values. A list of exact values to match the input against

<a id="policy-based-challenge-rule-list-rules-spec-domain-matcher-regex-values"></a>&#x2022; [`regex_values`](#policy-based-challenge-rule-list-rules-spec-domain-matcher-regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

#### Policy Based Challenge Rule List Rules Spec Headers

A [`headers`](#policy-based-challenge-rule-list-rules-spec-headers) block (within [`policy_based_challenge.rule_list.rules.spec`](#policy-based-challenge-rule-list-rules-spec)) supports the following:

<a id="policy-based-challenge-rule-list-rules-spec-headers-check-not-present"></a>&#x2022; [`check_not_present`](#policy-based-challenge-rule-list-rules-spec-headers-check-not-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="policy-based-challenge-rule-list-rules-spec-headers-check-present"></a>&#x2022; [`check_present`](#policy-based-challenge-rule-list-rules-spec-headers-check-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="policy-based-challenge-rule-list-rules-spec-headers-invert-matcher"></a>&#x2022; [`invert_matcher`](#policy-based-challenge-rule-list-rules-spec-headers-invert-matcher) - Optional Bool<br>Invert Header Matcher. Invert the match result

<a id="policy-based-challenge-rule-list-rules-spec-headers-item"></a>&#x2022; [`item`](#policy-based-challenge-rule-list-rules-spec-headers-item) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#policy-based-challenge-rule-list-rules-spec-headers-item) below.

<a id="policy-based-challenge-rule-list-rules-spec-headers-name"></a>&#x2022; [`name`](#policy-based-challenge-rule-list-rules-spec-headers-name) - Optional String<br>Header Name. A case-insensitive HTTP header name

#### Policy Based Challenge Rule List Rules Spec Headers Item

An [`item`](#policy-based-challenge-rule-list-rules-spec-headers-item) block (within [`policy_based_challenge.rule_list.rules.spec.headers`](#policy-based-challenge-rule-list-rules-spec-headers)) supports the following:

<a id="policy-based-challenge-rule-list-rules-spec-headers-item-exact-values"></a>&#x2022; [`exact_values`](#policy-based-challenge-rule-list-rules-spec-headers-item-exact-values) - Optional List<br>Exact Values. A list of exact values to match the input against

<a id="policy-based-challenge-rule-list-rules-spec-headers-item-regex-values"></a>&#x2022; [`regex_values`](#policy-based-challenge-rule-list-rules-spec-headers-item-regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="policy-based-challenge-rule-list-rules-spec-headers-item-transformers"></a>&#x2022; [`transformers`](#policy-based-challenge-rule-list-rules-spec-headers-item-transformers) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### Policy Based Challenge Rule List Rules Spec HTTP Method

A [`http_method`](#policy-based-challenge-rule-list-rules-spec-http-method) block (within [`policy_based_challenge.rule_list.rules.spec`](#policy-based-challenge-rule-list-rules-spec)) supports the following:

<a id="policy-based-challenge-rule-list-rules-spec-http-method-invert-matcher"></a>&#x2022; [`invert_matcher`](#policy-based-challenge-rule-list-rules-spec-http-method-invert-matcher) - Optional Bool<br>Invert Method Matcher. Invert the match result

<a id="policy-based-challenge-rule-list-rules-spec-http-method-methods"></a>&#x2022; [`methods`](#policy-based-challenge-rule-list-rules-spec-http-method-methods) - Optional List  Defaults to `ANY`<br>Possible values are `ANY`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`, `COPY`<br>Method List. List of methods values to match against

#### Policy Based Challenge Rule List Rules Spec IP Matcher

An [`ip_matcher`](#policy-based-challenge-rule-list-rules-spec-ip-matcher) block (within [`policy_based_challenge.rule_list.rules.spec`](#policy-based-challenge-rule-list-rules-spec)) supports the following:

<a id="policy-based-challenge-rule-list-rules-spec-ip-matcher-invert-matcher"></a>&#x2022; [`invert_matcher`](#policy-based-challenge-rule-list-rules-spec-ip-matcher-invert-matcher) - Optional Bool<br>Invert IP Matcher. Invert the match result

<a id="policy-based-challenge-rule-list-rules-spec-ip-matcher-prefix-sets"></a>&#x2022; [`prefix_sets`](#policy-based-challenge-rule-list-rules-spec-ip-matcher-prefix-sets) - Optional Block<br>IP Prefix Sets. A list of references to ip_prefix_set objects<br>See [Prefix Sets](#policy-based-challenge-rule-list-rules-spec-ip-matcher-prefix-sets) below.

#### Policy Based Challenge Rule List Rules Spec IP Matcher Prefix Sets

A [`prefix_sets`](#policy-based-challenge-rule-list-rules-spec-ip-matcher-prefix-sets) block (within [`policy_based_challenge.rule_list.rules.spec.ip_matcher`](#policy-based-challenge-rule-list-rules-spec-ip-matcher)) supports the following:

<a id="policy-based-challenge-rule-list-rules-spec-ip-matcher-prefix-sets-kind"></a>&#x2022; [`kind`](#policy-based-challenge-rule-list-rules-spec-ip-matcher-prefix-sets-kind) - Optional String<br>Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route')

<a id="policy-based-challenge-rule-list-rules-spec-ip-matcher-prefix-sets-name"></a>&#x2022; [`name`](#policy-based-challenge-rule-list-rules-spec-ip-matcher-prefix-sets-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="policy-based-challenge-rule-list-rules-spec-ip-matcher-prefix-sets-namespace"></a>&#x2022; [`namespace`](#policy-based-challenge-rule-list-rules-spec-ip-matcher-prefix-sets-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="policy-based-challenge-rule-list-rules-spec-ip-matcher-prefix-sets-tenant"></a>&#x2022; [`tenant`](#policy-based-challenge-rule-list-rules-spec-ip-matcher-prefix-sets-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="policy-based-challenge-rule-list-rules-spec-ip-matcher-prefix-sets-uid"></a>&#x2022; [`uid`](#policy-based-challenge-rule-list-rules-spec-ip-matcher-prefix-sets-uid) - Optional String<br>UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid

#### Policy Based Challenge Rule List Rules Spec IP Prefix List

An [`ip_prefix_list`](#policy-based-challenge-rule-list-rules-spec-ip-prefix-list) block (within [`policy_based_challenge.rule_list.rules.spec`](#policy-based-challenge-rule-list-rules-spec)) supports the following:

<a id="policy-based-challenge-rule-list-rules-spec-ip-prefix-list-invert-match"></a>&#x2022; [`invert_match`](#policy-based-challenge-rule-list-rules-spec-ip-prefix-list-invert-match) - Optional Bool<br>Invert Match Result. Invert the match result

<a id="policy-based-challenge-rule-list-rules-spec-ip-prefix-list-ip-prefixes"></a>&#x2022; [`ip_prefixes`](#policy-based-challenge-rule-list-rules-spec-ip-prefix-list-ip-prefixes) - Optional List<br>IPv4 Prefix List. List of IPv4 prefix strings

#### Policy Based Challenge Rule List Rules Spec Path

A [`path`](#policy-based-challenge-rule-list-rules-spec-path) block (within [`policy_based_challenge.rule_list.rules.spec`](#policy-based-challenge-rule-list-rules-spec)) supports the following:

<a id="policy-based-challenge-rule-list-rules-spec-path-exact-values"></a>&#x2022; [`exact_values`](#policy-based-challenge-rule-list-rules-spec-path-exact-values) - Optional List<br>Exact Values. A list of exact path values to match the input HTTP path against

<a id="policy-based-challenge-rule-list-rules-spec-path-invert-matcher"></a>&#x2022; [`invert_matcher`](#policy-based-challenge-rule-list-rules-spec-path-invert-matcher) - Optional Bool<br>Invert Path Matcher. Invert the match result

<a id="policy-based-challenge-rule-list-rules-spec-path-prefix-values"></a>&#x2022; [`prefix_values`](#policy-based-challenge-rule-list-rules-spec-path-prefix-values) - Optional List<br>Prefix Values. A list of path prefix values to match the input HTTP path against

<a id="policy-based-challenge-rule-list-rules-spec-path-regex-values"></a>&#x2022; [`regex_values`](#policy-based-challenge-rule-list-rules-spec-path-regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input HTTP path against

<a id="policy-based-challenge-rule-list-rules-spec-path-suffix-values"></a>&#x2022; [`suffix_values`](#policy-based-challenge-rule-list-rules-spec-path-suffix-values) - Optional List<br>Suffix Values. A list of path suffix values to match the input HTTP path against

<a id="policy-based-challenge-rule-list-rules-spec-path-transformers"></a>&#x2022; [`transformers`](#policy-based-challenge-rule-list-rules-spec-path-transformers) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### Policy Based Challenge Rule List Rules Spec Query Params

A [`query_params`](#policy-based-challenge-rule-list-rules-spec-query-params) block (within [`policy_based_challenge.rule_list.rules.spec`](#policy-based-challenge-rule-list-rules-spec)) supports the following:

<a id="policy-based-challenge-rule-list-rules-spec-query-params-check-not-present"></a>&#x2022; [`check_not_present`](#policy-based-challenge-rule-list-rules-spec-query-params-check-not-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="policy-based-challenge-rule-list-rules-spec-query-params-check-present"></a>&#x2022; [`check_present`](#policy-based-challenge-rule-list-rules-spec-query-params-check-present) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="policy-based-challenge-rule-list-rules-spec-query-params-invert-matcher"></a>&#x2022; [`invert_matcher`](#policy-based-challenge-rule-list-rules-spec-query-params-invert-matcher) - Optional Bool<br>Invert Query Parameter Matcher. Invert the match result

<a id="policy-based-challenge-rule-list-rules-spec-query-params-item"></a>&#x2022; [`item`](#policy-based-challenge-rule-list-rules-spec-query-params-item) - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#policy-based-challenge-rule-list-rules-spec-query-params-item) below.

<a id="policy-based-challenge-rule-list-rules-spec-query-params-key"></a>&#x2022; [`key`](#policy-based-challenge-rule-list-rules-spec-query-params-key) - Optional String<br>Query Parameter Name. A case-sensitive HTTP query parameter name

#### Policy Based Challenge Rule List Rules Spec Query Params Item

An [`item`](#policy-based-challenge-rule-list-rules-spec-query-params-item) block (within [`policy_based_challenge.rule_list.rules.spec.query_params`](#policy-based-challenge-rule-list-rules-spec-query-params)) supports the following:

<a id="policy-based-challenge-rule-list-rules-spec-query-params-item-exact-values"></a>&#x2022; [`exact_values`](#policy-based-challenge-rule-list-rules-spec-query-params-item-exact-values) - Optional List<br>Exact Values. A list of exact values to match the input against

<a id="policy-based-challenge-rule-list-rules-spec-query-params-item-regex-values"></a>&#x2022; [`regex_values`](#policy-based-challenge-rule-list-rules-spec-query-params-item-regex-values) - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="policy-based-challenge-rule-list-rules-spec-query-params-item-transformers"></a>&#x2022; [`transformers`](#policy-based-challenge-rule-list-rules-spec-query-params-item-transformers) - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

#### Policy Based Challenge Rule List Rules Spec TLS Fingerprint Matcher

A [`tls_fingerprint_matcher`](#policy-based-challenge-rule-list-rules-spec-tls-fingerprint-matcher) block (within [`policy_based_challenge.rule_list.rules.spec`](#policy-based-challenge-rule-list-rules-spec)) supports the following:

<a id="policy-based-challenge-rule-list-rules-spec-tls-fingerprint-matcher-classes"></a>&#x2022; [`classes`](#policy-based-challenge-rule-list-rules-spec-tls-fingerprint-matcher-classes) - Optional List  Defaults to `TLS_FINGERPRINT_NONE`<br>Possible values are `TLS_FINGERPRINT_NONE`, `ANY_MALICIOUS_FINGERPRINT`, `ADWARE`, `ADWIND`, `DRIDEX`, `GOOTKIT`, `GOZI`, `JBIFROST`, `QUAKBOT`, `RANSOMWARE`, `TROLDESH`, `TOFSEE`, `TORRENTLOCKER`, `TRICKBOT`<br>TLS fingerprint classes. A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against

<a id="policy-based-challenge-rule-list-rules-spec-tls-fingerprint-matcher-exact-values"></a>&#x2022; [`exact_values`](#policy-based-challenge-rule-list-rules-spec-tls-fingerprint-matcher-exact-values) - Optional List<br>Exact Values. A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against

<a id="policy-based-challenge-rule-list-rules-spec-tls-fingerprint-matcher-excluded-values"></a>&#x2022; [`excluded_values`](#policy-based-challenge-rule-list-rules-spec-tls-fingerprint-matcher-excluded-values) - Optional List<br>Excluded Values. A list of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can be used to skip known false positives when using one or more known TLS fingerprint classes in the enclosing matcher

#### Policy Based Challenge Temporary User Blocking

A [`temporary_user_blocking`](#policy-based-challenge-temporary-user-blocking) block (within [`policy_based_challenge`](#policy-based-challenge)) supports the following:

<a id="policy-based-challenge-temporary-user-blocking-custom-page"></a>&#x2022; [`custom_page`](#policy-based-challenge-temporary-user-blocking-custom-page) - Optional String<br>Custom Message for Temporary Blocking. Custom message is of type `uri_ref`. Currently supported URL schemes is `string:///`. For `string:///` scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. 'Blocked.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '`<p>` Blocked `</p>`'. Base64 encoded string for this HTML is 'PHA+IFBsZWFzZSBXYWl0IDwvcD4='

#### Protected Cookies

A [`protected_cookies`](#protected-cookies) block supports the following:

<a id="protected-cookies-add-httponly"></a>&#x2022; [`add_httponly`](#protected-cookies-add-httponly) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="protected-cookies-add-secure"></a>&#x2022; [`add_secure`](#protected-cookies-add-secure) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="protected-cookies-disable-tampering-protection"></a>&#x2022; [`disable_tampering_protection`](#protected-cookies-disable-tampering-protection) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="protected-cookies-enable-tampering-protection"></a>&#x2022; [`enable_tampering_protection`](#protected-cookies-enable-tampering-protection) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="protected-cookies-ignore-httponly"></a>&#x2022; [`ignore_httponly`](#protected-cookies-ignore-httponly) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="protected-cookies-ignore-max-age"></a>&#x2022; [`ignore_max_age`](#protected-cookies-ignore-max-age) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="protected-cookies-ignore-samesite"></a>&#x2022; [`ignore_samesite`](#protected-cookies-ignore-samesite) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="protected-cookies-ignore-secure"></a>&#x2022; [`ignore_secure`](#protected-cookies-ignore-secure) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="protected-cookies-max-age-value"></a>&#x2022; [`max_age_value`](#protected-cookies-max-age-value) - Optional Number<br>Add Max Age. Add max age attribute

<a id="protected-cookies-name"></a>&#x2022; [`name`](#protected-cookies-name) - Optional String<br>Cookie Name. Name of the Cookie

<a id="protected-cookies-samesite-lax"></a>&#x2022; [`samesite_lax`](#protected-cookies-samesite-lax) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="protected-cookies-samesite-none"></a>&#x2022; [`samesite_none`](#protected-cookies-samesite-none) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="protected-cookies-samesite-strict"></a>&#x2022; [`samesite_strict`](#protected-cookies-samesite-strict) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### Rate Limit

A [`rate_limit`](#rate-limit) block supports the following:

<a id="rate-limit-custom-ip-allowed-list"></a>&#x2022; [`custom_ip_allowed_list`](#rate-limit-custom-ip-allowed-list) - Optional Block<br>Custom IP Allowed List. IP Allowed list using existing ip_prefix_set objects<br>See [Custom IP Allowed List](#rate-limit-custom-ip-allowed-list) below.

<a id="rate-limit-ip-allowed-list"></a>&#x2022; [`ip_allowed_list`](#rate-limit-ip-allowed-list) - Optional Block<br>IPv4 Prefix List. x-example: '192.168.20.0/24' List of IPv4 prefixes that represent an endpoint<br>See [IP Allowed List](#rate-limit-ip-allowed-list) below.

<a id="rate-limit-no-ip-allowed-list"></a>&#x2022; [`no_ip_allowed_list`](#rate-limit-no-ip-allowed-list) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="rate-limit-no-policies"></a>&#x2022; [`no_policies`](#rate-limit-no-policies) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="rate-limit-policies"></a>&#x2022; [`policies`](#rate-limit-policies) - Optional Block<br>Rate Limiter Policy List. List of rate limiter policies to be applied<br>See [Policies](#rate-limit-policies) below.

<a id="rate-limit-rate-limiter"></a>&#x2022; [`rate_limiter`](#rate-limit-rate-limiter) - Optional Block<br>Rate Limit Value. A tuple consisting of a rate limit period unit and the total number of allowed requests for that period<br>See [Rate Limiter](#rate-limit-rate-limiter) below.

#### Rate Limit Custom IP Allowed List

A [`custom_ip_allowed_list`](#rate-limit-custom-ip-allowed-list) block (within [`rate_limit`](#rate-limit)) supports the following:

<a id="rate-limit-custom-ip-allowed-list-rate-limiter-allowed-prefixes"></a>&#x2022; [`rate_limiter_allowed_prefixes`](#rate-limit-custom-ip-allowed-list-rate-limiter-allowed-prefixes) - Optional Block<br>List of IP Prefix Sets. References to ip_prefix_set objects. Requests from source IP addresses that are covered by one of the allowed IP Prefixes are not subjected to rate limiting<br>See [Rate Limiter Allowed Prefixes](#rate-limit-custom-ip-allowed-list-rate-limiter-allowed-prefixes) below.

#### Rate Limit Custom IP Allowed List Rate Limiter Allowed Prefixes

A [`rate_limiter_allowed_prefixes`](#rate-limit-custom-ip-allowed-list-rate-limiter-allowed-prefixes) block (within [`rate_limit.custom_ip_allowed_list`](#rate-limit-custom-ip-allowed-list)) supports the following:

<a id="rate-limit-custom-ip-allowed-list-rate-limiter-allowed-prefixes-name"></a>&#x2022; [`name`](#rate-limit-custom-ip-allowed-list-rate-limiter-allowed-prefixes-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="rate-limit-custom-ip-allowed-list-rate-limiter-allowed-prefixes-namespace"></a>&#x2022; [`namespace`](#rate-limit-custom-ip-allowed-list-rate-limiter-allowed-prefixes-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="rate-limit-custom-ip-allowed-list-rate-limiter-allowed-prefixes-tenant"></a>&#x2022; [`tenant`](#rate-limit-custom-ip-allowed-list-rate-limiter-allowed-prefixes-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

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

<a id="rate-limit-rate-limiter-burst-multiplier"></a>&#x2022; [`burst_multiplier`](#rate-limit-rate-limiter-burst-multiplier) - Optional Number<br>Burst Multiplier. The maximum burst of requests to accommodate, expressed as a multiple of the rate

<a id="rate-limit-rate-limiter-disabled"></a>&#x2022; [`disabled`](#rate-limit-rate-limiter-disabled) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="rate-limit-rate-limiter-leaky-bucket"></a>&#x2022; [`leaky_bucket`](#rate-limit-rate-limiter-leaky-bucket) - Optional Block<br>Leaky Bucket Rate Limiter. Leaky-Bucket is the default rate limiter algorithm for F5

<a id="rate-limit-rate-limiter-period-multiplier"></a>&#x2022; [`period_multiplier`](#rate-limit-rate-limiter-period-multiplier) - Optional Number<br>Periods. This setting, combined with Per Period units, provides a duration

<a id="rate-limit-rate-limiter-token-bucket"></a>&#x2022; [`token_bucket`](#rate-limit-rate-limiter-token-bucket) - Optional Block<br>Token Bucket Rate Limiter. Token-Bucket is a rate limiter algorithm that is stricter with enforcing limits

<a id="rate-limit-rate-limiter-total-number"></a>&#x2022; [`total_number`](#rate-limit-rate-limiter-total-number) - Optional Number<br>Number Of Requests. The total number of allowed requests per rate-limiting period

<a id="rate-limit-rate-limiter-unit"></a>&#x2022; [`unit`](#rate-limit-rate-limiter-unit) - Optional String  Defaults to `SECOND`<br>Possible values are `SECOND`, `MINUTE`, `HOUR`<br>Rate Limit Period Unit. Unit for the period per which the rate limit is applied. - SECOND: Second Rate limit period unit is seconds - MINUTE: Minute Rate limit period unit is minutes - HOUR: Hour Rate limit period unit is hours - DAY: Day Rate limit period unit is days

#### Rate Limit Rate Limiter Action Block

An [`action_block`](#rate-limit-rate-limiter-action-block) block (within [`rate_limit.rate_limiter`](#rate-limit-rate-limiter)) supports the following:

<a id="rate-limit-rate-limiter-action-block-hours"></a>&#x2022; [`hours`](#rate-limit-rate-limiter-action-block-hours) - Optional Block<br>Hours. Input Duration Hours<br>See [Hours](#rate-limit-rate-limiter-action-block-hours) below.

<a id="rate-limit-rate-limiter-action-block-minutes"></a>&#x2022; [`minutes`](#rate-limit-rate-limiter-action-block-minutes) - Optional Block<br>Minutes. Input Duration Minutes<br>See [Minutes](#rate-limit-rate-limiter-action-block-minutes) below.

<a id="rate-limit-rate-limiter-action-block-seconds"></a>&#x2022; [`seconds`](#rate-limit-rate-limiter-action-block-seconds) - Optional Block<br>Seconds. Input Duration Seconds<br>See [Seconds](#rate-limit-rate-limiter-action-block-seconds) below.

#### Rate Limit Rate Limiter Action Block Hours

A [`hours`](#rate-limit-rate-limiter-action-block-hours) block (within [`rate_limit.rate_limiter.action_block`](#rate-limit-rate-limiter-action-block)) supports the following:

<a id="rate-limit-rate-limiter-action-block-hours-duration"></a>&#x2022; [`duration`](#rate-limit-rate-limiter-action-block-hours-duration) - Optional Number<br>Duration

#### Rate Limit Rate Limiter Action Block Minutes

A [`minutes`](#rate-limit-rate-limiter-action-block-minutes) block (within [`rate_limit.rate_limiter.action_block`](#rate-limit-rate-limiter-action-block)) supports the following:

<a id="rate-limit-rate-limiter-action-block-minutes-duration"></a>&#x2022; [`duration`](#rate-limit-rate-limiter-action-block-minutes-duration) - Optional Number<br>Duration

#### Rate Limit Rate Limiter Action Block Seconds

A [`seconds`](#rate-limit-rate-limiter-action-block-seconds) block (within [`rate_limit.rate_limiter.action_block`](#rate-limit-rate-limiter-action-block)) supports the following:

<a id="rate-limit-rate-limiter-action-block-seconds-duration"></a>&#x2022; [`duration`](#rate-limit-rate-limiter-action-block-seconds-duration) - Optional Number<br>Duration

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

<a id="ring-hash-hash-policy-cookie-add-httponly"></a>&#x2022; [`add_httponly`](#ring-hash-hash-policy-cookie-add-httponly) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="ring-hash-hash-policy-cookie-add-secure"></a>&#x2022; [`add_secure`](#ring-hash-hash-policy-cookie-add-secure) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="ring-hash-hash-policy-cookie-ignore-httponly"></a>&#x2022; [`ignore_httponly`](#ring-hash-hash-policy-cookie-ignore-httponly) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="ring-hash-hash-policy-cookie-ignore-samesite"></a>&#x2022; [`ignore_samesite`](#ring-hash-hash-policy-cookie-ignore-samesite) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="ring-hash-hash-policy-cookie-ignore-secure"></a>&#x2022; [`ignore_secure`](#ring-hash-hash-policy-cookie-ignore-secure) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="ring-hash-hash-policy-cookie-name"></a>&#x2022; [`name`](#ring-hash-hash-policy-cookie-name) - Optional String<br>Name. The name of the cookie that will be used to obtain the hash key. If the cookie is not present and TTL below is not set, no hash will be produced

<a id="ring-hash-hash-policy-cookie-path"></a>&#x2022; [`path`](#ring-hash-hash-policy-cookie-path) - Optional String<br>Path. The name of the path for the cookie. If no path is specified here, no path will be set for the cookie

<a id="ring-hash-hash-policy-cookie-samesite-lax"></a>&#x2022; [`samesite_lax`](#ring-hash-hash-policy-cookie-samesite-lax) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="ring-hash-hash-policy-cookie-samesite-none"></a>&#x2022; [`samesite_none`](#ring-hash-hash-policy-cookie-samesite-none) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="ring-hash-hash-policy-cookie-samesite-strict"></a>&#x2022; [`samesite_strict`](#ring-hash-hash-policy-cookie-samesite-strict) - Optional Block<br>Empty. This can be used for messages where no values are needed

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

<a id="routes-custom-route-object-route-ref-name"></a>&#x2022; [`name`](#routes-custom-route-object-route-ref-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="routes-custom-route-object-route-ref-namespace"></a>&#x2022; [`namespace`](#routes-custom-route-object-route-ref-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="routes-custom-route-object-route-ref-tenant"></a>&#x2022; [`tenant`](#routes-custom-route-object-route-ref-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Routes Direct Response Route

A [`direct_response_route`](#routes-direct-response-route) block (within [`routes`](#routes)) supports the following:

<a id="routes-direct-response-route-headers"></a>&#x2022; [`headers`](#routes-direct-response-route-headers) - Optional Block<br>Headers. List of (key, value) headers<br>See [Headers](#routes-direct-response-route-headers) below.

<a id="routes-direct-response-route-http-method"></a>&#x2022; [`http_method`](#routes-direct-response-route-http-method) - Optional String  Defaults to `ANY`<br>Possible values are `ANY`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`, `COPY`<br>HTTP Method. Specifies the HTTP method used to access a resource. Any HTTP Method

<a id="routes-direct-response-route-incoming-port"></a>&#x2022; [`incoming_port`](#routes-direct-response-route-incoming-port) - Optional Block<br>Port to Match. Port match of the request can be a range or a specific port<br>See [Incoming Port](#routes-direct-response-route-incoming-port) below.

<a id="routes-direct-response-route-path"></a>&#x2022; [`path`](#routes-direct-response-route-path) - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#routes-direct-response-route-path) below.

<a id="routes-direct-response-route-route-direct-response"></a>&#x2022; [`route_direct_response`](#routes-direct-response-route-route-direct-response) - Optional Block<br>Direct Response. Send this direct response in case of route match action is direct response<br>See [Route Direct Response](#routes-direct-response-route-route-direct-response) below.

#### Routes Direct Response Route Headers

A [`headers`](#routes-direct-response-route-headers) block (within [`routes.direct_response_route`](#routes-direct-response-route)) supports the following:

<a id="routes-direct-response-route-headers-exact"></a>&#x2022; [`exact`](#routes-direct-response-route-headers-exact) - Optional String<br>Exact. Header value to match exactly

<a id="routes-direct-response-route-headers-invert-match"></a>&#x2022; [`invert_match`](#routes-direct-response-route-headers-invert-match) - Optional Bool<br>NOT of match. Invert the result of the match to detect missing header or non-matching value

<a id="routes-direct-response-route-headers-name"></a>&#x2022; [`name`](#routes-direct-response-route-headers-name) - Optional String<br>Name. Name of the header

<a id="routes-direct-response-route-headers-presence"></a>&#x2022; [`presence`](#routes-direct-response-route-headers-presence) - Optional Bool<br>Presence. If true, check for presence of header

<a id="routes-direct-response-route-headers-regex"></a>&#x2022; [`regex`](#routes-direct-response-route-headers-regex) - Optional String<br>Regex. Regex match of the header value in re2 format

#### Routes Direct Response Route Incoming Port

An [`incoming_port`](#routes-direct-response-route-incoming-port) block (within [`routes.direct_response_route`](#routes-direct-response-route)) supports the following:

<a id="routes-direct-response-route-incoming-port-no-port-match"></a>&#x2022; [`no_port_match`](#routes-direct-response-route-incoming-port-no-port-match) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="routes-direct-response-route-incoming-port-port"></a>&#x2022; [`port`](#routes-direct-response-route-incoming-port-port) - Optional Number<br>Port. Exact Port to match

<a id="routes-direct-response-route-incoming-port-port-ranges"></a>&#x2022; [`port_ranges`](#routes-direct-response-route-incoming-port-port-ranges) - Optional String<br>Port range. Port range to match

#### Routes Direct Response Route Path

A [`path`](#routes-direct-response-route-path) block (within [`routes.direct_response_route`](#routes-direct-response-route)) supports the following:

<a id="routes-direct-response-route-path-path"></a>&#x2022; [`path`](#routes-direct-response-route-path-path) - Optional String<br>Exact. Exact path value to match

<a id="routes-direct-response-route-path-prefix"></a>&#x2022; [`prefix`](#routes-direct-response-route-path-prefix) - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

<a id="routes-direct-response-route-path-regex"></a>&#x2022; [`regex`](#routes-direct-response-route-path-regex) - Optional String<br>Regex. Regular expression of path match (e.g. the value .* will match on all paths)

#### Routes Direct Response Route Route Direct Response

A [`route_direct_response`](#routes-direct-response-route-route-direct-response) block (within [`routes.direct_response_route`](#routes-direct-response-route)) supports the following:

<a id="routes-direct-response-route-route-direct-response-response-body-encoded"></a>&#x2022; [`response_body_encoded`](#routes-direct-response-route-route-direct-response-response-body-encoded) - Optional String<br>Response Body. Response body to send. Currently supported URL schemes is string:/// for which message should be encoded in Base64 format. The message can be either plain text or HTML. E.g. '`<p>` Access Denied `</p>`'. Base64 encoded string URL for this is string:///PHA+IEFjY2VzcyBEZW5pZWQgPC9wPg==

<a id="routes-direct-response-route-route-direct-response-response-code"></a>&#x2022; [`response_code`](#routes-direct-response-route-route-direct-response-response-code) - Optional Number<br>Response Code. response code to send

#### Routes Redirect Route

A [`redirect_route`](#routes-redirect-route) block (within [`routes`](#routes)) supports the following:

<a id="routes-redirect-route-headers"></a>&#x2022; [`headers`](#routes-redirect-route-headers) - Optional Block<br>Headers. List of (key, value) headers<br>See [Headers](#routes-redirect-route-headers) below.

<a id="routes-redirect-route-http-method"></a>&#x2022; [`http_method`](#routes-redirect-route-http-method) - Optional String  Defaults to `ANY`<br>Possible values are `ANY`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`, `COPY`<br>HTTP Method. Specifies the HTTP method used to access a resource. Any HTTP Method

<a id="routes-redirect-route-incoming-port"></a>&#x2022; [`incoming_port`](#routes-redirect-route-incoming-port) - Optional Block<br>Port to Match. Port match of the request can be a range or a specific port<br>See [Incoming Port](#routes-redirect-route-incoming-port) below.

<a id="routes-redirect-route-path"></a>&#x2022; [`path`](#routes-redirect-route-path) - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#routes-redirect-route-path) below.

<a id="routes-redirect-route-route-redirect"></a>&#x2022; [`route_redirect`](#routes-redirect-route-route-redirect) - Optional Block<br>Redirect. route redirect parameters when match action is redirect<br>See [Route Redirect](#routes-redirect-route-route-redirect) below.

#### Routes Redirect Route Headers

A [`headers`](#routes-redirect-route-headers) block (within [`routes.redirect_route`](#routes-redirect-route)) supports the following:

<a id="routes-redirect-route-headers-exact"></a>&#x2022; [`exact`](#routes-redirect-route-headers-exact) - Optional String<br>Exact. Header value to match exactly

<a id="routes-redirect-route-headers-invert-match"></a>&#x2022; [`invert_match`](#routes-redirect-route-headers-invert-match) - Optional Bool<br>NOT of match. Invert the result of the match to detect missing header or non-matching value

<a id="routes-redirect-route-headers-name"></a>&#x2022; [`name`](#routes-redirect-route-headers-name) - Optional String<br>Name. Name of the header

<a id="routes-redirect-route-headers-presence"></a>&#x2022; [`presence`](#routes-redirect-route-headers-presence) - Optional Bool<br>Presence. If true, check for presence of header

<a id="routes-redirect-route-headers-regex"></a>&#x2022; [`regex`](#routes-redirect-route-headers-regex) - Optional String<br>Regex. Regex match of the header value in re2 format

#### Routes Redirect Route Incoming Port

An [`incoming_port`](#routes-redirect-route-incoming-port) block (within [`routes.redirect_route`](#routes-redirect-route)) supports the following:

<a id="routes-redirect-route-incoming-port-no-port-match"></a>&#x2022; [`no_port_match`](#routes-redirect-route-incoming-port-no-port-match) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="routes-redirect-route-incoming-port-port"></a>&#x2022; [`port`](#routes-redirect-route-incoming-port-port) - Optional Number<br>Port. Exact Port to match

<a id="routes-redirect-route-incoming-port-port-ranges"></a>&#x2022; [`port_ranges`](#routes-redirect-route-incoming-port-port-ranges) - Optional String<br>Port range. Port range to match

#### Routes Redirect Route Path

A [`path`](#routes-redirect-route-path) block (within [`routes.redirect_route`](#routes-redirect-route)) supports the following:

<a id="routes-redirect-route-path-path"></a>&#x2022; [`path`](#routes-redirect-route-path-path) - Optional String<br>Exact. Exact path value to match

<a id="routes-redirect-route-path-prefix"></a>&#x2022; [`prefix`](#routes-redirect-route-path-prefix) - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

<a id="routes-redirect-route-path-regex"></a>&#x2022; [`regex`](#routes-redirect-route-path-regex) - Optional String<br>Regex. Regular expression of path match (e.g. the value .* will match on all paths)

#### Routes Redirect Route Route Redirect

A [`route_redirect`](#routes-redirect-route-route-redirect) block (within [`routes.redirect_route`](#routes-redirect-route)) supports the following:

<a id="routes-redirect-route-route-redirect-host-redirect"></a>&#x2022; [`host_redirect`](#routes-redirect-route-route-redirect-host-redirect) - Optional String<br>Host. swap host part of incoming URL in redirect URL

<a id="routes-redirect-route-route-redirect-path-redirect"></a>&#x2022; [`path_redirect`](#routes-redirect-route-route-redirect-path-redirect) - Optional String<br>Path. swap path part of incoming URL in redirect URL

<a id="routes-redirect-route-route-redirect-prefix-rewrite"></a>&#x2022; [`prefix_rewrite`](#routes-redirect-route-route-redirect-prefix-rewrite) - Optional String<br>Prefix Rewrite. In Redirect response, the matched prefix (or path) should be swapped with this value. This option allows redirect URLs be dynamically created based on the request

<a id="routes-redirect-route-route-redirect-proto-redirect"></a>&#x2022; [`proto_redirect`](#routes-redirect-route-route-redirect-proto-redirect) - Optional String<br>Protocol. swap protocol part of incoming URL in redirect URL The protocol can be swapped with either HTTP or HTTPS When incoming-proto option is specified, swapping of protocol is not done

<a id="routes-redirect-route-route-redirect-remove-all-params"></a>&#x2022; [`remove_all_params`](#routes-redirect-route-route-redirect-remove-all-params) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="routes-redirect-route-route-redirect-replace-params"></a>&#x2022; [`replace_params`](#routes-redirect-route-route-redirect-replace-params) - Optional String<br>Replace All Parameters

<a id="routes-redirect-route-route-redirect-response-code"></a>&#x2022; [`response_code`](#routes-redirect-route-route-redirect-response-code) - Optional Number<br>Response Code. The HTTP status code to use in the redirect response

<a id="routes-redirect-route-route-redirect-retain-all-params"></a>&#x2022; [`retain_all_params`](#routes-redirect-route-route-redirect-retain-all-params) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### Routes Simple Route

A [`simple_route`](#routes-simple-route) block (within [`routes`](#routes)) supports the following:

<a id="routes-simple-route-advanced-options"></a>&#x2022; [`advanced_options`](#routes-simple-route-advanced-options) - Optional Block<br>Advanced Route Options. Configure advanced options for route like path rewrite, hash policy, etc<br>See [Advanced Options](#routes-simple-route-advanced-options) below.

<a id="routes-simple-route-auto-host-rewrite"></a>&#x2022; [`auto_host_rewrite`](#routes-simple-route-auto-host-rewrite) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="routes-simple-route-disable-host-rewrite"></a>&#x2022; [`disable_host_rewrite`](#routes-simple-route-disable-host-rewrite) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="routes-simple-route-headers"></a>&#x2022; [`headers`](#routes-simple-route-headers) - Optional Block<br>Headers. List of (key, value) headers<br>See [Headers](#routes-simple-route-headers) below.

<a id="routes-simple-route-host-rewrite"></a>&#x2022; [`host_rewrite`](#routes-simple-route-host-rewrite) - Optional String<br>Host Rewrite Value. Host header will be swapped with this value

<a id="routes-simple-route-http-method"></a>&#x2022; [`http_method`](#routes-simple-route-http-method) - Optional String  Defaults to `ANY`<br>Possible values are `ANY`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`, `COPY`<br>HTTP Method. Specifies the HTTP method used to access a resource. Any HTTP Method

<a id="routes-simple-route-incoming-port"></a>&#x2022; [`incoming_port`](#routes-simple-route-incoming-port) - Optional Block<br>Port to Match. Port match of the request can be a range or a specific port<br>See [Incoming Port](#routes-simple-route-incoming-port) below.

<a id="routes-simple-route-origin-pools"></a>&#x2022; [`origin_pools`](#routes-simple-route-origin-pools) - Optional Block<br>Origin Pools. Origin Pools for this route<br>See [Origin Pools](#routes-simple-route-origin-pools) below.

<a id="routes-simple-route-path"></a>&#x2022; [`path`](#routes-simple-route-path) - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#routes-simple-route-path) below.

<a id="routes-simple-route-query-params"></a>&#x2022; [`query_params`](#routes-simple-route-query-params) - Optional Block<br>Query Parameters. Handling of incoming query parameters in simple route<br>See [Query Params](#routes-simple-route-query-params) below.

#### Routes Simple Route Advanced Options

An [`advanced_options`](#routes-simple-route-advanced-options) block (within [`routes.simple_route`](#routes-simple-route)) supports the following:

<a id="routes-simple-route-advanced-options-app-firewall"></a>&#x2022; [`app_firewall`](#routes-simple-route-advanced-options-app-firewall) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [App Firewall](#routes-simple-route-advanced-options-app-firewall) below.

<a id="routes-simple-route-advanced-options-bot-defense-javascript-injection"></a>&#x2022; [`bot_defense_javascript_injection`](#routes-simple-route-advanced-options-bot-defense-javascript-injection) - Optional Block<br>Bot Defense Javascript Injection Configuration for inline deployments. Bot Defense Javascript Injection Configuration for inline bot defense deployments<br>See [Bot Defense Javascript Injection](#routes-simple-route-advanced-options-bot-defense-javascript-injection) below.

<a id="routes-simple-route-advanced-options-buffer-policy"></a>&#x2022; [`buffer_policy`](#routes-simple-route-advanced-options-buffer-policy) - Optional Block<br>Buffer Configuration. Some upstream applications are not capable of handling streamed data. This config enables buffering the entire request before sending to upstream application. We can specify the maximum buffer size and buffer interval with this config. Buffering can be enabled and disabled at VirtualHost and Route levels Route level buffer configuration takes precedence<br>See [Buffer Policy](#routes-simple-route-advanced-options-buffer-policy) below.

<a id="routes-simple-route-advanced-options-common-buffering"></a>&#x2022; [`common_buffering`](#routes-simple-route-advanced-options-common-buffering) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="routes-simple-route-advanced-options-common-hash-policy"></a>&#x2022; [`common_hash_policy`](#routes-simple-route-advanced-options-common-hash-policy) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="routes-simple-route-advanced-options-cors-policy"></a>&#x2022; [`cors_policy`](#routes-simple-route-advanced-options-cors-policy) - Optional Block<br>CORS Policy. Cross-Origin Resource Sharing requests configuration specified at Virtual-host or Route level. Route level configuration takes precedence. An example of an Cross origin HTTP request GET /resources/public-data/ HTTP/1.1 Host: bar.other User-Agent: Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.5; en-US; rv:1.9.1b3pre) Gecko/20081130 Minefield/3.1b3pre Accept: text/HTML,application/xhtml+XML,application/XML;q=0.9,*/*;q=0.8 Accept-Language: en-us,en;q=0.5 Accept-Encoding: gzip,deflate Accept-Charset: ISO-8859-1,utf-8;q=0.7,*;q=0.7 Connection: keep-alive Referrer: `HTTP://foo.example/examples/access-control/simpleXSInvocation.HTML` Origin: `HTTP://foo.example` HTTP/1.1 200 OK Date: Mon, 01 Dec 2008 00:23:53 GMT Server: Apache/2.0.61 Access-Control-Allow-Origin: * Keep-Alive: timeout=2, max=100 Connection: Keep-Alive Transfer-Encoding: chunked Content-Type: application/XML An example for cross origin HTTP OPTIONS request with Access-Control-Request-* header OPTIONS /resources/post-here/ HTTP/1.1 Host: bar.other User-Agent: Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.5; en-US; rv:1.9.1b3pre) Gecko/20081130 Minefield/3.1b3pre Accept: text/HTML,application/xhtml+XML,application/XML;q=0.9,*/*;q=0.8 Accept-Language: en-us,en;q=0.5 Accept-Encoding: gzip,deflate Accept-Charset: ISO-8859-1,utf-8;q=0.7,*;q=0.7 Connection: keep-alive Origin: `HTTP://foo.example` Access-Control-Request-Method: POST Access-Control-Request-Headers: X-PINGOTHER, Content-Type HTTP/1.1 204 No Content Date: Mon, 01 Dec 2008 01:15:39 GMT Server: Apache/2.0.61 (Unix) Access-Control-Allow-Origin: `HTTP://foo.example` Access-Control-Allow-Methods: POST, GET, OPTIONS Access-Control-Allow-Headers: X-PINGOTHER, Content-Type Access-Control-Max-Age: 86400 Vary: Accept-Encoding, Origin Keep-Alive: timeout=2, max=100 Connection: Keep-Alive<br>See [CORS Policy](#routes-simple-route-advanced-options-cors-policy) below.

<a id="routes-simple-route-advanced-options-csrf-policy"></a>&#x2022; [`csrf_policy`](#routes-simple-route-advanced-options-csrf-policy) - Optional Block<br>CSRF Policy. To mitigate CSRF attack , the policy checks where a request is coming from to determine if the request's origin is the same as its detination.The policy relies on two pieces of information used in determining if a request originated from the same host. 1. The origin that caused the user agent to issue the request (source origin). 2. The origin that the request is going to (target origin). When the policy evaluating a request, it ensures both pieces of information are present and compare their values. If the source origin is missing or origins do not match the request is rejected. The exception to this being if the source-origin has been added to they policy as valid. Because CSRF attacks specifically target state-changing requests, the policy only acts on the HTTP requests that have state-changing method (PUT,POST, etc.)<br>See [CSRF Policy](#routes-simple-route-advanced-options-csrf-policy) below.

<a id="routes-simple-route-advanced-options-default-retry-policy"></a>&#x2022; [`default_retry_policy`](#routes-simple-route-advanced-options-default-retry-policy) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="routes-simple-route-advanced-options-disable-location-add"></a>&#x2022; [`disable_location_add`](#routes-simple-route-advanced-options-disable-location-add) - Optional Bool<br>Disable Location Addition. disables append of x-volterra-location = `<RE-site-name>` at route level, if it is configured at virtual-host level. This configuration is ignored on CE sites

<a id="routes-simple-route-advanced-options-disable-mirroring"></a>&#x2022; [`disable_mirroring`](#routes-simple-route-advanced-options-disable-mirroring) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="routes-simple-route-advanced-options-disable-prefix-rewrite"></a>&#x2022; [`disable_prefix_rewrite`](#routes-simple-route-advanced-options-disable-prefix-rewrite) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="routes-simple-route-advanced-options-disable-spdy"></a>&#x2022; [`disable_spdy`](#routes-simple-route-advanced-options-disable-spdy) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="routes-simple-route-advanced-options-disable-waf"></a>&#x2022; [`disable_waf`](#routes-simple-route-advanced-options-disable-waf) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="routes-simple-route-advanced-options-disable-web-socket-config"></a>&#x2022; [`disable_web_socket_config`](#routes-simple-route-advanced-options-disable-web-socket-config) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="routes-simple-route-advanced-options-do-not-retract-cluster"></a>&#x2022; [`do_not_retract_cluster`](#routes-simple-route-advanced-options-do-not-retract-cluster) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="routes-simple-route-advanced-options-enable-spdy"></a>&#x2022; [`enable_spdy`](#routes-simple-route-advanced-options-enable-spdy) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="routes-simple-route-advanced-options-endpoint-subsets"></a>&#x2022; [`endpoint_subsets`](#routes-simple-route-advanced-options-endpoint-subsets) - Optional Block<br>Origin Servers Subsets. Upstream origin pool may be configured to divide its origin servers into subsets based on metadata attached to the origin servers. Routes may then specify the metadata that a endpoint must match in order to be selected by the load balancer For origin servers which are discovered in K8S or Consul cluster, the label of the service is merged with endpoint's labels. In case of Consul, the label is derived from the 'Tag' field. For labels that are common between configured endpoint and discovered service, labels from discovered service takes precedence. List of key-value pairs that will be used as matching metadata. Only those origin servers of upstream origin pool which match this metadata will be selected for load balancing

<a id="routes-simple-route-advanced-options-inherited-bot-defense-javascript-injection"></a>&#x2022; [`inherited_bot_defense_javascript_injection`](#routes-simple-route-advanced-options-inherited-bot-defense-javascript-injection) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="routes-simple-route-advanced-options-inherited-waf"></a>&#x2022; [`inherited_waf`](#routes-simple-route-advanced-options-inherited-waf) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="routes-simple-route-advanced-options-inherited-waf-exclusion"></a>&#x2022; [`inherited_waf_exclusion`](#routes-simple-route-advanced-options-inherited-waf-exclusion) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="routes-simple-route-advanced-options-mirror-policy"></a>&#x2022; [`mirror_policy`](#routes-simple-route-advanced-options-mirror-policy) - Optional Block<br>Mirror Policy. MirrorPolicy is used for shadowing traffic from one origin pool to another. The approach used is 'fire and forget', meaning it will not wait for the shadow origin pool to respond before returning the response from the primary origin pool. All normal statistics are collected for the shadow origin pool making this feature useful for testing and troubleshooting<br>See [Mirror Policy](#routes-simple-route-advanced-options-mirror-policy) below.

<a id="routes-simple-route-advanced-options-no-retry-policy"></a>&#x2022; [`no_retry_policy`](#routes-simple-route-advanced-options-no-retry-policy) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="routes-simple-route-advanced-options-prefix-rewrite"></a>&#x2022; [`prefix_rewrite`](#routes-simple-route-advanced-options-prefix-rewrite) - Optional String<br>Enable Prefix Rewrite. prefix_rewrite indicates that during forwarding, the matched prefix (or path) should be swapped with its value. When using regex path matching, the entire path (not including the query string) will be swapped with this value

<a id="routes-simple-route-advanced-options-priority"></a>&#x2022; [`priority`](#routes-simple-route-advanced-options-priority) - Optional String  Defaults to `DEFAULT`<br>Possible values are `DEFAULT`, `HIGH`<br>Routing Priority. Priority routing for each request. Different connection pools are used based on the priority selected for the request. Also, circuit-breaker configuration at destination cluster is chosen based on selected priority. Default routing mechanism High-Priority routing mechanism

<a id="routes-simple-route-advanced-options-regex-rewrite"></a>&#x2022; [`regex_rewrite`](#routes-simple-route-advanced-options-regex-rewrite) - Optional Block<br>Regex Match Rewrite. RegexMatchRewrite describes how to match a string and then produce a new string using a regular expression and a substitution string<br>See [Regex Rewrite](#routes-simple-route-advanced-options-regex-rewrite) below.

<a id="routes-simple-route-advanced-options-request-cookies-to-add"></a>&#x2022; [`request_cookies_to_add`](#routes-simple-route-advanced-options-request-cookies-to-add) - Optional Block<br>Add Cookies in Cookie Header. Cookies are key-value pairs to be added to HTTP request being routed towards upstream. Cookies specified at this level are applied after cookies from matched Route are applied<br>See [Request Cookies To Add](#routes-simple-route-advanced-options-request-cookies-to-add) below.

<a id="routes-simple-route-advanced-options-request-cookies-to-remove"></a>&#x2022; [`request_cookies_to_remove`](#routes-simple-route-advanced-options-request-cookies-to-remove) - Optional List<br>Remove Cookies from Cookie Header. List of keys of Cookies to be removed from the HTTP request being sent towards upstream

<a id="routes-simple-route-advanced-options-request-headers-to-add"></a>&#x2022; [`request_headers_to_add`](#routes-simple-route-advanced-options-request-headers-to-add) - Optional Block<br>Add Request Headers. Headers are key-value pairs to be added to HTTP request being routed towards upstream<br>See [Request Headers To Add](#routes-simple-route-advanced-options-request-headers-to-add) below.

<a id="routes-simple-route-advanced-options-request-headers-to-remove"></a>&#x2022; [`request_headers_to_remove`](#routes-simple-route-advanced-options-request-headers-to-remove) - Optional List<br>Remove Request Headers. List of keys of Headers to be removed from the HTTP request being sent towards upstream

<a id="routes-simple-route-advanced-options-response-cookies-to-add"></a>&#x2022; [`response_cookies_to_add`](#routes-simple-route-advanced-options-response-cookies-to-add) - Optional Block<br>Add Set-Cookie Headers. Cookies are name-value pairs along with optional attribute parameters to be added to HTTP response being sent towards downstream. Cookies specified at this level are applied after cookies from matched Route are applied<br>See [Response Cookies To Add](#routes-simple-route-advanced-options-response-cookies-to-add) below.

<a id="routes-simple-route-advanced-options-response-cookies-to-remove"></a>&#x2022; [`response_cookies_to_remove`](#routes-simple-route-advanced-options-response-cookies-to-remove) - Optional List<br>Remove Cookies from Set-Cookie Headers. List of name of Cookies to be removed from the HTTP response being sent towards downstream. Entire set-cookie header will be removed

<a id="routes-simple-route-advanced-options-response-headers-to-add"></a>&#x2022; [`response_headers_to_add`](#routes-simple-route-advanced-options-response-headers-to-add) - Optional Block<br>Add Response Headers. Headers are key-value pairs to be added to HTTP response being sent towards downstream<br>See [Response Headers To Add](#routes-simple-route-advanced-options-response-headers-to-add) below.

<a id="routes-simple-route-advanced-options-response-headers-to-remove"></a>&#x2022; [`response_headers_to_remove`](#routes-simple-route-advanced-options-response-headers-to-remove) - Optional List<br>Remove Response Headers. List of keys of Headers to be removed from the HTTP response being sent towards downstream

<a id="routes-simple-route-advanced-options-retract-cluster"></a>&#x2022; [`retract_cluster`](#routes-simple-route-advanced-options-retract-cluster) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="routes-simple-route-advanced-options-retry-policy"></a>&#x2022; [`retry_policy`](#routes-simple-route-advanced-options-retry-policy) - Optional Block<br>Retry Policy. Retry policy configuration for route destination<br>See [Retry Policy](#routes-simple-route-advanced-options-retry-policy) below.

<a id="routes-simple-route-advanced-options-specific-hash-policy"></a>&#x2022; [`specific_hash_policy`](#routes-simple-route-advanced-options-specific-hash-policy) - Optional Block<br>Hash Policy List. List of hash policy rules<br>See [Specific Hash Policy](#routes-simple-route-advanced-options-specific-hash-policy) below.

<a id="routes-simple-route-advanced-options-timeout"></a>&#x2022; [`timeout`](#routes-simple-route-advanced-options-timeout) - Optional Number<br>Timeout. The timeout for the route including all retries, in milliseconds. Should be set to a high value or 0 (infinite timeout) for server-side streaming

<a id="routes-simple-route-advanced-options-waf-exclusion-policy"></a>&#x2022; [`waf_exclusion_policy`](#routes-simple-route-advanced-options-waf-exclusion-policy) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [WAF Exclusion Policy](#routes-simple-route-advanced-options-waf-exclusion-policy) below.

<a id="routes-simple-route-advanced-options-web-socket-config"></a>&#x2022; [`web_socket_config`](#routes-simple-route-advanced-options-web-socket-config) - Optional Block<br>WebSocket Configuration. Configuration to allow WebSocket Request headers of such upgrade looks like below 'connection', 'Upgrade' 'upgrade', 'WebSocket' With configuration to allow WebSocket upgrade, ADC will produce following response 'HTTP/1.1 101 Switching Protocols 'Upgrade': 'WebSocket' 'Connection': 'Upgrade'<br>See [Web Socket Config](#routes-simple-route-advanced-options-web-socket-config) below.

#### Routes Simple Route Advanced Options App Firewall

An [`app_firewall`](#routes-simple-route-advanced-options-app-firewall) block (within [`routes.simple_route.advanced_options`](#routes-simple-route-advanced-options)) supports the following:

<a id="routes-simple-route-advanced-options-app-firewall-name"></a>&#x2022; [`name`](#routes-simple-route-advanced-options-app-firewall-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="routes-simple-route-advanced-options-app-firewall-namespace"></a>&#x2022; [`namespace`](#routes-simple-route-advanced-options-app-firewall-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="routes-simple-route-advanced-options-app-firewall-tenant"></a>&#x2022; [`tenant`](#routes-simple-route-advanced-options-app-firewall-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Routes Simple Route Advanced Options Bot Defense Javascript Injection

A [`bot_defense_javascript_injection`](#routes-simple-route-advanced-options-bot-defense-javascript-injection) block (within [`routes.simple_route.advanced_options`](#routes-simple-route-advanced-options)) supports the following:

<a id="routes-simple-route-advanced-options-bot-defense-javascript-injection-javascript-location"></a>&#x2022; [`javascript_location`](#routes-simple-route-advanced-options-bot-defense-javascript-injection-javascript-location) - Optional String  Defaults to `AFTER_HEAD`<br>Possible values are `AFTER_HEAD`, `AFTER_TITLE_END`, `BEFORE_SCRIPT`<br>JavaScript Location. All inside networks. Insert JavaScript after `<head>` tag Insert JavaScript after `</title>` tag. Insert JavaScript before first `<script>` tag

<a id="routes-simple-route-advanced-options-bot-defense-javascript-injection-javascript-tags"></a>&#x2022; [`javascript_tags`](#routes-simple-route-advanced-options-bot-defense-javascript-injection-javascript-tags) - Optional Block<br>JavaScript Tags. Select Add item to configure your javascript tag. If adding both Bot Adv and Fraud, the Bot Javascript should be added first<br>See [Javascript Tags](#routes-simple-route-advanced-options-bot-defense-javascript-injection-javascript-tags) below.

#### Routes Simple Route Advanced Options Bot Defense Javascript Injection Javascript Tags

A [`javascript_tags`](#routes-simple-route-advanced-options-bot-defense-javascript-injection-javascript-tags) block (within [`routes.simple_route.advanced_options.bot_defense_javascript_injection`](#routes-simple-route-advanced-options-bot-defense-javascript-injection)) supports the following:

<a id="routes-simple-route-advanced-options-bot-defense-javascript-injection-javascript-tags-javascript-url"></a>&#x2022; [`javascript_url`](#routes-simple-route-advanced-options-bot-defense-javascript-injection-javascript-tags-javascript-url) - Optional String<br>URL. Please enter the full URL (include domain and path), or relative path

<a id="routes-simple-route-advanced-options-bot-defense-javascript-injection-javascript-tags-tag-attributes"></a>&#x2022; [`tag_attributes`](#routes-simple-route-advanced-options-bot-defense-javascript-injection-javascript-tags-tag-attributes) - Optional Block<br>Tag Attributes. Add the tag attributes you want to include in your Javascript tag<br>See [Tag Attributes](#routes-simple-route-advanced-options-bot-defense-javascript-injection-javascript-tags-tag-attributes) below.

#### Routes Simple Route Advanced Options Bot Defense Javascript Injection Javascript Tags Tag Attributes

A [`tag_attributes`](#routes-simple-route-advanced-options-bot-defense-javascript-injection-javascript-tags-tag-attributes) block (within [`routes.simple_route.advanced_options.bot_defense_javascript_injection.javascript_tags`](#routes-simple-route-advanced-options-bot-defense-javascript-injection-javascript-tags)) supports the following:

<a id="routes-simple-route-advanced-options-bot-defense-javascript-injection-javascript-tags-tag-attributes-javascript-tag"></a>&#x2022; [`javascript_tag`](#routes-simple-route-advanced-options-bot-defense-javascript-injection-javascript-tags-tag-attributes-javascript-tag) - Optional String  Defaults to `JS_ATTR_ID`<br>Possible values are `JS_ATTR_ID`, `JS_ATTR_CID`, `JS_ATTR_CN`, `JS_ATTR_API_DOMAIN`, `JS_ATTR_API_URL`, `JS_ATTR_API_PATH`, `JS_ATTR_ASYNC`, `JS_ATTR_DEFER`<br>Tag Attribute Name. Select from one of the predefined tag attributes

<a id="routes-simple-route-advanced-options-bot-defense-javascript-injection-javascript-tags-tag-attributes-tag-value"></a>&#x2022; [`tag_value`](#routes-simple-route-advanced-options-bot-defense-javascript-injection-javascript-tags-tag-attributes-tag-value) - Optional String<br>Value. Add the tag attribute value

#### Routes Simple Route Advanced Options Buffer Policy

A [`buffer_policy`](#routes-simple-route-advanced-options-buffer-policy) block (within [`routes.simple_route.advanced_options`](#routes-simple-route-advanced-options)) supports the following:

<a id="routes-simple-route-advanced-options-buffer-policy-disabled"></a>&#x2022; [`disabled`](#routes-simple-route-advanced-options-buffer-policy-disabled) - Optional Bool<br>Disable. Disable buffering for a particular route. This is useful when virtual-host has buffering, but we need to disable it on a specific route. The value of this field is ignored for virtual-host

<a id="routes-simple-route-advanced-options-buffer-policy-max-request-bytes"></a>&#x2022; [`max_request_bytes`](#routes-simple-route-advanced-options-buffer-policy-max-request-bytes) - Optional Number<br>Max Request Bytes. The maximum request size that the filter will buffer before the connection manager will stop buffering and return a RequestEntityTooLarge (413) response

#### Routes Simple Route Advanced Options CORS Policy

A [`cors_policy`](#routes-simple-route-advanced-options-cors-policy) block (within [`routes.simple_route.advanced_options`](#routes-simple-route-advanced-options)) supports the following:

<a id="routes-simple-route-advanced-options-cors-policy-allow-credentials"></a>&#x2022; [`allow_credentials`](#routes-simple-route-advanced-options-cors-policy-allow-credentials) - Optional Bool<br>Allow Credentials. Specifies whether the resource allows credentials

<a id="routes-simple-route-advanced-options-cors-policy-allow-headers"></a>&#x2022; [`allow_headers`](#routes-simple-route-advanced-options-cors-policy-allow-headers) - Optional String<br>Allow Headers. Specifies the content for the access-control-allow-headers header

<a id="routes-simple-route-advanced-options-cors-policy-allow-methods"></a>&#x2022; [`allow_methods`](#routes-simple-route-advanced-options-cors-policy-allow-methods) - Optional String<br>Allow Methods. Specifies the content for the access-control-allow-methods header

<a id="routes-simple-route-advanced-options-cors-policy-allow-origin"></a>&#x2022; [`allow_origin`](#routes-simple-route-advanced-options-cors-policy-allow-origin) - Optional List<br>Allow Origin. Specifies the origins that will be allowed to do CORS requests. An origin is allowed if either allow_origin or allow_origin_regex match

<a id="routes-simple-route-advanced-options-cors-policy-allow-origin-regex"></a>&#x2022; [`allow_origin_regex`](#routes-simple-route-advanced-options-cors-policy-allow-origin-regex) - Optional List<br>Allow Origin Regex. Specifies regex patterns that match allowed origins. An origin is allowed if either allow_origin or allow_origin_regex match

<a id="routes-simple-route-advanced-options-cors-policy-disabled"></a>&#x2022; [`disabled`](#routes-simple-route-advanced-options-cors-policy-disabled) - Optional Bool<br>Disabled. Disable the CorsPolicy for a particular route. This is useful when virtual-host has CorsPolicy, but we need to disable it on a specific route. The value of this field is ignored for virtual-host

<a id="routes-simple-route-advanced-options-cors-policy-expose-headers"></a>&#x2022; [`expose_headers`](#routes-simple-route-advanced-options-cors-policy-expose-headers) - Optional String<br>Expose Headers. Specifies the content for the access-control-expose-headers header

<a id="routes-simple-route-advanced-options-cors-policy-maximum-age"></a>&#x2022; [`maximum_age`](#routes-simple-route-advanced-options-cors-policy-maximum-age) - Optional Number<br>Maximum Age. Specifies the content for the access-control-max-age header in seconds. This indicates the maximum number of seconds the results can be cached A value of -1 will disable caching. Maximum permitted value is 86400 seconds (24 hours)

#### Routes Simple Route Advanced Options CSRF Policy

A [`csrf_policy`](#routes-simple-route-advanced-options-csrf-policy) block (within [`routes.simple_route.advanced_options`](#routes-simple-route-advanced-options)) supports the following:

<a id="routes-simple-route-advanced-options-csrf-policy-all-load-balancer-domains"></a>&#x2022; [`all_load_balancer_domains`](#routes-simple-route-advanced-options-csrf-policy-all-load-balancer-domains) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="routes-simple-route-advanced-options-csrf-policy-custom-domain-list"></a>&#x2022; [`custom_domain_list`](#routes-simple-route-advanced-options-csrf-policy-custom-domain-list) - Optional Block<br>Domain name list. List of domain names used for Host header matching<br>See [Custom Domain List](#routes-simple-route-advanced-options-csrf-policy-custom-domain-list) below.

<a id="routes-simple-route-advanced-options-csrf-policy-disabled"></a>&#x2022; [`disabled`](#routes-simple-route-advanced-options-csrf-policy-disabled) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### Routes Simple Route Advanced Options CSRF Policy Custom Domain List

A [`custom_domain_list`](#routes-simple-route-advanced-options-csrf-policy-custom-domain-list) block (within [`routes.simple_route.advanced_options.csrf_policy`](#routes-simple-route-advanced-options-csrf-policy)) supports the following:

<a id="routes-simple-route-advanced-options-csrf-policy-custom-domain-list-domains"></a>&#x2022; [`domains`](#routes-simple-route-advanced-options-csrf-policy-custom-domain-list-domains) - Optional List<br>Domain names. A list of domain names that will be matched to loadbalancer. These domains are not used for SNI match. Wildcard names are supported in the suffix or prefix form

#### Routes Simple Route Advanced Options Mirror Policy

A [`mirror_policy`](#routes-simple-route-advanced-options-mirror-policy) block (within [`routes.simple_route.advanced_options`](#routes-simple-route-advanced-options)) supports the following:

<a id="routes-simple-route-advanced-options-mirror-policy-origin-pool"></a>&#x2022; [`origin_pool`](#routes-simple-route-advanced-options-mirror-policy-origin-pool) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Origin Pool](#routes-simple-route-advanced-options-mirror-policy-origin-pool) below.

<a id="routes-simple-route-advanced-options-mirror-policy-percent"></a>&#x2022; [`percent`](#routes-simple-route-advanced-options-mirror-policy-percent) - Optional Block<br>Fractional Percent. Fraction used where sampling percentages are needed. example sampled requests<br>See [Percent](#routes-simple-route-advanced-options-mirror-policy-percent) below.

#### Routes Simple Route Advanced Options Mirror Policy Origin Pool

An [`origin_pool`](#routes-simple-route-advanced-options-mirror-policy-origin-pool) block (within [`routes.simple_route.advanced_options.mirror_policy`](#routes-simple-route-advanced-options-mirror-policy)) supports the following:

<a id="routes-simple-route-advanced-options-mirror-policy-origin-pool-name"></a>&#x2022; [`name`](#routes-simple-route-advanced-options-mirror-policy-origin-pool-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="routes-simple-route-advanced-options-mirror-policy-origin-pool-namespace"></a>&#x2022; [`namespace`](#routes-simple-route-advanced-options-mirror-policy-origin-pool-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="routes-simple-route-advanced-options-mirror-policy-origin-pool-tenant"></a>&#x2022; [`tenant`](#routes-simple-route-advanced-options-mirror-policy-origin-pool-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Routes Simple Route Advanced Options Mirror Policy Percent

A [`percent`](#routes-simple-route-advanced-options-mirror-policy-percent) block (within [`routes.simple_route.advanced_options.mirror_policy`](#routes-simple-route-advanced-options-mirror-policy)) supports the following:

<a id="routes-simple-route-advanced-options-mirror-policy-percent-denominator"></a>&#x2022; [`denominator`](#routes-simple-route-advanced-options-mirror-policy-percent-denominator) - Optional String  Defaults to `HUNDRED`<br>Possible values are `HUNDRED`, `TEN_THOUSAND`, `MILLION`<br>Denominator. Denominator used in fraction where sampling percentages are needed. example sampled requests Use hundred as denominator Use ten thousand as denominator Use million as denominator

<a id="routes-simple-route-advanced-options-mirror-policy-percent-numerator"></a>&#x2022; [`numerator`](#routes-simple-route-advanced-options-mirror-policy-percent-numerator) - Optional Number<br>Numerator. sampled parts per denominator. If denominator was 10000, then value of 5 will be 5 in 10000

#### Routes Simple Route Advanced Options Regex Rewrite

A [`regex_rewrite`](#routes-simple-route-advanced-options-regex-rewrite) block (within [`routes.simple_route.advanced_options`](#routes-simple-route-advanced-options)) supports the following:

<a id="routes-simple-route-advanced-options-regex-rewrite-pattern"></a>&#x2022; [`pattern`](#routes-simple-route-advanced-options-regex-rewrite-pattern) - Optional String<br>Pattern. The regular expression used to find portions of a string that should be replaced

<a id="routes-simple-route-advanced-options-regex-rewrite-substitution"></a>&#x2022; [`substitution`](#routes-simple-route-advanced-options-regex-rewrite-substitution) - Optional String<br>Substitution. The string that should be substituted into matching portions of the subject string during a substitution operation to produce a new string

#### Routes Simple Route Advanced Options Request Cookies To Add

A [`request_cookies_to_add`](#routes-simple-route-advanced-options-request-cookies-to-add) block (within [`routes.simple_route.advanced_options`](#routes-simple-route-advanced-options)) supports the following:

<a id="routes-simple-route-advanced-options-request-cookies-to-add-name"></a>&#x2022; [`name`](#routes-simple-route-advanced-options-request-cookies-to-add-name) - Optional String<br>Name. Name of the cookie in Cookie header

<a id="routes-simple-route-advanced-options-request-cookies-to-add-overwrite"></a>&#x2022; [`overwrite`](#routes-simple-route-advanced-options-request-cookies-to-add-overwrite) - Optional Bool  Defaults to `do`<br>Overwrite. Should the value be overwritten? If true, the value is overwritten to existing values. not overwrite

<a id="routes-simple-route-advanced-options-request-cookies-to-add-secret-value"></a>&#x2022; [`secret_value`](#routes-simple-route-advanced-options-request-cookies-to-add-secret-value) - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Secret Value](#routes-simple-route-advanced-options-request-cookies-to-add-secret-value) below.

<a id="routes-simple-route-advanced-options-request-cookies-to-add-value"></a>&#x2022; [`value`](#routes-simple-route-advanced-options-request-cookies-to-add-value) - Optional String<br>Value. Value of the Cookie header

#### Routes Simple Route Advanced Options Request Cookies To Add Secret Value

A [`secret_value`](#routes-simple-route-advanced-options-request-cookies-to-add-secret-value) block (within [`routes.simple_route.advanced_options.request_cookies_to_add`](#routes-simple-route-advanced-options-request-cookies-to-add)) supports the following:

<a id="routes-simple-route-advanced-options-request-cookies-to-add-secret-value-blindfold-secret-info"></a>&#x2022; [`blindfold_secret_info`](#routes-simple-route-advanced-options-request-cookies-to-add-secret-value-blindfold-secret-info) - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#routes-simple-route-advanced-options-request-cookies-to-add-secret-value-blindfold-secret-info) below.

<a id="routes-simple-route-advanced-options-request-cookies-to-add-secret-value-clear-secret-info"></a>&#x2022; [`clear_secret_info`](#routes-simple-route-advanced-options-request-cookies-to-add-secret-value-clear-secret-info) - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#routes-simple-route-advanced-options-request-cookies-to-add-secret-value-clear-secret-info) below.

#### Routes Simple Route Advanced Options Request Cookies To Add Secret Value Blindfold Secret Info

A [`blindfold_secret_info`](#routes-simple-route-advanced-options-request-cookies-to-add-secret-value-blindfold-secret-info) block (within [`routes.simple_route.advanced_options.request_cookies_to_add.secret_value`](#routes-simple-route-advanced-options-request-cookies-to-add-secret-value)) supports the following:

<a id="routes-simple-route-advanced-options-request-cookies-to-add-secret-value-blindfold-secret-info-decryption-provider"></a>&#x2022; [`decryption_provider`](#routes-simple-route-advanced-options-request-cookies-to-add-secret-value-blindfold-secret-info-decryption-provider) - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

<a id="routes-simple-route-advanced-options-request-cookies-to-add-secret-value-blindfold-secret-info-location"></a>&#x2022; [`location`](#routes-simple-route-advanced-options-request-cookies-to-add-secret-value-blindfold-secret-info-location) - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

<a id="routes-simple-route-advanced-options-request-cookies-to-add-secret-value-blindfold-secret-info-store-provider"></a>&#x2022; [`store_provider`](#routes-simple-route-advanced-options-request-cookies-to-add-secret-value-blindfold-secret-info-store-provider) - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

#### Routes Simple Route Advanced Options Request Cookies To Add Secret Value Clear Secret Info

A [`clear_secret_info`](#routes-simple-route-advanced-options-request-cookies-to-add-secret-value-clear-secret-info) block (within [`routes.simple_route.advanced_options.request_cookies_to_add.secret_value`](#routes-simple-route-advanced-options-request-cookies-to-add-secret-value)) supports the following:

<a id="routes-simple-route-advanced-options-request-cookies-to-add-secret-value-clear-secret-info-provider-ref"></a>&#x2022; [`provider_ref`](#routes-simple-route-advanced-options-request-cookies-to-add-secret-value-clear-secret-info-provider-ref) - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

<a id="routes-simple-route-advanced-options-request-cookies-to-add-secret-value-clear-secret-info-url"></a>&#x2022; [`url`](#routes-simple-route-advanced-options-request-cookies-to-add-secret-value-clear-secret-info-url) - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

#### Routes Simple Route Advanced Options Request Headers To Add

A [`request_headers_to_add`](#routes-simple-route-advanced-options-request-headers-to-add) block (within [`routes.simple_route.advanced_options`](#routes-simple-route-advanced-options)) supports the following:

<a id="routes-simple-route-advanced-options-request-headers-to-add-append"></a>&#x2022; [`append`](#routes-simple-route-advanced-options-request-headers-to-add-append) - Optional Bool  Defaults to `do`<br>Append. Should the value be appended? If true, the value is appended to existing values. not append

<a id="routes-simple-route-advanced-options-request-headers-to-add-name"></a>&#x2022; [`name`](#routes-simple-route-advanced-options-request-headers-to-add-name) - Optional String<br>Name. Name of the HTTP header

<a id="routes-simple-route-advanced-options-request-headers-to-add-secret-value"></a>&#x2022; [`secret_value`](#routes-simple-route-advanced-options-request-headers-to-add-secret-value) - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Secret Value](#routes-simple-route-advanced-options-request-headers-to-add-secret-value) below.

<a id="routes-simple-route-advanced-options-request-headers-to-add-value"></a>&#x2022; [`value`](#routes-simple-route-advanced-options-request-headers-to-add-value) - Optional String<br>Value. Value of the HTTP header

#### Routes Simple Route Advanced Options Request Headers To Add Secret Value

A [`secret_value`](#routes-simple-route-advanced-options-request-headers-to-add-secret-value) block (within [`routes.simple_route.advanced_options.request_headers_to_add`](#routes-simple-route-advanced-options-request-headers-to-add)) supports the following:

<a id="routes-simple-route-advanced-options-request-headers-to-add-secret-value-blindfold-secret-info"></a>&#x2022; [`blindfold_secret_info`](#routes-simple-route-advanced-options-request-headers-to-add-secret-value-blindfold-secret-info) - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#routes-simple-route-advanced-options-request-headers-to-add-secret-value-blindfold-secret-info) below.

<a id="routes-simple-route-advanced-options-request-headers-to-add-secret-value-clear-secret-info"></a>&#x2022; [`clear_secret_info`](#routes-simple-route-advanced-options-request-headers-to-add-secret-value-clear-secret-info) - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#routes-simple-route-advanced-options-request-headers-to-add-secret-value-clear-secret-info) below.

#### Routes Simple Route Advanced Options Request Headers To Add Secret Value Blindfold Secret Info

A [`blindfold_secret_info`](#routes-simple-route-advanced-options-request-headers-to-add-secret-value-blindfold-secret-info) block (within [`routes.simple_route.advanced_options.request_headers_to_add.secret_value`](#routes-simple-route-advanced-options-request-headers-to-add-secret-value)) supports the following:

<a id="routes-simple-route-advanced-options-request-headers-to-add-secret-value-blindfold-secret-info-decryption-provider"></a>&#x2022; [`decryption_provider`](#routes-simple-route-advanced-options-request-headers-to-add-secret-value-blindfold-secret-info-decryption-provider) - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

<a id="routes-simple-route-advanced-options-request-headers-to-add-secret-value-blindfold-secret-info-location"></a>&#x2022; [`location`](#routes-simple-route-advanced-options-request-headers-to-add-secret-value-blindfold-secret-info-location) - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

<a id="routes-simple-route-advanced-options-request-headers-to-add-secret-value-blindfold-secret-info-store-provider"></a>&#x2022; [`store_provider`](#routes-simple-route-advanced-options-request-headers-to-add-secret-value-blindfold-secret-info-store-provider) - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

#### Routes Simple Route Advanced Options Request Headers To Add Secret Value Clear Secret Info

A [`clear_secret_info`](#routes-simple-route-advanced-options-request-headers-to-add-secret-value-clear-secret-info) block (within [`routes.simple_route.advanced_options.request_headers_to_add.secret_value`](#routes-simple-route-advanced-options-request-headers-to-add-secret-value)) supports the following:

<a id="routes-simple-route-advanced-options-request-headers-to-add-secret-value-clear-secret-info-provider-ref"></a>&#x2022; [`provider_ref`](#routes-simple-route-advanced-options-request-headers-to-add-secret-value-clear-secret-info-provider-ref) - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

<a id="routes-simple-route-advanced-options-request-headers-to-add-secret-value-clear-secret-info-url"></a>&#x2022; [`url`](#routes-simple-route-advanced-options-request-headers-to-add-secret-value-clear-secret-info-url) - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

#### Routes Simple Route Advanced Options Response Cookies To Add

A [`response_cookies_to_add`](#routes-simple-route-advanced-options-response-cookies-to-add) block (within [`routes.simple_route.advanced_options`](#routes-simple-route-advanced-options)) supports the following:

<a id="routes-simple-route-advanced-options-response-cookies-to-add-add-domain"></a>&#x2022; [`add_domain`](#routes-simple-route-advanced-options-response-cookies-to-add-add-domain) - Optional String<br>Add Domain. Add domain attribute

<a id="routes-simple-route-advanced-options-response-cookies-to-add-add-expiry"></a>&#x2022; [`add_expiry`](#routes-simple-route-advanced-options-response-cookies-to-add-add-expiry) - Optional String<br>Add expiry. Add expiry attribute

<a id="routes-simple-route-advanced-options-response-cookies-to-add-add-httponly"></a>&#x2022; [`add_httponly`](#routes-simple-route-advanced-options-response-cookies-to-add-add-httponly) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="routes-simple-route-advanced-options-response-cookies-to-add-add-partitioned"></a>&#x2022; [`add_partitioned`](#routes-simple-route-advanced-options-response-cookies-to-add-add-partitioned) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="routes-simple-route-advanced-options-response-cookies-to-add-add-path"></a>&#x2022; [`add_path`](#routes-simple-route-advanced-options-response-cookies-to-add-add-path) - Optional String<br>Add path. Add path attribute

<a id="routes-simple-route-advanced-options-response-cookies-to-add-add-secure"></a>&#x2022; [`add_secure`](#routes-simple-route-advanced-options-response-cookies-to-add-add-secure) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="routes-simple-route-advanced-options-response-cookies-to-add-ignore-domain"></a>&#x2022; [`ignore_domain`](#routes-simple-route-advanced-options-response-cookies-to-add-ignore-domain) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="routes-simple-route-advanced-options-response-cookies-to-add-ignore-expiry"></a>&#x2022; [`ignore_expiry`](#routes-simple-route-advanced-options-response-cookies-to-add-ignore-expiry) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="routes-simple-route-advanced-options-response-cookies-to-add-ignore-httponly"></a>&#x2022; [`ignore_httponly`](#routes-simple-route-advanced-options-response-cookies-to-add-ignore-httponly) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="routes-simple-route-advanced-options-response-cookies-to-add-ignore-max-age"></a>&#x2022; [`ignore_max_age`](#routes-simple-route-advanced-options-response-cookies-to-add-ignore-max-age) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="routes-simple-route-advanced-options-response-cookies-to-add-ignore-partitioned"></a>&#x2022; [`ignore_partitioned`](#routes-simple-route-advanced-options-response-cookies-to-add-ignore-partitioned) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="routes-simple-route-advanced-options-response-cookies-to-add-ignore-path"></a>&#x2022; [`ignore_path`](#routes-simple-route-advanced-options-response-cookies-to-add-ignore-path) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="routes-simple-route-advanced-options-response-cookies-to-add-ignore-samesite"></a>&#x2022; [`ignore_samesite`](#routes-simple-route-advanced-options-response-cookies-to-add-ignore-samesite) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="routes-simple-route-advanced-options-response-cookies-to-add-ignore-secure"></a>&#x2022; [`ignore_secure`](#routes-simple-route-advanced-options-response-cookies-to-add-ignore-secure) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="routes-simple-route-advanced-options-response-cookies-to-add-ignore-value"></a>&#x2022; [`ignore_value`](#routes-simple-route-advanced-options-response-cookies-to-add-ignore-value) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="routes-simple-route-advanced-options-response-cookies-to-add-max-age-value"></a>&#x2022; [`max_age_value`](#routes-simple-route-advanced-options-response-cookies-to-add-max-age-value) - Optional Number<br>Add Max Age. Add max age attribute

<a id="routes-simple-route-advanced-options-response-cookies-to-add-name"></a>&#x2022; [`name`](#routes-simple-route-advanced-options-response-cookies-to-add-name) - Optional String<br>Name. Name of the cookie in Cookie header

<a id="routes-simple-route-advanced-options-response-cookies-to-add-overwrite"></a>&#x2022; [`overwrite`](#routes-simple-route-advanced-options-response-cookies-to-add-overwrite) - Optional Bool  Defaults to `do`<br>Overwrite. Should the value be overwritten? If true, the value is overwritten to existing values. not overwrite

<a id="routes-simple-route-advanced-options-response-cookies-to-add-samesite-lax"></a>&#x2022; [`samesite_lax`](#routes-simple-route-advanced-options-response-cookies-to-add-samesite-lax) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="routes-simple-route-advanced-options-response-cookies-to-add-samesite-none"></a>&#x2022; [`samesite_none`](#routes-simple-route-advanced-options-response-cookies-to-add-samesite-none) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="routes-simple-route-advanced-options-response-cookies-to-add-samesite-strict"></a>&#x2022; [`samesite_strict`](#routes-simple-route-advanced-options-response-cookies-to-add-samesite-strict) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="routes-simple-route-advanced-options-response-cookies-to-add-secret-value"></a>&#x2022; [`secret_value`](#routes-simple-route-advanced-options-response-cookies-to-add-secret-value) - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Secret Value](#routes-simple-route-advanced-options-response-cookies-to-add-secret-value) below.

<a id="routes-simple-route-advanced-options-response-cookies-to-add-value"></a>&#x2022; [`value`](#routes-simple-route-advanced-options-response-cookies-to-add-value) - Optional String<br>Value. Value of the Cookie header

#### Routes Simple Route Advanced Options Response Cookies To Add Secret Value

A [`secret_value`](#routes-simple-route-advanced-options-response-cookies-to-add-secret-value) block (within [`routes.simple_route.advanced_options.response_cookies_to_add`](#routes-simple-route-advanced-options-response-cookies-to-add)) supports the following:

<a id="routes-simple-route-advanced-options-response-cookies-to-add-secret-value-blindfold-secret-info"></a>&#x2022; [`blindfold_secret_info`](#routes-simple-route-advanced-options-response-cookies-to-add-secret-value-blindfold-secret-info) - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#routes-simple-route-advanced-options-response-cookies-to-add-secret-value-blindfold-secret-info) below.

<a id="routes-simple-route-advanced-options-response-cookies-to-add-secret-value-clear-secret-info"></a>&#x2022; [`clear_secret_info`](#routes-simple-route-advanced-options-response-cookies-to-add-secret-value-clear-secret-info) - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#routes-simple-route-advanced-options-response-cookies-to-add-secret-value-clear-secret-info) below.

#### Routes Simple Route Advanced Options Response Cookies To Add Secret Value Blindfold Secret Info

A [`blindfold_secret_info`](#routes-simple-route-advanced-options-response-cookies-to-add-secret-value-blindfold-secret-info) block (within [`routes.simple_route.advanced_options.response_cookies_to_add.secret_value`](#routes-simple-route-advanced-options-response-cookies-to-add-secret-value)) supports the following:

<a id="routes-simple-route-advanced-options-response-cookies-to-add-secret-value-blindfold-secret-info-decryption-provider"></a>&#x2022; [`decryption_provider`](#routes-simple-route-advanced-options-response-cookies-to-add-secret-value-blindfold-secret-info-decryption-provider) - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

<a id="routes-simple-route-advanced-options-response-cookies-to-add-secret-value-blindfold-secret-info-location"></a>&#x2022; [`location`](#routes-simple-route-advanced-options-response-cookies-to-add-secret-value-blindfold-secret-info-location) - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

<a id="routes-simple-route-advanced-options-response-cookies-to-add-secret-value-blindfold-secret-info-store-provider"></a>&#x2022; [`store_provider`](#routes-simple-route-advanced-options-response-cookies-to-add-secret-value-blindfold-secret-info-store-provider) - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

#### Routes Simple Route Advanced Options Response Cookies To Add Secret Value Clear Secret Info

A [`clear_secret_info`](#routes-simple-route-advanced-options-response-cookies-to-add-secret-value-clear-secret-info) block (within [`routes.simple_route.advanced_options.response_cookies_to_add.secret_value`](#routes-simple-route-advanced-options-response-cookies-to-add-secret-value)) supports the following:

<a id="routes-simple-route-advanced-options-response-cookies-to-add-secret-value-clear-secret-info-provider-ref"></a>&#x2022; [`provider_ref`](#routes-simple-route-advanced-options-response-cookies-to-add-secret-value-clear-secret-info-provider-ref) - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

<a id="routes-simple-route-advanced-options-response-cookies-to-add-secret-value-clear-secret-info-url"></a>&#x2022; [`url`](#routes-simple-route-advanced-options-response-cookies-to-add-secret-value-clear-secret-info-url) - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

#### Routes Simple Route Advanced Options Response Headers To Add

A [`response_headers_to_add`](#routes-simple-route-advanced-options-response-headers-to-add) block (within [`routes.simple_route.advanced_options`](#routes-simple-route-advanced-options)) supports the following:

<a id="routes-simple-route-advanced-options-response-headers-to-add-append"></a>&#x2022; [`append`](#routes-simple-route-advanced-options-response-headers-to-add-append) - Optional Bool  Defaults to `do`<br>Append. Should the value be appended? If true, the value is appended to existing values. not append

<a id="routes-simple-route-advanced-options-response-headers-to-add-name"></a>&#x2022; [`name`](#routes-simple-route-advanced-options-response-headers-to-add-name) - Optional String<br>Name. Name of the HTTP header

<a id="routes-simple-route-advanced-options-response-headers-to-add-secret-value"></a>&#x2022; [`secret_value`](#routes-simple-route-advanced-options-response-headers-to-add-secret-value) - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Secret Value](#routes-simple-route-advanced-options-response-headers-to-add-secret-value) below.

<a id="routes-simple-route-advanced-options-response-headers-to-add-value"></a>&#x2022; [`value`](#routes-simple-route-advanced-options-response-headers-to-add-value) - Optional String<br>Value. Value of the HTTP header

#### Routes Simple Route Advanced Options Response Headers To Add Secret Value

A [`secret_value`](#routes-simple-route-advanced-options-response-headers-to-add-secret-value) block (within [`routes.simple_route.advanced_options.response_headers_to_add`](#routes-simple-route-advanced-options-response-headers-to-add)) supports the following:

<a id="routes-simple-route-advanced-options-response-headers-to-add-secret-value-blindfold-secret-info"></a>&#x2022; [`blindfold_secret_info`](#routes-simple-route-advanced-options-response-headers-to-add-secret-value-blindfold-secret-info) - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#routes-simple-route-advanced-options-response-headers-to-add-secret-value-blindfold-secret-info) below.

<a id="routes-simple-route-advanced-options-response-headers-to-add-secret-value-clear-secret-info"></a>&#x2022; [`clear_secret_info`](#routes-simple-route-advanced-options-response-headers-to-add-secret-value-clear-secret-info) - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#routes-simple-route-advanced-options-response-headers-to-add-secret-value-clear-secret-info) below.

#### Routes Simple Route Advanced Options Response Headers To Add Secret Value Blindfold Secret Info

A [`blindfold_secret_info`](#routes-simple-route-advanced-options-response-headers-to-add-secret-value-blindfold-secret-info) block (within [`routes.simple_route.advanced_options.response_headers_to_add.secret_value`](#routes-simple-route-advanced-options-response-headers-to-add-secret-value)) supports the following:

<a id="routes-simple-route-advanced-options-response-headers-to-add-secret-value-blindfold-secret-info-decryption-provider"></a>&#x2022; [`decryption_provider`](#routes-simple-route-advanced-options-response-headers-to-add-secret-value-blindfold-secret-info-decryption-provider) - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

<a id="routes-simple-route-advanced-options-response-headers-to-add-secret-value-blindfold-secret-info-location"></a>&#x2022; [`location`](#routes-simple-route-advanced-options-response-headers-to-add-secret-value-blindfold-secret-info-location) - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

<a id="routes-simple-route-advanced-options-response-headers-to-add-secret-value-blindfold-secret-info-store-provider"></a>&#x2022; [`store_provider`](#routes-simple-route-advanced-options-response-headers-to-add-secret-value-blindfold-secret-info-store-provider) - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

#### Routes Simple Route Advanced Options Response Headers To Add Secret Value Clear Secret Info

A [`clear_secret_info`](#routes-simple-route-advanced-options-response-headers-to-add-secret-value-clear-secret-info) block (within [`routes.simple_route.advanced_options.response_headers_to_add.secret_value`](#routes-simple-route-advanced-options-response-headers-to-add-secret-value)) supports the following:

<a id="routes-simple-route-advanced-options-response-headers-to-add-secret-value-clear-secret-info-provider-ref"></a>&#x2022; [`provider_ref`](#routes-simple-route-advanced-options-response-headers-to-add-secret-value-clear-secret-info-provider-ref) - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

<a id="routes-simple-route-advanced-options-response-headers-to-add-secret-value-clear-secret-info-url"></a>&#x2022; [`url`](#routes-simple-route-advanced-options-response-headers-to-add-secret-value-clear-secret-info-url) - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

#### Routes Simple Route Advanced Options Retry Policy

A [`retry_policy`](#routes-simple-route-advanced-options-retry-policy) block (within [`routes.simple_route.advanced_options`](#routes-simple-route-advanced-options)) supports the following:

<a id="routes-simple-route-advanced-options-retry-policy-back-off"></a>&#x2022; [`back_off`](#routes-simple-route-advanced-options-retry-policy-back-off) - Optional Block<br>Retry BackOff Interval. Specifies parameters that control retry back off<br>See [Back Off](#routes-simple-route-advanced-options-retry-policy-back-off) below.

<a id="routes-simple-route-advanced-options-retry-policy-num-retries"></a>&#x2022; [`num_retries`](#routes-simple-route-advanced-options-retry-policy-num-retries) - Optional Number  Defaults to `1`<br>Number of Retries. Specifies the allowed number of retries. Retries can be done any number of times. An exponential back-off algorithm is used between each retry

<a id="routes-simple-route-advanced-options-retry-policy-per-try-timeout"></a>&#x2022; [`per_try_timeout`](#routes-simple-route-advanced-options-retry-policy-per-try-timeout) - Optional Number<br>Per Try Timeout. Specifies a non-zero timeout per retry attempt. In milliseconds

<a id="routes-simple-route-advanced-options-retry-policy-retriable-status-codes"></a>&#x2022; [`retriable_status_codes`](#routes-simple-route-advanced-options-retry-policy-retriable-status-codes) - Optional List<br>Status Code to Retry. HTTP status codes that should trigger a retry in addition to those specified by retry_on

<a id="routes-simple-route-advanced-options-retry-policy-retry-condition"></a>&#x2022; [`retry_condition`](#routes-simple-route-advanced-options-retry-policy-retry-condition) - Optional List<br>Retry Condition. Specifies the conditions under which retry takes place. Retries can be on different types of condition depending on application requirements. For example, network failure, all 5xx response codes, idempotent 4xx response codes, etc The possible values are '5xx' : Retry will be done if the upstream server responds with any 5xx response code, or does not respond at all (disconnect/reset/read timeout). 'gateway-error' : Retry will be done only if the upstream server responds with 502, 503 or 504 responses (Included in 5xx) 'connect-failure' : Retry will be done if the request fails because of a connection failure to the upstream server (connect timeout, etc.). (Included in 5xx) 'refused-stream' : Retry is done if the upstream server resets the stream with a REFUSED_STREAM error code (Included in 5xx) 'retriable-4xx' : Retry is done if the upstream server responds with a retriable 4xx response code. The only response code in this category is HTTP CONFLICT (409) 'retriable-status-codes' : Retry is done if the upstream server responds with any response code matching one defined in retriable_status_codes field 'reset' : Retry is done if the upstream server does not respond at all (disconnect/reset/read timeout.)

#### Routes Simple Route Advanced Options Retry Policy Back Off

A [`back_off`](#routes-simple-route-advanced-options-retry-policy-back-off) block (within [`routes.simple_route.advanced_options.retry_policy`](#routes-simple-route-advanced-options-retry-policy)) supports the following:

<a id="routes-simple-route-advanced-options-retry-policy-back-off-base-interval"></a>&#x2022; [`base_interval`](#routes-simple-route-advanced-options-retry-policy-back-off-base-interval) - Optional Number<br>Base Retry Interval. Specifies the base interval between retries in milliseconds

<a id="routes-simple-route-advanced-options-retry-policy-back-off-max-interval"></a>&#x2022; [`max_interval`](#routes-simple-route-advanced-options-retry-policy-back-off-max-interval) - Optional Number  Defaults to `10`<br>Maximum Retry Interval. Specifies the maximum interval between retries in milliseconds. This parameter is optional, but must be greater than or equal to the base_interval if set. The times the base_interval

#### Routes Simple Route Advanced Options Specific Hash Policy

A [`specific_hash_policy`](#routes-simple-route-advanced-options-specific-hash-policy) block (within [`routes.simple_route.advanced_options`](#routes-simple-route-advanced-options)) supports the following:

<a id="routes-simple-route-advanced-options-specific-hash-policy-hash-policy"></a>&#x2022; [`hash_policy`](#routes-simple-route-advanced-options-specific-hash-policy-hash-policy) - Optional Block<br>Hash Policy. Specifies a list of hash policies to use for ring hash load balancing. Each hash policy is evaluated individually and the combined result is used to route the request<br>See [Hash Policy](#routes-simple-route-advanced-options-specific-hash-policy-hash-policy) below.

#### Routes Simple Route Advanced Options Specific Hash Policy Hash Policy

A [`hash_policy`](#routes-simple-route-advanced-options-specific-hash-policy-hash-policy) block (within [`routes.simple_route.advanced_options.specific_hash_policy`](#routes-simple-route-advanced-options-specific-hash-policy)) supports the following:

<a id="routes-simple-route-advanced-options-specific-hash-policy-hash-policy-cookie"></a>&#x2022; [`cookie`](#routes-simple-route-advanced-options-specific-hash-policy-hash-policy-cookie) - Optional Block<br>Hashing using Cookie. Two types of cookie affinity: 1. Passive. Takes a cookie that's present in the cookies header and hashes on its value. 2. Generated. Generates and sets a cookie with an expiration (TTL) on the first request from the client in its response to the client, based on the endpoint the request gets sent to. The client then presents this on the next and all subsequent requests. The hash of this is sufficient to ensure these requests get sent to the same endpoint. The cookie is generated by hashing the source and destination ports and addresses so that multiple independent HTTP2 streams on the same connection will independently receive the same cookie, even if they arrive simultaneously<br>See [Cookie](#routes-simple-route-advanced-options-specific-hash-policy-hash-policy-cookie) below.

<a id="routes-simple-route-advanced-options-specific-hash-policy-hash-policy-header-name"></a>&#x2022; [`header_name`](#routes-simple-route-advanced-options-specific-hash-policy-hash-policy-header-name) - Optional String<br>Header Name. The name or key of the request header that will be used to obtain the hash key

<a id="routes-simple-route-advanced-options-specific-hash-policy-hash-policy-source-ip"></a>&#x2022; [`source_ip`](#routes-simple-route-advanced-options-specific-hash-policy-hash-policy-source-ip) - Optional Bool<br>Source IP. Hash based on source IP address

<a id="routes-simple-route-advanced-options-specific-hash-policy-hash-policy-terminal"></a>&#x2022; [`terminal`](#routes-simple-route-advanced-options-specific-hash-policy-hash-policy-terminal) - Optional Bool<br>Terminal. Specify if its a terminal policy

#### Routes Simple Route Advanced Options Specific Hash Policy Hash Policy Cookie

A [`cookie`](#routes-simple-route-advanced-options-specific-hash-policy-hash-policy-cookie) block (within [`routes.simple_route.advanced_options.specific_hash_policy.hash_policy`](#routes-simple-route-advanced-options-specific-hash-policy-hash-policy)) supports the following:

<a id="routes-simple-route-advanced-options-specific-hash-policy-hash-policy-cookie-add-httponly"></a>&#x2022; [`add_httponly`](#routes-simple-route-advanced-options-specific-hash-policy-hash-policy-cookie-add-httponly) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="routes-simple-route-advanced-options-specific-hash-policy-hash-policy-cookie-add-secure"></a>&#x2022; [`add_secure`](#routes-simple-route-advanced-options-specific-hash-policy-hash-policy-cookie-add-secure) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="routes-simple-route-advanced-options-specific-hash-policy-hash-policy-cookie-ignore-httponly"></a>&#x2022; [`ignore_httponly`](#routes-simple-route-advanced-options-specific-hash-policy-hash-policy-cookie-ignore-httponly) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="routes-simple-route-advanced-options-specific-hash-policy-hash-policy-cookie-ignore-samesite"></a>&#x2022; [`ignore_samesite`](#routes-simple-route-advanced-options-specific-hash-policy-hash-policy-cookie-ignore-samesite) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="routes-simple-route-advanced-options-specific-hash-policy-hash-policy-cookie-ignore-secure"></a>&#x2022; [`ignore_secure`](#routes-simple-route-advanced-options-specific-hash-policy-hash-policy-cookie-ignore-secure) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="routes-simple-route-advanced-options-specific-hash-policy-hash-policy-cookie-name"></a>&#x2022; [`name`](#routes-simple-route-advanced-options-specific-hash-policy-hash-policy-cookie-name) - Optional String<br>Name. The name of the cookie that will be used to obtain the hash key. If the cookie is not present and TTL below is not set, no hash will be produced

<a id="routes-simple-route-advanced-options-specific-hash-policy-hash-policy-cookie-path"></a>&#x2022; [`path`](#routes-simple-route-advanced-options-specific-hash-policy-hash-policy-cookie-path) - Optional String<br>Path. The name of the path for the cookie. If no path is specified here, no path will be set for the cookie

<a id="routes-simple-route-advanced-options-specific-hash-policy-hash-policy-cookie-samesite-lax"></a>&#x2022; [`samesite_lax`](#routes-simple-route-advanced-options-specific-hash-policy-hash-policy-cookie-samesite-lax) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="routes-simple-route-advanced-options-specific-hash-policy-hash-policy-cookie-samesite-none"></a>&#x2022; [`samesite_none`](#routes-simple-route-advanced-options-specific-hash-policy-hash-policy-cookie-samesite-none) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="routes-simple-route-advanced-options-specific-hash-policy-hash-policy-cookie-samesite-strict"></a>&#x2022; [`samesite_strict`](#routes-simple-route-advanced-options-specific-hash-policy-hash-policy-cookie-samesite-strict) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="routes-simple-route-advanced-options-specific-hash-policy-hash-policy-cookie-ttl"></a>&#x2022; [`ttl`](#routes-simple-route-advanced-options-specific-hash-policy-hash-policy-cookie-ttl) - Optional Number<br>TTL. If specified, a cookie with the TTL will be generated if the cookie is not present. If the TTL is present and zero, the generated cookie will be a session cookie. TTL value is in milliseconds

#### Routes Simple Route Advanced Options WAF Exclusion Policy

A [`waf_exclusion_policy`](#routes-simple-route-advanced-options-waf-exclusion-policy) block (within [`routes.simple_route.advanced_options`](#routes-simple-route-advanced-options)) supports the following:

<a id="routes-simple-route-advanced-options-waf-exclusion-policy-name"></a>&#x2022; [`name`](#routes-simple-route-advanced-options-waf-exclusion-policy-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="routes-simple-route-advanced-options-waf-exclusion-policy-namespace"></a>&#x2022; [`namespace`](#routes-simple-route-advanced-options-waf-exclusion-policy-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="routes-simple-route-advanced-options-waf-exclusion-policy-tenant"></a>&#x2022; [`tenant`](#routes-simple-route-advanced-options-waf-exclusion-policy-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Routes Simple Route Advanced Options Web Socket Config

A [`web_socket_config`](#routes-simple-route-advanced-options-web-socket-config) block (within [`routes.simple_route.advanced_options`](#routes-simple-route-advanced-options)) supports the following:

<a id="routes-simple-route-advanced-options-web-socket-config-use-websocket"></a>&#x2022; [`use_websocket`](#routes-simple-route-advanced-options-web-socket-config-use-websocket) - Optional Bool<br>Use WebSocket. Specifies that the HTTP client connection to this route is allowed to upgrade to a WebSocket connection

#### Routes Simple Route Headers

A [`headers`](#routes-simple-route-headers) block (within [`routes.simple_route`](#routes-simple-route)) supports the following:

<a id="routes-simple-route-headers-exact"></a>&#x2022; [`exact`](#routes-simple-route-headers-exact) - Optional String<br>Exact. Header value to match exactly

<a id="routes-simple-route-headers-invert-match"></a>&#x2022; [`invert_match`](#routes-simple-route-headers-invert-match) - Optional Bool<br>NOT of match. Invert the result of the match to detect missing header or non-matching value

<a id="routes-simple-route-headers-name"></a>&#x2022; [`name`](#routes-simple-route-headers-name) - Optional String<br>Name. Name of the header

<a id="routes-simple-route-headers-presence"></a>&#x2022; [`presence`](#routes-simple-route-headers-presence) - Optional Bool<br>Presence. If true, check for presence of header

<a id="routes-simple-route-headers-regex"></a>&#x2022; [`regex`](#routes-simple-route-headers-regex) - Optional String<br>Regex. Regex match of the header value in re2 format

#### Routes Simple Route Incoming Port

An [`incoming_port`](#routes-simple-route-incoming-port) block (within [`routes.simple_route`](#routes-simple-route)) supports the following:

<a id="routes-simple-route-incoming-port-no-port-match"></a>&#x2022; [`no_port_match`](#routes-simple-route-incoming-port-no-port-match) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="routes-simple-route-incoming-port-port"></a>&#x2022; [`port`](#routes-simple-route-incoming-port-port) - Optional Number<br>Port. Exact Port to match

<a id="routes-simple-route-incoming-port-port-ranges"></a>&#x2022; [`port_ranges`](#routes-simple-route-incoming-port-port-ranges) - Optional String<br>Port range. Port range to match

#### Routes Simple Route Origin Pools

An [`origin_pools`](#routes-simple-route-origin-pools) block (within [`routes.simple_route`](#routes-simple-route)) supports the following:

<a id="routes-simple-route-origin-pools-cluster"></a>&#x2022; [`cluster`](#routes-simple-route-origin-pools-cluster) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Cluster](#routes-simple-route-origin-pools-cluster) below.

<a id="routes-simple-route-origin-pools-endpoint-subsets"></a>&#x2022; [`endpoint_subsets`](#routes-simple-route-origin-pools-endpoint-subsets) - Optional Block<br>Origin Servers Subsets. Upstream origin pool may be configured to divide its origin servers into subsets based on metadata attached to the origin servers. Routes may then specify the metadata that a endpoint must match in order to be selected by the load balancer For origin servers which are discovered in K8S or Consul cluster, the label of the service is merged with endpoint's labels. In case of Consul, the label is derived from the 'Tag' field. For labels that are common between configured endpoint and discovered service, labels from discovered service takes precedence. List of key-value pairs that will be used as matching metadata. Only those origin servers of upstream origin pool which match this metadata will be selected for load balancing

<a id="routes-simple-route-origin-pools-pool"></a>&#x2022; [`pool`](#routes-simple-route-origin-pools-pool) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Pool](#routes-simple-route-origin-pools-pool) below.

<a id="routes-simple-route-origin-pools-priority"></a>&#x2022; [`priority`](#routes-simple-route-origin-pools-priority) - Optional Number<br>Priority. Priority of this origin pool, valid only with multiple origin pools. Value of 0 will make the pool as lowest priority origin pool Priority of 1 means highest priority and is considered active. When active origin pool is not available, lower priority origin pools are made active as per the increasing priority

<a id="routes-simple-route-origin-pools-weight"></a>&#x2022; [`weight`](#routes-simple-route-origin-pools-weight) - Optional Number<br>Weight. Weight of this origin pool, valid only with multiple origin pool. Value of 0 will disable the pool

#### Routes Simple Route Origin Pools Cluster

A [`cluster`](#routes-simple-route-origin-pools-cluster) block (within [`routes.simple_route.origin_pools`](#routes-simple-route-origin-pools)) supports the following:

<a id="routes-simple-route-origin-pools-cluster-name"></a>&#x2022; [`name`](#routes-simple-route-origin-pools-cluster-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="routes-simple-route-origin-pools-cluster-namespace"></a>&#x2022; [`namespace`](#routes-simple-route-origin-pools-cluster-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="routes-simple-route-origin-pools-cluster-tenant"></a>&#x2022; [`tenant`](#routes-simple-route-origin-pools-cluster-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Routes Simple Route Origin Pools Pool

A [`pool`](#routes-simple-route-origin-pools-pool) block (within [`routes.simple_route.origin_pools`](#routes-simple-route-origin-pools)) supports the following:

<a id="routes-simple-route-origin-pools-pool-name"></a>&#x2022; [`name`](#routes-simple-route-origin-pools-pool-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="routes-simple-route-origin-pools-pool-namespace"></a>&#x2022; [`namespace`](#routes-simple-route-origin-pools-pool-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="routes-simple-route-origin-pools-pool-tenant"></a>&#x2022; [`tenant`](#routes-simple-route-origin-pools-pool-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Routes Simple Route Path

A [`path`](#routes-simple-route-path) block (within [`routes.simple_route`](#routes-simple-route)) supports the following:

<a id="routes-simple-route-path-path"></a>&#x2022; [`path`](#routes-simple-route-path-path) - Optional String<br>Exact. Exact path value to match

<a id="routes-simple-route-path-prefix"></a>&#x2022; [`prefix`](#routes-simple-route-path-prefix) - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

<a id="routes-simple-route-path-regex"></a>&#x2022; [`regex`](#routes-simple-route-path-regex) - Optional String<br>Regex. Regular expression of path match (e.g. the value .* will match on all paths)

#### Routes Simple Route Query Params

A [`query_params`](#routes-simple-route-query-params) block (within [`routes.simple_route`](#routes-simple-route)) supports the following:

<a id="routes-simple-route-query-params-remove-all-params"></a>&#x2022; [`remove_all_params`](#routes-simple-route-query-params-remove-all-params) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="routes-simple-route-query-params-replace-params"></a>&#x2022; [`replace_params`](#routes-simple-route-query-params-replace-params) - Optional String<br>Replace All Parameters

<a id="routes-simple-route-query-params-retain-all-params"></a>&#x2022; [`retain_all_params`](#routes-simple-route-query-params-retain-all-params) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### Sensitive Data Disclosure Rules

A [`sensitive_data_disclosure_rules`](#sensitive-data-disclosure-rules) block supports the following:

<a id="sensitive-data-disclosure-rules-sensitive-data-types-in-response"></a>&#x2022; [`sensitive_data_types_in_response`](#sensitive-data-disclosure-rules-sensitive-data-types-in-response) - Optional Block<br>Sensitive Data Exposure Rules. Sensitive Data Exposure Rules allows specifying rules to mask sensitive data fields in API responses<br>See [Sensitive Data Types In Response](#sensitive-data-disclosure-rules-sensitive-data-types-in-response) below.

#### Sensitive Data Disclosure Rules Sensitive Data Types In Response

A [`sensitive_data_types_in_response`](#sensitive-data-disclosure-rules-sensitive-data-types-in-response) block (within [`sensitive_data_disclosure_rules`](#sensitive-data-disclosure-rules)) supports the following:

<a id="sensitive-data-disclosure-rules-sensitive-data-types-in-response-api-endpoint"></a>&#x2022; [`api_endpoint`](#sensitive-data-disclosure-rules-sensitive-data-types-in-response-api-endpoint) - Optional Block<br>API Endpoint. This defines API endpoint<br>See [API Endpoint](#sensitive-data-disclosure-rules-sensitive-data-types-in-response-api-endpoint) below.

<a id="sensitive-data-disclosure-rules-sensitive-data-types-in-response-body"></a>&#x2022; [`body`](#sensitive-data-disclosure-rules-sensitive-data-types-in-response-body) - Optional Block<br>Body Section Masking Options. Options for HTTP Body Masking<br>See [Body](#sensitive-data-disclosure-rules-sensitive-data-types-in-response-body) below.

<a id="sensitive-data-disclosure-rules-sensitive-data-types-in-response-mask"></a>&#x2022; [`mask`](#sensitive-data-disclosure-rules-sensitive-data-types-in-response-mask) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="sensitive-data-disclosure-rules-sensitive-data-types-in-response-report"></a>&#x2022; [`report`](#sensitive-data-disclosure-rules-sensitive-data-types-in-response-report) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### Sensitive Data Disclosure Rules Sensitive Data Types In Response API Endpoint

An [`api_endpoint`](#sensitive-data-disclosure-rules-sensitive-data-types-in-response-api-endpoint) block (within [`sensitive_data_disclosure_rules.sensitive_data_types_in_response`](#sensitive-data-disclosure-rules-sensitive-data-types-in-response)) supports the following:

<a id="sensitive-data-disclosure-rules-sensitive-data-types-in-response-api-endpoint-methods"></a>&#x2022; [`methods`](#sensitive-data-disclosure-rules-sensitive-data-types-in-response-api-endpoint-methods) - Optional List  Defaults to `ANY`<br>Possible values are `ANY`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`, `COPY`<br>Methods. Methods to be matched

<a id="sensitive-data-disclosure-rules-sensitive-data-types-in-response-api-endpoint-path"></a>&#x2022; [`path`](#sensitive-data-disclosure-rules-sensitive-data-types-in-response-api-endpoint-path) - Optional String<br>Path. Path to be matched

#### Sensitive Data Disclosure Rules Sensitive Data Types In Response Body

A [`body`](#sensitive-data-disclosure-rules-sensitive-data-types-in-response-body) block (within [`sensitive_data_disclosure_rules.sensitive_data_types_in_response`](#sensitive-data-disclosure-rules-sensitive-data-types-in-response)) supports the following:

<a id="sensitive-data-disclosure-rules-sensitive-data-types-in-response-body-fields"></a>&#x2022; [`fields`](#sensitive-data-disclosure-rules-sensitive-data-types-in-response-body-fields) - Optional List<br>Values. List of JSON Path field values. Use square brackets with an underscore [_] to indicate array elements (e.g., person.emails[_]). To reference JSON keys that contain spaces, enclose the entire path in double quotes. For example: 'person.first name'

#### Sensitive Data Policy

A [`sensitive_data_policy`](#sensitive-data-policy) block supports the following:

<a id="sensitive-data-policy-sensitive-data-policy-ref"></a>&#x2022; [`sensitive_data_policy_ref`](#sensitive-data-policy-sensitive-data-policy-ref) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Sensitive Data Policy Ref](#sensitive-data-policy-sensitive-data-policy-ref) below.

#### Sensitive Data Policy Sensitive Data Policy Ref

A [`sensitive_data_policy_ref`](#sensitive-data-policy-sensitive-data-policy-ref) block (within [`sensitive_data_policy`](#sensitive-data-policy)) supports the following:

<a id="sensitive-data-policy-sensitive-data-policy-ref-name"></a>&#x2022; [`name`](#sensitive-data-policy-sensitive-data-policy-ref-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="sensitive-data-policy-sensitive-data-policy-ref-namespace"></a>&#x2022; [`namespace`](#sensitive-data-policy-sensitive-data-policy-ref-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="sensitive-data-policy-sensitive-data-policy-ref-tenant"></a>&#x2022; [`tenant`](#sensitive-data-policy-sensitive-data-policy-ref-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Single LB App

A [`single_lb_app`](#single-lb-app) block supports the following:

<a id="single-lb-app-disable-discovery"></a>&#x2022; [`disable_discovery`](#single-lb-app-disable-discovery) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="single-lb-app-disable-malicious-user-detection"></a>&#x2022; [`disable_malicious_user_detection`](#single-lb-app-disable-malicious-user-detection) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="single-lb-app-enable-discovery"></a>&#x2022; [`enable_discovery`](#single-lb-app-enable-discovery) - Optional Block<br>API Discovery Setting. Specifies the settings used for API discovery<br>See [Enable Discovery](#single-lb-app-enable-discovery) below.

<a id="single-lb-app-enable-malicious-user-detection"></a>&#x2022; [`enable_malicious_user_detection`](#single-lb-app-enable-malicious-user-detection) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### Single LB App Enable Discovery

An [`enable_discovery`](#single-lb-app-enable-discovery) block (within [`single_lb_app`](#single-lb-app)) supports the following:

<a id="single-lb-app-enable-discovery-api-crawler"></a>&#x2022; [`api_crawler`](#single-lb-app-enable-discovery-api-crawler) - Optional Block<br>API Crawling. API Crawler message<br>See [API Crawler](#single-lb-app-enable-discovery-api-crawler) below.

<a id="single-lb-app-enable-discovery-api-discovery-from-code-scan"></a>&#x2022; [`api_discovery_from_code_scan`](#single-lb-app-enable-discovery-api-discovery-from-code-scan) - Optional Block<br>Select Code Base and Repositories. x-required<br>See [API Discovery From Code Scan](#single-lb-app-enable-discovery-api-discovery-from-code-scan) below.

<a id="single-lb-app-enable-discovery-custom-api-auth-discovery"></a>&#x2022; [`custom_api_auth_discovery`](#single-lb-app-enable-discovery-custom-api-auth-discovery) - Optional Block<br>API Discovery Advanced Settings. API Discovery Advanced settings<br>See [Custom API Auth Discovery](#single-lb-app-enable-discovery-custom-api-auth-discovery) below.

<a id="single-lb-app-enable-discovery-default-api-auth-discovery"></a>&#x2022; [`default_api_auth_discovery`](#single-lb-app-enable-discovery-default-api-auth-discovery) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="single-lb-app-enable-discovery-disable-learn-from-redirect-traffic"></a>&#x2022; [`disable_learn_from_redirect_traffic`](#single-lb-app-enable-discovery-disable-learn-from-redirect-traffic) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="single-lb-app-enable-discovery-discovered-api-settings"></a>&#x2022; [`discovered_api_settings`](#single-lb-app-enable-discovery-discovered-api-settings) - Optional Block<br>Discovered API Settings. x-example: '2' Configure Discovered API Settings<br>See [Discovered API Settings](#single-lb-app-enable-discovery-discovered-api-settings) below.

<a id="single-lb-app-enable-discovery-enable-learn-from-redirect-traffic"></a>&#x2022; [`enable_learn_from_redirect_traffic`](#single-lb-app-enable-discovery-enable-learn-from-redirect-traffic) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### Single LB App Enable Discovery API Crawler

An [`api_crawler`](#single-lb-app-enable-discovery-api-crawler) block (within [`single_lb_app.enable_discovery`](#single-lb-app-enable-discovery)) supports the following:

<a id="single-lb-app-enable-discovery-api-crawler-api-crawler-config"></a>&#x2022; [`api_crawler_config`](#single-lb-app-enable-discovery-api-crawler-api-crawler-config) - Optional Block<br>Crawler Configure<br>See [API Crawler Config](#single-lb-app-enable-discovery-api-crawler-api-crawler-config) below.

<a id="single-lb-app-enable-discovery-api-crawler-disable-api-crawler"></a>&#x2022; [`disable_api_crawler`](#single-lb-app-enable-discovery-api-crawler-disable-api-crawler) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### Single LB App Enable Discovery API Crawler API Crawler Config

An [`api_crawler_config`](#single-lb-app-enable-discovery-api-crawler-api-crawler-config) block (within [`single_lb_app.enable_discovery.api_crawler`](#single-lb-app-enable-discovery-api-crawler)) supports the following:

<a id="single-lb-app-enable-discovery-api-crawler-api-crawler-config-domains"></a>&#x2022; [`domains`](#single-lb-app-enable-discovery-api-crawler-api-crawler-config-domains) - Optional Block<br>Domains to Crawl. Enter domains and their credentials to allow authenticated API crawling. You can only include domains you own that are associated with this Load Balancer<br>See [Domains](#single-lb-app-enable-discovery-api-crawler-api-crawler-config-domains) below.

#### Single LB App Enable Discovery API Crawler API Crawler Config Domains

A [`domains`](#single-lb-app-enable-discovery-api-crawler-api-crawler-config-domains) block (within [`single_lb_app.enable_discovery.api_crawler.api_crawler_config`](#single-lb-app-enable-discovery-api-crawler-api-crawler-config)) supports the following:

<a id="single-lb-app-enable-discovery-api-crawler-api-crawler-config-domains-domain"></a>&#x2022; [`domain`](#single-lb-app-enable-discovery-api-crawler-api-crawler-config-domains-domain) - Optional String<br>Domain. Select the domain to execute API Crawling with given credentials

<a id="single-lb-app-enable-discovery-api-crawler-api-crawler-config-domains-simple-login"></a>&#x2022; [`simple_login`](#single-lb-app-enable-discovery-api-crawler-api-crawler-config-domains-simple-login) - Optional Block<br>Simple Login<br>See [Simple Login](#single-lb-app-enable-discovery-api-crawler-api-crawler-config-domains-simple-login) below.

#### Single LB App Enable Discovery API Crawler API Crawler Config Domains Simple Login

A [`simple_login`](#single-lb-app-enable-discovery-api-crawler-api-crawler-config-domains-simple-login) block (within [`single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains`](#single-lb-app-enable-discovery-api-crawler-api-crawler-config-domains)) supports the following:

<a id="single-lb-app-enable-discovery-api-crawler-api-crawler-config-domains-simple-login-password"></a>&#x2022; [`password`](#single-lb-app-enable-discovery-api-crawler-api-crawler-config-domains-simple-login-password) - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Password](#single-lb-app-enable-discovery-api-crawler-api-crawler-config-domains-simple-login-password) below.

<a id="single-lb-app-enable-discovery-api-crawler-api-crawler-config-domains-simple-login-user"></a>&#x2022; [`user`](#single-lb-app-enable-discovery-api-crawler-api-crawler-config-domains-simple-login-user) - Optional String<br>User. Enter the username to assign credentials for the selected domain to crawl

#### Single LB App Enable Discovery API Crawler API Crawler Config Domains Simple Login Password

A [`password`](#single-lb-app-enable-discovery-api-crawler-api-crawler-config-domains-simple-login-password) block (within [`single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login`](#single-lb-app-enable-discovery-api-crawler-api-crawler-config-domains-simple-login)) supports the following:

<a id="single-lb-app-enable-discovery-api-crawler-api-crawler-config-domains-simple-login-password-blindfold-secret-info"></a>&#x2022; [`blindfold_secret_info`](#single-lb-app-enable-discovery-api-crawler-api-crawler-config-domains-simple-login-password-blindfold-secret-info) - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#single-lb-app-enable-discovery-api-crawler-api-crawler-config-domains-simple-login-password-blindfold-secret-info) below.

<a id="single-lb-app-enable-discovery-api-crawler-api-crawler-config-domains-simple-login-password-clear-secret-info"></a>&#x2022; [`clear_secret_info`](#single-lb-app-enable-discovery-api-crawler-api-crawler-config-domains-simple-login-password-clear-secret-info) - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#single-lb-app-enable-discovery-api-crawler-api-crawler-config-domains-simple-login-password-clear-secret-info) below.

#### Single LB App Enable Discovery API Crawler API Crawler Config Domains Simple Login Password Blindfold Secret Info

A [`blindfold_secret_info`](#single-lb-app-enable-discovery-api-crawler-api-crawler-config-domains-simple-login-password-blindfold-secret-info) block (within [`single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password`](#single-lb-app-enable-discovery-api-crawler-api-crawler-config-domains-simple-login-password)) supports the following:

<a id="single-lb-app-enable-discovery-api-crawler-api-crawler-config-domains-simple-login-password-blindfold-secret-info-decryption-provider"></a>&#x2022; [`decryption_provider`](#single-lb-app-enable-discovery-api-crawler-api-crawler-config-domains-simple-login-password-blindfold-secret-info-decryption-provider) - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

<a id="single-lb-app-enable-discovery-api-crawler-api-crawler-config-domains-simple-login-password-blindfold-secret-info-location"></a>&#x2022; [`location`](#single-lb-app-enable-discovery-api-crawler-api-crawler-config-domains-simple-login-password-blindfold-secret-info-location) - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

<a id="single-lb-app-enable-discovery-api-crawler-api-crawler-config-domains-simple-login-password-blindfold-secret-info-store-provider"></a>&#x2022; [`store_provider`](#single-lb-app-enable-discovery-api-crawler-api-crawler-config-domains-simple-login-password-blindfold-secret-info-store-provider) - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

#### Single LB App Enable Discovery API Crawler API Crawler Config Domains Simple Login Password Clear Secret Info

A [`clear_secret_info`](#single-lb-app-enable-discovery-api-crawler-api-crawler-config-domains-simple-login-password-clear-secret-info) block (within [`single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password`](#single-lb-app-enable-discovery-api-crawler-api-crawler-config-domains-simple-login-password)) supports the following:

<a id="single-lb-app-enable-discovery-api-crawler-api-crawler-config-domains-simple-login-password-clear-secret-info-provider-ref"></a>&#x2022; [`provider_ref`](#single-lb-app-enable-discovery-api-crawler-api-crawler-config-domains-simple-login-password-clear-secret-info-provider-ref) - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

<a id="single-lb-app-enable-discovery-api-crawler-api-crawler-config-domains-simple-login-password-clear-secret-info-url"></a>&#x2022; [`url`](#single-lb-app-enable-discovery-api-crawler-api-crawler-config-domains-simple-login-password-clear-secret-info-url) - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

#### Single LB App Enable Discovery API Discovery From Code Scan

An [`api_discovery_from_code_scan`](#single-lb-app-enable-discovery-api-discovery-from-code-scan) block (within [`single_lb_app.enable_discovery`](#single-lb-app-enable-discovery)) supports the following:

<a id="single-lb-app-enable-discovery-api-discovery-from-code-scan-code-base-integrations"></a>&#x2022; [`code_base_integrations`](#single-lb-app-enable-discovery-api-discovery-from-code-scan-code-base-integrations) - Optional Block<br>Select Code Base Integrations<br>See [Code Base Integrations](#single-lb-app-enable-discovery-api-discovery-from-code-scan-code-base-integrations) below.

#### Single LB App Enable Discovery API Discovery From Code Scan Code Base Integrations

A [`code_base_integrations`](#single-lb-app-enable-discovery-api-discovery-from-code-scan-code-base-integrations) block (within [`single_lb_app.enable_discovery.api_discovery_from_code_scan`](#single-lb-app-enable-discovery-api-discovery-from-code-scan)) supports the following:

<a id="single-lb-app-enable-discovery-api-discovery-from-code-scan-code-base-integrations-all-repos"></a>&#x2022; [`all_repos`](#single-lb-app-enable-discovery-api-discovery-from-code-scan-code-base-integrations-all-repos) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="single-lb-app-enable-discovery-api-discovery-from-code-scan-code-base-integrations-code-base-integration"></a>&#x2022; [`code_base_integration`](#single-lb-app-enable-discovery-api-discovery-from-code-scan-code-base-integrations-code-base-integration) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Code Base Integration](#single-lb-app-enable-discovery-api-discovery-from-code-scan-code-base-integrations-code-base-integration) below.

<a id="single-lb-app-enable-discovery-api-discovery-from-code-scan-code-base-integrations-selected-repos"></a>&#x2022; [`selected_repos`](#single-lb-app-enable-discovery-api-discovery-from-code-scan-code-base-integrations-selected-repos) - Optional Block<br>API Code Repositories. Select which API repositories represent the LB applications<br>See [Selected Repos](#single-lb-app-enable-discovery-api-discovery-from-code-scan-code-base-integrations-selected-repos) below.

#### Single LB App Enable Discovery API Discovery From Code Scan Code Base Integrations Code Base Integration

A [`code_base_integration`](#single-lb-app-enable-discovery-api-discovery-from-code-scan-code-base-integrations-code-base-integration) block (within [`single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations`](#single-lb-app-enable-discovery-api-discovery-from-code-scan-code-base-integrations)) supports the following:

<a id="single-lb-app-enable-discovery-api-discovery-from-code-scan-code-base-integrations-code-base-integration-name"></a>&#x2022; [`name`](#single-lb-app-enable-discovery-api-discovery-from-code-scan-code-base-integrations-code-base-integration-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="single-lb-app-enable-discovery-api-discovery-from-code-scan-code-base-integrations-code-base-integration-namespace"></a>&#x2022; [`namespace`](#single-lb-app-enable-discovery-api-discovery-from-code-scan-code-base-integrations-code-base-integration-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="single-lb-app-enable-discovery-api-discovery-from-code-scan-code-base-integrations-code-base-integration-tenant"></a>&#x2022; [`tenant`](#single-lb-app-enable-discovery-api-discovery-from-code-scan-code-base-integrations-code-base-integration-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Single LB App Enable Discovery API Discovery From Code Scan Code Base Integrations Selected Repos

A [`selected_repos`](#single-lb-app-enable-discovery-api-discovery-from-code-scan-code-base-integrations-selected-repos) block (within [`single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations`](#single-lb-app-enable-discovery-api-discovery-from-code-scan-code-base-integrations)) supports the following:

<a id="single-lb-app-enable-discovery-api-discovery-from-code-scan-code-base-integrations-selected-repos-api-code-repo"></a>&#x2022; [`api_code_repo`](#single-lb-app-enable-discovery-api-discovery-from-code-scan-code-base-integrations-selected-repos-api-code-repo) - Optional List<br>API Code Repository. Code repository which contain API endpoints

#### Single LB App Enable Discovery Custom API Auth Discovery

A [`custom_api_auth_discovery`](#single-lb-app-enable-discovery-custom-api-auth-discovery) block (within [`single_lb_app.enable_discovery`](#single-lb-app-enable-discovery)) supports the following:

<a id="single-lb-app-enable-discovery-custom-api-auth-discovery-api-discovery-ref"></a>&#x2022; [`api_discovery_ref`](#single-lb-app-enable-discovery-custom-api-auth-discovery-api-discovery-ref) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [API Discovery Ref](#single-lb-app-enable-discovery-custom-api-auth-discovery-api-discovery-ref) below.

#### Single LB App Enable Discovery Custom API Auth Discovery API Discovery Ref

An [`api_discovery_ref`](#single-lb-app-enable-discovery-custom-api-auth-discovery-api-discovery-ref) block (within [`single_lb_app.enable_discovery.custom_api_auth_discovery`](#single-lb-app-enable-discovery-custom-api-auth-discovery)) supports the following:

<a id="single-lb-app-enable-discovery-custom-api-auth-discovery-api-discovery-ref-name"></a>&#x2022; [`name`](#single-lb-app-enable-discovery-custom-api-auth-discovery-api-discovery-ref-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="single-lb-app-enable-discovery-custom-api-auth-discovery-api-discovery-ref-namespace"></a>&#x2022; [`namespace`](#single-lb-app-enable-discovery-custom-api-auth-discovery-api-discovery-ref-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="single-lb-app-enable-discovery-custom-api-auth-discovery-api-discovery-ref-tenant"></a>&#x2022; [`tenant`](#single-lb-app-enable-discovery-custom-api-auth-discovery-api-discovery-ref-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### Single LB App Enable Discovery Discovered API Settings

A [`discovered_api_settings`](#single-lb-app-enable-discovery-discovered-api-settings) block (within [`single_lb_app.enable_discovery`](#single-lb-app-enable-discovery)) supports the following:

<a id="single-lb-app-enable-discovery-discovered-api-settings-purge-duration-for-inactive-discovered-apis"></a>&#x2022; [`purge_duration_for_inactive_discovered_apis`](#single-lb-app-enable-discovery-discovered-api-settings-purge-duration-for-inactive-discovered-apis) - Optional Number<br>Purge Duration for Inactive Discovered APIs from Traffic. Inactive discovered API will be deleted after configured duration

#### Slow DDOS Mitigation

A [`slow_ddos_mitigation`](#slow-ddos-mitigation) block supports the following:

<a id="slow-ddos-mitigation-disable-request-timeout"></a>&#x2022; [`disable_request_timeout`](#slow-ddos-mitigation-disable-request-timeout) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="slow-ddos-mitigation-request-headers-timeout"></a>&#x2022; [`request_headers_timeout`](#slow-ddos-mitigation-request-headers-timeout) - Optional Number  Defaults to `10000`<br>Request Headers Timeout. The amount of time the client has to send only the headers on the request stream before the stream is cancelled. The milliseconds. This setting provides protection against Slowloris attacks

<a id="slow-ddos-mitigation-request-timeout"></a>&#x2022; [`request_timeout`](#slow-ddos-mitigation-request-timeout) - Optional Number<br>Custom Timeout

#### Timeouts

A [`timeouts`](#timeouts) block supports the following:

<a id="timeouts-create"></a>&#x2022; [`create`](#timeouts-create) - Optional String (Defaults to `10 minutes`)<br>Used when creating the resource

<a id="timeouts-delete"></a>&#x2022; [`delete`](#timeouts-delete) - Optional String (Defaults to `10 minutes`)<br>Used when deleting the resource

<a id="timeouts-read"></a>&#x2022; [`read`](#timeouts-read) - Optional String (Defaults to `5 minutes`)<br>Used when retrieving the resource

<a id="timeouts-update"></a>&#x2022; [`update`](#timeouts-update) - Optional String (Defaults to `10 minutes`)<br>Used when updating the resource

#### Trusted Clients

A [`trusted_clients`](#trusted-clients) block supports the following:

<a id="trusted-clients-actions"></a>&#x2022; [`actions`](#trusted-clients-actions) - Optional List  Defaults to `SKIP_PROCESSING_WAF`<br>Possible values are `SKIP_PROCESSING_WAF`, `SKIP_PROCESSING_BOT`, `SKIP_PROCESSING_MUM`, `SKIP_PROCESSING_IP_REPUTATION`, `SKIP_PROCESSING_API_PROTECTION`, `SKIP_PROCESSING_OAS_VALIDATION`, `SKIP_PROCESSING_DDOS_PROTECTION`, `SKIP_PROCESSING_THREAT_MESH`, `SKIP_PROCESSING_MALWARE_PROTECTION`<br>Actions. Actions that should be taken when client identifier matches the rule

<a id="trusted-clients-as-number"></a>&#x2022; [`as_number`](#trusted-clients-as-number) - Optional Number<br>AS Number. RFC 6793 defined 4-byte AS number

<a id="trusted-clients-bot-skip-processing"></a>&#x2022; [`bot_skip_processing`](#trusted-clients-bot-skip-processing) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="trusted-clients-expiration-timestamp"></a>&#x2022; [`expiration_timestamp`](#trusted-clients-expiration-timestamp) - Optional String<br>Expiration Timestamp. The expiration_timestamp is the RFC 3339 format timestamp at which the containing rule is considered to be logically expired. The rule continues to exist in the configuration but is not applied anymore

<a id="trusted-clients-http-header"></a>&#x2022; [`http_header`](#trusted-clients-http-header) - Optional Block<br>HTTP Header. Request header name and value pairs<br>See [HTTP Header](#trusted-clients-http-header) below.

<a id="trusted-clients-ip-prefix"></a>&#x2022; [`ip_prefix`](#trusted-clients-ip-prefix) - Optional String<br>IPv4 Prefix. IPv4 prefix string

<a id="trusted-clients-ipv6-prefix"></a>&#x2022; [`ipv6_prefix`](#trusted-clients-ipv6-prefix) - Optional String<br>IPv6 Prefix. IPv6 prefix string

<a id="trusted-clients-metadata"></a>&#x2022; [`metadata`](#trusted-clients-metadata) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#trusted-clients-metadata) below.

<a id="trusted-clients-skip-processing"></a>&#x2022; [`skip_processing`](#trusted-clients-skip-processing) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="trusted-clients-user-identifier"></a>&#x2022; [`user_identifier`](#trusted-clients-user-identifier) - Optional String<br>User Identifier. Identify user based on user identifier. User identifier value needs to be copied from security event

<a id="trusted-clients-waf-skip-processing"></a>&#x2022; [`waf_skip_processing`](#trusted-clients-waf-skip-processing) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### Trusted Clients HTTP Header

A [`http_header`](#trusted-clients-http-header) block (within [`trusted_clients`](#trusted-clients)) supports the following:

<a id="trusted-clients-http-header-headers"></a>&#x2022; [`headers`](#trusted-clients-http-header-headers) - Optional Block<br>HTTP Headers. List of HTTP header name and value pairs<br>See [Headers](#trusted-clients-http-header-headers) below.

#### Trusted Clients HTTP Header Headers

A [`headers`](#trusted-clients-http-header-headers) block (within [`trusted_clients.http_header`](#trusted-clients-http-header)) supports the following:

<a id="trusted-clients-http-header-headers-exact"></a>&#x2022; [`exact`](#trusted-clients-http-header-headers-exact) - Optional String<br>Exact. Header value to match exactly

<a id="trusted-clients-http-header-headers-invert-match"></a>&#x2022; [`invert_match`](#trusted-clients-http-header-headers-invert-match) - Optional Bool<br>NOT of match. Invert the result of the match to detect missing header or non-matching value

<a id="trusted-clients-http-header-headers-name"></a>&#x2022; [`name`](#trusted-clients-http-header-headers-name) - Optional String<br>Name. Name of the header

<a id="trusted-clients-http-header-headers-presence"></a>&#x2022; [`presence`](#trusted-clients-http-header-headers-presence) - Optional Bool<br>Presence. If true, check for presence of header

<a id="trusted-clients-http-header-headers-regex"></a>&#x2022; [`regex`](#trusted-clients-http-header-headers-regex) - Optional String<br>Regex. Regex match of the header value in re2 format

#### Trusted Clients Metadata

A [`metadata`](#trusted-clients-metadata) block (within [`trusted_clients`](#trusted-clients)) supports the following:

<a id="trusted-clients-metadata-description-spec"></a>&#x2022; [`description_spec`](#trusted-clients-metadata-description-spec) - Optional String<br>Description. Human readable description

<a id="trusted-clients-metadata-name"></a>&#x2022; [`name`](#trusted-clients-metadata-name) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### User Identification

An [`user_identification`](#user-identification) block supports the following:

<a id="user-identification-name"></a>&#x2022; [`name`](#user-identification-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="user-identification-namespace"></a>&#x2022; [`namespace`](#user-identification-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="user-identification-tenant"></a>&#x2022; [`tenant`](#user-identification-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

#### WAF Exclusion

A [`waf_exclusion`](#waf-exclusion) block supports the following:

<a id="waf-exclusion-waf-exclusion-inline-rules"></a>&#x2022; [`waf_exclusion_inline_rules`](#waf-exclusion-waf-exclusion-inline-rules) - Optional Block<br>WAF Exclusion Inline Rules. A list of WAF exclusion rules that will be applied inline<br>See [WAF Exclusion Inline Rules](#waf-exclusion-waf-exclusion-inline-rules) below.

<a id="waf-exclusion-waf-exclusion-policy"></a>&#x2022; [`waf_exclusion_policy`](#waf-exclusion-waf-exclusion-policy) - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [WAF Exclusion Policy](#waf-exclusion-waf-exclusion-policy) below.

#### WAF Exclusion WAF Exclusion Inline Rules

A [`waf_exclusion_inline_rules`](#waf-exclusion-waf-exclusion-inline-rules) block (within [`waf_exclusion`](#waf-exclusion)) supports the following:

<a id="waf-exclusion-waf-exclusion-inline-rules-rules"></a>&#x2022; [`rules`](#waf-exclusion-waf-exclusion-inline-rules-rules) - Optional Block<br>WAF Exclusion Rules. An ordered list of WAF Exclusions specific to this Load Balancer<br>See [Rules](#waf-exclusion-waf-exclusion-inline-rules-rules) below.

#### WAF Exclusion WAF Exclusion Inline Rules Rules

A [`rules`](#waf-exclusion-waf-exclusion-inline-rules-rules) block (within [`waf_exclusion.waf_exclusion_inline_rules`](#waf-exclusion-waf-exclusion-inline-rules)) supports the following:

<a id="waf-exclusion-waf-exclusion-inline-rules-rules-any-domain"></a>&#x2022; [`any_domain`](#waf-exclusion-waf-exclusion-inline-rules-rules-any-domain) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="waf-exclusion-waf-exclusion-inline-rules-rules-any-path"></a>&#x2022; [`any_path`](#waf-exclusion-waf-exclusion-inline-rules-rules-any-path) - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control"></a>&#x2022; [`app_firewall_detection_control`](#waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control) - Optional Block<br>App Firewall Detection Control. Define the list of Signature IDs, Violations, Attack Types and Bot Names that should be excluded from triggering on the defined match criteria<br>See [App Firewall Detection Control](#waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control) below.

<a id="waf-exclusion-waf-exclusion-inline-rules-rules-exact-value"></a>&#x2022; [`exact_value`](#waf-exclusion-waf-exclusion-inline-rules-rules-exact-value) - Optional String<br>Exact Value. Exact domain name

<a id="waf-exclusion-waf-exclusion-inline-rules-rules-expiration-timestamp"></a>&#x2022; [`expiration_timestamp`](#waf-exclusion-waf-exclusion-inline-rules-rules-expiration-timestamp) - Optional String<br>Expiration Timestamp. The expiration_timestamp is the RFC 3339 format timestamp at which the containing rule is considered to be logically expired. The rule continues to exist in the configuration but is not applied anymore

<a id="waf-exclusion-waf-exclusion-inline-rules-rules-metadata"></a>&#x2022; [`metadata`](#waf-exclusion-waf-exclusion-inline-rules-rules-metadata) - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#waf-exclusion-waf-exclusion-inline-rules-rules-metadata) below.

<a id="waf-exclusion-waf-exclusion-inline-rules-rules-methods"></a>&#x2022; [`methods`](#waf-exclusion-waf-exclusion-inline-rules-rules-methods) - Optional List  Defaults to `ANY`<br>Possible values are `ANY`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`, `COPY`<br>Methods. methods to be matched

<a id="waf-exclusion-waf-exclusion-inline-rules-rules-path-prefix"></a>&#x2022; [`path_prefix`](#waf-exclusion-waf-exclusion-inline-rules-rules-path-prefix) - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

<a id="waf-exclusion-waf-exclusion-inline-rules-rules-path-regex"></a>&#x2022; [`path_regex`](#waf-exclusion-waf-exclusion-inline-rules-rules-path-regex) - Optional String<br>Path Regex. Define the regex for the path. For example, the regex ^/.*$ will match on all paths

<a id="waf-exclusion-waf-exclusion-inline-rules-rules-suffix-value"></a>&#x2022; [`suffix_value`](#waf-exclusion-waf-exclusion-inline-rules-rules-suffix-value) - Optional String<br>Suffix Value. Suffix of domain name e.g 'xyz.com' will match '*.xyz.com' and 'xyz.com'

<a id="waf-exclusion-waf-exclusion-inline-rules-rules-waf-skip-processing"></a>&#x2022; [`waf_skip_processing`](#waf-exclusion-waf-exclusion-inline-rules-rules-waf-skip-processing) - Optional Block<br>Empty. This can be used for messages where no values are needed

#### WAF Exclusion WAF Exclusion Inline Rules Rules App Firewall Detection Control

An [`app_firewall_detection_control`](#waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control) block (within [`waf_exclusion.waf_exclusion_inline_rules.rules`](#waf-exclusion-waf-exclusion-inline-rules-rules)) supports the following:

<a id="waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control-exclude-attack-type-contexts"></a>&#x2022; [`exclude_attack_type_contexts`](#waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control-exclude-attack-type-contexts) - Optional Block<br>Attack Types. Attack Types to be excluded for the defined match criteria<br>See [Exclude Attack Type Contexts](#waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control-exclude-attack-type-contexts) below.

<a id="waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control-exclude-bot-name-contexts"></a>&#x2022; [`exclude_bot_name_contexts`](#waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control-exclude-bot-name-contexts) - Optional Block<br>Bot Names. Bot Names to be excluded for the defined match criteria<br>See [Exclude Bot Name Contexts](#waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control-exclude-bot-name-contexts) below.

<a id="waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control-exclude-signature-contexts"></a>&#x2022; [`exclude_signature_contexts`](#waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control-exclude-signature-contexts) - Optional Block<br>Signature IDs. Signature IDs to be excluded for the defined match criteria<br>See [Exclude Signature Contexts](#waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control-exclude-signature-contexts) below.

<a id="waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control-exclude-violation-contexts"></a>&#x2022; [`exclude_violation_contexts`](#waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control-exclude-violation-contexts) - Optional Block<br>Violations. Violations to be excluded for the defined match criteria<br>See [Exclude Violation Contexts](#waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control-exclude-violation-contexts) below.

#### WAF Exclusion WAF Exclusion Inline Rules Rules App Firewall Detection Control Exclude Attack Type Contexts

An [`exclude_attack_type_contexts`](#waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control-exclude-attack-type-contexts) block (within [`waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control`](#waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control)) supports the following:

<a id="waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control-exclude-attack-type-contexts-context"></a>&#x2022; [`context`](#waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control-exclude-attack-type-contexts-context) - Optional String  Defaults to `CONTEXT_ANY`<br>Possible values are `CONTEXT_ANY`, `CONTEXT_BODY`, `CONTEXT_REQUEST`, `CONTEXT_RESPONSE`, `CONTEXT_PARAMETER`, `CONTEXT_HEADER`, `CONTEXT_COOKIE`, `CONTEXT_URL`, `CONTEXT_URI`<br>WAF Exclusion Context Options. The available contexts for Exclusion rules. - CONTEXT_ANY: CONTEXT_ANY Detection will be excluded for all contexts. - CONTEXT_BODY: CONTEXT_BODY Detection will be excluded for the request body. - CONTEXT_REQUEST: CONTEXT_REQUEST Detection will be excluded for the request. - CONTEXT_RESPONSE: CONTEXT_RESPONSE - CONTEXT_PARAMETER: CONTEXT_PARAMETER Detection will be excluded for the parameters. The parameter name is required in the Context name field. If the field is left empty, the detection will be excluded for all parameters. - CONTEXT_HEADER: CONTEXT_HEADER Detection will be excluded for the headers. The header name is required in the Context name field. If the field is left empty, the detection will be excluded for all headers. - CONTEXT_COOKIE: CONTEXT_COOKIE Detection will be excluded for the cookies. The cookie name is required in the Context name field. If the field is left empty, the detection will be excluded for all cookies. - CONTEXT_URL: CONTEXT_URL Detection will be excluded for the request URL. - CONTEXT_URI: CONTEXT_URI

<a id="waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control-exclude-attack-type-contexts-context-name"></a>&#x2022; [`context_name`](#waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control-exclude-attack-type-contexts-context-name) - Optional String<br>Context Name. Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an wildcard asterisk (*)

<a id="waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control-exclude-attack-type-contexts-exclude-attack-type"></a>&#x2022; [`exclude_attack_type`](#waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control-exclude-attack-type-contexts-exclude-attack-type) - Optional String  Defaults to `ATTACK_TYPE_NONE`<br>Possible values are `ATTACK_TYPE_NONE`, `ATTACK_TYPE_NON_BROWSER_CLIENT`, `ATTACK_TYPE_OTHER_APPLICATION_ATTACKS`, `ATTACK_TYPE_TROJAN_BACKDOOR_SPYWARE`, `ATTACK_TYPE_DETECTION_EVASION`, `ATTACK_TYPE_VULNERABILITY_SCAN`, `ATTACK_TYPE_ABUSE_OF_FUNCTIONALITY`, `ATTACK_TYPE_AUTHENTICATION_AUTHORIZATION_ATTACKS`, `ATTACK_TYPE_BUFFER_OVERFLOW`, `ATTACK_TYPE_PREDICTABLE_RESOURCE_LOCATION`, `ATTACK_TYPE_INFORMATION_LEAKAGE`, `ATTACK_TYPE_DIRECTORY_INDEXING`, `ATTACK_TYPE_PATH_TRAVERSAL`, `ATTACK_TYPE_XPATH_INJECTION`, `ATTACK_TYPE_LDAP_INJECTION`, `ATTACK_TYPE_SERVER_SIDE_CODE_INJECTION`, `ATTACK_TYPE_COMMAND_EXECUTION`, `ATTACK_TYPE_SQL_INJECTION`, `ATTACK_TYPE_CROSS_SITE_SCRIPTING`, `ATTACK_TYPE_DENIAL_OF_SERVICE`, `ATTACK_TYPE_HTTP_PARSER_ATTACK`, `ATTACK_TYPE_SESSION_HIJACKING`, `ATTACK_TYPE_HTTP_RESPONSE_SPLITTING`, `ATTACK_TYPE_FORCEFUL_BROWSING`, `ATTACK_TYPE_REMOTE_FILE_INCLUDE`, `ATTACK_TYPE_MALICIOUS_FILE_UPLOAD`, `ATTACK_TYPE_GRAPHQL_PARSER_ATTACK`<br>Attack Types. List of all Attack Types ATTACK_TYPE_NONE ATTACK_TYPE_NON_BROWSER_CLIENT ATTACK_TYPE_OTHER_APPLICATION_ATTACKS ATTACK_TYPE_TROJAN_BACKDOOR_SPYWARE ATTACK_TYPE_DETECTION_EVASION ATTACK_TYPE_VULNERABILITY_SCAN ATTACK_TYPE_ABUSE_OF_FUNCTIONALITY ATTACK_TYPE_AUTHENTICATION_AUTHORIZATION_ATTACKS ATTACK_TYPE_BUFFER_OVERFLOW ATTACK_TYPE_PREDICTABLE_RESOURCE_LOCATION ATTACK_TYPE_INFORMATION_LEAKAGE ATTACK_TYPE_DIRECTORY_INDEXING ATTACK_TYPE_PATH_TRAVERSAL ATTACK_TYPE_XPATH_INJECTION ATTACK_TYPE_LDAP_INJECTION ATTACK_TYPE_SERVER_SIDE_CODE_INJECTION ATTACK_TYPE_COMMAND_EXECUTION ATTACK_TYPE_SQL_INJECTION ATTACK_TYPE_CROSS_SITE_SCRIPTING ATTACK_TYPE_DENIAL_OF_SERVICE ATTACK_TYPE_HTTP_PARSER_ATTACK ATTACK_TYPE_SESSION_HIJACKING ATTACK_TYPE_HTTP_RESPONSE_SPLITTING ATTACK_TYPE_FORCEFUL_BROWSING ATTACK_TYPE_REMOTE_FILE_INCLUDE ATTACK_TYPE_MALICIOUS_FILE_UPLOAD ATTACK_TYPE_GRAPHQL_PARSER_ATTACK

#### WAF Exclusion WAF Exclusion Inline Rules Rules App Firewall Detection Control Exclude Bot Name Contexts

An [`exclude_bot_name_contexts`](#waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control-exclude-bot-name-contexts) block (within [`waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control`](#waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control)) supports the following:

<a id="waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control-exclude-bot-name-contexts-bot-name"></a>&#x2022; [`bot_name`](#waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control-exclude-bot-name-contexts-bot-name) - Optional String<br>Bot Name

#### WAF Exclusion WAF Exclusion Inline Rules Rules App Firewall Detection Control Exclude Signature Contexts

An [`exclude_signature_contexts`](#waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control-exclude-signature-contexts) block (within [`waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control`](#waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control)) supports the following:

<a id="waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control-exclude-signature-contexts-context"></a>&#x2022; [`context`](#waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control-exclude-signature-contexts-context) - Optional String  Defaults to `CONTEXT_ANY`<br>Possible values are `CONTEXT_ANY`, `CONTEXT_BODY`, `CONTEXT_REQUEST`, `CONTEXT_RESPONSE`, `CONTEXT_PARAMETER`, `CONTEXT_HEADER`, `CONTEXT_COOKIE`, `CONTEXT_URL`, `CONTEXT_URI`<br>WAF Exclusion Context Options. The available contexts for Exclusion rules. - CONTEXT_ANY: CONTEXT_ANY Detection will be excluded for all contexts. - CONTEXT_BODY: CONTEXT_BODY Detection will be excluded for the request body. - CONTEXT_REQUEST: CONTEXT_REQUEST Detection will be excluded for the request. - CONTEXT_RESPONSE: CONTEXT_RESPONSE - CONTEXT_PARAMETER: CONTEXT_PARAMETER Detection will be excluded for the parameters. The parameter name is required in the Context name field. If the field is left empty, the detection will be excluded for all parameters. - CONTEXT_HEADER: CONTEXT_HEADER Detection will be excluded for the headers. The header name is required in the Context name field. If the field is left empty, the detection will be excluded for all headers. - CONTEXT_COOKIE: CONTEXT_COOKIE Detection will be excluded for the cookies. The cookie name is required in the Context name field. If the field is left empty, the detection will be excluded for all cookies. - CONTEXT_URL: CONTEXT_URL Detection will be excluded for the request URL. - CONTEXT_URI: CONTEXT_URI

<a id="waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control-exclude-signature-contexts-context-name"></a>&#x2022; [`context_name`](#waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control-exclude-signature-contexts-context-name) - Optional String<br>Context Name. Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an wildcard asterisk (*)

<a id="waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control-exclude-signature-contexts-signature-id"></a>&#x2022; [`signature_id`](#waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control-exclude-signature-contexts-signature-id) - Optional Number<br>SignatureID. The allowed values for signature id are 0 and in the range of 200000001-299999999. 0 implies that all signatures will be excluded for the specified context

#### WAF Exclusion WAF Exclusion Inline Rules Rules App Firewall Detection Control Exclude Violation Contexts

An [`exclude_violation_contexts`](#waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control-exclude-violation-contexts) block (within [`waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control`](#waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control)) supports the following:

<a id="waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control-exclude-violation-contexts-context"></a>&#x2022; [`context`](#waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control-exclude-violation-contexts-context) - Optional String  Defaults to `CONTEXT_ANY`<br>Possible values are `CONTEXT_ANY`, `CONTEXT_BODY`, `CONTEXT_REQUEST`, `CONTEXT_RESPONSE`, `CONTEXT_PARAMETER`, `CONTEXT_HEADER`, `CONTEXT_COOKIE`, `CONTEXT_URL`, `CONTEXT_URI`<br>WAF Exclusion Context Options. The available contexts for Exclusion rules. - CONTEXT_ANY: CONTEXT_ANY Detection will be excluded for all contexts. - CONTEXT_BODY: CONTEXT_BODY Detection will be excluded for the request body. - CONTEXT_REQUEST: CONTEXT_REQUEST Detection will be excluded for the request. - CONTEXT_RESPONSE: CONTEXT_RESPONSE - CONTEXT_PARAMETER: CONTEXT_PARAMETER Detection will be excluded for the parameters. The parameter name is required in the Context name field. If the field is left empty, the detection will be excluded for all parameters. - CONTEXT_HEADER: CONTEXT_HEADER Detection will be excluded for the headers. The header name is required in the Context name field. If the field is left empty, the detection will be excluded for all headers. - CONTEXT_COOKIE: CONTEXT_COOKIE Detection will be excluded for the cookies. The cookie name is required in the Context name field. If the field is left empty, the detection will be excluded for all cookies. - CONTEXT_URL: CONTEXT_URL Detection will be excluded for the request URL. - CONTEXT_URI: CONTEXT_URI

<a id="waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control-exclude-violation-contexts-context-name"></a>&#x2022; [`context_name`](#waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control-exclude-violation-contexts-context-name) - Optional String<br>Context Name. Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an wildcard asterisk (*)

<a id="waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control-exclude-violation-contexts-exclude-violation"></a>&#x2022; [`exclude_violation`](#waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control-exclude-violation-contexts-exclude-violation) - Optional String  Defaults to `VIOL_NONE`<br>Possible values are `VIOL_NONE`, `VIOL_FILETYPE`, `VIOL_METHOD`, `VIOL_MANDATORY_HEADER`, `VIOL_HTTP_RESPONSE_STATUS`, `VIOL_REQUEST_MAX_LENGTH`, `VIOL_FILE_UPLOAD`, `VIOL_FILE_UPLOAD_IN_BODY`, `VIOL_XML_MALFORMED`, `VIOL_JSON_MALFORMED`, `VIOL_ASM_COOKIE_MODIFIED`, `VIOL_HTTP_PROTOCOL_MULTIPLE_HOST_HEADERS`, `VIOL_HTTP_PROTOCOL_BAD_HOST_HEADER_VALUE`, `VIOL_HTTP_PROTOCOL_UNPARSABLE_REQUEST_CONTENT`, `VIOL_HTTP_PROTOCOL_NULL_IN_REQUEST`, `VIOL_HTTP_PROTOCOL_BAD_HTTP_VERSION`, `VIOL_HTTP_PROTOCOL_SEVERAL_CONTENT_LENGTH_HEADERS`, `VIOL_EVASION_DIRECTORY_TRAVERSALS`, `VIOL_MALFORMED_REQUEST`, `VIOL_EVASION_MULTIPLE_DECODING`, `VIOL_DATA_GUARD`, `VIOL_EVASION_APACHE_WHITESPACE`, `VIOL_COOKIE_MODIFIED`, `VIOL_EVASION_IIS_UNICODE_CODEPOINTS`, `VIOL_EVASION_IIS_BACKSLASHES`, `VIOL_EVASION_PERCENT_U_DECODING`, `VIOL_EVASION_BARE_BYTE_DECODING`, `VIOL_EVASION_BAD_UNESCAPE`, `VIOL_HTTP_PROTOCOL_BODY_IN_GET_OR_HEAD_REQUEST`, `VIOL_ENCODING`, `VIOL_COOKIE_MALFORMED`, `VIOL_GRAPHQL_FORMAT`, `VIOL_GRAPHQL_MALFORMED`, `VIOL_GRAPHQL_INTROSPECTION_QUERY`<br>App Firewall Violation Type. List of all supported Violation Types VIOL_NONE VIOL_FILETYPE VIOL_METHOD VIOL_MANDATORY_HEADER VIOL_HTTP_RESPONSE_STATUS VIOL_REQUEST_MAX_LENGTH VIOL_FILE_UPLOAD VIOL_FILE_UPLOAD_IN_BODY VIOL_XML_MALFORMED VIOL_JSON_MALFORMED VIOL_ASM_COOKIE_MODIFIED VIOL_HTTP_PROTOCOL_MULTIPLE_HOST_HEADERS VIOL_HTTP_PROTOCOL_BAD_HOST_HEADER_VALUE VIOL_HTTP_PROTOCOL_UNPARSABLE_REQUEST_CONTENT VIOL_HTTP_PROTOCOL_NULL_IN_REQUEST VIOL_HTTP_PROTOCOL_BAD_HTTP_VERSION VIOL_HTTP_PROTOCOL_CRLF_CHARACTERS_BEFORE_REQUEST_START VIOL_HTTP_PROTOCOL_NO_HOST_HEADER_IN_HTTP_1_1_REQUEST VIOL_HTTP_PROTOCOL_BAD_MULTIPART_PARAMETERS_PARSING VIOL_HTTP_PROTOCOL_SEVERAL_CONTENT_LENGTH_HEADERS VIOL_HTTP_PROTOCOL_CONTENT_LENGTH_SHOULD_BE_A_POSITIVE_NUMBER VIOL_EVASION_DIRECTORY_TRAVERSALS VIOL_MALFORMED_REQUEST VIOL_EVASION_MULTIPLE_DECODING VIOL_DATA_GUARD VIOL_EVASION_APACHE_WHITESPACE VIOL_COOKIE_MODIFIED VIOL_EVASION_IIS_UNICODE_CODEPOINTS VIOL_EVASION_IIS_BACKSLASHES VIOL_EVASION_PERCENT_U_DECODING VIOL_EVASION_BARE_BYTE_DECODING VIOL_EVASION_BAD_UNESCAPE VIOL_HTTP_PROTOCOL_BAD_MULTIPART_FORMDATA_REQUEST_PARSING VIOL_HTTP_PROTOCOL_BODY_IN_GET_OR_HEAD_REQUEST VIOL_HTTP_PROTOCOL_HIGH_ASCII_CHARACTERS_IN_HEADERS VIOL_ENCODING VIOL_COOKIE_MALFORMED VIOL_GRAPHQL_FORMAT VIOL_GRAPHQL_MALFORMED VIOL_GRAPHQL_INTROSPECTION_QUERY

#### WAF Exclusion WAF Exclusion Inline Rules Rules Metadata

A [`metadata`](#waf-exclusion-waf-exclusion-inline-rules-rules-metadata) block (within [`waf_exclusion.waf_exclusion_inline_rules.rules`](#waf-exclusion-waf-exclusion-inline-rules-rules)) supports the following:

<a id="waf-exclusion-waf-exclusion-inline-rules-rules-metadata-description-spec"></a>&#x2022; [`description_spec`](#waf-exclusion-waf-exclusion-inline-rules-rules-metadata-description-spec) - Optional String<br>Description. Human readable description

<a id="waf-exclusion-waf-exclusion-inline-rules-rules-metadata-name"></a>&#x2022; [`name`](#waf-exclusion-waf-exclusion-inline-rules-rules-metadata-name) - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

#### WAF Exclusion WAF Exclusion Policy

A [`waf_exclusion_policy`](#waf-exclusion-waf-exclusion-policy) block (within [`waf_exclusion`](#waf-exclusion)) supports the following:

<a id="waf-exclusion-waf-exclusion-policy-name"></a>&#x2022; [`name`](#waf-exclusion-waf-exclusion-policy-name) - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

<a id="waf-exclusion-waf-exclusion-policy-namespace"></a>&#x2022; [`namespace`](#waf-exclusion-waf-exclusion-policy-namespace) - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

<a id="waf-exclusion-waf-exclusion-policy-tenant"></a>&#x2022; [`tenant`](#waf-exclusion-waf-exclusion-policy-tenant) - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

## Import

Import is supported using the following syntax:

```shell
# Import using namespace/name format
terraform import f5xc_http_loadbalancer.example system/example
```
