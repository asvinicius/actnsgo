package routes

import (
	"github.com/asvinicius/actnsgo/internal/client"
	"github.com/asvinicius/actnsgo/internal/config"
	"github.com/asvinicius/actnsgo/internal/middleware"
	"github.com/asvinicius/actnsgo/internal/service/token"
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5/pgxpool"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func RegisterRoutes(app *fiber.App, pool *pgxpool.Pool, cfg config.Config) error {

	tokenService, err := token.NewTokenService(cfg.JWT.Secret, cfg.JWT.Expiration)

	if err != nil {
		return err
	}

	statusClient := client.NewClient(cfg.Cartola.CartolaURL)

	if err := AuthRoutes(app, pool, cfg, tokenService); err != nil {
		return err
	}

	protected := app.Group("/api/v1", middleware.RequireAuth(tokenService))

	if err := BankRoutes(protected, pool, cfg); err != nil {
		return err
	}

	if err := BackupRoutes(protected, pool, cfg); err != nil {
		return err
	}

	if err := InfoRoutes(protected, statusClient); err != nil {
		return err
	}

	return nil
}
