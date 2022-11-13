package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 错误信息
type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Ok 返回成功
// @msg: 返回消息
// @data: 返回成功的数据
func Ok(c *gin.Context, data ...interface{}) {
	resp := Response{
		Message: "OK",
	}

	if len(data) > 0 {
		resp.Data = data[0]
	}

	c.JSON(http.StatusOK, resp)
}

// Err 返回失败
// @param: httpCode 错误码
// @param: msg 错误消息
// @param: data 额外信息（可选）
func Err(c *gin.Context, httpCode int, msg string, data ...interface{}) {

	resp := Response{
		Message: msg,
	}

	if len(data) > 0 {
		resp.Data = data[0]
	}

	c.JSON(httpCode, resp)
}
