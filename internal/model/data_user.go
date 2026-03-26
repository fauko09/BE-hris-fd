package model

import (
	"time"

	"github.com/google/uuid"
)

type DataUser struct {
	DaID                uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"da_id"`
	UID                 uuid.UUID  `gorm:"type:uuid;not null;unique" json:"uid"`
	DivisiID            *uuid.UUID `gorm:"type:uuid" json:"divisi_id"`
	JabatanID           *uuid.UUID `gorm:"type:uuid" json:"jabatan_id"`
	Status              string     `gorm:"type:varchar(50)" json:"status"`
	AlamatTinggal       string     `gorm:"type:text" json:"alamat_tinggal"`
	TipeTinggal         string     `gorm:"type:varchar(50)" json:"tipe_tinggal"`
	DomisiliTinggal     string     `gorm:"type:varchar(100)" json:"domisili_tinggal"`
	ProvTinggal         string     `gorm:"type:varchar(100)" json:"prov_tinggal"`
	AlamatKTP           string     `gorm:"type:text" json:"alamat_ktp"`
	DomisiliKTP         string     `gorm:"type:varchar(100)" json:"domisili_ktp"`
	ProvKTP             string     `gorm:"type:varchar(100)" json:"prov_ktp"`
	DomisiliLahir       string     `gorm:"type:varchar(10)" json:"domisili_lahir"`
	TanggalLahir        *time.Time `json:"tanggal_lahir"`
	NoNPWP              string     `gorm:"type:varchar(20)" json:"no_npwp"`
	NoBPJS              string     `gorm:"type:varchar(20)" json:"no_bpjs"`
	IsBPJSTK            bool       `gorm:"default:false" json:"is_bpjs_tk"`
	NoAsuransi          string     `gorm:"type:varchar(20)" json:"no_asuransi"`
	JenisAsuransi       string     `gorm:"type:varchar(100)" json:"jenis_asuransi"`
	NamaAsuransi        string     `gorm:"type:varchar(100)" json:"nama_asuransi"`
	StatusPerkawinan    string     `gorm:"type:varchar(20)" json:"status_perkawinan"`
	IDFamily            *uuid.UUID `gorm:"type:uuid" json:"id_family"`
	NomorTelepon        string     `gorm:"type:varchar(15)" json:"nomor_telepon"`
	NomorTeleponSecond  *string    `gorm:"type:varchar(15)" json:"nomor_telepon_second"`
	NomorTeleponDarurat *string    `gorm:"type:varchar(15)" json:"nomor_telepon_darurat"`
	TanggalMasuk        *time.Time `json:"tanggal_masuk"`
	NoRekening          string     `gorm:"type:varchar(30)" json:"no_rekening"`
	TipeBank            string     `gorm:"type:varchar(50)" json:"tipe_bank"`

	// Relasi
	User    *User    `gorm:"foreignKey:UID" json:"user,omitempty"`
	Divisi  *Divisi  `gorm:"foreignKey:DivisiID" json:"divisi,omitempty"`
	Jabatan *Jabatan `gorm:"foreignKey:JabatanID" json:"jabatan,omitempty"`
}

type DataUserRequest struct {
	DivisiID            string  `json:"divisi_id"`
	JabatanID           string  `json:"jabatan_id"`
	Status              string  `json:"status"`
	AlamatTinggal       string  `json:"alamat_tinggal"`
	TipeTinggal         string  `json:"tipe_tinggal"`
	DomisiliTinggal     string  `json:"domisili_tinggal"`
	ProvTinggal         string  `json:"prov_tinggal"`
	AlamatKTP           string  `json:"alamat_ktp"`
	DomisiliKTP         string  `json:"domisili_ktp"`
	ProvKTP             string  `json:"prov_ktp"`
	DomisiliLahir       string  `json:"domisili_lahir"`
	TanggalLahir        string  `json:"tanggal_lahir"`
	NoNPWP              string  `json:"no_npwp"`
	NoBPJS              string  `json:"no_bpjs"`
	IsBPJSTK            bool    `json:"is_bpjs_tk"`
	NoAsuransi          string  `json:"no_asuransi"`
	JenisAsuransi       string  `json:"jenis_asuransi"`
	NamaAsuransi        string  `json:"nama_asuransi"`
	StatusPerkawinan    string  `json:"status_perkawinan"`
	NomorTelepon        string  `json:"nomor_telepon"`
	NomorTeleponSecond  *string `json:"nomor_telepon_second"`
	NomorTeleponDarurat *string `json:"nomor_telepon_darurat"`
	TanggalMasuk        string  `json:"tanggal_masuk"`
	NoRekening          string  `json:"no_rekening"`
	TipeBank            string  `json:"tipe_bank"`
}
