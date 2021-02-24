package app

import (
	"rank/internal/app/hanlder"

	"github.com/gin-gonic/gin"
)

func registerRouters(router *gin.Engine) {
	router.GET("/healthz", hanlder.HandleHealthz)
	router.POST("/match/score", hanlder.AddScore)
}
