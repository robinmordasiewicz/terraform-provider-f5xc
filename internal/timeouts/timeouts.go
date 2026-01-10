// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

// Package timeouts provides timeout configuration for F5 XC Terraform resources
// following Terraform Plugin Framework best practices.
package timeouts

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Default timeout values for F5 XC operations
const (
	// DefaultCreate is the default timeout for create operations
	DefaultCreate = 10 * time.Minute

	// DefaultRead is the default timeout for read operations
	DefaultRead = 5 * time.Minute

	// DefaultUpdate is the default timeout for update operations
	DefaultUpdate = 10 * time.Minute

	// DefaultDelete is the default timeout for delete operations
	DefaultDelete = 10 * time.Minute

	// LongRunningCreate is the timeout for long-running create operations (e.g., sites)
	LongRunningCreate = 30 * time.Minute

	// LongRunningDelete is the timeout for long-running delete operations (e.g., sites)
	LongRunningDelete = 30 * time.Minute
)

// ResourceType categorizes resources by their typical operation duration
type ResourceType int

const (
	// Standard resources have quick API operations
	Standard ResourceType = iota

	// LongRunning resources involve infrastructure provisioning (sites, clusters)
	LongRunning

	// Critical resources require extra time for safety
	Critical
)

// Config holds timeout configuration for a resource
type Config struct {
	Create time.Duration
	Read   time.Duration
	Update time.Duration
	Delete time.Duration
}

// DefaultConfig returns the standard timeout configuration
func DefaultConfig() Config {
	return Config{
		Create: DefaultCreate,
		Read:   DefaultRead,
		Update: DefaultUpdate,
		Delete: DefaultDelete,
	}
}

// LongRunningConfig returns timeout configuration for long-running resources
func LongRunningConfig() Config {
	return Config{
		Create: LongRunningCreate,
		Read:   DefaultRead,
		Update: LongRunningCreate,
		Delete: LongRunningDelete,
	}
}

// ConfigForResourceType returns appropriate timeout config based on resource type
func ConfigForResourceType(rt ResourceType) Config {
	switch rt {
	case LongRunning:
		return LongRunningConfig()
	case Critical:
		return Config{
			Create: 20 * time.Minute,
			Read:   DefaultRead,
			Update: 20 * time.Minute,
			Delete: 20 * time.Minute,
		}
	default:
		return DefaultConfig()
	}
}

// Block returns the timeouts block for schema definition
func Block(cfg Config) schema.Block {
	return timeouts.Block(context.Background(), timeouts.Opts{
		Create: true,
		Read:   true,
		Update: true,
		Delete: true,
	})
}

// BlockWithDefaults returns the timeouts block with default values displayed
func BlockWithDefaults(cfg Config) schema.Block {
	return timeouts.BlockAll(context.Background())
}

// StandardBlock returns the standard timeouts block for most resources
func StandardBlock() schema.Block {
	return Block(DefaultConfig())
}

// LongRunningBlock returns timeouts block for long-running resources
func LongRunningBlock() schema.Block {
	return Block(LongRunningConfig())
}

// WithContext creates a context with the specified timeout
func WithContext(ctx context.Context, timeout time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(ctx, timeout)
}

// CreateContext creates a context for create operations
func CreateContext(ctx context.Context, t timeouts.Value) (context.Context, context.CancelFunc) {
	timeout, diags := t.Create(ctx, DefaultCreate)
	if diags.HasError() {
		timeout = DefaultCreate
	}
	return context.WithTimeout(ctx, timeout)
}

// ReadContext creates a context for read operations
func ReadContext(ctx context.Context, t timeouts.Value) (context.Context, context.CancelFunc) {
	timeout, diags := t.Read(ctx, DefaultRead)
	if diags.HasError() {
		timeout = DefaultRead
	}
	return context.WithTimeout(ctx, timeout)
}

// UpdateContext creates a context for update operations
func UpdateContext(ctx context.Context, t timeouts.Value) (context.Context, context.CancelFunc) {
	timeout, diags := t.Update(ctx, DefaultUpdate)
	if diags.HasError() {
		timeout = DefaultUpdate
	}
	return context.WithTimeout(ctx, timeout)
}

// DeleteContext creates a context for delete operations
func DeleteContext(ctx context.Context, t timeouts.Value) (context.Context, context.CancelFunc) {
	timeout, diags := t.Delete(ctx, DefaultDelete)
	if diags.HasError() {
		timeout = DefaultDelete
	}
	return context.WithTimeout(ctx, timeout)
}

// LongRunningResourceTypes returns the list of resource types that need longer timeouts
var LongRunningResourceTypes = map[string]bool{
	"aws_vpc_site":       true,
	"azure_vnet_site":    true,
	"gcp_vpc_site":       true,
	"aws_tgw_site":       true,
	"voltstack_site":     true,
	"securemesh_site":    true,
	"securemesh_site_v2": true,
	"k8s_cluster":        true,
	"virtual_k8s":        true,
}

// IsLongRunning returns true if the resource type typically has long operations
func IsLongRunning(resourceType string) bool {
	return LongRunningResourceTypes[resourceType]
}
