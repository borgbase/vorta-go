package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"syscall"
)

//from https://gist.github.com/davidnewhall/3627895a9fc8fa0affbd747183abca39
func WritePidFile(pidFile string) error {
	if piddata, err := ioutil.ReadFile(pidFile); err == nil {
		if pid, err := strconv.Atoi(string(piddata)); err == nil {
			if process, err := os.FindProcess(pid); err == nil {
				if err := process.Signal(syscall.Signal(0)); err == nil {
					return fmt.Errorf("pid already running: %d", pid)
				}
			}
		}
	}
	return ioutil.WriteFile(pidFile, []byte(fmt.Sprintf("%d", os.Getpid())), 0664)
}
