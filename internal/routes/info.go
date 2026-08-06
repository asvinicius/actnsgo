package routes

import (
	"github.com/asvinicius/actnsgo/internal/client"
	infohandler "github.com/asvinicius/actnsgo/internal/handler/info"
	infoservice "github.com/asvinicius/actnsgo/internal/service/info"
	"github.com/gofiber/fiber/v3"
)

func InfoRoutes(router fiber.Router, statusClient *client.Client) error {
	serv := infoservice.NewStatusService(statusClient)
	hand := infohandler.NewStatusHandler(serv)

	router.Get("/info", hand.GetStatus)

	return nil
}
