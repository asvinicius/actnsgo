package repository

import (
	"context"
	"errors"

	"github.com/asvinicius/actnsgo/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrMktStatusNotFound = errors.New("market status not found")

type MarketStatusRepository struct {
	pool *pgxpool.Pool
}

func NewMarketStatusRepository(pool *pgxpool.Pool) *MarketStatusRepository {
	return &MarketStatusRepository{
		pool: pool,
	}
}

func (r *MarketStatusRepository) GetStatus() (*model.MarketStatus, error) {
	var ms model.MarketStatus

	query := `
		SELECT
			ms_id,
			ms_status,
			ms_current_round,
			ms_current_month,
			ms_current_season,
			ms_shutdown_day,
			ms_shutdown_month,
			ms_shutdown_year,
			ms_shutdown_hour,
			ms_shutdown_minute,
			ms_shutdown_timestamp,
			ms_post_round,
			ms_newmonth,
			ms_running
		FROM market_status
	`
	row := r.pool.QueryRow(context.Background(), query)

	err := row.Scan(
		&ms.MsID,
		&ms.MsStatus,
		&ms.MsCurrentRound,
		&ms.MsCurrentMonth,
		&ms.MsCurrentSeason,
		&ms.MsShutdownDay,
		&ms.MsShutdownMonth,
		&ms.MsShutdownYear,
		&ms.MsShutdownHour,
		&ms.MsShutdownMinute,
		&ms.MsShutdownTimestamp,
		&ms.MsPostRound,
		&ms.MsNewMonth,
		&ms.MsRunning,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrMktStatusNotFound
	}

	if err != nil {
		return nil, err
	}

	return &ms, nil
}

func (r *MarketStatusRepository) UpdateStatus(ms model.MarketStatus) error {
	query := `
		UPDATE market_status
		SET
			ms_status              = $1,
			ms_current_round       = $2,
			ms_current_month       = $3,
			ms_current_season      = $4,
			ms_shutdown_day        = $5,
			ms_shutdown_month      = $6,
			ms_shutdown_year       = $7,
			ms_shutdown_hour       = $8,
			ms_shutdown_minute     = $9,
			ms_shutdown_timestamp  = $10,
			ms_post_round          = $11,
			ms_newmonth            = $12,
			ms_running             = $13
		WHERE ms_id = $14
	`

	cmdTag, err := r.pool.Exec(context.Background(), query,
		ms.MsStatus,
		ms.MsCurrentRound,
		ms.MsCurrentMonth,
		ms.MsCurrentSeason,
		ms.MsShutdownDay,
		ms.MsShutdownMonth,
		ms.MsShutdownYear,
		ms.MsShutdownHour,
		ms.MsShutdownMinute,
		ms.MsShutdownTimestamp,
		ms.MsPostRound,
		ms.MsNewMonth,
		ms.MsRunning,
		ms.MsID,
	)

	if err != nil {
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return ErrMktStatusNotFound
	}

	return nil
}
