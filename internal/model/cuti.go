package model

import (
"time"

"github.com/google/uuid"
)

type Cuti struct {
CutiID       uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"cuti_id"`
UID          uuid.UUID  `gorm:"type:uuid;not null" json:"uid"`
TipeCuti     string     `gorm:"type:varchar(20)" json:"tipe_cuti"`
TanggalMulai time.Time  `gorm:"type:date" json:"tanggal_mulai"`
TanggalAkhir time.Time  `gorm:"type:date" json:"tanggal_akhir"`
TotalHari    int        `json:"total_hari"`
Alasan       string     `gorm:"type:text" json:"alasan"`
Status       string     `gorm:"type:varchar(20);default:pending" json:"status"`
CatatanHR    *string    `gorm:"type:text" json:"catatan_hr"`
ApprovedBy   *uuid.UUID `gorm:"type:uuid" json:"approved_by"`
ApprovedAt   *time.Time `json:"approved_at"`
CreatedAt    time.Time  `json:"created_at"`

User *User `gorm:"foreignKey:UID;references:UID" json:"user,omitempty"`
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
