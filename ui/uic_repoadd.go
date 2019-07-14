package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
)

type __repoadddialog struct{}

func (*__repoadddialog) init() {}

type RepoAddDialog struct {
	*__repoadddialog
	*widgets.QDialog
	GridLayout                 *widgets.QGridLayout
	TabWidget                  *widgets.QTabWidget
	TabWidgetPage1             *widgets.QWidget
	RepoDataFormLayout         *widgets.QFormLayout
	Title                      *widgets.QLabel
	RepoLabel                  *widgets.QLabel
	HorizontalLayout_3         *widgets.QHBoxLayout
	RepoURL                    *widgets.QLineEdit
	ChooseLocalFolderButton    *widgets.QToolButton
	UseRemoteRepoButton        *widgets.QToolButton
	Label_3                    *widgets.QLabel
	PasswordLineEdit           *widgets.QLineEdit
	TabWidgetPage2             *widgets.QWidget
	AdvancedFormLayout         *widgets.QFormLayout
	Label_4                    *widgets.QLabel
	SshComboBox                *widgets.QComboBox
	EncryptionLabel            *widgets.QLabel
	EncryptionComboBox         *widgets.QComboBox
	Label_31                   *widgets.QLabel
	ExtraBorgArgumentsLineEdit *widgets.QLineEdit
	HorizontalLayout_2         *widgets.QHBoxLayout
	HorizontalSpacer           *widgets.QSpacerItem
	SaveButton                 *widgets.QPushButton
	CloseButton                *widgets.QPushButton
	ErrorText                  *widgets.QLabel
}

func NewRepoAddDialog(p widgets.QWidget_ITF) *RepoAddDialog {
	var par *widgets.QWidget
	if p != nil {
		par = p.QWidget_PTR()
	}
	file := core.NewQFile2(":/ui/repoadd.ui")
	file.Open(core.QIODevice__ReadOnly)
	w := &RepoAddDialog{QDialog: widgets.NewQDialogFromPointer(uitools.NewQUiLoader(nil).Load(file, par).Pointer())}
	file.Close()
	w.setupUI()
	w.init()
	return w
}
func (w *RepoAddDialog) setupUI() {
	w.TabWidget = widgets.NewQTabWidgetFromPointer(w.FindChild("tabWidget", core.Qt__FindChildrenRecursively).Pointer())
	w.TabWidgetPage1 = widgets.NewQWidgetFromPointer(w.FindChild("tabWidgetPage1", core.Qt__FindChildrenRecursively).Pointer())
	w.HorizontalSpacer = widgets.NewQSpacerItemFromPointer(w.FindChild("horizontalSpacer", core.Qt__FindChildrenRecursively).Pointer())
	w.CloseButton = widgets.NewQPushButtonFromPointer(w.FindChild("closeButton", core.Qt__FindChildrenRecursively).Pointer())
	w.RepoDataFormLayout = widgets.NewQFormLayoutFromPointer(w.FindChild("repoDataFormLayout", core.Qt__FindChildrenRecursively).Pointer())
	w.Title = widgets.NewQLabelFromPointer(w.FindChild("title", core.Qt__FindChildrenRecursively).Pointer())
	w.UseRemoteRepoButton = widgets.NewQToolButtonFromPointer(w.FindChild("useRemoteRepoButton", core.Qt__FindChildrenRecursively).Pointer())
	w.Label_3 = widgets.NewQLabelFromPointer(w.FindChild("label_3", core.Qt__FindChildrenRecursively).Pointer())
	w.PasswordLineEdit = widgets.NewQLineEditFromPointer(w.FindChild("passwordLineEdit", core.Qt__FindChildrenRecursively).Pointer())
	w.Label_4 = widgets.NewQLabelFromPointer(w.FindChild("label_4", core.Qt__FindChildrenRecursively).Pointer())
	w.ExtraBorgArgumentsLineEdit = widgets.NewQLineEditFromPointer(w.FindChild("extraBorgArgumentsLineEdit", core.Qt__FindChildrenRecursively).Pointer())
	w.AdvancedFormLayout = widgets.NewQFormLayoutFromPointer(w.FindChild("advancedFormLayout", core.Qt__FindChildrenRecursively).Pointer())
	w.SshComboBox = widgets.NewQComboBoxFromPointer(w.FindChild("sshComboBox", core.Qt__FindChildrenRecursively).Pointer())
	w.EncryptionComboBox = widgets.NewQComboBoxFromPointer(w.FindChild("encryptionComboBox", core.Qt__FindChildrenRecursively).Pointer())
	w.HorizontalLayout_2 = widgets.NewQHBoxLayoutFromPointer(w.FindChild("horizontalLayout_2", core.Qt__FindChildrenRecursively).Pointer())
	w.Label_31 = widgets.NewQLabelFromPointer(w.FindChild("label_31", core.Qt__FindChildrenRecursively).Pointer())
	w.GridLayout = widgets.NewQGridLayoutFromPointer(w.FindChild("gridLayout", core.Qt__FindChildrenRecursively).Pointer())
	w.RepoLabel = widgets.NewQLabelFromPointer(w.FindChild("repoLabel", core.Qt__FindChildrenRecursively).Pointer())
	w.HorizontalLayout_3 = widgets.NewQHBoxLayoutFromPointer(w.FindChild("horizontalLayout_3", core.Qt__FindChildrenRecursively).Pointer())
	w.RepoURL = widgets.NewQLineEditFromPointer(w.FindChild("repoURL", core.Qt__FindChildrenRecursively).Pointer())
	w.ChooseLocalFolderButton = widgets.NewQToolButtonFromPointer(w.FindChild("chooseLocalFolderButton", core.Qt__FindChildrenRecursively).Pointer())
	w.TabWidgetPage2 = widgets.NewQWidgetFromPointer(w.FindChild("tabWidgetPage2", core.Qt__FindChildrenRecursively).Pointer())
	w.EncryptionLabel = widgets.NewQLabelFromPointer(w.FindChild("encryptionLabel", core.Qt__FindChildrenRecursively).Pointer())
	w.SaveButton = widgets.NewQPushButtonFromPointer(w.FindChild("saveButton", core.Qt__FindChildrenRecursively).Pointer())
	w.ErrorText = widgets.NewQLabelFromPointer(w.FindChild("errorText", core.Qt__FindChildrenRecursively).Pointer())
}
