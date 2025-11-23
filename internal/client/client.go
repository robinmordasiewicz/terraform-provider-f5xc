// Copyright (c) F5XC Community
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"math"
	"net/http"
	"time"

	f5xcerrors "github.com/f5xc/terraform-provider-f5xc/internal/errors"
)

// Client configuration defaults
const (
	DefaultTimeout        = 30 * time.Second
	DefaultMaxRetries     = 3
	DefaultRetryWaitMin   = 1 * time.Second
	DefaultRetryWaitMax   = 30 * time.Second
	DefaultRateLimitDelay = 60 * time.Second
)

// Client manages communication with the F5 Distributed Cloud API
type Client struct {
	BaseURL      string
	APIToken     string
	HTTPClient   *http.Client
	MaxRetries   int
	RetryWaitMin time.Duration
	RetryWaitMax time.Duration
}

// ClientOption allows customizing the client
type ClientOption func(*Client)

// WithTimeout sets the HTTP client timeout
func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) {
		c.HTTPClient.Timeout = timeout
	}
}

// WithMaxRetries sets the maximum number of retries
func WithMaxRetries(retries int) ClientOption {
	return func(c *Client) {
		c.MaxRetries = retries
	}
}

// WithRetryWait sets the retry wait bounds
func WithRetryWait(min, max time.Duration) ClientOption {
	return func(c *Client) {
		c.RetryWaitMin = min
		c.RetryWaitMax = max
	}
}

// NewClient creates a new F5 Distributed Cloud API client
func NewClient(baseURL, apiToken string, opts ...ClientOption) *Client {
	c := &Client{
		BaseURL:      baseURL,
		APIToken:     apiToken,
		MaxRetries:   DefaultMaxRetries,
		RetryWaitMin: DefaultRetryWaitMin,
		RetryWaitMax: DefaultRetryWaitMax,
		HTTPClient: &http.Client{
			Timeout: DefaultTimeout,
		},
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

// Metadata represents common metadata fields
type Metadata struct {
	Name        string            `json:"name"`
	Namespace   string            `json:"namespace,omitempty"`
	Labels      map[string]string `json:"labels,omitempty"`
	Annotations map[string]string `json:"annotations,omitempty"`
	UID         string            `json:"uid,omitempty"`
}

// isRetryableStatus returns true if the HTTP status code indicates a retryable error
func isRetryableStatus(statusCode int) bool {
	switch statusCode {
	case http.StatusTooManyRequests, // 429 - Rate limited
		http.StatusServiceUnavailable,   // 503
		http.StatusGatewayTimeout,       // 504
		http.StatusBadGateway:           // 502
		return true
	default:
		return statusCode >= 500
	}
}

// calculateBackoff returns the backoff duration for a given attempt
func (c *Client) calculateBackoff(attempt int) time.Duration {
	// Exponential backoff: min * 2^attempt, capped at max
	backoff := float64(c.RetryWaitMin) * math.Pow(2, float64(attempt))
	if backoff > float64(c.RetryWaitMax) {
		backoff = float64(c.RetryWaitMax)
	}
	return time.Duration(backoff)
}

// doRequest performs an HTTP request with retry logic and returns the response
func (c *Client) doRequest(ctx context.Context, method, path string, body interface{}) ([]byte, error) {
	var bodyBytes []byte
	if body != nil {
		var err error
		bodyBytes, err = json.Marshal(body)
		if err != nil {
			return nil, f5xcerrors.WrapError(err, "request", "marshal")
		}
	}

	url := c.BaseURL + path
	var lastErr error

	for attempt := 0; attempt <= c.MaxRetries; attempt++ {
		// Check context before making request
		if ctx.Err() != nil {
			return nil, f5xcerrors.NewTimeoutError("request", method+" "+path, ctx.Err())
		}

		var reqBody io.Reader
		if bodyBytes != nil {
			reqBody = bytes.NewBuffer(bodyBytes)
		}

		req, err := http.NewRequestWithContext(ctx, method, url, reqBody)
		if err != nil {
			return nil, f5xcerrors.WrapError(err, "request", "create")
		}

		req.Header.Set("Authorization", "APIToken "+c.APIToken)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/json")

		resp, err := c.HTTPClient.Do(req)
		if err != nil {
			lastErr = f5xcerrors.NewNetworkError(err)
			// Network errors are retryable
			if attempt < c.MaxRetries {
				select {
				case <-ctx.Done():
					return nil, f5xcerrors.NewTimeoutError("request", method+" "+path, ctx.Err())
				case <-time.After(c.calculateBackoff(attempt)):
					continue
				}
			}
			return nil, lastErr
		}

		respBody, err := io.ReadAll(resp.Body)
		_ = resp.Body.Close() // Error intentionally ignored - body already read, cleanup only
		if err != nil {
			return nil, f5xcerrors.WrapError(err, "response", "read")
		}

		// Success
		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			return respBody, nil
		}

		// Create structured error
		apiErr := f5xcerrors.NewAPIError(resp.StatusCode, respBody, path, method)
		lastErr = apiErr

		// Check if error is retryable
		if !isRetryableStatus(resp.StatusCode) || attempt >= c.MaxRetries {
			return nil, apiErr
		}

		// Calculate backoff (use longer delay for rate limiting)
		backoff := c.calculateBackoff(attempt)
		if resp.StatusCode == http.StatusTooManyRequests {
			backoff = DefaultRateLimitDelay
		}

		select {
		case <-ctx.Done():
			return nil, f5xcerrors.NewTimeoutError("request", method+" "+path, ctx.Err())
		case <-time.After(backoff):
			// Continue to next retry
		}
	}

	return nil, lastErr
}

// Get performs a GET request
func (c *Client) Get(ctx context.Context, path string, result interface{}) error {
	body, err := c.doRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return err
	}
	if result != nil && len(body) > 0 {
		return json.Unmarshal(body, result)
	}
	return nil
}

// Post performs a POST request
func (c *Client) Post(ctx context.Context, path string, data, result interface{}) error {
	body, err := c.doRequest(ctx, http.MethodPost, path, data)
	if err != nil {
		return err
	}
	if result != nil && len(body) > 0 {
		return json.Unmarshal(body, result)
	}
	return nil
}

// Put performs a PUT request
func (c *Client) Put(ctx context.Context, path string, data, result interface{}) error {
	body, err := c.doRequest(ctx, http.MethodPut, path, data)
	if err != nil {
		return err
	}
	if result != nil && len(body) > 0 {
		return json.Unmarshal(body, result)
	}
	return nil
}

// Delete performs a DELETE request
func (c *Client) Delete(ctx context.Context, path string) error {
	_, err := c.doRequest(ctx, http.MethodDelete, path, nil)
	return err
}
