package handler

import (
"hris-api/internal/model"
"hris-api/internal/repository"
"hris-api/utils"

"github.com/gin-gonic/gin"
"github.com/google/uuid"
)

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

func GetAllJabatan(c *gin.Context) {
list, err := repository.GetAllJabatan()
if err != nil {
utils.InternalError(c, err.Error())
return
}
utils.OK(c, "Success", list)
}

func UpdateJabatan(c *gin.Context) {
id, err := uuid.Parse(c.Param("id"))
if err != nil {
utils.BadRequest(c, "Invalid UUID")
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

func DeleteJabatan(c *gin.Context) {
id, err := uuid.Parse(c.Param("id"))
if err != nil {
utils.BadRequest(c, "Invalid UUID")
return
}
if err := repository.DeleteJabatan(id); err != nil {
utils.InternalError(c, err.Error())
return
}
utils.OK(c, "Jabatan berhasil dihapus", nil)
}
