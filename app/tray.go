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
	profileMenu := t.menu.AddMenu2("Backup Now")
	pp := []models.Profile{}
	models.DB.Find(&pp)
	for _, p := range pp {
		profile := p
		profileMenu.AddAction(p.Name).ConnectTriggered(func(_ bool) {
			utils.Log.Infof("Backup triggered for profile: %v", profile.Name)
			AppChan <- utils.VEvent{Topic: "StartBackup", Profile: &profile}
		})
	}

	t.menu.AddAction("Quit").ConnectTriggered(func(checked bool) { QtApp.Quit() })
}

func (t *systemTray) SetIcon(active bool) {
	var fileName string
	if active {
		fileName = ":qml/icons/hdd-o-active-dark.png"
	} else {
		fileName = ":qml/icons/hdd-o-dark.png"
	}
	trayIcon := gui.NewQIcon5(fileName)
	vortaTray.icon.Hide()
	vortaTray.icon.SetIcon(trayIcon)
	vortaTray.icon.Show()
}
