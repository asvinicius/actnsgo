package main

import (
	"github.com/asvinicius/actnsgo/super"
	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	api := app.Group("api")
	super.SetRoutes(api)

	app.Listen(":9001")
}
