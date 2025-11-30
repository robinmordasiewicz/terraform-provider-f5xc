// Copyright (c) F5XC Community
// SPDX-License-Identifier: MPL-2.0

// Package acctest provides acceptance test utilities including test resource cleanup.
//
// # Resource Cleanup Approaches
//
// This package provides two complementary approaches to test resource cleanup:
//
// ## 1. Tracked Cleanup (RECOMMENDED for multi-user environments)
//
// The ResourceTracker tracks specific resources created during a test run and only
// cleans up those exact resources. This is SAFE for environments where multiple
// users may be running tests against the same F5 XC tenant.
//
// Usage in tests:
//
//	func TestAccMyResource(t *testing.T) {
//	    tracker := acctest.GetTracker()
//	    resourceName := acctest.RandomName("myresource")
//	    tracker.TrackResource("f5xc_my_resource", resourceName, namespace)
//	    defer acctest.CleanupTracked()  // Only deletes resources WE created
//	    // ... test logic ...
//	}
//
// ## 2. Prefix-Based Sweepers (for orphaned resource cleanup)
//
// Sweepers delete ALL resources matching test prefixes (tf-acc-test-*, tf-test-*).
// Use this ONLY when you need to clean up orphaned resources from crashed tests
// or when you're certain no other users are running tests.
//
// Usage:
//
//	TF_ACC=1 go test ./internal/acctest -v -sweep=all -timeout 30m
//
// Or via Makefile:
//
//	make sweep
//
// WARNING: Prefix-based sweepers will delete ANY resource matching the test prefix,
// including resources created by other users running tests against the same tenant.
// Use CleanupTracked() for safe multi-user cleanup.
package acctest

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/f5xc/terraform-provider-f5xc/internal/client"
)

const (
	// TestResourcePrefix is the standard prefix for all acceptance test resources.
	// Resources matching this prefix are candidates for cleanup by sweepers.
	TestResourcePrefix = "tf-acc-test-"

	// LegacyTestPrefix is the old prefix used by some tests.
	// Included for backward compatibility during transition.
	LegacyTestPrefix = "tf-test-"

	// SweeperTimeout is the maximum time allowed for a single sweeper operation.
	SweeperTimeout = 5 * time.Minute
)

// sharedClient is the API client used by all sweepers
var sharedClient *client.Client

// GetSharedClient returns a client instance for sweepers, creating it if necessary.
// This is separate from GetTestClient to avoid import cycles.
func GetSharedClient() (*client.Client, error) {
	if sharedClient != nil {
		return sharedClient, nil
	}

	apiURL := os.Getenv(EnvF5XCURL)
	if apiURL == "" {
		return nil, fmt.Errorf("%s must be set for sweepers", EnvF5XCURL)
	}

	// Normalize URL
	apiURL = strings.TrimRight(apiURL, "/")
	if strings.HasSuffix(strings.ToLower(apiURL), "/api") {
		apiURL = apiURL[:len(apiURL)-4]
	}

	switch DetectAuthMethod() {
	case AuthMethodP12:
		c, err := client.NewClientWithP12(
			apiURL,
			os.Getenv(EnvF5XCP12File),
			os.Getenv(EnvF5XCP12Password),
		)
		if err != nil {
			return nil, fmt.Errorf("failed to create P12 client: %w", err)
		}
		sharedClient = c
	case AuthMethodPEM:
		c, err := client.NewClientWithCert(
			apiURL,
			os.Getenv(EnvF5XCCert),
			os.Getenv(EnvF5XCKey),
			"",
		)
		if err != nil {
			return nil, fmt.Errorf("failed to create PEM client: %w", err)
		}
		sharedClient = c
	case AuthMethodToken:
		sharedClient = client.NewClient(apiURL, os.Getenv(EnvF5XCToken))
	default:
		return nil, fmt.Errorf("no authentication configured for sweepers")
	}

	return sharedClient, nil
}

// isTestResource checks if a resource name matches test resource naming patterns.
func isTestResource(name string) bool {
	return strings.HasPrefix(name, TestResourcePrefix) ||
		strings.HasPrefix(name, LegacyTestPrefix)
}

