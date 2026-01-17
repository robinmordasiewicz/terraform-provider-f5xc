// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

//go:build ignore
// +build ignore

// generate-all-schemas.go - Batch generator for all F5 XC Terraform resources
// This tool processes all OpenAPI spec files and generates comprehensive Terraform schemas.
//
// CI/CD Integration:
//   Changes to this file trigger the generate.yml workflow which regenerates all
//   provider resources from the latest OpenAPI specifications.
//
// Pipeline Verification: 2025-12-15 - Full end-to-end workflow validation #488
//
// Usage: go run tools/generate-all-schemas.go [--spec-dir=/path/to/specs] [--dry-run]
//
// Environment Variables:
//   F5XC_SPEC_DIR - Directory containing OpenAPI spec files (default: /tmp)

package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"text/template"
	"time"

	"github.com/f5xc/terraform-provider-f5xc/tools/pkg/namespace"
	"github.com/f5xc/terraform-provider-f5xc/tools/pkg/naming"
	"github.com/f5xc/terraform-provider-f5xc/tools/pkg/openapi"
)

// Configuration
var (
	specDir   string
	dryRun    bool
	outputDir string
	clientDir string
	verbose   bool
)

// OpenAPI3Spec represents an OpenAPI 3.x specification
type OpenAPI3Spec struct {
	OpenAPI    string                 `json:"openapi"`
	Info       map[string]interface{} `json:"info"`
	Paths      map[string]interface{} `json:"paths"`
	Components Components             `json:"components"`
}

type Components struct {
	Schemas map[string]SchemaDefinition `json:"schemas"`
}

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
	AdditionalProperties interface{}                 `json:"additionalProperties"`

	// Original F5 vendor extensions (x-ves-*) - technical metadata from upstream
	XDisplayName        string            `json:"x-displayname"`
	XVesExample         string            `json:"x-ves-example"`
	XVesValidationRules map[string]string `json:"x-ves-validation-rules"`
	XVesProtoMessage    string            `json:"x-ves-proto-message"`

	// Enrichment extensions (x-f5xc-*) - added by f5xc-api-enriched repository
	XF5XCCategory         string   `json:"x-f5xc-category"`
	XF5XCRequiresTier     string   `json:"x-f5xc-requires-tier"`
	XF5XCComplexity       string   `json:"x-f5xc-complexity"`
	XF5XCExample          string   `json:"x-f5xc-example"`
	XF5XCDescriptionShort string   `json:"x-f5xc-description-short"`
	XF5XCDescriptionMed   string   `json:"x-f5xc-description-medium"`
	XF5XCUseCases         []string `json:"x-f5xc-use-cases"`
	XF5XCRelatedDomains   []string `json:"x-f5xc-related-domains"`
	XF5XCIsPreview        bool     `json:"x-f5xc-is-preview"`
	XF5XCServerDefault      bool        `json:"x-f5xc-server-default"`      // True if server applies default when omitted
	XF5XCRecommendedValue   interface{} `json:"x-f5xc-recommended-value"`   // Recommended value for required fields
}

type TerraformAttribute struct {
	Name               string
	GoName             string
	TfsdkTag           string
	Type               string
	ElementType        string
	Description        string
	Required           bool
	Optional           bool
	Computed           bool
	Sensitive          bool
	NestedAttributes   []TerraformAttribute
	NestedBlockType    string
	IsBlock            bool
	OneOfGroup         string
	PlanModifier       string
	MaxDepth           int    // Track recursion depth to prevent infinite loops
	IsSpecField        bool   // True if this is a spec field (not metadata)
	JsonName           string // JSON field name from OpenAPI for API marshaling
	GoType             string // Go type for client struct generation
	UseDomainValidator bool        // True if name field should use DomainValidator (for DNS resources)
	ServerDefault      bool        // True if server applies default when field is omitted (from x-f5xc-server-default)
	RecommendedValue   interface{} // Recommended value for required fields without server defaults
}

type ResourceTemplate struct {
	Name                   string
	TitleCase              string
	APIPath                string
	APIPathPlural          string
	APIPathItem            string // Path for single item operations (get/update/delete)
	HasNamespaceInPath     bool   // Whether API path contains namespace segment
	Description            string
	Attributes             []TerraformAttribute
	OneOfGroups            map[string][]string
	HasComplexSpec         bool
	RequiredAttributes     []string
	OptionalAttributes     []string
	ComputedAttributes     []string
	ExampleUsage           string // HCL example for documentation
	APIDocsURL             string // Link to F5 XC API documentation
	UsesBoolPlanModifier   bool   // True if any bool attribute uses a plan modifier
	UsesInt64PlanModifier  bool   // True if any int64 attribute uses a plan modifier
	UsesStringPlanModifier bool   // True if any string attribute uses a plan modifier
	HasBlocks              bool   // True if the resource has any nested blocks
}

type GenerationResult struct {
	ResourceName string
	Success      bool
	Error        string
	AttrCount    int
	BlockCount   int
}

// =============================================================================
// METADATA COLLECTION TYPES (for MCP Server consumption)
// =============================================================================

// MetadataCollection holds all extracted metadata for JSON output
type MetadataCollection struct {
	GeneratedAt string                       `json:"generated_at"`
	Version     string                       `json:"version"`
	Resources   map[string]*ResourceMetadata `json:"resources"`
}

// ResourceMetadata contains complete metadata for a single resource
type ResourceMetadata struct {
	Description  string                        `json:"description"`
	Category     string                        `json:"category,omitempty"`
	Tier         string                        `json:"tier,omitempty"`
	ImportFormat string                        `json:"import_format,omitempty"` // "namespace/name" or "name"
	OneOfGroups  map[string]*OneOfGroupInfo    `json:"oneof_groups,omitempty"`
	Attributes   map[string]*AttributeMetadata `json:"attributes"`
	Dependencies *DependencyInfo               `json:"dependencies,omitempty"`
}

// OneOfGroupInfo represents a mutually exclusive field group
type OneOfGroupInfo struct {
	Fields      []string `json:"fields"`
	Description string   `json:"description,omitempty"`
	Default     string   `json:"default,omitempty"` // Recommended default field
}

// AttributeMetadata contains metadata for a single attribute
type AttributeMetadata struct {
	Type         string      `json:"type"`
	Required     bool        `json:"required"`
	Optional     bool        `json:"optional,omitempty"`
	Computed     bool        `json:"computed,omitempty"`
	Sensitive    bool        `json:"sensitive,omitempty"`
	IsBlock      bool        `json:"is_block,omitempty"`
	PlanModifier string      `json:"plan_modifier,omitempty"`
	Validation   string      `json:"validation,omitempty"`
	Enum         []string    `json:"enum,omitempty"`
	Default      interface{} `json:"default,omitempty"`
	OneOfGroup        string      `json:"oneof_group,omitempty"`
	Description       string      `json:"description"`
	ServerDefault     bool        `json:"server_default,omitempty"`      // True if server applies default when field is omitted
	ServerDefaultDesc string      `json:"server_default_desc,omitempty"` // Description of what server default behavior is
	RecommendedValue  interface{} `json:"recommended_value,omitempty"`   // Suggested value for required fields
}

// DependencyInfo holds resource relationship information
type DependencyInfo struct {
	References   []string `json:"references,omitempty"`    // Resources this resource references
	ReferencedBy []string `json:"referenced_by,omitempty"` // Resources that reference this
}

// =============================================================================
// OPERATION METADATA TYPES (v2.0.33 extensions - for MCP Server consumption)
// =============================================================================

// OperationsMetadataCollection holds all operation-level metadata for JSON output
type OperationsMetadataCollection struct {
	GeneratedAt string                           `json:"generated_at"`
	Version     string                           `json:"version"`
	Resources   map[string]*ResourceOperationInfo `json:"resources"`
}

// ResourceOperationInfo contains operation metadata for a single resource
type ResourceOperationInfo struct {
	Resource     string                        `json:"resource"`
	BasePath     string                        `json:"base_path,omitempty"`
	Operations   map[string]*OperationMetadata `json:"operations"`             // key: "create", "read", "update", "delete", "list"
	BestPractices *BestPracticesInfo           `json:"best_practices,omitempty"`
	Workflows    []*GuidedWorkflowInfo         `json:"guided_workflows,omitempty"`
}

// OperationMetadata represents operation-level x-f5xc-* extensions
type OperationMetadata struct {
	Method               string            `json:"method"`                           // HTTP method (POST, GET, PUT, DELETE)
	Path                 string            `json:"path"`                             // API path
	DangerLevel          string            `json:"danger_level,omitempty"`           // low, medium, high, critical
	DiscoveredRespTime   *ResponseTimeInfo `json:"discovered_response_time,omitempty"`
	RequiredFields       []string          `json:"required_fields,omitempty"`
	ConfirmationRequired bool              `json:"confirmation_required,omitempty"`
	SideEffects          *SideEffectsInfo  `json:"side_effects,omitempty"`
	Purpose              string            `json:"purpose,omitempty"`                // From x-f5xc-operation-metadata.purpose
}

// ResponseTimeInfo represents the x-f5xc-discovered-response-time extension
type ResponseTimeInfo struct {
	P50Ms       int    `json:"p50_ms"`
	P95Ms       int    `json:"p95_ms"`
	P99Ms       int    `json:"p99_ms"`
	SampleCount int    `json:"sample_count"`
	Source      string `json:"source"` // "measured" or "estimate"
}

// SideEffectsInfo represents the x-f5xc-side-effects extension
type SideEffectsInfo struct {
	Creates  []string `json:"creates,omitempty"`
	Modifies []string `json:"modifies,omitempty"`
	Deletes  []string `json:"deletes,omitempty"`
}

// BestPracticesInfo represents the x-f5xc-best-practices extension
type BestPracticesInfo struct {
	CommonErrors []CommonErrorInfo `json:"common_errors,omitempty"`
}

// CommonErrorInfo represents a common error and its resolution
type CommonErrorInfo struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	Resolution string `json:"resolution"`
	Prevention string `json:"prevention,omitempty"`
}

// GuidedWorkflowInfo represents the x-f5xc-guided-workflows extension
type GuidedWorkflowInfo struct {
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Steps       []*WorkflowStepInfo `json:"steps"`
}

// WorkflowStepInfo represents a step in a guided workflow
type WorkflowStepInfo struct {
	Order       int      `json:"order"`
	Action      string   `json:"action"`
	Description string   `json:"description,omitempty"`
	Fields      []string `json:"fields,omitempty"`
	Validation  string   `json:"validation,omitempty"`
}

// Global operations metadata collection (populated during generation)
var operationsMetadataCollection = &OperationsMetadataCollection{
	Resources: make(map[string]*ResourceOperationInfo),
}

// Global metadata collection (populated during generation)
var metadataCollection = &MetadataCollection{
	Resources: make(map[string]*ResourceMetadata),
}

// Global maps from index.json for resource metadata enrichment
var (
	resourceTierMap         = make(map[string]string)                        // resourceName -> tier
	resourceDependencyMap   = make(map[string]*openapi.ResourceDependencies) // resourceName -> dependencies
	resourceReferencedByMap = make(map[string][]string)                      // resourceName -> resources that depend on it
	resourceCategoryMap     = make(map[string]string)                        // resourceName -> category
)

var schemaCache = make(map[string]SchemaDefinition)
var rawSpecCache = make(map[string]map[string]interface{}) // Store raw JSON for x-ves-oneof-field extraction

func init() {
	flag.StringVar(&specDir, "spec-dir", "", "Directory containing OpenAPI spec files")
	flag.BoolVar(&dryRun, "dry-run", false, "Show what would be generated without writing files")
	flag.StringVar(&outputDir, "output-dir", "internal/provider", "Output directory for provider files")
	flag.StringVar(&clientDir, "client-dir", "internal/client", "Output directory for client files")
	flag.BoolVar(&verbose, "verbose", false, "Enable verbose output")
}

func main() {
	flag.Parse()

	// Check for spec directory
	if specDir == "" {
		specDir = os.Getenv("F5XC_SPEC_DIR")
	}
	if specDir == "" {
		specDir = "docs/specifications/api"
	}

	fmt.Println("🔨 F5XC Terraform Provider - Batch Schema Generator")
	fmt.Println(strings.Repeat("=", 60))
	fmt.Printf("📁 Spec Directory: %s\n", specDir)
	fmt.Printf("📁 Output Directory: %s\n", outputDir)
	if dryRun {
		fmt.Println("🔍 DRY RUN MODE - No files will be written")
	}
	fmt.Println()

	// Detect spec version (expects v2 format)
	specVersion := openapi.GetSpecVersion(specDir)
	fmt.Printf("🔍 Detected spec version: %s\n\n", specVersion)

	var results []GenerationResult
	var successCount, failCount int

	switch specVersion {
	case openapi.SpecVersionV2:
		results, successCount, failCount = processV2Specs(specDir)
	default:
		fmt.Printf("❌ Unknown spec format in directory: %s\n", specDir)
		fmt.Println("💡 Expected v2 spec format: index.json + domains/*.json structure")
		os.Exit(1)
	}

	// Generate combined client types file
	if !dryRun {
		generateCombinedClientTypes(results)
	}

	// Print summary
	fmt.Println()
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("📊 Generation Summary")
	fmt.Println(strings.Repeat("=", 60))
	fmt.Printf("✅ Successfully generated: %d resources\n", successCount)
	fmt.Printf("⏭️  Skipped (no schema): %d\n", len(results)-successCount-failCount)
	fmt.Printf("❌ Failed: %d\n", failCount)

	if failCount > 0 {
		fmt.Println("\n❌ Failed resources:")
		for _, r := range results {
			if !r.Success && r.Error != "" {
				fmt.Printf("   - %s: %s\n", r.ResourceName, r.Error)
			}
		}
	}

	// Generate provider registration
	if !dryRun {
		generateProviderRegistration(results)
	}

	// Write metadata files for MCP server
	if !dryRun {
		if err := writeMetadataFiles(); err != nil {
			fmt.Printf("⚠️  Warning: Failed to write metadata files: %v\n", err)
		}
	}

	fmt.Println("\n🎉 Batch generation complete!")
}

