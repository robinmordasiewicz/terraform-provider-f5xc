// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/f5xc/terraform-provider-f5xc/internal/acctest"
)

func TestAccCertificateDataSource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test")
	nsName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_certificate.test"
	dataSourceName := "data.f5xc_certificate.test"

	// Generate test certificates dynamically for CI/CD compatibility
	certs := acctest.MustGenerateTestCertificates()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccCertificateDataSourceConfig_basic(nsName, rName, certs),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "namespace", resourceName, "namespace"),
					resource.TestCheckResourceAttrPair(dataSourceName, "id", resourceName, "id"),
				),
			},
		},
	})
}

func testAccCertificateDataSourceConfig_basic(nsName, name string, certs *acctest.TestCertificates) string {
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

resource "f5xc_certificate" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[2]q
  namespace  = f5xc_namespace.test.name

  certificate_url = "string:///%[3]s"

  private_key {
    clear_secret_info {
      url = "string:///%[4]s"
    }
  }

  disable_ocsp_stapling {}
}

data "f5xc_certificate" "test" {
  depends_on = [f5xc_certificate.test]
  name       = f5xc_certificate.test.name
  namespace  = f5xc_certificate.test.namespace
}
`, nsName, name, certs.ServerCertBase64, certs.ServerKeyBase64))
}
