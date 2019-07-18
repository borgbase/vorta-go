package borg

import (
	"vorta-go/models"
)

type InfoRun struct {
	BorgRun
}

func NewInfoRun(profile *models.Profile) (*InfoRun, error) {
	r := &InfoRun{}
	r.SubCommand = "info"
	r.SubCommandArgs = []string{"--json"}
	r.Profile = profile

	err := r.Prepare()
	if err != nil {
		return nil, err
	}

	r.SubCommandArgs = append(r.SubCommandArgs, r.Repo.Url)
	return r, nil
}

func (r *InfoRun) ProcessResult() {
	
}
