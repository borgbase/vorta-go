package main

import (
    "github.com/therecipe/qt/widgets"
    "vorta-go/app"
    "vorta-go/models"
    "vorta-go/ui"
    "vorta-go/utils"
)

func main() {
    app.InitApp()
    app.InitScheduler()
    app.AppChan = make(chan utils.VEvent)

    models.InitDb(app.ConfigDir.UserData())
    defer models.DB.Close()

    w := ui.NewMainWindow(nil)
    w.AddTabs()

    go w.RunUIEventHandler(app.AppChan)
    go app.RunAppEventHandler(ui.MainWindowChan)

    widgets.QApplication_Exec()
}
