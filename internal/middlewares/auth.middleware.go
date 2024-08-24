package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/vuvanthanhtb/go-ecommerce-backend-api/pkg/response"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "valid-token" {
			response.ErrorResponse(c, response.ErrInvalidToken)
			c.Abort()
		}
		c.Next()
	}
}
