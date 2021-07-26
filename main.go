package main

import (
	"github.com/tomiok/movies-lib/api"
	"log"
	"os"
)

func main() {
	app := api.Start()
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	log.Fatal(app.Listen(":" + port))
}
