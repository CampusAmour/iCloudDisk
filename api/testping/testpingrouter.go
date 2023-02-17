package testping

import (
	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {
	r := router.Group("/ping")
	r.GET("/", pingFunc)
}
