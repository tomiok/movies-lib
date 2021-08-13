package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tomiok/movies-lib/openmovie"
)

func Start() *fiber.App{
	app := fiber.New()
	s := openmovie.Service{
		Search: openmovie.NewOA(),
	}

	app.Get("/", s.HealthCheckHandler)
	app.Get("/query", s.SearchHandler)
	app.Get("/title", s.TitleHandler)

	return app
}
