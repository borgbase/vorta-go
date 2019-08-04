package app

import (
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"vorta/models"
	"vorta/utils"
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
	t.menu.AddAction("Vorta for Borg Backup").ConnectTriggered(
		func(checked bool) {
			AppChan <- utils.VEvent{Topic: "OpenMainWindow", Message: "From Tray"}
		})
	t.menu.AddSeparator()
	//t.menu.AddAction("Backup Now")
	profileMenu := t.menu.AddMenu2("Backup Now")
	pp := []models.Profile{}
	models.DB.Select(&pp, models.SqlAllProfiles)
	for _, p := range pp {
		profileName := p.Name
		profileMenu.AddAction(p.Name).ConnectTriggered(func(checked bool) { utils.Log.Info("Would backup profile", profileName) })
	}

	t.menu.AddAction("Quit").ConnectTriggered(func(checked bool) { QtApp.Quit() })
}
