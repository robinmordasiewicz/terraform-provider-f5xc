//go:build ignore
// +build ignore

// enrich-descriptions.go - LLM-based description enrichment for Terraform provider documentation
//
// This tool uses a local ollama instance to improve OpenAPI description quality.
// It caches results based on schema content hash to avoid unnecessary regeneration.
//
// Usage:
//
//	go run tools/enrich-descriptions.go --spec-dir=docs/specifications/api
package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

// CLI flags
var (
	specDir    = flag.String("spec-dir", "docs/specifications/api", "Directory containing OpenAPI spec files")
	cacheFile  = flag.String("cache-file", "tools/enriched-descriptions-cache.json", "Path to cache file")
	ollamaURL  = flag.String("ollama-url", "http://localhost:11434", "Ollama API URL")
	model      = flag.String("model", "qwen2.5-coder:7b", "Ollama model to use")
	force      = flag.Bool("force", false, "Force regeneration even if cache is valid")
	verbose    = flag.Bool("verbose", false, "Enable verbose output")
	dryRun     = flag.Bool("dry-run", false, "Show what would be enriched without calling LLM")
	maxWorkers = flag.Int("workers", 4, "Number of concurrent workers")
)

// DescriptionCache holds cached enriched descriptions
type DescriptionCache struct {
	SchemaHash   string            `json:"schema_hash"`
	Model        string            `json:"model"`
	GeneratedAt  string            `json:"generated_at"`
	Stats        CacheStats        `json:"stats"`
	Descriptions map[string]string `json:"descriptions"`
}

// CacheStats tracks enrichment statistics
type CacheStats struct {
	TotalEntries   int `json:"total_entries"`
	EnrichedCount  int `json:"enriched_count"`
	PassthroughCnt int `json:"passthrough_count"`
}

// SchemaDefinition matches the structure in generate-all-schemas.go
type SchemaDefinition struct {
	Type                string                      `json:"type"`
	Description         string                      `json:"description"`
	Title               string                      `json:"title"`
	Format              string                      `json:"format"`
	Enum                []interface{}               `json:"enum"`
	Properties          map[string]SchemaDefinition `json:"properties"`
	Items               *SchemaDefinition           `json:"items"`
	Ref                 string                      `json:"$ref"`
	XDisplayName        string                      `json:"x-displayname"`
	XVesExample         string                      `json:"x-ves-example"`
	XVesValidationRules map[string]string           `json:"x-ves-validation-rules"`
}

// OpenAPISpec represents the top-level OpenAPI document
type OpenAPISpec struct {
	Info       SpecInfo       `json:"info"`
	Components SpecComponents `json:"components"`
	// Paths is not unmarshaled - we only need schemas, and some specs have non-standard path values
}

// SpecInfo contains API metadata
type SpecInfo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// SpecComponents contains schema definitions
type SpecComponents struct {
	Schemas map[string]SchemaDefinition `json:"schemas"`
}

// DescriptionContext holds all context for a description
type DescriptionContext struct {
	ResourceName string
	FieldPath    string
	Description  string
	Title        string
	DisplayName  string
	Type         string
	Example      string
	Enum         []string
	Validation   map[string]string
}

// OllamaRequest is the request body for ollama API
type OllamaRequest struct {
	Model   string        `json:"model"`
	Prompt  string        `json:"prompt"`
	Stream  bool          `json:"stream"`
	Options OllamaOptions `json:"options"`
}

// OllamaOptions contains generation parameters
type OllamaOptions struct {
	Temperature float64 `json:"temperature"`
	Seed        int     `json:"seed"`
	NumCtx      int     `json:"num_ctx"`
}

// OllamaResponse is the response from ollama API
type OllamaResponse struct {
	Response string `json:"response"`
}

