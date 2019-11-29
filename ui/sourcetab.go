package ui

import (
	"vorta/models"
	"vorta/utils"
)

func (t *SourceTab) init() {
	t.SourceAddFile.ConnectClicked(func(_ bool) {
		utils.Log.Info("Add file triggered.")
		ChooseFileDialog(func(files []string) {
			utils.Log.Info(files)
			for _, file := range files {
				t.SourceFilesWidget.AddItem(file)
				models.DB.Model(&currentProfile).Association("SourceDirs").Append(models.SourceDir{Dir: file})
			}
			//var nExisting int
			//models.DB.Get(&nExisting, models.SqlCountSources, currentProfile.Id, files[0])
			//models.DB.Model(&currentProfile).Association("SourceDirs").Count()
			//
			//if nExisting == 0 {
			//	t.SourceFilesWidget.AddItem(files[0])
			//	models.DB.MustExec(models.SqlInsertSourceDir, files[0], currentProfile.Id)
			//}
		})
	})

	t.SourceRemove.ConnectClicked(func(_ bool) {
		item := t.SourceFilesWidget.TakeItem(t.SourceFilesWidget.CurrentRow())
		models.DB.Model(&currentProfile).Association("SourceDirs").Delete(models.SourceDir{Dir: item.Text()})
		utils.Log.Info(item.Text())
	})
}

func (t *SourceTab) Populate() {
	t.ExcludeIfPresentField.DisconnectTextChanged()
	t.ExcludePatternsField.DisconnectTextChanged()
	t.ExcludeIfPresentField.Clear()
	t.ExcludePatternsField.Clear()
	for i := t.SourceFilesWidget.Count(); i >= 0; i-- { // Clear() didn't work.
		t.SourceFilesWidget.TakeItem(i)
	}

	ss := []models.SourceDir{}
	models.DB.Model(&currentProfile).Related(&ss)

	for _, s := range ss {
		t.SourceFilesWidget.AddItem(s.Dir)
	}
	t.SourceFilesWidget.Repaint()
	t.ExcludePatternsField.AppendPlainText(currentProfile.ExcludePatterns.String)
	t.ExcludeIfPresentField.AppendPlainText(currentProfile.ExcludeIfPresent.String)

	t.ExcludePatternsField.ConnectTextChanged(t.saveExcludes)
	t.ExcludeIfPresentField.ConnectTextChanged(t.saveExcludes)
}

func (t *SourceTab) saveExcludes() {
	currentProfile.ExcludePatterns.String = t.ExcludePatternsField.ToPlainText()
	currentProfile.ExcludePatterns.Valid = true

	currentProfile.ExcludeIfPresent.String = t.ExcludeIfPresentField.ToPlainText()
	currentProfile.ExcludeIfPresent.Valid = true
	models.DB.Save(currentProfile)
}
