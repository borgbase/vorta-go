package ui

import (
	"github.com/therecipe/qt/core"
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
	currentRepo    *models.Repo
)

func (w *MainWindow) init() {
	w.SetWindowTitle("Vorta for Borg Backup")
	MainWindowChan = make(chan utils.VEvent)

	pp := []models.Profile{}
	models.DB.Select(&pp, models.SqlAllProfiles)
	for _, profile := range pp {
		w.ProfileSelector.AddItem(profile.Name, core.NewQVariant1(profile.Id))
	}

	// Set currentProfile and currentRepo
	currentProfile = &pp[0]
	currentRepo = &models.Repo{}
	models.DB.Get(currentRepo, models.SqlRepoById, pp[0].RepoId)

	w.CreateStartBtn.ConnectClicked(func(_ bool) {
		utils.Log.Info("clicked start-backup")
		MainWindowChan <- utils.VEvent{Topic: "StartBackup", Profile: currentProfile}
	})
	w.ProfileSelector.ConnectCurrentIndexChanged(w.profileSelectorChanged)
	w.ConnectClose(func() bool {w.Close(); return true})
	w.Show()
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

	Tabs.RepoTab.Populate()
}

func (w *MainWindow) profileSelectorChanged(ix int) {
	id := w.ProfileSelector.ItemData(ix, int(core.Qt__UserRole)).ToInt(nil)
	models.DB.Get(currentProfile, models.SqlProfileById, id)
	models.DB.Get(currentRepo, models.SqlRepoById, currentProfile.RepoId)
	utils.Log.Error(currentRepo.Url, id)
	Tabs.RepoTab.Populate()
}

func (w *MainWindow) displayLogMessage(m string) {
		w.CreateProgressText.SetText(m)
		w.CreateProgressText.Repaint()
}

func (w *MainWindow) RunUIEventHandler(appChan chan utils.VEvent) {
	for e := range MainWindowChan {
		switch e.Topic {
		case "StatusUpdate":
			w.displayLogMessage(e.Message)
		case "ChangeRepo":
			utils.Log.Info("Repo changed")
			models.DB.Get(currentRepo, models.SqlRepoById, e.Message)
			Tabs.RepoTab.Populate()
			// TODO: Save changed repoId to profile
		case "StartBackup":
			appChan <- e
		case "OpenMainWindow":
			w.Show()
		default:
			utils.Log.Info("Unhandled UI Channel Event")
		}

	}
}
