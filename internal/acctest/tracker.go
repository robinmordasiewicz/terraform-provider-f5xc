// Copyright (c) F5XC Community
// SPDX-License-Identifier: MPL-2.0

// Package acctest provides acceptance test utilities including resource tracking
// for safe multi-user test cleanup.
package acctest

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/f5xc/terraform-provider-f5xc/internal/client"
)

// TrackedResource represents a resource created during a test run
type TrackedResource struct {
	Type      string    // Resource type (e.g., "f5xc_namespace", "f5xc_http_loadbalancer")
	Name      string    // Resource name
	Namespace string    // Namespace (empty for namespace resources)
	CreatedAt time.Time // When the resource was registered
}

// ResourceTracker tracks resources created during test runs for safe cleanup.
// This ensures that cleanup only removes resources created by the current test
// session, not resources created by other users running tests concurrently.
type ResourceTracker struct {
	mu        sync.RWMutex
	resources []TrackedResource
}

// globalTracker is the singleton tracker for the current test session
var (
	globalTracker     *ResourceTracker
	globalTrackerOnce sync.Once
)

// GetTracker returns the global resource tracker instance
func GetTracker() *ResourceTracker {
	globalTrackerOnce.Do(func() {
		globalTracker = &ResourceTracker{
			resources: make([]TrackedResource, 0),
		}
	})
	return globalTracker
}

// Track registers a resource for cleanup after tests complete
func (t *ResourceTracker) Track(resourceType, name, namespace string) {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.resources = append(t.resources, TrackedResource{
		Type:      resourceType,
		Name:      name,
		Namespace: namespace,
		CreatedAt: time.Now(),
	})

	log.Printf("[TRACKER] Registered %s: %s (namespace: %s)", resourceType, name, namespace)
}

// TrackNamespace is a convenience method for tracking namespace resources
func (t *ResourceTracker) TrackNamespace(name string) {
	t.Track("f5xc_namespace", name, "")
}

// TrackResource is a convenience method for tracking namespaced resources
func (t *ResourceTracker) TrackResource(resourceType, name, namespace string) {
	t.Track(resourceType, name, namespace)
}

// GetTracked returns all tracked resources (thread-safe copy)
func (t *ResourceTracker) GetTracked() []TrackedResource {
	t.mu.RLock()
	defer t.mu.RUnlock()

	result := make([]TrackedResource, len(t.resources))
	copy(result, t.resources)
	return result
}

// GetTrackedByType returns tracked resources of a specific type
func (t *ResourceTracker) GetTrackedByType(resourceType string) []TrackedResource {
	t.mu.RLock()
	defer t.mu.RUnlock()

	var result []TrackedResource
	for _, r := range t.resources {
		if r.Type == resourceType {
			result = append(result, r)
		}
	}
	return result
}

// Clear removes all tracked resources (used after cleanup or at test end)
func (t *ResourceTracker) Clear() {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.resources = make([]TrackedResource, 0)
	log.Printf("[TRACKER] Cleared all tracked resources")
}

// Remove removes a specific resource from tracking (e.g., after successful delete)
func (t *ResourceTracker) Remove(resourceType, name, namespace string) {
	t.mu.Lock()
	defer t.mu.Unlock()

	for i, r := range t.resources {
		if r.Type == resourceType && r.Name == name && r.Namespace == namespace {
			t.resources = append(t.resources[:i], t.resources[i+1:]...)
			log.Printf("[TRACKER] Removed %s: %s from tracking", resourceType, name)
			return
		}
	}
}

// CleanupTracked deletes all tracked resources in the correct dependency order.
// This is the SAFE cleanup method that only removes resources created by the
// current test session.
func CleanupTracked() error {
	tracker := GetTracker()
	tracked := tracker.GetTracked()

	if len(tracked) == 0 {
		log.Printf("[CLEANUP] No tracked resources to clean up")
		return nil
	}

	log.Printf("[CLEANUP] Starting cleanup of %d tracked resources", len(tracked))

	c, err := GetSharedClient()
	if err != nil {
		return fmt.Errorf("error getting client for cleanup: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), SweeperTimeout)
	defer cancel()

	var errors []error

	// Delete in dependency order (reverse of creation order for most cases)
	// 1. First delete child resources (load balancers, etc.)
	// 2. Then delete parent resources (namespaces)

	// Order: Most dependent first, namespaces last
	deleteOrder := []string{
		"f5xc_http_loadbalancer",
		"f5xc_origin_pool",
		"f5xc_healthcheck",
		"f5xc_app_firewall",
		"f5xc_service_policy",
		"f5xc_ip_prefix_set",
		"f5xc_rate_limiter",
		"f5xc_user_identification",
		"f5xc_malicious_user_mitigation",
		"f5xc_namespace", // Always last
	}

	for _, resourceType := range deleteOrder {
		resources := tracker.GetTrackedByType(resourceType)
		for _, r := range resources {
			// Apply rate limiting before each cleanup operation
			WaitBeforeCleanup()

			if err := deleteTrackedResource(ctx, c, r); err != nil {
				log.Printf("[CLEANUP] Warning: failed to delete %s %s: %v", r.Type, r.Name, err)
				errors = append(errors, err)
			} else {
				tracker.Remove(r.Type, r.Name, r.Namespace)
			}
		}
	}

	// Handle any resource types not in the explicit order
	remaining := tracker.GetTracked()
	for _, r := range remaining {
		// Apply rate limiting before each cleanup operation
		WaitBeforeCleanup()

		if err := deleteTrackedResource(ctx, c, r); err != nil {
			log.Printf("[CLEANUP] Warning: failed to delete %s %s: %v", r.Type, r.Name, err)
			errors = append(errors, err)
		} else {
			tracker.Remove(r.Type, r.Name, r.Namespace)
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf("cleanup completed with %d errors", len(errors))
	}

	log.Printf("[CLEANUP] Successfully cleaned up all tracked resources")
	return nil
}

// deleteTrackedResource deletes a single tracked resource
func deleteTrackedResource(ctx context.Context, c *client.Client, r TrackedResource) error {
	deleteCtx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	log.Printf("[CLEANUP] Deleting %s: %s (namespace: %s)", r.Type, r.Name, r.Namespace)

	switch r.Type {
	case "f5xc_namespace":
		// Use cascade delete for namespaces (standard DELETE returns 501)
		return c.CascadeDeleteNamespace(deleteCtx, r.Name)
	case "f5xc_http_loadbalancer":
		return c.DeleteHTTPLoadBalancer(deleteCtx, r.Namespace, r.Name)
	case "f5xc_origin_pool":
		return c.DeleteOriginPool(deleteCtx, r.Namespace, r.Name)
	case "f5xc_healthcheck":
		return c.DeleteHealthcheck(deleteCtx, r.Namespace, r.Name)
	case "f5xc_app_firewall":
		return c.DeleteAppFirewall(deleteCtx, r.Namespace, r.Name)
	case "f5xc_service_policy":
		return c.DeleteServicePolicy(deleteCtx, r.Namespace, r.Name)
	case "f5xc_ip_prefix_set":
		return c.DeleteIPPrefixSet(deleteCtx, r.Namespace, r.Name)
	case "f5xc_rate_limiter":
		return c.DeleteRateLimiter(deleteCtx, r.Namespace, r.Name)
	case "f5xc_user_identification":
		return c.DeleteUserIdentification(deleteCtx, r.Namespace, r.Name)
	case "f5xc_malicious_user_mitigation":
		return c.DeleteMaliciousUserMitigation(deleteCtx, r.Namespace, r.Name)
	default:
		return fmt.Errorf("unknown resource type: %s", r.Type)
	}
}
