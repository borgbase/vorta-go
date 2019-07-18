package borg

import (
	"io/ioutil"
	"os"
	"strings"
	"vorta-go/models"
	"vorta-go/utils"
)

type CreateRun struct {
	BorgRun
}

func NewCreateRun(profile *models.Profile) (*CreateRun, error) {
	r := &CreateRun{}
	r.SubCommand = "create"
	r.SubCommandArgs = []string{"--json", "--list", "--filter=AM", "-C", profile.Compression}
	r.Profile = profile
	r.Repo = profile.GetRepo()

	// Do global preparations
	err := r.Prepare()
	if err != nil {
		return nil, err
	}

	// Write exclude patterns to temp file
	excludePatterns := []string{}
	excludeString := r.Profile.ExcludePatterns.String
	if excludeString != "" {
		lines := strings.Split(excludeString,`\n`)
		for  _, l := range lines {
			l = strings.TrimSpace(l)
			excludePatterns = append(excludePatterns, l)
		}

		tmpFile, _ := ioutil.TempFile(os.TempDir(), "prefix-")
		tmpFile.Write([]byte(strings.Join(excludePatterns[:], "\n")))
		r.SubCommandArgs = append(r.SubCommandArgs, "--exclude-from", tmpFile.Name())
	}

	// TODO: implement exclude-if-present
	//if profile.exclude_if_present is not None:
	//for f in profile.exclude_if_present.split('\n'):
	//if f.strip():
	//cmd.extend(['--exclude-if-present', f.strip()])

	newArchiveName := _formatArchiveName(r.Profile)
	r.SubCommandArgs = append(r.SubCommandArgs, r.Repo.Url+"::"+newArchiveName)

	ss := []models.SourceDir{}
	err = models.DB.Select(&ss, models.SqlAllSourcesByProfileId, r.Profile.Id)
	if err != nil {
		utils.Log.Error(err)
	}
	for _, s := range ss {
		r.SubCommandArgs = append(r.SubCommandArgs, s.Dir)
	}
	utils.Log.Debug(r.SubCommandArgs)
	return r, nil
}

func (r *CreateRun) ProcessResult() {
	
}
