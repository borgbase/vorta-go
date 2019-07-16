package models

var SqlEvenLogSchema = `
CREATE TABLE IF NOT EXISTS "eventlogmodel"
  (
     "id"         INTEGER NOT NULL PRIMARY KEY,
     "start_time" DATETIME NOT NULL,
     "category"   VARCHAR(255) NOT NULL,
     "subcommand" VARCHAR(255),
     "message"    VARCHAR(255),
     "returncode" INTEGER NOT NULL,
     "params"     TEXT,
     "repo_url"   VARCHAR(255),
     "profile"    VARCHAR(255)
  ) `
