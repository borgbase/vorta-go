package borg

import (
	"strings"
	"vorta-go/models"
)

type InfoRun struct {
	BorgRun
}

func NewInfoRun(profile *models.Profile, repoUrl, repoPassword, extraBorgArgs string) (*InfoRun, error) {
	r := &InfoRun{}
	r.SubCommand = "info"
	r.SubCommandArgs = []string{"--json"}
	r.Profile = profile
	if len(extraBorgArgs) > 0 {
		r.CommonBorgArgs = strings.Split(extraBorgArgs, " ")
	}

	// For unencrypted repos, set a dummy password.
	if repoPassword == "" {
		r.RepoPassword = "999"
	} else {
		r.RepoPassword = repoPassword
	}

	err := r.Prepare()
	if err != nil {
		return nil, err
	}
	r.SubCommandArgs = append(r.SubCommandArgs, repoUrl)
	return r, nil
}

func (r *InfoRun) ProcessResult() {
	
}
