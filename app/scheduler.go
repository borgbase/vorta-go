package app

import (
	"github.com/robfig/cron/v3"
	"vorta-go/utils"
)

func InitScheduler() {
	c := cron.New()
	c.AddFunc("*/1 * * * *", func() { utils.Log.Info("Every minute") })
	c.Start()
	utils.Log.Info("Started Scheduler.")
}
