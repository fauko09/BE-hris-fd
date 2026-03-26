package model

import "github.com/google/uuid"

type Jabatan struct {
JabatanID   uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"jabatan_id"`
NamaJabatan string    `gorm:"type:varchar(100);not null" json:"nama_jabatan"`
}

type JabatanRequest struct {
NamaJabatan string `json:"nama_jabatan" binding:"required"`
}
