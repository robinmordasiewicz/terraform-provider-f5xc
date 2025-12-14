// Copyright (c) F5XC Community
// SPDX-License-Identifier: MPL-2.0

// provider_helpers.go - Manually maintained helper functions for the provider.
// This file is NOT auto-generated and contains utility functions used by
// the provider implementation.

package provider

import (
	"strings"
)

// normalizeAPIURL cleans up the API URL to ensure consistent formatting.
// It removes trailing slashes and the /api suffix if present, since API paths
// already include the /api prefix (e.g., /api/web/namespaces).
func normalizeAPIURL(url string) (string, bool) {
	original := url

	// Remove trailing slashes
	url = strings.TrimRight(url, "/")

	// Remove /api suffix (case-insensitive check, preserve original case in removal)
	if strings.HasSuffix(strings.ToLower(url), "/api") {
		url = url[:len(url)-4]
	}

	// Remove any trailing slashes that might have been before /api
	url = strings.TrimRight(url, "/")

	return url, url != original
}

// filterSystemLabels removes F5 XC system-managed labels (ves.io/*) from the label map.
// These labels are injected by the platform and should not be managed by Terraform.
// nolint:unused // Used by generated resource/data source Read methods
func filterSystemLabels(labels map[string]string) map[string]string {
	filtered := make(map[string]string)
	for k, v := range labels {
		if !strings.HasPrefix(k, "ves.io/") {
			filtered[k] = v
		}
	}
	return filtered
}
