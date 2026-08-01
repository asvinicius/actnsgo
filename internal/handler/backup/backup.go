package backup

import (
	"github.com/asvinicius/actnsgo/internal/dto"
	backupservice "github.com/asvinicius/actnsgo/internal/service/backup"
	"github.com/gofiber/fiber/v3"
)

type BackupHandler struct {
	backupService *backupservice.BackupService
}

func NewBackupHandler(backupService *backupservice.BackupService) *BackupHandler {
	return &BackupHandler{
		backupService: backupService,
	}
}
func (bh *BackupHandler) Listing(c fiber.Ctx) error {

	listing, err := bh.backupService.Listing()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "erro ao listar backups",
		})
	}

	response := dto.ToBackupResponseList(listing)

	return c.Status(fiber.StatusOK).JSON(response)
}
