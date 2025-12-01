// Copyright (c) F5XC Community
// SPDX-License-Identifier: MPL-2.0

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/f5xc/terraform-provider-f5xc/internal/acctest"
)

func TestAccK8sClusterRoleBindingDataSource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test")
	resourceName := "f5xc_k8s_cluster_role_binding.test"
	dataSourceName := "data.f5xc_k8s_cluster_role_binding.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccK8sClusterRoleBindingDataSourceConfig_basic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "namespace", resourceName, "namespace"),
					resource.TestCheckResourceAttrPair(dataSourceName, "id", resourceName, "id"),
				),
			},
		},
	})
}

func testAccK8sClusterRoleBindingDataSourceConfig_basic(name string) string {
	// K8s cluster role binding must be in system namespace
	// Requires k8s_cluster_role reference and subjects
	// Using shared ves-io-admin-cluster-role from ves-io tenant
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_k8s_cluster_role_binding" "test" {
  name      = %[1]q
  namespace = "system"

  k8s_cluster_role {
    name      = "ves-io-admin-cluster-role"
    namespace = "shared"
    tenant    = "ves-io"
  }

  subjects {
    user = "test-user"
  }
}

data "f5xc_k8s_cluster_role_binding" "test" {
  depends_on = [f5xc_k8s_cluster_role_binding.test]
  name       = f5xc_k8s_cluster_role_binding.test.name
  namespace  = f5xc_k8s_cluster_role_binding.test.namespace
}
`, name))
}
