// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

// Package client provides the F5 Distributed Cloud API client with quota management support.
package client

import (
	"context"
	"fmt"
)

// QuotaUsageResponse represents the response from the quota usage API
type QuotaUsageResponse struct {
	Objects map[string]ObjectQuota `json:"objects"`
}

// ObjectQuota represents quota information for a single object type
type ObjectQuota struct {
	Limit LimitValue `json:"limit"`
	Usage UsageValue `json:"usage"`
}

// LimitValue represents the quota limit configuration
type LimitValue struct {
	Maximum int `json:"maximum"`
}

// UsageValue represents the current usage
type UsageValue struct {
	Current int `json:"current"`
}

// QuotaInfo provides a simplified view of quota for a resource type
type QuotaInfo struct {
	ResourceType string
	Limit        int
	Used         int
	Available    int
}

// GetQuotaUsage retrieves the quota usage for a namespace
func (c *Client) GetQuotaUsage(ctx context.Context, namespace string) (*QuotaUsageResponse, error) {
	path := fmt.Sprintf("/api/web/namespaces/%s/quota/usage", namespace)

	var result QuotaUsageResponse
	if err := c.Get(ctx, path, &result); err != nil {
		return nil, fmt.Errorf("failed to get quota usage for namespace %s: %w", namespace, err)
	}

	return &result, nil
}

// GetQuotaInfo retrieves simplified quota information for a specific resource type
func (c *Client) GetQuotaInfo(ctx context.Context, namespace, resourceType string) (*QuotaInfo, error) {
	usage, err := c.GetQuotaUsage(ctx, namespace)
	if err != nil {
		return nil, err
	}

	quota, ok := usage.Objects[resourceType]
	if !ok {
		return nil, fmt.Errorf("no quota information found for resource type %s", resourceType)
	}

	return &QuotaInfo{
		ResourceType: resourceType,
		Limit:        quota.Limit.Maximum,
		Used:         quota.Usage.Current,
		Available:    quota.Limit.Maximum - quota.Usage.Current,
	}, nil
}

// IsQuotaAvailable checks if there is available quota for a resource type
func (c *Client) IsQuotaAvailable(ctx context.Context, namespace, resourceType string) (bool, error) {
	info, err := c.GetQuotaInfo(ctx, namespace, resourceType)
	if err != nil {
		return false, err
	}

	return info.Available > 0, nil
}

// GetAllQuotaInfo retrieves quota information for all resource types in a namespace
func (c *Client) GetAllQuotaInfo(ctx context.Context, namespace string) (map[string]*QuotaInfo, error) {
	usage, err := c.GetQuotaUsage(ctx, namespace)
	if err != nil {
		return nil, err
	}

	result := make(map[string]*QuotaInfo)
	for resourceType, quota := range usage.Objects {
		result[resourceType] = &QuotaInfo{
			ResourceType: resourceType,
			Limit:        quota.Limit.Maximum,
			Used:         quota.Usage.Current,
			Available:    quota.Limit.Maximum - quota.Usage.Current,
		}
	}

	return result, nil
}