// processV2Specs processes v2 format specs (domain-organized files from f5xc-api-enriched)
func processV2Specs(specDir string) ([]GenerationResult, int, int) {
	// Parse the index.json to get domain information
	index, err := openapi.ParseIndexFromDir(specDir)
	if err != nil {
		fmt.Printf("❌ Error parsing index.json: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("📋 Spec version: %s\n", index.Version)
	fmt.Printf("📋 Generated at: %s\n", index.GeneratedAt)
	fmt.Printf("📄 Found %d domain specifications (v2 format)\n\n", len(index.Specifications))

	// Build global maps from index.json for metadata enrichment
	resourceTierMap = openapi.BuildResourceTierMap(index)
	resourceDependencyMap = openapi.BuildResourceDependencyMap(index)
	resourceReferencedByMap = openapi.BuildReferencedByMap(resourceDependencyMap)
	resourceCategoryMap = openapi.BuildResourceCategoryMap(index)
	if verbose {
		fmt.Printf("📊 Loaded metadata: %d resources with tier, %d with dependencies\n",
			len(resourceTierMap), len(resourceDependencyMap))
	}

	// Find all domain spec files
	domainFiles, err := openapi.FindDomainSpecFiles(specDir)
	if err != nil {
		fmt.Printf("❌ Error finding domain spec files: %v\n", err)
		os.Exit(1)
	}

	results := []GenerationResult{}
	successCount := 0
	failCount := 0
	skipCount := 0

	// Track processed resources to avoid duplicates across domain files
	// Some resources (like service_policy) appear in multiple domain specs
	processedResources := make(map[string]bool)

	// Build a map of domain metadata from index for quick lookup
	domainMetadata := make(map[string]openapi.DomainMetadata)
	for _, dm := range index.Specifications {
		domainMetadata[dm.Name] = dm
	}

	// Process each domain file
	for _, domainFile := range domainFiles {
		domainName := strings.TrimSuffix(filepath.Base(domainFile), ".json")
		fmt.Printf("🔄 Processing domain: %s\n", domainName)

		// Get domain metadata from index
		dm, hasMeta := domainMetadata[domainName]
		if hasMeta && verbose {
			fmt.Printf("   Category: %s, Tier: %s\n", dm.Category, dm.RequiresTier)
		}

		// Extract resources from the domain spec
		domainInfo, err := openapi.ExtractResourcesFromDomain(domainFile)
		if err != nil {
			fmt.Printf("   ⚠️  Error parsing domain: %v\n", err)
			results = append(results, GenerationResult{
				ResourceName: domainName,
				Success:      false,
				Error:        err.Error(),
			})
			failCount++
			continue
		}

		if len(domainInfo.Resources) == 0 {
			fmt.Printf("   ⏭️  No resources found in domain\n")
			continue
		}

		fmt.Printf("   📦 Found %d resources\n", len(domainInfo.Resources))

		// Process each resource in the domain
		for _, resource := range domainInfo.Resources {
			// Skip duplicate resources that appear in multiple domain specs
			if processedResources[resource.Name] {
				if verbose {
					fmt.Printf("      ⏭️  Skipping duplicate: %s (already processed)\n", resource.Name)
				}
				skipCount++
				continue
			}
			processedResources[resource.Name] = true

			// Create a virtual spec file path for compatibility with existing processing
			// The v2 domain spec contains all resources, so we process each individually
			result := processV2Resource(domainFile, resource, domainInfo)
			results = append(results, result)
			if result.Success {
				successCount++
			} else if result.Error != "" {
				failCount++
			}
		}
	}

	// Log duplicates if any were skipped
	if skipCount > 0 {
		fmt.Printf("\n⏭️  Skipped %d duplicate resources across domain files\n", skipCount)
	}

	return results, successCount, failCount
}

// processV2Resource processes a single resource from a v2 domain spec
func processV2Resource(domainFile string, resource openapi.ExtractedResource, domainInfo *openapi.DomainSpecInfo) GenerationResult {
	if verbose {
		fmt.Printf("      Processing resource: %s (category: %s, tier: %s)\n",
			resource.Name, resource.Category, resource.RequiresTier)
	}

	// For now, we'll use the same processSpecFile logic but with the domain file
	// This creates compatibility - the spec contains all schemas we need
	// We just need to focus on extracting the right schema for this resource

	// Use the existing processing with the domain file
	// The schema extraction will find the right Object schema based on resource name
	result := processSpecFileForResource(domainFile, resource.Name, resource.Category, resource.RequiresTier)

	return result
}

// processSpecFileForResource processes a spec file targeting a specific resource
// This is used for v2 specs where multiple resources exist in one domain file
func processSpecFileForResource(specFile string, resourceName string, category string, requiresTier string) GenerationResult {
	// Parse the spec file
	spec, err := parseOpenAPISpec(specFile)
	if err != nil {
		return GenerationResult{ResourceName: resourceName, Success: false, Error: err.Error()}
	}

	// Cache all schemas from the spec
	for name, schema := range spec.Components.Schemas {
		schemaCache[name] = schema
	}

	// Try to extract the resource schema
	schema, schemaName := extractResourceSchemaByName(spec, resourceName)
	if schema == nil {
		return GenerationResult{ResourceName: resourceName, Success: false, Error: ""}
	}

	// Extract API path from spec (or construct from resource name)
	apiPath := extractAPIPathForResource(spec, resourceName)
	if apiPath == "" {
		apiPath = fmt.Sprintf("/api/config/namespaces/{namespace}/%ss", resourceName)
	}

	// Generate the resource file using existing logic
	return generateResourceFromSchema(resourceName, schemaName, schema, apiPath, specFile, category, requiresTier)
}

// extractResourceSchemaByName extracts a specific resource schema by name from a spec
func extractResourceSchemaByName(spec *OpenAPI3Spec, resourceName string) (*SchemaDefinition, string) {
	// V2 schema naming patterns (e.g., app_firewallCreateSpecType)
	// The resource name is used as a prefix with CreateSpecType or GetSpecType suffix
	v2Patterns := []string{
		fmt.Sprintf("%sCreateSpecType", resourceName),
		fmt.Sprintf("%sGetSpecType", resourceName),
		fmt.Sprintf("%sReplaceSpecType", resourceName),
	}

	// Legacy schema naming patterns (ves.io.schema format still used in some v2 specs)
	legacyPatterns := []string{
		fmt.Sprintf("ves.io.schema.%s.Object", resourceName),
		fmt.Sprintf("%sType", toTitleCase(resourceName)),
		resourceName,
	}

	// Try v2 patterns first (CreateSpecType naming)
	for _, pattern := range v2Patterns {
		if schema, ok := spec.Components.Schemas[pattern]; ok {
			return &schema, pattern
		}
	}

	// Then try legacy patterns (ves.io.schema naming still present in some specs)
	for _, pattern := range legacyPatterns {
		if schema, ok := spec.Components.Schemas[pattern]; ok {
			return &schema, pattern
		}
	}

	// Try case-insensitive match for legacy .Object suffix
	lowerName := strings.ToLower(resourceName)
	for name, schema := range spec.Components.Schemas {
		if strings.Contains(strings.ToLower(name), lowerName) && strings.HasSuffix(name, ".Object") {
			return &schema, name
		}
	}

	// Try case-insensitive match for v2 CreateSpecType suffix
	for name, schema := range spec.Components.Schemas {
		if strings.Contains(strings.ToLower(name), lowerName) && strings.HasSuffix(name, "CreateSpecType") {
			return &schema, name
		}
	}

	return nil, ""
}

// extractAPIPathForResource extracts the API path for a specific resource from a spec
func extractAPIPathForResource(spec *OpenAPI3Spec, resourceName string) string {
	plural := resourceName + "s"
	// Handle special pluralization
	if strings.HasSuffix(resourceName, "y") {
		plural = strings.TrimSuffix(resourceName, "y") + "ies"
	}

	for path := range spec.Paths {
		if strings.Contains(path, "/"+plural) {
			return path
		}
	}
	return ""
}

// generateResourceFromSchema generates a resource file from an extracted schema
// This is factored out from processSpecFile to support both v1 and v2 processing
func generateResourceFromSchema(resourceName string, schemaName string, schema *SchemaDefinition, apiPath string, specFile string, category string, requiresTier string) GenerationResult {
	if verbose {
		fmt.Printf("      Generating: %s (schema: %s)\n", resourceName, schemaName)
		if category != "" {
			fmt.Printf("      Category: %s, Tier: %s\n", category, requiresTier)
		}
	}

	// Skip internal/utility schemas
	skipPatterns := []string{
		"object", "status", "spec", "metadata", "types", "common",
		"refs", "crudapi", "public", "private", "api", "empty",
	}
	for _, skip := range skipPatterns {
		if resourceName == skip {
			return GenerationResult{ResourceName: resourceName, Success: false}
		}
	}

	// Parse spec to get full schema information
	spec, err := parseOpenAPISpec(specFile)
	if err != nil {
		return GenerationResult{ResourceName: resourceName, Success: false, Error: err.Error()}
	}

	// Extract resource schema using the resource name we have
	resource, err := extractResourceSchema(spec, resourceName)
	if err != nil {
		if verbose {
			fmt.Printf("  ⏭️  Skipping %s: %v\n", resourceName, err)
		}
		return GenerationResult{ResourceName: resourceName, Success: false}
	}

	// Count attributes and blocks
	attrCount := 0
	blockCount := 0
	for _, attr := range resource.Attributes {
		if attr.IsBlock {
			blockCount++
		} else {
			attrCount++
		}
	}

	if !dryRun {
		// Generate resource file
		if err := generateResourceFile(resource); err != nil {
			return GenerationResult{ResourceName: resourceName, Success: false, Error: err.Error()}
		}

		// Generate client types
		if err := generateClientTypes(resource); err != nil {
			return GenerationResult{ResourceName: resourceName, Success: false, Error: err.Error()}
		}

		// Generate data source
		if err := generateDataSource(resource); err != nil {
			return GenerationResult{ResourceName: resourceName, Success: false, Error: err.Error()}
		}

		// Collect metadata for MCP server
		collectResourceMetadata(resource, category, requiresTier)

		// Collect operation-level metadata (v2.0.33 extensions)
		// Parse raw spec again to access path-level extensions
		if rawData, err := os.ReadFile(specFile); err == nil {
			var rawSpec map[string]interface{}
			if json.Unmarshal(rawData, &rawSpec) == nil {
				collectOperationMetadata(spec, resourceName, resource.APIPath, rawSpec)
			}
		}
	}

	fmt.Printf("✅ %s: %d attrs, %d blocks\n", resourceName, attrCount, blockCount)
	return GenerationResult{
		ResourceName: resourceName,
		Success:      true,
		AttrCount:    attrCount,
		BlockCount:   blockCount,
	}
}

func parseOpenAPISpec(specFile string) (*OpenAPI3Spec, error) {
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
		schemaCache[name] = schema
	}

	// Also parse raw JSON to extract x-ves-oneof-field annotations
	var rawSpec map[string]interface{}
	if err := json.Unmarshal(data, &rawSpec); err == nil {
		if components, ok := rawSpec["components"].(map[string]interface{}); ok {
			if schemas, ok := components["schemas"].(map[string]interface{}); ok {
				for name, schema := range schemas {
					if schemaMap, ok := schema.(map[string]interface{}); ok {
						rawSpecCache[name] = schemaMap
					}
				}
			}
		}
	}

	return &spec, nil
}

// extractAPIPath extracts the correct API path for CRUD operations from the OpenAPI spec
// It looks for paths containing POST (create) and returns the base path pattern
// Returns: basePath (for create/list), itemPath (for get/update/delete), hasNamespace (whether path has {namespace} segment)
func extractAPIPath(spec *OpenAPI3Spec, resourceName string) (basePath string, itemPath string, hasNamespace bool) {
	resourcePlural := resourceName + "s"

	// Look for CRUD paths in the spec
	for path, pathObj := range spec.Paths {
		pathMap, ok := pathObj.(map[string]interface{})
		if !ok {
			continue
		}

		// Check if this is a CRUD endpoint (has POST for create or GET for list)
		_, hasPost := pathMap["post"]
		_, hasGet := pathMap["get"]

		// Look for the base path (list/create endpoint) - ends with plural resource name
		// Pattern: /api/.../resource_names or /api/.../resource_names (with namespace)
		if (hasPost || hasGet) && strings.HasSuffix(path, "/"+resourcePlural) {
			// Check if path contains namespace segment
			hasNamespace = strings.Contains(path, "{namespace}") || strings.Contains(path, "{metadata.namespace}")

			// Convert {metadata.namespace} to %s for namespace substitution
			// Convert {metadata.name} or {name} for item paths
			basePath = path
			if hasNamespace {
				basePath = strings.ReplaceAll(basePath, "{metadata.namespace}", "%s")
				basePath = strings.ReplaceAll(basePath, "{namespace}", "%s")
			}

			// Item path is base path + /{name}
			itemPath = path + "/{name}"
			itemPath = strings.ReplaceAll(itemPath, "{metadata.namespace}", "%s")
			itemPath = strings.ReplaceAll(itemPath, "{namespace}", "%s")
			itemPath = strings.ReplaceAll(itemPath, "{metadata.name}", "%s")
			itemPath = strings.ReplaceAll(itemPath, "{name}", "%s")

			return basePath, itemPath, hasNamespace
		}
	}

	// Fallback to default pattern if no path found
	return fmt.Sprintf("/api/config/namespaces/%%s/%s", resourcePlural),
		fmt.Sprintf("/api/config/namespaces/%%s/%s/%%s", resourcePlural),
		true
}

func extractResourceSchema(spec *OpenAPI3Spec, resourceName string) (*ResourceTemplate, error) {
	// Find CreateSpecType schema
	var createSpec SchemaDefinition
	var found bool
	var createSpecKey string

	for key, schema := range spec.Components.Schemas {
		keyLower := strings.ToLower(key)
		if strings.Contains(keyLower, strings.ToLower(resourceName)) &&
			strings.Contains(keyLower, "createspectype") {
			createSpec = schema
			createSpecKey = key
			found = true
			break
		}
	}

	if !found {
		return nil, fmt.Errorf("no CreateSpecType found")
	}

	// Extract OneOf groups from x-ves-oneof-field annotations
	oneOfGroups := extractOneOfGroups(spec, createSpecKey)

	// Create reverse mapping: field -> group name + all fields in group
	// Also track which field should get the constraint (first alphabetically)
	fieldToOneOf := make(map[string][]string)
	fieldToGroupName := make(map[string]string) // Track the group name for AI-friendly defaults
	fieldIsFirst := make(map[string]bool)       // Only first field in each group gets the constraint
	for groupName, fields := range oneOfGroups {
		// Sort fields to determine which is first
		sortedFields := make([]string, len(fields))
		copy(sortedFields, fields)
		sort.Strings(sortedFields)
		firstField := sortedFields[0]

		for _, field := range fields {
			fieldToOneOf[field] = fields
			fieldToGroupName[field] = groupName
			if field == firstField {
				fieldIsFirst[field] = true
			}
		}
	}

	// Convert properties to Terraform attributes
	attributes := []TerraformAttribute{}
	requiredSet := make(map[string]bool)
	for _, r := range createSpec.Required {
		requiredSet[r] = true
	}

	for propName, propSchema := range createSpec.Properties {
		oneOfFields := fieldToOneOf[propName]
		groupName := fieldToGroupName[propName]
		attr := convertToTerraformAttribute(propName, propSchema, requiredSet[propName], "", spec)
		// Add OneOf constraint hint to description only for the first field in each group
		// Include group name for AI-friendly default recommendations
		if len(oneOfFields) > 1 && fieldIsFirst[propName] {
			attr.Description = addOneOfConstraintWithGroup(attr.Description, groupName, oneOfFields)
		}
		attributes = append(attributes, attr)
	}

	// Sort attributes per HashiCorp documentation standards:
	// Arguments: 1) ID components first, 2) Required alphabetically, 3) Optional alphabetically
	// Attributes: 1) id first, 2) remaining alphabetically
	sort.Slice(attributes, func(i, j int) bool {
		// Computed attributes go after arguments
		if attributes[i].Computed != attributes[j].Computed {
			return !attributes[i].Computed
		}
		// Required before optional
		if attributes[i].Required != attributes[j].Required {
			return attributes[i].Required
		}
		// Alphabetical within each group
		return attributes[i].Name < attributes[j].Name
	})

	// Extract correct API path from OpenAPI spec first to determine if namespace is required
	_, _, hasNamespace := extractAPIPath(spec, resourceName)

	// Add standard metadata attributes in HashiCorp-compliant order:
	// 1. ID components (name, namespace) - these form the resource ID
	// 2. Other required args alphabetically
	// 3. Optional args alphabetically (annotations, labels)
	// 4. Computed attributes (id first)

	// DNS zone and DNS domain resources use domain names (with dots), not standard names
	useDomainValidator := resourceName == "dns_zone" || resourceName == "dns_domain"
	nameDescription := fmt.Sprintf("Name of the %s. Must be unique within the namespace.", toHumanName(resourceName))
	if useDomainValidator {
		nameDescription = fmt.Sprintf("Domain name for the %s (e.g., example.com). Must be a valid DNS domain name.", toHumanName(resourceName))
	}

	idComponentAttrs := []TerraformAttribute{
		{Name: "name", GoName: "Name", TfsdkTag: "name", Type: "string",
			Description: nameDescription,
			Required:    true, PlanModifier: "RequiresReplace", UseDomainValidator: useDomainValidator},
	}

	// For resources without namespace in API path (like namespace itself), namespace is optional
	if hasNamespace {
		idComponentAttrs = append(idComponentAttrs, TerraformAttribute{
			Name: "namespace", GoName: "Namespace", TfsdkTag: "namespace", Type: "string",
			Description: fmt.Sprintf("Namespace where the %s will be created.", toHumanName(resourceName)),
			Required:    true, PlanModifier: "RequiresReplace"})
	} else {
		idComponentAttrs = append(idComponentAttrs, TerraformAttribute{
			Name: "namespace", GoName: "Namespace", TfsdkTag: "namespace", Type: "string",
			Description: fmt.Sprintf("Namespace for the %s. For this resource type, namespace should be empty or omitted.", toHumanName(resourceName)),
			Optional:    true, Computed: true, PlanModifier: "UseStateForUnknown"})
	}

	// Optional standard attrs will be sorted with other optionals
	// These match the F5XC schemaObjectCreateMetaType fields from OpenAPI specs
	optionalStdAttrs := []TerraformAttribute{
		{Name: "annotations", GoName: "Annotations", TfsdkTag: "annotations", Type: "map", ElementType: "string",
			Description: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.", Optional: true},
		{Name: "description", GoName: "Description", TfsdkTag: "description", Type: "string",
			Description: "Human readable description for the object.", Optional: true},
		{Name: "disable", GoName: "Disable", TfsdkTag: "disable", Type: "bool",
			Description: "A value of true will administratively disable the object.", Optional: true},
		{Name: "labels", GoName: "Labels", TfsdkTag: "labels", Type: "map", ElementType: "string",
			Description: "Labels is a user defined key value map that can be attached to resources for organization and filtering.", Optional: true},
	}

	// Computed attrs - id first per HashiCorp standards
	computedAttrs := []TerraformAttribute{
		{Name: "id", GoName: "ID", TfsdkTag: "id", Type: "string",
			Description: "Unique identifier for the resource.", Computed: true, PlanModifier: "UseStateForUnknown"},
	}

	// Combine: ID components first, then other required, then optional (incl. standard), then computed
	var sortedAttrs []TerraformAttribute
	sortedAttrs = append(sortedAttrs, idComponentAttrs...)

	// Add remaining required attributes (alphabetically)
	for _, attr := range attributes {
		if attr.Required && !attr.Computed {
			sortedAttrs = append(sortedAttrs, attr)
		}
	}

	// Add optional attributes (standard + schema-derived, alphabetically)
	// First, filter out standard attrs that already exist in schema-derived attrs to avoid duplicates
	schemaOptional := filterOptional(attributes)
	schemaAttrNames := make(map[string]bool)
	for _, attr := range schemaOptional {
		schemaAttrNames[attr.Name] = true
	}
	var filteredStdAttrs []TerraformAttribute
	for _, stdAttr := range optionalStdAttrs {
		if !schemaAttrNames[stdAttr.Name] {
			filteredStdAttrs = append(filteredStdAttrs, stdAttr)
		}
	}
	allOptional := append(filteredStdAttrs, schemaOptional...)
	sort.Slice(allOptional, func(i, j int) bool {
		return allOptional[i].Name < allOptional[j].Name
	})
	sortedAttrs = append(sortedAttrs, allOptional...)

	// Add computed attributes (id first, then others alphabetically)
	sortedAttrs = append(sortedAttrs, computedAttrs...)
	for _, attr := range attributes {
		if attr.Computed && attr.Name != "id" {
			sortedAttrs = append(sortedAttrs, attr)
		}
	}

	attributes = sortedAttrs

	// Get best description with enrichment extension priority:
	// 1. x-f5xc-description-medium (preferred - detailed but concise)
	// 2. x-f5xc-description-short (fallback - ultra-short)
	// 3. description (original)
	bestDescription := createSpec.XF5XCDescriptionMed
	if bestDescription == "" {
		bestDescription = createSpec.XF5XCDescriptionShort
	}
	if bestDescription == "" {
		bestDescription = createSpec.Description
	}
	description := transformResourceDescription(resourceName, bestDescription)

	// Generate example usage HCL
	exampleUsage := generateExampleUsage(resourceName, attributes)

	// Generate API docs URL
	apiDocsURL := fmt.Sprintf("https://docs.cloud.f5.com/docs/api/%s", strings.ReplaceAll(resourceName, "_", "-"))

	// Extract correct API path from OpenAPI spec
	apiPath, apiPathItem, hasNamespace := extractAPIPath(spec, resourceName)

	// Scan attributes to determine which plan modifier imports are needed
	usesBool, usesInt64, usesString := scanPlanModifierUsage(attributes)

	// Check if the resource has any nested models that would generate AttrTypes
	// AttrTypes (which use attr.Type) are generated for any block with nested attributes
	hasBlocks := hasNestedModelsWithAttrTypes(attributes)

	return &ResourceTemplate{
		Name:                   resourceName,
		TitleCase:              toTitleCase(resourceName),
		APIPath:                apiPath,
		APIPathPlural:          resourceName + "s",
		APIPathItem:            apiPathItem,
		HasNamespaceInPath:     hasNamespace,
		Description:            description,
		Attributes:             attributes,
		OneOfGroups:            oneOfGroups, // Now properly preserving extracted OneOf groups
		ExampleUsage:           exampleUsage,
		APIDocsURL:             apiDocsURL,
		UsesBoolPlanModifier:   usesBool,
		UsesInt64PlanModifier:  usesInt64,
		UsesStringPlanModifier: usesString,
		HasBlocks:              hasBlocks,
	}, nil
}

// hasNestedModelsWithAttrTypes checks recursively if any nested blocks would generate AttrTypes
// This is needed to determine if the attr import is required
// AttrTypes are generated for any nested model that has ANY nested attributes (block or non-block)
func hasNestedModelsWithAttrTypes(attributes []TerraformAttribute) bool {
	for _, attr := range attributes {
		if attr.IsBlock {
			// If this block has ANY nested attributes, AttrTypes will be generated for it
			if len(attr.NestedAttributes) > 0 {
				return true
			}
			// Note: Even if NestedAttributes is empty, we don't need to recurse
			// because empty blocks use EmptyModel which doesn't have AttrTypes
		}
	}
	return false
}

// scanPlanModifierUsage recursively scans attributes to determine which plan modifier imports are needed
func scanPlanModifierUsage(attributes []TerraformAttribute) (usesBool, usesInt64, usesString bool) {
	for _, attr := range attributes {
		if attr.PlanModifier != "" {
			switch attr.Type {
			case "bool":
				usesBool = true
			case "int64":
				usesInt64 = true
			case "string":
				usesString = true
			}
		}
		// Recursively scan nested attributes
		if len(attr.NestedAttributes) > 0 {
			nestedBool, nestedInt64, nestedString := scanPlanModifierUsage(attr.NestedAttributes)
			usesBool = usesBool || nestedBool
			usesInt64 = usesInt64 || nestedInt64
			usesString = usesString || nestedString
		}
	}
	return
}

// generateExampleUsage creates a sample HCL configuration for the resource
func generateExampleUsage(resourceName string, attributes []TerraformAttribute) string {
	// Get the correct namespace for this resource (fixes bug: was hardcoded to "system")
	_, ns := namespace.ForResource(resourceName)

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("resource \"f5xc_%s\" \"example\" {\n", resourceName))
	sb.WriteString("  name      = \"example\"\n")
	sb.WriteString(fmt.Sprintf("  namespace = %q\n", ns))
	sb.WriteString("\n")
	sb.WriteString("  labels = {\n")
	sb.WriteString("    \"env\" = \"production\"\n")
	sb.WriteString("  }\n")

	// Add example for optional scalar attributes (skip standard ones)
	standardAttrs := map[string]bool{"name": true, "namespace": true, "labels": true, "annotations": true, "id": true}
	for _, attr := range attributes {
		if standardAttrs[attr.Name] || attr.Computed || attr.IsBlock {
			continue
		}
		if attr.Optional {
			switch attr.Type {
			case "string":
				sb.WriteString(fmt.Sprintf("\n  %s = \"example-value\"\n", attr.Name))
			case "int64":
				sb.WriteString(fmt.Sprintf("\n  %s = 0\n", attr.Name))
			case "bool":
				sb.WriteString(fmt.Sprintf("\n  %s = false\n", attr.Name))
			}
			break // Just add one example optional attribute
		}
	}

	// Add example for first nested block with attributes
	for _, attr := range attributes {
		if attr.IsBlock && len(attr.NestedAttributes) > 0 {
			sb.WriteString(fmt.Sprintf("\n  %s {\n", attr.Name))
			for _, nested := range attr.NestedAttributes {
				if nested.IsBlock {
					continue // Skip deeply nested blocks in example
				}
				switch nested.Type {
				case "string":
					sb.WriteString(fmt.Sprintf("    %s = \"example\"\n", nested.Name))
				case "int64":
					sb.WriteString(fmt.Sprintf("    %s = 0\n", nested.Name))
				case "bool":
					sb.WriteString(fmt.Sprintf("    %s = false\n", nested.Name))
				}
				if len(sb.String()) > 500 { // Keep example concise
					break
				}
			}
			sb.WriteString("  }\n")
			break // Just add one example block
		}
	}

	sb.WriteString("}")
	return sb.String()
}

// Maximum recursion depth for nested schemas to prevent infinite loops
// Set high enough to capture all nested defaults while still preventing infinite recursion
const maxNestedDepth = 20

func convertToTerraformAttribute(name string, schema SchemaDefinition, required bool, oneOfGroup string, spec *OpenAPI3Spec) TerraformAttribute {
	return convertToTerraformAttributeWithDepth(name, schema, required, oneOfGroup, spec, 0, name)
}

func convertToTerraformAttributeWithDepth(name string, schema SchemaDefinition, required bool, oneOfGroup string, spec *OpenAPI3Spec, depth int, fieldPath string) TerraformAttribute {
	// Preserve extensions from original schema before resolving $ref
	// Extensions like x-f5xc-server-default are on the property, not the referenced schema
	serverDefault := schema.XF5XCServerDefault
	descShort := schema.XF5XCDescriptionShort
	descMed := schema.XF5XCDescriptionMed

	if schema.Ref != "" {
		schema = resolveRef(schema.Ref, spec)
		// Restore preserved extensions (property-level extensions take precedence)
		if serverDefault {
			schema.XF5XCServerDefault = serverDefault
		}
		if descShort != "" && schema.XF5XCDescriptionShort == "" {
			schema.XF5XCDescriptionShort = descShort
		}
		if descMed != "" && schema.XF5XCDescriptionMed == "" {
			schema.XF5XCDescriptionMed = descMed
		}
	}

	// Convert name to valid Terraform attribute name (lowercase with underscores)
	tfsdkName := toSnakeCase(name)

	// Handle reserved field names to avoid conflicts with standard metadata fields
	// "description" is used for resource metadata in TF schema, so rename spec fields
	goName := toTitleCase(name)
	if strings.ToLower(name) == "description" {
		goName = "DescriptionSpec"
		tfsdkName = "description_spec"
	}

	attr := TerraformAttribute{
		Name:        name,
		GoName:      goName,
		TfsdkTag:    tfsdkName,
		Required:    required,
		Optional:    !required,
		OneOfGroup:  oneOfGroup,
		MaxDepth:      depth,
		IsSpecField:   true, // Attributes from OpenAPI spec are spec fields
		JsonName:         name, // Original OpenAPI property name for JSON marshaling
		ServerDefault:    schema.XF5XCServerDefault,
		RecommendedValue: schema.XF5XCRecommendedValue,
	}

	// Build description with enrichment extension priority:
	// 1. x-f5xc-description-medium (preferred - detailed but concise)
	// 2. x-f5xc-description-short (fallback - ultra-short)
	// 3. x-displayname + description (original behavior)
	// 4. description only
	description := schema.XF5XCDescriptionMed
	if description == "" {
		description = schema.XF5XCDescriptionShort
	}
	if description == "" {
		if schema.XDisplayName != "" && schema.Description != "" {
			description = schema.XDisplayName + ". " + schema.Description
		} else if schema.XDisplayName != "" {
			description = schema.XDisplayName
		} else {
			description = schema.Description
		}
	}
	attr.Description = cleanDescription(description, fieldPath)
	if attr.Description == "" {
		attr.Description = fmt.Sprintf("Configuration for %s.", name)
	}

	// Format enum values per HashiCorp standards: "Possible values are `value1`, `value2`"
	if len(schema.Enum) > 0 {
		attr.Description = formatEnumDescription(attr.Description, schema.Enum)
	}

	// Format default values per HashiCorp standards: "Defaults to `value`."
	// First check for explicit default field, then try to extract from description text
	if schema.Default != nil {
		attr.Description = formatDefaultDescription(attr.Description, schema.Default)
	} else {
		// Try to extract default from description text (F5 XC often documents defaults this way)
		extractedDefault, cleanedDesc := extractDefaultFromDescription(attr.Description)
		if extractedDefault != nil {
			attr.Description = formatDefaultDescription(cleanedDesc, extractedDefault)
		}
	}

	// Add server default note to description (x-f5xc-server-default extension)
	// This indicates fields where F5XC server applies sensible defaults when omitted
	if schema.XF5XCServerDefault {
		attr.Description += " Server applies default when omitted."
	}

	// Add recommended value note to description (x-f5xc-recommended-value extension)
	// This indicates suggested values for fields that benefit from guidance
	if schema.XF5XCRecommendedValue != nil {
		attr.Description += fmt.Sprintf(" Recommended value: `%v`.", schema.XF5XCRecommendedValue)
	}

	// Determine type and extract nested attributes
	switch schema.Type {
	case "string":
		attr.Type = "string"
		attr.GoType = "string"
	case "integer", "number":
		attr.Type = "int64"
		attr.GoType = "int64"
	case "boolean":
		attr.Type = "bool"
		attr.GoType = "bool"
	case "array":
		attr.Type = "list"
		attr.GoType = "[]interface{}" // Default, refined below
		if schema.Items != nil {
			itemSchema := *schema.Items
			if itemSchema.Ref != "" {
				itemSchema = resolveRef(itemSchema.Ref, spec)
			}
			if itemSchema.Type == "object" || len(itemSchema.Properties) > 0 {
				attr.IsBlock = true
				attr.NestedBlockType = "list"
				attr.GoType = "[]map[string]interface{}"
				// Extract nested attributes if within depth limit
				if depth < maxNestedDepth {
					attr.NestedAttributes = extractNestedAttributes(itemSchema, spec, depth+1, fieldPath)
				}
			} else {
				attr.ElementType = mapSchemaType(itemSchema.Type)
				attr.GoType = "[]" + mapSchemaTypeToGo(itemSchema.Type)
				// Capture enum and default values from item schema for list element documentation
				if len(itemSchema.Enum) > 0 {
					attr.Description = formatEnumDescription(attr.Description, itemSchema.Enum)
				}
				if itemSchema.Default != nil {
					attr.Description = formatDefaultDescription(attr.Description, itemSchema.Default)
				}
			}
		}
	case "object":
		if len(schema.Properties) > 0 {
			attr.Type = "object"
			attr.IsBlock = true
			attr.NestedBlockType = "single"
			attr.GoType = "map[string]interface{}"
			// Extract nested attributes if within depth limit
			if depth < maxNestedDepth {
				attr.NestedAttributes = extractNestedAttributes(schema, spec, depth+1, fieldPath)
			}
		} else if schema.AdditionalProperties != nil {
			attr.Type = "map"
			attr.ElementType = "string"
			attr.GoType = "map[string]string"
		} else {
			attr.Type = "object"
			attr.IsBlock = true
			attr.NestedBlockType = "single"
			attr.GoType = "map[string]interface{}"
		}
	default:
		if len(schema.Properties) > 0 {
			attr.Type = "object"
			attr.IsBlock = true
			attr.NestedBlockType = "single"
			attr.GoType = "map[string]interface{}"
			// Extract nested attributes if within depth limit
			if depth < maxNestedDepth {
				attr.NestedAttributes = extractNestedAttributes(schema, spec, depth+1, fieldPath)
			}
		} else {
			attr.Type = "string"
			attr.GoType = "string"
		}
	}

	// For optional scalar attributes (string, bool, int64), mark as computed with
	// UseStateForUnknown to prevent drift from API defaults. The API may return
	// default values for optional fields that weren't set in the config, causing
	// spurious plan diffs. UseStateForUnknown tells Terraform to preserve the
	// prior state value when the config doesn't specify a value.
	// Note: We only apply this to top-level attributes (depth == 0) because:
	// 1. Only top-level attributes are rendered with plan modifiers in the template
	// 2. Nested attributes inside blocks don't need this protection as the block
	//    structure itself handles the state management
	if depth == 0 && attr.Optional && !attr.IsBlock && !attr.Required {
		if attr.Type == "string" || attr.Type == "bool" || attr.Type == "int64" {
			attr.Computed = true
			attr.PlanModifier = "UseStateForUnknown"
		}
	}

	return attr
}

// extractNestedAttributes extracts attributes from an object schema's properties
func extractNestedAttributes(schema SchemaDefinition, spec *OpenAPI3Spec, depth int, parentPath string) []TerraformAttribute {
	if depth > maxNestedDepth {
		return nil
	}

	requiredSet := make(map[string]bool)
	for _, r := range schema.Required {
		requiredSet[r] = true
	}

	var attrs []TerraformAttribute
	for propName, propSchema := range schema.Properties {
		nestedPath := propName
		if parentPath != "" {
			nestedPath = parentPath + "." + propName
		}
		attr := convertToTerraformAttributeWithDepth(propName, propSchema, requiredSet[propName], "", spec, depth, nestedPath)

		// Mark 'namespace', 'tenant', 'uid', and 'kind' fields as Computed in nested Object Reference blocks.
		// The API always returns these values even when not specified in config,
		// which causes state drift if not marked as Computed.
		// Object Reference types have pattern: kind, name, namespace, tenant, uid
		// Using UseStateForUnknown plan modifier prevents perpetual drift.
		propNameLower := strings.ToLower(propName)
		if (propNameLower == "namespace" || propNameLower == "tenant" || propNameLower == "uid" || propNameLower == "kind") && !attr.Required {
			attr.Computed = true
			attr.Optional = true
			attr.PlanModifier = "UseStateForUnknown"
		}

		attrs = append(attrs, attr)
	}

	// Sort attributes: required first, then alphabetically
	sort.Slice(attrs, func(i, j int) bool {
		if attrs[i].Required != attrs[j].Required {
			return attrs[i].Required
		}
		return attrs[i].Name < attrs[j].Name
	})

	return attrs
}

func resolveRef(ref string, spec *OpenAPI3Spec) SchemaDefinition {
	parts := strings.Split(ref, "/")
	schemaName := parts[len(parts)-1]

	if schema, ok := spec.Components.Schemas[schemaName]; ok {
		return schema
	}
	if schema, ok := schemaCache[schemaName]; ok {
		return schema
	}
	return SchemaDefinition{Type: "string"}
}

func mapSchemaType(t string) string {
	switch t {
	case "string":
		return "string"
	case "integer", "number":
		return "int64"
	case "boolean":
		return "bool"
	default:
		return "string"
	}
}

// mapSchemaTypeToGo converts OpenAPI types to Go types for client struct generation
func mapSchemaTypeToGo(t string) string {
	switch t {
	case "string":
		return "string"
	case "integer", "number":
		return "int64"
	case "boolean":
		return "bool"
	default:
		return "interface{}"
	}
}

func cleanDescription(desc string, fieldPath string) string {
	// Remove example and validation rules sections
	desc = regexp.MustCompile(`\s*Example:.*`).ReplaceAllString(desc, "")
	desc = regexp.MustCompile(`\s*Validation Rules:.*`).ReplaceAllString(desc, "")

	// Remove x-example annotations (OpenAPI 2.0 vendor extension for Swagger UI examples)
	// Pattern: x-example: "value" or x-example: 'value' embedded in description text
	desc = regexp.MustCompile(`\s*x-example:\s*["']?[^"'\n]*["']?`).ReplaceAllString(desc, "")
	// Also handle x-required annotations
	desc = regexp.MustCompile(`\s*x-required\s*`).ReplaceAllString(desc, "")

	// Remove ves.io validation annotations (common pattern in F5 XC specs)
	// These are internal protobuf validation rules that leaked into OpenAPI descriptions
	// Pattern: ves.io.schema.rules.xxx.yyy: value or ves.io.schema.xxx: value
	desc = regexp.MustCompile(`\s*ves\.io\.schema[^\s]*:\s*\S+`).ReplaceAllString(desc, "")
	desc = regexp.MustCompile(`\s*ves\.io\.[^\s]*:\s*\[.*?\]`).ReplaceAllString(desc, "")

	// Remove "Required: YES" or "Required: NO" annotations
	desc = regexp.MustCompile(`\s*Required:\s*(YES|NO)\s*`).ReplaceAllString(desc, " ")
	// Remove "Exclusive with [xxx]" patterns
	desc = regexp.MustCompile(`\s*Exclusive with\s*\[[^\]]*\]\s*`).ReplaceAllString(desc, " ")

	// Normalize generic empty message descriptions to user-friendly text
	// "Empty. This can be used for messages where no values are needed" → "Enable this option"
	desc = regexp.MustCompile(`(?i)Empty\.?\s*This can be used for messages where no values are needed\.?`).ReplaceAllString(desc, "Enable this option")
	// Also handle variations
	desc = regexp.MustCompile(`(?i)This can be used for messages where no values are needed\.?`).ReplaceAllString(desc, "Enable this option")

	// Normalize "Shape of the X specification" to "Configuration for X"
	// This converts internal F5 terminology to user-friendly Terraform terminology
	desc = regexp.MustCompile(`(?i)Shape of the ([^\s]+) specification`).ReplaceAllString(desc, "Configuration for $1")
	desc = regexp.MustCompile(`(?i)Shape of ([^\s]+) specification`).ReplaceAllString(desc, "Configuration for $1")

	// Remove escaped quotes and backslashes from raw spec data
	desc = strings.ReplaceAll(desc, `\"`, `"`)
	desc = strings.ReplaceAll(desc, `\\`, `\`)
	// Normalize whitespace
	desc = regexp.MustCompile(`[\n\r]+`).ReplaceAllString(desc, " ")
	desc = regexp.MustCompile(`\s+`).ReplaceAllString(desc, " ")
	// Escape quotes for Go string literals
	desc = strings.ReplaceAll(desc, `"`, "'")
	desc = strings.TrimSpace(desc)
	// Remove trailing periods that were left after cleanup
	desc = regexp.MustCompile(`\.\s*\.`).ReplaceAllString(desc, ".")
	// Normalize example names from F5 internal conventions ("my-*") to provider standard ("example-*")
	desc = naming.NormalizeExampleNames(desc)
	return desc
}

// filterOptional returns only optional (non-required, non-computed) attributes
func filterOptional(attrs []TerraformAttribute) []TerraformAttribute {
	var result []TerraformAttribute
	for _, attr := range attrs {
		if attr.Optional && !attr.Required && !attr.Computed {
			result = append(result, attr)
		}
	}
	return result
}

// extractDefaultFromDescription attempts to extract a default value from description text.
// This handles cases where F5 XC OpenAPI specs mention defaults in descriptions but don't use the default field.
// Returns the extracted default value (or nil) and the cleaned description with default text removed.
func extractDefaultFromDescription(desc string) (interface{}, string) {
	if desc == "" {
		return nil, desc
	}

	// Patterns to match defaults mentioned in description text
	// Pattern 1: "Defaults to X" or "Default to X" with optional units (including "or Ys" alternative)
	// Pattern 2: "Default value is X" or "Default value: X"
	// Pattern 3: "defaults to X" (lowercase)
	patterns := []string{
		// Match "Defaults to 30000ms or 30s" or "Defaults to 30000ms" or "Defaults to true"
		// Captures first value, removes optional " or Xs" alternative
		`[Dd]efaults?\s+to\s+(\d+(?:ms|s|%)?|\d+\.\d+|true|false)(?:\s+or\s+\d+(?:ms|s)?)?`,
		// Match "Default value is /graphql" or "Default value is true"
		`[Dd]efault\s+value\s+(?:is|:)\s+([^\s.,]+)`,
		// Match "default is 10" or "default: 10"
		`[Dd]efault\s+(?:is|:)\s+(\d+(?:ms|s|%)?|\d+\.\d+|true|false)`,
	}

	for _, pattern := range patterns {
		re := regexp.MustCompile(pattern)
		match := re.FindStringSubmatch(desc)
		if len(match) >= 2 {
			defaultVal := strings.TrimSpace(match[1])
			// Skip if it looks like a placeholder or invalid value
			upperVal := strings.ToUpper(defaultVal)
			if upperVal == "NONE" || upperVal == "INVALID" || upperVal == "UNKNOWN" || upperVal == "UNSPECIFIED" {
				continue
			}
			// Remove the matched text from description to avoid duplication
			cleanedDesc := re.ReplaceAllString(desc, "")
			cleanedDesc = strings.TrimSpace(cleanedDesc)
			// Clean up any trailing punctuation/whitespace artifacts
			cleanedDesc = regexp.MustCompile(`\s+\.`).ReplaceAllString(cleanedDesc, ".")
			cleanedDesc = regexp.MustCompile(`\.\s*\.`).ReplaceAllString(cleanedDesc, ".")
			return defaultVal, cleanedDesc
		}
	}

	return nil, desc
}

// formatDefaultDescription appends a default value to a description per HashiCorp standards.
// Format: "Defaults to `value`."
func formatDefaultDescription(desc string, defaultValue interface{}) string {
	if defaultValue == nil {
		return desc
	}

	// Convert default value to string
	defaultStr := fmt.Sprintf("%v", defaultValue)
	if defaultStr == "" || defaultStr == "<nil>" {
		return desc
	}

	// Skip invalid/placeholder defaults using EXACT match only
	// These are sentinel values that indicate "no value" rather than actual defaults
	invalidDefaults := map[string]bool{
		"INVALID":     true,
		"NONE":        true,
		"UNKNOWN":     true,
		"UNSPECIFIED": true,
		"0":           true, // Integer zero is often just the zero value
	}

	// Check if default is EXACTLY an invalid/placeholder value
	// Use exact matching to preserve valid defaults like XFCC_NONE, MTLS_NONE, etc.
	upperDefault := strings.ToUpper(defaultStr)
	if invalidDefaults[upperDefault] {
		return desc
	}

	// Ensure description ends properly before adding default info
	desc = strings.TrimSpace(desc)
	if desc != "" && !strings.HasSuffix(desc, ".") && !strings.HasSuffix(desc, ":") {
		desc += "."
	}

	return fmt.Sprintf("%s Defaults to `%s`.", desc, defaultStr)
}

// formatEnumDescription adds AI-parseable enum metadata and human-readable values to description.
// Format: "[Enum: val1|val2|val3] Human description. Possible values are `val1`, `val2`, `val3`."
// The [Enum: ...] prefix enables AI tools to deterministically extract valid values from
// `terraform providers schema -json` output without parsing natural language.
func formatEnumDescription(desc string, enumValues []interface{}) string {
	if len(enumValues) == 0 {
		return desc
	}

	// Convert enum values to strings (raw values for AI prefix, backtick for human)
	var rawValues []string
	var formattedValues []string
	for _, v := range enumValues {
		str := fmt.Sprintf("%v", v)
		// Skip empty or very long values
		if str == "" || len(str) > 50 {
			continue
		}
		rawValues = append(rawValues, str)
		formattedValues = append(formattedValues, fmt.Sprintf("`%s`", str))
	}

	if len(rawValues) == 0 {
		return desc
	}

	// Build AI-parseable prefix: [Enum: val1|val2|val3]
	aiPrefix := fmt.Sprintf("[Enum: %s]", strings.Join(rawValues, "|"))

	// Ensure description ends properly before adding enum info
	desc = strings.TrimSpace(desc)
	if desc != "" && !strings.HasSuffix(desc, ".") && !strings.HasSuffix(desc, ":") {
		desc += "."
	}

	// Format based on number of values per HashiCorp standards
	var humanSuffix string
	if len(formattedValues) == 1 {
		humanSuffix = fmt.Sprintf("The only possible value is %s.", formattedValues[0])
	} else {
		humanSuffix = fmt.Sprintf("Possible values are %s.", strings.Join(formattedValues, ", "))
	}

	// Combine: [AI prefix] Human description. Human enum list.
	if desc == "" {
		return fmt.Sprintf("%s %s", aiPrefix, humanSuffix)
	}
	return fmt.Sprintf("%s %s %s", aiPrefix, desc, humanSuffix)
}

// resourceCategories maps resource names to their functional categories for AI discovery
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

	// Networking
	"network_connector": "Networking",
	"virtual_network":   "Networking",
	"cloud_connect":     "Networking",
	"cloud_link":        "Networking",
	"bgp":               "Networking",
	"ip_prefix_set":     "Networking",
	"network_interface": "Networking",
	"virtual_site":      "Networking",

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
	"api_testing":    "API Security",
	"api_crawler":    "API Security",

	// Organization
	"namespace":      "Organization",
	"tenant":         "Organization",
	"role":           "Organization",
	"allowed_tenant": "Organization",
}

// resourceDependencies maps resource names to their common dependencies
// This helps AI tools understand creation order
var resourceDependencies = map[string][]string{
	"http_loadbalancer": {"namespace", "origin_pool"},
	"tcp_loadbalancer":  {"namespace", "origin_pool"},
	"udp_loadbalancer":  {"namespace", "origin_pool"},
	"origin_pool":       {"namespace", "healthcheck"},
	"healthcheck":       {"namespace"},
	"route":             {"namespace", "http_loadbalancer"},
	"app_firewall":      {"namespace"},
	"service_policy":    {"namespace"},
	"rate_limiter":      {"namespace"},
	"certificate":       {"namespace"},
	"api_definition":    {"namespace"},
	"dns_zone":          {},
	"virtual_site":      {},
	"namespace":         {},
}

// getResourceAIMetadata generates AI-parseable metadata prefix for a resource
// Format: [Category: X] [Namespace: required|optional] [DependsOn: res1, res2]
func getResourceAIMetadata(resourceName string) string {
	var parts []string

	// Add category if known
	if category, ok := resourceCategories[resourceName]; ok {
		parts = append(parts, fmt.Sprintf("[Category: %s]", category))
	}

	// Add namespace requirement
	// Most F5 XC resources require a namespace except for system-level resources
	systemResources := map[string]bool{
		"namespace": true, "tenant": true, "role": true, "allowed_tenant": true,
		"dns_zone": true, "dns_domain": true, "virtual_site": true,
		"cloud_credentials": true, "certificate": true, "trusted_ca_list": true,
	}
	if systemResources[resourceName] {
		parts = append(parts, "[Namespace: not_required]")
	} else {
		parts = append(parts, "[Namespace: required]")
	}

	// Add dependencies if known
	if deps, ok := resourceDependencies[resourceName]; ok && len(deps) > 0 {
		parts = append(parts, fmt.Sprintf("[DependsOn: %s]", strings.Join(deps, ", ")))
	}

	if len(parts) == 0 {
		return ""
	}
	return strings.Join(parts, " ")
}

// transformResourceDescription converts technical API descriptions into user-friendly
// Terraform resource descriptions following HashiCorp best practices.
// Pattern: "Manages a [Resource] in F5 Distributed Cloud [for purpose/capability]."
func transformResourceDescription(resourceName, rawDescription string) string {
	humanName := toHumanName(resourceName)

	// Clean and normalize the raw description first
	// Pass empty fieldPath since this is the resource-level description
	desc := cleanDescription(rawDescription, "")
	desc = strings.TrimSpace(desc)

	// Generate the human-readable description
	var humanDesc string

	// If empty, use default
	if desc == "" {
		humanDesc = fmt.Sprintf("Manages a %s resource in F5 Distributed Cloud.", humanName)
	} else {
		// Detect and transform common technical description patterns
		lowerDesc := strings.ToLower(desc)

		// Pattern 1: "Shape of the X specification" -> extract X and make user-friendly
		if strings.Contains(lowerDesc, "shape of") {
			humanDesc = generateCapabilityDescriptionOnly(resourceName, humanName, desc)
		} else if strings.HasSuffix(lowerDesc, " object") || strings.HasSuffix(lowerDesc, " configuration") ||
			strings.HasSuffix(lowerDesc, " spec") || strings.HasSuffix(lowerDesc, " specification") {
			// Pattern 2: "X object" or "X configuration" - technical object reference
			humanDesc = generateCapabilityDescriptionOnly(resourceName, humanName, desc)
		} else {
			// Pattern 3: Already starts with a verb like "Create", "Configure", "Define"
			// Transform to "Manages" for consistency
			actionVerbs := []string{"create", "configure", "define", "set up", "establish", "provision"}
			matched := false
			for _, verb := range actionVerbs {
				if strings.HasPrefix(lowerDesc, verb) {
					// Replace the action verb with "Manages" for Terraform consistency
					remainder := desc[len(verb):]
					remainder = strings.TrimPrefix(remainder, "s") // handle "Creates" -> "Create"
					remainder = strings.TrimSpace(remainder)
					if remainder != "" {
						// Clean up articles
						remainder = strings.TrimPrefix(remainder, "a ")
						remainder = strings.TrimPrefix(remainder, "an ")
						remainder = strings.TrimPrefix(remainder, "the ")
						humanDesc = fmt.Sprintf("Manages %s in F5 Distributed Cloud.", remainder)
						matched = true
						break
					}
				}
			}

			if !matched {
				// Pattern 4: Description is already decent but needs "Manages" prefix
				// If it doesn't start with a verb, add "Manages a X resource" prefix
				if !startsWithVerb(desc) {
					// Use the description as the capability explanation
					capability := extractCapabilityFromDescription(desc)
					if capability != "" {
						humanDesc = fmt.Sprintf("Manages a %s resource in F5 Distributed Cloud for %s.", humanName, capability)
					} else {
						humanDesc = fmt.Sprintf("Manages a %s resource in F5 Distributed Cloud. %s", humanName, desc)
					}
				} else {
					// If description already looks good, just ensure it ends properly
					if !strings.HasSuffix(desc, ".") {
						desc = desc + "."
					}
					humanDesc = desc
				}
			}
		}
	}

	// Return human-readable description only
	// Note: AI metadata was previously added here for the terraform-schema-ai tool,
	// but that tool was removed in commit 0c45bb3 as unused. The metadata was
	// polluting Terraform Registry documentation for human users.
	return humanDesc
}

// generateCapabilityDescriptionOnly is like generateCapabilityDescription but returns only the description
// without AI metadata (metadata is added by the caller)
func generateCapabilityDescriptionOnly(resourceName, humanName, rawDesc string) string {
	// Resource-specific capability mappings for common F5 XC resources
	capabilities := map[string]string{
		// Sites
		"securemesh_site":    "deploying secure mesh edge sites with distributed security capabilities",
		"securemesh_site_v2": "deploying secure mesh edge sites with enhanced security and networking features",
		"aws_vpc_site":       "deploying F5 sites within AWS VPC environments",
		"azure_vnet_site":    "deploying F5 sites within Azure Virtual Network environments",
		"gcp_vpc_site":       "deploying F5 sites within Google Cloud VPC environments",
		"aws_tgw_site":       "deploying F5 sites connected via AWS Transit Gateway",
		"voltstack_site":     "deploying Volterra stack sites for edge computing",
		"virtual_site":       "creating logical groupings of sites based on labels and selectors",

		// Load Balancing
		"http_loadbalancer": "load balancing HTTP/HTTPS traffic with advanced routing and security",
		"tcp_loadbalancer":  "load balancing TCP traffic across origin pools",
		"udp_loadbalancer":  "load balancing UDP traffic across origin pools",
		"dns_load_balancer": "intelligent DNS-based load balancing across multiple endpoints",
		"cdn_loadbalancer":  "content delivery and edge caching with load balancing",
		"origin_pool":       "defining backend server pools for load balancer targets",
		"healthcheck":       "monitoring backend server health and availability",
		"route":             "defining traffic routing rules for load balancers",

		// Security
		"app_firewall":                   "web application firewall (WAF) protection",
		"service_policy":                 "defining service-level access control and security policies",
		"network_firewall":               "network-level firewall rules and security controls",
		"rate_limiter":                   "protecting services from traffic spikes and DDoS attacks",
		"bot_defense_app_infrastructure": "bot detection and mitigation capabilities",
		"malicious_user_mitigation":      "identifying and blocking malicious user behavior",
		"waf_exclusion_policy":           "excluding specific requests from WAF inspection",

		// Networking
		"network_connector": "connecting networks across sites and cloud providers",
		"virtual_network":   "creating isolated virtual network segments",
		"cloud_connect":     "establishing connectivity to cloud provider networks",
		"cloud_link":        "linking F5 sites to cloud provider infrastructure",
		"bgp":               "BGP routing configuration for network connectivity",
		"ip_prefix_set":     "defining IP address prefix lists for network policies",
		"network_interface": "configuring network interfaces on sites",

		// DNS
		"dns_zone":              "DNS zone management and configuration",
		"dns_domain":            "DNS domain registration and management",
		"dns_lb_pool":           "DNS load balancer endpoint pools",
		"dns_lb_health_check":   "health monitoring for DNS load balanced endpoints",
		"dns_compliance_checks": "DNS security and compliance verification",

		// Kubernetes
		"k8s_cluster":              "Kubernetes cluster integration and management",
		"virtual_k8s":              "virtual Kubernetes cluster deployment",
		"k8s_cluster_role":         "Kubernetes RBAC cluster role definitions",
		"k8s_cluster_role_binding": "Kubernetes RBAC cluster role bindings",
		"k8s_pod_security_policy":  "Kubernetes pod security policy enforcement",
		"container_registry":       "container image registry configuration",

		// Authentication & Secrets
		"authentication":    "authentication methods and identity provider integration",
		"cloud_credentials": "cloud provider credential management for site deployment",
		"api_credential":    "API credential management for service authentication",
		"token":             "API token generation and management",
		"secret_policy":     "secret access policies and controls",

		// Certificates
		"certificate":       "TLS/SSL certificate management",
		"certificate_chain": "certificate chain configuration for TLS",
		"trusted_ca_list":   "trusted certificate authority list management",

		// Monitoring & Logging
		"log_receiver":        "log collection and forwarding configuration",
		"global_log_receiver": "global log aggregation settings",
		"alert_policy":        "alerting rules and notification policies",
		"alert_receiver":      "alert notification endpoints",

		// API Security
		"api_definition": "API schema and endpoint definitions for security",
		"api_discovery":  "automatic API endpoint discovery and inventory",
		"api_testing":    "API testing and validation capabilities",
		"api_crawler":    "API endpoint crawling and discovery",

		// Organization
		"namespace":      "logical namespace isolation for resources",
		"tenant":         "tenant configuration and management",
		"role":           "role-based access control definitions",
		"allowed_tenant": "tenant access permissions and restrictions",
	}

	// Check if we have a specific capability mapping
	if capability, ok := capabilities[resourceName]; ok {
		return fmt.Sprintf("Manages a %s resource in F5 Distributed Cloud for %s.", humanName, capability)
	}

	// Try to extract capability from raw description
	capability := extractCapabilityFromDescription(rawDesc)
	if capability != "" {
		return fmt.Sprintf("Manages a %s resource in F5 Distributed Cloud for %s.", humanName, capability)
	}

	// Default fallback
	return fmt.Sprintf("Manages a %s resource in F5 Distributed Cloud.", humanName)
}

// extractCapabilityFromDescription tries to extract a meaningful capability phrase
// from technical descriptions
func extractCapabilityFromDescription(desc string) string {
	lowerDesc := strings.ToLower(desc)

	// Remove common technical prefixes
	prefixes := []string{
		"shape of the ",
		"shape of ",
		"specification for ",
		"configuration for ",
		"defines the ",
		"defines a ",
		"represents the ",
		"represents a ",
		"the ",
		"a ",
		"an ",
	}
	for _, prefix := range prefixes {
		if strings.HasPrefix(lowerDesc, prefix) {
			desc = desc[len(prefix):]
			lowerDesc = strings.ToLower(desc)
			break
		}
	}

	// Remove common technical suffixes
	suffixes := []string{
		" specification",
		" spec",
		" configuration",
		" config",
		" object",
		" definition",
	}
	for _, suffix := range suffixes {
		if strings.HasSuffix(lowerDesc, suffix) {
			desc = desc[:len(desc)-len(suffix)]
			break
		}
	}

	desc = strings.TrimSpace(desc)

	// If what remains is meaningful (more than just a name), use it
	if len(desc) > 5 && !strings.Contains(strings.ToLower(desc), "shape") {
		// Convert to lowercase capability phrase
		return strings.ToLower(desc) + " configuration"
	}

	return ""
}

// startsWithVerb checks if the description starts with an action verb
func startsWithVerb(desc string) bool {
	verbs := []string{
		"manages", "creates", "configures", "defines", "sets", "establishes",
		"provisions", "deploys", "enables", "allows", "provides", "supports",
		"manage", "create", "configure", "define", "set", "establish",
		"provision", "deploy", "enable", "allow", "provide", "support",
	}
	lowerDesc := strings.ToLower(desc)
	for _, verb := range verbs {
		if strings.HasPrefix(lowerDesc, verb+" ") || strings.HasPrefix(lowerDesc, verb+"s ") {
			return true
		}
	}
	return false
}

// extractOneOfGroups extracts x-ves-oneof-field annotations from the raw schema JSON
func extractOneOfGroups(spec *OpenAPI3Spec, schemaKey string) map[string][]string {
	oneOfGroups := make(map[string][]string)

	// Get raw schema from cache
	rawSchema, ok := rawSpecCache[schemaKey]
	if !ok {
		return oneOfGroups
	}

	// Look for x-ves-oneof-field-* in the raw schema
	for key, value := range rawSchema {
		if strings.HasPrefix(key, "x-ves-oneof-field-") {
			groupName := strings.TrimPrefix(key, "x-ves-oneof-field-")
			// Value can be either a JSON array string or actual array
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

// =============================================================================
// METADATA COLLECTION FUNCTIONS (for MCP Server)
// =============================================================================

// collectResourceMetadata extracts and stores metadata for a resource
func collectResourceMetadata(resource *ResourceTemplate, category string, requiresTier string) {
	if resource == nil {
		return
	}

	// Determine tier: prefer resource-level from index.json, fall back to domain-level
	tier := requiresTier
	if resourceTier, ok := resourceTierMap[resource.Name]; ok && resourceTier != "" {
		tier = resourceTier
	}

	// Determine category: prefer resource-level from index.json, fall back to domain-level
	cat := category
	if resourceCat, ok := resourceCategoryMap[resource.Name]; ok && resourceCat != "" {
		cat = resourceCat
	}

	// Determine import format based on HasNamespaceInPath
	importFormat := "name"
	if resource.HasNamespaceInPath {
		importFormat = "namespace/name"
	}

	metadata := &ResourceMetadata{
		Description:  resource.Description,
		Category:     cat,
		Tier:         tier,
		ImportFormat: importFormat,
		OneOfGroups:  make(map[string]*OneOfGroupInfo),
		Attributes:   make(map[string]*AttributeMetadata),
	}

	// Convert OneOfGroups from ResourceTemplate format to metadata format
	for groupName, fields := range resource.OneOfGroups {
		defaultField := determineOneOfDefault(groupName, fields)
		metadata.OneOfGroups[groupName] = &OneOfGroupInfo{
			Fields:  fields,
			Default: defaultField,
		}
	}

	// Extract attribute metadata
	extractAttributeMetadata(resource.Attributes, metadata.Attributes, resource.OneOfGroups)

	// Add dependency information from index.json
	if deps, ok := resourceDependencyMap[resource.Name]; ok && deps != nil {
		metadata.Dependencies = &DependencyInfo{
			References: append(deps.Required, deps.Optional...),
		}
	}
	// Add referenced_by information (reverse dependencies)
	if referencedBy, ok := resourceReferencedByMap[resource.Name]; ok && len(referencedBy) > 0 {
		if metadata.Dependencies == nil {
			metadata.Dependencies = &DependencyInfo{}
		}
		metadata.Dependencies.ReferencedBy = referencedBy
	}

	// Store in global collection
	metadataCollection.Resources[resource.Name] = metadata
}

// extractAttributeMetadata recursively extracts metadata from attributes
func extractAttributeMetadata(attrs []TerraformAttribute, output map[string]*AttributeMetadata, oneOfGroups map[string][]string) {
	// Build reverse lookup: field -> group name
	fieldToGroup := make(map[string]string)
	for groupName, fields := range oneOfGroups {
		for _, field := range fields {
			fieldToGroup[field] = groupName
		}
	}

	for _, attr := range attrs {
		attrMeta := &AttributeMetadata{
			Type:             attr.Type,
			Required:         attr.Required,
			Optional:         attr.Optional,
			Computed:         attr.Computed,
			Sensitive:        attr.Sensitive,
			IsBlock:          attr.IsBlock,
			PlanModifier:     attr.PlanModifier,
			Description:      attr.Description,
			ServerDefault:    attr.ServerDefault,
			RecommendedValue: attr.RecommendedValue,
		}

		// Add OneOf group reference if applicable
		if group, ok := fieldToGroup[attr.Name]; ok {
			attrMeta.OneOfGroup = group
		}

		// Determine validation type based on field characteristics
		attrMeta.Validation = inferValidationType(attr)

		output[attr.Name] = attrMeta

		// Note: We don't recursively extract nested attributes to keep the JSON manageable
		// The descriptions contain constraint hints for nested fields
	}
}

// inferValidationType determines the validation pattern name for an attribute
func inferValidationType(attr TerraformAttribute) string {
	switch attr.Name {
	case "name":
		if attr.UseDomainValidator {
			return "domain"
		}
		return "name"
	case "namespace":
		return "name"
	case "port", "listen_port":
		return "port"
	}

	// Check type-based validations
	if attr.Type == "int64" && strings.Contains(strings.ToLower(attr.Description), "port") {
		return "port"
	}

	return ""
}

// =============================================================================
// OPERATION METADATA EXTRACTION (v2.0.33 extensions)
// =============================================================================

// collectOperationMetadata extracts operation-level metadata from a domain spec file for a resource
func collectOperationMetadata(spec *OpenAPI3Spec, resourceName string, basePath string, rawSpec map[string]interface{}) {
	if spec == nil || len(spec.Paths) == 0 {
		return
	}

	resourcePlural := resourceName + "s"
	// Handle special pluralization
	if strings.HasSuffix(resourceName, "y") && !strings.HasSuffix(resourceName, "ey") {
		resourcePlural = strings.TrimSuffix(resourceName, "y") + "ies"
	}

	// Initialize resource operation info
	resourceOps := &ResourceOperationInfo{
		Resource:   resourceName,
		BasePath:   basePath,
		Operations: make(map[string]*OperationMetadata),
	}

	// Get raw paths for extension extraction
	var rawPaths map[string]interface{}
	if components, ok := rawSpec["paths"].(map[string]interface{}); ok {
		rawPaths = components
	}

	// Process each path looking for this resource's operations
	for pathKey, pathValue := range spec.Paths {
		// Check if this path is for our resource
		if !strings.Contains(pathKey, "/"+resourcePlural) && !strings.Contains(pathKey, "/"+resourceName+"s") {
			continue
		}

		pathMap, ok := pathValue.(map[string]interface{})
		if !ok {
			continue
		}

		// Get raw path data for extension extraction
		var rawPathData map[string]interface{}
		if rawPaths != nil {
			if rp, ok := rawPaths[pathKey].(map[string]interface{}); ok {
				rawPathData = rp
			}
		}

		// Process each HTTP method
		for method, methodValue := range pathMap {
			if method == "parameters" {
				continue
			}

			// Skip if method value is not a map (sanity check)
			if _, ok := methodValue.(map[string]interface{}); !ok {
				continue
			}

			// Get raw method data for extension extraction
			var rawMethodData map[string]interface{}
			if rawPathData != nil {
				if rm, ok := rawPathData[method].(map[string]interface{}); ok {
					rawMethodData = rm
				}
			}

			// Determine operation type
			opType := mapMethodToOperationType(method, pathKey)
			if opType == "" {
				continue
			}

			// Extract operation metadata from extensions
			opMeta := extractOperationMetadataFromRaw(method, pathKey, rawMethodData)
			if opMeta != nil {
				resourceOps.Operations[opType] = opMeta
			}
		}

		// Extract best practices from any operation (typically on the base path)
		if rawPathData != nil {
			if bp := extractBestPracticesFromRaw(rawPathData); bp != nil {
				resourceOps.BestPractices = bp
			}
			if workflows := extractGuidedWorkflowsFromRaw(rawPathData); len(workflows) > 0 {
				resourceOps.Workflows = workflows
			}
		}
	}

	// Only store if we found operations
	if len(resourceOps.Operations) > 0 {
		operationsMetadataCollection.Resources[resourceName] = resourceOps
	}
}

// mapMethodToOperationType maps HTTP method and path to CRUD operation type
func mapMethodToOperationType(method string, path string) string {
	hasNameInPath := strings.Contains(path, "{name}") || strings.Contains(path, "{metadata.name}")

	switch strings.ToUpper(method) {
	case "POST":
		return "create"
	case "GET":
		if hasNameInPath {
			return "read"
		}
		return "list"
	case "PUT":
		return "update"
	case "DELETE":
		return "delete"
	default:
		return ""
	}
}

// extractOperationMetadataFromRaw extracts x-f5xc-* extensions from raw operation data
func extractOperationMetadataFromRaw(method string, path string, rawOp map[string]interface{}) *OperationMetadata {
	if rawOp == nil {
		return nil
	}

	opMeta := &OperationMetadata{
		Method: strings.ToUpper(method),
		Path:   path,
	}

	// Extract danger level
	if dl, ok := rawOp["x-f5xc-danger-level"].(string); ok {
		opMeta.DangerLevel = dl
	}

	// Extract required fields
	if rf, ok := rawOp["x-f5xc-required-fields"].([]interface{}); ok {
		for _, f := range rf {
			if s, ok := f.(string); ok {
				opMeta.RequiredFields = append(opMeta.RequiredFields, s)
			}
		}
	}

	// Extract confirmation required
	if cr, ok := rawOp["x-f5xc-confirmation-required"].(bool); ok {
		opMeta.ConfirmationRequired = cr
	}

	// Extract discovered response time
	if rt, ok := rawOp["x-f5xc-discovered-response-time"].(map[string]interface{}); ok {
		opMeta.DiscoveredRespTime = &ResponseTimeInfo{
			P50Ms:       getIntFromInterface(rt["p50_ms"]),
			P95Ms:       getIntFromInterface(rt["p95_ms"]),
			P99Ms:       getIntFromInterface(rt["p99_ms"]),
			SampleCount: getIntFromInterface(rt["sample_count"]),
			Source:      getStringFromInterface(rt["source"]),
		}
	}

	// Extract side effects
	if se, ok := rawOp["x-f5xc-side-effects"].(map[string]interface{}); ok {
		sideEffects := &SideEffectsInfo{}
		if creates, ok := se["creates"].([]interface{}); ok {
			for _, c := range creates {
				if s, ok := c.(string); ok {
					sideEffects.Creates = append(sideEffects.Creates, s)
				}
			}
		}
		if modifies, ok := se["modifies"].([]interface{}); ok {
			for _, m := range modifies {
				if s, ok := m.(string); ok {
					sideEffects.Modifies = append(sideEffects.Modifies, s)
				}
			}
		}
		if deletes, ok := se["deletes"].([]interface{}); ok {
			for _, d := range deletes {
				if s, ok := d.(string); ok {
					sideEffects.Deletes = append(sideEffects.Deletes, s)
				}
			}
		}
		if sideEffects.Creates != nil || sideEffects.Modifies != nil || sideEffects.Deletes != nil {
			opMeta.SideEffects = sideEffects
		}
	}

	// Extract purpose from operation metadata
	if om, ok := rawOp["x-f5xc-operation-metadata"].(map[string]interface{}); ok {
		if purpose, ok := om["purpose"].(string); ok {
			opMeta.Purpose = purpose
		}
	}

	// Only return if we extracted any meaningful metadata
	if opMeta.DangerLevel != "" || opMeta.DiscoveredRespTime != nil ||
		len(opMeta.RequiredFields) > 0 || opMeta.ConfirmationRequired ||
		opMeta.SideEffects != nil || opMeta.Purpose != "" {
		return opMeta
	}
	return nil
}

// extractBestPracticesFromRaw extracts x-f5xc-best-practices from path data
func extractBestPracticesFromRaw(rawPath map[string]interface{}) *BestPracticesInfo {
	// Check each method for best practices (typically on POST)
	for _, methodData := range rawPath {
		methodMap, ok := methodData.(map[string]interface{})
		if !ok {
			continue
		}

		if bp, ok := methodMap["x-f5xc-best-practices"].(map[string]interface{}); ok {
			bestPractices := &BestPracticesInfo{}

			if errors, ok := bp["common_errors"].([]interface{}); ok {
				for _, e := range errors {
					if errMap, ok := e.(map[string]interface{}); ok {
						errInfo := CommonErrorInfo{
							Code:       getIntFromInterface(errMap["code"]),
							Message:    getStringFromInterface(errMap["message"]),
							Resolution: getStringFromInterface(errMap["resolution"]),
							Prevention: getStringFromInterface(errMap["prevention"]),
						}
						bestPractices.CommonErrors = append(bestPractices.CommonErrors, errInfo)
					}
				}
			}

			if len(bestPractices.CommonErrors) > 0 {
				return bestPractices
			}
		}
	}
	return nil
}

// extractGuidedWorkflowsFromRaw extracts x-f5xc-guided-workflows from path data
func extractGuidedWorkflowsFromRaw(rawPath map[string]interface{}) []*GuidedWorkflowInfo {
	var workflows []*GuidedWorkflowInfo

	// Check each method for guided workflows
	for _, methodData := range rawPath {
		methodMap, ok := methodData.(map[string]interface{})
		if !ok {
			continue
		}

		if wf, ok := methodMap["x-f5xc-guided-workflows"].([]interface{}); ok {
			for _, w := range wf {
				if wfMap, ok := w.(map[string]interface{}); ok {
					workflow := &GuidedWorkflowInfo{
						Name:        getStringFromInterface(wfMap["name"]),
						Description: getStringFromInterface(wfMap["description"]),
					}

					if steps, ok := wfMap["steps"].([]interface{}); ok {
						for i, s := range steps {
							if stepMap, ok := s.(map[string]interface{}); ok {
								step := &WorkflowStepInfo{
									Order:       i + 1,
									Action:      getStringFromInterface(stepMap["action"]),
									Description: getStringFromInterface(stepMap["description"]),
									Validation:  getStringFromInterface(stepMap["validation"]),
								}
								if fields, ok := stepMap["fields"].([]interface{}); ok {
									for _, f := range fields {
										if s, ok := f.(string); ok {
											step.Fields = append(step.Fields, s)
										}
									}
								}
								workflow.Steps = append(workflow.Steps, step)
							}
						}
					}

					if workflow.Name != "" {
						workflows = append(workflows, workflow)
					}
				}
			}
		}
	}

	return workflows
}

// getIntFromInterface safely extracts an int from an interface{}
func getIntFromInterface(v interface{}) int {
	switch val := v.(type) {
	case int:
		return val
	case int64:
		return int(val)
	case float64:
		return int(val)
	default:
		return 0
	}
}

// getStringFromInterface safely extracts a string from an interface{}
func getStringFromInterface(v interface{}) string {
	if s, ok := v.(string); ok {
		return s
	}
	return ""
}

// writeMetadataFiles writes the collected metadata to JSON files for MCP server consumption
func writeMetadataFiles() error {
	// Create metadata directory
	metadataDir := filepath.Join("tools", "metadata")
	if err := os.MkdirAll(metadataDir, 0755); err != nil {
		return fmt.Errorf("failed to create metadata directory: %w", err)
	}

	// Set generation metadata
	metadataCollection.GeneratedAt = time.Now().UTC().Format(time.RFC3339)
	metadataCollection.Version = "1.0.0"

	// Write resource-metadata.json
	resourceMetadataPath := filepath.Join(metadataDir, "resource-metadata.json")
	resourceMetadataJSON, err := json.MarshalIndent(metadataCollection, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal resource metadata: %w", err)
	}
	if err := os.WriteFile(resourceMetadataPath, resourceMetadataJSON, 0644); err != nil {
		return fmt.Errorf("failed to write resource metadata: %w", err)
	}
	fmt.Printf("📝 Wrote metadata for %d resources to %s\n", len(metadataCollection.Resources), resourceMetadataPath)

	// Write validation-patterns.json
	validationPatterns := getValidationPatterns()
	validationPatternsPath := filepath.Join(metadataDir, "validation-patterns.json")
	validationPatternsJSON, err := json.MarshalIndent(validationPatterns, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal validation patterns: %w", err)
	}
	if err := os.WriteFile(validationPatternsPath, validationPatternsJSON, 0644); err != nil {
		return fmt.Errorf("failed to write validation patterns: %w", err)
	}
	fmt.Printf("📝 Wrote validation patterns to %s\n", validationPatternsPath)

	// Write operations-metadata.json (v2.0.33 extensions)
	if len(operationsMetadataCollection.Resources) > 0 {
		operationsMetadataCollection.GeneratedAt = time.Now().UTC().Format(time.RFC3339)
		operationsMetadataCollection.Version = "1.0.0"

		operationsMetadataPath := filepath.Join(metadataDir, "operations-metadata.json")
		operationsMetadataJSON, err := json.MarshalIndent(operationsMetadataCollection, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal operations metadata: %w", err)
		}
		if err := os.WriteFile(operationsMetadataPath, operationsMetadataJSON, 0644); err != nil {
			return fmt.Errorf("failed to write operations metadata: %w", err)
		}
		fmt.Printf("📝 Wrote operations metadata for %d resources to %s\n", len(operationsMetadataCollection.Resources), operationsMetadataPath)
	}

	return nil
}

// ValidationPatterns defines validation patterns for MCP server
type ValidationPatterns struct {
	Version  string                        `json:"version"`
	Patterns map[string]*ValidationPattern `json:"patterns"`
}

// ValidationPattern describes a single validation rule
type ValidationPattern struct {
	Type        string   `json:"type"`               // "regex" or "range"
	Pattern     string   `json:"pattern,omitempty"`  // Regex pattern
	Min         *int64   `json:"min,omitempty"`      // Minimum value for range
	Max         *int64   `json:"max,omitempty"`      // Maximum value for range
	Description string   `json:"description"`        // Human-readable description
	Examples    []string `json:"examples,omitempty"` // Valid example values
	Invalid     []string `json:"invalid,omitempty"`  // Invalid example values
}

// getValidationPatterns returns the validation patterns used by the provider
// These are extracted from internal/validators/validators.go
func getValidationPatterns() *ValidationPatterns {
	minPort := int64(1)
	maxPort := int64(65535)

	return &ValidationPatterns{
		Version: "1.0.0",
		Patterns: map[string]*ValidationPattern{
			"name": {
				Type:        "regex",
				Pattern:     `^[a-z][a-z0-9-]{0,62}[a-z0-9]$|^[a-z]$`,
				Description: "Resource name: lowercase alphanumeric with hyphens, 1-64 characters, must start with letter and end with letter or number",
				Examples:    []string{"my-resource", "example-lb", "ns1", "a"},
				Invalid:     []string{"My-Resource", "123-start", "-invalid", "too-long-name-that-exceeds-sixty-four-characters-limit-for-names"},
			},
			"domain": {
				Type:        "regex",
				Pattern:     `^(\*\.)?[a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(\.[a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`,
				Description: "Valid DNS domain name with optional wildcard prefix",
				Examples:    []string{"example.com", "*.example.com", "sub.domain.example.com"},
				Invalid:     []string{"-invalid.com", "example..com", ""},
			},
			"port": {
				Type:        "range",
				Min:         &minPort,
				Max:         &maxPort,
				Description: "Valid TCP/UDP port number between 1 and 65535",
				Examples:    []string{"80", "443", "8080", "65535"},
				Invalid:     []string{"0", "65536", "-1"},
			},
			"namespace": {
				Type:        "regex",
				Pattern:     `^[a-z][a-z0-9-]{0,62}[a-z0-9]$|^[a-z]$`,
				Description: "Namespace name: same rules as resource name",
				Examples:    []string{"default", "production", "dev-env"},
				Invalid:     []string{"Default", "123", "_invalid"},
			},
		},
	}
}

// addOneOfConstraint adds a OneOf constraint hint to the description with recommended default
// The format is AI-friendly: [OneOf: field1, field2; Default: recommended_field]
func addOneOfConstraint(desc string, oneOfFields []string) string {
	return addOneOfConstraintWithGroup(desc, "", oneOfFields)
}

// addOneOfConstraintWithGroup adds a OneOf constraint with group name for better AI parsing
func addOneOfConstraintWithGroup(desc string, groupName string, oneOfFields []string) string {
	if len(oneOfFields) < 2 {
		return desc
	}

	// Format fields
	quotedFields := make([]string, len(oneOfFields))
	for i, f := range oneOfFields {
		quotedFields[i] = f
	}

	// Determine recommended default choice using AI-friendly heuristics
	defaultChoice := determineOneOfDefault(groupName, oneOfFields)

	// Build AI-friendly constraint marker
	var constraint string
	if defaultChoice != "" {
		constraint = fmt.Sprintf("[OneOf: %s; Default: %s]", strings.Join(quotedFields, ", "), defaultChoice)
	} else {
		constraint = fmt.Sprintf("[OneOf: %s]", strings.Join(quotedFields, ", "))
	}

	// Add constraint at the beginning of description
	if desc == "" {
		return constraint
	}
	return constraint + " " + desc
}

// determineOneOfDefault determines the recommended default for a OneOf group
// This helps AI agents make informed decisions without trial-and-error
func determineOneOfDefault(groupName string, fields []string) string {
	// Common patterns for F5 XC resources - these are the safe, recommended defaults
	defaultPatterns := map[string]string{
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

	// Check if we have a known pattern for this group
	if groupName != "" {
		if defaultVal, ok := defaultPatterns[groupName]; ok {
			// Verify the default is actually in the fields
			for _, f := range fields {
				if f == defaultVal {
					return defaultVal
				}
			}
		}
	}

	// Heuristics for common patterns
	for _, f := range fields {
		// Prefer "default" variants (e.g., advertise_on_public_default_vip)
		if strings.Contains(f, "default") {
			return f
		}
	}

	for _, f := range fields {
		// Prefer "no_" or "disable_" for optional security features
		if strings.HasPrefix(f, "no_") || strings.HasPrefix(f, "disable_") {
			return f
		}
	}

	// No clear default - don't recommend one
	return ""
}

// toTitleCase wraps naming.ToResourceTypeName for backward compatibility.
// Converts a snake_case resource name to a PascalCase Go type name.
// Example: "http_loadbalancer" -> "HTTPLoadBalancer"
func toTitleCase(s string) string {
	return naming.ToResourceTypeName(s)
}

// toHumanName converts a resource name to human-readable format for documentation.
// Unlike toTitleCase which produces Go type names, this produces readable names with spaces.
// Used in transformResourceDescription() to generate user-friendly resource descriptions.
// Example: "http_loadbalancer" -> "HTTP Load Balancer"
func toHumanName(s string) string {
	return naming.ToHumanReadableName(s)
}

// toSnakeCase converts a string to lowercase snake_case for Terraform attribute names
func toSnakeCase(s string) string {
	// First, insert underscores before uppercase letters that follow lowercase letters
	// e.g., "AWAFPayG3Gbps" -> "AWAF_Pay_G3_Gbps"
	var result strings.Builder
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			prev := rune(s[i-1])
			// Add underscore if previous char is lowercase or a digit
			if (prev >= 'a' && prev <= 'z') || (prev >= '0' && prev <= '9') {
				result.WriteRune('_')
			}
		}
		result.WriteRune(r)
	}
	// Convert to lowercase and replace any existing underscores
	name := strings.ToLower(result.String())

	// Handle reserved Terraform attribute names
	reservedNames := map[string]string{
		"count":       "item_count",
		"for_each":    "for_each_items",
		"depends_on":  "depends_on_refs",
		"provider":    "provider_ref",
		"lifecycle":   "lifecycle_config",
		"provisioner": "provisioner_config",
		"connection":  "connection_config",
		"locals":      "locals_config",
	}

	if replacement, ok := reservedNames[name]; ok {
		return replacement
	}
	return name
}

// isMetadataField returns true if the field is a metadata field (not a spec field)
func isMetadataField(name string) bool {
	metadataFields := map[string]bool{
		"name":        true,
		"namespace":   true,
		"labels":      true,
		"annotations": true,
		"description": true,
		"id":          true,
		"timeouts":    true,
	}
	return metadataFields[name]
}

// filterSpecFields returns only the spec fields (non-metadata) from attributes
func filterSpecFields(attrs []TerraformAttribute) []TerraformAttribute {
	var specFields []TerraformAttribute
	for _, attr := range attrs {
		if !isMetadataField(attr.TfsdkTag) && attr.IsSpecField {
			specFields = append(specFields, attr)
		}
	}
	return specFields
}

// renderSpecStructFields generates Go struct fields for the client Spec type
func renderSpecStructFields(attrs []TerraformAttribute, indent string) string {
	specFields := filterSpecFields(attrs)
	if len(specFields) == 0 {
		return ""
	}

	var sb strings.Builder
	for _, attr := range specFields {
		goType := getGoClientType(attr)
		jsonTag := attr.JsonName
		if jsonTag == "" {
			jsonTag = attr.TfsdkTag
		}
		// For nested blocks, don't use omitempty - the API needs empty objects to be sent
		// For primitive fields, use omitempty to avoid sending zero values
		if attr.IsBlock {
			sb.WriteString(fmt.Sprintf("%s%s %s `json:\"%s\"`\n", indent, attr.GoName, goType, jsonTag))
		} else {
			sb.WriteString(fmt.Sprintf("%s%s %s `json:\"%s,omitempty\"`\n", indent, attr.GoName, goType, jsonTag))
		}
	}
	return sb.String()
}

// getGoClientType returns the Go type for use in client structs
func getGoClientType(attr TerraformAttribute) string {
	if attr.IsBlock {
		// Nested blocks become pointers to nested structs or slices
		if attr.NestedBlockType == "list" {
			return "[]map[string]interface{}"
		}
		return "map[string]interface{}"
	}

	switch attr.Type {
	case "string":
		return "string"
	case "int64":
		return "int64"
	case "bool":
		return "bool"
	case "list":
		if attr.ElementType == "string" {
			return "[]string"
		} else if attr.ElementType == "int64" {
			return "[]int64"
		}
		return "[]interface{}"
	case "map":
		return "map[string]string"
	default:
		return "interface{}"
	}
}

// renderSpecMarshalCodeForCreate generates Go code for Create (uses "createReq" variable)
func renderSpecMarshalCodeForCreate(attrs []TerraformAttribute, indent string, resourceTitleCase string) string {
	return renderSpecMarshalCodeWithVar(attrs, indent, "createReq", resourceTitleCase)
}

// renderSpecMarshalCode generates Go code to marshal spec fields from Terraform state to API struct (uses "apiResource" variable)
func renderSpecMarshalCode(attrs []TerraformAttribute, indent string, resourceTitleCase string) string {
	return renderSpecMarshalCodeWithVar(attrs, indent, "apiResource", resourceTitleCase)
}

// renderSpecMarshalCodeWithVar generates Go code to marshal spec fields with configurable variable name
func renderSpecMarshalCodeWithVar(attrs []TerraformAttribute, indent string, varName string, resourceTitleCase string) string {
	specFields := filterSpecFields(attrs)
	if len(specFields) == 0 {
		return ""
	}

	var sb strings.Builder
	for _, attr := range specFields {
		fieldName := attr.GoName
		tfsdkName := attr.TfsdkTag
		jsonName := attr.JsonName
		if jsonName == "" {
			jsonName = tfsdkName
		}

		if attr.IsBlock {
			// Handle list nested blocks
			if attr.NestedBlockType == "list" {
				// Check if there are any attributes that need the item variable
				// (primitives or nested blocks)
				needsItemVar := false
				for _, nestedAttr := range attr.NestedAttributes {
					switch nestedAttr.Type {
					case "string", "int64", "bool":
						needsItemVar = true
					}
					// Nested blocks also need access to item
					if nestedAttr.IsBlock {
						needsItemVar = true
					}
				}

				// Extract types.List to Go slice for iteration
				sb.WriteString(fmt.Sprintf("%sif !data.%s.IsNull() && !data.%s.IsUnknown() {\n", indent, fieldName, fieldName))
				sb.WriteString(fmt.Sprintf("%s\tvar %sItems []%s%sModel\n", indent, tfsdkName, resourceTitleCase, attr.GoName))
				sb.WriteString(fmt.Sprintf("%s\tdiags := data.%s.ElementsAs(ctx, &%sItems, false)\n", indent, fieldName, tfsdkName))
				sb.WriteString(fmt.Sprintf("%s\tresp.Diagnostics.Append(diags...)\n", indent))
				sb.WriteString(fmt.Sprintf("%s\tif !resp.Diagnostics.HasError() && len(%sItems) > 0 {\n", indent, tfsdkName))
				sb.WriteString(fmt.Sprintf("%s\t\tvar %sList []map[string]interface{}\n", indent, tfsdkName))
				if needsItemVar {
					sb.WriteString(fmt.Sprintf("%s\t\tfor _, item := range %sItems {\n", indent, tfsdkName))
				} else {
					sb.WriteString(fmt.Sprintf("%s\t\tfor range %sItems {\n", indent, tfsdkName))
				}
				sb.WriteString(fmt.Sprintf("%s\t\t\titemMap := make(map[string]interface{})\n", indent))
				// Marshal nested block fields
				for _, nestedAttr := range attr.NestedAttributes {
					nestedFieldName := nestedAttr.GoName
					nestedTfsdkName := nestedAttr.TfsdkTag
					nestedJsonName := nestedAttr.JsonName
					if nestedJsonName == "" {
						nestedJsonName = nestedTfsdkName
					}

					// Handle single nested blocks within list items
					if nestedAttr.IsBlock && nestedAttr.NestedBlockType == "single" {
						sb.WriteString(fmt.Sprintf("%s\t\tif item.%s != nil {\n", indent, nestedFieldName))
						if len(nestedAttr.NestedAttributes) == 0 {
							// Empty block - just send empty map
							sb.WriteString(fmt.Sprintf("%s\t\t\titemMap[\"%s\"] = map[string]interface{}{}\n", indent, nestedJsonName))
						} else {
							// Block with nested attributes - marshal them
							sb.WriteString(fmt.Sprintf("%s\t\t\t%sNestedMap := make(map[string]interface{})\n", indent, nestedTfsdkName))
							for _, deepAttr := range nestedAttr.NestedAttributes {
								deepFieldName := deepAttr.GoName
								deepTfsdkName := deepAttr.TfsdkTag
								deepJsonName := deepAttr.JsonName
								if deepJsonName == "" {
									deepJsonName = deepAttr.TfsdkTag
								}
								// Handle nested blocks within SingleNestedBlocks that are within ListNestedBlock items
								// e.g., allow{}, deny{}, ip_prefixes{} within action{} or match{} within rules[]
								if deepAttr.IsBlock && deepAttr.NestedBlockType == "single" {
									sb.WriteString(fmt.Sprintf("%s\t\t\tif item.%s.%s != nil {\n", indent, nestedFieldName, deepFieldName))
									if len(deepAttr.NestedAttributes) == 0 {
										// Empty nested block
										sb.WriteString(fmt.Sprintf("%s\t\t\t\t%sNestedMap[\"%s\"] = map[string]interface{}{}\n", indent, nestedTfsdkName, deepJsonName))
									} else {
										// Nested block with attributes - create a deep map
										sb.WriteString(fmt.Sprintf("%s\t\t\t\t%sDeepMap := make(map[string]interface{})\n", indent, deepTfsdkName))
										for _, leafAttr := range deepAttr.NestedAttributes {
											leafFieldName := leafAttr.GoName
											leafJsonName := leafAttr.JsonName
											if leafJsonName == "" {
												leafJsonName = leafAttr.TfsdkTag
											}
											// Handle empty blocks at leaf level
											if leafAttr.IsBlock && leafAttr.NestedBlockType == "single" && len(leafAttr.NestedAttributes) == 0 {
												sb.WriteString(fmt.Sprintf("%s\t\t\t\tif item.%s.%s.%s != nil {\n", indent, nestedFieldName, deepFieldName, leafFieldName))
												sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t%sDeepMap[\"%s\"] = map[string]interface{}{}\n", indent, deepTfsdkName, leafJsonName))
												sb.WriteString(fmt.Sprintf("%s\t\t\t\t}\n", indent))
											} else {
												switch leafAttr.Type {
												case "string":
													sb.WriteString(fmt.Sprintf("%s\t\t\t\tif !item.%s.%s.%s.IsNull() && !item.%s.%s.%s.IsUnknown() {\n", indent, nestedFieldName, deepFieldName, leafFieldName, nestedFieldName, deepFieldName, leafFieldName))
													sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t%sDeepMap[\"%s\"] = item.%s.%s.%s.ValueString()\n", indent, deepTfsdkName, leafJsonName, nestedFieldName, deepFieldName, leafFieldName))
													sb.WriteString(fmt.Sprintf("%s\t\t\t\t}\n", indent))
												case "int64":
													sb.WriteString(fmt.Sprintf("%s\t\t\t\tif !item.%s.%s.%s.IsNull() && !item.%s.%s.%s.IsUnknown() {\n", indent, nestedFieldName, deepFieldName, leafFieldName, nestedFieldName, deepFieldName, leafFieldName))
													sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t%sDeepMap[\"%s\"] = item.%s.%s.%s.ValueInt64()\n", indent, deepTfsdkName, leafJsonName, nestedFieldName, deepFieldName, leafFieldName))
													sb.WriteString(fmt.Sprintf("%s\t\t\t\t}\n", indent))
												case "bool":
													sb.WriteString(fmt.Sprintf("%s\t\t\t\tif !item.%s.%s.%s.IsNull() && !item.%s.%s.%s.IsUnknown() {\n", indent, nestedFieldName, deepFieldName, leafFieldName, nestedFieldName, deepFieldName, leafFieldName))
													sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t%sDeepMap[\"%s\"] = item.%s.%s.%s.ValueBool()\n", indent, deepTfsdkName, leafJsonName, nestedFieldName, deepFieldName, leafFieldName))
													sb.WriteString(fmt.Sprintf("%s\t\t\t\t}\n", indent))
												case "list":
													if leafAttr.ElementType == "string" {
														sb.WriteString(fmt.Sprintf("%s\t\t\t\tif !item.%s.%s.%s.IsNull() && !item.%s.%s.%s.IsUnknown() {\n", indent, nestedFieldName, deepFieldName, leafFieldName, nestedFieldName, deepFieldName, leafFieldName))
														sb.WriteString(fmt.Sprintf("%s\t\t\t\t\tvar %sItems []string\n", indent, leafFieldName))
														sb.WriteString(fmt.Sprintf("%s\t\t\t\t\tdiags := item.%s.%s.%s.ElementsAs(ctx, &%sItems, false)\n", indent, nestedFieldName, deepFieldName, leafFieldName, leafFieldName))
														sb.WriteString(fmt.Sprintf("%s\t\t\t\t\tif !diags.HasError() {\n", indent))
														sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t%sDeepMap[\"%s\"] = %sItems\n", indent, deepTfsdkName, leafJsonName, leafFieldName))
														sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t}\n", indent))
														sb.WriteString(fmt.Sprintf("%s\t\t\t\t}\n", indent))
													}
												}
											}
										}
										sb.WriteString(fmt.Sprintf("%s\t\t\t\t%sNestedMap[\"%s\"] = %sDeepMap\n", indent, nestedTfsdkName, deepJsonName, deepTfsdkName))
									}
									sb.WriteString(fmt.Sprintf("%s\t\t\t}\n", indent))
									continue
								}
								// Handle ListNestedBlocks within SingleNestedBlocks that are within ListNestedBlock items
								// e.g., prefixes[] within ip_prefixes{} within match{} within rules[]
								if deepAttr.IsBlock && deepAttr.NestedBlockType == "list" {
									sb.WriteString(fmt.Sprintf("%s\t\t\tif len(item.%s.%s) > 0 {\n", indent, nestedFieldName, deepFieldName))
									sb.WriteString(fmt.Sprintf("%s\t\t\t\tvar %sDeepList []map[string]interface{}\n", indent, deepTfsdkName))
									sb.WriteString(fmt.Sprintf("%s\t\t\t\tfor _, deepListItem := range item.%s.%s {\n", indent, nestedFieldName, deepFieldName))
									sb.WriteString(fmt.Sprintf("%s\t\t\t\t\tdeepListItemMap := make(map[string]interface{})\n", indent))
									for _, leafAttr := range deepAttr.NestedAttributes {
										leafFieldName := leafAttr.GoName
										leafJsonName := leafAttr.JsonName
										if leafJsonName == "" {
											leafJsonName = leafAttr.TfsdkTag
										}
										// Handle empty blocks at leaf level
										if leafAttr.IsBlock && leafAttr.NestedBlockType == "single" && len(leafAttr.NestedAttributes) == 0 {
											sb.WriteString(fmt.Sprintf("%s\t\t\t\t\tif deepListItem.%s != nil {\n", indent, leafFieldName))
											sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\tdeepListItemMap[\"%s\"] = map[string]interface{}{}\n", indent, leafJsonName))
											sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t}\n", indent))
										} else {
											switch leafAttr.Type {
											case "string":
												sb.WriteString(fmt.Sprintf("%s\t\t\t\t\tif !deepListItem.%s.IsNull() && !deepListItem.%s.IsUnknown() {\n", indent, leafFieldName, leafFieldName))
												sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\tdeepListItemMap[\"%s\"] = deepListItem.%s.ValueString()\n", indent, leafJsonName, leafFieldName))
												sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t}\n", indent))
											case "int64":
												sb.WriteString(fmt.Sprintf("%s\t\t\t\t\tif !deepListItem.%s.IsNull() && !deepListItem.%s.IsUnknown() {\n", indent, leafFieldName, leafFieldName))
												sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\tdeepListItemMap[\"%s\"] = deepListItem.%s.ValueInt64()\n", indent, leafJsonName, leafFieldName))
												sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t}\n", indent))
											case "bool":
												sb.WriteString(fmt.Sprintf("%s\t\t\t\t\tif !deepListItem.%s.IsNull() && !deepListItem.%s.IsUnknown() {\n", indent, leafFieldName, leafFieldName))
												sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\tdeepListItemMap[\"%s\"] = deepListItem.%s.ValueBool()\n", indent, leafJsonName, leafFieldName))
												sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t}\n", indent))
											}
										}
									}
									sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t%sDeepList = append(%sDeepList, deepListItemMap)\n", indent, deepTfsdkName, deepTfsdkName))
									sb.WriteString(fmt.Sprintf("%s\t\t\t\t}\n", indent))
									sb.WriteString(fmt.Sprintf("%s\t\t\t\t%sNestedMap[\"%s\"] = %sDeepList\n", indent, nestedTfsdkName, deepJsonName, deepTfsdkName))
									sb.WriteString(fmt.Sprintf("%s\t\t\t}\n", indent))
									continue
								}
								switch deepAttr.Type {
								case "string":
									sb.WriteString(fmt.Sprintf("%s\t\t\tif !item.%s.%s.IsNull() && !item.%s.%s.IsUnknown() {\n", indent, nestedFieldName, deepFieldName, nestedFieldName, deepFieldName))
									sb.WriteString(fmt.Sprintf("%s\t\t\t\t%sNestedMap[\"%s\"] = item.%s.%s.ValueString()\n", indent, nestedTfsdkName, deepJsonName, nestedFieldName, deepFieldName))
									sb.WriteString(fmt.Sprintf("%s\t\t\t}\n", indent))
								case "int64":
									sb.WriteString(fmt.Sprintf("%s\t\t\tif !item.%s.%s.IsNull() && !item.%s.%s.IsUnknown() {\n", indent, nestedFieldName, deepFieldName, nestedFieldName, deepFieldName))
									sb.WriteString(fmt.Sprintf("%s\t\t\t\t%sNestedMap[\"%s\"] = item.%s.%s.ValueInt64()\n", indent, nestedTfsdkName, deepJsonName, nestedFieldName, deepFieldName))
									sb.WriteString(fmt.Sprintf("%s\t\t\t}\n", indent))
								case "bool":
									sb.WriteString(fmt.Sprintf("%s\t\t\tif !item.%s.%s.IsNull() && !item.%s.%s.IsUnknown() {\n", indent, nestedFieldName, deepFieldName, nestedFieldName, deepFieldName))
									sb.WriteString(fmt.Sprintf("%s\t\t\t\t%sNestedMap[\"%s\"] = item.%s.%s.ValueBool()\n", indent, nestedTfsdkName, deepJsonName, nestedFieldName, deepFieldName))
									sb.WriteString(fmt.Sprintf("%s\t\t\t}\n", indent))
								case "list":
									// Handle list types inside single nested blocks within list items
									if deepAttr.ElementType == "string" {
										sb.WriteString(fmt.Sprintf("%s\t\t\tif !item.%s.%s.IsNull() && !item.%s.%s.IsUnknown() {\n", indent, nestedFieldName, deepFieldName, nestedFieldName, deepFieldName))
										sb.WriteString(fmt.Sprintf("%s\t\t\t\tvar %sItems []string\n", indent, deepFieldName))
										sb.WriteString(fmt.Sprintf("%s\t\t\t\tdiags := item.%s.%s.ElementsAs(ctx, &%sItems, false)\n", indent, nestedFieldName, deepFieldName, deepFieldName))
										sb.WriteString(fmt.Sprintf("%s\t\t\t\tif !diags.HasError() {\n", indent))
										sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t%sNestedMap[\"%s\"] = %sItems\n", indent, nestedTfsdkName, deepJsonName, deepFieldName))
										sb.WriteString(fmt.Sprintf("%s\t\t\t\t}\n", indent))
										sb.WriteString(fmt.Sprintf("%s\t\t\t}\n", indent))
									} else if deepAttr.ElementType == "int64" {
										sb.WriteString(fmt.Sprintf("%s\t\t\tif !item.%s.%s.IsNull() && !item.%s.%s.IsUnknown() {\n", indent, nestedFieldName, deepFieldName, nestedFieldName, deepFieldName))
										sb.WriteString(fmt.Sprintf("%s\t\t\t\tvar %sItems []int64\n", indent, deepFieldName))
										sb.WriteString(fmt.Sprintf("%s\t\t\t\tdiags := item.%s.%s.ElementsAs(ctx, &%sItems, false)\n", indent, nestedFieldName, deepFieldName, deepFieldName))
										sb.WriteString(fmt.Sprintf("%s\t\t\t\tif !diags.HasError() {\n", indent))
										sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t%sNestedMap[\"%s\"] = %sItems\n", indent, nestedTfsdkName, deepJsonName, deepFieldName))
										sb.WriteString(fmt.Sprintf("%s\t\t\t\t}\n", indent))
										sb.WriteString(fmt.Sprintf("%s\t\t\t}\n", indent))
									}
								}
							}
							sb.WriteString(fmt.Sprintf("%s\t\t\titemMap[\"%s\"] = %sNestedMap\n", indent, nestedJsonName, nestedTfsdkName))
						}
						sb.WriteString(fmt.Sprintf("%s\t\t}\n", indent))
						continue
					}

					// Handle list nested blocks within list items (e.g., app_type_ref within app_type_settings)
					if nestedAttr.IsBlock && nestedAttr.NestedBlockType == "list" {
						sb.WriteString(fmt.Sprintf("%s\t\tif len(item.%s) > 0 {\n", indent, nestedFieldName))
						sb.WriteString(fmt.Sprintf("%s\t\t\tvar %sNestedList []map[string]interface{}\n", indent, nestedTfsdkName))
						sb.WriteString(fmt.Sprintf("%s\t\t\tfor _, nestedItem := range item.%s {\n", indent, nestedFieldName))
						sb.WriteString(fmt.Sprintf("%s\t\t\t\tnestedItemMap := make(map[string]interface{})\n", indent))
						// Marshal fields from the nested list block
						for _, deepAttr := range nestedAttr.NestedAttributes {
							deepFieldName := deepAttr.GoName
							deepJsonName := deepAttr.JsonName
							if deepJsonName == "" {
								deepJsonName = deepAttr.TfsdkTag
							}
							switch deepAttr.Type {
							case "string":
								sb.WriteString(fmt.Sprintf("%s\t\t\t\tif !nestedItem.%s.IsNull() && !nestedItem.%s.IsUnknown() {\n", indent, deepFieldName, deepFieldName))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\tnestedItemMap[\"%s\"] = nestedItem.%s.ValueString()\n", indent, deepJsonName, deepFieldName))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t}\n", indent))
							case "int64":
								sb.WriteString(fmt.Sprintf("%s\t\t\t\tif !nestedItem.%s.IsNull() && !nestedItem.%s.IsUnknown() {\n", indent, deepFieldName, deepFieldName))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\tnestedItemMap[\"%s\"] = nestedItem.%s.ValueInt64()\n", indent, deepJsonName, deepFieldName))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t}\n", indent))
							case "bool":
								sb.WriteString(fmt.Sprintf("%s\t\t\t\tif !nestedItem.%s.IsNull() && !nestedItem.%s.IsUnknown() {\n", indent, deepFieldName, deepFieldName))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\tnestedItemMap[\"%s\"] = nestedItem.%s.ValueBool()\n", indent, deepJsonName, deepFieldName))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t}\n", indent))
							}
						}
						sb.WriteString(fmt.Sprintf("%s\t\t\t\t%sNestedList = append(%sNestedList, nestedItemMap)\n", indent, nestedTfsdkName, nestedTfsdkName))
						sb.WriteString(fmt.Sprintf("%s\t\t\t}\n", indent))
						sb.WriteString(fmt.Sprintf("%s\t\t\titemMap[\"%s\"] = %sNestedList\n", indent, nestedJsonName, nestedTfsdkName))
						sb.WriteString(fmt.Sprintf("%s\t\t}\n", indent))
						continue
					}

					switch nestedAttr.Type {
					case "string":
						sb.WriteString(fmt.Sprintf("%s\t\tif !item.%s.IsNull() && !item.%s.IsUnknown() {\n", indent, nestedFieldName, nestedFieldName))
						sb.WriteString(fmt.Sprintf("%s\t\t\titemMap[\"%s\"] = item.%s.ValueString()\n", indent, nestedJsonName, nestedFieldName))
						sb.WriteString(fmt.Sprintf("%s\t\t}\n", indent))
					case "int64":
						sb.WriteString(fmt.Sprintf("%s\t\tif !item.%s.IsNull() && !item.%s.IsUnknown() {\n", indent, nestedFieldName, nestedFieldName))
						sb.WriteString(fmt.Sprintf("%s\t\t\titemMap[\"%s\"] = item.%s.ValueInt64()\n", indent, nestedJsonName, nestedFieldName))
						sb.WriteString(fmt.Sprintf("%s\t\t}\n", indent))
					case "bool":
						sb.WriteString(fmt.Sprintf("%s\t\tif !item.%s.IsNull() && !item.%s.IsUnknown() {\n", indent, nestedFieldName, nestedFieldName))
						sb.WriteString(fmt.Sprintf("%s\t\t\titemMap[\"%s\"] = item.%s.ValueBool()\n", indent, nestedJsonName, nestedFieldName))
						sb.WriteString(fmt.Sprintf("%s\t\t}\n", indent))
					}
				}
				sb.WriteString(fmt.Sprintf("%s\t\t\t%sList = append(%sList, itemMap)\n", indent, tfsdkName, tfsdkName))
				sb.WriteString(fmt.Sprintf("%s\t\t}\n", indent))
				sb.WriteString(fmt.Sprintf("%s\t\t%s.Spec[\"%s\"] = %sList\n", indent, varName, jsonName, tfsdkName))
				sb.WriteString(fmt.Sprintf("%s\t}\n", indent))
				sb.WriteString(fmt.Sprintf("%s}\n", indent))
				continue
			}
			// Handle single nested blocks by converting model struct to map
			sb.WriteString(fmt.Sprintf("%sif data.%s != nil {\n", indent, fieldName))
			sb.WriteString(fmt.Sprintf("%s\t%sMap := make(map[string]interface{})\n", indent, tfsdkName))
			// Marshal nested block fields
			for _, nestedAttr := range attr.NestedAttributes {
				nestedFieldName := nestedAttr.GoName
				nestedTfsdkName := nestedAttr.TfsdkTag
				nestedJsonName := nestedAttr.JsonName
				if nestedJsonName == "" {
					nestedJsonName = nestedTfsdkName
				}

				// Handle single nested blocks within single nested blocks (not list nested blocks)
				if nestedAttr.IsBlock && nestedAttr.NestedBlockType == "single" {
					if len(nestedAttr.NestedAttributes) == 0 {
						// Empty nested block - send empty map if present
						sb.WriteString(fmt.Sprintf("%s\tif data.%s.%s != nil {\n", indent, fieldName, nestedFieldName))
						sb.WriteString(fmt.Sprintf("%s\t\t%sMap[\"%s\"] = map[string]interface{}{}\n", indent, tfsdkName, nestedJsonName))
						sb.WriteString(fmt.Sprintf("%s\t}\n", indent))
					} else {
						// Nested block with attributes - marshal them
						sb.WriteString(fmt.Sprintf("%s\tif data.%s.%s != nil {\n", indent, fieldName, nestedFieldName))
						sb.WriteString(fmt.Sprintf("%s\t\t%sNestedMap := make(map[string]interface{})\n", indent, nestedTfsdkName))
						for _, deepAttr := range nestedAttr.NestedAttributes {
							deepFieldName := deepAttr.GoName
							deepJsonName := deepAttr.JsonName
							if deepJsonName == "" {
								deepJsonName = deepAttr.TfsdkTag
							}
							switch deepAttr.Type {
							case "string":
								sb.WriteString(fmt.Sprintf("%s\t\tif !data.%s.%s.%s.IsNull() && !data.%s.%s.%s.IsUnknown() {\n", indent, fieldName, nestedFieldName, deepFieldName, fieldName, nestedFieldName, deepFieldName))
								sb.WriteString(fmt.Sprintf("%s\t\t\t%sNestedMap[\"%s\"] = data.%s.%s.%s.ValueString()\n", indent, nestedTfsdkName, deepJsonName, fieldName, nestedFieldName, deepFieldName))
								sb.WriteString(fmt.Sprintf("%s\t\t}\n", indent))
							case "int64":
								sb.WriteString(fmt.Sprintf("%s\t\tif !data.%s.%s.%s.IsNull() && !data.%s.%s.%s.IsUnknown() {\n", indent, fieldName, nestedFieldName, deepFieldName, fieldName, nestedFieldName, deepFieldName))
								sb.WriteString(fmt.Sprintf("%s\t\t\t%sNestedMap[\"%s\"] = data.%s.%s.%s.ValueInt64()\n", indent, nestedTfsdkName, deepJsonName, fieldName, nestedFieldName, deepFieldName))
								sb.WriteString(fmt.Sprintf("%s\t\t}\n", indent))
							case "bool":
								sb.WriteString(fmt.Sprintf("%s\t\tif !data.%s.%s.%s.IsNull() && !data.%s.%s.%s.IsUnknown() {\n", indent, fieldName, nestedFieldName, deepFieldName, fieldName, nestedFieldName, deepFieldName))
								sb.WriteString(fmt.Sprintf("%s\t\t\t%sNestedMap[\"%s\"] = data.%s.%s.%s.ValueBool()\n", indent, nestedTfsdkName, deepJsonName, fieldName, nestedFieldName, deepFieldName))
								sb.WriteString(fmt.Sprintf("%s\t\t}\n", indent))
							}
						}
						sb.WriteString(fmt.Sprintf("%s\t\t%sMap[\"%s\"] = %sNestedMap\n", indent, tfsdkName, nestedJsonName, nestedTfsdkName))
						sb.WriteString(fmt.Sprintf("%s\t}\n", indent))
					}
					continue
				}

				// Handle list nested blocks within single nested blocks (e.g., rules within mitigation_type)
				if nestedAttr.IsBlock && nestedAttr.NestedBlockType == "list" {
					sb.WriteString(fmt.Sprintf("%s\tif len(data.%s.%s) > 0 {\n", indent, fieldName, nestedFieldName))
					sb.WriteString(fmt.Sprintf("%s\t\tvar %sList []map[string]interface{}\n", indent, nestedTfsdkName))
					sb.WriteString(fmt.Sprintf("%s\t\tfor _, listItem := range data.%s.%s {\n", indent, fieldName, nestedFieldName))
					sb.WriteString(fmt.Sprintf("%s\t\t\tlistItemMap := make(map[string]interface{})\n", indent))

					// Marshal deep nested attributes within the list items
					for _, deepAttr := range nestedAttr.NestedAttributes {
						deepFieldName := deepAttr.GoName
						deepTfsdkName := deepAttr.TfsdkTag
						deepJsonName := deepAttr.JsonName
						if deepJsonName == "" {
							deepJsonName = deepTfsdkName
						}

						// Handle single nested blocks within list items
						if deepAttr.IsBlock && deepAttr.NestedBlockType == "single" {
							sb.WriteString(fmt.Sprintf("%s\t\t\tif listItem.%s != nil {\n", indent, deepFieldName))
							if len(deepAttr.NestedAttributes) == 0 {
								// Empty block - just send empty map
								sb.WriteString(fmt.Sprintf("%s\t\t\t\tlistItemMap[\"%s\"] = map[string]interface{}{}\n", indent, deepJsonName))
							} else {
								// Block with nested attributes (like choice blocks)
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t%sDeepMap := make(map[string]interface{})\n", indent, deepTfsdkName))
								for _, leafAttr := range deepAttr.NestedAttributes {
									leafFieldName := leafAttr.GoName
									leafTfsdkName := leafAttr.TfsdkTag
									leafJsonName := leafAttr.JsonName
									if leafJsonName == "" {
										leafJsonName = leafTfsdkName
									}
									// Handle empty choice blocks
									if leafAttr.IsBlock && leafAttr.NestedBlockType == "single" && len(leafAttr.NestedAttributes) == 0 {
										sb.WriteString(fmt.Sprintf("%s\t\t\t\tif listItem.%s.%s != nil {\n", indent, deepFieldName, leafFieldName))
										sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t%sDeepMap[\"%s\"] = map[string]interface{}{}\n", indent, deepTfsdkName, leafJsonName))
										sb.WriteString(fmt.Sprintf("%s\t\t\t\t}\n", indent))
									} else {
										// Handle primitive types at leaf level
										switch leafAttr.Type {
										case "string":
											sb.WriteString(fmt.Sprintf("%s\t\t\t\tif !listItem.%s.%s.IsNull() && !listItem.%s.%s.IsUnknown() {\n", indent, deepFieldName, leafFieldName, deepFieldName, leafFieldName))
											sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t%sDeepMap[\"%s\"] = listItem.%s.%s.ValueString()\n", indent, deepTfsdkName, leafJsonName, deepFieldName, leafFieldName))
											sb.WriteString(fmt.Sprintf("%s\t\t\t\t}\n", indent))
										case "int64":
											sb.WriteString(fmt.Sprintf("%s\t\t\t\tif !listItem.%s.%s.IsNull() && !listItem.%s.%s.IsUnknown() {\n", indent, deepFieldName, leafFieldName, deepFieldName, leafFieldName))
											sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t%sDeepMap[\"%s\"] = listItem.%s.%s.ValueInt64()\n", indent, deepTfsdkName, leafJsonName, deepFieldName, leafFieldName))
											sb.WriteString(fmt.Sprintf("%s\t\t\t\t}\n", indent))
										case "bool":
											sb.WriteString(fmt.Sprintf("%s\t\t\t\tif !listItem.%s.%s.IsNull() && !listItem.%s.%s.IsUnknown() {\n", indent, deepFieldName, leafFieldName, deepFieldName, leafFieldName))
											sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t%sDeepMap[\"%s\"] = listItem.%s.%s.ValueBool()\n", indent, deepTfsdkName, leafJsonName, deepFieldName, leafFieldName))
											sb.WriteString(fmt.Sprintf("%s\t\t\t\t}\n", indent))
										}
									}
								}
								sb.WriteString(fmt.Sprintf("%s\t\t\t\tlistItemMap[\"%s\"] = %sDeepMap\n", indent, deepJsonName, deepTfsdkName))
							}
							sb.WriteString(fmt.Sprintf("%s\t\t\t}\n", indent))
							continue
						}

						// Handle primitive types within list items
						switch deepAttr.Type {
						case "string":
							sb.WriteString(fmt.Sprintf("%s\t\t\tif !listItem.%s.IsNull() && !listItem.%s.IsUnknown() {\n", indent, deepFieldName, deepFieldName))
							sb.WriteString(fmt.Sprintf("%s\t\t\t\tlistItemMap[\"%s\"] = listItem.%s.ValueString()\n", indent, deepJsonName, deepFieldName))
							sb.WriteString(fmt.Sprintf("%s\t\t\t}\n", indent))
						case "int64":
							sb.WriteString(fmt.Sprintf("%s\t\t\tif !listItem.%s.IsNull() && !listItem.%s.IsUnknown() {\n", indent, deepFieldName, deepFieldName))
							sb.WriteString(fmt.Sprintf("%s\t\t\t\tlistItemMap[\"%s\"] = listItem.%s.ValueInt64()\n", indent, deepJsonName, deepFieldName))
							sb.WriteString(fmt.Sprintf("%s\t\t\t}\n", indent))
						case "bool":
							sb.WriteString(fmt.Sprintf("%s\t\t\tif !listItem.%s.IsNull() && !listItem.%s.IsUnknown() {\n", indent, deepFieldName, deepFieldName))
							sb.WriteString(fmt.Sprintf("%s\t\t\t\tlistItemMap[\"%s\"] = listItem.%s.ValueBool()\n", indent, deepJsonName, deepFieldName))
							sb.WriteString(fmt.Sprintf("%s\t\t\t}\n", indent))
						}
					}

					sb.WriteString(fmt.Sprintf("%s\t\t\t%sList = append(%sList, listItemMap)\n", indent, nestedTfsdkName, nestedTfsdkName))
					sb.WriteString(fmt.Sprintf("%s\t\t}\n", indent))
					sb.WriteString(fmt.Sprintf("%s\t\t%sMap[\"%s\"] = %sList\n", indent, tfsdkName, nestedJsonName, nestedTfsdkName))
					sb.WriteString(fmt.Sprintf("%s\t}\n", indent))
					continue
				}

				switch nestedAttr.Type {
				case "string":
					sb.WriteString(fmt.Sprintf("%s\tif !data.%s.%s.IsNull() && !data.%s.%s.IsUnknown() {\n", indent, fieldName, nestedFieldName, fieldName, nestedFieldName))
					sb.WriteString(fmt.Sprintf("%s\t\t%sMap[\"%s\"] = data.%s.%s.ValueString()\n", indent, tfsdkName, nestedTfsdkName, fieldName, nestedFieldName))
					sb.WriteString(fmt.Sprintf("%s\t}\n", indent))
				case "int64":
					sb.WriteString(fmt.Sprintf("%s\tif !data.%s.%s.IsNull() && !data.%s.%s.IsUnknown() {\n", indent, fieldName, nestedFieldName, fieldName, nestedFieldName))
					sb.WriteString(fmt.Sprintf("%s\t\t%sMap[\"%s\"] = data.%s.%s.ValueInt64()\n", indent, tfsdkName, nestedTfsdkName, fieldName, nestedFieldName))
					sb.WriteString(fmt.Sprintf("%s\t}\n", indent))
				case "bool":
					sb.WriteString(fmt.Sprintf("%s\tif !data.%s.%s.IsNull() && !data.%s.%s.IsUnknown() {\n", indent, fieldName, nestedFieldName, fieldName, nestedFieldName))
					sb.WriteString(fmt.Sprintf("%s\t\t%sMap[\"%s\"] = data.%s.%s.ValueBool()\n", indent, tfsdkName, nestedTfsdkName, fieldName, nestedFieldName))
					sb.WriteString(fmt.Sprintf("%s\t}\n", indent))
				case "list":
					// Handle list types (e.g., expressions: types.List) inside single nested blocks
					if nestedAttr.ElementType == "string" {
						sb.WriteString(fmt.Sprintf("%s\tif !data.%s.%s.IsNull() && !data.%s.%s.IsUnknown() {\n", indent, fieldName, nestedFieldName, fieldName, nestedFieldName))
						sb.WriteString(fmt.Sprintf("%s\t\tvar %sItems []string\n", indent, nestedTfsdkName))
						sb.WriteString(fmt.Sprintf("%s\t\tdiags := data.%s.%s.ElementsAs(ctx, &%sItems, false)\n", indent, fieldName, nestedFieldName, nestedTfsdkName))
						sb.WriteString(fmt.Sprintf("%s\t\tif !diags.HasError() {\n", indent))
						sb.WriteString(fmt.Sprintf("%s\t\t\t%sMap[\"%s\"] = %sItems\n", indent, tfsdkName, nestedTfsdkName, nestedTfsdkName))
						sb.WriteString(fmt.Sprintf("%s\t\t}\n", indent))
						sb.WriteString(fmt.Sprintf("%s\t}\n", indent))
					} else if nestedAttr.ElementType == "int64" {
						sb.WriteString(fmt.Sprintf("%s\tif !data.%s.%s.IsNull() && !data.%s.%s.IsUnknown() {\n", indent, fieldName, nestedFieldName, fieldName, nestedFieldName))
						sb.WriteString(fmt.Sprintf("%s\t\tvar %sItems []int64\n", indent, nestedTfsdkName))
						sb.WriteString(fmt.Sprintf("%s\t\tdiags := data.%s.%s.ElementsAs(ctx, &%sItems, false)\n", indent, fieldName, nestedFieldName, nestedTfsdkName))
						sb.WriteString(fmt.Sprintf("%s\t\tif !diags.HasError() {\n", indent))
						sb.WriteString(fmt.Sprintf("%s\t\t\t%sMap[\"%s\"] = %sItems\n", indent, tfsdkName, nestedTfsdkName, nestedTfsdkName))
						sb.WriteString(fmt.Sprintf("%s\t\t}\n", indent))
						sb.WriteString(fmt.Sprintf("%s\t}\n", indent))
					}
				}
			}
			sb.WriteString(fmt.Sprintf("%s\t%s.Spec[\"%s\"] = %sMap\n", indent, varName, jsonName, tfsdkName))
			sb.WriteString(fmt.Sprintf("%s}\n", indent))
			continue
		}

		switch attr.Type {
		case "string":
			sb.WriteString(fmt.Sprintf("%sif !data.%s.IsNull() && !data.%s.IsUnknown() {\n", indent, fieldName, fieldName))
			sb.WriteString(fmt.Sprintf("%s\t%s.Spec[\"%s\"] = data.%s.ValueString()\n", indent, varName, jsonName, fieldName))
			sb.WriteString(fmt.Sprintf("%s}\n", indent))
		case "int64":
			sb.WriteString(fmt.Sprintf("%sif !data.%s.IsNull() && !data.%s.IsUnknown() {\n", indent, fieldName, fieldName))
			sb.WriteString(fmt.Sprintf("%s\t%s.Spec[\"%s\"] = data.%s.ValueInt64()\n", indent, varName, jsonName, fieldName))
			sb.WriteString(fmt.Sprintf("%s}\n", indent))
		case "bool":
			sb.WriteString(fmt.Sprintf("%sif !data.%s.IsNull() && !data.%s.IsUnknown() {\n", indent, fieldName, fieldName))
			sb.WriteString(fmt.Sprintf("%s\t%s.Spec[\"%s\"] = data.%s.ValueBool()\n", indent, varName, jsonName, fieldName))
			sb.WriteString(fmt.Sprintf("%s}\n", indent))
		case "list":
			if attr.ElementType == "string" {
				sb.WriteString(fmt.Sprintf("%sif !data.%s.IsNull() && !data.%s.IsUnknown() {\n", indent, fieldName, fieldName))
				sb.WriteString(fmt.Sprintf("%s\tvar %sList []string\n", indent, tfsdkName))
				sb.WriteString(fmt.Sprintf("%s\tresp.Diagnostics.Append(data.%s.ElementsAs(ctx, &%sList, false)...)\n", indent, fieldName, tfsdkName))
				sb.WriteString(fmt.Sprintf("%s\tif !resp.Diagnostics.HasError() {\n", indent))
				sb.WriteString(fmt.Sprintf("%s\t\t%s.Spec[\"%s\"] = %sList\n", indent, varName, jsonName, tfsdkName))
				sb.WriteString(fmt.Sprintf("%s\t}\n", indent))
				sb.WriteString(fmt.Sprintf("%s}\n", indent))
			} else if attr.ElementType == "int64" {
				sb.WriteString(fmt.Sprintf("%sif !data.%s.IsNull() && !data.%s.IsUnknown() {\n", indent, fieldName, fieldName))
				sb.WriteString(fmt.Sprintf("%s\tvar %sList []int64\n", indent, tfsdkName))
				sb.WriteString(fmt.Sprintf("%s\tresp.Diagnostics.Append(data.%s.ElementsAs(ctx, &%sList, false)...)\n", indent, fieldName, tfsdkName))
				sb.WriteString(fmt.Sprintf("%s\tif !resp.Diagnostics.HasError() {\n", indent))
				sb.WriteString(fmt.Sprintf("%s\t\t%s.Spec[\"%s\"] = %sList\n", indent, varName, jsonName, tfsdkName))
				sb.WriteString(fmt.Sprintf("%s\t}\n", indent))
				sb.WriteString(fmt.Sprintf("%s}\n", indent))
			}
		}
	}
	return sb.String()
}

