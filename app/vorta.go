package app

import (
	"github.com/therecipe/qt/widgets"
	"os"
	"runtime"
	"strings"
	"vorta-go/utils"
)

var (
	AppChan   chan utils.VEvent
	QtApp	  *widgets.QApplication
)

func InitApp() {
	// Set up Qt App
	QtApp = widgets.NewQApplication(len(os.Args), os.Args)
	QtApp.SetQuitOnLastWindowClosed(false)

	if strings.HasPrefix(runtime.GOOS, "linux") {
		QtApp.SetStyle2("fusion")
	}

	InitTray()
}
