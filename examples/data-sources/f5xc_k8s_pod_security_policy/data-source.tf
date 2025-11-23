# K8s Pod Security Policy Data Source Example
# Retrieves information about an existing K8s Pod Security Policy

# Look up an existing K8s Pod Security Policy by name
data "f5xc_k8s_pod_security_policy" "example" {
  name      = "example-k8s-pod-security-policy"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "k8s_pod_security_policy_id" {
#   value = data.f5xc_k8s_pod_security_policy.example.id
# }
