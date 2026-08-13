package repository

import (
	"context"

	"github.com/asvinicius/actnsgo/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BankRepository struct {
	pool *pgxpool.Pool
}

func NewBankRepository(pool *pgxpool.Pool) *BankRepository {
	return &BankRepository{
		pool: pool,
	}
}

func (r *BankRepository) Insert(b model.Bank) (int64, error) {
	query := `
		INSERT INTO bank (
    		bank_name,
    		bank_logo,
    		bank_status
		) VALUES ($1, $2, $3)
		RETURNING bank_id
	`
	var bankID int64

	err := r.pool.QueryRow(context.Background(), query,
		b.BankName,
		b.BankLogo,
		b.BankStatus,
	).Scan(&bankID)

	if err != nil {
		return 0, err
	}

	return bankID, nil
}

func (r *BankRepository) Listing() ([]model.Bank, error) {
	query := `
		SELECT
			bank_id,
			bank_name,
			bank_logo,
			bank_status
		FROM bank
		ORDER BY bank_name ASC
	`

	rows, err := r.pool.Query(context.Background(), query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	listing := []model.Bank{}

	for rows.Next() {
		var bank model.Bank

		err := rows.Scan(
			&bank.BankID,
			&bank.BankName,
			&bank.BankLogo,
			&bank.BankStatus,
		)

		if err != nil {
			return nil, err
		}

		listing = append(listing, bank)

	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return listing, nil
}
