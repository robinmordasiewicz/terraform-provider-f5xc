// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/f5xc/terraform-provider-f5xc/internal/acctest"
)

func TestAccFleetDataSource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-fleet")
	resourceName := "f5xc_fleet.test"
	dataSourceName := "data.f5xc_fleet.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccFleetDataSourceConfig_basic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "namespace", resourceName, "namespace"),
					resource.TestCheckResourceAttrPair(dataSourceName, "id", resourceName, "id"),
				),
			},
		},
	})
}

func testAccFleetDataSourceConfig_basic(name string) string {
	// Fleet resources must be created in system namespace per F5XC API requirements
	// The fleet_label is required for fleet creation
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_fleet" "test" {
  name        = %[1]q
  namespace   = "system"
  fleet_label = %[1]q
}

data "f5xc_fleet" "test" {
  name      = f5xc_fleet.test.name
  namespace = f5xc_fleet.test.namespace
}
`, name))
}
