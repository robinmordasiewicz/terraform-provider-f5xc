// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/f5xc/terraform-provider-f5xc/internal/acctest"
	"github.com/f5xc/terraform-provider-f5xc/internal/mocks"
)

// =============================================================================
// TENANT CONFIGURATION MOCK TESTS
//
// The real tenant configuration API returns 500/501 errors in some environments
// and requires tenant admin permissions. These mock tests allow us to:
// 1. Test the provider's resource implementation logic
// 2. Verify schema correctness
// 3. Test state management
// 4. Run tests in CI/CD without special permissions
//
// Run with:
//   F5XC_MOCK_MODE=1 go test -v ./internal/provider/ -run TestMockTenantConfiguration -timeout 5m
// =============================================================================

// TestMockTenantConfigurationResource_basic tests basic tenant configuration operations
func TestMockTenantConfigurationResource_basic(t *testing.T) {
	acctest.SkipIfNoMockMode(t)

	resourceName := "f5xc_tenant_configuration.test"

	mockCfg := acctest.SetupMockTest(t)
	defer mockCfg.Cleanup()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: mockCfg.ProtoV6ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccMockTenantConfigurationConfig_basic(mockCfg),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
		},
	})
}

// TestMockTenantConfigurationDataSource_basic tests reading tenant configuration data source
func TestMockTenantConfigurationDataSource_basic(t *testing.T) {
	acctest.SkipIfNoMockMode(t)

	dataSourceName := "data.f5xc_tenant_configuration.test"

	mockCfg := acctest.SetupMockTest(t)
	defer mockCfg.Cleanup()

	// Pre-populate tenant configuration
	path := mocks.ResourcePath("system", "tenant_configurations", "tenant-config")
	mockCfg.PrePopulateResource(path, mocks.TenantConfigurationResponse("tenant-config"))

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: mockCfg.ProtoV6ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccMockTenantConfigurationDataSourceConfig(mockCfg),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(dataSourceName, "id"),
				),
			},
		},
	})
}

// TestMockTenantConfigurationResource_errorHandling tests error scenarios
func TestMockTenantConfigurationResource_errorHandling(t *testing.T) {
	acctest.SkipIfNoMockMode(t)

	mockCfg := acctest.SetupMockTest(t)
	defer mockCfg.Cleanup()

	// Simulate 501 NOT_IMPLEMENTED error (common in staging environments)
	mockCfg.Simulate501NotImplemented("system", "tenant_configurations")

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: mockCfg.ProtoV6ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccMockTenantConfigurationConfig_basic(mockCfg),
				// Verify that the provider properly reports the 501 error from the API
				ExpectError: acctest.MustCompileRegexp(`(?i)(501|not implemented|SERVER_ERROR)`),
			},
		},
	})
}

// =============================================================================
// Test Configuration Functions
// =============================================================================

func testAccMockTenantConfigurationConfig_basic(mockCfg *acctest.MockTestConfig) string {
	return acctest.ConfigCompose(
		mockCfg.MockProviderConfig(),
		`
resource "f5xc_tenant_configuration" "test" {
  name      = "tenant-config"
  namespace = "system"
}
`)
}

func testAccMockTenantConfigurationDataSourceConfig(mockCfg *acctest.MockTestConfig) string {
	return acctest.ConfigCompose(
		mockCfg.MockProviderConfig(),
		`
data "f5xc_tenant_configuration" "test" {
  name      = "tenant-config"
  namespace = "system"
}
`)
}
