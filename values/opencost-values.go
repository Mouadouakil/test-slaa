package values

const OpenCost = `
# -- Overwrite the default name of the chart
nameOverride: ""
# -- Overwrite all resources name created by the chart
fullnameOverride: ""

# -- List of secret names to use for pulling the images
imagePullSecrets: []

serviceAccount:
  # -- Specifies whether a service account should be created
  create: true
  # --  Annotations to add to the service account
  annotations: {}
  # eks.amazonaws.com/role-arn: arn:aws:iam::123456789012:role/eksctl-opencost
  # The name of the service account to use.
  # If not set and create is true, a name is generated using the fullname template
  name: ""
  # -- Whether pods running as this service account should have an API token automatically mounted
  automountServiceAccountToken: true

# --  Strategy to be used for the Deployment
updateStrategy:
  rollingUpdate:
    maxSurge: 1
    maxUnavailable: 1
  type: RollingUpdate
# --  Annotations to add to the all the resources
annotations: {}
# --  Annotations to add to the OpenCost Pod
podAnnotations: {}
# --  Annotations to add to the Secret
secretAnnotations: {}
# --  Labels to add to the OpenCost Pod
podLabels: {}
# --  Pod priority
priorityClassName: ~

# -- Holds pod-level security attributes and common container settings
podSecurityContext: {}
  # fsGroup: 2000

service:
  enabled: true
  # --  Annotations to add to the service
  annotations: {}
  # --  Labels to add to the service account
  labels: {}
  # --  Kubernetes Service type
  type: ClusterIP
  # -- extra ports.  Useful for sidecar pods such as oauth-proxy
  extraPorts: []
    # - name: oauth-proxy
    #   port: 8081
    #   targetPort: 8081
    # - name: oauth-metrics
    #   port: 8082
    #   targetPort: 8082

# Create cluster role policies
rbac:
  enabled: true

opencost:
  exporter:
    # -- The GCP Pricing API requires a key. This is supplied just for evaluation.
    cloudProviderApiKey: ""
    # -- Default cluster ID to use if cluster_id is not set in Prometheus metrics.
    defaultClusterId: 'default-cluster'
    image:
      # -- Exporter container image registry
      registry: quay.io
      # -- Exporter container image name
      repository: kubecost1/kubecost-cost-model
      # -- Exporter container image tag
      # @default -- "" (use appVersion in Chart.yaml)
      tag: "latest"
      # -- Exporter container image pull policy
      pullPolicy: IfNotPresent
    # -- List of extra arguments for the command, e.g.: log-format=json
    extraArgs: []
    # -- Number of OpenCost replicas to run
    replicas: 1
    resources:
      # -- CPU/Memory resource requests
      requests:
        cpu: '10m'
        memory: '55Mi'
      # -- CPU/Memory resource limits
      limits:
        cpu: '999m'
        memory: '1Gi'
    # Liveness probe configuration
    livenessProbe:
      # -- Whether probe is enabled
      enabled: true
      # -- Number of seconds before probe is initiated
      initialDelaySeconds: 120
      # -- Probe frequency in seconds
      periodSeconds: 10
      # -- Number of failures for probe to be considered failed
      failureThreshold: 3
    # Readiness probe configuration
    readinessProbe:
      # -- Whether probe is enabled
      enabled: true
      # -- Number of seconds before probe is initiated
      initialDelaySeconds: 120
      # -- Probe frequency in seconds
      periodSeconds: 10
      # -- Number of failures for probe to be considered failed
      failureThreshold: 3
    # -- The security options the container should be run with
    securityContext: {}
      # capabilities:
      #   drop:
      #   - ALL
      # readOnlyRootFilesystem: true
      # runAsNonRoot: true
      # runAsUser: 1000

    # Path of CSV file
    csv_path: ""

    # Persistent volume claim for storing the data. eg: csv file
    persistence:
      enabled: false
      # -- Annotations for persistent volume
      annotations: {}
      # -- Access mode for persistent volume
      accessMode: ""
      # -- Storage class for persistent volume
      storageClass: ""
      # -- Size for persistent volume
      size: ""

    aws:
      # -- AWS secret access key
      secret_access_key: ""
      # -- AWS secret key id
      access_key_id: ""
    # -- A list of volume mounts to be added to the pod
    extraVolumeMounts: []
    # -- List of additional environment variables to set in the container
    env: []
    # -- Any extra environment variables you would like to pass on to the pod
    extraEnv: {}
      # FOO: BAR
  customPricing:
    # -- Enables custom pricing configuration
    enabled: false
    # -- Customize the configmap name used for custom pricing
    configmapName: custom-pricing-model
    # -- Path for the pricing configuration.
    configPath: /tmp/custom-config
    # -- Configures the pricing model provided in the values file.
    createConfigmap: true
    # -- Sets the provider type for the custom pricing file.
    provider: custom
    # -- More information about these values here: https://www.opencost.io/docs/configuration/on-prem#custom-pricing-using-the-opencost-helm-chart
    costModel:
      description: Modified pricing configuration.
      CPU: 1.25
      spotCPU: 0.006655
      RAM: 0.50
      spotRAM: 0.000892
      GPU: 0.95
      storage: 0.25
      zoneNetworkEgress: 0.01
      regionNetworkEgress: 0.01
      internetNetworkEgress: 0.12

  dataRetention:
    dailyResolutionDays: 15

  cloudCost:
    # -- Enable cloud cost ingestion and querying, dependant on valid integration credentials
    enabled: false
    # -- Number of hours between each run of the Cloud Cost pipeline
    refreshRateHours: 6
    # -- Number of days into the past that a Cloud Cost standard run will query for
    runWindowDays: 3
    # -- The number of standard runs before a Month-to-Date run occurs
    monthToDateInterval: 6
    # -- The max number of days that any single query will be made to construct Cloud Costs
    queryWindowDays: 7


  metrics:
    serviceMonitor:
      # -- Create ServiceMonitor resource for scraping metrics using PrometheusOperator
      enabled: false
      # -- Additional labels to add to the ServiceMonitor
      additionalLabels: {}
      # -- Specify if the ServiceMonitor will be deployed into a different namespace (blank deploys into same namespace as chart)
      namespace: ""
      # -- Interval at which metrics should be scraped
      scrapeInterval: 30s
      # -- Timeout after which the scrape is ended
      scrapeTimeout: 10s
      # -- HonorLabels chooses the metric's labels on collisions with target labels
      honorLabels: true
      # -- RelabelConfigs to apply to samples before scraping. Prometheus Operator automatically adds relabelings for a few standard Kubernetes fields
      relabelings: []
      # -- MetricRelabelConfigs to apply to samples before ingestion
      metricRelabelings: []
      # -- extra Endpoints to add to the ServiceMonitor.  Useful for scraping sidecars
      extraEndpoints: []
        # - port: oauth-metrics
        #   path: /metrics
      # -- HTTP scheme used for scraping. Defaults to http
      scheme: http
      # -- TLS configuration for scraping metrics
      tlsConfig: {}
        # caFile: /etc/prom-certs/root-cert.pem
        # certFile: /etc/prom-certs/cert-chain.pem
        # insecureSkipVerify: true
        # keyFile: /etc/prom-certs/key.pem

  prometheus:
    # -- Secret name that contains credentials for Prometheus
    secret_name: ~
    # -- Prometheus Basic auth username
    username: ""
    # -- Key in the secret that references the username
    username_key: DB_BASIC_AUTH_USERNAME
    # -- Prometheus Basic auth password
    password: ""
    # -- Key in the secret that references the password
    password_key: DB_BASIC_AUTH_PW
    # -- Prometheus Bearer token
    bearer_token: ""
    bearer_token_key: DB_BEARER_TOKEN
    external:
      # -- Use external Prometheus (eg. Grafana Cloud)
      enabled: false
      # -- External Prometheus url
      url: "https://prometheus.example.com/prometheus"
    internal:
      # -- Use in-cluster Prometheus
      enabled: true
      # -- Service name of in-cluster Prometheus
      serviceName: my-prometheus
      # -- Namespace of in-cluster Prometheus
      namespaceName: opencost
      # -- Service port of in-cluster Prometheus
      port: 9090
    amp:
      # -- Use Amazon Managed Service for Prometheus (AMP)
      enabled: false  # If true, opencost will be configured to remote_write and query from Amazon Managed Service for Prometheus.
      # -- Workspace ID for AMP
      workspaceId: ""
    thanos:
      enabled: false
      queryOffset: ''
      maxSourceResolution: ''
      internal:
        enabled: true
        serviceName: my-thanos-query
        namespaceName: opencost
        port: 10901
      external:
        enabled: false
        url: 'https://thanos-query.example.com/thanos'

  ui:
    # -- Enable OpenCost UI
    enabled: true
    image:
      # -- UI container image registry
      registry: quay.io
      # -- UI container image name
      repository: kubecost1/opencost-ui
      # -- UI container image tag
      # @default -- "" (use appVersion in Chart.yaml)
      tag: "latest"
      # -- UI container image pull policy
      pullPolicy: IfNotPresent
    resources:
      # -- CPU/Memory resource requests
      requests:
        cpu: '10m'
        memory: '55Mi'
      # -- CPU/Memory resource limits
      limits:
        cpu: '999m'
        memory: '1Gi'
    # Liveness probe configuration
    livenessProbe:
      # -- Whether probe is enabled
      enabled: true
      # -- Number of seconds before probe is initiated
      initialDelaySeconds: 30
      # -- Probe frequency in seconds
      periodSeconds: 10
      # -- Number of failures for probe to be considered failed
      failureThreshold: 3
    # Readiness probe configuration
    readinessProbe:
      # -- Whether probe is enabled
      enabled: true
      # -- Number of seconds before probe is initiated
      initialDelaySeconds: 30
      # -- Probe frequency in seconds
      periodSeconds: 10
      # -- Number of failures for probe to be considered failed
      failureThreshold: 3
    # -- The security options the container should be run with
    securityContext: {}
      # capabilities:
      #   drop:
      #   - ALL
      # readOnlyRootFilesystem: true
      # runAsNonRoot: true
      # runAsUser: 1000

    # -- A list of volume mounts to be added to the pod
    extraVolumeMounts: []

    ingress:
      # -- Ingress for OpenCost UI
      enabled: false
      # -- Ingress controller which implements the resource
      ingressClassName: ""
      # -- Annotations for Ingress resource
      annotations: {}
        # kubernetes.io/tls-acme: "true"
      # -- A list of host rules used to configure the Ingress
      # @default -- See [values.yaml](values.yaml)
      hosts:
        - host: example.local
          paths:
            - /
      # -- Redirect ingress to an extraPort defined on the service such as oauth-proxy
      servicePort: http-ui
      # servicePort: oauth-proxy
      # -- Ingress TLS configuration
      tls: []
        #  - secretName: chart-example-tls
        #    hosts:
        #      - chart-example.local

  sigV4Proxy:
    image: public.ecr.aws/aws-observability/aws-sigv4-proxy:latest
    imagePullPolicy: IfNotPresent
    name: aps
    port: 8005
    region: us-west-2 # The AWS region
    host: aps-workspaces.us-west-2.amazonaws.com # The hostname for AMP service.
    # role_arn: arn:aws:iam::<account>:role/role-name # The AWS IAM role to assume.
    extraEnv: # Pass extra env variables to sigV4Proxy
    # - name: AWS_ACCESS_KEY_ID
    #   value: <access_key>
    # - name: AWS_SECRET_ACCESS_KEY
    #   value: <secret_key>
    resources: {}
      # limits:
      #   cpu: 200m
      #   memory: 500Mi
      # requests:
      #   cpu: 20m
      #   memory: 32Mi
    securityContext: {}
      # capabilities:
      #   drop:
      #   - ALL
      # readOnlyRootFilesystem: true
      # runAsNonRoot: true
      # runAsUser: 65534
  # -- Toleration labels for pod assignment
  tolerations: []
  # -- Node labels for pod assignment
  nodeSelector: {}
  # -- Affinity settings for pod assignment
  affinity: {}
  # -- Assign custom TopologySpreadConstraints rules
  topologySpreadConstraints: []

  # -- extra sidecars to add to the pod.  Useful for things like oauth-proxy for the UI
  extraContainers: []
    # - name: oauth-proxy
    #   image: quay.io/oauth2-proxy/oauth2-proxy:v7.5.1
    #   args:
    #     - --upstream=http://127.0.0.1:9090
    #     - --http-address=0.0.0.0:8081
    #     - --metrics-address=0.0.0.0:8082
    #     - ...
    #   ports:
    #     - name: oauth-proxy
    #       containerPort: 8081
    #       protocol: TCP
    #     - name: oauth-metrics
    #       containerPort: 8082
    #       protocol: TCP
    #   resources: {}

# -- A list of volumes to be added to the pod
extraVolumes: []`
