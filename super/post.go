package super

import (
	"net/http"

	"github.com/asvinicius/actnsgo/db"
	"github.com/gofiber/fiber/v3"
)

func addSuper(c fiber.Ctx) error {
	body := new(Super)

	if err := c.Bind().Body(body); err != nil {
		return c.Status(http.StatusBadRequest).JSON("invalid json")
	}

	id, err := db.Insert("super", body)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	body.Super_id = id

	return c.Status(http.StatusCreated).JSON(body)
}
