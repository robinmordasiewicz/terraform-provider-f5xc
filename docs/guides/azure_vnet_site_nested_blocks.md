---
page_title: "Azure Vnet Site Nested Blocks - f5xc Provider"
subcategory: "Sites"
description: |-
  Nested block reference for the Azure Vnet Site resource.
---

# Azure Vnet Site Nested Blocks

This page contains detailed documentation for nested blocks in the `f5xc_azure_vnet_site` resource.

For the main resource documentation, see [f5xc_azure_vnet_site](/docs/resources/azure_vnet_site).

## Contents

- [admin-password](#admin-password)
- [admin-password-blindfold-secret-info](#admin-password-blindfold-secret-info)
- [admin-password-clear-secret-info](#admin-password-clear-secret-info)
- [azure-cred](#azure-cred)
- [blocked-services](#blocked-services)
- [blocked-services-blocked-sevice](#blocked-services-blocked-sevice)
- [coordinates](#coordinates)
- [custom-dns](#custom-dns)
- [ingress-egress-gw](#ingress-egress-gw)
- [ingress-egress-gw-accelerated-networking](#ingress-egress-gw-accelerated-networking)
- [ingress-egress-gw-active-enhanced-firewall-policies](#ingress-egress-gw-active-enhanced-firewall-policies)
- [ingress-egress-gw-active-enhanced-firewall-policies-enhanced-firewall-policies](#ingress-egress-gw-active-enhanced-firewall-policies-enhanced-firewall-policies)
- [ingress-egress-gw-active-forward-proxy-policies](#ingress-egress-gw-active-forward-proxy-policies)
- [ingress-egress-gw-active-forward-proxy-policies-forward-proxy-policies](#ingress-egress-gw-active-forward-proxy-policies-forward-proxy-policies)
- [ingress-egress-gw-active-network-policies](#ingress-egress-gw-active-network-policies)
- [ingress-egress-gw-active-network-policies-network-policies](#ingress-egress-gw-active-network-policies-network-policies)
- [ingress-egress-gw-az-nodes](#ingress-egress-gw-az-nodes)
- [ingress-egress-gw-az-nodes-inside-subnet](#ingress-egress-gw-az-nodes-inside-subnet)
- [ingress-egress-gw-az-nodes-outside-subnet](#ingress-egress-gw-az-nodes-outside-subnet)
- [ingress-egress-gw-dc-cluster-group-inside-vn](#ingress-egress-gw-dc-cluster-group-inside-vn)
- [ingress-egress-gw-dc-cluster-group-outside-vn](#ingress-egress-gw-dc-cluster-group-outside-vn)
- [ingress-egress-gw-global-network-list](#ingress-egress-gw-global-network-list)
- [ingress-egress-gw-global-network-list-global-network-connections](#ingress-egress-gw-global-network-list-global-network-connections)
- [ingress-egress-gw-hub](#ingress-egress-gw-hub)
- [ingress-egress-gw-hub-express-route-enabled](#ingress-egress-gw-hub-express-route-enabled)
- [ingress-egress-gw-hub-spoke-vnets](#ingress-egress-gw-hub-spoke-vnets)
- [ingress-egress-gw-inside-static-routes](#ingress-egress-gw-inside-static-routes)
- [ingress-egress-gw-inside-static-routes-static-route-list](#ingress-egress-gw-inside-static-routes-static-route-list)
- [ingress-egress-gw-outside-static-routes](#ingress-egress-gw-outside-static-routes)
- [ingress-egress-gw-outside-static-routes-static-route-list](#ingress-egress-gw-outside-static-routes-static-route-list)
- [ingress-egress-gw-performance-enhancement-mode](#ingress-egress-gw-performance-enhancement-mode)
- [ingress-egress-gw-performance-enhancement-mode-perf-mode-l3-enhanced](#ingress-egress-gw-performance-enhancement-mode-perf-mode-l3-enhanced)
- [ingress-egress-gw-ar](#ingress-egress-gw-ar)
- [ingress-egress-gw-ar-accelerated-networking](#ingress-egress-gw-ar-accelerated-networking)
- [ingress-egress-gw-ar-active-enhanced-firewall-policies](#ingress-egress-gw-ar-active-enhanced-firewall-policies)
- [ingress-egress-gw-ar-active-enhanced-firewall-policies-enhanced-firewall-policies](#ingress-egress-gw-ar-active-enhanced-firewall-policies-enhanced-firewall-policies)
- [ingress-egress-gw-ar-active-forward-proxy-policies](#ingress-egress-gw-ar-active-forward-proxy-policies)
- [ingress-egress-gw-ar-active-forward-proxy-policies-forward-proxy-policies](#ingress-egress-gw-ar-active-forward-proxy-policies-forward-proxy-policies)
- [ingress-egress-gw-ar-active-network-policies](#ingress-egress-gw-ar-active-network-policies)
- [ingress-egress-gw-ar-active-network-policies-network-policies](#ingress-egress-gw-ar-active-network-policies-network-policies)
- [ingress-egress-gw-ar-dc-cluster-group-inside-vn](#ingress-egress-gw-ar-dc-cluster-group-inside-vn)
- [ingress-egress-gw-ar-dc-cluster-group-outside-vn](#ingress-egress-gw-ar-dc-cluster-group-outside-vn)
- [ingress-egress-gw-ar-global-network-list](#ingress-egress-gw-ar-global-network-list)
- [ingress-egress-gw-ar-global-network-list-global-network-connections](#ingress-egress-gw-ar-global-network-list-global-network-connections)
- [ingress-egress-gw-ar-hub](#ingress-egress-gw-ar-hub)
- [ingress-egress-gw-ar-hub-express-route-enabled](#ingress-egress-gw-ar-hub-express-route-enabled)
- [ingress-egress-gw-ar-hub-spoke-vnets](#ingress-egress-gw-ar-hub-spoke-vnets)
- [ingress-egress-gw-ar-inside-static-routes](#ingress-egress-gw-ar-inside-static-routes)
- [ingress-egress-gw-ar-inside-static-routes-static-route-list](#ingress-egress-gw-ar-inside-static-routes-static-route-list)
- [ingress-egress-gw-ar-node](#ingress-egress-gw-ar-node)
- [ingress-egress-gw-ar-node-inside-subnet](#ingress-egress-gw-ar-node-inside-subnet)
- [ingress-egress-gw-ar-node-outside-subnet](#ingress-egress-gw-ar-node-outside-subnet)
- [ingress-egress-gw-ar-outside-static-routes](#ingress-egress-gw-ar-outside-static-routes)
- [ingress-egress-gw-ar-outside-static-routes-static-route-list](#ingress-egress-gw-ar-outside-static-routes-static-route-list)
- [ingress-egress-gw-ar-performance-enhancement-mode](#ingress-egress-gw-ar-performance-enhancement-mode)
- [ingress-egress-gw-ar-performance-enhancement-mode-perf-mode-l3-enhanced](#ingress-egress-gw-ar-performance-enhancement-mode-perf-mode-l3-enhanced)
- [ingress-gw](#ingress-gw)
- [ingress-gw-accelerated-networking](#ingress-gw-accelerated-networking)
- [ingress-gw-az-nodes](#ingress-gw-az-nodes)
- [ingress-gw-az-nodes-local-subnet](#ingress-gw-az-nodes-local-subnet)
- [ingress-gw-performance-enhancement-mode](#ingress-gw-performance-enhancement-mode)
- [ingress-gw-performance-enhancement-mode-perf-mode-l3-enhanced](#ingress-gw-performance-enhancement-mode-perf-mode-l3-enhanced)
- [ingress-gw-ar](#ingress-gw-ar)
- [ingress-gw-ar-accelerated-networking](#ingress-gw-ar-accelerated-networking)
- [ingress-gw-ar-node](#ingress-gw-ar-node)
- [ingress-gw-ar-node-local-subnet](#ingress-gw-ar-node-local-subnet)
- [ingress-gw-ar-performance-enhancement-mode](#ingress-gw-ar-performance-enhancement-mode)
- [ingress-gw-ar-performance-enhancement-mode-perf-mode-l3-enhanced](#ingress-gw-ar-performance-enhancement-mode-perf-mode-l3-enhanced)
- [kubernetes-upgrade-drain](#kubernetes-upgrade-drain)
- [kubernetes-upgrade-drain-enable-upgrade-drain](#kubernetes-upgrade-drain-enable-upgrade-drain)
- [log-receiver](#log-receiver)
- [offline-survivability-mode](#offline-survivability-mode)
- [os](#os)
- [sw](#sw)
- [timeouts](#timeouts)
- [vnet](#vnet)
- [vnet-existing-vnet](#vnet-existing-vnet)
- [vnet-new-vnet](#vnet-new-vnet)
- [voltstack-cluster](#voltstack-cluster)
- [voltstack-cluster-accelerated-networking](#voltstack-cluster-accelerated-networking)
- [voltstack-cluster-active-enhanced-firewall-policies](#voltstack-cluster-active-enhanced-firewall-policies)
- [voltstack-cluster-active-enhanced-firewall-policies-enhanced-firewall-policies](#voltstack-cluster-active-enhanced-firewall-policies-enhanced-firewall-policies)
- [voltstack-cluster-active-forward-proxy-policies](#voltstack-cluster-active-forward-proxy-policies)
- [voltstack-cluster-active-forward-proxy-policies-forward-proxy-policies](#voltstack-cluster-active-forward-proxy-policies-forward-proxy-policies)
- [voltstack-cluster-active-network-policies](#voltstack-cluster-active-network-policies)
- [voltstack-cluster-active-network-policies-network-policies](#voltstack-cluster-active-network-policies-network-policies)
- [voltstack-cluster-az-nodes](#voltstack-cluster-az-nodes)
- [voltstack-cluster-az-nodes-local-subnet](#voltstack-cluster-az-nodes-local-subnet)
- [voltstack-cluster-dc-cluster-group](#voltstack-cluster-dc-cluster-group)
- [voltstack-cluster-global-network-list](#voltstack-cluster-global-network-list)
- [voltstack-cluster-global-network-list-global-network-connections](#voltstack-cluster-global-network-list-global-network-connections)
- [voltstack-cluster-k8s-cluster](#voltstack-cluster-k8s-cluster)
- [voltstack-cluster-outside-static-routes](#voltstack-cluster-outside-static-routes)
- [voltstack-cluster-outside-static-routes-static-route-list](#voltstack-cluster-outside-static-routes-static-route-list)
- [voltstack-cluster-storage-class-list](#voltstack-cluster-storage-class-list)
- [voltstack-cluster-storage-class-list-storage-classes](#voltstack-cluster-storage-class-list-storage-classes)
- [voltstack-cluster-ar](#voltstack-cluster-ar)
- [voltstack-cluster-ar-accelerated-networking](#voltstack-cluster-ar-accelerated-networking)
- [voltstack-cluster-ar-active-enhanced-firewall-policies](#voltstack-cluster-ar-active-enhanced-firewall-policies)
- [voltstack-cluster-ar-active-enhanced-firewall-policies-enhanced-firewall-policies](#voltstack-cluster-ar-active-enhanced-firewall-policies-enhanced-firewall-policies)
- [voltstack-cluster-ar-active-forward-proxy-policies](#voltstack-cluster-ar-active-forward-proxy-policies)
- [voltstack-cluster-ar-active-forward-proxy-policies-forward-proxy-policies](#voltstack-cluster-ar-active-forward-proxy-policies-forward-proxy-policies)
- [voltstack-cluster-ar-active-network-policies](#voltstack-cluster-ar-active-network-policies)
- [voltstack-cluster-ar-active-network-policies-network-policies](#voltstack-cluster-ar-active-network-policies-network-policies)
- [voltstack-cluster-ar-dc-cluster-group](#voltstack-cluster-ar-dc-cluster-group)
- [voltstack-cluster-ar-global-network-list](#voltstack-cluster-ar-global-network-list)
- [voltstack-cluster-ar-global-network-list-global-network-connections](#voltstack-cluster-ar-global-network-list-global-network-connections)
- [voltstack-cluster-ar-k8s-cluster](#voltstack-cluster-ar-k8s-cluster)
- [voltstack-cluster-ar-node](#voltstack-cluster-ar-node)
- [voltstack-cluster-ar-node-local-subnet](#voltstack-cluster-ar-node-local-subnet)
- [voltstack-cluster-ar-outside-static-routes](#voltstack-cluster-ar-outside-static-routes)
- [voltstack-cluster-ar-outside-static-routes-static-route-list](#voltstack-cluster-ar-outside-static-routes-static-route-list)
- [voltstack-cluster-ar-storage-class-list](#voltstack-cluster-ar-storage-class-list)
- [voltstack-cluster-ar-storage-class-list-storage-classes](#voltstack-cluster-ar-storage-class-list-storage-classes)

---

<a id="admin-password"></a>

### Admin Password

`blindfold_secret_info` - (Optional) Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management. See [Blindfold Secret Info](#admin-password-blindfold-secret-info) below.

`clear_secret_info` - (Optional) In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted. See [Clear Secret Info](#admin-password-clear-secret-info) below.

<a id="admin-password-blindfold-secret-info"></a>

### Admin Password Blindfold Secret Info

`decryption_provider` - (Optional) Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service (`String`).

`location` - (Optional) Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location (`String`).

`store_provider` - (Optional) Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:/// (`String`).

<a id="admin-password-clear-secret-info"></a>

### Admin Password Clear Secret Info

`provider_ref` - (Optional) Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:/// (`String`).

`url` - (Optional) URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding (`String`).

<a id="azure-cred"></a>

### Azure Cred

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

<a id="blocked-services"></a>

### Blocked Services

`blocked_sevice` - (Optional) Disable Node Local Services. See [Blocked Sevice](#blocked-services-blocked-sevice) below.

<a id="blocked-services-blocked-sevice"></a>

### Blocked Services Blocked Sevice

`dns` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`network_type` - (Optional) Virtual Network Type. Different types of virtual networks understood by the system Virtual-network of type VIRTUAL_NETWORK_SITE_LOCAL provides connectivity to public (outside) network. This is an insecure network and is connected to public internet via NAT Gateways/firwalls Virtual-network of this type is local to every site. Two virtual networks of this type on different sites are neither related nor connected. Constraints: There can be atmost one virtual network of this type in a given site. This network type is supported on CE sites. This network is created automatically and present on all sites Virtual-network of type VIRTUAL_NETWORK_SITE_LOCAL_INSIDE is a private network inside site. It is a secure network and is not connected to public network. Virtual-network of this type is local to every site. Two virtual networks of this type on different sites are neither related nor connected. Constraints: There can be atmost one virtual network of this type in a given site. This network type is supported on CE sites. This network is created during provisioning of site User defined per-site virtual network. Scope of this virtual network is limited to the site. This is not yet supported Virtual-network of type VIRTUAL_NETWORK_PUBLIC directly conects to the public internet. Virtual-network of this type is local to every site. Two virtual networks of this type on different sites are neither related nor connected. Constraints: There can be atmost one virtual network of this type in a given site. This network type is supported on RE sites only It is an internally created by the system. They must not be created by user Virtual Neworks with global scope across different sites in F5XC domain. An example global virtual-network called 'AIN Network' is created for every tenant. for volterra fabric Constraints: It is currently only supported as internally created by the system. vK8s service network for a given tenant. Used to advertise a virtual host only to vk8s pods for that tenant Constraints: It is an internally created by the system. Must not be created by user VER internal network for the site. It can only be used for virtual hosts with SMA_PROXY type proxy Constraints: It is an internally created by the system. Must not be created by user Virtual-network of type VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE represents both VIRTUAL_NETWORK_SITE_LOCAL and VIRTUAL_NETWORK_SITE_LOCAL_INSIDE Constraints: This network type is only meaningful in an advertise policy When virtual-network of type VIRTUAL_NETWORK_IP_AUTO is selected for an endpoint, VER will try to determine the network based on the provided IP address Constraints: This network type is only meaningful in an endpoint VoltADN Private Network is used on volterra RE(s) to connect to customer private networks This network is created by opening a support ticket This network is per site srv6 network VER IP Fabric network for the site. This Virtual network type is used for exposing virtual host on IP Fabric network on the VER site or for endpoint in IP Fabric network Constraints: It is an internally created by the system. Must not be created by user Network internally created for a segment Constraints: It is an internally created by the system. Must not be created by user. Possible values are `VIRTUAL_NETWORK_SITE_LOCAL`, `VIRTUAL_NETWORK_SITE_LOCAL_INSIDE`, `VIRTUAL_NETWORK_PER_SITE`, `VIRTUAL_NETWORK_PUBLIC`, `VIRTUAL_NETWORK_GLOBAL`, `VIRTUAL_NETWORK_SITE_SERVICE`, `VIRTUAL_NETWORK_VER_INTERNAL`, `VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE`, `VIRTUAL_NETWORK_IP_AUTO`, `VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK`, `VIRTUAL_NETWORK_SRV6_NETWORK`, `VIRTUAL_NETWORK_IP_FABRIC`, `VIRTUAL_NETWORK_SEGMENT`. Defaults to `VIRTUAL_NETWORK_SITE_LOCAL` (`String`).

`ssh` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`web_user_interface` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

<a id="coordinates"></a>

### Coordinates

`latitude` - (Optional) Latitude. Latitude of the site location (`Number`).

`longitude` - (Optional) Longitude. longitude of site location (`Number`).

<a id="custom-dns"></a>

### Custom DNS

`inside_nameserver` - (Optional) DNS Server for Inside Network. Optional DNS server IP to be used for name resolution in inside network (`String`).

`outside_nameserver` - (Optional) DNS Server for Outside Network. Optional DNS server IP to be used for name resolution in outside network (`String`).

<a id="ingress-egress-gw"></a>

### Ingress Egress Gw

`accelerated_networking` - (Optional) Accelerated Networking Type. x-required Accelerated Networking to reduce Latency, When Mode is toggled, traffic disruption will be seen. See [Accelerated Networking](#ingress-egress-gw-accelerated-networking) below.

`active_enhanced_firewall_policies` - (Optional) Active Enhanced Network Policies Type. List of Enhanced Firewall Policies These policies use session-based rules and provide all options available under firewall policies with an additional option for service insertion. See [Active Enhanced Firewall Policies](#ingress-egress-gw-active-enhanced-firewall-policies) below.

`active_forward_proxy_policies` - (Optional) Active Forward Proxy Policies Type. Ordered List of Forward Proxy Policies active. See [Active Forward Proxy Policies](#ingress-egress-gw-active-forward-proxy-policies) below.

`active_network_policies` - (Optional) Active Firewall Policies Type. List of firewall policy views. See [Active Network Policies](#ingress-egress-gw-active-network-policies) below.

`az_nodes` - (Optional) Ingress/Egress Gateway (two Interface) Nodes in AZ. Only Single AZ or Three AZ(s) nodes are supported currently. See [Az Nodes](#ingress-egress-gw-az-nodes) below.

`azure_certified_hw` - (Optional) Azure Certified Hardware. Name for Azure certified hardware (`String`).

`dc_cluster_group_inside_vn` - (Optional) Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name. See [Dc Cluster Group Inside Vn](#ingress-egress-gw-dc-cluster-group-inside-vn) below.

`dc_cluster_group_outside_vn` - (Optional) Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name. See [Dc Cluster Group Outside Vn](#ingress-egress-gw-dc-cluster-group-outside-vn) below.

`forward_proxy_allow_all` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`global_network_list` - (Optional) Global Network Connection List. List of global network connections. See [Global Network List](#ingress-egress-gw-global-network-list) below.

`hub` - (Optional) Hub VNet type. Hub VNet type. See [Hub](#ingress-egress-gw-hub) below.

`inside_static_routes` - (Optional) Static Route List Type. List of static routes. See [Inside Static Routes](#ingress-egress-gw-inside-static-routes) below.

`no_dc_cluster_group` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`no_forward_proxy` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`no_global_network` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`no_inside_static_routes` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`no_network_policy` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`no_outside_static_routes` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`not_hub` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`outside_static_routes` - (Optional) Static Route List Type. List of static routes. See [Outside Static Routes](#ingress-egress-gw-outside-static-routes) below.

`performance_enhancement_mode` - (Optional) Performance Enhancement Mode. x-required Optimize the site for L3 or L7 traffic processing. L7 optimized is the default. See [Performance Enhancement Mode](#ingress-egress-gw-performance-enhancement-mode) below.

`sm_connection_public_ip` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`sm_connection_pvt_ip` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

<a id="ingress-egress-gw-accelerated-networking"></a>

### Ingress Egress Gw Accelerated Networking

`disable` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`enable` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

<a id="ingress-egress-gw-active-enhanced-firewall-policies"></a>

### Ingress Egress Gw Active Enhanced Firewall Policies

`enhanced_firewall_policies` - (Optional) Enhanced Firewall Policy. Ordered List of Enhanced Firewall Policies active. See [Enhanced Firewall Policies](#ingress-egress-gw-active-enhanced-firewall-policies-enhanced-firewall-policies) below.

<a id="ingress-egress-gw-active-enhanced-firewall-policies-enhanced-firewall-policies"></a>

### Ingress Egress Gw Active Enhanced Firewall Policies Enhanced Firewall Policies

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

<a id="ingress-egress-gw-active-forward-proxy-policies"></a>

### Ingress Egress Gw Active Forward Proxy Policies

`forward_proxy_policies` - (Optional) Forward Proxy Policies. Ordered List of Forward Proxy Policies active. See [Forward Proxy Policies](#ingress-egress-gw-active-forward-proxy-policies-forward-proxy-policies) below.

<a id="ingress-egress-gw-active-forward-proxy-policies-forward-proxy-policies"></a>

### Ingress Egress Gw Active Forward Proxy Policies Forward Proxy Policies

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

<a id="ingress-egress-gw-active-network-policies"></a>

### Ingress Egress Gw Active Network Policies

`network_policies` - (Optional) Firewall Policy. Ordered List of Firewall Policies active for this network firewall. See [Network Policies](#ingress-egress-gw-active-network-policies-network-policies) below.

<a id="ingress-egress-gw-active-network-policies-network-policies"></a>

### Ingress Egress Gw Active Network Policies Network Policies

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

<a id="ingress-egress-gw-az-nodes"></a>

### Ingress Egress Gw Az Nodes

`azure_az` - (Optional) Azure Availability Zone. A zone depicting a grouping of datacenters within an Azure region. Expecting numeric input (`String`).

`inside_subnet` - (Optional) Azure Subnet. Parameters for Azure subnet. See [Inside Subnet](#ingress-egress-gw-az-nodes-inside-subnet) below.

`outside_subnet` - (Optional) Azure Subnet. Parameters for Azure subnet. See [Outside Subnet](#ingress-egress-gw-az-nodes-outside-subnet) below.

<a id="ingress-egress-gw-az-nodes-inside-subnet"></a>

### Ingress Egress Gw Az Nodes Inside Subnet

`subnet` - (Optional) Azure Subnet. Parameters for Azure subnet (`Block`).

`subnet_param` - (Optional) New Cloud Subnet Parameters. Parameters for creating a new cloud subnet (`Block`).

<a id="ingress-egress-gw-az-nodes-outside-subnet"></a>

### Ingress Egress Gw Az Nodes Outside Subnet

`subnet` - (Optional) Azure Subnet. Parameters for Azure subnet (`Block`).

`subnet_param` - (Optional) New Cloud Subnet Parameters. Parameters for creating a new cloud subnet (`Block`).

<a id="ingress-egress-gw-dc-cluster-group-inside-vn"></a>

### Ingress Egress Gw Dc Cluster Group Inside Vn

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

<a id="ingress-egress-gw-dc-cluster-group-outside-vn"></a>

### Ingress Egress Gw Dc Cluster Group Outside Vn

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

<a id="ingress-egress-gw-global-network-list"></a>

### Ingress Egress Gw Global Network List

`global_network_connections` - (Optional) Global Network Connections. Global network connections. See [Global Network Connections](#ingress-egress-gw-global-network-list-global-network-connections) below.

<a id="ingress-egress-gw-global-network-list-global-network-connections"></a>

### Ingress Egress Gw Global Network List Global Network Connections

`sli_to_global_dr` - (Optional) Global Network. Global network reference for direct connection (`Block`).

`slo_to_global_dr` - (Optional) Global Network. Global network reference for direct connection (`Block`).

<a id="ingress-egress-gw-hub"></a>

### Ingress Egress Gw Hub

`express_route_disabled` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`express_route_enabled` - (Optional) Express Route Configuration. Express Route Configuration. See [Express Route Enabled](#ingress-egress-gw-hub-express-route-enabled) below.

`spoke_vnets` - (Optional) Spoke VNet Peering (Legacy). Spoke VNet Peering. See [Spoke Vnets](#ingress-egress-gw-hub-spoke-vnets) below.

<a id="ingress-egress-gw-hub-express-route-enabled"></a>

### Ingress Egress Gw Hub Express Route Enabled

`advertise_to_route_server` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`auto_asn` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`connections` - (Optional) Connections. Add the ExpressRoute Circuit Connections to this site (`Block`).

`custom_asn` - (Optional) Custom ASN. Set custom ASN for F5XC Site (`Number`).

`do_not_advertise_to_route_server` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`gateway_subnet` - (Optional) Azure Subnet. Parameters for Azure subnet (`Block`).

`route_server_subnet` - (Optional) Azure Subnet. Parameters for Azure subnet (`Block`).

`site_registration_over_express_route` - (Optional) CloudLink ADN Network Config (`Block`).

`site_registration_over_internet` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`sku_ergw1az` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`sku_ergw2az` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`sku_high_perf` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`sku_standard` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

<a id="ingress-egress-gw-hub-spoke-vnets"></a>

### Ingress Egress Gw Hub Spoke Vnets

`auto` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`labels` - (Optional) Labels For VNets Peering. Add Labels for each of the VNets peered with transit VNet, these labels can be used in firewall policy These labels used must be from known key and label defined in shared namespace (`Block`).

`manual` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`vnet` - (Optional) Azure Existing Vnet Type. Resource group and name of existing Azure Vnet (`Block`).

<a id="ingress-egress-gw-inside-static-routes"></a>

### Ingress Egress Gw Inside Static Routes

`static_route_list` - (Optional) List of Static Routes. List of Static routes. See [Static Route List](#ingress-egress-gw-inside-static-routes-static-route-list) below.

<a id="ingress-egress-gw-inside-static-routes-static-route-list"></a>

### Ingress Egress Gw Inside Static Routes Static Route List

`custom_static_route` - (Optional) Static Route. Defines a static route, configuring a list of prefixes and a next-hop to be used for them (`Block`).

`simple_static_route` - (Optional) Simple Static Route. Use simple static route for prefix pointing to single interface in the network (`String`).

<a id="ingress-egress-gw-outside-static-routes"></a>

### Ingress Egress Gw Outside Static Routes

`static_route_list` - (Optional) List of Static Routes. List of Static routes. See [Static Route List](#ingress-egress-gw-outside-static-routes-static-route-list) below.

<a id="ingress-egress-gw-outside-static-routes-static-route-list"></a>

### Ingress Egress Gw Outside Static Routes Static Route List

`custom_static_route` - (Optional) Static Route. Defines a static route, configuring a list of prefixes and a next-hop to be used for them (`Block`).

`simple_static_route` - (Optional) Simple Static Route. Use simple static route for prefix pointing to single interface in the network (`String`).

<a id="ingress-egress-gw-performance-enhancement-mode"></a>

### Ingress Egress Gw Performance Enhancement Mode

`perf_mode_l3_enhanced` - (Optional) L3 Mode Enhanced Performance. x-required L3 enhanced performance mode options. See [Perf Mode L3 Enhanced](#ingress-egress-gw-performance-enhancement-mode-perf-mode-l3-enhanced) below.

`perf_mode_l7_enhanced` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

<a id="ingress-egress-gw-performance-enhancement-mode-perf-mode-l3-enhanced"></a>

### Ingress Egress Gw Performance Enhancement Mode Perf Mode L3 Enhanced

`jumbo` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`no_jumbo` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

<a id="ingress-egress-gw-ar"></a>

### Ingress Egress Gw Ar

`accelerated_networking` - (Optional) Accelerated Networking Type. x-required Accelerated Networking to reduce Latency, When Mode is toggled, traffic disruption will be seen. See [Accelerated Networking](#ingress-egress-gw-ar-accelerated-networking) below.

`active_enhanced_firewall_policies` - (Optional) Active Enhanced Network Policies Type. List of Enhanced Firewall Policies These policies use session-based rules and provide all options available under firewall policies with an additional option for service insertion. See [Active Enhanced Firewall Policies](#ingress-egress-gw-ar-active-enhanced-firewall-policies) below.

`active_forward_proxy_policies` - (Optional) Active Forward Proxy Policies Type. Ordered List of Forward Proxy Policies active. See [Active Forward Proxy Policies](#ingress-egress-gw-ar-active-forward-proxy-policies) below.

`active_network_policies` - (Optional) Active Firewall Policies Type. List of firewall policy views. See [Active Network Policies](#ingress-egress-gw-ar-active-network-policies) below.

`azure_certified_hw` - (Optional) Azure Certified Hardware. Name for Azure certified hardware (`String`).

`dc_cluster_group_inside_vn` - (Optional) Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name. See [Dc Cluster Group Inside Vn](#ingress-egress-gw-ar-dc-cluster-group-inside-vn) below.

`dc_cluster_group_outside_vn` - (Optional) Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name. See [Dc Cluster Group Outside Vn](#ingress-egress-gw-ar-dc-cluster-group-outside-vn) below.

`forward_proxy_allow_all` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`global_network_list` - (Optional) Global Network Connection List. List of global network connections. See [Global Network List](#ingress-egress-gw-ar-global-network-list) below.

`hub` - (Optional) Hub VNet type. Hub VNet type. See [Hub](#ingress-egress-gw-ar-hub) below.

`inside_static_routes` - (Optional) Static Route List Type. List of static routes. See [Inside Static Routes](#ingress-egress-gw-ar-inside-static-routes) below.

`no_dc_cluster_group` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`no_forward_proxy` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`no_global_network` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`no_inside_static_routes` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`no_network_policy` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`no_outside_static_routes` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`node` - (Optional) Two Interface Node. Parameters for creating two interface Node in one AZ. See [Node](#ingress-egress-gw-ar-node) below.

`not_hub` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`outside_static_routes` - (Optional) Static Route List Type. List of static routes. See [Outside Static Routes](#ingress-egress-gw-ar-outside-static-routes) below.

`performance_enhancement_mode` - (Optional) Performance Enhancement Mode. x-required Optimize the site for L3 or L7 traffic processing. L7 optimized is the default. See [Performance Enhancement Mode](#ingress-egress-gw-ar-performance-enhancement-mode) below.

`sm_connection_public_ip` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`sm_connection_pvt_ip` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

<a id="ingress-egress-gw-ar-accelerated-networking"></a>

### Ingress Egress Gw Ar Accelerated Networking

`disable` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`enable` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

<a id="ingress-egress-gw-ar-active-enhanced-firewall-policies"></a>

### Ingress Egress Gw Ar Active Enhanced Firewall Policies

`enhanced_firewall_policies` - (Optional) Enhanced Firewall Policy. Ordered List of Enhanced Firewall Policies active. See [Enhanced Firewall Policies](#ingress-egress-gw-ar-active-enhanced-firewall-policies-enhanced-firewall-policies) below.

<a id="ingress-egress-gw-ar-active-enhanced-firewall-policies-enhanced-firewall-policies"></a>

### Ingress Egress Gw Ar Active Enhanced Firewall Policies Enhanced Firewall Policies

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

<a id="ingress-egress-gw-ar-active-forward-proxy-policies"></a>

### Ingress Egress Gw Ar Active Forward Proxy Policies

`forward_proxy_policies` - (Optional) Forward Proxy Policies. Ordered List of Forward Proxy Policies active. See [Forward Proxy Policies](#ingress-egress-gw-ar-active-forward-proxy-policies-forward-proxy-policies) below.

<a id="ingress-egress-gw-ar-active-forward-proxy-policies-forward-proxy-policies"></a>

### Ingress Egress Gw Ar Active Forward Proxy Policies Forward Proxy Policies

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

<a id="ingress-egress-gw-ar-active-network-policies"></a>

### Ingress Egress Gw Ar Active Network Policies

`network_policies` - (Optional) Firewall Policy. Ordered List of Firewall Policies active for this network firewall. See [Network Policies](#ingress-egress-gw-ar-active-network-policies-network-policies) below.

<a id="ingress-egress-gw-ar-active-network-policies-network-policies"></a>

### Ingress Egress Gw Ar Active Network Policies Network Policies

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

<a id="ingress-egress-gw-ar-dc-cluster-group-inside-vn"></a>

### Ingress Egress Gw Ar Dc Cluster Group Inside Vn

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

<a id="ingress-egress-gw-ar-dc-cluster-group-outside-vn"></a>

### Ingress Egress Gw Ar Dc Cluster Group Outside Vn

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

<a id="ingress-egress-gw-ar-global-network-list"></a>

### Ingress Egress Gw Ar Global Network List

`global_network_connections` - (Optional) Global Network Connections. Global network connections. See [Global Network Connections](#ingress-egress-gw-ar-global-network-list-global-network-connections) below.

<a id="ingress-egress-gw-ar-global-network-list-global-network-connections"></a>

### Ingress Egress Gw Ar Global Network List Global Network Connections

`sli_to_global_dr` - (Optional) Global Network. Global network reference for direct connection (`Block`).

`slo_to_global_dr` - (Optional) Global Network. Global network reference for direct connection (`Block`).

<a id="ingress-egress-gw-ar-hub"></a>

### Ingress Egress Gw Ar Hub

`express_route_disabled` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`express_route_enabled` - (Optional) Express Route Configuration. Express Route Configuration. See [Express Route Enabled](#ingress-egress-gw-ar-hub-express-route-enabled) below.

`spoke_vnets` - (Optional) Spoke VNet Peering (Legacy). Spoke VNet Peering. See [Spoke Vnets](#ingress-egress-gw-ar-hub-spoke-vnets) below.

<a id="ingress-egress-gw-ar-hub-express-route-enabled"></a>

### Ingress Egress Gw Ar Hub Express Route Enabled

`advertise_to_route_server` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`auto_asn` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`connections` - (Optional) Connections. Add the ExpressRoute Circuit Connections to this site (`Block`).

`custom_asn` - (Optional) Custom ASN. Set custom ASN for F5XC Site (`Number`).

`do_not_advertise_to_route_server` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`gateway_subnet` - (Optional) Azure Subnet. Parameters for Azure subnet (`Block`).

`route_server_subnet` - (Optional) Azure Subnet. Parameters for Azure subnet (`Block`).

`site_registration_over_express_route` - (Optional) CloudLink ADN Network Config (`Block`).

`site_registration_over_internet` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`sku_ergw1az` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`sku_ergw2az` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`sku_high_perf` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`sku_standard` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

<a id="ingress-egress-gw-ar-hub-spoke-vnets"></a>

### Ingress Egress Gw Ar Hub Spoke Vnets

`auto` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`labels` - (Optional) Labels For VNets Peering. Add Labels for each of the VNets peered with transit VNet, these labels can be used in firewall policy These labels used must be from known key and label defined in shared namespace (`Block`).

`manual` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`vnet` - (Optional) Azure Existing Vnet Type. Resource group and name of existing Azure Vnet (`Block`).

<a id="ingress-egress-gw-ar-inside-static-routes"></a>

### Ingress Egress Gw Ar Inside Static Routes

`static_route_list` - (Optional) List of Static Routes. List of Static routes. See [Static Route List](#ingress-egress-gw-ar-inside-static-routes-static-route-list) below.

<a id="ingress-egress-gw-ar-inside-static-routes-static-route-list"></a>

### Ingress Egress Gw Ar Inside Static Routes Static Route List

`custom_static_route` - (Optional) Static Route. Defines a static route, configuring a list of prefixes and a next-hop to be used for them (`Block`).

`simple_static_route` - (Optional) Simple Static Route. Use simple static route for prefix pointing to single interface in the network (`String`).

<a id="ingress-egress-gw-ar-node"></a>

### Ingress Egress Gw Ar Node

`fault_domain` - (Optional) Number of fault domains. Namuber of fault domains to be used while creating the availability set (`Number`).

`inside_subnet` - (Optional) Azure Subnet. Parameters for Azure subnet. See [Inside Subnet](#ingress-egress-gw-ar-node-inside-subnet) below.

`node_number` - (Optional) Number of main nodes. Number of main nodes to create, either 1 or 3 (`Number`).

`outside_subnet` - (Optional) Azure Subnet. Parameters for Azure subnet. See [Outside Subnet](#ingress-egress-gw-ar-node-outside-subnet) below.

`update_domain` - (Optional) Number of update domains. Namuber of update domains to be used while creating the availability set (`Number`).

<a id="ingress-egress-gw-ar-node-inside-subnet"></a>

### Ingress Egress Gw Ar Node Inside Subnet

`subnet` - (Optional) Azure Subnet. Parameters for Azure subnet (`Block`).

`subnet_param` - (Optional) New Cloud Subnet Parameters. Parameters for creating a new cloud subnet (`Block`).

<a id="ingress-egress-gw-ar-node-outside-subnet"></a>

### Ingress Egress Gw Ar Node Outside Subnet

`subnet` - (Optional) Azure Subnet. Parameters for Azure subnet (`Block`).

`subnet_param` - (Optional) New Cloud Subnet Parameters. Parameters for creating a new cloud subnet (`Block`).

<a id="ingress-egress-gw-ar-outside-static-routes"></a>

### Ingress Egress Gw Ar Outside Static Routes

`static_route_list` - (Optional) List of Static Routes. List of Static routes. See [Static Route List](#ingress-egress-gw-ar-outside-static-routes-static-route-list) below.

<a id="ingress-egress-gw-ar-outside-static-routes-static-route-list"></a>

### Ingress Egress Gw Ar Outside Static Routes Static Route List

`custom_static_route` - (Optional) Static Route. Defines a static route, configuring a list of prefixes and a next-hop to be used for them (`Block`).

`simple_static_route` - (Optional) Simple Static Route. Use simple static route for prefix pointing to single interface in the network (`String`).

<a id="ingress-egress-gw-ar-performance-enhancement-mode"></a>

### Ingress Egress Gw Ar Performance Enhancement Mode

`perf_mode_l3_enhanced` - (Optional) L3 Mode Enhanced Performance. x-required L3 enhanced performance mode options. See [Perf Mode L3 Enhanced](#ingress-egress-gw-ar-performance-enhancement-mode-perf-mode-l3-enhanced) below.

`perf_mode_l7_enhanced` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

<a id="ingress-egress-gw-ar-performance-enhancement-mode-perf-mode-l3-enhanced"></a>

### Ingress Egress Gw Ar Performance Enhancement Mode Perf Mode L3 Enhanced

`jumbo` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`no_jumbo` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

<a id="ingress-gw"></a>

### Ingress Gw

`accelerated_networking` - (Optional) Accelerated Networking Type. x-required Accelerated Networking to reduce Latency, When Mode is toggled, traffic disruption will be seen. See [Accelerated Networking](#ingress-gw-accelerated-networking) below.

`az_nodes` - (Optional) Ingress Gateway (One Interface) Nodes in AZ. Only Single AZ or Three AZ(s) nodes are supported currently. See [Az Nodes](#ingress-gw-az-nodes) below.

`azure_certified_hw` - (Optional) Azure Certified Hardware. Name for Azure certified hardware (`String`).

`performance_enhancement_mode` - (Optional) Performance Enhancement Mode. x-required Optimize the site for L3 or L7 traffic processing. L7 optimized is the default. See [Performance Enhancement Mode](#ingress-gw-performance-enhancement-mode) below.

<a id="ingress-gw-accelerated-networking"></a>

### Ingress Gw Accelerated Networking

`disable` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`enable` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

<a id="ingress-gw-az-nodes"></a>

### Ingress Gw Az Nodes

`azure_az` - (Optional) Azure Availability Zone. A zone depicting a grouping of datacenters within an Azure region. Expecting numeric input (`String`).

`local_subnet` - (Optional) Azure Subnet. Parameters for Azure subnet. See [Local Subnet](#ingress-gw-az-nodes-local-subnet) below.

<a id="ingress-gw-az-nodes-local-subnet"></a>

### Ingress Gw Az Nodes Local Subnet

`subnet` - (Optional) Azure Subnet. Parameters for Azure subnet (`Block`).

`subnet_param` - (Optional) New Cloud Subnet Parameters. Parameters for creating a new cloud subnet (`Block`).

<a id="ingress-gw-performance-enhancement-mode"></a>

### Ingress Gw Performance Enhancement Mode

`perf_mode_l3_enhanced` - (Optional) L3 Mode Enhanced Performance. x-required L3 enhanced performance mode options. See [Perf Mode L3 Enhanced](#ingress-gw-performance-enhancement-mode-perf-mode-l3-enhanced) below.

`perf_mode_l7_enhanced` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

<a id="ingress-gw-performance-enhancement-mode-perf-mode-l3-enhanced"></a>

### Ingress Gw Performance Enhancement Mode Perf Mode L3 Enhanced

`jumbo` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`no_jumbo` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

<a id="ingress-gw-ar"></a>

### Ingress Gw Ar

`accelerated_networking` - (Optional) Accelerated Networking Type. x-required Accelerated Networking to reduce Latency, When Mode is toggled, traffic disruption will be seen. See [Accelerated Networking](#ingress-gw-ar-accelerated-networking) below.

`azure_certified_hw` - (Optional) Azure Certified Hardware. Name for Azure certified hardware (`String`).

`node` - (Optional) Single Interface Node for Alternate Region. Parameters for creating Single interface Node for Alternate Region. See [Node](#ingress-gw-ar-node) below.

`performance_enhancement_mode` - (Optional) Performance Enhancement Mode. x-required Optimize the site for L3 or L7 traffic processing. L7 optimized is the default. See [Performance Enhancement Mode](#ingress-gw-ar-performance-enhancement-mode) below.

<a id="ingress-gw-ar-accelerated-networking"></a>

### Ingress Gw Ar Accelerated Networking

`disable` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`enable` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

<a id="ingress-gw-ar-node"></a>

### Ingress Gw Ar Node

`fault_domain` - (Optional) Number of fault domains. Namuber of fault domains to be used while creating the availability set (`Number`).

`local_subnet` - (Optional) Azure Subnet. Parameters for Azure subnet. See [Local Subnet](#ingress-gw-ar-node-local-subnet) below.

`node_number` - (Optional) Number of main nodes. Number of main nodes to create, either 1 or 3 (`Number`).

`update_domain` - (Optional) Number of update domains. Namuber of update domains to be used while creating the availability set (`Number`).

<a id="ingress-gw-ar-node-local-subnet"></a>

### Ingress Gw Ar Node Local Subnet

`subnet` - (Optional) Azure Subnet. Parameters for Azure subnet (`Block`).

`subnet_param` - (Optional) New Cloud Subnet Parameters. Parameters for creating a new cloud subnet (`Block`).

<a id="ingress-gw-ar-performance-enhancement-mode"></a>

### Ingress Gw Ar Performance Enhancement Mode

`perf_mode_l3_enhanced` - (Optional) L3 Mode Enhanced Performance. x-required L3 enhanced performance mode options. See [Perf Mode L3 Enhanced](#ingress-gw-ar-performance-enhancement-mode-perf-mode-l3-enhanced) below.

`perf_mode_l7_enhanced` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

<a id="ingress-gw-ar-performance-enhancement-mode-perf-mode-l3-enhanced"></a>

### Ingress Gw Ar Performance Enhancement Mode Perf Mode L3 Enhanced

`jumbo` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`no_jumbo` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

<a id="kubernetes-upgrade-drain"></a>

### Kubernetes Upgrade Drain

`disable_upgrade_drain` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`enable_upgrade_drain` - (Optional) Enable Node by Node Upgrade. Specify batch upgrade settings for worker nodes within a site. See [Enable Upgrade Drain](#kubernetes-upgrade-drain-enable-upgrade-drain) below.

<a id="kubernetes-upgrade-drain-enable-upgrade-drain"></a>

### Kubernetes Upgrade Drain Enable Upgrade Drain

`disable_vega_upgrade_mode` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`drain_max_unavailable_node_count` - (Optional) Node Batch Size Count (`Number`).

`drain_node_timeout` - (Optional) Upgrade Wait Time. Seconds to wait before initiating upgrade on the next set of nodes. Setting it to 0 will wait indefinitely for all services on nodes to be upgraded gracefully before proceeding to the next set of nodes. (Warning: It may block upgrade if services on a node cannot be gracefully upgraded. It is recommended to use the default value) (`Number`).

`enable_vega_upgrade_mode` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

<a id="log-receiver"></a>

### Log Receiver

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

<a id="offline-survivability-mode"></a>

### Offline Survivability Mode

`enable_offline_survivability_mode` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`no_offline_survivability_mode` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

<a id="os"></a>

### OS

`default_os_version` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`operating_system_version` - (Optional) Operating System Version. Specify a OS version to be used e.g. 9.2024.6 (`String`).

<a id="sw"></a>

### Sw

`default_sw_version` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`volterra_software_version` - (Optional) F5XC Software Version. Specify a F5XC Software Version to be used e.g. crt-20210329-1002 (`String`).

<a id="timeouts"></a>

### Timeouts

`create` - (Optional) A string that can be [parsed as a duration](`HTTPS://pkg.go.dev/time#ParseDuration`) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours) (`String`).

`delete` - (Optional) A string that can be [parsed as a duration](`HTTPS://pkg.go.dev/time#ParseDuration`) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are saved into state before the destroy operation occurs (`String`).

`read` - (Optional) A string that can be [parsed as a duration](`HTTPS://pkg.go.dev/time#ParseDuration`) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh is enabled (`String`).

`update` - (Optional) A string that can be [parsed as a duration](`HTTPS://pkg.go.dev/time#ParseDuration`) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours) (`String`).

<a id="vnet"></a>

### Vnet

`existing_vnet` - (Optional) Azure Existing Vnet Type. Resource group and name of existing Azure Vnet. See [Existing Vnet](#vnet-existing-vnet) below.

`new_vnet` - (Optional) Azure Vnet Parameters. Parameters to create a new Azure Vnet. See [New Vnet](#vnet-new-vnet) below.

<a id="vnet-existing-vnet"></a>

### Vnet Existing Vnet

`f5_orchestrated_routing` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`manual_routing` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`resource_group` - (Optional) Existing Vnet Resource Group. Resource group of existing Vnet (`String`).

`vnet_name` - (Optional) Existing Vnet Name. Name of existing Vnet (`String`).

<a id="vnet-new-vnet"></a>

### Vnet New Vnet

`autogenerate` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`name` - (Optional) Choose Vnet Name. Specify the Vnet Name (`String`).

`primary_ipv4` - (Optional) IPv4 CIDR block. IPv4 CIDR block for this Vnet. It has to be private address space (`String`).

<a id="voltstack-cluster"></a>

### Voltstack Cluster

`accelerated_networking` - (Optional) Accelerated Networking Type. x-required Accelerated Networking to reduce Latency, When Mode is toggled, traffic disruption will be seen. See [Accelerated Networking](#voltstack-cluster-accelerated-networking) below.

`active_enhanced_firewall_policies` - (Optional) Active Enhanced Network Policies Type. List of Enhanced Firewall Policies These policies use session-based rules and provide all options available under firewall policies with an additional option for service insertion. See [Active Enhanced Firewall Policies](#voltstack-cluster-active-enhanced-firewall-policies) below.

`active_forward_proxy_policies` - (Optional) Active Forward Proxy Policies Type. Ordered List of Forward Proxy Policies active. See [Active Forward Proxy Policies](#voltstack-cluster-active-forward-proxy-policies) below.

`active_network_policies` - (Optional) Active Firewall Policies Type. List of firewall policy views. See [Active Network Policies](#voltstack-cluster-active-network-policies) below.

`az_nodes` - (Optional) App Stack Cluster (One Interface) Nodes in AZ. Only Single AZ or Three AZ(s) nodes are supported currently. See [Az Nodes](#voltstack-cluster-az-nodes) below.

`azure_certified_hw` - (Optional) Azure Certified Hardware. Name for Azure certified hardware (`String`).

`dc_cluster_group` - (Optional) Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name. See [Dc Cluster Group](#voltstack-cluster-dc-cluster-group) below.

`default_storage` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`forward_proxy_allow_all` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`global_network_list` - (Optional) Global Network Connection List. List of global network connections. See [Global Network List](#voltstack-cluster-global-network-list) below.

`k8s_cluster` - (Optional) Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name. See [K8s Cluster](#voltstack-cluster-k8s-cluster) below.

`no_dc_cluster_group` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`no_forward_proxy` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`no_global_network` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`no_k8s_cluster` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`no_network_policy` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`no_outside_static_routes` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`outside_static_routes` - (Optional) Static Route List Type. List of static routes. See [Outside Static Routes](#voltstack-cluster-outside-static-routes) below.

`sm_connection_public_ip` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`sm_connection_pvt_ip` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`storage_class_list` - (Optional) Custom Storage Class List. Add additional custom storage classes in kubernetes for this site. See [Storage Class List](#voltstack-cluster-storage-class-list) below.

<a id="voltstack-cluster-accelerated-networking"></a>

### Voltstack Cluster Accelerated Networking

`disable` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`enable` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

<a id="voltstack-cluster-active-enhanced-firewall-policies"></a>

### Voltstack Cluster Active Enhanced Firewall Policies

`enhanced_firewall_policies` - (Optional) Enhanced Firewall Policy. Ordered List of Enhanced Firewall Policies active. See [Enhanced Firewall Policies](#voltstack-cluster-active-enhanced-firewall-policies-enhanced-firewall-policies) below.

<a id="voltstack-cluster-active-enhanced-firewall-policies-enhanced-firewall-policies"></a>

### Voltstack Cluster Active Enhanced Firewall Policies Enhanced Firewall Policies

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

<a id="voltstack-cluster-active-forward-proxy-policies"></a>

### Voltstack Cluster Active Forward Proxy Policies

`forward_proxy_policies` - (Optional) Forward Proxy Policies. Ordered List of Forward Proxy Policies active. See [Forward Proxy Policies](#voltstack-cluster-active-forward-proxy-policies-forward-proxy-policies) below.

<a id="voltstack-cluster-active-forward-proxy-policies-forward-proxy-policies"></a>

### Voltstack Cluster Active Forward Proxy Policies Forward Proxy Policies

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

<a id="voltstack-cluster-active-network-policies"></a>

### Voltstack Cluster Active Network Policies

`network_policies` - (Optional) Firewall Policy. Ordered List of Firewall Policies active for this network firewall. See [Network Policies](#voltstack-cluster-active-network-policies-network-policies) below.

<a id="voltstack-cluster-active-network-policies-network-policies"></a>

### Voltstack Cluster Active Network Policies Network Policies

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

<a id="voltstack-cluster-az-nodes"></a>

### Voltstack Cluster Az Nodes

`azure_az` - (Optional) Azure Availability Zone. A zone depicting a grouping of datacenters within an Azure region. Expecting numeric input (`String`).

`local_subnet` - (Optional) Azure Subnet. Parameters for Azure subnet. See [Local Subnet](#voltstack-cluster-az-nodes-local-subnet) below.

<a id="voltstack-cluster-az-nodes-local-subnet"></a>

### Voltstack Cluster Az Nodes Local Subnet

`subnet` - (Optional) Azure Subnet. Parameters for Azure subnet (`Block`).

`subnet_param` - (Optional) New Cloud Subnet Parameters. Parameters for creating a new cloud subnet (`Block`).

<a id="voltstack-cluster-dc-cluster-group"></a>

### Voltstack Cluster Dc Cluster Group

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

<a id="voltstack-cluster-global-network-list"></a>

### Voltstack Cluster Global Network List

`global_network_connections` - (Optional) Global Network Connections. Global network connections. See [Global Network Connections](#voltstack-cluster-global-network-list-global-network-connections) below.

<a id="voltstack-cluster-global-network-list-global-network-connections"></a>

### Voltstack Cluster Global Network List Global Network Connections

`sli_to_global_dr` - (Optional) Global Network. Global network reference for direct connection (`Block`).

`slo_to_global_dr` - (Optional) Global Network. Global network reference for direct connection (`Block`).

<a id="voltstack-cluster-k8s-cluster"></a>

### Voltstack Cluster K8s Cluster

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

<a id="voltstack-cluster-outside-static-routes"></a>

### Voltstack Cluster Outside Static Routes

`static_route_list` - (Optional) List of Static Routes. List of Static routes. See [Static Route List](#voltstack-cluster-outside-static-routes-static-route-list) below.

<a id="voltstack-cluster-outside-static-routes-static-route-list"></a>

### Voltstack Cluster Outside Static Routes Static Route List

`custom_static_route` - (Optional) Static Route. Defines a static route, configuring a list of prefixes and a next-hop to be used for them (`Block`).

`simple_static_route` - (Optional) Simple Static Route. Use simple static route for prefix pointing to single interface in the network (`String`).

<a id="voltstack-cluster-storage-class-list"></a>

### Voltstack Cluster Storage Class List

`storage_classes` - (Optional) List of Storage Classes. List of custom storage classes. See [Storage Classes](#voltstack-cluster-storage-class-list-storage-classes) below.

<a id="voltstack-cluster-storage-class-list-storage-classes"></a>

### Voltstack Cluster Storage Class List Storage Classes

`default_storage_class` - (Optional) Default Storage Class. Make this storage class default storage class for the K8s cluster (`Bool`).

`storage_class_name` - (Optional) Storage Class Name. Name of the storage class as it will appear in K8s (`String`).

<a id="voltstack-cluster-ar"></a>

### Voltstack Cluster Ar

`accelerated_networking` - (Optional) Accelerated Networking Type. x-required Accelerated Networking to reduce Latency, When Mode is toggled, traffic disruption will be seen. See [Accelerated Networking](#voltstack-cluster-ar-accelerated-networking) below.

`active_enhanced_firewall_policies` - (Optional) Active Enhanced Network Policies Type. List of Enhanced Firewall Policies These policies use session-based rules and provide all options available under firewall policies with an additional option for service insertion. See [Active Enhanced Firewall Policies](#voltstack-cluster-ar-active-enhanced-firewall-policies) below.

`active_forward_proxy_policies` - (Optional) Active Forward Proxy Policies Type. Ordered List of Forward Proxy Policies active. See [Active Forward Proxy Policies](#voltstack-cluster-ar-active-forward-proxy-policies) below.

`active_network_policies` - (Optional) Active Firewall Policies Type. List of firewall policy views. See [Active Network Policies](#voltstack-cluster-ar-active-network-policies) below.

`azure_certified_hw` - (Optional) Azure Certified Hardware. Name for Azure certified hardware (`String`).

`dc_cluster_group` - (Optional) Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name. See [Dc Cluster Group](#voltstack-cluster-ar-dc-cluster-group) below.

`default_storage` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`forward_proxy_allow_all` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`global_network_list` - (Optional) Global Network Connection List. List of global network connections. See [Global Network List](#voltstack-cluster-ar-global-network-list) below.

`k8s_cluster` - (Optional) Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name. See [K8s Cluster](#voltstack-cluster-ar-k8s-cluster) below.

`no_dc_cluster_group` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`no_forward_proxy` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`no_global_network` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`no_k8s_cluster` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`no_network_policy` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`no_outside_static_routes` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`node` - (Optional) Single Interface Node for Alternate Region. Parameters for creating Single interface Node for Alternate Region. See [Node](#voltstack-cluster-ar-node) below.

`outside_static_routes` - (Optional) Static Route List Type. List of static routes. See [Outside Static Routes](#voltstack-cluster-ar-outside-static-routes) below.

`sm_connection_public_ip` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`sm_connection_pvt_ip` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`storage_class_list` - (Optional) Custom Storage Class List. Add additional custom storage classes in kubernetes for this site. See [Storage Class List](#voltstack-cluster-ar-storage-class-list) below.

<a id="voltstack-cluster-ar-accelerated-networking"></a>

### Voltstack Cluster Ar Accelerated Networking

`disable` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`enable` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

<a id="voltstack-cluster-ar-active-enhanced-firewall-policies"></a>

### Voltstack Cluster Ar Active Enhanced Firewall Policies

`enhanced_firewall_policies` - (Optional) Enhanced Firewall Policy. Ordered List of Enhanced Firewall Policies active. See [Enhanced Firewall Policies](#voltstack-cluster-ar-active-enhanced-firewall-policies-enhanced-firewall-policies) below.

<a id="voltstack-cluster-ar-active-enhanced-firewall-policies-enhanced-firewall-policies"></a>

### Voltstack Cluster Ar Active Enhanced Firewall Policies Enhanced Firewall Policies

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

<a id="voltstack-cluster-ar-active-forward-proxy-policies"></a>

### Voltstack Cluster Ar Active Forward Proxy Policies

`forward_proxy_policies` - (Optional) Forward Proxy Policies. Ordered List of Forward Proxy Policies active. See [Forward Proxy Policies](#voltstack-cluster-ar-active-forward-proxy-policies-forward-proxy-policies) below.

<a id="voltstack-cluster-ar-active-forward-proxy-policies-forward-proxy-policies"></a>

### Voltstack Cluster Ar Active Forward Proxy Policies Forward Proxy Policies

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

<a id="voltstack-cluster-ar-active-network-policies"></a>

### Voltstack Cluster Ar Active Network Policies

`network_policies` - (Optional) Firewall Policy. Ordered List of Firewall Policies active for this network firewall. See [Network Policies](#voltstack-cluster-ar-active-network-policies-network-policies) below.

<a id="voltstack-cluster-ar-active-network-policies-network-policies"></a>

### Voltstack Cluster Ar Active Network Policies Network Policies

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

<a id="voltstack-cluster-ar-dc-cluster-group"></a>

### Voltstack Cluster Ar Dc Cluster Group

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

<a id="voltstack-cluster-ar-global-network-list"></a>

### Voltstack Cluster Ar Global Network List

`global_network_connections` - (Optional) Global Network Connections. Global network connections. See [Global Network Connections](#voltstack-cluster-ar-global-network-list-global-network-connections) below.

<a id="voltstack-cluster-ar-global-network-list-global-network-connections"></a>

### Voltstack Cluster Ar Global Network List Global Network Connections

`sli_to_global_dr` - (Optional) Global Network. Global network reference for direct connection (`Block`).

`slo_to_global_dr` - (Optional) Global Network. Global network reference for direct connection (`Block`).

<a id="voltstack-cluster-ar-k8s-cluster"></a>

### Voltstack Cluster Ar K8s Cluster

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

<a id="voltstack-cluster-ar-node"></a>

### Voltstack Cluster Ar Node

`fault_domain` - (Optional) Number of fault domains. Namuber of fault domains to be used while creating the availability set (`Number`).

`local_subnet` - (Optional) Azure Subnet. Parameters for Azure subnet. See [Local Subnet](#voltstack-cluster-ar-node-local-subnet) below.

`node_number` - (Optional) Number of main nodes. Number of main nodes to create, either 1 or 3 (`Number`).

`update_domain` - (Optional) Number of update domains. Namuber of update domains to be used while creating the availability set (`Number`).

<a id="voltstack-cluster-ar-node-local-subnet"></a>

### Voltstack Cluster Ar Node Local Subnet

`subnet` - (Optional) Azure Subnet. Parameters for Azure subnet (`Block`).

`subnet_param` - (Optional) New Cloud Subnet Parameters. Parameters for creating a new cloud subnet (`Block`).

<a id="voltstack-cluster-ar-outside-static-routes"></a>

### Voltstack Cluster Ar Outside Static Routes

`static_route_list` - (Optional) List of Static Routes. List of Static routes. See [Static Route List](#voltstack-cluster-ar-outside-static-routes-static-route-list) below.

<a id="voltstack-cluster-ar-outside-static-routes-static-route-list"></a>

### Voltstack Cluster Ar Outside Static Routes Static Route List

`custom_static_route` - (Optional) Static Route. Defines a static route, configuring a list of prefixes and a next-hop to be used for them (`Block`).

`simple_static_route` - (Optional) Simple Static Route. Use simple static route for prefix pointing to single interface in the network (`String`).

<a id="voltstack-cluster-ar-storage-class-list"></a>

### Voltstack Cluster Ar Storage Class List

`storage_classes` - (Optional) List of Storage Classes. List of custom storage classes. See [Storage Classes](#voltstack-cluster-ar-storage-class-list-storage-classes) below.

<a id="voltstack-cluster-ar-storage-class-list-storage-classes"></a>

### Voltstack Cluster Ar Storage Class List Storage Classes

`default_storage_class` - (Optional) Default Storage Class. Make this storage class default storage class for the K8s cluster (`Bool`).

`storage_class_name` - (Optional) Storage Class Name. Name of the storage class as it will appear in K8s (`String`).
