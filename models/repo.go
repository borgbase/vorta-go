package models

import (
	"database/sql"
	"time"
)

var (
	SqlAllRepos = "SELECT * FROM repomodel ORDER BY url ASC"
	SqlRepoById = "SELECT * FROM repomodel WHERE id=?"
)

var SqlRepoSchema = `
CREATE TABLE IF NOT EXISTS "repomodel"
  (
     "id"                   INTEGER NOT NULL PRIMARY KEY,
     "url"                  VARCHAR(255) NOT NULL,
     "added_at"             DATETIME NOT NULL,
     "encryption"           VARCHAR(255),
     "unique_size"          INTEGER,
     "unique_csize"         INTEGER,
     "total_size"           INTEGER,
     "total_unique_chunks"  INTEGER,
     "extra_borg_arguments" VARCHAR(255)
  ) `

type Repo struct {
	Id int
	Url string
	AddedAt time.Time `db:"added_at"`
	Encryption sql.NullString
	UniqueSize sql.NullInt64 `db:"unique_size"`
	UniqueCsize sql.NullInt64 `db:"unique_csize"`
	TotalSize sql.NullInt64 `db:"total_size"`
	TotalUniqueChunks sql.NullInt64 `db:"total_unique_chunks"`
	ExtraBorgArguments sql.NullString `db:"extra_borg_arguments"`
}
