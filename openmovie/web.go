package openmovie

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type Service struct {
	Search
}

func (s *Service) SearchHandler(ctx *fiber.Ctx) error {
	query := ctx.Query("s")
	res, err := s.ByQueryTitle(query)

	if err != nil {
		fmt.Println(query)
		fmt.Println(err.Error())
		return err
	}

	return ctx.JSON(&res)
}

func (s *Service) TitleHandler(ctx *fiber.Ctx) error {
	title := ctx.Query("t")

	res, err := s.ByTitle(title)

	if err != nil {
		fmt.Println(title)
		fmt.Println(err.Error())
		return err
	}

	return ctx.JSON(&res)
}

func (s *Service) HealthCheckHandler(ctx *fiber.Ctx) error {
	ctx.Status(http.StatusOK)
	return nil
}
