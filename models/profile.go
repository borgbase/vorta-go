package models

import (
	"database/sql"
	"os"
	"os/user"
	"strings"
	"time"

	"github.com/gosimple/slug"
)

func NewProfile(name string) *Profile {
	profile := Profile{}
	profile.Name = name
	profile.AddedAt = time.Now()
	profile.Compression = "zstd,3"
	profile.ExcludePatterns = sql.NullString{String: "*/.DS_Store", Valid: true}
	profile.ExcludeIfPresent = sql.NullString{String: ".nobackup", Valid: true}
	profile.ScheduleMode = "off"
	profile.ScheduleIntervalHours = 1
	profile.ScheduleIntervalMinutes = 24
	profile.ScheduleFixedHour = 17
	profile.ScheduleFixedMinute = 54
	profile.ValidationOn = true
	profile.ValidationWeeks = 3
	profile.PruneOn = false
	profile.PruneHour = 2
	profile.PruneDay = 7
	profile.PruneWeek = 4
	profile.PruneMonth = 6
	profile.PruneKeepWithin = sql.NullString{String: "", Valid: true}
	profile.NewArchiveName = "{hostname}__{profile_slug}-{now}"
	profile.PrunePrefix = "{hostname}-{profile_slug}-"
	profile.PreBackupCmd = ""
	profile.PostBackupCmd = ""
	return &profile
}

type Profile struct {
	ID               int            `gorm:"not null;primary_key"`
	Name             string         `grom:"type:varchar(255)"`
	AddedAt          time.Time      `gorm:"column:added_at;not null"`
	RepoID           sql.NullInt64  `gorm:"column:repo_id"`
	Repo             Repo           `gorm:"foreignkey:RepoID"`
	SSHKey           sql.NullString `gorm:"column:ssh_key;type:varchar(255)"`
	Compression      string         `gorm:"type:varchar(255);not null"`
	ExcludePatterns  sql.NullString `gorm:"column:exclude_patterns"`
	ExcludeIfPresent sql.NullString `gorm:"column:exclude_if_present"`

	ScheduleMode            string `gorm:"column:schedule_mode;type:varchar(255);not null"`
	ScheduleIntervalHours   int    `gorm:"column:schedule_interval_hours;not null"`
	ScheduleIntervalMinutes int    `gorm:"column:schedule_interval_minutes;not null"`
	ScheduleFixedHour       int    `gorm:"column:schedule_fixed_hour;not null"`
	ScheduleFixedMinute     int    `gorm:"column:schedule_fixed_minute;not null"`

	ValidationOn    bool `gorm:"column:validation_on;not null"`
	ValidationWeeks int  `gorm:"column:validation_weeks;not null"`

	PruneOn         bool           `gorm:"column:prune_on;not null;default:0"`
	PruneHour       int            `gorm:"column:prune_hour;not null;default:2"`
	PruneDay        int            `gorm:"column:prune_day;not null;default:7"`
	PruneWeek       int            `gorm:"column:prune_week;not null;default:4"`
	PruneMonth      int            `gorm:"column:prune_month;not null;default:6"`
	PruneYear       int            `gorm:"column:prune_year;not null;default:2"`
	PruneKeepWithin sql.NullString `gorm:"column:prune_keep_within;type:varchar(255);default:''"`

	NewArchiveName string      `gorm:"column:new_archive_name;type:varchar(255);not null"`
	PrunePrefix    string      `gorm:"column:prune_prefix;type:varchar(255);not null"`
	PreBackupCmd   string      `gorm:"column:pre_backup_cmd;type:varchar(255);not null"`
	PostBackupCmd  string      `gorm:"column:post_backup_cmd;type:varchar(255);not null"`
	SourceDirs     []SourceDir `gorm:"foreignkey:ProfileId"`

	KnownWifis []KnownWifi `gorm:"foreignkey:ProfileID"`
}

func (Profile) TableName() string {
	return "backupprofilemodel"
}

func (p *Profile) Slug() string {
	return slug.Make(p.Name)
}

func (p *Profile) FormatArchiveName(archiveNameTemplate string) string {
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
	if archiveNameTemplate != "" {
		return r.Replace(archiveNameTemplate)
	} else {
		return time.Now().Format(timeFormat)
	}
}
