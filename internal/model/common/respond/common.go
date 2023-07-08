package respond

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Respond struct {
	Code int         `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data,omitempty"`
}

var (
	SUCCESS = 0
	ERROR   = 7
)

func Result(code int, msg string, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, Respond{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func Success(c *gin.Context) {
	Result(SUCCESS, "操作成功", nil, c)
}

func SuccessWithMsg(msg string, c *gin.Context) {
	Result(SUCCESS, msg, nil, c)
}

func SuccessWitData(msg string, data interface{}, c *gin.Context) {
	Result(SUCCESS, msg, data, c)
}

func Failure(c *gin.Context) {
	Result(ERROR, "操作失败", nil, c)
}

func FailureWithMsg(msg string, c *gin.Context) {
	Result(ERROR, msg, nil, c)
}

func FailureWithData(msg string, data interface{}, c *gin.Context) {
	Result(ERROR, msg, data, c)
}
