package api

import "github.com/gofiber/fiber/v2"

func Start() *fiber.App{
	app := fiber.New()
	s := Service{
		Search: newOA(),
	}

	app.Get("/", s.HealthCheckHandler)
	app.Get("/query", s.SearchHandler)
	app.Get("/title", s.TitleHandler)

	return app
}
