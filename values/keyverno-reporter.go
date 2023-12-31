package values

const Keyverno = `
image:
  registry: ghcr.io
  repository: kyverno/policy-reporter
  pullPolicy: IfNotPresent
  tag: 2.5.0

imagePullSecrets: []

# Deploy not more than one replica
# Policy Reporter doesn't scale yet.
# Each pod will report each change.
replicaCount: 1

deploymentStrategy: {}
  # rollingUpdate:
  #  maxSurge: 25%
  #  maxUnavailable: 25%
# type: RollingUpdate

# Key/value pairs that are attached to Deployment.
annotations: {}

# Create cluster role policies
rbac:
  enabled: true

serviceAccount:
  # Specifies whether a service account should be created
  create: true
  # Annotations to add to the service account
  annotations: {}
  # The name of the service account to use.
  # If not set and create is true, a name is generated using the fullname template
  name: ""

service:
  enabled: true
  ## configuration of service
  # key/value
  annotations: {}
  # key/value
  labels: {}
  type: ClusterIP
  # integer number. This is port for service
  port: 8080

podSecurityContext:
  fsGroup: 1234

securityContext:
  runAsUser: 1234
  runAsNonRoot: true
  privileged: false
  allowPrivilegeEscalation: false
  readOnlyRootFilesystem: true
  capabilities:
    drop:
      - ALL
  seccompProfile:
    type: RuntimeDefault

# Key/value pairs that are attached to pods.
podAnnotations: {}

# Key/value pairs that are attached to pods.
podLabels: {}

resources: {}
  # We usually recommend not to specify default resources and to leave this as a conscious
  # choice for the user. This also increases chances charts run on environments with little
  # resources, such as Minikube. If you do want to specify resources, uncomment the following
  # lines, adjust them as necessary, and remove the curly braces after 'resources:'.
  # limits:
  #   memory: 100Mi
  #   cpu: 10m
  # requests:
  #   memory: 75Mi
#   cpu: 5m

# Enable a NetworkPolicy for this chart. Useful on clusters where Network Policies are
# used and configured in a default-deny fashion.
networkPolicy:
  enabled: false
  # Kubernetes API Server
  egress:
    - to:
      ports:
        - protocol: TCP
          port: 6443
  ingress: []

# REST API
rest:
  enabled: false

# Prometheus Metrics API
metrics:
  enabled: false

# Filter PolicyReport resources to process
reportFilter:
  namespaces:
    # Process only PolicyReport resources from an included namespace, wildcards are supported
    include: []
    # Ignore all PolicyReport resources from a excluded namespace, wildcards are supported
    # exclude will be ignored if an include filter exists
    exclude: []
  clusterReports:
    # Disable the processing of ClusterPolicyReports
    disabled: false

# enable policy-report-ui
ui:
  enabled: true

kyvernoPlugin:
  enabled: true

# Settings for the monitoring subchart
monitoring:
  enabled: false

global:
  # available plugins
  plugins:
    # enable kyverno for Policy Reporter UI and monitoring
    kyverno: true
  # The name of service policy-report. Defaults to ReleaseName.
  backend: ""
  # Service Port number
  port: 8080
  fullnameOverride: ""
  # additional labels added on each resource
  labels: {}

# configure mappings from policy to priority
# you can use default to configure a default priority for fail results
# example mapping
#   default: warning
#   require-ns-labels: error
policyPriorities: {}

# Reference a configuration which already exists instead of creating one
existingTargetConfig:
  enabled: false
  # Name of the secret with the config
  name: ""
  # subPath within the secret (defaults to config.yaml)
  subPath: ""

# Supported targets for new PolicyReport Results
target:
  loki:
    # loki host address
    host: ""
    # minimum priority "" < info < warning < critical < error
    minimumPriority: ""
    # list of sources which should send to loki
    sources: []
    # Skip already existing PolicyReportResults on startup
    skipExistingOnStartup: true
    # Added as additional labels to each Loki event
    customLabels: {}
    # Filter Results which should send to this target by namespaces, priorities or policies
    # Wildcars for namespaces and policies are supported, you can either define exclude or include values
    # Filters are available for all targets except the UI
    filter: {}
    #      namespaces:
    #        include: ["develop"]
    #      priorities:
    #        exclude: ["debug", "info", "error"]
    channels: []
  #    - host: "http://loki.loki-stack:3100"
  #      sources: []
  #      customLabels: {}
  #      filter:
  #        namespaces:
  #          include: ["develop"]
  #        priorities:
  #          exclude: ["debug", "info", "error"]

  elasticsearch:
    # elasticsearch host address
    host: ""
    # elasticsearch index (default: policy-reporter)
    index: ""
    # elasticsearch index rotation and index suffix
    # possible values: daily, monthly, annually, none (default: daily)
    rotation: ""
    # minimum priority "" < info < warning < critical < error
    minimumPriority: ""
    # list of sources which should send to elasticsearch
    sources: []
    # Skip already existing PolicyReportResults on startup
    skipExistingOnStartup: true
    # filter results send by namespaces, policies and priorities
    filter: {}
    # add additional elasticsearch channels with different configurations and filters
    channels: []

  slack:
    # slack app webhook address
    webhook: ""
    # minimum priority "" < info < warning < critical < error
    minimumPriority: ""
    # list of sources which should send to slack
    sources: []
    # Skip already existing PolicyReportResults on startup
    skipExistingOnStartup: true
    # filter results send by namespaces, policies and priorities
    filter: {}
    # add additional slack channels with different configurations and filters
    channels: []
  #    - webhook: "https://slack.webhook1"
  #      filter:
  #        namespaces:
  #          include: ["develop"]
  #        priorities:
  #          exclude: ["debug", "info", "error"]
  #        policies:
  #          include: ["require-run-as-nonroot"]
  #    - webhook: "https://slack.webhook2"
  #      minimumPriority: "warning"
  #      filter:
  #        namespaces:
  #          include: ["team-a-*"]

  discord:
    # discord app webhook address
    webhook: ""
    # minimum priority "" < info < warning < critical < error
    minimumPriority: ""
    # list of sources which should send to discord
    sources: []
    # Skip already existing PolicyReportResults on startup
    skipExistingOnStartup: true
    # filter results send by namespaces, policies and priorities
    filter: {}
    # add additional discord channels with different configurations and filters
    channels: []

  teams:
    # teams webhook address
    webhook: ""
    # minimum priority "" < info < warning < critical < error
    minimumPriority: ""
    # list of sources which should send to teams
    sources: []
    # Skip already existing PolicyReportResults on startup
    skipExistingOnStartup: true
    # filter results send by namespaces, policies and priorities
    filter: {}
    # add additional teams channels with different configurations and filters
    channels: []

  ui:
    # ui host address
    host: ""
    # minimum priority "" < info < warning < critical < error
    minimumPriority: "warning"
    # list of sources which should send to the UI Log
    sources: []
    # Skip already existing PolicyReportResults on startup
    skipExistingOnStartup: true

  webhook:
    # webhook host address
    host: ""
    # additional http headers
    headers: {}
    # minimum priority "" < info < warning < critical < error
    minimumPriority: ""
    # list of sources which should send to the UI Log
    sources: []
    # Skip already existing PolicyReportResults on startup
    skipExistingOnStartup: true
    # filter results send by namespaces, policies and priorities
    filter: {}
    # add additional webhook channels with different configurations and filters
    channels: []

  s3:
    # S3 access key
    accessKeyID: ""
    # S3 secret access key
    secretAccessKey: ""
    # S3 storage region
    region: ""
    # S3 storage endpoint
    endpoint: ""
    # S3 storage, bucket name
    bucket: ""
    # name of prefix, keys will have format: s3://<bucket>/<prefix>/YYYY-MM-DD/YYYY-MM-DDTHH:mm:ss.s+01:00.json
    prefix: ""
    # minimum priority "" < info < warning < critical < error
    minimumPriority: ""
    # list of sources which should send to S3
    sources: []
    # Skip already existing PolicyReportResults on startup
    skipExistingOnStartup: true
    # filter results send by namespaces, policies and priorities
    filter: {}
    # add additional s3 channels with different configurations and filters
    channels: []

# Node labels for pod assignment
# ref: https://kubernetes.io/docs/user-guide/node-selection/
nodeSelector: {}

# Tolerations for pod assignment
# ref: https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/
tolerations: []

# Anti-affinity to disallow deploying client and master nodes on the same worker node
affinity: {}

# livenessProbe for policy-reporter
livenessProbe:
  httpGet:
    path: /ready
    port: http

# readinessProbe for policy-reporter
readinessProbe:
  httpGet:
    path: /healthz
    port: http
`
