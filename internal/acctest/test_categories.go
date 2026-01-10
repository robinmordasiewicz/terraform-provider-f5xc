// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package acctest

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
)

// TestCategory represents the type of test being run
type TestCategory string

const (
	// TestCategoryReal indicates tests running against the real F5 XC API
	TestCategoryReal TestCategory = "REAL_API"

	// TestCategoryMock indicates tests running against the mock server
	TestCategoryMock TestCategory = "MOCK_API"

	// TestCategoryUnit indicates unit tests (no external dependencies)
	TestCategoryUnit TestCategory = "UNIT"
)

// TestCategoryMarkerPrefix is the prefix used in log output for machine parsing
const TestCategoryMarkerPrefix = "[TEST_CATEGORY:"

// TestResult represents a structured test result for reporting
type TestResult struct {
	Name       string       `json:"name"`
	Category   TestCategory `json:"category"`
	Passed     bool         `json:"passed"`
	Duration   float64      `json:"duration_seconds"`
	SkipReason string       `json:"skip_reason,omitempty"`
}

// TestSummary provides aggregated test results by category
type TestSummary struct {
	Timestamp    time.Time                         `json:"timestamp"`
	TotalTests   int                               `json:"total_tests"`
	TotalPassed  int                               `json:"total_passed"`
	TotalFailed  int                               `json:"total_failed"`
	TotalSkipped int                               `json:"total_skipped"`
	ByCategory   map[TestCategory]*CategorySummary `json:"by_category"`
}

// CategorySummary provides test counts for a single category
type CategorySummary struct {
	Category TestCategory `json:"category"`
	Total    int          `json:"total"`
	Passed   int          `json:"passed"`
	Failed   int          `json:"failed"`
	Skipped  int          `json:"skipped"`
}

// LogTestCategory logs the test category for machine parsing and reporting.
// This should be called at the beginning of each test function.
//
// Usage:
//
//	func TestAccMyResource_basic(t *testing.T) {
//	    acctest.LogTestCategory(t, acctest.TestCategoryReal)
//	    // ... rest of test
//	}
//
//	func TestMockMyResource_basic(t *testing.T) {
//	    acctest.LogTestCategory(t, acctest.TestCategoryMock)
//	    // ... rest of test
//	}
func LogTestCategory(t *testing.T, category TestCategory) {
	t.Helper()
	t.Logf("%s%s]", TestCategoryMarkerPrefix, category)
}

// InferTestCategory determines the test category from the test name.
// This is useful for reporting tools that parse test output.
//
// Naming conventions:
//   - TestAcc* -> TestCategoryReal (real API acceptance tests)
//   - TestMock* -> TestCategoryMock (mock server tests)
//   - Test* (other) -> TestCategoryUnit (unit tests)
func InferTestCategory(testName string) TestCategory {
	switch {
	case strings.HasPrefix(testName, "TestAcc"):
		return TestCategoryReal
	case strings.HasPrefix(testName, "TestMock"):
		return TestCategoryMock
	default:
		return TestCategoryUnit
	}
}

// IsRealAPITest returns true if the test name indicates a real API test
func IsRealAPITest(testName string) bool {
	return strings.HasPrefix(testName, "TestAcc")
}

// IsMockAPITest returns true if the test name indicates a mock API test
func IsMockAPITest(testName string) bool {
	return strings.HasPrefix(testName, "TestMock")
}

// GetTestCategorySummary provides a formatted summary string for test reporting
func GetTestCategorySummary(summary *TestSummary) string {
	var sb strings.Builder

	divider := strings.Repeat("=", 70)
	subDivider := strings.Repeat("-", 70)

	sb.WriteString("\n")
	sb.WriteString(divider + "\n")
	sb.WriteString("                       TEST SUMMARY REPORT\n")
	sb.WriteString(divider + "\n")
	sb.WriteString(fmt.Sprintf("Timestamp: %s\n", summary.Timestamp.Format(time.RFC3339)))
	sb.WriteString(subDivider + "\n")

	// Overall totals
	sb.WriteString(fmt.Sprintf("TOTAL: %d tests | %d passed | %d failed | %d skipped\n",
		summary.TotalTests, summary.TotalPassed, summary.TotalFailed, summary.TotalSkipped))
	sb.WriteString(subDivider + "\n")

	// By category
	for _, cat := range []TestCategory{TestCategoryReal, TestCategoryMock, TestCategoryUnit} {
		if cs, ok := summary.ByCategory[cat]; ok && cs.Total > 0 {
			status := "PASS"
			if cs.Failed > 0 {
				status = "FAIL"
			}
			sb.WriteString(fmt.Sprintf("[%s] %-10s: %3d tests | %3d passed | %3d failed | %3d skipped\n",
				status, cat, cs.Total, cs.Passed, cs.Failed, cs.Skipped))
		}
	}

	sb.WriteString(divider + "\n")

	return sb.String()
}

// WriteTestResultJSON writes test results to a JSON file for further processing
func WriteTestResultJSON(filename string, summary *TestSummary) error {
	data, err := json.MarshalIndent(summary, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal test summary: %w", err)
	}

	if err := os.WriteFile(filename, data, 0644); err != nil {
		return fmt.Errorf("failed to write test summary file: %w", err)
	}

	return nil
}

// NewTestSummary creates a new empty TestSummary
func NewTestSummary() *TestSummary {
	return &TestSummary{
		Timestamp:  time.Now(),
		ByCategory: make(map[TestCategory]*CategorySummary),
	}
}

// AddResult adds a test result to the summary
func (s *TestSummary) AddResult(result TestResult) {
	s.TotalTests++
	if result.Passed {
		s.TotalPassed++
	} else if result.SkipReason != "" {
		s.TotalSkipped++
	} else {
		s.TotalFailed++
	}

	if _, ok := s.ByCategory[result.Category]; !ok {
		s.ByCategory[result.Category] = &CategorySummary{
			Category: result.Category,
		}
	}

	cs := s.ByCategory[result.Category]
	cs.Total++
	if result.Passed {
		cs.Passed++
	} else if result.SkipReason != "" {
		cs.Skipped++
	} else {
		cs.Failed++
	}
}

// Environment variable for test category reporting
const (
	// EnvTestReportFile specifies the output file for JSON test reports
	EnvTestReportFile = "TEST_REPORT_FILE"

	// EnvTestReportFormat specifies the output format (json, text, markdown)
	EnvTestReportFormat = "TEST_REPORT_FORMAT"
)
