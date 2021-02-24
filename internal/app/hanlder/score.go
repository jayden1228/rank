package hanlder

import (
	"net/http"
	"rank/internal/app/service"

	"github.com/gin-gonic/gin"
)

func AddScore(c *gin.Context) {
	type req struct {
		UserId int64
		MatchId int64
		Score int
	}

	var body req

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10001,
			"message": err.Error(),
		})
		return
	}

	if err := service.AddScore(body.UserId, body.MatchId, body.Score); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10003,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"message": "ok",
	})}
