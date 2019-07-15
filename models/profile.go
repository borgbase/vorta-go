package models

import (
	"time"
)

var (
	SqlAllProfiles = "SELECT * FROM backupprofilemodel ORDER BY name ASC"
	SqlProfileById = "SELECT * FROM repomodel WHERE id=?"
	SqlOneProfile =  "SELECT * FROM repomodel LIMIT 1"
)

type Profile struct {
	Id int
	Name string
	AddedAt time.Time `db:"added_at"`
	RepoId int `db:"repo_id"`
	SSHKey string `db:"ssh_key"`
	Compression string
	ExcludePatterns string `db:"exclude_patterns"`
	ExcludeIfPresent string `db:"exclude_if_present"`

	ScheduleMode string `db:"schedule_mode"`
	ScheduleIntervalHours int `db:"schedule_interval_hours"`
	ScheduleIntervalMinutes int `db:"schedule_interval_minutes"`
	ScheduleFixedHour int `db:"schedule_fixed_hour"`
	ScheduleFixedMinute int `db:"schedule_fixed_minute"`

	ValidationOn bool `db:"validation_on"`
	ValidationWeeks bool `db:"validation_weeks"`

	PruneOn bool `db:"prune_on"`
	PruneHour int `db:"prune_hour"`
	PruneDay int `db:"prune_day"`
	PruneWeek int `db:"prune_week"`
	PruneMonth int `db:"prune_month"`
	PruneYear int `db:"prune_year"`
	PruneKeepWithin string `db:"prune_keep_within"`

	//"{hostname}-{profile_slug}-{now:%Y-%m-%dT%H:%M:%S}"
	NewArchiveName string `db:"new_archive_name"`
	//"{hostname}-{profile_slug}-"
	PrunePrefix string `db:"prune_prefix"`
	PreBackupCmd string `db:"pre_backup_cmd"`
	PostBackupCmd string `db:"post_backup_cmd"`
}
