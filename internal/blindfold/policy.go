// This file is MANUALLY MAINTAINED and is NOT auto-generated.

package blindfold

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	// PolicyDocumentEndpointFmt is the F5XC API endpoint template for fetching policy documents.
	// Parameters: namespace, policy_name
	PolicyDocumentEndpointFmt = "/api/secret_management/namespaces/%s/secret_policys/%s/get_policy_document"
)

// GetSecretPolicyDocument fetches a secret policy document from F5XC.
// The policy document defines which clients are authorized to decrypt secrets
// encrypted under this policy.
//
// Parameters:
//   - ctx: Context for cancellation and timeouts
//   - httpClient: Configured HTTP client with authentication
//   - baseURL: F5XC API base URL (e.g., "https://tenant.console.ves.volterra.io")
//   - namespace: F5XC namespace containing the policy
//   - name: Name of the secret policy
//
// Returns:
//   - SecretPolicyDocument containing policy ID and rules
//   - Error if the request fails or response is invalid
func GetSecretPolicyDocument(ctx context.Context, httpClient *http.Client, baseURL, namespace, name string) (*SecretPolicyDocument, error) {
	if httpClient == nil {
		return nil, fmt.Errorf("HTTP client is required")
	}
	if baseURL == "" {
		return nil, fmt.Errorf("base URL is required")
	}
	if namespace == "" {
		return nil, fmt.Errorf("namespace is required")
	}
	if name == "" {
		return nil, fmt.Errorf("policy name is required")
	}

	// Build URL with URL-encoded path components
	endpoint := fmt.Sprintf(PolicyDocumentEndpointFmt,
		url.PathEscape(namespace),
		url.PathEscape(name),
	)
	fullURL := baseURL + endpoint

	// Create request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Accept", "application/json")

	// Execute request
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	// Check status code
	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("secret policy %q not found in namespace %q", name, namespace)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status %d: %s", resp.StatusCode, string(body))
	}

	// Parse response - F5XC wraps responses in an envelope
	var envelope APIEnvelope[SecretPolicyDocument]
	if err := json.Unmarshal(body, &envelope); err != nil {
		// Try parsing without envelope (direct response)
		var policy SecretPolicyDocument
		if err2 := json.Unmarshal(body, &policy); err2 != nil {
			return nil, fmt.Errorf("failed to parse response: %w (envelope: %v)", err2, err)
		}
		// Set name/namespace from parameters if not in response
		if policy.Name == "" {
			policy.Name = name
		}
		if policy.Namespace == "" {
			policy.Namespace = namespace
		}
		return &policy, nil
	}

	// Validate the response
	if envelope.Data.PolicyID == "" {
		return nil, fmt.Errorf("response missing policy_id")
	}

	// Ensure name/namespace are set
	if envelope.Data.Name == "" {
		envelope.Data.Name = name
	}
	if envelope.Data.Namespace == "" {
		envelope.Data.Namespace = namespace
	}

	return &envelope.Data, nil
}
