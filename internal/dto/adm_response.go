package dto

import (
	"time"

	"github.com/asvinicius/actnsgo/internal/model"
)

type AdmResponse struct {
	AdmID        int64      `json:"adm_id"`
	AdmName      string     `json:"adm_name"`
	AdmLogin     string     `json:"adm_login"`
	AdmStatus    bool       `json:"adm_status"`
	AdmCreatedAt time.Time  `json:"adm_created_at"`
	AdmUpdatedAt *time.Time `json:"adm_updated_at,omitempty"`
	AdmLastLogin *time.Time `json:"adm_last_login,omitempty"`
}

func ToAdmResponse(a model.UserAdm) AdmResponse {
	return AdmResponse{
		AdmID:        a.AdmID,
		AdmName:      a.AdmName,
		AdmLogin:     a.AdmLogin,
		AdmStatus:    a.AdmStatus,
		AdmCreatedAt: a.AdmCreatedAt,
		AdmUpdatedAt: a.AdmUpdatedAt,
		AdmLastLogin: a.AdmLastLogin,
	}
}

func ToAdmResponseList(adms []model.UserAdm) []AdmResponse {
	result := make([]AdmResponse, 0, len(adms))
	for _, a := range adms {
		result = append(result, ToAdmResponse(a))
	}

	return result
}
