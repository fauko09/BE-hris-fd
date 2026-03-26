package repository

import (
"hris-api/config"
"hris-api/internal/model"

"github.com/google/uuid"
)

func CreateJabatan(j *model.Jabatan) error {
j.JabatanID = uuid.New()
return config.DB.Create(j).Error
}

func GetAllJabatan() ([]model.Jabatan, error) {
var list []model.Jabatan
err := config.DB.Find(&list).Error
return list, err
}

func FindJabatanByID(id uuid.UUID) (*model.Jabatan, error) {
var j model.Jabatan
err := config.DB.Where("jabatan_id = ?", id).First(&j).Error
return &j, err
}

func UpdateJabatan(id uuid.UUID, updates map[string]interface{}) error {
return config.DB.Model(&model.Jabatan{}).Where("jabatan_id = ?", id).Updates(updates).Error
}

func DeleteJabatan(id uuid.UUID) error {
return config.DB.Delete(&model.Jabatan{}, "jabatan_id = ?", id).Error
}
