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

var _ datasource.DataSource = &OriginPoolDataSource{}

func NewOriginPoolDataSource() datasource.DataSource {
	return &OriginPoolDataSource{}
}

type OriginPoolDataSource struct {
	client *client.Client
}

type OriginPoolDataSourceModel struct {
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`
	Labels      types.Map    `tfsdk:"labels"`
	ID          types.String `tfsdk:"id"`
}

func (d *OriginPoolDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_origin_pool"
}

func (d *OriginPoolDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Fetches information about an existing Origin Pool in F5 Distributed Cloud.",

		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				MarkdownDescription: "Name of the Origin Pool to look up.",
				Required:            true,
			},
			"namespace": schema.StringAttribute{
				MarkdownDescription: "Namespace where the Origin Pool exists.",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description of the Origin Pool.",
				Computed:            true,
			},
			"labels": schema.MapAttribute{
				MarkdownDescription: "Labels applied to the Origin Pool.",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Unique identifier for the Origin Pool.",
			},
		},
	}
}

func (d *OriginPoolDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *OriginPoolDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data OriginPoolDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resource, err := d.client.GetOriginPool(ctx, data.Namespace.ValueString(), data.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Origin Pool, got error: %s", err))
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

	tflog.Trace(ctx, "read Origin Pool data source")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
