// Copyright (c) F5XC Community
// SPDX-License-Identifier: MPL-2.0

// compare-test-results compares mock and real API test results to identify discrepancies.
//
// Usage:
//
//	go run tools/compare-test-results/main.go \
//	  -mock=mock-tests.json \
//	  -real=real-tests.json \
//	  -output=comparison.md
//
// The tool identifies:
//   - Tests that PASS on real but FAIL on mock → Mock needs fixing
//   - Tests that FAIL on both → Resource/test issue (not mock problem)
//   - Tests that PASS on mock but FAIL on real → Mock is incorrect
//   - Tests that PASS on both → Mock is consistent (goal: 100%)
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

// TestResult represents a single test result
type TestResult struct {
	Name           string `json:"name"`
	Package        string `json:"package"`
	Category       string `json:"category"`
	Status         string `json:"status"` // pass, fail, skip
	Duration       float64 `json:"duration_seconds"`
	SkipReason     string `json:"skip_reason,omitempty"`
	TransientError string `json:"transient_error,omitempty"`
	FailureOutput  string `json:"failure_output,omitempty"`
}

// Report represents a test report from test-report tool
type Report struct {
	Timestamp    time.Time    `json:"timestamp"`
	TotalTests   int          `json:"total_tests"`
	TotalPassed  int          `json:"total_passed"`
	TotalFailed  int          `json:"total_failed"`
	TotalSkipped int          `json:"total_skipped"`
	Tests        []TestResult `json:"tests"`
}

// ComparisonCategory classifies the comparison result
type ComparisonCategory string

const (
	// MockConsistent - Test passes on both mock and real API (GOAL: 100%)
	MockConsistent ComparisonCategory = "CONSISTENT"
	// MockNeedsFix - Test passes on real but fails on mock (FIX MOCK)
	MockNeedsFix ComparisonCategory = "MOCK_NEEDS_FIX"
	// TestNeedsFix - Test fails on both real and mock (FIX TEST)
	TestNeedsFix ComparisonCategory = "TEST_NEEDS_FIX"
	// MockIncorrect - Test passes on mock but fails on real (MOCK INCORRECT)
	MockIncorrect ComparisonCategory = "MOCK_INCORRECT"
	// MockOnlyPass - Test only run on mock (skipped on real)
	MockOnlyPass ComparisonCategory = "MOCK_ONLY_PASS"
	// RealOnlyPass - Test only run on real (not in mock)
	RealOnlyPass ComparisonCategory = "REAL_ONLY_PASS"
	// BothSkipped - Test skipped on both
	BothSkipped ComparisonCategory = "BOTH_SKIPPED"
)

// ComparisonResult holds the comparison for a single test
type ComparisonResult struct {
	TestName     string             `json:"test_name"`
	Category     ComparisonCategory `json:"category"`
	RealStatus   string             `json:"real_status"`
	MockStatus   string             `json:"mock_status"`
	RealDuration float64            `json:"real_duration_seconds,omitempty"`
	MockDuration float64            `json:"mock_duration_seconds,omitempty"`
	MockError    string             `json:"mock_error,omitempty"`
	RealError    string             `json:"real_error,omitempty"`
}

// ComparisonReport is the full comparison output
type ComparisonReport struct {
	Timestamp          time.Time                          `json:"timestamp"`
	Summary            ComparisonSummary                  `json:"summary"`
	ByCategory         map[ComparisonCategory][]ComparisonResult `json:"by_category"`
	ConsistencyPercent float64                            `json:"consistency_percent"`
}

// ComparisonSummary holds aggregate statistics
type ComparisonSummary struct {
	TotalCompared     int `json:"total_compared"`
	Consistent        int `json:"consistent"`
	MockNeedsFix      int `json:"mock_needs_fix"`
	TestNeedsFix      int `json:"test_needs_fix"`
	MockIncorrect     int `json:"mock_incorrect"`
	MockOnlyPass      int `json:"mock_only_pass"`
	RealOnlyPass      int `json:"real_only_pass"`
	BothSkipped       int `json:"both_skipped"`
}

