# HTTP Load Balancer Guide - Outputs
# ===================================

output "namespace" {
  description = "The namespace where resources were created"
  value       = local.namespace
}

output "origin_pool_name" {
  description = "Name of the created origin pool"
  value       = f5xc_origin_pool.this.name
}

output "loadbalancer_name" {
  description = "Name of the created HTTP load balancer"
  value       = f5xc_http_loadbalancer.this.name
}

output "loadbalancer_domain" {
  description = "Domain configured on the load balancer"
  value       = var.domain
}

output "cname_target" {
  description = "CNAME target for DNS configuration. Point your domain to this value."
  value       = "ves-io-${replace(var.domain, ".", "-")}.ac.vh.ves.io"
}

output "security_features" {
  description = "Security features enabled on this deployment"
  value = {
    waf           = var.enable_waf
    bot_defense   = var.enable_bot_defense
    rate_limiting = var.enable_rate_limiting
    js_challenge  = true
    threat_mesh   = true
  }
}

output "next_steps" {
  description = "Instructions for completing the deployment"
  value       = <<-EOT

    =====================================================
    DEPLOYMENT COMPLETE - Next Steps:
    =====================================================

    1. Configure DNS:
       Create a CNAME record pointing your domain to:
       ${f5xc_http_loadbalancer.this.name}.${local.namespace}.tenant.vh.ves.io

    2. Wait for certificate provisioning:
       Auto-cert may take a few minutes to provision.

    3. Test your application:
       https://${var.domain}

    4. Monitor in F5 Distributed Cloud Console:
       - Load Balancers: Check traffic and health
       - Security Analytics: View WAF and bot defense events

    =====================================================
  EOT
}
