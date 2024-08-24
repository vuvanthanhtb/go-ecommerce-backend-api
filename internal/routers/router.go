package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/vuvanthanhtb/go-ecommerce-backend-api/internal/middlewares"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	// use the middlewares
	r.Use(middlewares.AuthMiddleware())

	return r
}
