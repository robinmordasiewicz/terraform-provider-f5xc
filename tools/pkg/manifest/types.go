// Package manifest provides types and utilities for generating AI-friendly
// constraint manifests from F5 XC OpenAPI specifications.
//
// The manifest format is designed to enable deterministic AI planning by
// providing machine-readable information about:
// - OneOf constraint groups (mutually exclusive fields)
// - Default values for attributes
// - Resource dependencies and ordering
// - Common usage patterns
package manifest

import (
	"time"
)

// AIManifest is the top-level structure for the AI-friendly constraint manifest.
// This manifest enables AI tools to deterministically create valid Terraform
// configurations without trial-and-error iteration.
type AIManifest struct {
	// Version of the manifest format (semver)
	Version string `json:"version"`

	// GeneratedAt is the timestamp when this manifest was generated
	GeneratedAt time.Time `json:"generated_at"`

	// Provider name (e.g., "f5xc")
	Provider string `json:"provider"`

	// ProviderVersion is the version of the provider this manifest was generated from
	ProviderVersion string `json:"provider_version,omitempty"`

	// TotalResources is the count of resources in this manifest
	TotalResources int `json:"total_resources"`

	// Resources maps resource names to their manifest entries
	Resources map[string]*ResourceManifest `json:"resources"`

	// Categories maps category names to resource lists for discovery
	Categories map[string][]string `json:"categories,omitempty"`

	// GlobalHints contains provider-wide AI hints
	GlobalHints *GlobalAIHints `json:"global_hints,omitempty"`
}

// ResourceManifest contains all AI-relevant information about a single resource.
type ResourceManifest struct {
	// Name is the Terraform resource name (e.g., "http_loadbalancer")
	Name string `json:"name"`

	// Category groups related resources (e.g., "Load Balancing", "Security")
	Category string `json:"category"`

	// Description is a human-readable description of the resource
	Description string `json:"description"`

	// APIPath is the F5 XC API endpoint for this resource
	APIPath string `json:"api_path"`

	// RequiresNamespace indicates if the resource requires a namespace
	RequiresNamespace bool `json:"requires_namespace"`

	// Constraints contains all constraint information for deterministic validation
	Constraints *ConstraintManifest `json:"constraints"`

	// Attributes maps attribute names to their manifest entries
	Attributes map[string]*AttributeManifest `json:"attributes,omitempty"`

	// Blocks maps nested block names to their manifest entries
	Blocks map[string]*BlockManifest `json:"blocks,omitempty"`

	// Defaults contains API-discovered default values
	Defaults map[string]interface{} `json:"defaults,omitempty"`

	// AIHints contains AI-specific guidance for this resource
	AIHints *ResourceAIHints `json:"ai_hints,omitempty"`

	// ImportSyntax documents how to import existing resources
	ImportSyntax string `json:"import_syntax,omitempty"`

	// DocsURL is the link to official F5 XC documentation
	DocsURL string `json:"docs_url,omitempty"`
}

// ConstraintManifest contains all constraint information for a resource.
type ConstraintManifest struct {
	// OneOfGroups lists all mutually exclusive field groups
	OneOfGroups []OneOfGroup `json:"one_of_groups,omitempty"`

	// ConflictsWith lists fields that cannot be used together
	ConflictsWith []ConflictConstraint `json:"conflicts_with,omitempty"`

	// RequiredWith lists fields that must appear together
	RequiredWith []RequiredWithConstraint `json:"required_with,omitempty"`

	// AtLeastOneOf lists groups where at least one field must be set
	AtLeastOneOf []AtLeastOneOfConstraint `json:"at_least_one_of,omitempty"`
}

// OneOfGroup represents a group of mutually exclusive fields.
// Exactly one field from this group must be set.
type OneOfGroup struct {
	// GroupID is a unique identifier for this constraint group
	GroupID string `json:"group_id"`

	// Description explains what this choice represents
	Description string `json:"description,omitempty"`

	// Fields lists all field names in this mutually exclusive group
	Fields []string `json:"fields"`

	// Required indicates if one of these fields MUST be set
	Required bool `json:"required"`

	// DefaultChoice is the recommended field to use when not specified
	DefaultChoice string `json:"default_choice,omitempty"`

	// BlockPath indicates the nested block path if this OneOf is within a block
	BlockPath string `json:"block_path,omitempty"`
}

// ConflictConstraint represents fields that cannot be used together.
type ConflictConstraint struct {
	// Field is the field that has conflicts
	Field string `json:"field"`

	// ConflictsWith lists fields that cannot be used with Field
	ConflictsWith []string `json:"conflicts_with"`
}

// RequiredWithConstraint represents fields that must appear together.
type RequiredWithConstraint struct {
	// Field is the field with dependencies
	Field string `json:"field"`

	// RequiredWith lists fields that must also be present
	RequiredWith []string `json:"required_with"`
}

// AtLeastOneOfConstraint represents a group where at least one must be set.
type AtLeastOneOfConstraint struct {
	// GroupID identifies this constraint
	GroupID string `json:"group_id"`

	// Fields lists all fields in this group
	Fields []string `json:"fields"`
}

