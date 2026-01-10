// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider_test


import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/f5xc/terraform-provider-f5xc/internal/acctest"
)

func TestAccDcClusterGroupDataSource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	// Skip: dc_cluster_group is used for connecting physical sites or App Stack sites
	// via underlay networks with MPLSoUDP encapsulation. Requires deployed sites with
	// Ingress/Egress Gateway (Two Interface) or App Stack Cluster configuration.
	t.Skip("Skipping: dc_cluster_group resource requires deployed site infrastructure with underlay network connectivity which is not available in standard acceptance tests")

	rName := acctest.RandomName("tf-acc-test")
	nsName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_dc_cluster_group.test"
	dataSourceName := "data.f5xc_dc_cluster_group.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccDcClusterGroupDataSourceConfig_basic(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "namespace", resourceName, "namespace"),
					resource.TestCheckResourceAttrPair(dataSourceName, "id", resourceName, "id"),
				),
			},
		},
	})
}


func testAccDcClusterGroupDataSourceConfig_basic(nsName, name string) string {
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

resource "f5xc_dc_cluster_group" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[2]q
  namespace  = f5xc_namespace.test.name

  type {
    data_plane_mesh {}
  }
}

data "f5xc_dc_cluster_group" "test" {
  depends_on = [f5xc_dc_cluster_group.test]
  name       = f5xc_dc_cluster_group.test.name
  namespace  = f5xc_dc_cluster_group.test.namespace
}
`, nsName, name))
}
