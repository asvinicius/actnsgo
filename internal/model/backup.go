package model

import "time"

type Backup struct {
	BackupID           int64
	BackupFile         string
	BackupPath         string
	BackupSize         int64
	BackupTrigger      string
	BackupStatus       string
	BackupErrorMessage *string
	BackupCreatedAt    time.Time
	BackupRestoredAt   *time.Time
}
