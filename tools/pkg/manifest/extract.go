package manifest

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

// OpenAPI3Spec represents an OpenAPI 3.x specification
type OpenAPI3Spec struct {
	OpenAPI    string                 `json:"openapi"`
	Info       map[string]interface{} `json:"info"`
	Paths      map[string]interface{} `json:"paths"`
	Components Components             `json:"components"`
}

// Components represents the components section of an OpenAPI spec
type Components struct {
	Schemas map[string]SchemaDefinition `json:"schemas"`
}

// SchemaDefinition represents a schema definition in OpenAPI
type SchemaDefinition struct {
	Type                 string                      `json:"type"`
	Description          string                      `json:"description"`
	Title                string                      `json:"title"`
	Format               string                      `json:"format"`
	Enum                 []interface{}               `json:"enum"`
	Default              interface{}                 `json:"default"`
	Properties           map[string]SchemaDefinition `json:"properties"`
	Items                *SchemaDefinition           `json:"items"`
	Ref                  string                      `json:"$ref"`
	Required             []string                    `json:"required"`
	XDisplayName         string                      `json:"x-displayname"`
	XVesExample          string                      `json:"x-ves-example"`
	XVesValidationRules  map[string]string           `json:"x-ves-validation-rules"`
	XVesProtoMessage     string                      `json:"x-ves-proto-message"`
	AdditionalProperties interface{}                 `json:"additionalProperties"`
}

// Extractor handles extracting constraint information from OpenAPI specs
type Extractor struct {
	specDir      string
	schemaCache  map[string]SchemaDefinition
	rawSpecCache map[string]map[string]interface{}
	verbose      bool
}

// NewExtractor creates a new constraint extractor
func NewExtractor(specDir string, verbose bool) *Extractor {
	return &Extractor{
		specDir:      specDir,
		schemaCache:  make(map[string]SchemaDefinition),
		rawSpecCache: make(map[string]map[string]interface{}),
		verbose:      verbose,
	}
}

// ExtractedResource contains all extracted information for a single resource
type ExtractedResource struct {
	Name              string
	Category          string
	Description       string
	APIPath           string
	RequiresNamespace bool
	OneOfGroups       map[string][]string // group name -> field names
	NestedOneOfGroups map[string]map[string][]string // block path -> group name -> field names
	Properties        map[string]*ExtractedProperty
	SpecFile          string
}

// ExtractedProperty contains information about a single property
type ExtractedProperty struct {
	Name        string
	Type        string
	Description string
	Required    bool
	OneOfGroup  string
	EnumValues  []string
	Default     interface{}
	Validation  *ValidationRules
}

// ExtractFromAllSpecs processes all OpenAPI spec files and extracts constraint information
func (e *Extractor) ExtractFromAllSpecs() ([]*ExtractedResource, error) {
	pattern := filepath.Join(e.specDir, "docs-cloud-f5-com.*.ves-swagger.json")
	specFiles, err := filepath.Glob(pattern)
	if err != nil {
		return nil, fmt.Errorf("error finding spec files: %w", err)
	}

	if len(specFiles) == 0 {
		return nil, fmt.Errorf("no spec files found matching pattern: %s", pattern)
	}

	if e.verbose {
		fmt.Printf("Found %d OpenAPI specification files\n", len(specFiles))
	}

	var resources []*ExtractedResource
	for _, specFile := range specFiles {
		resource, err := e.extractFromSpec(specFile)
		if err != nil {
			if e.verbose {
				fmt.Printf("Skipping %s: %v\n", filepath.Base(specFile), err)
			}
			continue
		}
		if resource != nil {
			resources = append(resources, resource)
		}
	}

	return resources, nil
}

