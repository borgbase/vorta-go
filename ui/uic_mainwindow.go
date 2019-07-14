package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
)

type __mainwindow struct{}

func (*__mainwindow) init() {}

type MainWindow struct {
	*__mainwindow
	*widgets.QMainWindow
	ActionLatest        *widgets.QAction
	ActionResetApp      *widgets.QAction
	Centralwidget       *widgets.QWidget
	VerticalLayout      *widgets.QVBoxLayout
	HorizontalLayout    *widgets.QHBoxLayout
	HorizontalSpacer    *widgets.QSpacerItem
	Label               *widgets.QLabel
	ProfileSelector     *widgets.QComboBox
	ProfileAddButton    *widgets.QToolButton
	ProfileRenameButton *widgets.QToolButton
	ProfileDeleteButton *widgets.QToolButton
	HorizontalSpacer_2  *widgets.QSpacerItem
	TabWidget           *widgets.QTabWidget
	GridLayout          *widgets.QGridLayout
	CreateProgressText  *widgets.QLabel
	CancelButton        *widgets.QPushButton
	CreateProgress      *widgets.QProgressBar
	CreateStartBtn      *widgets.QPushButton
	Statusbar           *widgets.QStatusBar
}

func NewMainWindow(p widgets.QWidget_ITF) *MainWindow {
	var par *widgets.QWidget
	if p != nil {
		par = p.QWidget_PTR()
	}
	file := core.NewQFile2(":/ui/mainwindow.ui")
	file.Open(core.QIODevice__ReadOnly)
	w := &MainWindow{QMainWindow: widgets.NewQMainWindowFromPointer(uitools.NewQUiLoader(nil).Load(file, par).Pointer())}
	file.Close()
	w.setupUI()
	w.init()
	return w
}
func (w *MainWindow) setupUI() {
	w.TabWidget = widgets.NewQTabWidgetFromPointer(w.FindChild("tabWidget", core.Qt__FindChildrenRecursively).Pointer())
	w.GridLayout = widgets.NewQGridLayoutFromPointer(w.FindChild("gridLayout", core.Qt__FindChildrenRecursively).Pointer())
	w.CancelButton = widgets.NewQPushButtonFromPointer(w.FindChild("cancelButton", core.Qt__FindChildrenRecursively).Pointer())
	w.CreateStartBtn = widgets.NewQPushButtonFromPointer(w.FindChild("createStartBtn", core.Qt__FindChildrenRecursively).Pointer())
	w.Statusbar = widgets.NewQStatusBarFromPointer(w.FindChild("statusbar", core.Qt__FindChildrenRecursively).Pointer())
	w.VerticalLayout = widgets.NewQVBoxLayoutFromPointer(w.FindChild("verticalLayout", core.Qt__FindChildrenRecursively).Pointer())
	w.HorizontalLayout = widgets.NewQHBoxLayoutFromPointer(w.FindChild("horizontalLayout", core.Qt__FindChildrenRecursively).Pointer())
	w.Label = widgets.NewQLabelFromPointer(w.FindChild("label", core.Qt__FindChildrenRecursively).Pointer())
	w.ProfileAddButton = widgets.NewQToolButtonFromPointer(w.FindChild("profileAddButton", core.Qt__FindChildrenRecursively).Pointer())
	w.HorizontalSpacer_2 = widgets.NewQSpacerItemFromPointer(w.FindChild("horizontalSpacer_2", core.Qt__FindChildrenRecursively).Pointer())
	w.ActionLatest = widgets.NewQActionFromPointer(w.FindChild("actionLatest", core.Qt__FindChildrenRecursively).Pointer())
	w.Centralwidget = widgets.NewQWidgetFromPointer(w.FindChild("centralwidget", core.Qt__FindChildrenRecursively).Pointer())
	w.ProfileDeleteButton = widgets.NewQToolButtonFromPointer(w.FindChild("profileDeleteButton", core.Qt__FindChildrenRecursively).Pointer())
	w.CreateProgressText = widgets.NewQLabelFromPointer(w.FindChild("createProgressText", core.Qt__FindChildrenRecursively).Pointer())
	w.CreateProgress = widgets.NewQProgressBarFromPointer(w.FindChild("createProgress", core.Qt__FindChildrenRecursively).Pointer())
	w.ActionResetApp = widgets.NewQActionFromPointer(w.FindChild("actionResetApp", core.Qt__FindChildrenRecursively).Pointer())
	w.HorizontalSpacer = widgets.NewQSpacerItemFromPointer(w.FindChild("horizontalSpacer", core.Qt__FindChildrenRecursively).Pointer())
	w.ProfileSelector = widgets.NewQComboBoxFromPointer(w.FindChild("profileSelector", core.Qt__FindChildrenRecursively).Pointer())
	w.ProfileRenameButton = widgets.NewQToolButtonFromPointer(w.FindChild("profileRenameButton", core.Qt__FindChildrenRecursively).Pointer())
}
