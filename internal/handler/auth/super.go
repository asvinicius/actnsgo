package auth

import (
	"errors"
	"strings"
	"github.com/asvinicius/actnsgo/internal/dto"
	"github.com/asvinicius/actnsgo/internal/service/auth"
	"github.com/gofiber/fiber/v3"
)

type AuthHandler struct {
	authService *auth.AuthService
}

func NewAuthHandler(authService *auth.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Login(c fiber.Ctx) error {
	var req dto.SuperLoginRequest

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "dados inválidos",
		})
	}

	super, token, err := h.authService.Authenticate(req.SuperLogin, req.SuperPassword)

	if errors.Is(err, auth.ErrSuperNotActive) {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "usuário inativo",
		})
	}

	if errors.Is(err, auth.ErrInvalidCredentials) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "usuário ou senha incorretos",
		})
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "erro interno de servidor",
		})
	}

	response := dto.SuperLoginResponse{
		SuperID:    super.SuperID,
		SuperName:  super.SuperName,
		SuperLogin: super.SuperLogin,
		Token:      token,
	}

	return c.Status(fiber.StatusOK).JSON(response)

}

func (h *AuthHandler) IsLogged(c fiber.Ctx) error {
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

    _, err := h.authService.IsValidSession(tokenString)

    if err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(dto.SuperTokenValidation{
            Valid: false,
        })
    }

    return c.JSON(dto.SuperTokenValidation{
        Valid: true,
    })
}