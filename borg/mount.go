package borg

import (
	"vorta/models"
	"os/user"
)

type MountRun struct {
	BorgRun
}

func NewMountRun(profile *models.Profile) (*MountRun, error) {
	r := &MountRun{}
	r.SubCommand = "mount"

	r.Profile = profile
	r.Repo = &models.Repo{}
	models.DB.Model(&profile).Related(&r.Repo)

	currentUser, _ := user.Current()
	r.SubCommandArgs = []string{"-o", "umask=0277,uid=" + currentUser.Uid, r.Repo.Url}

	err := r.Prepare()
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (r *MountRun) ProcessResult() {
}
