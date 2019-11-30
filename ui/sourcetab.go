package ui

import (
	"time"
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
				models.DB.Model(&currentProfile).Association("SourceDirs").Append(
					models.SourceDir{Dir: file, AddedAt: time.Now()})
			}
		})
	})

	t.SourceRemove.ConnectClicked(func(_ bool) {
		item := t.SourceFilesWidget.TakeItem(t.SourceFilesWidget.CurrentRow())
		models.DB.Where(&models.SourceDir{Dir: item.Text(), ProfileId: currentProfile.ID}).Delete(&models.SourceDir{})
	})
}

func (t *SourceTab) Populate() {
	//t.ExcludeIfPresentField.DisconnectTextChanged()
	//t.ExcludePatternsField.DisconnectTextChanged()
	t.ExcludePatternsField.SetDisabled(true)
	t.ExcludeIfPresentField.SetDisabled(true)
	t.ExcludeIfPresentField.Clear()
	t.ExcludePatternsField.Clear()
	for i := t.SourceFilesWidget.Count(); i >= 0; i-- { // Clear() didn't work.
		t.SourceFilesWidget.TakeItem(0)
	}

	ss := []models.SourceDir{}
	models.DB.Model(&currentProfile).Related(&ss)

	for _, s := range ss {
		t.SourceFilesWidget.AddItem(s.Dir)
		utils.Log.Info("Adding source item", s.Dir, s.ProfileId, s.ID)
	}
	t.SourceFilesWidget.Repaint()
	t.ExcludePatternsField.AppendPlainText(currentProfile.ExcludePatterns.String)
	t.ExcludeIfPresentField.AppendPlainText(currentProfile.ExcludeIfPresent.String)

	//t.ExcludePatternsField.ConnectTextChanged(t.saveExcludes)
	//t.ExcludeIfPresentField.ConnectTextChanged(t.saveExcludes)
	t.ExcludePatternsField.SetDisabled(false)
	t.ExcludeIfPresentField.SetDisabled(false)
}

func (t *SourceTab) saveExcludes() {
	currentProfile.ExcludePatterns.String = t.ExcludePatternsField.ToPlainText()
	currentProfile.ExcludePatterns.Valid = true

	currentProfile.ExcludeIfPresent.String = t.ExcludeIfPresentField.ToPlainText()
	currentProfile.ExcludeIfPresent.Valid = true
	models.DB.Save(currentProfile)
}
