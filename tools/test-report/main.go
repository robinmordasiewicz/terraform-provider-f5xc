// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

// test-report parses Go test JSON output and generates categorized test reports.
//
// Usage:
//
//	go test -json ./... | go run tools/test-report/main.go
//	go test -json ./... | go run tools/test-report/main.go -format=markdown
//	go test -json ./... | go run tools/test-report/main.go -format=junit
//	go test -json ./... | go run tools/test-report/main.go -output=report.json
//
// Test categories are determined by naming convention:
//   - TestAcc* -> REAL_API (tests against real F5 XC API)
//   - TestMock* -> MOCK_API (tests against mock server)
//   - Test* (other) -> UNIT (unit tests)
//
// Transient errors are detected by output patterns:
//   - RATE_LIMIT: HTTP 429, "Too Many Requests"
//   - TIMEOUT: "context deadline exceeded", "timeout"
//   - CONNECTION: "connection refused", "no such host"
//   - SERVER_ERROR: HTTP 5xx errors
package main

import (
	"bufio"
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"
	"strings"
	"time"
)

// TestEvent represents a single test event from go test -json output
type TestEvent struct {
	Time    time.Time `json:"Time"`
	Action  string    `json:"Action"`
	Package string    `json:"Package"`
	Test    string    `json:"Test"`
	Elapsed float64   `json:"Elapsed"`
	Output  string    `json:"Output"`
}

// TestCategory represents the type of test
type TestCategory string

const (
	CategoryRealAPI TestCategory = "REAL_API"
	CategoryMockAPI TestCategory = "MOCK_API"
	CategoryUnit    TestCategory = "UNIT"
)

// TransientErrorType classifies transient failures that may succeed on retry
type TransientErrorType string

const (
	TransientNone        TransientErrorType = ""
	TransientRateLimit   TransientErrorType = "RATE_LIMIT"
	TransientTimeout     TransientErrorType = "TIMEOUT"
	TransientConnection  TransientErrorType = "CONNECTION"
	TransientServerError TransientErrorType = "SERVER_ERROR"
)

// SkipReasonType classifies why tests were skipped
type SkipReasonType string

const (
	SkipReasonUnknown     SkipReasonType = "UNKNOWN"
	SkipReasonCredentials SkipReasonType = "MISSING_CREDENTIALS"
	SkipReasonMockMode    SkipReasonType = "MOCK_MODE_ONLY"
	SkipReasonPreCheck    SkipReasonType = "PRECHECK_FAILED"
	SkipReasonNotImpl     SkipReasonType = "NOT_IMPLEMENTED"
	SkipReasonEnvironment SkipReasonType = "ENVIRONMENT"
	SkipReasonInfra       SkipReasonType = "INFRASTRUCTURE"
)

// TestResult holds the result of a single test
type TestResult struct {
	Name           string             `json:"name"`
	Package        string             `json:"package"`
	Category       TestCategory       `json:"category"`
	Status         string             `json:"status"` // pass, fail, skip
	Duration       float64            `json:"duration_seconds"`
	SkipReason     string             `json:"skip_reason,omitempty"`
	SkipReasonType SkipReasonType     `json:"skip_reason_type,omitempty"`
	TransientError TransientErrorType `json:"transient_error,omitempty"`
	FailureOutput  string             `json:"failure_output,omitempty"`
}

// CategorySummary holds statistics for a test category
type CategorySummary struct {
	Total   int     `json:"total"`
	Passed  int     `json:"passed"`
	Failed  int     `json:"failed"`
	Skipped int     `json:"skipped"`
	Duration float64 `json:"duration_seconds"`
}

// TransientErrorSummary tracks transient error occurrences
type TransientErrorSummary struct {
	Type  TransientErrorType `json:"type"`
	Count int                `json:"count"`
	Tests []string           `json:"tests"`
}

// SkipReasonSummary tracks skip reason occurrences
type SkipReasonSummary struct {
	Type  SkipReasonType `json:"type"`
	Count int            `json:"count"`
}

