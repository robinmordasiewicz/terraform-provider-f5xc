//go:build ignore
// +build ignore

// terraform-schema-ai wraps `terraform providers schema -json` and enriches it
// with AI-friendly hints extracted from OpenAPI specifications.
//
// This tool maintains clean separation between:
//   - Provider Contract: The original Terraform schema (unchanged)
//   - Agent Heuristics: AI hints in a separate namespace
//
// Usage:
//
//	# Generate enriched schema (runs terraform internally)
//	go run tools/terraform-schema-ai/main.go
//
//	# Pipe existing schema (useful for testing)
//	terraform providers schema -json | go run tools/terraform-schema-ai/main.go -stdin
//
//	# Specify OpenAPI specs directory
//	go run tools/terraform-schema-ai/main.go -specs docs/specifications/api
//
// Output structure:
//
//	{
//	  "format_version": "1.0",
//	  "provider_schemas": { ... },  // Original Terraform schema (provider contract)
//	  "ai_hints": {                 // Agent heuristics (separate namespace)
//	    "f5xc_http_loadbalancer": {
//	      "category": "Load Balancing",
//	      "namespace_required": true,
//	      "depends_on": ["f5xc_namespace", "f5xc_origin_pool"],
//	      "one_of_groups": [...],
//	      "recommended_defaults": {...}
//	    }
//	  }
//	}
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

// TerraformSchema represents the terraform providers schema -json output
type TerraformSchema struct {
	FormatVersion   string                     `json:"format_version"`
	ProviderSchemas map[string]*ProviderSchema `json:"provider_schemas"`
}

// ProviderSchema represents a provider's schema
type ProviderSchema struct {
	Provider          *Schema            `json:"provider,omitempty"`
	ResourceSchemas   map[string]*Schema `json:"resource_schemas,omitempty"`
	DataSourceSchemas map[string]*Schema `json:"data_source_schemas,omitempty"`
}

// Schema represents a resource or data source schema
type Schema struct {
	Version int    `json:"version"`
	Block   *Block `json:"block"`
}

// Block represents a schema block
type Block struct {
	Attributes map[string]*Attribute `json:"attributes,omitempty"`
	BlockTypes map[string]*BlockType `json:"block_types,omitempty"`
}

// Attribute represents a schema attribute
type Attribute struct {
	Type        json.RawMessage `json:"type,omitempty"`
	Description string          `json:"description,omitempty"`
	Required    bool            `json:"required,omitempty"`
	Optional    bool            `json:"optional,omitempty"`
	Computed    bool            `json:"computed,omitempty"`
	Sensitive   bool            `json:"sensitive,omitempty"`
}

// BlockType represents a nested block type
type BlockType struct {
	NestingMode string `json:"nesting_mode"`
	Block       *Block `json:"block"`
	MinItems    int    `json:"min_items,omitempty"`
	MaxItems    int    `json:"max_items,omitempty"`
}

// EnrichedSchema is the output with ai_hints namespace
type EnrichedSchema struct {
	FormatVersion   string                     `json:"format_version"`
	ProviderSchemas map[string]*ProviderSchema `json:"provider_schemas"`
	AIHints         map[string]*ResourceHints  `json:"ai_hints"`
}

// ResourceHints contains AI-friendly hints for a resource
type ResourceHints struct {
	Category            string              `json:"category"`
	Description         string              `json:"description,omitempty"`
	NamespaceRequired   bool                `json:"namespace_required"`
	DependsOn           []string            `json:"depends_on,omitempty"`
	OneOfGroups         []OneOfGroup        `json:"one_of_groups,omitempty"`
	Enums               map[string][]string `json:"enums,omitempty"`
	RecommendedDefaults map[string]string   `json:"recommended_defaults,omitempty"`
	Complexity          string              `json:"complexity,omitempty"`
	UsageNotes          []string            `json:"usage_notes,omitempty"`
}

// OneOfGroup represents mutually exclusive fields
type OneOfGroup struct {
	Name        string   `json:"name"`
	Description string   `json:"description,omitempty"`
	Fields      []string `json:"fields"`
	Default     string   `json:"default,omitempty"`
	Required    bool     `json:"required"`
	BlockPath   string   `json:"block_path,omitempty"`
}

// OpenAPISpec represents a minimal OpenAPI 3.0 spec for extraction
type OpenAPISpec struct {
	Components struct {
		Schemas map[string]OpenAPISchema `json:"schemas"`
	} `json:"components"`
}

