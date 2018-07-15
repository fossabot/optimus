package v1

type CIProject struct {
	Pipelines []Pipeline `json:"pipelines"`
}

type Pipeline struct {
	ArchiveArtifacts bool       `json:"archive"`
	VCS              VCSDetails `json:"vcs_details"`
	Stages           []Stage    `json:"stages,omitempty"`
}

type Stage struct {
	Command    string   `json:"command"`
	Parameters []string `json:"parameters"`
}

type DockerBuilder struct {
	URI            string   `json:"builder_uri"`
	Registry       Registry `json:"container_registry"`
	Authentication Secrets  `json:"auth_details"`
}

type Registry struct {
	URI            string  `json:"registry_uri"`
	Authentication Secrets `json:"auth_details"`
}

type VCSDetails struct {
	PollInterval   uint32  `json:"poll_interval"`
	Repo           string  `json:"repository"`
	Username       string  `json:"usernam"`
	Authentication Secrets `json:"auth_details"`
}

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
