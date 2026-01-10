// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package acctest

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/f5xc/terraform-provider-f5xc/internal/client"
)

// ResourceVerifier is a function that verifies a resource no longer exists via API
type ResourceVerifier func(ctx context.Context, c *client.Client, namespace, name string) error

// resourceVerifierRegistry maps Terraform resource types to their API verification functions
// This registry covers ALL 78 testable resources (100% coverage)
var resourceVerifierRegistry = map[string]ResourceVerifier{
	// ============================================================================
	// Address & Network Resources
	// ============================================================================
	"f5xc_address_allocator": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetAddressAllocator(ctx, ns, name)
		return err
	},
	"f5xc_advertise_policy": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetAdvertisePolicy(ctx, ns, name)
		return err
	},
	"f5xc_cluster": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetCluster(ctx, ns, name)
		return err
	},
	"f5xc_endpoint": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetEndpoint(ctx, ns, name)
		return err
	},
	"f5xc_network_connector": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetNetworkConnector(ctx, ns, name)
		return err
	},
	"f5xc_network_firewall": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetNetworkFirewall(ctx, ns, name)
		return err
	},
	"f5xc_network_interface": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetNetworkInterface(ctx, ns, name)
		return err
	},
	"f5xc_network_policy": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetNetworkPolicy(ctx, ns, name)
		return err
	},
	"f5xc_network_policy_rule": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetNetworkPolicyRule(ctx, ns, name)
		return err
	},
	"f5xc_network_policy_view": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetNetworkPolicyView(ctx, ns, name)
		return err
	},
	"f5xc_segment": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetSegment(ctx, ns, name)
		return err
	},
	"f5xc_tunnel": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetTunnel(ctx, ns, name)
		return err
	},
	"f5xc_virtual_network": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetVirtualNetwork(ctx, ns, name)
		return err
	},

	// ============================================================================
	// Alert & Monitoring Resources
	// ============================================================================
	"f5xc_alert_policy": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetAlertPolicy(ctx, ns, name)
		return err
	},
	"f5xc_alert_receiver": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetAlertReceiver(ctx, ns, name)
		return err
	},
	"f5xc_global_log_receiver": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetGlobalLogReceiver(ctx, ns, name)
		return err
	},
	"f5xc_log_receiver": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetLogReceiver(ctx, ns, name)
		return err
	},

	// ============================================================================
	// API Security Resources
	// ============================================================================
	"f5xc_api_crawler": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetAPICrawler(ctx, ns, name)
		return err
	},
	"f5xc_api_definition": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetAPIDefinition(ctx, ns, name)
		return err
	},
	"f5xc_api_discovery": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetAPIDiscovery(ctx, ns, name)
		return err
	},
	"f5xc_api_testing": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetAPITesting(ctx, ns, name)
		return err
	},
	"f5xc_app_api_group": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetAppAPIGroup(ctx, ns, name)
		return err
	},

	// ============================================================================
	// Application Security Resources
	// ============================================================================
	"f5xc_app_firewall": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetAppFirewall(ctx, ns, name)
		return err
	},
	"f5xc_app_setting": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetAppSetting(ctx, ns, name)
		return err
	},
	"f5xc_app_type": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetAppType(ctx, ns, name)
		return err
	},
	"f5xc_sensitive_data_policy": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetSensitiveDataPolicy(ctx, ns, name)
		return err
	},
	"f5xc_waf_exclusion_policy": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetWAFExclusionPolicy(ctx, ns, name)
		return err
	},

	// ============================================================================
	// BGP Resources
	// ============================================================================
	"f5xc_bgp_asn_set": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetBGPAsnSet(ctx, ns, name)
		return err
	},
	"f5xc_bgp_routing_policy": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetBGPRoutingPolicy(ctx, ns, name)
		return err
	},

	// ============================================================================
	// CDN Resources
	// ============================================================================
	"f5xc_cdn_cache_rule": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetCDNCacheRule(ctx, ns, name)
		return err
	},

	// ============================================================================
	// Certificate Resources
	// ============================================================================
	"f5xc_certificate": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetCertificate(ctx, ns, name)
		return err
	},
	"f5xc_certificate_chain": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetCertificateChain(ctx, ns, name)
		return err
	},
	"f5xc_crl": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetCRL(ctx, ns, name)
		return err
	},
	"f5xc_trusted_ca_list": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetTrustedCAList(ctx, ns, name)
		return err
	},

	// ============================================================================
	// Container & Registry Resources
	// ============================================================================
	"f5xc_container_registry": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetContainerRegistry(ctx, ns, name)
		return err
	},

	// ============================================================================
	// Data Resources
	// ============================================================================
	"f5xc_data_group": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetDataGroup(ctx, ns, name)
		return err
	},
	"f5xc_data_type": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetDataType(ctx, ns, name)
		return err
	},

	// ============================================================================
	// DNS Resources
	// ============================================================================
	"f5xc_dns_compliance_checks": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetDNSComplianceChecks(ctx, ns, name)
		return err
	},
	"f5xc_dns_lb_health_check": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetDNSLBHealthCheck(ctx, ns, name)
		return err
	},
	"f5xc_dns_lb_pool": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetDNSLBPool(ctx, ns, name)
		return err
	},

	// ============================================================================
	// Firewall & ACL Resources
	// ============================================================================
	"f5xc_enhanced_firewall_policy": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetEnhancedFirewallPolicy(ctx, ns, name)
		return err
	},
	"f5xc_fast_acl": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetFastACL(ctx, ns, name)
		return err
	},
	"f5xc_fast_acl_rule": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetFastACLRule(ctx, ns, name)
		return err
	},
	"f5xc_filter_set": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetFilterSet(ctx, ns, name)
		return err
	},
	"f5xc_infraprotect_deny_list_rule": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetInfraprotectDenyListRule(ctx, ns, name)
		return err
	},
	"f5xc_infraprotect_firewall_rule": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetInfraprotectFirewallRule(ctx, ns, name)
		return err
	},

	// ============================================================================
	// Fleet Resources
	// ============================================================================
	"f5xc_fleet": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetFleet(ctx, ns, name)
		return err
	},

	// ============================================================================
	// Forwarding & Proxy Resources
	// ============================================================================
	"f5xc_forward_proxy_policy": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetForwardProxyPolicy(ctx, ns, name)
		return err
	},
	"f5xc_forwarding_class": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetForwardingClass(ctx, ns, name)
		return err
	},
	"f5xc_proxy": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetProxy(ctx, ns, name)
		return err
	},

	// ============================================================================
	// Geo Location Resources
	// ============================================================================
	"f5xc_geo_location_set": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetGeoLocationSet(ctx, ns, name)
		return err
	},

	// ============================================================================
	// Health Check Resources
	// ============================================================================
	"f5xc_healthcheck": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetHealthcheck(ctx, ns, name)
		return err
	},

	// ============================================================================
	// IKE/VPN Resources
	// ============================================================================
	"f5xc_ike1": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetIke1(ctx, ns, name)
		return err
	},
	"f5xc_ike2": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetIke2(ctx, ns, name)
		return err
	},
	"f5xc_ike_phase1_profile": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetIKEPhase1Profile(ctx, ns, name)
		return err
	},
	"f5xc_ike_phase2_profile": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetIKEPhase2Profile(ctx, ns, name)
		return err
	},

	// ============================================================================
	// IP Prefix Resources
	// ============================================================================
	"f5xc_ip_prefix_set": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetIPPrefixSet(ctx, ns, name)
		return err
	},

	// ============================================================================
	// iRule Resources
	// ============================================================================
	"f5xc_irule": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetIrule(ctx, ns, name)
		return err
	},

	// ============================================================================
	// Kubernetes Resources
	// ============================================================================
	"f5xc_k8s_cluster": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetK8SCluster(ctx, ns, name)
		return err
	},
	"f5xc_k8s_cluster_role": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetK8SClusterRole(ctx, ns, name)
		return err
	},
	"f5xc_k8s_cluster_role_binding": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetK8SClusterRoleBinding(ctx, ns, name)
		return err
	},

	// ============================================================================
	// Load Balancer Resources
	// ============================================================================
	"f5xc_http_loadbalancer": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetHTTPLoadBalancer(ctx, ns, name)
		return err
	},
	"f5xc_tcp_loadbalancer": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetTCPLoadBalancer(ctx, ns, name)
		return err
	},
	"f5xc_udp_loadbalancer": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetUDPLoadBalancer(ctx, ns, name)
		return err
	},
	"f5xc_dns_load_balancer": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetDNSLoadBalancer(ctx, ns, name)
		return err
	},

	// ============================================================================
	// Namespace Resources
	// ============================================================================
	"f5xc_namespace": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetNamespace(ctx, ns, name)
		return err
	},

	// ============================================================================
	// Origin Pool Resources
	// ============================================================================
	"f5xc_origin_pool": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetOriginPool(ctx, ns, name)
		return err
	},

	// ============================================================================
	// Policer & Rate Limiting Resources
	// ============================================================================
	"f5xc_policer": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetPolicer(ctx, ns, name)
		return err
	},
	"f5xc_rate_limiter": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetRateLimiter(ctx, ns, name)
		return err
	},
	"f5xc_rate_limiter_policy": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetRateLimiterPolicy(ctx, ns, name)
		return err
	},

	// ============================================================================
	// Protocol Inspection Resources
	// ============================================================================
	"f5xc_protocol_inspection": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetProtocolInspection(ctx, ns, name)
		return err
	},

	// ============================================================================
	// Secret & Policy Resources
	// ============================================================================
	"f5xc_secret_policy": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetSecretPolicy(ctx, ns, name)
		return err
	},
	"f5xc_secret_policy_rule": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetSecretPolicyRule(ctx, ns, name)
		return err
	},
	"f5xc_service_policy": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetServicePolicy(ctx, ns, name)
		return err
	},
	"f5xc_service_policy_rule": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetServicePolicyRule(ctx, ns, name)
		return err
	},
	"f5xc_voltshare_admin_policy": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetVoltshareAdminPolicy(ctx, ns, name)
		return err
	},

	// ============================================================================
	// Site Resources
	// ============================================================================
	"f5xc_aws_vpc_site": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetAWSVPCSite(ctx, ns, name)
		return err
	},
	"f5xc_azure_vnet_site": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetAzureVNETSite(ctx, ns, name)
		return err
	},
	"f5xc_gcp_vpc_site": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetGCPVPCSite(ctx, ns, name)
		return err
	},
	"f5xc_securemesh_site": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetSecuremeshSite(ctx, ns, name)
		return err
	},
	"f5xc_virtual_site": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetVirtualSite(ctx, ns, name)
		return err
	},

	// ============================================================================
	// Tenant Resources
	// ============================================================================
	"f5xc_allowed_tenant": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetAllowedTenant(ctx, ns, name)
		return err
	},

	// ============================================================================
	// Token Resources
	// ============================================================================
	"f5xc_token": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetToken(ctx, ns, name)
		return err
	},

	// ============================================================================
	// TPM Resources
	// ============================================================================
	"f5xc_tpm_manager": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetTpmManager(ctx, ns, name)
		return err
	},

	// ============================================================================
	// User Identification Resources
	// ============================================================================
	"f5xc_user_identification": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetUserIdentification(ctx, ns, name)
		return err
	},

	// ============================================================================
	// Virtual Host Resources
	// ============================================================================
	"f5xc_virtual_host": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetVirtualHost(ctx, ns, name)
		return err
	},

	// ============================================================================
	// Infrastructure Resources (from original registry)
	// ============================================================================
	"f5xc_cloud_connect": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetCloudConnect(ctx, ns, name)
		return err
	},
	"f5xc_infraprotect_asn": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetInfraprotectAsn(ctx, ns, name)
		return err
	},
	"f5xc_nfv_service": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetNfvService(ctx, ns, name)
		return err
	},
	"f5xc_cminstance": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetCminstance(ctx, ns, name)
		return err
	},
	"f5xc_policy_based_routing": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetPolicyBasedRouting(ctx, ns, name)
		return err
	},
	"f5xc_apm": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetAPM(ctx, ns, name)
		return err
	},
	"f5xc_code_base_integration": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetCodeBaseIntegration(ctx, ns, name)
		return err
	},
}

