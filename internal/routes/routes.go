package routes

import (
	"github.com/gofiber/fiber/v3"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func Setup(app *fiber.App) {

	app.Get("/api/v1/health", func(c fiber.Ctx) error {

		response := HealthResponse{
			Status: "ok",
		}

		return c.JSON(response)
	})
}
