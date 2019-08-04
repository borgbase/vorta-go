package app

import (
	"vorta/borg"
	"vorta/utils"
)

func RunAppEventHandler(UIChan chan utils.VEvent) {
	for e := range AppChan {
		switch e.Topic {
		case "StatusUpdate":
			UIChan <- e
		case "OpenMainWindow":
			UIChan <- e
		case "StartBackup":
			go func() {
				b, err := borg.NewCreateRun(e.Profile)
				if err != nil {
					utils.Log.Error(err)
				}
				AppChan <- utils.VEvent{Topic: "StatusUpdate", Message: "Started Backup"}
				b.Run()
				b.ProcessResult()
				UIChan <- utils.VEvent{Topic: "UpdateArchiveTab"}
			}()
		default:
			utils.Log.Info(e)
		}
	}
}
