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
	header.SetVisible(true)
	header.SetSectionResizeMode2(0, widgets.QHeaderView__ResizeToContents)
	header.SetSectionResizeMode2(1, widgets.QHeaderView__ResizeToContents)
	header.SetSectionResizeMode2(2, widgets.QHeaderView__ResizeToContents)
	header.SetSectionResizeMode2(3, widgets.QHeaderView__Interactive)
	header.SetSectionResizeMode2(4, widgets.QHeaderView__Stretch)
	header.SetStretchLastSection(true)

	w.ArchiveTable.SetSelectionBehavior(widgets.QAbstractItemView__SelectRows)
	w.ArchiveTable.SetSelectionMode(widgets.QAbstractItemView__SingleSelection)
	w.ArchiveTable.SetEditTriggers(widgets.QAbstractItemView__NoEditTriggers)
	w.ArchiveTable.SetWordWrap(false)
	w.ArchiveTable.SetAlternatingRowColors(true)
	w.ArchiveTable.SetTextElideMode(core.Qt__TextElideMode(core.Qt__ElideLeft))

	//self.archiveTable.cellDoubleClicked.connect(self.cell_double_clicked)
	//self.archiveTable.itemSelectionChanged.connect(self.update_mount_button_text)
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

/*
   def populate_from_profile(self):
       """Populate archive list and prune settings from profile."""
       profile = self.profile()
       if profile.repo is not None:
           self.mount_points = get_mount_points(profile.repo.url)
           archives = [s for s in profile.repo.archives.select().order_by(ArchiveModel.time.desc())]

           for row, archive in enumerate(archives):
               self.archiveTable.insertRow(row)

               formatted_time = archive.time.strftime('%Y-%m-%d %H:%M')
               self.archiveTable.setItem(row, 0, QTableWidgetItem(formatted_time))
               self.archiveTable.setItem(row, 1, QTableWidgetItem(pretty_bytes(archive.size)))
               if archive.duration is not None:
                   formatted_duration = str(timedelta(seconds=round(archive.duration)))
               else:
                   formatted_duration = ''

               self.archiveTable.setItem(row, 2, QTableWidgetItem(formatted_duration))

               mount_point = self.mount_points.get(archive.name)
               if mount_point is not None:
                   item = QTableWidgetItem(mount_point)
                   self.archiveTable.setItem(row, 3, item)

               self.archiveTable.setItem(row, 4, QTableWidgetItem(archive.name))

           self.archiveTable.setRowCount(len(archives))
           item = self.archiveTable.item(0, 0)
           self.archiveTable.scrollToItem(item)
           self._toggle_all_buttons(enabled=True)
       else:
           self.mount_points = {}
           self.archiveTable.setRowCount(0)
           self.toolBox.setItemText(0, self.tr('Archives'))
           self._toggle_all_buttons(enabled=False)

       self.archiveNameTemplate.setText(profile.new_archive_name)
       self.prunePrefixTemplate.setText(profile.prune_prefix)

 */

}
