package borg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"io"
	"vorta-go/utils"

	"errors"
	"golang.org/x/sync/semaphore"
	"os"
	"os/exec"
	"vorta-go/models"
)

var (
	borgProcessSlot = semaphore.NewWeighted(1)
	AppEventChan    chan utils.VEvent
)

type BorgRun struct {
	Bin             *BorgBin
	CommonBorgArgs  []string
	SubCommand      string
	SubCommandArgs  []string
	Repo            *models.Repo
	Env             []string
	Profile         *models.Profile
	Result          *simplejson.Json
	PlainTextResult string
}

// TODO: formatting function to print different log types.
type BorgLogMessage struct {
	LogType   string  `json:"type"` //log_message, file_status
	Message   string  `json:"message"`
	Levelname string  `json:"levelname"`
	Name      string  `json:"name"`
	Time      float32 `json:"time"`
	Status    string  `json:"status"`
	Path      string  `json:"path"`
}

func (r *BorgRun) Prepare() error {
	// checks: binary available,
	var err error
	r.Bin, err = NewBorgBin()
	if err != nil {
		return err
	}

	// backup in progress
	if !borgProcessSlot.TryAcquire(1) {
		return errors.New("Backup is already in progress.")
	}

	// Try to get repo password, else set dummy password to avoid prompt.
	password, err := r.Repo.GetPassword()
	if err != nil || password == "" {
		password = "999"
	}

	// TODO: deal with BORG_PASSCOMMAND
	r.Env = os.Environ()
	r.Env = append(r.Env, fmt.Sprintf("BORG_PASSPHRASE=%s", password))

	r.CommonBorgArgs = append(r.CommonBorgArgs, "--info", "--log-json")
	return nil
}

func (r *BorgRun) Run() error {
	mergedArgs := append(r.CommonBorgArgs, r.SubCommand)
	mergedArgs = append(mergedArgs, r.SubCommandArgs...)
	utils.Log.Info("Running command: ", r.Bin.Path, mergedArgs)
	cmd := exec.Command(
		r.Bin.Path,
		mergedArgs...,
	)

	cmd.Env = r.Env

	var stdOutBuf bytes.Buffer
	cmd.Stdout = &stdOutBuf

	stderr, err := cmd.StderrPipe()
	if err != nil {
		utils.Log.Fatal(err)
	}

	scanner := json.NewDecoder(stderr)
	go func() {
		l := BorgLogMessage{}
		for {
			err = scanner.Decode(&l)
			if err == io.EOF {
				return
			}
			if err != nil {
				continue
			}
			if l.LogType == "log_message" {
				AppEventChan <- utils.VEvent{Topic: "StatusUpdate", Message: l.Message}
			} else if l.LogType == "file_status" {
				AppEventChan <- utils.VEvent{Topic: "StatusUpdate", Message: l.Path}
			}
			utils.Log.Info(l.Message, l.Path)
		}
	}()

	if err := cmd.Start(); err != nil {
		utils.Log.Error(err)
	}
	AppEventChan <- utils.VEvent{Topic: "StatusUpdate", Message: "Started Command"}

	err = cmd.Wait()
	borgProcessSlot.Release(1)

	if err != nil { // TODO: return code 1 may only mean missing files. https://golang.org/pkg/os/exec/#ExitError
		utils.Log.Error(err)
		AppEventChan <- utils.VEvent{Topic: "StatusUpdate", Message: "Borg finished with errors."}
		return err
	}
	AppEventChan <- utils.VEvent{Topic: "StatusUpdate", Message: "Finished command without errors"}

	// Try to parse json stdout
	stdOutResult := stdOutBuf.Bytes()
	r.Result, err = simplejson.NewJson(stdOutResult)
	if err != nil {
		utils.Log.Error("Failed parsing JSON.", err)
		r.PlainTextResult = string(stdOutResult)
	}

	return nil
}

func (r *BorgRun) ProcessResult() {
	utils.Log.Error("not implemented")
}
