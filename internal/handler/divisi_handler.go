package handler

import (
"hris-api/internal/model"
"hris-api/internal/repository"
"hris-api/utils"

"github.com/gin-gonic/gin"
"github.com/google/uuid"
)

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

func GetAllDivisi(c *gin.Context) {
list, err := repository.GetAllDivisi()
if err != nil {
utils.InternalError(c, err.Error())
return
}
utils.OK(c, "Success", list)
}

func UpdateDivisi(c *gin.Context) {
id, err := uuid.Parse(c.Param("id"))
if err != nil {
utils.BadRequest(c, "Invalid UUID")
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

func DeleteDivisi(c *gin.Context) {
id, err := uuid.Parse(c.Param("id"))
if err != nil {
utils.BadRequest(c, "Invalid UUID")
return
}
if err := repository.DeleteDivisi(id); err != nil {
utils.InternalError(c, err.Error())
return
}
utils.OK(c, "Divisi berhasil dihapus", nil)
}
