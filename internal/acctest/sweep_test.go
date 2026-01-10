// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package acctest

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// TestMain handles setup and teardown for the acceptance test suite,
// including test resource sweepers.
//
// Run sweepers with:
//
//	TF_ACC=1 go test ./internal/acctest -v -sweep=all -timeout 30m
//
// Or to sweep specific resources:
//
//	TF_ACC=1 go test ./internal/acctest -v -sweep=f5xc_namespace -timeout 30m
func TestMain(m *testing.M) {
	resource.TestMain(m)
}
