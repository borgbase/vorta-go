package ui

import (
	"vorta-go/app"
	"vorta-go/borg"
)

func (w *MainWindow) init() {
	w.TabWidget.AddTab(NewRepoTab(w), "Repository")
	w.TabWidget.AddTab(NewSourceTab(w), "Sources")
	w.TabWidget.AddTab(NewScheduleTab(w), "Schedule")
	w.TabWidget.AddTab(NewArchiveTab(w), "Archives")
	w.TabWidget.AddTab(NewMiscTab(w), "Misc")
	w.SetWindowTitle("Vorta for Borg Backup")
	w.Show()

	w.CreateStartBtn.ConnectClicked(w.StartBackup)

	go w.displayLogMessages(app.App.StatusUpdateChannel)
}

func (w *MainWindow) StartBackup(checked bool) {
	app.App.StatusUpdateChannel <- "Running Borg"
	b := borg.BorgCommand{SubCommand: "info"}
	go b.Run()
}

func (w *MainWindow) displayLogMessages(c chan string) {
	for updateStr := range c {
		w.CreateProgressText.SetText(updateStr)
		w.CreateProgressText.Repaint()
	}

}
