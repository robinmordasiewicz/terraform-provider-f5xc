// Package openapi provides types and utilities for parsing OpenAPI specifications
// for the F5XC Terraform provider code generation tools.
package openapi

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

// ParseFile reads and parses an OpenAPI specification file.
func ParseFile(path string) (*Spec, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading file: %w", err)
	}

	var spec Spec
	if err := json.Unmarshal(data, &spec); err != nil {
		return nil, fmt.Errorf("parsing JSON: %w", err)
	}

	return &spec, nil
}

// ResolveRef resolves a $ref string to the actual schema.
// Only supports local references (#/components/schemas/SchemaName).
func (s *Spec) ResolveRef(ref string) (*Schema, error) {
	if !strings.HasPrefix(ref, "#/components/schemas/") {
		return nil, fmt.Errorf("unsupported ref format: %s", ref)
	}

	schemaName := strings.TrimPrefix(ref, "#/components/schemas/")
	schema, ok := s.Components.Schemas[schemaName]
	if !ok {
		return nil, fmt.Errorf("schema not found: %s", schemaName)
	}

	return &schema, nil
}

// GetRefName extracts the schema name from a $ref string.
func GetRefName(ref string) string {
	if ref == "" {
		return ""
	}
	parts := strings.Split(ref, "/")
	return parts[len(parts)-1]
}

// SpecFileInfo contains information extracted from an OpenAPI spec filename.
type SpecFileInfo struct {
	FilePath     string
	ResourceName string
	SchemaPath   string // e.g., "views.http_loadbalancer"
	URLPath      string // e.g., "views-http-loadbalancer"
}

// ParseSpecFilename extracts resource information from an OpenAPI spec filename.
// Pattern: docs-cloud-f5-com.XXXX.public.ves.io.schema.{path}.ves-swagger.json
func ParseSpecFilename(filename string) (*SpecFileInfo, error) {
	specRegex := regexp.MustCompile(`docs-cloud-f5-com\.\d+\.public\.ves\.io\.schema\.(.+)\.ves-swagger\.json`)

	base := filepath.Base(filename)
	matches := specRegex.FindStringSubmatch(base)
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("filename does not match expected pattern: %s", filename)
	}

	schemaPath := matches[1]
	parts := strings.Split(schemaPath, ".")
	resourceName := parts[len(parts)-1]

	// Convert schema path to URL format: dots -> hyphens, underscores -> hyphens
	urlPath := strings.ReplaceAll(schemaPath, ".", "-")
	urlPath = strings.ReplaceAll(urlPath, "_", "-")

	return &SpecFileInfo{
		FilePath:     filename,
		ResourceName: resourceName,
		SchemaPath:   schemaPath,
		URLPath:      urlPath,
	}, nil
}

// FindSpecFiles finds all OpenAPI spec files in a directory.
func FindSpecFiles(dir string) ([]string, error) {
	pattern := filepath.Join(dir, "*.json")
	return filepath.Glob(pattern)
}

// BuildResourceAPIPathMap scans the OpenAPI spec directory and builds a mapping
// from resource names to their API documentation paths.
// Example: "http_loadbalancer" -> "views-http-loadbalancer"
func BuildResourceAPIPathMap(specDir string) (map[string]string, error) {
	files, err := FindSpecFiles(specDir)
	if err != nil {
		return nil, fmt.Errorf("scanning spec directory: %w", err)
	}

	pathMap := make(map[string]string)
	for _, file := range files {
		info, err := ParseSpecFilename(file)
		if err != nil {
			// Skip files that don't match the pattern
			continue
		}
		pathMap[info.ResourceName] = info.URLPath
	}

	return pathMap, nil
}

// GetAPIDocURL returns the F5 API documentation URL for a resource.
func GetAPIDocURL(resourceName string, pathMap map[string]string) string {
	if urlPath, ok := pathMap[resourceName]; ok {
		return fmt.Sprintf("https://docs.cloud.f5.com/docs-v2/api/%s", urlPath)
	}
	return ""
}

// =============================================================================
// V2 Spec Parser Functions - For parsing enriched API specifications
// =============================================================================

