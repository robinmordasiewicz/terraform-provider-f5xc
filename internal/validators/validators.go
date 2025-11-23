// Copyright (c) F5XC Community
// SPDX-License-Identifier: MPL-2.0

// Package validators provides custom validators for F5 XC Terraform resources
// following Terraform Plugin Framework best practices.
package validators

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// Common regex patterns for F5 XC resources
var (
	// NamePattern validates F5 XC resource names: lowercase alphanumeric with hyphens, 1-64 chars
	NamePattern = regexp.MustCompile(`^[a-z][a-z0-9-]{0,62}[a-z0-9]$|^[a-z]$`)

	// NamespacePattern validates F5 XC namespaces
	NamespacePattern = regexp.MustCompile(`^[a-z][a-z0-9-]{0,62}[a-z0-9]$|^system$`)

	// LabelKeyPattern validates Kubernetes-style label keys
	LabelKeyPattern = regexp.MustCompile(`^([a-zA-Z0-9]([a-zA-Z0-9-_.]*[a-zA-Z0-9])?/)?[a-zA-Z0-9]([a-zA-Z0-9-_.]*[a-zA-Z0-9])?$`)

	// DomainPattern validates domain names
	DomainPattern = regexp.MustCompile(`^(\*\.)?[a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(\.[a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`)
)

// NameValidator returns a validator for F5 XC resource names
func NameValidator() validator.String {
	return &nameValidator{}
}

type nameValidator struct{}

func (v nameValidator) Description(ctx context.Context) string {
	return "name must be lowercase alphanumeric with hyphens, start with a letter, and be 1-64 characters"
}

func (v nameValidator) MarkdownDescription(ctx context.Context) string {
	return "name must be lowercase alphanumeric with hyphens, start with a letter, and be 1-64 characters"
}

func (v nameValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	value := req.ConfigValue.ValueString()

	if len(value) < 1 || len(value) > 64 {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid Name Length",
			fmt.Sprintf("Name must be between 1 and 64 characters, got %d characters.", len(value)),
		)
		return
	}

	if !NamePattern.MatchString(value) {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid Name Format",
			fmt.Sprintf("Name '%s' must start with a lowercase letter, contain only lowercase letters, numbers, and hyphens, and not end with a hyphen.", value),
		)
	}
}

// NamespaceValidator returns a validator for F5 XC namespaces
func NamespaceValidator() validator.String {
	return &namespaceValidator{}
}

type namespaceValidator struct{}

func (v namespaceValidator) Description(ctx context.Context) string {
	return "namespace must be 'system' or lowercase alphanumeric with hyphens"
}

func (v namespaceValidator) MarkdownDescription(ctx context.Context) string {
	return "namespace must be `system` or lowercase alphanumeric with hyphens"
}

func (v namespaceValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	value := req.ConfigValue.ValueString()

	if !NamespacePattern.MatchString(value) {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid Namespace Format",
			fmt.Sprintf("Namespace '%s' must be 'system' or start with a lowercase letter and contain only lowercase letters, numbers, and hyphens.", value),
		)
	}
}

// DomainValidator returns a validator for domain names
func DomainValidator() validator.String {
	return &domainValidator{}
}

type domainValidator struct{}

func (v domainValidator) Description(ctx context.Context) string {
	return "must be a valid domain name, optionally starting with *. for wildcard"
}

func (v domainValidator) MarkdownDescription(ctx context.Context) string {
	return "must be a valid domain name, optionally starting with `*.` for wildcard"
}

func (v domainValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	value := req.ConfigValue.ValueString()

	if !DomainPattern.MatchString(value) {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid Domain Format",
			fmt.Sprintf("Domain '%s' is not a valid domain name. Use format like 'example.com' or '*.example.com' for wildcards.", value),
		)
	}
}

// LabelKeyValidator returns a validator for Kubernetes-style label keys
func LabelKeyValidator() validator.String {
	return &labelKeyValidator{}
}

type labelKeyValidator struct{}

func (v labelKeyValidator) Description(ctx context.Context) string {
	return "must be a valid label key following Kubernetes naming conventions"
}

func (v labelKeyValidator) MarkdownDescription(ctx context.Context) string {
	return "must be a valid label key following Kubernetes naming conventions"
}

func (v labelKeyValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	value := req.ConfigValue.ValueString()

	if len(value) > 253 {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid Label Key Length",
			fmt.Sprintf("Label key must be at most 253 characters, got %d characters.", len(value)),
		)
		return
	}

	if !LabelKeyPattern.MatchString(value) {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid Label Key Format",
			fmt.Sprintf("Label key '%s' must follow Kubernetes naming conventions.", value),
		)
	}
}

// PortValidator returns a validator for port numbers (1-65535)
func PortValidator() validator.Int64 {
	return &portValidator{}
}

type portValidator struct{}

func (v portValidator) Description(ctx context.Context) string {
	return "must be a valid port number between 1 and 65535"
}

func (v portValidator) MarkdownDescription(ctx context.Context) string {
	return "must be a valid port number between 1 and 65535"
}

func (v portValidator) ValidateInt64(ctx context.Context, req validator.Int64Request, resp *validator.Int64Response) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	value := req.ConfigValue.ValueInt64()

	if value < 1 || value > 65535 {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid Port Number",
			fmt.Sprintf("Port must be between 1 and 65535, got %d.", value),
		)
	}
}

// NonEmptyStringValidator returns a validator that ensures string is not empty or whitespace-only
func NonEmptyStringValidator() validator.String {
	return &nonEmptyStringValidator{}
}

type nonEmptyStringValidator struct{}

func (v nonEmptyStringValidator) Description(ctx context.Context) string {
	return "must not be empty or whitespace-only"
}

func (v nonEmptyStringValidator) MarkdownDescription(ctx context.Context) string {
	return "must not be empty or whitespace-only"
}

func (v nonEmptyStringValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	value := req.ConfigValue.ValueString()

	if strings.TrimSpace(value) == "" {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Empty Value",
			"Value must not be empty or contain only whitespace.",
		)
	}
}

// Common validator combinations for convenience

// RequiredNameValidators returns validators for required name fields
func RequiredNameValidators() []validator.String {
	return []validator.String{
		stringvalidator.LengthBetween(1, 64),
		NameValidator(),
	}
}

// RequiredNamespaceValidators returns validators for required namespace fields
func RequiredNamespaceValidators() []validator.String {
	return []validator.String{
		stringvalidator.LengthBetween(1, 64),
		NamespaceValidator(),
	}
}

// OptionalDomainValidators returns validators for optional domain fields
func OptionalDomainValidators() []validator.String {
	return []validator.String{
		DomainValidator(),
	}
}
