package handler

import (
"time"

"hris-api/internal/model"
"hris-api/internal/repository"
"hris-api/utils"

"github.com/gin-gonic/gin"
"github.com/google/uuid"
)

func AjukanCuti(c *gin.Context) {
uid, _ := uuid.Parse(c.GetString("uid"))
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

func GetMyCuti(c *gin.Context) {
uid, _ := uuid.Parse(c.GetString("uid"))
list, err := repository.GetCutiByUID(uid)
if err != nil {
utils.InternalError(c, err.Error())
return
}
utils.OK(c, "Success", list)
}

func GetAllCuti(c *gin.Context) {
status := c.Query("status")
list, err := repository.GetAllCuti(status)
if err != nil {
utils.InternalError(c, err.Error())
return
}
utils.OK(c, "Success", list)
}

func ApproveCuti(c *gin.Context) {
approverUID, _ := uuid.Parse(c.GetString("uid"))
cutiID, err := uuid.Parse(c.Param("id"))
if err != nil {
utils.BadRequest(c, "Invalid UUID")
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
