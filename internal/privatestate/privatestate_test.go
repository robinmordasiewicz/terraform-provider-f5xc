// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package privatestate

import (
	"encoding/json"
	"testing"
	"time"
)

func TestKeyConstants(t *testing.T) {
	// Verify key constants are non-empty and distinct
	keys := []string{
		KeyAPIMetadata,
		KeyResourceUID,
		KeyLastModified,
		KeyETag,
		KeyResourceVersion,
	}

	seen := make(map[string]bool)
	for _, key := range keys {
		if key == "" {
			t.Error("Key constant should not be empty")
		}
		if seen[key] {
			t.Errorf("Duplicate key constant: %q", key)
		}
		seen[key] = true
	}
}

func TestNewPrivateStateData(t *testing.T) {
	psd := NewPrivateStateData()

	if psd == nil {
		t.Fatal("NewPrivateStateData() returned nil")
	}
	if psd.Metadata.Custom == nil {
		t.Error("NewPrivateStateData() should initialize Custom map")
	}
}

func TestPrivateStateDataSetters(t *testing.T) {
	psd := NewPrivateStateData()

	// Test method chaining
	result := psd.
		SetUID("test-uid").
		SetResourceVersion("v1").
		SetETag("etag-123").
		SetGeneration(5).
		SetSpecHash("abc123")

	if result != psd {
		t.Error("Setter methods should return the same instance for chaining")
	}

	if psd.Metadata.UID != "test-uid" {
		t.Errorf("SetUID() = %q, expected test-uid", psd.Metadata.UID)
	}
	if psd.Metadata.ResourceVersion != "v1" {
		t.Errorf("SetResourceVersion() = %q, expected v1", psd.Metadata.ResourceVersion)
	}
	if psd.Metadata.ETag != "etag-123" {
		t.Errorf("SetETag() = %q, expected etag-123", psd.Metadata.ETag)
	}
	if psd.Metadata.Generation != 5 {
		t.Errorf("SetGeneration() = %d, expected 5", psd.Metadata.Generation)
	}
	if psd.Metadata.SpecHash != "abc123" {
		t.Errorf("SetSpecHash() = %q, expected abc123", psd.Metadata.SpecHash)
	}
}

func TestSetLastModified(t *testing.T) {
	psd := NewPrivateStateData()
	now := time.Now()

	psd.SetLastModified(now)

	if !psd.Metadata.LastModified.Equal(now) {
		t.Errorf("SetLastModified() = %v, expected %v", psd.Metadata.LastModified, now)
	}
}

func TestSetCustom(t *testing.T) {
	psd := NewPrivateStateData()

	psd.SetCustom("key1", "value1").SetCustom("key2", "value2")

	if psd.Metadata.Custom["key1"] != "value1" {
		t.Errorf("SetCustom() key1 = %q, expected value1", psd.Metadata.Custom["key1"])
	}
	if psd.Metadata.Custom["key2"] != "value2" {
		t.Errorf("SetCustom() key2 = %q, expected value2", psd.Metadata.Custom["key2"])
	}
}

func TestSetCustomInitializesMap(t *testing.T) {
	// Create instance without Custom map initialized
	psd := &PrivateStateData{
		Metadata: APIMetadata{
			Custom: nil,
		},
	}

	psd.SetCustom("key", "value")

	if psd.Metadata.Custom == nil {
		t.Error("SetCustom() should initialize Custom map if nil")
	}
	if psd.Metadata.Custom["key"] != "value" {
		t.Error("SetCustom() should set the value")
	}
}

func TestToJSONAndFromJSON(t *testing.T) {
	psd := NewPrivateStateData()
	now := time.Now().UTC().Truncate(time.Second) // Truncate for JSON precision

	psd.
		SetUID("uid-123").
		SetResourceVersion("v2").
		SetETag("etag-456").
		SetGeneration(10).
		SetSpecHash("hash-789").
		SetLastModified(now).
		SetCustom("custom_key", "custom_value")

	// Serialize to JSON
	data, err := psd.ToJSON()
	if err != nil {
		t.Fatalf("ToJSON() error: %v", err)
	}

	// Deserialize from JSON
	psd2 := NewPrivateStateData()
	if err := psd2.FromJSON(data); err != nil {
		t.Fatalf("FromJSON() error: %v", err)
	}

	// Verify all fields
	if psd2.Metadata.UID != "uid-123" {
		t.Errorf("FromJSON() UID = %q, expected uid-123", psd2.Metadata.UID)
	}
	if psd2.Metadata.ResourceVersion != "v2" {
		t.Errorf("FromJSON() ResourceVersion = %q, expected v2", psd2.Metadata.ResourceVersion)
	}
	if psd2.Metadata.ETag != "etag-456" {
		t.Errorf("FromJSON() ETag = %q, expected etag-456", psd2.Metadata.ETag)
	}
	if psd2.Metadata.Generation != 10 {
		t.Errorf("FromJSON() Generation = %d, expected 10", psd2.Metadata.Generation)
	}
	if psd2.Metadata.SpecHash != "hash-789" {
		t.Errorf("FromJSON() SpecHash = %q, expected hash-789", psd2.Metadata.SpecHash)
	}
	if psd2.Metadata.Custom["custom_key"] != "custom_value" {
		t.Errorf("FromJSON() Custom[custom_key] = %q, expected custom_value", psd2.Metadata.Custom["custom_key"])
	}
}

