package repository

import (
"hris-api/config"
"hris-api/internal/model"

"github.com/google/uuid"
)

func CreateDataUser(data *model.DataUser) error {
data.DaID = uuid.New()
return config.DB.Create(data).Error
}

func FindDataUserByUID(uid uuid.UUID) (*model.DataUser, error) {
var data model.DataUser
err := config.DB.Preload("Divisi").Preload("Jabatan").Where("uid = ?", uid).First(&data).Error
return &data, err
}

func UpdateDataUser(uid uuid.UUID, updates map[string]interface{}) error {
return config.DB.Model(&model.DataUser{}).Where("uid = ?", uid).Updates(updates).Error
}

func GetAllDataUsers() ([]model.DataUser, error) {
var data []model.DataUser
err := config.DB.Preload("User").Preload("Divisi").Preload("Jabatan").Find(&data).Error
return data, err
}
