package borg

import (
	"strconv"
	"vorta/models"
	"vorta/utils"
)

type PruneRun struct {
	BorgRun
}

func NewPruneRun(profile *models.Profile) (*PruneRun, error) {
	r := &PruneRun{}
	r.SubCommand = "prune"

	r.Profile = profile
	r.Repo = &models.Repo{}
	models.DB.Model(&profile).Related(&r.Repo)

	r.SubCommandArgs = []string{
		"--list",
		"--keep-hourly", strconv.Itoa(profile.PruneHour),
		"--keep-daily", strconv.Itoa(profile.PruneDay),
		"--keep-weekly", strconv.Itoa(profile.PruneWeek),
		"--keep-monthly", strconv.Itoa(profile.PruneMonth),
		"--keep-yearly", strconv.Itoa(profile.PruneYear),
		"--prefix", profile.FormatArchiveName(profile.PrunePrefix),
	}

	if profile.PruneKeepWithin.Valid && len(profile.PruneKeepWithin.String) > 0 {
		r.SubCommandArgs = append(r.SubCommandArgs, "--keep-within", profile.PruneKeepWithin.String)
	}

	r.SubCommandArgs = append(r.SubCommandArgs, r.Repo.Url)

	err := r.Prepare()
	if err != nil {
		return nil, err
	}
	return r, nil
}

// List repo content to get rid of pruned archives.
func (r *PruneRun) ProcessResult() {
	listRun, err := NewListRepoRun(r.Profile)
	if err != nil {
		utils.Log.Error(err)
		return
	}
	err = listRun.Run()
	if err != nil {
		utils.Log.Error(err)
		return
	}
	listRun.ProcessResult()
}
