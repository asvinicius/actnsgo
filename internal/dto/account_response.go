package dto

type AccountResponse struct {
	AccountID     int64   `json:"account_id"`
	AccountBank   int64   `json:"account_bank"`
	AccountKey    string  `json:"account_key"`
	AccountStatus bool    `json:"account_status"`
	BankName      string  `json:"bank_name"`
	BankLogo      *string `json:"bank_logo,omitempty"`
}
