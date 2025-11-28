package summary

import (
	"testing"

	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"

	"github.com/rancher/wrangler/v3/pkg/yaml"
)

var nodeYaml = []byte(`apiVersion: v1
kind: Node
metadata:
  annotations:
    alpha.kubernetes.io/provided-node-ip: 10.115.50.51
    cluster.x-k8s.io/cluster-name: local
    cluster.x-k8s.io/cluster-namespace: fleet-local
    cluster.x-k8s.io/labels-from-machine: ""
    cluster.x-k8s.io/machine: custom-d8ea5ba6e941
    csi.volume.kubernetes.io/nodeid: '{"driver.longhorn.io":"epyc-milan"}'
    flannel.alpha.coreos.com/backend-data: '{"VNI":1,"VtepMAC":"0e:4d:1b:b8:34:c8"}'
    flannel.alpha.coreos.com/backend-type: vxlan
    flannel.alpha.coreos.com/kube-subnet-manager: "true"
    flannel.alpha.coreos.com/public-ip: 10.115.50.51
    kubevirt.io/heartbeat: "2025-08-29T03:34:18Z"
    kubevirt.io/ksm-handler-managed: "false"
    management.cattle.io/pod-limits: '{"cpu":"6115m","devices.kubevirt.io/kvm":"1","devices.kubevirt.io/tun":"1","devices.kubevirt.io/vhost-net":"1","memory":"9944924385"}'
    management.cattle.io/pod-requests: '{"cpu":"3523m","devices.kubevirt.io/kvm":"1","devices.kubevirt.io/tun":"1","devices.kubevirt.io/vhost-net":"1","ephemeral-storage":"50M","memory":"5537925281","pods":"25"}'
    node.alpha.kubernetes.io/ttl: "0"
    node.harvesterhci.io/ntp-service: '{"ntpSyncStatus":"synced","currentNtpServers":"0.suse.pool.ntp.org"}'
    rke2.io/hostname: epyc-milan
    rke2.io/internal-ip: 10.115.50.51
    rke2.io/node-args: '["agent","--node-label","harvesterhci.io/managed=true","--node-label","rke.cattle.io/machine=5c82e559-5ff1-4600-97ea-098403f53097","--private-registry","/etc/rancher/rke2/registries.yaml","--server","https://10.115.50.55:9345","--token","********","--kubelet-arg","max-pods=200","--kubelet-arg","system-reserved=cpu=408m","--kubelet-arg","kube-reserved=cpu=612m","--apiVersion","audit.k8s.io/v1","--kind","Policy","--omitStages","ResponseStarted","--omitStages","ResponseComplete","--rules","[{level
      Metadata} {verbs [create delete patch]} {omitStages [ResponseStarted ResponseComplete]}]"]'
    rke2.io/node-config-hash: PGNG6Z6KEZDRQGZ4UI7235YNQNHNPFBFFATWJ4E3OET3K7EREC7A====
    rke2.io/node-env: '{"RKE2_DATA_DIR":"/var/lib/rancher/rke2"}'
    volumes.kubernetes.io/controller-managed-attach-detach: "true"
  creationTimestamp: "2025-08-08T04:01:50Z"
  finalizers:
  - wrangler.cattle.io/managed-etcd-controller
  - wrangler.cattle.io/node
  - wrangler.cattle.io/maintain-node-controller
  - wrangler.cattle.io/node-remove-controller
  - wrangler.cattle.io/harvester-network-manager-node-controller
  labels:
    beta.kubernetes.io/arch: amd64
    beta.kubernetes.io/instance-type: rke2
    beta.kubernetes.io/os: linux
    cpu-feature.node.kubevirt.io/3dnowprefetch: "true"
  name: epyc-milan
  resourceVersion: "19536311"
  uid: e7ded6e8-088a-4222-bfb0-7bb92a5b7d70
spec:
  podCIDR: 10.52.1.0/24
  podCIDRs:
  - 10.52.1.0/24
  providerID: rke2://epyc-milan
status:
  addresses:
  - address: 10.115.50.51
    type: InternalIP
  - address: epyc-milan
    type: Hostname
  allocatable:
    cpu: 14980m
    devices.kubevirt.io/kvm: 1k
    devices.kubevirt.io/tun: 1k
    devices.kubevirt.io/vhost-net: 1k
    ephemeral-storage: "149527126718"
    hugepages-1Gi: "0"
    hugepages-2Mi: "0"
    memory: 65837932Ki
    pods: "200"
  capacity:
    cpu: "16"
    devices.kubevirt.io/kvm: 1k
    devices.kubevirt.io/tun: 1k
    devices.kubevirt.io/vhost-net: 1k
    ephemeral-storage: 153707984Ki
    hugepages-1Gi: "0"
    hugepages-2Mi: "0"
    memory: 65837932Ki
    pods: "200"
  conditions:
  - lastHeartbeatTime: "2025-08-08T05:05:02Z"
    lastTransitionTime: "2025-08-08T05:05:02Z"
    message: Flannel is running on this node
    reason: FlannelIsUp
    status: "False"
    type: NetworkUnavailable
  - lastHeartbeatTime: "2025-08-08T05:05:02Z"
    lastTransitionTime: "2025-08-08T05:05:02Z"
    message: error placing node in maintainance mode due to an actual error string
    reason: Error
    status: True
    type: MaintaianceModeOp
  - lastHeartbeatTime: "2025-08-29T03:32:39Z"
    lastTransitionTime: "2025-08-08T05:04:56Z"
    message: kubelet has sufficient memory available
    reason: KubeletHasSufficientMemory
    status: "False"
    type: MemoryPressure
  - lastHeartbeatTime: "2025-08-29T03:32:39Z"
    lastTransitionTime: "2025-08-08T05:04:56Z"
    message: kubelet has no disk pressure
    reason: KubeletHasNoDiskPressure
    status: "False"
    type: DiskPressure
  - lastHeartbeatTime: "2025-08-29T03:32:39Z"
    lastTransitionTime: "2025-08-08T05:04:56Z"
    message: kubelet has sufficient PID available
    reason: KubeletHasSufficientPID
    status: "False"
    type: PIDPressure
  - lastHeartbeatTime: "2025-08-29T03:32:39Z"
    lastTransitionTime: "2025-08-08T05:04:56Z"
    message: kubelet is posting ready status
    reason: KubeletReady
    status: "True"
    type: Ready
  daemonEndpoints:
    kubeletEndpoint:
      Port: 10250
  images:
  - names:
    - docker.io/rancher/rancher:v2.10.1
    sizeBytes: 2424463043
  - names:
    - docker.io/rancher/rancher-agent:v2.10.1
    sizeBytes: 1394986895
  - names:
    - docker.io/longhornio/longhorn-instance-manager:v1.7.3
    sizeBytes: 1016360728
  - names:
    - docker.io/rancher/nginx-ingress-controller:v1.12.1-hardened1
    sizeBytes: 761840672
  - names:
    - docker.io/rancher/hardened-kubernetes:v1.31.7-rke2r1-build20250312
    sizeBytes: 719658570
  - names:
    - registry.suse.com/suse/sles/15.6/libguestfs-tools:1.3.1-150600.5.9.1
    sizeBytes: 706322581
  - names:
    - docker.io/rancher/hardened-calico:v3.29.2-build20250306
    sizeBytes: 595156302
  - names: null
    sizeBytes: 525670380
  - names:
    - docker.io/rancher/harvester-upgrade:v1.4.3
    sizeBytes: 475809485
  - names:
    - docker.io/longhornio/longhorn-engine:v1.7.3
    sizeBytes: 401028329
  - names:
    - docker.io/rancher/shell:v0.1.24
    sizeBytes: 399338553
  - names:
    - docker.io/longhornio/backing-image-manager:v1.7.3
    sizeBytes: 389532868
  - names:
    - docker.io/rancher/harvester:v1.4.3
    sizeBytes: 332801554
  - names:
    - docker.io/longhornio/longhorn-manager:v1.7.3
    sizeBytes: 330174497
  - names:
    - docker.io/rancher/shell:v0.3.0
    sizeBytes: 312839214
  - names:
    - docker.io/rancher/fleet:v0.11.2
    sizeBytes: 312721109
  - names:
    - docker.io/rancher/mirrored-grafana-grafana:9.1.5
    sizeBytes: 305920285
  - names:
    - docker.io/rancher/shell:v0.1.26
    sizeBytes: 303834158
  - names:
    - docker.io/rancher/rke2-runtime:v1.31.7-rke2r1
    sizeBytes: 287513428
  - names:
    - docker.io/longhornio/support-bundle-kit:v0.0.51
    sizeBytes: 282763790
  - names:
    - docker.io/rancher/support-bundle-kit:v0.0.46
    sizeBytes: 281580819
  - names:
    - docker.io/rancher/harvester-vm-import-controller:v0.4.3
    sizeBytes: 280359912
  - names:
    - docker.io/rancher/harvester-pcidevices:v0.4.3
    sizeBytes: 259500321
  - names:
    - ghcr.io/kube-logging/fluentd:v1.15-ruby3
    sizeBytes: 252464715
  - names:
    - registry.suse.com/suse/sles/15.6/virt-handler:1.3.1-150600.5.9.1
    sizeBytes: 250388458
  - names:
    - docker.io/longhornio/longhorn-cli:v1.7.3
    sizeBytes: 248808574
  - names:
    - docker.io/rancher/harvester-webhook:v1.4.3
    sizeBytes: 239057209
  - names:
    - docker.io/rancher/mirrored-prometheus-prometheus:v2.45.0
    sizeBytes: 237682041
  - names:
    - docker.io/longhornio/longhorn-share-manager:v1.7.3
    sizeBytes: 232535508
  - names:
    - docker.io/rancher/gitjob:v0.9.8
    sizeBytes: 229952146
  - names:
    - docker.io/rancher/hardened-flannel:v0.26.5-build20250306
    sizeBytes: 225547041
  - names:
    - docker.io/rancher/mirrored-banzaicloud-fluentd:v1.14.6-alpine-5
    sizeBytes: 221240002
  - names:
    - docker.io/rancher/harvester-seeder:v0.4.3
    sizeBytes: 216884440
  - names:
    - docker.io/longhornio/longhorn-ui:v1.7.3
    sizeBytes: 212394241
  - names:
    - docker.io/rancher/harvester-network-controller:v0.5.6
    sizeBytes: 210608741
  - names:
    - docker.io/rancher/harvester-networkfs-manager:v0.1.2
    sizeBytes: 205283920
  - names:
    - registry.suse.com/suse/sles/15.6/virt-operator:1.3.1-150600.5.9.1
    sizeBytes: 202636514
  - names:
    - docker.io/rancher/harvester-node-disk-manager:v0.7.11
    sizeBytes: 202586710
  - names:
    - registry.suse.com/suse/sles/15.6/virt-api:1.3.1-150600.5.9.1
    sizeBytes: 201247330
  - names:
    - registry.suse.com/suse/sles/15.6/virt-controller:1.3.1-150600.5.9.1
    sizeBytes: 196248875
  - names:
    - docker.io/rancher/system-agent:v0.3.11-suc
    sizeBytes: 192713957
  - names:
    - docker.io/rancher/harvester-cluster-repo:v1.4.3
    sizeBytes: 190839535
  - names:
    - docker.io/rancher/klipper-helm:v0.9.4-build20250113
    sizeBytes: 190360365
  - names:
    - registry.suse.com/suse/sles/15.6/virt-launcher@sha256:a16fe8d5a19fa9191b236808cfd9abb10e43a4b58d720579973ff94aa5d1b831
    - registry.suse.com/suse/sles/15.6/virt-launcher:1.4.1-150600.5.21.2
    sizeBytes: 187677102
  - names:
    - registry.suse.com/suse/sles/15.7/virt-launcher@sha256:75bc5cd42d28f01e567c53e6f0f1386080902dd15037c2ee5f9febea646a3f87
    - registry.suse.com/suse/sles/15.7/virt-launcher:1.5.2-150700.3.5.2
    sizeBytes: 185646514
  - names:
    - docker.io/rancher/harvester-node-manager-webhook:v0.3.4
    sizeBytes: 176110656
  - names:
    - docker.io/rancher/harvester-node-manager:v0.3.4
    sizeBytes: 172317256
  - names:
    - docker.io/rancher/harvester-node-disk-manager-webhook:v0.7.11
    sizeBytes: 170517065
  - names:
    - quay.io/kubevirt/virt-launcher:v1.3.1
    - registry.suse.com/suse/sles/15.6/virt-launcher:1.3.1-150600.5.9.1
    sizeBytes: 164087662
  - names:
    - registry.suse.com/bci/bci-base:15.6
    sizeBytes: 125906507
  nodeInfo:
    architecture: amd64
    bootID: b4fa04ae-71c6-43d0-adf3-f1793042df8a
    containerRuntimeVersion: containerd://2.0.4-k3s2
    kernelVersion: 5.14.21-150500.55.103-default
    kubeProxyVersion: v1.31.7+rke2r1
    kubeletVersion: v1.31.7+rke2r1
    machineID: e26b1f43cecf3311d1f2cb5d682e47c6
    operatingSystem: linux
    osImage: Harvester v1.4.3
    systemUUID: 833d1356-08a1-57f6-aa0c-3053cd4319a0
  runtimeHandlers:
  - features:
      recursiveReadOnlyMounts: true
      userNamespaces: true
    name: runc
  - features:
      recursiveReadOnlyMounts: true
      userNamespaces: true
    name: ""
  - features:
      recursiveReadOnlyMounts: false
      userNamespaces: false
    name: runhcs-wcow-process
  volumesAttached:
  - devicePath: ""
    name: kubernetes.io/csi/driver.longhorn.io^pvc-6cfe681f-9a4f-4293-b1f4-853b26bd9848
  volumesInUse:
  - kubernetes.io/csi/driver.longhorn.io^pvc-6cfe681f-9a4f-4293-b1f4-853b26bd9848
`)

