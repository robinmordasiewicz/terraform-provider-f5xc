// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/f5xc/terraform-provider-f5xc/internal/acctest"
)

func TestAccFastAclDataSource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-facl")
	resourceName := "f5xc_fast_acl.test"
	dataSourceName := "data.f5xc_fast_acl.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccFastAclDataSourceConfig_basic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "namespace", resourceName, "namespace"),
					resource.TestCheckResourceAttrPair(dataSourceName, "id", resourceName, "id"),
				),
			},
		},
	})
}

func testAccFastAclDataSourceConfig_basic(name string) string {
	// Fast ACL must be in system namespace and requires site_choice (re_acl or site_acl)
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_fast_acl" "test" {
  name      = %[1]q
  namespace = "system"

  re_acl {}
}

data "f5xc_fast_acl" "test" {
  name      = f5xc_fast_acl.test.name
  namespace = f5xc_fast_acl.test.namespace
}
`, name))
}
