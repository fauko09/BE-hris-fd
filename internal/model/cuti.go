package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Cuti struct {
	CutiID       string     `gorm:"type:char(36);primaryKey" json:"cuti_id"`
	UID          string     `gorm:"type:char(36);not null" json:"uid"`
	TipeCuti     string     `gorm:"type:varchar(20)" json:"tipe_cuti"`
	TanggalMulai time.Time  `gorm:"type:date" json:"tanggal_mulai"`
	TanggalAkhir time.Time  `gorm:"type:date" json:"tanggal_akhir"`
	TotalHari    int        `json:"total_hari"`
	Alasan       string     `gorm:"type:text" json:"alasan"`
	Status       string     `gorm:"type:varchar(20);default:'pending'" json:"status"`
	CatatanHR    *string    `gorm:"type:text" json:"catatan_hr"`
	ApprovedBy   *string    `gorm:"type:char(36)" json:"approved_by"`
	ApprovedAt   *time.Time `json:"approved_at"`
	CreatedAt    time.Time  `json:"created_at"`

	User *User `gorm:"foreignKey:UID;references:UID" json:"user,omitempty"`
}

func (c *Cuti) BeforeCreate(tx *gorm.DB) error {
	if c.CutiID == "" {
		c.CutiID = uuid.New().String()
	}
	return nil
}

type CutiRequest struct {
	TipeCuti     string `json:"tipe_cuti" binding:"required"`
	TanggalMulai string `json:"tanggal_mulai" binding:"required"`
	TanggalAkhir string `json:"tanggal_akhir" binding:"required"`
	Alasan       string `json:"alasan" binding:"required"`
}

type ApproveCutiRequest struct {
	Status    string  `json:"status" binding:"required"`
	CatatanHR *string `json:"catatan_hr"`
}
