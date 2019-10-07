package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 通用返回模型
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// SendResp 返回自定义数据
func SendResp(c *gin.Context, httpCode, code int, message string, data interface{}) {
	c.JSON(httpCode, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

// SendFailResp 返回失败请求
func SendFailResp(c *gin.Context, message string) {
	c.JSON(http.StatusOK, Response{
		Code:    -1,
		Message: message,
	})
}

// SendSuccessResp 返回成功请求
func SendSuccessResp(c *gin.Context, message string) {
	c.JSON(http.StatusOK, Response{
		Code:    1,
		Message: message,
	})
}
