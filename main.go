package main

import (
	"github.com/therecipe/qt/widgets"
	"os"
	"path"
	"vorta/app"
	"vorta/borg"
	"vorta/models"
	"vorta/ui"
	"vorta/utils"
)

func main() {
	utils.InitLog()
	pidFile := path.Join(utils.ConfigDir.UserData(), "vorta-go.pid")
	err := utils.WritePidFile(pidFile)
	if err != nil {
		utils.Log.Error("Another instance of Vorta is already running.")
		os.Exit(1)
	}

	models.InitDb(utils.ConfigDir.UserData())
	app.InitApp()
	app.AppChan = make(chan utils.VEvent)
	borg.AppEventChan = app.AppChan //TODO: InitBorg and check version.
	utils.InitScheduler(app.AppChan)

	defer models.DB.Close()

	w := ui.NewMainWindow(nil)
	go w.RunUIEventHandler(app.AppChan)
	go app.RunAppEventHandler(ui.MainWindowChan)
	w.AddTabs()

	widgets.QApplication_Exec()
}
