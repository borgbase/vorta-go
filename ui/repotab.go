package ui

import (
	"database/sql"
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"os"
	"path"
	"vorta/models"
	"vorta/utils"
)

var availableCompression = map[string]string{
	"LZ4 (modern, default)":       "lz4",
	"Zstandard Level 3 (modern)":  "zstd,3",
	"Zstandard Level 8 (modern)":  "zstd,8",
	"ZLIB Level 6 (auto, legacy)": "auto,zlib,6",
	"LZMA Level 6 (auto, legacy)": "auto,lzma,6",
	"No Compression":              "none",
}

func (t *RepoTab) init() {
	// Populate Compression modes
	for desc, value := range availableCompression {
		t.RepoCompression.AddItem(desc, core.NewQVariant1(value))
	}
	t.RepoCompression.ConnectCurrentIndexChanged(t.compressionSelectorChanged)

	// Populate available Repos
	t.RepoSelector.AddItem("+ Initialize New Repository", core.NewQVariant1("new"))
	t.RepoSelector.AddItem("+ Add Existing Repository", core.NewQVariant1("existing"))
	t.RepoSelector.InsertSeparator(3)
	t.RepoSelector.ConnectCurrentIndexChanged(t.repoSelectorChanged)

	t.RepoRemoveToolbutton.ConnectClicked(t.unlinkRepo)

	t.SshComboBox.AddItem("Automatically choose SSH Key (default)", core.NewQVariant1(nil))
	t.SshComboBox.AddItem("Create new SSH Key", core.NewQVariant1("new"))
	t.SshComboBox.ConnectCurrentIndexChanged(t.sshSelectorChanged)

	t.SshKeyToClipboardButton.ConnectClicked(t.sshCopyToClipboard)
}

func (t *RepoTab) compressionSelectorChanged(ix int) {
	currentProfile.Compression = t.RepoCompression.ItemData(ix, int(core.Qt__UserRole)).ToString()
	models.DB.Save(currentProfile)
	//sql := fmt.Sprintf(models.SqlUpdateProfileFieldById, "compression")
	//_, err := models.DB.NamedExec(sql, currentProfile)
	//if err != nil {
	//	utils.Log.Error(err)
	//}
}

func (t *RepoTab) unlinkRepo(_ bool) {
	if currentRepo.ID != t.RepoSelector.CurrentData(int(core.Qt__UserRole)).ToInt(nil) {
		utils.Log.Panic("Not sure which repo to unlink.")
	}
	msgBox := widgets.QMessageBox_Question(nil, "Unlink Repo",
		fmt.Sprintf("Are you sure you want to unlinke the repo %v? This will not remove any data and you can always re-add the repo later.",
			currentRepo.Url), widgets.QMessageBox__Yes|widgets.QMessageBox__No, 0)

	if msgBox == widgets.QMessageBox__Yes {
		t.RepoSelector.DisconnectCurrentIndexChanged()
		utils.Log.Info("Unlinking repo %v", currentRepo.Url)
		//currentRepoId := currentRepo.ID
		currentProfile.RepoID = sql.NullInt64{Valid: false}
		//currentProfile.SaveField("repo_id")
		models.DB.Save(&currentProfile)
		//models.DB.MustExec(models.SqlRemoveRepoById, currentRepoId)
		models.DB.Delete(currentRepo)
		currentRepo = &models.Repo{}
		t.RepoSelector.RemoveItem(t.RepoSelector.CurrentIndex())
		t.RepoSelector.SetCurrentIndex(0)
		t.RepoSelector.ConnectCurrentIndexChanged(t.repoSelectorChanged)
	}
}

func (t *RepoTab) setCompression() {
	ix := t.RepoCompression.FindData(core.NewQVariant1(currentProfile.Compression), int(core.Qt__UserRole), core.Qt__MatchExactly)
	t.RepoCompression.SetCurrentIndex(ix)
}

func (t *RepoTab) setStats() {
	// Populate Repo Stats
	t.RepoEncryption.SetText(currentRepo.Encryption.String)
	t.SizeCompressed.SetText(humanize.Bytes(uint64(currentRepo.UniqueCsize.Int64)))
	t.SizeDeduplicated.SetText(humanize.Bytes(uint64(currentRepo.UniqueSize.Int64)))
	t.SizeOriginal.SetText(humanize.Bytes(uint64(currentRepo.TotalSize.Int64)))
}

