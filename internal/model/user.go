package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	UID       string         `gorm:"type:char(36);primaryKey" json:"uid"`
	Fullname  string         `gorm:"type:varchar(100);not null" json:"fullname"`
	NIK       string         `gorm:"type:varchar(16);unique;not null" json:"nik"`
	Email     string         `gorm:"type:varchar(255);unique;not null" json:"email"`
	Password  string         `gorm:"type:varchar(255)" json:"-"`
	PhotoURL  *string        `gorm:"type:varchar(255)" json:"photo_url"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	DataUser  *DataUser      `gorm:"foreignKey:UID" json:"data_user,omitempty"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.UID == "" {
		u.UID = uuid.New().String()
	}
	return nil
}

type RegisterRequest struct {
	Fullname string `json:"fullname" binding:"required"`
	NIK      string `json:"nik" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UpdateUserRequest struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email" binding:"omitempty,email"`
}
