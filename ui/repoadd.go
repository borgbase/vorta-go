package ui

import (
	"database/sql"
	"fmt"
	"github.com/therecipe/qt/core"
	"regexp"
	"vorta/borg"
	"vorta/models"
	"vorta/utils"
)

var encryptionModes = map[string]string{
	"repokey-blake2": "Repokey-Blake2 (Recommended, key stored in repository)",
	"repokey":        "Repokey",
	"keyfile-blake2": "Keyfile-Blake2 (Key stored in home directory)",
	"keyfile":        "Keyfile",
	"none":           "None (not recommended)",
} // TODO: populate list in order.

func (d *RepoAddDialog) init() {
	d.TabWidget.SetCurrentIndex(0)

	for k, v := range encryptionModes {
		d.EncryptionComboBox.AddItem(v, core.NewQVariant1(k))
	}

	// d.RepoURL accepts a remote SSH addr or a local folder.
	// d.UseRemoteRepoButton decides which one is currently used. Default = remote repo
	d.UseRemoteRepoButton.SetDisabled(true)
	d.UseRemoteRepoButton.ConnectClicked(d.useRemoteRepoUrl)
	d.ChooseLocalFolderButton.ConnectClicked(d.setLocalFolder)
	d.SaveButton.ConnectClicked(d.ProcessNewRepo)
	d.CloseButton.ConnectClicked(func(_ bool) {
		d.Close()
	})
}

func (d *RepoAddDialog) Validate() error {
	isRemoteRepo := !d.UseRemoteRepoButton.IsEnabled() // This button is disabled when we have a remote repo.
	isValidRemoteAddr, _ := regexp.MatchString(`.+:.+`, d.RepoURL.Text())

	if isRemoteRepo && !isValidRemoteAddr {
		return fmt.Errorf("Please enter a valid repo URL or select a local path.")
	}

	selectedEncryption := d.EncryptionComboBox.CurrentData(int(core.Qt__UserRole)).ToString()
	if selectedEncryption != "none" && len(d.PasswordLineEdit.Text()) < 8 {
		return fmt.Errorf("Please use a longer password.")
	}
	return nil
}

func (d *RepoAddDialog) setLocalFolder(_ bool) {
	ChooseFileDialog(func(files []string) {
		utils.Log.Info(files)
		d.RepoURL.SetText(files[0])
		d.UseRemoteRepoButton.SetDisabled(false)
	})
}

func (d *RepoAddDialog) useRemoteRepoUrl(_ bool) {
	d.UseRemoteRepoButton.SetDisabled(true)
	d.RepoURL.SetText("")
}

func (d *RepoAddDialog) DisableInput(isDisabled bool) {
	d.RepoURL.SetDisabled(isDisabled)
	d.PasswordLineEdit.SetDisabled(isDisabled)
	d.SaveButton.SetDisabled(isDisabled)
}

func (d *RepoAddDialog) UseForExistingRepo() {
	d.EncryptionComboBox.Hide()
	d.EncryptionLabel.Hide()
	d.Title.SetText("Connect to existing Repository")
	d.SaveButton.DisconnectClicked()
	d.SaveButton.ConnectClicked(d.ProcessExistingRepo)
}

// TODO: combine similar parts with ProcessExistingRepo(). Move parts to processResult?
func (d *RepoAddDialog) ProcessNewRepo(_ bool) {
	err := d.Validate()
	if err != nil {
		d.ErrorText.SetText(err.Error())
		return
	}
	b, err := borg.NewInitRun(currentProfile, d.RepoURL.Text(), d.PasswordLineEdit.Text(),
		d.ExtraBorgArgumentsLineEdit.Text(), d.EncryptionComboBox.CurrentData(int(core.Qt__UserRole)).ToString())
	if err != nil {
		d.ErrorText.SetText(err.Error())
		return
	}
	d.DisableInput(true)
	go func() {
		err := b.Run()
		if err != nil {
			d.ErrorText.SetText(err.Error())
			d.DisableInput(false)
		} else {
			utils.Log.Info(b.Result)
			newRepo := models.Repo{
				Url:                d.RepoURL.Text(),
				Encryption:         sql.NullString{d.EncryptionComboBox.CurrentData(int(core.Qt__UserRole)).ToString(), true},
				ExtraBorgArguments: sql.NullString{d.ExtraBorgArgumentsLineEdit.Text(), true},
			}
			rows, err := models.DB.NamedExec(models.SqlNewRepo, newRepo)
			if err != nil {
				utils.Log.Error(err)
			}
			newRepoId, _ := rows.LastInsertId()
			newRepo.Id = int(newRepoId)
			currentRepo = &newRepo
			currentProfile.RepoId = sql.NullInt64{int64(currentRepo.Id), true}
			currentProfile.SaveField("repo_id")
			d.Accept()
		}
	}()
	d.ErrorText.SetText("Setting up new Repository...")
}

func (d *RepoAddDialog) ProcessExistingRepo(_ bool) {
	err := d.Validate()
	if err != nil {
		d.ErrorText.SetText(err.Error())
		return
	}
	b, err := borg.NewInfoRun(currentProfile, d.RepoURL.Text(), d.PasswordLineEdit.Text(), d.ExtraBorgArgumentsLineEdit.Text())
	if err != nil {
		d.ErrorText.SetText(err.Error())
		return
	}
	d.DisableInput(true)
	go func() {
		err := b.Run()
		if err != nil {
			d.ErrorText.SetText(err.Error())
			d.DisableInput(false)
		} else {
			utils.Log.Info(b.Result)
			newRepo := models.Repo{
				Url:                d.RepoURL.Text(),
				Encryption:         sql.NullString{b.Result.GetPath("encryption", "mode").MustString(), true},
				UniqueSize:         sql.NullInt64{b.Result.GetPath("cache", "stats", "unique_size").MustInt64(), true},
				UniqueCsize:        sql.NullInt64{b.Result.GetPath("cache", "stats", "unique_csize").MustInt64(), true},
				TotalSize:          sql.NullInt64{b.Result.GetPath("cache", "stats", "total_size").MustInt64(), true},
				TotalUniqueChunks:  sql.NullInt64{b.Result.GetPath("cache", "stats", "total_unique_chunks").MustInt64(), true},
				ExtraBorgArguments: sql.NullString{d.ExtraBorgArgumentsLineEdit.Text(), true},
			}
			rows, err := models.DB.NamedExec(models.SqlNewRepo, newRepo)
			if err != nil {
				utils.Log.Error(err)
			}

			newRepoId, err := rows.LastInsertId()
			newRepo.Id = int(newRepoId)
			currentRepo = &newRepo
			currentProfile.RepoId = sql.NullInt64{int64(currentRepo.Id), true}
			currentProfile.SaveField("repo_id")
			d.Accept()
		}
	}()
	d.ErrorText.SetText("Connecting to Repository...")
}
