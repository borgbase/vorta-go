package borg

import (
	"strings"
	"vorta/models"
)

type InitRun struct {
	BorgRun
}

func NewInitRun(profile *models.Profile, repoUrl, repoPassword, extraBorgArgs, encryption string) (*InitRun, error) {
	r := &InitRun{}
	r.SubCommand = "init"
	r.SubCommandArgs = []string{"--encryption", encryption}
	r.Profile = profile
	if len(extraBorgArgs) > 0 {
		r.CommonBorgArgs = strings.Split(extraBorgArgs, " ")
	}

	r.Repo = &models.Repo{}
	r.Repo.Url = repoUrl
	err := r.Repo.SetPassword(repoPassword)
	if err != nil {
		return nil, err
	}

	err = r.Prepare()
	if err != nil {
		return nil, err
	}
	r.SubCommandArgs = append(r.SubCommandArgs, repoUrl)
	return r, nil
}

func (r *InitRun) ProcessResult() {
}
