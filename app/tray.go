package app

import (
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type systemTray struct {
	icon *widgets.QSystemTrayIcon
	menu *widgets.QMenu
}

var vortaTray systemTray

func InitTray() {
	trayMenu := widgets.NewQMenu(nil)
	trayIcon := gui.NewQIcon5(":qml/icons/hdd-o-dark.png")

	tray := widgets.NewQSystemTrayIcon(nil)
	tray.SetContextMenu(trayMenu)
	tray.SetIcon(trayIcon)
	tray.SetVisible(true)
	tray.Show()

	vortaTray = systemTray{
		icon: tray,
		menu: trayMenu,
	}
	trayMenu.ConnectAboutToShow(vortaTray.drawMenu)
}

func (t *systemTray) drawMenu() {
	t.menu.Clear()
	t.menu.AddAction("Vorta for Borg Backup")
	t.menu.AddSeparator()
	t.menu.AddAction("Backup Now")
	t.menu.AddAction("Quit")
}
