//go:build ignore
// +build ignore

// This tool queries the F5 Distributed Cloud catalog API to generate
// subscription tier metadata for resources. This metadata is used by
// documentation generators and the MCP server to indicate which resources
// require Advanced subscription tiers.
//
// Usage:
//   F5XC_API_URL="https://..." F5XC_P12_FILE="..." F5XC_P12_PASSWORD="..." go run tools/generate-subscription-metadata.go
//
// Environment Variables:
//   F5XC_API_URL      - F5 XC API URL (e.g., https://console.ves.volterra.io)
//   F5XC_P12_FILE     - Path to P12 certificate file
//   F5XC_P12_PASSWORD - Password for P12 file
//   F5XC_API_TOKEN    - API token (alternative to P12 auth)
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/f5xc/terraform-provider-f5xc/internal/client"
)

// CatalogListRequest is the request body for the catalog API
type CatalogListRequest struct {
	// Empty request to get full catalog
}

// CatalogListResponse is the response from the catalog API
type CatalogListResponse struct {
	AddonServices map[string]AddonServiceInfo `json:"addon_services"`
	Services      map[string]ServiceInfo      `json:"services"`
	Workspaces    map[string]WorkspaceInfo    `json:"workspaces"`
	UseCases      map[string]UseCaseInfo      `json:"use_cases"`
}

// AddonServiceInfo contains addon service details from catalog
type AddonServiceInfo struct {
	Tier                    string `json:"tier"`
	DisplayName             string `json:"display_name"`
	AddonServiceGroupName   string `json:"addon_service_group_name"`
}

// ServiceInfo contains service details from catalog
type ServiceInfo struct {
	Name                   string   `json:"name"`
	DisplayName            string   `json:"display_name"`
	Description            string   `json:"description"`
	Tier                   string   `json:"tier"`
	AccessStatus           string   `json:"access_status"`
	AddonServiceStatus     string   `json:"addon_service_status"`
	AddonServiceGroupStatus string   `json:"addon_service_group_status"`
	Tags                   []string `json:"tags"`
}

// WorkspaceInfo contains workspace details from catalog
type WorkspaceInfo struct {
	Name             string   `json:"name"`
	DisplayName      string   `json:"display_name"`
	Services         []string `json:"services"`
	RequiredServices []string `json:"required_services"`
	OptionalServices []string `json:"optional_services"`
}

// UseCaseInfo contains use case details from catalog
type UseCaseInfo struct {
	Name        string   `json:"name"`
	DisplayName string   `json:"display_name"`
	Workspaces  []string `json:"workspaces"`
	Order       int      `json:"order"`
}

// SubscriptionMetadata is the output format for subscription tier metadata
type SubscriptionMetadata struct {
	GeneratedAt string                       `json:"generated_at"`
	Source      string                       `json:"source"`
	Services    map[string]ServiceMetadata   `json:"services"`
	Resources   map[string]ResourceMetadata  `json:"resources"`
}

// ServiceMetadata contains metadata about an addon service
type ServiceMetadata struct {
	Tier        string `json:"tier"`
	DisplayName string `json:"display_name"`
	GroupName   string `json:"group_name,omitempty"`
}

// ResourceMetadata contains subscription metadata for a Terraform resource
type ResourceMetadata struct {
	Service         string   `json:"service"`
	MinimumTier     string   `json:"minimum_tier"`
	AdvancedFeatures []string `json:"advanced_features,omitempty"`
}