// extractFromSpec processes a single OpenAPI spec file
func (e *Extractor) extractFromSpec(specFile string) (*ExtractedResource, error) {
	data, err := os.ReadFile(specFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read spec: %w", err)
	}

	var spec OpenAPI3Spec
	if err := json.Unmarshal(data, &spec); err != nil {
		return nil, fmt.Errorf("failed to parse spec: %w", err)
	}

	// Cache schemas
	for name, schema := range spec.Components.Schemas {
		e.schemaCache[name] = schema
	}

	// Parse raw JSON for x-ves-oneof-field extraction
	var rawSpec map[string]interface{}
	if err := json.Unmarshal(data, &rawSpec); err == nil {
		if components, ok := rawSpec["components"].(map[string]interface{}); ok {
			if schemas, ok := components["schemas"].(map[string]interface{}); ok {
				for name, schema := range schemas {
					if schemaMap, ok := schema.(map[string]interface{}); ok {
						e.rawSpecCache[name] = schemaMap
					}
				}
			}
		}
	}

	// Find CreateSpecType schema
	var createSpecKey string
	for key := range spec.Components.Schemas {
		if strings.HasSuffix(key, "CreateSpecType") {
			createSpecKey = key
			break
		}
	}

	if createSpecKey == "" {
		return nil, fmt.Errorf("no CreateSpecType found")
	}

	// Extract resource name from schema key
	resourceName := extractResourceName(createSpecKey)
	if resourceName == "" {
		return nil, fmt.Errorf("could not extract resource name from %s", createSpecKey)
	}

	// Extract category from spec file path
	category := extractCategory(specFile)

	// Extract API path
	apiPath := extractAPIPath(spec.Paths, resourceName)

	// Extract OneOf groups
	oneOfGroups := e.extractOneOfGroups(createSpecKey)

	// Extract properties
	properties := e.extractProperties(createSpecKey, oneOfGroups)

	// Determine if namespace is required
	requiresNamespace := strings.Contains(apiPath, "namespaces")

	// Get description
	description := ""
	if schema, ok := spec.Components.Schemas[createSpecKey]; ok {
		description = schema.Description
	}

	return &ExtractedResource{
		Name:              resourceName,
		Category:          category,
		Description:       description,
		APIPath:           apiPath,
		RequiresNamespace: requiresNamespace,
		OneOfGroups:       oneOfGroups,
		Properties:        properties,
		SpecFile:          filepath.Base(specFile),
	}, nil
}

// extractOneOfGroups extracts x-ves-oneof-field annotations from the raw schema
func (e *Extractor) extractOneOfGroups(schemaKey string) map[string][]string {
	oneOfGroups := make(map[string][]string)

	rawSchema, ok := e.rawSpecCache[schemaKey]
	if !ok {
		return oneOfGroups
	}

	for key, value := range rawSchema {
		if strings.HasPrefix(key, "x-ves-oneof-field-") {
			groupName := strings.TrimPrefix(key, "x-ves-oneof-field-")
			switch v := value.(type) {
			case string:
				// Parse JSON array format: "[\"field1\",\"field2\"]"
				v = strings.Trim(v, "[]")
				fields := strings.Split(v, ",")
				for i, f := range fields {
					fields[i] = strings.Trim(strings.TrimSpace(f), "\"")
				}
				oneOfGroups[groupName] = fields
			case []interface{}:
				fields := make([]string, len(v))
				for i, f := range v {
					if s, ok := f.(string); ok {
						fields[i] = s
					}
				}
				oneOfGroups[groupName] = fields
			}
		}
	}

	return oneOfGroups
}

// extractProperties extracts property information from the schema
func (e *Extractor) extractProperties(schemaKey string, oneOfGroups map[string][]string) map[string]*ExtractedProperty {
	props := make(map[string]*ExtractedProperty)

	schema, ok := e.schemaCache[schemaKey]
	if !ok {
		return props
	}

	// Build reverse mapping: field -> OneOf group
	fieldToOneOf := make(map[string]string)
	for groupName, fields := range oneOfGroups {
		for _, field := range fields {
			fieldToOneOf[field] = groupName
		}
	}

	// Build required set
	requiredSet := make(map[string]bool)
	for _, r := range schema.Required {
		requiredSet[r] = true
	}

	for propName, propDef := range schema.Properties {
		prop := &ExtractedProperty{
			Name:        propName,
			Type:        propDef.Type,
			Description: propDef.Description,
			Required:    requiredSet[propName],
			Default:     propDef.Default,
		}

		// Check if property belongs to a OneOf group
		if group, ok := fieldToOneOf[propName]; ok {
			prop.OneOfGroup = group
		}

		// Extract enum values
		if len(propDef.Enum) > 0 {
			for _, e := range propDef.Enum {
				if s, ok := e.(string); ok {
					prop.EnumValues = append(prop.EnumValues, s)
				}
			}
		}

		// Extract validation rules
		if len(propDef.XVesValidationRules) > 0 {
			prop.Validation = &ValidationRules{}
			if pattern, ok := propDef.XVesValidationRules["pattern"]; ok {
				prop.Validation.Pattern = pattern
			}
		}

		props[propName] = prop
	}

	return props
}

