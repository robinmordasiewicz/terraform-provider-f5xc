//go:build ignore
// +build ignore

// This tool discovers API-applied default values by comparing request vs response
// when creating resources against the F5 XC API.
//
// Usage:
//
//	go run tools/discover-defaults.go -resource namespace
//	go run tools/discover-defaults.go -all
//	go run tools/discover-defaults.go -pattern "loadbalancer"
//	go run tools/discover-defaults.go -validate
//
// Environment Variables Required:
//
//	F5XC_API_URL - API URL (e.g., https://tenant.console.ves.volterra.io/api)
//	F5XC_API_P12_FILE + F5XC_P12_PASSWORD - P12 certificate authentication
//	or F5XC_API_TOKEN - Token authentication
package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"time"

	"github.com/f5xc/terraform-provider-f5xc/internal/client"
	"github.com/f5xc/terraform-provider-f5xc/tools/pkg/resource"
)

// ============================================================================
// Data Structures
// ============================================================================

// DefaultsDatabase stores all discovered defaults for the provider
type DefaultsDatabase struct {
	Version      string                      `json:"version"`
	GeneratedAt  string                      `json:"generated_at"`
	APIEndpoint  string                      `json:"api_endpoint"`
	TotalResources int                       `json:"total_resources"`
	Discovered   int                         `json:"discovered"`
	Skipped      int                         `json:"skipped"`
	Failed       int                         `json:"failed"`
	Resources    map[string]*DiscoveryResult `json:"resources"`
}

// DiscoveryResult holds the discovery result for a single resource type
type DiscoveryResult struct {
	ResourceName string                 `json:"resource_name"`
	Category     string                 `json:"category"`
	Status       string                 `json:"status"` // "discovered", "skipped", "failed"
	SkipReason   string                 `json:"skip_reason,omitempty"`
	Error        string                 `json:"error,omitempty"`
	DiscoveredAt string                 `json:"discovered_at,omitempty"`
	Defaults     map[string]FieldDefault `json:"defaults,omitempty"`
	RequestSent  map[string]interface{} `json:"request_sent,omitempty"`
	ResponseGot  map[string]interface{} `json:"response_got,omitempty"`
}

// FieldDefault represents a single field's default value
type FieldDefault struct {
	Path         string      `json:"path"`
	DefaultValue interface{} `json:"default_value"`
	Type         string      `json:"type"` // "string", "int", "bool", "object", "array"
	IsMarkerBlock bool       `json:"is_marker_block,omitempty"`
	Description  string      `json:"description,omitempty"`
}

// ResourceCategory defines complexity categories for discovery
type ResourceCategory string

const (
	CategorySimple      ResourceCategory = "simple"
	CategoryLoadBalancer ResourceCategory = "load_balancer"
	CategoryOriginPool  ResourceCategory = "origin_pool"
	CategorySecurity    ResourceCategory = "security"
	CategoryDNS         ResourceCategory = "dns"
	CategoryCloudSite   ResourceCategory = "cloud_site"
	CategoryTenant      ResourceCategory = "tenant"
)

// ============================================================================
// Skip List - Resources that cannot be tested
// ============================================================================

// Skip list is now centralized in tools/pkg/resource/resource.go
// Use resource.IsSkippedForAPITest() and resource.GetSkipReason() to check

// ============================================================================
// Resource Configuration Templates
// ============================================================================

// MinimalConfig defines the minimal configuration needed to create a resource
type MinimalConfig struct {
	Category    ResourceCategory
	Namespace   bool                   // true if resource requires namespace
	RequiredSpec map[string]interface{} // minimal spec fields required
}

