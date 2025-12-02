// Copyright (c) F5XC Community
// SPDX-License-Identifier: MPL-2.0

// Package mocks provides a mock F5 XC API server for testing.
// This enables testing resources that would otherwise require external credentials
// (AWS, Azure, GCP, etc.) or special infrastructure.
package mocks

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"
)

// debugLog logs a message if MOCK_DEBUG is set
func debugLog(format string, v ...interface{}) {
	if os.Getenv("MOCK_DEBUG") != "" {
		fmt.Fprintf(os.Stderr, "[MOCK DEBUG] "+format+"\n", v...)
	}
}

// Server is a mock F5 XC API server for testing
type Server struct {
	*httptest.Server

	mu sync.RWMutex
	// resources stores created resources by path (e.g., "/api/config/namespaces/system/aws_vpc_sites/my-site")
	resources map[string]interface{}
	// handlers allows custom response handlers for specific paths
	handlers map[string]http.HandlerFunc
	// responseDelay simulates API latency
	responseDelay time.Duration
	// errorResponses allows injecting errors for specific paths
	errorResponses map[string]*ErrorResponse
	// requestLog records all requests for verification
	requestLog []RequestRecord
}

// RequestRecord stores information about a request for verification
type RequestRecord struct {
	Method    string
	Path      string
	Body      string
	Timestamp time.Time
}

// ErrorResponse defines a custom error response
type ErrorResponse struct {
	StatusCode int
	Body       interface{}
}

// F5XCResponse is the standard F5 XC API response wrapper
type F5XCResponse struct {
	Metadata    Metadata               `json:"metadata,omitempty"`
	Spec        map[string]interface{} `json:"spec,omitempty"`
	SystemMeta  *SystemMetadata        `json:"system_metadata,omitempty"`
	Status      map[string]interface{} `json:"status,omitempty"`
	ObjectIndex int                    `json:"object_index,omitempty"`
}

// Metadata represents F5 XC object metadata
type Metadata struct {
	Name        string            `json:"name"`
	Namespace   string            `json:"namespace,omitempty"`
	Labels      map[string]string `json:"labels,omitempty"`
	Annotations map[string]string `json:"annotations,omitempty"`
	Description string            `json:"description,omitempty"`
	Disable     bool              `json:"disable,omitempty"`
	UID         string            `json:"uid,omitempty"`
}

// SystemMetadata represents F5 XC system-generated metadata
type SystemMetadata struct {
	UID              string    `json:"uid,omitempty"`
	CreationTime     time.Time `json:"creation_timestamp,omitempty"`
	ModificationTime time.Time `json:"modification_timestamp,omitempty"`
	CreatorClass     string    `json:"creator_class,omitempty"`
	CreatorID        string    `json:"creator_id,omitempty"`
	Tenant           string    `json:"tenant,omitempty"`
}

// ListResponse is the standard F5 XC list API response
type ListResponse struct {
	Items []F5XCResponse `json:"items"`
}

// NewServer creates a new mock F5 XC API server
func NewServer() *Server {
	s := &Server{
		resources:      make(map[string]interface{}),
		handlers:       make(map[string]http.HandlerFunc),
		errorResponses: make(map[string]*ErrorResponse),
		requestLog:     make([]RequestRecord, 0),
	}

	s.Server = httptest.NewServer(http.HandlerFunc(s.handleRequest))
	return s
}

// NewTLSServer creates a new mock F5 XC API server with TLS
func NewTLSServer() *Server {
	s := &Server{
		resources:      make(map[string]interface{}),
		handlers:       make(map[string]http.HandlerFunc),
		errorResponses: make(map[string]*ErrorResponse),
		requestLog:     make([]RequestRecord, 0),
	}

	s.Server = httptest.NewTLSServer(http.HandlerFunc(s.handleRequest))
	return s
}

// URL returns the mock server's URL (implementing the base URL for the client)
func (s *Server) URL() string {
	return s.Server.URL
}

// SetResponseDelay sets a delay for all responses (simulates network latency)
func (s *Server) SetResponseDelay(d time.Duration) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.responseDelay = d
}

// SetHandler registers a custom handler for a specific path pattern
func (s *Server) SetHandler(pathPattern string, handler http.HandlerFunc) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers[pathPattern] = handler
}

