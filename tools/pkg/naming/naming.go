// Package naming provides consistent case conversion and acronym handling
// for code generation tools in the F5XC Terraform provider.
package naming

import (
	"regexp"
	"strings"
	"unicode"
)

// ToTitleCase converts a snake_case or dot.separated string to Title Case,
// preserving acronym capitalization.
// Example: "http_load_balancer" -> "HTTP Load Balancer"
func ToTitleCase(s string) string {
	// Replace underscores and dots with spaces
	s = strings.ReplaceAll(s, "_", " ")
	s = strings.ReplaceAll(s, ".", " ")

	words := strings.Fields(s)
	for i, word := range words {
		if len(word) > 0 {
			// Apply standard title case first
			words[i] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
		}
	}
	result := strings.Join(words, " ")

	// Apply acronym normalization
	result = NormalizeAcronyms(result)

	return result
}

// ToTitleCaseFromAnchor converts an anchor name (kebab-case) to Title Case,
// preserving acronym capitalization.
// Example: "http-load-balancer" -> "HTTP Load Balancer"
func ToTitleCaseFromAnchor(anchor string) string {
	words := strings.Split(anchor, "-")
	for i, word := range words {
		upper := strings.ToUpper(word)
		if UppercaseAcronyms[upper] {
			words[i] = upper
		} else if mixed, ok := MixedCaseAcronyms[strings.ToLower(word)]; ok {
			words[i] = mixed
		} else if len(word) > 0 {
			words[i] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
		}
	}
	return strings.Join(words, " ")
}

// ToSnakeCase converts a CamelCase or PascalCase string to snake_case.
// Example: "HTTPLoadBalancer" -> "http_load_balancer"
func ToSnakeCase(s string) string {
	var result strings.Builder
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				// Don't add underscore if previous char was also uppercase (acronym)
				prev := rune(s[i-1])
				if !unicode.IsUpper(prev) {
					result.WriteRune('_')
				} else if i+1 < len(s) && unicode.IsLower(rune(s[i+1])) {
					// End of acronym followed by lowercase
					result.WriteRune('_')
				}
			}
			result.WriteRune(unicode.ToLower(r))
		} else {
			result.WriteRune(r)
		}
	}
	return result.String()
}

// ToHumanName converts a resource name to human-readable format.
// Example: "http_loadbalancer" -> "HTTP Loadbalancer"
func ToHumanName(resourceName string) string {
	// Split by underscore
	parts := strings.Split(resourceName, "_")
	result := make([]string, len(parts))

	for i, part := range parts {
		upper := strings.ToUpper(part)
		if UppercaseAcronyms[upper] {
			result[i] = upper
		} else if mixed, ok := MixedCaseAcronyms[strings.ToLower(part)]; ok {
			result[i] = mixed
		} else {
			// Capitalize first letter
			result[i] = strings.Title(part) //nolint:staticcheck // strings.Title is fine for single words
		}
	}

	return strings.Join(result, " ")
}

// ToHumanReadableName converts a resource name to human-readable format with proper spacing.
// Unlike ToHumanName, this function also handles compound words with spaces.
// Example: "http_loadbalancer" -> "HTTP Load Balancer"
func ToHumanReadableName(resourceName string) string {
	parts := strings.Split(resourceName, "_")
	var result []string

	for _, part := range parts {
		lower := strings.ToLower(part)
		upper := strings.ToUpper(part)

		if UppercaseAcronyms[upper] {
			result = append(result, upper)
		} else if compound, ok := CompoundWordsHumanReadable[lower]; ok {
			// Handle compound words with spaces (e.g., "loadbalancer" -> "Load Balancer")
			result = append(result, compound)
		} else if mixed, ok := MixedCaseAcronyms[lower]; ok {
			result = append(result, mixed)
		} else {
			// Capitalize first letter
			result = append(result, strings.Title(lower)) //nolint:staticcheck // strings.Title is fine for single words
		}
	}

	return strings.Join(result, " ")
}

// ToAnchorName converts a name to an anchor-friendly format (kebab-case).
// Example: "http_load_balancer" -> "http-load-balancer"
func ToAnchorName(name string) string {
	return strings.ToLower(strings.ReplaceAll(name, "_", "-"))
}

// NormalizeAcronyms corrects acronym capitalization in text.
// This function is idempotent - running it multiple times produces the same result.
func NormalizeAcronyms(text string) string {
	wordRegex := regexp.MustCompile(`\b([A-Za-z0-9]+)\b`)

	return wordRegex.ReplaceAllStringFunc(text, func(word string) string {
		upperWord := strings.ToUpper(word)
		lowerWord := strings.ToLower(word)

		// Check for mixed-case acronyms first (e.g., mTLS, OAuth)
		if corrected, ok := MixedCaseAcronyms[lowerWord]; ok {
			return corrected
		}

		// Check for uppercase acronyms (e.g., DNS, HTTP, TCP)
		if UppercaseAcronyms[upperWord] {
			return upperWord
		}

		// Return original word unchanged
		return word
	})
}

// StartsWithVowel checks if a string starts with a vowel (for "A" vs "An" grammar).
func StartsWithVowel(s string) bool {
	if len(s) == 0 {
		return false
	}
	firstChar := strings.ToLower(string(s[0]))
	return firstChar == "a" || firstChar == "e" || firstChar == "i" || firstChar == "o" || firstChar == "u"
}

// ToResourceTypeName converts a snake_case resource name to a Go type name.
// Example: "http_loadbalancer" -> "HTTPLoadBalancer"
func ToResourceTypeName(resourceName string) string {
	parts := strings.Split(resourceName, "_")
	var result strings.Builder

	for _, part := range parts {
		lower := strings.ToLower(part)
		upper := strings.ToUpper(part)

		if UppercaseAcronyms[upper] {
			result.WriteString(upper)
		} else if compound, ok := CompoundWords[lower]; ok {
			result.WriteString(compound)
		} else {
			result.WriteString(strings.Title(lower)) //nolint:staticcheck
		}
	}

	return result.String()
}