// ResourceConfigs maps resource names to their minimal configurations
var ResourceConfigs = map[string]MinimalConfig{
	// Simple resources - just name (and maybe namespace)
	"namespace": {
		Category:    CategorySimple,
		Namespace:   false,
		RequiredSpec: map[string]interface{}{},
	},
	"healthcheck": {
		Category:  CategorySimple,
		Namespace: true,
		RequiredSpec: map[string]interface{}{
			"http_health_check": map[string]interface{}{
				"path":                    "/healthz",
				"use_origin_server_name": map[string]interface{}{},
			},
			"timeout":              1,
			"interval":             10,
			"unhealthy_threshold":  2,
			"healthy_threshold":    3,
		},
	},
	"alert_policy": {
		Category:  CategorySimple,
		Namespace: true,
		RequiredSpec: map[string]interface{}{
			"policy": []interface{}{
				map[string]interface{}{
					"receivers": []interface{}{},
				},
			},
		},
	},
	"alert_receiver": {
		Category:  CategorySimple,
		Namespace: true,
		RequiredSpec: map[string]interface{}{
			"email": map[string]interface{}{
				"email": "test@example.com",
			},
		},
	},
	"contact": {
		Category:  CategorySimple,
		Namespace: false, // system namespace
		RequiredSpec: map[string]interface{}{
			"contact_info": map[string]interface{}{
				"email": "test@example.com",
			},
		},
	},
	"certificate": {
		Category:  CategorySimple,
		Namespace: true,
		RequiredSpec: map[string]interface{}{
			// Will use self-signed for testing
			"certificate_url": "string:///LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUJrakNDQVRlZ0F3SUJBZ0lSQU1aSEJLY0MvMEw0RnJXQy9oVlBWS0F3Q2dZSUtvWkl6ajBFQXdJd0VURVAKTUE4R0ExVUVDaE1JWkdselkyOTJaWEl3SGhjTk1qTXdNVEF4TURBd01EQXdXaGNOTWpRd01UQXhNREF3TURBdwpXakFSTVE4d0RRWURWUVFLRXdaa2FYTmpiM1psY2pCWk1CTUdCeXFHU000OUFnRUdDQ3FHU000OUF3RUhBMElBCkJMM2lqMG5oTEVhd0hMb29JRlBKT1JMVmRVRFFYYi9nTG0vWjJGekF6Y3VGRUtPWGdUV1RHK0tCMVdVSUZRNmYKY1VhTStGREJpZ2MxNStPUStLQ1A1dmFqWlRCak1BNEdBMVVkRHdFQi93UUVBd0lDaERBZEJnTlZIU1VFRmpBVQpCZ2dyQmdFRkJRY0RBZ1lJS3dZQkJRVUhBd0V3RHdZRFZSMFRBUUgvQkFVd0F3RUIvekFkQmdOVkhRNEVGZ1FVCnhGRHRDeE1BSzF4Q0NRSHFGcEFjT3ovMGptb3dDZ1lJS29aSXpqMEVBd0lEU1FBd1JnSWhBTnlLd3VkYWNxc3oKZjgrbk5WTGU5N1NsZXBHNHpOcFRjSHJ4WTIxam5SY1ZBaUVBcjk3TXQwK1VUdkVOdUJKSGVZcC9YOVFXOUxxegpJSHJtM2xmVGhHdENodHc9Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K",
			"private_key": map[string]interface{}{
				"blindfold_secret_info_internal": map[string]interface{}{
					"decryption_provider": "",
					"location":            "string:///LS0tLS1CRUdJTiBFQyBQUklWQVRFIEtFWS0tLS0tCk1IUUNBUUVFSUlFTUNZR2Yxd2gvZWUvSEt3Z0tnN3JSY0NSOE5IRFA0S0VZUDFDL3FCNjNvQWNHQlN1QkJBQU4Kb1VRRFFnQUV2ZUtQU2VFc1JyQWN1aWdnVThrNUV0VjFRTkJkditBdWI5bllYTUROeTRVUW81ZUJOWk1iNG9IVgpaUWdWRHA5eFJvejRVTUdLQnpYbjQ1RDRvSS9tOWc9PQotLS0tLUVORCBFQyBQUklWQVRFIEtFWS0tLS0tCg==",
					"store_provider":      "",
				},
			},
		},
	},

	// Load Balancer resources
	"http_loadbalancer": {
		Category:  CategoryLoadBalancer,
		Namespace: true,
		RequiredSpec: map[string]interface{}{
			"domains": []interface{}{"test.example.com"},
			"http": map[string]interface{}{
				"dns_volterra_managed": true,
				"port":                 80,
			},
			"advertise_on_public_default_vip": map[string]interface{}{},
			"disable_rate_limit":              map[string]interface{}{},
			"no_service_policies":             map[string]interface{}{},
			"round_robin":                     map[string]interface{}{},
			"disable_waf":                     map[string]interface{}{},
			"default_route_pools":             []interface{}{},
		},
	},
	"tcp_loadbalancer": {
		Category:  CategoryLoadBalancer,
		Namespace: true,
		RequiredSpec: map[string]interface{}{
			"domains":       []interface{}{"test.example.com"},
			"listen_port":   8080,
			"with_sni":      map[string]interface{}{},
			"tcp":           map[string]interface{}{},
			"advertise_on_public_default_vip": map[string]interface{}{},
			"origin_pools_weights":            []interface{}{},
		},
	},

	// Security resources
	"app_firewall": {
		Category:  CategorySecurity,
		Namespace: true,
		RequiredSpec: map[string]interface{}{
			"allow_all_response_codes": true,
			"default_anonymization":    true,
			"use_default_blocking_page": true,
			"default_bot_setting":      true,
			"default_detection_settings": true,
			"use_loadbalancer_setting": true,
			"blocking":                 map[string]interface{}{},
		},
	},
	"service_policy": {
		Category:  CategorySecurity,
		Namespace: true,
		RequiredSpec: map[string]interface{}{
			"algo":    "FIRST_MATCH",
			"any_server": map[string]interface{}{},
			"rule_list": map[string]interface{}{
				"rules": []interface{}{},
			},
		},
	},
	"rate_limiter": {
		Category:  CategorySecurity,
		Namespace: true,
		RequiredSpec: map[string]interface{}{
			"total_number": 100,
			"unit":         "SECOND",
			"burst_multiplier": 1,
		},
	},

	// Origin Pool
	"origin_pool": {
		Category:  CategoryOriginPool,
		Namespace: true,
		RequiredSpec: map[string]interface{}{
			"origin_servers": []interface{}{
				map[string]interface{}{
					"public_ip": map[string]interface{}{
						"ip": "1.2.3.4",
					},
				},
			},
			"port":                    80,
			"endpoint_selection":      "LOCAL_PREFERRED",
			"loadbalancer_algorithm":  "LB_OVERRIDE",
			"no_tls":                  map[string]interface{}{},
		},
	},

	// Virtual Site
	"virtual_site": {
		Category:  CategorySimple,
		Namespace: true,
		RequiredSpec: map[string]interface{}{
			"site_type": "REGIONAL_EDGE",
			"site_selector": map[string]interface{}{
				"expressions": []interface{}{"site.name in (ves-io-mia-1)"},
			},
		},
	},

	// Policer
	"policer": {
		Category:  CategorySimple,
		Namespace: true,
		RequiredSpec: map[string]interface{}{
			"committed_information_rate": 10000,
			"peak_information_rate":      20000,
			"burst_size":                 10000,
			"committed_burst_size":       10000,
			"single_rate_two_color":      map[string]interface{}{},
		},
	},

	// User Identification
	"user_identification": {
		Category:  CategorySecurity,
		Namespace: true,
		RequiredSpec: map[string]interface{}{
			"rules": []interface{}{
				map[string]interface{}{
					"client_identifier": map[string]interface{}{
						"ip_and_tls_fingerprint": map[string]interface{}{},
					},
				},
			},
		},
	},

	// IP Prefix Set
	"ip_prefix_set": {
		Category:  CategorySimple,
		Namespace: true,
		RequiredSpec: map[string]interface{}{
			"prefix": []interface{}{"10.0.0.0/8"},
		},
	},

	// BGP ASN Set
	"bgp_asn_set": {
		Category:  CategorySimple,
		Namespace: true,
		RequiredSpec: map[string]interface{}{
			"as_numbers": []interface{}{65000},
		},
	},

	// Data Group
	"data_group": {
		Category:  CategorySimple,
		Namespace: true,
		RequiredSpec: map[string]interface{}{
			"type": "IP_PREFIX",
			"data": []interface{}{
				map[string]interface{}{
					"ip_prefix": map[string]interface{}{
						"prefix": []interface{}{"10.0.0.0/8"},
					},
				},
			},
		},
	},

	// Data Type
	"data_type": {
		Category:  CategorySimple,
		Namespace: true,
		RequiredSpec: map[string]interface{}{
			"pattern": map[string]interface{}{
				"regex": ".*test.*",
			},
		},
	},

	// Filter Set
	"filter_set": {
		Category:  CategorySimple,
		Namespace: true,
		RequiredSpec: map[string]interface{}{
			"filters": []interface{}{
				map[string]interface{}{
					"metadata": map[string]interface{}{
						"name": "test-filter",
					},
					"spec": map[string]interface{}{
						"request_headers": map[string]interface{}{
							"headers": []interface{}{
								map[string]interface{}{
									"name":       "X-Test",
									"exact":      "value",
									"invert_matcher": false,
								},
							},
						},
					},
				},
			},
		},
	},

	// Geo Location Set
	"geo_location_set": {
		Category:  CategorySimple,
		Namespace: true,
		RequiredSpec: map[string]interface{}{
			"geoip": map[string]interface{}{
				"country": []interface{}{"US"},
			},
		},
	},

	// Forwarding Class
	"forwarding_class": {
		Category:  CategorySimple,
		Namespace: true,
		RequiredSpec: map[string]interface{}{
			"priority":    "HIGH",
			"dscp_value":  46,
		},
	},
}

