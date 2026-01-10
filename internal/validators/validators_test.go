// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package validators

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestNamePattern(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		matches bool
	}{
		{"valid simple name", "myname", true},
		{"valid with hyphens", "my-resource-name", true},
		{"valid with numbers", "my-resource-123", true},
		{"valid single char", "a", true},
		{"valid two chars", "ab", true},
		{"invalid starts with number", "1name", false},
		{"invalid starts with hyphen", "-name", false},
		{"invalid ends with hyphen", "name-", false},
		{"invalid uppercase", "MyName", false},
		{"invalid spaces", "my name", false},
		{"invalid underscores", "my_name", false},
		{"empty string", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NamePattern.MatchString(tt.input)
			if result != tt.matches {
				t.Errorf("NamePattern.MatchString(%q) = %v, expected %v", tt.input, result, tt.matches)
			}
		})
	}
}

func TestNamespacePattern(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		matches bool
	}{
		{"valid system namespace", "system", true},
		{"valid custom namespace", "my-namespace", true},
		{"valid with numbers", "namespace-123", true},
		{"invalid uppercase", "System", false},
		{"invalid starts with hyphen", "-namespace", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NamespacePattern.MatchString(tt.input)
			if result != tt.matches {
				t.Errorf("NamespacePattern.MatchString(%q) = %v, expected %v", tt.input, result, tt.matches)
			}
		})
	}
}

func TestDomainPattern(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		matches bool
	}{
		{"valid domain", "example.com", true},
		{"valid subdomain", "sub.example.com", true},
		{"valid wildcard", "*.example.com", true},
		{"valid with hyphens", "my-site.example.com", true},
		{"invalid double dot", "example..com", false},
		{"invalid starts with dot", ".example.com", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DomainPattern.MatchString(tt.input)
			if result != tt.matches {
				t.Errorf("DomainPattern.MatchString(%q) = %v, expected %v", tt.input, result, tt.matches)
			}
		})
	}
}

func TestNameValidator(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expectError bool
	}{
		{"valid name", "my-resource", false},
		{"valid single char", "a", false},
		{"invalid too long", string(make([]byte, 65)), true},
		{"invalid format", "Invalid_Name", true},
		{"null value", "", false}, // null values are allowed
	}

	ctx := context.Background()
	v := NameValidator()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := validator.StringRequest{
				Path:        path.Root("name"),
				ConfigValue: types.StringValue(tt.input),
			}
			if tt.input == "" {
				req.ConfigValue = types.StringNull()
			}

			resp := &validator.StringResponse{}
			v.ValidateString(ctx, req, resp)

			hasError := resp.Diagnostics.HasError()
			if hasError != tt.expectError {
				t.Errorf("NameValidator for %q: hasError = %v, expected %v", tt.input, hasError, tt.expectError)
				if hasError {
					for _, d := range resp.Diagnostics.Errors() {
						t.Logf("Error: %s - %s", d.Summary(), d.Detail())
					}
				}
			}
		})
	}
}

func TestNamespaceValidator(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expectError bool
	}{
		{"valid system", "system", false},
		{"valid custom", "my-namespace", false},
		{"invalid uppercase", "System", true},
	}

	ctx := context.Background()
	v := NamespaceValidator()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := validator.StringRequest{
				Path:        path.Root("namespace"),
				ConfigValue: types.StringValue(tt.input),
			}

			resp := &validator.StringResponse{}
			v.ValidateString(ctx, req, resp)

			hasError := resp.Diagnostics.HasError()
			if hasError != tt.expectError {
				t.Errorf("NamespaceValidator for %q: hasError = %v, expected %v", tt.input, hasError, tt.expectError)
			}
		})
	}
}

func TestDomainValidator(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expectError bool
	}{
		{"valid domain", "example.com", false},
		{"valid wildcard", "*.example.com", false},
		{"invalid format", "not a domain", true},
	}

	ctx := context.Background()
	v := DomainValidator()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := validator.StringRequest{
				Path:        path.Root("domain"),
				ConfigValue: types.StringValue(tt.input),
			}

			resp := &validator.StringResponse{}
			v.ValidateString(ctx, req, resp)

			hasError := resp.Diagnostics.HasError()
			if hasError != tt.expectError {
				t.Errorf("DomainValidator for %q: hasError = %v, expected %v", tt.input, hasError, tt.expectError)
			}
		})
	}
}

