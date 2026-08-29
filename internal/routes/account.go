package routes

import (
	"github.com/asvinicius/actnsgo/internal/config"
	financehandler "github.com/asvinicius/actnsgo/internal/handler/finance"
	"github.com/asvinicius/actnsgo/internal/repository"
	financeservice "github.com/asvinicius/actnsgo/internal/service/finance"
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5/pgxpool"
)

func AccountRoutes(router fiber.Router, pool *pgxpool.Pool, cfg config.Config) error {
	rep := repository.NewAccountRepository(pool)
	serv := financeservice.NewAccountService(rep)
	hand := financehandler.NewAccountHandler(serv)

	router.Post("/account/create", hand.Create)
	router.Get("/account/listing/:id", hand.Listing)
	router.Put("/account/update/:id", hand.Update)
	router.Delete("/account/delete/:id", hand.Delete)

	return nil
}