// ============================================================================
// CLI Flags
// ============================================================================

var (
	flagAll        = flag.Bool("all", false, "Discover defaults for all resources")
	flagResource   = flag.String("resource", "", "Discover defaults for a specific resource")
	flagPattern    = flag.String("pattern", "", "Discover defaults for resources matching pattern")
	flagOutput     = flag.String("output", "tools/api-defaults.json", "Output file path")
	flagValidate   = flag.Bool("validate", false, "Validate stored defaults against current API")
	flagVerbose    = flag.Bool("verbose", false, "Enable verbose output")
	flagDryRun     = flag.Bool("dry-run", false, "Show what would be done without making API calls")
	flagTestNS     = flag.String("test-namespace", "", "Use existing namespace for testing (skip namespace creation)")
	flagCreateNS   = flag.Bool("create-namespace", false, "Create test namespace automatically")
	flagNSPrefix   = flag.String("ns-prefix", "tf-discover", "Prefix for auto-created test namespace")
)

// ============================================================================
// Main
// ============================================================================

func main() {
	flag.Parse()

	if !*flagAll && *flagResource == "" && *flagPattern == "" && !*flagValidate {
		fmt.Println("Usage: go run tools/discover-defaults.go [options]")
		fmt.Println("")
		fmt.Println("Options:")
		flag.PrintDefaults()
		fmt.Println("")
		fmt.Println("Examples:")
		fmt.Println("  go run tools/discover-defaults.go -resource namespace")
		fmt.Println("  go run tools/discover-defaults.go -all")
		fmt.Println("  go run tools/discover-defaults.go -pattern loadbalancer")
		fmt.Println("  go run tools/discover-defaults.go -validate")
		fmt.Println("  go run tools/discover-defaults.go -dry-run -all")
		os.Exit(1)
	}

	// Determine which resources to discover (needed for dry-run)
	resourcesToDiscover := getResourcesToDiscover()

	// Handle dry-run mode without requiring API credentials
	if *flagDryRun {
		fmt.Println("Dry run - would discover defaults for:")
		fmt.Println("")
		discovered := 0
		skipped := 0
		noConfig := 0
		for _, name := range resourcesToDiscover {
			if resource.IsSkippedForAPITest(name) {
				fmt.Printf("  SKIP:      %s (%s)\n", name, resource.GetSkipReason(name))
				skipped++
			} else if _, hasConfig := ResourceConfigs[name]; hasConfig {
				fmt.Printf("  DISCOVER:  %s\n", name)
				discovered++
			} else {
				fmt.Printf("  NO CONFIG: %s (needs MinimalConfig template)\n", name)
				noConfig++
			}
		}
		fmt.Println("")
		fmt.Printf("Summary: %d to discover, %d skipped, %d need config templates\n", discovered, skipped, noConfig)
		return
	}

	// Create API client (only needed for actual discovery)
	apiClient, err := createClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating API client: %v\n", err)
		fmt.Fprintf(os.Stderr, "\nRequired environment variables:\n")
		fmt.Fprintf(os.Stderr, "  F5XC_API_URL - API URL\n")
		fmt.Fprintf(os.Stderr, "  F5XC_API_P12_FILE + F5XC_P12_PASSWORD - P12 cert auth\n")
		fmt.Fprintf(os.Stderr, "  or F5XC_API_TOKEN - Token auth\n")
		os.Exit(1)
	}

	// Initialize database
	db := &DefaultsDatabase{
		Version:     "1.0.0",
		GeneratedAt: time.Now().UTC().Format(time.RFC3339),
		APIEndpoint: os.Getenv("F5XC_API_URL"),
		Resources:   make(map[string]*DiscoveryResult),
	}

	// Load existing database if validating
	if *flagValidate {
		existingDB, err := loadDatabase(*flagOutput)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading database: %v\n", err)
			os.Exit(1)
		}
		validateDefaults(apiClient, existingDB)
		return
	}

	// Handle test namespace
	testNamespace := *flagTestNS
	createdNamespace := false

	if *flagCreateNS && testNamespace == "" {
		// Create a test namespace for discovery
		testNamespace = fmt.Sprintf("%s-%d", *flagNSPrefix, time.Now().Unix())
		fmt.Printf("Creating test namespace: %s\n", testNamespace)

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		nsRequest := map[string]interface{}{
			"metadata": map[string]interface{}{
				"name": testNamespace,
			},
			"spec": map[string]interface{}{},
		}
		var nsResp map[string]interface{}
		if err := apiClient.Post(ctx, "/api/web/namespaces", nsRequest, &nsResp); err != nil {
			cancel()
			fmt.Fprintf(os.Stderr, "Error creating test namespace: %v\n", err)
			os.Exit(1)
		}
		cancel()
		createdNamespace = true
		fmt.Printf("Created namespace: %s\n\n", testNamespace)

		// Small delay to let namespace propagate
		time.Sleep(2 * time.Second)
	}

	// Cleanup function for test namespace
	defer func() {
		if createdNamespace && testNamespace != "" {
			fmt.Printf("\nCleaning up test namespace: %s\n", testNamespace)
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()
			_ = apiClient.Delete(ctx, fmt.Sprintf("/api/web/namespaces/%s", testNamespace))
		}
	}()

	// Update the global test namespace for discoverResource
	if testNamespace != "" {
		*flagTestNS = testNamespace
	}

	// Run discovery
	fmt.Printf("Starting API default discovery for %d resources...\n\n", len(resourcesToDiscover))

	for i, resourceName := range resourcesToDiscover {
		fmt.Printf("[%d/%d] Processing %s...\n", i+1, len(resourcesToDiscover), resourceName)

		result := discoverResource(apiClient, resourceName)
		db.Resources[resourceName] = result

		switch result.Status {
		case "discovered":
			db.Discovered++
			fmt.Printf("  ✓ Discovered %d defaults\n", len(result.Defaults))
		case "skipped":
			db.Skipped++
			fmt.Printf("  ⊘ Skipped: %s\n", result.SkipReason)
		case "failed":
			db.Failed++
			fmt.Printf("  ✗ Failed: %s\n", result.Error)
		}
	}

	db.TotalResources = len(resourcesToDiscover)

	// Save database
	if err := saveDatabase(db, *flagOutput); err != nil {
		fmt.Fprintf(os.Stderr, "Error saving database: %v\n", err)
		os.Exit(1)
	}

	// Print summary
	fmt.Println("")
	fmt.Println("════════════════════════════════════════")
	fmt.Printf("Discovery Complete\n")
	fmt.Println("════════════════════════════════════════")
	fmt.Printf("Total Resources: %d\n", db.TotalResources)
	fmt.Printf("Discovered:      %d\n", db.Discovered)
	fmt.Printf("Skipped:         %d\n", db.Skipped)
	fmt.Printf("Failed:          %d\n", db.Failed)
	fmt.Printf("Output:          %s\n", *flagOutput)
}

