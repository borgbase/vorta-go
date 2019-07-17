package app

import (
	"github.com/ProtonMail/go-appdir"
	"github.com/sirupsen/logrus"
	"github.com/therecipe/qt/widgets"
	"io"
	"os"
	"path"
	"vorta-go/utils"
)

var (
	AppChan   chan utils.VEvent
	ConfigDir appdir.Dirs
	Log       *logrus.Logger
	QtApp	  *widgets.QApplication
)

func InitApp() {
	// Find and create required folders (settings and logs)
	ConfigDir = appdir.New("Vorta")

	requiredFolders := []string{ConfigDir.UserLogs(), ConfigDir.UserData()}
	for _, p := range requiredFolders {
		if _, err := os.Stat(p); os.IsNotExist(err) {
			err := os.Mkdir(p, os.ModePerm)
			if err != nil {
				panic("Unable to create required folder.")
			}
		}
	}

	// Set up logging
	Log = logrus.New()
	Formatter := new(logrus.TextFormatter)
	Formatter.TimestampFormat = "2006-01-02 15:04:05"
	Formatter.FullTimestamp = true
	Log.SetFormatter(Formatter)
	// TODO: make cli argument
	Log.SetLevel(logrus.DebugLevel)

	logFile, err := os.OpenFile(path.Join(ConfigDir.UserLogs(), "vorta-go.log"), os.O_WRONLY | os.O_CREATE, 0755)
	if err != nil {
		Log.Panic("Can't open log file.")
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	Log.SetOutput(mw)

	// Set up Qt App
	QtApp = widgets.NewQApplication(len(os.Args), os.Args)
	QtApp.SetQuitOnLastWindowClosed(false)

	InitTray()
}

func RunAppEventHandler(UIChan chan utils.VEvent) {
	for e := range AppChan {
		switch e.Topic {
		case "StatusUpdate":
			UIChan <- e
		case "OpenMainWindow":
			UIChan <- e
		default:
			Log.Info(e)
		}
	}
}
