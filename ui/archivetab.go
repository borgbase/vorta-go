package ui

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"time"
	"vorta-go/models"
	"vorta-go/utils"
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

	w.ArchiveNameTemplate.ConnectTextChanged(func(text string) {
		currentProfile.NewArchiveName = text
		currentProfile.SaveField("new_archive_name")
	})
	w.PrunePrefixTemplate.ConnectTextChanged(func(text string) {
		currentProfile.PrunePrefix = text
		currentProfile.SaveField("prune_prefix")
	})
}

func (w *ArchiveTab) Populate() {
	w.ToolBox.SetItemText(0, fmt.Sprintf("Archives for %s", currentRepo.Url))
	archives := []models.Archive{}
	err := models.DB.Select(&archives, models.SqlAllArchivesByRepoId, currentRepo.Id)
	if err != nil {
		utils.Log.Error(err)
	}
	w.ArchiveTable.SetRowCount(len(archives))

	for row, archive := range archives {
		w.ArchiveTable.SetItem(row, 0, widgets.NewQTableWidgetItem2(archive.CreatedAt.Format(_timeFormat), 0))
		w.ArchiveTable.SetItem(row, 1, widgets.NewQTableWidgetItem2(humanize.Bytes(uint64(archive.Size.Int64)), 0))
		formattedDuration := time.Duration(archive.Duration.Float64)*time.Second
		w.ArchiveTable.SetItem(row, 2, widgets.NewQTableWidgetItem2(formattedDuration.String(), 0))
		w.ArchiveTable.SetItem(row, 3, widgets.NewQTableWidgetItem2("", 0))
		w.ArchiveTable.SetItem(row, 4, widgets.NewQTableWidgetItem2(archive.Name, 0))
	}
	w.ArchiveTable.ResizeColumnsToContents()
	w.ArchiveNameTemplate.SetText(currentProfile.NewArchiveName)
	w.PrunePrefixTemplate.SetText(currentProfile.PrunePrefix)
}