// resourceServiceMapping maps OpenAPI schema path patterns to service groups
// This is based on the F5 XC API structure where resources are organized by service
var resourceServiceMapping = map[string]string{
	// WAAP (Web App & API Protection)
	"views.http_loadbalancer":       "waap",
	"views.tcp_loadbalancer":        "waap",
	"views.udp_loadbalancer":        "waap",
	"views.origin_pool":             "waap",
	"app_firewall":                  "waap",
	"waf":                           "waap",
	"waf_exclusion_policy":          "waap",
	"rate_limiter":                  "waap",
	"views.rate_limiter_policy":     "waap",
	"service_policy":                "waap",
	"service_policy_rule":           "waap",
	"service_policy_set":            "waap",
	"malicious_user_mitigation":     "waap",
	"user_identification":           "waap",
	"views.api_definition":          "waap",
	"api_sec.api_discovery":         "waap",
	"api_sec.api_crawler":           "waap",
	"api_sec.api_testing":           "waap",

	// Bot Defense (part of WAAP Advanced)
	"shape.bot_defense":                        "waap_advanced",
	"shape.bot_defense.protected_application":  "waap_advanced",
	"shape.bot_defense.bot_endpoint_policy":    "waap_advanced",
	"shape.bot_defense.bot_infrastructure":     "waap_advanced",
	"shape.bot_defense.bot_allowlist_policy":   "waap_advanced",
	"shape.bot_defense.bot_network_policy":     "waap_advanced",
	"shape.bot_defense.mobile_sdk":             "waap_advanced",
	"shape.bot_defense.mobile_base_config":     "waap_advanced",
	"shape_bot_defense_instance":               "waap_advanced",

	// CDN (Content Delivery Network)
	"views.cdn_loadbalancer":        "cdn",
	"cdn_cache_rule":                "cdn",

	// SecureMesh / Site Management
	"views.securemesh_site":         "securemesh",
	"views.securemesh_site_v2":      "securemesh",
	"views.voltstack_site":          "securemesh",
	"views.aws_vpc_site":            "site_management",
	"views.azure_vnet_site":         "site_management",
	"views.gcp_vpc_site":            "site_management",
	"views.aws_tgw_site":            "site_management",

	// Network Connect
	"network_connector":             "network_connect",
	"network_firewall":              "network_connect",
	"network_policy":                "network_connect",
	"network_policy_rule":           "network_connect",
	"network_policy_set":            "network_connect",
	"views.network_policy_view":     "network_connect",
	"enhanced_firewall_policy":      "network_connect",
	"fast_acl":                      "network_connect",
	"fast_acl_rule":                 "network_connect",

	// DNS Management
	"dns_zone":                      "dns",
	"dns_domain":                    "dns",
	"dns_load_balancer":             "dns",
	"dns_lb_health_check":           "dns",
	"dns_lb_pool":                   "dns",

	// App Stack
	"k8s_cluster":                   "appstack",
	"k8s_cluster_role":              "appstack",
	"k8s_cluster_role_binding":      "appstack",
	"k8s_pod_security_admission":    "appstack",
	"k8s_pod_security_policy":       "appstack",
	"virtual_k8s":                   "appstack",
	"views.workload":                "appstack",

	// Client Side Defense
	"shape.client_side_defense":                "client_side_defense",
	"shape.client_side_defense.allowed_domain": "client_side_defense",
	"shape.client_side_defense.protected_domain": "client_side_defense",
	"shape.client_side_defense.mitigated_domain": "client_side_defense",

	// Observability / Synthetic Monitoring
	"observability.synthetic_monitor":           "synthetic_monitoring",
	"observability.synthetic_monitor.v1_dns_monitor": "synthetic_monitoring",
	"observability.synthetic_monitor.v1_http_monitor": "synthetic_monitoring",

	// Core/Common (no specific tier required)
	"namespace":                     "core",
	"healthcheck":                   "core",
	"certificate":                   "core",
	"certificate_chain":             "core",
	"secret_policy":                 "core",
	"secret_policy_rule":            "core",
	"trusted_ca_list":               "core",
	"virtual_site":                  "core",
	"virtual_network":               "core",
	"virtual_host":                  "core",
	"cloud_credentials":             "core",
	"discovery":                     "core",
	"endpoint":                      "core",
	"known_label":                   "core",
	"known_label_key":               "core",
	"ip_prefix_set":                 "core",
	"bgp_asn_set":                   "core",
	"geo_location_set":              "core",
	"log_receiver":                  "core",
	"global_log_receiver":           "core",
	"alert_policy":                  "core",
	"alert_receiver":                "core",
}

