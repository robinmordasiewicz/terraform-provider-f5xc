//go:build ignore
// +build ignore

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
)

// OpenAPISpec represents the structure of an OpenAPI specification
type OpenAPISpec struct {
	Paths       map[string]interface{} `json:"paths"`
	Definitions map[string]interface{} `json:"definitions"`
}

// ResourceInfo holds information about a resource to generate
type ResourceInfo struct {
	Name          string   // e.g., "http_loadbalancer"
	TitleCase     string   // e.g., "HTTPLoadBalancer"
	CamelCase     string   // e.g., "HttpLoadbalancer"
	APIPath       string   // e.g., "/api/config/namespaces/{namespace}/http_loadbalancers"
	Attributes    []string // List of attribute names
	SpecFile      string   // OpenAPI spec file name
}

const resourceTemplate = `package provider

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
	_ resource.Resource                = &{{.TitleCase}}Resource{}
	_ resource.ResourceWithConfigure   = &{{.TitleCase}}Resource{}
	_ resource.ResourceWithImportState = &{{.TitleCase}}Resource{}
)

func New{{.TitleCase}}Resource() resource.Resource {
	return &{{.TitleCase}}Resource{}
}

type {{.TitleCase}}Resource struct {
	client *client.Client
}

type {{.TitleCase}}ResourceModel struct {
	Name        types.String ` + "`tfsdk:\"name\"`" + `
	Namespace   types.String ` + "`tfsdk:\"namespace\"`" + `
	Description types.String ` + "`tfsdk:\"description\"`" + `
	Labels      types.Map    ` + "`tfsdk:\"labels\"`" + `
	ID          types.String ` + "`tfsdk:\"id\"`" + `
}

func (r *{{.TitleCase}}Resource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_{{.Name}}"
}

func (r *{{.TitleCase}}Resource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages a {{.TitleCase}} in F5 Distributed Cloud.",
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				MarkdownDescription: "Name of the {{.TitleCase}}.",
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"namespace": schema.StringAttribute{
				MarkdownDescription: "Namespace where the {{.TitleCase}} will be created.",
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description of the {{.TitleCase}}.",
				Optional:            true,
			},
			"labels": schema.MapAttribute{
				MarkdownDescription: "Labels to apply to the {{.TitleCase}}.",
				ElementType:         types.StringType,
				Optional:            true,
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Unique identifier for the {{.TitleCase}}.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
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
			fmt.Sprintf("Expected *client.Client, got: %T", req.ProviderData),
		)
		return
	}
	r.client = client
}

func (r *{{.TitleCase}}Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data {{.TitleCase}}ResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resource := &client.{{.TitleCase}}{
		Metadata: client.Metadata{
			Name:      data.Name.ValueString(),
			Namespace: data.Namespace.ValueString(),
		},
		Spec: client.{{.TitleCase}}Spec{
			Description: data.Description.ValueString(),
		},
	}

	if !data.Labels.IsNull() {
		labels := make(map[string]string)
		resp.Diagnostics.Append(data.Labels.ElementsAs(ctx, &labels, false)...)
		if resp.Diagnostics.HasError() {
			return
		}
		resource.Metadata.Labels = labels
	}

	created, err := r.client.Create{{.TitleCase}}(ctx, resource)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create {{.TitleCase}}: %s", err))
		return
	}

	data.ID = types.StringValue(created.Metadata.Name)
	tflog.Trace(ctx, "created {{.TitleCase}} resource")
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *{{.TitleCase}}Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data {{.TitleCase}}ResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resource, err := r.client.Get{{.TitleCase}}(ctx, data.Namespace.ValueString(), data.Name.ValueString())
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

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *{{.TitleCase}}Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data {{.TitleCase}}ResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resource := &client.{{.TitleCase}}{
		Metadata: client.Metadata{
			Name:      data.Name.ValueString(),
			Namespace: data.Namespace.ValueString(),
		},
		Spec: client.{{.TitleCase}}Spec{
			Description: data.Description.ValueString(),
		},
	}

	if !data.Labels.IsNull() {
		labels := make(map[string]string)
		resp.Diagnostics.Append(data.Labels.ElementsAs(ctx, &labels, false)...)
		if resp.Diagnostics.HasError() {
			return
		}
		resource.Metadata.Labels = labels
	}

	updated, err := r.client.Update{{.TitleCase}}(ctx, resource)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update {{.TitleCase}}: %s", err))
		return
	}

	data.ID = types.StringValue(updated.Metadata.Name)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *{{.TitleCase}}Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data {{.TitleCase}}ResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

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

const clientTypesTemplate = `
// {{.TitleCase}} represents a F5XC {{.TitleCase}}
type {{.TitleCase}} struct {
	Metadata Metadata           ` + "`json:\"metadata\"`" + `
	Spec     {{.TitleCase}}Spec ` + "`json:\"spec\"`" + `
}

// {{.TitleCase}}Spec defines the specification for {{.TitleCase}}
type {{.TitleCase}}Spec struct {
	Description string ` + "`json:\"description,omitempty\"`" + `
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

func main() {
	fmt.Println("🔨 F5XC Terraform Provider Resource Generator")
	fmt.Println(strings.Repeat("=", 50))

	// Scan for OpenAPI spec files
	specFiles, err := filepath.Glob("/tmp/docs-cloud-f5-com.*.ves-swagger.json")
	if err != nil {
		fmt.Printf("Error finding spec files: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("📄 Found %d OpenAPI specification files\n\n", len(specFiles))

	resources := []ResourceInfo{}

	// Parse each spec file
	for _, specFile := range specFiles {
		resource := extractResourceInfo(specFile)
		if resource != nil {
			resources = append(resources, *resource)
		}
	}

	fmt.Printf("✅ Identified %d resources to generate\n\n", len(resources))

	// Generate resources
	for i, res := range resources {
		fmt.Printf("[%d/%d] Generating %s...\n", i+1, len(resources), res.Name)
		generateResource(res)
	}

	fmt.Println("\n🎉 Resource generation complete!")
	fmt.Printf("📊 Generated %d resources successfully\n", len(resources))
}

func extractResourceInfo(specFile string) *ResourceInfo {
	// Read spec file
	data, err := os.ReadFile(specFile)
	if err != nil {
		return nil
	}

	var spec OpenAPISpec
	if err := json.Unmarshal(data, &spec); err != nil {
		return nil
	}

	// Extract resource name from file name
	// Multiple formats:
	// - docs-cloud-f5-com.NNNN.public.ves.io.schema.views.RESOURCE.ves-swagger.json
	// - docs-cloud-f5-com.NNNN.public.ves.io.schema.RESOURCE.ves-swagger.json
	// - docs-cloud-f5-com.NNNN.public.ves.io.schema.SUBTYPE.RESOURCE.ves-swagger.json
	base := filepath.Base(specFile)

	// Try views pattern first
	re := regexp.MustCompile(`\.schema\.views\.([^.]+)\.ves-swagger`)
	matches := re.FindStringSubmatch(base)
	if len(matches) < 2 {
		// Try direct schema pattern: schema.RESOURCE.ves-swagger
		re = regexp.MustCompile(`\.schema\.([^.]+)\.ves-swagger`)
		matches = re.FindStringSubmatch(base)
	}
	if len(matches) < 2 {
		// Try subtype pattern: schema.SUBTYPE.RESOURCE.ves-swagger
		re = regexp.MustCompile(`\.schema\.[^.]+\.([^.]+)\.ves-swagger`)
		matches = re.FindStringSubmatch(base)
	}
	if len(matches) < 2 {
		return nil
	}

	resourceName := matches[1]

	// Skip certain resource names that are internal or not manageable
	skipResources := map[string]bool{
		"object": true, "status": true, "spec": true, "metadata": true,
		"types": true, "common": true, "refs": true, "crudapi": true,
	}
	if skipResources[resourceName] {
		return nil
	}

	// Find API path from paths
	var apiPath string
	for path := range spec.Paths {
		if strings.Contains(path, "/api/config/namespaces/") &&
		   strings.Contains(path, resourceName) &&
		   !strings.Contains(path, "{name}") &&
		   !strings.Contains(path, "{metadata.name}") {
			apiPath = path
			break
		}
	}

	if apiPath == "" {
		return nil
	}

	return &ResourceInfo{
		Name:      resourceName,
		TitleCase: toTitleCase(resourceName),
		CamelCase: toCamelCase(resourceName),
		APIPath:   apiPath,
		SpecFile:  specFile,
	}
}

func generateResource(res ResourceInfo) {
	// Generate resource file
	tmpl := template.Must(template.New("resource").Parse(resourceTemplate))

	resourceFile := fmt.Sprintf("/tmp/terraform-provider-f5xc/internal/provider/%s_resource.go", res.Name)
	f, err := os.Create(resourceFile)
	if err != nil {
		fmt.Printf("  ❌ Error creating resource file: %v\n", err)
		return
	}
	defer f.Close()

	if err := tmpl.Execute(f, res); err != nil {
		fmt.Printf("  ❌ Error generating resource: %v\n", err)
		return
	}

	fmt.Printf("  ✅ Generated %s_resource.go\n", res.Name)
}

func toTitleCase(s string) string {
	// List of acronyms that should be all uppercase
	acronyms := map[string]bool{
		"http": true, "https": true, "dns": true, "tcp": true, "udp": true,
		"tls": true, "ssl": true, "api": true, "url": true, "uri": true,
		"ip": true, "bgp": true, "jwt": true, "acl": true, "waf": true,
		"cdn": true, "aws": true, "gcp": true, "vpc": true, "tgw": true,
		"vnet": true, "ce": true, "re": true, "lb": true, "vip": true,
		"sni": true, "cors": true, "xss": true, "csrf": true, "oidc": true,
		"saml": true, "ssh": true, "nfs": true, "ntp": true, "pem": true,
		"rsa": true, "ecdsa": true, "id": true, "apm": true, "irule": true,
		"tpm": true, "ike": true, "vpn": true, "ha": true, "s2s": true,
		"sli": true, "slo": true, "oci": true, "kvm": true, "nfv": true,
		"crl": true, "ipv6": true, "ipv4": true, "mtls": true, "graphql": true,
	}

	// List of compound words that should have internal capitals
	compounds := map[string]string{
		"loadbalancer": "LoadBalancer",
		"bigip":        "BigIP",
		"websocket":    "WebSocket",
		"fastcgi":      "FastCGI",
	}

	parts := strings.Split(s, "_")
	for i, part := range parts {
		lowerPart := strings.ToLower(part)
		if acronyms[lowerPart] {
			parts[i] = strings.ToUpper(part)
		} else if replacement, ok := compounds[lowerPart]; ok {
			parts[i] = replacement
		} else {
			parts[i] = strings.Title(part)
		}
	}
	return strings.Join(parts, "")
}

func toCamelCase(s string) string {
	title := toTitleCase(s)
	if len(title) == 0 {
		return title
	}
	return strings.ToLower(string(title[0])) + title[1:]
}