// renderComputedFieldsCode generates Go code to set Computed+Optional fields from API response
// This ensures that fields with UseStateForUnknown plan modifier have known values after Create/Update
// The varName parameter specifies the API response variable name (e.g., "created" or "updated")
func renderComputedFieldsCode(attrs []TerraformAttribute, indent string, varName string) string {
	specFields := filterSpecFields(attrs)
	if len(specFields) == 0 {
		return ""
	}

	var sb strings.Builder
	sb.WriteString(indent + "// Set computed fields from API response\n")

	for _, attr := range specFields {
		// Only generate code for Computed+Optional scalar fields (UseStateForUnknown pattern)
		if !attr.Computed || !attr.Optional || attr.IsBlock {
			continue
		}

		fieldName := attr.GoName
		jsonName := attr.JsonName
		if jsonName == "" {
			jsonName = attr.TfsdkTag
		}

		switch attr.Type {
		case "bool":
			// Set value if API returns it; otherwise handle based on plan value:
			// - If plan was unknown, set to null (resolves unknown state after apply)
			// - If plan had a value, preserve it (user specified this value)
			sb.WriteString(fmt.Sprintf("%sif v, ok := %s.Spec[\"%s\"].(bool); ok {\n", indent, varName, jsonName))
			sb.WriteString(fmt.Sprintf("%s\tdata.%s = types.BoolValue(v)\n", indent, fieldName))
			sb.WriteString(fmt.Sprintf("%s} else if data.%s.IsUnknown() {\n", indent, fieldName))
			sb.WriteString(fmt.Sprintf("%s\t// API didn't return value and plan was unknown - set to null\n", indent))
			sb.WriteString(fmt.Sprintf("%s\tdata.%s = types.BoolNull()\n", indent, fieldName))
			sb.WriteString(fmt.Sprintf("%s}\n", indent))
			sb.WriteString(fmt.Sprintf("%s// If plan had a value, preserve it\n", indent))
		case "int64":
			// Set value if API returns it; otherwise handle based on plan value
			sb.WriteString(fmt.Sprintf("%sif v, ok := %s.Spec[\"%s\"].(float64); ok {\n", indent, varName, jsonName))
			sb.WriteString(fmt.Sprintf("%s\tdata.%s = types.Int64Value(int64(v))\n", indent, fieldName))
			sb.WriteString(fmt.Sprintf("%s} else if data.%s.IsUnknown() {\n", indent, fieldName))
			sb.WriteString(fmt.Sprintf("%s\t// API didn't return value and plan was unknown - set to null\n", indent))
			sb.WriteString(fmt.Sprintf("%s\tdata.%s = types.Int64Null()\n", indent, fieldName))
			sb.WriteString(fmt.Sprintf("%s}\n", indent))
			sb.WriteString(fmt.Sprintf("%s// If plan had a value, preserve it\n", indent))
		case "string":
			// Set value if API returns it; otherwise handle based on plan value
			sb.WriteString(fmt.Sprintf("%sif v, ok := %s.Spec[\"%s\"].(string); ok && v != \"\" {\n", indent, varName, jsonName))
			sb.WriteString(fmt.Sprintf("%s\tdata.%s = types.StringValue(v)\n", indent, fieldName))
			sb.WriteString(fmt.Sprintf("%s} else if data.%s.IsUnknown() {\n", indent, fieldName))
			sb.WriteString(fmt.Sprintf("%s\t// API didn't return value and plan was unknown - set to null\n", indent))
			sb.WriteString(fmt.Sprintf("%s\tdata.%s = types.StringNull()\n", indent, fieldName))
			sb.WriteString(fmt.Sprintf("%s}\n", indent))
			sb.WriteString(fmt.Sprintf("%s// If plan had a value, preserve it\n", indent))
		}
	}

	return sb.String()
}