// OpenAPISchema represents an OpenAPI schema object
type OpenAPISchema struct {
	Type                 string                  `json:"type"`
	Description          string                  `json:"description"`
	Properties           map[string]OpenAPISchema `json:"properties"`
	Enum                 []interface{}           `json:"enum"`
	XVesOneOfFieldPrefix map[string]interface{}  `json:"-"` // Captured dynamically
	AllFields            map[string]interface{}  `json:"-"` // Raw fields for x-ves parsing
}

func main() {
	specsDir := flag.String("specs", "docs/specifications/api", "Directory containing OpenAPI spec files")
	useStdin := flag.Bool("stdin", false, "Read terraform schema from stdin instead of running terraform")
	outputFile := flag.String("output", "", "Output file (default: stdout)")
	verbose := flag.Bool("verbose", false, "Enable verbose output to stderr")

	flag.Parse()

	// Get terraform schema
	var schemaJSON []byte
	var err error

	if *useStdin {
		schemaJSON, err = io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading stdin: %v\n", err)
			os.Exit(1)
		}
	} else {
		schemaJSON, err = runTerraformSchema()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error running terraform: %v\n", err)
			os.Exit(1)
		}
	}

	// Parse terraform schema
	var tfSchema TerraformSchema
	if err := json.Unmarshal(schemaJSON, &tfSchema); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing terraform schema: %v\n", err)
		os.Exit(1)
	}

	// Extract AI hints from OpenAPI specs
	aiHints, err := extractAIHints(*specsDir, *verbose)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Warning: Could not extract AI hints: %v\n", err)
		aiHints = make(map[string]*ResourceHints)
	}

	// Create enriched schema
	enriched := EnrichedSchema{
		FormatVersion:   tfSchema.FormatVersion,
		ProviderSchemas: tfSchema.ProviderSchemas,
		AIHints:         aiHints,
	}

	// Output
	output, err := json.MarshalIndent(enriched, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling output: %v\n", err)
		os.Exit(1)
	}

	if *outputFile != "" {
		if err := os.WriteFile(*outputFile, output, 0644); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing output file: %v\n", err)
			os.Exit(1)
		}
		if *verbose {
			fmt.Fprintf(os.Stderr, "Wrote enriched schema to %s\n", *outputFile)
		}
	} else {
		fmt.Println(string(output))
	}
}

func runTerraformSchema() ([]byte, error) {
	cmd := exec.Command("terraform", "providers", "schema", "-json")
	return cmd.Output()
}

func extractAIHints(specsDir string, verbose bool) (map[string]*ResourceHints, error) {
	hints := make(map[string]*ResourceHints)

	// Check if specs directory exists
	if _, err := os.Stat(specsDir); os.IsNotExist(err) {
		return hints, fmt.Errorf("specs directory not found: %s", specsDir)
	}

	// Find all OpenAPI spec files (pattern: docs-cloud-f5-com.*.ves-swagger.json)
	specFiles, err := filepath.Glob(filepath.Join(specsDir, "docs-cloud-f5-com.*.ves-swagger.json"))
	if err != nil {
		return hints, err
	}

	if verbose {
		fmt.Fprintf(os.Stderr, "Found %d OpenAPI spec files\n", len(specFiles))
	}

	for _, specFile := range specFiles {
		resourceName := extractResourceNameFromPath(specFile)
		if resourceName == "" {
			continue
		}

		resourceHints, err := extractResourceHints(specFile, resourceName, verbose)
		if err != nil {
			if verbose {
				fmt.Fprintf(os.Stderr, "Warning: Could not extract hints from %s: %v\n", specFile, err)
			}
			continue
		}

		// Use f5xc_ prefix for resource names
		hints["f5xc_"+resourceName] = resourceHints
	}

	return hints, nil
}

func extractResourceNameFromPath(specPath string) string {
	// Extract resource name from path like:
	// docs-cloud-f5-com.0019.public.ves.io.schema.app_firewall.ves-swagger.json
	// docs-cloud-f5-com.0060.public.ves.io.schema.views.http_loadbalancer.ves-swagger.json
	base := filepath.Base(specPath)

	// Remove suffix
	base = strings.TrimSuffix(base, ".ves-swagger.json")

	// Find the schema part - look for "ves.io.schema." and extract what follows
	schemaIdx := strings.Index(base, "ves.io.schema.")
	if schemaIdx == -1 {
		return ""
	}

	// Get everything after "ves.io.schema."
	remainder := base[schemaIdx+len("ves.io.schema."):]

	// Handle nested paths like "views.http_loadbalancer" or "api_sec.api_discovery"
	parts := strings.Split(remainder, ".")
	if len(parts) == 0 {
		return ""
	}

	// Return the last part as the resource name
	return parts[len(parts)-1]
}

