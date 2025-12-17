// Copyright (c) F5XC Community
// SPDX-License-Identifier: MPL-2.0

package errors

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TestF5XCErrorError(t *testing.T) {
	tests := []struct {
		name     string
		err      *F5XCError
		contains []string
	}{
		{
			name: "basic error",
			err: &F5XCError{
				Code:    ErrCodeNotFound,
				Message: "resource not found",
			},
			contains: []string{"NOT_FOUND", "resource not found"},
		},
		{
			name: "error with resource",
			err: &F5XCError{
				Code:     ErrCodeNotFound,
				Message:  "not found",
				Resource: "namespace",
			},
			contains: []string{"resource: namespace"},
		},
		{
			name: "error with operation",
			err: &F5XCError{
				Code:      ErrCodeTimeout,
				Message:   "timeout",
				Operation: "create",
			},
			contains: []string{"operation: create"},
		},
		{
			name: "error with status code",
			err: &F5XCError{
				Code:       ErrCodeServerError,
				Message:    "server error",
				StatusCode: 500,
			},
			contains: []string{"status: 500"},
		},
		{
			name: "error with wrapped error",
			err: &F5XCError{
				Code:    ErrCodeNetworkError,
				Message: "network error",
				Wrapped: errors.New("connection refused"),
			},
			contains: []string{"connection refused"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errStr := tt.err.Error()
			for _, s := range tt.contains {
				if !strings.Contains(errStr, s) {
					t.Errorf("Error() = %q, expected to contain %q", errStr, s)
				}
			}
		})
	}
}

func TestF5XCErrorUnwrap(t *testing.T) {
	innerErr := errors.New("inner error")
	err := &F5XCError{
		Code:    ErrCodeServerError,
		Message: "outer error",
		Wrapped: innerErr,
	}

	unwrapped := err.Unwrap()
	if unwrapped != innerErr {
		t.Errorf("Unwrap() = %v, expected %v", unwrapped, innerErr)
	}

	// Test with nil wrapped
	errNoWrap := &F5XCError{
		Code:    ErrCodeNotFound,
		Message: "no wrap",
	}
	if errNoWrap.Unwrap() != nil {
		t.Error("Unwrap() should return nil when no wrapped error")
	}
}

func TestF5XCErrorIsRetryable(t *testing.T) {
	tests := []struct {
		code     ErrorCode
		expected bool
	}{
		{ErrCodeRateLimit, true},
		{ErrCodeTimeout, true},
		{ErrCodeNetworkError, true},
		{ErrCodeServerError, true},
		{ErrCodeNotFound, false},
		{ErrCodeUnauthorized, false},
		{ErrCodeForbidden, false},
		{ErrCodeConflict, false},
		{ErrCodeBadRequest, false},
		{ErrCodeValidation, false},
	}

	for _, tt := range tests {
		t.Run(string(tt.code), func(t *testing.T) {
			err := &F5XCError{Code: tt.code}
			if err.IsRetryable() != tt.expected {
				t.Errorf("IsRetryable() for %s = %v, expected %v", tt.code, err.IsRetryable(), tt.expected)
			}
		})
	}
}

func TestF5XCErrorIsNotFound(t *testing.T) {
	tests := []struct {
		code     ErrorCode
		expected bool
	}{
		{ErrCodeNotFound, true},
		{ErrCodeUnauthorized, false},
		{ErrCodeServerError, false},
	}

	for _, tt := range tests {
		t.Run(string(tt.code), func(t *testing.T) {
			err := &F5XCError{Code: tt.code}
			if err.IsNotFound() != tt.expected {
				t.Errorf("IsNotFound() for %s = %v, expected %v", tt.code, err.IsNotFound(), tt.expected)
			}
		})
	}
}

