package repository

import (
	"context"
	"errors"

	"github.com/asvinicius/actnsgo/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrSuperNotFound = errors.New("super not found")

type SuperRepository struct {
	pool *pgxpool.Pool
}

func NewSuperRepository(pool *pgxpool.Pool) *SuperRepository {
	return &SuperRepository{
		pool: pool,
	}
}

func (r *SuperRepository) FindByLogin(superLogin string) (*model.UserSuper, error) {
	var super model.UserSuper

	query := `
		SELECT
			super_id,
			super_name,
			super_login,
			super_password,
			super_status
		FROM user_super
		WHERE super_login = $1
	`

	row := r.pool.QueryRow(context.Background(), query, superLogin)

	err := row.Scan(
		&super.SuperID,
		&super.SuperName,
		&super.SuperLogin,
		&super.SuperPassword,
		&super.SuperStatus,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrSuperNotFound
	}

	if err != nil {
		return nil, err
	}

	return &super, nil
}