// DurationStats holds duration statistics
type DurationStats struct {
	Min    float64 `json:"min_seconds"`
	Max    float64 `json:"max_seconds"`
	Avg    float64 `json:"avg_seconds"`
	P50    float64 `json:"p50_seconds"`
	P90    float64 `json:"p90_seconds"`
	P99    float64 `json:"p99_seconds"`
	Total  float64 `json:"total_seconds"`
}

// Report is the full test report
type Report struct {
	Timestamp       time.Time                         `json:"timestamp"`
	TotalDuration   float64                           `json:"total_duration_seconds"`
	TotalTests      int                               `json:"total_tests"`
	TotalPassed     int                               `json:"total_passed"`
	TotalFailed     int                               `json:"total_failed"`
	TotalSkipped    int                               `json:"total_skipped"`
	ByCategory      map[TestCategory]*CategorySummary `json:"by_category"`
	Tests           []TestResult                      `json:"tests"`
	FailedTests     []TestResult                      `json:"failed_tests"`
	TransientErrors []TransientErrorSummary           `json:"transient_errors,omitempty"`
	SkipReasons     []SkipReasonSummary               `json:"skip_reasons,omitempty"`
	SlowestTests    []TestResult                      `json:"slowest_tests,omitempty"`
	DurationStats   *DurationStats                    `json:"duration_stats,omitempty"`
}

// JUnit XML types
type JUnitTestSuites struct {
	XMLName   xml.Name         `xml:"testsuites"`
	Name      string           `xml:"name,attr"`
	Tests     int              `xml:"tests,attr"`
	Failures  int              `xml:"failures,attr"`
	Skipped   int              `xml:"skipped,attr"`
	Time      float64          `xml:"time,attr"`
	TestSuite []JUnitTestSuite `xml:"testsuite"`
}

type JUnitTestSuite struct {
	Name      string          `xml:"name,attr"`
	Tests     int             `xml:"tests,attr"`
	Failures  int             `xml:"failures,attr"`
	Skipped   int             `xml:"skipped,attr"`
	Time      float64         `xml:"time,attr"`
	TestCases []JUnitTestCase `xml:"testcase"`
}

type JUnitTestCase struct {
	Name      string        `xml:"name,attr"`
	ClassName string        `xml:"classname,attr"`
	Time      float64       `xml:"time,attr"`
	Failure   *JUnitFailure `xml:"failure,omitempty"`
	Skipped   *JUnitSkipped `xml:"skipped,omitempty"`
}

type JUnitFailure struct {
	Message string `xml:"message,attr"`
	Type    string `xml:"type,attr"`
	Content string `xml:",chardata"`
}

type JUnitSkipped struct {
	Message string `xml:"message,attr,omitempty"`
}

func main() {
	format := flag.String("format", "text", "Output format: text, json, markdown, junit")
	output := flag.String("output", "", "Output file (default: stdout)")
	showAll := flag.Bool("all", false, "Show all tests, not just summary")
	slowestCount := flag.Int("slowest", 10, "Number of slowest tests to show")
	flag.Parse()

	report := parseTestOutput(*slowestCount)

	var out *os.File
	var err error
	if *output != "" {
		out, err = os.Create(*output)
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
		outputJSON(out, report)
	case "markdown":
		outputMarkdown(out, report, *showAll)
	case "junit":
		outputJUnit(out, report)
	default:
		outputText(out, report, *showAll)
	}

	// Exit with non-zero status if there are failures
	if report.TotalFailed > 0 {
		os.Exit(1)
	}
}

