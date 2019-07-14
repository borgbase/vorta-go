package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
)

type __sourcetab struct{}

func (*__sourcetab) init() {}

type SourceTab struct {
	*__sourcetab
	*widgets.QWidget
	VerticalLayout_2      *widgets.QVBoxLayout
	Label_6               *widgets.QLabel
	HorizontalLayout      *widgets.QHBoxLayout
	SourceFilesWidget     *widgets.QListWidget
	VerticalLayout        *widgets.QVBoxLayout
	SourceAddFolder       *widgets.QPushButton
	SourceAddFile         *widgets.QPushButton
	SourceRemove          *widgets.QPushButton
	GridLayout            *widgets.QGridLayout
	Label                 *widgets.QLabel
	Label_2               *widgets.QLabel
	ExcludePatternsField  *widgets.QPlainTextEdit
	ExcludeIfPresentField *widgets.QPlainTextEdit
}

func NewSourceTab(p widgets.QWidget_ITF) *SourceTab {
	var par *widgets.QWidget
	if p != nil {
		par = p.QWidget_PTR()
	}
	file := core.NewQFile2(":/ui/sourcetab.ui")
	file.Open(core.QIODevice__ReadOnly)
	w := &SourceTab{QWidget: widgets.NewQWidgetFromPointer(uitools.NewQUiLoader(nil).Load(file, par).Pointer())}
	file.Close()
	w.setupUI()
	w.init()
	return w
}
func (w *SourceTab) setupUI() {
	w.SourceFilesWidget = widgets.NewQListWidgetFromPointer(w.FindChild("sourceFilesWidget", core.Qt__FindChildrenRecursively).Pointer())
	w.VerticalLayout = widgets.NewQVBoxLayoutFromPointer(w.FindChild("verticalLayout", core.Qt__FindChildrenRecursively).Pointer())
	w.SourceAddFile = widgets.NewQPushButtonFromPointer(w.FindChild("sourceAddFile", core.Qt__FindChildrenRecursively).Pointer())
	w.SourceRemove = widgets.NewQPushButtonFromPointer(w.FindChild("sourceRemove", core.Qt__FindChildrenRecursively).Pointer())
	w.Label_2 = widgets.NewQLabelFromPointer(w.FindChild("label_2", core.Qt__FindChildrenRecursively).Pointer())
	w.ExcludePatternsField = widgets.NewQPlainTextEditFromPointer(w.FindChild("excludePatternsField", core.Qt__FindChildrenRecursively).Pointer())
	w.ExcludeIfPresentField = widgets.NewQPlainTextEditFromPointer(w.FindChild("excludeIfPresentField", core.Qt__FindChildrenRecursively).Pointer())
	w.VerticalLayout_2 = widgets.NewQVBoxLayoutFromPointer(w.FindChild("verticalLayout_2", core.Qt__FindChildrenRecursively).Pointer())
	w.Label_6 = widgets.NewQLabelFromPointer(w.FindChild("label_6", core.Qt__FindChildrenRecursively).Pointer())
	w.HorizontalLayout = widgets.NewQHBoxLayoutFromPointer(w.FindChild("horizontalLayout", core.Qt__FindChildrenRecursively).Pointer())
	w.SourceAddFolder = widgets.NewQPushButtonFromPointer(w.FindChild("sourceAddFolder", core.Qt__FindChildrenRecursively).Pointer())
	w.GridLayout = widgets.NewQGridLayoutFromPointer(w.FindChild("gridLayout", core.Qt__FindChildrenRecursively).Pointer())
	w.Label = widgets.NewQLabelFromPointer(w.FindChild("label", core.Qt__FindChildrenRecursively).Pointer())
}
