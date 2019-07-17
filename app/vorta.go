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
)

func InitApp() {
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

	Log = logrus.New()
	Formatter := new(logrus.TextFormatter)
	Formatter.TimestampFormat = "2006-01-02 15:04:05"
	Formatter.FullTimestamp = true
	Log.SetFormatter(Formatter)

	logFile, err := os.OpenFile(path.Join(ConfigDir.UserLogs(), "vorta-go.log"), os.O_WRONLY | os.O_CREATE, 0755)
	if err != nil {
		Log.Panic("Can't open log file.")
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	Log.SetOutput(mw)
	Log.Info("Logging Ready.")

	widgets.NewQApplication(len(os.Args), os.Args)

	InitTray()
}

func RunAppEventHandler(UIChan chan utils.VEvent) {
	for e := range AppChan {
		Log.Info(e)
	}
}