var nodeFailureYaml = []byte(`apiVersion: v1
kind: Node
metadata:
  annotations:
    alpha.kubernetes.io/provided-node-ip: 10.115.50.51
    cluster.x-k8s.io/cluster-name: local
    cluster.x-k8s.io/cluster-namespace: fleet-local
    cluster.x-k8s.io/labels-from-machine: ""
    cluster.x-k8s.io/machine: custom-d8ea5ba6e941
    csi.volume.kubernetes.io/nodeid: '{"driver.longhorn.io":"epyc-milan"}'
    flannel.alpha.coreos.com/backend-data: '{"VNI":1,"VtepMAC":"0e:4d:1b:b8:34:c8"}'
    flannel.alpha.coreos.com/backend-type: vxlan
    flannel.alpha.coreos.com/kube-subnet-manager: "true"
    flannel.alpha.coreos.com/public-ip: 10.115.50.51
    kubevirt.io/heartbeat: "2025-08-29T03:56:55Z"
    kubevirt.io/ksm-handler-managed: "false"
    management.cattle.io/pod-limits: '{"cpu":"4100m","memory":"4948Mi"}'
    management.cattle.io/pod-requests: '{"cpu":"3393m","memory":"2135Mi","pods":"24"}'
    node.alpha.kubernetes.io/ttl: "0"
    node.harvesterhci.io/ntp-service: '{"ntpSyncStatus":"synced","currentNtpServers":"0.suse.pool.ntp.org"}'
    rke2.io/hostname: epyc-milan
    rke2.io/internal-ip: 10.115.50.51
    rke2.io/node-args: '["agent","--node-label","harvesterhci.io/managed=true","--node-label","rke.cattle.io/machine=5c82e559-5ff1-4600-97ea-098403f53097","--private-registry","/etc/rancher/rke2/registries.yaml","--server","https://10.115.50.55:9345","--token","********","--kubelet-arg","max-pods=200","--kubelet-arg","system-reserved=cpu=408m","--kubelet-arg","kube-reserved=cpu=612m","--apiVersion","audit.k8s.io/v1","--kind","Policy","--omitStages","ResponseStarted","--omitStages","ResponseComplete","--rules","[{level
      Metadata} {verbs [create delete patch]} {omitStages [ResponseStarted ResponseComplete]}]"]'
    rke2.io/node-config-hash: PGNG6Z6KEZDRQGZ4UI7235YNQNHNPFBFFATWJ4E3OET3K7EREC7A====
    rke2.io/node-env: '{"RKE2_DATA_DIR":"/var/lib/rancher/rke2"}'
    volumes.kubernetes.io/controller-managed-attach-detach: "true"
  creationTimestamp: "2025-08-08T04:01:50Z"
  finalizers:
  - wrangler.cattle.io/managed-etcd-controller
  - wrangler.cattle.io/node
  - wrangler.cattle.io/maintain-node-controller
  - wrangler.cattle.io/node-remove-controller
  - wrangler.cattle.io/harvester-network-manager-node-controller
  labels:
    beta.kubernetes.io/arch: amd64
    beta.kubernetes.io/instance-type: rke2
    beta.kubernetes.io/os: linux
    cpu-feature.node.kubevirt.io/3dnowprefetch: "true"
    cpu-feature.node.kubevirt.io/abm: "true"
    cpu-feature.node.kubevirt.io/adx: "true"
    cpu-feature.node.kubevirt.io/aes: "true"
  name: epyc-milan
  resourceVersion: "19552935"
  uid: e7ded6e8-088a-4222-bfb0-7bb92a5b7d70
spec:
  podCIDR: 10.52.1.0/24
  podCIDRs:
  - 10.52.1.0/24
  providerID: rke2://epyc-milan
  taints:
  - effect: NoSchedule
    key: node.kubernetes.io/unreachable
    timeAdded: "2025-08-29T03:56:39Z"
  - effect: NoExecute
    key: node.kubernetes.io/unreachable
    timeAdded: "2025-08-29T03:56:45Z"
status:
  addresses:
  - address: 10.115.50.51
    type: InternalIP
  - address: epyc-milan
    type: Hostname
  allocatable:
    cpu: 14980m
    devices.kubevirt.io/kvm: 1k
    devices.kubevirt.io/tun: 1k
    devices.kubevirt.io/vhost-net: 1k
    ephemeral-storage: "149527126718"
    hugepages-1Gi: "0"
    hugepages-2Mi: "0"
    memory: 65837932Ki
    pods: "200"
  capacity:
    cpu: "16"
    devices.kubevirt.io/kvm: 1k
    devices.kubevirt.io/tun: 1k
    devices.kubevirt.io/vhost-net: 1k
    ephemeral-storage: 153707984Ki
    hugepages-1Gi: "0"
    hugepages-2Mi: "0"
    memory: 65837932Ki
    pods: "200"
  conditions:
  - lastHeartbeatTime: "2025-08-08T05:05:02Z"
    lastTransitionTime: "2025-08-08T05:05:02Z"
    message: Flannel is running on this node
    reason: FlannelIsUp
    status: "False"
    type: NetworkUnavailable
  - lastHeartbeatTime: "2025-08-08T05:05:02Z"
    lastTransitionTime: "2025-08-08T05:05:02Z"
    message: ""
    reason: HarvesterNodeFailure
    status: "true"
    type: Drained
  daemonEndpoints:
    kubeletEndpoint:
      Port: 10250
  images:
  - names:
    - docker.io/rancher/rancher:v2.10.1
    sizeBytes: 2424463043
  - names:
    - docker.io/rancher/rancher-agent:v2.10.1
    sizeBytes: 1394986895
  - names:
    - docker.io/longhornio/longhorn-instance-manager:v1.7.3
    sizeBytes: 1016360728
  - names:
    - docker.io/rancher/nginx-ingress-controller:v1.12.1-hardened1
    sizeBytes: 761840672
  - names:
    - docker.io/rancher/hardened-kubernetes:v1.31.7-rke2r1-build20250312
    sizeBytes: 719658570
  - names:
    - registry.suse.com/suse/sles/15.6/libguestfs-tools:1.3.1-150600.5.9.1
    sizeBytes: 706322581
  - names:
    - docker.io/rancher/hardened-calico:v3.29.2-build20250306
    sizeBytes: 595156302
  - names: null
    sizeBytes: 525670380
  - names:
    - docker.io/rancher/harvester-upgrade:v1.4.3
    sizeBytes: 475809485
  - names:
    - docker.io/longhornio/longhorn-engine:v1.7.3
    sizeBytes: 401028329
  - names:
    - docker.io/rancher/shell:v0.1.24
    sizeBytes: 399338553
  - names:
    - docker.io/longhornio/backing-image-manager:v1.7.3
    sizeBytes: 389532868
  - names:
    - docker.io/rancher/harvester:v1.4.3
    sizeBytes: 332801554
  - names:
    - docker.io/longhornio/longhorn-manager:v1.7.3
    sizeBytes: 330174497
  - names:
    - docker.io/rancher/shell:v0.3.0
    sizeBytes: 312839214
  - names:
    - docker.io/rancher/fleet:v0.11.2
    sizeBytes: 312721109
  - names:
    - docker.io/rancher/mirrored-grafana-grafana:9.1.5
    sizeBytes: 305920285
  - names:
    - docker.io/rancher/shell:v0.1.26
    sizeBytes: 303834158
  - names:
    - docker.io/rancher/rke2-runtime:v1.31.7-rke2r1
    sizeBytes: 287513428
  - names:
    - docker.io/longhornio/support-bundle-kit:v0.0.51
    sizeBytes: 282763790
  - names:
    - docker.io/rancher/support-bundle-kit:v0.0.46
    sizeBytes: 281580819
  - names:
    - docker.io/rancher/harvester-vm-import-controller:v0.4.3
    sizeBytes: 280359912
  - names:
    - docker.io/rancher/harvester-pcidevices:v0.4.3
    sizeBytes: 259500321
  - names:
    - ghcr.io/kube-logging/fluentd:v1.15-ruby3
    sizeBytes: 252464715
  - names:
    - registry.suse.com/suse/sles/15.6/virt-handler:1.3.1-150600.5.9.1
    sizeBytes: 250388458
  - names:
    - docker.io/longhornio/longhorn-cli:v1.7.3
    sizeBytes: 248808574
  - names:
    - docker.io/rancher/harvester-webhook:v1.4.3
    sizeBytes: 239057209
  - names:
    - docker.io/rancher/mirrored-prometheus-prometheus:v2.45.0
    sizeBytes: 237682041
  - names:
    - docker.io/longhornio/longhorn-share-manager:v1.7.3
    sizeBytes: 232535508
  - names:
    - docker.io/rancher/gitjob:v0.9.8
    sizeBytes: 229952146
  - names:
    - docker.io/rancher/hardened-flannel:v0.26.5-build20250306
    sizeBytes: 225547041
  - names:
    - docker.io/rancher/mirrored-banzaicloud-fluentd:v1.14.6-alpine-5
    sizeBytes: 221240002
  - names:
    - docker.io/rancher/harvester-seeder:v0.4.3
    sizeBytes: 216884440
  - names:
    - docker.io/longhornio/longhorn-ui:v1.7.3
    sizeBytes: 212394241
  - names:
    - docker.io/rancher/harvester-network-controller:v0.5.6
    sizeBytes: 210608741
  - names:
    - docker.io/rancher/harvester-networkfs-manager:v0.1.2
    sizeBytes: 205283920
  - names:
    - registry.suse.com/suse/sles/15.6/virt-operator:1.3.1-150600.5.9.1
    sizeBytes: 202636514
  - names:
    - docker.io/rancher/harvester-node-disk-manager:v0.7.11
    sizeBytes: 202586710
  - names:
    - registry.suse.com/suse/sles/15.6/virt-api:1.3.1-150600.5.9.1
    sizeBytes: 201247330
  - names:
    - registry.suse.com/suse/sles/15.6/virt-controller:1.3.1-150600.5.9.1
    sizeBytes: 196248875
  - names:
    - docker.io/rancher/system-agent:v0.3.11-suc
    sizeBytes: 192713957
  - names:
    - docker.io/rancher/harvester-cluster-repo:v1.4.3
    sizeBytes: 190839535
  - names:
    - docker.io/rancher/klipper-helm:v0.9.4-build20250113
    sizeBytes: 190360365
  - names:
    - registry.suse.com/suse/sles/15.6/virt-launcher@sha256:a16fe8d5a19fa9191b236808cfd9abb10e43a4b58d720579973ff94aa5d1b831
    - registry.suse.com/suse/sles/15.6/virt-launcher:1.4.1-150600.5.21.2
    sizeBytes: 187677102
  - names:
    - registry.suse.com/suse/sles/15.7/virt-launcher@sha256:75bc5cd42d28f01e567c53e6f0f1386080902dd15037c2ee5f9febea646a3f87
    - registry.suse.com/suse/sles/15.7/virt-launcher:1.5.2-150700.3.5.2
    sizeBytes: 185646514
  - names:
    - docker.io/rancher/harvester-node-manager-webhook:v0.3.4
    sizeBytes: 176110656
  - names:
    - docker.io/rancher/harvester-node-manager:v0.3.4
    sizeBytes: 172317256
  - names:
    - docker.io/rancher/harvester-node-disk-manager-webhook:v0.7.11
    sizeBytes: 170517065
  - names:
    - quay.io/kubevirt/virt-launcher:v1.3.1
    - registry.suse.com/suse/sles/15.6/virt-launcher:1.3.1-150600.5.9.1
    sizeBytes: 164087662
  - names:
    - registry.suse.com/bci/bci-base:15.6
    sizeBytes: 125906507
  nodeInfo:
    architecture: amd64
    bootID: b4fa04ae-71c6-43d0-adf3-f1793042df8a
    containerRuntimeVersion: containerd://2.0.4-k3s2
    kernelVersion: 5.14.21-150500.55.103-default
    kubeProxyVersion: v1.31.7+rke2r1
    kubeletVersion: v1.31.7+rke2r1
    machineID: e26b1f43cecf3311d1f2cb5d682e47c6
    operatingSystem: linux
    osImage: Harvester v1.4.3
    systemUUID: 833d1356-08a1-57f6-aa0c-3053cd4319a0
  runtimeHandlers:
  - features:
      recursiveReadOnlyMounts: true
      userNamespaces: true
    name: runc
  - features:
      recursiveReadOnlyMounts: true
      userNamespaces: true
    name: ""
  - features:
      recursiveReadOnlyMounts: false
      userNamespaces: false
    name: runhcs-wcow-process`)

func Test_Summary(t *testing.T) {
	nodeObj := &corev1.Node{}
	assert := require.New(t)
	err := yaml.Unmarshal(nodeYaml, nodeObj)
	assert.NoError(err)
	summary := Summarize(nodeObj)
	t.Log(summary)
}

func Test_FailureSummary(t *testing.T) {
	nodeObj := &corev1.Node{}
	assert := require.New(t)
	err := yaml.Unmarshal(nodeFailureYaml, nodeObj)
	assert.NoError(err)
	summary := Summarize(nodeObj)
	t.Log(summary)
}
