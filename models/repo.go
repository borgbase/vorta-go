package models

import (
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
	Encryption string
	UniqueSize uint64 `db:"unique_size"`
	UniqueCsize uint64 `db:"unique_csize"`
	TotalSize uint64 `db:"total_size"`
	TotalUniqueChunks uint64 `db:"total_unique_chunks"`
	ExtraBorgArguments string `db:"extra_borg_arguments"`
}
