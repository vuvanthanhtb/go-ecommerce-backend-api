package user

import "github.com/gin-gonic/gin"

type UserRouter struct {
}

func (ur *UserRouter) InitUserRouter(rg *gin.RouterGroup) {
	// Public router
	userRouterPublic := rg.Group("/user")
	{
		userRouterPublic.GET("/register")
		userRouterPublic.GET("/otp")
	}

	// Private router
	userRouterPrivate := rg.Group("/user")
	// userRouterPrivate.Use(limiter())
	// userRouterPrivate.Use(permission())
	{
		userRouterPrivate.GET("/get_info")
	}
}
