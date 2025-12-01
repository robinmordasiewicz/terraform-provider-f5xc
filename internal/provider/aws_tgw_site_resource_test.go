// Copyright (c) F5XC Community
// SPDX-License-Identifier: MPL-2.0

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/f5xc/terraform-provider-f5xc/internal/acctest"
)

// TestAccAWSTGWSiteResource_basic validates basic AWS TGW Site resource operations
// Note: This test creates actual AWS infrastructure and requires:
// - Valid AWS credentials configured
// - F5 XC API access with AWS integration
// - Significant time for AWS resource provisioning (10-20 minutes)
func TestAccAWSTGWSiteResource_basic(t *testing.T) {
	t.Skip("Skipping: requires AWS cloud credentials and cloud_credentials resource")
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-tgw")
	nsName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_aws_tgw_site.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source: "hashicorp/time",
			},
		},
		CheckDestroy: acctest.CheckResourceDestroyed("f5xc_aws_tgw_site"),
		Steps: []resource.TestStep{
			{
				Config: testAccAWSTGWSiteConfig_basic(nsName, rName),
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
				ImportStateIdFunc:       testAccAWSTGWSiteImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccAWSTGWSiteImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
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

func testAccAWSTGWSiteConfig_basic(nsName, name string) string {
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

# Note: This is a minimal configuration for testing only
# Production deployments require extensive AWS configuration including:
# - aws_parameters block with region, VPC, credentials, subnets, etc.
# - Proper networking configuration
# - Security groups and access controls
resource "f5xc_aws_tgw_site" "test" {
  depends_on = [time_sleep.wait_for_namespace]

  name      = %[2]q
  namespace = f5xc_namespace.test.name

  # Minimal configuration - would need extensive aws_parameters block for real deployment
}
`, nsName, name))
}
