package handler

import (
	"strconv"

	"hris-api/internal/repository"
	"hris-api/utils"

	"github.com/gin-gonic/gin"
)

// @Summary      Clock In
// @Description  Karyawan melakukan clock in
// @Tags         Absensi
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      201   {object}  map[string]interface{}
// @Failure      400   {object}  map[string]interface{}
// @Router       /absensi/clock-in [post]
func ClockIn(c *gin.Context) {
	uid := c.GetString("uid")
	if uid == "" {
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

// @Summary      Clock Out
// @Description  Karyawan melakukan clock out
// @Tags         Absensi
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200   {object}  map[string]interface{}
// @Failure      400   {object}  map[string]interface{}
// @Router       /absensi/clock-out [post]
func ClockOut(c *gin.Context) {
	uid := c.GetString("uid")
	if uid == "" {
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

// @Summary      Get absensi saya
// @Description  Ambil riwayat absensi user yang sedang login
// @Tags         Absensi
// @Produce      json
// @Security     BearerAuth
// @Param        bulan  query     int  false  "Bulan (1-12)"
// @Param        tahun  query     int  false  "Tahun (contoh: 2026)"
// @Success      200    {object}  map[string]interface{}
// @Failure      500    {object}  map[string]interface{}
// @Router       /absensi/me [get]
func GetMyAbsensi(c *gin.Context) {
	uid := c.GetString("uid")
	bulan, _ := strconv.Atoi(c.Query("bulan"))
	tahun, _ := strconv.Atoi(c.Query("tahun"))
	list, err := repository.GetAbsensiByUID(uid, bulan, tahun)
	if err != nil {
		utils.InternalError(c, err.Error())
		return
	}
	utils.OK(c, "Success", list)
}

// @Summary      Get absensi hari ini
// @Description  Ambil semua absensi hari ini (admin)
// @Tags         Absensi
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /absensi/today [get]
func GetAbsensiToday(c *gin.Context) {
	list, err := repository.GetAllAbsensiToday()
	if err != nil {
		utils.InternalError(c, err.Error())
		return
	}
	utils.OK(c, "Success", list)
}