// renderCreateComputedFieldsCode generates code for Create method (uses "created" variable)
func renderCreateComputedFieldsCode(attrs []TerraformAttribute, indent string) string {
	return renderComputedFieldsCode(attrs, indent, "created")
}

// renderUpdateComputedFieldsCode generates code for Update method (uses "updated" variable)
// Deprecated: Use renderFetchedComputedFieldsCode instead since Update now uses GET after PUT
func renderUpdateComputedFieldsCode(attrs []TerraformAttribute, indent string) string {
	return renderComputedFieldsCode(attrs, indent, "updated")
}

// renderFetchedComputedFieldsCode generates code for Update method after GET (uses "fetched" variable)
func renderFetchedComputedFieldsCode(attrs []TerraformAttribute, indent string) string {
	return renderComputedFieldsCode(attrs, indent, "fetched")
}

// renderSpecUnmarshalCode generates Go code to unmarshal spec fields from API response to Terraform state
func renderSpecUnmarshalCode(attrs []TerraformAttribute, indent string, resourceTitleCase string) string {
	specFields := filterSpecFields(attrs)
	if len(specFields) == 0 {
		return ""
	}

	var sb strings.Builder
	for _, attr := range specFields {
		fieldName := attr.GoName
		jsonName := attr.JsonName
		if jsonName == "" {
			jsonName = attr.TfsdkTag
		}

		if attr.IsBlock {
			// Handle list nested blocks
			if attr.NestedBlockType == "list" {
				tfsdkName := attr.TfsdkTag

				// Check if there are any attributes that need the itemMap variable
				// (primitives, lists, or nested blocks with attributes)
				needsItemMap := false
				for _, nestedAttr := range attr.NestedAttributes {
					switch nestedAttr.Type {
					case "string", "int64", "bool", "list":
						needsItemMap = true
					}
					// Nested blocks with attributes need access to itemMap
					// Empty marker blocks (no nested attributes) don't need itemMap since
					// they're skipped to prevent state drift
					if nestedAttr.IsBlock && len(nestedAttr.NestedAttributes) > 0 {
						needsItemMap = true
					}
				}

				sb.WriteString(fmt.Sprintf("%sif listData, ok := apiResource.Spec[\"%s\"].([]interface{}); ok && len(listData) > 0 {\n", indent, jsonName))
				sb.WriteString(fmt.Sprintf("%s\tvar %sList []%s%sModel\n", indent, tfsdkName, resourceTitleCase, attr.GoName))
				// Extract current state's types.List to Go slice for preserving empty blocks
				sb.WriteString(fmt.Sprintf("%s\tvar existing%sItems []%s%sModel\n", indent, attr.GoName, resourceTitleCase, attr.GoName))
				sb.WriteString(fmt.Sprintf("%s\tif !data.%s.IsNull() && !data.%s.IsUnknown() {\n", indent, attr.GoName, attr.GoName))
				sb.WriteString(fmt.Sprintf("%s\t\tdata.%s.ElementsAs(ctx, &existing%sItems, false)\n", indent, attr.GoName, attr.GoName))
				sb.WriteString(fmt.Sprintf("%s\t}\n", indent))
				if needsItemMap {
					sb.WriteString(fmt.Sprintf("%s\tfor listIdx, item := range listData {\n", indent))
					sb.WriteString(fmt.Sprintf("%s\t\t_ = listIdx // May be unused if no empty marker blocks in list item\n", indent))
					sb.WriteString(fmt.Sprintf("%s\t\tif itemMap, ok := item.(map[string]interface{}); ok {\n", indent))
				} else {
					sb.WriteString(fmt.Sprintf("%s\tfor listIdx := range listData {\n", indent))
					sb.WriteString(fmt.Sprintf("%s\t\t_ = listIdx // May be unused if no empty marker blocks in list item\n", indent))
					sb.WriteString(fmt.Sprintf("%s\t\t{\n", indent))
				}
				sb.WriteString(fmt.Sprintf("%s\t\t\t%sList = append(%sList, %s%sModel{\n", indent, tfsdkName, tfsdkName, resourceTitleCase, attr.GoName))
				// Unmarshal nested block fields
				for _, nestedAttr := range attr.NestedAttributes {
					nestedFieldName := nestedAttr.GoName
					nestedJsonName := nestedAttr.JsonName
					if nestedJsonName == "" {
						nestedJsonName = nestedAttr.TfsdkTag
					}

					// Handle single nested blocks within list items
					if nestedAttr.IsBlock && nestedAttr.NestedBlockType == "single" {
						if len(nestedAttr.NestedAttributes) == 0 {
							// Empty block (marker block) - Preserve from existing state if user configured it
							// These "marker" blocks (like labels, endpoint_subsets, sni, tcp, etc.) should only
							// be in state if explicitly configured. We don't read from API as it doesn't return them.
							// Instead, we check if the user configured this block in the original state and preserve it.
							// Use extracted existingXxxItems slice instead of data.Xxx (which is now types.List)
							sb.WriteString(fmt.Sprintf("%s\t\t\t\t%s: func() *%sEmptyModel {\n", indent, nestedFieldName, resourceTitleCase))
							sb.WriteString(fmt.Sprintf("%s\t\t\t\t\tif !isImport && len(existing%sItems) > listIdx && existing%sItems[listIdx].%s != nil {\n", indent, attr.GoName, attr.GoName, nestedFieldName))
							sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\treturn &%sEmptyModel{}\n", indent, resourceTitleCase))
							sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t}\n", indent))
							sb.WriteString(fmt.Sprintf("%s\t\t\t\t\treturn nil\n", indent))
							sb.WriteString(fmt.Sprintf("%s\t\t\t\t}(),\n", indent))
						} else {
							// Block with nested attributes - unmarshal them
							// Model type includes parent list block name: Resource + ListBlock + NestedBlock + Model
							nestedModelType := resourceTitleCase + attr.GoName + nestedAttr.GoName + "Model"

							// Check if nested block has any primitive or list attributes that we can unmarshal
							hasDeepPrimitives := false
							for _, deepAttr := range nestedAttr.NestedAttributes {
								switch deepAttr.Type {
								case "string", "int64", "bool":
									hasDeepPrimitives = true
								case "list":
									// Only count lists with primitive element types that we can handle
									if deepAttr.ElementType == "string" || deepAttr.ElementType == "int64" {
										hasDeepPrimitives = true
									}
								}
							}

							sb.WriteString(fmt.Sprintf("%s\t\t\t\t%s: func() *%s {\n", indent, nestedFieldName, nestedModelType))
							if hasDeepPrimitives {
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\tif nestedMap, ok := itemMap[\"%s\"].(map[string]interface{}); ok {\n", indent, nestedJsonName))
							} else {
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\tif _, ok := itemMap[\"%s\"].(map[string]interface{}); ok {\n", indent, nestedJsonName))
							}
							sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\treturn &%s{\n", indent, nestedModelType))
							for _, deepAttr := range nestedAttr.NestedAttributes {
								deepFieldName := deepAttr.GoName
								deepJsonName := deepAttr.JsonName
								if deepJsonName == "" {
									deepJsonName = deepAttr.TfsdkTag
								}
								// Handle deeply nested empty blocks (e.g., allow{}, deny{} within action{} within rules[])
								// These are marker blocks that don't come back from API, so preserve from existing state
								// Use extracted existingXxxItems slice instead of data.Xxx (which is now types.List)
								if deepAttr.IsBlock && len(deepAttr.NestedAttributes) == 0 {
									sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t%s: func() *%sEmptyModel {\n", indent, deepFieldName, resourceTitleCase))
									sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\tif !isImport && len(existing%sItems) > listIdx && existing%sItems[listIdx].%s != nil && existing%sItems[listIdx].%s.%s != nil {\n", indent, attr.GoName, attr.GoName, nestedAttr.GoName, attr.GoName, nestedAttr.GoName, deepFieldName))
									sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\treturn &%sEmptyModel{}\n", indent, resourceTitleCase))
									sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t}\n", indent))
									sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\treturn nil\n", indent))
									sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t}(),\n", indent))
									continue
								}
								switch deepAttr.Type {
								case "string":
									sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t%s: func() types.String {\n", indent, deepFieldName))
									sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\tif v, ok := nestedMap[\"%s\"].(string); ok && v != \"\" {\n", indent, deepJsonName))
									sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\treturn types.StringValue(v)\n", indent))
									sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t}\n", indent))
									sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\treturn types.StringNull()\n", indent))
									sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t}(),\n", indent))
								case "int64":
									// Only set non-zero values to avoid drift from API defaults
									// Zero typically means "not set" for optional fields like refresh_interval
									sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t%s: func() types.Int64 {\n", indent, deepFieldName))
									sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\tif v, ok := nestedMap[\"%s\"].(float64); ok && v != 0 {\n", indent, deepJsonName))
									sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\treturn types.Int64Value(int64(v))\n", indent))
									sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t}\n", indent))
									sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\treturn types.Int64Null()\n", indent))
									sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t}(),\n", indent))
								case "bool":
									sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t%s: func() types.Bool {\n", indent, deepFieldName))
									sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\tif v, ok := nestedMap[\"%s\"].(bool); ok {\n", indent, deepJsonName))
									sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\treturn types.BoolValue(v)\n", indent))
									sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t}\n", indent))
									sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\treturn types.BoolNull()\n", indent))
									sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t}(),\n", indent))
								case "list":
									// Handle list types inside single nested blocks within list items
									if deepAttr.ElementType == "string" {
										sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t%s: func() types.List {\n", indent, deepFieldName))
										sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\tif v, ok := nestedMap[\"%s\"].([]interface{}); ok && len(v) > 0 {\n", indent, deepJsonName))
										sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\tvar items []string\n", indent))
										sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\tfor _, item := range v {\n", indent))
										sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\t\tif s, ok := item.(string); ok {\n", indent))
										sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\t\t\titems = append(items, s)\n", indent))
										sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\t\t}\n", indent))
										sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\t}\n", indent))
										sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\tlistVal, _ := types.ListValueFrom(ctx, types.StringType, items)\n", indent))
										sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\treturn listVal\n", indent))
										sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t}\n", indent))
										sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\treturn types.ListNull(types.StringType)\n", indent))
										sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t}(),\n", indent))
									} else if deepAttr.ElementType == "int64" {
										sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t%s: func() types.List {\n", indent, deepFieldName))
										sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\tif v, ok := nestedMap[\"%s\"].([]interface{}); ok && len(v) > 0 {\n", indent, deepJsonName))
										sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\tvar items []int64\n", indent))
										sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\tfor _, item := range v {\n", indent))
										sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\t\tif n, ok := item.(float64); ok {\n", indent))
										sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\t\t\titems = append(items, int64(n))\n", indent))
										sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\t\t}\n", indent))
										sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\t}\n", indent))
										sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\tlistVal, _ := types.ListValueFrom(ctx, types.Int64Type, items)\n", indent))
										sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\treturn listVal\n", indent))
										sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t}\n", indent))
										sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\treturn types.ListNull(types.Int64Type)\n", indent))
										sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t}(),\n", indent))
									}
								}
							}
							sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t}\n", indent))
							sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t}\n", indent))
							sb.WriteString(fmt.Sprintf("%s\t\t\t\t\treturn nil\n", indent))
							sb.WriteString(fmt.Sprintf("%s\t\t\t\t}(),\n", indent))
						}
						continue
					}

					// Handle list nested blocks within list items (e.g., app_type_ref within app_type_settings)
					if nestedAttr.IsBlock && nestedAttr.NestedBlockType == "list" {
						// Model type includes parent list block name: Resource + ListBlock + NestedListBlock + Model
						nestedListModelType := resourceTitleCase + attr.GoName + nestedAttr.GoName + "Model"

						sb.WriteString(fmt.Sprintf("%s\t\t\t\t%s: func() []%s {\n", indent, nestedFieldName, nestedListModelType))
						sb.WriteString(fmt.Sprintf("%s\t\t\t\t\tif nestedListData, ok := itemMap[\"%s\"].([]interface{}); ok && len(nestedListData) > 0 {\n", indent, nestedJsonName))
						sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\tvar result []%s\n", indent, nestedListModelType))
						sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\tfor _, nestedItem := range nestedListData {\n", indent))
						sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\tif nestedItemMap, ok := nestedItem.(map[string]interface{}); ok {\n", indent))
						sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\tresult = append(result, %s{\n", indent, nestedListModelType))
						// Generate field reading for each nested attribute
						for _, deepAttr := range nestedAttr.NestedAttributes {
							deepFieldName := deepAttr.GoName
							deepJsonName := deepAttr.JsonName
							if deepJsonName == "" {
								deepJsonName = deepAttr.TfsdkTag
							}
							switch deepAttr.Type {
							case "string":
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\t%s: func() types.String {\n", indent, deepFieldName))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\t\tif v, ok := nestedItemMap[\"%s\"].(string); ok && v != \"\" {\n", indent, deepJsonName))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\t\t\treturn types.StringValue(v)\n", indent))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\t\t}\n", indent))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\t\treturn types.StringNull()\n", indent))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\t}(),\n", indent))
							case "int64":
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\t%s: func() types.Int64 {\n", indent, deepFieldName))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\t\tif v, ok := nestedItemMap[\"%s\"].(float64); ok && v != 0 {\n", indent, deepJsonName))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\t\t\treturn types.Int64Value(int64(v))\n", indent))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\t\t}\n", indent))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\t\treturn types.Int64Null()\n", indent))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\t}(),\n", indent))
							case "bool":
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\t%s: func() types.Bool {\n", indent, deepFieldName))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\t\tif v, ok := nestedItemMap[\"%s\"].(bool); ok {\n", indent, deepJsonName))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\t\t\treturn types.BoolValue(v)\n", indent))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\t\t}\n", indent))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\t\treturn types.BoolNull()\n", indent))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\t}(),\n", indent))
							}
						}
						sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t})\n", indent))
						sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t}\n", indent))
						sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t}\n", indent))
						sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\treturn result\n", indent))
						sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t}\n", indent))
						sb.WriteString(fmt.Sprintf("%s\t\t\t\t\treturn nil\n", indent))
						sb.WriteString(fmt.Sprintf("%s\t\t\t\t}(),\n", indent))
						continue
					}

					switch nestedAttr.Type {
					case "string":
						sb.WriteString(fmt.Sprintf("%s\t\t\t\t%s: func() types.String {\n", indent, nestedFieldName))
						sb.WriteString(fmt.Sprintf("%s\t\t\t\t\tif v, ok := itemMap[\"%s\"].(string); ok && v != \"\" {\n", indent, nestedJsonName))
						sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\treturn types.StringValue(v)\n", indent))
						sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t}\n", indent))
						sb.WriteString(fmt.Sprintf("%s\t\t\t\t\treturn types.StringNull()\n", indent))
						sb.WriteString(fmt.Sprintf("%s\t\t\t\t}(),\n", indent))
					case "int64":
						// Only set non-zero values to avoid drift from API defaults
						sb.WriteString(fmt.Sprintf("%s\t\t\t\t%s: func() types.Int64 {\n", indent, nestedFieldName))
						sb.WriteString(fmt.Sprintf("%s\t\t\t\t\tif v, ok := itemMap[\"%s\"].(float64); ok && v != 0 {\n", indent, nestedJsonName))
						sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\treturn types.Int64Value(int64(v))\n", indent))
						sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t}\n", indent))
						sb.WriteString(fmt.Sprintf("%s\t\t\t\t\treturn types.Int64Null()\n", indent))
						sb.WriteString(fmt.Sprintf("%s\t\t\t\t}(),\n", indent))
					case "bool":
						sb.WriteString(fmt.Sprintf("%s\t\t\t\t%s: func() types.Bool {\n", indent, nestedFieldName))
						sb.WriteString(fmt.Sprintf("%s\t\t\t\t\tif v, ok := itemMap[\"%s\"].(bool); ok {\n", indent, nestedJsonName))
						sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\treturn types.BoolValue(v)\n", indent))
						sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t}\n", indent))
						sb.WriteString(fmt.Sprintf("%s\t\t\t\t\treturn types.BoolNull()\n", indent))
						sb.WriteString(fmt.Sprintf("%s\t\t\t\t}(),\n", indent))
					case "list":
						// Handle list types inside list nested blocks
						if nestedAttr.ElementType == "string" {
							sb.WriteString(fmt.Sprintf("%s\t\t\t\t%s: func() types.List {\n", indent, nestedFieldName))
							sb.WriteString(fmt.Sprintf("%s\t\t\t\t\tif v, ok := itemMap[\"%s\"].([]interface{}); ok && len(v) > 0 {\n", indent, nestedJsonName))
							sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\tvar items []string\n", indent))
							sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\tfor _, item := range v {\n", indent))
							sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\tif s, ok := item.(string); ok {\n", indent))
							sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\titems = append(items, s)\n", indent))
							sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t}\n", indent))
							sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t}\n", indent))
							sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\tlistVal, _ := types.ListValueFrom(ctx, types.StringType, items)\n", indent))
							sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\treturn listVal\n", indent))
							sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t}\n", indent))
							sb.WriteString(fmt.Sprintf("%s\t\t\t\t\treturn types.ListNull(types.StringType)\n", indent))
							sb.WriteString(fmt.Sprintf("%s\t\t\t\t}(),\n", indent))
						} else if nestedAttr.ElementType == "int64" {
							sb.WriteString(fmt.Sprintf("%s\t\t\t\t%s: func() types.List {\n", indent, nestedFieldName))
							sb.WriteString(fmt.Sprintf("%s\t\t\t\t\tif v, ok := itemMap[\"%s\"].([]interface{}); ok && len(v) > 0 {\n", indent, nestedJsonName))
							sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\tvar items []int64\n", indent))
							sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\tfor _, item := range v {\n", indent))
							sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\tif n, ok := item.(float64); ok {\n", indent))
							sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\titems = append(items, int64(n))\n", indent))
							sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t}\n", indent))
							sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t}\n", indent))
							sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\tlistVal, _ := types.ListValueFrom(ctx, types.Int64Type, items)\n", indent))
							sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\treturn listVal\n", indent))
							sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t}\n", indent))
							sb.WriteString(fmt.Sprintf("%s\t\t\t\t\treturn types.ListNull(types.Int64Type)\n", indent))
							sb.WriteString(fmt.Sprintf("%s\t\t\t\t}(),\n", indent))
						}
					}
				}
				sb.WriteString(fmt.Sprintf("%s\t\t\t})\n", indent))
				sb.WriteString(fmt.Sprintf("%s\t\t}\n", indent))
				sb.WriteString(fmt.Sprintf("%s\t}\n", indent))
				// Convert Go slice to types.List for proper unknown value handling
				sb.WriteString(fmt.Sprintf("%s\tlistVal, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: %s%sModelAttrTypes}, %sList)\n", indent, resourceTitleCase, attr.GoName, tfsdkName))
				sb.WriteString(fmt.Sprintf("%s\tresp.Diagnostics.Append(diags...)\n", indent))
				sb.WriteString(fmt.Sprintf("%s\tif !resp.Diagnostics.HasError() {\n", indent))
				sb.WriteString(fmt.Sprintf("%s\t\tdata.%s = listVal\n", indent, fieldName))
				sb.WriteString(fmt.Sprintf("%s\t}\n", indent))
				sb.WriteString(fmt.Sprintf("%s} else {\n", indent))
				sb.WriteString(fmt.Sprintf("%s\t// No data from API - set to null list\n", indent))
				sb.WriteString(fmt.Sprintf("%s\tdata.%s = types.ListNull(types.ObjectType{AttrTypes: %s%sModelAttrTypes})\n", indent, fieldName, resourceTitleCase, attr.GoName))
				sb.WriteString(fmt.Sprintf("%s}\n", indent))
				continue
			}

			// Empty blocks (no nested attributes) use EmptyModel
			// These are presence markers (choice blocks), so:
			// - During Import (psd is nil): read from API to populate state
			// - During normal Read (psd exists): preserve state to avoid drift
			if len(attr.NestedAttributes) == 0 {
				sb.WriteString(fmt.Sprintf("%sif _, ok := apiResource.Spec[\"%s\"].(map[string]interface{}); ok && isImport && data.%s == nil {\n", indent, jsonName, fieldName))
				sb.WriteString(fmt.Sprintf("%s\t// Import case: populate from API since state is nil and psd is empty\n", indent))
				sb.WriteString(fmt.Sprintf("%s\tdata.%s = &%sEmptyModel{}\n", indent, fieldName, resourceTitleCase))
				sb.WriteString(fmt.Sprintf("%s}\n", indent))
				sb.WriteString(fmt.Sprintf("%s// Normal Read: preserve existing state value\n", indent))
				continue
			}

			// Check if we need to read data (for primitives, lists, or nested list blocks)
			// Nested empty blocks (choice blocks) are NOT read back as they're presence markers and can cause drift
			hasReadableData := false
			for _, nestedAttr := range attr.NestedAttributes {
				switch nestedAttr.Type {
				case "string", "int64", "bool", "list":
					hasReadableData = true
				}
				// Check for nested list blocks that need to be read
				if nestedAttr.IsBlock && nestedAttr.NestedBlockType == "list" {
					hasReadableData = true
				}
			}

			// If the block contains ONLY nested empty blocks (no primitives, lists, or nested list blocks), handle specially:
			// - During Import (psd is nil): read from API to populate state
			// - During normal Read (psd exists): preserve state to avoid drift
			if !hasReadableData {
				sb.WriteString(fmt.Sprintf("%sif _, ok := apiResource.Spec[\"%s\"].(map[string]interface{}); ok && isImport && data.%s == nil {\n", indent, jsonName, fieldName))
				sb.WriteString(fmt.Sprintf("%s\t// Import case: populate from API since state is nil and psd is empty\n", indent))
				sb.WriteString(fmt.Sprintf("%s\tdata.%s = &%s%sModel{}\n", indent, fieldName, resourceTitleCase, attr.GoName))
				sb.WriteString(fmt.Sprintf("%s}\n", indent))
				sb.WriteString(fmt.Sprintf("%s// Normal Read: preserve existing state value\n", indent))
				continue
			}
			// Unmarshal single nested block from API response map
			// Only populate if this is an import (need full state) or user configured this block
			sb.WriteString(fmt.Sprintf("%sif blockData, ok := apiResource.Spec[\"%s\"].(map[string]interface{}); ok && (isImport || data.%s != nil) {\n", indent, jsonName, fieldName))
			sb.WriteString(fmt.Sprintf("%s\tdata.%s = &%s%sModel{\n", indent, fieldName, resourceTitleCase, attr.GoName))
			// Unmarshal nested block fields
			for _, nestedAttr := range attr.NestedAttributes {
				nestedFieldName := nestedAttr.GoName
				nestedJsonName := nestedAttr.JsonName
				if nestedJsonName == "" {
					nestedJsonName = nestedAttr.TfsdkTag
				}

				// Handle empty single nested blocks (presence markers like default_action_deny {})
				// These should be preserved from prior state during normal Read, and populated on Import
				if nestedAttr.IsBlock && nestedAttr.NestedBlockType == "single" && len(nestedAttr.NestedAttributes) == 0 {
					sb.WriteString(fmt.Sprintf("%s\t\t%s: func() *%sEmptyModel {\n", indent, nestedFieldName, resourceTitleCase))
					sb.WriteString(fmt.Sprintf("%s\t\t\tif !isImport && data.%s != nil {\n", indent, fieldName))
					sb.WriteString(fmt.Sprintf("%s\t\t\t\t// Normal Read: preserve existing state value (even if nil)\n", indent))
					sb.WriteString(fmt.Sprintf("%s\t\t\t\t// This prevents API returning empty objects from overwriting user's 'not configured' intent\n", indent))
					sb.WriteString(fmt.Sprintf("%s\t\t\t\treturn data.%s.%s\n", indent, fieldName, nestedFieldName))
					sb.WriteString(fmt.Sprintf("%s\t\t\t}\n", indent))
					sb.WriteString(fmt.Sprintf("%s\t\t\t// Import case: read from API\n", indent))
					sb.WriteString(fmt.Sprintf("%s\t\t\tif _, ok := blockData[\"%s\"].(map[string]interface{}); ok {\n", indent, nestedJsonName))
					sb.WriteString(fmt.Sprintf("%s\t\t\t\treturn &%sEmptyModel{}\n", indent, resourceTitleCase))
					sb.WriteString(fmt.Sprintf("%s\t\t\t}\n", indent))
					sb.WriteString(fmt.Sprintf("%s\t\t\treturn nil\n", indent))
					sb.WriteString(fmt.Sprintf("%s\t\t}(),\n", indent))
					continue
				}

				switch nestedAttr.Type {
				case "string":
					sb.WriteString(fmt.Sprintf("%s\t\t%s: func() types.String {\n", indent, nestedFieldName))
					sb.WriteString(fmt.Sprintf("%s\t\t\tif v, ok := blockData[\"%s\"].(string); ok && v != \"\" {\n", indent, nestedJsonName))
					sb.WriteString(fmt.Sprintf("%s\t\t\t\treturn types.StringValue(v)\n", indent))
					sb.WriteString(fmt.Sprintf("%s\t\t\t}\n", indent))
					sb.WriteString(fmt.Sprintf("%s\t\t\treturn types.StringNull()\n", indent))
					sb.WriteString(fmt.Sprintf("%s\t\t}(),\n", indent))
				case "int64":
					sb.WriteString(fmt.Sprintf("%s\t\t%s: func() types.Int64 {\n", indent, nestedFieldName))
					// For Optional nested int64 fields, preserve prior state during normal Read to avoid drift
					// This handles fields like connection_idle_timeout where API returns 0 as default
					if nestedAttr.Optional {
						sb.WriteString(fmt.Sprintf("%s\t\t\tif !isImport && data.%s != nil {\n", indent, fieldName))
						sb.WriteString(fmt.Sprintf("%s\t\t\t\t// Preserve existing state (null or user-set value)\n", indent))
						sb.WriteString(fmt.Sprintf("%s\t\t\t\t// This prevents API defaults (like 0) from overwriting user intent\n", indent))
						sb.WriteString(fmt.Sprintf("%s\t\t\t\treturn data.%s.%s\n", indent, fieldName, nestedFieldName))
						sb.WriteString(fmt.Sprintf("%s\t\t\t}\n", indent))
						// Fallback: block not in user config but not importing - return null, not API default
						sb.WriteString(fmt.Sprintf("%s\t\t\tif !isImport {\n", indent))
						sb.WriteString(fmt.Sprintf("%s\t\t\t\t// Block not in user config - return null, not API default\n", indent))
						sb.WriteString(fmt.Sprintf("%s\t\t\t\treturn types.Int64Null()\n", indent))
						sb.WriteString(fmt.Sprintf("%s\t\t\t}\n", indent))
					}
					sb.WriteString(fmt.Sprintf("%s\t\t\t// Import case: read from API\n", indent))
					sb.WriteString(fmt.Sprintf("%s\t\t\tif v, ok := blockData[\"%s\"].(float64); ok {\n", indent, nestedJsonName))
					sb.WriteString(fmt.Sprintf("%s\t\t\t\treturn types.Int64Value(int64(v))\n", indent))
					sb.WriteString(fmt.Sprintf("%s\t\t\t}\n", indent))
					sb.WriteString(fmt.Sprintf("%s\t\t\treturn types.Int64Null()\n", indent))
					sb.WriteString(fmt.Sprintf("%s\t\t}(),\n", indent))
				case "bool":
					sb.WriteString(fmt.Sprintf("%s\t\t%s: func() types.Bool {\n", indent, nestedFieldName))
					// For Optional nested bool fields, preserve prior state during normal Read to avoid drift
					// This handles both Computed+Optional and Optional-only fields where API returns defaults
					if nestedAttr.Optional {
						sb.WriteString(fmt.Sprintf("%s\t\t\tif !isImport && data.%s != nil {\n", indent, fieldName))
						sb.WriteString(fmt.Sprintf("%s\t\t\t\t// Preserve existing state (null or user-set value)\n", indent))
						sb.WriteString(fmt.Sprintf("%s\t\t\t\t// This prevents API defaults from overwriting user intent\n", indent))
						sb.WriteString(fmt.Sprintf("%s\t\t\t\treturn data.%s.%s\n", indent, fieldName, nestedFieldName))
						sb.WriteString(fmt.Sprintf("%s\t\t\t}\n", indent))
						// Fallback: block not in user config but not importing - return null, not API default
						sb.WriteString(fmt.Sprintf("%s\t\t\tif !isImport {\n", indent))
						sb.WriteString(fmt.Sprintf("%s\t\t\t\t// Block not in user config - return null, not API default\n", indent))
						sb.WriteString(fmt.Sprintf("%s\t\t\t\treturn types.BoolNull()\n", indent))
						sb.WriteString(fmt.Sprintf("%s\t\t\t}\n", indent))
					}
					sb.WriteString(fmt.Sprintf("%s\t\t\t// Import case: read from API\n", indent))
					sb.WriteString(fmt.Sprintf("%s\t\t\tif v, ok := blockData[\"%s\"].(bool); ok {\n", indent, nestedJsonName))
					sb.WriteString(fmt.Sprintf("%s\t\t\t\treturn types.BoolValue(v)\n", indent))
					sb.WriteString(fmt.Sprintf("%s\t\t\t}\n", indent))
					sb.WriteString(fmt.Sprintf("%s\t\t\treturn types.BoolNull()\n", indent))
					sb.WriteString(fmt.Sprintf("%s\t\t}(),\n", indent))
				case "list":
					// Handle list types inside single nested blocks
					if nestedAttr.ElementType == "string" {
						sb.WriteString(fmt.Sprintf("%s\t\t%s: func() types.List {\n", indent, nestedFieldName))
						sb.WriteString(fmt.Sprintf("%s\t\t\tif v, ok := blockData[\"%s\"].([]interface{}); ok && len(v) > 0 {\n", indent, nestedJsonName))
						sb.WriteString(fmt.Sprintf("%s\t\t\t\tvar items []string\n", indent))
						sb.WriteString(fmt.Sprintf("%s\t\t\t\tfor _, item := range v {\n", indent))
						sb.WriteString(fmt.Sprintf("%s\t\t\t\t\tif s, ok := item.(string); ok {\n", indent))
						sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\titems = append(items, s)\n", indent))
						sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t}\n", indent))
						sb.WriteString(fmt.Sprintf("%s\t\t\t\t}\n", indent))
						sb.WriteString(fmt.Sprintf("%s\t\t\t\tlistVal, _ := types.ListValueFrom(ctx, types.StringType, items)\n", indent))
						sb.WriteString(fmt.Sprintf("%s\t\t\t\treturn listVal\n", indent))
						sb.WriteString(fmt.Sprintf("%s\t\t\t}\n", indent))
						sb.WriteString(fmt.Sprintf("%s\t\t\treturn types.ListNull(types.StringType)\n", indent))
						sb.WriteString(fmt.Sprintf("%s\t\t}(),\n", indent))
					} else if nestedAttr.ElementType == "int64" {
						sb.WriteString(fmt.Sprintf("%s\t\t%s: func() types.List {\n", indent, nestedFieldName))
						sb.WriteString(fmt.Sprintf("%s\t\t\tif v, ok := blockData[\"%s\"].([]interface{}); ok && len(v) > 0 {\n", indent, nestedJsonName))
						sb.WriteString(fmt.Sprintf("%s\t\t\t\tvar items []int64\n", indent))
						sb.WriteString(fmt.Sprintf("%s\t\t\t\tfor _, item := range v {\n", indent))
						sb.WriteString(fmt.Sprintf("%s\t\t\t\t\tif n, ok := item.(float64); ok {\n", indent))
						sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\titems = append(items, int64(n))\n", indent))
						sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t}\n", indent))
						sb.WriteString(fmt.Sprintf("%s\t\t\t\t}\n", indent))
						sb.WriteString(fmt.Sprintf("%s\t\t\t\tlistVal, _ := types.ListValueFrom(ctx, types.Int64Type, items)\n", indent))
						sb.WriteString(fmt.Sprintf("%s\t\t\t\treturn listVal\n", indent))
						sb.WriteString(fmt.Sprintf("%s\t\t\t}\n", indent))
						sb.WriteString(fmt.Sprintf("%s\t\t\treturn types.ListNull(types.Int64Type)\n", indent))
						sb.WriteString(fmt.Sprintf("%s\t\t}(),\n", indent))
					}
				}
				// Handle nested list blocks within single nested blocks (e.g., rules in mitigation_type)
				if nestedAttr.IsBlock && nestedAttr.NestedBlockType == "list" {
					// Model type for nested list within single block: Resource + ParentBlock + NestedList + Model
					nestedListModelType := resourceTitleCase + attr.GoName + nestedAttr.GoName + "Model"

					sb.WriteString(fmt.Sprintf("%s\t\t%s: func() []%s {\n", indent, nestedFieldName, nestedListModelType))
					sb.WriteString(fmt.Sprintf("%s\t\t\tif listData, ok := blockData[\"%s\"].([]interface{}); ok && len(listData) > 0 {\n", indent, nestedJsonName))
					sb.WriteString(fmt.Sprintf("%s\t\t\t\tvar result []%s\n", indent, nestedListModelType))
					sb.WriteString(fmt.Sprintf("%s\t\t\t\tfor _, item := range listData {\n", indent))
					sb.WriteString(fmt.Sprintf("%s\t\t\t\t\tif itemMap, ok := item.(map[string]interface{}); ok {\n", indent))
					sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\tresult = append(result, %s{\n", indent, nestedListModelType))

					// Process nested attributes within the list items
					for _, deepAttr := range nestedAttr.NestedAttributes {
						deepFieldName := deepAttr.GoName
						deepJsonName := deepAttr.JsonName
						if deepJsonName == "" {
							deepJsonName = deepAttr.TfsdkTag
						}

						// Handle single nested blocks within list items (like threat_level, mitigation_action)
						if deepAttr.IsBlock && deepAttr.NestedBlockType == "single" {
							// Model type: Resource + ParentBlock + ListBlock + DeepBlock + Model
							deepModelType := resourceTitleCase + attr.GoName + nestedAttr.GoName + deepAttr.GoName + "Model"

							if len(deepAttr.NestedAttributes) == 0 {
								// Empty marker block - check if key exists
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t%s: func() *%sEmptyModel {\n", indent, deepFieldName, resourceTitleCase))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\tif _, ok := itemMap[\"%s\"].(map[string]interface{}); ok {\n", indent, deepJsonName))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\treturn &%sEmptyModel{}\n", indent, resourceTitleCase))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t}\n", indent))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\treturn nil\n", indent))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t}(),\n", indent))
							} else {
								// Block with nested attributes
								// Check if any leaf attributes will actually use deepMap
								needsDeepMap := false
								for _, checkAttr := range deepAttr.NestedAttributes {
									// Empty single blocks and primitives use deepMap
									if checkAttr.IsBlock && checkAttr.NestedBlockType == "single" && len(checkAttr.NestedAttributes) == 0 {
										needsDeepMap = true
										break
									}
									switch checkAttr.Type {
									case "string", "int64", "bool":
										needsDeepMap = true
									}
									if needsDeepMap {
										break
									}
								}

								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t%s: func() *%s {\n", indent, deepFieldName, deepModelType))
								if needsDeepMap {
									sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\tif deepMap, ok := itemMap[\"%s\"].(map[string]interface{}); ok {\n", indent, deepJsonName))
								} else {
									sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\tif _, ok := itemMap[\"%s\"].(map[string]interface{}); ok {\n", indent, deepJsonName))
								}
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\treturn &%s{\n", indent, deepModelType))

								// Handle the deepest level attributes (choice blocks like high, medium, low)
								for _, leafAttr := range deepAttr.NestedAttributes {
									leafFieldName := leafAttr.GoName
									leafJsonName := leafAttr.JsonName
									if leafJsonName == "" {
										leafJsonName = leafAttr.TfsdkTag
									}

									if leafAttr.IsBlock && leafAttr.NestedBlockType == "single" && len(leafAttr.NestedAttributes) == 0 {
										// Empty marker/choice block
										sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\t\t%s: func() *%sEmptyModel {\n", indent, leafFieldName, resourceTitleCase))
										sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\t\t\tif _, ok := deepMap[\"%s\"].(map[string]interface{}); ok {\n", indent, leafJsonName))
										sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\t\t\t\treturn &%sEmptyModel{}\n", indent, resourceTitleCase))
										sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\t\t\t}\n", indent))
										sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\t\t\treturn nil\n", indent))
										sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\t\t}(),\n", indent))
									} else {
										// Primitive types at the leaf level
										switch leafAttr.Type {
										case "string":
											sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\t\t%s: func() types.String {\n", indent, leafFieldName))
											sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\t\t\tif v, ok := deepMap[\"%s\"].(string); ok && v != \"\" {\n", indent, leafJsonName))
											sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\t\t\t\treturn types.StringValue(v)\n", indent))
											sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\t\t\t}\n", indent))
											sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\t\t\treturn types.StringNull()\n", indent))
											sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\t\t}(),\n", indent))
										case "int64":
											sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\t\t%s: func() types.Int64 {\n", indent, leafFieldName))
											sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\t\t\tif v, ok := deepMap[\"%s\"].(float64); ok {\n", indent, leafJsonName))
											sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\t\t\t\treturn types.Int64Value(int64(v))\n", indent))
											sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\t\t\t}\n", indent))
											sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\t\t\treturn types.Int64Null()\n", indent))
											sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\t\t}(),\n", indent))
										case "bool":
											sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\t\t%s: func() types.Bool {\n", indent, leafFieldName))
											sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\t\t\tif v, ok := deepMap[\"%s\"].(bool); ok {\n", indent, leafJsonName))
											sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\t\t\t\treturn types.BoolValue(v)\n", indent))
											sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\t\t\t}\n", indent))
											sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\t\t\treturn types.BoolNull()\n", indent))
											sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\t\t}(),\n", indent))
										}
									}
								}

								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\t}\n", indent))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t}\n", indent))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\treturn nil\n", indent))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t}(),\n", indent))
							}
						} else {
							// Handle primitive types within list items
							switch deepAttr.Type {
							case "string":
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t%s: func() types.String {\n", indent, deepFieldName))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\tif v, ok := itemMap[\"%s\"].(string); ok && v != \"\" {\n", indent, deepJsonName))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\treturn types.StringValue(v)\n", indent))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t}\n", indent))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\treturn types.StringNull()\n", indent))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t}(),\n", indent))
							case "int64":
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t%s: func() types.Int64 {\n", indent, deepFieldName))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\tif v, ok := itemMap[\"%s\"].(float64); ok {\n", indent, deepJsonName))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\treturn types.Int64Value(int64(v))\n", indent))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t}\n", indent))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\treturn types.Int64Null()\n", indent))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t}(),\n", indent))
							case "bool":
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t%s: func() types.Bool {\n", indent, deepFieldName))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\tif v, ok := itemMap[\"%s\"].(bool); ok {\n", indent, deepJsonName))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\treturn types.BoolValue(v)\n", indent))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t}\n", indent))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\treturn types.BoolNull()\n", indent))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t}(),\n", indent))
							}
						}
					}

					sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t})\n", indent))
					sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t}\n", indent))
					sb.WriteString(fmt.Sprintf("%s\t\t\t\t}\n", indent))
					sb.WriteString(fmt.Sprintf("%s\t\t\t\treturn result\n", indent))
					sb.WriteString(fmt.Sprintf("%s\t\t\t}\n", indent))
					sb.WriteString(fmt.Sprintf("%s\t\t\treturn nil\n", indent))
					sb.WriteString(fmt.Sprintf("%s\t\t}(),\n", indent))
				}
				// Handle single nested blocks with attributes within single nested blocks (e.g., prefix_list with prefixes)
				if nestedAttr.IsBlock && nestedAttr.NestedBlockType == "single" && len(nestedAttr.NestedAttributes) > 0 {
					// Model type for nested single within single block: Resource + ParentBlock + NestedSingle + Model
					nestedSingleModelType := resourceTitleCase + attr.GoName + nestedAttr.GoName + "Model"

					// Check if any attributes will actually use nestedBlockData
					hasUsableAttrs := false
					for _, da := range nestedAttr.NestedAttributes {
						switch da.Type {
						case "string", "int64", "bool":
							hasUsableAttrs = true
						case "list":
							if da.ElementType == "string" || da.ElementType == "int64" {
								hasUsableAttrs = true
							}
						}
					}

					sb.WriteString(fmt.Sprintf("%s\t\t%s: func() *%s {\n", indent, nestedFieldName, nestedSingleModelType))
					sb.WriteString(fmt.Sprintf("%s\t\t\tif !isImport && data.%s != nil && data.%s.%s != nil {\n", indent, fieldName, fieldName, nestedFieldName))
					sb.WriteString(fmt.Sprintf("%s\t\t\t\t// Normal Read: preserve existing state value\n", indent))
					sb.WriteString(fmt.Sprintf("%s\t\t\t\treturn data.%s.%s\n", indent, fieldName, nestedFieldName))
					sb.WriteString(fmt.Sprintf("%s\t\t\t}\n", indent))
					sb.WriteString(fmt.Sprintf("%s\t\t\t// Import case: read from API\n", indent))

					// Use _ if no attributes will use nestedBlockData
					varName := "nestedBlockData"
					if !hasUsableAttrs {
						varName = "_"
					}
					sb.WriteString(fmt.Sprintf("%s\t\t\tif %s, ok := blockData[\"%s\"].(map[string]interface{}); ok {\n", indent, varName, nestedJsonName))
					sb.WriteString(fmt.Sprintf("%s\t\t\t\treturn &%s{\n", indent, nestedSingleModelType))

					// Process attributes within the single nested block
					for _, deepAttr := range nestedAttr.NestedAttributes {
						deepFieldName := deepAttr.GoName
						deepJsonName := deepAttr.JsonName
						if deepJsonName == "" {
							deepJsonName = deepAttr.TfsdkTag
						}

						switch deepAttr.Type {
						case "string":
							sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t%s: func() types.String {\n", indent, deepFieldName))
							sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\tif v, ok := nestedBlockData[\"%s\"].(string); ok && v != \"\" {\n", indent, deepJsonName))
							sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\treturn types.StringValue(v)\n", indent))
							sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t}\n", indent))
							sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\treturn types.StringNull()\n", indent))
							sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t}(),\n", indent))
						case "int64":
							sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t%s: func() types.Int64 {\n", indent, deepFieldName))
							sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\tif v, ok := nestedBlockData[\"%s\"].(float64); ok {\n", indent, deepJsonName))
							sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\treturn types.Int64Value(int64(v))\n", indent))
							sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t}\n", indent))
							sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\treturn types.Int64Null()\n", indent))
							sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t}(),\n", indent))
						case "bool":
							sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t%s: func() types.Bool {\n", indent, deepFieldName))
							sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\tif v, ok := nestedBlockData[\"%s\"].(bool); ok {\n", indent, deepJsonName))
							sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\treturn types.BoolValue(v)\n", indent))
							sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t}\n", indent))
							sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\treturn types.BoolNull()\n", indent))
							sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t}(),\n", indent))
						case "list":
							if deepAttr.ElementType == "string" {
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t%s: func() types.List {\n", indent, deepFieldName))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\tif v, ok := nestedBlockData[\"%s\"].([]interface{}); ok && len(v) > 0 {\n", indent, deepJsonName))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\tvar items []string\n", indent))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\tfor _, item := range v {\n", indent))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\tif s, ok := item.(string); ok {\n", indent))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\titems = append(items, s)\n", indent))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t}\n", indent))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t}\n", indent))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\tlistVal, _ := types.ListValueFrom(ctx, types.StringType, items)\n", indent))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\treturn listVal\n", indent))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t}\n", indent))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\treturn types.ListNull(types.StringType)\n", indent))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t}(),\n", indent))
							} else if deepAttr.ElementType == "int64" {
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t%s: func() types.List {\n", indent, deepFieldName))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\tif v, ok := nestedBlockData[\"%s\"].([]interface{}); ok && len(v) > 0 {\n", indent, deepJsonName))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\tvar items []int64\n", indent))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\tfor _, item := range v {\n", indent))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\tif n, ok := item.(float64); ok {\n", indent))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t\titems = append(items, int64(n))\n", indent))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t\t}\n", indent))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\t}\n", indent))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\tlistVal, _ := types.ListValueFrom(ctx, types.Int64Type, items)\n", indent))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t\treturn listVal\n", indent))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\t}\n", indent))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t\treturn types.ListNull(types.Int64Type)\n", indent))
								sb.WriteString(fmt.Sprintf("%s\t\t\t\t\t}(),\n", indent))
							}
						}
					}

					sb.WriteString(fmt.Sprintf("%s\t\t\t\t}\n", indent))
					sb.WriteString(fmt.Sprintf("%s\t\t\t}\n", indent))
					sb.WriteString(fmt.Sprintf("%s\t\t\treturn nil\n", indent))
					sb.WriteString(fmt.Sprintf("%s\t\t}(),\n", indent))
				}
			}
			sb.WriteString(fmt.Sprintf("%s\t}\n", indent))
			sb.WriteString(fmt.Sprintf("%s}\n", indent))
			continue
		}

		switch attr.Type {
		case "string":
			sb.WriteString(fmt.Sprintf("%sif v, ok := apiResource.Spec[\"%s\"].(string); ok && v != \"\" {\n", indent, jsonName))
			sb.WriteString(fmt.Sprintf("%s\tdata.%s = types.StringValue(v)\n", indent, fieldName))
			sb.WriteString(fmt.Sprintf("%s} else {\n", indent))
			sb.WriteString(fmt.Sprintf("%s\tdata.%s = types.StringNull()\n", indent, fieldName))
			sb.WriteString(fmt.Sprintf("%s}\n", indent))
		case "int64":
			// JSON numbers are float64, so we need to convert
			// For Computed int64 fields, always set a known value after apply
			sb.WriteString(fmt.Sprintf("%sif v, ok := apiResource.Spec[\"%s\"].(float64); ok {\n", indent, jsonName))
			sb.WriteString(fmt.Sprintf("%s\tdata.%s = types.Int64Value(int64(v))\n", indent, fieldName))
			sb.WriteString(fmt.Sprintf("%s} else {\n", indent))
			sb.WriteString(fmt.Sprintf("%s\tdata.%s = types.Int64Null()\n", indent, fieldName))
			sb.WriteString(fmt.Sprintf("%s}\n", indent))
		case "bool":
			// For Optional top-level bool fields, preserve prior state during normal Read
			// This prevents drift when API returns defaults or doesn't return the field
			if attr.Optional {
				sb.WriteString(fmt.Sprintf("%s// Top-level Optional bool: preserve prior state to avoid API default drift\n", indent))
				sb.WriteString(fmt.Sprintf("%sif !isImport && !data.%s.IsNull() && !data.%s.IsUnknown() {\n", indent, fieldName, fieldName))
				sb.WriteString(fmt.Sprintf("%s\t// Normal Read: preserve existing state value (do nothing)\n", indent))
				sb.WriteString(fmt.Sprintf("%s} else {\n", indent))
				sb.WriteString(fmt.Sprintf("%s\t// Import case, null state, or unknown (after Create): read from API\n", indent))
				sb.WriteString(fmt.Sprintf("%s\tif v, ok := apiResource.Spec[\"%s\"].(bool); ok {\n", indent, jsonName))
				sb.WriteString(fmt.Sprintf("%s\t\tdata.%s = types.BoolValue(v)\n", indent, fieldName))
				sb.WriteString(fmt.Sprintf("%s\t} else {\n", indent))
				sb.WriteString(fmt.Sprintf("%s\t\tdata.%s = types.BoolNull()\n", indent, fieldName))
				sb.WriteString(fmt.Sprintf("%s\t}\n", indent))
				sb.WriteString(fmt.Sprintf("%s}\n", indent))
			} else {
				// Original behavior for Computed-only or Required fields
				sb.WriteString(fmt.Sprintf("%sif v, ok := apiResource.Spec[\"%s\"].(bool); ok {\n", indent, jsonName))
				sb.WriteString(fmt.Sprintf("%s\tdata.%s = types.BoolValue(v)\n", indent, fieldName))
				sb.WriteString(fmt.Sprintf("%s} else {\n", indent))
				sb.WriteString(fmt.Sprintf("%s\tdata.%s = types.BoolNull()\n", indent, fieldName))
				sb.WriteString(fmt.Sprintf("%s}\n", indent))
			}
		case "list":
			if attr.ElementType == "string" {
				sb.WriteString(fmt.Sprintf("%sif v, ok := apiResource.Spec[\"%s\"].([]interface{}); ok && len(v) > 0 {\n", indent, jsonName))
				sb.WriteString(fmt.Sprintf("%s\tvar %sList []string\n", indent, attr.TfsdkTag))
				sb.WriteString(fmt.Sprintf("%s\tfor _, item := range v {\n", indent))
				sb.WriteString(fmt.Sprintf("%s\t\tif s, ok := item.(string); ok {\n", indent))
				sb.WriteString(fmt.Sprintf("%s\t\t\t%sList = append(%sList, s)\n", indent, attr.TfsdkTag, attr.TfsdkTag))
				sb.WriteString(fmt.Sprintf("%s\t\t}\n", indent))
				sb.WriteString(fmt.Sprintf("%s\t}\n", indent))
				sb.WriteString(fmt.Sprintf("%s\tlistVal, diags := types.ListValueFrom(ctx, types.StringType, %sList)\n", indent, attr.TfsdkTag))
				sb.WriteString(fmt.Sprintf("%s\tresp.Diagnostics.Append(diags...)\n", indent))
				sb.WriteString(fmt.Sprintf("%s\tif !resp.Diagnostics.HasError() {\n", indent))
				sb.WriteString(fmt.Sprintf("%s\t\tdata.%s = listVal\n", indent, fieldName))
				sb.WriteString(fmt.Sprintf("%s\t}\n", indent))
				sb.WriteString(fmt.Sprintf("%s} else {\n", indent))
				sb.WriteString(fmt.Sprintf("%s\tdata.%s = types.ListNull(types.StringType)\n", indent, fieldName))
				sb.WriteString(fmt.Sprintf("%s}\n", indent))
			} else if attr.ElementType == "int64" {
				sb.WriteString(fmt.Sprintf("%sif v, ok := apiResource.Spec[\"%s\"].([]interface{}); ok && len(v) > 0 {\n", indent, jsonName))
				sb.WriteString(fmt.Sprintf("%s\tvar %sList []int64\n", indent, attr.TfsdkTag))
				sb.WriteString(fmt.Sprintf("%s\tfor _, item := range v {\n", indent))
				sb.WriteString(fmt.Sprintf("%s\t\tif n, ok := item.(float64); ok {\n", indent))
				sb.WriteString(fmt.Sprintf("%s\t\t\t%sList = append(%sList, int64(n))\n", indent, attr.TfsdkTag, attr.TfsdkTag))
				sb.WriteString(fmt.Sprintf("%s\t\t}\n", indent))
				sb.WriteString(fmt.Sprintf("%s\t}\n", indent))
				sb.WriteString(fmt.Sprintf("%s\tlistVal, diags := types.ListValueFrom(ctx, types.Int64Type, %sList)\n", indent, attr.TfsdkTag))
				sb.WriteString(fmt.Sprintf("%s\tresp.Diagnostics.Append(diags...)\n", indent))
				sb.WriteString(fmt.Sprintf("%s\tif !resp.Diagnostics.HasError() {\n", indent))
				sb.WriteString(fmt.Sprintf("%s\t\tdata.%s = listVal\n", indent, fieldName))
				sb.WriteString(fmt.Sprintf("%s\t}\n", indent))
				sb.WriteString(fmt.Sprintf("%s} else {\n", indent))
				sb.WriteString(fmt.Sprintf("%s\tdata.%s = types.ListNull(types.Int64Type)\n", indent, fieldName))
				sb.WriteString(fmt.Sprintf("%s}\n", indent))
			}
		}
	}
	return sb.String()
}

