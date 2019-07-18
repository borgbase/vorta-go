package models

import (
	"time"
)

var (
	SqlAllSourcesByProfileId = "SELECT * FROM sourcedirmodel WHERE profile_id=?"
	SqlCountSources = `SELECT count(*) FROM sourcedirmodel WHERE profile_id=? and dir=?`
	SqlInsertSourceDir = `INSERT INTO sourcedirmodel VALUES (NULL, ?, ?, DATETIME('now'))`
	SqlDeleteSourceDir = `DELETE FROM sourcedirmodel WHERE profile_id=? and dir=?`
)

var SqlSourceDirSchema = `
CREATE TABLE IF NOT EXISTS "sourcedirmodel"
  (
     "id"         INTEGER NOT NULL PRIMARY KEY,
     "dir"        VARCHAR(255) NOT NULL,
     "profile_id" INTEGER NOT NULL,
     "added_at"   DATETIME NOT NULL,
     FOREIGN KEY ("profile_id") REFERENCES "backupprofilemodel" ("id")
  );
`

type SourceDir struct {
	Id int `db:"id"`
	Dir string `db:"dir"`
	ProfileId int `db:"profile_id"`
	AddedAt time.Time `db:"added_at"`
}