// SetErrorResponse configures an error response for a specific path
func (s *Server) SetErrorResponse(path string, statusCode int, body interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.errorResponses[path] = &ErrorResponse{
		StatusCode: statusCode,
		Body:       body,
	}
}

// ClearErrorResponse removes an error response for a specific path
func (s *Server) ClearErrorResponse(path string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.errorResponses, path)
}

// SetResource pre-populates a resource in the mock server
func (s *Server) SetResource(path string, resource interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.resources[path] = resource
}

// GetResource retrieves a resource from the mock server
func (s *Server) GetResource(path string) (interface{}, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	r, ok := s.resources[path]
	return r, ok
}

// DeleteResource removes a resource from the mock server
func (s *Server) DeleteResource(path string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.resources, path)
}

// GetRequestLog returns all recorded requests
func (s *Server) GetRequestLog() []RequestRecord {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return append([]RequestRecord{}, s.requestLog...)
}

// ClearRequestLog clears all recorded requests
func (s *Server) ClearRequestLog() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.requestLog = make([]RequestRecord, 0)
}

// Reset clears all resources, handlers, errors, and request logs
func (s *Server) Reset() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.resources = make(map[string]interface{})
	s.handlers = make(map[string]http.HandlerFunc)
	s.errorResponses = make(map[string]*ErrorResponse)
	s.requestLog = make([]RequestRecord, 0)
	s.responseDelay = 0
}

// handleRequest is the main request handler
func (s *Server) handleRequest(w http.ResponseWriter, r *http.Request) {
	debugLog("handleRequest: %s %s", r.Method, r.URL.Path)
	s.mu.Lock()
	// Record request
	body := ""
	if r.Body != nil {
		// Read the entire body, handling both known and unknown content lengths
		bodyBytes, err := io.ReadAll(r.Body)
		if err == nil {
			body = string(bodyBytes)
		}
		// Reset body for later reading
		r.Body = &readCloser{data: bodyBytes}
	}
	s.requestLog = append(s.requestLog, RequestRecord{
		Method:    r.Method,
		Path:      r.URL.Path,
		Body:      body,
		Timestamp: time.Now(),
	})
	delay := s.responseDelay
	s.mu.Unlock()

	// Apply delay if configured
	if delay > 0 {
		time.Sleep(delay)
	}

	// Check for custom error response
	s.mu.RLock()
	if errResp, ok := s.errorResponses[r.URL.Path]; ok {
		s.mu.RUnlock()
		s.writeJSONResponse(w, errResp.StatusCode, errResp.Body)
		return
	}

	// Check for custom handler
	for pattern, handler := range s.handlers {
		if matched, _ := regexp.MatchString(pattern, r.URL.Path); matched {
			s.mu.RUnlock()
			handler(w, r)
			return
		}
	}
	s.mu.RUnlock()

	// Handle F5 XC special API patterns
	if s.handleSpecialEndpoints(w, r) {
		return
	}

	// Normalize path for resource storage - all API paths use the same pattern
	// /api/config/namespaces/{ns}/{type}/{name}
	// /api/register/namespaces/{ns}/{type}/{name}
	// /api/web/namespaces/{ns}/{type}/{name}
	// /api/system/{type}/{name}
	path := r.URL.Path

	// Default CRUD handling based on HTTP method
	switch r.Method {
	case http.MethodGet:
		s.handleGet(w, r, path)
	case http.MethodPost:
		s.handlePost(w, r, path)
	case http.MethodPut:
		s.handlePut(w, r, path)
	case http.MethodDelete:
		s.handleDelete(w, r, path)
	default:
		s.writeErrorResponse(w, http.StatusMethodNotAllowed, "METHOD_NOT_ALLOWED", "Method not allowed")
	}
}

