package auth

import (
	"github.com/asvinicius/actnsgo/internal/service/auth"
	"github.com/gofiber/fiber/v3"
)

type AuthHandler struct {
	authService *auth.AuthService
}

type SuperLoginRequest struct {
	SuperLogin    string `json:"super_login"`
	SuperPassword string `json:"super_password"`
}

func NewAuthHandler(authService *auth.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {

}
