// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

// Package acctest provides acceptance test utilities for F5 XC Terraform provider.
package acctest

import (
	"context"
	"testing"
	"time"
)

// DefaultQuotaCheckTimeout is the timeout for quota check API calls
const DefaultQuotaCheckTimeout = 30 * time.Second

// SkipIfQuotaExhausted skips the test if the specified resource type quota is exhausted
// in the given namespace. This allows tests to skip gracefully instead of failing
// when quota limits are reached.
//
// Example usage:
//
//	func TestAccHealthcheckResource_basic(t *testing.T) {
//	    acctest.SkipIfNotAccTest(t)
//	    acctest.PreCheck(t)
//	    acctest.SkipIfQuotaExhausted(t, "system", "healthcheck")
//	    // ... rest of test
//	}
func SkipIfQuotaExhausted(t *testing.T, namespace, resourceType string) {
	t.Helper()

	// Skip quota check in mock mode - mock server won't have real quota data
	if IsMockMode() {
		return
	}

	client, err := GetTestClient()
	if err != nil {
		t.Logf("Warning: Cannot check quota (client error): %v", err)
		return // Don't skip if we can't check - let the test run and potentially fail naturally
	}

	ctx, cancel := context.WithTimeout(context.Background(), DefaultQuotaCheckTimeout)
	defer cancel()

	info, err := client.GetQuotaInfo(ctx, namespace, resourceType)
	if err != nil {
		t.Logf("Warning: Cannot check quota for %s/%s: %v", namespace, resourceType, err)
		return // Don't skip if we can't check
	}

	if info.Available <= 0 {
		t.Skipf("Skipping: %s quota exhausted in namespace %s (%d/%d used)",
			resourceType, namespace, info.Used, info.Limit)
	}

	t.Logf("Quota check passed for %s/%s: %d/%d used (%d available)",
		namespace, resourceType, info.Used, info.Limit, info.Available)
}

// SkipIfQuotaNearLimit skips the test if the resource type usage exceeds
// the specified threshold percentage (0-100) of the limit.
// This is useful for tests that create multiple resources.
//
// Example usage:
//
//	func TestAccHealthcheckResource_multiple(t *testing.T) {
//	    acctest.SkipIfNotAccTest(t)
//	    acctest.PreCheck(t)
//	    // Skip if > 90% quota used (need headroom for multiple resources)
//	    acctest.SkipIfQuotaNearLimit(t, "system", "healthcheck", 90)
//	    // ... rest of test
//	}
func SkipIfQuotaNearLimit(t *testing.T, namespace, resourceType string, thresholdPercent int) {
	t.Helper()

	// Skip quota check in mock mode
	if IsMockMode() {
		return
	}

	client, err := GetTestClient()
	if err != nil {
		t.Logf("Warning: Cannot check quota (client error): %v", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), DefaultQuotaCheckTimeout)
	defer cancel()

	info, err := client.GetQuotaInfo(ctx, namespace, resourceType)
	if err != nil {
		t.Logf("Warning: Cannot check quota for %s/%s: %v", namespace, resourceType, err)
		return
	}

	usagePercent := 0
	if info.Limit > 0 {
		usagePercent = (info.Used * 100) / info.Limit
	}

	if usagePercent >= thresholdPercent {
		t.Skipf("Skipping: %s quota near limit in namespace %s (%d%% used, threshold: %d%%)",
			resourceType, namespace, usagePercent, thresholdPercent)
	}

	t.Logf("Quota check passed for %s/%s: %d%% used (threshold: %d%%)",
		namespace, resourceType, usagePercent, thresholdPercent)
}

// SkipIfInsufficientQuota skips the test if there are fewer than the required
// number of quota slots available. Useful for tests that need to create
// a specific number of resources.
//
// Example usage:
//
//	func TestAccHealthcheckResource_createMany(t *testing.T) {
//	    acctest.SkipIfNotAccTest(t)
//	    acctest.PreCheck(t)
//	    // Skip if we can't create at least 5 healthchecks
//	    acctest.SkipIfInsufficientQuota(t, "system", "healthcheck", 5)
//	    // ... rest of test
//	}
func SkipIfInsufficientQuota(t *testing.T, namespace, resourceType string, required int) {
	t.Helper()

	// Skip quota check in mock mode
	if IsMockMode() {
		return
	}

	client, err := GetTestClient()
	if err != nil {
		t.Logf("Warning: Cannot check quota (client error): %v", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), DefaultQuotaCheckTimeout)
	defer cancel()

	info, err := client.GetQuotaInfo(ctx, namespace, resourceType)
	if err != nil {
		t.Logf("Warning: Cannot check quota for %s/%s: %v", namespace, resourceType, err)
		return
	}

	if info.Available < required {
		t.Skipf("Skipping: insufficient %s quota in namespace %s (need %d, have %d available)",
			resourceType, namespace, required, info.Available)
	}

	t.Logf("Quota check passed for %s/%s: %d available (need %d)",
		namespace, resourceType, info.Available, required)
}

// LogQuotaUsage logs the current quota usage for a resource type without skipping.
// Useful for debugging or informational purposes in test output.
func LogQuotaUsage(t *testing.T, namespace, resourceType string) {
	t.Helper()

	// Skip in mock mode
	if IsMockMode() {
		t.Log("Quota check skipped (mock mode)")
		return
	}

	client, err := GetTestClient()
	if err != nil {
		t.Logf("Warning: Cannot check quota (client error): %v", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), DefaultQuotaCheckTimeout)
	defer cancel()

	info, err := client.GetQuotaInfo(ctx, namespace, resourceType)
	if err != nil {
		t.Logf("Warning: Cannot check quota for %s/%s: %v", namespace, resourceType, err)
		return
	}

	usagePercent := 0
	if info.Limit > 0 {
		usagePercent = (info.Used * 100) / info.Limit
	}

	t.Logf("Quota status for %s/%s: %d/%d used (%d%%), %d available",
		namespace, resourceType, info.Used, info.Limit, usagePercent, info.Available)
}
