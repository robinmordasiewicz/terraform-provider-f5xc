// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

// Package acctest provides acceptance test utilities including rate limiting
// to prevent exceeding F5 XC API rate limits during test execution.
package acctest

import (
	"log"
	"sync"
	"time"
)

// Rate limiting configuration for F5 XC API
// Based on observed API behavior and best practices for Terraform acceptance testing
const (
	// DefaultOperationDelay is the minimum delay between API operations
	// This helps prevent hitting rate limits during rapid test execution
	DefaultOperationDelay = 500 * time.Millisecond

	// DefaultTestDelay is the delay between test steps
	// Allows the API to process operations before verification
	DefaultTestDelay = 1 * time.Second

	// DefaultCleanupDelay is the delay between cleanup operations
	// Prevents overwhelming the API during test teardown
	DefaultCleanupDelay = 2 * time.Second

	// DefaultBurstLimit is the maximum number of rapid operations allowed
	// before enforcing a longer delay
	DefaultBurstLimit = 5

	// DefaultBurstDelay is the delay enforced after hitting the burst limit
	DefaultBurstDelay = 5 * time.Second

	// MaxParallelTests is the recommended maximum number of parallel tests
	// to avoid overwhelming the F5 XC API
	MaxParallelTests = 3
)

// RateLimiter provides coordinated rate limiting for acceptance tests
// to prevent exceeding F5 XC API rate limits
type RateLimiter struct {
	mu             sync.Mutex
	operationDelay time.Duration
	testDelay      time.Duration
	cleanupDelay   time.Duration
	burstLimit     int
	burstDelay     time.Duration
	operationCount int
	lastOperation  time.Time
	enabled        bool
}

// Global rate limiter instance
var (
	globalRateLimiter     *RateLimiter
	globalRateLimiterOnce sync.Once
)

// GetRateLimiter returns the global rate limiter instance
func GetRateLimiter() *RateLimiter {
	globalRateLimiterOnce.Do(func() {
		globalRateLimiter = &RateLimiter{
			operationDelay: DefaultOperationDelay,
			testDelay:      DefaultTestDelay,
			cleanupDelay:   DefaultCleanupDelay,
			burstLimit:     DefaultBurstLimit,
			burstDelay:     DefaultBurstDelay,
			enabled:        true,
		}
	})
	return globalRateLimiter
}

// SetEnabled enables or disables rate limiting
func (r *RateLimiter) SetEnabled(enabled bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.enabled = enabled
	if enabled {
		log.Printf("[RATE_LIMIT] Rate limiting enabled")
	} else {
		log.Printf("[RATE_LIMIT] Rate limiting disabled")
	}
}

// SetOperationDelay sets the delay between operations
func (r *RateLimiter) SetOperationDelay(d time.Duration) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.operationDelay = d
}

// SetTestDelay sets the delay between test steps
func (r *RateLimiter) SetTestDelay(d time.Duration) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.testDelay = d
}

// SetCleanupDelay sets the delay between cleanup operations
func (r *RateLimiter) SetCleanupDelay(d time.Duration) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.cleanupDelay = d
}

// SetBurstLimit sets the burst limit before enforcing longer delays
func (r *RateLimiter) SetBurstLimit(limit int) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.burstLimit = limit
}

// WaitForOperation should be called before each API operation
// It enforces rate limiting to prevent exceeding API limits
func (r *RateLimiter) WaitForOperation() {
	r.mu.Lock()
	defer r.mu.Unlock()

	if !r.enabled {
		return
	}

	// Check if we've hit the burst limit
	r.operationCount++
	if r.operationCount >= r.burstLimit {
		log.Printf("[RATE_LIMIT] Burst limit reached (%d operations), waiting %v",
			r.operationCount, r.burstDelay)
		time.Sleep(r.burstDelay)
		r.operationCount = 0
		r.lastOperation = time.Now()
		return
	}

	// Enforce minimum delay between operations
	elapsed := time.Since(r.lastOperation)
	if elapsed < r.operationDelay {
		sleepTime := r.operationDelay - elapsed
		time.Sleep(sleepTime)
	}
	r.lastOperation = time.Now()
}

// WaitForTestStep should be called between test steps
func (r *RateLimiter) WaitForTestStep() {
	r.mu.Lock()
	defer r.mu.Unlock()

	if !r.enabled {
		return
	}

	time.Sleep(r.testDelay)
	r.lastOperation = time.Now()
}

// WaitForCleanup should be called between cleanup operations
func (r *RateLimiter) WaitForCleanup() {
	r.mu.Lock()
	defer r.mu.Unlock()

	if !r.enabled {
		return
	}

	time.Sleep(r.cleanupDelay)
	r.lastOperation = time.Now()
}

// ResetBurstCounter resets the operation burst counter
// Call this at the start of each test to reset the counter
func (r *RateLimiter) ResetBurstCounter() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.operationCount = 0
}

// GetStats returns current rate limiter statistics
func (r *RateLimiter) GetStats() map[string]interface{} {
	r.mu.Lock()
	defer r.mu.Unlock()

	return map[string]interface{}{
		"enabled":         r.enabled,
		"operation_count": r.operationCount,
		"operation_delay": r.operationDelay,
		"test_delay":      r.testDelay,
		"cleanup_delay":   r.cleanupDelay,
		"burst_limit":     r.burstLimit,
		"burst_delay":     r.burstDelay,
	}
}

// WaitBeforeOperation is a convenience function to wait before an API operation
// using the global rate limiter
func WaitBeforeOperation() {
	GetRateLimiter().WaitForOperation()
}

// WaitBeforeTestStep is a convenience function to wait before a test step
// using the global rate limiter
func WaitBeforeTestStep() {
	GetRateLimiter().WaitForTestStep()
}

// WaitBeforeCleanup is a convenience function to wait before cleanup
// using the global rate limiter
func WaitBeforeCleanup() {
	GetRateLimiter().WaitForCleanup()
}

// ConfigureRateLimiting sets up rate limiting with custom parameters
// Call this at the start of your test suite to configure rate limiting
func ConfigureRateLimiting(operationDelay, testDelay, cleanupDelay time.Duration, burstLimit int) {
	rl := GetRateLimiter()
	if operationDelay > 0 {
		rl.SetOperationDelay(operationDelay)
	}
	if testDelay > 0 {
		rl.SetTestDelay(testDelay)
	}
	if cleanupDelay > 0 {
		rl.SetCleanupDelay(cleanupDelay)
	}
	if burstLimit > 0 {
		rl.SetBurstLimit(burstLimit)
	}
	log.Printf("[RATE_LIMIT] Configured: operation=%v, test=%v, cleanup=%v, burst=%d",
		operationDelay, testDelay, cleanupDelay, burstLimit)
}

// DisableRateLimiting disables rate limiting (useful for debugging)
func DisableRateLimiting() {
	GetRateLimiter().SetEnabled(false)
}

// EnableRateLimiting enables rate limiting
func EnableRateLimiting() {
	GetRateLimiter().SetEnabled(true)
}
