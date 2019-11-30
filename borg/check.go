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
	r.Repo = &models.Repo{}
	models.DB.Model(&profile).Related(&r.Repo)

	err := r.Prepare()
	if err != nil {
		return nil, err
	}
	r.SubCommandArgs = append(r.SubCommandArgs, r.Repo.Url)
	return r, nil
}

func (r *CheckRun) ProcessResult() {
}