func main() {
	flag.Parse()

	log.SetFlags(log.Ltime)
	log.Println("Starting description enrichment...")

	// Load existing cache
	cache := loadCache(*cacheFile)

	// Extract all descriptions from specs
	descriptions := extractDescriptions(*specDir)
	if len(descriptions) == 0 {
		log.Fatal("No descriptions found in spec files")
	}
	log.Printf("Extracted %d descriptions from specs", len(descriptions))

	// Compute hash of current descriptions
	currentHash := computeHash(descriptions)
	log.Printf("Current schema hash: %s", currentHash[:16])

	// Check if cache is valid
	if !*force && cache.SchemaHash == currentHash {
		log.Println("Cache is up-to-date, no enrichment needed")
		return
	}

	if cache.SchemaHash != "" {
		log.Printf("Cache hash mismatch (cached: %s), regenerating...", cache.SchemaHash[:16])
	} else {
		log.Println("No existing cache, generating from scratch...")
	}

	if *dryRun {
		log.Println("Dry run mode - showing descriptions that would be enriched:")
		count := 0
		for key, ctx := range descriptions {
			if needsEnrichment(ctx.Description) {
				count++
				if count <= 20 {
					log.Printf("  %s: %q", key, truncate(ctx.Description, 60))
				}
			}
		}
		log.Printf("Total descriptions needing enrichment: %d", count)
		return
	}

	// Check ollama availability
	if !checkOllama(*ollamaURL) {
		fmt.Println()
		fmt.Println("=" + strings.Repeat("=", 70))
		fmt.Println("ERROR: Ollama is not available")
		fmt.Println(strings.Repeat("=", 71))
		fmt.Println()
		fmt.Println("The description enrichment requires a local Ollama instance.")
		fmt.Println()
		fmt.Println("INSTALLATION:")
		fmt.Println("  macOS:   brew install ollama")
		fmt.Println("  Linux:   curl -fsSL https://ollama.com/install.sh | sh")
		fmt.Println("  Windows: Download from https://ollama.com/download")
		fmt.Println()
		fmt.Println("SETUP:")
		fmt.Println("  1. Start the Ollama server:")
		fmt.Println("     $ ollama serve")
		fmt.Println()
		fmt.Println("  2. Pull the required model (one-time):")
		fmt.Printf("     $ ollama pull %s\n", *model)
		fmt.Println()
		fmt.Println("  3. Re-run this command:")
		fmt.Println("     $ make enrich-descriptions")
		fmt.Println()
		fmt.Printf("Ollama URL: %s (use --ollama-url to change)\n", *ollamaURL)
		fmt.Println(strings.Repeat("=", 71))
		os.Exit(1)
	}

	// Ensure model is available
	if err := pullModel(*ollamaURL, *model); err != nil {
		log.Printf("Warning: Could not pull model: %v", err)
	}

	// Enrich descriptions
	enriched := enrichDescriptions(descriptions)

	// Update cache
	cache.SchemaHash = currentHash
	cache.Model = *model
	cache.GeneratedAt = time.Now().UTC().Format(time.RFC3339)
	cache.Descriptions = enriched
	cache.Stats = computeStats(descriptions, enriched)

	// Save cache
	if err := saveCache(*cacheFile, cache); err != nil {
		log.Fatalf("Failed to save cache: %v", err)
	}

	log.Printf("Enrichment complete: %d total, %d enriched, %d passthrough",
		cache.Stats.TotalEntries, cache.Stats.EnrichedCount, cache.Stats.PassthroughCnt)
}

// loadCache loads the cache from disk
func loadCache(path string) *DescriptionCache {
	cache := &DescriptionCache{
		Descriptions: make(map[string]string),
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return cache
		}
		log.Printf("Warning: Could not read cache: %v", err)
		return cache
	}

	if err := json.Unmarshal(data, cache); err != nil {
		log.Printf("Warning: Could not parse cache: %v", err)
		return &DescriptionCache{Descriptions: make(map[string]string)}
	}

	return cache
}

// saveCache saves the cache to disk
func saveCache(path string, cache *DescriptionCache) error {
	data, err := json.MarshalIndent(cache, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal cache: %w", err)
	}
	return os.WriteFile(path, data, 0644)
}

// extractDescriptions reads all OpenAPI specs and extracts descriptions
func extractDescriptions(dir string) map[string]DescriptionContext {
	descriptions := make(map[string]DescriptionContext)

	pattern := filepath.Join(dir, "docs-cloud-f5-com.*.ves-swagger.json")
	files, err := filepath.Glob(pattern)
	if err != nil {
		log.Fatalf("Failed to glob spec files: %v", err)
	}

	if len(files) == 0 {
		log.Printf("Warning: No spec files found matching %s", pattern)
		return descriptions
	}

	log.Printf("Processing %d spec files...", len(files))

	for _, file := range files {
		resourceName := extractResourceName(file)
		spec := parseSpec(file)
		if spec == nil {
			continue
		}

		// Extract resource-level description
		if spec.Info.Description != "" {
			key := resourceName
			descriptions[key] = DescriptionContext{
				ResourceName: resourceName,
				FieldPath:    "",
				Description:  spec.Info.Description,
				Title:        spec.Info.Title,
			}
		}

		// Extract schema descriptions
		for schemaName, schema := range spec.Components.Schemas {
			extractSchemaDescriptions(resourceName, schemaName, "", schema, descriptions)
		}
	}

	return descriptions
}

