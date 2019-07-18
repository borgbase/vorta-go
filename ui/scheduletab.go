package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"vorta-go/utils"
)

var schedulerRadioMap map[string]*widgets.QRadioButton

func (t *ScheduleTab) init() {
	t.ToolBox.SetCurrentIndex(0)

	t.ScheduleApplyButton.ConnectClicked(func(_ bool) {
		utils.Log.Info("Applying new schedule.")
		for k, v := range schedulerRadioMap {
			if v.IsChecked() {
				currentProfile.
			}
		}

		utils.Scheduler.ReloadJobs()
	})

	t.PreBackupCmdLineEdit.ConnectTextChanged(func(text string) {

	})
	t.PostBackupCmdLineEdit.ConnectTextChanged(func(text string) {

	})

	schedulerRadioMap = map[string]*widgets.QRadioButton{
		"off": t.ScheduleOffRadio,
		"interval": t.ScheduleIntervalRadio,
		"fixed": t.ScheduleFixedRadio,
	}
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
	s := utils.Scheduler.NextTimeForProfile(currentProfile.Id)
	t.NextBackupDateTimeLabel.SetText(s)
}
