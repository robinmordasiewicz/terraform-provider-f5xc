// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

// Addon Service Data Source for F5 XC
// Provides read-only access to addon service information

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
	_ datasource.DataSource              = &AddonServiceDataSource{}
	_ datasource.DataSourceWithConfigure = &AddonServiceDataSource{}
)

func NewAddonServiceDataSource() datasource.DataSource {
	return &AddonServiceDataSource{}
}

type AddonServiceDataSource struct {
	client *client.Client
}

type AddonServiceDataSourceModel struct {
	ID                           types.String `tfsdk:"id"`
	Name                         types.String `tfsdk:"name"`
	Namespace                    types.String `tfsdk:"namespace"`
	DisplayName                  types.String `tfsdk:"display_name"`
	Tier                         types.String `tfsdk:"tier"`
	ActivationType               types.String `tfsdk:"activation_type"`
	AddonServiceGroupName        types.String `tfsdk:"addon_service_group_name"`
	AddonServiceGroupDisplayName types.String `tfsdk:"addon_service_group_display_name"`
}

func (d *AddonServiceDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_addon_service"
}

func (d *AddonServiceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: `Retrieves information about an F5 Distributed Cloud Addon Service.

Addon services are system-managed resources that provide additional functionality
such as Bot Defense, Client Side Defense, and other security features. This data
source allows you to query addon service details including tier requirements and
activation type.

~> **Note:** Addon services cannot be created or modified via Terraform. Use the
` + "`f5xc_addon_subscription`" + ` resource to subscribe to an addon service.`,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Unique identifier for the data source.",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Name of the addon service (e.g., `bot_defense`, `client_side_defense`).",
				Required:            true,
			},
			"namespace": schema.StringAttribute{
				MarkdownDescription: "Namespace where the addon service is defined. Usually `shared`.",
				Optional:            true,
				Computed:            true,
			},
			"display_name": schema.StringAttribute{
				MarkdownDescription: "Human-readable display name of the addon service.",
				Computed:            true,
			},
			"tier": schema.StringAttribute{
				MarkdownDescription: "Subscription tier required for this addon service. Possible values: `NO_TIER`, `BASIC`, `STANDARD`, `ADVANCED`, `PREMIUM`.",
				Computed:            true,
			},
			"activation_type": schema.StringAttribute{
				MarkdownDescription: "How the addon service is activated. Possible values: `self` (user can activate directly), `partial` (requires partial SRE management), `managed` (requires full manual intervention).",
				Computed:            true,
			},
			"addon_service_group_name": schema.StringAttribute{
				MarkdownDescription: "Name of the addon service group this service belongs to.",
				Computed:            true,
			},
			"addon_service_group_display_name": schema.StringAttribute{
				MarkdownDescription: "Display name of the addon service group.",
				Computed:            true,
			},
		},
	}
}

func (d *AddonServiceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *AddonServiceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data AddonServiceDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Default namespace to "shared" if not specified
	namespace := data.Namespace.ValueString()
	if namespace == "" {
		namespace = "shared"
	}

	// Get detailed addon service information using the custom API
	details, err := d.client.GetAddonServiceDetails(ctx, data.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read addon service details: %s", err))
		return
	}

	// Set basic fields
	data.ID = types.StringValue(fmt.Sprintf("%s/%s", namespace, data.Name.ValueString()))
	data.Namespace = types.StringValue(namespace)

	// Set display name
	if details.DisplayName != "" {
		data.DisplayName = types.StringValue(details.DisplayName)
	} else {
		data.DisplayName = types.StringNull()
	}

	// Set tier
	if details.Tier != "" {
		data.Tier = types.StringValue(details.Tier)
	} else {
		data.Tier = types.StringValue("NO_TIER")
	}

	// Determine activation type based on which field is populated
	switch {
	case details.SelfActivation != nil:
		data.ActivationType = types.StringValue("self")
	case details.PartiallyManagedActivation != nil:
		data.ActivationType = types.StringValue("partial")
	case details.ManagedActivation != nil:
		data.ActivationType = types.StringValue("managed")
	default:
		data.ActivationType = types.StringNull()
	}

	// Set group information
	if details.AddonServiceGroupName != "" {
		data.AddonServiceGroupName = types.StringValue(details.AddonServiceGroupName)
	} else {
		data.AddonServiceGroupName = types.StringNull()
	}

	if details.AddonServiceGroupDisplayName != "" {
		data.AddonServiceGroupDisplayName = types.StringValue(details.AddonServiceGroupDisplayName)
	} else {
		data.AddonServiceGroupDisplayName = types.StringNull()
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
