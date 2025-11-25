---
page_title: "Workload Nested Blocks - f5xc Provider"
subcategory: "Kubernetes"
description: |-
  Nested block reference for the Workload resource.
---

# Workload Nested Blocks

This page contains detailed documentation for nested blocks in the `f5xc_workload` resource.

For the main resource documentation, see [f5xc_workload](/docs/resources/workload).

## Contents

- [job](#job)
- [job-configuration](#job-configuration)
- [job-configuration-parameters](#job-configuration-parameters)
- [job-containers](#job-containers)
- [job-containers-custom-flavor](#job-containers-custom-flavor)
- [job-containers-image](#job-containers-image)
- [job-containers-liveness-check](#job-containers-liveness-check)
- [job-containers-readiness-check](#job-containers-readiness-check)
- [job-deploy-options](#job-deploy-options)
- [job-deploy-options-deploy-ce-sites](#job-deploy-options-deploy-ce-sites)
- [job-deploy-options-deploy-ce-virtual-sites](#job-deploy-options-deploy-ce-virtual-sites)
- [job-deploy-options-deploy-re-sites](#job-deploy-options-deploy-re-sites)
- [job-deploy-options-deploy-re-virtual-sites](#job-deploy-options-deploy-re-virtual-sites)
- [job-volumes](#job-volumes)
- [job-volumes-empty-dir](#job-volumes-empty-dir)
- [job-volumes-host-path](#job-volumes-host-path)
- [job-volumes-persistent-volume](#job-volumes-persistent-volume)
- [service](#service)
- [service-advertise-options](#service-advertise-options)
- [service-advertise-options-advertise-custom](#service-advertise-options-advertise-custom)
- [service-advertise-options-advertise-in-cluster](#service-advertise-options-advertise-in-cluster)
- [service-advertise-options-advertise-on-public](#service-advertise-options-advertise-on-public)
- [service-configuration](#service-configuration)
- [service-configuration-parameters](#service-configuration-parameters)
- [service-containers](#service-containers)
- [service-containers-custom-flavor](#service-containers-custom-flavor)
- [service-containers-image](#service-containers-image)
- [service-containers-liveness-check](#service-containers-liveness-check)
- [service-containers-readiness-check](#service-containers-readiness-check)
- [service-deploy-options](#service-deploy-options)
- [service-deploy-options-deploy-ce-sites](#service-deploy-options-deploy-ce-sites)
- [service-deploy-options-deploy-ce-virtual-sites](#service-deploy-options-deploy-ce-virtual-sites)
- [service-deploy-options-deploy-re-sites](#service-deploy-options-deploy-re-sites)
- [service-deploy-options-deploy-re-virtual-sites](#service-deploy-options-deploy-re-virtual-sites)
- [service-volumes](#service-volumes)
- [service-volumes-empty-dir](#service-volumes-empty-dir)
- [service-volumes-host-path](#service-volumes-host-path)
- [service-volumes-persistent-volume](#service-volumes-persistent-volume)
- [simple-service](#simple-service)
- [simple-service-configuration](#simple-service-configuration)
- [simple-service-configuration-parameters](#simple-service-configuration-parameters)
- [simple-service-container](#simple-service-container)
- [simple-service-container-custom-flavor](#simple-service-container-custom-flavor)
- [simple-service-container-image](#simple-service-container-image)
- [simple-service-container-liveness-check](#simple-service-container-liveness-check)
- [simple-service-container-readiness-check](#simple-service-container-readiness-check)
- [simple-service-enabled](#simple-service-enabled)
- [simple-service-enabled-persistent-volume](#simple-service-enabled-persistent-volume)
- [simple-service-simple-advertise](#simple-service-simple-advertise)
- [stateful-service](#stateful-service)
- [stateful-service-advertise-options](#stateful-service-advertise-options)
- [stateful-service-advertise-options-advertise-custom](#stateful-service-advertise-options-advertise-custom)
- [stateful-service-advertise-options-advertise-in-cluster](#stateful-service-advertise-options-advertise-in-cluster)
- [stateful-service-advertise-options-advertise-on-public](#stateful-service-advertise-options-advertise-on-public)
- [stateful-service-configuration](#stateful-service-configuration)
- [stateful-service-configuration-parameters](#stateful-service-configuration-parameters)
- [stateful-service-containers](#stateful-service-containers)
- [stateful-service-containers-custom-flavor](#stateful-service-containers-custom-flavor)
- [stateful-service-containers-image](#stateful-service-containers-image)
- [stateful-service-containers-liveness-check](#stateful-service-containers-liveness-check)
- [stateful-service-containers-readiness-check](#stateful-service-containers-readiness-check)
- [stateful-service-deploy-options](#stateful-service-deploy-options)
- [stateful-service-deploy-options-deploy-ce-sites](#stateful-service-deploy-options-deploy-ce-sites)
- [stateful-service-deploy-options-deploy-ce-virtual-sites](#stateful-service-deploy-options-deploy-ce-virtual-sites)
- [stateful-service-deploy-options-deploy-re-sites](#stateful-service-deploy-options-deploy-re-sites)
- [stateful-service-deploy-options-deploy-re-virtual-sites](#stateful-service-deploy-options-deploy-re-virtual-sites)
- [stateful-service-persistent-volumes](#stateful-service-persistent-volumes)
- [stateful-service-persistent-volumes-persistent-volume](#stateful-service-persistent-volumes-persistent-volume)
- [stateful-service-volumes](#stateful-service-volumes)
- [stateful-service-volumes-empty-dir](#stateful-service-volumes-empty-dir)
- [stateful-service-volumes-host-path](#stateful-service-volumes-host-path)
- [timeouts](#timeouts)

---

<a id="job"></a>

### Job

`configuration` - (Optional) Configuration Parameters. Configuration parameters of the workload. See [Configuration](#job-configuration) below.

`containers` - (Optional) Containers. Containers to use for the job. See [Containers](#job-containers) below.

`deploy_options` - (Optional) Deploy Options. Deploy Options are used to configure the workload deployment options. See [Deploy Options](#job-deploy-options) below.

`num_replicas` - (Optional) Number of Replicas. Number of replicas of the batch job to spawn per site (`Number`).

`volumes` - (Optional) Volumes. Volumes for the job. See [Volumes](#job-volumes) below.

<a id="job-configuration"></a>

### Job Configuration

`parameters` - (Optional) Parameters. Parameters for the workload. See [Parameters](#job-configuration-parameters) below.

<a id="job-configuration-parameters"></a>

### Job Configuration Parameters

`env_var` - (Optional) Environment Variable. Environment Variable (`Block`).

`file` - (Optional) Configuration File. Configuration File for the workload (`Block`).

<a id="job-containers"></a>

### Job Containers

`args` - (Optional) Arguments. Arguments to the entrypoint. Overrides the docker image's CMD (`List`).

`command` - (Optional) Command. Command to execute. Overrides the docker image's ENTRYPOINT (`List`).

`custom_flavor` - (Optional) Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name. See [Custom Flavor](#job-containers-custom-flavor) below.

`default_flavor` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`flavor` - (Optional) Container Flavor Type. Container Flavor type - CONTAINER_FLAVOR_TYPE_TINY: Tiny Tiny containers have limit of 0.1 vCPU and 256 MiB (mebibyte) memory - CONTAINER_FLAVOR_TYPE_MEDIUM: Medium Medium containers have limit of 0.25 vCPU and 512 MiB (mebibyte) memory - CONTAINER_FLAVOR_TYPE_LARGE: Large Large containers have limit of 1 vCPU and 2048 MiB (mebibyte) memory. Possible values are `CONTAINER_FLAVOR_TYPE_TINY`, `CONTAINER_FLAVOR_TYPE_MEDIUM`, `CONTAINER_FLAVOR_TYPE_LARGE`. Defaults to `CONTAINER_FLAVOR_TYPE_TINY` (`String`).

`image` - (Optional) Image Configuration. ImageType configures the image to use, how to pull the image, and the associated secrets to use if any. See [Image](#job-containers-image) below.

`init_container` - (Optional) Initialization Container. Specialized container that runs before application container and runs to completion (`Bool`).

`liveness_check` - (Optional) Health Check. HealthCheckType describes a health check to be performed against a container to determine whether it has started up or is alive or ready to receive traffic. See [Liveness Check](#job-containers-liveness-check) below.

`name` - (Optional) Name. Name of the container (`String`).

`readiness_check` - (Optional) Health Check. HealthCheckType describes a health check to be performed against a container to determine whether it has started up or is alive or ready to receive traffic. See [Readiness Check](#job-containers-readiness-check) below.

<a id="job-containers-custom-flavor"></a>

### Job Containers Custom Flavor

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

<a id="job-containers-image"></a>

### Job Containers Image

`container_registry` - (Optional) Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name (`Block`).

`name` - (Optional) Image Name. Name is a container image which are usually given a name such as alpine, ubuntu, or quay.io/etcd:0.13. The format is registry/image:tag or registry/image@image-digest. If registry is not specified, the Docker public registry is assumed. If tag is not specified, latest is assumed (`String`).

`public` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`pull_policy` - (Optional) Image Pull Policy Type. Image pull policy type enumerates the policy choices to use for pulling the image prior to starting the workload - IMAGE_PULL_POLICY_DEFAULT: Default Default will always pull image if :latest tag is specified in image name. If :latest tag is not specified in image name, it will pull image only if it does not already exist on the node - IMAGE_PULL_POLICY_IF_NOT_PRESENT: IfNotPresent Only pull the image if it does not already exist on the node - IMAGE_PULL_POLICY_ALWAYS: Always Always pull the image - IMAGE_PULL_POLICY_NEVER: Never Never pull the image. Possible values are `IMAGE_PULL_POLICY_DEFAULT`, `IMAGE_PULL_POLICY_IF_NOT_PRESENT`, `IMAGE_PULL_POLICY_ALWAYS`, `IMAGE_PULL_POLICY_NEVER`. Defaults to `IMAGE_PULL_POLICY_DEFAULT` (`String`).

<a id="job-containers-liveness-check"></a>

### Job Containers Liveness Check

`exec_health_check` - (Optional) Exec Health Check. ExecHealthCheckType describes a health check based on 'run in container' action. Exit status of 0 is treated as live/healthy and non-zero is unhealthy (`Block`).

`healthy_threshold` - (Optional) Healthy Threshold. Number of consecutive successful responses after having failed before declaring healthy. In other words, this is the number of healthy health checks required before marking healthy. Note that during startup and liveliness, only a single successful health check is required to mark a container healthy (`Number`).

`http_health_check` - (Optional) HTTP Health Check. HTTPHealthCheckType describes a health check based on HTTP GET requests (`Block`).

`initial_delay` - (Optional) Initial Delay. Number of seconds after the container has started before health checks are initiated (`Number`).

`interval` - (Optional) Interval. Time interval in seconds between two health check requests (`Number`).

`tcp_health_check` - (Optional) TCP Health Check. TCPHealthCheckType describes a health check based on opening a TCP connection (`Block`).

`timeout` - (Optional) Timeout. Timeout in seconds to wait for successful response. In other words, it is the time to wait for a health check response. If the timeout is reached the health check attempt will be considered a failure (`Number`).

`unhealthy_threshold` - (Optional) Unhealthy Threshold. Number of consecutive failed responses before declaring unhealthy. In other words, this is the number of unhealthy health checks required before a container is marked unhealthy (`Number`).

<a id="job-containers-readiness-check"></a>

### Job Containers Readiness Check

`exec_health_check` - (Optional) Exec Health Check. ExecHealthCheckType describes a health check based on 'run in container' action. Exit status of 0 is treated as live/healthy and non-zero is unhealthy (`Block`).

`healthy_threshold` - (Optional) Healthy Threshold. Number of consecutive successful responses after having failed before declaring healthy. In other words, this is the number of healthy health checks required before marking healthy. Note that during startup and liveliness, only a single successful health check is required to mark a container healthy (`Number`).

`http_health_check` - (Optional) HTTP Health Check. HTTPHealthCheckType describes a health check based on HTTP GET requests (`Block`).

`initial_delay` - (Optional) Initial Delay. Number of seconds after the container has started before health checks are initiated (`Number`).

`interval` - (Optional) Interval. Time interval in seconds between two health check requests (`Number`).

`tcp_health_check` - (Optional) TCP Health Check. TCPHealthCheckType describes a health check based on opening a TCP connection (`Block`).

`timeout` - (Optional) Timeout. Timeout in seconds to wait for successful response. In other words, it is the time to wait for a health check response. If the timeout is reached the health check attempt will be considered a failure (`Number`).

`unhealthy_threshold` - (Optional) Unhealthy Threshold. Number of consecutive failed responses before declaring unhealthy. In other words, this is the number of unhealthy health checks required before a container is marked unhealthy (`Number`).

<a id="job-deploy-options"></a>

### Job Deploy Options

`all_res` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`default_virtual_sites` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`deploy_ce_sites` - (Optional) Customer Sites. This defines a way to deploy a workload on specific Customer sites. See [Deploy CE Sites](#job-deploy-options-deploy-ce-sites) below.

`deploy_ce_virtual_sites` - (Optional) Customer Virtual Sites. This defines a way to deploy a workload on specific Customer virtual sites. See [Deploy CE Virtual Sites](#job-deploy-options-deploy-ce-virtual-sites) below.

`deploy_re_sites` - (Optional) Regional Edge Sites. This defines a way to deploy a workload on specific Regional Edge sites. See [Deploy RE Sites](#job-deploy-options-deploy-re-sites) below.

`deploy_re_virtual_sites` - (Optional) Regional Edge Virtual Sites. This defines a way to deploy a workload on specific Regional Edge virtual sites. See [Deploy RE Virtual Sites](#job-deploy-options-deploy-re-virtual-sites) below.

<a id="job-deploy-options-deploy-ce-sites"></a>

### Job Deploy Options Deploy CE Sites

`site` - (Optional) List of Customer Sites to Deploy. Which customer sites should this workload be deployed (`Block`).

<a id="job-deploy-options-deploy-ce-virtual-sites"></a>

### Job Deploy Options Deploy CE Virtual Sites

`virtual_site` - (Optional) List of Customer Virtual Sites to Deploy. Which customer virtual sites should this workload be deployed (`Block`).

<a id="job-deploy-options-deploy-re-sites"></a>

### Job Deploy Options Deploy RE Sites

`site` - (Optional) List of Regional Edge Sites to Deploy. Which regional edge sites should this workload be deployed (`Block`).

<a id="job-deploy-options-deploy-re-virtual-sites"></a>

### Job Deploy Options Deploy RE Virtual Sites

`virtual_site` - (Optional) List of Regional Edge Virtual Sites to Deploy. Which regional edge virtual sites should this workload be deployed (`Block`).

<a id="job-volumes"></a>

### Job Volumes

`empty_dir` - (Optional) Empty Directory Volume. Volume containing a temporary directory whose lifetime is the same as a replica of a workload. See [Empty Dir](#job-volumes-empty-dir) below.

`host_path` - (Optional) HostPath Volume. Volume containing a host mapped path into the workload. See [Host Path](#job-volumes-host-path) below.

`name` - (Optional) Name. Name of the volume (`String`).

`persistent_volume` - (Optional) Persistent Storage Volume. Volume containing the Persistent Storage for the workload. See [Persistent Volume](#job-volumes-persistent-volume) below.

<a id="job-volumes-empty-dir"></a>

### Job Volumes Empty Dir

`mount` - (Optional) Volume Mount. Volume mount describes how volume is mounted inside a workload (`Block`).

`size_limit` - (Optional) Size Limit (in GiB) (`Number`).

<a id="job-volumes-host-path"></a>

### Job Volumes Host Path

`mount` - (Optional) Volume Mount. Volume mount describes how volume is mounted inside a workload (`Block`).

`path` - (Optional) Path. Path of the directory on the host (`String`).

<a id="job-volumes-persistent-volume"></a>

### Job Volumes Persistent Volume

`mount` - (Optional) Volume Mount. Volume mount describes how volume is mounted inside a workload (`Block`).

`storage` - (Optional) Persistence Storage Configuration. Persistent storage configuration is used to configure Persistent Volume Claim (PVC) (`Block`).

<a id="service"></a>

### Service

`advertise_options` - (Optional) Advertise Options. Advertise options are used to configure how and where to advertise the workload using load balancers. See [Advertise Options](#service-advertise-options) below.

`configuration` - (Optional) Configuration Parameters. Configuration parameters of the workload. See [Configuration](#service-configuration) below.

`containers` - (Optional) Containers. Containers to use for service. See [Containers](#service-containers) below.

`deploy_options` - (Optional) Deploy Options. Deploy Options are used to configure the workload deployment options. See [Deploy Options](#service-deploy-options) below.

`num_replicas` - (Optional) Number of Replicas. Number of replicas of service to spawn per site (`Number`).

`scale_to_zero` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`volumes` - (Optional) Volumes. Volumes for the service. See [Volumes](#service-volumes) below.

<a id="service-advertise-options"></a>

### Service Advertise Options

`advertise_custom` - (Optional) Advertise on specific sites. Advertise this workload via loadbalancer on specific sites. See [Advertise Custom](#service-advertise-options-advertise-custom) below.

`advertise_in_cluster` - (Optional) Advertise In Cluster. Advertise the workload locally in-cluster. See [Advertise In Cluster](#service-advertise-options-advertise-in-cluster) below.

`advertise_on_public` - (Optional) Advertise On Internet. Advertise this workload via loadbalancer on Internet with default VIP. See [Advertise On Public](#service-advertise-options-advertise-on-public) below.

`do_not_advertise` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

<a id="service-advertise-options-advertise-custom"></a>

### Service Advertise Options Advertise Custom

`advertise_where` - (Optional) List of Sites to Advertise. Where should this load balancer be available (`Block`).

`ports` - (Optional) Ports. Ports to advertise (`Block`).

<a id="service-advertise-options-advertise-in-cluster"></a>

### Service Advertise Options Advertise In Cluster

`multi_ports` - (Optional) Multiple Ports. Multiple ports (`Block`).

`port` - (Optional) Port. Single port (`Block`).

<a id="service-advertise-options-advertise-on-public"></a>

### Service Advertise Options Advertise On Public

`multi_ports` - (Optional) Advertise Multiple Ports. Advertise multiple ports (`Block`).

`port` - (Optional) Advertise Port. Advertise single port (`Block`).

<a id="service-configuration"></a>

### Service Configuration

`parameters` - (Optional) Parameters. Parameters for the workload. See [Parameters](#service-configuration-parameters) below.

<a id="service-configuration-parameters"></a>

### Service Configuration Parameters

`env_var` - (Optional) Environment Variable. Environment Variable (`Block`).

`file` - (Optional) Configuration File. Configuration File for the workload (`Block`).

<a id="service-containers"></a>

### Service Containers

`args` - (Optional) Arguments. Arguments to the entrypoint. Overrides the docker image's CMD (`List`).

`command` - (Optional) Command. Command to execute. Overrides the docker image's ENTRYPOINT (`List`).

`custom_flavor` - (Optional) Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name. See [Custom Flavor](#service-containers-custom-flavor) below.

`default_flavor` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`flavor` - (Optional) Container Flavor Type. Container Flavor type - CONTAINER_FLAVOR_TYPE_TINY: Tiny Tiny containers have limit of 0.1 vCPU and 256 MiB (mebibyte) memory - CONTAINER_FLAVOR_TYPE_MEDIUM: Medium Medium containers have limit of 0.25 vCPU and 512 MiB (mebibyte) memory - CONTAINER_FLAVOR_TYPE_LARGE: Large Large containers have limit of 1 vCPU and 2048 MiB (mebibyte) memory. Possible values are `CONTAINER_FLAVOR_TYPE_TINY`, `CONTAINER_FLAVOR_TYPE_MEDIUM`, `CONTAINER_FLAVOR_TYPE_LARGE`. Defaults to `CONTAINER_FLAVOR_TYPE_TINY` (`String`).

`image` - (Optional) Image Configuration. ImageType configures the image to use, how to pull the image, and the associated secrets to use if any. See [Image](#service-containers-image) below.

`init_container` - (Optional) Initialization Container. Specialized container that runs before application container and runs to completion (`Bool`).

`liveness_check` - (Optional) Health Check. HealthCheckType describes a health check to be performed against a container to determine whether it has started up or is alive or ready to receive traffic. See [Liveness Check](#service-containers-liveness-check) below.

`name` - (Optional) Name. Name of the container (`String`).

`readiness_check` - (Optional) Health Check. HealthCheckType describes a health check to be performed against a container to determine whether it has started up or is alive or ready to receive traffic. See [Readiness Check](#service-containers-readiness-check) below.

<a id="service-containers-custom-flavor"></a>

### Service Containers Custom Flavor

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

<a id="service-containers-image"></a>

### Service Containers Image

`container_registry` - (Optional) Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name (`Block`).

`name` - (Optional) Image Name. Name is a container image which are usually given a name such as alpine, ubuntu, or quay.io/etcd:0.13. The format is registry/image:tag or registry/image@image-digest. If registry is not specified, the Docker public registry is assumed. If tag is not specified, latest is assumed (`String`).

`public` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`pull_policy` - (Optional) Image Pull Policy Type. Image pull policy type enumerates the policy choices to use for pulling the image prior to starting the workload - IMAGE_PULL_POLICY_DEFAULT: Default Default will always pull image if :latest tag is specified in image name. If :latest tag is not specified in image name, it will pull image only if it does not already exist on the node - IMAGE_PULL_POLICY_IF_NOT_PRESENT: IfNotPresent Only pull the image if it does not already exist on the node - IMAGE_PULL_POLICY_ALWAYS: Always Always pull the image - IMAGE_PULL_POLICY_NEVER: Never Never pull the image. Possible values are `IMAGE_PULL_POLICY_DEFAULT`, `IMAGE_PULL_POLICY_IF_NOT_PRESENT`, `IMAGE_PULL_POLICY_ALWAYS`, `IMAGE_PULL_POLICY_NEVER`. Defaults to `IMAGE_PULL_POLICY_DEFAULT` (`String`).

<a id="service-containers-liveness-check"></a>

### Service Containers Liveness Check

`exec_health_check` - (Optional) Exec Health Check. ExecHealthCheckType describes a health check based on 'run in container' action. Exit status of 0 is treated as live/healthy and non-zero is unhealthy (`Block`).

`healthy_threshold` - (Optional) Healthy Threshold. Number of consecutive successful responses after having failed before declaring healthy. In other words, this is the number of healthy health checks required before marking healthy. Note that during startup and liveliness, only a single successful health check is required to mark a container healthy (`Number`).

`http_health_check` - (Optional) HTTP Health Check. HTTPHealthCheckType describes a health check based on HTTP GET requests (`Block`).

`initial_delay` - (Optional) Initial Delay. Number of seconds after the container has started before health checks are initiated (`Number`).

`interval` - (Optional) Interval. Time interval in seconds between two health check requests (`Number`).

`tcp_health_check` - (Optional) TCP Health Check. TCPHealthCheckType describes a health check based on opening a TCP connection (`Block`).

`timeout` - (Optional) Timeout. Timeout in seconds to wait for successful response. In other words, it is the time to wait for a health check response. If the timeout is reached the health check attempt will be considered a failure (`Number`).

`unhealthy_threshold` - (Optional) Unhealthy Threshold. Number of consecutive failed responses before declaring unhealthy. In other words, this is the number of unhealthy health checks required before a container is marked unhealthy (`Number`).

<a id="service-containers-readiness-check"></a>

### Service Containers Readiness Check

`exec_health_check` - (Optional) Exec Health Check. ExecHealthCheckType describes a health check based on 'run in container' action. Exit status of 0 is treated as live/healthy and non-zero is unhealthy (`Block`).

`healthy_threshold` - (Optional) Healthy Threshold. Number of consecutive successful responses after having failed before declaring healthy. In other words, this is the number of healthy health checks required before marking healthy. Note that during startup and liveliness, only a single successful health check is required to mark a container healthy (`Number`).

`http_health_check` - (Optional) HTTP Health Check. HTTPHealthCheckType describes a health check based on HTTP GET requests (`Block`).

`initial_delay` - (Optional) Initial Delay. Number of seconds after the container has started before health checks are initiated (`Number`).

`interval` - (Optional) Interval. Time interval in seconds between two health check requests (`Number`).

`tcp_health_check` - (Optional) TCP Health Check. TCPHealthCheckType describes a health check based on opening a TCP connection (`Block`).

`timeout` - (Optional) Timeout. Timeout in seconds to wait for successful response. In other words, it is the time to wait for a health check response. If the timeout is reached the health check attempt will be considered a failure (`Number`).

`unhealthy_threshold` - (Optional) Unhealthy Threshold. Number of consecutive failed responses before declaring unhealthy. In other words, this is the number of unhealthy health checks required before a container is marked unhealthy (`Number`).

<a id="service-deploy-options"></a>

### Service Deploy Options

`all_res` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`default_virtual_sites` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`deploy_ce_sites` - (Optional) Customer Sites. This defines a way to deploy a workload on specific Customer sites. See [Deploy CE Sites](#service-deploy-options-deploy-ce-sites) below.

`deploy_ce_virtual_sites` - (Optional) Customer Virtual Sites. This defines a way to deploy a workload on specific Customer virtual sites. See [Deploy CE Virtual Sites](#service-deploy-options-deploy-ce-virtual-sites) below.

`deploy_re_sites` - (Optional) Regional Edge Sites. This defines a way to deploy a workload on specific Regional Edge sites. See [Deploy RE Sites](#service-deploy-options-deploy-re-sites) below.

`deploy_re_virtual_sites` - (Optional) Regional Edge Virtual Sites. This defines a way to deploy a workload on specific Regional Edge virtual sites. See [Deploy RE Virtual Sites](#service-deploy-options-deploy-re-virtual-sites) below.

<a id="service-deploy-options-deploy-ce-sites"></a>

### Service Deploy Options Deploy CE Sites

`site` - (Optional) List of Customer Sites to Deploy. Which customer sites should this workload be deployed (`Block`).

<a id="service-deploy-options-deploy-ce-virtual-sites"></a>

### Service Deploy Options Deploy CE Virtual Sites

`virtual_site` - (Optional) List of Customer Virtual Sites to Deploy. Which customer virtual sites should this workload be deployed (`Block`).

<a id="service-deploy-options-deploy-re-sites"></a>

### Service Deploy Options Deploy RE Sites

`site` - (Optional) List of Regional Edge Sites to Deploy. Which regional edge sites should this workload be deployed (`Block`).

<a id="service-deploy-options-deploy-re-virtual-sites"></a>

### Service Deploy Options Deploy RE Virtual Sites

`virtual_site` - (Optional) List of Regional Edge Virtual Sites to Deploy. Which regional edge virtual sites should this workload be deployed (`Block`).

<a id="service-volumes"></a>

### Service Volumes

`empty_dir` - (Optional) Empty Directory Volume. Volume containing a temporary directory whose lifetime is the same as a replica of a workload. See [Empty Dir](#service-volumes-empty-dir) below.

`host_path` - (Optional) HostPath Volume. Volume containing a host mapped path into the workload. See [Host Path](#service-volumes-host-path) below.

`name` - (Optional) Name. Name of the volume (`String`).

`persistent_volume` - (Optional) Persistent Storage Volume. Volume containing the Persistent Storage for the workload. See [Persistent Volume](#service-volumes-persistent-volume) below.

<a id="service-volumes-empty-dir"></a>

### Service Volumes Empty Dir

`mount` - (Optional) Volume Mount. Volume mount describes how volume is mounted inside a workload (`Block`).

`size_limit` - (Optional) Size Limit (in GiB) (`Number`).

<a id="service-volumes-host-path"></a>

### Service Volumes Host Path

`mount` - (Optional) Volume Mount. Volume mount describes how volume is mounted inside a workload (`Block`).

`path` - (Optional) Path. Path of the directory on the host (`String`).

<a id="service-volumes-persistent-volume"></a>

### Service Volumes Persistent Volume

`mount` - (Optional) Volume Mount. Volume mount describes how volume is mounted inside a workload (`Block`).

`storage` - (Optional) Persistence Storage Configuration. Persistent storage configuration is used to configure Persistent Volume Claim (PVC) (`Block`).

<a id="simple-service"></a>

### Simple Service

`configuration` - (Optional) Configuration Parameters. Configuration parameters of the workload. See [Configuration](#simple-service-configuration) below.

`container` - (Optional) Container Configuration. ContainerType configures the container information. See [Container](#simple-service-container) below.

`disabled` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`do_not_advertise` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`enabled` - (Optional) Persistent Storage Volume. Persistent storage volume configuration for the workload. See [Enabled](#simple-service-enabled) below.

`scale_to_zero` - (Optional) Scale Down to Zero. Scale down replicas of the service to zero (`Bool`).

`simple_advertise` - (Optional) Advertise Options For Simple Service. Advertise options for Simple Service. See [Simple Advertise](#simple-service-simple-advertise) below.

<a id="simple-service-configuration"></a>

### Simple Service Configuration

`parameters` - (Optional) Parameters. Parameters for the workload. See [Parameters](#simple-service-configuration-parameters) below.

<a id="simple-service-configuration-parameters"></a>

### Simple Service Configuration Parameters

`env_var` - (Optional) Environment Variable. Environment Variable (`Block`).

`file` - (Optional) Configuration File. Configuration File for the workload (`Block`).

<a id="simple-service-container"></a>

### Simple Service Container

`args` - (Optional) Arguments. Arguments to the entrypoint. Overrides the docker image's CMD (`List`).

`command` - (Optional) Command. Command to execute. Overrides the docker image's ENTRYPOINT (`List`).

`custom_flavor` - (Optional) Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name. See [Custom Flavor](#simple-service-container-custom-flavor) below.

`default_flavor` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`flavor` - (Optional) Container Flavor Type. Container Flavor type - CONTAINER_FLAVOR_TYPE_TINY: Tiny Tiny containers have limit of 0.1 vCPU and 256 MiB (mebibyte) memory - CONTAINER_FLAVOR_TYPE_MEDIUM: Medium Medium containers have limit of 0.25 vCPU and 512 MiB (mebibyte) memory - CONTAINER_FLAVOR_TYPE_LARGE: Large Large containers have limit of 1 vCPU and 2048 MiB (mebibyte) memory. Possible values are `CONTAINER_FLAVOR_TYPE_TINY`, `CONTAINER_FLAVOR_TYPE_MEDIUM`, `CONTAINER_FLAVOR_TYPE_LARGE`. Defaults to `CONTAINER_FLAVOR_TYPE_TINY` (`String`).

`image` - (Optional) Image Configuration. ImageType configures the image to use, how to pull the image, and the associated secrets to use if any. See [Image](#simple-service-container-image) below.

`init_container` - (Optional) Initialization Container. Specialized container that runs before application container and runs to completion (`Bool`).

`liveness_check` - (Optional) Health Check. HealthCheckType describes a health check to be performed against a container to determine whether it has started up or is alive or ready to receive traffic. See [Liveness Check](#simple-service-container-liveness-check) below.

`name` - (Optional) Name. Name of the container (`String`).

`readiness_check` - (Optional) Health Check. HealthCheckType describes a health check to be performed against a container to determine whether it has started up or is alive or ready to receive traffic. See [Readiness Check](#simple-service-container-readiness-check) below.

<a id="simple-service-container-custom-flavor"></a>

### Simple Service Container Custom Flavor

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

<a id="simple-service-container-image"></a>

### Simple Service Container Image

`container_registry` - (Optional) Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name (`Block`).

`name` - (Optional) Image Name. Name is a container image which are usually given a name such as alpine, ubuntu, or quay.io/etcd:0.13. The format is registry/image:tag or registry/image@image-digest. If registry is not specified, the Docker public registry is assumed. If tag is not specified, latest is assumed (`String`).

`public` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`pull_policy` - (Optional) Image Pull Policy Type. Image pull policy type enumerates the policy choices to use for pulling the image prior to starting the workload - IMAGE_PULL_POLICY_DEFAULT: Default Default will always pull image if :latest tag is specified in image name. If :latest tag is not specified in image name, it will pull image only if it does not already exist on the node - IMAGE_PULL_POLICY_IF_NOT_PRESENT: IfNotPresent Only pull the image if it does not already exist on the node - IMAGE_PULL_POLICY_ALWAYS: Always Always pull the image - IMAGE_PULL_POLICY_NEVER: Never Never pull the image. Possible values are `IMAGE_PULL_POLICY_DEFAULT`, `IMAGE_PULL_POLICY_IF_NOT_PRESENT`, `IMAGE_PULL_POLICY_ALWAYS`, `IMAGE_PULL_POLICY_NEVER`. Defaults to `IMAGE_PULL_POLICY_DEFAULT` (`String`).

<a id="simple-service-container-liveness-check"></a>

### Simple Service Container Liveness Check

`exec_health_check` - (Optional) Exec Health Check. ExecHealthCheckType describes a health check based on 'run in container' action. Exit status of 0 is treated as live/healthy and non-zero is unhealthy (`Block`).

`healthy_threshold` - (Optional) Healthy Threshold. Number of consecutive successful responses after having failed before declaring healthy. In other words, this is the number of healthy health checks required before marking healthy. Note that during startup and liveliness, only a single successful health check is required to mark a container healthy (`Number`).

`http_health_check` - (Optional) HTTP Health Check. HTTPHealthCheckType describes a health check based on HTTP GET requests (`Block`).

`initial_delay` - (Optional) Initial Delay. Number of seconds after the container has started before health checks are initiated (`Number`).

`interval` - (Optional) Interval. Time interval in seconds between two health check requests (`Number`).

`tcp_health_check` - (Optional) TCP Health Check. TCPHealthCheckType describes a health check based on opening a TCP connection (`Block`).

`timeout` - (Optional) Timeout. Timeout in seconds to wait for successful response. In other words, it is the time to wait for a health check response. If the timeout is reached the health check attempt will be considered a failure (`Number`).

`unhealthy_threshold` - (Optional) Unhealthy Threshold. Number of consecutive failed responses before declaring unhealthy. In other words, this is the number of unhealthy health checks required before a container is marked unhealthy (`Number`).

<a id="simple-service-container-readiness-check"></a>

### Simple Service Container Readiness Check

`exec_health_check` - (Optional) Exec Health Check. ExecHealthCheckType describes a health check based on 'run in container' action. Exit status of 0 is treated as live/healthy and non-zero is unhealthy (`Block`).

`healthy_threshold` - (Optional) Healthy Threshold. Number of consecutive successful responses after having failed before declaring healthy. In other words, this is the number of healthy health checks required before marking healthy. Note that during startup and liveliness, only a single successful health check is required to mark a container healthy (`Number`).

`http_health_check` - (Optional) HTTP Health Check. HTTPHealthCheckType describes a health check based on HTTP GET requests (`Block`).

`initial_delay` - (Optional) Initial Delay. Number of seconds after the container has started before health checks are initiated (`Number`).

`interval` - (Optional) Interval. Time interval in seconds between two health check requests (`Number`).

`tcp_health_check` - (Optional) TCP Health Check. TCPHealthCheckType describes a health check based on opening a TCP connection (`Block`).

`timeout` - (Optional) Timeout. Timeout in seconds to wait for successful response. In other words, it is the time to wait for a health check response. If the timeout is reached the health check attempt will be considered a failure (`Number`).

`unhealthy_threshold` - (Optional) Unhealthy Threshold. Number of consecutive failed responses before declaring unhealthy. In other words, this is the number of unhealthy health checks required before a container is marked unhealthy (`Number`).

<a id="simple-service-enabled"></a>

### Simple Service Enabled

`name` - (Optional) Name. Name of the volume (`String`).

`persistent_volume` - (Optional) Persistent Storage Volume. Volume containing the Persistent Storage for the workload. See [Persistent Volume](#simple-service-enabled-persistent-volume) below.

<a id="simple-service-enabled-persistent-volume"></a>

### Simple Service Enabled Persistent Volume

`mount` - (Optional) Volume Mount. Volume mount describes how volume is mounted inside a workload (`Block`).

`storage` - (Optional) Persistence Storage Configuration. Persistent storage configuration is used to configure Persistent Volume Claim (PVC) (`Block`).

<a id="simple-service-simple-advertise"></a>

### Simple Service Simple Advertise

`domains` - (Optional) Domains. A list of Domains (host/authority header) that will be matched to Load Balancer. Wildcard hosts are supported in the suffix or prefix form Supported Domains and search order: 1. Exact Domain names: `www.foo.com.` 2. Domains starting with a Wildcard: *.foo.com. Not supported Domains: - Just a Wildcard: * - A Wildcard and TLD with no root Domain: *.com. - A Wildcard not matching a whole DNS label. e.g. *.foo.com and *.bar.foo.com are valid Wildcards however *bar.foo.com, *-bar.foo.com, and bar*.foo.com are all invalid. Additional notes: A Wildcard will not match empty string. e.g. *.foo.com will match bar.foo.com and baz-bar.foo.com but not .foo.com. The longest Wildcards match first. Only a single virtual host in the entire route configuration can match on *. Also a Domain must be unique across all virtual hosts within an advertise policy. Domains are also used for SNI matching if the Load Balancer type is HTTPS. Domains also indicate the list of names for which DNS resolution will be automatically resolved to IP addresses by the system (`List`).

`service_port` - (Optional) Service Port. Service port to advertise on Internet via HTTP loadbalancer using port 80 (`Number`).

<a id="stateful-service"></a>

### Stateful Service

`advertise_options` - (Optional) Advertise Options. Advertise options are used to configure how and where to advertise the workload using load balancers. See [Advertise Options](#stateful-service-advertise-options) below.

`configuration` - (Optional) Configuration Parameters. Configuration parameters of the workload. See [Configuration](#stateful-service-configuration) below.

`containers` - (Optional) Containers. Containers to use for service. See [Containers](#stateful-service-containers) below.

`deploy_options` - (Optional) Deploy Options. Deploy Options are used to configure the workload deployment options. See [Deploy Options](#stateful-service-deploy-options) below.

`num_replicas` - (Optional) Number of Replicas. Number of replicas of service to spawn per site (`Number`).

`persistent_volumes` - (Optional) Persistent Storage Configuration. Persistent storage configuration for the service. See [Persistent Volumes](#stateful-service-persistent-volumes) below.

`scale_to_zero` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`volumes` - (Optional) Ephemeral Volumes. Ephemeral volumes for the service. See [Volumes](#stateful-service-volumes) below.

<a id="stateful-service-advertise-options"></a>

### Stateful Service Advertise Options

`advertise_custom` - (Optional) Advertise on specific sites. Advertise this workload via loadbalancer on specific sites. See [Advertise Custom](#stateful-service-advertise-options-advertise-custom) below.

`advertise_in_cluster` - (Optional) Advertise In Cluster. Advertise the workload locally in-cluster. See [Advertise In Cluster](#stateful-service-advertise-options-advertise-in-cluster) below.

`advertise_on_public` - (Optional) Advertise On Internet. Advertise this workload via loadbalancer on Internet with default VIP. See [Advertise On Public](#stateful-service-advertise-options-advertise-on-public) below.

`do_not_advertise` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

<a id="stateful-service-advertise-options-advertise-custom"></a>

### Stateful Service Advertise Options Advertise Custom

`advertise_where` - (Optional) List of Sites to Advertise. Where should this load balancer be available (`Block`).

`ports` - (Optional) Ports. Ports to advertise (`Block`).

<a id="stateful-service-advertise-options-advertise-in-cluster"></a>

### Stateful Service Advertise Options Advertise In Cluster

`multi_ports` - (Optional) Multiple Ports. Multiple ports (`Block`).

`port` - (Optional) Port. Single port (`Block`).

<a id="stateful-service-advertise-options-advertise-on-public"></a>

### Stateful Service Advertise Options Advertise On Public

`multi_ports` - (Optional) Advertise Multiple Ports. Advertise multiple ports (`Block`).

`port` - (Optional) Advertise Port. Advertise single port (`Block`).

<a id="stateful-service-configuration"></a>

### Stateful Service Configuration

`parameters` - (Optional) Parameters. Parameters for the workload. See [Parameters](#stateful-service-configuration-parameters) below.

<a id="stateful-service-configuration-parameters"></a>

### Stateful Service Configuration Parameters

`env_var` - (Optional) Environment Variable. Environment Variable (`Block`).

`file` - (Optional) Configuration File. Configuration File for the workload (`Block`).

<a id="stateful-service-containers"></a>

### Stateful Service Containers

`args` - (Optional) Arguments. Arguments to the entrypoint. Overrides the docker image's CMD (`List`).

`command` - (Optional) Command. Command to execute. Overrides the docker image's ENTRYPOINT (`List`).

`custom_flavor` - (Optional) Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name. See [Custom Flavor](#stateful-service-containers-custom-flavor) below.

`default_flavor` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`flavor` - (Optional) Container Flavor Type. Container Flavor type - CONTAINER_FLAVOR_TYPE_TINY: Tiny Tiny containers have limit of 0.1 vCPU and 256 MiB (mebibyte) memory - CONTAINER_FLAVOR_TYPE_MEDIUM: Medium Medium containers have limit of 0.25 vCPU and 512 MiB (mebibyte) memory - CONTAINER_FLAVOR_TYPE_LARGE: Large Large containers have limit of 1 vCPU and 2048 MiB (mebibyte) memory. Possible values are `CONTAINER_FLAVOR_TYPE_TINY`, `CONTAINER_FLAVOR_TYPE_MEDIUM`, `CONTAINER_FLAVOR_TYPE_LARGE`. Defaults to `CONTAINER_FLAVOR_TYPE_TINY` (`String`).

`image` - (Optional) Image Configuration. ImageType configures the image to use, how to pull the image, and the associated secrets to use if any. See [Image](#stateful-service-containers-image) below.

`init_container` - (Optional) Initialization Container. Specialized container that runs before application container and runs to completion (`Bool`).

`liveness_check` - (Optional) Health Check. HealthCheckType describes a health check to be performed against a container to determine whether it has started up or is alive or ready to receive traffic. See [Liveness Check](#stateful-service-containers-liveness-check) below.

`name` - (Optional) Name. Name of the container (`String`).

`readiness_check` - (Optional) Health Check. HealthCheckType describes a health check to be performed against a container to determine whether it has started up or is alive or ready to receive traffic. See [Readiness Check](#stateful-service-containers-readiness-check) below.

<a id="stateful-service-containers-custom-flavor"></a>

### Stateful Service Containers Custom Flavor

`name` - (Optional) Name. When a configuration object(e.g. virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. route's) name (`String`).

`namespace` - (Optional) Namespace. When a configuration object(e.g. virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. route's) namespace (`String`).

`tenant` - (Optional) Tenant. When a configuration object(e.g. virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. route's) tenant (`String`).

<a id="stateful-service-containers-image"></a>

### Stateful Service Containers Image

`container_registry` - (Optional) Object reference. This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name (`Block`).

`name` - (Optional) Image Name. Name is a container image which are usually given a name such as alpine, ubuntu, or quay.io/etcd:0.13. The format is registry/image:tag or registry/image@image-digest. If registry is not specified, the Docker public registry is assumed. If tag is not specified, latest is assumed (`String`).

`public` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`pull_policy` - (Optional) Image Pull Policy Type. Image pull policy type enumerates the policy choices to use for pulling the image prior to starting the workload - IMAGE_PULL_POLICY_DEFAULT: Default Default will always pull image if :latest tag is specified in image name. If :latest tag is not specified in image name, it will pull image only if it does not already exist on the node - IMAGE_PULL_POLICY_IF_NOT_PRESENT: IfNotPresent Only pull the image if it does not already exist on the node - IMAGE_PULL_POLICY_ALWAYS: Always Always pull the image - IMAGE_PULL_POLICY_NEVER: Never Never pull the image. Possible values are `IMAGE_PULL_POLICY_DEFAULT`, `IMAGE_PULL_POLICY_IF_NOT_PRESENT`, `IMAGE_PULL_POLICY_ALWAYS`, `IMAGE_PULL_POLICY_NEVER`. Defaults to `IMAGE_PULL_POLICY_DEFAULT` (`String`).

<a id="stateful-service-containers-liveness-check"></a>

### Stateful Service Containers Liveness Check

`exec_health_check` - (Optional) Exec Health Check. ExecHealthCheckType describes a health check based on 'run in container' action. Exit status of 0 is treated as live/healthy and non-zero is unhealthy (`Block`).

`healthy_threshold` - (Optional) Healthy Threshold. Number of consecutive successful responses after having failed before declaring healthy. In other words, this is the number of healthy health checks required before marking healthy. Note that during startup and liveliness, only a single successful health check is required to mark a container healthy (`Number`).

`http_health_check` - (Optional) HTTP Health Check. HTTPHealthCheckType describes a health check based on HTTP GET requests (`Block`).

`initial_delay` - (Optional) Initial Delay. Number of seconds after the container has started before health checks are initiated (`Number`).

`interval` - (Optional) Interval. Time interval in seconds between two health check requests (`Number`).

`tcp_health_check` - (Optional) TCP Health Check. TCPHealthCheckType describes a health check based on opening a TCP connection (`Block`).

`timeout` - (Optional) Timeout. Timeout in seconds to wait for successful response. In other words, it is the time to wait for a health check response. If the timeout is reached the health check attempt will be considered a failure (`Number`).

`unhealthy_threshold` - (Optional) Unhealthy Threshold. Number of consecutive failed responses before declaring unhealthy. In other words, this is the number of unhealthy health checks required before a container is marked unhealthy (`Number`).

<a id="stateful-service-containers-readiness-check"></a>

### Stateful Service Containers Readiness Check

`exec_health_check` - (Optional) Exec Health Check. ExecHealthCheckType describes a health check based on 'run in container' action. Exit status of 0 is treated as live/healthy and non-zero is unhealthy (`Block`).

`healthy_threshold` - (Optional) Healthy Threshold. Number of consecutive successful responses after having failed before declaring healthy. In other words, this is the number of healthy health checks required before marking healthy. Note that during startup and liveliness, only a single successful health check is required to mark a container healthy (`Number`).

`http_health_check` - (Optional) HTTP Health Check. HTTPHealthCheckType describes a health check based on HTTP GET requests (`Block`).

`initial_delay` - (Optional) Initial Delay. Number of seconds after the container has started before health checks are initiated (`Number`).

`interval` - (Optional) Interval. Time interval in seconds between two health check requests (`Number`).

`tcp_health_check` - (Optional) TCP Health Check. TCPHealthCheckType describes a health check based on opening a TCP connection (`Block`).

`timeout` - (Optional) Timeout. Timeout in seconds to wait for successful response. In other words, it is the time to wait for a health check response. If the timeout is reached the health check attempt will be considered a failure (`Number`).

`unhealthy_threshold` - (Optional) Unhealthy Threshold. Number of consecutive failed responses before declaring unhealthy. In other words, this is the number of unhealthy health checks required before a container is marked unhealthy (`Number`).

<a id="stateful-service-deploy-options"></a>

### Stateful Service Deploy Options

`all_res` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`default_virtual_sites` - (Optional) Empty. This can be used for messages where no values are needed (`Block`).

`deploy_ce_sites` - (Optional) Customer Sites. This defines a way to deploy a workload on specific Customer sites. See [Deploy CE Sites](#stateful-service-deploy-options-deploy-ce-sites) below.

`deploy_ce_virtual_sites` - (Optional) Customer Virtual Sites. This defines a way to deploy a workload on specific Customer virtual sites. See [Deploy CE Virtual Sites](#stateful-service-deploy-options-deploy-ce-virtual-sites) below.

`deploy_re_sites` - (Optional) Regional Edge Sites. This defines a way to deploy a workload on specific Regional Edge sites. See [Deploy RE Sites](#stateful-service-deploy-options-deploy-re-sites) below.

`deploy_re_virtual_sites` - (Optional) Regional Edge Virtual Sites. This defines a way to deploy a workload on specific Regional Edge virtual sites. See [Deploy RE Virtual Sites](#stateful-service-deploy-options-deploy-re-virtual-sites) below.

<a id="stateful-service-deploy-options-deploy-ce-sites"></a>

### Stateful Service Deploy Options Deploy CE Sites

`site` - (Optional) List of Customer Sites to Deploy. Which customer sites should this workload be deployed (`Block`).

<a id="stateful-service-deploy-options-deploy-ce-virtual-sites"></a>

### Stateful Service Deploy Options Deploy CE Virtual Sites

`virtual_site` - (Optional) List of Customer Virtual Sites to Deploy. Which customer virtual sites should this workload be deployed (`Block`).

<a id="stateful-service-deploy-options-deploy-re-sites"></a>

### Stateful Service Deploy Options Deploy RE Sites

`site` - (Optional) List of Regional Edge Sites to Deploy. Which regional edge sites should this workload be deployed (`Block`).

<a id="stateful-service-deploy-options-deploy-re-virtual-sites"></a>

### Stateful Service Deploy Options Deploy RE Virtual Sites

`virtual_site` - (Optional) List of Regional Edge Virtual Sites to Deploy. Which regional edge virtual sites should this workload be deployed (`Block`).

<a id="stateful-service-persistent-volumes"></a>

### Stateful Service Persistent Volumes

`name` - (Optional) Name. Name of the volume (`String`).

`persistent_volume` - (Optional) Persistent Storage Volume. Volume containing the Persistent Storage for the workload. See [Persistent Volume](#stateful-service-persistent-volumes-persistent-volume) below.

<a id="stateful-service-persistent-volumes-persistent-volume"></a>

### Stateful Service Persistent Volumes Persistent Volume

`mount` - (Optional) Volume Mount. Volume mount describes how volume is mounted inside a workload (`Block`).

`storage` - (Optional) Persistence Storage Configuration. Persistent storage configuration is used to configure Persistent Volume Claim (PVC) (`Block`).

<a id="stateful-service-volumes"></a>

### Stateful Service Volumes

`empty_dir` - (Optional) Empty Directory Volume. Volume containing a temporary directory whose lifetime is the same as a replica of a workload. See [Empty Dir](#stateful-service-volumes-empty-dir) below.

`host_path` - (Optional) HostPath Volume. Volume containing a host mapped path into the workload. See [Host Path](#stateful-service-volumes-host-path) below.

`name` - (Optional) Name. Name of the volume (`String`).

<a id="stateful-service-volumes-empty-dir"></a>

### Stateful Service Volumes Empty Dir

`mount` - (Optional) Volume Mount. Volume mount describes how volume is mounted inside a workload (`Block`).

`size_limit` - (Optional) Size Limit (in GiB) (`Number`).

<a id="stateful-service-volumes-host-path"></a>

### Stateful Service Volumes Host Path

`mount` - (Optional) Volume Mount. Volume mount describes how volume is mounted inside a workload (`Block`).

`path` - (Optional) Path. Path of the directory on the host (`String`).

<a id="timeouts"></a>

### Timeouts

`create` - (Optional) A string that can be [parsed as a duration](`HTTPS://pkg.go.dev/time#ParseDuration`) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours) (`String`).

`delete` - (Optional) A string that can be [parsed as a duration](`HTTPS://pkg.go.dev/time#ParseDuration`) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are saved into state before the destroy operation occurs (`String`).

`read` - (Optional) A string that can be [parsed as a duration](`HTTPS://pkg.go.dev/time#ParseDuration`) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh is enabled (`String`).

`update` - (Optional) A string that can be [parsed as a duration](`HTTPS://pkg.go.dev/time#ParseDuration`) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours) (`String`).
