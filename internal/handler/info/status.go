package info

import (
	"github.com/asvinicius/actnsgo/internal/dto"
	"github.com/asvinicius/actnsgo/internal/service/info"
	"github.com/gofiber/fiber/v3"
)

type StatusHandler struct {
	statusService *info.StatusService
}

func NewStatusHandler(statusService *info.StatusService) *StatusHandler {
	return &StatusHandler{
		statusService: statusService,
	}
}

func (sh *StatusHandler) GetStatus(c fiber.Ctx) error {
	status, err := sh.statusService.GetStatus()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "erro ao buscar api/status",
		})
	}

	response := dto.ToHeaderInfoResponse(*status)

	return c.Status(fiber.StatusOK).JSON(response)
}
