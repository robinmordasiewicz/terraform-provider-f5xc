---
page_title: "Virtual Host Nested Blocks - f5xc Provider"
subcategory: "Load Balancing"
description: |-
  Nested block reference for the Virtual Host resource.
---

# Virtual Host Nested Blocks

This page contains detailed documentation for nested blocks in the `f5xc_virtual_host` resource.

For the main resource documentation, see [f5xc_virtual_host](./resources/virtual_host).

## Contents

- [advertise-policies](#advertise-policies)
- [authentication](#authentication)
- [authentication-auth-config](#authentication-auth-config)
- [authentication-cookie-params](#authentication-cookie-params)
- [authentication-cookie-params-auth-hmac](#authentication-cookie-params-auth-hmac)
- [buffer-policy](#buffer-policy)
- [captcha-challenge](#captcha-challenge)
- [coalescing-options](#coalescing-options)
- [compression-params](#compression-params)
- [cors-policy](#cors-policy)
- [csrf-policy](#csrf-policy)
- [csrf-policy-custom-domain-list](#csrf-policy-custom-domain-list)
- [dynamic-reverse-proxy](#dynamic-reverse-proxy)
- [dynamic-reverse-proxy-resolution-network](#dynamic-reverse-proxy-resolution-network)
- [http-protocol-options](#http-protocol-options)
- [http-protocol-options-http-protocol-enable-v1-only](#http-protocol-options-http-protocol-enable-v1-only)
- [http-protocol-options-http-protocol-enable-v1-only-header-transformation](#http-protocol-options-http-protocol-enable-v1-only-header-transformation)
- [js-challenge](#js-challenge)
- [rate-limiter-allowed-prefixes](#rate-limiter-allowed-prefixes)
- [request-cookies-to-add](#request-cookies-to-add)
- [request-cookies-to-add-secret-value](#request-cookies-to-add-secret-value)
- [request-cookies-to-add-secret-value-blindfold-secret-info](#request-cookies-to-add-secret-value-blindfold-secret-info)
- [request-cookies-to-add-secret-value-clear-secret-info](#request-cookies-to-add-secret-value-clear-secret-info)
- [request-headers-to-add](#request-headers-to-add)
- [request-headers-to-add-secret-value](#request-headers-to-add-secret-value)
- [request-headers-to-add-secret-value-blindfold-secret-info](#request-headers-to-add-secret-value-blindfold-secret-info)
- [request-headers-to-add-secret-value-clear-secret-info](#request-headers-to-add-secret-value-clear-secret-info)
- [response-cookies-to-add](#response-cookies-to-add)
- [response-cookies-to-add-secret-value](#response-cookies-to-add-secret-value)
- [response-cookies-to-add-secret-value-blindfold-secret-info](#response-cookies-to-add-secret-value-blindfold-secret-info)
- [response-cookies-to-add-secret-value-clear-secret-info](#response-cookies-to-add-secret-value-clear-secret-info)
- [response-headers-to-add](#response-headers-to-add)
- [response-headers-to-add-secret-value](#response-headers-to-add-secret-value)
- [response-headers-to-add-secret-value-blindfold-secret-info](#response-headers-to-add-secret-value-blindfold-secret-info)
- [response-headers-to-add-secret-value-clear-secret-info](#response-headers-to-add-secret-value-clear-secret-info)
- [retry-policy](#retry-policy)
- [retry-policy-back-off](#retry-policy-back-off)
- [routes](#routes)
- [sensitive-data-policy](#sensitive-data-policy)
- [slow-ddos-mitigation](#slow-ddos-mitigation)
- [timeouts](#timeouts)
- [tls-cert-params](#tls-cert-params)
- [tls-cert-params-certificates](#tls-cert-params-certificates)
- [tls-cert-params-validation-params](#tls-cert-params-validation-params)
- [tls-cert-params-validation-params-trusted-ca](#tls-cert-params-validation-params-trusted-ca)
- [tls-parameters](#tls-parameters)
- [tls-parameters-common-params](#tls-parameters-common-params)
- [tls-parameters-common-params-tls-certificates](#tls-parameters-common-params-tls-certificates)
- [tls-parameters-common-params-validation-params](#tls-parameters-common-params-validation-params)
- [user-identification](#user-identification)
- [waf-type](#waf-type)
- [waf-type-app-firewall](#waf-type-app-firewall)
- [waf-type-app-firewall-app-firewall](#waf-type-app-firewall-app-firewall)

---

<a id="advertise-policies"></a>

**Advertise Policies**

`kind` - (Optional) Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route') (`String`).

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

`uid` - (Optional) UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid (`String`).

---

<a id="authentication"></a>

**Authentication**

`auth_config` - (Optional) Reference to Authentication Object. Reference to Authentication Config Object. See [Auth Config](#authentication-auth-config) below.

`cookie_params` - (Optional) Cookie Parameters. Specifies different cookie related config parameters for authentication. See [Cookie Params](#authentication-cookie-params) below.

`redirect_dynamic` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`redirect_url` - (Optional) Configure Redirect URL. user can provide a URL for e.g `HTTPS://abc.xyz.com` where user gets redirected. This URL configured here must match with the redirect URL configured with the OIDC provider (`String`).

`use_auth_object_config` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

---

<a id="authentication-auth-config"></a>

**Authentication Auth Config**

`kind` - (Optional) Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route') (`String`).

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

`uid` - (Optional) UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid (`String`).

---

<a id="authentication-cookie-params"></a>

**Authentication Cookie Params**

`auth_hmac` - (Optional) HMAC Key Pair. HMAC primary and secondary keys to be used for hashing the Cookie. Each key also have an associated expiry timestamp, beyond which key is invalid. See [Auth HMAC](#authentication-cookie-params-auth-hmac) below.

`cookie_expiry` - (Optional) Cookie Expiry duration. specifies in seconds max duration of the allocated cookie. This maps to “Max-Age” attribute in the session cookie. This will act as an expiry duration on the client side after which client will not be setting the cookie as part of the request. Default cookie expiry is 3600 seconds (`Number`).

`cookie_refresh_interval` - (Optional) Cookie Refresh Interval. Specifies in seconds refresh interval for session cookie. This is used to keep the active user active and reduce RE-login. When an incoming cookie's session expiry is still valid, and time to expire falls behind this interval, RE-issue a cookie with new expiry and with the same original session expiry. Default refresh interval is 3000 seconds (`Number`).

`kms_key_hmac` - (Optional) KMS Key Reference. Reference to KMS Key Object (`Block`).

`session_expiry` - (Optional) Session Expiry duration. specifies in seconds max lifetime of an authenticated session after which the user will be forced to login again. Default session expiry is 86400 seconds(24 hours) (`Number`).

---

<a id="authentication-cookie-params-auth-hmac"></a>

**Authentication Cookie Params Auth HMAC**

`prim_key` - (Optional) Secret. SecretType is used in an object to indicate a sensitive/confidential field (`Block`).

`prim_key_expiry` - (Optional) HMAC Primary Key Expiry. Primary HMAC Key Expiry time (`String`).

`sec_key` - (Optional) Secret. SecretType is used in an object to indicate a sensitive/confidential field (`Block`).

`sec_key_expiry` - (Optional) HMAC Secondary Key Expiry. Secondary HMAC Key Expiry time (`String`).

---

<a id="buffer-policy"></a>

**Buffer Policy**

`disabled` - (Optional) Disable. Disable buffering for a particular route. This is useful when virtual-host has buffering, but we need to disable it on a specific route. The value of this field is ignored for virtual-host (`Bool`).

`max_request_bytes` - (Optional) Max Request Bytes. The maximum request size that the filter will buffer before the connection manager will stop buffering and return a RequestEntityTooLarge (413) response (`Number`).

---

<a id="captcha-challenge"></a>

**Captcha Challenge**

`cookie_expiry` - (Optional) Cookie Expiration Period. Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new challenge (`Number`).

`custom_page` - (Optional) Custom message for Captcha Challenge. Custom message is of type uri_ref. Currently supported URL schemes is string:///. For string:/// scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. 'Please Wait.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '<p> Please Wait </p>'. Base64 encoded string for this HTML is 'PHA+IFBsZWFzZSBXYWl0IDwvcD4=' (`String`).

---

<a id="coalescing-options"></a>

**Coalescing Options**

`default_coalescing` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`strict_coalescing` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

---

<a id="compression-params"></a>

**Compression Params**

`content_length` - (Optional) Content Length. Minimum response length, in bytes, which will trigger compression. The default value is 30 (`Number`).

`content_type` - (Optional) Content Type. Set of strings that allows specifying which mime-types yield compression When this field is not defined, compression will be applied to the following mime-types: 'application/javascript' 'application/JSON', 'application/xhtml+XML' 'image/svg+XML' 'text/CSS' 'text/HTML' 'text/plain' 'text/XML' (`List`).

`disable_on_etag_header` - (Optional) Disable On Etag Header. If true, disables compression when the response contains an etag header. When it is false, weak etags will be preserved and the ones that require strong validation will be removed (`Bool`).

`remove_accept_encoding_header` - (Optional) Remove Accept-Encoding Header. If true, removes accept-encoding from the request headers before dispatching it to the upstream so that responses do not get compressed before reaching the filter (`Bool`).

---

<a id="cors-policy"></a>

**CORS Policy**

`allow_credentials` - (Optional) Allow Credentials. Specifies whether the resource allows credentials (`Bool`).

`allow_headers` - (Optional) Allow Headers. Specifies the content for the access-control-allow-headers header (`String`).

`allow_methods` - (Optional) Allow Methods. Specifies the content for the access-control-allow-methods header (`String`).

`allow_origin` - (Optional) Allow Origin. Specifies the origins that will be allowed to do CORS requests. An origin is allowed if either allow_origin or allow_origin_regex match (`List`).

`allow_origin_regex` - (Optional) Allow Origin Regex. Specifies regex patterns that match allowed origins. An origin is allowed if either allow_origin or allow_origin_regex match (`List`).

`disabled` - (Optional) Disabled. Disable the CorsPolicy for a particular route. This is useful when virtual-host has CorsPolicy, but we need to disable it on a specific route. The value of this field is ignored for virtual-host (`Bool`).

`expose_headers` - (Optional) Expose Headers. Specifies the content for the access-control-expose-headers header (`String`).

`maximum_age` - (Optional) Maximum Age. Specifies the content for the access-control-max-age header in seconds. This indicates the maximum number of seconds the results can be cached A value of -1 will disable caching. Maximum permitted value is 86400 seconds (24 hours) (`Number`).

---

<a id="csrf-policy"></a>

**CSRF Policy**

`all_load_balancer_domains` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`custom_domain_list` - (Optional) Domain name list. List of domain names used for Host header matching. See [Custom Domain List](#csrf-policy-custom-domain-list) below.

`disabled` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

---

<a id="csrf-policy-custom-domain-list"></a>

**CSRF Policy Custom Domain List**

`domains` - (Optional) Domain names. A list of domain names that will be matched to loadbalancer. These domains are not used for SNI match. Wildcard names are supported in the suffix or prefix form (`List`).

---

<a id="dynamic-reverse-proxy"></a>

**Dynamic Reverse Proxy**

`connection_timeout` - (Optional) Connection Timeout. The timeout for new network connections to upstream server. This is specified in milliseconds. The default value is 2000 (2 seconds) (`Number`).

`resolution_network` - (Optional) Resolution Network. Reference to virtual network where the endpoint is resolved. Reference is valid only when the network type is VIRTUAL_NETWORK_PER_SITE or VIRTUAL_NETWORK_GLOBAL. It is ignored for all other network types. See [Resolution Network](#dynamic-reverse-proxy-resolution-network) below.

`resolution_network_type` - (Optional) Virtual Network Type. Different types of virtual networks understood by the system Virtual-network of type VIRTUAL_NETWORK_SITE_LOCAL provides connectivity to public (outside) network. This is an insecure network and is connected to public internet via NAT Gateways/firwalls Virtual-network of this type is local to every site. Two virtual networks of this type on different sites are neither related nor connected. Constraints: There can be atmost one virtual network of this type in a given site. This network type is supported on CE sites. This network is created automatically and present on all sites Virtual-network of type VIRTUAL_NETWORK_SITE_LOCAL_INSIDE is a private network inside site. It is a secure network and is not connected to public network. Virtual-network of this type is local to every site. Two virtual networks of this type on different sites are neither related nor connected. Constraints: There can be atmost one virtual network of this type in a given site. This network type is supported on CE sites. This network is created during provisioning of site User defined per-site virtual network. Scope of this virtual network is limited to the site. This is not yet supported Virtual-network of type VIRTUAL_NETWORK_PUBLIC directly conects to the public internet. Virtual-network of this type is local to every site. Two virtual networks of this type on different sites are neither related nor connected. Constraints: There can be atmost one virtual network of this type in a given site. This network type is supported on RE sites only It is an internally created by the system. They must not be created by user Virtual Neworks with global scope across different sites in F5XC domain. An example global virtual-network called 'AIN Network' is created for every tenant. for volterra fabric Constraints: It is currently only supported as internally created by the system. vK8s service network for a given tenant. Used to advertise a virtual host only to vk8s pods for that tenant Constraints: It is an internally created by the system. Must not be created by user VER internal network for the site. It can only be used for virtual hosts with SMA_PROXY type proxy Constraints: It is an internally created by the system. Must not be created by user Virtual-network of type VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE represents both VIRTUAL_NETWORK_SITE_LOCAL and VIRTUAL_NETWORK_SITE_LOCAL_INSIDE Constraints: This network type is only meaningful in an advertise policy When virtual-network of type VIRTUAL_NETWORK_IP_AUTO is selected for an endpoint, VER will try to determine the network based on the provided IP address Constraints: This network type is only meaningful in an endpoint VoltADN Private Network is used on volterra RE(s) to connect to customer private networks This network is created by opening a support ticket This network is per site srv6 network VER IP Fabric network for the site. This Virtual network type is used for exposing virtual host on IP Fabric network on the VER site or for endpoint in IP Fabric network Constraints: It is an internally created by the system. Must not be created by user Network internally created for a segment Constraints: It is an internally created by the system. Must not be created by user. Possible values are `VIRTUAL_NETWORK_SITE_LOCAL`, `VIRTUAL_NETWORK_SITE_LOCAL_INSIDE`, `VIRTUAL_NETWORK_PER_SITE`, `VIRTUAL_NETWORK_PUBLIC`, `VIRTUAL_NETWORK_GLOBAL`, `VIRTUAL_NETWORK_SITE_SERVICE`, `VIRTUAL_NETWORK_VER_INTERNAL`, `VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE`, `VIRTUAL_NETWORK_IP_AUTO`, `VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK`, `VIRTUAL_NETWORK_SRV6_NETWORK`, `VIRTUAL_NETWORK_IP_FABRIC`, `VIRTUAL_NETWORK_SEGMENT`. Defaults to `VIRTUAL_NETWORK_SITE_LOCAL` (`String`).

`resolve_endpoint_dynamically` - (Optional) Dynamic Endpoint Resolution. x-example : true In this mode of proxy, virtual host will resolve the destination endpoint dynamically. The dynamic resolution is done using a predefined field in the request. This predefined field depends on the ProxyType configured on the Virtual Host. For HTTP traffic, i.e. with ProxyType as HTTP_PROXY or HTTPS_PROXY, virtual host will use the 'HOST' HTTP header from the request and perform DNS resolution to select destination endpoint. For TCP traffic with SNI, (If the ProxyType is TCP_PROXY_WITH_SNI), virtual host will perform DNS resolution using the SNI. The DNS resolution is performed in the virtual network specified in outside_network_type or outside_network In both modes of operation(either using Host header or SNI), the DNS resolution could return multiple addresses. First IPv4 address from such returned list is used as endpoint for the request. The DNS response is cached for 60s by default (`Bool`).

---

<a id="dynamic-reverse-proxy-resolution-network"></a>

**Dynamic Reverse Proxy Resolution Network**

`kind` - (Optional) Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route') (`String`).

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

`uid` - (Optional) UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid (`String`).

---

<a id="http-protocol-options"></a>

**HTTP Protocol Options**

`http_protocol_enable_v1_only` - (Optional) HTTP/1.1 Protocol Options. HTTP/1.1 Protocol options for downstream connections. See [HTTP Protocol Enable V1 Only](#http-protocol-options-http-protocol-enable-v1-only) below.

`http_protocol_enable_v1_v2` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`http_protocol_enable_v2_only` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

---

<a id="http-protocol-options-http-protocol-enable-v1-only"></a>

**HTTP Protocol Options HTTP Protocol Enable V1 Only**

`header_transformation` - (Optional) Header Transformation. Header Transformation options for HTTP/1.1 request/response headers. See [Header Transformation](#http-protocol-options-http-protocol-enable-v1-only-header-transformation) below.

---

<a id="http-protocol-options-http-protocol-enable-v1-only-header-transformation"></a>

**HTTP Protocol Options HTTP Protocol Enable V1 Only Header Transformation**

`default_header_transformation` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`legacy_header_transformation` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`preserve_case_header_transformation` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`proper_case_header_transformation` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

---

<a id="js-challenge"></a>

**Js Challenge**

`cookie_expiry` - (Optional) Cookie Expiration Period. Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new challenge (`Number`).

`custom_page` - (Optional) Custom Message for Javascript Challenge. Custom message is of type uri_ref. Currently supported URL schemes is string:///. For string:/// scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. 'Please Wait.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '<p> Please Wait </p>'. Base64 encoded string for this HTML is 'PHA+IFBsZWFzZSBXYWl0IDwvcD4=' (`String`).

`js_script_delay` - (Optional) Javascript Delay. Delay introduced by Javascript, in milliseconds (`Number`).

---

<a id="rate-limiter-allowed-prefixes"></a>

**Rate Limiter Allowed Prefixes**

`kind` - (Optional) Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route') (`String`).

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

`uid` - (Optional) UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid (`String`).

---

<a id="request-cookies-to-add"></a>

**Request Cookies To Add**

`name` - (Optional) Name. Name of the cookie in Cookie header (`String`).

`overwrite` - (Optional) Overwrite. Should the value be overwritten? If true, the value is overwritten to existing values. Default value is do not overwrite (`Bool`).

`secret_value` - (Optional) Secret. SecretType is used in an object to indicate a sensitive/confidential field. See [Secret Value](#request-cookies-to-add-secret-value) below.

`value` - (Optional) Value. Value of the Cookie header (`String`).

---

<a id="request-cookies-to-add-secret-value"></a>

**Request Cookies To Add Secret Value**

`blindfold_secret_info` - (Optional) Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management. See [Blindfold Secret Info](#request-cookies-to-add-secret-value-blindfold-secret-info) below.

`clear_secret_info` - (Optional) In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted. See [Clear Secret Info](#request-cookies-to-add-secret-value-clear-secret-info) below.

---

<a id="request-cookies-to-add-secret-value-blindfold-secret-info"></a>

**Request Cookies To Add Secret Value Blindfold Secret Info**

`decryption_provider` - (Optional) Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service (`String`).

`location` - (Optional) Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location (`String`).

`store_provider` - (Optional) Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:/// (`String`).

---

<a id="request-cookies-to-add-secret-value-clear-secret-info"></a>

**Request Cookies To Add Secret Value Clear Secret Info**

`provider_ref` - (Optional) Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:/// (`String`).

`url` - (Optional) URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding (`String`).

---

<a id="request-headers-to-add"></a>

**Request Headers To Add**

`append` - (Optional) Append. Should the value be appended? If true, the value is appended to existing values. Default value is do not append (`Bool`).

`name` - (Optional) Name. Name of the HTTP header (`String`).

`secret_value` - (Optional) Secret. SecretType is used in an object to indicate a sensitive/confidential field. See [Secret Value](#request-headers-to-add-secret-value) below.

`value` - (Optional) Value. Value of the HTTP header (`String`).

---

<a id="request-headers-to-add-secret-value"></a>

**Request Headers To Add Secret Value**

`blindfold_secret_info` - (Optional) Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management. See [Blindfold Secret Info](#request-headers-to-add-secret-value-blindfold-secret-info) below.

`clear_secret_info` - (Optional) In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted. See [Clear Secret Info](#request-headers-to-add-secret-value-clear-secret-info) below.

---

<a id="request-headers-to-add-secret-value-blindfold-secret-info"></a>

**Request Headers To Add Secret Value Blindfold Secret Info**

`decryption_provider` - (Optional) Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service (`String`).

`location` - (Optional) Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location (`String`).

`store_provider` - (Optional) Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:/// (`String`).

---

<a id="request-headers-to-add-secret-value-clear-secret-info"></a>

**Request Headers To Add Secret Value Clear Secret Info**

`provider_ref` - (Optional) Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:/// (`String`).

`url` - (Optional) URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding (`String`).

---

<a id="response-cookies-to-add"></a>

**Response Cookies To Add**

`add_domain` - (Optional) Add Domain. Add domain attribute (`String`).

`add_expiry` - (Optional) Add expiry. Add expiry attribute (`String`).

`add_httponly` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`add_partitioned` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`add_path` - (Optional) Add path. Add path attribute (`String`).

`add_secure` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`ignore_domain` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`ignore_expiry` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`ignore_httponly` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`ignore_max_age` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`ignore_partitioned` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`ignore_path` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`ignore_samesite` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`ignore_secure` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`ignore_value` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`max_age_value` - (Optional) Add Max Age. Add max age attribute (`Number`).

`name` - (Optional) Name. Name of the cookie in Cookie header (`String`).

`overwrite` - (Optional) Overwrite. Should the value be overwritten? If true, the value is overwritten to existing values. Default value is do not overwrite (`Bool`).

`samesite_lax` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`samesite_none` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`samesite_strict` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`secret_value` - (Optional) Secret. SecretType is used in an object to indicate a sensitive/confidential field. See [Secret Value](#response-cookies-to-add-secret-value) below.

`value` - (Optional) Value. Value of the Cookie header (`String`).

---

<a id="response-cookies-to-add-secret-value"></a>

**Response Cookies To Add Secret Value**

`blindfold_secret_info` - (Optional) Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management. See [Blindfold Secret Info](#response-cookies-to-add-secret-value-blindfold-secret-info) below.

`clear_secret_info` - (Optional) In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted. See [Clear Secret Info](#response-cookies-to-add-secret-value-clear-secret-info) below.

---

<a id="response-cookies-to-add-secret-value-blindfold-secret-info"></a>

**Response Cookies To Add Secret Value Blindfold Secret Info**

`decryption_provider` - (Optional) Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service (`String`).

`location` - (Optional) Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location (`String`).

`store_provider` - (Optional) Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:/// (`String`).

---

<a id="response-cookies-to-add-secret-value-clear-secret-info"></a>

**Response Cookies To Add Secret Value Clear Secret Info**

`provider_ref` - (Optional) Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:/// (`String`).

`url` - (Optional) URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding (`String`).

---

<a id="response-headers-to-add"></a>

**Response Headers To Add**

`append` - (Optional) Append. Should the value be appended? If true, the value is appended to existing values. Default value is do not append (`Bool`).

`name` - (Optional) Name. Name of the HTTP header (`String`).

`secret_value` - (Optional) Secret. SecretType is used in an object to indicate a sensitive/confidential field. See [Secret Value](#response-headers-to-add-secret-value) below.

`value` - (Optional) Value. Value of the HTTP header (`String`).

---

<a id="response-headers-to-add-secret-value"></a>

**Response Headers To Add Secret Value**

`blindfold_secret_info` - (Optional) Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management. See [Blindfold Secret Info](#response-headers-to-add-secret-value-blindfold-secret-info) below.

`clear_secret_info` - (Optional) In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted. See [Clear Secret Info](#response-headers-to-add-secret-value-clear-secret-info) below.

---

<a id="response-headers-to-add-secret-value-blindfold-secret-info"></a>

**Response Headers To Add Secret Value Blindfold Secret Info**

`decryption_provider` - (Optional) Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service (`String`).

`location` - (Optional) Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location (`String`).

`store_provider` - (Optional) Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:/// (`String`).

---

<a id="response-headers-to-add-secret-value-clear-secret-info"></a>

**Response Headers To Add Secret Value Clear Secret Info**

`provider_ref` - (Optional) Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:/// (`String`).

`url` - (Optional) URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding (`String`).

---

<a id="retry-policy"></a>

**Retry Policy**

`back_off` - (Optional) Retry BackOff Interval. Specifies parameters that control retry back off. See [Back Off](#retry-policy-back-off) below.

`num_retries` - (Optional) Number of Retries. Specifies the allowed number of retries. Defaults to 1. Retries can be done any number of times. An exponential back-off algorithm is used between each retry (`Number`).

`per_try_timeout` - (Optional) Per Try Timeout. Specifies a non-zero timeout per retry attempt. In milliseconds (`Number`).

`retriable_status_codes` - (Optional) Status Code to Retry. HTTP status codes that should trigger a retry in addition to those specified by retry_on (`List`).

`retry_condition` - (Optional) Retry Condition. Specifies the conditions under which retry takes place. Retries can be on different types of condition depending on application requirements. For example, network failure, all 5xx response codes, idempotent 4xx response codes, etc The possible values are '5xx' : Retry will be done if the upstream server responds with any 5xx response code, or does not respond at all (disconnect/reset/read timeout). 'gateway-error' : Retry will be done only if the upstream server responds with 502, 503 or 504 responses (Included in 5xx) 'connect-failure' : Retry will be done if the request fails because of a connection failure to the upstream server (connect timeout, etc.). (Included in 5xx) 'refused-stream' : Retry is done if the upstream server resets the stream with a REFUSED_STREAM error code (Included in 5xx) 'retriable-4xx' : Retry is done if the upstream server responds with a retriable 4xx response code. The only response code in this category is HTTP CONFLICT (409) 'retriable-status-codes' : Retry is done if the upstream server responds with any response code matching one defined in retriable_status_codes field 'reset' : Retry is done if the upstream server does not respond at all (disconnect/reset/read timeout.) (`List`).

---

<a id="retry-policy-back-off"></a>

**Retry Policy Back Off**

`base_interval` - (Optional) Base Retry Interval. Specifies the base interval between retries in milliseconds (`Number`).

`max_interval` - (Optional) Maximum Retry Interval. Specifies the maximum interval between retries in milliseconds. This parameter is optional, but must be greater than or equal to the base_interval if set. The default is 10 times the base_interval (`Number`).

---

<a id="routes"></a>

**Routes**

`kind` - (Optional) Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route') (`String`).

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

`uid` - (Optional) UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid (`String`).

---

<a id="sensitive-data-policy"></a>

**Sensitive Data Policy**

`kind` - (Optional) Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route') (`String`).

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

`uid` - (Optional) UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid (`String`).

---

<a id="slow-ddos-mitigation"></a>

**Slow DDOS Mitigation**

`disable_request_timeout` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`request_headers_timeout` - (Optional) Request Headers Timeout. The amount of time the client has to send only the headers on the request stream before the stream is cancelled. The default value is 10000 milliseconds. This setting provides protection against Slowloris attacks (`Number`).

`request_timeout` - (Optional) Custom Timeout (`Number`).

---

<a id="timeouts"></a>

**Timeouts**

`create` - (Optional) A string that can be [parsed as a duration](`HTTPS://pkg.go.dev/time#ParseDuration`) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours) (`String`).

`delete` - (Optional) A string that can be [parsed as a duration](`HTTPS://pkg.go.dev/time#ParseDuration`) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are saved into state before the destroy operation occurs (`String`).

`read` - (Optional) A string that can be [parsed as a duration](`HTTPS://pkg.go.dev/time#ParseDuration`) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh is enabled (`String`).

`update` - (Optional) A string that can be [parsed as a duration](`HTTPS://pkg.go.dev/time#ParseDuration`) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours) (`String`).

---

<a id="tls-cert-params"></a>

**TLS Cert Params**

`certificates` - (Optional) Certificates. Set of certificates. See [Certificates](#tls-cert-params-certificates) below.

`cipher_suites` - (Optional) Cipher Suites. The following list specifies the supported cipher suite TLS_AES_128_GCM_SHA256 TLS_AES_256_GCM_SHA384 TLS_CHACHA20_POLY1305_SHA256 TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256 TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384 TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256 TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256 TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384 TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256 TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA TLS_RSA_WITH_AES_128_CBC_SHA TLS_RSA_WITH_AES_128_GCM_SHA256 TLS_RSA_WITH_AES_256_CBC_SHA TLS_RSA_WITH_AES_256_GCM_SHA384 If not specified, the default list: TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256 TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256 TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256 TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256 TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384 TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384 will be used (`List`).

`client_certificate_optional` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`client_certificate_required` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`maximum_protocol_version` - (Optional) TLS Protocol. TlsProtocol is enumeration of supported TLS versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are `TLS_AUTO`, `TLSv1_0`, `TLSv1_1`, `TLSv1_2`, `TLSv1_3`. Defaults to `TLS_AUTO` (`String`).

`minimum_protocol_version` - (Optional) TLS Protocol. TlsProtocol is enumeration of supported TLS versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are `TLS_AUTO`, `TLSv1_0`, `TLSv1_1`, `TLSv1_2`, `TLSv1_3`. Defaults to `TLS_AUTO` (`String`).

`no_client_certificate` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`validation_params` - (Optional) TLS Certificate Validation Parameters. This includes URL for a trust store, whether SAN verification is required and list of Subject Alt Names for verification. See [Validation Params](#tls-cert-params-validation-params) below.

`xfcc_header_elements` - (Optional) XFCC Header. X-Forwarded-Client-Cert header elements to be set in an mTLS enabled connections. If none are defined, the header will not be added (`List`).

---

<a id="tls-cert-params-certificates"></a>

**TLS Cert Params Certificates**

`kind` - (Optional) Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route') (`String`).

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

`uid` - (Optional) UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid (`String`).

---

<a id="tls-cert-params-validation-params"></a>

**TLS Cert Params Validation Params**

`skip_hostname_verification` - (Optional) Skip verification of hostname. When True, skip verification of hostname i.e. CN/Subject Alt Name of certificate is not matched to the connecting hostname (`Bool`).

`trusted_ca` - (Optional) Root CA Certificate Reference. Reference to Root CA Certificate. See [Trusted CA](#tls-cert-params-validation-params-trusted-ca) below.

`trusted_ca_url` - (Optional) Inline Root CA Certificate (legacy). Inline Root CA Certificate (`String`).

`verify_subject_alt_names` - (Optional) List of SANs for matching. List of acceptable Subject Alt Names/CN in the peer's certificate. When skip_hostname_verification is false and verify_subject_alt_names is empty, the hostname of the peer will be used for matching against SAN/CN of peer's certificate (`List`).

---

<a id="tls-cert-params-validation-params-trusted-ca"></a>

**TLS Cert Params Validation Params Trusted CA**

`trusted_ca_list` - (Optional) Root CA Certificate Reference. Reference to Root CA Certificate (`Block`).

---

<a id="tls-parameters"></a>

**TLS Parameters**

`client_certificate_optional` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`client_certificate_required` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`common_params` - (Optional) TLS Parameters. Information of different aspects for TLS authentication related to ciphers, certificates and trust store. See [Common Params](#tls-parameters-common-params) below.

`no_client_certificate` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`xfcc_header_elements` - (Optional) XFCC Header. X-Forwarded-Client-Cert header elements to be set in an mTLS enabled connections. If none are defined, the header will not be added (`List`).

---

<a id="tls-parameters-common-params"></a>

**TLS Parameters Common Params**

`cipher_suites` - (Optional) Cipher Suites. The following list specifies the supported cipher suite TLS_AES_128_GCM_SHA256 TLS_AES_256_GCM_SHA384 TLS_CHACHA20_POLY1305_SHA256 TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256 TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384 TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256 TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256 TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384 TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256 TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA TLS_RSA_WITH_AES_128_CBC_SHA TLS_RSA_WITH_AES_128_GCM_SHA256 TLS_RSA_WITH_AES_256_CBC_SHA TLS_RSA_WITH_AES_256_GCM_SHA384 If not specified, the default list: TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256 TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256 TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256 TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256 TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384 TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384 will be used (`List`).

`maximum_protocol_version` - (Optional) TLS Protocol. TlsProtocol is enumeration of supported TLS versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are `TLS_AUTO`, `TLSv1_0`, `TLSv1_1`, `TLSv1_2`, `TLSv1_3`. Defaults to `TLS_AUTO` (`String`).

`minimum_protocol_version` - (Optional) TLS Protocol. TlsProtocol is enumeration of supported TLS versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are `TLS_AUTO`, `TLSv1_0`, `TLSv1_1`, `TLSv1_2`, `TLSv1_3`. Defaults to `TLS_AUTO` (`String`).

`tls_certificates` - (Optional) TLS Certificates. Set of TLS certificates. See [TLS Certificates](#tls-parameters-common-params-tls-certificates) below.

`validation_params` - (Optional) TLS Certificate Validation Parameters. This includes URL for a trust store, whether SAN verification is required and list of Subject Alt Names for verification. See [Validation Params](#tls-parameters-common-params-validation-params) below.

---

<a id="tls-parameters-common-params-tls-certificates"></a>

**TLS Parameters Common Params TLS Certificates**

`certificate_url` - (Optional) Certificate. TLS certificate. Certificate or certificate chain in PEM format including the PEM headers (`String`).

`custom_hash_algorithms` - (Optional) Hash Algorithms. Specifies the hash algorithms to be used (`Block`).

`description` - (Optional) Description. Description for the certificate (`String`).

`disable_ocsp_stapling` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`private_key` - (Optional) Secret. SecretType is used in an object to indicate a sensitive/confidential field (`Block`).

`use_system_defaults` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

---

<a id="tls-parameters-common-params-validation-params"></a>

**TLS Parameters Common Params Validation Params**

`skip_hostname_verification` - (Optional) Skip verification of hostname. When True, skip verification of hostname i.e. CN/Subject Alt Name of certificate is not matched to the connecting hostname (`Bool`).

`trusted_ca` - (Optional) Root CA Certificate Reference. Reference to Root CA Certificate (`Block`).

`trusted_ca_url` - (Optional) Inline Root CA Certificate (legacy). Inline Root CA Certificate (`String`).

`verify_subject_alt_names` - (Optional) List of SANs for matching. List of acceptable Subject Alt Names/CN in the peer's certificate. When skip_hostname_verification is false and verify_subject_alt_names is empty, the hostname of the peer will be used for matching against SAN/CN of peer's certificate (`List`).

---

<a id="user-identification"></a>

**User Identification**

`kind` - (Optional) Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route') (`String`).

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

`uid` - (Optional) UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid (`String`).

---

<a id="waf-type"></a>

**WAF Type**

`app_firewall` - (Optional) App Firewall Reference. A list of references to the app_firewall configuration objects. See [App Firewall](#waf-type-app-firewall) below.

`disable_waf` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`inherit_waf` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

---

<a id="waf-type-app-firewall"></a>

**WAF Type App Firewall**

`app_firewall` - (Optional) Application Firewall. References to an Application Firewall configuration object. See [App Firewall](#waf-type-app-firewall-app-firewall) below.

---

<a id="waf-type-app-firewall-app-firewall"></a>

**WAF Type App Firewall App Firewall**

`kind` - (Optional) Kind. When a configuration object(e.g. virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route') (`String`).

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

`uid` - (Optional) UID. When a configuration object(e.g. virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. route's) uid (`String`).

---
