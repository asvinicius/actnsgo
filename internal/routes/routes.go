package routes

import (
	"github.com/asvinicius/actnsgo/internal/config"
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5/pgxpool"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func RegisterRoutes(app *fiber.App, pool *pgxpool.Pool, cfg config.Config) error {

	if err := AuthRoutes(app, pool, cfg); err != nil {
		return err
	}

	return nil
}
