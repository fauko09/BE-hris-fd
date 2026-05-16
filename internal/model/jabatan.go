package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Jabatan struct {
	JabatanID   string `gorm:"type:char(36);primaryKey" json:"jabatan_id"`
	NamaJabatan string `gorm:"type:varchar(100);not null" json:"nama_jabatan"`
}

func (j *Jabatan) BeforeCreate(tx *gorm.DB) error {
	if j.JabatanID == "" {
		j.JabatanID = uuid.New().String()
	}
	return nil
}

type JabatanRequest struct {
	NamaJabatan string `json:"nama_jabatan" binding:"required"`
}
