package models

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
