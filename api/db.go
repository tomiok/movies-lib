package api

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var SDB *gorm.DB

func Get() *gorm.DB {
	if SDB == nil {
		SDB = get()
	}
	return SDB
}

func get() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("movies.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	return db
}

func Migrate() {
	db := Get()

	err := db.AutoMigrate(&Review{})

	if err != nil {
		log.Println(err.Error())
	}
}
