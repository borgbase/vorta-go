package app

import (
	"vorta-go/borg"
	"vorta-go/utils"
)

func RunAppEventHandler(UIChan chan utils.VEvent) {
	for e := range AppChan {
		switch e.Topic {
		case "StatusUpdate":
			UIChan <- e
		case "OpenMainWindow":
			UIChan <- e
		case "StartBackup":
			go StartBackupEventHandler(e)
		default:
			utils.Log.Info(e)
		}
	}
}

func StartBackupEventHandler(e utils.VEvent) {
	b, err := borg.NewCreateRun(e.Profile)
	if err != nil {
		utils.Log.Error(err)
	}
	AppChan <- utils.VEvent{Topic: "StatusUpdate", Message: "Started Backup"}
	b.Run()
}
