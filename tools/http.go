package tools

import "github.com/gin-gonic/gin"

type Data struct {
	Code int
	Data interface{}
}

func ResponseSuccess(ctx *gin.Context, code int, data interface{}) {
	ctx.JSON(code, gin.H{
		"data": Data{
			Code: code,
			Data: data,
		},
	})
}

func ResponseFailed(ctx *gin.Context, code int, err error) {
	ctx.JSON(code, gin.H{
		"data": Data{
			Code: code,
			Data: err.Error(),
		},
	})
}
