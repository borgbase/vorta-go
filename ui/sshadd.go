package ui

import (
	"github.com/therecipe/qt/core"
	"os"
	"os/exec"
	"path"
	"vorta/utils"
)

func (d *SshAddDialog) init() {
	d.FormatSelect.AddItem("ED25519 (Recommended)", core.NewQVariant1("ed25519"))
	d.FormatSelect.AddItem("RSA (Legacy)", core.NewQVariant1("rsa"))
	d.FormatSelect.AddItem("ECDSA", core.NewQVariant1("ecdsa"))
	d.FormatSelect.ConnectCurrentIndexChanged(d.setOutputFilename)
	d.setOutputFilename(0)

	d.LengthSelect.AddItem("High (Recommended)", core.NewQVariant1("high"))
	d.LengthSelect.AddItem("Medium", core.NewQVariant1("medium"))

	d.GenerateButton.ConnectClicked(d.generateKey)
	d.CloseButton.ConnectClicked(func(_ bool) {
		d.Close()
	})
}

func (d *SshAddDialog) generateKey(_ bool) {
	format := d.FormatSelect.CurrentData(int(core.Qt__UserRole)).ToString()
	length := d.LengthSelect.CurrentData(int(core.Qt__UserRole)).ToString()
	var lengthInt string
	if format == "rsa" {
		switch length {
		case "high":
			lengthInt = "4096"
		case "medium":
			lengthInt = "2048"
		}
	} else {
		switch length {
		case "high":
			lengthInt = "521"
		case "medium":
			lengthInt = "384"
		}
	}
	sshDir, _ := utils.GetSSHDir()
	keyPath := path.Join(sshDir, d.OutputFileTextBox.Text())
	if _, err := os.Stat(keyPath); err == nil {
		d.Errors.SetText("Key file already exists. Not overwriting.")
		return
	}
	err := exec.Command("ssh-keygen", "-t", format, "-b", lengthInt, "-f", keyPath, "-N", "").Run()
	if err != nil {
		d.Errors.SetText("Error during key generation.")
		utils.Log.Error(err)
		return
	}
	utils.CopyPublicKeyToClipboard(keyPath + ".pub")
	d.Errors.SetText("New key was copied to clipboard and written to " + keyPath)
}

func (d *SshAddDialog) setOutputFilename(_ int) {
	selectedFormat := d.FormatSelect.CurrentData(int(core.Qt__UserRole))
	d.OutputFileTextBox.SetText(path.Join("id_" + selectedFormat.ToString()))
}