// ============================================================================
// Client Creation
// ============================================================================

func createClient() (*client.Client, error) {
	apiURL := os.Getenv("F5XC_API_URL")
	if apiURL == "" {
		return nil, fmt.Errorf("F5XC_API_URL environment variable not set")
	}

	// Try P12 authentication first
	p12File := os.Getenv("F5XC_API_P12_FILE")
	p12Password := os.Getenv("F5XC_P12_PASSWORD")
	if p12File != "" && p12Password != "" {
		return client.NewClientWithP12(apiURL, p12File, p12Password)
	}

	// Fall back to token authentication
	apiToken := os.Getenv("F5XC_API_TOKEN")
	if apiToken != "" {
		return client.NewClient(apiURL, apiToken), nil
	}

	return nil, fmt.Errorf("no authentication credentials found")
}

// ============================================================================
// Resource Discovery
// ============================================================================

func getResourcesToDiscover() []string {
	var resources []string

	// Get all available resources from configs
	allResources := make([]string, 0, len(ResourceConfigs))
	for name := range ResourceConfigs {
		allResources = append(allResources, name)
	}
	sort.Strings(allResources)

	if *flagResource != "" {
		resources = []string{*flagResource}
	} else if *flagPattern != "" {
		for _, name := range allResources {
			if strings.Contains(name, *flagPattern) {
				resources = append(resources, name)
			}
		}
	} else if *flagAll {
		resources = allResources
	}

	return resources
}

