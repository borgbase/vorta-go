package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
)

type __scheduletab struct{}

func (*__scheduletab) init() {}

type ScheduleTab struct {
	*__scheduletab
	*widgets.QWidget
	GridLayout              *widgets.QGridLayout
	ToolBox                 *widgets.QToolBox
	Schedule                *widgets.QWidget
	VerticalLayout          *widgets.QVBoxLayout
	HorizontalLayout_2      *widgets.QHBoxLayout
	ScheduleOffRadio        *widgets.QRadioButton
	HorizontalLayout_3      *widgets.QHBoxLayout
	ScheduleIntervalRadio   *widgets.QRadioButton
	ScheduleIntervalHours   *widgets.QSpinBox
	Label                   *widgets.QLabel
	ScheduleIntervalMinutes *widgets.QSpinBox
	Label_2                 *widgets.QLabel
	HorizontalSpacer_2      *widgets.QSpacerItem
	HorizontalLayout_5      *widgets.QHBoxLayout
	ScheduleFixedRadio      *widgets.QRadioButton
	ScheduleFixedTime       *widgets.QTimeEdit
	HorizontalSpacer        *widgets.QSpacerItem
	HorizontalLayout_6      *widgets.QHBoxLayout
	ValidationCheckBox      *widgets.QCheckBox
	ValidationSpinBox       *widgets.QSpinBox
	Label_5                 *widgets.QLabel
	HorizontalSpacer_4      *widgets.QSpacerItem
	HorizontalLayout_8      *widgets.QHBoxLayout
	PruneCheckBox           *widgets.QCheckBox
	HorizontalSpacer_6      *widgets.QSpacerItem
	Line_2                  *widgets.QFrame
	HorizontalLayout        *widgets.QHBoxLayout
	ScheduleApplyButton     *widgets.QPushButton
	HorizontalSpacer_5      *widgets.QSpacerItem
	Label_3                 *widgets.QLabel
	NextBackupDateTimeLabel *widgets.QLabel
	HorizontalSpacer_3      *widgets.QSpacerItem
	VerticalSpacer          *widgets.QSpacerItem
	Page_2                  *widgets.QWidget
	GridLayout_2            *widgets.QGridLayout
	WifiListWidget          *widgets.QListWidget
	WifiListLabel           *widgets.QLabel
	Page                    *widgets.QWidget
	GridLayout_3            *widgets.QGridLayout
	LogTableWidget          *widgets.QTableWidget
	Page_3                  *widgets.QWidget
	VerticalLayout_2        *widgets.QVBoxLayout
	Label_4                 *widgets.QLabel
	PreBackupCmdLineEdit    *widgets.QLineEdit
	PostBackupCmdLineEdit   *widgets.QLineEdit
	Label_6                 *widgets.QLabel
	VerticalSpacer_2        *widgets.QSpacerItem
}