func TestPortValidator(t *testing.T) {
	tests := []struct {
		name        string
		input       int64
		expectError bool
	}{
		{"valid port 80", 80, false},
		{"valid port 443", 443, false},
		{"valid port 1", 1, false},
		{"valid port 65535", 65535, false},
		{"invalid port 0", 0, true},
		{"invalid port negative", -1, true},
		{"invalid port too high", 65536, true},
	}

	ctx := context.Background()
	v := PortValidator()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := validator.Int64Request{
				Path:        path.Root("port"),
				ConfigValue: types.Int64Value(tt.input),
			}

			resp := &validator.Int64Response{}
			v.ValidateInt64(ctx, req, resp)

			hasError := resp.Diagnostics.HasError()
			if hasError != tt.expectError {
				t.Errorf("PortValidator for %d: hasError = %v, expected %v", tt.input, hasError, tt.expectError)
			}
		})
	}
}

func TestNonEmptyStringValidator(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expectError bool
	}{
		{"valid string", "hello", false},
		{"valid with spaces", "hello world", false},
		{"invalid empty", "", true},
		{"invalid whitespace only", "   ", true},
		{"invalid tabs only", "\t\t", true},
	}

	ctx := context.Background()
	v := NonEmptyStringValidator()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := validator.StringRequest{
				Path:        path.Root("value"),
				ConfigValue: types.StringValue(tt.input),
			}

			resp := &validator.StringResponse{}
			v.ValidateString(ctx, req, resp)

			hasError := resp.Diagnostics.HasError()
			if hasError != tt.expectError {
				t.Errorf("NonEmptyStringValidator for %q: hasError = %v, expected %v", tt.input, hasError, tt.expectError)
			}
		})
	}
}

func TestLabelKeyValidator(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expectError bool
	}{
		{"valid simple key", "app", false},
		{"valid with prefix", "kubernetes.io/name", false},
		{"valid with dots", "app.kubernetes.io/name", false},
		{"invalid too long", string(make([]byte, 254)), true},
	}

	ctx := context.Background()
	v := LabelKeyValidator()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := validator.StringRequest{
				Path:        path.Root("label_key"),
				ConfigValue: types.StringValue(tt.input),
			}

			resp := &validator.StringResponse{}
			v.ValidateString(ctx, req, resp)

			hasError := resp.Diagnostics.HasError()
			if hasError != tt.expectError {
				t.Errorf("LabelKeyValidator for %q: hasError = %v, expected %v", tt.input, hasError, tt.expectError)
			}
		})
	}
}

func TestRequiredNameValidators(t *testing.T) {
	validators := RequiredNameValidators()
	if len(validators) != 2 {
		t.Errorf("RequiredNameValidators() returned %d validators, expected 2", len(validators))
	}
}

func TestRequiredNamespaceValidators(t *testing.T) {
	validators := RequiredNamespaceValidators()
	if len(validators) != 2 {
		t.Errorf("RequiredNamespaceValidators() returned %d validators, expected 2", len(validators))
	}
}

func TestOptionalDomainValidators(t *testing.T) {
	validators := OptionalDomainValidators()
	if len(validators) != 1 {
		t.Errorf("OptionalDomainValidators() returned %d validators, expected 1", len(validators))
	}
}

func TestValidatorDescriptions(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name      string
		validator interface {
			Description(context.Context) string
			MarkdownDescription(context.Context) string
		}
	}{
		{"NameValidator", NameValidator()},
		{"NamespaceValidator", NamespaceValidator()},
		{"DomainValidator", DomainValidator()},
		{"LabelKeyValidator", LabelKeyValidator()},
		{"NonEmptyStringValidator", NonEmptyStringValidator()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			desc := tt.validator.Description(ctx)
			mdDesc := tt.validator.MarkdownDescription(ctx)

			if desc == "" {
				t.Errorf("%s.Description() returned empty string", tt.name)
			}
			if mdDesc == "" {
				t.Errorf("%s.MarkdownDescription() returned empty string", tt.name)
			}
		})
	}
}
