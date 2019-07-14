package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
)

type __sshadddialog struct{}

func (*__sshadddialog) init() {}

type SshAddDialog struct {
	*__sshadddialog
	*widgets.QDialog
	VerticalLayout_2  *widgets.QVBoxLayout
	VerticalLayout    *widgets.QVBoxLayout
	Label_3           *widgets.QLabel
	FormLayout        *widgets.QFormLayout
	Label             *widgets.QLabel
	FormatSelect      *widgets.QComboBox
	Label_2           *widgets.QLabel
	LengthSelect      *widgets.QComboBox
	Label_4           *widgets.QLabel
	OutputFileTextBox *widgets.QLineEdit
	Label_5           *widgets.QLabel
	Label_6           *widgets.QLabel
	HorizontalLayout  *widgets.QHBoxLayout
	CloseButton       *widgets.QPushButton
	GenerateButton    *widgets.QPushButton
	Errors            *widgets.QLabel
}

func NewSshAddDialog(p widgets.QWidget_ITF) *SshAddDialog {
	var par *widgets.QWidget
	if p != nil {
		par = p.QWidget_PTR()
	}
	file := core.NewQFile2(":/ui/sshadd.ui")
	file.Open(core.QIODevice__ReadOnly)
	w := &SshAddDialog{QDialog: widgets.NewQDialogFromPointer(uitools.NewQUiLoader(nil).Load(file, par).Pointer())}
	file.Close()
	w.setupUI()
	w.init()
	return w
}
func (w *SshAddDialog) setupUI() {
	w.GenerateButton = widgets.NewQPushButtonFromPointer(w.FindChild("generateButton", core.Qt__FindChildrenRecursively).Pointer())
	w.FormatSelect = widgets.NewQComboBoxFromPointer(w.FindChild("formatSelect", core.Qt__FindChildrenRecursively).Pointer())
	w.LengthSelect = widgets.NewQComboBoxFromPointer(w.FindChild("lengthSelect", core.Qt__FindChildrenRecursively).Pointer())
	w.Label_4 = widgets.NewQLabelFromPointer(w.FindChild("label_4", core.Qt__FindChildrenRecursively).Pointer())
	w.VerticalLayout_2 = widgets.NewQVBoxLayoutFromPointer(w.FindChild("verticalLayout_2", core.Qt__FindChildrenRecursively).Pointer())
	w.VerticalLayout = widgets.NewQVBoxLayoutFromPointer(w.FindChild("verticalLayout", core.Qt__FindChildrenRecursively).Pointer())
	w.Label = widgets.NewQLabelFromPointer(w.FindChild("label", core.Qt__FindChildrenRecursively).Pointer())
	w.CloseButton = widgets.NewQPushButtonFromPointer(w.FindChild("closeButton", core.Qt__FindChildrenRecursively).Pointer())
	w.Errors = widgets.NewQLabelFromPointer(w.FindChild("errors", core.Qt__FindChildrenRecursively).Pointer())
	w.OutputFileTextBox = widgets.NewQLineEditFromPointer(w.FindChild("outputFileTextBox", core.Qt__FindChildrenRecursively).Pointer())
	w.Label_5 = widgets.NewQLabelFromPointer(w.FindChild("label_5", core.Qt__FindChildrenRecursively).Pointer())
	w.HorizontalLayout = widgets.NewQHBoxLayoutFromPointer(w.FindChild("horizontalLayout", core.Qt__FindChildrenRecursively).Pointer())
	w.Label_6 = widgets.NewQLabelFromPointer(w.FindChild("label_6", core.Qt__FindChildrenRecursively).Pointer())
	w.Label_3 = widgets.NewQLabelFromPointer(w.FindChild("label_3", core.Qt__FindChildrenRecursively).Pointer())
	w.FormLayout = widgets.NewQFormLayoutFromPointer(w.FindChild("formLayout", core.Qt__FindChildrenRecursively).Pointer())
	w.Label_2 = widgets.NewQLabelFromPointer(w.FindChild("label_2", core.Qt__FindChildrenRecursively).Pointer())
}
