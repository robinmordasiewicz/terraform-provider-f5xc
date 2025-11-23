# K8s Pod Security Admission Data Source Example
# Retrieves information about an existing K8s Pod Security Admission

# Look up an existing K8s Pod Security Admission by name
data "f5xc_k8s_pod_security_admission" "example" {
  name      = "example-k8s-pod-security-admission"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "k8s_pod_security_admission_id" {
#   value = data.f5xc_k8s_pod_security_admission.example.id
# }
