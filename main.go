package main

import (
	"github.com/tomiok/movies-lib/api"
	"github.com/tomiok/movies-lib/config"
	"github.com/tomiok/movies-lib/movies"
	"github.com/tomiok/movies-lib/storage"
	"go.uber.org/zap"
	"os"
)

func main() {
	config.InitLogs()
	storage.Migrate(false, movies.Movie{}, movies.Review{})
	app := api.Start()
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	zap.L().Error("", zap.Error(app.Listen(":"+port)))
}
