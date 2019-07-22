package ui

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/therecipe/qt/core"
	"vorta-go/models"
	"vorta-go/utils"
)

var availableCompression = map[string]string{
	"LZ4 (modern, default)": "lz4",
	"Zstandard Level 3 (modern)": "zstd,3",
	"Zstandard Level 8 (modern)": "zstd,8",
	"ZLIB Level 6 (auto, legacy)": "auto,zlib,6",
	"LZMA Level 6 (auto, legacy)": "auto,lzma,6",
	"No Compression": "none",
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

}

func (t *RepoTab) compressionSelectorChanged(ix int) {
	currentProfile.Compression = t.RepoCompression.ItemData(ix, int(core.Qt__UserRole)).ToString()
	sql := fmt.Sprintf(models.SqlUpdateProfileFieldById, "compression")
	_, err := models.DB.NamedExec(sql, currentProfile)
	if err != nil {
		utils.Log.Error(err)
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
	rr := []models.Repo{}
	models.DB.Select(&rr, models.SqlAllRepos)
	for _, repo := range rr {
		// see if repo already exists, otherwise add it.
		existingIx := t.RepoSelector.FindData(core.NewQVariant1(repo.Id), int(core.Qt__UserRole), core.Qt__MatchExactly)
		if existingIx == -1 {
			t.RepoSelector.AddItem(repo.Url, core.NewQVariant1(repo.Id))
		}
	}
	ix := t.RepoSelector.FindData(core.NewQVariant1(currentRepo.Id), int(core.Qt__UserRole), core.Qt__MatchExactly)
	t.RepoSelector.SetCurrentIndex(ix)

	t.setStats()
	t.setCompression()
}

func (t *RepoTab) repoSelectorChanged(index int) {
	itemData := t.RepoSelector.ItemData(index, int(core.Qt__UserRole)).ToString()
	if index == 0 {
		return
	} else if itemData == "new" {
		dialog := NewRepoAddDialog(t)
		dialog.SetParent2(t, core.Qt__Sheet)
		dialog.ConnectAccepted(func(){
			utils.Log.Info("New repo added.")
			MainWindowChan <- utils.VEvent{Topic: "ChangeRepo", Message: string(currentRepo.Id)}
		})
		dialog.ConnectRejected(func(){
			utils.Log.Info("Dialog Rejected")
		})
		dialog.Show()
	} else if itemData == "existing" {
		dialog := NewRepoAddDialog(t)
		dialog.UseForExistingRepo()
		dialog.SetParent2(t, core.Qt__Sheet)
		dialog.ConnectAccepted(func(){
			utils.Log.Info("Existing repo added.")
			MainWindowChan <- utils.VEvent{Topic: "ChangeRepo", Message: string(currentRepo.Id)}

		})
		dialog.ConnectRejected(func(){
			utils.Log.Info("Dialog Rejected")
		})
		dialog.Show()
	} else {
		MainWindowChan <- utils.VEvent{Topic: "ChangeRepo", Message: itemData}
	}
}
