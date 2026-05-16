package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Penggajian struct {
	GajiID       string     `gorm:"type:char(36);primaryKey" json:"gaji_id"`
	UID          string     `gorm:"type:char(36);not null" json:"uid"`
	Periode      string     `gorm:"type:varchar(7)" json:"periode"`
	GajiPokok    float64    `gorm:"type:decimal(15,2)" json:"gaji_pokok"`
	Tunjangan    float64    `gorm:"type:decimal(15,2);default:0" json:"tunjangan"`
	Potongan     float64    `gorm:"type:decimal(15,2);default:0" json:"potongan"`
	PotonganBPJS float64    `gorm:"type:decimal(15,2);default:0" json:"potongan_bpjs"`
	TotalGaji    float64    `gorm:"type:decimal(15,2)" json:"total_gaji"`
	StatusBayar  string     `gorm:"type:varchar(20);default:'unpaid'" json:"status_bayar"`
	TanggalBayar *time.Time `json:"tanggal_bayar"`
	Keterangan   *string    `gorm:"type:text" json:"keterangan"`
	CreatedAt    time.Time  `json:"created_at"`

	User *User `gorm:"foreignKey:UID;references:UID" json:"user,omitempty"`
}

func (p *Penggajian) BeforeCreate(tx *gorm.DB) error {
	if p.GajiID == "" {
		p.GajiID = uuid.New().String()
	}
	return nil
}

type PenggajianRequest struct {
	UID        string  `json:"uid" binding:"required"`
	Periode    string  `json:"periode" binding:"required"`
	GajiPokok  float64 `json:"gaji_pokok" binding:"required"`
	Tunjangan  float64 `json:"tunjangan"`
	Potongan   float64 `json:"potongan"`
	Keterangan *string `json:"keterangan"`
}