// GetSpecVersion detects whether a spec directory contains v1 or v2 format.
// v2 has index.json and domains/ subdirectory.
// v1 has individual docs-cloud-f5-com.*.ves-swagger.json files.
func GetSpecVersion(specDir string) SpecVersion {
	// Check for v2 indicators
	indexPath := filepath.Join(specDir, "index.json")
	domainsDir := filepath.Join(specDir, "domains")

	indexExists := fileExists(indexPath)
	domainsExists := dirExists(domainsDir)

	if indexExists && domainsExists {
		return SpecVersionV2
	}

	// Check for v1 pattern
	v1Pattern := filepath.Join(specDir, "docs-cloud-f5-com.*.ves-swagger.json")
	v1Files, _ := filepath.Glob(v1Pattern)
	if len(v1Files) > 0 {
		return SpecVersionV1
	}

	return SpecVersionUnknown
}

// IsV2SpecDirectory returns true if the directory contains v2 spec structure.
func IsV2SpecDirectory(specDir string) bool {
	return GetSpecVersion(specDir) == SpecVersionV2
}

// ParseIndex reads and parses the index.json manifest file.
func ParseIndex(path string) (*Index, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading index file: %w", err)
	}

	var index Index
	if err := json.Unmarshal(data, &index); err != nil {
		return nil, fmt.Errorf("parsing index JSON: %w", err)
	}

	return &index, nil
}

// ParseIndexFromDir reads the index.json from a spec directory.
func ParseIndexFromDir(specDir string) (*Index, error) {
	indexPath := filepath.Join(specDir, "index.json")
	return ParseIndex(indexPath)
}

// FindDomainSpecFiles finds all domain specification files in a v2 spec directory.
// Returns paths to all .json files in the domains/ subdirectory.
func FindDomainSpecFiles(specDir string) ([]string, error) {
	domainsDir := filepath.Join(specDir, "domains")
	if !dirExists(domainsDir) {
		return nil, fmt.Errorf("domains directory not found: %s", domainsDir)
	}

	pattern := filepath.Join(domainsDir, "*.json")
	files, err := filepath.Glob(pattern)
	if err != nil {
		return nil, fmt.Errorf("globbing domain files: %w", err)
	}

	return files, nil
}

// ParseDomainSpec reads and parses a domain specification file.
func ParseDomainSpec(path string) (*DomainSpec, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading domain spec: %w", err)
	}

	var spec DomainSpec
	if err := json.Unmarshal(data, &spec); err != nil {
		return nil, fmt.Errorf("parsing domain spec JSON: %w", err)
	}

	return &spec, nil
}

// DomainSpecInfo contains parsed information about a domain and its resources.
type DomainSpecInfo struct {
	DomainName    string
	Category      string
	RequiresTier  string
	Complexity    string
	IsPreview     bool
	Resources     []ExtractedResource
	SourceFile    string
	Spec          *DomainSpec
}

// ExtractedResource contains information about a single resource extracted from a domain.
type ExtractedResource struct {
	Name             string
	SchemaName       string  // Full schema name in components/schemas
	Description      string
	APIPath          string
	RequiresTier     string
	Complexity       string
	Category         string  // Inherited from domain if not specified
	IsPreview        bool    // Inherited from domain if not specified
	DomainName       string  // Parent domain
}

// ExtractResourcesFromDomain parses a domain spec and extracts resource information.
// It identifies CRUD resources by looking for Create/Get/Update/Delete operations.
func ExtractResourcesFromDomain(specPath string) (*DomainSpecInfo, error) {
	spec, err := ParseDomainSpec(specPath)
	if err != nil {
		return nil, err
	}

	domainName := filepath.Base(specPath)
	domainName = strings.TrimSuffix(domainName, ".json")

	info := &DomainSpecInfo{
		DomainName:   domainName,
		Category:     spec.XF5XCCategory,
		RequiresTier: spec.XF5XCRequiresTier,
		Complexity:   spec.XF5XCComplexity,
		IsPreview:    spec.XF5XCIsPreview,
		SourceFile:   specPath,
		Spec:         spec,
		Resources:    []ExtractedResource{},
	}

	// Extract resources from paths by looking for CRUD operations
	// Pattern: /api/config/namespaces/{namespace}/{resource_type} with POST (Create)
	resourcePaths := extractResourcePathsFromPaths(spec.Paths)

	for _, rp := range resourcePaths {
		resource := ExtractedResource{
			Name:         rp.ResourceName,
			SchemaName:   rp.SchemaName,
			APIPath:      rp.APIPath,
			DomainName:   domainName,
			Category:     info.Category,    // Inherit from domain
			IsPreview:    info.IsPreview,   // Inherit from domain
			RequiresTier: info.RequiresTier, // Inherit from domain
			Complexity:   info.Complexity,  // Inherit from domain
		}

		// Try to get description from schema if available
		if schema, ok := spec.Components.Schemas[rp.SchemaName]; ok {
			resource.Description = schema.Description
			// Override with resource-level x-f5xc-* if specified
			if schema.XF5XCRequiresTier != "" {
				resource.RequiresTier = schema.XF5XCRequiresTier
			}
			if schema.XF5XCComplexity != "" {
				resource.Complexity = schema.XF5XCComplexity
			}
			if schema.XF5XCCategory != "" {
				resource.Category = schema.XF5XCCategory
			}
		}

		info.Resources = append(info.Resources, resource)
	}

	return info, nil
}

