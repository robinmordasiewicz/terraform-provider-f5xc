---
page_title: "AWS Tgw Site Nested Blocks - f5xc Provider"
subcategory: "Sites"
description: |-
  Nested block reference for the AWS Tgw Site resource.
---

# AWS Tgw Site Nested Blocks

This page contains detailed documentation for nested blocks in the `f5xc_aws_tgw_site` resource.

For the main resource documentation, see [f5xc_aws_tgw_site](./resources/aws_tgw_site).

## Contents

- [aws-parameters](#aws-parameters)
- [aws-parameters-admin-password](#aws-parameters-admin-password)
- [aws-parameters-admin-password-blindfold-secret-info](#aws-parameters-admin-password-blindfold-secret-info)
- [aws-parameters-admin-password-clear-secret-info](#aws-parameters-admin-password-clear-secret-info)
- [aws-parameters-aws-cred](#aws-parameters-aws-cred)
- [aws-parameters-az-nodes](#aws-parameters-az-nodes)
- [aws-parameters-az-nodes-inside-subnet](#aws-parameters-az-nodes-inside-subnet)
- [aws-parameters-az-nodes-outside-subnet](#aws-parameters-az-nodes-outside-subnet)
- [aws-parameters-az-nodes-workload-subnet](#aws-parameters-az-nodes-workload-subnet)
- [aws-parameters-custom-security-group](#aws-parameters-custom-security-group)
- [aws-parameters-existing-tgw](#aws-parameters-existing-tgw)
- [aws-parameters-new-tgw](#aws-parameters-new-tgw)
- [aws-parameters-new-tgw-user-assigned](#aws-parameters-new-tgw-user-assigned)
- [aws-parameters-new-vpc](#aws-parameters-new-vpc)
- [aws-parameters-tgw-cidr](#aws-parameters-tgw-cidr)
- [blocked-services](#blocked-services)
- [blocked-services-blocked-sevice](#blocked-services-blocked-sevice)
- [coordinates](#coordinates)
- [custom-dns](#custom-dns)
- [direct-connect-enabled](#direct-connect-enabled)
- [direct-connect-enabled-hosted-vifs](#direct-connect-enabled-hosted-vifs)
- [direct-connect-enabled-hosted-vifs-site-registration-over-direct-connect](#direct-connect-enabled-hosted-vifs-site-registration-over-direct-connect)
- [direct-connect-enabled-hosted-vifs-vif-list](#direct-connect-enabled-hosted-vifs-vif-list)
- [kubernetes-upgrade-drain](#kubernetes-upgrade-drain)
- [kubernetes-upgrade-drain-enable-upgrade-drain](#kubernetes-upgrade-drain-enable-upgrade-drain)
- [log-receiver](#log-receiver)
- [offline-survivability-mode](#offline-survivability-mode)
- [os](#os)
- [performance-enhancement-mode](#performance-enhancement-mode)
- [performance-enhancement-mode-perf-mode-l3-enhanced](#performance-enhancement-mode-perf-mode-l3-enhanced)
- [private-connectivity](#private-connectivity)
- [private-connectivity-cloud-link](#private-connectivity-cloud-link)
- [sw](#sw)
- [tgw-security](#tgw-security)
- [tgw-security-active-east-west-service-policies](#tgw-security-active-east-west-service-policies)
- [tgw-security-active-east-west-service-policies-service-policies](#tgw-security-active-east-west-service-policies-service-policies)
- [tgw-security-active-enhanced-firewall-policies](#tgw-security-active-enhanced-firewall-policies)
- [tgw-security-active-enhanced-firewall-policies-enhanced-firewall-policies](#tgw-security-active-enhanced-firewall-policies-enhanced-firewall-policies)
- [tgw-security-active-forward-proxy-policies](#tgw-security-active-forward-proxy-policies)
- [tgw-security-active-forward-proxy-policies-forward-proxy-policies](#tgw-security-active-forward-proxy-policies-forward-proxy-policies)
- [tgw-security-active-network-policies](#tgw-security-active-network-policies)
- [tgw-security-active-network-policies-network-policies](#tgw-security-active-network-policies-network-policies)
- [timeouts](#timeouts)
- [vn-config](#vn-config)
- [vn-config-allowed-vip-port](#vn-config-allowed-vip-port)
- [vn-config-allowed-vip-port-custom-ports](#vn-config-allowed-vip-port-custom-ports)
- [vn-config-allowed-vip-port-sli](#vn-config-allowed-vip-port-sli)
- [vn-config-allowed-vip-port-sli-custom-ports](#vn-config-allowed-vip-port-sli-custom-ports)
- [vn-config-dc-cluster-group-inside-vn](#vn-config-dc-cluster-group-inside-vn)
- [vn-config-dc-cluster-group-outside-vn](#vn-config-dc-cluster-group-outside-vn)
- [vn-config-global-network-list](#vn-config-global-network-list)
- [vn-config-global-network-list-global-network-connections](#vn-config-global-network-list-global-network-connections)
- [vn-config-inside-static-routes](#vn-config-inside-static-routes)
- [vn-config-inside-static-routes-static-route-list](#vn-config-inside-static-routes-static-route-list)
- [vn-config-outside-static-routes](#vn-config-outside-static-routes)
- [vn-config-outside-static-routes-static-route-list](#vn-config-outside-static-routes-static-route-list)
- [vpc-attachments](#vpc-attachments)
- [vpc-attachments-vpc-list](#vpc-attachments-vpc-list)

---

<a id="aws-parameters"></a>

### AWS Parameters

`admin_password` - (Optional) Secret. SecretType is used in an object to indicate a sensitive/confidential field. See [Admin Password](#aws-parameters-admin-password) below.

`aws_cred` - (Optional) Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name. See [AWS Cred](#aws-parameters-aws-cred) below.

`aws_region` - (Optional) AWS Region. AWS Region of your services vpc, where F5XC site will be deployed (`String`).

`az_nodes` - (Optional) Ingress/Egress Gateway (two Interface) Nodes in AZ. Only Single AZ or Three AZ(s) nodes are supported currently. See [Az Nodes](#aws-parameters-az-nodes) below.

`custom_security_group` - (Optional) Security Group IDS. Enter pre created security groups for slo(Site Local Outside) and sli(Site Local Inside) interface. Supported only for sites deployed on existing VPC. See [Custom Security Group](#aws-parameters-custom-security-group) below.

`disable_internet_vip` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`disk_size` - (Optional) Node Disk Size. Node disk size for all node in the F5XC site. Unit is GiB (`Number`).

`enable_internet_vip` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`existing_tgw` - (Optional) Existing TGW Type. Information needed for existing TGW. See [Existing Tgw](#aws-parameters-existing-tgw) below.

`f5xc_security_group` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`instance_type` - (Optional) AWS Instance Type for Node. Instance size based on the performance (`String`).

`new_tgw` - (Optional) TGWParamsType. See [New Tgw](#aws-parameters-new-tgw) below.

`new_vpc` - (Optional) AWS VPC Parameters. Parameters to create new AWS VPC. See [New Vpc](#aws-parameters-new-vpc) below.

`no_worker_nodes` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`nodes_per_az` - (Optional) Desired Worker Nodes Per AZ. Desired Worker Nodes Per AZ. Max limit is up to 21 (`Number`).

`reserved_tgw_cidr` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`ssh_key` - (Optional) Public SSH key. Public SSH key for accessing nodes of the site (`String`).

`tgw_cidr` - (Optional) New Cloud Subnet Parameters. Parameters for creating a new cloud subnet. See [Tgw CIDR](#aws-parameters-tgw-cidr) below.

`total_nodes` - (Optional) Total Number of Worker Nodes for a Site. Total number of worker nodes to be deployed across all AZ's used in the Site (`Number`).

`vpc_id` - (Optional) Existing VPC ID. Existing VPC ID (`String`).

<a id="aws-parameters-admin-password"></a>

### AWS Parameters Admin Password

`blindfold_secret_info` - (Optional) Blindfold Secret. BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management. See [Blindfold Secret Info](#aws-parameters-admin-password-blindfold-secret-info) below.

`clear_secret_info` - (Optional) In-Clear Secret. ClearSecretInfoType specifies information about the Secret that is not encrypted. See [Clear Secret Info](#aws-parameters-admin-password-clear-secret-info) below.

<a id="aws-parameters-admin-password-blindfold-secret-info"></a>

### AWS Parameters Admin Password Blindfold Secret Info

`decryption_provider` - (Optional) Decryption Provider. Name of the Secret Management Access object that contains information about the backend Secret Management service (`String`).

`location` - (Optional) Location. Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location (`String`).

`store_provider` - (Optional) Store Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:/// (`String`).

<a id="aws-parameters-admin-password-clear-secret-info"></a>

### AWS Parameters Admin Password Clear Secret Info

`provider_ref` - (Optional) Provider. Name of the Secret Management Access object that contains information about the store to get encrypted bytes This field needs to be provided only if the URL scheme is not string:/// (`String`).

`url` - (Optional) URL. URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will get Secret bytes after Base64 decoding (`String`).

<a id="aws-parameters-aws-cred"></a>

### AWS Parameters AWS Cred

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

<a id="aws-parameters-az-nodes"></a>

### AWS Parameters Az Nodes

`aws_az_name` - (Optional) AWS AZ Name. AWS availability zone, must be consistent with the selected AWS region (`String`).

`inside_subnet` - (Optional) AWS Subnet. Parameters for AWS subnet. See [Inside Subnet](#aws-parameters-az-nodes-inside-subnet) below.

`outside_subnet` - (Optional) AWS Subnet. Parameters for AWS subnet. See [Outside Subnet](#aws-parameters-az-nodes-outside-subnet) below.

`reserved_inside_subnet` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`workload_subnet` - (Optional) AWS Subnet. Parameters for AWS subnet. See [Workload Subnet](#aws-parameters-az-nodes-workload-subnet) below.

<a id="aws-parameters-az-nodes-inside-subnet"></a>

### AWS Parameters Az Nodes Inside Subnet

`existing_subnet_id` - (Optional) Existing Subnet ID. Information about existing subnet ID (`String`).

`subnet_param` - (Optional) New Cloud Subnet Parameters. Parameters for creating a new cloud subnet (`Block`).

<a id="aws-parameters-az-nodes-outside-subnet"></a>

### AWS Parameters Az Nodes Outside Subnet

`existing_subnet_id` - (Optional) Existing Subnet ID. Information about existing subnet ID (`String`).

`subnet_param` - (Optional) New Cloud Subnet Parameters. Parameters for creating a new cloud subnet (`Block`).

<a id="aws-parameters-az-nodes-workload-subnet"></a>

### AWS Parameters Az Nodes Workload Subnet

`existing_subnet_id` - (Optional) Existing Subnet ID. Information about existing subnet ID (`String`).

`subnet_param` - (Optional) New Cloud Subnet Parameters. Parameters for creating a new cloud subnet (`Block`).

<a id="aws-parameters-custom-security-group"></a>

### AWS Parameters Custom Security Group

`inside_security_group_id` - (Optional) Inside Security Group ID. Security Group ID to be attached to SLI(Site Local Inside) Interface (`String`).

`outside_security_group_id` - (Optional) Outside Security Group ID. Security Group ID to be attached to SLO(Site Local Outside) Interface (`String`).

<a id="aws-parameters-existing-tgw"></a>

### AWS Parameters Existing Tgw

`tgw_asn` - (Optional) Enter TGW ASN. TGW ASN (`Number`).

`tgw_id` - (Optional) Existing TGW ID. Existing TGW ID (`String`).

`volterra_site_asn` - (Optional) Enter F5XC Site ASN. F5XC Site ASN (`Number`).

<a id="aws-parameters-new-tgw"></a>

### AWS Parameters New Tgw

`system_generated` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`user_assigned` - (Optional) TGW Assigned ASN Type. Information needed when ASNs are assigned by the user. See [User Assigned](#aws-parameters-new-tgw-user-assigned) below.

<a id="aws-parameters-new-tgw-user-assigned"></a>

### AWS Parameters New Tgw User Assigned

`tgw_asn` - (Optional) Enter TGW ASN. TGW ASN. Allowed range for 16-bit private ASNs include 64512 to 65534 (`Number`).

`volterra_site_asn` - (Optional) Enter F5XC Site ASN. F5XC Site ASN (`Number`).

<a id="aws-parameters-new-vpc"></a>

### AWS Parameters New Vpc

`autogenerate` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`name_tag` - (Optional) Choose VPC Name. Specify the VPC Name (`String`).

`primary_ipv4` - (Optional) Primary IPv4 CIDR block. IPv4 CIDR block for this VPC. It has to be private address space. The Primary IPv4 block cannot be modified. All subnets prefixes in this VPC must be part of this CIDR block (`String`).

<a id="aws-parameters-tgw-cidr"></a>

### AWS Parameters Tgw CIDR

`ipv4` - (Optional) IPv4 Subnet. IPv4 subnet prefix for this subnet (`String`).

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

<a id="direct-connect-enabled"></a>

### Direct Connect Enabled

`auto_asn` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`custom_asn` - (Optional) Custom ASN. Custom Autonomous System Number (`Number`).

`hosted_vifs` - (Optional) AWS Direct Connect Hosted VIF Config. x-example: 'value' AWS Direct Connect Hosted VIF Configuration. See [Hosted Vifs](#direct-connect-enabled-hosted-vifs) below.

`standard_vifs` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

<a id="direct-connect-enabled-hosted-vifs"></a>

### Direct Connect Enabled Hosted Vifs

`site_registration_over_direct_connect` - (Optional) CloudLink ADN Network Config. See [Site Registration Over Direct Connect](#direct-connect-enabled-hosted-vifs-site-registration-over-direct-connect) below.

`site_registration_over_internet` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`vif_list` - (Optional) List of Hosted VIF Config. List of Hosted VIF Config. See [Vif List](#direct-connect-enabled-hosted-vifs-vif-list) below.

<a id="direct-connect-enabled-hosted-vifs-site-registration-over-direct-connect"></a>

### Direct Connect Enabled Hosted Vifs Site Registration Over Direct Connect

`cloudlink_network_name` - (Optional) Private ADN Network. Establish private connectivity with the F5 Distributed Cloud Global Network using a Private ADN network. To provision a Private ADN network, please contact F5 Distributed Cloud support (`String`).

<a id="direct-connect-enabled-hosted-vifs-vif-list"></a>

### Direct Connect Enabled Hosted Vifs Vif List

`other_region` - (Optional) Other Region. Other Region (`String`).

`same_as_site_region` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`vif_id` - (Optional) VIF ID. AWS Direct Connect VIF ID that needs to be connected to the site (`String`).

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

<a id="performance-enhancement-mode"></a>

### Performance Enhancement Mode

`perf_mode_l3_enhanced` - (Optional) L3 Mode Enhanced Performance. x-required L3 enhanced performance mode options. See [Perf Mode L3 Enhanced](#performance-enhancement-mode-perf-mode-l3-enhanced) below.

`perf_mode_l7_enhanced` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

<a id="performance-enhancement-mode-perf-mode-l3-enhanced"></a>

### Performance Enhancement Mode Perf Mode L3 Enhanced

`jumbo` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`no_jumbo` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

<a id="private-connectivity"></a>

### Private Connectivity

`cloud_link` - (Optional) Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name. See [Cloud Link](#private-connectivity-cloud-link) below.

`inside` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`outside` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

<a id="private-connectivity-cloud-link"></a>

### Private Connectivity Cloud Link

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

<a id="sw"></a>

### Sw

`default_sw_version` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`volterra_software_version` - (Optional) F5XC Software Version. Specify a F5XC Software Version to be used e.g. crt-20210329-1002 (`String`).

<a id="tgw-security"></a>

### Tgw Security

`active_east_west_service_policies` - (Optional) Active Service Policies. Active service policies for the east-west proxy. See [Active East West Service Policies](#tgw-security-active-east-west-service-policies) below.

`active_enhanced_firewall_policies` - (Optional) Active Enhanced Network Policies Type. List of Enhanced Firewall Policies These policies use session-based rules and provide all options available under firewall policies with an additional option for service insertion. See [Active Enhanced Firewall Policies](#tgw-security-active-enhanced-firewall-policies) below.

`active_forward_proxy_policies` - (Optional) Active Forward Proxy Policies Type. Ordered List of Forward Proxy Policies active. See [Active Forward Proxy Policies](#tgw-security-active-forward-proxy-policies) below.

`active_network_policies` - (Optional) Active Firewall Policies Type. List of firewall policy views. See [Active Network Policies](#tgw-security-active-network-policies) below.

`east_west_service_policy_allow_all` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`forward_proxy_allow_all` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`no_east_west_policy` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`no_forward_proxy` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`no_network_policy` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

<a id="tgw-security-active-east-west-service-policies"></a>

### Tgw Security Active East West Service Policies

`service_policies` - (Optional) Service Policies. A list of references to service_policy objects. See [Service Policies](#tgw-security-active-east-west-service-policies-service-policies) below.

<a id="tgw-security-active-east-west-service-policies-service-policies"></a>

### Tgw Security Active East West Service Policies Service Policies

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

<a id="tgw-security-active-enhanced-firewall-policies"></a>

### Tgw Security Active Enhanced Firewall Policies

`enhanced_firewall_policies` - (Optional) Enhanced Firewall Policy. Ordered List of Enhanced Firewall Policies active. See [Enhanced Firewall Policies](#tgw-security-active-enhanced-firewall-policies-enhanced-firewall-policies) below.

<a id="tgw-security-active-enhanced-firewall-policies-enhanced-firewall-policies"></a>

### Tgw Security Active Enhanced Firewall Policies Enhanced Firewall Policies

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

<a id="tgw-security-active-forward-proxy-policies"></a>

### Tgw Security Active Forward Proxy Policies

`forward_proxy_policies` - (Optional) Forward Proxy Policies. Ordered List of Forward Proxy Policies active. See [Forward Proxy Policies](#tgw-security-active-forward-proxy-policies-forward-proxy-policies) below.

<a id="tgw-security-active-forward-proxy-policies-forward-proxy-policies"></a>

### Tgw Security Active Forward Proxy Policies Forward Proxy Policies

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

<a id="tgw-security-active-network-policies"></a>

### Tgw Security Active Network Policies

`network_policies` - (Optional) Firewall Policy. Ordered List of Firewall Policies active for this network firewall. See [Network Policies](#tgw-security-active-network-policies-network-policies) below.

<a id="tgw-security-active-network-policies-network-policies"></a>

### Tgw Security Active Network Policies Network Policies

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

<a id="timeouts"></a>

### Timeouts

`create` - (Optional) A string that can be [parsed as a duration](`HTTPS://pkg.go.dev/time#ParseDuration`) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours) (`String`).

`delete` - (Optional) A string that can be [parsed as a duration](`HTTPS://pkg.go.dev/time#ParseDuration`) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are saved into state before the destroy operation occurs (`String`).

`read` - (Optional) A string that can be [parsed as a duration](`HTTPS://pkg.go.dev/time#ParseDuration`) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh is enabled (`String`).

`update` - (Optional) A string that can be [parsed as a duration](`HTTPS://pkg.go.dev/time#ParseDuration`) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours) (`String`).

<a id="vn-config"></a>

### Vn Config

`allowed_vip_port` - (Optional) Allowed VIP Ports. This defines the TCP port(s) which will be opened on the cloud loadbalancer. Such that the client can use the cloud VIP IP and port combination to reach TCP/HTTP LB configured on the F5XC Site. See [Allowed VIP Port](#vn-config-allowed-vip-port) below.

`allowed_vip_port_sli` - (Optional) Allowed VIP Ports. This defines the TCP port(s) which will be opened on the cloud loadbalancer. Such that the client can use the cloud VIP IP and port combination to reach TCP/HTTP LB configured on the F5XC Site. See [Allowed VIP Port Sli](#vn-config-allowed-vip-port-sli) below.

`dc_cluster_group_inside_vn` - (Optional) Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name. See [Dc Cluster Group Inside Vn](#vn-config-dc-cluster-group-inside-vn) below.

`dc_cluster_group_outside_vn` - (Optional) Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name. See [Dc Cluster Group Outside Vn](#vn-config-dc-cluster-group-outside-vn) below.

`global_network_list` - (Optional) Global Network Connection List. List of global network connections. See [Global Network List](#vn-config-global-network-list) below.

`inside_static_routes` - (Optional) Static Route List Type. List of static routes. See [Inside Static Routes](#vn-config-inside-static-routes) below.

`no_dc_cluster_group` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`no_global_network` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`no_inside_static_routes` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`no_outside_static_routes` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`outside_static_routes` - (Optional) Static Route List Type. List of static routes. See [Outside Static Routes](#vn-config-outside-static-routes) below.

`sm_connection_public_ip` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`sm_connection_pvt_ip` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

<a id="vn-config-allowed-vip-port"></a>

### Vn Config Allowed VIP Port

`custom_ports` - (Optional) Custom Ports. List of Custom port. See [Custom Ports](#vn-config-allowed-vip-port-custom-ports) below.

`disable_allowed_vip_port` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`use_http_https_port` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`use_http_port` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`use_https_port` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

<a id="vn-config-allowed-vip-port-custom-ports"></a>

### Vn Config Allowed VIP Port Custom Ports

`port_ranges` - (Optional) Port Ranges. Port Ranges (`String`).

<a id="vn-config-allowed-vip-port-sli"></a>

### Vn Config Allowed VIP Port Sli

`custom_ports` - (Optional) Custom Ports. List of Custom port. See [Custom Ports](#vn-config-allowed-vip-port-sli-custom-ports) below.

`disable_allowed_vip_port` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`use_http_https_port` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`use_http_port` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`use_https_port` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

<a id="vn-config-allowed-vip-port-sli-custom-ports"></a>

### Vn Config Allowed VIP Port Sli Custom Ports

`port_ranges` - (Optional) Port Ranges. Port Ranges (`String`).

<a id="vn-config-dc-cluster-group-inside-vn"></a>

### Vn Config Dc Cluster Group Inside Vn

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

<a id="vn-config-dc-cluster-group-outside-vn"></a>

### Vn Config Dc Cluster Group Outside Vn

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

<a id="vn-config-global-network-list"></a>

### Vn Config Global Network List

`global_network_connections` - (Optional) Global Network Connections. Global network connections. See [Global Network Connections](#vn-config-global-network-list-global-network-connections) below.

<a id="vn-config-global-network-list-global-network-connections"></a>

### Vn Config Global Network List Global Network Connections

`sli_to_global_dr` - (Optional) Global Network. Global network reference for direct connection (`Block`).

`slo_to_global_dr` - (Optional) Global Network. Global network reference for direct connection (`Block`).

<a id="vn-config-inside-static-routes"></a>

### Vn Config Inside Static Routes

`static_route_list` - (Optional) List of Static Routes. List of Static routes. See [Static Route List](#vn-config-inside-static-routes-static-route-list) below.

<a id="vn-config-inside-static-routes-static-route-list"></a>

### Vn Config Inside Static Routes Static Route List

`custom_static_route` - (Optional) Static Route. Defines a static route, configuring a list of prefixes and a next-hop to be used for them (`Block`).

`simple_static_route` - (Optional) Simple Static Route. Use simple static route for prefix pointing to single interface in the network (`String`).

<a id="vn-config-outside-static-routes"></a>

### Vn Config Outside Static Routes

`static_route_list` - (Optional) List of Static Routes. List of Static routes. See [Static Route List](#vn-config-outside-static-routes-static-route-list) below.

<a id="vn-config-outside-static-routes-static-route-list"></a>

### Vn Config Outside Static Routes Static Route List

`custom_static_route` - (Optional) Static Route. Defines a static route, configuring a list of prefixes and a next-hop to be used for them (`Block`).

`simple_static_route` - (Optional) Simple Static Route. Use simple static route for prefix pointing to single interface in the network (`String`).

<a id="vpc-attachments"></a>

### Vpc Attachments

`vpc_list` - (Optional) VPC List. List of VPC attachments to transit gateway. See [Vpc List](#vpc-attachments-vpc-list) below.

<a id="vpc-attachments-vpc-list"></a>

### Vpc Attachments Vpc List

`labels` - (Optional) Labels. Add labels for the VPC attachment. These labels can then be used in policies such as enhanced firewall (`Block`).

`vpc_id` - (Optional) VPC ID. Information about existing VPC (`String`).
