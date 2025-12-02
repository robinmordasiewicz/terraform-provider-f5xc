// Copyright (c) F5XC Community
// SPDX-License-Identifier: MPL-2.0

package mocks

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"
)

func TestNewServer(t *testing.T) {
	s := NewServer()
	defer s.Close()

	if s.URL() == "" {
		t.Error("Expected non-empty URL")
	}
}

func TestServerCRUD(t *testing.T) {
	s := NewServer()
	defer s.Close()

	client := s.Client()

	// Test Create (POST)
	t.Run("Create", func(t *testing.T) {
		createBody := map[string]interface{}{
			"metadata": map[string]interface{}{
				"name":      "test-resource",
				"namespace": "test-ns",
			},
			"spec": map[string]interface{}{
				"field1": "value1",
			},
		}

		bodyBytes, _ := json.Marshal(createBody)
		resp, err := client.Post(s.URL()+"/api/config/namespaces/test-ns/test_resources", "application/json", bytes.NewReader(bodyBytes))
		if err != nil {
			t.Fatalf("POST failed: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			t.Errorf("Expected status 200, got %d: %s", resp.StatusCode, string(body))
		}

		var result map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&result)

		metadata, ok := result["metadata"].(map[string]interface{})
		if !ok {
			t.Fatal("Expected metadata in response")
		}
		if metadata["name"] != "test-resource" {
			t.Errorf("Expected name 'test-resource', got %v", metadata["name"])
		}
	})

	// Test Read (GET)
	t.Run("Read", func(t *testing.T) {
		resp, err := client.Get(s.URL() + "/api/config/namespaces/test-ns/test_resources/test-resource")
		if err != nil {
			t.Fatalf("GET failed: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			t.Errorf("Expected status 200, got %d: %s", resp.StatusCode, string(body))
		}
	})

	// Test Update (PUT)
	t.Run("Update", func(t *testing.T) {
		updateBody := map[string]interface{}{
			"metadata": map[string]interface{}{
				"name":        "test-resource",
				"namespace":   "test-ns",
				"description": "updated description",
			},
			"spec": map[string]interface{}{
				"field1": "updated-value",
			},
		}

		bodyBytes, _ := json.Marshal(updateBody)
		req, _ := http.NewRequest(http.MethodPut, s.URL()+"/api/config/namespaces/test-ns/test_resources/test-resource", bytes.NewReader(bodyBytes))
		req.Header.Set("Content-Type", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("PUT failed: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			t.Errorf("Expected status 200, got %d: %s", resp.StatusCode, string(body))
		}
	})

	// Test Delete (DELETE)
	t.Run("Delete", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodDelete, s.URL()+"/api/config/namespaces/test-ns/test_resources/test-resource", nil)
		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("DELETE failed: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			t.Errorf("Expected status 200, got %d: %s", resp.StatusCode, string(body))
		}

		// Verify resource is deleted
		resp2, err := client.Get(s.URL() + "/api/config/namespaces/test-ns/test_resources/test-resource")
		if err != nil {
			t.Fatalf("GET after DELETE failed: %v", err)
		}
		defer resp2.Body.Close()

		if resp2.StatusCode != http.StatusNotFound {
			t.Errorf("Expected status 404 after delete, got %d", resp2.StatusCode)
		}
	})
}

