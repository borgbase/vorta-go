package main

import (
    "github.com/therecipe/qt/widgets"
    "vorta-go/app"
    "vorta-go/borg"
    "vorta-go/models"
    "vorta-go/ui"
    "vorta-go/utils"
)

func main() {
    utils.InitLog()
    app.InitApp()
    app.InitScheduler()
    app.AppChan = make(chan utils.VEvent)
    borg.AppEventChan = app.AppChan

    models.InitDb(utils.ConfigDir.UserData())
    defer models.DB.Close()

    w := ui.NewMainWindow(nil)
    go w.RunUIEventHandler(app.AppChan)
    go app.RunAppEventHandler(ui.MainWindowChan)
    w.AddTabs()

    widgets.QApplication_Exec()
}
