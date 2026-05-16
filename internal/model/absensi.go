package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Absensi struct {
	AbsensiID    string     `gorm:"type:char(36);primaryKey" json:"absensi_id"`
	UID          string     `gorm:"type:char(36);not null" json:"uid"`
	Tanggal      time.Time  `gorm:"type:date;not null" json:"tanggal"`
	JamMasuk     *time.Time `json:"jam_masuk"`
	JamKeluar    *time.Time `json:"jam_keluar"`
	Status       string     `gorm:"type:varchar(20);default:'hadir'" json:"status"`
	Keterangan   *string    `gorm:"type:text" json:"keterangan"`
	LokasiMasuk  *string    `gorm:"type:varchar(255)" json:"lokasi_masuk"`
	LokasiKeluar *string    `gorm:"type:varchar(255)" json:"lokasi_keluar"`
	CreatedAt    time.Time  `json:"created_at"`

	User *User `gorm:"foreignKey:UID;references:UID" json:"user,omitempty"`
}

func (a *Absensi) BeforeCreate(tx *gorm.DB) error {
	if a.AbsensiID == "" {
		a.AbsensiID = uuid.New().String()
	}
	return nil
}

type ClockInRequest struct {
	Lokasi string `json:"lokasi"`
}

type ClockOutRequest struct {
	Lokasi string `json:"lokasi"`
}
