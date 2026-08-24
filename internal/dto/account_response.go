package dto

import "github.com/asvinicius/actnsgo/internal/model"

type AccountResponse struct {
	AccountID     int64   `json:"account_id"`
	AccountBank   int64   `json:"account_bank"`
	AccountKey    string  `json:"account_key"`
	AccountStatus bool    `json:"account_status"`
	BankName      string  `json:"bank_name"`
	BankLogo      *string `json:"bank_logo,omitempty"`
}

func ToAccountResponse(a model.AccountWithBank) AccountResponse {
	return AccountResponse{
		AccountID:     a.AccountID,
		AccountBank:   a.AccountBank,
		AccountKey:    a.AccountKey,
		AccountStatus: a.AccountStatus,
		BankName:      a.BankName,
		BankLogo:      a.BankLogo,
	}
}

func ToAccountResponseList(accounts []model.AccountWithBank) []AccountResponse {
	result := make([]AccountResponse, 0, len(accounts))
	for _, a := range accounts {
		result = append(result, ToAccountResponse(a))
	}

	return result
}
