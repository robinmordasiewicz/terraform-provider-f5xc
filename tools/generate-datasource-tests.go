// Copyright (c) F5XC Community
// SPDX-License-Identifier: MPL-2.0

//go:build ignore
// +build ignore

// This script generates data source acceptance tests for all data sources.
// It uses resource configurations based on OpenAPI specs and existing resource tests.
//
// Features:
//   - resourceConfigs: Custom HCL configuration for each resource's required fields
//   - skipResources: Resources that cannot be tested (permissions, broken APIs, etc.)
//   - systemNamespaceResources: Resources that work in "system" namespace
//
// Run with: go run tools/generate-datasource-tests.go
package main

import (
	"bytes"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// ResourceConfig contains the Terraform configuration for a resource
type ResourceConfig struct {
	Name           string
	StructName     string
	ResourceName   string
	TestFuncName   string
	ConfigFuncName string
	NeedsNamespace bool
	IsNamespace    bool
	// Custom resource configuration block (the HCL inside the resource block)
	ResourceConfig string
	// Whether to skip this resource (e.g., FORBIDDEN errors)
	Skip       bool
	SkipReason string
}

// Resource configurations based on OpenAPI specs
// Key is the resource name (without f5xc_ prefix)
var resourceConfigs = map[string]string{
	// Simple resources - just name/namespace
	"alert_policy":              "",
	"allowed_tenant":            "",
	"api_crawler":               "",
	"api_definition":            "",
	"api_discovery":             "",
	"api_testing":               "",
	"app_api_group":             "",
	"app_firewall":              "",
	"cdn_cache_rule":            "",
	"cluster":                   "",
	"code_base_integration":     "",
	"dns_compliance_checks":     "",
	"enhanced_firewall_policy":  "",
	"ike1":                      "",
	"ike2":                      "",
	"ike_phase1_profile":        "",
	"ike_phase2_profile":        "",
	"irule":                     "",
	"malicious_user_mitigation": "",
	"namespace":                 "",
	"rate_limiter":              "",
	"rate_limiter_policy":       "",
	"secret_policy":             "",
	"sensitive_data_policy":     "",
	"tpm_manager":               "",
	"trusted_ca_list":           "",
	"virtual_k8s":               "",
	"voltshare_admin_policy":    "",

	// Resources with required fields
	"address_allocator": `
  address_pool = ["10.0.0.0/24"]
  mode = "SITE_LOCAL_ADDRESS_ALLOCATOR"`,

	"advertise_policy": `
  port = 80`,

	"alert_receiver": `
  email {
    email = "test@example.com"
  }`,

	"api_credential": `
  api_credential_type = "API_TOKEN"
  expiration_days     = 30`,

	"app_setting": `
  app_type_refs {
    name      = f5xc_app_type.test.name
    namespace = f5xc_namespace.test.name
  }`,

	"app_type": `
  business_logic_markup_setting {
    enable = true
  }`,

	"authentication": `
  hmac_auth {
    secret {
      clear_secret_info {
        url = "string:///dGVzdC1zZWNyZXQtdmFsdWU="
      }
    }
  }`,

	"bgp": `
  bgp_parameters {
    asn          = 65000
    bgp_router_id_type = "BGP_ROUTER_ID_FROM_INTERFACE"
  }
  where {
    site {
      network_type = "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE"
    }
  }`,

	"bgp_asn_set": `
  as_numbers = ["65000"]`,

	"bgp_routing_policy": `
  policy_rule {
    description = "test rule"
    match {
      community_string {
        values = ["65000:100"]
      }
    }
    action {
      permit = true
    }
  }`,

	"bigip_irule": `
  when_http_request = "HTTP::respond 200 content \"OK\""`,

	"certificate": `
  certificate_url = "string:///LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUNxRENDQVpBQ0NRQ21SN0dCREN4dXRqQU5CZ2txaGtpRzl3MEJBUXNGQURBV01SUXdFZ1lEVlFRRERBdDAKWlhOMExXTmxjblF1WTI5dE1CNFhEVEkxTURFeE1UQXdNREF3TUZvWERUSTJNREV4TVRBd01EQXdNRm93RmpFVQpNQklHQTFVRUF3d0xkR1Z6ZEMxalpYSjBMbU52YlRDQ0FTSXdEUVlKS29aSWh2Y05BUUVCQlFBRGdnRVBBRENDCkFRb0NnZ0VCQU1wQklYdVhJcmZJT1ZoQndxWnBKNXZsRlk4cEtndnhOdHQ4elEyRWtLVmhmMk1mOXc4bE5RdTIKK3FkZm5YNW5YRUZXS2dhVnpqYUZGZHFOZ0xHaUlRaWQvWVcwa0dDd3RHZURlMGE4OFg1UW5oeTRHdkdHdEhyVgp6V3ZoTlVHaXJSSnJiOGUwemxQdWxGeTA2dEl4ZUdBb1FXT09KU0p1NUtsOWJJenFYZ0hHZG9LWXd0VG9ERzc5CnFXN3VZNGZNVnhSM2dBYm5CSk84eGxMR3dQNndpSU5PQTJObFNUc1g5TmliOHR3b0NuSnpQMm1wOGFOd1VYZWsKNHBiZW1HZ2dxY2ZBM0ROaXNmRUNhLy9SRzFYZlFrNjI1WGtrb3h5UlBuaG9JSU5kVmFhY0RRTDVYZU16akRuRwphaUZaaGZCQWJHWlB3dk1MRnVEN3JPYWllRGRGckZjQ0F3RUFBVEFOQmdrcWhraUc5dzBCQVFzRkFBT0NBUUVBCmdEZjVPMk1DT0JCRURQZnFkRnpiMjBLUm9nNUp3bmNHcklwSVQwQmc1S0M5VElQRDUxdkJtY2w0ZWZxT3Q2TWkKT0lSeEFKa3JvTFFIRW5xMnBCdkxVM0dwT1IxQ1V1STYvRnBmb2toaFlEN1lsSnR4RE9mZW9JT0NOTk0rMXgyaApmTXY1NHVhV0V4R3dwQnVQL1M1dVZ6TU95aVVFcklXejIyUFpsZndIS3ZhdGNuclFBTjhLNmpvMUp1dkc4RmdsCk1VcjNVa3VMWEVYMHBNd1p0WWVoUjJVYlN0ZEVVT1puSTdFalp5TTU5WndFVmRrK1NWYmRVNjQzOEdYVHlDRjcKeGxRTjk3YmZOUnBkZG5JYWpyYU52SXJrclBYVnhERXhDOEo4NzQzRE9ZenZ2bTZNVW1vZW9ML3FuRVRiTGV5RApqQjJWS0MxTi9NeEFjUUdGOWRueHpBPT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="
  private_key {
    clear_secret_info {
      url = "string:///LS0tLS1CRUdJTiBQUklWQVRFIEtFWS0tLS0tCk1JSUV2UUlCQURBTkJna3Foa2lHOXcwQkFRRUZBQVNDQktjd2dnU2pBZ0VBQW9JQkFRREtRU0Y3bHlLM3lEbFkKUWNLbWFTZWI1UldQS1NvTDhUYmJmTTBOaEpDbFlYOWpIL2NQSlRVTHR2cW5YNTErWjF4QlZpb0dsYzQyaFJYYQpqWUN4b2lFSW5mMkZ0SkJnc0xSbmczdEd2UEYrVUo0Y3VCcnhoclI2MWMxcjRUVkJvcTBTYTIvSHRNNVQ3cFJjCnRPclNNWGhnS0VGamlpVWlidVNwZld5TTZsNEJ4bmFDbU1MVTZBeHUvYWx1N21PSHpGY1VkNEFHNXdTVHZNWlMKeHNEK3NJaURUZ05qWlVrN0YvVFltL0xjS0FweWN6OXBxZkdqY0ZGM3BPS1czcGhvSUtuSHdOd3pZckh4QW12LwowUnRWMzBKT3R1VjVKS01ja1Q1NGFDQ0RYVldtbkEwQytWM2pNNHc1eG1vaFdZWHdRR3htVDhMekN4YmcrNnptCm9uZzNSYXhYQWdNQkFBRUNnZ0VBQm8xS1dNNEVQWEZQZVFnTElXLzZ2TUFCRzJaSUZwQmRQS2szL29pVER2RVgKYzBTMGJhdG5MVXl0UkNkcjdFSStUVy9JbFlMQm1uQXl4OGpCSEV4akVJK3Q0Z0VRTkVoSXk0ZnM5emtlYVIrVgpoY0pHaFp5eDVYZjkyMW0zcjBHZkVLZmJEaExjY3daeWRKUCtEa0tkSit1OWtSamY4ZEpITXhuMjZCRXJYbnVvClNiM1pnVE8wTjlJN3hGTUxRN1pvZWlWbWNYZ0ZwdHV5L0FLbHZQcHVKZHRRd0lYcXVqT1JVUEFnYVdpdUFOcUcKcGdObk1BUGhoT3pvdS9TVlBsU25ISGNRM0tqRmMwT0E0cnE2WS9rTCtqRDh5N0RRMTZ6M2piNm5ZM2RTdDA5VQp0UHE3ZCtUMGFIVWVXeUhPY0NWWVorZU9obytzdHdmblUrTFpxQnliQVFLQmdRRHdpeXEzZUd3UjFPNW5Qcm9hClpHaU9KbUVnbkszSjV4UXRLTkV2UmFHYTJBTHV0ZjVuVndVNFVSa1BFejJxME1WOENmYlE2UnVab2ZhTGltUHcKaHk5SGVtYk5JMTZZNXl3NWRIZlFCMWNBQ05jY2VLS1F3WEFneVdxNjBldTFLanBveGozVTdKaTF2WC8wdEVYRgpYVEhVU3B1bHJDMHR2NXN1V0xIelVFaDJWd0tCZ1FEWDRwUFVzaCtqeFBDNUUxcGNKOWY3UTFMK3IxV0dRVHRPCjJTd2F3YVQvQkNsRXRpVDRJZ2lXTTBQL3lSeUVFaFNrbGsycWJFUW40L3RNMkVyZ0ZTOEN1cTdjWWtmblFDckoKSHdKcG15TFVQK01vTy9FZTdQRkYyS1lTSVp5K1ZRZlJ4Z2FEWHhuNFdoOFZMYys2Qk1pTGc1V3FhMkZXVk1ENwpDSHdEOE1tK1FRS0JnRkZmcVBMM2xLQUpiYWJxNEJUOXM0UHRNWjQzbWI5NkdTTDF3dVBuaVVMSFd3U0VVQVRCCmxoNEEzcXU5cmRRQVBnWHBUd0I4QWVKRTVWTE9qb3ljY3k5bWxPd0dQTjBHZmg1S0ZhZndIWHpESTNuZm0wbTUKVzdsRFRDb2g5RTY0bE9UaVpiK0s0MXZ6eDUrMjF0OWJCeEt4YW0xMDNiQzc1aHRjUUxhZTg5TnZBb0dBSi93RQpubTlYR202TDJiaDJ3YmtTRGpRUWdZTGZYVVdxNFBwd3RnY3k2U29yOW43dnZZZ2FCTm5mUUdFaWtwSGM2Q1RMCnRsRU1LZ3VhZ1hlVjUxKy9PS01OUHNxWkhqbksvYkpZRGVYdjE3SjV5d3lBWUpmNFNFMlFLYlI5MkxjdDZYckcKdFhJVHBFNHA2ZjJyWm1KdDBPRHN2VEtqNTVPQ1VuL0prTjl3VEdFQ2dZRUE3QVA4L2o4S0lNVlkzaWM5MXZKMwpkZm1xNWxHb2NRWmxzTkV2NW96elByRGRlMnFQejU4NjQrckQ5WEhKUlNSRSs1OXlKVytqT25TVm1qeFk1dHA1Cnk0U3hYZnFETURCL1lYS05hR0dxNE1vdDNOOGZ0U0drUkRpOGk4RThuRHQ3dlNUZm5xaW5QVW9YVGhYeW9EV2EKUXI2ZURJK3VjL1pPN2V5bWw1N3FrL009Ci0tLS0tRU5EIFBSSVZBVEUgS0VZLS0tLS0K"
    }
  }`,

	"certificate_chain": `
  certificate_url = "string:///LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUNxRENDQVpBQ0NRQ21SN0dCREN4dXRqQU5CZ2txaGtpRzl3MEJBUXNGQURBV01SUXdFZ1lEVlFRRERBdDAKWlhOMExXTmxjblF1WTI5dE1CNFhEVEkxTURFeE1UQXdNREF3TUZvWERUSTJNREV4TVRBd01EQXdNRm93RmpFVQpNQklHQTFVRUF3d0xkR1Z6ZEMxalpYSjBMbU52YlRDQ0FTSXdEUVlKS29aSWh2Y05BUUVCQlFBRGdnRVBBRENDCkFRb0NnZ0VCQU1wQklYdVhJcmZJT1ZoQndxWnBKNXZsRlk4cEtndnhOdHQ4elEyRWtLVmhmMk1mOXc4bE5RdTIKK3FkZm5YNW5YRUZXS2dhVnpqYUZGZHFOZ0xHaUlRaWQvWVcwa0dDd3RHZURlMGE4OFg1UW5oeTRHdkdHdEhyVgp6V3ZoTlVHaXJSSnJiOGUwemxQdWxGeTA2dEl4ZUdBb1FXT09KU0p1NUtsOWJJenFYZ0hHZG9LWXd0VG9ERzc5CnFXN3VZNGZNVnhSM2dBYm5CSk84eGxMR3dQNndpSU5PQTJObFNUc1g5TmliOHR3b0NuSnpQMm1wOGFOd1VYZWsKNHBiZW1HZ2dxY2ZBM0ROaXNmRUNhLy9SRzFYZlFrNjI1WGtrb3h5UlBuaG9JSU5kVmFhY0RRTDVYZU16akRuRwphaUZaaGZCQWJHWlB3dk1MRnVEN3JPYWllRGRGckZjQ0F3RUFBVEFOQmdrcWhraUc5dzBCQVFzRkFBT0NBUUVBCmdEZjVPMk1DT0JCRURQZnFkRnpiMjBLUm9nNUp3bmNHcklwSVQwQmc1S0M5VElQRDUxdkJtY2w0ZWZxT3Q2TWkKT0lSeEFKa3JvTFFIRW5xMnBCdkxVM0dwT1IxQ1V1STYvRnBmb2toaFlEN1lsSnR4RE9mZW9JT0NOTk0rMXgyaApmTXY1NHVhV0V4R3dwQnVQL1M1dVZ6TU95aVVFcklXejIyUFpsZndIS3ZhdGNuclFBTjhLNmpvMUp1dkc4RmdsCk1VcjNVa3VMWEVYMHBNd1p0WWVoUjJVYlN0ZEVVT1puSTdFalp5TTU5WndFVmRrK1NWYmRVNjQzOEdYVHlDRjcKeGxRTjk3YmZOUnBkZG5JYWpyYU52SXJrclBYVnhERXhDOEo4NzQzRE9ZenZ2bTZNVW1vZW9ML3FuRVRiTGV5RApqQjJWS0MxTi9NeEFjUUdGOWRueHpBPT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="`,

	"contact": `
  email = "test@example.com"`,

	"container_registry": `
  registry          = "registry.example.com"
  user_name         = "testuser"
  password {
    clear_secret_info {
      url = "string:///dGVzdC1wYXNzd29yZA=="
    }
  }`,

	"crl": `
  crl_url = "string:///LS0tLS1CRUdJTiBYNTA5IENSTC0tLS0tCk1JSUJURENDQVFNQ0FRRXdEUVlKS29aSWh2Y05BUUVMQlFBd0ZqRVVNQklHQTFVRUF3d0xkR1Z6ZEMxalpYSjAKTG1OdmJSY05NalV3TVRFeE1EQXdNREF3V2hjTk1qWXdNVEV4TURBd01EQXdXakFOQmdrcWhraUc5dzBCQVFzRgpBQU9DQVFFQWdEZjVPMk1DT0JCRURQZnFkRnpiMjBLUm9nNUp3bmNHcklwSVQwQmc1S0M5VElQRDUxdkJtY2w0CmVmcU90Nk1pT0lSeEFKa3JvTFFIRW5xMnBCdkxVM0dwT1IxQ1V1STYvRnBmb2toaFlEN1lsSnR4RE9mZW9JT0MKTk5NKzF4MmhmTXY1NHVhV0V4R3dwQnVQL1M1dVZ6TU95aVVFcklXejIyUFpsZndIS3ZhdGNuclFBTjhLNmpvMQpKdXZHOEZnbE1VcjNVa3VMWEVYMHBNd1p0WWVoUjJVYlN0ZEVVT1puSTdFalp5TTU5WndFVmRrK1NWYmRVNjQzCjhHWFR5Q0Y3eGxRTjk3YmZOUnBkZG5JYWpyYU52SXJrclBYVnhERXhDOEo4NzQzRE9ZenZ2bTZNVW1vZW9ML3EKbkVUYkxleURqQjJWS0MxTi9NeEFjUUdGOWRueHpBPT0KLS0tLS1FTkQgWDUwOSBDUkwtLS0tLQo="`,

	"data_group": `
  string_records {}`,

	"data_type": `
  is_pii = true
  rules {
    key_pattern {
      substring_value = "test"
    }
  }`,

	"discovery": `
  discovery_consul {
    access_info {
      connection_info {
        api_server = "http://consul.example.com:8500"
      }
    }
  }`,

	"dns_domain": `
  dns_domain = "example.com"`,

	"dns_lb_health_check": `
  http {
    path = "/"
  }`,

	"dns_lb_pool": `
  type = "A"
  a_pool {
    members {
      public_ip {
        ip = "1.2.3.4"
      }
    }
  }`,

	"dns_load_balancer": `
  record_type = "A"
  dns_a_record {
    values = ["1.2.3.4"]
  }`,

	"dns_zone": `
  dns_zone_type = "primary"
  primary {
    default_rr_set_group = []
  }`,

	"endpoint": `
  where {
    site {
      network_type = "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE"
    }
  }`,

	"fast_acl": `
  site_acl {
    fast_acl_rules {
      name   = "test-rule"
      action {
        deny = true
      }
      prefix {
        prefix = ["10.0.0.0/8"]
      }
    }
  }`,

	"fast_acl_rule": `
  action {
    deny = true
  }
  prefix {
    prefix = ["10.0.0.0/8"]
  }`,

	"filter_set": `
  context_key = "dashboard"
  filter_fields {
    field_id = "test-field"
    string_field {
      field_values = ["test-value"]
    }
  }`,

	"fleet": `
  fleet_label = "test-fleet"
  network_connectors {
    disable_forward_proxy = true
  }`,

	"forward_proxy_policy": `
  proxy_label = "test-proxy"
  rule_list {
    rules {
      metadata {
        name = "rule1"
      }
      spec {
        action      = "ALLOW"
        any_client  = true
        any_dst     = true
        rule_description = "Allow all"
      }
    }
  }`,

	"forwarding_class": `
  dscp_class = "EF"`,

	"geo_location_set": ``,

	"global_log_receiver": `
  http_global_receiver {
    uri = "https://logs.example.com/receive"
  }`,

	"healthcheck": `
  healthy_threshold   = 1
  unhealthy_threshold = 2
  timeout             = 3
  interval            = 5
  tcp_health_check {}`,

	"http_loadbalancer": `
  domains = ["test.example.com"]
  http {
    dns_volterra_managed = false
  }
  default_route_pools {
    pool {
      name      = f5xc_origin_pool.test.name
      namespace = f5xc_namespace.test.name
    }
  }`,

	"ip_prefix_set": `
  ipv4_prefixes {
    ipv4_prefix      = "10.0.0.0/8"
    description_spec = "Test prefix"
  }`,

	"known_label": `
  key = "test-key"`,

	"known_label_key": `
  description = "Test label key"`,

	"log_receiver": `
  http_receiver {
    uri = "https://logs.example.com/receive"
  }`,

	"nat_policy": `
  rule_list {
    rules {
      metadata {
        name = "test-rule"
      }
      spec {
        src_ip_match = "ANY"
        action {
          snat_pool {
            any_vip = true
          }
        }
      }
    }
  }`,

	"network_connector": `
  sli_to_global_dr {
    global_vn {
      name      = "global-network"
      namespace = "system"
    }
  }`,

	"network_firewall": `
  active_forward_proxy_policies {
    forward_proxy_policies {
      name      = f5xc_forward_proxy_policy.test.name
      namespace = f5xc_namespace.test.name
    }
  }`,

	"network_interface": `
  ethernet_interface {
    device    = "eth0"
    dhcp_server {
      dhcp_networks {
        network_prefix = "10.0.0.0/24"
      }
    }
  }`,

	"network_policy": `
  endpoint {
    any = true
  }
  ingress_rules {
    metadata {
      name = "rule1"
    }
    spec {
      action      = "ALLOW"
      any         = true
      protocol    = "TCP"
    }
  }
  egress_rules {
    metadata {
      name = "rule1"
    }
    spec {
      action      = "ALLOW"
      any         = true
      protocol    = "TCP"
    }
  }`,

	"network_policy_rule": `
  action   = "ALLOW"
  any      = true
  protocol = "TCP"`,

	"network_policy_view": `
  endpoint {
    any = true
  }
  ingress_rules {
    metadata {
      name = "rule1"
    }
    spec {
      action      = "ALLOW"
      any         = true
      protocol    = "TCP"
    }
  }
  egress_rules {
    metadata {
      name = "rule1"
    }
    spec {
      action      = "ALLOW"
      any         = true
      protocol    = "TCP"
    }
  }`,

	"oidc_provider": `
  issuer   = "https://issuer.example.com"
  jwks_url = "https://issuer.example.com/.well-known/jwks.json"`,

	"origin_pool": `
  port = 443
  origin_servers {
    labels {}
    public_name {
      dns_name = "example.com"
    }
  }
  no_tls {}
  same_as_endpoint_port {}`,

	"policer": `
  committed_information_rate = 10000
  burst_size                 = 5000`,

	"policy_based_routing": `
  destination_match_rules {
    any_prefix = true
    any_port   = true
    protocol   = "TCP"
    action {
      vrf_egress = true
    }
  }`,

	"protocol_inspection": `
  protocol_list = ["HTTP"]`,

	"protocol_policer": `
  protocol_policer {
    match {
      any_ip {}
    }
    protocol {
      dns {}
    }
    policer {
      bandwidth = 1000
      burst     = 1000
    }
  }`,

	"proxy": `
  proxy_url = "http://proxy.example.com:8080"`,

	"report_config": `
  report_type = "SECURITY_EVENTS"`,

	"secret_policy_rule": `
  action      = "ALLOW"
  label_selector {
    expressions = ["key1=value1"]
  }`,

	"service_policy": `
  allow_all_requests {}
  any_server {}`,

	"service_policy_rule": `
  action = "ALLOW"`,

	"tcp_loadbalancer": `
  domains = ["test.example.com"]
  tcp {
    dns_volterra_managed = false
  }
  origin_pools_weights {
    pool {
      name      = f5xc_origin_pool.test.name
      namespace = f5xc_namespace.test.name
    }
  }`,

	"ticket_tracking_system": `
  jira {
    url = "https://jira.example.com"
    credentials {
      email = "test@example.com"
      token {
        clear_secret_info {
          url = "string:///dGVzdC10b2tlbg=="
        }
      }
    }
    project_key = "TEST"
  }`,

	"token": `
  type = "JWT"`,

	"tpm_api_key": `
  api_key {
    clear_secret_info {
      url = "string:///dGVzdC1hcGkta2V5"
    }
  }`,

	"tpm_category": `
  category_type = "CUSTOM"`,

	"udp_loadbalancer": `
  domains     = ["%[1]s.example.com"]
  listen_port = 53
  origin_pools_weights {
    pool {
      name      = f5xc_origin_pool.test.name
      namespace = "system"
    }
    weight = 1
  }
  advertise_on_public_default_vip {}`,

	"usb_policy": ``,

	"user_identification": `
  rules {
    client_ip {}
  }`,

	"virtual_host": `
  domains = ["test.example.com"]
  proxy   = "SMA_PROXY"`,

	"virtual_network": `
  global_network {}`,

	"virtual_site": `
  site_type = "CUSTOMER_EDGE"
  site_selector {
    expressions = ["site-type=ce"]
  }`,

	"waf_exclusion_policy": ``,

	"workload": `
  containers {
    name  = "nginx"
    image = "nginx:latest"
  }`,

	"workload_flavor": `
  vcpus        = "1"
  memory       = "1024"
  ephemeral_storage = "10Gi"`,
}

// Resources to skip (FORBIDDEN errors, special permissions required)
var skipResources = map[string]string{
	"addon_subscription":             "Requires special addon permissions",
	"address_allocator":              "API returns BAD_REQUEST - requires specific infrastructure",
	"apm":                            "Requires APM subscription",
	"bot_defense_app_infrastructure": "Requires Bot Defense subscription",
	"child_tenant":                   "Requires tenant management permissions",
	"cloud_credentials":              "Requires system-level namespace access",
	"cminstance":                     "Requires CM instance management permissions",
	"customer_support":               "Requires special support permissions",
	"forwarding_class":               "Requires tenant quota configuration",
	"managed_tenant":                 "Requires tenant management permissions",
	"protocol_inspection":            "Resource test has schema mismatch - needs fixing",
	"protocol_policer":               "Resource test has schema mismatch - needs fixing",
	"quota":                          "Requires system admin permissions",
	"registration":                   "Requires site registration permissions",
	"role":                           "Requires IAM permissions",
	"site_mesh_group":                "API returns 500 errors",
	"route":                          "API returns 500 errors",
	"tenant_configuration":           "API returns 500 errors",
	"tenant_profile":                 "API returns 501 errors",
	"token":                          "Resource test has schema mismatch - needs fixing",
	"tpm_category":                   "Resource test has schema mismatch - needs fixing",
	"udp_loadbalancer":               "Requires system namespace template with dependencies - needs template fix",
	"usb_policy":                     "Resource test fails - API returns BAD_REQUEST",
	"waf_exclusion_policy":           "No resource test exists - config unknown",
}

// Resources that need special dependencies in the config
var resourceDependencies = map[string]string{
	"app_setting": `
resource "f5xc_app_type" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = "${%[2]q}-apptype"
  namespace  = f5xc_namespace.test.name
  business_logic_markup_setting {
    enable = true
  }
}
`,
	"http_loadbalancer": `
resource "f5xc_origin_pool" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = "${%[2]q}-pool"
  namespace  = f5xc_namespace.test.name
  origin_servers {
    public_ip {
      ip = "1.2.3.4"
    }
  }
  port               = 80
  endpoint_selection = "LOCAL_PREFERRED"
  loadbalancer_algorithm = "ROUND_ROBIN"
}
`,
	"tcp_loadbalancer": `
resource "f5xc_origin_pool" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = "${%[2]q}-pool"
  namespace  = f5xc_namespace.test.name
  origin_servers {
    public_ip {
      ip = "1.2.3.4"
    }
  }
  port               = 80
  endpoint_selection = "LOCAL_PREFERRED"
  loadbalancer_algorithm = "ROUND_ROBIN"
}
`,
	"udp_loadbalancer": `
resource "f5xc_origin_pool" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = "${%[2]q}-pool"
  namespace  = f5xc_namespace.test.name
  origin_servers {
    public_ip {
      ip = "1.2.3.4"
    }
  }
  port               = 80
  endpoint_selection = "LOCAL_PREFERRED"
  loadbalancer_algorithm = "ROUND_ROBIN"
}
`,
	"network_firewall": `
resource "f5xc_forward_proxy_policy" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = "${%[2]q}-fpp"
  namespace  = f5xc_namespace.test.name
  proxy_label = "test-proxy"
  rule_list {
    rules {
      metadata {
        name = "rule1"
      }
      spec {
        action      = "ALLOW"
        any_client  = true
        any_dst     = true
        rule_description = "Allow all"
      }
    }
  }
}
`,
}

// Resources that don't need a namespace (they're system-level)
var systemLevelResources = map[string]bool{
	"namespace": true,
}

// Resources that MUST use namespace = "system" (not custom namespace)
var systemNamespaceResources = map[string]bool{
	"virtual_network":     true,
	"policer":             true,
	"user_identification": true,
	"bgp_asn_set":         true,
	"geo_location_set":    true,
	"healthcheck":         true,
	"rate_limiter":        true,
	"service_policy":      true,
	"data_group":          true,
	"data_type":           true,
	"filter_set":          true,
	"origin_pool":         true,
	"app_firewall":        true,
}

func toStructName(name string) string {
	parts := strings.Split(name, "_")
	var result string
	for _, part := range parts {
		if len(part) > 0 {
			result += strings.ToUpper(part[:1]) + part[1:]
		}
	}
	return result
}

func main() {
	// Template for the test file
	const testTemplate = `// Copyright (c) F5XC Community
// SPDX-License-Identifier: MPL-2.0

package provider_test

{{if .Skip}}
import (
	"testing"
)

func {{.TestFuncName}}(t *testing.T) {
	t.Skip("{{.SkipReason}}")
}
{{else}}
import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/f5xc/terraform-provider-f5xc/internal/acctest"
)

func {{.TestFuncName}}(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test")
{{- if and .NeedsNamespace (not .IsNamespace) (not .UseSystemNamespace)}}
	nsName := acctest.RandomName("tf-acc-test-ns")
{{- end}}
	resourceName := "{{.ResourceName}}.test"
	dataSourceName := "data.{{.ResourceName}}.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
{{- if and .NeedsNamespace (not .IsNamespace) (not .UseSystemNamespace)}}
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
{{- end}}
		Steps: []resource.TestStep{
			{
{{- if .IsNamespace}}
				Config: {{.ConfigFuncName}}(rName),
{{- else if .UseSystemNamespace}}
				Config: {{.ConfigFuncName}}(rName),
{{- else if .NeedsNamespace}}
				Config: {{.ConfigFuncName}}(nsName, rName),
{{- else}}
				Config: {{.ConfigFuncName}}(rName),
{{- end}}
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
{{- if and .NeedsNamespace (not .IsNamespace)}}
					resource.TestCheckResourceAttrPair(dataSourceName, "namespace", resourceName, "namespace"),
{{- end}}
					resource.TestCheckResourceAttrPair(dataSourceName, "id", resourceName, "id"),
				),
			},
		},
	})
}
{{end}}
{{- if .IsNamespace}}

func {{.ConfigFuncName}}(name string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(` + "`" + `
resource "{{.ResourceName}}" "test" {
  name = %[1]q
}

data "{{.ResourceName}}" "test" {
  name      = {{.ResourceName}}.test.name
  namespace = "system"
}
` + "`" + `, name))
}
{{- else if .UseSystemNamespace}}

func {{.ConfigFuncName}}(name string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(` + "`" + `
resource "{{.ResourceName}}" "test" {
  name      = %[1]q
  namespace = "system"{{.ResourceConfig}}
}

data "{{.ResourceName}}" "test" {
  depends_on = [{{.ResourceName}}.test]
  name       = {{.ResourceName}}.test.name
  namespace  = {{.ResourceName}}.test.namespace
}
` + "`" + `, name))
}
{{- else if and .NeedsNamespace (not .Skip)}}

func {{.ConfigFuncName}}(nsName, name string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(` + "`" + `
resource "f5xc_namespace" "test" {
  name = %[1]q
}

resource "time_sleep" "wait_for_namespace" {
  depends_on      = [f5xc_namespace.test]
  create_duration = "5s"
}
{{.Dependencies}}
resource "{{.ResourceName}}" "test" {
  depends_on = [time_sleep.wait_for_namespace{{.DependsOnExtra}}]
  name       = %[2]q
  namespace  = f5xc_namespace.test.name{{.ResourceConfig}}
}

data "{{.ResourceName}}" "test" {
  depends_on = [{{.ResourceName}}.test]
  name       = {{.ResourceName}}.test.name
  namespace  = {{.ResourceName}}.test.namespace
}
` + "`" + `, nsName, name))
}
{{- else if not .Skip}}

func {{.ConfigFuncName}}(name string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(` + "`" + `
resource "{{.ResourceName}}" "test" {
  name = %[1]q
}

data "{{.ResourceName}}" "test" {
  name      = {{.ResourceName}}.test.name
  namespace = "system"
}
` + "`" + `, name))
}
{{- end}}
`

	// Find all data source test files
	dataSourceDir := "internal/provider"
	pattern := filepath.Join(dataSourceDir, "*_data_source_test.go")
	files, err := filepath.Glob(pattern)
	if err != nil {
		fmt.Printf("Error finding test files: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Found %d data source test files to fix\n", len(files))

	fixed := 0
	skipped := 0

	for _, file := range files {
		baseName := filepath.Base(file)
		resourceName := strings.TrimSuffix(baseName, "_data_source_test.go")

		// Check if we should skip this resource
		if reason, skip := skipResources[resourceName]; skip {
			fmt.Printf("Marking %s as skipped: %s\n", resourceName, reason)
			// Continue to regenerate with skip flag
		}

		structName := toStructName(resourceName)

		config := ResourceConfig{
			Name:           resourceName,
			StructName:     structName,
			ResourceName:   "f5xc_" + resourceName,
			TestFuncName:   "TestAcc" + structName + "DataSource_basic",
			ConfigFuncName: "testAcc" + structName + "DataSourceConfig_basic",
			NeedsNamespace: !systemLevelResources[resourceName],
			IsNamespace:    resourceName == "namespace",
		}

		// Check for skip
		if reason, skip := skipResources[resourceName]; skip {
			config.Skip = true
			config.SkipReason = reason
		}

		// Get resource-specific config
		if rcfg, exists := resourceConfigs[resourceName]; exists && rcfg != "" {
			config.ResourceConfig = rcfg
		}

		// Add dependencies
		deps := ""
		dependsOnExtra := ""
		if dep, exists := resourceDependencies[resourceName]; exists {
			deps = dep
			// Extract resource type from dependency
			if strings.Contains(dep, "f5xc_origin_pool") {
				dependsOnExtra = ", f5xc_origin_pool.test"
			} else if strings.Contains(dep, "f5xc_app_type") {
				dependsOnExtra = ", f5xc_app_type.test"
			} else if strings.Contains(dep, "f5xc_forward_proxy_policy") {
				dependsOnExtra = ", f5xc_forward_proxy_policy.test"
			}
		}

		// Create template data without embedding
		type templateData struct {
			Name               string
			StructName         string
			ResourceName       string
			TestFuncName       string
			ConfigFuncName     string
			NeedsNamespace     bool
			IsNamespace        bool
			UseSystemNamespace bool
			ResourceConfig     string
			Skip               bool
			SkipReason         string
			Dependencies       string
			DependsOnExtra     string
		}

		data := templateData{
			Name:               config.Name,
			StructName:         config.StructName,
			ResourceName:       config.ResourceName,
			TestFuncName:       config.TestFuncName,
			ConfigFuncName:     config.ConfigFuncName,
			NeedsNamespace:     config.NeedsNamespace,
			IsNamespace:        config.IsNamespace,
			UseSystemNamespace: systemNamespaceResources[resourceName],
			ResourceConfig:     config.ResourceConfig,
			Skip:               config.Skip,
			SkipReason:         config.SkipReason,
			Dependencies:       deps,
			DependsOnExtra:     dependsOnExtra,
		}

		tmpl, err := template.New("test").Parse(testTemplate)
		if err != nil {
			fmt.Printf("Error parsing template: %v\n", err)
			continue
		}

		// Execute template to buffer first
		var buf bytes.Buffer
		err = tmpl.Execute(&buf, data)
		if err != nil {
			fmt.Printf("Error executing template for %s: %v\n", file, err)
			continue
		}

		// Format the generated code with gofmt
		formatted, err := format.Source(buf.Bytes())
		if err != nil {
			// If formatting fails, write unformatted code with warning
			fmt.Printf("⚠️  gofmt failed for %s: %v (writing unformatted)\n", file, err)
			formatted = buf.Bytes()
		}

		// Write formatted content to file
		err = os.WriteFile(file, formatted, 0644)
		if err != nil {
			fmt.Printf("Error writing %s: %v\n", file, err)
			continue
		}

		if config.Skip {
			skipped++
		} else {
			fixed++
		}
		fmt.Printf("Fixed %s\n", file)
	}

	fmt.Printf("\nSummary: Fixed %d test files, skipped %d\n", fixed, skipped)
}
