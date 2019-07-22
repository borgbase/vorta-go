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

	r.Repo = &models.Repo{}
	r.Repo.Url = repoUrl
	err := r.Repo.SetPassword(repoPassword)
	if err != nil {
		return r, err
	}

	if len(extraBorgArgs) > 0 {
		r.CommonBorgArgs = strings.Split(extraBorgArgs, " ")
	}

	err = r.Prepare()
	if err != nil {
		return nil, err
	}
	r.SubCommandArgs = append(r.SubCommandArgs, repoUrl)
	return r, nil
}

func (r *InfoRun) ProcessResult() {
}
