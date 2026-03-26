package handler

import (
"hris-api/internal/model"
"hris-api/internal/repository"
"hris-api/utils"

"github.com/gin-gonic/gin"
"github.com/google/uuid"
)

func CreatePenggajian(c *gin.Context) {
var req model.PenggajianRequest
if err := c.ShouldBindJSON(&req); err != nil {
utils.BadRequest(c, err.Error())
return
}
uid, err := uuid.Parse(req.UID)
if err != nil {
utils.BadRequest(c, "UID tidak valid")
return
}
p := &model.Penggajian{
UID:        uid,
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

func GetMyGaji(c *gin.Context) {
uid, _ := uuid.Parse(c.GetString("uid"))
list, err := repository.GetPenggajianByUID(uid)
if err != nil {
utils.InternalError(c, err.Error())
return
}
utils.OK(c, "Success", list)
}

func GetAllPenggajian(c *gin.Context) {
periode := c.Query("periode")
list, err := repository.GetAllPenggajian(periode)
if err != nil {
utils.InternalError(c, err.Error())
return
}
utils.OK(c, "Success", list)
}

func BayarGaji(c *gin.Context) {
gajiID, err := uuid.Parse(c.Param("id"))
if err != nil {
utils.BadRequest(c, "Invalid UUID")
return
}
if err := repository.BayarGaji(gajiID); err != nil {
utils.InternalError(c, err.Error())
return
}
utils.OK(c, "Gaji berhasil dibayar", nil)
}
