package utils

import (
	"github.com/ProtonMail/go-appdir"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
)

var (
	Log       *logrus.Logger
	ConfigDir appdir.Dirs
)

func InitLog() {
	// Find and create required folders
	ConfigDir = appdir.New("Vorta")
	requiredFolders := []string{ConfigDir.UserLogs(), ConfigDir.UserData()}
	for _, p := range requiredFolders {
		if _, err := os.Stat(p); os.IsNotExist(err) {
			err := os.MkdirAll(p, os.ModePerm)
			if err != nil {
				Log.Panicf("error while creating required folder: %v", err)
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

	logFile, err := os.OpenFile(path.Join(ConfigDir.UserLogs(), "vorta-go.log"), os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		Log.Panic("Can't open log file.")
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	Log.SetOutput(mw)
}
