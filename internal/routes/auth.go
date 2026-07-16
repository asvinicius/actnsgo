package routes

import (
	authhandler "github.com/asvinicius/actnsgo/internal/handler/auth"
	"github.com/asvinicius/actnsgo/internal/repository"
	authservice "github.com/asvinicius/actnsgo/internal/service/auth"
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5/pgxpool"
)

func AuthRoutes(app *fiber.App, pool *pgxpool.Pool) {

	rep := repository.NewSuperRepository(pool)
	serv := authservice.NewAuthService(rep)
	hand := authhandler.NewAuthHandler(serv)

	app.Post("/api/v1/auth/super/login", hand.Login)
}
