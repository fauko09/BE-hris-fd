package repository

import (
"time"

"hris-api/config"
"hris-api/internal/model"

"github.com/google/uuid"
)

func CreatePenggajian(p *model.Penggajian) error {
p.GajiID = uuid.New()
p.PotonganBPJS = p.GajiPokok * 0.01
p.TotalGaji = p.GajiPokok + p.Tunjangan - p.Potongan - p.PotonganBPJS
return config.DB.Create(p).Error
}

func GetPenggajianByUID(uid uuid.UUID) ([]model.Penggajian, error) {
var list []model.Penggajian
err := config.DB.Where("uid = ?", uid).Order("periode DESC").Find(&list).Error
return list, err
}

func GetAllPenggajian(periode string) ([]model.Penggajian, error) {
var list []model.Penggajian
query := config.DB.Preload("User")
if periode != "" {
query = query.Where("periode = ?", periode)
}
err := query.Order("periode DESC").Find(&list).Error
return list, err
}

func BayarGaji(gajiID uuid.UUID) error {
now := time.Now()
return config.DB.Model(&model.Penggajian{}).
Where("gaji_id = ?", gajiID).
Updates(map[string]interface{}{
"status_bayar":  "paid",
"tanggal_bayar": now,
}).Error
}
