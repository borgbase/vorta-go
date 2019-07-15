package main

import (
    "github.com/therecipe/qt/widgets"
    "vorta-go/app"
    "vorta-go/models"
    "vorta-go/ui"
)

func main() {
    app.InitApp()
    app.InitScheduler()
    models.InitDb(app.ConfigDir.UserData())
    defer models.DB.Close()

    ui.NewMainWindow(nil)

    widgets.QApplication_Exec()
}
