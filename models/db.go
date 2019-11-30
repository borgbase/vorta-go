package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"path"
)

var DB *gorm.DB

func InitDb(dbPath string) {
	var err error
	DB, err = gorm.Open("sqlite3", path.Join(dbPath, "settings.db"))
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	DB.AutoMigrate(&Archive{}, &Profile{}, &Repo{}, &SourceDir{})
	DB.LogMode(true)
	var nProfiles int
	DB.Model(&Profile{}).Count(&nProfiles)

	if nProfiles == 0 {
		var defaultProfile = NewProfile("Default")
		DB.Create(&defaultProfile)
	}
}
