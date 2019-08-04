package ui

import (
	"database/sql"
	"fmt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"vorta/models"
	"vorta/utils"
)

type MainWindowTabs struct {
	RepoTab     *RepoTab
	SourceTab   *SourceTab
	ScheduleTab *ScheduleTab
	ArchiveTab  *ArchiveTab
	MiscTab     *MiscTab
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
	w.ProfileAddButton.ConnectClicked(w.addProfile)
	w.ProfileDeleteButton.ConnectClicked(w.removeProfile)
	w.ProfileRenameButton.ConnectClicked(w.renameProfile)
	w.ConnectClose(func() bool { w.Close(); return true })
	w.Show()
}

func (w *MainWindow) AddTabs() {
	// Keep reference of tabs
	Tabs = MainWindowTabs{
		RepoTab:     NewRepoTab(w),
		SourceTab:   NewSourceTab(w),
		ScheduleTab: NewScheduleTab(w),
		ArchiveTab:  NewArchiveTab(w),
		MiscTab:     NewMiscTab(w),
	}
	w.TabWidget.AddTab(Tabs.RepoTab, "Repository")
	w.TabWidget.AddTab(Tabs.SourceTab, "Sources")
	w.TabWidget.AddTab(Tabs.ScheduleTab, "Schedule")
	w.TabWidget.AddTab(Tabs.ArchiveTab, "Archives")
	w.TabWidget.AddTab(Tabs.MiscTab, "Misc")

	w.refreshAllTabs()
}

func (w *MainWindow) profileSelectorChanged(ix int) {
	id := w.ProfileSelector.ItemData(ix, int(core.Qt__UserRole)).ToInt(nil)
	models.DB.Get(currentProfile, models.SqlProfileById, id)
	models.DB.Get(currentRepo, models.SqlRepoById, currentProfile.RepoId)
	utils.Log.Error(currentRepo.Url, id)
	w.refreshAllTabs()
}

func (w *MainWindow) refreshAllTabs() {
	Tabs.RepoTab.Populate()
	Tabs.SourceTab.Populate()
	Tabs.ScheduleTab.Populate()
	Tabs.ArchiveTab.Populate()
}

func (w *MainWindow) displayLogMessage(m string) {
	w.CreateProgressText.SetText(m)
	w.CreateProgressText.Repaint()
}

func (w *MainWindow) addProfile(_ bool) {
	dialog := NewProfileAddDialog(nil)
	dialog.SetParent2(w, core.Qt__Sheet)
	dialog.ConnectAccepted(func() {
		utils.Log.Info("profile added")
		rows, _ := models.DB.Exec(models.SqlNewProfile, dialog.ProfileNameField.Text())
		newProfileId, _ := rows.LastInsertId()
		models.DB.Get(currentProfile, models.SqlProfileById, newProfileId)
		w.ProfileSelector.AddItem(currentProfile.Name, core.NewQVariant1(currentProfile.Id))
		ix := w.ProfileSelector.FindData(core.NewQVariant1(currentProfile.Id), int(core.Qt__UserRole), core.Qt__MatchExactly)
		w.ProfileSelector.SetCurrentIndex(ix)
		currentRepo = &models.Repo{}
		w.refreshAllTabs()
	})
	dialog.Show()
}

func (w *MainWindow) renameProfile(_ bool) {
	dialog := NewProfileAddDialog(nil)
	dialog.SetParent2(w, core.Qt__Sheet)
	dialog.ModalTitle.SetText("Rename Profile")
	dialog.IntroTextLabel.Hide()
	dialog.AdjustSize()
	dialog.ProfileNameField.SetText(currentProfile.Name)
	dialog.ConnectAccepted(func() {
		utils.Log.Info("profile renamed")
		currentProfile.Name = dialog.ProfileNameField.Text()
		currentProfile.SaveField("name")
		ix := w.ProfileSelector.FindData(core.NewQVariant1(currentProfile.Id), int(core.Qt__UserRole), core.Qt__MatchExactly)
		w.ProfileSelector.SetItemText(ix, currentProfile.Name)
	})
	dialog.Show()
}

func (w *MainWindow) removeProfile(_ bool) {
	msgBox := widgets.QMessageBox_Question(nil, "Remove Profile",
		fmt.Sprintf("Are you sure you want to remove the profile %v?",
			currentProfile.Name), widgets.QMessageBox__Yes|widgets.QMessageBox__No, 0)

	if msgBox == widgets.QMessageBox__Yes {
		ix := w.ProfileSelector.FindData(core.NewQVariant1(currentProfile.Id), int(core.Qt__UserRole), core.Qt__MatchExactly)
		w.ProfileSelector.RemoveItem(ix)
		models.DB.MustExec(models.SqlRemoveProfileById, currentProfile.Id)
		var nProfiles int
		models.DB.Get(&nProfiles, models.SqlCountProfiles)
		if nProfiles == 0 {
			rows, _ := models.DB.Exec(models.SqlNewProfile, "Default")
			newID, _ := rows.LastInsertId()
			w.ProfileSelector.AddItem("Default", core.NewQVariant1(newID))
		}
	}
}

func (w *MainWindow) RunUIEventHandler(appChan chan utils.VEvent) {
	for e := range MainWindowChan {
		switch e.Topic {
		case "StatusUpdate": // TODO: Use enums here.
			w.displayLogMessage(e.Message)
		case "ChangeRepo":
			utils.Log.Info("Repo changed")
			err := models.DB.Get(currentRepo, models.SqlRepoById, e.Message)
			utils.Log.Info("currentRepo val:", currentRepo)
			if err == nil {
				currentProfile.RepoId = sql.NullInt64{int64(currentRepo.Id), true}
				models.DB.NamedExec(fmt.Sprintf(models.SqlUpdateProfileFieldById, "repo_id"), currentProfile)
			}
			w.refreshAllTabs()
		case "StartBackup":
			appChan <- e
		case "OpenMainWindow":
			w.Show()
			w.Raise()
			w.ActivateWindow()
		case "UpdateArchiveTab":
			Tabs.ArchiveTab.Populate()
		default:
			utils.Log.Info("Unhandled UI Channel Event")
		}

	}
}

func ChooseFileDialog(callback func(files []string)) {
	fd := widgets.NewQFileDialog(nil, 0)
	fd.SetFileMode(widgets.QFileDialog__Directory)
	fd.SetWindowModality(core.Qt__WindowModal) //TODO: not working on macOS?
	fd.ConnectFilesSelected(callback)
	fd.Exec() //TODO: what happens if user cancels?
}
