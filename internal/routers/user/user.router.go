package user

import (
	"github.com/gin-gonic/gin"
	"github.com/vuvanthanhtb/go-ecommerce-backend-api/internal/wire"
)

type UserRouter struct {
}

func (ur *UserRouter) InitUserRouter(rg *gin.RouterGroup) {
	// Public router
	// this is non-dependency
	// ure := repo.NewUserRepository()
	// us := service.NewUserService(ure)
	// userHandlerNonDependency := controller.NewUserController(us)

	// DI use wire go
	userController, _ := wire.InitUserRouterHandler()
	userRouterPublic := rg.Group("/user")
	{
		userRouterPublic.GET("/register", userController.Register)
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