// AttributeManifest contains information about a single attribute.
type AttributeManifest struct {
	// Name is the attribute name in Terraform schema
	Name string `json:"name"`

	// Type is the Terraform type (string, number, bool, list, map, etc.)
	Type string `json:"type"`

	// ElementType is the type of elements for list/set/map types
	ElementType string `json:"element_type,omitempty"`

	// Description is a human-readable description
	Description string `json:"description,omitempty"`

	// Required indicates if this attribute must be set
	Required bool `json:"required,omitempty"`

	// Optional indicates if this attribute can be set
	Optional bool `json:"optional,omitempty"`

	// Computed indicates if this attribute is computed by the provider
	Computed bool `json:"computed,omitempty"`

	// Sensitive indicates if this attribute contains sensitive data
	Sensitive bool `json:"sensitive,omitempty"`

	// ForceNew indicates if changing this attribute requires recreation
	ForceNew bool `json:"force_new,omitempty"`

	// DefaultValue is the default value if not specified
	DefaultValue interface{} `json:"default_value,omitempty"`

	// OneOfGroup is the group ID if this attribute belongs to a OneOf constraint
	OneOfGroup string `json:"one_of_group,omitempty"`

	// Validation contains validation rules for this attribute
	Validation *ValidationRules `json:"validation,omitempty"`

	// EnumValues lists valid values for enum attributes
	EnumValues []string `json:"enum_values,omitempty"`
}

// BlockManifest contains information about a nested block.
type BlockManifest struct {
	// Name is the block name in Terraform schema
	Name string `json:"name"`

	// Type is the block type (single, list, set, map)
	Type string `json:"type"`

	// Description is a human-readable description
	Description string `json:"description,omitempty"`

	// Required indicates if this block must be present
	Required bool `json:"required,omitempty"`

	// Optional indicates if this block can be present
	Optional bool `json:"optional,omitempty"`

	// MinItems is the minimum number of block instances
	MinItems int `json:"min_items,omitempty"`

	// MaxItems is the maximum number of block instances
	MaxItems int `json:"max_items,omitempty"`

	// OneOfGroup is the group ID if this block belongs to a OneOf constraint
	OneOfGroup string `json:"one_of_group,omitempty"`

	// Attributes maps attribute names within this block
	Attributes map[string]*AttributeManifest `json:"attributes,omitempty"`

	// NestedBlocks maps nested block names within this block
	NestedBlocks map[string]*BlockManifest `json:"nested_blocks,omitempty"`

	// Constraints contains constraints specific to this block
	Constraints *ConstraintManifest `json:"constraints,omitempty"`
}

// ValidationRules contains validation constraints for an attribute.
type ValidationRules struct {
	// Pattern is a regex pattern the value must match
	Pattern string `json:"pattern,omitempty"`

	// MinLength is the minimum string length
	MinLength int `json:"min_length,omitempty"`

	// MaxLength is the maximum string length
	MaxLength int `json:"max_length,omitempty"`

	// Min is the minimum numeric value
	Min *float64 `json:"min,omitempty"`

	// Max is the maximum numeric value
	Max *float64 `json:"max,omitempty"`

	// MinItems is the minimum number of items for lists
	MinItems int `json:"min_items,omitempty"`

	// MaxItems is the maximum number of items for lists
	MaxItems int `json:"max_items,omitempty"`
}

// ResourceAIHints contains AI-specific guidance for a resource.
type ResourceAIHints struct {
	// CommonPatterns lists common configuration combinations
	CommonPatterns []string `json:"common_patterns,omitempty"`

	// RecommendedDefaults lists recommended default choices for OneOf groups
	RecommendedDefaults map[string]string `json:"recommended_defaults,omitempty"`

	// DependencyOrder lists resources that should be created before this one
	DependencyOrder []string `json:"dependency_order,omitempty"`

	// Complexity indicates relative complexity (simple, moderate, complex)
	Complexity string `json:"complexity,omitempty"`

	// UsageNotes contains important notes for using this resource
	UsageNotes []string `json:"usage_notes,omitempty"`

	// RelatedResources lists commonly used related resources
	RelatedResources []string `json:"related_resources,omitempty"`
}

// GlobalAIHints contains provider-wide AI guidance.
type GlobalAIHints struct {
	// CommonWorkflows lists common multi-resource workflows
	CommonWorkflows []WorkflowHint `json:"common_workflows,omitempty"`

	// ResourcePriority lists resources by usage frequency
	ResourcePriority []string `json:"resource_priority,omitempty"`

	// NamingConventions documents naming patterns
	NamingConventions map[string]string `json:"naming_conventions,omitempty"`
}

// WorkflowHint describes a common multi-resource workflow.
type WorkflowHint struct {
	// Name of the workflow
	Name string `json:"name"`

	// Description of what this workflow accomplishes
	Description string `json:"description"`

	// ResourceOrder lists resources in creation order
	ResourceOrder []string `json:"resource_order"`
}
