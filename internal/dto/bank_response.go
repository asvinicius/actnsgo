package dto

import "github.com/asvinicius/actnsgo/internal/model"

type BankResponse struct {
	BankID     int64   `json:"bank_id"`
	BankName   string  `json:"bank_name"`
	BankLogo   *string `json:"bank_logo,omitempty"`
	BankStatus bool    `json:"bank_status"`
}

func ToBankResponse(b model.Bank) BankResponse {
	return BankResponse{
		BankID:     b.BankID,
		BankName:   b.BankName,
		BankLogo:   b.BankLogo,
		BankStatus: b.BankStatus,
	}
}

func ToBankResponseList(banks []model.Bank) []BankResponse {
	result := make([]BankResponse, 0, len(banks))
	for _, b := range banks {
		result = append(result, ToBankResponse(b))
	}
	return result
}
