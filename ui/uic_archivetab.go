package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
)

type __archivetab struct{}

func (*__archivetab) init() {}

type ArchiveTab struct {
	*__archivetab
	*widgets.QWidget
	VerticalLayout      *widgets.QVBoxLayout
	ToolBox             *widgets.QToolBox
	ArchivesToolBox     *widgets.QWidget
	VerticalLayout_2    *widgets.QVBoxLayout
	ArchiveTable        *widgets.QTableWidget
	HorizontalLayout_3  *widgets.QHBoxLayout
	ExtractButton       *widgets.QPushButton
	MountButton         *widgets.QPushButton
	CheckButton         *widgets.QPushButton
	DeleteButton        *widgets.QPushButton
	HorizontalSpacer    *widgets.QSpacerItem
	PruneButton         *widgets.QPushButton
	ListButton          *widgets.QPushButton
	MountErrors         *widgets.QLabel
	PruningToolBox      *widgets.QWidget
	VerticalLayout_3    *widgets.QVBoxLayout
	Label_5             *widgets.QLabel
	HorizontalLayout_7  *widgets.QHBoxLayout
	Label_12            *widgets.QLabel
	Prune_hour          *widgets.QSpinBox
	Label_6             *widgets.QLabel
	Prune_day           *widgets.QSpinBox
	Label_7             *widgets.QLabel
	Prune_week          *widgets.QSpinBox
	Label_8             *widgets.QLabel
	Prune_month         *widgets.QSpinBox
	Label_9             *widgets.QLabel
	Prune_year          *widgets.QSpinBox
	Label_10            *widgets.QLabel
	HorizontalSpacer_5  *widgets.QSpacerItem
	HorizontalLayout_2  *widgets.QHBoxLayout
	Label_2             *widgets.QLabel
	Prune_keep_within   *widgets.QLineEdit
	HorizontalSpacer_3  *widgets.QSpacerItem
	Line                *widgets.QFrame
	GridLayout          *widgets.QGridLayout
	Label_3             *widgets.QLabel
	PrunePrefixTemplate *widgets.QLineEdit
	ArchiveNamePreview  *widgets.QLabel
	ArchiveNameTemplate *widgets.QLineEdit
	Label_4             *widgets.QLabel
	PrunePrefixPreview  *widgets.QLabel
	VerticalSpacer      *widgets.QSpacerItem
}

