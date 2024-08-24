package admin

import "github.com/gin-gonic/gin"

type AdminRouter struct {
}

func (ur *AdminRouter) InitAdminRouter(rg *gin.RouterGroup) {
	// Public router
	adminRouterPublic := rg.Group("/admin")
	{
		adminRouterPublic.POST("/login")
	}

	// Private router
	adminRouterPrivate := rg.Group("/admin")
	// adminRouterPrivate.Use(limiter())
	// adminRouterPrivate.Use(Authen())
	// adminRouterPrivate.Use(Permission())
	{
		adminRouterPrivate.POST("/active_user")
	}
}
