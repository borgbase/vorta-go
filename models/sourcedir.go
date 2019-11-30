package models

import (
	"time"
)

type SourceDir struct {
	ID        int       `gorm:"column:id;not null;primary_key"`
	Dir       string    `gorm:"column:dir;type:varchar(255);not null"`
	ProfileId int       `gorm:"column:profile_id;not null"`
	AddedAt   time.Time `gorm:"column:added_at;not null"`
}

func (SourceDir) TableName() string {
	return "sourcedirmodel"
}
