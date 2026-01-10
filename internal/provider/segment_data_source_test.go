// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider_test


import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/f5xc/terraform-provider-f5xc/internal/acctest"
)

func TestAccSegmentDataSource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test")
	resourceName := "f5xc_segment.test"
	dataSourceName := "data.f5xc_segment.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSegmentDataSourceConfig_basic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "namespace", resourceName, "namespace"),
					resource.TestCheckResourceAttrPair(dataSourceName, "id", resourceName, "id"),
				),
			},
		},
	})
}


func testAccSegmentDataSourceConfig_basic(name string) string {
	// Segment must be in system namespace
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_segment" "test" {
  name      = %[1]q
  namespace = "system"
}

data "f5xc_segment" "test" {
  depends_on = [f5xc_segment.test]
  name       = f5xc_segment.test.name
  namespace  = f5xc_segment.test.namespace
}
`, name))
}