// extractResourceName extracts the resource name from a schema key
func extractResourceName(schemaKey string) string {
	// Schema keys look like: viewshttp_loadbalancerCreateSpecType
	// We need to extract: http_loadbalancer

	// Remove common suffixes
	name := schemaKey
	for _, suffix := range []string{"CreateSpecType", "GetSpecType", "ReplaceSpecType", "GlobalSpecType"} {
		name = strings.TrimSuffix(name, suffix)
	}

	// Remove common prefixes
	for _, prefix := range []string{"views", "schema", "config"} {
		name = strings.TrimPrefix(name, prefix)
	}

	// Convert camelCase to snake_case if needed
	if !strings.Contains(name, "_") {
		name = toSnakeCase(name)
	}

	return name
}

// toSnakeCase converts CamelCase to snake_case
func toSnakeCase(s string) string {
	var result strings.Builder
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result.WriteRune('_')
		}
		result.WriteRune(r)
	}
	return strings.ToLower(result.String())
}

// extractCategory determines the resource category from the spec file path
func extractCategory(specFile string) string {
	filename := filepath.Base(specFile)

	// Pattern: docs-cloud-f5-com.XXXX.public.ves.io.schema.CATEGORY.RESOURCE.ves-swagger.json
	re := regexp.MustCompile(`ves\.io\.schema\.([^.]+)\.`)
	matches := re.FindStringSubmatch(filename)
	if len(matches) > 1 {
		category := matches[1]
		// Map common prefixes to friendly names
		categoryMap := map[string]string{
			"views":              "Load Balancing",
			"app_firewall":       "Security",
			"service_policy":     "Security",
			"rate_limiter":       "Security",
			"network":            "Networking",
			"dns":                "DNS",
			"certificate":        "Certificates",
			"namespace":          "Organization",
			"site":               "Sites",
			"cloud":              "Cloud Resources",
			"tenant_management":  "Organization",
			"pbac":               "Access Control",
			"api_sec":            "API Security",
			"bot_defense":        "Security",
			"malicious_user":     "Security",
			"waf":                "Security",
		}
		if mapped, ok := categoryMap[category]; ok {
			return mapped
		}
		// Title case the category
		return strings.Title(strings.ReplaceAll(category, "_", " "))
	}

	return "Uncategorized"
}

// extractAPIPath extracts the API path from the paths section
func extractAPIPath(paths map[string]interface{}, resourceName string) string {
	// Look for the create endpoint
	for path := range paths {
		if strings.Contains(path, resourceName) || strings.Contains(path, strings.ReplaceAll(resourceName, "_", "")) {
			return path
		}
	}

	// Try to construct a default path
	return fmt.Sprintf("/api/config/namespaces/{namespace}/%ss", resourceName)
}

