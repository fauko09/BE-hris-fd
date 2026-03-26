package repository

import (
"hris-api/config"
"hris-api/internal/model"

"github.com/google/uuid"
)

func CreateDivisi(d *model.Divisi) error {
d.DivisiID = uuid.New()
return config.DB.Create(d).Error
}

func GetAllDivisi() ([]model.Divisi, error) {
var list []model.Divisi
err := config.DB.Find(&list).Error
return list, err
}

func FindDivisiByID(id uuid.UUID) (*model.Divisi, error) {
var d model.Divisi
err := config.DB.Where("divisi_id = ?", id).First(&d).Error
return &d, err
}

func UpdateDivisi(id uuid.UUID, updates map[string]interface{}) error {
return config.DB.Model(&model.Divisi{}).Where("divisi_id = ?", id).Updates(updates).Error
}

func DeleteDivisi(id uuid.UUID) error {
return config.DB.Delete(&model.Divisi{}, "divisi_id = ?", id).Error
}
