package handler

import (
	"hris-api/internal/model"
	"hris-api/internal/repository"
	"hris-api/utils"

	"github.com/gin-gonic/gin"
)

// @Summary      Buat jabatan
// @Description  Tambah jabatan baru
// @Tags         Jabatan
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        body  body      model.JabatanRequest  true  "Jabatan Request"
// @Success      201   {object}  map[string]interface{}
// @Failure      400   {object}  map[string]interface{}
// @Router       /jabatan [post]
func CreateJabatan(c *gin.Context) {
	var req model.JabatanRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}
	j := &model.Jabatan{NamaJabatan: req.NamaJabatan}
	if err := repository.CreateJabatan(j); err != nil {
		utils.InternalError(c, err.Error())
		return
	}
	utils.Created(c, "Jabatan berhasil dibuat", j)
}

// @Summary      Get semua jabatan
// @Description  Ambil daftar semua jabatan
// @Tags         Jabatan
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /jabatan [get]
func GetAllJabatan(c *gin.Context) {
	list, err := repository.GetAllJabatan()
	if err != nil {
		utils.InternalError(c, err.Error())
		return
	}
	utils.OK(c, "Success", list)
}

// @Summary      Update jabatan
// @Description  Update nama jabatan berdasarkan ID
// @Tags         Jabatan
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id    path      string                true  "Jabatan ID"
// @Param        body  body      model.JabatanRequest  true  "Jabatan Request"
// @Success      200   {object}  map[string]interface{}
// @Failure      500   {object}  map[string]interface{}
// @Router       /jabatan/{id} [put]
func UpdateJabatan(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		utils.BadRequest(c, "Invalid ID")
		return
	}
	var req model.JabatanRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}
	if err := repository.UpdateJabatan(id, map[string]interface{}{"nama_jabatan": req.NamaJabatan}); err != nil {
		utils.InternalError(c, err.Error())
		return
	}
	utils.OK(c, "Jabatan berhasil diupdate", nil)
}

// @Summary      Delete jabatan
// @Description  Hapus jabatan berdasarkan ID
// @Tags         Jabatan
// @Produce      json
// @Security     BearerAuth
// @Param        id  path      string  true  "Jabatan ID"
// @Success      200  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /jabatan/{id} [delete]
func DeleteJabatan(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		utils.BadRequest(c, "Invalid ID")
		return
	}
	if err := repository.DeleteJabatan(id); err != nil {
		utils.InternalError(c, err.Error())
		return
	}
	utils.OK(c, "Jabatan berhasil dihapus", nil)
}
