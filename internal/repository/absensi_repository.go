package repository

import (
"fmt"
"time"

"hris-api/config"
"hris-api/internal/model"

"github.com/google/uuid"
)

func ClockIn(uid string, lokasi string) (*model.Absensi, error) {
now := time.Now()
today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

var existing model.Absensi
err := config.DB.Where("uid = ? AND tanggal = ?", uid, today).First(&existing).Error
if err == nil {
return nil, fmt.Errorf("sudah clock in hari ini")
}

absensi := &model.Absensi{
AbsensiID:   uuid.New().String(),
UID:         uid,
Tanggal:     today,
JamMasuk:    &now,
Status:      "hadir",
LokasiMasuk: &lokasi,
}
return absensi, config.DB.Create(absensi).Error
}

func ClockOut(uid string, lokasi string) (*model.Absensi, error) {
now := time.Now()
today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

var absensi model.Absensi
err := config.DB.Where("uid = ? AND tanggal = ?", uid, today).First(&absensi).Error
if err != nil {
return nil, fmt.Errorf("belum clock in hari ini")
}
if absensi.JamKeluar != nil {
return nil, fmt.Errorf("sudah clock out hari ini")
}

absensi.JamKeluar = &now
absensi.LokasiKeluar = &lokasi
return &absensi, config.DB.Save(&absensi).Error
}

func GetAbsensiByUID(uid string, bulan, tahun int) ([]model.Absensi, error) {
var list []model.Absensi
query := config.DB.Where("uid = ?", uid)
if bulan > 0 && tahun > 0 {
query = query.Where("EXTRACT(MONTH FROM tanggal) = ? AND EXTRACT(YEAR FROM tanggal) = ?", bulan, tahun)
}
err := query.Order("tanggal DESC").Find(&list).Error
return list, err
}

func GetAllAbsensiToday() ([]model.Absensi, error) {
var list []model.Absensi
today := time.Now().Format("2006-01-02")
err := config.DB.Preload("User").Where("tanggal = ?", today).Find(&list).Error
return list, err
}
