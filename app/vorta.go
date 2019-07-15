package app

import (
	"github.com/ProtonMail/go-appdir"
	"github.com/sirupsen/logrus"
	"github.com/therecipe/qt/widgets"
	"io"
	"os"
	"os/exec"
	"path"
	"vorta-go/models"
)

var (
	QtApp *widgets.QApplication
	StatusUpdateChannel chan string
	CurrentCommand *exec.Cmd
	ConfigDir appdir.Dirs
	Log *logrus.Logger
	CurrentProfile *models.Profile
)

func InitApp() {
	ConfigDir = appdir.New("Vorta")

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

	QtApp = widgets.NewQApplication(len(os.Args), os.Args)
	StatusUpdateChannel = make(chan string)
	InitTray()
}