func discoverResource(apiClient *client.Client, resourceName string) *DiscoveryResult {
	result := &DiscoveryResult{
		ResourceName: resourceName,
		Category:     resource.GetCategory(resourceName),
	}

	// Check skip list (centralized in tools/pkg/resource)
	if resource.IsSkippedForAPITest(resourceName) {
		result.Status = "skipped"
		result.SkipReason = resource.GetSkipReason(resourceName)
		return result
	}

	// Check if we have a config template
	config, hasConfig := ResourceConfigs[resourceName]
	if !hasConfig {
		result.Status = "skipped"
		result.SkipReason = "No MinimalConfig template defined"
		return result
	}

	// Build minimal request
	testName := fmt.Sprintf("tf-discover-%d", time.Now().UnixNano())
	testNamespace := *flagTestNS
	if testNamespace == "" && config.Namespace {
		// We need a test namespace - skip for now unless one is provided
		result.Status = "skipped"
		result.SkipReason = "Requires namespace - use -test-namespace flag"
		return result
	}

	request := buildRequest(testName, testNamespace, config)
	result.RequestSent = request

	// Make API call
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	response, err := createAndGetResource(ctx, apiClient, resourceName, testNamespace, request)
	if err != nil {
		result.Status = "failed"
		result.Error = err.Error()
		return result
	}

	result.ResponseGot = response
	result.Status = "discovered"
	result.DiscoveredAt = time.Now().UTC().Format(time.RFC3339)

	// Extract defaults by comparing request and response
	result.Defaults = extractDefaults(request, response, "")

	// Clean up - delete the test resource
	cleanupResource(ctx, apiClient, resourceName, testNamespace, testName)

	return result
}

