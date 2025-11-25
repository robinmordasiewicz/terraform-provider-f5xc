---
page_title: "Securemesh Site V2 Nested Blocks - f5xc Provider"
subcategory: "Sites"
description: |-
  Nested block reference for the Securemesh Site V2 resource.
---

# Securemesh Site V2 Nested Blocks

This page contains detailed documentation for nested blocks in the `f5xc_securemesh_site_v2` resource.

For the main resource documentation, see [f5xc_securemesh_site_v2](/docs/resources/securemesh_site_v2).

## Contents

- [active-enhanced-firewall-policies](#active-enhanced-firewall-policies)
- [active-enhanced-firewall-policies-enhanced-firewall-policies](#active-enhanced-firewall-policies-enhanced-firewall-policies)
- [active-forward-proxy-policies](#active-forward-proxy-policies)
- [active-forward-proxy-policies-forward-proxy-policies](#active-forward-proxy-policies-forward-proxy-policies)
- [admin-user-credentials](#admin-user-credentials)
- [admin-user-credentials-admin-password](#admin-user-credentials-admin-password)
- [admin-user-credentials-admin-password-blindfold-secret-info](#admin-user-credentials-admin-password-blindfold-secret-info)
- [admin-user-credentials-admin-password-clear-secret-info](#admin-user-credentials-admin-password-clear-secret-info)
- [aws](#aws)
- [aws-not-managed](#aws-not-managed)
- [aws-not-managed-node-list](#aws-not-managed-node-list)
- [azure](#azure)
- [azure-not-managed](#azure-not-managed)
- [azure-not-managed-node-list](#azure-not-managed-node-list)
- [baremetal](#baremetal)
- [baremetal-not-managed](#baremetal-not-managed)
- [baremetal-not-managed-node-list](#baremetal-not-managed-node-list)
- [blocked-services](#blocked-services)
- [blocked-services-blocked-sevice](#blocked-services-blocked-sevice)
- [custom-proxy](#custom-proxy)
- [custom-proxy-password](#custom-proxy-password)
- [custom-proxy-password-blindfold-secret-info](#custom-proxy-password-blindfold-secret-info)
- [custom-proxy-password-clear-secret-info](#custom-proxy-password-clear-secret-info)
- [custom-proxy-bypass](#custom-proxy-bypass)
- [dc-cluster-group-sli](#dc-cluster-group-sli)
- [dc-cluster-group-slo](#dc-cluster-group-slo)
- [dns-ntp-config](#dns-ntp-config)
- [dns-ntp-config-custom-dns](#dns-ntp-config-custom-dns)
- [dns-ntp-config-custom-ntp](#dns-ntp-config-custom-ntp)
- [equinix](#equinix)
- [equinix-not-managed](#equinix-not-managed)
- [equinix-not-managed-node-list](#equinix-not-managed-node-list)
- [gcp](#gcp)
- [gcp-not-managed](#gcp-not-managed)
- [gcp-not-managed-node-list](#gcp-not-managed-node-list)
- [kvm](#kvm)
- [kvm-not-managed](#kvm-not-managed)
- [kvm-not-managed-node-list](#kvm-not-managed-node-list)
- [load-balancing](#load-balancing)
- [local-vrf](#local-vrf)
- [local-vrf-sli-config](#local-vrf-sli-config)
- [local-vrf-sli-config-static-routes](#local-vrf-sli-config-static-routes)
- [local-vrf-sli-config-static-v6-routes](#local-vrf-sli-config-static-v6-routes)
- [local-vrf-slo-config](#local-vrf-slo-config)
- [local-vrf-slo-config-static-routes](#local-vrf-slo-config-static-routes)
- [local-vrf-slo-config-static-v6-routes](#local-vrf-slo-config-static-v6-routes)
- [log-receiver](#log-receiver)
- [nutanix](#nutanix)
- [nutanix-not-managed](#nutanix-not-managed)
- [nutanix-not-managed-node-list](#nutanix-not-managed-node-list)
- [oci](#oci)
- [oci-not-managed](#oci-not-managed)
- [oci-not-managed-node-list](#oci-not-managed-node-list)
- [offline-survivability-mode](#offline-survivability-mode)
- [openstack](#openstack)
- [openstack-not-managed](#openstack-not-managed)
- [openstack-not-managed-node-list](#openstack-not-managed-node-list)
- [performance-enhancement-mode](#performance-enhancement-mode)
- [performance-enhancement-mode-perf-mode-l3-enhanced](#performance-enhancement-mode-perf-mode-l3-enhanced)
- [re-select](#re-select)
- [re-select-specific-re](#re-select-specific-re)
- [site-mesh-group-on-slo](#site-mesh-group-on-slo)
- [site-mesh-group-on-slo-site-mesh-group](#site-mesh-group-on-slo-site-mesh-group)
- [software-settings](#software-settings)
- [software-settings-os](#software-settings-os)
- [software-settings-sw](#software-settings-sw)
- [timeouts](#timeouts)
- [upgrade-settings](#upgrade-settings)
- [upgrade-settings-kubernetes-upgrade-drain](#upgrade-settings-kubernetes-upgrade-drain)
- [upgrade-settings-kubernetes-upgrade-drain-enable-upgrade-drain](#upgrade-settings-kubernetes-upgrade-drain-enable-upgrade-drain)
- [vmware](#vmware)
- [vmware-not-managed](#vmware-not-managed)
- [vmware-not-managed-node-list](#vmware-not-managed-node-list)

---

<a id="active-enhanced-firewall-policies"></a>

### Active Enhanced Firewall Policies

`enhanced_firewall_policies` - (Optional) Enhanced Firewall Policy. Ordered List of Enhanced Firewall Policies active. See [Enhanced Firewall Policies](#active-enhanced-firewall-policies-enhanced-firewall-policies) below.

<a id="active-enhanced-firewall-policies-enhanced-firewall-policies"></a>

### Active Enhanced Firewall Policies Enhanced Firewall Policies

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

<a id="active-forward-proxy-policies"></a>

### Active Forward Proxy Policies

`forward_proxy_policies` - (Optional) Forward Proxy Policies. Ordered List of Forward Proxy Policies active. See [Forward Proxy Policies](#active-forward-proxy-policies-forward-proxy-policies) below.

<a id="active-forward-proxy-policies-forward-proxy-policies"></a>

### Active Forward Proxy Policies Forward Proxy Policies

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

<a id="admin-user-credentials"></a>

### Admin User Credentials

`admin_password` - (Optional) Secret. SecretType is used in an object to indicate a sensitive/confidential field. See [Admin Password](#admin-user-credentials-admin-password) below.

`ssh_key` - (Optional) Public SSH key. Provided Public SSH key can be used for accessing nodes of the site. When provided, customers can SSH to the nodes of this Customer Edge site using admin as the user (`String`).

<a id="admin-user-credentials-admin-password"></a>

### Admin User Credentials Admin Password

`blindfold_secret_info` - (Optional) Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management. See [Blindfold Secret Info](#admin-user-credentials-admin-password-blindfold-secret-info) below.

`clear_secret_info` - (Optional) In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted. See [Clear Secret Info](#admin-user-credentials-admin-password-clear-secret-info) below.

<a id="admin-user-credentials-admin-password-blindfold-secret-info"></a>

### Admin User Credentials Admin Password Blindfold Secret Info

`decryption_provider` - (Optional) Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service (`String`).

`location` - (Optional) Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location (`String`).

`store_provider` - (Optional) Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:/// (`String`).

<a id="admin-user-credentials-admin-password-clear-secret-info"></a>

### Admin User Credentials Admin Password Clear Secret Info

`provider_ref` - (Optional) Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:/// (`String`).

`url` - (Optional) URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding (`String`).

<a id="aws"></a>

### AWS

`not_managed` - (Optional) List of Nodes. This section will show nodes associated with this site. Note: For sites that are not orchestrated by F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it will be shown in this section. See [Not Managed](#aws-not-managed) below.

<a id="aws-not-managed"></a>

### AWS Not Managed

`node_list` - (Optional) Nodes. This section will show nodes associated with this site. Note: For sites that are not orchestrated by F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it will be shown in this section. See [Node List](#aws-not-managed-node-list) below.

<a id="aws-not-managed-node-list"></a>

### AWS Not Managed Node List

`hostname` - (Optional) Hostname. Hostname for this Node (`String`).

`interface_list` - (Optional) Interfaces. Manage interfaces belonging to this node (`Block`).

`public_ip` - (Optional) Public IP. Public IP for this Node (`String`).

`type` - (Optional) Type. Type for this Node, can be Control or Worker (`String`).

<a id="azure"></a>

### Azure

`not_managed` - (Optional) List of Nodes. This section will show nodes associated with this site. Note: For sites that are not orchestrated by F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it will be shown in this section. See [Not Managed](#azure-not-managed) below.

<a id="azure-not-managed"></a>

### Azure Not Managed

`node_list` - (Optional) Nodes. This section will show nodes associated with this site. Note: For sites that are not orchestrated by F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it will be shown in this section. See [Node List](#azure-not-managed-node-list) below.

<a id="azure-not-managed-node-list"></a>

### Azure Not Managed Node List

`hostname` - (Optional) Hostname. Hostname for this Node (`String`).

`interface_list` - (Optional) Interfaces. Manage interfaces belonging to this node (`Block`).

`public_ip` - (Optional) Public IP. Public IP for this Node (`String`).

`type` - (Optional) Type. Type for this Node, can be Control or Worker (`String`).

<a id="baremetal"></a>

### Baremetal

`not_managed` - (Optional) List of Nodes. This section will show nodes associated with this site. Note: For sites that are not orchestrated by F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it will be shown in this section. See [Not Managed](#baremetal-not-managed) below.

<a id="baremetal-not-managed"></a>

### Baremetal Not Managed

`node_list` - (Optional) Nodes. This section will show nodes associated with this site. Note: For sites that are not orchestrated by F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it will be shown in this section. See [Node List](#baremetal-not-managed-node-list) below.

<a id="baremetal-not-managed-node-list"></a>

### Baremetal Not Managed Node List

`hostname` - (Optional) Hostname. Hostname for this Node (`String`).

`interface_list` - (Optional) Interfaces. Manage interfaces belonging to this node (`Block`).

`public_ip` - (Optional) Public IP. Public IP for this Node (`String`).

`type` - (Optional) Type. Type for this Node, can be Control or Worker (`String`).

<a id="blocked-services"></a>

### Blocked Services

`blocked_sevice` - (Optional) Disable Node Local Services. See [Blocked Sevice](#blocked-services-blocked-sevice) below.

<a id="blocked-services-blocked-sevice"></a>

### Blocked Services Blocked Sevice

`dns` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`network_type` - (Optional) Virtual Network Type. Different types of virtual networks understood by the system Virtual-network of type VIRTUAL_NETWORK_SITE_LOCAL provides connectivity to public (outside) network. This is an insecure network and is connected to public internet via NAT Gateways/firwalls Virtual-network of this type is local to every site. Two virtual networks of this type on different sites are neither related nor connected. Constraints: There can be atmost one virtual network of this type in a given site. This network type is supported on CE sites. This network is created automatically and present on all sites Virtual-network of type VIRTUAL_NETWORK_SITE_LOCAL_INSIDE is a private network inside site. It is a secure network and is not connected to public network. Virtual-network of this type is local to every site. Two virtual networks of this type on different sites are neither related nor connected. Constraints: There can be atmost one virtual network of this type in a given site. This network type is supported on CE sites. This network is created during provisioning of site User defined per-site virtual network. Scope of this virtual network is limited to the site. This is not yet supported Virtual-network of type VIRTUAL_NETWORK_PUBLIC directly conects to the public internet. Virtual-network of this type is local to every site. Two virtual networks of this type on different sites are neither related nor connected. Constraints: There can be atmost one virtual network of this type in a given site. This network type is supported on RE sites only It is an internally created by the system. They must not be created by user Virtual Neworks with global scope across different sites in F5XC domain. An example global virtual-network called 'AIN Network' is created for every tenant. for volterra fabric Constraints: It is currently only supported as internally created by the system. vK8s service network for a given tenant. Used to advertise a virtual host only to vk8s pods for that tenant Constraints: It is an internally created by the system. Must not be created by user VER internal network for the site. It can only be used for virtual hosts with SMA_PROXY type proxy Constraints: It is an internally created by the system. Must not be created by user Virtual-network of type VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE represents both VIRTUAL_NETWORK_SITE_LOCAL and VIRTUAL_NETWORK_SITE_LOCAL_INSIDE Constraints: This network type is only meaningful in an advertise policy When virtual-network of type VIRTUAL_NETWORK_IP_AUTO is selected for an endpoint, VER will try to determine the network based on the provided IP address Constraints: This network type is only meaningful in an endpoint VoltADN Private Network is used on volterra RE(s) to connect to customer private networks This network is created by opening a support ticket This network is per site srv6 network VER IP Fabric network for the site. This Virtual network type is used for exposing virtual host on IP Fabric network on the VER site or for endpoint in IP Fabric network Constraints: It is an internally created by the system. Must not be created by user Network internally created for a segment Constraints: It is an internally created by the system. Must not be created by user. Possible values are `VIRTUAL_NETWORK_SITE_LOCAL`, `VIRTUAL_NETWORK_SITE_LOCAL_INSIDE`, `VIRTUAL_NETWORK_PER_SITE`, `VIRTUAL_NETWORK_PUBLIC`, `VIRTUAL_NETWORK_GLOBAL`, `VIRTUAL_NETWORK_SITE_SERVICE`, `VIRTUAL_NETWORK_VER_INTERNAL`, `VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE`, `VIRTUAL_NETWORK_IP_AUTO`, `VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK`, `VIRTUAL_NETWORK_SRV6_NETWORK`, `VIRTUAL_NETWORK_IP_FABRIC`, `VIRTUAL_NETWORK_SEGMENT`. Defaults to `VIRTUAL_NETWORK_SITE_LOCAL` (`String`).

`ssh` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`web_user_interface` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

<a id="custom-proxy"></a>

### Custom Proxy

`disable_re_tunnel` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`enable_re_tunnel` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`password` - (Optional) Secret. SecretType is used in an object to indicate a sensitive/confidential field. See [Password](#custom-proxy-password) below.

`proxy_ip_address` - (Optional) Proxy IPv4 Address. Specify the IPv4 Address of the internal Enterprise Proxy (`String`).

`proxy_port` - (Optional) Proxy Port. Specify the Port of the internal Enterprise Proxy (`Number`).

`username` - (Optional) Username. If the internal Enterprise Proxy is using basic authentication, specify the username. This is an optional field (`String`).

<a id="custom-proxy-password"></a>

### Custom Proxy Password

`blindfold_secret_info` - (Optional) Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management. See [Blindfold Secret Info](#custom-proxy-password-blindfold-secret-info) below.

`clear_secret_info` - (Optional) In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted. See [Clear Secret Info](#custom-proxy-password-clear-secret-info) below.

<a id="custom-proxy-password-blindfold-secret-info"></a>

### Custom Proxy Password Blindfold Secret Info

`decryption_provider` - (Optional) Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service (`String`).

`location` - (Optional) Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location (`String`).

`store_provider` - (Optional) Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:/// (`String`).

<a id="custom-proxy-password-clear-secret-info"></a>

### Custom Proxy Password Clear Secret Info

`provider_ref` - (Optional) Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:/// (`String`).

`url` - (Optional) URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding (`String`).

<a id="custom-proxy-bypass"></a>

### Custom Proxy Bypass

`proxy_bypass` - (Optional) Proxy Bypass. List of domains to bypass the proxy (`List`).

<a id="dc-cluster-group-sli"></a>

### Dc Cluster Group Sli

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

<a id="dc-cluster-group-slo"></a>

### Dc Cluster Group Slo

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

<a id="dns-ntp-config"></a>

### DNS NTP Config

`custom_dns` - (Optional) DNS Servers. DNS Servers. See [Custom DNS](#dns-ntp-config-custom-dns) below.

`custom_ntp` - (Optional) NTP Servers. NTP Servers. See [Custom NTP](#dns-ntp-config-custom-ntp) below.

`f5_dns_default` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`f5_ntp_default` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

<a id="dns-ntp-config-custom-dns"></a>

### DNS NTP Config Custom DNS

`dns_servers` - (Optional) DNS Servers. DNS Servers (`List`).

<a id="dns-ntp-config-custom-ntp"></a>

### DNS NTP Config Custom NTP

`ntp_servers` - (Optional) NTP Servers. NTP Servers (`List`).

<a id="equinix"></a>

### Equinix

`not_managed` - (Optional) List of Nodes. This section will show nodes associated with this site. Note: For sites that are not orchestrated by F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it will be shown in this section. See [Not Managed](#equinix-not-managed) below.

<a id="equinix-not-managed"></a>

### Equinix Not Managed

`node_list` - (Optional) Nodes. This section will show nodes associated with this site. Note: For sites that are not orchestrated by F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it will be shown in this section. See [Node List](#equinix-not-managed-node-list) below.

<a id="equinix-not-managed-node-list"></a>

### Equinix Not Managed Node List

`hostname` - (Optional) Hostname. Hostname for this Node (`String`).

`interface_list` - (Optional) Interfaces. Manage interfaces belonging to this node (`Block`).

`public_ip` - (Optional) Public IP. Public IP for this Node (`String`).

`type` - (Optional) Type. Type for this Node, can be Control or Worker (`String`).

<a id="gcp"></a>

### GCP

`not_managed` - (Optional) List of Nodes. This section will show nodes associated with this site. Note: For sites that are not orchestrated by F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it will be shown in this section. See [Not Managed](#gcp-not-managed) below.

<a id="gcp-not-managed"></a>

### GCP Not Managed

`node_list` - (Optional) Nodes. This section will show nodes associated with this site. Note: For sites that are not orchestrated by F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it will be shown in this section. See [Node List](#gcp-not-managed-node-list) below.

<a id="gcp-not-managed-node-list"></a>

### GCP Not Managed Node List

`hostname` - (Optional) Hostname. Hostname for this Node (`String`).

`interface_list` - (Optional) Interfaces. Manage interfaces belonging to this node (`Block`).

`public_ip` - (Optional) Public IP. Public IP for this Node (`String`).

`type` - (Optional) Type. Type for this Node, can be Control or Worker (`String`).

<a id="kvm"></a>

### Kvm

`not_managed` - (Optional) List of Nodes. This section will show nodes associated with this site. Note: For sites that are not orchestrated by F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it will be shown in this section. See [Not Managed](#kvm-not-managed) below.

<a id="kvm-not-managed"></a>

### Kvm Not Managed

`node_list` - (Optional) Nodes. This section will show nodes associated with this site. Note: For sites that are not orchestrated by F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it will be shown in this section. See [Node List](#kvm-not-managed-node-list) below.

<a id="kvm-not-managed-node-list"></a>

### Kvm Not Managed Node List

`hostname` - (Optional) Hostname. Hostname for this Node (`String`).

`interface_list` - (Optional) Interfaces. Manage interfaces belonging to this node (`Block`).

`public_ip` - (Optional) Public IP. Public IP for this Node (`String`).

`type` - (Optional) Type. Type for this Node, can be Control or Worker (`String`).

<a id="load-balancing"></a>

### Load Balancing

`vip_vrrp_mode` - (Optional) VRRP Virtual-IP. VRRP advertisement mode for VIP Invalid VRRP mode. Possible values are `VIP_VRRP_INVALID`, `VIP_VRRP_ENABLE`, `VIP_VRRP_DISABLE` (`String`).

<a id="local-vrf"></a>

### Local Vrf

`default_config` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`default_sli_config` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`sli_config` - (Optional) Site Local Network Configuration. Site local network configuration. See [Sli Config](#local-vrf-sli-config) below.

`slo_config` - (Optional) Site Local Network Configuration. Site local network configuration. See [Slo Config](#local-vrf-slo-config) below.

<a id="local-vrf-sli-config"></a>

### Local Vrf Sli Config

`labels` - (Optional) Network Labels. Add Labels for this network, these labels can be used in firewall policy (`Block`).

`nameserver` - (Optional) DNS V4 Server. Optional DNS V4 server IP to be used for name resolution (`String`).

`no_static_routes` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`no_v6_static_routes` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`static_routes` - (Optional) Static Routes List. See [Static Routes](#local-vrf-sli-config-static-routes) below.

`static_v6_routes` - (Optional) Static IPv6 Routes List. List of IPv6 static routes. See [Static V6 Routes](#local-vrf-sli-config-static-v6-routes) below.

`vip` - (Optional) Common V4 VIP. Optional common virtual V4 IP across all nodes to be used as automatic VIP (`String`).

<a id="local-vrf-sli-config-static-routes"></a>

### Local Vrf Sli Config Static Routes

`static_routes` - (Optional) Static Routes (`Block`).

<a id="local-vrf-sli-config-static-v6-routes"></a>

### Local Vrf Sli Config Static V6 Routes

`static_routes` - (Optional) Static IPv6 Routes. List of IPv6 static routes (`Block`).

<a id="local-vrf-slo-config"></a>

### Local Vrf Slo Config

`labels` - (Optional) Network Labels. Add Labels for this network, these labels can be used in firewall policy (`Block`).

`nameserver` - (Optional) DNS V4 Server. Optional DNS V4 server IP to be used for name resolution (`String`).

`no_static_routes` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`no_v6_static_routes` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`static_routes` - (Optional) Static Routes List. See [Static Routes](#local-vrf-slo-config-static-routes) below.

`static_v6_routes` - (Optional) Static IPv6 Routes List. List of IPv6 static routes. See [Static V6 Routes](#local-vrf-slo-config-static-v6-routes) below.

`vip` - (Optional) Common V4 VIP. Optional common virtual V4 IP across all nodes to be used as automatic VIP (`String`).

<a id="local-vrf-slo-config-static-routes"></a>

### Local Vrf Slo Config Static Routes

`static_routes` - (Optional) Static Routes (`Block`).

<a id="local-vrf-slo-config-static-v6-routes"></a>

### Local Vrf Slo Config Static V6 Routes

`static_routes` - (Optional) Static IPv6 Routes. List of IPv6 static routes (`Block`).

<a id="log-receiver"></a>

### Log Receiver

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

<a id="nutanix"></a>

### Nutanix

`not_managed` - (Optional) List of Nodes. This section will show nodes associated with this site. Note: For sites that are not orchestrated by F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it will be shown in this section. See [Not Managed](#nutanix-not-managed) below.

<a id="nutanix-not-managed"></a>

### Nutanix Not Managed

`node_list` - (Optional) Nodes. This section will show nodes associated with this site. Note: For sites that are not orchestrated by F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it will be shown in this section. See [Node List](#nutanix-not-managed-node-list) below.

<a id="nutanix-not-managed-node-list"></a>

### Nutanix Not Managed Node List

`hostname` - (Optional) Hostname. Hostname for this Node (`String`).

`interface_list` - (Optional) Interfaces. Manage interfaces belonging to this node (`Block`).

`public_ip` - (Optional) Public IP. Public IP for this Node (`String`).

`type` - (Optional) Type. Type for this Node, can be Control or Worker (`String`).

<a id="oci"></a>

### Oci

`not_managed` - (Optional) List of Nodes. This section will show nodes associated with this site. Note: For sites that are not orchestrated by F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it will be shown in this section. See [Not Managed](#oci-not-managed) below.

<a id="oci-not-managed"></a>

### Oci Not Managed

`node_list` - (Optional) Nodes. This section will show nodes associated with this site. Note: For sites that are not orchestrated by F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it will be shown in this section. See [Node List](#oci-not-managed-node-list) below.

<a id="oci-not-managed-node-list"></a>

### Oci Not Managed Node List

`hostname` - (Optional) Hostname. Hostname for this Node (`String`).

`interface_list` - (Optional) Interfaces. Manage interfaces belonging to this node (`Block`).

`public_ip` - (Optional) Public IP. Public IP for this Node (`String`).

`type` - (Optional) Type. Type for this Node, can be Control or Worker (`String`).

<a id="offline-survivability-mode"></a>

### Offline Survivability Mode

`enable_offline_survivability_mode` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`no_offline_survivability_mode` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

<a id="openstack"></a>

### Openstack

`not_managed` - (Optional) List of Nodes. This section will show nodes associated with this site. Note: For sites that are not orchestrated by F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it will be shown in this section. See [Not Managed](#openstack-not-managed) below.

<a id="openstack-not-managed"></a>

### Openstack Not Managed

`node_list` - (Optional) Nodes. This section will show nodes associated with this site. Note: For sites that are not orchestrated by F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it will be shown in this section. See [Node List](#openstack-not-managed-node-list) below.

<a id="openstack-not-managed-node-list"></a>

### Openstack Not Managed Node List

`hostname` - (Optional) Hostname. Hostname for this Node (`String`).

`interface_list` - (Optional) Interfaces. Manage interfaces belonging to this node (`Block`).

`public_ip` - (Optional) Public IP. Public IP for this Node (`String`).

`type` - (Optional) Type. Type for this Node, can be Control or Worker (`String`).

<a id="performance-enhancement-mode"></a>

### Performance Enhancement Mode

`perf_mode_l3_enhanced` - (Optional) L3 Mode Enhanced Performance. x-required L3 enhanced performance mode options. See [Perf Mode L3 Enhanced](#performance-enhancement-mode-perf-mode-l3-enhanced) below.

`perf_mode_l7_enhanced` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

<a id="performance-enhancement-mode-perf-mode-l3-enhanced"></a>

### Performance Enhancement Mode Perf Mode L3 Enhanced

`jumbo` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`no_jumbo` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

<a id="re-select"></a>

### RE Select

`geo_proximity` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`specific_re` - (Optional) Specific RE. Select specific REs. This is useful when a site needs to deterministically connect to a set of REs. A site will always be connected to 2 REs. See [Specific RE](#re-select-specific-re) below.

<a id="re-select-specific-re"></a>

### RE Select Specific RE

`primary_re` - (Optional) Primary RE Geography. Select primary RE for this site (`String`).

<a id="site-mesh-group-on-slo"></a>

### Site Mesh Group On Slo

`no_site_mesh_group` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`site_mesh_group` - (Optional) Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name. See [Site Mesh Group](#site-mesh-group-on-slo-site-mesh-group) below.

`sm_connection_public_ip` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`sm_connection_pvt_ip` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

<a id="site-mesh-group-on-slo-site-mesh-group"></a>

### Site Mesh Group On Slo Site Mesh Group

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

<a id="software-settings"></a>

### Software Settings

`os` - (Optional) Operating System Version. Select the F5XC Operating System Version for the site. By default, latest available OS Version will be used. Refer to release notes to find required released OS versions. See [OS](#software-settings-os) below.

`sw` - (Optional) F5XC Software Version. Select the F5XC Software Version for the site. By default, latest available F5XC Software Version will be used. Refer to release notes to find required released SW versions. See [Sw](#software-settings-sw) below.

<a id="software-settings-os"></a>

### Software Settings OS

`default_os_version` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`operating_system_version` - (Optional) Operating System Version. Specify a OS version to be used e.g. 9.2024.6 (`String`).

<a id="software-settings-sw"></a>

### Software Settings Sw

`default_sw_version` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`volterra_software_version` - (Optional) F5XC Software Version. Specify a F5XC Software Version to be used e.g. crt-20210329-1002 (`String`).

<a id="timeouts"></a>

### Timeouts

`create` - (Optional) A string that can be [parsed as a duration](`HTTPS://pkg.go.dev/time#ParseDuration`) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours) (`String`).

`delete` - (Optional) A string that can be [parsed as a duration](`HTTPS://pkg.go.dev/time#ParseDuration`) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are saved into state before the destroy operation occurs (`String`).

`read` - (Optional) A string that can be [parsed as a duration](`HTTPS://pkg.go.dev/time#ParseDuration`) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh is enabled (`String`).

`update` - (Optional) A string that can be [parsed as a duration](`HTTPS://pkg.go.dev/time#ParseDuration`) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours) (`String`).

<a id="upgrade-settings"></a>

### Upgrade Settings

`kubernetes_upgrade_drain` - (Optional) Node by Node Upgrade. Specify how worker nodes within a site will be upgraded. See [Kubernetes Upgrade Drain](#upgrade-settings-kubernetes-upgrade-drain) below.

<a id="upgrade-settings-kubernetes-upgrade-drain"></a>

### Upgrade Settings Kubernetes Upgrade Drain

`disable_upgrade_drain` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`enable_upgrade_drain` - (Optional) Enable Node by Node Upgrade. Specify batch upgrade settings for worker nodes within a site. See [Enable Upgrade Drain](#upgrade-settings-kubernetes-upgrade-drain-enable-upgrade-drain) below.

<a id="upgrade-settings-kubernetes-upgrade-drain-enable-upgrade-drain"></a>

### Upgrade Settings Kubernetes Upgrade Drain Enable Upgrade Drain

`disable_vega_upgrade_mode` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`drain_max_unavailable_node_count` - (Optional) Node Batch Size Count (`Number`).

`drain_node_timeout` - (Optional) Upgrade Wait Time. Seconds to wait before initiating upgrade on the next set of nodes. Setting it to 0 will wait indefinitely for all services on nodes to be upgraded gracefully before proceeding to the next set of nodes. (Warning: It may block upgrade if services on a node cannot be gracefully upgraded. It is recommended to use the default value) (`Number`).

`enable_vega_upgrade_mode` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

<a id="vmware"></a>

### Vmware

`not_managed` - (Optional) List of Nodes. This section will show nodes associated with this site. Note: For sites that are not orchestrated by F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it will be shown in this section. See [Not Managed](#vmware-not-managed) below.

<a id="vmware-not-managed"></a>

### Vmware Not Managed

`node_list` - (Optional) Nodes. This section will show nodes associated with this site. Note: For sites that are not orchestrated by F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it will be shown in this section. See [Node List](#vmware-not-managed-node-list) below.

<a id="vmware-not-managed-node-list"></a>

### Vmware Not Managed Node List

`hostname` - (Optional) Hostname. Hostname for this Node (`String`).

`interface_list` - (Optional) Interfaces. Manage interfaces belonging to this node (`Block`).

`public_ip` - (Optional) Public IP. Public IP for this Node (`String`).

`type` - (Optional) Type. Type for this Node, can be Control or Worker (`String`).
