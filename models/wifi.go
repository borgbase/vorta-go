package models

const SqlWifiSchema = `
CREATE TABLE "wifisettingmodel" (
  "id" integer NOT NULL PRIMARY KEY,
  "ssid" varchar(255) NOT NULL,
  "last_connected" datetime,
  "allowed" integer NOT NULL,
  "profile_id" integer NOT NULL,
  FOREIGN KEY ("profile_id") REFERENCES "backupprofilemodel" ("id")
)`

