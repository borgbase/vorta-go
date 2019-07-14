package ui

import (
	"vorta-go/models"
)

func (t *RepoTab) init() {
	models.QueryAllRepos()
}
