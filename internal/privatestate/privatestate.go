// Copyright (c) F5XC Community
// SPDX-License-Identifier: MPL-2.0

// Package privatestate provides private state management for F5 XC Terraform resources.
// Private state stores API metadata for drift detection following Terraform Plugin Framework best practices.
package privatestate

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Keys for private state data
const (
	KeyAPIMetadata     = "api_metadata"
	KeyResourceUID     = "resource_uid"
	KeyLastModified    = "last_modified"
	KeyETag            = "etag"
	KeyResourceVersion = "resource_version"
)

// APIMetadata stores metadata from the F5 XC API for drift detection
type APIMetadata struct {
	// UID is the unique identifier from the API
	UID string `json:"uid,omitempty"`

	// ResourceVersion is used for optimistic locking
	ResourceVersion string `json:"resource_version,omitempty"`

	// ETag for HTTP caching
	ETag string `json:"etag,omitempty"`

	// LastModified timestamp from the API
	LastModified time.Time `json:"last_modified,omitempty"`

	// Generation is incremented by the API on spec changes
	Generation int64 `json:"generation,omitempty"`

	// Hash of the spec for detecting external changes
	SpecHash string `json:"spec_hash,omitempty"`

	// CreatedAt when the resource was created
	CreatedAt time.Time `json:"created_at,omitempty"`

	// ModifiedAt when the resource was last modified
	ModifiedAt time.Time `json:"modified_at,omitempty"`

	// Owner information
	Owner string `json:"owner,omitempty"`

	// Labels snapshot for drift detection
	LabelsHash string `json:"labels_hash,omitempty"`

	// Custom fields for resource-specific metadata
	Custom map[string]string `json:"custom,omitempty"`
}

// PrivateStateData provides methods for reading and writing private state
type PrivateStateData struct {
	Metadata APIMetadata
}

// NewPrivateStateData creates a new private state data instance
func NewPrivateStateData() *PrivateStateData {
	return &PrivateStateData{
		Metadata: APIMetadata{
			Custom: make(map[string]string),
		},
	}
}

// SetUID sets the API UID
func (p *PrivateStateData) SetUID(uid string) *PrivateStateData {
	p.Metadata.UID = uid
	return p
}

// SetResourceVersion sets the resource version for optimistic locking
func (p *PrivateStateData) SetResourceVersion(version string) *PrivateStateData {
	p.Metadata.ResourceVersion = version
	return p
}

// SetETag sets the ETag for HTTP caching
func (p *PrivateStateData) SetETag(etag string) *PrivateStateData {
	p.Metadata.ETag = etag
	return p
}

// SetLastModified sets the last modified timestamp
func (p *PrivateStateData) SetLastModified(t time.Time) *PrivateStateData {
	p.Metadata.LastModified = t
	return p
}

// SetGeneration sets the generation number
func (p *PrivateStateData) SetGeneration(gen int64) *PrivateStateData {
	p.Metadata.Generation = gen
	return p
}

// SetSpecHash sets the spec hash for drift detection
func (p *PrivateStateData) SetSpecHash(hash string) *PrivateStateData {
	p.Metadata.SpecHash = hash
	return p
}

// SetCustom sets a custom metadata field
func (p *PrivateStateData) SetCustom(key, value string) *PrivateStateData {
	if p.Metadata.Custom == nil {
		p.Metadata.Custom = make(map[string]string)
	}
	p.Metadata.Custom[key] = value
	return p
}

// ToJSON serializes the private state to JSON
func (p *PrivateStateData) ToJSON() ([]byte, error) {
	return json.Marshal(p.Metadata)
}

// FromJSON deserializes the private state from JSON
func (p *PrivateStateData) FromJSON(data []byte) error {
	return json.Unmarshal(data, &p.Metadata)
}

