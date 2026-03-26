package model

import "github.com/google/uuid"

type Divisi struct {
	DivisiID   uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"divisi_id"`
	NameDivisi string    `gorm:"type:varchar(100);not null" json:"name_divisi"`
}

type DivisiRequest struct {
	NameDivisi string `json:"name_divisi" binding:"required"`
}
