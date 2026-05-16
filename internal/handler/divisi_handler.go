package handler

import (
	"hris-api/internal/model"
	"hris-api/internal/repository"
	"hris-api/utils"

	"github.com/gin-gonic/gin"
)

// @Summary      Buat divisi
// @Description  Tambah divisi baru
// @Tags         Divisi
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        body  body      model.DivisiRequest  true  "Divisi Request"
// @Success      201   {object}  map[string]interface{}
// @Failure      400   {object}  map[string]interface{}
// @Router       /divisi [post]
func CreateDivisi(c *gin.Context) {
	var req model.DivisiRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}
	d := &model.Divisi{NameDivisi: req.NameDivisi}
	if err := repository.CreateDivisi(d); err != nil {
		utils.InternalError(c, err.Error())
		return
	}
	utils.Created(c, "Divisi berhasil dibuat", d)
}

// @Summary      Get semua divisi
// @Description  Ambil daftar semua divisi
// @Tags         Divisi
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /divisi [get]
func GetAllDivisi(c *gin.Context) {
	list, err := repository.GetAllDivisi()
	if err != nil {
		utils.InternalError(c, err.Error())
		return
	}
	utils.OK(c, "Success", list)
}

// @Summary      Update divisi
// @Description  Update nama divisi berdasarkan ID
// @Tags         Divisi
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id    path      string               true  "Divisi ID"
// @Param        body  body      model.DivisiRequest  true  "Divisi Request"
// @Success      200   {object}  map[string]interface{}
// @Failure      500   {object}  map[string]interface{}
// @Router       /divisi/{id} [put]
func UpdateDivisi(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		utils.BadRequest(c, "Invalid ID")
		return
	}
	var req model.DivisiRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}
	if err := repository.UpdateDivisi(id, map[string]interface{}{"name_divisi": req.NameDivisi}); err != nil {
		utils.InternalError(c, err.Error())
		return
	}
	utils.OK(c, "Divisi berhasil diupdate", nil)
}

// @Summary      Delete divisi
// @Description  Hapus divisi berdasarkan ID
// @Tags         Divisi
// @Produce      json
// @Security     BearerAuth
// @Param        id  path      string  true  "Divisi ID"
// @Success      200  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /divisi/{id} [delete]
func DeleteDivisi(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		utils.BadRequest(c, "Invalid ID")
		return
	}
	if err := repository.DeleteDivisi(id); err != nil {
		utils.InternalError(c, err.Error())
		return
	}
	utils.OK(c, "Divisi berhasil dihapus", nil)
}
