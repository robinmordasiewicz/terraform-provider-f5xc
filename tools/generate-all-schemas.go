//go:build ignore
// +build ignore

// generate-all-schemas.go - Batch generator for all F5 XC Terraform resources
// This tool processes all OpenAPI spec files and generates comprehensive Terraform schemas.
//
// Usage: go run tools/generate-all-schemas.go [--spec-dir=/path/to/specs] [--dry-run]
//
// Environment Variables:
//   F5XC_SPEC_DIR - Directory containing OpenAPI spec files (default: /tmp)

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"text/template"
)

// Configuration
var (
	specDir    string
	dryRun     bool
	outputDir  string
	clientDir  string
	verbose    bool
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
	XDisplayName         string                      `json:"x-displayname"`
	XVesExample          string                      `json:"x-ves-example"`
	XVesValidationRules  map[string]string           `json:"x-ves-validation-rules"`
	XVesProtoMessage     string                      `json:"x-ves-proto-message"`
	AdditionalProperties interface{}                 `json:"additionalProperties"`
}

type TerraformAttribute struct {
	Name             string
	GoName           string
	TfsdkTag         string
	Type             string
	ElementType      string
	Description      string
	Required         bool
	Optional         bool
	Computed         bool
	Sensitive        bool
	NestedAttributes []TerraformAttribute
	NestedBlockType  string
	IsBlock          bool
	OneOfGroup       string
	PlanModifier     string
	MaxDepth         int // Track recursion depth to prevent infinite loops
}

type ResourceTemplate struct {
	Name               string
	TitleCase          string
	APIPath            string
	APIPathPlural      string
	Description        string
	Attributes         []TerraformAttribute
	OneOfGroups        map[string][]string
	HasComplexSpec     bool
	RequiredAttributes []string
	OptionalAttributes []string
	ComputedAttributes []string
	ExampleUsage       string // HCL example for documentation
	APIDocsURL         string // Link to F5 XC API documentation
}

type GenerationResult struct {
	ResourceName string
	Success      bool
	Error        string
	AttrCount    int
	BlockCount   int
}

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
		specDir = "/tmp"
	}

	fmt.Println("🔨 F5XC Terraform Provider - Batch Schema Generator")
	fmt.Println(strings.Repeat("=", 60))
	fmt.Printf("📁 Spec Directory: %s\n", specDir)
	fmt.Printf("📁 Output Directory: %s\n", outputDir)
	if dryRun {
		fmt.Println("🔍 DRY RUN MODE - No files will be written")
	}
	fmt.Println()

	// Find all spec files
	pattern := filepath.Join(specDir, "docs-cloud-f5-com.*.ves-swagger.json")
	specFiles, err := filepath.Glob(pattern)
	if err != nil {
		fmt.Printf("❌ Error finding spec files: %v\n", err)
		os.Exit(1)
	}

	if len(specFiles) == 0 {
		fmt.Printf("❌ No spec files found matching pattern: %s\n", pattern)
		fmt.Println("💡 Tip: Download specs from docs.cloud.f5.com or set F5XC_SPEC_DIR")
		os.Exit(1)
	}

	fmt.Printf("📄 Found %d OpenAPI specification files\n\n", len(specFiles))

	// Process each spec file
	results := []GenerationResult{}
	successCount := 0
	failCount := 0

	for _, specFile := range specFiles {
		result := processSpecFile(specFile)
		results = append(results, result)
		if result.Success {
			successCount++
		} else if result.Error != "" {
			failCount++
		}
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

	fmt.Println("\n🎉 Batch generation complete!")
}

func processSpecFile(specFile string) GenerationResult {
	// Extract resource name from filename
	resourceName := extractResourceName(specFile)
	if resourceName == "" {
		return GenerationResult{ResourceName: filepath.Base(specFile), Success: false}
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

	if verbose {
		fmt.Printf("Processing: %s\n", resourceName)
	}

	// Parse spec
	spec, err := parseOpenAPISpec(specFile)
	if err != nil {
		return GenerationResult{ResourceName: resourceName, Success: false, Error: err.Error()}
	}

	// Extract resource schema
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
	}

	fmt.Printf("✅ %s: %d attrs, %d blocks\n", resourceName, attrCount, blockCount)
	return GenerationResult{
		ResourceName: resourceName,
		Success:      true,
		AttrCount:    attrCount,
		BlockCount:   blockCount,
	}
}

