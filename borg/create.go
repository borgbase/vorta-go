package borg

import (
	"database/sql"
	"github.com/mitchellh/go-homedir"
	"io/ioutil"
	"os"
	"strings"
	"vorta/models"
	"vorta/utils"
)

type CreateRun struct {
	BorgRun
}

func NewCreateRun(profile *models.Profile) (*CreateRun, error) {
	r := &CreateRun{}
	r.SubCommand = "create"
	r.SubCommandArgs = []string{"--json", "--list", "--filter=AM", "-C", profile.Compression}
	r.Profile = profile
	r.Repo = &models.Repo{}
	models.DB.Model(profile).Related(r.Repo)

	// Do global preparations
	err := r.Prepare()
	if err != nil {
		return nil, err
	}

	// Write exclude patterns to temp file
	excludePatterns := []string{}
	excludeString := r.Profile.ExcludePatterns.String
	if excludeString != "" {
		lines := strings.Split(excludeString, "\n")
		for _, l := range lines {
			l = strings.TrimSpace(l)
			l, err := homedir.Expand(l)
			utils.Log.Info(l)
			if err != nil {
				utils.Log.Error(err)
				continue
			}
			excludePatterns = append(excludePatterns, l)
		}

		tmpFile, _ := ioutil.TempFile(os.TempDir(), "borg-exclude-from-")
		tmpFile.Write([]byte(strings.Join(excludePatterns[:], "\n")))
		r.SubCommandArgs = append(r.SubCommandArgs, "--exclude-from", tmpFile.Name())
		utils.Log.Infof("Writing exclude file to %v", tmpFile.Name())
	}

	// Append exclude-if-present patterns
	if r.Profile.ExcludeIfPresent.Valid && len(r.Profile.ExcludeIfPresent.String) > 3 {
		for _, e := range strings.Split(r.Profile.ExcludeIfPresent.String, "\n") {
			trimmed := strings.TrimSpace(e)
			if len(trimmed) > 3 {
				r.SubCommandArgs = append(r.SubCommandArgs, "--exclude-if-present", trimmed)
			}
		}
	}

	newArchiveName := r.Profile.FormatArchiveName(r.Profile.NewArchiveName)
	r.SubCommandArgs = append(r.SubCommandArgs, r.Repo.Url+"::"+newArchiveName)

	ss := []models.SourceDir{}
	models.DB.Model(r.Profile).Related(&ss, "SourceDirs")

	for _, s := range ss {
		r.SubCommandArgs = append(r.SubCommandArgs, s.Dir)
	}
	utils.Log.Debug(r.SubCommandArgs)
	return r, nil
}

func (r *CreateRun) ProcessResult() {
	// Save new archive
	newArchive := models.Archive{}
	newArchive.ArchiveID = r.Result.GetPath("archive", "id").MustString()
	newArchive.Name = r.Result.GetPath("archive", "name").MustString()
	newArchive.RepoID = r.Repo.ID
	newArchive.Duration = sql.NullFloat64{r.Result.GetPath("archive", "duration").MustFloat64(), true}
	newArchive.Size = sql.NullInt64{r.Result.GetPath("archive", "stats", "deduplicated_size").MustInt64(), true}
	models.DB.Create(&newArchive)

	// Update repo space stats
	r.Repo.UniqueSize = sql.NullInt64{r.Result.GetPath("cache", "stats", "unique_size").MustInt64(), true}
	r.Repo.UniqueCsize = sql.NullInt64{r.Result.GetPath("cache", "stats", "unique_csize").MustInt64(), true}
	r.Repo.TotalSize = sql.NullInt64{r.Result.GetPath("cache", "stats", "total_size").MustInt64(), true}
	r.Repo.TotalUniqueChunks = sql.NullInt64{r.Result.GetPath("cache", "stats", "total_unique_chunks").MustInt64(), true}
	models.DB.Save(&r.Repo)
}
