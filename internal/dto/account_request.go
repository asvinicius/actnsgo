package dto

type AccountRequest struct {
	AccountBank   int64  `json:"account_bank"`
	AccountKey    string `json:"account_key"`
	AccountStatus bool   `json:"account_status"`
}

type AccountCreateRequest struct {
	AccountAdm  int64  `json:"account_adm"`
	AccountBank int64  `json:"account_bank"`
	AccountKey  string `json:"account_key"`
}

type AccountUpdateRequest struct {
	AccountBank   int64  `json:"account_bank"`
	AccountKey    string `json:"account_key"`
	AccountStatus bool   `json:"account_status"`
}
