// Copyright (c) F5XC Community
// SPDX-License-Identifier: MPL-2.0

package provider_test


import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/f5xc/terraform-provider-f5xc/internal/acctest"
)

func TestAccCrlDataSource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test")
	nsName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_crl.test"
	dataSourceName := "data.f5xc_crl.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccCrlDataSourceConfig_basic(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "namespace", resourceName, "namespace"),
					resource.TestCheckResourceAttrPair(dataSourceName, "id", resourceName, "id"),
				),
			},
		},
	})
}


func testAccCrlDataSourceConfig_basic(nsName, name string) string {
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

resource "f5xc_crl" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[2]q
  namespace  = f5xc_namespace.test.name
  crl_url = "string:///LS0tLS1CRUdJTiBYNTA5IENSTC0tLS0tCk1JSUJURENDQVFNQ0FRRXdEUVlKS29aSWh2Y05BUUVMQlFBd0ZqRVVNQklHQTFVRUF3d0xkR1Z6ZEMxalpYSjAKTG1OdmJSY05NalV3TVRFeE1EQXdNREF3V2hjTk1qWXdNVEV4TURBd01EQXdXakFOQmdrcWhraUc5dzBCQVFzRgpBQU9DQVFFQWdEZjVPMk1DT0JCRURQZnFkRnpiMjBLUm9nNUp3bmNHcklwSVQwQmc1S0M5VElQRDUxdkJtY2w0CmVmcU90Nk1pT0lSeEFKa3JvTFFIRW5xMnBCdkxVM0dwT1IxQ1V1STYvRnBmb2toaFlEN1lsSnR4RE9mZW9JT0MKTk5NKzF4MmhmTXY1NHVhV0V4R3dwQnVQL1M1dVZ6TU95aVVFcklXejIyUFpsZndIS3ZhdGNuclFBTjhLNmpvMQpKdXZHOEZnbE1VcjNVa3VMWEVYMHBNd1p0WWVoUjJVYlN0ZEVVT1puSTdFalp5TTU5WndFVmRrK1NWYmRVNjQzCjhHWFR5Q0Y3eGxRTjk3YmZOUnBkZG5JYWpyYU52SXJrclBYVnhERXhDOEo4NzQzRE9ZenZ2bTZNVW1vZW9ML3EKbkVUYkxleURqQjJWS0MxTi9NeEFjUUdGOWRueHpBPT0KLS0tLS1FTkQgWDUwOSBDUkwtLS0tLQo="
}

data "f5xc_crl" "test" {
  depends_on = [f5xc_crl.test]
  name       = f5xc_crl.test.name
  namespace  = f5xc_crl.test.namespace
}
`, nsName, name))
}
