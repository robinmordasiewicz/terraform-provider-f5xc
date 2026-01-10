// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

// Addon Service client types for F5 XC
// This file provides read-only access to addon services (system-managed resources)

package client

import (
	"context"
	"fmt"
)

// AddonService represents a F5XC Addon Service
type AddonService struct {
	Metadata Metadata               `json:"metadata"`
	Spec     map[string]interface{} `json:"spec"`
}

// AddonServiceDetails represents detailed addon service information
// from the GetAddonServiceDetails API
type AddonServiceDetails struct {
	DisplayName                  string                 `json:"display_name,omitempty"`
	AddonServiceGroupName        string                 `json:"addon_service_group_name,omitempty"`
	AddonServiceGroupDisplayName string                 `json:"addon_service_group_display_name,omitempty"`
	Tier                         string                 `json:"tier,omitempty"`
	SelfActivation               map[string]interface{} `json:"self_activation,omitempty"`
	PartiallyManagedActivation   map[string]interface{} `json:"partially_managed_activation,omitempty"`
	ManagedActivation            map[string]interface{} `json:"managed_activation,omitempty"`
}

// AddonServiceActivationStatus represents the activation status response
type AddonServiceActivationStatus struct {
	State string `json:"state,omitempty"`
}

// AddonServiceActivationStatusFull represents full activation status with access info
type AddonServiceActivationStatusFull struct {
	Name               string `json:"name,omitempty"`
	DisplayName        string `json:"display_name,omitempty"`
	AddonServiceStatus string `json:"addon_service_status,omitempty"`
	AccessStatus       string `json:"access_status,omitempty"`
	Tier               string `json:"tier,omitempty"`
}

// GetAddonService retrieves an AddonService by namespace and name
func (c *Client) GetAddonService(ctx context.Context, namespace, name string) (*AddonService, error) {
	var result AddonService
	path := fmt.Sprintf("/api/web/namespaces/%s/addon_services/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}

// GetAddonServiceDetails retrieves detailed addon service information
// Uses the custom API endpoint for richer information
func (c *Client) GetAddonServiceDetails(ctx context.Context, name string) (*AddonServiceDetails, error) {
	var result AddonServiceDetails
	path := fmt.Sprintf("/api/web/custom/namespaces/shared/addon_services/%s", name)
	err := c.Get(ctx, path, &result)
	return &result, err
}

// GetAddonServiceActivationStatus retrieves the activation status for an addon service
func (c *Client) GetAddonServiceActivationStatus(ctx context.Context, addonService string) (*AddonServiceActivationStatus, error) {
	var result AddonServiceActivationStatus
	path := fmt.Sprintf("/api/web/namespaces/system/addon_services/%s/activation-status", addonService)
	err := c.Get(ctx, path, &result)
	return &result, err
}

// ListAddonServices lists all addon services in a namespace
func (c *Client) ListAddonServices(ctx context.Context, namespace string) ([]AddonService, error) {
	var result struct {
		Items []AddonService `json:"items"`
	}
	path := fmt.Sprintf("/api/web/namespaces/%s/addon_services", namespace)
	err := c.Get(ctx, path, &result)
	return result.Items, err
}
