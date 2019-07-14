package main

import (
    "os"
    "vorta-go/ui"
    //"github.com/therecipe/qt/core"
    //"github.com/therecipe/qt/uitools"
    "github.com/therecipe/qt/widgets"
    "github.com/therecipe/qt/gui"
)

func main() {
    widgets.NewQApplication(len(os.Args), os.Args)
    
    w := ui.NewMainWindow(nil)

    w.TabWidget.AddTab(ui.NewRepoTab(w), "Repository")
    w.TabWidget.AddTab(ui.NewSourceTab(w), "Sources")
    w.TabWidget.AddTab(ui.NewScheduleTab(w), "Schedule")
    w.TabWidget.AddTab(ui.NewArchiveTab(w), "Archives")
    w.TabWidget.AddTab(ui.NewMiscTab(w), "Misc")
    w.Show()

    trayMenu := widgets.NewQMenu(nil)
    trayMenu.AddAction("Vorta for Borg Backup")
    trayMenu.AddSeparator()
    trayMenu.AddAction("Backup Now")
    trayMenu.AddAction("Quit")

    //pixmap := gui.NewQPixmap5(":/balloon.jpg", "", core.Qt__AutoColor)
    trayIcon := gui.NewQIcon5(":qml/icons/hdd-o-dark.png")

    tray := widgets.NewQSystemTrayIcon(nil)
    tray.SetContextMenu(trayMenu)
    tray.SetIcon(trayIcon)
    tray.SetVisible(true)
    tray.Show()

    widgets.QApplication_Exec()
}
