package repository

import (
	"context"
	"errors"

	"github.com/asvinicius/actnsgo/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrAdmNotFound = errors.New("adm not found")

type AdmRepository struct {
	pool *pgxpool.Pool
}

func NewAdmRepository(pool *pgxpool.Pool) *AdmRepository {
	return &AdmRepository{
		pool: pool,
	}
}

func (r *AdmRepository) Insert(a model.UserAdm) (int64, error) {
	query := `
		INSERT INTO user_adm (
			adm_name,
			adm_login,
			adm_password,
			adm_status,
			adm_created_at
		) VALUES ($1, $2, $3, $4, $5)
		RETURNING adm_id
	`
	var admID int64

	err := r.pool.QueryRow(context.Background(), query,
		a.AdmName,
		a.AdmLogin,
		a.AdmPassword,
		a.AdmStatus,
		a.AdmCreatedAt,
	).Scan(&admID)

	if err != nil {
		return 0, err
	}

	return admID, nil

}

func (r *AdmRepository) Update(a model.UserAdm) error {
	query := `
		UPDATE user_adm
		SET
			adm_name = $1,
			adm_login = $2,
			adm_status = $3,
			adm_updated_at = $4
		WHERE adm_id = $5
	`

	cmdTag, err := r.pool.Exec(context.Background(), query,
		a.AdmName,
		a.AdmLogin,
		a.AdmStatus,
		a.AdmUpdatedAt,
		a.AdmID,
	)

	if err != nil {
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return ErrAdmNotFound
	}

	return nil
}

func (r *AdmRepository) Delete(admID int64) error {
	query := `
		DELETE FROM user_adm
		WHERE adm_id = $1
	`

	cmdTag, err := r.pool.Exec(context.Background(), query, admID)

	if err != nil {
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return ErrAdmNotFound
	}

	return nil
}

func (r *AdmRepository) FindByID(admID int64) (*model.UserAdm, error) {
	var adm model.UserAdm

	query := `
		SELECT
			adm_id,
			adm_name,
			adm_login,
			adm_status,
			adm_updated_at
		FROM user_adm
		WHERE adm_id = $1
	`

	row := r.pool.QueryRow(context.Background(), query, admID)

	err := row.Scan(
		&adm.AdmID,
		&adm.AdmName,
		&adm.AdmLogin,
		&adm.AdmStatus,
		&adm.AdmUpdatedAt,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrAdmNotFound
	}

	if err != nil {
		return nil, err
	}

	return &adm, nil

}

func (r *AdmRepository) Listing() ([]model.UserAdm, error) {
	query := `
		SELECT
			adm_id,
			adm_name,
			adm_status,
			adm_created_at,
			adm_updated_at,
			adm_last_login
		FROM user_adm
		ORDER BY adm_name ASC
	`

	rows, err := r.pool.Query(context.Background(), query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	listing := []model.UserAdm{}

	for rows.Next() {
		var adm model.UserAdm

		err := rows.Scan(
			&adm.AdmID,
			&adm.AdmName,
			&adm.AdmStatus,
			&adm.AdmCreatedAt,
			&adm.AdmUpdatedAt,
			&adm.AdmLastLogin,
		)

		if err != nil {
			return nil, err
		}

		listing = append(listing, adm)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return listing, nil
}
