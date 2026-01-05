// Package resource provides resource metadata, categories, and skip lists
// for code generation tools in the F5XC Terraform provider.
package resource

import (
	"strings"
	"sync"
)

// =============================================================================
// V2 Category Cache - For categories from x-f5xc-category extensions
// =============================================================================

// v2CategoryCache stores categories from v2 spec x-f5xc-category extensions.
// This takes priority over overrides and pattern matching.
var v2CategoryCache = make(map[string]string)
var v2CategoryMutex sync.RWMutex

// SetV2Category sets the category for a resource from v2 spec metadata.
// This should be called during spec parsing when x-f5xc-category is found.
func SetV2Category(resourceName string, category string) {
	if category == "" {
		return
	}
	v2CategoryMutex.Lock()
	defer v2CategoryMutex.Unlock()
	v2CategoryCache[resourceName] = category
}

// GetV2Category retrieves the v2 category for a resource if set.
func GetV2Category(resourceName string) (string, bool) {
	v2CategoryMutex.RLock()
	defer v2CategoryMutex.RUnlock()
	cat, ok := v2CategoryCache[resourceName]
	return cat, ok
}

// ClearV2Categories clears all v2 category cache entries (for testing).
func ClearV2Categories() {
	v2CategoryMutex.Lock()
	defer v2CategoryMutex.Unlock()
	v2CategoryCache = make(map[string]string)
}

// V2CategoryCount returns the number of v2 categories in cache.
func V2CategoryCount() int {
	v2CategoryMutex.RLock()
	defer v2CategoryMutex.RUnlock()
	return len(v2CategoryCache)
}

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
// 1. Check v2 spec x-f5xc-category first (from API enrichment, most specific)
// 2. Check explicit SubcategoryOverrides (for exceptions and manual overrides)
// 3. Apply pattern matching (for automatic categorization)
// 4. Default to "Uncategorized"
func GetCategory(resourceName string) string {
	// Tier 1: Check v2 spec x-f5xc-category (preferred, from enriched API specs)
	if category, ok := GetV2Category(resourceName); ok {
		return category
	}

	// Tier 2: Check explicit overrides
	if category, ok := SubcategoryOverrides[resourceName]; ok {
		return category
	}

	// Tier 3: Apply pattern matching
	for _, pattern := range CategoryPatterns {
		if strings.Contains(resourceName, pattern.Pattern) {
			return pattern.Category
		}
	}

	// Default to Uncategorized
	return "Uncategorized"
}

// GetCategorySource returns both the category and its source tier for debugging.
// Returns (category, source) where source is "v2", "override", "pattern", or "default".
func GetCategorySource(resourceName string) (string, string) {
	// Tier 1: Check v2 spec x-f5xc-category
	if category, ok := GetV2Category(resourceName); ok {
		return category, "v2"
	}

	// Tier 2: Check explicit overrides
	if category, ok := SubcategoryOverrides[resourceName]; ok {
		return category, "override"
	}

	// Tier 3: Apply pattern matching
	for _, pattern := range CategoryPatterns {
		if strings.Contains(resourceName, pattern.Pattern) {
			return pattern.Category, "pattern"
		}
	}

	// Default to Uncategorized
	return "Uncategorized", "default"
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
