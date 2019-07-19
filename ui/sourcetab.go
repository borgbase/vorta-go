package ui

import (
	"vorta-go/models"
	"vorta-go/utils"
)

func (t *SourceTab) init() {
	t.ExcludePatternsField.ConnectTextChanged(func() {
		currentProfile.ExcludePatterns.String = t.ExcludePatternsField.ToPlainText()
		currentProfile.ExcludePatterns.Valid = true
		currentProfile.UpdateField("exclude_patterns")
	})

	t.ExcludeIfPresentField.ConnectTextChanged(func() {
		currentProfile.ExcludeIfPresent.String = t.ExcludeIfPresentField.ToPlainText()
		currentProfile.ExcludeIfPresent.Valid = true
		currentProfile.UpdateField("exclude_if_present")
	})

	t.SourceAddFile.ConnectClicked(func(_ bool) {
		utils.Log.Info("Add file triggered.")
		ChooseFileDialog(func(files []string) {
			utils.Log.Info(files)
			var nExisting int
			models.DB.Get(&nExisting, models.SqlCountSources, currentProfile.Id, files[0])
			if nExisting == 0 {
				t.SourceFilesWidget.AddItem(files[0])
				models.DB.MustExec(models.SqlInsertSourceDir, files[0], currentProfile.Id)
			}
		})
	})
	
	t.SourceRemove.ConnectClicked(func(_ bool) {
		item := t.SourceFilesWidget.TakeItem(t.SourceFilesWidget.CurrentRow())
		utils.Log.Info(item.Text())
		models.DB.MustExec(models.SqlDeleteSourceDir, currentProfile.Id, item.Text())
	})
}

func (t *SourceTab) Populate() {
	t.ExcludeIfPresentField.Clear()
	t.ExcludePatternsField.Clear()
	for i := t.SourceFilesWidget.Count(); i >= 0; i-- {  // Clear() didn't work.
		t.SourceFilesWidget.TakeItem(i)
	}

	ss := []models.SourceDir{}
	models.DB.Select(&ss, models.SqlAllSourcesByProfileId, currentProfile.Id)
	for _, s := range ss {
		t.SourceFilesWidget.AddItem(s.Dir)
	}
	t.ExcludePatternsField.AppendPlainText(currentProfile.ExcludePatterns.String)
	t.ExcludeIfPresentField.AppendPlainText(currentProfile.ExcludeIfPresent.String)

}


