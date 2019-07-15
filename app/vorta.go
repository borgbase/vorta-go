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

type VortaApp struct {
	QtApp *widgets.QApplication
	StatusUpdateChannel chan string
	CurrentCommand *exec.Cmd
	Appdir appdir.Dirs
	Log *logrus.Logger
	CurrentProfile *models.Profile
}

var App *VortaApp

func InitApp() {
	vortaAppdir := appdir.New("Vorta")

	log := logrus.New()
	Formatter := new(logrus.TextFormatter)
	Formatter.TimestampFormat = "2006-01-02 15:04:05"
	Formatter.FullTimestamp = true
	log.SetFormatter(Formatter)


	logFile, err := os.OpenFile(path.Join(vortaAppdir.UserLogs(), "vorta-go.log"), os.O_WRONLY | os.O_CREATE, 0755)
	if err != nil {
		log.Panic("Can't open log file.")
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
	log.Info("Logging Ready.")

	p := models.Profile{}
	models.DB.Get(&p, models.SqlOneProfile)

	app := VortaApp{
		QtApp: widgets.NewQApplication(len(os.Args), os.Args),
		StatusUpdateChannel: make(chan string),
		Appdir: vortaAppdir,
		Log: log,
		CurrentProfile: &p,
	}
	App = &app
	InitTray()
}