func parseTestOutput(slowestCount int) *Report {
	report := &Report{
		Timestamp:  time.Now(),
		ByCategory: make(map[TestCategory]*CategorySummary),
		Tests:      make([]TestResult, 0),
		FailedTests: make([]TestResult, 0),
	}

	// Initialize categories
	for _, cat := range []TestCategory{CategoryRealAPI, CategoryMockAPI, CategoryUnit} {
		report.ByCategory[cat] = &CategorySummary{}
	}

	// Track test results (tests can have multiple events)
	testResults := make(map[string]*TestResult)
	testOutputs := make(map[string][]string)

	scanner := bufio.NewScanner(os.Stdin)
	// Increase buffer size for large output lines
	buf := make([]byte, 0, 1024*1024)
	scanner.Buffer(buf, 10*1024*1024)

	for scanner.Scan() {
		var event TestEvent
		if err := json.Unmarshal(scanner.Bytes(), &event); err != nil {
			continue // Skip non-JSON lines
		}

		// Only process test-level events (not package-level)
		if event.Test == "" {
			continue
		}

		key := event.Package + "/" + event.Test

		switch event.Action {
		case "run":
			testResults[key] = &TestResult{
				Name:     event.Test,
				Package:  event.Package,
				Category: inferCategory(event.Test),
			}
			testOutputs[key] = make([]string, 0)

		case "output":
			if _, ok := testOutputs[key]; ok {
				testOutputs[key] = append(testOutputs[key], event.Output)
			}

		case "pass":
			if result, ok := testResults[key]; ok {
				result.Status = "pass"
				result.Duration = event.Elapsed
			}

		case "fail":
			if result, ok := testResults[key]; ok {
				result.Status = "fail"
				result.Duration = event.Elapsed
				// Capture failure output and detect transient errors
				output := strings.Join(testOutputs[key], "")
				result.FailureOutput = truncateOutput(output, 2000)
				result.TransientError = detectTransientError(output)
			}

		case "skip":
			if result, ok := testResults[key]; ok {
				result.Status = "skip"
				result.Duration = event.Elapsed
				// Extract and classify skip reason
				output := strings.Join(testOutputs[key], "")
				result.SkipReason = extractSkipReason(output)
				result.SkipReasonType = classifySkipReason(result.SkipReason)
			}
		}
	}

	// Aggregate results
	var durations []float64
	for _, result := range testResults {
		if result.Status == "" {
			continue // Test didn't complete
		}

		report.Tests = append(report.Tests, *result)
		report.TotalTests++
		durations = append(durations, result.Duration)
		report.TotalDuration += result.Duration

		cat := result.Category
		summary := report.ByCategory[cat]
		summary.Duration += result.Duration

		switch result.Status {
		case "pass":
			report.TotalPassed++
			summary.Passed++
		case "fail":
			report.TotalFailed++
			summary.Failed++
			report.FailedTests = append(report.FailedTests, *result)
		case "skip":
			report.TotalSkipped++
			summary.Skipped++
		}
		summary.Total++
	}

	// Calculate duration statistics
	if len(durations) > 0 {
		report.DurationStats = calculateDurationStats(durations)
	}

	// Sort tests by category then name
	sort.Slice(report.Tests, func(i, j int) bool {
		if report.Tests[i].Category != report.Tests[j].Category {
			return report.Tests[i].Category < report.Tests[j].Category
		}
		return report.Tests[i].Name < report.Tests[j].Name
	})

	// Extract slowest tests
	report.SlowestTests = extractSlowestTests(report.Tests, slowestCount)

	// Aggregate transient errors
	report.TransientErrors = aggregateTransientErrors(report.FailedTests)

	// Aggregate skip reasons
	report.SkipReasons = aggregateSkipReasons(report.Tests)

	return report
}

func inferCategory(testName string) TestCategory {
	switch {
	case strings.HasPrefix(testName, "TestAcc"):
		return CategoryRealAPI
	case strings.HasPrefix(testName, "TestMock"):
		return CategoryMockAPI
	default:
		return CategoryUnit
	}
}

