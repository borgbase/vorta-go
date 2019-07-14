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
	for _, repo := range models.QueryAllRepos() {
		t.RepoSelector.AddItem(repo.Url, core.NewQVariant1(repo.Id))
	}

	t.RepoSelector.ConnectCurrentIndexChanged(t.repoSelectorChanged)
}

func (t *RepoTab) repoSelectorChanged(newIndex int) {
	fmt.Println("Index Changed", newIndex)
}
