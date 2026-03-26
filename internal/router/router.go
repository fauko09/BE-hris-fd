package router

import (
"net/http"

"hris-api/internal/handler"
"hris-api/middleware"

"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
r := gin.Default()

r.StaticFS("/uploads", http.Dir("storage/uploads"))

api := r.Group("/api/v1")

auth := api.Group("/auth")
{
auth.POST("/register", handler.Register)
auth.POST("/login", handler.Login)
}

p := api.Group("/")
p.Use(middleware.AuthMiddleware())
{
p.GET("/users", handler.GetAllUsers)
p.GET("/users/:id", handler.GetUserByID)
p.DELETE("/users/:id", handler.DeleteUser)

p.POST("/employees/:uid/photo", handler.UploadPhoto)
p.POST("/employees/:uid/data", handler.CreateDataUser)
p.GET("/employees/:uid/data", handler.GetDataUserByUID)
p.GET("/employees", handler.GetAllDataUsers)

p.POST("/absensi/clock-in", handler.ClockIn)
p.POST("/absensi/clock-out", handler.ClockOut)
p.GET("/absensi/me", handler.GetMyAbsensi)
p.GET("/absensi/today", handler.GetAbsensiToday)

p.POST("/cuti", handler.AjukanCuti)
p.GET("/cuti/me", handler.GetMyCuti)
p.GET("/cuti", handler.GetAllCuti)
p.PUT("/cuti/:id/approve", handler.ApproveCuti)

p.POST("/penggajian", handler.CreatePenggajian)
p.GET("/penggajian/me", handler.GetMyGaji)
p.GET("/penggajian", handler.GetAllPenggajian)
p.PUT("/penggajian/:id/bayar", handler.BayarGaji)

p.GET("/divisi", handler.GetAllDivisi)
p.POST("/divisi", handler.CreateDivisi)
p.PUT("/divisi/:id", handler.UpdateDivisi)
p.DELETE("/divisi/:id", handler.DeleteDivisi)

p.GET("/jabatan", handler.GetAllJabatan)
p.POST("/jabatan", handler.CreateJabatan)
p.PUT("/jabatan/:id", handler.UpdateJabatan)
p.DELETE("/jabatan/:id", handler.DeleteJabatan)
}

return r
}