package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"vorta/models"
	"vorta/utils"
)

var schedulerRadioMap map[string]*widgets.QRadioButton

func (t *ScheduleTab) init() {
	t.ToolBox.SetCurrentIndex(0)

	schedulerRadioMap = map[string]*widgets.QRadioButton{
		"off":      t.ScheduleOffRadio,
		"interval": t.ScheduleIntervalRadio,
		"fixed":    t.ScheduleFixedRadio,
	}

	t.ScheduleApplyButton.ConnectClicked(func(_ bool) {
		utils.Log.Info("Applying new schedule.")
		for k, v := range schedulerRadioMap {
			if v.IsChecked() {
				currentProfile.ScheduleMode = k
				models.DB.Save(currentProfile)
				break
			}
		}
		utils.Scheduler.ReloadJobs()
		t.setNextBackupTime()
	})

	t.PreBackupCmdLineEdit.ConnectTextChanged(func(text string) {
		currentProfile.PreBackupCmd = text
		models.DB.Save(currentProfile)
	})
	t.PostBackupCmdLineEdit.ConnectTextChanged(func(text string) {
		currentProfile.PostBackupCmd = text
		models.DB.Save(currentProfile)
	})
}

func (t *ScheduleTab) Populate() {
	schedulerRadioMap[currentProfile.ScheduleMode].SetChecked(true)

	t.ScheduleIntervalHours.SetValue(currentProfile.ScheduleIntervalHours)
	t.ScheduleIntervalMinutes.SetValue(currentProfile.ScheduleIntervalMinutes)

	t.ScheduleFixedTime.SetTime(
		core.NewQTime3(currentProfile.ScheduleFixedHour, currentProfile.ScheduleFixedMinute, 0, 0))

	t.ValidationCheckBox.SetChecked(currentProfile.ValidationOn)
	t.ValidationSpinBox.SetValue(currentProfile.ValidationWeeks)
	t.PruneCheckBox.SetChecked(currentProfile.PruneOn)
	t.ValidationCheckBox.SetTristate(false)
	t.PruneCheckBox.SetTristate(false)

	t.PreBackupCmdLineEdit.SetText(currentProfile.PreBackupCmd)
	t.PostBackupCmdLineEdit.SetText(currentProfile.PostBackupCmd)

	t.setNextBackupTime()
}

func (t *ScheduleTab) setNextBackupTime() {
	s := utils.Scheduler.NextTimeForProfile(currentProfile.ID)
	t.NextBackupDateTimeLabel.SetText(s)
}