// serviceToTierMapping maps service groups to their tier requirements
// This defines which features/resources require Advanced tier
var serviceToTierMapping = map[string]string{
	"waap":              "STANDARD",  // Basic WAAP is Standard
	"waap_advanced":     "ADVANCED",  // Bot defense, API security is Advanced
	"cdn":               "STANDARD",  // CDN has both Standard and Advanced
	"securemesh":        "STANDARD",  // SecureMesh has both Standard and Advanced
	"site_management":   "STANDARD",
	"network_connect":   "STANDARD",
	"dns":               "STANDARD",
	"appstack":          "STANDARD",
	"client_side_defense": "ADVANCED",  // CSD is part of advanced
	"synthetic_monitoring": "STANDARD",
	"core":              "NO_TIER",   // Core resources don't require specific tier
}

// advancedFeaturesList lists features that are only in Advanced tier
var advancedFeaturesList = map[string][]string{
	"http_loadbalancer": {"api_discovery", "bot_defense", "malicious_user_detection", "sensitive_data_policy"},
	"app_firewall":      {"advanced_detection_settings", "ai_bot_defense"},
	"cdn_loadbalancer":  {"advanced_caching", "edge_compute"},
}

func main() {
	// Get credentials from environment
	apiURL := os.Getenv("F5XC_API_URL")
	p12File := os.Getenv("F5XC_P12_FILE")
	p12Password := os.Getenv("F5XC_P12_PASSWORD")
	apiToken := os.Getenv("F5XC_API_TOKEN")

	if apiURL == "" {
		fmt.Fprintln(os.Stderr, "Error: F5XC_API_URL environment variable is required")
		os.Exit(1)
	}

	var apiClient *client.Client
	var err error

	// Try P12 authentication first, then fall back to token
	if p12File != "" && p12Password != "" {
		apiClient, err = client.NewClientWithP12(apiURL, p12File, p12Password)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating client with P12: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Using P12 certificate authentication")
	} else if apiToken != "" {
		apiClient = client.NewClient(apiURL, apiToken)
		fmt.Println("Using API token authentication")
	} else {
		fmt.Fprintln(os.Stderr, "Error: Either F5XC_P12_FILE+F5XC_P12_PASSWORD or F5XC_API_TOKEN is required")
		os.Exit(1)
	}

	// Query the catalog API
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var catalogResp CatalogListResponse
	err = apiClient.Put(ctx, "/api/web/namespaces/system/catalogs", CatalogListRequest{}, &catalogResp)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error querying catalog API: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Retrieved catalog with %d addon services, %d services, %d workspaces\n",
		len(catalogResp.AddonServices), len(catalogResp.Services), len(catalogResp.Workspaces))

	// Build subscription metadata
	metadata := buildSubscriptionMetadata(catalogResp)

	// Write output file
	outputPath := "tools/subscription-tiers.json"

	// Check if existing file has same content (ignoring timestamp)
	// to avoid unnecessary changes during pre-commit hooks
	if existingData, err := os.ReadFile(outputPath); err == nil {
		var existingMeta SubscriptionMetadata
		if json.Unmarshal(existingData, &existingMeta) == nil {
			// Compare content without timestamps
			existingMeta.GeneratedAt = ""
			metadataCompare := metadata
			metadataCompare.GeneratedAt = ""

			existingJSON, _ := json.Marshal(existingMeta)
			newJSON, _ := json.Marshal(metadataCompare)

			if string(existingJSON) == string(newJSON) {
				fmt.Printf("No changes detected in %s (keeping existing file)\n", outputPath)
				fmt.Printf("Existing file has %d services and %d resources\n",
					len(existingMeta.Services), len(existingMeta.Resources))
				return
			}
		}
	}

	outputData, err := json.MarshalIndent(metadata, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling metadata: %v\n", err)
		os.Exit(1)
	}

	// Add trailing newline for POSIX compliance
	outputData = append(outputData, '\n')

	if err := os.WriteFile(outputPath, outputData, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing output file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Generated %s with %d services and %d resources\n",
		outputPath, len(metadata.Services), len(metadata.Resources))
}

