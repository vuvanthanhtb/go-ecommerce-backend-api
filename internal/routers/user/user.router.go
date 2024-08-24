package user

import (
	"github.com/gin-gonic/gin"
	"github.com/vuvanthanhtb/go-ecommerce-backend-api/internal/controller"
	"github.com/vuvanthanhtb/go-ecommerce-backend-api/internal/repo"
	"github.com/vuvanthanhtb/go-ecommerce-backend-api/internal/service"
)

type UserRouter struct {
}

func (ur *UserRouter) InitUserRouter(rg *gin.RouterGroup) {
	// Public router
	// this is non-dependency
	ure := repo.NewUserRepository()
	us := service.NewUserService(ure)
	userHandlerNonDependency := controller.NewUserController(us)
	userRouterPublic := rg.Group("/user")
	{
		userRouterPublic.GET("/register", userHandlerNonDependency.Register)
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
