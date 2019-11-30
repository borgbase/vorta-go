package models

import (
	"database/sql"
	"github.com/zalando/go-keyring"
	"time"
)

type Repo struct {
	ID                 int            `gorm:"column:id;not null;primary_key"`
	Url                string         `gorm:"column:url;not null"`
	AddedAt            time.Time      `gorm:"column:added_at;not null"`
	Encryption         sql.NullString `gorm:"column:encryption;type:varchar(255)"`
	UniqueSize         sql.NullInt64  `gorm:"column:unique_size"`
	UniqueCsize        sql.NullInt64  `gorm:"column:unique_csize"`
	TotalSize          sql.NullInt64  `gorm:"column:total_size"`
	TotalUniqueChunks  sql.NullInt64  `gorm:"column:total_unique_chunks"`
	ExtraBorgArguments sql.NullString `gorm:"column:extra_borg_arguments;type:varchar(255)"`
	Archives           []Archive      `gorm:"foreignkey:RepoID"`
}

func (Repo) TableName() string {
	return "repomodel"
}

func (r *Repo) SetPassword(password string) error {
	return keyring.Set("vorta-repo", r.Url, password)
}

func (r *Repo) GetPassword() (password string, err error) {
	password, err = keyring.Get("vorta-repo", r.Url)
	return
}
