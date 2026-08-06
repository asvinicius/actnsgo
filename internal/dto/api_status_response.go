package dto

import (
	"time"

	"github.com/asvinicius/actnsgo/internal/response"
)

type ApiStatusResponse struct {
	MarketStatus      int       `json:"market_status"`
	CurrentRound      int       `json:"current_round"`
	ShutdownTimestamp time.Time `json:"shutdown_timestamp"`
}

func ToHeaderInfoResponse(status response.StatusResponse) ApiStatusResponse {
	return ApiStatusResponse{
		MarketStatus:      status.MarketStatus,
		CurrentRound:      status.CurrentRound,
		ShutdownTimestamp: time.Unix(status.Shutdown.Timestamp, 0),
	}
}
