// Copyright (c) F5XC Community
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"testing"
)

func TestNormalizeAPIURL(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedURL    string
		expectedNorm   bool
	}{
		{
			name:         "clean URL - no changes",
			input:        "https://console.ves.volterra.io",
			expectedURL:  "https://console.ves.volterra.io",
			expectedNorm: false,
		},
		{
			name:         "URL with trailing slash",
			input:        "https://console.ves.volterra.io/",
			expectedURL:  "https://console.ves.volterra.io",
			expectedNorm: true,
		},
		{
			name:         "URL with /api suffix",
			input:        "https://console.ves.volterra.io/api",
			expectedURL:  "https://console.ves.volterra.io",
			expectedNorm: true,
		},
		{
			name:         "URL with /api/ suffix",
			input:        "https://console.ves.volterra.io/api/",
			expectedURL:  "https://console.ves.volterra.io",
			expectedNorm: true,
		},
		{
			name:         "URL with multiple trailing slashes",
			input:        "https://console.ves.volterra.io///",
			expectedURL:  "https://console.ves.volterra.io",
			expectedNorm: true,
		},
		{
			name:         "tenant URL - no changes",
			input:        "https://nferreira.staging.volterra.us",
			expectedURL:  "https://nferreira.staging.volterra.us",
			expectedNorm: false,
		},
		{
			name:         "tenant URL with /api",
			input:        "https://nferreira.staging.volterra.us/api",
			expectedURL:  "https://nferreira.staging.volterra.us",
			expectedNorm: true,
		},
		{
			name:         "tenant URL with trailing slash",
			input:        "https://nferreira.staging.volterra.us/",
			expectedURL:  "https://nferreira.staging.volterra.us",
			expectedNorm: true,
		},
		{
			name:         "URL with /API (uppercase)",
			input:        "https://console.ves.volterra.io/API",
			expectedURL:  "https://console.ves.volterra.io",
			expectedNorm: true,
		},
		{
			name:         "URL with /Api (mixed case)",
			input:        "https://console.ves.volterra.io/Api",
			expectedURL:  "https://console.ves.volterra.io",
			expectedNorm: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotURL, gotNorm := normalizeAPIURL(tt.input)
			if gotURL != tt.expectedURL {
				t.Errorf("normalizeAPIURL(%q) URL = %q, want %q", tt.input, gotURL, tt.expectedURL)
			}
			if gotNorm != tt.expectedNorm {
				t.Errorf("normalizeAPIURL(%q) normalized = %v, want %v", tt.input, gotNorm, tt.expectedNorm)
			}
		})
	}
}
