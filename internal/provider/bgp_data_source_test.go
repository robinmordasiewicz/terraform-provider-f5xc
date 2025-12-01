// Copyright (c) F5XC Community
// SPDX-License-Identifier: MPL-2.0

package provider_test

import (
	"testing"

	"github.com/f5xc/terraform-provider-f5xc/internal/acctest"
)

func TestAccBgpDataSource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	// Skip: BGP requires a real site reference which is infrastructure-dependent
	t.Skip("Skipping: bgp resource requires site infrastructure (CE/RE site) which is not available in acceptance tests")
}
