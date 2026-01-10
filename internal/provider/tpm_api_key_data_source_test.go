// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider_test


import (
	"testing"
)

func TestAccTpmApiKeyDataSource_basic(t *testing.T) {
	t.Skip("Skipping: tpm_api_key requires TPM (Trusted Platform Module) infrastructure which is not available in standard test environments")
}
