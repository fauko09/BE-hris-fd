package handler

import (
	"os"
	"time"

	"hris-api/internal/model"
	"hris-api/internal/repository"
	"hris-api/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// @Summary      Register user baru
// @Description  Membuat akun user baru
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        body  body      model.RegisterRequest  true  "Register Request"
// @Success      201   {object}  map[string]interface{}
// @Failure      400   {object}  map[string]interface{}
// @Router       /auth/register [post]
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

// @Summary      Login user
// @Description  Login dan mendapatkan JWT token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        body  body      model.LoginRequest  true  "Login Request"
// @Success      200   {object}  map[string]interface{}
// @Failure      401   {object}  map[string]interface{}
// @Router       /auth/login [post]
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
		"uid": user.UID,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	})
	tokenStr, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	utils.OK(c, "Login berhasil", gin.H{"token": tokenStr, "user": user})
}

// @Summary      Get semua user
// @Description  Ambil daftar semua user
// @Tags         Users
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /users [get]
func GetAllUsers(c *gin.Context) {
	users, err := repository.GetAllUsers()
	if err != nil {
		utils.InternalError(c, err.Error())
		return
	}
	utils.OK(c, "Success", users)
}

// @Summary      Get user by ID
// @Description  Ambil data user berdasarkan ID
// @Tags         Users
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      string  true  "User ID"
// @Success      200  {object}  map[string]interface{}
// @Failure      404  {object}  map[string]interface{}
// @Router       /users/{id} [get]
func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		utils.BadRequest(c, "Invalid ID")
		return
	}
	user, err := repository.FindUserByID(id)
	if err != nil {
		utils.NotFound(c, "User tidak ditemukan")
		return
	}
	utils.OK(c, "Success", user)
}

// @Summary      Delete user
// @Description  Hapus user berdasarkan ID
// @Tags         Users
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      string  true  "User ID"
// @Success      200  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /users/{id} [delete]
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		utils.BadRequest(c, "Invalid ID")
		return
	}
	if err := repository.DeleteUser(id); err != nil {
		utils.InternalError(c, err.Error())
		return
	}
	utils.OK(c, "User berhasil dihapus", nil)
}
