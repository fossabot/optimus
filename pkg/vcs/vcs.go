package vcs

import (
	_ "github.com/src-d/go-git"
)

type VCS struct {
	URL    string
	Creds  VCSCreds
	Branch string
	Tag    string
}

type VCSCreds struct {
	Username string
	SSHKey   string
}
