package model

type Account struct {
	AccountID     int64
	AccountAdm    int64
	AccountBank   int64
	AccountKey    string
	AccountStatus bool
}

type AccountWithBank struct {
	AccountID     int64
	AccountBank   int64
	AccountKey    string
	AccountStatus bool
	BankName      string
	BankLogo      *string
}
