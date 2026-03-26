package main

import (
	"log"
	"os"

	"hris-api/config"
	"hris-api/internal/model"
	"hris-api/internal/router"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env")
	}

	config.ConnectDB()

	config.DB.AutoMigrate(
		&model.User{},
		&model.Divisi{},
		&model.Jabatan{},
		&model.DataUser{},
		&model.Absensi{},
		&model.Cuti{},
		&model.Penggajian{},
	)

	r := router.SetupRouter()

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on port %s", port)
	r.Run(":" + port)
}
