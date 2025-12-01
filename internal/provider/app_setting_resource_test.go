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

func TestAccAppSettingResource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	// app_setting can only be created in namespaces with namespace_type: "app"
	// Using "hipster-shop" which is an existing app-type namespace in the staging environment
	resourceName := "f5xc_app_setting.test"
	rName := acctest.RandomName("tf-test-appset")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_app_setting"),
		Steps: []resource.TestStep{
			{
				Config: testAccAppSettingConfig_basic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "namespace", "hipster-shop"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"timeouts", "app_type_settings"},
				ImportStateIdFunc:       testAccAppSettingImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccAppSettingImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
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

func testAccAppSettingConfig_basic(name string) string {
	// App setting requires app_type_settings with at least one app_type_ref
	// Must be created in a namespace with namespace_type: "app" (e.g., hipster-shop)
	return fmt.Sprintf(`
resource "f5xc_app_setting" "test" {
  name      = %[1]q
  namespace = "hipster-shop"

  app_type_settings {
    app_type_ref {
      name      = "hipster-shop"
      namespace = "shared"
    }
  }
}
`, name)
}
