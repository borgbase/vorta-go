package ui

import (
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

func (w *MainWindow) init() {
	w.TabWidget.AddTab(NewRepoTab(w), "Repository")
	w.TabWidget.AddTab(NewSourceTab(w), "Sources")
	w.TabWidget.AddTab(NewScheduleTab(w), "Schedule")
	w.TabWidget.AddTab(NewArchiveTab(w), "Archives")
	w.TabWidget.AddTab(NewMiscTab(w), "Misc")
	w.Show()

	trayMenu := widgets.NewQMenu(nil)
	trayMenu.AddAction("Vorta for Borg Backup")
	trayMenu.AddSeparator()
	trayMenu.AddAction("Backup Now")
	trayMenu.AddAction("Quit")

	trayIcon := gui.NewQIcon5(":qml/icons/hdd-o-dark.png")

	tray := widgets.NewQSystemTrayIcon(nil)
	tray.SetContextMenu(trayMenu)
	tray.SetIcon(trayIcon)
	tray.SetVisible(true)
	tray.Show()
}
