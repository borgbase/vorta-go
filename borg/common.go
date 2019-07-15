package borg

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"

	"vorta-go/app"
)

type BorgCommand struct {
	SubCommand string
}

func (c *BorgCommand) Prepare() {
	path, err := exec.LookPath("borg")
	if err != nil {
		log.Fatal("Borg binary not found.")
	}
	fmt.Printf("borg is available at %s\n", path)
}

func (c *BorgCommand) Run() {
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
