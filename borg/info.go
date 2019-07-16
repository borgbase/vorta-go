package borg

import (
	"vorta-go/models"
)

type InfoRun struct {
	BorgRun
}

func NewInfoRun(profile *models.Profile, repo *models.Repo) (*InfoRun, error) {
	r := &InfoRun{}
	r.SubCommand = "info"
	r.SubCommandArgs = []string{"--json"}
	r.Profile = profile
	r.Repo = repo

	err := r.Prepare()
	if err != nil {
		return nil, err
	}

	r.SubCommandArgs = append(r.SubCommandArgs, repo.Url)
	return r, nil
}

func (r *InfoRun) ProcessResult() {
	
}
