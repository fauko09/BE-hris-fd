package handler

import (
"strconv"

"hris-api/internal/repository"
"hris-api/utils"

"github.com/gin-gonic/gin"
"github.com/google/uuid"
)

func ClockIn(c *gin.Context) {
uid, err := uuid.Parse(c.GetString("uid"))
if err != nil {
utils.Unauthorized(c, "Token tidak valid")
return
}
var req struct {
Lokasi string `json:"lokasi"`
}
c.ShouldBindJSON(&req)
absensi, err := repository.ClockIn(uid, req.Lokasi)
if err != nil {
utils.BadRequest(c, err.Error())
return
}
utils.Created(c, "Clock in berhasil", absensi)
}

func ClockOut(c *gin.Context) {
uid, err := uuid.Parse(c.GetString("uid"))
if err != nil {
utils.Unauthorized(c, "Token tidak valid")
return
}
var req struct {
Lokasi string `json:"lokasi"`
}
c.ShouldBindJSON(&req)
absensi, err := repository.ClockOut(uid, req.Lokasi)
if err != nil {
utils.BadRequest(c, err.Error())
return
}
utils.OK(c, "Clock out berhasil", absensi)
}

func GetMyAbsensi(c *gin.Context) {
uid, _ := uuid.Parse(c.GetString("uid"))
bulan, _ := strconv.Atoi(c.Query("bulan"))
tahun, _ := strconv.Atoi(c.Query("tahun"))
list, err := repository.GetAbsensiByUID(uid, bulan, tahun)
if err != nil {
utils.InternalError(c, err.Error())
return
}
utils.OK(c, "Success", list)
}

func GetAbsensiToday(c *gin.Context) {
list, err := repository.GetAllAbsensiToday()
if err != nil {
utils.InternalError(c, err.Error())
return
}
utils.OK(c, "Success", list)
}
