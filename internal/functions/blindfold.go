// This file is MANUALLY MAINTAINED and is NOT auto-generated.

package functions

import (
	"context"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/function"

	"github.com/f5xc/terraform-provider-f5xc/internal/blindfold"
)

var _ function.Function = &BlindfoldFunction{}

// BlindfoldFunction implements the blindfold() provider function.
// It encrypts base64-encoded plaintext using F5XC Secret Management.
type BlindfoldFunction struct{}

// NewBlindfoldFunction creates a new blindfold function instance.
func NewBlindfoldFunction() function.Function {
	return &BlindfoldFunction{}
}

// Metadata returns the function name.
func (f *BlindfoldFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "blindfold"
}

// Definition describes the function signature and documentation.
func (f *BlindfoldFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Encrypt a secret using F5XC blindfold",
		Description: "Encrypts base64-encoded plaintext using F5 Distributed Cloud Secret Management (blindfold). " +
			"Returns a sealed secret string suitable for use in blindfold_secret_info.location fields. " +
			"The encryption happens locally using the public key fetched from F5XC - the plaintext is never sent to F5XC.",
		MarkdownDescription: `Encrypts base64-encoded plaintext using F5 Distributed Cloud Secret Management (blindfold).

Returns a sealed secret string suitable for use in ` + "`blindfold_secret_info.location`" + ` fields.

**Security**: The encryption happens locally using the public key fetched from F5XC.
The plaintext secret is **never** transmitted to F5XC during encryption.

### Example

` + "```hcl" + `
resource "f5xc_http_loadbalancer" "example" {
  name = "secure-lb"

  tls_parameters {
    private_key {
      blindfold_secret_info {
        location = provider::f5xc::blindfold(
          base64encode(file("${path.module}/private.key")),
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
				Name:        "plaintext",
				Description: "Base64-encoded plaintext to encrypt. Use Terraform's base64encode() function for raw strings or file contents.",
				MarkdownDescription: "Base64-encoded plaintext to encrypt. Use Terraform's `base64encode()` function for raw strings or file contents.\n\n" +
					"Example: `base64encode(file(\"private.key\"))`",
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

// Run executes the blindfold encryption.
func (f *BlindfoldFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var plaintextBase64, policyName, namespace string

	// Extract arguments
	resp.Error = function.ConcatFuncErrors(
		req.Arguments.Get(ctx, &plaintextBase64, &policyName, &namespace),
	)
	if resp.Error != nil {
		return
	}

	// Validate arguments
	if plaintextBase64 == "" {
		resp.Error = function.NewArgumentFuncError(0, "plaintext cannot be empty")
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

	// Decode base64 plaintext to validate it
	plaintext, err := base64.StdEncoding.DecodeString(plaintextBase64)
	if err != nil {
		resp.Error = function.NewArgumentFuncError(0,
			fmt.Sprintf("Invalid base64 encoding in plaintext: %s. Use base64encode() to encode your secret.", err),
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
			fmt.Sprintf("Plaintext too large (%d bytes). Maximum size for this key is %d bytes. Consider using a symmetric key or splitting the data.", len(plaintext), maxSize),
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

// bearerTokenTransport is an http.RoundTripper that adds Bearer token authentication.
type bearerTokenTransport struct {
	token     string
	transport http.RoundTripper
}

func (t *bearerTokenTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Clone the request to avoid modifying the original
	reqClone := req.Clone(req.Context())
	reqClone.Header.Set("Authorization", "APIToken "+t.token)
	reqClone.Header.Set("Content-Type", "application/json")

	transport := t.transport
	if transport == nil {
		transport = &http.Transport{
			TLSClientConfig: &tls.Config{MinVersion: tls.VersionTLS12},
		}
	}

	return transport.RoundTrip(reqClone)
}
