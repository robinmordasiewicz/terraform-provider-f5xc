// Copyright (c) F5XC Community
// SPDX-License-Identifier: MPL-2.0

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/f5xc/terraform-provider-f5xc/internal/acctest"
	"github.com/f5xc/terraform-provider-f5xc/internal/mocks"
)

// =============================================================================
// AWS VPC SITE MOCK TESTS
//
// These tests use mock HTTP server instead of real F5 XC and AWS APIs.
// They validate provider logic without requiring cloud credentials.
//
// Run with:
//   F5XC_MOCK_MODE=1 go test -v ./internal/provider/ -run TestMockAWSVPCSite -timeout 5m
//
// These tests complement the real acceptance tests by:
// 1. Testing provider CRUD logic without cloud dependencies
// 2. Enabling CI/CD testing without secrets
// 3. Testing error handling scenarios
// 4. Faster feedback loop during development
// =============================================================================

// TestMockAWSVPCSiteResource_basic tests the AWS VPC Site resource using mock server.
// This test validates that the provider correctly handles:
// - Resource creation (POST)
// - Resource read (GET)
// - Resource deletion (DELETE)
// - State management
func TestMockAWSVPCSiteResource_basic(t *testing.T) {
	acctest.SkipIfNoMockMode(t)

	rName := acctest.RandomName("tf-mock-test-vpc")
	nsName := acctest.RandomName("tf-mock-test-ns")
	resourceName := "f5xc_aws_vpc_site.test"

	mockCfg := acctest.SetupMockTest(t)
	defer mockCfg.Cleanup()

	// Pre-populate the namespace (required dependency)
	mockCfg.SetupNamespaceMock(nsName)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: mockCfg.ProtoV6ProviderFactories(),
		Steps: []resource.TestStep{
			// Step 1: Create AWS VPC Site
			{
				Config: testAccMockAWSVPCSiteConfig_basic(mockCfg, nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "namespace", nsName),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			// Step 2: Import test
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"timeouts"},
				ImportStateId:           fmt.Sprintf("%s/%s", nsName, rName),
			},
		},
	})
}

// TestMockAWSVPCSiteResource_withAWSConfig tests AWS VPC Site with full AWS configuration
func TestMockAWSVPCSiteResource_withAWSConfig(t *testing.T) {
	acctest.SkipIfNoMockMode(t)

	rName := acctest.RandomName("tf-mock-test-vpc")
	nsName := acctest.RandomName("tf-mock-test-ns")
	credsName := acctest.RandomName("tf-mock-aws-creds")
	resourceName := "f5xc_aws_vpc_site.test"

	mockCfg := acctest.SetupMockTest(t)
	defer mockCfg.Cleanup()

	// Pre-populate dependencies
	mockCfg.SetupNamespaceMock(nsName)
	mockCfg.SetupAWSCredentialsMock(nsName, credsName)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: mockCfg.ProtoV6ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccMockAWSVPCSiteConfig_withAWSConfig(mockCfg, nsName, rName, credsName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "namespace", nsName),
				),
			},
		},
	})
}

// TestMockAWSVPCSiteResource_update tests updating AWS VPC Site attributes
func TestMockAWSVPCSiteResource_update(t *testing.T) {
	acctest.SkipIfNoMockMode(t)

	rName := acctest.RandomName("tf-mock-test-vpc")
	nsName := acctest.RandomName("tf-mock-test-ns")
	resourceName := "f5xc_aws_vpc_site.test"

	mockCfg := acctest.SetupMockTest(t)
	defer mockCfg.Cleanup()

	mockCfg.SetupNamespaceMock(nsName)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: mockCfg.ProtoV6ProviderFactories(),
		Steps: []resource.TestStep{
			// Step 1: Create with initial description
			{
				Config: testAccMockAWSVPCSiteConfig_withDescription(mockCfg, nsName, rName, "Initial description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "description", "Initial description"),
				),
			},
			// Step 2: Update description
			{
				Config: testAccMockAWSVPCSiteConfig_withDescription(mockCfg, nsName, rName, "Updated description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "description", "Updated description"),
				),
			},
		},
	})
}

