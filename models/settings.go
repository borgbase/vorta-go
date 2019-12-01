package models

var SqlSettingsSchema = `
CREATE TABLE IF NOT EXISTS "settingsmodel"
  (
     "id"    INTEGER NOT NULL PRIMARY KEY,
     "key"   VARCHAR(255) NOT NULL,
     "value" INTEGER NOT NULL,
     "label" VARCHAR(255) NOT NULL,
     "type"  VARCHAR(255) NOT NULL
  );
CREATE UNIQUE INDEX IF NOT EXISTS "settingsmodel_key" ON "settingsmodel" ("key");
INSERT OR IGNORE INTO "settingsmodel" VALUES (NULL, 'use_light_icon', 0, 'Use light system tray icon (applies after restart)', 'checkbox');
INSERT OR IGNORE INTO "settingsmodel" VALUES (NULL, 'enable_notifications', 1, 'Display notifications when background tasks fail', 'checkbox');
INSERT OR IGNORE INTO "settingsmodel" VALUES (NULL, 'check_for_updates', 1, 'Check for updates on startup', 'checkbox');
INSERT OR IGNORE INTO "settingsmodel" VALUES (NULL, 'enable_notifications_success', 0, 'Also notify about successful background tasks', 'checkbox');
INSERT OR IGNORE INTO "settingsmodel" VALUES (NULL, 'use_dark_theme', 0, 'Use dark theme (applies after restart)', 'checkbox');
`
