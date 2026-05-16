package repository

import (
"hris-api/config"
"hris-api/internal/model"

"github.com/google/uuid"
)

func CreateUser(user *model.User) error {
user.UID = uuid.New().String()
return config.DB.Create(user).Error
}

func FindUserByEmail(email string) (*model.User, error) {
var user model.User
err := config.DB.Where("email = ?", email).First(&user).Error
return &user, err
}

func FindUserByID(uid string) (*model.User, error) {
var user model.User
err := config.DB.Preload("DataUser").Where("uid = ?", uid).First(&user).Error
return &user, err
}

func GetAllUsers() ([]model.User, error) {
var users []model.User
err := config.DB.Preload("DataUser").Find(&users).Error
return users, err
}

func UpdateUser(uid string, updates map[string]interface{}) error {
return config.DB.Model(&model.User{}).Where("uid = ?", uid).Updates(updates).Error
}

func DeleteUser(uid string) error {
return config.DB.Delete(&model.User{}, "uid = ?", uid).Error
}
