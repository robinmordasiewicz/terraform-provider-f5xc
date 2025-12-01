// Copyright (c) F5XC Community
// SPDX-License-Identifier: MPL-2.0

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/f5xc/terraform-provider-f5xc/internal/acctest"
)

func TestAccAppTypeDataSource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	// app_type resources must be created in the "shared" namespace
	rName := acctest.RandomName("tf-acc-test-apptype")
	resourceName := "f5xc_app_type.test"
	dataSourceName := "data.f5xc_app_type.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAppTypeDataSourceConfig_basic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "namespace", resourceName, "namespace"),
					resource.TestCheckResourceAttrPair(dataSourceName, "id", resourceName, "id"),
				),
			},
		},
	})
}

func testAccAppTypeDataSourceConfig_basic(name string) string {
	// app_type resources must be created in the "shared" namespace
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_app_type" "test" {
  name      = %[1]q
  namespace = "shared"
}

data "f5xc_app_type" "test" {
  depends_on = [f5xc_app_type.test]
  name       = f5xc_app_type.test.name
  namespace  = f5xc_app_type.test.namespace
}
`, name))
}
