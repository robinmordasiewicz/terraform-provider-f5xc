// Copyright (c) F5XC Community
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client manages communication with the F5 Distributed Cloud API
type Client struct {
	BaseURL    string
	APIToken   string
	HTTPClient *http.Client
}

// NewClient creates a new F5 Distributed Cloud API client
func NewClient(baseURL, apiToken string) *Client {
	return &Client{
		BaseURL:  baseURL,
		APIToken: apiToken,
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// Metadata represents common metadata fields
type Metadata struct {
	Name        string            `json:"name"`
	Namespace   string            `json:"namespace,omitempty"`
	Labels      map[string]string `json:"labels,omitempty"`
	Annotations map[string]string `json:"annotations,omitempty"`
	UID         string            `json:"uid,omitempty"`
}

// doRequest performs an HTTP request and returns the response
func (c *Client) doRequest(ctx context.Context, method, path string, body interface{}) ([]byte, error) {
	var reqBody io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonBody)
	}

	url := c.BaseURL + path
	req, err := http.NewRequestWithContext(ctx, method, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "APIToken "+c.APIToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("API error (status %d): %s", resp.StatusCode, string(respBody))
	}

	return respBody, nil
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

// ===== NAMESPACE =====
type Namespace struct {
	Metadata Metadata      `json:"metadata"`
	Spec     NamespaceSpec `json:"spec"`
}
type NamespaceSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateNamespace(ctx context.Context, resource *Namespace) (*Namespace, error) {
	var result Namespace
	path := fmt.Sprintf("/api/web/namespaces")
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetNamespace(ctx context.Context, name string) (*Namespace, error) {
	var result Namespace
	path := fmt.Sprintf("/api/web/namespaces/%s", name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateNamespace(ctx context.Context, resource *Namespace) (*Namespace, error) {
	var result Namespace
	path := fmt.Sprintf("/api/web/namespaces/%s", resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteNamespace(ctx context.Context, name string) error {
	path := fmt.Sprintf("/api/web/namespaces/%s", name)
	return c.Delete(ctx, path)
}

// ===== HTTP LOAD BALANCER =====
type HTTPLoadBalancer struct {
	Metadata Metadata             `json:"metadata"`
	Spec     HTTPLoadBalancerSpec `json:"spec"`
}
type HTTPLoadBalancerSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateHTTPLoadBalancer(ctx context.Context, resource *HTTPLoadBalancer) (*HTTPLoadBalancer, error) {
	var result HTTPLoadBalancer
	path := fmt.Sprintf("/api/config/namespaces/%s/http_loadbalancers", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetHTTPLoadBalancer(ctx context.Context, namespace, name string) (*HTTPLoadBalancer, error) {
	var result HTTPLoadBalancer
	path := fmt.Sprintf("/api/config/namespaces/%s/http_loadbalancers/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateHTTPLoadBalancer(ctx context.Context, resource *HTTPLoadBalancer) (*HTTPLoadBalancer, error) {
	var result HTTPLoadBalancer
	path := fmt.Sprintf("/api/config/namespaces/%s/http_loadbalancers/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteHTTPLoadBalancer(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/http_loadbalancers/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== ORIGIN POOL =====
type OriginPool struct {
	Metadata Metadata       `json:"metadata"`
	Spec     OriginPoolSpec `json:"spec"`
}
type OriginPoolSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateOriginPool(ctx context.Context, resource *OriginPool) (*OriginPool, error) {
	var result OriginPool
	path := fmt.Sprintf("/api/config/namespaces/%s/origin_pools", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetOriginPool(ctx context.Context, namespace, name string) (*OriginPool, error) {
	var result OriginPool
	path := fmt.Sprintf("/api/config/namespaces/%s/origin_pools/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateOriginPool(ctx context.Context, resource *OriginPool) (*OriginPool, error) {
	var result OriginPool
	path := fmt.Sprintf("/api/config/namespaces/%s/origin_pools/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteOriginPool(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/origin_pools/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== AddressAllocator =====
type AddressAllocator struct {
	Metadata Metadata         `json:"metadata"`
	Spec     AddressAllocatorSpec `json:"spec"`
}
type AddressAllocatorSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateAddressAllocator(ctx context.Context, resource *AddressAllocator) (*AddressAllocator, error) {
	var result AddressAllocator
	path := fmt.Sprintf("/api/config/namespaces/%s/address_allocators", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetAddressAllocator(ctx context.Context, namespace, name string) (*AddressAllocator, error) {
	var result AddressAllocator
	path := fmt.Sprintf("/api/config/namespaces/%s/address_allocators/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateAddressAllocator(ctx context.Context, resource *AddressAllocator) (*AddressAllocator, error) {
	var result AddressAllocator
	path := fmt.Sprintf("/api/config/namespaces/%s/address_allocators/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteAddressAllocator(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/address_allocators/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== AdvertisePolicy =====
type AdvertisePolicy struct {
	Metadata Metadata         `json:"metadata"`
	Spec     AdvertisePolicySpec `json:"spec"`
}
type AdvertisePolicySpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateAdvertisePolicy(ctx context.Context, resource *AdvertisePolicy) (*AdvertisePolicy, error) {
	var result AdvertisePolicy
	path := fmt.Sprintf("/api/config/namespaces/%s/advertise_policies", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetAdvertisePolicy(ctx context.Context, namespace, name string) (*AdvertisePolicy, error) {
	var result AdvertisePolicy
	path := fmt.Sprintf("/api/config/namespaces/%s/advertise_policies/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateAdvertisePolicy(ctx context.Context, resource *AdvertisePolicy) (*AdvertisePolicy, error) {
	var result AdvertisePolicy
	path := fmt.Sprintf("/api/config/namespaces/%s/advertise_policies/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteAdvertisePolicy(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/advertise_policies/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== AlertPolicy =====
type AlertPolicy struct {
	Metadata Metadata         `json:"metadata"`
	Spec     AlertPolicySpec `json:"spec"`
}
type AlertPolicySpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateAlertPolicy(ctx context.Context, resource *AlertPolicy) (*AlertPolicy, error) {
	var result AlertPolicy
	path := fmt.Sprintf("/api/config/namespaces/%s/alert_policies", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetAlertPolicy(ctx context.Context, namespace, name string) (*AlertPolicy, error) {
	var result AlertPolicy
	path := fmt.Sprintf("/api/config/namespaces/%s/alert_policies/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateAlertPolicy(ctx context.Context, resource *AlertPolicy) (*AlertPolicy, error) {
	var result AlertPolicy
	path := fmt.Sprintf("/api/config/namespaces/%s/alert_policies/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteAlertPolicy(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/alert_policies/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== AlertReceiver =====
type AlertReceiver struct {
	Metadata Metadata         `json:"metadata"`
	Spec     AlertReceiverSpec `json:"spec"`
}
type AlertReceiverSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateAlertReceiver(ctx context.Context, resource *AlertReceiver) (*AlertReceiver, error) {
	var result AlertReceiver
	path := fmt.Sprintf("/api/config/namespaces/%s/alert_receivers", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetAlertReceiver(ctx context.Context, namespace, name string) (*AlertReceiver, error) {
	var result AlertReceiver
	path := fmt.Sprintf("/api/config/namespaces/%s/alert_receivers/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateAlertReceiver(ctx context.Context, resource *AlertReceiver) (*AlertReceiver, error) {
	var result AlertReceiver
	path := fmt.Sprintf("/api/config/namespaces/%s/alert_receivers/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteAlertReceiver(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/alert_receivers/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== APICrawler =====
type APICrawler struct {
	Metadata Metadata         `json:"metadata"`
	Spec     APICrawlerSpec `json:"spec"`
}
type APICrawlerSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateAPICrawler(ctx context.Context, resource *APICrawler) (*APICrawler, error) {
	var result APICrawler
	path := fmt.Sprintf("/api/config/namespaces/%s/api_crawlers", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetAPICrawler(ctx context.Context, namespace, name string) (*APICrawler, error) {
	var result APICrawler
	path := fmt.Sprintf("/api/config/namespaces/%s/api_crawlers/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateAPICrawler(ctx context.Context, resource *APICrawler) (*APICrawler, error) {
	var result APICrawler
	path := fmt.Sprintf("/api/config/namespaces/%s/api_crawlers/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteAPICrawler(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/api_crawlers/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== APIDefinition =====
type APIDefinition struct {
	Metadata Metadata         `json:"metadata"`
	Spec     APIDefinitionSpec `json:"spec"`
}
type APIDefinitionSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateAPIDefinition(ctx context.Context, resource *APIDefinition) (*APIDefinition, error) {
	var result APIDefinition
	path := fmt.Sprintf("/api/config/namespaces/%s/api_definitions", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetAPIDefinition(ctx context.Context, namespace, name string) (*APIDefinition, error) {
	var result APIDefinition
	path := fmt.Sprintf("/api/config/namespaces/%s/api_definitions/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateAPIDefinition(ctx context.Context, resource *APIDefinition) (*APIDefinition, error) {
	var result APIDefinition
	path := fmt.Sprintf("/api/config/namespaces/%s/api_definitions/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteAPIDefinition(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/api_definitions/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== APIDiscovery =====
type APIDiscovery struct {
	Metadata Metadata         `json:"metadata"`
	Spec     APIDiscoverySpec `json:"spec"`
}
type APIDiscoverySpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateAPIDiscovery(ctx context.Context, resource *APIDiscovery) (*APIDiscovery, error) {
	var result APIDiscovery
	path := fmt.Sprintf("/api/config/namespaces/%s/api_discoveries", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetAPIDiscovery(ctx context.Context, namespace, name string) (*APIDiscovery, error) {
	var result APIDiscovery
	path := fmt.Sprintf("/api/config/namespaces/%s/api_discoveries/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateAPIDiscovery(ctx context.Context, resource *APIDiscovery) (*APIDiscovery, error) {
	var result APIDiscovery
	path := fmt.Sprintf("/api/config/namespaces/%s/api_discoveries/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteAPIDiscovery(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/api_discoveries/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== APITesting =====
type APITesting struct {
	Metadata Metadata         `json:"metadata"`
	Spec     APITestingSpec `json:"spec"`
}
type APITestingSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateAPITesting(ctx context.Context, resource *APITesting) (*APITesting, error) {
	var result APITesting
	path := fmt.Sprintf("/api/config/namespaces/%s/api_testings", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetAPITesting(ctx context.Context, namespace, name string) (*APITesting, error) {
	var result APITesting
	path := fmt.Sprintf("/api/config/namespaces/%s/api_testings/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateAPITesting(ctx context.Context, resource *APITesting) (*APITesting, error) {
	var result APITesting
	path := fmt.Sprintf("/api/config/namespaces/%s/api_testings/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteAPITesting(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/api_testings/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== Apm =====
type Apm struct {
	Metadata Metadata         `json:"metadata"`
	Spec     ApmSpec `json:"spec"`
}
type ApmSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateApm(ctx context.Context, resource *Apm) (*Apm, error) {
	var result Apm
	path := fmt.Sprintf("/api/config/namespaces/%s/apms", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetApm(ctx context.Context, namespace, name string) (*Apm, error) {
	var result Apm
	path := fmt.Sprintf("/api/config/namespaces/%s/apms/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateApm(ctx context.Context, resource *Apm) (*Apm, error) {
	var result Apm
	path := fmt.Sprintf("/api/config/namespaces/%s/apms/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteApm(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/apms/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== AppAPIGroup =====
type AppAPIGroup struct {
	Metadata Metadata         `json:"metadata"`
	Spec     AppAPIGroupSpec `json:"spec"`
}
type AppAPIGroupSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateAppAPIGroup(ctx context.Context, resource *AppAPIGroup) (*AppAPIGroup, error) {
	var result AppAPIGroup
	path := fmt.Sprintf("/api/config/namespaces/%s/app_api_groups", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetAppAPIGroup(ctx context.Context, namespace, name string) (*AppAPIGroup, error) {
	var result AppAPIGroup
	path := fmt.Sprintf("/api/config/namespaces/%s/app_api_groups/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateAppAPIGroup(ctx context.Context, resource *AppAPIGroup) (*AppAPIGroup, error) {
	var result AppAPIGroup
	path := fmt.Sprintf("/api/config/namespaces/%s/app_api_groups/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteAppAPIGroup(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/app_api_groups/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== AppFirewall =====
type AppFirewall struct {
	Metadata Metadata         `json:"metadata"`
	Spec     AppFirewallSpec `json:"spec"`
}
type AppFirewallSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateAppFirewall(ctx context.Context, resource *AppFirewall) (*AppFirewall, error) {
	var result AppFirewall
	path := fmt.Sprintf("/api/config/namespaces/%s/app_firewalls", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetAppFirewall(ctx context.Context, namespace, name string) (*AppFirewall, error) {
	var result AppFirewall
	path := fmt.Sprintf("/api/config/namespaces/%s/app_firewalls/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateAppFirewall(ctx context.Context, resource *AppFirewall) (*AppFirewall, error) {
	var result AppFirewall
	path := fmt.Sprintf("/api/config/namespaces/%s/app_firewalls/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteAppFirewall(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/app_firewalls/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== AppSetting =====
type AppSetting struct {
	Metadata Metadata         `json:"metadata"`
	Spec     AppSettingSpec `json:"spec"`
}
type AppSettingSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateAppSetting(ctx context.Context, resource *AppSetting) (*AppSetting, error) {
	var result AppSetting
	path := fmt.Sprintf("/api/config/namespaces/%s/app_settings", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetAppSetting(ctx context.Context, namespace, name string) (*AppSetting, error) {
	var result AppSetting
	path := fmt.Sprintf("/api/config/namespaces/%s/app_settings/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateAppSetting(ctx context.Context, resource *AppSetting) (*AppSetting, error) {
	var result AppSetting
	path := fmt.Sprintf("/api/config/namespaces/%s/app_settings/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteAppSetting(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/app_settings/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== AppType =====
type AppType struct {
	Metadata Metadata         `json:"metadata"`
	Spec     AppTypeSpec `json:"spec"`
}
type AppTypeSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateAppType(ctx context.Context, resource *AppType) (*AppType, error) {
	var result AppType
	path := fmt.Sprintf("/api/config/namespaces/%s/app_types", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetAppType(ctx context.Context, namespace, name string) (*AppType, error) {
	var result AppType
	path := fmt.Sprintf("/api/config/namespaces/%s/app_types/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateAppType(ctx context.Context, resource *AppType) (*AppType, error) {
	var result AppType
	path := fmt.Sprintf("/api/config/namespaces/%s/app_types/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteAppType(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/app_types/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== Authentication =====
type Authentication struct {
	Metadata Metadata         `json:"metadata"`
	Spec     AuthenticationSpec `json:"spec"`
}
type AuthenticationSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateAuthentication(ctx context.Context, resource *Authentication) (*Authentication, error) {
	var result Authentication
	path := fmt.Sprintf("/api/config/namespaces/%s/authentications", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetAuthentication(ctx context.Context, namespace, name string) (*Authentication, error) {
	var result Authentication
	path := fmt.Sprintf("/api/config/namespaces/%s/authentications/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateAuthentication(ctx context.Context, resource *Authentication) (*Authentication, error) {
	var result Authentication
	path := fmt.Sprintf("/api/config/namespaces/%s/authentications/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteAuthentication(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/authentications/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== AWSTGWSite =====
type AWSTGWSite struct {
	Metadata Metadata         `json:"metadata"`
	Spec     AWSTGWSiteSpec `json:"spec"`
}
type AWSTGWSiteSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateAWSTGWSite(ctx context.Context, resource *AWSTGWSite) (*AWSTGWSite, error) {
	var result AWSTGWSite
	path := fmt.Sprintf("/api/config/namespaces/%s/aws_tgw_sites", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetAWSTGWSite(ctx context.Context, namespace, name string) (*AWSTGWSite, error) {
	var result AWSTGWSite
	path := fmt.Sprintf("/api/config/namespaces/%s/aws_tgw_sites/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateAWSTGWSite(ctx context.Context, resource *AWSTGWSite) (*AWSTGWSite, error) {
	var result AWSTGWSite
	path := fmt.Sprintf("/api/config/namespaces/%s/aws_tgw_sites/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteAWSTGWSite(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/aws_tgw_sites/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== AWSVPCSite =====
type AWSVPCSite struct {
	Metadata Metadata         `json:"metadata"`
	Spec     AWSVPCSiteSpec `json:"spec"`
}
type AWSVPCSiteSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateAWSVPCSite(ctx context.Context, resource *AWSVPCSite) (*AWSVPCSite, error) {
	var result AWSVPCSite
	path := fmt.Sprintf("/api/config/namespaces/%s/aws_vpc_sites", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetAWSVPCSite(ctx context.Context, namespace, name string) (*AWSVPCSite, error) {
	var result AWSVPCSite
	path := fmt.Sprintf("/api/config/namespaces/%s/aws_vpc_sites/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateAWSVPCSite(ctx context.Context, resource *AWSVPCSite) (*AWSVPCSite, error) {
	var result AWSVPCSite
	path := fmt.Sprintf("/api/config/namespaces/%s/aws_vpc_sites/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteAWSVPCSite(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/aws_vpc_sites/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== AzureVNETSite =====
type AzureVNETSite struct {
	Metadata Metadata         `json:"metadata"`
	Spec     AzureVNETSiteSpec `json:"spec"`
}
type AzureVNETSiteSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateAzureVNETSite(ctx context.Context, resource *AzureVNETSite) (*AzureVNETSite, error) {
	var result AzureVNETSite
	path := fmt.Sprintf("/api/config/namespaces/%s/azure_vnet_sites", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetAzureVNETSite(ctx context.Context, namespace, name string) (*AzureVNETSite, error) {
	var result AzureVNETSite
	path := fmt.Sprintf("/api/config/namespaces/%s/azure_vnet_sites/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateAzureVNETSite(ctx context.Context, resource *AzureVNETSite) (*AzureVNETSite, error) {
	var result AzureVNETSite
	path := fmt.Sprintf("/api/config/namespaces/%s/azure_vnet_sites/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteAzureVNETSite(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/azure_vnet_sites/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== BgpAsnSet =====
type BgpAsnSet struct {
	Metadata Metadata         `json:"metadata"`
	Spec     BgpAsnSetSpec `json:"spec"`
}
type BgpAsnSetSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateBgpAsnSet(ctx context.Context, resource *BgpAsnSet) (*BgpAsnSet, error) {
	var result BgpAsnSet
	path := fmt.Sprintf("/api/config/namespaces/%s/bgp_asn_sets", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetBgpAsnSet(ctx context.Context, namespace, name string) (*BgpAsnSet, error) {
	var result BgpAsnSet
	path := fmt.Sprintf("/api/config/namespaces/%s/bgp_asn_sets/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateBgpAsnSet(ctx context.Context, resource *BgpAsnSet) (*BgpAsnSet, error) {
	var result BgpAsnSet
	path := fmt.Sprintf("/api/config/namespaces/%s/bgp_asn_sets/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteBgpAsnSet(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/bgp_asn_sets/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== Bgp =====
type Bgp struct {
	Metadata Metadata         `json:"metadata"`
	Spec     BgpSpec `json:"spec"`
}
type BgpSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateBgp(ctx context.Context, resource *Bgp) (*Bgp, error) {
	var result Bgp
	path := fmt.Sprintf("/api/config/namespaces/%s/bgps", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetBgp(ctx context.Context, namespace, name string) (*Bgp, error) {
	var result Bgp
	path := fmt.Sprintf("/api/config/namespaces/%s/bgps/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateBgp(ctx context.Context, resource *Bgp) (*Bgp, error) {
	var result Bgp
	path := fmt.Sprintf("/api/config/namespaces/%s/bgps/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteBgp(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/bgps/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== BgpRoutingPolicy =====
type BgpRoutingPolicy struct {
	Metadata Metadata         `json:"metadata"`
	Spec     BgpRoutingPolicySpec `json:"spec"`
}
type BgpRoutingPolicySpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateBgpRoutingPolicy(ctx context.Context, resource *BgpRoutingPolicy) (*BgpRoutingPolicy, error) {
	var result BgpRoutingPolicy
	path := fmt.Sprintf("/api/config/namespaces/%s/bgp_routing_policies", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetBgpRoutingPolicy(ctx context.Context, namespace, name string) (*BgpRoutingPolicy, error) {
	var result BgpRoutingPolicy
	path := fmt.Sprintf("/api/config/namespaces/%s/bgp_routing_policies/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateBgpRoutingPolicy(ctx context.Context, resource *BgpRoutingPolicy) (*BgpRoutingPolicy, error) {
	var result BgpRoutingPolicy
	path := fmt.Sprintf("/api/config/namespaces/%s/bgp_routing_policies/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteBgpRoutingPolicy(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/bgp_routing_policies/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== BigIPVirtualServer =====
type BigIPVirtualServer struct {
	Metadata Metadata         `json:"metadata"`
	Spec     BigIPVirtualServerSpec `json:"spec"`
}
type BigIPVirtualServerSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateBigIPVirtualServer(ctx context.Context, resource *BigIPVirtualServer) (*BigIPVirtualServer, error) {
	var result BigIPVirtualServer
	path := fmt.Sprintf("/api/config/namespaces/%s/bigip_virtual_servers", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetBigIPVirtualServer(ctx context.Context, namespace, name string) (*BigIPVirtualServer, error) {
	var result BigIPVirtualServer
	path := fmt.Sprintf("/api/config/namespaces/%s/bigip_virtual_servers/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateBigIPVirtualServer(ctx context.Context, resource *BigIPVirtualServer) (*BigIPVirtualServer, error) {
	var result BigIPVirtualServer
	path := fmt.Sprintf("/api/config/namespaces/%s/bigip_virtual_servers/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteBigIPVirtualServer(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/bigip_virtual_servers/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== BotDefenseAppInfrastructure =====
type BotDefenseAppInfrastructure struct {
	Metadata Metadata         `json:"metadata"`
	Spec     BotDefenseAppInfrastructureSpec `json:"spec"`
}
type BotDefenseAppInfrastructureSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateBotDefenseAppInfrastructure(ctx context.Context, resource *BotDefenseAppInfrastructure) (*BotDefenseAppInfrastructure, error) {
	var result BotDefenseAppInfrastructure
	path := fmt.Sprintf("/api/config/namespaces/%s/bot_defense_app_infrastructures", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetBotDefenseAppInfrastructure(ctx context.Context, namespace, name string) (*BotDefenseAppInfrastructure, error) {
	var result BotDefenseAppInfrastructure
	path := fmt.Sprintf("/api/config/namespaces/%s/bot_defense_app_infrastructures/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateBotDefenseAppInfrastructure(ctx context.Context, resource *BotDefenseAppInfrastructure) (*BotDefenseAppInfrastructure, error) {
	var result BotDefenseAppInfrastructure
	path := fmt.Sprintf("/api/config/namespaces/%s/bot_defense_app_infrastructures/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteBotDefenseAppInfrastructure(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/bot_defense_app_infrastructures/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== CdnCacheRule =====
type CdnCacheRule struct {
	Metadata Metadata         `json:"metadata"`
	Spec     CdnCacheRuleSpec `json:"spec"`
}
type CdnCacheRuleSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateCdnCacheRule(ctx context.Context, resource *CdnCacheRule) (*CdnCacheRule, error) {
	var result CdnCacheRule
	path := fmt.Sprintf("/api/config/namespaces/%s/cdn_cache_rules", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetCdnCacheRule(ctx context.Context, namespace, name string) (*CdnCacheRule, error) {
	var result CdnCacheRule
	path := fmt.Sprintf("/api/config/namespaces/%s/cdn_cache_rules/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateCdnCacheRule(ctx context.Context, resource *CdnCacheRule) (*CdnCacheRule, error) {
	var result CdnCacheRule
	path := fmt.Sprintf("/api/config/namespaces/%s/cdn_cache_rules/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteCdnCacheRule(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/cdn_cache_rules/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== CdnLoadBalancer =====
type CdnLoadBalancer struct {
	Metadata Metadata         `json:"metadata"`
	Spec     CdnLoadBalancerSpec `json:"spec"`
}
type CdnLoadBalancerSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateCdnLoadBalancer(ctx context.Context, resource *CdnLoadBalancer) (*CdnLoadBalancer, error) {
	var result CdnLoadBalancer
	path := fmt.Sprintf("/api/config/namespaces/%s/cdn_loadbalancers", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetCdnLoadBalancer(ctx context.Context, namespace, name string) (*CdnLoadBalancer, error) {
	var result CdnLoadBalancer
	path := fmt.Sprintf("/api/config/namespaces/%s/cdn_loadbalancers/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateCdnLoadBalancer(ctx context.Context, resource *CdnLoadBalancer) (*CdnLoadBalancer, error) {
	var result CdnLoadBalancer
	path := fmt.Sprintf("/api/config/namespaces/%s/cdn_loadbalancers/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteCdnLoadBalancer(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/cdn_loadbalancers/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== CertificateChain =====
type CertificateChain struct {
	Metadata Metadata         `json:"metadata"`
	Spec     CertificateChainSpec `json:"spec"`
}
type CertificateChainSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateCertificateChain(ctx context.Context, resource *CertificateChain) (*CertificateChain, error) {
	var result CertificateChain
	path := fmt.Sprintf("/api/config/namespaces/%s/certificate_chains", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetCertificateChain(ctx context.Context, namespace, name string) (*CertificateChain, error) {
	var result CertificateChain
	path := fmt.Sprintf("/api/config/namespaces/%s/certificate_chains/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateCertificateChain(ctx context.Context, resource *CertificateChain) (*CertificateChain, error) {
	var result CertificateChain
	path := fmt.Sprintf("/api/config/namespaces/%s/certificate_chains/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteCertificateChain(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/certificate_chains/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== Certificate =====
type Certificate struct {
	Metadata Metadata         `json:"metadata"`
	Spec     CertificateSpec `json:"spec"`
}
type CertificateSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateCertificate(ctx context.Context, resource *Certificate) (*Certificate, error) {
	var result Certificate
	path := fmt.Sprintf("/api/config/namespaces/%s/certificates", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetCertificate(ctx context.Context, namespace, name string) (*Certificate, error) {
	var result Certificate
	path := fmt.Sprintf("/api/config/namespaces/%s/certificates/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateCertificate(ctx context.Context, resource *Certificate) (*Certificate, error) {
	var result Certificate
	path := fmt.Sprintf("/api/config/namespaces/%s/certificates/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteCertificate(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/certificates/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== CertifiedHardware =====
type CertifiedHardware struct {
	Metadata Metadata         `json:"metadata"`
	Spec     CertifiedHardwareSpec `json:"spec"`
}
type CertifiedHardwareSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateCertifiedHardware(ctx context.Context, resource *CertifiedHardware) (*CertifiedHardware, error) {
	var result CertifiedHardware
	path := fmt.Sprintf("/api/config/namespaces/%s/certified_hardwares", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetCertifiedHardware(ctx context.Context, namespace, name string) (*CertifiedHardware, error) {
	var result CertifiedHardware
	path := fmt.Sprintf("/api/config/namespaces/%s/certified_hardwares/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateCertifiedHardware(ctx context.Context, resource *CertifiedHardware) (*CertifiedHardware, error) {
	var result CertifiedHardware
	path := fmt.Sprintf("/api/config/namespaces/%s/certified_hardwares/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteCertifiedHardware(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/certified_hardwares/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== CloudConnect =====
type CloudConnect struct {
	Metadata Metadata         `json:"metadata"`
	Spec     CloudConnectSpec `json:"spec"`
}
type CloudConnectSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateCloudConnect(ctx context.Context, resource *CloudConnect) (*CloudConnect, error) {
	var result CloudConnect
	path := fmt.Sprintf("/api/config/namespaces/%s/cloud_connects", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetCloudConnect(ctx context.Context, namespace, name string) (*CloudConnect, error) {
	var result CloudConnect
	path := fmt.Sprintf("/api/config/namespaces/%s/cloud_connects/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateCloudConnect(ctx context.Context, resource *CloudConnect) (*CloudConnect, error) {
	var result CloudConnect
	path := fmt.Sprintf("/api/config/namespaces/%s/cloud_connects/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteCloudConnect(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/cloud_connects/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== CloudCredentials =====
type CloudCredentials struct {
	Metadata Metadata         `json:"metadata"`
	Spec     CloudCredentialsSpec `json:"spec"`
}
type CloudCredentialsSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateCloudCredentials(ctx context.Context, resource *CloudCredentials) (*CloudCredentials, error) {
	var result CloudCredentials
	path := fmt.Sprintf("/api/config/namespaces/%s/cloud_credentialses", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetCloudCredentials(ctx context.Context, namespace, name string) (*CloudCredentials, error) {
	var result CloudCredentials
	path := fmt.Sprintf("/api/config/namespaces/%s/cloud_credentialses/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateCloudCredentials(ctx context.Context, resource *CloudCredentials) (*CloudCredentials, error) {
	var result CloudCredentials
	path := fmt.Sprintf("/api/config/namespaces/%s/cloud_credentialses/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteCloudCredentials(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/cloud_credentialses/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== CloudElasticIP =====
type CloudElasticIP struct {
	Metadata Metadata         `json:"metadata"`
	Spec     CloudElasticIPSpec `json:"spec"`
}
type CloudElasticIPSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateCloudElasticIP(ctx context.Context, resource *CloudElasticIP) (*CloudElasticIP, error) {
	var result CloudElasticIP
	path := fmt.Sprintf("/api/config/namespaces/%s/cloud_elastic_ips", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetCloudElasticIP(ctx context.Context, namespace, name string) (*CloudElasticIP, error) {
	var result CloudElasticIP
	path := fmt.Sprintf("/api/config/namespaces/%s/cloud_elastic_ips/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateCloudElasticIP(ctx context.Context, resource *CloudElasticIP) (*CloudElasticIP, error) {
	var result CloudElasticIP
	path := fmt.Sprintf("/api/config/namespaces/%s/cloud_elastic_ips/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteCloudElasticIP(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/cloud_elastic_ips/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== CloudLink =====
type CloudLink struct {
	Metadata Metadata         `json:"metadata"`
	Spec     CloudLinkSpec `json:"spec"`
}
type CloudLinkSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateCloudLink(ctx context.Context, resource *CloudLink) (*CloudLink, error) {
	var result CloudLink
	path := fmt.Sprintf("/api/config/namespaces/%s/cloud_links", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetCloudLink(ctx context.Context, namespace, name string) (*CloudLink, error) {
	var result CloudLink
	path := fmt.Sprintf("/api/config/namespaces/%s/cloud_links/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateCloudLink(ctx context.Context, resource *CloudLink) (*CloudLink, error) {
	var result CloudLink
	path := fmt.Sprintf("/api/config/namespaces/%s/cloud_links/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteCloudLink(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/cloud_links/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== CloudRegion =====
type CloudRegion struct {
	Metadata Metadata         `json:"metadata"`
	Spec     CloudRegionSpec `json:"spec"`
}
type CloudRegionSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateCloudRegion(ctx context.Context, resource *CloudRegion) (*CloudRegion, error) {
	var result CloudRegion
	path := fmt.Sprintf("/api/config/namespaces/%s/cloud_regions", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetCloudRegion(ctx context.Context, namespace, name string) (*CloudRegion, error) {
	var result CloudRegion
	path := fmt.Sprintf("/api/config/namespaces/%s/cloud_regions/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateCloudRegion(ctx context.Context, resource *CloudRegion) (*CloudRegion, error) {
	var result CloudRegion
	path := fmt.Sprintf("/api/config/namespaces/%s/cloud_regions/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteCloudRegion(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/cloud_regions/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== Cluster =====
type Cluster struct {
	Metadata Metadata         `json:"metadata"`
	Spec     ClusterSpec `json:"spec"`
}
type ClusterSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateCluster(ctx context.Context, resource *Cluster) (*Cluster, error) {
	var result Cluster
	path := fmt.Sprintf("/api/config/namespaces/%s/clusters", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetCluster(ctx context.Context, namespace, name string) (*Cluster, error) {
	var result Cluster
	path := fmt.Sprintf("/api/config/namespaces/%s/clusters/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateCluster(ctx context.Context, resource *Cluster) (*Cluster, error) {
	var result Cluster
	path := fmt.Sprintf("/api/config/namespaces/%s/clusters/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteCluster(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/clusters/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== Cminstance =====
type Cminstance struct {
	Metadata Metadata         `json:"metadata"`
	Spec     CminstanceSpec `json:"spec"`
}
type CminstanceSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateCminstance(ctx context.Context, resource *Cminstance) (*Cminstance, error) {
	var result Cminstance
	path := fmt.Sprintf("/api/config/namespaces/%s/cminstances", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetCminstance(ctx context.Context, namespace, name string) (*Cminstance, error) {
	var result Cminstance
	path := fmt.Sprintf("/api/config/namespaces/%s/cminstances/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateCminstance(ctx context.Context, resource *Cminstance) (*Cminstance, error) {
	var result Cminstance
	path := fmt.Sprintf("/api/config/namespaces/%s/cminstances/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteCminstance(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/cminstances/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== CodeBaseIntegration =====
type CodeBaseIntegration struct {
	Metadata Metadata         `json:"metadata"`
	Spec     CodeBaseIntegrationSpec `json:"spec"`
}
type CodeBaseIntegrationSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateCodeBaseIntegration(ctx context.Context, resource *CodeBaseIntegration) (*CodeBaseIntegration, error) {
	var result CodeBaseIntegration
	path := fmt.Sprintf("/api/config/namespaces/%s/code_base_integrations", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetCodeBaseIntegration(ctx context.Context, namespace, name string) (*CodeBaseIntegration, error) {
	var result CodeBaseIntegration
	path := fmt.Sprintf("/api/config/namespaces/%s/code_base_integrations/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateCodeBaseIntegration(ctx context.Context, resource *CodeBaseIntegration) (*CodeBaseIntegration, error) {
	var result CodeBaseIntegration
	path := fmt.Sprintf("/api/config/namespaces/%s/code_base_integrations/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteCodeBaseIntegration(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/code_base_integrations/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== ContainerRegistry =====
type ContainerRegistry struct {
	Metadata Metadata         `json:"metadata"`
	Spec     ContainerRegistrySpec `json:"spec"`
}
type ContainerRegistrySpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateContainerRegistry(ctx context.Context, resource *ContainerRegistry) (*ContainerRegistry, error) {
	var result ContainerRegistry
	path := fmt.Sprintf("/api/config/namespaces/%s/container_registries", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetContainerRegistry(ctx context.Context, namespace, name string) (*ContainerRegistry, error) {
	var result ContainerRegistry
	path := fmt.Sprintf("/api/config/namespaces/%s/container_registries/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateContainerRegistry(ctx context.Context, resource *ContainerRegistry) (*ContainerRegistry, error) {
	var result ContainerRegistry
	path := fmt.Sprintf("/api/config/namespaces/%s/container_registries/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteContainerRegistry(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/container_registries/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== Crl =====
type Crl struct {
	Metadata Metadata         `json:"metadata"`
	Spec     CrlSpec `json:"spec"`
}
type CrlSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateCrl(ctx context.Context, resource *Crl) (*Crl, error) {
	var result Crl
	path := fmt.Sprintf("/api/config/namespaces/%s/crls", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetCrl(ctx context.Context, namespace, name string) (*Crl, error) {
	var result Crl
	path := fmt.Sprintf("/api/config/namespaces/%s/crls/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateCrl(ctx context.Context, resource *Crl) (*Crl, error) {
	var result Crl
	path := fmt.Sprintf("/api/config/namespaces/%s/crls/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteCrl(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/crls/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== DataGroup =====
type DataGroup struct {
	Metadata Metadata         `json:"metadata"`
	Spec     DataGroupSpec `json:"spec"`
}
type DataGroupSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateDataGroup(ctx context.Context, resource *DataGroup) (*DataGroup, error) {
	var result DataGroup
	path := fmt.Sprintf("/api/config/namespaces/%s/data_groups", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetDataGroup(ctx context.Context, namespace, name string) (*DataGroup, error) {
	var result DataGroup
	path := fmt.Sprintf("/api/config/namespaces/%s/data_groups/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateDataGroup(ctx context.Context, resource *DataGroup) (*DataGroup, error) {
	var result DataGroup
	path := fmt.Sprintf("/api/config/namespaces/%s/data_groups/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteDataGroup(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/data_groups/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== DataType =====
type DataType struct {
	Metadata Metadata         `json:"metadata"`
	Spec     DataTypeSpec `json:"spec"`
}
type DataTypeSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateDataType(ctx context.Context, resource *DataType) (*DataType, error) {
	var result DataType
	path := fmt.Sprintf("/api/config/namespaces/%s/data_types", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetDataType(ctx context.Context, namespace, name string) (*DataType, error) {
	var result DataType
	path := fmt.Sprintf("/api/config/namespaces/%s/data_types/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateDataType(ctx context.Context, resource *DataType) (*DataType, error) {
	var result DataType
	path := fmt.Sprintf("/api/config/namespaces/%s/data_types/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteDataType(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/data_types/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== DcClusterGroup =====
type DcClusterGroup struct {
	Metadata Metadata         `json:"metadata"`
	Spec     DcClusterGroupSpec `json:"spec"`
}
type DcClusterGroupSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateDcClusterGroup(ctx context.Context, resource *DcClusterGroup) (*DcClusterGroup, error) {
	var result DcClusterGroup
	path := fmt.Sprintf("/api/config/namespaces/%s/dc_cluster_groups", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetDcClusterGroup(ctx context.Context, namespace, name string) (*DcClusterGroup, error) {
	var result DcClusterGroup
	path := fmt.Sprintf("/api/config/namespaces/%s/dc_cluster_groups/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateDcClusterGroup(ctx context.Context, resource *DcClusterGroup) (*DcClusterGroup, error) {
	var result DcClusterGroup
	path := fmt.Sprintf("/api/config/namespaces/%s/dc_cluster_groups/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteDcClusterGroup(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/dc_cluster_groups/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== Discovery =====
type Discovery struct {
	Metadata Metadata         `json:"metadata"`
	Spec     DiscoverySpec `json:"spec"`
}
type DiscoverySpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateDiscovery(ctx context.Context, resource *Discovery) (*Discovery, error) {
	var result Discovery
	path := fmt.Sprintf("/api/config/namespaces/%s/discoveries", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetDiscovery(ctx context.Context, namespace, name string) (*Discovery, error) {
	var result Discovery
	path := fmt.Sprintf("/api/config/namespaces/%s/discoveries/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateDiscovery(ctx context.Context, resource *Discovery) (*Discovery, error) {
	var result Discovery
	path := fmt.Sprintf("/api/config/namespaces/%s/discoveries/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteDiscovery(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/discoveries/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== DNSComplianceChecks =====
type DNSComplianceChecks struct {
	Metadata Metadata         `json:"metadata"`
	Spec     DNSComplianceChecksSpec `json:"spec"`
}
type DNSComplianceChecksSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateDNSComplianceChecks(ctx context.Context, resource *DNSComplianceChecks) (*DNSComplianceChecks, error) {
	var result DNSComplianceChecks
	path := fmt.Sprintf("/api/config/namespaces/%s/dns_compliance_checkses", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetDNSComplianceChecks(ctx context.Context, namespace, name string) (*DNSComplianceChecks, error) {
	var result DNSComplianceChecks
	path := fmt.Sprintf("/api/config/namespaces/%s/dns_compliance_checkses/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateDNSComplianceChecks(ctx context.Context, resource *DNSComplianceChecks) (*DNSComplianceChecks, error) {
	var result DNSComplianceChecks
	path := fmt.Sprintf("/api/config/namespaces/%s/dns_compliance_checkses/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteDNSComplianceChecks(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/dns_compliance_checkses/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== DNSDomain =====
type DNSDomain struct {
	Metadata Metadata         `json:"metadata"`
	Spec     DNSDomainSpec `json:"spec"`
}
type DNSDomainSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateDNSDomain(ctx context.Context, resource *DNSDomain) (*DNSDomain, error) {
	var result DNSDomain
	path := fmt.Sprintf("/api/config/namespaces/%s/dns_domains", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetDNSDomain(ctx context.Context, namespace, name string) (*DNSDomain, error) {
	var result DNSDomain
	path := fmt.Sprintf("/api/config/namespaces/%s/dns_domains/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateDNSDomain(ctx context.Context, resource *DNSDomain) (*DNSDomain, error) {
	var result DNSDomain
	path := fmt.Sprintf("/api/config/namespaces/%s/dns_domains/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteDNSDomain(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/dns_domains/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== Endpoint =====
type Endpoint struct {
	Metadata Metadata         `json:"metadata"`
	Spec     EndpointSpec `json:"spec"`
}
type EndpointSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateEndpoint(ctx context.Context, resource *Endpoint) (*Endpoint, error) {
	var result Endpoint
	path := fmt.Sprintf("/api/config/namespaces/%s/endpoints", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetEndpoint(ctx context.Context, namespace, name string) (*Endpoint, error) {
	var result Endpoint
	path := fmt.Sprintf("/api/config/namespaces/%s/endpoints/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateEndpoint(ctx context.Context, resource *Endpoint) (*Endpoint, error) {
	var result Endpoint
	path := fmt.Sprintf("/api/config/namespaces/%s/endpoints/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteEndpoint(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/endpoints/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== EnhancedFirewallPolicy =====
type EnhancedFirewallPolicy struct {
	Metadata Metadata         `json:"metadata"`
	Spec     EnhancedFirewallPolicySpec `json:"spec"`
}
type EnhancedFirewallPolicySpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateEnhancedFirewallPolicy(ctx context.Context, resource *EnhancedFirewallPolicy) (*EnhancedFirewallPolicy, error) {
	var result EnhancedFirewallPolicy
	path := fmt.Sprintf("/api/config/namespaces/%s/enhanced_firewall_policies", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetEnhancedFirewallPolicy(ctx context.Context, namespace, name string) (*EnhancedFirewallPolicy, error) {
	var result EnhancedFirewallPolicy
	path := fmt.Sprintf("/api/config/namespaces/%s/enhanced_firewall_policies/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateEnhancedFirewallPolicy(ctx context.Context, resource *EnhancedFirewallPolicy) (*EnhancedFirewallPolicy, error) {
	var result EnhancedFirewallPolicy
	path := fmt.Sprintf("/api/config/namespaces/%s/enhanced_firewall_policies/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteEnhancedFirewallPolicy(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/enhanced_firewall_policies/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== ExternalConnector =====
type ExternalConnector struct {
	Metadata Metadata         `json:"metadata"`
	Spec     ExternalConnectorSpec `json:"spec"`
}
type ExternalConnectorSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateExternalConnector(ctx context.Context, resource *ExternalConnector) (*ExternalConnector, error) {
	var result ExternalConnector
	path := fmt.Sprintf("/api/config/namespaces/%s/external_connectors", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetExternalConnector(ctx context.Context, namespace, name string) (*ExternalConnector, error) {
	var result ExternalConnector
	path := fmt.Sprintf("/api/config/namespaces/%s/external_connectors/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateExternalConnector(ctx context.Context, resource *ExternalConnector) (*ExternalConnector, error) {
	var result ExternalConnector
	path := fmt.Sprintf("/api/config/namespaces/%s/external_connectors/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteExternalConnector(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/external_connectors/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== FastAcl =====
type FastAcl struct {
	Metadata Metadata         `json:"metadata"`
	Spec     FastAclSpec `json:"spec"`
}
type FastAclSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateFastAcl(ctx context.Context, resource *FastAcl) (*FastAcl, error) {
	var result FastAcl
	path := fmt.Sprintf("/api/config/namespaces/%s/fast_acls", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetFastAcl(ctx context.Context, namespace, name string) (*FastAcl, error) {
	var result FastAcl
	path := fmt.Sprintf("/api/config/namespaces/%s/fast_acls/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateFastAcl(ctx context.Context, resource *FastAcl) (*FastAcl, error) {
	var result FastAcl
	path := fmt.Sprintf("/api/config/namespaces/%s/fast_acls/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteFastAcl(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/fast_acls/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== FastAclRule =====
type FastAclRule struct {
	Metadata Metadata         `json:"metadata"`
	Spec     FastAclRuleSpec `json:"spec"`
}
type FastAclRuleSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateFastAclRule(ctx context.Context, resource *FastAclRule) (*FastAclRule, error) {
	var result FastAclRule
	path := fmt.Sprintf("/api/config/namespaces/%s/fast_acl_rules", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetFastAclRule(ctx context.Context, namespace, name string) (*FastAclRule, error) {
	var result FastAclRule
	path := fmt.Sprintf("/api/config/namespaces/%s/fast_acl_rules/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateFastAclRule(ctx context.Context, resource *FastAclRule) (*FastAclRule, error) {
	var result FastAclRule
	path := fmt.Sprintf("/api/config/namespaces/%s/fast_acl_rules/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteFastAclRule(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/fast_acl_rules/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== FilterSet =====
type FilterSet struct {
	Metadata Metadata         `json:"metadata"`
	Spec     FilterSetSpec `json:"spec"`
}
type FilterSetSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateFilterSet(ctx context.Context, resource *FilterSet) (*FilterSet, error) {
	var result FilterSet
	path := fmt.Sprintf("/api/config/namespaces/%s/filter_sets", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetFilterSet(ctx context.Context, namespace, name string) (*FilterSet, error) {
	var result FilterSet
	path := fmt.Sprintf("/api/config/namespaces/%s/filter_sets/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateFilterSet(ctx context.Context, resource *FilterSet) (*FilterSet, error) {
	var result FilterSet
	path := fmt.Sprintf("/api/config/namespaces/%s/filter_sets/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteFilterSet(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/filter_sets/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== Fleet =====
type Fleet struct {
	Metadata Metadata         `json:"metadata"`
	Spec     FleetSpec `json:"spec"`
}
type FleetSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateFleet(ctx context.Context, resource *Fleet) (*Fleet, error) {
	var result Fleet
	path := fmt.Sprintf("/api/config/namespaces/%s/fleets", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetFleet(ctx context.Context, namespace, name string) (*Fleet, error) {
	var result Fleet
	path := fmt.Sprintf("/api/config/namespaces/%s/fleets/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateFleet(ctx context.Context, resource *Fleet) (*Fleet, error) {
	var result Fleet
	path := fmt.Sprintf("/api/config/namespaces/%s/fleets/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteFleet(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/fleets/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== FlowAnomaly =====
type FlowAnomaly struct {
	Metadata Metadata         `json:"metadata"`
	Spec     FlowAnomalySpec `json:"spec"`
}
type FlowAnomalySpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateFlowAnomaly(ctx context.Context, resource *FlowAnomaly) (*FlowAnomaly, error) {
	var result FlowAnomaly
	path := fmt.Sprintf("/api/config/namespaces/%s/flow_anomalies", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetFlowAnomaly(ctx context.Context, namespace, name string) (*FlowAnomaly, error) {
	var result FlowAnomaly
	path := fmt.Sprintf("/api/config/namespaces/%s/flow_anomalies/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateFlowAnomaly(ctx context.Context, resource *FlowAnomaly) (*FlowAnomaly, error) {
	var result FlowAnomaly
	path := fmt.Sprintf("/api/config/namespaces/%s/flow_anomalies/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteFlowAnomaly(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/flow_anomalies/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== Flow =====
type Flow struct {
	Metadata Metadata         `json:"metadata"`
	Spec     FlowSpec `json:"spec"`
}
type FlowSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateFlow(ctx context.Context, resource *Flow) (*Flow, error) {
	var result Flow
	path := fmt.Sprintf("/api/config/namespaces/%s/flows", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetFlow(ctx context.Context, namespace, name string) (*Flow, error) {
	var result Flow
	path := fmt.Sprintf("/api/config/namespaces/%s/flows/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateFlow(ctx context.Context, resource *Flow) (*Flow, error) {
	var result Flow
	path := fmt.Sprintf("/api/config/namespaces/%s/flows/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteFlow(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/flows/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== ForwardProxyPolicy =====
type ForwardProxyPolicy struct {
	Metadata Metadata         `json:"metadata"`
	Spec     ForwardProxyPolicySpec `json:"spec"`
}
type ForwardProxyPolicySpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateForwardProxyPolicy(ctx context.Context, resource *ForwardProxyPolicy) (*ForwardProxyPolicy, error) {
	var result ForwardProxyPolicy
	path := fmt.Sprintf("/api/config/namespaces/%s/forward_proxy_policies", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetForwardProxyPolicy(ctx context.Context, namespace, name string) (*ForwardProxyPolicy, error) {
	var result ForwardProxyPolicy
	path := fmt.Sprintf("/api/config/namespaces/%s/forward_proxy_policies/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateForwardProxyPolicy(ctx context.Context, resource *ForwardProxyPolicy) (*ForwardProxyPolicy, error) {
	var result ForwardProxyPolicy
	path := fmt.Sprintf("/api/config/namespaces/%s/forward_proxy_policies/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteForwardProxyPolicy(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/forward_proxy_policies/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== ForwardingClass =====
type ForwardingClass struct {
	Metadata Metadata         `json:"metadata"`
	Spec     ForwardingClassSpec `json:"spec"`
}
type ForwardingClassSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateForwardingClass(ctx context.Context, resource *ForwardingClass) (*ForwardingClass, error) {
	var result ForwardingClass
	path := fmt.Sprintf("/api/config/namespaces/%s/forwarding_classes", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetForwardingClass(ctx context.Context, namespace, name string) (*ForwardingClass, error) {
	var result ForwardingClass
	path := fmt.Sprintf("/api/config/namespaces/%s/forwarding_classes/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateForwardingClass(ctx context.Context, resource *ForwardingClass) (*ForwardingClass, error) {
	var result ForwardingClass
	path := fmt.Sprintf("/api/config/namespaces/%s/forwarding_classes/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteForwardingClass(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/forwarding_classes/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== GCPVPCSite =====
type GCPVPCSite struct {
	Metadata Metadata         `json:"metadata"`
	Spec     GCPVPCSiteSpec `json:"spec"`
}
type GCPVPCSiteSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateGCPVPCSite(ctx context.Context, resource *GCPVPCSite) (*GCPVPCSite, error) {
	var result GCPVPCSite
	path := fmt.Sprintf("/api/config/namespaces/%s/gcp_vpc_sites", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetGCPVPCSite(ctx context.Context, namespace, name string) (*GCPVPCSite, error) {
	var result GCPVPCSite
	path := fmt.Sprintf("/api/config/namespaces/%s/gcp_vpc_sites/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateGCPVPCSite(ctx context.Context, resource *GCPVPCSite) (*GCPVPCSite, error) {
	var result GCPVPCSite
	path := fmt.Sprintf("/api/config/namespaces/%s/gcp_vpc_sites/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteGCPVPCSite(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/gcp_vpc_sites/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== GlobalLogReceiver =====
type GlobalLogReceiver struct {
	Metadata Metadata         `json:"metadata"`
	Spec     GlobalLogReceiverSpec `json:"spec"`
}
type GlobalLogReceiverSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateGlobalLogReceiver(ctx context.Context, resource *GlobalLogReceiver) (*GlobalLogReceiver, error) {
	var result GlobalLogReceiver
	path := fmt.Sprintf("/api/config/namespaces/%s/global_log_receivers", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetGlobalLogReceiver(ctx context.Context, namespace, name string) (*GlobalLogReceiver, error) {
	var result GlobalLogReceiver
	path := fmt.Sprintf("/api/config/namespaces/%s/global_log_receivers/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateGlobalLogReceiver(ctx context.Context, resource *GlobalLogReceiver) (*GlobalLogReceiver, error) {
	var result GlobalLogReceiver
	path := fmt.Sprintf("/api/config/namespaces/%s/global_log_receivers/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteGlobalLogReceiver(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/global_log_receivers/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== Healthcheck =====
type Healthcheck struct {
	Metadata Metadata         `json:"metadata"`
	Spec     HealthcheckSpec `json:"spec"`
}
type HealthcheckSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateHealthcheck(ctx context.Context, resource *Healthcheck) (*Healthcheck, error) {
	var result Healthcheck
	path := fmt.Sprintf("/api/config/namespaces/%s/healthchecks", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetHealthcheck(ctx context.Context, namespace, name string) (*Healthcheck, error) {
	var result Healthcheck
	path := fmt.Sprintf("/api/config/namespaces/%s/healthchecks/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateHealthcheck(ctx context.Context, resource *Healthcheck) (*Healthcheck, error) {
	var result Healthcheck
	path := fmt.Sprintf("/api/config/namespaces/%s/healthchecks/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteHealthcheck(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/healthchecks/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== Ike1 =====
type Ike1 struct {
	Metadata Metadata         `json:"metadata"`
	Spec     Ike1Spec `json:"spec"`
}
type Ike1Spec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateIke1(ctx context.Context, resource *Ike1) (*Ike1, error) {
	var result Ike1
	path := fmt.Sprintf("/api/config/namespaces/%s/ike1s", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetIke1(ctx context.Context, namespace, name string) (*Ike1, error) {
	var result Ike1
	path := fmt.Sprintf("/api/config/namespaces/%s/ike1s/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateIke1(ctx context.Context, resource *Ike1) (*Ike1, error) {
	var result Ike1
	path := fmt.Sprintf("/api/config/namespaces/%s/ike1s/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteIke1(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/ike1s/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== Ike2 =====
type Ike2 struct {
	Metadata Metadata         `json:"metadata"`
	Spec     Ike2Spec `json:"spec"`
}
type Ike2Spec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateIke2(ctx context.Context, resource *Ike2) (*Ike2, error) {
	var result Ike2
	path := fmt.Sprintf("/api/config/namespaces/%s/ike2s", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetIke2(ctx context.Context, namespace, name string) (*Ike2, error) {
	var result Ike2
	path := fmt.Sprintf("/api/config/namespaces/%s/ike2s/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateIke2(ctx context.Context, resource *Ike2) (*Ike2, error) {
	var result Ike2
	path := fmt.Sprintf("/api/config/namespaces/%s/ike2s/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteIke2(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/ike2s/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== IKEPhase1Profile =====
type IKEPhase1Profile struct {
	Metadata Metadata         `json:"metadata"`
	Spec     IKEPhase1ProfileSpec `json:"spec"`
}
type IKEPhase1ProfileSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateIKEPhase1Profile(ctx context.Context, resource *IKEPhase1Profile) (*IKEPhase1Profile, error) {
	var result IKEPhase1Profile
	path := fmt.Sprintf("/api/config/namespaces/%s/ike_phase1_profiles", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetIKEPhase1Profile(ctx context.Context, namespace, name string) (*IKEPhase1Profile, error) {
	var result IKEPhase1Profile
	path := fmt.Sprintf("/api/config/namespaces/%s/ike_phase1_profiles/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateIKEPhase1Profile(ctx context.Context, resource *IKEPhase1Profile) (*IKEPhase1Profile, error) {
	var result IKEPhase1Profile
	path := fmt.Sprintf("/api/config/namespaces/%s/ike_phase1_profiles/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteIKEPhase1Profile(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/ike_phase1_profiles/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== IKEPhase2Profile =====
type IKEPhase2Profile struct {
	Metadata Metadata         `json:"metadata"`
	Spec     IKEPhase2ProfileSpec `json:"spec"`
}
type IKEPhase2ProfileSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateIKEPhase2Profile(ctx context.Context, resource *IKEPhase2Profile) (*IKEPhase2Profile, error) {
	var result IKEPhase2Profile
	path := fmt.Sprintf("/api/config/namespaces/%s/ike_phase2_profiles", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetIKEPhase2Profile(ctx context.Context, namespace, name string) (*IKEPhase2Profile, error) {
	var result IKEPhase2Profile
	path := fmt.Sprintf("/api/config/namespaces/%s/ike_phase2_profiles/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateIKEPhase2Profile(ctx context.Context, resource *IKEPhase2Profile) (*IKEPhase2Profile, error) {
	var result IKEPhase2Profile
	path := fmt.Sprintf("/api/config/namespaces/%s/ike_phase2_profiles/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteIKEPhase2Profile(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/ike_phase2_profiles/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== ImplicitLabel =====
type ImplicitLabel struct {
	Metadata Metadata         `json:"metadata"`
	Spec     ImplicitLabelSpec `json:"spec"`
}
type ImplicitLabelSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateImplicitLabel(ctx context.Context, resource *ImplicitLabel) (*ImplicitLabel, error) {
	var result ImplicitLabel
	path := fmt.Sprintf("/api/config/namespaces/%s/implicit_labels", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetImplicitLabel(ctx context.Context, namespace, name string) (*ImplicitLabel, error) {
	var result ImplicitLabel
	path := fmt.Sprintf("/api/config/namespaces/%s/implicit_labels/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateImplicitLabel(ctx context.Context, resource *ImplicitLabel) (*ImplicitLabel, error) {
	var result ImplicitLabel
	path := fmt.Sprintf("/api/config/namespaces/%s/implicit_labels/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteImplicitLabel(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/implicit_labels/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== IPPrefixSet =====
type IPPrefixSet struct {
	Metadata Metadata         `json:"metadata"`
	Spec     IPPrefixSetSpec `json:"spec"`
}
type IPPrefixSetSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateIPPrefixSet(ctx context.Context, resource *IPPrefixSet) (*IPPrefixSet, error) {
	var result IPPrefixSet
	path := fmt.Sprintf("/api/config/namespaces/%s/ip_prefix_sets", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetIPPrefixSet(ctx context.Context, namespace, name string) (*IPPrefixSet, error) {
	var result IPPrefixSet
	path := fmt.Sprintf("/api/config/namespaces/%s/ip_prefix_sets/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateIPPrefixSet(ctx context.Context, resource *IPPrefixSet) (*IPPrefixSet, error) {
	var result IPPrefixSet
	path := fmt.Sprintf("/api/config/namespaces/%s/ip_prefix_sets/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteIPPrefixSet(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/ip_prefix_sets/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== Irule =====
type Irule struct {
	Metadata Metadata         `json:"metadata"`
	Spec     IruleSpec `json:"spec"`
}
type IruleSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateIrule(ctx context.Context, resource *Irule) (*Irule, error) {
	var result Irule
	path := fmt.Sprintf("/api/config/namespaces/%s/irules", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetIrule(ctx context.Context, namespace, name string) (*Irule, error) {
	var result Irule
	path := fmt.Sprintf("/api/config/namespaces/%s/irules/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateIrule(ctx context.Context, resource *Irule) (*Irule, error) {
	var result Irule
	path := fmt.Sprintf("/api/config/namespaces/%s/irules/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteIrule(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/irules/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== K8sCluster =====
type K8sCluster struct {
	Metadata Metadata         `json:"metadata"`
	Spec     K8sClusterSpec `json:"spec"`
}
type K8sClusterSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateK8sCluster(ctx context.Context, resource *K8sCluster) (*K8sCluster, error) {
	var result K8sCluster
	path := fmt.Sprintf("/api/config/namespaces/%s/k8s_clusters", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetK8sCluster(ctx context.Context, namespace, name string) (*K8sCluster, error) {
	var result K8sCluster
	path := fmt.Sprintf("/api/config/namespaces/%s/k8s_clusters/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateK8sCluster(ctx context.Context, resource *K8sCluster) (*K8sCluster, error) {
	var result K8sCluster
	path := fmt.Sprintf("/api/config/namespaces/%s/k8s_clusters/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteK8sCluster(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/k8s_clusters/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== K8sClusterRoleBinding =====
type K8sClusterRoleBinding struct {
	Metadata Metadata         `json:"metadata"`
	Spec     K8sClusterRoleBindingSpec `json:"spec"`
}
type K8sClusterRoleBindingSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateK8sClusterRoleBinding(ctx context.Context, resource *K8sClusterRoleBinding) (*K8sClusterRoleBinding, error) {
	var result K8sClusterRoleBinding
	path := fmt.Sprintf("/api/config/namespaces/%s/k8s_cluster_role_bindings", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetK8sClusterRoleBinding(ctx context.Context, namespace, name string) (*K8sClusterRoleBinding, error) {
	var result K8sClusterRoleBinding
	path := fmt.Sprintf("/api/config/namespaces/%s/k8s_cluster_role_bindings/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateK8sClusterRoleBinding(ctx context.Context, resource *K8sClusterRoleBinding) (*K8sClusterRoleBinding, error) {
	var result K8sClusterRoleBinding
	path := fmt.Sprintf("/api/config/namespaces/%s/k8s_cluster_role_bindings/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteK8sClusterRoleBinding(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/k8s_cluster_role_bindings/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== K8sClusterRole =====
type K8sClusterRole struct {
	Metadata Metadata         `json:"metadata"`
	Spec     K8sClusterRoleSpec `json:"spec"`
}
type K8sClusterRoleSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateK8sClusterRole(ctx context.Context, resource *K8sClusterRole) (*K8sClusterRole, error) {
	var result K8sClusterRole
	path := fmt.Sprintf("/api/config/namespaces/%s/k8s_cluster_roles", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetK8sClusterRole(ctx context.Context, namespace, name string) (*K8sClusterRole, error) {
	var result K8sClusterRole
	path := fmt.Sprintf("/api/config/namespaces/%s/k8s_cluster_roles/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateK8sClusterRole(ctx context.Context, resource *K8sClusterRole) (*K8sClusterRole, error) {
	var result K8sClusterRole
	path := fmt.Sprintf("/api/config/namespaces/%s/k8s_cluster_roles/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteK8sClusterRole(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/k8s_cluster_roles/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== K8sPodSecurityAdmission =====
type K8sPodSecurityAdmission struct {
	Metadata Metadata         `json:"metadata"`
	Spec     K8sPodSecurityAdmissionSpec `json:"spec"`
}
type K8sPodSecurityAdmissionSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateK8sPodSecurityAdmission(ctx context.Context, resource *K8sPodSecurityAdmission) (*K8sPodSecurityAdmission, error) {
	var result K8sPodSecurityAdmission
	path := fmt.Sprintf("/api/config/namespaces/%s/k8s_pod_security_admissions", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetK8sPodSecurityAdmission(ctx context.Context, namespace, name string) (*K8sPodSecurityAdmission, error) {
	var result K8sPodSecurityAdmission
	path := fmt.Sprintf("/api/config/namespaces/%s/k8s_pod_security_admissions/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateK8sPodSecurityAdmission(ctx context.Context, resource *K8sPodSecurityAdmission) (*K8sPodSecurityAdmission, error) {
	var result K8sPodSecurityAdmission
	path := fmt.Sprintf("/api/config/namespaces/%s/k8s_pod_security_admissions/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteK8sPodSecurityAdmission(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/k8s_pod_security_admissions/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== K8sPodSecurityPolicy =====
type K8sPodSecurityPolicy struct {
	Metadata Metadata         `json:"metadata"`
	Spec     K8sPodSecurityPolicySpec `json:"spec"`
}
type K8sPodSecurityPolicySpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateK8sPodSecurityPolicy(ctx context.Context, resource *K8sPodSecurityPolicy) (*K8sPodSecurityPolicy, error) {
	var result K8sPodSecurityPolicy
	path := fmt.Sprintf("/api/config/namespaces/%s/k8s_pod_security_policies", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetK8sPodSecurityPolicy(ctx context.Context, namespace, name string) (*K8sPodSecurityPolicy, error) {
	var result K8sPodSecurityPolicy
	path := fmt.Sprintf("/api/config/namespaces/%s/k8s_pod_security_policies/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateK8sPodSecurityPolicy(ctx context.Context, resource *K8sPodSecurityPolicy) (*K8sPodSecurityPolicy, error) {
	var result K8sPodSecurityPolicy
	path := fmt.Sprintf("/api/config/namespaces/%s/k8s_pod_security_policies/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteK8sPodSecurityPolicy(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/k8s_pod_security_policies/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== KnownLabelKey =====
type KnownLabelKey struct {
	Metadata Metadata         `json:"metadata"`
	Spec     KnownLabelKeySpec `json:"spec"`
}
type KnownLabelKeySpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateKnownLabelKey(ctx context.Context, resource *KnownLabelKey) (*KnownLabelKey, error) {
	var result KnownLabelKey
	path := fmt.Sprintf("/api/config/namespaces/%s/known_label_keys", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetKnownLabelKey(ctx context.Context, namespace, name string) (*KnownLabelKey, error) {
	var result KnownLabelKey
	path := fmt.Sprintf("/api/config/namespaces/%s/known_label_keys/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateKnownLabelKey(ctx context.Context, resource *KnownLabelKey) (*KnownLabelKey, error) {
	var result KnownLabelKey
	path := fmt.Sprintf("/api/config/namespaces/%s/known_label_keys/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteKnownLabelKey(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/known_label_keys/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== KnownLabel =====
type KnownLabel struct {
	Metadata Metadata         `json:"metadata"`
	Spec     KnownLabelSpec `json:"spec"`
}
type KnownLabelSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateKnownLabel(ctx context.Context, resource *KnownLabel) (*KnownLabel, error) {
	var result KnownLabel
	path := fmt.Sprintf("/api/config/namespaces/%s/known_labels", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetKnownLabel(ctx context.Context, namespace, name string) (*KnownLabel, error) {
	var result KnownLabel
	path := fmt.Sprintf("/api/config/namespaces/%s/known_labels/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateKnownLabel(ctx context.Context, resource *KnownLabel) (*KnownLabel, error) {
	var result KnownLabel
	path := fmt.Sprintf("/api/config/namespaces/%s/known_labels/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteKnownLabel(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/known_labels/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== LmaRegion =====
type LmaRegion struct {
	Metadata Metadata         `json:"metadata"`
	Spec     LmaRegionSpec `json:"spec"`
}
type LmaRegionSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateLmaRegion(ctx context.Context, resource *LmaRegion) (*LmaRegion, error) {
	var result LmaRegion
	path := fmt.Sprintf("/api/config/namespaces/%s/lma_regions", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetLmaRegion(ctx context.Context, namespace, name string) (*LmaRegion, error) {
	var result LmaRegion
	path := fmt.Sprintf("/api/config/namespaces/%s/lma_regions/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateLmaRegion(ctx context.Context, resource *LmaRegion) (*LmaRegion, error) {
	var result LmaRegion
	path := fmt.Sprintf("/api/config/namespaces/%s/lma_regions/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteLmaRegion(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/lma_regions/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== LogReceiver =====
type LogReceiver struct {
	Metadata Metadata         `json:"metadata"`
	Spec     LogReceiverSpec `json:"spec"`
}
type LogReceiverSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateLogReceiver(ctx context.Context, resource *LogReceiver) (*LogReceiver, error) {
	var result LogReceiver
	path := fmt.Sprintf("/api/config/namespaces/%s/log_receivers", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetLogReceiver(ctx context.Context, namespace, name string) (*LogReceiver, error) {
	var result LogReceiver
	path := fmt.Sprintf("/api/config/namespaces/%s/log_receivers/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateLogReceiver(ctx context.Context, resource *LogReceiver) (*LogReceiver, error) {
	var result LogReceiver
	path := fmt.Sprintf("/api/config/namespaces/%s/log_receivers/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteLogReceiver(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/log_receivers/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== MaliciousUserMitigation =====
type MaliciousUserMitigation struct {
	Metadata Metadata         `json:"metadata"`
	Spec     MaliciousUserMitigationSpec `json:"spec"`
}
type MaliciousUserMitigationSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateMaliciousUserMitigation(ctx context.Context, resource *MaliciousUserMitigation) (*MaliciousUserMitigation, error) {
	var result MaliciousUserMitigation
	path := fmt.Sprintf("/api/config/namespaces/%s/malicious_user_mitigations", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetMaliciousUserMitigation(ctx context.Context, namespace, name string) (*MaliciousUserMitigation, error) {
	var result MaliciousUserMitigation
	path := fmt.Sprintf("/api/config/namespaces/%s/malicious_user_mitigations/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateMaliciousUserMitigation(ctx context.Context, resource *MaliciousUserMitigation) (*MaliciousUserMitigation, error) {
	var result MaliciousUserMitigation
	path := fmt.Sprintf("/api/config/namespaces/%s/malicious_user_mitigations/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteMaliciousUserMitigation(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/malicious_user_mitigations/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== ModuleManagement =====
type ModuleManagement struct {
	Metadata Metadata         `json:"metadata"`
	Spec     ModuleManagementSpec `json:"spec"`
}
type ModuleManagementSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateModuleManagement(ctx context.Context, resource *ModuleManagement) (*ModuleManagement, error) {
	var result ModuleManagement
	path := fmt.Sprintf("/api/config/namespaces/%s/module_managements", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetModuleManagement(ctx context.Context, namespace, name string) (*ModuleManagement, error) {
	var result ModuleManagement
	path := fmt.Sprintf("/api/config/namespaces/%s/module_managements/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateModuleManagement(ctx context.Context, resource *ModuleManagement) (*ModuleManagement, error) {
	var result ModuleManagement
	path := fmt.Sprintf("/api/config/namespaces/%s/module_managements/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteModuleManagement(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/module_managements/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== NatPolicy =====
type NatPolicy struct {
	Metadata Metadata         `json:"metadata"`
	Spec     NatPolicySpec `json:"spec"`
}
type NatPolicySpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateNatPolicy(ctx context.Context, resource *NatPolicy) (*NatPolicy, error) {
	var result NatPolicy
	path := fmt.Sprintf("/api/config/namespaces/%s/nat_policies", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetNatPolicy(ctx context.Context, namespace, name string) (*NatPolicy, error) {
	var result NatPolicy
	path := fmt.Sprintf("/api/config/namespaces/%s/nat_policies/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateNatPolicy(ctx context.Context, resource *NatPolicy) (*NatPolicy, error) {
	var result NatPolicy
	path := fmt.Sprintf("/api/config/namespaces/%s/nat_policies/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteNatPolicy(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/nat_policies/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== NetworkConnector =====
type NetworkConnector struct {
	Metadata Metadata         `json:"metadata"`
	Spec     NetworkConnectorSpec `json:"spec"`
}
type NetworkConnectorSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateNetworkConnector(ctx context.Context, resource *NetworkConnector) (*NetworkConnector, error) {
	var result NetworkConnector
	path := fmt.Sprintf("/api/config/namespaces/%s/network_connectors", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetNetworkConnector(ctx context.Context, namespace, name string) (*NetworkConnector, error) {
	var result NetworkConnector
	path := fmt.Sprintf("/api/config/namespaces/%s/network_connectors/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateNetworkConnector(ctx context.Context, resource *NetworkConnector) (*NetworkConnector, error) {
	var result NetworkConnector
	path := fmt.Sprintf("/api/config/namespaces/%s/network_connectors/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteNetworkConnector(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/network_connectors/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== NetworkFirewall =====
type NetworkFirewall struct {
	Metadata Metadata         `json:"metadata"`
	Spec     NetworkFirewallSpec `json:"spec"`
}
type NetworkFirewallSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateNetworkFirewall(ctx context.Context, resource *NetworkFirewall) (*NetworkFirewall, error) {
	var result NetworkFirewall
	path := fmt.Sprintf("/api/config/namespaces/%s/network_firewalls", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetNetworkFirewall(ctx context.Context, namespace, name string) (*NetworkFirewall, error) {
	var result NetworkFirewall
	path := fmt.Sprintf("/api/config/namespaces/%s/network_firewalls/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateNetworkFirewall(ctx context.Context, resource *NetworkFirewall) (*NetworkFirewall, error) {
	var result NetworkFirewall
	path := fmt.Sprintf("/api/config/namespaces/%s/network_firewalls/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteNetworkFirewall(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/network_firewalls/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== NetworkInterface =====
type NetworkInterface struct {
	Metadata Metadata         `json:"metadata"`
	Spec     NetworkInterfaceSpec `json:"spec"`
}
type NetworkInterfaceSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateNetworkInterface(ctx context.Context, resource *NetworkInterface) (*NetworkInterface, error) {
	var result NetworkInterface
	path := fmt.Sprintf("/api/config/namespaces/%s/network_interfaces", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetNetworkInterface(ctx context.Context, namespace, name string) (*NetworkInterface, error) {
	var result NetworkInterface
	path := fmt.Sprintf("/api/config/namespaces/%s/network_interfaces/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateNetworkInterface(ctx context.Context, resource *NetworkInterface) (*NetworkInterface, error) {
	var result NetworkInterface
	path := fmt.Sprintf("/api/config/namespaces/%s/network_interfaces/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteNetworkInterface(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/network_interfaces/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== NetworkPolicy =====
type NetworkPolicy struct {
	Metadata Metadata         `json:"metadata"`
	Spec     NetworkPolicySpec `json:"spec"`
}
type NetworkPolicySpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateNetworkPolicy(ctx context.Context, resource *NetworkPolicy) (*NetworkPolicy, error) {
	var result NetworkPolicy
	path := fmt.Sprintf("/api/config/namespaces/%s/network_policies", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetNetworkPolicy(ctx context.Context, namespace, name string) (*NetworkPolicy, error) {
	var result NetworkPolicy
	path := fmt.Sprintf("/api/config/namespaces/%s/network_policies/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateNetworkPolicy(ctx context.Context, resource *NetworkPolicy) (*NetworkPolicy, error) {
	var result NetworkPolicy
	path := fmt.Sprintf("/api/config/namespaces/%s/network_policies/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteNetworkPolicy(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/network_policies/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== NetworkPolicyRule =====
type NetworkPolicyRule struct {
	Metadata Metadata         `json:"metadata"`
	Spec     NetworkPolicyRuleSpec `json:"spec"`
}
type NetworkPolicyRuleSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateNetworkPolicyRule(ctx context.Context, resource *NetworkPolicyRule) (*NetworkPolicyRule, error) {
	var result NetworkPolicyRule
	path := fmt.Sprintf("/api/config/namespaces/%s/network_policy_rules", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetNetworkPolicyRule(ctx context.Context, namespace, name string) (*NetworkPolicyRule, error) {
	var result NetworkPolicyRule
	path := fmt.Sprintf("/api/config/namespaces/%s/network_policy_rules/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateNetworkPolicyRule(ctx context.Context, resource *NetworkPolicyRule) (*NetworkPolicyRule, error) {
	var result NetworkPolicyRule
	path := fmt.Sprintf("/api/config/namespaces/%s/network_policy_rules/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteNetworkPolicyRule(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/network_policy_rules/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== NetworkPolicySet =====
type NetworkPolicySet struct {
	Metadata Metadata         `json:"metadata"`
	Spec     NetworkPolicySetSpec `json:"spec"`
}
type NetworkPolicySetSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateNetworkPolicySet(ctx context.Context, resource *NetworkPolicySet) (*NetworkPolicySet, error) {
	var result NetworkPolicySet
	path := fmt.Sprintf("/api/config/namespaces/%s/network_policy_sets", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetNetworkPolicySet(ctx context.Context, namespace, name string) (*NetworkPolicySet, error) {
	var result NetworkPolicySet
	path := fmt.Sprintf("/api/config/namespaces/%s/network_policy_sets/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateNetworkPolicySet(ctx context.Context, resource *NetworkPolicySet) (*NetworkPolicySet, error) {
	var result NetworkPolicySet
	path := fmt.Sprintf("/api/config/namespaces/%s/network_policy_sets/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteNetworkPolicySet(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/network_policy_sets/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== NetworkPolicyView =====
type NetworkPolicyView struct {
	Metadata Metadata         `json:"metadata"`
	Spec     NetworkPolicyViewSpec `json:"spec"`
}
type NetworkPolicyViewSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateNetworkPolicyView(ctx context.Context, resource *NetworkPolicyView) (*NetworkPolicyView, error) {
	var result NetworkPolicyView
	path := fmt.Sprintf("/api/config/namespaces/%s/network_policy_views", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetNetworkPolicyView(ctx context.Context, namespace, name string) (*NetworkPolicyView, error) {
	var result NetworkPolicyView
	path := fmt.Sprintf("/api/config/namespaces/%s/network_policy_views/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateNetworkPolicyView(ctx context.Context, resource *NetworkPolicyView) (*NetworkPolicyView, error) {
	var result NetworkPolicyView
	path := fmt.Sprintf("/api/config/namespaces/%s/network_policy_views/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteNetworkPolicyView(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/network_policy_views/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== NfvService =====
type NfvService struct {
	Metadata Metadata         `json:"metadata"`
	Spec     NfvServiceSpec `json:"spec"`
}
type NfvServiceSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateNfvService(ctx context.Context, resource *NfvService) (*NfvService, error) {
	var result NfvService
	path := fmt.Sprintf("/api/config/namespaces/%s/nfv_services", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetNfvService(ctx context.Context, namespace, name string) (*NfvService, error) {
	var result NfvService
	path := fmt.Sprintf("/api/config/namespaces/%s/nfv_services/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateNfvService(ctx context.Context, resource *NfvService) (*NfvService, error) {
	var result NfvService
	path := fmt.Sprintf("/api/config/namespaces/%s/nfv_services/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteNfvService(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/nfv_services/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== Policer =====
type Policer struct {
	Metadata Metadata         `json:"metadata"`
	Spec     PolicerSpec `json:"spec"`
}
type PolicerSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreatePolicer(ctx context.Context, resource *Policer) (*Policer, error) {
	var result Policer
	path := fmt.Sprintf("/api/config/namespaces/%s/policers", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetPolicer(ctx context.Context, namespace, name string) (*Policer, error) {
	var result Policer
	path := fmt.Sprintf("/api/config/namespaces/%s/policers/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdatePolicer(ctx context.Context, resource *Policer) (*Policer, error) {
	var result Policer
	path := fmt.Sprintf("/api/config/namespaces/%s/policers/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeletePolicer(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/policers/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== PolicyBasedRouting =====
type PolicyBasedRouting struct {
	Metadata Metadata         `json:"metadata"`
	Spec     PolicyBasedRoutingSpec `json:"spec"`
}
type PolicyBasedRoutingSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreatePolicyBasedRouting(ctx context.Context, resource *PolicyBasedRouting) (*PolicyBasedRouting, error) {
	var result PolicyBasedRouting
	path := fmt.Sprintf("/api/config/namespaces/%s/policy_based_routings", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetPolicyBasedRouting(ctx context.Context, namespace, name string) (*PolicyBasedRouting, error) {
	var result PolicyBasedRouting
	path := fmt.Sprintf("/api/config/namespaces/%s/policy_based_routings/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdatePolicyBasedRouting(ctx context.Context, resource *PolicyBasedRouting) (*PolicyBasedRouting, error) {
	var result PolicyBasedRouting
	path := fmt.Sprintf("/api/config/namespaces/%s/policy_based_routings/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeletePolicyBasedRouting(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/policy_based_routings/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== ProtocolInspection =====
type ProtocolInspection struct {
	Metadata Metadata         `json:"metadata"`
	Spec     ProtocolInspectionSpec `json:"spec"`
}
type ProtocolInspectionSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateProtocolInspection(ctx context.Context, resource *ProtocolInspection) (*ProtocolInspection, error) {
	var result ProtocolInspection
	path := fmt.Sprintf("/api/config/namespaces/%s/protocol_inspections", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetProtocolInspection(ctx context.Context, namespace, name string) (*ProtocolInspection, error) {
	var result ProtocolInspection
	path := fmt.Sprintf("/api/config/namespaces/%s/protocol_inspections/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateProtocolInspection(ctx context.Context, resource *ProtocolInspection) (*ProtocolInspection, error) {
	var result ProtocolInspection
	path := fmt.Sprintf("/api/config/namespaces/%s/protocol_inspections/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteProtocolInspection(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/protocol_inspections/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== ProtocolPolicer =====
type ProtocolPolicer struct {
	Metadata Metadata         `json:"metadata"`
	Spec     ProtocolPolicerSpec `json:"spec"`
}
type ProtocolPolicerSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateProtocolPolicer(ctx context.Context, resource *ProtocolPolicer) (*ProtocolPolicer, error) {
	var result ProtocolPolicer
	path := fmt.Sprintf("/api/config/namespaces/%s/protocol_policers", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetProtocolPolicer(ctx context.Context, namespace, name string) (*ProtocolPolicer, error) {
	var result ProtocolPolicer
	path := fmt.Sprintf("/api/config/namespaces/%s/protocol_policers/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateProtocolPolicer(ctx context.Context, resource *ProtocolPolicer) (*ProtocolPolicer, error) {
	var result ProtocolPolicer
	path := fmt.Sprintf("/api/config/namespaces/%s/protocol_policers/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteProtocolPolicer(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/protocol_policers/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== Proxy =====
type Proxy struct {
	Metadata Metadata         `json:"metadata"`
	Spec     ProxySpec `json:"spec"`
}
type ProxySpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateProxy(ctx context.Context, resource *Proxy) (*Proxy, error) {
	var result Proxy
	path := fmt.Sprintf("/api/config/namespaces/%s/proxies", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetProxy(ctx context.Context, namespace, name string) (*Proxy, error) {
	var result Proxy
	path := fmt.Sprintf("/api/config/namespaces/%s/proxies/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateProxy(ctx context.Context, resource *Proxy) (*Proxy, error) {
	var result Proxy
	path := fmt.Sprintf("/api/config/namespaces/%s/proxies/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteProxy(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/proxies/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== PublicIP =====
type PublicIP struct {
	Metadata Metadata         `json:"metadata"`
	Spec     PublicIPSpec `json:"spec"`
}
type PublicIPSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreatePublicIP(ctx context.Context, resource *PublicIP) (*PublicIP, error) {
	var result PublicIP
	path := fmt.Sprintf("/api/config/namespaces/%s/public_ips", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetPublicIP(ctx context.Context, namespace, name string) (*PublicIP, error) {
	var result PublicIP
	path := fmt.Sprintf("/api/config/namespaces/%s/public_ips/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdatePublicIP(ctx context.Context, resource *PublicIP) (*PublicIP, error) {
	var result PublicIP
	path := fmt.Sprintf("/api/config/namespaces/%s/public_ips/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeletePublicIP(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/public_ips/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== RateLimiterPolicy =====
type RateLimiterPolicy struct {
	Metadata Metadata         `json:"metadata"`
	Spec     RateLimiterPolicySpec `json:"spec"`
}
type RateLimiterPolicySpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateRateLimiterPolicy(ctx context.Context, resource *RateLimiterPolicy) (*RateLimiterPolicy, error) {
	var result RateLimiterPolicy
	path := fmt.Sprintf("/api/config/namespaces/%s/rate_limiter_policies", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetRateLimiterPolicy(ctx context.Context, namespace, name string) (*RateLimiterPolicy, error) {
	var result RateLimiterPolicy
	path := fmt.Sprintf("/api/config/namespaces/%s/rate_limiter_policies/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateRateLimiterPolicy(ctx context.Context, resource *RateLimiterPolicy) (*RateLimiterPolicy, error) {
	var result RateLimiterPolicy
	path := fmt.Sprintf("/api/config/namespaces/%s/rate_limiter_policies/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteRateLimiterPolicy(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/rate_limiter_policies/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== RateLimiter =====
type RateLimiter struct {
	Metadata Metadata         `json:"metadata"`
	Spec     RateLimiterSpec `json:"spec"`
}
type RateLimiterSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateRateLimiter(ctx context.Context, resource *RateLimiter) (*RateLimiter, error) {
	var result RateLimiter
	path := fmt.Sprintf("/api/config/namespaces/%s/rate_limiters", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetRateLimiter(ctx context.Context, namespace, name string) (*RateLimiter, error) {
	var result RateLimiter
	path := fmt.Sprintf("/api/config/namespaces/%s/rate_limiters/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateRateLimiter(ctx context.Context, resource *RateLimiter) (*RateLimiter, error) {
	var result RateLimiter
	path := fmt.Sprintf("/api/config/namespaces/%s/rate_limiters/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteRateLimiter(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/rate_limiters/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== Route =====
type Route struct {
	Metadata Metadata         `json:"metadata"`
	Spec     RouteSpec `json:"spec"`
}
type RouteSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateRoute(ctx context.Context, resource *Route) (*Route, error) {
	var result Route
	path := fmt.Sprintf("/api/config/namespaces/%s/routes", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetRoute(ctx context.Context, namespace, name string) (*Route, error) {
	var result Route
	path := fmt.Sprintf("/api/config/namespaces/%s/routes/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateRoute(ctx context.Context, resource *Route) (*Route, error) {
	var result Route
	path := fmt.Sprintf("/api/config/namespaces/%s/routes/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteRoute(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/routes/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== RuleSuggestion =====
type RuleSuggestion struct {
	Metadata Metadata         `json:"metadata"`
	Spec     RuleSuggestionSpec `json:"spec"`
}
type RuleSuggestionSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateRuleSuggestion(ctx context.Context, resource *RuleSuggestion) (*RuleSuggestion, error) {
	var result RuleSuggestion
	path := fmt.Sprintf("/api/config/namespaces/%s/rule_suggestions", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetRuleSuggestion(ctx context.Context, namespace, name string) (*RuleSuggestion, error) {
	var result RuleSuggestion
	path := fmt.Sprintf("/api/config/namespaces/%s/rule_suggestions/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateRuleSuggestion(ctx context.Context, resource *RuleSuggestion) (*RuleSuggestion, error) {
	var result RuleSuggestion
	path := fmt.Sprintf("/api/config/namespaces/%s/rule_suggestions/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteRuleSuggestion(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/rule_suggestions/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== SecretManagementAccess =====
type SecretManagementAccess struct {
	Metadata Metadata         `json:"metadata"`
	Spec     SecretManagementAccessSpec `json:"spec"`
}
type SecretManagementAccessSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateSecretManagementAccess(ctx context.Context, resource *SecretManagementAccess) (*SecretManagementAccess, error) {
	var result SecretManagementAccess
	path := fmt.Sprintf("/api/config/namespaces/%s/secret_management_accesses", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetSecretManagementAccess(ctx context.Context, namespace, name string) (*SecretManagementAccess, error) {
	var result SecretManagementAccess
	path := fmt.Sprintf("/api/config/namespaces/%s/secret_management_accesses/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateSecretManagementAccess(ctx context.Context, resource *SecretManagementAccess) (*SecretManagementAccess, error) {
	var result SecretManagementAccess
	path := fmt.Sprintf("/api/config/namespaces/%s/secret_management_accesses/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteSecretManagementAccess(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/secret_management_accesses/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== SecuremeshSite =====
type SecuremeshSite struct {
	Metadata Metadata         `json:"metadata"`
	Spec     SecuremeshSiteSpec `json:"spec"`
}
type SecuremeshSiteSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateSecuremeshSite(ctx context.Context, resource *SecuremeshSite) (*SecuremeshSite, error) {
	var result SecuremeshSite
	path := fmt.Sprintf("/api/config/namespaces/%s/securemesh_sites", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetSecuremeshSite(ctx context.Context, namespace, name string) (*SecuremeshSite, error) {
	var result SecuremeshSite
	path := fmt.Sprintf("/api/config/namespaces/%s/securemesh_sites/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateSecuremeshSite(ctx context.Context, resource *SecuremeshSite) (*SecuremeshSite, error) {
	var result SecuremeshSite
	path := fmt.Sprintf("/api/config/namespaces/%s/securemesh_sites/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteSecuremeshSite(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/securemesh_sites/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== SecuremeshSiteV2 =====
type SecuremeshSiteV2 struct {
	Metadata Metadata         `json:"metadata"`
	Spec     SecuremeshSiteV2Spec `json:"spec"`
}
type SecuremeshSiteV2Spec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateSecuremeshSiteV2(ctx context.Context, resource *SecuremeshSiteV2) (*SecuremeshSiteV2, error) {
	var result SecuremeshSiteV2
	path := fmt.Sprintf("/api/config/namespaces/%s/securemesh_site_v2s", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetSecuremeshSiteV2(ctx context.Context, namespace, name string) (*SecuremeshSiteV2, error) {
	var result SecuremeshSiteV2
	path := fmt.Sprintf("/api/config/namespaces/%s/securemesh_site_v2s/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateSecuremeshSiteV2(ctx context.Context, resource *SecuremeshSiteV2) (*SecuremeshSiteV2, error) {
	var result SecuremeshSiteV2
	path := fmt.Sprintf("/api/config/namespaces/%s/securemesh_site_v2s/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteSecuremeshSiteV2(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/securemesh_site_v2s/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== SegmentConnection =====
type SegmentConnection struct {
	Metadata Metadata         `json:"metadata"`
	Spec     SegmentConnectionSpec `json:"spec"`
}
type SegmentConnectionSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateSegmentConnection(ctx context.Context, resource *SegmentConnection) (*SegmentConnection, error) {
	var result SegmentConnection
	path := fmt.Sprintf("/api/config/namespaces/%s/segment_connections", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetSegmentConnection(ctx context.Context, namespace, name string) (*SegmentConnection, error) {
	var result SegmentConnection
	path := fmt.Sprintf("/api/config/namespaces/%s/segment_connections/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateSegmentConnection(ctx context.Context, resource *SegmentConnection) (*SegmentConnection, error) {
	var result SegmentConnection
	path := fmt.Sprintf("/api/config/namespaces/%s/segment_connections/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteSegmentConnection(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/segment_connections/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== Segment =====
type Segment struct {
	Metadata Metadata         `json:"metadata"`
	Spec     SegmentSpec `json:"spec"`
}
type SegmentSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateSegment(ctx context.Context, resource *Segment) (*Segment, error) {
	var result Segment
	path := fmt.Sprintf("/api/config/namespaces/%s/segments", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetSegment(ctx context.Context, namespace, name string) (*Segment, error) {
	var result Segment
	path := fmt.Sprintf("/api/config/namespaces/%s/segments/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateSegment(ctx context.Context, resource *Segment) (*Segment, error) {
	var result Segment
	path := fmt.Sprintf("/api/config/namespaces/%s/segments/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteSegment(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/segments/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== SensitiveDataPolicy =====
type SensitiveDataPolicy struct {
	Metadata Metadata         `json:"metadata"`
	Spec     SensitiveDataPolicySpec `json:"spec"`
}
type SensitiveDataPolicySpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateSensitiveDataPolicy(ctx context.Context, resource *SensitiveDataPolicy) (*SensitiveDataPolicy, error) {
	var result SensitiveDataPolicy
	path := fmt.Sprintf("/api/config/namespaces/%s/sensitive_data_policies", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetSensitiveDataPolicy(ctx context.Context, namespace, name string) (*SensitiveDataPolicy, error) {
	var result SensitiveDataPolicy
	path := fmt.Sprintf("/api/config/namespaces/%s/sensitive_data_policies/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateSensitiveDataPolicy(ctx context.Context, resource *SensitiveDataPolicy) (*SensitiveDataPolicy, error) {
	var result SensitiveDataPolicy
	path := fmt.Sprintf("/api/config/namespaces/%s/sensitive_data_policies/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteSensitiveDataPolicy(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/sensitive_data_policies/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== ServicePolicy =====
type ServicePolicy struct {
	Metadata Metadata         `json:"metadata"`
	Spec     ServicePolicySpec `json:"spec"`
}
type ServicePolicySpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateServicePolicy(ctx context.Context, resource *ServicePolicy) (*ServicePolicy, error) {
	var result ServicePolicy
	path := fmt.Sprintf("/api/config/namespaces/%s/service_policies", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetServicePolicy(ctx context.Context, namespace, name string) (*ServicePolicy, error) {
	var result ServicePolicy
	path := fmt.Sprintf("/api/config/namespaces/%s/service_policies/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateServicePolicy(ctx context.Context, resource *ServicePolicy) (*ServicePolicy, error) {
	var result ServicePolicy
	path := fmt.Sprintf("/api/config/namespaces/%s/service_policies/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteServicePolicy(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/service_policies/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== ServicePolicyRule =====
type ServicePolicyRule struct {
	Metadata Metadata         `json:"metadata"`
	Spec     ServicePolicyRuleSpec `json:"spec"`
}
type ServicePolicyRuleSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateServicePolicyRule(ctx context.Context, resource *ServicePolicyRule) (*ServicePolicyRule, error) {
	var result ServicePolicyRule
	path := fmt.Sprintf("/api/config/namespaces/%s/service_policy_rules", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetServicePolicyRule(ctx context.Context, namespace, name string) (*ServicePolicyRule, error) {
	var result ServicePolicyRule
	path := fmt.Sprintf("/api/config/namespaces/%s/service_policy_rules/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateServicePolicyRule(ctx context.Context, resource *ServicePolicyRule) (*ServicePolicyRule, error) {
	var result ServicePolicyRule
	path := fmt.Sprintf("/api/config/namespaces/%s/service_policy_rules/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteServicePolicyRule(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/service_policy_rules/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== ServicePolicySet =====
type ServicePolicySet struct {
	Metadata Metadata         `json:"metadata"`
	Spec     ServicePolicySetSpec `json:"spec"`
}
type ServicePolicySetSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateServicePolicySet(ctx context.Context, resource *ServicePolicySet) (*ServicePolicySet, error) {
	var result ServicePolicySet
	path := fmt.Sprintf("/api/config/namespaces/%s/service_policy_sets", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetServicePolicySet(ctx context.Context, namespace, name string) (*ServicePolicySet, error) {
	var result ServicePolicySet
	path := fmt.Sprintf("/api/config/namespaces/%s/service_policy_sets/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateServicePolicySet(ctx context.Context, resource *ServicePolicySet) (*ServicePolicySet, error) {
	var result ServicePolicySet
	path := fmt.Sprintf("/api/config/namespaces/%s/service_policy_sets/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteServicePolicySet(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/service_policy_sets/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== ShapeBotDefenseInstance =====
type ShapeBotDefenseInstance struct {
	Metadata Metadata         `json:"metadata"`
	Spec     ShapeBotDefenseInstanceSpec `json:"spec"`
}
type ShapeBotDefenseInstanceSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateShapeBotDefenseInstance(ctx context.Context, resource *ShapeBotDefenseInstance) (*ShapeBotDefenseInstance, error) {
	var result ShapeBotDefenseInstance
	path := fmt.Sprintf("/api/config/namespaces/%s/shape_bot_defense_instances", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetShapeBotDefenseInstance(ctx context.Context, namespace, name string) (*ShapeBotDefenseInstance, error) {
	var result ShapeBotDefenseInstance
	path := fmt.Sprintf("/api/config/namespaces/%s/shape_bot_defense_instances/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateShapeBotDefenseInstance(ctx context.Context, resource *ShapeBotDefenseInstance) (*ShapeBotDefenseInstance, error) {
	var result ShapeBotDefenseInstance
	path := fmt.Sprintf("/api/config/namespaces/%s/shape_bot_defense_instances/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteShapeBotDefenseInstance(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/shape_bot_defense_instances/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== SiteMeshGroup =====
type SiteMeshGroup struct {
	Metadata Metadata         `json:"metadata"`
	Spec     SiteMeshGroupSpec `json:"spec"`
}
type SiteMeshGroupSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateSiteMeshGroup(ctx context.Context, resource *SiteMeshGroup) (*SiteMeshGroup, error) {
	var result SiteMeshGroup
	path := fmt.Sprintf("/api/config/namespaces/%s/site_mesh_groups", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetSiteMeshGroup(ctx context.Context, namespace, name string) (*SiteMeshGroup, error) {
	var result SiteMeshGroup
	path := fmt.Sprintf("/api/config/namespaces/%s/site_mesh_groups/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateSiteMeshGroup(ctx context.Context, resource *SiteMeshGroup) (*SiteMeshGroup, error) {
	var result SiteMeshGroup
	path := fmt.Sprintf("/api/config/namespaces/%s/site_mesh_groups/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteSiteMeshGroup(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/site_mesh_groups/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== Site =====
type Site struct {
	Metadata Metadata         `json:"metadata"`
	Spec     SiteSpec `json:"spec"`
}
type SiteSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateSite(ctx context.Context, resource *Site) (*Site, error) {
	var result Site
	path := fmt.Sprintf("/api/config/namespaces/%s/sites", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetSite(ctx context.Context, namespace, name string) (*Site, error) {
	var result Site
	path := fmt.Sprintf("/api/config/namespaces/%s/sites/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateSite(ctx context.Context, resource *Site) (*Site, error) {
	var result Site
	path := fmt.Sprintf("/api/config/namespaces/%s/sites/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteSite(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/sites/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== Srv6NetworkSlice =====
type Srv6NetworkSlice struct {
	Metadata Metadata         `json:"metadata"`
	Spec     Srv6NetworkSliceSpec `json:"spec"`
}
type Srv6NetworkSliceSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateSrv6NetworkSlice(ctx context.Context, resource *Srv6NetworkSlice) (*Srv6NetworkSlice, error) {
	var result Srv6NetworkSlice
	path := fmt.Sprintf("/api/config/namespaces/%s/srv6_network_slices", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetSrv6NetworkSlice(ctx context.Context, namespace, name string) (*Srv6NetworkSlice, error) {
	var result Srv6NetworkSlice
	path := fmt.Sprintf("/api/config/namespaces/%s/srv6_network_slices/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateSrv6NetworkSlice(ctx context.Context, resource *Srv6NetworkSlice) (*Srv6NetworkSlice, error) {
	var result Srv6NetworkSlice
	path := fmt.Sprintf("/api/config/namespaces/%s/srv6_network_slices/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteSrv6NetworkSlice(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/srv6_network_slices/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== Subnet =====
type Subnet struct {
	Metadata Metadata         `json:"metadata"`
	Spec     SubnetSpec `json:"spec"`
}
type SubnetSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateSubnet(ctx context.Context, resource *Subnet) (*Subnet, error) {
	var result Subnet
	path := fmt.Sprintf("/api/config/namespaces/%s/subnets", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetSubnet(ctx context.Context, namespace, name string) (*Subnet, error) {
	var result Subnet
	path := fmt.Sprintf("/api/config/namespaces/%s/subnets/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateSubnet(ctx context.Context, resource *Subnet) (*Subnet, error) {
	var result Subnet
	path := fmt.Sprintf("/api/config/namespaces/%s/subnets/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteSubnet(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/subnets/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== TCPLoadBalancer =====
type TCPLoadBalancer struct {
	Metadata Metadata         `json:"metadata"`
	Spec     TCPLoadBalancerSpec `json:"spec"`
}
type TCPLoadBalancerSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateTCPLoadBalancer(ctx context.Context, resource *TCPLoadBalancer) (*TCPLoadBalancer, error) {
	var result TCPLoadBalancer
	path := fmt.Sprintf("/api/config/namespaces/%s/tcp_loadbalancers", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetTCPLoadBalancer(ctx context.Context, namespace, name string) (*TCPLoadBalancer, error) {
	var result TCPLoadBalancer
	path := fmt.Sprintf("/api/config/namespaces/%s/tcp_loadbalancers/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateTCPLoadBalancer(ctx context.Context, resource *TCPLoadBalancer) (*TCPLoadBalancer, error) {
	var result TCPLoadBalancer
	path := fmt.Sprintf("/api/config/namespaces/%s/tcp_loadbalancers/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteTCPLoadBalancer(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/tcp_loadbalancers/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== TenantConfiguration =====
type TenantConfiguration struct {
	Metadata Metadata         `json:"metadata"`
	Spec     TenantConfigurationSpec `json:"spec"`
}
type TenantConfigurationSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateTenantConfiguration(ctx context.Context, resource *TenantConfiguration) (*TenantConfiguration, error) {
	var result TenantConfiguration
	path := fmt.Sprintf("/api/config/namespaces/%s/tenant_configurations", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetTenantConfiguration(ctx context.Context, namespace, name string) (*TenantConfiguration, error) {
	var result TenantConfiguration
	path := fmt.Sprintf("/api/config/namespaces/%s/tenant_configurations/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateTenantConfiguration(ctx context.Context, resource *TenantConfiguration) (*TenantConfiguration, error) {
	var result TenantConfiguration
	path := fmt.Sprintf("/api/config/namespaces/%s/tenant_configurations/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteTenantConfiguration(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/tenant_configurations/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== TerraformParameters =====
type TerraformParameters struct {
	Metadata Metadata         `json:"metadata"`
	Spec     TerraformParametersSpec `json:"spec"`
}
type TerraformParametersSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateTerraformParameters(ctx context.Context, resource *TerraformParameters) (*TerraformParameters, error) {
	var result TerraformParameters
	path := fmt.Sprintf("/api/config/namespaces/%s/terraform_parameterses", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetTerraformParameters(ctx context.Context, namespace, name string) (*TerraformParameters, error) {
	var result TerraformParameters
	path := fmt.Sprintf("/api/config/namespaces/%s/terraform_parameterses/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateTerraformParameters(ctx context.Context, resource *TerraformParameters) (*TerraformParameters, error) {
	var result TerraformParameters
	path := fmt.Sprintf("/api/config/namespaces/%s/terraform_parameterses/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteTerraformParameters(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/terraform_parameterses/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== ThirdPartyApplication =====
type ThirdPartyApplication struct {
	Metadata Metadata         `json:"metadata"`
	Spec     ThirdPartyApplicationSpec `json:"spec"`
}
type ThirdPartyApplicationSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateThirdPartyApplication(ctx context.Context, resource *ThirdPartyApplication) (*ThirdPartyApplication, error) {
	var result ThirdPartyApplication
	path := fmt.Sprintf("/api/config/namespaces/%s/third_party_applications", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetThirdPartyApplication(ctx context.Context, namespace, name string) (*ThirdPartyApplication, error) {
	var result ThirdPartyApplication
	path := fmt.Sprintf("/api/config/namespaces/%s/third_party_applications/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateThirdPartyApplication(ctx context.Context, resource *ThirdPartyApplication) (*ThirdPartyApplication, error) {
	var result ThirdPartyApplication
	path := fmt.Sprintf("/api/config/namespaces/%s/third_party_applications/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteThirdPartyApplication(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/third_party_applications/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== TrustedCaList =====
type TrustedCaList struct {
	Metadata Metadata         `json:"metadata"`
	Spec     TrustedCaListSpec `json:"spec"`
}
type TrustedCaListSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateTrustedCaList(ctx context.Context, resource *TrustedCaList) (*TrustedCaList, error) {
	var result TrustedCaList
	path := fmt.Sprintf("/api/config/namespaces/%s/trusted_ca_lists", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetTrustedCaList(ctx context.Context, namespace, name string) (*TrustedCaList, error) {
	var result TrustedCaList
	path := fmt.Sprintf("/api/config/namespaces/%s/trusted_ca_lists/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateTrustedCaList(ctx context.Context, resource *TrustedCaList) (*TrustedCaList, error) {
	var result TrustedCaList
	path := fmt.Sprintf("/api/config/namespaces/%s/trusted_ca_lists/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteTrustedCaList(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/trusted_ca_lists/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== Tunnel =====
type Tunnel struct {
	Metadata Metadata         `json:"metadata"`
	Spec     TunnelSpec `json:"spec"`
}
type TunnelSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateTunnel(ctx context.Context, resource *Tunnel) (*Tunnel, error) {
	var result Tunnel
	path := fmt.Sprintf("/api/config/namespaces/%s/tunnels", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetTunnel(ctx context.Context, namespace, name string) (*Tunnel, error) {
	var result Tunnel
	path := fmt.Sprintf("/api/config/namespaces/%s/tunnels/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateTunnel(ctx context.Context, resource *Tunnel) (*Tunnel, error) {
	var result Tunnel
	path := fmt.Sprintf("/api/config/namespaces/%s/tunnels/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteTunnel(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/tunnels/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== UDPLoadBalancer =====
type UDPLoadBalancer struct {
	Metadata Metadata         `json:"metadata"`
	Spec     UDPLoadBalancerSpec `json:"spec"`
}
type UDPLoadBalancerSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateUDPLoadBalancer(ctx context.Context, resource *UDPLoadBalancer) (*UDPLoadBalancer, error) {
	var result UDPLoadBalancer
	path := fmt.Sprintf("/api/config/namespaces/%s/udp_loadbalancers", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetUDPLoadBalancer(ctx context.Context, namespace, name string) (*UDPLoadBalancer, error) {
	var result UDPLoadBalancer
	path := fmt.Sprintf("/api/config/namespaces/%s/udp_loadbalancers/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateUDPLoadBalancer(ctx context.Context, resource *UDPLoadBalancer) (*UDPLoadBalancer, error) {
	var result UDPLoadBalancer
	path := fmt.Sprintf("/api/config/namespaces/%s/udp_loadbalancers/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteUDPLoadBalancer(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/udp_loadbalancers/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== UsbPolicy =====
type UsbPolicy struct {
	Metadata Metadata         `json:"metadata"`
	Spec     UsbPolicySpec `json:"spec"`
}
type UsbPolicySpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateUsbPolicy(ctx context.Context, resource *UsbPolicy) (*UsbPolicy, error) {
	var result UsbPolicy
	path := fmt.Sprintf("/api/config/namespaces/%s/usb_policies", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetUsbPolicy(ctx context.Context, namespace, name string) (*UsbPolicy, error) {
	var result UsbPolicy
	path := fmt.Sprintf("/api/config/namespaces/%s/usb_policies/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateUsbPolicy(ctx context.Context, resource *UsbPolicy) (*UsbPolicy, error) {
	var result UsbPolicy
	path := fmt.Sprintf("/api/config/namespaces/%s/usb_policies/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteUsbPolicy(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/usb_policies/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== UserIdentification =====
type UserIdentification struct {
	Metadata Metadata         `json:"metadata"`
	Spec     UserIdentificationSpec `json:"spec"`
}
type UserIdentificationSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateUserIdentification(ctx context.Context, resource *UserIdentification) (*UserIdentification, error) {
	var result UserIdentification
	path := fmt.Sprintf("/api/config/namespaces/%s/user_identifications", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetUserIdentification(ctx context.Context, namespace, name string) (*UserIdentification, error) {
	var result UserIdentification
	path := fmt.Sprintf("/api/config/namespaces/%s/user_identifications/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateUserIdentification(ctx context.Context, resource *UserIdentification) (*UserIdentification, error) {
	var result UserIdentification
	path := fmt.Sprintf("/api/config/namespaces/%s/user_identifications/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteUserIdentification(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/user_identifications/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== UserToken =====
type UserToken struct {
	Metadata Metadata         `json:"metadata"`
	Spec     UserTokenSpec `json:"spec"`
}
type UserTokenSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateUserToken(ctx context.Context, resource *UserToken) (*UserToken, error) {
	var result UserToken
	path := fmt.Sprintf("/api/config/namespaces/%s/user_tokens", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetUserToken(ctx context.Context, namespace, name string) (*UserToken, error) {
	var result UserToken
	path := fmt.Sprintf("/api/config/namespaces/%s/user_tokens/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateUserToken(ctx context.Context, resource *UserToken) (*UserToken, error) {
	var result UserToken
	path := fmt.Sprintf("/api/config/namespaces/%s/user_tokens/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteUserToken(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/user_tokens/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== ViewInternal =====
type ViewInternal struct {
	Metadata Metadata         `json:"metadata"`
	Spec     ViewInternalSpec `json:"spec"`
}
type ViewInternalSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateViewInternal(ctx context.Context, resource *ViewInternal) (*ViewInternal, error) {
	var result ViewInternal
	path := fmt.Sprintf("/api/config/namespaces/%s/view_internals", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetViewInternal(ctx context.Context, namespace, name string) (*ViewInternal, error) {
	var result ViewInternal
	path := fmt.Sprintf("/api/config/namespaces/%s/view_internals/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateViewInternal(ctx context.Context, resource *ViewInternal) (*ViewInternal, error) {
	var result ViewInternal
	path := fmt.Sprintf("/api/config/namespaces/%s/view_internals/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteViewInternal(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/view_internals/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== VirtualHost =====
type VirtualHost struct {
	Metadata Metadata         `json:"metadata"`
	Spec     VirtualHostSpec `json:"spec"`
}
type VirtualHostSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateVirtualHost(ctx context.Context, resource *VirtualHost) (*VirtualHost, error) {
	var result VirtualHost
	path := fmt.Sprintf("/api/config/namespaces/%s/virtual_hosts", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetVirtualHost(ctx context.Context, namespace, name string) (*VirtualHost, error) {
	var result VirtualHost
	path := fmt.Sprintf("/api/config/namespaces/%s/virtual_hosts/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateVirtualHost(ctx context.Context, resource *VirtualHost) (*VirtualHost, error) {
	var result VirtualHost
	path := fmt.Sprintf("/api/config/namespaces/%s/virtual_hosts/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteVirtualHost(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/virtual_hosts/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== VirtualK8s =====
type VirtualK8s struct {
	Metadata Metadata         `json:"metadata"`
	Spec     VirtualK8sSpec `json:"spec"`
}
type VirtualK8sSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateVirtualK8s(ctx context.Context, resource *VirtualK8s) (*VirtualK8s, error) {
	var result VirtualK8s
	path := fmt.Sprintf("/api/config/namespaces/%s/virtual_k8ses", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetVirtualK8s(ctx context.Context, namespace, name string) (*VirtualK8s, error) {
	var result VirtualK8s
	path := fmt.Sprintf("/api/config/namespaces/%s/virtual_k8ses/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateVirtualK8s(ctx context.Context, resource *VirtualK8s) (*VirtualK8s, error) {
	var result VirtualK8s
	path := fmt.Sprintf("/api/config/namespaces/%s/virtual_k8ses/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteVirtualK8s(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/virtual_k8ses/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== VirtualNetwork =====
type VirtualNetwork struct {
	Metadata Metadata         `json:"metadata"`
	Spec     VirtualNetworkSpec `json:"spec"`
}
type VirtualNetworkSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateVirtualNetwork(ctx context.Context, resource *VirtualNetwork) (*VirtualNetwork, error) {
	var result VirtualNetwork
	path := fmt.Sprintf("/api/config/namespaces/%s/virtual_networks", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetVirtualNetwork(ctx context.Context, namespace, name string) (*VirtualNetwork, error) {
	var result VirtualNetwork
	path := fmt.Sprintf("/api/config/namespaces/%s/virtual_networks/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateVirtualNetwork(ctx context.Context, resource *VirtualNetwork) (*VirtualNetwork, error) {
	var result VirtualNetwork
	path := fmt.Sprintf("/api/config/namespaces/%s/virtual_networks/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteVirtualNetwork(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/virtual_networks/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== VirtualSite =====
type VirtualSite struct {
	Metadata Metadata         `json:"metadata"`
	Spec     VirtualSiteSpec `json:"spec"`
}
type VirtualSiteSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateVirtualSite(ctx context.Context, resource *VirtualSite) (*VirtualSite, error) {
	var result VirtualSite
	path := fmt.Sprintf("/api/config/namespaces/%s/virtual_sites", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetVirtualSite(ctx context.Context, namespace, name string) (*VirtualSite, error) {
	var result VirtualSite
	path := fmt.Sprintf("/api/config/namespaces/%s/virtual_sites/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateVirtualSite(ctx context.Context, resource *VirtualSite) (*VirtualSite, error) {
	var result VirtualSite
	path := fmt.Sprintf("/api/config/namespaces/%s/virtual_sites/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteVirtualSite(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/virtual_sites/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== VoltstackSite =====
type VoltstackSite struct {
	Metadata Metadata         `json:"metadata"`
	Spec     VoltstackSiteSpec `json:"spec"`
}
type VoltstackSiteSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateVoltstackSite(ctx context.Context, resource *VoltstackSite) (*VoltstackSite, error) {
	var result VoltstackSite
	path := fmt.Sprintf("/api/config/namespaces/%s/voltstack_sites", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetVoltstackSite(ctx context.Context, namespace, name string) (*VoltstackSite, error) {
	var result VoltstackSite
	path := fmt.Sprintf("/api/config/namespaces/%s/voltstack_sites/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateVoltstackSite(ctx context.Context, resource *VoltstackSite) (*VoltstackSite, error) {
	var result VoltstackSite
	path := fmt.Sprintf("/api/config/namespaces/%s/voltstack_sites/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteVoltstackSite(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/voltstack_sites/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== WAFExclusionPolicy =====
type WAFExclusionPolicy struct {
	Metadata Metadata         `json:"metadata"`
	Spec     WAFExclusionPolicySpec `json:"spec"`
}
type WAFExclusionPolicySpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateWAFExclusionPolicy(ctx context.Context, resource *WAFExclusionPolicy) (*WAFExclusionPolicy, error) {
	var result WAFExclusionPolicy
	path := fmt.Sprintf("/api/config/namespaces/%s/waf_exclusion_policies", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetWAFExclusionPolicy(ctx context.Context, namespace, name string) (*WAFExclusionPolicy, error) {
	var result WAFExclusionPolicy
	path := fmt.Sprintf("/api/config/namespaces/%s/waf_exclusion_policies/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateWAFExclusionPolicy(ctx context.Context, resource *WAFExclusionPolicy) (*WAFExclusionPolicy, error) {
	var result WAFExclusionPolicy
	path := fmt.Sprintf("/api/config/namespaces/%s/waf_exclusion_policies/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteWAFExclusionPolicy(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/waf_exclusion_policies/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== WorkloadFlavor =====
type WorkloadFlavor struct {
	Metadata Metadata         `json:"metadata"`
	Spec     WorkloadFlavorSpec `json:"spec"`
}
type WorkloadFlavorSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateWorkloadFlavor(ctx context.Context, resource *WorkloadFlavor) (*WorkloadFlavor, error) {
	var result WorkloadFlavor
	path := fmt.Sprintf("/api/config/namespaces/%s/workload_flavors", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetWorkloadFlavor(ctx context.Context, namespace, name string) (*WorkloadFlavor, error) {
	var result WorkloadFlavor
	path := fmt.Sprintf("/api/config/namespaces/%s/workload_flavors/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateWorkloadFlavor(ctx context.Context, resource *WorkloadFlavor) (*WorkloadFlavor, error) {
	var result WorkloadFlavor
	path := fmt.Sprintf("/api/config/namespaces/%s/workload_flavors/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteWorkloadFlavor(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/workload_flavors/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== Workload =====
type Workload struct {
	Metadata Metadata         `json:"metadata"`
	Spec     WorkloadSpec `json:"spec"`
}
type WorkloadSpec struct {
	Description string `json:"description,omitempty"`
}
func (c *Client) CreateWorkload(ctx context.Context, resource *Workload) (*Workload, error) {
	var result Workload
	path := fmt.Sprintf("/api/config/namespaces/%s/workloads", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetWorkload(ctx context.Context, namespace, name string) (*Workload, error) {
	var result Workload
	path := fmt.Sprintf("/api/config/namespaces/%s/workloads/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateWorkload(ctx context.Context, resource *Workload) (*Workload, error) {
	var result Workload
	path := fmt.Sprintf("/api/config/namespaces/%s/workloads/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteWorkload(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/workloads/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== DNSLoadBalancer =====
type DNSLoadBalancer struct {
	Metadata Metadata            `json:"metadata"`
	Spec     DNSLoadBalancerSpec `json:"spec"`
}
type DNSLoadBalancerSpec struct {
	Description string `json:"description,omitempty"`
}

func (c *Client) CreateDNSLoadBalancer(ctx context.Context, resource *DNSLoadBalancer) (*DNSLoadBalancer, error) {
	var result DNSLoadBalancer
	path := fmt.Sprintf("/api/config/namespaces/%s/dns_load_balancers", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetDNSLoadBalancer(ctx context.Context, namespace, name string) (*DNSLoadBalancer, error) {
	var result DNSLoadBalancer
	path := fmt.Sprintf("/api/config/namespaces/%s/dns_load_balancers/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateDNSLoadBalancer(ctx context.Context, resource *DNSLoadBalancer) (*DNSLoadBalancer, error) {
	var result DNSLoadBalancer
	path := fmt.Sprintf("/api/config/namespaces/%s/dns_load_balancers/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteDNSLoadBalancer(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/dns_load_balancers/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== Role =====
type Role struct {
	Metadata Metadata `json:"metadata"`
	Spec     RoleSpec `json:"spec"`
}
type RoleSpec struct {
	Description string `json:"description,omitempty"`
}

func (c *Client) CreateRole(ctx context.Context, resource *Role) (*Role, error) {
	var result Role
	path := fmt.Sprintf("/api/config/namespaces/%s/roles", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetRole(ctx context.Context, namespace, name string) (*Role, error) {
	var result Role
	path := fmt.Sprintf("/api/config/namespaces/%s/roles/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateRole(ctx context.Context, resource *Role) (*Role, error) {
	var result Role
	path := fmt.Sprintf("/api/config/namespaces/%s/roles/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteRole(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/roles/%s", namespace, name)
	return c.Delete(ctx, path)
}

// ===== SiteState =====
type SiteState struct {
	Metadata Metadata      `json:"metadata"`
	Spec     SiteStateSpec `json:"spec"`
}
type SiteStateSpec struct {
	Description string `json:"description,omitempty"`
	State       string `json:"state,omitempty"`
}

func (c *Client) CreateSiteState(ctx context.Context, resource *SiteState) (*SiteState, error) {
	var result SiteState
	path := fmt.Sprintf("/api/config/namespaces/%s/site_states", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) GetSiteState(ctx context.Context, namespace, name string) (*SiteState, error) {
	var result SiteState
	path := fmt.Sprintf("/api/config/namespaces/%s/site_states/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) UpdateSiteState(ctx context.Context, resource *SiteState) (*SiteState, error) {
	var result SiteState
	path := fmt.Sprintf("/api/config/namespaces/%s/site_states/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) DeleteSiteState(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/site_states/%s", namespace, name)
	return c.Delete(ctx, path)
}
