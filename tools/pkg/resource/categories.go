// Package resource provides resource metadata, categories, and skip lists
// for code generation tools in the F5XC Terraform provider.
package resource

import "strings"

// SubcategoryOverrides provides explicit category assignments for resources
// that don't match any pattern or need a specific override.
var SubcategoryOverrides = map[string]string{
	// Explicit overrides for resources that don't match patterns well
	"apm":                 "Monitoring",
	"crl":                 "Certificates",
	"bgp":                 "Networking",
	"proxy":               "Networking",
	"tunnel":              "Networking",
	"segment":             "Networking",
	"subnet":              "Networking",
	"fleet":               "Sites",
	"cluster":             "Load Balancing",
	"endpoint":            "Load Balancing",
	"route":               "Load Balancing",
	"healthcheck":         "Load Balancing",
	"origin_pool":         "Load Balancing",
	"virtual_host":        "Load Balancing",
	"discovery":           "Applications",
	"filter_set":          "Applications",
	"policer":             "Service Mesh",
	"quota":               "Organization",
	"contact":             "Organization",
	"role":                "Organization",
	"token":               "Authentication",
	"registration":        "Sites",
	"namespace":           "Organization",
	"data_type":           "Security",
	"data_group":          "BIG-IP Integration",
	"irule":               "BIG-IP Integration",
	"nfv_service":         "Networking",
	"workload":            "Kubernetes",
	"workload_flavor":     "Kubernetes",
	"cminstance":          "Subscriptions",
	"user_identification": "Security",
	"virtual_network":     "Networking",
}

// CategoryPattern defines a pattern-to-category mapping.
type CategoryPattern struct {
	Pattern  string
	Category string
}

// CategoryPatterns defines patterns to auto-categorize resources.
// Order matters: more specific patterns should come first.
var CategoryPatterns = []CategoryPattern{
	// Sites - cloud provider sites and site-related
	{"_vpc_site", "Sites"},
	{"_vnet_site", "Sites"},
	{"_tgw_site", "Sites"},
	{"securemesh_site", "Sites"},
	{"voltstack_site", "Sites"},
	{"virtual_site", "Sites"},
	{"site_mesh", "Sites"},

	// Load Balancing
	{"loadbalancer", "Load Balancing"},
	{"cdn_", "Load Balancing"},
	{"advertise_policy", "Load Balancing"},

	// Security - firewall, policies, WAF
	{"firewall", "Security"},
	{"_policy", "Security"},
	{"_acl", "Security"},
	{"rate_limiter", "Security"},
	{"malicious_user", "Security"},
	{"bot_defense", "Security"},
	{"waf_", "Security"},
	{"protocol_", "Security"},
	{"sensitive_data", "Security"},

	// Networking
	{"network_", "Networking"},
	{"bgp_", "Networking"},
	{"ip_prefix", "Networking"},
	{"cloud_link", "Networking"},
	{"cloud_connect", "Networking"},
	{"_connector", "Networking"},
	{"dc_cluster", "Networking"},
	{"srv6_", "Networking"},
	{"forwarding_", "Networking"},
	{"routing", "Networking"},
	{"nat_", "Networking"},

	// DNS
	{"dns_", "DNS"},

	// Authentication & Secrets
	{"credential", "Authentication"},
	{"authentication", "Authentication"},
	{"oidc_", "Authentication"},
	{"secret_", "Authentication"},

	// Certificates
	{"certificate", "Certificates"},
	{"trusted_ca", "Certificates"},

	// API Security
	{"api_", "API Security"},
	{"app_api", "API Security"},

	// Monitoring & Logging
	{"log_receiver", "Monitoring"},
	{"alert_", "Monitoring"},
	{"report_", "Monitoring"},

	// Organization & Tenants
	{"tenant", "Organization"},
	{"_support", "Organization"},
	{"voltshare", "Organization"},
	{"allowed_", "Organization"},

	// Kubernetes
	{"k8s_", "Kubernetes"},
	{"virtual_k8s", "Kubernetes"},
	{"container_registry", "Kubernetes"},

	// VPN & IPSec
	{"ike", "VPN"},

	// Infrastructure Protection (DDoS)
	{"infraprotect_", "Infrastructure Protection"},

	// Applications
	{"app_setting", "Applications"},
	{"app_type", "Applications"},

	// BIG-IP Integration
	{"bigip_", "BIG-IP Integration"},

	// Cloud Resources
	{"cloud_elastic", "Cloud Resources"},
	{"address_allocator", "Cloud Resources"},
	{"geo_location", "Cloud Resources"},

	// Integrations
	{"ticket_tracking", "Integrations"},
	{"code_base", "Integrations"},
	{"tpm_", "Integrations"},

	// Subscriptions
	{"subscription", "Subscriptions"},

	// Service Mesh (catch remaining mesh-related)
	{"usb_policy", "Service Mesh"},
}

// GetCategory returns the category for a resource based on its name.
// Uses a three-tier approach:
// 1. Check explicit overrides first (for exceptions)
// 2. Apply pattern matching (for automatic categorization)
// 3. Default to "Uncategorized"
func GetCategory(resourceName string) string {
	// Check explicit overrides first
	if category, ok := SubcategoryOverrides[resourceName]; ok {
		return category
	}

	// Apply pattern matching
	for _, pattern := range CategoryPatterns {
		if strings.Contains(resourceName, pattern.Pattern) {
			return pattern.Category
		}
	}

	// Default to Uncategorized
	return "Uncategorized"
}

// AllCategories returns a list of all unique categories.
func AllCategories() []string {
	seen := make(map[string]bool)
	var categories []string

	// Add override categories
	for _, cat := range SubcategoryOverrides {
		if !seen[cat] {
			seen[cat] = true
			categories = append(categories, cat)
		}
	}

	// Add pattern categories
	for _, pattern := range CategoryPatterns {
		if !seen[pattern.Category] {
			seen[pattern.Category] = true
			categories = append(categories, pattern.Category)
		}
	}

	return categories
}