func extractResourceName(specFile string) string {
	base := filepath.Base(specFile)

	// Try views pattern: schema.views.RESOURCE.ves-swagger
	re := regexp.MustCompile(`\.schema\.views\.([^.]+)\.ves-swagger`)
	matches := re.FindStringSubmatch(base)
	if len(matches) >= 2 {
		return matches[1]
	}

	// Try subtype pattern: schema.SUBTYPE.RESOURCE.ves-swagger
	re = regexp.MustCompile(`\.schema\.[^.]+\.([^.]+)\.ves-swagger`)
	matches = re.FindStringSubmatch(base)
	if len(matches) >= 2 {
		return matches[1]
	}

	// Try direct pattern: schema.RESOURCE.ves-swagger
	re = regexp.MustCompile(`\.schema\.([^.]+)\.ves-swagger`)
	matches = re.FindStringSubmatch(base)
	if len(matches) >= 2 {
		return matches[1]
	}

	return ""
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
	fieldIsFirst := make(map[string]bool) // Only first field in each group gets the constraint
	for _, fields := range oneOfGroups {
		// Sort fields to determine which is first
		sortedFields := make([]string, len(fields))
		copy(sortedFields, fields)
		sort.Strings(sortedFields)
		firstField := sortedFields[0]

		for _, field := range fields {
			fieldToOneOf[field] = fields
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
		attr := convertToTerraformAttribute(propName, propSchema, requiredSet[propName], "", spec)
		// Add OneOf constraint hint to description only for the first field in each group
		if len(oneOfFields) > 1 && fieldIsFirst[propName] {
			attr.Description = addOneOfConstraint(attr.Description, oneOfFields)
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

	// Add standard metadata attributes in HashiCorp-compliant order:
	// 1. ID components (name, namespace) - these form the resource ID
	// 2. Other required args alphabetically
	// 3. Optional args alphabetically (annotations, labels)
	// 4. Computed attributes (id first)
	idComponentAttrs := []TerraformAttribute{
		{Name: "name", GoName: "Name", TfsdkTag: "name", Type: "string",
			Description: fmt.Sprintf("Name of the %s. Must be unique within the namespace.", toTitleCase(resourceName)),
			Required: true, PlanModifier: "RequiresReplace"},
		{Name: "namespace", GoName: "Namespace", TfsdkTag: "namespace", Type: "string",
			Description: fmt.Sprintf("Namespace where the %s will be created.", toTitleCase(resourceName)),
			Required: true, PlanModifier: "RequiresReplace"},
	}

	// Optional standard attrs will be sorted with other optionals
	optionalStdAttrs := []TerraformAttribute{
		{Name: "annotations", GoName: "Annotations", TfsdkTag: "annotations", Type: "map", ElementType: "string",
			Description: "Annotations to apply to this resource.", Optional: true},
		{Name: "labels", GoName: "Labels", TfsdkTag: "labels", Type: "map", ElementType: "string",
			Description: "Labels to apply to this resource.", Optional: true},
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
	allOptional := append(optionalStdAttrs, filterOptional(attributes)...)
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

	// Transform raw API description into user-friendly Terraform description
	description := transformResourceDescription(resourceName, createSpec.Description)

	// Generate example usage HCL
	exampleUsage := generateExampleUsage(resourceName, attributes)

	// Generate API docs URL
	apiDocsURL := fmt.Sprintf("https://docs.cloud.f5.com/docs/api/%s", strings.ReplaceAll(resourceName, "_", "-"))

	return &ResourceTemplate{
		Name:          resourceName,
		TitleCase:     toTitleCase(resourceName),
		APIPath:       fmt.Sprintf("/api/config/namespaces/%%s/%ss", resourceName),
		APIPathPlural: resourceName + "s",
		Description:   description,
		Attributes:    attributes,
		OneOfGroups:   make(map[string][]string),
		ExampleUsage:  exampleUsage,
		APIDocsURL:    apiDocsURL,
	}, nil
}

// generateExampleUsage creates a sample HCL configuration for the resource
func generateExampleUsage(resourceName string, attributes []TerraformAttribute) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("resource \"f5xc_%s\" \"example\" {\n", resourceName))
	sb.WriteString("  name      = \"example\"\n")
	sb.WriteString("  namespace = \"system\"\n")
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
const maxNestedDepth = 3

func convertToTerraformAttribute(name string, schema SchemaDefinition, required bool, oneOfGroup string, spec *OpenAPI3Spec) TerraformAttribute {
	return convertToTerraformAttributeWithDepth(name, schema, required, oneOfGroup, spec, 0)
}

func convertToTerraformAttributeWithDepth(name string, schema SchemaDefinition, required bool, oneOfGroup string, spec *OpenAPI3Spec, depth int) TerraformAttribute {
	if schema.Ref != "" {
		schema = resolveRef(schema.Ref, spec)
	}

	// Convert name to valid Terraform attribute name (lowercase with underscores)
	tfsdkName := toSnakeCase(name)

	attr := TerraformAttribute{
		Name:       name,
		GoName:     toTitleCase(name),
		TfsdkTag:   tfsdkName,
		Required:   required,
		Optional:   !required,
		OneOfGroup: oneOfGroup,
		MaxDepth:   depth,
	}

	// Build description
	description := schema.Description
	if schema.XDisplayName != "" {
		description = schema.XDisplayName + ". " + description
	}
	attr.Description = cleanDescription(description)
	if attr.Description == "" {
		attr.Description = fmt.Sprintf("Configuration for %s.", name)
	}

	// Format enum values per HashiCorp standards: "Possible values are `value1`, `value2`"
	if len(schema.Enum) > 0 {
		attr.Description = formatEnumDescription(attr.Description, schema.Enum)
	}

	// Format default values per HashiCorp standards: "Defaults to `value`."
	if schema.Default != nil {
		attr.Description = formatDefaultDescription(attr.Description, schema.Default)
	}

	// Determine type and extract nested attributes
	switch schema.Type {
	case "string":
		attr.Type = "string"
	case "integer", "number":
		attr.Type = "int64"
	case "boolean":
		attr.Type = "bool"
	case "array":
		attr.Type = "list"
		if schema.Items != nil {
			itemSchema := *schema.Items
			if itemSchema.Ref != "" {
				itemSchema = resolveRef(itemSchema.Ref, spec)
			}
			if itemSchema.Type == "object" || len(itemSchema.Properties) > 0 {
				attr.IsBlock = true
				attr.NestedBlockType = "list"
				// Extract nested attributes if within depth limit
				if depth < maxNestedDepth {
					attr.NestedAttributes = extractNestedAttributes(itemSchema, spec, depth+1)
				}
			} else {
				attr.ElementType = mapSchemaType(itemSchema.Type)
			}
		}
	case "object":
		if len(schema.Properties) > 0 {
			attr.Type = "object"
			attr.IsBlock = true
			attr.NestedBlockType = "single"
			// Extract nested attributes if within depth limit
			if depth < maxNestedDepth {
				attr.NestedAttributes = extractNestedAttributes(schema, spec, depth+1)
			}
		} else if schema.AdditionalProperties != nil {
			attr.Type = "map"
			attr.ElementType = "string"
		} else {
			attr.Type = "object"
			attr.IsBlock = true
			attr.NestedBlockType = "single"
		}
	default:
		if len(schema.Properties) > 0 {
			attr.Type = "object"
			attr.IsBlock = true
			attr.NestedBlockType = "single"
			// Extract nested attributes if within depth limit
			if depth < maxNestedDepth {
				attr.NestedAttributes = extractNestedAttributes(schema, spec, depth+1)
			}
		} else {
			attr.Type = "string"
		}
	}

	return attr
}

// extractNestedAttributes extracts attributes from an object schema's properties
func extractNestedAttributes(schema SchemaDefinition, spec *OpenAPI3Spec, depth int) []TerraformAttribute {
	if depth > maxNestedDepth {
		return nil
	}

	requiredSet := make(map[string]bool)
	for _, r := range schema.Required {
		requiredSet[r] = true
	}

	var attrs []TerraformAttribute
	for propName, propSchema := range schema.Properties {
		attr := convertToTerraformAttributeWithDepth(propName, propSchema, requiredSet[propName], "", spec, depth)
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

func cleanDescription(desc string) string {
	// Remove example and validation rules sections
	desc = regexp.MustCompile(`\s*Example:.*`).ReplaceAllString(desc, "")
	desc = regexp.MustCompile(`\s*Validation Rules:.*`).ReplaceAllString(desc, "")
	// Remove ves.io validation annotations (common pattern in F5 XC specs)
	// Pattern: ves.io.schema.rules.xxx.yyy: value or ves.io.schema.xxx: value
	desc = regexp.MustCompile(`\s*ves\.io\.schema[^\s]*:\s*\S+`).ReplaceAllString(desc, "")
	desc = regexp.MustCompile(`\s*ves\.io\.[^\s]*:\s*\[.*?\]`).ReplaceAllString(desc, "")
	// Remove "Required: YES" or "Required: NO" annotations
	desc = regexp.MustCompile(`\s*Required:\s*(YES|NO)\s*`).ReplaceAllString(desc, " ")
	// Remove "Exclusive with [xxx]" patterns
	desc = regexp.MustCompile(`\s*Exclusive with\s*\[[^\]]*\]\s*`).ReplaceAllString(desc, " ")
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
	// Limit description length to prevent overly long strings
	if len(desc) > 500 {
		desc = desc[:497] + "..."
	}
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

	// Skip invalid/placeholder defaults
	invalidDefaults := map[string]bool{
		"INVALID":      true,
		"NONE":         true,
		"UNKNOWN":      true,
		"UNSPECIFIED":  true,
		"0":            true, // Often placeholder
		"false":        true, // Boolean false is often just the default state
	}

	// Check if default contains "INVALID" or similar markers
	upperDefault := strings.ToUpper(defaultStr)
	for invalid := range invalidDefaults {
		if strings.Contains(upperDefault, invalid) {
			return desc
		}
	}

	// Ensure description ends properly before adding default info
	desc = strings.TrimSpace(desc)
	if desc != "" && !strings.HasSuffix(desc, ".") && !strings.HasSuffix(desc, ":") {
		desc += "."
	}

	return fmt.Sprintf("%s Defaults to `%s`.", desc, defaultStr)
}

// formatEnumDescription appends enum values to a description per HashiCorp standards.
// Format: "Possible values are `value1`, `value2`" or "The only possible value is `value1`"
func formatEnumDescription(desc string, enumValues []interface{}) string {
	if len(enumValues) == 0 {
		return desc
	}

	// Convert enum values to strings with backtick formatting
	var formattedValues []string
	for _, v := range enumValues {
		str := fmt.Sprintf("%v", v)
		// Skip empty or very long values
		if str == "" || len(str) > 50 {
			continue
		}
		formattedValues = append(formattedValues, fmt.Sprintf("`%s`", str))
	}

	if len(formattedValues) == 0 {
		return desc
	}

	// Ensure description ends properly before adding enum info
	desc = strings.TrimSpace(desc)
	if desc != "" && !strings.HasSuffix(desc, ".") && !strings.HasSuffix(desc, ":") {
		desc += "."
	}

	// Format based on number of values per HashiCorp standards
	if len(formattedValues) == 1 {
		return fmt.Sprintf("%s The only possible value is %s.", desc, formattedValues[0])
	}

	// Limit to first 10 values to avoid overly long descriptions
	if len(formattedValues) > 10 {
		formattedValues = formattedValues[:10]
		return fmt.Sprintf("%s Possible values include %s, and others.", desc, strings.Join(formattedValues, ", "))
	}

	return fmt.Sprintf("%s Possible values are %s.", desc, strings.Join(formattedValues, ", "))
}

// transformResourceDescription converts technical API descriptions into user-friendly
// Terraform resource descriptions following HashiCorp best practices.
// Pattern: "Manages a [Resource] in F5 Distributed Cloud [for purpose/capability]."
func transformResourceDescription(resourceName, rawDescription string) string {
	titleCase := toTitleCase(resourceName)

	// Clean and normalize the raw description first
	desc := cleanDescription(rawDescription)
	desc = strings.TrimSpace(desc)

	// If empty, use default
	if desc == "" {
		return fmt.Sprintf("Manages a %s resource in F5 Distributed Cloud.", titleCase)
	}

	// Detect and transform common technical description patterns
	lowerDesc := strings.ToLower(desc)

	// Pattern 1: "Shape of the X specification" -> extract X and make user-friendly
	if strings.Contains(lowerDesc, "shape of") {
		return generateCapabilityDescription(resourceName, titleCase, desc)
	}

	// Pattern 2: "X object" or "X configuration" - technical object reference
	if strings.HasSuffix(lowerDesc, " object") || strings.HasSuffix(lowerDesc, " configuration") ||
		strings.HasSuffix(lowerDesc, " spec") || strings.HasSuffix(lowerDesc, " specification") {
		return generateCapabilityDescription(resourceName, titleCase, desc)
	}

	// Pattern 3: Already starts with a verb like "Create", "Configure", "Define"
	// Transform to "Manages" for consistency
	actionVerbs := []string{"create", "configure", "define", "set up", "establish", "provision"}
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
				return fmt.Sprintf("Manages %s in F5 Distributed Cloud.", remainder)
			}
		}
	}

	// Pattern 4: Description is already decent but needs "Manages" prefix
	// If it doesn't start with a verb, add "Manages a X resource" prefix
	if !startsWithVerb(desc) {
		// Use the description as the capability explanation
		capability := extractCapabilityFromDescription(desc)
		if capability != "" {
			return fmt.Sprintf("Manages a %s resource in F5 Distributed Cloud for %s.", titleCase, capability)
		}
		return fmt.Sprintf("Manages a %s resource in F5 Distributed Cloud. %s", titleCase, desc)
	}

	// If description already looks good, just ensure it ends properly
	if !strings.HasSuffix(desc, ".") {
		desc = desc + "."
	}
	return desc
}

// generateCapabilityDescription creates a user-friendly description based on resource name
// and any capability hints from the original description
func generateCapabilityDescription(resourceName, titleCase, rawDesc string) string {
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
		"app_firewall":              "web application firewall (WAF) protection",
		"service_policy":            "defining service-level access control and security policies",
		"network_firewall":          "network-level firewall rules and security controls",
		"rate_limiter":              "protecting services from traffic spikes and DDoS attacks",
		"bot_defense_app_infrastructure": "bot detection and mitigation capabilities",
		"malicious_user_mitigation": "identifying and blocking malicious user behavior",
		"waf_exclusion_policy":      "excluding specific requests from WAF inspection",

		// Networking
		"network_connector":  "connecting networks across sites and cloud providers",
		"virtual_network":    "creating isolated virtual network segments",
		"cloud_connect":      "establishing connectivity to cloud provider networks",
		"cloud_link":         "linking F5 sites to cloud provider infrastructure",
		"bgp":                "BGP routing configuration for network connectivity",
		"ip_prefix_set":      "defining IP address prefix lists for network policies",
		"network_interface":  "configuring network interfaces on sites",

		// DNS
		"dns_zone":              "DNS zone management and configuration",
		"dns_domain":            "DNS domain registration and management",
		"dns_lb_pool":           "DNS load balancer endpoint pools",
		"dns_lb_health_check":   "health monitoring for DNS load balanced endpoints",
		"dns_compliance_checks": "DNS security and compliance verification",

		// Kubernetes
		"k8s_cluster":             "Kubernetes cluster integration and management",
		"virtual_k8s":             "virtual Kubernetes cluster deployment",
		"k8s_cluster_role":        "Kubernetes RBAC cluster role definitions",
		"k8s_cluster_role_binding": "Kubernetes RBAC cluster role bindings",
		"k8s_pod_security_policy": "Kubernetes pod security policy enforcement",
		"container_registry":      "container image registry configuration",

		// Authentication & Secrets
		"authentication":   "authentication methods and identity provider integration",
		"cloud_credentials": "cloud provider credential management for site deployment",
		"api_credential":   "API credential management for service authentication",
		"token":            "API token generation and management",
		"secret_policy":    "secret access policies and controls",

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
		return fmt.Sprintf("Manages a %s resource in F5 Distributed Cloud for %s.", titleCase, capability)
	}

	// Try to extract capability from raw description
	capability := extractCapabilityFromDescription(rawDesc)
	if capability != "" {
		return fmt.Sprintf("Manages a %s resource in F5 Distributed Cloud for %s.", titleCase, capability)
	}

	// Default fallback
	return fmt.Sprintf("Manages a %s resource in F5 Distributed Cloud.", titleCase)
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

// addOneOfConstraint adds a OneOf constraint hint to the description
func addOneOfConstraint(desc string, oneOfFields []string) string {
	if len(oneOfFields) < 2 {
		return desc
	}

	// Format fields with quotes
	quotedFields := make([]string, len(oneOfFields))
	for i, f := range oneOfFields {
		quotedFields[i] = f
	}
	constraint := fmt.Sprintf("[OneOf: %s]", strings.Join(quotedFields, ", "))

	// Add constraint at the beginning of description
	if desc == "" {
		return constraint
	}
	return constraint + " " + desc
}

func toTitleCase(s string) string {
	acronyms := map[string]bool{
		"http": true, "https": true, "api": true, "dns": true, "waf": true,
		"tls": true, "tcp": true, "udp": true, "ssl": true, "aws": true,
		"gcp": true, "vpc": true, "vnet": true, "tgw": true, "ike": true,
		"vpn": true, "ip": true, "id": true, "url": true, "uri": true,
		"ntp": true, "ssh": true, "ha": true, "s2s": true, "sli": true,
		"slo": true, "oci": true, "kvm": true, "nfv": true, "bgp": true,
		"cdn": true, "crl": true, "apm": true, "ipv6": true, "ipv4": true,
		"k8s": true, "acl": true,
	}

	compounds := map[string]string{
		"loadbalancer": "LoadBalancer",
		"bigip":        "BigIP",
	}

	parts := strings.Split(s, "_")
	for i, part := range parts {
		lower := strings.ToLower(part)
		if acronyms[lower] {
			parts[i] = strings.ToUpper(part)
		} else if replacement, ok := compounds[lower]; ok {
			parts[i] = replacement
		} else {
			parts[i] = strings.Title(lower)
		}
	}
	return strings.Join(parts, "")
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

func generateResourceFile(resource *ResourceTemplate) error {
	outputPath := filepath.Join(outputDir, resource.Name+"_resource.go")

	// Create template with custom functions
	funcMap := template.FuncMap{
		"renderNestedAttrs": renderNestedAttributes,
		"renderNestedBlocks": renderNestedBlocks,
	}

	tmpl, err := template.New("resource").Funcs(funcMap).Parse(resourceTemplate)
	if err != nil {
		return fmt.Errorf("template parse error: %w", err)
	}

	f, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("create file error: %w", err)
	}
	defer f.Close()

	return tmpl.Execute(f, resource)
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
		} else if attr.Optional {
			sb.WriteString(fmt.Sprintf("%s\t\tOptional: true,\n", indent))
		} else if attr.Computed {
			sb.WriteString(fmt.Sprintf("%s\t\tComputed: true,\n", indent))
		}

		if attr.Type == "map" || attr.Type == "list" {
			sb.WriteString(fmt.Sprintf("%s\t\tElementType: types.StringType,\n", indent))
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

	tmpl, err := template.New("client").Parse(clientTemplate)
	if err != nil {
		return fmt.Errorf("template parse error: %w", err)
	}

	f, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("create file error: %w", err)
	}
	defer f.Close()

	return tmpl.Execute(f, resource)
}

func generateDataSource(resource *ResourceTemplate) error {
	outputPath := filepath.Join(outputDir, resource.Name+"_data_source.go")

	tmpl, err := template.New("datasource").Parse(dataSourceTemplate)
	if err != nil {
		return fmt.Errorf("template parse error: %w", err)
	}

	f, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("create file error: %w", err)
	}
	defer f.Close()

	return tmpl.Execute(f, resource)
}

