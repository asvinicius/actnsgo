package dto

type AccountRequest struct {
	AccountBank   int64  `json:"account_bank"`
	AccountKey    string `json:"account_key"`
	AccountStatus bool   `json:"account_status"`
}