// ConvertToManifest converts extracted resources to the AI manifest format
func ConvertToManifest(resources []*ExtractedResource, providerVersion string) *AIManifest {
	manifest := &AIManifest{
		Version:         "1.0.0",
		Provider:        "f5xc",
		ProviderVersion: providerVersion,
		TotalResources:  len(resources),
		Resources:       make(map[string]*ResourceManifest),
		Categories:      make(map[string][]string),
	}

	// Sort resources by name for consistent output
	sort.Slice(resources, func(i, j int) bool {
		return resources[i].Name < resources[j].Name
	})

	for _, res := range resources {
		rm := &ResourceManifest{
			Name:              res.Name,
			Category:          res.Category,
			Description:       res.Description,
			APIPath:           res.APIPath,
			RequiresNamespace: res.RequiresNamespace,
			Constraints:       &ConstraintManifest{},
			Attributes:        make(map[string]*AttributeManifest),
		}

		// Convert OneOf groups
		for groupName, fields := range res.OneOfGroups {
			oneOf := OneOfGroup{
				GroupID:  groupName,
				Fields:   fields,
				Required: true, // Most F5XC OneOf groups are required
			}

			// Set default choice (first field alphabetically as a heuristic)
			if len(fields) > 0 {
				sortedFields := make([]string, len(fields))
				copy(sortedFields, fields)
				sort.Strings(sortedFields)
				oneOf.DefaultChoice = determineDefaultChoice(groupName, sortedFields)
			}

			rm.Constraints.OneOfGroups = append(rm.Constraints.OneOfGroups, oneOf)
		}

		// Sort OneOf groups by ID for consistent output
		sort.Slice(rm.Constraints.OneOfGroups, func(i, j int) bool {
			return rm.Constraints.OneOfGroups[i].GroupID < rm.Constraints.OneOfGroups[j].GroupID
		})

		// Convert properties to attributes
		for propName, prop := range res.Properties {
			attr := &AttributeManifest{
				Name:        propName,
				Type:        prop.Type,
				Description: prop.Description,
				Required:    prop.Required,
				Optional:    !prop.Required,
				OneOfGroup:  prop.OneOfGroup,
				EnumValues:  prop.EnumValues,
			}
			if prop.Default != nil {
				attr.DefaultValue = prop.Default
			}
			if prop.Validation != nil {
				attr.Validation = prop.Validation
			}
			rm.Attributes[propName] = attr
		}

		// Add AI hints with common patterns
		rm.AIHints = generateAIHints(res)

		manifest.Resources[res.Name] = rm

		// Track categories
		manifest.Categories[res.Category] = append(manifest.Categories[res.Category], res.Name)
	}

	// Sort category entries
	for cat := range manifest.Categories {
		sort.Strings(manifest.Categories[cat])
	}

	// Add global hints
	manifest.GlobalHints = generateGlobalHints(resources)

	return manifest
}

// determineDefaultChoice determines a sensible default for a OneOf group
func determineDefaultChoice(groupName string, fields []string) string {
	// Common patterns for F5 XC resources
	defaultPatterns := map[string]string{
		"advertise_choice":     "advertise_on_public_default_vip",
		"loadbalancer_type":    "https_auto_cert",
		"hash_policy_choice":   "round_robin",
		"waf_choice":           "disable_waf",
		"challenge_type":       "no_challenge",
		"rate_limit_choice":    "disable_rate_limit",
		"service_policy_choice": "no_service_policies",
		"tls_choice":           "no_tls",
	}

	if defaultVal, ok := defaultPatterns[groupName]; ok {
		// Verify the default is actually in the fields
		for _, f := range fields {
			if f == defaultVal {
				return defaultVal
			}
		}
	}

	// Heuristics for common patterns
	for _, f := range fields {
		// Prefer "default" variants
		if strings.Contains(f, "default") {
			return f
		}
		// Prefer "no_" or "disable_" for optional features
		if strings.HasPrefix(f, "no_") || strings.HasPrefix(f, "disable_") {
			return f
		}
	}

	// Fall back to first field
	if len(fields) > 0 {
		return fields[0]
	}

	return ""
}

// generateAIHints generates AI hints for a resource
func generateAIHints(res *ExtractedResource) *ResourceAIHints {
	hints := &ResourceAIHints{
		RecommendedDefaults: make(map[string]string),
	}

	// Add recommended defaults for each OneOf group
	for groupName, fields := range res.OneOfGroups {
		sortedFields := make([]string, len(fields))
		copy(sortedFields, fields)
		sort.Strings(sortedFields)
		hints.RecommendedDefaults[groupName] = determineDefaultChoice(groupName, sortedFields)
	}

	// Determine complexity based on number of OneOf groups and properties
	oneOfCount := len(res.OneOfGroups)
	propCount := len(res.Properties)

	if oneOfCount > 10 || propCount > 50 {
		hints.Complexity = "complex"
	} else if oneOfCount > 3 || propCount > 20 {
		hints.Complexity = "moderate"
	} else {
		hints.Complexity = "simple"
	}

	// Add common patterns based on resource type
	hints.CommonPatterns = getCommonPatterns(res.Name)

	// Add dependency order
	hints.DependencyOrder = getDependencyOrder(res.Name)

	// Add related resources
	hints.RelatedResources = getRelatedResources(res.Name)

	return hints
}

