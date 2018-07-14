package v1

type CIProject struct {
	Pipelines []Pipeline `json:"pipelines"`
}

type Pipeline struct {
	ArchiveArtifacts bool       `json:"archive,omitempty"`
	VCS              VCSDetails `json:"vcs_details"`
	Stages           []Stage    `json:"stages,omitempty"`
}

type Stage struct {
	Command    string   `json:"command"`
	Parameters []string `json:"parameters"`
}

type VCSDetails struct {
	PollInterval uint32 `json:"poll_interval"`
	Repo         string `json:"repository"`
	Username     string `json:"usernam"`
	SSHKey       string `json:"ssh_key"`
}

type Storage struct {
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
	Token     string `json:"token"`
}
