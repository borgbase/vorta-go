package app

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"os"
	"runtime"
	"strings"
	"vorta-go/utils"
)

var (
	AppChan chan utils.VEvent
	QtApp   *widgets.QApplication
)

func InitApp() {
	// Set up Qt App
	QtApp = widgets.NewQApplication(len(os.Args), os.Args)
	QtApp.SetQuitOnLastWindowClosed(false)

	// TODO: will this prevent proper dark theme? https://www.linuxuprising.com/2018/05/get-qt5-apps-to-use-native-gtk-style-in.html
	if strings.HasPrefix(runtime.GOOS, "linux") {
		QtApp.SetStyle2("fusion")
	}

	// Load translations
	qtTranslator := core.NewQTranslator(nil)
	success := qtTranslator.Load("ui_de", ":/qml/i18n/", "", "") //+core.QLocale_System().Name()
	//success := qtTranslator.Load(":/qml/i18n/ui_de.qm", "", "", "") //+core.QLocale_System().Name()
	//success := qtTranslator.Load2(core.NewQLocale2("de"), "ui_de", "", ":/qml/i18n/", "")
	utils.Log.Info("loaded translations:", success)
	QtApp.InstallTranslator(qtTranslator)

	InitTray()
}