func main() {
	mockFile := flag.String("mock", "", "Path to mock test results JSON")
	realFile := flag.String("real", "", "Path to real API test results JSON")
	outputFile := flag.String("output", "", "Output file (default: stdout)")
	format := flag.String("format", "markdown", "Output format: markdown, json")
	flag.Parse()

	if *mockFile == "" || *realFile == "" {
		fmt.Fprintln(os.Stderr, "Usage: compare-test-results -mock=mock.json -real=real.json")
		fmt.Fprintln(os.Stderr, "Both -mock and -real flags are required")
		os.Exit(1)
	}

	mockReport, err := loadReport(*mockFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading mock report: %v\n", err)
		os.Exit(1)
	}

	realReport, err := loadReport(*realFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading real report: %v\n", err)
		os.Exit(1)
	}

	comparison := compare(mockReport, realReport)

	var out *os.File
	if *outputFile != "" {
		out, err = os.Create(*outputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating output file: %v\n", err)
			os.Exit(1)
		}
		defer out.Close()
	} else {
		out = os.Stdout
	}

	switch *format {
	case "json":
		outputJSON(out, comparison)
	default:
		outputMarkdown(out, comparison)
	}

	// Exit with non-zero if mock needs fixing
	if comparison.Summary.MockNeedsFix > 0 || comparison.Summary.MockIncorrect > 0 {
		os.Exit(1)
	}
}

func loadReport(path string) (*Report, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading file: %w", err)
	}

	var report Report
	if err := json.Unmarshal(data, &report); err != nil {
		return nil, fmt.Errorf("parsing JSON: %w", err)
	}

	return &report, nil
}

func compare(mock, real *Report) *ComparisonReport {
	report := &ComparisonReport{
		Timestamp:  time.Now(),
		ByCategory: make(map[ComparisonCategory][]ComparisonResult),
	}

	// Build maps for quick lookup
	mockTests := make(map[string]TestResult)
	for _, t := range mock.Tests {
		mockTests[t.Name] = t
	}

	realTests := make(map[string]TestResult)
	for _, t := range real.Tests {
		realTests[t.Name] = t
	}

	// Find all unique test names
	allTests := make(map[string]bool)
	for name := range mockTests {
		allTests[name] = true
	}
	for name := range realTests {
		allTests[name] = true
	}

	// Compare each test
	for testName := range allTests {
		mockTest, hasMock := mockTests[testName]
		realTest, hasReal := realTests[testName]

		result := ComparisonResult{
			TestName: testName,
		}

		if hasMock {
			result.MockStatus = mockTest.Status
			result.MockDuration = mockTest.Duration
			result.MockError = mockTest.FailureOutput
		}

		if hasReal {
			result.RealStatus = realTest.Status
			result.RealDuration = realTest.Duration
			result.RealError = realTest.FailureOutput
		}

		// Determine category
		switch {
		case !hasMock && hasReal:
			// Only in real, not in mock
			if realTest.Status == "pass" {
				result.Category = RealOnlyPass
				report.Summary.RealOnlyPass++
			} else if realTest.Status == "skip" {
				result.Category = BothSkipped
				report.Summary.BothSkipped++
			} else {
				result.Category = TestNeedsFix
				report.Summary.TestNeedsFix++
			}

		case hasMock && !hasReal:
			// Only in mock, not in real
			if mockTest.Status == "pass" {
				result.Category = MockOnlyPass
				report.Summary.MockOnlyPass++
			} else {
				result.Category = BothSkipped
				report.Summary.BothSkipped++
			}

		case mockTest.Status == "skip" && realTest.Status == "skip":
			result.Category = BothSkipped
			report.Summary.BothSkipped++

		case mockTest.Status == "pass" && realTest.Status == "pass":
			result.Category = MockConsistent
			report.Summary.Consistent++

		case mockTest.Status == "fail" && realTest.Status == "pass":
			// Real passes but mock fails - MOCK NEEDS FIXING
			result.Category = MockNeedsFix
			report.Summary.MockNeedsFix++

		case mockTest.Status == "pass" && realTest.Status == "fail":
			// Mock passes but real fails - MOCK IS INCORRECT
			result.Category = MockIncorrect
			report.Summary.MockIncorrect++

		case mockTest.Status == "fail" && realTest.Status == "fail":
			// Both fail - test needs fixing
			result.Category = TestNeedsFix
			report.Summary.TestNeedsFix++

		case mockTest.Status == "pass" && realTest.Status == "skip":
			result.Category = MockOnlyPass
			report.Summary.MockOnlyPass++

		case mockTest.Status == "skip" && realTest.Status == "pass":
			result.Category = RealOnlyPass
			report.Summary.RealOnlyPass++

		default:
			result.Category = BothSkipped
			report.Summary.BothSkipped++
		}

		report.ByCategory[result.Category] = append(report.ByCategory[result.Category], result)
	}

	report.Summary.TotalCompared = len(allTests)

	// Calculate consistency percentage (only for tests that ran on both)
	testedOnBoth := report.Summary.Consistent + report.Summary.MockNeedsFix +
		report.Summary.MockIncorrect + report.Summary.TestNeedsFix
	if testedOnBoth > 0 {
		report.ConsistencyPercent = float64(report.Summary.Consistent) / float64(testedOnBoth) * 100
	}

	// Sort results by test name
	for cat := range report.ByCategory {
		sort.Slice(report.ByCategory[cat], func(i, j int) bool {
			return report.ByCategory[cat][i].TestName < report.ByCategory[cat][j].TestName
		})
	}

	return report
}

