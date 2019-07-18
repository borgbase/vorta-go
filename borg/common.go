package borg

import (
	"encoding/json"
	"io"

	//"io"
	"os/user"
	"strings"
	"time"
	"vorta-go/utils"

	"errors"
	"fmt"
	"os"
	"os/exec"
	"golang.org/x/sync/semaphore"
	"github.com/zalando/go-keyring"
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
	// Get Repo object from Profile
	r.Repo = &models.Repo{}
	models.DB.Get(r.Repo, models.SqlRepoById, r.Profile.RepoId)
	utils.Log.Info("Backing up to repo", r.Repo.Url)

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
	utils.Log.Info(mergedArgs)
	cmd := exec.Command(
		r.Bin.Path,
		mergedArgs...
		)

	secret, err := keyring.Get("vorta-repo", r.Repo.Url)
	if err != nil {
		utils.Log.Fatal(err)
	}

	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, fmt.Sprintf("BORG_PASSPHRASE=%s", secret))


	stdout, err := cmd.StdoutPipe()
	if err != nil {
		utils.Log.Fatal(err)
	}

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
				utils.Log.Error(err)
				utils.Log.Error(l)
				continue
			}
			AppEventChan <- utils.VEvent{Topic: "StatusUpdate", Message: l.Message}
			utils.Log.Info(l.Message)
		}
	}()

	if err := cmd.Start(); err != nil {
		utils.Log.Fatal(err)
	}
	AppEventChan <- utils.VEvent{Topic: "StatusUpdate", Message: "Started Command"}

	var result map[string]interface{}
	if err := json.NewDecoder(stdout).Decode(&result); err != nil {
		utils.Log.Fatal(err)
	}

	if err := cmd.Wait(); err != nil {
		utils.Log.Fatal(err)
	}
	AppEventChan <- utils.VEvent{Topic: "StatusUpdate", Message: "Finished Command"}
	borgProcessSlot.Release(1)

	fmt.Println(result)
}

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
