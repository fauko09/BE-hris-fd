package handler

import (
	"hris-api/internal/model"
	"hris-api/internal/repository"
	"hris-api/utils"

	"github.com/gin-gonic/gin"
)

// @Summary      Buat data penggajian
// @Description  Tambah data gaji karyawan
// @Tags         Penggajian
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        body  body      model.PenggajianRequest  true  "Penggajian Request"
// @Success      201   {object}  map[string]interface{}
// @Failure      400   {object}  map[string]interface{}
// @Router       /penggajian [post]
func CreatePenggajian(c *gin.Context) {
	var req model.PenggajianRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}
	p := &model.Penggajian{
		UID:        req.UID,
		Periode:    req.Periode,
		GajiPokok:  req.GajiPokok,
		Tunjangan:  req.Tunjangan,
		Potongan:   req.Potongan,
		Keterangan: req.Keterangan,
	}
	if err := repository.CreatePenggajian(p); err != nil {
		utils.InternalError(c, err.Error())
		return
	}
	utils.Created(c, "Data gaji berhasil disimpan", p)
}

// @Summary      Get gaji saya
// @Description  Ambil riwayat gaji user yang sedang login
// @Tags         Penggajian
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /penggajian/me [get]
func GetMyGaji(c *gin.Context) {
	uid := c.GetString("uid")
	list, err := repository.GetPenggajianByUID(uid)
	if err != nil {
		utils.InternalError(c, err.Error())
		return
	}
	utils.OK(c, "Success", list)
}

// @Summary      Get semua penggajian
// @Description  Ambil semua data penggajian (admin)
// @Tags         Penggajian
// @Produce      json
// @Security     BearerAuth
// @Param        periode  query     string  false  "Filter periode (contoh: 2026-05)"
// @Success      200      {object}  map[string]interface{}
// @Failure      500      {object}  map[string]interface{}
// @Router       /penggajian [get]
func GetAllPenggajian(c *gin.Context) {
	periode := c.Query("periode")
	list, err := repository.GetAllPenggajian(periode)
	if err != nil {
		utils.InternalError(c, err.Error())
		return
	}
	utils.OK(c, "Success", list)
}

// @Summary      Bayar gaji
// @Description  Update status gaji menjadi paid
// @Tags         Penggajian
// @Produce      json
// @Security     BearerAuth
// @Param        id  path      string  true  "Gaji ID"
// @Success      200  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /penggajian/{id}/bayar [put]
func BayarGaji(c *gin.Context) {
	gajiID := c.Param("id")
	if gajiID == "" {
		utils.BadRequest(c, "Invalid ID")
		return
	}
	if err := repository.BayarGaji(gajiID); err != nil {
		utils.InternalError(c, err.Error())
		return
	}
	utils.OK(c, "Gaji berhasil dibayar", nil)
}
