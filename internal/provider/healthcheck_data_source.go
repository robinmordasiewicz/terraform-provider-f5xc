// Copyright (c) F5XC Community
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/f5xc/terraform-provider-f5xc/internal/client"
)

var _ datasource.DataSource = &HealthcheckDataSource{}

func NewHealthcheckDataSource() datasource.DataSource {
	return &HealthcheckDataSource{}
}

type HealthcheckDataSource struct {
	client *client.Client
}

type HealthcheckDataSourceModel struct {
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`
	Labels      types.Map    `tfsdk:"labels"`
	ID          types.String `tfsdk:"id"`
}

func (d *HealthcheckDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_healthcheck"
}

func (d *HealthcheckDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Fetches information about an existing Healthcheck in F5 Distributed Cloud.",

		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				MarkdownDescription: "Name of the Healthcheck to look up.",
				Required:            true,
			},
			"namespace": schema.StringAttribute{
				MarkdownDescription: "Namespace where the Healthcheck exists.",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description of the Healthcheck.",
				Computed:            true,
			},
			"labels": schema.MapAttribute{
				MarkdownDescription: "Labels applied to the Healthcheck.",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Unique identifier for the Healthcheck.",
			},
		},
	}
}

func (d *HealthcheckDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

func (d *HealthcheckDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data HealthcheckDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resource, err := d.client.GetHealthcheck(ctx, data.Namespace.ValueString(), data.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Healthcheck, got error: %s", err))
		return
	}

	data.ID = types.StringValue(resource.Metadata.Name)
	data.Name = types.StringValue(resource.Metadata.Name)
	data.Namespace = types.StringValue(resource.Metadata.Namespace)
	data.Description = types.StringValue(resource.Spec.Description)

	if len(resource.Metadata.Labels) > 0 {
		labels, diags := types.MapValueFrom(ctx, types.StringType, resource.Metadata.Labels)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
		data.Labels = labels
	} else {
		data.Labels = types.MapNull(types.StringType)
	}

	tflog.Trace(ctx, "read Healthcheck data source")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