func generateResourceFile(resource *ResourceTemplate) error {
	outputPath := filepath.Join(outputDir, resource.Name+"_resource.go")

	// Create template with custom functions
	funcMap := template.FuncMap{
		"renderNestedAttrs":               renderNestedAttributes,
		"renderNestedBlocks":              renderNestedBlocks,
		"renderNestedModelTypes":          renderNestedModelTypes,
		"renderBlockFields":               renderBlockFields,
		"renderSpecStructFields":          renderSpecStructFields,
		"renderSpecMarshalCode":           renderSpecMarshalCode,
		"renderSpecMarshalCodeForCreate":  renderSpecMarshalCodeForCreate,
		"renderSpecUnmarshalCode":         renderSpecUnmarshalCode,
		"renderCreateComputedFieldsCode":  renderCreateComputedFieldsCode,
		"renderUpdateComputedFieldsCode":  renderUpdateComputedFieldsCode,
		"renderFetchedComputedFieldsCode": renderFetchedComputedFieldsCode,
		"filterSpecFields":                filterSpecFields,
	}

	tmpl, err := template.New("resource").Funcs(funcMap).Parse(resourceTemplate)
	if err != nil {
		return fmt.Errorf("template parse error: %w", err)
	}

	// Execute template to buffer first
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, resource); err != nil {
		return fmt.Errorf("template execute error: %w", err)
	}

	// Format the generated code with gofmt
	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		// If formatting fails, write unformatted code with warning
		fmt.Printf("⚠️  gofmt failed for %s: %v (writing unformatted)\n", outputPath, err)
		formatted = buf.Bytes()
	}

	return os.WriteFile(outputPath, formatted, 0644)
}

