package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/dustin/go-humanize"
	"vorta-go/app"
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

var currentRepo *models.Repo

func (t *RepoTab) init() {

	// Populate available Repos
	t.RepoSelector.AddItem("+ Initialize New Repository", core.NewQVariant1("new"))
	t.RepoSelector.AddItem("+ Add Existing Repository", core.NewQVariant1("existing"))
	t.RepoSelector.InsertSeparator(3)
	rr := []models.Repo{}
	models.DB.Select(&rr, models.SqlAllRepos)
	for _, repo := range rr {
		t.RepoSelector.AddItem(repo.Url, core.NewQVariant1(repo.Id))
	}
	t.RepoSelector.ConnectCurrentIndexChanged(t.repoSelectorChanged)

	// Populate Compression modes
	for desc, value := range availableCompression {
		t.RepoCompression.AddItem(desc, core.NewQVariant1(value))
	}
	t.RepoCompression.ConnectCurrentIndexChanged(t.compressionSelectorChanged)
}

func (t *RepoTab) compressionSelectorChanged(newIndex int) {

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

func (t *RepoTab) Update() {
	ix := t.RepoSelector.FindData(core.NewQVariant1(currentRepo.Id), int(core.Qt__UserRole), core.Qt__MatchExactly)
	t.RepoSelector.SetCurrentIndex(ix)
}

func (t *RepoTab) repoSelectorChanged(newIndex int) {
	// Save new repo to profile
	// Set global currentRepo
	// Repaint stuff in other tabs.
	app.Log.Info("Repo changed.")
	MainWindowChan <- utils.VEvent{Topic: "ChangeRepo", Data: "ii"}
	//t.setStats()
	//t.setCompression()
}