func buildRequest(name, namespace string, config MinimalConfig) map[string]interface{} {
	metadata := map[string]interface{}{
		"name": name,
	}
	if config.Namespace {
		metadata["namespace"] = namespace
	}

	request := map[string]interface{}{
		"metadata": metadata,
		"spec":     config.RequiredSpec,
	}

	return request
}

// createAndGetResource creates a resource and returns the API response
func createAndGetResource(ctx context.Context, apiClient *client.Client, resourceName, namespace string, request map[string]interface{}) (map[string]interface{}, error) {
	// Get the resource name from metadata
	metadata := request["metadata"].(map[string]interface{})
	name := metadata["name"].(string)

	// Build API paths based on resource type
	var createPath, getPath string

	switch resourceName {
	case "namespace":
		createPath = "/api/web/namespaces"
		getPath = fmt.Sprintf("/api/web/namespaces/%s", name)
	case "contact":
		createPath = "/api/web/system/contacts"
		getPath = fmt.Sprintf("/api/web/system/contacts/%s", name)
	default:
		// Most resources follow this pattern
		pluralName := pluralizeResource(resourceName)
		createPath = fmt.Sprintf("/api/config/namespaces/%s/%s", namespace, pluralName)
		getPath = fmt.Sprintf("/api/config/namespaces/%s/%s/%s", namespace, pluralName, name)
	}

	if *flagVerbose {
		fmt.Printf("  POST %s\n", createPath)
		reqJSON, _ := json.MarshalIndent(request, "    ", "  ")
		fmt.Printf("    Request: %s\n", string(reqJSON))
	}

	// Create the resource
	var createResp map[string]interface{}
	if err := apiClient.Post(ctx, createPath, request, &createResp); err != nil {
		return nil, fmt.Errorf("create failed: %w", err)
	}

	// Small delay to ensure consistency
	time.Sleep(500 * time.Millisecond)

	if *flagVerbose {
		fmt.Printf("  GET %s\n", getPath)
	}

	// Read back the resource to get all defaults
	var getResp map[string]interface{}
	if err := apiClient.Get(ctx, getPath, &getResp); err != nil {
		return nil, fmt.Errorf("get failed: %w", err)
	}

	return getResp, nil
}

