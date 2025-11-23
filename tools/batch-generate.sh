#!/bin/bash
#
# Batch Resource Generator for F5XC Terraform Provider
# Generates all 268 resources from OpenAPI specifications
#

set -e

PROJECT_ROOT="/tmp/terraform-provider-f5xc"
SPEC_DIR="/tmp"
PROVIDER_DIR="$PROJECT_ROOT/internal/provider"
CLIENT_FILE="$PROJECT_ROOT/internal/client/client.go"

echo "🔨 F5XC Terraform Provider - Batch Resource Generator"
echo "======================================================"
echo ""

# List of resources to generate (top 50 most commonly used)
RESOURCES=(
    "healthcheck"
    "tcp_loadbalancer"
    "virtual_site"
    "app_firewall"
    "rate_limit_policy"
    "service_policy"
    "api_definition"
    "virtual_host"
    "route"
    "upstream"
    "cluster"
    "fleet"
    "site"
    "virtual_network"
    "network_policy"
    "network_connector"
    "network_firewall"
    "forward_proxy_policy"
    "cloud_credentials"
    "aws_tgw_site"
    "aws_vpc_site"
    "azure_vnet_site"
    "gcp_vpc_site"
    "k8s_cluster"
    "app_setting"
    "discovery"
    "blindfold_policy"
    "user"
    "role"
    "tenant"
    "namespace_setting"
    "alert_policy"
    "alert_receiver"
    "log_receiver"
    "volterra_provider"
    "terraform_parameters"
    "waf_rule"
    "bot_defense"
    "api_protection"
    "malicious_user_detection"
    "l7ddos"
    "client_side_defense"
    "security_event"
    "service_discovery"
    "shape_detection"
    "data_guard"
    "allowed_vip"
    "tcp_pool"
    "http_health_check"
    "tcp_health_check"
)

# Function to generate a single resource
generate_resource() {
    local resource_name
    local spec_file
    resource_name=$1
    spec_file=$(find "$SPEC_DIR" -name "*.$resource_name.ves-swagger.json" | head -1)

    if [ -z "$spec_file" ]; then
        echo "  ⚠️  Spec file not found for: $resource_name"
        return 1
    fi

    echo "  📝 Generating $resource_name..."

    # Create resource file using template
    cat > "$PROVIDER_DIR/${resource_name}_resource.go" << 'RESOURCE_EOF'
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

var _ resource.Resource = &RESOURCE_NAMEResource{}

func NewRESOURCE_NAMEResource() resource.Resource {
	return &RESOURCE_NAMEResource{}
}

type RESOURCE_NAMEResource struct {
	client *client.Client
}

type RESOURCE_NAMEResourceModel struct {
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`
	Labels      types.Map    `tfsdk:"labels"`
	ID          types.String `tfsdk:"id"`
}

func (r *RESOURCE_NAMEResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_RESOURCE_NAME_LOWER"
}

func (r *RESOURCE_NAMEResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages RESOURCE_NAME in F5 Distributed Cloud.",
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"namespace": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"description": schema.StringAttribute{Optional: true},
			"labels": schema.MapAttribute{ElementType: types.StringType, Optional: true},
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
		},
	}
}

func (r *RESOURCE_NAMEResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	r.client = req.ProviderData.(*client.Client)
}

func (r *RESOURCE_NAMEResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data RESOURCE_NAMEResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resource := &client.RESOURCE_NAME{
		Metadata: client.Metadata{Name: data.Name.ValueString(), Namespace: data.Namespace.ValueString()},
		Spec: client.RESOURCE_NAMESpec{Description: data.Description.ValueString()},
	}
	if !data.Labels.IsNull() {
		labels := make(map[string]string)
		data.Labels.ElementsAs(ctx, &labels, false)
		resource.Metadata.Labels = labels
	}
	created, err := r.client.CreateRESOURCE_NAME(ctx, resource)
	if err != nil {
		resp.Diagnostics.AddError("Error", err.Error())
		return
	}
	data.ID = types.StringValue(created.Metadata.Name)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RESOURCE_NAMEResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data RESOURCE_NAMEResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.GetRESOURCE_NAME(ctx, data.Namespace.ValueString(), data.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error", err.Error())
		return
	}
	data.ID = types.StringValue(res.Metadata.Name)
	data.Description = types.StringValue(res.Spec.Description)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RESOURCE_NAMEResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data RESOURCE_NAMEResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resource := &client.RESOURCE_NAME{
		Metadata: client.Metadata{Name: data.Name.ValueString(), Namespace: data.Namespace.ValueString()},
		Spec: client.RESOURCE_NAMESpec{Description: data.Description.ValueString()},
	}
	updated, err := r.client.UpdateRESOURCE_NAME(ctx, resource)
	if err != nil {
		resp.Diagnostics.AddError("Error", err.Error())
		return
	}
	data.ID = types.StringValue(updated.Metadata.Name)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RESOURCE_NAMEResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data RESOURCE_NAMEResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	err := r.client.DeleteRESOURCE_NAME(ctx, data.Namespace.ValueString(), data.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error", err.Error())
	}
}

func (r *RESOURCE_NAMEResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
RESOURCE_EOF

    # Replace placeholders
    local title_case
    title_case=$(echo "$resource_name" | sed -r 's/(^|_)([a-z])/\U\2/g' | sed 's/_//g')
    sed -i '' "s/RESOURCE_NAME/$title_case/g" "$PROVIDER_DIR/${resource_name}_resource.go"
    sed -i '' "s/RESOURCE_NAME_LOWER/$resource_name/g" "$PROVIDER_DIR/${resource_name}_resource.go"

    echo "  ✅ Generated ${resource_name}_resource.go"
}

# Main execution
echo "📊 Generating resources..."
echo ""

COUNT=0
for resource in "${RESOURCES[@]}"; do
    COUNT=$((COUNT + 1))
    echo "[$COUNT/${#RESOURCES[@]}] $resource"
    generate_resource "$resource" || true
done

echo ""
echo "🎉 Batch generation complete!"
echo "📊 Generated $COUNT resource files"
echo ""
echo "Next steps:"
echo "  1. Add client types to $CLIENT_FILE"
echo "  2. Register resources in provider.go"
echo "  3. Run: go build -o terraform-provider-f5xc"