func NewScheduleTab(p widgets.QWidget_ITF) *ScheduleTab {
	var par *widgets.QWidget
	if p != nil {
		par = p.QWidget_PTR()
	}
	file := core.NewQFile2(":/ui/scheduletab.ui")
	file.Open(core.QIODevice__ReadOnly)
	w := &ScheduleTab{QWidget: widgets.NewQWidgetFromPointer(uitools.NewQUiLoader(nil).Load(file, par).Pointer())}
	file.Close()
	w.setupUI()
	w.init()
	return w
}
func (w *ScheduleTab) setupUI() {
	w.ScheduleIntervalHours = widgets.NewQSpinBoxFromPointer(w.FindChild("scheduleIntervalHours", core.Qt__FindChildrenRecursively).Pointer())
	w.HorizontalSpacer_5 = widgets.NewQSpacerItemFromPointer(w.FindChild("horizontalSpacer_5", core.Qt__FindChildrenRecursively).Pointer())
	w.NextBackupDateTimeLabel = widgets.NewQLabelFromPointer(w.FindChild("nextBackupDateTimeLabel", core.Qt__FindChildrenRecursively).Pointer())
	w.Page = widgets.NewQWidgetFromPointer(w.FindChild("page", core.Qt__FindChildrenRecursively).Pointer())
	w.ScheduleIntervalRadio = widgets.NewQRadioButtonFromPointer(w.FindChild("scheduleIntervalRadio", core.Qt__FindChildrenRecursively).Pointer())
	w.Label = widgets.NewQLabelFromPointer(w.FindChild("label", core.Qt__FindChildrenRecursively).Pointer())
	w.Line_2 = widgets.NewQFrameFromPointer(w.FindChild("line_2", core.Qt__FindChildrenRecursively).Pointer())
	w.Page_3 = widgets.NewQWidgetFromPointer(w.FindChild("page_3", core.Qt__FindChildrenRecursively).Pointer())
	w.Label_4 = widgets.NewQLabelFromPointer(w.FindChild("label_4", core.Qt__FindChildrenRecursively).Pointer())
	w.Label_6 = widgets.NewQLabelFromPointer(w.FindChild("label_6", core.Qt__FindChildrenRecursively).Pointer())
	w.PreBackupCmdLineEdit = widgets.NewQLineEditFromPointer(w.FindChild("preBackupCmdLineEdit", core.Qt__FindChildrenRecursively).Pointer())
	w.Label_2 = widgets.NewQLabelFromPointer(w.FindChild("label_2", core.Qt__FindChildrenRecursively).Pointer())
	w.ScheduleFixedTime = widgets.NewQTimeEditFromPointer(w.FindChild("scheduleFixedTime", core.Qt__FindChildrenRecursively).Pointer())
	w.Label_5 = widgets.NewQLabelFromPointer(w.FindChild("label_5", core.Qt__FindChildrenRecursively).Pointer())
	w.HorizontalSpacer_6 = widgets.NewQSpacerItemFromPointer(w.FindChild("horizontalSpacer_6", core.Qt__FindChildrenRecursively).Pointer())
	w.HorizontalLayout = widgets.NewQHBoxLayoutFromPointer(w.FindChild("horizontalLayout", core.Qt__FindChildrenRecursively).Pointer())
	w.GridLayout_2 = widgets.NewQGridLayoutFromPointer(w.FindChild("gridLayout_2", core.Qt__FindChildrenRecursively).Pointer())
	w.VerticalLayout_2 = widgets.NewQVBoxLayoutFromPointer(w.FindChild("verticalLayout_2", core.Qt__FindChildrenRecursively).Pointer())
	w.ScheduleOffRadio = widgets.NewQRadioButtonFromPointer(w.FindChild("scheduleOffRadio", core.Qt__FindChildrenRecursively).Pointer())
	w.ScheduleIntervalMinutes = widgets.NewQSpinBoxFromPointer(w.FindChild("scheduleIntervalMinutes", core.Qt__FindChildrenRecursively).Pointer())
	w.ScheduleFixedRadio = widgets.NewQRadioButtonFromPointer(w.FindChild("scheduleFixedRadio", core.Qt__FindChildrenRecursively).Pointer())
	w.HorizontalLayout_8 = widgets.NewQHBoxLayoutFromPointer(w.FindChild("horizontalLayout_8", core.Qt__FindChildrenRecursively).Pointer())
	w.Label_3 = widgets.NewQLabelFromPointer(w.FindChild("label_3", core.Qt__FindChildrenRecursively).Pointer())
	w.VerticalSpacer = widgets.NewQSpacerItemFromPointer(w.FindChild("verticalSpacer", core.Qt__FindChildrenRecursively).Pointer())
	w.GridLayout = widgets.NewQGridLayoutFromPointer(w.FindChild("gridLayout", core.Qt__FindChildrenRecursively).Pointer())
	w.HorizontalSpacer_4 = widgets.NewQSpacerItemFromPointer(w.FindChild("horizontalSpacer_4", core.Qt__FindChildrenRecursively).Pointer())
	w.HorizontalSpacer_3 = widgets.NewQSpacerItemFromPointer(w.FindChild("horizontalSpacer_3", core.Qt__FindChildrenRecursively).Pointer())
	w.GridLayout_3 = widgets.NewQGridLayoutFromPointer(w.FindChild("gridLayout_3", core.Qt__FindChildrenRecursively).Pointer())
	w.ToolBox = widgets.NewQToolBoxFromPointer(w.FindChild("toolBox", core.Qt__FindChildrenRecursively).Pointer())
	w.HorizontalSpacer_2 = widgets.NewQSpacerItemFromPointer(w.FindChild("horizontalSpacer_2", core.Qt__FindChildrenRecursively).Pointer())
	w.HorizontalLayout_6 = widgets.NewQHBoxLayoutFromPointer(w.FindChild("horizontalLayout_6", core.Qt__FindChildrenRecursively).Pointer())
	w.PruneCheckBox = widgets.NewQCheckBoxFromPointer(w.FindChild("pruneCheckBox", core.Qt__FindChildrenRecursively).Pointer())
	w.ScheduleApplyButton = widgets.NewQPushButtonFromPointer(w.FindChild("scheduleApplyButton", core.Qt__FindChildrenRecursively).Pointer())
	w.LogTableWidget = widgets.NewQTableWidgetFromPointer(w.FindChild("logTableWidget", core.Qt__FindChildrenRecursively).Pointer())
	w.Schedule = widgets.NewQWidgetFromPointer(w.FindChild("schedule", core.Qt__FindChildrenRecursively).Pointer())
	w.HorizontalLayout_2 = widgets.NewQHBoxLayoutFromPointer(w.FindChild("horizontalLayout_2", core.Qt__FindChildrenRecursively).Pointer())
	w.HorizontalSpacer = widgets.NewQSpacerItemFromPointer(w.FindChild("horizontalSpacer", core.Qt__FindChildrenRecursively).Pointer())
	w.ValidationSpinBox = widgets.NewQSpinBoxFromPointer(w.FindChild("validationSpinBox", core.Qt__FindChildrenRecursively).Pointer())
	w.Page_2 = widgets.NewQWidgetFromPointer(w.FindChild("page_2", core.Qt__FindChildrenRecursively).Pointer())
	w.WifiListWidget = widgets.NewQListWidgetFromPointer(w.FindChild("wifiListWidget", core.Qt__FindChildrenRecursively).Pointer())
	w.WifiListLabel = widgets.NewQLabelFromPointer(w.FindChild("wifiListLabel", core.Qt__FindChildrenRecursively).Pointer())
	w.PostBackupCmdLineEdit = widgets.NewQLineEditFromPointer(w.FindChild("postBackupCmdLineEdit", core.Qt__FindChildrenRecursively).Pointer())
	w.VerticalLayout = widgets.NewQVBoxLayoutFromPointer(w.FindChild("verticalLayout", core.Qt__FindChildrenRecursively).Pointer())
	w.HorizontalLayout_3 = widgets.NewQHBoxLayoutFromPointer(w.FindChild("horizontalLayout_3", core.Qt__FindChildrenRecursively).Pointer())
	w.HorizontalLayout_5 = widgets.NewQHBoxLayoutFromPointer(w.FindChild("horizontalLayout_5", core.Qt__FindChildrenRecursively).Pointer())
	w.ValidationCheckBox = widgets.NewQCheckBoxFromPointer(w.FindChild("validationCheckBox", core.Qt__FindChildrenRecursively).Pointer())
	w.VerticalSpacer_2 = widgets.NewQSpacerItemFromPointer(w.FindChild("verticalSpacer_2", core.Qt__FindChildrenRecursively).Pointer())
}
