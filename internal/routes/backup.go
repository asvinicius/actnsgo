package routes

import (
	"github.com/asvinicius/actnsgo/internal/config"
	backuphandler "github.com/asvinicius/actnsgo/internal/handler/backup"
	"github.com/asvinicius/actnsgo/internal/repository"
	backupservice "github.com/asvinicius/actnsgo/internal/service/backup"
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5/pgxpool"
)

func BackupRoutes(router fiber.Router, pool *pgxpool.Pool, cfg config.Config) error {
	rep := repository.NewBackupRepository(pool)
	serv := backupservice.NewBackupService(rep)
	hand := backuphandler.NewBackupHandler(serv)

	router.Get("/backup/listing", hand.Listing)

	return nil
}
