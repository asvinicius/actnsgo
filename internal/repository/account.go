package repository

import (
	"context"
	"errors"

	"github.com/asvinicius/actnsgo/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrAccountNotFound = errors.New("account not found")

type AccountRepository struct {
	pool *pgxpool.Pool
}

func NewAccountRepository(pool *pgxpool.Pool) *AccountRepository {
	return &AccountRepository{
		pool: pool,
	}
}

func (r *AccountRepository) Insert(a model.Account) (int64, error) {
	query := `
		INSERT INTO account (
			account_adm,
			account_bank,
			account_key,
			account_status
		) VALUES ($1, $2, $3, $4)
		RETURNING account_id
	`

	var accountID int64

	err := r.pool.QueryRow(context.Background(), query,
		a.AccountAdm,
		a.AccountBank,
		a.AccountKey,
		a.AccountStatus,
	).Scan(&accountID)

	if err != nil {
		return 0, err
	}

	return accountID, nil
}

func (r *AccountRepository) Update(a model.Account) error {
	query := `
		UPDATE account
		SET
			account_bank = $1,
			account_key = $2,
			account_status = $3
		WHERE account_id = $4
	`

	cmdTag, err := r.pool.Exec(context.Background(), query,
		a.AccountBank,
		a.AccountKey,
		a.AccountStatus,
		a.AccountID,
	)

	if err != nil {
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return ErrAccountNotFound
	}

	return nil
}

func (r *AccountRepository) Delete(accountID int64) error {
	query := `
		DELETE FROM account
		WHERE account_id = $1
	`

	cmdTag, err := r.pool.Exec(context.Background(), query, accountID)

	if err != nil {
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return ErrAccountNotFound
	}

	return nil
}

func (r *AccountRepository) FindByID(accountID int64) (*model.Account, error) {
	var account model.Account

	query := `
		SELECT
			account_id,
			account_bank,
			account_key,
			account_status
		FROM account
		WHERE account_id = $1
	`

	row := r.pool.QueryRow(context.Background(), query, accountID)

	err := row.Scan(
		&account.AccountID,
		&account.AccountBank,
		&account.AccountKey,
		&account.AccountStatus,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrAccountNotFound
	}

	if err != nil {
		return nil, err
	}

	return &account, nil

}

func (r *AccountRepository) Listing(admID int64) ([]model.AccountWithBank, error) {
	query := `
		SELECT
			account.account_id,
			account.account_bank,
			account.account_key,
			account.account_status,
			bank.bank_name,
			bank.bank_logo
		FROM account
		JOIN bank ON bank.bank_id = account_bank
		WHERE account.account_adm = $1
		ORDER BY bank.bank_name ASC
	`

	rows, err := r.pool.Query(context.Background(), query, admID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	listing := []model.AccountWithBank{}

	for rows.Next() {
		var account model.AccountWithBank

		err := rows.Scan(
			&account.AccountID,
			&account.AccountBank,
			&account.AccountKey,
			&account.AccountStatus,
			&account.BankName,
			&account.BankLogo,
		)

		if err != nil {
			return nil, err
		}

		listing = append(listing, account)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return listing, nil

}
