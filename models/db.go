package models

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sqlx.DB

func InitDb() {
	var err error
	DB, err = sqlx.Open("sqlite3", "/Users/manu/Library/Application Support/Vorta/settings.db")
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	DB.Ping()
}

