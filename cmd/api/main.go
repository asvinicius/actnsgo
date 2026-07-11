package main

import (
	"log"

	"github.com/asvinicius/actnsgo/internal/routes"
	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	routes.Setup(app)

	log.Fatal(app.Listen(":3000"))
}
