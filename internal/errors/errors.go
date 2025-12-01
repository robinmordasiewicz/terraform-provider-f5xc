// Copyright (c) F5XC Community
// SPDX-License-Identifier: MPL-2.0

// Package errors provides structured error types for F5 XC Terraform provider
// following Terraform Plugin Framework best practices.
package errors

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// ErrorCode represents specific error types for better handling
type ErrorCode string

const (
	// API errors
	ErrCodeNotFound     ErrorCode = "NOT_FOUND"
	ErrCodeUnauthorized ErrorCode = "UNAUTHORIZED"
	ErrCodeForbidden    ErrorCode = "FORBIDDEN"
	ErrCodeConflict     ErrorCode = "CONFLICT"
	ErrCodeRateLimit    ErrorCode = "RATE_LIMIT"
	ErrCodeServerError  ErrorCode = "SERVER_ERROR"
	ErrCodeBadRequest   ErrorCode = "BAD_REQUEST"
	ErrCodeTimeout      ErrorCode = "TIMEOUT"
	ErrCodeNetworkError ErrorCode = "NETWORK_ERROR"

	// Resource errors
	ErrCodeValidation    ErrorCode = "VALIDATION"
	ErrCodeStateRead     ErrorCode = "STATE_READ"
	ErrCodeStateWrite    ErrorCode = "STATE_WRITE"
	ErrCodeConfiguration ErrorCode = "CONFIGURATION"
)

// F5XCError is a structured error type for the provider
type F5XCError struct {
	Code       ErrorCode
	Message    string
	Resource   string
	Operation  string
	StatusCode int
	Details    map[string]interface{}
	Wrapped    error
}

// Error implements the error interface
func (e *F5XCError) Error() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("[%s] %s", e.Code, e.Message))

	if e.Resource != "" {
		sb.WriteString(fmt.Sprintf(" (resource: %s)", e.Resource))
	}
	if e.Operation != "" {
		sb.WriteString(fmt.Sprintf(" (operation: %s)", e.Operation))
	}
	if e.StatusCode != 0 {
		sb.WriteString(fmt.Sprintf(" (status: %d)", e.StatusCode))
	}
	if e.Wrapped != nil {
		sb.WriteString(fmt.Sprintf(": %v", e.Wrapped))
	}

	return sb.String()
}

// Unwrap returns the wrapped error
func (e *F5XCError) Unwrap() error {
	return e.Wrapped
}

// IsRetryable returns true if the error is potentially transient
func (e *F5XCError) IsRetryable() bool {
	switch e.Code {
	case ErrCodeRateLimit, ErrCodeTimeout, ErrCodeNetworkError, ErrCodeServerError:
		return true
	default:
		return false
	}
}

// IsNotFound returns true if the resource was not found
func (e *F5XCError) IsNotFound() bool {
	return e.Code == ErrCodeNotFound
}

// APIErrorResponse represents the F5 XC API error response structure
type APIErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details []struct {
		Type   string `json:"@type"`
		Reason string `json:"reason"`
		Domain string `json:"domain"`
	} `json:"details"`
}

// NewAPIError creates an error from an API response
func NewAPIError(statusCode int, body []byte, resource, operation string) *F5XCError {
	err := &F5XCError{
		Resource:   resource,
		Operation:  operation,
		StatusCode: statusCode,
		Details:    make(map[string]interface{}),
	}

	// Set error code based on status
	switch statusCode {
	case http.StatusNotFound:
		err.Code = ErrCodeNotFound
		err.Message = fmt.Sprintf("%s not found", resource)
	case http.StatusUnauthorized:
		err.Code = ErrCodeUnauthorized
		err.Message = "API authentication failed - check your API token"
	case http.StatusForbidden:
		err.Code = ErrCodeForbidden
		err.Message = "Access denied - insufficient permissions"
	case http.StatusConflict:
		err.Code = ErrCodeConflict
		err.Message = fmt.Sprintf("%s already exists or has a conflicting state", resource)
	case http.StatusTooManyRequests:
		err.Code = ErrCodeRateLimit
		err.Message = "API rate limit exceeded - retry after a delay"
	case http.StatusBadRequest:
		err.Code = ErrCodeBadRequest
		err.Message = "Invalid request parameters"
	default:
		if statusCode >= 500 {
			err.Code = ErrCodeServerError
			err.Message = "F5 XC API server error"
		} else {
			err.Code = ErrCodeBadRequest
			err.Message = fmt.Sprintf("API request failed with status %d", statusCode)
		}
	}

	// Try to parse API error response for more details
	if len(body) > 0 {
		var apiErr APIErrorResponse
		if jsonErr := json.Unmarshal(body, &apiErr); jsonErr == nil {
			if apiErr.Message != "" {
				err.Message = apiErr.Message
			}
			if apiErr.Code != "" {
				err.Details["api_code"] = apiErr.Code
			}
			if len(apiErr.Details) > 0 {
				err.Details["api_details"] = apiErr.Details
			}
		} else {
			// Store raw response if JSON parsing fails
			err.Details["raw_response"] = string(body)
		}
	}

	return err
}

