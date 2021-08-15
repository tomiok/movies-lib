package storage

import (
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var MoviesDB *gorm.DB

func Get(isTesting bool) *gorm.DB {
	if MoviesDB == nil {
		MoviesDB = get(isTesting)
	}
	return MoviesDB
}

func get(isTesting bool) *gorm.DB {
	var dsn string

	if isTesting {
		dsn = "movies-testing.db"
	} else {
		dsn = "movies.db"
	}

	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func Migrate(isTesting bool, dest ...interface{}) {
	db := Get(isTesting)
	if isTesting {
		err := db.Migrator().DropTable(dest...)
		if err != nil {
			zap.L().Warn("cannot drop tables", zap.Error(err))
		}
	}
	err := db.AutoMigrate(dest...)

	if err != nil {
		zap.L().Error("cannot perform migration", zap.Error(err))
		return
	}
	zap.L().Info("migration done")
}
