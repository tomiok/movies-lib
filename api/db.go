package api

import (
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var MoviesDB *gorm.DB

func Get() *gorm.DB {
	if MoviesDB == nil {
		MoviesDB = get()
	}
	return MoviesDB
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

	err := db.AutoMigrate(&Review{}, &Movie{})

	if err != nil {
		zap.L().Error("cannot perform migration", zap.Error(err))
		return
	}
	zap.L().Info("migration done")
}
