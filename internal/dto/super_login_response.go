package dto

type SuperLoginResponse struct {
	SuperID    int64  `json:"super_id"`
	SuperName  string `json:"super_name"`
	SuperLogin string `json:"super_login"`
	Token      string `json:"token"`
}
