// Copyright (c) F5XC Community
// SPDX-License-Identifier: MPL-2.0

package timeouts

import (
	"context"
	"testing"
	"time"
)

func TestDefaultConstants(t *testing.T) {
	tests := []struct {
		name     string
		constant time.Duration
		expected time.Duration
	}{
		{"DefaultCreate", DefaultCreate, 10 * time.Minute},
		{"DefaultRead", DefaultRead, 5 * time.Minute},
		{"DefaultUpdate", DefaultUpdate, 10 * time.Minute},
		{"DefaultDelete", DefaultDelete, 10 * time.Minute},
		{"LongRunningCreate", LongRunningCreate, 30 * time.Minute},
		{"LongRunningDelete", LongRunningDelete, 30 * time.Minute},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.constant != tt.expected {
				t.Errorf("%s = %v, expected %v", tt.name, tt.constant, tt.expected)
			}
		})
	}
}

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()

	if cfg.Create != DefaultCreate {
		t.Errorf("DefaultConfig().Create = %v, expected %v", cfg.Create, DefaultCreate)
	}
	if cfg.Read != DefaultRead {
		t.Errorf("DefaultConfig().Read = %v, expected %v", cfg.Read, DefaultRead)
	}
	if cfg.Update != DefaultUpdate {
		t.Errorf("DefaultConfig().Update = %v, expected %v", cfg.Update, DefaultUpdate)
	}
	if cfg.Delete != DefaultDelete {
		t.Errorf("DefaultConfig().Delete = %v, expected %v", cfg.Delete, DefaultDelete)
	}
}

func TestLongRunningConfig(t *testing.T) {
	cfg := LongRunningConfig()

	if cfg.Create != LongRunningCreate {
		t.Errorf("LongRunningConfig().Create = %v, expected %v", cfg.Create, LongRunningCreate)
	}
	if cfg.Read != DefaultRead {
		t.Errorf("LongRunningConfig().Read = %v, expected %v", cfg.Read, DefaultRead)
	}
	if cfg.Update != LongRunningCreate {
		t.Errorf("LongRunningConfig().Update = %v, expected %v", cfg.Update, LongRunningCreate)
	}
	if cfg.Delete != LongRunningDelete {
		t.Errorf("LongRunningConfig().Delete = %v, expected %v", cfg.Delete, LongRunningDelete)
	}
}

func TestConfigForResourceType(t *testing.T) {
	tests := []struct {
		name         string
		resourceType ResourceType
		expectedCfg  Config
	}{
		{
			name:         "Standard resource",
			resourceType: Standard,
			expectedCfg:  DefaultConfig(),
		},
		{
			name:         "LongRunning resource",
			resourceType: LongRunning,
			expectedCfg:  LongRunningConfig(),
		},
		{
			name:         "Critical resource",
			resourceType: Critical,
			expectedCfg: Config{
				Create: 20 * time.Minute,
				Read:   DefaultRead,
				Update: 20 * time.Minute,
				Delete: 20 * time.Minute,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := ConfigForResourceType(tt.resourceType)
			if cfg != tt.expectedCfg {
				t.Errorf("ConfigForResourceType(%v) = %+v, expected %+v", tt.resourceType, cfg, tt.expectedCfg)
			}
		})
	}
}

func TestIsLongRunning(t *testing.T) {
	tests := []struct {
		resourceType string
		expected     bool
	}{
		{"aws_vpc_site", true},
		{"azure_vnet_site", true},
		{"gcp_vpc_site", true},
		{"aws_tgw_site", true},
		{"voltstack_site", true},
		{"securemesh_site", true},
		{"securemesh_site_v2", true},
		{"k8s_cluster", true},
		{"virtual_k8s", true},
		{"namespace", false},
		{"http_loadbalancer", false},
		{"origin_pool", false},
		{"unknown_resource", false},
	}

	for _, tt := range tests {
		t.Run(tt.resourceType, func(t *testing.T) {
			result := IsLongRunning(tt.resourceType)
			if result != tt.expected {
				t.Errorf("IsLongRunning(%q) = %v, expected %v", tt.resourceType, result, tt.expected)
			}
		})
	}
}

func TestWithContext(t *testing.T) {
	ctx := context.Background()
	timeout := 5 * time.Second

	newCtx, cancel := WithContext(ctx, timeout)
	defer cancel()

	if newCtx == nil {
		t.Error("WithContext returned nil context")
	}

	deadline, ok := newCtx.Deadline()
	if !ok {
		t.Error("WithContext did not set deadline")
	}

	// Deadline should be approximately timeout from now
	expectedDeadline := time.Now().Add(timeout)
	if deadline.Before(expectedDeadline.Add(-1*time.Second)) || deadline.After(expectedDeadline.Add(1*time.Second)) {
		t.Errorf("Deadline %v not within expected range around %v", deadline, expectedDeadline)
	}
}

func TestLongRunningResourceTypes(t *testing.T) {
	// Verify the map is populated
	if len(LongRunningResourceTypes) == 0 {
		t.Error("LongRunningResourceTypes map is empty")
	}

	// Verify known resources are in the map
	expectedResources := []string{
		"aws_vpc_site",
		"azure_vnet_site",
		"gcp_vpc_site",
		"k8s_cluster",
		"virtual_k8s",
	}

	for _, resource := range expectedResources {
		if !LongRunningResourceTypes[resource] {
			t.Errorf("Expected %q to be in LongRunningResourceTypes", resource)
		}
	}
}

func TestResourceTypeConstants(t *testing.T) {
	// Verify resource type constants are distinct
	if Standard == LongRunning {
		t.Error("Standard and LongRunning should be distinct values")
	}
	if Standard == Critical {
		t.Error("Standard and Critical should be distinct values")
	}
	if LongRunning == Critical {
		t.Error("LongRunning and Critical should be distinct values")
	}
}

func TestBlock(t *testing.T) {
	cfg := DefaultConfig()
	block := Block(cfg)

	if block == nil {
		t.Error("Block() returned nil")
	}
}

func TestStandardBlock(t *testing.T) {
	block := StandardBlock()

	if block == nil {
		t.Error("StandardBlock() returned nil")
	}
}

func TestLongRunningBlock(t *testing.T) {
	block := LongRunningBlock()

	if block == nil {
		t.Error("LongRunningBlock() returned nil")
	}
}
