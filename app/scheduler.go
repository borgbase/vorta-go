package app

import (
	"github.com/robfig/cron/v3"
)

func InitScheduler() {
	c := cron.New()
	c.AddFunc("*/1 * * * *", func() { Log.Info("Every minute") })
	c.Start()
	Log.Info("Started Scheduler.")
}
