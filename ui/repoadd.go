package ui

import (
	"database/sql"
	"github.com/therecipe/qt/core"
	"vorta-go/borg"
	"vorta-go/models"
	"vorta-go/utils"
)

var encryptionModes = map[string]string{
	"repokey-blake2": "Repokey-Blake2 (Recommended, key stored in repository)",
	"repokey": "Repokey",
	"keyfile-blake2": "Keyfile-Blake2 (Key stored in home directory)",
	"keyfile": "Keyfile",
	"none": "None (not recommended)",
}

func (d *RepoAddDialog) init() {
	/*
	    def __init__(self, parent=None):
	        super().__init__(parent)
	        self.result = None
	        self.is_remote_repo = True
	        self.chooseLocalFolderButton.clicked.connect(self.choose_local_backup_folder)
	        self.useRemoteRepoButton.clicked.connect(self.use_remote_repo_action)
	*/
	for k, v := range encryptionModes {
		d.EncryptionComboBox.AddItem(v, core.NewQVariant1(k))
	}
	d.SaveButton.ConnectClicked(func(_ bool) {
		d.Accept()
	})

	d.CloseButton.ConnectClicked(func(_ bool) {
		d.Close()
	})

	d.TabWidget.SetCurrentIndex(0)
}

func (d *RepoAddDialog) Validate() {
	// Check if URL exists.
/*
   if self.is_remote_repo and not re.match(r'.+:.+', self.values['repo_url']):
       self._set_status(self.tr('Please enter a valid repo URL or select a local path.'))
       return False

   if self.__class__ == AddRepoWindow:
       if self.values['encryption'] != 'none':
           if len(self.values['password']) < 8:
               self._set_status(self.tr('Please use a longer password.'))
               return False
 */
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

func (d *RepoAddDialog) ProcessExistingRepo(_ bool) {
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
				Url: d.RepoURL.Text(),
				Encryption: sql.NullString{b.Result.GetPath("encryption", "mode").MustString(), true},
				UniqueSize: sql.NullInt64{b.Result.GetPath("cache", "stats", "unique_size").MustInt64(), true},
				UniqueCsize: sql.NullInt64{b.Result.GetPath("cache", "stats", "unique_csize").MustInt64(), true},
				TotalSize: sql.NullInt64{b.Result.GetPath("cache", "stats", "total_size").MustInt64(), true},
				TotalUniqueChunks: sql.NullInt64{b.Result.GetPath("cache", "stats", "total_unique_chunks").MustInt64(), true},
				ExtraBorgArguments: sql.NullString{d.ExtraBorgArgumentsLineEdit.Text(), true},
			}
			rows, err := models.DB.NamedExec(models.SqlNewRepo, newRepo)
			if err != nil {
				utils.Log.Error(err)
			}

			newRepoId, err := rows.LastInsertId()
			newRepo.Id = int(newRepoId)
			currentRepo = &newRepo
			currentProfile.RepoId = currentRepo.Id
			currentProfile.UpdateField("repo_id")
			MainWindowChan <- utils.VEvent{Topic: "ChangeRepo", Message: string(currentRepo.Id)}
			d.Accept()
		}
	}()
	d.ErrorText.SetText("Connecting to Repository...")
}