// detectTransientError checks failure output for patterns indicating transient failures
func detectTransientError(output string) TransientErrorType {
	lower := strings.ToLower(output)

	// Rate limiting patterns
	if strings.Contains(lower, "429") ||
		strings.Contains(lower, "too many requests") ||
		strings.Contains(lower, "rate limit") {
		return TransientRateLimit
	}

	// Timeout patterns
	if strings.Contains(lower, "context deadline exceeded") ||
		strings.Contains(lower, "timeout") ||
		strings.Contains(lower, "client.timeout") ||
		strings.Contains(lower, "operation timed out") {
		return TransientTimeout
	}

	// Connection patterns
	if strings.Contains(lower, "connection refused") ||
		strings.Contains(lower, "no such host") ||
		strings.Contains(lower, "network is unreachable") ||
		strings.Contains(lower, "dial tcp") {
		return TransientConnection
	}

	// Server error patterns
	if regexp.MustCompile(`\b50[0-9]\b`).MatchString(output) ||
		strings.Contains(lower, "internal server error") ||
		strings.Contains(lower, "service unavailable") {
		return TransientServerError
	}

	return TransientNone
}

// extractSkipReason finds the skip reason from test output
func extractSkipReason(output string) string {
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		// Look for common skip patterns
		if strings.Contains(line, "SKIP") || strings.Contains(line, "skip") ||
			strings.Contains(line, "Skipping") || strings.Contains(line, "skipping") {
			return strings.TrimSpace(line)
		}
		// Look for t.Skip() calls
		if strings.Contains(line, "--- SKIP:") {
			return strings.TrimSpace(line)
		}
	}
	return ""
}

// classifySkipReason categorizes the skip reason
func classifySkipReason(reason string) SkipReasonType {
	lower := strings.ToLower(reason)

	// Credential-related skips
	if strings.Contains(lower, "credential") ||
		strings.Contains(lower, "api_token") ||
		strings.Contains(lower, "api token") ||
		strings.Contains(lower, "authentication") ||
		strings.Contains(lower, "f5xc_api") ||
		strings.Contains(lower, "cloud credentials") {
		return SkipReasonCredentials
	}

	// Mock mode skips
	if strings.Contains(lower, "mock") ||
		strings.Contains(lower, "mock_mode") ||
		strings.Contains(lower, "mock mode") {
		return SkipReasonMockMode
	}

	// PreCheck failures
	if strings.Contains(lower, "precheck") ||
		strings.Contains(lower, "pre-check") ||
		strings.Contains(lower, "precondition") {
		return SkipReasonPreCheck
	}

	// Not implemented
	if strings.Contains(lower, "not implemented") ||
		strings.Contains(lower, "todo") ||
		strings.Contains(lower, "pending") {
		return SkipReasonNotImpl
	}

	// Infrastructure requirements
	if strings.Contains(lower, "site") ||
		strings.Contains(lower, "vk8s") ||
		strings.Contains(lower, "virtual k8s") ||
		strings.Contains(lower, "fleet") ||
		strings.Contains(lower, "aws") ||
		strings.Contains(lower, "azure") ||
		strings.Contains(lower, "gcp") ||
		strings.Contains(lower, "physical") ||
		strings.Contains(lower, "tenant admin") {
		return SkipReasonInfra
	}

	// Environment requirements
	if strings.Contains(lower, "environment") ||
		strings.Contains(lower, "env var") ||
		strings.Contains(lower, "configuration") {
		return SkipReasonEnvironment
	}

	return SkipReasonUnknown
}

// truncateOutput limits output to a maximum length
func truncateOutput(output string, maxLen int) string {
	if len(output) <= maxLen {
		return output
	}
	return output[:maxLen] + "... (truncated)"
}

// calculateDurationStats computes statistical measures for durations
func calculateDurationStats(durations []float64) *DurationStats {
	if len(durations) == 0 {
		return nil
	}

	// Sort for percentile calculation
	sorted := make([]float64, len(durations))
	copy(sorted, durations)
	sort.Float64s(sorted)

	var total float64
	min := sorted[0]
	max := sorted[len(sorted)-1]

	for _, d := range sorted {
		total += d
	}

	return &DurationStats{
		Min:   min,
		Max:   max,
		Avg:   total / float64(len(sorted)),
		P50:   percentile(sorted, 50),
		P90:   percentile(sorted, 90),
		P99:   percentile(sorted, 99),
		Total: total,
	}
}

