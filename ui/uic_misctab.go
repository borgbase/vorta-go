package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
)

type __misctab struct{}

func (*__misctab) init() {}

type MiscTab struct {
	*__misctab
	*widgets.QWidget
	VerticalLayout_2   *widgets.QVBoxLayout
	CheckboxLayout     *widgets.QVBoxLayout
	VerticalSpacer     *widgets.QSpacerItem
	HorizontalLayout   *widgets.QHBoxLayout
	HorizontalSpacer   *widgets.QSpacerItem
	Label_2            *widgets.QLabel
	VersionLabel       *widgets.QLabel
	Label              *widgets.QLabel
	LogLink            *widgets.QLabel
	HorizontalLayout_2 *widgets.QHBoxLayout
	HorizontalSpacer_2 *widgets.QSpacerItem
	Label_3            *widgets.QLabel
	BorgVersion        *widgets.QLabel
	BorgPath           *widgets.QLabel
}

func NewMiscTab(p widgets.QWidget_ITF) *MiscTab {
	var par *widgets.QWidget
	if p != nil {
		par = p.QWidget_PTR()
	}
	file := core.NewQFile2(":/ui/misctab.ui")
	file.Open(core.QIODevice__ReadOnly)
	w := &MiscTab{QWidget: widgets.NewQWidgetFromPointer(uitools.NewQUiLoader(nil).Load(file, par).Pointer())}
	file.Close()
	w.setupUI()
	w.init()
	return w
}
func (w *MiscTab) setupUI() {
	w.VerticalSpacer = widgets.NewQSpacerItemFromPointer(w.FindChild("verticalSpacer", core.Qt__FindChildrenRecursively).Pointer())
	w.VerticalLayout_2 = widgets.NewQVBoxLayoutFromPointer(w.FindChild("verticalLayout_2", core.Qt__FindChildrenRecursively).Pointer())
	w.Label_2 = widgets.NewQLabelFromPointer(w.FindChild("label_2", core.Qt__FindChildrenRecursively).Pointer())
	w.LogLink = widgets.NewQLabelFromPointer(w.FindChild("logLink", core.Qt__FindChildrenRecursively).Pointer())
	w.HorizontalLayout_2 = widgets.NewQHBoxLayoutFromPointer(w.FindChild("horizontalLayout_2", core.Qt__FindChildrenRecursively).Pointer())
	w.HorizontalSpacer_2 = widgets.NewQSpacerItemFromPointer(w.FindChild("horizontalSpacer_2", core.Qt__FindChildrenRecursively).Pointer())
	w.BorgPath = widgets.NewQLabelFromPointer(w.FindChild("borgPath", core.Qt__FindChildrenRecursively).Pointer())
	w.HorizontalSpacer = widgets.NewQSpacerItemFromPointer(w.FindChild("horizontalSpacer", core.Qt__FindChildrenRecursively).Pointer())
	w.VersionLabel = widgets.NewQLabelFromPointer(w.FindChild("versionLabel", core.Qt__FindChildrenRecursively).Pointer())
	w.Label_3 = widgets.NewQLabelFromPointer(w.FindChild("label_3", core.Qt__FindChildrenRecursively).Pointer())
	w.BorgVersion = widgets.NewQLabelFromPointer(w.FindChild("borgVersion", core.Qt__FindChildrenRecursively).Pointer())
	w.CheckboxLayout = widgets.NewQVBoxLayoutFromPointer(w.FindChild("checkboxLayout", core.Qt__FindChildrenRecursively).Pointer())
	w.HorizontalLayout = widgets.NewQHBoxLayoutFromPointer(w.FindChild("horizontalLayout", core.Qt__FindChildrenRecursively).Pointer())
	w.Label = widgets.NewQLabelFromPointer(w.FindChild("label", core.Qt__FindChildrenRecursively).Pointer())
}
