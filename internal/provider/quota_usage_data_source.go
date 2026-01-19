// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

// Quota Usage Data Source for F5 XC
// Provides read-only access to namespace quota information for resource planning
// and lifecycle preconditions.
//
// This is a manually maintained data source - not auto-generated from OpenAPI specs.

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
	_ datasource.DataSource              = &QuotaUsageDataSource{}
	_ datasource.DataSourceWithConfigure = &QuotaUsageDataSource{}
)

// NewQuotaUsageDataSource creates a new quota usage data source
func NewQuotaUsageDataSource() datasource.DataSource {
	return &QuotaUsageDataSource{}
}

// QuotaUsageDataSource provides quota information for a namespace
type QuotaUsageDataSource struct {
	client *client.Client
}

// QuotaUsageDataSourceModel describes the data source model
type QuotaUsageDataSourceModel struct {
	ID        types.String                   `tfsdk:"id"`
	Namespace types.String                   `tfsdk:"namespace"`
	Objects   map[string]ObjectQuotaModel    `tfsdk:"objects"`
}

// ObjectQuotaModel represents quota information for a single object type
type ObjectQuotaModel struct {
	Limit     types.Int64 `tfsdk:"limit"`
	Usage     types.Int64 `tfsdk:"usage"`
	Available types.Int64 `tfsdk:"available"`
}

func (d *QuotaUsageDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_quota_usage"
}

func (d *QuotaUsageDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: `Retrieves quota usage information for an F5 Distributed Cloud namespace.

This data source allows you to query current quota usage for all resource types
in a namespace. Use it with Terraform lifecycle preconditions to prevent resource
creation failures due to quota exhaustion.

## Example Usage

` + "```hcl" + `
# Query quota usage for the system namespace
data "f5xc_quota_usage" "system" {
  namespace = "system"
}

# Output healthcheck quota information
output "healthcheck_quota" {
  value = {
    used      = data.f5xc_quota_usage.system.objects["healthcheck"].usage
    limit     = data.f5xc_quota_usage.system.objects["healthcheck"].limit
    available = data.f5xc_quota_usage.system.objects["healthcheck"].available
  }
}

# Use with lifecycle preconditions
resource "f5xc_healthcheck" "example" {
  name      = "my-healthcheck"
  namespace = "system"
  # ... other configuration

  lifecycle {
    precondition {
      condition     = data.f5xc_quota_usage.system.objects["healthcheck"].available > 0
      error_message = "No healthcheck quota available in system namespace"
    }
  }
}
` + "```" + `

~> **Note:** The quota API returns information for all resource types in the namespace.
Not all resource types may have quota limits enforced.`,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Unique identifier for the data source (namespace name).",
				Computed:            true,
			},
			"namespace": schema.StringAttribute{
				MarkdownDescription: "Name of the namespace to query quota usage for. Common values include `system`, `default`, or custom namespace names.",
				Required:            true,
			},
			"objects": schema.MapNestedAttribute{
				MarkdownDescription: "Map of resource type names to their quota information. Keys are resource type names (e.g., `healthcheck`, `http_loadbalancer`, `origin_pool`).",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"limit": schema.Int64Attribute{
							MarkdownDescription: "Maximum number of objects allowed for this resource type.",
							Computed:            true,
						},
						"usage": schema.Int64Attribute{
							MarkdownDescription: "Current number of objects of this resource type.",
							Computed:            true,
						},
						"available": schema.Int64Attribute{
							MarkdownDescription: "Number of objects that can still be created (limit - usage).",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *QuotaUsageDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *QuotaUsageDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data QuotaUsageDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	namespace := data.Namespace.ValueString()

	// Get quota usage from the API
	usage, err := d.client.GetQuotaUsage(ctx, namespace)
	if err != nil {
		resp.Diagnostics.AddError(
			"Client Error",
			fmt.Sprintf("Unable to read quota usage for namespace %s: %s", namespace, err),
		)
		return
	}

	// Set the ID to the namespace
	data.ID = types.StringValue(namespace)

	// Convert API response to model
	data.Objects = make(map[string]ObjectQuotaModel)
	for resourceType, quota := range usage.Objects {
		available := quota.Limit.Maximum - quota.Usage.Current
		data.Objects[resourceType] = ObjectQuotaModel{
			Limit:     types.Int64Value(int64(quota.Limit.Maximum)),
			Usage:     types.Int64Value(int64(quota.Usage.Current)),
			Available: types.Int64Value(int64(available)),
		}
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