func outputJSON(out *os.File, report *ComparisonReport) {
	encoder := json.NewEncoder(out)
	encoder.SetIndent("", "  ")
	encoder.Encode(report)
}

func outputMarkdown(out *os.File, report *ComparisonReport) {
	// Header
	fmt.Fprintln(out, "# Mock vs Real API Test Comparison Report")
	fmt.Fprintln(out)
	fmt.Fprintf(out, "**Generated:** %s\n\n", report.Timestamp.Format(time.RFC3339))

	// Consistency score
	consistencyIcon := "✅"
	if report.ConsistencyPercent < 100 {
		consistencyIcon = "⚠️"
	}
	if report.Summary.MockNeedsFix > 0 || report.Summary.MockIncorrect > 0 {
		consistencyIcon = "❌"
	}

	fmt.Fprintf(out, "## %s Mock Server Consistency: %.1f%%\n\n", consistencyIcon, report.ConsistencyPercent)

	// Summary table
	fmt.Fprintln(out, "## Summary")
	fmt.Fprintln(out)
	fmt.Fprintln(out, "| Category | Count | Description |")
	fmt.Fprintln(out, "|----------|------:|-------------|")
	fmt.Fprintf(out, "| ✅ Consistent | %d | Tests pass on both mock and real API |\n", report.Summary.Consistent)
	fmt.Fprintf(out, "| 🔧 Mock Needs Fix | %d | Tests pass on real but fail on mock |\n", report.Summary.MockNeedsFix)
	fmt.Fprintf(out, "| 🐛 Test Needs Fix | %d | Tests fail on both mock and real API |\n", report.Summary.TestNeedsFix)
	fmt.Fprintf(out, "| ⚠️ Mock Incorrect | %d | Tests pass on mock but fail on real API |\n", report.Summary.MockIncorrect)
	fmt.Fprintf(out, "| 🟡 Mock Only Pass | %d | Tests only pass on mock (skipped on real) |\n", report.Summary.MockOnlyPass)
	fmt.Fprintf(out, "| 🔵 Real Only Pass | %d | Tests only pass on real (not in mock) |\n", report.Summary.RealOnlyPass)
	fmt.Fprintf(out, "| ⏭️ Both Skipped | %d | Tests skipped on both |\n", report.Summary.BothSkipped)
	fmt.Fprintf(out, "| **Total** | **%d** | All unique tests |\n", report.Summary.TotalCompared)
	fmt.Fprintln(out)

	// Mock Needs Fix section (CRITICAL)
	if len(report.ByCategory[MockNeedsFix]) > 0 {
		fmt.Fprintln(out, "## 🔧 Mock Needs Fixing (Real PASS, Mock FAIL)")
		fmt.Fprintln(out)
		fmt.Fprintln(out, "These tests pass on the real API but fail on the mock server. The mock server needs to be updated to match real API behavior.")
		fmt.Fprintln(out)
		fmt.Fprintln(out, "| Test Name | Mock Error |")
		fmt.Fprintln(out, "|-----------|------------|")
		for _, r := range report.ByCategory[MockNeedsFix] {
			mockErr := truncate(r.MockError, 100)
			fmt.Fprintf(out, "| `%s` | %s |\n", r.TestName, mockErr)
		}
		fmt.Fprintln(out)
	}

	// Mock Incorrect section (CRITICAL)
	if len(report.ByCategory[MockIncorrect]) > 0 {
		fmt.Fprintln(out, "## ⚠️ Mock Incorrect (Mock PASS, Real FAIL)")
		fmt.Fprintln(out)
		fmt.Fprintln(out, "**WARNING:** These tests pass on mock but fail on real API. This means the mock is TOO PERMISSIVE and doesn't accurately simulate the real API behavior.")
		fmt.Fprintln(out)
		fmt.Fprintln(out, "| Test Name | Real Error |")
		fmt.Fprintln(out, "|-----------|------------|")
		for _, r := range report.ByCategory[MockIncorrect] {
			realErr := truncate(r.RealError, 100)
			fmt.Fprintf(out, "| `%s` | %s |\n", r.TestName, realErr)
		}
		fmt.Fprintln(out)
	}

	// Test Needs Fix section
	if len(report.ByCategory[TestNeedsFix]) > 0 {
		fmt.Fprintln(out, "## 🐛 Tests Failing on Both")
		fmt.Fprintln(out)
		fmt.Fprintln(out, "These tests fail on both mock and real API. The test or resource implementation needs to be fixed.")
		fmt.Fprintln(out)
		fmt.Fprintln(out, "| Test Name |")
		fmt.Fprintln(out, "|-----------|")
		for _, r := range report.ByCategory[TestNeedsFix] {
			fmt.Fprintf(out, "| `%s` |\n", r.TestName)
		}
		fmt.Fprintln(out)
	}

	// Consistent tests section
	if len(report.ByCategory[MockConsistent]) > 0 {
		fmt.Fprintln(out, "## ✅ Consistent Tests (Behavioral Contract)")
		fmt.Fprintln(out)
		fmt.Fprintln(out, "These tests pass on both mock and real API, proving the mock server accurately simulates real API behavior.")
		fmt.Fprintln(out)
		fmt.Fprintln(out, "| Test Name | Real Duration | Mock Duration | Speedup |")
		fmt.Fprintln(out, "|-----------|---------------|---------------|---------|")
		for _, r := range report.ByCategory[MockConsistent] {
			speedup := "N/A"
			if r.MockDuration > 0 && r.RealDuration > 0 {
				speedup = fmt.Sprintf("%.1fx", r.RealDuration/r.MockDuration)
			}
			fmt.Fprintf(out, "| `%s` | %.2fs | %.2fs | %s |\n",
				r.TestName, r.RealDuration, r.MockDuration, speedup)
		}
		fmt.Fprintln(out)
	}

	// Real Only Pass section (opportunity for mock expansion)
	if len(report.ByCategory[RealOnlyPass]) > 0 {
		fmt.Fprintln(out, "## 🔵 Real Only Pass (Mock Expansion Opportunity)")
		fmt.Fprintln(out)
		fmt.Fprintln(out, "These tests pass on real API but weren't run on mock. Running these on mock would expand coverage.")
		fmt.Fprintln(out)
		fmt.Fprintln(out, "<details>")
		fmt.Fprintf(out, "<summary>%d tests (click to expand)</summary>\n\n", len(report.ByCategory[RealOnlyPass]))
		for _, r := range report.ByCategory[RealOnlyPass] {
			fmt.Fprintf(out, "- `%s`\n", r.TestName)
		}
		fmt.Fprintln(out, "</details>")
		fmt.Fprintln(out)
	}
}

func truncate(s string, maxLen int) string {
	// Remove newlines and extra spaces
	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.TrimSpace(s)
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
}
