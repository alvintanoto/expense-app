package entity

import "time"

type GeneralData struct {
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy *string   `json:"updated_by"`
	IsDeleted bool      `json:"is_deleted"`
}
