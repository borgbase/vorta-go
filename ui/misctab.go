package ui

var version string

func (w *MiscTab) init() {
	w.VersionLabel.SetText(version)
}
