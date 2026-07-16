package routes

import (
	"github.com/asvinicius/actnsgo/internal/config"
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5/pgxpool"
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

func RegisterRoutes(app *fiber.App, pool *pgxpool.Pool, cfg config.Config) error {

	if err := AuthRoutes(app, pool, cfg); err != nil {
		return err
	}

	return nil
}
