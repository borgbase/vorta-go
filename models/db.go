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
	DB, err = sqlx.Open("sqlite3", path.Join(dbPath, "settings.db"))

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	err = DB.Ping()
	if err != nil {
		panic(err)
	}
}
