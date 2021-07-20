package main

import (
	"github.com/tomiok/movies-lib/api"
	"log"
)

func main() {
	app := api.Start()

	log.Fatal(app.Listen(":5000"))
}
