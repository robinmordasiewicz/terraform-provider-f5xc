// Copyright (c) F5XC Community
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/f5xc/terraform-provider-f5xc/internal/client"
)

// Ensure F5XCProvider satisfies various provider interfaces.
var _ provider.Provider = &F5XCProvider{}

// F5XCProvider defines the provider implementation.
type F5XCProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// F5XCProviderModel describes the provider data model.
type F5XCProviderModel struct {
	APIToken types.String `tfsdk:"api_token"`
	APIURL   types.String `tfsdk:"api_url"`
}

func (p *F5XCProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "f5xc"
	resp.Version = p.version
}

func (p *F5XCProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Terraform provider for F5 Distributed Cloud (F5XC). " +
			"This is an open source community provider built from public F5 API documentation.",
		Attributes: map[string]schema.Attribute{
			"api_token": schema.StringAttribute{
				MarkdownDescription: "F5 Distributed Cloud API Token. Can also be set via F5XC_API_TOKEN environment variable.",
				Optional:            true,
				Sensitive:           true,
			},
			"api_url": schema.StringAttribute{
				MarkdownDescription: "F5 Distributed Cloud API URL. Defaults to https://console.ves.volterra.io/api. " +
					"Can also be set via F5XC_API_URL environment variable.",
				Optional: true,
			},
		},
	}
}

func (p *F5XCProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Info(ctx, "Configuring F5XC client")

	var config F5XCProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Check for environment variables if not set in configuration
	apiToken := os.Getenv("F5XC_API_TOKEN")
	apiURL := os.Getenv("F5XC_API_URL")

	// Configuration values override environment variables
	if !config.APIToken.IsNull() {
		apiToken = config.APIToken.ValueString()
	}

	if !config.APIURL.IsNull() {
		apiURL = config.APIURL.ValueString()
	}

	// Set default API URL if not provided
	if apiURL == "" {
		apiURL = "https://console.ves.volterra.io/api"
	}

	// Validate that API token is provided
	if apiToken == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("api_token"),
			"Missing F5XC API Token",
			"The provider cannot create the F5XC API client as there is a missing or empty value for the F5XC API token. "+
				"Set the api_token value in the configuration or use the F5XC_API_TOKEN environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
		return
	}

	// Create the F5XC client
	c := client.NewClient(apiURL, apiToken)

	// Make the client available during DataSource and Resource type Configure methods
	resp.DataSourceData = c
	resp.ResourceData = c

	tflog.Info(ctx, "Configured F5XC client", map[string]any{"success": true, "api_url": apiURL})
}

func (p *F5XCProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewAddressAllocatorResource,
		NewAdvertisePolicyResource,
		NewAlertPolicyResource,
		NewAlertReceiverResource,
		NewAPICrawlerResource,
		NewAPIDefinitionResource,
		NewAPIDiscoveryResource,
		NewAPITestingResource,
		NewApmResource,
		NewAppAPIGroupResource,
		NewAppFirewallResource,
		NewAppSettingResource,
		NewAppTypeResource,
		NewAuthenticationResource,
		NewAWSTGWSiteResource,
		NewAWSVPCSiteResource,
		NewAzureVNETSiteResource,
		NewBgpAsnSetResource,
		NewBgpResource,
		NewBgpRoutingPolicyResource,
		NewBigIPVirtualServerResource,
		NewBotDefenseAppInfrastructureResource,
		NewCdnCacheRuleResource,
		NewCdnLoadBalancerResource,
		NewCertificateChainResource,
		NewCertificateResource,
		NewCertifiedHardwareResource,
		NewCloudConnectResource,
		NewCloudCredentialsResource,
		NewCloudElasticIPResource,
		NewCloudLinkResource,
		NewCloudRegionResource,
		NewClusterResource,
		NewCminstanceResource,
		NewCodeBaseIntegrationResource,
		NewContainerRegistryResource,
		NewCrlResource,
		NewDataGroupResource,
		NewDataTypeResource,
		NewDcClusterGroupResource,
		NewDiscoveryResource,
		NewDNSComplianceChecksResource,
		NewDNSDomainResource,
		NewEndpointResource,
		NewEnhancedFirewallPolicyResource,
		NewExternalConnectorResource,
		NewFastAclResource,
		NewFastAclRuleResource,
		NewFilterSetResource,
		NewFleetResource,
		NewFlowAnomalyResource,
		NewFlowResource,
		NewForwardingClassResource,
		NewForwardProxyPolicyResource,
		NewGCPVPCSiteResource,
		NewGlobalLogReceiverResource,
		NewHealthcheckResource,
		NewHTTPLoadBalancerResource,
		NewIke1Resource,
		NewIke2Resource,
		NewIKEPhase1ProfileResource,
		NewIKEPhase2ProfileResource,
		NewImplicitLabelResource,
		NewIPPrefixSetResource,
		NewIruleResource,
		NewK8sClusterResource,
		NewK8sClusterRoleBindingResource,
		NewK8sClusterRoleResource,
		NewK8sPodSecurityAdmissionResource,
		NewK8sPodSecurityPolicyResource,
		NewKnownLabelKeyResource,
		NewKnownLabelResource,
		NewLmaRegionResource,
		NewLogReceiverResource,
		NewMaliciousUserMitigationResource,
		NewModuleManagementResource,
		NewNamespaceResource,
		NewNatPolicyResource,
		NewNetworkConnectorResource,
		NewNetworkFirewallResource,
		NewNetworkInterfaceResource,
		NewNetworkPolicyResource,
		NewNetworkPolicyRuleResource,
		NewNetworkPolicySetResource,
		NewNetworkPolicyViewResource,
		NewNfvServiceResource,
		NewOriginPoolResource,
		NewPolicerResource,
		NewPolicyBasedRoutingResource,
		NewProtocolInspectionResource,
		NewProtocolPolicerResource,
		NewProxyResource,
		NewPublicIPResource,
		NewRateLimiterPolicyResource,
		NewRateLimiterResource,
		NewRouteResource,
		NewRuleSuggestionResource,
		NewSecretManagementAccessResource,
		NewSecuremeshSiteResource,
		NewSecuremeshSiteV2Resource,
		NewSegmentConnectionResource,
		NewSegmentResource,
		NewSensitiveDataPolicyResource,
		NewServicePolicyResource,
		NewServicePolicyRuleResource,
		NewServicePolicySetResource,
		NewShapeBotDefenseInstanceResource,
		NewSiteMeshGroupResource,
		NewSiteResource,
		NewSrv6NetworkSliceResource,
		NewSubnetResource,
		NewTCPLoadBalancerResource,
		NewTenantConfigurationResource,
		NewTerraformParametersResource,
		NewThirdPartyApplicationResource,
		NewTrustedCaListResource,
		NewTunnelResource,
		NewUDPLoadBalancerResource,
		NewUsbPolicyResource,
		NewUserIdentificationResource,
		NewUserTokenResource,
		NewViewInternalResource,
		NewVirtualHostResource,
		NewVirtualK8sResource,
		NewVirtualNetworkResource,
		NewVirtualSiteResource,
		NewVoltstackSiteResource,
		NewWAFExclusionPolicyResource,
		NewWorkloadFlavorResource,
		NewWorkloadResource,
	}
}

func (p *F5XCProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewNamespaceDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &F5XCProvider{
			version: version,
		}
	}
}
