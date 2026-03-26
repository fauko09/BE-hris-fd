package handler

import (
"time"

"hris-api/internal/model"
"hris-api/internal/repository"
"hris-api/utils"

"github.com/gin-gonic/gin"
"github.com/google/uuid"
)

func CreateDataUser(c *gin.Context) {
uidStr := c.Param("uid")
uid, err := uuid.Parse(uidStr)
if err != nil {
utils.BadRequest(c, "Invalid UUID")
return
}
var req model.DataUserRequest
if err := c.ShouldBindJSON(&req); err != nil {
utils.BadRequest(c, err.Error())
return
}
data := &model.DataUser{
UID:                 uid,
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
id, _ := uuid.Parse(req.DivisiID)
data.DivisiID = &id
}
if req.JabatanID != "" {
id, _ := uuid.Parse(req.JabatanID)
data.JabatanID = &id
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

func GetDataUserByUID(c *gin.Context) {
uid, err := uuid.Parse(c.Param("uid"))
if err != nil {
utils.BadRequest(c, "Invalid UUID")
return
}
data, err := repository.FindDataUserByUID(uid)
if err != nil {
utils.NotFound(c, "Data user tidak ditemukan")
return
}
utils.OK(c, "Success", data)
}

func GetAllDataUsers(c *gin.Context) {
data, err := repository.GetAllDataUsers()
if err != nil {
utils.InternalError(c, err.Error())
return
}
utils.OK(c, "Success", data)
}