// renderNestedAttributes generates the Attributes map for nested blocks
func renderNestedAttributes(attrs []TerraformAttribute, indent string) string {
	if len(attrs) == 0 {
		return ""
	}

	var sb strings.Builder
	sb.WriteString(indent + "Attributes: map[string]schema.Attribute{\n")

	for _, attr := range attrs {
		if attr.IsBlock {
			continue // Blocks are handled separately
		}

		attrType := "String"
		switch attr.Type {
		case "int64":
			attrType = "Int64"
		case "bool":
			attrType = "Bool"
		case "map":
			attrType = "Map"
		case "list":
			attrType = "List"
		}

		// Escape backslashes and quotes in descriptions for Go string literals
		desc := escapeGoString(attr.Description)

		sb.WriteString(fmt.Sprintf("%s\t\"%s\": schema.%sAttribute{\n", indent, attr.TfsdkTag, attrType))
		sb.WriteString(fmt.Sprintf("%s\t\tMarkdownDescription: \"%s\",\n", indent, desc))

		if attr.Required {
			sb.WriteString(fmt.Sprintf("%s\t\tRequired: true,\n", indent))
		} else {
			// Handle Optional and Computed flags (can be both true for fields like 'tenant')
			if attr.Optional {
				sb.WriteString(fmt.Sprintf("%s\t\tOptional: true,\n", indent))
			}
			if attr.Computed {
				sb.WriteString(fmt.Sprintf("%s\t\tComputed: true,\n", indent))
			}
		}

		// Add PlanModifiers for attributes that need UseStateForUnknown (e.g., tenant, uid, kind in nested blocks)
		if attr.PlanModifier != "" && attr.Type == "string" {
			sb.WriteString(fmt.Sprintf("%s\t\tPlanModifiers: []planmodifier.String{\n", indent))
			sb.WriteString(fmt.Sprintf("%s\t\t\tstringplanmodifier.UseStateForUnknown(),\n", indent))
			sb.WriteString(fmt.Sprintf("%s\t\t},\n", indent))
		}

		if attr.Type == "map" || attr.Type == "list" {
			// Map ElementType to Terraform attr type
			elementTfType := "types.StringType"
			switch attr.ElementType {
			case "int64":
				elementTfType = "types.Int64Type"
			case "bool":
				elementTfType = "types.BoolType"
			}
			sb.WriteString(fmt.Sprintf("%s\t\tElementType: %s,\n", indent, elementTfType))
		}

		sb.WriteString(fmt.Sprintf("%s\t},\n", indent))
	}

	sb.WriteString(indent + "},\n")
	return sb.String()
}

