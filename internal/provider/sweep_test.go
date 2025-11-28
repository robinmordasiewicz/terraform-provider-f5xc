// Copyright (c) F5XC Community
// SPDX-License-Identifier: MPL-2.0

package provider_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/f5xc/terraform-provider-f5xc/internal/acctest"
)

// =============================================================================
// TEST SWEEPERS
//
// Sweepers clean up orphaned test resources that may have been left behind
// from failed or interrupted test runs. They are NOT run during normal test
// execution - they must be explicitly invoked.
//
// Run sweepers with:
//   go test ./internal/provider/ -v -sweep=all -sweep-allow-failures
//
// Or for a specific sweeper:
//   go test ./internal/provider/ -v -sweep=all -sweep-run=f5xc_namespace
//
// IMPORTANT: Sweepers only delete resources with the "tf-acc-test-" prefix
// to avoid accidentally deleting production resources.
// =============================================================================

func init() {
	resource.AddTestSweepers("f5xc_namespace", &resource.Sweeper{
		Name: "f5xc_namespace",
		F:    sweepNamespaces,
	})

	resource.AddTestSweepers("f5xc_healthcheck", &resource.Sweeper{
		Name:         "f5xc_healthcheck",
		F:            sweepHealthchecks,
		Dependencies: []string{"f5xc_namespace"}, // Sweep healthchecks before namespaces
	})
}

// sweepNamespaces deletes orphaned test namespaces from the F5 XC API.
// Only namespaces with the "tf-acc-test-" prefix are deleted.
func sweepNamespaces(region string) error {
	// region parameter is not used for F5 XC (it's cloud-agnostic)
	// but is required by the sweeper interface
	_ = region

	// Check if required environment variables are set
	apiURL := os.Getenv("F5XC_API_URL")
	if apiURL == "" {
		log.Println("[WARN] F5XC_API_URL not set, skipping namespace sweeper")
		return nil
	}

	// Get test client
	client, err := acctest.GetTestClient()
	if err != nil {
		return fmt.Errorf("error getting test client for sweeper: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	// List all namespaces
	namespaces, err := client.ListNamespaces(ctx, "system")
	if err != nil {
		return fmt.Errorf("error listing namespaces for sweeper: %w", err)
	}

	var sweepErrors []string
	var sweptCount int

	for _, ns := range namespaces {
		// Only delete test resources (those with our test prefix)
		if !strings.HasPrefix(ns.Metadata.Name, "tf-acc-test-") {
			continue
		}

		log.Printf("[INFO] Sweeping namespace: %s", ns.Metadata.Name)

		err := client.DeleteNamespace(ctx, "system", ns.Metadata.Name)
		if err != nil {
			// Check if already deleted
			if strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "not found") {
				log.Printf("[DEBUG] Namespace %s already deleted", ns.Metadata.Name)
				continue
			}
			sweepErrors = append(sweepErrors, fmt.Sprintf("error deleting namespace %s: %s", ns.Metadata.Name, err))
			continue
		}

		sweptCount++
		log.Printf("[INFO] Successfully swept namespace: %s", ns.Metadata.Name)
	}

	log.Printf("[INFO] Namespace sweeper completed: %d namespaces swept", sweptCount)

	if len(sweepErrors) > 0 {
		return fmt.Errorf("sweeper encountered errors:\n%s", strings.Join(sweepErrors, "\n"))
	}

	return nil
}

// sweepHealthchecks deletes orphaned test healthchecks from the F5 XC API.
// Only healthchecks with the "tf-acc-test-" prefix are deleted.
// This sweeper runs across all test namespaces.
func sweepHealthchecks(region string) error {
	// region parameter is not used for F5 XC (it's cloud-agnostic)
	// but is required by the sweeper interface
	_ = region

	// Check if required environment variables are set
	apiURL := os.Getenv("F5XC_API_URL")
	if apiURL == "" {
		log.Println("[WARN] F5XC_API_URL not set, skipping healthcheck sweeper")
		return nil
	}

	// Get test client
	client, err := acctest.GetTestClient()
	if err != nil {
		return fmt.Errorf("error getting test client for sweeper: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	// First, get all test namespaces to search for healthchecks
	namespaces, err := client.ListNamespaces(ctx, "system")
	if err != nil {
		return fmt.Errorf("error listing namespaces for healthcheck sweeper: %w", err)
	}

	var sweepErrors []string
	var sweptCount int

	// Sweep healthchecks from each test namespace
	for _, ns := range namespaces {
		// Only check namespaces that might contain test resources
		if !strings.HasPrefix(ns.Metadata.Name, "tf-acc-test-") {
			continue
		}

		// List healthchecks in this namespace
		healthchecks, err := client.ListHealthchecks(ctx, ns.Metadata.Name)
		if err != nil {
			// Namespace might be deleted or inaccessible
			if strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "not found") {
				continue
			}
			log.Printf("[WARN] Error listing healthchecks in namespace %s: %v", ns.Metadata.Name, err)
			continue
		}

		for _, hc := range healthchecks {
			// Only delete test resources (those with our test prefix)
			if !strings.HasPrefix(hc.Metadata.Name, "tf-acc-test-") {
				continue
			}

			log.Printf("[INFO] Sweeping healthcheck: %s/%s", ns.Metadata.Name, hc.Metadata.Name)

			err := client.DeleteHealthcheck(ctx, ns.Metadata.Name, hc.Metadata.Name)
			if err != nil {
				// Check if already deleted
				if strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "not found") {
					log.Printf("[DEBUG] Healthcheck %s already deleted", hc.Metadata.Name)
					continue
				}
				sweepErrors = append(sweepErrors, fmt.Sprintf("error deleting healthcheck %s/%s: %s", ns.Metadata.Name, hc.Metadata.Name, err))
				continue
			}

			sweptCount++
			log.Printf("[INFO] Successfully swept healthcheck: %s/%s", ns.Metadata.Name, hc.Metadata.Name)
		}
	}

	log.Printf("[INFO] Healthcheck sweeper completed: %d healthchecks swept", sweptCount)

	if len(sweepErrors) > 0 {
		return fmt.Errorf("sweeper encountered errors:\n%s", strings.Join(sweepErrors, "\n"))
	}

	return nil
}
