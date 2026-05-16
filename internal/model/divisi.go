package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Divisi struct {
	DivisiID   string `gorm:"type:char(36);primaryKey" json:"divisi_id"`
	NameDivisi string `gorm:"type:varchar(100);not null" json:"name_divisi"`
}

func (d *Divisi) BeforeCreate(tx *gorm.DB) error {
	if d.DivisiID == "" {
		d.DivisiID = uuid.New().String()
	}
	return nil
}

type DivisiRequest struct {
	NameDivisi string `json:"name_divisi" binding:"required"`
}
