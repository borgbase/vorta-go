package models

var SqlSchemaVersionSchema = `
CREATE TABLE IF NOT EXISTS "schemaversion"
  (
     "id"         INTEGER NOT NULL PRIMARY KEY,
     "version"    INTEGER NOT NULL,
     "changed_at" DATETIME NOT NULL
  );
INSERT OR IGNORE INTO "schemaversion" VALUES (1, 13, DATETIME('now'));
`
