package super

import "github.com/gofiber/fiber/v3"

func SetRoutes(r fiber.Router) {
	super := r.Group("/supers")

	super.Post("/", addSuper())
}
