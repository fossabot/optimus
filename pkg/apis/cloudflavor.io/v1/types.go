package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type PipelineList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Pipelines []Pipeline `json:"pipelines"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Pipeline struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	ArchiveArtifacts bool             `json:"archive"`
	Project          Project          `json:"vcs_details"`
	ContainerBuilder ContainerBuilder `json:"container_builder"`
}

type Project struct {
	Stages         []Stage `json:"stages"`
	PollInterval   uint32  `json:"poll_interval"`
	Repository     string  `json:"repository"`
	Username       string  `json:"usernam"`
	Authentication Secrets `json:"auth_details"`
}

type ContainerBuilder struct {
	URI            string   `json:"builder_uri"`
	Registry       Registry `json:"container_registry"`
	Authentication Secrets  `json:"auth_details"`
}

type Stage struct {
	ExitOnError string    `json:"exit_on_error"`
	Name        string    `json:"stage_name"`
	Parallel    bool      `json:"parallel"`
	Commands    []Command `json:"commands"`
}

type Commands struct {
	ExitOnError string   `json:"exit_on_error"`
	Cmd         string   `json:"command"`
	Parameters  []string `json:"parameters"`
}

type Registry struct {
	URI            string  `json:"registry_uri"`
	Authentication Secrets `json:"auth_details"`
}

// NOTE: if we use k8s native secrets, do we need secrets at all?
type Secrets struct {
	SSHKey          string `json:"ssh_key,omitempty"`
	AccessToken     string `json:"access_token,omitempty"`
	SecretAccessKey string `json:"secret_key,omitempty"`
	AccessKey       string `json:"access_key,omitempty"`
	TLSCert         string `json:"tls_cert,omitempty"`
	TLSCACert       string `json:"tls_ca_cert,omitempty"`
	TLSCertKey      string `json:"tls_cert_key,omitempty"`
}

type Storage struct {
	URI            string  `json:"uri"`
	Authentication Secrets `json:"auth_details"`
}

type NotificationDetails struct {
	URI            string  `json:"notification_uri"`
	Authentication Secrets `json:"auth_details"`
}
