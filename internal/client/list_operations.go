// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package client

import (
	"context"
)

// ListResponse represents the response from a list API call
type ListResponse struct {
	Items []struct {
		Metadata Metadata               `json:"metadata"`
		Spec     map[string]interface{} `json:"spec,omitempty"`
	} `json:"items"`
}

// ListNamespaces retrieves all namespaces from the F5 XC API
func (c *Client) ListNamespaces(ctx context.Context) (*ListResponse, error) {
	var result ListResponse
	path := "/api/web/namespaces"
	err := c.Get(ctx, path, &result)
	return &result, err
}

// ListHTTPLoadBalancers retrieves all HTTP load balancers in a namespace
func (c *Client) ListHTTPLoadBalancers(ctx context.Context, namespace string) (*ListResponse, error) {
	var result ListResponse
	path := "/api/config/namespaces/" + namespace + "/http_loadbalancers"
	err := c.Get(ctx, path, &result)
	return &result, err
}

// ListOriginPools retrieves all origin pools in a namespace
func (c *Client) ListOriginPools(ctx context.Context, namespace string) (*ListResponse, error) {
	var result ListResponse
	path := "/api/config/namespaces/" + namespace + "/origin_pools"
	err := c.Get(ctx, path, &result)
	return &result, err
}

// ListHealthchecks retrieves all healthchecks in a namespace
func (c *Client) ListHealthchecks(ctx context.Context, namespace string) (*ListResponse, error) {
	var result ListResponse
	path := "/api/config/namespaces/" + namespace + "/healthchecks"
	err := c.Get(ctx, path, &result)
	return &result, err
}

// ListAppFirewalls retrieves all app firewalls in a namespace
func (c *Client) ListAppFirewalls(ctx context.Context, namespace string) (*ListResponse, error) {
	var result ListResponse
	path := "/api/config/namespaces/" + namespace + "/app_firewalls"
	err := c.Get(ctx, path, &result)
	return &result, err
}

// ListServicePolicies retrieves all service policies in a namespace
func (c *Client) ListServicePolicies(ctx context.Context, namespace string) (*ListResponse, error) {
	var result ListResponse
	path := "/api/config/namespaces/" + namespace + "/service_policys"
	err := c.Get(ctx, path, &result)
	return &result, err
}

// ListIPPrefixSets retrieves all IP prefix sets in a namespace
func (c *Client) ListIPPrefixSets(ctx context.Context, namespace string) (*ListResponse, error) {
	var result ListResponse
	path := "/api/config/namespaces/" + namespace + "/ip_prefix_sets"
	err := c.Get(ctx, path, &result)
	return &result, err
}

// ListRateLimiters retrieves all rate limiters in a namespace
func (c *Client) ListRateLimiters(ctx context.Context, namespace string) (*ListResponse, error) {
	var result ListResponse
	path := "/api/config/namespaces/" + namespace + "/rate_limiters"
	err := c.Get(ctx, path, &result)
	return &result, err
}

// ListUserIdentifications retrieves all user identifications in a namespace
func (c *Client) ListUserIdentifications(ctx context.Context, namespace string) (*ListResponse, error) {
	var result ListResponse
	path := "/api/config/namespaces/" + namespace + "/user_identifications"
	err := c.Get(ctx, path, &result)
	return &result, err
}

// ListMaliciousUserMitigations retrieves all malicious user mitigations in a namespace
func (c *Client) ListMaliciousUserMitigations(ctx context.Context, namespace string) (*ListResponse, error) {
	var result ListResponse
	path := "/api/config/namespaces/" + namespace + "/malicious_user_mitigations"
	err := c.Get(ctx, path, &result)
	return &result, err
}

// CascadeDeleteNamespace deletes a namespace and all its contained resources.
// This is used instead of the regular DELETE endpoint which returns 501 Not Implemented.
// The cascade_delete endpoint is a POST request that deletes the namespace and all
// resources within it.
func (c *Client) CascadeDeleteNamespace(ctx context.Context, name string) error {
	path := "/api/web/namespaces/" + name + "/cascade_delete"
	// POST with empty body triggers cascade delete
	return c.Post(ctx, path, struct{}{}, nil)
}
