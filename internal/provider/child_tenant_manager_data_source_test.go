// Copyright (c) F5XC Community
// SPDX-License-Identifier: MPL-2.0

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/f5xc/terraform-provider-f5xc/internal/acctest"
)

// =============================================================================
// CHILD TENANT MANAGER DATA SOURCE ACCEPTANCE TESTS
//
// These tests verify the f5xc_child_tenant_manager data source implementation.
// Child Tenant Manager is a partner-only feature that requires a partner account
// with child tenant management capabilities.
//
// Run with:
//   TF_ACC=1 F5XC_API_URL="..." F5XC_P12_FILE="..." F5XC_P12_PASSWORD="..." \
//   go test -v ./internal/provider/ -run TestAccChildTenantManagerDataSource -timeout 30m
// =============================================================================

func TestAccChildTenantManagerDataSource_basic(t *testing.T) {
	// Skip: Child Tenant Manager requires an F5 XC partner account with child tenant
	// management capabilities, which is not available in standard tenant accounts
	t.Skip("Skipping: Child Tenant Manager requires F5 XC partner account with child tenant management capabilities")

	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test")
	nsName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_child_tenant_manager.test"
	dataSourceName := "data.f5xc_child_tenant_manager.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccChildTenantManagerDataSourceConfig_basic(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "namespace", resourceName, "namespace"),
					resource.TestCheckResourceAttrPair(dataSourceName, "id", resourceName, "id"),
				),
			},
		},
	})
}

func testAccChildTenantManagerDataSourceConfig_basic(nsName, name string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_namespace" "test" {
  name = %[1]q
}

resource "time_sleep" "wait_for_namespace" {
  depends_on      = [f5xc_namespace.test]
  create_duration = "5s"
}

resource "f5xc_child_tenant_manager" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[2]q
  namespace  = f5xc_namespace.test.name
}

data "f5xc_child_tenant_manager" "test" {
  depends_on = [f5xc_child_tenant_manager.test]
  name       = f5xc_child_tenant_manager.test.name
  namespace  = f5xc_child_tenant_manager.test.namespace
}
`, nsName, name))
}
