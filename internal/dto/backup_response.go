package dto

import (
	"time"

	"github.com/asvinicius/actnsgo/internal/model"
)

type BackupResponse struct {
	BackupID         int64      `json:"backup_id"`
	BackupFile       string     `json:"backup_file"`
	BackupSize       int64      `json:"backup_size"`
	BackupTrigger    string     `json:"backup_trigger"`
	BackupStatus     string     `json:"backup_status"`
	BackupCreatedAt  time.Time  `json:"backup_created_at"`
	BackupRestoredAt *time.Time `json:"backup_restored_at,omitempty"`
}

func ToBackupResponse(b model.Backup) BackupResponse {
	return BackupResponse{
		BackupID:         b.BackupID,
		BackupFile:       b.BackupFile,
		BackupSize:       b.BackupSize,
		BackupTrigger:    b.BackupTrigger,
		BackupStatus:     b.BackupStatus,
		BackupCreatedAt:  b.BackupCreatedAt,
		BackupRestoredAt: b.BackupRestoredAt,
	}
}

func ToBackupResponseList(backups []model.Backup) []BackupResponse {
	result := make([]BackupResponse, 0, len(backups))
	for _, b := range backups {
		result = append(result, ToBackupResponse(b))
	}
	return result
}