func buildSubscriptionMetadata(catalog CatalogListResponse) SubscriptionMetadata {
	metadata := SubscriptionMetadata{
		GeneratedAt: time.Now().UTC().Format(time.RFC3339),
		Source:      "F5 XC Catalog API",
		Services:    make(map[string]ServiceMetadata),
		Resources:   make(map[string]ResourceMetadata),
	}

	// Extract service tier information from catalog
	for name, svc := range catalog.AddonServices {
		metadata.Services[name] = ServiceMetadata{
			Tier:        svc.Tier,
			DisplayName: svc.DisplayName,
			GroupName:   svc.AddonServiceGroupName,
		}
	}

	// Also extract from services if addon_services is empty
	for name, svc := range catalog.Services {
		if _, exists := metadata.Services[name]; !exists {
			metadata.Services[name] = ServiceMetadata{
				Tier:        svc.Tier,
				DisplayName: svc.DisplayName,
			}
		}
	}

	// Map resources to services based on OpenAPI spec files
	specDir := "docs/specifications/api"
	files, err := filepath.Glob(filepath.Join(specDir, "*.json"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Warning: Could not scan OpenAPI specs: %v\n", err)
		return metadata
	}

	// Pattern: docs-cloud-f5-com.XXXX.public.ves.io.schema.{path}.ves-swagger.json
	specRegex := regexp.MustCompile(`docs-cloud-f5-com\.\d+\.public\.ves\.io\.schema\.(.+)\.ves-swagger\.json`)

	for _, file := range files {
		base := filepath.Base(file)
		matches := specRegex.FindStringSubmatch(base)
		if matches == nil || len(matches) < 2 {
			continue
		}

		schemaPath := matches[1]
		resourceName := getResourceName(schemaPath)

		// Skip non-resource schemas
		if shouldSkipSchema(schemaPath) {
			continue
		}

		// Find matching service
		service := findServiceForSchema(schemaPath)
		tier := serviceToTierMapping[service]
		if tier == "" {
			tier = "STANDARD" // Default to standard if unknown
		}

		// Get advanced features if any
		advFeatures := advancedFeaturesList[resourceName]

		metadata.Resources[resourceName] = ResourceMetadata{
			Service:         service,
			MinimumTier:     tier,
			AdvancedFeatures: advFeatures,
		}
	}

	return metadata
}

// getResourceName extracts the resource name from a schema path
func getResourceName(schemaPath string) string {
	parts := strings.Split(schemaPath, ".")
	return parts[len(parts)-1]
}

// findServiceForSchema finds the service group for a schema path
func findServiceForSchema(schemaPath string) string {
	// Try exact match first
	if service, ok := resourceServiceMapping[schemaPath]; ok {
		return service
	}

	// Try prefix matching for nested schemas
	for pattern, service := range resourceServiceMapping {
		if strings.HasPrefix(schemaPath, pattern) {
			return service
		}
	}

	// Default to core for unknown schemas
	return "core"
}

// shouldSkipSchema returns true for schemas that shouldn't be mapped to resources
func shouldSkipSchema(schemaPath string) bool {
	skipPatterns := []string{
		"operate.",      // Operational APIs, not resources
		"graph.",        // Graph/visualization APIs
		"usage.",        // Usage tracking
		"billing.",      // Billing
		"signup.",       // Signup flow
		"pbac.",         // Permission-based access control internals
		"tenant_management.", // Tenant management (MSP only)
		"marketplace.",  // Marketplace
		"user.setting",  // User settings
		"ui.",           // UI components
	}

	for _, pattern := range skipPatterns {
		if strings.Contains(schemaPath, pattern) {
			return true
		}
	}

	return false
}
