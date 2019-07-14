package borg

import (
	"bufio"
	"log"
	"os/exec"

	"vorta-go/app"
)

type BorgCommand struct {
	SubCommand string
}

//func (c *BorgCommand) Prepare() {
//
//}

func (c *BorgCommand) Run() {
	cmd := exec.Command("borg", "info", "--debug", "uy5cg8ky@uy5cg8ky.repo.borgbase.com:repo")
	app.App.CurrentCommand = cmd

	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(stderr)
	go func() {
		for scanner.Scan() {
			app.App.StatusUpdateChannel <- scanner.Text()
		}
	}()

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	app.App.StatusUpdateChannel <- "Started Command"

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
	app.App.StatusUpdateChannel <- "Finished Command"
	app.App.CurrentCommand = nil
}
