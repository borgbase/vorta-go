package borg

import (
	"vorta/models"
)

type CheckRun struct {
	BorgRun
}

func NewCheckRun(profile *models.Profile) (*CheckRun, error) {
	r := &CheckRun{}
	r.SubCommand = "check"
	r.Profile = profile
	r.Repo = profile.GetRepo()

	err := r.Prepare()
	if err != nil {
		return nil, err
	}
	r.SubCommandArgs = append(r.SubCommandArgs, r.Repo.Url)
	return r, nil
}

func (r *CheckRun) ProcessResult() {
}