func cleanupResource(ctx context.Context, apiClient *client.Client, resourceName, namespace, name string) {
	var deletePath string

	switch resourceName {
	case "namespace":
		deletePath = fmt.Sprintf("/api/web/namespaces/%s", name)
	case "contact":
		deletePath = fmt.Sprintf("/api/web/system/contacts/%s", name)
	default:
		pluralName := pluralizeResource(resourceName)
		deletePath = fmt.Sprintf("/api/config/namespaces/%s/%s/%s", namespace, pluralName, name)
	}

	if *flagVerbose {
		fmt.Printf("  DELETE %s\n", deletePath)
	}

	_ = apiClient.Delete(ctx, deletePath)
}

func pluralizeResource(name string) string {
	// Special cases
	switch name {
	case "healthcheck":
		return "healthchecks"
	case "certificate":
		return "certificates"
	case "alert_policy":
		return "alert_policys" // F5 API uses this
	case "service_policy":
		return "service_policys"
	case "rate_limiter_policy":
		return "rate_limiter_policys"
	}

	// Default: add 's'
	if strings.HasSuffix(name, "s") || strings.HasSuffix(name, "x") || strings.HasSuffix(name, "ch") {
		return name + "es"
	}
	return name + "s"
}

// ============================================================================
// Default Extraction
// ============================================================================

func extractDefaults(request, response map[string]interface{}, prefix string) map[string]FieldDefault {
	defaults := make(map[string]FieldDefault)

	// Get spec from both
	reqSpec, reqOk := request["spec"].(map[string]interface{})
	respSpec, respOk := response["spec"].(map[string]interface{})

	if !reqOk || !respOk {
		return defaults
	}

	// Find fields in response that weren't in request
	compareObjects(reqSpec, respSpec, "spec", defaults)

	return defaults
}

