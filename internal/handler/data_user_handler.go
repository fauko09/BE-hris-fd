package handler

import (
	"time"

	"hris-api/internal/model"
	"hris-api/internal/repository"
	"hris-api/utils"

	"github.com/gin-gonic/gin"
)

// @Summary      Buat data user
// @Description  Simpan data detail karyawan
// @Tags         Employees
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        uid   path      string                true  "User ID"
// @Param        body  body      model.DataUserRequest  true  "Data User Request"
// @Success      201   {object}  map[string]interface{}
// @Failure      400   {object}  map[string]interface{}
// @Router       /employees/{uid}/data [post]
func CreateDataUser(c *gin.Context) {
	uidStr := c.Param("uid")
	if uidStr == "" {
		utils.BadRequest(c, "Invalid UID")
		return
	}
	var req model.DataUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}
	data := &model.DataUser{
		UID:                 uidStr,
		Status:              req.Status,
		AlamatTinggal:       req.AlamatTinggal,
		TipeTinggal:         req.TipeTinggal,
		DomisiliTinggal:     req.DomisiliTinggal,
		ProvTinggal:         req.ProvTinggal,
		AlamatKTP:           req.AlamatKTP,
		DomisiliKTP:         req.DomisiliKTP,
		ProvKTP:             req.ProvKTP,
		DomisiliLahir:       req.DomisiliLahir,
		NoNPWP:              req.NoNPWP,
		NoBPJS:              req.NoBPJS,
		IsBPJSTK:            req.IsBPJSTK,
		NoAsuransi:          req.NoAsuransi,
		JenisAsuransi:       req.JenisAsuransi,
		NamaAsuransi:        req.NamaAsuransi,
		StatusPerkawinan:    req.StatusPerkawinan,
		NomorTelepon:        req.NomorTelepon,
		NomorTeleponSecond:  req.NomorTeleponSecond,
		NomorTeleponDarurat: req.NomorTeleponDarurat,
		NoRekening:          req.NoRekening,
		TipeBank:            req.TipeBank,
	}
	if req.DivisiID != "" {
		data.DivisiID = &req.DivisiID
	}
	if req.JabatanID != "" {
		data.JabatanID = &req.JabatanID
	}
	if req.TanggalLahir != "" {
		t, _ := time.Parse("2006-01-02", req.TanggalLahir)
		data.TanggalLahir = &t
	}
	if req.TanggalMasuk != "" {
		t, _ := time.Parse("2006-01-02", req.TanggalMasuk)
		data.TanggalMasuk = &t
	}
	if err := repository.CreateDataUser(data); err != nil {
		utils.InternalError(c, err.Error())
		return
	}
	utils.Created(c, "Data user berhasil disimpan", data)
}

// @Summary      Get data user by UID
// @Description  Ambil data detail karyawan berdasarkan UID
// @Tags         Employees
// @Produce      json
// @Security     BearerAuth
// @Param        uid  path      string  true  "User ID"
// @Success      200  {object}  map[string]interface{}
// @Failure      404  {object}  map[string]interface{}
// @Router       /employees/{uid}/data [get]
func GetDataUserByUID(c *gin.Context) {
	uid := c.Param("uid")
	if uid == "" {
		utils.BadRequest(c, "Invalid UID")
		return
	}
	data, err := repository.FindDataUserByUID(uid)
	if err != nil {
		utils.NotFound(c, "Data user tidak ditemukan")
		return
	}
	utils.OK(c, "Success", data)
}

// @Summary      Get semua data karyawan
// @Description  Ambil semua data detail karyawan
// @Tags         Employees
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /employees [get]
func GetAllDataUsers(c *gin.Context) {
	data, err := repository.GetAllDataUsers()
	if err != nil {
		utils.InternalError(c, err.Error())
		return
	}
	utils.OK(c, "Success", data)
}
