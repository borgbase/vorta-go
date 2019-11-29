package models

import (
	"database/sql"
	"os"
	"os/user"
	"strings"
	"time"

	"github.com/gosimple/slug"
)

type Profile struct {
	ID               int            `gorm:"not null;primary_key"`
	Name             string         `grom:"type:varchar(255)"`
	AddedAt          time.Time      `gorm:"column:added_at;not null;default:CURRENT_TIMESTAMP"`
	RepoId           sql.NullInt64  `gorm:"column:repo_id"`
	Repo             Repo           `gorm:"foreignkey:RepoId"`
	SSHKey           sql.NullString `gorm:"column:ssh_key;type:varchar(255)"`
	Compression      string         `gorm:"type:varchar(255);not null;default:'zstd,3'"`
	ExcludePatterns  sql.NullString `gorm:"column:exclude_patterns;default:'*/.DS_Store'"`
	ExcludeIfPresent sql.NullString `gorm:"column:exclude_if_present;default:'.nobackup'"`

	ScheduleMode            string `gorm:"column:schedule_mode;type:varchar(255);not null;default:'off'"`
	ScheduleIntervalHours   int    `gorm:"column:schedule_interval_hours;not null;default:1"`
	ScheduleIntervalMinutes int    `gorm:"column:schedule_interval_minutes;not null;default:24"`
	ScheduleFixedHour       int    `gorm:"column:schedule_fixed_hour;not null;default:17"`
	ScheduleFixedMinute     int    `gorm:"column:schedule_fixed_minute;not null;default:54"`

	ValidationOn    bool `gorm:"column:validation_on;not null;default:1"`
	ValidationWeeks int  `gorm:"column:validation_weeks;not null;default:3"`

	PruneOn         bool           `gorm:"column:prune_on;not null;default:0"`
	PruneHour       int            `gorm:"column:prune_hour;not null;default:2"`
	PruneDay        int            `gorm:"column:prune_day;not null;default:7"`
	PruneWeek       int            `gorm:"column:prune_week;not null;default:4"`
	PruneMonth      int            `gorm:"column:prune_month;not null;default:6"`
	PruneYear       int            `gorm:"column:prune_year;not null;default:2"`
	PruneKeepWithin sql.NullString `gorm:"column:prune_keep_within;type:varchar(255);default:''"`

	NewArchiveName string      `gorm:"column:new_archive_name;type:varchar(255);not null;default:'{hostname}__{profile_slug}-{now}'"`
	PrunePrefix    string      `gorm:"column:prune_prefix;type:varchar(255);not null;default:'{hostname}-{profile_slug}-'"`
	PreBackupCmd   string      `gorm:"column:pre_backup_cmd;type:varchar(255);not null;default:''"`
	PostBackupCmd  string      `gorm:"column:post_backup_cmd;type:varchar(255);not null;default:''"`
	SourceDirs     []SourceDir `gorm:"foreignkey:ProfileId"`
}

//func (p *Profile) GetRepo() *Repo {
//	r := Repo{}
//	DB.Get(&r, SqlRepoById, p.RepoId)
//	return &r
//}

func (Profile) TableName() string {
	return "backupprofilemodel"
}

func (p *Profile) Slug() string {
	return slug.Make(p.Name)
}

func (p *Profile) FormatArchiveName() string {
	// Time formatting: https://stackoverflow.com/a/20234207/3983708
	// TODO: fully support time formatting?
	timeFormat := "2006-01-02T15:04:05"
	hostname, _ := os.Hostname()
	user, _ := user.Current()
	r := strings.NewReplacer(
		"{hostname}", hostname,
		"{profile_id}", string(p.ID),
		"{profile_slug}", p.Slug(),
		"{now}", time.Now().Format(timeFormat),
		"{now:%Y-%m-%dT%H:%M:%S}", time.Now().Format(timeFormat),
		"{utc_now}", time.Now().UTC().Format(timeFormat),
		"{user}", user.Username,
	)

	// Fallback if no archive name is set.
	if p.NewArchiveName != "" {
		return r.Replace(p.NewArchiveName)
	} else {
		return time.Now().Format(timeFormat)
	}
}
