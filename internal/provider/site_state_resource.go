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
	_ resource.Resource                = &SiteStateResource{}
	_ resource.ResourceWithConfigure   = &SiteStateResource{}
	_ resource.ResourceWithImportState = &SiteStateResource{}
)

func NewSiteStateResource() resource.Resource {
	return &SiteStateResource{}
}

type SiteStateResource struct {
	client *client.Client
}

type SiteStateResourceModel struct {
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`
	State       types.String `tfsdk:"state"`
	Labels      types.Map    `tfsdk:"labels"`
	ID          types.String `tfsdk:"id"`
}

func (r *SiteStateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_site_state"
}

func (r *SiteStateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages a Site State in F5 Distributed Cloud.",
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				MarkdownDescription: "Name of the Site State.",
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"namespace": schema.StringAttribute{
				MarkdownDescription: "Namespace where the Site State will be created.",
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description of the Site State.",
				Optional:            true,
			},
			"state": schema.StringAttribute{
				MarkdownDescription: "State of the site (e.g., ONLINE, OFFLINE).",
				Optional:            true,
			},
			"labels": schema.MapAttribute{
				MarkdownDescription: "Labels to apply to the Site State.",
				ElementType:         types.StringType,
				Optional:            true,
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Unique identifier for the Site State.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *SiteStateResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *SiteStateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SiteStateResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resource := &client.SiteState{
		Metadata: client.Metadata{
			Name:      data.Name.ValueString(),
			Namespace: data.Namespace.ValueString(),
		},
		Spec: client.SiteStateSpec{
			Description: data.Description.ValueString(),
			State:       data.State.ValueString(),
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

	created, err := r.client.CreateSiteState(ctx, resource)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create Site State: %s", err))
		return
	}

	data.ID = types.StringValue(created.Metadata.Name)
	tflog.Trace(ctx, "created Site State resource")
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SiteStateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SiteStateResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resource, err := r.client.GetSiteState(ctx, data.Namespace.ValueString(), data.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Site State: %s", err))
		return
	}

	data.ID = types.StringValue(resource.Metadata.Name)
	data.Name = types.StringValue(resource.Metadata.Name)
	data.Namespace = types.StringValue(resource.Metadata.Namespace)
	data.Description = types.StringValue(resource.Spec.Description)
	data.State = types.StringValue(resource.Spec.State)

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

	tflog.Trace(ctx, "read Site State resource")
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SiteStateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data SiteStateResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resource := &client.SiteState{
		Metadata: client.Metadata{
			Name:      data.Name.ValueString(),
			Namespace: data.Namespace.ValueString(),
		},
		Spec: client.SiteStateSpec{
			Description: data.Description.ValueString(),
			State:       data.State.ValueString(),
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

	updated, err := r.client.UpdateSiteState(ctx, resource)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update Site State: %s", err))
		return
	}

	data.ID = types.StringValue(updated.Metadata.Name)
	tflog.Trace(ctx, "updated Site State resource")
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SiteStateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SiteStateResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteSiteState(ctx, data.Namespace.ValueString(), data.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete Site State: %s", err))
		return
	}

	tflog.Trace(ctx, "deleted Site State resource")
}

func (r *SiteStateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