func NewArchiveTab(p widgets.QWidget_ITF) *ArchiveTab {
	var par *widgets.QWidget
	if p != nil {
		par = p.QWidget_PTR()
	}
	file := core.NewQFile2(":/ui/archivetab.ui")
	file.Open(core.QIODevice__ReadOnly)
	w := &ArchiveTab{QWidget: widgets.NewQWidgetFromPointer(uitools.NewQUiLoader(nil).Load(file, par).Pointer())}
	file.Close()
	w.setupUI()
	w.init()
	return w
}
func (w *ArchiveTab) setupUI() {
	w.Label_4 = widgets.NewQLabelFromPointer(w.FindChild("label_4", core.Qt__FindChildrenRecursively).Pointer())
	w.ArchivesToolBox = widgets.NewQWidgetFromPointer(w.FindChild("archivesToolBox", core.Qt__FindChildrenRecursively).Pointer())
	w.DeleteButton = widgets.NewQPushButtonFromPointer(w.FindChild("deleteButton", core.Qt__FindChildrenRecursively).Pointer())
	w.PruneButton = widgets.NewQPushButtonFromPointer(w.FindChild("pruneButton", core.Qt__FindChildrenRecursively).Pointer())
	w.Prune_hour = widgets.NewQSpinBoxFromPointer(w.FindChild("prune_hour", core.Qt__FindChildrenRecursively).Pointer())
	w.Label_10 = widgets.NewQLabelFromPointer(w.FindChild("label_10", core.Qt__FindChildrenRecursively).Pointer())
	w.ArchiveNameTemplate = widgets.NewQLineEditFromPointer(w.FindChild("archiveNameTemplate", core.Qt__FindChildrenRecursively).Pointer())
	w.PruningToolBox = widgets.NewQWidgetFromPointer(w.FindChild("pruningToolBox", core.Qt__FindChildrenRecursively).Pointer())
	w.Prune_year = widgets.NewQSpinBoxFromPointer(w.FindChild("prune_year", core.Qt__FindChildrenRecursively).Pointer())
	w.Line = widgets.NewQFrameFromPointer(w.FindChild("line", core.Qt__FindChildrenRecursively).Pointer())
	w.HorizontalSpacer = widgets.NewQSpacerItemFromPointer(w.FindChild("horizontalSpacer", core.Qt__FindChildrenRecursively).Pointer())
	w.Label_12 = widgets.NewQLabelFromPointer(w.FindChild("label_12", core.Qt__FindChildrenRecursively).Pointer())
	w.Label_6 = widgets.NewQLabelFromPointer(w.FindChild("label_6", core.Qt__FindChildrenRecursively).Pointer())
	w.HorizontalSpacer_3 = widgets.NewQSpacerItemFromPointer(w.FindChild("horizontalSpacer_3", core.Qt__FindChildrenRecursively).Pointer())
	w.PrunePrefixPreview = widgets.NewQLabelFromPointer(w.FindChild("prunePrefixPreview", core.Qt__FindChildrenRecursively).Pointer())
	w.CheckButton = widgets.NewQPushButtonFromPointer(w.FindChild("checkButton", core.Qt__FindChildrenRecursively).Pointer())
	w.HorizontalLayout_7 = widgets.NewQHBoxLayoutFromPointer(w.FindChild("horizontalLayout_7", core.Qt__FindChildrenRecursively).Pointer())
	w.HorizontalLayout_2 = widgets.NewQHBoxLayoutFromPointer(w.FindChild("horizontalLayout_2", core.Qt__FindChildrenRecursively).Pointer())
	w.Label_7 = widgets.NewQLabelFromPointer(w.FindChild("label_7", core.Qt__FindChildrenRecursively).Pointer())
	w.Label_2 = widgets.NewQLabelFromPointer(w.FindChild("label_2", core.Qt__FindChildrenRecursively).Pointer())
	w.Label_3 = widgets.NewQLabelFromPointer(w.FindChild("label_3", core.Qt__FindChildrenRecursively).Pointer())
	w.PrunePrefixTemplate = widgets.NewQLineEditFromPointer(w.FindChild("prunePrefixTemplate", core.Qt__FindChildrenRecursively).Pointer())
	w.ArchiveNamePreview = widgets.NewQLabelFromPointer(w.FindChild("archiveNamePreview", core.Qt__FindChildrenRecursively).Pointer())
	w.ToolBox = widgets.NewQToolBoxFromPointer(w.FindChild("toolBox", core.Qt__FindChildrenRecursively).Pointer())
	w.MountButton = widgets.NewQPushButtonFromPointer(w.FindChild("mountButton", core.Qt__FindChildrenRecursively).Pointer())
	w.ListButton = widgets.NewQPushButtonFromPointer(w.FindChild("listButton", core.Qt__FindChildrenRecursively).Pointer())
	w.MountErrors = widgets.NewQLabelFromPointer(w.FindChild("mountErrors", core.Qt__FindChildrenRecursively).Pointer())
	w.Prune_month = widgets.NewQSpinBoxFromPointer(w.FindChild("prune_month", core.Qt__FindChildrenRecursively).Pointer())
	w.ExtractButton = widgets.NewQPushButtonFromPointer(w.FindChild("extractButton", core.Qt__FindChildrenRecursively).Pointer())
	w.Label_9 = widgets.NewQLabelFromPointer(w.FindChild("label_9", core.Qt__FindChildrenRecursively).Pointer())
	w.HorizontalSpacer_5 = widgets.NewQSpacerItemFromPointer(w.FindChild("horizontalSpacer_5", core.Qt__FindChildrenRecursively).Pointer())
	w.Prune_keep_within = widgets.NewQLineEditFromPointer(w.FindChild("prune_keep_within", core.Qt__FindChildrenRecursively).Pointer())
	w.Prune_day = widgets.NewQSpinBoxFromPointer(w.FindChild("prune_day", core.Qt__FindChildrenRecursively).Pointer())
	w.Prune_week = widgets.NewQSpinBoxFromPointer(w.FindChild("prune_week", core.Qt__FindChildrenRecursively).Pointer())
	w.VerticalLayout = widgets.NewQVBoxLayoutFromPointer(w.FindChild("verticalLayout", core.Qt__FindChildrenRecursively).Pointer())
	w.VerticalLayout_2 = widgets.NewQVBoxLayoutFromPointer(w.FindChild("verticalLayout_2", core.Qt__FindChildrenRecursively).Pointer())
	w.ArchiveTable = widgets.NewQTableWidgetFromPointer(w.FindChild("archiveTable", core.Qt__FindChildrenRecursively).Pointer())
	w.HorizontalLayout_3 = widgets.NewQHBoxLayoutFromPointer(w.FindChild("horizontalLayout_3", core.Qt__FindChildrenRecursively).Pointer())
	w.VerticalLayout_3 = widgets.NewQVBoxLayoutFromPointer(w.FindChild("verticalLayout_3", core.Qt__FindChildrenRecursively).Pointer())
	w.Label_5 = widgets.NewQLabelFromPointer(w.FindChild("label_5", core.Qt__FindChildrenRecursively).Pointer())
	w.Label_8 = widgets.NewQLabelFromPointer(w.FindChild("label_8", core.Qt__FindChildrenRecursively).Pointer())
	w.GridLayout = widgets.NewQGridLayoutFromPointer(w.FindChild("gridLayout", core.Qt__FindChildrenRecursively).Pointer())
	w.VerticalSpacer = widgets.NewQSpacerItemFromPointer(w.FindChild("verticalSpacer", core.Qt__FindChildrenRecursively).Pointer())
}
