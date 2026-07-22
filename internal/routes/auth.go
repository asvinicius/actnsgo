package routes

import (
	"github.com/asvinicius/actnsgo/internal/config"
	authhandler "github.com/asvinicius/actnsgo/internal/handler/auth"
	"github.com/asvinicius/actnsgo/internal/repository"
	authservice "github.com/asvinicius/actnsgo/internal/service/auth"
	"github.com/asvinicius/actnsgo/internal/service/token"
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5/pgxpool"
)

func AuthRoutes(app *fiber.App, pool *pgxpool.Pool, cfg config.Config) error {

	tokenService, err := token.NewTokenService(
		cfg.JWT.Secret,
		cfg.JWT.Expiration,
	)

	if err != nil {
		return err
	}

	rep := repository.NewSuperRepository(pool)
	serv := authservice.NewAuthService(rep, tokenService)
	hand := authhandler.NewAuthHandler(serv)

	app.Post("/api/v1/auth/super/login", hand.Login)
	app.Get("/api/v1/auth/super/islogged", hand.IsLogged)

	return nil
}
