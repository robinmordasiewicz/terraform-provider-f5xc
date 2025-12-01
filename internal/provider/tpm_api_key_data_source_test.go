// Copyright (c) F5XC Community
// SPDX-License-Identifier: MPL-2.0

package provider_test


import (
	"testing"
)

func TestAccTpmApiKeyDataSource_basic(t *testing.T) {
	t.Skip("Skipping: tpm_api_key requires TPM (Trusted Platform Module) infrastructure which is not available in standard test environments")
}