// extractSchemaDescriptions recursively extracts descriptions from a schema
func extractSchemaDescriptions(resourceName, schemaName, path string, schema SchemaDefinition, descriptions map[string]DescriptionContext) {
	// Build the key
	var key string
	if path == "" {
		key = fmt.Sprintf("%s:%s", resourceName, schemaName)
	} else {
		key = fmt.Sprintf("%s:%s.%s", resourceName, schemaName, path)
	}

	// Add this description if present
	if schema.Description != "" {
		descriptions[key] = DescriptionContext{
			ResourceName: resourceName,
			FieldPath:    path,
			Description:  schema.Description,
			Title:        schema.Title,
			DisplayName:  schema.XDisplayName,
			Type:         schema.Type,
			Example:      schema.XVesExample,
			Enum:         toStringSlice(schema.Enum),
			Validation:   schema.XVesValidationRules,
		}
	}

	// Recurse into properties
	for propName, propSchema := range schema.Properties {
		newPath := propName
		if path != "" {
			newPath = path + "." + propName
		}
		extractSchemaDescriptions(resourceName, schemaName, newPath, propSchema, descriptions)
	}

	// Recurse into items
	if schema.Items != nil {
		extractSchemaDescriptions(resourceName, schemaName, path+"[]", *schema.Items, descriptions)
	}
}

// extractResourceName extracts the resource name from a spec file path
func extractResourceName(path string) string {
	// Pattern: docs-cloud-f5-com.XXXX.public.ves.io.schema.RESOURCE_NAME.ves-swagger.json
	base := filepath.Base(path)
	parts := strings.Split(base, ".")

	// Find "schema" and take the next part
	for i, part := range parts {
		if part == "schema" && i+1 < len(parts) {
			return parts[i+1]
		}
	}

	return strings.TrimSuffix(base, ".ves-swagger.json")
}

// parseSpec parses an OpenAPI spec file
func parseSpec(path string) *OpenAPISpec {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Printf("Warning: Could not read %s: %v", path, err)
		return nil
	}

	var spec OpenAPISpec
	if err := json.Unmarshal(data, &spec); err != nil {
		log.Printf("Warning: Could not parse %s: %v", path, err)
		return nil
	}

	return &spec
}

