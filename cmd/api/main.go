package main

import (
	"log"

	"github.com/asvinicius/actnsgo/internal/config"
	"github.com/asvinicius/actnsgo/internal/db"
	"github.com/asvinicius/actnsgo/internal/routes"
	"github.com/gofiber/fiber/v3"
)

func main() {
	cfg, err := config.Load()

	if err != nil {
		log.Fatal(err)
	}

	pool, err := db.Connect(cfg)

	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	routes.RegisterRoutes(app, pool)

	log.Fatal(app.Listen(":" + cfg.App.Port))
}