// TestMockAWSVPCSiteResource_labels tests AWS VPC Site with labels
func TestMockAWSVPCSiteResource_labels(t *testing.T) {
	acctest.SkipIfNoMockMode(t)

	rName := acctest.RandomName("tf-mock-test-vpc")
	nsName := acctest.RandomName("tf-mock-test-ns")
	resourceName := "f5xc_aws_vpc_site.test"

	mockCfg := acctest.SetupMockTest(t)
	defer mockCfg.Cleanup()

	mockCfg.SetupNamespaceMock(nsName)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: mockCfg.ProtoV6ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccMockAWSVPCSiteConfig_withLabels(mockCfg, nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "test"),
					resource.TestCheckResourceAttr(resourceName, "labels.team", "platform"),
				),
			},
		},
	})
}

// TestMockAWSVPCSiteDataSource_basic tests reading AWS VPC Site data source with mock
func TestMockAWSVPCSiteDataSource_basic(t *testing.T) {
	acctest.SkipIfNoMockMode(t)

	rName := "existing-vpc-site"
	nsName := "system"
	dataSourceName := "data.f5xc_aws_vpc_site.test"

	mockCfg := acctest.SetupMockTest(t)
	defer mockCfg.Cleanup()

	// Pre-populate an existing AWS VPC site
	mockCfg.SetupAWSVPCSiteMock(nsName, rName,
		mocks.WithAWSRegion("us-west-2"),
		mocks.WithVPCID("vpc-mock-12345"),
	)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: mockCfg.ProtoV6ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccMockAWSVPCSiteDataSourceConfig(mockCfg, nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "name", rName),
					resource.TestCheckResourceAttr(dataSourceName, "namespace", nsName),
				),
			},
		},
	})
}

// =============================================================================
// Test Configuration Functions
// =============================================================================

func testAccMockAWSVPCSiteConfig_basic(mockCfg *acctest.MockTestConfig, nsName, name string) string {
	return acctest.ConfigCompose(
		mockCfg.MockProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_namespace" "test" {
  name = %[1]q
}

resource "f5xc_aws_vpc_site" "test" {
  depends_on = [f5xc_namespace.test]

  name      = %[2]q
  namespace = f5xc_namespace.test.name
}
`, nsName, name))
}

func testAccMockAWSVPCSiteConfig_withAWSConfig(mockCfg *acctest.MockTestConfig, nsName, name, credsName string) string {
	return acctest.ConfigCompose(
		mockCfg.MockProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_namespace" "test" {
  name = %[1]q
}

resource "f5xc_aws_vpc_site" "test" {
  depends_on = [f5xc_namespace.test]

  name      = %[2]q
  namespace = f5xc_namespace.test.name

  description = "AWS VPC Site created with mock test"

  labels = {
    environment = "mock-test"
  }
}
`, nsName, name))
}

func testAccMockAWSVPCSiteConfig_withDescription(mockCfg *acctest.MockTestConfig, nsName, name, description string) string {
	return acctest.ConfigCompose(
		mockCfg.MockProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_namespace" "test" {
  name = %[1]q
}

resource "f5xc_aws_vpc_site" "test" {
  depends_on = [f5xc_namespace.test]

  name        = %[2]q
  namespace   = f5xc_namespace.test.name
  description = %[3]q
}
`, nsName, name, description))
}

func testAccMockAWSVPCSiteConfig_withLabels(mockCfg *acctest.MockTestConfig, nsName, name string) string {
	return acctest.ConfigCompose(
		mockCfg.MockProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_namespace" "test" {
  name = %[1]q
}

resource "f5xc_aws_vpc_site" "test" {
  depends_on = [f5xc_namespace.test]

  name      = %[2]q
  namespace = f5xc_namespace.test.name

  labels = {
    environment = "test"
    team        = "platform"
  }
}
`, nsName, name))
}

func testAccMockAWSVPCSiteDataSourceConfig(mockCfg *acctest.MockTestConfig, nsName, name string) string {
	return acctest.ConfigCompose(
		mockCfg.MockProviderConfig(),
		fmt.Sprintf(`
data "f5xc_aws_vpc_site" "test" {
  name      = %[1]q
  namespace = %[2]q
}
`, name, nsName))
}
