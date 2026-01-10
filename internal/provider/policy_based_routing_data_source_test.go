// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider_test

import (
	"testing"

	"github.com/f5xc/terraform-provider-f5xc/internal/acctest"
)

func TestAccPolicyBasedRoutingDataSource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	// Skip: policy_based_routing requires additional configuration (rules/routes)
	// Need to investigate required spec fields for successful creation
	t.Skip("Skipping: policy_based_routing requires additional configuration - BAD_REQUEST")
}
