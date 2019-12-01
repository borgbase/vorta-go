package ui

import (
	"fmt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"vorta/app"
	"vorta/borg"
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
	models.DB.Find(&pp)
	for _, profile := range pp {
		w.ProfileSelector.AddItem(profile.Name, core.NewQVariant1(profile.ID))
	}

	// Set currentProfile and currentRepo
	currentProfile = &pp[0]
	currentRepo = &models.Repo{}
	models.DB.Model(&currentProfile).Related(&currentRepo)

	w.CreateStartBtn.ConnectClicked(func(_ bool) {
		utils.Log.Info("clicked start-backup")
		MainWindowChan <- utils.VEvent{Topic: "StartBackup", Profile: currentProfile}
	})
	w.ProfileSelector.ConnectCurrentIndexChanged(w.profileSelectorChanged)
	w.ProfileAddButton.ConnectClicked(w.addProfile)
	w.ProfileDeleteButton.ConnectClicked(w.removeProfile)
	w.ProfileRenameButton.ConnectClicked(w.renameProfile)
	w.ConnectClose(func() bool { w.Close(); return true })
	w.CancelButton.ConnectClicked(func(_ bool) {
		app.AppChan <- utils.VEvent{Topic: "CancelBorgRun"}
	})
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
	utils.Log.Info("Running profileselctorchanged")
	id := w.ProfileSelector.ItemData(ix, int(core.Qt__UserRole)).ToInt(nil)
	currentProfile = &models.Profile{ID: id}
	models.DB.Take(&currentProfile)
	currentRepo = &models.Repo{}
	models.DB.Model(&currentProfile).Related(&currentRepo)
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
		currentProfile = models.NewProfile(dialog.ProfileNameField.Text())
		models.DB.Create(currentProfile)
		w.ProfileSelector.AddItem(currentProfile.Name, core.NewQVariant1(currentProfile.ID))
		ix := w.ProfileSelector.FindData(core.NewQVariant1(currentProfile.ID), int(core.Qt__UserRole), core.Qt__MatchExactly)
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
		models.DB.Save(currentProfile)
		ix := w.ProfileSelector.FindData(core.NewQVariant1(currentProfile.ID), int(core.Qt__UserRole), core.Qt__MatchExactly)
		w.ProfileSelector.SetItemText(ix, currentProfile.Name)
	})
	dialog.Show()
}

func (w *MainWindow) removeProfile(_ bool) {
	msgBox := widgets.QMessageBox_Question(nil, "Remove Profile",
		fmt.Sprintf("Are you sure you want to remove the profile %v?",
			currentProfile.Name), widgets.QMessageBox__Yes|widgets.QMessageBox__No, 0)

	if msgBox == widgets.QMessageBox__Yes {
		ix := w.ProfileSelector.FindData(core.NewQVariant1(currentProfile.ID), int(core.Qt__UserRole), core.Qt__MatchExactly)
		w.ProfileSelector.RemoveItem(ix)
		//models.DB.MustExec(models.SqlRemoveProfileById, currentProfile.ID)
		models.DB.Delete(&currentProfile)
		var nProfiles int
		//models.DB.Get(&nProfiles, models.SqlCountProfiles)
		models.DB.Model(&models.Profile{}).Count(&nProfiles)
		if nProfiles == 0 {
			//rows, _ := models.DB.Exec(models.SqlNewProfile, "Default")
			//newID, _ := rows.LastInsertId()
			var currentProfile = models.NewProfile("Default")
			models.DB.Create(&currentProfile)
			w.ProfileSelector.AddItem("Default", core.NewQVariant1(currentProfile.ID))
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
			err := models.DB.Where("id = ?", e.Message).First(&currentRepo)
			utils.Log.Info("currentRepo val:", currentRepo)
			if err == nil {
				models.DB.Model(&currentProfile).Association("Repo").Replace(currentRepo)
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
		case "BorgNotFound":
			w.displayLogMessage("Couldn't find Borg binary.")
		case "CheckVersion":
			if !borg.FeatureIsSupported("JSON_LOG") {
				w.displayLogMessage("Your Borg version is too old.")
			} else {
				w.displayLogMessage("Borg binary was found and is ready for use.")
			}
		case "BorgRunStart":
			w.CreateStartBtn.SetDisabled(true)
			w.CancelButton.SetDisabled(false)
			Tabs.ArchiveTab.ToggleButtons(true)
		case "BorgRunStop":
			w.CreateStartBtn.SetDisabled(false)
			w.CancelButton.SetDisabled(true)
			Tabs.ArchiveTab.ToggleButtons(false)
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
