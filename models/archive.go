package models

import (
	"database/sql"
	"time"
)

type Archive struct {
	ID        int             `gorm:"column:id;primary_key"`
	ArchiveID string          `gorm:"column:snapshot_id;type:varchar(255);not null"` // use this as primary key?
	Name      string          `gorm:"column:name;type:varchar(255);not null"`
	RepoID    int             `gorm:"column:repo_id;not null"`
	CreatedAt time.Time       `gorm:"column:time;not null"`
	Duration  sql.NullFloat64 `gorm:"column:duration"`
	Size      sql.NullInt64   `gorm:"column:size"`
}

func (Archive) TableName() string {
	return "archivemodel"
}