// getCommonPatterns returns common configuration patterns for a resource
func getCommonPatterns(resourceName string) []string {
	patterns := map[string][]string{
		"http_loadbalancer": {
			"https_auto_cert + app_firewall + round_robin",
			"advertise_on_public_default_vip + default_route_pools",
			"rate_limiter + service_policies for security",
		},
		"origin_pool": {
			"public_ip or private_ip origin servers",
			"no_tls for internal, use_tls for external",
			"healthcheck for production deployments",
		},
		"app_firewall": {
			"blocking mode for production",
			"default_detection_settings for standard protection",
		},
		"namespace": {
			"create before other namespaced resources",
		},
	}

	if p, ok := patterns[resourceName]; ok {
		return p
	}
	return nil
}

// getDependencyOrder returns the creation order for a resource
func getDependencyOrder(resourceName string) []string {
	dependencies := map[string][]string{
		"http_loadbalancer": {"namespace", "origin_pool", "healthcheck", "app_firewall"},
		"tcp_loadbalancer":  {"namespace", "origin_pool", "healthcheck"},
		"origin_pool":       {"namespace", "healthcheck"},
		"app_firewall":      {"namespace"},
		"service_policy":    {"namespace"},
		"healthcheck":       {"namespace"},
	}

	if deps, ok := dependencies[resourceName]; ok {
		return deps
	}
	return nil
}

// getRelatedResources returns commonly used related resources
func getRelatedResources(resourceName string) []string {
	related := map[string][]string{
		"http_loadbalancer": {"origin_pool", "healthcheck", "app_firewall", "service_policy", "rate_limiter"},
		"origin_pool":       {"healthcheck", "http_loadbalancer", "tcp_loadbalancer"},
		"app_firewall":      {"http_loadbalancer", "service_policy"},
		"namespace":         {"http_loadbalancer", "origin_pool", "app_firewall"},
	}

	if rel, ok := related[resourceName]; ok {
		return rel
	}
	return nil
}

// generateGlobalHints generates provider-wide AI hints
func generateGlobalHints(resources []*ExtractedResource) *GlobalAIHints {
	return &GlobalAIHints{
		CommonWorkflows: []WorkflowHint{
			{
				Name:          "Basic HTTP Load Balancer",
				Description:   "Create a simple HTTP load balancer with origin pool",
				ResourceOrder: []string{"namespace", "healthcheck", "origin_pool", "http_loadbalancer"},
			},
			{
				Name:          "Secure HTTP Load Balancer with WAF",
				Description:   "HTTP load balancer with TLS, WAF protection, and rate limiting",
				ResourceOrder: []string{"namespace", "healthcheck", "origin_pool", "app_firewall", "rate_limiter", "http_loadbalancer"},
			},
			{
				Name:          "Multi-Cloud Site Deployment",
				Description:   "Deploy F5 XC sites across cloud providers",
				ResourceOrder: []string{"cloud_credentials", "aws_vpc_site", "azure_vnet_site", "gcp_vpc_site"},
			},
		},
		ResourcePriority: []string{
			"namespace",
			"http_loadbalancer",
			"origin_pool",
			"healthcheck",
			"app_firewall",
			"service_policy",
			"tcp_loadbalancer",
			"dns_zone",
			"certificate",
		},
		NamingConventions: map[string]string{
			"pattern":     "^[a-z][a-z0-9-]*[a-z0-9]$",
			"max_length":  "63",
			"description": "Lowercase alphanumeric with hyphens, must start with letter",
		},
	}
}
