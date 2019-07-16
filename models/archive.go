package models

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
