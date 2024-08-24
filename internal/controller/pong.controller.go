package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PongController struct {
}

func NewPongController() *PongController {
	return &PongController{}
}

func (p *PongController) Pong(c *gin.Context) {
	name := c.DefaultQuery("name", "Thanh")
	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{
		"name": name,
		"uid":  uid,
	})
}
