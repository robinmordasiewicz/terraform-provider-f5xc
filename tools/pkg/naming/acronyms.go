// Package naming provides consistent case conversion and acronym handling
// for code generation tools in the F5XC Terraform provider.
package naming

// UppercaseAcronyms defines acronyms that should always be uppercase.
// Based on RFC 4949, IEEE standards, and industry style guides (Google, Microsoft, Apple).
var UppercaseAcronyms = map[string]bool{
	// Networking protocols
	"DNS": true, "HTTP": true, "HTTPS": true, "TCP": true, "UDP": true,
	"TLS": true, "SSL": true, "SSH": true, "FTP": true, "SFTP": true,
	"SMTP": true, "IMAP": true, "POP": true, "LDAP": true, "DHCP": true,
	"ARP": true, "ICMP": true, "SNMP": true, "NTP": true, "SIP": true,
	"RTP": true, "RTSP": true, "QUIC": true, "IP": true, "GRPC": true,
	// Web/API
	"API": true, "URL": true, "URI": true, "REST": true, "SOAP": true,
	"JSON": true, "XML": true, "HTML": true, "CSS": true, "CORS": true,
	"CDN": true, "WAF": true, "JWT": true, "SAML": true,
	// Network infrastructure
	"VPN": true, "NAT": true, "VLAN": true, "BGP": true, "OSPF": true,
	"QOS": true, "MTU": true, "TTL": true, "ACL": true, "CIDR": true,
	"VIP": true, "LB": true, "HA": true, "DR": true,
	// Security
	"PKI": true, "CA": true, "CSR": true, "CRL": true, "OCSP": true,
	"PEM": true, "AES": true, "RSA": true, "SHA": true, "MD5": true,
	"HMAC": true, "MFA": true, "SSO": true, "RBAC": true, "IAM": true,
	"DDOS": true, "DOS": true, "XSS": true, "CSRF": true, "SQL": true,
	// Cloud/Infrastructure
	"AWS": true, "GCP": true, "CPU": true, "RAM": true, "SSD": true,
	"HDD": true, "GPU": true, "RAID": true, "VM": true, "OS": true,
	"SLA": true, "RPO": true, "RTO": true, "VPC": true, "VNET": true,
	"TGW": true, "IKE": true, "ID": true, "SLI": true, "S2S": true,
	// F5/Volterra specific
	// Note: ASN is intentionally NOT included to maintain backward compatibility
	// with existing code that uses "Asn" in type names (e.g., BGPAsnSet)
	"RE": true, "CE": true, "SPO": true, "SMG": true,
	"APM": true, "PII": true, "OIDC": true, "K8S": true,
}

// MixedCaseAcronyms defines acronyms with specific mixed-case conventions.
// These should be preserved exactly as specified.
var MixedCaseAcronyms = map[string]string{
	"mtls":      "mTLS",
	"oauth":     "OAuth",
	"graphql":   "GraphQL",
	"websocket": "WebSocket",
	"iscsi":     "iSCSI",
	"ipv4":      "IPv4",
	"ipv6":      "IPv6",
	"macos":     "macOS",
	"ios":       "iOS",
	"nosql":     "NoSQL",
}

// CompoundWords defines compound words that should have specific PascalCase formatting
// when converting resource names to Go type names.
// Example: "loadbalancer" -> "LoadBalancer" (not "Loadbalancer")
var CompoundWords = map[string]string{
	"loadbalancer": "LoadBalancer",
	"bigip":        "BigIP",
	"websocket":    "WebSocket",
	"fastcgi":      "FastCGI",
}

// IsUppercaseAcronym returns true if the given string (in any case) is a known uppercase acronym.
func IsUppercaseAcronym(s string) bool {
	return UppercaseAcronyms[ToUpper(s)]
}

// GetMixedCaseAcronym returns the correct mixed-case form if the string is a known
// mixed-case acronym, otherwise returns empty string.
func GetMixedCaseAcronym(s string) string {
	return MixedCaseAcronyms[ToLower(s)]
}

// ToUpper is a helper that converts string to uppercase (avoids import in callers).
func ToUpper(s string) string {
	result := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= 'a' && c <= 'z' {
			result[i] = c - 32
		} else {
			result[i] = c
		}
	}
	return string(result)
}

// ToLower is a helper that converts string to lowercase (avoids import in callers).
func ToLower(s string) string {
	result := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= 'A' && c <= 'Z' {
			result[i] = c + 32
		} else {
			result[i] = c
		}
	}
	return string(result)
}
