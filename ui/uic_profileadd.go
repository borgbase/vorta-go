package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
)

type __profileadddialog struct{}

func (*__profileadddialog) init() {}

type ProfileAddDialog struct {
	*__profileadddialog
	*widgets.QDialog
	VerticalLayout   *widgets.QVBoxLayout
	ModalTitle       *widgets.QLabel
	Label_2          *widgets.QLabel
	FormLayout       *widgets.QFormLayout
	ProfileNameField *widgets.QLineEdit
	Label_3          *widgets.QLabel
	ErrorText        *widgets.QLabel
	ButtonBox        *widgets.QDialogButtonBox
}

func NewProfileAddDialog(p widgets.QWidget_ITF) *ProfileAddDialog {
	var par *widgets.QWidget
	if p != nil {
		par = p.QWidget_PTR()
	}
	file := core.NewQFile2(":/ui/profileadd.ui")
	file.Open(core.QIODevice__ReadOnly)
	w := &ProfileAddDialog{QDialog: widgets.NewQDialogFromPointer(uitools.NewQUiLoader(nil).Load(file, par).Pointer())}
	file.Close()
	w.setupUI()
	w.init()
	return w
}
func (w *ProfileAddDialog) setupUI() {
	w.ButtonBox = widgets.NewQDialogButtonBoxFromPointer(w.FindChild("buttonBox", core.Qt__FindChildrenRecursively).Pointer())
	w.VerticalLayout = widgets.NewQVBoxLayoutFromPointer(w.FindChild("verticalLayout", core.Qt__FindChildrenRecursively).Pointer())
	w.ModalTitle = widgets.NewQLabelFromPointer(w.FindChild("modalTitle", core.Qt__FindChildrenRecursively).Pointer())
	w.Label_2 = widgets.NewQLabelFromPointer(w.FindChild("label_2", core.Qt__FindChildrenRecursively).Pointer())
	w.FormLayout = widgets.NewQFormLayoutFromPointer(w.FindChild("formLayout", core.Qt__FindChildrenRecursively).Pointer())
	w.ProfileNameField = widgets.NewQLineEditFromPointer(w.FindChild("profileNameField", core.Qt__FindChildrenRecursively).Pointer())
	w.Label_3 = widgets.NewQLabelFromPointer(w.FindChild("label_3", core.Qt__FindChildrenRecursively).Pointer())
	w.ErrorText = widgets.NewQLabelFromPointer(w.FindChild("errorText", core.Qt__FindChildrenRecursively).Pointer())
}
