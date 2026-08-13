package routes

import (
	"github.com/asvinicius/actnsgo/internal/config"
	financehandler "github.com/asvinicius/actnsgo/internal/handler/finance"
	"github.com/asvinicius/actnsgo/internal/repository"
	financeservice "github.com/asvinicius/actnsgo/internal/service/finance"
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5/pgxpool"
)

func BankRoutes(router fiber.Router, pool *pgxpool.Pool, cfg config.Config) error {
	rep := repository.NewBankRepository(pool)
	serv := financeservice.NewBankService(rep)
	hand := financehandler.NewBankHandler(serv, cfg.Upload.BankLogoDir)

	router.Post("/bank/create", hand.Create)
	router.Get("/bank/listing", hand.Listing)

	return nil
}
