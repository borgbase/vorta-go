package models

import (
	"database/sql"
	"time"
)

var (
	SqlAllRepos = "SELECT * FROM repomodel ORDER BY url ASC"
	SqlRepoById = "SELECT * FROM repomodel WHERE id=?"
	SqlNewRepo = `INSERT INTO repomodel VALUES (NULL, :url, DATETIME('now'), :encryption, 
					:unique_size, :unique_csize, :total_size, :total_unique_chunks, :extra_borg_arguments)`
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
	Id int `db:"id"`
	Url string `db:"url"`
	AddedAt time.Time `db:"added_at"`
	Encryption sql.NullString `db:"encryption"`
	UniqueSize sql.NullInt64 `db:"unique_size"`
	UniqueCsize sql.NullInt64 `db:"unique_csize"`
	TotalSize sql.NullInt64 `db:"total_size"`
	TotalUniqueChunks sql.NullInt64 `db:"total_unique_chunks"`
	ExtraBorgArguments sql.NullString `db:"extra_borg_arguments"`
}
