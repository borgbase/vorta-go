package ui

import (
	"fmt"
	"github.com/therecipe/qt/core"
	"vorta-go/models"
)

func (t *RepoTab) init() {
	t.RepoSelector.AddItem("+ Initialize New Repository", core.NewQVariant1("new"))
	t.RepoSelector.AddItem("+ Add Existing Repository", core.NewQVariant1("existing"))
	t.RepoSelector.InsertSeparator(3)
	rr := []models.Repo{}
	models.DB.Select(&rr, models.SqlAllRepos)
	for _, repo := range rr {
		t.RepoSelector.AddItem(repo.Url, core.NewQVariant1(repo.Id))
	}

	t.RepoSelector.ConnectCurrentIndexChanged(t.repoSelectorChanged)
}

func (t *RepoTab) repoSelectorChanged(newIndex int) {
	fmt.Println("Index Changed", newIndex)
}
