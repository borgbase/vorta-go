package ui

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"time"
	"vorta/borg"
	"vorta/models"
)

var (
	_timeFormat = "2006-01-02 15:04"
)

func (w *ArchiveTab) init() {
	w.ToolBox.SetCurrentIndex(0)
	header := w.ArchiveTable.HorizontalHeader()
	header.SetSectionResizeMode2(0, widgets.QHeaderView__ResizeToContents)
	header.SetSectionResizeMode2(1, widgets.QHeaderView__ResizeToContents)
	header.SetSectionResizeMode2(2, widgets.QHeaderView__ResizeToContents)
	header.SetSectionResizeMode2(3, widgets.QHeaderView__Interactive)
	header.SetSectionResizeMode2(4, widgets.QHeaderView__Stretch)
	header.SetStretchLastSection(true)
	header.SetVisible(true)

	w.ArchiveTable.SetSelectionBehavior(widgets.QAbstractItemView__SelectRows)
	w.ArchiveTable.SetSelectionMode(widgets.QAbstractItemView__SingleSelection)
	w.ArchiveTable.SetEditTriggers(widgets.QAbstractItemView__NoEditTriggers)
	w.ArchiveTable.SetWordWrap(false)
	w.ArchiveTable.SetAlternatingRowColors(true)
	w.ArchiveTable.SetTextElideMode(core.Qt__TextElideMode(core.Qt__ElideLeft))

	w.ArchiveNameTemplate.ConnectTextChanged(w.archiveNameTemplateChanged)
	w.PrunePrefixTemplate.ConnectTextChanged(w.prunePrefixTemplateChanged)

	w.CheckButton.ConnectClicked(func(_ bool) {
		r, err := borg.NewCheckRun(currentProfile)
		if err != nil {
			w.MountErrors.SetText(err.Error())
			return
		}
		w.MountErrors.SetText("Starting repo check...")
		go func() {
			err := r.Run()
			if err != nil {
				w.MountErrors.SetText(err.Error())
			}
		}()
	})
}

func (w *ArchiveTab) archiveNameTemplateChanged(text string) {
	currentProfile.NewArchiveName = text
	models.DB.Save(currentProfile)
}

func (w *ArchiveTab) prunePrefixTemplateChanged(text string) {
	currentProfile.PrunePrefix = text
	models.DB.Save(currentProfile)
}

func (w *ArchiveTab) Populate() {
	// Deal with archive name options
	w.ArchiveNameTemplate.DisconnectTextChanged()
	w.PrunePrefixTemplate.DisconnectTextChanged()
	w.ArchiveNameTemplate.SetText(currentProfile.NewArchiveName)
	w.PrunePrefixTemplate.SetText(currentProfile.PrunePrefix)
	w.ArchiveNameTemplate.ConnectTextChanged(w.archiveNameTemplateChanged)
	w.PrunePrefixTemplate.ConnectTextChanged(w.prunePrefixTemplateChanged)

	// Populate archive table
	w.ToolBox.SetItemText(0, fmt.Sprintf("Archives for %s", currentRepo.Url))
	archives := []models.Archive{}
	models.DB.Model(&currentRepo).Related(&archives)

	if len(archives) == 0 {
		return
	}
	w.ArchiveTable.SetRowCount(len(archives))

	for row, archive := range archives {
		w.ArchiveTable.SetItem(row, 0, widgets.NewQTableWidgetItem2(archive.CreatedAt.Format(_timeFormat), 0))
		w.ArchiveTable.SetItem(row, 1, widgets.NewQTableWidgetItem2(humanize.Bytes(uint64(archive.Size.Int64)), 0))
		formattedDuration := time.Duration(archive.Duration.Float64) * time.Second
		w.ArchiveTable.SetItem(row, 2, widgets.NewQTableWidgetItem2(formattedDuration.String(), 0))
		w.ArchiveTable.SetItem(row, 3, widgets.NewQTableWidgetItem2("", 0))
		w.ArchiveTable.SetItem(row, 4, widgets.NewQTableWidgetItem2(archive.Name, 0))
	}
	w.ArchiveTable.ResizeColumnsToContents()


}
