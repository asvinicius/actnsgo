package dto

type AdmRequest struct {
	AdmName     string `json:"adm_name"`
	AdmLogin    string `json:"adm_login"`
	AdmPassword string `json:"adm_password"`
}

type AdmUpdateRequest struct {
	AdmName   string `json:"adm_name"`
	AdmLogin  string `json:"adm_login"`
	AdmStatus bool   `json:"adm_status"`
}

type AdmChangePassRequest struct {
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"`
}
