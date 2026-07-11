package main

import (
	"log"

	"github.com/asvinicius/actnsgo/internal/config"
	"github.com/asvinicius/actnsgo/internal/routes"
	"github.com/gofiber/fiber/v3"
)

func main() {
	cfg, err := config.Load()

	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	routes.Setup(app)

	log.Fatal(app.Listen(":" + cfg.AppPort))
}
