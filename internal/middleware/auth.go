package middleware

import (
	"strings"

	"github.com/asvinicius/actnsgo/internal/service/token"
	"github.com/gofiber/fiber/v3"
)

func RequireAuth(ts *token.TokenService) fiber.Handler {
	return func(c fiber.Ctx) error {
		authorization := c.Get("Authorization")

		if authorization == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "token não informado",
			})
		}

		const bearerPrefix = "Bearer "

		if !strings.HasPrefix(authorization, bearerPrefix) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "token inválido",
			})
		}

		tokenString := strings.TrimPrefix(authorization, bearerPrefix)

		claims, err := ts.ValidateToken(tokenString)

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "token inválido",
			})
		}

		c.Locals("claims", claims)

		return c.Next()
	}
}
