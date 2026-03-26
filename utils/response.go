package utils

import "github.com/gin-gonic/gin"

type Response struct {
Success bool        `json:"success"`
Message string      `json:"message"`
Data    interface{} `json:"data,omitempty"`
Error   string      `json:"error,omitempty"`
}

func OK(c *gin.Context, message string, data interface{}) {
c.JSON(200, Response{Success: true, Message: message, Data: data})
}

func Created(c *gin.Context, message string, data interface{}) {
c.JSON(201, Response{Success: true, Message: message, Data: data})
}

func BadRequest(c *gin.Context, err string) {
c.JSON(400, Response{Success: false, Error: err})
}

func Unauthorized(c *gin.Context, err string) {
c.JSON(401, Response{Success: false, Error: err})
}

func NotFound(c *gin.Context, err string) {
c.JSON(404, Response{Success: false, Error: err})
}

func InternalError(c *gin.Context, err string) {
c.JSON(500, Response{Success: false, Error: err})
}