// escapeGoString escapes a string for use in a Go string literal
func escapeGoString(s string) string {
	// Replace backslashes first, then other special characters
	s = strings.ReplaceAll(s, `\`, `\\`)
	s = strings.ReplaceAll(s, `"`, `\"`)
	return s
}

// NestedModelInfo holds information needed to generate a nested model type
type NestedModelInfo struct {
	TypeName    string
	Description string
	Prefix      string // Full prefix path for generating nested type references
	Attributes  []TerraformAttribute
	IsEmpty     bool
}

// collectNestedModelTypes recursively collects all nested model type definitions
func collectNestedModelTypes(resourceTitleCase string, attrs []TerraformAttribute, prefix string, collected *[]NestedModelInfo) {
	for _, attr := range attrs {
		if !attr.IsBlock {
			continue
		}

		// Build the type name: ResourceName + Prefix + AttributeName + Model
		currentPrefix := prefix + toTitleCase(attr.TfsdkTag)
		typeName := resourceTitleCase + currentPrefix + "Model"

		// Check if this block has any nested attributes (non-empty)
		hasContent := false
		for _, nested := range attr.NestedAttributes {
			if !nested.IsBlock {
				hasContent = true
				break
			}
		}

		// Also check for nested blocks
		for _, nested := range attr.NestedAttributes {
			if nested.IsBlock {
				hasContent = true
				break
			}
		}

		*collected = append(*collected, NestedModelInfo{
			TypeName:    typeName,
			Description: attr.TfsdkTag,
			Prefix:      currentPrefix, // Store the full prefix for generating nested type references
			Attributes:  attr.NestedAttributes,
			IsEmpty:     !hasContent && len(attr.NestedAttributes) == 0,
		})

		// Recursively collect nested types
		if len(attr.NestedAttributes) > 0 {
			collectNestedModelTypes(resourceTitleCase, attr.NestedAttributes, currentPrefix, collected)
		}
	}
}

// renderNestedModelTypes generates all nested model struct definitions for a resource
func renderNestedModelTypes(resourceTitleCase string, attrs []TerraformAttribute) string {
	// First check if there are any blocks
	hasBlocks := false
	for _, attr := range attrs {
		if attr.IsBlock {
			hasBlocks = true
			break
		}
	}
	if !hasBlocks {
		return ""
	}

	var sb strings.Builder

	// Add empty model type for blocks with no attributes
	sb.WriteString(fmt.Sprintf("// %sEmptyModel represents empty nested blocks\n", resourceTitleCase))
	sb.WriteString(fmt.Sprintf("type %sEmptyModel struct {\n}\n\n", resourceTitleCase))

	// Collect all nested model types
	var models []NestedModelInfo
	collectNestedModelTypes(resourceTitleCase, attrs, "", &models)

	// Generate each model type
	for _, model := range models {
		if model.IsEmpty {
			continue // Empty models use the shared EmptyModel
		}

		sb.WriteString(fmt.Sprintf("// %s represents %s block\n", model.TypeName, model.Description))
		sb.WriteString(fmt.Sprintf("type %s struct {\n", model.TypeName))

		// Generate fields for non-block attributes
		for _, attr := range model.Attributes {
			if attr.IsBlock {
				continue
			}
			goType := "String"
			switch attr.Type {
			case "int64":
				goType = "Int64"
			case "bool":
				goType = "Bool"
			case "map":
				goType = "Map"
			case "list":
				goType = "List"
			}
			sb.WriteString(fmt.Sprintf("\t%s types.%s `tfsdk:\"%s\"`\n", attr.GoName, goType, attr.TfsdkTag))
		}

		// Generate pointer fields for nested block attributes
		for _, attr := range model.Attributes {
			if !attr.IsBlock {
				continue
			}
			// Use the model's full prefix to build the nested type name
			nestedTypeName := resourceTitleCase + model.Prefix + toTitleCase(attr.TfsdkTag) + "Model"

			// Check if this nested block is empty
			isNestedEmpty := len(attr.NestedAttributes) == 0
			if !isNestedEmpty {
				hasNonBlockAttrs := false
				for _, nested := range attr.NestedAttributes {
					if !nested.IsBlock {
						hasNonBlockAttrs = true
						break
					}
				}
				isNestedEmpty = !hasNonBlockAttrs && len(attr.NestedAttributes) == 0
			}

			if isNestedEmpty {
				nestedTypeName = resourceTitleCase + "EmptyModel"
			}

			// For nested blocks within other models, use Go slices/pointers (not types.List)
			// The types.List migration only applies to TOP-LEVEL blocks on the ResourceModel
			// because those are the ones that can be dynamically set to unknown during planning.
			// Nested blocks inside other blocks are already within a known structure.
			if attr.NestedBlockType == "list" {
				sb.WriteString(fmt.Sprintf("\t%s []%s `tfsdk:\"%s\"`\n", attr.GoName, nestedTypeName, attr.TfsdkTag))
			} else {
				sb.WriteString(fmt.Sprintf("\t%s *%s `tfsdk:\"%s\"`\n", attr.GoName, nestedTypeName, attr.TfsdkTag))
			}
		}

		sb.WriteString("}\n\n")

		// Generate AttrTypes map for this model (needed for types.List conversion)
		sb.WriteString(fmt.Sprintf("// %sAttrTypes defines the attribute types for %s\n", model.TypeName, model.TypeName))
		sb.WriteString(fmt.Sprintf("var %sAttrTypes = map[string]attr.Type{\n", model.TypeName))

		// Generate attr.Type for non-block attributes
		for _, attr := range model.Attributes {
			if attr.IsBlock {
				continue
			}
			var attrType string
			switch attr.Type {
			case "string":
				attrType = "types.StringType"
			case "int64":
				attrType = "types.Int64Type"
			case "bool":
				attrType = "types.BoolType"
			case "map":
				attrType = "types.MapType{ElemType: types.StringType}"
			case "list":
				elemType := "types.StringType"
				switch attr.ElementType {
				case "int64":
					elemType = "types.Int64Type"
				case "bool":
					elemType = "types.BoolType"
				}
				attrType = fmt.Sprintf("types.ListType{ElemType: %s}", elemType)
			default:
				attrType = "types.StringType"
			}
			sb.WriteString(fmt.Sprintf("\t\"%s\": %s,\n", attr.TfsdkTag, attrType))
		}

		// Generate attr.Type for block attributes
		for _, attr := range model.Attributes {
			if !attr.IsBlock {
				continue
			}

			// Check if this nested block is empty (has no nested content)
			// If NestedAttributes has any content (blocks or attributes), a model with
			// AttrTypes is generated and should be referenced. Only truly empty blocks
			// should use inline empty maps.
			// Fix for Issue #452: Blocks containing only sub-blocks are NOT empty -
			// they need to reference their model's AttrTypes to avoid Value Conversion Errors.
			isNestedEmpty := len(attr.NestedAttributes) == 0

			if isNestedEmpty {
				// Empty nested block - use inline empty map
				if attr.NestedBlockType == "list" {
					sb.WriteString(fmt.Sprintf("\t\"%s\": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{}}},\n", attr.TfsdkTag))
				} else {
					sb.WriteString(fmt.Sprintf("\t\"%s\": types.ObjectType{AttrTypes: map[string]attr.Type{}},\n", attr.TfsdkTag))
				}
			} else {
				// Non-empty nested block - reference its AttrTypes variable
				nestedAttrTypesName := resourceTitleCase + model.Prefix + toTitleCase(attr.TfsdkTag) + "ModelAttrTypes"
				if attr.NestedBlockType == "list" {
					// ListNestedBlock → types.ListType{ElemType: types.ObjectType{AttrTypes: ...}}
					sb.WriteString(fmt.Sprintf("\t\"%s\": types.ListType{ElemType: types.ObjectType{AttrTypes: %s}},\n", attr.TfsdkTag, nestedAttrTypesName))
				} else {
					// SingleNestedBlock → types.ObjectType{AttrTypes: ...}
					sb.WriteString(fmt.Sprintf("\t\"%s\": types.ObjectType{AttrTypes: %s},\n", attr.TfsdkTag, nestedAttrTypesName))
				}
			}
		}

		sb.WriteString("}\n\n")
	}

	return sb.String()
}

// renderBlockFields generates the block fields for the main ResourceModel struct
func renderBlockFields(resourceTitleCase string, attrs []TerraformAttribute) string {
	var sb strings.Builder

	for _, attr := range attrs {
		if !attr.IsBlock {
			continue
		}

		// Determine the type name
		typeName := resourceTitleCase + toTitleCase(attr.TfsdkTag) + "Model"

		// Check if this block is empty
		isBlockEmpty := len(attr.NestedAttributes) == 0
		if !isBlockEmpty {
			hasNonBlockAttrs := false
			for _, nested := range attr.NestedAttributes {
				if !nested.IsBlock {
					hasNonBlockAttrs = true
					break
				}
			}
			// Also check for any nested blocks
			hasNestedBlocks := false
			for _, nested := range attr.NestedAttributes {
				if nested.IsBlock {
					hasNestedBlocks = true
					break
				}
			}
			isBlockEmpty = !hasNonBlockAttrs && !hasNestedBlocks
		}

		if isBlockEmpty {
			typeName = resourceTitleCase + "EmptyModel"
		}

		// For list nested blocks, use types.List to properly handle unknown values during planning.
		// Go slices cannot represent "unknown" state, causing errors with dynamic blocks.
		// For single nested blocks, use pointer type.
		if attr.NestedBlockType == "list" {
			sb.WriteString(fmt.Sprintf("\t%s types.List `tfsdk:\"%s\"`\n", attr.GoName, attr.TfsdkTag))
		} else {
			sb.WriteString(fmt.Sprintf("\t%s *%s `tfsdk:\"%s\"`\n", attr.GoName, typeName, attr.TfsdkTag))
		}
	}

	return sb.String()
}

// renderNestedBlocks generates the Blocks map for nested blocks within a block
func renderNestedBlocks(attrs []TerraformAttribute, indent string) string {
	var hasBlocks bool
	for _, attr := range attrs {
		if attr.IsBlock {
			hasBlocks = true
			break
		}
	}

	if !hasBlocks {
		return ""
	}

	var sb strings.Builder
	sb.WriteString(indent + "Blocks: map[string]schema.Block{\n")

	for _, attr := range attrs {
		if !attr.IsBlock {
			continue
		}

		blockType := "SingleNestedBlock"
		if attr.NestedBlockType == "list" {
			blockType = "ListNestedBlock"
		}

		// Escape backslashes and quotes in descriptions
		desc := escapeGoString(attr.Description)

		sb.WriteString(fmt.Sprintf("%s\t\"%s\": schema.%s{\n", indent, attr.TfsdkTag, blockType))
		sb.WriteString(fmt.Sprintf("%s\t\tMarkdownDescription: \"%s\",\n", indent, desc))

		if attr.NestedBlockType == "list" {
			sb.WriteString(fmt.Sprintf("%s\t\tNestedObject: schema.NestedBlockObject{\n", indent))
			if len(attr.NestedAttributes) > 0 {
				sb.WriteString(renderNestedAttributes(attr.NestedAttributes, indent+"\t\t\t"))
				sb.WriteString(renderNestedBlocks(attr.NestedAttributes, indent+"\t\t\t"))
			} else {
				sb.WriteString(fmt.Sprintf("%s\t\t\tAttributes: map[string]schema.Attribute{},\n", indent))
			}
			sb.WriteString(fmt.Sprintf("%s\t\t},\n", indent))
		} else {
			// SingleNestedBlock
			if len(attr.NestedAttributes) > 0 {
				sb.WriteString(renderNestedAttributes(attr.NestedAttributes, indent+"\t\t"))
				sb.WriteString(renderNestedBlocks(attr.NestedAttributes, indent+"\t\t"))
			}
		}

		sb.WriteString(fmt.Sprintf("%s\t},\n", indent))
	}

	sb.WriteString(indent + "},\n")
	return sb.String()
}

func generateClientTypes(resource *ResourceTemplate) error {
	outputPath := filepath.Join(clientDir, resource.Name+"_types.go")

	// Create template with custom functions for spec field generation
	funcMap := template.FuncMap{
		"renderSpecStructFields": func(attrs []TerraformAttribute) string {
			return renderSpecStructFields(attrs, "\t")
		},
	}

	tmpl, err := template.New("client").Funcs(funcMap).Parse(clientTemplate)
	if err != nil {
		return fmt.Errorf("template parse error: %w", err)
	}

	// Execute template to buffer first
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, resource); err != nil {
		return fmt.Errorf("template execute error: %w", err)
	}

	// Format the generated code with gofmt
	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		// If formatting fails, write unformatted code with warning
		fmt.Printf("⚠️  gofmt failed for %s: %v (writing unformatted)\n", outputPath, err)
		formatted = buf.Bytes()
	}

	return os.WriteFile(outputPath, formatted, 0644)
}

func generateDataSource(resource *ResourceTemplate) error {
	outputPath := filepath.Join(outputDir, resource.Name+"_data_source.go")

	tmpl, err := template.New("datasource").Parse(dataSourceTemplate)
	if err != nil {
		return fmt.Errorf("template parse error: %w", err)
	}

	// Execute template to buffer first
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, resource); err != nil {
		return fmt.Errorf("template execute error: %w", err)
	}

	// Format the generated code with gofmt
	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		// If formatting fails, write unformatted code with warning
		fmt.Printf("⚠️  gofmt failed for %s: %v (writing unformatted)\n", outputPath, err)
		formatted = buf.Bytes()
	}

	return os.WriteFile(outputPath, formatted, 0644)
}

func generateCombinedClientTypes(results []GenerationResult) {
	// This is handled by individual client type files
}

// coreResources are resources that must always be registered in the provider,
// even if they're not present in the current OpenAPI specifications.
// These resources have working implementations that were generated previously
// but may not be included in the enriched API specs.
// Note: namespace was removed in v3.0.0 as part of backwards compatibility cleanup
var coreResources = []string{}

