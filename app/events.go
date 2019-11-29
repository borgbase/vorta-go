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
				vortaTray.SetIcon(true)
				err = b.Run()
				if err != nil {
					utils.Log.Errorf("Error during backup run: %v", err)
				} else {
					b.ProcessResult()
					UIChan <- utils.VEvent{Topic: "UpdateArchiveTab"}

				}
				vortaTray.SetIcon(false)
			}()
		case "CancelBorgRun":
			borg.CancelBorgRun()
		default:
			utils.Log.Info(e)
		}
	}
}
