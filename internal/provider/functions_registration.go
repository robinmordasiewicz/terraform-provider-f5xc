// This file is MANUALLY MAINTAINED and is NOT auto-generated from OpenAPI specifications.
// It registers provider-defined functions for utility operations that are not part of
// the F5XC API specification.
//
// DO NOT DELETE OR MODIFY during code generation. This file is preserved by the
// generate-all-schemas.go tool.

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"

	"github.com/f5xc/terraform-provider-f5xc/internal/functions"
)

// Ensure F5XCProvider satisfies the provider.ProviderWithFunctions interface.
var _ provider.ProviderWithFunctions = &F5XCProvider{}

// Functions returns utility functions provided by this provider.
// These include blindfold encryption functions for F5XC Secret Management.
//
// Available functions:
//   - blindfold: Encrypts base64-encoded plaintext using F5XC Secret Management
//   - blindfold_file: Reads a file and encrypts its contents using F5XC Secret Management
//
// These functions require Terraform 1.8.0 or later and are called using the provider namespace:
//
//	provider::f5xc::blindfold(plaintext, policy_name, namespace)
//	provider::f5xc::blindfold_file(path, policy_name, namespace)
func (p *F5XCProvider) Functions(ctx context.Context) []func() function.Function {
	return []func() function.Function{
		functions.NewBlindfoldFunction,
		functions.NewBlindfoldFileFunction,
	}
}
