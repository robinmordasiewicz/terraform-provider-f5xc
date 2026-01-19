// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

// Package openapi provides types and utilities for parsing OpenAPI specifications
// for the F5XC Terraform provider code generation tools.
package openapi

// Spec represents an OpenAPI 3.x specification.
type Spec struct {
	OpenAPI    string                 `json:"openapi"`
	Info       map[string]interface{} `json:"info"`
	Paths      map[string]interface{} `json:"paths"`
	Components Components             `json:"components"`
}

// Components contains the reusable components of an OpenAPI spec.
type Components struct {
	Schemas map[string]Schema `json:"schemas"`
}

// Schema represents a schema definition in an OpenAPI spec.
type Schema struct {
	Type                 string            `json:"type"`
	Description          string            `json:"description"`
	Title                string            `json:"title"`
	Format               string            `json:"format"`
	Enum                 []interface{}     `json:"enum"`
	Default              interface{}       `json:"default"`
	Properties           map[string]Schema `json:"properties"`
	Items                *Schema           `json:"items"`
	Ref                  string            `json:"$ref"`
	Required             []string          `json:"required"`
	AdditionalProperties interface{}       `json:"additionalProperties"`

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
	XF5XCIcon             string   `json:"x-f5xc-icon"`
	XF5XCCLIDomain        string   `json:"x-f5xc-cli-domain"`
	XF5XCConflictsWith    []string `json:"x-f5xc-conflicts-with"`
}

// TerraformAttribute represents an attribute in a Terraform resource schema.
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
	MaxDepth           int      // Track recursion depth to prevent infinite loops
	IsSpecField        bool     // True if this is a spec field (not metadata)
	JsonName           string   // JSON field name from OpenAPI for API marshaling
	GoType             string   // Go type for client struct generation
	UseDomainValidator bool     // True if name field should use DomainValidator (for DNS resources)
	ConflictsWith      []string // Field names that conflict with this attribute (from x-f5xc-conflicts-with)
}

// ResourceTemplate contains data for generating a Terraform resource.
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
}

// GenerationResult tracks the result of generating a resource.
type GenerationResult struct {
	ResourceName string
	Success      bool
	Error        string
	FilePath     string
}

// IsRef returns true if the schema is a reference to another schema.
func (s *Schema) IsRef() bool {
	return s.Ref != ""
}

// IsArray returns true if the schema type is array.
func (s *Schema) IsArray() bool {
	return s.Type == "array"
}

// IsObject returns true if the schema type is object.
func (s *Schema) IsObject() bool {
	return s.Type == "object"
}

// IsRequired checks if a property name is in the required list.
func (s *Schema) IsRequired(propertyName string) bool {
	for _, r := range s.Required {
		if r == propertyName {
			return true
		}
	}
	return false
}

// HasProperties returns true if the schema has properties defined.
func (s *Schema) HasProperties() bool {
	return len(s.Properties) > 0
}

// =============================================================================
// V2 Spec Types - For parsing enriched API specifications from f5xc-api-enriched
// =============================================================================

// Index represents the index.json manifest file in v2 spec structure.
// This file provides metadata about all domain specifications.
type Index struct {
	Version        string           `json:"version"`
	GeneratedAt    string           `json:"generated_at"`
	Specifications []DomainMetadata `json:"specifications"`
}

// DomainMetadata represents metadata about a domain specification file.
// Field names map to the x-f5xc-* extensions in index.json.
type DomainMetadata struct {
	Name              string                      `json:"domain"` // Domain name from "domain" field
	File              string                      `json:"file"`
	Category          string                      `json:"x-f5xc-category"`
	Description       string                      `json:"description"`
	DescriptionShort  string                      `json:"x-f5xc-description-short"`
	DescriptionMedium string                      `json:"x-f5xc-description-medium"`
	Icon              string                      `json:"x-f5xc-icon"`
	RequiresTier      string                      `json:"x-f5xc-requires-tier"`
	Complexity        string                      `json:"x-f5xc-complexity"`
	IsPreview         bool                        `json:"x-f5xc-is-preview"`
	CLIDomain         string                      `json:"x-f5xc-cli-domain"`
	RelatedDomains    []string                    `json:"x-f5xc-related-domains"`
	UseCases          []string                    `json:"x-f5xc-use-cases"`
	PrimaryResources  []PrimaryResourceMetadata   `json:"x-f5xc-primary-resources"`
}

// PrimaryResourceMetadata represents resource-level metadata from x-f5xc-primary-resources in index.json.
// This is extracted from index.json and provides per-resource tier and dependency info.
type PrimaryResourceMetadata struct {
	Name             string               `json:"name"`
	Description      string               `json:"description"`
	DescriptionShort string               `json:"description_short"`
	Tier             string               `json:"tier"`
	Icon             string               `json:"icon"`
	Category         string               `json:"category"`
	SupportsLogs     bool                 `json:"supports_logs"`
	SupportsMetrics  bool                 `json:"supports_metrics"`
	Dependencies     ResourceDependencies `json:"dependencies"`
	RelationshipHints []string            `json:"relationship_hints"`
}

// ResourceDependencies represents the dependencies of a resource.
type ResourceDependencies struct {
	Required []string `json:"required"`
	Optional []string `json:"optional"`
}

// ResourceMetadata represents metadata about a resource within a domain.
// Used by ExtractResourcesFromDomain for processing.
type ResourceMetadata struct {
	Name                 string   `json:"name"`
	Description          string   `json:"description"`
	APIPath              string   `json:"api_path"`
	RequiresTier         string   `json:"requires_tier"`
	Complexity           string   `json:"complexity"`
	Dependencies         []string `json:"dependencies"`
	MinimumConfiguration string   `json:"minimum_configuration"`
}

// DomainSpec represents a parsed domain specification file (v2 format).
// Each domain file contains multiple related resources.
type DomainSpec struct {
	OpenAPI    string                 `json:"openapi"`
	Info       DomainInfo             `json:"info"`
	Paths      map[string]interface{} `json:"paths"`
	Components Components             `json:"components"`

	// Domain-level enrichment metadata
	XF5XCCategory       string   `json:"x-f5xc-category"`
	XF5XCRequiresTier   string   `json:"x-f5xc-requires-tier"`
	XF5XCComplexity     string   `json:"x-f5xc-complexity"`
	XF5XCIsPreview      bool     `json:"x-f5xc-is-preview"`
	XF5XCRelatedDomains []string `json:"x-f5xc-related-domains"`
	XF5XCUseCases       []string `json:"x-f5xc-use-cases"`
}

// DomainInfo represents the info section of a domain spec.
type DomainInfo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Version     string `json:"version"`

	// Enrichment extensions at info level
	XF5XCDescriptionShort  string `json:"x-f5xc-description-short"`
	XF5XCDescriptionMedium string `json:"x-f5xc-description-medium"`
	XF5XCIcon              string `json:"x-f5xc-icon"`
	XF5XCLogoSVG           string `json:"x-f5xc-logo-svg"`
}

// SpecVersion represents the detected specification version.
type SpecVersion string

const (
	// SpecVersionV2 represents the v2 spec format (38 domain-organized files).
	SpecVersionV2 SpecVersion = "v2"
	// SpecVersionUnknown represents an unrecognized spec format.
	SpecVersionUnknown SpecVersion = "unknown"
)

// V2Categories maps domain categories from v2 specs.
var V2Categories = map[string]string{
	"security":       "Security",
	"networking":     "Networking",
	"infrastructure": "Infrastructure",
	"platform":       "Platform",
	"operations":     "Operations",
	"ai":             "AI Services",
}
