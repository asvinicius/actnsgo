package super

import "github.com/gofiber/fiber/v2"

func SetRoutes(r fiber.Router) {
	super := r.Group("/super")

	super.Post("/", addSuper)
}