// resourcePath holds information about a resource path extracted from OpenAPI paths.
type resourcePath struct {
	ResourceName string
	SchemaName   string
	APIPath      string
}

// extractResourcePathsFromPaths analyzes OpenAPI paths to find CRUD resource patterns.
func extractResourcePathsFromPaths(paths map[string]interface{}) []resourcePath {
	var results []resourcePath
	seen := make(map[string]bool)

	// Pattern: /api/config/namespaces/{namespace}/{resource_plural}
	configPathRegex := regexp.MustCompile(`^/api/config/namespaces/\{namespace\}/([a-z_]+s)$`)

	for path := range paths {
		matches := configPathRegex.FindStringSubmatch(path)
		if matches == nil || len(matches) < 2 {
			continue
		}

		resourcePlural := matches[1]
		// Convert plural to singular (simple heuristic)
		resourceName := strings.TrimSuffix(resourcePlural, "s")
		if strings.HasSuffix(resourceName, "ie") {
			resourceName = strings.TrimSuffix(resourceName, "ie") + "y"
		}

		if seen[resourceName] {
			continue
		}
		seen[resourceName] = true

		// Schema name is typically the request/response type
		// Convention: ves.io.schema.{resource}.Object or similar
		schemaName := fmt.Sprintf("ves.io.schema.%s.Object", resourceName)

		results = append(results, resourcePath{
			ResourceName: resourceName,
			SchemaName:   schemaName,
			APIPath:      path,
		})
	}

	return results
}

// GetExampleValue returns the best available example for a schema field.
// Priority: x-f5xc-example > x-ves-example > empty string.
func (s *Schema) GetExampleValue() string {
	if s.XF5XCExample != "" {
		return s.XF5XCExample
	}
	if s.XVesExample != "" {
		return s.XVesExample
	}
	return ""
}

// GetBestDescription returns the best available description for a schema.
// Priority: x-f5xc-description-medium > x-f5xc-description-short > description.
func (s *Schema) GetBestDescription() string {
	if s.XF5XCDescriptionMed != "" {
		return s.XF5XCDescriptionMed
	}
	if s.XF5XCDescriptionShort != "" {
		return s.XF5XCDescriptionShort
	}
	return s.Description
}

// GetCategory returns the category for a schema (from x-f5xc-category).
func (s *Schema) GetCategory() string {
	return s.XF5XCCategory
}

// Helper functions

func fileExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

func dirExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// =============================================================================
// V2 Example Cache - For examples from x-f5xc-example extensions
// =============================================================================

// v2ExampleCache stores examples from v2 spec x-f5xc-example extensions.
// Maps: resourceName -> fieldPath -> exampleValue
// This is populated during spec parsing and used during example generation.
var v2ExampleCache = make(map[string]map[string]string)
var v2ExampleMutex sync.RWMutex

// SetV2Example sets an example for a specific resource field from v2 spec metadata.
// This should be called during spec parsing when x-f5xc-example is found.
func SetV2Example(resourceName, fieldPath, example string) {
	if example == "" {
		return
	}
	v2ExampleMutex.Lock()
	defer v2ExampleMutex.Unlock()
	if v2ExampleCache[resourceName] == nil {
		v2ExampleCache[resourceName] = make(map[string]string)
	}
	v2ExampleCache[resourceName][fieldPath] = example
}

// GetV2Example retrieves the v2 example for a specific resource field if set.
func GetV2Example(resourceName, fieldPath string) (string, bool) {
	v2ExampleMutex.RLock()
	defer v2ExampleMutex.RUnlock()
	if resourceExamples, ok := v2ExampleCache[resourceName]; ok {
		if example, found := resourceExamples[fieldPath]; found {
			return example, true
		}
	}
	return "", false
}

