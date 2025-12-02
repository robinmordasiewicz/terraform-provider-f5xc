//go:build ignore
// +build ignore

// generate-comprehensive-schema.go - Comprehensive Terraform schema generator from OpenAPI specs
// This tool generates complete Terraform Provider Framework resources with full attribute schemas,
// nested blocks, OneOf handling, descriptions, and validation rules.
//
// Usage: go run tools/generate-comprehensive-schema.go [resource-name]
// Example: go run tools/generate-comprehensive-schema.go securemesh_site_v2

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"text/template"

	"github.com/f5xc/terraform-provider-f5xc/tools/pkg/naming"
)

// OpenAPI3Spec represents an OpenAPI 3.x specification
type OpenAPI3Spec struct {
	OpenAPI    string                 `json:"openapi"`
	Info       map[string]interface{} `json:"info"`
	Paths      map[string]interface{} `json:"paths"`
	Components Components             `json:"components"`
}

// Components contains the schemas section
type Components struct {
	Schemas map[string]SchemaDefinition `json:"schemas"`
}

// SchemaDefinition represents a schema in OpenAPI
type SchemaDefinition struct {
	Type                 string                      `json:"type"`
	Description          string                      `json:"description"`
	Title                string                      `json:"title"`
	Format               string                      `json:"format"`
	Enum                 []interface{}               `json:"enum"`
	Properties           map[string]SchemaDefinition `json:"properties"`
	Items                *SchemaDefinition           `json:"items"`
	Ref                  string                      `json:"$ref"`
	Required             []string                    `json:"required"`
	XDisplayName         string                      `json:"x-displayname"`
	XVesExample          string                      `json:"x-ves-example"`
	XVesValidationRules  map[string]string           `json:"x-ves-validation-rules"`
	XVesProtoMessage     string                      `json:"x-ves-proto-message"`
	AdditionalProperties interface{}                 `json:"additionalProperties"`
	// OneOf field annotations (F5 XC specific)
	XVesOneOfFields map[string]string `json:"-"` // Populated during parsing
}

// TerraformAttribute represents a Terraform schema attribute
type TerraformAttribute struct {
	Name             string
	GoName           string
	TfsdkTag         string
	Type             string // string, int64, bool, list, set, map, object
	ElementType      string // for list/set/map
	Description      string
	Required         bool
	Optional         bool
	Computed         bool
	Sensitive        bool
	NestedAttributes []TerraformAttribute
	NestedBlockType  string // single, list, set
	IsBlock          bool   // true for nested blocks
	OneOfGroup       string // group name if part of oneof
	ValidationMin    string
	ValidationMax    string
	PlanModifier     string // RequiresReplace, UseStateForUnknown
}

// ResourceTemplate contains data for generating a resource
type ResourceTemplate struct {
	Name               string // snake_case
	TitleCase          string // TitleCase
	APIPath            string
	APIPathPlural      string
	Description        string
	Attributes         []TerraformAttribute
	OneOfGroups        map[string][]string // group name -> field names
	HasComplexSpec     bool
	RequiredAttributes []string
	OptionalAttributes []string
	ComputedAttributes []string
}

