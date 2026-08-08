package backup

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/asvinicius/actnsgo/internal/model"
	"github.com/asvinicius/actnsgo/internal/repository"
)

type BackupService struct {
	backupRepository *repository.BackupRepository
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func NewBackupService(backupRepository *repository.BackupRepository) *BackupService {
	return &BackupService{
		backupRepository: backupRepository,
	}
}

func (bs *BackupService) RunBackup(trigger string, dbConfig DBConfig, backupDir string) (*model.Backup, error) {
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	fileName := fmt.Sprintf("backup_%s.sql", timestamp)
	filePath := fmt.Sprintf("%s/%s", backupDir, fileName)

	status := "success"
	var errorMessage *string

	cmd := exec.Command(
		"pg_dump",
		"-h", dbConfig.Host,
		"-p", dbConfig.Port,
		"-U", dbConfig.User,
		"-d", dbConfig.Database,
		"-f", filePath,
	)

	cmd.Env = append(os.Environ(), "PGPASSWORD="+dbConfig.Password)

	if output, err := cmd.CombinedOutput(); err != nil {
		msg := string(output)
		errorMessage = &msg
		status = "failed"
	}

	var sizeBytes int64
	if info, err := os.Stat(filePath); err == nil {
		sizeBytes = info.Size()
	}

	backup := model.Backup{
		BackupFile:         fileName,
		BackupPath:         filePath,
		BackupSize:         sizeBytes,
		BackupTrigger:      trigger,
		BackupStatus:       status,
		BackupErrorMessage: errorMessage,
		BackupCreatedAt:    time.Now(),
	}

	backupID, err := bs.backupRepository.Insert(backup)
	if err != nil {
		return nil, err
	}

	backup.BackupID = backupID

	return &backup, nil
}

func (bs *BackupService) Listing() ([]model.Backup, error) {
	return bs.backupRepository.Listing()
}
