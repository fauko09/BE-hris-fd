package handler

import (
	"time"

	"hris-api/internal/model"
	"hris-api/internal/repository"
	"hris-api/utils"

	"github.com/gin-gonic/gin"
)

// @Summary      Ajukan cuti
// @Description  Karyawan mengajukan permohonan cuti
// @Tags         Cuti
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        body  body      model.CutiRequest  true  "Cuti Request"
// @Success      201   {object}  map[string]interface{}
// @Failure      400   {object}  map[string]interface{}
// @Router       /cuti [post]
func AjukanCuti(c *gin.Context) {
	uid := c.GetString("uid")
	var req model.CutiRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}
	mulai, _ := time.Parse("2006-01-02", req.TanggalMulai)
	akhir, _ := time.Parse("2006-01-02", req.TanggalAkhir)
	totalHari := int(akhir.Sub(mulai).Hours()/24) + 1
	cuti := &model.Cuti{
		UID:          uid,
		TipeCuti:     req.TipeCuti,
		TanggalMulai: mulai,
		TanggalAkhir: akhir,
		TotalHari:    totalHari,
		Alasan:       req.Alasan,
		Status:       "pending",
	}
	if err := repository.CreateCuti(cuti); err != nil {
		utils.InternalError(c, err.Error())
		return
	}
	utils.Created(c, "Pengajuan cuti berhasil", cuti)
}

// @Summary      Get cuti saya
// @Description  Ambil riwayat cuti user yang sedang login
// @Tags         Cuti
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /cuti/me [get]
func GetMyCuti(c *gin.Context) {
	uid := c.GetString("uid")
	list, err := repository.GetCutiByUID(uid)
	if err != nil {
		utils.InternalError(c, err.Error())
		return
	}
	utils.OK(c, "Success", list)
}

// @Summary      Get semua cuti
// @Description  Ambil semua pengajuan cuti (admin)
// @Tags         Cuti
// @Produce      json
// @Security     BearerAuth
// @Param        status  query     string  false  "Filter status (pending/approved/rejected)"
// @Success      200     {object}  map[string]interface{}
// @Failure      500     {object}  map[string]interface{}
// @Router       /cuti [get]
func GetAllCuti(c *gin.Context) {
	status := c.Query("status")
	list, err := repository.GetAllCuti(status)
	if err != nil {
		utils.InternalError(c, err.Error())
		return
	}
	utils.OK(c, "Success", list)
}

// @Summary      Approve/Reject cuti
// @Description  HR menyetujui atau menolak pengajuan cuti
// @Tags         Cuti
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id    path      string                    true  "Cuti ID"
// @Param        body  body      model.ApproveCutiRequest  true  "Approve Request"
// @Success      200   {object}  map[string]interface{}
// @Failure      500   {object}  map[string]interface{}
// @Router       /cuti/{id}/approve [put]
func ApproveCuti(c *gin.Context) {
	approverUID := c.GetString("uid")
	cutiID := c.Param("id")
	if cutiID == "" {
		utils.BadRequest(c, "Invalid ID")
		return
	}
	var req model.ApproveCutiRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}
	if err := repository.ApproveCuti(cutiID, approverUID, req); err != nil {
		utils.InternalError(c, err.Error())
		return
	}
	utils.OK(c, "Status cuti berhasil diupdate", nil)
}
