package model

import "time"

type UserAdm struct {
	AdmID        int64
	AdmName      string
	AdmLogin     string
	AdmPassword  string
	AdmStatus    bool
	AdmCreatedAt time.Time
	AdmUpdatedAt *time.Time
	AdmLastLogin *time.Time
}