// Global spec cache
var specCache = make(map[string]*OpenAPI3Spec)
var schemaCache = make(map[string]SchemaDefinition)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run generate-comprehensive-schema.go <resource-name>")
		fmt.Println("Example: go run generate-comprehensive-schema.go securemesh_site_v2")
		os.Exit(1)
	}

	resourceName := os.Args[1]
	fmt.Printf("🔨 Generating comprehensive schema for: %s\n", resourceName)
	fmt.Println(strings.Repeat("=", 60))

	// Find the spec file for this resource
	specFile, err := findSpecFile(resourceName)
	if err != nil {
		fmt.Printf("❌ Error finding spec file: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("📄 Found spec file: %s\n", filepath.Base(specFile))

	// Parse the spec
	spec, err := parseOpenAPISpec(specFile)
	if err != nil {
		fmt.Printf("❌ Error parsing spec: %v\n", err)
		os.Exit(1)
	}

	// Extract the resource schema
	resource, err := extractResourceSchema(spec, resourceName)
	if err != nil {
		fmt.Printf("❌ Error extracting schema: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("✅ Extracted %d attributes\n", len(resource.Attributes))
	fmt.Printf("   - Required: %d\n", len(resource.RequiredAttributes))
	fmt.Printf("   - Optional: %d\n", len(resource.OptionalAttributes))
	fmt.Printf("   - Computed: %d\n", len(resource.ComputedAttributes))
	fmt.Printf("   - OneOf Groups: %d\n", len(resource.OneOfGroups))

	// Generate the resource file
	if err := generateResourceFile(resource); err != nil {
		fmt.Printf("❌ Error generating resource file: %v\n", err)
		os.Exit(1)
	}

	// Generate the client types
	if err := generateClientTypes(resource); err != nil {
		fmt.Printf("❌ Error generating client types: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\n🎉 Generation complete!")
}

func findSpecFile(resourceName string) (string, error) {
	// Try different patterns for the spec file
	patterns := []string{
		fmt.Sprintf("/tmp/docs-cloud-f5-com.*.schema.views.%s.ves-swagger.json", resourceName),
		fmt.Sprintf("/tmp/docs-cloud-f5-com.*.schema.%s.ves-swagger.json", resourceName),
		fmt.Sprintf("/tmp/docs-cloud-f5-com.*%s*.ves-swagger.json", resourceName),
	}

	for _, pattern := range patterns {
		matches, err := filepath.Glob(pattern)
		if err != nil {
			continue
		}
		if len(matches) > 0 {
			return matches[0], nil
		}
	}

	return "", fmt.Errorf("no spec file found for resource: %s", resourceName)
}

func parseOpenAPISpec(specFile string) (*OpenAPI3Spec, error) {
	if cached, ok := specCache[specFile]; ok {
		return cached, nil
	}

	data, err := os.ReadFile(specFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read spec file: %w", err)
	}

	var spec OpenAPI3Spec
	if err := json.Unmarshal(data, &spec); err != nil {
		return nil, fmt.Errorf("failed to parse spec: %w", err)
	}

	// Cache all schemas for reference resolution
	for name, schema := range spec.Components.Schemas {
		schemaCache[name] = schema
	}

	specCache[specFile] = &spec
	return &spec, nil
}

func extractResourceSchema(spec *OpenAPI3Spec, resourceName string) (*ResourceTemplate, error) {
	// Find the CreateSpecType schema
	createSpecKey := ""
	for key := range spec.Components.Schemas {
		if strings.Contains(strings.ToLower(key), resourceName) &&
			strings.Contains(strings.ToLower(key), "createspectype") {
			createSpecKey = key
			break
		}
	}

	if createSpecKey == "" {
		return nil, fmt.Errorf("CreateSpecType schema not found for %s", resourceName)
	}

	createSpec := spec.Components.Schemas[createSpecKey]
	fmt.Printf("📋 Using schema: %s\n", createSpecKey)

	// Extract OneOf groups from x-ves-oneof-field annotations
	oneOfGroups := make(map[string][]string)
	rawJSON, _ := json.Marshal(spec.Components.Schemas[createSpecKey])
	var rawSchema map[string]interface{}
	json.Unmarshal(rawJSON, &rawSchema)

	for key, value := range rawSchema {
		if strings.HasPrefix(key, "x-ves-oneof-field-") {
			groupName := strings.TrimPrefix(key, "x-ves-oneof-field-")
			if strVal, ok := value.(string); ok {
				// Parse the JSON array format: "[\"field1\",\"field2\"]"
				strVal = strings.Trim(strVal, "[]")
				fields := strings.Split(strVal, ",")
				for i, f := range fields {
					fields[i] = strings.Trim(strings.TrimSpace(f), "\"")
				}
				oneOfGroups[groupName] = fields
			}
		}
	}

	// Convert properties to Terraform attributes
	attributes := []TerraformAttribute{}
	requiredSet := make(map[string]bool)
	for _, r := range createSpec.Required {
		requiredSet[r] = true
	}

	// Determine which OneOf group each field belongs to
	fieldToOneOf := make(map[string]string)
	for groupName, fields := range oneOfGroups {
		for _, field := range fields {
			fieldToOneOf[field] = groupName
		}
	}

	for propName, propSchema := range createSpec.Properties {
		attr := convertToTerraformAttribute(propName, propSchema, requiredSet[propName], fieldToOneOf[propName], spec)
		attributes = append(attributes, attr)
	}

	// Sort attributes: required first, then optional, then computed
	sort.Slice(attributes, func(i, j int) bool {
		if attributes[i].Required != attributes[j].Required {
			return attributes[i].Required
		}
		if attributes[i].Optional != attributes[j].Optional {
			return attributes[i].Optional
		}
		return attributes[i].Name < attributes[j].Name
	})

	// Categorize attributes
	var required, optional, computed []string
	for _, attr := range attributes {
		if attr.Required {
			required = append(required, attr.Name)
		} else if attr.Optional {
			optional = append(optional, attr.Name)
		} else if attr.Computed {
			computed = append(computed, attr.Name)
		}
	}

	// Add standard metadata attributes
	standardAttrs := []TerraformAttribute{
		{
			Name:         "name",
			GoName:       "Name",
			TfsdkTag:     "name",
			Type:         "string",
			Description:  fmt.Sprintf("Name of the %s. Must be unique within the namespace.", toTitleCase(resourceName)),
			Required:     true,
			PlanModifier: "RequiresReplace",
		},
		{
			Name:         "namespace",
			GoName:       "Namespace",
			TfsdkTag:     "namespace",
			Type:         "string",
			Description:  fmt.Sprintf("Namespace where the %s will be created.", toTitleCase(resourceName)),
			Required:     true,
			PlanModifier: "RequiresReplace",
		},
		{
			Name:        "labels",
			GoName:      "Labels",
			TfsdkTag:    "labels",
			Type:        "map",
			ElementType: "string",
			Description: "Labels to apply to this resource.",
			Optional:    true,
		},
		{
			Name:        "annotations",
			GoName:      "Annotations",
			TfsdkTag:    "annotations",
			Type:        "map",
			ElementType: "string",
			Description: "Annotations to apply to this resource.",
			Optional:    true,
		},
		{
			Name:         "id",
			GoName:       "ID",
			TfsdkTag:     "id",
			Type:         "string",
			Description:  "Unique identifier for the resource.",
			Computed:     true,
			PlanModifier: "UseStateForUnknown",
		},
	}

	// Prepend standard attributes
	attributes = append(standardAttrs, attributes...)

	return &ResourceTemplate{
		Name:               resourceName,
		TitleCase:          toTitleCase(resourceName),
		APIPath:            fmt.Sprintf("/api/config/namespaces/%%s/%ss", resourceName),
		APIPathPlural:      resourceName + "s",
		Description:        createSpec.Description,
		Attributes:         attributes,
		OneOfGroups:        oneOfGroups,
		HasComplexSpec:     len(attributes) > 10,
		RequiredAttributes: required,
		OptionalAttributes: optional,
		ComputedAttributes: computed,
	}, nil
}

func convertToTerraformAttribute(name string, schema SchemaDefinition, required bool, oneOfGroup string, spec *OpenAPI3Spec) TerraformAttribute {
	// Resolve $ref if present
	if schema.Ref != "" {
		schema = resolveRef(schema.Ref, spec)
	}

	attr := TerraformAttribute{
		Name:       name,
		GoName:     toTitleCase(name),
		TfsdkTag:   name,
		Required:   required,
		Optional:   !required,
		OneOfGroup: oneOfGroup,
	}

	// Clean up description
	description := schema.Description
	if schema.XDisplayName != "" {
		description = schema.XDisplayName + ". " + description
	}
	// Remove example annotations and extra whitespace
	description = regexp.MustCompile(`\s*Example:.*`).ReplaceAllString(description, "")
	description = regexp.MustCompile(`\s*Validation Rules:.*`).ReplaceAllString(description, "")
	// Replace newlines with spaces for single-line strings in Go code
	description = regexp.MustCompile(`[\n\r]+`).ReplaceAllString(description, " ")
	// Escape quotes
	description = strings.ReplaceAll(description, `"`, `\"`)
	description = strings.TrimSpace(description)
	if description == "" {
		description = fmt.Sprintf("Configuration for %s.", name)
	}
	attr.Description = description

	// Extract validation rules
	if schema.XVesValidationRules != nil {
		for rule, value := range schema.XVesValidationRules {
			if strings.Contains(rule, "gte") || strings.Contains(rule, "min") {
				attr.ValidationMin = value
			}
			if strings.Contains(rule, "lte") || strings.Contains(rule, "max") {
				attr.ValidationMax = value
			}
		}
	}

	// Determine type
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
				// It's a nested block
				attr.IsBlock = true
				attr.NestedBlockType = "list"
				attr.NestedAttributes = convertNestedProperties(itemSchema, spec)
			} else {
				attr.ElementType = mapSchemaTypeToTerraform(itemSchema.Type)
			}
		}
	case "object":
		if len(schema.Properties) > 0 {
			// It's a nested block
			attr.Type = "object"
			attr.IsBlock = true
			attr.NestedBlockType = "single"
			attr.NestedAttributes = convertNestedProperties(schema, spec)
		} else if schema.AdditionalProperties != nil {
			// It's a map
			attr.Type = "map"
			attr.ElementType = "string"
		} else {
			// Generic object, treat as single nested block
			attr.Type = "object"
			attr.IsBlock = true
			attr.NestedBlockType = "single"
		}
	default:
		// If no type but has properties, it's an object
		if len(schema.Properties) > 0 {
			attr.Type = "object"
			attr.IsBlock = true
			attr.NestedBlockType = "single"
			attr.NestedAttributes = convertNestedProperties(schema, spec)
		} else {
			attr.Type = "string" // Default fallback
		}
	}

	return attr
}

func convertNestedProperties(schema SchemaDefinition, spec *OpenAPI3Spec) []TerraformAttribute {
	attrs := []TerraformAttribute{}
	requiredSet := make(map[string]bool)
	for _, r := range schema.Required {
		requiredSet[r] = true
	}

	for propName, propSchema := range schema.Properties {
		attr := convertToTerraformAttribute(propName, propSchema, requiredSet[propName], "", spec)
		attrs = append(attrs, attr)
	}

	// Sort by name
	sort.Slice(attrs, func(i, j int) bool {
		return attrs[i].Name < attrs[j].Name
	})

	return attrs
}

func resolveRef(ref string, spec *OpenAPI3Spec) SchemaDefinition {
	// Extract schema name from ref like "#/components/schemas/SchemaName"
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

func mapSchemaTypeToTerraform(schemaType string) string {
	switch schemaType {
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

func generateResourceFile(resource *ResourceTemplate) error {
	outputPath := fmt.Sprintf("internal/provider/%s_resource_generated.go", resource.Name)

	tmpl, err := template.New("resource").Funcs(template.FuncMap{
		"lower": strings.ToLower,
		"title": toTitleCase,
	}).Parse(resourceFileTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	f, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer f.Close()

	if err := tmpl.Execute(f, resource); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	fmt.Printf("📝 Generated: %s\n", outputPath)
	return nil
}

func generateClientTypes(resource *ResourceTemplate) error {
	outputPath := fmt.Sprintf("internal/client/%s_types_generated.go", resource.Name)

	tmpl, err := template.New("client").Parse(clientTypesTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	f, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer f.Close()

	if err := tmpl.Execute(f, resource); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	fmt.Printf("📝 Generated: %s\n", outputPath)
	return nil
}

// toTitleCase wraps naming.ToResourceTypeName for backward compatibility.
func toTitleCase(s string) string {
	return naming.ToResourceTypeName(s)
}

// Resource file template - generates comprehensive Terraform resource
const resourceFileTemplate = `// Code generated by generate-comprehensive-schema.go. DO NOT EDIT.
// Source: OpenAPI specification for {{.Name}}

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/f5xc/terraform-provider-f5xc/internal/client"
)

var (
	_ resource.Resource                = &{{.TitleCase}}GeneratedResource{}
	_ resource.ResourceWithConfigure   = &{{.TitleCase}}GeneratedResource{}
	_ resource.ResourceWithImportState = &{{.TitleCase}}GeneratedResource{}
)

func New{{.TitleCase}}GeneratedResource() resource.Resource {
	return &{{.TitleCase}}GeneratedResource{}
}

type {{.TitleCase}}GeneratedResource struct {
	client *client.Client
}

type {{.TitleCase}}GeneratedResourceModel struct {
{{- range .Attributes}}
{{- if not .IsBlock}}
	{{.GoName}} types.{{if eq .Type "string"}}String{{else if eq .Type "int64"}}Int64{{else if eq .Type "bool"}}Bool{{else if eq .Type "map"}}Map{{else if eq .Type "list"}}List{{else}}String{{end}} ` + "`" + `tfsdk:"{{.TfsdkTag}}"` + "`" + `
{{- end}}
{{- end}}
}

func (r *{{.TitleCase}}GeneratedResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_{{.Name}}"
}

func (r *{{.TitleCase}}GeneratedResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
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
			},
{{- end}}
{{- end}}
		},
		Blocks: map[string]schema.Block{
{{- range .Attributes}}
{{- if .IsBlock}}
			"{{.TfsdkTag}}": schema.{{if eq .NestedBlockType "single"}}SingleNestedBlock{{else if eq .NestedBlockType "list"}}ListNestedBlock{{else}}SingleNestedBlock{{end}}{
				MarkdownDescription: "{{.Description}}",
{{- if eq .NestedBlockType "list"}}
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						// TODO: Add nested attributes from schema
					},
				},
{{- end}}
			},
{{- end}}
{{- end}}
		},
	}
}

func (r *{{.TitleCase}}GeneratedResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T", req.ProviderData),
		)
		return
	}
	r.client = client
}

func (r *{{.TitleCase}}GeneratedResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data {{.TitleCase}}GeneratedResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resource := &client.{{.TitleCase}}Generated{
		Metadata: client.Metadata{
			Name:      data.Name.ValueString(),
			Namespace: data.Namespace.ValueString(),
		},
		Spec: client.{{.TitleCase}}GeneratedSpec{},
	}

	if !data.Labels.IsNull() {
		labels := make(map[string]string)
		resp.Diagnostics.Append(data.Labels.ElementsAs(ctx, &labels, false)...)
		if resp.Diagnostics.HasError() {
			return
		}
		resource.Metadata.Labels = labels
	}

	if !data.Annotations.IsNull() {
		annotations := make(map[string]string)
		resp.Diagnostics.Append(data.Annotations.ElementsAs(ctx, &annotations, false)...)
		if resp.Diagnostics.HasError() {
			return
		}
		resource.Metadata.Annotations = annotations
	}

	created, err := r.client.Create{{.TitleCase}}Generated(ctx, resource)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create {{.TitleCase}}: %s", err))
		return
	}

	data.ID = types.StringValue(created.Metadata.Name)
	tflog.Trace(ctx, "created {{.TitleCase}} resource")
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *{{.TitleCase}}GeneratedResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data {{.TitleCase}}GeneratedResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resource, err := r.client.Get{{.TitleCase}}Generated(ctx, data.Namespace.ValueString(), data.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read {{.TitleCase}}: %s", err))
		return
	}

	data.ID = types.StringValue(resource.Metadata.Name)
	data.Name = types.StringValue(resource.Metadata.Name)
	data.Namespace = types.StringValue(resource.Metadata.Namespace)

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

func (r *{{.TitleCase}}GeneratedResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data {{.TitleCase}}GeneratedResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resource := &client.{{.TitleCase}}Generated{
		Metadata: client.Metadata{
			Name:      data.Name.ValueString(),
			Namespace: data.Namespace.ValueString(),
		},
		Spec: client.{{.TitleCase}}GeneratedSpec{},
	}

	if !data.Labels.IsNull() {
		labels := make(map[string]string)
		resp.Diagnostics.Append(data.Labels.ElementsAs(ctx, &labels, false)...)
		if resp.Diagnostics.HasError() {
			return
		}
		resource.Metadata.Labels = labels
	}

	if !data.Annotations.IsNull() {
		annotations := make(map[string]string)
		resp.Diagnostics.Append(data.Annotations.ElementsAs(ctx, &annotations, false)...)
		if resp.Diagnostics.HasError() {
			return
		}
		resource.Metadata.Annotations = annotations
	}

	updated, err := r.client.Update{{.TitleCase}}Generated(ctx, resource)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update {{.TitleCase}}: %s", err))
		return
	}
	_ = updated // Suppress unused variable warning

	// Use plan data for ID since API response may not include metadata.name
	data.ID = types.StringValue(data.Name.ValueString())
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *{{.TitleCase}}GeneratedResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data {{.TitleCase}}GeneratedResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.Delete{{.TitleCase}}Generated(ctx, data.Namespace.ValueString(), data.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete {{.TitleCase}}: %s", err))
		return
	}
}

func (r *{{.TitleCase}}GeneratedResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
`

// Client types template
const clientTypesTemplate = `// Code generated by generate-comprehensive-schema.go. DO NOT EDIT.
// Source: OpenAPI specification for {{.Name}}

package client

import (
	"context"
	"fmt"
)

// {{.TitleCase}}Generated represents a F5XC {{.TitleCase}}
type {{.TitleCase}}Generated struct {
	Metadata Metadata                 ` + "`" + `json:"metadata"` + "`" + `
	Spec     {{.TitleCase}}GeneratedSpec ` + "`" + `json:"spec"` + "`" + `
}

// {{.TitleCase}}GeneratedSpec defines the specification for {{.TitleCase}}
type {{.TitleCase}}GeneratedSpec struct {
	// TODO: Add full spec fields based on OpenAPI schema
	Description string ` + "`" + `json:"description,omitempty"` + "`" + `
}

// Create{{.TitleCase}}Generated creates a new {{.TitleCase}}
func (c *Client) Create{{.TitleCase}}Generated(ctx context.Context, resource *{{.TitleCase}}Generated) (*{{.TitleCase}}Generated, error) {
	var result {{.TitleCase}}Generated
	path := fmt.Sprintf("{{.APIPath}}", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}

// Get{{.TitleCase}}Generated retrieves a {{.TitleCase}}
func (c *Client) Get{{.TitleCase}}Generated(ctx context.Context, namespace, name string) (*{{.TitleCase}}Generated, error) {
	var result {{.TitleCase}}Generated
	path := fmt.Sprintf("{{.APIPath}}/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}

// Update{{.TitleCase}}Generated updates a {{.TitleCase}}
func (c *Client) Update{{.TitleCase}}Generated(ctx context.Context, resource *{{.TitleCase}}Generated) (*{{.TitleCase}}Generated, error) {
	var result {{.TitleCase}}Generated
	path := fmt.Sprintf("{{.APIPath}}/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}

// Delete{{.TitleCase}}Generated deletes a {{.TitleCase}}
func (c *Client) Delete{{.TitleCase}}Generated(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("{{.APIPath}}/%s", namespace, name)
	return c.Delete(ctx, path)
}
`