func TestNewAPIError(t *testing.T) {
	tests := []struct {
		name         string
		statusCode   int
		body         []byte
		resource     string
		operation    string
		expectedCode ErrorCode
	}{
		{
			name:         "not found",
			statusCode:   http.StatusNotFound,
			body:         nil,
			resource:     "namespace",
			operation:    "read",
			expectedCode: ErrCodeNotFound,
		},
		{
			name:         "unauthorized",
			statusCode:   http.StatusUnauthorized,
			body:         nil,
			resource:     "namespace",
			operation:    "create",
			expectedCode: ErrCodeUnauthorized,
		},
		{
			name:         "forbidden",
			statusCode:   http.StatusForbidden,
			body:         nil,
			resource:     "namespace",
			operation:    "delete",
			expectedCode: ErrCodeForbidden,
		},
		{
			name:         "conflict",
			statusCode:   http.StatusConflict,
			body:         nil,
			resource:     "namespace",
			operation:    "create",
			expectedCode: ErrCodeConflict,
		},
		{
			name:         "rate limit",
			statusCode:   http.StatusTooManyRequests,
			body:         nil,
			resource:     "namespace",
			operation:    "list",
			expectedCode: ErrCodeRateLimit,
		},
		{
			name:         "bad request",
			statusCode:   http.StatusBadRequest,
			body:         nil,
			resource:     "namespace",
			operation:    "create",
			expectedCode: ErrCodeBadRequest,
		},
		{
			name:         "server error",
			statusCode:   http.StatusInternalServerError,
			body:         nil,
			resource:     "namespace",
			operation:    "read",
			expectedCode: ErrCodeServerError,
		},
		{
			name:         "with JSON body",
			statusCode:   http.StatusBadRequest,
			body:         []byte(`{"code": "INVALID_ARGUMENT", "message": "field is required"}`),
			resource:     "namespace",
			operation:    "create",
			expectedCode: ErrCodeBadRequest,
		},
		{
			name:         "with invalid JSON body",
			statusCode:   http.StatusBadRequest,
			body:         []byte(`not json`),
			resource:     "namespace",
			operation:    "create",
			expectedCode: ErrCodeBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := NewAPIError(tt.statusCode, tt.body, tt.resource, tt.operation)

			if err.Code != tt.expectedCode {
				t.Errorf("NewAPIError() code = %v, expected %v", err.Code, tt.expectedCode)
			}
			if err.Resource != tt.resource {
				t.Errorf("NewAPIError() resource = %v, expected %v", err.Resource, tt.resource)
			}
			if err.Operation != tt.operation {
				t.Errorf("NewAPIError() operation = %v, expected %v", err.Operation, tt.operation)
			}
			if err.StatusCode != tt.statusCode {
				t.Errorf("NewAPIError() statusCode = %v, expected %v", err.StatusCode, tt.statusCode)
			}
		})
	}
}

func TestNewNotFoundError(t *testing.T) {
	err := NewNotFoundError("namespace", "test-ns", "system")

	if err.Code != ErrCodeNotFound {
		t.Errorf("NewNotFoundError() code = %v, expected %v", err.Code, ErrCodeNotFound)
	}
	if !strings.Contains(err.Message, "test-ns") {
		t.Errorf("NewNotFoundError() message should contain name, got %v", err.Message)
	}
	if !strings.Contains(err.Message, "system") {
		t.Errorf("NewNotFoundError() message should contain namespace, got %v", err.Message)
	}
}

func TestNewValidationError(t *testing.T) {
	err := NewValidationError("namespace", "name", "must not be empty")

	if err.Code != ErrCodeValidation {
		t.Errorf("NewValidationError() code = %v, expected %v", err.Code, ErrCodeValidation)
	}
	if !strings.Contains(err.Message, "name") {
		t.Errorf("NewValidationError() message should contain field name, got %v", err.Message)
	}
}

func TestNewTimeoutError(t *testing.T) {
	wrapped := errors.New("context deadline exceeded")
	err := NewTimeoutError("namespace", "create", wrapped)

	if err.Code != ErrCodeTimeout {
		t.Errorf("NewTimeoutError() code = %v, expected %v", err.Code, ErrCodeTimeout)
	}
	if err.Wrapped != wrapped {
		t.Error("NewTimeoutError() should wrap the original error")
	}
}

func TestNewNetworkError(t *testing.T) {
	wrapped := errors.New("connection refused")
	err := NewNetworkError(wrapped)

	if err.Code != ErrCodeNetworkError {
		t.Errorf("NewNetworkError() code = %v, expected %v", err.Code, ErrCodeNetworkError)
	}
	if err.Wrapped != wrapped {
		t.Error("NewNetworkError() should wrap the original error")
	}
}

func TestNewConfigurationError(t *testing.T) {
	err := NewConfigurationError("API URL is required")

	if err.Code != ErrCodeConfiguration {
		t.Errorf("NewConfigurationError() code = %v, expected %v", err.Code, ErrCodeConfiguration)
	}
	if err.Message != "API URL is required" {
		t.Errorf("NewConfigurationError() message = %v, expected %v", err.Message, "API URL is required")
	}
}

func TestAddError(t *testing.T) {
	var diags diag.Diagnostics
	err := &F5XCError{
		Code:    ErrCodeNotFound,
		Message: "resource not found",
	}

	AddError(&diags, err)

	if !diags.HasError() {
		t.Error("AddError() should add an error to diagnostics")
	}
	if len(diags.Errors()) != 1 {
		t.Errorf("AddError() should add exactly one error, got %d", len(diags.Errors()))
	}
}

