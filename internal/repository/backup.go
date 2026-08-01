package repository

import (
	"context"

	"github.com/asvinicius/actnsgo/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BackupRepository struct {
	pool *pgxpool.Pool
}

func NewBackupRepository(pool *pgxpool.Pool) *BackupRepository {
	return &BackupRepository{
		pool: pool,
	}
}

func (r *BackupRepository) Listing() ([]model.Backup, error) {

	query := `
		SELECT
			backup_id,
			backup_file,
			backup_path,
			backup_size,
			backup_trigger,
			backup_status,
			backup_error_message,
			backup_created_at,
			backup_restored_at
		FROM backup
		ORDER BY backup_created_at DESC
	`

	rows, err := r.pool.Query(context.Background(), query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	listing := []model.Backup{}

	for rows.Next() {

		var backup model.Backup

		err := rows.Scan(
			&backup.BackupID,
			&backup.BackupFile,
			&backup.BackupPath,
			&backup.BackupSize,
			&backup.BackupTrigger,
			&backup.BackupStatus,
			&backup.BackupErrorMessage,
			&backup.BackupCreatedAt,
			&backup.BackupRestoredAt,
		)

		if err != nil {
			return nil, err
		}

		listing = append(listing, backup)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return listing, nil

}