// percentile calculates the pth percentile of a sorted slice
func percentile(sorted []float64, p int) float64 {
	if len(sorted) == 0 {
		return 0
	}
	if len(sorted) == 1 {
		return sorted[0]
	}

	index := float64(p) / 100.0 * float64(len(sorted)-1)
	lower := int(math.Floor(index))
	upper := int(math.Ceil(index))

	if lower == upper {
		return sorted[lower]
	}

	fraction := index - float64(lower)
	return sorted[lower] + fraction*(sorted[upper]-sorted[lower])
}

// extractSlowestTests returns the N slowest tests
func extractSlowestTests(tests []TestResult, n int) []TestResult {
	if len(tests) == 0 {
		return nil
	}

	// Copy and sort by duration descending
	sorted := make([]TestResult, len(tests))
	copy(sorted, tests)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Duration > sorted[j].Duration
	})

	if n > len(sorted) {
		n = len(sorted)
	}
	return sorted[:n]
}

// aggregateTransientErrors groups transient errors by type
func aggregateTransientErrors(failedTests []TestResult) []TransientErrorSummary {
	counts := make(map[TransientErrorType][]string)

	for _, test := range failedTests {
		if test.TransientError != TransientNone {
			counts[test.TransientError] = append(counts[test.TransientError], test.Name)
		}
	}

	var result []TransientErrorSummary
	for errType, tests := range counts {
		result = append(result, TransientErrorSummary{
			Type:  errType,
			Count: len(tests),
			Tests: tests,
		})
	}

	// Sort by count descending
	sort.Slice(result, func(i, j int) bool {
		return result[i].Count > result[j].Count
	})

	return result
}

// aggregateSkipReasons groups skip reasons by type
func aggregateSkipReasons(tests []TestResult) []SkipReasonSummary {
	counts := make(map[SkipReasonType]int)

	for _, test := range tests {
		if test.Status == "skip" {
			counts[test.SkipReasonType]++
		}
	}

	var result []SkipReasonSummary
	for reasonType, count := range counts {
		result = append(result, SkipReasonSummary{
			Type:  reasonType,
			Count: count,
		})
	}

	// Sort by count descending
	sort.Slice(result, func(i, j int) bool {
		return result[i].Count > result[j].Count
	})

	return result
}

func outputJSON(out *os.File, report *Report) {
	encoder := json.NewEncoder(out)
	encoder.SetIndent("", "  ")
	encoder.Encode(report)
}

func outputJUnit(out *os.File, report *Report) {
	suites := JUnitTestSuites{
		Name:     "F5XC Provider Tests",
		Tests:    report.TotalTests,
		Failures: report.TotalFailed,
		Skipped:  report.TotalSkipped,
		Time:     report.TotalDuration,
	}

	// Group tests by category into test suites
	testsByCategory := make(map[TestCategory][]TestResult)
	for _, test := range report.Tests {
		testsByCategory[test.Category] = append(testsByCategory[test.Category], test)
	}

	for _, cat := range []TestCategory{CategoryRealAPI, CategoryMockAPI, CategoryUnit} {
		tests := testsByCategory[cat]
		if len(tests) == 0 {
			continue
		}

		summary := report.ByCategory[cat]
		suite := JUnitTestSuite{
			Name:     string(cat),
			Tests:    summary.Total,
			Failures: summary.Failed,
			Skipped:  summary.Skipped,
			Time:     summary.Duration,
		}

		for _, test := range tests {
			tc := JUnitTestCase{
				Name:      test.Name,
				ClassName: test.Package,
				Time:      test.Duration,
			}

			if test.Status == "fail" {
				failureType := "AssertionError"
				if test.TransientError != TransientNone {
					failureType = string(test.TransientError)
				}
				tc.Failure = &JUnitFailure{
					Message: fmt.Sprintf("Test failed: %s", test.Name),
					Type:    failureType,
					Content: test.FailureOutput,
				}
			} else if test.Status == "skip" {
				tc.Skipped = &JUnitSkipped{
					Message: test.SkipReason,
				}
			}

			suite.TestCases = append(suite.TestCases, tc)
		}

		suites.TestSuite = append(suites.TestSuite, suite)
	}

	out.WriteString(xml.Header)
	encoder := xml.NewEncoder(out)
	encoder.Indent("", "  ")
	encoder.Encode(suites)
	out.WriteString("\n")
}

