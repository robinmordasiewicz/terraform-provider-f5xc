// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

// Addon Service Activation Status Data Source for F5 XC
// Provides information about whether an addon service can be activated

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
	_ datasource.DataSource              = &AddonServiceActivationStatusDataSource{}
	_ datasource.DataSourceWithConfigure = &AddonServiceActivationStatusDataSource{}
)

func NewAddonServiceActivationStatusDataSource() datasource.DataSource {
	return &AddonServiceActivationStatusDataSource{}
}

type AddonServiceActivationStatusDataSource struct {
	client *client.Client
}

type AddonServiceActivationStatusDataSourceModel struct {
	ID           types.String `tfsdk:"id"`
	AddonService types.String `tfsdk:"addon_service"`
	State        types.String `tfsdk:"state"`
	CanActivate  types.Bool   `tfsdk:"can_activate"`
	Message      types.String `tfsdk:"message"`
}

func (d *AddonServiceActivationStatusDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_addon_service_activation_status"
}

func (d *AddonServiceActivationStatusDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: `Checks the activation status of an F5 Distributed Cloud Addon Service.

Use this data source to determine if an addon service can be activated for your tenant
and what the current subscription state is.

**Possible state values:**

| State           | Description                              |
| --------------- | ---------------------------------------- |
| ` + "`AS_NONE`" + `       | Default state, service not subscribed    |
| ` + "`AS_PENDING`" + `    | Subscription request pending activation  |
| ` + "`AS_SUBSCRIBED`" + ` | Service is active and subscribed         |
| ` + "`AS_ERROR`" + `      | Subscription in error state              |`,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Unique identifier for the data source.",
				Computed:            true,
			},
			"addon_service": schema.StringAttribute{
				MarkdownDescription: "Name of the addon service to check (e.g., `bot_defense`, `client_side_defense`).",
				Required:            true,
			},
			"state": schema.StringAttribute{
				MarkdownDescription: "Current state of the addon service subscription. Possible values: `AS_NONE`, `AS_PENDING`, `AS_SUBSCRIBED`, `AS_ERROR`.",
				Computed:            true,
			},
			"can_activate": schema.BoolAttribute{
				MarkdownDescription: "Whether the addon service can be activated. True if state is `AS_NONE` (not yet subscribed) or `AS_SUBSCRIBED` (already active).",
				Computed:            true,
			},
			"message": schema.StringAttribute{
				MarkdownDescription: "Human-readable message describing the current activation status.",
				Computed:            true,
			},
		},
	}
}

func (d *AddonServiceActivationStatusDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *AddonServiceActivationStatusDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data AddonServiceActivationStatusDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get activation status
	status, err := d.client.GetAddonServiceActivationStatus(ctx, data.AddonService.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read addon service activation status: %s", err))
		return
	}

	// Set basic fields
	data.ID = types.StringValue(data.AddonService.ValueString())

	// Set state
	state := status.State
	if state == "" {
		state = "AS_NONE"
	}
	data.State = types.StringValue(state)

	// Determine if activation is possible and set message
	switch state {
	case "AS_NONE":
		data.CanActivate = types.BoolValue(true)
		data.Message = types.StringValue("Addon service is available for activation. Create an addon_subscription to activate.")
	case "AS_PENDING":
		data.CanActivate = types.BoolValue(false)
		data.Message = types.StringValue("Addon service activation is pending. Wait for the subscription to be enabled.")
	case "AS_SUBSCRIBED":
		data.CanActivate = types.BoolValue(true)
		data.Message = types.StringValue("Addon service is already subscribed and active.")
	case "AS_ERROR":
		data.CanActivate = types.BoolValue(false)
		data.Message = types.StringValue("Addon service subscription is in error state. Contact support for assistance.")
	default:
		data.CanActivate = types.BoolValue(false)
		data.Message = types.StringValue(fmt.Sprintf("Unknown activation state: %s", state))
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