func TestAddWarning(t *testing.T) {
	var diags diag.Diagnostics

	AddWarning(&diags, "Warning Title", "Warning detail")

	if diags.HasError() {
		t.Error("AddWarning() should not add an error")
	}
	if len(diags.Warnings()) != 1 {
		t.Errorf("AddWarning() should add exactly one warning, got %d", len(diags.Warnings()))
	}
}

func TestAddAttributeError(t *testing.T) {
	var diags diag.Diagnostics

	AddAttributeError(&diags, "name", "Invalid Value", "Name must not be empty")

	if !diags.HasError() {
		t.Error("AddAttributeError() should add an error")
	}
}

func TestCreateDiagnostic(t *testing.T) {
	t.Run("with F5XCError", func(t *testing.T) {
		f5xcErr := &F5XCError{
			Code:    ErrCodeNotFound,
			Message: "not found",
		}

		diags := CreateDiagnostic("reading", "namespace", f5xcErr)

		if !diags.HasError() {
			t.Error("CreateDiagnostic() should create error diagnostic")
		}
	})

	t.Run("with standard error", func(t *testing.T) {
		stdErr := errors.New("standard error")

		diags := CreateDiagnostic("creating", "namespace", stdErr)

		if !diags.HasError() {
			t.Error("CreateDiagnostic() should create error diagnostic")
		}
	})
}

func TestWrapError(t *testing.T) {
	t.Run("wrap standard error", func(t *testing.T) {
		stdErr := errors.New("original error")

		wrapped := WrapError(stdErr, "namespace", "create")

		if wrapped.Resource != "namespace" {
			t.Errorf("WrapError() resource = %v, expected namespace", wrapped.Resource)
		}
		if wrapped.Operation != "create" {
			t.Errorf("WrapError() operation = %v, expected create", wrapped.Operation)
		}
		if wrapped.Wrapped != stdErr {
			t.Error("WrapError() should preserve original error")
		}
	})

	t.Run("wrap F5XCError", func(t *testing.T) {
		f5xcErr := &F5XCError{
			Code:    ErrCodeNotFound,
			Message: "not found",
		}

		wrapped := WrapError(f5xcErr, "namespace", "read")

		if wrapped.Resource != "namespace" {
			t.Errorf("WrapError() should update resource, got %v", wrapped.Resource)
		}
		if wrapped.Operation != "read" {
			t.Errorf("WrapError() should update operation, got %v", wrapped.Operation)
		}
		if wrapped.Code != ErrCodeNotFound {
			t.Error("WrapError() should preserve original code")
		}
	})
}

func TestAPIErrorResponseParsing(t *testing.T) {
	jsonBody := `{
		"code": "INVALID_ARGUMENT",
		"message": "Field validation failed",
		"details": [
			{
				"@type": "type.googleapis.com/google.rpc.BadRequest",
				"reason": "name is required",
				"domain": "f5xc.io"
			}
		]
	}`

	var apiErr APIErrorResponse
	if err := json.Unmarshal([]byte(jsonBody), &apiErr); err != nil {
		t.Fatalf("Failed to unmarshal APIErrorResponse: %v", err)
	}

	if apiErr.Code != "INVALID_ARGUMENT" {
		t.Errorf("APIErrorResponse.Code = %v, expected INVALID_ARGUMENT", apiErr.Code)
	}
	if apiErr.Message != "Field validation failed" {
		t.Errorf("APIErrorResponse.Message = %v, expected Field validation failed", apiErr.Message)
	}
	if len(apiErr.Details) != 1 {
		t.Errorf("APIErrorResponse.Details length = %d, expected 1", len(apiErr.Details))
	}
}

func TestErrorCodeConstants(t *testing.T) {
	// Verify error codes are distinct
	codes := []ErrorCode{
		ErrCodeNotFound,
		ErrCodeUnauthorized,
		ErrCodeForbidden,
		ErrCodeConflict,
		ErrCodeRateLimit,
		ErrCodeServerError,
		ErrCodeBadRequest,
		ErrCodeTimeout,
		ErrCodeNetworkError,
		ErrCodeValidation,
		ErrCodeStateRead,
		ErrCodeStateWrite,
		ErrCodeConfiguration,
	}

	seen := make(map[ErrorCode]bool)
	for _, code := range codes {
		if seen[code] {
			t.Errorf("Duplicate error code: %v", code)
		}
		seen[code] = true
	}
}
