package app

import (
	"github.com/robfig/cron/v3"
)

func InitScheduler() {
	c := cron.New()
	c.AddFunc("*/1 * * * *", func() { App.Log.Info("Every minute") })
	c.Start()
	App.Log.Info("Started Scheduler.")
}