// handleSpecialEndpoints handles F5 XC specific API patterns
func (s *Server) handleSpecialEndpoints(w http.ResponseWriter, r *http.Request) bool {
	path := r.URL.Path

	// Handle cascade_delete: POST /api/.../namespaces/{name}/cascade_delete
	// or POST /api/.../resource_type/{name}/cascade_delete
	if strings.Contains(path, "/cascade_delete") && r.Method == http.MethodPost {
		debugLog("cascade_delete: processing path=%s", path)
		// Extract the resource path by removing /cascade_delete suffix
		resourcePath := strings.TrimSuffix(path, "/cascade_delete")
		debugLog("cascade_delete: resourcePath=%s", resourcePath)

		s.mu.Lock()
		// Direct match - delete the resource at this exact path
		if _, exists := s.resources[resourcePath]; exists {
			debugLog("cascade_delete: found direct match, deleting %s", resourcePath)
			delete(s.resources, resourcePath)
		}

		// Also delete any child resources (resources that have this path as a prefix)
		for key := range s.resources {
			if strings.HasPrefix(key, resourcePath+"/") {
				debugLog("cascade_delete: deleting child resource %s", key)
				delete(s.resources, key)
			}
		}
		s.mu.Unlock()
		s.writeJSONResponse(w, http.StatusOK, map[string]string{"status": "deleted"})
		return true
	}

	// Handle list endpoints: GET /api/config/namespaces/{ns}/{resource_type}
	if r.Method == http.MethodGet && !strings.HasSuffix(path, "/") {
		parts := strings.Split(strings.TrimPrefix(path, "/"), "/")
		// Check if this is a list operation (path ends with resource type, not resource name)
		if len(parts) == 5 && parts[0] == "api" && parts[1] == "config" && parts[2] == "namespaces" {
			// This looks like a list request
			s.mu.RLock()
			items := make([]interface{}, 0)
			prefix := path + "/"
			for key, val := range s.resources {
				if strings.HasPrefix(key, prefix) {
					items = append(items, val)
				}
			}
			s.mu.RUnlock()

			// Return list response (empty list if no items)
			s.writeJSONResponse(w, http.StatusOK, map[string]interface{}{
				"items": items,
			})
			return true
		}
	}

	return false
}

// handleGet handles GET requests (Read operations)
func (s *Server) handleGet(w http.ResponseWriter, r *http.Request, path string) {
	s.mu.RLock()
	resource, exists := s.resources[path]
	s.mu.RUnlock()

	if !exists {
		s.writeErrorResponse(w, http.StatusNotFound, "NOT_FOUND", fmt.Sprintf("Resource not found: %s", path))
		return
	}

	s.writeJSONResponse(w, http.StatusOK, resource)
}

// handlePost handles POST requests (Create operations)
func (s *Server) handlePost(w http.ResponseWriter, r *http.Request, path string) {
	var requestBody map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		s.writeErrorResponse(w, http.StatusBadRequest, "BAD_REQUEST", "Invalid JSON body")
		return
	}

	// Extract name from metadata
	metadata, ok := requestBody["metadata"].(map[string]interface{})
	if !ok {
		s.writeErrorResponse(w, http.StatusBadRequest, "BAD_REQUEST", "Missing metadata")
		return
	}

	name, ok := metadata["name"].(string)
	if !ok || name == "" {
		s.writeErrorResponse(w, http.StatusBadRequest, "BAD_REQUEST", "Missing or invalid name in metadata")
		return
	}

	// Extract resource type from path for validation and defaults
	resourceType := extractResourceTypeFromPath(path)

	// Validate name based on resource type (DNS domain requires valid domain format)
	if err := s.validateResourceName(resourceType, name); err != nil {
		s.writeErrorResponse(w, http.StatusBadRequest, "BAD_REQUEST", err.Error())
		return
	}

	// Build resource path
	resourcePath := path + "/" + name

	// Extract namespace from path if present
	namespace := extractNamespaceFromPath(path)

	s.mu.Lock()
	// Check if resource already exists
	if _, exists := s.resources[resourcePath]; exists {
		s.mu.Unlock()
		s.writeErrorResponse(w, http.StatusConflict, "ALREADY_EXISTS", fmt.Sprintf("Resource already exists: %s", name))
		return
	}

	// Create response with system metadata and resource-specific defaults
	response := s.buildResponse(requestBody, namespace, resourceType)
	s.resources[resourcePath] = response
	s.mu.Unlock()

	s.writeJSONResponse(w, http.StatusOK, response)
}

