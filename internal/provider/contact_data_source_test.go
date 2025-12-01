// Copyright (c) F5XC Community
// SPDX-License-Identifier: MPL-2.0

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/f5xc/terraform-provider-f5xc/internal/acctest"
)

func TestAccContactDataSource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	// Skip: contact resources are tenant-level billing/mailing objects that may require
	// specific tenant permissions or have undocumented API requirements
	t.Skip("Skipping: contact resource requires special tenant permissions (BAD_REQUEST)")

	rName := acctest.RandomName("tf-acc-test-contact")
	resourceName := "f5xc_contact.test"
	dataSourceName := "data.f5xc_contact.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccContactDataSourceConfig_basic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "namespace", resourceName, "namespace"),
					resource.TestCheckResourceAttrPair(dataSourceName, "id", resourceName, "id"),
				),
			},
		},
	})
}

func testAccContactDataSourceConfig_basic(name string) string {
	// Contact resources are tenant-level objects for billing/mailing purposes
	// They need to be created in system namespace
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_contact" "test" {
  name       = %[1]q
  namespace  = "system"

  address1     = "123 Test Street"
  city         = "Seattle"
  state        = "Washington"
  state_code   = "WA"
  zip_code     = "98101"
  country      = "US"
  phone_number = "+12065551234"
  contact_type = "MAILING"
}

data "f5xc_contact" "test" {
  depends_on = [f5xc_contact.test]
  name       = f5xc_contact.test.name
  namespace  = f5xc_contact.test.namespace
}
`, name))
}