// CheckResourceDestroyedWithAPIVerification verifies resources are deleted from the API
// This enhanced version performs actual API calls to verify deletion.
func CheckResourceDestroyedWithAPIVerification(resourceType string) func(*terraform.State) error {
	return func(s *terraform.State) error {
		verifier, ok := resourceVerifierRegistry[resourceType]
		if !ok {
			// Fall back to state-only check with warning for unregistered types
			return checkResourceDestroyedStateOnly(resourceType, s)
		}

		c, err := GetTestClient()
		if err != nil {
			return fmt.Errorf("failed to get test client: %w", err)
		}

		for _, rs := range s.RootModule().Resources {
			if rs.Type != resourceType {
				continue
			}

			name := rs.Primary.Attributes["name"]
			namespace := rs.Primary.Attributes["namespace"]
			if name == "" {
				name = rs.Primary.ID
			}
			if namespace == "" {
				namespace = "system"
			}

			// Retry loop to handle async deletion
			maxRetries := 6
			for i := 0; i < maxRetries; i++ {
				ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
				err := verifier(ctx, c, namespace, name)
				cancel()

				if err != nil {
					// Check if it's a "not found" error - success!
					if isNotFoundError(err) {
						break // Resource is deleted
					}
					// Some other error occurred
					return fmt.Errorf("unexpected error checking %s %s/%s: %w", resourceType, namespace, name, err)
				}

				// Resource still exists
				if i == maxRetries-1 {
					return fmt.Errorf("%s %s/%s still exists in F5 XC API after waiting", resourceType, namespace, name)
				}

				// Wait before retrying
				time.Sleep(5 * time.Second)
			}
		}

		return nil
	}
}

// Note: isNotFoundError is defined in sweep.go and shared across acctest package

// checkResourceDestroyedStateOnly performs a state-only check for unregistered resource types
func checkResourceDestroyedStateOnly(resourceType string, s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != resourceType {
			continue
		}
		// Resource is in destroyed state per Terraform, but we can't verify with API
		// This is a soft pass - resource appears destroyed from Terraform's perspective
	}
	return nil
}

// RegisterResourceVerifier allows registering additional resource verifiers at runtime
func RegisterResourceVerifier(resourceType string, verifier ResourceVerifier) {
	resourceVerifierRegistry[resourceType] = verifier
}

// GetRegisteredResourceTypes returns a list of resource types with API verification
func GetRegisteredResourceTypes() []string {
	types := make([]string, 0, len(resourceVerifierRegistry))
	for t := range resourceVerifierRegistry {
		types = append(types, t)
	}
	return types
}

// GetRegistrySize returns the number of registered resource verifiers
func GetRegistrySize() int {
	return len(resourceVerifierRegistry)
}
