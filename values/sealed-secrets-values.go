package values

const SealedSecret = `
## @section Common parameters

## @param kubeVersion Override Kubernetes version
##
kubeVersion: ""
## @param nameOverride String to partially override sealed-secrets.fullname
##
nameOverride: ""
## @param fullnameOverride String to fully override sealed-secrets.fullname
##
fullnameOverride: ""
## @param namespace Namespace where to deploy the Sealed Secrets controller
##
namespace: ""
## @param extraDeploy [array] Array of extra objects to deploy with the release
##
extraDeploy: []

## @section Sealed Secrets Parameters

## Sealed Secrets image
## ref: https://quay.io/repository/bitnami/sealed-secrets-controller?tab=tags
## @param image.registry Sealed Secrets image registry
## @param image.repository Sealed Secrets image repository
## @param image.tag Sealed Secrets image tag (immutable tags are recommended)
## @param image.pullPolicy Sealed Secrets image pull policy
## @param image.pullSecrets [array]  Sealed Secrets image pull secrets
##
image:
  registry: docker.io
  repository: bitnami/sealed-secrets-controller
  tag: v0.17.3
  ## Specify a imagePullPolicy
  ## Defaults to 'Always' if image tag is 'latest', else set to 'IfNotPresent'
  ## ref: http://kubernetes.io/docs/user-guide/images/#pre-pulling-images
  ##
  pullPolicy: IfNotPresent
  ## Optionally specify an array of imagePullSecrets.
  ## Secrets must be manually created in the namespace.
  ## ref: https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/
  ## e.g:
  ## pullSecrets:
  ##   - myRegistryKeySecretName
  ##
  pullSecrets: []
## @param createController Specifies whether the Sealed Secrets controller should be created
##
createController: true
## @param secretName The name of an existing TLS secret containing the key used to encrypt secrets
##
secretName: "sealed-secrets-key"
## @param updateStatus Specifies whether the Sealed Secrets controller should update the status subresource
##
updateStatus: true
## @param keyrenewperiod Specifies key renewal period. Default 30 days
## e.g
## keyrenewperiod: "720h30m"
##
keyrenewperiod: ""
## @param command Override default container command
##
command: []
## @param args Override default container args
##
args: []
## Configure extra options for Sealed Secret containers' liveness, readiness and startup probes
## ref: https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/#configure-probes
## @param livenessProbe.enabled Enable livenessProbe on Sealed Secret containers
## @param livenessProbe.initialDelaySeconds Initial delay seconds for livenessProbe
## @param livenessProbe.periodSeconds Period seconds for livenessProbe
## @param livenessProbe.timeoutSeconds Timeout seconds for livenessProbe
## @param livenessProbe.failureThreshold Failure threshold for livenessProbe
## @param livenessProbe.successThreshold Success threshold for livenessProbe
##
livenessProbe:
  enabled: true
  initialDelaySeconds: 0
  periodSeconds: 10
  timeoutSeconds: 1
  failureThreshold: 3
  successThreshold: 1
## @param readinessProbe.enabled Enable readinessProbe on Sealed Secret containers
## @param readinessProbe.initialDelaySeconds Initial delay seconds for readinessProbe
## @param readinessProbe.periodSeconds Period seconds for readinessProbe
## @param readinessProbe.timeoutSeconds Timeout seconds for readinessProbe
## @param readinessProbe.failureThreshold Failure threshold for readinessProbe
## @param readinessProbe.successThreshold Success threshold for readinessProbe
##
readinessProbe:
  enabled: true
  initialDelaySeconds: 0
  periodSeconds: 10
  timeoutSeconds: 1
  failureThreshold: 3
  successThreshold: 1
## @param startupProbe.enabled Enable startupProbe on Sealed Secret containers
## @param startupProbe.initialDelaySeconds Initial delay seconds for startupProbe
## @param startupProbe.periodSeconds Period seconds for startupProbe
## @param startupProbe.timeoutSeconds Timeout seconds for startupProbe
## @param startupProbe.failureThreshold Failure threshold for startupProbe
## @param startupProbe.successThreshold Success threshold for startupProbe
##
startupProbe:
  enabled: false
  initialDelaySeconds: 0
  periodSeconds: 10
  timeoutSeconds: 1
  failureThreshold: 3
  successThreshold: 1
## @param customLivenessProbe Custom livenessProbe that overrides the default one
##
customLivenessProbe: {}
## @param customReadinessProbe Custom readinessProbe that overrides the default one
##
customReadinessProbe: {}
## @param customStartupProbe Custom startupProbe that overrides the default one
##
customStartupProbe: {}
## Sealed Secret resource requests and limits
## ref: http://kubernetes.io/docs/user-guide/compute-resources/
## @param resources.limits [object] The resources limits for the Sealed Secret containers
## @param resources.requests [object] The requested resources for the Sealed Secret containers
##
resources:
  limits: {}
  requests: {}
## Configure Pods Security Context
## ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/#set-the-security-context-for-a-pod
## @param podSecurityContext.enabled Enabled Sealed Secret pods' Security Context
## @param podSecurityContext.fsGroup Set Sealed Secret pod's Security Context fsGroup
##
podSecurityContext:
  enabled: true
  fsGroup: 65534
## Configure Container Security Context
## ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/#set-the-security-context-for-a-pod
## @param containerSecurityContext.enabled Enabled Sealed Secret containers' Security Context
## @param containerSecurityContext.readOnlyRootFilesystem Whether the Sealed Secret container has a read-only root filesystem
## @param containerSecurityContext.runAsNonRoot Indicates that the Sealed Secret container must run as a non-root user
## @param containerSecurityContext.runAsUser Set Sealed Secret containers' Security Context runAsUser
##
containerSecurityContext:
  enabled: true
  readOnlyRootFilesystem: true
  runAsNonRoot: true
  runAsUser: 1001
## @param podLabels [object] Extra labels for Sealed Secret pods
## ref: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/
##
podLabels: {}
## @param podAnnotations [object] Annotations for Sealed Secret pods
## ref: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/
##
podAnnotations: {}
## @param priorityClassName Sealed Secret pods' priorityClassName
##
priorityClassName: ""
## @param affinity [object] Affinity for Sealed Secret pods assignment
## ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#affinity-and-anti-affinity
##
affinity: {}
## @param nodeSelector [object] Node labels for Sealed Secret pods assignment
## ref: https://kubernetes.io/docs/user-guide/node-selection/
##
nodeSelector: {}
## @param tolerations [array] Tolerations for Sealed Secret pods assignment
## ref: https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/
##
tolerations: []

## @section Traffic Exposure Parameters

## Sealed Secret service parameters
##
service:
  ## @param service.type Sealed Secret service type
  ##
  type: ClusterIP
  ## @param service.port Sealed Secret service HTTP port
  ##
  port: 8080
  ## @param service.nodePort Node port for HTTP
  ## Specify the nodePort value for the LoadBalancer and NodePort service types
  ## ref: https://kubernetes.io/docs/concepts/services-networking/service/#type-nodeport
  ## NOTE: choose port between <30000-32767>
  ##
  nodePort: ""
  ## @param service.annotations [object] Additional custom annotations for Sealed Secret service
  ##
  annotations: {}
## Sealed Secret ingress parameters
## ref: http://kubernetes.io/docs/user-guide/ingress/
##
ingress:
  ## @param ingress.enabled Enable ingress record generation for Sealed Secret
  ##
  enabled: true
  ## @param ingress.pathType Ingress path type
  ##
  pathType: ImplementationSpecific
  ## @param ingress.apiVersion Force Ingress API version (automatically detected if not set)
  ##
  apiVersion: ""
  ## @param ingress.ingressClassName IngressClass that will be be used to implement the Ingress
  ## This is supported in Kubernetes 1.18+ and required if you have more than one IngressClass marked as the default for your cluster.
  ## ref: https://kubernetes.io/blog/2020/04/02/improvements-to-the-ingress-api-in-kubernetes-1.18/
  ##
  ingressClassName: ""
  ## @param ingress.hostname Default host for the ingress record
  ##
  hostname: sealed-secrets.campus.clusterdiali.me
  ## @param ingress.path Default path for the ingress record
  ##
  path: /v1/cert.pem
  ## @param ingress.annotations [object] Additional annotations for the Ingress resource. To enable certificate autogeneration, place here your cert-manager annotations.
  ## Use this parameter to set the required annotations for cert-manager, see
  ## ref: https://cert-manager.io/docs/usage/ingress/#supported-annotations
  ## e.g:
  ## annotations:
  ##   kubernetes.io/ingress.class: nginx
  ##   cert-manager.io/cluster-issuer: cluster-issuer-name
  ##
  annotations:
    kubernetes.io/ingress.class: nginx
    kubernetes.io/tls-acme: "true"
    cert-manager.io/issuer: letsencrypt-campus
  ## @param ingress.tls Enable TLS configuration for the host defined at ingress.hostname parameter
  ## TLS certificates will be retrieved from a TLS secret with name: {{- printf "%s-tls" .Values.ingress.hostname }}
  ## You can:
  ##   - Use the ingress.secrets parameter to create this TLS secret
  ##   - Relay on cert-manager to create it by setting the corresponding annotations
  ##   - Relay on Helm to create self-signed certificates by setting ingress.selfSigned=true
  ##
  tls: true
  ## @param ingress.selfSigned Create a TLS secret for this ingress record using self-signed certificates generated by Helm
  ##
  selfSigned: false
  ## @param ingress.extraHosts [array] An array with additional hostname(s) to be covered with the ingress record
  ## e.g:
  ## extraHosts:
  ##   - name: sealed-secrets.local
  ##     path: /
  ##
  extraHosts: []
  ## @param ingress.extraPaths [array] An array with additional arbitrary paths that may need to be added to the ingress under the main host
  ## e.g:
  ## extraPaths:
  ## - path: /*
  ##   backend:
  ##     serviceName: ssl-redirect
  ##     servicePort: use-annotation
  ##
  extraPaths: []
  ## @param ingress.extraTls [array] TLS configuration for additional hostname(s) to be covered with this ingress record
  ## ref: https://kubernetes.io/docs/concepts/services-networking/ingress/#tls
  ## e.g:
  ## extraTls:
  ## - hosts:
  ##     - sealed-secrets.local
  ##   secretName: sealed-secrets.local-tls
  ##
  extraTls: []
  ## @param ingress.secrets [array] Custom TLS certificates as secrets
  ## NOTE: 'key' and 'certificate' are expected in PEM format
  ## NOTE: 'name' should line up with a 'secretName' set further up
  ## If it is not set and you're using cert-manager, this is unneeded, as it will create a secret for you with valid certificates
  ## If it is not set and you're NOT using cert-manager either, self-signed certificates will be created valid for 365 days
  ## It is also possible to create and manage the certificates outside of this helm chart
  ## Please see README.md for more information
  ## e.g:
  ## secrets:
  ##   - name: sealed-secrets.local-tls
  ##     key: |-
  ##       -----BEGIN RSA PRIVATE KEY-----
  ##       ...
  ##       -----END RSA PRIVATE KEY-----
  ##     certificate: |-
  ##       -----BEGIN CERTIFICATE-----
  ##       ...
  ##       -----END CERTIFICATE-----
  ##
  secrets: []
## Network policies
## Ref: https://kubernetes.io/docs/concepts/services-networking/network-policies/
##
networkPolicy:
  ## @param networkPolicy.enabled Specifies whether a NetworkPolicy should be created
  ##
  enabled: false

## @section Other Parameters

## ServiceAccount configuration
##
serviceAccount:
  ## @param serviceAccount.create Specifies whether a ServiceAccount should be created
  ##
  create: true
  ## @param serviceAccount.labels Extra labels to be added to the ServiceAccount
  ##
  labels: {}
  ## @param serviceAccount.name The name of the ServiceAccount to use.
  ## If not set and create is true, a name is generated using the sealed-secrets.fullname template
  ##
  name: ""
## RBAC configuration
##
rbac:
  ## @param rbac.create Specifies whether RBAC resources should be created
  ##
  create: true
  ## @param rbac.labels Extra labels to be added to RBAC resources
  ##
  labels: {}
  ## @param rbac.pspEnabled PodSecurityPolicy
  ##
  pspEnabled: false

## @section Metrics parameters

metrics:
  ## Prometheus Operator ServiceMonitor configuration
  ##
  serviceMonitor:
    ## @param metrics.serviceMonitor.enabled Specify if a ServiceMonitor will be deployed for Prometheus Operator
    ##
    enabled: false
    ## @param metrics.serviceMonitor.namespace Namespace where Prometheus Operator is running in
    ##
    namespace: ""
    ## @param metrics.serviceMonitor.labels Extra labels for the ServiceMonitor
    ##
    labels: {}
    ## @param metrics.serviceMonitor.annotations Extra annotations for the ServiceMonitor
    ##
    annotations: {}
    ## @param metrics.serviceMonitor.interval How frequently to scrape metrics
    ## e.g:
    ## interval: 10s
    ##
    interval: ""
    ## @param metrics.serviceMonitor.scrapeTimeout Timeout after which the scrape is ended
    ## e.g:
    ## scrapeTimeout: 10s
    ##
    scrapeTimeout: ""
    ## @param metrics.serviceMonitor.metricRelabelings [array] Specify additional relabeling of metrics
    ##
    metricRelabelings: []
    ## @param metrics.serviceMonitor.relabelings [array] Specify general relabeling
    ##
    relabelings: []
  ## Grafana dashboards configuration
  ##
  dashboards:
    ## @param metrics.dashboards.create Specifies whether a ConfigMap with a Grafana dashboard configuration should be created
    ## ref https://github.com/helm/charts/tree/master/stable/grafana#configuration
    ##
    create: false
    ## @param metrics.dashboards.labels Extra labels to be added to the Grafana dashboard ConfigMap
    ##
    labels: {}
    ## @param metrics.dashboards.namespace Namespace where Grafana dashboard ConfigMap is deployed
    ##
    namespace: ""`
