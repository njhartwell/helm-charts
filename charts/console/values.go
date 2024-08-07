// +gotohelm:ignore=true
package console

import (
	_ "embed"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
)

//go:embed values.yaml
var DefaultValuesYAML []byte

type Values struct {
	ReplicaCount                 int                               `json:"replicaCount"`
	Image                        Image                             `json:"image"`
	ImagePullSecrets             []corev1.LocalObjectReference     `json:"imagePullSecrets"`
	NameOverride                 string                            `json:"nameOverride"`
	FullnameOverride             string                            `json:"fullnameOverride"`
	AutomountServiceAccountToken bool                              `json:"automountServiceAccountToken"`
	ServiceAccount               ServiceAccount                    `json:"serviceAccount"`
	CommonLabels                 map[string]string                 `json:"commonLabels"`
	Annotations                  map[string]string                 `json:"annotations"`
	PodAnnotations               map[string]string                 `json:"podAnnotations"`
	PodLabels                    map[string]string                 `json:"podLabels"`
	PodSecurityContext           corev1.PodSecurityContext         `json:"podSecurityContext"`
	SecurityContext              corev1.SecurityContext            `json:"securityContext"`
	Service                      Service                           `json:"service"`
	Ingress                      Ingress                           `json:"ingress"`
	Resources                    corev1.ResourceRequirements       `json:"resources"`
	Autoscaling                  AutoScaling                       `json:"autoscaling"`
	NodeSelector                 map[string]string                 `json:"nodeSelector"`
	Tolerations                  []corev1.Toleration               `json:"tolerations"`
	Affinity                     corev1.Affinity                   `json:"affinity"`
	TopologySpreadConstraints    []corev1.TopologySpreadConstraint `json:"topologySpreadConstraints"`
	PriorityClassName            string                            `json:"priorityClassName"`
	Console                      Console                           `json:"console"`
	ExtraEnv                     []corev1.EnvVar                   `json:"extraEnv"`
	ExtraEnvFrom                 []corev1.EnvFromSource            `json:"extraEnvFrom"`
	ExtraVolumes                 []corev1.Volume                   `json:"extraVolumes"`
	ExtraVolumeMounts            []corev1.VolumeMount              `json:"extraVolumeMounts"`
	ExtraContainers              []corev1.Container                `json:"extraContainers"`
	InitContainers               InitContainers                    `json:"initContainers"`
	SecretMounts                 []SecretMount                     `json:"secretMounts"`
	Secret                       Secret                            `json:"secret"`
	Enterprise                   Enterprise                        `json:"enterprise"`
	LivenessProbe                corev1.Probe                      `json:"livenessProbe"`
	ReadinessProbe               corev1.Probe                      `json:"readinessProbe"`
	ConfigMap                    Creatable                         `json:"configmap"`
	Deployment                   Creatable                         `json:"deployment"`
	Strategy                     appsv1.DeploymentStrategy         `json:"strategy"`
	Tests                        Enableable                        `json:"tests"`
}

type Enterprise struct {
	LicenseSecretRef SecretKeyRef `json:"licenseSecretRef"`
}

type ServiceAccount struct {
	Create                       bool              `json:"create"`
	AutomountServiceAccountToken bool              `json:"automountServiceAccountToken"`
	Annotations                  map[string]string `json:"annotations"`
	Name                         string            `json:"name"`
}

type Service struct {
	Type        string            `json:"type"`
	Port        int               `json:"port"`
	NodePort    *int              `json:"nodePort,omitempty"`
	TargetPort  *int              `json:"targetPort"`
	Annotations map[string]string `json:"annotations"`
}

type Ingress struct {
	Enabled     bool                      `json:"enabled"`
	ClassName   string                    `json:"className"`
	Annotations map[string]string         `json:"annotations"`
	Hosts       []IngressHost             `json:"hosts"`
	TLS         []networkingv1.IngressTLS `json:"tls"`
}

type IngressHost struct {
	Host  string        `json:"host"`
	Paths []IngressPath `json:"paths"`
}

type IngressPath struct {
	Path     string `json:"path"`
	PathType string `json:"pathType"`
}

