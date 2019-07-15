package models

import (
	"time"
)

var (
	SqlAllRepos = "SELECT * FROM repomodel ORDER BY url ASC"
	SqlRepoById = "SELECT * FROM repomodel WHERE id=?"
)

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
