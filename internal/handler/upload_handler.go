package handler

import (
"fmt"
"os"
"path/filepath"
"time"

"hris-api/internal/repository"
"hris-api/utils"

"github.com/gin-gonic/gin"
"github.com/google/uuid"
)

func UploadPhoto(c *gin.Context) {
uid, err := uuid.Parse(c.Param("uid"))
if err != nil {
utils.BadRequest(c, "Invalid UUID")
return
}
file, err := c.FormFile("photo")
if err != nil {
utils.BadRequest(c, "File tidak ditemukan")
return
}
ext := filepath.Ext(file.Filename)
allowed := map[string]bool{".jpg": true, ".jpeg": true, ".png": true}
if !allowed[ext] {
utils.BadRequest(c, "Format file harus jpg/jpeg/png")
return
}
if file.Size > 2*1024*1024 {
utils.BadRequest(c, "Ukuran file maksimal 2MB")
return
}
filename := fmt.Sprintf("%s_%d%s", uid.String(), time.Now().Unix(), ext)
savePath := filepath.Join("storage", "uploads", "photos", filename)
os.MkdirAll(filepath.Dir(savePath), os.ModePerm)
if err := c.SaveUploadedFile(file, savePath); err != nil {
utils.InternalError(c, "Gagal menyimpan file")
return
}
photoURL := fmt.Sprintf("/uploads/photos/%s", filename)
if err := repository.UpdateUser(uid, map[string]interface{}{"photo_url": photoURL}); err != nil {
utils.InternalError(c, err.Error())
return
}
utils.OK(c, "Foto berhasil diupload", gin.H{"photo_url": photoURL})
}