func (t *RepoTab) Populate() {
	t.RepoSelector.DisconnectCurrentIndexChanged()
	t.RepoCompression.DisconnectCurrentIndexChanged()
	rr := []models.Repo{}
	models.DB.Find(&rr)
	for _, repo := range rr {
		// see if repo already exists, otherwise add it.
		existingIx := t.RepoSelector.FindData(core.NewQVariant1(repo.ID), int(core.Qt__UserRole), core.Qt__MatchExactly)
		if existingIx == -1 {
			t.RepoSelector.AddItem(repo.Url, core.NewQVariant1(repo.ID))
		}
	}
	ix := t.RepoSelector.FindData(core.NewQVariant1(currentRepo.ID), int(core.Qt__UserRole), core.Qt__MatchExactly)
	if ix < 0 {
		ix = 0 // if currentRepo is empty, set to first row.
	}
	t.RepoSelector.SetCurrentIndex(ix)

	for i := t.SshComboBox.Count(); i > 1; i-- {
		t.SshComboBox.RemoveItem(i)
	}
	localSshKeys, err := utils.FindSSHKeysInStandardFolder()
	if err != nil {
		utils.Log.Errorf("Error reading users SSH keys in ~/.ssh: %v", err)
	}
	if len(localSshKeys) > 0 {
		t.SshComboBox.InsertSeparator(2)
		for _, key := range localSshKeys {
			t.SshComboBox.AddItem(key, core.NewQVariant1(key))
		}
	}

	t.setStats()
	t.setCompression()
	t.RepoCompression.ConnectCurrentIndexChanged(t.compressionSelectorChanged)
	t.RepoSelector.ConnectCurrentIndexChanged(t.repoSelectorChanged)
}

func (t *RepoTab) sshSelectorChanged(index int) {
	switch index {
	case 0:
		currentProfile.SSHKey = sql.NullString{Valid: false}
		models.DB.Save(currentProfile)
	case 1:
		dialog := NewSshAddDialog(t)
		dialog.SetParent2(t, core.Qt__Sheet)
		dialog.ConnectRejected(func() {
			utils.Log.Info("SSH Dialog closed")
			t.Populate()
			ix := t.SshComboBox.FindData(core.NewQVariant1(dialog.OutputFileTextBox.Text()), int(core.Qt__UserRole), core.Qt__MatchExactly)
			t.SshComboBox.SetCurrentIndex(ix)
		})
		dialog.Show()
	}
}

func (t *RepoTab) sshCopyToClipboard(_ bool) {
	keyName := t.SshComboBox.CurrentData(int(core.Qt__UserRole)).ToString()
	sshDir, _ := utils.GetSSHDir()
	keyFullPath := path.Join(sshDir, keyName+".pub")
	if _, err := os.Stat(keyFullPath); err != nil {
		widgets.QMessageBox_Critical(nil,
			"Public key not found.", fmt.Sprintf("Didn't find the public key for %v.", keyFullPath),
			widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
		return
	}
	utils.CopyPublicKeyToClipboard(keyFullPath)
	widgets.QMessageBox_Information(nil,
		"Public key copied to clipboard", fmt.Sprintf("The public key part of %v was copied to the clipboard.", keyFullPath),
		widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
}

func (t *RepoTab) repoSelectorChanged(index int) {
	if index == 0 {
		return
	}
	itemData := t.RepoSelector.ItemData(index, int(core.Qt__UserRole)).ToString()
	switch itemData {
	case "new":
		dialog := NewRepoAddDialog(t)
		dialog.SetParent2(t, core.Qt__Sheet)
		dialog.ConnectAccepted(func() {
			utils.Log.Info("New repo added.")
			MainWindowChan <- utils.VEvent{Topic: "ChangeRepo", Message: string(currentRepo.ID)}
		})
		dialog.ConnectRejected(func() {
			utils.Log.Info("Dialog Rejected")
		})
		dialog.Show()
	case "existing":
		dialog := NewRepoAddDialog(t)
		dialog.UseForExistingRepo()
		dialog.SetParent2(t, core.Qt__Sheet)
		dialog.ConnectAccepted(func() {
			utils.Log.Info("Existing repo added.")
			MainWindowChan <- utils.VEvent{Topic: "ChangeRepo", Message: string(currentRepo.ID)}

		})
		dialog.ConnectRejected(func() {
			utils.Log.Info("Dialog Rejected")
		})
		dialog.Show()
	default:
		MainWindowChan <- utils.VEvent{Topic: "ChangeRepo", Message: itemData}
	}
}
