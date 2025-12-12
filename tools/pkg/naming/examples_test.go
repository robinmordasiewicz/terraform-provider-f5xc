package naming

import (
	"testing"
)

func TestNormalizeExampleNames(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		// Exact replacements from OpenAPI specs
		{"my-bucket", "my-bucket", "example-bucket"},
		{"my-log-bucket", "my-log-bucket", "example-log-bucket"},
		{"my-resources", "my-resources", "example-resources"},
		{"my-tsig-key", "my-tsig-key", "example-tsig-key"},
		{"my-website", "my-website", "example-website"},

		// Additional common patterns
		{"my-file", "my-file", "example-file"},
		{"my-key", "my-key", "example-key"},
		{"my-ns", "my-ns", "example-ns"},
		{"my-namespace", "my-namespace", "example-namespace"},
		{"my-app", "my-app", "example-app"},
		{"my-config", "my-config", "example-config"},
		{"my-secret", "my-secret", "example-secret"},
		{"my-cert", "my-cert", "example-cert"},
		{"my-policy", "my-policy", "example-policy"},
		{"my-pool", "my-pool", "example-pool"},
		{"my-origin", "my-origin", "example-origin"},
		{"my-site", "my-site", "example-site"},
		{"my-cluster", "my-cluster", "example-cluster"},

		// Prefix-based replacement (fallback for unregistered patterns)
		{"my-custom-name fallback", "my-custom-name", "example-custom-name"},
		{"my-other-resource fallback", "my-other-resource", "example-other-resource"},

		// Underscore variants
		{"my_bucket underscore", "my_bucket", "example_bucket"},
		{"my_custom_name underscore", "my_custom_name", "example_custom_name"},

		// Case insensitivity
		{"MY-BUCKET uppercase", "MY-BUCKET", "example-bucket"},
		{"My-Bucket mixed case", "My-Bucket", "example-bucket"},
		{"MY_BUCKET uppercase underscore", "MY_BUCKET", "example_bucket"},

		// Domain replacements
		{"my.example.com domain", "my.example.com", "example.example.com"},
		{"my.tenant.domain", "my.tenant.domain", "example.tenant.domain"},

		// In context (description text)
		{"in description hyphen", "Configure the bucket name, e.g. my-bucket",
			"Configure the bucket name, e.g. example-bucket"},
		{"multiple in text", "Use my-bucket and my-file for storage.",
			"Use example-bucket and example-file for storage."},
		{"with domain in URL", "https://my.tenant.domain/api/v1",
			"https://example.tenant.domain/api/v1"},

		// Complex example from OpenAPI specs
		{"x-ves-example format", "my-file, shared/my-file, my-ns/my-file",
			"example-file, shared/example-file, example-ns/example-file"},

		// No change cases
		{"already example- prefix", "example-bucket", "example-bucket"},
		{"unrelated text", "this is normal text", "this is normal text"},
		{"empty string", "", ""},
		{"contains my but not prefix", "my configuration", "my configuration"},
		{"myapp not my-app", "myapp", "myapp"}, // no separator

		// Edge cases
		{"multiple patterns", "my-bucket and my-file and my-key",
			"example-bucket and example-file and example-key"},
		{"at start of string", "my-bucket is the name", "example-bucket is the name"},
		{"at end of string", "name is my-bucket", "name is example-bucket"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NormalizeExampleNames(tt.input)
			if result != tt.expected {
				t.Errorf("NormalizeExampleNames(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestNormalizeExampleNames_Idempotent(t *testing.T) {
	inputs := []string{
		"my-bucket",
		"my-log-bucket",
		"https://my.tenant.domain/api/test",
		"Configure my-file storage",
		"example-bucket", // already normalized
		"Use my-custom-resource for testing",
	}

	for _, input := range inputs {
		t.Run(input, func(t *testing.T) {
			first := NormalizeExampleNames(input)
			second := NormalizeExampleNames(first)
			if first != second {
				t.Errorf("NormalizeExampleNames is not idempotent:\n  input:  %q\n  first:  %q\n  second: %q",
					input, first, second)
			}
		})
	}
}

func TestNormalizeExampleInDescription(t *testing.T) {
	input := "The bucket name. Default is my-bucket. See my-log-bucket for logs."
	expected := "The bucket name. Default is example-bucket. See example-log-bucket for logs."

	result := NormalizeExampleInDescription(input)
	if result != expected {
		t.Errorf("NormalizeExampleInDescription(%q) = %q, want %q", input, result, expected)
	}
}

func TestNormalizeXVesExample(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"simple", "my-bucket", "example-bucket"},
		{"comma separated", "my-file, shared/my-file, my-ns/my-file",
			"example-file, shared/example-file, example-ns/example-file"},
		{"URL with domain", "https://my.tenant.domain/api/object_store/namespaces/my-ns/stored_objects",
			"https://example.tenant.domain/api/object_store/namespaces/example-ns/stored_objects"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NormalizeXVesExample(tt.input)
			if result != tt.expected {
				t.Errorf("NormalizeXVesExample(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

// TestExampleNameReplacementsComplete verifies all patterns found in OpenAPI specs are covered
func TestExampleNameReplacementsComplete(t *testing.T) {
	// Patterns found in docs/specifications/api/*.json
	patternsFromOpenAPI := []string{
		"my-bucket",
		"my-log-bucket",
		"my-resources",
		"my-tsig-key",
		"my-website",
	}

	for _, pattern := range patternsFromOpenAPI {
		if _, exists := ExampleNameReplacements[pattern]; !exists {
			t.Errorf("OpenAPI pattern %q is not in ExampleNameReplacements", pattern)
		}
	}
}

// TestExampleDomainReplacementsComplete verifies domain patterns are covered
func TestExampleDomainReplacementsComplete(t *testing.T) {
	// Domain patterns found in OpenAPI specs
	domainsFromOpenAPI := []string{
		"my.example.com",
		"my.tenant.domain",
	}

	for _, domain := range domainsFromOpenAPI {
		if _, exists := ExampleDomainReplacements[domain]; !exists {
			t.Errorf("OpenAPI domain pattern %q is not in ExampleDomainReplacements", domain)
		}
	}
}
