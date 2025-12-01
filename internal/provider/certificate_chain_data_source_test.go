// Copyright (c) F5XC Community
// SPDX-License-Identifier: MPL-2.0

package provider_test


import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/f5xc/terraform-provider-f5xc/internal/acctest"
)

func TestAccCertificateChainDataSource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test")
	nsName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_certificate_chain.test"
	dataSourceName := "data.f5xc_certificate_chain.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccCertificateChainDataSourceConfig_basic(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "namespace", resourceName, "namespace"),
					resource.TestCheckResourceAttrPair(dataSourceName, "id", resourceName, "id"),
				),
			},
		},
	})
}


func testAccCertificateChainDataSourceConfig_basic(nsName, name string) string {
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

resource "f5xc_certificate_chain" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[2]q
  namespace  = f5xc_namespace.test.name
  certificate_url = "string:///LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUNxRENDQVpBQ0NRQ21SN0dCREN4dXRqQU5CZ2txaGtpRzl3MEJBUXNGQURBV01SUXdFZ1lEVlFRRERBdDAKWlhOMExXTmxjblF1WTI5dE1CNFhEVEkxTURFeE1UQXdNREF3TUZvWERUSTJNREV4TVRBd01EQXdNRm93RmpFVQpNQklHQTFVRUF3d0xkR1Z6ZEMxalpYSjBMbU52YlRDQ0FTSXdEUVlKS29aSWh2Y05BUUVCQlFBRGdnRVBBRENDCkFRb0NnZ0VCQU1wQklYdVhJcmZJT1ZoQndxWnBKNXZsRlk4cEtndnhOdHQ4elEyRWtLVmhmMk1mOXc4bE5RdTIKK3FkZm5YNW5YRUZXS2dhVnpqYUZGZHFOZ0xHaUlRaWQvWVcwa0dDd3RHZURlMGE4OFg1UW5oeTRHdkdHdEhyVgp6V3ZoTlVHaXJSSnJiOGUwemxQdWxGeTA2dEl4ZUdBb1FXT09KU0p1NUtsOWJJenFYZ0hHZG9LWXd0VG9ERzc5CnFXN3VZNGZNVnhSM2dBYm5CSk84eGxMR3dQNndpSU5PQTJObFNUc1g5TmliOHR3b0NuSnpQMm1wOGFOd1VYZWsKNHBiZW1HZ2dxY2ZBM0ROaXNmRUNhLy9SRzFYZlFrNjI1WGtrb3h5UlBuaG9JSU5kVmFhY0RRTDVYZU16akRuRwphaUZaaGZCQWJHWlB3dk1MRnVEN3JPYWllRGRGckZjQ0F3RUFBVEFOQmdrcWhraUc5dzBCQVFzRkFBT0NBUUVBCmdEZjVPMk1DT0JCRURQZnFkRnpiMjBLUm9nNUp3bmNHcklwSVQwQmc1S0M5VElQRDUxdkJtY2w0ZWZxT3Q2TWkKT0lSeEFKa3JvTFFIRW5xMnBCdkxVM0dwT1IxQ1V1STYvRnBmb2toaFlEN1lsSnR4RE9mZW9JT0NOTk0rMXgyaApmTXY1NHVhV0V4R3dwQnVQL1M1dVZ6TU95aVVFcklXejIyUFpsZndIS3ZhdGNuclFBTjhLNmpvMUp1dkc4RmdsCk1VcjNVa3VMWEVYMHBNd1p0WWVoUjJVYlN0ZEVVT1puSTdFalp5TTU5WndFVmRrK1NWYmRVNjQzOEdYVHlDRjcKeGxRTjk3YmZOUnBkZG5JYWpyYU52SXJrclBYVnhERXhDOEo4NzQzRE9ZenZ2bTZNVW1vZW9ML3FuRVRiTGV5RApqQjJWS0MxTi9NeEFjUUdGOWRueHpBPT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="
}

data "f5xc_certificate_chain" "test" {
  depends_on = [f5xc_certificate_chain.test]
  name       = f5xc_certificate_chain.test.name
  namespace  = f5xc_certificate_chain.test.namespace
}
`, nsName, name))
}