func generateCombinedClientTypes(results []GenerationResult) {
	// This is handled by individual client type files
}

func generateProviderRegistration(results []GenerationResult) {
	// Collect successful resources
	var resources []string
	var dataSources []string
	for _, r := range results {
		if r.Success {
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
	APIToken types.String `+"`"+`tfsdk:"api_token"`+"`"+`
	APIURL   types.String `+"`"+`tfsdk:"api_url"`+"`"+`
}

func (p *F5XCProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "f5xc"
	resp.Version = p.version
}

func (p *F5XCProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Terraform provider for F5 Distributed Cloud (F5XC). " +
			"This is an open source community provider built from public F5 API documentation.",
		Attributes: map[string]schema.Attribute{
			"api_token": schema.StringAttribute{
				MarkdownDescription: "F5 Distributed Cloud API Token. Can also be set via F5XC_API_TOKEN environment variable.",
				Optional:            true,
				Sensitive:           true,
			},
			"api_url": schema.StringAttribute{
				MarkdownDescription: "F5 Distributed Cloud API URL. Defaults to https://console.ves.volterra.io/api. " +
					"Can also be set via F5XC_API_URL environment variable.",
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

	// Check for environment variables if not set in configuration
	apiToken := os.Getenv("F5XC_API_TOKEN")
	apiURL := os.Getenv("F5XC_API_URL")

	// Configuration values override environment variables
	if !config.APIToken.IsNull() {
		apiToken = config.APIToken.ValueString()
	}

	if !config.APIURL.IsNull() {
		apiURL = config.APIURL.ValueString()
	}

	// Set default API URL if not provided
	if apiURL == "" {
		apiURL = "https://console.ves.volterra.io/api"
	}

	// Validate that API token is provided
	if apiToken == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("api_token"),
			"Missing F5XC API Token",
			"The provider cannot create the F5XC API client as there is a missing or empty value for the F5XC API token. "+
				"Set the api_token value in the configuration or use the F5XC_API_TOKEN environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
		return
	}

	// Create the F5XC client
	c := client.NewClient(apiURL, apiToken)

	// Make the client available during DataSource and Resource type Configure methods
	resp.DataSourceData = c
	resp.ResourceData = c

	tflog.Info(ctx, "Configured F5XC client", map[string]any{"success": true, "api_url": apiURL})
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

	if err := os.WriteFile(providerPath, []byte(providerContent), 0644); err != nil {
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

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/f5xc/terraform-provider-f5xc/internal/client"
	"github.com/f5xc/terraform-provider-f5xc/internal/privatestate"
	inttimeouts "github.com/f5xc/terraform-provider-f5xc/internal/timeouts"
	"github.com/f5xc/terraform-provider-f5xc/internal/validators"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                   = &{{.TitleCase}}Resource{}
	_ resource.ResourceWithConfigure      = &{{.TitleCase}}Resource{}
	_ resource.ResourceWithImportState    = &{{.TitleCase}}Resource{}
	_ resource.ResourceWithModifyPlan     = &{{.TitleCase}}Resource{}
	_ resource.ResourceWithUpgradeState   = &{{.TitleCase}}Resource{}
	_ resource.ResourceWithValidateConfig = &{{.TitleCase}}Resource{}
)

// {{.Name}}SchemaVersion is the schema version for state upgrades
const {{.Name}}SchemaVersion int64 = 1

func New{{.TitleCase}}Resource() resource.Resource {
	return &{{.TitleCase}}Resource{}
}

type {{.TitleCase}}Resource struct {
	client *client.Client
}

type {{.TitleCase}}ResourceModel struct {
{{- range .Attributes}}
{{- if not .IsBlock}}
	{{.GoName}} types.{{if eq .Type "string"}}String{{else if eq .Type "int64"}}Int64{{else if eq .Type "bool"}}Bool{{else if eq .Type "map"}}Map{{else if eq .Type "list"}}List{{else}}String{{end}} ` + "`" + `tfsdk:"{{.TfsdkTag}}"` + "`" + `
{{- end}}
{{- end}}
	Timeouts timeouts.Value ` + "`" + `tfsdk:"timeouts"` + "`" + `
}

func (r *{{.TitleCase}}Resource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_{{.Name}}"
}

func (r *{{.TitleCase}}Resource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version:             {{.Name}}SchemaVersion,
		MarkdownDescription: "{{.Description}}",
		Attributes: map[string]schema.Attribute{
{{- range .Attributes}}
{{- if not .IsBlock}}
			"{{.TfsdkTag}}": schema.{{if eq .Type "string"}}String{{else if eq .Type "int64"}}Int64{{else if eq .Type "bool"}}Bool{{else if eq .Type "map"}}Map{{else if eq .Type "list"}}List{{else}}String{{end}}Attribute{
				MarkdownDescription: "{{.Description}}",
{{- if .Required}}
				Required: true,
{{- else if .Optional}}
				Optional: true,
{{- else if .Computed}}
				Computed: true,
{{- end}}
{{- if eq .Type "map"}}
				ElementType: types.StringType,
{{- end}}
{{- if eq .Type "list"}}
				ElementType: types.StringType,
{{- end}}
{{- if .PlanModifier}}
				PlanModifiers: []planmodifier.String{
{{- if eq .PlanModifier "RequiresReplace"}}
					stringplanmodifier.RequiresReplace(),
{{- else if eq .PlanModifier "UseStateForUnknown"}}
					stringplanmodifier.UseStateForUnknown(),
{{- end}}
				},
{{- end}}
{{- if eq .TfsdkTag "name"}}
				Validators: []validator.String{
					validators.NameValidator(),
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

// UpgradeState implements resource.ResourceWithUpgradeState
func (r *{{.TitleCase}}Resource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema: &schema.Schema{
				Attributes: map[string]schema.Attribute{
					"name":        schema.StringAttribute{Required: true},
					"namespace":   schema.StringAttribute{Required: true},
					"annotations": schema.MapAttribute{Optional: true, ElementType: types.StringType},
					"labels":      schema.MapAttribute{Optional: true, ElementType: types.StringType},
					"id":          schema.StringAttribute{Computed: true},
				},
			},
			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var priorState struct {
					Name        types.String ` + "`" + `tfsdk:"name"` + "`" + `
					Namespace   types.String ` + "`" + `tfsdk:"namespace"` + "`" + `
					Annotations types.Map    ` + "`" + `tfsdk:"annotations"` + "`" + `
					Labels      types.Map    ` + "`" + `tfsdk:"labels"` + "`" + `
					ID          types.String ` + "`" + `tfsdk:"id"` + "`" + `
				}

				resp.Diagnostics.Append(req.State.Get(ctx, &priorState)...)
				if resp.Diagnostics.HasError() {
					return
				}

				upgradedState := {{.TitleCase}}ResourceModel{
					Name:        priorState.Name,
					Namespace:   priorState.Namespace,
					Annotations: priorState.Annotations,
					Labels:      priorState.Labels,
					ID:          priorState.ID,
					Timeouts:    timeouts.Value{},
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, upgradedState)...)
			},
		},
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

	apiResource := &client.{{.TitleCase}}{
		Metadata: client.Metadata{
			Name:      data.Name.ValueString(),
			Namespace: data.Namespace.ValueString(),
		},
		Spec: client.{{.TitleCase}}Spec{},
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

	created, err := r.client.Create{{.TitleCase}}(ctx, apiResource)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create {{.TitleCase}}: %s", err))
		return
	}

	data.ID = types.StringValue(created.Metadata.Name)

	psd := privatestate.NewPrivateStateData()
	psd.SetUID(created.Metadata.UID)
	resp.Diagnostics.Append(psd.SaveToPrivateState(ctx, resp)...)

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

	psd, psDiags := privatestate.LoadFromPrivateState(ctx, &req)
	resp.Diagnostics.Append(psDiags...)

	apiResource, err := r.client.Get{{.TitleCase}}(ctx, data.Namespace.ValueString(), data.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read {{.TitleCase}}: %s", err))
		return
	}

	if psd != nil && psd.Metadata.UID != "" && apiResource.Metadata.UID != psd.Metadata.UID {
		resp.Diagnostics.AddWarning(
			"Resource Drift Detected",
			"The {{.Name}} may have been recreated outside of Terraform.",
		)
	}

	data.ID = types.StringValue(apiResource.Metadata.Name)
	data.Name = types.StringValue(apiResource.Metadata.Name)
	data.Namespace = types.StringValue(apiResource.Metadata.Namespace)

	if len(apiResource.Metadata.Labels) > 0 {
		labels, diags := types.MapValueFrom(ctx, types.StringType, apiResource.Metadata.Labels)
		resp.Diagnostics.Append(diags...)
		if !resp.Diagnostics.HasError() {
			data.Labels = labels
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

	psd = privatestate.NewPrivateStateData()
	psd.SetUID(apiResource.Metadata.UID)
	resp.Diagnostics.Append(psd.SaveToPrivateState(ctx, resp)...)

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
		Spec: client.{{.TitleCase}}Spec{},
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

	updated, err := r.client.Update{{.TitleCase}}(ctx, apiResource)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update {{.TitleCase}}: %s", err))
		return
	}

	data.ID = types.StringValue(updated.Metadata.Name)

	psd := privatestate.NewPrivateStateData()
	psd.SetUID(updated.Metadata.UID)
	resp.Diagnostics.Append(psd.SaveToPrivateState(ctx, resp)...)

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

	err := r.client.Delete{{.TitleCase}}(ctx, data.Namespace.ValueString(), data.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete {{.TitleCase}}: %s", err))
		return
	}
}

func (r *{{.TitleCase}}Resource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
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
	Metadata Metadata       ` + "`" + `json:"metadata"` + "`" + `
	Spec     {{.TitleCase}}Spec ` + "`" + `json:"spec"` + "`" + `
}

// {{.TitleCase}}Spec defines the specification for {{.TitleCase}}
type {{.TitleCase}}Spec struct {
	Description string ` + "`" + `json:"description,omitempty"` + "`" + `
}

// Create{{.TitleCase}} creates a new {{.TitleCase}}
func (c *Client) Create{{.TitleCase}}(ctx context.Context, resource *{{.TitleCase}}) (*{{.TitleCase}}, error) {
	var result {{.TitleCase}}
	path := fmt.Sprintf("{{.APIPath}}", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}

// Get{{.TitleCase}} retrieves a {{.TitleCase}}
func (c *Client) Get{{.TitleCase}}(ctx context.Context, namespace, name string) (*{{.TitleCase}}, error) {
	var result {{.TitleCase}}
	path := fmt.Sprintf("{{.APIPath}}/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}

// Update{{.TitleCase}} updates a {{.TitleCase}}
func (c *Client) Update{{.TitleCase}}(ctx context.Context, resource *{{.TitleCase}}) (*{{.TitleCase}}, error) {
	var result {{.TitleCase}}
	path := fmt.Sprintf("{{.APIPath}}/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}

// Delete{{.TitleCase}} deletes a {{.TitleCase}}
func (c *Client) Delete{{.TitleCase}}(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("{{.APIPath}}/%s", namespace, name)
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
	data.Description = types.StringValue(resource.Spec.Description)

	if len(resource.Metadata.Labels) > 0 {
		labels, diags := types.MapValueFrom(ctx, types.StringType, resource.Metadata.Labels)
		resp.Diagnostics.Append(diags...)
		if !resp.Diagnostics.HasError() {
			data.Labels = labels
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
