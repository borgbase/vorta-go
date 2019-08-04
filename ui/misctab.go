package ui

import (
	"vorta/borg"
	"vorta/utils"
)

var version string

func (w *MiscTab) init() {
	w.VersionLabel.SetText(version)

	go func() {
		r, err := borg.NewVersionRun(currentProfile)
		if err != nil {
			MainWindowChan <- utils.VEvent{Topic: "BorgNotFound", Message: ""}
			return
		}
		err = r.Run()
		r.ProcessResult()
		w.BorgVersion.SetText(borg.BorgVersion)
		w.BorgPath.SetText(r.Bin.Path)
		MainWindowChan <- utils.VEvent{Topic: "CheckVersion", Message: ""}
	}()
}
