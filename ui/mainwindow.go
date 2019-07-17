package ui

import (
	"github.com/therecipe/qt/core"
	"vorta-go/app"
	"vorta-go/borg"
	"vorta-go/models"
	"vorta-go/utils"
)

type MainWindowTabs struct {
	RepoTab *RepoTab
	SourceTab *SourceTab
	ScheduleTab *ScheduleTab
	ArchiveTab *ArchiveTab
	MiscTab *MiscTab
}

var (
	currentProfile *models.Profile
	Tabs           MainWindowTabs
	MainWindowChan chan utils.VEvent
)

func (w *MainWindow) init() {
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
	// Set currentProfile and currentRepo
	currentProfile = &pp[0]
	models.DB.Get(currentRepo, models.SqlRepoById, pp[0].RepoId)

	w.CreateStartBtn.ConnectClicked(w.StartBackup)
	w.ProfileSelector.ConnectCurrentIndexChanged(w.profileSelectorChanged)
	w.Show()
	MainWindowChan = make(chan utils.VEvent)
}

func (w *MainWindow) AddTabs() {
	// Keep reference of tabs
	Tabs = MainWindowTabs{
		RepoTab: NewRepoTab(w),
		SourceTab: NewSourceTab(w),
		ScheduleTab: NewScheduleTab(w),
		ArchiveTab: NewArchiveTab(w),
		MiscTab: NewMiscTab(w),
	}
	w.TabWidget.AddTab(Tabs.RepoTab, "Repository")
	w.TabWidget.AddTab(Tabs.SourceTab, "Sources")
	w.TabWidget.AddTab(Tabs.ScheduleTab, "Schedule")
	w.TabWidget.AddTab(Tabs.ArchiveTab, "Archives")
	w.TabWidget.AddTab(Tabs.MiscTab, "Misc")
}

func (w *MainWindow) profileSelectorChanged(ix int) {
	id := w.ProfileSelector.ItemData(ix, int(core.Qt__UserRole))
	models.DB.Get(currentProfile, models.SqlProfileById, id)
	models.DB.Get(currentRepo, models.SqlRepoById, currentProfile.RepoId)
	Tabs.RepoTab.Update()
}

func (w *MainWindow) StartBackup(checked bool) {
	MainWindowChan <- utils.VEvent{Topic: "StatusUpdate", Data: "Started Backup"}

	b, err := borg.NewInfoRun(currentProfile, currentRepo)
	if err != nil {
		app.Log.Error(err)
	}
	go b.Run()
}

func (w *MainWindow) displayLogMessage(m string) {
		w.CreateProgressText.SetText(m)
		w.CreateProgressText.Repaint()
}

func (w *MainWindow) RunUIEventHandler(appChan chan utils.VEvent) {
	for e := range MainWindowChan {
		app.Log.Info("Processing event")
		switch e.Topic {
		case "StatusUpdate":
			w.displayLogMessage(e.Data)
		case "ChangeRepo":
			app.Log.Info("Repo changed")
		case "EventForApp":
			appChan <- e
		default:
			app.Log.Info("Unhandled UI Event")
		}

	}
}
