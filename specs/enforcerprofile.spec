# Model
model:
  rest_name: enforcerprofile
  resource_name: enforcerprofiles
  entity_name: EnforcerProfile
  package: squall
  description: |-
    Allows to create reusable configuration profile for your enforcers. Enforcer
    Profiles contains various startup information that can (for some) be updated
    live. Enforcer Profiles are assigned to some Enforcer using a Enforcer Profile
    Mapping Policy.
  aliases:
  - profile
  - profiles
  indexes:
  - - :shard
    - zone
    - zHash
  - - namespace
  - - namespace
    - name
  - - namespace
    - normalizedTags
  get:
    description: Retrieves the object with the given ID.
  update:
    description: Updates the object with the given ID.
  delete:
    description: Deletes the object with the given ID.
    global_parameters:
    - $filtering
  extends:
  - '@base'
  - '@described'
  - '@identifiable-pk-stored'
  - '@named'
  - '@metadatable'
  - '@zonable'
  validations:
  - $enforcerprofile

# Attributes
attributes:
  v1:
  - name: IPTablesMarkValue
    description: IptablesMarkValue is the mark value to be used in an iptables implementation.
    type: integer
    exposed: true
    stored: true
    default_value: 1000
    max_value: 65000
    orderable: true

  - name: PUBookkeepingInterval
    description: PUBookkeepingInterval configures how often the PU will be synchronized.
    type: string
    exposed: true
    stored: true
    allowed_chars: ^[0-9]+[smh]$
    allowed_chars_message: must be a valid duration like <n>s or <n>s or <n>h
    default_value: 15m
    orderable: true

  - name: PUHeartbeatInterval
    description: PUHeartbeatInterval configures the heart beat interval.
    type: string
    exposed: true
    stored: true
    allowed_chars: ^[0-9]+[smh]$
    allowed_chars_message: must be a valid duration like <n>s or <n>s or <n>h
    default_value: 5s
    orderable: true

  - name: applicationProxyPort
    description: Port used by aporeto application proxy.
    type: integer
    exposed: true
    stored: true
    default_value: 20992
    max_value: 65535
    min_value: 1
    orderable: true

  - name: auditProfileSelectors
    description: |-
      AuditProfileSelectors is the list of tags (key/value pairs) that define the
      audit policies that must be implemented by this enforcer. The enforcer will
      implement all policies that match any of these tags.
    type: list
    exposed: true
    subtype: string
    stored: true

  - name: auditProfiles
    description: |-
      AuditProfiles returns the audit rules associated with the enforcer profile. This
      is a read only attribute when an enforcer profile is resolved for an enforcer.
    type: refList
    exposed: true
    subtype: auditprofile
    read_only: true
    autogenerated: true

  - name: auditSocketBufferSize
    description: AuditSocketBufferSize is the size of the audit socket buffer. Default
      16384.
    type: integer
    exposed: true
    stored: true
    default_value: 16384
    max_value: 262144
    orderable: true

  - name: dockerSocketAddress
    description: DockerSocketAddress is the address of the docker daemon.
    type: string
    exposed: true
    stored: true
    allowed_chars: ^(:([1-9]|[1-9][0-9]|[1-9][0-9]{1,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|65535))$|(unix://(/[^/]{1,16}){1,5}/?)$
    allowed_chars_message: must be a valid url or path starting by unix://
    default_value: unix:///var/run/docker.sock
    orderable: true

  - name: excludedInterfaces
    description: ExcludedInterfaces is a list of interfaces that must be excluded.
    type: list
    exposed: true
    subtype: string
    stored: true
    orderable: true

  - name: excludedNetworks
    description: |-
      ExcludedNetworks is the list of networks that must be excluded for this
      enforcer.
    type: list
    exposed: true
    subtype: string
    stored: true
    orderable: true

  - name: hostModeEnabled
    description: |-
      hostModeEnabled enables protection of the complete host. When this option is
      turned on, all incoming and outgoing flows will be monitored. Flows will
      be allowed if and only if a network policy has been created to allow the flow.
    type: boolean
    exposed: true
    stored: true
    default_value: false
    orderable: true

  - name: hostServices
    description: |-
      HostServices is a list of services that must be activated by default to all
      enforcers matching this profile.
    type: refList
    exposed: true
    subtype: hostservice
    stored: true
    validations:
    - $host_services_list
    extensions:
      refMode: pointer

  - name: ignoreExpression
    description: |-
      IgnoreExpression allows to set a tag expression that will make Aporeto to ignore
      docker container started with labels matching the rule.
    type: external
    exposed: true
    subtype: list_of_lists_of_strings
    stored: true

  - name: killContainersOnFailure
    description: |-
      KillContainersOnFailure will configure the enforcers to kill any containers if
      there are policy failures.
    type: boolean
    exposed: true
    stored: true

  - name: kubernetesMetadataExtractor
    description: |-
      Select which metadata extractor to use to process new processing units from
      Kubernetes.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - KubeSquall
    - PodAtomic
    - PodContainers
    default_value: PodAtomic
    orderable: true

  - name: kubernetesSupportEnabled
    description: KubernetesSupportEnabled enables kubernetes mode for the enforcer.
    type: boolean
    exposed: true
    stored: true
    default_value: false
    orderable: true

  - name: linuxProcessesSupportEnabled
    description: LinuxProcessesSupportEnabled configures support for Linux processes.
    type: boolean
    exposed: true
    stored: true
    default_value: true

  - name: metadataExtractor
    description: Select which metadata extractor to use to process new processing
      units.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - Docker
    - ECS
    - Kubernetes
    default_value: Docker
    orderable: true

  - name: policySynchronizationInterval
    description: |-
      PolicySynchronizationInterval configures how often the policy will be
      resynchronized.
    type: string
    exposed: true
    stored: true
    allowed_chars: ^[0-9]+[smh]$
    allowed_chars_message: must be a valid duration like <n>s or <n>s or <n>h
    default_value: 10m
    orderable: true

  - name: proxyListenAddress
    description: |-
      ProxyListenAddress is the address the enforcer should use to listen for API
      calls. It can be a port (example :9443) or socket path
      example:
        unix:///var/run/aporeto.sock.
    type: string
    exposed: true
    stored: true
    allowed_chars: ^(:([1-9]|[1-9][0-9]|[1-9][0-9]{1,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|65535))$|(unix://(/[^/]{1,16}){1,5}/?)$
    allowed_chars_message: must be a valid url or path starting by unix://
    default_value: unix:///var/run/aporeto.sock
    orderable: true

  - name: receiverNumberOfQueues
    description: |-
      ReceiverNumberOfQueues is the number of queues for the NFQUEUE of the network
      receiver starting at the ReceiverQueue.
    type: integer
    exposed: true
    stored: true
    default_value: 4
    max_value: 16
    min_value: 1
    orderable: true

  - name: receiverQueue
    description: ReceiverQueue is the base queue number for traffic from the network.
    type: integer
    exposed: true
    stored: true
    max_value: 1000
    orderable: true

  - name: receiverQueueSize
    description: ReceiverQueueSize is the queue size of the receiver.
    type: integer
    exposed: true
    stored: true
    default_value: 500
    max_value: 5000
    min_value: 1
    orderable: true

  - name: remoteEnforcerEnabled
    description: |-
      RemoteEnforcerEnabled inidicates whether a single enforcer should be used or a
      distributed enforcer. True means distributed.
    type: boolean
    exposed: true
    stored: true
    default_value: true
    orderable: true

  - name: targetNetworks
    description: TargetNetworks is the list of networks that authorization should
      be applied.
    type: list
    exposed: true
    subtype: string
    stored: true
    orderable: true

  - name: targetUDPNetworks
    description: |-
      TargetUDPNetworks is the list of UDP networks that authorization should be
      applied.
    type: list
    exposed: true
    subtype: string
    stored: true
    orderable: true

  - name: transmitterNumberOfQueues
    description: TransmitterNumberOfQueues is the number of queues for application
      traffic.
    type: integer
    exposed: true
    stored: true
    default_value: 4
    max_value: 16
    min_value: 1
    orderable: true

  - name: transmitterQueue
    description: |-
      TransmitterQueue is the queue number for traffic from the applications starting
      at the transmitterQueue.
    type: integer
    exposed: true
    stored: true
    default_value: 4
    max_value: 1000
    min_value: 1
    orderable: true

  - name: transmitterQueueSize
    description: TransmitterQueueSize is the size of the queue for application traffic.
    type: integer
    exposed: true
    stored: true
    default_value: 500
    max_value: 1000
    min_value: 1
    orderable: true

  - name: trustedCAs
    description: List of trusted CA. If empty the main chain of trust will be used.
    type: list
    exposed: true
    subtype: string
    stored: true

# Relations
relations:
- rest_name: auditprofile
  get:
    description: Returns the list of AuditProfiles used by an enforcer profile.
