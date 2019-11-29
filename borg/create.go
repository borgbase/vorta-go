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
	models.DB.Model(&profile).Related(&r.Repo)

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

	// TODO: implement exclude-if-present
	//if profile.exclude_if_present is not None:
	//for f in profile.exclude_if_present.split('\n'):
	//if f.strip():
	//cmd.extend(['--exclude-if-present', f.strip()])

	newArchiveName := r.Profile.FormatArchiveName()
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
	newArchive.ArchiveId = r.Result.GetPath("archive", "id").MustString()
	newArchive.Name = r.Result.GetPath("archive", "name").MustString()
	newArchive.RepoId = r.Repo.ID
	newArchive.Duration = sql.NullFloat64{r.Result.GetPath("archive", "duration").MustFloat64(), true}
	newArchive.Size = sql.NullInt64{r.Result.GetPath("archive", "duration").MustInt64(), true}
	models.DB.Create(newArchive)

	// Update repo space stats
	r.Repo.UniqueSize = sql.NullInt64{r.Result.GetPath("cache", "stats", "unique_size").MustInt64(), true}
	r.Repo.UniqueCsize = sql.NullInt64{r.Result.GetPath("cache", "stats", "unique_csize").MustInt64(), true}
	r.Repo.TotalSize = sql.NullInt64{r.Result.GetPath("cache", "stats", "total_size").MustInt64(), true}
	r.Repo.TotalUniqueChunks = sql.NullInt64{r.Result.GetPath("cache", "stats", "total_unique_chunks").MustInt64(), true}
	models.DB.Save(&r.Repo)
}
