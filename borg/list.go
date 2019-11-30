package borg

import (
	"time"
	"vorta/models"
	"vorta/utils"
)

type ListRepoRun struct {
	BorgRun
}

func NewListRepoRun(profile *models.Profile) (*ListRepoRun, error) {
	r := &ListRepoRun{}
	r.SubCommand = "list"

	r.Profile = profile
	r.Repo = &models.Repo{}
	models.DB.Model(&profile).Related(&r.Repo)

	r.SubCommandArgs = []string{"--json", r.Repo.Url}

	err := r.Prepare()
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (r *ListRepoRun) ProcessResult() {
	utils.Log.Info(r.Result)
	for i, _ := range r.Result.GetPath("archives").MustArray(){
		jsArchive := r.Result.Get("archives").GetIndex(i)
		newArchive := models.Archive{}
		newArchive.ArchiveID = jsArchive.Get("id").MustString()
		newArchive.Name = jsArchive.Get("name").MustString()
		newArchive.CreatedAt, _ = time.Parse(time.RFC3339Nano, jsArchive.Get("time").MustString())
		newArchive.RepoID = r.Repo.ID

		var existingArchives int
		models.DB.Where("snapshot_id = ?", newArchive.ArchiveID).Find(&models.Archive{}).Count(&existingArchives)
		if existingArchives == 0 {
			models.DB.Create(&newArchive)
		}
	}
}