// handlePut handles PUT requests (Update operations)
func (s *Server) handlePut(w http.ResponseWriter, r *http.Request, path string) {
	debugLog("handlePut: path=%s", path)
	var requestBody map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		s.writeErrorResponse(w, http.StatusBadRequest, "BAD_REQUEST", "Invalid JSON body")
		return
	}
	debugLog("handlePut: request body=%+v", requestBody)

	// Extract namespace and resource type from path
	namespace := extractNamespaceFromPath(path)
	resourceType := extractResourceTypeFromPath(path)

	s.mu.Lock()
	// Check if resource exists
	if _, exists := s.resources[path]; !exists {
		s.mu.Unlock()
		s.writeErrorResponse(w, http.StatusNotFound, "NOT_FOUND", fmt.Sprintf("Resource not found: %s", path))
		return
	}

	// Update resource with resource-specific defaults
	response := s.buildResponse(requestBody, namespace, resourceType)
	s.resources[path] = response
	s.mu.Unlock()

	debugLog("handlePut: response=%+v", response)
	s.writeJSONResponse(w, http.StatusOK, response)
}

// handleDelete handles DELETE requests
func (s *Server) handleDelete(w http.ResponseWriter, r *http.Request, path string) {
	s.mu.Lock()
	_, exists := s.resources[path]
	if !exists {
		s.mu.Unlock()
		s.writeErrorResponse(w, http.StatusNotFound, "NOT_FOUND", fmt.Sprintf("Resource not found: %s", path))
		return
	}

	delete(s.resources, path)
	s.mu.Unlock()

	s.writeJSONResponse(w, http.StatusOK, map[string]string{"status": "deleted"})
}

// extractNamespaceFromPath extracts the namespace from F5 XC API paths
func extractNamespaceFromPath(path string) string {
	// Expected format: /api/{apiType}/namespaces/{namespace}/{resourceType}
	// e.g., /api/config/namespaces/system/http_loadbalancers
	parts := strings.Split(strings.TrimPrefix(path, "/"), "/")
	for i, part := range parts {
		if part == "namespaces" && i+1 < len(parts) {
			return parts[i+1]
		}
	}
	return "system" // Default to system namespace
}

// buildResponse creates a standard F5 XC API response from a request body
func (s *Server) buildResponse(requestBody map[string]interface{}, namespace string, resourceType string) map[string]interface{} {
	response := make(map[string]interface{})
	tenant := "mock-tenant"

	// Copy metadata and ensure namespace is set correctly
	if metadata, ok := requestBody["metadata"].(map[string]interface{}); ok {
		metaCopy := make(map[string]interface{})
		for k, v := range metadata {
			metaCopy[k] = v
		}
		// Add UID if not present
		if _, ok := metaCopy["uid"]; !ok {
			metaCopy["uid"] = generateUID()
		}
		// Ensure namespace is set correctly from path
		if namespace != "" {
			metaCopy["namespace"] = namespace
		}
		response["metadata"] = metaCopy
	}

	// Copy spec and add tenant to nested objects
	if spec, ok := requestBody["spec"].(map[string]interface{}); ok {
		specCopy := s.deepCopyWithTenant(spec, tenant)
		// Apply resource-specific defaults AFTER copying from request
		s.applyResourceDefaults(specCopy, resourceType)
		response["spec"] = specCopy
	} else {
		// If no spec was provided, create one with defaults
		specCopy := make(map[string]interface{})
		s.applyResourceDefaults(specCopy, resourceType)
		if len(specCopy) > 0 {
			response["spec"] = specCopy
		}
	}

	// Add system metadata
	response["system_metadata"] = map[string]interface{}{
		"uid":                     generateUID(),
		"creation_timestamp":     time.Now().UTC().Format(time.RFC3339),
		"modification_timestamp": time.Now().UTC().Format(time.RFC3339),
		"creator_class":          "API",
		"creator_id":             "mock-server",
		"tenant":                 tenant,
	}

	return response
}

// applyResourceDefaults adds computed default values based on resource type
// These mimic the real F5 XC API behavior of adding default values to responses
func (s *Server) applyResourceDefaults(spec map[string]interface{}, resourceType string) {
	switch resourceType {
	case "secret_policy_rules", "service_policy_rules":
		// Policy rules have a default action of DENY
		if _, ok := spec["action"]; !ok {
			spec["action"] = "DENY"
		}
	case "rate_limiter_policys":
		// Rate limiter has default action
		if _, ok := spec["action"]; !ok {
			spec["action"] = "DENY"
		}
	}
}

// validateResourceName validates the resource name based on resource type
// Returns an error if the name is invalid for the resource type
func (s *Server) validateResourceName(resourceType, name string) error {
	switch resourceType {
	case "dns_domains":
		// DNS domains must be valid domain names (contain at least one dot)
		if !isValidDomainName(name) {
			return fmt.Errorf("Invalid domain name: %s - must be a valid DNS domain", name)
		}
	}
	return nil
}

