package admin

import "github.com/gin-gonic/gin"

type UserRouter struct {
}

func (ur *UserRouter) InitUserRouter(rg *gin.RouterGroup) {
	// Public router
	// userRouterPublic := rg.Group("/admin/user")
	// {
	// 	userRouterPublic.GET("/register")
	// 	userRouterPublic.GET("/otp")
	// }

	// Private router
	userRouterPrivate := rg.Group("/admin/user")
	// userRouterPrivate.Use(limiter())
	// userRouterPrivate.Use(permission())
	{
		userRouterPrivate.POST("/active_user")
	}
}