func outputText(out *os.File, report *Report, showAll bool) {
	divider := strings.Repeat("=", 78)
	subDivider := strings.Repeat("-", 78)

	fmt.Fprintln(out)
	fmt.Fprintln(out, divider)
	fmt.Fprintln(out, "                           TEST SUMMARY REPORT")
	fmt.Fprintln(out, divider)
	fmt.Fprintf(out, "Timestamp: %s\n", report.Timestamp.Format(time.RFC3339))
	fmt.Fprintf(out, "Total Duration: %.2fs\n", report.TotalDuration)
	fmt.Fprintln(out, subDivider)

	// Overall totals
	overallStatus := "PASS"
	if report.TotalFailed > 0 {
		overallStatus = "FAIL"
	}
	fmt.Fprintf(out, "[%s] TOTAL: %d tests | %d passed | %d failed | %d skipped\n",
		overallStatus, report.TotalTests, report.TotalPassed, report.TotalFailed, report.TotalSkipped)
	fmt.Fprintln(out, subDivider)

	// By category
	fmt.Fprintln(out, "BY CATEGORY:")
	for _, cat := range []TestCategory{CategoryRealAPI, CategoryMockAPI, CategoryUnit} {
		summary := report.ByCategory[cat]
		if summary.Total == 0 {
			continue
		}
		status := "PASS"
		if summary.Failed > 0 {
			status = "FAIL"
		}
		fmt.Fprintf(out, "  [%s] %-10s: %4d tests | %4d passed | %4d failed | %4d skipped (%.2fs)\n",
			status, cat, summary.Total, summary.Passed, summary.Failed, summary.Skipped, summary.Duration)
	}
	fmt.Fprintln(out, subDivider)

	// Duration statistics
	if report.DurationStats != nil {
		fmt.Fprintln(out, "DURATION STATISTICS:")
		stats := report.DurationStats
		fmt.Fprintf(out, "  Min: %.2fs | Max: %.2fs | Avg: %.2fs\n", stats.Min, stats.Max, stats.Avg)
		fmt.Fprintf(out, "  P50: %.2fs | P90: %.2fs | P99: %.2fs\n", stats.P50, stats.P90, stats.P99)
		fmt.Fprintln(out, subDivider)
	}

	// Transient errors
	if len(report.TransientErrors) > 0 {
		fmt.Fprintln(out, "TRANSIENT ERRORS (Retry Candidates):")
		for _, te := range report.TransientErrors {
			fmt.Fprintf(out, "  [%s] %d tests: %s\n", te.Type, te.Count, strings.Join(te.Tests, ", "))
		}
		fmt.Fprintln(out, subDivider)
	}

	// Failed tests
	if len(report.FailedTests) > 0 {
		fmt.Fprintln(out, "FAILED TESTS:")
		for _, test := range report.FailedTests {
			transient := ""
			if test.TransientError != TransientNone {
				transient = fmt.Sprintf(" [%s]", test.TransientError)
			}
			fmt.Fprintf(out, "  [%s] %s (%.2fs)%s\n", test.Category, test.Name, test.Duration, transient)
		}
		fmt.Fprintln(out, subDivider)
	}

	// Skip reasons
	if len(report.SkipReasons) > 0 {
		fmt.Fprintln(out, "SKIP REASONS:")
		for _, sr := range report.SkipReasons {
			fmt.Fprintf(out, "  %-25s: %d tests\n", sr.Type, sr.Count)
		}
		fmt.Fprintln(out, subDivider)
	}

	// Slowest tests
	if len(report.SlowestTests) > 0 {
		fmt.Fprintln(out, "SLOWEST TESTS:")
		for i, test := range report.SlowestTests {
			fmt.Fprintf(out, "  %2d. %-50s %.2fs\n", i+1, test.Name, test.Duration)
		}
		fmt.Fprintln(out, subDivider)
	}

	// All tests (if requested)
	if showAll {
		fmt.Fprintln(out, "ALL TESTS:")
		currentCat := TestCategory("")
		for _, test := range report.Tests {
			if test.Category != currentCat {
				currentCat = test.Category
				fmt.Fprintf(out, "\n  --- %s ---\n", currentCat)
			}
			statusIcon := "✓"
			if test.Status == "fail" {
				statusIcon = "✗"
			} else if test.Status == "skip" {
				statusIcon = "○"
			}
			fmt.Fprintf(out, "  %s %s (%.2fs)\n", statusIcon, test.Name, test.Duration)
		}
		fmt.Fprintln(out, subDivider)
	}

	fmt.Fprintln(out, divider)
}

