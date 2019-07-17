package borg

import (
	"encoding/json"
	"io"
	"vorta-go/utils"

	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"golang.org/x/sync/semaphore"
	"vorta-go/models"

	"vorta-go/app"
)


var borgProcessSlot = semaphore.NewWeighted(1)

type BorgRun struct {
	Bin *BorgBin
	CommonBorgArgs []string
	SubCommand string
	SubCommandArgs []string
	Repo *models.Repo
	Profile *models.Profile
}

type BorgLogMessage struct {
	LogType string `json:"type"`
	Message string `json:"message"`
	Levelname string `json:"levelname"`
	Name string `json:"name"`
	Time float32 `json:"time"`
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

	r.CommonBorgArgs = append(r.CommonBorgArgs, "--info", "--log-json")
	return nil
}


func (r *BorgRun) Run() {
	mergedArgs := append(r.CommonBorgArgs, r.SubCommand)
	mergedArgs = append(mergedArgs, r.SubCommandArgs...)
	cmd := exec.Command(
		r.Bin.Path,
		mergedArgs...
		)
	//app.CurrentCommand = cmd
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "BORG_PASSPHRASE=xxx")

	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
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
				app.Log.Error(err)
			}
			app.AppChan <- utils.VEvent{Topic: "StatusUpdate", Data: l.Message}
			app.Log.Info(l.Message)
		}
	}()

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	app.AppChan <- utils.VEvent{Topic: "StatusUpdate", Data: "Started Command"}

	var result map[string]interface{}
	if err := json.NewDecoder(stdout).Decode(&result); err != nil {
		log.Fatal(err)
	}

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
	app.AppChan <- utils.VEvent{Topic: "StatusUpdate", Data: "Finished Command"}
	borgProcessSlot.Release(1)

	fmt.Println(result)
}