func extractResourceHints(specPath, resourceName string, verbose bool) (*ResourceHints, error) {
	data, err := os.ReadFile(specPath)
	if err != nil {
		return nil, err
	}

	// Parse as generic JSON to handle x-ves-* fields
	var rawSpec map[string]interface{}
	if err := json.Unmarshal(data, &rawSpec); err != nil {
		return nil, err
	}

	hints := &ResourceHints{
		Category:            getResourceCategory(resourceName),
		NamespaceRequired:   !isSystemResource(resourceName),
		DependsOn:           getResourceDependencies(resourceName),
		OneOfGroups:         extractOneOfGroups(rawSpec, resourceName),
		Enums:               extractEnums(rawSpec),
		RecommendedDefaults: getRecommendedDefaults(resourceName),
		Complexity:          getResourceComplexity(resourceName),
		UsageNotes:          getUsageNotes(resourceName),
	}

	if verbose {
		fmt.Fprintf(os.Stderr, "Extracted hints for %s: %d OneOf groups, %d enums\n",
			resourceName, len(hints.OneOfGroups), len(hints.Enums))
	}

	return hints, nil
}

func extractOneOfGroups(spec map[string]interface{}, resourceName string) []OneOfGroup {
	var groups []OneOfGroup

	components, ok := spec["components"].(map[string]interface{})
	if !ok {
		return groups
	}

	schemas, ok := components["schemas"].(map[string]interface{})
	if !ok {
		return groups
	}

	// Look for CreateSpecType or GlobalSpecType schemas
	for schemaName, schemaVal := range schemas {
		if !strings.Contains(schemaName, "CreateSpecType") && !strings.Contains(schemaName, "GlobalSpecType") {
			continue
		}

		schema, ok := schemaVal.(map[string]interface{})
		if !ok {
			continue
		}

		// Find x-ves-oneof-field-* annotations
		for key, val := range schema {
			if strings.HasPrefix(key, "x-ves-oneof-field-") {
				groupName := strings.TrimPrefix(key, "x-ves-oneof-field-")
				fields := extractOneOfFields(val)
				if len(fields) > 1 {
					groups = append(groups, OneOfGroup{
						Name:     groupName,
						Fields:   fields,
						Default:  determineOneOfDefault(groupName, fields),
						Required: isRequiredOneOf(groupName),
					})
				}
			}
		}
	}

	// Sort groups by name for consistent output
	sort.Slice(groups, func(i, j int) bool {
		return groups[i].Name < groups[j].Name
	})

	return groups
}

func extractOneOfFields(val interface{}) []string {
	var fields []string

	// The value can be a JSON-encoded string array or an actual array
	switch v := val.(type) {
	case string:
		// Try to parse as JSON array (e.g., "[\"field1\",\"field2\"]")
		var parsed []string
		if err := json.Unmarshal([]byte(v), &parsed); err == nil {
			fields = parsed
		}
	case []interface{}:
		for _, item := range v {
			switch field := item.(type) {
			case string:
				fields = append(fields, field)
			case map[string]interface{}:
				// Look for field name in the object
				if name, ok := field["name"].(string); ok {
					fields = append(fields, name)
				} else if fieldName, ok := field["field"].(string); ok {
					fields = append(fields, fieldName)
				}
			}
		}
	}

	sort.Strings(fields)
	return fields
}

func extractEnums(spec map[string]interface{}) map[string][]string {
	enums := make(map[string][]string)

	components, ok := spec["components"].(map[string]interface{})
	if !ok {
		return enums
	}

	schemas, ok := components["schemas"].(map[string]interface{})
	if !ok {
		return enums
	}

	// Find enum types
	for schemaName, schemaVal := range schemas {
		schema, ok := schemaVal.(map[string]interface{})
		if !ok {
			continue
		}

		if enumVals, ok := schema["enum"].([]interface{}); ok {
			var values []string
			for _, v := range enumVals {
				if s, ok := v.(string); ok && s != "" {
					values = append(values, s)
				}
			}
			if len(values) > 0 {
				// Use schema name as enum identifier
				enums[schemaName] = values
			}
		}
	}

	return enums
}

