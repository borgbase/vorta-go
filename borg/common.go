package borg

import (
	"bufio"
	"errors"
	"log"
	"os/exec"
	"golang.org/x/sync/semaphore"
	"vorta-go/models"

	"vorta-go/app"
)

var borgProcessSlot = semaphore.NewWeighted(1)

type BorgRun struct {
	Bin *BorgBin
	SubCommand string
	SubCommandArgs []string
	ExtraBorgArgs []string
	Repo *models.Repo
	Profile *models.Profile
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

	return nil
}


func (r *BorgRun) Run() {
	cmd := exec.Command(
		"/Users/manu/.pyenv/shims/borg",
		"info", "--debug", "uy5cg8ky@uy5cg8ky.repo.borgbase.com:repo")
	app.CurrentCommand = cmd

	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(stderr)
	go func() {
		for scanner.Scan() {
			app.StatusUpdateChannel <- scanner.Text()
		}
	}()

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	app.StatusUpdateChannel <- "Started Command"

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
	app.StatusUpdateChannel <- "Finished Command"
	app.CurrentCommand = nil
}
