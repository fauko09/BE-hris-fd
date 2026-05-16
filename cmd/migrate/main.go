package main

import (
	"fmt"
	"log"
	"os"

	"hris-api/internal/model"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatal("Failed to connect:", err)
	}

	log.Println("Database connected!")

	err = db.AutoMigrate(
		&model.Divisi{},
		&model.Jabatan{},
		&model.User{},
		&model.DataUser{},
		&model.Absensi{},
		&model.Cuti{},
		&model.Penggajian{},
	)
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("Migration success! Semua tabel berhasil dibuat.")
}