// init registers all sweepers when the package is loaded.
// Sweepers are only executed when running with -sweep flag.
func init() {
	// Register sweepers with dependency ordering
	// Namespace sweeper runs last since it can cascade delete child resources

	// Leaf resources (no dependencies on other test resources)
	resource.AddTestSweepers("f5xc_healthcheck", &resource.Sweeper{
		Name: "f5xc_healthcheck",
		F:    sweepHealthchecks,
	})

	resource.AddTestSweepers("f5xc_ip_prefix_set", &resource.Sweeper{
		Name: "f5xc_ip_prefix_set",
		F:    sweepIPPrefixSets,
	})

	resource.AddTestSweepers("f5xc_rate_limiter", &resource.Sweeper{
		Name: "f5xc_rate_limiter",
		F:    sweepRateLimiters,
	})

	resource.AddTestSweepers("f5xc_user_identification", &resource.Sweeper{
		Name: "f5xc_user_identification",
		F:    sweepUserIdentifications,
	})

	resource.AddTestSweepers("f5xc_malicious_user_mitigation", &resource.Sweeper{
		Name: "f5xc_malicious_user_mitigation",
		F:    sweepMaliciousUserMitigations,
	})

	resource.AddTestSweepers("f5xc_app_firewall", &resource.Sweeper{
		Name: "f5xc_app_firewall",
		F:    sweepAppFirewalls,
	})

	// Resources with dependencies
	resource.AddTestSweepers("f5xc_origin_pool", &resource.Sweeper{
		Name:         "f5xc_origin_pool",
		F:            sweepOriginPools,
		Dependencies: []string{"f5xc_http_loadbalancer"},
	})

	resource.AddTestSweepers("f5xc_http_loadbalancer", &resource.Sweeper{
		Name: "f5xc_http_loadbalancer",
		F:    sweepHTTPLoadbalancers,
	})

	resource.AddTestSweepers("f5xc_service_policy", &resource.Sweeper{
		Name: "f5xc_service_policy",
		F:    sweepServicePolicies,
	})

	// Namespace sweeper runs last - depends on all other sweepers
	// Deleting a namespace can cascade delete child resources
	resource.AddTestSweepers("f5xc_namespace", &resource.Sweeper{
		Name: "f5xc_namespace",
		F:    sweepNamespaces,
		Dependencies: []string{
			"f5xc_http_loadbalancer",
			"f5xc_origin_pool",
			"f5xc_healthcheck",
			"f5xc_app_firewall",
			"f5xc_service_policy",
			"f5xc_ip_prefix_set",
			"f5xc_rate_limiter",
			"f5xc_user_identification",
			"f5xc_malicious_user_mitigation",
		},
	})
}

