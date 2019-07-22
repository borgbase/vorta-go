package models

import (
	"database/sql"
	"time"
)

var (
	SqlAllArchivesByRepoId = `SELECT * FROM archivemodel WHERE repo_id=? ORDER BY time DESC`
	SqlCreateArchive = `INSERT INTO archivemodel VALUES (NULL, :snapshot_id, :name, :repo_id, DATETIME('now'), :duration, :size)`
)

var SqlArchiveSchema = `
CREATE TABLE IF NOT EXISTS "archivemodel"
  (
     "id"          INTEGER NOT NULL PRIMARY KEY,
     "snapshot_id" VARCHAR(255) NOT NULL,
     "name"        VARCHAR(255) NOT NULL,
     "repo_id"     INTEGER NOT NULL,
     "time"        DATETIME NOT NULL,
     "duration"    REAL,
     "size"        INTEGER,
     FOREIGN KEY ("repo_id") REFERENCES "repomodel" ("id")
  ) 
`

type Archive struct {
	Id int `db:"id"`
	ArchiveId string `db:"snapshot_id"`  // use this as primary key?
	Name string `db:"name"`
	RepoId int `db:"repo_id"`
	CreatedAt time.Time `db:"time"`
	Duration sql.NullFloat64 `db:"duration"`
	Size sql.NullInt64 `db:"size"`
}
