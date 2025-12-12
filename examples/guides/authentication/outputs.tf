# F5 Distributed Cloud Provider - Authentication Outputs
# =======================================================
#
# These outputs help verify that authentication is working correctly
# and provide useful information about the connection.

output "api_url" {
  description = "The F5XC API URL used for authentication (from environment or configuration)"
  value       = var.ves_api_url != "" ? var.ves_api_url : "Set via VES_API_URL environment variable"
}

output "authentication_method" {
  description = "The authentication method being used based on configuration"
  value       = var.ves_api_token != "" ? "API Token" : (var.ves_p12_file != "" ? "P12 Certificate" : (var.ves_cert != "" ? "PEM Certificate" : "Environment Variables"))
}

output "system_namespace" {
  description = "Information about the system namespace (confirms authentication is working)"
  value = {
    name        = data.f5xc_namespace.system.name
    description = data.f5xc_namespace.system.description
  }
}
