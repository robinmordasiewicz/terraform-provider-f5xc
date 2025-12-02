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
	XDisplayName         string            `json:"x-displayname"`
	XVesExample          string            `json:"x-ves-example"`
	XVesValidationRules  map[string]string `json:"x-ves-validation-rules"`
	XVesProtoMessage     string            `json:"x-ves-proto-message"`
	AdditionalProperties interface{}       `json:"additionalProperties"`
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
	MaxDepth           int    // Track recursion depth to prevent infinite loops
	IsSpecField        bool   // True if this is a spec field (not metadata)
	JsonName           string // JSON field name from OpenAPI for API marshaling
	GoType             string // Go type for client struct generation
	UseDomainValidator bool   // True if name field should use DomainValidator (for DNS resources)
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
