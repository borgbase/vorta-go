package models

import (
	"fmt"
	"time"
)

var (
	sqlAllRepos = "SELECT * FROM repomodel"
	sqlRepoById = "SELECT * FROM repomodel WHERE id=:id"
)

type Repo struct {
	Id uint
	Url string
	AddedAt time.Time `db:"added_at"`
	Encryption string
	UniqueSize uint64 `db:"unique_size"`
	UniqueCsize uint64 `db:"unique_csize"`
	TotalSize uint64 `db:"total_size"`
	TotalUniqueChunks uint64 `db:"total_unique_chunks"`
	ExtraBorgArguments string `db:"extra_borg_arguments"`
}

func QueryAllRepos() {
	rr := []Repo{}
	err := DB.Select(&rr, sqlAllRepos)
	if err != nil {
		panic("Error while loading data.")
	}
	for _, repo := range rr {

		fmt.Println("url: %s", repo.Url)
	}
}
