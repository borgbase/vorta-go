package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"strconv"
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
		currentProfile.ScheduleIntervalHours, _ = strconv.Atoi(t.ScheduleIntervalHours.Text())
		currentProfile.ScheduleIntervalMinutes, _ = strconv.Atoi(t.ScheduleIntervalMinutes.Text())
		currentProfile.ScheduleFixedHour = t.ScheduleFixedTime.Time().Hour()
		currentProfile.ScheduleFixedMinute = t.ScheduleFixedTime.Time().Minute()
		models.DB.Save(currentProfile)

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

	t.WifiListWidget.DisconnectItemChanged()
	t.WifiListWidget.Clear()
	ww := []models.KnownWifi{}
	models.DB.Model(currentProfile).Related(&ww)
	for _, wifi := range ww {
		item := widgets.NewQListWidgetItem(t.WifiListWidget, 0)
		item.SetText(wifi.SSID)
		item.SetFlags(
				core.Qt__ItemIsUserCheckable |
				core.Qt__ItemIsEditable |
				core.Qt__ItemIsEnabled |
				core.Qt__ItemIsSelectable)
		item.SetData(int(core.Qt__UserRole), core.NewQVariant1(wifi.ID))
		if wifi.Allowed {
			item.SetCheckState(core.Qt__Checked)
		} else {
			item.SetCheckState(core.Qt__Unchecked)
		}
		t.WifiListWidget.AddItem2(item)
	}
	t.WifiListWidget.ConnectItemChanged(func(item *widgets.QListWidgetItem) {
		wifiToChange := models.KnownWifi{}
		models.DB.First(&wifiToChange, item.Data(int(core.Qt__UserRole)).ToInt(nil))
		wifiToChange.Allowed = item.CheckState() == core.Qt__Checked
		models.DB.Save(&wifiToChange)
	})
}

func (t *ScheduleTab) setNextBackupTime() {
	s := utils.Scheduler.NextTimeForProfile(currentProfile.ID)
	t.NextBackupDateTimeLabel.SetText(s)
}
