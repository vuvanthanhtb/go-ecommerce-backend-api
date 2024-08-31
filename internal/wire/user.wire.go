//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/vuvanthanhtb/go-ecommerce-backend-api/internal/controller"
	"github.com/vuvanthanhtb/go-ecommerce-backend-api/internal/repo"
	"github.com/vuvanthanhtb/go-ecommerce-backend-api/internal/service"
)

func InitUserRouterHandler() (*controller.UserController, error) {
	wire.Build(
		repo.NewUserRepository,
		repo.NewAuthRepository,
		service.NewUserService,
		controller.NewUserController,
	)

	return new(controller.UserController), nil
}