func TestDetectDrift(t *testing.T) {
	tests := []struct {
		name             string
		setup            func(*PrivateStateData)
		apiUID           string
		apiSpecHash      string
		apiGeneration    int64
		expectDrift      bool
		expectUIDChange  bool
		expectSpecChange bool
		expectGenChange  bool
	}{
		{
			name:          "no drift - empty metadata",
			setup:         func(psd *PrivateStateData) {},
			apiUID:        "uid-123",
			apiSpecHash:   "hash-456",
			apiGeneration: 1,
			expectDrift:   false,
		},
		{
			name: "UID changed",
			setup: func(psd *PrivateStateData) {
				psd.SetUID("old-uid")
			},
			apiUID:          "new-uid",
			apiSpecHash:     "",
			apiGeneration:   0,
			expectDrift:     true,
			expectUIDChange: true,
		},
		{
			name: "spec hash changed",
			setup: func(psd *PrivateStateData) {
				psd.SetSpecHash("old-hash")
			},
			apiUID:           "",
			apiSpecHash:      "new-hash",
			apiGeneration:    0,
			expectDrift:      true,
			expectSpecChange: true,
		},
		{
			name: "generation changed",
			setup: func(psd *PrivateStateData) {
				psd.SetGeneration(5)
			},
			apiUID:          "",
			apiSpecHash:     "",
			apiGeneration:   10,
			expectDrift:     true,
			expectGenChange: true,
		},
		{
			name: "no drift - same values",
			setup: func(psd *PrivateStateData) {
				psd.SetUID("uid-123").SetSpecHash("hash-456").SetGeneration(5)
			},
			apiUID:        "uid-123",
			apiSpecHash:   "hash-456",
			apiGeneration: 5,
			expectDrift:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			psd := NewPrivateStateData()
			tt.setup(psd)

			drift := psd.DetectDrift(tt.apiUID, tt.apiSpecHash, tt.apiGeneration)

			if drift.HasDrift != tt.expectDrift {
				t.Errorf("DetectDrift() HasDrift = %v, expected %v", drift.HasDrift, tt.expectDrift)
			}
			if drift.UIDChanged != tt.expectUIDChange {
				t.Errorf("DetectDrift() UIDChanged = %v, expected %v", drift.UIDChanged, tt.expectUIDChange)
			}
			if drift.SpecChanged != tt.expectSpecChange {
				t.Errorf("DetectDrift() SpecChanged = %v, expected %v", drift.SpecChanged, tt.expectSpecChange)
			}
			if drift.GenerationChanged != tt.expectGenChange {
				t.Errorf("DetectDrift() GenerationChanged = %v, expected %v", drift.GenerationChanged, tt.expectGenChange)
			}
		})
	}
}

func TestDriftInfoAsWarning(t *testing.T) {
	t.Run("with drift", func(t *testing.T) {
		drift := &DriftInfo{
			HasDrift: true,
			Details:  []string{"UID changed", "Spec changed"},
		}

		diags := drift.AsWarning()

		if len(diags.Warnings()) != 1 {
			t.Errorf("AsWarning() returned %d warnings, expected 1", len(diags.Warnings()))
		}
	})

	t.Run("without drift", func(t *testing.T) {
		drift := &DriftInfo{
			HasDrift: false,
		}

		diags := drift.AsWarning()

		if len(diags.Warnings()) != 0 {
			t.Errorf("AsWarning() returned %d warnings, expected 0", len(diags.Warnings()))
		}
	})
}

func TestHashSpec(t *testing.T) {
	t.Run("consistent hashing", func(t *testing.T) {
		spec := map[string]string{
			"name":      "test",
			"namespace": "default",
		}

		hash1, err1 := HashSpec(spec)
		hash2, err2 := HashSpec(spec)

		if err1 != nil || err2 != nil {
			t.Fatalf("HashSpec() error: %v, %v", err1, err2)
		}

		if hash1 != hash2 {
			t.Error("HashSpec() should return consistent hash for same input")
		}
	})

	t.Run("different specs different hashes", func(t *testing.T) {
		spec1 := map[string]string{"name": "test1"}
		spec2 := map[string]string{"name": "test2"}

		hash1, _ := HashSpec(spec1)
		hash2, _ := HashSpec(spec2)

		if hash1 == hash2 {
			t.Error("HashSpec() should return different hashes for different inputs")
		}
	})

	t.Run("non-empty hash", func(t *testing.T) {
		spec := map[string]string{"key": "value"}

		hash, err := HashSpec(spec)
		if err != nil {
			t.Fatalf("HashSpec() error: %v", err)
		}

		if hash == "" {
			t.Error("HashSpec() should return non-empty hash")
		}
	})
}

func TestAPIMetadataJSONSerialization(t *testing.T) {
	now := time.Now().UTC().Truncate(time.Second)

	metadata := APIMetadata{
		UID:             "test-uid",
		ResourceVersion: "v1",
		ETag:            "etag-123",
		LastModified:    now,
		Generation:      5,
		SpecHash:        "hash-abc",
		CreatedAt:       now,
		ModifiedAt:      now,
		Owner:           "test-owner",
		LabelsHash:      "labels-hash",
		Custom: map[string]string{
			"key": "value",
		},
	}

	// Serialize
	data, err := json.Marshal(metadata)
	if err != nil {
		t.Fatalf("json.Marshal() error: %v", err)
	}

	// Deserialize
	var metadata2 APIMetadata
	if err := json.Unmarshal(data, &metadata2); err != nil {
		t.Fatalf("json.Unmarshal() error: %v", err)
	}

	// Verify
	if metadata2.UID != metadata.UID {
		t.Errorf("UID mismatch: got %q, expected %q", metadata2.UID, metadata.UID)
	}
	if metadata2.Generation != metadata.Generation {
		t.Errorf("Generation mismatch: got %d, expected %d", metadata2.Generation, metadata.Generation)
	}
	if metadata2.Owner != metadata.Owner {
		t.Errorf("Owner mismatch: got %q, expected %q", metadata2.Owner, metadata.Owner)
	}
}