// Resource categorization
var resourceCategories = map[string]string{
	// Load Balancing
	"http_loadbalancer": "Load Balancing",
	"tcp_loadbalancer":  "Load Balancing",
	"udp_loadbalancer":  "Load Balancing",
	"dns_load_balancer": "Load Balancing",
	"cdn_loadbalancer":  "Load Balancing",
	"origin_pool":       "Load Balancing",
	"healthcheck":       "Load Balancing",
	"route":             "Load Balancing",

	// Security
	"app_firewall":                   "Security",
	"service_policy":                 "Security",
	"network_firewall":               "Security",
	"rate_limiter":                   "Security",
	"bot_defense_app_infrastructure": "Security",
	"malicious_user_mitigation":      "Security",
	"waf_exclusion_policy":           "Security",
	"enhanced_firewall_policy":       "Security",
	"forward_proxy_policy":           "Security",
	"ip_prefix_set":                  "Security",

	// Networking
	"network_connector": "Networking",
	"virtual_network":   "Networking",
	"cloud_connect":     "Networking",
	"cloud_link":        "Networking",
	"bgp":               "Networking",
	"network_interface": "Networking",
	"virtual_site":      "Networking",
	"site_mesh_group":   "Networking",

	// Sites & Infrastructure
	"securemesh_site":    "Sites",
	"securemesh_site_v2": "Sites",
	"aws_vpc_site":       "Sites",
	"azure_vnet_site":    "Sites",
	"gcp_vpc_site":       "Sites",
	"aws_tgw_site":       "Sites",
	"voltstack_site":     "Sites",

	// DNS
	"dns_zone":              "DNS",
	"dns_domain":            "DNS",
	"dns_lb_pool":           "DNS",
	"dns_lb_health_check":   "DNS",
	"dns_compliance_checks": "DNS",

	// Kubernetes
	"k8s_cluster":              "Kubernetes",
	"virtual_k8s":              "Kubernetes",
	"k8s_cluster_role":         "Kubernetes",
	"k8s_cluster_role_binding": "Kubernetes",
	"k8s_pod_security_policy":  "Kubernetes",
	"container_registry":       "Kubernetes",

	// Authentication & Credentials
	"authentication":    "Authentication",
	"cloud_credentials": "Authentication",
	"api_credential":    "Authentication",
	"token":             "Authentication",
	"secret_policy":     "Authentication",

	// Certificates
	"certificate":       "Certificates",
	"certificate_chain": "Certificates",
	"trusted_ca_list":   "Certificates",

	// Monitoring
	"log_receiver":        "Monitoring",
	"global_log_receiver": "Monitoring",
	"alert_policy":        "Monitoring",
	"alert_receiver":      "Monitoring",

	// API Security
	"api_definition": "API Security",
	"api_discovery":  "API Security",

	// Organization
	"namespace":      "Organization",
	"tenant":         "Organization",
	"role":           "Organization",
	"user":           "Organization",
	"allowed_tenant": "Organization",
}

func getResourceCategory(resourceName string) string {
	if cat, ok := resourceCategories[resourceName]; ok {
		return cat
	}
	return "Other"
}

// System resources that don't require namespace
var systemResources = map[string]bool{
	"namespace":         true,
	"tenant":            true,
	"role":              true,
	"user":              true,
	"cloud_credentials": true,
	"token":             true,
}

func isSystemResource(resourceName string) bool {
	return systemResources[resourceName]
}

// Resource dependencies
var resourceDependencies = map[string][]string{
	"http_loadbalancer": {"f5xc_namespace", "f5xc_origin_pool"},
	"tcp_loadbalancer":  {"f5xc_namespace", "f5xc_origin_pool"},
	"udp_loadbalancer":  {"f5xc_namespace", "f5xc_origin_pool"},
	"origin_pool":       {"f5xc_namespace", "f5xc_healthcheck"},
	"healthcheck":       {"f5xc_namespace"},
	"route":             {"f5xc_namespace", "f5xc_http_loadbalancer"},
	"app_firewall":      {"f5xc_namespace"},
	"service_policy":    {"f5xc_namespace"},
	"rate_limiter":      {"f5xc_namespace"},
	"certificate":       {"f5xc_namespace"},
	"api_definition":    {"f5xc_namespace"},
	"dns_zone":          {"f5xc_namespace"},
	"virtual_network":   {"f5xc_namespace"},
	"aws_vpc_site":      {"f5xc_cloud_credentials"},
	"azure_vnet_site":   {"f5xc_cloud_credentials"},
	"gcp_vpc_site":      {"f5xc_cloud_credentials"},
}

