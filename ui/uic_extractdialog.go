package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
)

type __extractdialog struct{}

func (*__extractdialog) init() {}

type ExtractDialog struct {
	*__extractdialog
	*widgets.QDialog
	VerticalLayout     *widgets.QVBoxLayout
	HorizontalLayout_2 *widgets.QHBoxLayout
	Label_2            *widgets.QLabel
	ArchiveNameLabel   *widgets.QLabel
	HorizontalSpacer_2 *widgets.QSpacerItem
	TreeView           *widgets.QTreeView
	Label              *widgets.QLabel
	HorizontalLayout   *widgets.QHBoxLayout
	HorizontalSpacer   *widgets.QSpacerItem
	CancelButton       *widgets.QPushButton
	ExtractButton      *widgets.QPushButton
}

func NewExtractDialog(p widgets.QWidget_ITF) *ExtractDialog {
	var par *widgets.QWidget
	if p != nil {
		par = p.QWidget_PTR()
	}
	file := core.NewQFile2(":/ui/extractdialog.ui")
	file.Open(core.QIODevice__ReadOnly)
	w := &ExtractDialog{QDialog: widgets.NewQDialogFromPointer(uitools.NewQUiLoader(nil).Load(file, par).Pointer())}
	file.Close()
	w.setupUI()
	w.init()
	return w
}
func (w *ExtractDialog) setupUI() {
	w.Label_2 = widgets.NewQLabelFromPointer(w.FindChild("label_2", core.Qt__FindChildrenRecursively).Pointer())
	w.ArchiveNameLabel = widgets.NewQLabelFromPointer(w.FindChild("archiveNameLabel", core.Qt__FindChildrenRecursively).Pointer())
	w.HorizontalSpacer_2 = widgets.NewQSpacerItemFromPointer(w.FindChild("horizontalSpacer_2", core.Qt__FindChildrenRecursively).Pointer())
	w.Label = widgets.NewQLabelFromPointer(w.FindChild("label", core.Qt__FindChildrenRecursively).Pointer())
	w.HorizontalLayout = widgets.NewQHBoxLayoutFromPointer(w.FindChild("horizontalLayout", core.Qt__FindChildrenRecursively).Pointer())
	w.ExtractButton = widgets.NewQPushButtonFromPointer(w.FindChild("extractButton", core.Qt__FindChildrenRecursively).Pointer())
	w.HorizontalLayout_2 = widgets.NewQHBoxLayoutFromPointer(w.FindChild("horizontalLayout_2", core.Qt__FindChildrenRecursively).Pointer())
	w.TreeView = widgets.NewQTreeViewFromPointer(w.FindChild("treeView", core.Qt__FindChildrenRecursively).Pointer())
	w.HorizontalSpacer = widgets.NewQSpacerItemFromPointer(w.FindChild("horizontalSpacer", core.Qt__FindChildrenRecursively).Pointer())
	w.CancelButton = widgets.NewQPushButtonFromPointer(w.FindChild("cancelButton", core.Qt__FindChildrenRecursively).Pointer())
	w.VerticalLayout = widgets.NewQVBoxLayoutFromPointer(w.FindChild("verticalLayout", core.Qt__FindChildrenRecursively).Pointer())
}
