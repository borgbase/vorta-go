package borg

import (
	"errors"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"vorta/utils"
)

type BorgBin struct {
	Path    string
	Version string
}

func NewBorgBin() (*BorgBin, error) {
	// First check in PATH
	pathBin, err := exec.LookPath("borg")
	if err == nil {
		return &BorgBin{Path: pathBin}, nil
	} else {
		utils.Log.Info("Borg binary not found in path.")
	}

	// Check in Resources folder (macOS)
	ex, err := os.Executable()
	dir := filepath.Dir(ex)
	resourceBin := path.Join(filepath.Dir(dir), "Resources", "borg")
	if _, err := os.Stat(resourceBin); err == nil {
		return &BorgBin{Path: resourceBin}, nil
	} else {
		utils.Log.Info("Borg binary not found in Resources folder.")
	}
	return nil, errors.New("Couldn't find borg binary to use.")
}
