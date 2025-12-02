// Package namespace provides consistent namespace classification for F5XC resources.
// This package centralizes the logic for determining which namespace type a resource
// should use, following F5XC best practices.
package namespace

// Type represents the type of namespace a resource should use.
type Type int

const (
	// System is for infrastructure objects (sites, networks, fleet, cluster, cloud credentials).
	System Type = iota
	// Shared is for cross-app security policies (app_firewall, certificates, rate_limiters).
	Shared
	// Application is for app-specific workloads (load balancers, origin pools, healthchecks).
	Application
)

// String returns the string representation of the namespace type.
func (t Type) String() string {
	switch t {
	case System:
		return "system"
	case Shared:
		return "shared"
	case Application:
		return "staging"
	default:
		return "staging"
	}
}

// systemResources lists resources that belong in the "system" namespace.
// These are infrastructure-level objects.
var systemResources = map[string]bool{
	// Site resources
	"aws_vpc_site": true, "azure_vnet_site": true, "gcp_vpc_site": true,
	"securemesh_site": true, "securemesh_site_v2": true, "voltstack_site": true,
	"aws_tgw_site": true,
	// Network infrastructure
	"virtual_network": true, "network_interface": true, "network_connector": true,
	"subnet": true, "tunnel": true, "network_firewall": true,
	// Fleet and cluster management
	"fleet": true, "cluster": true, "dc_cluster_group": true, "site_mesh_group": true,
	// Cloud infrastructure
	"cloud_credentials": true, "cloud_connect": true, "cloud_link": true, "cloud_elastic_ip": true,
	// BGP and routing
	"bgp": true, "bgp_asn_set": true, "bgp_routing_policy": true, "route": true,
	// Core system resources
	"namespace": true, "virtual_site": true, "global_log_receiver": true,
	"tenant_configuration": true, "tenant_profile": true, "role": true,
	"k8s_cluster": true, "k8s_cluster_role": true, "k8s_cluster_role_binding": true,
	// Additional system resources
	"token": true, "api_credential": true, "user": true, "service_credential": true,
	"known_label": true, "known_label_key": true,
	"fast_acl": true, "fast_acl_rule": true,
	"network_policy_rule": true, "network_policy_view": true,
	"view_internal_ref": true,
}

// sharedResources lists resources that belong in the "shared" namespace.
// These are cross-team/application reusable resources.
var sharedResources = map[string]bool{
	// Security policies
	"app_firewall": true, "waf_exclusion_policy": true,
	"service_policy": true, "service_policy_rule": true,
	"sensitive_data_policy": true, "secret_policy": true, "secret_policy_rule": true,
	// Certificates and trust
	"certificate": true, "certificate_chain": true, "trusted_ca_list": true, "crl": true,
	// Rate limiting
	"rate_limiter": true, "rate_limiter_policy": true,
	// User identification and mitigation
	"user_identification": true, "malicious_user_mitigation": true,
	// Bot defense
	"bot_defense_app_infrastructure": true,
	// API definitions (shared across apps)
	"api_definition": true, "data_type": true,
	// Network policy sets
	"ip_prefix_set": true, "geo_location_set": true,
	// Protocol inspection
	"protocol_inspection": true, "protocol_policer": true, "policer": true,
	// Forward proxy
	"forward_proxy_policy": true,
	// Alert configuration
	"alert_policy": true, "alert_receiver": true,
	// Additional shared resources
	"data_group": true, "filter_set": true,
	"forwarding_class": true, "log_receiver": true,
}

// ForResource returns the appropriate namespace type and string based on F5XC best practices:
// - system: Infrastructure objects (sites, networks, fleet, cluster, cloud credentials)
// - shared: Cross-app security policies (app_firewall, certificates, rate_limiters)
// - staging: App-specific workloads (load balancers, origin pools, healthchecks)
func ForResource(resourceName string) (Type, string) {
	if systemResources[resourceName] {
		return System, "system"
	}
	if sharedResources[resourceName] {
		return Shared, "shared"
	}
	// Default to application namespace for workload resources
	return Application, "staging"
}

// ForReference returns the appropriate namespace for a resource reference
// based on what type of resource is being referenced.
func ForReference(referencedResourceType string) string {
	_, ns := ForResource(referencedResourceType)
	return ns
}

// IsSystem returns true if the resource belongs in the system namespace.
func IsSystem(resourceName string) bool {
	return systemResources[resourceName]
}

// IsShared returns true if the resource belongs in the shared namespace.
func IsShared(resourceName string) bool {
	return sharedResources[resourceName]
}

// IsApplication returns true if the resource belongs in an application namespace.
func IsApplication(resourceName string) bool {
	return !systemResources[resourceName] && !sharedResources[resourceName]
}

// GetSystemResources returns a copy of the system resources map.
func GetSystemResources() map[string]bool {
	result := make(map[string]bool, len(systemResources))
	for k, v := range systemResources {
		result[k] = v
	}
	return result
}

// GetSharedResources returns a copy of the shared resources map.
func GetSharedResources() map[string]bool {
	result := make(map[string]bool, len(sharedResources))
	for k, v := range sharedResources {
		result[k] = v
	}
	return result
}