func getResourceDependencies(resourceName string) []string {
	if deps, ok := resourceDependencies[resourceName]; ok {
		return deps
	}
	if !isSystemResource(resourceName) {
		return []string{"f5xc_namespace"}
	}
	return nil
}

// Recommended defaults for OneOf groups
var oneOfDefaults = map[string]string{
	"advertise_choice":           "advertise_on_public_default_vip",
	"loadbalancer_type":          "https_auto_cert",
	"hash_policy_choice":         "round_robin",
	"waf_choice":                 "disable_waf",
	"challenge_type":             "no_challenge",
	"rate_limit_choice":          "disable_rate_limit",
	"service_policy_choice":      "no_service_policies",
	"tls_choice":                 "no_tls",
	"bot_defense_choice":         "disable_bot_defense",
	"api_definition_choice":      "disable_api_definition",
	"api_discovery_choice":       "disable_api_discovery",
	"ip_reputation_choice":       "disable_ip_reputation",
	"malware_protection":         "disable_malware_protection",
	"client_side_defense_choice": "disable_client_side_defense",
}

func determineOneOfDefault(groupName string, fields []string) string {
	// Check explicit defaults first
	if defaultVal, ok := oneOfDefaults[groupName]; ok {
		for _, f := range fields {
			if f == defaultVal {
				return defaultVal
			}
		}
	}

	// Heuristics for common patterns
	for _, f := range fields {
		if strings.Contains(f, "default") {
			return f
		}
	}

	for _, f := range fields {
		if strings.HasPrefix(f, "no_") || strings.HasPrefix(f, "disable_") {
			return f
		}
	}

	return ""
}

func isRequiredOneOf(groupName string) bool {
	// These groups typically require a choice
	requiredGroups := map[string]bool{
		"advertise_choice":   true,
		"loadbalancer_type":  true,
		"hash_policy_choice": true,
	}
	return requiredGroups[groupName]
}

// Resource-specific recommended defaults
func getRecommendedDefaults(resourceName string) map[string]string {
	defaults := map[string]map[string]string{
		"http_loadbalancer": {
			"advertise_choice":  "advertise_on_public_default_vip",
			"loadbalancer_type": "https_auto_cert",
		},
		"origin_pool": {
			"loadbalancer_algorithm": "ROUND_ROBIN",
		},
	}

	if d, ok := defaults[resourceName]; ok {
		return d
	}
	return nil
}

// Resource complexity rating
func getResourceComplexity(resourceName string) string {
	complex := map[string]bool{
		"http_loadbalancer": true,
		"tcp_loadbalancer":  true,
		"aws_vpc_site":      true,
		"azure_vnet_site":   true,
		"gcp_vpc_site":      true,
		"service_policy":    true,
	}

	simple := map[string]bool{
		"namespace":   true,
		"healthcheck": true,
		"certificate": true,
	}

	if complex[resourceName] {
		return "complex"
	}
	if simple[resourceName] {
		return "simple"
	}
	return "moderate"
}

// Usage notes for specific resources
func getUsageNotes(resourceName string) []string {
	notes := map[string][]string{
		"http_loadbalancer": {
			"Use https_auto_cert for automatic TLS certificate management",
			"advertise_on_public_default_vip uses shared Volterra VIP",
			"Origin pools must be created before the load balancer",
		},
		"origin_pool": {
			"Health checks are optional but recommended for production",
			"Private origins require site connectivity",
		},
		"namespace": {
			"Most resources require a namespace to be created first",
			"Namespace names must be unique within a tenant",
		},
		"app_firewall": {
			"WAF policies can be attached to HTTP load balancers",
			"Consider using allow_all_response_codes for initial testing",
		},
	}

	return notes[resourceName]
}

// Regex for extracting resource name patterns
var resourceNamePattern = regexp.MustCompile(`[a-z][a-z0-9_]*`)

func init() {
	// Ensure regex is compiled
	_ = resourceNamePattern
}
