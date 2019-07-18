package app

import (
	"github.com/therecipe/qt/widgets"
	"os"
	"vorta-go/utils"
)

var (
	AppChan   chan utils.VEvent
	QtApp	  *widgets.QApplication
)

func InitApp() {
	requiredFolders := []string{utils.ConfigDir.UserLogs(), utils.ConfigDir.UserData()}
	for _, p := range requiredFolders {
		if _, err := os.Stat(p); os.IsNotExist(err) {
			err := os.Mkdir(p, os.ModePerm)
			if err != nil {
				panic("Unable to create required folder.")
			}
		}
	}

	// Set up Qt App
	QtApp = widgets.NewQApplication(len(os.Args), os.Args)
	QtApp.SetQuitOnLastWindowClosed(false)

	InitTray()
}