func compareObjects(request, response map[string]interface{}, path string, defaults map[string]FieldDefault) {
	for key, respValue := range response {
		fullPath := path + "." + key

		reqValue, existsInReq := request[key]

		if !existsInReq {
			// This is a default value - wasn't sent in request
			fd := FieldDefault{
				Path:         fullPath,
				DefaultValue: respValue,
				Type:         getValueType(respValue),
			}

			// Check if it's an empty marker block
			if m, ok := respValue.(map[string]interface{}); ok && len(m) == 0 {
				fd.IsMarkerBlock = true
			}

			defaults[fullPath] = fd
		} else {
			// Both have the key - recursively compare if both are maps
			if reqMap, reqIsMap := reqValue.(map[string]interface{}); reqIsMap {
				if respMap, respIsMap := respValue.(map[string]interface{}); respIsMap {
					// Only recurse if request map was empty (marker block)
					if len(reqMap) == 0 {
						// Check if response added any fields
						for subKey, subValue := range respMap {
							subPath := fullPath + "." + subKey
							fd := FieldDefault{
								Path:         subPath,
								DefaultValue: subValue,
								Type:         getValueType(subValue),
							}
							defaults[subPath] = fd
						}
					} else {
						compareObjects(reqMap, respMap, fullPath, defaults)
					}
				}
			}
		}
	}
}

func getValueType(v interface{}) string {
	if v == nil {
		return "null"
	}
	switch reflect.TypeOf(v).Kind() {
	case reflect.String:
		return "string"
	case reflect.Int, reflect.Int64, reflect.Float64:
		return "number"
	case reflect.Bool:
		return "bool"
	case reflect.Map:
		return "object"
	case reflect.Slice, reflect.Array:
		return "array"
	default:
		return "unknown"
	}
}

// ============================================================================
// Database Operations
// ============================================================================

func loadDatabase(path string) (*DefaultsDatabase, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var db DefaultsDatabase
	if err := json.Unmarshal(data, &db); err != nil {
		return nil, err
	}

	return &db, nil
}

func saveDatabase(db *DefaultsDatabase, path string) error {
	// Ensure directory exists
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(db, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

// ============================================================================
// Validation
// ============================================================================

func validateDefaults(apiClient *client.Client, db *DefaultsDatabase) {
	fmt.Printf("Validating stored defaults against API...\n\n")

	valid := 0
	invalid := 0

	for name, result := range db.Resources {
		if result.Status != "discovered" {
			continue
		}

		fmt.Printf("Validating %s... ", name)

		// Re-run discovery
		newResult := discoverResource(apiClient, name)

		if newResult.Status != "discovered" {
			fmt.Printf("SKIP (%s)\n", newResult.SkipReason)
			continue
		}

		// Compare defaults
		if compareDefaults(result.Defaults, newResult.Defaults) {
			fmt.Printf("✓ VALID\n")
			valid++
		} else {
			fmt.Printf("✗ CHANGED\n")
			invalid++

			// Show differences
			showDifferences(result.Defaults, newResult.Defaults)
		}
	}

	fmt.Println("")
	fmt.Printf("Validation complete: %d valid, %d changed\n", valid, invalid)
}

func compareDefaults(old, new map[string]FieldDefault) bool {
	if len(old) != len(new) {
		return false
	}

	for path, oldDefault := range old {
		newDefault, exists := new[path]
		if !exists {
			return false
		}

		if !reflect.DeepEqual(oldDefault.DefaultValue, newDefault.DefaultValue) {
			return false
		}
	}

	return true
}

func showDifferences(old, new map[string]FieldDefault) {
	// Find added
	for path := range new {
		if _, exists := old[path]; !exists {
			fmt.Printf("  + %s\n", path)
		}
	}

	// Find removed
	for path := range old {
		if _, exists := new[path]; !exists {
			fmt.Printf("  - %s\n", path)
		}
	}

	// Find changed
	for path, oldDefault := range old {
		if newDefault, exists := new[path]; exists {
			if !reflect.DeepEqual(oldDefault.DefaultValue, newDefault.DefaultValue) {
				fmt.Printf("  ~ %s: %v -> %v\n", path, oldDefault.DefaultValue, newDefault.DefaultValue)
			}
		}
	}
}
