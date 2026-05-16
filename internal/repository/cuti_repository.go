package repository

import (
"time"

"hris-api/config"
"hris-api/internal/model"

"github.com/google/uuid"
)

func CreateCuti(cuti *model.Cuti) error {
cuti.CutiID = uuid.New().String()
return config.DB.Create(cuti).Error
}

func GetCutiByUID(uid string) ([]model.Cuti, error) {
var list []model.Cuti
err := config.DB.Where("uid = ?", uid).Order("created_at DESC").Find(&list).Error
return list, err
}

func GetAllCuti(status string) ([]model.Cuti, error) {
var list []model.Cuti
query := config.DB.Preload("User")
if status != "" {
query = query.Where("status = ?", status)
}
err := query.Order("created_at DESC").Find(&list).Error
return list, err
}

func ApproveCuti(cutiID, approverUID string, req model.ApproveCutiRequest) error {
now := time.Now()
updates := map[string]interface{}{
"status":      req.Status,
"catatan_hr":  req.CatatanHR,
"approved_by": approverUID,
"approved_at": now,
}
return config.DB.Model(&model.Cuti{}).Where("cuti_id = ?", cutiID).Updates(updates).Error
}
