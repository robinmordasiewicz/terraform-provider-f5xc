---
page_title: "Voltstack Site Nested Blocks - f5xc Provider"
subcategory: "Sites"
description: |-
  Nested block reference for the Voltstack Site resource.
---

# Voltstack Site Nested Blocks

This page contains detailed documentation for nested blocks in the `f5xc_voltstack_site` resource.

For the main resource documentation, see [f5xc_voltstack_site](./resources/voltstack_site).

## Contents

- [blocked-services](#blocked-services)
- [blocked-services-blocked-sevice](#blocked-services-blocked-sevice)
- [bond-device-list](#bond-device-list)
- [bond-device-list-bond-devices](#bond-device-list-bond-devices)
- [bond-device-list-bond-devices-lacp](#bond-device-list-bond-devices-lacp)
- [coordinates](#coordinates)
- [custom-dns](#custom-dns)
- [custom-network-config](#custom-network-config)
- [custom-network-config-active-enhanced-firewall-policies](#custom-network-config-active-enhanced-firewall-policies)
- [custom-network-config-active-enhanced-firewall-policies-enhanced-firewall-policies](#custom-network-config-active-enhanced-firewall-policies-enhanced-firewall-policies)
- [custom-network-config-active-forward-proxy-policies](#custom-network-config-active-forward-proxy-policies)
- [custom-network-config-active-forward-proxy-policies-forward-proxy-policies](#custom-network-config-active-forward-proxy-policies-forward-proxy-policies)
- [custom-network-config-active-network-policies](#custom-network-config-active-network-policies)
- [custom-network-config-active-network-policies-network-policies](#custom-network-config-active-network-policies-network-policies)
- [custom-network-config-global-network-list](#custom-network-config-global-network-list)
- [custom-network-config-global-network-list-global-network-connections](#custom-network-config-global-network-list-global-network-connections)
- [custom-network-config-interface-list](#custom-network-config-interface-list)
- [custom-network-config-interface-list-interfaces](#custom-network-config-interface-list-interfaces)
- [custom-network-config-sli-config](#custom-network-config-sli-config)
- [custom-network-config-sli-config-static-routes](#custom-network-config-sli-config-static-routes)
- [custom-network-config-sli-config-static-v6-routes](#custom-network-config-sli-config-static-v6-routes)
- [custom-network-config-slo-config](#custom-network-config-slo-config)
- [custom-network-config-slo-config-dc-cluster-group](#custom-network-config-slo-config-dc-cluster-group)
- [custom-network-config-slo-config-static-routes](#custom-network-config-slo-config-static-routes)
- [custom-network-config-slo-config-static-v6-routes](#custom-network-config-slo-config-static-v6-routes)
- [custom-storage-config](#custom-storage-config)
- [custom-storage-config-static-routes](#custom-storage-config-static-routes)
- [custom-storage-config-static-routes-static-routes](#custom-storage-config-static-routes-static-routes)
- [custom-storage-config-storage-class-list](#custom-storage-config-storage-class-list)
- [custom-storage-config-storage-class-list-storage-classes](#custom-storage-config-storage-class-list-storage-classes)
- [custom-storage-config-storage-device-list](#custom-storage-config-storage-device-list)
- [custom-storage-config-storage-device-list-storage-devices](#custom-storage-config-storage-device-list-storage-devices)
- [custom-storage-config-storage-interface-list](#custom-storage-config-storage-interface-list)
- [custom-storage-config-storage-interface-list-storage-interfaces](#custom-storage-config-storage-interface-list-storage-interfaces)
- [enable-vgpu](#enable-vgpu)
- [k8s-cluster](#k8s-cluster)
- [kubernetes-upgrade-drain](#kubernetes-upgrade-drain)
- [kubernetes-upgrade-drain-enable-upgrade-drain](#kubernetes-upgrade-drain-enable-upgrade-drain)
- [local-control-plane](#local-control-plane)
- [local-control-plane-bgp-config](#local-control-plane-bgp-config)
- [local-control-plane-bgp-config-peers](#local-control-plane-bgp-config-peers)
- [log-receiver](#log-receiver)
- [master-node-configuration](#master-node-configuration)
- [offline-survivability-mode](#offline-survivability-mode)
- [os](#os)
- [sriov-interfaces](#sriov-interfaces)
- [sriov-interfaces-sriov-interface](#sriov-interfaces-sriov-interface)
- [sw](#sw)
- [timeouts](#timeouts)
- [usb-policy](#usb-policy)

---

<a id="blocked-services"></a>

### Blocked Services

`blocked_sevice` - (Optional) Disable Node Local Services. See [Blocked Sevice](#blocked-services-blocked-sevice) below.

<a id="blocked-services-blocked-sevice"></a>

### Blocked Services Blocked Sevice

`dns` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`network_type` - (Optional) Virtual Network Type. Different types of virtual networks understood by the system Virtual-network of type VIRTUAL_NETWORK_SITE_LOCAL provides connectivity to public (outside) network. This is an insecure network and is connected to public internet via NAT Gateways/firwalls Virtual-network of this type is local to every site. Two virtual networks of this type on different sites are neither related nor connected. Constraints: There can be atmost one virtual network of this type in a given site. This network type is supported on CE sites. This network is created automatically and present on all sites Virtual-network of type VIRTUAL_NETWORK_SITE_LOCAL_INSIDE is a private network inside site. It is a secure network and is not connected to public network. Virtual-network of this type is local to every site. Two virtual networks of this type on different sites are neither related nor connected. Constraints: There can be atmost one virtual network of this type in a given site. This network type is supported on CE sites. This network is created during provisioning of site User defined per-site virtual network. Scope of this virtual network is limited to the site. This is not yet supported Virtual-network of type VIRTUAL_NETWORK_PUBLIC directly conects to the public internet. Virtual-network of this type is local to every site. Two virtual networks of this type on different sites are neither related nor connected. Constraints: There can be atmost one virtual network of this type in a given site. This network type is supported on RE sites only It is an internally created by the system. They must not be created by user Virtual Neworks with global scope across different sites in F5XC domain. An example global virtual-network called 'AIN Network' is created for every tenant. for volterra fabric Constraints: It is currently only supported as internally created by the system. vK8s service network for a given tenant. Used to advertise a virtual host only to vk8s pods for that tenant Constraints: It is an internally created by the system. Must not be created by user VER internal network for the site. It can only be used for virtual hosts with SMA_PROXY type proxy Constraints: It is an internally created by the system. Must not be created by user Virtual-network of type VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE represents both VIRTUAL_NETWORK_SITE_LOCAL and VIRTUAL_NETWORK_SITE_LOCAL_INSIDE Constraints: This network type is only meaningful in an advertise policy When virtual-network of type VIRTUAL_NETWORK_IP_AUTO is selected for an endpoint, VER will try to determine the network based on the provided IP address Constraints: This network type is only meaningful in an endpoint VoltADN Private Network is used on volterra RE(s) to connect to customer private networks This network is created by opening a support ticket This network is per site srv6 network VER IP Fabric network for the site. This Virtual network type is used for exposing virtual host on IP Fabric network on the VER site or for endpoint in IP Fabric network Constraints: It is an internally created by the system. Must not be created by user Network internally created for a segment Constraints: It is an internally created by the system. Must not be created by user. Possible values are `VIRTUAL_NETWORK_SITE_LOCAL`, `VIRTUAL_NETWORK_SITE_LOCAL_INSIDE`, `VIRTUAL_NETWORK_PER_SITE`, `VIRTUAL_NETWORK_PUBLIC`, `VIRTUAL_NETWORK_GLOBAL`, `VIRTUAL_NETWORK_SITE_SERVICE`, `VIRTUAL_NETWORK_VER_INTERNAL`, `VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE`, `VIRTUAL_NETWORK_IP_AUTO`, `VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK`, `VIRTUAL_NETWORK_SRV6_NETWORK`, `VIRTUAL_NETWORK_IP_FABRIC`, `VIRTUAL_NETWORK_SEGMENT`. Defaults to `VIRTUAL_NETWORK_SITE_LOCAL` (`String`).

`ssh` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`web_user_interface` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

<a id="bond-device-list"></a>

### Bond Device List

`bond_devices` - (Optional) Bond Devices. List of bond devices. See [Bond Devices](#bond-device-list-bond-devices) below.

<a id="bond-device-list-bond-devices"></a>

### Bond Device List Bond Devices

`active_backup` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`devices` - (Optional) Member Ethernet Devices. Ethernet devices that will make up this bond (`List`).

`lacp` - (Optional) LACP parameters. LACP parameters for the bond device. See [Lacp](#bond-device-list-bond-devices-lacp) below.

`link_polling_interval` - (Optional) Link Polling Interval. Link polling interval in milliseconds (`Number`).

`link_up_delay` - (Optional) Link Up Delay. Milliseconds wait before link is declared up (`Number`).

`name` - (Optional) Bond Device Name. Name for the Bond. Ex 'bond0' (`String`).

<a id="bond-device-list-bond-devices-lacp"></a>

### Bond Device List Bond Devices Lacp

`rate` - (Optional) LACP Packet Interval. Interval in seconds to transmit LACP packets (`Number`).

<a id="coordinates"></a>

### Coordinates

`latitude` - (Optional) Latitude. Latitude of the site location (`Number`).

`longitude` - (Optional) Longitude. longitude of site location (`Number`).

<a id="custom-dns"></a>

### Custom DNS

`inside_nameserver` - (Optional) DNS Server for Inside Network. Optional DNS server IP to be used for name resolution in inside network (`String`).

`outside_nameserver` - (Optional) DNS Server for Outside Network. Optional DNS server IP to be used for name resolution in outside network (`String`).

<a id="custom-network-config"></a>

### Custom Network Config

`active_enhanced_firewall_policies` - (Optional) Active Enhanced Network Policies Type. List of Enhanced Firewall Policies These policies use session-based rules and provide all options available under firewall policies with an additional option for service insertion. See [Active Enhanced Firewall Policies](#custom-network-config-active-enhanced-firewall-policies) below.

`active_forward_proxy_policies` - (Optional) Active Forward Proxy Policies Type. Ordered List of Forward Proxy Policies active. See [Active Forward Proxy Policies](#custom-network-config-active-forward-proxy-policies) below.

`active_network_policies` - (Optional) Active Firewall Policies Type. List of firewall policy views. See [Active Network Policies](#custom-network-config-active-network-policies) below.

`bgp_peer_address` - (Optional) BGP Peer Address. Optional BGP peer address that can be used as parameter for BGP configuration when BGP is configured to fetch BGP peer address from site Object. This can be used to change peer address per site in fleet (`String`).

`bgp_router_id` - (Optional) BGP Router ID. Optional BGP router id that can be used as parameter for BGP configuration when BGP is configured to fetch BGP router ID from site object (`String`).

`default_config` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`default_interface_config` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`default_sli_config` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`forward_proxy_allow_all` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`global_network_list` - (Optional) Global Network Connection List. List of global network connections. See [Global Network List](#custom-network-config-global-network-list) below.

`interface_list` - (Optional) List of Interface. Configure network interfaces for this App Stack site. See [Interface List](#custom-network-config-interface-list) below.

`no_forward_proxy` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`no_global_network` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`no_network_policy` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`outside_nameserver` - (Optional) DNS V4 Server for Local Network. Optional DNS server V4 IP to be used for name resolution in local network (`String`).

`outside_vip` - (Optional) Common V4 VIP. Optional common virtual V4 IP across all nodes to be used as automatic VIP for site local network (`String`).

`site_to_site_tunnel_ip` - (Optional) Site Mesh Group Connection Via Virtual IP. Site Mesh Group Connection Via Virtual IP. This option will use the Virtual IP provided for creating ipsec between two sites which are part of the site mesh group (`String`).

`sli_config` - (Optional) Site Local Inside Network Configuration. Site local inside network configuration. See [Sli Config](#custom-network-config-sli-config) below.

`slo_config` - (Optional) Site Local Network Configuration. Site local network configuration. See [Slo Config](#custom-network-config-slo-config) below.

`sm_connection_public_ip` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`sm_connection_pvt_ip` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`tunnel_dead_timeout` - (Optional) Tunnel Dead Timeout (msec). Time interval, in millisec, within which any ipsec / SSL connection from the site going down is detected. When not set (== 0), a default value of 10000 msec will be used (`Number`).

`vip_vrrp_mode` - (Optional) VRRP Virtual-IP. VRRP advertisement mode for VIP Invalid VRRP mode. Possible values are `VIP_VRRP_INVALID`, `VIP_VRRP_ENABLE`, `VIP_VRRP_DISABLE` (`String`).

<a id="custom-network-config-active-enhanced-firewall-policies"></a>

### Custom Network Config Active Enhanced Firewall Policies

`enhanced_firewall_policies` - (Optional) Enhanced Firewall Policy. Ordered List of Enhanced Firewall Policies active. See [Enhanced Firewall Policies](#custom-network-config-active-enhanced-firewall-policies-enhanced-firewall-policies) below.

<a id="custom-network-config-active-enhanced-firewall-policies-enhanced-firewall-policies"></a>

### Custom Network Config Active Enhanced Firewall Policies Enhanced Firewall Policies

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

<a id="custom-network-config-active-forward-proxy-policies"></a>

### Custom Network Config Active Forward Proxy Policies

`forward_proxy_policies` - (Optional) Forward Proxy Policies. Ordered List of Forward Proxy Policies active. See [Forward Proxy Policies](#custom-network-config-active-forward-proxy-policies-forward-proxy-policies) below.

<a id="custom-network-config-active-forward-proxy-policies-forward-proxy-policies"></a>

### Custom Network Config Active Forward Proxy Policies Forward Proxy Policies

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

<a id="custom-network-config-active-network-policies"></a>

### Custom Network Config Active Network Policies

`network_policies` - (Optional) Firewall Policy. Ordered List of Firewall Policies active for this network firewall. See [Network Policies](#custom-network-config-active-network-policies-network-policies) below.

<a id="custom-network-config-active-network-policies-network-policies"></a>

### Custom Network Config Active Network Policies Network Policies

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

<a id="custom-network-config-global-network-list"></a>

### Custom Network Config Global Network List

`global_network_connections` - (Optional) Global Network Connections. Global network connections. See [Global Network Connections](#custom-network-config-global-network-list-global-network-connections) below.

<a id="custom-network-config-global-network-list-global-network-connections"></a>

### Custom Network Config Global Network List Global Network Connections

`sli_to_global_dr` - (Optional) Global Network. Global network reference for direct connection (`Block`).

`slo_to_global_dr` - (Optional) Global Network. Global network reference for direct connection (`Block`).

<a id="custom-network-config-interface-list"></a>

### Custom Network Config Interface List

`interfaces` - (Optional) List of Interface. Configure network interfaces for this App Stack site. See [Interfaces](#custom-network-config-interface-list-interfaces) below.

<a id="custom-network-config-interface-list-interfaces"></a>

### Custom Network Config Interface List Interfaces

`dc_cluster_group_connectivity_interface_disabled` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`dc_cluster_group_connectivity_interface_enabled` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`dedicated_interface` - (Optional) Dedicated Interface. Dedicated Interface Configuration (`Block`).

`dedicated_management_interface` - (Optional) Dedicated Management Interface. Dedicated Interface Configuration (`Block`).

`description` - (Optional) Interface Description. Description for this Interface (`String`).

`ethernet_interface` - (Optional) Ethernet Interface. Ethernet Interface Configuration (`Block`).

`labels` - (Optional) Interface Labels. Add Labels for this Interface, these labels can be used in firewall policy (`Block`).

`tunnel_interface` - (Optional) Tunnel Interface. Tunnel Interface Configuration (`Block`).

<a id="custom-network-config-sli-config"></a>

### Custom Network Config Sli Config

`no_static_routes` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`no_v6_static_routes` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`static_routes` - (Optional) Static Routes List. List of static routes. See [Static Routes](#custom-network-config-sli-config-static-routes) below.

`static_v6_routes` - (Optional) Static IPv6 Routes List. List of IPv6 static routes. See [Static V6 Routes](#custom-network-config-sli-config-static-v6-routes) below.

<a id="custom-network-config-sli-config-static-routes"></a>

### Custom Network Config Sli Config Static Routes

`static_routes` - (Optional) Static Routes. List of static routes (`Block`).

<a id="custom-network-config-sli-config-static-v6-routes"></a>

### Custom Network Config Sli Config Static V6 Routes

`static_routes` - (Optional) Static IPv6 Routes. List of IPv6 static routes (`Block`).

<a id="custom-network-config-slo-config"></a>

### Custom Network Config Slo Config

`dc_cluster_group` - (Optional) Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name. See [Dc Cluster Group](#custom-network-config-slo-config-dc-cluster-group) below.

`labels` - (Optional) Network Labels. Add Labels for this network, these labels can be used in firewall policy (`Block`).

`no_dc_cluster_group` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`no_static_routes` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`no_static_v6_routes` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`static_routes` - (Optional) Static Routes List. List of static routes. See [Static Routes](#custom-network-config-slo-config-static-routes) below.

`static_v6_routes` - (Optional) Static IPv6 Routes List. List of IPv6 static routes. See [Static V6 Routes](#custom-network-config-slo-config-static-v6-routes) below.

<a id="custom-network-config-slo-config-dc-cluster-group"></a>

### Custom Network Config Slo Config Dc Cluster Group

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

<a id="custom-network-config-slo-config-static-routes"></a>

### Custom Network Config Slo Config Static Routes

`static_routes` - (Optional) Static Routes. List of static routes (`Block`).

<a id="custom-network-config-slo-config-static-v6-routes"></a>

### Custom Network Config Slo Config Static V6 Routes

`static_routes` - (Optional) Static IPv6 Routes. List of IPv6 static routes (`Block`).

<a id="custom-storage-config"></a>

### Custom Storage Config

`default_storage_class` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`no_static_routes` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`no_storage_device` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`no_storage_interfaces` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`static_routes` - (Optional) Static Routes List. List of static routes. See [Static Routes](#custom-storage-config-static-routes) below.

`storage_class_list` - (Optional) Custom Storage Class List. Add additional custom storage classes in kubernetes for this fleet. See [Storage Class List](#custom-storage-config-storage-class-list) below.

`storage_device_list` - (Optional) Custom Storage Device List. Add additional custom storage classes in kubernetes for this fleet. See [Storage Device List](#custom-storage-config-storage-device-list) below.

`storage_interface_list` - (Optional) List of Interface. Configure storage interfaces for this App Stack site. See [Storage Interface List](#custom-storage-config-storage-interface-list) below.

<a id="custom-storage-config-static-routes"></a>

### Custom Storage Config Static Routes

`static_routes` - (Optional) Static Routes. List of static routes. See [Static Routes](#custom-storage-config-static-routes-static-routes) below.

<a id="custom-storage-config-static-routes-static-routes"></a>

### Custom Storage Config Static Routes Static Routes

`attrs` - (Optional) Attributes. List of attributes that control forwarding, dynamic routing and control plane (host) reachability (`List`).

`default_gateway` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`ip_address` - (Optional) IP Address. Traffic matching the IP prefixes is sent to this IP Address (`String`).

`ip_prefixes` - (Optional) IP Prefixes. List of route prefixes that have common next hop and attributes (`List`).

`node_interface` - (Optional) NodeInterfaceType. On multinode site, this type holds the information about per node interfaces (`Block`).

<a id="custom-storage-config-storage-class-list"></a>

### Custom Storage Config Storage Class List

`storage_classes` - (Optional) List of Storage Classes. List of custom storage classes. See [Storage Classes](#custom-storage-config-storage-class-list-storage-classes) below.

<a id="custom-storage-config-storage-class-list-storage-classes"></a>

### Custom Storage Config Storage Class List Storage Classes

`advanced_storage_parameters` - (Optional) Advanced Parameters. Map of parameter name and string value (`Block`).

`allow_volume_expansion` - (Optional) Allow Volume Expansion. Allow volume expansion (`Bool`).

`custom_storage` - (Optional) Custom StorageClass. Custom Storage Class allows to insert Kubernetes storageclass definition which will be applied into given site (`Block`).

`default_storage_class` - (Optional) Default Storage Class. Make this storage class default storage class for the K8s cluster (`Bool`).

`description` - (Optional) Storage Class Description. Description for this storage class (`String`).

`hpe_storage` - (Optional) HPE Storage. Storage class Device configuration for HPE Storage (`Block`).

`netapp_trident` - (Optional) NetApp Trident Storage. Storage class Device configuration for NetApp Trident (`Block`).

`pure_service_orchestrator` - (Optional) Pure Storage Service Orchestrator. Storage class Device configuration for Pure Service Orchestrator (`Block`).

`reclaim_policy` - (Optional) Reclaim Policy. Reclaim Policy (`String`).

`storage_class_name` - (Optional) Storage Class Name. Name of the storage class as it will appear in K8s (`String`).

`storage_device` - (Optional) Storage Device. Storage device that this class will use. The Device name defined at previous step (`String`).

<a id="custom-storage-config-storage-device-list"></a>

### Custom Storage Config Storage Device List

`storage_devices` - (Optional) List of Storage Devices. List of custom storage devices. See [Storage Devices](#custom-storage-config-storage-device-list-storage-devices) below.

<a id="custom-storage-config-storage-device-list-storage-devices"></a>

### Custom Storage Config Storage Device List Storage Devices

`advanced_advanced_parameters` - (Optional) Advanced Parameters. Map of parameter name and string value (`Block`).

`custom_storage` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`hpe_storage` - (Optional) HPE Storage. Device configuration for HPE Storage (`Block`).

`netapp_trident` - (Optional) NetApp Trident. Device configuration for NetApp Trident Storage (`Block`).

`pure_service_orchestrator` - (Optional) Pure Storage Service Orchestrator. Device configuration for Pure Storage Service Orchestrator (`Block`).

`storage_device` - (Optional) Storage Device. Storage device and device unit (`String`).

<a id="custom-storage-config-storage-interface-list"></a>

### Custom Storage Config Storage Interface List

`storage_interfaces` - (Optional) List of Interface. Configure storage interfaces for this App Stack site. See [Storage Interfaces](#custom-storage-config-storage-interface-list-storage-interfaces) below.

<a id="custom-storage-config-storage-interface-list-storage-interfaces"></a>

### Custom Storage Config Storage Interface List Storage Interfaces

`description` - (Optional) Interface Description. Description for this Interface (`String`).

`labels` - (Optional) Interface Labels. Add Labels for this Interface, these labels can be used in firewall policy (`Block`).

`storage_interface` - (Optional) Ethernet Interface. Ethernet Interface Configuration (`Block`).

<a id="enable-vgpu"></a>

### Enable Vgpu

`feature_type` - (Optional) Feature Type. Set feature to be enabled Operate with a degraded vGPU performance Enable NVIDIA vGPU Enable NVIDIA RTX Virtual Workstation Enable NVIDIA Virtual Compute Server. Possible values are `UNLICENSED`, `VGPU`, `VWS`, `VCS`. Defaults to `UNLICENSED` (`String`).

`server_address` - (Optional) License Server Address. Set License Server Address (`String`).

`server_port` - (Optional) License Server Port Number. Set License Server port number (`Number`).

<a id="k8s-cluster"></a>

### K8s Cluster

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

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

<a id="local-control-plane"></a>

### Local Control Plane

`bgp_config` - (Optional) BGP Configuration. BGP configuration parameters. See [BGP Config](#local-control-plane-bgp-config) below.

`inside_vn` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`outside_vn` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

<a id="local-control-plane-bgp-config"></a>

### Local Control Plane BGP Config

`asn` - (Optional) ASN. Autonomous System Number (`Number`).

`peers` - (Optional) Peers. BGP parameters for peer. See [Peers](#local-control-plane-bgp-config-peers) below.

<a id="local-control-plane-bgp-config-peers"></a>

### Local Control Plane BGP Config Peers

`bfd_disabled` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`bfd_enabled` - (Optional) BFD. BFD parameters (`Block`).

`disable` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`external` - (Optional) External BGP Peer. External BGP Peer parameters (`Block`).

`label` - (Optional) Label. Specify whether this peer should be (`String`).

`metadata` - (Optional) Message Metadata. MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs (`Block`).

`passive_mode_disabled` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`passive_mode_enabled` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`routing_policies` - (Optional) BGP Routing Policy. List of rules which can be applied on all or particular nodes (`Block`).

<a id="log-receiver"></a>

### Log Receiver

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

<a id="master-node-configuration"></a>

### Master Node Configuration

`name` - (Optional) Name. Names of master node (`String`).

`public_ip` - (Optional) Public IP. IP Address of the master node. This IP will be used when other sites connect via Site Mesh Group (`String`).

<a id="offline-survivability-mode"></a>

### Offline Survivability Mode

`enable_offline_survivability_mode` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`no_offline_survivability_mode` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

<a id="os"></a>

### OS

`default_os_version` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`operating_system_version` - (Optional) Operating System Version. Specify a OS version to be used e.g. 9.2024.6 (`String`).

<a id="sriov-interfaces"></a>

### Sriov Interfaces

`sriov_interface` - (Optional) Custom SR-IOV interfaces Configuration. Use custom SR-IOV interfaces Configuration. See [Sriov Interface](#sriov-interfaces-sriov-interface) below.

<a id="sriov-interfaces-sriov-interface"></a>

### Sriov Interfaces Sriov Interface

`interface_name` - (Optional) Name of physical interface. Name of SR-IOV physical interface (`String`).

`number_of_vfio_vfs` - (Optional) Number of virtual functions reserved for vfio. Number of virtual functions reserved for VNFs and DPDK-based CNFs (`Number`).

`number_of_vfs` - (Optional) Total number of virtual functions. Total number of virtual functions (`Number`).

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

<a id="usb-policy"></a>

### Usb Policy

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).
