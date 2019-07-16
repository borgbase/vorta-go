package ui

import (
	"github.com/therecipe/qt/core"
	"vorta-go/app"
	"vorta-go/borg"
	"vorta-go/models"
)

func (w *MainWindow) init() {
	w.TabWidget.AddTab(NewRepoTab(w), "Repository")
	w.TabWidget.AddTab(NewSourceTab(w), "Sources")
	w.TabWidget.AddTab(NewScheduleTab(w), "Schedule")
	w.TabWidget.AddTab(NewArchiveTab(w), "Archives")
	w.TabWidget.AddTab(NewMiscTab(w), "Misc")
	w.SetWindowTitle("Vorta for Borg Backup")

	//# Init profile list
	//for profile in BackupProfileModel.select():
	//	self.profileSelector.addItem(profile.name, profile.id)
	//	self.profileSelector.setCurrentIndex(0)
	//	self.profileSelector.currentIndexChanged.connect(self.profile_select_action)
	//	self.profileRenameButton.clicked.connect(self.profile_rename_action)
	//	self.profileDeleteButton.clicked.connect(self.profile_delete_action)
	//	self.profileAddButton.clicked.connect(self.profile_add_action)

	pp := []models.Profile{}
	models.DB.Select(&pp, models.SqlAllProfiles)
	for _, profile := range pp {
		w.ProfileSelector.AddItem(profile.Name, core.NewQVariant1(profile.Id))
	}

	if app.CurrentProfile == nil {
		app.CurrentProfile = &pp[0]
	}

	w.CreateStartBtn.ConnectClicked(w.StartBackup)

	go w.displayLogMessages(app.StatusUpdateChannel)
	w.Show()
}

func (w *MainWindow) StartBackup(checked bool) {
	app.StatusUpdateChannel <- "Running Borg"
	b := borg.BorgRun{SubCommand: "info"}
	b.Prepare()
	go b.Run()
}

func (w *MainWindow) displayLogMessages(c chan string) {
	for updateStr := range c {
		w.CreateProgressText.SetText(updateStr)
		w.CreateProgressText.Repaint()
	}

}