type AutoScaling struct {
	Enabled                           bool `json:"enabled"`
	MinReplicas                       int  `json:"minReplicas"`
	MaxReplicas                       int  `json:"maxReplicas"`
	TargetCPUUtilizationPercentage    int  `json:"targetCPUUtilizationPercentage"`
	TargetMemoryUtilizationPercentage *int `json:"targetMemoryUtilizationPercentage,omitempty"`
}

// TODO the typing of these values are unclear. All of them get marshalled to
// YAML and then run through tpl which gives no indication of what they are
// aside from YAML marshal-able.
type Console struct {
	Config       any              `json:"config"`
	Roles        []map[string]any `json:"roles,omitempty"`
	RoleBindings []map[string]any `json:"roleBindings,omitempty"`
}

type InitContainers struct {
	ExtraInitContainers *string `json:"extraInitContainers"` // XXX Templated YAML
}

type Secret struct {
	Create     bool              `json:"create"`
	Kafka      KafkaSecrets      `json:"kafka"`
	Login      LoginSecrets      `json:"login"`
	Enterprise EnterpriseSecrets `json:"enterprise"`
	Redpanda   RedpandaSecrets   `json:"redpanda"`
}

type SecretMount struct {
	Name        string `json:"name"`
	SecretName  string `json:"secretName"`
	Path        string `json:"path"`
	DefaultMode int    `json:"defaultMode"`
}

type KafkaSecrets struct {
	SASLPssword                  string `json:"saslPassword,omitempty"`
	AWSMSKIAMSecretKey           string `json:"awsMskIamSecretKey,omitempty"`
	TLSCA                        string `json:"tlsCa,omitempty"`
	TLSCert                      string `json:"tlsCert,omitempty"`
	TLSPassphrase                string `json:"tlsPassphrase,omitempty"`
	SchemaRegistryPassword       string `json:"schemaRegistryPassword,omitempty"`
	SchemaRegistryTLSCA          string `json:"schemaRegistryTlsCa,omitempty"`
	SchemaRegistryTLSCert        string `json:"schemaRegistryTlsCert,omitempty"`
	SchemaRegistryTLSKey         string `json:"schemaRegistryTlsKey,omitempty"`
	ProtobufGitBasicAuthPassword string `json:"protobufGitBasicAuthPassword,omitempty"`
}

type LoginSecrets struct {
	JWTSecret string             `json:"jwtSecret"`
	Google    GoogleLoginSecrets `json:"google"`
	Github    GithubLoginSecrets `json:"github"`
	Okta      OktaLoginSecrets   `json:"okta"`
	OIDC      OIDCLoginSecrets   `json:"oidc"`
}

type GoogleLoginSecrets struct {
	ClientSecret         *string `json:"clientSecret,omitempty"`
	GroupsServiceAccount *string `json:"groupsServiceAccount,omitempty"`
}

type GithubLoginSecrets struct {
	ClientSecret        *string `json:"clientSecret,omitempty"`
	PersonalAccessToken *string `json:"personalAccessToken,omitempty"`
}

type OktaLoginSecrets struct {
	ClientSecret      *string `json:"clientSecret,omitempty"`
	DirectoryAPIToken *string `json:"directoryApiToken,omitempty"`
}

type OIDCLoginSecrets struct {
	ClientSecret *string `json:"clientSecret,omitempty"`
}

type EnterpriseSecrets struct {
	LicenseSecretRef *SecretKeyRef `json:"licenseSecretRef,omitempty"`
}

type RedpandaSecrets struct {
	AdminAPI RedpandaAdminAPISecrets `json:"adminApi"`
}

type RedpandaAdminAPISecrets struct {
	Password *string `json:"password,omitempty"`
	TLSCA    *string `json:"tlsCa,omitempty"`
	TLSCert  *string `json:"tlsCert,omitempty"`
	TLSKey   *string `json:"tlsKey,omitempty"`
}

type SecretKeyRef struct {
	Name string `json:"name"`
	Key  string `json:"key"`
}

type Enableable struct {
	Enabled bool `json:"enabled"`
}

type Creatable struct {
	Create bool `json:"create"`
}

type Image struct {
	Registry   string            `json:"registry"`
	Repository string            `json:"repository"`
	PullPolicy corev1.PullPolicy `json:"pullPolicy"`
	Tag        string            `json:"tag"`
}