func generateProviderRegistration(results []GenerationResult) {
	// Collect successful resources
	var resources []string
	var dataSources []string

	// First, add core resources that must always be registered
	// These have working implementations but may not be in current OpenAPI specs
	added := make(map[string]bool)
	for _, core := range coreResources {
		titleCase := toTitleCase(core)
		resources = append(resources, fmt.Sprintf("\t\tNew%sResource,", titleCase))
		dataSources = append(dataSources, fmt.Sprintf("\t\tNew%sDataSource,", titleCase))
		added[core] = true
	}

	// Then add resources from spec generation results (avoiding duplicates)
	for _, r := range results {
		if r.Success && !added[r.ResourceName] {
			titleCase := toTitleCase(r.ResourceName)
			resources = append(resources, fmt.Sprintf("\t\tNew%sResource,", titleCase))
			dataSources = append(dataSources, fmt.Sprintf("\t\tNew%sDataSource,", titleCase))
		}
	}

	// Sort for consistent output
	sort.Strings(resources)
	sort.Strings(dataSources)

	// Generate provider.go file
	providerPath := filepath.Join(outputDir, "provider.go")
	fmt.Printf("\n📝 Updating provider.go with %d resources and %d data sources...\n", len(resources), len(dataSources))

	providerContent := fmt.Sprintf(`// Code generated by generate-all-schemas.go. DO NOT EDIT.
// Source: F5 XC OpenAPI specification

package provider

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/f5xc/terraform-provider-f5xc/internal/client"
)

// Ensure F5XCProvider satisfies various provider interfaces.
var _ provider.Provider = &F5XCProvider{}

// F5XCProvider defines the provider implementation.
type F5XCProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// F5XCProviderModel describes the provider data model.
type F5XCProviderModel struct {
	APIToken     types.String `+"`"+`tfsdk:"api_token"`+"`"+`
	APIURL       types.String `+"`"+`tfsdk:"api_url"`+"`"+`
	APIP12File   types.String `+"`"+`tfsdk:"api_p12_file"`+"`"+`
	P12Password  types.String `+"`"+`tfsdk:"p12_password"`+"`"+`
	APICert      types.String `+"`"+`tfsdk:"api_cert"`+"`"+`
	APIKey       types.String `+"`"+`tfsdk:"api_key"`+"`"+`
	APICACert    types.String `+"`"+`tfsdk:"api_ca_cert"`+"`"+`
}

func (p *F5XCProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "f5xc"
	resp.Version = p.version
}

func (p *F5XCProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Terraform provider for F5 Distributed Cloud (F5XC) enabling infrastructure as code " +
			"for load balancers, security policies, sites, and networking. Community-maintained provider " +
			"built from public F5 API documentation.",
		Attributes: map[string]schema.Attribute{
			"api_url": schema.StringAttribute{
				MarkdownDescription: "F5 Distributed Cloud API URL. " +
					"Defaults to https://console.ves.volterra.io. " +
					"Example: https://tenant.console.ves.volterra.io. " +
					"Can also be set via F5XC_API_URL environment variable.",
				Optional: true,
			},
			"api_token": schema.StringAttribute{
				MarkdownDescription: "F5 Distributed Cloud API Token for token-based authentication. " +
					"Can also be set via F5XC_API_TOKEN environment variable. " +
					"Either api_token or api_p12_file/api_cert must be specified.",
				Optional:  true,
				Sensitive: true,
			},
			"api_p12_file": schema.StringAttribute{
				MarkdownDescription: "Path to PKCS#12 certificate bundle file for certificate-based authentication. " +
					"Can also be set via F5XC_P12_FILE environment variable. " +
					"When using P12 authentication, p12_password must also be provided.",
				Optional:  true,
				Sensitive: false,
			},
			"p12_password": schema.StringAttribute{
				MarkdownDescription: "Password for the PKCS#12 certificate bundle. " +
					"Can also be set via F5XC_P12_PASSWORD environment variable.",
				Optional:  true,
				Sensitive: true,
			},
			"api_cert": schema.StringAttribute{
				MarkdownDescription: "Path to PEM-encoded client certificate file for certificate-based authentication. " +
					"Can also be set via F5XC_CERT environment variable. " +
					"When using certificate authentication, api_key must also be provided.",
				Optional: true,
			},
			"api_key": schema.StringAttribute{
				MarkdownDescription: "Path to PEM-encoded client private key file for certificate-based authentication. " +
					"Can also be set via F5XC_KEY environment variable.",
				Optional:  true,
				Sensitive: true,
			},
			"api_ca_cert": schema.StringAttribute{
				MarkdownDescription: "Path to PEM-encoded CA certificate file for verifying the F5XC API server. " +
					"Can also be set via F5XC_CACERT environment variable. Optional.",
				Optional: true,
			},
		},
	}
}

func (p *F5XCProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Info(ctx, "Configuring F5XC client")

	var config F5XCProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Get configuration values from environment variables first
	apiURL := os.Getenv("F5XC_API_URL")
	apiToken := os.Getenv("F5XC_API_TOKEN")
	apiP12File := os.Getenv("F5XC_P12_FILE")
	p12Password := os.Getenv("F5XC_P12_PASSWORD")
	apiCert := os.Getenv("F5XC_CERT")
	apiKey := os.Getenv("F5XC_KEY")
	apiCACert := os.Getenv("F5XC_CACERT")

	// Configuration values override environment variables
	if !config.APIURL.IsNull() {
		apiURL = config.APIURL.ValueString()
	}
	if !config.APIToken.IsNull() {
		apiToken = config.APIToken.ValueString()
	}
	if !config.APIP12File.IsNull() {
		apiP12File = config.APIP12File.ValueString()
	}
	if !config.P12Password.IsNull() {
		p12Password = config.P12Password.ValueString()
	}
	if !config.APICert.IsNull() {
		apiCert = config.APICert.ValueString()
	}
	if !config.APIKey.IsNull() {
		apiKey = config.APIKey.ValueString()
	}
	if !config.APICACert.IsNull() {
		apiCACert = config.APICACert.ValueString()
	}

	// Set default API URL if not provided
	if apiURL == "" {
		apiURL = "https://console.ves.volterra.io"
	}

	// Normalize the API URL (removes /api suffix and trailing slashes)
	apiURL, _ = normalizeAPIURL(apiURL)

	var c *client.Client
	var err error

	// Determine authentication method
	switch {
	case apiP12File != "":
		// P12 certificate authentication
		if p12Password == "" {
			resp.Diagnostics.AddAttributeError(
				path.Root("p12_password"),
				"Missing P12 Password",
				"When using P12 certificate authentication (api_p12_file), the p12_password must be provided. "+
					"Set the p12_password value in the configuration or use the F5XC_P12_PASSWORD environment variable.",
			)
			return
		}
		c, err = client.NewClientWithP12(apiURL, apiP12File, p12Password)
		if err != nil {
			resp.Diagnostics.AddError(
				"Failed to Create F5XC Client",
				"Could not create F5XC client with P12 certificate: "+err.Error(),
			)
			return
		}
		tflog.Info(ctx, "Configured F5XC client with P12 certificate authentication", map[string]any{"success": true, "api_url": apiURL})

	case apiCert != "" && apiKey != "":
		// PEM certificate/key authentication
		c, err = client.NewClientWithCert(apiURL, apiCert, apiKey, apiCACert)
		if err != nil {
			resp.Diagnostics.AddError(
				"Failed to Create F5XC Client",
				"Could not create F5XC client with certificate: "+err.Error(),
			)
			return
		}
		tflog.Info(ctx, "Configured F5XC client with certificate authentication", map[string]any{"success": true, "api_url": apiURL})

	case apiToken != "":
		// API token authentication
		c = client.NewClient(apiURL, apiToken)
		tflog.Info(ctx, "Configured F5XC client with API token authentication", map[string]any{"success": true, "api_url": apiURL})

	default:
		resp.Diagnostics.AddError(
			"Missing Authentication Configuration",
			"The provider requires authentication. Please configure one of the following:\n"+
				"  - api_token (or F5XC_API_TOKEN environment variable) for API token authentication\n"+
				"  - api_p12_file and p12_password (or F5XC_P12_FILE and F5XC_P12_PASSWORD environment variables) for P12 certificate authentication\n"+
				"  - api_cert and api_key (or F5XC_CERT and F5XC_KEY environment variables) for PEM certificate authentication",
		)
		return
	}

	// Make the client available during DataSource and Resource type Configure methods
	resp.DataSourceData = c
	resp.ResourceData = c
}

func (p *F5XCProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
%s
	}
}

func (p *F5XCProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
%s
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &F5XCProvider{
			version: version,
		}
	}
}
`, strings.Join(resources, "\n"), strings.Join(dataSources, "\n"))

	// Format the generated code with gofmt
	formatted, err := format.Source([]byte(providerContent))
	if err != nil {
		// If formatting fails, write unformatted code with warning
		fmt.Printf("⚠️  gofmt failed for %s: %v (writing unformatted)\n", providerPath, err)
		formatted = []byte(providerContent)
	}

	if err := os.WriteFile(providerPath, formatted, 0644); err != nil {
		fmt.Printf("❌ Error writing provider.go: %v\n", err)
		return
	}

	fmt.Printf("✅ Updated %s\n", providerPath)
}

const resourceTemplate = `// Code generated by generate-all-schemas.go. DO NOT EDIT.
// Source: F5 XC OpenAPI specification

package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
{{- if .UsesBoolPlanModifier}}
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
{{- end}}
{{- if .UsesInt64PlanModifier}}
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
{{- end}}
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
{{- if .HasBlocks}}
	"github.com/hashicorp/terraform-plugin-framework/attr"
{{- end}}
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/f5xc/terraform-provider-f5xc/internal/client"
	inttimeouts "github.com/f5xc/terraform-provider-f5xc/internal/timeouts"
	"github.com/f5xc/terraform-provider-f5xc/internal/validators"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                   = &{{.TitleCase}}Resource{}
	_ resource.ResourceWithConfigure      = &{{.TitleCase}}Resource{}
	_ resource.ResourceWithImportState    = &{{.TitleCase}}Resource{}
	_ resource.ResourceWithModifyPlan     = &{{.TitleCase}}Resource{}
	_ resource.ResourceWithValidateConfig = &{{.TitleCase}}Resource{}
)

func New{{.TitleCase}}Resource() resource.Resource {
	return &{{.TitleCase}}Resource{}
}

type {{.TitleCase}}Resource struct {
	client *client.Client
}

{{renderNestedModelTypes .TitleCase .Attributes}}type {{.TitleCase}}ResourceModel struct {
{{- range .Attributes}}
{{- if not .IsBlock}}
	{{.GoName}} types.{{if eq .Type "string"}}String{{else if eq .Type "int64"}}Int64{{else if eq .Type "bool"}}Bool{{else if eq .Type "map"}}Map{{else if eq .Type "list"}}List{{else}}String{{end}} ` + "`" + `tfsdk:"{{.TfsdkTag}}"` + "`" + `
{{- end}}
{{- end}}
	Timeouts timeouts.Value ` + "`" + `tfsdk:"timeouts"` + "`" + `
{{renderBlockFields .TitleCase .Attributes}}}

func (r *{{.TitleCase}}Resource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_{{.Name}}"
}

func (r *{{.TitleCase}}Resource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "{{.Description}}",
		Attributes: map[string]schema.Attribute{
{{- range .Attributes}}
{{- if not .IsBlock}}
			"{{.TfsdkTag}}": schema.{{if eq .Type "string"}}String{{else if eq .Type "int64"}}Int64{{else if eq .Type "bool"}}Bool{{else if eq .Type "map"}}Map{{else if eq .Type "list"}}List{{else}}String{{end}}Attribute{
				MarkdownDescription: "{{.Description}}",
{{- if .Required}}
				Required: true,
{{- end}}
{{- if and .Optional (not .Required)}}
				Optional: true,
{{- end}}
{{- if .Computed}}
				Computed: true,
{{- end}}
{{- if eq .Type "map"}}
				ElementType: types.StringType,
{{- end}}
{{- if eq .Type "list"}}
				ElementType: {{if eq .ElementType "int64"}}types.Int64Type{{else if eq .ElementType "bool"}}types.BoolType{{else}}types.StringType{{end}},
{{- end}}
{{- if .PlanModifier}}
				PlanModifiers: []planmodifier.{{if eq .Type "string"}}String{{else if eq .Type "bool"}}Bool{{else if eq .Type "int64"}}Int64{{else if eq .Type "list"}}List{{else if eq .Type "map"}}Map{{else}}String{{end}}{
{{- if eq .PlanModifier "RequiresReplace"}}
					{{if eq .Type "string"}}stringplanmodifier{{else if eq .Type "bool"}}boolplanmodifier{{else if eq .Type "int64"}}int64planmodifier{{else if eq .Type "list"}}listplanmodifier{{else if eq .Type "map"}}mapplanmodifier{{else}}stringplanmodifier{{end}}.RequiresReplace(),
{{- else if eq .PlanModifier "UseStateForUnknown"}}
					{{if eq .Type "string"}}stringplanmodifier{{else if eq .Type "bool"}}boolplanmodifier{{else if eq .Type "int64"}}int64planmodifier{{else if eq .Type "list"}}listplanmodifier{{else if eq .Type "map"}}mapplanmodifier{{else}}stringplanmodifier{{end}}.UseStateForUnknown(),
{{- end}}
				},
{{- end}}
{{- if eq .TfsdkTag "name"}}
				Validators: []validator.String{
{{- if .UseDomainValidator}}
					validators.DomainValidator(),
{{- else}}
					validators.NameValidator(),
{{- end}}
				},
{{- else if eq .TfsdkTag "namespace"}}
				Validators: []validator.String{
					validators.NamespaceValidator(),
				},
{{- end}}
			},
{{- end}}
{{- end}}
		},
		Blocks: map[string]schema.Block{
			"timeouts": timeouts.Block(ctx, timeouts.Opts{
				Create: true,
				Read:   true,
				Update: true,
				Delete: true,
			}),
{{- range .Attributes}}
{{- if .IsBlock}}
			"{{.TfsdkTag}}": schema.{{if eq .NestedBlockType "single"}}SingleNestedBlock{{else if eq .NestedBlockType "list"}}ListNestedBlock{{else}}SingleNestedBlock{{end}}{
				MarkdownDescription: "{{.Description}}",
{{- if eq .NestedBlockType "list"}}
				NestedObject: schema.NestedBlockObject{
{{- if .NestedAttributes}}
{{renderNestedAttrs .NestedAttributes "\t\t\t\t\t"}}{{renderNestedBlocks .NestedAttributes "\t\t\t\t\t"}}{{- else}}
					Attributes: map[string]schema.Attribute{},
{{- end}}
				},
{{- else}}
{{- if .NestedAttributes}}
{{renderNestedAttrs .NestedAttributes "\t\t\t\t"}}{{renderNestedBlocks .NestedAttributes "\t\t\t\t"}}{{- end}}
{{- end}}
			},
{{- end}}
{{- end}}
		},
	}
}

func (r *{{.TitleCase}}Resource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}
	r.client = client
}

// ValidateConfig implements resource.ResourceWithValidateConfig
func (r *{{.TitleCase}}Resource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data {{.TitleCase}}ResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// ModifyPlan implements resource.ResourceWithModifyPlan
func (r *{{.TitleCase}}Resource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.Plan.Raw.IsNull() {
		resp.Diagnostics.AddWarning(
			"Resource Destruction",
			"This will permanently delete the {{.Name}} from F5 Distributed Cloud.",
		)
		return
	}

	if req.State.Raw.IsNull() {
		var plan {{.TitleCase}}ResourceModel
		resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
		if resp.Diagnostics.HasError() {
			return
		}

		if plan.Name.IsUnknown() {
			resp.Diagnostics.AddWarning(
				"Unknown Resource Name",
				"The resource name is not yet known. This may affect planning for dependent resources.",
			)
		}
	}
}

func (r *{{.TitleCase}}Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data {{.TitleCase}}ResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createTimeout, diags := data.Timeouts.Create(ctx, inttimeouts.DefaultCreate)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	ctx, cancel := context.WithTimeout(ctx, createTimeout)
	defer cancel()

	tflog.Debug(ctx, "Creating {{.Name}}", map[string]interface{}{
		"name":      data.Name.ValueString(),
		"namespace": data.Namespace.ValueString(),
	})

	createReq := &client.{{.TitleCase}}{
		Metadata: client.Metadata{
			Name:      data.Name.ValueString(),
			Namespace: data.Namespace.ValueString(),
		},
		Spec: make(map[string]interface{}),
	}

	if !data.Description.IsNull() {
		createReq.Metadata.Description = data.Description.ValueString()
	}

	if !data.Labels.IsNull() {
		labels := make(map[string]string)
		resp.Diagnostics.Append(data.Labels.ElementsAs(ctx, &labels, false)...)
		if resp.Diagnostics.HasError() {
			return
		}
		createReq.Metadata.Labels = labels
	}

	if !data.Annotations.IsNull() {
		annotations := make(map[string]string)
		resp.Diagnostics.Append(data.Annotations.ElementsAs(ctx, &annotations, false)...)
		if resp.Diagnostics.HasError() {
			return
		}
		createReq.Metadata.Annotations = annotations
	}

	// Marshal spec fields from Terraform state to API struct
{{renderSpecMarshalCodeForCreate .Attributes "\t" .TitleCase}}

	apiResource, err := r.client.Create{{.TitleCase}}(ctx, createReq)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create {{.TitleCase}}: %s", err))
		return
	}

	data.ID = types.StringValue(apiResource.Metadata.Name)
{{- if not .HasNamespaceInPath}}
	// For resources without namespace in API path, namespace is computed from API response
	data.Namespace = types.StringValue(apiResource.Metadata.Namespace)
{{- end}}

	// Unmarshal spec fields from API response to Terraform state
	// This ensures computed nested fields (like tenant in Object Reference blocks) have known values
	isImport := false // Create is never an import
	_ = isImport // May be unused if resource has no blocks needing import detection
{{renderSpecUnmarshalCode .Attributes "\t" .TitleCase}}

	tflog.Trace(ctx, "created {{.TitleCase}} resource")
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *{{.TitleCase}}Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data {{.TitleCase}}ResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	readTimeout, diags := data.Timeouts.Read(ctx, inttimeouts.DefaultRead)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	ctx, cancel := context.WithTimeout(ctx, readTimeout)
	defer cancel()

	apiResource, err := r.client.Get{{.TitleCase}}(ctx, data.Namespace.ValueString(), data.Name.ValueString())
	if err != nil {
		// Check if the resource was deleted outside Terraform
		if strings.Contains(err.Error(), "NOT_FOUND") || strings.Contains(err.Error(), "404") {
			tflog.Warn(ctx, "{{.TitleCase}} not found, removing from state", map[string]interface{}{
				"name":      data.Name.ValueString(),
				"namespace": data.Namespace.ValueString(),
			})
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read {{.TitleCase}}: %s", err))
		return
	}

	data.ID = types.StringValue(apiResource.Metadata.Name)
	data.Name = types.StringValue(apiResource.Metadata.Name)
	data.Namespace = types.StringValue(apiResource.Metadata.Namespace)

	// Read description from metadata
	if apiResource.Metadata.Description != "" {
		data.Description = types.StringValue(apiResource.Metadata.Description)
	} else {
		data.Description = types.StringNull()
	}

	// Filter out system-managed labels (ves.io/*) that are injected by the platform
	if len(apiResource.Metadata.Labels) > 0 {
		filteredLabels := filterSystemLabels(apiResource.Metadata.Labels)
		if len(filteredLabels) > 0 {
			labels, diags := types.MapValueFrom(ctx, types.StringType, filteredLabels)
			resp.Diagnostics.Append(diags...)
			if !resp.Diagnostics.HasError() {
				data.Labels = labels
			}
		} else {
			data.Labels = types.MapNull(types.StringType)
		}
	} else {
		data.Labels = types.MapNull(types.StringType)
	}

	if len(apiResource.Metadata.Annotations) > 0 {
		annotations, diags := types.MapValueFrom(ctx, types.StringType, apiResource.Metadata.Annotations)
		resp.Diagnostics.Append(diags...)
		if !resp.Diagnostics.HasError() {
			data.Annotations = annotations
		}
	} else {
		data.Annotations = types.MapNull(types.StringType)
	}

	// Check if this Read is triggered by an import operation
	// Import sets a private state marker so we know to populate all nested blocks from API response
	isImport := false
	if importMarker, diags := req.Private.GetKey(ctx, "isImport"); diags.HasError() == false && string(importMarker) == "true" {
		isImport = true
	}
	_ = isImport // May be unused if resource has no blocks needing import detection
{{renderSpecUnmarshalCode .Attributes "\t" .TitleCase}}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *{{.TitleCase}}Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data {{.TitleCase}}ResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	updateTimeout, diags := data.Timeouts.Update(ctx, inttimeouts.DefaultUpdate)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	ctx, cancel := context.WithTimeout(ctx, updateTimeout)
	defer cancel()

	apiResource := &client.{{.TitleCase}}{
		Metadata: client.Metadata{
			Name:      data.Name.ValueString(),
			Namespace: data.Namespace.ValueString(),
		},
		Spec: make(map[string]interface{}),
	}

	if !data.Description.IsNull() {
		apiResource.Metadata.Description = data.Description.ValueString()
	}

	if !data.Labels.IsNull() {
		labels := make(map[string]string)
		resp.Diagnostics.Append(data.Labels.ElementsAs(ctx, &labels, false)...)
		if resp.Diagnostics.HasError() {
			return
		}
		apiResource.Metadata.Labels = labels
	}

	if !data.Annotations.IsNull() {
		annotations := make(map[string]string)
		resp.Diagnostics.Append(data.Annotations.ElementsAs(ctx, &annotations, false)...)
		if resp.Diagnostics.HasError() {
			return
		}
		apiResource.Metadata.Annotations = annotations
	}

	// Marshal spec fields from Terraform state to API struct
{{renderSpecMarshalCode .Attributes "\t" .TitleCase}}

	_, err := r.client.Update{{.TitleCase}}(ctx, apiResource)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update {{.TitleCase}}: %s", err))
		return
	}

	// Use plan data for ID since API response may not include metadata.name
	data.ID = types.StringValue(data.Name.ValueString())

	// Fetch the resource to get complete state including computed fields
	// PUT responses may not include all computed nested fields (like tenant in Object Reference blocks)
	fetched, fetchErr := r.client.Get{{.TitleCase}}(ctx, data.Namespace.ValueString(), data.Name.ValueString())
	if fetchErr != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read {{.TitleCase}} after update: %s", fetchErr))
		return
	}

{{renderFetchedComputedFieldsCode .Attributes "\t"}}

	// Unmarshal spec fields from fetched resource to Terraform state
	apiResource = fetched // Use GET response which includes all computed fields
	isImport := false // Update is never an import
	_ = isImport // May be unused if resource has no blocks needing import detection
{{renderSpecUnmarshalCode .Attributes "\t" .TitleCase}}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *{{.TitleCase}}Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data {{.TitleCase}}ResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	deleteTimeout, diags := data.Timeouts.Delete(ctx, inttimeouts.DefaultDelete)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	ctx, cancel := context.WithTimeout(ctx, deleteTimeout)
	defer cancel()

{{- if eq .TitleCase "Namespace"}}
	// Namespace requires cascade_delete endpoint (standard DELETE returns 501)
	err := r.client.CascadeDeleteNamespace(ctx, data.Name.ValueString())
{{- else}}
	err := r.client.Delete{{.TitleCase}}(ctx, data.Namespace.ValueString(), data.Name.ValueString())
{{- end}}
	if err != nil {
		// If the resource is already gone, consider deletion successful (idempotent delete)
		if strings.Contains(err.Error(), "NOT_FOUND") || strings.Contains(err.Error(), "404") {
			tflog.Warn(ctx, "{{.TitleCase}} already deleted, removing from state", map[string]interface{}{
				"name":      data.Name.ValueString(),
				"namespace": data.Namespace.ValueString(),
			})
			return
		}
		// If delete is not implemented (501), warn and remove from state
		// Some F5 XC resources don't support deletion via API
		if strings.Contains(err.Error(), "501") {
			tflog.Warn(ctx, "{{.TitleCase}} delete not supported by API (501), removing from state only", map[string]interface{}{
				"name":      data.Name.ValueString(),
				"namespace": data.Namespace.ValueString(),
			})
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete {{.TitleCase}}: %s", err))
		return
	}
}

func (r *{{.TitleCase}}Resource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
{{- if .HasNamespaceInPath}}
	// Import ID format: namespace/name
	parts := strings.Split(req.ID, "/")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		resp.Diagnostics.AddError(
			"Invalid Import ID",
			fmt.Sprintf("Expected import ID format: namespace/name, got: %s", req.ID),
		)
		return
	}
	namespace := parts[0]
	name := parts[1]

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("namespace"), namespace)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("name"), name)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), name)...)
{{- else}}
	// Import ID format: name (no namespace for this resource type)
	name := req.ID
	if name == "" {
		resp.Diagnostics.AddError(
			"Invalid Import ID",
			"Expected import ID to be the resource name, got empty string",
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("namespace"), "")...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("name"), name)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), name)...)
{{- end}}

	// Set private state marker to indicate this is an import operation
	// This allows Read to populate all nested blocks from API response
	diags := resp.Private.SetKey(ctx, "isImport", []byte("true"))
	resp.Diagnostics.Append(diags...)
}
`

const clientTemplate = `// Code generated by generate-all-schemas.go. DO NOT EDIT.
// Source: F5 XC OpenAPI specification

package client

import (
	"context"
	"fmt"
)

// {{.TitleCase}} represents a F5XC {{.TitleCase}}
type {{.TitleCase}} struct {
	Metadata Metadata               ` + "`" + `json:"metadata"` + "`" + `
	Spec     map[string]interface{} ` + "`" + `json:"spec"` + "`" + `
}

// Create{{.TitleCase}} creates a new {{.TitleCase}}
func (c *Client) Create{{.TitleCase}}(ctx context.Context, resource *{{.TitleCase}}) (*{{.TitleCase}}, error) {
	var result {{.TitleCase}}
{{- if .HasNamespaceInPath}}
	path := fmt.Sprintf("{{.APIPath}}", resource.Metadata.Namespace)
{{- else}}
	path := "{{.APIPath}}"
	_ = resource.Metadata.Namespace // Namespace not required in API path for this resource
{{- end}}
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}

// Get{{.TitleCase}} retrieves a {{.TitleCase}}
func (c *Client) Get{{.TitleCase}}(ctx context.Context, namespace, name string) (*{{.TitleCase}}, error) {
	var result {{.TitleCase}}
{{- if .HasNamespaceInPath}}
	path := fmt.Sprintf("{{.APIPathItem}}", namespace, name)
{{- else}}
	path := fmt.Sprintf("{{.APIPathItem}}", name)
	_ = namespace // Namespace not required in API path for this resource
{{- end}}
	err := c.Get(ctx, path, &result)
	return &result, err
}

// Update{{.TitleCase}} updates a {{.TitleCase}}
func (c *Client) Update{{.TitleCase}}(ctx context.Context, resource *{{.TitleCase}}) (*{{.TitleCase}}, error) {
	var result {{.TitleCase}}
{{- if .HasNamespaceInPath}}
	path := fmt.Sprintf("{{.APIPathItem}}", resource.Metadata.Namespace, resource.Metadata.Name)
{{- else}}
	path := fmt.Sprintf("{{.APIPathItem}}", resource.Metadata.Name)
	_ = resource.Metadata.Namespace // Namespace not required in API path for this resource
{{- end}}
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}

// Delete{{.TitleCase}} deletes a {{.TitleCase}}
func (c *Client) Delete{{.TitleCase}}(ctx context.Context, namespace, name string) error {
{{- if .HasNamespaceInPath}}
	path := fmt.Sprintf("{{.APIPathItem}}", namespace, name)
{{- else}}
	path := fmt.Sprintf("{{.APIPathItem}}", name)
	_ = namespace // Namespace not required in API path for this resource
{{- end}}
	return c.Delete(ctx, path)
}
`

const dataSourceTemplate = `// Code generated by generate-all-schemas.go. DO NOT EDIT.
// Source: F5 XC OpenAPI specification

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/f5xc/terraform-provider-f5xc/internal/client"
)

var (
	_ datasource.DataSource              = &{{.TitleCase}}DataSource{}
	_ datasource.DataSourceWithConfigure = &{{.TitleCase}}DataSource{}
)

func New{{.TitleCase}}DataSource() datasource.DataSource {
	return &{{.TitleCase}}DataSource{}
}

type {{.TitleCase}}DataSource struct {
	client *client.Client
}

type {{.TitleCase}}DataSourceModel struct {
	ID          types.String ` + "`" + `tfsdk:"id"` + "`" + `
	Name        types.String ` + "`" + `tfsdk:"name"` + "`" + `
	Namespace   types.String ` + "`" + `tfsdk:"namespace"` + "`" + `
	Description types.String ` + "`" + `tfsdk:"description"` + "`" + `
	Labels      types.Map    ` + "`" + `tfsdk:"labels"` + "`" + `
	Annotations types.Map    ` + "`" + `tfsdk:"annotations"` + "`" + `
}

func (d *{{.TitleCase}}DataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_{{.Name}}"
}

func (d *{{.TitleCase}}DataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "{{.Description}}",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Unique identifier for the resource.",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Name of the {{.TitleCase}}.",
				Required:            true,
			},
			"namespace": schema.StringAttribute{
				MarkdownDescription: "Namespace where the {{.TitleCase}} exists.",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description of the {{.TitleCase}}.",
				Computed:            true,
			},
			"labels": schema.MapAttribute{
				MarkdownDescription: "Labels applied to this resource.",
				Computed:            true,
				ElementType:         types.StringType,
			},
			"annotations": schema.MapAttribute{
				MarkdownDescription: "Annotations applied to this resource.",
				Computed:            true,
				ElementType:         types.StringType,
			},
		},
	}
}

func (d *{{.TitleCase}}DataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Data Source Configure Type", "Expected *client.Client")
		return
	}
	d.client = client
}

func (d *{{.TitleCase}}DataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data {{.TitleCase}}DataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resource, err := d.client.Get{{.TitleCase}}(ctx, data.Namespace.ValueString(), data.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read {{.TitleCase}}: %s", err))
		return
	}

	data.ID = types.StringValue(resource.Metadata.Name)
	data.Name = types.StringValue(resource.Metadata.Name)
	data.Namespace = types.StringValue(resource.Metadata.Namespace)
	if resource.Metadata.Description != "" {
		data.Description = types.StringValue(resource.Metadata.Description)
	} else {
		data.Description = types.StringNull()
	}

	// Filter out system-managed labels (ves.io/*) that are injected by the platform
	if len(resource.Metadata.Labels) > 0 {
		filteredLabels := filterSystemLabels(resource.Metadata.Labels)
		if len(filteredLabels) > 0 {
			labels, diags := types.MapValueFrom(ctx, types.StringType, filteredLabels)
			resp.Diagnostics.Append(diags...)
			if !resp.Diagnostics.HasError() {
				data.Labels = labels
			}
		} else {
			data.Labels = types.MapNull(types.StringType)
		}
	} else {
		data.Labels = types.MapNull(types.StringType)
	}

	if len(resource.Metadata.Annotations) > 0 {
		annotations, diags := types.MapValueFrom(ctx, types.StringType, resource.Metadata.Annotations)
		resp.Diagnostics.Append(diags...)
		if !resp.Diagnostics.HasError() {
			data.Annotations = annotations
		}
	} else {
		data.Annotations = types.MapNull(types.StringType)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
`
