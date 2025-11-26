---
page_title: "f5xc_http_loadbalancer Resource - terraform-provider-f5xc"
subcategory: "Load Balancing"
description: |-
  Manages a HTTPLoadBalancer resource in F5 Distributed Cloud for load balancing HTTP/HTTPS traffic with advanced routing and security.
---

# f5xc_http_loadbalancer (Resource)

Manages a HTTPLoadBalancer resource in F5 Distributed Cloud for load balancing HTTP/HTTPS traffic with advanced routing and security.

~> **Note** For more information about this resource, please refer to the [F5 XC API Documentation](https://docs.cloud.f5.com/docs/api/).

## Example Usage

```terraform
# Http Loadbalancer Resource Example
# Manages a HTTPLoadBalancer resource in F5 Distributed Cloud for load balancing HTTP/HTTPS traffic with advanced routing and security.

# Basic Http Loadbalancer configuration
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

  // One of the arguments from this list "api_definition api_definitions api_specification disable_api_definition" must be set

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

&#x2022; `name` - Required String<br>Name of the HTTPLoadBalancer. Must be unique within the namespace

&#x2022; `namespace` - Required String<br>Namespace where the HTTPLoadBalancer will be created

&#x2022; `annotations` - Optional Map<br>Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata

&#x2022; `description` - Optional String<br>Human readable description for the object

&#x2022; `disable` - Optional Bool<br>A value of true will administratively disable the object

&#x2022; `labels` - Optional Map<br>Labels is a user defined key value map that can be attached to resources for organization and filtering

### Spec Argument Reference

-> **One of the following:**
&#x2022; `active_service_policies` - Optional Block<br>Service Policy List. List of service policies<br>See [Active Service Policies](#active-service-policies) below for details.
<br><br>&#x2022; `no_service_policies` - Optional Block<br>Empty. This can be used for messages where no values are needed
<br><br>&#x2022; `service_policies_from_namespace` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `add_location` - Optional Bool<br>Add Location. x-example: true Appends header x-volterra-location = <RE-site-name> in responses. This configuration is ignored on CE sites

-> **One of the following:**
&#x2022; `advertise_custom` - Optional Block<br>Advertise Custom. This defines a way to advertise a VIP on specific sites<br>See [Advertise Custom](#advertise-custom) below for details.
<br><br>&#x2022; `advertise_on_public` - Optional Block<br>Advertise Public. This defines a way to advertise a load balancer on public. If optional public_ip is provided, it will only be advertised on RE sites where that public_ip is available<br>See [Advertise On Public](#advertise-on-public) below for details.
<br><br>&#x2022; `advertise_on_public_default_vip` - Optional Block<br>Empty. This can be used for messages where no values are needed
<br><br>&#x2022; `do_not_advertise` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `api_protection_rules` - Optional Block<br>API Protection Rules. API Protection Rules<br>See [API Protection Rules](#api-protection-rules) below for details.

-> **One of the following:**
&#x2022; `api_rate_limit` - Optional Block<br>APIRateLimit
<br><br>&#x2022; `disable_rate_limit` - Optional Block<br>Empty. This can be used for messages where no values are needed
<br><br>&#x2022; `rate_limit` - Optional Block<br>RateLimitConfigType

-> **One of the following:**
&#x2022; `api_specification` - Optional Block<br>API Specification and Validation. Settings for API specification (API definition, OpenAPI validation, etc.)
<br><br>&#x2022; `disable_api_definition` - Optional Block<br>Empty. This can be used for messages where no values are needed

-> **One of the following:**
&#x2022; `api_testing` - Optional Block<br>API Testing
<br><br>&#x2022; `disable_api_testing` - Optional Block<br>Empty. This can be used for messages where no values are needed

-> **One of the following:**
&#x2022; `app_firewall` - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name
<br><br>&#x2022; `disable_waf` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `blocked_clients` - Optional Block<br>Client Blocking Rules. Define rules to block IP Prefixes or AS numbers

-> **One of the following:**
&#x2022; `bot_defense` - Optional Block<br>Bot Defense. This defines various configuration options for Bot Defense Policy
<br><br>&#x2022; `bot_defense_advanced` - Optional Block<br>Bot Defense Advanced. Bot Defense Advanced
<br><br>&#x2022; `disable_bot_defense` - Optional Block<br>Empty. This can be used for messages where no values are needed

-> **One of the following:**
&#x2022; `caching_policy` - Optional Block<br>Caching Policies. x-required Caching Policies for the CDN
<br><br>&#x2022; `disable_caching` - Optional Block<br>Empty. This can be used for messages where no values are needed

-> **One of the following:**
&#x2022; `captcha_challenge` - Optional Block<br>Captcha Challenge Parameters. Enables loadbalancer to perform captcha challenge Captcha challenge will be based on Google Recaptcha. With this feature enabled, only clients that pass the captcha challenge will be allowed to complete the HTTP request. When loadbalancer is configured to do Captcha Challenge, it will redirect the browser to an HTML page on every new HTTP request. This HTML page will have captcha challenge embedded in it. Client will be allowed to make the request only if the captcha challenge is successful. Loadbalancer will tag response header with a cookie to avoid Captcha challenge for subsequent requests. CAPTCHA is mainly used as a security check to ensure only human users can pass through. Generally, computers or bots are not capable of solving a captcha. You can enable either Javascript challenge or Captcha challenge on a virtual host
<br><br>&#x2022; `enable_challenge` - Optional Block<br>Enable Malicious User Challenge. Configure auto mitigation i.e risk based challenges for malicious users
<br><br>&#x2022; `js_challenge` - Optional Block<br>Javascript Challenge Parameters. Enables loadbalancer to perform client browser compatibility test by redirecting to a page with Javascript. With this feature enabled, only clients that are capable of executing Javascript(mostly browsers) will be allowed to complete the HTTP request. When loadbalancer is configured to do Javascript Challenge, it will redirect the browser to an HTML page on every new HTTP request. This HTML page will have Javascript embedded in it. Loadbalancer chooses a set of random numbers for every new client and sends these numbers along with an encrypted answer with the request such that it embed these numbers as input in the Javascript. Javascript will run on the requestor browser and perform a complex Math operation. Script will submit the answer to loadbalancer. Loadbalancer will validate the answer by comparing the calculated answer with the decrypted answer (which was encrypted when it was sent back as reply) and allow the request to the upstream server only if the answer is correct. Loadbalancer will tag response header with a cookie to avoid Javascript challenge for subsequent requests. Javascript challenge serves following purposes * Validate that the request is coming via a browser that is capable for running Javascript * Force the browser to run a complex operation, f(X), that requires it to spend a large number of CPU cycles. This is to slow down a potential DOS attacker by making it difficult to launch a large request flood without having to spend even larger CPU cost at their end. You can enable either Javascript challenge or Captcha challenge on a virtual host
<br><br>&#x2022; `no_challenge` - Optional Block<br>Empty. This can be used for messages where no values are needed
<br><br>&#x2022; `policy_based_challenge` - Optional Block<br>Policy Based Challenge. Specifies the settings for policy rule based challenge

-> **One of the following:**
&#x2022; `client_side_defense` - Optional Block<br>Client-Side Defense. This defines various configuration options for Client-Side Defense Policy
<br><br>&#x2022; `disable_client_side_defense` - Optional Block<br>Empty. This can be used for messages where no values are needed

-> **One of the following:**
&#x2022; `cookie_stickiness` - Optional Block<br>Hashing using Cookie. Two types of cookie affinity: 1. Passive. Takes a cookie that's present in the cookies header and hashes on its value. 2. Generated. Generates and sets a cookie with an expiration (TTL) on the first request from the client in its response to the client, based on the endpoint the request gets sent to. The client then presents this on the next and all subsequent requests. The hash of this is sufficient to ensure these requests get sent to the same endpoint. The cookie is generated by hashing the source and destination ports and addresses so that multiple independent HTTP2 streams on the same connection will independently receive the same cookie, even if they arrive simultaneously
<br><br>&#x2022; `least_active` - Optional Block<br>Empty. This can be used for messages where no values are needed
<br><br>&#x2022; `random` - Optional Block<br>Empty. This can be used for messages where no values are needed
<br><br>&#x2022; `ring_hash` - Optional Block<br>Hash Policy List. List of hash policy rules
<br><br>&#x2022; `round_robin` - Optional Block<br>Empty. This can be used for messages where no values are needed
<br><br>&#x2022; `source_ip_stickiness` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `cors_policy` - Optional Block<br>CORS Policy. Cross-Origin Resource Sharing requests configuration specified at Virtual-host or Route level. Route level configuration takes precedence. An example of an Cross origin HTTP request GET /resources/public-data/ HTTP/1.1 Host: bar.other User-Agent: Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.5; en-US; rv:1.9.1b3pre) Gecko/20081130 Minefield/3.1b3pre Accept: text/HTML,application/xhtml+XML,application/XML;q=0.9,*/*;q=0.8 Accept-Language: en-us,en;q=0.5 Accept-Encoding: gzip,deflate Accept-Charset: ISO-8859-1,utf-8;q=0.7,*;q=0.7 Connection: keep-alive Referrer: `HTTP://foo.example/examples/access-control/simpleXSInvocation.HTML` Origin: `HTTP://foo.example` HTTP/1.1 200 OK Date: Mon, 01 Dec 2008 00:23:53 GMT Server: Apache/2.0.61 Access-Control-Allow-Origin: * Keep-Alive: timeout=2, max=100 Connection: Keep-Alive Transfer-Encoding: chunked Content-Type: application/XML An example for cross origin HTTP OPTIONS request with Access-Control-Request-* header OPTIONS /resources/post-here/ HTTP/1.1 Host: bar.other User-Agent: Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.5; en-US; rv:1.9.1b3pre) Gecko/20081130 Minefield/3.1b3pre Accept: text/HTML,application/xhtml+XML,application/XML;q=0.9,*/*;q=0.8 Accept-Language: en-us,en;q=0.5 Accept-Encoding: gzip,deflate Accept-Charset: ISO-8859-1,utf-8;q=0.7,*;q=0.7 Connection: keep-alive Origin: `HTTP://foo.example` Access-Control-Request-Method: POST Access-Control-Request-Headers: X-PINGOTHER, Content-Type HTTP/1.1 204 No Content Date: Mon, 01 Dec 2008 01:15:39 GMT Server: Apache/2.0.61 (Unix) Access-Control-Allow-Origin: `HTTP://foo.example` Access-Control-Allow-Methods: POST, GET, OPTIONS Access-Control-Allow-Headers: X-PINGOTHER, Content-Type Access-Control-Max-Age: 86400 Vary: Accept-Encoding, Origin Keep-Alive: timeout=2, max=100 Connection: Keep-Alive

&#x2022; `csrf_policy` - Optional Block<br>CSRF Policy. To mitigate CSRF attack , the policy checks where a request is coming from to determine if the request's origin is the same as its detination.The policy relies on two pieces of information used in determining if a request originated from the same host. 1. The origin that caused the user agent to issue the request (source origin). 2. The origin that the request is going to (target origin). When the policy evaluating a request, it ensures both pieces of information are present and compare their values. If the source origin is missing or origins do not match the request is rejected. The exception to this being if the source-origin has been added to they policy as valid. Because CSRF attacks specifically target state-changing requests, the policy only acts on the HTTP requests that have state-changing method (PUT,POST, etc.)

&#x2022; `data_guard_rules` - Optional Block<br>Data Guard Rules. Data Guard prevents responses from exposing sensitive information by masking the data. The system masks credit card numbers and social security numbers leaked from the application from within the HTTP response with a string of asterisks (*). Note: App Firewall should be enabled, to use Data Guard feature

&#x2022; `ddos_mitigation_rules` - Optional Block<br>DDOS Mitigation Rules. Define manual mitigation rules to block L7 DDOS attacks

-> **One of the following:**
&#x2022; `default_pool` - Optional Block<br>Global Specification. Shape of the origin pool specification
<br><br>&#x2022; `default_pool_list` - Optional Block<br>Origin Pool List Type. List of Origin Pools

&#x2022; `default_route_pools` - Optional Block<br>Origin Pools. Origin Pools used when no route is specified (default route)

-> **One of the following:**
&#x2022; `default_sensitive_data_policy` - Optional Block<br>Empty. This can be used for messages where no values are needed
<br><br>&#x2022; `sensitive_data_policy` - Optional Block<br>Sensitive Data Discovery. Settings for data type policy

-> **One of the following:**
&#x2022; `disable_api_discovery` - Optional Block<br>Empty. This can be used for messages where no values are needed
<br><br>&#x2022; `enable_api_discovery` - Optional Block<br>API Discovery Setting. Specifies the settings used for API discovery

-> **One of the following:**
&#x2022; `disable_ip_reputation` - Optional Block<br>Empty. This can be used for messages where no values are needed
<br><br>&#x2022; `enable_ip_reputation` - Optional Block<br>IP Threat Category List. List of IP threat categories

-> **One of the following:**
&#x2022; `disable_malicious_user_detection` - Optional Block<br>Empty. This can be used for messages where no values are needed
<br><br>&#x2022; `enable_malicious_user_detection` - Optional Block<br>Empty. This can be used for messages where no values are needed

-> **One of the following:**
&#x2022; `disable_malware_protection` - Optional Block<br>Empty. This can be used for messages where no values are needed
<br><br>&#x2022; `malware_protection_settings` - Optional Block<br>Malware Protection Policy. Malware Protection protects Web Apps and APIs, from malicious file uploads by scanning files in real-time

-> **One of the following:**
&#x2022; `disable_threat_mesh` - Optional Block<br>Empty. This can be used for messages where no values are needed
<br><br>&#x2022; `enable_threat_mesh` - Optional Block<br>Empty. This can be used for messages where no values are needed

-> **One of the following:**
&#x2022; `disable_trust_client_ip_headers` - Optional Block<br>Empty. This can be used for messages where no values are needed
<br><br>&#x2022; `enable_trust_client_ip_headers` - Optional Block<br>Trust Client IP Headers List. List of Client IP Headers

&#x2022; `domains` - Optional List<br>Domains. A list of Domains (host/authority header) that will be matched to load balancer. Supported Domains and search order: 1. Exact Domain names: `www.foo.com.` 2. Domains starting with a Wildcard: *.foo.com. Not supported Domains: - Just a Wildcard: * - A Wildcard and TLD with no root Domain: *.com. - A Wildcard not matching a whole DNS label. e.g. *.foo.com and *.bar.foo.com are valid Wildcards however *bar.foo.com, *-bar.foo.com, and bar*.foo.com are all invalid. Additional notes: A Wildcard will not match empty string. e.g. *.foo.com will match bar.foo.com and baz-bar.foo.com but not .foo.com. The longest Wildcards match first. Only a single virtual host in the entire route configuration can match on *. Also a Domain must be unique across all virtual hosts within an advertise policy. Domains are also used for SNI matching if the Loadbalancer type is HTTPS. Domains also indicate the list of names for which DNS resolution will be automatically resolved to IP addresses by the system

&#x2022; `graphql_rules` - Optional Block<br>GraphQL Inspection. GraphQL is a query language and server-side runtime for APIs which provides a complete and understandable description of the data in API. GraphQL gives clients the power to ask for exactly what they need, makes it easier to evolve APIs over time, and enables powerful developer tools. Policy configuration to analyze GraphQL queries and prevent GraphQL tailored attacks

-> **One of the following:**
&#x2022; `http` - Optional Block<br>HTTP Choice. Choice for selecting HTTP proxy
<br><br>&#x2022; `https` - Optional Block<br>BYOC HTTPS Choice. Choice for selecting HTTP proxy with bring your own certificates
<br><br>&#x2022; `https_auto_cert` - Optional Block<br>HTTPS with Auto Certs Choice. Choice for selecting HTTP proxy with bring your own certificates

&#x2022; `jwt_validation` - Optional Block<br>JWT Validation. JWT Validation stops JWT replay attacks and JWT tampering by cryptographically verifying incoming JWTs before they are passed to your API origin. JWT Validation will also stop requests with expired tokens or tokens that are not yet valid

-> **One of the following:**
&#x2022; `l7_ddos_action_block` - Optional Block<br>Empty. This can be used for messages where no values are needed
<br><br>&#x2022; `l7_ddos_action_default` - Optional Block<br>Empty. This can be used for messages where no values are needed
<br><br>&#x2022; `l7_ddos_action_js_challenge` - Optional Block<br>Javascript Challenge Parameters. Enables loadbalancer to perform client browser compatibility test by redirecting to a page with Javascript. With this feature enabled, only clients that are capable of executing Javascript(mostly browsers) will be allowed to complete the HTTP request. When loadbalancer is configured to do Javascript Challenge, it will redirect the browser to an HTML page on every new HTTP request. This HTML page will have Javascript embedded in it. Loadbalancer chooses a set of random numbers for every new client and sends these numbers along with an encrypted answer with the request such that it embed these numbers as input in the Javascript. Javascript will run on the requestor browser and perform a complex Math operation. Script will submit the answer to loadbalancer. Loadbalancer will validate the answer by comparing the calculated answer with the decrypted answer (which was encrypted when it was sent back as reply) and allow the request to the upstream server only if the answer is correct. Loadbalancer will tag response header with a cookie to avoid Javascript challenge for subsequent requests. Javascript challenge serves following purposes * Validate that the request is coming via a browser that is capable for running Javascript * Force the browser to run a complex operation, f(X), that requires it to spend a large number of CPU cycles. This is to slow down a potential DOS attacker by making it difficult to launch a large request flood without having to spend even larger CPU cost at their end. You can enable either Javascript challenge or Captcha challenge on a virtual host

&#x2022; `l7_ddos_protection` - Optional Block<br>L7 DDOS Protection Settings. L7 DDOS protection is critical for safeguarding web applications, APIs, and services that are exposed to the internet from sophisticated, volumetric, application-level threats. Configure actions, thresholds and policies to apply during L7 DDOS attack

&#x2022; `more_option` - Optional Block<br>Advanced Options. This defines various options to define a route

-> **One of the following:**
&#x2022; `multi_lb_app` - Optional Block<br>Empty. This can be used for messages where no values are needed
<br><br>&#x2022; `single_lb_app` - Optional Block<br>Single Load Balancer App Setting. Specific settings for Machine learning analysis on this HTTP LB, independently from other LBs

&#x2022; `origin_server_subset_rule_list` - Optional Block<br>Origin Server Subset Rule List Type. List of Origin Pools

&#x2022; `protected_cookies` - Optional Block<br>Cookie Protection. Allows setting attributes (SameSite, Secure, and HttpOnly) on cookies in responses. Cookie Tampering Protection prevents attackers from modifying the value of session cookies. For Cookie Tampering Protection, enabling a web app firewall (WAF) is a prerequisite. The configured mode of WAF (monitoring or blocking) will be enforced on the request when cookie tampering is identified. Note: We recommend enabling Secure and HttpOnly attributes along with cookie tampering protection

&#x2022; `routes` - Optional Block<br>Routes. Routes allow users to define match condition on a path and/or HTTP method to either forward matching traffic to origin pool or redirect matching traffic to a different URL or respond directly to matching traffic

&#x2022; `sensitive_data_disclosure_rules` - Optional Block<br>Sensitive Data Exposure Rules. Sensitive Data Exposure Rules allows specifying rules to mask sensitive data fields in API responses

-> **One of the following:**
&#x2022; `slow_ddos_mitigation` - Optional Block<br>Slow DDOS Mitigation. 'Slow and low' attacks tie up server resources, leaving none available for servicing requests from actual users
<br><br>&#x2022; `system_default_timeouts` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `timeouts` - Optional Block

&#x2022; `trusted_clients` - Optional Block<br>Trusted Client Rules. Define rules to skip processing of one or more features such as WAF, Bot Defense etc. for clients

-> **One of the following:**
&#x2022; `user_id_client_ip` - Optional Block<br>Empty. This can be used for messages where no values are needed
<br><br>&#x2022; `user_identification` - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name

&#x2022; `waf_exclusion` - Optional Block<br>WAF Exclusion

### Attributes Reference

In addition to all arguments above, the following attributes are exported:

&#x2022; `id` - Optional String<br>Unique identifier for the resource

---

<a id="active-service-policies"></a>

**Active Service Policies**

&#x2022; `policies` - Optional Block<br>Policies. Service Policies is a sequential engine where policies (and rules within the policy) are evaluated one after the other. It's important to define the correct order (policies evaluated from top to bottom in the list) for service policies, to get the intended result. For each request, its characteristics are evaluated based on the match criteria in each service policy starting at the top. If there is a match in the current policy, then the policy takes effect, and no more policies are evaluated. Otherwise, the next policy is evaluated. If all policies are evaluated and none match, then the request will be denied by default<br>See [Policies](#active-service-policies-policies) below.

<a id="active-service-policies-policies"></a>

**Active Service Policies Policies**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="advertise-custom"></a>

**Advertise Custom**

&#x2022; `advertise_where` - Optional Block<br>List of Sites to Advertise. Where should this load balancer be available<br>See [Advertise Where](#advertise-custom-advertise-where) below.

<a id="advertise-custom-advertise-where"></a>

**Advertise Custom Advertise Where**

&#x2022; `advertise_on_public` - Optional Block<br>Advertise Public. This defines a way to advertise a load balancer on public. If optional public_ip is provided, it will only be advertised on RE sites where that public_ip is available<br>See [Advertise On Public](#advertise-custom-advertise-where-advertise-on-public) below.

&#x2022; `port` - Optional Number<br>Listen Port. Port to Listen

&#x2022; `port_ranges` - Optional String<br>Listen Port Ranges. A string containing a comma separated list of port ranges. Each port range consists of a single port or two ports separated by '-'

&#x2022; `site` - Optional Block<br>Site. This defines a reference to a CE site along with network type and an optional IP address where a load balancer could be advertised<br>See [Site](#advertise-custom-advertise-where-site) below.

&#x2022; `use_default_port` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `virtual_network` - Optional Block<br>Virtual Network. Parameters to advertise on a given virtual network<br>See [Virtual Network](#advertise-custom-advertise-where-virtual-network) below.

&#x2022; `virtual_site` - Optional Block<br>Virtual Site. This defines a reference to a customer site virtual site along with network type where a load balancer could be advertised<br>See [Virtual Site](#advertise-custom-advertise-where-virtual-site) below.

&#x2022; `virtual_site_with_vip` - Optional Block<br>Virtual Site with Specified VIP. This defines a reference to a customer site virtual site along with network type and IP where a load balancer could be advertised<br>See [Virtual Site With VIP](#advertise-custom-advertise-where-virtual-site-with-vip) below.

&#x2022; `vk8s_service` - Optional Block<br>vK8s Services on RE. This defines a reference to a RE site or virtual site where a load balancer could be advertised in the vK8s service network<br>See [Vk8s Service](#advertise-custom-advertise-where-vk8s-service) below.

<a id="advertise-custom-advertise-where-advertise-on-public"></a>

**Advertise Custom Advertise Where Advertise On Public**

&#x2022; `public_ip` - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Public IP](#advertise-custom-advertise-where-advertise-on-public-public-ip) below.

<a id="advertise-custom-advertise-where-advertise-on-public-public-ip"></a>

**Advertise Custom Advertise Where Advertise On Public Public IP**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="advertise-custom-advertise-where-site"></a>

**Advertise Custom Advertise Where Site**

&#x2022; `ip` - Optional String<br>IP Address. Use given IP address as VIP on the site

&#x2022; `network` - Optional String  Defaults to `SITE_NETWORK_INSIDE_AND_OUTSIDE`<br>Possible values are `SITE_NETWORK_INSIDE_AND_OUTSIDE`, `SITE_NETWORK_INSIDE`, `SITE_NETWORK_OUTSIDE`, `SITE_NETWORK_SERVICE`, `SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP`, `SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP`, `SITE_NETWORK_IP_FABRIC`<br>Site Network. This defines network types to be used on site All inside and outside networks. All inside and outside networks with internet VIP support. All inside networks. All outside networks. All outside networks with internet VIP support. vK8s service network. - SITE_NETWORK_IP_FABRIC: VER IP Fabric network for the site This Virtual network type is used for exposing virtual host on IP Fabric network on the VER site or for endpoint in IP Fabric network

&#x2022; `site` - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Site](#advertise-custom-advertise-where-site-site) below.

<a id="advertise-custom-advertise-where-site-site"></a>

**Advertise Custom Advertise Where Site Site**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="advertise-custom-advertise-where-virtual-network"></a>

**Advertise Custom Advertise Where Virtual Network**

&#x2022; `default_v6_vip` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `default_vip` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `specific_v6_vip` - Optional String<br>Specific V6 VIP. Use given IPv6 address as VIP on virtual Network

&#x2022; `specific_vip` - Optional String<br>Specific V4 VIP. Use given IPv4 address as VIP on virtual Network

&#x2022; `virtual_network` - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Virtual Network](#advertise-custom-advertise-where-virtual-network-virtual-network) below.

<a id="advertise-custom-advertise-where-virtual-network-virtual-network"></a>

**Advertise Custom Advertise Where Virtual Network Virtual Network**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="advertise-custom-advertise-where-virtual-site"></a>

**Advertise Custom Advertise Where Virtual Site**

&#x2022; `network` - Optional String  Defaults to `SITE_NETWORK_INSIDE_AND_OUTSIDE`<br>Possible values are `SITE_NETWORK_INSIDE_AND_OUTSIDE`, `SITE_NETWORK_INSIDE`, `SITE_NETWORK_OUTSIDE`, `SITE_NETWORK_SERVICE`, `SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP`, `SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP`, `SITE_NETWORK_IP_FABRIC`<br>Site Network. This defines network types to be used on site All inside and outside networks. All inside and outside networks with internet VIP support. All inside networks. All outside networks. All outside networks with internet VIP support. vK8s service network. - SITE_NETWORK_IP_FABRIC: VER IP Fabric network for the site This Virtual network type is used for exposing virtual host on IP Fabric network on the VER site or for endpoint in IP Fabric network

&#x2022; `virtual_site` - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Virtual Site](#advertise-custom-advertise-where-virtual-site-virtual-site) below.

<a id="advertise-custom-advertise-where-virtual-site-virtual-site"></a>

**Advertise Custom Advertise Where Virtual Site Virtual Site**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="advertise-custom-advertise-where-virtual-site-with-vip"></a>

**Advertise Custom Advertise Where Virtual Site With VIP**

&#x2022; `ip` - Optional String<br>IP Address. Use given IP address as VIP on the site

&#x2022; `network` - Optional String  Defaults to `SITE_NETWORK_SPECIFIED_VIP_OUTSIDE`<br>Possible values are `SITE_NETWORK_SPECIFIED_VIP_OUTSIDE`, `SITE_NETWORK_SPECIFIED_VIP_INSIDE`<br>Site Network. This defines network types to be used on virtual-site with specified VIP All outside networks. All inside networks

&#x2022; `virtual_site` - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Virtual Site](#advertise-custom-advertise-where-virtual-site-with-vip-virtual-site) below.

<a id="advertise-custom-advertise-where-virtual-site-with-vip-virtual-site"></a>

**Advertise Custom Advertise Where Virtual Site With VIP Virtual Site**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="advertise-custom-advertise-where-vk8s-service"></a>

**Advertise Custom Advertise Where Vk8s Service**

&#x2022; `site` - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Site](#advertise-custom-advertise-where-vk8s-service-site) below.

&#x2022; `virtual_site` - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Virtual Site](#advertise-custom-advertise-where-vk8s-service-virtual-site) below.

<a id="advertise-custom-advertise-where-vk8s-service-site"></a>

**Advertise Custom Advertise Where Vk8s Service Site**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="advertise-custom-advertise-where-vk8s-service-virtual-site"></a>

**Advertise Custom Advertise Where Vk8s Service Virtual Site**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="advertise-on-public"></a>

**Advertise On Public**

&#x2022; `public_ip` - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Public IP](#advertise-on-public-public-ip) below.

<a id="advertise-on-public-public-ip"></a>

**Advertise On Public Public IP**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="api-protection-rules"></a>

**API Protection Rules**

&#x2022; `api_endpoint_rules` - Optional Block<br>API Endpoints. This category defines specific rules per API endpoints. If request matches any of these rules, skipping second category rules<br>See [API Endpoint Rules](#api-protection-rules-api-endpoint-rules) below.

&#x2022; `api_groups_rules` - Optional Block<br>Server URLs and API Groups. This category includes rules per API group or Server URL. For API groups, refer to API Definition which includes API groups derived from uploaded swaggers<br>See [API Groups Rules](#api-protection-rules-api-groups-rules) below.

<a id="api-protection-rules-api-endpoint-rules"></a>

**API Protection Rules API Endpoint Rules**

&#x2022; `action` - Optional Block<br>API Protection Rule Action. The action to take if the input request matches the rule<br>See [Action](#api-protection-rules-api-endpoint-rules-action) below.

&#x2022; `any_domain` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `api_endpoint_method` - Optional Block<br>HTTP Method Matcher. A HTTP method matcher specifies a list of methods to match an input HTTP method. The match is considered successful if the input method is a member of the list. The result of the match based on the method list is inverted if invert_matcher is true<br>See [API Endpoint Method](#api-protection-rules-api-endpoint-rules-api-endpoint-method) below.

&#x2022; `api_endpoint_path` - Optional String<br>API Endpoint. The endpoint (path) of the request

&#x2022; `client_matcher` - Optional Block<br>Client Matcher. Client conditions for matching a rule<br>See [Client Matcher](#api-protection-rules-api-endpoint-rules-client-matcher) below.

&#x2022; `metadata` - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#api-protection-rules-api-endpoint-rules-metadata) below.

&#x2022; `request_matcher` - Optional Block<br>Request Matcher. Request conditions for matching a rule<br>See [Request Matcher](#api-protection-rules-api-endpoint-rules-request-matcher) below.

&#x2022; `specific_domain` - Optional String<br>Specific Domain. The rule will apply for a specific domain. For example: API.example.com

<a id="api-protection-rules-api-endpoint-rules-action"></a>

**API Protection Rules API Endpoint Rules Action**

&#x2022; `allow` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `deny` - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-protection-rules-api-endpoint-rules-api-endpoint-method"></a>

**API Protection Rules API Endpoint Rules API Endpoint Method**

&#x2022; `invert_matcher` - Optional Bool<br>Invert Method Matcher. Invert the match result

&#x2022; `methods` - Optional List  Defaults to `ANY`<br>Possible values are `ANY`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`, `COPY`<br>Method List. List of methods values to match against

<a id="api-protection-rules-api-endpoint-rules-client-matcher"></a>

**API Protection Rules API Endpoint Rules Client Matcher**

&#x2022; `any_client` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `any_ip` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `asn_list` - Optional Block<br>ASN Match List. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer<br>See [Asn List](#api-protection-rules-api-endpoint-rules-client-matcher-asn-list) below.

&#x2022; `asn_matcher` - Optional Block<br>ASN Matcher. Match any AS number contained in the list of bgp_asn_sets<br>See [Asn Matcher](#api-protection-rules-api-endpoint-rules-client-matcher-asn-matcher) below.

&#x2022; `client_selector` - Optional Block<br>Label Selector. This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expresssions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects. Label selector is immutable. expressions is a list of strings of label selection expression. Each string has ',' separated values which are 'AND' and all strings are logically 'OR'. BNF for expression string <selector-syntax> ::= <requirement> | <requirement> ',' <selector-syntax> <requirement> ::= [!] KEY [ <set-based-restriction> | <exact-match-restriction> ] <set-based-restriction> ::= '' | <inclusion-exclusion> <value-set> <inclusion-exclusion> ::= <inclusion> | <exclusion> <exclusion> ::= 'notin' <inclusion> ::= 'in' <value-set> ::= '(' <values> ')' <values> ::= VALUE | VALUE ',' <values> <exact-match-restriction> ::= ['='|'=='|'!='] VALUE<br>See [Client Selector](#api-protection-rules-api-endpoint-rules-client-matcher-client-selector) below.

&#x2022; `ip_matcher` - Optional Block<br>IP Prefix Matcher. Match any IP prefix contained in the list of ip_prefix_sets. The result of the match is inverted if invert_matcher is true<br>See [IP Matcher](#api-protection-rules-api-endpoint-rules-client-matcher-ip-matcher) below.

&#x2022; `ip_prefix_list` - Optional Block<br>IP Prefix Match List. List of IP Prefix strings to match against<br>See [IP Prefix List](#api-protection-rules-api-endpoint-rules-client-matcher-ip-prefix-list) below.

&#x2022; `ip_threat_category_list` - Optional Block<br>IP Threat Category List Type. List of IP threat categories<br>See [IP Threat Category List](#api-protection-rules-api-endpoint-rules-client-matcher-ip-threat-category-list) below.

&#x2022; `tls_fingerprint_matcher` - Optional Block<br>TLS Fingerprint Matcher. A TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of supported positve match criteria includes a list of known classes of TLS fingerprints and a list of exact values. The match is considered successful if either of these positive criteria are satisfied and the input fingerprint is not one of the excluded values<br>See [TLS Fingerprint Matcher](#api-protection-rules-api-endpoint-rules-client-matcher-tls-fingerprint-matcher) below.

<a id="api-protection-rules-api-endpoint-rules-client-matcher-asn-list"></a>

**API Protection Rules API Endpoint Rules Client Matcher Asn List**

&#x2022; `as_numbers` - Optional List<br>AS Numbers. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer

<a id="api-protection-rules-api-endpoint-rules-client-matcher-asn-matcher"></a>

**API Protection Rules API Endpoint Rules Client Matcher Asn Matcher**

&#x2022; `asn_sets` - Optional Block<br>BGP ASN Sets. A list of references to bgp_asn_set objects<br>See [Asn Sets](#api-protection-rules-api-endpoint-rules-client-matcher-asn-matcher-asn-sets) below.

<a id="api-protection-rules-api-endpoint-rules-client-matcher-asn-matcher-asn-sets"></a>

**API Protection Rules API Endpoint Rules Client Matcher Asn Matcher Asn Sets**

&#x2022; `kind` - Optional String<br>Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route')

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

&#x2022; `uid` - Optional String<br>UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid

<a id="api-protection-rules-api-endpoint-rules-client-matcher-client-selector"></a>

**API Protection Rules API Endpoint Rules Client Matcher Client Selector**

&#x2022; `expressions` - Optional List<br>Selector Expression. expressions contains the kubernetes style label expression for selections

<a id="api-protection-rules-api-endpoint-rules-client-matcher-ip-matcher"></a>

**API Protection Rules API Endpoint Rules Client Matcher IP Matcher**

&#x2022; `invert_matcher` - Optional Bool<br>Invert IP Matcher. Invert the match result

&#x2022; `prefix_sets` - Optional Block<br>IP Prefix Sets. A list of references to ip_prefix_set objects<br>See [Prefix Sets](#api-protection-rules-api-endpoint-rules-client-matcher-ip-matcher-prefix-sets) below.

<a id="api-protection-rules-api-endpoint-rules-client-matcher-ip-matcher-prefix-sets"></a>

**API Protection Rules API Endpoint Rules Client Matcher IP Matcher Prefix Sets**

&#x2022; `kind` - Optional String<br>Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route')

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

&#x2022; `uid` - Optional String<br>UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid

<a id="api-protection-rules-api-endpoint-rules-client-matcher-ip-prefix-list"></a>

**API Protection Rules API Endpoint Rules Client Matcher IP Prefix List**

&#x2022; `invert_match` - Optional Bool<br>Invert Match Result. Invert the match result

&#x2022; `ip_prefixes` - Optional List<br>IPv4 Prefix List. List of IPv4 prefix strings

<a id="api-protection-rules-api-endpoint-rules-client-matcher-ip-threat-category-list"></a>

**API Protection Rules API Endpoint Rules Client Matcher IP Threat Category List**

&#x2022; `ip_threat_categories` - Optional List  Defaults to `SPAM_SOURCES`<br>Possible values are `SPAM_SOURCES`, `WINDOWS_EXPLOITS`, `WEB_ATTACKS`, `BOTNETS`, `SCANNERS`, `REPUTATION`, `PHISHING`, `PROXY`, `MOBILE_THREATS`, `TOR_PROXY`, `DENIAL_OF_SERVICE`, `NETWORK`<br>List of IP Threat Categories to choose. The IP threat categories is obtained from the list and is used to auto-generate equivalent label selection expressions

<a id="api-protection-rules-api-endpoint-rules-client-matcher-tls-fingerprint-matcher"></a>

**API Protection Rules API Endpoint Rules Client Matcher TLS Fingerprint Matcher**

&#x2022; `classes` - Optional List  Defaults to `TLS_FINGERPRINT_NONE`<br>Possible values are `TLS_FINGERPRINT_NONE`, `ANY_MALICIOUS_FINGERPRINT`, `ADWARE`, `ADWIND`, `DRIDEX`, `GOOTKIT`, `GOZI`, `JBIFROST`, `QUAKBOT`, `RANSOMWARE`, `TROLDESH`, `TOFSEE`, `TORRENTLOCKER`, `TRICKBOT`<br>TLS fingerprint classes. A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against

&#x2022; `exact_values` - Optional List<br>Exact Values. A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against

&#x2022; `excluded_values` - Optional List<br>Excluded Values. A list of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can be used to skip known false positives when using one or more known TLS fingerprint classes in the enclosing matcher

<a id="api-protection-rules-api-endpoint-rules-metadata"></a>

**API Protection Rules API Endpoint Rules Metadata**

&#x2022; `description` - Optional String<br>Description. Human readable description

&#x2022; `name` - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

<a id="api-protection-rules-api-endpoint-rules-request-matcher"></a>

**API Protection Rules API Endpoint Rules Request Matcher**

&#x2022; `cookie_matchers` - Optional Block<br>Cookie Matchers. A list of predicates for all cookies that need to be matched. The criteria for matching each cookie is described in individual instances of CookieMatcherType. The actual cookie values are extracted from the request API as a list of strings for each cookie name. Note that all specified cookie matcher predicates must evaluate to true<br>See [Cookie Matchers](#api-protection-rules-api-endpoint-rules-request-matcher-cookie-matchers) below.

&#x2022; `headers` - Optional Block<br>HTTP Headers. A list of predicates for various HTTP headers that need to match. The criteria for matching each HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values are extracted from the request API as a list of strings for each HTTP header type. Note that all specified header predicates must evaluate to true<br>See [Headers](#api-protection-rules-api-endpoint-rules-request-matcher-headers) below.

&#x2022; `jwt_claims` - Optional Block<br>JWT Claims. A list of predicates for various JWT claims that need to match. The criteria for matching each JWT claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are extracted from the JWT payload as a list of strings. Note that all specified JWT claim predicates must evaluate to true. Note that this feature only works on LBs with JWT Validation feature enabled<br>See [JWT Claims](#api-protection-rules-api-endpoint-rules-request-matcher-jwt-claims) below.

&#x2022; `query_params` - Optional Block<br>HTTP Query Parameters. A list of predicates for all query parameters that need to be matched. The criteria for matching each query parameter are described in individual instances of QueryParameterMatcherType. The actual query parameter values are extracted from the request API as a list of strings for each query parameter name. Note that all specified query parameter predicates must evaluate to true<br>See [Query Params](#api-protection-rules-api-endpoint-rules-request-matcher-query-params) below.

<a id="api-protection-rules-api-endpoint-rules-request-matcher-cookie-matchers"></a>

**API Protection Rules API Endpoint Rules Request Matcher Cookie Matchers**

&#x2022; `check_not_present` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `check_present` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `invert_matcher` - Optional Bool<br>Invert Matcher. Invert Match of the expression defined

&#x2022; `item` - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#api-protection-rules-api-endpoint-rules-request-matcher-cookie-matchers-item) below.

&#x2022; `name` - Optional String<br>Cookie Name. A case-sensitive cookie name

<a id="api-protection-rules-api-endpoint-rules-request-matcher-cookie-matchers-item"></a>

**API Protection Rules API Endpoint Rules Request Matcher Cookie Matchers Item**

&#x2022; `exact_values` - Optional List<br>Exact Values. A list of exact values to match the input against

&#x2022; `regex_values` - Optional List<br>Regex Values. A list of regular expressions to match the input against

&#x2022; `transformers` - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

<a id="api-protection-rules-api-endpoint-rules-request-matcher-headers"></a>

**API Protection Rules API Endpoint Rules Request Matcher Headers**

&#x2022; `check_not_present` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `check_present` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `invert_matcher` - Optional Bool<br>Invert Header Matcher. Invert the match result

&#x2022; `item` - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#api-protection-rules-api-endpoint-rules-request-matcher-headers-item) below.

&#x2022; `name` - Optional String<br>Header Name. A case-insensitive HTTP header name

<a id="api-protection-rules-api-endpoint-rules-request-matcher-headers-item"></a>

**API Protection Rules API Endpoint Rules Request Matcher Headers Item**

&#x2022; `exact_values` - Optional List<br>Exact Values. A list of exact values to match the input against

&#x2022; `regex_values` - Optional List<br>Regex Values. A list of regular expressions to match the input against

&#x2022; `transformers` - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

<a id="api-protection-rules-api-endpoint-rules-request-matcher-jwt-claims"></a>

**API Protection Rules API Endpoint Rules Request Matcher JWT Claims**

&#x2022; `check_not_present` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `check_present` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `invert_matcher` - Optional Bool<br>Invert Matcher. Invert the match result

&#x2022; `item` - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#api-protection-rules-api-endpoint-rules-request-matcher-jwt-claims-item) below.

&#x2022; `name` - Optional String<br>JWT Claim Name. JWT claim name

<a id="api-protection-rules-api-endpoint-rules-request-matcher-jwt-claims-item"></a>

**API Protection Rules API Endpoint Rules Request Matcher JWT Claims Item**

&#x2022; `exact_values` - Optional List<br>Exact Values. A list of exact values to match the input against

&#x2022; `regex_values` - Optional List<br>Regex Values. A list of regular expressions to match the input against

&#x2022; `transformers` - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

<a id="api-protection-rules-api-endpoint-rules-request-matcher-query-params"></a>

**API Protection Rules API Endpoint Rules Request Matcher Query Params**

&#x2022; `check_not_present` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `check_present` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `invert_matcher` - Optional Bool<br>Invert Query Parameter Matcher. Invert the match result

&#x2022; `item` - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#api-protection-rules-api-endpoint-rules-request-matcher-query-params-item) below.

&#x2022; `key` - Optional String<br>Query Parameter Name. A case-sensitive HTTP query parameter name

<a id="api-protection-rules-api-endpoint-rules-request-matcher-query-params-item"></a>

**API Protection Rules API Endpoint Rules Request Matcher Query Params Item**

&#x2022; `exact_values` - Optional List<br>Exact Values. A list of exact values to match the input against

&#x2022; `regex_values` - Optional List<br>Regex Values. A list of regular expressions to match the input against

&#x2022; `transformers` - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

<a id="api-protection-rules-api-groups-rules"></a>

**API Protection Rules API Groups Rules**

&#x2022; `action` - Optional Block<br>API Protection Rule Action. The action to take if the input request matches the rule<br>See [Action](#api-protection-rules-api-groups-rules-action) below.

&#x2022; `any_domain` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `api_group` - Optional String<br>API Group. API groups derived from API Definition swaggers. For example oas-all-operations including all paths and methods from the swaggers, oas-base-urls covering all requests under base-paths from the swaggers. Custom groups can be created if user tags paths or operations with 'x-volterra-API-group' extensions inside swaggers

&#x2022; `base_path` - Optional String<br>Base Path. Prefix of the request path. For example: /v1

&#x2022; `client_matcher` - Optional Block<br>Client Matcher. Client conditions for matching a rule<br>See [Client Matcher](#api-protection-rules-api-groups-rules-client-matcher) below.

&#x2022; `metadata` - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#api-protection-rules-api-groups-rules-metadata) below.

&#x2022; `request_matcher` - Optional Block<br>Request Matcher. Request conditions for matching a rule<br>See [Request Matcher](#api-protection-rules-api-groups-rules-request-matcher) below.

&#x2022; `specific_domain` - Optional String<br>Specific Domain. The rule will apply for a specific domain. For example: API.example.com

<a id="api-protection-rules-api-groups-rules-action"></a>

**API Protection Rules API Groups Rules Action**

&#x2022; `allow` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `deny` - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-protection-rules-api-groups-rules-client-matcher"></a>

**API Protection Rules API Groups Rules Client Matcher**

&#x2022; `any_client` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `any_ip` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `asn_list` - Optional Block<br>ASN Match List. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer<br>See [Asn List](#api-protection-rules-api-groups-rules-client-matcher-asn-list) below.

&#x2022; `asn_matcher` - Optional Block<br>ASN Matcher. Match any AS number contained in the list of bgp_asn_sets<br>See [Asn Matcher](#api-protection-rules-api-groups-rules-client-matcher-asn-matcher) below.

&#x2022; `client_selector` - Optional Block<br>Label Selector. This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expresssions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects. Label selector is immutable. expressions is a list of strings of label selection expression. Each string has ',' separated values which are 'AND' and all strings are logically 'OR'. BNF for expression string <selector-syntax> ::= <requirement> | <requirement> ',' <selector-syntax> <requirement> ::= [!] KEY [ <set-based-restriction> | <exact-match-restriction> ] <set-based-restriction> ::= '' | <inclusion-exclusion> <value-set> <inclusion-exclusion> ::= <inclusion> | <exclusion> <exclusion> ::= 'notin' <inclusion> ::= 'in' <value-set> ::= '(' <values> ')' <values> ::= VALUE | VALUE ',' <values> <exact-match-restriction> ::= ['='|'=='|'!='] VALUE<br>See [Client Selector](#api-protection-rules-api-groups-rules-client-matcher-client-selector) below.

&#x2022; `ip_matcher` - Optional Block<br>IP Prefix Matcher. Match any IP prefix contained in the list of ip_prefix_sets. The result of the match is inverted if invert_matcher is true<br>See [IP Matcher](#api-protection-rules-api-groups-rules-client-matcher-ip-matcher) below.

&#x2022; `ip_prefix_list` - Optional Block<br>IP Prefix Match List. List of IP Prefix strings to match against<br>See [IP Prefix List](#api-protection-rules-api-groups-rules-client-matcher-ip-prefix-list) below.

&#x2022; `ip_threat_category_list` - Optional Block<br>IP Threat Category List Type. List of IP threat categories<br>See [IP Threat Category List](#api-protection-rules-api-groups-rules-client-matcher-ip-threat-category-list) below.

&#x2022; `tls_fingerprint_matcher` - Optional Block<br>TLS Fingerprint Matcher. A TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of supported positve match criteria includes a list of known classes of TLS fingerprints and a list of exact values. The match is considered successful if either of these positive criteria are satisfied and the input fingerprint is not one of the excluded values<br>See [TLS Fingerprint Matcher](#api-protection-rules-api-groups-rules-client-matcher-tls-fingerprint-matcher) below.

<a id="api-protection-rules-api-groups-rules-client-matcher-asn-list"></a>

**API Protection Rules API Groups Rules Client Matcher Asn List**

&#x2022; `as_numbers` - Optional List<br>AS Numbers. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer

<a id="api-protection-rules-api-groups-rules-client-matcher-asn-matcher"></a>

**API Protection Rules API Groups Rules Client Matcher Asn Matcher**

&#x2022; `asn_sets` - Optional Block<br>BGP ASN Sets. A list of references to bgp_asn_set objects<br>See [Asn Sets](#api-protection-rules-api-groups-rules-client-matcher-asn-matcher-asn-sets) below.

<a id="api-protection-rules-api-groups-rules-client-matcher-asn-matcher-asn-sets"></a>

**API Protection Rules API Groups Rules Client Matcher Asn Matcher Asn Sets**

&#x2022; `kind` - Optional String<br>Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route')

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

&#x2022; `uid` - Optional String<br>UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid

<a id="api-protection-rules-api-groups-rules-client-matcher-client-selector"></a>

**API Protection Rules API Groups Rules Client Matcher Client Selector**

&#x2022; `expressions` - Optional List<br>Selector Expression. expressions contains the kubernetes style label expression for selections

<a id="api-protection-rules-api-groups-rules-client-matcher-ip-matcher"></a>

**API Protection Rules API Groups Rules Client Matcher IP Matcher**

&#x2022; `invert_matcher` - Optional Bool<br>Invert IP Matcher. Invert the match result

&#x2022; `prefix_sets` - Optional Block<br>IP Prefix Sets. A list of references to ip_prefix_set objects<br>See [Prefix Sets](#api-protection-rules-api-groups-rules-client-matcher-ip-matcher-prefix-sets) below.

<a id="api-protection-rules-api-groups-rules-client-matcher-ip-matcher-prefix-sets"></a>

**API Protection Rules API Groups Rules Client Matcher IP Matcher Prefix Sets**

&#x2022; `kind` - Optional String<br>Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route')

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

&#x2022; `uid` - Optional String<br>UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid

<a id="api-protection-rules-api-groups-rules-client-matcher-ip-prefix-list"></a>

**API Protection Rules API Groups Rules Client Matcher IP Prefix List**

&#x2022; `invert_match` - Optional Bool<br>Invert Match Result. Invert the match result

&#x2022; `ip_prefixes` - Optional List<br>IPv4 Prefix List. List of IPv4 prefix strings

<a id="api-protection-rules-api-groups-rules-client-matcher-ip-threat-category-list"></a>

**API Protection Rules API Groups Rules Client Matcher IP Threat Category List**

&#x2022; `ip_threat_categories` - Optional List  Defaults to `SPAM_SOURCES`<br>Possible values are `SPAM_SOURCES`, `WINDOWS_EXPLOITS`, `WEB_ATTACKS`, `BOTNETS`, `SCANNERS`, `REPUTATION`, `PHISHING`, `PROXY`, `MOBILE_THREATS`, `TOR_PROXY`, `DENIAL_OF_SERVICE`, `NETWORK`<br>List of IP Threat Categories to choose. The IP threat categories is obtained from the list and is used to auto-generate equivalent label selection expressions

<a id="api-protection-rules-api-groups-rules-client-matcher-tls-fingerprint-matcher"></a>

**API Protection Rules API Groups Rules Client Matcher TLS Fingerprint Matcher**

&#x2022; `classes` - Optional List  Defaults to `TLS_FINGERPRINT_NONE`<br>Possible values are `TLS_FINGERPRINT_NONE`, `ANY_MALICIOUS_FINGERPRINT`, `ADWARE`, `ADWIND`, `DRIDEX`, `GOOTKIT`, `GOZI`, `JBIFROST`, `QUAKBOT`, `RANSOMWARE`, `TROLDESH`, `TOFSEE`, `TORRENTLOCKER`, `TRICKBOT`<br>TLS fingerprint classes. A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against

&#x2022; `exact_values` - Optional List<br>Exact Values. A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against

&#x2022; `excluded_values` - Optional List<br>Excluded Values. A list of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can be used to skip known false positives when using one or more known TLS fingerprint classes in the enclosing matcher

<a id="api-protection-rules-api-groups-rules-metadata"></a>

**API Protection Rules API Groups Rules Metadata**

&#x2022; `description` - Optional String<br>Description. Human readable description

&#x2022; `name` - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

<a id="api-protection-rules-api-groups-rules-request-matcher"></a>

**API Protection Rules API Groups Rules Request Matcher**

&#x2022; `cookie_matchers` - Optional Block<br>Cookie Matchers. A list of predicates for all cookies that need to be matched. The criteria for matching each cookie is described in individual instances of CookieMatcherType. The actual cookie values are extracted from the request API as a list of strings for each cookie name. Note that all specified cookie matcher predicates must evaluate to true<br>See [Cookie Matchers](#api-protection-rules-api-groups-rules-request-matcher-cookie-matchers) below.

&#x2022; `headers` - Optional Block<br>HTTP Headers. A list of predicates for various HTTP headers that need to match. The criteria for matching each HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values are extracted from the request API as a list of strings for each HTTP header type. Note that all specified header predicates must evaluate to true<br>See [Headers](#api-protection-rules-api-groups-rules-request-matcher-headers) below.

&#x2022; `jwt_claims` - Optional Block<br>JWT Claims. A list of predicates for various JWT claims that need to match. The criteria for matching each JWT claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are extracted from the JWT payload as a list of strings. Note that all specified JWT claim predicates must evaluate to true. Note that this feature only works on LBs with JWT Validation feature enabled<br>See [JWT Claims](#api-protection-rules-api-groups-rules-request-matcher-jwt-claims) below.

&#x2022; `query_params` - Optional Block<br>HTTP Query Parameters. A list of predicates for all query parameters that need to be matched. The criteria for matching each query parameter are described in individual instances of QueryParameterMatcherType. The actual query parameter values are extracted from the request API as a list of strings for each query parameter name. Note that all specified query parameter predicates must evaluate to true<br>See [Query Params](#api-protection-rules-api-groups-rules-request-matcher-query-params) below.

<a id="api-protection-rules-api-groups-rules-request-matcher-cookie-matchers"></a>

**API Protection Rules API Groups Rules Request Matcher Cookie Matchers**

&#x2022; `check_not_present` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `check_present` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `invert_matcher` - Optional Bool<br>Invert Matcher. Invert Match of the expression defined

&#x2022; `item` - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#api-protection-rules-api-groups-rules-request-matcher-cookie-matchers-item) below.

&#x2022; `name` - Optional String<br>Cookie Name. A case-sensitive cookie name

<a id="api-protection-rules-api-groups-rules-request-matcher-cookie-matchers-item"></a>

**API Protection Rules API Groups Rules Request Matcher Cookie Matchers Item**

&#x2022; `exact_values` - Optional List<br>Exact Values. A list of exact values to match the input against

&#x2022; `regex_values` - Optional List<br>Regex Values. A list of regular expressions to match the input against

&#x2022; `transformers` - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

<a id="api-protection-rules-api-groups-rules-request-matcher-headers"></a>

**API Protection Rules API Groups Rules Request Matcher Headers**

&#x2022; `check_not_present` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `check_present` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `invert_matcher` - Optional Bool<br>Invert Header Matcher. Invert the match result

&#x2022; `item` - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#api-protection-rules-api-groups-rules-request-matcher-headers-item) below.

&#x2022; `name` - Optional String<br>Header Name. A case-insensitive HTTP header name

<a id="api-protection-rules-api-groups-rules-request-matcher-headers-item"></a>

**API Protection Rules API Groups Rules Request Matcher Headers Item**

&#x2022; `exact_values` - Optional List<br>Exact Values. A list of exact values to match the input against

&#x2022; `regex_values` - Optional List<br>Regex Values. A list of regular expressions to match the input against

&#x2022; `transformers` - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

<a id="api-protection-rules-api-groups-rules-request-matcher-jwt-claims"></a>

**API Protection Rules API Groups Rules Request Matcher JWT Claims**

&#x2022; `check_not_present` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `check_present` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `invert_matcher` - Optional Bool<br>Invert Matcher. Invert the match result

&#x2022; `item` - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#api-protection-rules-api-groups-rules-request-matcher-jwt-claims-item) below.

&#x2022; `name` - Optional String<br>JWT Claim Name. JWT claim name

<a id="api-protection-rules-api-groups-rules-request-matcher-jwt-claims-item"></a>

**API Protection Rules API Groups Rules Request Matcher JWT Claims Item**

&#x2022; `exact_values` - Optional List<br>Exact Values. A list of exact values to match the input against

&#x2022; `regex_values` - Optional List<br>Regex Values. A list of regular expressions to match the input against

&#x2022; `transformers` - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

<a id="api-protection-rules-api-groups-rules-request-matcher-query-params"></a>

**API Protection Rules API Groups Rules Request Matcher Query Params**

&#x2022; `check_not_present` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `check_present` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `invert_matcher` - Optional Bool<br>Invert Query Parameter Matcher. Invert the match result

&#x2022; `item` - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#api-protection-rules-api-groups-rules-request-matcher-query-params-item) below.

&#x2022; `key` - Optional String<br>Query Parameter Name. A case-sensitive HTTP query parameter name

<a id="api-protection-rules-api-groups-rules-request-matcher-query-params-item"></a>

**API Protection Rules API Groups Rules Request Matcher Query Params Item**

&#x2022; `exact_values` - Optional List<br>Exact Values. A list of exact values to match the input against

&#x2022; `regex_values` - Optional List<br>Regex Values. A list of regular expressions to match the input against

&#x2022; `transformers` - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

<a id="api-rate-limit"></a>

**API Rate Limit**

&#x2022; `api_endpoint_rules` - Optional Block<br>API Endpoints. Sets of rules for a specific endpoints. Order is matter as it uses first match policy. For creating rule that contain a whole domain or group of endpoints, please use the server URL rules above<br>See [API Endpoint Rules](#api-rate-limit-api-endpoint-rules) below.

&#x2022; `bypass_rate_limiting_rules` - Optional Block<br>Bypass Rate Limiting. This category defines rules per URL or API group. If request matches any of these rules, skip Rate Limiting<br>See [Bypass Rate Limiting Rules](#api-rate-limit-bypass-rate-limiting-rules) below.

&#x2022; `custom_ip_allowed_list` - Optional Block<br>Custom IP Allowed List. IP Allowed list using existing ip_prefix_set objects<br>See [Custom IP Allowed List](#api-rate-limit-custom-ip-allowed-list) below.

&#x2022; `ip_allowed_list` - Optional Block<br>IPv4 Prefix List. x-example: '192.168.20.0/24' List of IPv4 prefixes that represent an endpoint<br>See [IP Allowed List](#api-rate-limit-ip-allowed-list) below.

&#x2022; `no_ip_allowed_list` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `server_url_rules` - Optional Block<br>Server URLs. Set of rules for entire domain or base path that contain multiple endpoints. Order is matter as it uses first match policy. For matching also specific endpoints you can use the API endpoint rules set bellow<br>See [Server URL Rules](#api-rate-limit-server-url-rules) below.

<a id="api-rate-limit-api-endpoint-rules"></a>

**API Rate Limit API Endpoint Rules**

&#x2022; `any_domain` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `api_endpoint_method` - Optional Block<br>HTTP Method Matcher. A HTTP method matcher specifies a list of methods to match an input HTTP method. The match is considered successful if the input method is a member of the list. The result of the match based on the method list is inverted if invert_matcher is true<br>See [API Endpoint Method](#api-rate-limit-api-endpoint-rules-api-endpoint-method) below.

&#x2022; `api_endpoint_path` - Optional String<br>API Endpoint. The endpoint (path) of the request

&#x2022; `client_matcher` - Optional Block<br>Client Matcher. Client conditions for matching a rule<br>See [Client Matcher](#api-rate-limit-api-endpoint-rules-client-matcher) below.

&#x2022; `inline_rate_limiter` - Optional Block<br>InlineRateLimiter<br>See [Inline Rate Limiter](#api-rate-limit-api-endpoint-rules-inline-rate-limiter) below.

&#x2022; `ref_rate_limiter` - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Ref Rate Limiter](#api-rate-limit-api-endpoint-rules-ref-rate-limiter) below.

&#x2022; `request_matcher` - Optional Block<br>Request Matcher. Request conditions for matching a rule<br>See [Request Matcher](#api-rate-limit-api-endpoint-rules-request-matcher) below.

&#x2022; `specific_domain` - Optional String<br>Specific Domain. The rule will apply for a specific domain

<a id="api-rate-limit-api-endpoint-rules-api-endpoint-method"></a>

**API Rate Limit API Endpoint Rules API Endpoint Method**

&#x2022; `invert_matcher` - Optional Bool<br>Invert Method Matcher. Invert the match result

&#x2022; `methods` - Optional List  Defaults to `ANY`<br>Possible values are `ANY`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`, `COPY`<br>Method List. List of methods values to match against

<a id="api-rate-limit-api-endpoint-rules-client-matcher"></a>

**API Rate Limit API Endpoint Rules Client Matcher**

&#x2022; `any_client` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `any_ip` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `asn_list` - Optional Block<br>ASN Match List. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer<br>See [Asn List](#api-rate-limit-api-endpoint-rules-client-matcher-asn-list) below.

&#x2022; `asn_matcher` - Optional Block<br>ASN Matcher. Match any AS number contained in the list of bgp_asn_sets<br>See [Asn Matcher](#api-rate-limit-api-endpoint-rules-client-matcher-asn-matcher) below.

&#x2022; `client_selector` - Optional Block<br>Label Selector. This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expresssions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects. Label selector is immutable. expressions is a list of strings of label selection expression. Each string has ',' separated values which are 'AND' and all strings are logically 'OR'. BNF for expression string <selector-syntax> ::= <requirement> | <requirement> ',' <selector-syntax> <requirement> ::= [!] KEY [ <set-based-restriction> | <exact-match-restriction> ] <set-based-restriction> ::= '' | <inclusion-exclusion> <value-set> <inclusion-exclusion> ::= <inclusion> | <exclusion> <exclusion> ::= 'notin' <inclusion> ::= 'in' <value-set> ::= '(' <values> ')' <values> ::= VALUE | VALUE ',' <values> <exact-match-restriction> ::= ['='|'=='|'!='] VALUE<br>See [Client Selector](#api-rate-limit-api-endpoint-rules-client-matcher-client-selector) below.

&#x2022; `ip_matcher` - Optional Block<br>IP Prefix Matcher. Match any IP prefix contained in the list of ip_prefix_sets. The result of the match is inverted if invert_matcher is true<br>See [IP Matcher](#api-rate-limit-api-endpoint-rules-client-matcher-ip-matcher) below.

&#x2022; `ip_prefix_list` - Optional Block<br>IP Prefix Match List. List of IP Prefix strings to match against<br>See [IP Prefix List](#api-rate-limit-api-endpoint-rules-client-matcher-ip-prefix-list) below.

&#x2022; `ip_threat_category_list` - Optional Block<br>IP Threat Category List Type. List of IP threat categories<br>See [IP Threat Category List](#api-rate-limit-api-endpoint-rules-client-matcher-ip-threat-category-list) below.

&#x2022; `tls_fingerprint_matcher` - Optional Block<br>TLS Fingerprint Matcher. A TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of supported positve match criteria includes a list of known classes of TLS fingerprints and a list of exact values. The match is considered successful if either of these positive criteria are satisfied and the input fingerprint is not one of the excluded values<br>See [TLS Fingerprint Matcher](#api-rate-limit-api-endpoint-rules-client-matcher-tls-fingerprint-matcher) below.

<a id="api-rate-limit-api-endpoint-rules-client-matcher-asn-list"></a>

**API Rate Limit API Endpoint Rules Client Matcher Asn List**

&#x2022; `as_numbers` - Optional List<br>AS Numbers. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer

<a id="api-rate-limit-api-endpoint-rules-client-matcher-asn-matcher"></a>

**API Rate Limit API Endpoint Rules Client Matcher Asn Matcher**

&#x2022; `asn_sets` - Optional Block<br>BGP ASN Sets. A list of references to bgp_asn_set objects<br>See [Asn Sets](#api-rate-limit-api-endpoint-rules-client-matcher-asn-matcher-asn-sets) below.

<a id="api-rate-limit-api-endpoint-rules-client-matcher-asn-matcher-asn-sets"></a>

**API Rate Limit API Endpoint Rules Client Matcher Asn Matcher Asn Sets**

&#x2022; `kind` - Optional String<br>Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route')

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

&#x2022; `uid` - Optional String<br>UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid

<a id="api-rate-limit-api-endpoint-rules-client-matcher-client-selector"></a>

**API Rate Limit API Endpoint Rules Client Matcher Client Selector**

&#x2022; `expressions` - Optional List<br>Selector Expression. expressions contains the kubernetes style label expression for selections

<a id="api-rate-limit-api-endpoint-rules-client-matcher-ip-matcher"></a>

**API Rate Limit API Endpoint Rules Client Matcher IP Matcher**

&#x2022; `invert_matcher` - Optional Bool<br>Invert IP Matcher. Invert the match result

&#x2022; `prefix_sets` - Optional Block<br>IP Prefix Sets. A list of references to ip_prefix_set objects<br>See [Prefix Sets](#api-rate-limit-api-endpoint-rules-client-matcher-ip-matcher-prefix-sets) below.

<a id="api-rate-limit-api-endpoint-rules-client-matcher-ip-matcher-prefix-sets"></a>

**API Rate Limit API Endpoint Rules Client Matcher IP Matcher Prefix Sets**

&#x2022; `kind` - Optional String<br>Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route')

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

&#x2022; `uid` - Optional String<br>UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid

<a id="api-rate-limit-api-endpoint-rules-client-matcher-ip-prefix-list"></a>

**API Rate Limit API Endpoint Rules Client Matcher IP Prefix List**

&#x2022; `invert_match` - Optional Bool<br>Invert Match Result. Invert the match result

&#x2022; `ip_prefixes` - Optional List<br>IPv4 Prefix List. List of IPv4 prefix strings

<a id="api-rate-limit-api-endpoint-rules-client-matcher-ip-threat-category-list"></a>

**API Rate Limit API Endpoint Rules Client Matcher IP Threat Category List**

&#x2022; `ip_threat_categories` - Optional List  Defaults to `SPAM_SOURCES`<br>Possible values are `SPAM_SOURCES`, `WINDOWS_EXPLOITS`, `WEB_ATTACKS`, `BOTNETS`, `SCANNERS`, `REPUTATION`, `PHISHING`, `PROXY`, `MOBILE_THREATS`, `TOR_PROXY`, `DENIAL_OF_SERVICE`, `NETWORK`<br>List of IP Threat Categories to choose. The IP threat categories is obtained from the list and is used to auto-generate equivalent label selection expressions

<a id="api-rate-limit-api-endpoint-rules-client-matcher-tls-fingerprint-matcher"></a>

**API Rate Limit API Endpoint Rules Client Matcher TLS Fingerprint Matcher**

&#x2022; `classes` - Optional List  Defaults to `TLS_FINGERPRINT_NONE`<br>Possible values are `TLS_FINGERPRINT_NONE`, `ANY_MALICIOUS_FINGERPRINT`, `ADWARE`, `ADWIND`, `DRIDEX`, `GOOTKIT`, `GOZI`, `JBIFROST`, `QUAKBOT`, `RANSOMWARE`, `TROLDESH`, `TOFSEE`, `TORRENTLOCKER`, `TRICKBOT`<br>TLS fingerprint classes. A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against

&#x2022; `exact_values` - Optional List<br>Exact Values. A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against

&#x2022; `excluded_values` - Optional List<br>Excluded Values. A list of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can be used to skip known false positives when using one or more known TLS fingerprint classes in the enclosing matcher

<a id="api-rate-limit-api-endpoint-rules-inline-rate-limiter"></a>

**API Rate Limit API Endpoint Rules Inline Rate Limiter**

&#x2022; `ref_user_id` - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Ref User Id](#api-rate-limit-api-endpoint-rules-inline-rate-limiter-ref-user-id) below.

&#x2022; `threshold` - Optional Number<br>Threshold. The total number of allowed requests for 1 unit (e.g. SECOND/MINUTE/HOUR etc.) of the specified period

&#x2022; `unit` - Optional String  Defaults to `SECOND`<br>Possible values are `SECOND`, `MINUTE`, `HOUR`<br>Rate Limit Period Unit. Unit for the period per which the rate limit is applied. - SECOND: Second Rate limit period unit is seconds - MINUTE: Minute Rate limit period unit is minutes - HOUR: Hour Rate limit period unit is hours - DAY: Day Rate limit period unit is days

&#x2022; `use_http_lb_user_id` - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-rate-limit-api-endpoint-rules-inline-rate-limiter-ref-user-id"></a>

**API Rate Limit API Endpoint Rules Inline Rate Limiter Ref User Id**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="api-rate-limit-api-endpoint-rules-ref-rate-limiter"></a>

**API Rate Limit API Endpoint Rules Ref Rate Limiter**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="api-rate-limit-api-endpoint-rules-request-matcher"></a>

**API Rate Limit API Endpoint Rules Request Matcher**

&#x2022; `cookie_matchers` - Optional Block<br>Cookie Matchers. A list of predicates for all cookies that need to be matched. The criteria for matching each cookie is described in individual instances of CookieMatcherType. The actual cookie values are extracted from the request API as a list of strings for each cookie name. Note that all specified cookie matcher predicates must evaluate to true<br>See [Cookie Matchers](#api-rate-limit-api-endpoint-rules-request-matcher-cookie-matchers) below.

&#x2022; `headers` - Optional Block<br>HTTP Headers. A list of predicates for various HTTP headers that need to match. The criteria for matching each HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values are extracted from the request API as a list of strings for each HTTP header type. Note that all specified header predicates must evaluate to true<br>See [Headers](#api-rate-limit-api-endpoint-rules-request-matcher-headers) below.

&#x2022; `jwt_claims` - Optional Block<br>JWT Claims. A list of predicates for various JWT claims that need to match. The criteria for matching each JWT claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are extracted from the JWT payload as a list of strings. Note that all specified JWT claim predicates must evaluate to true. Note that this feature only works on LBs with JWT Validation feature enabled<br>See [JWT Claims](#api-rate-limit-api-endpoint-rules-request-matcher-jwt-claims) below.

&#x2022; `query_params` - Optional Block<br>HTTP Query Parameters. A list of predicates for all query parameters that need to be matched. The criteria for matching each query parameter are described in individual instances of QueryParameterMatcherType. The actual query parameter values are extracted from the request API as a list of strings for each query parameter name. Note that all specified query parameter predicates must evaluate to true<br>See [Query Params](#api-rate-limit-api-endpoint-rules-request-matcher-query-params) below.

<a id="api-rate-limit-api-endpoint-rules-request-matcher-cookie-matchers"></a>

**API Rate Limit API Endpoint Rules Request Matcher Cookie Matchers**

&#x2022; `check_not_present` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `check_present` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `invert_matcher` - Optional Bool<br>Invert Matcher. Invert Match of the expression defined

&#x2022; `item` - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#api-rate-limit-api-endpoint-rules-request-matcher-cookie-matchers-item) below.

&#x2022; `name` - Optional String<br>Cookie Name. A case-sensitive cookie name

<a id="api-rate-limit-api-endpoint-rules-request-matcher-cookie-matchers-item"></a>

**API Rate Limit API Endpoint Rules Request Matcher Cookie Matchers Item**

&#x2022; `exact_values` - Optional List<br>Exact Values. A list of exact values to match the input against

&#x2022; `regex_values` - Optional List<br>Regex Values. A list of regular expressions to match the input against

&#x2022; `transformers` - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

<a id="api-rate-limit-api-endpoint-rules-request-matcher-headers"></a>

**API Rate Limit API Endpoint Rules Request Matcher Headers**

&#x2022; `check_not_present` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `check_present` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `invert_matcher` - Optional Bool<br>Invert Header Matcher. Invert the match result

&#x2022; `item` - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#api-rate-limit-api-endpoint-rules-request-matcher-headers-item) below.

&#x2022; `name` - Optional String<br>Header Name. A case-insensitive HTTP header name

<a id="api-rate-limit-api-endpoint-rules-request-matcher-headers-item"></a>

**API Rate Limit API Endpoint Rules Request Matcher Headers Item**

&#x2022; `exact_values` - Optional List<br>Exact Values. A list of exact values to match the input against

&#x2022; `regex_values` - Optional List<br>Regex Values. A list of regular expressions to match the input against

&#x2022; `transformers` - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

<a id="api-rate-limit-api-endpoint-rules-request-matcher-jwt-claims"></a>

**API Rate Limit API Endpoint Rules Request Matcher JWT Claims**

&#x2022; `check_not_present` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `check_present` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `invert_matcher` - Optional Bool<br>Invert Matcher. Invert the match result

&#x2022; `item` - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#api-rate-limit-api-endpoint-rules-request-matcher-jwt-claims-item) below.

&#x2022; `name` - Optional String<br>JWT Claim Name. JWT claim name

<a id="api-rate-limit-api-endpoint-rules-request-matcher-jwt-claims-item"></a>

**API Rate Limit API Endpoint Rules Request Matcher JWT Claims Item**

&#x2022; `exact_values` - Optional List<br>Exact Values. A list of exact values to match the input against

&#x2022; `regex_values` - Optional List<br>Regex Values. A list of regular expressions to match the input against

&#x2022; `transformers` - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

<a id="api-rate-limit-api-endpoint-rules-request-matcher-query-params"></a>

**API Rate Limit API Endpoint Rules Request Matcher Query Params**

&#x2022; `check_not_present` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `check_present` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `invert_matcher` - Optional Bool<br>Invert Query Parameter Matcher. Invert the match result

&#x2022; `item` - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#api-rate-limit-api-endpoint-rules-request-matcher-query-params-item) below.

&#x2022; `key` - Optional String<br>Query Parameter Name. A case-sensitive HTTP query parameter name

<a id="api-rate-limit-api-endpoint-rules-request-matcher-query-params-item"></a>

**API Rate Limit API Endpoint Rules Request Matcher Query Params Item**

&#x2022; `exact_values` - Optional List<br>Exact Values. A list of exact values to match the input against

&#x2022; `regex_values` - Optional List<br>Regex Values. A list of regular expressions to match the input against

&#x2022; `transformers` - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

<a id="api-rate-limit-bypass-rate-limiting-rules"></a>

**API Rate Limit Bypass Rate Limiting Rules**

&#x2022; `bypass_rate_limiting_rules` - Optional Block<br>Bypass Rate Limiting. This category defines rules per URL or API group. If request matches any of these rules, skip Rate Limiting<br>See [Bypass Rate Limiting Rules](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules) below.

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules"></a>

**API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules**

&#x2022; `any_domain` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `any_url` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `api_endpoint` - Optional Block<br>API Endpoint. This defines API endpoint<br>See [API Endpoint](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-api-endpoint) below.

&#x2022; `api_groups` - Optional Block<br>API Groups<br>See [API Groups](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-api-groups) below.

&#x2022; `base_path` - Optional String<br>Base Path. The base path which this validation applies to

&#x2022; `client_matcher` - Optional Block<br>Client Matcher. Client conditions for matching a rule<br>See [Client Matcher](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher) below.

&#x2022; `request_matcher` - Optional Block<br>Request Matcher. Request conditions for matching a rule<br>See [Request Matcher](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher) below.

&#x2022; `specific_domain` - Optional String<br>Specific Domain. The rule will apply for a specific domain. For example: API.example.com

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-api-endpoint"></a>

**API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules API Endpoint**

&#x2022; `methods` - Optional List  Defaults to `ANY`<br>Possible values are `ANY`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`, `COPY`<br>Methods. Methods to be matched

&#x2022; `path` - Optional String<br>Path. Path to be matched

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-api-groups"></a>

**API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules API Groups**

&#x2022; `api_groups` - Optional List<br>API Groups

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher"></a>

**API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Client Matcher**

&#x2022; `any_client` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `any_ip` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `asn_list` - Optional Block<br>ASN Match List. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer<br>See [Asn List](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-asn-list) below.

&#x2022; `asn_matcher` - Optional Block<br>ASN Matcher. Match any AS number contained in the list of bgp_asn_sets<br>See [Asn Matcher](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-asn-matcher) below.

&#x2022; `client_selector` - Optional Block<br>Label Selector. This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expresssions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects. Label selector is immutable. expressions is a list of strings of label selection expression. Each string has ',' separated values which are 'AND' and all strings are logically 'OR'. BNF for expression string <selector-syntax> ::= <requirement> | <requirement> ',' <selector-syntax> <requirement> ::= [!] KEY [ <set-based-restriction> | <exact-match-restriction> ] <set-based-restriction> ::= '' | <inclusion-exclusion> <value-set> <inclusion-exclusion> ::= <inclusion> | <exclusion> <exclusion> ::= 'notin' <inclusion> ::= 'in' <value-set> ::= '(' <values> ')' <values> ::= VALUE | VALUE ',' <values> <exact-match-restriction> ::= ['='|'=='|'!='] VALUE<br>See [Client Selector](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-client-selector) below.

&#x2022; `ip_matcher` - Optional Block<br>IP Prefix Matcher. Match any IP prefix contained in the list of ip_prefix_sets. The result of the match is inverted if invert_matcher is true<br>See [IP Matcher](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-ip-matcher) below.

&#x2022; `ip_prefix_list` - Optional Block<br>IP Prefix Match List. List of IP Prefix strings to match against<br>See [IP Prefix List](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-ip-prefix-list) below.

&#x2022; `ip_threat_category_list` - Optional Block<br>IP Threat Category List Type. List of IP threat categories<br>See [IP Threat Category List](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-ip-threat-category-list) below.

&#x2022; `tls_fingerprint_matcher` - Optional Block<br>TLS Fingerprint Matcher. A TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of supported positve match criteria includes a list of known classes of TLS fingerprints and a list of exact values. The match is considered successful if either of these positive criteria are satisfied and the input fingerprint is not one of the excluded values<br>See [TLS Fingerprint Matcher](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-tls-fingerprint-matcher) below.

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-asn-list"></a>

**API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Client Matcher Asn List**

&#x2022; `as_numbers` - Optional List<br>AS Numbers. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-asn-matcher"></a>

**API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Client Matcher Asn Matcher**

&#x2022; `asn_sets` - Optional Block<br>BGP ASN Sets. A list of references to bgp_asn_set objects<br>See [Asn Sets](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-asn-matcher-asn-sets) below.

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-asn-matcher-asn-sets"></a>

**API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Client Matcher Asn Matcher Asn Sets**

&#x2022; `kind` - Optional String<br>Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route')

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

&#x2022; `uid` - Optional String<br>UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-client-selector"></a>

**API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Client Matcher Client Selector**

&#x2022; `expressions` - Optional List<br>Selector Expression. expressions contains the kubernetes style label expression for selections

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-ip-matcher"></a>

**API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Client Matcher IP Matcher**

&#x2022; `invert_matcher` - Optional Bool<br>Invert IP Matcher. Invert the match result

&#x2022; `prefix_sets` - Optional Block<br>IP Prefix Sets. A list of references to ip_prefix_set objects<br>See [Prefix Sets](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-ip-matcher-prefix-sets) below.

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-ip-matcher-prefix-sets"></a>

**API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Client Matcher IP Matcher Prefix Sets**

&#x2022; `kind` - Optional String<br>Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route')

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

&#x2022; `uid` - Optional String<br>UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-ip-prefix-list"></a>

**API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Client Matcher IP Prefix List**

&#x2022; `invert_match` - Optional Bool<br>Invert Match Result. Invert the match result

&#x2022; `ip_prefixes` - Optional List<br>IPv4 Prefix List. List of IPv4 prefix strings

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-ip-threat-category-list"></a>

**API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Client Matcher IP Threat Category List**

&#x2022; `ip_threat_categories` - Optional List  Defaults to `SPAM_SOURCES`<br>Possible values are `SPAM_SOURCES`, `WINDOWS_EXPLOITS`, `WEB_ATTACKS`, `BOTNETS`, `SCANNERS`, `REPUTATION`, `PHISHING`, `PROXY`, `MOBILE_THREATS`, `TOR_PROXY`, `DENIAL_OF_SERVICE`, `NETWORK`<br>List of IP Threat Categories to choose. The IP threat categories is obtained from the list and is used to auto-generate equivalent label selection expressions

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-client-matcher-tls-fingerprint-matcher"></a>

**API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Client Matcher TLS Fingerprint Matcher**

&#x2022; `classes` - Optional List  Defaults to `TLS_FINGERPRINT_NONE`<br>Possible values are `TLS_FINGERPRINT_NONE`, `ANY_MALICIOUS_FINGERPRINT`, `ADWARE`, `ADWIND`, `DRIDEX`, `GOOTKIT`, `GOZI`, `JBIFROST`, `QUAKBOT`, `RANSOMWARE`, `TROLDESH`, `TOFSEE`, `TORRENTLOCKER`, `TRICKBOT`<br>TLS fingerprint classes. A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against

&#x2022; `exact_values` - Optional List<br>Exact Values. A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against

&#x2022; `excluded_values` - Optional List<br>Excluded Values. A list of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can be used to skip known false positives when using one or more known TLS fingerprint classes in the enclosing matcher

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher"></a>

**API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Request Matcher**

&#x2022; `cookie_matchers` - Optional Block<br>Cookie Matchers. A list of predicates for all cookies that need to be matched. The criteria for matching each cookie is described in individual instances of CookieMatcherType. The actual cookie values are extracted from the request API as a list of strings for each cookie name. Note that all specified cookie matcher predicates must evaluate to true<br>See [Cookie Matchers](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-cookie-matchers) below.

&#x2022; `headers` - Optional Block<br>HTTP Headers. A list of predicates for various HTTP headers that need to match. The criteria for matching each HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values are extracted from the request API as a list of strings for each HTTP header type. Note that all specified header predicates must evaluate to true<br>See [Headers](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-headers) below.

&#x2022; `jwt_claims` - Optional Block<br>JWT Claims. A list of predicates for various JWT claims that need to match. The criteria for matching each JWT claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are extracted from the JWT payload as a list of strings. Note that all specified JWT claim predicates must evaluate to true. Note that this feature only works on LBs with JWT Validation feature enabled<br>See [JWT Claims](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-jwt-claims) below.

&#x2022; `query_params` - Optional Block<br>HTTP Query Parameters. A list of predicates for all query parameters that need to be matched. The criteria for matching each query parameter are described in individual instances of QueryParameterMatcherType. The actual query parameter values are extracted from the request API as a list of strings for each query parameter name. Note that all specified query parameter predicates must evaluate to true<br>See [Query Params](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-query-params) below.

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-cookie-matchers"></a>

**API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Request Matcher Cookie Matchers**

&#x2022; `check_not_present` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `check_present` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `invert_matcher` - Optional Bool<br>Invert Matcher. Invert Match of the expression defined

&#x2022; `item` - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-cookie-matchers-item) below.

&#x2022; `name` - Optional String<br>Cookie Name. A case-sensitive cookie name

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-cookie-matchers-item"></a>

**API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Request Matcher Cookie Matchers Item**

&#x2022; `exact_values` - Optional List<br>Exact Values. A list of exact values to match the input against

&#x2022; `regex_values` - Optional List<br>Regex Values. A list of regular expressions to match the input against

&#x2022; `transformers` - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-headers"></a>

**API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Request Matcher Headers**

&#x2022; `check_not_present` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `check_present` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `invert_matcher` - Optional Bool<br>Invert Header Matcher. Invert the match result

&#x2022; `item` - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-headers-item) below.

&#x2022; `name` - Optional String<br>Header Name. A case-insensitive HTTP header name

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-headers-item"></a>

**API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Request Matcher Headers Item**

&#x2022; `exact_values` - Optional List<br>Exact Values. A list of exact values to match the input against

&#x2022; `regex_values` - Optional List<br>Regex Values. A list of regular expressions to match the input against

&#x2022; `transformers` - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-jwt-claims"></a>

**API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Request Matcher JWT Claims**

&#x2022; `check_not_present` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `check_present` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `invert_matcher` - Optional Bool<br>Invert Matcher. Invert the match result

&#x2022; `item` - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-jwt-claims-item) below.

&#x2022; `name` - Optional String<br>JWT Claim Name. JWT claim name

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-jwt-claims-item"></a>

**API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Request Matcher JWT Claims Item**

&#x2022; `exact_values` - Optional List<br>Exact Values. A list of exact values to match the input against

&#x2022; `regex_values` - Optional List<br>Regex Values. A list of regular expressions to match the input against

&#x2022; `transformers` - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-query-params"></a>

**API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Request Matcher Query Params**

&#x2022; `check_not_present` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `check_present` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `invert_matcher` - Optional Bool<br>Invert Query Parameter Matcher. Invert the match result

&#x2022; `item` - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-query-params-item) below.

&#x2022; `key` - Optional String<br>Query Parameter Name. A case-sensitive HTTP query parameter name

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules-request-matcher-query-params-item"></a>

**API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules Request Matcher Query Params Item**

&#x2022; `exact_values` - Optional List<br>Exact Values. A list of exact values to match the input against

&#x2022; `regex_values` - Optional List<br>Regex Values. A list of regular expressions to match the input against

&#x2022; `transformers` - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

<a id="api-rate-limit-custom-ip-allowed-list"></a>

**API Rate Limit Custom IP Allowed List**

&#x2022; `rate_limiter_allowed_prefixes` - Optional Block<br>List of IP Prefix Sets. References to ip_prefix_set objects. Requests from source IP addresses that are covered by one of the allowed IP Prefixes are not subjected to rate limiting<br>See [Rate Limiter Allowed Prefixes](#api-rate-limit-custom-ip-allowed-list-rate-limiter-allowed-prefixes) below.

<a id="api-rate-limit-custom-ip-allowed-list-rate-limiter-allowed-prefixes"></a>

**API Rate Limit Custom IP Allowed List Rate Limiter Allowed Prefixes**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="api-rate-limit-ip-allowed-list"></a>

**API Rate Limit IP Allowed List**

&#x2022; `prefixes` - Optional List<br>IPv4 Prefix List. List of IPv4 prefixes that represent an endpoint

<a id="api-rate-limit-server-url-rules"></a>

**API Rate Limit Server URL Rules**

&#x2022; `any_domain` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `api_group` - Optional String<br>API Group. API groups derived from API Definition swaggers. For example oas-all-operations including all paths and methods from the swaggers, oas-base-urls covering all requests under base-paths from the swaggers. Custom groups can be created if user tags paths or operations with 'x-volterra-API-group' extensions inside swaggers

&#x2022; `base_path` - Optional String<br>Base Path. Prefix of the request path

&#x2022; `client_matcher` - Optional Block<br>Client Matcher. Client conditions for matching a rule<br>See [Client Matcher](#api-rate-limit-server-url-rules-client-matcher) below.

&#x2022; `inline_rate_limiter` - Optional Block<br>InlineRateLimiter<br>See [Inline Rate Limiter](#api-rate-limit-server-url-rules-inline-rate-limiter) below.

&#x2022; `ref_rate_limiter` - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Ref Rate Limiter](#api-rate-limit-server-url-rules-ref-rate-limiter) below.

&#x2022; `request_matcher` - Optional Block<br>Request Matcher. Request conditions for matching a rule<br>See [Request Matcher](#api-rate-limit-server-url-rules-request-matcher) below.

&#x2022; `specific_domain` - Optional String<br>Specific Domain. The rule will apply for a specific domain

<a id="api-rate-limit-server-url-rules-client-matcher"></a>

**API Rate Limit Server URL Rules Client Matcher**

&#x2022; `any_client` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `any_ip` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `asn_list` - Optional Block<br>ASN Match List. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer<br>See [Asn List](#api-rate-limit-server-url-rules-client-matcher-asn-list) below.

&#x2022; `asn_matcher` - Optional Block<br>ASN Matcher. Match any AS number contained in the list of bgp_asn_sets<br>See [Asn Matcher](#api-rate-limit-server-url-rules-client-matcher-asn-matcher) below.

&#x2022; `client_selector` - Optional Block<br>Label Selector. This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expresssions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects. Label selector is immutable. expressions is a list of strings of label selection expression. Each string has ',' separated values which are 'AND' and all strings are logically 'OR'. BNF for expression string <selector-syntax> ::= <requirement> | <requirement> ',' <selector-syntax> <requirement> ::= [!] KEY [ <set-based-restriction> | <exact-match-restriction> ] <set-based-restriction> ::= '' | <inclusion-exclusion> <value-set> <inclusion-exclusion> ::= <inclusion> | <exclusion> <exclusion> ::= 'notin' <inclusion> ::= 'in' <value-set> ::= '(' <values> ')' <values> ::= VALUE | VALUE ',' <values> <exact-match-restriction> ::= ['='|'=='|'!='] VALUE<br>See [Client Selector](#api-rate-limit-server-url-rules-client-matcher-client-selector) below.

&#x2022; `ip_matcher` - Optional Block<br>IP Prefix Matcher. Match any IP prefix contained in the list of ip_prefix_sets. The result of the match is inverted if invert_matcher is true<br>See [IP Matcher](#api-rate-limit-server-url-rules-client-matcher-ip-matcher) below.

&#x2022; `ip_prefix_list` - Optional Block<br>IP Prefix Match List. List of IP Prefix strings to match against<br>See [IP Prefix List](#api-rate-limit-server-url-rules-client-matcher-ip-prefix-list) below.

&#x2022; `ip_threat_category_list` - Optional Block<br>IP Threat Category List Type. List of IP threat categories<br>See [IP Threat Category List](#api-rate-limit-server-url-rules-client-matcher-ip-threat-category-list) below.

&#x2022; `tls_fingerprint_matcher` - Optional Block<br>TLS Fingerprint Matcher. A TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of supported positve match criteria includes a list of known classes of TLS fingerprints and a list of exact values. The match is considered successful if either of these positive criteria are satisfied and the input fingerprint is not one of the excluded values<br>See [TLS Fingerprint Matcher](#api-rate-limit-server-url-rules-client-matcher-tls-fingerprint-matcher) below.

<a id="api-rate-limit-server-url-rules-client-matcher-asn-list"></a>

**API Rate Limit Server URL Rules Client Matcher Asn List**

&#x2022; `as_numbers` - Optional List<br>AS Numbers. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer

<a id="api-rate-limit-server-url-rules-client-matcher-asn-matcher"></a>

**API Rate Limit Server URL Rules Client Matcher Asn Matcher**

&#x2022; `asn_sets` - Optional Block<br>BGP ASN Sets. A list of references to bgp_asn_set objects<br>See [Asn Sets](#api-rate-limit-server-url-rules-client-matcher-asn-matcher-asn-sets) below.

<a id="api-rate-limit-server-url-rules-client-matcher-asn-matcher-asn-sets"></a>

**API Rate Limit Server URL Rules Client Matcher Asn Matcher Asn Sets**

&#x2022; `kind` - Optional String<br>Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route')

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

&#x2022; `uid` - Optional String<br>UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid

<a id="api-rate-limit-server-url-rules-client-matcher-client-selector"></a>

**API Rate Limit Server URL Rules Client Matcher Client Selector**

&#x2022; `expressions` - Optional List<br>Selector Expression. expressions contains the kubernetes style label expression for selections

<a id="api-rate-limit-server-url-rules-client-matcher-ip-matcher"></a>

**API Rate Limit Server URL Rules Client Matcher IP Matcher**

&#x2022; `invert_matcher` - Optional Bool<br>Invert IP Matcher. Invert the match result

&#x2022; `prefix_sets` - Optional Block<br>IP Prefix Sets. A list of references to ip_prefix_set objects<br>See [Prefix Sets](#api-rate-limit-server-url-rules-client-matcher-ip-matcher-prefix-sets) below.

<a id="api-rate-limit-server-url-rules-client-matcher-ip-matcher-prefix-sets"></a>

**API Rate Limit Server URL Rules Client Matcher IP Matcher Prefix Sets**

&#x2022; `kind` - Optional String<br>Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route')

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

&#x2022; `uid` - Optional String<br>UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid

<a id="api-rate-limit-server-url-rules-client-matcher-ip-prefix-list"></a>

**API Rate Limit Server URL Rules Client Matcher IP Prefix List**

&#x2022; `invert_match` - Optional Bool<br>Invert Match Result. Invert the match result

&#x2022; `ip_prefixes` - Optional List<br>IPv4 Prefix List. List of IPv4 prefix strings

<a id="api-rate-limit-server-url-rules-client-matcher-ip-threat-category-list"></a>

**API Rate Limit Server URL Rules Client Matcher IP Threat Category List**

&#x2022; `ip_threat_categories` - Optional List  Defaults to `SPAM_SOURCES`<br>Possible values are `SPAM_SOURCES`, `WINDOWS_EXPLOITS`, `WEB_ATTACKS`, `BOTNETS`, `SCANNERS`, `REPUTATION`, `PHISHING`, `PROXY`, `MOBILE_THREATS`, `TOR_PROXY`, `DENIAL_OF_SERVICE`, `NETWORK`<br>List of IP Threat Categories to choose. The IP threat categories is obtained from the list and is used to auto-generate equivalent label selection expressions

<a id="api-rate-limit-server-url-rules-client-matcher-tls-fingerprint-matcher"></a>

**API Rate Limit Server URL Rules Client Matcher TLS Fingerprint Matcher**

&#x2022; `classes` - Optional List  Defaults to `TLS_FINGERPRINT_NONE`<br>Possible values are `TLS_FINGERPRINT_NONE`, `ANY_MALICIOUS_FINGERPRINT`, `ADWARE`, `ADWIND`, `DRIDEX`, `GOOTKIT`, `GOZI`, `JBIFROST`, `QUAKBOT`, `RANSOMWARE`, `TROLDESH`, `TOFSEE`, `TORRENTLOCKER`, `TRICKBOT`<br>TLS fingerprint classes. A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against

&#x2022; `exact_values` - Optional List<br>Exact Values. A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against

&#x2022; `excluded_values` - Optional List<br>Excluded Values. A list of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can be used to skip known false positives when using one or more known TLS fingerprint classes in the enclosing matcher

<a id="api-rate-limit-server-url-rules-inline-rate-limiter"></a>

**API Rate Limit Server URL Rules Inline Rate Limiter**

&#x2022; `ref_user_id` - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Ref User Id](#api-rate-limit-server-url-rules-inline-rate-limiter-ref-user-id) below.

&#x2022; `threshold` - Optional Number<br>Threshold. The total number of allowed requests for 1 unit (e.g. SECOND/MINUTE/HOUR etc.) of the specified period

&#x2022; `unit` - Optional String  Defaults to `SECOND`<br>Possible values are `SECOND`, `MINUTE`, `HOUR`<br>Rate Limit Period Unit. Unit for the period per which the rate limit is applied. - SECOND: Second Rate limit period unit is seconds - MINUTE: Minute Rate limit period unit is minutes - HOUR: Hour Rate limit period unit is hours - DAY: Day Rate limit period unit is days

&#x2022; `use_http_lb_user_id` - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-rate-limit-server-url-rules-inline-rate-limiter-ref-user-id"></a>

**API Rate Limit Server URL Rules Inline Rate Limiter Ref User Id**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="api-rate-limit-server-url-rules-ref-rate-limiter"></a>

**API Rate Limit Server URL Rules Ref Rate Limiter**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="api-rate-limit-server-url-rules-request-matcher"></a>

**API Rate Limit Server URL Rules Request Matcher**

&#x2022; `cookie_matchers` - Optional Block<br>Cookie Matchers. A list of predicates for all cookies that need to be matched. The criteria for matching each cookie is described in individual instances of CookieMatcherType. The actual cookie values are extracted from the request API as a list of strings for each cookie name. Note that all specified cookie matcher predicates must evaluate to true<br>See [Cookie Matchers](#api-rate-limit-server-url-rules-request-matcher-cookie-matchers) below.

&#x2022; `headers` - Optional Block<br>HTTP Headers. A list of predicates for various HTTP headers that need to match. The criteria for matching each HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values are extracted from the request API as a list of strings for each HTTP header type. Note that all specified header predicates must evaluate to true<br>See [Headers](#api-rate-limit-server-url-rules-request-matcher-headers) below.

&#x2022; `jwt_claims` - Optional Block<br>JWT Claims. A list of predicates for various JWT claims that need to match. The criteria for matching each JWT claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are extracted from the JWT payload as a list of strings. Note that all specified JWT claim predicates must evaluate to true. Note that this feature only works on LBs with JWT Validation feature enabled<br>See [JWT Claims](#api-rate-limit-server-url-rules-request-matcher-jwt-claims) below.

&#x2022; `query_params` - Optional Block<br>HTTP Query Parameters. A list of predicates for all query parameters that need to be matched. The criteria for matching each query parameter are described in individual instances of QueryParameterMatcherType. The actual query parameter values are extracted from the request API as a list of strings for each query parameter name. Note that all specified query parameter predicates must evaluate to true<br>See [Query Params](#api-rate-limit-server-url-rules-request-matcher-query-params) below.

<a id="api-rate-limit-server-url-rules-request-matcher-cookie-matchers"></a>

**API Rate Limit Server URL Rules Request Matcher Cookie Matchers**

&#x2022; `check_not_present` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `check_present` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `invert_matcher` - Optional Bool<br>Invert Matcher. Invert Match of the expression defined

&#x2022; `item` - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#api-rate-limit-server-url-rules-request-matcher-cookie-matchers-item) below.

&#x2022; `name` - Optional String<br>Cookie Name. A case-sensitive cookie name

<a id="api-rate-limit-server-url-rules-request-matcher-cookie-matchers-item"></a>

**API Rate Limit Server URL Rules Request Matcher Cookie Matchers Item**

&#x2022; `exact_values` - Optional List<br>Exact Values. A list of exact values to match the input against

&#x2022; `regex_values` - Optional List<br>Regex Values. A list of regular expressions to match the input against

&#x2022; `transformers` - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

<a id="api-rate-limit-server-url-rules-request-matcher-headers"></a>

**API Rate Limit Server URL Rules Request Matcher Headers**

&#x2022; `check_not_present` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `check_present` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `invert_matcher` - Optional Bool<br>Invert Header Matcher. Invert the match result

&#x2022; `item` - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#api-rate-limit-server-url-rules-request-matcher-headers-item) below.

&#x2022; `name` - Optional String<br>Header Name. A case-insensitive HTTP header name

<a id="api-rate-limit-server-url-rules-request-matcher-headers-item"></a>

**API Rate Limit Server URL Rules Request Matcher Headers Item**

&#x2022; `exact_values` - Optional List<br>Exact Values. A list of exact values to match the input against

&#x2022; `regex_values` - Optional List<br>Regex Values. A list of regular expressions to match the input against

&#x2022; `transformers` - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

<a id="api-rate-limit-server-url-rules-request-matcher-jwt-claims"></a>

**API Rate Limit Server URL Rules Request Matcher JWT Claims**

&#x2022; `check_not_present` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `check_present` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `invert_matcher` - Optional Bool<br>Invert Matcher. Invert the match result

&#x2022; `item` - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#api-rate-limit-server-url-rules-request-matcher-jwt-claims-item) below.

&#x2022; `name` - Optional String<br>JWT Claim Name. JWT claim name

<a id="api-rate-limit-server-url-rules-request-matcher-jwt-claims-item"></a>

**API Rate Limit Server URL Rules Request Matcher JWT Claims Item**

&#x2022; `exact_values` - Optional List<br>Exact Values. A list of exact values to match the input against

&#x2022; `regex_values` - Optional List<br>Regex Values. A list of regular expressions to match the input against

&#x2022; `transformers` - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

<a id="api-rate-limit-server-url-rules-request-matcher-query-params"></a>

**API Rate Limit Server URL Rules Request Matcher Query Params**

&#x2022; `check_not_present` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `check_present` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `invert_matcher` - Optional Bool<br>Invert Query Parameter Matcher. Invert the match result

&#x2022; `item` - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#api-rate-limit-server-url-rules-request-matcher-query-params-item) below.

&#x2022; `key` - Optional String<br>Query Parameter Name. A case-sensitive HTTP query parameter name

<a id="api-rate-limit-server-url-rules-request-matcher-query-params-item"></a>

**API Rate Limit Server URL Rules Request Matcher Query Params Item**

&#x2022; `exact_values` - Optional List<br>Exact Values. A list of exact values to match the input against

&#x2022; `regex_values` - Optional List<br>Regex Values. A list of regular expressions to match the input against

&#x2022; `transformers` - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

<a id="api-specification"></a>

**API Specification**

&#x2022; `api_definition` - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [API Definition](#api-specification-api-definition) below.

&#x2022; `validation_all_spec_endpoints` - Optional Block<br>API Inventory. Settings for API Inventory validation<br>See [Validation All Spec Endpoints](#api-specification-validation-all-spec-endpoints) below.

&#x2022; `validation_custom_list` - Optional Block<br>Custom List. Define API groups, base paths, or API endpoints and their OpenAPI validation modes. Any other API-endpoint not listed will act according to 'Fall Through Mode'<br>See [Validation Custom List](#api-specification-validation-custom-list) below.

&#x2022; `validation_disabled` - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-specification-api-definition"></a>

**API Specification API Definition**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="api-specification-validation-all-spec-endpoints"></a>

**API Specification Validation All Spec Endpoints**

&#x2022; `fall_through_mode` - Optional Block<br>Fall Through Mode. x-required Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a. swagger) or doesn't have a specific rule in custom rules)<br>See [Fall Through Mode](#api-specification-validation-all-spec-endpoints-fall-through-mode) below.

&#x2022; `settings` - Optional Block<br>Common Settings. OpenAPI specification validation settings relevant for 'API Inventory' enforcement and for 'Custom list' enforcement<br>See [Settings](#api-specification-validation-all-spec-endpoints-settings) below.

&#x2022; `validation_mode` - Optional Block<br>Validation Mode. x-required Validation mode of OpenAPI specification. When a validation mismatch occurs on a request to one of the endpoints listed on the OpenAPI specification file (a.k.a. swagger)<br>See [Validation Mode](#api-specification-validation-all-spec-endpoints-validation-mode) below.

<a id="api-specification-validation-all-spec-endpoints-fall-through-mode"></a>

**API Specification Validation All Spec Endpoints Fall Through Mode**

&#x2022; `fall_through_mode_allow` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `fall_through_mode_custom` - Optional Block<br>Custom Fall Through Mode. Define the fall through settings<br>See [Fall Through Mode Custom](#api-specification-validation-all-spec-endpoints-fall-through-mode-fall-through-mode-custom) below.

<a id="api-specification-validation-all-spec-endpoints-fall-through-mode-fall-through-mode-custom"></a>

**API Specification Validation All Spec Endpoints Fall Through Mode Fall Through Mode Custom**

&#x2022; `open_api_validation_rules` - Optional Block<br>Custom Fall Through Rule List<br>See [Open API Validation Rules](#api-specification-validation-all-spec-endpoints-fall-through-mode-fall-through-mode-custom-open-api-validation-rules) below.

<a id="api-specification-validation-all-spec-endpoints-fall-through-mode-fall-through-mode-custom-open-api-validation-rules"></a>

**API Specification Validation All Spec Endpoints Fall Through Mode Fall Through Mode Custom Open API Validation Rules**

&#x2022; `action_block` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `action_report` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `action_skip` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `api_endpoint` - Optional Block<br>API Endpoint. This defines API endpoint<br>See [API Endpoint](#api-specification-validation-all-spec-endpoints-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-api-endpoint) below.

&#x2022; `api_group` - Optional String<br>API Group. The API group which this validation applies to

&#x2022; `base_path` - Optional String<br>Base Path. The base path which this validation applies to

&#x2022; `metadata` - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#api-specification-validation-all-spec-endpoints-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-metadata) below.

<a id="api-specification-validation-all-spec-endpoints-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-api-endpoint"></a>

**API Specification Validation All Spec Endpoints Fall Through Mode Fall Through Mode Custom Open API Validation Rules API Endpoint**

&#x2022; `methods` - Optional List  Defaults to `ANY`<br>Possible values are `ANY`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`, `COPY`<br>Methods. Methods to be matched

&#x2022; `path` - Optional String<br>Path. Path to be matched

<a id="api-specification-validation-all-spec-endpoints-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-metadata"></a>

**API Specification Validation All Spec Endpoints Fall Through Mode Fall Through Mode Custom Open API Validation Rules Metadata**

&#x2022; `description` - Optional String<br>Description. Human readable description

&#x2022; `name` - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

<a id="api-specification-validation-all-spec-endpoints-settings"></a>

**API Specification Validation All Spec Endpoints Settings**

&#x2022; `oversized_body_fail_validation` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `oversized_body_skip_validation` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `property_validation_settings_custom` - Optional Block<br>Validation Property Settings. Custom property validation settings<br>See [Property Validation Settings Custom](#api-specification-validation-all-spec-endpoints-settings-property-validation-settings-custom) below.

&#x2022; `property_validation_settings_default` - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-specification-validation-all-spec-endpoints-settings-property-validation-settings-custom"></a>

**API Specification Validation All Spec Endpoints Settings Property Validation Settings Custom**

&#x2022; `query_parameters` - Optional Block<br>Validation Settings For Query Parameters. Custom settings for query parameters validation<br>See [Query Parameters](#api-specification-validation-all-spec-endpoints-settings-property-validation-settings-custom-query-parameters) below.

<a id="api-specification-validation-all-spec-endpoints-settings-property-validation-settings-custom-query-parameters"></a>

**API Specification Validation All Spec Endpoints Settings Property Validation Settings Custom Query Parameters**

&#x2022; `allow_additional_parameters` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `disallow_additional_parameters` - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-specification-validation-all-spec-endpoints-validation-mode"></a>

**API Specification Validation All Spec Endpoints Validation Mode**

&#x2022; `response_validation_mode_active` - Optional Block<br>Open API Validation Mode Active. Validation mode properties of response<br>See [Response Validation Mode Active](#api-specification-validation-all-spec-endpoints-validation-mode-response-validation-mode-active) below.

&#x2022; `skip_response_validation` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `skip_validation` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `validation_mode_active` - Optional Block<br>Open API Validation Mode Active. Validation mode properties of request<br>See [Validation Mode Active](#api-specification-validation-all-spec-endpoints-validation-mode-validation-mode-active) below.

<a id="api-specification-validation-all-spec-endpoints-validation-mode-response-validation-mode-active"></a>

**API Specification Validation All Spec Endpoints Validation Mode Response Validation Mode Active**

&#x2022; `enforcement_block` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `enforcement_report` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `response_validation_properties` - Optional List  Defaults to `PROPERTY_QUERY_PARAMETERS`<br>Possible values are `PROPERTY_QUERY_PARAMETERS`, `PROPERTY_PATH_PARAMETERS`, `PROPERTY_CONTENT_TYPE`, `PROPERTY_COOKIE_PARAMETERS`, `PROPERTY_HTTP_HEADERS`, `PROPERTY_HTTP_BODY`, `PROPERTY_SECURITY_SCHEMA`, `PROPERTY_RESPONSE_CODE`<br>Response Validation Properties. List of properties of the response to validate according to the OpenAPI specification file (a.k.a. swagger)

<a id="api-specification-validation-all-spec-endpoints-validation-mode-validation-mode-active"></a>

**API Specification Validation All Spec Endpoints Validation Mode Validation Mode Active**

&#x2022; `enforcement_block` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `enforcement_report` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `request_validation_properties` - Optional List  Defaults to `PROPERTY_QUERY_PARAMETERS`<br>Possible values are `PROPERTY_QUERY_PARAMETERS`, `PROPERTY_PATH_PARAMETERS`, `PROPERTY_CONTENT_TYPE`, `PROPERTY_COOKIE_PARAMETERS`, `PROPERTY_HTTP_HEADERS`, `PROPERTY_HTTP_BODY`, `PROPERTY_SECURITY_SCHEMA`, `PROPERTY_RESPONSE_CODE`<br>Request Validation Properties. List of properties of the request to validate according to the OpenAPI specification file (a.k.a. swagger)

<a id="api-specification-validation-custom-list"></a>

**API Specification Validation Custom List**

&#x2022; `fall_through_mode` - Optional Block<br>Fall Through Mode. x-required Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a. swagger) or doesn't have a specific rule in custom rules)<br>See [Fall Through Mode](#api-specification-validation-custom-list-fall-through-mode) below.

&#x2022; `open_api_validation_rules` - Optional Block<br>Validation List<br>See [Open API Validation Rules](#api-specification-validation-custom-list-open-api-validation-rules) below.

&#x2022; `settings` - Optional Block<br>Common Settings. OpenAPI specification validation settings relevant for 'API Inventory' enforcement and for 'Custom list' enforcement<br>See [Settings](#api-specification-validation-custom-list-settings) below.

<a id="api-specification-validation-custom-list-fall-through-mode"></a>

**API Specification Validation Custom List Fall Through Mode**

&#x2022; `fall_through_mode_allow` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `fall_through_mode_custom` - Optional Block<br>Custom Fall Through Mode. Define the fall through settings<br>See [Fall Through Mode Custom](#api-specification-validation-custom-list-fall-through-mode-fall-through-mode-custom) below.

<a id="api-specification-validation-custom-list-fall-through-mode-fall-through-mode-custom"></a>

**API Specification Validation Custom List Fall Through Mode Fall Through Mode Custom**

&#x2022; `open_api_validation_rules` - Optional Block<br>Custom Fall Through Rule List<br>See [Open API Validation Rules](#api-specification-validation-custom-list-fall-through-mode-fall-through-mode-custom-open-api-validation-rules) below.

<a id="api-specification-validation-custom-list-fall-through-mode-fall-through-mode-custom-open-api-validation-rules"></a>

**API Specification Validation Custom List Fall Through Mode Fall Through Mode Custom Open API Validation Rules**

&#x2022; `action_block` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `action_report` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `action_skip` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `api_endpoint` - Optional Block<br>API Endpoint. This defines API endpoint<br>See [API Endpoint](#api-specification-validation-custom-list-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-api-endpoint) below.

&#x2022; `api_group` - Optional String<br>API Group. The API group which this validation applies to

&#x2022; `base_path` - Optional String<br>Base Path. The base path which this validation applies to

&#x2022; `metadata` - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#api-specification-validation-custom-list-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-metadata) below.

<a id="api-specification-validation-custom-list-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-api-endpoint"></a>

**API Specification Validation Custom List Fall Through Mode Fall Through Mode Custom Open API Validation Rules API Endpoint**

&#x2022; `methods` - Optional List  Defaults to `ANY`<br>Possible values are `ANY`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`, `COPY`<br>Methods. Methods to be matched

&#x2022; `path` - Optional String<br>Path. Path to be matched

<a id="api-specification-validation-custom-list-fall-through-mode-fall-through-mode-custom-open-api-validation-rules-metadata"></a>

**API Specification Validation Custom List Fall Through Mode Fall Through Mode Custom Open API Validation Rules Metadata**

&#x2022; `description` - Optional String<br>Description. Human readable description

&#x2022; `name` - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

<a id="api-specification-validation-custom-list-open-api-validation-rules"></a>

**API Specification Validation Custom List Open API Validation Rules**

&#x2022; `any_domain` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `api_endpoint` - Optional Block<br>API Endpoint. This defines API endpoint<br>See [API Endpoint](#api-specification-validation-custom-list-open-api-validation-rules-api-endpoint) below.

&#x2022; `api_group` - Optional String<br>API Group. The API group which this validation applies to

&#x2022; `base_path` - Optional String<br>Base Path. The base path which this validation applies to

&#x2022; `metadata` - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#api-specification-validation-custom-list-open-api-validation-rules-metadata) below.

&#x2022; `specific_domain` - Optional String<br>Specific Domain. The rule will apply for a specific domain

&#x2022; `validation_mode` - Optional Block<br>Validation Mode. x-required Validation mode of OpenAPI specification. When a validation mismatch occurs on a request to one of the endpoints listed on the OpenAPI specification file (a.k.a. swagger)<br>See [Validation Mode](#api-specification-validation-custom-list-open-api-validation-rules-validation-mode) below.

<a id="api-specification-validation-custom-list-open-api-validation-rules-api-endpoint"></a>

**API Specification Validation Custom List Open API Validation Rules API Endpoint**

&#x2022; `methods` - Optional List  Defaults to `ANY`<br>Possible values are `ANY`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`, `COPY`<br>Methods. Methods to be matched

&#x2022; `path` - Optional String<br>Path. Path to be matched

<a id="api-specification-validation-custom-list-open-api-validation-rules-metadata"></a>

**API Specification Validation Custom List Open API Validation Rules Metadata**

&#x2022; `description` - Optional String<br>Description. Human readable description

&#x2022; `name` - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

<a id="api-specification-validation-custom-list-open-api-validation-rules-validation-mode"></a>

**API Specification Validation Custom List Open API Validation Rules Validation Mode**

&#x2022; `response_validation_mode_active` - Optional Block<br>Open API Validation Mode Active. Validation mode properties of response<br>See [Response Validation Mode Active](#api-specification-validation-custom-list-open-api-validation-rules-validation-mode-response-validation-mode-active) below.

&#x2022; `skip_response_validation` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `skip_validation` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `validation_mode_active` - Optional Block<br>Open API Validation Mode Active. Validation mode properties of request<br>See [Validation Mode Active](#api-specification-validation-custom-list-open-api-validation-rules-validation-mode-validation-mode-active) below.

<a id="api-specification-validation-custom-list-open-api-validation-rules-validation-mode-response-validation-mode-active"></a>

**API Specification Validation Custom List Open API Validation Rules Validation Mode Response Validation Mode Active**

&#x2022; `enforcement_block` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `enforcement_report` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `response_validation_properties` - Optional List  Defaults to `PROPERTY_QUERY_PARAMETERS`<br>Possible values are `PROPERTY_QUERY_PARAMETERS`, `PROPERTY_PATH_PARAMETERS`, `PROPERTY_CONTENT_TYPE`, `PROPERTY_COOKIE_PARAMETERS`, `PROPERTY_HTTP_HEADERS`, `PROPERTY_HTTP_BODY`, `PROPERTY_SECURITY_SCHEMA`, `PROPERTY_RESPONSE_CODE`<br>Response Validation Properties. List of properties of the response to validate according to the OpenAPI specification file (a.k.a. swagger)

<a id="api-specification-validation-custom-list-open-api-validation-rules-validation-mode-validation-mode-active"></a>

**API Specification Validation Custom List Open API Validation Rules Validation Mode Validation Mode Active**

&#x2022; `enforcement_block` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `enforcement_report` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `request_validation_properties` - Optional List  Defaults to `PROPERTY_QUERY_PARAMETERS`<br>Possible values are `PROPERTY_QUERY_PARAMETERS`, `PROPERTY_PATH_PARAMETERS`, `PROPERTY_CONTENT_TYPE`, `PROPERTY_COOKIE_PARAMETERS`, `PROPERTY_HTTP_HEADERS`, `PROPERTY_HTTP_BODY`, `PROPERTY_SECURITY_SCHEMA`, `PROPERTY_RESPONSE_CODE`<br>Request Validation Properties. List of properties of the request to validate according to the OpenAPI specification file (a.k.a. swagger)

<a id="api-specification-validation-custom-list-settings"></a>

**API Specification Validation Custom List Settings**

&#x2022; `oversized_body_fail_validation` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `oversized_body_skip_validation` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `property_validation_settings_custom` - Optional Block<br>Validation Property Settings. Custom property validation settings<br>See [Property Validation Settings Custom](#api-specification-validation-custom-list-settings-property-validation-settings-custom) below.

&#x2022; `property_validation_settings_default` - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-specification-validation-custom-list-settings-property-validation-settings-custom"></a>

**API Specification Validation Custom List Settings Property Validation Settings Custom**

&#x2022; `query_parameters` - Optional Block<br>Validation Settings For Query Parameters. Custom settings for query parameters validation<br>See [Query Parameters](#api-specification-validation-custom-list-settings-property-validation-settings-custom-query-parameters) below.

<a id="api-specification-validation-custom-list-settings-property-validation-settings-custom-query-parameters"></a>

**API Specification Validation Custom List Settings Property Validation Settings Custom Query Parameters**

&#x2022; `allow_additional_parameters` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `disallow_additional_parameters` - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-testing"></a>

**API Testing**

&#x2022; `custom_header_value` - Optional String<br>Custom Header. Add x-f5-API-testing-identifier header value to prevent security flags on API testing traffic

&#x2022; `domains` - Optional Block<br>Testing Environments. Add and configure testing domains and credentials<br>See [Domains](#api-testing-domains) below.

&#x2022; `every_day` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `every_month` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `every_week` - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-testing-domains"></a>

**API Testing Domains**

&#x2022; `allow_destructive_methods` - Optional Bool<br>Use Destructive Methods (e.g., DELETE, PUT). Enable to allow API test to execute destructive methods. Be cautious as these can alter or delete data

&#x2022; `credentials` - Optional Block<br>Credentials. Add credentials for API testing to use in the selected environment<br>See [Credentials](#api-testing-domains-credentials) below.

&#x2022; `domain` - Optional String<br>Domain. Add your testing environment domain. Be aware that running tests on a production domain can impact live applications, as API testing cannot distinguish between production and testing environments

<a id="api-testing-domains-credentials"></a>

**API Testing Domains Credentials**

&#x2022; `admin` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `api_key` - Optional Block<br>API Key<br>See [API Key](#api-testing-domains-credentials-api-key) below.

&#x2022; `basic_auth` - Optional Block<br>Basic Authentication<br>See [Basic Auth](#api-testing-domains-credentials-basic-auth) below.

&#x2022; `bearer_token` - Optional Block<br>Bearer<br>See [Bearer Token](#api-testing-domains-credentials-bearer-token) below.

&#x2022; `credential_name` - Optional String<br>Name. Enter a unique name for the credentials used in API testing

&#x2022; `login_endpoint` - Optional Block<br>Login Endpoint<br>See [Login Endpoint](#api-testing-domains-credentials-login-endpoint) below.

&#x2022; `standard` - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="api-testing-domains-credentials-api-key"></a>

**API Testing Domains Credentials API Key**

&#x2022; `key` - Optional String<br>Key

&#x2022; `value` - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Value](#api-testing-domains-credentials-api-key-value) below.

<a id="api-testing-domains-credentials-api-key-value"></a>

**API Testing Domains Credentials API Key Value**

&#x2022; `blindfold_secret_info` - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#api-testing-domains-credentials-api-key-value-blindfold-secret-info) below.

&#x2022; `clear_secret_info` - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#api-testing-domains-credentials-api-key-value-clear-secret-info) below.

<a id="api-testing-domains-credentials-api-key-value-blindfold-secret-info"></a>

**API Testing Domains Credentials API Key Value Blindfold Secret Info**

&#x2022; `decryption_provider` - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

&#x2022; `location` - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

&#x2022; `store_provider` - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

<a id="api-testing-domains-credentials-api-key-value-clear-secret-info"></a>

**API Testing Domains Credentials API Key Value Clear Secret Info**

&#x2022; `provider_ref` - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

&#x2022; `url` - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

<a id="api-testing-domains-credentials-basic-auth"></a>

**API Testing Domains Credentials Basic Auth**

&#x2022; `password` - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Password](#api-testing-domains-credentials-basic-auth-password) below.

&#x2022; `user` - Optional String<br>User

<a id="api-testing-domains-credentials-basic-auth-password"></a>

**API Testing Domains Credentials Basic Auth Password**

&#x2022; `blindfold_secret_info` - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#api-testing-domains-credentials-basic-auth-password-blindfold-secret-info) below.

&#x2022; `clear_secret_info` - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#api-testing-domains-credentials-basic-auth-password-clear-secret-info) below.

<a id="api-testing-domains-credentials-basic-auth-password-blindfold-secret-info"></a>

**API Testing Domains Credentials Basic Auth Password Blindfold Secret Info**

&#x2022; `decryption_provider` - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

&#x2022; `location` - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

&#x2022; `store_provider` - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

<a id="api-testing-domains-credentials-basic-auth-password-clear-secret-info"></a>

**API Testing Domains Credentials Basic Auth Password Clear Secret Info**

&#x2022; `provider_ref` - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

&#x2022; `url` - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

<a id="api-testing-domains-credentials-bearer-token"></a>

**API Testing Domains Credentials Bearer Token**

&#x2022; `token` - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Token](#api-testing-domains-credentials-bearer-token-token) below.

<a id="api-testing-domains-credentials-bearer-token-token"></a>

**API Testing Domains Credentials Bearer Token Token**

&#x2022; `blindfold_secret_info` - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#api-testing-domains-credentials-bearer-token-token-blindfold-secret-info) below.

&#x2022; `clear_secret_info` - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#api-testing-domains-credentials-bearer-token-token-clear-secret-info) below.

<a id="api-testing-domains-credentials-bearer-token-token-blindfold-secret-info"></a>

**API Testing Domains Credentials Bearer Token Token Blindfold Secret Info**

&#x2022; `decryption_provider` - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

&#x2022; `location` - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

&#x2022; `store_provider` - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

<a id="api-testing-domains-credentials-bearer-token-token-clear-secret-info"></a>

**API Testing Domains Credentials Bearer Token Token Clear Secret Info**

&#x2022; `provider_ref` - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

&#x2022; `url` - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

<a id="api-testing-domains-credentials-login-endpoint"></a>

**API Testing Domains Credentials Login Endpoint**

&#x2022; `json_payload` - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [JSON Payload](#api-testing-domains-credentials-login-endpoint-json-payload) below.

&#x2022; `method` - Optional String  Defaults to `ANY`<br>Possible values are `ANY`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`, `COPY`<br>HTTP Method. Specifies the HTTP method used to access a resource. Any HTTP Method

&#x2022; `path` - Optional String<br>Path

&#x2022; `token_response_key` - Optional String<br>Token Response Key. Specifies how to handle the API response, extracting authentication tokens

<a id="api-testing-domains-credentials-login-endpoint-json-payload"></a>

**API Testing Domains Credentials Login Endpoint JSON Payload**

&#x2022; `blindfold_secret_info` - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#api-testing-domains-credentials-login-endpoint-json-payload-blindfold-secret-info) below.

&#x2022; `clear_secret_info` - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#api-testing-domains-credentials-login-endpoint-json-payload-clear-secret-info) below.

<a id="api-testing-domains-credentials-login-endpoint-json-payload-blindfold-secret-info"></a>

**API Testing Domains Credentials Login Endpoint JSON Payload Blindfold Secret Info**

&#x2022; `decryption_provider` - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

&#x2022; `location` - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

&#x2022; `store_provider` - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

<a id="api-testing-domains-credentials-login-endpoint-json-payload-clear-secret-info"></a>

**API Testing Domains Credentials Login Endpoint JSON Payload Clear Secret Info**

&#x2022; `provider_ref` - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

&#x2022; `url` - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

<a id="app-firewall"></a>

**App Firewall**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="blocked-clients"></a>

**Blocked Clients**

&#x2022; `actions` - Optional List  Defaults to `SKIP_PROCESSING_WAF`<br>Possible values are `SKIP_PROCESSING_WAF`, `SKIP_PROCESSING_BOT`, `SKIP_PROCESSING_MUM`, `SKIP_PROCESSING_IP_REPUTATION`, `SKIP_PROCESSING_API_PROTECTION`, `SKIP_PROCESSING_OAS_VALIDATION`, `SKIP_PROCESSING_DDOS_PROTECTION`, `SKIP_PROCESSING_THREAT_MESH`, `SKIP_PROCESSING_MALWARE_PROTECTION`<br>Actions. Actions that should be taken when client identifier matches the rule

&#x2022; `as_number` - Optional Number<br>AS Number. RFC 6793 defined 4-byte AS number

&#x2022; `bot_skip_processing` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `expiration_timestamp` - Optional String<br>Expiration Timestamp. The expiration_timestamp is the RFC 3339 format timestamp at which the containing rule is considered to be logically expired. The rule continues to exist in the configuration but is not applied anymore

&#x2022; `http_header` - Optional Block<br>HTTP Header. Request header name and value pairs<br>See [HTTP Header](#blocked-clients-http-header) below.

&#x2022; `ip_prefix` - Optional String<br>IPv4 Prefix. IPv4 prefix string

&#x2022; `ipv6_prefix` - Optional String<br>IPv6 Prefix. IPv6 prefix string

&#x2022; `metadata` - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#blocked-clients-metadata) below.

&#x2022; `skip_processing` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `user_identifier` - Optional String<br>User Identifier. Identify user based on user identifier. User identifier value needs to be copied from security event

&#x2022; `waf_skip_processing` - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="blocked-clients-http-header"></a>

**Blocked Clients HTTP Header**

&#x2022; `headers` - Optional Block<br>HTTP Headers. List of HTTP header name and value pairs<br>See [Headers](#blocked-clients-http-header-headers) below.

<a id="blocked-clients-http-header-headers"></a>

**Blocked Clients HTTP Header Headers**

&#x2022; `exact` - Optional String<br>Exact. Header value to match exactly

&#x2022; `invert_match` - Optional Bool<br>NOT of match. Invert the result of the match to detect missing header or non-matching value

&#x2022; `name` - Optional String<br>Name. Name of the header

&#x2022; `presence` - Optional Bool<br>Presence. If true, check for presence of header

&#x2022; `regex` - Optional String<br>Regex. Regex match of the header value in re2 format

<a id="blocked-clients-metadata"></a>

**Blocked Clients Metadata**

&#x2022; `description` - Optional String<br>Description. Human readable description

&#x2022; `name` - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

<a id="bot-defense"></a>

**Bot Defense**

&#x2022; `disable_cors_support` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `enable_cors_support` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `policy` - Optional Block<br>Bot Defense Policy. This defines various configuration options for Bot Defense policy<br>See [Policy](#bot-defense-policy) below.

&#x2022; `regional_endpoint` - Optional String  Defaults to `AUTO`<br>Possible values are `AUTO`, `US`, `EU`, `ASIA`<br>Bot Defense Region. Defines a selection for Bot Defense region - AUTO: AUTO Automatic selection based on client IP address - US: US US region - EU: EU European Union region - ASIA: ASIA Asia region

&#x2022; `timeout` - Optional Number<br>Timeout. The timeout for the inference check, in milliseconds

<a id="bot-defense-policy"></a>

**Bot Defense Policy**

&#x2022; `disable_js_insert` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `disable_mobile_sdk` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `javascript_mode` - Optional String  Defaults to `ASYNC_JS_NO_CACHING`<br>Possible values are `ASYNC_JS_NO_CACHING`, `ASYNC_JS_CACHING`, `SYNC_JS_NO_CACHING`, `SYNC_JS_CACHING`<br>Web Client JavaScript Mode. Web Client JavaScript Mode. Bot Defense JavaScript for telemetry collection is requested asynchronously, and it is non-cacheable Bot Defense JavaScript for telemetry collection is requested asynchronously, and it is cacheable Bot Defense JavaScript for telemetry collection is requested synchronously, and it is non-cacheable Bot Defense JavaScript for telemetry collection is requested synchronously, and it is cacheable

&#x2022; `js_download_path` - Optional String<br>JavaScript Download Path. Customize Bot Defense Client JavaScript path. If not specified, default `/common.js`

&#x2022; `js_insert_all_pages` - Optional Block<br>Insert Bot Defense JavaScript in All Pages. Insert Bot Defense JavaScript in all pages<br>See [Js Insert All Pages](#bot-defense-policy-js-insert-all-pages) below.

&#x2022; `js_insert_all_pages_except` - Optional Block<br>Insert JavaScript in All Pages with the Exceptions. Insert Bot Defense JavaScript in all pages with the exceptions<br>See [Js Insert All Pages Except](#bot-defense-policy-js-insert-all-pages-except) below.

&#x2022; `js_insertion_rules` - Optional Block<br>JavaScript Custom Insertion Rules. This defines custom JavaScript insertion rules for Bot Defense Policy<br>See [Js Insertion Rules](#bot-defense-policy-js-insertion-rules) below.

&#x2022; `mobile_sdk_config` - Optional Block<br>Mobile SDK Configuration. Mobile SDK configuration<br>See [Mobile Sdk Config](#bot-defense-policy-mobile-sdk-config) below.

&#x2022; `protected_app_endpoints` - Optional Block<br>App Endpoint Type. List of protected endpoints. Limit: Approx '128 endpoints per Load Balancer (LB)' upto 4 LBs, '32 endpoints per LB' after 4 LBs<br>See [Protected App Endpoints](#bot-defense-policy-protected-app-endpoints) below.

<a id="bot-defense-policy-js-insert-all-pages"></a>

**Bot Defense Policy Js Insert All Pages**

&#x2022; `javascript_location` - Optional String  Defaults to `AFTER_HEAD`<br>Possible values are `AFTER_HEAD`, `AFTER_TITLE_END`, `BEFORE_SCRIPT`<br>JavaScript Location. All inside networks. Insert JavaScript after <head> tag Insert JavaScript after </title> tag. Insert JavaScript before first <script> tag

<a id="bot-defense-policy-js-insert-all-pages-except"></a>

**Bot Defense Policy Js Insert All Pages Except**

&#x2022; `exclude_list` - Optional Block<br>Exclude Pages. Optional JavaScript insertions exclude list of domain and path matchers<br>See [Exclude List](#bot-defense-policy-js-insert-all-pages-except-exclude-list) below.

&#x2022; `javascript_location` - Optional String  Defaults to `AFTER_HEAD`<br>Possible values are `AFTER_HEAD`, `AFTER_TITLE_END`, `BEFORE_SCRIPT`<br>JavaScript Location. All inside networks. Insert JavaScript after <head> tag Insert JavaScript after </title> tag. Insert JavaScript before first <script> tag

<a id="bot-defense-policy-js-insert-all-pages-except-exclude-list"></a>

**Bot Defense Policy Js Insert All Pages Except Exclude List**

&#x2022; `any_domain` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `domain` - Optional Block<br>Domains. Domains names<br>See [Domain](#bot-defense-policy-js-insert-all-pages-except-exclude-list-domain) below.

&#x2022; `metadata` - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#bot-defense-policy-js-insert-all-pages-except-exclude-list-metadata) below.

&#x2022; `path` - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#bot-defense-policy-js-insert-all-pages-except-exclude-list-path) below.

<a id="bot-defense-policy-js-insert-all-pages-except-exclude-list-domain"></a>

**Bot Defense Policy Js Insert All Pages Except Exclude List Domain**

&#x2022; `exact_value` - Optional String<br>Exact Value. Exact domain name

&#x2022; `regex_value` - Optional String<br>Regex Values of Domains. Regular Expression value for the domain name

&#x2022; `suffix_value` - Optional String<br>Suffix Value. Suffix of domain name e.g 'xyz.com' will match '*.xyz.com' and 'xyz.com'

<a id="bot-defense-policy-js-insert-all-pages-except-exclude-list-metadata"></a>

**Bot Defense Policy Js Insert All Pages Except Exclude List Metadata**

&#x2022; `description` - Optional String<br>Description. Human readable description

&#x2022; `name` - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

<a id="bot-defense-policy-js-insert-all-pages-except-exclude-list-path"></a>

**Bot Defense Policy Js Insert All Pages Except Exclude List Path**

&#x2022; `path` - Optional String<br>Exact. Exact path value to match

&#x2022; `prefix` - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

&#x2022; `regex` - Optional String<br>Regex. Regular expression of path match (e.g. the value .* will match on all paths)

<a id="bot-defense-policy-js-insertion-rules"></a>

**Bot Defense Policy Js Insertion Rules**

&#x2022; `exclude_list` - Optional Block<br>Exclude Paths. Optional JavaScript insertions exclude list of domain and path matchers<br>See [Exclude List](#bot-defense-policy-js-insertion-rules-exclude-list) below.

&#x2022; `rules` - Optional Block<br>JavaScript Insertions. Required list of pages to insert Bot Defense client JavaScript<br>See [Rules](#bot-defense-policy-js-insertion-rules-rules) below.

<a id="bot-defense-policy-js-insertion-rules-exclude-list"></a>

**Bot Defense Policy Js Insertion Rules Exclude List**

&#x2022; `any_domain` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `domain` - Optional Block<br>Domains. Domains names<br>See [Domain](#bot-defense-policy-js-insertion-rules-exclude-list-domain) below.

&#x2022; `metadata` - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#bot-defense-policy-js-insertion-rules-exclude-list-metadata) below.

&#x2022; `path` - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#bot-defense-policy-js-insertion-rules-exclude-list-path) below.

<a id="bot-defense-policy-js-insertion-rules-exclude-list-domain"></a>

**Bot Defense Policy Js Insertion Rules Exclude List Domain**

&#x2022; `exact_value` - Optional String<br>Exact Value. Exact domain name

&#x2022; `regex_value` - Optional String<br>Regex Values of Domains. Regular Expression value for the domain name

&#x2022; `suffix_value` - Optional String<br>Suffix Value. Suffix of domain name e.g 'xyz.com' will match '*.xyz.com' and 'xyz.com'

<a id="bot-defense-policy-js-insertion-rules-exclude-list-metadata"></a>

**Bot Defense Policy Js Insertion Rules Exclude List Metadata**

&#x2022; `description` - Optional String<br>Description. Human readable description

&#x2022; `name` - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

<a id="bot-defense-policy-js-insertion-rules-exclude-list-path"></a>

**Bot Defense Policy Js Insertion Rules Exclude List Path**

&#x2022; `path` - Optional String<br>Exact. Exact path value to match

&#x2022; `prefix` - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

&#x2022; `regex` - Optional String<br>Regex. Regular expression of path match (e.g. the value .* will match on all paths)

<a id="bot-defense-policy-js-insertion-rules-rules"></a>

**Bot Defense Policy Js Insertion Rules Rules**

&#x2022; `any_domain` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `domain` - Optional Block<br>Domains. Domains names<br>See [Domain](#bot-defense-policy-js-insertion-rules-rules-domain) below.

&#x2022; `javascript_location` - Optional String  Defaults to `AFTER_HEAD`<br>Possible values are `AFTER_HEAD`, `AFTER_TITLE_END`, `BEFORE_SCRIPT`<br>JavaScript Location. All inside networks. Insert JavaScript after <head> tag Insert JavaScript after </title> tag. Insert JavaScript before first <script> tag

&#x2022; `metadata` - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#bot-defense-policy-js-insertion-rules-rules-metadata) below.

&#x2022; `path` - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#bot-defense-policy-js-insertion-rules-rules-path) below.

<a id="bot-defense-policy-js-insertion-rules-rules-domain"></a>

**Bot Defense Policy Js Insertion Rules Rules Domain**

&#x2022; `exact_value` - Optional String<br>Exact Value. Exact domain name

&#x2022; `regex_value` - Optional String<br>Regex Values of Domains. Regular Expression value for the domain name

&#x2022; `suffix_value` - Optional String<br>Suffix Value. Suffix of domain name e.g 'xyz.com' will match '*.xyz.com' and 'xyz.com'

<a id="bot-defense-policy-js-insertion-rules-rules-metadata"></a>

**Bot Defense Policy Js Insertion Rules Rules Metadata**

&#x2022; `description` - Optional String<br>Description. Human readable description

&#x2022; `name` - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

<a id="bot-defense-policy-js-insertion-rules-rules-path"></a>

**Bot Defense Policy Js Insertion Rules Rules Path**

&#x2022; `path` - Optional String<br>Exact. Exact path value to match

&#x2022; `prefix` - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

&#x2022; `regex` - Optional String<br>Regex. Regular expression of path match (e.g. the value .* will match on all paths)

<a id="bot-defense-policy-mobile-sdk-config"></a>

**Bot Defense Policy Mobile Sdk Config**

&#x2022; `mobile_identifier` - Optional Block<br>Mobile Traffic Identifier. Mobile traffic identifier type<br>See [Mobile Identifier](#bot-defense-policy-mobile-sdk-config-mobile-identifier) below.

<a id="bot-defense-policy-mobile-sdk-config-mobile-identifier"></a>

**Bot Defense Policy Mobile Sdk Config Mobile Identifier**

&#x2022; `headers` - Optional Block<br>Headers. Headers that can be used to identify mobile traffic<br>See [Headers](#bot-defense-policy-mobile-sdk-config-mobile-identifier-headers) below.

<a id="bot-defense-policy-mobile-sdk-config-mobile-identifier-headers"></a>

**Bot Defense Policy Mobile Sdk Config Mobile Identifier Headers**

&#x2022; `check_not_present` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `check_present` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `item` - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#bot-defense-policy-mobile-sdk-config-mobile-identifier-headers-item) below.

&#x2022; `name` - Optional String<br>Header Name. A case-insensitive HTTP header name

<a id="bot-defense-policy-mobile-sdk-config-mobile-identifier-headers-item"></a>

**Bot Defense Policy Mobile Sdk Config Mobile Identifier Headers Item**

&#x2022; `exact_values` - Optional List<br>Exact Values. A list of exact values to match the input against

&#x2022; `regex_values` - Optional List<br>Regex Values. A list of regular expressions to match the input against

&#x2022; `transformers` - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

<a id="bot-defense-policy-protected-app-endpoints"></a>

**Bot Defense Policy Protected App Endpoints**

&#x2022; `allow_good_bots` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `any_domain` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `domain` - Optional Block<br>Domains. Domains names<br>See [Domain](#bot-defense-policy-protected-app-endpoints-domain) below.

&#x2022; `flow_label` - Optional Block<br>Bot Defense Flow Label Category. Bot Defense Flow Label Category allows to associate traffic with selected category<br>See [Flow Label](#bot-defense-policy-protected-app-endpoints-flow-label) below.

&#x2022; `headers` - Optional Block<br>HTTP Headers. A list of predicates for various HTTP headers that need to match. The criteria for matching each HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values are extracted from the request API as a list of strings for each HTTP header type. Note that all specified header predicates must evaluate to true<br>See [Headers](#bot-defense-policy-protected-app-endpoints-headers) below.

&#x2022; `http_methods` - Optional List  Defaults to `METHOD_ANY`<br>Possible values are `METHOD_ANY`, `METHOD_GET`, `METHOD_POST`, `METHOD_PUT`, `METHOD_PATCH`, `METHOD_DELETE`, `METHOD_GET_DOCUMENT`<br>HTTP Methods. List of HTTP methods

&#x2022; `metadata` - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#bot-defense-policy-protected-app-endpoints-metadata) below.

&#x2022; `mitigate_good_bots` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `mitigation` - Optional Block<br>Bot Mitigation Action. Modify Bot Defense behavior for a matching request<br>See [Mitigation](#bot-defense-policy-protected-app-endpoints-mitigation) below.

&#x2022; `mobile` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `path` - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#bot-defense-policy-protected-app-endpoints-path) below.

&#x2022; `protocol` - Optional String  Defaults to `BOTH`<br>Possible values are `BOTH`, `HTTP`, `HTTPS`<br>URL Scheme. SchemeType is used to indicate URL scheme. - BOTH: BOTH URL scheme for HTTPS:// or `HTTP://.` - HTTP: HTTP URL scheme HTTP:// only. - HTTPS: HTTPS URL scheme HTTPS:// only

&#x2022; `query_params` - Optional Block<br>HTTP Query Parameters. A list of predicates for all query parameters that need to be matched. The criteria for matching each query parameter are described in individual instances of QueryParameterMatcherType. The actual query parameter values are extracted from the request API as a list of strings for each query parameter name. Note that all specified query parameter predicates must evaluate to true<br>See [Query Params](#bot-defense-policy-protected-app-endpoints-query-params) below.

&#x2022; `undefined_flow_label` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `web` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `web_mobile` - Optional Block<br>Web and Mobile traffic type. Web and Mobile traffic type<br>See [Web Mobile](#bot-defense-policy-protected-app-endpoints-web-mobile) below.

<a id="bot-defense-policy-protected-app-endpoints-domain"></a>

**Bot Defense Policy Protected App Endpoints Domain**

&#x2022; `exact_value` - Optional String<br>Exact Value. Exact domain name

&#x2022; `regex_value` - Optional String<br>Regex Values of Domains. Regular Expression value for the domain name

&#x2022; `suffix_value` - Optional String<br>Suffix Value. Suffix of domain name e.g 'xyz.com' will match '*.xyz.com' and 'xyz.com'

<a id="bot-defense-policy-protected-app-endpoints-flow-label"></a>

**Bot Defense Policy Protected App Endpoints Flow Label**

&#x2022; `account_management` - Optional Block<br>Bot Defense Flow Label Account Management Category. Bot Defense Flow Label Account Management Category<br>See [Account Management](#bot-defense-policy-protected-app-endpoints-flow-label-account-management) below.

&#x2022; `authentication` - Optional Block<br>Bot Defense Flow Label Authentication Category. Bot Defense Flow Label Authentication Category<br>See [Authentication](#bot-defense-policy-protected-app-endpoints-flow-label-authentication) below.

&#x2022; `financial_services` - Optional Block<br>Bot Defense Flow Label Financial Services Category. Bot Defense Flow Label Financial Services Category<br>See [Financial Services](#bot-defense-policy-protected-app-endpoints-flow-label-financial-services) below.

&#x2022; `flight` - Optional Block<br>Bot Defense Flow Label Flight Category. Bot Defense Flow Label Flight Category<br>See [Flight](#bot-defense-policy-protected-app-endpoints-flow-label-flight) below.

&#x2022; `profile_management` - Optional Block<br>Bot Defense Flow Label Profile Management Category. Bot Defense Flow Label Profile Management Category<br>See [Profile Management](#bot-defense-policy-protected-app-endpoints-flow-label-profile-management) below.

&#x2022; `search` - Optional Block<br>Bot Defense Flow Label Search Category. Bot Defense Flow Label Search Category<br>See [Search](#bot-defense-policy-protected-app-endpoints-flow-label-search) below.

&#x2022; `shopping_gift_cards` - Optional Block<br>Bot Defense Flow Label Shopping & Gift Cards Category. Bot Defense Flow Label Shopping & Gift Cards Category<br>See [Shopping Gift Cards](#bot-defense-policy-protected-app-endpoints-flow-label-shopping-gift-cards) below.

<a id="bot-defense-policy-protected-app-endpoints-flow-label-account-management"></a>

**Bot Defense Policy Protected App Endpoints Flow Label Account Management**

&#x2022; `create` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `password_reset` - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="bot-defense-policy-protected-app-endpoints-flow-label-authentication"></a>

**Bot Defense Policy Protected App Endpoints Flow Label Authentication**

&#x2022; `login` - Optional Block<br>Bot Defense Transaction Result. Bot Defense Transaction Result<br>See [Login](#bot-defense-policy-protected-app-endpoints-flow-label-authentication-login) below.

&#x2022; `login_mfa` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `login_partner` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `logout` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `token_refresh` - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="bot-defense-policy-protected-app-endpoints-flow-label-authentication-login"></a>

**Bot Defense Policy Protected App Endpoints Flow Label Authentication Login**

&#x2022; `disable_transaction_result` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `transaction_result` - Optional Block<br>Bot Defense Transaction Result Type. Bot Defense Transaction ResultType<br>See [Transaction Result](#bot-defense-policy-protected-app-endpoints-flow-label-authentication-login-transaction-result) below.

<a id="bot-defense-policy-protected-app-endpoints-flow-label-authentication-login-transaction-result"></a>

**Bot Defense Policy Protected App Endpoints Flow Label Authentication Login Transaction Result**

&#x2022; `failure_conditions` - Optional Block<br>Failure Conditions. Failure Conditions<br>See [Failure Conditions](#bot-defense-policy-protected-app-endpoints-flow-label-authentication-login-transaction-result-failure-conditions) below.

&#x2022; `success_conditions` - Optional Block<br>Success Conditions. Success Conditions<br>See [Success Conditions](#bot-defense-policy-protected-app-endpoints-flow-label-authentication-login-transaction-result-success-conditions) below.

<a id="bot-defense-policy-protected-app-endpoints-flow-label-authentication-login-transaction-result-failure-conditions"></a>

**Bot Defense Policy Protected App Endpoints Flow Label Authentication Login Transaction Result Failure Conditions**

&#x2022; `name` - Optional String<br>Header Name. A case-insensitive HTTP header name

&#x2022; `regex_values` - Optional List<br>Regex Values. A list of regular expressions to match the input against

&#x2022; `status` - Optional String  Defaults to `EmptyStatusCode`<br>Possible values are `EmptyStatusCode`, `Continue`, `OK`, `Created`, `Accepted`, `NonAuthoritativeInformation`, `NoContent`, `ResetContent`, `PartialContent`, `MultiStatus`, `AlreadyReported`, `IMUsed`, `MultipleChoices`, `MovedPermanently`, `Found`, `SeeOther`, `NotModified`, `UseProxy`, `TemporaryRedirect`, `PermanentRedirect`, `BadRequest`, `Unauthorized`, `PaymentRequired`, `Forbidden`, `NotFound`, `MethodNotAllowed`, `NotAcceptable`, `ProxyAuthenticationRequired`, `RequestTimeout`, `Conflict`, `Gone`, `LengthRequired`, `PreconditionFailed`, `PayloadTooLarge`, `URITooLong`, `UnsupportedMediaType`, `RangeNotSatisfiable`, `ExpectationFailed`, `MisdirectedRequest`, `UnprocessableEntity`, `Locked`, `FailedDependency`, `UpgradeRequired`, `PreconditionRequired`, `TooManyRequests`, `RequestHeaderFieldsTooLarge`, `InternalServerError`, `NotImplemented`, `BadGateway`, `ServiceUnavailable`, `GatewayTimeout`, `HTTPVersionNotSupported`, `VariantAlsoNegotiates`, `InsufficientStorage`, `LoopDetected`, `NotExtended`, `NetworkAuthenticationRequired`<br>HTTP Status Code. HTTP response status codes EmptyStatusCode response codes means it is not specified Continue status code OK status code Created status code Accepted status code Non Authoritative Information status code No Content status code Reset Content status code Partial Content status code Multi Status status code Already Reported status code Im Used status code Multiple Choices status code Moved Permanently status code Found status code See Other status code Not Modified status code Use Proxy status code Temporary Redirect status code Permanent Redirect status code Bad Request status code Unauthorized status code Payment Required status code Forbidden status code Not Found status code Method Not Allowed status code Not Acceptable status code Proxy Authentication Required status code Request Timeout status code Conflict status code Gone status code Length Required status code Precondition Failed status code Payload Too Large status code URI Too Long status code Unsupported Media Type status code Range Not Satisfiable status code Expectation Failed status code Misdirected Request status code Unprocessable Entity status code Locked status code Failed Dependency status code Upgrade Required status code Precondition Required status code Too Many Requests status code Request Header Fields Too Large status code Internal Server Error status code Not Implemented status code Bad Gateway status code Service Unavailable status code Gateway Timeout status code HTTP Version Not Supported status code Variant Also Negotiates status code Insufficient Storage status code Loop Detected status code Not Extended status code Network Authentication Required status code

<a id="bot-defense-policy-protected-app-endpoints-flow-label-authentication-login-transaction-result-success-conditions"></a>

**Bot Defense Policy Protected App Endpoints Flow Label Authentication Login Transaction Result Success Conditions**

&#x2022; `name` - Optional String<br>Header Name. A case-insensitive HTTP header name

&#x2022; `regex_values` - Optional List<br>Regex Values. A list of regular expressions to match the input against

&#x2022; `status` - Optional String  Defaults to `EmptyStatusCode`<br>Possible values are `EmptyStatusCode`, `Continue`, `OK`, `Created`, `Accepted`, `NonAuthoritativeInformation`, `NoContent`, `ResetContent`, `PartialContent`, `MultiStatus`, `AlreadyReported`, `IMUsed`, `MultipleChoices`, `MovedPermanently`, `Found`, `SeeOther`, `NotModified`, `UseProxy`, `TemporaryRedirect`, `PermanentRedirect`, `BadRequest`, `Unauthorized`, `PaymentRequired`, `Forbidden`, `NotFound`, `MethodNotAllowed`, `NotAcceptable`, `ProxyAuthenticationRequired`, `RequestTimeout`, `Conflict`, `Gone`, `LengthRequired`, `PreconditionFailed`, `PayloadTooLarge`, `URITooLong`, `UnsupportedMediaType`, `RangeNotSatisfiable`, `ExpectationFailed`, `MisdirectedRequest`, `UnprocessableEntity`, `Locked`, `FailedDependency`, `UpgradeRequired`, `PreconditionRequired`, `TooManyRequests`, `RequestHeaderFieldsTooLarge`, `InternalServerError`, `NotImplemented`, `BadGateway`, `ServiceUnavailable`, `GatewayTimeout`, `HTTPVersionNotSupported`, `VariantAlsoNegotiates`, `InsufficientStorage`, `LoopDetected`, `NotExtended`, `NetworkAuthenticationRequired`<br>HTTP Status Code. HTTP response status codes EmptyStatusCode response codes means it is not specified Continue status code OK status code Created status code Accepted status code Non Authoritative Information status code No Content status code Reset Content status code Partial Content status code Multi Status status code Already Reported status code Im Used status code Multiple Choices status code Moved Permanently status code Found status code See Other status code Not Modified status code Use Proxy status code Temporary Redirect status code Permanent Redirect status code Bad Request status code Unauthorized status code Payment Required status code Forbidden status code Not Found status code Method Not Allowed status code Not Acceptable status code Proxy Authentication Required status code Request Timeout status code Conflict status code Gone status code Length Required status code Precondition Failed status code Payload Too Large status code URI Too Long status code Unsupported Media Type status code Range Not Satisfiable status code Expectation Failed status code Misdirected Request status code Unprocessable Entity status code Locked status code Failed Dependency status code Upgrade Required status code Precondition Required status code Too Many Requests status code Request Header Fields Too Large status code Internal Server Error status code Not Implemented status code Bad Gateway status code Service Unavailable status code Gateway Timeout status code HTTP Version Not Supported status code Variant Also Negotiates status code Insufficient Storage status code Loop Detected status code Not Extended status code Network Authentication Required status code

<a id="bot-defense-policy-protected-app-endpoints-flow-label-financial-services"></a>

**Bot Defense Policy Protected App Endpoints Flow Label Financial Services**

&#x2022; `apply` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `money_transfer` - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="bot-defense-policy-protected-app-endpoints-flow-label-flight"></a>

**Bot Defense Policy Protected App Endpoints Flow Label Flight**

&#x2022; `checkin` - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="bot-defense-policy-protected-app-endpoints-flow-label-profile-management"></a>

**Bot Defense Policy Protected App Endpoints Flow Label Profile Management**

&#x2022; `create` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `update` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `view` - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="bot-defense-policy-protected-app-endpoints-flow-label-search"></a>

**Bot Defense Policy Protected App Endpoints Flow Label Search**

&#x2022; `flight_search` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `product_search` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `reservation_search` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `room_search` - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="bot-defense-policy-protected-app-endpoints-flow-label-shopping-gift-cards"></a>

**Bot Defense Policy Protected App Endpoints Flow Label Shopping Gift Cards**

&#x2022; `gift_card_make_purchase_with_gift_card` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `gift_card_validation` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `shop_add_to_cart` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `shop_checkout` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `shop_choose_seat` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `shop_enter_drawing_submission` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `shop_make_payment` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `shop_order` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `shop_price_inquiry` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `shop_promo_code_validation` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `shop_purchase_gift_card` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `shop_update_quantity` - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="bot-defense-policy-protected-app-endpoints-headers"></a>

**Bot Defense Policy Protected App Endpoints Headers**

&#x2022; `check_not_present` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `check_present` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `invert_matcher` - Optional Bool<br>Invert Header Matcher. Invert the match result

&#x2022; `item` - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#bot-defense-policy-protected-app-endpoints-headers-item) below.

&#x2022; `name` - Optional String<br>Header Name. A case-insensitive HTTP header name

<a id="bot-defense-policy-protected-app-endpoints-headers-item"></a>

**Bot Defense Policy Protected App Endpoints Headers Item**

&#x2022; `exact_values` - Optional List<br>Exact Values. A list of exact values to match the input against

&#x2022; `regex_values` - Optional List<br>Regex Values. A list of regular expressions to match the input against

&#x2022; `transformers` - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

<a id="bot-defense-policy-protected-app-endpoints-metadata"></a>

**Bot Defense Policy Protected App Endpoints Metadata**

&#x2022; `description` - Optional String<br>Description. Human readable description

&#x2022; `name` - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

<a id="bot-defense-policy-protected-app-endpoints-mitigation"></a>

**Bot Defense Policy Protected App Endpoints Mitigation**

&#x2022; `block` - Optional Block<br>Block bot mitigation. Block request and respond with custom content<br>See [Block](#bot-defense-policy-protected-app-endpoints-mitigation-block) below.

&#x2022; `flag` - Optional Block<br>Select Flag Bot Mitigation Action. Flag mitigation action<br>See [Flag](#bot-defense-policy-protected-app-endpoints-mitigation-flag) below.

&#x2022; `redirect` - Optional Block<br>Redirect bot mitigation. Redirect request to a custom URI<br>See [Redirect](#bot-defense-policy-protected-app-endpoints-mitigation-redirect) below.

<a id="bot-defense-policy-protected-app-endpoints-mitigation-block"></a>

**Bot Defense Policy Protected App Endpoints Mitigation Block**

&#x2022; `body` - Optional String<br>Body. Custom body message is of type uri_ref. Currently supported URL schemes is string:///. For string:/// scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. 'Your request was blocked' or it can be HTML paragraph or a body string encoded as base64 string E.g. '<p> Your request was blocked </p>'. Base64 encoded string for this HTML is 'LzxwPiBZb3VyIHJlcXVlc3Qgd2FzIGJsb2NrZWQgPC9wPg=='

&#x2022; `status` - Optional String  Defaults to `EmptyStatusCode`<br>Possible values are `EmptyStatusCode`, `Continue`, `OK`, `Created`, `Accepted`, `NonAuthoritativeInformation`, `NoContent`, `ResetContent`, `PartialContent`, `MultiStatus`, `AlreadyReported`, `IMUsed`, `MultipleChoices`, `MovedPermanently`, `Found`, `SeeOther`, `NotModified`, `UseProxy`, `TemporaryRedirect`, `PermanentRedirect`, `BadRequest`, `Unauthorized`, `PaymentRequired`, `Forbidden`, `NotFound`, `MethodNotAllowed`, `NotAcceptable`, `ProxyAuthenticationRequired`, `RequestTimeout`, `Conflict`, `Gone`, `LengthRequired`, `PreconditionFailed`, `PayloadTooLarge`, `URITooLong`, `UnsupportedMediaType`, `RangeNotSatisfiable`, `ExpectationFailed`, `MisdirectedRequest`, `UnprocessableEntity`, `Locked`, `FailedDependency`, `UpgradeRequired`, `PreconditionRequired`, `TooManyRequests`, `RequestHeaderFieldsTooLarge`, `InternalServerError`, `NotImplemented`, `BadGateway`, `ServiceUnavailable`, `GatewayTimeout`, `HTTPVersionNotSupported`, `VariantAlsoNegotiates`, `InsufficientStorage`, `LoopDetected`, `NotExtended`, `NetworkAuthenticationRequired`<br>HTTP Status Code. HTTP response status codes EmptyStatusCode response codes means it is not specified Continue status code OK status code Created status code Accepted status code Non Authoritative Information status code No Content status code Reset Content status code Partial Content status code Multi Status status code Already Reported status code Im Used status code Multiple Choices status code Moved Permanently status code Found status code See Other status code Not Modified status code Use Proxy status code Temporary Redirect status code Permanent Redirect status code Bad Request status code Unauthorized status code Payment Required status code Forbidden status code Not Found status code Method Not Allowed status code Not Acceptable status code Proxy Authentication Required status code Request Timeout status code Conflict status code Gone status code Length Required status code Precondition Failed status code Payload Too Large status code URI Too Long status code Unsupported Media Type status code Range Not Satisfiable status code Expectation Failed status code Misdirected Request status code Unprocessable Entity status code Locked status code Failed Dependency status code Upgrade Required status code Precondition Required status code Too Many Requests status code Request Header Fields Too Large status code Internal Server Error status code Not Implemented status code Bad Gateway status code Service Unavailable status code Gateway Timeout status code HTTP Version Not Supported status code Variant Also Negotiates status code Insufficient Storage status code Loop Detected status code Not Extended status code Network Authentication Required status code

<a id="bot-defense-policy-protected-app-endpoints-mitigation-flag"></a>

**Bot Defense Policy Protected App Endpoints Mitigation Flag**

&#x2022; `append_headers` - Optional Block<br>Append Flag Mitigation Headers. Append flag mitigation headers to forwarded request<br>See [Append Headers](#bot-defense-policy-protected-app-endpoints-mitigation-flag-append-headers) below.

&#x2022; `no_headers` - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="bot-defense-policy-protected-app-endpoints-mitigation-flag-append-headers"></a>

**Bot Defense Policy Protected App Endpoints Mitigation Flag Append Headers**

&#x2022; `auto_type_header_name` - Optional String<br>Automation Type Header Name. A case-insensitive HTTP header name

&#x2022; `inference_header_name` - Optional String<br>Inference Header Name. A case-insensitive HTTP header name

<a id="bot-defense-policy-protected-app-endpoints-mitigation-redirect"></a>

**Bot Defense Policy Protected App Endpoints Mitigation Redirect**

&#x2022; `uri` - Optional String<br>URI. URI location for redirect may be relative or absolute

<a id="bot-defense-policy-protected-app-endpoints-path"></a>

**Bot Defense Policy Protected App Endpoints Path**

&#x2022; `path` - Optional String<br>Exact. Exact path value to match

&#x2022; `prefix` - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

&#x2022; `regex` - Optional String<br>Regex. Regular expression of path match (e.g. the value .* will match on all paths)

<a id="bot-defense-policy-protected-app-endpoints-query-params"></a>

**Bot Defense Policy Protected App Endpoints Query Params**

&#x2022; `check_not_present` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `check_present` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `invert_matcher` - Optional Bool<br>Invert Query Parameter Matcher. Invert the match result

&#x2022; `item` - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#bot-defense-policy-protected-app-endpoints-query-params-item) below.

&#x2022; `key` - Optional String<br>Query Parameter Name. A case-sensitive HTTP query parameter name

<a id="bot-defense-policy-protected-app-endpoints-query-params-item"></a>

**Bot Defense Policy Protected App Endpoints Query Params Item**

&#x2022; `exact_values` - Optional List<br>Exact Values. A list of exact values to match the input against

&#x2022; `regex_values` - Optional List<br>Regex Values. A list of regular expressions to match the input against

&#x2022; `transformers` - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

<a id="bot-defense-policy-protected-app-endpoints-web-mobile"></a>

**Bot Defense Policy Protected App Endpoints Web Mobile**

&#x2022; `mobile_identifier` - Optional String  Defaults to `HEADERS`<br>Mobile Identifier. Mobile identifier type - HEADERS: Headers Headers. The only possible value is `HEADERS`

<a id="bot-defense-advanced"></a>

**Bot Defense Advanced**

&#x2022; `disable_js_insert` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `disable_mobile_sdk` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `js_insert_all_pages` - Optional Block<br>Insert Bot Defense JavaScript in All Pages. Insert Bot Defense JavaScript in all pages<br>See [Js Insert All Pages](#bot-defense-advanced-js-insert-all-pages) below.

&#x2022; `js_insert_all_pages_except` - Optional Block<br>Insert JavaScript in All Pages with the Exceptions. Insert Bot Defense JavaScript in all pages with the exceptions<br>See [Js Insert All Pages Except](#bot-defense-advanced-js-insert-all-pages-except) below.

&#x2022; `js_insertion_rules` - Optional Block<br>JavaScript Custom Insertion Rules. This defines custom JavaScript insertion rules for Bot Defense Policy<br>See [Js Insertion Rules](#bot-defense-advanced-js-insertion-rules) below.

&#x2022; `mobile` - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Mobile](#bot-defense-advanced-mobile) below.

&#x2022; `mobile_sdk_config` - Optional Block<br>Mobile Request Identifier Headers. Mobile Request Identifier Headers<br>See [Mobile Sdk Config](#bot-defense-advanced-mobile-sdk-config) below.

&#x2022; `web` - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Web](#bot-defense-advanced-web) below.

<a id="bot-defense-advanced-js-insert-all-pages"></a>

**Bot Defense Advanced Js Insert All Pages**

&#x2022; `javascript_location` - Optional String  Defaults to `AFTER_HEAD`<br>Possible values are `AFTER_HEAD`, `AFTER_TITLE_END`, `BEFORE_SCRIPT`<br>JavaScript Location. All inside networks. Insert JavaScript after <head> tag Insert JavaScript after </title> tag. Insert JavaScript before first <script> tag

<a id="bot-defense-advanced-js-insert-all-pages-except"></a>

**Bot Defense Advanced Js Insert All Pages Except**

&#x2022; `exclude_list` - Optional Block<br>Exclude Pages. Optional JavaScript insertions exclude list of domain and path matchers<br>See [Exclude List](#bot-defense-advanced-js-insert-all-pages-except-exclude-list) below.

&#x2022; `javascript_location` - Optional String  Defaults to `AFTER_HEAD`<br>Possible values are `AFTER_HEAD`, `AFTER_TITLE_END`, `BEFORE_SCRIPT`<br>JavaScript Location. All inside networks. Insert JavaScript after <head> tag Insert JavaScript after </title> tag. Insert JavaScript before first <script> tag

<a id="bot-defense-advanced-js-insert-all-pages-except-exclude-list"></a>

**Bot Defense Advanced Js Insert All Pages Except Exclude List**

&#x2022; `any_domain` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `domain` - Optional Block<br>Domains. Domains names<br>See [Domain](#bot-defense-advanced-js-insert-all-pages-except-exclude-list-domain) below.

&#x2022; `metadata` - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#bot-defense-advanced-js-insert-all-pages-except-exclude-list-metadata) below.

&#x2022; `path` - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#bot-defense-advanced-js-insert-all-pages-except-exclude-list-path) below.

<a id="bot-defense-advanced-js-insert-all-pages-except-exclude-list-domain"></a>

**Bot Defense Advanced Js Insert All Pages Except Exclude List Domain**

&#x2022; `exact_value` - Optional String<br>Exact Value. Exact domain name

&#x2022; `regex_value` - Optional String<br>Regex Values of Domains. Regular Expression value for the domain name

&#x2022; `suffix_value` - Optional String<br>Suffix Value. Suffix of domain name e.g 'xyz.com' will match '*.xyz.com' and 'xyz.com'

<a id="bot-defense-advanced-js-insert-all-pages-except-exclude-list-metadata"></a>

**Bot Defense Advanced Js Insert All Pages Except Exclude List Metadata**

&#x2022; `description` - Optional String<br>Description. Human readable description

&#x2022; `name` - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

<a id="bot-defense-advanced-js-insert-all-pages-except-exclude-list-path"></a>

**Bot Defense Advanced Js Insert All Pages Except Exclude List Path**

&#x2022; `path` - Optional String<br>Exact. Exact path value to match

&#x2022; `prefix` - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

&#x2022; `regex` - Optional String<br>Regex. Regular expression of path match (e.g. the value .* will match on all paths)

<a id="bot-defense-advanced-js-insertion-rules"></a>

**Bot Defense Advanced Js Insertion Rules**

&#x2022; `exclude_list` - Optional Block<br>Exclude Paths. Optional JavaScript insertions exclude list of domain and path matchers<br>See [Exclude List](#bot-defense-advanced-js-insertion-rules-exclude-list) below.

&#x2022; `rules` - Optional Block<br>JavaScript Insertions. Required list of pages to insert Bot Defense client JavaScript<br>See [Rules](#bot-defense-advanced-js-insertion-rules-rules) below.

<a id="bot-defense-advanced-js-insertion-rules-exclude-list"></a>

**Bot Defense Advanced Js Insertion Rules Exclude List**

&#x2022; `any_domain` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `domain` - Optional Block<br>Domains. Domains names<br>See [Domain](#bot-defense-advanced-js-insertion-rules-exclude-list-domain) below.

&#x2022; `metadata` - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#bot-defense-advanced-js-insertion-rules-exclude-list-metadata) below.

&#x2022; `path` - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#bot-defense-advanced-js-insertion-rules-exclude-list-path) below.

<a id="bot-defense-advanced-js-insertion-rules-exclude-list-domain"></a>

**Bot Defense Advanced Js Insertion Rules Exclude List Domain**

&#x2022; `exact_value` - Optional String<br>Exact Value. Exact domain name

&#x2022; `regex_value` - Optional String<br>Regex Values of Domains. Regular Expression value for the domain name

&#x2022; `suffix_value` - Optional String<br>Suffix Value. Suffix of domain name e.g 'xyz.com' will match '*.xyz.com' and 'xyz.com'

<a id="bot-defense-advanced-js-insertion-rules-exclude-list-metadata"></a>

**Bot Defense Advanced Js Insertion Rules Exclude List Metadata**

&#x2022; `description` - Optional String<br>Description. Human readable description

&#x2022; `name` - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

<a id="bot-defense-advanced-js-insertion-rules-exclude-list-path"></a>

**Bot Defense Advanced Js Insertion Rules Exclude List Path**

&#x2022; `path` - Optional String<br>Exact. Exact path value to match

&#x2022; `prefix` - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

&#x2022; `regex` - Optional String<br>Regex. Regular expression of path match (e.g. the value .* will match on all paths)

<a id="bot-defense-advanced-js-insertion-rules-rules"></a>

**Bot Defense Advanced Js Insertion Rules Rules**

&#x2022; `any_domain` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `domain` - Optional Block<br>Domains. Domains names<br>See [Domain](#bot-defense-advanced-js-insertion-rules-rules-domain) below.

&#x2022; `javascript_location` - Optional String  Defaults to `AFTER_HEAD`<br>Possible values are `AFTER_HEAD`, `AFTER_TITLE_END`, `BEFORE_SCRIPT`<br>JavaScript Location. All inside networks. Insert JavaScript after <head> tag Insert JavaScript after </title> tag. Insert JavaScript before first <script> tag

&#x2022; `metadata` - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#bot-defense-advanced-js-insertion-rules-rules-metadata) below.

&#x2022; `path` - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#bot-defense-advanced-js-insertion-rules-rules-path) below.

<a id="bot-defense-advanced-js-insertion-rules-rules-domain"></a>

**Bot Defense Advanced Js Insertion Rules Rules Domain**

&#x2022; `exact_value` - Optional String<br>Exact Value. Exact domain name

&#x2022; `regex_value` - Optional String<br>Regex Values of Domains. Regular Expression value for the domain name

&#x2022; `suffix_value` - Optional String<br>Suffix Value. Suffix of domain name e.g 'xyz.com' will match '*.xyz.com' and 'xyz.com'

<a id="bot-defense-advanced-js-insertion-rules-rules-metadata"></a>

**Bot Defense Advanced Js Insertion Rules Rules Metadata**

&#x2022; `description` - Optional String<br>Description. Human readable description

&#x2022; `name` - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

<a id="bot-defense-advanced-js-insertion-rules-rules-path"></a>

**Bot Defense Advanced Js Insertion Rules Rules Path**

&#x2022; `path` - Optional String<br>Exact. Exact path value to match

&#x2022; `prefix` - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

&#x2022; `regex` - Optional String<br>Regex. Regular expression of path match (e.g. the value .* will match on all paths)

<a id="bot-defense-advanced-mobile"></a>

**Bot Defense Advanced Mobile**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="bot-defense-advanced-mobile-sdk-config"></a>

**Bot Defense Advanced Mobile Sdk Config**

&#x2022; `mobile_identifier` - Optional Block<br>Mobile Traffic Identifier. Mobile traffic identifier type<br>See [Mobile Identifier](#bot-defense-advanced-mobile-sdk-config-mobile-identifier) below.

<a id="bot-defense-advanced-mobile-sdk-config-mobile-identifier"></a>

**Bot Defense Advanced Mobile Sdk Config Mobile Identifier**

&#x2022; `headers` - Optional Block<br>Headers. Headers that can be used to identify mobile traffic<br>See [Headers](#bot-defense-advanced-mobile-sdk-config-mobile-identifier-headers) below.

<a id="bot-defense-advanced-mobile-sdk-config-mobile-identifier-headers"></a>

**Bot Defense Advanced Mobile Sdk Config Mobile Identifier Headers**

&#x2022; `check_not_present` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `check_present` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `item` - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#bot-defense-advanced-mobile-sdk-config-mobile-identifier-headers-item) below.

&#x2022; `name` - Optional String<br>Header Name. A case-insensitive HTTP header name

<a id="bot-defense-advanced-mobile-sdk-config-mobile-identifier-headers-item"></a>

**Bot Defense Advanced Mobile Sdk Config Mobile Identifier Headers Item**

&#x2022; `exact_values` - Optional List<br>Exact Values. A list of exact values to match the input against

&#x2022; `regex_values` - Optional List<br>Regex Values. A list of regular expressions to match the input against

&#x2022; `transformers` - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

<a id="bot-defense-advanced-web"></a>

**Bot Defense Advanced Web**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="caching-policy"></a>

**Caching Policy**

&#x2022; `custom_cache_rule` - Optional Block<br>Custom Cache Rules. Caching policies for CDN<br>See [Custom Cache Rule](#caching-policy-custom-cache-rule) below.

&#x2022; `default_cache_action` - Optional Block<br>Default Cache Behaviour. This defines a Default Cache Action<br>See [Default Cache Action](#caching-policy-default-cache-action) below.

<a id="caching-policy-custom-cache-rule"></a>

**Caching Policy Custom Cache Rule**

&#x2022; `cdn_cache_rules` - Optional Block<br>CDN Cache Rule. Reference to CDN Cache Rule configuration object<br>See [CDN Cache Rules](#caching-policy-custom-cache-rule-cdn-cache-rules) below.

<a id="caching-policy-custom-cache-rule-cdn-cache-rules"></a>

**Caching Policy Custom Cache Rule CDN Cache Rules**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="caching-policy-default-cache-action"></a>

**Caching Policy Default Cache Action**

&#x2022; `cache_disabled` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `cache_ttl_default` - Optional String<br>Fallback Cache TTL (d/ h/ m). Use Cache TTL Provided by Origin, and set a contigency TTL value in case one is not provided

&#x2022; `cache_ttl_override` - Optional String<br>Override Cache TTL (d/ h/ m/ s). Always override the Cahce TTL provided by Origin

<a id="captcha-challenge"></a>

**Captcha Challenge**

&#x2022; `cookie_expiry` - Optional Number<br>Cookie Expiration Period. Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new challenge

&#x2022; `custom_page` - Optional String<br>Custom message for Captcha Challenge. Custom message is of type uri_ref. Currently supported URL schemes is string:///. For string:/// scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. 'Please Wait.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '<p> Please Wait </p>'. Base64 encoded string for this HTML is 'PHA+IFBsZWFzZSBXYWl0IDwvcD4='

<a id="client-side-defense"></a>

**Client Side Defense**

&#x2022; `policy` - Optional Block<br>Client-Side Defense Policy. This defines various configuration options for Client-Side Defense policy<br>See [Policy](#client-side-defense-policy) below.

<a id="client-side-defense-policy"></a>

**Client Side Defense Policy**

&#x2022; `disable_js_insert` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `js_insert_all_pages` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `js_insert_all_pages_except` - Optional Block<br>Insert JavaScript in All Pages with the Exceptions. Insert Client-Side Defense JavaScript in all pages with the exceptions<br>See [Js Insert All Pages Except](#client-side-defense-policy-js-insert-all-pages-except) below.

&#x2022; `js_insertion_rules` - Optional Block<br>JavaScript Custom Insertion Rules. This defines custom JavaScript insertion rules for Client-Side Defense Policy<br>See [Js Insertion Rules](#client-side-defense-policy-js-insertion-rules) below.

<a id="client-side-defense-policy-js-insert-all-pages-except"></a>

**Client Side Defense Policy Js Insert All Pages Except**

&#x2022; `exclude_list` - Optional Block<br>Exclude Pages. Optional JavaScript insertions exclude list of domain and path matchers<br>See [Exclude List](#client-side-defense-policy-js-insert-all-pages-except-exclude-list) below.

<a id="client-side-defense-policy-js-insert-all-pages-except-exclude-list"></a>

**Client Side Defense Policy Js Insert All Pages Except Exclude List**

&#x2022; `any_domain` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `domain` - Optional Block<br>Domains. Domains names<br>See [Domain](#client-side-defense-policy-js-insert-all-pages-except-exclude-list-domain) below.

&#x2022; `metadata` - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#client-side-defense-policy-js-insert-all-pages-except-exclude-list-metadata) below.

&#x2022; `path` - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#client-side-defense-policy-js-insert-all-pages-except-exclude-list-path) below.

<a id="client-side-defense-policy-js-insert-all-pages-except-exclude-list-domain"></a>

**Client Side Defense Policy Js Insert All Pages Except Exclude List Domain**

&#x2022; `exact_value` - Optional String<br>Exact Value. Exact domain name

&#x2022; `regex_value` - Optional String<br>Regex Values of Domains. Regular Expression value for the domain name

&#x2022; `suffix_value` - Optional String<br>Suffix Value. Suffix of domain name e.g 'xyz.com' will match '*.xyz.com' and 'xyz.com'

<a id="client-side-defense-policy-js-insert-all-pages-except-exclude-list-metadata"></a>

**Client Side Defense Policy Js Insert All Pages Except Exclude List Metadata**

&#x2022; `description` - Optional String<br>Description. Human readable description

&#x2022; `name` - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

<a id="client-side-defense-policy-js-insert-all-pages-except-exclude-list-path"></a>

**Client Side Defense Policy Js Insert All Pages Except Exclude List Path**

&#x2022; `path` - Optional String<br>Exact. Exact path value to match

&#x2022; `prefix` - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

&#x2022; `regex` - Optional String<br>Regex. Regular expression of path match (e.g. the value .* will match on all paths)

<a id="client-side-defense-policy-js-insertion-rules"></a>

**Client Side Defense Policy Js Insertion Rules**

&#x2022; `exclude_list` - Optional Block<br>Exclude Paths. Optional JavaScript insertions exclude list of domain and path matchers<br>See [Exclude List](#client-side-defense-policy-js-insertion-rules-exclude-list) below.

&#x2022; `rules` - Optional Block<br>JavaScript Insertions. Required list of pages to insert Client-Side Defense client JavaScript<br>See [Rules](#client-side-defense-policy-js-insertion-rules-rules) below.

<a id="client-side-defense-policy-js-insertion-rules-exclude-list"></a>

**Client Side Defense Policy Js Insertion Rules Exclude List**

&#x2022; `any_domain` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `domain` - Optional Block<br>Domains. Domains names<br>See [Domain](#client-side-defense-policy-js-insertion-rules-exclude-list-domain) below.

&#x2022; `metadata` - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#client-side-defense-policy-js-insertion-rules-exclude-list-metadata) below.

&#x2022; `path` - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#client-side-defense-policy-js-insertion-rules-exclude-list-path) below.

<a id="client-side-defense-policy-js-insertion-rules-exclude-list-domain"></a>

**Client Side Defense Policy Js Insertion Rules Exclude List Domain**

&#x2022; `exact_value` - Optional String<br>Exact Value. Exact domain name

&#x2022; `regex_value` - Optional String<br>Regex Values of Domains. Regular Expression value for the domain name

&#x2022; `suffix_value` - Optional String<br>Suffix Value. Suffix of domain name e.g 'xyz.com' will match '*.xyz.com' and 'xyz.com'

<a id="client-side-defense-policy-js-insertion-rules-exclude-list-metadata"></a>

**Client Side Defense Policy Js Insertion Rules Exclude List Metadata**

&#x2022; `description` - Optional String<br>Description. Human readable description

&#x2022; `name` - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

<a id="client-side-defense-policy-js-insertion-rules-exclude-list-path"></a>

**Client Side Defense Policy Js Insertion Rules Exclude List Path**

&#x2022; `path` - Optional String<br>Exact. Exact path value to match

&#x2022; `prefix` - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

&#x2022; `regex` - Optional String<br>Regex. Regular expression of path match (e.g. the value .* will match on all paths)

<a id="client-side-defense-policy-js-insertion-rules-rules"></a>

**Client Side Defense Policy Js Insertion Rules Rules**

&#x2022; `any_domain` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `domain` - Optional Block<br>Domains. Domains names<br>See [Domain](#client-side-defense-policy-js-insertion-rules-rules-domain) below.

&#x2022; `metadata` - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#client-side-defense-policy-js-insertion-rules-rules-metadata) below.

&#x2022; `path` - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#client-side-defense-policy-js-insertion-rules-rules-path) below.

<a id="client-side-defense-policy-js-insertion-rules-rules-domain"></a>

**Client Side Defense Policy Js Insertion Rules Rules Domain**

&#x2022; `exact_value` - Optional String<br>Exact Value. Exact domain name

&#x2022; `regex_value` - Optional String<br>Regex Values of Domains. Regular Expression value for the domain name

&#x2022; `suffix_value` - Optional String<br>Suffix Value. Suffix of domain name e.g 'xyz.com' will match '*.xyz.com' and 'xyz.com'

<a id="client-side-defense-policy-js-insertion-rules-rules-metadata"></a>

**Client Side Defense Policy Js Insertion Rules Rules Metadata**

&#x2022; `description` - Optional String<br>Description. Human readable description

&#x2022; `name` - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

<a id="client-side-defense-policy-js-insertion-rules-rules-path"></a>

**Client Side Defense Policy Js Insertion Rules Rules Path**

&#x2022; `path` - Optional String<br>Exact. Exact path value to match

&#x2022; `prefix` - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

&#x2022; `regex` - Optional String<br>Regex. Regular expression of path match (e.g. the value .* will match on all paths)

<a id="cookie-stickiness"></a>

**Cookie Stickiness**

&#x2022; `add_httponly` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `add_secure` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `ignore_httponly` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `ignore_samesite` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `ignore_secure` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `name` - Optional String<br>Name. The name of the cookie that will be used to obtain the hash key. If the cookie is not present and TTL below is not set, no hash will be produced

&#x2022; `path` - Optional String<br>Path. The name of the path for the cookie. If no path is specified here, no path will be set for the cookie

&#x2022; `samesite_lax` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `samesite_none` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `samesite_strict` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `ttl` - Optional Number<br>TTL. If specified, a cookie with the TTL will be generated if the cookie is not present. If the TTL is present and zero, the generated cookie will be a session cookie. TTL value is in milliseconds

<a id="cors-policy"></a>

**CORS Policy**

&#x2022; `allow_credentials` - Optional Bool<br>Allow Credentials. Specifies whether the resource allows credentials

&#x2022; `allow_headers` - Optional String<br>Allow Headers. Specifies the content for the access-control-allow-headers header

&#x2022; `allow_methods` - Optional String<br>Allow Methods. Specifies the content for the access-control-allow-methods header

&#x2022; `allow_origin` - Optional List<br>Allow Origin. Specifies the origins that will be allowed to do CORS requests. An origin is allowed if either allow_origin or allow_origin_regex match

&#x2022; `allow_origin_regex` - Optional List<br>Allow Origin Regex. Specifies regex patterns that match allowed origins. An origin is allowed if either allow_origin or allow_origin_regex match

&#x2022; `disabled` - Optional Bool<br>Disabled. Disable the CorsPolicy for a particular route. This is useful when virtual-host has CorsPolicy, but we need to disable it on a specific route. The value of this field is ignored for virtual-host

&#x2022; `expose_headers` - Optional String<br>Expose Headers. Specifies the content for the access-control-expose-headers header

&#x2022; `maximum_age` - Optional Number<br>Maximum Age. Specifies the content for the access-control-max-age header in seconds. This indicates the maximum number of seconds the results can be cached A value of -1 will disable caching. Maximum permitted value is 86400 seconds (24 hours)

<a id="csrf-policy"></a>

**CSRF Policy**

&#x2022; `all_load_balancer_domains` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `custom_domain_list` - Optional Block<br>Domain name list. List of domain names used for Host header matching<br>See [Custom Domain List](#csrf-policy-custom-domain-list) below.

&#x2022; `disabled` - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="csrf-policy-custom-domain-list"></a>

**CSRF Policy Custom Domain List**

&#x2022; `domains` - Optional List<br>Domain names. A list of domain names that will be matched to loadbalancer. These domains are not used for SNI match. Wildcard names are supported in the suffix or prefix form

<a id="data-guard-rules"></a>

**Data Guard Rules**

&#x2022; `any_domain` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `apply_data_guard` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `exact_value` - Optional String<br>Exact Value. Exact domain name

&#x2022; `metadata` - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#data-guard-rules-metadata) below.

&#x2022; `path` - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#data-guard-rules-path) below.

&#x2022; `skip_data_guard` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `suffix_value` - Optional String<br>Suffix Value. Suffix of domain name e.g 'xyz.com' will match '*.xyz.com' and 'xyz.com'

<a id="data-guard-rules-metadata"></a>

**Data Guard Rules Metadata**

&#x2022; `description` - Optional String<br>Description. Human readable description

&#x2022; `name` - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

<a id="data-guard-rules-path"></a>

**Data Guard Rules Path**

&#x2022; `path` - Optional String<br>Exact. Exact path value to match

&#x2022; `prefix` - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

&#x2022; `regex` - Optional String<br>Regex. Regular expression of path match (e.g. the value .* will match on all paths)

<a id="ddos-mitigation-rules"></a>

**DDOS Mitigation Rules**

&#x2022; `block` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `ddos_client_source` - Optional Block<br>DDOS Client Source Choice. DDOS Mitigation sources to be blocked<br>See [DDOS Client Source](#ddos-mitigation-rules-ddos-client-source) below.

&#x2022; `expiration_timestamp` - Optional String<br>Expiration Timestamp. The expiration_timestamp is the RFC 3339 format timestamp at which the containing rule is considered to be logically expired. The rule continues to exist in the configuration but is not applied anymore

&#x2022; `ip_prefix_list` - Optional Block<br>IP Prefix Match List. List of IP Prefix strings to match against<br>See [IP Prefix List](#ddos-mitigation-rules-ip-prefix-list) below.

&#x2022; `metadata` - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#ddos-mitigation-rules-metadata) below.

<a id="ddos-mitigation-rules-ddos-client-source"></a>

**DDOS Mitigation Rules DDOS Client Source**

&#x2022; `asn_list` - Optional Block<br>ASN Match List. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer<br>See [Asn List](#ddos-mitigation-rules-ddos-client-source-asn-list) below.

&#x2022; `country_list` - Optional List  Defaults to `COUNTRY_NONE`<br>Possible values are `COUNTRY_NONE`, `COUNTRY_AD`, `COUNTRY_AE`, `COUNTRY_AF`, `COUNTRY_AG`, `COUNTRY_AI`, `COUNTRY_AL`, `COUNTRY_AM`, `COUNTRY_AN`, `COUNTRY_AO`, `COUNTRY_AQ`, `COUNTRY_AR`, `COUNTRY_AS`, `COUNTRY_AT`, `COUNTRY_AU`, `COUNTRY_AW`, `COUNTRY_AX`, `COUNTRY_AZ`, `COUNTRY_BA`, `COUNTRY_BB`, `COUNTRY_BD`, `COUNTRY_BE`, `COUNTRY_BF`, `COUNTRY_BG`, `COUNTRY_BH`, `COUNTRY_BI`, `COUNTRY_BJ`, `COUNTRY_BL`, `COUNTRY_BM`, `COUNTRY_BN`, `COUNTRY_BO`, `COUNTRY_BQ`, `COUNTRY_BR`, `COUNTRY_BS`, `COUNTRY_BT`, `COUNTRY_BV`, `COUNTRY_BW`, `COUNTRY_BY`, `COUNTRY_BZ`, `COUNTRY_CA`, `COUNTRY_CC`, `COUNTRY_CD`, `COUNTRY_CF`, `COUNTRY_CG`, `COUNTRY_CH`, `COUNTRY_CI`, `COUNTRY_CK`, `COUNTRY_CL`, `COUNTRY_CM`, `COUNTRY_CN`, `COUNTRY_CO`, `COUNTRY_CR`, `COUNTRY_CS`, `COUNTRY_CU`, `COUNTRY_CV`, `COUNTRY_CW`, `COUNTRY_CX`, `COUNTRY_CY`, `COUNTRY_CZ`, `COUNTRY_DE`, `COUNTRY_DJ`, `COUNTRY_DK`, `COUNTRY_DM`, `COUNTRY_DO`, `COUNTRY_DZ`, `COUNTRY_EC`, `COUNTRY_EE`, `COUNTRY_EG`, `COUNTRY_EH`, `COUNTRY_ER`, `COUNTRY_ES`, `COUNTRY_ET`, `COUNTRY_FI`, `COUNTRY_FJ`, `COUNTRY_FK`, `COUNTRY_FM`, `COUNTRY_FO`, `COUNTRY_FR`, `COUNTRY_GA`, `COUNTRY_GB`, `COUNTRY_GD`, `COUNTRY_GE`, `COUNTRY_GF`, `COUNTRY_GG`, `COUNTRY_GH`, `COUNTRY_GI`, `COUNTRY_GL`, `COUNTRY_GM`, `COUNTRY_GN`, `COUNTRY_GP`, `COUNTRY_GQ`, `COUNTRY_GR`, `COUNTRY_GS`, `COUNTRY_GT`, `COUNTRY_GU`, `COUNTRY_GW`, `COUNTRY_GY`, `COUNTRY_HK`, `COUNTRY_HM`, `COUNTRY_HN`, `COUNTRY_HR`, `COUNTRY_HT`, `COUNTRY_HU`, `COUNTRY_ID`, `COUNTRY_IE`, `COUNTRY_IL`, `COUNTRY_IM`, `COUNTRY_IN`, `COUNTRY_IO`, `COUNTRY_IQ`, `COUNTRY_IR`, `COUNTRY_IS`, `COUNTRY_IT`, `COUNTRY_JE`, `COUNTRY_JM`, `COUNTRY_JO`, `COUNTRY_JP`, `COUNTRY_KE`, `COUNTRY_KG`, `COUNTRY_KH`, `COUNTRY_KI`, `COUNTRY_KM`, `COUNTRY_KN`, `COUNTRY_KP`, `COUNTRY_KR`, `COUNTRY_KW`, `COUNTRY_KY`, `COUNTRY_KZ`, `COUNTRY_LA`, `COUNTRY_LB`, `COUNTRY_LC`, `COUNTRY_LI`, `COUNTRY_LK`, `COUNTRY_LR`, `COUNTRY_LS`, `COUNTRY_LT`, `COUNTRY_LU`, `COUNTRY_LV`, `COUNTRY_LY`, `COUNTRY_MA`, `COUNTRY_MC`, `COUNTRY_MD`, `COUNTRY_ME`, `COUNTRY_MF`, `COUNTRY_MG`, `COUNTRY_MH`, `COUNTRY_MK`, `COUNTRY_ML`, `COUNTRY_MM`, `COUNTRY_MN`, `COUNTRY_MO`, `COUNTRY_MP`, `COUNTRY_MQ`, `COUNTRY_MR`, `COUNTRY_MS`, `COUNTRY_MT`, `COUNTRY_MU`, `COUNTRY_MV`, `COUNTRY_MW`, `COUNTRY_MX`, `COUNTRY_MY`, `COUNTRY_MZ`, `COUNTRY_NA`, `COUNTRY_NC`, `COUNTRY_NE`, `COUNTRY_NF`, `COUNTRY_NG`, `COUNTRY_NI`, `COUNTRY_NL`, `COUNTRY_NO`, `COUNTRY_NP`, `COUNTRY_NR`, `COUNTRY_NU`, `COUNTRY_NZ`, `COUNTRY_OM`, `COUNTRY_PA`, `COUNTRY_PE`, `COUNTRY_PF`, `COUNTRY_PG`, `COUNTRY_PH`, `COUNTRY_PK`, `COUNTRY_PL`, `COUNTRY_PM`, `COUNTRY_PN`, `COUNTRY_PR`, `COUNTRY_PS`, `COUNTRY_PT`, `COUNTRY_PW`, `COUNTRY_PY`, `COUNTRY_QA`, `COUNTRY_RE`, `COUNTRY_RO`, `COUNTRY_RS`, `COUNTRY_RU`, `COUNTRY_RW`, `COUNTRY_SA`, `COUNTRY_SB`, `COUNTRY_SC`, `COUNTRY_SD`, `COUNTRY_SE`, `COUNTRY_SG`, `COUNTRY_SH`, `COUNTRY_SI`, `COUNTRY_SJ`, `COUNTRY_SK`, `COUNTRY_SL`, `COUNTRY_SM`, `COUNTRY_SN`, `COUNTRY_SO`, `COUNTRY_SR`, `COUNTRY_SS`, `COUNTRY_ST`, `COUNTRY_SV`, `COUNTRY_SX`, `COUNTRY_SY`, `COUNTRY_SZ`, `COUNTRY_TC`, `COUNTRY_TD`, `COUNTRY_TF`, `COUNTRY_TG`, `COUNTRY_TH`, `COUNTRY_TJ`, `COUNTRY_TK`, `COUNTRY_TL`, `COUNTRY_TM`, `COUNTRY_TN`, `COUNTRY_TO`, `COUNTRY_TR`, `COUNTRY_TT`, `COUNTRY_TV`, `COUNTRY_TW`, `COUNTRY_TZ`, `COUNTRY_UA`, `COUNTRY_UG`, `COUNTRY_UM`, `COUNTRY_US`, `COUNTRY_UY`, `COUNTRY_UZ`, `COUNTRY_VA`, `COUNTRY_VC`, `COUNTRY_VE`, `COUNTRY_VG`, `COUNTRY_VI`, `COUNTRY_VN`, `COUNTRY_VU`, `COUNTRY_WF`, `COUNTRY_WS`, `COUNTRY_XK`, `COUNTRY_XT`, `COUNTRY_YE`, `COUNTRY_YT`, `COUNTRY_ZA`, `COUNTRY_ZM`, `COUNTRY_ZW`<br>Country List. Sources that are located in one of the countries in the given list

&#x2022; `ja4_tls_fingerprint_matcher` - Optional Block<br>JA4 TLS Fingerprint Matcher. An extended version of JA3 that includes additional fields for more comprehensive fingerprinting of SSL/TLS clients and potentially has a different structure and length<br>See [Ja4 TLS Fingerprint Matcher](#ddos-mitigation-rules-ddos-client-source-ja4-tls-fingerprint-matcher) below.

&#x2022; `tls_fingerprint_matcher` - Optional Block<br>TLS Fingerprint Matcher. A TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of supported positve match criteria includes a list of known classes of TLS fingerprints and a list of exact values. The match is considered successful if either of these positive criteria are satisfied and the input fingerprint is not one of the excluded values<br>See [TLS Fingerprint Matcher](#ddos-mitigation-rules-ddos-client-source-tls-fingerprint-matcher) below.

<a id="ddos-mitigation-rules-ddos-client-source-asn-list"></a>

**DDOS Mitigation Rules DDOS Client Source Asn List**

&#x2022; `as_numbers` - Optional List<br>AS Numbers. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer

<a id="ddos-mitigation-rules-ddos-client-source-ja4-tls-fingerprint-matcher"></a>

**DDOS Mitigation Rules DDOS Client Source Ja4 TLS Fingerprint Matcher**

&#x2022; `exact_values` - Optional List<br>Exact Values. A list of exact JA4 TLS fingerprint to match the input JA4 TLS fingerprint against

<a id="ddos-mitigation-rules-ddos-client-source-tls-fingerprint-matcher"></a>

**DDOS Mitigation Rules DDOS Client Source TLS Fingerprint Matcher**

&#x2022; `classes` - Optional List  Defaults to `TLS_FINGERPRINT_NONE`<br>Possible values are `TLS_FINGERPRINT_NONE`, `ANY_MALICIOUS_FINGERPRINT`, `ADWARE`, `ADWIND`, `DRIDEX`, `GOOTKIT`, `GOZI`, `JBIFROST`, `QUAKBOT`, `RANSOMWARE`, `TROLDESH`, `TOFSEE`, `TORRENTLOCKER`, `TRICKBOT`<br>TLS fingerprint classes. A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against

&#x2022; `exact_values` - Optional List<br>Exact Values. A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against

&#x2022; `excluded_values` - Optional List<br>Excluded Values. A list of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can be used to skip known false positives when using one or more known TLS fingerprint classes in the enclosing matcher

<a id="ddos-mitigation-rules-ip-prefix-list"></a>

**DDOS Mitigation Rules IP Prefix List**

&#x2022; `invert_match` - Optional Bool<br>Invert Match Result. Invert the match result

&#x2022; `ip_prefixes` - Optional List<br>IPv4 Prefix List. List of IPv4 prefix strings

<a id="ddos-mitigation-rules-metadata"></a>

**DDOS Mitigation Rules Metadata**

&#x2022; `description` - Optional String<br>Description. Human readable description

&#x2022; `name` - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

<a id="default-pool"></a>

**Default Pool**

&#x2022; `advanced_options` - Optional Block<br>Origin Pool Advanced Options. Configure Advanced options for origin pool<br>See [Advanced Options](#default-pool-advanced-options) below.

&#x2022; `automatic_port` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `endpoint_selection` - Optional String  Defaults to `DISTRIBUTED`<br>Possible values are `DISTRIBUTED`, `LOCAL_ONLY`, `LOCAL_PREFERRED`<br>Endpoint Selection Policy. Policy for selection of endpoints from local site/remote site/both Consider both remote and local endpoints for load balancing LOCAL_ONLY: Consider only local endpoints for load balancing Enable this policy to load balance ONLY among locally discovered endpoints Prefer the local endpoints for load balancing. If local endpoints are not present remote endpoints will be considered

&#x2022; `health_check_port` - Optional Number<br>Health check port. Port used for performing health check

&#x2022; `healthcheck` - Optional Block<br>Health Check object. Reference to healthcheck configuration objects<br>See [Healthcheck](#default-pool-healthcheck) below.

&#x2022; `lb_port` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `loadbalancer_algorithm` - Optional String  Defaults to `ROUND_ROBIN`<br>Possible values are `ROUND_ROBIN`, `LEAST_REQUEST`, `RING_HASH`, `RANDOM`, `LB_OVERRIDE`<br>Load Balancer Algorithm. Different load balancing algorithms supported When a connection to a endpoint in an upstream cluster is required, the load balancer uses loadbalancer_algorithm to determine which host is selected. - ROUND_ROBIN: ROUND_ROBIN Policy in which each healthy/available upstream endpoint is selected in round robin order. - LEAST_REQUEST: LEAST_REQUEST Policy in which loadbalancer picks the upstream endpoint which has the fewest active requests - RING_HASH: RING_HASH Policy implements consistent hashing to upstream endpoints using ring hash of endpoint names Hash of the incoming request is calculated using request hash policy. The ring/modulo hash load balancer implements consistent hashing to upstream hosts. The algorithm is based on mapping all hosts onto a circle such that the addition or removal of a host from the host set changes only affect 1/N requests. This technique is also commonly known as “ketama” hashing. A consistent hashing load balancer is only effective when protocol routing is used that specifies a value to hash on. The minimum ring size governs the replication factor for each host in the ring. For example, if the minimum ring size is 1024 and there are 16 hosts, each host will be replicated 64 times. - RANDOM: RANDOM Policy in which each available upstream endpoint is selected in random order. The random load balancer selects a random healthy host. The random load balancer generally performs better than round robin if no health checking policy is configured. Random selection avoids bias towards the host in the set that comes after a failed host. - LB_OVERRIDE: Load Balancer Override Hash policy is taken from from the load balancer which is using this origin pool

&#x2022; `no_tls` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `origin_servers` - Optional Block<br>Origin Servers. List of origin servers in this pool<br>See [Origin Servers](#default-pool-origin-servers) below.

&#x2022; `port` - Optional Number<br>Port. Endpoint service is available on this port

&#x2022; `same_as_endpoint_port` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `upstream_conn_pool_reuse_type` - Optional Block<br>Select upstream connection pool reuse state. Select upstream connection pool reuse state for every downstream connection. This configuration choice is for HTTP(S) LB only<br>See [Upstream Conn Pool Reuse Type](#default-pool-upstream-conn-pool-reuse-type) below.

&#x2022; `use_tls` - Optional Block<br>TLS Parameters for Origin Servers. Upstream TLS Parameters<br>See [Use TLS](#default-pool-use-tls) below.

&#x2022; `view_internal` - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [View Internal](#default-pool-view-internal) below.

<a id="default-pool-advanced-options"></a>

**Default Pool Advanced Options**

&#x2022; `auto_http_config` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `circuit_breaker` - Optional Block<br>Circuit Breaker. CircuitBreaker provides a mechanism for watching failures in upstream connections or requests and if the failures reach a certain threshold, automatically fail subsequent requests which allows to apply back pressure on downstream quickly<br>See [Circuit Breaker](#default-pool-advanced-options-circuit-breaker) below.

&#x2022; `connection_timeout` - Optional Number  Defaults to `2`  Specified in milliseconds<br>Connection Timeout. The timeout for new network connections to endpoints in the cluster.  The seconds

&#x2022; `default_circuit_breaker` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `disable_circuit_breaker` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `disable_lb_source_ip_persistance` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `disable_outlier_detection` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `disable_proxy_protocol` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `disable_subsets` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `enable_lb_source_ip_persistance` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `enable_subsets` - Optional Block<br>Origin Pool Subset Options. Configure subset options for origin pool<br>See [Enable Subsets](#default-pool-advanced-options-enable-subsets) below.

&#x2022; `http1_config` - Optional Block<br>HTTP/1.1 Protocol Options. HTTP/1.1 Protocol options for upstream connections<br>See [Http1 Config](#default-pool-advanced-options-http1-config) below.

&#x2022; `http2_options` - Optional Block<br>Http2 Protocol Options. Http2 Protocol options for upstream connections<br>See [Http2 Options](#default-pool-advanced-options-http2-options) below.

&#x2022; `http_idle_timeout` - Optional Number  Defaults to `5`  Specified in milliseconds<br>HTTP Idle Timeout. The idle timeout for upstream connection pool connections. The idle timeout is defined as the period in which there are no active requests. When the idle timeout is reached the connection will be closed. Note that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive.  The minutes

&#x2022; `no_panic_threshold` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `outlier_detection` - Optional Block<br>Outlier Detection. Outlier detection and ejection is the process of dynamically determining whether some number of hosts in an upstream cluster are performing unlike the others and removing them from the healthy load balancing set. Outlier detection is a form of passive health checking. Algorithm 1. A endpoint is determined to be an outlier (based on configured number of consecutive_5xx or consecutive_gateway_failures) . 2. If no endpoints have been ejected, loadbalancer will eject the host immediately. Otherwise, it checks to make sure the number of ejected hosts is below the allowed threshold (specified via max_ejection_percent setting). If the number of ejected hosts is above the threshold, the host is not ejected. 3. The endpoint is ejected for some number of milliseconds. Ejection means that the endpoint is marked unhealthy and will not be used during load balancing. The number of milliseconds is equal to the base_ejection_time value multiplied by the number of times the host has been ejected. 4. An ejected endpoint will automatically be brought back into service after the ejection time has been satisfied<br>See [Outlier Detection](#default-pool-advanced-options-outlier-detection) below.

&#x2022; `panic_threshold` - Optional Number<br>Panic threshold. x-example:'25' Configure a threshold (percentage of unhealthy endpoints) below which all endpoints will be considered for load balancing ignoring its health status

&#x2022; `proxy_protocol_v1` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `proxy_protocol_v2` - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="default-pool-advanced-options-circuit-breaker"></a>

**Default Pool Advanced Options Circuit Breaker**

&#x2022; `connection_limit` - Optional Number<br>Connection Limit. The maximum number of connections that loadbalancer will establish to all hosts in an upstream cluster. In practice this is only applicable to TCP and HTTP/1.1 clusters since HTTP/2 uses a single connection to each host. Remove endpoint out of load balancing decision, if number of connections reach connection limit

&#x2022; `max_requests` - Optional Number<br>Maximum Request Count. The maximum number of requests that can be outstanding to all hosts in a cluster at any given time. In practice this is applicable to HTTP/2 clusters since HTTP/1.1 clusters are governed by the maximum connections (connection_limit). Remove endpoint out of load balancing decision, if requests exceed this count

&#x2022; `pending_requests` - Optional Number<br>Pending Requests. The maximum number of requests that will be queued while waiting for a ready connection pool connection. Since HTTP/2 requests are sent over a single connection, this circuit breaker only comes into play as the initial connection is created, as requests will be multiplexed immediately afterwards. For HTTP/1.1, requests are added to the list of pending requests whenever there aren’t enough upstream connections available to immediately dispatch the request, so this circuit breaker will remain in play for the lifetime of the process. Remove endpoint out of load balancing decision, if pending request reach pending_request

&#x2022; `priority` - Optional String  Defaults to `DEFAULT`<br>Possible values are `DEFAULT`, `HIGH`<br>Routing Priority. Priority routing for each request. Different connection pools are used based on the priority selected for the request. Also, circuit-breaker configuration at destination cluster is chosen based on selected priority. Default routing mechanism High-Priority routing mechanism

&#x2022; `retries` - Optional Number<br>Retry Count. The maximum number of retries that can be outstanding to all hosts in a cluster at any given time. Remove endpoint out of load balancing decision, if retries for request exceed this count

<a id="default-pool-advanced-options-enable-subsets"></a>

**Default Pool Advanced Options Enable Subsets**

&#x2022; `any_endpoint` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `default_subset` - Optional Block<br>Origin Pool Default Subset. Default Subset definition<br>See [Default Subset](#default-pool-advanced-options-enable-subsets-default-subset) below.

&#x2022; `endpoint_subsets` - Optional Block<br>Origin Server Subsets Classes. List of subset class. Subsets class is defined using list of keys. Every unique combination of values of these keys form a subset withing the class<br>See [Endpoint Subsets](#default-pool-advanced-options-enable-subsets-endpoint-subsets) below.

&#x2022; `fail_request` - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="default-pool-advanced-options-enable-subsets-default-subset"></a>

**Default Pool Advanced Options Enable Subsets Default Subset**

&#x2022; `default_subset` - Optional Block<br>Default Subset for Origin Pool. List of key-value pairs that define default subset. which gets used when route specifies no metadata or no subset matching the metadata exists

<a id="default-pool-advanced-options-enable-subsets-endpoint-subsets"></a>

**Default Pool Advanced Options Enable Subsets Endpoint Subsets**

&#x2022; `keys` - Optional List<br>Keys. List of keys that define a cluster subset class

<a id="default-pool-advanced-options-http1-config"></a>

**Default Pool Advanced Options Http1 Config**

&#x2022; `header_transformation` - Optional Block<br>Header Transformation. Header Transformation options for HTTP/1.1 request/response headers<br>See [Header Transformation](#default-pool-advanced-options-http1-config-header-transformation) below.

<a id="default-pool-advanced-options-http1-config-header-transformation"></a>

**Default Pool Advanced Options Http1 Config Header Transformation**

&#x2022; `default_header_transformation` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `legacy_header_transformation` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `preserve_case_header_transformation` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `proper_case_header_transformation` - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="default-pool-advanced-options-http2-options"></a>

**Default Pool Advanced Options Http2 Options**

&#x2022; `enabled` - Optional Bool<br>HTTP2 Enabled. Enable/disable HTTP2 Protocol for upstream connections

<a id="default-pool-advanced-options-outlier-detection"></a>

**Default Pool Advanced Options Outlier Detection**

&#x2022; `base_ejection_time` - Optional Number  Defaults to `30000ms`  Specified in milliseconds<br>Base Ejection Time. The base time that a host is ejected for. The real time is equal to the base time multiplied by the number of times the host has been ejected. This causes hosts to get ejected for longer periods if they continue to fail

&#x2022; `consecutive_5xx` - Optional Number  Defaults to `5`<br>Consecutive 5xx Count. If an upstream endpoint returns some number of consecutive 5xx, it will be ejected. Note that in this case a 5xx means an actual 5xx respond code, or an event that would cause the HTTP router to return one on the upstream’s behalf(reset, connection failure, etc.) consecutive_5xx indicates the number of consecutive 5xx responses required before a consecutive 5xx ejection occurs

&#x2022; `consecutive_gateway_failure` - Optional Number  Defaults to `5`<br>Consecutive Gateway Failure. If an upstream endpoint returns some number of consecutive “gateway errors” (502, 503 or 504 status code), it will be ejected. Note that this includes events that would cause the HTTP router to return one of these status codes on the upstream’s behalf (reset, connection failure, etc.). consecutive_gateway_failure indicates the number of consecutive gateway failures before a consecutive gateway failure ejection occurs

&#x2022; `interval` - Optional Number  Defaults to `10000ms`  Specified in milliseconds<br>Interval. The time interval between ejection analysis sweeps. This can result in both new ejections as well as endpoints being returned to service

&#x2022; `max_ejection_percent` - Optional Number  Defaults to `10%`<br>Max Ejection Percentage. The maximum % of an upstream cluster that can be ejected due to outlier detection. but will eject at least one host regardless of the value

<a id="default-pool-healthcheck"></a>

**Default Pool Healthcheck**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="default-pool-origin-servers"></a>

**Default Pool Origin Servers**

&#x2022; `cbip_service` - Optional Block<br>Discovered Classic BIG-IP Service Name. Specify origin server with Classic BIG-IP Service (Virtual Server)<br>See [Cbip Service](#default-pool-origin-servers-cbip-service) below.

&#x2022; `consul_service` - Optional Block<br>Consul Service Name on given Sites. Specify origin server with Hashi Corp Consul service name and site information<br>See [Consul Service](#default-pool-origin-servers-consul-service) below.

&#x2022; `custom_endpoint_object` - Optional Block<br>Custom Endpoint Object for Origin Server. Specify origin server with a reference to endpoint object<br>See [Custom Endpoint Object](#default-pool-origin-servers-custom-endpoint-object) below.

&#x2022; `k8s_service` - Optional Block<br>K8s Service Name on given Sites. Specify origin server with K8s service name and site information<br>See [K8s Service](#default-pool-origin-servers-k8s-service) below.

&#x2022; `labels` - Optional Block<br>Origin Server Labels. Add Labels for this origin server, these labels can be used to form subset

&#x2022; `private_ip` - Optional Block<br>IP address on given Sites. Specify origin server with private or public IP address and site information<br>See [Private IP](#default-pool-origin-servers-private-ip) below.

&#x2022; `private_name` - Optional Block<br>DNS Name on given Sites. Specify origin server with private or public DNS name and site information<br>See [Private Name](#default-pool-origin-servers-private-name) below.

&#x2022; `public_ip` - Optional Block<br>Public IP. Specify origin server with public IP address<br>See [Public IP](#default-pool-origin-servers-public-ip) below.

&#x2022; `public_name` - Optional Block<br>Public DNS Name. Specify origin server with public DNS name<br>See [Public Name](#default-pool-origin-servers-public-name) below.

&#x2022; `vn_private_ip` - Optional Block<br>IP address Virtual Network. Specify origin server with IP on Virtual Network<br>See [Vn Private IP](#default-pool-origin-servers-vn-private-ip) below.

&#x2022; `vn_private_name` - Optional Block<br>DNS Name on Virtual Network. Specify origin server with DNS name on Virtual Network<br>See [Vn Private Name](#default-pool-origin-servers-vn-private-name) below.

<a id="default-pool-origin-servers-cbip-service"></a>

**Default Pool Origin Servers Cbip Service**

&#x2022; `service_name` - Optional String<br>Service Name. Name of the discovered Classic BIG-IP virtual server to be used as origin

<a id="default-pool-origin-servers-consul-service"></a>

**Default Pool Origin Servers Consul Service**

&#x2022; `inside_network` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `outside_network` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `service_name` - Optional String<br>Service Name. Consul service name of this origin server will be listed, including cluster-id. The format is servicename:cluster-id

&#x2022; `site_locator` - Optional Block<br>Site or Virtual Site. This message defines a reference to a site or virtual site object<br>See [Site Locator](#default-pool-origin-servers-consul-service-site-locator) below.

&#x2022; `snat_pool` - Optional Block<br>Snat Pool. Snat Pool configuration<br>See [Snat Pool](#default-pool-origin-servers-consul-service-snat-pool) below.

<a id="default-pool-origin-servers-consul-service-site-locator"></a>

**Default Pool Origin Servers Consul Service Site Locator**

&#x2022; `site` - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Site](#default-pool-origin-servers-consul-service-site-locator-site) below.

&#x2022; `virtual_site` - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Virtual Site](#default-pool-origin-servers-consul-service-site-locator-virtual-site) below.

<a id="default-pool-origin-servers-consul-service-site-locator-site"></a>

**Default Pool Origin Servers Consul Service Site Locator Site**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="default-pool-origin-servers-consul-service-site-locator-virtual-site"></a>

**Default Pool Origin Servers Consul Service Site Locator Virtual Site**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="default-pool-origin-servers-consul-service-snat-pool"></a>

**Default Pool Origin Servers Consul Service Snat Pool**

&#x2022; `no_snat_pool` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `snat_pool` - Optional Block<br>IPv4 Prefix List. x-example: '192.168.20.0/24' List of IPv4 prefixes that represent an endpoint<br>See [Snat Pool](#default-pool-origin-servers-consul-service-snat-pool-snat-pool) below.

<a id="default-pool-origin-servers-consul-service-snat-pool-snat-pool"></a>

**Default Pool Origin Servers Consul Service Snat Pool Snat Pool**

&#x2022; `prefixes` - Optional List<br>IPv4 Prefix List. List of IPv4 prefixes that represent an endpoint

<a id="default-pool-origin-servers-custom-endpoint-object"></a>

**Default Pool Origin Servers Custom Endpoint Object**

&#x2022; `endpoint` - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Endpoint](#default-pool-origin-servers-custom-endpoint-object-endpoint) below.

<a id="default-pool-origin-servers-custom-endpoint-object-endpoint"></a>

**Default Pool Origin Servers Custom Endpoint Object Endpoint**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="default-pool-origin-servers-k8s-service"></a>

**Default Pool Origin Servers K8s Service**

&#x2022; `inside_network` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `outside_network` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `protocol` - Optional String  Defaults to `PROTOCOL_TCP`<br>Possible values are `PROTOCOL_TCP`, `PROTOCOL_UDP`<br>Protocol Type. Type of protocol - PROTOCOL_TCP: TCP - PROTOCOL_UDP: UDP

&#x2022; `service_name` - Optional String<br>Service Name. K8s service name of the origin server will be listed, including the namespace and cluster-id. For vK8s services, you need to enter a string with the format servicename.namespace:cluster-id. If the servicename is 'frontend', namespace is 'speedtest' and cluster-id is 'prod', then you will enter 'frontend.speedtest:prod'. Both namespace and cluster-id are optional

&#x2022; `site_locator` - Optional Block<br>Site or Virtual Site. This message defines a reference to a site or virtual site object<br>See [Site Locator](#default-pool-origin-servers-k8s-service-site-locator) below.

&#x2022; `snat_pool` - Optional Block<br>Snat Pool. Snat Pool configuration<br>See [Snat Pool](#default-pool-origin-servers-k8s-service-snat-pool) below.

&#x2022; `vk8s_networks` - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="default-pool-origin-servers-k8s-service-site-locator"></a>

**Default Pool Origin Servers K8s Service Site Locator**

&#x2022; `site` - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Site](#default-pool-origin-servers-k8s-service-site-locator-site) below.

&#x2022; `virtual_site` - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Virtual Site](#default-pool-origin-servers-k8s-service-site-locator-virtual-site) below.

<a id="default-pool-origin-servers-k8s-service-site-locator-site"></a>

**Default Pool Origin Servers K8s Service Site Locator Site**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="default-pool-origin-servers-k8s-service-site-locator-virtual-site"></a>

**Default Pool Origin Servers K8s Service Site Locator Virtual Site**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="default-pool-origin-servers-k8s-service-snat-pool"></a>

**Default Pool Origin Servers K8s Service Snat Pool**

&#x2022; `no_snat_pool` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `snat_pool` - Optional Block<br>IPv4 Prefix List. x-example: '192.168.20.0/24' List of IPv4 prefixes that represent an endpoint<br>See [Snat Pool](#default-pool-origin-servers-k8s-service-snat-pool-snat-pool) below.

<a id="default-pool-origin-servers-k8s-service-snat-pool-snat-pool"></a>

**Default Pool Origin Servers K8s Service Snat Pool Snat Pool**

&#x2022; `prefixes` - Optional List<br>IPv4 Prefix List. List of IPv4 prefixes that represent an endpoint

<a id="default-pool-origin-servers-private-ip"></a>

**Default Pool Origin Servers Private IP**

&#x2022; `inside_network` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `ip` - Optional String<br>IP. Private IPv4 address

&#x2022; `outside_network` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `segment` - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Segment](#default-pool-origin-servers-private-ip-segment) below.

&#x2022; `site_locator` - Optional Block<br>Site or Virtual Site. This message defines a reference to a site or virtual site object<br>See [Site Locator](#default-pool-origin-servers-private-ip-site-locator) below.

&#x2022; `snat_pool` - Optional Block<br>Snat Pool. Snat Pool configuration<br>See [Snat Pool](#default-pool-origin-servers-private-ip-snat-pool) below.

<a id="default-pool-origin-servers-private-ip-segment"></a>

**Default Pool Origin Servers Private IP Segment**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="default-pool-origin-servers-private-ip-site-locator"></a>

**Default Pool Origin Servers Private IP Site Locator**

&#x2022; `site` - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Site](#default-pool-origin-servers-private-ip-site-locator-site) below.

&#x2022; `virtual_site` - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Virtual Site](#default-pool-origin-servers-private-ip-site-locator-virtual-site) below.

<a id="default-pool-origin-servers-private-ip-site-locator-site"></a>

**Default Pool Origin Servers Private IP Site Locator Site**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="default-pool-origin-servers-private-ip-site-locator-virtual-site"></a>

**Default Pool Origin Servers Private IP Site Locator Virtual Site**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="default-pool-origin-servers-private-ip-snat-pool"></a>

**Default Pool Origin Servers Private IP Snat Pool**

&#x2022; `no_snat_pool` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `snat_pool` - Optional Block<br>IPv4 Prefix List. x-example: '192.168.20.0/24' List of IPv4 prefixes that represent an endpoint<br>See [Snat Pool](#default-pool-origin-servers-private-ip-snat-pool-snat-pool) below.

<a id="default-pool-origin-servers-private-ip-snat-pool-snat-pool"></a>

**Default Pool Origin Servers Private IP Snat Pool Snat Pool**

&#x2022; `prefixes` - Optional List<br>IPv4 Prefix List. List of IPv4 prefixes that represent an endpoint

<a id="default-pool-origin-servers-private-name"></a>

**Default Pool Origin Servers Private Name**

&#x2022; `dns_name` - Optional String<br>DNS Name. DNS Name

&#x2022; `inside_network` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `outside_network` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `refresh_interval` - Optional Number<br>DNS Refresh Interval. Interval for DNS refresh in seconds. Max value is 7 days as per `HTTPS://datatracker.ietf.org/doc/HTML/rfc8767`

&#x2022; `segment` - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Segment](#default-pool-origin-servers-private-name-segment) below.

&#x2022; `site_locator` - Optional Block<br>Site or Virtual Site. This message defines a reference to a site or virtual site object<br>See [Site Locator](#default-pool-origin-servers-private-name-site-locator) below.

&#x2022; `snat_pool` - Optional Block<br>Snat Pool. Snat Pool configuration<br>See [Snat Pool](#default-pool-origin-servers-private-name-snat-pool) below.

<a id="default-pool-origin-servers-private-name-segment"></a>

**Default Pool Origin Servers Private Name Segment**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="default-pool-origin-servers-private-name-site-locator"></a>

**Default Pool Origin Servers Private Name Site Locator**

&#x2022; `site` - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Site](#default-pool-origin-servers-private-name-site-locator-site) below.

&#x2022; `virtual_site` - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Virtual Site](#default-pool-origin-servers-private-name-site-locator-virtual-site) below.

<a id="default-pool-origin-servers-private-name-site-locator-site"></a>

**Default Pool Origin Servers Private Name Site Locator Site**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="default-pool-origin-servers-private-name-site-locator-virtual-site"></a>

**Default Pool Origin Servers Private Name Site Locator Virtual Site**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="default-pool-origin-servers-private-name-snat-pool"></a>

**Default Pool Origin Servers Private Name Snat Pool**

&#x2022; `no_snat_pool` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `snat_pool` - Optional Block<br>IPv4 Prefix List. x-example: '192.168.20.0/24' List of IPv4 prefixes that represent an endpoint<br>See [Snat Pool](#default-pool-origin-servers-private-name-snat-pool-snat-pool) below.

<a id="default-pool-origin-servers-private-name-snat-pool-snat-pool"></a>

**Default Pool Origin Servers Private Name Snat Pool Snat Pool**

&#x2022; `prefixes` - Optional List<br>IPv4 Prefix List. List of IPv4 prefixes that represent an endpoint

<a id="default-pool-origin-servers-public-ip"></a>

**Default Pool Origin Servers Public IP**

&#x2022; `ip` - Optional String<br>Public IPv4. Public IPv4 address

<a id="default-pool-origin-servers-public-name"></a>

**Default Pool Origin Servers Public Name**

&#x2022; `dns_name` - Optional String<br>DNS Name. DNS Name

&#x2022; `refresh_interval` - Optional Number<br>DNS Refresh Interval. Interval for DNS refresh in seconds. Max value is 7 days as per `HTTPS://datatracker.ietf.org/doc/HTML/rfc8767`

<a id="default-pool-origin-servers-vn-private-ip"></a>

**Default Pool Origin Servers Vn Private IP**

&#x2022; `ip` - Optional String<br>IPv4. IPv4 address

&#x2022; `virtual_network` - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Virtual Network](#default-pool-origin-servers-vn-private-ip-virtual-network) below.

<a id="default-pool-origin-servers-vn-private-ip-virtual-network"></a>

**Default Pool Origin Servers Vn Private IP Virtual Network**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="default-pool-origin-servers-vn-private-name"></a>

**Default Pool Origin Servers Vn Private Name**

&#x2022; `dns_name` - Optional String<br>DNS Name. DNS Name

&#x2022; `private_network` - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Private Network](#default-pool-origin-servers-vn-private-name-private-network) below.

<a id="default-pool-origin-servers-vn-private-name-private-network"></a>

**Default Pool Origin Servers Vn Private Name Private Network**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="default-pool-upstream-conn-pool-reuse-type"></a>

**Default Pool Upstream Conn Pool Reuse Type**

&#x2022; `disable_conn_pool_reuse` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `enable_conn_pool_reuse` - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="default-pool-use-tls"></a>

**Default Pool Use TLS**

&#x2022; `default_session_key_caching` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `disable_session_key_caching` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `disable_sni` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `max_session_keys` - Optional Number<br>Max Session Keys Cached. x-example:'25' Number of session keys that are cached

&#x2022; `no_mtls` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `skip_server_verification` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `sni` - Optional String<br>SNI Value. SNI value to be used

&#x2022; `tls_config` - Optional Block<br>TLS Config. This defines various options to configure TLS configuration parameters<br>See [TLS Config](#default-pool-use-tls-tls-config) below.

&#x2022; `use_host_header_as_sni` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `use_mtls` - Optional Block<br>mTLS Certificate. mTLS Client Certificate<br>See [Use mTLS](#default-pool-use-tls-use-mtls) below.

&#x2022; `use_mtls_obj` - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Use mTLS Obj](#default-pool-use-tls-use-mtls-obj) below.

&#x2022; `use_server_verification` - Optional Block<br>TLS Validation Context for Origin Servers. Upstream TLS Validation Context<br>See [Use Server Verification](#default-pool-use-tls-use-server-verification) below.

&#x2022; `volterra_trusted_ca` - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="default-pool-use-tls-tls-config"></a>

**Default Pool Use TLS TLS Config**

&#x2022; `custom_security` - Optional Block<br>Custom Ciphers. This defines TLS protocol config including min/max versions and allowed ciphers<br>See [Custom Security](#default-pool-use-tls-tls-config-custom-security) below.

&#x2022; `default_security` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `low_security` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `medium_security` - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="default-pool-use-tls-tls-config-custom-security"></a>

**Default Pool Use TLS TLS Config Custom Security**

&#x2022; `cipher_suites` - Optional List<br>Cipher Suites. The TLS listener will only support the specified cipher list

&#x2022; `max_version` - Optional String  Defaults to `TLS_AUTO`<br>Possible values are `TLS_AUTO`, `TLSv1_0`, `TLSv1_1`, `TLSv1_2`, `TLSv1_3`<br>TLS Protocol. TlsProtocol is enumeration of supported TLS versions F5 Distributed Cloud will choose the optimal TLS version

&#x2022; `min_version` - Optional String  Defaults to `TLS_AUTO`<br>Possible values are `TLS_AUTO`, `TLSv1_0`, `TLSv1_1`, `TLSv1_2`, `TLSv1_3`<br>TLS Protocol. TlsProtocol is enumeration of supported TLS versions F5 Distributed Cloud will choose the optimal TLS version

<a id="default-pool-use-tls-use-mtls"></a>

**Default Pool Use TLS Use mTLS**

&#x2022; `tls_certificates` - Optional Block<br>mTLS Client Certificate. mTLS Client Certificate<br>See [TLS Certificates](#default-pool-use-tls-use-mtls-tls-certificates) below.

<a id="default-pool-use-tls-use-mtls-tls-certificates"></a>

**Default Pool Use TLS Use mTLS TLS Certificates**

&#x2022; `certificate_url` - Optional String<br>Certificate. TLS certificate. Certificate or certificate chain in PEM format including the PEM headers

&#x2022; `custom_hash_algorithms` - Optional Block<br>Hash Algorithms. Specifies the hash algorithms to be used<br>See [Custom Hash Algorithms](#default-pool-use-tls-use-mtls-tls-certificates-custom-hash-algorithms) below.

&#x2022; `description` - Optional String<br>Description. Description for the certificate

&#x2022; `disable_ocsp_stapling` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `private_key` - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Private Key](#default-pool-use-tls-use-mtls-tls-certificates-private-key) below.

&#x2022; `use_system_defaults` - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="default-pool-use-tls-use-mtls-tls-certificates-custom-hash-algorithms"></a>

**Default Pool Use TLS Use mTLS TLS Certificates Custom Hash Algorithms**

&#x2022; `hash_algorithms` - Optional List  Defaults to `INVALID_HASH_ALGORITHM`<br>Possible values are `INVALID_HASH_ALGORITHM`, `SHA256`, `SHA1`<br>Hash Algorithms. Ordered list of hash algorithms to be used

<a id="default-pool-use-tls-use-mtls-tls-certificates-private-key"></a>

**Default Pool Use TLS Use mTLS TLS Certificates Private Key**

&#x2022; `blindfold_secret_info` - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#default-pool-use-tls-use-mtls-tls-certificates-private-key-blindfold-secret-info) below.

&#x2022; `clear_secret_info` - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#default-pool-use-tls-use-mtls-tls-certificates-private-key-clear-secret-info) below.

<a id="default-pool-use-tls-use-mtls-tls-certificates-private-key-blindfold-secret-info"></a>

**Default Pool Use TLS Use mTLS TLS Certificates Private Key Blindfold Secret Info**

&#x2022; `decryption_provider` - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

&#x2022; `location` - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

&#x2022; `store_provider` - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

<a id="default-pool-use-tls-use-mtls-tls-certificates-private-key-clear-secret-info"></a>

**Default Pool Use TLS Use mTLS TLS Certificates Private Key Clear Secret Info**

&#x2022; `provider_ref` - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

&#x2022; `url` - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

<a id="default-pool-use-tls-use-mtls-obj"></a>

**Default Pool Use TLS Use mTLS Obj**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="default-pool-use-tls-use-server-verification"></a>

**Default Pool Use TLS Use Server Verification**

&#x2022; `trusted_ca` - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Trusted CA](#default-pool-use-tls-use-server-verification-trusted-ca) below.

&#x2022; `trusted_ca_url` - Optional String<br>Inline Root CA Certificate (legacy). Upload a Root CA Certificate specifically for this Origin Pool for verification of server's certificate

<a id="default-pool-use-tls-use-server-verification-trusted-ca"></a>

**Default Pool Use TLS Use Server Verification Trusted CA**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="default-pool-view-internal"></a>

**Default Pool View Internal**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="default-pool-list"></a>

**Default Pool List**

&#x2022; `pools` - Optional Block<br>Origin Pools. List of Origin Pools<br>See [Pools](#default-pool-list-pools) below.

<a id="default-pool-list-pools"></a>

**Default Pool List Pools**

&#x2022; `cluster` - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Cluster](#default-pool-list-pools-cluster) below.

&#x2022; `endpoint_subsets` - Optional Block<br>Origin Servers Subsets. Upstream origin pool may be configured to divide its origin servers into subsets based on metadata attached to the origin servers. Routes may then specify the metadata that a endpoint must match in order to be selected by the load balancer For origin servers which are discovered in K8S or Consul cluster, the label of the service is merged with endpoint's labels. In case of Consul, the label is derived from the 'Tag' field. For labels that are common between configured endpoint and discovered service, labels from discovered service takes precedence. List of key-value pairs that will be used as matching metadata. Only those origin servers of upstream origin pool which match this metadata will be selected for load balancing

&#x2022; `pool` - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Pool](#default-pool-list-pools-pool) below.

&#x2022; `priority` - Optional Number<br>Priority. Priority of this origin pool, valid only with multiple origin pools. Value of 0 will make the pool as lowest priority origin pool Priority of 1 means highest priority and is considered active. When active origin pool is not available, lower priority origin pools are made active as per the increasing priority

&#x2022; `weight` - Optional Number<br>Weight. Weight of this origin pool, valid only with multiple origin pool. Value of 0 will disable the pool

<a id="default-pool-list-pools-cluster"></a>

**Default Pool List Pools Cluster**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="default-pool-list-pools-pool"></a>

**Default Pool List Pools Pool**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="default-route-pools"></a>

**Default Route Pools**

&#x2022; `cluster` - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Cluster](#default-route-pools-cluster) below.

&#x2022; `endpoint_subsets` - Optional Block<br>Origin Servers Subsets. Upstream origin pool may be configured to divide its origin servers into subsets based on metadata attached to the origin servers. Routes may then specify the metadata that a endpoint must match in order to be selected by the load balancer For origin servers which are discovered in K8S or Consul cluster, the label of the service is merged with endpoint's labels. In case of Consul, the label is derived from the 'Tag' field. For labels that are common between configured endpoint and discovered service, labels from discovered service takes precedence. List of key-value pairs that will be used as matching metadata. Only those origin servers of upstream origin pool which match this metadata will be selected for load balancing

&#x2022; `pool` - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Pool](#default-route-pools-pool) below.

&#x2022; `priority` - Optional Number<br>Priority. Priority of this origin pool, valid only with multiple origin pools. Value of 0 will make the pool as lowest priority origin pool Priority of 1 means highest priority and is considered active. When active origin pool is not available, lower priority origin pools are made active as per the increasing priority

&#x2022; `weight` - Optional Number<br>Weight. Weight of this origin pool, valid only with multiple origin pool. Value of 0 will disable the pool

<a id="default-route-pools-cluster"></a>

**Default Route Pools Cluster**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="default-route-pools-pool"></a>

**Default Route Pools Pool**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="enable-api-discovery"></a>

**Enable API Discovery**

&#x2022; `api_crawler` - Optional Block<br>API Crawling. API Crawler message<br>See [API Crawler](#enable-api-discovery-api-crawler) below.

&#x2022; `api_discovery_from_code_scan` - Optional Block<br>Select Code Base and Repositories. x-required<br>See [API Discovery From Code Scan](#enable-api-discovery-api-discovery-from-code-scan) below.

&#x2022; `custom_api_auth_discovery` - Optional Block<br>API Discovery Advanced Settings. API Discovery Advanced settings<br>See [Custom API Auth Discovery](#enable-api-discovery-custom-api-auth-discovery) below.

&#x2022; `default_api_auth_discovery` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `disable_learn_from_redirect_traffic` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `discovered_api_settings` - Optional Block<br>Discovered API Settings. x-example: '2' Configure Discovered API Settings<br>See [Discovered API Settings](#enable-api-discovery-discovered-api-settings) below.

&#x2022; `enable_learn_from_redirect_traffic` - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="enable-api-discovery-api-crawler"></a>

**Enable API Discovery API Crawler**

&#x2022; `api_crawler_config` - Optional Block<br>Crawler Configure<br>See [API Crawler Config](#enable-api-discovery-api-crawler-api-crawler-config) below.

&#x2022; `disable_api_crawler` - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="enable-api-discovery-api-crawler-api-crawler-config"></a>

**Enable API Discovery API Crawler API Crawler Config**

&#x2022; `domains` - Optional Block<br>Domains to Crawl. Enter domains and their credentials to allow authenticated API crawling. You can only include domains you own that are associated with this Load Balancer<br>See [Domains](#enable-api-discovery-api-crawler-api-crawler-config-domains) below.

<a id="enable-api-discovery-api-crawler-api-crawler-config-domains"></a>

**Enable API Discovery API Crawler API Crawler Config Domains**

&#x2022; `domain` - Optional String<br>Domain. Select the domain to execute API Crawling with given credentials

&#x2022; `simple_login` - Optional Block<br>Simple Login<br>See [Simple Login](#enable-api-discovery-api-crawler-api-crawler-config-domains-simple-login) below.

<a id="enable-api-discovery-api-crawler-api-crawler-config-domains-simple-login"></a>

**Enable API Discovery API Crawler API Crawler Config Domains Simple Login**

&#x2022; `password` - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Password](#enable-api-discovery-api-crawler-api-crawler-config-domains-simple-login-password) below.

&#x2022; `user` - Optional String<br>User. Enter the username to assign credentials for the selected domain to crawl

<a id="enable-api-discovery-api-crawler-api-crawler-config-domains-simple-login-password"></a>

**Enable API Discovery API Crawler API Crawler Config Domains Simple Login Password**

&#x2022; `blindfold_secret_info` - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#enable-api-discovery-api-crawler-api-crawler-config-domains-simple-login-password-blindfold-secret-info) below.

&#x2022; `clear_secret_info` - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#enable-api-discovery-api-crawler-api-crawler-config-domains-simple-login-password-clear-secret-info) below.

<a id="enable-api-discovery-api-crawler-api-crawler-config-domains-simple-login-password-blindfold-secret-info"></a>

**Enable API Discovery API Crawler API Crawler Config Domains Simple Login Password Blindfold Secret Info**

&#x2022; `decryption_provider` - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

&#x2022; `location` - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

&#x2022; `store_provider` - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

<a id="enable-api-discovery-api-crawler-api-crawler-config-domains-simple-login-password-clear-secret-info"></a>

**Enable API Discovery API Crawler API Crawler Config Domains Simple Login Password Clear Secret Info**

&#x2022; `provider_ref` - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

&#x2022; `url` - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

<a id="enable-api-discovery-api-discovery-from-code-scan"></a>

**Enable API Discovery API Discovery From Code Scan**

&#x2022; `code_base_integrations` - Optional Block<br>Select Code Base Integrations<br>See [Code Base Integrations](#enable-api-discovery-api-discovery-from-code-scan-code-base-integrations) below.

<a id="enable-api-discovery-api-discovery-from-code-scan-code-base-integrations"></a>

**Enable API Discovery API Discovery From Code Scan Code Base Integrations**

&#x2022; `all_repos` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `code_base_integration` - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Code Base Integration](#enable-api-discovery-api-discovery-from-code-scan-code-base-integrations-code-base-integration) below.

&#x2022; `selected_repos` - Optional Block<br>API Code Repositories. Select which API repositories represent the LB applications<br>See [Selected Repos](#enable-api-discovery-api-discovery-from-code-scan-code-base-integrations-selected-repos) below.

<a id="enable-api-discovery-api-discovery-from-code-scan-code-base-integrations-code-base-integration"></a>

**Enable API Discovery API Discovery From Code Scan Code Base Integrations Code Base Integration**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="enable-api-discovery-api-discovery-from-code-scan-code-base-integrations-selected-repos"></a>

**Enable API Discovery API Discovery From Code Scan Code Base Integrations Selected Repos**

&#x2022; `api_code_repo` - Optional List<br>API Code Repository. Code repository which contain API endpoints

<a id="enable-api-discovery-custom-api-auth-discovery"></a>

**Enable API Discovery Custom API Auth Discovery**

&#x2022; `api_discovery_ref` - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [API Discovery Ref](#enable-api-discovery-custom-api-auth-discovery-api-discovery-ref) below.

<a id="enable-api-discovery-custom-api-auth-discovery-api-discovery-ref"></a>

**Enable API Discovery Custom API Auth Discovery API Discovery Ref**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="enable-api-discovery-discovered-api-settings"></a>

**Enable API Discovery Discovered API Settings**

&#x2022; `purge_duration_for_inactive_discovered_apis` - Optional Number<br>Purge Duration for Inactive Discovered APIs from Traffic. Inactive discovered API will be deleted after configured duration

<a id="enable-challenge"></a>

**Enable Challenge**

&#x2022; `captcha_challenge_parameters` - Optional Block<br>Captcha Challenge Parameters. Enables loadbalancer to perform captcha challenge Captcha challenge will be based on Google Recaptcha. With this feature enabled, only clients that pass the captcha challenge will be allowed to complete the HTTP request. When loadbalancer is configured to do Captcha Challenge, it will redirect the browser to an HTML page on every new HTTP request. This HTML page will have captcha challenge embedded in it. Client will be allowed to make the request only if the captcha challenge is successful. Loadbalancer will tag response header with a cookie to avoid Captcha challenge for subsequent requests. CAPTCHA is mainly used as a security check to ensure only human users can pass through. Generally, computers or bots are not capable of solving a captcha. You can enable either Javascript challenge or Captcha challenge on a virtual host<br>See [Captcha Challenge Parameters](#enable-challenge-captcha-challenge-parameters) below.

&#x2022; `default_captcha_challenge_parameters` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `default_js_challenge_parameters` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `default_mitigation_settings` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `js_challenge_parameters` - Optional Block<br>Javascript Challenge Parameters. Enables loadbalancer to perform client browser compatibility test by redirecting to a page with Javascript. With this feature enabled, only clients that are capable of executing Javascript(mostly browsers) will be allowed to complete the HTTP request. When loadbalancer is configured to do Javascript Challenge, it will redirect the browser to an HTML page on every new HTTP request. This HTML page will have Javascript embedded in it. Loadbalancer chooses a set of random numbers for every new client and sends these numbers along with an encrypted answer with the request such that it embed these numbers as input in the Javascript. Javascript will run on the requestor browser and perform a complex Math operation. Script will submit the answer to loadbalancer. Loadbalancer will validate the answer by comparing the calculated answer with the decrypted answer (which was encrypted when it was sent back as reply) and allow the request to the upstream server only if the answer is correct. Loadbalancer will tag response header with a cookie to avoid Javascript challenge for subsequent requests. Javascript challenge serves following purposes * Validate that the request is coming via a browser that is capable for running Javascript * Force the browser to run a complex operation, f(X), that requires it to spend a large number of CPU cycles. This is to slow down a potential DOS attacker by making it difficult to launch a large request flood without having to spend even larger CPU cost at their end. You can enable either Javascript challenge or Captcha challenge on a virtual host<br>See [Js Challenge Parameters](#enable-challenge-js-challenge-parameters) below.

&#x2022; `malicious_user_mitigation` - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Malicious User Mitigation](#enable-challenge-malicious-user-mitigation) below.

<a id="enable-challenge-captcha-challenge-parameters"></a>

**Enable Challenge Captcha Challenge Parameters**

&#x2022; `cookie_expiry` - Optional Number<br>Cookie Expiration Period. Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new challenge

&#x2022; `custom_page` - Optional String<br>Custom message for Captcha Challenge. Custom message is of type uri_ref. Currently supported URL schemes is string:///. For string:/// scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. 'Please Wait.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '<p> Please Wait </p>'. Base64 encoded string for this HTML is 'PHA+IFBsZWFzZSBXYWl0IDwvcD4='

<a id="enable-challenge-js-challenge-parameters"></a>

**Enable Challenge Js Challenge Parameters**

&#x2022; `cookie_expiry` - Optional Number<br>Cookie Expiration Period. Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new challenge

&#x2022; `custom_page` - Optional String<br>Custom Message for Javascript Challenge. Custom message is of type uri_ref. Currently supported URL schemes is string:///. For string:/// scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. 'Please Wait.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '<p> Please Wait </p>'. Base64 encoded string for this HTML is 'PHA+IFBsZWFzZSBXYWl0IDwvcD4='

&#x2022; `js_script_delay` - Optional Number<br>Javascript Delay. Delay introduced by Javascript, in milliseconds

<a id="enable-challenge-malicious-user-mitigation"></a>

**Enable Challenge Malicious User Mitigation**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="enable-ip-reputation"></a>

**Enable IP Reputation**

&#x2022; `ip_threat_categories` - Optional List  Defaults to `SPAM_SOURCES`<br>Possible values are `SPAM_SOURCES`, `WINDOWS_EXPLOITS`, `WEB_ATTACKS`, `BOTNETS`, `SCANNERS`, `REPUTATION`, `PHISHING`, `PROXY`, `MOBILE_THREATS`, `TOR_PROXY`, `DENIAL_OF_SERVICE`, `NETWORK`<br>List of IP Threat Categories to choose. If the source IP matches on atleast one of the enabled IP threat categories, the request will be denied

<a id="enable-trust-client-ip-headers"></a>

**Enable Trust Client IP Headers**

&#x2022; `client_ip_headers` - Optional List<br>Client IP Headers. Define the list of one or more Client IP Headers. Headers will be used in order from top to bottom, meaning if the first header is not present in the request, the system will proceed to check for the second header, and so on, until one of the listed headers is found. If none of the defined headers exist, or the value is not an IP address, then the system will use the source IP of the packet. If multiple defined headers with different names are present in the request, the value of the first header name in the configuration will be used. If multiple defined headers with the same name are present in the request, values of all those headers will be combined. The system will read the right-most IP address from header, if there are multiple IP addresses in the header value. For X-Forwarded-For header, the system will read the IP address(rightmost - 1), as the client IP

<a id="graphql-rules"></a>

**GraphQL Rules**

&#x2022; `any_domain` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `exact_path` - Optional String  Defaults to `/GraphQL`<br>Path. Specifies the exact path to GraphQL endpoint

&#x2022; `exact_value` - Optional String<br>Exact Value. Exact domain name

&#x2022; `graphql_settings` - Optional Block<br>GraphQL Settings. GraphQL configuration<br>See [GraphQL Settings](#graphql-rules-graphql-settings) below.

&#x2022; `metadata` - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#graphql-rules-metadata) below.

&#x2022; `method_get` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `method_post` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `suffix_value` - Optional String<br>Suffix Value. Suffix of domain name e.g 'xyz.com' will match '*.xyz.com' and 'xyz.com'

<a id="graphql-rules-graphql-settings"></a>

**GraphQL Rules GraphQL Settings**

&#x2022; `disable_introspection` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `enable_introspection` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `max_batched_queries` - Optional Number<br>Maximum Batched Queries. Specify maximum number of queries in a single batched request

&#x2022; `max_depth` - Optional Number<br>Maximum Structure Depth. Specify maximum depth for the GraphQL query

&#x2022; `max_total_length` - Optional Number<br>Maximum Total Length. Specify maximum length in bytes for the GraphQL query

<a id="graphql-rules-metadata"></a>

**GraphQL Rules Metadata**

&#x2022; `description` - Optional String<br>Description. Human readable description

&#x2022; `name` - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

<a id="http"></a>

**HTTP**

&#x2022; `dns_volterra_managed` - Optional Bool<br>Automatically Manage DNS Records. DNS records for domains will be managed automatically by F5 Distributed Cloud. As a prerequisite, the domain must be delegated to F5 Distributed Cloud using Delegated domain feature or a DNS CNAME record should be created in your DNS provider's portal

&#x2022; `port` - Optional Number<br>HTTP Listen Port. HTTP port to Listen

&#x2022; `port_ranges` - Optional String<br>Port Ranges. A string containing a comma separated list of port ranges. Each port range consists of a single port or two ports separated by '-'

<a id="https"></a>

**HTTPS**

&#x2022; `add_hsts` - Optional Bool<br>Add HSTS Header. Add HTTP Strict-Transport-Security response header

&#x2022; `append_server_name` - Optional String<br>Append header value. Define the header value for the header name “server”. If header value is already present, it is not overwritten and passed as-is

&#x2022; `coalescing_options` - Optional Block<br>TLS Coalescing Options. TLS connection coalescing configuration (not compatible with mTLS)<br>See [Coalescing Options](#https-coalescing-options) below.

&#x2022; `connection_idle_timeout` - Optional Number  Defaults to `2`  Specified in milliseconds<br>Connection Idle Timeout. The idle timeout for downstream connections. The idle timeout is defined as the period in which there are no active requests. When the idle timeout is reached the connection will be closed. Note that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive.  The minutes

&#x2022; `default_header` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `default_loadbalancer` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `disable_path_normalize` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `enable_path_normalize` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `http_protocol_options` - Optional Block<br>HTTP Protocol Configuration Options. HTTP protocol configuration options for downstream connections<br>See [HTTP Protocol Options](#https-http-protocol-options) below.

&#x2022; `http_redirect` - Optional Bool<br>HTTP Redirect to HTTPS. Redirect HTTP traffic to HTTPS

&#x2022; `non_default_loadbalancer` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `pass_through` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `port` - Optional Number<br>HTTPS Port. HTTPS port to Listen

&#x2022; `port_ranges` - Optional String<br>Port Ranges. A string containing a comma separated list of port ranges. Each port range consists of a single port or two ports separated by '-'

&#x2022; `server_name` - Optional String<br>Modify header value. Define the header value for the header name “server”. This will overwrite existing values, if any, for the server header

&#x2022; `tls_cert_params` - Optional Block<br>TLS Parameters. Select TLS Parameters and Certificates<br>See [TLS Cert Params](#https-tls-cert-params) below.

&#x2022; `tls_parameters` - Optional Block<br>Inline TLS Parameters. Inline TLS parameters<br>See [TLS Parameters](#https-tls-parameters) below.

<a id="https-coalescing-options"></a>

**HTTPS Coalescing Options**

&#x2022; `default_coalescing` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `strict_coalescing` - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="https-http-protocol-options"></a>

**HTTPS HTTP Protocol Options**

&#x2022; `http_protocol_enable_v1_only` - Optional Block<br>HTTP/1.1 Protocol Options. HTTP/1.1 Protocol options for downstream connections<br>See [HTTP Protocol Enable V1 Only](#https-http-protocol-options-http-protocol-enable-v1-only) below.

&#x2022; `http_protocol_enable_v1_v2` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `http_protocol_enable_v2_only` - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="https-http-protocol-options-http-protocol-enable-v1-only"></a>

**HTTPS HTTP Protocol Options HTTP Protocol Enable V1 Only**

&#x2022; `header_transformation` - Optional Block<br>Header Transformation. Header Transformation options for HTTP/1.1 request/response headers<br>See [Header Transformation](#https-http-protocol-options-http-protocol-enable-v1-only-header-transformation) below.

<a id="https-http-protocol-options-http-protocol-enable-v1-only-header-transformation"></a>

**HTTPS HTTP Protocol Options HTTP Protocol Enable V1 Only Header Transformation**

&#x2022; `default_header_transformation` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `legacy_header_transformation` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `preserve_case_header_transformation` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `proper_case_header_transformation` - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="https-tls-cert-params"></a>

**HTTPS TLS Cert Params**

&#x2022; `certificates` - Optional Block<br>Certificates. Select one or more certificates with any domain names<br>See [Certificates](#https-tls-cert-params-certificates) below.

&#x2022; `no_mtls` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `tls_config` - Optional Block<br>TLS Config. This defines various options to configure TLS configuration parameters<br>See [TLS Config](#https-tls-cert-params-tls-config) below.

&#x2022; `use_mtls` - Optional Block<br>Clients TLS validation context. Validation context for downstream client TLS connections<br>See [Use mTLS](#https-tls-cert-params-use-mtls) below.

<a id="https-tls-cert-params-certificates"></a>

**HTTPS TLS Cert Params Certificates**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="https-tls-cert-params-tls-config"></a>

**HTTPS TLS Cert Params TLS Config**

&#x2022; `custom_security` - Optional Block<br>Custom Ciphers. This defines TLS protocol config including min/max versions and allowed ciphers<br>See [Custom Security](#https-tls-cert-params-tls-config-custom-security) below.

&#x2022; `default_security` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `low_security` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `medium_security` - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="https-tls-cert-params-tls-config-custom-security"></a>

**HTTPS TLS Cert Params TLS Config Custom Security**

&#x2022; `cipher_suites` - Optional List<br>Cipher Suites. The TLS listener will only support the specified cipher list

&#x2022; `max_version` - Optional String  Defaults to `TLS_AUTO`<br>Possible values are `TLS_AUTO`, `TLSv1_0`, `TLSv1_1`, `TLSv1_2`, `TLSv1_3`<br>TLS Protocol. TlsProtocol is enumeration of supported TLS versions F5 Distributed Cloud will choose the optimal TLS version

&#x2022; `min_version` - Optional String  Defaults to `TLS_AUTO`<br>Possible values are `TLS_AUTO`, `TLSv1_0`, `TLSv1_1`, `TLSv1_2`, `TLSv1_3`<br>TLS Protocol. TlsProtocol is enumeration of supported TLS versions F5 Distributed Cloud will choose the optimal TLS version

<a id="https-tls-cert-params-use-mtls"></a>

**HTTPS TLS Cert Params Use mTLS**

&#x2022; `client_certificate_optional` - Optional Bool<br>Client Certificate Optional. Client certificate is optional. If the client has provided a certificate, the load balancer will verify it. If certification verification fails, the connection will be terminated. If the client does not provide a certificate, the connection will be accepted

&#x2022; `crl` - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [CRL](#https-tls-cert-params-use-mtls-crl) below.

&#x2022; `no_crl` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `trusted_ca` - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Trusted CA](#https-tls-cert-params-use-mtls-trusted-ca) below.

&#x2022; `trusted_ca_url` - Optional String<br>Inline Root CA Certificate (legacy). Upload a Root CA Certificate specifically for this Load Balancer

&#x2022; `xfcc_disabled` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `xfcc_options` - Optional Block<br>XFCC Header Elements. X-Forwarded-Client-Cert header elements to be added to requests<br>See [Xfcc Options](#https-tls-cert-params-use-mtls-xfcc-options) below.

<a id="https-tls-cert-params-use-mtls-crl"></a>

**HTTPS TLS Cert Params Use mTLS CRL**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="https-tls-cert-params-use-mtls-trusted-ca"></a>

**HTTPS TLS Cert Params Use mTLS Trusted CA**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="https-tls-cert-params-use-mtls-xfcc-options"></a>

**HTTPS TLS Cert Params Use mTLS Xfcc Options**

&#x2022; `xfcc_header_elements` - Optional List  Defaults to `XFCC_NONE`<br>Possible values are `XFCC_NONE`, `XFCC_CERT`, `XFCC_CHAIN`, `XFCC_SUBJECT`, `XFCC_URI`, `XFCC_DNS`<br>XFCC Header Elements. X-Forwarded-Client-Cert header elements to be added to requests

<a id="https-tls-parameters"></a>

**HTTPS TLS Parameters**

&#x2022; `no_mtls` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `tls_certificates` - Optional Block<br>TLS Certificates. Users can add one or more certificates that share the same set of domains. for example, domain.com and *.domain.com - but use different signature algorithms<br>See [TLS Certificates](#https-tls-parameters-tls-certificates) below.

&#x2022; `tls_config` - Optional Block<br>TLS Config. This defines various options to configure TLS configuration parameters<br>See [TLS Config](#https-tls-parameters-tls-config) below.

&#x2022; `use_mtls` - Optional Block<br>Clients TLS validation context. Validation context for downstream client TLS connections<br>See [Use mTLS](#https-tls-parameters-use-mtls) below.

<a id="https-tls-parameters-tls-certificates"></a>

**HTTPS TLS Parameters TLS Certificates**

&#x2022; `certificate_url` - Optional String<br>Certificate. TLS certificate. Certificate or certificate chain in PEM format including the PEM headers

&#x2022; `custom_hash_algorithms` - Optional Block<br>Hash Algorithms. Specifies the hash algorithms to be used<br>See [Custom Hash Algorithms](#https-tls-parameters-tls-certificates-custom-hash-algorithms) below.

&#x2022; `description` - Optional String<br>Description. Description for the certificate

&#x2022; `disable_ocsp_stapling` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `private_key` - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Private Key](#https-tls-parameters-tls-certificates-private-key) below.

&#x2022; `use_system_defaults` - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="https-tls-parameters-tls-certificates-custom-hash-algorithms"></a>

**HTTPS TLS Parameters TLS Certificates Custom Hash Algorithms**

&#x2022; `hash_algorithms` - Optional List  Defaults to `INVALID_HASH_ALGORITHM`<br>Possible values are `INVALID_HASH_ALGORITHM`, `SHA256`, `SHA1`<br>Hash Algorithms. Ordered list of hash algorithms to be used

<a id="https-tls-parameters-tls-certificates-private-key"></a>

**HTTPS TLS Parameters TLS Certificates Private Key**

&#x2022; `blindfold_secret_info` - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#https-tls-parameters-tls-certificates-private-key-blindfold-secret-info) below.

&#x2022; `clear_secret_info` - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#https-tls-parameters-tls-certificates-private-key-clear-secret-info) below.

<a id="https-tls-parameters-tls-certificates-private-key-blindfold-secret-info"></a>

**HTTPS TLS Parameters TLS Certificates Private Key Blindfold Secret Info**

&#x2022; `decryption_provider` - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

&#x2022; `location` - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

&#x2022; `store_provider` - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

<a id="https-tls-parameters-tls-certificates-private-key-clear-secret-info"></a>

**HTTPS TLS Parameters TLS Certificates Private Key Clear Secret Info**

&#x2022; `provider_ref` - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

&#x2022; `url` - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

<a id="https-tls-parameters-tls-config"></a>

**HTTPS TLS Parameters TLS Config**

&#x2022; `custom_security` - Optional Block<br>Custom Ciphers. This defines TLS protocol config including min/max versions and allowed ciphers<br>See [Custom Security](#https-tls-parameters-tls-config-custom-security) below.

&#x2022; `default_security` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `low_security` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `medium_security` - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="https-tls-parameters-tls-config-custom-security"></a>

**HTTPS TLS Parameters TLS Config Custom Security**

&#x2022; `cipher_suites` - Optional List<br>Cipher Suites. The TLS listener will only support the specified cipher list

&#x2022; `max_version` - Optional String  Defaults to `TLS_AUTO`<br>Possible values are `TLS_AUTO`, `TLSv1_0`, `TLSv1_1`, `TLSv1_2`, `TLSv1_3`<br>TLS Protocol. TlsProtocol is enumeration of supported TLS versions F5 Distributed Cloud will choose the optimal TLS version

&#x2022; `min_version` - Optional String  Defaults to `TLS_AUTO`<br>Possible values are `TLS_AUTO`, `TLSv1_0`, `TLSv1_1`, `TLSv1_2`, `TLSv1_3`<br>TLS Protocol. TlsProtocol is enumeration of supported TLS versions F5 Distributed Cloud will choose the optimal TLS version

<a id="https-tls-parameters-use-mtls"></a>

**HTTPS TLS Parameters Use mTLS**

&#x2022; `client_certificate_optional` - Optional Bool<br>Client Certificate Optional. Client certificate is optional. If the client has provided a certificate, the load balancer will verify it. If certification verification fails, the connection will be terminated. If the client does not provide a certificate, the connection will be accepted

&#x2022; `crl` - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [CRL](#https-tls-parameters-use-mtls-crl) below.

&#x2022; `no_crl` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `trusted_ca` - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Trusted CA](#https-tls-parameters-use-mtls-trusted-ca) below.

&#x2022; `trusted_ca_url` - Optional String<br>Inline Root CA Certificate (legacy). Upload a Root CA Certificate specifically for this Load Balancer

&#x2022; `xfcc_disabled` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `xfcc_options` - Optional Block<br>XFCC Header Elements. X-Forwarded-Client-Cert header elements to be added to requests<br>See [Xfcc Options](#https-tls-parameters-use-mtls-xfcc-options) below.

<a id="https-tls-parameters-use-mtls-crl"></a>

**HTTPS TLS Parameters Use mTLS CRL**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="https-tls-parameters-use-mtls-trusted-ca"></a>

**HTTPS TLS Parameters Use mTLS Trusted CA**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="https-tls-parameters-use-mtls-xfcc-options"></a>

**HTTPS TLS Parameters Use mTLS Xfcc Options**

&#x2022; `xfcc_header_elements` - Optional List  Defaults to `XFCC_NONE`<br>Possible values are `XFCC_NONE`, `XFCC_CERT`, `XFCC_CHAIN`, `XFCC_SUBJECT`, `XFCC_URI`, `XFCC_DNS`<br>XFCC Header Elements. X-Forwarded-Client-Cert header elements to be added to requests

<a id="https-auto-cert"></a>

**HTTPS Auto Cert**

&#x2022; `add_hsts` - Optional Bool<br>Add HSTS Header. Add HTTP Strict-Transport-Security response header

&#x2022; `append_server_name` - Optional String<br>Append header value. Define the header value for the header name “server”. If header value is already present, it is not overwritten and passed as-is

&#x2022; `coalescing_options` - Optional Block<br>TLS Coalescing Options. TLS connection coalescing configuration (not compatible with mTLS)<br>See [Coalescing Options](#https-auto-cert-coalescing-options) below.

&#x2022; `connection_idle_timeout` - Optional Number  Defaults to `2`  Specified in milliseconds<br>Connection Idle Timeout. The idle timeout for downstream connections. The idle timeout is defined as the period in which there are no active requests. When the idle timeout is reached the connection will be closed. Note that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive.  The minutes

&#x2022; `default_header` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `default_loadbalancer` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `disable_path_normalize` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `enable_path_normalize` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `http_protocol_options` - Optional Block<br>HTTP Protocol Configuration Options. HTTP protocol configuration options for downstream connections<br>See [HTTP Protocol Options](#https-auto-cert-http-protocol-options) below.

&#x2022; `http_redirect` - Optional Bool<br>HTTP Redirect to HTTPS. Redirect HTTP traffic to HTTPS

&#x2022; `no_mtls` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `non_default_loadbalancer` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `pass_through` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `port` - Optional Number<br>HTTPS Listen Port. HTTPS port to Listen

&#x2022; `port_ranges` - Optional String<br>Port Ranges. A string containing a comma separated list of port ranges. Each port range consists of a single port or two ports separated by '-'

&#x2022; `server_name` - Optional String<br>Modify header value. Define the header value for the header name “server”. This will overwrite existing values, if any, for the server header

&#x2022; `tls_config` - Optional Block<br>TLS Config. This defines various options to configure TLS configuration parameters<br>See [TLS Config](#https-auto-cert-tls-config) below.

&#x2022; `use_mtls` - Optional Block<br>Clients TLS validation context. Validation context for downstream client TLS connections<br>See [Use mTLS](#https-auto-cert-use-mtls) below.

<a id="https-auto-cert-coalescing-options"></a>

**HTTPS Auto Cert Coalescing Options**

&#x2022; `default_coalescing` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `strict_coalescing` - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="https-auto-cert-http-protocol-options"></a>

**HTTPS Auto Cert HTTP Protocol Options**

&#x2022; `http_protocol_enable_v1_only` - Optional Block<br>HTTP/1.1 Protocol Options. HTTP/1.1 Protocol options for downstream connections<br>See [HTTP Protocol Enable V1 Only](#https-auto-cert-http-protocol-options-http-protocol-enable-v1-only) below.

&#x2022; `http_protocol_enable_v1_v2` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `http_protocol_enable_v2_only` - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="https-auto-cert-http-protocol-options-http-protocol-enable-v1-only"></a>

**HTTPS Auto Cert HTTP Protocol Options HTTP Protocol Enable V1 Only**

&#x2022; `header_transformation` - Optional Block<br>Header Transformation. Header Transformation options for HTTP/1.1 request/response headers<br>See [Header Transformation](#https-auto-cert-http-protocol-options-http-protocol-enable-v1-only-header-transformation) below.

<a id="https-auto-cert-http-protocol-options-http-protocol-enable-v1-only-header-transformation"></a>

**HTTPS Auto Cert HTTP Protocol Options HTTP Protocol Enable V1 Only Header Transformation**

&#x2022; `default_header_transformation` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `legacy_header_transformation` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `preserve_case_header_transformation` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `proper_case_header_transformation` - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="https-auto-cert-tls-config"></a>

**HTTPS Auto Cert TLS Config**

&#x2022; `custom_security` - Optional Block<br>Custom Ciphers. This defines TLS protocol config including min/max versions and allowed ciphers<br>See [Custom Security](#https-auto-cert-tls-config-custom-security) below.

&#x2022; `default_security` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `low_security` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `medium_security` - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="https-auto-cert-tls-config-custom-security"></a>

**HTTPS Auto Cert TLS Config Custom Security**

&#x2022; `cipher_suites` - Optional List<br>Cipher Suites. The TLS listener will only support the specified cipher list

&#x2022; `max_version` - Optional String  Defaults to `TLS_AUTO`<br>Possible values are `TLS_AUTO`, `TLSv1_0`, `TLSv1_1`, `TLSv1_2`, `TLSv1_3`<br>TLS Protocol. TlsProtocol is enumeration of supported TLS versions F5 Distributed Cloud will choose the optimal TLS version

&#x2022; `min_version` - Optional String  Defaults to `TLS_AUTO`<br>Possible values are `TLS_AUTO`, `TLSv1_0`, `TLSv1_1`, `TLSv1_2`, `TLSv1_3`<br>TLS Protocol. TlsProtocol is enumeration of supported TLS versions F5 Distributed Cloud will choose the optimal TLS version

<a id="https-auto-cert-use-mtls"></a>

**HTTPS Auto Cert Use mTLS**

&#x2022; `client_certificate_optional` - Optional Bool<br>Client Certificate Optional. Client certificate is optional. If the client has provided a certificate, the load balancer will verify it. If certification verification fails, the connection will be terminated. If the client does not provide a certificate, the connection will be accepted

&#x2022; `crl` - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [CRL](#https-auto-cert-use-mtls-crl) below.

&#x2022; `no_crl` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `trusted_ca` - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Trusted CA](#https-auto-cert-use-mtls-trusted-ca) below.

&#x2022; `trusted_ca_url` - Optional String<br>Inline Root CA Certificate (legacy). Upload a Root CA Certificate specifically for this Load Balancer

&#x2022; `xfcc_disabled` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `xfcc_options` - Optional Block<br>XFCC Header Elements. X-Forwarded-Client-Cert header elements to be added to requests<br>See [Xfcc Options](#https-auto-cert-use-mtls-xfcc-options) below.

<a id="https-auto-cert-use-mtls-crl"></a>

**HTTPS Auto Cert Use mTLS CRL**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="https-auto-cert-use-mtls-trusted-ca"></a>

**HTTPS Auto Cert Use mTLS Trusted CA**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="https-auto-cert-use-mtls-xfcc-options"></a>

**HTTPS Auto Cert Use mTLS Xfcc Options**

&#x2022; `xfcc_header_elements` - Optional List  Defaults to `XFCC_NONE`<br>Possible values are `XFCC_NONE`, `XFCC_CERT`, `XFCC_CHAIN`, `XFCC_SUBJECT`, `XFCC_URI`, `XFCC_DNS`<br>XFCC Header Elements. X-Forwarded-Client-Cert header elements to be added to requests

<a id="js-challenge"></a>

**Js Challenge**

&#x2022; `cookie_expiry` - Optional Number<br>Cookie Expiration Period. Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new challenge

&#x2022; `custom_page` - Optional String<br>Custom Message for Javascript Challenge. Custom message is of type uri_ref. Currently supported URL schemes is string:///. For string:/// scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. 'Please Wait.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '<p> Please Wait </p>'. Base64 encoded string for this HTML is 'PHA+IFBsZWFzZSBXYWl0IDwvcD4='

&#x2022; `js_script_delay` - Optional Number<br>Javascript Delay. Delay introduced by Javascript, in milliseconds

<a id="jwt-validation"></a>

**JWT Validation**

&#x2022; `action` - Optional Block<br>Action<br>See [Action](#jwt-validation-action) below.

&#x2022; `jwks_config` - Optional Block<br>JSON Web Key Set (JWKS). The JSON Web Key Set (JWKS) is a set of keys used to verify JSON Web Token (JWT) issued by the Authorization Server. See RFC 7517 for more details<br>See [Jwks Config](#jwt-validation-jwks-config) below.

&#x2022; `mandatory_claims` - Optional Block<br>Mandatory Claims. Configurable Validation of mandatory Claims<br>See [Mandatory Claims](#jwt-validation-mandatory-claims) below.

&#x2022; `reserved_claims` - Optional Block<br>Reserved claims configuration. Configurable Validation of reserved Claims<br>See [Reserved Claims](#jwt-validation-reserved-claims) below.

&#x2022; `target` - Optional Block<br>Target. Define endpoints for which JWT token validation will be performed<br>See [Target](#jwt-validation-target) below.

&#x2022; `token_location` - Optional Block<br>Token Location. Location of JWT in HTTP request<br>See [Token Location](#jwt-validation-token-location) below.

<a id="jwt-validation-action"></a>

**JWT Validation Action**

&#x2022; `block` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `report` - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="jwt-validation-jwks-config"></a>

**JWT Validation Jwks Config**

&#x2022; `cleartext` - Optional String<br>JSON Web Key Set (JWKS). The JSON Web Key Set (JWKS) is a set of keys used to verify JSON Web Token (JWT) issued by the Authorization Server. See RFC 7517 for more details

<a id="jwt-validation-mandatory-claims"></a>

**JWT Validation Mandatory Claims**

&#x2022; `claim_names` - Optional List<br>Claim Names

<a id="jwt-validation-reserved-claims"></a>

**JWT Validation Reserved Claims**

&#x2022; `audience` - Optional Block<br>Audiences<br>See [Audience](#jwt-validation-reserved-claims-audience) below.

&#x2022; `audience_disable` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `issuer` - Optional String<br>Exact Match

&#x2022; `issuer_disable` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `validate_period_disable` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `validate_period_enable` - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="jwt-validation-reserved-claims-audience"></a>

**JWT Validation Reserved Claims Audience**

&#x2022; `audiences` - Optional List<br>Values

<a id="jwt-validation-target"></a>

**JWT Validation Target**

&#x2022; `all_endpoint` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `api_groups` - Optional Block<br>API Groups<br>See [API Groups](#jwt-validation-target-api-groups) below.

&#x2022; `base_paths` - Optional Block<br>Base Paths<br>See [Base Paths](#jwt-validation-target-base-paths) below.

<a id="jwt-validation-target-api-groups"></a>

**JWT Validation Target API Groups**

&#x2022; `api_groups` - Optional List<br>API Groups

<a id="jwt-validation-target-base-paths"></a>

**JWT Validation Target Base Paths**

&#x2022; `base_paths` - Optional List<br>Prefix Values

<a id="jwt-validation-token-location"></a>

**JWT Validation Token Location**

&#x2022; `bearer_token` - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="l7-ddos-action-js-challenge"></a>

**L7 DDOS Action Js Challenge**

&#x2022; `cookie_expiry` - Optional Number<br>Cookie Expiration Period. Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new challenge

&#x2022; `custom_page` - Optional String<br>Custom Message for Javascript Challenge. Custom message is of type uri_ref. Currently supported URL schemes is string:///. For string:/// scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. 'Please Wait.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '<p> Please Wait </p>'. Base64 encoded string for this HTML is 'PHA+IFBsZWFzZSBXYWl0IDwvcD4='

&#x2022; `js_script_delay` - Optional Number<br>Javascript Delay. Delay introduced by Javascript, in milliseconds

<a id="l7-ddos-protection"></a>

**L7 DDOS Protection**

&#x2022; `clientside_action_captcha_challenge` - Optional Block<br>Captcha Challenge Parameters. Enables loadbalancer to perform captcha challenge Captcha challenge will be based on Google Recaptcha. With this feature enabled, only clients that pass the captcha challenge will be allowed to complete the HTTP request. When loadbalancer is configured to do Captcha Challenge, it will redirect the browser to an HTML page on every new HTTP request. This HTML page will have captcha challenge embedded in it. Client will be allowed to make the request only if the captcha challenge is successful. Loadbalancer will tag response header with a cookie to avoid Captcha challenge for subsequent requests. CAPTCHA is mainly used as a security check to ensure only human users can pass through. Generally, computers or bots are not capable of solving a captcha. You can enable either Javascript challenge or Captcha challenge on a virtual host<br>See [Clientside Action Captcha Challenge](#l7-ddos-protection-clientside-action-captcha-challenge) below.

&#x2022; `clientside_action_js_challenge` - Optional Block<br>Javascript Challenge Parameters. Enables loadbalancer to perform client browser compatibility test by redirecting to a page with Javascript. With this feature enabled, only clients that are capable of executing Javascript(mostly browsers) will be allowed to complete the HTTP request. When loadbalancer is configured to do Javascript Challenge, it will redirect the browser to an HTML page on every new HTTP request. This HTML page will have Javascript embedded in it. Loadbalancer chooses a set of random numbers for every new client and sends these numbers along with an encrypted answer with the request such that it embed these numbers as input in the Javascript. Javascript will run on the requestor browser and perform a complex Math operation. Script will submit the answer to loadbalancer. Loadbalancer will validate the answer by comparing the calculated answer with the decrypted answer (which was encrypted when it was sent back as reply) and allow the request to the upstream server only if the answer is correct. Loadbalancer will tag response header with a cookie to avoid Javascript challenge for subsequent requests. Javascript challenge serves following purposes * Validate that the request is coming via a browser that is capable for running Javascript * Force the browser to run a complex operation, f(X), that requires it to spend a large number of CPU cycles. This is to slow down a potential DOS attacker by making it difficult to launch a large request flood without having to spend even larger CPU cost at their end. You can enable either Javascript challenge or Captcha challenge on a virtual host<br>See [Clientside Action Js Challenge](#l7-ddos-protection-clientside-action-js-challenge) below.

&#x2022; `clientside_action_none` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `ddos_policy_custom` - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [DDOS Policy Custom](#l7-ddos-protection-ddos-policy-custom) below.

&#x2022; `ddos_policy_none` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `default_rps_threshold` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `mitigation_block` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `mitigation_captcha_challenge` - Optional Block<br>Captcha Challenge Parameters. Enables loadbalancer to perform captcha challenge Captcha challenge will be based on Google Recaptcha. With this feature enabled, only clients that pass the captcha challenge will be allowed to complete the HTTP request. When loadbalancer is configured to do Captcha Challenge, it will redirect the browser to an HTML page on every new HTTP request. This HTML page will have captcha challenge embedded in it. Client will be allowed to make the request only if the captcha challenge is successful. Loadbalancer will tag response header with a cookie to avoid Captcha challenge for subsequent requests. CAPTCHA is mainly used as a security check to ensure only human users can pass through. Generally, computers or bots are not capable of solving a captcha. You can enable either Javascript challenge or Captcha challenge on a virtual host<br>See [Mitigation Captcha Challenge](#l7-ddos-protection-mitigation-captcha-challenge) below.

&#x2022; `mitigation_js_challenge` - Optional Block<br>Javascript Challenge Parameters. Enables loadbalancer to perform client browser compatibility test by redirecting to a page with Javascript. With this feature enabled, only clients that are capable of executing Javascript(mostly browsers) will be allowed to complete the HTTP request. When loadbalancer is configured to do Javascript Challenge, it will redirect the browser to an HTML page on every new HTTP request. This HTML page will have Javascript embedded in it. Loadbalancer chooses a set of random numbers for every new client and sends these numbers along with an encrypted answer with the request such that it embed these numbers as input in the Javascript. Javascript will run on the requestor browser and perform a complex Math operation. Script will submit the answer to loadbalancer. Loadbalancer will validate the answer by comparing the calculated answer with the decrypted answer (which was encrypted when it was sent back as reply) and allow the request to the upstream server only if the answer is correct. Loadbalancer will tag response header with a cookie to avoid Javascript challenge for subsequent requests. Javascript challenge serves following purposes * Validate that the request is coming via a browser that is capable for running Javascript * Force the browser to run a complex operation, f(X), that requires it to spend a large number of CPU cycles. This is to slow down a potential DOS attacker by making it difficult to launch a large request flood without having to spend even larger CPU cost at their end. You can enable either Javascript challenge or Captcha challenge on a virtual host<br>See [Mitigation Js Challenge](#l7-ddos-protection-mitigation-js-challenge) below.

&#x2022; `rps_threshold` - Optional Number<br>Custom. Configure custom RPS threshold

<a id="l7-ddos-protection-clientside-action-captcha-challenge"></a>

**L7 DDOS Protection Clientside Action Captcha Challenge**

&#x2022; `cookie_expiry` - Optional Number<br>Cookie Expiration Period. Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new challenge

&#x2022; `custom_page` - Optional String<br>Custom message for Captcha Challenge. Custom message is of type uri_ref. Currently supported URL schemes is string:///. For string:/// scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. 'Please Wait.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '<p> Please Wait </p>'. Base64 encoded string for this HTML is 'PHA+IFBsZWFzZSBXYWl0IDwvcD4='

<a id="l7-ddos-protection-clientside-action-js-challenge"></a>

**L7 DDOS Protection Clientside Action Js Challenge**

&#x2022; `cookie_expiry` - Optional Number<br>Cookie Expiration Period. Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new challenge

&#x2022; `custom_page` - Optional String<br>Custom Message for Javascript Challenge. Custom message is of type uri_ref. Currently supported URL schemes is string:///. For string:/// scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. 'Please Wait.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '<p> Please Wait </p>'. Base64 encoded string for this HTML is 'PHA+IFBsZWFzZSBXYWl0IDwvcD4='

&#x2022; `js_script_delay` - Optional Number<br>Javascript Delay. Delay introduced by Javascript, in milliseconds

<a id="l7-ddos-protection-ddos-policy-custom"></a>

**L7 DDOS Protection DDOS Policy Custom**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="l7-ddos-protection-mitigation-captcha-challenge"></a>

**L7 DDOS Protection Mitigation Captcha Challenge**

&#x2022; `cookie_expiry` - Optional Number<br>Cookie Expiration Period. Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new challenge

&#x2022; `custom_page` - Optional String<br>Custom message for Captcha Challenge. Custom message is of type uri_ref. Currently supported URL schemes is string:///. For string:/// scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. 'Please Wait.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '<p> Please Wait </p>'. Base64 encoded string for this HTML is 'PHA+IFBsZWFzZSBXYWl0IDwvcD4='

<a id="l7-ddos-protection-mitigation-js-challenge"></a>

**L7 DDOS Protection Mitigation Js Challenge**

&#x2022; `cookie_expiry` - Optional Number<br>Cookie Expiration Period. Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new challenge

&#x2022; `custom_page` - Optional String<br>Custom Message for Javascript Challenge. Custom message is of type uri_ref. Currently supported URL schemes is string:///. For string:/// scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. 'Please Wait.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '<p> Please Wait </p>'. Base64 encoded string for this HTML is 'PHA+IFBsZWFzZSBXYWl0IDwvcD4='

&#x2022; `js_script_delay` - Optional Number<br>Javascript Delay. Delay introduced by Javascript, in milliseconds

<a id="malware-protection-settings"></a>

**Malware Protection Settings**

&#x2022; `malware_protection_rules` - Optional Block<br>Malware Detection Rules. Configure the match criteria to trigger Malware Protection Scan<br>See [Malware Protection Rules](#malware-protection-settings-malware-protection-rules) below.

<a id="malware-protection-settings-malware-protection-rules"></a>

**Malware Protection Settings Malware Protection Rules**

&#x2022; `action` - Optional Block<br>Action<br>See [Action](#malware-protection-settings-malware-protection-rules-action) below.

&#x2022; `domain` - Optional Block<br>Domain to Match. Domain to be matched<br>See [Domain](#malware-protection-settings-malware-protection-rules-domain) below.

&#x2022; `http_methods` - Optional List  Defaults to `ANY`<br>Possible values are `ANY`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`, `COPY`<br>HTTP Methods. Methods to be matched

&#x2022; `metadata` - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#malware-protection-settings-malware-protection-rules-metadata) below.

&#x2022; `path` - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#malware-protection-settings-malware-protection-rules-path) below.

<a id="malware-protection-settings-malware-protection-rules-action"></a>

**Malware Protection Settings Malware Protection Rules Action**

&#x2022; `block` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `report` - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="malware-protection-settings-malware-protection-rules-domain"></a>

**Malware Protection Settings Malware Protection Rules Domain**

&#x2022; `any_domain` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `domain` - Optional Block<br>Domains. Domains names<br>See [Domain](#malware-protection-settings-malware-protection-rules-domain-domain) below.

<a id="malware-protection-settings-malware-protection-rules-domain-domain"></a>

**Malware Protection Settings Malware Protection Rules Domain Domain**

&#x2022; `exact_value` - Optional String<br>Exact Value. Exact domain name

&#x2022; `regex_value` - Optional String<br>Regex Values of Domains. Regular Expression value for the domain name

&#x2022; `suffix_value` - Optional String<br>Suffix Value. Suffix of domain name e.g 'xyz.com' will match '*.xyz.com' and 'xyz.com'

<a id="malware-protection-settings-malware-protection-rules-metadata"></a>

**Malware Protection Settings Malware Protection Rules Metadata**

&#x2022; `description` - Optional String<br>Description. Human readable description

&#x2022; `name` - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

<a id="malware-protection-settings-malware-protection-rules-path"></a>

**Malware Protection Settings Malware Protection Rules Path**

&#x2022; `path` - Optional String<br>Exact. Exact path value to match

&#x2022; `prefix` - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

&#x2022; `regex` - Optional String<br>Regex. Regular expression of path match (e.g. the value .* will match on all paths)

<a id="more-option"></a>

**More Option**

&#x2022; `buffer_policy` - Optional Block<br>Buffer Configuration. Some upstream applications are not capable of handling streamed data. This config enables buffering the entire request before sending to upstream application. We can specify the maximum buffer size and buffer interval with this config. Buffering can be enabled and disabled at VirtualHost and Route levels Route level buffer configuration takes precedence<br>See [Buffer Policy](#more-option-buffer-policy) below.

&#x2022; `compression_params` - Optional Block<br>Compression Parameters. Enables loadbalancer to compress dispatched data from an upstream service upon client request. The content is compressed and then sent to the client with the appropriate headers if either response and request allow. Only GZIP compression is supported. By default compression will be skipped when: A request does NOT contain accept-encoding header. A request includes accept-encoding header, but it does not contain “gzip” or “*”. A request includes accept-encoding with “gzip” or “*” with the weight “q=0”. Note that the “gzip” will have a higher weight then “*”. For example, if accept-encoding is “gzip;q=0,*;q=1”, the filter will not compress. But if the header is set to “*;q=0,gzip;q=1”, the filter will compress. A request whose accept-encoding header includes “identity”. A response contains a content-encoding header. A response contains a cache-control header whose value includes “no-transform”. A response contains a transfer-encoding header whose value includes “gzip”. A response does not contain a content-type value that matches one of the selected mime-types, which default to application/javascript, application/JSON, application/xhtml+XML, image/svg+XML, text/CSS, text/HTML, text/plain, text/XML. Neither content-length nor transfer-encoding headers are present in the response. Response size is smaller than 30 bytes (only applicable when transfer-encoding is not chunked). When compression is applied: The content-length is removed from response headers. Response headers contain “transfer-encoding: chunked” and do not contain “content-encoding” header. The “vary: accept-encoding” header is inserted on every response. GZIP Compression Level: A value which is optimal balance between speed of compression and amount of compression is chosen<br>See [Compression Params](#more-option-compression-params) below.

&#x2022; `custom_errors` - Optional Block<br>Custom Error Responses. Map of integer error codes as keys and string values that can be used to provide custom HTTP pages for each error code. Key of the map can be either response code class or HTTP Error code. Response code classes for key is configured as follows 3 -- for 3xx response code class 4 -- for 4xx response code class 5 -- for 5xx response code class Value of the map is string which represents custom HTTP responses. Specific response code takes preference when both response code and response code class matches for a request

&#x2022; `disable_default_error_pages` - Optional Bool<br>Disable Default Error Pages. Disable the use of default F5XC error pages

&#x2022; `disable_path_normalize` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `enable_path_normalize` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `idle_timeout` - Optional Number<br>Idle Timeout. The amount of time that a stream can exist without upstream or downstream activity, in milliseconds. The stream is terminated with a HTTP 504 (Gateway Timeout) error code if no upstream response header has been received, otherwise the stream is reset

&#x2022; `max_request_header_size` - Optional Number<br>Maximum Request Header Size. The maximum request header size for downstream connections, in KiB. A HTTP 431 (Request Header Fields Too Large) error code is sent for requests that exceed this size. If multiple load balancers share the same advertise_policy, the highest value configured across all such load balancers is used for all the load balancers in question

&#x2022; `request_cookies_to_add` - Optional Block<br>Add Cookies in Cookie Header. Cookies are key-value pairs to be added to HTTP request being routed towards upstream. Cookies specified at this level are applied after cookies from matched Route are applied<br>See [Request Cookies To Add](#more-option-request-cookies-to-add) below.

&#x2022; `request_cookies_to_remove` - Optional List<br>Remove Cookies from Cookie Header. List of keys of Cookies to be removed from the HTTP request being sent towards upstream

&#x2022; `request_headers_to_add` - Optional Block<br>Add Request Headers. Headers are key-value pairs to be added to HTTP request being routed towards upstream. Headers specified at this level are applied after headers from matched Route are applied<br>See [Request Headers To Add](#more-option-request-headers-to-add) below.

&#x2022; `request_headers_to_remove` - Optional List<br>Remove Request Headers. List of keys of Headers to be removed from the HTTP request being sent towards upstream

&#x2022; `response_cookies_to_add` - Optional Block<br>Add Set-Cookie Headers. Cookies are name-value pairs along with optional attribute parameters to be added to HTTP response being sent towards downstream. Cookies specified at this level are applied after cookies from matched Route are applied<br>See [Response Cookies To Add](#more-option-response-cookies-to-add) below.

&#x2022; `response_cookies_to_remove` - Optional List<br>Remove Cookies from Set-Cookie Headers. List of name of Cookies to be removed from the HTTP response being sent towards downstream. Entire set-cookie header will be removed

&#x2022; `response_headers_to_add` - Optional Block<br>Add Response Headers. Headers are key-value pairs to be added to HTTP response being sent towards downstream. Headers specified at this level are applied after headers from matched Route are applied<br>See [Response Headers To Add](#more-option-response-headers-to-add) below.

&#x2022; `response_headers_to_remove` - Optional List<br>Remove Response Headers. List of keys of Headers to be removed from the HTTP response being sent towards downstream

<a id="more-option-buffer-policy"></a>

**More Option Buffer Policy**

&#x2022; `disabled` - Optional Bool<br>Disable. Disable buffering for a particular route. This is useful when virtual-host has buffering, but we need to disable it on a specific route. The value of this field is ignored for virtual-host

&#x2022; `max_request_bytes` - Optional Number<br>Max Request Bytes. The maximum request size that the filter will buffer before the connection manager will stop buffering and return a RequestEntityTooLarge (413) response

<a id="more-option-compression-params"></a>

**More Option Compression Params**

&#x2022; `content_length` - Optional Number  Defaults to `30`<br>Content Length. Minimum response length, in bytes, which will trigger compression. The

&#x2022; `content_type` - Optional List<br>Content Type. Set of strings that allows specifying which mime-types yield compression When this field is not defined, compression will be applied to the following mime-types: 'application/javascript' 'application/JSON', 'application/xhtml+XML' 'image/svg+XML' 'text/CSS' 'text/HTML' 'text/plain' 'text/XML'

&#x2022; `disable_on_etag_header` - Optional Bool<br>Disable On Etag Header. If true, disables compression when the response contains an etag header. When it is false, weak etags will be preserved and the ones that require strong validation will be removed

&#x2022; `remove_accept_encoding_header` - Optional Bool<br>Remove Accept-Encoding Header. If true, removes accept-encoding from the request headers before dispatching it to the upstream so that responses do not get compressed before reaching the filter

<a id="more-option-request-cookies-to-add"></a>

**More Option Request Cookies To Add**

&#x2022; `name` - Optional String<br>Name. Name of the cookie in Cookie header

&#x2022; `overwrite` - Optional Bool  Defaults to `do`<br>Overwrite. Should the value be overwritten? If true, the value is overwritten to existing values. not overwrite

&#x2022; `secret_value` - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Secret Value](#more-option-request-cookies-to-add-secret-value) below.

&#x2022; `value` - Optional String<br>Value. Value of the Cookie header

<a id="more-option-request-cookies-to-add-secret-value"></a>

**More Option Request Cookies To Add Secret Value**

&#x2022; `blindfold_secret_info` - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#more-option-request-cookies-to-add-secret-value-blindfold-secret-info) below.

&#x2022; `clear_secret_info` - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#more-option-request-cookies-to-add-secret-value-clear-secret-info) below.

<a id="more-option-request-cookies-to-add-secret-value-blindfold-secret-info"></a>

**More Option Request Cookies To Add Secret Value Blindfold Secret Info**

&#x2022; `decryption_provider` - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

&#x2022; `location` - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

&#x2022; `store_provider` - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

<a id="more-option-request-cookies-to-add-secret-value-clear-secret-info"></a>

**More Option Request Cookies To Add Secret Value Clear Secret Info**

&#x2022; `provider_ref` - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

&#x2022; `url` - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

<a id="more-option-request-headers-to-add"></a>

**More Option Request Headers To Add**

&#x2022; `append` - Optional Bool  Defaults to `do`<br>Append. Should the value be appended? If true, the value is appended to existing values. not append

&#x2022; `name` - Optional String<br>Name. Name of the HTTP header

&#x2022; `secret_value` - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Secret Value](#more-option-request-headers-to-add-secret-value) below.

&#x2022; `value` - Optional String<br>Value. Value of the HTTP header

<a id="more-option-request-headers-to-add-secret-value"></a>

**More Option Request Headers To Add Secret Value**

&#x2022; `blindfold_secret_info` - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#more-option-request-headers-to-add-secret-value-blindfold-secret-info) below.

&#x2022; `clear_secret_info` - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#more-option-request-headers-to-add-secret-value-clear-secret-info) below.

<a id="more-option-request-headers-to-add-secret-value-blindfold-secret-info"></a>

**More Option Request Headers To Add Secret Value Blindfold Secret Info**

&#x2022; `decryption_provider` - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

&#x2022; `location` - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

&#x2022; `store_provider` - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

<a id="more-option-request-headers-to-add-secret-value-clear-secret-info"></a>

**More Option Request Headers To Add Secret Value Clear Secret Info**

&#x2022; `provider_ref` - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

&#x2022; `url` - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

<a id="more-option-response-cookies-to-add"></a>

**More Option Response Cookies To Add**

&#x2022; `add_domain` - Optional String<br>Add Domain. Add domain attribute

&#x2022; `add_expiry` - Optional String<br>Add expiry. Add expiry attribute

&#x2022; `add_httponly` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `add_partitioned` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `add_path` - Optional String<br>Add path. Add path attribute

&#x2022; `add_secure` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `ignore_domain` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `ignore_expiry` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `ignore_httponly` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `ignore_max_age` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `ignore_partitioned` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `ignore_path` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `ignore_samesite` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `ignore_secure` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `ignore_value` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `max_age_value` - Optional Number<br>Add Max Age. Add max age attribute

&#x2022; `name` - Optional String<br>Name. Name of the cookie in Cookie header

&#x2022; `overwrite` - Optional Bool  Defaults to `do`<br>Overwrite. Should the value be overwritten? If true, the value is overwritten to existing values. not overwrite

&#x2022; `samesite_lax` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `samesite_none` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `samesite_strict` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `secret_value` - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Secret Value](#more-option-response-cookies-to-add-secret-value) below.

&#x2022; `value` - Optional String<br>Value. Value of the Cookie header

<a id="more-option-response-cookies-to-add-secret-value"></a>

**More Option Response Cookies To Add Secret Value**

&#x2022; `blindfold_secret_info` - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#more-option-response-cookies-to-add-secret-value-blindfold-secret-info) below.

&#x2022; `clear_secret_info` - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#more-option-response-cookies-to-add-secret-value-clear-secret-info) below.

<a id="more-option-response-cookies-to-add-secret-value-blindfold-secret-info"></a>

**More Option Response Cookies To Add Secret Value Blindfold Secret Info**

&#x2022; `decryption_provider` - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

&#x2022; `location` - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

&#x2022; `store_provider` - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

<a id="more-option-response-cookies-to-add-secret-value-clear-secret-info"></a>

**More Option Response Cookies To Add Secret Value Clear Secret Info**

&#x2022; `provider_ref` - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

&#x2022; `url` - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

<a id="more-option-response-headers-to-add"></a>

**More Option Response Headers To Add**

&#x2022; `append` - Optional Bool  Defaults to `do`<br>Append. Should the value be appended? If true, the value is appended to existing values. not append

&#x2022; `name` - Optional String<br>Name. Name of the HTTP header

&#x2022; `secret_value` - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Secret Value](#more-option-response-headers-to-add-secret-value) below.

&#x2022; `value` - Optional String<br>Value. Value of the HTTP header

<a id="more-option-response-headers-to-add-secret-value"></a>

**More Option Response Headers To Add Secret Value**

&#x2022; `blindfold_secret_info` - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#more-option-response-headers-to-add-secret-value-blindfold-secret-info) below.

&#x2022; `clear_secret_info` - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#more-option-response-headers-to-add-secret-value-clear-secret-info) below.

<a id="more-option-response-headers-to-add-secret-value-blindfold-secret-info"></a>

**More Option Response Headers To Add Secret Value Blindfold Secret Info**

&#x2022; `decryption_provider` - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

&#x2022; `location` - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

&#x2022; `store_provider` - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

<a id="more-option-response-headers-to-add-secret-value-clear-secret-info"></a>

**More Option Response Headers To Add Secret Value Clear Secret Info**

&#x2022; `provider_ref` - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

&#x2022; `url` - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

<a id="origin-server-subset-rule-list"></a>

**Origin Server Subset Rule List**

&#x2022; `origin_server_subset_rules` - Optional Block<br>Origin Server Subset Rules. Origin Server Subset Rules allow users to define match condition on Client (IP address, ASN, Country), IP Reputation, Regional Edge names, Request for subset selection of origin servers. Origin Server Subset is a sequential engine where rules are evaluated one after the other. It's important to define the correct order for Origin Server Subset to get the intended result, rules are evaluated from top to bottom in the list. When an Origin server subset rule is matched, then this selection rule takes effect and no more rules are evaluated<br>See [Origin Server Subset Rules](#origin-server-subset-rule-list-origin-server-subset-rules) below.

<a id="origin-server-subset-rule-list-origin-server-subset-rules"></a>

**Origin Server Subset Rule List Origin Server Subset Rules**

&#x2022; `any_asn` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `any_ip` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `asn_list` - Optional Block<br>ASN Match List. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer<br>See [Asn List](#origin-server-subset-rule-list-origin-server-subset-rules-asn-list) below.

&#x2022; `asn_matcher` - Optional Block<br>ASN Matcher. Match any AS number contained in the list of bgp_asn_sets<br>See [Asn Matcher](#origin-server-subset-rule-list-origin-server-subset-rules-asn-matcher) below.

&#x2022; `client_selector` - Optional Block<br>Label Selector. This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expresssions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects. Label selector is immutable. expressions is a list of strings of label selection expression. Each string has ',' separated values which are 'AND' and all strings are logically 'OR'. BNF for expression string <selector-syntax> ::= <requirement> | <requirement> ',' <selector-syntax> <requirement> ::= [!] KEY [ <set-based-restriction> | <exact-match-restriction> ] <set-based-restriction> ::= '' | <inclusion-exclusion> <value-set> <inclusion-exclusion> ::= <inclusion> | <exclusion> <exclusion> ::= 'notin' <inclusion> ::= 'in' <value-set> ::= '(' <values> ')' <values> ::= VALUE | VALUE ',' <values> <exact-match-restriction> ::= ['='|'=='|'!='] VALUE<br>See [Client Selector](#origin-server-subset-rule-list-origin-server-subset-rules-client-selector) below.

&#x2022; `country_codes` - Optional List  Defaults to `COUNTRY_NONE`<br>Possible values are `COUNTRY_NONE`, `COUNTRY_AD`, `COUNTRY_AE`, `COUNTRY_AF`, `COUNTRY_AG`, `COUNTRY_AI`, `COUNTRY_AL`, `COUNTRY_AM`, `COUNTRY_AN`, `COUNTRY_AO`, `COUNTRY_AQ`, `COUNTRY_AR`, `COUNTRY_AS`, `COUNTRY_AT`, `COUNTRY_AU`, `COUNTRY_AW`, `COUNTRY_AX`, `COUNTRY_AZ`, `COUNTRY_BA`, `COUNTRY_BB`, `COUNTRY_BD`, `COUNTRY_BE`, `COUNTRY_BF`, `COUNTRY_BG`, `COUNTRY_BH`, `COUNTRY_BI`, `COUNTRY_BJ`, `COUNTRY_BL`, `COUNTRY_BM`, `COUNTRY_BN`, `COUNTRY_BO`, `COUNTRY_BQ`, `COUNTRY_BR`, `COUNTRY_BS`, `COUNTRY_BT`, `COUNTRY_BV`, `COUNTRY_BW`, `COUNTRY_BY`, `COUNTRY_BZ`, `COUNTRY_CA`, `COUNTRY_CC`, `COUNTRY_CD`, `COUNTRY_CF`, `COUNTRY_CG`, `COUNTRY_CH`, `COUNTRY_CI`, `COUNTRY_CK`, `COUNTRY_CL`, `COUNTRY_CM`, `COUNTRY_CN`, `COUNTRY_CO`, `COUNTRY_CR`, `COUNTRY_CS`, `COUNTRY_CU`, `COUNTRY_CV`, `COUNTRY_CW`, `COUNTRY_CX`, `COUNTRY_CY`, `COUNTRY_CZ`, `COUNTRY_DE`, `COUNTRY_DJ`, `COUNTRY_DK`, `COUNTRY_DM`, `COUNTRY_DO`, `COUNTRY_DZ`, `COUNTRY_EC`, `COUNTRY_EE`, `COUNTRY_EG`, `COUNTRY_EH`, `COUNTRY_ER`, `COUNTRY_ES`, `COUNTRY_ET`, `COUNTRY_FI`, `COUNTRY_FJ`, `COUNTRY_FK`, `COUNTRY_FM`, `COUNTRY_FO`, `COUNTRY_FR`, `COUNTRY_GA`, `COUNTRY_GB`, `COUNTRY_GD`, `COUNTRY_GE`, `COUNTRY_GF`, `COUNTRY_GG`, `COUNTRY_GH`, `COUNTRY_GI`, `COUNTRY_GL`, `COUNTRY_GM`, `COUNTRY_GN`, `COUNTRY_GP`, `COUNTRY_GQ`, `COUNTRY_GR`, `COUNTRY_GS`, `COUNTRY_GT`, `COUNTRY_GU`, `COUNTRY_GW`, `COUNTRY_GY`, `COUNTRY_HK`, `COUNTRY_HM`, `COUNTRY_HN`, `COUNTRY_HR`, `COUNTRY_HT`, `COUNTRY_HU`, `COUNTRY_ID`, `COUNTRY_IE`, `COUNTRY_IL`, `COUNTRY_IM`, `COUNTRY_IN`, `COUNTRY_IO`, `COUNTRY_IQ`, `COUNTRY_IR`, `COUNTRY_IS`, `COUNTRY_IT`, `COUNTRY_JE`, `COUNTRY_JM`, `COUNTRY_JO`, `COUNTRY_JP`, `COUNTRY_KE`, `COUNTRY_KG`, `COUNTRY_KH`, `COUNTRY_KI`, `COUNTRY_KM`, `COUNTRY_KN`, `COUNTRY_KP`, `COUNTRY_KR`, `COUNTRY_KW`, `COUNTRY_KY`, `COUNTRY_KZ`, `COUNTRY_LA`, `COUNTRY_LB`, `COUNTRY_LC`, `COUNTRY_LI`, `COUNTRY_LK`, `COUNTRY_LR`, `COUNTRY_LS`, `COUNTRY_LT`, `COUNTRY_LU`, `COUNTRY_LV`, `COUNTRY_LY`, `COUNTRY_MA`, `COUNTRY_MC`, `COUNTRY_MD`, `COUNTRY_ME`, `COUNTRY_MF`, `COUNTRY_MG`, `COUNTRY_MH`, `COUNTRY_MK`, `COUNTRY_ML`, `COUNTRY_MM`, `COUNTRY_MN`, `COUNTRY_MO`, `COUNTRY_MP`, `COUNTRY_MQ`, `COUNTRY_MR`, `COUNTRY_MS`, `COUNTRY_MT`, `COUNTRY_MU`, `COUNTRY_MV`, `COUNTRY_MW`, `COUNTRY_MX`, `COUNTRY_MY`, `COUNTRY_MZ`, `COUNTRY_NA`, `COUNTRY_NC`, `COUNTRY_NE`, `COUNTRY_NF`, `COUNTRY_NG`, `COUNTRY_NI`, `COUNTRY_NL`, `COUNTRY_NO`, `COUNTRY_NP`, `COUNTRY_NR`, `COUNTRY_NU`, `COUNTRY_NZ`, `COUNTRY_OM`, `COUNTRY_PA`, `COUNTRY_PE`, `COUNTRY_PF`, `COUNTRY_PG`, `COUNTRY_PH`, `COUNTRY_PK`, `COUNTRY_PL`, `COUNTRY_PM`, `COUNTRY_PN`, `COUNTRY_PR`, `COUNTRY_PS`, `COUNTRY_PT`, `COUNTRY_PW`, `COUNTRY_PY`, `COUNTRY_QA`, `COUNTRY_RE`, `COUNTRY_RO`, `COUNTRY_RS`, `COUNTRY_RU`, `COUNTRY_RW`, `COUNTRY_SA`, `COUNTRY_SB`, `COUNTRY_SC`, `COUNTRY_SD`, `COUNTRY_SE`, `COUNTRY_SG`, `COUNTRY_SH`, `COUNTRY_SI`, `COUNTRY_SJ`, `COUNTRY_SK`, `COUNTRY_SL`, `COUNTRY_SM`, `COUNTRY_SN`, `COUNTRY_SO`, `COUNTRY_SR`, `COUNTRY_SS`, `COUNTRY_ST`, `COUNTRY_SV`, `COUNTRY_SX`, `COUNTRY_SY`, `COUNTRY_SZ`, `COUNTRY_TC`, `COUNTRY_TD`, `COUNTRY_TF`, `COUNTRY_TG`, `COUNTRY_TH`, `COUNTRY_TJ`, `COUNTRY_TK`, `COUNTRY_TL`, `COUNTRY_TM`, `COUNTRY_TN`, `COUNTRY_TO`, `COUNTRY_TR`, `COUNTRY_TT`, `COUNTRY_TV`, `COUNTRY_TW`, `COUNTRY_TZ`, `COUNTRY_UA`, `COUNTRY_UG`, `COUNTRY_UM`, `COUNTRY_US`, `COUNTRY_UY`, `COUNTRY_UZ`, `COUNTRY_VA`, `COUNTRY_VC`, `COUNTRY_VE`, `COUNTRY_VG`, `COUNTRY_VI`, `COUNTRY_VN`, `COUNTRY_VU`, `COUNTRY_WF`, `COUNTRY_WS`, `COUNTRY_XK`, `COUNTRY_XT`, `COUNTRY_YE`, `COUNTRY_YT`, `COUNTRY_ZA`, `COUNTRY_ZM`, `COUNTRY_ZW`<br>Country Codes List. List of Country Codes

&#x2022; `ip_matcher` - Optional Block<br>IP Prefix Matcher. Match any IP prefix contained in the list of ip_prefix_sets. The result of the match is inverted if invert_matcher is true<br>See [IP Matcher](#origin-server-subset-rule-list-origin-server-subset-rules-ip-matcher) below.

&#x2022; `ip_prefix_list` - Optional Block<br>IP Prefix Match List. List of IP Prefix strings to match against<br>See [IP Prefix List](#origin-server-subset-rule-list-origin-server-subset-rules-ip-prefix-list) below.

&#x2022; `metadata` - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#origin-server-subset-rule-list-origin-server-subset-rules-metadata) below.

&#x2022; `none` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `origin_server_subsets_action` - Optional Block<br>Action. Add labels to select one or more origin servers. Note: The pre-requisite settings to be configured in the origin pool are: 1. Add labels to origin servers 2. Enable subset load balancing in the Origin Server Subsets section and configure keys in origin server subsets classes

&#x2022; `re_name_list` - Optional List<br>RE Names. List of RE names for match

<a id="origin-server-subset-rule-list-origin-server-subset-rules-asn-list"></a>

**Origin Server Subset Rule List Origin Server Subset Rules Asn List**

&#x2022; `as_numbers` - Optional List<br>AS Numbers. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer

<a id="origin-server-subset-rule-list-origin-server-subset-rules-asn-matcher"></a>

**Origin Server Subset Rule List Origin Server Subset Rules Asn Matcher**

&#x2022; `asn_sets` - Optional Block<br>BGP ASN Sets. A list of references to bgp_asn_set objects<br>See [Asn Sets](#origin-server-subset-rule-list-origin-server-subset-rules-asn-matcher-asn-sets) below.

<a id="origin-server-subset-rule-list-origin-server-subset-rules-asn-matcher-asn-sets"></a>

**Origin Server Subset Rule List Origin Server Subset Rules Asn Matcher Asn Sets**

&#x2022; `kind` - Optional String<br>Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route')

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

&#x2022; `uid` - Optional String<br>UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid

<a id="origin-server-subset-rule-list-origin-server-subset-rules-client-selector"></a>

**Origin Server Subset Rule List Origin Server Subset Rules Client Selector**

&#x2022; `expressions` - Optional List<br>Selector Expression. expressions contains the kubernetes style label expression for selections

<a id="origin-server-subset-rule-list-origin-server-subset-rules-ip-matcher"></a>

**Origin Server Subset Rule List Origin Server Subset Rules IP Matcher**

&#x2022; `invert_matcher` - Optional Bool<br>Invert IP Matcher. Invert the match result

&#x2022; `prefix_sets` - Optional Block<br>IP Prefix Sets. A list of references to ip_prefix_set objects<br>See [Prefix Sets](#origin-server-subset-rule-list-origin-server-subset-rules-ip-matcher-prefix-sets) below.

<a id="origin-server-subset-rule-list-origin-server-subset-rules-ip-matcher-prefix-sets"></a>

**Origin Server Subset Rule List Origin Server Subset Rules IP Matcher Prefix Sets**

&#x2022; `kind` - Optional String<br>Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route')

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

&#x2022; `uid` - Optional String<br>UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid

<a id="origin-server-subset-rule-list-origin-server-subset-rules-ip-prefix-list"></a>

**Origin Server Subset Rule List Origin Server Subset Rules IP Prefix List**

&#x2022; `invert_match` - Optional Bool<br>Invert Match Result. Invert the match result

&#x2022; `ip_prefixes` - Optional List<br>IPv4 Prefix List. List of IPv4 prefix strings

<a id="origin-server-subset-rule-list-origin-server-subset-rules-metadata"></a>

**Origin Server Subset Rule List Origin Server Subset Rules Metadata**

&#x2022; `description` - Optional String<br>Description. Human readable description

&#x2022; `name` - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

<a id="policy-based-challenge"></a>

**Policy Based Challenge**

&#x2022; `always_enable_captcha_challenge` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `always_enable_js_challenge` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `captcha_challenge_parameters` - Optional Block<br>Captcha Challenge Parameters. Enables loadbalancer to perform captcha challenge Captcha challenge will be based on Google Recaptcha. With this feature enabled, only clients that pass the captcha challenge will be allowed to complete the HTTP request. When loadbalancer is configured to do Captcha Challenge, it will redirect the browser to an HTML page on every new HTTP request. This HTML page will have captcha challenge embedded in it. Client will be allowed to make the request only if the captcha challenge is successful. Loadbalancer will tag response header with a cookie to avoid Captcha challenge for subsequent requests. CAPTCHA is mainly used as a security check to ensure only human users can pass through. Generally, computers or bots are not capable of solving a captcha. You can enable either Javascript challenge or Captcha challenge on a virtual host<br>See [Captcha Challenge Parameters](#policy-based-challenge-captcha-challenge-parameters) below.

&#x2022; `default_captcha_challenge_parameters` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `default_js_challenge_parameters` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `default_mitigation_settings` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `default_temporary_blocking_parameters` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `js_challenge_parameters` - Optional Block<br>Javascript Challenge Parameters. Enables loadbalancer to perform client browser compatibility test by redirecting to a page with Javascript. With this feature enabled, only clients that are capable of executing Javascript(mostly browsers) will be allowed to complete the HTTP request. When loadbalancer is configured to do Javascript Challenge, it will redirect the browser to an HTML page on every new HTTP request. This HTML page will have Javascript embedded in it. Loadbalancer chooses a set of random numbers for every new client and sends these numbers along with an encrypted answer with the request such that it embed these numbers as input in the Javascript. Javascript will run on the requestor browser and perform a complex Math operation. Script will submit the answer to loadbalancer. Loadbalancer will validate the answer by comparing the calculated answer with the decrypted answer (which was encrypted when it was sent back as reply) and allow the request to the upstream server only if the answer is correct. Loadbalancer will tag response header with a cookie to avoid Javascript challenge for subsequent requests. Javascript challenge serves following purposes * Validate that the request is coming via a browser that is capable for running Javascript * Force the browser to run a complex operation, f(X), that requires it to spend a large number of CPU cycles. This is to slow down a potential DOS attacker by making it difficult to launch a large request flood without having to spend even larger CPU cost at their end. You can enable either Javascript challenge or Captcha challenge on a virtual host<br>See [Js Challenge Parameters](#policy-based-challenge-js-challenge-parameters) below.

&#x2022; `malicious_user_mitigation` - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Malicious User Mitigation](#policy-based-challenge-malicious-user-mitigation) below.

&#x2022; `no_challenge` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `rule_list` - Optional Block<br>Challenge Rule List. List of challenge rules to be used in policy based challenge<br>See [Rule List](#policy-based-challenge-rule-list) below.

&#x2022; `temporary_user_blocking` - Optional Block<br>Temporary User Blocking. Specifies configuration for temporary user blocking resulting from user behavior analysis. When Malicious User Mitigation is enabled from service policy rules, users' accessing the application will be analyzed for malicious activity and the configured mitigation actions will be taken on identified malicious users. These mitigation actions include setting up temporary blocking on that user. This configuration specifies settings on how that blocking should be done by the loadbalancer<br>See [Temporary User Blocking](#policy-based-challenge-temporary-user-blocking) below.

<a id="policy-based-challenge-captcha-challenge-parameters"></a>

**Policy Based Challenge Captcha Challenge Parameters**

&#x2022; `cookie_expiry` - Optional Number<br>Cookie Expiration Period. Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new challenge

&#x2022; `custom_page` - Optional String<br>Custom message for Captcha Challenge. Custom message is of type uri_ref. Currently supported URL schemes is string:///. For string:/// scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. 'Please Wait.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '<p> Please Wait </p>'. Base64 encoded string for this HTML is 'PHA+IFBsZWFzZSBXYWl0IDwvcD4='

<a id="policy-based-challenge-js-challenge-parameters"></a>

**Policy Based Challenge Js Challenge Parameters**

&#x2022; `cookie_expiry` - Optional Number<br>Cookie Expiration Period. Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new challenge

&#x2022; `custom_page` - Optional String<br>Custom Message for Javascript Challenge. Custom message is of type uri_ref. Currently supported URL schemes is string:///. For string:/// scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. 'Please Wait.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '<p> Please Wait </p>'. Base64 encoded string for this HTML is 'PHA+IFBsZWFzZSBXYWl0IDwvcD4='

&#x2022; `js_script_delay` - Optional Number<br>Javascript Delay. Delay introduced by Javascript, in milliseconds

<a id="policy-based-challenge-malicious-user-mitigation"></a>

**Policy Based Challenge Malicious User Mitigation**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="policy-based-challenge-rule-list"></a>

**Policy Based Challenge Rule List**

&#x2022; `rules` - Optional Block<br>Rules. Rules that specify the match conditions and challenge type to be launched. When a challenge type is selected to be always enabled, these rules can be used to disable challenge or launch a different challenge for requests that match the specified conditions<br>See [Rules](#policy-based-challenge-rule-list-rules) below.

<a id="policy-based-challenge-rule-list-rules"></a>

**Policy Based Challenge Rule List Rules**

&#x2022; `metadata` - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#policy-based-challenge-rule-list-rules-metadata) below.

&#x2022; `spec` - Optional Block<br>Challenge Rule Specification. A Challenge Rule consists of an unordered list of predicates and an action. The predicates are evaluated against a set of input fields that are extracted from or derived from an L7 request API. A request API is considered to match the rule if all predicates in the rule evaluate to true for that request. Any predicates that are not specified in a rule are implicitly considered to be true. If a request API matches a challenge rule, the configured challenge is enforced<br>See [Spec](#policy-based-challenge-rule-list-rules-spec) below.

<a id="policy-based-challenge-rule-list-rules-metadata"></a>

**Policy Based Challenge Rule List Rules Metadata**

&#x2022; `description` - Optional String<br>Description. Human readable description

&#x2022; `name` - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

<a id="policy-based-challenge-rule-list-rules-spec"></a>

**Policy Based Challenge Rule List Rules Spec**

&#x2022; `any_asn` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `any_client` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `any_ip` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `arg_matchers` - Optional Block<br>A list of predicates for all POST args that need to be matched. The criteria for matching each arg are described in individual instances of ArgMatcherType. The actual arg values are extracted from the request API as a list of strings for each arg selector name. Note that all specified arg matcher predicates must evaluate to true<br>See [Arg Matchers](#policy-based-challenge-rule-list-rules-spec-arg-matchers) below.

&#x2022; `asn_list` - Optional Block<br>ASN Match List. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer<br>See [Asn List](#policy-based-challenge-rule-list-rules-spec-asn-list) below.

&#x2022; `asn_matcher` - Optional Block<br>ASN Matcher. Match any AS number contained in the list of bgp_asn_sets<br>See [Asn Matcher](#policy-based-challenge-rule-list-rules-spec-asn-matcher) below.

&#x2022; `body_matcher` - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Body Matcher](#policy-based-challenge-rule-list-rules-spec-body-matcher) below.

&#x2022; `client_selector` - Optional Block<br>Label Selector. This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expresssions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects. Label selector is immutable. expressions is a list of strings of label selection expression. Each string has ',' separated values which are 'AND' and all strings are logically 'OR'. BNF for expression string <selector-syntax> ::= <requirement> | <requirement> ',' <selector-syntax> <requirement> ::= [!] KEY [ <set-based-restriction> | <exact-match-restriction> ] <set-based-restriction> ::= '' | <inclusion-exclusion> <value-set> <inclusion-exclusion> ::= <inclusion> | <exclusion> <exclusion> ::= 'notin' <inclusion> ::= 'in' <value-set> ::= '(' <values> ')' <values> ::= VALUE | VALUE ',' <values> <exact-match-restriction> ::= ['='|'=='|'!='] VALUE<br>See [Client Selector](#policy-based-challenge-rule-list-rules-spec-client-selector) below.

&#x2022; `cookie_matchers` - Optional Block<br>A list of predicates for all cookies that need to be matched. The criteria for matching each cookie is described in individual instances of CookieMatcherType. The actual cookie values are extracted from the request API as a list of strings for each cookie name. Note that all specified cookie matcher predicates must evaluate to true<br>See [Cookie Matchers](#policy-based-challenge-rule-list-rules-spec-cookie-matchers) below.

&#x2022; `disable_challenge` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `domain_matcher` - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Domain Matcher](#policy-based-challenge-rule-list-rules-spec-domain-matcher) below.

&#x2022; `enable_captcha_challenge` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `enable_javascript_challenge` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `expiration_timestamp` - Optional String<br>The expiration_timestamp is the RFC 3339 format timestamp at which the containing rule is considered to be logically expired. The rule continues to exist in the configuration but is not applied anymore

&#x2022; `headers` - Optional Block<br>A list of predicates for various HTTP headers that need to match. The criteria for matching each HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values are extracted from the request API as a list of strings for each HTTP header type. Note that all specified header predicates must evaluate to true<br>See [Headers](#policy-based-challenge-rule-list-rules-spec-headers) below.

&#x2022; `http_method` - Optional Block<br>HTTP Method Matcher. A HTTP method matcher specifies a list of methods to match an input HTTP method. The match is considered successful if the input method is a member of the list. The result of the match based on the method list is inverted if invert_matcher is true<br>See [HTTP Method](#policy-based-challenge-rule-list-rules-spec-http-method) below.

&#x2022; `ip_matcher` - Optional Block<br>IP Prefix Matcher. Match any IP prefix contained in the list of ip_prefix_sets. The result of the match is inverted if invert_matcher is true<br>See [IP Matcher](#policy-based-challenge-rule-list-rules-spec-ip-matcher) below.

&#x2022; `ip_prefix_list` - Optional Block<br>IP Prefix Match List. List of IP Prefix strings to match against<br>See [IP Prefix List](#policy-based-challenge-rule-list-rules-spec-ip-prefix-list) below.

&#x2022; `path` - Optional Block<br>Path Matcher. A path matcher specifies multiple criteria for matching an HTTP path string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of path prefixes, a list of exact path values and a list of regular expressions<br>See [Path](#policy-based-challenge-rule-list-rules-spec-path) below.

&#x2022; `query_params` - Optional Block<br>A list of predicates for all query parameters that need to be matched. The criteria for matching each query parameter are described in individual instances of QueryParameterMatcherType. The actual query parameter values are extracted from the request API as a list of strings for each query parameter name. Note that all specified query parameter predicates must evaluate to true<br>See [Query Params](#policy-based-challenge-rule-list-rules-spec-query-params) below.

&#x2022; `tls_fingerprint_matcher` - Optional Block<br>TLS Fingerprint Matcher. A TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of supported positve match criteria includes a list of known classes of TLS fingerprints and a list of exact values. The match is considered successful if either of these positive criteria are satisfied and the input fingerprint is not one of the excluded values<br>See [TLS Fingerprint Matcher](#policy-based-challenge-rule-list-rules-spec-tls-fingerprint-matcher) below.

<a id="policy-based-challenge-rule-list-rules-spec-arg-matchers"></a>

**Policy Based Challenge Rule List Rules Spec Arg Matchers**

&#x2022; `check_not_present` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `check_present` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `invert_matcher` - Optional Bool<br>Invert Matcher. Invert Match of the expression defined

&#x2022; `item` - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#policy-based-challenge-rule-list-rules-spec-arg-matchers-item) below.

&#x2022; `name` - Optional String<br>Argument Name. x-example: 'phones[_]' x-example: 'cars.make.toyota.models[1]' x-example: 'cars.make.honda.models[_]' x-example: 'cars.make[_].models[_]' A case-sensitive JSON path in the HTTP request body

<a id="policy-based-challenge-rule-list-rules-spec-arg-matchers-item"></a>

**Policy Based Challenge Rule List Rules Spec Arg Matchers Item**

&#x2022; `exact_values` - Optional List<br>Exact Values. A list of exact values to match the input against

&#x2022; `regex_values` - Optional List<br>Regex Values. A list of regular expressions to match the input against

&#x2022; `transformers` - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

<a id="policy-based-challenge-rule-list-rules-spec-asn-list"></a>

**Policy Based Challenge Rule List Rules Spec Asn List**

&#x2022; `as_numbers` - Optional List<br>AS Numbers. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer

<a id="policy-based-challenge-rule-list-rules-spec-asn-matcher"></a>

**Policy Based Challenge Rule List Rules Spec Asn Matcher**

&#x2022; `asn_sets` - Optional Block<br>BGP ASN Sets. A list of references to bgp_asn_set objects<br>See [Asn Sets](#policy-based-challenge-rule-list-rules-spec-asn-matcher-asn-sets) below.

<a id="policy-based-challenge-rule-list-rules-spec-asn-matcher-asn-sets"></a>

**Policy Based Challenge Rule List Rules Spec Asn Matcher Asn Sets**

&#x2022; `kind` - Optional String<br>Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route')

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

&#x2022; `uid` - Optional String<br>UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid

<a id="policy-based-challenge-rule-list-rules-spec-body-matcher"></a>

**Policy Based Challenge Rule List Rules Spec Body Matcher**

&#x2022; `exact_values` - Optional List<br>Exact Values. A list of exact values to match the input against

&#x2022; `regex_values` - Optional List<br>Regex Values. A list of regular expressions to match the input against

&#x2022; `transformers` - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

<a id="policy-based-challenge-rule-list-rules-spec-client-selector"></a>

**Policy Based Challenge Rule List Rules Spec Client Selector**

&#x2022; `expressions` - Optional List<br>Selector Expression. expressions contains the kubernetes style label expression for selections

<a id="policy-based-challenge-rule-list-rules-spec-cookie-matchers"></a>

**Policy Based Challenge Rule List Rules Spec Cookie Matchers**

&#x2022; `check_not_present` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `check_present` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `invert_matcher` - Optional Bool<br>Invert Matcher. Invert Match of the expression defined

&#x2022; `item` - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#policy-based-challenge-rule-list-rules-spec-cookie-matchers-item) below.

&#x2022; `name` - Optional String<br>Cookie Name. A case-sensitive cookie name

<a id="policy-based-challenge-rule-list-rules-spec-cookie-matchers-item"></a>

**Policy Based Challenge Rule List Rules Spec Cookie Matchers Item**

&#x2022; `exact_values` - Optional List<br>Exact Values. A list of exact values to match the input against

&#x2022; `regex_values` - Optional List<br>Regex Values. A list of regular expressions to match the input against

&#x2022; `transformers` - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

<a id="policy-based-challenge-rule-list-rules-spec-domain-matcher"></a>

**Policy Based Challenge Rule List Rules Spec Domain Matcher**

&#x2022; `exact_values` - Optional List<br>Exact Values. A list of exact values to match the input against

&#x2022; `regex_values` - Optional List<br>Regex Values. A list of regular expressions to match the input against

<a id="policy-based-challenge-rule-list-rules-spec-headers"></a>

**Policy Based Challenge Rule List Rules Spec Headers**

&#x2022; `check_not_present` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `check_present` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `invert_matcher` - Optional Bool<br>Invert Header Matcher. Invert the match result

&#x2022; `item` - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#policy-based-challenge-rule-list-rules-spec-headers-item) below.

&#x2022; `name` - Optional String<br>Header Name. A case-insensitive HTTP header name

<a id="policy-based-challenge-rule-list-rules-spec-headers-item"></a>

**Policy Based Challenge Rule List Rules Spec Headers Item**

&#x2022; `exact_values` - Optional List<br>Exact Values. A list of exact values to match the input against

&#x2022; `regex_values` - Optional List<br>Regex Values. A list of regular expressions to match the input against

&#x2022; `transformers` - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

<a id="policy-based-challenge-rule-list-rules-spec-http-method"></a>

**Policy Based Challenge Rule List Rules Spec HTTP Method**

&#x2022; `invert_matcher` - Optional Bool<br>Invert Method Matcher. Invert the match result

&#x2022; `methods` - Optional List  Defaults to `ANY`<br>Possible values are `ANY`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`, `COPY`<br>Method List. List of methods values to match against

<a id="policy-based-challenge-rule-list-rules-spec-ip-matcher"></a>

**Policy Based Challenge Rule List Rules Spec IP Matcher**

&#x2022; `invert_matcher` - Optional Bool<br>Invert IP Matcher. Invert the match result

&#x2022; `prefix_sets` - Optional Block<br>IP Prefix Sets. A list of references to ip_prefix_set objects<br>See [Prefix Sets](#policy-based-challenge-rule-list-rules-spec-ip-matcher-prefix-sets) below.

<a id="policy-based-challenge-rule-list-rules-spec-ip-matcher-prefix-sets"></a>

**Policy Based Challenge Rule List Rules Spec IP Matcher Prefix Sets**

&#x2022; `kind` - Optional String<br>Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route')

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

&#x2022; `uid` - Optional String<br>UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid

<a id="policy-based-challenge-rule-list-rules-spec-ip-prefix-list"></a>

**Policy Based Challenge Rule List Rules Spec IP Prefix List**

&#x2022; `invert_match` - Optional Bool<br>Invert Match Result. Invert the match result

&#x2022; `ip_prefixes` - Optional List<br>IPv4 Prefix List. List of IPv4 prefix strings

<a id="policy-based-challenge-rule-list-rules-spec-path"></a>

**Policy Based Challenge Rule List Rules Spec Path**

&#x2022; `exact_values` - Optional List<br>Exact Values. A list of exact path values to match the input HTTP path against

&#x2022; `invert_matcher` - Optional Bool<br>Invert Path Matcher. Invert the match result

&#x2022; `prefix_values` - Optional List<br>Prefix Values. A list of path prefix values to match the input HTTP path against

&#x2022; `regex_values` - Optional List<br>Regex Values. A list of regular expressions to match the input HTTP path against

&#x2022; `suffix_values` - Optional List<br>Suffix Values. A list of path suffix values to match the input HTTP path against

&#x2022; `transformers` - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

<a id="policy-based-challenge-rule-list-rules-spec-query-params"></a>

**Policy Based Challenge Rule List Rules Spec Query Params**

&#x2022; `check_not_present` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `check_present` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `invert_matcher` - Optional Bool<br>Invert Query Parameter Matcher. Invert the match result

&#x2022; `item` - Optional Block<br>Matcher. A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions<br>See [Item](#policy-based-challenge-rule-list-rules-spec-query-params-item) below.

&#x2022; `key` - Optional String<br>Query Parameter Name. A case-sensitive HTTP query parameter name

<a id="policy-based-challenge-rule-list-rules-spec-query-params-item"></a>

**Policy Based Challenge Rule List Rules Spec Query Params Item**

&#x2022; `exact_values` - Optional List<br>Exact Values. A list of exact values to match the input against

&#x2022; `regex_values` - Optional List<br>Regex Values. A list of regular expressions to match the input against

&#x2022; `transformers` - Optional List  Defaults to `TRANSFORMER_NONE`<br>Possible values are `LOWER_CASE`, `UPPER_CASE`, `BASE64_DECODE`, `NORMALIZE_PATH`, `REMOVE_WHITESPACE`, `URL_DECODE`, `TRIM_LEFT`, `TRIM_RIGHT`, `TRIM`<br>Transformers. An ordered list of transformers (starting from index 0) to be applied to the path before matching

<a id="policy-based-challenge-rule-list-rules-spec-tls-fingerprint-matcher"></a>

**Policy Based Challenge Rule List Rules Spec TLS Fingerprint Matcher**

&#x2022; `classes` - Optional List  Defaults to `TLS_FINGERPRINT_NONE`<br>Possible values are `TLS_FINGERPRINT_NONE`, `ANY_MALICIOUS_FINGERPRINT`, `ADWARE`, `ADWIND`, `DRIDEX`, `GOOTKIT`, `GOZI`, `JBIFROST`, `QUAKBOT`, `RANSOMWARE`, `TROLDESH`, `TOFSEE`, `TORRENTLOCKER`, `TRICKBOT`<br>TLS fingerprint classes. A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against

&#x2022; `exact_values` - Optional List<br>Exact Values. A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against

&#x2022; `excluded_values` - Optional List<br>Excluded Values. A list of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can be used to skip known false positives when using one or more known TLS fingerprint classes in the enclosing matcher

<a id="policy-based-challenge-temporary-user-blocking"></a>

**Policy Based Challenge Temporary User Blocking**

&#x2022; `custom_page` - Optional String<br>Custom Message for Temporary Blocking. Custom message is of type `uri_ref`. Currently supported URL schemes is `string:///`. For `string:///` scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. 'Blocked.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '<p> Blocked </p>'. Base64 encoded string for this HTML is 'PHA+IFBsZWFzZSBXYWl0IDwvcD4='

<a id="protected-cookies"></a>

**Protected Cookies**

&#x2022; `add_httponly` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `add_secure` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `disable_tampering_protection` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `enable_tampering_protection` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `ignore_httponly` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `ignore_max_age` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `ignore_samesite` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `ignore_secure` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `max_age_value` - Optional Number<br>Add Max Age. Add max age attribute

&#x2022; `name` - Optional String<br>Cookie Name. Name of the Cookie

&#x2022; `samesite_lax` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `samesite_none` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `samesite_strict` - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="rate-limit"></a>

**Rate Limit**

&#x2022; `custom_ip_allowed_list` - Optional Block<br>Custom IP Allowed List. IP Allowed list using existing ip_prefix_set objects<br>See [Custom IP Allowed List](#rate-limit-custom-ip-allowed-list) below.

&#x2022; `ip_allowed_list` - Optional Block<br>IPv4 Prefix List. x-example: '192.168.20.0/24' List of IPv4 prefixes that represent an endpoint<br>See [IP Allowed List](#rate-limit-ip-allowed-list) below.

&#x2022; `no_ip_allowed_list` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `no_policies` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `policies` - Optional Block<br>Rate Limiter Policy List. List of rate limiter policies to be applied<br>See [Policies](#rate-limit-policies) below.

&#x2022; `rate_limiter` - Optional Block<br>Rate Limit Value. A tuple consisting of a rate limit period unit and the total number of allowed requests for that period<br>See [Rate Limiter](#rate-limit-rate-limiter) below.

<a id="rate-limit-custom-ip-allowed-list"></a>

**Rate Limit Custom IP Allowed List**

&#x2022; `rate_limiter_allowed_prefixes` - Optional Block<br>List of IP Prefix Sets. References to ip_prefix_set objects. Requests from source IP addresses that are covered by one of the allowed IP Prefixes are not subjected to rate limiting<br>See [Rate Limiter Allowed Prefixes](#rate-limit-custom-ip-allowed-list-rate-limiter-allowed-prefixes) below.

<a id="rate-limit-custom-ip-allowed-list-rate-limiter-allowed-prefixes"></a>

**Rate Limit Custom IP Allowed List Rate Limiter Allowed Prefixes**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="rate-limit-ip-allowed-list"></a>

**Rate Limit IP Allowed List**

&#x2022; `prefixes` - Optional List<br>IPv4 Prefix List. List of IPv4 prefixes that represent an endpoint

<a id="rate-limit-policies"></a>

**Rate Limit Policies**

&#x2022; `policies` - Optional Block<br>Rate Limiter Policies. Ordered list of rate limiter policies<br>See [Policies](#rate-limit-policies-policies) below.

<a id="rate-limit-policies-policies"></a>

**Rate Limit Policies Policies**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="rate-limit-rate-limiter"></a>

**Rate Limit Rate Limiter**

&#x2022; `action_block` - Optional Block<br>Rate Limit Block Action. Action where a user is blocked from making further requests after exceeding rate limit threshold<br>See [Action Block](#rate-limit-rate-limiter-action-block) below.

&#x2022; `burst_multiplier` - Optional Number<br>Burst Multiplier. The maximum burst of requests to accommodate, expressed as a multiple of the rate

&#x2022; `disabled` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `leaky_bucket` - Optional Block<br>Leaky Bucket Rate Limiter. Leaky-Bucket is the default rate limiter algorithm for F5

&#x2022; `period_multiplier` - Optional Number<br>Periods. This setting, combined with Per Period units, provides a duration

&#x2022; `token_bucket` - Optional Block<br>Token Bucket Rate Limiter. Token-Bucket is a rate limiter algorithm that is stricter with enforcing limits

&#x2022; `total_number` - Optional Number<br>Number Of Requests. The total number of allowed requests per rate-limiting period

&#x2022; `unit` - Optional String  Defaults to `SECOND`<br>Possible values are `SECOND`, `MINUTE`, `HOUR`<br>Rate Limit Period Unit. Unit for the period per which the rate limit is applied. - SECOND: Second Rate limit period unit is seconds - MINUTE: Minute Rate limit period unit is minutes - HOUR: Hour Rate limit period unit is hours - DAY: Day Rate limit period unit is days

<a id="rate-limit-rate-limiter-action-block"></a>

**Rate Limit Rate Limiter Action Block**

&#x2022; `hours` - Optional Block<br>Hours. Input Duration Hours<br>See [Hours](#rate-limit-rate-limiter-action-block-hours) below.

&#x2022; `minutes` - Optional Block<br>Minutes. Input Duration Minutes<br>See [Minutes](#rate-limit-rate-limiter-action-block-minutes) below.

&#x2022; `seconds` - Optional Block<br>Seconds. Input Duration Seconds<br>See [Seconds](#rate-limit-rate-limiter-action-block-seconds) below.

<a id="rate-limit-rate-limiter-action-block-hours"></a>

**Rate Limit Rate Limiter Action Block Hours**

&#x2022; `duration` - Optional Number<br>Duration

<a id="rate-limit-rate-limiter-action-block-minutes"></a>

**Rate Limit Rate Limiter Action Block Minutes**

&#x2022; `duration` - Optional Number<br>Duration

<a id="rate-limit-rate-limiter-action-block-seconds"></a>

**Rate Limit Rate Limiter Action Block Seconds**

&#x2022; `duration` - Optional Number<br>Duration

<a id="ring-hash"></a>

**Ring Hash**

&#x2022; `hash_policy` - Optional Block<br>Hash Policy. Specifies a list of hash policies to use for ring hash load balancing. Each hash policy is evaluated individually and the combined result is used to route the request<br>See [Hash Policy](#ring-hash-hash-policy) below.

<a id="ring-hash-hash-policy"></a>

**Ring Hash Hash Policy**

&#x2022; `cookie` - Optional Block<br>Hashing using Cookie. Two types of cookie affinity: 1. Passive. Takes a cookie that's present in the cookies header and hashes on its value. 2. Generated. Generates and sets a cookie with an expiration (TTL) on the first request from the client in its response to the client, based on the endpoint the request gets sent to. The client then presents this on the next and all subsequent requests. The hash of this is sufficient to ensure these requests get sent to the same endpoint. The cookie is generated by hashing the source and destination ports and addresses so that multiple independent HTTP2 streams on the same connection will independently receive the same cookie, even if they arrive simultaneously<br>See [Cookie](#ring-hash-hash-policy-cookie) below.

&#x2022; `header_name` - Optional String<br>Header Name. The name or key of the request header that will be used to obtain the hash key

&#x2022; `source_ip` - Optional Bool<br>Source IP. Hash based on source IP address

&#x2022; `terminal` - Optional Bool<br>Terminal. Specify if its a terminal policy

<a id="ring-hash-hash-policy-cookie"></a>

**Ring Hash Hash Policy Cookie**

&#x2022; `add_httponly` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `add_secure` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `ignore_httponly` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `ignore_samesite` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `ignore_secure` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `name` - Optional String<br>Name. The name of the cookie that will be used to obtain the hash key. If the cookie is not present and TTL below is not set, no hash will be produced

&#x2022; `path` - Optional String<br>Path. The name of the path for the cookie. If no path is specified here, no path will be set for the cookie

&#x2022; `samesite_lax` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `samesite_none` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `samesite_strict` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `ttl` - Optional Number<br>TTL. If specified, a cookie with the TTL will be generated if the cookie is not present. If the TTL is present and zero, the generated cookie will be a session cookie. TTL value is in milliseconds

<a id="routes"></a>

**Routes**

&#x2022; `custom_route_object` - Optional Block<br>Custom Route Object. A custom route uses a route object created outside of this view<br>See [Custom Route Object](#routes-custom-route-object) below.

&#x2022; `direct_response_route` - Optional Block<br>Direct Response Route. A direct response route matches on path, incoming header, incoming port and/or HTTP method and responds directly to the matching traffic<br>See [Direct Response Route](#routes-direct-response-route) below.

&#x2022; `redirect_route` - Optional Block<br>Redirect Route. A redirect route matches on path, incoming header, incoming port and/or HTTP method and redirects the matching traffic to a different URL<br>See [Redirect Route](#routes-redirect-route) below.

&#x2022; `simple_route` - Optional Block<br>Simple Route. A simple route matches on path, incoming header, incoming port and/or HTTP method and forwards the matching traffic to the associated pools<br>See [Simple Route](#routes-simple-route) below.

<a id="routes-custom-route-object"></a>

**Routes Custom Route Object**

&#x2022; `route_ref` - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Route Ref](#routes-custom-route-object-route-ref) below.

<a id="routes-custom-route-object-route-ref"></a>

**Routes Custom Route Object Route Ref**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="routes-direct-response-route"></a>

**Routes Direct Response Route**

&#x2022; `headers` - Optional Block<br>Headers. List of (key, value) headers<br>See [Headers](#routes-direct-response-route-headers) below.

&#x2022; `http_method` - Optional String  Defaults to `ANY`<br>Possible values are `ANY`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`, `COPY`<br>HTTP Method. Specifies the HTTP method used to access a resource. Any HTTP Method

&#x2022; `incoming_port` - Optional Block<br>Port to Match. Port match of the request can be a range or a specific port<br>See [Incoming Port](#routes-direct-response-route-incoming-port) below.

&#x2022; `path` - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#routes-direct-response-route-path) below.

&#x2022; `route_direct_response` - Optional Block<br>Direct Response. Send this direct response in case of route match action is direct response<br>See [Route Direct Response](#routes-direct-response-route-route-direct-response) below.

<a id="routes-direct-response-route-headers"></a>

**Routes Direct Response Route Headers**

&#x2022; `exact` - Optional String<br>Exact. Header value to match exactly

&#x2022; `invert_match` - Optional Bool<br>NOT of match. Invert the result of the match to detect missing header or non-matching value

&#x2022; `name` - Optional String<br>Name. Name of the header

&#x2022; `presence` - Optional Bool<br>Presence. If true, check for presence of header

&#x2022; `regex` - Optional String<br>Regex. Regex match of the header value in re2 format

<a id="routes-direct-response-route-incoming-port"></a>

**Routes Direct Response Route Incoming Port**

&#x2022; `no_port_match` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `port` - Optional Number<br>Port. Exact Port to match

&#x2022; `port_ranges` - Optional String<br>Port range. Port range to match

<a id="routes-direct-response-route-path"></a>

**Routes Direct Response Route Path**

&#x2022; `path` - Optional String<br>Exact. Exact path value to match

&#x2022; `prefix` - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

&#x2022; `regex` - Optional String<br>Regex. Regular expression of path match (e.g. the value .* will match on all paths)

<a id="routes-direct-response-route-route-direct-response"></a>

**Routes Direct Response Route Route Direct Response**

&#x2022; `response_body_encoded` - Optional String<br>Response Body. Response body to send. Currently supported URL schemes is string:/// for which message should be encoded in Base64 format. The message can be either plain text or HTML. E.g. '<p> Access Denied </p>'. Base64 encoded string URL for this is string:///PHA+IEFjY2VzcyBEZW5pZWQgPC9wPg==

&#x2022; `response_code` - Optional Number<br>Response Code. response code to send

<a id="routes-redirect-route"></a>

**Routes Redirect Route**

&#x2022; `headers` - Optional Block<br>Headers. List of (key, value) headers<br>See [Headers](#routes-redirect-route-headers) below.

&#x2022; `http_method` - Optional String  Defaults to `ANY`<br>Possible values are `ANY`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`, `COPY`<br>HTTP Method. Specifies the HTTP method used to access a resource. Any HTTP Method

&#x2022; `incoming_port` - Optional Block<br>Port to Match. Port match of the request can be a range or a specific port<br>See [Incoming Port](#routes-redirect-route-incoming-port) below.

&#x2022; `path` - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#routes-redirect-route-path) below.

&#x2022; `route_redirect` - Optional Block<br>Redirect. route redirect parameters when match action is redirect<br>See [Route Redirect](#routes-redirect-route-route-redirect) below.

<a id="routes-redirect-route-headers"></a>

**Routes Redirect Route Headers**

&#x2022; `exact` - Optional String<br>Exact. Header value to match exactly

&#x2022; `invert_match` - Optional Bool<br>NOT of match. Invert the result of the match to detect missing header or non-matching value

&#x2022; `name` - Optional String<br>Name. Name of the header

&#x2022; `presence` - Optional Bool<br>Presence. If true, check for presence of header

&#x2022; `regex` - Optional String<br>Regex. Regex match of the header value in re2 format

<a id="routes-redirect-route-incoming-port"></a>

**Routes Redirect Route Incoming Port**

&#x2022; `no_port_match` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `port` - Optional Number<br>Port. Exact Port to match

&#x2022; `port_ranges` - Optional String<br>Port range. Port range to match

<a id="routes-redirect-route-path"></a>

**Routes Redirect Route Path**

&#x2022; `path` - Optional String<br>Exact. Exact path value to match

&#x2022; `prefix` - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

&#x2022; `regex` - Optional String<br>Regex. Regular expression of path match (e.g. the value .* will match on all paths)

<a id="routes-redirect-route-route-redirect"></a>

**Routes Redirect Route Route Redirect**

&#x2022; `host_redirect` - Optional String<br>Host. swap host part of incoming URL in redirect URL

&#x2022; `path_redirect` - Optional String<br>Path. swap path part of incoming URL in redirect URL

&#x2022; `prefix_rewrite` - Optional String<br>Prefix Rewrite. In Redirect response, the matched prefix (or path) should be swapped with this value. This option allows redirect URLs be dynamically created based on the request

&#x2022; `proto_redirect` - Optional String<br>Protocol. swap protocol part of incoming URL in redirect URL The protocol can be swapped with either HTTP or HTTPS When incoming-proto option is specified, swapping of protocol is not done

&#x2022; `remove_all_params` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `replace_params` - Optional String<br>Replace All Parameters

&#x2022; `response_code` - Optional Number<br>Response Code. The HTTP status code to use in the redirect response

&#x2022; `retain_all_params` - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="routes-simple-route"></a>

**Routes Simple Route**

&#x2022; `advanced_options` - Optional Block<br>Advanced Route Options. Configure advanced options for route like path rewrite, hash policy, etc<br>See [Advanced Options](#routes-simple-route-advanced-options) below.

&#x2022; `auto_host_rewrite` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `disable_host_rewrite` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `headers` - Optional Block<br>Headers. List of (key, value) headers<br>See [Headers](#routes-simple-route-headers) below.

&#x2022; `host_rewrite` - Optional String<br>Host Rewrite Value. Host header will be swapped with this value

&#x2022; `http_method` - Optional String  Defaults to `ANY`<br>Possible values are `ANY`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`, `COPY`<br>HTTP Method. Specifies the HTTP method used to access a resource. Any HTTP Method

&#x2022; `incoming_port` - Optional Block<br>Port to Match. Port match of the request can be a range or a specific port<br>See [Incoming Port](#routes-simple-route-incoming-port) below.

&#x2022; `origin_pools` - Optional Block<br>Origin Pools. Origin Pools for this route<br>See [Origin Pools](#routes-simple-route-origin-pools) below.

&#x2022; `path` - Optional Block<br>Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match<br>See [Path](#routes-simple-route-path) below.

&#x2022; `query_params` - Optional Block<br>Query Parameters. Handling of incoming query parameters in simple route<br>See [Query Params](#routes-simple-route-query-params) below.

<a id="routes-simple-route-advanced-options"></a>

**Routes Simple Route Advanced Options**

&#x2022; `app_firewall` - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [App Firewall](#routes-simple-route-advanced-options-app-firewall) below.

&#x2022; `bot_defense_javascript_injection` - Optional Block<br>Bot Defense Javascript Injection Configuration for inline deployments. Bot Defense Javascript Injection Configuration for inline bot defense deployments<br>See [Bot Defense Javascript Injection](#routes-simple-route-advanced-options-bot-defense-javascript-injection) below.

&#x2022; `buffer_policy` - Optional Block<br>Buffer Configuration. Some upstream applications are not capable of handling streamed data. This config enables buffering the entire request before sending to upstream application. We can specify the maximum buffer size and buffer interval with this config. Buffering can be enabled and disabled at VirtualHost and Route levels Route level buffer configuration takes precedence<br>See [Buffer Policy](#routes-simple-route-advanced-options-buffer-policy) below.

&#x2022; `common_buffering` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `common_hash_policy` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `cors_policy` - Optional Block<br>CORS Policy. Cross-Origin Resource Sharing requests configuration specified at Virtual-host or Route level. Route level configuration takes precedence. An example of an Cross origin HTTP request GET /resources/public-data/ HTTP/1.1 Host: bar.other User-Agent: Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.5; en-US; rv:1.9.1b3pre) Gecko/20081130 Minefield/3.1b3pre Accept: text/HTML,application/xhtml+XML,application/XML;q=0.9,*/*;q=0.8 Accept-Language: en-us,en;q=0.5 Accept-Encoding: gzip,deflate Accept-Charset: ISO-8859-1,utf-8;q=0.7,*;q=0.7 Connection: keep-alive Referrer: `HTTP://foo.example/examples/access-control/simpleXSInvocation.HTML` Origin: `HTTP://foo.example` HTTP/1.1 200 OK Date: Mon, 01 Dec 2008 00:23:53 GMT Server: Apache/2.0.61 Access-Control-Allow-Origin: * Keep-Alive: timeout=2, max=100 Connection: Keep-Alive Transfer-Encoding: chunked Content-Type: application/XML An example for cross origin HTTP OPTIONS request with Access-Control-Request-* header OPTIONS /resources/post-here/ HTTP/1.1 Host: bar.other User-Agent: Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.5; en-US; rv:1.9.1b3pre) Gecko/20081130 Minefield/3.1b3pre Accept: text/HTML,application/xhtml+XML,application/XML;q=0.9,*/*;q=0.8 Accept-Language: en-us,en;q=0.5 Accept-Encoding: gzip,deflate Accept-Charset: ISO-8859-1,utf-8;q=0.7,*;q=0.7 Connection: keep-alive Origin: `HTTP://foo.example` Access-Control-Request-Method: POST Access-Control-Request-Headers: X-PINGOTHER, Content-Type HTTP/1.1 204 No Content Date: Mon, 01 Dec 2008 01:15:39 GMT Server: Apache/2.0.61 (Unix) Access-Control-Allow-Origin: `HTTP://foo.example` Access-Control-Allow-Methods: POST, GET, OPTIONS Access-Control-Allow-Headers: X-PINGOTHER, Content-Type Access-Control-Max-Age: 86400 Vary: Accept-Encoding, Origin Keep-Alive: timeout=2, max=100 Connection: Keep-Alive<br>See [CORS Policy](#routes-simple-route-advanced-options-cors-policy) below.

&#x2022; `csrf_policy` - Optional Block<br>CSRF Policy. To mitigate CSRF attack , the policy checks where a request is coming from to determine if the request's origin is the same as its detination.The policy relies on two pieces of information used in determining if a request originated from the same host. 1. The origin that caused the user agent to issue the request (source origin). 2. The origin that the request is going to (target origin). When the policy evaluating a request, it ensures both pieces of information are present and compare their values. If the source origin is missing or origins do not match the request is rejected. The exception to this being if the source-origin has been added to they policy as valid. Because CSRF attacks specifically target state-changing requests, the policy only acts on the HTTP requests that have state-changing method (PUT,POST, etc.)<br>See [CSRF Policy](#routes-simple-route-advanced-options-csrf-policy) below.

&#x2022; `default_retry_policy` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `disable_location_add` - Optional Bool<br>Disable Location Addition. disables append of x-volterra-location = <RE-site-name> at route level, if it is configured at virtual-host level. This configuration is ignored on CE sites

&#x2022; `disable_mirroring` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `disable_prefix_rewrite` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `disable_spdy` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `disable_waf` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `disable_web_socket_config` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `do_not_retract_cluster` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `enable_spdy` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `endpoint_subsets` - Optional Block<br>Origin Servers Subsets. Upstream origin pool may be configured to divide its origin servers into subsets based on metadata attached to the origin servers. Routes may then specify the metadata that a endpoint must match in order to be selected by the load balancer For origin servers which are discovered in K8S or Consul cluster, the label of the service is merged with endpoint's labels. In case of Consul, the label is derived from the 'Tag' field. For labels that are common between configured endpoint and discovered service, labels from discovered service takes precedence. List of key-value pairs that will be used as matching metadata. Only those origin servers of upstream origin pool which match this metadata will be selected for load balancing

&#x2022; `inherited_bot_defense_javascript_injection` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `inherited_waf` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `inherited_waf_exclusion` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `mirror_policy` - Optional Block<br>Mirror Policy. MirrorPolicy is used for shadowing traffic from one origin pool to another. The approach used is 'fire and forget', meaning it will not wait for the shadow origin pool to respond before returning the response from the primary origin pool. All normal statistics are collected for the shadow origin pool making this feature useful for testing and troubleshooting<br>See [Mirror Policy](#routes-simple-route-advanced-options-mirror-policy) below.

&#x2022; `no_retry_policy` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `prefix_rewrite` - Optional String<br>Enable Prefix Rewrite. prefix_rewrite indicates that during forwarding, the matched prefix (or path) should be swapped with its value. When using regex path matching, the entire path (not including the query string) will be swapped with this value

&#x2022; `priority` - Optional String  Defaults to `DEFAULT`<br>Possible values are `DEFAULT`, `HIGH`<br>Routing Priority. Priority routing for each request. Different connection pools are used based on the priority selected for the request. Also, circuit-breaker configuration at destination cluster is chosen based on selected priority. Default routing mechanism High-Priority routing mechanism

&#x2022; `regex_rewrite` - Optional Block<br>Regex Match Rewrite. RegexMatchRewrite describes how to match a string and then produce a new string using a regular expression and a substitution string<br>See [Regex Rewrite](#routes-simple-route-advanced-options-regex-rewrite) below.

&#x2022; `request_cookies_to_add` - Optional Block<br>Add Cookies in Cookie Header. Cookies are key-value pairs to be added to HTTP request being routed towards upstream. Cookies specified at this level are applied after cookies from matched Route are applied<br>See [Request Cookies To Add](#routes-simple-route-advanced-options-request-cookies-to-add) below.

&#x2022; `request_cookies_to_remove` - Optional List<br>Remove Cookies from Cookie Header. List of keys of Cookies to be removed from the HTTP request being sent towards upstream

&#x2022; `request_headers_to_add` - Optional Block<br>Add Request Headers. Headers are key-value pairs to be added to HTTP request being routed towards upstream<br>See [Request Headers To Add](#routes-simple-route-advanced-options-request-headers-to-add) below.

&#x2022; `request_headers_to_remove` - Optional List<br>Remove Request Headers. List of keys of Headers to be removed from the HTTP request being sent towards upstream

&#x2022; `response_cookies_to_add` - Optional Block<br>Add Set-Cookie Headers. Cookies are name-value pairs along with optional attribute parameters to be added to HTTP response being sent towards downstream. Cookies specified at this level are applied after cookies from matched Route are applied<br>See [Response Cookies To Add](#routes-simple-route-advanced-options-response-cookies-to-add) below.

&#x2022; `response_cookies_to_remove` - Optional List<br>Remove Cookies from Set-Cookie Headers. List of name of Cookies to be removed from the HTTP response being sent towards downstream. Entire set-cookie header will be removed

&#x2022; `response_headers_to_add` - Optional Block<br>Add Response Headers. Headers are key-value pairs to be added to HTTP response being sent towards downstream<br>See [Response Headers To Add](#routes-simple-route-advanced-options-response-headers-to-add) below.

&#x2022; `response_headers_to_remove` - Optional List<br>Remove Response Headers. List of keys of Headers to be removed from the HTTP response being sent towards downstream

&#x2022; `retract_cluster` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `retry_policy` - Optional Block<br>Retry Policy. Retry policy configuration for route destination<br>See [Retry Policy](#routes-simple-route-advanced-options-retry-policy) below.

&#x2022; `specific_hash_policy` - Optional Block<br>Hash Policy List. List of hash policy rules<br>See [Specific Hash Policy](#routes-simple-route-advanced-options-specific-hash-policy) below.

&#x2022; `timeout` - Optional Number<br>Timeout. The timeout for the route including all retries, in milliseconds. Should be set to a high value or 0 (infinite timeout) for server-side streaming

&#x2022; `waf_exclusion_policy` - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [WAF Exclusion Policy](#routes-simple-route-advanced-options-waf-exclusion-policy) below.

&#x2022; `web_socket_config` - Optional Block<br>WebSocket Configuration. Configuration to allow WebSocket Request headers of such upgrade looks like below 'connection', 'Upgrade' 'upgrade', 'WebSocket' With configuration to allow WebSocket upgrade, ADC will produce following response 'HTTP/1.1 101 Switching Protocols 'Upgrade': 'WebSocket' 'Connection': 'Upgrade'<br>See [Web Socket Config](#routes-simple-route-advanced-options-web-socket-config) below.

<a id="routes-simple-route-advanced-options-app-firewall"></a>

**Routes Simple Route Advanced Options App Firewall**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="routes-simple-route-advanced-options-bot-defense-javascript-injection"></a>

**Routes Simple Route Advanced Options Bot Defense Javascript Injection**

&#x2022; `javascript_location` - Optional String  Defaults to `AFTER_HEAD`<br>Possible values are `AFTER_HEAD`, `AFTER_TITLE_END`, `BEFORE_SCRIPT`<br>JavaScript Location. All inside networks. Insert JavaScript after <head> tag Insert JavaScript after </title> tag. Insert JavaScript before first <script> tag

&#x2022; `javascript_tags` - Optional Block<br>JavaScript Tags. Select Add item to configure your javascript tag. If adding both Bot Adv and Fraud, the Bot Javascript should be added first<br>See [Javascript Tags](#routes-simple-route-advanced-options-bot-defense-javascript-injection-javascript-tags) below.

<a id="routes-simple-route-advanced-options-bot-defense-javascript-injection-javascript-tags"></a>

**Routes Simple Route Advanced Options Bot Defense Javascript Injection Javascript Tags**

&#x2022; `javascript_url` - Optional String<br>URL. Please enter the full URL (include domain and path), or relative path

&#x2022; `tag_attributes` - Optional Block<br>Tag Attributes. Add the tag attributes you want to include in your Javascript tag<br>See [Tag Attributes](#routes-simple-route-advanced-options-bot-defense-javascript-injection-javascript-tags-tag-attributes) below.

<a id="routes-simple-route-advanced-options-bot-defense-javascript-injection-javascript-tags-tag-attributes"></a>

**Routes Simple Route Advanced Options Bot Defense Javascript Injection Javascript Tags Tag Attributes**

&#x2022; `javascript_tag` - Optional String  Defaults to `JS_ATTR_ID`<br>Possible values are `JS_ATTR_ID`, `JS_ATTR_CID`, `JS_ATTR_CN`, `JS_ATTR_API_DOMAIN`, `JS_ATTR_API_URL`, `JS_ATTR_API_PATH`, `JS_ATTR_ASYNC`, `JS_ATTR_DEFER`<br>Tag Attribute Name. Select from one of the predefined tag attributes

&#x2022; `tag_value` - Optional String<br>Value. Add the tag attribute value

<a id="routes-simple-route-advanced-options-buffer-policy"></a>

**Routes Simple Route Advanced Options Buffer Policy**

&#x2022; `disabled` - Optional Bool<br>Disable. Disable buffering for a particular route. This is useful when virtual-host has buffering, but we need to disable it on a specific route. The value of this field is ignored for virtual-host

&#x2022; `max_request_bytes` - Optional Number<br>Max Request Bytes. The maximum request size that the filter will buffer before the connection manager will stop buffering and return a RequestEntityTooLarge (413) response

<a id="routes-simple-route-advanced-options-cors-policy"></a>

**Routes Simple Route Advanced Options CORS Policy**

&#x2022; `allow_credentials` - Optional Bool<br>Allow Credentials. Specifies whether the resource allows credentials

&#x2022; `allow_headers` - Optional String<br>Allow Headers. Specifies the content for the access-control-allow-headers header

&#x2022; `allow_methods` - Optional String<br>Allow Methods. Specifies the content for the access-control-allow-methods header

&#x2022; `allow_origin` - Optional List<br>Allow Origin. Specifies the origins that will be allowed to do CORS requests. An origin is allowed if either allow_origin or allow_origin_regex match

&#x2022; `allow_origin_regex` - Optional List<br>Allow Origin Regex. Specifies regex patterns that match allowed origins. An origin is allowed if either allow_origin or allow_origin_regex match

&#x2022; `disabled` - Optional Bool<br>Disabled. Disable the CorsPolicy for a particular route. This is useful when virtual-host has CorsPolicy, but we need to disable it on a specific route. The value of this field is ignored for virtual-host

&#x2022; `expose_headers` - Optional String<br>Expose Headers. Specifies the content for the access-control-expose-headers header

&#x2022; `maximum_age` - Optional Number<br>Maximum Age. Specifies the content for the access-control-max-age header in seconds. This indicates the maximum number of seconds the results can be cached A value of -1 will disable caching. Maximum permitted value is 86400 seconds (24 hours)

<a id="routes-simple-route-advanced-options-csrf-policy"></a>

**Routes Simple Route Advanced Options CSRF Policy**

&#x2022; `all_load_balancer_domains` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `custom_domain_list` - Optional Block<br>Domain name list. List of domain names used for Host header matching<br>See [Custom Domain List](#routes-simple-route-advanced-options-csrf-policy-custom-domain-list) below.

&#x2022; `disabled` - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="routes-simple-route-advanced-options-csrf-policy-custom-domain-list"></a>

**Routes Simple Route Advanced Options CSRF Policy Custom Domain List**

&#x2022; `domains` - Optional List<br>Domain names. A list of domain names that will be matched to loadbalancer. These domains are not used for SNI match. Wildcard names are supported in the suffix or prefix form

<a id="routes-simple-route-advanced-options-mirror-policy"></a>

**Routes Simple Route Advanced Options Mirror Policy**

&#x2022; `origin_pool` - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Origin Pool](#routes-simple-route-advanced-options-mirror-policy-origin-pool) below.

&#x2022; `percent` - Optional Block<br>Fractional Percent. Fraction used where sampling percentages are needed. example sampled requests<br>See [Percent](#routes-simple-route-advanced-options-mirror-policy-percent) below.

<a id="routes-simple-route-advanced-options-mirror-policy-origin-pool"></a>

**Routes Simple Route Advanced Options Mirror Policy Origin Pool**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="routes-simple-route-advanced-options-mirror-policy-percent"></a>

**Routes Simple Route Advanced Options Mirror Policy Percent**

&#x2022; `denominator` - Optional String  Defaults to `HUNDRED`<br>Possible values are `HUNDRED`, `TEN_THOUSAND`, `MILLION`<br>Denominator. Denominator used in fraction where sampling percentages are needed. example sampled requests Use hundred as denominator Use ten thousand as denominator Use million as denominator

&#x2022; `numerator` - Optional Number<br>Numerator. sampled parts per denominator. If denominator was 10000, then value of 5 will be 5 in 10000

<a id="routes-simple-route-advanced-options-regex-rewrite"></a>

**Routes Simple Route Advanced Options Regex Rewrite**

&#x2022; `pattern` - Optional String<br>Pattern. The regular expression used to find portions of a string that should be replaced

&#x2022; `substitution` - Optional String<br>Substitution. The string that should be substituted into matching portions of the subject string during a substitution operation to produce a new string

<a id="routes-simple-route-advanced-options-request-cookies-to-add"></a>

**Routes Simple Route Advanced Options Request Cookies To Add**

&#x2022; `name` - Optional String<br>Name. Name of the cookie in Cookie header

&#x2022; `overwrite` - Optional Bool  Defaults to `do`<br>Overwrite. Should the value be overwritten? If true, the value is overwritten to existing values. not overwrite

&#x2022; `secret_value` - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Secret Value](#routes-simple-route-advanced-options-request-cookies-to-add-secret-value) below.

&#x2022; `value` - Optional String<br>Value. Value of the Cookie header

<a id="routes-simple-route-advanced-options-request-cookies-to-add-secret-value"></a>

**Routes Simple Route Advanced Options Request Cookies To Add Secret Value**

&#x2022; `blindfold_secret_info` - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#routes-simple-route-advanced-options-request-cookies-to-add-secret-value-blindfold-secret-info) below.

&#x2022; `clear_secret_info` - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#routes-simple-route-advanced-options-request-cookies-to-add-secret-value-clear-secret-info) below.

<a id="routes-simple-route-advanced-options-request-cookies-to-add-secret-value-blindfold-secret-info"></a>

**Routes Simple Route Advanced Options Request Cookies To Add Secret Value Blindfold Secret Info**

&#x2022; `decryption_provider` - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

&#x2022; `location` - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

&#x2022; `store_provider` - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

<a id="routes-simple-route-advanced-options-request-cookies-to-add-secret-value-clear-secret-info"></a>

**Routes Simple Route Advanced Options Request Cookies To Add Secret Value Clear Secret Info**

&#x2022; `provider_ref` - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

&#x2022; `url` - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

<a id="routes-simple-route-advanced-options-request-headers-to-add"></a>

**Routes Simple Route Advanced Options Request Headers To Add**

&#x2022; `append` - Optional Bool  Defaults to `do`<br>Append. Should the value be appended? If true, the value is appended to existing values. not append

&#x2022; `name` - Optional String<br>Name. Name of the HTTP header

&#x2022; `secret_value` - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Secret Value](#routes-simple-route-advanced-options-request-headers-to-add-secret-value) below.

&#x2022; `value` - Optional String<br>Value. Value of the HTTP header

<a id="routes-simple-route-advanced-options-request-headers-to-add-secret-value"></a>

**Routes Simple Route Advanced Options Request Headers To Add Secret Value**

&#x2022; `blindfold_secret_info` - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#routes-simple-route-advanced-options-request-headers-to-add-secret-value-blindfold-secret-info) below.

&#x2022; `clear_secret_info` - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#routes-simple-route-advanced-options-request-headers-to-add-secret-value-clear-secret-info) below.

<a id="routes-simple-route-advanced-options-request-headers-to-add-secret-value-blindfold-secret-info"></a>

**Routes Simple Route Advanced Options Request Headers To Add Secret Value Blindfold Secret Info**

&#x2022; `decryption_provider` - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

&#x2022; `location` - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

&#x2022; `store_provider` - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

<a id="routes-simple-route-advanced-options-request-headers-to-add-secret-value-clear-secret-info"></a>

**Routes Simple Route Advanced Options Request Headers To Add Secret Value Clear Secret Info**

&#x2022; `provider_ref` - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

&#x2022; `url` - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

<a id="routes-simple-route-advanced-options-response-cookies-to-add"></a>

**Routes Simple Route Advanced Options Response Cookies To Add**

&#x2022; `add_domain` - Optional String<br>Add Domain. Add domain attribute

&#x2022; `add_expiry` - Optional String<br>Add expiry. Add expiry attribute

&#x2022; `add_httponly` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `add_partitioned` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `add_path` - Optional String<br>Add path. Add path attribute

&#x2022; `add_secure` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `ignore_domain` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `ignore_expiry` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `ignore_httponly` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `ignore_max_age` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `ignore_partitioned` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `ignore_path` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `ignore_samesite` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `ignore_secure` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `ignore_value` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `max_age_value` - Optional Number<br>Add Max Age. Add max age attribute

&#x2022; `name` - Optional String<br>Name. Name of the cookie in Cookie header

&#x2022; `overwrite` - Optional Bool  Defaults to `do`<br>Overwrite. Should the value be overwritten? If true, the value is overwritten to existing values. not overwrite

&#x2022; `samesite_lax` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `samesite_none` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `samesite_strict` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `secret_value` - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Secret Value](#routes-simple-route-advanced-options-response-cookies-to-add-secret-value) below.

&#x2022; `value` - Optional String<br>Value. Value of the Cookie header

<a id="routes-simple-route-advanced-options-response-cookies-to-add-secret-value"></a>

**Routes Simple Route Advanced Options Response Cookies To Add Secret Value**

&#x2022; `blindfold_secret_info` - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#routes-simple-route-advanced-options-response-cookies-to-add-secret-value-blindfold-secret-info) below.

&#x2022; `clear_secret_info` - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#routes-simple-route-advanced-options-response-cookies-to-add-secret-value-clear-secret-info) below.

<a id="routes-simple-route-advanced-options-response-cookies-to-add-secret-value-blindfold-secret-info"></a>

**Routes Simple Route Advanced Options Response Cookies To Add Secret Value Blindfold Secret Info**

&#x2022; `decryption_provider` - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

&#x2022; `location` - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

&#x2022; `store_provider` - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

<a id="routes-simple-route-advanced-options-response-cookies-to-add-secret-value-clear-secret-info"></a>

**Routes Simple Route Advanced Options Response Cookies To Add Secret Value Clear Secret Info**

&#x2022; `provider_ref` - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

&#x2022; `url` - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

<a id="routes-simple-route-advanced-options-response-headers-to-add"></a>

**Routes Simple Route Advanced Options Response Headers To Add**

&#x2022; `append` - Optional Bool  Defaults to `do`<br>Append. Should the value be appended? If true, the value is appended to existing values. not append

&#x2022; `name` - Optional String<br>Name. Name of the HTTP header

&#x2022; `secret_value` - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Secret Value](#routes-simple-route-advanced-options-response-headers-to-add-secret-value) below.

&#x2022; `value` - Optional String<br>Value. Value of the HTTP header

<a id="routes-simple-route-advanced-options-response-headers-to-add-secret-value"></a>

**Routes Simple Route Advanced Options Response Headers To Add Secret Value**

&#x2022; `blindfold_secret_info` - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#routes-simple-route-advanced-options-response-headers-to-add-secret-value-blindfold-secret-info) below.

&#x2022; `clear_secret_info` - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#routes-simple-route-advanced-options-response-headers-to-add-secret-value-clear-secret-info) below.

<a id="routes-simple-route-advanced-options-response-headers-to-add-secret-value-blindfold-secret-info"></a>

**Routes Simple Route Advanced Options Response Headers To Add Secret Value Blindfold Secret Info**

&#x2022; `decryption_provider` - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

&#x2022; `location` - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

&#x2022; `store_provider` - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

<a id="routes-simple-route-advanced-options-response-headers-to-add-secret-value-clear-secret-info"></a>

**Routes Simple Route Advanced Options Response Headers To Add Secret Value Clear Secret Info**

&#x2022; `provider_ref` - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

&#x2022; `url` - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

<a id="routes-simple-route-advanced-options-retry-policy"></a>

**Routes Simple Route Advanced Options Retry Policy**

&#x2022; `back_off` - Optional Block<br>Retry BackOff Interval. Specifies parameters that control retry back off<br>See [Back Off](#routes-simple-route-advanced-options-retry-policy-back-off) below.

&#x2022; `num_retries` - Optional Number  Defaults to `1`<br>Number of Retries. Specifies the allowed number of retries. Retries can be done any number of times. An exponential back-off algorithm is used between each retry

&#x2022; `per_try_timeout` - Optional Number<br>Per Try Timeout. Specifies a non-zero timeout per retry attempt. In milliseconds

&#x2022; `retriable_status_codes` - Optional List<br>Status Code to Retry. HTTP status codes that should trigger a retry in addition to those specified by retry_on

&#x2022; `retry_condition` - Optional List<br>Retry Condition. Specifies the conditions under which retry takes place. Retries can be on different types of condition depending on application requirements. For example, network failure, all 5xx response codes, idempotent 4xx response codes, etc The possible values are '5xx' : Retry will be done if the upstream server responds with any 5xx response code, or does not respond at all (disconnect/reset/read timeout). 'gateway-error' : Retry will be done only if the upstream server responds with 502, 503 or 504 responses (Included in 5xx) 'connect-failure' : Retry will be done if the request fails because of a connection failure to the upstream server (connect timeout, etc.). (Included in 5xx) 'refused-stream' : Retry is done if the upstream server resets the stream with a REFUSED_STREAM error code (Included in 5xx) 'retriable-4xx' : Retry is done if the upstream server responds with a retriable 4xx response code. The only response code in this category is HTTP CONFLICT (409) 'retriable-status-codes' : Retry is done if the upstream server responds with any response code matching one defined in retriable_status_codes field 'reset' : Retry is done if the upstream server does not respond at all (disconnect/reset/read timeout.)

<a id="routes-simple-route-advanced-options-retry-policy-back-off"></a>

**Routes Simple Route Advanced Options Retry Policy Back Off**

&#x2022; `base_interval` - Optional Number<br>Base Retry Interval. Specifies the base interval between retries in milliseconds

&#x2022; `max_interval` - Optional Number  Defaults to `10`<br>Maximum Retry Interval. Specifies the maximum interval between retries in milliseconds. This parameter is optional, but must be greater than or equal to the base_interval if set. The times the base_interval

<a id="routes-simple-route-advanced-options-specific-hash-policy"></a>

**Routes Simple Route Advanced Options Specific Hash Policy**

&#x2022; `hash_policy` - Optional Block<br>Hash Policy. Specifies a list of hash policies to use for ring hash load balancing. Each hash policy is evaluated individually and the combined result is used to route the request<br>See [Hash Policy](#routes-simple-route-advanced-options-specific-hash-policy-hash-policy) below.

<a id="routes-simple-route-advanced-options-specific-hash-policy-hash-policy"></a>

**Routes Simple Route Advanced Options Specific Hash Policy Hash Policy**

&#x2022; `cookie` - Optional Block<br>Hashing using Cookie. Two types of cookie affinity: 1. Passive. Takes a cookie that's present in the cookies header and hashes on its value. 2. Generated. Generates and sets a cookie with an expiration (TTL) on the first request from the client in its response to the client, based on the endpoint the request gets sent to. The client then presents this on the next and all subsequent requests. The hash of this is sufficient to ensure these requests get sent to the same endpoint. The cookie is generated by hashing the source and destination ports and addresses so that multiple independent HTTP2 streams on the same connection will independently receive the same cookie, even if they arrive simultaneously<br>See [Cookie](#routes-simple-route-advanced-options-specific-hash-policy-hash-policy-cookie) below.

&#x2022; `header_name` - Optional String<br>Header Name. The name or key of the request header that will be used to obtain the hash key

&#x2022; `source_ip` - Optional Bool<br>Source IP. Hash based on source IP address

&#x2022; `terminal` - Optional Bool<br>Terminal. Specify if its a terminal policy

<a id="routes-simple-route-advanced-options-specific-hash-policy-hash-policy-cookie"></a>

**Routes Simple Route Advanced Options Specific Hash Policy Hash Policy Cookie**

&#x2022; `add_httponly` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `add_secure` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `ignore_httponly` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `ignore_samesite` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `ignore_secure` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `name` - Optional String<br>Name. The name of the cookie that will be used to obtain the hash key. If the cookie is not present and TTL below is not set, no hash will be produced

&#x2022; `path` - Optional String<br>Path. The name of the path for the cookie. If no path is specified here, no path will be set for the cookie

&#x2022; `samesite_lax` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `samesite_none` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `samesite_strict` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `ttl` - Optional Number<br>TTL. If specified, a cookie with the TTL will be generated if the cookie is not present. If the TTL is present and zero, the generated cookie will be a session cookie. TTL value is in milliseconds

<a id="routes-simple-route-advanced-options-waf-exclusion-policy"></a>

**Routes Simple Route Advanced Options WAF Exclusion Policy**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="routes-simple-route-advanced-options-web-socket-config"></a>

**Routes Simple Route Advanced Options Web Socket Config**

&#x2022; `use_websocket` - Optional Bool<br>Use WebSocket. Specifies that the HTTP client connection to this route is allowed to upgrade to a WebSocket connection

<a id="routes-simple-route-headers"></a>

**Routes Simple Route Headers**

&#x2022; `exact` - Optional String<br>Exact. Header value to match exactly

&#x2022; `invert_match` - Optional Bool<br>NOT of match. Invert the result of the match to detect missing header or non-matching value

&#x2022; `name` - Optional String<br>Name. Name of the header

&#x2022; `presence` - Optional Bool<br>Presence. If true, check for presence of header

&#x2022; `regex` - Optional String<br>Regex. Regex match of the header value in re2 format

<a id="routes-simple-route-incoming-port"></a>

**Routes Simple Route Incoming Port**

&#x2022; `no_port_match` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `port` - Optional Number<br>Port. Exact Port to match

&#x2022; `port_ranges` - Optional String<br>Port range. Port range to match

<a id="routes-simple-route-origin-pools"></a>

**Routes Simple Route Origin Pools**

&#x2022; `cluster` - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Cluster](#routes-simple-route-origin-pools-cluster) below.

&#x2022; `endpoint_subsets` - Optional Block<br>Origin Servers Subsets. Upstream origin pool may be configured to divide its origin servers into subsets based on metadata attached to the origin servers. Routes may then specify the metadata that a endpoint must match in order to be selected by the load balancer For origin servers which are discovered in K8S or Consul cluster, the label of the service is merged with endpoint's labels. In case of Consul, the label is derived from the 'Tag' field. For labels that are common between configured endpoint and discovered service, labels from discovered service takes precedence. List of key-value pairs that will be used as matching metadata. Only those origin servers of upstream origin pool which match this metadata will be selected for load balancing

&#x2022; `pool` - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Pool](#routes-simple-route-origin-pools-pool) below.

&#x2022; `priority` - Optional Number<br>Priority. Priority of this origin pool, valid only with multiple origin pools. Value of 0 will make the pool as lowest priority origin pool Priority of 1 means highest priority and is considered active. When active origin pool is not available, lower priority origin pools are made active as per the increasing priority

&#x2022; `weight` - Optional Number<br>Weight. Weight of this origin pool, valid only with multiple origin pool. Value of 0 will disable the pool

<a id="routes-simple-route-origin-pools-cluster"></a>

**Routes Simple Route Origin Pools Cluster**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="routes-simple-route-origin-pools-pool"></a>

**Routes Simple Route Origin Pools Pool**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="routes-simple-route-path"></a>

**Routes Simple Route Path**

&#x2022; `path` - Optional String<br>Exact. Exact path value to match

&#x2022; `prefix` - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

&#x2022; `regex` - Optional String<br>Regex. Regular expression of path match (e.g. the value .* will match on all paths)

<a id="routes-simple-route-query-params"></a>

**Routes Simple Route Query Params**

&#x2022; `remove_all_params` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `replace_params` - Optional String<br>Replace All Parameters

&#x2022; `retain_all_params` - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="sensitive-data-disclosure-rules"></a>

**Sensitive Data Disclosure Rules**

&#x2022; `sensitive_data_types_in_response` - Optional Block<br>Sensitive Data Exposure Rules. Sensitive Data Exposure Rules allows specifying rules to mask sensitive data fields in API responses<br>See [Sensitive Data Types In Response](#sensitive-data-disclosure-rules-sensitive-data-types-in-response) below.

<a id="sensitive-data-disclosure-rules-sensitive-data-types-in-response"></a>

**Sensitive Data Disclosure Rules Sensitive Data Types In Response**

&#x2022; `api_endpoint` - Optional Block<br>API Endpoint. This defines API endpoint<br>See [API Endpoint](#sensitive-data-disclosure-rules-sensitive-data-types-in-response-api-endpoint) below.

&#x2022; `body` - Optional Block<br>Body Section Masking Options. Options for HTTP Body Masking<br>See [Body](#sensitive-data-disclosure-rules-sensitive-data-types-in-response-body) below.

&#x2022; `mask` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `report` - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="sensitive-data-disclosure-rules-sensitive-data-types-in-response-api-endpoint"></a>

**Sensitive Data Disclosure Rules Sensitive Data Types In Response API Endpoint**

&#x2022; `methods` - Optional List  Defaults to `ANY`<br>Possible values are `ANY`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`, `COPY`<br>Methods. Methods to be matched

&#x2022; `path` - Optional String<br>Path. Path to be matched

<a id="sensitive-data-disclosure-rules-sensitive-data-types-in-response-body"></a>

**Sensitive Data Disclosure Rules Sensitive Data Types In Response Body**

&#x2022; `fields` - Optional List<br>Values. List of JSON Path field values. Use square brackets with an underscore [_] to indicate array elements (e.g., person.emails[_]). To reference JSON keys that contain spaces, enclose the entire path in double quotes. For example: 'person.first name'

<a id="sensitive-data-policy"></a>

**Sensitive Data Policy**

&#x2022; `sensitive_data_policy_ref` - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Sensitive Data Policy Ref](#sensitive-data-policy-sensitive-data-policy-ref) below.

<a id="sensitive-data-policy-sensitive-data-policy-ref"></a>

**Sensitive Data Policy Sensitive Data Policy Ref**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="single-lb-app"></a>

**Single LB App**

&#x2022; `disable_discovery` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `disable_malicious_user_detection` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `enable_discovery` - Optional Block<br>API Discovery Setting. Specifies the settings used for API discovery<br>See [Enable Discovery](#single-lb-app-enable-discovery) below.

&#x2022; `enable_malicious_user_detection` - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="single-lb-app-enable-discovery"></a>

**Single LB App Enable Discovery**

&#x2022; `api_crawler` - Optional Block<br>API Crawling. API Crawler message<br>See [API Crawler](#single-lb-app-enable-discovery-api-crawler) below.

&#x2022; `api_discovery_from_code_scan` - Optional Block<br>Select Code Base and Repositories. x-required<br>See [API Discovery From Code Scan](#single-lb-app-enable-discovery-api-discovery-from-code-scan) below.

&#x2022; `custom_api_auth_discovery` - Optional Block<br>API Discovery Advanced Settings. API Discovery Advanced settings<br>See [Custom API Auth Discovery](#single-lb-app-enable-discovery-custom-api-auth-discovery) below.

&#x2022; `default_api_auth_discovery` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `disable_learn_from_redirect_traffic` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `discovered_api_settings` - Optional Block<br>Discovered API Settings. x-example: '2' Configure Discovered API Settings<br>See [Discovered API Settings](#single-lb-app-enable-discovery-discovered-api-settings) below.

&#x2022; `enable_learn_from_redirect_traffic` - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="single-lb-app-enable-discovery-api-crawler"></a>

**Single LB App Enable Discovery API Crawler**

&#x2022; `api_crawler_config` - Optional Block<br>Crawler Configure<br>See [API Crawler Config](#single-lb-app-enable-discovery-api-crawler-api-crawler-config) below.

&#x2022; `disable_api_crawler` - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="single-lb-app-enable-discovery-api-crawler-api-crawler-config"></a>

**Single LB App Enable Discovery API Crawler API Crawler Config**

&#x2022; `domains` - Optional Block<br>Domains to Crawl. Enter domains and their credentials to allow authenticated API crawling. You can only include domains you own that are associated with this Load Balancer<br>See [Domains](#single-lb-app-enable-discovery-api-crawler-api-crawler-config-domains) below.

<a id="single-lb-app-enable-discovery-api-crawler-api-crawler-config-domains"></a>

**Single LB App Enable Discovery API Crawler API Crawler Config Domains**

&#x2022; `domain` - Optional String<br>Domain. Select the domain to execute API Crawling with given credentials

&#x2022; `simple_login` - Optional Block<br>Simple Login<br>See [Simple Login](#single-lb-app-enable-discovery-api-crawler-api-crawler-config-domains-simple-login) below.

<a id="single-lb-app-enable-discovery-api-crawler-api-crawler-config-domains-simple-login"></a>

**Single LB App Enable Discovery API Crawler API Crawler Config Domains Simple Login**

&#x2022; `password` - Optional Block<br>Secret. SecretType is used in an object to indicate a sensitive/confidential field<br>See [Password](#single-lb-app-enable-discovery-api-crawler-api-crawler-config-domains-simple-login-password) below.

&#x2022; `user` - Optional String<br>User. Enter the username to assign credentials for the selected domain to crawl

<a id="single-lb-app-enable-discovery-api-crawler-api-crawler-config-domains-simple-login-password"></a>

**Single LB App Enable Discovery API Crawler API Crawler Config Domains Simple Login Password**

&#x2022; `blindfold_secret_info` - Optional Block<br>Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management<br>See [Blindfold Secret Info](#single-lb-app-enable-discovery-api-crawler-api-crawler-config-domains-simple-login-password-blindfold-secret-info) below.

&#x2022; `clear_secret_info` - Optional Block<br>In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted<br>See [Clear Secret Info](#single-lb-app-enable-discovery-api-crawler-api-crawler-config-domains-simple-login-password-clear-secret-info) below.

<a id="single-lb-app-enable-discovery-api-crawler-api-crawler-config-domains-simple-login-password-blindfold-secret-info"></a>

**Single LB App Enable Discovery API Crawler API Crawler Config Domains Simple Login Password Blindfold Secret Info**

&#x2022; `decryption_provider` - Optional String<br>Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service

&#x2022; `location` - Optional String<br>Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location

&#x2022; `store_provider` - Optional String<br>Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

<a id="single-lb-app-enable-discovery-api-crawler-api-crawler-config-domains-simple-login-password-clear-secret-info"></a>

**Single LB App Enable Discovery API Crawler API Crawler Config Domains Simple Login Password Clear Secret Info**

&#x2022; `provider_ref` - Optional String<br>Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:///

&#x2022; `url` - Optional String<br>URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding

<a id="single-lb-app-enable-discovery-api-discovery-from-code-scan"></a>

**Single LB App Enable Discovery API Discovery From Code Scan**

&#x2022; `code_base_integrations` - Optional Block<br>Select Code Base Integrations<br>See [Code Base Integrations](#single-lb-app-enable-discovery-api-discovery-from-code-scan-code-base-integrations) below.

<a id="single-lb-app-enable-discovery-api-discovery-from-code-scan-code-base-integrations"></a>

**Single LB App Enable Discovery API Discovery From Code Scan Code Base Integrations**

&#x2022; `all_repos` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `code_base_integration` - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [Code Base Integration](#single-lb-app-enable-discovery-api-discovery-from-code-scan-code-base-integrations-code-base-integration) below.

&#x2022; `selected_repos` - Optional Block<br>API Code Repositories. Select which API repositories represent the LB applications<br>See [Selected Repos](#single-lb-app-enable-discovery-api-discovery-from-code-scan-code-base-integrations-selected-repos) below.

<a id="single-lb-app-enable-discovery-api-discovery-from-code-scan-code-base-integrations-code-base-integration"></a>

**Single LB App Enable Discovery API Discovery From Code Scan Code Base Integrations Code Base Integration**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="single-lb-app-enable-discovery-api-discovery-from-code-scan-code-base-integrations-selected-repos"></a>

**Single LB App Enable Discovery API Discovery From Code Scan Code Base Integrations Selected Repos**

&#x2022; `api_code_repo` - Optional List<br>API Code Repository. Code repository which contain API endpoints

<a id="single-lb-app-enable-discovery-custom-api-auth-discovery"></a>

**Single LB App Enable Discovery Custom API Auth Discovery**

&#x2022; `api_discovery_ref` - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [API Discovery Ref](#single-lb-app-enable-discovery-custom-api-auth-discovery-api-discovery-ref) below.

<a id="single-lb-app-enable-discovery-custom-api-auth-discovery-api-discovery-ref"></a>

**Single LB App Enable Discovery Custom API Auth Discovery API Discovery Ref**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="single-lb-app-enable-discovery-discovered-api-settings"></a>

**Single LB App Enable Discovery Discovered API Settings**

&#x2022; `purge_duration_for_inactive_discovered_apis` - Optional Number<br>Purge Duration for Inactive Discovered APIs from Traffic. Inactive discovered API will be deleted after configured duration

<a id="slow-ddos-mitigation"></a>

**Slow DDOS Mitigation**

&#x2022; `disable_request_timeout` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `request_headers_timeout` - Optional Number  Defaults to `10000`<br>Request Headers Timeout. The amount of time the client has to send only the headers on the request stream before the stream is cancelled. The milliseconds. This setting provides protection against Slowloris attacks

&#x2022; `request_timeout` - Optional Number<br>Custom Timeout

<a id="timeouts"></a>

**Timeouts**

&#x2022; `create` - Optional String<br>A string that can be [parsed as a duration](`HTTPS://pkg.go.dev/time#ParseDuration`) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours)

&#x2022; `delete` - Optional String<br>A string that can be [parsed as a duration](`HTTPS://pkg.go.dev/time#ParseDuration`) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are saved into state before the destroy operation occurs

&#x2022; `read` - Optional String<br>A string that can be [parsed as a duration](`HTTPS://pkg.go.dev/time#ParseDuration`) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh is enabled

&#x2022; `update` - Optional String<br>A string that can be [parsed as a duration](`HTTPS://pkg.go.dev/time#ParseDuration`) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours)

<a id="trusted-clients"></a>

**Trusted Clients**

&#x2022; `actions` - Optional List  Defaults to `SKIP_PROCESSING_WAF`<br>Possible values are `SKIP_PROCESSING_WAF`, `SKIP_PROCESSING_BOT`, `SKIP_PROCESSING_MUM`, `SKIP_PROCESSING_IP_REPUTATION`, `SKIP_PROCESSING_API_PROTECTION`, `SKIP_PROCESSING_OAS_VALIDATION`, `SKIP_PROCESSING_DDOS_PROTECTION`, `SKIP_PROCESSING_THREAT_MESH`, `SKIP_PROCESSING_MALWARE_PROTECTION`<br>Actions. Actions that should be taken when client identifier matches the rule

&#x2022; `as_number` - Optional Number<br>AS Number. RFC 6793 defined 4-byte AS number

&#x2022; `bot_skip_processing` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `expiration_timestamp` - Optional String<br>Expiration Timestamp. The expiration_timestamp is the RFC 3339 format timestamp at which the containing rule is considered to be logically expired. The rule continues to exist in the configuration but is not applied anymore

&#x2022; `http_header` - Optional Block<br>HTTP Header. Request header name and value pairs<br>See [HTTP Header](#trusted-clients-http-header) below.

&#x2022; `ip_prefix` - Optional String<br>IPv4 Prefix. IPv4 prefix string

&#x2022; `ipv6_prefix` - Optional String<br>IPv6 Prefix. IPv6 prefix string

&#x2022; `metadata` - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#trusted-clients-metadata) below.

&#x2022; `skip_processing` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `user_identifier` - Optional String<br>User Identifier. Identify user based on user identifier. User identifier value needs to be copied from security event

&#x2022; `waf_skip_processing` - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="trusted-clients-http-header"></a>

**Trusted Clients HTTP Header**

&#x2022; `headers` - Optional Block<br>HTTP Headers. List of HTTP header name and value pairs<br>See [Headers](#trusted-clients-http-header-headers) below.

<a id="trusted-clients-http-header-headers"></a>

**Trusted Clients HTTP Header Headers**

&#x2022; `exact` - Optional String<br>Exact. Header value to match exactly

&#x2022; `invert_match` - Optional Bool<br>NOT of match. Invert the result of the match to detect missing header or non-matching value

&#x2022; `name` - Optional String<br>Name. Name of the header

&#x2022; `presence` - Optional Bool<br>Presence. If true, check for presence of header

&#x2022; `regex` - Optional String<br>Regex. Regex match of the header value in re2 format

<a id="trusted-clients-metadata"></a>

**Trusted Clients Metadata**

&#x2022; `description` - Optional String<br>Description. Human readable description

&#x2022; `name` - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

<a id="user-identification"></a>

**User Identification**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

<a id="waf-exclusion"></a>

**WAF Exclusion**

&#x2022; `waf_exclusion_inline_rules` - Optional Block<br>WAF Exclusion Inline Rules. A list of WAF exclusion rules that will be applied inline<br>See [WAF Exclusion Inline Rules](#waf-exclusion-waf-exclusion-inline-rules) below.

&#x2022; `waf_exclusion_policy` - Optional Block<br>Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name<br>See [WAF Exclusion Policy](#waf-exclusion-waf-exclusion-policy) below.

<a id="waf-exclusion-waf-exclusion-inline-rules"></a>

**WAF Exclusion WAF Exclusion Inline Rules**

&#x2022; `rules` - Optional Block<br>WAF Exclusion Rules. An ordered list of WAF Exclusions specific to this Load Balancer<br>See [Rules](#waf-exclusion-waf-exclusion-inline-rules-rules) below.

<a id="waf-exclusion-waf-exclusion-inline-rules-rules"></a>

**WAF Exclusion WAF Exclusion Inline Rules Rules**

&#x2022; `any_domain` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `any_path` - Optional Block<br>Empty. This can be used for messages where no values are needed

&#x2022; `app_firewall_detection_control` - Optional Block<br>App Firewall Detection Control. Define the list of Signature IDs, Violations, Attack Types and Bot Names that should be excluded from triggering on the defined match criteria<br>See [App Firewall Detection Control](#waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control) below.

&#x2022; `exact_value` - Optional String<br>Exact Value. Exact domain name

&#x2022; `expiration_timestamp` - Optional String<br>Expiration Timestamp. The expiration_timestamp is the RFC 3339 format timestamp at which the containing rule is considered to be logically expired. The rule continues to exist in the configuration but is not applied anymore

&#x2022; `metadata` - Optional Block<br>Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs<br>See [Metadata](#waf-exclusion-waf-exclusion-inline-rules-rules-metadata) below.

&#x2022; `methods` - Optional List  Defaults to `ANY`<br>Possible values are `ANY`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`, `COPY`<br>Methods. methods to be matched

&#x2022; `path_prefix` - Optional String<br>Prefix. Path prefix to match (e.g. the value / will match on all paths)

&#x2022; `path_regex` - Optional String<br>Path Regex. Define the regex for the path. For example, the regex ^/.*$ will match on all paths

&#x2022; `suffix_value` - Optional String<br>Suffix Value. Suffix of domain name e.g 'xyz.com' will match '*.xyz.com' and 'xyz.com'

&#x2022; `waf_skip_processing` - Optional Block<br>Empty. This can be used for messages where no values are needed

<a id="waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control"></a>

**WAF Exclusion WAF Exclusion Inline Rules Rules App Firewall Detection Control**

&#x2022; `exclude_attack_type_contexts` - Optional Block<br>Attack Types. Attack Types to be excluded for the defined match criteria<br>See [Exclude Attack Type Contexts](#waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control-exclude-attack-type-contexts) below.

&#x2022; `exclude_bot_name_contexts` - Optional Block<br>Bot Names. Bot Names to be excluded for the defined match criteria<br>See [Exclude Bot Name Contexts](#waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control-exclude-bot-name-contexts) below.

&#x2022; `exclude_signature_contexts` - Optional Block<br>Signature IDs. Signature IDs to be excluded for the defined match criteria<br>See [Exclude Signature Contexts](#waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control-exclude-signature-contexts) below.

&#x2022; `exclude_violation_contexts` - Optional Block<br>Violations. Violations to be excluded for the defined match criteria<br>See [Exclude Violation Contexts](#waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control-exclude-violation-contexts) below.

<a id="waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control-exclude-attack-type-contexts"></a>

**WAF Exclusion WAF Exclusion Inline Rules Rules App Firewall Detection Control Exclude Attack Type Contexts**

&#x2022; `context` - Optional String  Defaults to `CONTEXT_ANY`<br>Possible values are `CONTEXT_ANY`, `CONTEXT_BODY`, `CONTEXT_REQUEST`, `CONTEXT_RESPONSE`, `CONTEXT_PARAMETER`, `CONTEXT_HEADER`, `CONTEXT_COOKIE`, `CONTEXT_URL`, `CONTEXT_URI`<br>WAF Exclusion Context Options. The available contexts for Exclusion rules. - CONTEXT_ANY: CONTEXT_ANY Detection will be excluded for all contexts. - CONTEXT_BODY: CONTEXT_BODY Detection will be excluded for the request body. - CONTEXT_REQUEST: CONTEXT_REQUEST Detection will be excluded for the request. - CONTEXT_RESPONSE: CONTEXT_RESPONSE - CONTEXT_PARAMETER: CONTEXT_PARAMETER Detection will be excluded for the parameters. The parameter name is required in the Context name field. If the field is left empty, the detection will be excluded for all parameters. - CONTEXT_HEADER: CONTEXT_HEADER Detection will be excluded for the headers. The header name is required in the Context name field. If the field is left empty, the detection will be excluded for all headers. - CONTEXT_COOKIE: CONTEXT_COOKIE Detection will be excluded for the cookies. The cookie name is required in the Context name field. If the field is left empty, the detection will be excluded for all cookies. - CONTEXT_URL: CONTEXT_URL Detection will be excluded for the request URL. - CONTEXT_URI: CONTEXT_URI

&#x2022; `context_name` - Optional String<br>Context Name. Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an wildcard asterisk (*)

&#x2022; `exclude_attack_type` - Optional String  Defaults to `ATTACK_TYPE_NONE`<br>Possible values are `ATTACK_TYPE_NONE`, `ATTACK_TYPE_NON_BROWSER_CLIENT`, `ATTACK_TYPE_OTHER_APPLICATION_ATTACKS`, `ATTACK_TYPE_TROJAN_BACKDOOR_SPYWARE`, `ATTACK_TYPE_DETECTION_EVASION`, `ATTACK_TYPE_VULNERABILITY_SCAN`, `ATTACK_TYPE_ABUSE_OF_FUNCTIONALITY`, `ATTACK_TYPE_AUTHENTICATION_AUTHORIZATION_ATTACKS`, `ATTACK_TYPE_BUFFER_OVERFLOW`, `ATTACK_TYPE_PREDICTABLE_RESOURCE_LOCATION`, `ATTACK_TYPE_INFORMATION_LEAKAGE`, `ATTACK_TYPE_DIRECTORY_INDEXING`, `ATTACK_TYPE_PATH_TRAVERSAL`, `ATTACK_TYPE_XPATH_INJECTION`, `ATTACK_TYPE_LDAP_INJECTION`, `ATTACK_TYPE_SERVER_SIDE_CODE_INJECTION`, `ATTACK_TYPE_COMMAND_EXECUTION`, `ATTACK_TYPE_SQL_INJECTION`, `ATTACK_TYPE_CROSS_SITE_SCRIPTING`, `ATTACK_TYPE_DENIAL_OF_SERVICE`, `ATTACK_TYPE_HTTP_PARSER_ATTACK`, `ATTACK_TYPE_SESSION_HIJACKING`, `ATTACK_TYPE_HTTP_RESPONSE_SPLITTING`, `ATTACK_TYPE_FORCEFUL_BROWSING`, `ATTACK_TYPE_REMOTE_FILE_INCLUDE`, `ATTACK_TYPE_MALICIOUS_FILE_UPLOAD`, `ATTACK_TYPE_GRAPHQL_PARSER_ATTACK`<br>Attack Types. List of all Attack Types ATTACK_TYPE_NONE ATTACK_TYPE_NON_BROWSER_CLIENT ATTACK_TYPE_OTHER_APPLICATION_ATTACKS ATTACK_TYPE_TROJAN_BACKDOOR_SPYWARE ATTACK_TYPE_DETECTION_EVASION ATTACK_TYPE_VULNERABILITY_SCAN ATTACK_TYPE_ABUSE_OF_FUNCTIONALITY ATTACK_TYPE_AUTHENTICATION_AUTHORIZATION_ATTACKS ATTACK_TYPE_BUFFER_OVERFLOW ATTACK_TYPE_PREDICTABLE_RESOURCE_LOCATION ATTACK_TYPE_INFORMATION_LEAKAGE ATTACK_TYPE_DIRECTORY_INDEXING ATTACK_TYPE_PATH_TRAVERSAL ATTACK_TYPE_XPATH_INJECTION ATTACK_TYPE_LDAP_INJECTION ATTACK_TYPE_SERVER_SIDE_CODE_INJECTION ATTACK_TYPE_COMMAND_EXECUTION ATTACK_TYPE_SQL_INJECTION ATTACK_TYPE_CROSS_SITE_SCRIPTING ATTACK_TYPE_DENIAL_OF_SERVICE ATTACK_TYPE_HTTP_PARSER_ATTACK ATTACK_TYPE_SESSION_HIJACKING ATTACK_TYPE_HTTP_RESPONSE_SPLITTING ATTACK_TYPE_FORCEFUL_BROWSING ATTACK_TYPE_REMOTE_FILE_INCLUDE ATTACK_TYPE_MALICIOUS_FILE_UPLOAD ATTACK_TYPE_GRAPHQL_PARSER_ATTACK

<a id="waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control-exclude-bot-name-contexts"></a>

**WAF Exclusion WAF Exclusion Inline Rules Rules App Firewall Detection Control Exclude Bot Name Contexts**

&#x2022; `bot_name` - Optional String<br>Bot Name

<a id="waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control-exclude-signature-contexts"></a>

**WAF Exclusion WAF Exclusion Inline Rules Rules App Firewall Detection Control Exclude Signature Contexts**

&#x2022; `context` - Optional String  Defaults to `CONTEXT_ANY`<br>Possible values are `CONTEXT_ANY`, `CONTEXT_BODY`, `CONTEXT_REQUEST`, `CONTEXT_RESPONSE`, `CONTEXT_PARAMETER`, `CONTEXT_HEADER`, `CONTEXT_COOKIE`, `CONTEXT_URL`, `CONTEXT_URI`<br>WAF Exclusion Context Options. The available contexts for Exclusion rules. - CONTEXT_ANY: CONTEXT_ANY Detection will be excluded for all contexts. - CONTEXT_BODY: CONTEXT_BODY Detection will be excluded for the request body. - CONTEXT_REQUEST: CONTEXT_REQUEST Detection will be excluded for the request. - CONTEXT_RESPONSE: CONTEXT_RESPONSE - CONTEXT_PARAMETER: CONTEXT_PARAMETER Detection will be excluded for the parameters. The parameter name is required in the Context name field. If the field is left empty, the detection will be excluded for all parameters. - CONTEXT_HEADER: CONTEXT_HEADER Detection will be excluded for the headers. The header name is required in the Context name field. If the field is left empty, the detection will be excluded for all headers. - CONTEXT_COOKIE: CONTEXT_COOKIE Detection will be excluded for the cookies. The cookie name is required in the Context name field. If the field is left empty, the detection will be excluded for all cookies. - CONTEXT_URL: CONTEXT_URL Detection will be excluded for the request URL. - CONTEXT_URI: CONTEXT_URI

&#x2022; `context_name` - Optional String<br>Context Name. Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an wildcard asterisk (*)

&#x2022; `signature_id` - Optional Number<br>SignatureID. The allowed values for signature id are 0 and in the range of 200000001-299999999. 0 implies that all signatures will be excluded for the specified context

<a id="waf-exclusion-waf-exclusion-inline-rules-rules-app-firewall-detection-control-exclude-violation-contexts"></a>

**WAF Exclusion WAF Exclusion Inline Rules Rules App Firewall Detection Control Exclude Violation Contexts**

&#x2022; `context` - Optional String  Defaults to `CONTEXT_ANY`<br>Possible values are `CONTEXT_ANY`, `CONTEXT_BODY`, `CONTEXT_REQUEST`, `CONTEXT_RESPONSE`, `CONTEXT_PARAMETER`, `CONTEXT_HEADER`, `CONTEXT_COOKIE`, `CONTEXT_URL`, `CONTEXT_URI`<br>WAF Exclusion Context Options. The available contexts for Exclusion rules. - CONTEXT_ANY: CONTEXT_ANY Detection will be excluded for all contexts. - CONTEXT_BODY: CONTEXT_BODY Detection will be excluded for the request body. - CONTEXT_REQUEST: CONTEXT_REQUEST Detection will be excluded for the request. - CONTEXT_RESPONSE: CONTEXT_RESPONSE - CONTEXT_PARAMETER: CONTEXT_PARAMETER Detection will be excluded for the parameters. The parameter name is required in the Context name field. If the field is left empty, the detection will be excluded for all parameters. - CONTEXT_HEADER: CONTEXT_HEADER Detection will be excluded for the headers. The header name is required in the Context name field. If the field is left empty, the detection will be excluded for all headers. - CONTEXT_COOKIE: CONTEXT_COOKIE Detection will be excluded for the cookies. The cookie name is required in the Context name field. If the field is left empty, the detection will be excluded for all cookies. - CONTEXT_URL: CONTEXT_URL Detection will be excluded for the request URL. - CONTEXT_URI: CONTEXT_URI

&#x2022; `context_name` - Optional String<br>Context Name. Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an wildcard asterisk (*)

&#x2022; `exclude_violation` - Optional String  Defaults to `VIOL_NONE`<br>Possible values are `VIOL_NONE`, `VIOL_FILETYPE`, `VIOL_METHOD`, `VIOL_MANDATORY_HEADER`, `VIOL_HTTP_RESPONSE_STATUS`, `VIOL_REQUEST_MAX_LENGTH`, `VIOL_FILE_UPLOAD`, `VIOL_FILE_UPLOAD_IN_BODY`, `VIOL_XML_MALFORMED`, `VIOL_JSON_MALFORMED`, `VIOL_ASM_COOKIE_MODIFIED`, `VIOL_HTTP_PROTOCOL_MULTIPLE_HOST_HEADERS`, `VIOL_HTTP_PROTOCOL_BAD_HOST_HEADER_VALUE`, `VIOL_HTTP_PROTOCOL_UNPARSABLE_REQUEST_CONTENT`, `VIOL_HTTP_PROTOCOL_NULL_IN_REQUEST`, `VIOL_HTTP_PROTOCOL_BAD_HTTP_VERSION`, `VIOL_HTTP_PROTOCOL_SEVERAL_CONTENT_LENGTH_HEADERS`, `VIOL_EVASION_DIRECTORY_TRAVERSALS`, `VIOL_MALFORMED_REQUEST`, `VIOL_EVASION_MULTIPLE_DECODING`, `VIOL_DATA_GUARD`, `VIOL_EVASION_APACHE_WHITESPACE`, `VIOL_COOKIE_MODIFIED`, `VIOL_EVASION_IIS_UNICODE_CODEPOINTS`, `VIOL_EVASION_IIS_BACKSLASHES`, `VIOL_EVASION_PERCENT_U_DECODING`, `VIOL_EVASION_BARE_BYTE_DECODING`, `VIOL_EVASION_BAD_UNESCAPE`, `VIOL_HTTP_PROTOCOL_BODY_IN_GET_OR_HEAD_REQUEST`, `VIOL_ENCODING`, `VIOL_COOKIE_MALFORMED`, `VIOL_GRAPHQL_FORMAT`, `VIOL_GRAPHQL_MALFORMED`, `VIOL_GRAPHQL_INTROSPECTION_QUERY`<br>App Firewall Violation Type. List of all supported Violation Types VIOL_NONE VIOL_FILETYPE VIOL_METHOD VIOL_MANDATORY_HEADER VIOL_HTTP_RESPONSE_STATUS VIOL_REQUEST_MAX_LENGTH VIOL_FILE_UPLOAD VIOL_FILE_UPLOAD_IN_BODY VIOL_XML_MALFORMED VIOL_JSON_MALFORMED VIOL_ASM_COOKIE_MODIFIED VIOL_HTTP_PROTOCOL_MULTIPLE_HOST_HEADERS VIOL_HTTP_PROTOCOL_BAD_HOST_HEADER_VALUE VIOL_HTTP_PROTOCOL_UNPARSABLE_REQUEST_CONTENT VIOL_HTTP_PROTOCOL_NULL_IN_REQUEST VIOL_HTTP_PROTOCOL_BAD_HTTP_VERSION VIOL_HTTP_PROTOCOL_CRLF_CHARACTERS_BEFORE_REQUEST_START VIOL_HTTP_PROTOCOL_NO_HOST_HEADER_IN_HTTP_1_1_REQUEST VIOL_HTTP_PROTOCOL_BAD_MULTIPART_PARAMETERS_PARSING VIOL_HTTP_PROTOCOL_SEVERAL_CONTENT_LENGTH_HEADERS VIOL_HTTP_PROTOCOL_CONTENT_LENGTH_SHOULD_BE_A_POSITIVE_NUMBER VIOL_EVASION_DIRECTORY_TRAVERSALS VIOL_MALFORMED_REQUEST VIOL_EVASION_MULTIPLE_DECODING VIOL_DATA_GUARD VIOL_EVASION_APACHE_WHITESPACE VIOL_COOKIE_MODIFIED VIOL_EVASION_IIS_UNICODE_CODEPOINTS VIOL_EVASION_IIS_BACKSLASHES VIOL_EVASION_PERCENT_U_DECODING VIOL_EVASION_BARE_BYTE_DECODING VIOL_EVASION_BAD_UNESCAPE VIOL_HTTP_PROTOCOL_BAD_MULTIPART_FORMDATA_REQUEST_PARSING VIOL_HTTP_PROTOCOL_BODY_IN_GET_OR_HEAD_REQUEST VIOL_HTTP_PROTOCOL_HIGH_ASCII_CHARACTERS_IN_HEADERS VIOL_ENCODING VIOL_COOKIE_MALFORMED VIOL_GRAPHQL_FORMAT VIOL_GRAPHQL_MALFORMED VIOL_GRAPHQL_INTROSPECTION_QUERY

<a id="waf-exclusion-waf-exclusion-inline-rules-rules-metadata"></a>

**WAF Exclusion WAF Exclusion Inline Rules Rules Metadata**

&#x2022; `description` - Optional String<br>Description. Human readable description

&#x2022; `name` - Optional String<br>Name. This is the name of the message. The value of name has to follow DNS-1035 format

<a id="waf-exclusion-waf-exclusion-policy"></a>

**WAF Exclusion WAF Exclusion Policy**

&#x2022; `name` - Optional String<br>Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name

&#x2022; `namespace` - Optional String<br>Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace

&#x2022; `tenant` - Optional String<br>Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant

## Import

Import is supported using the following syntax:

```shell
# Import using namespace/name format
terraform import f5xc_http_loadbalancer.example system/example
```
