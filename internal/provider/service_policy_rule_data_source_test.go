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
// SERVICE POLICY RULE DATA SOURCE ACCEPTANCE TESTS
//
// These tests verify the f5xc_service_policy_rule data source implementation.
// Service Policy Rule requires system namespace and waf_action block.
//
// Run with:
//   TF_ACC=1 F5XC_API_URL="..." F5XC_API_P12_FILE="..." F5XC_P12_PASSWORD="..." \
//   go test -v ./internal/provider/ -run TestAccServicePolicyRuleDataSource -timeout 30m
// =============================================================================

func TestAccServicePolicyRuleDataSource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test")
	resourceName := "f5xc_service_policy_rule.test"
	dataSourceName := "data.f5xc_service_policy_rule.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccServicePolicyRuleDataSourceConfig_basic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "namespace", resourceName, "namespace"),
					resource.TestCheckResourceAttrPair(dataSourceName, "id", resourceName, "id"),
				),
			},
		},
	})
}

// testAccServicePolicyRuleDataSourceConfig_basic creates a service policy rule in system namespace
// with the required waf_action block, then reads it via data source.
func testAccServicePolicyRuleDataSourceConfig_basic(name string) string {
	// Service Policy Rule must be in system namespace and requires waf_action block
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_service_policy_rule" "test" {
  name      = %[1]q
  namespace = "system"

  waf_action {
    none {}
  }
}

data "f5xc_service_policy_rule" "test" {
  depends_on = [f5xc_service_policy_rule.test]
  name       = f5xc_service_policy_rule.test.name
  namespace  = f5xc_service_policy_rule.test.namespace
}
`, name))
}
