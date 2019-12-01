package borg

import (
	"time"
	"vorta/models"
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
	archiveIDs := map[string]struct{}{}
	for i, _ := range r.Result.GetPath("archives").MustArray(){
		jsArchive := r.Result.Get("archives").GetIndex(i)
		newArchive := models.Archive{}
		newArchive.ArchiveID = jsArchive.Get("id").MustString()
		newArchive.Name = jsArchive.Get("name").MustString()
		newArchive.CreatedAt, _ = time.Parse(time.RFC3339Nano, jsArchive.Get("time").MustString())
		newArchive.RepoID = r.Repo.ID

		models.DB.FirstOrCreate(&newArchive, models.Archive{ArchiveID: newArchive.ArchiveID})

		// build list of archives to remove deleted ones later
		archiveIDs[newArchive.ArchiveID] = struct{}{}
		}

	// Remove entries for archives no found in the repo
	aa := []models.Archive{}
	models.DB.Model(r.Repo).Related(&aa)
	for _, a := range aa {
		if _, ok := archiveIDs[a.ArchiveID]; !ok {
			models.DB.Delete(&a)
		}
	}
}
