package routes

import (
	"github.com/asvinicius/actnsgo/internal/config"
	admhandler "github.com/asvinicius/actnsgo/internal/handler/adm"
	"github.com/asvinicius/actnsgo/internal/repository"
	admservice "github.com/asvinicius/actnsgo/internal/service/adm"
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5/pgxpool"
)

func AdmRoutes(router fiber.Router, pool *pgxpool.Pool, cfg config.Config) error {
	rep := repository.NewAdmRepository(pool)
	serv := admservice.NewAdmService(rep)
	hand := admhandler.NewAdmHandler(serv)

	router.Post("/adm/create", hand.Create)
	router.Get("/adm/listing", hand.Listing)
	router.Put("/adm/update/:id", hand.Update)
	router.Delete("/adm/delete/:id", hand.Delete)

	return nil
}
