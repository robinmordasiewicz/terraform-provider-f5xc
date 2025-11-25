---
page_title: "CDN Loadbalancer Nested Blocks - f5xc Provider"
subcategory: "Load Balancing"
description: |-
  Nested block reference for the CDN Loadbalancer resource.
---

# CDN Loadbalancer Nested Blocks

This page contains detailed documentation for nested blocks in the `f5xc_cdn_loadbalancer` resource.

For the main resource documentation, see [f5xc_cdn_loadbalancer](/docs/resources/cdn_loadbalancer).

## Contents

- [active-service-policies](#active-service-policies)
- [active-service-policies-policies](#active-service-policies-policies)
- [api-rate-limit](#api-rate-limit)
- [api-rate-limit-api-endpoint-rules](#api-rate-limit-api-endpoint-rules)
- [api-rate-limit-api-endpoint-rules-api-endpoint-method](#api-rate-limit-api-endpoint-rules-api-endpoint-method)
- [api-rate-limit-api-endpoint-rules-client-matcher](#api-rate-limit-api-endpoint-rules-client-matcher)
- [api-rate-limit-api-endpoint-rules-inline-rate-limiter](#api-rate-limit-api-endpoint-rules-inline-rate-limiter)
- [api-rate-limit-api-endpoint-rules-ref-rate-limiter](#api-rate-limit-api-endpoint-rules-ref-rate-limiter)
- [api-rate-limit-api-endpoint-rules-request-matcher](#api-rate-limit-api-endpoint-rules-request-matcher)
- [api-rate-limit-bypass-rate-limiting-rules](#api-rate-limit-bypass-rate-limiting-rules)
- [api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules)
- [api-rate-limit-custom-ip-allowed-list](#api-rate-limit-custom-ip-allowed-list)
- [api-rate-limit-custom-ip-allowed-list-rate-limiter-allowed-prefixes](#api-rate-limit-custom-ip-allowed-list-rate-limiter-allowed-prefixes)
- [api-rate-limit-ip-allowed-list](#api-rate-limit-ip-allowed-list)
- [api-rate-limit-server-url-rules](#api-rate-limit-server-url-rules)
- [api-rate-limit-server-url-rules-client-matcher](#api-rate-limit-server-url-rules-client-matcher)
- [api-rate-limit-server-url-rules-inline-rate-limiter](#api-rate-limit-server-url-rules-inline-rate-limiter)
- [api-rate-limit-server-url-rules-ref-rate-limiter](#api-rate-limit-server-url-rules-ref-rate-limiter)
- [api-rate-limit-server-url-rules-request-matcher](#api-rate-limit-server-url-rules-request-matcher)
- [api-specification](#api-specification)
- [api-specification-api-definition](#api-specification-api-definition)
- [api-specification-validation-all-spec-endpoints](#api-specification-validation-all-spec-endpoints)
- [api-specification-validation-all-spec-endpoints-fall-through-mode](#api-specification-validation-all-spec-endpoints-fall-through-mode)
- [api-specification-validation-all-spec-endpoints-settings](#api-specification-validation-all-spec-endpoints-settings)
- [api-specification-validation-all-spec-endpoints-validation-mode](#api-specification-validation-all-spec-endpoints-validation-mode)
- [api-specification-validation-custom-list](#api-specification-validation-custom-list)
- [api-specification-validation-custom-list-fall-through-mode](#api-specification-validation-custom-list-fall-through-mode)
- [api-specification-validation-custom-list-open-api-validation-rules](#api-specification-validation-custom-list-open-api-validation-rules)
- [api-specification-validation-custom-list-settings](#api-specification-validation-custom-list-settings)
- [app-firewall](#app-firewall)
- [blocked-clients](#blocked-clients)
- [blocked-clients-http-header](#blocked-clients-http-header)
- [blocked-clients-http-header-headers](#blocked-clients-http-header-headers)
- [blocked-clients-metadata](#blocked-clients-metadata)
- [bot-defense](#bot-defense)
- [bot-defense-policy](#bot-defense-policy)
- [bot-defense-policy-js-insert-all-pages](#bot-defense-policy-js-insert-all-pages)
- [bot-defense-policy-js-insert-all-pages-except](#bot-defense-policy-js-insert-all-pages-except)
- [bot-defense-policy-js-insertion-rules](#bot-defense-policy-js-insertion-rules)
- [bot-defense-policy-mobile-sdk-config](#bot-defense-policy-mobile-sdk-config)
- [bot-defense-policy-protected-app-endpoints](#bot-defense-policy-protected-app-endpoints)
- [captcha-challenge](#captcha-challenge)
- [client-side-defense](#client-side-defense)
- [client-side-defense-policy](#client-side-defense-policy)
- [client-side-defense-policy-js-insert-all-pages-except](#client-side-defense-policy-js-insert-all-pages-except)
- [client-side-defense-policy-js-insertion-rules](#client-side-defense-policy-js-insertion-rules)
- [cors-policy](#cors-policy)
- [csrf-policy](#csrf-policy)
- [csrf-policy-custom-domain-list](#csrf-policy-custom-domain-list)
- [custom-cache-rule](#custom-cache-rule)
- [custom-cache-rule-cdn-cache-rules](#custom-cache-rule-cdn-cache-rules)
- [data-guard-rules](#data-guard-rules)
- [data-guard-rules-metadata](#data-guard-rules-metadata)
- [data-guard-rules-path](#data-guard-rules-path)
- [ddos-mitigation-rules](#ddos-mitigation-rules)
- [ddos-mitigation-rules-ddos-client-source](#ddos-mitigation-rules-ddos-client-source)
- [ddos-mitigation-rules-ddos-client-source-asn-list](#ddos-mitigation-rules-ddos-client-source-asn-list)
- [ddos-mitigation-rules-ddos-client-source-ja4-tls-fingerprint-matcher](#ddos-mitigation-rules-ddos-client-source-ja4-tls-fingerprint-matcher)
- [ddos-mitigation-rules-ddos-client-source-tls-fingerprint-matcher](#ddos-mitigation-rules-ddos-client-source-tls-fingerprint-matcher)
- [ddos-mitigation-rules-ip-prefix-list](#ddos-mitigation-rules-ip-prefix-list)
- [ddos-mitigation-rules-metadata](#ddos-mitigation-rules-metadata)
- [default-cache-action](#default-cache-action)
- [enable-api-discovery](#enable-api-discovery)
- [enable-api-discovery-api-crawler](#enable-api-discovery-api-crawler)
- [enable-api-discovery-api-crawler-api-crawler-config](#enable-api-discovery-api-crawler-api-crawler-config)
- [enable-api-discovery-api-discovery-from-code-scan](#enable-api-discovery-api-discovery-from-code-scan)
- [enable-api-discovery-api-discovery-from-code-scan-code-base-integrations](#enable-api-discovery-api-discovery-from-code-scan-code-base-integrations)
- [enable-api-discovery-custom-api-auth-discovery](#enable-api-discovery-custom-api-auth-discovery)
- [enable-api-discovery-custom-api-auth-discovery-api-discovery-ref](#enable-api-discovery-custom-api-auth-discovery-api-discovery-ref)
- [enable-api-discovery-discovered-api-settings](#enable-api-discovery-discovered-api-settings)
- [enable-challenge](#enable-challenge)
- [enable-challenge-captcha-challenge-parameters](#enable-challenge-captcha-challenge-parameters)
- [enable-challenge-js-challenge-parameters](#enable-challenge-js-challenge-parameters)
- [enable-challenge-malicious-user-mitigation](#enable-challenge-malicious-user-mitigation)
- [enable-ip-reputation](#enable-ip-reputation)
- [graphql-rules](#graphql-rules)
- [graphql-rules-graphql-settings](#graphql-rules-graphql-settings)
- [graphql-rules-metadata](#graphql-rules-metadata)
- [http](#http)
- [https](#https)
- [https-tls-cert-options](#https-tls-cert-options)
- [https-tls-cert-options-tls-cert-params](#https-tls-cert-options-tls-cert-params)
- [https-tls-cert-options-tls-inline-params](#https-tls-cert-options-tls-inline-params)
- [https-auto-cert](#https-auto-cert)
- [https-auto-cert-tls-config](#https-auto-cert-tls-config)
- [js-challenge](#js-challenge)
- [jwt-validation](#jwt-validation)
- [jwt-validation-action](#jwt-validation-action)
- [jwt-validation-jwks-config](#jwt-validation-jwks-config)
- [jwt-validation-mandatory-claims](#jwt-validation-mandatory-claims)
- [jwt-validation-reserved-claims](#jwt-validation-reserved-claims)
- [jwt-validation-reserved-claims-audience](#jwt-validation-reserved-claims-audience)
- [jwt-validation-target](#jwt-validation-target)
- [jwt-validation-target-api-groups](#jwt-validation-target-api-groups)
- [jwt-validation-target-base-paths](#jwt-validation-target-base-paths)
- [jwt-validation-token-location](#jwt-validation-token-location)
- [l7-ddos-action-js-challenge](#l7-ddos-action-js-challenge)
- [origin-pool](#origin-pool)
- [origin-pool-more-origin-options](#origin-pool-more-origin-options)
- [origin-pool-origin-servers](#origin-pool-origin-servers)
- [origin-pool-origin-servers-public-ip](#origin-pool-origin-servers-public-ip)
- [origin-pool-origin-servers-public-name](#origin-pool-origin-servers-public-name)
- [origin-pool-public-name](#origin-pool-public-name)
- [origin-pool-use-tls](#origin-pool-use-tls)
- [origin-pool-use-tls-tls-config](#origin-pool-use-tls-tls-config)
- [origin-pool-use-tls-use-mtls](#origin-pool-use-tls-use-mtls)
- [origin-pool-use-tls-use-mtls-obj](#origin-pool-use-tls-use-mtls-obj)
- [origin-pool-use-tls-use-server-verification](#origin-pool-use-tls-use-server-verification)
- [other-settings](#other-settings)
- [other-settings-header-options](#other-settings-header-options)
- [other-settings-header-options-request-headers-to-add](#other-settings-header-options-request-headers-to-add)
- [other-settings-header-options-response-headers-to-add](#other-settings-header-options-response-headers-to-add)
- [other-settings-logging-options](#other-settings-logging-options)
- [other-settings-logging-options-client-log-options](#other-settings-logging-options-client-log-options)
- [other-settings-logging-options-origin-log-options](#other-settings-logging-options-origin-log-options)
- [policy-based-challenge](#policy-based-challenge)
- [policy-based-challenge-captcha-challenge-parameters](#policy-based-challenge-captcha-challenge-parameters)
- [policy-based-challenge-js-challenge-parameters](#policy-based-challenge-js-challenge-parameters)
- [policy-based-challenge-malicious-user-mitigation](#policy-based-challenge-malicious-user-mitigation)
- [policy-based-challenge-rule-list](#policy-based-challenge-rule-list)
- [policy-based-challenge-rule-list-rules](#policy-based-challenge-rule-list-rules)
- [policy-based-challenge-temporary-user-blocking](#policy-based-challenge-temporary-user-blocking)
- [protected-cookies](#protected-cookies)
- [rate-limit](#rate-limit)
- [rate-limit-custom-ip-allowed-list](#rate-limit-custom-ip-allowed-list)
- [rate-limit-custom-ip-allowed-list-rate-limiter-allowed-prefixes](#rate-limit-custom-ip-allowed-list-rate-limiter-allowed-prefixes)
- [rate-limit-ip-allowed-list](#rate-limit-ip-allowed-list)
- [rate-limit-policies](#rate-limit-policies)
- [rate-limit-policies-policies](#rate-limit-policies-policies)
- [rate-limit-rate-limiter](#rate-limit-rate-limiter)
- [rate-limit-rate-limiter-action-block](#rate-limit-rate-limiter-action-block)
- [sensitive-data-policy](#sensitive-data-policy)
- [sensitive-data-policy-sensitive-data-policy-ref](#sensitive-data-policy-sensitive-data-policy-ref)
- [slow-ddos-mitigation](#slow-ddos-mitigation)
- [timeouts](#timeouts)
- [trusted-clients](#trusted-clients)
- [trusted-clients-http-header](#trusted-clients-http-header)
- [trusted-clients-http-header-headers](#trusted-clients-http-header-headers)
- [trusted-clients-metadata](#trusted-clients-metadata)
- [user-identification](#user-identification)
- [waf-exclusion](#waf-exclusion)
- [waf-exclusion-waf-exclusion-inline-rules](#waf-exclusion-waf-exclusion-inline-rules)
- [waf-exclusion-waf-exclusion-inline-rules-rules](#waf-exclusion-waf-exclusion-inline-rules-rules)
- [waf-exclusion-waf-exclusion-policy](#waf-exclusion-waf-exclusion-policy)

---

<a id="active-service-policies"></a>

### Active Service Policies

`policies` - (Optional) Policies. Service Policies is a sequential engine where policies (and rules within the policy) are evaluated one after the other. It's important to define the correct order (policies evaluated from top to bottom in the list) for service policies, to get the intended result. For each request, its characteristics are evaluated based on the match criteria in each service policy starting at the top. If there is a match in the current policy, then the policy takes effect, and no more policies are evaluated. Otherwise, the next policy is evaluated. If all policies are evaluated and none match, then the request will be denied by default. See [Policies](#active-service-policies-policies) below.

<a id="active-service-policies-policies"></a>

### Active Service Policies Policies

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

<a id="api-rate-limit"></a>

### API Rate Limit

`api_endpoint_rules` - (Optional) API Endpoints. Sets of rules for a specific endpoints. Order is matter as it uses first match policy. For creating rule that contain a whole domain or group of endpoints, please use the server URL rules above. See [API Endpoint Rules](#api-rate-limit-api-endpoint-rules) below.

`bypass_rate_limiting_rules` - (Optional) Bypass Rate Limiting. This category defines rules per URL or API group. If request matches any of these rules, skip Rate Limiting. See [Bypass Rate Limiting Rules](#api-rate-limit-bypass-rate-limiting-rules) below.

`custom_ip_allowed_list` - (Optional) Custom IP Allowed List. IP Allowed list using existing ip_prefix_set objects. See [Custom IP Allowed List](#api-rate-limit-custom-ip-allowed-list) below.

`ip_allowed_list` - (Optional) IPv4 Prefix List. x-example: '192.168.20.0/24' List of IPv4 prefixes that represent an endpoint. See [IP Allowed List](#api-rate-limit-ip-allowed-list) below.

`no_ip_allowed_list` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`server_url_rules` - (Optional) Server URLs. Set of rules for entire domain or base path that contain multiple endpoints. Order is matter as it uses first match policy. For matching also specific endpoints you can use the API endpoint rules set bellow. See [Server URL Rules](#api-rate-limit-server-url-rules) below.

<a id="api-rate-limit-api-endpoint-rules"></a>

### API Rate Limit API Endpoint Rules

`any_domain` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`api_endpoint_method` - (Optional) HTTP Method Matcher. A HTTP method matcher specifies a list of methods to match an input HTTP method. The match is considered successful if the input method is a member of the list. The result of the match based on the method list is inverted if invert_matcher is true. See [API Endpoint Method](#api-rate-limit-api-endpoint-rules-api-endpoint-method) below.

`api_endpoint_path` - (Optional) API Endpoint. The endpoint (path) of the request (`String`).

`client_matcher` - (Optional) Client Matcher. Client conditions for matching a rule. See [Client Matcher](#api-rate-limit-api-endpoint-rules-client-matcher) below.

`inline_rate_limiter` - (Optional) InlineRateLimiter. See [Inline Rate Limiter](#api-rate-limit-api-endpoint-rules-inline-rate-limiter) below.

`ref_rate_limiter` - (Optional) Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name. See [Ref Rate Limiter](#api-rate-limit-api-endpoint-rules-ref-rate-limiter) below.

`request_matcher` - (Optional) Request Matcher. Request conditions for matching a rule. See [Request Matcher](#api-rate-limit-api-endpoint-rules-request-matcher) below.

`specific_domain` - (Optional) Specific Domain. The rule will apply for a specific domain (`String`).

<a id="api-rate-limit-api-endpoint-rules-api-endpoint-method"></a>

### API Rate Limit API Endpoint Rules API Endpoint Method

`invert_matcher` - (Optional) Invert Method Matcher. Invert the match result (`Bool`).

`methods` - (Optional) Method List. List of methods values to match against (`List`).

<a id="api-rate-limit-api-endpoint-rules-client-matcher"></a>

### API Rate Limit API Endpoint Rules Client Matcher

`any_client` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`any_ip` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`asn_list` - (Optional) ASN Match List. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer (`Block`).

`asn_matcher` - (Optional) ASN Matcher. Match any AS number contained in the list of bgp_asn_sets (`Block`).

`client_selector` - (Optional) Label Selector. This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expresssions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects. Label selector is immutable. expressions is a list of strings of label selection expression. Each string has ',' separated values which are 'AND' and all strings are logically 'OR'. BNF for expression string <selector-syntax> ::= <requirement> | <requirement> ',' <selector-syntax> <requirement> ::= [!] KEY [ <set-based-restriction> | <exact-match-restriction> ] <set-based-restriction> ::= '' | <inclusion-exclusion> <value-set> <inclusion-exclusion> ::= <inclusion> | <exclusion> <exclusion> ::= 'notin' <inclusion> ::= 'in' <value-set> ::= '(' <values> ')' <values> ::= VALUE | VALUE ',' <values> <exact-match-restriction> ::= ['='|'=='|'!='] VALUE (`Block`).

`ip_matcher` - (Optional) IP Prefix Matcher. Match any IP prefix contained in the list of ip_prefix_sets. The result of the match is inverted if invert_matcher is true (`Block`).

`ip_prefix_list` - (Optional) IP Prefix Match List. List of IP Prefix strings to match against (`Block`).

`ip_threat_category_list` - (Optional) IP Threat Category List Type. List of IP threat categories (`Block`).

`tls_fingerprint_matcher` - (Optional) TLS Fingerprint Matcher. A TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of supported positve match criteria includes a list of known classes of TLS fingerprints and a list of exact values. The match is considered successful if either of these positive criteria are satisfied and the input fingerprint is not one of the excluded values (`Block`).

<a id="api-rate-limit-api-endpoint-rules-inline-rate-limiter"></a>

### API Rate Limit API Endpoint Rules Inline Rate Limiter

`ref_user_id` - (Optional) Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name (`Block`).

`threshold` - (Optional) Threshold. The total number of allowed requests for 1 unit (e.g. SECOND/MINUTE/HOUR etc.) of the specified period (`Number`).

`unit` - (Optional) Rate Limit Period Unit. Unit for the period per which the rate limit is applied. - SECOND: Second Rate limit period unit is seconds - MINUTE: Minute Rate limit period unit is minutes - HOUR: Hour Rate limit period unit is hours - DAY: Day Rate limit period unit is days. Possible values are `SECOND`, `MINUTE`, `HOUR`. Defaults to `SECOND` (`String`).

`use_http_lb_user_id` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

<a id="api-rate-limit-api-endpoint-rules-ref-rate-limiter"></a>

### API Rate Limit API Endpoint Rules Ref Rate Limiter

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

<a id="api-rate-limit-api-endpoint-rules-request-matcher"></a>

### API Rate Limit API Endpoint Rules Request Matcher

`cookie_matchers` - (Optional) Cookie Matchers. A list of predicates for all cookies that need to be matched. The criteria for matching each cookie is described in individual instances of CookieMatcherType. The actual cookie values are extracted from the request API as a list of strings for each cookie name. Note that all specified cookie matcher predicates must evaluate to true (`Block`).

`headers` - (Optional) HTTP Headers. A list of predicates for various HTTP headers that need to match. The criteria for matching each HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values are extracted from the request API as a list of strings for each HTTP header type. Note that all specified header predicates must evaluate to true (`Block`).

`jwt_claims` - (Optional) JWT Claims. A list of predicates for various JWT claims that need to match. The criteria for matching each JWT claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are extracted from the JWT payload as a list of strings. Note that all specified JWT claim predicates must evaluate to true. Note that this feature only works on LBs with JWT Validation feature enabled (`Block`).

`query_params` - (Optional) HTTP Query Parameters. A list of predicates for all query parameters that need to be matched. The criteria for matching each query parameter are described in individual instances of QueryParameterMatcherType. The actual query parameter values are extracted from the request API as a list of strings for each query parameter name. Note that all specified query parameter predicates must evaluate to true (`Block`).

<a id="api-rate-limit-bypass-rate-limiting-rules"></a>

### API Rate Limit Bypass Rate Limiting Rules

`bypass_rate_limiting_rules` - (Optional) Bypass Rate Limiting. This category defines rules per URL or API group. If request matches any of these rules, skip Rate Limiting. See [Bypass Rate Limiting Rules](#api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules) below.

<a id="api-rate-limit-bypass-rate-limiting-rules-bypass-rate-limiting-rules"></a>

### API Rate Limit Bypass Rate Limiting Rules Bypass Rate Limiting Rules

`any_domain` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`any_url` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`api_endpoint` - (Optional) API Endpoint. This defines API endpoint (`Block`).

`api_groups` - (Optional) API Groups (`Block`).

`base_path` - (Optional) Base Path. The base path which this validation applies to (`String`).

`client_matcher` - (Optional) Client Matcher. Client conditions for matching a rule (`Block`).

`request_matcher` - (Optional) Request Matcher. Request conditions for matching a rule (`Block`).

`specific_domain` - (Optional) Specific Domain. The rule will apply for a specific domain. For example: API.example.com (`String`).

<a id="api-rate-limit-custom-ip-allowed-list"></a>

### API Rate Limit Custom IP Allowed List

`rate_limiter_allowed_prefixes` - (Optional) List of IP Prefix Sets. References to ip_prefix_set objects. Requests from source IP addresses that are covered by one of the allowed IP Prefixes are not subjected to rate limiting. See [Rate Limiter Allowed Prefixes](#api-rate-limit-custom-ip-allowed-list-rate-limiter-allowed-prefixes) below.

<a id="api-rate-limit-custom-ip-allowed-list-rate-limiter-allowed-prefixes"></a>

### API Rate Limit Custom IP Allowed List Rate Limiter Allowed Prefixes

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

<a id="api-rate-limit-ip-allowed-list"></a>

### API Rate Limit IP Allowed List

`prefixes` - (Optional) IPv4 Prefix List. List of IPv4 prefixes that represent an endpoint (`List`).

<a id="api-rate-limit-server-url-rules"></a>

### API Rate Limit Server URL Rules

`any_domain` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`api_group` - (Optional) API Group. API groups derived from API Definition swaggers. For example oas-all-operations including all paths and methods from the swaggers, oas-base-urls covering all requests under base-paths from the swaggers. Custom groups can be created if user tags paths or operations with 'x-volterra-API-group' extensions inside swaggers (`String`).

`base_path` - (Optional) Base Path. Prefix of the request path (`String`).

`client_matcher` - (Optional) Client Matcher. Client conditions for matching a rule. See [Client Matcher](#api-rate-limit-server-url-rules-client-matcher) below.

`inline_rate_limiter` - (Optional) InlineRateLimiter. See [Inline Rate Limiter](#api-rate-limit-server-url-rules-inline-rate-limiter) below.

`ref_rate_limiter` - (Optional) Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name. See [Ref Rate Limiter](#api-rate-limit-server-url-rules-ref-rate-limiter) below.

`request_matcher` - (Optional) Request Matcher. Request conditions for matching a rule. See [Request Matcher](#api-rate-limit-server-url-rules-request-matcher) below.

`specific_domain` - (Optional) Specific Domain. The rule will apply for a specific domain (`String`).

<a id="api-rate-limit-server-url-rules-client-matcher"></a>

### API Rate Limit Server URL Rules Client Matcher

`any_client` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`any_ip` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`asn_list` - (Optional) ASN Match List. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer (`Block`).

`asn_matcher` - (Optional) ASN Matcher. Match any AS number contained in the list of bgp_asn_sets (`Block`).

`client_selector` - (Optional) Label Selector. This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expresssions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects. Label selector is immutable. expressions is a list of strings of label selection expression. Each string has ',' separated values which are 'AND' and all strings are logically 'OR'. BNF for expression string <selector-syntax> ::= <requirement> | <requirement> ',' <selector-syntax> <requirement> ::= [!] KEY [ <set-based-restriction> | <exact-match-restriction> ] <set-based-restriction> ::= '' | <inclusion-exclusion> <value-set> <inclusion-exclusion> ::= <inclusion> | <exclusion> <exclusion> ::= 'notin' <inclusion> ::= 'in' <value-set> ::= '(' <values> ')' <values> ::= VALUE | VALUE ',' <values> <exact-match-restriction> ::= ['='|'=='|'!='] VALUE (`Block`).

`ip_matcher` - (Optional) IP Prefix Matcher. Match any IP prefix contained in the list of ip_prefix_sets. The result of the match is inverted if invert_matcher is true (`Block`).

`ip_prefix_list` - (Optional) IP Prefix Match List. List of IP Prefix strings to match against (`Block`).

`ip_threat_category_list` - (Optional) IP Threat Category List Type. List of IP threat categories (`Block`).

`tls_fingerprint_matcher` - (Optional) TLS Fingerprint Matcher. A TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of supported positve match criteria includes a list of known classes of TLS fingerprints and a list of exact values. The match is considered successful if either of these positive criteria are satisfied and the input fingerprint is not one of the excluded values (`Block`).

<a id="api-rate-limit-server-url-rules-inline-rate-limiter"></a>

### API Rate Limit Server URL Rules Inline Rate Limiter

`ref_user_id` - (Optional) Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name (`Block`).

`threshold` - (Optional) Threshold. The total number of allowed requests for 1 unit (e.g. SECOND/MINUTE/HOUR etc.) of the specified period (`Number`).

`unit` - (Optional) Rate Limit Period Unit. Unit for the period per which the rate limit is applied. - SECOND: Second Rate limit period unit is seconds - MINUTE: Minute Rate limit period unit is minutes - HOUR: Hour Rate limit period unit is hours - DAY: Day Rate limit period unit is days. Possible values are `SECOND`, `MINUTE`, `HOUR`. Defaults to `SECOND` (`String`).

`use_http_lb_user_id` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

<a id="api-rate-limit-server-url-rules-ref-rate-limiter"></a>

### API Rate Limit Server URL Rules Ref Rate Limiter

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

<a id="api-rate-limit-server-url-rules-request-matcher"></a>

### API Rate Limit Server URL Rules Request Matcher

`cookie_matchers` - (Optional) Cookie Matchers. A list of predicates for all cookies that need to be matched. The criteria for matching each cookie is described in individual instances of CookieMatcherType. The actual cookie values are extracted from the request API as a list of strings for each cookie name. Note that all specified cookie matcher predicates must evaluate to true (`Block`).

`headers` - (Optional) HTTP Headers. A list of predicates for various HTTP headers that need to match. The criteria for matching each HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values are extracted from the request API as a list of strings for each HTTP header type. Note that all specified header predicates must evaluate to true (`Block`).

`jwt_claims` - (Optional) JWT Claims. A list of predicates for various JWT claims that need to match. The criteria for matching each JWT claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are extracted from the JWT payload as a list of strings. Note that all specified JWT claim predicates must evaluate to true. Note that this feature only works on LBs with JWT Validation feature enabled (`Block`).

`query_params` - (Optional) HTTP Query Parameters. A list of predicates for all query parameters that need to be matched. The criteria for matching each query parameter are described in individual instances of QueryParameterMatcherType. The actual query parameter values are extracted from the request API as a list of strings for each query parameter name. Note that all specified query parameter predicates must evaluate to true (`Block`).

<a id="api-specification"></a>

### API Specification

`api_definition` - (Optional) Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name. See [API Definition](#api-specification-api-definition) below.

`validation_all_spec_endpoints` - (Optional) API Inventory. Settings for API Inventory validation. See [Validation All Spec Endpoints](#api-specification-validation-all-spec-endpoints) below.

`validation_custom_list` - (Optional) Custom List. Define API groups, base paths, or API endpoints and their OpenAPI validation modes. Any other API-endpoint not listed will act according to 'Fall Through Mode'. See [Validation Custom List](#api-specification-validation-custom-list) below.

`validation_disabled` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

<a id="api-specification-api-definition"></a>

### API Specification API Definition

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

<a id="api-specification-validation-all-spec-endpoints"></a>

### API Specification Validation All Spec Endpoints

`fall_through_mode` - (Optional) Fall Through Mode. x-required Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a. swagger) or doesn't have a specific rule in custom rules). See [Fall Through Mode](#api-specification-validation-all-spec-endpoints-fall-through-mode) below.

`settings` - (Optional) Common Settings. OpenAPI specification validation settings relevant for 'API Inventory' enforcement and for 'Custom list' enforcement. See [Settings](#api-specification-validation-all-spec-endpoints-settings) below.

`validation_mode` - (Optional) Validation Mode. x-required Validation mode of OpenAPI specification. When a validation mismatch occurs on a request to one of the endpoints listed on the OpenAPI specification file (a.k.a. swagger). See [Validation Mode](#api-specification-validation-all-spec-endpoints-validation-mode) below.

<a id="api-specification-validation-all-spec-endpoints-fall-through-mode"></a>

### API Specification Validation All Spec Endpoints Fall Through Mode

`fall_through_mode_allow` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`fall_through_mode_custom` - (Optional) Custom Fall Through Mode. Define the fall through settings (`Block`).

<a id="api-specification-validation-all-spec-endpoints-settings"></a>

### API Specification Validation All Spec Endpoints Settings

`oversized_body_fail_validation` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`oversized_body_skip_validation` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`property_validation_settings_custom` - (Optional) Validation Property Settings. Custom property validation settings (`Block`).

`property_validation_settings_default` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

<a id="api-specification-validation-all-spec-endpoints-validation-mode"></a>

### API Specification Validation All Spec Endpoints Validation Mode

`response_validation_mode_active` - (Optional) Open API Validation Mode Active. Validation mode properties of response (`Block`).

`skip_response_validation` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`skip_validation` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`validation_mode_active` - (Optional) Open API Validation Mode Active. Validation mode properties of request (`Block`).

<a id="api-specification-validation-custom-list"></a>

### API Specification Validation Custom List

`fall_through_mode` - (Optional) Fall Through Mode. x-required Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a. swagger) or doesn't have a specific rule in custom rules). See [Fall Through Mode](#api-specification-validation-custom-list-fall-through-mode) below.

`open_api_validation_rules` - (Optional) Validation List. See [Open API Validation Rules](#api-specification-validation-custom-list-open-api-validation-rules) below.

`settings` - (Optional) Common Settings. OpenAPI specification validation settings relevant for 'API Inventory' enforcement and for 'Custom list' enforcement. See [Settings](#api-specification-validation-custom-list-settings) below.

<a id="api-specification-validation-custom-list-fall-through-mode"></a>

### API Specification Validation Custom List Fall Through Mode

`fall_through_mode_allow` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`fall_through_mode_custom` - (Optional) Custom Fall Through Mode. Define the fall through settings (`Block`).

<a id="api-specification-validation-custom-list-open-api-validation-rules"></a>

### API Specification Validation Custom List Open API Validation Rules

`any_domain` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`api_endpoint` - (Optional) API Endpoint. This defines API endpoint (`Block`).

`api_group` - (Optional) API Group. The API group which this validation applies to (`String`).

`base_path` - (Optional) Base Path. The base path which this validation applies to (`String`).

`metadata` - (Optional) Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs (`Block`).

`specific_domain` - (Optional) Specific Domain. The rule will apply for a specific domain (`String`).

`validation_mode` - (Optional) Validation Mode. x-required Validation mode of OpenAPI specification. When a validation mismatch occurs on a request to one of the endpoints listed on the OpenAPI specification file (a.k.a. swagger) (`Block`).

<a id="api-specification-validation-custom-list-settings"></a>

### API Specification Validation Custom List Settings

`oversized_body_fail_validation` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`oversized_body_skip_validation` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`property_validation_settings_custom` - (Optional) Validation Property Settings. Custom property validation settings (`Block`).

`property_validation_settings_default` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

<a id="app-firewall"></a>

### App Firewall

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

<a id="blocked-clients"></a>

### Blocked Clients

`actions` - (Optional) Actions. Actions that should be taken when client identifier matches the rule (`List`).

`as_number` - (Optional) AS Number. RFC 6793 defined 4-byte AS number (`Number`).

`bot_skip_processing` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`expiration_timestamp` - (Optional) Expiration Timestamp. The expiration_timestamp is the RFC 3339 format timestamp at which the containing rule is considered to be logically expired. The rule continues to exist in the configuration but is not applied anymore (`String`).

`http_header` - (Optional) HTTP Header. Request header name and value pairs. See [HTTP Header](#blocked-clients-http-header) below.

`ip_prefix` - (Optional) IPv4 Prefix. IPv4 prefix string (`String`).

`ipv6_prefix` - (Optional) IPv6 Prefix. IPv6 prefix string (`String`).

`metadata` - (Optional) Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs. See [Metadata](#blocked-clients-metadata) below.

`skip_processing` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`user_identifier` - (Optional) User Identifier. Identify user based on user identifier. User identifier value needs to be copied from security event (`String`).

`waf_skip_processing` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

<a id="blocked-clients-http-header"></a>

### Blocked Clients HTTP Header

`headers` - (Optional) HTTP Headers. List of HTTP header name and value pairs. See [Headers](#blocked-clients-http-header-headers) below.

<a id="blocked-clients-http-header-headers"></a>

### Blocked Clients HTTP Header Headers

`exact` - (Optional) Exact. Header value to match exactly (`String`).

`invert_match` - (Optional) NOT of match. Invert the result of the match to detect missing header or non-matching value (`Bool`).

`name` - (Optional) Name. Name of the header (`String`).

`presence` - (Optional) Presence. If true, check for presence of header (`Bool`).

`regex` - (Optional) Regex. Regex match of the header value in re2 format (`String`).

<a id="blocked-clients-metadata"></a>

### Blocked Clients Metadata

`description` - (Optional) Description. Human readable description (`String`).

`name` - (Optional) Name. This is the name of the message. The value of name has to follow DNS-1035 format (`String`).

<a id="bot-defense"></a>

### Bot Defense

`disable_cors_support` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`enable_cors_support` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`policy` - (Optional) Bot Defense Policy. This defines various configuration options for Bot Defense policy. See [Policy](#bot-defense-policy) below.

`regional_endpoint` - (Optional) Bot Defense Region. Defines a selection for Bot Defense region - AUTO: AUTO Automatic selection based on client IP address - US: US US region - EU: EU European Union region - ASIA: ASIA Asia region. Possible values are `AUTO`, `US`, `EU`, `ASIA`. Defaults to `AUTO` (`String`).

`timeout` - (Optional) Timeout. The timeout for the inference check, in milliseconds (`Number`).

<a id="bot-defense-policy"></a>

### Bot Defense Policy

`disable_js_insert` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`disable_mobile_sdk` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`javascript_mode` - (Optional) Web Client JavaScript Mode. Web Client JavaScript Mode. Bot Defense JavaScript for telemetry collection is requested asynchronously, and it is non-cacheable Bot Defense JavaScript for telemetry collection is requested asynchronously, and it is cacheable Bot Defense JavaScript for telemetry collection is requested synchronously, and it is non-cacheable Bot Defense JavaScript for telemetry collection is requested synchronously, and it is cacheable. Possible values are `ASYNC_JS_NO_CACHING`, `ASYNC_JS_CACHING`, `SYNC_JS_NO_CACHING`, `SYNC_JS_CACHING`. Defaults to `ASYNC_JS_NO_CACHING` (`String`).

`js_download_path` - (Optional) JavaScript Download Path. Customize Bot Defense Client JavaScript path. If not specified, default `/common.js` (`String`).

`js_insert_all_pages` - (Optional) Insert Bot Defense JavaScript in All Pages. Insert Bot Defense JavaScript in all pages. See [Js Insert All Pages](#bot-defense-policy-js-insert-all-pages) below.

`js_insert_all_pages_except` - (Optional) Insert JavaScript in All Pages with the Exceptions. Insert Bot Defense JavaScript in all pages with the exceptions. See [Js Insert All Pages Except](#bot-defense-policy-js-insert-all-pages-except) below.

`js_insertion_rules` - (Optional) JavaScript Custom Insertion Rules. This defines custom JavaScript insertion rules for Bot Defense Policy. See [Js Insertion Rules](#bot-defense-policy-js-insertion-rules) below.

`mobile_sdk_config` - (Optional) Mobile SDK Configuration. Mobile SDK configuration. See [Mobile Sdk Config](#bot-defense-policy-mobile-sdk-config) below.

`protected_app_endpoints` - (Optional) App Endpoint Type. List of protected endpoints. Limit: Approx '128 endpoints per Load Balancer (LB)' upto 4 LBs, '32 endpoints per LB' after 4 LBs. See [Protected App Endpoints](#bot-defense-policy-protected-app-endpoints) below.

<a id="bot-defense-policy-js-insert-all-pages"></a>

### Bot Defense Policy Js Insert All Pages

`javascript_location` - (Optional) JavaScript Location. All inside networks. Insert JavaScript after <head> tag Insert JavaScript after </title> tag. Insert JavaScript before first <script> tag. Possible values are `AFTER_HEAD`, `AFTER_TITLE_END`, `BEFORE_SCRIPT`. Defaults to `AFTER_HEAD` (`String`).

<a id="bot-defense-policy-js-insert-all-pages-except"></a>

### Bot Defense Policy Js Insert All Pages Except

`exclude_list` - (Optional) Exclude Pages. Optional JavaScript insertions exclude list of domain and path matchers (`Block`).

`javascript_location` - (Optional) JavaScript Location. All inside networks. Insert JavaScript after <head> tag Insert JavaScript after </title> tag. Insert JavaScript before first <script> tag. Possible values are `AFTER_HEAD`, `AFTER_TITLE_END`, `BEFORE_SCRIPT`. Defaults to `AFTER_HEAD` (`String`).

<a id="bot-defense-policy-js-insertion-rules"></a>

### Bot Defense Policy Js Insertion Rules

`exclude_list` - (Optional) Exclude Paths. Optional JavaScript insertions exclude list of domain and path matchers (`Block`).

`rules` - (Optional) JavaScript Insertions. Required list of pages to insert Bot Defense client JavaScript (`Block`).

<a id="bot-defense-policy-mobile-sdk-config"></a>

### Bot Defense Policy Mobile Sdk Config

`mobile_identifier` - (Optional) Mobile Traffic Identifier. Mobile traffic identifier type (`Block`).

<a id="bot-defense-policy-protected-app-endpoints"></a>

### Bot Defense Policy Protected App Endpoints

`allow_good_bots` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`any_domain` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`domain` - (Optional) Domains. Domains names (`Block`).

`flow_label` - (Optional) Bot Defense Flow Label Category. Bot Defense Flow Label Category allows to associate traffic with selected category (`Block`).

`headers` - (Optional) HTTP Headers. A list of predicates for various HTTP headers that need to match. The criteria for matching each HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values are extracted from the request API as a list of strings for each HTTP header type. Note that all specified header predicates must evaluate to true (`Block`).

`http_methods` - (Optional) HTTP Methods. List of HTTP methods (`List`).

`metadata` - (Optional) Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs (`Block`).

`mitigate_good_bots` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`mitigation` - (Optional) Bot Mitigation Action. Modify Bot Defense behavior for a matching request (`Block`).

`mobile` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`path` - (Optional) Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match (`Block`).

`protocol` - (Optional) URL Scheme. SchemeType is used to indicate URL scheme. - BOTH: BOTH URL scheme for HTTPS:// or `HTTP://.` - HTTP: HTTP URL scheme HTTP:// only. - HTTPS: HTTPS URL scheme HTTPS:// only. Possible values are `BOTH`, `HTTP`, `HTTPS`. Defaults to `BOTH` (`String`).

`query_params` - (Optional) HTTP Query Parameters. A list of predicates for all query parameters that need to be matched. The criteria for matching each query parameter are described in individual instances of QueryParameterMatcherType. The actual query parameter values are extracted from the request API as a list of strings for each query parameter name. Note that all specified query parameter predicates must evaluate to true (`Block`).

`undefined_flow_label` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`web` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`web_mobile` - (Optional) Web and Mobile traffic type. Web and Mobile traffic type (`Block`).

<a id="captcha-challenge"></a>

### Captcha Challenge

`cookie_expiry` - (Optional) Cookie Expiration Period. Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new challenge (`Number`).

`custom_page` - (Optional) Custom message for Captcha Challenge. Custom message is of type uri_ref. Currently supported URL schemes is string:///. For string:/// scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. 'Please Wait.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '<p> Please Wait </p>'. Base64 encoded string for this HTML is 'PHA+IFBsZWFzZSBXYWl0IDwvcD4=' (`String`).

<a id="client-side-defense"></a>

### Client Side Defense

`policy` - (Optional) Client-Side Defense Policy. This defines various configuration options for Client-Side Defense policy. See [Policy](#client-side-defense-policy) below.

<a id="client-side-defense-policy"></a>

### Client Side Defense Policy

`disable_js_insert` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`js_insert_all_pages` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`js_insert_all_pages_except` - (Optional) Insert JavaScript in All Pages with the Exceptions. Insert Client-Side Defense JavaScript in all pages with the exceptions. See [Js Insert All Pages Except](#client-side-defense-policy-js-insert-all-pages-except) below.

`js_insertion_rules` - (Optional) JavaScript Custom Insertion Rules. This defines custom JavaScript insertion rules for Client-Side Defense Policy. See [Js Insertion Rules](#client-side-defense-policy-js-insertion-rules) below.

<a id="client-side-defense-policy-js-insert-all-pages-except"></a>

### Client Side Defense Policy Js Insert All Pages Except

`exclude_list` - (Optional) Exclude Pages. Optional JavaScript insertions exclude list of domain and path matchers (`Block`).

<a id="client-side-defense-policy-js-insertion-rules"></a>

### Client Side Defense Policy Js Insertion Rules

`exclude_list` - (Optional) Exclude Paths. Optional JavaScript insertions exclude list of domain and path matchers (`Block`).

`rules` - (Optional) JavaScript Insertions. Required list of pages to insert Client-Side Defense client JavaScript (`Block`).

<a id="cors-policy"></a>

### CORS Policy

`allow_credentials` - (Optional) Allow Credentials. Specifies whether the resource allows credentials (`Bool`).

`allow_headers` - (Optional) Allow Headers. Specifies the content for the access-control-allow-headers header (`String`).

`allow_methods` - (Optional) Allow Methods. Specifies the content for the access-control-allow-methods header (`String`).

`allow_origin` - (Optional) Allow Origin. Specifies the origins that will be allowed to do CORS requests. An origin is allowed if either allow_origin or allow_origin_regex match (`List`).

`allow_origin_regex` - (Optional) Allow Origin Regex. Specifies regex patterns that match allowed origins. An origin is allowed if either allow_origin or allow_origin_regex match (`List`).

`disabled` - (Optional) Disabled. Disable the CorsPolicy for a particular route. This is useful when virtual-host has CorsPolicy, but we need to disable it on a specific route. The value of this field is ignored for virtual-host (`Bool`).

`expose_headers` - (Optional) Expose Headers. Specifies the content for the access-control-expose-headers header (`String`).

`maximum_age` - (Optional) Maximum Age. Specifies the content for the access-control-max-age header in seconds. This indicates the maximum number of seconds the results can be cached A value of -1 will disable caching. Maximum permitted value is 86400 seconds (24 hours) (`Number`).

<a id="csrf-policy"></a>

### CSRF Policy

`all_load_balancer_domains` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`custom_domain_list` - (Optional) Domain name list. List of domain names used for Host header matching. See [Custom Domain List](#csrf-policy-custom-domain-list) below.

`disabled` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

<a id="csrf-policy-custom-domain-list"></a>

### CSRF Policy Custom Domain List

`domains` - (Optional) Domain names. A list of domain names that will be matched to loadbalancer. These domains are not used for SNI match. Wildcard names are supported in the suffix or prefix form (`List`).

<a id="custom-cache-rule"></a>

### Custom Cache Rule

`cdn_cache_rules` - (Optional) CDN Cache Rule. Reference to CDN Cache Rule configuration object. See [CDN Cache Rules](#custom-cache-rule-cdn-cache-rules) below.

<a id="custom-cache-rule-cdn-cache-rules"></a>

### Custom Cache Rule CDN Cache Rules

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

<a id="data-guard-rules"></a>

### Data Guard Rules

`any_domain` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`apply_data_guard` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`exact_value` - (Optional) Exact Value. Exact domain name (`String`).

`metadata` - (Optional) Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs. See [Metadata](#data-guard-rules-metadata) below.

`path` - (Optional) Path to Match. Path match of the URI can be either be, Prefix match or exact match or regular expression match. See [Path](#data-guard-rules-path) below.

`skip_data_guard` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`suffix_value` - (Optional) Suffix Value. Suffix of domain name e.g 'xyz.com' will match '*.xyz.com' and 'xyz.com' (`String`).

<a id="data-guard-rules-metadata"></a>

### Data Guard Rules Metadata

`description` - (Optional) Description. Human readable description (`String`).

`name` - (Optional) Name. This is the name of the message. The value of name has to follow DNS-1035 format (`String`).

<a id="data-guard-rules-path"></a>

### Data Guard Rules Path

`path` - (Optional) Exact. Exact path value to match (`String`).

`prefix` - (Optional) Prefix. Path prefix to match (e.g. the value / will match on all paths) (`String`).

`regex` - (Optional) Regex. Regular expression of path match (e.g. the value .* will match on all paths) (`String`).

<a id="ddos-mitigation-rules"></a>

### DDOS Mitigation Rules

`block` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`ddos_client_source` - (Optional) DDOS Client Source Choice. DDOS Mitigation sources to be blocked. See [DDOS Client Source](#ddos-mitigation-rules-ddos-client-source) below.

`expiration_timestamp` - (Optional) Expiration Timestamp. The expiration_timestamp is the RFC 3339 format timestamp at which the containing rule is considered to be logically expired. The rule continues to exist in the configuration but is not applied anymore (`String`).

`ip_prefix_list` - (Optional) IP Prefix Match List. List of IP Prefix strings to match against. See [IP Prefix List](#ddos-mitigation-rules-ip-prefix-list) below.

`metadata` - (Optional) Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs. See [Metadata](#ddos-mitigation-rules-metadata) below.

<a id="ddos-mitigation-rules-ddos-client-source"></a>

### DDOS Mitigation Rules DDOS Client Source

`asn_list` - (Optional) ASN Match List. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer. See [Asn List](#ddos-mitigation-rules-ddos-client-source-asn-list) below.

`country_list` - (Optional) Country List. Sources that are located in one of the countries in the given list (`List`).

`ja4_tls_fingerprint_matcher` - (Optional) JA4 TLS Fingerprint Matcher. An extended version of JA3 that includes additional fields for more comprehensive fingerprinting of SSL/TLS clients and potentially has a different structure and length. See [Ja4 TLS Fingerprint Matcher](#ddos-mitigation-rules-ddos-client-source-ja4-tls-fingerprint-matcher) below.

`tls_fingerprint_matcher` - (Optional) TLS Fingerprint Matcher. A TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of supported positve match criteria includes a list of known classes of TLS fingerprints and a list of exact values. The match is considered successful if either of these positive criteria are satisfied and the input fingerprint is not one of the excluded values. See [TLS Fingerprint Matcher](#ddos-mitigation-rules-ddos-client-source-tls-fingerprint-matcher) below.

<a id="ddos-mitigation-rules-ddos-client-source-asn-list"></a>

### DDOS Mitigation Rules DDOS Client Source Asn List

`as_numbers` - (Optional) AS Numbers. An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer (`List`).

<a id="ddos-mitigation-rules-ddos-client-source-ja4-tls-fingerprint-matcher"></a>

### DDOS Mitigation Rules DDOS Client Source Ja4 TLS Fingerprint Matcher

`exact_values` - (Optional) Exact Values. A list of exact JA4 TLS fingerprint to match the input JA4 TLS fingerprint against (`List`).

<a id="ddos-mitigation-rules-ddos-client-source-tls-fingerprint-matcher"></a>

### DDOS Mitigation Rules DDOS Client Source TLS Fingerprint Matcher

`classes` - (Optional) TLS fingerprint classes. A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against (`List`).

`exact_values` - (Optional) Exact Values. A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against (`List`).

`excluded_values` - (Optional) Excluded Values. A list of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can be used to skip known false positives when using one or more known TLS fingerprint classes in the enclosing matcher (`List`).

<a id="ddos-mitigation-rules-ip-prefix-list"></a>

### DDOS Mitigation Rules IP Prefix List

`invert_match` - (Optional) Invert Match Result. Invert the match result (`Bool`).

`ip_prefixes` - (Optional) IPv4 Prefix List. List of IPv4 prefix strings (`List`).

<a id="ddos-mitigation-rules-metadata"></a>

### DDOS Mitigation Rules Metadata

`description` - (Optional) Description. Human readable description (`String`).

`name` - (Optional) Name. This is the name of the message. The value of name has to follow DNS-1035 format (`String`).

<a id="default-cache-action"></a>

### Default Cache Action

`cache_disabled` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`cache_ttl_default` - (Optional) Fallback Cache TTL (d/ h/ m). Use Cache TTL Provided by Origin, and set a contigency TTL value in case one is not provided (`String`).

`cache_ttl_override` - (Optional) Override Cache TTL (d/ h/ m/ s). Always override the Cahce TTL provided by Origin (`String`).

<a id="enable-api-discovery"></a>

### Enable API Discovery

`api_crawler` - (Optional) API Crawling. API Crawler message. See [API Crawler](#enable-api-discovery-api-crawler) below.

`api_discovery_from_code_scan` - (Optional) Select Code Base and Repositories. x-required. See [API Discovery From Code Scan](#enable-api-discovery-api-discovery-from-code-scan) below.

`custom_api_auth_discovery` - (Optional) API Discovery Advanced Settings. API Discovery Advanced settings. See [Custom API Auth Discovery](#enable-api-discovery-custom-api-auth-discovery) below.

`default_api_auth_discovery` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`disable_learn_from_redirect_traffic` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`discovered_api_settings` - (Optional) Discovered API Settings. x-example: '2' Configure Discovered API Settings. See [Discovered API Settings](#enable-api-discovery-discovered-api-settings) below.

`enable_learn_from_redirect_traffic` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

<a id="enable-api-discovery-api-crawler"></a>

### Enable API Discovery API Crawler

`api_crawler_config` - (Optional) Crawler Configure. See [API Crawler Config](#enable-api-discovery-api-crawler-api-crawler-config) below.

`disable_api_crawler` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

<a id="enable-api-discovery-api-crawler-api-crawler-config"></a>

### Enable API Discovery API Crawler API Crawler Config

`domains` - (Optional) Domains to Crawl. Enter domains and their credentials to allow authenticated API crawling. You can only include domains you own that are associated with this Load Balancer (`Block`).

<a id="enable-api-discovery-api-discovery-from-code-scan"></a>

### Enable API Discovery API Discovery From Code Scan

`code_base_integrations` - (Optional) Select Code Base Integrations. See [Code Base Integrations](#enable-api-discovery-api-discovery-from-code-scan-code-base-integrations) below.

<a id="enable-api-discovery-api-discovery-from-code-scan-code-base-integrations"></a>

### Enable API Discovery API Discovery From Code Scan Code Base Integrations

`all_repos` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`code_base_integration` - (Optional) Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name (`Block`).

`selected_repos` - (Optional) API Code Repositories. Select which API repositories represent the LB applications (`Block`).

<a id="enable-api-discovery-custom-api-auth-discovery"></a>

### Enable API Discovery Custom API Auth Discovery

`api_discovery_ref` - (Optional) Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name. See [API Discovery Ref](#enable-api-discovery-custom-api-auth-discovery-api-discovery-ref) below.

<a id="enable-api-discovery-custom-api-auth-discovery-api-discovery-ref"></a>

### Enable API Discovery Custom API Auth Discovery API Discovery Ref

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

<a id="enable-api-discovery-discovered-api-settings"></a>

### Enable API Discovery Discovered API Settings

`purge_duration_for_inactive_discovered_apis` - (Optional) Purge Duration for Inactive Discovered APIs from Traffic. Inactive discovered API will be deleted after configured duration (`Number`).

<a id="enable-challenge"></a>

### Enable Challenge

`captcha_challenge_parameters` - (Optional) Captcha Challenge Parameters. Enables loadbalancer to perform captcha challenge Captcha challenge will be based on Google Recaptcha. With this feature enabled, only clients that pass the captcha challenge will be allowed to complete the HTTP request. When loadbalancer is configured to do Captcha Challenge, it will redirect the browser to an HTML page on every new HTTP request. This HTML page will have captcha challenge embedded in it. Client will be allowed to make the request only if the captcha challenge is successful. Loadbalancer will tag response header with a cookie to avoid Captcha challenge for subsequent requests. CAPTCHA is mainly used as a security check to ensure only human users can pass through. Generally, computers or bots are not capable of solving a captcha. You can enable either Javascript challenge or Captcha challenge on a virtual host. See [Captcha Challenge Parameters](#enable-challenge-captcha-challenge-parameters) below.

`default_captcha_challenge_parameters` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`default_js_challenge_parameters` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`default_mitigation_settings` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`js_challenge_parameters` - (Optional) Javascript Challenge Parameters. Enables loadbalancer to perform client browser compatibility test by redirecting to a page with Javascript. With this feature enabled, only clients that are capable of executing Javascript(mostly browsers) will be allowed to complete the HTTP request. When loadbalancer is configured to do Javascript Challenge, it will redirect the browser to an HTML page on every new HTTP request. This HTML page will have Javascript embedded in it. Loadbalancer chooses a set of random numbers for every new client and sends these numbers along with an encrypted answer with the request such that it embed these numbers as input in the Javascript. Javascript will run on the requestor browser and perform a complex Math operation. Script will submit the answer to loadbalancer. Loadbalancer will validate the answer by comparing the calculated answer with the decrypted answer (which was encrypted when it was sent back as reply) and allow the request to the upstream server only if the answer is correct. Loadbalancer will tag response header with a cookie to avoid Javascript challenge for subsequent requests. Javascript challenge serves following purposes * Validate that the request is coming via a browser that is capable for running Javascript * Force the browser to run a complex operation, f(X), that requires it to spend a large number of CPU cycles. This is to slow down a potential DOS attacker by making it difficult to launch a large request flood without having to spend even larger CPU cost at their end. You can enable either Javascript challenge or Captcha challenge on a virtual host. See [Js Challenge Parameters](#enable-challenge-js-challenge-parameters) below.

`malicious_user_mitigation` - (Optional) Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name. See [Malicious User Mitigation](#enable-challenge-malicious-user-mitigation) below.

<a id="enable-challenge-captcha-challenge-parameters"></a>

### Enable Challenge Captcha Challenge Parameters

`cookie_expiry` - (Optional) Cookie Expiration Period. Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new challenge (`Number`).

`custom_page` - (Optional) Custom message for Captcha Challenge. Custom message is of type uri_ref. Currently supported URL schemes is string:///. For string:/// scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. 'Please Wait.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '<p> Please Wait </p>'. Base64 encoded string for this HTML is 'PHA+IFBsZWFzZSBXYWl0IDwvcD4=' (`String`).

<a id="enable-challenge-js-challenge-parameters"></a>

### Enable Challenge Js Challenge Parameters

`cookie_expiry` - (Optional) Cookie Expiration Period. Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new challenge (`Number`).

`custom_page` - (Optional) Custom Message for Javascript Challenge. Custom message is of type uri_ref. Currently supported URL schemes is string:///. For string:/// scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. 'Please Wait.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '<p> Please Wait </p>'. Base64 encoded string for this HTML is 'PHA+IFBsZWFzZSBXYWl0IDwvcD4=' (`String`).

`js_script_delay` - (Optional) Javascript Delay. Delay introduced by Javascript, in milliseconds (`Number`).

<a id="enable-challenge-malicious-user-mitigation"></a>

### Enable Challenge Malicious User Mitigation

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

<a id="enable-ip-reputation"></a>

### Enable IP Reputation

`ip_threat_categories` - (Optional) List of IP Threat Categories to choose. If the source IP matches on atleast one of the enabled IP threat categories, the request will be denied (`List`).

<a id="graphql-rules"></a>

### GraphQL Rules

`any_domain` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`exact_path` - (Optional) Path. Specifies the exact path to GraphQL endpoint. Default value is /GraphQL (`String`).

`exact_value` - (Optional) Exact Value. Exact domain name (`String`).

`graphql_settings` - (Optional) GraphQL Settings. GraphQL configuration. See [GraphQL Settings](#graphql-rules-graphql-settings) below.

`metadata` - (Optional) Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs. See [Metadata](#graphql-rules-metadata) below.

`method_get` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`method_post` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`suffix_value` - (Optional) Suffix Value. Suffix of domain name e.g 'xyz.com' will match '*.xyz.com' and 'xyz.com' (`String`).

<a id="graphql-rules-graphql-settings"></a>

### GraphQL Rules GraphQL Settings

`disable_introspection` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`enable_introspection` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`max_batched_queries` - (Optional) Maximum Batched Queries. Specify maximum number of queries in a single batched request (`Number`).

`max_depth` - (Optional) Maximum Structure Depth. Specify maximum depth for the GraphQL query (`Number`).

`max_total_length` - (Optional) Maximum Total Length. Specify maximum length in bytes for the GraphQL query (`Number`).

<a id="graphql-rules-metadata"></a>

### GraphQL Rules Metadata

`description` - (Optional) Description. Human readable description (`String`).

`name` - (Optional) Name. This is the name of the message. The value of name has to follow DNS-1035 format (`String`).

<a id="http"></a>

### HTTP

`dns_volterra_managed` - (Optional) Automatically Manage DNS Records. DNS records for domains will be managed automatically by F5 Distributed Cloud. As a prerequisite, the domain must be delegated to F5 Distributed Cloud using Delegated domain feature or a DNS CNAME record should be created in your DNS provider's portal (`Bool`).

`port` - (Optional) HTTP Listen Port. HTTP port to Listen (`Number`).

`port_ranges` - (Optional) Port Ranges. A string containing a comma separated list of port ranges. Each port range consists of a single port or two ports separated by '-' (`String`).

<a id="https"></a>

### HTTPS

`add_hsts` - (Optional) Add HSTS Header. Add HTTP Strict-Transport-Security response header (`Bool`).

`http_redirect` - (Optional) HTTP Redirect to HTTPS. Redirect HTTP traffic to HTTPS (`Bool`).

`tls_cert_options` - (Optional) TLS Options. TLS Certificate Options. See [TLS Cert Options](#https-tls-cert-options) below.

<a id="https-tls-cert-options"></a>

### HTTPS TLS Cert Options

`tls_cert_params` - (Optional) TLS Parameters. Select TLS Parameters and Certificates. See [TLS Cert Params](#https-tls-cert-options-tls-cert-params) below.

`tls_inline_params` - (Optional) Inline TLS Parameters. Inline TLS parameters. See [TLS Inline Params](#https-tls-cert-options-tls-inline-params) below.

<a id="https-tls-cert-options-tls-cert-params"></a>

### HTTPS TLS Cert Options TLS Cert Params

`certificates` - (Optional) Certificates. Select one or more certificates with any domain names (`Block`).

`no_mtls` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`tls_config` - (Optional) TLS Config. This defines various options to configure TLS configuration parameters (`Block`).

`use_mtls` - (Optional) Clients TLS validation context. Validation context for downstream client TLS connections (`Block`).

<a id="https-tls-cert-options-tls-inline-params"></a>

### HTTPS TLS Cert Options TLS Inline Params

`no_mtls` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`tls_certificates` - (Optional) TLS Certificates. Users can add one or more certificates that share the same set of domains. for example, domain.com and *.domain.com - but use different signature algorithms (`Block`).

`tls_config` - (Optional) TLS Config. This defines various options to configure TLS configuration parameters (`Block`).

`use_mtls` - (Optional) Clients TLS validation context. Validation context for downstream client TLS connections (`Block`).

<a id="https-auto-cert"></a>

### HTTPS Auto Cert

`add_hsts` - (Optional) Add HSTS Header. Add HTTP Strict-Transport-Security response header (`Bool`).

`http_redirect` - (Optional) HTTP Redirect to HTTPS. Redirect HTTP traffic to HTTPS (`Bool`).

`tls_config` - (Optional) TLS Config. This defines various options to configure TLS configuration parameters. See [TLS Config](#https-auto-cert-tls-config) below.

<a id="https-auto-cert-tls-config"></a>

### HTTPS Auto Cert TLS Config

`tls_11_plus` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`tls_12_plus` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

<a id="js-challenge"></a>

### Js Challenge

`cookie_expiry` - (Optional) Cookie Expiration Period. Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new challenge (`Number`).

`custom_page` - (Optional) Custom Message for Javascript Challenge. Custom message is of type uri_ref. Currently supported URL schemes is string:///. For string:/// scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. 'Please Wait.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '<p> Please Wait </p>'. Base64 encoded string for this HTML is 'PHA+IFBsZWFzZSBXYWl0IDwvcD4=' (`String`).

`js_script_delay` - (Optional) Javascript Delay. Delay introduced by Javascript, in milliseconds (`Number`).

<a id="jwt-validation"></a>

### JWT Validation

`action` - (Optional) Action. See [Action](#jwt-validation-action) below.

`jwks_config` - (Optional) JSON Web Key Set (JWKS). The JSON Web Key Set (JWKS) is a set of keys used to verify JSON Web Token (JWT) issued by the Authorization Server. See RFC 7517 for more details. See [Jwks Config](#jwt-validation-jwks-config) below.

`mandatory_claims` - (Optional) Mandatory Claims. Configurable Validation of mandatory Claims. See [Mandatory Claims](#jwt-validation-mandatory-claims) below.

`reserved_claims` - (Optional) Reserved claims configuration. Configurable Validation of reserved Claims. See [Reserved Claims](#jwt-validation-reserved-claims) below.

`target` - (Optional) Target. Define endpoints for which JWT token validation will be performed. See [Target](#jwt-validation-target) below.

`token_location` - (Optional) Token Location. Location of JWT in HTTP request. See [Token Location](#jwt-validation-token-location) below.

<a id="jwt-validation-action"></a>

### JWT Validation Action

`block` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`report` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

<a id="jwt-validation-jwks-config"></a>

### JWT Validation Jwks Config

`cleartext` - (Optional) JSON Web Key Set (JWKS). The JSON Web Key Set (JWKS) is a set of keys used to verify JSON Web Token (JWT) issued by the Authorization Server. See RFC 7517 for more details (`String`).

<a id="jwt-validation-mandatory-claims"></a>

### JWT Validation Mandatory Claims

`claim_names` - (Optional) Claim Names (`List`).

<a id="jwt-validation-reserved-claims"></a>

### JWT Validation Reserved Claims

`audience` - (Optional) Audiences. See [Audience](#jwt-validation-reserved-claims-audience) below.

`audience_disable` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`issuer` - (Optional) Exact Match (`String`).

`issuer_disable` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`validate_period_disable` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`validate_period_enable` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

<a id="jwt-validation-reserved-claims-audience"></a>

### JWT Validation Reserved Claims Audience

`audiences` - (Optional) Values (`List`).

<a id="jwt-validation-target"></a>

### JWT Validation Target

`all_endpoint` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`api_groups` - (Optional) API Groups. See [API Groups](#jwt-validation-target-api-groups) below.

`base_paths` - (Optional) Base Paths. See [Base Paths](#jwt-validation-target-base-paths) below.

<a id="jwt-validation-target-api-groups"></a>

### JWT Validation Target API Groups

`api_groups` - (Optional) API Groups (`List`).

<a id="jwt-validation-target-base-paths"></a>

### JWT Validation Target Base Paths

`base_paths` - (Optional) Prefix Values (`List`).

<a id="jwt-validation-token-location"></a>

### JWT Validation Token Location

`bearer_token` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

<a id="l7-ddos-action-js-challenge"></a>

### L7 DDOS Action Js Challenge

`cookie_expiry` - (Optional) Cookie Expiration Period. Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new challenge (`Number`).

`custom_page` - (Optional) Custom Message for Javascript Challenge. Custom message is of type uri_ref. Currently supported URL schemes is string:///. For string:/// scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. 'Please Wait.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '<p> Please Wait </p>'. Base64 encoded string for this HTML is 'PHA+IFBsZWFzZSBXYWl0IDwvcD4=' (`String`).

`js_script_delay` - (Optional) Javascript Delay. Delay introduced by Javascript, in milliseconds (`Number`).

<a id="origin-pool"></a>

### Origin Pool

`more_origin_options` - (Optional) Origin Byte Range Request Config. See [More Origin Options](#origin-pool-more-origin-options) below.

`no_tls` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`origin_request_timeout` - (Optional) Origin Request Timeout Duration. Configures the time after which a request to the origin will time out waiting for a response (`String`).

`origin_servers` - (Optional) List Of Origin Servers. List of original servers. See [Origin Servers](#origin-pool-origin-servers) below.

`public_name` - (Optional) Public DNS Name. Specify origin server with public DNS name. See [Public Name](#origin-pool-public-name) below.

`use_tls` - (Optional) TLS Parameters for Origin Servers. Upstream TLS Parameters. See [Use TLS](#origin-pool-use-tls) below.

<a id="origin-pool-more-origin-options"></a>

### Origin Pool More Origin Options

`enable_byte_range_request` - (Optional) Enable Origin Byte Range Requests. Choice to enable/disable byte range requests towards origin (`Bool`).

`websocket_proxy` - (Optional) Enable WebSocket proxy to the origin. Option to enable proxying of WebSocket connections to the origin server (`Bool`).

<a id="origin-pool-origin-servers"></a>

### Origin Pool Origin Servers

`port` - (Optional) Origin Server Port. Port the workload can be reached on (`Number`).

`public_ip` - (Optional) Public IP. Specify origin server with public IP address. See [Public IP](#origin-pool-origin-servers-public-ip) below.

`public_name` - (Optional) Public DNS Name. Specify origin server with public DNS name. See [Public Name](#origin-pool-origin-servers-public-name) below.

<a id="origin-pool-origin-servers-public-ip"></a>

### Origin Pool Origin Servers Public IP

`ip` - (Optional) Public IPv4. Public IPv4 address (`String`).

<a id="origin-pool-origin-servers-public-name"></a>

### Origin Pool Origin Servers Public Name

`dns_name` - (Optional) DNS Name. DNS Name (`String`).

`refresh_interval` - (Optional) DNS Refresh Interval. Interval for DNS refresh in seconds. Max value is 7 days as per `HTTPS://datatracker.ietf.org/doc/HTML/rfc8767` (`Number`).

<a id="origin-pool-public-name"></a>

### Origin Pool Public Name

`dns_name` - (Optional) DNS Name. DNS Name (`String`).

`refresh_interval` - (Optional) DNS Refresh Interval. Interval for DNS refresh in seconds. Max value is 7 days as per `HTTPS://datatracker.ietf.org/doc/HTML/rfc8767` (`Number`).

<a id="origin-pool-use-tls"></a>

### Origin Pool Use TLS

`default_session_key_caching` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`disable_session_key_caching` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`disable_sni` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`max_session_keys` - (Optional) Max Session Keys Cached. x-example:'25' Number of session keys that are cached (`Number`).

`no_mtls` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`skip_server_verification` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`sni` - (Optional) SNI Value. SNI value to be used (`String`).

`tls_config` - (Optional) TLS Config. This defines various options to configure TLS configuration parameters. See [TLS Config](#origin-pool-use-tls-tls-config) below.

`use_host_header_as_sni` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`use_mtls` - (Optional) mTLS Certificate. mTLS Client Certificate. See [Use mTLS](#origin-pool-use-tls-use-mtls) below.

`use_mtls_obj` - (Optional) Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name. See [Use mTLS Obj](#origin-pool-use-tls-use-mtls-obj) below.

`use_server_verification` - (Optional) TLS Validation Context for Origin Servers. Upstream TLS Validation Context. See [Use Server Verification](#origin-pool-use-tls-use-server-verification) below.

`volterra_trusted_ca` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

<a id="origin-pool-use-tls-tls-config"></a>

### Origin Pool Use TLS TLS Config

`custom_security` - (Optional) Custom Ciphers. This defines TLS protocol config including min/max versions and allowed ciphers (`Block`).

`default_security` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`low_security` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`medium_security` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

<a id="origin-pool-use-tls-use-mtls"></a>

### Origin Pool Use TLS Use mTLS

`tls_certificates` - (Optional) mTLS Client Certificate. mTLS Client Certificate (`Block`).

<a id="origin-pool-use-tls-use-mtls-obj"></a>

### Origin Pool Use TLS Use mTLS Obj

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

<a id="origin-pool-use-tls-use-server-verification"></a>

### Origin Pool Use TLS Use Server Verification

`trusted_ca` - (Optional) Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name (`Block`).

`trusted_ca_url` - (Optional) Inline Root CA Certificate (legacy). Upload a Root CA Certificate specifically for this Origin Pool for verification of server's certificate (`String`).

<a id="other-settings"></a>

### Other Settings

`add_location` - (Optional) Add Location. x-example: true Appends header x-volterra-location = <RE-site-name> in responses (`Bool`).

`header_options` - (Optional) Header Control. This defines various options related to request/response headers. See [Header Options](#other-settings-header-options) below.

`logging_options` - (Optional) Logging Options. This defines various options related to logging. See [Logging Options](#other-settings-logging-options) below.

<a id="other-settings-header-options"></a>

### Other Settings Header Options

`request_headers_to_add` - (Optional) Add Origin Request Headers. Headers are key-value pairs to be added to HTTP request being routed towards upstream. Headers specified at this level are applied after headers from matched Route are applied. See [Request Headers To Add](#other-settings-header-options-request-headers-to-add) below.

`request_headers_to_remove` - (Optional) Remove Origin Request Headers. List of keys of Headers to be removed from the HTTP request being sent towards upstream (`List`).

`response_headers_to_add` - (Optional) Add Response Headers. Headers are key-value pairs to be added to HTTP response being sent towards downstream. Headers specified at this level are applied after headers from matched Route are applied. See [Response Headers To Add](#other-settings-header-options-response-headers-to-add) below.

`response_headers_to_remove` - (Optional) Remove Response Headers. List of keys of Headers to be removed from the HTTP response being sent towards downstream (`List`).

<a id="other-settings-header-options-request-headers-to-add"></a>

### Other Settings Header Options Request Headers To Add

`append` - (Optional) Append. Should the value be appended? If true, the value is appended to existing values. Default value is do not append (`Bool`).

`name` - (Optional) Name. Name of the HTTP header (`String`).

`secret_value` - (Optional) Secret. SecretType is used in an object to indicate a sensitive/confidential field (`Block`).

`value` - (Optional) Value. Value of the HTTP header (`String`).

<a id="other-settings-header-options-response-headers-to-add"></a>

### Other Settings Header Options Response Headers To Add

`append` - (Optional) Append. Should the value be appended? If true, the value is appended to existing values. Default value is do not append (`Bool`).

`name` - (Optional) Name. Name of the HTTP header (`String`).

`secret_value` - (Optional) Secret. SecretType is used in an object to indicate a sensitive/confidential field (`Block`).

`value` - (Optional) Value. Value of the HTTP header (`String`).

<a id="other-settings-logging-options"></a>

### Other Settings Logging Options

`client_log_options` - (Optional) Headers to Log. List of headers to Log. See [Client Log Options](#other-settings-logging-options-client-log-options) below.

`origin_log_options` - (Optional) Headers to Log. List of headers to Log. See [Origin Log Options](#other-settings-logging-options-origin-log-options) below.

<a id="other-settings-logging-options-client-log-options"></a>

### Other Settings Logging Options Client Log Options

`header_list` - (Optional) Headers. List of headers (`List`).

<a id="other-settings-logging-options-origin-log-options"></a>

### Other Settings Logging Options Origin Log Options

`header_list` - (Optional) Headers. List of headers (`List`).

<a id="policy-based-challenge"></a>

### Policy Based Challenge

`always_enable_captcha_challenge` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`always_enable_js_challenge` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`captcha_challenge_parameters` - (Optional) Captcha Challenge Parameters. Enables loadbalancer to perform captcha challenge Captcha challenge will be based on Google Recaptcha. With this feature enabled, only clients that pass the captcha challenge will be allowed to complete the HTTP request. When loadbalancer is configured to do Captcha Challenge, it will redirect the browser to an HTML page on every new HTTP request. This HTML page will have captcha challenge embedded in it. Client will be allowed to make the request only if the captcha challenge is successful. Loadbalancer will tag response header with a cookie to avoid Captcha challenge for subsequent requests. CAPTCHA is mainly used as a security check to ensure only human users can pass through. Generally, computers or bots are not capable of solving a captcha. You can enable either Javascript challenge or Captcha challenge on a virtual host. See [Captcha Challenge Parameters](#policy-based-challenge-captcha-challenge-parameters) below.

`default_captcha_challenge_parameters` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`default_js_challenge_parameters` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`default_mitigation_settings` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`default_temporary_blocking_parameters` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`js_challenge_parameters` - (Optional) Javascript Challenge Parameters. Enables loadbalancer to perform client browser compatibility test by redirecting to a page with Javascript. With this feature enabled, only clients that are capable of executing Javascript(mostly browsers) will be allowed to complete the HTTP request. When loadbalancer is configured to do Javascript Challenge, it will redirect the browser to an HTML page on every new HTTP request. This HTML page will have Javascript embedded in it. Loadbalancer chooses a set of random numbers for every new client and sends these numbers along with an encrypted answer with the request such that it embed these numbers as input in the Javascript. Javascript will run on the requestor browser and perform a complex Math operation. Script will submit the answer to loadbalancer. Loadbalancer will validate the answer by comparing the calculated answer with the decrypted answer (which was encrypted when it was sent back as reply) and allow the request to the upstream server only if the answer is correct. Loadbalancer will tag response header with a cookie to avoid Javascript challenge for subsequent requests. Javascript challenge serves following purposes * Validate that the request is coming via a browser that is capable for running Javascript * Force the browser to run a complex operation, f(X), that requires it to spend a large number of CPU cycles. This is to slow down a potential DOS attacker by making it difficult to launch a large request flood without having to spend even larger CPU cost at their end. You can enable either Javascript challenge or Captcha challenge on a virtual host. See [Js Challenge Parameters](#policy-based-challenge-js-challenge-parameters) below.

`malicious_user_mitigation` - (Optional) Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name. See [Malicious User Mitigation](#policy-based-challenge-malicious-user-mitigation) below.

`no_challenge` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`rule_list` - (Optional) Challenge Rule List. List of challenge rules to be used in policy based challenge. See [Rule List](#policy-based-challenge-rule-list) below.

`temporary_user_blocking` - (Optional) Temporary User Blocking. Specifies configuration for temporary user blocking resulting from user behavior analysis. When Malicious User Mitigation is enabled from service policy rules, users' accessing the application will be analyzed for malicious activity and the configured mitigation actions will be taken on identified malicious users. These mitigation actions include setting up temporary blocking on that user. This configuration specifies settings on how that blocking should be done by the loadbalancer. See [Temporary User Blocking](#policy-based-challenge-temporary-user-blocking) below.

<a id="policy-based-challenge-captcha-challenge-parameters"></a>

### Policy Based Challenge Captcha Challenge Parameters

`cookie_expiry` - (Optional) Cookie Expiration Period. Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new challenge (`Number`).

`custom_page` - (Optional) Custom message for Captcha Challenge. Custom message is of type uri_ref. Currently supported URL schemes is string:///. For string:/// scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. 'Please Wait.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '<p> Please Wait </p>'. Base64 encoded string for this HTML is 'PHA+IFBsZWFzZSBXYWl0IDwvcD4=' (`String`).

<a id="policy-based-challenge-js-challenge-parameters"></a>

### Policy Based Challenge Js Challenge Parameters

`cookie_expiry` - (Optional) Cookie Expiration Period. Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new challenge (`Number`).

`custom_page` - (Optional) Custom Message for Javascript Challenge. Custom message is of type uri_ref. Currently supported URL schemes is string:///. For string:/// scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. 'Please Wait.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '<p> Please Wait </p>'. Base64 encoded string for this HTML is 'PHA+IFBsZWFzZSBXYWl0IDwvcD4=' (`String`).

`js_script_delay` - (Optional) Javascript Delay. Delay introduced by Javascript, in milliseconds (`Number`).

<a id="policy-based-challenge-malicious-user-mitigation"></a>

### Policy Based Challenge Malicious User Mitigation

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

<a id="policy-based-challenge-rule-list"></a>

### Policy Based Challenge Rule List

`rules` - (Optional) Rules. Rules that specify the match conditions and challenge type to be launched. When a challenge type is selected to be always enabled, these rules can be used to disable challenge or launch a different challenge for requests that match the specified conditions. See [Rules](#policy-based-challenge-rule-list-rules) below.

<a id="policy-based-challenge-rule-list-rules"></a>

### Policy Based Challenge Rule List Rules

`metadata` - (Optional) Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs (`Block`).

`spec` - (Optional) Challenge Rule Specification. A Challenge Rule consists of an unordered list of predicates and an action. The predicates are evaluated against a set of input fields that are extracted from or derived from an L7 request API. A request API is considered to match the rule if all predicates in the rule evaluate to true for that request. Any predicates that are not specified in a rule are implicitly considered to be true. If a request API matches a challenge rule, the configured challenge is enforced (`Block`).

<a id="policy-based-challenge-temporary-user-blocking"></a>

### Policy Based Challenge Temporary User Blocking

`custom_page` - (Optional) Custom Message for Temporary Blocking. Custom message is of type `uri_ref`. Currently supported URL schemes is `string:///`. For `string:///` scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. 'Blocked.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '<p> Blocked </p>'. Base64 encoded string for this HTML is 'PHA+IFBsZWFzZSBXYWl0IDwvcD4=' (`String`).

<a id="protected-cookies"></a>

### Protected Cookies

`add_httponly` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`add_secure` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`disable_tampering_protection` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`enable_tampering_protection` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`ignore_httponly` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`ignore_max_age` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`ignore_samesite` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`ignore_secure` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`max_age_value` - (Optional) Add Max Age. Add max age attribute (`Number`).

`name` - (Optional) Cookie Name. Name of the Cookie (`String`).

`samesite_lax` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`samesite_none` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`samesite_strict` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

<a id="rate-limit"></a>

### Rate Limit

`custom_ip_allowed_list` - (Optional) Custom IP Allowed List. IP Allowed list using existing ip_prefix_set objects. See [Custom IP Allowed List](#rate-limit-custom-ip-allowed-list) below.

`ip_allowed_list` - (Optional) IPv4 Prefix List. x-example: '192.168.20.0/24' List of IPv4 prefixes that represent an endpoint. See [IP Allowed List](#rate-limit-ip-allowed-list) below.

`no_ip_allowed_list` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`no_policies` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`policies` - (Optional) Rate Limiter Policy List. List of rate limiter policies to be applied. See [Policies](#rate-limit-policies) below.

`rate_limiter` - (Optional) Rate Limit Value. A tuple consisting of a rate limit period unit and the total number of allowed requests for that period. See [Rate Limiter](#rate-limit-rate-limiter) below.

<a id="rate-limit-custom-ip-allowed-list"></a>

### Rate Limit Custom IP Allowed List

`rate_limiter_allowed_prefixes` - (Optional) List of IP Prefix Sets. References to ip_prefix_set objects. Requests from source IP addresses that are covered by one of the allowed IP Prefixes are not subjected to rate limiting. See [Rate Limiter Allowed Prefixes](#rate-limit-custom-ip-allowed-list-rate-limiter-allowed-prefixes) below.

<a id="rate-limit-custom-ip-allowed-list-rate-limiter-allowed-prefixes"></a>

### Rate Limit Custom IP Allowed List Rate Limiter Allowed Prefixes

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

<a id="rate-limit-ip-allowed-list"></a>

### Rate Limit IP Allowed List

`prefixes` - (Optional) IPv4 Prefix List. List of IPv4 prefixes that represent an endpoint (`List`).

<a id="rate-limit-policies"></a>

### Rate Limit Policies

`policies` - (Optional) Rate Limiter Policies. Ordered list of rate limiter policies. See [Policies](#rate-limit-policies-policies) below.

<a id="rate-limit-policies-policies"></a>

### Rate Limit Policies Policies

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

<a id="rate-limit-rate-limiter"></a>

### Rate Limit Rate Limiter

`action_block` - (Optional) Rate Limit Block Action. Action where a user is blocked from making further requests after exceeding rate limit threshold. See [Action Block](#rate-limit-rate-limiter-action-block) below.

`burst_multiplier` - (Optional) Burst Multiplier. The maximum burst of requests to accommodate, expressed as a multiple of the rate (`Number`).

`disabled` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`leaky_bucket` - (Optional) Leaky Bucket Rate Limiter. Leaky-Bucket is the default rate limiter algorithm for F5 (`Block`).

`period_multiplier` - (Optional) Periods. This setting, combined with Per Period units, provides a duration (`Number`).

`token_bucket` - (Optional) Token Bucket Rate Limiter. Token-Bucket is a rate limiter algorithm that is stricter with enforcing limits (`Block`).

`total_number` - (Optional) Number Of Requests. The total number of allowed requests per rate-limiting period (`Number`).

`unit` - (Optional) Rate Limit Period Unit. Unit for the period per which the rate limit is applied. - SECOND: Second Rate limit period unit is seconds - MINUTE: Minute Rate limit period unit is minutes - HOUR: Hour Rate limit period unit is hours - DAY: Day Rate limit period unit is days. Possible values are `SECOND`, `MINUTE`, `HOUR`. Defaults to `SECOND` (`String`).

<a id="rate-limit-rate-limiter-action-block"></a>

### Rate Limit Rate Limiter Action Block

`hours` - (Optional) Hours. Input Duration Hours (`Block`).

`minutes` - (Optional) Minutes. Input Duration Minutes (`Block`).

`seconds` - (Optional) Seconds. Input Duration Seconds (`Block`).

<a id="sensitive-data-policy"></a>

### Sensitive Data Policy

`sensitive_data_policy_ref` - (Optional) Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name. See [Sensitive Data Policy Ref](#sensitive-data-policy-sensitive-data-policy-ref) below.

<a id="sensitive-data-policy-sensitive-data-policy-ref"></a>

### Sensitive Data Policy Sensitive Data Policy Ref

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

<a id="slow-ddos-mitigation"></a>

### Slow DDOS Mitigation

`disable_request_timeout` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`request_headers_timeout` - (Optional) Request Headers Timeout. The amount of time the client has to send only the headers on the request stream before the stream is cancelled. The default value is 10000 milliseconds. This setting provides protection against Slowloris attacks (`Number`).

`request_timeout` - (Optional) Custom Timeout (`Number`).

<a id="timeouts"></a>

### Timeouts

`create` - (Optional) A string that can be [parsed as a duration](`HTTPS://pkg.go.dev/time#ParseDuration`) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours) (`String`).

`delete` - (Optional) A string that can be [parsed as a duration](`HTTPS://pkg.go.dev/time#ParseDuration`) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are saved into state before the destroy operation occurs (`String`).

`read` - (Optional) A string that can be [parsed as a duration](`HTTPS://pkg.go.dev/time#ParseDuration`) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh is enabled (`String`).

`update` - (Optional) A string that can be [parsed as a duration](`HTTPS://pkg.go.dev/time#ParseDuration`) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours) (`String`).

<a id="trusted-clients"></a>

### Trusted Clients

`actions` - (Optional) Actions. Actions that should be taken when client identifier matches the rule (`List`).

`as_number` - (Optional) AS Number. RFC 6793 defined 4-byte AS number (`Number`).

`bot_skip_processing` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`expiration_timestamp` - (Optional) Expiration Timestamp. The expiration_timestamp is the RFC 3339 format timestamp at which the containing rule is considered to be logically expired. The rule continues to exist in the configuration but is not applied anymore (`String`).

`http_header` - (Optional) HTTP Header. Request header name and value pairs. See [HTTP Header](#trusted-clients-http-header) below.

`ip_prefix` - (Optional) IPv4 Prefix. IPv4 prefix string (`String`).

`ipv6_prefix` - (Optional) IPv6 Prefix. IPv6 prefix string (`String`).

`metadata` - (Optional) Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs. See [Metadata](#trusted-clients-metadata) below.

`skip_processing` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`user_identifier` - (Optional) User Identifier. Identify user based on user identifier. User identifier value needs to be copied from security event (`String`).

`waf_skip_processing` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

<a id="trusted-clients-http-header"></a>

### Trusted Clients HTTP Header

`headers` - (Optional) HTTP Headers. List of HTTP header name and value pairs. See [Headers](#trusted-clients-http-header-headers) below.

<a id="trusted-clients-http-header-headers"></a>

### Trusted Clients HTTP Header Headers

`exact` - (Optional) Exact. Header value to match exactly (`String`).

`invert_match` - (Optional) NOT of match. Invert the result of the match to detect missing header or non-matching value (`Bool`).

`name` - (Optional) Name. Name of the header (`String`).

`presence` - (Optional) Presence. If true, check for presence of header (`Bool`).

`regex` - (Optional) Regex. Regex match of the header value in re2 format (`String`).

<a id="trusted-clients-metadata"></a>

### Trusted Clients Metadata

`description` - (Optional) Description. Human readable description (`String`).

`name` - (Optional) Name. This is the name of the message. The value of name has to follow DNS-1035 format (`String`).

<a id="user-identification"></a>

### User Identification

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

<a id="waf-exclusion"></a>

### WAF Exclusion

`waf_exclusion_inline_rules` - (Optional) WAF Exclusion Inline Rules. A list of WAF exclusion rules that will be applied inline. See [WAF Exclusion Inline Rules](#waf-exclusion-waf-exclusion-inline-rules) below.

`waf_exclusion_policy` - (Optional) Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name. See [WAF Exclusion Policy](#waf-exclusion-waf-exclusion-policy) below.

<a id="waf-exclusion-waf-exclusion-inline-rules"></a>

### WAF Exclusion WAF Exclusion Inline Rules

`rules` - (Optional) WAF Exclusion Rules. An ordered list of WAF Exclusions specific to this Load Balancer. See [Rules](#waf-exclusion-waf-exclusion-inline-rules-rules) below.

<a id="waf-exclusion-waf-exclusion-inline-rules-rules"></a>

### WAF Exclusion WAF Exclusion Inline Rules Rules

`any_domain` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`any_path` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`app_firewall_detection_control` - (Optional) App Firewall Detection Control. Define the list of Signature IDs, Violations, Attack Types and Bot Names that should be excluded from triggering on the defined match criteria (`Block`).

`exact_value` - (Optional) Exact Value. Exact domain name (`String`).

`expiration_timestamp` - (Optional) Expiration Timestamp. The expiration_timestamp is the RFC 3339 format timestamp at which the containing rule is considered to be logically expired. The rule continues to exist in the configuration but is not applied anymore (`String`).

`metadata` - (Optional) Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs (`Block`).

`methods` - (Optional) Methods. methods to be matched (`List`).

`path_prefix` - (Optional) Prefix. Path prefix to match (e.g. the value / will match on all paths) (`String`).

`path_regex` - (Optional) Path Regex. Define the regex for the path. For example, the regex ^/.*$ will match on all paths (`String`).

`suffix_value` - (Optional) Suffix Value. Suffix of domain name e.g 'xyz.com' will match '*.xyz.com' and 'xyz.com' (`String`).

`waf_skip_processing` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

<a id="waf-exclusion-waf-exclusion-policy"></a>

### WAF Exclusion WAF Exclusion Policy

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).
