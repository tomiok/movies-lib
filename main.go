package main

import (
	"github.com/tomiok/movies-lib/api"
	"github.com/tomiok/movies-lib/config"
	"go.uber.org/zap"
	"os"
)

func main() {
	config.InitLogs()
	api.Migrate()
	app := api.Start()
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	zap.L().Error("", zap.Error(app.Listen(":"+port)))
}
