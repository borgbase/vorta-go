package utils

import (
	"io/ioutil"
	"os/user"
	"path"
	"strings"
)

func isMiscSSHFile(category string) bool {
	switch category {
	case
		"config",
		"authorized_keys",
		"known_hosts",
		".DS_Store":
		return true
	}
	return false
}

func GetSSHDir() (string, error){
	usr, _ := user.Current()
	return path.Join(usr.HomeDir, ".ssh"), nil
}

func FindSSHKeysInStandardFolder() (keys []string, err error) {
	usr, _ := user.Current()
	files, err := ioutil.ReadDir(path.Join(usr.HomeDir, ".ssh"))
	if err != nil {
		return keys, err
	}

	for _, file := range files {
		if file.IsDir() || strings.HasSuffix(file.Name(), ".pub") || isMiscSSHFile(file.Name()) {
			continue
		}
		keys = append(keys, file.Name())
	}
	return keys, nil
}
