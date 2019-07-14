package main

import (
    "github.com/therecipe/qt/widgets"
    "os"
    "vorta-go/models"
    "vorta-go/ui"
)

func main() {
    widgets.NewQApplication(len(os.Args), os.Args)
    models.InitDb()
    defer models.DB.Close()

    ui.NewMainWindow(nil)

    widgets.QApplication_Exec()
}
