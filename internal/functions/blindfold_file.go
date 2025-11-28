// This file is MANUALLY MAINTAINED and is NOT auto-generated.

package functions

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/function"

	"github.com/f5xc/terraform-provider-f5xc/internal/blindfold"
)

var _ function.Function = &BlindfoldFileFunction{}

// BlindfoldFileFunction implements the blindfold_file() provider function.
// It reads a file and encrypts its contents using F5XC Secret Management.
type BlindfoldFileFunction struct{}

// NewBlindfoldFileFunction creates a new blindfold_file function instance.
func NewBlindfoldFileFunction() function.Function {
	return &BlindfoldFileFunction{}
}

// Metadata returns the function name.
func (f *BlindfoldFileFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "blindfold_file"
}

// Definition describes the function signature and documentation.
func (f *BlindfoldFileFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Encrypt a file using F5XC blindfold",
		Description: "Reads a file and encrypts its contents using F5 Distributed Cloud Secret Management (blindfold). " +
			"Returns a sealed secret string suitable for use in blindfold_secret_info.location fields. " +
			"This is a convenience function equivalent to blindfold(base64encode(file(path)), policy_name, namespace).",
		MarkdownDescription: `Reads a file and encrypts its contents using F5 Distributed Cloud Secret Management (blindfold).

Returns a sealed secret string suitable for use in ` + "`blindfold_secret_info.location`" + ` fields.

This is a convenience function equivalent to:
` + "```hcl" + `
provider::f5xc::blindfold(base64encode(file(path)), policy_name, namespace)
` + "```" + `

**Security**: The encryption happens locally using the public key fetched from F5XC.
The file contents are **never** transmitted to F5XC during encryption.

### Example

` + "```hcl" + `
resource "f5xc_http_loadbalancer" "example" {
  name = "secure-lb"

  tls_parameters {
    private_key {
      blindfold_secret_info {
        location = provider::f5xc::blindfold_file(
          "${path.module}/certs/private.key",
          "example-secret-policy",
          "shared"
        )
      }
    }
  }
}
` + "```",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "path",
				Description: "Path to the file to encrypt. Can be absolute or relative to the Terraform working directory.",
				MarkdownDescription: "Path to the file to encrypt. Can be absolute or relative to the Terraform working directory.\n\n" +
					"Use `${path.module}` for paths relative to the current module.",
			},
			function.StringParameter{
				Name:        "policy_name",
				Description: "Name of the SecretPolicy that controls which clients can decrypt this secret.",
				MarkdownDescription: "Name of the SecretPolicy that controls which clients can decrypt this secret.\n\n" +
					"The policy must exist in the specified namespace before encryption.",
			},
			function.StringParameter{
				Name:        "namespace",
				Description: "F5XC namespace containing the SecretPolicy.",
				MarkdownDescription: "F5XC namespace containing the SecretPolicy.\n\n" +
					"Common values: `shared`, `system`, or your application namespace.",
			},
		},
		Return: function.StringReturn{},
	}
}

// Run executes the blindfold file encryption.
func (f *BlindfoldFileFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var path, policyName, namespace string

	// Extract arguments
	resp.Error = function.ConcatFuncErrors(
		req.Arguments.Get(ctx, &path, &policyName, &namespace),
	)
	if resp.Error != nil {
		return
	}

	// Validate arguments
	if path == "" {
		resp.Error = function.NewArgumentFuncError(0, "path cannot be empty")
		return
	}
	if policyName == "" {
		resp.Error = function.NewArgumentFuncError(1, "policy_name cannot be empty")
		return
	}
	if namespace == "" {
		resp.Error = function.NewArgumentFuncError(2, "namespace cannot be empty")
		return
	}

	// Get API configuration from environment variables
	apiURL := os.Getenv("F5XC_API_URL")
	if apiURL == "" {
		apiURL = "https://console.ves.volterra.io/api"
	}
	apiToken := os.Getenv("F5XC_API_TOKEN")
	if apiToken == "" {
		resp.Error = function.NewFuncError(
			"F5XC_API_TOKEN environment variable is required for blindfold encryption. " +
				"Set this variable to your F5 Distributed Cloud API token.",
		)
		return
	}

	// Create HTTP client with bearer token authentication
	httpClient := &http.Client{
		Timeout: 30 * time.Second,
		Transport: &bearerTokenTransport{
			token:     apiToken,
			transport: http.DefaultTransport,
		},
	}

	// Clean and validate the path
	cleanPath := filepath.Clean(path)

	// Check if file exists and is readable
	info, err := os.Stat(cleanPath)
	if err != nil {
		if os.IsNotExist(err) {
			resp.Error = function.NewArgumentFuncError(0,
				fmt.Sprintf("File not found: %s", cleanPath),
			)
		} else if os.IsPermission(err) {
			resp.Error = function.NewArgumentFuncError(0,
				fmt.Sprintf("Permission denied reading file: %s", cleanPath),
			)
		} else {
			resp.Error = function.NewArgumentFuncError(0,
				fmt.Sprintf("Error accessing file %s: %s", cleanPath, err),
			)
		}
		return
	}

	// Validate it's a regular file
	if !info.Mode().IsRegular() {
		resp.Error = function.NewArgumentFuncError(0,
			fmt.Sprintf("Path is not a regular file: %s", cleanPath),
		)
		return
	}

	// Read file contents
	plaintext, err := os.ReadFile(cleanPath)
	if err != nil {
		resp.Error = function.NewArgumentFuncError(0,
			fmt.Sprintf("Failed to read file %s: %s", cleanPath, err),
		)
		return
	}

	if len(plaintext) == 0 {
		resp.Error = function.NewArgumentFuncError(0,
			fmt.Sprintf("File is empty: %s", cleanPath),
		)
		return
	}

	// Fetch public key from F5XC
	pubKey, err := blindfold.GetPublicKey(ctx, httpClient, apiURL)
	if err != nil {
		resp.Error = function.NewFuncError(
			fmt.Sprintf("Failed to fetch public key from F5XC: %s", err),
		)
		return
	}

	// Fetch policy document from F5XC
	policy, err := blindfold.GetSecretPolicyDocument(ctx, httpClient, apiURL, namespace, policyName)
	if err != nil {
		resp.Error = function.NewFuncError(
			fmt.Sprintf("Failed to fetch secret policy %q in namespace %q: %s", policyName, namespace, err),
		)
		return
	}

	// Check plaintext size against RSA limits
	maxSize, err := blindfold.MaxPlaintextSize(pubKey)
	if err != nil {
		resp.Error = function.NewFuncError(
			fmt.Sprintf("Failed to determine maximum plaintext size: %s", err),
		)
		return
	}
	if len(plaintext) > maxSize {
		resp.Error = function.NewArgumentFuncError(0,
			fmt.Sprintf("File too large (%d bytes). Maximum size for this key is %d bytes. File: %s", len(plaintext), maxSize, cleanPath),
		)
		return
	}

	// Perform encryption
	sealed, err := blindfold.Seal(plaintext, pubKey, policy)
	if err != nil {
		resp.Error = function.NewFuncError(
			fmt.Sprintf("Encryption failed: %s", err),
		)
		return
	}

	// Set the result
	resp.Error = resp.Result.Set(ctx, sealed)
}
