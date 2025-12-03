#!/bin/bash
# Sequential test runner with rate limit protection
#
# Required environment variables:
#   F5XC_API_URL - F5 XC API URL (e.g., https://console.ves.volterra.io/api)
#   F5XC_API_P12_FILE - Path to P12 certificate file
#   F5XC_P12_PASSWORD - Password for P12 certificate
#   TF_ACC=1 - Enable acceptance tests

set -e

# Validate required environment variables
if [[ -z "$F5XC_API_URL" ]]; then
    echo "ERROR: F5XC_API_URL environment variable is required"
    exit 1
fi

if [[ -z "$F5XC_API_P12_FILE" ]]; then
    echo "ERROR: F5XC_API_P12_FILE environment variable is required"
    exit 1
fi

if [[ -z "$F5XC_P12_PASSWORD" ]]; then
    echo "ERROR: F5XC_P12_PASSWORD environment variable is required"
    exit 1
fi

export TF_ACC=1

RESULTS_FILE="test-results-$(date +%Y%m%d-%H%M%S).log"
DELAY_BETWEEN_BATCHES=20

# All non-skipped resources (80 total)
RESOURCES=(
    "AddressAllocator"
    "AdvertisePolicy"
    "AlertPolicy"
    "AlertReceiver"
    "AllowedTenant"
    "APICrawler"
    "APIDefinition"
    "APIDiscovery"
    "APITesting"
    "AppAPIGroup"
    "AppFirewall"
    "AppSetting"
    "AppType"
    "BGPAsnSet"
    "BGPRoutingPolicy"
    "CDNCacheRule"
    "CertificateChain"
    "Certificate"
    "Cluster"
    "ContainerRegistry"
    "CRL"
    "DataGroup"
    "DataType"
    "DNSComplianceChecks"
    "DNSLBHealthCheck"
    "DNSLBPool"
    "Endpoint"
    "EnhancedFirewallPolicy"
    "FastACL"
    "FastACLRule"
    "FilterSet"
    "Fleet"
    "ForwardProxyPolicy"
    "GeoLocationSet"
    "GlobalLogReceiver"
    "Healthcheck"
    "HTTPLoadBalancer"
    "IKEPhase1Profile"
    "IKEPhase2Profile"
    "Ike1"
    "Ike2"
    "InfraprotectDenyListRule"
    "InfraprotectFirewallRule"
    "IPPrefixSet"
    "Irule"
    "K8SCluster"
    "K8SClusterRoleBinding"
    "K8SPodSecurityAdmission"
    "LogReceiver"
    "Namespace"
    "NetworkConnector"
    "NetworkFirewall"
    "NetworkInterface"
    "NetworkPolicy"
    "NetworkPolicyRule"
    "NetworkPolicyView"
    "OriginPool"
    "Policer"
    "ProtocolInspection"
    "Proxy"
    "RateLimiterPolicy"
    "RateLimiter"
    "SecretPolicy"
    "SecretPolicyRule"
    "Segment"
    "SensitiveDataPolicy"
    "ServicePolicy"
    "ServicePolicyRule"
    "TCPLoadBalancer"
    "Token"
    "TpmManager"
    "TrustedCAList"
    "Tunnel"
    "UserIdentification"
    "VirtualHost"
    "VirtualK8S"
    "VirtualNetwork"
    "VirtualSite"
    "VoltshareAdminPolicy"
    "WAFExclusionPolicy"
)

echo "========================================" | tee "$RESULTS_FILE"
echo "Sequential Test Run Started: $(date)" | tee -a "$RESULTS_FILE"
echo "Total resources to test: ${#RESOURCES[@]}" | tee -a "$RESULTS_FILE"
echo "Delay between batches: ${DELAY_BETWEEN_BATCHES}s" | tee -a "$RESULTS_FILE"
echo "========================================" | tee -a "$RESULTS_FILE"

PASSED=0
FAILED=0
SKIPPED=0
batch_count=0

for resource in "${RESOURCES[@]}"; do
    test_name="TestAcc${resource}Resource_basic"
    echo "" | tee -a "$RESULTS_FILE"
    echo "Testing: $test_name" | tee -a "$RESULTS_FILE"

    result=$(go test -v -timeout 10m -parallel 1 -run "$test_name" ./internal/provider/... 2>&1)

    if echo "$result" | grep -q "PASS:.*$test_name"; then
        echo "  RESULT: PASS" | tee -a "$RESULTS_FILE"
        ((PASSED++))
    elif echo "$result" | grep -q "FAIL:.*$test_name"; then
        echo "  RESULT: FAIL" | tee -a "$RESULTS_FILE"
        echo "$result" | tail -20 >> "$RESULTS_FILE"
        ((FAILED++))
    elif echo "$result" | grep -q "no tests to run"; then
        echo "  RESULT: NO_TEST_FOUND" | tee -a "$RESULTS_FILE"
        ((SKIPPED++))
    else
        echo "  RESULT: SKIPPED/OTHER" | tee -a "$RESULTS_FILE"
        ((SKIPPED++))
    fi

    ((batch_count++))

    # Delay every 5 tests
    if [ $((batch_count % 5)) -eq 0 ]; then
        echo "  [Rate limit protection: sleeping ${DELAY_BETWEEN_BATCHES}s]" | tee -a "$RESULTS_FILE"
        sleep $DELAY_BETWEEN_BATCHES
    fi
done

echo "" | tee -a "$RESULTS_FILE"
echo "========================================" | tee -a "$RESULTS_FILE"
echo "Test Run Completed: $(date)" | tee -a "$RESULTS_FILE"
echo "PASSED:  $PASSED" | tee -a "$RESULTS_FILE"
echo "FAILED:  $FAILED" | tee -a "$RESULTS_FILE"
echo "SKIPPED: $SKIPPED" | tee -a "$RESULTS_FILE"
echo "========================================" | tee -a "$RESULTS_FILE"
