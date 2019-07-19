package borg

import (
	"bytes"
	"encoding/json"
	"github.com/zalando/go-keyring"
	"github.com/bitly/go-simplejson"
	"io"
	"fmt"
	"os/user"
	"strings"
	"time"
	"vorta-go/utils"

	"errors"
	"golang.org/x/sync/semaphore"
	"os"
	"os/exec"
	"vorta-go/models"
)


var (
	borgProcessSlot = semaphore.NewWeighted(1)
	AppEventChan chan utils.VEvent
)


type BorgRun struct {
	Bin *BorgBin
	CommonBorgArgs []string
	SubCommand string
	SubCommandArgs []string
	Repo *models.Repo
	RepoPassword string
	Env []string
	Profile *models.Profile
	Result *simplejson.Json
	PlainTextResult string
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

	// Set password if not yet defined
	if r.RepoPassword == "" {
		secret, err := keyring.Get("vorta-repo", r.Repo.Url)
		if err != nil {
			utils.Log.Error(err)
		} else {
			r.RepoPassword = secret
		}
	}

	// TODO: deal with BORG_PASSCOMMAND
	r.Env = os.Environ()
	r.Env = append(r.Env, fmt.Sprintf("BORG_PASSPHRASE=%s", r.RepoPassword))

	r.CommonBorgArgs = append(r.CommonBorgArgs, "--info", "--log-json")
	return nil
}


func (r *BorgRun) Run() error {
	mergedArgs := append(r.CommonBorgArgs, r.SubCommand)
	mergedArgs = append(mergedArgs, r.SubCommandArgs...)
	utils.Log.Info("Running command: ", r.Bin.Path, mergedArgs)
	cmd := exec.Command(
		r.Bin.Path,
		mergedArgs...
		)

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
			AppEventChan <- utils.VEvent{Topic: "StatusUpdate", Message: l.Message}
			utils.Log.Info(l.Message)
		}
	}()

	if err := cmd.Start(); err != nil {
		utils.Log.Error(err)
	}
	AppEventChan <- utils.VEvent{Topic: "StatusUpdate", Message: "Started Command"}

	err = cmd.Wait()
	borgProcessSlot.Release(1)

	if err != nil {
		utils.Log.Error(err)
		AppEventChan <- utils.VEvent{Topic: "StatusUpdate", Message: "Borg finished with errors."}
		return err
	}
	AppEventChan <- utils.VEvent{Topic: "StatusUpdate", Message: "Finished Command"}

	// Try to parse json stdout
	stdOutResult := stdOutBuf.Bytes()
	r.Result, err= simplejson.NewJson(stdOutResult)
	if err != nil {
		utils.Log.Error("Failed parsing JSON.", err)
		r.PlainTextResult = string(stdOutResult)
	}

	return nil
}

func (r *BorgRun) ProcessResult(result map[string]interface{}) {}

func _formatArchiveName(p *models.Profile) string {
	// Time formatting: https://stackoverflow.com/a/20234207/3983708
	// TODO: fully support time formatting?
	timeFormat := "2006-01-02T15:04:05"
	hostname, _ := os.Hostname()
	user, _ := user.Current()
	r := strings.NewReplacer(
		"{hostname}", hostname,
		"{profile_id}", string(p.Id),
		"{profile_slug}", p.Slug(),
		"{now}", time.Now().Format(timeFormat),
		"{now:%Y-%m-%dT%H:%M:%S}", time.Now().Format(timeFormat),
		"{utc_now}", time.Now().UTC().Format(timeFormat),
		"{user}", user.Username,
		)
	return r.Replace(p.NewArchiveName)
}
