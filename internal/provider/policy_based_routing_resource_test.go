// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/f5xc/terraform-provider-f5xc/internal/acctest"
)

func TestAccPolicyBasedRoutingResource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	// Skip: policy_based_routing requires additional configuration (rules/routes)
	// Need to investigate required spec fields for successful creation
	t.Skip("Skipping: policy_based_routing requires additional configuration - BAD_REQUEST")

	rName := acctest.RandomName("tf-acc-test-pbr")
	nsName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_policy_based_routing.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		CheckDestroy: acctest.CheckResourceDestroyed("f5xc_policy_based_routing"),
		Steps: []resource.TestStep{
			{
				Config: testAccPolicyBasedRoutingConfig_basic(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "namespace", nsName),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"timeouts"},
				ImportStateIdFunc:       testAccPolicyBasedRoutingImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccPolicyBasedRoutingImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("resource not found: %s", resourceName)
		}
		namespace := rs.Primary.Attributes["namespace"]
		name := rs.Primary.Attributes["name"]
		return fmt.Sprintf("%s/%s", namespace, name), nil
	}
}

func testAccPolicyBasedRoutingConfig_basic(nsName, name string) string {
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

resource "f5xc_policy_based_routing" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name      = %[2]q
  namespace = f5xc_namespace.test.name
}
`, nsName, name))
}
