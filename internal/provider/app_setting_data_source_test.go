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
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test")
	nsName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_app_setting.test"
	dataSourceName := "data.f5xc_app_setting.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccAppSettingDataSourceConfig_basic(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "namespace", resourceName, "namespace"),
					resource.TestCheckResourceAttrPair(dataSourceName, "id", resourceName, "id"),
				),
			},
		},
	})
}


func testAccAppSettingDataSourceConfig_basic(nsName, name string) string {
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

resource "f5xc_app_type" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = "${%[2]q}-apptype"
  namespace  = f5xc_namespace.test.name
  business_logic_markup_setting {
    enable = true
  }
}

resource "f5xc_app_setting" "test" {
  depends_on = [time_sleep.wait_for_namespace, f5xc_app_type.test]
  name       = %[2]q
  namespace  = f5xc_namespace.test.name
  app_type_refs {
    name      = f5xc_app_type.test.name
    namespace = f5xc_namespace.test.name
  }
}

data "f5xc_app_setting" "test" {
  depends_on = [f5xc_app_setting.test]
  name       = f5xc_app_setting.test.name
  namespace  = f5xc_app_setting.test.namespace
}
`, nsName, name))
}