// SaveToPrivateState saves the metadata to the response's private state
func (p *PrivateStateData) SaveToPrivateState(ctx context.Context, private interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	data, err := p.ToJSON()
	if err != nil {
		diags.AddError(
			"Error Saving Private State",
			fmt.Sprintf("Failed to serialize private state: %s", err),
		)
		return diags
	}

	// Handle different response types
	switch resp := private.(type) {
	case *resource.CreateResponse:
		diags.Append(resp.Private.SetKey(ctx, KeyAPIMetadata, data)...)
	case *resource.UpdateResponse:
		diags.Append(resp.Private.SetKey(ctx, KeyAPIMetadata, data)...)
	case *resource.ReadResponse:
		diags.Append(resp.Private.SetKey(ctx, KeyAPIMetadata, data)...)
	default:
		diags.AddError(
			"Error Saving Private State",
			"Unsupported response type for private state",
		)
	}

	return diags
}

// LoadFromPrivateState loads metadata from the request's private state
func LoadFromPrivateState(ctx context.Context, private interface{}) (*PrivateStateData, diag.Diagnostics) {
	var diags diag.Diagnostics
	psd := NewPrivateStateData()

	var data []byte
	var getDiags diag.Diagnostics

	// Handle different request types
	switch req := private.(type) {
	case *resource.ReadRequest:
		data, getDiags = req.Private.GetKey(ctx, KeyAPIMetadata)
	case *resource.UpdateRequest:
		data, getDiags = req.Private.GetKey(ctx, KeyAPIMetadata)
	case *resource.DeleteRequest:
		data, getDiags = req.Private.GetKey(ctx, KeyAPIMetadata)
	default:
		diags.AddWarning(
			"No Private State Available",
			"Private state is not available for this request type",
		)
		return psd, diags
	}

	diags.Append(getDiags...)
	if diags.HasError() {
		return psd, diags
	}

	if len(data) == 0 {
		return psd, diags
	}

	if err := psd.FromJSON(data); err != nil {
		diags.AddWarning(
			"Error Loading Private State",
			fmt.Sprintf("Failed to deserialize private state: %s. Using empty state.", err),
		)
	}

	return psd, diags
}

// DetectDrift compares API state with private state and returns drift information
func (p *PrivateStateData) DetectDrift(apiUID, apiSpecHash string, apiGeneration int64) *DriftInfo {
	drift := &DriftInfo{
		HasDrift: false,
	}

	// Check UID change (shouldn't happen but indicates serious drift)
	if p.Metadata.UID != "" && p.Metadata.UID != apiUID {
		drift.HasDrift = true
		drift.UIDChanged = true
		drift.Details = append(drift.Details, fmt.Sprintf("UID changed from %s to %s", p.Metadata.UID, apiUID))
	}

	// Check spec hash for external modifications
	if p.Metadata.SpecHash != "" && p.Metadata.SpecHash != apiSpecHash {
		drift.HasDrift = true
		drift.SpecChanged = true
		drift.Details = append(drift.Details, "Resource spec was modified externally")
	}

	// Check generation for any modifications
	if p.Metadata.Generation > 0 && p.Metadata.Generation != apiGeneration {
		drift.HasDrift = true
		drift.GenerationChanged = true
		drift.Details = append(drift.Details, fmt.Sprintf("Generation changed from %d to %d", p.Metadata.Generation, apiGeneration))
	}

	return drift
}

// DriftInfo contains information about detected drift
type DriftInfo struct {
	HasDrift          bool
	UIDChanged        bool
	SpecChanged       bool
	GenerationChanged bool
	Details           []string
}

// AsWarning returns drift information as a diagnostic warning
func (d *DriftInfo) AsWarning() diag.Diagnostics {
	var diags diag.Diagnostics

	if d.HasDrift {
		diags.AddWarning(
			"Resource Drift Detected",
			fmt.Sprintf("External changes detected: %v. Consider running 'terraform refresh' to update state.", d.Details),
		)
	}

	return diags
}

// HashSpec creates a simple hash of a spec for drift detection
// In production, use a proper hashing library like crypto/sha256
func HashSpec(spec interface{}) (string, error) {
	data, err := json.Marshal(spec)
	if err != nil {
		return "", err
	}

	// Simple hash for demonstration - in production use crypto/sha256
	var hash uint64
	for _, b := range data {
		hash = hash*31 + uint64(b)
	}
	return fmt.Sprintf("%x", hash), nil
}
