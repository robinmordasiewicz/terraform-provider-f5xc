// Copyright (c) F5XC Community
// SPDX-License-Identifier: MPL-2.0

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/f5xc/terraform-provider-f5xc/internal/acctest"
)

func TestAccAppSettingDataSource_basic(t *testing.T) {
	t.Skip("Skipping: requires hipster-shop namespace with namespace_type 'app' and specific app_type configuration not available in standard test environments")
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	// app_setting can only be created in namespaces with namespace_type: "app"
	// Using "hipster-shop" which is an existing app-type namespace in the staging environment
	rName := acctest.RandomName("tf-acc-test-appset")
	resourceName := "f5xc_app_setting.test"
	dataSourceName := "data.f5xc_app_setting.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAppSettingDataSourceConfig_basic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "namespace", resourceName, "namespace"),
					resource.TestCheckResourceAttrPair(dataSourceName, "id", resourceName, "id"),
				),
			},
		},
	})
}

func testAccAppSettingDataSourceConfig_basic(name string) string {
	// App setting requires app_type_settings with at least one app_type_ref
	// Must be created in a namespace with namespace_type: "app" (e.g., hipster-shop)
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
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

data "f5xc_app_setting" "test" {
  depends_on = [f5xc_app_setting.test]
  name       = f5xc_app_setting.test.name
  namespace  = f5xc_app_setting.test.namespace
}
`, name))
}
