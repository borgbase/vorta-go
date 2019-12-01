package ui

import (
	"database/sql"
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

	w.MountButton.ConnectClicked(func(_ bool) {
		row := w.ArchiveTable.CurrentRow()
		archiveName := w.ArchiveTable.Item(row, 4).Text()
		ChooseFileDialog(func(files []string) {
			r, err := borg.NewMountRun(currentProfile)
			r.SubCommandArgs[len(r.SubCommandArgs)-1] += "::" + archiveName
			r.SubCommandArgs = append(r.SubCommandArgs, files[0])
			if err != nil {
				w.MountErrors.SetText(err.Error())
				return
			}
			w.MountErrors.SetText("Mounting archive...")
			go func() {
				err := r.Run()
				if err != nil {
					w.MountErrors.SetText(err.Error())
				}
				w.MountErrors.SetText("Archive mounted.")
			}()
		})
	})

	w.PruneButton.ConnectClicked(func(_ bool) {
		r, err :=  borg.NewPruneRun(currentProfile)
		if err != nil {
			w.MountErrors.SetText(err.Error())
			return
		}
		go func() {
			err := r.Run()
			if err != nil {
				w.MountErrors.SetText(err.Error())
			}
			r.ProcessResult()
			w.Populate()
		}()
	})

	w.ListButton.ConnectClicked(func(_ bool) {
		r, err := borg.NewListRepoRun(currentProfile)
		if err != nil {
			w.MountErrors.SetText(err.Error())
			return
		}
		go func() {
			err := r.Run()
			if err != nil {
				w.MountErrors.SetText(err.Error())
			}
			r.ProcessResult()
			w.Populate()
		}()
	})
}

func (w *ArchiveTab) archiveNameTemplateChanged(text string) {
	currentProfile.NewArchiveName = text
	models.DB.Save(currentProfile)

	preview := currentProfile.FormatArchiveName(currentProfile.NewArchiveName)
	w.ArchiveNamePreview.SetText(preview)
}

func (w *ArchiveTab) prunePrefixTemplateChanged(text string) {
	currentProfile.PrunePrefix = text
	models.DB.Save(currentProfile)

	preview := currentProfile.FormatArchiveName(currentProfile.PrunePrefix)
	w.PrunePrefixPreview.SetText(preview)
}

func (w *ArchiveTab) pruneSettingsChanged(_ int) {
	currentProfile.PruneHour = w.Prune_hour.Value()
	currentProfile.PruneDay = w.Prune_day.Value()
	currentProfile.PruneWeek = w.Prune_week.Value()
	currentProfile.PruneMonth = w.Prune_month.Value()
	currentProfile.PruneYear = w.Prune_year.Value()
	models.DB.Save(currentProfile)
}

func (w *ArchiveTab) Populate() {
	// Deal with archive name options
	w.ArchiveNameTemplate.DisconnectTextChanged()  // Disconnect first to avoid accidential changes
	w.PrunePrefixTemplate.DisconnectTextChanged()
	w.ArchiveNameTemplate.SetText(currentProfile.NewArchiveName)
	w.PrunePrefixTemplate.SetText(currentProfile.PrunePrefix)
	w.ArchiveNameTemplate.ConnectTextChanged(w.archiveNameTemplateChanged)
	w.PrunePrefixTemplate.ConnectTextChanged(w.prunePrefixTemplateChanged)

	// Populate prune options
	pruneFields := []*widgets.QSpinBox{w.Prune_hour, w.Prune_day, w.Prune_week, w.Prune_month, w.Prune_year}
	for _, field := range pruneFields {
		field.DisconnectValueChanged()
	}
	w.Prune_hour.SetValue(currentProfile.PruneHour)
	w.Prune_day.SetValue(currentProfile.PruneDay)
	w.Prune_week.SetValue(currentProfile.PruneWeek)
	w.Prune_month.SetValue(currentProfile.PruneMonth)
	w.Prune_year.SetValue(currentProfile.PruneYear)
	w.Prune_keep_within.SetText(currentProfile.PruneKeepWithin.String)

	for _, field := range pruneFields {
		field.ConnectValueChanged(w.pruneSettingsChanged)
	}
	w.Prune_keep_within.ConnectTextChanged(func(text string) {
		currentProfile.PruneKeepWithin = sql.NullString{w.Prune_keep_within.Text(), true}
		models.DB.Save(currentProfile)
	})

	// Populate archive table
	w.ToolBox.SetItemText(0, fmt.Sprintf("Archives for %s", currentRepo.Url))
	archives := []models.Archive{}
	models.DB.Order("time desc").Model(&currentRepo).Related(&archives)

	w.ArchiveTable.SetRowCount(len(archives))

	if len(archives) == 0 {
		return
	}

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

func (w *ArchiveTab) ToggleButtons(SetDisabled bool) {
	Tabs.ArchiveTab.ListButton.SetDisabled(SetDisabled)
	Tabs.ArchiveTab.MountButton.SetDisabled(SetDisabled)
	Tabs.ArchiveTab.DeleteButton.SetDisabled(SetDisabled)
	Tabs.ArchiveTab.ExtractButton.SetDisabled(SetDisabled)
	Tabs.ArchiveTab.CheckButton.SetDisabled(SetDisabled)
	Tabs.ArchiveTab.PruneButton.SetDisabled(SetDisabled)
}
