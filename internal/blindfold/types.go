// This file is MANUALLY MAINTAINED and is NOT auto-generated.

package blindfold

// PublicKey represents the F5XC Secret Management public key used for encryption.
// The key is retrieved from the F5XC API and contains RSA key components.
type PublicKey struct {
	// KeyVersion is the version identifier for key rotation support
	KeyVersion int `json:"key_version" yaml:"keyVersion"`

	// ModulusBase64 is the RSA modulus (n) encoded in base64
	ModulusBase64 string `json:"modulus_base64" yaml:"modulusBase64"`

	// PublicExponentBase64 is the RSA public exponent (e) encoded in base64
	// Typically 65537 (0x10001)
	PublicExponentBase64 string `json:"public_exponent_base64" yaml:"publicExponentBase64"`

	// Tenant is the F5XC tenant identifier
	Tenant string `json:"tenant" yaml:"tenant"`
}

// SecretPolicyDocument represents an F5XC secret policy that controls
// which clients are authorized to decrypt secrets encrypted under this policy.
type SecretPolicyDocument struct {
	// Name is the policy name
	Name string `json:"name" yaml:"name"`

	// Namespace is the F5XC namespace containing the policy
	Namespace string `json:"namespace" yaml:"namespace"`

	// Tenant is the F5XC tenant identifier
	Tenant string `json:"tenant" yaml:"tenant"`

	// PolicyID is the unique identifier for this policy
	PolicyID string `json:"policy_id" yaml:"policyId"`

	// PolicyInfo contains the policy rules and algorithm
	PolicyInfo SecretPolicyInfo `json:"policy_info" yaml:"policyInfo"`
}

// SecretPolicyInfo contains the encryption algorithm and access rules.
type SecretPolicyInfo struct {
	// Algo is the encryption algorithm (e.g., "RSA-OAEP")
	Algo string `json:"algo" yaml:"algo"`

	// Rules defines the access control rules for decryption
	Rules []SecretPolicyRule `json:"rules" yaml:"rules"`
}

// SecretPolicyRule defines a single access control rule for secret decryption.
type SecretPolicyRule struct {
	// Action is the policy action (e.g., "ALLOW", "DENY")
	Action string `json:"action" yaml:"action"`

	// ClientName is the specific client name to match (optional)
	ClientName string `json:"client_name,omitempty" yaml:"clientName,omitempty"`

	// ClientNameMatcher provides pattern matching for client names (optional)
	ClientNameMatcher *MatcherType `json:"client_name_matcher,omitempty" yaml:"clientNameMatcher,omitempty"`

	// ClientSelector provides label-based selection for clients (optional)
	ClientSelector *LabelSelectorType `json:"client_selector,omitempty" yaml:"clientSelector,omitempty"`
}

// MatcherType provides flexible string matching capabilities.
type MatcherType struct {
	// ExactValues matches exact string values
	ExactValues []string `json:"exact_values,omitempty" yaml:"exactValues,omitempty"`

	// RegexValues matches using regular expressions
	RegexValues []string `json:"regex_values,omitempty" yaml:"regexValues,omitempty"`

	// Transformers applies transformations before matching
	Transformers []string `json:"transformers,omitempty" yaml:"transformers,omitempty"`
}

// LabelSelectorType provides Kubernetes-style label selection.
type LabelSelectorType struct {
	// Expressions are label selector expressions
	Expressions []string `json:"expressions,omitempty" yaml:"expressions,omitempty"`
}

// SealedSecret represents the encrypted output from the blindfold operation.
// This structure uses envelope encryption (hybrid encryption):
// - A random AES-256 key (DEK) encrypts the actual data
// - The DEK is encrypted with RSA-OAEP using the F5XC public key (KEK)
// This allows encrypting data of any size (up to 128KB API limit).
type SealedSecret struct {
	// KeyVersion identifies which public key version was used for encryption
	KeyVersion int `json:"key_version"`

	// PolicyID identifies the policy document used
	PolicyID string `json:"policy_id"`

	// Tenant is the F5XC tenant identifier
	Tenant string `json:"tenant"`

	// EncryptedKey is the base64-encoded RSA-OAEP encrypted AES-256 key (DEK)
	EncryptedKey string `json:"encrypted_key"`

	// Nonce is the base64-encoded 12-byte nonce used for AES-GCM
	Nonce string `json:"nonce"`

	// Ciphertext is the base64-encoded AES-GCM encrypted data (includes auth tag)
	Ciphertext string `json:"ciphertext"`
}

// SealedSecretLegacy represents the old direct RSA-OAEP format.
// Kept for backwards compatibility with secrets < 190 bytes.
// Deprecated: Use SealedSecret with envelope encryption instead.
type SealedSecretLegacy struct {
	// KeyVersion identifies which public key version was used for encryption
	KeyVersion int `json:"key_version"`

	// PolicyID identifies the policy document used
	PolicyID string `json:"policy_id"`

	// Tenant is the F5XC tenant identifier
	Tenant string `json:"tenant"`

	// Data is the base64-encoded RSA-OAEP encrypted ciphertext
	Data string `json:"data"`
}

// APIEnvelope wraps API responses from F5XC.
type APIEnvelope[T any] struct {
	Data T `json:"data"`
}
