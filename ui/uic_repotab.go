package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
)

type __repotab struct{}

func (*__repotab) init() {}

type RepoTab struct {
	*__repotab
	*widgets.QWidget
	VerticalLayout          *widgets.QVBoxLayout
	FormLayout              *widgets.QFormLayout
	Label                   *widgets.QLabel
	GridLayout_2            *widgets.QGridLayout
	Label_5                 *widgets.QLabel
	RepoSelector            *widgets.QComboBox
	RepoRemoveToolbutton    *widgets.QToolButton
	Label_6                 *widgets.QLabel
	GridLayout_4            *widgets.QGridLayout
	SshKeyToClipboardButton *widgets.QToolButton
	SshComboBox             *widgets.QComboBox
	Label_2                 *widgets.QLabel
	Label_7                 *widgets.QLabel
	GridLayout_5            *widgets.QGridLayout
	RepoCompression         *widgets.QComboBox
	Label_4                 *widgets.QLabel
	VerticalSpacer_3        *widgets.QSpacerItem
	RepoStats               *widgets.QFormLayout
	Label_10                *widgets.QLabel
	RepoEncryption          *widgets.QLabel
	Line                    *widgets.QFrame
	Line_2                  *widgets.QFrame
	Label_11                *widgets.QLabel
	SizeOriginal            *widgets.QLabel
	Label_13                *widgets.QLabel
	Label_12                *widgets.QLabel
	SizeCompressed          *widgets.QLabel
	SizeDeduplicated        *widgets.QLabel
}

func NewRepoTab(p widgets.QWidget_ITF) *RepoTab {
	var par *widgets.QWidget
	if p != nil {
		par = p.QWidget_PTR()
	}
	file := core.NewQFile2(":/ui/repotab.ui")
	file.Open(core.QIODevice__ReadOnly)
	w := &RepoTab{QWidget: widgets.NewQWidgetFromPointer(uitools.NewQUiLoader(nil).Load(file, par).Pointer())}
	file.Close()
	w.setupUI()
	w.init()
	return w
}
func (w *RepoTab) setupUI() {
	w.SshComboBox = widgets.NewQComboBoxFromPointer(w.FindChild("sshComboBox", core.Qt__FindChildrenRecursively).Pointer())
	w.GridLayout_5 = widgets.NewQGridLayoutFromPointer(w.FindChild("gridLayout_5", core.Qt__FindChildrenRecursively).Pointer())
	w.SizeCompressed = widgets.NewQLabelFromPointer(w.FindChild("sizeCompressed", core.Qt__FindChildrenRecursively).Pointer())
	w.VerticalLayout = widgets.NewQVBoxLayoutFromPointer(w.FindChild("verticalLayout", core.Qt__FindChildrenRecursively).Pointer())
	w.GridLayout_4 = widgets.NewQGridLayoutFromPointer(w.FindChild("gridLayout_4", core.Qt__FindChildrenRecursively).Pointer())
	w.RepoSelector = widgets.NewQComboBoxFromPointer(w.FindChild("repoSelector", core.Qt__FindChildrenRecursively).Pointer())
	w.RepoCompression = widgets.NewQComboBoxFromPointer(w.FindChild("repoCompression", core.Qt__FindChildrenRecursively).Pointer())
	w.RepoEncryption = widgets.NewQLabelFromPointer(w.FindChild("repoEncryption", core.Qt__FindChildrenRecursively).Pointer())
	w.Label_13 = widgets.NewQLabelFromPointer(w.FindChild("label_13", core.Qt__FindChildrenRecursively).Pointer())
	w.Label_12 = widgets.NewQLabelFromPointer(w.FindChild("label_12", core.Qt__FindChildrenRecursively).Pointer())
	w.SizeDeduplicated = widgets.NewQLabelFromPointer(w.FindChild("sizeDeduplicated", core.Qt__FindChildrenRecursively).Pointer())
	w.SshKeyToClipboardButton = widgets.NewQToolButtonFromPointer(w.FindChild("sshKeyToClipboardButton", core.Qt__FindChildrenRecursively).Pointer())
	w.Label_5 = widgets.NewQLabelFromPointer(w.FindChild("label_5", core.Qt__FindChildrenRecursively).Pointer())
	w.Label_2 = widgets.NewQLabelFromPointer(w.FindChild("label_2", core.Qt__FindChildrenRecursively).Pointer())
	w.Label_4 = widgets.NewQLabelFromPointer(w.FindChild("label_4", core.Qt__FindChildrenRecursively).Pointer())
	w.VerticalSpacer_3 = widgets.NewQSpacerItemFromPointer(w.FindChild("verticalSpacer_3", core.Qt__FindChildrenRecursively).Pointer())
	w.Line = widgets.NewQFrameFromPointer(w.FindChild("line", core.Qt__FindChildrenRecursively).Pointer())
	w.Label_11 = widgets.NewQLabelFromPointer(w.FindChild("label_11", core.Qt__FindChildrenRecursively).Pointer())
	w.GridLayout_2 = widgets.NewQGridLayoutFromPointer(w.FindChild("gridLayout_2", core.Qt__FindChildrenRecursively).Pointer())
	w.Label_6 = widgets.NewQLabelFromPointer(w.FindChild("label_6", core.Qt__FindChildrenRecursively).Pointer())
	w.FormLayout = widgets.NewQFormLayoutFromPointer(w.FindChild("formLayout", core.Qt__FindChildrenRecursively).Pointer())
	w.RepoRemoveToolbutton = widgets.NewQToolButtonFromPointer(w.FindChild("repoRemoveToolbutton", core.Qt__FindChildrenRecursively).Pointer())
	w.Label_7 = widgets.NewQLabelFromPointer(w.FindChild("label_7", core.Qt__FindChildrenRecursively).Pointer())
	w.Label = widgets.NewQLabelFromPointer(w.FindChild("label", core.Qt__FindChildrenRecursively).Pointer())
	w.Line_2 = widgets.NewQFrameFromPointer(w.FindChild("line_2", core.Qt__FindChildrenRecursively).Pointer())
	w.SizeOriginal = widgets.NewQLabelFromPointer(w.FindChild("sizeOriginal", core.Qt__FindChildrenRecursively).Pointer())
	w.Label_10 = widgets.NewQLabelFromPointer(w.FindChild("label_10", core.Qt__FindChildrenRecursively).Pointer())
	w.RepoStats = widgets.NewQFormLayoutFromPointer(w.FindChild("repoStats", core.Qt__FindChildrenRecursively).Pointer())
}