// isValidDomainName checks if the name is a valid domain name
func isValidDomainName(name string) bool {
	// Domain names must contain at least one dot and have valid characters
	// F5 XC requires lowercase domain names
	if !strings.Contains(name, ".") {
		return false
	}
	// Check for any uppercase letters - F5 XC rejects uppercase in domain names
	for _, c := range name {
		if c >= 'A' && c <= 'Z' {
			return false
		}
	}
	// Check for valid domain name characters (lowercase only)
	domainPattern := regexp.MustCompile(`^[a-z0-9]([a-z0-9\-]*[a-z0-9])?(\.[a-z0-9]([a-z0-9\-]*[a-z0-9])?)*$`)
	return domainPattern.MatchString(name)
}

// extractResourceTypeFromPath extracts the resource type from the API path
// e.g., /api/config/namespaces/system/secret_policy_rules -> secret_policy_rules
func extractResourceTypeFromPath(path string) string {
	parts := strings.Split(strings.TrimPrefix(path, "/"), "/")
	// The resource type is typically the last path segment for collection paths
	// or second-to-last for individual resource paths
	if len(parts) >= 1 {
		// For paths like /api/config/namespaces/ns/resource_type
		// the resource type is the last element
		return parts[len(parts)-1]
	}
	return ""
}

// deepCopyWithTenant recursively copies a map and adds tenant to objects that reference other resources
func (s *Server) deepCopyWithTenant(input map[string]interface{}, tenant string) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range input {
		result[k] = s.addTenantToValue(v, tenant)
	}
	return result
}

// addTenantToValue adds tenant field to maps that look like resource references
func (s *Server) addTenantToValue(value interface{}, tenant string) interface{} {
	switch v := value.(type) {
	case map[string]interface{}:
		result := make(map[string]interface{})
		for key, val := range v {
			result[key] = s.addTenantToValue(val, tenant)
		}
		// If this looks like a resource reference (has name field)
		// add tenant if not present - F5 XC API adds tenant to any object with a name
		if _, hasName := result["name"]; hasName {
			if _, hasTenant := result["tenant"]; !hasTenant {
				result["tenant"] = tenant
			}
		}
		return result
	case []interface{}:
		resultSlice := make([]interface{}, len(v))
		for i, item := range v {
			resultSlice[i] = s.addTenantToValue(item, tenant)
		}
		return resultSlice
	default:
		return value
	}
}

// writeJSONResponse writes a JSON response
func (s *Server) writeJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

// writeErrorResponse writes an F5 XC-style error response
func (s *Server) writeErrorResponse(w http.ResponseWriter, statusCode int, code, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"code":    code,
		"message": message,
	})
}

// generateUID generates a mock UID
func generateUID() string {
	return fmt.Sprintf("mock-uid-%d", time.Now().UnixNano())
}

// readCloser is a helper to reset request body after reading
type readCloser struct {
	data []byte
	pos  int
}

func (rc *readCloser) Read(p []byte) (n int, err error) {
	if rc.pos >= len(rc.data) {
		return 0, io.EOF
	}
	n = copy(p, rc.data[rc.pos:])
	rc.pos += n
	return n, nil
}

func (rc *readCloser) Close() error {
	return nil
}

// ResourcePath builds a resource API path
func ResourcePath(namespace, resourceType, name string) string {
	return fmt.Sprintf("/api/config/namespaces/%s/%s/%s", namespace, resourceType, name)
}

// ListPath builds a list API path
func ListPath(namespace, resourceType string) string {
	return fmt.Sprintf("/api/config/namespaces/%s/%s", namespace, resourceType)
}

// DataSourcePath builds a data source API path (same as resource path for GET)
func DataSourcePath(namespace, resourceType, name string) string {
	return ResourcePath(namespace, resourceType, name)
}

// ExtractResourceInfo extracts namespace, resource type, and name from a path
func ExtractResourceInfo(path string) (namespace, resourceType, name string, ok bool) {
	// Expected format: /api/config/namespaces/{namespace}/{resourceType}/{name}
	parts := strings.Split(strings.TrimPrefix(path, "/"), "/")
	if len(parts) < 5 {
		return "", "", "", false
	}
	return parts[3], parts[4], parts[len(parts)-1], true
}
