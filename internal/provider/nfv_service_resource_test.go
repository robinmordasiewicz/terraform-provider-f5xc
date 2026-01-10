// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/f5xc/terraform-provider-f5xc/internal/acctest"
)

func TestAccNFVServiceResource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	// Skip: NFV Service requires either f5_big_ip_aws_service or palo_alto_fw_service
	// which need real AWS TGW Site infrastructure with deployed BIG-IP or Palo Alto firewalls.
	// This cannot be tested without actual cloud infrastructure.
	t.Skip("Skipping: nfv_service requires AWS TGW Site with BIG-IP or Palo Alto firewall infrastructure")

	rName := acctest.RandomName("tf-acc-test-nfv")
	nsName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_nfv_service.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		CheckDestroy: acctest.CheckResourceDestroyed("f5xc_nfv_service"),
		Steps: []resource.TestStep{
			{
				Config: testAccNFVServiceConfig_basic(nsName, rName),
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
				ImportStateIdFunc:       testAccNFVServiceImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccNFVServiceImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
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

func testAccNFVServiceConfig_basic(nsName, name string) string {
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

resource "f5xc_nfv_service" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name      = %[2]q
  namespace = f5xc_namespace.test.name
}
`, nsName, name))
}