func outputMarkdown(out *os.File, report *Report, showAll bool) {
	fmt.Fprintln(out, "# Test Summary Report")
	fmt.Fprintln(out)
	fmt.Fprintf(out, "**Timestamp:** %s\n", report.Timestamp.Format(time.RFC3339))
	fmt.Fprintf(out, "**Duration:** %s (%.2fs)\n\n", formatDuration(report.TotalDuration), report.TotalDuration)

	// Overall summary
	overallStatus := "✅ PASS"
	if report.TotalFailed > 0 {
		overallStatus = "❌ FAIL"
	}
	fmt.Fprintf(out, "## Overall: %s\n\n", overallStatus)
	fmt.Fprintln(out, "| Metric | Count |")
	fmt.Fprintln(out, "|--------|------:|")
	fmt.Fprintf(out, "| Total | %d |\n", report.TotalTests)
	fmt.Fprintf(out, "| Passed | %d |\n", report.TotalPassed)
	fmt.Fprintf(out, "| Failed | %d |\n", report.TotalFailed)
	fmt.Fprintf(out, "| Skipped | %d |\n", report.TotalSkipped)
	fmt.Fprintln(out)

	// By category
	fmt.Fprintln(out, "## Results by Category")
	fmt.Fprintln(out)
	fmt.Fprintln(out, "| Category | Status | Total | Passed | Failed | Skipped |")
	fmt.Fprintln(out, "|----------|:------:|------:|-------:|-------:|--------:|")
	for _, cat := range []TestCategory{CategoryRealAPI, CategoryMockAPI, CategoryUnit} {
		summary := report.ByCategory[cat]
		if summary.Total == 0 {
			continue
		}
		status := "✅"
		if summary.Failed > 0 {
			status = "❌"
		}
		catName := string(cat)
		switch cat {
		case CategoryRealAPI:
			catName = "🔴 Real API"
		case CategoryMockAPI:
			catName = "🟡 Mock API"
		case CategoryUnit:
			catName = "🟢 Unit"
		}
		fmt.Fprintf(out, "| %s | %s | %d | %d | %d | %d |\n",
			catName, status, summary.Total, summary.Passed, summary.Failed, summary.Skipped)
	}
	fmt.Fprintln(out)

	// Transient errors
	if len(report.TransientErrors) > 0 {
		fmt.Fprintln(out, "## Transient Errors (Retry Candidates)")
		fmt.Fprintln(out)
		fmt.Fprintln(out, "| Error Type | Count | Tests |")
		fmt.Fprintln(out, "|------------|------:|-------|")
		for _, te := range report.TransientErrors {
			testList := strings.Join(te.Tests, "`, `")
			fmt.Fprintf(out, "| %s | %d | `%s` |\n", te.Type, te.Count, testList)
		}
		fmt.Fprintln(out)
	}

	// Failed tests
	if len(report.FailedTests) > 0 {
		fmt.Fprintln(out, "## Failed Tests")
		fmt.Fprintln(out)
		fmt.Fprintln(out, "| Category | Test Name | Duration | Transient? |")
		fmt.Fprintln(out, "|----------|-----------|----------|------------|")
		for _, test := range report.FailedTests {
			transient := "-"
			if test.TransientError != TransientNone {
				transient = string(test.TransientError)
			}
			fmt.Fprintf(out, "| %s | `%s` | %.2fs | %s |\n", test.Category, test.Name, test.Duration, transient)
		}
		fmt.Fprintln(out)
	}

	// Skip reasons
	if len(report.SkipReasons) > 0 {
		fmt.Fprintln(out, "## Skip Reasons")
		fmt.Fprintln(out)
		fmt.Fprintln(out, "| Reason | Count |")
		fmt.Fprintln(out, "|--------|------:|")
		for _, sr := range report.SkipReasons {
			fmt.Fprintf(out, "| %s | %d |\n", sr.Type, sr.Count)
		}
		fmt.Fprintln(out)
	}

	// Slowest tests
	if len(report.SlowestTests) > 0 {
		fmt.Fprintln(out, "## Slowest Tests (Top 10)")
		fmt.Fprintln(out)
		fmt.Fprintln(out, "| Rank | Test Name | Duration |")
		fmt.Fprintln(out, "|-----:|-----------|----------|")
		for i, test := range report.SlowestTests {
			fmt.Fprintf(out, "| %d | `%s` | %.2fs |\n", i+1, test.Name, test.Duration)
		}
		fmt.Fprintln(out)
	}

	// Duration statistics
	if report.DurationStats != nil {
		fmt.Fprintln(out, "## Duration Statistics")
		fmt.Fprintln(out)
		stats := report.DurationStats
		fmt.Fprintln(out, "| Metric | Value |")
		fmt.Fprintln(out, "|--------|------:|")
		fmt.Fprintf(out, "| Min | %.2fs |\n", stats.Min)
		fmt.Fprintf(out, "| Max | %.2fs |\n", stats.Max)
		fmt.Fprintf(out, "| Avg | %.2fs |\n", stats.Avg)
		fmt.Fprintf(out, "| P50 | %.2fs |\n", stats.P50)
		fmt.Fprintf(out, "| P90 | %.2fs |\n", stats.P90)
		fmt.Fprintf(out, "| P99 | %.2fs |\n", stats.P99)
		fmt.Fprintln(out)
	}

	// All tests (if requested)
	if showAll {
		fmt.Fprintln(out, "## All Tests")
		fmt.Fprintln(out)
		currentCat := TestCategory("")
		for _, test := range report.Tests {
			if test.Category != currentCat {
				currentCat = test.Category
				fmt.Fprintf(out, "\n### %s\n\n", currentCat)
				fmt.Fprintln(out, "| Status | Test Name | Duration |")
				fmt.Fprintln(out, "|--------|-----------|----------|")
			}
			status := "✅"
			if test.Status == "fail" {
				status = "❌"
			} else if test.Status == "skip" {
				status = "⏭️"
			}
			fmt.Fprintf(out, "| %s | `%s` | %.2fs |\n", status, test.Name, test.Duration)
		}
	}
}

// formatDuration formats seconds into human-readable duration
func formatDuration(seconds float64) string {
	if seconds < 60 {
		return fmt.Sprintf("%.0fs", seconds)
	}
	minutes := int(seconds) / 60
	secs := int(seconds) % 60
	if minutes < 60 {
		return fmt.Sprintf("%dm %ds", minutes, secs)
	}
	hours := minutes / 60
	minutes = minutes % 60
	return fmt.Sprintf("%dh %dm %ds", hours, minutes, secs)
}
