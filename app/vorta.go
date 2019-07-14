package app

import (
	"github.com/therecipe/qt/widgets"
	"os"
	"os/exec"
)

type VortaApp struct {
	QtApp *widgets.QApplication
	StatusUpdateChannel chan string
	CurrentCommand *exec.Cmd
}

var App *VortaApp

func InitApp() {
	app := VortaApp{
		QtApp: widgets.NewQApplication(len(os.Args), os.Args),
		StatusUpdateChannel: make(chan string),
	}

	App = &app
	InitTray()
}