// sweepNamespaces removes all test namespaces.
// This is the most important sweeper as deleting a namespace will cascade
// delete all resources within it.
func sweepNamespaces(_ string) error {
	c, err := GetSharedClient()
	if err != nil {
		return fmt.Errorf("error getting client: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), SweeperTimeout)
	defer cancel()

	log.Printf("[INFO] Sweeping namespaces with prefix %q or %q", TestResourcePrefix, LegacyTestPrefix)

	resp, err := c.ListNamespaces(ctx)
	if err != nil {
		return fmt.Errorf("error listing namespaces: %w", err)
	}

	var errs []string
	swept := 0

	for _, item := range resp.Items {
		name := item.Metadata.Name

		if !isTestResource(name) {
			continue
		}

		// Apply rate limiting before delete operation
		WaitBeforeCleanup()

		log.Printf("[INFO] Deleting namespace: %s", name)

		deleteCtx, deleteCancel := context.WithTimeout(ctx, 60*time.Second)
		// Use cascade delete for namespaces (standard DELETE returns 501)
		err := c.CascadeDeleteNamespace(deleteCtx, name)
		deleteCancel()

		if err != nil {
			// Check for "not found" errors which indicate already deleted
			if strings.Contains(err.Error(), "404") ||
				strings.Contains(err.Error(), "not found") ||
				strings.Contains(err.Error(), "NOT_FOUND") {
				log.Printf("[INFO] Namespace %s already deleted", name)
				continue
			}
			errs = append(errs, fmt.Sprintf("error deleting namespace %s: %v", name, err))
			continue
		}

		swept++
		log.Printf("[INFO] Successfully deleted namespace: %s", name)
	}

	log.Printf("[INFO] Swept %d namespaces", swept)

	if len(errs) > 0 {
		return fmt.Errorf("errors during namespace sweep:\n%s", strings.Join(errs, "\n"))
	}

	return nil
}

// sweepHTTPLoadbalancers removes all test HTTP load balancers.
func sweepHTTPLoadbalancers(_ string) error {
	c, err := GetSharedClient()
	if err != nil {
		return fmt.Errorf("error getting client: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), SweeperTimeout)
	defer cancel()

	log.Printf("[INFO] Sweeping HTTP load balancers")

	// Get list of test namespaces to search in
	namespaces, err := getTestNamespaces(ctx, c)
	if err != nil {
		return fmt.Errorf("error getting namespaces: %w", err)
	}

	// Also check system namespace
	namespaces = append(namespaces, "system")

	var errs []string
	swept := 0

	for _, ns := range namespaces {
		resp, err := c.ListHTTPLoadBalancers(ctx, ns)
		if err != nil {
			// Skip namespaces that don't exist or have errors
			continue
		}

		for _, item := range resp.Items {
			name := item.Metadata.Name

			if !isTestResource(name) {
				continue
			}

			// Apply rate limiting before delete operation
			WaitBeforeCleanup()

			log.Printf("[INFO] Deleting HTTP load balancer: %s/%s", ns, name)

			deleteCtx, deleteCancel := context.WithTimeout(ctx, 60*time.Second)
			err := c.DeleteHTTPLoadBalancer(deleteCtx, ns, name)
			deleteCancel()

			if err != nil {
				if isNotFoundError(err) {
					continue
				}
				errs = append(errs, fmt.Sprintf("error deleting http_loadbalancer %s/%s: %v", ns, name, err))
				continue
			}

			swept++
		}
	}

	log.Printf("[INFO] Swept %d HTTP load balancers", swept)

	if len(errs) > 0 {
		return fmt.Errorf("errors during HTTP load balancer sweep:\n%s", strings.Join(errs, "\n"))
	}

	return nil
}

// sweepOriginPools removes all test origin pools.
func sweepOriginPools(_ string) error {
	c, err := GetSharedClient()
	if err != nil {
		return fmt.Errorf("error getting client: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), SweeperTimeout)
	defer cancel()

	log.Printf("[INFO] Sweeping origin pools")

	namespaces, err := getTestNamespaces(ctx, c)
	if err != nil {
		return fmt.Errorf("error getting namespaces: %w", err)
	}
	namespaces = append(namespaces, "system")

	var errs []string
	swept := 0

	for _, ns := range namespaces {
		resp, err := c.ListOriginPools(ctx, ns)
		if err != nil {
			continue
		}

		for _, item := range resp.Items {
			name := item.Metadata.Name

			if !isTestResource(name) {
				continue
			}

			// Apply rate limiting before delete operation
			WaitBeforeCleanup()

			log.Printf("[INFO] Deleting origin pool: %s/%s", ns, name)

			deleteCtx, deleteCancel := context.WithTimeout(ctx, 60*time.Second)
			err := c.DeleteOriginPool(deleteCtx, ns, name)
			deleteCancel()

			if err != nil {
				if isNotFoundError(err) {
					continue
				}
				errs = append(errs, fmt.Sprintf("error deleting origin_pool %s/%s: %v", ns, name, err))
				continue
			}

			swept++
		}
	}

	log.Printf("[INFO] Swept %d origin pools", swept)

	if len(errs) > 0 {
		return fmt.Errorf("errors during origin pool sweep:\n%s", strings.Join(errs, "\n"))
	}

	return nil
}

// sweepHealthchecks removes all test healthchecks.
func sweepHealthchecks(_ string) error {
	c, err := GetSharedClient()
	if err != nil {
		return fmt.Errorf("error getting client: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), SweeperTimeout)
	defer cancel()

	log.Printf("[INFO] Sweeping healthchecks")

	namespaces, err := getTestNamespaces(ctx, c)
	if err != nil {
		return fmt.Errorf("error getting namespaces: %w", err)
	}
	namespaces = append(namespaces, "system")

	var errs []string
	swept := 0

	for _, ns := range namespaces {
		resp, err := c.ListHealthchecks(ctx, ns)
		if err != nil {
			continue
		}

		for _, item := range resp.Items {
			name := item.Metadata.Name

			if !isTestResource(name) {
				continue
			}

			// Apply rate limiting before delete operation
			WaitBeforeCleanup()

			log.Printf("[INFO] Deleting healthcheck: %s/%s", ns, name)

			deleteCtx, deleteCancel := context.WithTimeout(ctx, 60*time.Second)
			err := c.DeleteHealthcheck(deleteCtx, ns, name)
			deleteCancel()

			if err != nil {
				if isNotFoundError(err) {
					continue
				}
				errs = append(errs, fmt.Sprintf("error deleting healthcheck %s/%s: %v", ns, name, err))
				continue
			}

			swept++
		}
	}

	log.Printf("[INFO] Swept %d healthchecks", swept)

	if len(errs) > 0 {
		return fmt.Errorf("errors during healthcheck sweep:\n%s", strings.Join(errs, "\n"))
	}

	return nil
}

// sweepAppFirewalls removes all test app firewalls.
func sweepAppFirewalls(_ string) error {
	c, err := GetSharedClient()
	if err != nil {
		return fmt.Errorf("error getting client: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), SweeperTimeout)
	defer cancel()

	log.Printf("[INFO] Sweeping app firewalls")

	namespaces, err := getTestNamespaces(ctx, c)
	if err != nil {
		return fmt.Errorf("error getting namespaces: %w", err)
	}
	namespaces = append(namespaces, "system")

	var errs []string
	swept := 0

	for _, ns := range namespaces {
		resp, err := c.ListAppFirewalls(ctx, ns)
		if err != nil {
			continue
		}

		for _, item := range resp.Items {
			name := item.Metadata.Name

			if !isTestResource(name) {
				continue
			}

			// Apply rate limiting before delete operation
			WaitBeforeCleanup()

			log.Printf("[INFO] Deleting app firewall: %s/%s", ns, name)

			deleteCtx, deleteCancel := context.WithTimeout(ctx, 60*time.Second)
			err := c.DeleteAppFirewall(deleteCtx, ns, name)
			deleteCancel()

			if err != nil {
				if isNotFoundError(err) {
					continue
				}
				errs = append(errs, fmt.Sprintf("error deleting app_firewall %s/%s: %v", ns, name, err))
				continue
			}

			swept++
		}
	}

	log.Printf("[INFO] Swept %d app firewalls", swept)

	if len(errs) > 0 {
		return fmt.Errorf("errors during app firewall sweep:\n%s", strings.Join(errs, "\n"))
	}

	return nil
}

// sweepServicePolicies removes all test service policies.
func sweepServicePolicies(_ string) error {
	c, err := GetSharedClient()
	if err != nil {
		return fmt.Errorf("error getting client: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), SweeperTimeout)
	defer cancel()

	log.Printf("[INFO] Sweeping service policies")

	namespaces, err := getTestNamespaces(ctx, c)
	if err != nil {
		return fmt.Errorf("error getting namespaces: %w", err)
	}
	namespaces = append(namespaces, "system")

	var errs []string
	swept := 0

	for _, ns := range namespaces {
		resp, err := c.ListServicePolicies(ctx, ns)
		if err != nil {
			continue
		}

		for _, item := range resp.Items {
			name := item.Metadata.Name

			if !isTestResource(name) {
				continue
			}

			// Apply rate limiting before delete operation
			WaitBeforeCleanup()

			log.Printf("[INFO] Deleting service policy: %s/%s", ns, name)

			deleteCtx, deleteCancel := context.WithTimeout(ctx, 60*time.Second)
			err := c.DeleteServicePolicy(deleteCtx, ns, name)
			deleteCancel()

			if err != nil {
				if isNotFoundError(err) {
					continue
				}
				errs = append(errs, fmt.Sprintf("error deleting service_policy %s/%s: %v", ns, name, err))
				continue
			}

			swept++
		}
	}

	log.Printf("[INFO] Swept %d service policies", swept)

	if len(errs) > 0 {
		return fmt.Errorf("errors during service policy sweep:\n%s", strings.Join(errs, "\n"))
	}

	return nil
}

// sweepIPPrefixSets removes all test IP prefix sets.
func sweepIPPrefixSets(_ string) error {
	c, err := GetSharedClient()
	if err != nil {
		return fmt.Errorf("error getting client: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), SweeperTimeout)
	defer cancel()

	log.Printf("[INFO] Sweeping IP prefix sets")

	namespaces, err := getTestNamespaces(ctx, c)
	if err != nil {
		return fmt.Errorf("error getting namespaces: %w", err)
	}
	namespaces = append(namespaces, "system")

	var errs []string
	swept := 0

	for _, ns := range namespaces {
		resp, err := c.ListIPPrefixSets(ctx, ns)
		if err != nil {
			continue
		}

		for _, item := range resp.Items {
			name := item.Metadata.Name

			if !isTestResource(name) {
				continue
			}

			// Apply rate limiting before delete operation
			WaitBeforeCleanup()

			log.Printf("[INFO] Deleting IP prefix set: %s/%s", ns, name)

			deleteCtx, deleteCancel := context.WithTimeout(ctx, 60*time.Second)
			err := c.DeleteIPPrefixSet(deleteCtx, ns, name)
			deleteCancel()

			if err != nil {
				if isNotFoundError(err) {
					continue
				}
				errs = append(errs, fmt.Sprintf("error deleting ip_prefix_set %s/%s: %v", ns, name, err))
				continue
			}

			swept++
		}
	}

	log.Printf("[INFO] Swept %d IP prefix sets", swept)

	if len(errs) > 0 {
		return fmt.Errorf("errors during IP prefix set sweep:\n%s", strings.Join(errs, "\n"))
	}

	return nil
}

// sweepRateLimiters removes all test rate limiters.
func sweepRateLimiters(_ string) error {
	c, err := GetSharedClient()
	if err != nil {
		return fmt.Errorf("error getting client: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), SweeperTimeout)
	defer cancel()

	log.Printf("[INFO] Sweeping rate limiters")

	namespaces, err := getTestNamespaces(ctx, c)
	if err != nil {
		return fmt.Errorf("error getting namespaces: %w", err)
	}
	namespaces = append(namespaces, "system")

	var errs []string
	swept := 0

	for _, ns := range namespaces {
		resp, err := c.ListRateLimiters(ctx, ns)
		if err != nil {
			continue
		}

		for _, item := range resp.Items {
			name := item.Metadata.Name

			if !isTestResource(name) {
				continue
			}

			// Apply rate limiting before delete operation
			WaitBeforeCleanup()

			log.Printf("[INFO] Deleting rate limiter: %s/%s", ns, name)

			deleteCtx, deleteCancel := context.WithTimeout(ctx, 60*time.Second)
			err := c.DeleteRateLimiter(deleteCtx, ns, name)
			deleteCancel()

			if err != nil {
				if isNotFoundError(err) {
					continue
				}
				errs = append(errs, fmt.Sprintf("error deleting rate_limiter %s/%s: %v", ns, name, err))
				continue
			}

			swept++
		}
	}

	log.Printf("[INFO] Swept %d rate limiters", swept)

	if len(errs) > 0 {
		return fmt.Errorf("errors during rate limiter sweep:\n%s", strings.Join(errs, "\n"))
	}

	return nil
}

// sweepUserIdentifications removes all test user identifications.
func sweepUserIdentifications(_ string) error {
	c, err := GetSharedClient()
	if err != nil {
		return fmt.Errorf("error getting client: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), SweeperTimeout)
	defer cancel()

	log.Printf("[INFO] Sweeping user identifications")

	namespaces, err := getTestNamespaces(ctx, c)
	if err != nil {
		return fmt.Errorf("error getting namespaces: %w", err)
	}
	namespaces = append(namespaces, "system")

	var errs []string
	swept := 0

	for _, ns := range namespaces {
		resp, err := c.ListUserIdentifications(ctx, ns)
		if err != nil {
			continue
		}

		for _, item := range resp.Items {
			name := item.Metadata.Name

			if !isTestResource(name) {
				continue
			}

			// Apply rate limiting before delete operation
			WaitBeforeCleanup()

			log.Printf("[INFO] Deleting user identification: %s/%s", ns, name)

			deleteCtx, deleteCancel := context.WithTimeout(ctx, 60*time.Second)
			err := c.DeleteUserIdentification(deleteCtx, ns, name)
			deleteCancel()

			if err != nil {
				if isNotFoundError(err) {
					continue
				}
				errs = append(errs, fmt.Sprintf("error deleting user_identification %s/%s: %v", ns, name, err))
				continue
			}

			swept++
		}
	}

	log.Printf("[INFO] Swept %d user identifications", swept)

	if len(errs) > 0 {
		return fmt.Errorf("errors during user identification sweep:\n%s", strings.Join(errs, "\n"))
	}

	return nil
}

// sweepMaliciousUserMitigations removes all test malicious user mitigations.
func sweepMaliciousUserMitigations(_ string) error {
	c, err := GetSharedClient()
	if err != nil {
		return fmt.Errorf("error getting client: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), SweeperTimeout)
	defer cancel()

	log.Printf("[INFO] Sweeping malicious user mitigations")

	namespaces, err := getTestNamespaces(ctx, c)
	if err != nil {
		return fmt.Errorf("error getting namespaces: %w", err)
	}
	namespaces = append(namespaces, "system")

	var errs []string
	swept := 0

	for _, ns := range namespaces {
		resp, err := c.ListMaliciousUserMitigations(ctx, ns)
		if err != nil {
			continue
		}

		for _, item := range resp.Items {
			name := item.Metadata.Name

			if !isTestResource(name) {
				continue
			}

			// Apply rate limiting before delete operation
			WaitBeforeCleanup()

			log.Printf("[INFO] Deleting malicious user mitigation: %s/%s", ns, name)

			deleteCtx, deleteCancel := context.WithTimeout(ctx, 60*time.Second)
			err := c.DeleteMaliciousUserMitigation(deleteCtx, ns, name)
			deleteCancel()

			if err != nil {
				if isNotFoundError(err) {
					continue
				}
				errs = append(errs, fmt.Sprintf("error deleting malicious_user_mitigation %s/%s: %v", ns, name, err))
				continue
			}

			swept++
		}
	}

	log.Printf("[INFO] Swept %d malicious user mitigations", swept)

	if len(errs) > 0 {
		return fmt.Errorf("errors during malicious user mitigation sweep:\n%s", strings.Join(errs, "\n"))
	}

	return nil
}

// getTestNamespaces returns a list of namespace names that match test patterns.
func getTestNamespaces(ctx context.Context, c *client.Client) ([]string, error) {
	resp, err := c.ListNamespaces(ctx)
	if err != nil {
		return nil, err
	}

	var namespaces []string
	for _, item := range resp.Items {
		if isTestResource(item.Metadata.Name) {
			namespaces = append(namespaces, item.Metadata.Name)
		}
	}

	return namespaces, nil
}

// isNotFoundError checks if an error indicates the resource was not found.
func isNotFoundError(err error) bool {
	if err == nil {
		return false
	}
	errStr := err.Error()
	return strings.Contains(errStr, "404") ||
		strings.Contains(errStr, "not found") ||
		strings.Contains(errStr, "NOT_FOUND")
}
