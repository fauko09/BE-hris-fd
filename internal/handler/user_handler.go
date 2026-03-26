package handler

import (
"os"
"time"

"hris-api/internal/model"
"hris-api/internal/repository"
"hris-api/utils"

"github.com/gin-gonic/gin"
"github.com/golang-jwt/jwt/v5"
"github.com/google/uuid"
"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
var req model.RegisterRequest
if err := c.ShouldBindJSON(&req); err != nil {
utils.BadRequest(c, err.Error())
return
}
hashed, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
user := &model.User{
Fullname: req.Fullname,
NIK:      req.NIK,
Email:    req.Email,
Password: string(hashed),
}
if err := repository.CreateUser(user); err != nil {
utils.BadRequest(c, "Email atau NIK sudah terdaftar")
return
}
utils.Created(c, "Registrasi berhasil", user)
}

func Login(c *gin.Context) {
var req model.LoginRequest
if err := c.ShouldBindJSON(&req); err != nil {
utils.BadRequest(c, err.Error())
return
}
user, err := repository.FindUserByEmail(req.Email)
if err != nil {
utils.Unauthorized(c, "Email tidak ditemukan")
return
}
if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
utils.Unauthorized(c, "Password salah")
return
}
token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
"uid": user.UID.String(),
"exp": time.Now().Add(24 * time.Hour).Unix(),
})
tokenStr, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
utils.OK(c, "Login berhasil", gin.H{"token": tokenStr, "user": user})
}

func GetAllUsers(c *gin.Context) {
users, err := repository.GetAllUsers()
if err != nil {
utils.InternalError(c, err.Error())
return
}
utils.OK(c, "Success", users)
}

func GetUserByID(c *gin.Context) {
id, err := uuid.Parse(c.Param("id"))
if err != nil {
utils.BadRequest(c, "Invalid UUID")
return
}
user, err := repository.FindUserByID(id)
if err != nil {
utils.NotFound(c, "User tidak ditemukan")
return
}
utils.OK(c, "Success", user)
}

func DeleteUser(c *gin.Context) {
id, err := uuid.Parse(c.Param("id"))
if err != nil {
utils.BadRequest(c, "Invalid UUID")
return
}
if err := repository.DeleteUser(id); err != nil {
utils.InternalError(c, err.Error())
return
}
utils.OK(c, "User berhasil dihapus", nil)
}
