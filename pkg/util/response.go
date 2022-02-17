package util

import "github.com/gin-gonic/gin"

type GIN struct {
	context *gin.Context
}

func (g *GIN) SendResponse(httpStatus int, data interface{}) {
	g.context.JSON(httpStatus, data)
}
