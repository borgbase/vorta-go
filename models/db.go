package models

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"path"
)

var DB *sqlx.DB

func InitDb(dbPath string) {
	var err error
	DB, err = sqlx.Connect("sqlite3", path.Join(dbPath, "settings.db"))

	DB.MustExec(SqlArchiveSchema)
	DB.MustExec(SqlEvenLogSchema)
	DB.MustExec(SqlProfileSchema)
	DB.MustExec(SqlRepoSchema)
	DB.MustExec(SqlSchemaVersionSchema)
	DB.MustExec(SqlSettingsSchema)
	DB.MustExec(SqlSourceDirSchema)

	var nProfiles int
	DB.Get(&nProfiles, SqlCountProfiles)
	if nProfiles == 0 {
		DB.MustExec(SqlNewProfile, "Default")
	}

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
}
