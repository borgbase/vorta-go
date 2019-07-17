package models

import (
	"database/sql"
	"time"
)

var (
	SqlAllProfiles = "SELECT * FROM backupprofilemodel ORDER BY name ASC"
	SqlProfileById = "SELECT * FROM backupprofilemodel WHERE id=?"
	SqlCountProfiles = "SELECT count(*) from backupprofilemodel"
	SqlUpdateProfile = `UPDATE backupprofilemodel SET x = y WHERE z = y;`
    SqlNewProfile = `INSERT INTO "backupprofilemodel"
					  VALUES (1, ?, (DATETIME('now')), 14, NULL, 'zstd,3', '*/.DS_Store', '.nobackup', 
							  'off', 1, 24, 17, 54, 1, 3, 0, 2, 7, 4, 6, 2, '', '{hostname}__{profile_slug}-{now:%Y-%m-%dT%H:%M:%S}', 
							  '{hostname}-{profile_slug}-', '', ''
					)`
)

type Profile struct {
	Id int
	Name string
	AddedAt time.Time `db:"added_at"`
	RepoId int `db:"repo_id"`
	SSHKey sql.NullString `db:"ssh_key"`
	Compression string
	ExcludePatterns sql.NullString `db:"exclude_patterns"`
	ExcludeIfPresent sql.NullString `db:"exclude_if_present"`

	ScheduleMode string `db:"schedule_mode"`
	ScheduleIntervalHours int `db:"schedule_interval_hours"`
	ScheduleIntervalMinutes int `db:"schedule_interval_minutes"`
	ScheduleFixedHour int `db:"schedule_fixed_hour"`
	ScheduleFixedMinute int `db:"schedule_fixed_minute"`

	ValidationOn bool `db:"validation_on"`
	ValidationWeeks int `db:"validation_weeks"`

	PruneOn bool `db:"prune_on"`
	PruneHour int `db:"prune_hour"`
	PruneDay int `db:"prune_day"`
	PruneWeek int `db:"prune_week"`
	PruneMonth int `db:"prune_month"`
	PruneYear int `db:"prune_year"`
	PruneKeepWithin sql.NullString `db:"prune_keep_within"`

	NewArchiveName string `db:"new_archive_name"`
	PrunePrefix string `db:"prune_prefix"`
	PreBackupCmd string `db:"pre_backup_cmd"`
	PostBackupCmd string `db:"post_backup_cmd"`
}

func (p *Profile) GetRepo() *Repo {
	r := Repo{}
	DB.Get(&r, SqlRepoById, p.RepoId)
	return &r
}


var SqlProfileSchema = `
CREATE TABLE IF NOT EXISTS "backupprofilemodel"
  (
     "id"                        INTEGER NOT NULL PRIMARY KEY,
     "name"                      VARCHAR(255) NOT NULL,
     "added_at"                  DATETIME NOT NULL,
     "repo_id"                   INTEGER,
     "ssh_key"                   VARCHAR(255),
     "compression"               VARCHAR(255) NOT NULL,
     "exclude_patterns"          TEXT,
     "exclude_if_present"        TEXT,
     "schedule_mode"             VARCHAR(255) NOT NULL,
     "schedule_interval_hours"   INTEGER NOT NULL,
     "schedule_interval_minutes" INTEGER NOT NULL,
     "schedule_fixed_hour"       INTEGER NOT NULL,
     "schedule_fixed_minute"     INTEGER NOT NULL,
     "validation_on"             INTEGER NOT NULL,
     "validation_weeks"          INTEGER NOT NULL,
     "prune_on"                  INTEGER NOT NULL,
     "prune_hour"                INTEGER NOT NULL,
     "prune_day"                 INTEGER NOT NULL,
     "prune_week"                INTEGER NOT NULL,
     "prune_month"               INTEGER NOT NULL,
     "prune_year"                INTEGER NOT NULL,
     "prune_keep_within"         VARCHAR(255),
     "new_archive_name"          VARCHAR(255) NOT NULL,
     "prune_prefix"              VARCHAR(255) NOT NULL,
     "pre_backup_cmd"            VARCHAR(255) NOT NULL,
     "post_backup_cmd"           VARCHAR(255) NOT NULL,
     FOREIGN KEY ("repo_id") REFERENCES "repomodel" ("id")
  );
`

