package utils

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"vorta/models"
)

type SchedulerCls struct {
	Cron           *cron.Cron
	AppChan        chan VEvent
	IdToProfileMap map[int]cron.EntryID
}

var Scheduler *SchedulerCls

func InitScheduler(ac chan VEvent) {
	Scheduler = &SchedulerCls{
		AppChan:        ac,
		Cron:           cron.New(),
		IdToProfileMap: make(map[int]cron.EntryID),
	}
	Scheduler.ReloadJobs()
}

func (s *SchedulerCls) ReloadJobs() {
	s.Cron.Stop()
	s.Cron = cron.New()
	pp := []models.Profile{}
	models.DB.Select(&pp, models.SqlAllProfiles)
	for _, p := range pp {
		var cronStr string
		var newJob VortaJob
		switch p.ScheduleMode {
		case "interval":
			newJob = VortaJob{ProfileId: p.Id}
			cronStr = fmt.Sprintf("%d */%d * * *", p.ScheduleIntervalMinutes, p.ScheduleIntervalHours)
		case "fixed":
			newJob = VortaJob{ProfileId: p.Id}
			cronStr = fmt.Sprintf("%d %d * * *", p.ScheduleFixedMinute, p.ScheduleFixedHour)
		default:
			continue
		}
		jobId, err := s.Cron.AddJob(cronStr, newJob)
		if err != nil {
			Log.Error(err)
		}
		s.IdToProfileMap[p.Id] = jobId
		Log.Info("Scheduled job for profile ", p.Name)
	}
	s.Cron.Start()
	Log.Info("Reloaded Scheduler.")
}

func (s *SchedulerCls) NextTimeForProfile(profileId int) string {
	entry, jobExists := s.IdToProfileMap[profileId]
	if !jobExists {
		return "None found"
	}

	ee := s.Cron.Entries()
	for _, e := range ee {
		Log.Info(e, entry, e.ID)
		if e.ID == entry {
			return e.Next.Format("2006-01-02 15:04")
		}
	}
	return "None scheduled" //TODO: could be more elegant
}

type VortaJob struct {
	cron.Job
	ProfileId int
}

func (j VortaJob) Run() {
	Scheduler.AppChan <- VEvent{Topic: "StartBackupxx", Message: string(j.ProfileId)}
}