// NewNotFoundError creates a not found error
func NewNotFoundError(resource, name, namespace string) *F5XCError {
	return &F5XCError{
		Code:     ErrCodeNotFound,
		Message:  fmt.Sprintf("%s '%s' not found in namespace '%s'", resource, name, namespace),
		Resource: resource,
		Details: map[string]interface{}{
			"name":      name,
			"namespace": namespace,
		},
	}
}

// NewValidationError creates a validation error
func NewValidationError(resource, field, message string) *F5XCError {
	return &F5XCError{
		Code:     ErrCodeValidation,
		Message:  fmt.Sprintf("validation failed for %s.%s: %s", resource, field, message),
		Resource: resource,
		Details: map[string]interface{}{
			"field": field,
		},
	}
}

// NewTimeoutError creates a timeout error
func NewTimeoutError(resource, operation string, wrapped error) *F5XCError {
	return &F5XCError{
		Code:      ErrCodeTimeout,
		Message:   fmt.Sprintf("operation timed out: %s %s", operation, resource),
		Resource:  resource,
		Operation: operation,
		Wrapped:   wrapped,
	}
}

// NewNetworkError creates a network error
func NewNetworkError(wrapped error) *F5XCError {
	return &F5XCError{
		Code:    ErrCodeNetworkError,
		Message: "network error communicating with F5 XC API",
		Wrapped: wrapped,
	}
}

// NewConfigurationError creates a configuration error
func NewConfigurationError(message string) *F5XCError {
	return &F5XCError{
		Code:    ErrCodeConfiguration,
		Message: message,
	}
}

// DiagnosticHelpers provides methods to add errors to diagnostics

// AddError adds a structured error to diagnostics
func AddError(diags *diag.Diagnostics, err *F5XCError) {
	caser := cases.Title(language.English)
	summary := fmt.Sprintf("%s Error", caser.String(string(err.Code)))
	diags.AddError(summary, err.Error())
}

// AddWarning adds a warning to diagnostics
func AddWarning(diags *diag.Diagnostics, summary, detail string) {
	diags.AddWarning(summary, detail)
}

// AddAttributeError adds an attribute-specific error
func AddAttributeError(diags *diag.Diagnostics, path, summary, detail string) {
	diags.AddError(fmt.Sprintf("%s: %s", path, summary), detail)
}

// CreateDiagnostic creates a simple error diagnostic from a resource operation error
func CreateDiagnostic(operation, resourceType string, err error) diag.Diagnostics {
	var diags diag.Diagnostics

	if f5xcErr, ok := err.(*F5XCError); ok {
		AddError(&diags, f5xcErr)
	} else {
		diags.AddError(
			fmt.Sprintf("Error %s %s", operation, resourceType),
			err.Error(),
		)
	}

	return diags
}

// WrapError wraps an error with additional context
func WrapError(err error, resource, operation string) *F5XCError {
	if f5xcErr, ok := err.(*F5XCError); ok {
		f5xcErr.Resource = resource
		f5xcErr.Operation = operation
		return f5xcErr
	}

	return &F5XCError{
		Code:      ErrCodeServerError,
		Message:   err.Error(),
		Resource:  resource,
		Operation: operation,
		Wrapped:   err,
	}
}