// computeHash computes a SHA256 hash of all descriptions
func computeHash(descriptions map[string]DescriptionContext) string {
	// Sort keys for deterministic hashing
	keys := make([]string, 0, len(descriptions))
	for k := range descriptions {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	h := sha256.New()
	for _, k := range keys {
		ctx := descriptions[k]
		h.Write([]byte(k))
		h.Write([]byte(ctx.Description))
		h.Write([]byte(ctx.Title))
		h.Write([]byte(ctx.DisplayName))
		h.Write([]byte(ctx.Type))
		h.Write([]byte(ctx.Example))
	}

	return fmt.Sprintf("sha256:%x", h.Sum(nil))
}

// needsEnrichment determines if a description should be enriched
func needsEnrichment(desc string) bool {
	desc = strings.TrimSpace(desc)

	// Empty - needs enrichment
	if len(desc) == 0 {
		return true
	}

	// Very short single word - needs enrichment
	if len(desc) < 20 && !strings.Contains(desc, " ") {
		return true
	}

	// Already good quality (long, well-formed) - skip
	if len(desc) > 100 && strings.Contains(desc, ".") {
		return false
	}

	// Terse patterns that need enrichment
	tersePatterns := []string{
		"x-displayname",
		"x-ves-",
		"Enable this option",
		"Configuration for",
		"Shape of",
	}
	descLower := strings.ToLower(desc)
	for _, pattern := range tersePatterns {
		if strings.Contains(descLower, strings.ToLower(pattern)) {
			return true
		}
	}

	// Single word title-case - needs enrichment
	if len(strings.Fields(desc)) <= 2 {
		return true
	}

	return false
}

// enrichDescriptions enriches all descriptions using ollama
func enrichDescriptions(descriptions map[string]DescriptionContext) map[string]string {
	result := make(map[string]string)

	total := len(descriptions)
	enriched := 0
	passthrough := 0

	for key, ctx := range descriptions {
		if !needsEnrichment(ctx.Description) {
			// Pass through already-good descriptions
			result[key] = ctx.Description
			passthrough++
			continue
		}

		// Enrich with LLM
		improved := callOllama(ctx)
		if improved != "" && validateEnrichment(ctx.Description, improved) {
			result[key] = improved
			enriched++
		} else {
			result[key] = ctx.Description
			passthrough++
		}

		// Progress
		if *verbose || (enriched+passthrough)%100 == 0 {
			log.Printf("Progress: %d/%d (enriched: %d, passthrough: %d)",
				enriched+passthrough, total, enriched, passthrough)
		}
	}

	return result
}

// callOllama calls the ollama API to enrich a description
func callOllama(ctx DescriptionContext) string {
	prompt := buildPrompt(ctx)

	reqBody := OllamaRequest{
		Model:  *model,
		Prompt: prompt,
		Stream: false,
		Options: OllamaOptions{
			Temperature: 0,
			Seed:        42,
			NumCtx:      4096,
		},
	}

	data, err := json.Marshal(reqBody)
	if err != nil {
		log.Printf("Error marshaling request: %v", err)
		return ""
	}

	resp, err := http.Post(*ollamaURL+"/api/generate", "application/json", bytes.NewReader(data))
	if err != nil {
		log.Printf("Error calling ollama: %v", err)
		return ""
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		log.Printf("Ollama error (status %d): %s", resp.StatusCode, string(body))
		return ""
	}

	var ollamaResp OllamaResponse
	if err := json.NewDecoder(resp.Body).Decode(&ollamaResp); err != nil {
		log.Printf("Error decoding ollama response: %v", err)
		return ""
	}

	return strings.TrimSpace(ollamaResp.Response)
}

// buildPrompt constructs the LLM prompt
func buildPrompt(ctx DescriptionContext) string {
	var sb strings.Builder

	sb.WriteString("Improve this Terraform field description. Return ONLY the improved text (1-2 sentences, under 100 words).\n\n")

	sb.WriteString(fmt.Sprintf("Resource: %s\n", ctx.ResourceName))
	if ctx.FieldPath != "" {
		sb.WriteString(fmt.Sprintf("Field: %s\n", ctx.FieldPath))
	}
	sb.WriteString(fmt.Sprintf("Original: %s\n", ctx.Description))

	if ctx.Type != "" {
		sb.WriteString(fmt.Sprintf("Type: %s\n", ctx.Type))
	}
	if ctx.DisplayName != "" {
		sb.WriteString(fmt.Sprintf("Display Name: %s\n", ctx.DisplayName))
	}
	if ctx.Example != "" {
		sb.WriteString(fmt.Sprintf("Example: %s\n", ctx.Example))
	}
	if len(ctx.Enum) > 0 {
		sb.WriteString(fmt.Sprintf("Allowed values: %s\n", strings.Join(ctx.Enum, ", ")))
	}

	sb.WriteString("\nImproved description:")

	return sb.String()
}

// validateEnrichment checks if the enriched description is valid
func validateEnrichment(original, enriched string) bool {
	// Reject if shorter than original
	if len(enriched) < len(original) {
		return false
	}

	// Reject if too short
	if len(enriched) < 10 {
		return false
	}

	// Reject hallucination patterns
	hallucinations := []string{
		"I ", "we ", "you should", "please note",
		"as mentioned", "in this case", "here is",
		"the improved", "improved description",
	}
	enrichedLower := strings.ToLower(enriched)
	for _, h := range hallucinations {
		if strings.Contains(enrichedLower, h) {
			return false
		}
	}

	return true
}

// checkOllama checks if ollama is available
func checkOllama(url string) bool {
	resp, err := http.Get(url + "/api/tags")
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusOK
}

// pullModel ensures the model is available
func pullModel(url, modelName string) error {
	reqBody := map[string]string{"name": modelName}
	data, _ := json.Marshal(reqBody)

	resp, err := http.Post(url+"/api/pull", "application/json", bytes.NewReader(data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read and discard response (streaming)
	io.Copy(io.Discard, resp.Body)
	return nil
}

// computeStats computes cache statistics
func computeStats(original map[string]DescriptionContext, enriched map[string]string) CacheStats {
	stats := CacheStats{
		TotalEntries: len(enriched),
	}

	for key, enrichedDesc := range enriched {
		if ctx, ok := original[key]; ok {
			if enrichedDesc != ctx.Description {
				stats.EnrichedCount++
			} else {
				stats.PassthroughCnt++
			}
		}
	}

	return stats
}

// toStringSlice converts []interface{} to []string
func toStringSlice(vals []interface{}) []string {
	result := make([]string, 0, len(vals))
	for _, v := range vals {
		if s, ok := v.(string); ok {
			result = append(result, s)
		}
	}
	return result
}

// truncate truncates a string to max length
func truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max-3] + "..."
}
