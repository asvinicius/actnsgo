package backup

import (
	"github.com/asvinicius/actnsgo/internal/model"
	"github.com/asvinicius/actnsgo/internal/repository"
)

type BackupService struct {
	backupRepository *repository.BackupRepository
}

func NewBackupService(backupRepository *repository.BackupRepository) *BackupService {
	return &BackupService{
		backupRepository: backupRepository,
	}
}

func (bs *BackupService) Listing() ([]model.Backup, error) {
	return bs.backupRepository.Listing()
}