// GetV2ResourceExamples retrieves all v2 examples for a specific resource.
func GetV2ResourceExamples(resourceName string) (map[string]string, bool) {
	v2ExampleMutex.RLock()
	defer v2ExampleMutex.RUnlock()
	examples, ok := v2ExampleCache[resourceName]
	if !ok {
		return nil, false
	}
	// Return a copy to prevent external modification
	result := make(map[string]string, len(examples))
	for k, v := range examples {
		result[k] = v
	}
	return result, true
}

// ClearV2Examples clears all v2 example cache entries (for testing).
func ClearV2Examples() {
	v2ExampleMutex.Lock()
	defer v2ExampleMutex.Unlock()
	v2ExampleCache = make(map[string]map[string]string)
}

// V2ExampleCount returns the number of resources with v2 examples in cache.
func V2ExampleCount() int {
	v2ExampleMutex.RLock()
	defer v2ExampleMutex.RUnlock()
	return len(v2ExampleCache)
}

// V2ExampleFieldCount returns the total number of field examples in cache.
func V2ExampleFieldCount() int {
	v2ExampleMutex.RLock()
	defer v2ExampleMutex.RUnlock()
	count := 0
	for _, fields := range v2ExampleCache {
		count += len(fields)
	}
	return count
}

// =============================================================================
// V2 Ves Example Cache - For examples from original x-ves-example extensions
// =============================================================================

// vesExampleCache stores examples from original x-ves-example extensions.
// Maps: resourceName -> fieldPath -> exampleValue
// This is the second priority after x-f5xc-example.
var vesExampleCache = make(map[string]map[string]string)
var vesExampleMutex sync.RWMutex

// SetVesExample sets an example for a specific resource field from x-ves-example.
func SetVesExample(resourceName, fieldPath, example string) {
	if example == "" {
		return
	}
	vesExampleMutex.Lock()
	defer vesExampleMutex.Unlock()
	if vesExampleCache[resourceName] == nil {
		vesExampleCache[resourceName] = make(map[string]string)
	}
	vesExampleCache[resourceName][fieldPath] = example
}

// GetVesExample retrieves the ves example for a specific resource field if set.
func GetVesExample(resourceName, fieldPath string) (string, bool) {
	vesExampleMutex.RLock()
	defer vesExampleMutex.RUnlock()
	if resourceExamples, ok := vesExampleCache[resourceName]; ok {
		if example, found := resourceExamples[fieldPath]; found {
			return example, true
		}
	}
	return "", false
}

// ClearVesExamples clears all ves example cache entries (for testing).
func ClearVesExamples() {
	vesExampleMutex.Lock()
	defer vesExampleMutex.Unlock()
	vesExampleCache = make(map[string]map[string]string)
}

// =============================================================================
// Unified Example Lookup - Priority-based resolution
// =============================================================================

// GetBestExample returns the best available example for a resource field.
// Priority: x-f5xc-example (v2) > x-ves-example (original) > empty string.
// This function should be used by example generators for priority-based lookup.
func GetBestExample(resourceName, fieldPath string) (string, string) {
	// Priority 1: v2 x-f5xc-example
	if example, ok := GetV2Example(resourceName, fieldPath); ok {
		return example, "v2"
	}
	// Priority 2: original x-ves-example
	if example, ok := GetVesExample(resourceName, fieldPath); ok {
		return example, "ves"
	}
	return "", ""
}

// PopulateExamplesFromSchema extracts and caches examples from a parsed schema.
// This should be called during spec parsing to populate the example caches.
func PopulateExamplesFromSchema(resourceName string, schema *Schema, fieldPath string) {
	if schema == nil {
		return
	}

	// Cache x-f5xc-example if present
	if schema.XF5XCExample != "" {
		SetV2Example(resourceName, fieldPath, schema.XF5XCExample)
	}

	// Cache x-ves-example if present
	if schema.XVesExample != "" {
		SetVesExample(resourceName, fieldPath, schema.XVesExample)
	}

	// Recursively process properties
	for propName, propSchema := range schema.Properties {
		childPath := fieldPath
		if childPath != "" {
			childPath = childPath + "." + propName
		} else {
			childPath = propName
		}
		// Copy to avoid mutation issues
		propSchemaCopy := propSchema
		PopulateExamplesFromSchema(resourceName, &propSchemaCopy, childPath)
	}

	// Process array items
	if schema.Items != nil {
		PopulateExamplesFromSchema(resourceName, schema.Items, fieldPath)
	}
}
