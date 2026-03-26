package model

import (
	"time"

	"github.com/google/uuid"
)

type Absensi struct {
	AbsensiID    uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"absensi_id"`
	UID          uuid.UUID  `gorm:"type:uuid;not null" json:"uid"`
	Tanggal      time.Time  `gorm:"type:date;not null" json:"tanggal"`
	JamMasuk     *time.Time `json:"jam_masuk"`
	JamKeluar    *time.Time `json:"jam_keluar"`
	Status       string     `gorm:"type:varchar(20);default:'hadir'" json:"status"`
	Keterangan   *string    `gorm:"type:text" json:"keterangan"`
	LokasiMasuk  *string    `gorm:"type:varchar(255)" json:"lokasi_masuk"`
	LokasiKeluar *string    `gorm:"type:varchar(255)" json:"lokasi_keluar"`
	CreatedAt    time.Time  `json:"created_at"`

	User *User `gorm:"foreignKey:UID" json:"user,omitempty"`
}

type ClockInRequest struct {
	Lokasi string `json:"lokasi"`
}

type ClockOutRequest struct {
	Lokasi string `json:"lokasi"`
}
