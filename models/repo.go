package models

import (
	"time"
)

var (
	sqlAllRepos = "SELECT * FROM repomodel ORDER BY url ASC"
	sqlRepoById = "SELECT * FROM repomodel WHERE id=:id"
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

func QueryAllRepos() []Repo {
	rr := []Repo{}
	err := DB.Select(&rr, sqlAllRepos)
	if err != nil {
		panic("Error while loading data.")
	}
	return rr
}