func TestServerSetResource(t *testing.T) {
	s := NewServer()
	defer s.Close()

	// Pre-populate a resource
	resource := NamespaceResponse("test-ns", nil, nil, "test description")
	s.SetResource("/api/config/namespaces/system/namespaces/test-ns", resource)

	// Verify we can retrieve it
	client := s.Client()
	resp, err := client.Get(s.URL() + "/api/config/namespaces/system/namespaces/test-ns")
	if err != nil {
		t.Fatalf("GET failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
}

func TestServerErrorResponse(t *testing.T) {
	s := NewServer()
	defer s.Close()

	// Configure error response
	s.SetErrorResponse("/api/config/namespaces/system/failing_resource/test", http.StatusInternalServerError, map[string]string{
		"code":    "INTERNAL_ERROR",
		"message": "Simulated server error",
	})

	client := s.Client()
	resp, err := client.Get(s.URL() + "/api/config/namespaces/system/failing_resource/test")
	if err != nil {
		t.Fatalf("GET failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("Expected status 500, got %d", resp.StatusCode)
	}
}

func TestServerCustomHandler(t *testing.T) {
	s := NewServer()
	defer s.Close()

	// Set custom handler
	s.SetHandler("/api/custom/.*", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"custom": "response",
		})
	})

	client := s.Client()
	resp, err := client.Get(s.URL() + "/api/custom/path")
	if err != nil {
		t.Fatalf("GET failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	var result map[string]string
	json.NewDecoder(resp.Body).Decode(&result)
	if result["custom"] != "response" {
		t.Errorf("Expected custom response, got %v", result)
	}
}

func TestServerRequestLog(t *testing.T) {
	s := NewServer()
	defer s.Close()

	// Pre-populate a resource for the GET to succeed
	s.SetResource("/api/config/namespaces/system/test/resource1", map[string]interface{}{})

	client := s.Client()
	client.Get(s.URL() + "/api/config/namespaces/system/test/resource1")

	log := s.GetRequestLog()
	if len(log) == 0 {
		t.Error("Expected at least one request in log")
	}

	if log[0].Method != "GET" {
		t.Errorf("Expected GET method, got %s", log[0].Method)
	}
}

func TestServerResponseDelay(t *testing.T) {
	s := NewServer()
	defer s.Close()

	s.SetResponseDelay(100 * time.Millisecond)

	// Pre-populate resource
	s.SetResource("/api/test", map[string]interface{}{})

	client := s.Client()
	start := time.Now()
	resp, _ := client.Get(s.URL() + "/api/test")
	elapsed := time.Since(start)
	resp.Body.Close()

	if elapsed < 100*time.Millisecond {
		t.Errorf("Expected delay of at least 100ms, got %v", elapsed)
	}
}

func TestServerReset(t *testing.T) {
	s := NewServer()
	defer s.Close()

	// Add resource and error
	s.SetResource("/api/test", map[string]interface{}{})
	s.SetErrorResponse("/api/error", 500, nil)

	// Reset
	s.Reset()

	// Verify resource is gone
	_, exists := s.GetResource("/api/test")
	if exists {
		t.Error("Expected resource to be cleared after reset")
	}

	// Verify log is cleared
	log := s.GetRequestLog()
	if len(log) != 0 {
		t.Error("Expected request log to be cleared after reset")
	}
}

func TestResourcePath(t *testing.T) {
	path := ResourcePath("my-namespace", "origin_pools", "my-pool")
	expected := "/api/config/namespaces/my-namespace/origin_pools/my-pool"
	if path != expected {
		t.Errorf("Expected %s, got %s", expected, path)
	}
}

func TestListPath(t *testing.T) {
	path := ListPath("my-namespace", "origin_pools")
	expected := "/api/config/namespaces/my-namespace/origin_pools"
	if path != expected {
		t.Errorf("Expected %s, got %s", expected, path)
	}
}

func TestExtractResourceInfo(t *testing.T) {
	tests := []struct {
		path     string
		wantNS   string
		wantType string
		wantName string
		wantOK   bool
	}{
		{
			path:     "/api/config/namespaces/system/aws_vpc_sites/my-site",
			wantNS:   "system",
			wantType: "aws_vpc_sites",
			wantName: "my-site",
			wantOK:   true,
		},
		{
			path:     "/api/config/namespaces/my-ns/origin_pools/pool1",
			wantNS:   "my-ns",
			wantType: "origin_pools",
			wantName: "pool1",
			wantOK:   true,
		},
		{
			path:   "/invalid/path",
			wantOK: false,
		},
	}

	for _, tt := range tests {
		ns, rt, name, ok := ExtractResourceInfo(tt.path)
		if ok != tt.wantOK {
			t.Errorf("ExtractResourceInfo(%s) ok = %v, want %v", tt.path, ok, tt.wantOK)
			continue
		}
		if ok {
			if ns != tt.wantNS || rt != tt.wantType || name != tt.wantName {
				t.Errorf("ExtractResourceInfo(%s) = (%s, %s, %s), want (%s, %s, %s)",
					tt.path, ns, rt, name, tt.wantNS, tt.wantType, tt.wantName)
			}
		}
	}
}

func TestFixtures(t *testing.T) {
	t.Run("NamespaceResponse", func(t *testing.T) {
		resp := NamespaceResponse("test-ns", map[string]string{"env": "test"}, nil, "Test namespace")
		metadata := resp["metadata"].(map[string]interface{})
		if metadata["name"] != "test-ns" {
			t.Errorf("Expected name 'test-ns', got %v", metadata["name"])
		}
		if metadata["labels"].(map[string]string)["env"] != "test" {
			t.Error("Expected label 'env: test'")
		}
	})

	t.Run("AWSVPCSiteResponse", func(t *testing.T) {
		resp := AWSVPCSiteResponse("system", "my-site", WithAWSRegion("us-west-2"))
		spec := resp["spec"].(map[string]interface{})
		if spec["aws_region"] != "us-west-2" {
			t.Errorf("Expected region 'us-west-2', got %v", spec["aws_region"])
		}
	})

	t.Run("CloudCredentialsResponse", func(t *testing.T) {
		resp := CloudCredentialsResponse("system", "aws-creds", "aws")
		spec := resp["spec"].(map[string]interface{})
		if _, ok := spec["aws_secret_credentials"]; !ok {
			t.Error("Expected aws_secret_credentials in spec")
		}
	})

	t.Run("GenericResourceResponse", func(t *testing.T) {
		resp := GenericResourceResponse("my-ns", "my-resource", "custom_type", map[string]interface{}{
			"custom_field": "custom_value",
		})
		metadata := resp["metadata"].(map[string]interface{})
		if metadata["name"] != "my-resource" {
			t.Errorf("Expected name 'my-resource', got %v", metadata["name"])
		}
		spec := resp["spec"].(map[string]interface{})
		if spec["custom_field"] != "custom_value" {
			t.Errorf("Expected custom_field 'custom_value', got %v", spec["custom_field"])
		}
	})
}
